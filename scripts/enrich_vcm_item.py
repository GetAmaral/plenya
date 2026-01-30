#!/usr/bin/env python3
"""
Enriquecimento do item VCM (MCV) - Volume Corpuscular Médio
Inclui pesquisa científica, salvamento de artigos e atualização do banco
"""

import psycopg2
import json
from datetime import datetime

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

# ID do item VCM
ITEM_ID = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'

# Artigos científicos coletados
ARTICLES = [
    {
        'title': 'Mean Corpuscular Volume - StatPearls',
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK545275/',
        'authors': 'StatPearls Publishing',
        'journal': 'NCBI Bookshelf',
        'year': 2024,
        'type': 'review',
        'summary': 'Revisão abrangente sobre VCM como medida crítica para identificar a causa subjacente de anemia. Descreve valores normais (80-100 fL), classificação de anemias (microcítica, normocítica, macrocítica), causas comuns e abordagem diagnóstica.'
    },
    {
        'title': 'Anemia - StatPearls',
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK499994/',
        'authors': 'StatPearls Publishing',
        'journal': 'NCBI Bookshelf',
        'year': 2024,
        'type': 'review',
        'summary': 'Classificação completa de anemia baseada em VCM. Detalha fisiopatologia, diagnóstico diferencial por categoria de VCM e princípios de manejo. Inclui contagem corrigida de reticulócitos para diferenciar produção inadequada de hemólise.'
    },
    {
        'title': 'Normal and Abnormal Complete Blood Count With Differential',
        'url': 'https://www.ncbi.nlm.nih.gov/books/NBK604207/',
        'authors': 'StatPearls Publishing',
        'journal': 'NCBI Bookshelf',
        'year': 2024,
        'type': 'review',
        'summary': 'Abordagem prática para interpretação de hemograma completo, incluindo uso de VCM com diretrizes WHO 2024 para definição de anemia. Enfatiza abordagem multiparamétrica para reduzir erros de classificação.'
    }
]

# Conteúdo clínico em PT-BR
CLINICAL_RELEVANCE = """**VCM (Volume Corpuscular Médio) - Classificação de Anemias**

O VCM é uma medida fundamental no hemograma que indica o tamanho médio das hemácias, sendo crítico para identificar a causa subjacente de anemia e orientar estratégias de tratamento.

**Valores de Referência:**
• Normal: 80-100 fL (femtolitros)
• Microcitose: VCM < 80 fL (hemácias pequenas)
• Macrocitose: VCM > 100 fL (hemácias grandes)

**Classificação por VCM:**

**Anemia Microcítica (VCM < 80 fL):**
• Deficiência de ferro (causa mais comum)
• Anemia de doença crônica
• Talassemia
• Anemia sideroblástica
• Intoxicação por chumbo

**Anemia Normocítica (VCM 80-100 fL):**
• Anemia de doença renal
• Anemia aplástica
• Perda sanguínea aguda
• Anemia hemolítica
• Mielofibrose

**Anemia Macrocítica (VCM > 100 fL):**
• Deficiência de vitamina B12
• Deficiência de ácido fólico
• Doença hepática/cirrose
• Hipotireoidismo
• Síndrome mielodisplásica
• Medicamentos (metotrexato, zidovudina)

**Interpretação Clínica:**
• VCM 100-110 fL (macrocitose leve): mais provável causa benigna
• VCM > 110 fL (macrocitose marcada): sugere doença primária da medula óssea ou deficiência megaloblástica

**Abordagem Diagnóstica:**
1. Confirmar anemia (WHO 2024: Hb < 13 g/dL homens, < 12 g/dL mulheres)
2. Hemograma completo com VCM, RDW, reticulócitos
3. Contagem corrigida de reticulócitos: >2% sugere hemólise, <2% indica hipoproliferação
4. Testes adicionais baseados na categoria de VCM:
   - Microcítica: ferritina, saturação de transferrina, eletroforese de Hb
   - Macrocítica: B12, folato, TSH, função hepática
   - Normocítica: função renal, esfregaço, mielograma

**Considerações Importantes:**
• Usar VCM em conjunto com história clínica, exame físico e outros parâmetros
• Um terço dos idosos pode ter VCM elevado sem causa identificável
• Abordagem multiparamétrica reduz erros de classificação
• Evitar interpretação isolada do VCM

**Fórmula de Cálculo:**
VCM (fL) = (Hematócrito %) / (Contagem de hemácias × 10¹²/L) × 10

**Referências Científicas:**
• StatPearls - Mean Corpuscular Volume (2024)
• StatPearls - Anemia (2024)
• WHO Guidelines for Hemoglobin Cutoffs (2024)
"""

