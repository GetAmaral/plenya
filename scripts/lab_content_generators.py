"""
Geradores de conteúdo clínico para Score Items laboratoriais
Medicina Funcional Integrativa
"""

from typing import Dict


def generate_colesterol_total_texts() -> Dict[str, str]:
    """Gera textos para Colesterol Total"""

    clinical_relevance = """O colesterol total representa a soma de todas as frações de colesterol circulante: LDL, HDL, VLDL e IDL. É uma molécula essencial à vida, precursora de hormônios esteroidais (testosterona, estrogênio, progesterona, cortisol), ácidos biliares, vitamina D e componente estrutural de todas as membranas celulares, incluindo a bainha de mielina cerebral.

Na medicina funcional integrativa, o paradigma de "quanto menor, melhor" para colesterol total está ultrapassado. Estudos epidemiológicos mostram que a mortalidade total apresenta curva em U: tanto valores muito baixos (<160 mg/dL) quanto muito elevados (>300 mg/dL) associam-se a maior mortalidade. A faixa de menor mortalidade situa-se entre 200-240 mg/dL, desafiando diretrizes convencionais.

Colesterol total isolado é marcador insuficiente para estratificação de risco cardiovascular. A medicina funcional prioriza:

1. **Qualidade das partículas:** Subtipos de LDL (partículas pequenas e densas vs. grandes e flutuantes), subtipos de HDL (HDL grande e funcional vs. pequeno e disfuncional)

2. **Razões e marcadores avançados:** ApoB/ApoA1 (preditor superior ao colesterol isolado), relação triglicerídeos/HDL (<2 ideal), colesterol não-HDL, LDL oxidado (anti-LDL oxidado)

3. **Contexto metabólico:** Resistência insulínica, inflamação (PCR-us), perfil de ácidos graxos, estresse oxidativo

4. **Carga aterosclerótica real:** Escore de cálcio coronariano (CAC) - pacientes com CAC=0 têm risco cardiovascular muito baixo mesmo com LDL elevado (estudo "Power of Zero")

Valores muito baixos de colesterol total (<160 mg/dL) associam-se a maior risco de depressão, declínio cognitivo, câncer, hemorragia cerebral e mortalidade por todas as causas, especialmente em idosos. A privação de colesterol compromete a síntese de hormônios e a integridade das membranas neuronais.

A elevação isolada do colesterol total deve ser investigada contextualmente: padrão alimentar, genética (hipercolesterolemia familiar), função tireoidiana (hipotireoidismo eleva colesterol), síndrome metabólica, inflamação crônica."""

    patient_explanation = """O colesterol total é a soma de todos os tipos de colesterol no sangue. Diferente do que muita gente pensa, o colesterol não é um vilão - ele é essencial para a vida. Seu corpo usa colesterol para fabricar hormônios importantes (como testosterona e estrogênio), vitamina D quando você toma sol, e para construir as paredes de todas as células do seu corpo, incluindo as do cérebro.

Cerca de 80% do seu colesterol é produzido pelo próprio fígado, e apenas 20% vem da alimentação. Quando você come menos colesterol, o fígado produz mais, e vice-versa.

O número do colesterol total sozinho não diz muita coisa. O que realmente importa é: que tipos de colesterol você tem? Eles são de boa qualidade? Há inflamação no seu corpo? Por isso, o médico precisa olhar outros exames junto, como o HDL (colesterol "bom"), LDL (colesterol "ruim"), triglicérides e marcadores de inflamação.

Ter colesterol muito baixo (abaixo de 160) pode ser tão ruim quanto ter muito alto, aumentando risco de depressão, problemas de memória e outras doenças. O ideal é um equilíbrio."""

    conduct = """**Interpretação Funcional:**

- **Muito baixo (<160 mg/dL):** Investigar deficiência nutricional, má absorção, hipertireoidismo, doença hepática, uso excessivo de estatinas. Avaliar sintomas neuropsiquiátricos, fadiga, baixa libido.

- **Ótimo funcional (180-240 mg/dL):** Faixa associada a menor mortalidade em estudos populacionais.

- **Elevado (240-300 mg/dL):** Investigar contexto: se HDL alto, triglicerídeos baixos e sem resistência insulínica, pode ser perfil genético protetor. Solicitar painel avançado.

- **Muito elevado (>300 mg/dL):** Investigar hipercolesterolemia familiar, hipotireoidismo, síndrome nefrótica, colestase.

**Avaliação Complementar Obrigatória:**

NUNCA tomar decisão baseada em colesterol total isolado. Solicitar:

**Painel Lipídico Avançado:**
- HDL, LDL, triglicerídeos, VLDL
- Relação triglicerídeos/HDL (ideal <2)
- Colesterol não-HDL (Total - HDL)
- ApoB e ApoA1 → relação ApoB/ApoA1 (melhor preditor de risco)
- Subfracionamento de LDL (partículas pequenas vs. grandes)
- LDL oxidado (anti-LDL oxidado)

**Marcadores Metabólicos:**
- Glicemia, insulina de jejum, HOMA-IR, HbA1c
- PCR ultrassensível (inflamação)
- Homocisteína
- Lipoproteína(a) - Lp(a)
- Perfil de ácidos graxos ômega-3

**Função Tireoidiana:**
- TSH, T4 livre, T3 livre (hipotireoidismo eleva colesterol)

**Estratificação de Risco Cardiovascular:**
- **Escore de cálcio coronariano (CAC):** ESSENCIAL antes de prescrever estatinas em prevenção primária
  - CAC = 0: risco muito baixo, estatina geralmente não indicada
  - CAC 1-100: risco baixo-moderado
  - CAC 100-400: risco moderado-alto
  - CAC >400: risco alto

**Abordagem Terapêutica:**

**Se Colesterol Total Elevado COM resistência insulínica/inflamação:**

1. **Nutrição:**
   - Reduzir carboidratos refinados e açúcares (principal driver de dislipidemia aterogênica)
   - Dieta low carb ou mediterrânea
   - Aumentar ômega-3 (peixes gordos, suplementação 2-3g EPA+DHA)
   - Fibras solúveis (aveia, psyllium, leguminosas)
   - Evitar gorduras trans e óleos vegetais refinados

2. **Exercício:**
   - Musculação + aeróbico: melhora perfil lipídico e sensibilidade insulínica

3. **Suplementação:**
   - **Bergamota (Vasguard):** 1.000 mg/dia - reduz LDL ~30%, aumenta HDL, melhora subtipos
   - **Ômega-3:** EPA+DHA 2-3g/dia
   - **Policosanol:** 10-20 mg/dia
   - **Fitosteróis:** 1,6-2g/dia
   - **Red Yeast Rice (Redisrise):** 600-1.200 mg/dia (contém monacolina K) + CoQ10 100 mg
   - **Magnésio, vitamina D, CoQ10**

4. **Correção de disbiose:** Prebióticos, probióticos

**Se Colesterol Total Normal-Baixo + Sintomas (fadiga, baixa libido, depressão):**

Investigar deficiência de gorduras boas, má absorção, excesso de estatina. Considerar redução/suspensão de estatina se em uso. Aumentar ingestão de gorduras saudáveis (abacate, azeite, oleaginosas, ovos, peixes).

**Estatinas:**

Uso criterioso, NUNCA baseado apenas em colesterol total. Indicações:

- Prevenção secundária (pós-infarto, AVC)
- CAC >100-400 + LDL elevado + múltiplos fatores de risco
- Hipercolesterolemia familiar

Evitar em: CAC=0, baixo risco, idosos frágeis, colesterol total <200 mg/dL sem outros fatores.

**Monitoramento:**

- Perfil lipídico a cada 12 semanas durante intervenção
- Reavaliar ApoB/ApoA1, PCR-us em 3-6 meses
- CAC não precisa repetir anualmente (progressão lenta)"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_hdl_texts() -> Dict[str, str]:
    """Gera textos para HDL Colesterol"""

    clinical_relevance = """O HDL (lipoproteína de alta densidade) é conhecido como "colesterol bom" por mediar o transporte reverso de colesterol: remove colesterol dos tecidos periféricos e placas ateroscleróticas, levando-o ao fígado para excreção biliar. Porém, na medicina funcional, o conceito vai além da quantidade: a FUNCIONALIDADE e o TAMANHO das partículas de HDL são determinantes.

