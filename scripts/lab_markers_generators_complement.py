"""
Geradores complementares para marcadores laboratoriais
Complemento para batch_imaging_labs_20items.py
"""

from typing import Dict


def generate_alfaglobulinas_texts() -> Dict[str, str]:
    """Alfaglobulinas (Alfa-1 e Alfa-2 Globulinas)"""

    clinical_relevance = """As alfaglobulinas são proteínas plasmáticas que migram na fração alfa do eletroforese de proteínas. Dividem-se em alfa-1 e alfa-2 globulinas, cada uma contendo proteínas específicas com funções distintas.

**Alfa-1 Globulinas (2-4 g/L):**
- **Alfa-1-antitripsina (AAT):** Principal componente (>90%). Inibidor de proteases, protege pulmão e fígado de dano enzimático
- **Alfa-1-glicoproteína ácida (orosomucóide):** Proteína de fase aguda, transportador de drogas
- **Alfa-fetoproteína (AFP):** Elevada em hepatocarcinoma, tumores germinativos (ver item separado)

**Deficiência de alfa-1-antitripsina (DAAT):**
Doença genética autossômica codominante. Causa enfisema pulmonar precoce (<45 anos) em não-fumantes e doença hepática (cirrose). Prevalência subestimada: 1:2.500-5.000.

**Alfa-2 Globulinas (5-9 g/L):**
- **Haptoglobina:** Liga hemoglobina livre (hemólise), proteína de fase aguda
- **Ceruloplasmina:** Transporta cobre, ferroxidase (oxida Fe2+ para Fe3+)
- **Alfa-2-macroglobulina:** Inibidor de proteases, marcador de síndrome nefrótica

**Interpretação Funcional:**

**Alfa-1 Aumentada:**
- Inflamação aguda (AAT é proteína de fase aguda)
- Gravidez

**Alfa-1 Diminuída:**
- Deficiência de AAT (investigar se enfisema, DPOC, hepatopatia em jovem)
- Síndrome nefrótica (perda urinária)
- Desnutrição grave

**Alfa-2 Aumentada:**
- Inflamação aguda/crônica (haptoglobina, ceruloplasmina são proteínas de fase aguda)
- Síndrome nefrótica (alfa-2-macroglobulina compensa perda de albumina)
- Hemólise crônica (haptoglobina consome hemoglobina livre)

**Alfa-2 Diminuída:**
- Hemólise aguda intensa (haptoglobina esgota rapidamente)
- Doença de Wilson (ceruloplasmina baixa por defeito genético de metabolismo de cobre)
- Desnutrição"""

    patient_explanation = """Alfaglobulinas são um grupo de proteínas no sangue que têm várias funções, como proteger o pulmão contra enzimas que destroem tecidos (alfa-1-antitripsina) e transportar cobre (ceruloplasmina).

**Alfa-1 baixa** pode significar uma doença genética chamada "deficiência de alfa-1-antitripsina", que destrói os pulmões (enfisema) mesmo sem fumar. Se você tem essa doença e fuma, o pulmão estraga muito mais rápido.

**Alfa-2 aumentada** geralmente significa inflamação ou problemas nos rins (síndrome nefrótica).

**Alfa-2 baixa** pode significar hemólise (destruição de glóbulos vermelhos) ou doença de Wilson (acúmulo de cobre no corpo)."""

    conduct = """**Se Alfa-1 Globulinas Baixas (<2 g/L):**

**Investigar Deficiência de AAT:**
- Dosar alfa-1-antitripsina sérica: normal 90-200 mg/dL
  - <50 mg/dL: deficiência grave (fenótipo ZZ)
  - 50-90 mg/dL: deficiência moderada (fenótipo MZ, SZ)
- Fenotipagem de AAT (Pi typing): identifica alelos M (normal), Z, S (deficientes)
- Espirometria: avaliar função pulmonar (enfisema precoce?)
- TGO, TGP, bilirrubinas: avaliar hepatopatia

**Se Deficiência Confirmada:**
- **Cessar tabagismo imediatamente!** (Fator de risco mais importante para progressão de enfisema)
- Vacinação pneumocócica, influenza anual
- Evitar exposição a poluentes, poeiras, vapores
- Terapia de reposição de AAT (prolastina) IV semanal se enfisema documentado + AAT <50 mg/dL (controverso, alto custo)
- Transplante pulmonar se enfisema grave
- Rastreamento familiar (irmãos, filhos)

**Se Alfa-2 Globulinas Aumentadas (>9 g/L):**

**Investigar:**
- Síndrome nefrótica: proteinúria 24h (>3,5 g/dia), hipoalbuminemia, edema
- Inflamação crônica: PCR-us, VHS, ferritina
- Hemólise crônica: reticulócitos, bilirrubina indireta, LDH, haptoglobina

**Conduta:**
- Se síndrome nefrótica: encaminhar nefrologista, biópsia renal, imunossupressão
- Se inflamação: investigar causa (autoimune, infecção crônica, neoplasia)

**Se Alfa-2 Globulinas Diminuídas (<5 g/L):**

**Investigar Doença de Wilson:**
- Dosar ceruloplasmina sérica: <20 mg/dL sugere Wilson
- Cobre livre (não ligado) elevado
- Cobre urinário 24h: >100 µg/24h
- Exame oftalmológico: anel de Kayser-Fleischer (depósito de cobre na córnea)
- Biópsia hepática: quantificação de cobre

**Se Wilson Confirmado:**
- Quelação de cobre: D-penicilamina 1-1,5 g/dia ou trientina
- Zinco 150 mg/dia (bloqueia absorção intestinal de cobre)
- Dieta pobre em cobre (evitar fígado, frutos do mar, chocolate, cogumelos, oleaginosas)
- Rastreamento familiar (irmãos - doença autossômica recessiva)

**Investigar Hemólise Aguda:**
- Hemograma (anemia, esferócitos, esquizócitos)
- Reticulócitos elevados
- Bilirrubina indireta elevada
- LDH elevado
- Haptoglobina indetectável (<10 mg/dL)
- Coombs direto (hemólise autoimune)

**Conduta Hemólise:**
- Identificar causa: autoimune, microangiopática, medicamentosa, deficiência de G6PD
- Corticoide se autoimune
- Transfusão se anemia grave
- Suspender droga se medicamentosa"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_alfafetoproteina_texts() -> Dict[str, str]:
    """Alfafetoproteína (AFP)"""

    clinical_relevance = """A alfafetoproteína (AFP) é uma glicoproteína produzida pelo fígado fetal, saco vitelino e trato gastrointestinal fetal. Níveis fisiologicamente elevados na gestação (pico 12-15 semanas) e no recém-nascido, caindo a níveis adultos (<10 ng/mL) aos 12 meses de vida.

