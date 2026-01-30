#!/usr/bin/env python3
"""
Batch Social Items Enrichment Script
Enriquece 30 items do grupo SOCIAL com conteúdo clínico baseado em medicina funcional
Foco: Determinantes sociais da saúde, impacto ambiental, exposição a toxinas, hobbies e longevidade
"""

import os
import sys
import json
import anthropic
import requests
from typing import Dict, List, Optional

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# IDs dos 30 items SOCIAL
SOCIAL_ITEM_IDS = [
    "c84412f7-393f-41d0-8bd7-0a28824dbeb0",
    "91e450db-29df-4a78-8741-441f89630ff7",
    "7be6a88b-1fe4-4e98-9dea-203e042cf010",
    "da38d7a5-f201-4801-bb58-5cf2e4209288",
    "ac2c2df9-e65f-4156-8c39-d296c613f85c",
    "d6bc405f-bba0-4189-bc33-2b965fa75aa1",
    "66ffd35a-5ab1-42bd-a87b-414c0c91dfef",
    "eea096d8-d7d0-45ed-a757-a8fb171db7df",
    "f36df124-c3fd-4cd0-97fd-96910942289f",
    "79b3d41b-966b-4d3f-af69-8932c2b0f23a",
    "b7736f9d-5498-4ae7-8a07-f92d9edccc4d",
    "bd9b8332-7bf1-4429-8639-2ae1e4a6f749",
    "cd6b1f77-00a3-4817-9a38-3e34313af80b",
    "beabe594-930f-4b4b-83ef-d9332f4ebe31",
    "6d02ca36-a0ed-4997-a1e5-186cd9f2c0da",
    "83bb29be-8511-42af-96f2-cbefd62a6ff2",
    "df5e8c1f-168a-4984-88c6-b45afcbaebc6",
    "f335cb3d-df6b-4f37-822d-377fdaaf2fca",
    "46102736-5199-4bda-9133-8a8614dae7d8",
    "43eb760a-0917-4b00-903e-30092bc7d749",
    "95712ede-16a9-4a11-a205-06af77bfaa45",
    "c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5",
    "bc0b347d-2253-427e-859b-02ecdfbb351d",
    "3f22aa6a-632d-4dee-8d81-2166e3b44339",
    "b4776b1b-e4c1-4ca6-8f35-04d15156e7a3",
    "d85e11c0-60b5-4b7f-8252-cd1b7d54da99",
    "832c4ac0-128f-4979-b32d-ffa4afe26b88",
    "4ecdf8e0-f608-4a32-9c15-4279e86e9f92",
    "e5d1b42f-8b54-40eb-b2b1-70783bb29bd6",
    "554ee12a-9fb9-425f-9c08-22ae2b0dc9ad"
]


