-- BATCH FINAL 3 - EXAMES C (Primeiros 5 items)
-- Gerado em: 2026-01-28
-- Padrão MFI com valores ótimos e condutas específicas

BEGIN;


UPDATE score_items
SET
  optimal_range_min = NULL,
  optimal_range_max = NULL,
  unit = 'grau (0-4)',
  optimal_explanation = 'Ausência de retinopatia (grau 0) indica controle glicêmico adequado, prevenindo microangiopatia retiniana e reduzindo risco de cegueira em 90%',
  clinical_interpretation = '## Fundamento Fisiológico

A retinopatia diabética resulta da hiperglicemia crônica causando lesão microvascular progressiva na retina. O excesso de glicose ativa vias de dano celular (poliol, AGEs, estresse oxidativo) que comprometem pericitos, células endoteliais e causam neovascularização patológica.

## Classificação e Interpretação

**Grau 0 (Sem retinopatia):** Controle glicêmico adequado, HbA1c <7%, ausência de lesões microvasculares. Meta terapêutica em medicina funcional.

**Grau 1-2 (Não proliferativa leve/moderada):** Microaneurismas, exsudatos, hemorragias pontuais. Indica descompensação glicêmica crônica, necessita intensificação de controle metabólico.

**Grau 3-4 (Não proliferativa severa/proliferativa):** Neovascularização, descolamento de retina. Requer intervenção oftalmológica urgente (fotocoagulação laser).

## Contexto Clínico Funcional

Correlacionar com HbA1c, frutosamina, glicemia pós-prandial, marcadores inflamatórios (PCR-us, IL-6) e estresse oxidativo (8-OHdG). A ausência de retinopatia com HbA1c >8% sugere proteção vascular (investigar polimorfismos VEGF, SOD2). Rastreamento anual obrigatório em diabéticos, semestral em graus 2+.',
  medical_conduct = '## Prevenção e Controle (Grau 0-1)

- Manter HbA1c <6,5% através de dieta low-carb e exercício
- Suplementação antioxidante: Ácido alfa-lipóico 600mg/dia, Luteína/Zeaxantina 10mg/dia
- Controle pressórico rigoroso: PA <130/80 mmHg
- Fundoscopia anual para monitoramento

## Retinopatia Estabelecida (Grau 2-4)

- Encaminhamento imediato ao retinólogo
- Fotocoagulação laser panretiniana se proliferativa
- Anti-VEGF intravítreo (ranibizumabe) em edema macular
- Intensificar controle glicêmico gradualmente (evitar hipoglicemias)

## Abordagem Funcional Integrativa

- N-acetilcisteína 1.200mg/dia para reduzir estresse oxidativo
- Ômega-3 EPA/DHA 2-3g/dia (reduz neovascularização)
- Coenzima Q10 200mg/dia (proteção mitocondrial)
- Vitamina D3 manter >40 ng/mL
- Dieta anti-inflamatória: eliminar açúcares refinados, priorizar vegetais coloridos
- Monitorar pressão arterial domiciliar diariamente',
  related_articles = ARRAY['Wong TY, et al. Diabetic retinopathy. Nat Rev Dis Primers. 2016;2:16012. PMID: 27159554', 'Simó R, et al. Neurodegeneration in diabetic retinopathy: does it really matter? Diabetologia. 2018;61(9):1902-1912. PMID: 29392494', 'Stern JH, et al. Emerging therapies for diabetic retinopathy. Curr Opin Ophthalmol. 2023;34(3):214-220. PMID: 36938849'],
  updated_at = NOW()
WHERE id = 'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52'::uuid;


UPDATE score_items
SET
  optimal_range_min = NULL,
  optimal_range_max = 4.0,
  unit = 'mm',
  optimal_explanation = 'Endométrio ≤4mm na pós-menopausa exclui hiperplasia/câncer com VPN 99%, dispensando biópsia. Espessura >4mm requer investigação histológica',
  clinical_interpretation = '## Fundamento Fisiológico

Após a menopausa, a queda estrogênica leva à atrofia endometrial fisiológica (eco <4mm). Espessamento acima deste valor sugere estimulação estrogênica anormal (endógena ou exógena), hiperplasia ou malignidade. A ultrassonografia transvaginal é método inicial de rastreamento pela excelente correlação com achados histológicos.