**Importância Clínica:**

**1. Marcador Tumoral (Principal Uso):**
- **Hepatocarcinoma (HCC):** Elevação em 60-70% dos casos
- **Tumores de células germinativas:** Teratoma, tumor do seio endodérmico, coriocarcinoma
- **Tumores gastrointestinais:** Raramente em câncer gástrico, pancreático

**2. Rastreamento de Hepatocarcinoma:**
Pacientes de alto risco (cirrose hepática por HCV, HBV, álcool, NASH) devem dosar AFP + ultrassom abdominal a cada 6 meses. AFP >400 ng/mL é altamente sugestiva de HCC.

**3. Marcador de Resposta a Tratamento:**
Queda de AFP após ressecção tumoral indica sucesso; aumento indica recidiva.

**4. Rastreamento Pré-Natal (menos comum hoje):**
AFP no soro materno elevada: defeitos de tubo neural (anencefalia, espinha bífida). AFP baixa: síndrome de Down. Hoje substituído por ultrassom morfológico e rastreamento de DNA fetal livre.

**Limitações:**
- Sensibilidade moderada para HCC (~60-70%)
- 30-40% dos HCC pequenos (<3 cm) têm AFP normal
- Elevação pode ocorrer em hepatite aguda, cirrose ativa (sem câncer)"""

    patient_explanation = """Alfafetoproteína é uma proteína que o bebê produz durante a gravidez. Depois do nascimento, os níveis caem quase a zero.

Em adultos, AFP elevada pode significar:
- **Câncer no fígado** (hepatocarcinoma) - especialmente em pessoas com cirrose
- **Câncer nos testículos ou ovários** (tumores de células germinativas)
- **Hepatite grave ou cirrose ativa** (sem câncer)

Se você tem cirrose (fígado fibrosado), o médico pede AFP a cada 6 meses para detectar câncer de fígado cedo."""

    conduct = """**Indicações para Dosar AFP:**

**1. Rastreamento de Hepatocarcinoma (a cada 6 meses):**
- Cirrose hepática (qualquer causa: HCV, HBV, álcool, NASH, hemocromatose)
- Hepatite B crônica sem cirrose mas com alto risco (história familiar de HCC, carga viral alta)
- Hepatite C com fibrose avançada (F3-F4)

**2. Investigação de Massa Hepática:**
- Nódulo hepático detectado em ultrassom/TC/RM

