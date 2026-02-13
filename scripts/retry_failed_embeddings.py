#!/usr/bin/env python3
"""
Script para buscar PDFs de artigos com embeddings falhados.

Tenta múltiplas fontes:
1. Unpaywall API (Open Access)
2. PubMed Central
3. Sci-Hub (último recurso)

Atualiza o banco de dados automaticamente quando encontra PDF.
"""

import os
import sys
import time
import requests
from typing import Optional, Dict, List, Tuple
from dataclasses import dataclass
from pathlib import Path

try:
    import psycopg2
    from psycopg2.extras import RealDictCursor
except ImportError:
    print("❌ Erro: psycopg2 não instalado")
    print("Instalando psycopg2...")
    os.system("pip install psycopg2-binary requests")
    import psycopg2
    from psycopg2.extras import RealDictCursor


@dataclass
class Article:
    """Representa um artigo científico."""
    id: str
    title: str
    doi: Optional[str]
    pmid: Optional[str]


class PDFRetriever:
    """Busca PDFs de artigos científicos em múltiplas fontes."""

    UNPAYWALL_EMAIL = "research@plenya.com"
    RATE_LIMIT_SECONDS = 2
    TIMEOUT = 30

    def __init__(self, upload_dir: str):
        self.upload_dir = Path(upload_dir)
        self.upload_dir.mkdir(parents=True, exist_ok=True)
        self.session = requests.Session()
        self.session.headers.update({
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        })

    def try_unpaywall(self, doi: str) -> Optional[str]:
        """Tenta buscar PDF via Unpaywall API."""
        if not doi:
            return None

        try:
            url = f"https://api.unpaywall.org/v2/{doi}"
            params = {'email': self.UNPAYWALL_EMAIL}

            response = self.session.get(url, params=params, timeout=self.TIMEOUT)
            response.raise_for_status()

            data = response.json()

            # Tenta best_oa_location primeiro
            if data.get('best_oa_location'):
                pdf_url = data['best_oa_location'].get('url_for_pdf')
                if pdf_url:
                    return pdf_url

                # Se não tem url_for_pdf, tenta url genérica
                url = data['best_oa_location'].get('url')
                if url and url.endswith('.pdf'):
                    return url

            # Tenta oa_locations alternativas
            for location in data.get('oa_locations', []):
                pdf_url = location.get('url_for_pdf')
                if pdf_url:
                    return pdf_url

            return None

        except Exception as e:
            print(f"    ⚠️  Unpaywall error: {e}")
            return None

    def try_pubmed_central(self, pmid: str) -> Optional[str]:
        """Tenta buscar PDF via PubMed Central."""
        if not pmid:
            return None

        # Remove prefixo PMC se existir
        pmid_clean = pmid.replace('PMC', '')

        try:
            # Tenta formato PMC direto
            pdf_url = f"https://www.ncbi.nlm.nih.gov/pmc/articles/PMC{pmid_clean}/pdf/"

            response = self.session.head(pdf_url, timeout=self.TIMEOUT, allow_redirects=True)

            if response.status_code == 200:
                return pdf_url

            return None

        except Exception as e:
            print(f"    ⚠️  PubMed Central error: {e}")
            return None

    def try_scihub(self, doi: str) -> Optional[str]:
        """Tenta buscar PDF via Sci-Hub (último recurso)."""
        if not doi:
            return None

        scihub_domains = [
            'https://sci-hub.se',
            'https://sci-hub.st',
            'https://sci-hub.ru'
        ]

        for domain in scihub_domains:
            try:
                url = f"{domain}/{doi}"
                response = self.session.get(url, timeout=self.TIMEOUT, allow_redirects=True)

                if response.status_code == 200 and 'pdf' in response.headers.get('Content-Type', '').lower():
                    return response.url

            except Exception:
                continue

        return None

    def download_pdf(self, url: str, article_id: str) -> bool:
        """Baixa PDF e salva no diretório de uploads."""
        try:
            response = self.session.get(url, timeout=self.TIMEOUT, stream=True)
            response.raise_for_status()

            # Verifica se é realmente um PDF
            content_type = response.headers.get('Content-Type', '').lower()
            if 'pdf' not in content_type and not url.endswith('.pdf'):
                print(f"    ⚠️  URL não retornou PDF: {content_type}")
                return False

            # Salva o arquivo
            pdf_path = self.upload_dir / f"{article_id}.pdf"

            with open(pdf_path, 'wb') as f:
                for chunk in response.iter_content(chunk_size=8192):
                    if chunk:
                        f.write(chunk)

            # Verifica tamanho mínimo (1KB)
            if pdf_path.stat().st_size < 1024:
                pdf_path.unlink()
                print(f"    ⚠️  PDF muito pequeno (< 1KB)")
                return False

            return True

        except Exception as e:
            print(f"    ⚠️  Download error: {e}")
            return False

    def find_and_download(self, article: Article) -> Tuple[bool, Optional[str]]:
        """
        Busca e baixa PDF do artigo.

        Returns:
            (sucesso, fonte)
        """
        # 1. Tenta Unpaywall
        if article.doi:
            print(f"    🔍 Trying Unpaywall...")
            pdf_url = self.try_unpaywall(article.doi)
            if pdf_url:
                if self.download_pdf(pdf_url, article.id):
                    return True, "Unpaywall"
                time.sleep(1)

        # 2. Tenta PubMed Central
        if article.pmid:
            print(f"    🔍 Trying PubMed Central...")
            pdf_url = self.try_pubmed_central(article.pmid)
            if pdf_url:
                if self.download_pdf(pdf_url, article.id):
                    return True, "PubMed Central"
                time.sleep(1)

        # 3. Tenta Sci-Hub (último recurso)
        if article.doi:
            print(f"    🔍 Trying Sci-Hub...")
            pdf_url = self.try_scihub(article.doi)
            if pdf_url:
                if self.download_pdf(pdf_url, article.id):
                    return True, "Sci-Hub"

        return False, None


