#!/usr/bin/env python3
"""
Script para enriquecer 20 Score Items de ALIMENTAÇÃO usando lectures MFI
Foco: Glúten, Histamina, Programação Metabólica
"""

import os
import sys
import json
import requests
import anthropic
from typing import Dict, List, Optional

# Configurações
API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY")

# IDs dos 20 itens a enriquecer
TARGET_ITEMS = [
    "dc29ca47-3fa1-4ada-8cb0-85e83bcd38eb",  # Gluten 1
    "503b91b1-50d5-41f8-bc2d-772822096033",  # Gluten 2
    "a4bd778b-63f1-4b5f-9435-08411b954ffe",  # Gluten 3
    "21128901-d9bb-422f-8a95-5d4f2202fa49",  # Histamina 1
    "e76cf533-e63a-4602-98ff-5dbfa54c6079",  # Histamina 2
    "2436e6d5-945f-4646-885d-8c2df6ed6da7",  # Histamina 3
    "2a12cebb-111e-492a-8388-55222be9c346",  # Histórico familiar alimentação 1
    "64f67026-3778-461f-b5eb-4674bbb46c58",  # Histórico familiar alimentação 2
    "207b3a03-01a4-4165-9b48-a1fd0c42ae35",  # Histórico familiar alimentação 3
    "2a2e420e-1a0c-44b4-bb86-67d46e33c572",  # Hábitos mãe 1
    "73050428-0624-4b9b-8ebe-f920606d2335",  # Hábitos mãe 2
    "f946af49-7962-4371-b56e-794fcfb1d505",  # Hábitos mãe 3
    "b8401d7b-edb1-4c06-8dbd-dfa3e8754e41",  # Hábitos pai 1
    "92df69af-cf17-43fc-8a18-251bc6b8ebdf",  # Hábitos pai 2
    "8f4bbe8c-41ee-45c0-8455-c758fcc22bb0",  # Hábitos pai 3
    "2cf7a504-b4ce-4b50-9471-30fe89b19758",  # Problemas alimentares 1
    "e75bec23-9bc7-4fdc-9911-c9c9a2f48411",  # Problemas alimentares 2
    "f267b2b2-8a63-4801-9029-166aabb83176",  # Problemas alimentares 3
    "1daa1598-27a4-4750-9330-dc9967c345c6",  # Qualidade alimentação parentes 1
    "b781e715-372d-4f05-b01e-4db68c05d8db",  # Qualidade alimentação parentes 2
]


def login() -> str:
    """Faz login e retorna o access token"""
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    response.raise_for_status()
    return response.json()["accessToken"]


def get_score_item(token: str, item_id: str) -> Dict:
    """Busca um score item pelo ID"""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.get(
        f"{API_URL}/score-items/{item_id}",
        headers=headers
    )
    response.raise_for_status()
    return response.json()


def get_relevant_lectures(token: str, topic: str) -> List[Dict]:
    """Busca lectures MFI relevantes para o tópico"""
    headers = {"Authorization": f"Bearer {token}"}

    # Mapear tópico para termos de busca
    search_terms = {
        "Gluten": ["gluten", "MFI", "imunolog", "autoimun", "inflam", "intestin"],
        "Histamina": ["histamin", "MFI", "alergi", "imunolog", "inflam"],
        "Histórico familiar": ["MFI - PROGRAMAÇÃO METABÓLICA", "epigen", "genética", "familiar"],
        "Hábitos vícios mãe": ["MFI - PROGRAMAÇÃO METABÓLICA", "gestação", "maternal", "epigen"],
        "Hábitos vícios pai": ["MFI - PROGRAMAÇÃO METABÓLICA", "paternal", "epigen", "genética"],
        "Problemas alimentares": ["MFI - PROGRAMAÇÃO METABÓLICA", "nutrição", "aliment", "desenvolvimento"],
        "Qualidade alimentação": ["MFI - PROGRAMAÇÃO METABÓLICA", "nutrição", "aliment", "familiar"]
    }

    # Buscar todas as lectures (já que temos apenas 38)
    response = requests.get(
        f"{API_URL}/articles?limit=100",
        headers=headers
    )
    response.raise_for_status()
    all_articles = response.json()

    # Filtrar relevantes baseado no tópico
    relevant = []
    for article in all_articles:
        title = article.get("title", "").lower()
        # Priorizar MFI e lectures sobre programação metabólica
        if "mfi" in title or "programação metabólica" in title:
            relevant.append(article)
        elif "genética" in title or "epigenética" in title:
            relevant.append(article)
        elif "histamin" in title and "histamin" in topic.lower():
            relevant.append(article)
        elif "gluten" in title and "gluten" in topic.lower():
            relevant.append(article)

    return relevant


