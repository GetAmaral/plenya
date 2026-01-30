#!/usr/bin/env python3
"""
Batch Final 3 - Exames C (35 items)
Enriquecimento com padrão MFI completo
"""

import json
import anthropic
import os
from pathlib import Path

# Configurar cliente Anthropic
client = anthropic.Anthropic(api_key=os.environ.get("ANTHROPIC_API_KEY"))

def generate_enrichment(item: dict) -> dict:
    """Gera conteúdo MFI enriquecido via Claude Sonnet 4.5"""

    prompt = f"""Você é especialista em Medicina Laboratorial Funcional Integrativa (MFI).

**ITEM A ENRIQUECER:**
- ID: {item['id']}
- Nome: {item['name']}
- Subgrupo: {item['subgroup']}

**INSTRUÇÕES CRÍTICAS:**

1. **VALORES DE REFERÊNCIA FUNCIONAIS (optimal_range_min/max):**
   - Use faixas MFI otimizadas (não apenas valores laboratoriais padrão)
   - Se houver diferença entre homens/mulheres, especifique no nome do item
   - Para exames qualitativos, use NULL nos campos numéricos

2. **FAIXA ÓTIMA (optimal_explanation):**
   - Justificativa científica da faixa funcional
   - 2-3 linhas, máximo 200 caracteres
   - Exemplo: "Valores entre 40-60% reduzem mortalidade cardiovascular em 35% (Framingham, 2023)"

3. **INTERPRETAÇÃO CLÍNICA FUNCIONAL (clinical_interpretation):**
   - 3-4 parágrafos detalhados (400-600 palavras)
   - Subseções: Fundamento Fisiológico, Interpretação Valores Baixos, Interpretação Valores Altos, Contexto Clínico
   - Correlações entre múltiplos marcadores
   - Diferenças etárias/sexuais quando aplicável

4. **CONDUTAS MÉDICAS (medical_conduct):**
   - Estrutura markdown com bullet points
   - Subdivisões: Valores Baixos, Valores Altos, Abordagem Funcional
   - SEMPRE especificar doses exatas (mg, UI, g/dia)
   - Exemplo: "Vitamina D3: 5.000-10.000 UI/dia se <30 ng/mL"
   - Incluir suplementação, estilo de vida, dieta, acompanhamento

5. **ARTIGOS CIENTÍFICOS (related_articles):**
   - 3-5 artigos PubMed recentes (2020-2025)
   - Foco em medicina funcional/preventiva
   - Formato: Array de strings com citação completa

**EXIGÊNCIAS DE QUALIDADE:**
- Zero placeholders genéricos
- Doses e valores numéricos específicos
- Fundamentação científica em cada afirmação
- Linguagem técnica mas clara

**FORMATO DE SAÍDA (JSON válido):**
```json
{{
  "item_id": "{item['id']}",
  "optimal_range_min": 45.0,
  "optimal_range_max": 60.0,
  "unit": "unidade",
  "optimal_explanation": "Texto justificativa...",
  "clinical_interpretation": "## Fundamento Fisiológico\\n\\n...\\n\\n## Interpretação Valores Baixos\\n\\n...\\n\\n## Interpretação Valores Altos\\n\\n...\\n\\n## Contexto Clínico\\n\\n...",
  "medical_conduct": "## Valores Baixos\\n\\n- Conduta 1 (dose específica)\\n- Conduta 2\\n\\n## Valores Altos\\n\\n- Conduta 3\\n\\n## Abordagem Funcional\\n\\n- Monitoramento...",
  "related_articles": [
    "Autor A et al. Título. Journal. 2024;Vol(Issue):pages. PMID: 12345678",
    "Autor B et al. Título. Journal. 2023;Vol(Issue):pages. PMID: 87654321"
  ]
}}
```

IMPORTANTE: Retorne APENAS o JSON, sem markdown code blocks."""

    try:
        message = client.messages.create(
            model="claude-sonnet-4-5-20250929",
            max_tokens=4000,
            temperature=0.7,
            messages=[{"role": "user", "content": prompt}]
        )

        response_text = message.content[0].text.strip()

        # Remover markdown code blocks se existirem
        if response_text.startswith("```"):
            lines = response_text.split("\n")
            response_text = "\n".join(lines[1:-1])

        enrichment = json.loads(response_text)

        print(f"✓ {item['name']}")
        return enrichment

    except Exception as e:
        print(f"✗ ERRO em {item['name']}: {e}")
        return None

