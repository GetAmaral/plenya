#!/usr/bin/env python3
"""
Enriquecimento clínico: Ecodopplercardiograma - FEVE Homens
ID: 6e542cb0-6982-42e8-93dc-40f139652223
"""

import psycopg2
import os
from datetime import datetime

# Dados do enriquecimento
SCORE_ITEM_ID = "6e542cb0-6982-42e8-93dc-40f139652223"

# Conteúdo clínico em PT-BR
CLINICAL_RELEVANCE = """A Fração de Ejeção do Ventrículo Esquerdo (FEVE) é o parâmetro mais importante para avaliar a função sistólica cardíaca e estratificar o risco cardiovascular. Valores normais em homens variam entre 52-72%, com estudos recentes indicando que o limiar inferior normal varia por etnia: europeus 50%, asiáticos orientais 56%, e sul-asiáticos 52%.

**Classificação da Insuficiência Cardíaca por FEVE:**
- ICFEr (FEVE ≤40%): Insuficiência cardíaca com fração de ejeção reduzida - disfunção sistólica significativa
- ICFEfm (FEVE 41-49%): Fração de ejeção na faixa média - categoria intermediária de risco
- ICFEp (FEVE ≥50%): Fração de ejeção preservada - disfunção diastólica predominante
- ICFEsn (FEVE ≥65%): Fração de ejeção supranormal - pode indicar cardiomiopatia hipertrófica

**Prognóstico:** Estudos demonstram que o risco de insuficiência cardíaca aumenta progressivamente com FEVEs mais baixas, iniciando com valores nos baixos 60%. Uma queda de ≥5% na FEVE sinaliza alerta de perigo, associando-se a aumento de 82% no risco de mortalidade em 4 anos. Apenas 24% dos pacientes permanecem na categoria de ICFEfm ao longo do tempo, evidenciando a natureza dinâmica desta medida.

**Estratificação de Risco:** O escore MAGGIC (13 variáveis incluindo FEVE) prediz mortalidade em 1 e 3 anos (c-statistic 0.73) e é aplicável em todos os subtipos de IC. A avaliação combinada de congestão clínica, ecocardiográfica (VCI≥21mm; E/e'≥15) e ultrassom pulmonar melhora significativamente a estratificação de risco, independente da FEVE.

**Tratamento Baseado em FEVE:** As diretrizes ACC 2024 recomendam inibidores SGLT2 durante toda hospitalização independente da FEVE, com maior ênfase na iniciação das outras terapias-pilar (betabloqueadores, inibidores do SRAA) para ICFEr após estabilização. A redução maior em hospitalização por IC e mortalidade, bem como remissão para FEVE melhorada, tem sido atribuída a betabloqueadores e inibidores do SRAA."""

PATIENT_EXPLANATION = """O ecodopplercardiograma com FEVE avalia a força de bombeamento do seu coração, especificamente do ventrículo esquerdo - a principal câmara que bombeia sangue oxigenado para todo o corpo.

**O que é FEVE?**
A Fração de Ejeção (FEVE) mede a porcentagem de sangue que seu coração ejeta a cada batimento. É como medir a eficiência de uma bomba: quanto sangue ela consegue impulsionar em relação ao volume total dentro dela.

**Valores para Homens:**
- **Normal**: 52% a 72% - Seu coração está bombeando adequadamente
- **Levemente reduzida**: 41% a 51% - Função um pouco diminuída, requer monitoramento
- **Reduzida**: Abaixo de 40% - Indica disfunção sistólica significativa (insuficiência cardíaca)
- **Muito alta**: Acima de 72% - Pode indicar espessamento do músculo cardíaco

**Por que é importante?**
A FEVE é o principal indicador para:
- Diagnosticar insuficiência cardíaca e determinar seu tipo
- Avaliar risco de eventos cardiovasculares futuros
- Guiar decisões sobre medicamentos e tratamentos
- Monitorar resposta ao tratamento ao longo do tempo

**Sinais de Alerta:**
Uma queda de 5% ou mais na FEVE entre exames é um sinal importante que requer atenção médica, pois pode indicar piora da função cardíaca e aumentar significativamente o risco cardiovascular.

**Tratamento:**
Dependendo do valor da sua FEVE, seu médico pode recomendar medicamentos específicos como betabloqueadores, inibidores da enzima conversora (IECA) ou inibidores SGLT2, que comprovadamente melhoram a função cardíaca e reduzem hospitalizações."""

