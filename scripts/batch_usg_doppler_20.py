#!/usr/bin/env python3
"""
Batch de 20 exames USG e Doppler - Medicina Funcional Integrativa
IDs fornecidos pelo usuário
"""
import requests
import json
import time

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

def login():
    response = requests.post(f"{API_URL}/auth/login", json={"email": EMAIL, "password": PASSWORD})
    response.raise_for_status()
    return response.json()["accessToken"]

def update_item(token, item_id, clinical, patient, conduct):
    headers = {"Authorization": f"Bearer {token}"}
    data = {
        "clinicalRelevance": clinical,
        "patientExplanation": patient,
        "conduct": conduct
    }
    response = requests.put(f"{API_URL}/score-items/{item_id}", json=data, headers=headers)
    return response.status_code == 200

# Lista dos 20 items fornecidos
ITEMS = [
    ("92891a23-02aa-4d4a-a3c2-f770b6505759", "40901114 - USG mamas"),
    ("ca3db864-3407-41a2-b317-07fb2e53253e", "40901114 - USG mamas"),
    ("9f306531-7f98-4b87-80fc-8fbdc46d952d", "40901114 - USG mamas"),
    ("e44680b0-138a-4a8f-812b-eedca2838c94", "40901122 - USG abdome total"),
    ("a87ee825-298c-49eb-aea2-183981f49a89", "40901122 - USG abdome total"),
    ("1b1e689d-14af-4b66-a269-be27ddcce26d", "40901122 - USG abdome total"),
    ("b29bbeb0-b0bd-4f71-acaa-fab10b2ac45d", "40901300 - USG transvaginal"),
    ("5a315bc4-8f86-412f-be04-315636a9e3ec", "40901300 - USG transvaginal"),
    ("1178ea33-d57f-419d-a9b2-a90474bb30d1", "40901300 - USG transvaginal"),
    ("f22b2474-b22e-4fc9-bf44-4706378bf88a", "40901360 - Doppler carótidas"),
    ("e7b2ab85-9ac7-4edf-b888-cc373bcf7b4e", "40901360 - Doppler carótidas"),
    ("bf1e1008-8a30-41f8-b39a-067ce5ef782a", "40901360 - Doppler carótidas"),
    ("01f1d5e4-4b02-45e5-8a07-3830a94474ce", "40901394 - Doppler aorta/renais"),
    ("7fe1681d-b0e7-485f-ae06-7391d98c6df6", "40901394 - Doppler aorta/renais"),
    ("ee825218-ca7a-47c6-b0d9-78e086bbbf08", "40901394 - Doppler aorta/renais"),
    ("8815b8b2-8a3f-4d04-bdb6-4b9f2b8a488a", "40901750 - USG próstata"),
    ("d149c9ab-f69b-4ca2-afa1-62d8b096c055", "40901750 - USG próstata"),
    ("4bbaf683-98af-4947-a99e-5a99e96ced70", "40901750 - USG próstata"),
    ("a7368290-1c34-4376-b7ea-b49351283189", "40901114 - USG mamas"),
    ("c0ee3c8a-1a83-4853-867a-e6f66893af90", "40901122 - USG abdome total"),
]

def get_content_usg_mamas():
    return {
        "clinical": """A ultrassonografia mamária é exame complementar fundamental na propedêutica das mamas, especialmente em mulheres <40 anos (mama densa), como adjunto à mamografia em rastreamento combinado, e para caracterização de nódulos detectados em mamografia ou exame físico. Utiliza ondas sonoras de alta frequência, sem radiação ionizante, permitindo diferenciação entre lesões sólidas e císticas com alta acurácia.

Na medicina funcional integrativa, o USG mamas é componente de estratégia preventiva multimodal contra câncer de mama (segunda causa de morte por câncer em mulheres no Brasil). Avaliamos não apenas achados ultrassonográficos, mas o contexto hormonal-metabólico completo: metabolismo estrogênico (ratio 2-OH-estrone/16-OH-estrone idealmente >2.0, refletindo metabolização favorável vs carcinogênica), equilíbrio progesterona/estrogênio, resistência insulínica (hiperinsulinemia crônica é fator de risco independente para Ca mama via ativação IGF-1), inflamação sistêmica (PCR-us, interleucinas), status de vitamina D (receptores VDR no tecido mamário - níveis >50 ng/mL associados a 50% menor risco de Ca mama).

Achados ultrassonográficos classificados pelo BI-RADS US: 1 (normal), 2 (benigno definitivo - cistos simples, linfonodos intramamários, lipomas), 3 (provavelmente benigno - <2% malignidade, seguimento em 6 meses), 4 (suspeito - 2-95% malignidade, requer biópsia core), 5 (altamente sugestivo de malignidade >95%, biópsia urgente), 6 (malignidade provada por biópsia). Nódulos sólidos hipoecóicos com margens irregulares, orientação não-paralela, sombra acústica posterior, aumento da vascularização ao Doppler colorido são sinais de alerta.

Cistos mamários simples são achados benignos extremamente comuns (50% das mulheres), relacionados a flutuações hormonais (excesso estrogênico relativo, deficiência de progesterona). Fibroadenomas são tumores benignos sólidos comuns em mulheres jovens. Alterações fibrocísticas difusas refletem sensibilidade a estrogênio e deficiências nutricionais (iodo, vitamina E, ácidos graxos essenciais).

A ultrassonografia permite também avaliação axilar (linfonodos - arquitetura preservada vs cortical espessada suspeita) e guia procedimentos intervencionistas (biópsia core guiada por USG tem taxa de sucesso >95% com baixa complicação).""",

        "patient": """A ultrassonografia (USG) de mamas é um exame de imagem que usa ondas sonoras para ver o interior das mamas. É parecido com o ultrassom feito em grávidas, totalmente indolor e sem radiação.

O USG das mamas é muito útil porque consegue diferenciar um caroço sólido (que pode ser benigno ou maligno) de um cisto simples cheio de líquido (sempre benigno). Muitas vezes, quando a mamografia detecta uma alteração, o ultrassom é solicitado para esclarecer melhor do que se trata.

Em mulheres jovens (abaixo de 40 anos), a mama costuma ser muito densa, o que dificulta a mamografia. Nesses casos, o ultrassom é o exame preferido. Também é excelente para avaliar nódulos palpáveis (caroços que você ou seu médico sentiram no autoexame ou exame clínico).

Cistos mamários são extremamente comuns e benignos - são "bolsinhas" de líquido que aparecem devido a alterações hormonais. Não viram câncer e muitas vezes não precisam de tratamento. Nódulos sólidos precisam de avaliação mais cuidadosa: alguns são benignos (fibroadenomas), outros requerem biópsia.

O exame não previne câncer, mas detecta alterações precocemente, quando o tratamento é mais eficaz. É sempre recomendável combinar com mamografia após os 40 anos, autoexame mensal e exame clínico anual com seu médico.""",

        "conduct": """**Indicações USG Mamas:**
Rastreamento em mulheres <40 anos (mama densa), complemento à mamografia em rastreamento combinado (aumenta detecção em 2-4 casos/1.000 mulheres), avaliação de nódulo palpável, caracterização de achados mamográficos (BI-RADS 0 ou 3), mastite/abscesso, orientação de biópsia, seguimento de lesões BI-RADS 3, avaliação axilar.

**Interpretação BI-RADS US:**
- **BI-RADS 1-2:** Tranquilizar, manter rastreamento habitual
- **BI-RADS 3:** Seguimento em 6, 12 e 24 meses (alternativa à biópsia imediata em lesões provavelmente benignas)
- **BI-RADS 4 ou 5:** Biópsia core guiada por USG (14-16G, mínimo 5 fragmentos) é obrigatória
- **Cistos simples:** Benignos, não requerem seguimento; se dolorosos/grandes, considerar aspiração

**Abordagem Funcional Integrativa Preventiva:**

**1. Otimização Hormonal Estrogênica:**
Testar metabolitos urinários de estrogênio (2-OH-E1, 4-OH-E1, 16-OH-E1). Ratio 2-OH/16-OH ideal >2.0 (metabolização protetora). Se desfavorável:
- DIM (3,3'-diindolylmethane) 200-400 mg/dia ou I3C (indole-3-carbinol) 400-800 mg
- Vegetais crucíferos diários (brócolis, couve-flor, couve, repolho - fonte natural de I3C)
- Sulforafano (brócolis germinado) 30-60 mg
- Cálcio-D-glucarato 500-1.000 mg (inibe beta-glucuronidase, reduz recirculação entero-hepática de estrogênio)

**2. Equilíbrio Progesterona/Estrogênio:**
Testar progesterona luteal (dia 19-21 ciclo). Se deficiência lútea: Vitex agnus-castus 400 mg/dia (aumenta progesterona endógena). Em sintomas graves ou pós-menopausa, considerar progesterona bioidêntica micronizada 100-200 mg transdérmica ou oral (noturno).

**3. Iodo (crítico para saúde mamária):**
Tecido mamário concentra iodo (receptores NIS). Deficiência associada a mastalgia, alterações fibrocísticas e risco aumentado de Ca. Testar iodo urinário 24h (meta >150 mcg/L). Suplementar iodeto de potássio 3-6 mg/dia (ou Lugol 2% - 2 gotas = 5 mg) + selênio 200 mcg (cofator). Monitorar TSH (auto-imunidade tireoidiana).

**4. Modulação Inflamatória:**
- Ômega-3 EPA+DHA 2-3 g/dia (estudos mostram redução de 32% risco Ca mama)
- Curcumina 1.000 mg/dia (inibe NF-kB, COX-2)
- Vitamina D3: meta >50 ng/mL (suplementar 5.000-10.000 UI/dia conforme nível basal). Receptores VDR no tecido mamário - ação antiproliferativa

**5. Resistência Insulínica (fator de risco subestimado):**
Hiperinsulinemia crônica estimula proliferação celular via IGF-1. Testar insulina jejum (ideal <5 µU/mL), HOMA-IR (<1.0), HbA1c (<5.3%). Se alterado: dieta low-carb, jejum intermitente 16/8, berberina 500 mg 3x/dia, metformina 500-1.000 mg 2x/dia, exercício 150 min/sem.

**6. Detoxificação Hepática (fase II):**
Suportar conjugação e excreção de estrogênio:
- NAC (N-acetilcisteína) 1.200 mg/dia
- Glicina 3-5 g/dia
- Glutationa lipossomal 500 mg
- Evitar xenoestrogênios: plásticos (BPA, ftalatos), parabenos (cosméticos), pesticidas organofosforados

**7. Microbioma Intestinal (Estroboloma):**
Bactérias intestinais metabolizam estrogênio conjugado via beta-glucuronidase. Disbiose aumenta recirculação de estrogênio ativo. Probióticos específicos: Lactobacillus rhamnosus GG, Lactobacillus acidophilus, Bifidobacterium longum. Prebióticos (FOS, inulina).

**8. Composição Corporal:**
Tecido adiposo (especialmente visceral) expressa aromatase, convertendo andrógenos em estrogênio. Perda de excesso de gordura reduz carga estrogênica. Meta IMC 18,5-24,9, gordura corporal <30%.

**9. Estilo de Vida:**
- Exercício aeróbico + resistido 150 min/sem (reduz risco 25%)
- Evitar álcool (aumenta estrogênio, reduz metilação) ou limitar a <3 doses/sem
- Cessar tabagismo
- Sono 7-9h/noite (melatonina tem ação antiestrogênica e antioxidante)

**10. Rastreamento:**
Mamografia anual 40-74 anos (iniciar aos 35 se histórico familiar forte). USG mamas como complemento se mama densa (ACR tipo C ou D). Autoexame mensal, exame clínico anual.

**Se BI-RADS 3:** Além de seguimento ultrassonográfico em 6 meses, implementar protocolo funcional completo acima - objetivo é não apenas monitorar, mas reduzir risco e reverter fatores promotores.

**Se BI-RADS 4-5:** Biópsia core urgente (não retardar). Se confirmado Ca: tratamento oncológico convencional (cirurgia ± quimio/rádio/hormônio) + suporte funcional integrativo adjuvante."""
    }

