#!/usr/bin/env python3
"""
Script de correção da estrutura de scores:
1. Remove duplicatas de score_items (mantém apenas com literatura)
2. Reorganiza grupo Genética (exames como items, não como grupos)
"""

import psycopg2
from psycopg2.extras import RealDictCursor
import json
from datetime import datetime

# Configuração do banco
DB_CONFIG = {
    "dbname": "plenya_db",
    "user": "plenya_user",
    "password": "Plenya@2024!Secure",
    "host": "localhost",
    "port": "5432"
}

def get_connection():
    return psycopg2.connect(**DB_CONFIG)

def log(msg):
    print(f"[{datetime.now().strftime('%H:%M:%S')}] {msg}")

# ============================================================================
# PARTE 1: IDENTIFICAR E REMOVER DUPLICATAS
# ============================================================================

def find_duplicates(cursor):
    """Identifica items duplicados por nome"""

    query = """
    WITH duplicates AS (
        SELECT
            name,
            COUNT(*) as count,
            ARRAY_AGG(id ORDER BY
                CASE
                    WHEN LENGTH(COALESCE(clinical_relevance, '')) > 100 THEN 0  -- Com conteúdo primeiro
                    ELSE 1  -- Vazios depois
                END,
                created_at NULLS LAST  -- Mais antigos primeiro
            ) as ids,
            ARRAY_AGG(
                LENGTH(COALESCE(clinical_relevance, ''))
                ORDER BY
                    CASE
                        WHEN LENGTH(COALESCE(clinical_relevance, '')) > 100 THEN 0
                        ELSE 1
                    END,
                    created_at NULLS LAST
            ) as content_lengths
        FROM score_items
        GROUP BY name
        HAVING COUNT(*) > 1
    )
    SELECT
        name,
        count as total_copies,
        ids,
        content_lengths,
        ids[1] as keep_id,  -- Primeiro da lista (com conteúdo se existir)
        ids[2:] as delete_ids  -- Resto para excluir
    FROM duplicates
    ORDER BY count DESC, name;
    """

    cursor.execute(query)
    return cursor.fetchall()

def count_levels_to_delete(cursor, item_ids):
    """Conta quantos levels serão deletados"""

    if not item_ids:
        return 0

    placeholders = ','.join(['%s'] * len(item_ids))
    query = f"""
    SELECT COUNT(*)
    FROM score_levels
    WHERE item_id IN ({placeholders});
    """

    cursor.execute(query, item_ids)
    return cursor.fetchone()[0]

def delete_duplicates(cursor, dry_run=True):
    """Remove items e levels duplicados"""

    duplicates = find_duplicates(cursor)

    log(f"Encontrados {len(duplicates)} conjuntos de duplicatas")

    stats = {
        "total_duplicates": len(duplicates),
        "items_to_delete": 0,
        "levels_to_delete": 0,
        "items_to_keep": 0
    }

    all_delete_ids = []

    for dup in duplicates:
        name = dup['name']
        total = dup['total_copies']
        keep_id = dup['keep_id']
        delete_ids = dup['delete_ids']
        content_lengths = dup['content_lengths']

        keep_has_content = content_lengths[0] > 100 if content_lengths else False

        log(f"\n  '{name[:60]}': {total} cópias")
        log(f"    Manter: {keep_id} (conteúdo: {content_lengths[0] if content_lengths else 0} chars)")
        log(f"    Excluir: {len(delete_ids)} cópias")

        all_delete_ids.extend(delete_ids)
        stats["items_to_delete"] += len(delete_ids)
        stats["items_to_keep"] += 1

    # Contar levels a deletar
    if all_delete_ids:
        stats["levels_to_delete"] = count_levels_to_delete(cursor, all_delete_ids)

    log(f"\n{'='*80}")
    log(f"RESUMO DA LIMPEZA:")
    log(f"  Items a manter: {stats['items_to_keep']}")
    log(f"  Items a excluir: {stats['items_to_delete']}")
    log(f"  Levels a excluir: {stats['levels_to_delete']}")
    log(f"{'='*80}\n")

    if dry_run:
        log("DRY RUN - Nenhuma alteração foi feita")
        return stats

    # Executar deleções
    if all_delete_ids:
        # 1. Deletar levels primeiro (FK constraint)
        placeholders = ','.join(['%s'] * len(all_delete_ids))
        cursor.execute(f"""
            DELETE FROM score_levels
            WHERE item_id IN ({placeholders});
        """, all_delete_ids)
        log(f"✓ {cursor.rowcount} levels deletados")

        # 2. Deletar items
        cursor.execute(f"""
            DELETE FROM score_items
            WHERE id IN ({placeholders});
        """, all_delete_ids)
        log(f"✓ {cursor.rowcount} items deletados")

    return stats

# ============================================================================
# PARTE 2: REORGANIZAR GENÉTICA
# ============================================================================