## Interpretação de Valores

**≤4mm (Atrofia fisiológica):** Padrão esperado na pós-menopausa sem terapia hormonal. Valor preditivo negativo de 99% para câncer endometrial. Não requer biópsia na ausência de sangramento.

**5-10mm (Zona cinzenta):** Considerar fatores de risco (obesidade, SOP prévia, tamoxifeno, HRT). Em mulheres com sangramento, biópsia endometrial indicada. Sem sangramento, repetir USG em 3-6 meses.

**>10mm (Espessamento patológico):** Alta suspeita de hiperplasia com/sem atipia ou adenocarcinoma. Biópsia endometrial ou histeroscopia com biópsia dirigida obrigatória, independente de sintomas.

## Contexto Clínico

Em usuárias de terapia hormonal (estrogênio+progesterona), espessura até 8mm pode ser aceitável se assintomática. Tamoxifeno causa efeitos estrogênicos uterinos (limite 8mm). Obesidade e diabetes aumentam risco de hiperplasia/câncer (hiperestrogenismo por aromatização periférica).',
  medical_conduct = '## Endométrio ≤4mm (Sem Sangramento)

- Reavaliação anual com USG transvaginal
- Manter peso saudável (IMC <25 kg/m²)
- Suplementação: Vitamina D3 5.000 UI/dia, Ômega-3 2g/dia
- Evitar fontes exógenas de estrogênio (fitoestrogênios em excesso)

## Endométrio 5-10mm ou >4mm com Sangramento

- Biópsia endometrial ambulatorial (Pipelle) ou histeroscopia
- Suspender TRH temporariamente antes da biópsia (lavar out hormonal)
- Investigar causas de hiperestrogenismo: obesidade, tumores ovarianos produtores
- Perfil hormonal: estradiol, FSH, LH

## Endométrio >10mm ou Hiperplasia Confirmada

- Hiperplasia simples sem atipia: Progesterona micronizada 200mg/dia VO dias 1-14/mês por 6 meses, repetir biópsia
- Hiperplasia com atipia: Encaminhamento cirúrgico (histerectomia)
- Modulação estrogênica: DIM (diindolilmetano) 300mg/dia, Cálcio-D-glucarato 500mg 2x/dia
- Perda de peso se IMC >25: reduz aromatização periférica

## Monitoramento

- USG transvaginal 3-6 meses após intervenção
- Atenção a sangramento pós-menopausa (sempre investigar)',
  related_articles = ARRAY['Smith-Bindman R, et al. Endometrial thickness for diagnosing endometrial cancer in symptomatic postmenopausal women. Obstet Gynecol. 2020;135(2):e16-e36. PMID: 31977782', 'Clarke MA, et al. Association of endometrial cancer risk with postmenopausal bleeding. JAMA Intern Med. 2018;178(9):1210-1222. PMID: 30083701', 'Giannella L, et al. Accuracy of ultrasound in predicting endometrial pathology. J Minim Invasive Gynecol. 2023;30(4):289-297. PMID: 36581012'],
  updated_at = NOW()
WHERE id = '30af9809-7316-47e6-b363-8279c7bd3a69'::uuid;


UPDATE score_items
SET
  optimal_range_min = 3.0,
  optimal_range_max = 10.0,
  unit = 'cm³',
  optimal_explanation = 'Volume ovariano 3-10 cm³ indica função folicular preservada. <3 cm³ sugere insuficiência ovariana; >10 cm³ pode indicar SOP ou tumor',
  clinical_interpretation = '## Fundamento Fisiológico

O volume ovariano reflete a reserva folicular e atividade endócrina. É calculado pela fórmula do elipsóide: 0,523 × comprimento × largura × altura. Varia com idade, fase do ciclo menstrual e status reprodutivo. Valores reduzidos correlacionam com reserva ovariana diminuída (RFO baixa), enquanto volumes aumentados sugerem síndrome dos ovários policísticos (SOP) ou neoplasias.

## Interpretação por Faixa Etária