def get_content_usg_abdome():
    return {
        "clinical": """A ultrassonografia de abdome total é método de imagem não-invasivo, sem radiação ionizante, que avalia morfologia, dimensões e ecotextura de órgãos abdominais sólidos (fígado, vesícula biliar, pâncreas, baço, rins, retroperitônio) e estruturas vasculares (aorta, veia cava, vasos portais). É exame de primeira linha na investigação de dor abdominal, alterações laboratoriais hepatobiliares, massas abdominais e rastreamento de litíase biliar/renal.

Na medicina funcional integrativa, o USG abdominal revela achados comuns mas altamente relevantes metabolicamente, muitas vezes subvalorizados: **esteatose hepática não-alcoólica (DHGNA/NAFLD)** - presente em 25-35% da população geral e >70% dos obesos/diabéticos - é manifestação hepática da síndrome metabólica e resistência insulínica sistêmica. DHGNA não é condição benigna: 20-30% progridem para esteato-hepatite (NASH) com inflamação e fibrose, podendo evoluir para cirrose e carcinoma hepatocelular. USG detecta esteatose moderada-grave (>30% hepatócitos com gordura), porém subestima casos leves e não diferencia esteatose simples de NASH (requer elastografia ou biópsia).

**Litíase biliar (colelitíase):** cálculos em vesícula detectados em 10-20% adultos, maioria assintomática. Fatores de risco: "4 Fs" - Female, Forty, Fatty, Fertile + dieta rica em carboidratos refinados, jejum prolongado, perda de peso rápida, resistência insulínica, hipertrigliceridemia, disbiose intestinal. Complicações: colecistite aguda, coledocolitíase, pancreatite aguda, colangite.

**Litíase renal (nefrolitíase):** prevalência 10% população, recorrência 50% em 5 anos. Composição: 80% cálcio (oxalato ou fosfato), 10% ácido úrico, 10% estruvita/cistina. Investigação funcional identifica causas metabólicas corrigíveis: hipercalciúria (excesso proteína animal, sódio, deficiência magnésio), hiperoxalúria (dieta rica em oxalato, disbiose intestinal com redução Oxalobacter formigenes, deficiência B6), hiperuricosúria (excesso purinas, resistência insulínica), hipocitratúria (acidose metabólica crônica, deficiência potássio), cistinúria (genética).

**Cistos renais simples:** extremamente comuns (>50% >50 anos), benignos, não requerem seguimento. Cistos complexos classificados por Bosniak (I-IV) - categorias III-IV requerem investigação adicional (TC/RM) por risco de malignidade.

**Esplenomegalia:** baço >12 cm sugere hipertensão portal (cirrose, trombose veia porta), doenças hematológicas (leucemias, linfomas, anemias hemolíticas), infecções crônicas (mononucleose, malária), doenças de depósito.

O USG abdominal também avalia aorta abdominal (rastreamento de aneurisma em homens >65 anos, especialmente tabagistas - diâmetro >3 cm define aneurisma, >5,5 cm indica alto risco de ruptura).""",

        "patient": """A ultrassonografia (USG) de abdome total é um exame de imagem que usa ondas sonoras para visualizar órgãos dentro da barriga: fígado, vesícula, pâncreas, baço, rins e grandes vasos sanguíneos. É o mesmo tipo de exame feito em grávidas - totalmente indolor, sem radiação.

O exame detecta uma série de condições: **fígado gorduroso** (acúmulo de gordura no fígado, muito comum em quem tem sobrepeso, diabetes ou colesterol alterado), **pedras na vesícula** (pequenas "pedrinhas" que podem causar dor intensa do lado direito após refeições gordurosas), **pedras nos rins** (causam dor lombar intensa quando se movem), **cistos** nos rins ou fígado (geralmente benignos), **aumento do baço**, e até **dilatação da aorta** (artéria principal do abdome).

Fígado gorduroso é um achado muito comum e importante: significa que seu corpo está com dificuldade em processar açúcares e gorduras. Se não tratado, pode evoluir para inflamação do fígado, fibrose e até cirrose. A boa notícia é que é totalmente reversível com mudanças na alimentação e perda de peso.

Pedras na vesícula são muito comuns (1 em cada 5 adultos tem) e a maioria nunca causa sintomas. Mas se você tiver dor forte do lado direito após comer, pode ser necessário remover a vesícula.

O exame requer jejum de 6-8 horas para que a vesícula fique cheia de bile e seja bem visualizada.""",

        "conduct": """**Indicações USG Abdome Total:**
Dor abdominal (especialmente hipocôndrio direito - suspeita de colecistite/colelitíase), alterações de enzimas hepáticas (TGO, TGP, GGT, FA, bilirrubinas), icterícia, hepatomegalia/esplenomegalia palpável, massa abdominal, suspeita de litíase renal (dor lombar, hematúria), rastreamento de aneurisma de aorta (homens >65 anos, tabagistas), seguimento de achados prévios.

**Interpretação e Conduta Funcional por Achado:**

**1. ESTEATOSE HEPÁTICA (Fígado Gorduroso):**

Classificação ultrassonográfica: leve, moderada ou acentuada (baseado em aumento de ecogenicidade, atenuação posterior, borramento vascular).

Investigação complementar obrigatória:
- Painel metabólico: glicemia jejum, HbA1c, insulina jejum, HOMA-IR, peptídeo C
- Perfil lipídico avançado: TG, HDL, LDL, VLDL, ApoB, partículas LDL, ratio TG/HDL (<2.0 ideal)
- Função hepática: TGO, TGP, GGT, FA, bilirrubinas, albumina, TAP/INR
- Marcadores de fibrose não-invasivos: FIB-4, APRI score, elastografia hepática transitória (Fibroscan - quantifica fibrose em kPa: <7 kPa normal, 7-9,5 kPa F2, 9,5-12,5 kPa F3, >12,5 kPa F4/cirrose)
- Descartar causas secundárias: sorologias hepatite B/C, ferritina/saturação transferrina (hemocromatose), anti-músculo liso/FAN/IgG (hepatite autoimune), ceruloplasmina (doença de Wilson <40 anos)

**Abordagem Funcional Integrativa para DHGNA/NASH:**

**Pilar 1 - Perda de Peso (mais eficaz):**
Meta: 7-10% peso corporal (normaliza DHGNA em 90%, reverte NASH em 45%). Estratégias:
- Dieta low-carb (<100 g/dia) ou cetogênica (<50 g/dia): reduz lipogênese de novo hepática
- Jejum intermitente 16/8 (janela alimentar 8h): ativa autofagia, mitofagia, reduz gordura visceral
- Evitar **frutose** (xarope de milho, refrigerantes, sucos industrializados): frutose metabolizada exclusivamente no fígado → DNL (de novo lipogenesis) → TG hepáticos
- Proteína 1,2-1,6 g/kg/dia (preserva massa magra durante perda de peso)
- Eliminar ultraprocessados, óleos vegetais refinados (soja, milho, canola - pró-inflamatórios)

**Pilar 2 - Exercício Físico:**
Musculação 3-4x/sem (aumenta captação de glicose/ácidos graxos pelo músculo, reduz aporte hepático) + aeróbico 150 min/sem. Exercício sozinho (sem perda peso) melhora esteatose em 20-30%.

**Pilar 3 - Suplementação Hepatoprotetora:**
- **Silimarina** (milk thistle) 420-600 mg/dia: regeneração hepatocelular, antioxidante, anti-fibrótico
- **NAC** (N-acetilcisteína) 1.200-1.800 mg/dia: precursor glutationa, detoxificação fase II
- **Ácido alfa-lipoico** 600-1.200 mg/dia: antioxidante mitocondrial, melhora sensibilidade insulínica
- **Vitamina E** 800 UI/dia: apenas se NASH confirmado por biópsia (estudos mostram melhora histológica, mas uso prolongado controverso)
- **Colina** 500-1.000 mg/dia + **inositol** 2-4 g/dia: exportação de TG hepáticos via VLDL
- **Berberina** 500 mg 3x/dia: melhora resistência insulínica, reduz DNL hepática (eficácia similar à metformina)
- **Ômega-3** EPA+DHA 2-4 g/dia: anti-inflamatório, reduz TG hepáticos
- **Curcumina** 1.000 mg/dia: anti-inflamatório potente (inibe NF-kB)
- **Probióticos** multi-cepas: corrige disbiose, reduz endotoxemia metabólica (LPS bacteriano via portal perpetua inflamação hepática)

**Pilar 4 - Correção de Resistência Insulínica:**
Metformina 500-1.000 mg 2x/dia (off-label, melhora DHGNA via AMPK). Pioglitazona 15-30 mg (se diabetes/pré-diabetes + NASH - evidência histológica de melhora, mas ganho peso).

**Pilar 5 - Evitar Hepatotoxinas:**
Álcool (zero ou <1 dose/dia mulheres, <2 homens), paracetamol em excesso, AINEs crônicos.

**Seguimento:**
- Exames lab 8-12 semanas
- Repetir USG abdome em 6-12 meses (esperar melhora ecogenicidade)
- Se TGP >2x normal persistente: considerar biópsia hepática (padrão-ouro para estadiar NASH/fibrose)

---

**2. LITÍASE BILIAR (Cálculos na Vesícula):**

**Assintomáticos (maioria):** Conduta expectante. Prevenir crescimento/sintomas:
- Evitar jejum prolongado (vesícula estagnada favorece precipitação)
- Perder peso gradualmente (1-2 kg/mês - perda rápida precipita cálculos)
- Fibras solúveis (reduz colesterol biliar): psyllium, aveia, linhaça
- Lecitina de soja 1.200 mg 3x/dia (emulsifica colesterol)
- Vitamina C 500-1.000 mg/dia (converte colesterol em ácidos biliares)
- Curcumina 500 mg/dia (estimula contração vesicular)

**Sintomáticos (cólica biliar, colecistite aguda):**
Colecistectomia videolaparoscópica é tratamento definitivo. Pós-operatório: suplementar sais biliares (ox bile 500 mg) + enzimas digestivas se dispepsia, garantir digestão de gorduras.

---

**3. LITÍASE RENAL (Cálculos Renais):**

Investigação metabólica: urina 24h (cálcio, oxalato, ácido úrico, citrato, sódio, magnésio, creatinina, volume), pH urinário, exame de cálculo (análise composição).

**Conduta Funcional por Tipo:**

**Cálculos de Oxalato de Cálcio (80%):**
- **Hidratação:** 2,5-3L água/dia (meta volume urinário >2L/dia - dilui urina)
- **Citrato de potássio:** 30-60 mEq/dia (citrato inibe cristalização, alcaliniza urina)
- **Magnésio:** 400-600 mg/dia (liga oxalato intestinal, inibe cristalização)
- **Vitamina B6:** 50-100 mg/dia (reduz produção endógena de oxalato)
- Reduzir oxalato dietético: evitar espinafre, acelga, beterraba, chocolate, nozes, chá preto
- Reduzir sódio (<2.300 mg/dia - sódio aumenta calciúria)
- Moderar proteína animal (1,0-1,2 g/kg/dia - excesso acidifica urina, aumenta calciúria)
- Cálcio alimentar adequado 1.000-1.200 mg/dia (liga oxalato intestinal, **não** restringir!)
- Probióticos com **Lactobacillus** e **Oxalobacter formigenes** (degradam oxalato intestinal)

**Cálculos de Ácido Úrico (10%):**
- Alcalinizar urina: citrato de potássio (meta pH 6,5-7,0)
- Reduzir purinas: carnes vermelhas, vísceras, frutos do mar, álcool
- Alopurinol 100-300 mg se hiperuricemia persistente
- Hidratação 2,5-3L/dia
- Reverter resistência insulínica (hiperinsulinemia reduz excreção renal de ácido úrico)

**Seguimento:**
USG ou RX abdome anual. Meta: prevenir novos cálculos (recorrência 50% em 5 anos sem profilaxia).

---

**4. CISTOS RENAIS SIMPLES:**
Bosniak I-II: benignos, sem seguimento. Bosniak III-IV: TC/RM contrastada para descartar neoplasia.

**5. ANEURISMA DE AORTA ABDOMINAL:**
Diâmetro >3 cm: seguimento anual. >5,5 cm ou expansão >0,5 cm/ano: correção cirúrgica/endovascular. Controle rigoroso HAS, cessar tabagismo, estatina."""
    }

