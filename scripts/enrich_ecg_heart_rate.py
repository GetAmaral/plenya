#!/usr/bin/env python3
"""
Enrichment script para ECG - Frequência Cardíaca
Item ID: 79a44be5-fa22-4ecf-8900-9cb51f007292
"""

import psycopg2
from datetime import datetime

# Database connection
DB_CONFIG = {
    'host': 'localhost',
    'port': '5432',
    'database': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_dev_password'
}

# Item ID
ITEM_ID = '79a44be5-fa22-4ecf-8900-9cb51f007292'

# Clinical content in Portuguese
CLINICAL_RELEVANCE = """A frequência cardíaca (FC) medida pelo ECG é um biomarcador cardiovascular fundamental que reflete a atividade do sistema nervoso autônomo e a função do nó sinusal.

VALORES DE REFERÊNCIA:
- Normal: 60-100 bpm em repouso
- Bradicardia: < 60 bpm
- Taquicardia: > 100 bpm

SIGNIFICADO CLÍNICO:

Bradicardia Sinusal:
- Pode ser fisiológica em atletas (adaptação cardiovascular)
- Patológica quando associada a disfunção do nó sinusal, bloqueios AV, hipotireoidismo, hipotermia ou medicações (betabloqueadores, bloqueadores de canais de cálcio)
- Guideline ACC/AHA/HRS 2018: bradicardia sintomática requer avaliação por ECG de 12 derivações e monitoramento ambulatorial

Taquicardia Sinusal:
- Resposta fisiológica a exercício, estresse, febre, desidratação
- Patológica quando persistente em repouso: hipertireoidismo, anemia, insuficiência cardíaca, embolia pulmonar
- Requer avaliação do QRS (estreito vs. largo) para diagnóstico diferencial

VARIABILIDADE DA FREQUÊNCIA CARDÍACA (HRV):
- Evidências 2025 mostram que HRV reduzida (SDNN < 70 ms ou LF/HF > 2.5) associa-se a risco 1,5-2,3x maior de eventos cardiovasculares adversos maiores (MACE)
- Hazard ratio para mortalidade por todas as causas: 2,27 (IC 95%: 1,72-3,00)
- Hazard ratio para eventos cardiovasculares: 1,41 (IC 95%: 1,16-1,72)

VALOR PROGNÓSTICO:
- FC de repouso elevada (> 80-85 bpm) é preditor independente de mortalidade cardiovascular
- Redução de HRV prediz complicações cardiovasculares em UTI com 24-48h de antecedência
- Monitoramento contínuo com wearables expande aplicações clínicas"""

PATIENT_EXPLANATION = """O eletrocardiograma (ECG) mede a frequência cardíaca, ou seja, quantas vezes seu coração bate por minuto. Este é um dos sinais vitais mais importantes para avaliar sua saúde cardiovascular.

O QUE É NORMAL:
- Em repouso, o coração deve bater entre 60 e 100 vezes por minuto
- Atletas podem ter valores mais baixos (50-60 bpm) e isso é saudável
- Durante exercício, a frequência pode subir até 150-180 bpm dependendo da idade e condicionamento

QUANDO A FREQUÊNCIA ESTÁ BAIXA (BRADICARDIA):
- Menos de 60 bpm em repouso
- Pode ser normal se você é atleta ou tem boa condição física
- Precisa atenção se acompanhada de: tontura, desmaios, cansaço excessivo, falta de ar

QUANDO A FREQUÊNCIA ESTÁ ALTA (TAQUICARDIA):
- Mais de 100 bpm em repouso
- Pode acontecer por: ansiedade, febre, desidratação, anemia, problemas de tireoide
- Precisa investigação se persistir mesmo em repouso completo

VARIABILIDADE DA FREQUÊNCIA:
- O coração saudável não bate como um metrônomo - há pequenas variações entre cada batimento
- Essa variabilidade é sinal de saúde cardiovascular
- Estudos recentes (2025) mostram que baixa variabilidade pode indicar risco aumentado para problemas cardíacos

O QUE PODE AFETAR SUA FREQUÊNCIA:
- Atividade física e nível de condicionamento
- Medicamentos (especialmente para pressão e coração)
- Temperatura corporal e ambiente
- Níveis de hidratação
- Estado emocional (estresse, ansiedade)
- Qualidade do sono
- Consumo de cafeína, álcool ou tabaco"""

