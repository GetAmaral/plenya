#!/usr/bin/env python3
"""
Batch Final 1 - Exames Laboratoriais Parte A (45 items)
Processador MFI com geração de SQL único para execução via Docker
"""

import anthropic
import json
import os
import sys
from datetime import datetime

# Configuração
API_KEY = os.getenv("ANTHROPIC_API_KEY")
if not API_KEY:
    print("❌ ANTHROPIC_API_KEY não configurada")
    sys.exit(1)

client = anthropic.Anthropic(api_key=API_KEY)

# Carregar dados
with open("/home/user/plenya/scripts/enrichment_data/batch_final_1_exames_A.json", "r", encoding="utf-8") as f:
    data = json.load(f)

print(f"🎯 BATCH FINAL 1 - EXAMES LABORATORIAIS PARTE A")
print(f"📊 Total de items: {data['total']}")
print(f"🏥 Grupo: {data['group']}")
print(f"⏰ Início: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")

# System prompt MFI
SYSTEM_PROMPT = """Você é um especialista em Medicina Laboratorial Funcional Integrativa (MFI).

**CONTEXTO:** Sistema de prontuário eletrônico com score items para avaliação de saúde.

**PADRÃO MFI OBRIGATÓRIO:**

1. **Valores Ótimos Funcionais:** Não apenas "normal", mas faixas ótimas para longevidade
2. **Interpretação Funcional:** Significado clínico dos valores alterados
3. **Condutas Específicas:** Intervenções com doses exatas (suplementos, dieta, exercício)
4. **Artigos Científicos:** Mínimo 3 referências PubMed/journals relevantes

**FORMATO DE RESPOSTA (JSON):**

```json
{
  "interpretation": "Interpretação funcional detalhada com faixas ótimas. Ex: 'Valores ótimos: 50-100 ng/mL (não apenas >30). Valores <50 indicam deficiência subclínica...'",
  "low_level_description": "Consequências de valores baixos com mecanismos fisiopatológicos",
  "medium_level_description": "Estado intermediário e riscos associados",
  "high_level_description": "Consequências de valores elevados e ações necessárias",
  "low_level_recommendation": "Condutas específicas: suplementação (doses), dieta (alimentos), exercício (tipo/intensidade)",
  "medium_level_recommendation": "Ajustes para otimização funcional",
  "high_level_recommendation": "Intervenções para normalização com follow-up",
  "articles": [
    {
      "title": "Título completo do artigo",
      "url": "https://pubmed.ncbi.nlm.nih.gov/PMID ou DOI"
    }
  ]
}
```

**EXIGÊNCIAS:**
- Interpretation: mínimo 200 caracteres
- Descriptions: mínimo 150 caracteres cada
- Recommendations: condutas específicas com doses/protocolos
- Artigos: 3-5 referências válidas (PubMed preferencial)

**ADAPTAR PARA TIPO DE EXAME:**
- Exames de Imagem: interpretar scores/medidas, relação com desfechos clínicos
- Exames Laboratoriais: valores ótimos, não apenas referência laboratorial
- Exames com faixas etárias/gênero: especificar contexto na interpretação
"""

# Processar todos os items
results = []
sql_statements = []

