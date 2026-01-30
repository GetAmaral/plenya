#!/usr/bin/env python3
"""
Enriquecimento do item ECG - Sokolow-Lyon
ID: 631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb

Baseado em:
- LITFL ECG Library
- Meta-análise PubMed (RR 1.62 para MACE)
- AHA Hypertension Journal
- Estudos de predição de mortalidade cardiovascular
"""

import psycopg2
from datetime import datetime

# Conexão com banco
conn = psycopg2.connect(
    host="localhost",
    port=5432,
    database="plenya_db",
    user="plenya_user",
    password="senha_segura"
)
cur = conn.cursor()

# ID do item
item_id = '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'

# Conteúdo clínico enriquecido
clinical_relevance = """O critério de Sokolow-Lyon é um dos métodos eletrocardiográficos mais utilizados para diagnóstico de hipertrofia ventricular esquerda (HVE), calculado pela soma da onda S em V1 com a maior onda R em V5 ou V6.

**Interpretação:**
• Valor normal: ≤35 mm (≤3.5 mV)
• HVE provável: >35 mm (>3.5 mV)
• HVE significativa: >40 mm (>4.0 mV)

**Características diagnósticas:**
• Sensibilidade: 20-25% (baixa)
• Especificidade: >85% (alta)
• Valor preditivo: Alta especificidade torna o teste confirmatório quando positivo

**Significado clínico:**
Meta-análise populacional demonstrou que HVE pelo Sokolow-Lyon está associada a:
• Risco relativo de 1.62 (IC95% 1.40-1.89) para eventos cardiovasculares maiores (MACE)
• Risco relativo de 1.47 (IC95% 1.10-1.97) para mortalidade por todas as causas
• Risco relativo de 1.38 (IC95% 1.19-1.60) para mortalidade cardiovascular
• Cada aumento de 0.1 mV prediz aumento de 1.4-3.9% no risco de AVC, doença coronariana e morte cardiovascular

**Limitações importantes:**
• IMC elevado reduz amplitude do QRS (falso negativo em obesos)
• Pacientes jovens e magros podem ter valores elevados sem HVE verdadeira
• Menor sensibilidade que Cornell voltage em algumas populações
• Não detecta HVE concêntrica com mesma eficácia que HVE excêntrica

**Quando suspeitar de HVE:**
• Hipertensão arterial não controlada
• Estenose aórtica
• Miocardiopatia hipertrófica
• Atletas de alta performance (fisiológica)
• Insuficiência aórtica crônica"""

patient_explanation = """O **Critério de Sokolow-Lyon** é uma medida no eletrocardiograma (ECG) que avalia se o músculo do seu coração está aumentado de tamanho (hipertrofia ventricular esquerda).

**Como funciona:**
O médico soma a profundidade de uma onda (onda S) em uma parte do ECG com a altura de outra onda (onda R) em outra parte. Se essa soma for maior que 35 mm, pode indicar que o coração está trabalhando mais do que deveria.

**O que significa:**
• **Normal (≤35 mm):** O músculo do coração tem espessura adequada
• **Aumentado (>35 mm):** O coração pode estar "bombando" com mais força do que o normal, aumentando o músculo

**Por que o coração aumenta?**
1. **Pressão alta:** O coração precisa fazer mais força para bombear sangue, então o músculo cresce
2. **Problemas nas válvulas:** Válvula aórtica estreita ou com vazamento
3. **Atletas:** Pode ser normal em quem pratica esporte intenso
4. **Outras doenças:** Algumas condições cardíacas específicas

**Por que isso é importante?**
Quando o critério está alterado (positivo), estudos mostram que há:
• 62% mais chance de ter problemas cardíacos graves
• 47% mais risco de morte por qualquer causa
• 38% mais risco de morte por doenças do coração

**Importante saber:**
• Se você tem sobrepeso, o exame pode não detectar o problema (falso negativo)
• Se você é jovem e magro, pode aparecer alterado mesmo sem problema (falso positivo)
• Um resultado positivo é bem confiável (85% de certeza)
• Um resultado normal não descarta completamente o problema

**O que fazer se estiver alterado:**
Seu médico pode pedir outros exames (como ecocardiograma) para confirmar e avaliar melhor o coração."""

