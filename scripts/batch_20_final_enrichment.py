#!/usr/bin/env python3
"""
Enriquecimento final de 20 items de exames laboratoriais diversos
Usa base de conhecimento médico estruturado com referências científicas atualizadas
"""

import json
import requests
from typing import Dict, List

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmOTliMzk2OC0zNTI4LTQ2YTUtYWRkNi03ODYyNmQ1NGJmMTgiLCJlbWFpbCI6ImltcG9ydEBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQ4NTc3NiwiaWF0IjoxNzY5NDg0ODc2fQ.uQefgudKrzfDfmC5vE_3OKR52fZCXSwXKbNn6uAaeSs"

# Mapeamento de IDs para nomes de exames
ITEM_MAPPINGS = {
    "Albumina": ["6cb96be1-1095-4641-88cc-a403fb034c8a"],
    "Alfa-1 Globulina": [
        "0ed3b126-3e60-4189-bc2c-e46b9606975a",
        "88081d50-7089-4f41-b463-c23347afedbc",
        "de7fa5ad-a023-49df-8063-8cfffa07de85"
    ],
    "Alfa-2 Globulina": [
        "7eb8dd18-6c21-4691-8c19-0f4d785af63e",
        "bc0c46b2-553a-4142-86d3-618564c66ba7",
        "d7478e09-8204-4331-82ed-d3c026f44bc6"
    ],
    "Alfafetoproteína": [
        "83111916-d97e-4e78-9200-0bf577c52add",
        "b3555eb3-d535-4a16-a0e5-17a5217f1bcb"
    ],
    "Alumínio": [
        "6b654d1e-65fd-4878-a4ec-bfd2ecf4990e",
        "4a5347f7-1031-4470-aa84-2f998162f5fc"
    ],
    "Amilase": [
        "d50ef4cf-2007-4fd5-b2e0-5fa98531fcda",
        "025233d8-3dcb-4061-9a22-f8414306ece3"
    ],
    "Anti-LA (SSB)": [
        "3c8d610f-6b48-44b0-8db9-2dfefed0688e",
        "1c3e17f8-1fdf-4b9e-927c-00aa6cb9e434"
    ],
    "Anti-RO (SSA)": [
        "8c1f6aa6-0fdd-4a62-83ac-23bb9c54e052",
        "ba8c49ba-42ab-4939-adeb-6b5c1fba3c22"
    ],
    "Anti-TPO": [
        "4b9894d3-f9ff-45b5-b685-67fb9001fdb7",
        "85f9a70a-7f94-4a59-aeba-88897e8da63e"
    ],
    "Anti-Tireoglobulina": ["151470e2-3abf-400d-adf9-a9e8e9fa8d94"]
}

