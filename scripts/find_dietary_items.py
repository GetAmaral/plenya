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

# Buscar árvore completa
headers = {"Authorization": f"Bearer {token}"}
response = requests.get(f"{API_URL}/score-groups/tree", headers=headers)
groups = response.json()

print("=== BUSCANDO ITEMS DE ALIMENTAÇÃO ===\n")

# Procurar itens que contenham palavras-chave
target_names = ["Estratégia macro atual", "Livre", "Low carb", "Vegetariana", "Vegana"]

found_items = []

for group in groups:
    for subgroup in group.get('scoreSubgroups', []):
        for item in subgroup.get('scoreItems', []):
            item_name = item.get('name', '')
            # Verificar se o nome do item contém alguma das palavras-chave
            for target in target_names:
                if target.lower() in item_name.lower() or item_name.lower() in target.lower():
                    found_items.append({
                        'group': group['name'],
                        'subgroup': subgroup['name'],
                        'item': item_name,
                        'id': item['id']
                    })
                    print(f"✓ ENCONTRADO:")
                    print(f"  Grupo: {group['name']}")
                    print(f"  Subgrupo: {subgroup['name']}")
                    print(f"  Item: {item_name}")
                    print(f"  ID: {item['id']}")
                    print()
                    break

if not found_items:
    print("Nenhum item encontrado. Listando TODOS os items:")
    for group in groups:
        print(f"\nGRUPO: {group['name']}")
        for subgroup in group.get('scoreSubgroups', []):
            print(f"  SUBGRUPO: {subgroup['name']}")
            for item in subgroup.get('scoreItems', []):
                print(f"    - {item['name']} (ID: {item['id']})")
