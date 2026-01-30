#!/usr/bin/env python3
"""
HISTÓRICO FAMILIAR - 30 Items Enrichment
Conteúdo baseado em epigenética, polimorfismos e medicina de precisão
"""

import requests
import json
import time

API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# Conteúdo científico por doença
DISEASE_CONTENT = {
    "diabetes": {
        "clinicalRelevance": """O histórico familiar de diabetes mellitus tipo 2 (DM2), pré-diabetes ou resistência insulínica representa um dos mais importantes fatores de risco modificáveis através de intervenções epigenéticas. A predisposição genética ao DM2 é poligênica, envolvendo mais de 400 variantes genéticas identificadas por estudos de GWAS (Genome-Wide Association Studies), sendo os polimorfismos mais relevantes: TCF7L2 (aumenta risco em 1,4-1,5x), PPARG Pro12Ala (modulador de adipogênese), KCNJ11 (canais de potássio pancreáticos), IRS1 (sinalização insulínica), e FTO (gene da massa gorda e obesidade).

A arquitetura genética do DM2 familiar demonstra que ter um parente de primeiro grau afetado aumenta o risco em 2-3 vezes, enquanto dois parentes elevam para 5-6 vezes comparado à população geral. Estudos de gêmeos monozigóticos mostram concordância de 70-90%, indicando forte componente hereditário, porém os 10-30% de discordância revelam a crítica importância dos fatores epigenéticos modificáveis: dieta, atividade física, sono, estresse crônico e exposição a disruptores endócrinos.

Na perspectiva da Medicina Funcional Integrativa e Epigenética, o histórico familiar não determina destino, mas sinaliza necessidade de vigilância intensificada e intervenções preventivas personalizadas. Mecanismos epigenéticos-chave incluem: (1) metilação do DNA em genes reguladores da insulina (INS, PPARGC1A), modulada por folato, betaína e colina dietary; (2) modificações de histonas (acetilação/desacetilação) influenciadas por jejum intermitente e restrição calórica; (3) expressão de microRNAs (miR-375, miR-9, miR-126) que regulam secreção de insulina e são alterados por ácidos graxos saturados vs insaturados.

Estudos prospectivos demonstram que indivíduos com alto risco genético (escore poligênico elevado + histórico familiar) podem reduzir incidência de DM2 em 58-71% através do Diabetes Prevention Program: perda de 7% do peso corporal + 150min/semana de atividade física moderada + dieta mediterrânea de baixa carga glicêmica. A intervenção é mais efetiva que metformina (redução de 31%) e seus efeitos são mediados por reversão de marcas epigenéticas patológicas, restauração de sensibilidade insulínica muscular e hepática, e modulação da microbiota intestinal (aumento de Akkermansia muciniphila e Faecalibacterium prausnitzii).""",

        "conduct": """**RASTREAMENTO PRECOCE PERSONALIZADO:**
Iniciar screening anual 10 anos antes da idade de diagnóstico do familiar mais jovem afetado, ou aos 25 anos (o que ocorrer primeiro). Painel metabólico completo: glicemia de jejum, HbA1c, insulina basal, peptídeo-C, HOMA-IR (índice de resistência insulínica), teste oral de tolerância à glicose com medição de insulina em 0-30-60-120min (identifica hiperinsulinemia compensatória precoce), frutosamina. Considerar teste genético de polimorfismos de risco (TCF7L2, PPARG, FTO) para estratificação precisa.

**INTERVENÇÕES NUTRICIONAIS EPIGENÉTICAS:**
Dieta mediterrânea de baixo índice glicêmico (<55) com carboidratos complexos integrais, rica em polifenóis anti-inflamatórios (azeite extravirgem, oleaginosas, frutas vermelhas, vegetais crucíferos). Implementar janela alimentar de 10-12h (jejum intermitente 14/10) para otimizar autofagia celular e sensibilidade insulínica. Suplementação baseada em evidências: inositol (mio-inositol 2g + d-chiro-inositol 50mg), berberina 1500mg/dia (ativa AMPK similar a metformina), cromo picolinato 200-400mcg, ácido alfa-lipoico 600mg, magnésio glicinato 400mg, ômega-3 EPA/DHA 2-3g. Evitar rigorosamente açúcares adicionados, xarope de milho, carboidratos refinados.

**ATIVIDADE FÍSICA ESTRATÉGICA:**
Mínimo 150min/semana aeróbico moderado + 2-3x/semana treino resistido para ganho de massa muscular (músculo é principal sítio de captação de glicose). Exercícios de alta intensidade intervalados (HIIT) 2x/semana melhoram marcadores epigenéticos de sensibilidade insulínica. Reduzir comportamento sedentário: levantar e movimentar a cada 30min.

**MODULAÇÃO DO ESTRESSE E SONO:**
Hipercortisolemia crônica induz resistência insulínica via epigenética. Implementar técnicas de redução de estresse validadas: meditação mindfulness 20min/dia, respiração diafragmática, yoga. Sono de 7-8h/noite com higiene rigorosa: bloquear luz azul 2h antes, quarto escuro (<0,3 lux), temperatura 18-20°C. Sono fragmentado aumenta HbA1c independentemente de outros fatores.

**MONITORAMENTO CONTÍNUO:**
Glicose contínua (CGM) por 14 dias/ano para identificar excursões glicêmicas pós-prandiais ocultas. Composição corporal por bioimpedância ou DEXA: redução de gordura visceral é preditor mais forte que IMC. Marcadores inflamatórios: hs-CRP <1,0 mg/L, IL-6, adiponectina (>10 mcg/mL protetor)."""
    },

    "dislipidemia": {
        "clinicalRelevance": """O histórico familiar de dislipidemia abrange um espectro desde hipercolesterolemia familiar monogênica (HF, prevalência 1:250) até dislipidemias poligênicas multifatoriais. A HF é causada por mutações em genes-chave: LDLR (receptor de LDL, 85-90% dos casos), APOB (apolipoproteína B-100, 5-10%), PCSK9 (pró-proteína convertase, 1-3%), e LDLRAP1 (forma recessiva rara). Mutações heterozigotas em LDLR resultam em LDL-c 190-400 mg/dL desde infância, com risco de eventos cardiovasculares prematuros (homens <55 anos, mulheres <65 anos) aumentado em 10-20 vezes.

Além das formas monogênicas, a dislipidemia familiar comum é altamente poligênica, envolvendo dezenas de variantes genéticas de pequeno efeito individual: APOE ε4 (aumenta LDL e risco cardiovascular), CETP (modulador de HDL), LIPC (lipase hepática), LPL (lipase lipoproteica), APOA5 (hipertrigliceridemia). O escore de risco genético poligênico pode estratificar indivíduos em quintis de risco, com o quintil superior apresentando LDL-c médio 40-50 mg/dL maior que o quintil inferior, independentemente de fatores ambientais.

A interação gene-ambiente na dislipidemia é exemplar da epigenética nutricional: o polimorfismo APOE ε4 aumenta significativamente a resposta do LDL-c a gorduras saturadas dietárias, enquanto portadores de ε2 são mais responsivos a carboidratos refinados (hipertrigliceridemia). Mecanismos epigenéticos incluem metilação de promotores de genes lipogênicos (SREBP-1c, FAS) modulada por ácidos graxos ômega-3, acetilação de histonas em APOA1 (HDL) influenciada por polifenóis, e microRNAs reguladores (miR-33, miR-122, miR-370) alterados por excesso de frutose e gorduras trans.

Estudos longitudinais demonstram que indivíduos com histórico familiar positivo para dislipidemia e eventos cardiovasculares precoces apresentam progressão acelerada de aterosclerose subclínica (espessamento médio-intimal carotídeo, escore de cálcio coronariano) décadas antes de manifestações clínicas, enfatizando a importância do screening e intervenção precoces baseados em risco genético-familiar.""",

        "conduct": """**RASTREAMENTO INTENSIFICADO:**
Iniciar perfil lipídico completo aos 20 anos (ou 10 anos se HF familiar conhecida): colesterol total, LDL-c, HDL-c, VLDL, triglicerídeos, não-HDL-c, Apo B, Apo A1, relação Apo B/Apo A1, Lp(a) (lipoproteína(a), fator de risco independente genético). Repetir anualmente até 30 anos, depois a cada 2-3 anos se normal. LDL-c >190 mg/dL sem causas secundárias sugere HF, considerar teste genético de LDLR/APOB/PCSK9. Escore de cálcio coronariano aos 40-45 anos para estratificação de risco se LDL-c limítrofe.

**INTERVENÇÕES NUTRICIONAIS BASEADAS EM GENOTIPO:**
Dieta portfolio para redução de LDL (-25-30% sem medicação): fibras solúveis 10-25g/dia (aveia, psyllium, leguminosas), fitoesteróis 2g/dia (esteróis vegetais fortificados), proteína de soja 25g/dia, oleaginosas 30-45g/dia (amêndoas, nozes ricas em ômega-3 ALA). Substituir gorduras saturadas (<7% calorias) por monoinsaturadas (azeite, abacate) e poliinsaturadas (peixes gordos 3-4x/semana). Eliminar completamente gorduras trans industriais. Se APOE ε4+, restrição ainda mais rigorosa de saturadas (<5%).

Para hipertrigliceridemia familiar: redução drástica de carboidratos refinados e álcool, aumentar ômega-3 marinho EPA/DHA 3-4g/dia (dose farmacológica para triglicerídeos >500 mg/dL), MCT oil (triglicerídeos de cadeia média) 15-30ml/dia, evitar frutose adicionada que induz lipogênese hepática via ativação de SREBP-1c.

**SUPLEMENTAÇÃO EVIDENCIADA:**
Niacina (ácido nicotínico) liberação prolongada 1-2g/dia aumenta HDL 15-35% e reduz LDL 10-20%, porém requer monitoramento hepático. Berberina 1500mg/dia reduz LDL 20-25% via upreg

ulação de receptores LDL. Arroz vermelho fermentado (contém monacolina K/lovastatina natural) 1200-2400mg, atenção a variabilidade de dose e interações. Probióticos específicos: Lactobacillus reuteri NCIMB 30242 reduz LDL até 11,6%. Suplementos antioxidantes: CoQ10 100-200mg se estatinas futuras, vitamina E natural (não sintética).

**ATIVIDADE FÍSICA PARA LIPÍDIOS:**
Exercícios aeróbicos 150-300min/semana aumentam HDL 5-10% e reduzem triglicerídeos 20-30%. Intensidade moderada-alta mais efetiva. Treinamento resistido complementar melhora perfil lipídico via ganho de massa magra.

**MONITORAMENTO E META-TERAPIA:**
Meta de LDL-c baseada em risco: se HF ou risco muito alto <55 mg/dL, risco alto <70 mg/dL, risco intermediário <100 mg/dL. Considerar não-HDL-c e Apo B como alvos superiores ao LDL-c. Se intervenção lifestyle intensiva por 6-12 meses não atingir metas, considerar estatinas moderadas-altas doses + ezetimiba (inibidor absorção). Lp(a) elevado (>50 mg/dL) sem terapia comprovada, foco em controle rigoroso de outros fatores."""
    },

    "cardiovascular": {
        "clinicalRelevance": """O histórico familiar de doença cardiovascular prematura (DCV) - definida como IAM, revascularização coronariana, AVC isquêmico, ou morte súbita cardíaca em parente de primeiro grau masculino <55 anos ou feminino <65 anos - representa um dos mais fortes preditores independentes de risco cardiovascular, duplicando a 4-5x o risco de eventos mesmo após ajuste para fatores de risco tradicionais.

A base genética da DCV familiar é altamente heterogênea e multifatorial, envolvendo variantes em centenas de genes que regulam lipoproteínas (LDLR, APOB, PCSK9, APOE), trombose e fibrinólise (F2, F5 Leiden aumenta risco trombótico 3-8x, MTHFR C677T homozigose eleva homocisteína), metabolismo endotelial (eNOS, ACE I/D polimorfismo), inflamação vascular (IL6, TNF-α, CRP high-sensitivity), e metabolismo glicêmico (TCF7L2, PPARG). Estudos de escore de risco poligênico para doença arterial coronariana identificaram indivíduos no quintil superior de risco genético com incidência de eventos 3-4x maior que quintil inferior.

Mecanismos epigenéticos cruciais na DCV incluem: (1) metilação aberrante de genes anti-aterogênicos (ABCA1, APOA1, eNOS) induzida por inflamação crônica de baixo grau, estresse oxidativo e hiperglicemia; (2) modificações de histonas em genes pró-inflamatórios endoteliais (VCAM-1, ICAM-1, E-selectina) perpetuadas por dieta aterogênica; (3) microRNAs disfuncionais (miR-33, miR-146a, miR-126, miR-155) que regulam metabolismo lipídico, função endotelial, e estabilidade de placa aterosclerótica.

A interação gene-ambiente é exemplificada pelo polimorfismo 9p21 (locus de maior risco genético para DCV, OR 1,3-1,5), cujo efeito deletério é amplificado por tabagismo, dieta pró-inflamatória e sedentarismo, mas significativamente atenuado por dieta prudente rica em frutas e vegetais. Estudos prospectivos demonstram que indivíduos com alto risco genético (histórico familiar + escore poligênico elevado) que aderem a lifestyle saudável (4-5 fatores: não fumar, IMC<30, atividade física regular, dieta saudável) reduzem risco de eventos em 46-50%, evidenciando potencial modificador da epigenética lifestyle.""",

        "conduct": """**SCREENING CARDIOVASCULAR AGRESSIVO:**
Iniciar aos 30 anos (ou 10 anos antes do evento familiar mais precoce): perfil lipídico avançado (LDL-c, não-HDL-c, Apo B, Lp(a), LDL partícula pequena-densa), glicemia/HbA1c/insulina/HOMA-IR, hs-CRP (<1 mg/L ideal, 1-3 intermediário, >3 alto risco), homocisteína (<10 µmol/L), fibrinogênio, hemograma (hematócrito elevado = viscosidade aumentada), função renal, microalbuminúria. Considerar teste genético de variantes alto-risco (F5 Leiden, F2 protrombina, APOE, 9p21, PCSK9 gain-of-function). Escore de cálcio coronariano (CAC score) aos 40-45 anos para reclassificação de risco: CAC=0 baixo risco, 1-100 moderado, >100 alto, >300 muito alto.

**INTERVENÇÃO NUTRICIONAL CARDIOPROTETORA:**
Dieta mediterrânea rigorosa (evidência nível I do estudo PREDIMED: redução 30% eventos cardiovasculares): azeite extravirgem 40-50ml/dia (polifenóis anti-inflamatórios), oleaginosas 30g/dia (nozes reduzem eventos 28% em meta-análises), peixes gordos água fria 3-4x/semana (EPA/DHA 2-3g/dia), abundância de vegetais coloridos (9-11 porções/dia, polifenóis), frutas vermelhas (antocianinas), leguminosas diárias, grãos integrais 100% (aveia reduz LDL), vinho tinto moderado opcional (resveratrol 1 taça/dia mulheres, 2 homens). Eliminar: carnes processadas (cada 50g/dia aumenta risco 18%), refrigerantes e bebidas açucaradas, gorduras trans, sódio <2g/dia.

Suplementação estratégica: ômega-3 EPA/DHA 2-4g/dia (redução 25% mortalidade cardíaca em estudos), CoQ10 100-300mg/dia (melhora função endotelial, essencial se estatinas), magnésio 400-600mg (anti-arrítmico, reduz PA), vitamina K2 MK-7 180-360mcg (previne calcificação arterial), vitamina D3 para manter 40-60 ng/mL, alho envelhecido 600-1200mg.

**ATIVIDADE FÍSICA TERAPÊUTICA:**
Mínimo 150min/semana moderado ou 75min vigoroso, ideal 200-300min. Exercício aeróbico regular melhora função endotelial, aumenta HDL, reduz inflamação, melhora variabilidade cardíaca (marcador SNA). Treinamento intervalado de alta intensidade (HIIT) superior para fitness cardiorrespiratório. Musculação 2-3x/semana complementar.

**CONTROLE RIGOROSO DE FATORES DE RISCO:**
PA <130/80 mmHg (se alto risco <120/80), LDL-c <70 mg/dL (se risco muito alto <55), HbA1c <5,7%, IMC <25 ou circunferência abdominal <94cm homens/<80cm mulheres. Cessar tabagismo absoluto (risco 2-4x). Manejo estresse: cortisol elevado crônico = inflamação vascular, praticar meditação/yoga/biofeedback. Sono 7-8h/noite, apneia do sono tratada.

**FARMACOTERAPIA PREVENTIVA:**
Se LDL-c >100 mg/dL apesar de lifestyle ou CAC score >100: estatina moderada-alta intensidade (atorvastatina 20-40mg, rosuvastatina 10-20mg). Se Lp(a) >50 mg/dL + LDL alto: considerar inibidores PCSK9. AAS 81-100mg/dia se >50 anos + risco intermediário-alto (discutir risco-benefício sangramento). Metformina 500-1000mg se pré-diabetes + histórico familiar DM2."""
    },

    "renal": {
        "clinicalRelevance": """O histórico familiar de doença renal crônica (DRC), nefrites, síndrome nefrótica, litíase renal recorrente ou infecções urinárias de repetição (ITU) sinaliza múltiplas condições hereditárias com espectros distintos de transmissão genética e potencial de intervenção preventiva.

DRC familiar pode ser monogênica - doença renal policística autossômica dominante (PKD1/PKD2, prevalência 1:1000, principal causa de DRC hereditária), síndrome de Alport (mutações em colágeno tipo IV COL4A3/A4/A5, hematúria familiar evoluindo para insuficiência renal), nefropatia associada a APOL1 (duas cópias de variantes G1/G2 aumentam risco 7-30x em afrodescendentes) - ou poligênica associada a diabetes, hipertensão e obesidade familiar com variantes em UMOD (nefropatia intersticial), SHROOM3, DAB2 (função podocitária).

Litíase renal tem herdabilidade de 45-60%, com formas monogênicas raras (cistinúria, hiperoxalúria primária) e formas multifatoriais comuns associadas a polimorfismos em transportadores iônicos (VDR receptor vitamina D, CASR receptor cálcio-sensível, SLC genes de transporte). Estudos de gêmeos mostram concordância 32% para cálculos renais, indicando forte contribuição ambiental modificável.

ITU recorrente familiar em mulheres está associada a variantes em genes de resposta imune inata (TLR4, CXCR1) e receptores de adesão bacteriana (FUT2 secretor status influencia binding de E.coli uropatogênica). Mecanismos epigenéticos na doença renal incluem metilação de genes fibrogênicos (TGF-β, CTGF) modulada por hiperglicemia e proteinúria, modificações de histonas em podócitos (células filtrantes glomerulares) influenciadas por AGEs (advanced glycation end-products), e microRNAs reguladores de fibrose (miR-21, miR-29, miR-192).

A progressão de DRC pode ser significativamente desacelerada por intervenções preventivas: controle rigoroso de PA (<130/80, <125/75 se proteinúria), controle glicêmico estrito (HbA1c <7,0%), restrição proteica moderada (0,8-1,0 g/kg/dia), redução de sódio, e uso de IECA/BRA (inibem sistema renina-angiotensina, reduzem proteinúria e fibrose intersticial via efeitos pleiotrópicos).""",

        "conduct": """**RASTREAMENTO NEFROLÓGICO DIRECIONADO:**
Baseline aos 20-25 anos: creatinina sérica, TFG estimada (eGFR, normal >90 ml/min/1,73m²), ureia, eletrólitos (Na, K, Ca, P, Mg), urina I completa (densidade, pH, proteínas, hemácias, leucócitos, cilindros), relação proteína/creatinina urinária (<0,2 normal), microalbuminúria (30-300 mg/dia indica lesão glomerular incipiente). Se história de litíase: cálcio urinário 24h, oxalato, ácido úrico, citrato, cistina, volume urinário. Repetir anualmente se alterações, a cada 2-3 anos se normal.

Se história familiar de PKD: ultrassom renal aos 18-20 anos (cistos bilaterais múltiplos), considerar RM se ultrassom inconclusivo. Se hematúria familiar persistente: investigar síndrome de Alport (biópsia renal com microscopia eletrônica, teste genético COL4A). Se ITU recorrente (>2 episódios/ano): urinocultura, ultrassom rins/vias urinárias para excluir anomalias anatômicas, considerar cistoscopia se refratário.

**NEFROPROTEÇÃO NUTRICIONAL:**
Hidratação otimizada: 2,5-3L líquidos/dia (água preferencialmente), volume urinário >2L/dia previne litíase e ITU. Redução de sódio <2g/dia (previne HAS e proteinúria). Restrição proteica moderada 0,8-1,0 g/kg/dia se eGFR <60, evitar excesso de proteína animal que aumenta carga ácida renal e hiperfiltração glomerular.

Para prevenção de litíase cálcica: dieta normocálcica 1000-1200mg/dia (baixo cálcio aumenta oxalato), reduzir oxalato (espinafre, beterraba, chocolate, nozes em excesso), reduzir proteína animal, aumentar citrato (limonada natural 120ml/dia de suco limão), reduzir sódio. Se litíase úrica: alcalinizar urina (pH 6,5-7,0) com citrato de potássio, reduzir purinas (carnes vermelhas, frutos do mar, álcool).

Suplementação: citrato de potássio 30-60 mEq/dia (litíase cálcica recorrente, alcaliniza urina), magnésio 400-600mg (inibe cristalização cálcio-oxalato), probióticos Lactobacillus/Bifidobacterium degradam oxalato intestinal, vitamina B6 (piridoxina) 100-200mg/dia (reduz oxalato endógeno se hiperoxalúria), cranberry 36mg PACs/dia (prevenção ITU recorrente, inibe adesão E.coli).

**CONTROLE PRESSÓRICO E GLICÊMICO RIGOROSO:**
PA meta <130/80 (se proteinúria <125/75), preferencialmente com IECA (enalapril, ramipril) ou BRA (losartana, valsartana) mesmo se PA normal mas proteinúria presente - efeito nefroprotetor independente via bloqueio RAAS. HbA1c <7,0% se diabético/pré-diabético. Evitar AINEs crônicos (anti-inflamatórios lesam rins), contraste iodado se eGFR <60.

**MONITORAMENTO LONGITUDINAL:**
TFG e proteinúria anuais, declínio >5 ml/min/ano indica progressão acelerada (normal declínio 1 ml/min/ano após 40 anos). Encaminhar nefrologista se eGFR <45, proteinúria persistente, hematúria glomerular, HAS refratária, ou hipercalemia recorrente."""
    },

    "autoimune": {
        "clinicalRelevance": """O histórico familiar de doenças autoimunes - lúpus eritematoso sistêmico (LES), artrite reumatoide (AR), esclerodermia, doença de Crohn, retocolite ulcerativa (RCU), asma alérgica - indica predisposição genética para desregulação imunológica caracterizada por perda de tolerância a auto-antígenos e inflamação crônica mediada por citocinas.

A arquitetura genética das doenças autoimunes é complexa e poligênica, com forte contribuição do complexo principal de histocompatibilidade (MHC/HLA): HLA-DR3/DR4 aumenta risco de LES 2-3x e diabetes tipo 1, HLA-DRB1 shared epitope confere risco de AR 3-6x, HLA-B27 associa-se a espondilite anquilosante (OR 50-100), HLA-DQ2/DQ8 são necessários mas não suficientes para doença celíaca (30% população geral tem, mas apenas 1% desenvolve doença). Além do HLA, centenas de variantes em genes imunorregulatórios contribuem: PTPN22 (regulador de células T, risco para múltiplas autoimunes), IL23R (Crohn e psoríase), CTLA4 (checkpoint imune), NOD2 (sensoriamento bacteriano, Crohn).

A concordância em gêmeos monozigóticos é parcial - LES 25-50%, AR 15-30%, Crohn 35-50%, RCU 10-15% - demonstrando que genética é necessária mas não suficiente, enfatizando o papel crítico de gatilhos ambientais e fatores epigenéticos. Mecanismos epigenéticos-chave incluem: (1) hipometilação global do DNA em células T CD4+ de pacientes com LES, levando a hiperexpressão de genes inflamatórios; (2) acetilação aberrante de histonas em promotores de citocinas pró-inflamatórias (IL-6, TNF-α, IL-17); (3) microRNAs desregulados (miR-146a, miR-155, miR-21) que amplificam cascatas inflamatórias.

Gatilhos ambientais validados incluem: infecções virais (EBV fortemente associado a LES e esclerose múltipla via molecular mimicry), disbiose intestinal (depleção de Faecalibacterium e Roseburia, expansão de Proteobacteria em DII), deficiência de vitamina D (<20 ng/mL dobra risco de autoimunidade), tabagismo (AR e Crohn), exposição a sílica e solventes orgânicos (esclerodermia), estresse crônico (desregulação eixo HPA e inflamação), e permeabilidade intestinal aumentada (leaky gut permite translocação de LPS e ativação imune sistêmica).""",

        "conduct": """**VIGILÂNCIA IMUNOLÓGICA PROATIVA:**
Screening baseline aos 18-20 anos: hemograma completo (citopenias autoimunes), VHS e hs-CRP (inflamação sistêmica), FAN (anticorpos anti-núcleo, se positivo refinar com anti-DNA, anti-Sm, anti-Ro/La), fator reumatoide e anti-CCP (AR), TSH e anti-TPO (tireoidite Hashimoto, prevalência 5-10% em mulheres), anti-transglutaminase IgA (doença celíaca), vitamina D 25-OH. Repetir a cada 2-3 anos ou se sintomas sugestivos (artralgia migratória, rash fotossensível, fadiga inexplicada, sintomas GI crônicos).

**INTERVENÇÃO NUTRICIONAL IMUNOMODULATÓRIA:**
Protocolo autoimune (AIP - Autoimmune Protocol): eliminar alimentos pró-inflamatórios e permeabilizadores intestinais por 6-12 semanas: glúten (mesmo sem doença celíaca, gliadina aumenta zonulina e permeabilidade), laticínios A1 (caseína pró-inflamatória), grãos e leguminosas (lectinas e fitatos), nightshades (solanáceas: tomate, berinjela, pimentas se sensível), ovos (albúmina), oleaginosas/sementes, álcool, açúcares refinados, aditivos.

Alimentos terapêuticos: caldo de ossos (colágeno, glicina, glutamina restauram barreira intestinal), peixes gordos 4-5x/semana (EPA/DHA anti-inflamatório potente, meta-análises mostram redução de atividade de AR), vegetais fermentados (probióticos naturais), azeite extravirgem (oleocanthal inibe COX-2 similar a ibuprofeno), cúrcuma (curcumina inibe NF-κB, 1000mg/dia com piperina), gengibre (gingerois anti-inflamatórios), chá verde (EGCG imunomodulador).

Suplementação evidenciada: vitamina D3 5000-10000 UI/dia para manter 40-60 ng/mL (estudos mostram redução 20-50% risco autoimunidade), ômega-3 EPA/DHA 3-4g/dia (dose anti-inflamatória), glutamina 5-10g/dia (combustível enterócitos, reduz permeabilidade), N-acetilcisteína 600-1200mg (antioxidante, modula glutationa), probióticos multi-cepas 50-100 bilhões UFC (Lactobacillus plantarum, rhamnosus, Bifidobacterium longum), prebióticos (inulina 10g/dia alimenta Faecalibacterium).

**RESTAURAÇÃO DA BARREIRA INTESTINAL:**
Protocolo 5R: (1) Remover patógenos (SIBO, Candida, parasitas via testes respiratório/PCR fezes), (2) Repovoar com probióticos, (3) Reparar com glutamina + zinco carnosina + DGL (licorice deglicirrizado) + slippery elm, (4) Reinocular microbiota diversa, (5) Rebalancear com enzimas digestivas e HCl se hipocloridria.

**MODULAÇÃO EPIGENÉTICA DO ESTRESSE:**
Hipercortisolemia crônica -> inflamação sistêmica via NF-κB. Implementar redução de estresse validada cientificamente: meditação mindfulness 20-30min/dia (reduz citocinas inflamatórias 25-40%), yoga terapêutico, tai chi (melhora sintomas AR), biofeedback de variabilidade cardíaca, sono 8-9h/noite (privação aumenta IL-6 e TNF-α).

**EXPOSIÇÃO SOLAR E VITAMINA D:**
20-30min exposição solar diária (braços/pernas) sem protetor antes das 10h ou após 16h para síntese vitamina D (controle epigenético de >1000 genes). Evitar deficiência absoluta (<20 ng/mL aumenta risco autoimune 50-100%).

**EXERCÍCIO IMUNOMODULADOR:**
Atividade física regular moderada (150min/semana) reduz citocinas pró-inflamatórias, melhora microbiota, restaura tolerância imune. Evitar overtraining (imunossupressor). Yoga e tai chi particularmente benéficos para doenças autoimunes reumáticas."""
    },

    "viral_cronica": {
        "clinicalRelevance": """O histórico familiar de doenças virais crônicas - HIV, hepatites B e C - apresenta implicações distintas de transmissão, risco genético de progressão e estratégias preventivas. Embora infecções virais crônicas não sejam hereditárias no senso clássico (não há transmissão via DNA germinativo), o contexto familiar é relevante por: (1) transmissão vertical mãe-filho (HIV, HBV, HCV via parto/aleitamento se não tratada), (2) transmissão horizontal intrafamiliar (HBV via contato com sangue/fluidos), e (3) predisposição genética à suscetibilidade de infecção e progressão de doença.

Polimorfismos genéticos influenciam significativamente o curso de infecções virais crônicas: HLA-B*57:01 e B*27 associam-se a progressão lenta de HIV (elite controllers com carga viral indetectável sem terapia), deleção de 32 pares de base em CCR5 (CCR5-Δ32 homozigose confere resistência quase completa a HIV, heterozigose retarda progressão), polimorfismos em IL28B (interferon-lambda) predizem resposta a tratamento de HCV (CC genótipo tem 2-3x maior chance de clearance viral espontâneo e resposta a interferon). Para HBV, variantes em HLA-DP, HLA-DQ influenciam risco de cronificação vs clearance.

Mecanismos epigenéticos nas infecções virais crônicas são bidirecionais: (1) vírus modulam epigenoma do hospedeiro para favorecer replicação (HIV induz hipometilação de genes pró-virais, metilação de genes antivirais; HBV e HCV alteram metilação de supressores tumorais contribuindo para hepatocarcinoma); (2) estado epigenético do hospedeiro influencia resposta imune e progressão (metilação de promotores de interferon associa-se a cronificação de HCV, acetilação de histonas em genes de células T CD8+ citotóxicas determina controle de HIV).

A transmissão vertical de HIV/HBV/HCV é quase completamente prevenível com intervenções modernas: terapia antirretroviral (TARV) em gestantes HIV+ reduz transmissão de 25-30% para <1%, vacinação de recém-nascidos de mães HBsAg+ com imunoglobulina + vacina reduz transmissão HBV de 90% para <5%, screening e tratamento de HCV em gestantes (DAAs - antivirais de ação direta) podem eliminar transmissão vertical. Portanto, o histórico familiar deve motivar vigilância intensificada, testagem regular, e implementação rigorosa de medidas preventivas.""",

        "conduct": """**RASTREAMENTO SISTEMÁTICO E PREVENÇÃO PRIMÁRIA:**
Testagem baseline aos 15-18 anos (início atividade sexual): HIV (teste combo Ag/Ac 4ª geração, janela 18-45 dias), HBsAg + anti-HBs + anti-HBc total (status imunológico HBV), anti-HCV (se risco percutâneo). Repetir anualmente se comportamentos de risco (múltiplos parceiros, não uso consistente de preservativos, usuário de drogas injetáveis, tatuagens/piercings sem biossegurança), ou a cada 3 anos se baixo risco.

**VACINAÇÃO UNIVERSAL HBV:**
Esquema 0-1-6 meses (Engerix-B ou equivalente), dose dobrada se imunossupressão. Checar anti-HBs 1-2 meses após série: >10 mIU/mL = proteção, <10 = não-respondedor (5-10% população, considerar revacinação ou dose dobrada). Se contato familiar HBsAg+: vacinação urgente + imunoglobulina anti-HBV (HBIG) dentro de 24-48h de exposição.

**PROFILAXIA PRÉ-EXPOSIÇÃO (PrEP) HIV:**
Se parceiro HIV+ ou risco substancial: tenofovir/emtricitabina (Truvada) diária reduz risco 92-99% se aderência >4 doses/semana. Monitoramento trimestral: HIV, função renal, triagem ISTs.

**OTIMIZAÇÃO IMUNOLÓGICA NUTRICIONAL:**
Sistema imune robusto reduz risco de aquisição e progressão de virais crônicas. Dieta anti-inflamatória mediterrânea rica em antioxidantes: vegetais coloridos (carotenoides, vitamina C 500-1000mg/dia), frutas vermelhas (antocianinas), alho (alicina antiviral), cogumelos medicinais (shiitake, reishi: beta-glucanos imunoestimulantes), chá verde (EGCG antiviral in vitro para HCV), cúrcuma (curcumina modula NF-κB).

Suplementação imunomoduladora: vitamina D 5000 UI/dia (deficiência <20 ng/mL associa-se a maior risco infecções virais e progressão HIV), zinco 30-50mg/dia (essencial para células T, deficiência prejudica imunidade celular), vitamina C 1000-2000mg/dia (antioxidante, suporta neutrófilos), N-acetilcisteína 600-1200mg (precursor glutationa, antioxidante hepatoprotetor para HBV/HCV), probióticos (modulam microbiota e imunidade mucosa), EGCG chá verde 400-800mg, quercetina 1000mg (flavonoide antiviral).

**HEPATOPROTEÇÃO SE HBV/HCV FAMILIAR:**
Evitar rigorosamente hepatotóxicos: álcool (zero se hepatite crônica), paracetamol >2g/dia, metabólitos de drogas, aflatoxinas (mofo em oleaginosas/grãos mal armazenados). Suplementos hepatoprotetores: silimarina (cardo mariano) 420mg/dia (antioxidante, anti-fibrótico, estudos mistos mas seguro), SAMe 800-1600mg/dia (doador de metil, regeneração hepatocitária), ácido alfa-lipoico 600mg (antioxidante mitocondrial).

**MONITORAMENTO SE INFECÇÃO ESTABELECIDA:**
HIV: carga viral indetectável (<50 cópias/mL) com TARV, CD4 >500, triagem anual tuberculose, Papanicolau, lipídios (TARV pode causar dislipidemia). HBV: HBsAg, HBV DNA quantitativo, ALT, AFP (alfa-fetoproteína para hepatocarcinoma) a cada 6-12 meses, ultrassom hepático anual se cirrose/alta replicação. HCV: RNA HCV quantitativo, tratamento com DAAs (sofosbuvir/velpatasvir ou equivalente) cura >95%, pós-tratamento monitorar fibrose hepática (Fibroscan) e hepatocarcinoma se cirrose prévia.

**REDUÇÃO DE DANOS:**
Nunca compartilhar agulhas, seringas, equipamento de injeção. Preservativo consistente. Testagem de parceiros. PEP (profilaxia pós-exposição HIV) em 72h se exposição de alto risco."""
    },

    "hipertensao": {
        "clinicalRelevance": """A hipertensão arterial sistêmica (HAS) com histórico familiar representa uma das condições cardiovasculares com maior componente hereditário, exibindo herdabilidade estimada de 30-60% por estudos de gêmeos e famílias. A arquitetura genética da HAS essencial (95% dos casos) é complexa e poligênica, envolvendo centenas de variantes genéticas de pequeno efeito individual que regulam sistemas fisiológicos críticos para homeostase pressórica.

Genes e vias-chave incluem: (1) Sistema renina-angiotensina-aldosterona (RAAS): polimorfismo AGT M235T (angiotensinogênio) aumenta risco 20-30%, ACE I/D (enzima conversora de angiotensina, alelo D associado a maior atividade enzimática e risco), AGTR1 (receptor AT1), CYP11B2 (aldosterona sintase); (2) Homeostase de sódio: WNK1, WNK4 (quinases reguladoras de canais renais), SLC12A3 (transportador Na-Cl cotransporter), NPPA/NPPB (peptídeos natriuréticos); (3) Função endotelial: eNOS (óxido nítrico sintase endotelial, vasodilatador), ADD1 (aducina, regula bomba Na/K); (4) Modulação adrenérgica: ADRB1, ADRB2 (receptores beta-adrenérgicos).

Estudos de escore de risco genético poligênico identificam indivíduos no quintil superior com PA sistólica 3-4 mmHg maior que quintil inferior, traduzindo-se em risco 16-20% maior de eventos cardiovasculares. Crítico: mesmo com alto risco genético, fatores lifestyle modulam dramaticamente a expressão fenotípica - o estudo DASH demonstrou que dieta otimizada reduz PA sistólica em média 11 mmHg em hipertensos, independentemente de genética.

Mecanismos epigenéticos na HAS incluem: (1) metilação de genes do RAAS (AGT, ACE) modulada por ingestão de sódio e obesidade visceral; (2) acetilação de histonas em genes de canais iônicos renais influenciada por aldosterona e hipercortisolemia; (3) microRNAs vasculares (miR-21, miR-155, miR-143/145) desregulados por inflamação vascular e disfunção endotelial, alterando proliferação de células musculares lisas e rigidez arterial.

A interação gene-ambiente é exemplificada pela sensibilidade ao sódio: indivíduos com variantes de risco em genes de transporte renal de sódio exibem elevação pressórica exagerada com alto consumo de sal (>6g/dia), mas PA normaliza com restrição (<3g/dia). Obesidade visceral amplifica risco genético via hiperinsulinemia (retenção renal de sódio), hiperativação simpática, e inflamação de baixo grau.""",

        "conduct": """**SCREENING PRESSÓRICO INTENSIFICADO:**
Iniciar monitoramento aos 18 anos (ou mais cedo se obesidade infantil): PA em todas consultas médicas. Técnica correta: repouso 5min, braço apoiado altura cardíaca, manguito apropriado (cobrir 80% circunferência), média de 2-3 medidas. Diagnóstico HAS: PA ≥130/80 mmHg em ≥2 ocasiões (nova diretriz ACC/AHA 2017) ou ≥140/90 mmHg (diretriz brasileira). Se limítrofe ou HAS grau I: MAPA (monitoramento ambulatorial 24h) ou MRPA (medidas residenciais) para confirmar e excluir hipertensão do avental branco.

Se HAS confirmada ou PA 120-129/<80 (elevada): investigação secundária em jovens <30 anos: função renal, eletrólitos, aldosterona/renina (hiperaldosteronismo primário 5-10% HAS), metanefrinas (feocromocitoma), TSH, cortisol livre urinário (Cushing), ultrassom renal (estenose artéria renal, PKD). Avaliação de lesão órgão-alvo: ECG (hipertrofia ventricular esquerda), ecocardiograma se HVE ou sintomas, microalbuminúria, fundoscopia (retinopatia hipertensiva).

**INTERVENÇÃO NUTRICIONAL ANTI-HIPERTENSIVA:**
Dieta DASH (Dietary Approaches to Stop Hypertension, evidência nível 1A): rica em potássio (4700mg/dia de frutas, vegetais, leguminosas), magnésio (400-600mg, oleaginosas, vegetais folhosos), cálcio (1200mg, laticínios low-fat ou vegetais), fibras (25-35g), proteína vegetal, gorduras insaturadas. Reduz PA sistólica média 11 mmHg em hipertensos, 6 mmHg em normotensos.

Restrição rigorosa de sódio <2000mg/dia (ideal 1500mg): eliminar alimentos processados (70-80% sódio dietário), não adicionar sal, usar especiarias/ervas. Cada redução de 1000mg sódio diminui PA 5-6 mmHg em hipertensos sensíveis ao sal. Relação potássio/sódio >2:1 ideal.

Suplementação evidenciada: magnésio 400-600mg/dia (vasodilatador via relaxamento muscular liso, meta-análises mostram redução PA 2-3 mmHg), potássio 1000-2000mg suplementar (se dieta insuficiente e sem IRC), CoQ10 100-200mg/dia (antioxidante, melhora função endotelial, reduz PA 11/7 mmHg em meta-análises), alho envelhecido 600-1200mg (contém alicina vasodilatadora, reduz PA 8-9/5-6 mmHg), hibiscus (chá de hibisco 3 xícaras/dia reduz PA comparável a captopril 25mg em estudos), ômega-3 EPA/DHA 2-3g/dia (reduz PA 3-4/2-3 mmHg).

**PERDA DE PESO E COMPOSIÇÃO CORPORAL:**
Cada 1kg redução de peso diminui PA 1 mmHg. Meta IMC <25 ou circunferência abdominal <94cm homens/<80cm mulheres (gordura visceral é crítica). Estratégias: restrição calórica moderada 500kcal/dia, jejum intermitente 16/8, dieta mediterrânea baixa em carboidratos refinados.

**ATIVIDADE FÍSICA TERAPÊUTICA:**
Exercício aeróbico 150-300min/semana intensidade moderada (caminhada rápida, ciclismo, natação) reduz PA sistólica 5-8 mmHg, diastólica 3-5 mmHg - efeito comparável a monoterapia farmacológica leve. Mecanismo: vasodilatação mediada por óxido nítrico, redução resistência vascular periférica, modulação simpática. Treinamento resistido complementar 2-3x/semana (não exclusivo). Exercícios isométricos (handgrip) 3x/semana emergem como adjuvantes potentes.

**MODERAÇÃO DE ÁLCOOL E CESSAÇÃO DE TABAGISMO:**
Álcool <2 doses/dia homens, <1 mulher (excesso eleva PA dose-dependente). Tabagismo: cada cigarro eleva PA agudamente 10-15 mmHg por 15-30min, promove rigidez arterial, aterosclerose. Cessar absolutamente.

**MANEJO DE ESTRESSE:**
Estresse crônico -> hiperativação simpática + hipercortisolemia -> HAS. Técnicas validadas: meditação mindfulness, respiração lenta (6 resp/min, device RESPeRATE aprovado FDA), biofeedback, yoga, tai chi (reduções de 5-10 mmHg). Sono 7-8h/noite, tratar apneia se presente (CPAP reduz PA 4-6 mmHg).

**FARMACOTERAPIA SE NECESSÁRIO:**
Meta PA <130/80 (se risco CV alto <120/80). Se PA >140/90 ou >130/80 com RCVG ≥10% ou DM/DRC: iniciar IECA/BRA (primeira linha se <60 anos) ou BCC/tiazídico (se >60 anos ou raça negra). Monoterapia controla apenas 40-50%, maioria necessita 2-3 drogas. Titular até meta, monitorar efeitos adversos e eletrólitos."""
    },

    "obesidade": {
        "clinicalRelevance": """A obesidade familiar exibe uma das maiores herdabilidades entre condições crônicas complexas, com estudos de gêmeos estimando 40-70% de variância do IMC atribuível a fatores genéticos. A arquitetura genética é heterogênea, abrangendo formas monogênicas raras (1-5% dos casos de obesidade grave) e formas poligênicas comuns (maioria absoluta).

Obesidade monogênica severa resulta de mutações em genes do eixo leptina-melanocortina que regulam saciedade e gasto energético: LEP (leptina, hormônio da saciedade produzido por adipócitos, mutações causam hiperfagia insaciável), LEPR (receptor de leptina), POMC (pró-opiomelanocortina, precursor de α-MSH), MC4R (receptor melanocortina-4, mutações responsáveis por 2-5% obesidade severa infantil, herança autossômica dominante com penetrância variável), PC1/3 (pró-hormônio convertase), BDNF/TrkB (fator neurotrófico derivado do cérebro). Identificação de formas monogênicas é crítica pois algumas são tratáveis (deficiência de leptina responde dramaticamente a reposição recombinante).

Obesidade poligênica comum envolve >1000 variantes genéticas de pequeno efeito identificadas por GWAS, com loci principais: FTO (fat mass and obesity-associated gene, variante rs9939609 aumenta risco obesidade 1,2-1,7x por alelo, afeta 43% população europeia), MC4R, TMEM18, GNPDA2, SH2B1, MTCH2, NEGR1, BDNF. O escore de risco genético poligênico pode identificar indivíduos no quintil superior com IMC médio 3-4 kg/m² maior que quintil inferior.

Crítico: alto risco genético NÃO é determinista. O clássico estudo de Qi et al. (2012) em >20.000 indivíduos demonstrou que atividade física regular atenua significativamente o efeito genético do FTO: indivíduos fisicamente ativos com genótipo de alto risco apresentam IMC similar a inativos com genótipo de baixo risco, evidenciando potente interação gene-ambiente. Similarmente, aderência a dieta saudável reduz em 40-50% o efeito de escores poligênicos de obesidade.

Mecanismos epigenéticos cruciais: (1) metilação do DNA em genes reguladores de apetite (POMC, NPY) e adipogênese (PPARγ, LEP) modulada por nutrição materna durante gestação (programação fetal, hipótese de Barker), dieta high-fat pós-natal, e exercício; (2) modificações de histonas em tecido adiposo influenciadas por ácidos graxos dietários (palmitato induz acetilação pró-inflamatória); (3) microRNAs reguladores de adipogênese (miR-27, miR-143) e browning de tecido adiposo (miR-133, miR-196a) alterados por supernutrição.

Obesidade visceral (androide, central) é particularmente hereditária e metabolicamente deletéria, associando-se fortemente a resistência insulínica, dislipidemia aterogênica (LDL pequeno-denso, triglicerídeos elevados, HDL baixo), NASH (esteato-hepatite), HAS, e inflamação crônica de baixo grau (elevação IL-6, TNF-α, leptina, redução adiponectina). O tecido adiposo visceral funciona como órgão endócrino disfuncional, secretando adipocinas pró-inflamatórias.""",

        "conduct": """**AVALIAÇÃO ANTROPOMÉTRICA E METABÓLICA COMPLETA:**
Baseline aos 18 anos (mais cedo se sobrepeso infantil): IMC (normal 18,5-24,9, sobrepeso 25-29,9, obesidade ≥30), circunferência abdominal (risco metabólico homens >94cm, muito alto >102cm; mulheres >80cm, muito alto >88cm), relação cintura/quadril (androide se >0,9 homens, >0,85 mulheres). Composição corporal por bioimpedância ou DEXA: % gordura corporal, massa magra, gordura visceral.

Perfil metabólico: glicemia jejum, HbA1c, insulina basal, HOMA-IR (resistência insulínica), perfil lipídico avançado, hs-CRP, ALT/AST/GGT (NAFLD), TSH (hipotireoidismo), cortisol (Cushing), leptina e adiponectina (marcadores endócrinos adipócitos). Considerar teste genético de MC4R se obesidade severa infantil, ou painel poligênico se disponível e custo-efetivo.

**INTERVENÇÃO NUTRICIONAL PERSONALIZADA:**
Não existe dieta única ideal - personalizar baseado em genótipo, preferências, sustentabilidade. Abordagens evidenciadas:

(1) Dieta mediterrânea hipocalórica: rica em vegetais, frutas, grãos integrais, leguminosas, azeite, peixe, pobre em carne vermelha/processada e doces. Déficit 500-750 kcal/dia → perda 0,5-1 kg/semana. Superior a low-fat para perda peso e manutenção em estudos longos.

(2) Low-carb/cetogênica: <50-130g carboidratos/dia, aumenta gorduras saudáveis e proteína. Efetiva para perda peso (especialmente visceral) e reversão resistência insulínica em 6-12 meses. Individualizar: portadores FTO rs9939609 AA podem responder melhor a low-carb que low-fat.

(3) Jejum intermitente: 16/8 (janela alimentar 8h) ou 5:2 (2 dias/semana 500-600 kcal). Potencializa autofagia, melhora sensibilidade insulínica, preserva massa magra. Atenção contraindicações (gestantes, diabéticos em insulina).

Princípios universais: densidade calórica baixa (alimentos volume grande, calorias baixas), índice glicêmico baixo, proteína adequada 1,2-1,6 g/kg (preserva massa magra, saciedade), fibras 30-40g/dia (saciedade, microbiota), eliminar bebidas calóricas, ultraprocessados, açúcares adicionados.

**ATIVIDADE FÍSICA PARA PERDA E MANUTENÇÃO PESO:**
Exercício aeróbico 200-300min/semana intensidade moderada (necessário para manutenção pós-perda, prevenção reganho). Treinamento resistido 3-4x/semana crítico para preservar massa muscular (metabolismo basal, sensibilidade insulínica). HIIT 2-3x/semana (eficiência tempo). NEAT (termogênese atividade não-exercício): aumentar movimentos cotidianos, passos >10.000/dia.

**MODULAÇÃO COMPORTAMENTAL E PSICOLÓGICA:**
Obesidade tem componente psico-comportamental significativo: comer emocional, compulsão alimentar, distorção de porção. Terapia cognitivo-comportamental (TCC) especializada em obesidade melhora resultados 15-20%. Monitoramento de ingestão (diário alimentar, apps) aumenta sucesso. Suporte social/grupos. Atenção a transtornos alimentares (BED - binge eating disorder presente em 20-30% obesos em clínicas).

**SONO E MANEJO DE ESTRESSE:**
Privação de sono (<6h/noite) associa-se a ganho peso via desregulação hormonal (↑grelina fome, ↓leptina saciedade, ↑cortisol, resistência insulínica). Meta 7-9h/noite. Estresse crônico -> hipercortisolemia -> obesidade visceral, hyperphagia. Técnicas de redução estresse essenciais.

**FARMACOTERAPIA ADJUVANTE SE INDICADO:**
Considerar se IMC ≥30 ou ≥27 com comorbidades após 6 meses de lifestyle intensivo sem sucesso. Opções: liraglutida 3mg/dia (GLP-1 agonista, perda média 5-8%), semaglutida 2,4mg/sem (perda 12-15%, revolucionário), naltrexona/bupropiona, orlistat (menos efetivo). Setmelanotide para deficiência POMC/LEPR.

**CIRURGIA BARIÁTRICA:**
Se IMC ≥40 ou ≥35 com comorbidades graves (DM2, HAS, apneia) refratárias a tratamento clínico. Bypass gástrico Y-Roux ou gastrectomia vertical. Perda peso média 25-30%, remissão DM2 em 60-80%, mas requer acompanhamento vitalício e suplementação."""
    }
}