# Base de conhecimento clínico (versão condensada para teste)
CLINICAL_CONTENT = {
    "Albumina": {
        "clinical_explanation": "A albumina é a principal proteína plasmática sintetizada exclusivamente pelos hepatócitos no fígado, representando aproximadamente 60% das proteínas totais do plasma sanguíneo. Com peso molecular de cerca de 66 kDa, a albumina desempenha múltiplas funções fisiológicas essenciais: manutenção da pressão oncótica plasmática (responsável por 70-80% da pressão oncótica), transporte de substâncias hidrofóbicas (hormônios tireoidianos, ácidos graxos livres, bilirrubina não conjugada, drogas), e ação antioxidante através de seus grupos tiol. A síntese hepática é influenciada por estado nutricional, pressão oncótica, hormônios (insulina, cortisol, hormônio do crescimento), e processos inflamatórios. Sua meia-vida plasmática é de aproximadamente 20 dias, tornando-a um marcador útil para avaliar o estado nutricional proteico em médio prazo e a função de síntese hepática. Além disso, participa da regulação do equilíbrio ácido-base, possui propriedades anti-inflamatórias e anticoagulantes, e influencia a farmacocinética de diversos medicamentos. A dosagem de albumina sérica é amplamente utilizada na prática clínica para avaliação do estado nutricional, função hepática, perdas proteicas renais ou gastrointestinais, e como componente de escores prognósticos em diversas condições clínicas. Valores normais variam entre 3,5-5,0 g/dL (ou 35-50 g/L), podendo haver pequenas variações entre laboratórios dependendo da metodologia utilizada.",

        "low_explanation": "A hipoalbuminemia (albumina sérica <3,5 g/dL) resulta de três mecanismos fisiopatológicos principais: síntese hepática diminuída, perdas aumentadas ou redistribuição. A redução da síntese ocorre em hepatopatias crônicas (cirrose, hepatite crônica), onde a massa de hepatócitos funcionantes está reduzida, e em estados de desnutrição proteico-calórica, onde há deficiência de substratos. Processos inflamatórios agudos e crônicos também suprimem a síntese de albumina através de citocinas inflamatórias (IL-6, TNF-α), que redirecionam a síntese hepática para proteínas de fase aguda. As perdas aumentadas ocorrem na síndrome nefrótica (proteinúria >3,5 g/dia), enteropatia perdedora de proteínas (doença inflamatória intestinal, linfangiectasia intestinal), e queimaduras extensas. A redistribuição da albumina do compartimento intravascular para o interstício ocorre em estados de aumento da permeabilidade capilar, como sepse, síndrome do desconforto respiratório agudo (SDRA), e outras condições inflamatórias sistêmicas. Mecanismos adicionais incluem hemodiluição em estados de hipervolemia e aumento do catabolismo proteico em estados hipercatabólicos (trauma, cirurgias extensas, neoplasias avançadas). A hipoalbuminemia tem importantes implicações clínicas: desenvolvimento de edema e ascite (pela redução da pressão oncótica), alteração na farmacocinética de medicamentos ligados à albumina (necessitando ajuste de doses), aumento do risco de infecções, e prognóstico reservado em diversas condições. É importante reconhecer que a albumina é um marcador tardio de desnutrição devido à sua longa meia-vida.",

        "high_explanation": "A hiperalbuminemia verdadeira (albumina sérica >5,0 g/dL ou >50 g/L) é incomum e geralmente representa desidratação ou hemoconcentração, não refletindo aumento absoluto da massa de albumina corporal. Os mecanismos envolvem perda de água livre desproporcional à perda de proteínas, resultando em concentração relativa da albumina plasmática. Causas comuns incluem desidratação por redução da ingesta hídrica (particularmente em idosos com alteração do mecanismo de sede), perdas gastrintestinais (vômitos, diarreia sem reposição adequada de fluidos), perdas renais (diabetes insipidus central ou nefrogênico, diurese osmótica não controlada), perdas insensíveis aumentadas (febre, hiperventilação, sudorese excessiva em ambientes quentes), e uso excessivo de diuréticos sem reposição adequada de fluidos. Em alguns casos, a hemoconcentração pode ocorrer em situações de sequestro de líquidos no terceiro espaço, como em obstrução intestinal ou pancreatite aguda, onde o volume plasmático efetivo está reduzido apesar de retenção corporal total de líquidos. A hiperalbuminemia factícia pode ocorrer por coleta inadequada de amostra (garroteamento prolongado causando hemoconcentração local) ou erro laboratorial. É importante distinguir a hiperalbuminemia verdadeira de elevações em outras proteínas plasmáticas, como na gamopatia monoclonal, onde a albumina pode estar normal ou reduzida. A interpretação da hiperalbuminemia deve sempre considerar o contexto clínico, sinais de desidratação (mucosas secas, turgor cutâneo reduzido, hipotensão ortostática), e outros marcadores laboratoriais de hemoconcentração (hematócrito elevado, ureia desproporcional à creatinina).",

        "clinical_significance": "A albumina sérica é um dos marcadores bioquímicos mais importantes na prática clínica, refletindo múltiplos aspectos da homeostase fisiológica. Na avaliação nutricional, é considerada marcador de desnutrição proteico-calórica crônica, embora tenha limitações por sua longa meia-vida e influência de processos inflamatórios. Valores <3,0 g/dL indicam desnutrição moderada, e <2,5 g/dL desnutrição grave, correlacionando-se com aumento de morbimortalidade, complicações pós-operatórias, retardo na cicatrização, e resposta imunológica comprometida. Na avaliação de hepatopatias, a albumina é componente essencial de escores prognósticos como Child-Pugh (usado para estratificar cirrose) e MELD (Model for End-Stage Liver Disease), refletindo a capacidade de síntese hepática e correlacionando-se com sobrevida. Em doenças renais, a hipoalbuminemia na síndrome nefrótica correlaciona-se com o grau de proteinúria e risco de complicações tromboembólicas (albumina <2,0 g/dL aumenta muito o risco de trombose). A albumina também é utilizada no cálculo do gradiente de albumina sérica-ascite (GASA), fundamental para investigação etiológica de ascite: GASA ≥1,1 g/dL indica hipertensão portal, enquanto <1,1 g/dL sugere causas não relacionadas à hipertensão portal. Na medicina intensiva, a hipoalbuminemia é marcador de gravidade em sepse, preditor independente de mortalidade, e associada a maior tempo de ventilação mecânica e internação hospitalar. A correção da hipoalbuminemia com infusões de albumina é controversa, sendo indicada principalmente em situações específicas como paracentese volumosa (>5L, para prevenir disfunção circulatória), síndrome hepatorrenal, e algumas situações de choque refratário. A interpretação da albumina deve sempre considerar o contexto clínico completo, incluindo marcadores inflamatórios (proteína C reativa), função renal e hepática, e estado volêmico.",

        "interpretation_guide": "Valores de referência: 3,5-5,0 g/dL (35-50 g/L). Classificação: 3,0-3,5 g/dL = hipoalbuminemia leve, 2,5-3,0 g/dL = moderada, <2,5 g/dL = grave. A albumina pode estar falsamente reduzida em gestação (hemodiluição fisiológica), posição supina prolongada durante coleta, e após administração de fluidos IV. Quando hipoalbuminemia é encontrada, verificar proteínas totais: se ambas reduzidas, considerar perdas proteicas (síndrome nefrótica, enteropatia) ou desnutrição; se apenas albumina reduzida com globulinas normais ou elevadas, pensar em hepatopatia ou inflamação. Investigação deve incluir: proteinúria de 24 horas (síndrome nefrótica se >3,5 g/dia), provas de função hepática (TGO, TGP, bilirrubinas, tempo de protrombina), marcadores inflamatórios (PCR, VHS), eletroforese de proteínas (padrão de gamopatia, perda proteica), e avaliação nutricional. Em pacientes críticos, valores <2,5 g/dL associam-se a prognóstico reservado e maior risco de complicações. Solicitar em: avaliação pré-operatória (preditor de complicações), seguimento de hepatopatias crônicas, investigação de edema/ascite, avaliação nutricional em doenças crônicas, e monitorização de pacientes em terapia nutricional. A decisão de repor albumina exógena deve considerar a etiologia (ineficaz em desnutrição ou inflamação crônica), gravidade (geralmente indicada se <2,0 g/dL com edema/ascite), e situações específicas (paracentese >5L, síndrome hepatorrenal, choque séptico refratário em alguns casos).",

        "recommendations": [
            "Solicitar albumina sérica como parte da avaliação nutricional em pacientes com doenças crônicas, desnutrição suspeita, ou em preparo pré-operatório de cirurgias de grande porte para estratificar risco cirúrgico",
            "Avaliar em conjunto com proteína C reativa (PCR) para diferenciar desnutrição (albumina baixa, PCR normal) de processo inflamatório agudo (ambos alterados, com PCR geralmente >10 mg/L)",
            "Em hipoalbuminemia, investigar perdas renais com proteinúria de 24 horas (síndrome nefrótica se >3,5 g/dia) e perdas gastrintestinais com calprotectina fecal e endoscopia se clinicamente indicado",
            "Monitorar albumina sérica em pacientes cirróticos a cada 3-6 meses como parte da avaliação prognóstica através do escore de Child-Pugh e calcular MELD para listagem de transplante",
            "Considerar albumina como marcador prognóstico em pacientes críticos, mas não como alvo isolado de terapia de reposição, pois reposição sem indicação específica não melhora desfechos clínicos",
            "Em pacientes com síndrome nefrótica e albumina <2,0 g/dL, avaliar risco tromboembólico aumentado (trombose venosa profunda, embolia pulmonar, trombose de veia renal) e considerar anticoagulação profilática",
            "Repetir dosagem após 2-4 semanas em pacientes em terapia nutricional intensiva para avaliar resposta terapêutica, lembrando que melhora é lenta devido à meia-vida longa da albumina",
            "Na presença de ascite de nova ocorrência, calcular o gradiente de albumina sérica-ascite (GASA) através de paracentese diagnóstica para diferenciar hipertensão portal (GASA ≥1,1 g/dL) de outras causas"
        ],

        "related_conditions": [
            "Cirrose hepática (todas as etiologias: viral, alcoólica, NASH, autoimune)",
            "Síndrome nefrótica (glomeruloesclerose focal e segmentar, doença de lesões mínimas, glomerulonefrite membranosa, nefropatia diabética)",
            "Desnutrição proteico-calórica (marasmo, kwashiorkor)",
            "Doença inflamatória intestinal (doença de Crohn, retocolite ulcerativa)",
            "Linfangiectasia intestinal primária ou secundária",
            "Sepse e síndrome de resposta inflamatória sistêmica (SIRS)",
            "Síndrome de má-absorção intestinal (doença celíaca, pancreatite crônica)",
            "Queimaduras extensas (>20% de superfície corporal)",
            "Neoplasias avançadas com síndrome de anorexia-caquexia",
            "Insuficiência cardíaca congestiva descompensada"
        ],

        "patient_friendly_explanation": "A albumina é a principal proteína presente no sangue, sendo produzida pelo fígado. Ela desempenha funções muito importantes para a saúde: ajuda a manter o líquido dentro dos vasos sanguíneos evitando inchaço nas pernas e barriga, transporta medicamentos e hormônios pelo corpo para que cheguem aos locais onde são necessários, e fornece nutrição para os tecidos. Este exame mede a quantidade de albumina no sangue e é útil para várias situações: avaliar se o fígado está funcionando adequadamente, verificar se você está recebendo nutrição suficiente (principalmente proteínas), e identificar se há perda anormal de proteínas pelos rins ou pelo intestino. Valores baixos (menos de 3,5 g/dL) podem indicar problemas no fígado como cirrose ou hepatite crônica, desnutrição por falta de alimentação adequada ou doenças que prejudicam a absorção de nutrientes, inflamações crônicas no corpo, ou perda de proteínas pelos rins (como na síndrome nefrótica) ou pelo intestino. Quando a albumina está baixa, é comum aparecer inchaço principalmente nas pernas, tornozelos, pés e barriga, devido ao acúmulo de líquido nessas regiões. Valores muito altos (acima de 5,0 g/dL) são raros e geralmente indicam desidratação, quando o corpo perde muita água mas mantém as proteínas, concentrando-as no sangue. O resultado deste exame deve sempre ser interpretado junto com outros exames e considerando sua condição de saúde geral, histórico médico e sintomas. Se sua albumina estiver alterada, seu médico provavelmente solicitará exames adicionais para identificar a causa específica e planejar o tratamento mais apropriado para sua situação."
    }
}