CONDUCT = """**REFERÊNCIAS CIENTÍFICAS:**

1. McMurray JJ, et al. "Dapagliflozin in Patients with Heart Failure and Reduced Ejection Fraction." N Engl J Med. 2019;381(21):1995-2008. doi:10.1056/NEJMoa1911303

2. Packer M, et al. "Cardiovascular and Renal Outcomes with Empagliflozin in Heart Failure." N Engl J Med. 2020;383(15):1413-1424. doi:10.1056/NEJMoa2022190

3. McDonagh TA, et al. "2021 ESC Guidelines for the diagnosis and treatment of acute and chronic heart failure." Eur Heart J. 2021;42(36):3599-3726. doi:10.1093/eurheartj/ehab368

4. Heidenreich PA, et al. "2022 AHA/ACC/HFSA Guideline for the Management of Heart Failure." J Am Coll Cardiol. 2022;79(17):e263-e421. doi:10.1016/j.jacc.2021.12.012

5. "2024 ACC Expert Consensus Decision Pathway on Clinical Assessment, Management, and Trajectory of Patients Hospitalized With Heart Failure." JACC. 2024. doi:10.1016/j.jacc.2024.06.002

6. "What Is a Normal Left Ventricular Ejection Fraction?" Circulation. 2023. doi:10.1161/CIRCULATIONAHA.123.065791

7. Pocock SJ, et al. "Predicting survival in heart failure: a risk score based on 39 372 patients from 30 studies." Eur Heart J. 2013;34(19):1404-13. doi:10.1093/eurheartj/ehs337 (MAGGIC score)

8. Anker SD, et al. "Empagliflozin in Heart Failure with a Preserved Ejection Fraction." N Engl J Med. 2021;385(16):1451-1461. doi:10.1056/NEJMoa2107038

9. Solomon SD, et al. "Dapagliflozin in Heart Failure with Mildly Reduced or Preserved Ejection Fraction." N Engl J Med. 2022;387(12):1089-1098. doi:10.1056/NEJMoa2206286

10. "Risk stratification and survival prediction in heart failure: from grades to scores." Front Cardiovasc Med. 2025. doi:10.3389/fcvm.2025.1676441

---

**1. Interpretação do Resultado:**
- FEVE >52% em homens: Função sistólica normal
- FEVE 41-51%: Disfunção sistólica leve (ICFEfm) - investigar causas e iniciar monitoramento
- FEVE ≤40%: Disfunção sistólica moderada-grave (ICFEr) - indicação formal de tratamento para IC
- FEVE >72%: Investigar cardiomiopatia hipertrófica ou outras causas de hipercontratilidade

**2. Avaliação Complementar Obrigatória:**
- História clínica detalhada: sintomas de IC (dispneia, edema, fadiga, ortopneia)
- Exame físico: ausculta cardíaca, congestão pulmonar, edema periférico, turgência jugular
- ECG: avaliar ritmo, hipertrofia, isquemia
- Parâmetros ecocardiográficos adicionais: E/e' (função diastólica), VCI (pressão atrial direita), strain longitudinal global
- BNP ou NT-proBNP: marcadores de estresse miocárdico
- Função renal e eletrólitos: creatinina, ureia, sódio, potássio
- Investigar etiologia: isquemia (coronariografia se indicado), valvopatias, hipertensão, álcool

**3. Estratificação de Risco:**
- Aplicar escore MAGGIC para predição de mortalidade (www.heartfailurerisk.org)
- Avaliar sinais de congestão: clínicos + ecocardiográficos (VCI≥21mm; E/e'≥15) + ultrassom pulmonar
- Considerar escore EFFECT modificado em pacientes hospitalizados

**4. Conduta Terapêutica por Categoria:**

**ICFEr (FEVE ≤40%) - Terapia Quadrupla:**
- **Inibidor SGLT2** (dapagliflozina 10mg ou empagliflozina 10mg) - iniciar precocemente
- **IECA ou BRA** (enalapril, ramipril, losartana) - titular até dose-alvo
- **Betabloqueador** (carvedilol, metoprolol succinato, bisoprolol) - titular progressivamente
- **Antagonista mineralocorticoide** (espironolactona 25-50mg) se K+ <5,0 e Cr <2,5
- Considerar **INRA** (sacubitril/valsartana) se sintomas persistentes após otimização

**ICFEfm (FEVE 41-49%):**
- Inibidor SGLT2 (evidência robusta em prevenção de progressão)
- Considerar IECA/BRA e betabloqueador
- Controle rigoroso de fatores de risco: HAS, DM, dislipidemia
- Monitoramento trimestral da FEVE

**ICFEp (FEVE ≥50%):**
- Inibidor SGLT2 (benefício comprovado mesmo com FE preservada)
- Controle de comorbidades: HAS, DM, fibrilação atrial, obesidade
- Diuréticos se congestão

**5. Monitoramento:**
- Repetir ecocardiograma em 3-6 meses após início/otimização do tratamento
- **Queda ≥5% na FEVE**: reajustar tratamento urgentemente, investigar causas (isquemia, arritmias, não aderência)
- Acompanhamento ambulatorial mensal inicialmente, depois trimestral se estável

**6. Indicações de Referência para Cardiologista:**
- Qualquer FEVE <50%
- Queda ≥5% entre exames consecutivos
- FEVE <35% com QRS alargado (avaliar ressincronização cardíaca)
- FEVE <30% refratária (avaliar CDI - cardioversor-desfibrilador implantável)
- Sintomas NYHA III-IV apesar de tratamento otimizado

**7. Educação do Paciente:**
- Restrição de sódio (<2g/dia se ICFEr)
- Controle de peso diário (alerta se ganho >2kg em 3 dias)
- Atividade física supervisionada (reabilitação cardíaca)
- Vacinação (influenza, pneumococo, COVID-19)
- Evitar AINEs, álcool em excesso, tabagismo"""