def generate_sql(enrichments: list) -> str:
    """Gera SQL único com todos os 35 items"""

    sql_lines = [
        "-- BATCH FINAL 3 - EXAMES C (35 items)",
        "-- Gerado automaticamente com padrão MFI",
        "-- Data: 2026-01-28",
        "",
        "BEGIN;",
        ""
    ]

    for enr in enrichments:
        if not enr:
            continue

        # Escapar strings SQL
        def escape_sql(text):
            if text is None:
                return "NULL"
            return "'" + str(text).replace("'", "''") + "'"

        # Converter related_articles para array PostgreSQL
        articles_array = "NULL"
        if enr.get("related_articles"):
            articles_escaped = [f'"{art.replace(chr(34), chr(34)+chr(34))}"' for art in enr["related_articles"]]
            articles_array = f"ARRAY[{', '.join([escape_sql(art) for art in enr['related_articles']])}]"

        sql_lines.append(f"""
UPDATE score_items
SET
  optimal_range_min = {enr.get('optimal_range_min', 'NULL')},
  optimal_range_max = {enr.get('optimal_range_max', 'NULL')},
  unit = {escape_sql(enr.get('unit'))},
  optimal_explanation = {escape_sql(enr.get('optimal_explanation'))},
  clinical_interpretation = {escape_sql(enr.get('clinical_interpretation'))},
  medical_conduct = {escape_sql(enr.get('medical_conduct'))},
  related_articles = {articles_array},
  updated_at = NOW()
WHERE id = '{enr['item_id']}'::uuid;
""")

    sql_lines.extend([
        "",
        "COMMIT;",
        "",
        f"-- Total de items atualizados: {len([e for e in enrichments if e])}",
        ""
    ])

    return "\n".join(sql_lines)

def main():
    print("=" * 80)
    print("BATCH FINAL 3 - EXAMES C (35 items)")
    print("=" * 80)
    print()

    # Carregar dados
    data_file = Path(__file__).parent / "enrichment_data" / "batch_final_3_exames_C.json"
    with open(data_file) as f:
        data = json.load(f)

    items = data["items"]
    print(f"Total de items: {len(items)}")
    print()

    # Processar cada item
    enrichments = []
    for i, item in enumerate(items, 1):
        print(f"[{i}/{len(items)}] ", end="")
        enrichment = generate_enrichment(item)
        enrichments.append(enrichment)

    print()
    print("=" * 80)
    print("GERANDO SQL...")
    print("=" * 80)

    # Gerar SQL
    sql_content = generate_sql(enrichments)

    # Salvar SQL
    output_file = Path(__file__).parent / "batch_final_3_exames_C.sql"
    with open(output_file, "w", encoding="utf-8") as f:
        f.write(sql_content)

    print(f"✓ SQL gerado: {output_file}")
    print()

    # Salvar JSON de backup
    json_output = Path(__file__).parent / "batch_final_3_exames_C_results.json"
    with open(json_output, "w", encoding="utf-8") as f:
        json.dump({
            "batch": "final_3_exames_C",
            "total": len(items),
            "processed": len([e for e in enrichments if e]),
            "failed": len([e for e in enrichments if not e]),
            "enrichments": enrichments
        }, f, indent=2, ensure_ascii=False)

    print(f"✓ Backup JSON: {json_output}")
    print()

    # Estatísticas
    successful = len([e for e in enrichments if e])
    print("=" * 80)
    print("ESTATÍSTICAS")
    print("=" * 80)
    print(f"Total items: {len(items)}")
    print(f"Processados com sucesso: {successful}")
    print(f"Falhas: {len(items) - successful}")
    print()

    print("PRÓXIMO PASSO:")
    print(f"  docker compose exec -T db psql -U plenya_user -d plenya_db < {output_file}")
    print()

if __name__ == "__main__":
    main()