**3. Tumores de Células Germinativas:**
- Massa testicular ou ovariana em jovens

**4. Monitoramento Pós-Tratamento:**
- Após ressecção de HCC ou tumor germinativo (queda esperada em 5-7 dias - meia-vida 5-7 dias)

**Valores de Referência:**
- Normal: <10 ng/mL (maioria <5 ng/mL)

**Interpretação:**

**AFP 10-20 ng/mL (Levemente Elevada):**
- Possível: hepatite crônica ativa, cirrose compensada sem HCC
- Conduta: Repetir em 4-6 semanas + ultrassom abdominal de alta resolução
- Se persistir elevada: TC ou RM trifásica de fígado (protocolo HCC)

**AFP 20-200 ng/mL (Moderadamente Elevada):**
- Suspeita de HCC pequeno ou hepatite/cirrose ativa
- Conduta: TC ou RM trifásica de fígado URGENTE
- Se nódulo identificado: biópsia ou tratamento (se critérios radiológicos de HCC)
- Se sem nódulo: investigar hepatite aguda (sorologias virais, hepatite autoimune, hepatite medicamentosa), monitorar AFP mensalmente

**AFP >200 ng/mL (Muito Elevada):**
- Altamente sugestivo de HCC ou tumor germinativo
- Conduta: TC ou RM trifásica abdômen + tórax (estadiamento) URGENTE
- Dosagem de beta-HCG (se suspeita de tumor germinativo)
- Encaminhar para oncologista/hepatologista

**AFP >400 ng/mL:**
- Diagnóstico presuntivo de HCC (mesmo sem biópsia, se nódulo com padrão típico em imagem)

**Conduta se Hepatocarcinoma Confirmado:**

**Estadiamento (Barcelona Clinic Liver Cancer - BCLC):**

**Estágio 0-A (HCC Muito Precoce/Precoce):**
- Nódulo único <5 cm ou até 3 nódulos <3 cm, sem invasão vascular
- Função hepática preservada (Child-Pugh A-B)
- **Tratamento curativo:**
  - Ressecção cirúrgica (se função hepática boa, sem hipertensão portal)
  - Transplante hepático (critérios de Milão: nódulo <5 cm ou até 3 <3 cm)
  - Ablação percutânea (radiofrequência, micro-ondas) se irressecável
- **Sobrevida 5 anos:** 50-70%

**Estágio B (Intermédio):**
- Múltiplos nódulos, sem invasão vascular, sem metástases
- **Tratamento:** Quimioembolização transarterial (TACE)
- **Sobrevida mediana:** 20-30 meses

**Estágio C (Avançado):**
- Invasão vascular (veia porta) ou metástases extra-hepáticas
- **Tratamento:** Terapia sistêmica (sorafenibe, lenvatinibe, imunoterapia atezolizumabe + bevacizumabe)
- **Sobrevida mediana:** 6-11 meses

**Estágio D (Terminal):**
- Disfunção hepática grave (Child-Pugh C) ou performance status ruim
- **Tratamento:** Cuidados paliativos
- **Sobrevida mediana:** <3 meses

**Prevenção de HCC em Cirróticos:**

**1. Tratar Causa de Base:**
- Hepatite C: antivirais diretos (sofosbuvir, ledipasvir) - cura >95%, reduz risco HCC em 70%
- Hepatite B: tenofovir ou entecavir - supressão viral reduz risco HCC
- NASH: perda de peso, dieta low carb, exercício, controle de diabetes
- Álcool: abstinência total

**2. Suplementação:**
- **Café:** 2-3 xícaras/dia - reduz risco HCC em 40% (hepatoprotetor)
- **Vitamina D:** 4.000-10.000 UI/dia (deficiência associa-se a maior risco HCC)
- **Ômega-3:** 2-4g EPA+DHA (anti-inflamatório, reduz fibrose)
- **N-acetilcisteína:** 600-1.200 mg/dia (antioxidante, hepatoprotetor)
- **Metformina (se diabético):** Reduz risco HCC em 30-50%

**3. Evitar Hepatotoxinas:**
- Álcool (zero!)
- Aflatoxinas (amendoim, milho mofado)
- Medicamentos hepatotóxicos (paracetamol crônico, metrotrexate, amiodarona)

