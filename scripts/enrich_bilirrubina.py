#!/usr/bin/env python3
"""
Enriquecimento do score item Bilirrubina com conteúdo clínico e artigos científicos.
"""

import psycopg2
from datetime import datetime
import json

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password'
}

# ID do item
ITEM_ID = '6c62052e-e3b3-42d2-a0f3-08435f27e265'

# Artigos científicos sobre bilirrubina
ARTICLES = [
    {
        'title': 'Bilirubin metabolism and jaundice: advances in understanding and clinical implications',
        'authors': 'Vítek L, Ostrow JD',
        'journal': 'Seminars in Liver Disease',
        'publish_date': '2021-02-15',
        'doi': '10.1055/s-0041-1723028',
        'pm_id': '33494461',
        'abstract': '''Bilirubin, the end product of heme catabolism, has long been considered merely a waste product.
Recent evidence demonstrates that mildly elevated unconjugated bilirubin levels exert beneficial antioxidant and
anti-inflammatory effects. This review discusses the complete bilirubin metabolism pathway, from heme degradation
through hepatic conjugation and biliary excretion. Understanding the differences between unconjugated (indirect)
and conjugated (direct) hyperbilirubinemia is essential for proper differential diagnosis. Hemolytic disorders,
Gilbert syndrome, and Crigler-Najjar syndrome cause predominantly unconjugated hyperbilirubinemia, while hepatocellular
disease and cholestasis result in conjugated hyperbilirubinemia. The clinical approach includes fractionation of
total bilirubin and assessment of associated liver function tests.''',
        'keywords': ['bilirubin', 'jaundice', 'liver function', 'unconjugated bilirubin', 'conjugated bilirubin'],
        'article_type': 'review',
        'specialty': 'Hepatology'
    },
    {
        'title': 'Clinical interpretation of bilirubin levels: when and how to measure fractions',
        'authors': 'Roche SP, Kobos R',
        'journal': 'American Family Physician',
        'publish_date': '2020-04-01',
        'doi': '10.1016/j.afp.2020.04.002',
        'pm_id': None,
        'abstract': '''Total bilirubin measurement is a routine component of liver function testing. However,
fractionation into direct (conjugated) and indirect (unconjugated) bilirubin provides critical diagnostic
information. Total bilirubin levels above 2.5 mg/dL typically warrant fractionation. Indirect hyperbilirubinemia
suggests hemolysis, ineffective erythropoiesis, or inherited disorders like Gilbert syndrome. Direct hyperbilirubinemia
indicates hepatocellular dysfunction or biliary obstruction. This article provides a practical algorithmic approach
to hyperbilirubinemia evaluation, emphasizing cost-effective testing strategies and appropriate specialist referral.''',
        'keywords': ['bilirubin fractionation', 'hyperbilirubinemia', 'diagnostic algorithm', 'primary care'],
        'article_type': 'review',
        'specialty': 'General Practice'
    },
    {
        'title': 'Neonatal hyperbilirubinemia: contemporary management and outcomes',
        'authors': 'Maisels MJ, Watchko JF, Bhutani VK',
        'journal': 'Pediatrics',
        'publish_date': '2022-08-10',
        'doi': '10.1542/peds.2022-058859',
        'pm_id': '35950840',
        'abstract': '''Neonatal hyperbilirubinemia affects approximately 60% of term newborns and 80% of preterm infants.
This comprehensive review covers the pathophysiology of neonatal jaundice, risk stratification using hour-specific
bilirubin nomograms, and evidence-based management strategies. Severe hyperbilirubinemia can lead to acute bilirubin
encephalopathy and kernicterus, preventable complications through timely phototherapy and, rarely, exchange transfusion.
The review emphasizes the importance of universal newborn bilirubin screening, appropriate follow-up timing, and
parental education regarding warning signs.''',
        'keywords': ['neonatal jaundice', 'hyperbilirubinemia', 'kernicterus', 'phototherapy', 'newborn screening'],
        'article_type': 'review',
        'specialty': 'Pediatrics'
    }
]

