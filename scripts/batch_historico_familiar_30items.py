#!/usr/bin/env python3
"""
Batch enrichment: 30 items Histórico Familiar de Doenças
Foco: epigenética, predisposição genética, medicina de precisão
"""

import anthropic
import requests
import json
import time
import os
from typing import Dict, Any, Optional

# Config
API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"
CLAUDE_MODEL = "claude-opus-4-5-20251101"

# Verificar se API key existe
if not os.environ.get("ANTHROPIC_API_KEY"):
    print("ERROR: ANTHROPIC_API_KEY environment variable not set")
    print("Please export ANTHROPIC_API_KEY=your-key-here")
    exit(1)

# 30 IDs do grupo Histórico Familiar
ITEM_IDS = [
    "4e960d9f-a04a-4524-9049-f4559170db14",
    "26bccb28-c694-48af-8f6a-0ce6c5f73e52",
    "90d067be-f3cd-48ad-9cf9-e48fd8a2027a",
    "1c15d32d-7e79-44b8-976d-8e90ecd896be",
    "4f76ef4e-1b23-47d2-bb41-549463ad3cdf",
    "a856d80b-1443-4c16-94db-1f99508b1a9c",
    "8fb2e7de-c341-4ff5-814f-25d7695fd1cf",
    "515052d6-c7c1-40c7-942e-a1134e3aa05e",
    "8342ba54-4c2d-4b31-8619-05f4a2e86719",
    "d489c80a-50f2-4606-9579-66015e62649e",
    "394414b6-0538-4162-a68c-1e1e9d8cffef",
    "307d0403-8648-431d-88cb-2ac2422f8e86",
    "62da018d-faf4-490f-9a7b-47e0f0881bbf",
    "f4baae49-221d-428e-a218-a29391e1e26f",
    "c2c1d736-45a4-45ed-be76-e9b4704d4b1d",
    "8a552fbb-862c-4845-98dc-303f46922403",
    "12da9917-7311-4a87-89de-a1f0c8d918e7",
    "ac07f7de-eef0-420a-bdd9-cc0a3a41fbd8",
    "6491db47-381b-45fe-b483-ea0183478225",
    "35a7dbf8-2a43-4ee9-bc6b-c13122ed36c5",
    "e257d4b5-c0b2-471b-8693-defe0b2055a3",
    "77f78798-6a3a-4eef-b092-7ae1c277df4e",
    "3c52ad44-7049-40b8-b444-b05d15b96a57",
    "e2b287df-015a-4f2e-a5bf-0d99c8b24a97",
    "56fd5913-4b41-4d56-976a-b6339b4482a6",
    "1da8069d-ed8a-4b5f-83fe-0089909dd630",
    "ed899a60-9067-4a6a-aeb1-b79a08ea6062",
    "b4c8d1fe-07a2-4186-a020-051a07b98618",
    "4dc130ae-9c84-4f5f-9813-561389776254",
    "18a7fccf-4515-4807-a125-2f3e918a1d6b"
]


