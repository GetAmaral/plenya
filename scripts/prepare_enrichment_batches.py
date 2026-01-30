#!/usr/bin/env python3
"""
Prepara batches de items pendentes para enriquecimento com literatura científica
"""

import subprocess
import json
import sys

def run_sql(sql):
    """Executa SQL e retorna resultado"""
    cmd = ['docker', 'compose', 'exec', '-T', 'db', 'psql',
           '-U', 'plenya_user', '-d', 'plenya_db',
           '-t', '-A', '-F', '|', '-c', sql]
    result = subprocess.run(cmd, capture_output=True, text=True)
    if result.returncode != 0:
        print(f"Erro SQL: {result.stderr}", file=sys.stderr)
        return []
    return result.stdout.strip().split('\n') if result.stdout.strip() else []

def get_pending_items(group_name):
    """Busca items pendentes de um grupo"""
    sql = f"""
    SELECT
      si.id,
      si.name,
      sg.name as subgrupo
    FROM score_items si
    JOIN score_subgroups sg ON si.subgroup_id = sg.id
    JOIN score_groups g ON sg.group_id = g.id
    WHERE (LENGTH(COALESCE(si.clinical_relevance, '')) <= 100 OR si.clinical_relevance IS NULL)
      AND g.name = '{group_name}'
    ORDER BY si.created_at;
    """

    lines = run_sql(sql)
    items = []

    for line in lines:
        if not line.strip():
            continue
        parts = line.split('|')
        if len(parts) >= 3:
            items.append({
                'id': parts[0].strip(),
                'name': parts[1].strip(),
                'subgroup': parts[2].strip()
            })

    return items

def main():
    # Definir batches
    batches = {
        'batch1_vida_sexual': ('Vida Sexual', 2),
        'batch1_cognicao': ('Cognição', 6),
        'batch1_historico_familiar': ('Histórico Familiar de Doenças', 3),
        'batch1_social': ('Social', 9),
        'batch2_alimentacao': ('Alimentação', 23),
        'batch2_sono': ('Sono', 24),
        'batch3_composicao': ('Composição corporal', 46),
        'batch4_genetica': ('Genética', 81),
    }

    print("=" * 80)
    print("PREPARANDO BATCHES DE ENRIQUECIMENTO")
    print("=" * 80)
    print()

    total_items = 0

    for batch_key, (group_name, expected_count) in batches.items():
        print(f"📦 {batch_key}: {group_name}")

        items = get_pending_items(group_name)

        if len(items) == 0:
            print(f"  ⚠️  Nenhum item pendente encontrado!")
            continue

        if len(items) != expected_count:
            print(f"  ⚠️  Esperado: {expected_count}, Encontrado: {len(items)}")

        # Salvar JSON
        filename = f"scripts/enrichment_data/{batch_key}.json"
        with open(filename, 'w', encoding='utf-8') as f:
            json.dump({
                'group': group_name,
                'total': len(items),
                'items': items
            }, f, ensure_ascii=False, indent=2)

        print(f"  ✓ {len(items)} items salvos em {filename}")
        total_items += len(items)

        # Mostrar amostra
        if len(items) <= 5:
            for item in items:
                print(f"    - {item['name']}")
        else:
            for item in items[:3]:
                print(f"    - {item['name']}")
            print(f"    ... e mais {len(items) - 3} items")

        print()

    print("=" * 80)
    print(f"TOTAL: {total_items} items preparados em {len(batches)} batches")
    print("=" * 80)

if __name__ == '__main__':
    # Criar diretório
    import os
    os.makedirs('scripts/enrichment_data', exist_ok=True)
    main()
