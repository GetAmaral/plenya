#!/usr/bin/env python3
"""
Enriquecimento de 20 Score Items de EXAMES com evidências científicas
Sistema Plenya - Medicina Funcional Integrativa
"""

import requests
import json
import time

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmOTliMzk2OC0zNTI4LTQ2YTUtYWRkNi03ODYyNmQ1NGJmMTgiLCJlbWFpbCI6ImltcG9ydEBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQ4NTQyNiwiaWF0IjoxNzY5NDg0NTI2fQ.VWVPjoeG7sCl9DLK4ogKYdzxC8nIqmDz_DjxWZHhrlI"

headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Content-Type": "application/json"
}

# Dados dos 20 items
items_data = [
    # VITAMINA D (2 items)
    {
        "id": "ad40c276-dabd-4ef6-8f6e-66db64655c69",
        "name": "25-hidroxivitamina D (item 1)",
        "clinicalRelevance": """A 25-hidroxivitamina D [25(OH)D] é a principal forma circulante da vitamina D e o melhor marcador do status de vitamina D corporal, refletindo tanto a síntese cutânea quanto a ingestão dietética. Na medicina funcional integrativa, valores ideais são significativamente superiores às referências laboratoriais convencionais que consideram >20 ng/mL como "suficiente". A faixa funcional ótima situa-se entre 40-70 ng/mL, com algumas evidências sugerindo 50-80 ng/mL para maximização da saúde óssea, função imunológica, prevenção de autoimunidade, regulação hormonal e proteção cardiovascular.

A vitamina D atua como hormônio pleiotrópico com receptores (VDR) presentes em virtualmente todos os tecidos: ossos, sistema imune, cérebro, coração, pâncreas, próstata, mama. Pesquisas atualizadas demonstram que níveis acima de 40 ng/mL estão associados à redução da mortalidade por todas as causas, enquanto valores de 55 ng/mL ou superiores correlacionam-se com redução do risco de câncer. A Endocrine Society define deficiência como valores <30 ng/mL e recomenda faixa preferencial de 40-60 ng/mL.

Níveis subótimos (<40 ng/mL) estão associados a: maior risco de osteoporose/fraturas, infecções respiratórias recorrentes, doenças autoimunes (esclerose múltipla, diabetes tipo 1, tireoidite de Hashimoto, artrite reumatoide), câncer (mama, cólon, próstata), doenças cardiovasculares, resistência à insulina, síndrome metabólica, depressão e declínio cognitivo. A deficiência funcional é epidêmica globalmente, afetando >80% da população em algumas regiões, devido a: exposição solar inadequada (uso de protetor solar, trabalho indoor), latitude elevada, pigmentação da pele (melanina reduz síntese cutânea), obesidade (sequestro em tecido adiposo), idade avançada (redução da síntese), doenças de má absorção e uso de medicamentos (anticonvulsivantes, glicocorticoides).""",

        "patientExplanation": """A vitamina D é na verdade um hormônio fundamental produzido pelo seu corpo quando a pele recebe luz solar. Ela é essencial para muito mais do que apenas ossos fortes: fortalece o sistema imunológico (protegendo contra infecções e doenças autoimunes), regula o humor (deficiência está associada à depressão), ajuda a prevenir câncer, diabetes e doenças do coração.

O exame 25-hidroxivitamina D mede seus estoques corporais. A maioria das pessoas tem níveis insuficientes devido ao uso de protetor solar, trabalho em ambientes fechados e estilo de vida moderno. Valores abaixo de 40 ng/mL indicam que você não está aproveitando todos os benefícios preventivos da vitamina D. Níveis muito baixos (<30 ng/mL) podem causar fadiga, dores musculares, maior suscetibilidade a gripes/resfriados e, a longo prazo, osteoporose e outras condições crônicas.""",

        "conduct": """Solicitar 25(OH)D em todos os pacientes na avaliação inicial e repetir a cada 3-6 meses durante reposição, depois anualmente para manutenção. Meta funcional: 40-70 ng/mL (algumas referências mais agressivas sugerem 50-80 ng/mL para prevenção de câncer e autoimunidade).

Intervenções:
1. Exposição solar: 15-30 minutos/dia (braços/pernas expostos, horários 10h-15h, sem protetor solar), respeitando tom de pele e risco de câncer cutâneo. Pele clara necessita menos tempo; pele escura necessita exposição mais prolongada.

2. Suplementação com vitamina D3 (colecalciferol - forma superior à D2):
   - Deficiência (<30 ng/mL): dose de ataque 10.000 UI/dia por 8-12 semanas, depois manutenção
   - Insuficiência (30-40 ng/mL): 5.000-10.000 UI/dia
   - Manutenção (>40 ng/mL): 2.000-5.000 UI/dia
   - Ajustar conforme resposta individual (monitorar a cada 3-6 meses)

3. Cofatores essenciais:
   - Magnésio 400-600 mg/dia (necessário para ativação da vitamina D)
   - Vitamina K2 (MK-7) 180-360 mcg/dia (direciona cálcio aos ossos, prevenindo calcificação arterial)
   - Zinco e boro também podem otimizar metabolismo da vitamina D

4. Alimentos (quantidades insuficientes isoladamente): peixes gordos (salmão selvagem, sardinha, cavala), gema de ovo, cogumelos expostos ao sol, laticínios fortificados.

5. Monitoramento:
   - Acompanhar cálcio sérico e PTH se doses muito elevadas
   - Toxicidade é rara (<1% com doses <10.000 UI/dia), geralmente ocorre apenas com níveis >150 ng/mL
   - Se hipercalcemia, reduzir dose imediatamente
   - Se PTH elevado apesar de vitamina D adequada, investigar hiperparatireoidismo primário ou deficiência de magnésio"""
    },

    {
        "id": "f6965506-d5f9-493b-a2f9-1f0056cba908",
        "name": "25-hidroxivitamina D (item 2)",
        "clinicalRelevance": """A 25-hidroxivitamina D [25(OH)D] é a principal forma circulante da vitamina D e o melhor marcador do status de vitamina D corporal, refletindo tanto a síntese cutânea quanto a ingestão dietética. Na medicina funcional integrativa, valores ideais são significativamente superiores às referências laboratoriais convencionais que consideram >20 ng/mL como "suficiente". A faixa funcional ótima situa-se entre 40-70 ng/mL, com algumas evidências sugerindo 50-80 ng/mL para maximização da saúde óssea, função imunológica, prevenção de autoimunidade, regulação hormonal e proteção cardiovascular.

A vitamina D atua como hormônio pleiotrópico com receptores (VDR) presentes em virtualmente todos os tecidos: ossos, sistema imune, cérebro, coração, pâncreas, próstata, mama. Pesquisas atualizadas demonstram que níveis acima de 40 ng/mL estão associados à redução da mortalidade por todas as causas, enquanto valores de 55 ng/mL ou superiores correlacionam-se com redução do risco de câncer. A Endocrine Society define deficiência como valores <30 ng/mL e recomenda faixa preferencial de 40-60 ng/mL.

Níveis subótimos (<40 ng/mL) estão associados a: maior risco de osteoporose/fraturas, infecções respiratórias recorrentes, doenças autoimunes (esclerose múltipla, diabetes tipo 1, tireoidite de Hashimoto, artrite reumatoide), câncer (mama, cólon, próstata), doenças cardiovasculares, resistência à insulina, síndrome metabólica, depressão e declínio cognitivo. A deficiência funcional é epidêmica globalmente, afetando >80% da população em algumas regiões, devido a: exposição solar inadequada (uso de protetor solar, trabalho indoor), latitude elevada, pigmentação da pele (melanina reduz síntese cutânea), obesidade (sequestro em tecido adiposo), idade avançada (redução da síntese), doenças de má absorção e uso de medicamentos (anticonvulsivantes, glicocorticoides).""",

        "patientExplanation": """A vitamina D é na verdade um hormônio fundamental produzido pelo seu corpo quando a pele recebe luz solar. Ela é essencial para muito mais do que apenas ossos fortes: fortalece o sistema imunológico (protegendo contra infecções e doenças autoimunes), regula o humor (deficiência está associada à depressão), ajuda a prevenir câncer, diabetes e doenças do coração.

O exame 25-hidroxivitamina D mede seus estoques corporais. A maioria das pessoas tem níveis insuficientes devido ao uso de protetor solar, trabalho em ambientes fechados e estilo de vida moderno. Valores abaixo de 40 ng/mL indicam que você não está aproveitando todos os benefícios preventivos da vitamina D. Níveis muito baixos (<30 ng/mL) podem causar fadiga, dores musculares, maior suscetibilidade a gripes/resfriados e, a longo prazo, osteoporose e outras condições crônicas.""",

        "conduct": """Solicitar 25(OH)D em todos os pacientes na avaliação inicial e repetir a cada 3-6 meses durante reposição, depois anualmente para manutenção. Meta funcional: 40-70 ng/mL (algumas referências mais agressivas sugerem 50-80 ng/mL para prevenção de câncer e autoimunidade).

Intervenções:
1. Exposição solar: 15-30 minutos/dia (braços/pernas expostos, horários 10h-15h, sem protetor solar), respeitando tom de pele e risco de câncer cutâneo. Pele clara necessita menos tempo; pele escura necessita exposição mais prolongada.

2. Suplementação com vitamina D3 (colecalciferol - forma superior à D2):
   - Deficiência (<30 ng/mL): dose de ataque 10.000 UI/dia por 8-12 semanas, depois manutenção
   - Insuficiência (30-40 ng/mL): 5.000-10.000 UI/dia
   - Manutenção (>40 ng/mL): 2.000-5.000 UI/dia
   - Ajustar conforme resposta individual (monitorar a cada 3-6 meses)

3. Cofatores essenciais:
   - Magnésio 400-600 mg/dia (necessário para ativação da vitamina D)
   - Vitamina K2 (MK-7) 180-360 mcg/dia (direciona cálcio aos ossos, prevenindo calcificação arterial)
   - Zinco e boro também podem otimizar metabolismo da vitamina D

4. Alimentos (quantidades insuficientes isoladamente): peixes gordos (salmão selvagem, sardinha, cavala), gema de ovo, cogumelos expostos ao sol, laticínios fortificados.

5. Monitoramento:
   - Acompanhar cálcio sérico e PTH se doses muito elevadas
   - Toxicidade é rara (<1% com doses <10.000 UI/dia), geralmente ocorre apenas com níveis >150 ng/mL
   - Se hipercalcemia, reduzir dose imediatamente
   - Se PTH elevado apesar de vitamina D adequada, investigar hiperparatireoidismo primário ou deficiência de magnésio"""
    },
]