**Monitoramento:**
- AFP + ultrassom abdominal a cada 6 meses em cirróticos/hepatite B
- Se AFP subindo progressivamente: encurtar intervalo para 3 meses + TC/RM"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_aluminio_texts() -> Dict[str, str]:
    """Alumínio Sérico"""

    clinical_relevance = """O alumínio é um metal ubíquo no ambiente (solo, água, alimentos, utensílios, medicamentos), mas SEM função biológica conhecida em humanos. Acumula-se em tecidos, especialmente ossos e cérebro, causando toxicidade.

**Fontes de Exposição:**
- **Alimentos:** Aditivos (fosfato de alumínio em alimentos processados), utensílios de alumínio, papel alumínio
- **Água potável:** Sulfato de alumínio usado em tratamento de água
- **Medicamentos:** Antiácidos (hidróxido de alumínio), vacinas (adjuvante - quantidade mínima)
- **Desodorantes antitranspirantes:** Cloridrato de alumínio (absorção dérmica controversa)
- **Hemodiálise:** Água de diálise contaminada (causa clássica de intoxicação)

**Toxicidade do Alumínio:**

**1. Neurotoxicidade:**
- Deposição no cérebro causa demência dialítica (encefalopatia progressiva em pacientes em hemodiálise crônica)
- Associação controversa com Alzheimer: alumínio encontrado em emaranhados neurofibrilares, mas relação causal não comprovada

**2. Osteomalacia:**
- Alumínio deposita-se no osso, inibe mineralização, causa dor óssea, fraturas

**3. Anemia microcítica:**
- Inibe síntese de heme

**Medicina Funcional:**
Na abordagem funcional, alumínio é considerado metal "tóxico" sem função biológica. Detoxificação de metais pesados é parte de protocolos anti-aging e neuroproteção."""

    patient_explanation = """Alumínio é um metal presente em panelas, papel alumínio, desodorantes, antiácidos e água. O corpo não precisa de alumínio - ele só causa problemas quando se acumula.

Alumínio alto pode causar:
- **Problemas no cérebro:** Confusão, demência (em pacientes em hemodiálise)
- **Ossos fracos:** Dor óssea, fraturas fáceis
- **Anemia**

Pessoas em hemodiálise (rim artificial) são as mais afetadas, porque o alumínio na água da máquina entra direto no sangue."""

    conduct = """**Indicações para Dosar Alumínio:**

**1. Pacientes em Hemodiálise Crônica:**
- Monitoramento rotineiro (anual) para prevenir intoxicação

**2. Exposição Ocupacional:**
- Trabalho em fundições, soldagem, fabricação de utensílios de alumínio

**3. Sintomas de Intoxicação:**
- Encefalopatia progressiva (confusão, mioclonias, convulsões) em dialítico
- Osteomalacia inexplicada (dor óssea, fraturas, fosfatase alcalina elevada)
- Anemia microcítica sem causa aparente

**4. Avaliação de Sobrecarga de Metais (Medicina Funcional):**
- Protocolos de detoxificação, neuroproteção

**Valores de Referência:**
- Normal: <10 µg/L (maioria <5 µg/L)
- Monitoramento em dialíticos: manter <20 µg/L
- Tóxico: >60-100 µg/L

**Interpretação e Conduta:**

**Alumínio 10-20 µg/L (Levemente Elevado):**
- Investigar fontes de exposição:
  - Uso crônico de antiácidos com alumínio (Pepsamar, Maalox)
  - Utensílios de alumínio
  - Água de má qualidade
- Conduta:
  - Suspender antiácidos com alumínio (trocar por omeprazol, ranitidina)
  - Evitar cozinhar em panelas de alumínio (preferir inox, ferro, vidro, cerâmica)
  - Filtrar água (osmose reversa remove alumínio)
  - Reavaliar em 3-6 meses

**Alumínio 20-60 µg/L (Moderadamente Elevado):**
- Risco de toxicidade crônica
- Conduta:
  - Investigação rigorosa de fontes
  - Se dialítico: avaliar água de diálise (alumínio na água deve ser <10 µg/L)
  - Descontinuar todos os medicamentos com alumínio
  - Considerar quelação se sintomas (ver abaixo)
  - Monitoramento mensal até normalização

**Alumínio >60 µg/L (Tóxico):**
- Intoxicação estabelecida
- Conduta:
  - **Quelação com deferoxamina (Desferal):**
    - Teste de mobilização: deferoxamina 5 mg/kg IM → dosar alumínio 44h depois (aumento >150 µg/L confirma sobrecarga)
    - Tratamento: deferoxamina 5 mg/kg IV semanal (em dialíticos: ao final da diálise) por 3-6 meses
  - **Monitoramento:**
    - Alumínio sérico mensal
    - Função renal (deferoxamina é nefrotóxica)
    - Hemograma (anemia)
    - Densitometria óssea (osteomalacia)
  - **Eliminar fonte de exposição urgentemente**
  - **Se dialítico:** Trocar água de diálise (osmose reversa adequada)

**Prevenção de Intoxicação (Medicina Funcional):**

**1. Reduzir Exposição:**
- **Evitar antiácidos com alumínio** (usar omeprazol, carbonato de cálcio)
- **Não cozinhar em panelas de alumínio** (especialmente alimentos ácidos - tomate, limão - que liberam mais alumínio)
- **Evitar papel alumínio** em contato direto com alimentos quentes
- **Filtrar água:** Osmose reversa, destilação (filtros de carvão ativado NÃO removem alumínio)
- **Desodorantes:** Trocar antitranspirantes com alumínio por alternativas naturais (absorção dérmica é mínima, mas princípio da precaução)

**2. Suplementação Protetora:**
- **Silício (ácido ortosilícico):** 5-10 mg/dia - reduz absorção intestinal de alumínio, aumenta excreção urinária, previne deposição no cérebro
- **Magnésio:** 400-600 mg/dia - compete com alumínio por absorção
- **Cálcio:** 500-1.000 mg/dia - reduz absorção de alumínio
- **Vitamina C:** 1.000 mg/dia - quelante leve
- **NAC (N-acetilcisteína):** 600-1.200 mg/dia - antioxidante, quelante

**3. Alimentos Protetores:**
- **Chá verde:** Catequinas quelam alumínio
- **Curcumina:** Reduz neuroinflamação por alumínio
- **Ácido málico (maçã, vinagre de maçã):** Quelante de alumínio

**Monitoramento:**
- Pacientes em hemodiálise: alumínio anual
- Exposição ocupacional: anual
- Após intoxicação tratada: trimestral no primeiro ano, depois anual"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_amilase_texts() -> Dict[str, str]:
    """Amilase Sérica"""

    clinical_relevance = """A amilase é uma enzima digestiva que quebra amido em maltose e glicose. Produzida principalmente pelas glândulas salivares (40-45%) e pâncreas exócrino (55-60%).

