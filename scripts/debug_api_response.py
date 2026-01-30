#!/usr/bin/env python3
"""Debug completo da resposta da API"""

import requests
import json

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

def login():
    response = requests.post(
        f"{API_BASE}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )
    response.raise_for_status()
    return response.json()['accessToken']

def main():
    token = login()
    print("Token obtido\n")

    item_id = "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0"

    print(f"Buscando item: {item_id}\n")
    response = requests.get(
        f"{API_BASE}/score-items/{item_id}",
        headers={'Authorization': f'Bearer {token}'}
    )

    print(f"Status Code: {response.status_code}")
    print(f"\nResposta completa (JSON):")
    print(json.dumps(response.json(), indent=2, ensure_ascii=False))

if __name__ == "__main__":
    main()
