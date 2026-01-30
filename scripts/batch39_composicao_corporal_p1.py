#!/usr/bin/env python3
"""
BATCH 39 - COMPOSIÇÃO CORPORAL - PARTE 1
30 items com enrichment profundo via Claude API
"""

import requests
import json
import time
import anthropic
import os
from typing import Dict, List, Optional

# API Configuration
API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# IDs dos 30 itens
ITEM_IDS = [
    "299e22bd-184e-4513-8f21-26f26d91d737",
    "9ff7a160-8ad1-4c2a-857a-53677355a631",
    "35f8405e-5bd5-4803-bade-c50e6d615582",
    "29ec8df2-5b0f-442d-aa13-1647a9759a0d",
    "119de191-7399-47e4-8c4a-e3e5f21623ff",
    "a2688922-5b23-4f7c-a842-ffff6acf081e",
    "b73bd6c3-f529-41c2-b3de-7b322663f22e",
    "924e147d-eb99-4539-a1c3-424d79f577b7",
    "9a97c090-ce0b-4dbd-a030-e652542afc6c",
    "08c992b6-3bc9-425a-b848-a9fa0bc0c0f2",
    "8fef997c-73d3-44fa-9aec-bc15f625068c",
    "f6a6515f-5488-455d-9459-8c606a13029e",
    "3be0bbdb-15bb-4899-ae09-35a811ba6c62",
    "48b082bf-3697-4c8a-a183-b2fc4396d270",
    "40a6fbbc-c6c1-4086-9622-b66d4cb67d17",
    "405b0018-088c-4a35-8688-11de766fc557",
    "01e84baf-d377-4cc4-ab36-3661b3868f1a",
    "9cbf71b3-83de-4b71-b987-120e525c790e",
    "779d1bf5-d183-4607-9c86-672badfb78e0",
    "7fe0b34c-151d-407f-9de8-6ff2492dde4b",
    "30fa255c-83d9-4cd4-8b06-75f70e2fb3eb",
    "371f12db-5a68-4092-a4c3-1430ec21f18a",
    "9e815787-7948-4030-a470-6f050a91b2f8",
    "8e0efcc9-8281-48dc-856e-fa86a845e97d",
    "ac2da1dc-80aa-4035-8703-5e669caa37a6",
    "088b4d4d-1873-45bb-8a3a-19e4463de7a5",
    "0b58623e-63d3-4685-b864-c41a0058ed56",
    "731d8307-6295-469e-95b6-361373449dcf",
    "8062469b-618a-4cfc-9239-046fd1f882e2",
    "f72f4d3f-3af8-408a-9b45-4a2988bc778a"
]

class PlenyaAPIClient:
    def __init__(self):
        self.token = None
        self.login()

    def login(self):
        """Faz login e armazena o token"""
        response = requests.post(
            f"{API_BASE}/auth/login",
            json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
        )
        response.raise_for_status()
        self.token = response.json()["accessToken"]
        print(f"✅ Login realizado com sucesso")

    def get_headers(self):
        return {
            "Authorization": f"Bearer {self.token}",
            "Content-Type": "application/json"
        }

    def get_score_item(self, item_id: str) -> Dict:
        """Busca um score item por ID"""
        response = requests.get(
            f"{API_BASE}/score-items/{item_id}",
            headers=self.get_headers()
        )
        response.raise_for_status()
        return response.json()

    def update_score_item(self, item_id: str, data: Dict) -> Dict:
        """Atualiza um score item"""
        response = requests.put(
            f"{API_BASE}/score-items/{item_id}",
            headers=self.get_headers(),
            json=data
        )
        response.raise_for_status()
        return response.json()

def generate_enrichment_prompt(item: Dict) -> str:
    """Gera prompt para enrichment do item"""

    question = item.get('question', '')
    description = item.get('description', '')

    return f"""Você é um especialista em Medicina Funcional e Composição Corporal. Enriqueça o seguinte item do sistema de avaliação com conteúdo clínico de alta qualidade em português brasileiro.

**ITEM ATUAL:**
- Pergunta: {question}
- Descrição: {description}

**SUAS TAREFAS:**

1. **CLINICAL_CONTEXT** (600-800 palavras):
   - Fisiologia e bioquímica relevante
   - Valores de referência funcionais (não apenas patológicos)
   - Diferenças por sexo, idade, etnia quando relevante
   - Relação com sarcopenia, obesidade sarcopênica, composição corporal
   - Métodos de avaliação (bioimpedância, DEXA, antropometria)
   - Impacto funcional (força, mobilidade, metabolismo)

2. **FUNCTIONAL_IMPLICATIONS** (400-500 palavras):
   - Consequências de valores subótimos
   - Relação com performance física e cognitiva
   - Impacto em longevidade e healthspan
   - Risco cardiovascular e metabólico

3. **INTERPRETATION_GUIDE** (300-400 palavras):
   - Como interpretar valores na prática clínica
   - Quando valores "normais" podem ser subótimos
   - Fatores confundidores a considerar
   - Necessidade de avaliações complementares

4. **INTERVENTION_STRATEGIES** (500-600 palavras):
   - Intervenções nutricionais (proteína, creatina, HMB, leucina)
   - Exercício resistido e HIIT
   - Suplementação baseada em evidências
   - Monitoramento e reavaliação
   - Metas funcionais realistas por faixa etária

5. **SCIENTIFIC_SOURCES** (5-8 referências):
   - Foque em: sarcopenia, obesidade sarcopênica, composição corporal, medicina funcional
   - Inclua revisões sistemáticas e guidelines recentes
   - Formate: Autor(es). Título. Revista. Ano.

**DIRETRIZES CRÍTICAS:**
- Português brasileiro técnico mas acessível
- Foco em VALORES FUNCIONAIS IDEAIS (não apenas evitar doença)
- Baseado em evidências científicas robustas
- Aplicável à prática clínica de medicina funcional
- Cite valores numéricos específicos quando possível

Retorne APENAS um JSON válido com as chaves:
{{"clinical_context": "...", "functional_implications": "...", "interpretation_guide": "...", "intervention_strategies": "...", "scientific_sources": "..."}}"""

