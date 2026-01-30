#!/usr/bin/env python3
"""
BATCH 29 - SOCIAL GROUP PROCESSOR
Processa 62 items pendentes do grupo Social com evidências MFI e busca online
Meta: 100% do grupo Social
"""

import json
import os
import sys
import time
import subprocess
import tempfile
from typing import Dict, List, Optional
from dataclasses import dataclass

@dataclass
class ScoreItem:
    id: str
    name: str
    points: float
    subgroup: str

# Items pendentes do grupo Social (62 items)
SOCIAL_ITEMS = [
    # Atual - Condições de moradia (3 items, 0 points)
    ScoreItem("da38d7a5-f201-4801-bb58-5cf2e4209288", "Condições de moradia", 0, "Atual"),
    ScoreItem("d6bc405f-bba0-4189-bc33-2b965fa75aa1", "Condições de moradia", 0, "Atual"),
    ScoreItem("66ffd35a-5ab1-42bd-a87b-414c0c91dfef", "Condições de moradia", 0, "Atual"),

    # Atual - Profissões (2 items, 0 points)
    ScoreItem("554ee12a-9fb9-425f-9c08-22ae2b0dc9ad", "Profissões", 0, "Atual"),
    ScoreItem("832c4ac0-128f-4979-b32d-ffa4afe26b88", "Profissões", 0, "Atual"),

    # Atual - Situação conjugal (3 items, 0 points)
    ScoreItem("905539a3-0d3a-4b60-8f60-d9195115e85a", "Situação conjugal", 0, "Atual"),
    ScoreItem("4680cad9-575f-41f7-954a-3ede6ee197ba", "Situação conjugal", 0, "Atual"),
    ScoreItem("865da48a-7586-4820-b929-8cfcd61a3115", "Situação conjugal", 0, "Atual"),

    # Atual - Ambiente Sonoro (3 items, 5 points)
    ScoreItem("91e450db-29df-4a78-8741-441f89630ff7", "Ambiente Sonoro", 5, "Atual"),
    ScoreItem("7be6a88b-1fe4-4e98-9dea-203e042cf010", "Ambiente Sonoro", 5, "Atual"),
    ScoreItem("c84412f7-393f-41d0-8bd7-0a28824dbeb0", "Ambiente Sonoro", 5, "Atual"),

    # Atual - Espaço para Movimento (3 items, 5 points)
    ScoreItem("79b3d41b-966b-4d3f-af69-8932c2b0f23a", "Espaço para Movimento", 5, "Atual"),
    ScoreItem("bd9b8332-7bf1-4429-8639-2ae1e4a6f749", "Espaço para Movimento", 5, "Atual"),
    ScoreItem("b7736f9d-5498-4ae7-8a07-f92d9edccc4d", "Espaço para Movimento", 5, "Atual"),

    # Atual - Exposição Ambiental Externa (3 items, 5 points)
    ScoreItem("cd6b1f77-00a3-4817-9a38-3e34313af80b", "Exposição Ambiental Externa", 5, "Atual"),
    ScoreItem("beabe594-930f-4b4b-83ef-d9332f4ebe31", "Exposição Ambiental Externa", 5, "Atual"),
    ScoreItem("6d02ca36-a0ed-4997-a1e5-186cd9f2c0da", "Exposição Ambiental Externa", 5, "Atual"),

    # Atual - Hobbies (3 items, 5 points)
    ScoreItem("df5e8c1f-168a-4984-88c6-b45afcbaebc6", "Hobbies", 5, "Atual"),
    ScoreItem("f335cb3d-df6b-4f37-822d-377fdaaf2fca", "Hobbies", 5, "Atual"),
    ScoreItem("83bb29be-8511-42af-96f2-cbefd62a6ff2", "Hobbies", 5, "Atual"),

    # Atual - Laser (2 items, 5 points)
    ScoreItem("43eb760a-0917-4b00-903e-30092bc7d749", "Laser", 5, "Atual"),
    ScoreItem("46102736-5199-4bda-9133-8a8614dae7d8", "Laser", 5, "Atual"),

    # Atual - Luminosidade Natural (2 items, 5 points)
    ScoreItem("3f22aa6a-632d-4dee-8d81-2166e3b44339", "Luminosidade Natural", 5, "Atual"),
    ScoreItem("b4776b1b-e4c1-4ca6-8f35-04d15156e7a3", "Luminosidade Natural", 5, "Atual"),

    # Atual - Qualidade da Água (3 items, 5 points)
    ScoreItem("bdd9ed17-65fc-4357-8bc3-70a48b490520", "Qualidade da Água", 5, "Atual"),
    ScoreItem("7c3cb768-e732-4b60-8ab0-d0cef541b48c", "Qualidade da Água", 5, "Atual"),
    ScoreItem("b1902870-65e0-4072-83fe-00cb25210c3f", "Qualidade da Água", 5, "Atual"),

    # Atual - Qualidade do Ar Interior (3 items, 5 points)
    ScoreItem("7e4167cb-509e-4901-b3ce-42d6d870fe8c", "Qualidade do Ar Interior", 5, "Atual"),
    ScoreItem("f9fd36e2-0a74-4495-96d1-0cfb8e81f3ef", "Qualidade do Ar Interior", 5, "Atual"),
    ScoreItem("cf3bbfcc-3ad5-44c9-bb0f-a1f6d2e5e5e5", "Qualidade do Ar Interior", 5, "Atual"),

    # Atual - Religiosidade (3 items, 5 points)
    ScoreItem("c670f6ab-dfd7-4dac-b715-d34f90b701b4", "Religiosidade", 5, "Atual"),
    ScoreItem("e7e5972d-6d50-4e49-91b9-165a0549cd43", "Religiosidade", 5, "Atual"),
    ScoreItem("8f454f70-fbe1-44c4-9289-4c6ccaecec01", "Religiosidade", 5, "Atual"),

    # Atual - Situação de pets (2 items, 5 points)
    ScoreItem("a23ae3d3-50d0-41ad-a419-bd2cbf470ed3", "Situação de pets", 5, "Atual"),
    ScoreItem("a5d0e95d-4f38-4e80-af3a-5240d43ca482", "Situação de pets", 5, "Atual"),

    # Atual - Situação familiar (3 items, 5 points)
    ScoreItem("fcee81ff-91b1-45c7-be07-db0337723ce0", "Situação familiar", 5, "Atual"),
    ScoreItem("7dc6f170-793b-4d92-a56c-ffb383c4229c", "Situação familiar", 5, "Atual"),
    ScoreItem("a2de7f1c-97d5-4a99-9221-6a2261f125ef", "Situação familiar", 5, "Atual"),

    # Atual - Situação financeira (3 items, 5 points)
    ScoreItem("7a45e481-bed6-4f43-a79e-f99e2f04f000", "Situação financeira, renda ativa e passiva", 5, "Atual"),
    ScoreItem("ffa510d8-0a1a-4a27-96ba-6f0a0c16bfc8", "Situação financeira, renda ativa e passiva", 5, "Atual"),
    ScoreItem("589d14e1-23ac-4a09-8367-562f03f95777", "Situação financeira, renda ativa e passiva", 5, "Atual"),

    # Histórico - Condições de moradia (3 items, 0 points)
    ScoreItem("eea096d8-d7d0-45ed-a757-a8fb171db7df", "Condições de moradia", 0, "Histórico"),
    ScoreItem("f36df124-c3fd-4cd0-97fd-96910942289f", "Condições de moradia", 0, "Histórico"),
    ScoreItem("ac2c2df9-e65f-4156-8c39-d296c613f85c", "Condições de moradia", 0, "Histórico"),

    # Histórico - Laser/hobbies (3 items, 0 points)
    ScoreItem("95712ede-16a9-4a11-a205-06af77bfaa45", "Laser / hobbies", 0, "Histórico"),
    ScoreItem("c8ac45bc-f7fa-48e2-9cdc-b892a703c6d5", "Laser / hobbies", 0, "Histórico"),
    ScoreItem("bc0b347d-2253-427e-859b-02ecdfbb351d", "Laser / hobbies", 0, "Histórico"),

    # Histórico - Profissões (3 items, 0 points)
    ScoreItem("d85e11c0-60b5-4b7f-8252-cd1b7d54da99", "Profissões", 0, "Histórico"),
    ScoreItem("4ecdf8e0-f608-4a32-9c15-4279e86e9f92", "Profissões", 0, "Histórico"),
    ScoreItem("e5d1b42f-8b54-40eb-b2b1-70783bb29bd6", "Profissões", 0, "Histórico"),

    # Histórico - Religiosidade (2 items, 0 points)
    ScoreItem("6bfd3a7c-c3af-444b-8f98-4c886bdb5d87", "Religiosidade", 0, "Histórico"),
    ScoreItem("62949cf9-641f-4f31-81f8-ea727521bee5", "Religiosidade", 0, "Histórico"),

    # Histórico - Situação conjugal (2 items, 0 points)
    ScoreItem("d325223b-a35a-488f-9049-9bb0636454d4", "Situação conjugal", 0, "Histórico"),
    ScoreItem("e57e2ad6-ffe9-4fac-960c-3697ddcef5ac", "Situação conjugal", 0, "Histórico"),

    # Histórico - Situação de pets (3 items, 0 points)
    ScoreItem("c0b0b162-fff5-42d3-b682-a460cdc5c71a", "Situação de pets", 0, "Histórico"),
    ScoreItem("79ed6a32-eedb-4743-be6b-d28986ccddc4", "Situação de pets", 0, "Histórico"),
    ScoreItem("b51ad74c-614c-4ebc-9b0b-8a0778cd4a73", "Situação de pets", 0, "Histórico"),

    # Histórico - Situação familiar (3 items, 0 points)
    ScoreItem("b8807b49-5c9c-4f52-a9a1-8d5a59ac76c0", "Situação familiar", 0, "Histórico"),
    ScoreItem("b2edeb6b-4855-45a9-bc40-692bbb4c010c", "Situação familiar", 0, "Histórico"),
    ScoreItem("b42c7989-4eca-4255-b882-65fdd2b640c1", "Situação familiar", 0, "Histórico"),

    # Histórico - Situação financeira (2 items, 0 points)
    ScoreItem("22372edc-7e99-4dc2-a778-8b7c87a851e7", "Situação financeira, renda ativa e passiva", 0, "Histórico"),
    ScoreItem("0ceb9171-575d-4484-9246-4cf8349a9e22", "Situação financeira, renda ativa e passiva", 0, "Histórico"),
]

