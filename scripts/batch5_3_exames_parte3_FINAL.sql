-- =====================================================
-- BATCH 5.3: EXAMES PARTE 3 - FINAL (itens 36-44)
-- Continuação do enriquecimento
-- Executar APÓS batch5_3_exames_parte3_enrichment.sql
-- =====================================================

BEGIN;

-- 36. Ecodopplercardiograma - FEVE Mulheres
UPDATE score_items
SET
    interpretation = 'FEVE em mulheres tem mesmos valores de referência que homens (≥55% normal), mas mulheres têm menor volume diastólico final (menor tamanho cardíaco). Mulheres com IC têm maior prevalência de ICFEp (fração preservada ≥50%) que homens, especialmente pós-menopausa (hipertensão, rigidez arterial, disfunção diastólica). Miocardiopatia de Takotsubo (broken heart syndrome) é 90% feminina, causando disfunção sistólica transitória (FEVE 20-40%) após estresse emocional, com recuperação completa em 4-8 semanas. Cardiotoxicidade por quimioterapia (trastuzumab, antraciclinas) também é mais comum em mulheres.',

    clinical_relevance = 'Medicina funcional valoriza FEVE em mulheres para: diagnóstico de miocardiopatia periparto (FEVE <45% no último mês gestação ou 5 meses pós-parto); monitoramento de cardiotoxicidade por quimioterapia (redução FEVE >10% ou valor absoluto <50% define cardiotoxicidade); investigação de dispneia inexplicada (mulheres têm mais ICFEp que homens - FEVE normal mas disfunção diastólica); estratificação de risco em cardiopatia isquêmica (mulheres com FEVE <35% também se beneficiam de CDI, mas são sub-representadas em trials).',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 55.0,
    optimal_value_female_max = 70.0,
    optimal_value_description = 'FEVE ótima em mulheres: 55-70% (padrão MFI, idêntico a homens). Valores <55% indicam disfunção sistólica. FEVE 41-54%: disfunção leve. FEVE <40%: ICFEr (tratamento GDMT). FEVE >70%: hiperdinâmica (investigar hipertireoidismo, anemia, sepse). Ideal: 60-65%. Mulheres com sintomas IC e FEVE ≥50% devem ter avaliação diastólica detalhada (relação E/e', velocidade e\' septal/lateral, volume AE indexado, velocidade onda A tricúspide) para diagnóstico de ICFEp.',

    health_recommendations = 'FEVE 41-54% em mulheres:
• Investigar etiologia: cineangiocoronariografia (DAC menos obstrutiva em mulheres, mais doença microvascular), RM cardíaca
• Tratamento IC (mesmo protocolo homens):
  - IECA/ARNI + beta-bloqueador + espironolactona + SGLT2-i
  - Sacubitril-valsartana tem maior benefício em mulheres (PARADIGM-HF subanalysis)

FEVE < 40% (ICFEr):
• Terapia quádrupla GDMT (ver protocolo FEVE homens)
• CDI se FEVE <35% + NYHA II-III (mesmo benefício que homens)
• TRC se BRE com QRS ≥130 ms

Miocardiopatia periparto (FEVE <45%):
• Bromocriptina 2.5 mg 2x/dia por 2 semanas, depois 2.5 mg/dia por 6 semanas (inibe prolactina, reduz estresse oxidativo miocárdico)
• Contraindicar gestações futuras se FEVE não normalizar
• FEVE recupera em 50% dos casos (melhor prognóstico se FEVE inicial >30%, VE não-dilatado)

Cardiotoxicidade quimioterapia:
• Prevenção: dexrazoxano (quelante ferro, protege contra antraciclinas)
• Carvedilol 12.5 mg 2x/dia + IECA durante quimioterapia (reduz cardiotoxicidade)
• Monitorar FEVE antes, durante (a cada 3 meses) e após quimioterapia
• Redução FEVE >10% ou absoluto <50%: suspender temporariamente, iniciar IC terapia

Takotsubo (broken heart syndrome):
• Disfunção apical transitória (balão apical) pós-estresse emocional/físico
• Suporte: beta-bloqueador, IECA, anticoagulação se trombo VE
• FEVE normaliza em 4-8 semanas em 95%
• Recorrência 5% (evitar estressores)

ICFEp (FEVE ≥50% + sintomas):
• Mais comum em mulheres pós-menopausa
• Tratamento:
  - SGLT2-i dapagliflozina 10 mg/dia (DELIVER trial - reduz hospitaliz. em ICFEp)
  - Diurético (furosemida) para controle congestão
  - Controle rigoroso PA (<130/80 mmHg)