# Mapeamento de IDs para doenças (baseado no JSON analisado)
ITEM_DISEASE_MAPPING = {
    # Diabetes (3x)
    "4e960d9f-a04a-4524-9049-f4559170db14": "diabetes",
    "26bccb28-c694-48af-8f6a-0ce6c5f73e52": "diabetes",
    "90d067be-f3cd-48ad-9cf9-e48fd8a2027a": "diabetes",

    # Dislipidemia (3x)
    "1c15d32d-7e79-44b8-976d-8e90ecd896be": "dislipidemia",
    "4f76ef4e-1b23-47d2-bb41-549463ad3cdf": "dislipidemia",
    "a856d80b-1443-4c16-94db-1f99508b1a9c": "dislipidemia",

    # Cardiovascular (3x)
    "8fb2e7de-c341-4ff5-814f-25d7695fd1cf": "cardiovascular",
    "515052d6-c7c1-40c7-942e-a1134e3aa05e": "cardiovascular",
    "8342ba54-4c2d-4b31-8619-05f4a2e86719": "cardiovascular",

    # Renal (3x)
    "d489c80a-50f2-4606-9579-66015e62649e": "renal",
    "394414b6-0538-4162-a68c-1e1e9d8cffef": "renal",
    "307d0403-8648-431d-88cb-2ac2422f8e86": "renal",

    # Autoimune (3x)
    "62da018d-faf4-490f-9a7b-47e0f0881bbf": "autoimune",
    "f4baae49-221d-428e-a218-a29391e1e26f": "autoimune",
    "c2c1d736-45a4-45ed-be76-e9b4704d4b1d": "autoimune",

    # Viral crônica (3x)
    "8a552fbb-862c-4845-98dc-303f46922403": "viral_cronica",
    "12da9917-7311-4a87-89de-a1f0c8d918e7": "viral_cronica",
    "ac07f7de-eef0-420a-bdd9-cc0a3a41fbd8": "viral_cronica",

    # Hipertensão (3x)
    "6491db47-381b-45fe-b483-ea0183478225": "hipertensao",
    "35a7dbf8-2a43-4ee9-bc6b-c13122ed36c5": "hipertensao",
    "e257d4b5-c0b2-471b-8693-defe0b2055a3": "hipertensao",

    # Obesidade (3x)
    "77f78798-6a3a-4eef-b092-7ae1c277df4e": "obesidade",
    "3c52ad44-7049-40b8-b444-b05d15b96a57": "obesidade",
    "e2b287df-015a-4f2e-a5bf-0d99c8b24a97": "obesidade"
}