def get_content_usg_transvaginal():
    return {
        "clinical": """A ultrassonografia transvaginal (USTV) é método de imagem de alta resolução para avaliação detalhada de órgãos pélvicos femininos: útero (miométrio, endométrio, colo), ovários, trompas, fundo de saco de Douglas, região anexial. Sonda intracavitária de alta frequência (5-9 MHz) fornece imagens superiores ao USG abdominal, especialmente em mulheres obesas ou com bexiga vazia.

Na medicina funcional integrativa, a USTV detecta achados que refletem desequilíbrios hormonais-metabólicos subjacentes: **síndrome dos ovários policísticos (SOP)**, **miomas uterinos**, **adenomiose**, **endometriose**, **pólipos endometriais**, **hiperplasia endometrial**, **cistos ovarianos funcionais/patológicos**, **câncer de endométrio/ovário**.

**SÍNDROME DOS OVÁRIOS POLICÍSTICOS (SOP):** prevalência 10-15% mulheres em idade reprodutiva, principal causa de infertilidade anovulatória. Critérios diagnósticos Rotterdam (2/3): 1) oligo/anovulação, 2) hiperandrogenismo clínico/laboratorial, 3) morfologia policística ovariana ao USG (≥12 folículos 2-9 mm e/ou volume ovariano ≥10 mL). Na perspectiva funcional, SOP é manifestação ovariana de **resistência insulínica sistêmica** em 70-80% casos (incluindo mulheres magras). Hiperinsulinemia estimula células tecais ovarianas a produzir andrógenos (testosterona, androstenediona) + inibe SHBG hepática (aumenta testosterona livre) + bloqueia enzima aromatase (reduz conversão andrógenos→estrogênios). Resulta em: irregularidade menstrual, acne, hirsutismo, alopecia androgenética, infertilidade, risco aumentado de diabetes tipo 2 (35-50% em 10 anos), doença cardiovascular, câncer de endométrio (exposição estrogênica sem oposição de progesterona por anovulação crônica).

**MIOMAS UTERINOS (leiomiomas):** tumores benignos do músculo liso uterino, presentes em 70-80% mulheres até 50 anos (maioria assintomática). Fatores de risco: nuliparidade, menarca precoce, obesidade (aromatização periférica de andrógenos em estrogênio), raça negra, histórico familiar. Sintomas dependem de localização/tamanho: menorragia (sangramento aumentado - anemia ferropriva), dismenorreia, dor pélvica, compressão vesical/retal, infertilidade. USTV classifica miomas: **subserosos** (crescem para fora do útero), **intramurais** (dentro da parede), **submucosos** (projetam-se para cavidade - maior impacto em sangramento/fertilidade). Patogênese: receptores estrogênio/progesterona em miomas (crescem sob influência hormonal, regridem na menopausa), fatores de crescimento (IGF-1, EGF), desregulação apoptose.

**ADENOMIOSE:** presença de glândulas endometriais e estroma dentro do miométrio (músculo uterino). Afeta 20-35% mulheres, mais comum >40 anos, multíparas. Sintomas: dismenorreia severa progressiva, menorragia, dispareunia profunda, dor pélvica crônica. USTV: útero globoso aumentado, estrias lineares/cistos miometriais, assimetria parede anterior/posterior, zonas juncionais mal definidas. Diferenciação de miomas: adenomiose = bordas mal definidas, difusa; miomas = lesões circunscritas.

**PÓLIPOS ENDOMETRIAIS:** crescimentos focais do endométrio, prevalência 10-30% (aumenta com idade). Causas: excesso estrogênico (obesidade, TRH, tamoxifeno), hipertensão, diabetes. Sintomas: sangramento uterino anormal (intermenstrual, pós-menopausa). USTV: lesão hiperecóica bem definida projetando-se para cavidade, vascularização central ao Doppler (vaso pedicular único - sinal diferencial vs coágulo/pólipo fibrinoso). Risco malignidade: 0,5-5% (maior em pós-menopausa, sangramento, >1,5 cm). Conduta: polipectomia histeroscópica com biópsia.

**ESPESSAMENTO ENDOMETRIAL:** medido em corte longitudinal (dupla camada). Valores: **menacme** - proliferativa 5-7 mm, secretora 7-16 mm; **pós-menopausa** - idealmente <4-5 mm. Espessamento >5 mm pós-menopausa com sangramento OU >11 mm sem sangramento = investigação obrigatória (histeroscopia + biópsia) para descartar hiperplasia/câncer endometrial. Causas: excesso estrogênico (obesidade, anovulação crônica/SOP, TRH sem progesterona, tumor secretor estrogênio), tamoxifeno, hiperplasia endometrial (simples/complexa, com/sem atipia - atipia = precursor câncer), câncer de endométrio.

**CISTOS OVARIANOS:**
- **Funcionais** (foliculares, corpo lúteo): benignos, relacionados a ciclo menstrual, regridem espontaneamente em 1-3 ciclos. USG: uniloculares, anecóicos, paredes finas, <5 cm, sem vascularização interna
- **Patológicos** (cistoadenomas seroso/mucinoso, teratomas/dermóides, endometriomas): persistentes, podem crescer. Endometriomas ("cistos de chocolate"): conteúdo hemorrágico denso, ecogenicidade em vidro fosco, associados a endometriose

Classificação de risco: escore IOTA (International Ovarian Tumor Analysis) ou O-RADS US identifica cistos com risco de malignidade. Marcadores tumorais (CA-125) complementam. Cistoadenomas mucinosos podem atingir grandes dimensões (>10 cm), risco de ruptura/torção. Teratomas maduros (dermóides) têm conteúdo heterogêneo (cabelo, dente, gordura) - calcificações, ecogenicidade mista.""",

        "patient": """A ultrassonografia transvaginal é um exame de imagem em que uma pequena sonda (coberta com preservativo e gel) é introduzida suavemente na vagina para visualizar útero, ovários e região pélvica de perto. É mais detalhado que o ultrassom abdominal e não requer bexiga cheia. Pode causar leve desconforto, mas não é doloroso.

O exame investiga causas de: sangramento uterino anormal, dor pélvica, irregularidade menstrual, infertilidade, alterações no exame de toque ginecológico. Detecta condições como:

**Ovários policísticos (SOP):** ovários com múltiplos pequenos cistos, associado a resistência à insulina (mesmo em mulheres magras). Causa menstruação irregular, dificuldade para engravidar, acne, aumento de pelos, queda de cabelo. Tratamento foca em reverter resistência insulínica com alimentação baixa em carboidratos, exercícios e suplementos específicos.

**Miomas:** "caroços" benignos no útero, muito comuns (7 em cada 10 mulheres terão miomas ao longo da vida). Maioria não causa sintomas, mas alguns provocam sangramento menstrual intenso (levando à anemia), cólicas fortes e dificuldade para engravidar.

**Adenomiose:** tecido que normalmente reveste o útero cresce dentro da parede muscular, causando cólicas progressivamente piores, sangramento aumentado e dor pélvica crônica.

**Pólipos endometriais:** pequenos crescimentos dentro do útero que podem sangrar entre menstruações ou após menopausa.

**Espessamento do endométrio:** camada interna do útero está mais espessa que o normal, podendo indicar excesso de estrogênio. Na pós-menopausa com sangramento, sempre requer investigação para descartar câncer.

**Cistos ovarianos:** "bolsas" com líquido nos ovários. Cistos pequenos e simples são normais e desaparecem sozinhos. Cistos grandes, complexos ou que persistem precisam de acompanhamento.""",

        "conduct": """**Indicações USG Transvaginal:**
Sangramento uterino anormal (menorragia, metrorragia, sangramento pós-menopausa), dismenorreia, dor pélvica crônica, irregularidade menstrual/amenorreia, infertilidade, massa pélvica palpável, rastreamento em uso de tamoxifeno, seguimento de achados prévios (miomas, cistos).

**Interpretação e Conduta Funcional por Achado:**

**1. SÍNDROME DOS OVÁRIOS POLICÍSTICOS (SOP):**

Diagnóstico Rotterdam: 2 de 3 critérios (oligo/anovulação + hiperandrogenismo + USG policístico).

Investigação laboratorial obrigatória:
- Hormônios: LH, FSH (ratio LH/FSH >2:1 sugere SOP), testosterona total/livre, DHEA-S, androstenediona, SHBG, prolactina, TSH (excluir hipotireoidismo), 17-OH-progesterona (excluir hiperplasia adrenal congênita)
- Metabólico: glicemia jejum, insulina jejum, HOMA-IR, HbA1c, TOTG com curva de insulina (70% SOP têm resistência insulínica)
- Perfil lipídico: TG, HDL, LDL, ApoB
- Marcadores inflamatórios: PCR-us, homocisteína
- Vitamina D (deficiência comum em SOP)

**Abordagem Funcional Integrativa para SOP (foco em CAUSA-RAIZ = resistência insulínica):**

**Pilar 1 - Reversão de Resistência Insulínica:**
- Dieta low-carb (50-100 g/dia) ou cetogênica: reduz insulina, restaura ovulação em 60-70%
- Jejum intermitente 16/8
- Eliminar açúcares, farinhas refinadas, ultraprocessados
- Priorizar proteínas, gorduras saudáveis, vegetais fibrosos
- Exercício: musculação 3-4x/sem + aeróbico 150 min/sem

**Pilar 2 - Suplementação:**
- **Mio-inositol 2 g + D-quiro-inositol 50 mg** (ratio 40:1): melhora sensibilidade insulínica, restaura ovulação (40-60%), reduz testosterona, equivalente/superior a metformina sem efeitos colaterais
- **Berberina** 500 mg 3x/dia: eficácia similar à metformina, restaura ovulação, reduz andrógenos
- **NAC** (N-acetilcisteína) 1.200-1.800 mg/dia: melhora taxa ovulação, qualidade oócitos, antioxidante
- **Vitamina D3**: meta >50 ng/mL (suplementar 5.000-10.000 UI/dia) - receptores VDR em ovários, melhora função reprodutiva
- **Magnésio** 400-600 mg/dia: sensibilidade insulínica
- **Ômega-3** 2-3 g/dia: anti-inflamatório, reduz andrógenos
- **Cromo** 200-400 mcg/dia: potencializa insulina
- **Canela de Ceilão** 1-3 g/dia: sensibilizador insulínico

**Pilar 3 - Modulação Hormonal:**
- **Vitex agnus-castus** 400 mg/dia (se progesterona baixa): normaliza LH/FSH
- **Saw palmetto** 160 mg 2x/dia: bloqueia 5-alfa-redutase (conversão testosterona→DHT), reduz hirsutismo
- **Spearmint tea** (chá hortelã) 2 xícaras/dia: anti-androgênico, reduz hirsutismo

**Pilar 4 - Correção de Disbiose/Permeabilidade Intestinal:**
Endotoxemia (LPS) perpetua resistência insulínica e inflamação. Probióticos, prebióticos, glutamina, zinco-carnosina.

**Pilar 5 - Farmacoterapia (se necessário):**
- **Metformina** 500-1.000 mg 2x/dia: restaura ovulação em 50%, reduz risco diabetes. Combinação metformina + mio-inositol > monoterapia
- **Anticoncepcional oral** (se não desejar gravidez): reduz andrógenos, regulariza ciclo, protege endométrio (anovulação crônica = risco Ca endométrio). Preferir formulações com progestinas anti-androgênicas (drospirenona, ciproterona). **ATENÇÃO:** ACO mascara problema, não trata causa-raiz - sempre associar intervenção funcional

**Meta terapêutica:** HOMA-IR <1.0, testosterona livre normalizada, restauração de ciclos ovulatórios regulares (28-35 dias), melhora hirsutismo/acne.

---

**2. MIOMAS UTERINOS:**

Conduta depende de: sintomas, tamanho, localização, desejo reprodutivo, idade.

**Assintomáticos:** Seguimento anual com USTV. Abordagem funcional preventiva:
- Reduzir estrogênio: DIM 200-400 mg/dia, I3C, vegetais crucíferos, cálcio-D-glucarato 500-1.000 mg
- Reduzir IGF-1: dieta low-carb, berberina, metformina
- EGCG (chá verde) 800 mg/dia: inibe proliferação células mioma
- Vitamina D3: meta >50 ng/mL (níveis adequados associados a menor tamanho miomas)
- Curcumina 1 g/dia: anti-proliferativa

**Sintomáticos (menorragia → anemia, dor, compressão):**
Tratamento médico:
- Ácido tranexâmico 1-1,5 g 3x/dia (dias de sangramento): reduz perda sanguínea 40-50% (inibe fibrinólise)
- DIU de levonorgestrel (Mirena): reduz sangramento 90%, atrofia endometrial
- Agonistas GnRH (leuprorelina, gos erelina): induz menopausa temporária, reduz miomas 30-60% (pré-operatório, máximo 6 meses - perda óssea)

Se anemia: ferro bisglicinato 50-100 mg/dia, vitamina C 500 mg, B12, folato.

Tratamento definitivo (se refratário, miomas grandes >6 cm, desejo de preservar útero vs não):
- **Miomectomia** (cirúrgica/laparoscópica): remove miomas, preserva útero
- **Embolização artérias uterinas**: reduz fluxo sanguíneo, miomas necrosam/diminuem 40-60%
- **Histerectomia**: cura definitiva (se prole completa, >40 anos, refratário)

---

**3. ADENOMIOSE:**

Não há cura, apenas controle sintomático. Abordagem funcional:
- **Modulação estrogênica:** DIM, I3C, cálcio-D-glucarato
- **Anti-inflamatórios naturais:** ômega-3 4 g/dia, curcumina 1 g/dia, gengibre 1-2 g/dia, resveratrol 500 mg
- **Suporte hormonal:** progesterona micronizada 200-400 mg/dia (dias 14-28 ciclo) ou DIU Mirena (reduz sangramento, dor)
- **AINEs:** ibuprofeno 600 mg ou naproxeno 500 mg (início menstruação, 3-5 dias)
- Se refratário e prole completa: histerectomia é curativa

---

**4. PÓLIPOS ENDOMETRIAIS:**

**Qualquer pólipo em mulher com sangramento OU >1,5 cm OU pós-menopausa:** polipectomia histeroscópica com biópsia (descartar malignidade).

Pós-ressecção: modular estrogênio (DIM, perda peso se obesa, reverter resistência insulínica).

---

**5. ESPESSAMENTO ENDOMETRIAL:**

**Pós-menopausa com sangramento + eco >5 mm OU sem sangramento + eco >11 mm:**
**Histeroscopia + biópsia endometrial obrigatória** (descartar hiperplasia/câncer).

Se hiperplasia sem atipia: progesterona micronizada 200-400 mg/dia (14 dias/mês) + abordagem funcional (DIM, perda peso). Repetir biópsia em 3-6 meses.

Se hiperplasia com atipia: risco 25-40% Ca - histerectomia se prole completa OU progestina alta dose (acetato medroxiprogesterona 10-20 mg/dia) se deseja fertilidade.

Se câncer: estadiamento + histerectomia ± quimio/rádio conforme estágio.

---

**6. CISTOS OVARIANOS:**

**Cistos funcionais simples <5 cm:** Reavaliar USG em 6-12 semanas (pós-menstrual). Geralmente regridem. ACO pode suprimir formação de novos cistos (mas não trata os existentes).

**Cistos persistentes/complexos/grandes >5 cm ou sintomáticos:**
- Marcadores tumorais: CA-125, HE4, índice ROMA (pós-menopausa)
- Se baixo risco: seguimento USG 3-6 meses
- Se alto risco malignidade (escore IOTA, O-RADS 4-5): encaminhar ginecologista-oncologista, considerar cirurgia

**Endometriomas:** Tratamento endometriose (ACO contínuo, progestinas, agonistas GnRH). Se >4 cm, dor refratária ou infertilidade: cistectomia laparoscópica.

**Dermoides (teratomas):** Cirurgia (cistectomia) - risco de torção/ruptura."""
    }