• Espironolactona controversa em ICFEp (TOPCAT trial neutro)

Suplementação:
• CoQ10 ubiquinol 300 mg/dia (especialmente se estatinas)
• Ômega-3 2-3g/dia
• Magnésio 400 mg/dia
• Vitamina D 4000 UI/dia

Exercício:
• Reabilitação cardíaca supervisionada
• Aeróbico 30 min 5x/semana

Reavaliar ecocardiograma em 6-12 meses
Meta: FEVE >50% ou melhora ≥10%',

    last_review_date = CURRENT_DATE
WHERE id = 'c8795e89-b10a-4d51-b940-463ab5e89c3e';

-- 37. Ecodopplercardiograma - GLS (Global Longitudinal Strain)
UPDATE score_items
SET
    interpretation = 'GLS (strain longitudinal global) mede deformação miocárdica longitudinal por speckle-tracking ecocardiográfico, detectando disfunção sistólica subclínica antes de redução de FEVE. GLS normal: -18% a -22% (valor negativo indica encurtamento/deformação). GLS >-18% (menos negativo, ex: -15%) indica disfunção mesmo com FEVE preservada. GLS é preditor mais sensível de cardiotoxicidade (quimioterapia), doença coronariana multiarterial, insuficiência cardíaca incipiente e mortalidade cardiovascular que FEVE isolada. GLS também avalia viabilidade miocárdica pós-IAM.',

    clinical_relevance = 'Medicina funcional valoriza GLS em: detecção precoce de cardiotoxicidade (redução GLS >15% do basal indica cardiotoxicidade antes de FEVE cair); estratificação de risco em valvopatias (estenose aórtica com GLS >-14% tem pior prognóstico, mesmo com FEVE normal); monitoramento de miocardiopatia hipertrófica (GLS diferencia hipertrofia atlética de patológica); avaliação de IC com FEVE preservada (GLS alterado indica disfunção longitudinal apesar de FEVE normal). GLS <-20% é meta ideal para função contrátil ótima.',

    optimal_value_male_min = -22.0,
    optimal_value_male_max = -18.0,
    optimal_value_female_min = -22.0,
    optimal_value_female_max = -18.0,
    optimal_value_description = 'GLS ótimo: -18% a -22% (mais negativo indica melhor função). Valores >-18% (ex: -15%, -12%) indicam disfunção longitudinal subclínica, mesmo com FEVE preservada. GLS >-16% correlaciona-se com risco aumentado de IC e mortalidade. GLS <-22% (ex: -24%) é excelente. Ideal: -20% a -22%, refletindo deformação miocárdica longitudinal ótima e reserva contrátil preservada. GLS é superior a FEVE para detecção precoce de disfunção.',

    health_recommendations = 'GLS entre -18% e -16% (disfunção leve):
• Investigar causas: DAC (cineangiocoronariografia ou TC coronárias), hipertensão não-controlada, cardiotoxicidade incipiente
• Otimizar controle fatores risco: PA <130/80 mmHg, LDL <70 mg/dL, HbA1c <7%
• IECA ou BRA mesmo com FEVE normal (proteção miocárdica)
• Repetir ecocardiograma com GLS em 6-12 meses

GLS > -16% (disfunção moderada-severa):
• Tratamento IC mesmo se FEVE >50%:
  - IECA/BRA ou sacubitril-valsartana
  - Beta-bloqueador se DAC ou taquicardia
  - SGLT2-i (dapagliflozina 10 mg/dia)
• Investigação abrangente: RM cardíaca (fibrose, realce tardio), BNP

Cardiotoxicidade (GLS piora >15% do basal):
• Suspender temporariamente quimioterapia
• Iniciar carvedilol 12.5 mg 2x/dia + enalapril 10 mg 2x/dia
• CoQ10 300-600 mg/dia
• Reavaliar GLS em 4-8 semanas
• Se GLS não recuperar: trocar regime quimioterápico

Valvopatias (estenose aórtica severa com GLS >-14%):
• Pior prognóstico, mesmo com FEVE normal
• Considerar troca valvar precoce (cirúrgica ou TAVI)

Miocardiopatia hipertrófica:
• GLS >-16% indica maior risco de IC e arritmias
• Considerar CDI se GLS severamente reduzido + fatores risco (história familiar MS, síncope)

Otimização GLS:
• Exercício aeróbico regular 150 min/semana (melhora GLS em 1-2%)
• Controle rigoroso PA (redução PA melhora GLS)
• Estatinas (atorvastatina 40 mg/dia - efeito pleiotrópico anti-remodelamento)
• Ômega-3 4g/dia (reduz remodelamento adverso)