def enrich_with_claude(item: Dict, lectures: List[Dict]) -> Dict:
    """Usa Claude para gerar conteúdo clínico baseado nas lectures"""

    client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    # Preparar contexto das lectures
    lectures_context = ""
    for i, lecture in enumerate(lectures[:10], 1):  # Limitar a 10 lectures mais relevantes
        title = lecture.get("title", "Sem título")
        abstract = lecture.get("abstract", "")
        full_content = lecture.get("full_content", "")

        # Usar full_content se disponível, senão abstract
        content = full_content if full_content else abstract
        if content:
            # Limitar tamanho do conteúdo (primeiros 3000 chars por lecture)
            content = content[:3000]
            lectures_context += f"\n### LECTURE {i}: {title}\n{content}\n"

    if not lectures_context:
        lectures_context = "Nenhuma lecture disponível. Use conhecimento geral de medicina funcional integrativa."

    # Prompt para Claude
    prompt = f"""Você é um especialista em Medicina Funcional Integrativa. Baseado nas lectures do MFI abaixo, enriqueça este Score Item com conteúdo clínico profundo e baseado em evidências.

**SCORE ITEM:**
- Nome: {item.get('name')}
- ID: {item.get('id')}

**LECTURES DISPONÍVEIS:**
{lectures_context}

**INSTRUÇÕES:**

Gere 3 campos de texto em formato JSON:

1. **clinicalRelevance** (200-400 palavras):
   - Fisiopatologia na perspectiva da medicina funcional
   - Mecanismos bioquímicos e moleculares
   - Impacto na saúde (curto e longo prazo)
   - Epidemiologia quando relevante
   - Conexões sistêmicas (eixo intestino-cérebro, inflamação, etc)
   - USAR EVIDÊNCIAS DAS LECTURES MFI

2. **patientExplanation** (100-200 palavras):
   - Linguagem simples e acessível
   - Explicar "o que é" e "por que importa"
   - Sem jargões técnicos
   - Tom empático e educativo

3. **conduct** (150-300 palavras):
   - Protocolo clínico estruturado
   - Investigação (anamnese, exames)
   - Intervenções (dieta, estilo de vida, suplementação)
   - Monitoramento e reavaliação
   - Critérios de sucesso

**IMPORTANTE:**
- Foque especificamente no tema: {item.get('name')}
- Se o item for sobre GLÚTEN: discuta sensibilidade ao glúten não-celíaca, doença celíaca, permeabilidade intestinal, zonulina, inflamação sistêmica
- Se o item for sobre HISTAMINA: discuta intolerância à histamina, DAO, mastócitos, alimentos ricos em histamina, MCAS
- Se o item for sobre PROGRAMAÇÃO METABÓLICA (histórico familiar, hábitos parentais): discuta epigenética, janelas críticas de desenvolvimento, influências perinatais, transgeracionalidade

Retorne APENAS um objeto JSON válido com as 3 chaves: clinicalRelevance, patientExplanation, conduct.
Não inclua markdown, não inclua explicações extras. APENAS O JSON."""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4000,
            temperature=0.7,
            messages=[
                {"role": "user", "content": prompt}
            ]
        )

        # Extrair JSON da resposta
        response_text = message.content[0].text.strip()

        # Tentar parsear JSON
        # Remover possíveis markdown code blocks
        if response_text.startswith("```"):
            # Remover ```json e ```
            lines = response_text.split("\n")
            response_text = "\n".join(lines[1:-1])

        enriched_data = json.loads(response_text)

        return enriched_data

    except Exception as e:
        print(f"   ⚠️  Erro ao chamar Claude: {str(e)}")
        return None


