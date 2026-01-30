#!/usr/bin/env python3
"""
Enriquecimento do score item "Magnésio Sérico"
ID: 2813e3c1-a2d0-464d-ad05-35aec75a3d7d
Data: 2026-01-29
Baseado em artigos científicos 2024-2025
"""

import psycopg2
from datetime import datetime

# Database connection
conn = psycopg2.connect(
    host="localhost",
    port=5432,
    database="plenya_db",
    user="plenya_user",
    password="plenya_dev_password"
)
conn.autocommit = False
cur = conn.cursor()

try:
    print("Iniciando enriquecimento de Magnésio Sérico...")

    # 1. Update score_item
    clinical_relevance = """O magnésio sérico é um marcador essencial que reflete o status de magnésio no organismo, mineral envolvido em mais de 300 processos metabólicos fundamentais. Desempenha papel crítico na produção de energia celular, regulação do sistema nervoso, função cardiovascular e metabolismo ósseo. Valores normais variam entre 1,7-2,2 mg/dL, porém é importante destacar que níveis séricos normais não excluem estados de deficiência moderada a severa, uma vez que apenas 1% do magnésio corporal está presente no sangue. A hipomagnesemia é uma condição frequentemente subdiagnosticada, afetando aproximadamente 50% dos adultos que não consomem quantidades adequadas através da dieta. Estudos epidemiológicos, ensaios clínicos randomizados e meta-análises demonstram relação inversa consistente entre status de magnésio e doenças cardiovasculares, incluindo hipertensão, calcificação arterial coronariana, AVC, doença cardíaca isquêmica, fibrilação atrial, insuficiência cardíaca e mortalidade cardíaca. No sistema nervoso, o magnésio modula a neuroinflamação, previne excitotoxicidade por manter o equilíbrio de cálcio neuronal, inibe a ativação de NF-κB reduzindo citocinas pró-inflamatórias, e regula a função sináptica essencial para memória e aprendizado. Para cada aumento de 0,5 mg/dL no magnésio sérico, observa-se redução de 7% no risco de hipertensão. A deficiência está particularmente presente em indivíduos com diabetes mellitus (30% dos pacientes), idosos, pacientes com doenças crônicas, usuários de diuréticos, inibidores de bomba de prótons e antibióticos. O conceito de "deficiência latente crônica" descreve indivíduos com magnésio corporal total reduzido mas níveis séricos normais, contribuindo silenciosamente para desenvolvimento de doenças cardiovasculares e metabólicas."""

    patient_explanation = """O magnésio é um mineral essencial que seu corpo precisa para funcionar adequadamente, participando de centenas de processos importantes como produção de energia, funcionamento dos músculos e nervos, saúde do coração e fortalecimento dos ossos. Este exame mede a quantidade de magnésio no seu sangue. Valores normais ficam geralmente entre 1,7 e 2,2 mg/dL. Porém, é importante saber que mesmo com resultados normais neste exame, você pode ter deficiência de magnésio no corpo, pois apenas uma pequena parte do magnésio total está no sangue - a maior parte fica dentro das células e nos ossos. A falta de magnésio é mais comum do que se imagina e pode causar diversos sintomas como fraqueza muscular, cãibras, cansaço, alterações no batimento cardíaco, dificuldade para dormir, ansiedade e até mudanças no humor. Cerca de metade dos adultos não consome magnésio suficiente pela alimentação. A deficiência é mais frequente em pessoas com diabetes, idosos, quem usa certos medicamentos (como diuréticos e remédios para estômago), e quem tem problemas intestinais que dificultam a absorção. Estudos científicos mostram que ter magnésio adequado protege seu coração, ajuda a controlar a pressão arterial, melhora o funcionamento do cérebro e pode reduzir o risco de derrame e problemas cardíacos. Para cada aumento de 0,5 mg/dL no magnésio, o risco de pressão alta diminui 7%. Boas fontes de magnésio incluem vegetais verdes escuros, grãos integrais, leguminosas, oleaginosas (castanhas, amêndoas), sementes, abacate e cacau."""

    conduct = """AVALIAÇÃO INICIAL: Solicitar dosagem de magnésio sérico em pacientes com fatores de risco (diabetes mellitus, idosos, uso crônico de diuréticos/IBP/antibióticos, doenças gastrointestinais, hipertensão refratária, arritmias cardíacas inexplicadas, cãibras musculares frequentes, osteoporose, síndrome metabólica). Em casos de suspeita clínica forte mas magnésio sérico normal, considerar dosagem de magnésio eritrocitário ou magnésio urinário de 24h para melhor avaliação do status corporal total. HIPOMAGNESEMIA LEVE A MODERADA (1,2-1,7 mg/dL): Iniciar suplementação oral com magnésio quelato, citrato, glicinato ou cloreto 300-450 mg/dia (evitar óxido de magnésio pela baixa biodisponibilidade). Ajustar dose conforme tolerância gastrointestinal. Reavaliar níveis séricos após 8-12 semanas. Investigar e tratar causas subjacentes: ajustar diuréticos se possível, tratar má-absorção intestinal, corrigir deficiência de vitamina D. Enfatizar aumento de ingestão alimentar de fontes ricas em magnésio: vegetais folhosos verde-escuros (espinafre, couve), grãos integrais, leguminosas (feijão, lentilha), oleaginosas (amêndoas, castanhas), sementes (abóbora, girassol), abacate, banana, cacau. HIPOMAGNESEMIA GRAVE (<1,2 mg/dL) OU SINTOMÁTICA: Reposição intravenosa com sulfato de magnésio 1-2g em 15-30 minutos sob monitorização cardíaca, seguido de manutenção IV ou transição para via oral assim que possível. Monitorar cálcio e potássio séricos concomitantemente, pois hipomagnesemia frequentemente coexiste com hipocalcemia e hipocalemia refratárias. MANEJO CARDIOVASCULAR: Para pacientes hipertensos com hipomagnesemia, suplementação de 400-500 mg/dia pode reduzir pressão sistólica em 2-3 mmHg e diastólica em 1,78 mmHg. Manter nível sérico alvo ≥0,85 mmol/L (≈2,06 mg/dL). Em pacientes com risco cardiovascular elevado, considerar suplementação preventiva mesmo com níveis limítrofes. SEGUIMENTO: Reavaliar níveis a cada 3-6 meses durante reposição. Monitorar função renal antes e durante suplementação (contraindicada em insuficiência renal grave). Ajustar tratamento conforme resposta clínica e laboratorial."""

    cur.execute("""
        UPDATE score_items
        SET clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = CURRENT_TIMESTAMP
        WHERE id = '2813e3c1-a2d0-464d-ad05-35aec75a3d7d'
    """, (clinical_relevance, patient_explanation, conduct))

    print(f"✓ Score item atualizado ({cur.rowcount} linha)")

    # 2. Insert articles
    articles = [
        {
            'id': 'f8e9a3d2-5c7b-4e1f-9d8a-3c4b5a6e7f8d',
            'title': 'The Role of Dietary Magnesium in Cardiovascular Disease',
            'authors': 'Forrest H. Nielsen',
            'journal': 'Nutrients',
            'publish_date': '2024-12-06',
            'doi': '10.3390/nu16234223',
            'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11644202/',
            'abstract': 'Estudo de revisão de 2024 demonstrando relação inversa entre ingestão/status de magnésio e doenças cardiovasculares. Aproximadamente 50% dos adultos americanos consomem magnésio abaixo dos níveis recomendados. Status inadequado de magnésio está associado a hipertensão, calcificação arterial coronariana, AVC, doença cardíaca isquêmica, fibrilação atrial, insuficiência cardíaca e mortalidade cardíaca. O conceito de "deficiência latente crônica" (magnésio corporal total reduzido com níveis séricos normais) contribui significativamente para desenvolvimento de DCV. Suplementação ≥400 mg/dia por ≥12 semanas mostrou eficácia na redução da pressão arterial. Recomendações atualizadas sugerem aproximadamente 250 mg/dia para adultos de peso médio, com ajustes conforme peso corporal.'
        },
        {
            'id': 'a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d',
            'title': 'Neuroprotective effects of magnesium: implications for neuroinflammation and cognitive decline',
            'authors': 'Veer Patel, Mohammed S. Razzaque, Nuraly S Akimbekov, William B Grant, Carolyn Dean, Xiaoqian Fang',
            'journal': 'Frontiers in Endocrinology',
            'publish_date': '2024-01-01',
            'doi': '10.3389/fendo.2024.1461281',
            'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11461281/',
            'abstract': 'Artigo de 2024 sobre efeitos neuroprotetores do magnésio. O mineral modula neuroinflamação através de múltiplos mecanismos: mantém homeostase de cálcio neuronal prevenindo excitotoxicidade; inibe ativação de NF-κB reduzindo expressão de citocinas pró-inflamatórias; estabiliza superóxido dismutase (SOD) potencializando defesa antioxidante endógena; regula processos intracelulares afetando densidade e plasticidade sináptica, essenciais para memória e aprendizado. Deficiência de magnésio correlaciona-se com ativação neuroinflamatória e comprometimento cognitivo. Estudos pré-clínicos demonstram benefícios em modelos de Alzheimer e Parkinson, embora ensaios clínicos em humanos apresentem resultados heterogêneos, enfatizando necessidade de pesquisas mais rigorosas.'
        },
        {
            'id': 'b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e',
            'title': 'Hypomagnesemia: exploring its multifaceted health impacts and associations with blood pressure regulation and metabolic syndrome',
            'authors': 'Wenlong Wu et al.',
            'journal': 'Diabetology & Metabolic Syndrome',
            'publish_date': '2025-06-16',
            'doi': '10.1186/s13098-025-01519-4',
            'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12168336/',
            'abstract': 'Publicação recente de 2025 investigando impactos multifacetados da hipomagnesemia. Identifica múltiplos mecanismos ligando deficiência de magnésio à hipertensão: potencializa vasodilatação via produção de óxido nítrico e ativação de proteína quinase B; regula canais de cálcio e Na+/K+-ATPase mantendo relaxamento vascular; deficiência promove liberação de catecolaminas aumentando vasoconstrição; baixo magnésio upregula aldosterona promovendo retenção de sódio. Para cada aumento de 0,5 mg/dL no magnésio sérico, risco de hipertensão diminui 7%. Hipomagnesemia associa-se com resistência insulínica, exacerba inflamação relacionada à obesidade e dislipidemia. Recomendações clínicas: suplementação diária de 300-450 mg (citrato/glicinato de magnésio), redução esperada de 2-3 mmHg sistólica e 1,78 mmHg diastólica. Limiar sérico ótimo revisado para ≥0,85 mmol/L. Para síndrome metabólica: 400-500 mg/dia por mínimo 3 meses.'
        },
        {
            'id': 'c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f',
            'title': 'A Comprehensive Review on Understanding Magnesium Disorders: Pathophysiology, Clinical Manifestations, and Management Strategies',
            'authors': 'Kothari M, Wanjari A, Shaikh SM, Tantia P, Waghmare BV, Parepalli A, Hamdulay KF, Nelakuditi M',
            'journal': 'Cureus',
            'publish_date': '2024-01-01',
            'doi': '10.7759/cureus.69516',
            'original_link': 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11444808/',
            'abstract': 'Revisão abrangente de 2024 sobre distúrbios do magnésio. Fisiopatologia da hipomagnesemia decorre de ingestão dietética inadequada, perdas gastrointestinais ou excreção renal excessiva. Aproximadamente 20-30% do magnésio filtrado é reabsorvido no túbulo proximal, com taxas variáveis ao longo do néfron. Hormônios como paratormônio e vitamina D influenciam homeostase do magnésio. Manifestações clínicas: neuromusculares (cãibras, tetania); cardiovasculares (arritmias, hipertensão); metabólicas (resistência insulínica, hipocalcemia); reduz limiar convulsivo em populações suscetíveis. Estratégias de manejo: suplementação oral como terapia de primeira linha para deficiência leve-moderada (óxido ou citrato de magnésio); terapia intravenosa reservada para casos graves (1-2g de sulfato de magnésio em 15-30 minutos); tratar causas subjacentes (síndromes de má-absorção, ajustar medicações) essencial para prevenir recorrência.'
        }
    ]

    articles_inserted = 0
    for article in articles:
        cur.execute("""
            INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, abstract, created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
            ON CONFLICT (id) DO NOTHING
        """, (
            article['id'], article['title'], article['authors'], article['journal'],
            article['publish_date'], article['doi'], article['original_link'], article['abstract']
        ))
        if cur.rowcount > 0:
            articles_inserted += 1

    print(f"✓ Artigos inseridos: {articles_inserted}/4")

    # 3. Link articles to score_item
    score_item_id = '2813e3c1-a2d0-464d-ad05-35aec75a3d7d'
    links_created = 0

    for article in articles:
        cur.execute("""
            INSERT INTO article_score_items (article_id, score_item_id)
            VALUES (%s, %s)
            ON CONFLICT (score_item_id, article_id) DO NOTHING
        """, (article['id'], score_item_id))
        if cur.rowcount > 0:
            links_created += 1

    print(f"✓ Links artigo-item criados: {links_created}/4")

    # Commit transaction
    conn.commit()
    print("\n✓ Transação commitada com sucesso!")

    # Verification
    print("\n=== VERIFICAÇÃO ===")
    cur.execute("""
        SELECT
            si.name,
            si.unit,
            LENGTH(si.clinical_relevance) as clinical_rel_len,
            LENGTH(si.patient_explanation) as patient_exp_len,
            LENGTH(si.conduct) as conduct_len,
            si.last_review,
            COUNT(asi.article_id) as article_count
        FROM score_items si
        LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
        WHERE si.id = %s
        GROUP BY si.id, si.name, si.unit, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review
    """, (score_item_id,))

    result = cur.fetchone()
    print(f"\nNome: {result[0]}")
    print(f"Unidade: {result[1]}")
    print(f"Relevância Clínica: {result[2]} caracteres")
    print(f"Explicação Paciente: {result[3]} caracteres")
    print(f"Conduta: {result[4]} caracteres")
    print(f"Última Revisão: {result[5]}")
    print(f"Artigos Vinculados: {result[6]}")

    print("\n✓ Enriquecimento concluído com sucesso!")

except Exception as e:
    conn.rollback()
    print(f"\n✗ ERRO: {e}")
    raise
finally:
    cur.close()
    conn.close()
