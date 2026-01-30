#!/usr/bin/env python3
"""
Batch de 100 exames de imagem - Medicina Funcional Integrativa
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

def get_content(name):
    """Gera conteúdo baseado no tipo de exame"""

    # Templates por categoria
    if "ECG" in name or "Eletrocardiograma" in name:
        return {
            "clinical": f"""O {name} é ferramenta fundamental na avaliação cardiovascular funcional. Na medicina integrativa, avaliamos não apenas arritmias agudas, mas sinais precoces de disfunção: hipertrofia ventricular (sobrecarga pressórica crônica), alterações isquêmicas subclínicas, prolongamento QTc (risco de morte súbita associado a deficiências de magnésio, potássio, hipocalcemia), alterações de repolarização (disfunção mitocondrial, inflamação). O ECG fornece insights sobre eixo HPA (taquicardia sinusal em excesso simpático), deficiências nutricionais (magnésio, potássio, cálcio afetam condução) e impacto de toxinas (metais pesados prolongam QRS).""",

            "patient": f"""O {name} avalia a atividade elétrica do coração. É exame simples, rápido e indolor que detecta problemas no ritmo cardíaco, sinais de sobrecarga do coração, áreas com pouca oxigenação e risco de doenças cardíacas. Alterações podem indicar desde falta de minerais (magnésio, potássio) até problemas mais sérios que precisam de tratamento.""",

            "conduct": f"""Solicitar {name} como triagem cardiovascular em: adultos >40 anos (baseline), hipertensão, diabetes, histórico familiar de doença cardíaca, palpitações, síncope, dispneia. Avaliar: ritmo (sinusal vs arritmias), frequência (ideal 60-80 bpm em repouso), eixo cardíaco, hipertrofia ventricular (Cornell voltage), alterações ST-T, intervalo QTc corrigido (normal <450ms homens, <460ms mulheres). Se alterações: investigar deficiências (magnésio eritrocitário, potássio, cálcio iônico, vitamina D), solicitar ecocardiograma, teste ergométrico. Otimização funcional: ômega-3 2-4g/dia, CoQ10 200-400mg, magnésio 400-600mg, taurina 1-2g, reduzir inflamação sistêmica."""
        }

    elif "Colonoscopia" in name:
        return {
            "clinical": f"""A {name} é padrão-ouro para detecção e prevenção de câncer colorretal (CCR) e avaliação de doenças inflamatórias intestinais (DII). Na medicina funcional, avaliamos além da patologia macroscópica: inflamação de baixo grau, permeabilidade intestinal, disbiose refletida em alterações mucosas. Scores específicos ({name}) quantificam atividade inflamatória em Crohn (SES-CD) e retocolite ulcerativa (Mayo Score), orientando tratamento. Adenomas são lesões pré-malignas - remoção reduz mortalidade por CCR em 50%. Investigamos causas-raiz: dieta pró-inflamatória, disbiose, deficiências nutricionais (vitamina D, folato, B12), exposição a toxinas.""",

            "patient": f"""A {name} examina o intestino grosso usando câmera flexível. Detecta pólipos (que podem virar câncer), inflamação, sangramentos, úlceras e outras alterações. É o melhor exame para prevenir câncer de intestino, pois permite remover pólipos antes de virarem câncer. Geralmente solicitado após 45-50 anos ou antes se houver sintomas (sangramento, dor, mudança no hábito intestinal) ou histórico familiar de câncer intestinal.""",

            "conduct": f"""Indicações {name}: triagem CCR (início 45 anos, 40 se histórico familiar), sangramento retal, anemia ferropriva sem causa, diarreia crônica, suspeita DII. Preparo adequado fundamental (dieta líquida, laxativos). Avaliação funcional: se adenomas/pólipos removidos - investigar causas: vitamina D (meta >50 ng/ml, suplementar 5.000-10.000 UI/dia), folato, cálcio, magnésio; dieta anti-inflamatória (eliminar ultraprocessados, açúcares, carnes processadas; aumentar fibras, vegetais crucíferos); microbioma intestinal (probióticos específicos: Lactobacillus rhamnosus GG, Bifidobacterium longum). Se DII ativa: tratamento médico + suporte funcional (vitamina D, ômega-3 4g/dia, curcumina 1-2g/dia, butirato, glutamina 5-10g/dia, dieta específica de carboidratos). Seguimento conforme achados."""
        }

    elif "Endoscopia" in name:
        return {
            "clinical": f"""A {name} avalia esôfago, estômago e duodeno, detectando: doença do refluxo (esofagite), gastrite/úlceras (H. pylori, AINEs, estresse), doença celíaca (atrofia vilositária duodenal), neoplasias. Na medicina funcional integrativa, investigamos causas-raiz: hipocloridria (produção insuficiente de ácido gástrico, comum >50 anos, deficiência B12/ferro), disbiose (H. pylori, SIBO), permeabilidade intestinal, deficiências nutricionais consequentes (B12, ferro, cálcio, zinco por má absorção). Biópsias duodenais essenciais se suspeita de doença celíaca (sorologia positiva) ou má absorção.""",

            "patient": f"""A {name} examina esôfago, estômago e início do intestino delgado usando câmera flexível. Investiga: azia persistente, dor no estômago, dificuldade de engolir, náuseas, vômitos, sangramento digestivo, anemia sem causa aparente. Pode detectar inflamação, úlceras, infecção por bactéria H. pylori, doença celíaca e tumores. Permite também colher amostras (biópsias) e tratar algumas condições.""",

            "conduct": f"""Indicações {name}: DRGE refratária, dispepsia persistente, disfagia, sangramento digestivo alto, anemia ferropriva, rastreamento Ca gástrico (histórico familiar, etnia asiática >50 anos), triagem celíaca (anti-tTG+). Com biópsias se suspeita celíaca/má absorção. Se H. pylori+: erradicação (terapia quádrupla: IBP + bismuto + tetraciclina + metronidazol 14 dias) + suporte funcional (probióticos durante/após antibióticos, glutamina, DGL, zinco-carnosina). Se gastrite atrófica: suplementar B12 (metilcobalamina 1.000mcg sublingual), ferro se deficiência, considerar HCl-betaína com refeições se hipocloridria sintomática. Se esofagite: elevar cabeceira, evitar refeições 3h antes de dormir, perder peso, suplementar melatonina 3-6mg (protege mucosa esofágica), DGL, alginato. Otimização funcional pós-erradicação H. pylori: restaurar microbioma (probióticos multi-cepas, prebióticos), reparar mucosa (glutamina 5g 2x/dia, zinco-carnosina 75mg 2x/dia, vitamina A 10.000 UI/dia)."""
        }

    elif "Mamografia" in name or "USG mamas" in name:
        return {
            "clinical": f"""O {name} é ferramenta essencial no rastreamento de câncer de mama, segunda causa de morte por câncer em mulheres. Mamografia detecta microcalcificações e massas não palpáveis; USG complementa avaliando natureza de nódulos (sólido vs cístico). Na medicina funcional integrativa, rastreamento é parte de estratégia preventiva abrangente: avaliar fatores de risco modificáveis (obesidade, resistência insulínica, exposição a xenoestrogênios, deficiência vitamina D), biomarcadores hormonais (metabolitos estrogênio - 2OH vs 16OH, ratio progesterona/estrogênio), inflamação sistêmica (PCR-us), função de detoxificação (polimorfismos COMT, MTHFR, GST).""",

            "patient": f"""O {name} é usado para detectar câncer de mama precocemente, quando ainda é pequeno e mais fácil de tratar. Quanto mais cedo o câncer é descoberto, maiores as chances de cura. Geralmente começa aos 40-50 anos e é repetido a cada 1-2 anos. Não previne câncer, mas detecta precocemente. Complementar com autoexame mensal e exame clínico anual. Hábitos saudáveis reduzem risco: peso adequado, exercícios, alimentação anti-inflamatória, evitar álcool/tabaco.""",

            "conduct": f"""Rastreamento com {name}: início 40-50 anos (individualizar baseado em risco), anual ou bienal conforme guidelines. Se histórico familiar forte (mãe/irmã <50 anos): iniciar 10 anos antes da idade do diagnóstico familiar. Se mutação BRCA: RM mamas anual + mamografia. Abordagem funcional preventiva: 1) Otimização hormonal: testar metabolitos de estrogênio (2OH-E1, 16OH-E1) - ratio ideal >2.0, suplementar DIM 200-400mg ou I3C 400mg, brócolis germinado; aumentar progesterona natural (Vitex 400mg se deficiência lútea). 2) Modulação inflamatória: ômega-3 2g/dia, curcumina 1g/dia, vitamina D (meta >50 ng/ml). 3) Detoxificação: suportar fase II hepática (sulforafano, NAC 1.200mg, glicina, glutationa), evitar xenoestrogênios (plásticos, parabenos, ftalatos). 4) Metabolismo: perder excesso de gordura (aromatase em tecido adiposo converte andrógenos em estrogênio), exercício regular 150min/sem. 5) Microbioma: probióticos com Lactobacillus (metabolizam estrogênio via estroboloma). Se BI-RADS 4/5: biópsia urgente."""
        }

    elif "Densitometria" in name or "T-Score" in name:
        return {
            "clinical": f"""A densitometria óssea ({name}) avalia densidade mineral óssea (DMO), diagnosticando osteopenia (T-score -1,0 a -2,5) e osteoporose (T-score ≤-2,5). Na medicina funcional integrativa, avaliamos causas-raiz da perda óssea: deficiência vitamina D (>80% dos osteoporóticos), vitamina K2 (direciona cálcio aos ossos), magnésio (60% do magnésio corporal está nos ossos), proteína inadequada, acidose metabólica crônica (dieta pró-inflamatória mobiliza cálcio ósseo), hipercortisolismo, hipogonadismo, malabsorção, hiperparatireoidismo secundário, inflamação sistêmica (IL-6, TNF-alfa ativam osteoclastos), deficiência de hormônios anabólicos (testosterona, DHEA, estrogênio). Screening de causas secundárias obrigatório.""",

            "patient": f"""A densitometria ({name}) mede a densidade (força) dos seus ossos, principalmente coluna e quadril. T-score acima de -1,0 é normal; entre -1,0 e -2,5 indica ossos enfraquecidos (osteopenia); abaixo de -2,5 indica osteoporose (ossos frágeis com alto risco de fratura). Osteoporose não dói até ocorrer fratura - por isso é importante rastrear e prevenir. Fatores que enfraquecem ossos: falta de cálcio/vitamina D, sedentarismo, tabagismo, excesso de álcool, menopausa, medicamentos (corticoides), baixo peso.""",

            "conduct": f"""Indicações densitometria ({name}): mulheres ≥65 anos, homens ≥70 anos, menopausa precoce (<45 anos), fraturas por trauma mínimo, uso prolongado corticoides (>3 meses), histórico familiar osteoporose, doenças associadas (hiperparatireoidismo, hipertireoidismo, doença celíaca, DII, IRC). Se osteopenia/osteoporose: investigar causas secundárias - exames: vitamina D (meta >50 ng/ml), PTH (excluir hiperparatireoidismo), cálcio iônico, magnésio, TSH, testosterona (homens), função renal, marcadores de reabsorção óssea (CTX). Intervenção funcional integrativa: 1) Vitamina D3: 5.000-10.000 UI/dia (meta >50 ng/ml). 2) Vitamina K2 (MK-7): 180-360mcg/dia (ativa osteocalcina). 3) Magnésio: 400-800mg/dia (glicina to/treonato). 4) Cálcio: 1.000-1.200mg/dia (preferir alimentos - laticínios fermentados, vegetais, sardinhas; suplementar apenas se deficiência). 5) Proteína: 1,2-1,6g/kg/dia (aminoácidos essenciais para matriz óssea). 6) Colágeno tipo I: 5-10g/dia. 7) Exercício resistido 3x/sem + impacto (caminhada, dança). 8) Corrigir acidose: dieta alcalinizante (vegetais, frutas low-glycemic). 9) Se mulher pós-menopausa: considerar TRH bioidêntica (estrogênio + progesterona protegem ossos). 10) Evitar excesso sódio, cafeína, álcool, tabaco. Farmacoterapia (bisfosfonatos, denosumab) se T-score <-2,5 ou fraturas prévias - sempre associar suporte funcional."""
        }

    elif "Doppler" in name:
        return {
            "clinical": f"""O {name} avalia fluxo sanguíneo arterial/venoso usando ultrassom, detectando estenoses, obstruções, insuficiências. Na medicina funcional, Doppler de carótidas/aorta/renais rastreia aterosclerose subclínica (estenose indica doença avançada), hipertensão renovascular. CIMT (espessura íntima-média carotídea) é marcador precoce de aterosclerose - valores >0,9mm indicam risco cardiovascular aumentado. Avaliamos causas-raiz: dislipidemia aterogênica (LDL oxidado, partículas pequenas densas, ApoB elevado), inflamação crônica (PCR-us, Lp-PLA2), resistência insulínica, homocisteína elevada, disfunção endotelial, estresse oxidativo.""",

            "patient": f"""O {name} usa ultrassom para ver como o sangue está fluindo nas artérias e veias. Detecta entupimentos, estreitamentos ou outros problemas nos vasos sanguíneos. Pode avaliar artérias do pescoço (risco de derrame), rins (causa de pressão alta), aorta, pernas. É exame indolor, sem radiação. Importante para prevenir derrame, infarto e outros problemas vasculares graves.""",

            "conduct": f"""Indicações {name}: triagem aterosclerose (>50 anos, fatores de risco CV), AVC/AIT prévio, sopro carotídeo, hipertensão resistente (avaliar artérias renais), claudicação intermitente. Interpretar estenose carotídea: <50% (leve), 50-69% (moderada - AAS + estatina + controle fatores risco), 70-99% (severa - considerar endarterectomia se sintomático). Se CIMT elevado ou estenose: abordagem funcional agressiva: 1) Otimização lipídica avançada: estatina alta potência (rosuvastatina 20-40mg ou atorvastatina 40-80mg) + ezetimiba 10mg (meta: LDL <70 mg/dL, idealmente <55 se alto risco); ômega-3 EPA 4g/dia (Vascepa - reduz eventos CV); bergamota (Vasguard) 1.000mg/dia (reduz LDL oxidado); niacina 500-2.000mg se HDL baixo. 2) Anti-inflamatórios: curcumina 1g/dia, resveratrol 500mg, ômega-3, dieta anti-inflamatória (eliminar açúcares, ultraprocessados). 3) Suporte endotelial: L-arginina 3-6g/dia (precursor NO), citrulina 3g, CoQ10 200mg, vitamina K2 (remove cálcio vascular). 4) Reduzir homocisteína: metilfolato 1-5mg, metilcobalamina 1.000mcg, P-5-P 50mg (meta homocisteína <7 µmol/L). 5) Exercício aeróbico 150min/sem, parar tabagismo, controlar hipertensão/diabetes. Repetir Doppler anualmente se estenose moderada."""
        }

    elif "TC" in name or "Tomografia" in name or "Angiotomografia" in name:
        return {
            "clinical": f"""A {name} fornece imagens transversais de alta resolução, superior à radiografia convencional. Escore de cálcio coronariano (CAC score) quantifica aterosclerose coronariana: 0 (sem placa), 1-100 (placa mínima), 101-400 (placa moderada), >400 (placa extensa - risco CV alto). CAC score >100 em <60 anos ou >400 qualquer idade indica risco significativo. Angiotomografia coronariana visualiza placas não calcificadas (vulneráveis). Na medicina funcional, CAC score >0 desencadeia intervenção intensiva preventiva: controle agressivo de fatores de risco, anti-inflamatórios, antioxidantes, suporte endotelial, reverter resistência insulínica.""",

            "patient": f"""A {name} usa raios-X em múltiplos ângulos para criar imagens detalhadas de órgãos internos. O escore de cálcio coronariano mede depósitos de cálcio nas artérias do coração - quanto maior o escore, maior o risco de infarto. Angiotomografia mostra as artérias coronárias em detalhe, detectando entupimentos. Envolve radiação e contraste iodado (risco para rins e alergias). Reservado para situações específicas, não para rastreamento rotineiro.""",

            "conduct": f"""Indicações {name}: dor torácica atípica (risco intermediário), triagem selecionada (risco intermediário Framingham 10-20%, histórico familiar forte, CAC score para estratificação). Se CAC score >0: tratamento agressivo mesmo se assintomático: 1) Estatina alta potência (rosuvastatina 20-40mg - reduz progressão placa em 5-10%), meta LDL <70 mg/dL. 2) AAS 81-100mg se CAC >100. 3) Controle rigoroso PA (<130/80), glicemia (HbA1c <5,7%). 4) Suporte funcional: vitamina K2 (MK-7) 360mcg/dia (inibe calcificação vascular, ativa MGP), magnésio 600mg (antagonista cálcio), vitamina D (meta 50-80 ng/ml), ômega-3 EPA 2-4g, curcumina 1g, CoQ10 200mg (se estatina - reposição obrigatória), bergamota 1.000mg, nattokinase 2.000 FU. 5) Reverter resistência insulínica: dieta low-carb, jejum intermitente, berberina 500mg 3x/dia ou metformina 500-1.000mg 2x/dia. 6) Exercício 150min/sem, cessar tabagismo. Repetir CAC em 5-7 anos para monitorar progressão (ideal <15% ao ano)."""
        }

    elif "USG" in name or "Ultrassom" in name or "Ultrassonografia" in name:
        return {
            "clinical": f"""O {name} utiliza ondas sonoras de alta frequência para visualizar estruturas internas em tempo real, sem radiação ionizante. Avalia morfologia, dimensões, textura, vascularização de órgãos. Na medicina funcional integrativa, ultrassom abdominal rastreia esteatose hepática não alcoólica (DHGNA - presente em 25-30% da população, associada a resistência insulínica), litíase biliar, alterações pancreáticas, esplenomegalia. USG transvaginal avalia ovários (SOP, cistos), útero (miomas, adenomiose, espessura endometrial). USG mamas complementa mamografia (nódulos, cistos). USG próstata triagem câncer (menos sensível que PSA/toque retal).""",

            "patient": f"""O {name} é exame de imagem que usa ondas sonoras (não usa radiação) para ver órgãos internos. É seguro, indolor e pode ser feito em grávidas. Avalia fígado, vesícula, pâncreas, rins, útero, ovários, próstata, mamas, tireoide e outras estruturas. Detecta nódulos, cistos, cálculos (pedras), inflamações, tumores, acúmulo de gordura no fígado. Exame básico de triagem muito útil.""",

            "conduct": f"""Indicações {name}: dor abdominal, alterações em exames lab (enzimas hepáticas, bilirrubina), triagem câncer (útero/ovários em mulheres, próstata em homens), avaliação de nódulos/massas. Se esteatose hepática (fígado gorduroso): não é inócua - DHGNA pode progredir para esteato-hepatite (NASH), fibrose, cirrose. Abordagem funcional: 1) Reverter resistência insulínica: dieta low-carb (<100g/dia), jejum intermitente, perder 7-10% peso (melhora DHGNA em 90%), exercício 150min/sem. 2) Suplementação hepatoprotetora: silimarina 420mg/dia (milk thistle - regeneração hepatocelular), NAC 1.200mg/dia (glutationa), alfa-lipóico 600mg, colina 500mg, inositol 2g, vitamina E 800 UI/dia (se NASH confirmado por biópsia). 3) Berberina 500mg 3x/dia ou metformina (melhoram DHGNA). 4) Evitar frutose (xarope de milho), álcool, triglicerídeos médios (TCM em excesso). Repetir USG em 6-12 meses para monitorar melhora. Se SOP (ovários policísticos): resistência insulínica é causa-raiz - mesma abordagem + inositol 4g/dia, NAC, vitamina D. Se litíase biliar sintomática: colecistectomia; assintomática: observation + prevenir crescimento (evitar jejum prolongado, perder peso gradualmente, fibras, lecitina)."""
        }

    elif "Fundoscopia" in name:
        return {
            "clinical": f"""A fundoscopia ({name}) examina retina, nervo óptico, mácula e vasos retinianos. É janela única para visualizar microvasculatura in vivo, refletindo saúde vascular sistêmica. Detecta: retinopatia diabética (microaneurismas, exsudatos, hemorragias - principal causa de cegueira em adultos), retinopatia hipertensiva (estreitamento arteriolar, cruzamentos AV), edema de papila (hipertensão intracraniana), atrofia óptica, degeneração macular. Na medicina funcional, achados vasculares retinianos indicam dano microvascular cerebral, renal e cardíaco. Prevenção via controle glicêmico rigoroso, PA, lipídios e suporte antioxidante.""",

            "patient": f"""A fundoscopia ({name}) examina o fundo do olho (retina, nervos e vasos sanguíneos) através da pupila dilatada. Detecta danos causados por diabetes (pode levar à cegueira), pressão alta, glaucoma, degeneração macular e outras doenças oculares. É exame importante porque muitas doenças dos olhos não causam sintomas iniciais - quando a visão começa a ficar embaçada, o dano já é avançado. Prevenir é mais fácil que reverter.""",

            "conduct": f"""Indicações fundoscopia ({name}): diabetes (anual desde diagnóstico), hipertensão mal controlada, cefaleia + alterações visuais (excluir papiledema), rastreamento glaucoma >40 anos. Se retinopatia diabética: classificar severidade (não proliferativa leve/moderada/severa vs proliferativa). Abordagem funcional preventiva intensiva: 1) Controle glicêmico rigoroso: HbA1c <6,5% (idealmente <6,0%) - cada 1% redução = 25-30% menos retinopatia. Estratégias: dieta low-carb, berberina 500mg 3x/dia, metformina, atividade física, sono adequado. 2) Controle PA: <130/80 mmHg. 3) Perfil lipídico: fenofibrato 160mg/dia (se TG >200 - reduz progressão retinopatia em 40%). 4) Antioxidantes oculares: luteína 10mg + zeaxantina 2mg (protegem mácula), vitamina C 500mg, vitamina E 400 UI, zinco 40mg, cobre 2mg (fórmula AREDS2). 5) Ômega-3 DHA 1.000mg (DHA concentra-se na retina). 6) Alfa-lipóico 600mg (neuroprotetor, antioxidante). 7) Parar tabagismo (risco retinopatia dobra). Se retinopatia proliferativa: encaminhar oftalmologista urgente (fotocoagulação laser, anti-VEGF intravítreo). Fundoscopia anual obrigatória em diabéticos."""
        }

    else:  # Fallback genérico para outros exames de imagem
        return {
            "clinical": f"""O {name} é método diagnóstico por imagem essencial na medicina moderna. Na abordagem funcional integrativa, exames de imagem complementam avaliação clínica e laboratorial, permitindo visualizar estrutura e função de órgãos, detectar alterações anatômicas, inflamação, tumores, calcificações, obstruções. Resultados devem ser interpretados em contexto clínico completo, correlacionando com sintomas, história, biomarcadores inflamatórios/metabólicos. Achados orientam investigação de causas-raiz e monitoramento de intervenções terapêuticas.""",

            "patient": f"""O {name} é exame que cria imagens de órgãos e estruturas internas do corpo, ajudando médicos a detectar problemas de saúde, diagnosticar doenças e acompanhar tratamentos. Cada tipo de exame tem indicações específicas - seu médico solicitará quando houver suspeita clínica ou necessidade de rastreamento preventivo. Siga orientações de preparo (jejum, hidratação, medicações) para garantir qualidade do exame.""",

            "conduct": f"""Solicitar {name} conforme indicação clínica específica, considerando: sintomas apresentados, alterações em exames prévios, necessidade de rastreamento (idade, fatores de risco), monitoramento de condições conhecidas. Interpretar resultados no contexto funcional integrativo: se alterações detectadas, investigar causas-raiz através de: marcadores inflamatórios (PCR-us, VHS), perfil metabólico (glicemia, insulina, HbA1c, perfil lipídico avançado), hormônios (tireoide, cortisol, sexuais), micronutrientes (vitamina D, magnésio, zinco, B12), marcadores de estresse oxidativo. Intervir com abordagem multimodal: otimização nutricional, suplementação direcionada, modulação de estilo de vida, manejo de estresse, correção de deficiências. Repetir exame conforme evolução clínica e resposta terapêutica."""
        }

    return {
        "clinical": content["clinical"],
        "patient": content["patient"],
        "conduct": content["conduct"]
    }

# Lista dos 100 items (primeiros 100 do resultado anterior)
ITEMS = [
    ("5ef6017d-ef5d-41a3-bca2-109cc75fb394", "40101010 - Eletrocardiograma (ECG 12 derivações)"),
    ("5f62a5d0-f0fe-4b40-8ddc-a1265fd5f3d0", "40101010 - Eletrocardiograma (ECG 12 derivações)"),
    ("4773bed8-5a9c-4875-9e78-23b29ed86405", "40101010 - Eletrocardiograma (ECG 12 derivações)"),
    ("a201fc10-b1a8-49a6-8281-9126f5835d38", "40201082 - Colonoscopia"),
    ("3126328e-6bdd-44a9-ac37-2e57e5b73169", "40201082 - Colonoscopia"),
    ("52991341-2fde-486d-9240-16bdd6263086", "40201082 - Colonoscopia"),
    ("d865f95e-329a-442f-b509-1b654d6b828e", "40201120 - Endoscopia digestiva alta"),
    ("480b8a24-d4c3-43af-a2a1-1fb6e0634025", "40201120 - Endoscopia digestiva alta"),
    ("d922be27-270f-44d2-8b66-8feb9fdabf97", "40201120 - Endoscopia digestiva alta"),
    ("89a6df89-d343-4255-9e45-10bf0cb55d97", "40808041 - Mamografia digital bilateral"),
    ("a53b8c72-c148-43c6-86f7-34bbefdc9544", "40808041 - Mamografia digital bilateral"),
    ("03123378-ca32-4088-8536-31666f049f75", "40808041 - Mamografia digital bilateral"),
    ("ca7da1cd-2b71-4352-92ba-3f5690078d51", "40808149 - Densitometria corpo inteiro (composição corporal)"),
    ("e9ad56fa-b958-4583-874b-f55c58f30b97", "40808149 - Densitometria corpo inteiro (composição corporal)"),
    ("c725c0e8-ff20-48d4-a7b0-1d0475a92cb2", "40808149 - Densitometria corpo inteiro (composição corporal)"),
    ("26bfb650-ae71-47cb-8a5d-f672e2777dce", "40901106 - Ecodopplercardiograma transtorácico"),
    ("ec83822f-0999-49b1-b0e5-7d23e64616c7", "40901106 - Ecodopplercardiograma transtorácico"),
    ("49171389-d0c7-4a92-a8a7-e4f3d48246c3", "40901106 - Ecodopplercardiograma transtorácico"),
    ("9f306531-7f98-4b87-80fc-8fbdc46d952d", "40901114 - USG mamas"),
    ("ca3db864-3407-41a2-b317-07fb2e53253e", "40901114 - USG mamas"),
    ("92891a23-02aa-4d4a-a3c2-f770b6505759", "40901114 - USG mamas"),
    ("e44680b0-138a-4a8f-812b-eedca2838c94", "40901122 - USG abdome total"),
    ("a87ee825-298c-49eb-aea2-183981f49a89", "40901122 - USG abdome total"),
    ("1b1e689d-14af-4b66-a269-be27ddcce26d", "40901122 - USG abdome total"),
    ("b29bbeb0-b0bd-4f71-acaa-fab10b2ac45d", "40901300 - USG transvaginal (útero, ovário, anexos e vagina)"),
    ("5a315bc4-8f86-412f-be04-315636a9e3ec", "40901300 - USG transvaginal (útero, ovário, anexos e vagina)"),
    ("1178ea33-d57f-419d-a9b2-a90474bb30d1", "40901300 - USG transvaginal (útero, ovário, anexos e vagina)"),
    ("bf1e1008-8a30-41f8-b39a-067ce5ef782a", "40901360 - Doppler colorido de vasos cervicais arteriais bilateral (carótidas e vertebrais)"),
    ("f22b2474-b22e-4fc9-bf44-4706378bf88a", "40901360 - Doppler colorido de vasos cervicais arteriais bilateral (carótidas e vertebrais)"),
    ("e7b2ab85-9ac7-4edf-b888-cc373bcf7b4e", "40901360 - Doppler colorido de vasos cervicais arteriais bilateral (carótidas e vertebrais)"),
    ("7fe1681d-b0e7-485f-ae06-7391d98c6df6", "40901394 - Doppler colorido de aorta e artérias renais"),
    ("ee825218-ca7a-47c6-b0d9-78e086bbbf08", "40901394 - Doppler colorido de aorta e artérias renais"),
    ("01f1d5e4-4b02-45e5-8a07-3830a94474ce", "40901394 - Doppler colorido de aorta e artérias renais"),
    ("8815b8b2-8a3f-4d04-bdb6-4b9f2b8a488a", "40901750 - USG próstata (via abdominal)"),
    ("d149c9ab-f69b-4ca2-afa1-62d8b096c055", "40901750 - USG próstata (via abdominal)"),
    ("4bbaf683-98af-4947-a99e-5a99e96ced70", "40901750 - USG próstata (via abdominal)"),
    ("c0ee3c8a-1a83-4853-867a-e6f66893af90", "41001079 - TC tórax"),
    ("d028fed6-20d0-47c5-a220-3ccab547f5cb", "41001079 - TC tórax"),
    ("a7368290-1c34-4376-b7ea-b49351283189", "41001079 - TC tórax"),
    ("fd6ae113-0e05-4ade-90a6-1d9a1bbb34e5", "41001087 - TC coração para escore de cálcio coronariano"),
    ("acae175f-0a75-4a2b-958f-dbddc4e3af14", "41001087 - TC coração para escore de cálcio coronariano"),
    ("8e5d5e56-9a2f-46a7-b671-5d0f54885bf5", "41001087 - TC coração para escore de cálcio coronariano"),
    ("2130a92b-fb3b-4fa5-a042-66a9176520ab", "41001230 - Angiotomografia coronariana"),
    ("3e097ac7-75d3-4b01-951c-8c30783d00bc", "41001230 - Angiotomografia coronariana"),
    ("70942d7c-f82f-4080-8d79-00f228455a75", "41001230 - Angiotomografia coronariana"),
    ("f7c0f0bb-fb7c-4af0-8935-1d484eb159aa", "41301439 - Fundoscopia sob midríase (binocular)"),
    ("fbd27704-676b-40cf-8d0a-063dcef8f7ed", "41301439 - Fundoscopia sob midríase (binocular)"),
    ("f9725605-3a58-4753-8502-4c66be7b566b", "41301439 - Fundoscopia sob midríase (binocular)"),
    ("d2e30831-413b-4b2e-af81-2ad905858c32", "81000405 - Radiografia panorâmica de mandíbula e maxila"),
    ("fa4875a2-11a7-42e4-b41e-847373016e1e", "81000405 - Radiografia panorâmica de mandíbula e maxila"),
    ("6fc549a3-4413-4741-a55b-0ac021259105", "81000405 - Radiografia panorâmica de mandíbula e maxila"),
    ("9064aa05-9f21-402f-9c36-49887d2b8ec9", "CIMT Carótidas (Espessura Íntima-Média)"),
    ("562a9ae4-52d4-465a-aa78-b73cd14f1ca1", "CIMT Carótidas (Espessura Íntima-Média)"),
    ("8f5f3d75-c847-4e98-a74a-87966e600252", "CIMT Carótidas (Espessura Íntima-Média)"),
    ("59585a72-8bc9-4ab7-a81f-2cf21706bcee", "Colonoscopia - Mayo Score UC"),
    ("4e601b8f-e3e0-484e-9d16-d59e7f1682b2", "Colonoscopia - Mayo Score UC"),
    ("c5d52258-9298-44f9-978a-92717e6dba4b", "Colonoscopia - Mayo Score UC"),
    ("3e0353cd-a928-4a5e-b96f-c7cea26262d2", "Colonoscopia - Número Total Adenomas"),
    ("0296e16f-52e2-411a-b4bb-9948ea1eb598", "Colonoscopia - Número Total Adenomas"),
    ("7a182c48-8b53-49e7-baeb-71f282325420", "Colonoscopia - Número Total Adenomas"),
    ("14dec504-0061-446b-b6f1-4ac464b0e463", "Colonoscopia - SES-CD Crohn"),
    ("7c8d2198-a150-4604-8ba9-32478ce36697", "Colonoscopia - SES-CD Crohn"),
    ("25b3da57-83df-4673-b0c8-914fa21e6f1e", "Colonoscopia - SES-CD Crohn"),
    ("29659962-4210-4f38-a9be-fb2278bc1bdc", "Delta Hidrogênio (Δ H₂)"),
    ("430384f8-52d9-4775-b46b-730db6567458", "Delta Hidrogênio (Δ H₂)"),
    ("3f6fb4d7-efa4-4dac-93ef-95da8fbee038", "Delta Hidrogênio (Δ H₂)"),
    ("e164ee28-7a9d-4750-8b65-9ce7799cd84d", "Delta Metano (Δ CH₄)"),
    ("bda631d4-0a5d-42d0-bc41-2d9516ec2bef", "Delta Metano (Δ CH₄)"),
    ("b1b0b19f-12ef-40bd-9d11-16cecb78863c", "Delta Metano (Δ CH₄)"),
    ("4ae093c8-b6ac-4ac1-b242-dcd64194e6ad", "Densitometria - T-Score Colo Femoral"),
    ("4e272227-9881-4f3b-bd87-e7cebbb7ef63", "Densitometria - T-Score Colo Femoral"),
    ("0f9291af-8a3c-4aa6-b51e-8ddbee42dab7", "Densitometria - T-Score Colo Femoral"),
    ("6cd4edf1-6392-493b-8086-dc0a89d0afeb", "Densitometria - T-Score Coluna Lombar"),
    ("e46dce84-5755-48c2-978b-e5b12aa7bdd4", "Densitometria - T-Score Coluna Lombar"),
    ("73f6f0f6-1794-484a-8fdd-1245aa35701e", "Densitometria - T-Score Coluna Lombar"),
    ("9f01023a-a7ec-449d-b18d-ca46245038cb", "Doppler Aorta - PSV Artéria Renal"),
    ("0c052bec-e860-4068-927c-db125905c0d3", "Doppler Aorta - PSV Artéria Renal"),
    ("2af91e72-4889-47a0-a75c-19445289e827", "Doppler Aorta - PSV Artéria Renal"),
    ("b7c85ad4-ece5-4a21-b23d-8edf321f7874", "Doppler Aorta - RAR (Renal-Aortic Ratio)"),
    ("2e33bc6a-52e7-4d5e-9de7-5f6f92b5236f", "Doppler Aorta - RAR (Renal-Aortic Ratio)"),
    ("dc0ee2c3-6d48-4094-a423-e942f4d3c07b", "Doppler Aorta - RAR (Renal-Aortic Ratio)"),
    ("3e7e0ea2-6ad1-403e-acc1-2ca4d575e4d4", "Doppler Carótidas - Estenose Carotídea"),
    ("cff78506-db7e-4682-915c-6ae7db0fb0b8", "Doppler Carótidas - Estenose Carotídea"),
    ("579a961c-e160-417f-9371-418284386f35", "Doppler Carótidas - Estenose Carotídea"),
    ("932f426c-7a98-437d-8de5-e0f1c1b29e7a", "Doppler Carótidas - PSV Carótida Interna"),
    ("b3d1bfe4-35c6-4509-a41a-ddb014beae43", "Doppler Carótidas - PSV Carótida Interna"),
    ("f4bd7c59-f1d9-41f0-81c2-65d5c7bcd098", "Doppler Carótidas - PSV Carótida Interna"),
    ("6f149dac-b336-4383-acde-fdb102a8d68c", "ECG - Cornell Voltage - Homens (R aVL + S V3)"),
    ("10655c96-4de2-4c05-9ab4-7ab5fcf1cefb", "ECG - Cornell Voltage - Homens (R aVL + S V3)"),
    ("77c90eb9-d99c-4347-9c48-51e0d9ead5d6", "ECG - Cornell Voltage - Homens (R aVL + S V3)"),
    ("eccba072-18aa-4cea-99a6-70fd183624e8", "ECG - Cornell Voltage - Mulheres (R aVL + S V3)"),
    ("3be39560-d2f1-424c-9f1d-bfbd9472558a", "ECG - Cornell Voltage - Mulheres (R aVL + S V3)"),
    ("1a3debb7-f986-4f0c-86f0-fa05fcc1891e", "ECG - Cornell Voltage - Mulheres (R aVL + S V3)"),
    ("7646f122-6c7b-46a3-8f75-d96e42bd1a3f", "ECG - Duração QRS"),
    ("7c35d0f9-f48b-4be1-808e-c32a975ccf84", "ECG - Duração QRS"),
    ("ac440fb9-7452-40c3-b896-a0d50a9e3054", "ECG - Duração QRS"),
    ("c14642d1-a657-4d62-8881-b8f8ea6def2f", "ECG - Eixo Cardíaco"),
    ("2e9e9628-b7d3-49b6-ae60-61aca75de8f8", "ECG - Eixo Cardíaco"),
    ("91dec098-5918-4bf9-9cbb-2967847d2284", "ECG - Eixo Cardíaco"),
    ("79a44be5-fa22-4ecf-8900-9cb51f007292", "ECG - Frequência Cardíaca"),
]

def main():
    print("=" * 80)
    print("BATCH: 100 EXAMES DE IMAGEM - MEDICINA FUNCIONAL INTEGRATIVA")
    print("=" * 80)

    print("\n[1/3] Autenticando...")
    token = login()
    print("✓ Login realizado")

    print(f"\n[2/3] Processando {len(ITEMS)} items...")
    success = 0
    errors = []

    for idx, (item_id, name) in enumerate(ITEMS, 1):
        try:
            content = get_content(name)
            if update_item(token, item_id, content["clinical"], content["patient"], content["conduct"]):
                success += 1
                if idx % 10 == 0:
                    print(f"  ✓ {idx}/{len(ITEMS)} processados...")
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
        for idx, err in errors[:5]:
            print(f"  - Item {idx}: {err}")
    print("=" * 80)

if __name__ == "__main__":
    main()