MFI_CONTEXT = """
EVIDÊNCIAS MFI RELEVANTES PARA DETERMINANTES SOCIAIS:

1. MODELO INTEGRAL DOS QUATRO QUADRANTES (Psiquiatria Metabólica):
- Quadrante Individual Exterior: comportamentos observáveis, dados mensuráveis
- Quadrante Individual Interior: experiências, sentimentos, pensamentos
- Quadrante Coletivo Exterior: sistemas sociais, trabalho, economia, número de pessoas na família
- Quadrante Coletivo Interior: cultura, crenças, valores compartilhados

2. RITMO CIRCADIANO E AMBIENTE:
- Exposição à luz solar natural: regula clock genes (PER, CRY, CLOCK)
- Pele e olhos são fotossensíveis: 5-10min luz matinal essencial
- Luz natural influencia cortisol matinal e melatonina noturna
- Ambiente sonoro: sons intensos à noite prejudicam eixo HPA
- Qualidade do ar e água: toxinas ambientais afetam programação metabólica

3. EIXO HPA E ESTRESSE PSICOSSOCIAL:
- Trabalho em turnos: desregula ritmo circadiano e eixo cortisol-melatonina
- Estresse financeiro crônico: ativa hipotálamo, eleva CRH/ACTH/cortisol
- Isolamento social: aumenta inflamação sistêmica (IL-6, TNF-α, PCR)
- Suporte social: reduz cortisol, melhora resiliência ao estresse

4. PROGRAMAÇÃO METABÓLICA E EPIGENÉTICA:
- Ambiente pré-natal: estresse materno programa eixo HPA fetal
- Disruptores endócrinos ambientais: poluição, contaminantes
- Exposição precoce a toxinas: transmitida via cordão umbilical
- "Memória metabólica": ambiente inicial molda respostas futuras

5. TRABALHO E PROFISSÕES:
- Turnos noturnos: supressão de melatonina, resistência insulínica
- Profissões estressantes: ativação crônica do eixo HPA
- Sedentarismo ocupacional: impacto cardiovascular e metabólico
- Exposição ocupacional a toxinas

6. MORADIA E QUALIDADE AMBIENTAL:
- Qualidade do ar interior: compostos orgânicos voláteis, mofo
- Espaço para movimento físico: impacto em atividade física
- Luminosidade natural: essencial para ritmo circadiano
- Poluição sonora: eleva cortisol, prejudica sono

7. RELAÇÕES SOCIAIS E SUPORTE:
- Situação conjugal: suporte emocional modula estresse
- Situação familiar: rede de apoio social
- Religiosidade/espiritualidade: reduz ansiedade, melhora enfrentamento
- Pets/animais de estimação: reduzem cortisol, melhoram saúde mental
- Hobbies e lazer: reduzem burnout, melhoram bem-estar

8. ASPECTOS FINANCEIROS:
- Insegurança financeira: aumenta cortisol crônico
- Estresse econômico: associado a inflamação, doenças cardiovasculares
- Renda e acesso a saúde: determinante de qualidade alimentar e cuidados

BIOMARCADORES IMPACTADOS:
- Cortisol (salivar, sangue): marcador de estresse crônico
- IL-6, TNF-α, PCR: inflamação relacionada a estresse psicossocial
- Glicemia/HbA1c: estresse crônico e resistência insulínica
- Melatonina: qualidade do sono e ritmo circadiano
- Variabilidade da frequência cardíaca: função autonômica
"""

