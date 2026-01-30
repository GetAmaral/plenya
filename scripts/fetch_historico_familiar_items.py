#!/usr/bin/env python3
"""Busca os 30 items de histórico familiar para análise"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"

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

# Login
response = requests.post(
    f"{API_BASE}/auth/login",
    json={"email": "import@plenya.com", "password": "Import@123456"}
)

if response.status_code != 200:
    print(f"Login failed: {response.status_code}")
    exit(1)

token = response.json()["accessToken"]
print("✓ Login successful\n")

# Fetch all items
items_data = []
for i, item_id in enumerate(ITEM_IDS, 1):
    response = requests.get(
        f"{API_BASE}/score-items/{item_id}",
        headers={"Authorization": f"Bearer {token}"}
    )

    if response.status_code == 200:
        item = response.json()
        group_name = "N/A"
        if item.get("subgroup") and item["subgroup"].get("group"):
            group_name = item["subgroup"]["group"].get("name", "N/A")

        items_data.append({
            "id": item_id,
            "name": item.get("name", "N/A"),
            "points": item.get("points", 0),
            "group_name": group_name,
            "subgroup_name": item.get("subgroup", {}).get("name", "N/A"),
            "parent_item_id": item.get("parentItemId"),
            "parent_item_name": item.get("parentItem", {}).get("name"),
            "levels": item.get("levels", []),
            "current_clinical_explanation": item.get("clinicalExplanation"),
            "current_clinical_recommendation": item.get("clinicalRecommendation")
        })
        print(f"[{i}/30] {item.get('name', 'N/A')[:80]}")
    else:
        print(f"[{i}/30] ERROR fetching {item_id}")

# Save to JSON
output_file = "/home/user/plenya/historico_familiar_30_items.json"
with open(output_file, "w", encoding="utf-8") as f:
    json.dump(items_data, f, indent=2, ensure_ascii=False)

print(f"\n✓ Saved {len(items_data)} items to: {output_file}")
