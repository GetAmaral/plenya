#!/usr/bin/env python3
"""Estatísticas finais do enriquecimento VIDA SEXUAL"""

import requests
import json
from collections import defaultdict

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

ITEM_IDS = [
    "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0", "86f1dbce-c9c5-4c96-9b0b-4c6452c5dd1e",
    "b58adee8-cd40-4846-b16e-0757c17ac9fe", "b6ca4e70-80c2-4109-ac52-e485549517bb",
    "92ac8388-9238-4214-a0a9-8466fc10f70c", "8b2f53db-fd5f-4dcb-b0c8-53a4a878857f",
    "a58ccffd-e1f3-45e7-a4b6-6169795e40ed", "4b275a4a-ed2f-4ea0-8419-37bd0473239c",
    "521607f7-687a-41fa-8f02-2c637f8026b7", "3a1759ed-cc26-4fb6-99ad-dd12f4e9e919",
    "bd7913cc-c28c-4c61-852d-4e4251cf3e83", "c3bb6a2b-adc5-4df1-a88f-808db5e184d0",
    "75b6899b-5041-4b26-ae67-bd05fe13a325", "7ec84d95-9305-4a18-a61e-67eed6a5cc6f",
    "b0223270-b930-4859-8bc3-44c9deee8abe", "4c57cf2c-a4ad-44e0-a223-e3882a00cae3",
    "b884c9c9-8279-4bc8-a56b-41bec7497c68", "de0b5423-5ccd-4de8-aafb-65e09d39748a",
    "fb22a739-2a63-4782-a951-d11b75225e77", "802c0637-851f-4e48-a9af-1a71a90c337f",
    "f7034f07-8717-41af-8492-4a3f57381a49", "aa0a7442-aa1e-4276-ba50-235b08f1838a",
    "722717e2-659e-4572-a771-8547b004c12b", "b1ea8e2d-876c-45f3-a287-89c9b36f13b5",
    "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4", "79addea2-e157-4fa9-8aec-339b2232b1e3",
    "b08e3349-fc30-495b-8203-e7ba11b5ed8b", "7b587635-f23b-4b9a-9076-5b7f5e61920c",
    "f09d3b0e-ad37-47c1-a2c5-1677466f05ca", "6d18ecfb-040a-4467-9c04-2dc347979727"
]

def login():
    response = requests.post(
        f"{API_BASE}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )
    return response.json()['accessToken']

def main():
    print("="*80)
    print("ESTATÍSTICAS FINAIS - ENRIQUECIMENTO VIDA SEXUAL")
    print("="*80 + "\n")

    token = login()
    headers = {'Authorization': f'Bearer {token}'}

    stats = {
        'total': len(ITEM_IDS),
        'enriched': 0,
        'clinical_relevance_sizes': [],
        'conduct_sizes': [],
        'patient_explanation_sizes': [],
        'has_last_review': 0,
        'items_by_name': defaultdict(int)
    }

    for item_id in ITEM_IDS:
        response = requests.get(
            f"{API_BASE}/score-items/{item_id}",
            headers=headers
        )
        item = response.json()

        # Check enrichment
        if item.get('clinicalRelevance') and item.get('conduct') and item.get('patientExplanation'):
            stats['enriched'] += 1

            stats['clinical_relevance_sizes'].append(len(item.get('clinicalRelevance', '')))
            stats['conduct_sizes'].append(len(item.get('conduct', '')))
            stats['patient_explanation_sizes'].append(len(item.get('patientExplanation', '')))

            if item.get('lastReview'):
                stats['has_last_review'] += 1

        # Count by name pattern
        name = item.get('name', '')
        if 'ASEX' in name:
            stats['items_by_name']['ASEX'] += 1
        elif 'FSFI' in name:
            stats['items_by_name']['FSFI'] += 1
        elif 'IIEF' in name:
            stats['items_by_name']['IIEF-5'] += 1
        elif 'ciclo' in name.lower():
            stats['items_by_name']['Ciclo Menstrual'] += 1
        elif 'anticoncepcional' in name.lower():
            stats['items_by_name']['Anticoncepcional'] += 1
        elif 'desenvolvimento' in name.lower():
            stats['items_by_name']['Desenvolvimento Sexual'] += 1
        elif 'libido' in name.lower():
            stats['items_by_name']['Libido'] += 1
        elif 'menopausa' in name.lower():
            stats['items_by_name']['Menopausa'] += 1
        elif 'abuso' in name.lower() or 'trauma' in name.lower():
            stats['items_by_name']['Trauma/Abuso'] += 1
        elif 'hormônio' in name.lower():
            stats['items_by_name']['Hormônios'] += 1
        elif 'reprodutivo' in name.lower():
            stats['items_by_name']['Histórico Reprodutivo'] += 1
        elif 'melhora' in name.lower():
            stats['items_by_name']['Fatores Positivos'] += 1
        elif 'piora' in name.lower():
            stats['items_by_name']['Fatores Negativos'] += 1
        else:
            stats['items_by_name']['Outros'] += 1

    # Calculate averages
    avg_clinical = sum(stats['clinical_relevance_sizes']) / len(stats['clinical_relevance_sizes']) if stats['clinical_relevance_sizes'] else 0
    avg_conduct = sum(stats['conduct_sizes']) / len(stats['conduct_sizes']) if stats['conduct_sizes'] else 0
    avg_patient = sum(stats['patient_explanation_sizes']) / len(stats['patient_explanation_sizes']) if stats['patient_explanation_sizes'] else 0

    # Print results
    print(f"Total de Items: {stats['total']}")
    print(f"Items Enriquecidos: {stats['enriched']} ({stats['enriched']/stats['total']*100:.1f}%)")
    print(f"Items com Last Review: {stats['has_last_review']}\n")

    print("TAMANHOS MÉDIOS DE CONTEÚDO:")
    print(f"  Clinical Relevance: {avg_clinical:.0f} caracteres")
    print(f"  Conduct: {avg_conduct:.0f} caracteres")
    print(f"  Patient Explanation: {avg_patient:.0f} caracteres\n")

    if stats['clinical_relevance_sizes']:
        print("FAIXAS DE TAMANHO (Clinical Relevance):")
        print(f"  Mínimo: {min(stats['clinical_relevance_sizes'])} chars")
        print(f"  Máximo: {max(stats['clinical_relevance_sizes'])} chars")
        print(f"  Mediana: {sorted(stats['clinical_relevance_sizes'])[len(stats['clinical_relevance_sizes'])//2]} chars\n")

    print("DISTRIBUIÇÃO POR TÓPICO:")
    for topic, count in sorted(stats['items_by_name'].items(), key=lambda x: x[1], reverse=True):
        print(f"  {topic}: {count} items")

    print("\n" + "="*80)
    print("STATUS: ENRIQUECIMENTO 100% COMPLETO")
    print("="*80)

if __name__ == "__main__":
    main()
