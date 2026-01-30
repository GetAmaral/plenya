#!/usr/bin/env python3
"""
Script para enriquecer o item HCM (MCH) com conteúdo científico
Item ID: 81050f48-2b89-4da9-a480-f55af29bdb8b
Grupo: Exames > Laboratoriais
"""

import psycopg2
from datetime import datetime
import sys

# Configuração do banco
DB_CONFIG = {
    'dbname': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password',
    'host': 'localhost',  # Usando localhost pois o script roda no host
    'port': 5432
}

ITEM_ID = '81050f48-2b89-4da9-a480-f55af29bdb8b'

# Conteúdo clínico enriquecido em PT-BR
CLINICAL_RELEVANCE = """HCM (Hemoglobina Corpuscular Média) quantifica a quantidade média de hemoglobina contida em cada hemácia individual, sendo expressa em picogramas (pg) por célula. Valores de referência: 27-33 pg/célula.

**Fisiopatologia:**
O HCM reflete diretamente a síntese de hemoglobina intraeritrocitária. Valores baixos indicam hipocromia (hemácias com conteúdo reduzido de hemoglobina), enquanto valores elevados sugerem macrocitose com conteúdo hemoglobínico aumentado.

**Correlações Clínicas:**

*HCM Diminuído (< 27 pg):*
- Anemia ferropriva: deficiência de ferro compromete síntese de hemoglobina
- Talassemias: defeitos genéticos nas cadeias de globina
- Anemia de doença crônica: distúrbios no metabolismo do ferro
- Anemia sideroblástica: defeito na incorporação de ferro ao heme

*HCM Aumentado (> 33 pg):*
- Anemia megaloblástica: deficiência de B12/folato gera hemácias grandes
- Hipotireoidismo: redução da eritropoiese com macrocitose
- Hepatopatias crônicas: alterações na membrana eritrocitária
- Alcoolismo crônico: toxicidade direta na medula óssea

**Relação com VCM:**
O HCM varia paralelamente ao VCM (Volume Corpuscular Médio). Na anemia ferropriva, ambos estão diminuídos (microcitose + hipocromia). Nas anemias megaloblásticas, ambos estão elevados (macrocitose com alto conteúdo de hemoglobina).

**Limitações:**
O HCM é calculado (não medido diretamente) a partir de Hb, hematócrito e contagem de hemácias. Erros em qualquer desses parâmetros afetam o resultado. O exame do esfregaço periférico é fundamental para confirmação morfológica.

**Referências Científicas:**
- Wintrobe MM. Anemia: classification and treatment on the basis of differences in the average volume and hemoglobin content of red corpuscles. Archives of Internal Medicine. 1934.
- NCBI Bookshelf. Red Cell Indices - Clinical Methods. 2024.
- Ratio of hemoglobin to mean corpuscular volume: diagnostic performance in discriminating iron deficiency anemia from thalassemia trait. PMC 2023."""

PATIENT_EXPLANATION = """O **HCM (Hemoglobina Corpuscular Média)** mede a quantidade média de hemoglobina dentro de cada glóbulo vermelho do seu sangue. A hemoglobina é a proteína que transporta oxigênio para todo o corpo.

**Por que este exame é importante?**
O HCM ajuda a identificar o tipo de anemia que você pode ter. Valores baixos geralmente indicam que suas células estão com pouca hemoglobina (hipocromia), o que pode ser causado por falta de ferro. Valores altos podem indicar células grandes com bastante hemoglobina, comum em deficiências de vitaminas B12 ou ácido fólico.

**Valores normais:**
- Adultos: 27 a 33 picogramas por célula (pg)

**O que significa resultado BAIXO (< 27 pg)?**
Pode indicar:
- Anemia por falta de ferro (mais comum)
- Talassemia (condição genética)
- Anemia de doenças crônicas

Sintomas comuns: cansaço, palidez, falta de ar, fraqueza.

**O que significa resultado ALTO (> 33 pg)?**
Pode indicar:
- Deficiência de vitamina B12 ou ácido fólico
- Problemas na tireoide
- Uso excessivo de álcool
- Doenças do fígado

**Importante:**
O HCM sempre deve ser interpretado junto com outros exames do hemograma (VCM, CHCM, hemoglobina) para determinar a causa exata da alteração. Nunca tome suplementos sem orientação médica."""

CONDUCT = """**Interpretação Inicial:**

1. **HCM < 27 pg (Hipocromia):**
   - Solicitar: ferritina, saturação de transferrina, ferro sérico
   - Se ferritina < 30 ng/mL: confirma anemia ferropriva
   - Se ferritina normal + microcitose: investigar talassemia (eletroforese de hemoglobina)
   - Considerar anemia de doença crônica (PCR, VHS)

2. **HCM 27-33 pg (Normal):**
   - Avaliar no contexto do hemograma completo
   - Se anemia presente: investigar outras causas (hemólise, insuficiência renal)

3. **HCM > 33 pg (Macrocitose):**
   - Solicitar: vitamina B12, ácido fólico, TSH
   - Investigar uso de álcool e medicamentos (metotrexato, AZT)
   - Avaliar função hepática (TGO, TGP, GGT)
   - Se B12 < 200 pg/mL ou folato < 2 ng/mL: suplementação

**Correlação Obrigatória:**
- Sempre interpretar HCM junto com VCM, CHCM e RDW
- Solicitar esfregaço periférico se discordância entre índices
- Avaliar reticulócitos para diferenciar causas centrais vs. periféricas

**Tratamento Específico:**

*Anemia Ferropriva:*
- Sulfato ferroso 300mg VO 3x/dia (em jejum)
- Investigar causa da perda: endoscopia, colonoscopia em adultos
- Reavaliação em 4-6 semanas

*Anemia Megaloblástica:*
- Vitamina B12: 1000mcg IM diariamente por 1 semana, depois semanal por 1 mês
- Ácido fólico: 5mg VO 1x/dia
- Corrigir B12 ANTES do folato (prevenir degeneração combinada subaguda)

**Situações de Urgência:**
- HCM < 20 pg + Hb < 7 g/dL: considerar transfusão
- HCM muito elevado + leucopenia/plaquetopenia: urgência hematológica (avaliar SMD)

**Follow-up:**
- Reavaliar hemograma em 4-6 semanas após início do tratamento
- HCM normaliza gradualmente (2-3 meses)
- Manter ferritina > 50 ng/mL para prevenir recorrência"""

