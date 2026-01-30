#!/usr/bin/env python3
"""
Script para importar arquivos Markdown da pasta mfi_pos como artigos (lectures)
"""

import os
import sys
import json
import requests
from pathlib import Path
from datetime import datetime
import re

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
MFI_FOLDER = "/home/user/plenya/mfi_pos"

# Função para limpar e formatar o título do arquivo
def format_title(filename):
    """Remove .md e formata o nome do arquivo como título"""
    title = filename.replace(".md", "")
    title = title.replace("_", " ")
    # Remove múltiplos espaços
    title = re.sub(r'\s+', ' ', title)
    return title.strip()

# Função para criar artigo via API
def create_article(title, full_content, auth_token):
    """Cria um artigo via API POST"""

    url = f"{API_BASE_URL}/articles"

    # Preparar payload
    payload = {
        "title": title,
        "authors": "Equipe MFI",
        "journal": "Pos Graduacao MFI",
        "publishDate": "2024-01-01T00:00:00Z",  # Data padrão
        "articleType": "lecture",
        "specialty": "Medicina Funcional",
        "fullContent": full_content,
        "abstract": f"Conteúdo da aula: {title[:200]}...",  # Primeiros 200 chars como abstract
    }

    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {auth_token}"
    }

    try:
        response = requests.post(url, json=payload, headers=headers)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"❌ Erro ao criar artigo '{title}': {e}")
        if hasattr(e.response, 'text'):
            print(f"   Resposta: {e.response.text[:500]}")
        return None

# Função para obter token de autenticação
def get_auth_token():
    """Obtém token de autenticação do admin"""
    # Fazer login como admin
    url = f"{API_BASE_URL}/auth/login"

    payload = {
        "email": "import@plenya.com",
        "password": "Import@123456"
    }

    try:
        response = requests.post(url, json=payload)
        response.raise_for_status()
        data = response.json()
        return data.get("accessToken")
    except Exception as e:
        print(f"❌ Erro ao obter token: {e}")
        if hasattr(e, 'response') and e.response is not None:
            print(f"   Status: {e.response.status_code}")
            print(f"   Resposta: {e.response.text[:500]}")
        sys.exit(1)

def main():
    print("=" * 80)
    print("📚 IMPORTAÇÃO DE LECTURES MFI - Medicina Funcional Integrativa")
    print("=" * 80)
    print()

    # 1. Obter token de autenticação
    print("🔐 Obtendo token de autenticação...")
    auth_token = get_auth_token()
    print(f"✅ Token obtido com sucesso\n")

    # 2. Listar todos os arquivos .md
    md_files = sorted(Path(MFI_FOLDER).glob("*.md"))
    total_files = len(md_files)

    print(f"📂 Encontrados {total_files} arquivos Markdown\n")

    # 3. Importar cada arquivo
    success_count = 0
    error_count = 0

    for idx, md_file in enumerate(md_files, 1):
        filename = md_file.name
        title = format_title(filename)

        print(f"[{idx}/{total_files}] {filename}")
        print(f"   Título: {title}")

        # Ler conteúdo completo do arquivo
        try:
            with open(md_file, 'r', encoding='utf-8') as f:
                full_content = f.read()

            # Verificar se arquivo está vazio
            if not full_content.strip():
                print(f"   ⚠️  Arquivo vazio, pulando...")
                error_count += 1
                continue

            print(f"   Tamanho: {len(full_content):,} caracteres")

            # Criar artigo via API
            result = create_article(title, full_content, auth_token)

            if result:
                article_id = result.get('id', 'unknown')
                print(f"   ✅ Criado com sucesso (ID: {article_id})")
                success_count += 1
            else:
                print(f"   ❌ Falha ao criar")
                error_count += 1

        except Exception as e:
            print(f"   ❌ Erro ao processar: {e}")
            error_count += 1

        print()

    # 4. Resumo final
    print("=" * 80)
    print("📊 RESUMO DA IMPORTAÇÃO")
    print("=" * 80)
    print(f"Total de arquivos: {total_files}")
    print(f"✅ Importados com sucesso: {success_count}")
    print(f"❌ Erros: {error_count}")
    print(f"Taxa de sucesso: {(success_count/total_files*100):.1f}%")
    print()

    if success_count > 0:
        print("🎉 Importação concluída! Acesse http://localhost:3000/articles para ver.")

    return 0 if error_count == 0 else 1

if __name__ == "__main__":
    sys.exit(main())