class SocialItemEnricher:
    def __init__(self):
        self.api_key = os.environ.get("ANTHROPIC_API_KEY")
        if not self.api_key:
            raise ValueError("ANTHROPIC_API_KEY not found in environment")

        self.client = anthropic.Anthropic(api_key=self.api_key)
        self.access_token = None

    def login(self) -> str:
        """Autentica e retorna access token"""
        response = requests.post(
            f"{API_BASE_URL}/auth/login",
            json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
        )
        response.raise_for_status()
        self.access_token = response.json()["accessToken"]
        return self.access_token

    def get_item(self, item_id: str) -> Dict:
        """Busca um score item por ID"""
        headers = {"Authorization": f"Bearer {self.access_token}"}
        response = requests.get(
            f"{API_BASE_URL}/score-items/{item_id}",
            headers=headers
        )
        response.raise_for_status()
        return response.json()

    def update_item(self, item_id: str, data: Dict) -> Dict:
        """Atualiza um score item"""
        headers = {"Authorization": f"Bearer {self.access_token}"}
        response = requests.put(
            f"{API_BASE_URL}/score-items/{item_id}",
            headers=headers,
            json=data
        )
        response.raise_for_status()
        return response.json()

    def generate_clinical_content(self, item: Dict) -> Dict:
        """
        Gera conteúdo clínico usando Claude com foco em determinantes sociais da saúde
        """

        prompt = f"""Você é um especialista em medicina funcional e determinantes sociais da saúde.

ITEM DO QUESTIONÁRIO:
- Nome: {item['name']}
- Descrição: {item.get('description', 'N/A')}
- Tipo de input: {item.get('input_type', 'N/A')}
- Grupo: SOCIAL (Determinantes Sociais da Saúde)

CONTEXTO CLÍNICO:
Este item faz parte de um questionário de medicina funcional que avalia determinantes sociais da saúde:
- Ambiente Sonoro (poluição sonora, impacto no cortisol e sono)
- Condições de Moradia (mofo, umidade, toxinas domésticas)
- Espaço para Movimento (sedentarismo ambiental)
- Exposição Ambiental Externa (poluição do ar, metais pesados)
- Hobbies e Lazer (impacto na longevidade e saúde mental)
- Luminosidade Natural (ritmo circadiano, vitamina D)
- Profissões (estresse ocupacional, exposições)

MISSÃO:
Gere conteúdo clínico profundo e baseado em evidências para este item, seguindo EXATAMENTE esta estrutura JSON:

{{
  "clinicalRelevance": "2-3 parágrafos explicando POR QUE este aspecto social/ambiental importa na medicina funcional. Cite mecanismos fisiopatológicos específicos (ex: poluição sonora > cortisol > resistência insulínica; mofo > inflamação > doenças autoimunes; falta de hobbies > isolamento social > mortalidade). Inclua dados epidemiológicos quando relevante.",

  "interpretationGuidelines": "Guia PRÁTICO para o médico interpretar as respostas. Para cada padrão de resposta, explique: (1) O que significa clinicamente, (2) Quais sistemas podem estar comprometidos, (3) Que investigações complementares considerar (ex: se exposição a mofo > dosar IgG/IgE para fungos; se trabalho noturno > avaliar cortisol e melatonina).",

  "actionableInsights": "Lista de 5-8 intervenções PRÁTICAS baseadas nas respostas. Priorize mudanças ambientais/comportamentais com melhor custo-benefício. Exemplos: 'Se ruído >70dB: protetores auriculares, cortinas acústicas, avaliar mudança'; 'Se mofo visível: desumidificador, correção estrutural, suplementação antifúngica'; 'Se <2h lazer/semana: prescrever hobbies como tratamento, grupos sociais'.",

  "redFlags": "3-5 sinais de ALERTA que exigem ação imediata ou investigação aprofundada. Exemplos: 'Exposição ocupacional a solventes orgânicos (risco neurotoxicidade)', 'Moradia com mofo + sintomas respiratórios (aspergilose?)', 'Isolamento social completo (risco suicídio)', 'Trabalho noturno >5 anos (risco câncer)'."
}}

INSTRUÇÕES:
1. Seja ESPECÍFICO: cite mecanismos bioquímicos, não generalidades
2. Base-se em EVIDÊNCIAS: mencione estudos quando relevante (ex: "Estudo Framingham mostrou...")
3. Seja ACIONÁVEL: médicos devem saber o que fazer com a informação
4. Foque em MEDICINA FUNCIONAL: causas raiz, não apenas sintomas
5. Considere DETERMINANTES SOCIAIS: como ambiente/trabalho/lazer impactam saúde

Retorne APENAS o JSON válido, sem markdown ou explicações extras."""

        try:
            message = self.client.messages.create(
                model="claude-sonnet-4-20250514",
                max_tokens=4000,
                temperature=1,
                messages=[{
                    "role": "user",
                    "content": prompt
                }]
            )

            response_text = message.content[0].text.strip()

            # Remove markdown code blocks se presentes
            if response_text.startswith("```"):
                lines = response_text.split("\n")
                response_text = "\n".join(lines[1:-1])

            content = json.loads(response_text)

            # Validação
            required_fields = [
                "clinicalRelevance",
                "interpretationGuidelines",
                "actionableInsights",
                "redFlags"
            ]

            for field in required_fields:
                if field not in content:
                    raise ValueError(f"Missing required field: {field}")

            return content

        except json.JSONDecodeError as e:
            print(f"❌ JSON decode error: {e}")
            print(f"Response: {response_text[:500]}")
            raise
        except Exception as e:
            print(f"❌ Claude API error: {e}")
            raise

    def enrich_item(self, item_id: str, index: int, total: int) -> bool:
        """Enriquece um item individual"""
        try:
            print(f"\n{'='*80}")
            print(f"[{index}/{total}] Processing item {item_id}")
            print(f"{'='*80}")

            # Busca item atual
            item = self.get_item(item_id)
            print(f"✓ Item: {item['name']}")
            print(f"  Group: {item.get('group', 'N/A')}")
            print(f"  Description: {item.get('description', 'N/A')[:100]}...")

            # Verifica se já tem conteúdo clínico
            has_content = bool(
                item.get('clinical_relevance') or
                item.get('interpretation_guidelines') or
                item.get('actionable_insights') or
                item.get('red_flags')
            )

            if has_content:
                print(f"⚠️  Item já possui conteúdo clínico. Sobrescrevendo...")

            # Gera conteúdo clínico
            print(f"🤖 Gerando conteúdo clínico com Claude...")
            clinical_content = self.generate_clinical_content(item)

            # Prepara payload de atualização
            update_payload = {
                "clinical_relevance": clinical_content["clinicalRelevance"],
                "interpretation_guidelines": clinical_content["interpretationGuidelines"],
                "actionable_insights": clinical_content["actionableInsights"],
                "red_flags": clinical_content["redFlags"]
            }

            # Atualiza item
            updated_item = self.update_item(item_id, update_payload)

            print(f"✅ Item atualizado com sucesso!")
            print(f"   Clinical Relevance: {len(clinical_content['clinicalRelevance'])} chars")
            print(f"   Interpretation: {len(clinical_content['interpretationGuidelines'])} chars")
            print(f"   Actionable Insights: {len(clinical_content['actionableInsights'])} chars")
            print(f"   Red Flags: {len(clinical_content['redFlags'])} chars")

            return True

        except Exception as e:
            print(f"❌ Erro ao processar item {item_id}: {e}")
            return False

    def run_batch(self):
        """Executa o enriquecimento em batch de todos os 30 items"""
        print(f"{'='*80}")
        print(f"BATCH SOCIAL ITEMS ENRICHMENT")
        print(f"{'='*80}")
        print(f"Total items: {len(SOCIAL_ITEM_IDS)}")
        print(f"API: {API_BASE_URL}")
        print(f"{'='*80}\n")

        # Login
        print("🔐 Autenticando...")
        self.login()
        print("✓ Autenticado com sucesso\n")

        # Processar items
        results = {
            "success": [],
            "failed": []
        }

        for index, item_id in enumerate(SOCIAL_ITEM_IDS, start=1):
            success = self.enrich_item(item_id, index, len(SOCIAL_ITEM_IDS))

            if success:
                results["success"].append(item_id)
            else:
                results["failed"].append(item_id)

        # Relatório final
        print(f"\n{'='*80}")
        print(f"RELATÓRIO FINAL")
        print(f"{'='*80}")
        print(f"✅ Sucesso: {len(results['success'])}/{len(SOCIAL_ITEM_IDS)}")
        print(f"❌ Falhas: {len(results['failed'])}/{len(SOCIAL_ITEM_IDS)}")

        if results["failed"]:
            print(f"\nItems com falha:")
            for item_id in results["failed"]:
                print(f"  - {item_id}")

        print(f"\n{'='*80}")

        # Salvar relatório
        report_path = "/home/user/plenya/SOCIAL-BATCH-REPORT.json"
        with open(report_path, "w") as f:
            json.dump(results, f, indent=2)

        print(f"📄 Relatório salvo em: {report_path}")

        return results


def main():
    try:
        enricher = SocialItemEnricher()
        results = enricher.run_batch()

        # Exit code baseado no sucesso
        if len(results["failed"]) > 0:
            sys.exit(1)
        else:
            sys.exit(0)

    except Exception as e:
        print(f"❌ Erro fatal: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