def get_content_doppler_carotidas():
    return {
        "clinical": """O Doppler colorido de carótidas e vertebrais é método não-invasivo de ultrassom que combina imagem em modo B (morfologia vascular) com Doppler pulsado/colorido (análise hemodinâmica) para avaliar artérias carótidas comuns, carótidas internas/externas e artérias vertebrais bilateralmente. É ferramenta fundamental na estratificação de risco cardiovascular e prevenção primária/secundária de AVC isquêmico (85% dos AVCs são isquêmicos, 15-20% por doença carotídea).

Na medicina funcional integrativa, o Doppler carotídeo transcende o rastreamento de estenose hemodinamicamente significativa: avaliamos **aterosclerose subclínica** através da **espessura íntima-média carotídea (CIMT - carotid intima-media thickness)**, marcador precoce de risco cardiovascular validado que detecta alterações arteriais décadas antes de eventos clínicos. CIMT é medido na parede posterior da carótida comum (1 cm proximal à bifurcação) bilateralmente. Valores normais: <0,9 mm; limítrofe 0,9-1,0 mm; aumentado >1,0 mm (indica risco CV 2-4x maior). CIMT reflete exposição cumulativa a fatores de risco aterogênicos: LDL oxidado, hipertensão, tabagismo, diabetes, inflamação crônica, homocisteína elevada.

**Placas ateroscleróticas** são detectadas como espessamentos focais >1,5 mm ou >50% do CIMT adjacente. Classificação: **estáveis** (ricas em tecido fibroso, calcificadas, ecogênicas, superfície regular) vs **instáveis/vulneráveis** (ricas em lipídio, núcleo necrótico, capa fibrosa fina, hipoecóicas, superfície irregular/ulcerada - alto risco de ruptura → trombo → AVC). Ulcerações/irregularidades superficiais aumentam risco embólico 5-7x independente do grau de estenose.

**Grau de estenose carotídea** (redução diâmetro luminal) é quantificado por velocidade de pico sistólico (PSV - peak systolic velocity):
- **Normal:** PSV <125 cm/s, estenose <50%
- **Estenose moderada (50-69%):** PSV 125-230 cm/s
- **Estenose severa (70-99%):** PSV >230 cm/s, ratio PSV carótida interna/comum >4,0
- **Oclusão:** ausência de fluxo

Fisiopatologicamente, estenose >50% raramente causa sintomas (circulação colateral compensa). Estenose >70% reduz perfusão cerebral sob estresse hemodinâmico (hipotensão, exercício). Placas instáveis causam eventos através de **embolização distal** (fragmentos de placa/trombo migram → oclusão artérias cerebrais médias/anteriores → AVC) ou **trombose in situ** (ruptura placa → cascata coagulação → oclusão aguda).

Pacientes com estenose carotídea 70-99% **sintomáticos** (AIT ou AVC nos últimos 6 meses) têm risco de AVC ipsilateral 13% ao ano - benefício absoluto de endarterectomia é 16% em 5 anos (NNT=6). Estenose severa **assintomática** tem risco ~2%/ano - endarterectomia reduz para 1%/ano (NNT=50 em 5 anos, benefício menor).

Na abordagem funcional, pacientes com CIMT aumentado ou placas não-obstrutivas (<50%) requerem **otimização agressiva de fatores de risco** e **terapia anti-aterosclerótica intensiva** para estabilizar/regredir placas e prevenir progressão.""",

        "patient": """O Doppler de carótidas é um ultrassom que avalia as artérias do pescoço (carótidas e vertebrais) que levam sangue para o cérebro. É exame indolor, sem radiação, que mede: 1) espessura da parede das artérias, 2) presença de placas de gordura/cálcio, 3) estreitamento (estenose) que pode reduzir fluxo sanguíneo para o cérebro.

**Por que é importante:** Derrame (AVC) é a segunda maior causa de morte no Brasil e principal causa de incapacidade. 80-90% dos AVCs podem ser prevenidos identificando e tratando fatores de risco antes que ocorra o evento. O Doppler de carótidas detecta aterosclerose ("entupimento" das artérias) décadas antes do derrame.

**CIMT (espessura da parede):** Valor normal <0,9 mm. Acima disso indica que suas artérias estão "envelhecendo" mais rápido, aumentando risco de infarto e derrame 2-4 vezes.

**Placas:** Depósitos de gordura e cálcio na parede arterial. Podem ser estáveis (duras, calcificadas) ou instáveis (moles, com risco de romper e causar derrame). Algumas placas estreitam a artéria (<50% = leve; 50-70% = moderada; >70% = severa).

**Quando suspeitar:** Se você tem pressão alta, diabetes, colesterol alterado, histórico familiar de derrame/infarto, idade >50 anos, tabagismo, ou já teve sintomas neurológicos (fraqueza/formigamento em um lado do corpo, dificuldade de fala, perda de visão temporária).

O exame ajuda a decidir intensidade do tratamento preventivo. Pessoas com placas ou CIMT aumentado precisam de controle muito mais rigoroso do colesterol, pressão e outros fatores de risco.""",

        "conduct": """**Indicações Doppler Carótidas:**
AVC/AIT prévio (investigação etiológica), sopro carotídeo ausculta, sintomas neurológicos focais ou hemisféricos, triagem em pacientes alto risco CV (diabetes >50 anos, doença aterosclerótica em outro território - coronário/periférico, múltiplos fatores de risco), homens >65 anos + tabagismo + ≥1 fator risco adicional.

**Interpretação e Conduta:**

---

**1. CIMT AUMENTADO (>0,9-1,0 mm) SEM ESTENOSE SIGNIFICATIVA:**

Indica aterosclerose subclínica - paciente está em **alto risco CV** mesmo sem obstrução. Intervenção funcional intensiva:

**A. Otimização Lipídica Avançada:**
- **Meta:** LDL <70 mg/dL (idealmente <55 se alto risco), ApoB <80 mg/dL, partículas LDL-P <1.000 nmol/L, triglicerídeos <100 mg/dL, HDL >50 mg/dL (mulheres) ou >40 (homens), ratio TG/HDL <2,0, Lp(a) <30 mg/dL
- **Farmacoterapia:** Estatina alta potência (rosuvastatina 20-40 mg ou atorvastatina 40-80 mg) + ezetimiba 10 mg se LDL não na meta. Estudos mostram regressão/estabilização de CIMT com estatinas em 2-3 anos
- **Suplementação lipídica:**
  - **Ômega-3 EPA** (icosapent etil/Vascepa) 4 g/dia: reduz eventos CV 25% (estudo REDUCE-IT), anti-inflamatório potente
  - **Bergamota** (citrus bergamia) 500-1.000 mg/dia: reduz LDL oxidado, aumenta HDL
  - **Niacina** 500-2.000 mg/dia (liberação estendida): aumenta HDL 20-30%, reduz Lp(a), TG. Monitorar TGO/TGP
  - **Berberina** 500 mg 3x/dia: reduz LDL 20-30% (via up-regulation receptores LDL hepáticos)

**B. Controle Pressórico Rigoroso:**
Meta <130/80 mmHg (ou <120/80 se tolerado). Hipertensão acelera CIMT 0,01-0,02 mm/ano. IECA ou BRA preferíveis (bloqueiam angiotensina II - pró-inflamatória, pró-fibrótica). Suplementação: magnésio 600 mg, ômega-3, CoQ10 200 mg, alho envelhecido 1.200 mg, L-arginina 3-6 g (precursor NO - vasodilatador endógeno).

**C. Reversão de Resistência Insulínica:**
Hiperinsulinemia promove aterosclerose via múltiplas vias (inflamação, proliferação células musculares lisas, oxidação LDL). Testar insulina jejum, HOMA-IR, HbA1c. Meta: HOMA-IR <1,0, HbA1c <5,3%. Intervenção: dieta low-carb/cetogênica, jejum intermitente 16/8, berberina 500 mg 3x/dia ou metformina 500-1.000 mg 2x/dia, exercício 150 min/sem.

**D. Redução de Homocisteína:**
Homocisteína >10 µmol/L é fator de risco independente para aterosclerose (disfunção endotelial, estresse oxidativo). Meta <7 µmol/L. Suplementação: metilfolato (5-MTHF) 1-5 mg/dia, metilcobalamina (B12) 1.000-2.000 mcg/dia, P-5-P (vitamina B6 ativa) 50-100 mg/dia.

**E. Anti-inflamatórios Potentes:**
Inflamação crônica (PCR-us >3 mg/L) promove instabilidade placa. Testar PCR-us, Lp-PLA2 (marcador específico de inflamação vascular). Meta PCR-us <1 mg/L.
- **Curcumina** 1.000 mg/dia (inibe NF-kB, reduz PCR 50-60%)
- **Resveratrol** 500 mg/dia (ativa sirtuínas, antioxidante)
- **Ômega-3** 2-4 g/dia (reduz TNF-alfa, IL-6, PCR)
- Dieta anti-inflamatória (eliminar ultraprocessados, açúcares, óleos vegetais refinados)

**F. Suporte Endotelial (melhora função vasodilatadora, reduz disfunção endotelial):**
- **L-arginina** 3-6 g/dia (precursor óxido nítrico - NO)
- **L-citrulina** 3-6 g/dia (convertida em arginina endogenamente, biodisponibilidade superior)
- **CoQ10 (ubiquinol)** 200-400 mg/dia (antioxidante mitocondrial, melhora função endotelial; reposição obrigatória em uso de estatinas)
- **Vitamina K2 (MK-7)** 180-360 mcg/dia (ativa MGP - matrix Gla protein - remove cálcio vascular, inibe calcificação arterial)

**G. Antioxidantes:**
- **Vitamina E** (mixed tocoferóis) 400 UI/dia
- **Vitamina C** 1.000 mg/dia
- **Ácido alfa-lipóico** 600 mg/dia
- **NAC** 1.200 mg/dia (precursor glutationa)

**H. Estilo de Vida:**
- **Exercício aeróbico:** 150 min/sem (melhora função endotelial, reduz CIMT)
- **Cessar tabagismo:** tabaco acelera CIMT 0,02-0,03 mm/ano
- **Dieta mediterrânea/low-carb:** vegetais, azeite extravirgem, peixes gordos, oleaginosas
- **Manejo estresse:** cortisol crônico elevado promove aterosclerose
- **Sono 7-9h/noite**

**Seguimento:** Repetir Doppler carótidas em 12-24 meses para avaliar estabilização/regressão de CIMT e placas.

---

**2. ESTENOSE CAROTÍDEA 50-69% (MODERADA):**

**Farmacoterapia intensiva obrigatória (reduz progressão e eventos):**
- **AAS 81-100 mg/dia** (antiagregante plaquetário)
- **Estatina alta potência** (rosuvastatina 20-40 mg): meta LDL <55 mg/dL
- **IECA ou BRA:** controle PA <130/80 mmHg
- Todas as estratégias funcionais acima (seção 1)

**Endarterectomia NÃO recomendada** em assintomáticos (risco cirúrgico >benefício). Apenas tratamento clínico otimizado + seguimento Doppler anual.

---

**3. ESTENOSE CAROTÍDEA 70-99% (SEVERA):**

**Se SINTOMÁTICO (AVC/AIT nos últimos 6 meses no território carotídeo):**
- **Endarterectomia carotídea (CEA)** ou **angioplastia + stent carotídeo (CAS)** dentro de 2 semanas (quanto antes, melhor - janela terapêutica)
- CEA preferível se: anatomia favorável, idade <70 anos, ausência de comorbidades graves. CAS se: anatomia desfavorável (carótida alta, ressecção cervical prévia), múltiplas comorbidades
- Peri-operatório: AAS + clopidogrel
- Pós-procedimento: AAS indefinidamente + estatina + controle fatores risco

**Se ASSINTOMÁTICO:**
Considerar CEA/CAS **seletivamente** em:
- Expectativa de vida >5 anos
- Centro cirúrgico com taxa complicação <3%
- Estenose >80% com placa instável (hipoecóica, irregular, ulcerada)
- Progressão documentada em exames seriados
- Reserva cerebrovascular reduzida (teste CO2)

**Alternativa:** Tratamento clínico otimizado (AAS + estatina + IECA + otimização funcional) + seguimento Doppler 6-12 meses. Risco anual de AVC ~2%/ano com tratamento médico atual.

---

**4. PLACAS NÃO-OBSTRUTIVAS (<50%) MAS PRESENTES:**

Mesmo sem estenose significativa, presença de placas indica doença aterosclerótica sistêmica. Implementar estratégia preventiva agressiva (mesma da seção 1). Objetivo: estabilizar placas (aumentar capa fibrosa, reduzir núcleo lipídico), prevenir novas placas, reduzir inflamação.

Repetir Doppler em 12-24 meses."""
    }