conduct = """**1. CONFIRMAÇÃO DIAGNÓSTICA:**

Se Sokolow-Lyon >35 mm:
• Solicitar ecocardiograma transtorácico para confirmar HVE anatômica
• Avaliar massa ventricular esquerda indexada (MVEI):
  - Homens: normal <115 g/m²
  - Mulheres: normal <95 g/m²
• Classificar padrão geométrico (concêntrico vs excêntrico)
• Avaliar função sistólica (FEVE) e diastólica (relação E/A, E/e')

Considerar outros critérios ECG de HVE:
• Cornell voltage (mais sensível em mulheres)
• Cornell product
• Romhilt-Estes score (mais específico)

**2. INVESTIGAÇÃO ETIOLÓGICA:**

A. Avaliação cardiovascular:
• MAPA 24h ou medidas seriadas de PA (meta <130/80 mmHg)
• Pesquisa de lesão de órgão-alvo:
  - Microalbuminúria / creatinina sérica / TFG
  - Fundo de olho (retinopatia hipertensiva)
  - ECG basal (arritmias, isquemia silenciosa)

B. Avaliação valvular (se ecocardiograma anormal):
• Ecocardiograma com Doppler para gradientes
• TC ou RNM cardíaca se estenose/insuficiência valvular

C. Descartar causas secundárias de hipertensão (se hipertensão refratária):
• Aldosterona/renina (hiperaldosteronismo primário)
• Metanefrinas (feocromocitoma)
• Ultrassonografia de artérias renais (estenose renal)

**3. ESTRATIFICAÇÃO DE RISCO:**

Calcular escore de risco cardiovascular (Framingham ou DAC):
• HVE pelo Sokolow-Lyon aumenta risco em 62% (considerar como alto risco)
• Pesquisar fatores modificáveis:
  - Dislipidemia (LDL, HDL, triglicerídeos)
  - Diabetes mellitus (HbA1c, glicemia)
  - Tabagismo
  - Obesidade (IMC, circunferência abdominal)

**4. TRATAMENTO:**

A. Medidas não-farmacológicas (fundamentais):
• Restrição de sódio (<2g/dia ou 5g sal/dia)
• Dieta DASH (rica em frutas, vegetais, laticínios desnatados)
• Perda de peso se IMC >25 kg/m² (cada 5 kg reduz PA em 5 mmHg)
• Exercício aeróbico moderado 150 min/semana
• Cessar tabagismo
• Limitar álcool (≤2 doses/dia homens, ≤1 mulheres)

B. Tratamento farmacológico de HAS:
Primeira linha (regressão de HVE comprovada):
• **IECA** (enalapril, ramipril) ou **BRA** (losartana, valsartana)
  - Bloqueio do SRAA → reduz hipertrofia em 6-12 meses
  - Meta: PA <130/80 mmHg

Associações se necessário:
• Bloqueadores de canal de cálcio (amlodipina)
• Diuréticos tiazídicos (hidroclorotiazida)
• Beta-bloqueadores se indicação adicional (angina, IC, arritmia)

C. Manejo de comorbidades:
• Estatina se LDL elevado (meta <70 mg/dL se alto risco)
• Antiagregação se DAC estabelecida
• Controle glicêmico rigoroso se diabetes (HbA1c <7%)

**5. MONITORAMENTO:**

Reavaliação seriada:
• ECG a cada 12 meses para acompanhar Sokolow-Lyon
• Ecocardiograma a cada 2-3 anos (ou antes se sintomas)
• Regressão da HVE é marcador de redução de risco cardiovascular

Critérios para regressão:
• Redução de MVEI >10% no ecocardiograma
• Normalização do Sokolow-Lyon (<35 mm)

**6. ENCAMINHAMENTO:**

Considerar cardiologista se:
• Sokolow-Lyon >50 mm (HVE importante)
• HVE + sintomas (dispneia, angina, síncope)
• Hipertensão resistente (≥3 fármacos incluindo diurético)
• Suspeita de miocardiopatia hipertrófica (história familiar, genética)
• Valvulopatia moderada/grave no ecocardiograma

**7. SEGUIMENTO ESPECIAL:**

Atletas com Sokolow-Lyon elevado:
• Diferenciar HVE fisiológica vs patológica
• Critérios adicionais: padrão strain, disfunção diastólica ausentes
• RNM cardíaca se dúvida diagnóstica
• Clearance para esporte competitivo após avaliação completa"""

# Atualizar o item no banco
print(f"Atualizando item {item_id}...")

update_query = """
UPDATE score_items
SET
    clinical_relevance = %s,
    patient_explanation = %s,
    conduct = %s,
    last_review = %s,
    updated_at = NOW()
WHERE id = %s
"""

cur.execute(update_query, (
    clinical_relevance,
    patient_explanation,
    conduct,
    datetime.now(),
    item_id
))

# Commit
conn.commit()

# Verificar atualização
cur.execute("""
    SELECT name,
           LENGTH(clinical_relevance) as len_clinical,
           LENGTH(patient_explanation) as len_patient,
           LENGTH(conduct) as len_conduct,
           last_review
    FROM score_items
    WHERE id = %s
""", (item_id,))

result = cur.fetchone()
print(f"\n✅ Item atualizado com sucesso!")
print(f"Nome: {result[0]}")
print(f"Tamanho clinical_relevance: {result[1]} caracteres")
print(f"Tamanho patient_explanation: {result[2]} caracteres")
print(f"Tamanho conduct: {result[3]} caracteres")
print(f"Última revisão: {result[4]}")

# Fechar conexão
cur.close()
conn.close()

print("\n" + "="*80)
print("ENRIQUECIMENTO CONCLUÍDO - ECG Sokolow-Lyon")
print("="*80)
print("\nReferências científicas utilizadas:")
print("• LITFL - Left Ventricular Hypertrophy ECG Library")
print("• Meta-análise: Predictive value of ECG LVH (RR 1.62 for MACE)")
print("• AHA Hypertension: Gender-Specific Partition Values")
print("• PubMed: LVH by Sokolow-Lyon predicts mortality")
print("• ECGwaves: Clinical characteristics and implications")
print("="*80)
