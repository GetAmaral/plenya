#!/usr/bin/env python3
import requests
import json

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# Login
response = requests.post(
    f"{API_URL}/auth/login",
    json={"email": EMAIL, "password": PASSWORD}
)
token = response.json()["accessToken"]

# Buscar árvore
headers = {"Authorization": f"Bearer {token}"}
response = requests.get(f"{API_URL}/score-groups/tree", headers=headers)
groups = response.json()

print("=== ESTRUTURA DO SCORE ===\n")

for group in groups:
    print(f"GROUP: {group['name']} (ID: {group['id']})")
    for subgroup in group.get('scoreSubgroups', []):
        print(f"  SUBGROUP: {subgroup['name']} (ID: {subgroup['id']})")
        for item in subgroup.get('scoreItems', []):
            print(f"    ITEM: {item['name']} (ID: {item['id']})")
    print()