def get_content_doppler_aorta_renais():
    return {
        "clinical": """O Doppler colorido de aorta abdominal e artérias renais combina ultrassom modo B (imagem anatômica) com Doppler pulsado/colorido (análise hemodinâmica) para avaliar: 1) aorta abdominal (diâmetro, presença de aneurisma, aterosclerose, dissecção), 2) artérias renais principais (estenose, fluxo, índices de resistividade), 3) perfusão renal intraparenquimatosa.

**ANEURISMA DE AORTA ABDOMINAL (AAA):** dilatação focal >50% do diâmetro normal ou diâmetro absoluto >3,0 cm. Prevalência: 4-8% homens >65 anos, 1% mulheres. Fatores de risco: tabagismo (RR 5-7x - principal fator modificável), idade >65 anos, sexo masculino, histórico familiar (RR 2-4x se parente de primeiro grau), hipertensão, dislipidemia, aterosclerose em outros territórios, DPOC. Patogênese: degradação de elastina e colágeno na camada média por metaloproteinases (MMPs), inflamação crônica de baixo grau, estresse oxidativo, disfunção endotelial.

Classificação por diâmetro: **pequeno** 3,0-3,9 cm (risco ruptura <1%/ano), **médio** 4,0-5,4 cm (1-5%/ano), **grande** ≥5,5 cm (risco ruptura 10-20%/ano). Taxa de expansão média: 2-3 mm/ano (acelerada em tabagistas, hipertensos, diâmetro inicial maior). Ruptura de AAA tem mortalidade 80-90% (50% morrem antes de chegar ao hospital, mortalidade cirúrgica emergencial 40-50%).

Rastreamento com USG abdominal **recomendado uma vez** em homens 65-75 anos com histórico de tabagismo (USPSTF grau B - reduz mortalidade específica por AAA em 40%). Mulheres: rastreamento individualizado se tabagismo + histórico familiar.

**ESTENOSE DE ARTÉRIA RENAL (EAR):** estreitamento ≥50% do lúmen arterial, causa de **hipertensão renovascular** (5-10% de todas HAS, mais comum em HAS resistente a 3+ drogas). Etiologias: **aterosclerose** (90% - proximal/óstio, bilateral em 30%, pacientes >50 anos com aterosclerose sistêmica), **displasia fibromuscular** (10% - mulheres jovens <50 anos, terço médio/distal, aspecto "colar de contas", não-progressiva).

Fisiopatologia: hipoperfusão renal → ativação eixo renina-angiotensina-aldosterona (SRAA) → hipertensão, retenção Na+, aldosteronismo secundário, hipertrofia ventricular esquerda, insuficiência cardíaca. Estenose bilateral ou em rim único → insuficiência renal progressiva por isquemia crônica.

Doppler detecta EAR através de critérios hemodinâmicos: **PSV artéria renal** (peak systolic velocity - velocidade pico sistólico) >180-200 cm/s indica estenose ≥60%; **RAR** (renal-aortic ratio = ratio PSV artéria renal/PSV aorta) ≥3,5 sugere estenose ≥60%; **aceleração tardio** (tardus parvus) no Doppler intraparenquimatoso (tempo de aceleração >70 ms). Sensibilidade 85%, especificidade 90% para EAR ≥60%.

Limitações do Doppler: operador-dependente, difícil em obesos, gases intestinais. Angiotomografia ou angioressonância são padrões-ouro se Doppler inconclusivo.""",

        "patient": """O Doppler de aorta e artérias renais é um ultrassom que avalia: 1) a aorta abdominal (artéria principal que leva sangue do coração para pernas e órgãos abdominais), 2) as artérias renais (artérias que levam sangue para os rins).

**Aneurisma de aorta:** É uma dilatação ("balão") na parede da aorta. Se pequeno, não causa sintomas e cresce lentamente ao longo dos anos. O perigo é a ruptura: se o aneurisma crescer muito (geralmente >5,5 cm) e romper, causa hemorragia interna massiva, com risco de morte de 80-90%. Por isso, aneurismas grandes precisam ser corrigidos preventivamente com cirurgia.

Fatores de risco: cigarro (principal), pressão alta, idade >65 anos, histórico familiar. O exame é recomendado uma vez para todos os homens entre 65-75 anos que fumaram na vida.

**Estreitamento (estenose) das artérias renais:** Pode causar pressão alta difícil de controlar (mesmo com vários remédios) e prejudicar a função dos rins. O estreitamento reduz o fluxo de sangue para os rins, que "pensam" que a pressão está baixa e liberam hormônios que aumentam a pressão do corpo todo.

**Quando suspeitar de estenose renal:** Pressão alta que começou antes dos 30 anos ou depois dos 55 anos, pressão alta resistente (não melhora com 3+ remédios), piora súbita de pressão previamente controlada, sopro abdominal na ausculta médica, assimetria no tamanho dos rins.

Exame indolor, sem radiação, dura 20-30 minutos. Requer jejum de 6-8 horas (gases intestinais atrapalham visualização).""",

        "conduct": """**Indicações Doppler Aorta Abdominal + Artérias Renais:**
- Rastreamento AAA: homens 65-75 anos com histórico tabagismo (rastreamento único)
- Sopro abdominal, massa pulsátil palpável, dor abdominal/lombar sugestiva
- HAS resistente (PA não controlada com ≥3 anti-hipertensivos incluindo diurético), HAS de início recente <30 anos ou >55 anos, HAS acelerada/maligna, piora abrupta de HAS previamente controlada
- Assimetria renal (rim >1,5 cm menor que contralateral), IRC progressiva após IECA/BRA, edema pulmonar flash recorrente

---

**INTERPRETAÇÃO E CONDUTA:**

**1. ANEURISMA DE AORTA ABDOMINAL (AAA):**

**Diâmetro 2,5-2,9 cm (aorta normal-alta):**
Seguimento não necessário, mas reforçar prevenção primária (cessar tabagismo, controle HAS/dislipidemia).

**Diâmetro 3,0-3,9 cm (AAA pequeno):**
- Seguimento USG a cada 3 anos
- Abordagem funcional preventiva (objetivo: reduzir taxa de expansão e risco ruptura)

**Diâmetro 4,0-5,4 cm (AAA médio):**
- Seguimento USG a cada 6-12 meses
- Considerar angioTC se morfologia complexa/trombo
- Abordagem funcional intensiva

**Diâmetro ≥5,5 cm (homens) ou ≥5,0 cm (mulheres) OU expansão >0,5 cm/6 meses OU sintomático:**
- **Correção cirúrgica eletiva indicada** (benefício sobrevida comprovado)
- Duas opções:
  - **EVAR (endovascular):** stent-graft via femoral, menos invasivo, recuperação rápida, mortalidade 1-2%, preferível se anatomia favorável + comorbidades. Requer seguimento imaging vida toda (endoleaks)
  - **Cirurgia aberta:** substituição segmento aórtico por prótese, mais invasiva, mortalidade 3-5%, durável sem necessidade de seguimento imaging
- Pré-operatório: otimizar controle CV (PA, betabloqueador, estatina), cessar tabagismo ≥4 semanas antes

**Abordagem Funcional Integrativa para AAA (reduzir progressão):**

**A. Cessar Tabagismo (MAIS IMPORTANTE):**
Tabagismo acelera expansão AAA 2-3x. Cessação reduz taxa expansão em 30-50%. Suporte: terapia de reposição nicotínica, vareniclina, bupropiona, aconselhamento.

**B. Controle Pressórico:**
Meta PA <130/80 mmHg. Betabloqueadores reduzem estresse parietal (força de cisalhamento na parede aórtica). IECA/BRA podem reduzir atividade MMPs (estudos conflitantes).

**C. Estatinas (evidência robusta):**
Rosuvastatina 20-40 mg ou atorvastatina 40-80 mg: reduzem taxa de expansão AAA em 15-30%, reduzem mortalidade perioperatória, estabilizam placa aterosclerótica. Mecanismo: anti-inflamatório (reduz PCR, MMPs), estabilização parede arterial.

**D. Inibição de Metaloproteinases (MMPs) - estratégia funcional:**
MMPs degradam elastina/colágeno → expansão AAA. Inibidores naturais:
- **Doxiciclina** 100-200 mg/dia (off-label): inibe MMPs independente de ação antibiótica. Estudos preliminares mostram redução expansão AAA
- **Curcumina** 1.000 mg/dia: inibe NF-kB e MMPs
- **EGCG (chá verde)** 400-800 mg/dia: inibe MMPs
- **Resveratrol** 500 mg/dia

**E. Antioxidantes e Anti-inflamatórios:**
Estresse oxidativo e inflamação crônica promovem degradação parede aórtica.
- **Vitamina C** 1.000 mg/dia (síntese colágeno)
- **Vitamina E** 400 UI/dia
- **Ômega-3** 2-4 g/dia (anti-inflamatório, reduz PCR)
- **NAC** 1.200 mg/dia (antioxidante, reduz MMPs)
- **Ácido alfa-lipóico** 600 mg/dia

**F. Suporte Estrutural da Parede Arterial:**
- **Vitamina K2 (MK-7)** 180-360 mcg/dia (inibe calcificação, ativa proteínas matriz)
- **Magnésio** 600 mg/dia (antagonista cálcio, relaxamento vascular)
- **Vitamina D3:** meta >50 ng/mL (imunomodulação, reduz inflamação)
- **Colágeno tipo I/III** hidrolisado 10 g/dia (suporte matriz extracelular)
- **Vitamina C + lisina + prolina:** síntese colágeno (protocolo Rath/Pauling - controverso, sem evidência robusta)

**G. Estilo de Vida:**
- Evitar exercícios isométricos intensos (levantamento peso >50% 1-RM), Valsalva forçada (aumentam pressão intra-abdominal e estresse parietal)
- Exercício aeróbico moderado (caminhada, natação) seguro e recomendado

---

**2. ESTENOSE DE ARTÉRIA RENAL (EAR) ≥60%:**

**Investigação Complementar:**
- Dosagem de renina/aldosterona plasmáticas (renina elevada sugere EAR funcional)
- Função renal: creatinina, TFG (CKD-EPI), ureia, eletrólitos, relação albumina/creatinina urinária
- AngioTC ou angioRM renal (confirmar estenose, avaliar anatomia pré-intervenção)
- Cintilografia renal com captopril (avalia função renal diferencial)

**Conduta Geral:**

**Tratamento Clínico Otimizado (TODOS os pacientes):**
- **Anti-hipertensivos:** IECA ou BRA (bloqueiam SRAA - mecanismo fisiopatológico da HAS renovascular). **ATENÇÃO:** IECA/BRA podem reduzir TFG se estenose bilateral ou rim único - monitorar creatinina/potássio em 1-2 semanas; se creatinina ↑>30%, suspender. Adicionar bloqueadores canais cálcio (amlodipina), diuréticos tiazídicos, betabloqueadores conforme necessário
- **Estatina alta potência:** aterosclerose é causa em 90% - tratar doença de base
- **AAS 81-100 mg/dia**
- Controle glicêmico rigoroso se diabetes
- Cessar tabagismo

**Revascularização (angioplastia + stent renal):**

**Indicações atuais (estudos CORAL, ASTRAL mostram benefício limitado em maioria):**
- **Indicação forte (consenso):**
  - Estenose bilateral >70% ou estenose >70% em rim único + IRC progressiva apesar de tratamento clínico
  - HAS refratária a ≥3 drogas em doses otimizadas + estenose >70%
  - Edema pulmonar flash recorrente (síndrome Pickering)
  - Displasia fibromuscular sintomática (angioplastia SEM stent - alta taxa sucesso)

- **Não indicado (sem benefício):**
  - EAR assintomática descoberta incidentalmente
  - HAS controlada com medicação
  - Função renal estável

**Pós-revascularização:** AAS + clopidogrel 1-3 meses, depois AAS indefinido. Manter IECA/BRA, estatina. Doppler follow-up 6-12 meses (re-estenose 10-20%).

**Abordagem Funcional Adjuvante em EAR:**
Mesma estratégia anti-aterosclerótica descrita para Doppler carótidas (otimização lipídica, anti-inflamatórios, antioxidantes, suporte endotelial) para prevenir progressão e aterosclerose em outros territórios."""
    }