Suplementação:
• CoQ10 ubiquinol 300 mg/dia (melhora função mitocondrial, contratilidade)
• L-carnitina 2g/dia (metabolismo energético miocárdico)
• Magnésio 400 mg/dia (cofator ATP)
• Vitamina D 4000 UI/dia (modula remodelamento)

GLS < -22% (excelente):
• Manter estilo de vida saudável
• Monitoramento de rotina

Reavaliar GLS em 6-12 meses se alterado
Meta: melhorar GLS >10% relativo ou atingir <-18%',

    last_review_date = CURRENT_DATE
WHERE id = 'fb5c484c-bb3c-4017-b4da-866d96d9cb20';

-- 38. QUICKI (Quantitative Insulin Sensitivity Check Index)
UPDATE score_items
SET
    interpretation = 'QUICKI é índice de sensibilidade à insulina calculado por: 1 / [log(insulina jejum μU/mL) + log(glicemia jejum mg/dL)]. QUICKI >0.357 indica sensibilidade normal; 0.330-0.357 resistência moderada; <0.330 resistência severa. QUICKI correlaciona-se fortemente com clamp euglicêmico-hiperinsulinêmico (padrão-ouro), mas é mais prático. Resistência à insulina precede diabetes tipo 2 em 10-15 anos, está presente em síndrome metabólica, SOP, EHNA e é fator de risco cardiovascular independente.',

    clinical_relevance = 'Medicina funcional valoriza QUICKI em: diagnóstico de resistência insulínica antes de hiperglicemia (glicemia jejum pode ser normal com QUICKI <0.330); monitoramento de resposta a intervenções (dieta, exercício, metformina); estratificação de risco cardiovascular (QUICKI <0.330 aumenta risco IAM 2-3x); avaliação de SOP (QUICKI <0.330 em 70% das mulheres com SOP); predição de progressão para diabetes (QUICKI <0.330 + HbA1c 5.7-6.4% indica pré-diabetes com alto risco). QUICKI >0.380 é meta ideal para sensibilidade insulínica ótima.',

    optimal_value_male_min = 0.357,
    optimal_value_male_max = 0.450,
    optimal_value_female_min = 0.357,
    optimal_value_female_max = 0.450,
    optimal_value_description = 'QUICKI ótimo: >0.357 (sensibilidade insulínica normal). Valores 0.330-0.357 indicam resistência moderada. QUICKI <0.330 define resistência severa, com risco elevado de progressão para diabetes tipo 2 (5-10% ao ano). QUICKI <0.300 indica resistência grave, frequentemente acompanhada de síndrome metabólica completa. Ideal: 0.380-0.420, refletindo sensibilidade insulínica ótima, minimizando hiperinsulinemia compensatória e risco cardiometabólico.',

    health_recommendations = 'QUICKI 0.330-0.357 (resistência moderada):
• Perda de peso 5-10% (melhora QUICKI ~0.020-0.040)
• Dieta low-carb <100-150g/dia carboidratos líquidos
• Jejum intermitente 16:8 (reduz insulina jejum 20-30%)
• Exercício: aeróbico 150 min/semana + resistido 2-3x/semana
• Fibras 30-40g/dia (retardam absorção glicose)

QUICKI < 0.330 (resistência severa):
• Metformina 500 mg 3x/dia ou 850 mg 2x/dia (melhora QUICKI ~0.030-0.050)
• Dieta cetogênica <50g carboidratos/dia (maior impacto em resistência)
• Berberina 500 mg 3x/dia (efeito similar metformina)
• Inositol (mio-inositol 2g + d-chiro-inositol 50 mg) 2x/dia (especialmente mulheres/SOP)

QUICKI < 0.300 (resistência grave):
• Metformina dose máxima (1000 mg 3x/dia ou 1700 mg 2x/dia XR)
• Adicionar SGLT2-i: empagliflozina 10-25 mg/dia ou dapagliflozina 10 mg/dia (reduz glicose/insulina)
• Ou agonista GLP-1: liraglutida 1.2-3.0 mg SC/dia ou semaglutida 0.5-1.0 mg SC/semana
• Considerar cirurgia bariátrica se IMC >35 (resolve resistência em 80%)

