#!/usr/bin/env python3
"""
Migração completa de genes: grupos → items nos subgrupos de Genética
"""

import subprocess
import json

# Mapeamento gene → subgrupo
GENE_MAPPING = {
    # METABOLISMO
    'MTHFR C677T rs1801133 (Homocisteína)': 'Metabolismo',
    'FTO rs9939609 (Obesidade)': 'Metabolismo',
    'MC4R rs17782313 (Obesidade)': 'Metabolismo',
    'TCF7L2 rs7903146 (Diabetes)': 'Metabolismo',
    'PPARG Pro12Ala rs1801282 (Diabetes)': 'Metabolismo',
    'VDR FokI rs2228570 (Vitamina D)': 'Metabolismo',
    'ABCC8 rs757110 (Diabetes)': 'Metabolismo',
    'CDKAL1 rs7754840 (Diabetes)': 'Metabolismo',
    'HHEX rs1111875 (Diabetes)': 'Metabolismo',
    'IGF2BP2 rs4402960 (Diabetes)': 'Metabolismo',
    'INS VNTR (Diabetes Tipo 1)': 'Metabolismo',
    'IRS1 rs1801278 (Resistência Insulina)': 'Metabolismo',
    'KCNJ11 rs5219 (Diabetes)': 'Metabolismo',
    'SLC30A8 rs13266634 (Diabetes)': 'Metabolismo',
    'PPARGC1A rs8192678 Gly482Ser (Metabolismo)': 'Metabolismo',
    'LEPR rs1137101 Gln223Arg (Obesidade)': 'Metabolismo',
    'POMC rs6713532 (Obesidade)': 'Metabolismo',
    'PPARA rs4253778 (Resistência)': 'Metabolismo',
    'GCK (MODY2)': 'Metabolismo',
    'HNF1A (MODY3)': 'Metabolismo',
    'HNF1B (MODY5)': 'Metabolismo',
    'HNF4A (MODY1)': 'Metabolismo',
    'BCO1 rs6564851 (Vitamina A)': 'Metabolismo',
    'FABP2 Ala54Thr rs1799883 (Gordura)': 'Metabolismo',
    'FADS1 rs174546 (Ômega-3)': 'Metabolismo',
    'FADS2 rs174575 (Ômega-3/DHA)': 'Metabolismo',
    'MCM6 rs4988235 (Lactose)': 'Metabolismo',
    'SLC23A1 rs33972313 (Vitamina C)': 'Metabolismo',

    # CARDIOVASCULAR
    'APOE (Alzheimer e Lipídios)': 'Cardiovascular',
    'ACE I/D rs4646994 (Hipertensão)': 'Cardiovascular',
    'AGT rs699 M235T (Hipertensão)': 'Cardiovascular',
    'AGTR1 rs5186 A1166C (Hipertensão)': 'Cardiovascular',
    'NOS3 rs1799983 Glu298Asp (Hipertensão)': 'Cardiovascular',
    'ADD1 rs4961 Gly460Trp (Hipertensão)': 'Cardiovascular',
    'CYP11B2 rs1799998 (Hipertensão)': 'Cardiovascular',
    'GNB3 rs5443 C825T (Hipertensão)': 'Cardiovascular',
    'LDLR rs688 (Colesterol LDL)': 'Cardiovascular',
    'PCSK9 rs11591147 R46L (Colesterol)': 'Cardiovascular',
    'ABCA1 rs9282541 (HDL)': 'Cardiovascular',
    'APOA1 rs670 (HDL)': 'Cardiovascular',
    'APOA5 rs662799 (Triglicerídeos)': 'Cardiovascular',
    'LCAT rs5923 (HDL)': 'Cardiovascular',
    'LIPC rs1800588 (HDL)': 'Cardiovascular',
    'LPL rs328 (Triglicerídeos)': 'Cardiovascular',
    'APOC2 (Deficiência)': 'Cardiovascular',
    'GPIHBP1 (Quilomicronemia)': 'Cardiovascular',
    'LMF1 (Hipertrigliceridemia)': 'Cardiovascular',

    # NEURODEGENERAÇÃO
    'APP A673T rs63750847 (Alzheimer)': 'Neurodegeneração',
    'PSEN1 E280A (Alzheimer Familial)': 'Neurodegeneração',
    'PSEN2 (Alzheimer Familial)': 'Neurodegeneração',
    'LRRK2 G2019S rs34637584 (Parkinson)': 'Neurodegeneração',
    'SNCA rs356219 (Parkinson)': 'Neurodegeneração',
    'PARK2 (Parkinson Precoce)': 'Neurodegeneração',
    'PARK7/DJ-1 (Parkinson Precoce)': 'Neurodegeneração',
    'PINK1 (Parkinson Precoce)': 'Neurodegeneração',
    'C9orf72 (DFT/ELA)': 'Neurodegeneração',
    'MAPT rs1467967 H1/H2 (Demência FTD)': 'Neurodegeneração',
    'GRN rs5848 (Demência FTD)': 'Neurodegeneração',

    # DETOXIFICAÇÃO
    'CYP1A1 rs4646903 MspI (Detoxificação)': 'Detoxificação',
    'CYP1A2 rs762551 (Cafeína)': 'Detoxificação',
    'CYP2A6 rs1801272 (Nicotina)': 'Detoxificação',
    'GSTM1 (Detoxificação)': 'Detoxificação',
    'GSTP1 rs1695 Ile105Val (Detoxificação)': 'Detoxificação',
    'GSTT1 (Detoxificação)': 'Detoxificação',
    'NAT2 (Acetilador)': 'Detoxificação',
    'EPHX1 rs1051740 Tyr113His (Detoxificação)': 'Detoxificação',
    'ADH1B rs1229984 Arg48His (Álcool)': 'Detoxificação',
    'ALDH2 rs671 Glu504Lys (Álcool)': 'Detoxificação',
    'GPX1 rs1050450 Pro198Leu (Glutationa)': 'Detoxificação',
    'SOD2 rs4880 Ala16Val (Antioxidante)': 'Detoxificação',
    'CAT rs1001179 -262C>T (Catalase)': 'Detoxificação',

    # IMUNIDADE
    'HLA-DQ2/DQ8 (Doença Celíaca)': 'Imunidade',
    'IL1B rs16944 -511C>T (Inflamação)': 'Imunidade',
    'IL6 rs1800795 -174G>C (Inflamação)': 'Imunidade',
    'TNF rs1800629 -308G>A (Inflamação)': 'Imunidade',
    'CRP rs1130864 (Proteína C Reativa)': 'Imunidade',

    # PERFORMANCE
    'ACTN3 rs1815739 R577X (Performance)': 'Performance',
    'COL5A1 rs12722 (Lesão Tendão)': 'Performance',
    'COL1A1 rs1800012 Sp1 (Osteoporose)': 'Performance',
    'ESR1 rs2234693 PvuII (Osteoporose)': 'Performance',

    # OUTROS
    'ALPL (Hipofosfatasia)': 'Outros',
}

