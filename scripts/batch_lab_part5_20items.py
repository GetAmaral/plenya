#!/usr/bin/env python3
"""
BATCH LAB PART 5 - 20 Items de Exames Laboratoriais
Items: CA-125, CEA, CHCM, CIMT Carótidas, Calprotectina fecal, IST,
       Chumbo, Cilindros (Hialinos/Patológicos), Cistatina C, Cobre,
       CoQ10, Colesterol Total, Colesterol não-HDL, Colonoscopia scores
"""

import anthropic
import requests
import json
import time
from typing import Dict, Any, List

# Config
API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"
ANTHROPIC_API_KEY = "YOUR_ANTHROPIC_API_KEY"  # Will be passed via env

# IDs dos 20 items (IDs reais do banco)
ITEM_IDS = [
    "62d59b40-6dab-46c8-a6f5-53bafe3b7cae",  # CA-125
    "f36dd23c-3aa9-434f-bc86-18a3b5ee7020",  # CEA
    "eb19354a-d925-4301-bec5-2bf488501cea",  # CHCM (MCHC)
    "5c8cdefa-2c8d-4fbf-a5ca-f0ad586e1cda",  # CIMT Carótidas (Espessura Íntima-Média)
    "9de603e9-3027-4fd4-b375-5d1b9298506d",  # Calprotectina fecal
    "b6b25826-63d5-4698-87c6-b23d1390e90f",  # IST (Capacidade de fixação de ferro)
    "2a8611ab-0d4a-4a35-b3f8-569eb041ab9d",  # Chumbo
    "e4a51d97-6e52-4745-8c61-bae621bf0d3e",  # Cilindros Hialinos
    "0498349f-58d5-4c77-a25b-76c7c15c9496",  # Cilindros Patológicos
    "ae49e9b7-e1b2-4177-b8b1-d0e30f85e7fc",  # Cistatina C
    "45b91bea-accb-41a5-87aa-ac932eda0ce1",  # Cobre
    "4755b545-6065-4e41-8b7b-cb3fe0b81388",  # Coenzima Q10
    "995f731f-f891-4928-89e4-b53f27e1bb5c",  # Colesterol Total
    "aeee7b2d-0aa9-490b-bc26-e5b4ea5bc37c",  # Colesterol não-HDL
    "dd16df2d-deb3-46bc-bb4a-fdd7139863c9",  # Colonoscopia - Mayo Score UC
    "c1ea91b5-132e-419a-ae6e-5105d4df205c",  # Colonoscopia - Número Total Adenomas
    "6255db6b-9dfb-4941-a1a6-e79c8acbab17",  # Colonoscopia - SES-CD Crohn
    "fed861ee-d777-4d0a-b06d-ee1733c903e9",  # Colonoscopia (geral)
    "593303eb-371c-4b81-946c-4529d81c7e09",  # Relação Colesterol Total/HDL
    "8bf5780d-7c03-4a36-bd5b-2562ee525ab6",  # Cristais Patológicos
]