def get_genetic_groups(cursor):
    """Lista grupos que deveriam ser exames genéticos"""

    # Genes conhecidos que estão como grupos
    genetic_patterns = [
        'MTHFR', 'APOE', 'FTO', 'MC4R', 'TCF7L2', 'PPARG', 'VDR', 'ACE',
        'CYP', 'GSTT', 'GSTM', 'SOD', 'NAT2', 'HLA', 'COMT', 'MAOA'
    ]

    pattern_condition = ' OR '.join([f"name ILIKE '%{p}%'" for p in genetic_patterns])

    query = f"""
    SELECT id, name
    FROM score_groups
    WHERE {pattern_condition}
       OR name ~ '[A-Z]+[0-9]+'  -- Padrão de gene (ex: CYP1A2)
    ORDER BY name;
    """

    cursor.execute(query)
    return cursor.fetchall()

def reorganize_genetics(cursor, dry_run=True):
    """Reorganiza estrutura do grupo Genética"""

    # 1. Verificar se grupo Genética existe
    cursor.execute("SELECT id FROM score_groups WHERE name = 'Genética';")
    genetics_group = cursor.fetchone()

    if not genetics_group:
        log("❌ Grupo 'Genética' não encontrado!")
        return

    genetics_group_id = genetics_group['id']
    log(f"✓ Grupo Genética encontrado: {genetics_group_id}")

    # 2. Listar grupos que deveriam ser exames
    wrong_groups = get_genetic_groups(cursor)
    log(f"\n✓ Encontrados {len(wrong_groups)} grupos genéticos incorretos")

    for group in wrong_groups[:10]:  # Mostrar apenas 10 primeiros
        log(f"  - {group['name']}")

    if len(wrong_groups) > 10:
        log(f"  ... e mais {len(wrong_groups) - 10}")

    # 3. Criar subgrupos lógicos no grupo Genética
    subgroups_to_create = [
        ("Metabolismo", "Genes relacionados ao metabolismo de nutrientes e energia"),
        ("Cardiovascular", "Genes relacionados à saúde cardiovascular"),
        ("Neurodegeneração", "Genes relacionados a doenças neurodegenerativas"),
        ("Desintoxicação", "Genes de fase I e II de detoxificação"),
        ("Imunidade", "Genes relacionados ao sistema imune"),
        ("Hormônios", "Genes relacionados a hormônios e receptores"),
        ("Performance", "Genes relacionados a performance física")
    ]

    log(f"\n{'='*80}")
    log("REORGANIZAÇÃO DO GRUPO GENÉTICA:")
    log(f"  Grupos incorretos a mover: {len(wrong_groups)}")
    log(f"  Subgrupos a criar: {len(subgroups_to_create)}")
    log(f"{'='*80}\n")

    if dry_run:
        log("DRY RUN - Nenhuma alteração foi feita")
        return

    # Criar subgrupos
    created_subgroups = {}
    for name, description in subgroups_to_create:
        cursor.execute("""
            INSERT INTO score_subgroups (group_id, name, description)
            VALUES (%s, %s, %s)
            ON CONFLICT DO NOTHING
            RETURNING id;
        """, (genetics_group_id, name, description))

        result = cursor.fetchone()
        if result:
            created_subgroups[name] = result['id']
            log(f"✓ Subgrupo criado: {name}")
        else:
            # Já existe, buscar ID
            cursor.execute("""
                SELECT id FROM score_subgroups
                WHERE group_id = %s AND name = %s;
            """, (genetics_group_id, name))
            created_subgroups[name] = cursor.fetchone()['id']
            log(f"⚠ Subgrupo já existe: {name}")

    log(f"\n✓ {len(created_subgroups)} subgrupos prontos")

    # TODO: Mover grupos genéticos para items nos subgrupos apropriados
    # Isso requer análise manual de cada gene para categorizá-lo corretamente
    log("\n⚠ Categorização manual necessária para completar reorganização")
    log("  Execute a query abaixo para mapear genes aos subgrupos:\n")

    for group in wrong_groups[:5]:
        log(f"  -- {group['name']} → qual subgrupo?")

# ============================================================================
# MAIN
# ============================================================================

def main():
    import argparse

    parser = argparse.ArgumentParser(description="Corrige estrutura de scores")
    parser.add_argument('--dry-run', action='store_true', help='Simula sem fazer mudanças')
    parser.add_argument('--duplicates', action='store_true', help='Remove duplicatas')
    parser.add_argument('--genetics', action='store_true', help='Reorganiza genética')
    parser.add_argument('--all', action='store_true', help='Faz tudo')

    args = parser.parse_args()

    conn = get_connection()
    cursor = conn.cursor(cursor_factory=RealDictCursor)

    try:
        if args.duplicates or args.all:
            log("="*80)
            log("PARTE 1: REMOÇÃO DE DUPLICATAS")
            log("="*80)
            stats = delete_duplicates(cursor, dry_run=args.dry_run)

            # Salvar estatísticas
            with open('/home/user/plenya/duplicate_cleanup_stats.json', 'w') as f:
                json.dump(stats, f, indent=2)

        if args.genetics or args.all:
            log("\n" + "="*80)
            log("PARTE 2: REORGANIZAÇÃO GENÉTICA")
            log("="*80)
            reorganize_genetics(cursor, dry_run=args.dry_run)

        if not args.dry_run:
            conn.commit()
            log("\n✓ Alterações commitadas!")
        else:
            conn.rollback()
            log("\n⚠ DRY RUN - Nenhuma alteração foi salva")

    except Exception as e:
        conn.rollback()
        log(f"\n❌ Erro: {e}")
        import traceback
        traceback.print_exc()

    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    main()