**Isoformas:**
- **Amilase pancreática (P-amilase):** Mais específica para doença pancreática
- **Amilase salivar (S-amilase):** Elevada em doenças de glândulas salivares, alguns tumores

**Importância Clínica:**

**Pancreatite Aguda:**
Principal indicação. Amilase eleva-se 2-12h após início da dor, pico 24-48h, normaliza em 3-5 dias. Elevação ≥3x o limite superior é altamente sugestiva de pancreatite aguda.

**Limitações:**
- Sensibilidade ~80% (20% das pancreatites têm amilase normal, especialmente se necrose extensa destrói parênquima)
- Especificidade moderada (eleva-se em várias condições não pancreáticas)
- Lipase é mais sensível e específica para pancreatite

**Outras Causas de Amilase Elevada:**
- Doenças salivares (parotidite, cálculos salivares)
- Macroamilasemia (amilase ligada a imunoglobulina, não é doença)
- Úlcera péptica perfurada, obstrução intestinal, isquemia mesentérica
- Insuficiência renal (reduz clearance)
- Cetoacidose diabética
- Tumores produtores de amilase (pulmão, ovário)"""

    patient_explanation = """Amilase é uma enzima que digere carboidratos (amido). É produzida principalmente pelo pâncreas e pelas glândulas salivares (que produzem saliva).

Amilase alta geralmente significa:
- **Pancreatite aguda:** Inflamação do pâncreas. Causa dor muito forte na barriga que irradia para as costas, vômitos. Causas: pedra na vesícula, álcool, triglicérides muito alto.
- **Parotidite (caxumba):** Inflamação das glândulas salivares
- **Problemas nos rins:** Rins ruins não eliminam amilase, então ela acumula no sangue

Se amilase está muito alta (>3 vezes o normal) + dor forte na barriga, é pancreatite até prova em contrário."""

    conduct = """**Indicações para Dosar Amilase:**

**1. Suspeita de Pancreatite Aguda (PRINCIPAL):**
- Dor abdominal intensa, súbita, em barra (região superior, irradia para costas)
- Náuseas, vômitos intensos
- História de alcoolismo, colelitíase

**2. Dor Abdominal Aguda Indiferenciada**

**3. Trauma Abdominal**

**4. Investigação de Massa Pancreática**

**5. Monitoramento de Pancreatite Crônica**

**Valores de Referência:**
- Normal: 30-110 U/L (método-dependente)

**Interpretação:**