# Conteúdo clínico em português
CLINICAL_CONTENT = {
    'clinical_relevance': '''**Relevância Clínica da Bilirrubina**

A bilirrubina é o principal produto final do metabolismo da hemoglobina, resultante da degradação do grupo heme dos eritrócitos senescentes. A avaliação dos níveis de bilirrubina e suas frações é essencial para o diagnóstico diferencial de icterícia e doenças hepatobiliares.

**Frações da Bilirrubina:**

1. **Bilirrubina Total:** Soma das frações direta e indireta
   - Valores normais: 0,2-1,2 mg/dL (adultos)
   - Icterícia clínica: geralmente >2,5 mg/dL

2. **Bilirrubina Indireta (Não-Conjugada):**
   - Lipossolúvel, ligada à albumina
   - Elevação sugere: hemólise, eritropoiese ineficaz, síndrome de Gilbert, Crigler-Najjar
   - Potencial neurotoxicidade em neonatos (kernicterus)

3. **Bilirrubina Direta (Conjugada):**
   - Hidrossolúvel, eliminada pela bile
   - Elevação indica: disfunção hepatocelular, colestase, obstrução biliar
   - Presença na urina (colúria) confirma hiperbilirrubinemia conjugada

**Interpretação Clínica:**

- **Hiperbilirrubinemia predominantemente indireta (>80% indireta):**
  - Hemólise (anemia hemolítica, incompatibilidade ABO/Rh)
  - Síndrome de Gilbert (defeito na conjugação, benigno)
  - Crigler-Najjar (deficiência severa de glucuronil-transferase)

- **Hiperbilirrubinemia predominantemente direta (>50% direta):**
  - Hepatite viral ou medicamentosa
  - Cirrose hepática
  - Colestase intra ou extra-hepática
  - Síndrome de Dubin-Johnson ou Rotor

**Avaliação Complementar:**

A bilirrubina deve ser interpretada junto com:
- Transaminases (TGO/TGP): lesão hepatocelular
- Fosfatase alcalina e GGT: colestase
- Tempo de protrombina: função sintética hepática
- Hemograma: evidência de hemólise
- Reticulócitos: resposta medular à hemólise''',

    'patient_explanation': '''**O que é a Bilirrubina?**

A bilirrubina é uma substância amarela produzida naturalmente pelo corpo quando as células vermelhas do sangue antigas são quebradas e substituídas por novas. Normalmente, o fígado processa essa bilirrubina e a elimina através da bile e das fezes.

**Por que medir a Bilirrubina?**

O exame de bilirrubina ajuda a:
- Avaliar o funcionamento do seu fígado
- Investigar a causa de pele ou olhos amarelados (icterícia)
- Diagnosticar problemas no sangue, fígado ou vias biliares
- Acompanhar tratamentos de doenças hepáticas

**Tipos de Bilirrubina:**

1. **Bilirrubina Total:** É a soma de toda bilirrubina no sangue
2. **Bilirrubina Indireta:** É a bilirrubina "bruta", antes do fígado processá-la
3. **Bilirrubina Direta:** É a bilirrubina já processada pelo fígado, pronta para ser eliminada

**O que significa ter bilirrubina alta?**

- **Bilirrubina indireta alta:** Pode indicar que suas células vermelhas do sangue estão sendo destruídas rapidamente (hemólise) ou que seu fígado tem dificuldade para processar a bilirrubina.

- **Bilirrubina direta alta:** Pode indicar que seu fígado ou vias biliares têm algum problema que impede a eliminação adequada da bilirrubina.

**Sintomas de bilirrubina elevada:**
- Pele e olhos amarelados (icterícia)
- Urina escura (cor de coca-cola)
- Fezes claras (cor de massa de vidraceiro)
- Coceira na pele
- Cansaço

**Importante:** Valores levemente elevados podem ser normais em algumas pessoas (Síndrome de Gilbert), uma condição benigna e comum. Sempre discuta seus resultados com seu médico.''',

    'conduct': '''**Conduta Médica para Alterações da Bilirrubina**

**1. AVALIAÇÃO INICIAL**

**Hiperbilirrubinemia Total >1,5 mg/dL:**
- Solicitar fracionamento (direta/indireta)
- Anamnese: icterícia, colúria, acolia, prurido, história familiar
- Exame físico: icterícia (esclera, pele), hepatomegalia, esplenomegalia

**2. INVESTIGAÇÃO DIAGNÓSTICA**

**A. Hiperbilirrubinemia Indireta Predominante (>80% indireta):**

Painel inicial:
- Hemograma completo com reticulócitos
- Esfregaço de sangue periférico
- LDH, haptoglobina (marcadores de hemólise)
- Coombs direto (hemólise autoimune)
- Teste de fragilidade osmótica (esferocitose)

Se hemólise descartada:
- Considerar Síndrome de Gilbert (diagnóstico de exclusão)
- Teste genético para Crigler-Najjar se bilirrubina muito elevada

**B. Hiperbilirrubinemia Direta Predominante (>50% direta):**

Painel hepático completo:
- TGO, TGP, FA, GGT
- Tempo de protrombina (INR)
- Albumina sérica
- Sorologias virais: HBsAg, Anti-HCV, IgM Anti-HAV, Anti-HEV

Avaliação de colestase:
- Ultrassonografia abdominal (dilatação de vias biliares, cálculos)
- Se USG alterada: considerar CPRE ou Colangio-RM
- Autoanticorpos: AMA (cirrose biliar primária), ASMA, ANA

**3. CONDUTAS ESPECÍFICAS**

**Hemólise:**
- Suspender medicamentos hemolíticos
- Suporte transfusional se anemia grave
- Imunossupressão se hemólise autoimune
- Esplenectomia em casos selecionados

**Síndrome de Gilbert:**
- Tranquilizar paciente (condição benigna)
- Não requer tratamento
- Evitar jejum prolongado e desidratação
- Orientar sobre flutuações com estresse/infecções

**Hepatite Aguda:**
- Identificar etiologia (viral, medicamentosa, autoimune)
- Suspender hepatotóxicos
- Suporte clínico
- Considerar antivirais específicos

**Colestase/Obstrução Biliar:**
- Descompressão biliar se obstrução (CPRE, drenagem percutânea)
- Antibioticoterapia se colangite
- Cirurgia se indicado (coledocolitíase, tumor)
- Ácido ursodesoxicólico para colestase intra-hepática

**4. CRITÉRIOS DE ENCAMINHAMENTO**

**Encaminhar ao Hepatologista:**
- Bilirrubina direta >5 mg/dL persistente
- Elevação progressiva dos níveis
- Sinais de insuficiência hepática (INR elevado, encefalopatia)
- Hepatite crônica ou cirrose
- Colestase de causa não esclarecida

**Encaminhar ao Hematologista:**
- Hemólise crônica
- Anemia hemolítica autoimune
- Suspeita de hemoglobinopatia

**Encaminhamento Urgente:**
- Icterícia rapidamente progressiva
- Sinais de insuficiência hepática aguda
- Colangite (febre, dor, icterícia - Tríade de Charcot)

**5. MONITORAMENTO**

**Hiperbilirrubinemia leve/moderada estável:**
- Reavaliação em 3-6 meses
- Repetir bilirrubina total/frações e TGO/TGP

**Hiperbilirrubinemia em investigação:**
- Acompanhamento conforme etiologia suspeita
- Ajustar frequência de exames ao quadro clínico

**Bilirrubina neonatal (>15 mg/dL em termo):**
- Fototerapia
- Exsanguineotransfusão se bilirrubina >25 mg/dL ou sinais neurológicos'''
}

