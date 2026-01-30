#!/usr/bin/env python3
"""
HISTÓRICO FAMILIAR - 30 Items Enrichment
Gerador de SQL updates com conteúdo científico sobre predisposição genética
"""

import json

# IDs dos 30 items
ITEM_IDS = [
    "4e960d9f-a04a-4524-9049-f4559170db14",
    "26bccb28-c694-48af-8f6a-0ce6c5f73e52",
    "90d067be-f3cd-48ad-9cf9-e48fd8a2027a",
    "1c15d32d-7e79-44b8-976d-8e90ecd896be",
    "4f76ef4e-1b23-47d2-bb41-549463ad3cdf",
    "a856d80b-1443-4c16-94db-1f99508b1a9c",
    "8fb2e7de-c341-4ff5-814f-25d7695fd1cf",
    "515052d6-c7c1-40c7-942e-a1134e3aa05e",
    "8342ba54-4c2d-4b31-8619-05f4a2e86719",
    "d489c80a-50f2-4606-9579-66015e62649e",
    "394414b6-0538-4162-a68c-1e1e9d8cffef",
    "307d0403-8648-431d-88cb-2ac2422f8e86",
    "62da018d-faf4-490f-9a7b-47e0f0881bbf",
    "f4baae49-221d-428e-a218-a29391e1e26f",
    "c2c1d736-45a4-45ed-be76-e9b4704d4b1d",
    "8a552fbb-862c-4845-98dc-303f46922403",
    "12da9917-7311-4a87-89de-a1f0c8d918e7",
    "ac07f7de-eef0-420a-bdd9-cc0a3a41fbd8",
    "6491db47-381b-45fe-b483-ea0183478225",
    "35a7dbf8-2a43-4ee9-bc6b-c13122ed36c5",
    "e257d4b5-c0b2-471b-8693-defe0b2055a3",
    "77f78798-6a3a-4eef-b092-7ae1c277df4e",
    "3c52ad44-7049-40b8-b444-b05d15b96a57",
    "e2b287df-015a-4f2e-a5bf-0d99c8b24a97",
    "56fd5913-4b41-4d56-976a-b6339b4482a6",
    "1da8069d-ed8a-4b5f-83fe-0089909dd630",
    "ed899a60-9067-4a6a-aeb1-b79a08ea6062",
    "b4c8d1fe-07a2-4186-a020-051a07b98618",
    "4dc130ae-9c84-4f5f-9813-561389776254",
    "18a7fccf-4515-4807-a125-2f3e918a1d6b"
]

# Mapeamento manual que preciso completar após identificar cada item
# Vou criar um guia de processamento
print("""
==================================================================================
GUIA DE PROCESSAMENTO - HISTÓRICO FAMILIAR 30 ITEMS
==================================================================================

Este script gera instruções para enriquecer 30 items sobre histórico familiar.

METODOLOGIA:
1. Buscar cada item na API para identificar a pergunta
2. Para cada item, gerar conteúdo sobre:
   - Predisposição genética (polimorfismos, herança)
   - Fatores epigenéticos modificáveis
   - Medicina de precisão / rastreamento personalizado
   - Prevenção primária baseada em histórico familiar

FOCO CIENTÍFICO:
- Epigenética e modulação de expressão gênica
- Polimorfismos relevantes (APOE, MTHFR, TCF7L2, FTO, etc)
- Risco relativo vs risco absoluto
- Intervenções baseadas em evidências para redução de risco

==================================================================================
""")

# Vou gerar comandos curl para cada item
print("\n# COMANDOS PARA IDENTIFICAR ITEMS:")
print("# Copie e execute para ver as perguntas de cada item\n")

for i, item_id in enumerate(ITEM_IDS, 1):
    print(f"echo '\\n[{i}/30] Item {item_id}:'")
    print(f"curl -s http://localhost:3001/api/v1/score-items/{item_id} \\")
    print(f"  -H 'Authorization: Bearer $TOKEN' | \\")
    print(f"  python3 -c \"import sys, json; d=json.load(sys.stdin); print(f\\\"Question: {{d.get('question','N/A')}}\\\\nType: {{d.get('question_type','N/A')}}\\\\nGroup: {{d.get('group_name','N/A')}}\\\")\"\n")

print("\n# Para obter o token:")
print("TOKEN=$(curl -s -X POST http://localhost:3001/api/v1/auth/login \\")
print("  -H 'Content-Type: application/json' \\")
print("  -d '{\"email\":\"import@plenya.com\",\"password\":\"Import@123456\"}' | \\")
print("  python3 -c 'import sys, json; print(json.load(sys.stdin)[\"accessToken\"])')")

print("\n==================================================================================")
print("PRÓXIMOS PASSOS:")
print("==================================================================================")
print("1. Execute os comandos acima para identificar todos os 30 items")
print("2. Crie arquivo JSON com conteúdo para cada item")
print("3. Use script de atualização via API")
print("==================================================================================\n")

# Salvar lista de IDs
output_file = "/home/user/plenya/historico_familiar_30_ids.json"
with open(output_file, "w") as f:
    json.dump({"ids": ITEM_IDS, "total": len(ITEM_IDS)}, f, indent=2)

print(f"✓ Lista de IDs salva em: {output_file}")