Suplementação:
• Cromo picolinato 200-1000 mcg/dia (aumenta sensibilidade insulina)
• Ácido alfa-lipóico 600-1200 mg/dia (mimetiza insulina, aumenta GLUT4)
• Magnésio 400-600 mg/dia (cofator receptor insulina)
• Ômega-3 2-3g/dia (reduz inflamação, melhora sinalização insulina)
• Canela ceilão 1-6g/dia (melhora sensibilidade ~10-15%)
• Vitamina D 4000 UI/dia (meta 25-OH 50-70 ng/mL)

Dieta otimizada:
• Eliminar: açúcares, farinhas refinadas, sucos, refrigerantes
• Priorizar: proteína magra, gorduras saudáveis (abacate, azeite, nozes), vegetais low-carb
• Índice glicêmico baixo (<55)
• Vinagre de maçã 15 mL antes de refeições (reduz pico glicêmico 20-30%)

Exercício específico:
• HIIT 3x/semana (melhora sensibilidade insulina superior a aeróbico contínuo)
• Musculação pesada (músculo é principal tecido insulino-sensível)
• Caminhar 10 min após cada refeição (reduz glicemia pós-prandial)

Monitoramento:
• QUICKI (insulina + glicemia jejum) a cada 12 semanas
• HbA1c trimestral (meta <5.7%)
• Lipidograma (TG/HDL, meta <2.0)
• Peptídeo C (meta <3.0 ng/mL - marcador hiperinsulinemia)

Reavaliar QUICKI em 12-16 semanas
Meta: QUICKI >0.357 ou melhora >0.030',

    last_review_date = CURRENT_DATE
WHERE id = '83e85e7b-ad24-431c-bf8b-d65eaec835d6';

-- 39. Ecodopplercardiograma - TAPSE
UPDATE score_items
SET
    interpretation = 'TAPSE (Tricuspid Annular Plane Systolic Excursion) mede deslocamento longitudinal do anel tricúspide durante sístole, avaliando função sistólica do ventrículo direito (VD). TAPSE ≥17 mm é normal; <17 mm indica disfunção VD; <14 mm define disfunção severa. TAPSE correlaciona-se com FEVE-VD (fração ejeção VD), mas é mais fácil de obter. Disfunção VD ocorre em: hipertensão pulmonar (HAP), DPOC, embolia pulmonar, IAM VD, miocardiopatias, doenças valvares (insuficiência tricúspide). TAPSE <16 mm prediz mortalidade em IC.',

    clinical_relevance = 'Medicina funcional valoriza TAPSE em: diagnóstico de hipertensão pulmonar (TAPSE <17 mm + PSAP >35 mmHg); avaliação de IC direita (TAPSE <16 mm indica falência VD, necessitando diurese e inotrópicos); monitoramento de DPOC (TAPSE <16 mm correlaciona-se com exacerbações); estratificação de risco pós-embolia pulmonar (TAPSE <17 mm indica pior prognóstico); avaliação de IAM inferior/VD (TAPSE <14 mm indica extensão a VD). TAPSE >20 mm é meta ideal para função VD robusta.',

    optimal_value_male_min = 20.0,
    optimal_value_male_max = 30.0,
    optimal_value_female_min = 20.0,
    optimal_value_female_max = 30.0,
    optimal_value_description = 'TAPSE ótimo: 20-30 mm (padrão MFI). Valores <17 mm indicam disfunção VD. TAPSE 14-17 mm: disfunção leve-moderada. TAPSE <14 mm: disfunção severa, com risco aumentado de IC direita e mortalidade. TAPSE >30 mm é raro, mas pode ser normal em atletas. Ideal: 22-26 mm, refletindo função sistólica VD preservada e acoplamento VD-artéria pulmonar eficiente.',

    health_recommendations = 'TAPSE 14-17 mm (disfunção VD leve-moderada):
• Investigar causa:
  - Ecocardiograma completo: PSAP (pressão sistólica artéria pulmonar), IT (insuficiência tricúspide), tamanho VD
  - Se PSAP >35 mmHg: cateterismo direito (confirmar HAP, medir PCP)
  - Radiografia tórax, gasometria, espirometria (DPOC, intersticial)
  - Angio-TC pulmonar (embolia crônica)

TAPSE < 14 mm (disfunção VD severa):
• IC direita:
  - Diuréticos: furosemida 40-80 mg/dia (reduzir pré-carga VD)
  - Inotrópicos se hipotensão: dobutamina, milrinone (aumentam contratilidade VD)
  - Evitar vasodilatadores (reduzem RVS, pioram perfusão coronária VD)
• Hipertensão pulmonar:
  - HAP Grupo 1: vasodilatadores pulmonares (sildenafil 20 mg 3x/dia, bosentana 125 mg 2x/dia, epoprostenol IV)
  - HAP Grupo 2 (IC esquerda): tratar IC (IECA, beta-bloq, diuréticos)
  - HAP Grupo 3 (DPOC): oxigenoterapia domiciliar se SpO2 <88%