PATIENT_EXPLANATION = """**VCM - O Que É e Por Que É Importante**

O VCM (Volume Corpuscular Médio) mede o tamanho médio das suas hemácias (glóbulos vermelhos). É como medir o tamanho dos seus glóbulos vermelhos para entender por que você pode estar com anemia.

**Valores Normais:**
• Normal: 80 a 100 fL
• Abaixo de 80: suas hemácias estão menores que o normal
• Acima de 100: suas hemácias estão maiores que o normal

**Por Que o Tamanho das Hemácias Importa?**

O tamanho das hemácias ajuda seu médico a descobrir a CAUSA da anemia:

**Hemácias Pequenas (VCM baixo):**
Geralmente significa que você está com falta de ferro ou tem talassemia (problema genético nas hemácias). É como construir uma casa sem tijolos suficientes - ela fica menor.

**Hemácias Grandes (VCM alto):**
Pode indicar falta de vitaminas B12 ou ácido fólico, problemas no fígado ou tireoide. É como células que não conseguem se dividir direito e ficam grandes demais.

**Hemácias Normais (VCM normal) mas com Anemia:**
Pode ser perda de sangue recente, problema nos rins ou destruição das hemácias.

**O Que Fazer Se Meu VCM Estiver Alterado?**

1. **Não se preocupe sozinho**: VCM alterado não é um diagnóstico final, é uma pista
2. **Seu médico vai pedir outros exames**: ferritina (ferro), vitamina B12, ácido fólico
3. **O tratamento depende da causa**:
   - VCM baixo: geralmente suplementação de ferro
   - VCM alto: vitamina B12 ou ácido fólico
   - Alguns casos precisam de avaliação especializada

**Dica Importante:**
Pessoas idosas podem ter VCM um pouco elevado sem problema grave. Sempre analise junto com seu médico considerando seu histórico completo de saúde.

**Quando o VCM Volta ao Normal?**
Depende da causa:
• Falta de ferro: 2-3 meses de tratamento
• Falta de B12: 4-8 semanas
• Problemas crônicos: pode precisar tratamento contínuo
"""