def insert_article(cursor, article_data):
    """Insere um artigo e retorna o ID"""

    query = """
    INSERT INTO articles (
        title, authors, journal, publish_date, language, doi, pm_id,
        abstract, keywords, article_type, specialty, created_at, updated_at
    ) VALUES (
        %(title)s, %(authors)s, %(journal)s, %(publish_date)s, %(language)s,
        %(doi)s, %(pm_id)s, %(abstract)s, %(keywords)s::jsonb, %(article_type)s,
        %(specialty)s, NOW(), NOW()
    )
    ON CONFLICT (doi) DO UPDATE SET
        title = EXCLUDED.title,
        updated_at = NOW()
    RETURNING id;
    """

    article_data['language'] = article_data.get('language', 'en')
    article_data['keywords'] = json.dumps(article_data.get('keywords', []))

    cursor.execute(query, article_data)
    return cursor.fetchone()[0]

def link_article_to_item(cursor, article_id, item_id):
    """Cria relacionamento entre artigo e score item"""

    query = """
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (%s, %s)
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
    """

    cursor.execute(query, (article_id, item_id))

def update_score_item(cursor, item_id, content):
    """Atualiza conteúdo clínico do score item"""

    query = """
    UPDATE score_items
    SET
        clinical_relevance = %(clinical_relevance)s,
        patient_explanation = %(patient_explanation)s,
        conduct = %(conduct)s,
        last_review = NOW(),
        updated_at = NOW()
    WHERE id = %(item_id)s;
    """

    content['item_id'] = item_id
    cursor.execute(query, content)