def run_sql(sql, description):
    """Executa SQL via Docker"""
    print(f"  → {description}")
    cmd = ['docker', 'compose', 'exec', '-T', 'db', 'psql', '-U', 'plenya_user', '-d', 'plenya_db', '-c', sql]
    result = subprocess.run(cmd, capture_output=True, text=True)
    if result.returncode != 0:
        print(f"    ❌ Erro: {result.stderr}")
        return None
    return result.stdout

def get_subgroup_ids():
    """Busca IDs dos subgrupos de Genética"""
    sql = """
    SELECT sg.name, sg.id::text
    FROM score_subgroups sg
    JOIN score_groups g ON sg.group_id = g.id
    WHERE g.name = 'Genética'
    ORDER BY sg.name;
    """
    output = run_sql(sql, "Buscando IDs dos subgrupos")

    subgroups = {}
    for line in output.strip().split('\n')[2:]:  # Pular cabeçalho e separador
        line = line.strip()
        if not line or line.startswith('('):  # Ignorar linha vazia ou contador
            continue
        parts = line.split('|')
        if len(parts) >= 2:
            name = parts[0].strip()
            id = parts[1].strip()
            subgroups[name] = id

    return subgroups

def get_gene_groups():
    """Lista grupos que são genes"""
    genes_list = "','".join(GENE_MAPPING.keys())
    sql = f"""
    SELECT name, id::text
    FROM score_groups
    WHERE name IN ('{genes_list}')
    ORDER BY name;
    """
    output = run_sql(sql, "Buscando grupos genéticos")

    groups = {}
    for line in output.strip().split('\n')[2:]:  # Pular cabeçalho e separador
        line = line.strip()
        if not line or line.startswith('('):  # Ignorar linha vazia ou contador
            continue
        parts = line.split('|')
        if len(parts) >= 2:
            name = parts[0].strip()
            id = parts[1].strip()
            groups[name] = id

    return groups

