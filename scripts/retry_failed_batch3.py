#!/usr/bin/env python3
"""
Reprocessar 3 items que falharam no batch 3
"""

import requests
import json
import time

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

FAILED_ITEMS = [
    {"id": "35225e38-b0d2-4636-a5e9-bd52087a3550", "name": "Tendinite"},
    {"id": "0846cfba-7290-4f85-b7b8-d265f252fb99", "name": "Dor lombar"},
    {"id": "e6f0b0eb-07de-469d-874f-ade1376eceea", "name": "Fibromialgia"}
]

def login():
    try:
        response = requests.post(
            f"{API_BASE_URL}/auth/login",
            json={"email": EMAIL, "password": PASSWORD}
        )
        response.raise_for_status()
        token = response.json()["accessToken"]
        print("✓ Login realizado")
        return token
    except Exception as e:
        print(f"✗ Erro login: {e}")
        return None

def update_item(token, item_id, content):
    try:
        response = requests.put(
            f"{API_BASE_URL}/score-items/{item_id}",
            json=content,
            headers={"Authorization": f"Bearer {token}"}
        )
        response.raise_for_status()
        return True
    except Exception as e:
        print(f"✗ Erro: {e}")
        return False

token = login()
if not token:
    exit(1)

contents = {
    "35225e38-b0d2-4636-a5e9-bd52087a3550": {  # Tendinite
        "clinicalRelevance": "Tendinite é inflamação/degeneração de tendões por sobrecarga, trauma repetitivo ou biomecânica inadequada. Locais comuns: ombro (manguito rotador), cotovelo (epicondilite), punho, aquileu. Tratamento: repouso, modificação de atividade, fisioterapia, fortalecimento progressivo. Na medicina funcional, suporte com nutrientes para síntese de colágeno (vitamina C, cobre, manganês), anti-inflamatórios naturais e correção de deficiências.",
        "patientExplanation": "Tendinite é inflamação dos tendões (estruturas que conectam músculos aos ossos) geralmente por uso excessivo ou movimentos repetitivos. Tratamento: repouso, gelo, fisioterapia e fortalecimento gradual. Suplementação com colágeno, vitamina C e anti-inflamatórios naturais pode acelerar recuperação.",
        "conduct": "1. Repouso relativo, modificação de atividade 2. Fisioterapia (fortalecimento excêntrico) 3. Colágeno hidrolisado (10-15g/dia) 4. Vitamina C (1-2g/dia) 5. Ômega-3, cúrcuma 6. Enzimas proteolíticas (bromelina, papaína) 7. Correção biomecânica"
    },
    "0846cfba-7290-4f85-b7b8-d265f252fb99": {  # Dor lombar
        "clinicalRelevance": "Dor lombar é extremamente prevalente (80% da população em algum momento da vida). Causas: mecânicas (tensão muscular, hérnia discal, estenose, espondilolistese), inflamatórias, infecciosas, neoplásicas, referidas. Maioria é inespecífica/mecânica, autolimitada em 4-6 semanas. Bandeiras vermelhas (déficit neurológico, síndrome de cauda equina, febre, perda de peso) exigem investigação urgente. Tratamento: manutenção de atividade, fisioterapia, exercícios de fortalecimento core. Na medicina funcional, avaliação postural, biomecânica, inflamação sistêmica.",
        "patientExplanation": "Dor lombar (dor nas costas) é muito comum e geralmente melhora em semanas. Causas: má postura, fraqueza muscular, sobrecarga, hérnia de disco. Importante: manter-se ativo (repouso prolongado piora), fazer fisioterapia, fortalecer abdômen e costas, corrigir postura. Sinais de alerta: fraqueza nas pernas, perda de controle da bexiga, febre.",
        "conduct": "1. Avaliação clínica (bandeiras vermelhas) 2. Imagem se indicação específica 3. Manter atividade, evitar repouso 4. Fisioterapia, exercícios de core 5. Quiropraxia/osteopatia 6. Magnésio, cúrcuma 7. Peso saudável 8. Ergonomia 9. Técnicas de relaxamento"
    },
    "e6f0b0eb-07de-469d-874f-ade1376eceea": {  # Fibromialgia
        "clinicalRelevance": "Fibromialgia é síndrome de dor crônica generalizada com sensibilização central do sistema nervoso, fadiga, distúrbios do sono, disfunção cognitiva (\"fibrofog\") e sintomas somáticos variados. Fisiopatologia envolve processamento anormal da dor, neuroinflamação, disfunção mitocondrial, alterações neurotransmissoras (serotonina, noradrenalina), eixo HPA. Diagnóstico clínico (critérios ACR). Tratamento multimodal: exercícios aeróbicos graduais, terapia cognitivo-comportamental, pregabalina/duloxetina. Na medicina funcional, foco em sono, estresse, inflamação sistêmica, deficiências nutricionais, saúde intestinal.",
        "patientExplanation": "Fibromialgia é dor generalizada crônica no corpo todo, acompanhada de cansaço, sono ruim e dificuldade de concentração. O sistema nervoso fica \"supersensível\" à dor. Tratamento combina: exercícios leves regulares, melhora do sono, manejo de estresse, medicamentos quando necessário. Na medicina funcional, corrigimos deficiências de vitamina D, magnésio, coenzima Q10, melhoramos sono e reduzimos inflamação.",
        "conduct": "1. Diagnóstico clínico (critérios ACR) 2. Excluir outras causas 3. Exercícios aeróbicos graduais (essencial) 4. Higiene do sono rigorosa 5. TCC, mindfulness 6. Vitamina D (>40ng/mL) 7. Magnésio 8. CoQ10 9. 5-HTP ou SAMe 10. Ômega-3 11. Dieta anti-inflamatória 12. Pregabalina ou duloxetina se necessário"
    }
}

success_count = 0
for item in FAILED_ITEMS:
    print(f"\nProcessando: {item['name']}")
    if update_item(token, item['id'], contents[item['id']]):
        print(f"✓ {item['name']} atualizado")
        success_count += 1
    time.sleep(1)

print(f"\n{'='*80}")
print(f"Resultado: {success_count}/{len(FAILED_ITEMS)} items atualizados com sucesso")
print(f"{'='*80}")