**Menacme (20-40 anos):**
- Normal: 5-10 cm³ por ovário
- <3 cm³: Insuficiência ovariana prematura (IOP), baixa reserva ovariana
- >10 cm³: SOP (critério Rotterdam: volume >10 cm³ OU ≥12 folículos de 2-9mm)

**Perimenopausa (40-50 anos):**
- Normal: 3-8 cm³ (redução fisiológica da reserva)
- <2 cm³: Falência ovariana iminente/instalada

**Pós-menopausa (>50 anos):**
- Normal: 1,5-3 cm³ (atrofia fisiológica)
- >5 cm³: Investigar neoplasia (CA-125, doppler, RM)

## Contexto Clínico

Correlacionar volume com contagem de folículos antrais (CFA), hormônio antimülleriano (AMH), FSH e LH. SOP: volume >10 cm³ + CFA ≥12 + hiperandrogenismo clínico/laboratorial. IOP: volume <3 cm³ + AMH <0,5 ng/mL + FSH >25 UI/L.',
  medical_conduct = '## Volume Reduzido (<3 cm³) - Baixa Reserva Ovariana

- Dosagem de AMH, FSH basal (dia 3 do ciclo), estradiol
- DHEA 25-50mg/dia (melhora resposta ovariana)
- Coenzima Q10 ubiquinol 300-600mg/dia (qualidade oocitária)
- Mio-inositol 2g + Ácido fólico 400mcg/dia
- Vitamina D3 manter >40 ng/mL
- Se desejo de gravidez: encaminhamento rápido para reprodução assistida

## Volume Aumentado (>10 cm³) - Suspeita de SOP

- Perfil andrógenos: Testosterona total e livre, DHEA-S, Androstenediona
- Resistência insulínica: HOMA-IR, Teste de tolerância à glicose
- Metformina 1.500-2.000mg/dia se HOMA-IR >2,5
- Inositol (mio+d-quiro 40:1) 4g/dia - reduz volume ovariano
- Dieta low-carb, exercício resistido 3-4x/semana
- NAC 1.800mg/dia (melhora ovulação em 30%)
- Berberina 1.500mg/dia se intolerância à metformina

## Volume Aumentado Pós-Menopausa (>5 cm³)

- Doppler colorido para avaliar vascularização (IR <0,4 = suspeito)
- CA-125, HE4, índice ROMA
- RM pelve com contraste se achados suspeitos
- Encaminhamento oncológico se massa complexa/sólida

## Monitoramento

- Menacme: USG anual ou conforme sintomas
- SOP: USG 6-12 meses após intervenções
- Pós-menopausa: USG anual se volume normal, imediato se alteração',
  related_articles = ARRAY['Balen AH, et al. Ultrasound assessment of the polycystic ovary: international consensus definitions. Hum Reprod Update. 2003;9(6):505-514. PMID: 14714587', 'Souter I, et al. Ovarian volume and antral follicle count in women with polycystic ovary syndrome. Fertil Steril. 2021;115(6):1478-1485. PMID: 33832672', 'Amer SA, et al. Ultrasonographic evaluation of ovarian reserve in women of reproductive age. Ultrasound Obstet Gynecol. 2022;60(3):291-301. PMID: 35191089'],
  updated_at = NOW()
WHERE id = 'afdd9d67-3f42-4e6e-a77d-0e75475ca72d'::uuid;


UPDATE score_items
SET
  optimal_range_min = 350.0,
  optimal_range_max = 430.0,
  unit = 'ms',
  optimal_explanation = 'QTc 350-430ms em homens reduz risco de morte súbita. QTc >450ms aumenta risco de torsades de pointes; <350ms pode indicar QT curto congênito',
  clinical_interpretation = '## Fundamento Fisiológico

O intervalo QT corrigido (QTc) representa o tempo de repolarização ventricular ajustado pela frequência cardíaca. A fórmula de Bazett (QTc = QT/√RR) é amplamente utilizada mas superestima em taquicardia e subestima em bradicardia. Prolongamento do QTc indica heterogeneidade da repolarização, predispondo a arritmias ventriculares malignas (torsades de pointes, fibrilação ventricular).

## Interpretação por Faixa de Valores

**350-430ms (Normal em homens):** Repolarização ventricular homogênea, baixo risco arrítmico. Meta funcional para otimização cardiovascular.

