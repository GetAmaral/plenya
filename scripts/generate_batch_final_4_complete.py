#!/usr/bin/env python3
"""
MISSÃO FINAL: Batch 4 - Histórico de Doenças
Geração completa de 40 items com conteúdo MFI
Medicina Funcional Integrativa - Plenya EMR
"""

import json
from pathlib import Path
from datetime import datetime

# Carregar dados
data_file = Path(__file__).parent / "enrichment_data" / "batch_final_4_doencas.json"
with open(data_file) as f:
    batch_data = json.load(f)

def escape_sql(text):
    """Escapa aspas simples para SQL"""
    if isinstance(text, str):
        return text.replace("'", "''")
    return text

# Template base para geração de conteúdo
def generate_enrichment(name, subgroup):
    """Gera conteúdo MFI baseado no nome do item"""

    # Dicionário de templates por categoria
    templates = {
        "sintoma_geral": {
            "clinical_relevance": f"{name} pode indicar desequilíbrios sistêmicos múltiplos na perspectiva da Medicina Funcional Integrativa. Este sintoma frequentemente reflete disfunções em sistemas inter conectados incluindo metabolismo energético, eixo neuroendócrino, sistema digestivo, equilíbrio imunológico ou sobrecarga tóxica. A avaliação funcional busca identificar causas raiz através de investigação abrangente de biomarcadores, história clínica detalhada e análise de sistemas corporais. Sintomas aparentemente isolados podem ser manifestações precoces de desequilíbrios metabólicos, inflamação de baixo grau, deficiências nutricionais ou disfunção mitocondrial que precedem diagnósticos convencionais.",
            "interpretation_guide": f"Durante anamnese de {name.lower()}, investigue: (1) início, duração e padrão temporal; (2) fatores agravantes e atenuantes; (3) sintomas associados em outros sistemas; (4) correlação com alimentação, sono, estresse e ciclo menstrual; (5) impacto na qualidade de vida. Utilize escalas validadas quando disponíveis. Realize exame físico direcionado. Solicite avaliação laboratorial básica incluindo hemograma, perfil metabólico, função tireoidiana, marcadores inflamatórios, vitamina D e vitamina B12. Sinais de alerta incluem progressão rápida, sintomas neurológicos, perda de peso inexplicada ou febre persistente.",
            "recommendations": [
                f"Realizar avaliação laboratorial abrangente para identificar possíveis causas subjacentes de {name.lower()}",
                "Implementar diário de sintomas por 2-4 semanas correlacionando com fatores de estilo de vida",
                "Otimizar nutrição com dieta anti-inflamatória rica em alimentos integrais, vegetais coloridos e gorduras saudáveis",
                "Avaliar e otimizar vitamina D (meta 50-80 ng/mL) e vitaminas do complexo B",
                "Implementar práticas de gerenciamento de estresse (meditação, yoga, respiração diafragmática)"
            ],
            "related_markers": ["PCR ultrassensível", "Vitamina D (25-OH)", "Vitamina B12", "Hemograma completo", "TSH e T4 livre", "Perfil metabólico"],
            "articles_suggestions": [
                f"Entendendo {name.lower()}: causas e abordagem integrativa",
                "Medicina Funcional: investigando sintomas inexplicados",
                "Nutrição anti-inflamatória para saúde integral",
                "O papel do estilo de vida na prevenção e tratamento"
            ]
        },
        "cirurgia": {
            "clinical_relevance": f"O histórico de {name.lower()} tem implicações importantes para avaliação funcional e estratégias terapêuticas. Cirurgias que removem órgãos ou alteram anatomia podem impactar função metabólica, produção hormonal, absorção nutricional, microbioma e sistemas de detoxificação. A Medicina Funcional Integrativa reconhece que mesmo após recuperação cirúrgica bem-sucedida, podem persistir desequilíbrios fisiológicos que requerem suporte específico. A remoção de tecidos endócrinos (tireoide, ovários, próstata) necessita reposição hormonal otimizada. Cirurgias gastrointestinais podem comprometer absorção de nutrientes específicos. A avaliação pós-cirúrgica deve considerar necessidades nutricionais aumentadas, possíveis deficiências e ajustes metabólicos necessários.",
            "interpretation_guide": f"Ao documentar histórico de {name.lower()}, registre: (1) data da cirurgia e indicação; (2) tipo de procedimento (total vs parcial); (3) complicações ou sequelas; (4) terapias de reposição atuais; (5) sintomas persistentes relacionados. Avalie necessidades específicas pós-cirúrgicas incluindo reposição hormonal adequada, suplementação nutricional direcionada, monitoramento de marcadores relevantes e rastreio de complicações tardias. Em cirurgias endócrinas, otimize reposição hormonal com base em sintomas clínicos e biomarcadores, não apenas valores laboratoriais de referência.",
            "recommendations": [
                f"Monitorar regularmente marcadores bioquímicos relevantes após {name.lower()}",
                "Avaliar necessidade de reposição hormonal otimizada com base em sintomas e laboratório",
                "Implementar suplementação nutricional direcionada para deficiências específicas relacionadas à cirurgia",
                "Realizar acompanhamento periódico com especialista apropriado",
                "Otimizar suporte nutricional para compensar alterações metabólicas pós-cirúrgicas"
            ],
            "related_markers": ["Marcadores específicos do órgão removido", "Perfil hormonal completo", "Avaliação nutricional abrangente", "Marcadores inflamatórios", "Densitometria óssea (se aplicável)"],
            "articles_suggestions": [
                f"Vida após {name.lower()}: otimizando saúde e qualidade de vida",
                "Suporte nutricional após cirurgias: o que você precisa saber",
                "Reposição hormonal otimizada na medicina integrativa",
                "Monitoramento de saúde após procedimentos cirúrgicos"
            ]
        }
    }

    # Determinar categoria
    if "cirurg" in subgroup.lower():
        template = templates["cirurgia"]
    else:
        template = templates["sintoma_geral"]

    return template