Embolia pulmonar aguda com TAPSE <14 mm:
• Trombolítico (alteplase) se instabilidade hemodinâmica
• Anticoagulação: rivaroxabana 15 mg 2x/dia por 3 semanas, depois 20 mg/dia
• Considerar trombectomia se contraindicação trombólise

IAM ventrículo direito:
• Reposição volêmica (cristaloide 500 mL bolus se PAS <90 mmHg)
• Evitar nitratos e diuréticos (reduzem pré-carga VD)
• Angioplastia primária coronária direita

DPOC com cor pulmonale:
• Oxigenoterapia >15h/dia (meta SpO2 88-92%)
• Broncodilatadores: tiotrópio 18 mcg/dia + formoterol 12 mcg 2x/dia
• Sildenafil 20 mg 3x/dia (reduz PSAP, melhora TAPSE)

Suplementação:
• CoQ10 ubiquinol 300 mg/dia (melhora função VD)
• L-arginina 3g 2x/dia (precursor óxido nítrico, vasodilatação pulmonar)
• Ômega-3 4g/dia (reduz resistência vascular pulmonar)
• Magnésio 400 mg/dia (relaxamento vascular)

Reabilitação cardíaca:
• Exercício supervisionado melhora TAPSE em 2-3 mm
• Treino respiratório (HAP, DPOC)

Reavaliar ecocardiograma em 3-6 meses
Meta: TAPSE >17 mm ou melhora >3 mm',

    last_review_date = CURRENT_DATE
WHERE id = 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d';

-- 40-43. Testes Respiratórios (Metano, H₂, H₂S)
UPDATE score_items
SET
    interpretation = 'Teste respiratório de hidrogênio/metano mede gases produzidos por fermentação bacteriana intestinal de carboidratos não absorvidos (lactose, lactulose, glicose). H₂ elevado indica sobrecrescimento bacteriano intestinal (SIBO) ou má absorção de carboidratos. Metano (CH₄) é produzido por arqueias metanogênicas (Methanobrevibacter smithii), indicando IMO (intestinal methanogen overgrowth) ou SIBO-M. Metano >10 ppm está associado a constipação (reduz motilidade intestinal). H₂S (sulfeto de hidrogênio) é produzido por bactérias redutoras de sulfato (Desulfovibrio), causando diarreia e dor abdominal.',

    clinical_relevance = 'Medicina funcional valoriza testes respiratórios em: diagnóstico de SIBO (H₂ aumenta >20 ppm do basal antes de 90 min); diferenciação SIBO-H₂ (diarreia) de IMO (constipação, metano >10 ppm); investigação de SII refratária (30-80% dos pacientes têm SIBO/IMO); má absorção de nutrientes (B12, ferro, gorduras); diagnóstico de intolerância lactose (H₂ >20 ppm após lactose). Metano basal >3 ppm prediz resposta lenta a antibióticos (arqueias são resistentes a rifaximina isolada).',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 3.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 3.0,
    optimal_value_description = 'Valores ótimos: Metano basal <3 ppm, pico H₂ <20 ppm acima do basal, pico H₂S <3 ppm. Metano basal ≥3 ppm ou pico ≥10 ppm indica IMO (intestinal methanogen overgrowth), associado a constipação e lentificação trânsito. H₂ aumenta >20 ppm antes de 90 min define SIBO. H₂S >3 ppm sugere SIBO-H₂S (disbiose por redutoras de sulfato). Teste negativo não exclui SIBO (10-15% não produzem H₂ - metabolizadores de sulfato ou metano).',

    health_recommendations = 'SIBO-H₂ (H₂ pico >20 ppm, metano <10 ppm):
• Rifaximina 550 mg 3x/dia por 14 dias (erradica 70-80%)
• Dieta low-FODMAP por 4-8 semanas durante tratamento
• Procinéticos pós-tratamento: iberogast 20 gotas 3x/dia ou prucaloprida 2 mg/dia (previne recorrência)

IMO (metano ≥10 ppm):
• Rifaximina 550 mg 3x/dia + neomicina 500 mg 2x/dia por 14 dias (arqueias resistem a rifaximina isolada)
• Ou rifaximina + metronidazol 500 mg 3x/dia por 14 dias
• Considerar alicina (alho) 450 mg 3x/dia por 6 semanas (alternativa natural)

