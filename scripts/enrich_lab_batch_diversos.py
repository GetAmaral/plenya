#!/usr/bin/env python3
"""
Enriquecimento de 20 items de exames laboratoriais diversos com conteúdo clínico completo.
Usa Claude para gerar conteúdo científico e estruturado.
"""

import os
import json
import anthropic
import requests
from typing import Dict, Any

# Configuração
API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJmOTliMzk2OC0zNTI4LTQ2YTUtYWRkNi03ODYyNmQ1NGJmMTgiLCJlbWFpbCI6ImltcG9ydEBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQ4NTc3NiwiaWF0IjoxNzY5NDg0ODc2fQ.uQefgudKrzfDfmC5vE_3OKR52fZCXSwXKbNn6uAaeSs"

# IDs dos 20 items
ITEM_IDS = [
    "6cb96be1-1095-4641-88cc-a403fb034c8a",  # 1
    "0ed3b126-3e60-4189-bc2c-e46b9606975a",  # 2
    "88081d50-7089-4f41-b463-c23347afedbc",  # 3
    "de7fa5ad-a023-49df-8063-8cfffa07de85",  # 4
    "7eb8dd18-6c21-4691-8c19-0f4d785af63e",  # 5
    "bc0c46b2-553a-4142-86d3-618564c66ba7",  # 6
    "d7478e09-8204-4331-82ed-d3c026f44bc6",  # 7
    "83111916-d97e-4e78-9200-0bf577c52add",  # 8
    "b3555eb3-d535-4a16-a0e5-17a5217f1bcb",  # 9
    "6b654d1e-65fd-4878-a4ec-bfd2ecf4990e",  # 10
    "4a5347f7-1031-4470-aa84-2f998162f5fc",  # 11
    "d50ef4cf-2007-4fd5-b2e0-5fa98531fcda",  # 12
    "025233d8-3dcb-4061-9a22-f8414306ece3",  # 13
    "3c8d610f-6b48-44b0-8db9-2dfefed0688e",  # 14
    "1c3e17f8-1fdf-4b9e-927c-00aa6cb9e434",  # 15
    "8c1f6aa6-0fdd-4a62-83ac-23bb9c54e052",  # 16
    "ba8c49ba-42ab-4939-adeb-6b5c1fba3c22",  # 17
    "4b9894d3-f9ff-45b5-b685-67fb9001fdb7",  # 18
    "85f9a70a-7f94-4a59-aeba-88897e8da63e",  # 19
    "151470e2-3abf-400d-adf9-a9e8e9fa8d94",  # 20
]

def get_item(item_id: str) -> Dict[str, Any]:
    """Busca item da API"""
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers={"Authorization": f"Bearer {TOKEN}"}
    )
    response.raise_for_status()
    return response.json()

def update_item(item_id: str, data: Dict[str, Any]) -> Dict[str, Any]:
    """Atualiza item na API"""
    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        headers={
            "Authorization": f"Bearer {TOKEN}",
            "Content-Type": "application/json"
        },
        json=data
    )
    response.raise_for_status()
    return response.json()

def generate_clinical_content(item_name: str, item_type: str) -> Dict[str, Any]:
    """Gera conteúdo clínico usando Claude"""

    client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

    prompt = f"""Você é um especialista em medicina laboratorial. Gere conteúdo clínico COMPLETO e CIENTIFICAMENTE PRECISO para o exame laboratorial: "{item_name}".

IMPORTANTE: Seja EXTREMAMENTE ESPECÍFICO e DETALHADO. Este conteúdo será usado por médicos em prontuários eletrônicos.

Retorne um JSON com:

{{
  "clinical_explanation": "Explicação científica COMPLETA do que é o exame, sua importância clínica, mecanismos bioquímicos/fisiológicos envolvidos (3-5 parágrafos DENSOS e TÉCNICOS com terminologia médica apropriada)",

  "low_explanation": "O que significa quando o valor está BAIXO/DIMINUÍDO (não apenas listar causas, mas EXPLICAR os mecanismos fisiopatológicos - 2-3 parágrafos DETALHADOS)",

  "high_explanation": "O que significa quando o valor está ALTO/ELEVADO (não apenas listar causas, mas EXPLICAR os mecanismos fisiopatológicos - 2-3 parágrafos DETALHADOS)",

  "clinical_significance": "Significância clínica PROFUNDA: como interpretar resultados, correlações com outros exames, importância no diagnóstico diferencial, seguimento clínico (2-3 parágrafos TÉCNICOS)",

  "interpretation_guide": "Guia PRÁTICO e DETALHADO de interpretação para médicos: valores de referência típicos, variações fisiológicas, quando solicitar, como interpretar em diferentes contextos clínicos (2-3 parágrafos APLICADOS)",

  "recommendations": [
    "Recomendação específica 1 baseada em evidências (ex: 'Solicitar em pacientes com suspeita de hepatopatia junto com TGO e TGP')",
    "Recomendação específica 2 (ex: 'Repetir exame em 3-6 meses se alterado para avaliar tendência')",
    "Recomendação específica 3 (ex: 'Correlacionar com clínica e outros marcadores antes de iniciar tratamento')",
    "Recomendação específica 4 (ex: 'Considerar interferentes: medicações, jejum, exercício físico')",
    "Recomendação específica 5 (mínimo 5 recomendações PRÁTICAS e ACIONÁVEIS)"
  ],

  "related_conditions": [
    "Condição clínica 1 relacionada (ex: 'Insuficiência hepática')",
    "Condição clínica 2 (ex: 'Desnutrição proteico-calórica')",
    "Condição clínica 3 (ex: 'Síndrome nefrótica')",
    "Condição clínica 4 (ex: 'Doença inflamatória intestinal')",
    "Condição clínica 5 (mínimo 5-8 condições RELEVANTES)"
  ],

  "patient_friendly_explanation": "Explicação para o PACIENTE em linguagem CLARA mas RESPEITOSA e COMPLETA (não infantilizada). Explicar o que é o exame, por que é importante, o que os resultados podem indicar (2-3 parágrafos acessíveis mas informativos)"
}}

CRITÉRIOS DE QUALIDADE:
- Terminologia médica PRECISA (use termos técnicos corretos)
- Informações baseadas em EVIDÊNCIAS científicas atuais
- Contexto clínico PRÁTICO e APLICÁVEL
- Profundidade ADEQUADA para profissionais de saúde
- Clareza na explicação para pacientes SEM simplificar demais

RETORNE APENAS O JSON, sem texto adicional."""

    message = client.messages.create(
        model="claude-3-7-sonnet-20250219",
        max_tokens=16000,
        temperature=1,
        messages=[{
            "role": "user",
            "content": prompt
        }]
    )

    # Extrair JSON da resposta
    content = message.content[0].text.strip()

    # Remover markdown se presente
    if content.startswith("```json"):
        content = content[7:]
    if content.startswith("```"):
        content = content[3:]
    if content.endswith("```"):
        content = content[:-3]

    return json.loads(content.strip())

