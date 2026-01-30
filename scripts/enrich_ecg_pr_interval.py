#!/usr/bin/env python3
"""
Enriquecimento científico: ECG - Intervalo PR
ID: b2dd0c76-7bce-4beb-a8e2-52d70d467241
Grupo: Exames > Imagem
"""

import psycopg2
from datetime import datetime
import sys

# Configuração do banco
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

# ID do item
ITEM_ID = 'b2dd0c76-7bce-4beb-a8e2-52d70d467241'

# IDs dos artigos relacionados (cardiologia)
ARTICLE_IDS = [
    '5f6a3374-d88d-4f9f-9abd-97906a74919d',  # ACC/AHA/HRS Guideline - Bradycardia and Cardiac Conduction Delay
    'eddc9921-0f50-406b-aea4-b2b37594385c',  # Heart Rate Variability in Cardiovascular Disease
    'd90edaac-a622-42f3-b02a-2de1ccd77a10'   # Heart Rate Variability: Standards of Measurement
]

# Conteúdo clínico em PT-BR
CLINICAL_RELEVANCE = """O intervalo PR no eletrocardiograma representa o tempo de condução atrioventricular, refletindo a propagação do impulso elétrico desde o início da despolarização atrial até o início da despolarização ventricular. O valor normal situa-se entre 120 e 200 ms (0,12-0,20 segundos), correspondendo a 3-5 quadradinhos no ECG padrão.

**Valores e Interpretação Clínica:**

• **Normal (120-200 ms):** Condução AV preservada, sem bloqueios ou vias acessórias
• **Prolongado (>200 ms):** Bloqueio atrioventricular de 1º grau, indicando condução lentificada através do nó AV
• **Muito prolongado (>300 ms):** BAV 1º grau marcado, podendo causar sintomas de dessincronização AV (pseudo-síndrome do marcapasso)
• **Curto (<120 ms):** Sugere pré-excitação ventricular (síndrome de Wolff-Parkinson-White) ou ritmo juncional

**Significância Fisiopatológica:**

O nó atrioventricular funciona como um "filtro" fisiológico que:
1. Retarda a condução elétrica entre átrios e ventrículos
2. Permite enchimento ventricular adequado durante a diástole
3. Protege os ventrículos de frequências atriais excessivas (ex: fibrilação atrial)

Estudos de coorte demonstraram que o BAV de 1º grau não é totalmente benigno: pacientes com PR prolongado apresentam maior incidência de fibrilação atrial, necessidade de marcapasso permanente, insuficiência cardíaca e mortalidade por todas as causas. Em casos extremos (PR >300 ms), a perda da sincronização AV pode causar intolerância ao exercício, dispneia e fadiga semelhante à síndrome do marcapasso.

A síndrome de WPW caracteriza-se pela presença de uma via acessória que bypassa o nó AV, resultando em PR curto (<120 ms), onda delta (upstroke inicial do QRS lentificado) e QRS alargado (>120 ms). Esta condição está associada a risco de taquiarritmias paroxísticas potencialmente fatais.

**Referências Científicas:**
• ACC/AHA/HRS 2018 Guidelines on Management of Bradycardia and Cardiac Conduction Delay (Circulation)
• LITFL ECG Library: PR Interval - Normal values and clinical interpretation
• StatPearls: First-Degree Heart Block - Clinical significance and management
• JAMA: Long-term Outcomes in Individuals With Prolonged PR Interval"""