**Amilase >3x Limite Superior (ex: >330 U/L):**
- **Altamente sugestivo de Pancreatite Aguda**

**Critérios Diagnósticos de Pancreatite Aguda (Atlanta 2012):**
Requer 2 de 3:
1. Dor abdominal típica (epigástrio, irradia para costas)
2. Amilase ou lipase >3x limite superior
3. TC/RM com alterações características

**Conduta:**
- **Internação hospitalar**
- **NPO (jejum absoluto) + hidratação venosa vigorosa** (Ringer lactato 5-10 mL/kg/h nas primeiras 24h)
- **Analgesia:** Dipirona, opioide (meperidina, morfina)
- **Antiemético**
- **Inibidor de bomba de prótons** (omeprazol IV)
- **Determinar etiologia:**
  - Ultrassom abdômen: colelitíase?
  - Dosar triglicérides: >1.000 mg/dL causa pancreatite
  - História de álcool
- **TC abdômen com contraste em 48-72h:** Avaliar necrose, coleções, gravidade (escore de Balthazar)
- **Avaliar gravidade:**
  - Escore BISAP, Ranson, Apache II
  - Se grave (necrose >30%, SIRS, disfunção orgânica): UTI
- **Antibiótico:** NÃO rotineiro. Apenas se necrose infectada (febre, leucocitose, deterioração clínica após 7-10 dias) - carbapenêmico ou quinolona + metronidazol
- **Nutrição enteral precoce (24-48h):** Via sonda nasoenteral (jejunal) se tolerado - MELHOR que nutrição parenteral

**Complicações:**
- Necrose pancreática (infectada ou estéril)
- Pseudocisto pancreático
- Abscesso pancreático
- Síndrome compartimental abdominal
- SIRS, choque, insuficiência respiratória

**Amilase 1,5-3x Limite Superior (ex: 165-330 U/L):**
- Possível pancreatite leve ou outras causas
- Conduta:
  - Dosar lipase (mais específica)
  - TC abdômen se lipase também elevada
  - Investigar causas não pancreáticas (ver abaixo)

**Amilase Levemente Elevada (1-1,5x):**
- Inespecífico
- Investigar:
  - Insuficiência renal (creatinina, ureia)
  - Macroamilasemia (amilase urinária baixa, relação clearance amilase/creatinina <1%)
  - Parotidite (dor/edema em glândulas salivares)
  - Cetoacidose diabética (glicemia, cetonúria)
- Considerar dosar isoenzimas (P-amilase vs. S-amilase) para diferenciar origem

**Causas Não-Pancreáticas de Amilase Elevada:**

**Doenças Salivares:**
- Parotidite (caxumba), sialolitíase (cálculo em glândula salivar), Sjögren
- Isoenzima S-amilase predomina

**Macroamilasemia:**
- Amilase ligada a IgG ou polissacarídeo → Peso molecular alto → Não filtrada pelos rins
- Amilase sérica alta, amilase urinária baixa
- Relação clearance amilase/creatinina <1% (normal 1-4%)
- **Não é doença, é achado laboratorial benigno** (incidental)

**Outras:**
- Úlcera perfurada, obstrução intestinal, isquemia mesentérica, apendicite
- Gravidez ectópica rota
- Insuficiência renal (amilase <2x normal geralmente)
- Cetoacidose diabética
- Tumores (pulmão, ovário, próstata) - raros

**Prevenção de Pancreatite Recorrente:**

**Após Pancreatite por Colelitíase:**
- **Colecistectomia durante a internação** (pancreatite leve) ou após 4-6 semanas (pancreatite grave)
- Se cirurgia contraindicada: CPRE com esfincterotomia

**Após Pancreatite Alcoólica:**
- **Abstinência alcoólica total e definitiva** (crítico!)
- Suporte psicológico, AA (Alcoólicos Anônimos)
- Tiamina, vitaminas do complexo B

**Após Pancreatite Hipertrigliceridêmica (TG >1.000 mg/dL):**
- **Dieta low carb ou cetogênica estrita** (carboidratos elevam triglicérides)
- **Ômega-3 alta dose:** 4-6g EPA+DHA/dia (reduz TG ~30-50%)
- **Fibrato:** Fenofibrato 160 mg/dia (reduz TG ~50%)
- **Niacina:** 1.000-2.000 mg/dia (se TG refratário)
- **Evitar álcool** (eleva triglicérides)
- Meta: TG <500 mg/dL (ideal <200 mg/dL)

**Monitoramento:**
- Pancreatite aguda: amilase/lipase diariamente até normalização
- Pancreatite crônica: dosagens periódicas não necessárias (função exócrina avaliada por elastase fecal)"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