HDL funcional possui propriedades anti-inflamatórias, antioxidantes, vasodilatadoras e antitrombóticas. No entanto, em estados de inflamação crônica, resistência insulínica ou estresse oxidativo intenso, o HDL pode se tornar DISFUNCIONAL, perdendo suas capacidades protetoras e até adquirindo propriedades pró-inflamatórias.

**Valores e Interpretação Funcional:**

- **Baixo (<40 mg/dL homens, <50 mg/dL mulheres):** Associa-se fortemente a síndrome metabólica, resistência insulínica, risco cardiovascular elevado. Frequentemente acompanha triglicerídeos altos (relação TG/HDL >3).

- **Ideal funcional (50-70 mg/dL homens; 60-90 mg/dL mulheres):** Faixa de melhor proteção cardiovascular.

- **Muito alto (>90-100 mg/dL):** Paradoxalmente, pode indicar HDL disfuncional, especialmente se acompanhado de LDL elevado ou inflamação. Estudos mostram que HDL extremamente alto não confere proteção adicional e pode estar associado a mutações genéticas (deficiência de CETP) com maior risco cardiovascular.

**Subtipos de HDL:**

- **HDL grande (HDL2):** Mais eficiente no transporte reverso, anti-inflamatório, protetor. Ideal >6.500 nmol/L (mulheres), >7.000 nmol/L (homens).