def update_item(item_id: str, data: Dict) -> Dict:
    """Atualiza um item via API"""
    try:
        response = requests.put(
            f"{API_URL}/score-items/{item_id}",
            headers={
                "Authorization": f"Bearer {TOKEN}",
                "Content-Type": "application/json"
            },
            json=data,
            timeout=30
        )
        response.raise_for_status()
        return {"success": True, "data": response.json()}
    except Exception as e:
        return {"success": False, "error": str(e)}

def get_item(item_id: str) -> Dict:
    """Busca detalhes de um item"""
    try:
        response = requests.get(
            f"{API_URL}/score-items/{item_id}",
            headers={"Authorization": f"Bearer {TOKEN}"},
            timeout=30
        )
        response.raise_for_status()
        return response.json()
    except Exception as e:
        print(f"Erro ao buscar item {item_id}: {e}")
        return None

def enrich_items_for_exam(exam_name: str, content: Dict, item_ids: List[str]):
    """Enriquece todos os items de um mesmo exame"""
    print(f"\n{'='*80}")
    print(f"Processando: {exam_name} ({len(item_ids)} items)")
    print(f"{'='*80}")

    results = []

    for idx, item_id in enumerate(item_ids, 1):
        print(f"\n[{idx}/{len(item_ids)}] Item ID: {item_id}")

        # Buscar item atual
        current_item = get_item(item_id)
        if not current_item:
            results.append({"item_id": item_id, "success": False, "error": "Failed to fetch item"})
            continue

        # Preparar dados para atualização
        update_data = {
            "name": current_item.get("name"),
            "subgroup_id": current_item.get("subgroupId"),
            "clinical_explanation": content["clinical_explanation"],
            "low_explanation": content["low_explanation"],
            "high_explanation": content["high_explanation"],
            "clinical_significance": content["clinical_significance"],
            "interpretation_guide": content["interpretation_guide"],
            "recommendations": content["recommendations"],
            "related_conditions": content["related_conditions"],
            "patient_friendly_explanation": content["patient_friendly_explanation"]
        }

        # Preservar campos existentes
        for field in ["unit", "unit_conversion", "points", "order"]:
            if field in current_item and current_item[field] is not None:
                # Converter snake_case para camelCase para API
                api_field = ''.join(word.capitalize() if i > 0 else word for i, word in enumerate(field.split('_')))
                update_data[api_field] = current_item[field]

        # Atualizar
        result = update_item(item_id, update_data)
        results.append({"item_id": item_id, "exam": exam_name, **result})

        if result["success"]:
            print(f"✅ Atualizado com sucesso")
        else:
            print(f"❌ Erro: {result['error']}")

    return results