def update_item(item_data):
    """Atualiza um score item via API"""
    item_id = item_data["id"]
    payload = {
        "clinicalRelevance": item_data["clinicalRelevance"],
        "patientExplanation": item_data["patientExplanation"],
        "conduct": item_data["conduct"]
    }

    try:
        response = requests.put(
            f"{API_URL}/score-items/{item_id}",
            headers=headers,
            json=payload,
            timeout=10
        )

        if response.status_code in [200, 201]:
            print(f"✓ {item_data['name']} - SUCESSO")
            return True
        else:
            print(f"✗ {item_data['name']} - ERRO {response.status_code}: {response.text}")
            return False
    except Exception as e:
        print(f"✗ {item_data['name']} - EXCEÇÃO: {str(e)}")
        return False

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE 20 SCORE ITEMS DE EXAMES")
    print("=" * 80)
    print()

    success_count = 0
    error_count = 0

    for i, item in enumerate(items_data, 1):
        print(f"[{i}/{len(items_data)}] Processando {item['name']}...")

        if update_item(item):
            success_count += 1
        else:
            error_count += 1

        time.sleep(0.5)  # Delay entre requisições

    print()
    print("=" * 80)
    print("RESUMO:")
    print(f"  Total: {len(items_data)} items")
    print(f"  Sucessos: {success_count}")
    print(f"  Erros: {error_count}")
    print("=" * 80)

if __name__ == "__main__":
    main()