def enrich_items():
    """Enriquece os 24 items de doenças (exclui 6 parent items)"""

    # Login
    response = requests.post(
        f"{API_BASE}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )

    if response.status_code != 200:
        print(f"Login failed: {response.status_code}")
        return

    token = response.json()["accessToken"]
    print("✓ Login successful\n")

    results = []
    total = len(ITEM_DISEASE_MAPPING)

    for index, (item_id, disease_key) in enumerate(ITEM_DISEASE_MAPPING.items(), 1):
        print(f"[{index}/{total}] Processing {item_id} ({disease_key})...")

        content = DISEASE_CONTENT[disease_key]

        # Update via API
        response = requests.put(
            f"{API_BASE}/score-items/{item_id}",
            headers={
                "Authorization": f"Bearer {token}",
                "Content-Type": "application/json"
            },
            json=content
        )

        if response.status_code == 200:
            print(f"  ✓ Successfully updated")
            results.append({"id": item_id, "disease": disease_key, "success": True})
        else:
            print(f"  ✗ Failed: {response.status_code} - {response.text}")
            results.append({"id": item_id, "disease": disease_key, "success": False, "error": response.text})

        # Rate limiting
        if index < total:
            time.sleep(1)

    # Summary
    print("\n" + "=" * 60)
    print("SUMMARY")
    print("=" * 60)
    successful = [r for r in results if r["success"]]
    failed = [r for r in results if not r["success"]]

    print(f"\nTotal items: {total}")
    print(f"Successful: {len(successful)}")
    print(f"Failed: {len(failed)}")

    if failed:
        print("\nFailed items:")
        for r in failed:
            print(f"  - {r['id']} ({r['disease']})")

    # Save results
    output_file = "/home/user/plenya/historico_familiar_enrichment_results.json"
    with open(output_file, "w", encoding="utf-8") as f:
        json.dump(results, f, indent=2, ensure_ascii=False)

    print(f"\n✓ Results saved to: {output_file}")


if __name__ == "__main__":
    enrich_items()