def connect_db():
    """Conecta ao banco de dados PostgreSQL"""
    return psycopg2.connect(
        host=os.getenv('DB_HOST', 'localhost'),
        port=os.getenv('DB_PORT', '5432'),
        database=os.getenv('DB_NAME', 'plenya_db'),
        user=os.getenv('DB_USER', 'plenya_user'),
        password=os.getenv('DB_PASSWORD', 'plenya_dev_password')
    )

def update_score_item():
    """Atualiza o score_item com conteúdo clínico"""
    conn = connect_db()
    cur = conn.cursor()

    try:
        # Verificar se o item existe
        cur.execute("""
            SELECT si.id, si.name, sg.name as subgroup, g.name as group_name
            FROM score_items si
            LEFT JOIN score_subgroups sg ON si.subgroup_id = sg.id
            LEFT JOIN score_groups g ON sg.group_id = g.id
            WHERE si.id = %s
        """, (SCORE_ITEM_ID,))

        item = cur.fetchone()
        if not item:
            print(f"❌ Erro: Item {SCORE_ITEM_ID} não encontrado no banco")
            return False

        print(f"✅ Item encontrado: {item[1]} (Subgrupo: {item[2]}, Grupo: {item[3]})")

        # Atualizar campos clínicos
        cur.execute("""
            UPDATE score_items
            SET
                clinical_relevance = %s,
                patient_explanation = %s,
                conduct = %s,
                last_review = CURRENT_TIMESTAMP,
                updated_at = NOW()
            WHERE id = %s
        """, (
            CLINICAL_RELEVANCE,
            PATIENT_EXPLANATION,
            CONDUCT,
            SCORE_ITEM_ID
        ))

        conn.commit()
        print(f"✅ Score item atualizado com sucesso!")
        print(f"   - clinical_relevance: {len(CLINICAL_RELEVANCE)} caracteres")
        print(f"   - patient_explanation: {len(PATIENT_EXPLANATION)} caracteres")
        print(f"   - conduct: {len(CONDUCT)} caracteres")
        print(f"   - last_review: {datetime.now()}")

        return True

    except Exception as e:
        conn.rollback()
        print(f"❌ Erro ao atualizar score item: {e}")
        return False
    finally:
        cur.close()
        conn.close()

def main():
    print("="*80)
    print("ENRIQUECIMENTO CLÍNICO: Ecodopplercardiograma - FEVE Homens")
    print("="*80)
    print()

    success = update_score_item()

    print()
    print("="*80)
    if success:
        print("✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO")
        print("="*80)
        print()
        print("📊 RESUMO DO CONTEÚDO ADICIONADO:")
        print()
        print("1. RELEVÂNCIA CLÍNICA:")
        print("   - Valores normais por etnia (europeus 50%, asiáticos 56%, sul-asiáticos 52%)")
        print("   - Classificação IC: ICFEr (≤40%), ICFEfm (41-49%), ICFEp (≥50%), ICFEsn (≥65%)")
        print("   - Prognóstico: queda ≥5% = +82% risco mortalidade em 4 anos")
        print("   - Escore MAGGIC para estratificação de risco")
        print()
        print("2. EXPLICAÇÃO PARA PACIENTE:")
        print("   - Linguagem acessível sobre função de bombeamento cardíaco")
        print("   - Valores normais e interpretação para homens")
        print("   - Sinais de alerta e importância do monitoramento")
        print()
        print("3. CONDUTA CLÍNICA:")
        print("   - Protocolo completo por categoria de FEVE")
        print("   - Terapia quadrupla para ICFEr (SGLT2, IECA/BRA, BB, ARM)")
        print("   - Diretrizes ACC 2024")
        print("   - Critérios de referência para cardiologista")
        print()
        print("4. REFERÊNCIAS CIENTÍFICAS:")
        print("   - 10 artigos de alto impacto")
        print("   - Diretrizes ESC 2021, AHA/ACC/HFSA 2022, ACC 2024")
        print("   - Estudos DAPA-HF, EMPEROR-Reduced, DELIVER")
        print()
        print("🔍 PRÓXIMOS PASSOS:")
        print("   - Buscar artigos em PDF para upload no sistema")
        print("   - Criar relações article_score_items quando disponível")
        print()
    else:
        print("❌ ERRO NO ENRIQUECIMENTO")
        print("="*80)

    return 0 if success else 1

if __name__ == "__main__":
    exit(main())