- **HDL pequeno (HDL3):** Menos eficiente, pode estar elevado em estados pró-inflamatórios.

**Marcadores de Funcionalidade:**

- **ApoA1:** Principal apoproteína do HDL. Níveis elevados indicam HDL funcional.
- **Relação ApoB/ApoA1:** Melhor preditor de risco que colesterol isolado. Ideal <0,7 (homens), <0,6 (mulheres).
- **Capacidade de efluxo de colesterol:** Teste funcional que mede a eficiência do HDL em remover colesterol das células.

HDL baixo é frequentemente consequência de: resistência insulínica, dieta rica em carboidratos refinados, sedentarismo, obesidade visceral, tabagismo, inflamação crônica. HDL disfuncional ocorre em diabetes, doenças autoimunes, infecções crônicas."""

    patient_explanation = """O HDL é chamado de "colesterol bom" porque funciona como um caminhão de lixo: ele recolhe o excesso de colesterol das artérias e leva de volta para o fígado, onde é eliminado. Isso ajuda a prevenir entupimento das artérias e infarto.

Níveis ideais são acima de 40 mg/dL para homens e 50 mg/dL para mulheres. Quanto mais HDL você tiver (até certo ponto), melhor. Ter HDL baixo aumenta muito o risco de doenças do coração.

Mas atenção: não é só a quantidade que importa, é a qualidade. Se você tem muito açúcar no sangue, inflamação ou obesidade, mesmo que o número do HDL pareça bom, ele pode não estar funcionando direito - é como ter um caminhão de lixo quebrado.

Para aumentar o HDL de verdade, você precisa: fazer exercícios (especialmente musculação e corrida), comer gorduras boas (azeite, abacate, castanhas, peixes), tomar ômega-3, perder barriga e parar de fumar. Açúcar e comida industrializada abaixam o HDL."""

    conduct = """**Interpretação Funcional:**

- **Muito baixo (<30 mg/dL):** Risco cardiovascular muito alto. Investigar síndrome metabólica, resistência insulínica grave, genética (mutações ABCA1).

- **Baixo (30-40 mg/dL homens; 30-50 mg/dL mulheres):** Sinal de risco cardiovascular. Intervenção intensiva necessária.

- **Ótimo (50-70 mg/dL homens; 60-90 mg/dL mulheres):** Faixa de proteção ideal.

- **Muito alto (>90-100 mg/dL):** Investigar funcionalidade. Se LDL também alto ou presença de inflamação, HDL pode estar disfuncional. Solicitar ApoA1, subfracionamento de HDL, capacidade de efluxo.

**Investigação Complementar:**

**Painel Metabólico:**
- Glicemia, insulina, HOMA-IR, HbA1c
- Triglicerídeos → relação TG/HDL (ideal <2)
- LDL, colesterol não-HDL
- ApoB e ApoA1 → relação ApoB/ApoA1

**Marcadores de Inflamação:**
- PCR ultrassensível
- Homocisteína

**Avaliação de Funcionalidade (quando disponível):**
- Subfracionamento de HDL (HDL grande vs. pequeno)
- Capacidade de efluxo de colesterol

