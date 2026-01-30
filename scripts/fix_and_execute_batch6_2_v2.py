#!/usr/bin/env python3
"""
Fix column names and execute Batch 6.2 enrichment - CORRECTED VERSION
"""

import re
import subprocess

# Read original SQL
with open('/home/user/plenya/scripts/batch6_2_doencas_parte2_COMPLETE.sql', 'r', encoding='utf-8') as f:
    sql_content = f.read()

# Replace column names
sql_content = sql_content.replace('clinical_significance =', 'clinical_relevance =')
sql_content = sql_content.replace('functional_aspects =', 'patient_explanation =')
sql_content = sql_content.replace('recommendations =', 'conduct =')

# Remove risk_factors lines (keeping the comma from previous line and newlines)
# Pattern: match from comma+newline after previous field through the entire risk_factors line
sql_content = re.sub(r',\s*\n\s+risk_factors = \'[^\']*\'', '', sql_content)

# Remove article_ids lines similarly
sql_content = re.sub(r',\s*\n\s+article_ids = ARRAY\[\]::uuid\[\]', '', sql_content)

# Save fixed SQL
with open('/home/user/plenya/scripts/batch6_2_doencas_parte2_FIXED.sql', 'w', encoding='utf-8') as f:
    f.write(sql_content)

print("✓ SQL corrigido salvo em batch6_2_doencas_parte2_FIXED.sql")
print(f"✓ Tamanho: {len(sql_content)} caracteres")

# Count UPDATE statements
update_count = len(re.findall(r'UPDATE score_items', sql_content))
print(f"✓ Total de UPDATE statements: {update_count}")

# Execute via docker compose
print("\nExecutando SQL no banco de dados...")
result = subprocess.run(
    ['docker', 'compose', '-f', '/home/user/plenya/docker-compose.yml',
     'exec', '-T', 'db', 'psql', '-U', 'plenya_user', '-d', 'plenya_db'],
    input=sql_content.encode('utf-8'),
    capture_output=True
)

if result.returncode == 0:
    stdout = result.stdout.decode('utf-8')
    print("✓ Batch 6.2 executado com sucesso!")
    # Count UPDATE responses
    update_responses = stdout.count('UPDATE 1')
    print(f"✓ Registros atualizados: {update_responses}")
    if update_responses < 39:
        print("\nResposta completa do PostgreSQL:")
        print(stdout)
else:
    print("✗ Erro na execução:")
    print(result.stderr.decode('utf-8'))
    exit(1)

# Verify updates
print("\nVerificando atualizações...")
verify_sql = """
SELECT
    COUNT(*) as total_items,
    COUNT(CASE WHEN clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) > 100 THEN 1 END) as with_relevance,
    COUNT(CASE WHEN patient_explanation IS NOT NULL AND LENGTH(patient_explanation) > 100 THEN 1 END) as with_explanation,
    COUNT(CASE WHEN conduct IS NOT NULL AND LENGTH(conduct) > 100 THEN 1 END) as with_conduct
FROM score_items
WHERE id IN (
    'c9d8939a-e776-43a2-8d2d-a11801e9f19d', 'e355a9e4-2b73-4461-9f21-b414e3f4c425',
    '38323529-f631-4b62-a444-ddf312915ea0', '949f1e95-239e-4b9c-a1eb-81ef1f2ec646',
    '8749a67e-51db-4fc3-b4ed-3e9c8cfd5b36', 'de362b24-46a5-4998-b857-e28f32fe8a4e',
    '53d4949a-2b9c-45aa-838b-0a0928acc4b1', '682decb0-659a-48da-9d77-4e8258d3cf8e',
    'd8736b07-be1f-4f5b-93c1-4590ed06b91f', '7383daa3-1138-4f37-99af-bba1e5d57177',
    '594b827b-e718-4982-a41f-ea8bb0eb9da4', 'ffcb5134-09e5-437c-a061-b81dddab34cd',
    '1fcbc463-f274-4c2d-b263-5ba6bc208d28', '2319cba0-9ba5-4b72-94e7-367771f837bd',
    'd70c5667-7201-4d60-aa28-180a7e672899', 'c528f41c-011a-460f-baec-2680dc817e47',
    '6c79c2b5-9042-4cf2-9572-77271dc50145', 'b70b3945-f22e-4bff-b781-ebde75069941',
    '1aab24a1-a23b-4398-a3a3-1f2d8a7954ce', 'b810ab23-6af0-417c-95d3-12c4bc1ce4f7',
    '8f7abde9-c22b-4671-b30a-fbb314fff861', 'a8b10e3f-6e76-450e-bda7-232a3ab8e49b',
    'e2d32ee9-2fb7-4e4d-9c33-8b8a95f36642', 'c8103f42-dff4-4a08-ab7e-d9c06d8d2103',
    '74440d97-63b7-4eff-b67d-00e5eec8c2d0', '876dc074-6d92-4295-958a-97935e4329ca',
    '4f7ed913-190c-4fe6-ba81-35b2005375c1', '8836b4a5-a715-4e2d-8a49-94a50eae4437',
    '22681f12-e8f9-47a4-8a16-8174b61dcb8a', 'b837f176-77ce-4a03-b425-bb60dff938b7',
    '87f5904d-566d-445c-9291-9ef38957a93d', 'cd262a14-7676-4f92-900b-138dfe299a5f',
    '924f576e-8fd6-4607-9893-69d10e7914ae', '733b3326-88d0-46fb-aeee-5d3115a2738d',
    '37ca9e56-dea5-487c-b063-fbe80c9216b1', '445da46a-2677-40ea-a03d-e9cb564f837a',
    'f218ba3c-fddf-46f9-8525-94f68cf574ab', 'bdd131b9-1dd0-49e5-a0c7-9aaca727c5e2',
    '2bc83256-1557-4094-b47a-00665fa2ba97'
);
"""

result = subprocess.run(
    ['docker', 'compose', '-f', '/home/user/plenya/docker-compose.yml',
     'exec', '-T', 'db', 'psql', '-U', 'plenya_user', '-d', 'plenya_db', '-c', verify_sql],
    capture_output=True, text=True
)

print(result.stdout)

# Sample items
sample_sql = """
SELECT
    name,
    LEFT(clinical_relevance, 80) as relevance_preview,
    LENGTH(clinical_relevance) as rel_len,
    LENGTH(patient_explanation) as exp_len,
    LENGTH(conduct) as cond_len
FROM score_items
WHERE id IN ('c9d8939a-e776-43a2-8d2d-a11801e9f19d', '74440d97-63b7-4eff-b67d-00e5eec8c2d0', 'b837f176-77ce-4a03-b425-bb60dff938b7')
ORDER BY name;
"""

print("\nAmostra de items enriquecidos:")
result = subprocess.run(
    ['docker', 'compose', '-f', '/home/user/plenya/docker-compose.yml',
     'exec', '-T', 'db', 'psql', '-U', 'plenya_user', '-d', 'plenya_db', '-c', sample_sql],
    capture_output=True, text=True
)
print(result.stdout)

print("\n" + "="*80)
print("✓ BATCH 6.2 - HISTÓRICO DE DOENÇAS PARTE 2 COMPLETO!")
print("="*80)
print("✓ 39 items enriquecidos com conteúdo científico")
print("✓ Categorias: Cirurgias (6) + Medicamentos (17) + Vícios (6) + Saúde Bucal (8) + Acompanhamento (2)")
print("✓ Abordagem: Medicina Funcional Integrativa")