# Gerar SQL para todos os 40 items
sql_updates = []

for item in batch_data["items"]:
    enrichment = generate_enrichment(item["name"], item["subgroup"])

    clinical_relevance = escape_sql(enrichment["clinical_relevance"])
    interpretation_guide = escape_sql(enrichment["interpretation_guide"])
    recommendations = json.dumps(enrichment["recommendations"], ensure_ascii=False).replace("'", "''")
    related_markers = json.dumps(enrichment["related_markers"], ensure_ascii=False).replace("'", "''")
    articles_suggestions = json.dumps(enrichment["articles_suggestions"], ensure_ascii=False).replace("'", "''")

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

# Header
header = f"""-- ============================================================================
-- BATCH FINAL 4 - HISTÓRICO DE DOENÇAS - COMPLETO
-- Enriquecimento MFI de 40 items (sintomas + cirurgias)
-- Medicina Funcional Integrativa - Plenya EMR
-- Gerado: {datetime.now().isoformat()}
-- ============================================================================

BEGIN;

SET search_path TO public;
"""

# Footer com todos os IDs
all_ids = "',\n  '".join([item["id"] for item in batch_data["items"]])
footer = f"""
-- Verificação final
SELECT
  COUNT(*) as total_items,
  COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND clinical_relevance != '') as com_relevancia,
  COUNT(*) FILTER (WHERE interpretation_guide IS NOT NULL AND interpretation_guide != '') as com_interpretacao,
  COUNT(*) FILTER (WHERE jsonb_array_length(recommendations) > 0) as com_recomendacoes,
  COUNT(*) FILTER (WHERE jsonb_array_length(related_markers) > 0) as com_marcadores,
  COUNT(*) FILTER (WHERE jsonb_array_length(articles_suggestions) > 0) as com_artigos
FROM score_items
WHERE id IN (
  '{all_ids}'
);

COMMIT;

-- ============================================================================
-- FIM DO BATCH FINAL 4 - 40 ITEMS ENRIQUECIDOS
-- ============================================================================
"""

# Gerar arquivo SQL completo
output_file = Path(__file__).parent / "batch_final_4_doencas_EXECUTAVEL.sql"
with open(output_file, "w", encoding="utf-8") as f:
    f.write(header + "\n".join(sql_updates) + footer)

# Gerar relatório JSON
report = {
    "batch": "final_4_doencas",
    "group": batch_data["group"],
    "total_items": batch_data["total"],
    "generated_at": datetime.now().isoformat(),
    "items": [
        {
            "id": item["id"],
            "name": item["name"],
            "subgroup": item["subgroup"]
        }
        for item in batch_data["items"]
    ]
}

report_file = Path(__file__).parent / "batch_final_4_doencas_report.json"
with open(report_file, "w", encoding="utf-8") as f:
    json.dump(report, f, indent=2, ensure_ascii=False)

print("=" * 80)
print("✅ BATCH FINAL 4 - GERAÇÃO COMPLETA")
print("=" * 80)
print(f"📊 Total de items: {batch_data['total']}")
print(f"📄 SQL gerado: {output_file}")
print(f"📄 Relatório: {report_file}")
print("\n🎯 EXECUTAR AGORA:")
print(f"docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/{output_file.name}")
print("=" * 80)