**431-450ms (Limítrofe):** Zona cinzenta, investigar causas adquiridas (medicamentos, distúrbios eletrolíticos). Risco intermediário de arritmias.

**451-470ms (Prolongado leve):** QTc >450ms em homens define prolongamento patológico. Risco 2-3x maior de morte súbita. Investigar causas secundárias antes de assumir síndrome congênita.

**>470ms (Prolongado grave):** Alto risco de torsades de pointes (especialmente >500ms). Requer avaliação cardiológica urgente, suspensão de QT-prolongadores, correção de eletrólitos.

**<350ms (QT curto):** Raro, sugere síndrome do QT curto congênito (canalopatia). Também associado a hipercalcemia, digoxina. Risco de fibrilação atrial e morte súbita.

## Causas Adquiridas de QTc Prolongado

Medicamentos: antiarrítmicos (amiodarona, sotalol), antipsicóticos (haloperidol, quetiapina), antibióticos (azitromicina, fluoroquinolonas), antidepressivos (citalopram). Distúrbios eletrolíticos: hipocalemia (<3,5 mEq/L), hipomagnesemia (<1,5 mg/dL), hipocalcemia. Outros: hipotireoidismo, hipotermia, AVC agudo.',
  medical_conduct = '## QTc Normal (350-430ms)

- Manutenção: hidratação adequada, dieta rica em potássio (banana, abacate, batata doce)
- Magnésio quelado 400mg/dia para otimização eletrolítica
- Evitar início de medicamentos QT-prolongadores sem ECG prévio

## QTc Limítrofe/Prolongado (431-470ms)

- Dosagem urgente: Potássio, Magnésio, Cálcio iônico, TSH
- Revisar lista de medicamentos (suspender QT-prolongadores se possível)
- Reposição agressiva se hipocalemia/hipomagnesemia:
  - K <3,5: KCl 40-80 mEq VO/dia fracionado
  - Mg <1,7: Magnésio quelado 600-800mg/dia
- ECG de controle após 3-7 dias de correção eletrolítica
- Avaliação cardiológica se QTc >460ms persistente

## QTc >470ms (Alto Risco)

- Suspensão imediata de todos QT-prolongadores
- Internação se sintomas (síncope, pré-síncope, palpitações)
- Magnésio IV 2g em bolus se torsades de pointes
- Reposição IV de K+/Mg++ se deficiência grave
- Teste genético para canalopatias (genes KCNQ1, KCNH2, SCN5A) se QTc >500ms
- Considerar beta-bloqueador (nadolol 40-80mg/dia) em síndrome congênita

## QTc <350ms

- Dosagem de cálcio iônico, digoxinemia se em uso
- Teste genético para síndrome do QT curto (genes KCNH2, KCNQ1)
- Avaliação cardiológica para risco de FA/morte súbita

## Monitoramento

- QTc limítrofe: ECG mensal até normalização
- QTc prolongado grave: Holter 24h, ECG semanal inicial',
  related_articles = ARRAY['Rautaharju PM, et al. AHA/ACCF/HRS recommendations for the standardization and interpretation of the electrocardiogram. Circulation. 2009;119(10):e241-250. PMID: 19228821', 'Vandenberk B, et al. Which QT correction formulae to use for QT monitoring? J Am Heart Assoc. 2016;5(6):e003264. PMID: 27317349', 'Priori SG, et al. 2015 ESC Guidelines for management of patients with ventricular arrhythmias and prevention of sudden cardiac death. Eur Heart J. 2015;36(41):2793-2867. PMID: 26320108'],
  updated_at = NOW()
WHERE id = '3a0d7e6b-c53d-47de-a50c-d7774e542835'::uuid;


UPDATE score_items
SET
  optimal_range_min = 360.0,
  optimal_range_max = 450.0,
  unit = 'ms',
  optimal_explanation = 'QTc 360-450ms em mulheres é fisiológico. QTc >460ms aumenta risco arrítmico; influência hormonal (estrogênio prolonga QT naturalmente)',
  clinical_interpretation = '## Fundamento Fisiológico