def update_score_item(token: str, item_id: str, data: Dict) -> Dict:
    """Atualiza um score item com o conteúdo enriquecido"""
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }

    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        headers=headers,
        json=data
    )
    response.raise_for_status()
    return response.json()


def main():
    print("=" * 80)
    print("ENRIQUECIMENTO DE 20 SCORE ITEMS DE ALIMENTAÇÃO")
    print("Usando lectures MFI sobre Glúten, Histamina e Programação Metabólica")
    print("=" * 80)
    print()

    if not ANTHROPIC_API_KEY:
        print("❌ ERRO: Variável ANTHROPIC_API_KEY não configurada")
        sys.exit(1)

    # Login
    print("🔐 Fazendo login...")
    token = login()
    print("✅ Login realizado com sucesso\n")

    # Processar cada item
    results = []

    for idx, item_id in enumerate(TARGET_ITEMS, 1):
        print(f"\n{'=' * 80}")
        print(f"ITEM {idx}/{len(TARGET_ITEMS)}")
        print(f"{'=' * 80}")

        try:
            # 1. Buscar item
            print(f"📖 Buscando item {item_id}...")
            item = get_score_item(token, item_id)
            item_name = item.get("name", "Sem nome")
            print(f"   Nome: {item_name}")

            # 2. Buscar lectures relevantes
            print(f"📚 Buscando lectures MFI relevantes...")
            lectures = get_relevant_lectures(token, item_name)
            print(f"   Encontradas: {len(lectures)} lectures")

            if lectures:
                for i, lec in enumerate(lectures[:5], 1):
                    print(f"      {i}. {lec.get('title', 'Sem título')}")

            # 3. Enriquecer com Claude
            print(f"🤖 Gerando conteúdo clínico com Claude Sonnet 4.5...")
            enriched = enrich_with_claude(item, lectures)

            if not enriched:
                print(f"   ❌ Falha ao gerar conteúdo")
                results.append({
                    "item_id": item_id,
                    "name": item_name,
                    "status": "error",
                    "error": "Falha ao gerar conteúdo com Claude"
                })
                continue

            # Verificar tamanhos
            cr_len = len(enriched.get("clinicalRelevance", ""))
            pe_len = len(enriched.get("patientExplanation", ""))
            cd_len = len(enriched.get("conduct", ""))

            print(f"   ✅ Conteúdo gerado:")
            print(f"      - Clinical Relevance: {cr_len} caracteres")
            print(f"      - Patient Explanation: {pe_len} caracteres")
            print(f"      - Conduct: {cd_len} caracteres")

            # 4. Atualizar no banco
            print(f"💾 Atualizando score item...")
            updated = update_score_item(token, item_id, enriched)
            print(f"   ✅ Item atualizado com sucesso!")

            results.append({
                "item_id": item_id,
                "name": item_name,
                "status": "success",
                "lectures_used": len(lectures),
                "content_sizes": {
                    "clinical_relevance": cr_len,
                    "patient_explanation": pe_len,
                    "conduct": cd_len
                }
            })

        except Exception as e:
            print(f"   ❌ ERRO: {str(e)}")
            results.append({
                "item_id": item_id,
                "name": item.get("name", "Desconhecido") if 'item' in locals() else "Desconhecido",
                "status": "error",
                "error": str(e)
            })

    # Relatório final
    print(f"\n{'=' * 80}")
    print("RELATÓRIO FINAL")
    print(f"{'=' * 80}\n")

    success_count = sum(1 for r in results if r["status"] == "success")
    error_count = sum(1 for r in results if r["status"] == "error")

    print(f"✅ Sucesso: {success_count}/{len(TARGET_ITEMS)}")
    print(f"❌ Erros: {error_count}/{len(TARGET_ITEMS)}")
    print()

    if error_count > 0:
        print("ERROS DETALHADOS:")
        for r in results:
            if r["status"] == "error":
                print(f"  - {r['name']} ({r['item_id']}): {r.get('error', 'Erro desconhecido')}")
        print()

    # Salvar relatório JSON
    report_path = "/home/user/plenya/DIETARY-BATCH-20-REPORT.json"
    with open(report_path, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"📄 Relatório salvo em: {report_path}")
    print()

    return 0 if error_count == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
