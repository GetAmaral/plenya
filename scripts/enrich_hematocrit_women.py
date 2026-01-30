#!/usr/bin/env python3
"""
Enriquecimento: Hematócrito - Mulheres
ID: 32037968-f263-4e7d-ab85-ea83efd61c7b
Grupo: Exames > Laboratoriais
"""

import os
import sys
import anthropic
from datetime import datetime

# Configuração
ITEM_ID = "32037968-f263-4e7d-ab85-ea83efd61c7b"
ITEM_NAME = "Hematócrito - Mulheres"

def enrich_item():
    """Enriquece o item com conteúdo clínico baseado em evidências científicas"""

    api_key = os.environ.get("ANTHROPIC_API_KEY")
    if not api_key:
        print("❌ ANTHROPIC_API_KEY não encontrada")
        sys.exit(1)

    client = anthropic.Anthropic(api_key=api_key)

    prompt = f"""Você é um médico especialista em hematologia e patologia clínica, com expertise em fisiologia feminina.

**TAREFA:** Enriquecer o item "{ITEM_NAME}" com conteúdo clínico baseado em evidências científicas.

**CONTEXTO CLÍNICO:**
- Hematócrito é a porcentagem do volume sanguíneo ocupado por eritrócitos
- Valores normais em mulheres adultas: 36-48% (menores que em homens)
- Diferenças de gênero devido a: menstruação, hormônios, menor massa muscular
- Reflete capacidade de transporte de oxigênio
- Alterações indicam: anemia (comum em mulheres), desidratação, policitemia
- Impacto na viscosidade sanguínea e oxigenação tecidual

**INSTRUÇÕES:**

1. **BUSCAR ARTIGOS CIENTÍFICOS:**
   - Use web_search para encontrar 2-3 artigos científicos relevantes
   - Termos: "hematocrit women normal values", "anemia women menstruation", "iron deficiency hematocrit", "hematocrit reference ranges gender"
   - Priorize: PubMed, UpToDate, journals médicos (2020-2025)
   - Foque em: valores de referência femininos, diferenças de gênero, menstruação, ferro

2. **CLINICAL_RELEVANCE** (PT-BR, 150-200 palavras):
   - Valores de referência em mulheres (36-48%)
   - Por que valores são menores que em homens (menstruação, hormônios)
   - Interpretação: valores baixos (<36%), normais, elevados (>48%)
   - Causas de alteração específicas: menstruação, deficiência de ferro, gravidez, anemia ferropriva
   - Correlação com outros marcadores (hemoglobina, ferritina)
   - Impacto na oxigenação e sintomas (fadiga, palidez)
   - Cite fontes científicas encontradas

3. **PATIENT_EXPLANATION** (PT-BR, linguagem leiga, 100-120 palavras):
   - O que é hematócrito (percentual de glóbulos vermelhos)
   - Por que mulheres têm valores menores (menstruação)
   - O que valores baixos podem significar (anemia, cansaço)
   - Sintomas de alerta: fadiga, palidez, falta de ar
   - Quando procurar médico
   - Evite termos técnicos complexos

4. **CONDUCT** (PT-BR, bullet points, acionável):
   - Protocolo para hematócrito baixo (<36%):
     * Investigar ferritina, ferro sérico, capacidade de ligação
     * Avaliar perdas menstruais (menorragia)
     * Dosagem de vitamina B12 e ácido fólico
     * Considerar suplementação de ferro
   - Protocolo para hematócrito elevado (>48%):
     * Avaliar hidratação
     * Investigar policitemia
     * Verificar tabagismo, apneia do sono
   - Quando encaminhar ao hematologista
   - Orientações sobre coleta (evitar período menstrual)

5. **RETORNAR ARTIGOS ENCONTRADOS:**
   - Liste os 2-3 artigos científicos com:
     * Título completo
     * Autores
     * URL completo
     * Ano de publicação
     * Breve resumo da relevância (1-2 frases)

**FORMATO DE SAÍDA:**
```json
{{
  "clinical_relevance": "texto em PT-BR...",
  "patient_explanation": "texto em PT-BR...",
  "conduct": "texto em PT-BR...",
  "articles": [
    {{
      "title": "título completo",
      "authors": "autores",
      "url": "https://...",
      "year": 2024,
      "relevance": "por que este artigo é relevante"
    }}
  ]
}}
```

**IMPORTANTE:**
- Todo texto em português brasileiro
- Baseie-se em evidências científicas dos artigos encontrados
- Cite fontes quando mencionar dados específicos
- Linguagem clara e objetiva
- Foco em diferenças de gênero (menstruação, ferro, hormônios)
- Aplicabilidade clínica prática para população feminina
"""

    print(f"🔬 Enriquecendo: {ITEM_NAME}")
    print(f"📋 ID: {ITEM_ID}")
    print(f"🤖 Buscando artigos científicos e gerando conteúdo...\n")

    try:
        message = client.messages.create(
            model="claude-opus-4-5-20251101",
            max_tokens=16000,
            temperature=0.3,
            tools=[
                {
                    "type": "web_search_tool"
                }
            ],
            messages=[
                {
                    "role": "user",
                    "content": prompt
                }
            ]
        )

        # Extrair resposta
        response_text = ""
        for block in message.content:
            if hasattr(block, 'text'):
                response_text += block.text

        print("✅ Conteúdo gerado com sucesso!\n")
        print("=" * 80)
        print(response_text)
        print("=" * 80)

        # Extrair JSON do response
        import json
        import re

        json_match = re.search(r'```json\s*(\{.*?\})\s*```', response_text, re.DOTALL)
        if json_match:
            content = json.loads(json_match.group(1))

            # Preparar SQL
            clinical_relevance = content['clinical_relevance'].replace("'", "''")
            patient_explanation = content['patient_explanation'].replace("'", "''")
            conduct = content['conduct'].replace("'", "''")

            sql = f"""
-- Atualizar score_item com conteúdo clínico: Hematócrito - Mulheres
UPDATE score_items
SET
  clinical_relevance = '{clinical_relevance}',
  patient_explanation = '{patient_explanation}',
  conduct = '{conduct}',
  last_review = NOW()
WHERE id = '{ITEM_ID}';
"""

            # Salvar SQL
            sql_file = "/home/user/plenya/scripts/hematocrit_women_update.sql"
            with open(sql_file, 'w', encoding='utf-8') as f:
                f.write(sql)

            print(f"\n✅ SQL gerado: {sql_file}")

            # Salvar artigos para processamento posterior
            if 'articles' in content and content['articles']:
                articles_file = "/home/user/plenya/scripts/hematocrit_women_articles.json"
                with open(articles_file, 'w', encoding='utf-8') as f:
                    json.dump(content['articles'], f, indent=2, ensure_ascii=False)

                print(f"📚 Artigos encontrados: {len(content['articles'])}")
                print(f"📄 Salvo em: {articles_file}")

                for i, article in enumerate(content['articles'], 1):
                    print(f"\n{i}. {article['title']}")
                    print(f"   Autores: {article.get('authors', 'N/A')}")
                    print(f"   Ano: {article.get('year', 'N/A')}")
                    print(f"   URL: {article['url']}")

            return sql_file, content.get('articles', [])
        else:
            print("❌ Não foi possível extrair JSON da resposta")
            return None, []

    except Exception as e:
        print(f"❌ Erro: {e}")
        import traceback
        traceback.print_exc()
        return None, []

if __name__ == "__main__":
    sql_file, articles = enrich_item()

    if sql_file:
        print("\n" + "=" * 80)
        print("📊 PRÓXIMOS PASSOS:")
        print("=" * 80)
        print(f"1. Revisar SQL: {sql_file}")
        print(f"2. Aplicar no banco:")
        print(f"   docker compose exec -T db psql -U plenya_user -d plenya_db < {sql_file}")
        print(f"\n3. Upload de artigos (se houver):")
        if articles:
            print(f"   - {len(articles)} artigos encontrados")
            print(f"   - Executar upload_articles.py com hematocrit_women_articles.json")
        print("=" * 80)
    else:
        print("\n❌ Falha no enriquecimento")
        sys.exit(1)