Mulheres apresentam QTc fisiologicamente 10-20ms mais longo que homens devido ao efeito do estrogênio sobre canais de potássio cardíacos (IKr), prolongando a repolarização. Essa diferença hormonal explica maior incidência de torsades de pointes em mulheres (2-3x). O QTc varia durante ciclo menstrual (mais longo na fase lútea), gravidez e menopausa.

## Interpretação por Faixa de Valores

**360-450ms (Normal em mulheres):** Repolarização fisiológica, sem risco arrítmico aumentado. Variação até 30ms entre fases do ciclo menstrual é esperada.

**451-460ms (Limítrofe):** Zona cinzenta, comum em mulheres jovens saudáveis. Investigar causas secundárias se sintomática ou mudança recente em ECGs prévios.

**461-480ms (Prolongado leve):** QTc >460ms define prolongamento em mulheres. Risco moderado de torsades, especialmente se hipocalemia/hipomagnesemia concomitante.

**>480ms (Prolongado grave):** Alto risco arrítmico (>500ms = risco crítico). Requer investigação urgente e manejo especializado. Evitar absolutamente QT-prolongadores.

**<360ms (QT curto):** Incomum, sugere canalopatia ou distúrbio metabólico (hipercalcemia, hipermagnesemia).

## Variações Hormonais Fisiológicas

**Ciclo menstrual:** QTc aumenta 10-15ms na fase lútea (dias 14-28) versus fase folicular. Maior risco de arritmias no período pré-menstrual.

**Gravidez:** QTc encurta discretamente no 1º-2º trimestre, retorna ao basal no 3º. Monitorar se síndrome do QT longo conhecida.

**Menopausa:** Queda estrogênica pode encurtar QTc em 10-20ms. TRH pode prolongar novamente.',
  medical_conduct = '## QTc Normal (360-450ms)

- Manter equilíbrio eletrolítico: Magnésio 400mg/dia, Potássio via dieta (3-4 porções frutas/dia)
- Hidratação adequada (2L/dia)
- Cautela ao iniciar medicamentos QT-prolongadores (antibióticos, antidepressivos)

## QTc Limítrofe/Prolongado (451-480ms)

- Dosagem imediata: K+, Mg++, Ca++ iônico, TSH
- Revisar medicações: priorizar suspensão de QT-prolongadores
- Reposição eletrolítica agressiva:
  - Magnésio quelado 600-800mg/dia (alvo >2,0 mg/dL)
  - K+ se <3,8: banana, abacate, água de coco, KCl suplementar
- ECG de controle em 1 semana
- Se QTc >470ms persistente: Holter 24h, avaliação cardiológica

## QTc >480ms (Alto Risco)

- Suspensão imediata de todos QT-prolongadores
- Se sintomas (síncope, palpitações): emergência cardiológica
- Reposição IV de Mg++ (2g) se torsades de pointes
- Teste genético para LQTS (genes KCNQ1, KCNH2, SCN5A) se >500ms
- Beta-bloqueador (propranolol 40-80mg 2x/dia ou nadolol 40-80mg/dia) se LQTS congênita
- Evitar exercício intenso até estabilização

## Considerações Específicas em Mulheres

- Monitorar QTc na fase lútea (dias 21-28) se história de palpitações pré-menstruais
- Suplementação de Magnésio pode reduzir variação cíclica
- Em TRH: ECG basal antes de iniciar, repetir 3 meses depois
- Gravidez planejada com LQTS: otimizar beta-bloqueador pré-concepção

## Monitoramento

- QTc limítrofe: ECG a cada 3-6 meses
- QTc >470ms: ECG mensal, Holter 24h anual
- Após correção eletrolítica: ECG 7-14 dias',
  related_articles = ARRAY['Rautaharju PM, et al. Sex differences in the evolution of the electrocardiographic QT interval with age. Can J Cardiol. 1992;8(7):690-695. PMID: 1422988', 'Sedlak T, et al. Sex hormones and the QT interval: a review. J Womens Health. 2012;21(9):933-941. PMID: 22663191', 'Bai CX, et al. Distinct effects of 17β-estradiol on cardiac repolarization. Circulation. 2005;111(11):1365-1373. PMID: 15753228'],
  updated_at = NOW()
WHERE id = '2e3c06ce-bcb6-4649-984e-8c30a92e26f4'::uuid;


COMMIT;

-- Total de items atualizados: 5