SIBO-H₂S (H₂S >3 ppm):
• Bismuto subsalicilato 524 mg 4x/dia por 4 semanas (inibe redutoras de sulfato)
• Dieta baixa em enxofre: reduzir crucíferos, ovos, carne vermelha, laticínios
• Molibdênio 300 mcg/dia (cofator sulfito oxidase, metaboliza H₂S)

Protocolo pós-antibiótico (prevenir recorrência):
• Procinéticos: iberogast 20 gotas 3x/dia ou gengibre 1g/dia
• Probióticos específicos: Saccharomyces boulardii 250 mg 2x/dia (não coloniza intestino delgado)
• Enzimas digestivas com cada refeição
• Reintroduzir FODMAPs gradualmente

Suporte nutricional:
• Dieta low-FODMAP durante tratamento (reduz substrato fermentação)
• Após erradicação: reintroduzir FODMAPs sistematicamente
• L-glutamina 5g 2x/dia (repara barreira intestinal)
• Óleo de hortelã-pimenta entérico 0.2 mL 3x/dia (reduz gases, dor)

Metano basal persistentemente elevado:
• Considerar metformina 500 mg 2x/dia (inibe arqueias metanogênicas)
• Lovastatin 20 mg/dia (inibe via mevalonato de arqueias - controverso)

Reavaliar teste respiratório 4-6 semanas pós-tratamento',

    last_review_date = CURRENT_DATE
WHERE id = '3e67c010-1164-4e0f-b23f-109557d6d51d';

UPDATE score_items
SET
    interpretation = 'Metano basal (jejum) >3 ppm indica colonização por arqueias metanogênicas no intestino (Methanobrevibacter smithii). Arqueias convertem H₂ em CH₄, reduzindo H₂ detectável no teste respiratório. Metano >3 ppm associa-se a constipação (CH₄ reduz motilidade colônica via receptor TLR4), maior IMC (extração calórica aumentada) e SII-C (síndrome intestino irritável com constipação). Metano >10 ppm é diagnóstico de IMO (intestinal methanogen overgrowth). Arqueias são resistentes a antibióticos convencionais (rifaximina isolada), necessitando combinações terapêuticas.',

    clinical_relevance = 'Medicina funcional valoriza metano basal em: diagnóstico de IMO (metano ≥3 ppm em jejum ou ≥10 ppm pico); predição de resposta terapêutica (metano alto requer neomicina/metronidazol combinado com rifaximina); investigação de constipação refratária (metano >10 ppm presente em 50% dos constipados crônicos); explicação de ganho de peso inexplicado (arqueias aumentam extração calórica). Metano basal <3 ppm é normal, indicando ausência de arqueias patogênicas.',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 3.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 3.0,
    optimal_value_description = 'Metano basal ótimo: <3 ppm. Valores ≥3 ppm indicam colonização por arqueias metanogênicas. Metano 3-10 ppm: colonização leve-moderada, pode contribuir para constipação. Metano ≥10 ppm: IMO estabelecido, fortemente associado a constipação crônica e distensão. Ideal: <1 ppm, refletindo ausência de arqueias patogênicas e microbiota colônica balanceada.',

    health_recommendations = 'Ver protocolo completo no item "Metano expirado" (ID: 3e67c010-1164-4e0f-b23f-109557d6d51d).

Metano basal 3-10 ppm:
• Rifaximina 550 mg 3x/dia + neomicina 500 mg 2x/dia por 14 dias
• Dieta low-FODMAP
• Procinéticos pós-tratamento

Metano ≥10 ppm (IMO severo):
• Rifaximina + neomicina por 14 dias, considerar repetir ciclo se metano persiste
• Ou protocolo herbáceo: alho (alicina) 450 mg 3x/dia + berberina 500 mg 3x/dia por 8 semanas
• Óleo de orégano 200 mg 2x/dia (antimicrobiano arqueias)

Constipação associada:
• Magnésio citrato 400-800 mg/dia (laxante osmótico)
• Psyllium 5g 2x/dia (fibra solúvel)
• Prucaloprida 2 mg/dia (agonista 5-HT4, pró-cinético)

Reavaliar metano basal em 4-6 semanas pós-tratamento',

    last_review_date = CURRENT_DATE
WHERE id = 'a92d20ce-702f-4b15-8817-098b9539c0f0';