def get_content_usg_prostata():
    return {
        "clinical": """A ultrassonografia de próstata via abdominal (suprapúbica) avalia a glândula prostática através da bexiga cheia, que funciona como "janela acústica". Fornece informações sobre volume prostático, morfologia, ecotextura, resíduo pós-miccional (RPM), bexiga e parede vesical. É exame complementar na investigação de sintomas do trato urinário inferior (STUI), hiperplasia prostática benigna (HPB), rastreamento de câncer de próstata (porém PSA + toque retal são mais sensíveis), retenção urinária.

**USG transretal (USTR)** é mais acurado que USG abdominal para avaliação detalhada da próstata (anatomia zonal, lesões focais, volume preciso, guia de biópsia), mas USG abdominal é não-invasivo e suficiente para triagem inicial e seguimento de HPB.

**HIPERPLASIA PROSTÁTICA BENIGNA (HPB):** crescimento benigno da zona de transição prostática, praticamente universal em homens >50 anos (50% aos 60 anos, 90% aos 85 anos). Volume prostático normal: 20-25 mL. HPB definida como volume >30 mL. Crescimento médio: 0,6 mL/ano (acelerado em obesos, resistência insulínica, síndrome metabólica).

Fisiopatologia: conversão de testosterona em **di-hidrotestosterona (DHT)** pela enzima **5-alfa-redutase** (tipo 2, expressa na próstata). DHT é 5-10x mais potente que testosterona em receptores androgênicos, estimulando proliferação celular prostática. Fatores contribuintes: envelhecimento (↓ testosterona, ↑ estrogênio relativo - aromatização periférica), resistência insulínica (IGF-1 estimula crescimento prostático), inflamação crônica de baixo grau, estresse oxidativo.

**Sintomas STUI:** divididos em **obstrutivos** (jato fraco, hesitação, gotejamento terminal, esvaziamento incompleto, esforço miccional) e **irritativos** (urgência, frequência aumentada, noctúria, incontinência de urgência). Escore IPSS (International Prostate Symptom Score): 0-7 leve, 8-19 moderado, 20-35 severo. IPSS ≥8 + volume >30 mL indica HPB sintomática.

**Resíduo pós-miccional (RPM):** volume urinário residual após micção. Normal <50 mL. RPM 50-200 mL sugere obstrução leve-moderada; >200 mL indica obstrução significativa com risco de retenção aguda, ITUs recorrentes, hidronefrose, insuficiência renal.

**CÂNCER DE PRÓSTATA (CaP):** neoplasia mais comum em homens (exceto pele), segunda causa de morte oncológica masculina. Rastreamento controverso: PSA + toque retal a partir dos 50 anos (45 se afrodescendente ou histórico familiar). PSA >4 ng/mL OU toque retal alterado OU PSA velocity >0,75 ng/mL/ano = indica biópsia prostática (USTR-guiada, 12+ fragmentos).

USG abdominal detecta próstata aumentada/heterogênea mas NÃO diagnostica câncer (60-70% CaPs são isoecóicos, indistinguíveis). RM multiparamétrica (mpMRI) + biópsia-fusão é padrão atual para detecção de lesões suspeitas.

Na medicina funcional integrativa, avaliamos não apenas HPB/CaP, mas fatores de risco modificáveis: resistência insulínica (IGF-1, insulina estimulam crescimento prostático), inflamação crônica (prostatite crônica assintomática presente em 20-30% biópsias), desequilíbrios hormonais (excesso estrogênio/DHT, deficiência testosterona livre), exposição a xenoestrogênios, deficiências nutricionais (zinco, selênio, vitamina D, licopeno).""",

        "patient": """A ultrassonografia de próstata via abdominal é um exame que visualiza a próstata através da barriga, usando a bexiga cheia como "janela". A próstata é uma glândula do tamanho de uma noz, localizada abaixo da bexiga, que envolve a uretra (canal por onde sai a urina). Sua função é produzir parte do líquido seminal.

Com o envelhecimento, a próstata tende a crescer (hiperplasia prostática benigna - HPB), podendo apertar a uretra e dificultar a passagem da urina. Sintomas comuns de próstata aumentada:
- Jato urinário fraco
- Dificuldade para começar a urinar
- Sensação de esvaziamento incompleto da bexiga
- Acordar várias vezes à noite para urinar (noctúria)
- Urgência (vontade súbita e forte de urinar)

O exame mede: 1) tamanho da próstata (volume em mL - normal ~20 mL, aumentada >30 mL), 2) resíduo pós-miccional (quanto de urina fica "presa" após urinar - normal <50 mL), 3) espessura da parede da bexiga (se muito espessa, indica esforço crônico contra obstrução).

**Importante:** USG abdominal NÃO detecta câncer de próstata. Para rastreamento de câncer são necessários: PSA (exame de sangue) + toque retal. Se PSA elevado ou toque alterado → biópsia prostática guiada por ultrassom transretal.

O exame requer bexiga cheia (beber 4-6 copos água 1h antes e não urinar). Totalmente indolor, sem radiação.""",

        "conduct": """**Indicações USG Próstata Abdominal:**
STUI (sintomas do trato urinário inferior), avaliação de HPB (volume prostático, resposta terapêutica), resíduo pós-miccional, hematúria, ITUs de repetição em homens, retenção urinária, seguimento pós-RTU de próstata.

**Interpretação e Conduta:**

---

**1. PRÓSTATA NORMAL (volume <30 mL, RPM <50 mL, sem STUI significativos):**

Seguimento de rotina. Rastreamento CaP conforme guidelines (PSA + toque retal anual a partir 50 anos, ou 45 anos se afrodescendente/histórico familiar).

**Prevenção Funcional Integrativa (reduzir risco HPB e CaP):**

**A. Modulação Hormonal (reduzir DHT/estrogênio, otimizar testosterona livre):**
- **Saw palmetto** (Serenoa repens) 320 mg/dia: inibe 5-alfa-redutase (conversão testosterona→DHT), reduz sintomas STUI leve-moderado em 20-40% (eficácia similar a finasterida 5 mg em alguns estudos, sem efeitos colaterais sexuais)
- **Beta-sitosterol** 60-130 mg/dia: fitosterol que melhora fluxo urinário
- **Pygeum africanum** 100-200 mg/dia: anti-inflamatório, anti-proliferativo prostático
- **Urtica dioica** (urtiga) 300-600 mg/dia: inibe ligação DHT a receptores androgênicos
- **DIM** 200-300 mg/dia: metaboliza estrogênio em formas menos ativas (2-OH vs 16-OH), reduz ratio estrogênio/testosterona

**B. Anti-inflamatórios (prostatite crônica subclínica comum):**
- **Quercetina** 500 mg 2x/dia: flavonoide com potente ação anti-inflamatória prostática (estudos mostram melhora significativa STUI/dor pélvica crônica)
- **Curcumina** 500-1.000 mg/dia: inibe NF-kB, COX-2
- **Ômega-3** EPA+DHA 2-3 g/dia: reduz inflamação prostática
- **Cernitin** (pólen): reduz inflamação, melhora STUI

**C. Antioxidantes (proteção contra dano oxidativo e CaP):**
- **Licopeno** 15-30 mg/dia: carotenoide do tomate, concentra-se na próstata. Estudos epidemiológicos: ↓40-50% risco CaP agressivo. Fontes: tomate cozido, molho tomate (biodisponibilidade superior a tomate cru)
- **Selênio** 200 mcg/dia: antioxidante, estudos (SELECT trial) controversos - possível benefício em homens com selênio basal baixo. Não exceder 200 mcg (toxicidade)
- **Vitamina E** (mixed tocoferóis) 400 UI/dia: antioxidante. SELECT trial mostrou leve ↑ risco CaP com alfa-tocoferol isolado - usar mixed tocoferóis
- **Zinco** 30-50 mg/dia: concentração prostática é mais alta que qualquer tecido. Deficiência associada a HPB e CaP. Suplementar se deficiente

**D. Vitamina D:**
Receptores VDR expressos na próstata. Níveis adequados (>50 ng/mL) associados a menor risco CaP agressivo. Suplementar D3 5.000-10.000 UI/dia conforme nível basal.

**E. Chá Verde (EGCG):**
400-800 mg EGCG/dia ou 3-5 xícaras chá verde. Estudos mostram redução progressão CaP low-grade (vigilância ativa).

**F. Reversão de Resistência Insulínica:**
Hiperinsulinemia/IGF-1 elevado promovem proliferação prostática. Dieta low-carb, jejum intermitente, exercício, berberina 500 mg 3x/dia. Meta HOMA-IR <1,0.

**G. Evitar Xenoestrogênios:**
Plásticos (BPA), pesticidas, parabenos (disruptores endócrinos que mimetizam estrogênio).

**H. Atividade Física:**
Exercício regular (150 min/sem) reduz risco HPB sintomática em 25% e CaP agressivo em 30-50%. Mecanismo: reduz insulina, IGF-1, inflamação, obesidade.

**I. Ejaculação Regular:**
Estudos sugerem ejaculação frequente (≥21x/mês) associada a 20-30% ↓ risco CaP (possível mecanismo: clearance de carcinógenos/células senescentes via fluido prostático).

---

**2. HPB SINTOMÁTICA (volume 30-80 mL, IPSS 8-19 moderado, RPM <200 mL):**

**Primeira Linha - Tratamento Funcional + Fitoterapia:**
Implementar TODAS as estratégias da seção 1 (A-I). Ensaio clínico 12 semanas:
- Saw palmetto 320 mg/dia
- Beta-sitosterol 130 mg/dia
- Quercetina 500 mg 2x/dia
- Ômega-3 2 g/dia
- Licopeno 30 mg/dia
- Vitamina D3 (meta >50 ng/mL)
- DIM 300 mg/dia
- Perda de peso se obeso (reduz volume prostático)

Reavaliar IPSS em 12 semanas. Se melhora ≥30% → continuar. Se resposta insuficiente → adicionar farmacoterapia.

**Segunda Linha - Farmacoterapia:**
- **Alfabloqueadores** (tansulosina 0,4 mg/dia, alfuzosina 10 mg/dia, doxazosina 4-8 mg/dia, silodosina 8 mg/dia): relaxam musculatura lisa prostática/colo vesical → melhora sintomas em 3-7 dias, sem reduzir volume. Efeitos colaterais: hipotensão ortostática (iniciar dose baixa), ejaculação retrógrada (silodosina/tansulosina)

- **Inibidores 5-alfa-redutase** (finasterida 5 mg/dia, dutasterida 0,5 mg/dia): bloqueiam conversão testosterona→DHT → reduzem volume prostático 20-30% em 6-12 meses, reduzem PSA ~50%, previnem progressão/retenção aguda. Efeitos colaterais: disfunção erétil/libido (5-10%), ginecomastia. Dutasterida inibe tipo 1+2 (finasterida só tipo 2) - mais potente

**Terapia Combinada (alfabloqueador + inibidor 5-alfa-redutase):**
Estudo CombAT: combinação superior a monoterapia em HPB moderada-severa (volume >30 mL + IPSS ≥12). Reduz progressão em 66% vs alfabloqueador isolado.

**Terceira Linha - Procedimentos Minimamente Invasivos (se refratário a farmacoterapia):**
- **TURP** (ressecção transuretral): padrão-ouro, melhora >90%, mas invasivo (anestesia, internação, risco sangramento/estenose/incontinência)
- **Laser** (HoLEP, GreenLight PVP): vaporização/enucleação, menos sangramento, alta no mesmo dia
- **Embolização artérias prostáticas**: reduz volume 30-40%, melhora sintomas, preserva função sexual
- **Rezūm** (vapor de água): ablação térmica, ambulatorial
- **UroLift**: implantes que afastam lobos prostáticos, preserva função sexual (sem ejaculação retrógrada)

---

**3. HPB SEVERA (volume >80 mL, IPSS ≥20, RPM >200 mL, complicações):**

Complicações: retenção urinária aguda, hidronefrose, insuficiência renal, ITUs recorrentes, hematúria recorrente, cálculos vesicais.

**Conduta:**
- Se retenção aguda: sondagem vesical alívio + avaliação urológica urgente
- Farmacoterapia combinada (alfabloqueador + dutasterida/finasterida) por 3-6 meses
- Se resposta inadequada ou complicações persistentes: **cirurgia** (TURP, laser, enucleação prostática aberta se >80-100 mL)

---

**4. PSA ELEVADO ou TOQUE RETAL ALTERADO (suspeita CaP):**

**Encaminhar urologista** para:
- RM multiparamétrica de próstata (detecta lesões PI-RADS 4-5 suspeitas)
- Biópsia prostática guiada por USTR ± fusão MRI (12+ fragmentos)

Enquanto aguarda procedimentos, implementar estratégias funcionais anti-CaP (seção 1)."""
    }