CONDUCT = """**Conduta Clínica Baseada em VCM**

**1. ANEMIA MICROCÍTICA (VCM < 80 fL)**

**Investigação Inicial:**
□ Ferritina sérica (< 30 ng/mL sugere deficiência de ferro)
□ Saturação de transferrina (< 20% sugere deficiência)
□ Proteína C reativa (avaliar inflamação)
□ Hemoglobina A2 e F (suspeita de talassemia)

**Conduta:**
• Deficiência de ferro confirmada:
  - Sulfato ferroso 300mg VO 1-2x/dia (ou formulação IV se intolerância)
  - Investigar causa da deficiência (sangramento GI, menstruação, dieta)
  - Reavaliar hemograma em 4-8 semanas
• Talassemia suspeita: encaminhar para hematologia + aconselhamento genético
• Anemia de doença crônica: tratar doença de base

**2. ANEMIA MACROCÍTICA (VCM > 100 fL)**

**Investigação Inicial:**
□ Vitamina B12 sérica (< 200 pg/mL = deficiência)
□ Ácido fólico sérico (< 4 ng/mL = deficiência)
□ TSH (rastreio hipotireoidismo)
□ Função hepática (TGO, TGP, bilirrubinas, albumina)
□ Contagem de reticulócitos corrigida
□ Esfregaço de sangue periférico

**Classificação:**
• **Macrocitose leve (100-110 fL)**: investigar causas benignas (álcool, medicamentos)
• **Macrocitose marcada (> 110 fL)**: suspeitar deficiência megaloblástica ou síndrome mielodisplásica

**Conduta:**
• Deficiência de B12:
  - Cianocobalamina 1000 mcg IM/SC diário por 1 semana
  - Depois: 1000 mcg semanal por 4 semanas
  - Manutenção: 1000 mcg mensal VO ou IM
  - Investigar causa (anemia perniciosa, malabsorção, dieta)
• Deficiência de folato:
  - Ácido fólico 1-5 mg VO diário
  - SEMPRE repor B12 antes ou junto com folato (risco de degeneração medular)
• VCM > 115 fL persistente: encaminhar para hematologia (avaliar mielograma)

**3. ANEMIA NORMOCÍTICA (VCM 80-100 fL)**

**Investigação Inicial:**
□ Contagem de reticulócitos corrigida (diferenciar produção vs destruição)
□ Creatinina e ureia (função renal)
□ Bilirrubinas e LDH (hemólise)
□ Esfregaço de sangue periférico
□ Haptoglobina (se suspeita de hemólise)

**Conduta:**
• Reticulócitos > 2% (anemia hemolítica):
  - Encaminhar para hematologia
  - Investigar causa: autoimune, medicamentos, hemoglobinopatias
• Reticulócitos < 2% (hipoproliferativa):
  - Doença renal crônica: considerar eritropoetina se TFG < 30 mL/min
  - Anemia aplástica suspeita: URGENTE - encaminhar hematologia + mielograma
  - Mielofibrose/mieloma: encaminhar para hematologia

**4. ALGORITMO GERAL DE SEGUIMENTO**

**Semana 0 (Diagnóstico):**
• Hemograma completo + reticulócitos
• Testes específicos baseados em VCM
• Iniciar tratamento empírico se indicado

**Semana 4-8:**
• Reavaliar hemograma
• Resposta adequada: aumento Hb > 1 g/dL ou reticulócitos aumentados
• Sem resposta: considerar diagnóstico alternativo ou encaminhar hematologia

**Semana 12-16:**
• Confirmar normalização dos parâmetros
• Avaliar necessidade de manutenção (B12, ferro)

**5. CRITÉRIOS DE ENCAMINHAMENTO PARA HEMATOLOGIA**

**Urgente (< 1 semana):**
• VCM > 115 fL sem causa óbvia
• Pancitopenia associada
• Blastos ou células anormais no esfregaço
• Anemia grave (Hb < 7 g/dL) sem sangramento evidente

**Eletivo (< 4 semanas):**
• Anemia refratária ao tratamento após 8-12 semanas
• Talassemia ou hemoglobinopatia suspeita
• Macrocitose persistente sem causa identificável
• História familiar de anemia ou malignidade hematológica

**6. MONITORAMENTO EM LONGO PRAZO**

• Deficiência de ferro: hemograma anual + ferritina (manter > 50 ng/mL)
• Deficiência de B12: hemograma semestral (se em reposição contínua)
• Anemia de doença crônica: seguir protocolo da doença de base
• Idosos com VCM limítrofe: hemograma anual de rotina

**7. CONSIDERAÇÕES ESPECIAIS**

**Gestantes:**
• VCM pode diminuir fisiologicamente (hemodiluição)
• Iniciar ferro profilático se ferritina < 30 ng/mL

**Idosos:**
• VCM elevado isolado (< 105 fL) pode ser fisiológico
• Investigar se sintomático ou progressivo

**Pacientes em Quimioterapia:**
• Macrocitose esperada com alguns agentes (metotrexato, hidroxiureia)
• Seguir protocolo oncológico

**Referências:**
• WHO 2024 - Hemoglobin Cutoffs for Anemia Definition
• UpToDate 2024 - Approach to the Adult with Anemia
• ASH Guidelines - Iron Deficiency Anemia (2024)
"""


def create_connection():
    """Cria conexão com o banco de dados"""
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        return conn
    except Exception as e:
        print(f"❌ Erro ao conectar ao banco: {e}")
        return None


def insert_article(conn, article):
    """Insere um artigo científico no banco"""
    cursor = conn.cursor()

    try:
        # Verificar se artigo já existe
        cursor.execute(
            "SELECT id FROM articles WHERE original_link = %s",
            (article['url'],)
        )
        result = cursor.fetchone()

        if result:
            article_id = result[0]
            print(f"  ℹ️  Artigo já existe: {article['title'][:50]}... (ID: {article_id})")
            return article_id

        # Inserir novo artigo
        # publish_date: usando 01/01/year
        publish_date = f"{article['year']}-01-01"

        cursor.execute("""
            INSERT INTO articles (
                title, original_link, authors, journal, publish_date,
                article_type, abstract, created_at, updated_at
            ) VALUES (%s, %s, %s, %s, %s, %s, %s, NOW(), NOW())
            RETURNING id
        """, (
            article['title'],
            article['url'],
            article['authors'],
            article['journal'],
            publish_date,
            article['type'],
            article['summary']
        ))

        article_id = cursor.fetchone()[0]
        conn.commit()
        print(f"  ✅ Artigo inserido: {article['title'][:50]}... (ID: {article_id})")
        return article_id

    except Exception as e:
        conn.rollback()
        print(f"  ❌ Erro ao inserir artigo: {e}")
        return None