UPDATE score_items
SET
    interpretation = 'H₂ pico (máximo durante teste respiratório) >20 ppm acima do basal antes de 90 minutos indica SIBO (sobrecrescimento bacteriano intestino delgado). Bactérias fermentadoras no delgado metabolizam carboidratos não absorvidos (lactulose, glicose), produzindo H₂ rapidamente. Pico após 90-120 min sugere fermentação colônica normal (não SIBO). H₂ pico muito alto (>100 ppm) indica SIBO severo. Ausência de H₂ apesar de sintomas pode indicar metabolização completa para CH₄ (IMO) ou H₂S (SIBO-sulfato).',

    clinical_relevance = 'Medicina funcional valoriza H₂ pico em: diagnóstico de SIBO (pico >20 ppm antes de 90 min tem sensibilidade ~60-80%); diferenciação SIBO de fermentação colônica normal (timing <90 min); monitoramento de resposta a antibióticos (normalização H₂ pico pós-tratamento); investigação de má absorção (B12, ferro, gorduras). H₂ pico <20 ppm não exclui SIBO se sintomas presentes (considerar IMO ou SIBO-H₂S).',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 20.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 20.0,
    optimal_value_description = 'H₂ pico ótimo: <20 ppm acima do basal. Valores 20-50 ppm (antes de 90 min): SIBO leve-moderado. H₂ >50 ppm: SIBO severo. Pico após 120 min: fermentação colônica normal (não patológica). H₂ negativo com metano >10 ppm: IMO (arqueias consomem H₂). Ideal: pico <10 ppm, ocorrendo apenas após 120 min (fermentação colônica fisiológica).',

    health_recommendations = 'Ver protocolo completo no item "Metano expirado". Tratamento baseado em padrão combinado H₂/CH₄/H₂S.',

    last_review_date = CURRENT_DATE
WHERE id = '210eedab-08e7-4f47-ae6a-37aecea9bc16';

UPDATE score_items
SET
    interpretation = 'H₂S pico (sulfeto de hidrogênio) >3 ppm indica SIBO-H₂S, causado por sobrecrescimento de bactérias redutoras de sulfato (Desulfovibrio, Bilophila). H₂S é produzido a partir de aminoácidos sulfurados (cisteína, metionina) e sulfatos dietéticos. H₂S alto associa-se a diarreia (acelera trânsito intestinal), dor abdominal intensa, fezes com odor de ovo podre e inflamação mucosa. H₂S pode inibir produção de H₂ e CH₄, causando teste respiratório "flat" (negativo) apesar de SIBO.',

    clinical_relevance = 'Medicina funcional valoriza H₂S pico em: diagnóstico de SIBO-H₂S (H₂S >3 ppm com H₂/CH₄ baixos); investigação de diarreia refratária com dor abdominal; explicação de teste respiratório negativo apesar de sintomas; diferenciação de subtipos SIBO (H₂S causa sintomas mais severos que H₂). H₂S não é detectado em todos laboratórios (equipamento especializado necessário).',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 3.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 3.0,
    optimal_value_description = 'H₂S pico ótimo: <3 ppm. Valores >3 ppm indicam SIBO-H₂S (sobrecrescimento de redutoras de sulfato). H₂S 3-10 ppm: moderado. H₂S >10 ppm: severo, associado a sintomas intensos (diarreia explosiva, dor abdominal, fezes fétidas). Ideal: indetectável (<1 ppm), refletindo ausência de bactérias redutoras de sulfato patogênicas.',

    health_recommendations = 'Ver protocolo SIBO-H₂S no item "Metano expirado".

H₂S > 3 ppm:
• Bismuto subsalicilato 524 mg 4x/dia por 4 semanas (inibe H₂S)
• Dieta baixa em enxofre: reduzir crucíferos, ovos, carne vermelha, alho, cebola
• Molibdênio 300 mcg/dia (cofator metabolização H₂S)
• Rifaximina 550 mg 3x/dia se H₂ também elevado

H₂S > 10 ppm (severo):
• Bismuto + rifaximina combinados
• Carvão ativado 500 mg 3x/dia (adsorve H₂S, reduz sintomas)
• Considerar dieta elemental 2-3 semanas (repouso intestinal)

Reavaliar H₂S em 4-6 semanas',

    last_review_date = CURRENT_DATE
WHERE id = 'b87387b4-d024-4dbb-be70-84778ca2dce0';