def migrate_gene(gene_name, gene_group_id, subgroup_id, order_num):
    """Migra um gene: grupo → item"""

    # SQL de migração
    sql = f"""
    BEGIN;

    -- 1. Criar item no subgrupo correto de Genética
    INSERT INTO score_items (subgroup_id, name, "order")
    VALUES ('{subgroup_id}', '{gene_name.replace("'", "''")}', {order_num})
    RETURNING id::text;

    -- 2. Buscar subgrupos do gene-grupo (que virarão levels)
    -- 3. Criar levels a partir desses subgrupos
    INSERT INTO score_levels (item_id, name, level)
    SELECT
        (SELECT id FROM score_items WHERE subgroup_id = '{subgroup_id}' AND name = '{gene_name.replace("'", "''")}' ORDER BY created_at DESC LIMIT 1),
        sg.name,
        (ROW_NUMBER() OVER (ORDER BY sg."order") - 1)::integer
    FROM score_subgroups sg
    WHERE sg.group_id = '{gene_group_id}';

    -- 4. Migrar article links de items antigos para o novo item
    INSERT INTO article_score_items (article_id, score_item_id)
    SELECT DISTINCT
        asi.article_id,
        (SELECT id FROM score_items WHERE subgroup_id = '{subgroup_id}' AND name = '{gene_name.replace("'", "''")}' ORDER BY created_at DESC LIMIT 1)
    FROM article_score_items asi
    JOIN score_items si ON asi.score_item_id = si.id
    JOIN score_subgroups sg ON si.subgroup_id = sg.id
    WHERE sg.group_id = '{gene_group_id}'
    ON CONFLICT (article_id, score_item_id) DO NOTHING;

    -- 5. Deletar article links antigos
    DELETE FROM article_score_items
    WHERE score_item_id IN (
        SELECT si.id
        FROM score_items si
        JOIN score_subgroups sg ON si.subgroup_id = sg.id
        WHERE sg.group_id = '{gene_group_id}'
    );

    -- 6. Deletar levels antigos (se houver) - ANTES de deletar items
    DELETE FROM score_levels
    WHERE item_id IN (
        SELECT si.id
        FROM score_items si
        JOIN score_subgroups sg ON si.subgroup_id = sg.id
        WHERE sg.group_id = '{gene_group_id}'
    );

    -- 7. Deletar items antigos dos subgrupos do gene
    DELETE FROM score_items
    WHERE subgroup_id IN (
        SELECT id FROM score_subgroups WHERE group_id = '{gene_group_id}'
    );

    -- 8. Deletar subgrupos do gene-grupo
    DELETE FROM score_subgroups WHERE group_id = '{gene_group_id}';

    -- 9. Deletar gene-grupo
    DELETE FROM score_groups WHERE id = '{gene_group_id}';

    COMMIT;

    -- Verificar
    SELECT
        'MIGRADO' as status,
        si.name as item_name,
        COUNT(sl.id) as levels_criados
    FROM score_items si
    LEFT JOIN score_levels sl ON sl.item_id = si.id
    WHERE si.subgroup_id = '{subgroup_id}' AND si.name = '{gene_name.replace("'", "''")}'
    GROUP BY si.name;
    """

    output = run_sql(sql, f"Migrando {gene_name}")
    return output

def main():
    print("="*80)
    print("MIGRAÇÃO COMPLETA DE EXAMES GENÉTICOS")
    print("="*80)
    print()

    # 1. Buscar IDs dos subgrupos
    print("PASSO 1: Buscar subgrupos de Genética")
    subgroups = get_subgroup_ids()
    print(f"  ✓ {len(subgroups)} subgrupos encontrados")
    for name, id in subgroups.items():
        print(f"    - {name}: {id}")
    print()

    # 2. Buscar grupos genéticos
    print("PASSO 2: Buscar grupos genéticos")
    gene_groups = get_gene_groups()
    print(f"  ✓ {len(gene_groups)} grupos genéticos encontrados")
    print()

    # 3. Migrar cada gene
    print("PASSO 3: Migrar genes (grupos → items)")
    print()

    migrated = 0
    failed = []

    for gene_name, target_subgroup in sorted(GENE_MAPPING.items()):
        if gene_name not in gene_groups:
            print(f"  ⚠ {gene_name}: GRUPO NÃO ENCONTRADO (pode não existir no banco)")
            continue

        if target_subgroup not in subgroups:
            print(f"  ❌ {gene_name}: SUBGRUPO '{target_subgroup}' NÃO EXISTE")
            failed.append(gene_name)
            continue

        gene_group_id = gene_groups[gene_name]
        subgroup_id = subgroups[target_subgroup]

        print(f"  Migrando: {gene_name}")
        print(f"    De: grupo {gene_group_id}")
        print(f"    Para: subgrupo {target_subgroup} ({subgroup_id})")

        result = migrate_gene(gene_name, gene_group_id, subgroup_id, migrated + 1)

        if result and 'MIGRADO' in result:
            migrated += 1
            print(f"    ✓ Sucesso ({migrated}/{len(GENE_MAPPING)})")
        else:
            failed.append(gene_name)
            print(f"    ❌ Falhou")

        print()

    # 4. Resumo
    print("="*80)
    print("RESUMO DA MIGRAÇÃO")
    print("="*80)
    print(f"  Total no mapeamento: {len(GENE_MAPPING)}")
    print(f"  Migrados com sucesso: {migrated}")
    print(f"  Falharam: {len(failed)}")

    if failed:
        print(f"\n  Genes que falharam:")
        for gene in failed:
            print(f"    - {gene}")

    print()
    print("✓ Migração concluída!")
    print()

if __name__ == "__main__":
    main()
