#!/usr/bin/env python3
"""Verify character counts for LDL oxidada enrichment content."""

clinical_relevance = '''A lipoproteína de baixa densidade oxidada (LDL oxidada ou oxLDL) representa um marcador fundamental na avaliação do risco cardiovascular aterosclerótico. A oxidação da LDL nativa ocorre quando radicais livres modificam seus componentes lipídicos e proteicos, transformando-a em uma partícula altamente aterogênica. Estudos recentes demonstram que níveis elevados de oxLDL estão diretamente associados à presença, progressão e gravidade da doença cardiovascular aterosclerótica. Meta-análises de 2023 confirmam que a oxLDL participa de um ciclo vicioso entre aterosclerose e inflamação, sendo mais causalmente relacionada aos desfechos cardiovasculares do que o LDL-C convencional. Pesquisas de 2024 evidenciam que valores de oxLDL ≥7358,82 µg/mL associam-se a risco 4,6 vezes maior de múltiplos eventos isquêmicos e acometimento de múltiplos vasos coronários. A oxLDL induz disfunção endotelial precoce, ativa cascatas inflamatórias, promove formação de células espumosas e estimula proliferação de células musculares lisas na parede arterial. Seu potencial antigênico desencadeia respostas imunes inatas e adaptativas que perpetuam o processo aterosclerótico. A dosagem de oxLDL mostra-se especialmente útil na identificação de aterosclerose subclínica, estratificação de risco em populações aparentemente saudáveis e em condições de inflamação crônica como diabetes mellitus tipo 2. Valores elevados correlacionam-se com instabilidade de placa aterosclerótica, risco de eventos coronários agudos e extensão da doença arterial coronariana. Embora ainda não padronizada para uso clínico rotineiro, a oxLDL representa um biomarcador promissor para refinamento da avaliação de risco cardiovascular e personalização de estratégias preventivas e terapêuticas.'''

patient_explanation = '''A LDL oxidada é uma forma "danificada" do colesterol ruim (LDL) que se torna muito mais perigosa para suas artérias. Imagine o colesterol LDL como um caminhão transportando gordura pelo sangue. Quando esse caminhão sofre oxidação (como se "enferrujasse"), ele se torna extremamente prejudicial, grudando nas paredes das artérias e iniciando um processo inflamatório que forma placas de gordura. Quanto maior o nível de LDL oxidada no seu sangue, maior o risco de entupimento das artérias do coração (aterosclerose). Pesquisas recentes mostram que pessoas com níveis muito elevados de LDL oxidada têm mais de 4 vezes maior chance de sofrer infartos ou ter várias artérias comprometidas. A LDL oxidada não apenas deposita gordura nas artérias, mas também causa inflamação constante, danifica a camada interna dos vasos sanguíneos e facilita a formação de coágulos. É como ter um processo de enferrujamento contínuo dentro das artérias. Esse exame é especialmente importante se você tem diabetes, pressão alta, histórico familiar de doença cardíaca ou outros fatores de risco cardiovascular. Níveis elevados indicam que medidas preventivas mais intensivas podem ser necessárias, como ajustes na alimentação, exercícios regulares, uso de antioxidantes e, em alguns casos, medicamentos específicos. O objetivo é reduzir tanto a quantidade de LDL quanto protegê-la da oxidação.'''

conduct = '''VALORES DE REFERÊNCIA: Ainda não há consenso internacional absoluto, mas estudos recentes sugerem que valores <7358,82 µg/mL apresentam menor risco cardiovascular, enquanto valores ≥7358,82 µg/mL associam-se significativamente a maior risco de eventos isquêmicos múltiplos e acometimento vascular extenso. A interpretação deve considerar o contexto clínico individual e outros marcadores de risco.

VALORES ELEVADOS (≥7358 µg/mL ou acima do percentil 75 para laboratório de referência):
1. Investigação complementar: Perfil lipídico completo com LDL-C, HDL-C, triglicerídeos, apolipoproteínas A1 e B, lipoproteína(a), PCR ultrassensível, hemoglobina glicada
2. Avaliação de aterosclerose subclínica: Escore de cálcio coronário, espessura médio-intimal de carótidas, índice tornozelo-braquial conforme indicação clínica
3. Estratificação de risco cardiovascular global utilizando calculadoras validadas (Framingham, ASCVD Risk Estimator)
4. Modificações intensivas do estilo de vida: Dieta mediterrânea ou DASH com ênfase em antioxidantes naturais (frutas vermelhas, vegetais folhosos, azeite extra-virgem, oleaginosas, peixes ricos em ômega-3), exercícios aeróbicos 150-300min/semana, cessação tabagismo, controle do estresse
5. Suplementação antioxidante individualizada: Avaliar vitamina E (até 400 UI/dia), vitamina C (500-1000mg/dia), coenzima Q10 (100-200mg/dia) conforme perfil do paciente
6. Terapia hipolipemiante otimizada: Estatinas de alta intensidade visando LDL-C <70mg/dL (<55mg/dL se alto risco ou <40mg/dL se muito alto risco), considerar associação com ezetimiba ou inibidores de PCSK9
7. Controle rigoroso de comorbidades: HbA1c <7% em diabéticos, PA <130/80mmHg, obesidade (meta IMC <25 kg/m²)
8. Reavaliação em 3-6 meses após intervenções intensivas

VALORES NORMAIS/CONTROLADOS:
1. Manutenção de estilo de vida cardioprotetor
2. Monitoramento anual da oxLDL em pacientes de risco moderado-alto
3. Perfil lipídico anual
4. Reforço de medidas preventivas primárias

SEMPRE: Interpretação integrada com quadro clínico, fatores de risco cardiovascular, histórico familiar e demais exames complementares. Considerar encaminhamento para cardiologista ou especialista em dislipidemia em casos de risco elevado ou refratariedade terapêutica.'''

print("=== CHARACTER COUNT VERIFICATION ===")
print(f"clinical_relevance: {len(clinical_relevance)} chars (target: 1500-2000)")
print(f"patient_explanation: {len(patient_explanation)} chars (target: 1000-1500)")
print(f"conduct: {len(conduct)} chars (target: 1500-2500)")
print()

# Check ranges
in_range = True
if not (1500 <= len(clinical_relevance) <= 2000):
    print(f"WARNING: clinical_relevance outside target range!")
    in_range = False

if not (1000 <= len(patient_explanation) <= 1500):
    print(f"WARNING: patient_explanation outside target range!")
    in_range = False

if not (1500 <= len(conduct) <= 2500):
    print(f"WARNING: conduct outside target range!")
    in_range = False

if in_range:
    print("SUCCESS: All character counts within target ranges!")