CONDUCT = """INTERPRETAÇÃO DOS RESULTADOS:

1. FREQUÊNCIA CARDÍACA NORMAL (60-100 BPM):
   - Sem necessidade de intervenção imediata
   - Manter acompanhamento de rotina
   - Considerar análise de HRV se disponível

2. BRADICARDIA SINTOMÁTICA (< 60 BPM + SINTOMAS):
   AVALIAÇÃO IMEDIATA:
   - ECG 12 derivações completo
   - Avaliar medicações em uso (betabloqueadores, digitálicos, bloqueadores de cálcio)
   - Descartar causas reversíveis: hipotireoidismo, hipotermia, hipóxia, isquemia
   - Eletrólitos (K+, Mg2+, Ca2+)
   - TSH se suspeita de disfunção tireoidiana

   CONDUTA SEGUNDO ACC/AHA/HRS 2018:
   - Monitoramento Holter 24-48h se sintomas intermitentes
   - Encaminhar para cardiologista se bradicardia sintomática persistente
   - Considerar marcapasso se disfunção sinusal ou bloqueio AV avançado

3. BRADICARDIA ASSINTOMÁTICA (< 60 BPM SEM SINTOMAS):
   - Comum em atletas: FC até 40-50 bpm pode ser normal
   - Investigar história de atividade física regular
   - Reavaliação em 3-6 meses
   - Atentar para desenvolvimento de sintomas

4. TAQUICARDIA SINUSAL (> 100 BPM):
   PRIMEIRA ABORDAGEM:
   - Identificar causas fisiológicas: ansiedade, dor, febre, desidratação
   - Avaliar uso de substâncias: cafeína, descongestionantes, broncodilatadores
   - ECG 12 derivações para caracterizar ritmo (QRS estreito vs. largo)

   INVESTIGAÇÃO COMPLEMENTAR:
   - Hemograma completo (anemia)
   - TSH e T4 livre (hipertireoidismo)
   - Eletrólitos
   - Gasometria se suspeita de hipóxia
   - Ecocardiograma se suspeita de disfunção cardíaca estrutural

   ENCAMINHAMENTO:
   - Cardiologista se FC repouso persistentemente > 90-100 bpm sem causa aparente
   - Urgência se taquicardia com instabilidade hemodinâmica

5. VARIABILIDADE DA FREQUÊNCIA CARDÍACA (HRV):
   SE HRV DISPONÍVEL:
   - SDNN < 70 ms ou LF/HF > 2.5: risco aumentado para eventos CV
   - Considerar estratificação de risco cardiovascular completa
   - Otimizar controle de fatores de risco: HAS, DM, dislipidemia
   - Reforçar modificações de estilo de vida

   INTERVENÇÕES PARA MELHORAR HRV:
   - Exercício físico aeróbico regular (150 min/semana)
   - Controle de estresse (mindfulness, meditação)
   - Sono adequado (7-9 horas/noite)
   - Cessação de tabagismo
   - Redução de consumo de álcool

6. MONITORAMENTO CONTÍNUO:
   - Considerar uso de wearables para acompanhamento de tendências
   - Atenção: dispositivos PPG têm menor acurácia que ECG
   - Reavaliação por ECG profissional a cada 6-12 meses
   - Mais frequente se alterações detectadas

7. SITUAÇÕES DE EMERGÊNCIA (ACLS 2025):
   BRADICARDIA COM INSTABILIDADE:
   - Atropina 0,5-1 mg IV (até 3 mg)
   - Considerar marcapasso transcutâneo
   - Dopamina 5-20 mcg/kg/min ou epinefrina 2-10 mcg/min

   TAQUICARDIA COM INSTABILIDADE:
   - Cardioversão sincronizada conforme protocolo ACLS
   - Avaliar necessidade de sedação/anestesia

FOLLOW-UP:
- Bradicardia sintomática: retorno em 2-4 semanas após ajustes
- Taquicardia não explicada: retorno em 4-6 semanas com exames
- FC normal mas HRV reduzida: retorno em 3-6 meses para reavaliação
- Atletas com bradicardia fisiológica: retorno anual de rotina"""

def enrich_item():
    """Enrich ECG - Frequência Cardíaca item"""

    try:
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        # Update score_items with clinical content
        update_query = """
        UPDATE score_items
        SET
            clinical_relevance = %s,
            patient_explanation = %s,
            conduct = %s,
            last_review = %s,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = %s
        RETURNING id, name;
        """

        cur.execute(update_query, (
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            datetime.now().date(),
            ITEM_ID
        ))

        result = cur.fetchone()

        if result:
            print(f"✅ Successfully enriched item: {result[1]} (ID: {result[0]})")
        else:
            print(f"❌ Item not found: {ITEM_ID}")
            return

        # Commit changes
        conn.commit()

        # Verify update
        cur.execute("""
            SELECT
                id,
                name,
                clinical_relevance IS NOT NULL as has_clinical,
                patient_explanation IS NOT NULL as has_explanation,
                conduct IS NOT NULL as has_conduct,
                last_review
            FROM score_items
            WHERE id = %s;
        """, (ITEM_ID,))

        verification = cur.fetchone()
        print("\n📊 Verification:")
        print(f"   ID: {verification[0]}")
        print(f"   Name: {verification[1]}")
        print(f"   Clinical Relevance: {'✅' if verification[2] else '❌'}")
        print(f"   Patient Explanation: {'✅' if verification[3] else '❌'}")
        print(f"   Conduct: {'✅' if verification[4] else '❌'}")
        print(f"   Last Review: {verification[5]}")

        cur.close()
        conn.close()

        print("\n✨ Enrichment completed successfully!")

    except Exception as e:
        print(f"❌ Error: {e}")
        try:
            if 'conn' in locals() and conn:
                conn.rollback()
        except:
            pass
        raise

if __name__ == "__main__":
    print("🔬 Starting enrichment for ECG - Frequência Cardíaca")
    print(f"📝 Item ID: {ITEM_ID}\n")
    enrich_item()