def main():
    print("=" * 80)
    print("BATCH: 20 EXAMES USG E DOPPLER - MEDICINA FUNCIONAL INTEGRATIVA")
    print("=" * 80)

    print("\n[1/3] Autenticando...")
    token = login()
    print("✓ Login realizado")

    print(f"\n[2/3] Processando {len(ITEMS)} items...")
    success = 0
    errors = []

    for idx, (item_id, name) in enumerate(ITEMS, 1):
        try:
            # Determinar tipo de exame e obter conteúdo
            if "USG mamas" in name:
                content = get_content_usg_mamas()
            elif "USG abdome" in name:
                content = get_content_usg_abdome()
            elif "USG transvaginal" in name:
                content = get_content_usg_transvaginal()
            elif "Doppler" in name and ("carótidas" in name or "cervicais" in name):
                content = get_content_doppler_carotidas()
            elif "Doppler" in name and ("aorta" in name or "renais" in name):
                content = get_content_doppler_aorta_renais()
            elif "USG próstata" in name:
                content = get_content_usg_prostata()
            else:
                errors.append((idx, f"{name}: Tipo não mapeado"))
                continue

            if update_item(token, item_id, content["clinical"], content["patient"], content["conduct"]):
                success += 1
                print(f"  ✓ {idx}/{len(ITEMS)}: {name}")
            else:
                errors.append((idx, name))
        except Exception as e:
            errors.append((idx, f"{name}: {str(e)}"))

        time.sleep(0.3)

    print(f"\n[3/3] Finalizado!")
    print("=" * 80)
    print(f"SUCESSO: {success}/{len(ITEMS)} items ({success/len(ITEMS)*100:.1f}%)")
    if errors:
        print(f"ERROS: {len(errors)}")
        for idx, err in errors[:10]:
            print(f"  - Item {idx}: {err}")
    print("=" * 80)

if __name__ == "__main__":
    main()