def link_article_to_item(conn, article_id, item_id):
    """Cria relação entre artigo e score_item"""
    cursor = conn.cursor()

    try:
        # Verificar se relação já existe
        cursor.execute("""
            SELECT 1 FROM article_score_items
            WHERE score_item_id = %s AND article_id = %s
        """, (item_id, article_id))

        if cursor.fetchone():
            print(f"    ℹ️  Relação já existe")
            return True

        # Criar relação
        cursor.execute("""
            INSERT INTO article_score_items (score_item_id, article_id)
            VALUES (%s, %s)
        """, (item_id, article_id))

        conn.commit()
        print(f"    ✅ Relação criada: item -> artigo")
        return True

    except Exception as e:
        conn.rollback()
        print(f"    ❌ Erro ao criar relação: {e}")
        return False


def update_score_item(conn, item_id):
    """Atualiza os campos clínicos do score_item"""
    cursor = conn.cursor()

    try:
        cursor.execute("""
            UPDATE score_items
            SET
                clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                updated_at = NOW()
            WHERE id = %s
        """, (
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            item_id
        ))

        rows_affected = cursor.rowcount
        conn.commit()

        if rows_affected > 0:
            print(f"✅ Score item atualizado com sucesso!")
            return True
        else:
            print(f"❌ Nenhum item encontrado com ID: {item_id}")
            return False

    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao atualizar score item: {e}")
        return False


def verify_update(conn, item_id):
    """Verifica se a atualização foi bem-sucedida"""
    cursor = conn.cursor()

    try:
        cursor.execute("""
            SELECT
                name,
                clinical_relevance IS NOT NULL as has_relevance,
                patient_explanation IS NOT NULL as has_explanation,
                conduct IS NOT NULL as has_conduct,
                LENGTH(clinical_relevance) as relevance_length,
                LENGTH(patient_explanation) as explanation_length,
                LENGTH(conduct) as conduct_length,
                updated_at
            FROM score_items
            WHERE id = %s
        """, (item_id,))

        result = cursor.fetchone()

        if result:
            print("\n📊 VERIFICAÇÃO DO ITEM ATUALIZADO:")
            print(f"  Nome: {result[0]}")
            print(f"  Clinical Relevance: {'✅' if result[1] else '❌'} ({result[4]} caracteres)")
            print(f"  Patient Explanation: {'✅' if result[2] else '❌'} ({result[5]} caracteres)")
            print(f"  Conduct: {'✅' if result[3] else '❌'} ({result[6]} caracteres)")
            print(f"  Atualizado em: {result[7]}")

            # Verificar artigos vinculados
            cursor.execute("""
                SELECT COUNT(*)
                FROM article_score_items
                WHERE score_item_id = %s
            """, (item_id,))

            article_count = cursor.fetchone()[0]
            print(f"  Artigos vinculados: {article_count}")

            return True
        else:
            print(f"❌ Item não encontrado: {item_id}")
            return False

    except Exception as e:
        print(f"❌ Erro ao verificar: {e}")
        return False


def main():
    """Função principal"""
    print("=" * 80)
    print("ENRIQUECIMENTO DO ITEM: VCM (MCV)")
    print("=" * 80)

    # Conectar ao banco
    print("\n1️⃣  Conectando ao banco de dados...")
    conn = create_connection()

    if not conn:
        print("❌ Falha na conexão. Abortando.")
        return

    print("✅ Conectado com sucesso!")

    # Inserir artigos
    print(f"\n2️⃣  Inserindo {len(ARTICLES)} artigos científicos...")
    article_ids = []

    for i, article in enumerate(ARTICLES, 1):
        print(f"\n  [{i}/{len(ARTICLES)}] Processando: {article['title'][:60]}...")
        article_id = insert_article(conn, article)

        if article_id:
            article_ids.append(article_id)
            # Criar relação com o item
            link_article_to_item(conn, article_id, ITEM_ID)

    print(f"\n✅ {len(article_ids)} artigos processados com sucesso!")

    # Atualizar score_item
    print(f"\n3️⃣  Atualizando campos clínicos do item VCM...")
    success = update_score_item(conn, ITEM_ID)

    if not success:
        print("❌ Falha ao atualizar item. Abortando.")
        conn.close()
        return

    # Verificar resultados
    print(f"\n4️⃣  Verificando atualização...")
    verify_update(conn, ITEM_ID)

    # Fechar conexão
    conn.close()

    print("\n" + "=" * 80)
    print("✅ PROCESSO CONCLUÍDO COM SUCESSO!")
    print("=" * 80)
    print(f"\nResumo:")
    print(f"  • Item ID: {ITEM_ID}")
    print(f"  • Artigos inseridos/vinculados: {len(article_ids)}")
    print(f"  • Campos atualizados: clinical_relevance, patient_explanation, conduct")
    print(f"  • Status: ✅ COMPLETO")
    print("\n" + "=" * 80)


if __name__ == "__main__":
    main()