def enrich_item(item_id: str, index: int) -> Dict[str, Any]:
    """Enriquece um item com conteúdo clínico"""

    print(f"\n{'='*80}")
    print(f"[{index}/20] Processando item: {item_id}")
    print(f"{'='*80}")

    # Buscar item atual
    item = get_item(item_id)
    item_name = item.get("name", "")
    item_type = item.get("type", "")

    print(f"Nome: {item_name}")
    print(f"Tipo: {item_type}")

    # Gerar conteúdo clínico
    print(f"\nGerando conteúdo clínico com Claude...")
    clinical_content = generate_clinical_content(item_name, item_type)

    # Preparar dados para atualização
    update_data = {
        "name": item_name,
        "type": item_type,
        "group_id": item.get("group_id"),
        "clinical_explanation": clinical_content["clinical_explanation"],
        "low_explanation": clinical_content["low_explanation"],
        "high_explanation": clinical_content["high_explanation"],
        "clinical_significance": clinical_content["clinical_significance"],
        "interpretation_guide": clinical_content["interpretation_guide"],
        "recommendations": clinical_content["recommendations"],
        "related_conditions": clinical_content["related_conditions"],
        "patient_friendly_explanation": clinical_content["patient_friendly_explanation"]
    }

    # Preservar campos existentes
    if item.get("options"):
        update_data["options"] = item["options"]
    if item.get("unit"):
        update_data["unit"] = item["unit"]
    if item.get("min_value") is not None:
        update_data["min_value"] = item["min_value"]
    if item.get("max_value") is not None:
        update_data["max_value"] = item["max_value"]
    if item.get("display_order") is not None:
        update_data["display_order"] = item["display_order"]

    # Atualizar item
    print(f"\nAtualizando item na API...")
    updated_item = update_item(item_id, update_data)

    print(f"✅ Item enriquecido com sucesso!")
    print(f"   - Clinical explanation: {len(clinical_content['clinical_explanation'])} chars")
    print(f"   - Low explanation: {len(clinical_content['low_explanation'])} chars")
    print(f"   - High explanation: {len(clinical_content['high_explanation'])} chars")
    print(f"   - Recommendations: {len(clinical_content['recommendations'])} items")
    print(f"   - Related conditions: {len(clinical_content['related_conditions'])} items")

    return {
        "item_id": item_id,
        "item_name": item_name,
        "success": True,
        "content": clinical_content
    }

def main():
    """Função principal"""

    print("="*80)
    print("ENRIQUECIMENTO DE 20 ITEMS DE EXAMES LABORATORIAIS DIVERSOS")
    print("="*80)
    print(f"Total de items: {len(ITEM_IDS)}")
    print(f"API: {API_URL}")
    print("="*80)

    results = []
    errors = []

    for index, item_id in enumerate(ITEM_IDS, 1):
        try:
            result = enrich_item(item_id, index)
            results.append(result)

        except Exception as e:
            print(f"\n❌ ERRO ao processar item {item_id}: {str(e)}")
            errors.append({
                "item_id": item_id,
                "error": str(e)
            })
            continue

    # Salvar resultados
    output = {
        "total_items": len(ITEM_IDS),
        "successful": len(results),
        "errors": len(errors),
        "results": results,
        "error_details": errors
    }

    output_file = "/home/user/plenya/lab_batch_diversos_results.json"
    with open(output_file, "w", encoding="utf-8") as f:
        json.dump(output, f, indent=2, ensure_ascii=False)

    # Relatório final
    print(f"\n{'='*80}")
    print("RELATÓRIO FINAL")
    print(f"{'='*80}")
    print(f"✅ Sucesso: {len(results)}/{len(ITEM_IDS)} items")
    print(f"❌ Erros: {len(errors)}/{len(ITEM_IDS)} items")
    print(f"\nResultados salvos em: {output_file}")
    print(f"{'='*80}")

if __name__ == "__main__":
    main()
