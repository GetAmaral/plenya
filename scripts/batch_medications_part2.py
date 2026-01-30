#!/usr/bin/env python3
"""
Script para enriquecer 15 Score Items de MEDICAMENTOS - Parte 2
Usa Claude API para gerar conteúdo clínico de alta qualidade
"""

import requests
import json
import time
import os
from typing import Dict, List, Optional
from anthropic import Anthropic

# Configuração
API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 15 items de MEDICAMENTOS (sem o Anti-hipertensivos que já está feito)
MEDICATION_ITEMS = [
    {"id": "84d98093-e62f-470f-9e08-6ab525f7d357", "name": "Antiarrítmicos"},
    {"id": "1c173e0e-3f0d-48a9-805d-b206544d27b9", "name": "Antibióticos de uso prolongado"},
    {"id": "9bc0c04e-dffe-4356-b350-182b7d4fc25a", "name": "Anticoagulantes"},
    {"id": "36575fbe-db94-4640-b590-3019cc7e12c2", "name": "Anticoncepcionais"},
    {"id": "2799ebd7-4c10-4362-9119-d0238a5c4890", "name": "Anticonvulsivantes"},
    {"id": "c77ebbe3-b387-4427-842d-04ab6960c01d", "name": "Antidepressivos"},
    {"id": "ae0c5229-76f9-4c8a-8b48-761739ac2268", "name": "Antidiabéticos orais"},
    {"id": "18faeba8-2ccb-4adf-afed-64e725f1f4c7", "name": "Antiosteoporóticos"},
    {"id": "8c91bc11-46ed-4f20-b5f0-eb1edf5f13be", "name": "Antiparkinsonianos"},
    {"id": "3a5048fe-1ff2-4f39-a8fd-1332fa906fa5", "name": "Antipsicóticos"},
    {"id": "0a0507a5-ce18-4c92-99e3-48fb31d32b5f", "name": "Antivirais"},
    {"id": "1647a98b-9883-4dad-baf7-f92b6f987100", "name": "Agonistas de GLP-1"},
    {"id": "e94f9287-d972-43ad-b6d1-ee59ed991d31", "name": "Arritmia (condição)"},
    {"id": "8b232987-3a9a-48ab-8051-04356897d97b", "name": "Artrite"},
    {"id": "2fd99065-70ad-4d33-86b8-0ac8be744547", "name": "Asma"}
]

class MedicationEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()
        self.claude_client = Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

    def login(self) -> bool:
        """Faz login e obtém token JWT"""
        try:
            response = self.session.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": EMAIL, "password": PASSWORD}
            )
            response.raise_for_status()
            data = response.json()
            self.token = data["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login realizado com sucesso!")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            return False

    def generate_clinical_content(self, medication_name: str) -> Dict[str, str]:
        """Usa Claude API para gerar conteúdo clínico de alta qualidade"""

        prompt = f"""Você é um médico especialista em medicina funcional integrativa. Gere conteúdo clínico DENSO e CIENTÍFICO sobre o uso de {medication_name}.

FORMATO DE SAÍDA (JSON):
{{
  "clinical_relevance": "...",
  "patient_explanation": "...",
  "conduct": "..."
}}

INSTRUÇÕES:

1. **clinical_relevance** (1200-1500 palavras):
   - Visão médica funcional integrativa sobre o medicamento/condição
   - Classes específicas de medicamentos e mecanismos de ação
   - Indicações principais e benefícios terapêuticos
   - Efeitos adversos DETALHADOS (comuns e graves)
   - Interações medicamentosas relevantes
   - Deficiências nutricionais induzidas pelo medicamento
   - Impacto metabólico, mitocondrial, inflamatório
   - Monitoramento necessário
   - Contraindicações absolutas e relativas
   - Cuidados especiais (idosos, gestantes, renais, hepáticos)
   - Tom: científico, técnico, denso
   - Público: médicos e profissionais de saúde

2. **patient_explanation** (400-600 palavras):
   - Explicação clara e acessível sobre o medicamento/condição
   - Por que o medicamento é usado
   - Como funciona (linguagem simples)
   - Principais benefícios esperados
   - Efeitos colaterais possíveis (sem alarmar)
   - Importância da adesão ao tratamento
   - Quando procurar o médico
   - Tom: educativo, empático, claro
   - Público: pacientes e familiares

3. **conduct** (conduta clínica - bullet points estruturados):
   - Avaliação inicial necessária
   - Exames de monitoramento (periodicidade)
   - Ajustes de dose conforme função renal/hepática
   - Suplementação para prevenir deficiências
   - Recomendações de estilo de vida
   - Sinais de alerta para suspensão
   - Interações a evitar
   - Follow-up recomendado

IMPORTANTE:
- Use terminologia médica precisa em clinical_relevance
- Seja ESPECÍFICO sobre doses, nomes genéricos, classes farmacológicas
- Inclua evidências e guidelines quando relevante
- Seja COMPASSIVO mas HONESTO em patient_explanation
- A conduct deve ser PRÁTICA e ACIONÁVEL

Retorne APENAS o JSON válido, sem markdown ou formatação extra."""

        try:
            message = self.claude_client.messages.create(
                model="claude-sonnet-4-5-20250929",
                max_tokens=8000,
                temperature=0.3,
                messages=[
                    {"role": "user", "content": prompt}
                ]
            )

            # Extrair JSON da resposta
            response_text = message.content[0].text.strip()

            # Remover markdown se presente
            if response_text.startswith("```json"):
                response_text = response_text.split("```json")[1].split("```")[0].strip()
            elif response_text.startswith("```"):
                response_text = response_text.split("```")[1].split("```")[0].strip()

            content = json.loads(response_text)

            return content

        except Exception as e:
            print(f"✗ Erro ao gerar conteúdo com Claude: {e}")
            return None

    def update_score_item(self, item_id: str, clinical_relevance: str,
                         patient_explanation: str, conduct: str) -> bool:
        """Atualiza um score item com os textos clínicos"""
        try:
            payload = {
                "clinicalRelevance": clinical_relevance,
                "patientExplanation": patient_explanation,
                "conduct": conduct
            }

            response = self.session.put(
                f"{API_BASE_URL}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"✗ Erro ao atualizar item {item_id}: {e}")
            return False

    def enrich_item(self, item: Dict) -> bool:
        """Enriquece um item específico usando Claude API"""
        print(f"\n{'='*80}")
        print(f"Processando: {item['name']}")
        print(f"ID: {item['id']}")
        print(f"{'='*80}")

        # Gerar conteúdo via Claude API
        print("  → Gerando conteúdo clínico com Claude API...")
        content = self.generate_clinical_content(item['name'])

        if not content:
            print(f"✗ Falha ao gerar conteúdo para '{item['name']}'")
            return False

        # Validar campos obrigatórios
        required_fields = ['clinical_relevance', 'patient_explanation', 'conduct']
        if not all(field in content for field in required_fields):
            print(f"✗ Conteúdo gerado incompleto para '{item['name']}'")
            return False

        # Atualizar via API
        print("  → Atualizando item via API...")
        success = self.update_score_item(
            item['id'],
            content['clinical_relevance'],
            content['patient_explanation'],
            content['conduct']
        )

        if success:
            print(f"✓ Item '{item['name']}' enriquecido com sucesso!")
            print(f"  - Clinical Relevance: {len(content['clinical_relevance'])} caracteres")
            print(f"  - Patient Explanation: {len(content['patient_explanation'])} caracteres")
            print(f"  - Conduct: {len(content['conduct'])} caracteres")
        else:
            print(f"✗ Falha ao enriquecer item '{item['name']}'")

        return success

    def process_all_items(self):
        """Processa todos os 15 items de medicamentos"""
        results = {
            "success": [],
            "failed": [],
            "total": len(MEDICATION_ITEMS)
        }

        print(f"\n{'='*80}")
        print(f"MISSÃO: ENRIQUECER 15 ITEMS DE MEDICAMENTOS - PARTE 2")
        print(f"{'='*80}\n")

        for idx, item in enumerate(MEDICATION_ITEMS, 1):
            print(f"\n[{idx}/{len(MEDICATION_ITEMS)}] Processando item...")

            success = self.enrich_item(item)

            if success:
                results["success"].append(item['name'])
            else:
                results["failed"].append(item['name'])

            # Delay entre requests para respeitar rate limits
            if idx < len(MEDICATION_ITEMS):
                print("  → Aguardando 3 segundos...")
                time.sleep(3)

        # Sumário final
        print(f"\n{'='*80}")
        print("SUMÁRIO DO PROCESSAMENTO")
        print(f"{'='*80}")
        print(f"Total de items: {results['total']}")
        print(f"Sucesso: {len(results['success'])} ({len(results['success'])/results['total']*100:.1f}%)")
        print(f"Falhas: {len(results['failed'])} ({len(results['failed'])/results['total']*100:.1f}%)")

        if results['success']:
            print(f"\n✓ Items enriquecidos:")
            for name in results['success']:
                print(f"  - {name}")

        if results['failed']:
            print(f"\n✗ Items com falha:")
            for name in results['failed']:
                print(f"  - {name}")

        print(f"{'='*80}\n")

        return results

def main():
    # Verificar se ANTHROPIC_API_KEY está configurada
    if not os.environ.get("ANTHROPIC_API_KEY"):
        print("✗ ERRO: Variável de ambiente ANTHROPIC_API_KEY não configurada!")
        print("Configure com: export ANTHROPIC_API_KEY='sua-api-key'")
        return

    enricher = MedicationEnricher()

    # Login
    if not enricher.login():
        print("✗ Erro: Não foi possível fazer login. Encerrando.")
        return

    # Processar todos items
    results = enricher.process_all_items()

    # Salvar resultados em arquivo
    output_file = '/home/user/plenya/batch_medications_part2_results.json'
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"✓ Resultados salvos em: {output_file}")

if __name__ == "__main__":
    main()