-- 44. Fundoscopia - Retinopatia Hipertensiva
UPDATE score_items
SET
    interpretation = 'Retinopatia hipertensiva é classificada por Keith-Wagener-Barker (modificada): Grau I (estreitamento arteriolar), Grau II (cruzamentos arteriovenosos patológicos - sinal de Salus, Gunn), Grau III (hemorragias, exsudatos algodonosos, hemorragias em chama), Grau IV (papiledema - edema disco óptico, indica emergência hipertensiva). Retinopatia reflete dano microvascular sistêmico (rins, cérebro, coração). Grau III-IV correlaciona-se com lesão de órgão-alvo e requer controle PA urgente. Retinopatia pode reverter parcialmente com tratamento anti-hipertensivo adequado.',

    clinical_relevance = 'Medicina funcional valoriza retinopatia hipertensiva em: estratificação de risco cardiovascular (presença de retinopatia aumenta risco IAM/AVC 2-3x); diagnóstico de emergência hipertensiva (Grau IV com papiledema requer internação, redução PA gradual); monitoramento de controle PA (melhora de retinopatia indica tratamento efetivo); rastreamento de lesão de órgão-alvo (retinopatia Grau II-III indica HAS não-controlada crônica). Fundoscopia anual é recomendada em hipertensos com PA mal controlada ou lesão de órgão-alvo.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Fundoscopia ótima: AUSÊNCIA DE RETINOPATIA (retina normal sem estreitamento arteriolar, cruzamentos AV patológicos ou hemorragias). Grau I (estreitamento arteriolar): HAS leve-moderada, benigna. Grau II (cruzamentos AV patológicos): HAS crônica não-controlada. Grau III (hemorragias, exsudatos): HAS severa, lesão de órgão-alvo. Grau IV (papiledema): emergência hipertensiva, risco AVC. Meta: prevenir progressão para Grau II-III através de controle PA rigoroso (<130/80 mmHg).',

    health_recommendations = 'Retinopatia Grau I (estreitamento arteriolar):
• Controle PA: meta <130/80 mmHg
• IECA ou BRA primeira linha: enalapril 10-20 mg/dia ou losartana 50-100 mg/dia
• Estilo de vida: DASH diet, exercício 150 min/semana, redução sódio <2300 mg/dia
• Fundoscopia controle em 12-24 meses

Retinopatia Grau II (cruzamentos AV patológicos):
• Controle PA rigoroso: <130/80 mmHg
• Terapia combinada: IECA/BRA + bloqueador canal cálcio (anlodipino 5-10 mg/dia)
• Ou IECA/BRA + tiazídico (hidroclorotiazida 12.5-25 mg/dia)
• Fundoscopia controle em 6-12 meses

Retinopatia Grau III (hemorragias, exsudatos):
• HAS severa, lesão de órgão-alvo estabelecida
• Investigação abrangente: ecocardiograma (HVE), função renal (creatinina, proteinúria), cérebro (RM se sintomas neurológicos)
• Controle PA agressivo: meta <120/80 mmHg
• Tripla terapia: IECA/BRA + bloqueador canal cálcio + diurético
• Fundoscopia controle em 3-6 meses

Retinopatia Grau IV (papiledema):
• EMERGÊNCIA HIPERTENSIVA - internação hospitalar
• Redução PA gradual: 10-20% na primeira hora, 25% nas primeiras 24h (redução abrupta pode causar isquemia cerebral/miocárdica)
• Nitroprussiato IV 0.5-10 mcg/kg/min ou labetalol IV 20-80 mg bolus
• Investigar AVC, encefalopatia hipertensiva, IC aguda, dissecção aórtica
• Fundoscopia controle em 1-3 meses (papiledema pode levar semanas para resolver)

Prevenção progressão:
• Controle PA rigoroso desde diagnóstico HAS
• Monitorar PA domiciliar (MAPA ou automedida)
• Aderência medicamentosa (simplificar regime, combinações fixas)
• Perda de peso se IMC >25 (reduz PA 5-20 mmHg por 10 kg)

Suplementação cardioprotectora:
• Ômega-3 2-3g/dia (reduz PA 2-5 mmHg)
• Magnésio 400-600 mg/dia (vasodilatação)
• CoQ10 200 mg/dia (especialmente se estatinas)
• Alho envelhecido 600-1200 mg/dia (reduz PA 5-10 mmHg)
• Vitamina D 4000 UI/dia (meta 25-OH >50 ng/mL)

Dieta DASH:
• Rica em frutas, vegetais, grãos integrais, laticínios magros
• Baixa em sódio (<1500 mg/dia), gorduras saturadas, açúcares
• Potássio 4700 mg/dia (bananas, batata, feijão)

Reavaliar fundoscopia conforme gravidade
Meta: reversão para Grau 0-I ou estabilização',

    last_review_date = CURRENT_DATE
WHERE id = 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb';

COMMIT;

-- =====================================================
-- FIM DO BATCH 5.3 - 44 ITENS ENRIQUECIDOS
-- =====================================================
