#!/usr/bin/env python3
"""
Exporta todos os Score Items com contexto completo para processamento por IA
"""

import json
import psycopg2
from psycopg2.extras import RealDictCursor

# Conexão com banco
conn = psycopg2.connect(
    host="localhost",
    port=5432,
    database="plenya_db",
    user="plenya_user",
    password="plenya_password"
)

cursor = conn.cursor(cursor_factory=RealDictCursor)

# Query para obter todos os items com contexto
query = """
SELECT
  si.id,
  si.name as item_name,
  si.unit,
  si.unit_conversion,
  si.points,
  si.clinical_relevance,
  si.patient_explanation,
  si.conduct,
  si.last_review,
  sg.id as subgroup_id,
  sg.name as subgroup_name,
  g.id as group_id,
  g.name as group_name,
  COUNT(DISTINCT sl.id) as level_count,
  COUNT(DISTINCT asi.article_id) as article_count
FROM score_items si
LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
LEFT JOIN score_groups g ON sg.group_id = g.id
LEFT JOIN score_levels sl ON si.id = sl.item_id
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
GROUP BY si.id, si.name, si.unit, si.unit_conversion, si.points,
         si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review,
         sg.id, sg.name, g.id, g.name
ORDER BY g.name, sg.name, si.order, si.name;
"""

cursor.execute(query)
items = cursor.fetchall()

# Converter para JSON serializável
items_json = []
for item in items:
    item_dict = dict(item)
    # Converter datetime para string
    if item_dict.get('last_review'):
        item_dict['last_review'] = item_dict['last_review'].isoformat()
    items_json.append(item_dict)

# Salvar como JSON
output_file = '/home/user/plenya/score_items_context.json'
with open(output_file, 'w', encoding='utf-8') as f:
    json.dump(items_json, f, ensure_ascii=False, indent=2)

print(f"✅ Exportados {len(items_json)} Score Items para {output_file}")

# Estatísticas
total = len(items_json)
with_clinical = sum(1 for i in items_json if i['clinical_relevance'])
with_patient = sum(1 for i in items_json if i['patient_explanation'])
with_conduct = sum(1 for i in items_json if i['conduct'])
pending = sum(1 for i in items_json if not (i['clinical_relevance'] and i['patient_explanation'] and i['conduct']))

print(f"\n📊 Estatísticas:")
print(f"Total de items: {total}")
print(f"Com clinical_relevance: {with_clinical}")
print(f"Com patient_explanation: {with_patient}")
print(f"Com conduct: {with_conduct}")
print(f"Pendentes (sem todos os campos): {pending}")
print(f"\nArquivo: {output_file}")

cursor.close()
conn.close()