class DatabaseManager:
    """Gerencia conexão e queries do PostgreSQL."""

    def __init__(self, host: str, user: str, password: str, database: str):
        self.conn = psycopg2.connect(
            host=host,
            user=user,
            password=password,
            database=database,
            cursor_factory=RealDictCursor
        )
        self.conn.autocommit = False

    def get_failed_articles(self) -> List[Article]:
        """Busca artigos com embeddings falhados."""
        query = """
            SELECT DISTINCT a.id, a.title, a.doi, a.pm_id as pmid, a.publish_date
            FROM articles a
            JOIN embedding_queue eq ON a.id = eq.entity_id
            WHERE eq.status = 'failed'
              AND a.deleted_at IS NULL
              AND (a.doi IS NOT NULL OR a.pm_id IS NOT NULL)
            ORDER BY a.publish_date DESC NULLS LAST
        """

        with self.conn.cursor() as cur:
            cur.execute(query)
            rows = cur.fetchall()

            return [
                Article(
                    id=row['id'],
                    title=row['title'],
                    doi=row['doi'],
                    pmid=row['pmid']
                )
                for row in rows
            ]

    def update_article_pdf(self, article_id: str) -> bool:
        """Atualiza article com internal_link e reseta embedding_queue."""
        try:
            with self.conn.cursor() as cur:
                # Update article
                cur.execute(
                    """
                    UPDATE articles
                    SET internal_link = %s
                    WHERE id = %s
                    """,
                    (f"/uploads/articles/{article_id}.pdf", article_id)
                )

                # Reset embedding queue
                cur.execute(
                    """
                    UPDATE embedding_queue
                    SET status = 'pending',
                        retry_count = 0,
                        error_message = NULL
                    WHERE entity_id = %s
                    """,
                    (article_id,)
                )

                self.conn.commit()
                return True

        except Exception as e:
            self.conn.rollback()
            print(f"    ❌ Database error: {e}")
            return False

    def close(self):
        """Fecha conexão."""
        self.conn.close()


def print_banner():
    """Imprime banner do script."""
    print("\n" + "="*80)
    print("📚 RETRY FAILED EMBEDDINGS - PDF RETRIEVER")
    print("="*80 + "\n")


