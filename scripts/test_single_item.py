#!/usr/bin/env python3
import requests
import json

API_BASE = "http://localhost:3001/api/v1"

# Login
response = requests.post(
    f"{API_BASE}/auth/login",
    json={"email": "import@plenya.com", "password": "Import@123456"}
)
token = response.json()["accessToken"]

# Test first ID
test_id = "4e960d9f-a04a-4524-9049-f4559170db14"
response = requests.get(
    f"{API_BASE}/score-items/{test_id}",
    headers={"Authorization": f"Bearer {token}"}
)

print(f"Status: {response.status_code}")
print(f"Response: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