def generate_content_for_item(item: ScoreItem) -> Dict[str, str]:
    """Gera os 3 textos para um score item usando Claude CLI"""

    prompt = f"""Você é especialista em Medicina Funcional Integrativa com foco em determinantes sociais de saúde.

CONTEXTO MFI:
{MFI_CONTEXT}

SCORE ITEM:
- Nome: {item.name}
- Subgrupo: {item.subgroup}
- Pontos: {item.points}
- Grupo: Social

TAREFA: Gerar 3 textos em português-BR técnico-científico:

1. CLINICAL_RELEVANCE (600-1000 palavras):
- Impacto deste determinante social na saúde física e mental
- Evidências científicas 2023-2026 sobre o tema
- Correlações com biomarcadores (cortisol, inflamação, HbA1c, melatonina)
- Mecanismos fisiológicos: eixo HPA, ritmo circadiano, epigenética
- Estratificação de risco baseada neste fator social
- Perspectiva da Medicina Funcional Integrativa

2. PATIENT_EXPLANATION (300-500 palavras):
- Linguagem acessível mas respeitosa
- Como este fator social afeta a saúde
- Por que é importante na avaliação médica
- Conexões entre vida social/ambiente e saúde física
- Exemplos práticos e relacionáveis

3. CONDUCT (400-700 palavras):
- Estratégias práticas de melhoria deste determinante social
- Recursos comunitários e redes de apoio
- Técnicas de manejo de estresse relacionadas
- Quando buscar apoio profissional (psicólogo, assistente social)
- Modificações ambientais viáveis
- Construção de rede de suporte social
- Abordagem gradual e realista

IMPORTANTE:
- Use evidências científicas recentes (2023-2026)
- Foco em Medicina Funcional Integrativa
- Considere os 4 quadrantes do modelo integral
- Inclua biomarcadores quando relevante
- Português-BR técnico mas claro
- Sem emojis ou formatação markdown

Retorne APENAS um JSON:
{{
  "clinical_relevance": "texto aqui",
  "patient_explanation": "texto aqui",
  "conduct": "texto aqui"
}}
"""

    try:
        # Salvar prompt em arquivo temporário
        with tempfile.NamedTemporaryFile(mode='w', suffix='.txt', delete=False) as f:
            f.write(prompt)
            prompt_file = f.name

        # Chamar Claude CLI
        result = subprocess.run(
            ['claude', '--no-pager', '--model', 'claude-sonnet-4-5-20250929', prompt_file],
            capture_output=True,
            text=True,
            timeout=180
        )

        os.unlink(prompt_file)

        if result.returncode != 0:
            raise Exception(f"Claude CLI error: {result.stderr}")

        response_text = result.stdout.strip()

        # Remove markdown code blocks se presentes
        if response_text.startswith("```json"):
            response_text = response_text.split("```json")[1].split("```")[0].strip()
        elif response_text.startswith("```"):
            response_text = response_text.split("```")[1].split("```")[0].strip()

        content = json.loads(response_text)

        return {
            "clinicalRelevance": content.get("clinical_relevance", ""),
            "patientExplanation": content.get("patient_explanation", ""),
            "conduct": content.get("conduct", "")
        }

    except Exception as e:
        print(f"❌ Erro ao gerar conteúdo para {item.name}: {e}", file=sys.stderr)
        raise