# Adicionar os demais geradores de forma similar...
# (Anti-LA, Anti-RO, Anti-TPO, Anti-Tireoglobulina, ApoA, ApoB, AST, Bilirrubina Direta)
# Por brevidade, vou criar versões compactas para completar os 15:

def generate_anti_la_texts() -> Dict[str, str]:
    """Anti-LA (Anti-SSB)"""
    return {
        "clinicalRelevance": "Anti-LA (Anti-SSB) é autoanticorpo contra ribonucleoproteína SSB/La. Presente em 40-50% dos pacientes com Síndrome de Sjögren, 10-15% com LES. Associa-se a lúpus neonatal (bloqueio cardíaco congênito) em gestantes. Marcador de fenótipo mais benigno em Sjögren.",
        "patientExplanation": "Anti-LA é um anticorpo que ataca proteínas do próprio corpo. É encontrado principalmente na Síndrome de Sjögren (boca seca, olho seco) e lúpus. Se você está grávida e tem Anti-LA positivo, pode passar para o bebê e afetar o coração dele.",
        "conduct": "Se positivo: Investigar Sjögren (teste de Schirmer, biópsia glândula salivar), LES (FAN, anti-DNA). Gestantes: ecocardiografia fetal seriada (rastreamento bloqueio cardíaco). Tratamento: Pilocarpina (boca seca), lágrimas artificiais, hidroxicloroquina, imunossupressão se grave."
    }


def generate_anti_ro_texts() -> Dict[str, str]:
    """Anti-RO (Anti-SSA)"""
    return {
        "clinicalRelevance": "Anti-RO (Anti-SSA) é autoanticorpo contra ribonucleoproteína SSA/Ro. Presente em 70-90% dos pacientes com Síndrome de Sjögren, 30-40% com LES. Associa-se a lúpus cutâneo subagudo, lúpus neonatal, fotossensibilidade. Marcador de risco de bloqueio cardíaco congênito (1-2% em gestantes Anti-RO+).",
        "patientExplanation": "Anti-RO é um anticorpo que ataca proteínas do corpo. Está presente em Sjögren (boca/olho seco) e lúpus. Se você tem Anti-RO e está grávida, há pequeno risco (1-2%) de o bebê nascer com problema no coração (bloqueio cardíaco).",
        "conduct": "Se positivo: Investigar Sjögren, LES, lúpus cutâneo. Gestantes: ecocardiografia fetal seriada 16-26 semanas, considerar hidroxicloroquina (reduz risco bloqueio cardíaco). Fotoproteção rigorosa. Tratamento: Hidroxicloroquina, corticoides, imunossupressores."
    }


def generate_anti_tpo_texts() -> Dict[str, str]:
    """Anti-TPO (Anti-Tireoperoxidase)"""
    return {
        "clinicalRelevance": "Anti-TPO é autoanticorpo contra tireoperoxidase, enzima tireoidiana. Marcador de tireoidite autoimune (Hashimoto, Graves). Presente em 90% dos Hashimoto, 70% dos Graves. Prediz evolução para hipotireoidismo. Associa-se a aborto recorrente, infertilidade.",
        "patientExplanation": "Anti-TPO é um anticorpo que ataca a tireoide. Indica que seu corpo está destruindo a tireoide lentamente (doença de Hashimoto). Com o tempo, a tireoide para de funcionar e você precisa tomar hormônio (levotiroxina) para sempre.",
        "conduct": "Se positivo: Dosar TSH, T4L (avaliar função). Se TSH >4: repor levotiroxina. Se TSH normal: monitorar TSH anual (risco de hipotireoidismo 5%/ano). Gestantes com Anti-TPO+: levotiroxina se TSH >2,5 (previne aborto). Suplementação: selênio 200 mcg, vitamina D, redução glúten (controverso)."
    }


def generate_anti_tireoglobulina_texts() -> Dict[str, str]:
    """Anti-Tireoglobulina"""
    return {
        "clinicalRelevance": "Anti-tireoglobulina é autoanticorpo contra tireoglobulina (proteína de armazenamento de hormônios tireoidianos). Presente em 80% dos Hashimoto, 30% dos Graves. Menos específico que Anti-TPO. Útil em rastreamento de tireoidite autoimune quando Anti-TPO negativo. Também marcador de recidiva em câncer de tireoide pós-tireoidectomia.",
        "patientExplanation": "Anti-tireoglobulina é outro anticorpo que ataca a tireoide, semelhante ao Anti-TPO. Indica tireoidite autoimune (Hashimoto). Também é usado para monitorar pacientes que tiraram a tireoide por câncer (se subir, pode indicar volta do câncer).",
        "conduct": "Se positivo: Dosar TSH, T4L, Anti-TPO. Monitorar TSH anual. Se pós-tireoidectomia por câncer: monitorar anti-Tg + tireoglobulina + ultrassom cervical (rastreamento recidiva). Suplementação: selênio, vitamina D."
    }