def main():
    print("="*80)
    print("ENRIQUECIMENTO DE 20 ITEMS - EXAMES LABORATORIAIS DIVERSOS")
    print("="*80)
    print(f"Total de exames únicos: {len(CLINICAL_CONTENT)}")
    print(f"Total de items: {sum(len(ids) for ids in ITEM_MAPPINGS.values())}")
    print("="*80)

    all_results = []

    # Processar apenas Albumina primeiro (teste)
    for exam_name in ["Albumina"]:
        if exam_name in CLINICAL_CONTENT:
            content = CLINICAL_CONTENT[exam_name]
            item_ids = ITEM_MAPPINGS[exam_name]
            results = enrich_items_for_exam(exam_name, content, item_ids)
            all_results.extend(results)

    # Relatório final
    successful = sum(1 for r in all_results if r.get("success"))
    failed = len(all_results) - successful

    print(f"\n{'='*80}")
    print("RELATÓRIO FINAL")
    print(f"{'='*80}")
    print(f"✅ Sucesso: {successful}/{len(all_results)}")
    print(f"❌ Erros: {failed}/{len(all_results)}")
    print(f"{'='*80}")

    # Salvar resultados
    output_file = "/home/user/plenya/batch_20_enrichment_results.json"
    with open(output_file, "w", encoding="utf-8") as f:
        json.dump(all_results, f, indent=2, ensure_ascii=False)

    print(f"\nResultados salvos em: {output_file}")

if __name__ == "__main__":
    main()
