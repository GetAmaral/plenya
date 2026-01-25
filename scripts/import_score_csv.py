#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Script de Importação CSV → Score System Database

Importa dados de "Escore Plenya SPL v1.3.csv" para as tabelas:
- score_groups
- score_subgroups
- score_items
- score_levels

Autor: Claude Code
Data: 2026-01-24
"""

import csv
import re
import sys
import argparse
import logging
from typing import Optional, Tuple, Dict
from uuid import UUID

try:
    import psycopg2
    from psycopg2.extras import RealDictCursor
except ImportError:
    print("ERRO: psycopg2 não instalado. Execute: pip3 install psycopg2-binary")
    sys.exit(1)


# ============================================================================
# CONFIGURAÇÃO DE LOGGING
# ============================================================================

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s [%(levelname)s] %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S'
)
logger = logging.getLogger(__name__)


# ============================================================================
# FUNÇÕES DE PARSING
# ============================================================================

def parse_float(value: str) -> float:
    """Converte string para float, retorna 0 se vazio/inválido"""
    if not value or not value.strip():
        return 0.0
    try:
        return float(value.strip().replace(',', '.'))
    except ValueError:
        return 0.0


def normalize_decimal(text: str) -> str:
    """Normaliza vírgula para ponto em decimais"""
    return text.replace(',', '.')


def parse_range(text: str) -> Tuple[str, Optional[str], Optional[str]]:
    """
    Parse de range de valores retornando (operator, lower_limit, upper_limit)

    Operadores:
        '=' - Igualdade (valor exato ou texto)
        '>' - Maior que
        '>=' - Maior ou igual
        '<' - Menor que
        '<=' - Menor ou igual
        'between' - Entre dois valores (range)

    Exemplos:
        "≥7,0" -> ('>=', "7.0", None)
        "6,0–6,9" -> ('between', "6.0", "6.9")
        "<10" -> ('<', None, "10")
        ">-16" -> ('>', "-16", None)
        "-2.5 a -2.0" -> ('between', "-2.5", "-2.0")
        "Remissão" -> ('=', None, None)
        "5.5" -> ('=', "5.5", "5.5")
    """
    if not text or not text.strip():
        return ('=', None, None)

    # Normalizar
    text = normalize_decimal(text)
    text = text.strip()

    # Se tem "ou", pegar apenas primeira parte
    if " ou " in text:
        text = text.split(" ou ")[0].strip()

    # Padrão numérico: aceita +/- opcional, dígitos, ponto decimal opcional
    NUM_PATTERN = r'[+\-]?[0-9]+\.?[0-9]*'

    # Padrão: >=X ou ≥X (com número positivo ou negativo)
    match = re.match(rf'^(≥|>=)\s*({NUM_PATTERN})', text)
    if match:
        return ('>=', match.group(2), None)

    # Padrão: >X (mas não >=, com número positivo ou negativo)
    match = re.match(rf'^>\s*({NUM_PATTERN})', text)
    if match:
        return ('>', match.group(1), None)

    # Padrão: <=X ou ≤X (com número positivo ou negativo)
    match = re.match(rf'^(≤|<=)\s*({NUM_PATTERN})', text)
    if match:
        return ('<=', None, match.group(2))

    # Padrão: <X (mas não <=, com número positivo ou negativo)
    match = re.match(rf'^<\s*({NUM_PATTERN})', text)
    if match:
        return ('<', None, match.group(1))

    # Padrão: X a Y (com palavra "a")
    match = re.match(rf'^({NUM_PATTERN})\s+a\s+({NUM_PATTERN})', text)
    if match:
        return ('between', match.group(1), match.group(2))

    # Padrão: X–Y, X—Y, X-Y (en-dash, em-dash, ou hífen normal)
    # IMPORTANTE: Testar en-dash/em-dash primeiro, depois hífen
    # En-dash: –, Em-dash: —, Hífen normal: -
    match = re.match(rf'^({NUM_PATTERN})\s*[–—\-]\s*({NUM_PATTERN})', text)
    if match:
        lower = match.group(1)
        upper = match.group(2)
        # Validar que não é um número negativo sendo interpretado errado
        # Exemplo: "-5" não deve virar range, mas "-5 - 10" ou "5-10" sim
        if lower and upper:
            return ('between', lower, upper)

    # Número único exato (positivo ou negativo)
    match = re.match(rf'^({NUM_PATTERN})$', text)
    if match:
        num = match.group(1)
        return ('=', num, num)

    # Se não tem número, é texto puro (equals)
    if not re.search(r'[0-9]', text):
        return ('=', None, None)

    # Não conseguiu parsear - retornar equals por padrão
    logger.debug(f"Não conseguiu parsear range: '{text}' - usando operator '='")
    return ('=', None, None)


def parse_unit_conversion(text: str) -> Tuple[Optional[str], Optional[str]]:
    """
    Separa unit e conversion de string "unit | conversion"

    Exemplo:
        "mg/dL | mg/dL × 59,48 = μmol/L" -> ("mg/dL", "mg/dL × 59,48 = μmol/L")
    """
    if not text or " | " not in text:
        return (None, None)

    parts = text.split(" | ", 1)
    unit = parts[0].strip() if parts[0] else None
    conversion = parts[1].strip() if len(parts) > 1 and parts[1] else None

    return (unit, conversion)


# ============================================================================
# FUNÇÕES DE BANCO DE DADOS
# ============================================================================

class ScoreImporter:
    def __init__(self, conn):
        self.conn = conn
        self.cursor = conn.cursor(cursor_factory=RealDictCursor)

        # Cache de IDs para evitar re-queries
        self.group_cache: Dict[str, UUID] = {}
        self.subgroup_cache: Dict[Tuple[str, UUID], UUID] = {}

        # Contadores
        self.stats = {
            'groups': 0,
            'subgroups': 0,
            'items': 0,
            'items_with_points': 0,
            'items_without_points': 0,
            'levels': 0,
            'rows_processed': 0,
            'rows_skipped': 0,
        }

    def insert_group(self, name: str, order: int) -> UUID:
        """Insere ou retorna ID de grupo existente"""
        # Verificar cache
        if name in self.group_cache:
            return self.group_cache[name]

        # Verificar se já existe
        self.cursor.execute(
            "SELECT id FROM score_groups WHERE name = %s AND deleted_at IS NULL",
            (name,)
        )
        row = self.cursor.fetchone()

        if row:
            group_id = row['id']
        else:
            # Inserir novo
            self.cursor.execute(
                """
                INSERT INTO score_groups (name, "order")
                VALUES (%s, %s)
                RETURNING id
                """,
                (name, order)
            )
            group_id = self.cursor.fetchone()['id']
            self.stats['groups'] += 1
            logger.info(f"✅ Grupo criado: {name} (ordem {order})")

        self.group_cache[name] = group_id
        return group_id

    def insert_subgroup(self, name: str, group_id: UUID, order: int) -> UUID:
        """Insere ou retorna ID de subgrupo existente"""
        cache_key = (name, group_id)

        # Verificar cache
        if cache_key in self.subgroup_cache:
            return self.subgroup_cache[cache_key]

        # Verificar se já existe
        self.cursor.execute(
            """
            SELECT id FROM score_subgroups
            WHERE name = %s AND group_id = %s AND deleted_at IS NULL
            """,
            (name, group_id)
        )
        row = self.cursor.fetchone()

        if row:
            subgroup_id = row['id']
        else:
            # Inserir novo
            self.cursor.execute(
                """
                INSERT INTO score_subgroups (name, "order", group_id)
                VALUES (%s, %s, %s)
                RETURNING id
                """,
                (name, order, group_id)
            )
            subgroup_id = self.cursor.fetchone()['id']
            self.stats['subgroups'] += 1
            logger.info(f"  ✅ Subgrupo criado: {name} (ordem {order})")

        self.subgroup_cache[cache_key] = subgroup_id
        return subgroup_id

    def insert_item(
        self,
        name: str,
        unit: Optional[str],
        unit_conversion: Optional[str],
        points: float,
        order: int,
        subgroup_id: UUID,
        parent_item_id: Optional[UUID]
    ) -> UUID:
        """Insere item de score"""
        self.cursor.execute(
            """
            INSERT INTO score_items
            (name, unit, unit_conversion, points, "order", subgroup_id, parent_item_id)
            VALUES (%s, %s, %s, %s, %s, %s, %s)
            RETURNING id
            """,
            (name, unit, unit_conversion, points, order, subgroup_id, parent_item_id)
        )
        item_id = self.cursor.fetchone()['id']
        self.stats['items'] += 1

        if points > 0:
            self.stats['items_with_points'] += 1
            parent_info = f" (filho de {parent_item_id})" if parent_item_id else ""
            logger.info(f"    ✅ Item: {name} ({points} pts){parent_info}")
        else:
            self.stats['items_without_points'] += 1
            logger.debug(f"    ℹ️  Item organizador: {name} (0 pts)")

        return item_id

    def insert_level(
        self,
        level: int,
        name: str,
        operator: str,
        lower_limit: Optional[str],
        upper_limit: Optional[str],
        item_id: UUID
    ):
        """Insere nível de score"""
        self.cursor.execute(
            """
            INSERT INTO score_levels
            (level, name, operator, lower_limit, upper_limit, item_id)
            VALUES (%s, %s, %s, %s, %s, %s)
            """,
            (level, name, operator, lower_limit, upper_limit, item_id)
        )
        self.stats['levels'] += 1

    def validate_import(self):
        """Valida dados importados"""
        logger.info("\n" + "="*70)
        logger.info("VALIDAÇÃO DE IMPORTAÇÃO")
        logger.info("="*70)

        # 1. Count total
        self.cursor.execute("""
            SELECT
                (SELECT COUNT(*) FROM score_groups) as groups,
                (SELECT COUNT(*) FROM score_subgroups) as subgroups,
                (SELECT COUNT(*) FROM score_items) as items,
                (SELECT COUNT(*) FROM score_levels) as levels
        """)
        counts = self.cursor.fetchone()
        logger.info(f"\n📊 Registros no banco:")
        logger.info(f"  - Grupos: {counts['groups']}")
        logger.info(f"  - Subgrupos: {counts['subgroups']}")
        logger.info(f"  - Itens: {counts['items']}")
        logger.info(f"  - Níveis: {counts['levels']}")

        # 2. Itens sem 6 níveis (que têm points > 0)
        self.cursor.execute("""
            SELECT i.id, i.name, i.points, COUNT(l.id) as nivel_count
            FROM score_items i
            LEFT JOIN score_levels l ON l.item_id = i.id
            WHERE i.points > 0
            GROUP BY i.id, i.name, i.points
            HAVING COUNT(l.id) != 6
            ORDER BY i.name
            LIMIT 10
        """)
        problemas = self.cursor.fetchall()

        if problemas:
            logger.warning(f"\n⚠️  {len(problemas)} itens COM PONTOS sem exatamente 6 níveis:")
            for p in problemas[:10]:
                logger.warning(f"  - {p['name']}: {p['nivel_count']} níveis (esperado: 6)")
        else:
            logger.info("\n✅ Todos os itens com pontos têm exatamente 6 níveis")

        # 3. Níveis duplicados
        self.cursor.execute("""
            SELECT item_id, level, COUNT(*) as count
            FROM score_levels
            GROUP BY item_id, level
            HAVING COUNT(*) > 1
        """)
        duplicados = self.cursor.fetchall()

        if duplicados:
            logger.warning(f"\n⚠️  {len(duplicados)} níveis duplicados encontrados!")
            for d in duplicados:
                logger.warning(f"  - Item {d['item_id']}, Nível {d['level']}: {d['count']} ocorrências")
        else:
            logger.info("✅ Sem níveis duplicados")

        # 4. Itens hierárquicos
        self.cursor.execute("""
            SELECT COUNT(*) as count
            FROM score_items
            WHERE parent_item_id IS NOT NULL
        """)
        hierarquicos = self.cursor.fetchone()['count']
        logger.info(f"\n📁 Itens hierárquicos (com parent): {hierarquicos}")

        # 5. Amostra de itens
        self.cursor.execute("""
            SELECT
                g.name as grupo,
                s.name as subgrupo,
                i.name as item,
                i.points,
                COUNT(l.id) as niveis
            FROM score_groups g
            JOIN score_subgroups s ON s.group_id = g.id
            JOIN score_items i ON i.subgroup_id = s.id
            LEFT JOIN score_levels l ON l.item_id = i.id
            GROUP BY g.name, g."order", s.name, s."order", i.name, i."order", i.points
            ORDER BY g."order", s."order", i."order"
            LIMIT 5
        """)
        amostra = self.cursor.fetchall()

        logger.info("\n📋 Amostra de itens importados:")
        for a in amostra:
            logger.info(f"  {a['grupo']} → {a['subgrupo']} → {a['item']} ({a['points']} pts, {a['niveis']} níveis)")


# ============================================================================
# PROCESSAMENTO DO CSV
# ============================================================================

def process_csv(filepath: str, importer: ScoreImporter):
    """Processa arquivo CSV e popula banco de dados"""

    logger.info(f"📂 Lendo CSV: {filepath}")

    # Estado global (sticky)
    current_group_name = None
    current_group_id = None
    current_subgroup_name = None
    current_subgroup_id = None
    last_item_col3_id = None  # Para hierarquia

    # Contadores de ordem
    group_order = 0
    subgroup_order = {}  # dict[group_id] -> int
    item_order = {}      # dict[subgroup_id] -> int

    # Ler CSV com encoding UTF-8
    with open(filepath, 'r', encoding='utf-8') as f:
        reader = csv.reader(f, delimiter=';')

        # Pular header
        next(reader)

        for row_num, row in enumerate(reader, start=2):
            importer.stats['rows_processed'] += 1

            # Garantir 11 colunas (padding se necessário)
            while len(row) < 11:
                row.append('')

            col1_grupo = row[0].strip()
            col2_subgrupo = row[1].strip()
            col3_item = row[2].strip()
            col4_subitem_ou_unit = row[3].strip()
            col5_points = parse_float(row[4])
            niveis = [row[5], row[6], row[7], row[8], row[9], row[10]]

            # ================================================================
            # 1. PROCESSAR GRUPO (sticky)
            # ================================================================
            if col1_grupo:
                current_group_name = col1_grupo
                group_order += 1
                current_group_id = importer.insert_group(col1_grupo, group_order)
                subgroup_order[current_group_id] = 0

            # ================================================================
            # 2. PROCESSAR SUBGRUPO (sticky)
            # ================================================================
            if col2_subgrupo:
                current_subgroup_name = col2_subgrupo
                subgroup_order[current_group_id] += 1
                current_subgroup_id = importer.insert_subgroup(
                    col2_subgrupo,
                    current_group_id,
                    subgroup_order[current_group_id]
                )
                item_order[current_subgroup_id] = 0

            # ================================================================
            # 3. PULAR LINHAS SEM CONTEÚDO
            # ================================================================
            if not col3_item and not col4_subitem_ou_unit:
                importer.stats['rows_skipped'] += 1
                continue

            if not current_subgroup_id:
                logger.warning(f"Linha {row_num}: Item sem subgrupo definido, pulando")
                importer.stats['rows_skipped'] += 1
                continue

            # ================================================================
            # 4. PROCESSAR ITEM
            # ================================================================
            item_order[current_subgroup_id] += 1
            parent_id = None
            unit = None
            unit_conversion = None
            item_name = None

            # CASO: Col3 preenchida
            if col3_item:
                item_name = col3_item

                # Verificar se col4 tem "unit | conversion"
                if col4_subitem_ou_unit and " | " in col4_subitem_ou_unit:
                    # EXAME LABORATORIAL
                    unit, unit_conversion = parse_unit_conversion(col4_subitem_ou_unit)

                # Inserir item
                last_item_col3_id = importer.insert_item(
                    name=item_name,
                    unit=unit,
                    unit_conversion=unit_conversion,
                    points=col5_points,
                    order=item_order[current_subgroup_id],
                    subgroup_id=current_subgroup_id,
                    parent_item_id=None
                )
                item_id = last_item_col3_id

            # CASO: Col3 vazia, Col4 preenchida = SUBITEM HIERÁRQUICO
            elif col4_subitem_ou_unit:
                item_name = col4_subitem_ou_unit
                parent_id = last_item_col3_id

                item_id = importer.insert_item(
                    name=item_name,
                    unit=None,
                    unit_conversion=None,
                    points=col5_points,
                    order=item_order[current_subgroup_id],
                    subgroup_id=current_subgroup_id,
                    parent_item_id=parent_id
                )

            else:
                # Não deveria acontecer
                logger.warning(f"Linha {row_num}: Lógica inesperada")
                importer.stats['rows_skipped'] += 1
                continue

            # ================================================================
            # 5. PROCESSAR NÍVEIS (0-5)
            # ================================================================
            # SÓ criar níveis se item tem pontos > 0
            if col5_points > 0:
                for level_idx, nivel_text in enumerate(niveis):
                    if not nivel_text or not nivel_text.strip():
                        continue

                    nivel_text = nivel_text.strip()
                    operator, lower, upper = parse_range(nivel_text)

                    importer.insert_level(
                        level=level_idx,
                        name=nivel_text,
                        operator=operator,
                        lower_limit=lower,
                        upper_limit=upper,
                        item_id=item_id
                    )

    logger.info("\n" + "="*70)
    logger.info("ESTATÍSTICAS DE IMPORTAÇÃO")
    logger.info("="*70)
    logger.info(f"Linhas processadas: {importer.stats['rows_processed']}")
    logger.info(f"Linhas puladas: {importer.stats['rows_skipped']}")
    logger.info(f"Grupos criados: {importer.stats['groups']}")
    logger.info(f"Subgrupos criados: {importer.stats['subgroups']}")
    logger.info(f"Itens criados: {importer.stats['items']}")
    logger.info(f"  - Com pontos: {importer.stats['items_with_points']}")
    logger.info(f"  - Sem pontos (organizadores): {importer.stats['items_without_points']}")
    logger.info(f"Níveis criados: {importer.stats['levels']}")


# ============================================================================
# MAIN
# ============================================================================

def main():
    parser = argparse.ArgumentParser(
        description='Importar CSV de escores para banco de dados Plenya'
    )
    parser.add_argument('--csv', required=True, help='Caminho do arquivo CSV')
    parser.add_argument('--host', default='localhost', help='PostgreSQL host')
    parser.add_argument('--port', type=int, default=5432, help='PostgreSQL port')
    parser.add_argument('--user', default='plenya_user', help='PostgreSQL user')
    parser.add_argument('--password', default='plenya_dev_password', help='PostgreSQL password')
    parser.add_argument('--database', default='plenya_db', help='PostgreSQL database')
    parser.add_argument('--dry-run', action='store_true', help='Executar sem commit (rollback)')

    args = parser.parse_args()

    logger.info("="*70)
    logger.info("IMPORTAÇÃO CSV → SCORE SYSTEM DATABASE")
    logger.info("="*70)
    logger.info(f"Arquivo: {args.csv}")
    logger.info(f"Database: {args.user}@{args.host}:{args.port}/{args.database}")

    if args.dry_run:
        logger.warning("⚠️  MODO DRY-RUN: Mudanças serão revertidas!")

    # Conectar ao banco
    try:
        conn = psycopg2.connect(
            host=args.host,
            port=args.port,
            user=args.user,
            password=args.password,
            database=args.database
        )
        logger.info("✅ Conectado ao PostgreSQL\n")
    except Exception as e:
        logger.error(f"❌ Erro ao conectar: {e}")
        sys.exit(1)

    try:
        importer = ScoreImporter(conn)

        # Processar CSV
        process_csv(args.csv, importer)

        # Validar
        importer.validate_import()

        # Commit ou Rollback
        if args.dry_run:
            logger.warning("\n⚠️  ROLLBACK executado (dry-run mode)")
            conn.rollback()
        else:
            logger.info("\n✅ COMMIT executado - dados salvos!")
            conn.commit()

    except Exception as e:
        logger.error(f"\n❌ ERRO durante importação: {e}")
        logger.exception(e)
        conn.rollback()
        sys.exit(1)

    finally:
        conn.close()
        logger.info("🔒 Conexão fechada")


if __name__ == '__main__':
    main()