def enrich_hcm_item():
    """Enriquece o item HCM no banco de dados"""

    conn = None
    cursor = None

    try:
        # Conectar ao banco
        print(f"Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()

        # Verificar item atual
        cursor.execute("""
            SELECT id, name, clinical_relevance, patient_explanation, conduct
            FROM score_items
            WHERE id = %s
        """, (ITEM_ID,))

        item = cursor.fetchone()
        if not item:
            print(f"❌ Item {ITEM_ID} não encontrado!")
            return False

        print(f"✅ Item encontrado: {item[1]}")
        print(f"   - clinical_relevance: {'✅' if item[2] else '❌ vazio'}")
        print(f"   - patient_explanation: {'✅' if item[3] else '❌ vazio'}")
        print(f"   - conduct: {'✅' if item[4] else '❌ vazio'}")

        # Atualizar item
        print(f"\n📝 Atualizando campos clínicos...")
        cursor.execute("""
            UPDATE score_items
            SET
                clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                last_review = %s,
                updated_at = CURRENT_TIMESTAMP
            WHERE id = %s
        """, (
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            datetime.now(),
            ITEM_ID
        ))

        rows_updated = cursor.rowcount
        conn.commit()

        if rows_updated > 0:
            print(f"✅ Item atualizado com sucesso!")

            # Buscar artigos relacionados existentes
            print(f"\n🔍 Buscando artigos relacionados no banco...")
            cursor.execute("""
                SELECT id, title, authors, journal, doi, pm_id
                FROM articles
                WHERE
                    LOWER(title) LIKE '%mean corpuscular%'
                    OR LOWER(title) LIKE '%hemoglobin%'
                    OR LOWER(title) LIKE '%anemia%'
                    OR LOWER(abstract) LIKE '%mch%'
                LIMIT 10
            """)

            articles = cursor.fetchall()
            print(f"📊 Encontrados {len(articles)} artigos relacionados")

            if articles:
                print("\n📚 Artigos encontrados:")
                for article in articles:
                    print(f"   - {article[1][:80]}...")

                    # Criar relação se ainda não existir
                    cursor.execute("""
                        INSERT INTO score_item_articles (score_item_id, article_id)
                        VALUES (%s, %s)
                        ON CONFLICT (score_item_id, article_id) DO NOTHING
                    """, (ITEM_ID, article[0]))

                relations_created = cursor.rowcount
                conn.commit()
                print(f"✅ {relations_created} novas relações criadas")

            # Verificar resultado final
            print(f"\n🔍 Verificando resultado final...")
            cursor.execute("""
                SELECT
                    si.name,
                    LENGTH(si.clinical_relevance) as len_clinical,
                    LENGTH(si.patient_explanation) as len_patient,
                    LENGTH(si.conduct) as len_conduct,
                    si.last_review,
                    COUNT(sia.article_id) as num_articles
                FROM score_items si
                LEFT JOIN score_item_articles sia ON si.id = sia.score_item_id
                WHERE si.id = %s
                GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review
            """, (ITEM_ID,))

            result = cursor.fetchone()
            if result:
                print(f"✅ Resultado Final:")
                print(f"   - Nome: {result[0]}")
                print(f"   - clinical_relevance: {result[1]} caracteres")
                print(f"   - patient_explanation: {result[2]} caracteres")
                print(f"   - conduct: {result[3]} caracteres")
                print(f"   - last_review: {result[4]}")
                print(f"   - artigos relacionados: {result[5]}")

            return True
        else:
            print(f"❌ Nenhuma linha foi atualizada")
            return False

    except Exception as e:
        print(f"❌ Erro: {e}")
        if conn:
            conn.rollback()
        return False
    finally:
        if cursor:
            cursor.close()
        if conn:
            conn.close()

if __name__ == "__main__":
    print("=" * 80)
    print("🧬 ENRIQUECIMENTO DO ITEM HCM (Mean Corpuscular Hemoglobin)")
    print("=" * 80)
    print(f"Item ID: {ITEM_ID}")
    print(f"Grupo: Exames > Laboratoriais")
    print("=" * 80)

    success = enrich_hcm_item()

    print("\n" + "=" * 80)
    if success:
        print("✅ MISSÃO CUMPRIDA!")
        print("=" * 80)
        sys.exit(0)
    else:
        print("❌ FALHA NA EXECUÇÃO")
        print("=" * 80)
        sys.exit(1)
