#!/usr/bin/env python3
"""Test single social item processing"""

import json
import subprocess
import tempfile
import os

MFI_CONTEXT = """
EVIDÊNCIAS MFI - DETERMINANTES SOCIAIS DE SAÚDE:

1. RITMO CIRCADIANO (Prêmio Nobel 2017):
- Clock genes (PER, CRY, CLOCK) regulam expressão gênica
- Luz solar matinal: essencial para liberação de clock genes
- Cortisol: pico matinal, declínio ao longo do dia
- Melatonina: aumento a partir das 20h

2. EIXO HPA E ESTRESSE:
- Estresse crônico: eleva CRH→ACTH→cortisol
- Isolamento social: aumenta IL-6, TNF-α, PCR
- Suporte social: reduz cortisol, melhora resiliência

3. AMBIENTE E SAÚDE:
- Qualidade do ar: compostos orgânicos voláteis afetam saúde
- Poluição sonora: eleva cortisol, prejudica sono
- Espaço para movimento: impacto em atividade física
- Luminosidade natural: regula ritmo circadiano
"""

def test_generate():
    prompt = f"""Você é especialista em Medicina Funcional Integrativa.

{MFI_CONTEXT}

ITEM: Qualidade do Ar Interior (Atual, 5 pontos)

Gere 3 textos em português-BR:

1. CLINICAL_RELEVANCE (600-1000 palavras):
- Impacto da qualidade do ar interior na saúde
- Compostos orgânicos voláteis (VOCs), mofo, umidade
- Correlação com biomarcadores inflamatórios
- Evidências 2023-2026

2. PATIENT_EXPLANATION (300-500 palavras):
- Como ar interior afeta saúde
- Fontes de poluição doméstica
- Importância da ventilação

3. CONDUCT (400-700 palavras):
- Melhorias práticas: ventilação, plantas, filtros
- Redução de VOCs: produtos de limpeza naturais
- Controle de umidade e mofo

Retorne JSON:
{{
  "clinical_relevance": "texto",
  "patient_explanation": "texto",
  "conduct": "texto"
}}
"""

    with tempfile.NamedTemporaryFile(mode='w', suffix='.txt', delete=False) as f:
        f.write(prompt)
        prompt_file = f.name

    try:
        result = subprocess.run(
            ['claude', '-p', '--model', 'claude-sonnet-4-5-20250929', '--output-format', 'text', prompt_file],
            capture_output=True,
            text=True,
            timeout=180
        )

        os.unlink(prompt_file)

        if result.returncode != 0:
            print(f"❌ Error: {result.stderr}")
            return None

        response = result.stdout.strip()

        # Remove markdown
        if response.startswith("```json"):
            response = response.split("```json")[1].split("```")[0].strip()
        elif response.startswith("```"):
            response = response.split("```")[1].split("```")[0].strip()

        content = json.loads(response)

        print("✅ SUCCESS!")
        print(f"Clinical Relevance: {len(content['clinical_relevance'])} chars")
        print(f"Patient Explanation: {len(content['patient_explanation'])} chars")
        print(f"Conduct: {len(content['conduct'])} chars")

        return content

    except Exception as e:
        print(f"❌ Exception: {e}")
        return None

if __name__ == "__main__":
    test_generate()