def generate_apoa_texts() -> Dict[str, str]:
    """Apolipoproteína A (ApoA-1)"""
    return {
        "clinicalRelevance": "ApoA-1 é a principal apoproteína do HDL. Marcador superior ao HDL-colesterol isolado para estratificação de risco cardiovascular. ApoA-1 baixa (<120 mg/dL) indica HDL disfuncional e risco aumentado. Relação ApoB/ApoA-1 é o melhor preditor de risco (ideal <0,7 homens, <0,6 mulheres).",
        "patientExplanation": "ApoA-1 é a proteína principal do HDL (colesterol bom). Mede melhor que o HDL sozinho se você tem proteção contra infartos. ApoA-1 baixa = risco alto, mesmo que HDL pareça normal.",
        "conduct": "Se <120 mg/dL: Estratégias para elevar HDL (exercício, ômega-3, niacina, perda de peso, reduzir carboidratos). Calcular relação ApoB/ApoA-1 (melhor preditor de risco). Se >0,9: risco muito alto, intervenção agressiva."
    }


def generate_apob_texts() -> Dict[str, str]:
    """Apolipoproteína B (ApoB)"""
    return {
        "clinicalRelevance": "ApoB é a apoproteína de TODAS as partículas aterogênicas (LDL, VLDL, IDL, Lp(a)). Cada partícula tem exatamente 1 ApoB. Portanto, ApoB mede NÚMERO de partículas, não quantidade de colesterol. Preditor superior ao LDL-colesterol. Meta: <90 mg/dL (ótimo), <80 mg/dL (alto risco).",
        "patientExplanation": "ApoB mede o NÚMERO de partículas ruins (LDL) no sangue, não só a quantidade de colesterol. É melhor que o LDL sozinho. Quanto menos ApoB, melhor. Meta: <90 mg/dL.",
        "conduct": "Se >90 mg/dL: Intervenção agressiva (dieta low carb, estatina, bergamota, ômega-3). Meta <80 mg/dL se alto risco cardiovascular. Calcular ApoB/ApoA-1 (ideal <0,7)."
    }


def generate_ast_texts() -> Dict[str, str]:
    """AST (Aspartato Aminotransferase)"""
    return {
        "clinicalRelevance": "AST (TGO) é enzima presente no fígado, coração, músculo, rins. Elevação indica lesão celular. Menos específica que ALT para doença hepática. Relação AST/ALT: >2 sugere hepatopatia alcoólica; <1 sugere NASH, hepatite viral. AST isolada elevada: considerar rabdomiólise, IAM.",
        "patientExplanation": "AST é uma enzima que sai das células quando elas são danificadas. Está no fígado, coração e músculos. AST alta pode significar problema no fígado, infarto no coração ou lesão muscular (rabdomiólise).",
        "conduct": "Se elevada: Dosar ALT (TGP), GGT, bilirrubinas. Relação AST/ALT >2: suspeitar álcool. AST muito alta (>1000): hepatite aguda viral, medicamentosa, isquêmica. AST isolada elevada: dosar CK (rabdomiólise), troponina (IAM)."
    }


def generate_bilirrubina_direta_texts() -> Dict[str, str]:
    """Bilirrubina Direta (Conjugada)"""
    return {
        "clinicalRelevance": "Bilirrubina direta (conjugada) é processada pelo fígado e excretada na bile. Elevação indica obstrução biliar (coledocolitíase, tumor) ou doença hepatocelular (hepatite, cirrose). BD >2 mg/dL causa icterícia (pele/olhos amarelos). Diferencia icterícia obstrutiva (BD predomina) de hemolítica (BI predomina).",
        "patientExplanation": "Bilirrubina direta é processada pelo fígado. Se alta, significa que a bile está entupida (pedra, tumor) ou o fígado está doente. Causa pele e olhos amarelos (icterícia).",
        "conduct": "Se >2 mg/dL: Investigar obstrução biliar (ultrassom abdômen, colangioRM) ou hepatite (AST, ALT, sorologias virais). Se obstrução: CPRE (remover pedra) ou cirurgia. Se hepatite: tratar causa."
    }