def enrich_item_with_claude(item: Dict, client: anthropic.Anthropic) -> Dict:
    """Enriquece um item usando Claude API"""

    prompt = generate_enrichment_prompt(item)

    try:
        message = client.messages.create(
            model="claude-opus-4-5-20251101",
            max_tokens=8000,
            temperature=0.7,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        # Extrai JSON do response
        content = message.content[0].text

        # Tenta encontrar JSON no texto
        start = content.find('{')
        end = content.rfind('}') + 1

        if start != -1 and end > start:
            json_str = content[start:end]
            enrichment = json.loads(json_str)
            return enrichment
        else:
            raise ValueError("JSON não encontrado na resposta")

    except Exception as e:
        print(f"❌ Erro ao enriquecer com Claude: {e}")
        return None

def main():
    """Função principal"""

    print("=" * 80)
    print("BATCH 39 - COMPOSIÇÃO CORPORAL - PARTE 1 (30 items)")
    print("=" * 80)
    print()

    # Inicializa clientes
    api_client = PlenyaAPIClient()

    # Inicializa Claude
    claude_api_key = os.getenv("ANTHROPIC_API_KEY")
    if not claude_api_key:
        print("❌ ANTHROPIC_API_KEY não configurada")
        return

    claude_client = anthropic.Anthropic(api_key=claude_api_key)

    # Processa cada item
    results = {
        "success": [],
        "failed": [],
        "skipped": []
    }

    for idx, item_id in enumerate(ITEM_IDS, 1):
        print(f"\n[{idx}/30] Processando item {item_id}...")

        try:
            # Busca item atual
            item = api_client.get_score_item(item_id)
            question = item.get('question', 'N/A')

            print(f"  📋 Pergunta: {question[:80]}...")

            # Verifica se já tem conteúdo
            if item.get('clinical_context') and len(item.get('clinical_context', '')) > 200:
                print(f"  ⏭️  Item já enriquecido, pulando...")
                results["skipped"].append({
                    "id": item_id,
                    "question": question
                })
                continue

            # Enriquece com Claude
            print(f"  🤖 Enriquecendo com Claude Opus 4.5...")
            enrichment = enrich_item_with_claude(item, claude_client)

            if not enrichment:
                results["failed"].append({
                    "id": item_id,
                    "question": question,
                    "error": "Falha ao gerar enrichment"
                })
                continue

            # Atualiza no banco
            print(f"  💾 Salvando no banco...")
            updated = api_client.update_score_item(item_id, enrichment)

            print(f"  ✅ Item atualizado com sucesso!")
            print(f"     Clinical Context: {len(enrichment.get('clinical_context', ''))} chars")
            print(f"     Functional Implications: {len(enrichment.get('functional_implications', ''))} chars")

            results["success"].append({
                "id": item_id,
                "question": question,
                "enrichment_size": sum(len(str(v)) for v in enrichment.values())
            })

            # Rate limiting
            time.sleep(3)

        except Exception as e:
            print(f"  ❌ Erro: {e}")
            results["failed"].append({
                "id": item_id,
                "question": question if 'question' in locals() else 'N/A',
                "error": str(e)
            })

    # Relatório final
    print("\n" + "=" * 80)
    print("RELATÓRIO FINAL")
    print("=" * 80)
    print(f"✅ Sucesso: {len(results['success'])}")
    print(f"⏭️  Pulados: {len(results['skipped'])}")
    print(f"❌ Falhas: {len(results['failed'])}")

    # Salva relatório
    with open('/home/user/plenya/batch39_composicao_corporal_p1_report.json', 'w', encoding='utf-8') as f:
        json.dump(results, f, ensure_ascii=False, indent=2)

    print(f"\n📄 Relatório salvo em: batch39_composicao_corporal_p1_report.json")

if __name__ == "__main__":
    main()