for idx, item in enumerate(data["items"], 1):
    print(f"\n{'='*80}")
    print(f"[{idx}/{data['total']}] Processando: {item['name']}")
    print(f"ID: {item['id']}")
    print(f"Subgrupo: {item['subgroup']}")

    try:
        # Chamar Claude
        message = client.messages.create(
            model="claude-opus-4-5-20251101",
            max_tokens=4000,
            temperature=1,
            system=SYSTEM_PROMPT,
            messages=[{
                "role": "user",
                "content": f"""Enriqueça este score item com padrão MFI:

**Nome:** {item['name']}
**Subgrupo:** {item['subgroup']}
**Contexto:** Exame {'de imagem' if item['subgroup'] == 'Imagem' else 'laboratorial'}

Retorne APENAS o JSON conforme especificado no system prompt."""
            }]
        )

        # Parse resposta
        content = message.content[0].text.strip()

        # Limpar markdown se presente
        if content.startswith("```json"):
            content = content.replace("```json", "").replace("```", "").strip()
        elif content.startswith("```"):
            content = content.replace("```", "").strip()

        enrichment = json.loads(content)

        # Validação
        required_fields = [
            "interpretation", "low_level_description", "medium_level_description",
            "high_level_description", "low_level_recommendation",
            "medium_level_recommendation", "high_level_recommendation", "articles"
        ]

        missing = [f for f in required_fields if f not in enrichment or not enrichment[f]]
        if missing:
            print(f"⚠️  Campos faltando: {missing}")
            continue

        # Validar tamanho mínimo
        if len(enrichment["interpretation"]) < 200:
            print(f"⚠️  Interpretation muito curto: {len(enrichment['interpretation'])} chars")
            continue

        # Validar artigos
        if len(enrichment["articles"]) < 3:
            print(f"⚠️  Apenas {len(enrichment['articles'])} artigos (mínimo 3)")
            continue

        # Gerar SQL
        articles_json = json.dumps(enrichment["articles"], ensure_ascii=False).replace("'", "''")

        sql = f"""
-- {item['name']} ({item['id']})
UPDATE score_items SET
  interpretation = '{enrichment['interpretation'].replace("'", "''")}',
  low_level_description = '{enrichment['low_level_description'].replace("'", "''")}',
  medium_level_description = '{enrichment['medium_level_description'].replace("'", "''")}',
  high_level_description = '{enrichment['high_level_description'].replace("'", "''")}',
  low_level_recommendation = '{enrichment['low_level_recommendation'].replace("'", "''")}',
  medium_level_recommendation = '{enrichment['medium_level_recommendation'].replace("'", "''")}',
  high_level_recommendation = '{enrichment['high_level_recommendation'].replace("'", "''")}',
  articles = '{articles_json}'::jsonb,
  last_review = NOW()
WHERE id = '{item['id']}';
"""

        sql_statements.append(sql)

        # Resultado
        result = {
            "id": item["id"],
            "name": item["name"],
            "subgroup": item["subgroup"],
            "enrichment": enrichment,
            "status": "success",
            "chars": {
                "interpretation": len(enrichment["interpretation"]),
                "low_desc": len(enrichment["low_level_description"]),
                "medium_desc": len(enrichment["medium_level_description"]),
                "high_desc": len(enrichment["high_level_description"]),
                "low_rec": len(enrichment["low_level_recommendation"]),
                "medium_rec": len(enrichment["medium_level_recommendation"]),
                "high_rec": len(enrichment["high_level_recommendation"])
            },
            "articles_count": len(enrichment["articles"])
        }

        results.append(result)

        print(f"✅ Enriquecido com sucesso")
        print(f"   Interpretation: {result['chars']['interpretation']} chars")
        print(f"   Artigos: {result['articles_count']}")

    except Exception as e:
        print(f"❌ Erro: {str(e)}")
        results.append({
            "id": item["id"],
            "name": item["name"],
            "status": "error",
            "error": str(e)
        })
        continue

# Gerar SQL final
print(f"\n{'='*80}")
print(f"📝 GERANDO SQL ÚNICO PARA EXECUÇÃO")

sql_file = "/home/user/plenya/batch_final_1_exames_A.sql"
with open(sql_file, "w", encoding="utf-8") as f:
    f.write("-- BATCH FINAL 1 - EXAMES LABORATORIAIS PARTE A\n")
    f.write(f"-- Gerado em: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
    f.write(f"-- Total de items: {len(sql_statements)}/{data['total']}\n\n")
    f.write("BEGIN;\n\n")
    f.write("\n".join(sql_statements))
    f.write("\n\nCOMMIT;\n")

print(f"✅ SQL gerado: {sql_file}")

# Salvar resultados JSON
results_file = "/home/user/plenya/batch_final_1_exames_A_results.json"
with open(results_file, "w", encoding="utf-8") as f:
    json.dump({
        "batch": "final_1_exames_A",
        "generated_at": datetime.now().isoformat(),
        "total_items": data["total"],
        "processed": len(results),
        "successful": len([r for r in results if r.get("status") == "success"]),
        "failed": len([r for r in results if r.get("status") == "error"]),
        "results": results
    }, f, indent=2, ensure_ascii=False)

print(f"✅ Resultados salvos: {results_file}")

# Estatísticas finais
successful = [r for r in results if r.get("status") == "success"]
failed = [r for r in results if r.get("status") == "error"]

print(f"\n{'='*80}")
print(f"📊 ESTATÍSTICAS FINAIS")
print(f"✅ Sucesso: {len(successful)}/{data['total']}")
print(f"❌ Falhas: {len(failed)}/{data['total']}")

if successful:
    avg_chars = {
        "interpretation": sum(r["chars"]["interpretation"] for r in successful) / len(successful),
        "descriptions": sum(r["chars"]["low_desc"] + r["chars"]["medium_desc"] + r["chars"]["high_desc"] for r in successful) / len(successful) / 3,
        "recommendations": sum(r["chars"]["low_rec"] + r["chars"]["medium_rec"] + r["chars"]["high_rec"] for r in successful) / len(successful) / 3
    }
    avg_articles = sum(r["articles_count"] for r in successful) / len(successful)

    print(f"\n📈 Médias:")
    print(f"   Interpretation: {avg_chars['interpretation']:.0f} chars")
    print(f"   Descriptions: {avg_chars['descriptions']:.0f} chars")
    print(f"   Recommendations: {avg_chars['recommendations']:.0f} chars")
    print(f"   Artigos: {avg_articles:.1f}")

if failed:
    print(f"\n⚠️  Items com falha:")
    for r in failed:
        print(f"   - {r['name']}: {r.get('error', 'Unknown')}")

print(f"\n⏰ Fim: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
print(f"\n🎯 PRÓXIMO PASSO: Executar SQL via Docker")
print(f"   docker compose exec db psql -U plenya_user -d plenya_db -f /tmp/batch_final_1_exames_A.sql")