**Função Hepática e Tireoidiana:**
- TGO, TGP, GGT
- TSH, T4L, T3L

**Estratégias para Aumentar HDL Funcional:**

**1. Exercício Físico (MAIS EFETIVO):**
   - **Musculação:** Aumenta HDL e melhora subtipos
   - **HIIT:** Eleva HDL rapidamente
   - **Aeróbico moderado:** 150 min/semana mínimo
   - Meta: exercício intenso 4-5x/semana

**2. Nutrição:**
   - **Aumentar gorduras monoinsaturadas:** Azeite extra-virgem, abacate, oleaginosas (castanhas, amêndoas, nozes)
   - **Ômega-3:** Peixes gordos (salmão, sardinha, cavala) 3-4x/semana
   - **Reduzir carboidratos refinados e açúcares:** Principal inimigo do HDL
   - **Fibras solúveis:** Aveia, linhaça, psyllium
   - **Consumo moderado de álcool:** 1 taça de vinho tinto pode elevar HDL (mas excessos prejudicam)

**3. Perda de Peso (especialmente gordura visceral):**
   - Cada 1 kg de gordura perdida aumenta HDL ~0,35 mg/dL

**4. Suplementação:**
   - **Ômega-3 (EPA+DHA):** 2-4g/dia
   - **Niacina (vitamina B3):** 500-2.000 mg/dia (flush comum; usar formulação de liberação prolongada). CUIDADO: monitorar função hepática
   - **Bergamota (Vasguard):** 1.000 mg/dia - aumenta HDL grande
   - **Policosanol:** 10-20 mg/dia - estudos mostram aumento de HDL ~17%
   - **Resveratrol:** 200-500 mg/dia
   - **CoQ10:** 100-200 mg/dia

**5. Estilo de Vida:**
   - **Parar de fumar:** Tabagismo reduz HDL significativamente
   - **Reduzir estresse crônico:** Cortisol elevado prejudica HDL
   - **Sono adequado:** 7-9h/noite

**6. Correção de Resistência Insulínica:**
   - Jejum intermitente 16/8
   - Dieta low carb ou cetogênica
   - Berberina 500 mg 2-3x/dia
   - Magnésio, cromo

**Se HDL >90 mg/dL com LDL alto ou inflamação:**

- Solicitar subfracionamento de HDL e ApoA1
- Avaliar capacidade de efluxo (se disponível)
- Focar em reduzir inflamação (PCR-us), LDL oxidado
- Considerar reduzir consumo excessivo de álcool (se presente)
- Monitorar em 12 semanas

**Monitoramento:**