PATIENT_EXPLANATION = """O intervalo PR é uma medida que o eletrocardiograma (ECG) mostra sobre como o estímulo elétrico do seu coração viaja desde os átrios (câmaras superiores) até os ventrículos (câmaras inferiores). Pense nisso como o tempo que a "mensagem elétrica" leva para passar de cima para baixo.

**O que significa cada resultado:**

• **Normal (0,12 a 0,20 segundos):** Seu coração está conduzindo os impulsos elétricos de forma adequada. A sincronia entre átrios e ventrículos está perfeita, permitindo um bombeamento eficiente do sangue.

• **Prolongado (mais de 0,20 segundos):** Chamamos isso de "bloqueio de primeiro grau". Significa que o impulso elétrico está demorando um pouco mais para passar pelo nó atrioventricular (uma espécie de "estação" entre os átrios e ventrículos). Na maioria dos casos, isso não causa sintomas e muitas pessoas vivem normalmente com esta condição.

• **Muito prolongado (mais de 0,30 segundos):** Em casos raros, quando o atraso é muito grande, você pode sentir cansaço, falta de ar ao se exercitar ou tonturas. Isso acontece porque os átrios e ventrículos ficam fora de sincronia, prejudicando a eficiência do bombeamento do sangue.

• **Curto (menos de 0,12 segundos):** Pode indicar uma condição chamada síndrome de Wolff-Parkinson-White, onde existe um "atalho elétrico" no coração que faz o impulso chegar mais rápido aos ventrículos. Pessoas com esta condição podem ter episódios de palpitações rápidas.

**Por que é importante monitorar:**

Embora muitas vezes o intervalo PR alterado não cause problemas imediatos, estudos mostram que pessoas com PR prolongado têm maior risco de desenvolver fibrilação atrial (um tipo de arritmia) ou precisar de marcapasso no futuro. Por isso, é importante fazer acompanhamento regular com seu cardiologista.

**Quando se preocupar:**

Procure atendimento médico se você sentir:
• Palpitações (sensação de coração acelerado ou batendo irregular)
• Tonturas frequentes ou desmaios
• Falta de ar ao fazer esforços que antes não causavam sintomas
• Dor no peito ou cansaço excessivo"""

CONDUCT = """**Avaliação Inicial:**

1. **História Clínica Detalhada:**
   - Sintomas: palpitações, síncope, pré-síncope, dispneia aos esforços, fadiga
   - Medicações em uso (betabloqueadores, bloqueadores de canal de cálcio, digitálicos, antiarrítmicos)
   - Antecedentes: cardiopatia estrutural, miocardite, distúrbios metabólicos
   - História familiar de morte súbita ou cardiomiopatias

2. **Exame Físico:**
   - Sinais vitais: FC, PA, saturação O2
   - Ausculta cardíaca: sopros, B3/B4, ritmo irregular
   - Sinais de insuficiência cardíaca: edema, turgência jugular, crepitações

**Estratificação por Valores do Intervalo PR:**

**PR Normal (120-200 ms):**
• Nenhuma conduta adicional necessária
• Manter seguimento cardiológico de rotina

**PR Prolongado (200-300 ms) - BAV 1º Grau:**
• Assintomático:
  - Observação clínica e ECG anual
  - Investigar causas reversíveis: hipocalemia, hipermagnesemia, hipotireoidismo
  - Revisar medicações cronotrópicas negativas
  - Ecocardiograma se não realizado recentemente

• Sintomático (fadiga, dispneia aos esforços):
  - Holter 24h para avaliar progressão do bloqueio
  - Teste ergométrico para avaliar resposta cronotrópica
  - Considerar estudo eletrofisiológico se QRS também alargado
  - Discussão sobre indicação de marcapasso (Classe IIa se PR >300 ms + sintomas)

**PR Muito Prolongado (>300 ms) - BAV 1º Grau Marcado:**
• Avaliação especializada obrigatória (arritmologista)
• Holter 24h + teste ergométrico
• Ecocardiograma com Doppler para avaliar sincronização AV
• Considerar marcapasso definitivo se:
  - Sintomas claramente relacionados à dessincronização AV
  - Comprometimento hemodinâmico (pseudo-síndrome do marcapasso)
  - Progressão para BAV de grau superior

**PR Curto (<120 ms) - Suspeita de Pré-excitação:**
• ECG de 12 derivações de alta qualidade
• Procurar onda delta (upstroke inicial do QRS lentificado)
• Se onda delta presente → Wolff-Parkinson-White:
  - Referência urgente para arritmologista
  - Estudo eletrofisiológico para mapear via acessória
  - Discussão sobre ablação por radiofrequência (tratamento curativo)
  - Evitar adenosina, betabloqueadores e bloqueadores de canal de cálcio se FA presente
• Se onda delta ausente → Provável ritmo juncional ou variante normal

**BAV de 2º ou 3º Grau (se evolução):**
• Internação hospitalar com monitorização contínua
• Marcapasso transcutâneo de prontidão
• Marcapasso definitivo conforme guidelines ACC/AHA/HRS 2018

**Monitoramento de Longo Prazo:**
• ECG anual para pacientes com BAV 1º grau
• Orientar paciente sobre sinais de alerta
• Reavaliar medicações a cada consulta
• Ecocardiograma a cada 2-3 anos se cardiopatia estrutural associada

**Indicações de Marcapasso (ACC/AHA/HRS 2018):**
• Classe I: BAV 2º grau tipo II, BAV 3º grau sintomático
• Classe IIa: BAV 1º grau com PR >300 ms + sintomas de dessincronização AV
• Classe IIb: BAV 1º grau assintomático com cardiomiopatia e FEVE reduzida

**Condutas NÃO Recomendadas:**
• Marcapasso profilático em BAV 1º grau assintomático sem progressão (evidência insuficiente)
• Suspensão automática de betabloqueadores/BCC sem avaliar benefício vs. risco
• Negligenciar seguimento em pacientes com PR >200 ms"""

