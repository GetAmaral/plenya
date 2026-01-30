#!/usr/bin/env python3
"""
Enriquecimento FINAL de 20 Score Items de EXAMES
Sistema Plenya - Medicina Funcional Integrativa
Baseado em evidências científicas 2026
"""
import requests
import json
import time

API = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmOTliMzk2OC0zNTI4LTQ2YTUtYWRkNi03ODYyNmQ1NGJmMTgiLCJlbWFpbCI6ImltcG9ydEBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQ4NTQyNiwiaWF0IjoxNzY5NDg0NTI2fQ.VWVPjoeG7sCl9DLK4ogKYdzxC8nIqmDz_DjxWZHhrlI"

headers = {"Authorization": f"Bearer {TOKEN}", "Content-Type": "application/json"}

# Carregar items do JSON
with open('/home/user/plenya/complete_20_items.json', 'r') as f:
    items = json.load(f)

print("=" * 80)
print("ENRIQUECIMENTO DE 20 SCORE ITEMS DE EXAMES")
print("Sistema Plenya - Medicina Funcional Integrativa")
print("=" * 80)
print(f"\nTotal de items: {len(items)}")
print("\nIniciando processamento...\n")

success_count = 0
errors = []

for i, item in enumerate(items, 1):
    print(f"[{i}/20] {item['name']:40s}", end=" ... ")

    payload = {
        "clinicalRelevance": item["clinicalRelevance"],
        "patientExplanation": item["patientExplanation"],
        "conduct": item["conduct"]
    }

    try:
        response = requests.put(
            f"{API}/score-items/{item['id']}",
            headers=headers,
            json=payload,
            timeout=15
        )

        if response.status_code in [200, 201]:
            print("✓ SUCESSO")
            success_count += 1
        else:
            error_msg = response.text[:100] if response.text else "Sem mensagem"
            print(f"✗ ERRO {response.status_code}")
            errors.append({
                "item": item["name"],
                "id": item["id"],
                "status": response.status_code,
                "message": error_msg
            })
    except requests.exceptions.Timeout:
        print("✗ TIMEOUT")
        errors.append({"item": item["name"], "id": item["id"], "error": "Timeout após 15s"})
    except Exception as e:
        print(f"✗ EXCEÇÃO: {str(e)[:50]}")
        errors.append({"item": item["name"], "id": item["id"], "error": str(e)[:100]})

    time.sleep(0.4)  # 400ms entre requisições

print("\n" + "=" * 80)
print("RESUMO FINAL")
print("=" * 80)
print(f"  Total processado: {len(items)} items")
print(f"  Sucessos: {success_count} ({success_count/len(items)*100:.1f}%)")
print(f"  Erros: {len(errors)} ({len(errors)/len(items)*100:.1f}%)")

if errors:
    print("\n" + "-" * 80)
    print("DETALHES DOS ERROS:")
    print("-" * 80)
    for idx, err in enumerate(errors, 1):
        print(f"\n{idx}. {err['item']}")
        print(f"   ID: {err['id']}")
        if 'status' in err:
            print(f"   Status HTTP: {err['status']}")
            print(f"   Mensagem: {err['message']}")
        else:
            print(f"   Erro: {err['error']}")

print("\n" + "=" * 80)
print("EXAMES ENRIQUECIDOS:")
print("  - Vitamina D (2 items)")
print("  - ECG 12 derivações (3 items)")
print("  - Colonoscopia (3 items)")
print("  - Endoscopia digestiva alta (3 items)")
print("  - Mamografia digital bilateral (3 items)")
print("  - DEXA composição corporal (3 items)")
print("  - Ecocardiograma transtorácico (3 items)")
print("=" * 80)

# Salvar relatório
report = {
    "timestamp": time.strftime("%Y-%m-%d %H:%M:%S"),
    "total": len(items),
    "success": success_count,
    "errors": len(errors),
    "error_details": errors
}

with open('/home/user/plenya/enrichment_report_20_items.json', 'w') as f:
    json.dump(report, f, indent=2, ensure_ascii=False)

print("\nRelatório salvo em: /home/user/plenya/enrichment_report_20_items.json")