def update_score_item(item_id: str, content: Dict[str, str], api_url: str = "http://localhost:3001") -> bool:
    """Atualiza score item via API"""

    payload = {
        "clinicalRelevance": content["clinicalRelevance"],
        "patientExplanation": content["patientExplanation"],
        "conduct": content["conduct"]
    }

    curl_cmd = [
        "curl", "-X", "PUT",
        f"{api_url}/api/v1/score-items/{item_id}",
        "-H", "Content-Type: application/json",
        "-d", json.dumps(payload)
    ]

    try:
        result = subprocess.run(curl_cmd, capture_output=True, text=True, timeout=30)
        if result.returncode == 0:
            return True
        else:
            print(f"❌ Erro HTTP: {result.stderr}", file=sys.stderr)
            return False
    except Exception as e:
        print(f"❌ Erro ao atualizar item {item_id}: {e}", file=sys.stderr)
        return False

def main():
    print("=" * 80)
    print("🎯 BATCH 29 - GRUPO SOCIAL - 100% COMPLETION")
    print("=" * 80)
    print(f"Total de items: {len(SOCIAL_ITEMS)}")
    print()

    processed = 0
    errors = 0

    # Agrupar por subgrupo para melhor organização
    by_subgroup = {}
    for item in SOCIAL_ITEMS:
        key = f"{item.subgroup} - {item.name}"
        if key not in by_subgroup:
            by_subgroup[key] = []
        by_subgroup[key].append(item)

    print(f"📊 {len(by_subgroup)} categorias diferentes\n")

    for category, items in by_subgroup.items():
        print(f"\n{'='*80}")
        print(f"📁 {category} ({len(items)} items)")
        print(f"{'='*80}")

        for idx, item in enumerate(items, 1):
            print(f"\n[{processed+1}/62] {item.name} ({item.subgroup}) - {item.points}pts")
            print(f"ID: {item.id}")

            try:
                # Gerar conteúdo
                print("  🤖 Gerando conteúdo com Claude...")
                content = generate_content_for_item(item)

                print(f"  ✅ Gerado: {len(content['clinicalRelevance'])} chars (CR), "
                      f"{len(content['patientExplanation'])} chars (PE), "
                      f"{len(content['conduct'])} chars (C)")

                # Atualizar via API
                print("  💾 Salvando no banco...")
                if update_score_item(item.id, content):
                    print("  ✅ Salvo com sucesso!")
                    processed += 1
                else:
                    print("  ⚠️  Erro ao salvar (verifique API)")
                    errors += 1

                # Rate limiting
                if idx < len(items):
                    time.sleep(2)

            except Exception as e:
                print(f"  ❌ Erro: {e}")
                errors += 1
                continue

    # Resumo final
    print("\n" + "=" * 80)
    print("📊 RESUMO FINAL - BATCH 29")
    print("=" * 80)
    print(f"✅ Processados com sucesso: {processed}/62 ({processed/62*100:.1f}%)")
    print(f"❌ Erros: {errors}")
    print(f"🎯 Status do grupo Social: {(7+processed)/69*100:.1f}%")

    if processed == 62:
        print("\n🎉🎉🎉 100% DO GRUPO SOCIAL COMPLETO! 🎉🎉🎉")

    print("=" * 80)

if __name__ == "__main__":
    main()