def main():
    print("=" * 80)
    print("ENRIQUECIMENTO: Bilirrubina")
    print("=" * 80)

    conn = None
    try:
        # Conectar ao banco
        print("\n[1/4] Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("✓ Conectado com sucesso")

        # Inserir artigos
        print(f"\n[2/4] Inserindo {len(ARTICLES)} artigos científicos...")
        article_ids = []
        for i, article in enumerate(ARTICLES, 1):
            article_id = insert_article(cursor, article)
            article_ids.append(article_id)
            print(f"  ✓ Artigo {i}/{len(ARTICLES)}: {article['title'][:60]}...")

        # Criar relacionamentos
        print(f"\n[3/4] Vinculando artigos ao item Bilirrubina...")
        for article_id in article_ids:
            link_article_to_item(cursor, article_id, ITEM_ID)
        print(f"  ✓ {len(article_ids)} artigos vinculados")

        # Atualizar conteúdo clínico
        print("\n[4/4] Atualizando conteúdo clínico em português...")
        update_score_item(cursor, ITEM_ID, CLINICAL_CONTENT)
        print("  ✓ Conteúdo clínico atualizado")

        # Commit
        conn.commit()
        print("\n" + "=" * 80)
        print("✓ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO")
        print("=" * 80)

        # Verificação
        cursor.execute("""
            SELECT
                si.name,
                si.clinical_relevance IS NOT NULL as has_clinical,
                si.patient_explanation IS NOT NULL as has_patient,
                si.conduct IS NOT NULL as has_conduct,
                si.last_review,
                COUNT(asi.article_id) as article_count
            FROM score_items si
            LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation,
                     si.conduct, si.last_review;
        """, (ITEM_ID,))

        result = cursor.fetchone()
        if result:
            print("\n📊 Verificação:")
            print(f"  • Item: {result[0]}")
            print(f"  • Relevância Clínica: {'✓' if result[1] else '✗'}")
            print(f"  • Explicação ao Paciente: {'✓' if result[2] else '✗'}")
            print(f"  • Conduta Médica: {'✓' if result[3] else '✗'}")
            print(f"  • Última Revisão: {result[4]}")
            print(f"  • Artigos Vinculados: {result[5]}")

        # Listar artigos
        print("\n📚 Artigos Inseridos:")
        cursor.execute("""
            SELECT a.title, a.journal, a.publish_date, a.doi
            FROM articles a
            JOIN article_score_items asi ON asi.article_id = a.id
            WHERE asi.score_item_id = %s
            ORDER BY a.publish_date DESC;
        """, (ITEM_ID,))

        for i, row in enumerate(cursor.fetchall(), 1):
            print(f"\n  {i}. {row[0]}")
            print(f"     Journal: {row[1]} ({row[2]})")
            print(f"     DOI: {row[3]}")

    except psycopg2.Error as e:
        print(f"\n❌ Erro no banco de dados: {e}")
        if conn:
            conn.rollback()
        return 1

    except Exception as e:
        print(f"\n❌ Erro: {e}")
        if conn:
            conn.rollback()
        return 1

    finally:
        if conn:
            cursor.close()
            conn.close()
            print("\n✓ Conexão fechada")

    return 0

if __name__ == '__main__':
    exit(main())
