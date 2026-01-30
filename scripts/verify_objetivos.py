#!/usr/bin/env python3
"""Verificar se items foram enriquecidos corretamente"""

import requests
import subprocess

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login():
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    return response.json()["accessToken"]

def get_item(token, item_id):
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers=headers
    )
    return response.json()

# Login
token = login()

# Testar alguns items
test_items = [
    "c0a05918-804e-4550-b182-66b83c7fc141",  # Disciplina moderada
    "1318016c-736c-45c9-aca4-a98fdebd5996",  # 5 anos
    "676721d6-cc37-4ee2-a2df-5e188b542326",  # Objetivos iniciais
]

print("Verificando items via API:\n")

for item_id in test_items:
    item = get_item(token, item_id)
    name = item.get("name", "N/A")

    clinical = item.get("clinicalRelevance", "")
    patient = item.get("patientExplanation", "")
    conduct = item.get("conduct", "")

    print(f"Item: {name}")
    print(f"  clinical_relevance: {len(clinical) if clinical else 0} chars")
    print(f"  patient_explanation: {len(patient) if patient else 0} chars")
    print(f"  conduct: {len(conduct) if conduct else 0} chars")
    print()

# Verificar banco
print("\nVerificando banco de dados:\n")

cmd = [
    "docker", "compose", "exec", "-T", "db",
    "psql", "-U", "plenya_user", "-d", "plenya_db",
    "-c", """
    SELECT
        name,
        LENGTH(clinical_relevance) as len_clinical,
        LENGTH(patient_explanation) as len_patient,
        LENGTH(conduct) as len_conduct
    FROM score_items
    WHERE id = 'c0a05918-804e-4550-b182-66b83c7fc141';
    """
]

result = subprocess.run(cmd, capture_output=True, text=True)
print(result.stdout)