class HistoricoFamiliarEnricher:
    def __init__(self):
        self.token = None
        self.client = anthropic.Anthropic()
        self.results = []

    def login(self) -> bool:
        """Autentica na API"""
        try:
            response = requests.post(
                f"{API_BASE}/auth/login",
                json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
            )
            if response.status_code == 200:
                self.token = response.json()["accessToken"]
                print("✓ Login successful")
                return True
            print(f"✗ Login failed: {response.status_code}")
            return False
        except Exception as e:
            print(f"✗ Login error: {e}")
            return False

    def get_item(self, item_id: str) -> Optional[Dict[str, Any]]:
        """Busca item atual"""
        try:
            response = requests.get(
                f"{API_BASE}/score-items/{item_id}",
                headers={"Authorization": f"Bearer {self.token}"}
            )
            if response.status_code == 200:
                return response.json()
            print(f"✗ Failed to fetch item {item_id}: {response.status_code}")
            return None
        except Exception as e:
            print(f"✗ Error fetching item {item_id}: {e}")
            return None

    def enrich_with_claude(self, item: Dict[str, Any]) -> Dict[str, str]:
        """Enriquece item com Claude Opus 4.5"""

        prompt = f"""Você é um especialista em medicina de precisão, genética médica e epigenética.

ITEM ATUAL:
- Pergunta: {item.get('question', 'N/A')}
- Tipo: {item.get('question_type', 'N/A')}
- Categoria: Histórico Familiar de Doenças
- Grupo: {item.get('group_name', 'N/A')}

CONTEXTO CIENTÍFICO:
Este item avalia histórico familiar de doenças crônicas. Foco em:
1. Predisposição genética vs fatores epigenéticos modificáveis
2. Polimorfismos relevantes (APOE, MTHFR, TCF7L2, etc)
3. Interação gene-ambiente
4. Medicina preventiva personalizada baseada em histórico familiar
5. Rastreamento precoce e prevenção primária

INSTRUÇÕES:
Gere conteúdo clínico de alta qualidade (200-300 palavras) para os campos:

1. **clinical_explanation**: Explicação científica sobre a doença familiar, mecanismos de herança (monogênica vs poligênica), penetrância variável, impacto epigenético. Cite polimorfismos específicos quando aplicável (ex: APOE4 para Alzheimer, TCF7L2 para diabetes tipo 2).

2. **clinical_recommendation**: Recomendações personalizadas baseadas no histórico familiar:
   - Rastreamento precoce específico (idade de início, exames)
   - Modificações de estilo de vida para reduzir risco epigenético
   - Testes genéticos quando indicados
   - Frequência de acompanhamento médico
   - Prevenção primária baseada em evidências

3. **scientific_references**: 3-5 referências científicas REAIS e RECENTES (2020-2025):
   - Priorize: Nature Genetics, Cell, NEJM, Lancet, JAMA
   - Foco em: epigenética, polimorfismos, medicina de precisão
   - Inclua pelo menos 1 meta-análise ou guideline

FORMATO DE SAÍDA (JSON):
{{
  "clinical_explanation": "texto aqui",
  "clinical_recommendation": "texto aqui",
  "scientific_references": "Referência 1\\n\\nReferência 2\\n\\nReferência 3"
}}

IMPORTANTE:
- Seja específico sobre polimorfismos e mecanismos genéticos
- Diferencie risco genético fixo vs epigenético modificável
- Cite valores de risco relativo quando disponíveis
- Recomendações devem ser acionáveis e personalizadas
- Referências devem ser verificáveis (DOI ou PMID)
"""

        try:
            message = self.client.messages.create(
                model=CLAUDE_MODEL,
                max_tokens=4000,
                temperature=0.7,
                messages=[{"role": "user", "content": prompt}]
            )

            content = message.content[0].text

            # Parse JSON do response
            if "```json" in content:
                content = content.split("```json")[1].split("```")[0].strip()
            elif "```" in content:
                content = content.split("```")[1].split("```")[0].strip()

            enriched = json.loads(content)

            print(f"  ✓ Claude enrichment successful")
            return enriched

        except Exception as e:
            print(f"  ✗ Claude error: {e}")
            return {}

    def update_item(self, item_id: str, enriched: Dict[str, str]) -> bool:
        """Atualiza item via API"""
        try:
            response = requests.put(
                f"{API_BASE}/score-items/{item_id}",
                headers={
                    "Authorization": f"Bearer {self.token}",
                    "Content-Type": "application/json"
                },
                json=enriched
            )

            if response.status_code == 200:
                print(f"  ✓ Item updated successfully")
                return True

            print(f"  ✗ Update failed: {response.status_code} - {response.text}")
            return False

        except Exception as e:
            print(f"  ✗ Update error: {e}")
            return False

    def process_item(self, item_id: str, index: int, total: int) -> Dict[str, Any]:
        """Processa um item completo"""
        print(f"\n[{index}/{total}] Processing item {item_id}")

        result = {
            "item_id": item_id,
            "success": False,
            "error": None,
            "question": None
        }

        # 1. Buscar item
        item = self.get_item(item_id)
        if not item:
            result["error"] = "Failed to fetch item"
            return result

        result["question"] = item.get("question", "N/A")
        print(f"  Question: {result['question'][:80]}...")

        # 2. Enriquecer com Claude
        enriched = self.enrich_with_claude(item)
        if not enriched:
            result["error"] = "Claude enrichment failed"
            return result

        # 3. Atualizar via API
        if self.update_item(item_id, enriched):
            result["success"] = True
        else:
            result["error"] = "API update failed"

        return result

    def run(self):
        """Executa batch completo"""
        print("=" * 60)
        print("BATCH: 30 Items Histórico Familiar de Doenças")
        print("=" * 60)

        # Login
        if not self.login():
            print("\n✗ Authentication failed. Aborting.")
            return

        # Process all items
        total = len(ITEM_IDS)
        for index, item_id in enumerate(ITEM_IDS, 1):
            result = self.process_item(item_id, index, total)
            self.results.append(result)

            # Rate limiting
            if index < total:
                time.sleep(2)

        # Summary
        self.print_summary()
        self.save_results()

    def print_summary(self):
        """Imprime resumo"""
        print("\n" + "=" * 60)
        print("SUMMARY")
        print("=" * 60)

        successful = [r for r in self.results if r["success"]]
        failed = [r for r in self.results if not r["success"]]

        print(f"\nTotal items: {len(self.results)}")
        print(f"Successful: {len(successful)}")
        print(f"Failed: {len(failed)}")

        if failed:
            print("\nFailed items:")
            for r in failed:
                print(f"  - {r['item_id']}: {r['error']}")

    def save_results(self):
        """Salva resultados em JSON"""
        filename = "/home/user/plenya/batch_historico_familiar_30_results.json"
        with open(filename, "w", encoding="utf-8") as f:
            json.dump(self.results, f, indent=2, ensure_ascii=False)
        print(f"\n✓ Results saved to: {filename}")


if __name__ == "__main__":
    enricher = HistoricoFamiliarEnricher()
    enricher.run()