class LabEnricher:
    def __init__(self):
        self.token = None
        self.client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

    def login(self) -> str:
        """Autentica e retorna JWT token"""
        print("Fazendo login...")
        response = requests.post(
            f"{API_BASE}/auth/login",
            json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
        )
        response.raise_for_status()
        self.token = response.json()["data"]["accessToken"]
        print(f"✓ Login realizado com sucesso")
        return self.token

    def get_headers(self) -> Dict[str, str]:
        """Retorna headers com autenticação"""
        return {
            "Authorization": f"Bearer {self.token}",
            "Content-Type": "application/json"
        }

    def get_lab_item(self, item_id: str) -> Dict[str, Any]:
        """Busca um item de exame pelo ID"""
        response = requests.get(
            f"{API_BASE}/lab-test-definitions/{item_id}",
            headers=self.get_headers()
        )
        response.raise_for_status()
        return response.json()["data"]

    def search_articles(self, query: str) -> List[Dict[str, Any]]:
        """Busca artigos científicos relacionados"""
        response = requests.get(
            f"{API_BASE}/articles/search",
            params={"q": query, "limit": 10},
            headers=self.get_headers()
        )
        if response.status_code == 200:
            return response.json().get("data", [])
        return []

    def enrich_with_ai(self, item: Dict[str, Any], articles: List[Dict[str, Any]]) -> Dict[str, str]:
        """Gera conteúdo enriquecido usando Claude API"""

        # Preparar contexto dos artigos
        articles_context = "\n\n".join([
            f"Artigo: {art.get('title', 'N/A')}\nResumo: {art.get('summary', 'N/A')}"
            for art in articles[:5]
        ])

        prompt = f"""Você é um especialista em medicina laboratorial e longevidade. Analise o seguinte exame:

EXAME: {item.get('name', 'N/A')}
DESCRIÇÃO ATUAL: {item.get('description', 'N/A')}
CATEGORIA: {item.get('category', 'N/A')}

ARTIGOS CIENTÍFICOS RELACIONADOS:
{articles_context if articles_context else "Nenhum artigo disponível no sistema"}

Gere 3 textos clínicos de alta qualidade:

1. CLINICAL_SIGNIFICANCE (200-400 palavras):
   - O que este exame mede especificamente
   - Mecanismos fisiológicos/bioquímicos
   - Aplicações clínicas principais
   - Interpretação de valores alterados

2. LONGEVITY_CONTEXT (100-200 palavras):
   - Relação com envelhecimento saudável
   - Marcadores de longevidade
   - Implicações preventivas
   - Estudos de referência (se aplicável)

3. CLINICAL_RECOMMENDATIONS (150-300 palavras):
   - Quando solicitar este exame
   - Interpretação de resultados
   - Fatores que afetam valores
   - Intervenções baseadas em resultados
   - Frequência de monitoramento

IMPORTANTE:
- Linguagem técnica mas acessível a médicos
- Citar mecanismos biológicos específicos
- Focar em medicina preventiva e longevidade
- Sem emojis, sem listas numeradas excessivas
- Parágrafos coesos, texto fluído

Retorne APENAS um JSON válido:
{{
  "clinical_significance": "texto aqui...",
  "longevity_context": "texto aqui...",
  "clinical_recommendations": "texto aqui..."
}}"""

        print(f"   Gerando conteúdo AI para: {item.get('name', 'N/A')}")

        try:
            message = self.client.messages.create(
                model="claude-sonnet-4-20250514",
                max_tokens=4000,
                temperature=0.7,
                messages=[{
                    "role": "user",
                    "content": prompt
                }]
            )

            response_text = message.content[0].text.strip()

            # Extrair JSON da resposta
            if "```json" in response_text:
                response_text = response_text.split("```json")[1].split("```")[0].strip()
            elif "```" in response_text:
                response_text = response_text.split("```")[1].split("```")[0].strip()

            return json.loads(response_text)

        except Exception as e:
            print(f"   ⚠ Erro ao gerar conteúdo AI: {e}")
            return {
                "clinical_significance": "",
                "longevity_context": "",
                "clinical_recommendations": ""
            }

    def update_lab_item(self, item_id: str, enriched_data: Dict[str, str]) -> bool:
        """Atualiza item com conteúdo enriquecido"""
        try:
            response = requests.patch(
                f"{API_BASE}/lab-test-definitions/{item_id}",
                headers=self.get_headers(),
                json=enriched_data
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"   ⚠ Erro ao atualizar item: {e}")
            return False

    def process_batch(self):
        """Processa todos os 20 items"""
        print("\n" + "="*80)
        print("BATCH LAB PART 5 - 20 Items de Exames Laboratoriais")
        print("="*80 + "\n")

        results = {
            "success": [],
            "failed": [],
            "skipped": []
        }

        for idx, item_id in enumerate(ITEM_IDS, 1):
            print(f"\n[{idx}/20] Processando item: {item_id}")

            try:
                # 1. Buscar item atual
                item = self.get_lab_item(item_id)
                print(f"   Nome: {item.get('name', 'N/A')}")
                print(f"   Categoria: {item.get('category', 'N/A')}")

                # Verificar se já está enriquecido
                if item.get('clinical_significance') and len(item.get('clinical_significance', '')) > 100:
                    print(f"   ⊘ Item já possui conteúdo clínico extenso, pulando...")
                    results["skipped"].append(item_id)
                    continue

                # 2. Buscar artigos relacionados
                search_query = item.get('name', '')
                articles = self.search_articles(search_query)
                print(f"   Artigos encontrados: {len(articles)}")

                # 3. Gerar conteúdo enriquecido com AI
                enriched_data = self.enrich_with_ai(item, articles)

                # 4. Atualizar no banco
                if self.update_lab_item(item_id, enriched_data):
                    print(f"   ✓ Item atualizado com sucesso")
                    results["success"].append(item_id)
                else:
                    results["failed"].append(item_id)

                # Rate limiting
                time.sleep(2)

            except Exception as e:
                print(f"   ✗ Erro ao processar item: {e}")
                results["failed"].append(item_id)

        # Relatório final
        print("\n" + "="*80)
        print("RELATÓRIO FINAL")
        print("="*80)
        print(f"✓ Sucesso: {len(results['success'])}")
        print(f"⊘ Pulados: {len(results['skipped'])}")
        print(f"✗ Falhas: {len(results['failed'])}")
        print("="*80 + "\n")

        # Salvar relatório
        with open("/home/user/plenya/batch_lab_part5_report.json", "w") as f:
            json.dump(results, f, indent=2)

        return results


def main():
    import os

    # Pegar API key do ambiente
    api_key = os.getenv("ANTHROPIC_API_KEY")
    if not api_key:
        print("ERRO: ANTHROPIC_API_KEY não definida")
        return

    global ANTHROPIC_API_KEY
    ANTHROPIC_API_KEY = api_key

    enricher = LabEnricher()
    enricher.login()
    enricher.process_batch()


if __name__ == "__main__":
    main()
