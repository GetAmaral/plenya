#!/usr/bin/env python3
"""
Batch de enriquecimento - ALIMENTAÇÃO Parte 3
20 items sobre: Ordem alimentos, Perfil metabólico parental, Piores períodos, Preferências, Proteína do leite
"""

import os
import json
from anthropic import Anthropic

# Configuração
API_KEY = os.getenv("ANTHROPIC_API_KEY", "sk-ant-api03-9gfxXKJOzQOPpTEK7j0NvZl4BVgaMKSO6T1ivfvXRBRg-mHmD14jPLt8IqQXqYKE1YoBOr_dAG3v2DgEeWcT8w-cJ0ZLAAA")

client = Anthropic(api_key=API_KEY)

# Items agrupados por tema
ITEMS = {
    "ordem_alimentos": [
        "4f6132f2-ab68-4806-83e1-e5834eedd1b7",
        "07943b6c-e3e1-4c12-8d6a-646d9371ac28",
        "20ab6972-c807-4138-bc39-4f0bf027205f"
    ],
    "outros": [
        "028d0bc0-63bc-42f9-a7d1-d12720f1adf0",
        "5a4202a3-67d3-470b-b1ed-163d099ccc5a",
        "96186d86-e1d7-4c7f-be37-1da17755d6cb"
    ],
    "perfil_metabolico_mae": [
        "06f06535-b63a-4cc0-a851-8741fda0f082",
        "808a8a68-91e9-48a6-9756-58b01b8c8714",
        "1be1f170-6e5a-4372-bb2f-086e146821a7"
    ],
    "perfil_metabolico_pai": [
        "3c598f1d-67b2-4b08-9b9b-a6cc8f6e5f25",
        "dc6c7bf0-abc3-42b1-a6cc-b857b577810f",
        "bebe4239-cdf0-409c-82a2-cd6d3ef5df55"
    ],
    "piores_periodos": [
        "7f15bca0-9060-41b1-b457-5ecfae0ce087",
        "6fda89f8-00d3-4cb0-b262-de74d244308d",
        "4f8057e2-a99f-4565-83ad-f6ac3335460d"
    ],
    "preferencias": [
        "f1911a45-4e01-4d88-b409-304088f50254",
        "fd13405b-264e-4640-89f5-07496160171e",
        "1743b33b-79c9-4ef7-9b31-afec3e4562e2"
    ],
    "proteina_leite": [
        "91067f64-2acf-47b3-b390-90707003035d",
        "3d3335c8-39d0-4664-a20d-37666632c3ba"
    ]
}