- Perfil lipídico a cada 8-12 semanas durante intervenção intensiva
- Reavaliar ApoB/ApoA1, TG/HDL em 3 meses
- Exercício + ômega-3 + perda de peso: esperado aumento de 10-20% em 12 semanas"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_ldl_texts() -> Dict[str, str]:
    """Gera textos para LDL Colesterol"""

    clinical_relevance = """O LDL (lipoproteína de baixa densidade) transporta colesterol do fígado aos tecidos periféricos para funções essenciais: síntese de membranas celulares, hormônios esteroides e bainha de mielina. A demonização do LDL como "colesterol ruim" é uma simplificação excessiva que ignora nuances críticas.

**O Paradigma Funcional sobre LDL:**

Na medicina funcional integrativa, o foco não está apenas na QUANTIDADE de LDL, mas em:

1. **QUALIDADE (subtipos):** Partículas pequenas e densas (padrão B) vs. grandes e flutuantes (padrão A)
2. **OXIDAÇÃO:** LDL oxidado, glicado, eletronegativo
3. **CONTEXTO INFLAMATÓRIO:** LDL só penetra no endotélio quando há dano vascular prévio
4. **CARGA ATEROSCLERÓTICA REAL:** Escore de cálcio coronariano (CAC)

**Subtipos de LDL:**

Existem 11 subtipos de LDL, classificados em dois padrões:

- **Padrão A (partículas grandes):** 8 subtipos grandes e flutuantes, pouco aterogênicos. Aumentam com dietas ricas em gorduras saudáveis.

- **Padrão B (partículas pequenas e densas):** 3 subtipos pequenos, densos, penetrantes no endotélio, altamente oxidáveis e aterogênicos. Aumentam com dietas ricas em carboidratos refinados e resistência insulínica.

**Tamanho Pico de LDL:**
- Ideal: ≥223 Å
- Moderado: 218-222 Å
- Ruim: <218 Å (quanto menor, pior)

**LDL "Estragado" (Modificado):**

- **LDL oxidado:** Sofreu peroxidação lipídica por radicais livres. É altamente inflamatório e aterogênico. Marcador: anti-LDL oxidado (ideal <25 U/L).
- **LDL glicado:** Exposto cronicamente a hiperglicemia, forma AGEs (advanced glycation end products).
- **LDL eletronegativo:** Forma altamente modificada, pró-inflamatória.

**A Aterosclerose como Doença da "Porta Aberta":**

O colesterol LDL só se deposita no endotélio quando este já está danificado por:
- Hipertensão (estresse mecânico)
- Hiperglicemia/glicação
- Inflamação crônica (PCR elevada, citocinas)
- Estresse oxidativo (tabagismo, poluição)

Portanto, o LDL é um "oportunista", não um agressor primário. Tratar apenas o LDL sem corrigir o ambiente inflamatório é insuficiente.

**Escore de Cálcio Coronariano (CAC):**

Ferramenta CRUCIAL para decisão terapêutica. Estudo "Power of Zero" demonstrou que pacientes com CAC=0 têm risco cardiovascular muito baixo mesmo com LDL>190 mg/dL. Prescrever estatina nesses casos pode ser desnecessário e até prejudicial."""

    patient_explanation = """O LDL é chamado de "colesterol ruim", mas isso não é totalmente correto. O LDL é um transporte que leva colesterol do fígado para todo o corpo construir células, hormônios e outras coisas importantes.

O problema não é o LDL em si, mas quando ele fica "estragado" (oxidado pelo açúcar, inflamação e radicais livres) e quando a parede das artérias está machucada. Imagine que o LDL é um caminhão de entrega: se a estrada está boa, ele passa tranquilo. Mas se a estrada está cheia de buracos (inflamação), o caminhão derrama a carga e entope tudo.

Além disso, existem dois tipos principais de LDL:
- **LDL grande:** Partículas grandes, como bolas de basquete. São benéficas e não causam problema.
- **LDL pequeno:** Partículas pequenas e densas, como bolinhas de gude. Essas sim penetram fácil nas artérias e causam entupimento.

O que faz você ter mais LDL pequeno e perigoso? Comer muito açúcar, pão, massa e comida industrializada. O que faz você ter mais LDL grande e seguro? Comer gorduras boas (azeite, abacate, castanhas, ovos) e reduzir carboidratos.

Então, o número do LDL sozinho não conta a história toda. É mais importante saber: que TIPO de LDL você tem? Tem inflamação no corpo? Tem açúcar alto? Tem placas nas artérias?"""

    conduct = """**Interpretação Funcional:**

**Valores de LDL (mg/dL):**
- **Ótimo funcional:** 80-120 mg/dL (em contexto metabólico saudável)
- **Limítrofe:** 120-160 mg/dL (depende do contexto: partículas, inflamação, CAC)
- **Alto:** 160-190 mg/dL (investigar qualidade, contexto)
- **Muito alto:** >190 mg/dL (investigar hipercolesterolemia familiar)

**FUNDAMENTAL: NUNCA avaliar LDL isoladamente!**

**Investigação Complementar OBRIGATÓRIA:**

**Painel Lipídico Avançado:**
- Subfracionamento de LDL (tamanho pico, LDL pequeno e denso)
- LDL oxidado (anti-LDL oxidado): ideal <25 U/L
- ApoB (número de partículas aterogênicas): ideal <90 mg/dL
- Relação ApoB/ApoA1: ideal <0,7 (homens), <0,6 (mulheres)
- Triglicerídeos, HDL → relação TG/HDL: ideal <2
- Colesterol não-HDL: (Total - HDL), ideal <130 mg/dL
- Lipoproteína(a) - Lp(a): genética, altamente aterogênica

**Painel Metabólico e Inflamatório:**
- Glicemia, insulina, HOMA-IR, HbA1c
- PCR ultrassensível: ideal <1 mg/L
- Homocisteína: ideal <8 µmol/L
- Fibrinogênio

**Estratificação de Risco Cardiovascular:**
- **Escore de Cálcio Coronariano (CAC):** ESSENCIAL em prevenção primária antes de decidir estatina
  - CAC = 0: Risco muito baixo. Estatina raramente indicada mesmo com LDL alto.
  - CAC 1-100: Risco baixo-moderado
  - CAC 100-400: Risco moderado-alto
  - CAC >400: Risco alto. Estatina + intervenção agressiva.

**Abordagem Terapêutica:**

**Cenário 1: LDL elevado + CAC=0 + Sem resistência insulínica/inflamação:**

- **Conduta:** Observação + intervenção no estilo de vida
- Focar em melhorar QUALIDADE do LDL (aumentar partículas grandes)
- Reduzir inflamação e oxidação
- Reavaliar em 6-12 meses
- Estatina NÃO indicada

**Cenário 2: LDL elevado + Resistência insulínica/inflamação:**

**1. Nutrição (PILAR CENTRAL):**
   - **Dieta Low Carb ou Mediterrânea:** Reduz partículas pequenas de LDL, aumenta partículas grandes
   - Eliminar carboidratos refinados, açúcares, ultraprocessados
   - Aumentar gorduras saudáveis: azeite extra-virgem, abacate, oleaginosas, peixes gordos
   - Fibras solúveis: aveia, psyllium, leguminosas (reduzem LDL ~10%)
   - Evitar óleos vegetais refinados, gorduras trans

**2. Exercício:**
   - Musculação + aeróbico: melhora partículas de LDL, reduz inflamação

**3. Suplementação:**
   - **Bergamota (Vasguard):** 1.000 mg/dia - reduz LDL ~30%, melhora subtipos (equivalente a rosuvastatina 10 mg)
   - **Red Yeast Rice (Redisrise):** 600-1.200 mg/dia (contém monacolina K precursora de estatina natural) + CoQ10 100 mg
   - **Policosanol:** 10-20 mg/dia - reduz LDL ~28%
   - **Fitosteróis:** 1,6-2g/dia - reduz LDL ~10% (bloqueia absorção)
   - **Ômega-3 (EPA+DHA):** 2-4g/dia - reduz triglicerídeos, melhora partículas
   - **Ácido Alfa-Lipoico:** 600 mg/dia - reduz LDL oxidado
   - **NAC (N-acetilcisteína):** 600-1.200 mg/dia - antioxidante
   - **Vitamina E, C, Selênio:** Reduzem oxidação de LDL

**4. Correção de Resistência Insulínica:**
   - Berberina 500 mg 2-3x/dia
   - Jejum intermitente
   - Magnésio, cromo

**Cenário 3: LDL elevado + CAC >100 ou Prevenção Secundária:**

- **Estatina indicada** (efeitos pleotrópicos: anti-inflamatório, estabiliza placa)
- Meta LDL: 80-100 mg/dL (não obrigatoriamente <70 mg/dL)
- Combinar estatina dose moderada + suplementação (bergamota, Red Yeast Rice) para reduzir dose de estatina
- SEMPRE suplementar CoQ10 100-200 mg/dia (estatina depleta CoQ10)
- Monitorar enzimas hepáticas, CK (risco de miopatia)

**Cenário 4: LDL muito alto (>190 mg/dL):**

- Investigar hipercolesterolemia familiar (HF):
  - História familiar de infarto precoce (<55 anos homens, <65 anos mulheres)
  - Xantomas (depósitos de colesterol na pele/tendões)
  - Arco corneano precoce
  - Teste genético se disponível
- HF confirmada: estatina de alta potência + ezetimiba + intervenção agressiva no estilo de vida

**Foco em Reduzir LDL Oxidado:**

- Reduzir estresse oxidativo: antioxidantes, ômega-3, vitaminas
- Controlar glicemia (glicação)
- Reduzir inflamação (PCR-us)
- Parar de fumar (maior oxidante)

**Monitoramento:**

- Perfil lipídico + ApoB + LDL oxidado a cada 8-12 semanas durante intervenção
- Reavaliar CAC apenas se >5 anos (progressão lenta)
- Se estatina: monitorar TGO, TGP, CK a cada 3 meses no primeiro ano"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }
