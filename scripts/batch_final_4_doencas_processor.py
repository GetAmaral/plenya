#!/usr/bin/env python3
"""
Batch Final 4 - Histórico de Doenças
Enriquecimento de 40 items com padrão MFI
Medicina Funcional Integrativa - Plenya EMR
"""

import json
import anthropic
import os
from pathlib import Path

# Configuração
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY")
if not ANTHROPIC_API_KEY:
    raise ValueError("ANTHROPIC_API_KEY não encontrada")

client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)

# Carregar dados
data_file = Path(__file__).parent / "enrichment_data" / "batch_final_4_doencas.json"
with open(data_file) as f:
    batch_data = json.load(f)

print(f"📋 Batch Final 4 - {batch_data['group']}")
print(f"📊 Total de items: {batch_data['total']}")
print("=" * 80)

# Template de enriquecimento MFI para histórico de doenças/sintomas
ENRICHMENT_PROMPT = """Você é um especialista em Medicina Funcional Integrativa.

**ITEM:** {name}
**SUBGRUPO:** {subgroup}
**CONTEXTO:** Histórico de doenças/sintomas ou cirurgias realizadas

Gere conteúdo clínico estruturado seguindo EXATAMENTE este padrão MFI:

**1. clinical_relevance (200-300 palavras):**
- Descreva a relevância clínica do sintoma/doença/cirurgia
- Fisiopatologia sob perspectiva funcional (causas raiz sistêmicas)
- Conexões com outros sistemas corporais
- Impacto na saúde global e qualidade de vida
- Linguagem acessível mas tecnicamente precisa

**2. interpretation_guide (150-250 palavras):**
- Como investigar este item na anamnese
- Sinais de alerta e red flags
- Perguntas essenciais para aprofundamento
- Correlações com outros sintomas/sistemas
- Quando referenciar para especialista

**3. recommendations (3-5 itens práticos):**
- Protocolos terapêuticos baseados em evidências
- Modificações de estilo de vida específicas
- Suplementação funcional quando aplicável
- Terapias integrativas validadas
- Monitoramento e follow-up

**4. related_markers (lista JSON):**
["Biomarcador 1", "Biomarcador 2", "Biomarcador 3"]
- Exames laboratoriais relevantes
- Biomarcadores funcionais específicos
- Marcadores de risco associados

**5. articles_suggestions (3-5 tópicos):**
- Temas para artigos educacionais
- Foco em empoderamento do paciente
- Linguagem não técnica

**FORMATO DE SAÍDA (JSON puro, sem markdown):**
{{
  "clinical_relevance": "texto aqui...",
  "interpretation_guide": "texto aqui...",
  "recommendations": [
    "Recomendação 1 completa e específica",
    "Recomendação 2 com detalhes práticos",
    "Recomendação 3 baseada em evidências"
  ],
  "related_markers": ["Marcador1", "Marcador2", "Marcador3"],
  "articles_suggestions": [
    "Tópico 1 para artigo educacional",
    "Tópico 2 focado em prevenção",
    "Tópico 3 sobre manejo integrativo"
  ]
}}

**IMPORTANTE:**
- JSON válido, sem comentários
- Textos em português brasileiro
- Abordagem funcional integrativa
- Evidências científicas quando possível
- Linguagem acessível mas precisa"""

def enrich_item(item: dict) -> dict:
    """Enriquece um item usando Claude API"""

    print(f"\n🔄 Enriquecendo: {item['name']}")
    print(f"   Subgrupo: {item['subgroup']}")

    prompt = ENRICHMENT_PROMPT.format(
        name=item['name'],
        subgroup=item['subgroup']
    )

    try:
        message = client.messages.create(
            model="claude-opus-4-5-20251101",
            max_tokens=4000,
            temperature=0.7,
            messages=[{
                "role": "user",
                "content": prompt
            }]
        )

        content = message.content[0].text.strip()

        # Remove markdown se presente
        if content.startswith("```json"):
            content = content.split("```json")[1].split("```")[0].strip()
        elif content.startswith("```"):
            content = content.split("```")[1].split("```")[0].strip()

        enrichment = json.loads(content)

        print(f"   ✅ Enriquecido com sucesso")
        return {
            "id": item["id"],
            "name": item["name"],
            "subgroup": item["subgroup"],
            "enrichment": enrichment
        }

    except json.JSONDecodeError as e:
        print(f"   ❌ Erro JSON: {e}")
        print(f"   Conteúdo: {content[:200]}")
        return None
    except Exception as e:
        print(f"   ❌ Erro: {e}")
        return None

def generate_sql(enriched_items: list) -> str:
    """Gera SQL consolidado para todos os items"""

    sql_updates = []

    for item in enriched_items:
        if not item or "enrichment" not in item:
            continue

        e = item["enrichment"]

        # Escapar aspas simples para SQL
        def escape_sql(text):
            if isinstance(text, str):
                return text.replace("'", "''")
            return text

        clinical_relevance = escape_sql(e.get("clinical_relevance", ""))
        interpretation_guide = escape_sql(e.get("interpretation_guide", ""))
        recommendations = json.dumps(e.get("recommendations", []), ensure_ascii=False).replace("'", "''")
        related_markers = json.dumps(e.get("related_markers", []), ensure_ascii=False).replace("'", "''")
        articles_suggestions = json.dumps(e.get("articles_suggestions", []), ensure_ascii=False).replace("'", "''")

        sql = f"""
-- {item['name']} ({item['subgroup']})
UPDATE score_items
SET
  clinical_relevance = '{clinical_relevance}',
  interpretation_guide = '{interpretation_guide}',
  recommendations = '{recommendations}'::jsonb,
  related_markers = '{related_markers}'::jsonb,
  articles_suggestions = '{articles_suggestions}'::jsonb,
  updated_at = NOW()
WHERE id = '{item['id']}';
"""
        sql_updates.append(sql)

    # Header do SQL
    header = """-- ============================================================================
-- BATCH FINAL 4 - HISTÓRICO DE DOENÇAS
-- Enriquecimento MFI de 40 items (sintomas + cirurgias)
-- Medicina Funcional Integrativa - Plenya EMR
-- Gerado automaticamente em: {date}
-- ============================================================================

BEGIN;

SET search_path TO public;
"""

    # Footer do SQL
    footer = """
-- Verificação final
SELECT
  COUNT(*) as total_enriquecidos,
  COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND clinical_relevance != '') as com_relevancia,
  COUNT(*) FILTER (WHERE interpretation_guide IS NOT NULL AND interpretation_guide != '') as com_interpretacao,
  COUNT(*) FILTER (WHERE jsonb_array_length(recommendations) > 0) as com_recomendacoes
FROM score_items
WHERE id IN (
  '1176540d-cefa-4d2c-b5e2-4a992060de4d',
  '360d1e6a-84c5-4763-a743-0fce76fe2686',
  'adcb06a3-4791-4ce2-ae05-c01f8fef9ead',
  '65a28dbf-9466-4a23-9410-9d6034c3e141',
  'b7370158-c34e-449f-8305-537a8fb85d98',
  'e24dae19-4cb0-4d83-a6db-9571aabf9bde',
  'a7fdd0e5-9aa7-4604-b5ae-731c05c90af0',
  '40837ae0-790a-4a1a-a7ea-fb9d0cc21878',
  'b4cda178-b5a9-4193-b496-afe7ce51ed91',
  'deb8e2ee-27a7-4ac0-a743-d6cb400f4b27',
  'fdb0c742-1db1-48ba-95bc-91e454256d84',
  'e375191b-d8be-4044-9d59-b71fa222cbaa',
  '44f501b3-6c1d-4a31-b8c9-0f54b201dbcb',
  '14881a00-579e-4297-bd6d-4c83d8081d2d',
  'a7698447-ddbf-4bee-92b4-81a1559562ad',
  '76578167-9f69-4981-b73d-46aa7c71ec6b',
  '23b90c6c-6f44-481a-8255-108158a64239',
  '4d1372ee-5bee-4ce8-845a-a4e176109ad2',
  'fe938ed9-3cf9-4ddf-a1ae-2f67c03e54a4',
  '5ddc4af9-dd19-4ea1-8117-3d3e30377dab',
  'e010b2a7-6e9e-4b42-9d37-487e91411a18',
  '933b7816-fe10-4e35-94d7-0d232cc258ce',
  '2b01e2f0-76c7-4616-8f7b-33b1c9d11279',
  '33ffce34-e1fb-4e31-b393-f2783549d874',
  'f54dbfaa-599a-40ee-907a-28431ce4858a',
  '30a7764c-4c11-4078-b942-860ee56136a4',
  '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd',
  '4511ae8d-2ad3-40e0-a47d-2cef065d39e9',
  'cbb42389-9a9d-48fe-a64a-61215107d5e3',
  'e65c56dc-5c07-4270-8a8b-017b293ca147',
  'ad36ad5d-e587-43a8-b1ae-3e27305e25d8',
  '8115cf13-eab0-4db2-9844-d04c37701d92',
  '19b1c83e-f2a7-47e9-9045-5c994cfcbf94',
  'fca251a2-5aa2-42f9-b93c-0f0a852d9d51',
  '925c2474-689e-486a-a6b1-3e1b4ca8cc6e',
  'e5ea8f41-05e0-48e0-b6a8-2323b3896449',
  'f691cda1-bd2d-449e-9e66-045319ceaaa3',
  'ec0bbb80-873a-4946-bf90-bc759eddb080',
  '53dbbf23-25bf-4b91-9670-b17fab93c6e9',
  '8c2456de-1ba1-4a1f-ab7c-964f8c8114e6'
);

COMMIT;

-- ============================================================================
-- FIM DO BATCH FINAL 4
-- ============================================================================
"""

    from datetime import datetime
    full_sql = header.format(date=datetime.now().isoformat()) + "\n".join(sql_updates) + footer

    return full_sql

def main():
    """Executa o processamento completo"""

    enriched_items = []

    # Enriquecer todos os items
    for idx, item in enumerate(batch_data["items"], 1):
        print(f"\n[{idx}/{batch_data['total']}]")
        enriched = enrich_item(item)
        if enriched:
            enriched_items.append(enriched)

    # Estatísticas
    success_count = len(enriched_items)
    print("\n" + "=" * 80)
    print(f"✅ Items enriquecidos: {success_count}/{batch_data['total']}")
    print(f"❌ Falhas: {batch_data['total'] - success_count}")

    if enriched_items:
        # Gerar SQL
        sql_output = generate_sql(enriched_items)

        # Salvar arquivos
        output_dir = Path(__file__).parent

        # SQL file
        sql_file = output_dir / "batch_final_4_doencas.sql"
        with open(sql_file, "w", encoding="utf-8") as f:
            f.write(sql_output)
        print(f"\n📄 SQL gerado: {sql_file}")

        # JSON com dados enriquecidos
        json_file = output_dir / "batch_final_4_doencas_results.json"
        with open(json_file, "w", encoding="utf-8") as f:
            json.dump({
                "batch": "final_4_doencas",
                "group": batch_data["group"],
                "total_items": batch_data["total"],
                "enriched_count": success_count,
                "items": enriched_items
            }, f, indent=2, ensure_ascii=False)
        print(f"📄 JSON gerado: {json_file}")

        print("\n" + "=" * 80)
        print("🎯 PRÓXIMOS PASSOS:")
        print("1. Revisar o SQL gerado")
        print("2. Executar via Docker:")
        print(f"   docker compose exec -T db psql -U plenya_user -d plenya_db < {sql_file.name}")
        print("3. Verificar resultados no banco")
        print("=" * 80)
    else:
        print("\n❌ Nenhum item foi enriquecido com sucesso")

if __name__ == "__main__":
    main()