def generate_content(topic: str, item_name: str) -> dict:
    """Gera conteúdo clínico para um item específico usando Claude"""

    prompts = {
        "ordem_alimentos": f"""Você é um nutricionista clínico especializado. Gere conteúdo educacional sobre a ordem de consumo dos alimentos nas refeições.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a importância clínica da ordem de ingestão dos alimentos, incluindo:
   - Impacto na glicemia pós-prandial e curva glicêmica
   - Efeito na saciedade e controle do apetite
   - Influência na absorção de nutrientes
   - Relação com síndrome metabólica e diabetes tipo 2
   - Conceito de "food sequencing" baseado em evidências

2. **patient_explanation** (100-150 palavras): Linguagem acessível explicando:
   - Por que a ordem importa na prática
   - Exemplo prático de sequência ideal (fibras → proteínas → carboidratos)
   - Benefícios para controle de peso e energia
   - Como implementar no dia a dia

3. **conduct** (120-180 palavras): Orientações práticas incluindo:
   - Sequência recomendada: salada/vegetais primeiro, depois proteína, por último carboidratos
   - Tempo entre cada grupo alimentar (pausas de 5-10 minutos)
   - Estratégias para refeições mistas
   - Adaptações para diferentes perfis (diabéticos, pré-diabéticos, obesidade)
   - Importância da mastigação adequada
   - Monitoramento da glicemia pós-prandial quando indicado

Retorne JSON com as 3 chaves.""",

        "outros": f"""Você é um nutricionista clínico. Gere conteúdo sobre aspectos complementares da alimentação não cobertos por categorias específicas.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a importância de avaliar aspectos diversos da alimentação que podem impactar a saúde metabólica, incluindo:
   - Padrões alimentares únicos ou culturais
   - Restrições ou preferências não mencionadas em outras categorias
   - Comportamentos alimentares atípicos
   - Uso de produtos específicos (suplementos, alimentos funcionais)

2. **patient_explanation** (100-150 palavras): Linguagem clara sobre:
   - Por que detalhes específicos da alimentação são importantes
   - Como informações adicionais ajudam na personalização do plano
   - Exemplos de informações relevantes a compartilhar

3. **conduct** (120-180 palavras): Orientações incluindo:
   - Avaliação individualizada de padrões alimentares únicos
   - Investigação de restrições, alergias ou intolerâncias não mencionadas
   - Análise de uso de suplementos e produtos específicos
   - Identificação de comportamentos que requerem atenção
   - Referência para especialistas quando necessário

Retorne JSON com as 3 chaves.""",

        "perfil_metabolico_mae": f"""Você é um médico especialista em medicina preventiva e genética metabólica. Gere conteúdo sobre o perfil metabólico materno durante a concepção/gestação.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a relevância clínica do perfil metabólico materno, incluindo:
   - Programação fetal e epigenética (DOHaD - Developmental Origins of Health and Disease)
   - Impacto de diabetes gestacional, obesidade, hipertensão e dislipidemia materna
   - Risco aumentado de síndrome metabólica na prole
   - Mecanismos de transmissão intergeracional (não apenas genética)
   - Importância da avaliação preconcepcional

2. **patient_explanation** (100-150 palavras): Linguagem acessível sobre:
   - Como a saúde metabólica da mãe durante a gestação influencia a saúde futura do filho
   - Conceito de "memória metabólica" desde o útero
   - Por que essas informações são valiosas para o acompanhamento atual
   - Não é culpa, mas sim oportunidade de prevenção

3. **conduct** (120-180 palavras): Orientações práticas:
   - Investigação detalhada do histórico metabólico materno durante a gestação
   - Rastreamento mais rigoroso para diabetes, dislipidemia, hipertensão
   - Estratégias preventivas personalizadas (nutrição, exercício)
   - Monitoramento laboratorial mais frequente quando indicado
   - Educação sobre risco aumentado e medidas de mitigação
   - Acompanhamento interdisciplinar quando necessário

Retorne JSON com as 3 chaves.""",

        "perfil_metabolico_pai": f"""Você é um médico especialista em medicina preventiva e genética metabólica. Gere conteúdo sobre o perfil metabólico paterno durante a concepção.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a relevância do perfil metabólico paterno, incluindo:
   - Contribuição epigenética paterna (modificações no esperma)
   - Evidências de transmissão de risco metabólico via linha paterna
   - Impacto de obesidade, diabetes, dislipidemia e hipertensão paterna
   - Mecanismos moleculares (metilação de DNA, RNA não codificantes)
   - Janela crítica pré-concepcional paterna

2. **patient_explanation** (100-150 palavras): Linguagem clara sobre:
   - Como a saúde do pai também influencia a predisposição metabólica dos filhos
   - Conceito de "herança além dos genes" (epigenética)
   - Importância dessas informações para estratificação de risco
   - Oportunidade de prevenção personalizada

3. **conduct** (120-180 palavras): Orientações práticas:
   - Levantamento do histórico metabólico paterno na época da concepção
   - Avaliação de risco cumulativo (materno + paterno)
   - Rastreamento preventivo para síndrome metabólica
   - Estratégias de modificação de estilo de vida específicas
   - Monitoramento laboratorial quando história familiar positiva
   - Aconselhamento sobre redução de risco

Retorne JSON com as 3 chaves.""",

        "piores_periodos": f"""Você é um nutricionista clínico especializado em comportamento alimentar. Gere conteúdo sobre períodos de pior alimentação ao longo da vida.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a importância clínica de identificar períodos de má alimentação, incluindo:
   - Janelas críticas de vulnerabilidade metabólica
   - Impacto cumulativo de períodos prolongados de má alimentação
   - Relação com desenvolvimento de doenças crônicas
   - Padrões de comportamento alimentar recorrentes
   - Fatores desencadeantes (estresse, transições de vida)

2. **patient_explanation** (100-150 palavras): Linguagem acessível sobre:
   - Por que identificar fases difíceis ajuda no tratamento atual
   - Compreensão de padrões e gatilhos
   - Oportunidade de aprender com experiências passadas
   - Não é sobre julgamento, mas sobre autoconhecimento

3. **conduct** (120-180 palavras): Orientações práticas:
   - Mapeamento de períodos críticos e fatores associados
   - Identificação de gatilhos (emocionais, sociais, ambientais)
   - Estratégias de prevenção de recaídas
   - Desenvolvimento de plano de contingência para momentos de risco
   - Apoio psicológico quando padrões recorrentes
   - Monitoramento mais próximo durante situações similares

Retorne JSON com as 3 chaves.""",

        "preferencias": f"""Você é um nutricionista clínico. Gere conteúdo sobre preferências alimentares do paciente.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a importância de conhecer preferências alimentares, incluindo:
   - Adesão ao plano nutricional (fator crítico de sucesso)
   - Identificação de padrões alimentares predominantes
   - Detecção de possíveis desequilíbrios nutricionais
   - Relação entre preferências e perfil metabólico
   - Oportunidades de substituições inteligentes

2. **patient_explanation** (100-150 palavras): Linguagem clara sobre:
   - Por que suas preferências são importantes e respeitadas
   - Como adaptar recomendações aos seus gostos
   - Possibilidade de expandir repertório alimentar gradualmente
   - Equilíbrio entre prazer e saúde

3. **conduct** (120-180 palavras): Orientações práticas:
   - Mapeamento detalhado de alimentos preferidos e evitados
   - Planejamento nutricional personalizado baseado em preferências
   - Estratégias de substituição para alimentos menos saudáveis preferidos
   - Educação nutricional para escolhas mais equilibradas
   - Técnicas para introdução gradual de novos alimentos
   - Respeito às restrições culturais e religiosas

Retorne JSON com as 3 chaves.""",

        "proteina_leite": f"""Você é um nutricionista clínico especializado em alergias e intolerâncias alimentares. Gere conteúdo sobre proteína do leite.

Item: {item_name}

Forneça:
1. **clinical_relevance** (150-200 palavras): Explique a importância clínica da avaliação de consumo/reação à proteína do leite, incluindo:
   - Diferença entre alergia à proteína do leite (APLV) e intolerância à lactose
   - Prevalência e manifestações clínicas (digestivas, cutâneas, respiratórias)
   - Impacto nutricional da exclusão de laticínios (cálcio, vitamina D, proteína)
   - Relação com permeabilidade intestinal e inflamação
   - Diagnóstico diferencial (IgE-mediada vs não-IgE)

2. **patient_explanation** (100-150 palavras): Linguagem acessível sobre:
   - Diferença entre alergia (reação imunológica) e intolerância (deficiência enzimática)
   - Sintomas comuns a observar
   - Importância de investigação adequada antes de exclusão
   - Fontes alternativas de cálcio e proteína

3. **conduct** (120-180 palavras): Orientações práticas:
   - Investigação detalhada de sintomas relacionados a laticínios
   - Testes diagnósticos quando indicado (IgE específico, teste de provocação oral)
   - Dieta de exclusão supervisionada quando suspeita confirmada
   - Planejamento nutricional para garantir aporte adequado de cálcio e proteína
   - Suplementação quando necessário (cálcio, vitamina D)
   - Orientação sobre leitura de rótulos e fontes ocultas
   - Reavaliação periódica (possibilidade de tolerância ao longo do tempo)

Retorne JSON com as 3 chaves."""
    }

    prompt = prompts.get(topic, prompts["outros"])

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=2000,
            temperature=0.7,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        response_text = message.content[0].text

        # Extrair JSON da resposta
        if "```json" in response_text:
            json_str = response_text.split("```json")[1].split("```")[0].strip()
        elif "```" in response_text:
            json_str = response_text.split("```")[1].split("```")[0].strip()
        else:
            json_str = response_text.strip()

        return json.loads(json_str)

    except Exception as e:
        print(f"Erro ao gerar conteúdo: {e}")
        return None