def print_summary(found: List[Tuple[Article, str]], not_found: List[Article], errors: List[Tuple[Article, str]]):
    """Imprime resumo da execução."""
    total = len(found) + len(not_found) + len(errors)

    print("\n" + "="*80)
    print("📊 SUMMARY")
    print("="*80)
    print(f"\n✅ Found PDFs:     {len(found)}/{total} ({len(found)/total*100:.1f}%)")
    print(f"❌ Not found:      {len(not_found)}/{total} ({len(not_found)/total*100:.1f}%)")
    print(f"⚠️  Errors:         {len(errors)}/{total} ({len(errors)/total*100:.1f}%)")

    if found:
        print("\n" + "-"*80)
        print("✅ FOUND PDFs:")
        print("-"*80)
        for article, source in found:
            print(f"  • {article.title[:70]}... [{source}]")

    if not_found:
        print("\n" + "-"*80)
        print("❌ NOT FOUND (manual search needed):")
        print("-"*80)
        for article in not_found:
            identifiers = []
            if article.doi:
                identifiers.append(f"DOI: {article.doi}")
            if article.pmid:
                identifiers.append(f"PMID: {article.pmid}")

            print(f"\n  📄 {article.title}")
            print(f"     {' | '.join(identifiers)}")

            # Links para busca manual
            if article.doi:
                print(f"     🔗 https://doi.org/{article.doi}")
            if article.pmid:
                print(f"     🔗 https://pubmed.ncbi.nlm.nih.gov/{article.pmid}/")

    if errors:
        print("\n" + "-"*80)
        print("⚠️  ERRORS:")
        print("-"*80)
        for article, error in errors:
            print(f"  • {article.title[:70]}... - {error}")

    print("\n" + "="*80 + "\n")


def main():
    """Função principal."""
    print_banner()

    # Configurações
    DB_CONFIG = {
        'host': 'localhost',  # Porta 5432 exposta do Docker
        'user': 'plenya_user',
        'password': 'plenya_dev_password',
        'database': 'plenya_db'
    }

    UPLOAD_DIR = '/home/user/plenya/uploads/articles'

    # Inicializa
    print("🔌 Connecting to database...")
    try:
        db = DatabaseManager(**DB_CONFIG)
    except Exception as e:
        print(f"❌ Failed to connect: {e}")
        return 1

    print("✅ Connected to PostgreSQL\n")

    # Busca artigos falhados
    print("🔍 Fetching failed articles...")
    articles = db.get_failed_articles()
    total = len(articles)

    if total == 0:
        print("✅ No failed articles found!\n")
        db.close()
        return 0

    print(f"📋 Found {total} failed articles\n")

    # Inicializa retriever
    retriever = PDFRetriever(UPLOAD_DIR)

    # Listas de resultados
    found: List[Tuple[Article, str]] = []
    not_found: List[Article] = []
    errors: List[Tuple[Article, str]] = []

    # Processa cada artigo
    for idx, article in enumerate(articles, 1):
        print(f"\n[{idx}/{total}] 📄 {article.title[:60]}...")

        if article.doi:
            print(f"    DOI: {article.doi}")
        if article.pmid:
            print(f"    PMID: {article.pmid}")

        try:
            success, source = retriever.find_and_download(article)

            if success:
                # Atualiza banco de dados
                if db.update_article_pdf(article.id):
                    print(f"    ✅ Found on {source} - Database updated")
                    found.append((article, source))
                else:
                    print(f"    ⚠️  PDF downloaded but database update failed")
                    errors.append((article, "Database update failed"))
            else:
                print(f"    ❌ Not found")
                not_found.append(article)

        except Exception as e:
            print(f"    ❌ Error: {e}")
            errors.append((article, str(e)))

        # Rate limiting
        if idx < total:
            time.sleep(PDFRetriever.RATE_LIMIT_SECONDS)

    # Fecha conexão
    db.close()

    # Imprime resumo
    print_summary(found, not_found, errors)

    return 0


if __name__ == '__main__':
    sys.exit(main())