def enrich_item():
    """Enriquece o item ECG - Intervalo PR no banco de dados"""

    conn = None
    try:
        # Conectar ao banco
        print("Conectando ao banco de dados...")
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        # 1. Verificar se o item existe
        cur.execute("SELECT name FROM score_items WHERE id = %s", (ITEM_ID,))
        result = cur.fetchone()
        if not result:
            print(f"ERRO: Item {ITEM_ID} não encontrado!")
            return False

        print(f"✓ Item encontrado: {result[0]}")

        # 2. Atualizar campos clínicos
        print("\nAtualizando campos clínicos...")
        cur.execute("""
            UPDATE score_items
            SET clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                last_review = %s,
                updated_at = CURRENT_TIMESTAMP
            WHERE id = %s
        """, (CLINICAL_RELEVANCE, PATIENT_EXPLANATION, CONDUCT, datetime.now(), ITEM_ID))

        print(f"✓ Campos clínicos atualizados")

        # 3. Verificar quais artigos existem
        print("\nVerificando artigos...")
        existing_articles = []
        for article_id in ARTICLE_IDS:
            cur.execute("SELECT title FROM articles WHERE id = %s AND deleted_at IS NULL", (article_id,))
            result = cur.fetchone()
            if result:
                existing_articles.append(article_id)
                print(f"  ✓ {result[0][:80]}...")
            else:
                print(f"  ✗ Article {article_id} não encontrado")

        # 4. Criar relações article_score_items
        print("\nCriando relações article_score_items...")
        relations_created = 0
        for article_id in existing_articles:
            # Verificar se relação já existe
            cur.execute("""
                SELECT 1 FROM article_score_items
                WHERE article_id = %s AND score_item_id = %s
            """, (article_id, ITEM_ID))

            if cur.fetchone():
                print(f"  - Relação com {article_id} já existe")
            else:
                cur.execute("""
                    INSERT INTO article_score_items (article_id, score_item_id)
                    VALUES (%s, %s)
                """, (article_id, ITEM_ID))
                relations_created += 1
                print(f"  ✓ Relação criada com {article_id}")

        # 5. Commit
        conn.commit()

        # 6. Relatório final
        print("\n" + "="*80)
        print("ENRIQUECIMENTO CONCLUÍDO COM SUCESSO")
        print("="*80)
        print(f"Item ID: {ITEM_ID}")
        print(f"Item: ECG - Intervalo PR")
        print(f"Artigos relacionados: {len(existing_articles)}")
        print(f"Novas relações criadas: {relations_created}")
        print(f"Campos atualizados: clinical_relevance, patient_explanation, conduct, last_review")
        print(f"Timestamp: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

        # 7. Verificar resultado
        print("\nVerificando resultado final...")
        cur.execute("""
            SELECT
                si.name,
                LENGTH(si.clinical_relevance) as clin_len,
                LENGTH(si.patient_explanation) as pat_len,
                LENGTH(si.conduct) as cond_len,
                si.last_review,
                COUNT(asi.article_id) as num_articles
            FROM score_items si
            LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
            WHERE si.id = %s
            GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review
        """, (ITEM_ID,))

        result = cur.fetchone()
        if result:
            print(f"  Nome: {result[0]}")
            print(f"  clinical_relevance: {result[1]} caracteres")
            print(f"  patient_explanation: {result[2]} caracteres")
            print(f"  conduct: {result[3]} caracteres")
            print(f"  last_review: {result[4]}")
            print(f"  Artigos relacionados: {result[5]}")

        cur.close()
        return True

    except Exception as e:
        print(f"\nERRO: {e}")
        if conn:
            conn.rollback()
        return False
    finally:
        if conn:
            conn.close()

if __name__ == "__main__":
    success = enrich_item()
    sys.exit(0 if success else 1)