def main():
    results = {}
    total_items = sum(len(items) for items in ITEMS.values())
    current_item = 0

    print(f"Iniciando enriquecimento de {total_items} items de ALIMENTAÇÃO - Parte 3\n")

    for topic, item_ids in ITEMS.items():
        print(f"\n{'='*80}")
        print(f"Tópico: {topic.upper()} ({len(item_ids)} items)")
        print(f"{'='*80}\n")

        topic_name = {
            "ordem_alimentos": "Ordem dos alimentos na refeição",
            "outros": "Outros aspectos da alimentação",
            "perfil_metabolico_mae": "Perfil metabólico da mãe na época",
            "perfil_metabolico_pai": "Perfil metabólico do pai na época",
            "piores_periodos": "Piores períodos de alimentação",
            "preferencias": "Preferências alimentares",
            "proteina_leite": "Proteína do leite"
        }.get(topic, topic)

        # Gera conteúdo único para o tópico
        print(f"Gerando conteúdo clínico para {topic}...")
        content = generate_content(topic, topic_name)

        if content:
            # Aplica o mesmo conteúdo para todos os items do grupo
            for item_id in item_ids:
                current_item += 1
                results[item_id] = content
                print(f"[{current_item}/{total_items}] Item {item_id}: OK")
        else:
            print(f"ERRO ao gerar conteúdo para {topic}")
            for item_id in item_ids:
                current_item += 1
                print(f"[{current_item}/{total_items}] Item {item_id}: FALHOU")

    # Salvar resultados
    output_file = "/home/user/plenya/batch_dietary_part3_results.json"
    with open(output_file, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"\n{'='*80}")
    print(f"Resultados salvos em: {output_file}")
    print(f"Total de items processados: {len(results)}")
    print(f"{'='*80}\n")

    # Gerar SQL
    sql_statements = []
    sql_statements.append("-- Batch de Enriquecimento: ALIMENTAÇÃO - Parte 3")
    sql_statements.append("-- Total de items: 20")
    sql_statements.append("-- Gerado automaticamente\n")

    for item_id, content in results.items():
        clinical = content.get("clinical_relevance", "").replace("'", "''")
        explanation = content.get("patient_explanation", "").replace("'", "''")
        conduct = content.get("conduct", "").replace("'", "''")

        sql = f"""UPDATE score_items SET
    clinical_relevance = '{clinical}',
    patient_explanation = '{explanation}',
    conduct = '{conduct}',
    last_review = CURRENT_TIMESTAMP
WHERE id = '{item_id}';
"""
        sql_statements.append(sql)

    sql_file = "/home/user/plenya/batch_dietary_part3.sql"
    with open(sql_file, "w", encoding="utf-8") as f:
        f.write("\n".join(sql_statements))

    print(f"SQL gerado em: {sql_file}\n")

if __name__ == "__main__":
    main()
