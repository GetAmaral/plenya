-- ============================================================================
-- BATCH 5.4 - ENRIQUECIMENTO DE EXAMES LABORATORIAIS - PARTE 4
-- Total: 20 items
-- Foco: Metano, Ecocardiografia, TYG Index, Enzimas, Hormônios, Proteínas
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. Metano Pico (CH₄ Máximo)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Medida do pico de metano no teste respiratório (lactulose ou glicose). Valores ≥10 ppm indicam supercrescimento bacteriano intestinal (SIBO) com predomínio de arqueias metanogênicas (Methanobrevibacter smithii). Associado a constipação crônica, flatulência e distensão abdominal. Metano reduz motilidade intestinal e está correlacionado com obesidade e disbiose. Identifica SIBO metanogênico com especificidade >90%. Valores >20 ppm indicam sobrecarga metanogênica severa requerendo antibioticoterapia combinada.$cr$,
  patient_explanation = $pe$Metano é um gás produzido por microorganismos específicos (arqueias) no seu intestino. Quando o teste mostra metano acima de 10 ppm, indica que você tem um supercrescimento dessas arqueias no intestino delgado (SIBO metanogênico). Isso explica sintomas como constipação crônica difícil de tratar, barriga inchada, gases e desconforto. O metano literalmente "freia" os movimentos intestinais. Níveis acima de 20 ppm são considerados severos. O tratamento envolve antibióticos específicos, dieta low-FODMAP temporária e medicamentos que estimulam o intestino a se mover melhor.$pe$,
  conduct = $cd$Valores ótimos MFI: < 10 ppm (negativo para SIBO metanogênico). Interpretação: < 10 ppm: Negativo → manter prebióticos e fibras fermentáveis; 10-20 ppm: SIBO metanogênico leve → rifaximina 550mg 3x/dia + neomicina 500mg 2x/dia por 14 dias, dieta low-FODMAP 4-6 semanas; > 20 ppm: SIBO severo → antibioticoterapia combinada prolongada, procinéticos (prucaloprida 2mg/dia), dieta low-FODMAP rigorosa. Investigar: hipocloridria (betaína HCl), motilidade (TSH, glicemia), anatomia (aderências, divertículos). Pós-tratamento: reavaliar em 4 semanas, reintroduzir prebióticos gradualmente (PHGG 5g/dia), corrigir causas subjacentes.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '47222b99-19cd-466f-9443-2ec818a30625';

-- ============================================================================
-- 2. Ecodopplercardiograma - E/e' Média
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Razão entre velocidade de fluxo mitral (onda E) e velocidade do anel mitral (e'). Indicador acurado de pressão diastólica final do ventrículo esquerdo (PDFVE). E/e' média >14 indica disfunção diastólica com elevação de pressões de enchimento. Valores intermediários (8-14) requerem avaliação adicional. Preditor independente de mortalidade cardiovascular. Essencial para diagnóstico de insuficiência cardíaca com fração de ejeção preservada (ICFEp). E/e' >14 correlaciona com pressão atrial esquerda elevada, risco de fibrilação atrial e prognóstico cardiovascular adverso.$cr$,
  patient_explanation = $pe$Este é um exame de ecocardiograma que avalia como seu coração relaxa e enche de sangue entre as batidas (função diastólica). E/e' mede a relação entre duas velocidades do fluxo sanguíneo. Valores acima de 14 indicam que seu coração está "rígido" e tem dificuldade para relaxar, aumentando a pressão interna. Isso pode causar falta de ar, cansaço e retenção de líquidos, mesmo que a "força de contração" esteja normal. Valores entre 8-14 são zona de atenção. Abaixo de 8 é ótimo. Controlar pressão arterial rigorosamente e perder peso são essenciais.$pe$,
  conduct = $cd$Valores ótimos MFI: E/e' média < 8 (função diastólica normal). Interpretação: E/e' < 8: Ótimo → manutenção com exercício aeróbico moderado 150min/semana, controle PA <130/80; E/e' 8-14: Zona intermediária → otimizar controle pressórico rigoroso, IECA/BRA, considerar espironolactona 25mg se HAS resistente, perda de peso 5-10%; E/e' > 14: Disfunção diastólica → diuréticos (furosemida 20-40mg ajustado), IECA/BRA titulado, controle volemia, restrição sal <2g/dia, BNP. Investigar causas: hipertrofia VE (RMC), amiloidose (troponina, BNP, cintilografia com pirofosfato), isquemia miocárdica (cintilografia).$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '09a6852e-bf3c-44ae-b307-8cb78e736f5f';

-- ============================================================================
-- 3. TYG INDEX (Índice Triglicerídeos-Glicose)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Índice calculado: ln[TG (mg/dL) × Glicemia jejum (mg/dL) / 2]. Marcador robusto de resistência insulínica, superior ao HOMA-IR em predição de risco cardiovascular e diabetes tipo 2. TYG >8.8 equivale a resistência insulínica moderada-severa. Correlaciona fortemente com adiposidade visceral, disfunção endotelial e DHGNA. Preditor de risco cardiovascular, síndrome metabólica e progressão de aterosclerose subclínica. TYG Index >9.0 aumenta risco de eventos cardiovasculares em 2-3x e de diabetes tipo 2 em 4-5x em 5 anos.$cr$,
  patient_explanation = $pe$TYG Index é um cálculo matemático simples usando seus triglicerídeos e glicemia de jejum. É um indicador poderoso de resistência insulínica - quando suas células não respondem bem à insulina. Valores acima de 8.8 indicam resistência moderada-severa, mesmo que sua glicemia pareça "normal". Acima de 9.0 multiplica por 3-4x seu risco de diabetes e infarto nos próximos anos. É uma ferramenta de prevenção excelente. A boa notícia: responde muito bem a mudanças no estilo de vida - dieta low-carb, jejum intermitente, exercício intenso e suplementos como berberina ou metformina.$pe$,
  conduct = $cd$Valores ótimos MFI: TYG Index < 8.5 (baixo risco de resistência insulínica). Interpretação: TYG < 8.5: Baixo risco → manter atividade física 150min/semana, dieta mediterrânea; TYG 8.5-9.0: Resistência insulínica leve → jejum intermitente 16:8, berberina 500mg 3x/dia ou mio-inositol 2g 2x/dia, reduzir carboidratos <150g/dia; TYG > 9.0: Resistência moderada-severa → metformina 850mg 2x/dia, dieta low-carb <100g/dia, HIIT 3x/semana, perda de peso 7-10%. Avaliar: circunferência abdominal, ALT, elastografia hepática (FibroScan), microalbuminúria, insulina basal.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '87b7f94c-4ef6-4451-9a12-b0d99b16817f';

-- ============================================================================
-- 4. Desidrogenase lática (LDH)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Enzima citosólica envolvida na conversão lactato-piruvato. Marcador de lesão celular inespecífico, presente em tecidos de alta atividade metabólica (músculo, fígado, hemácias, miocárdio). LDH elevado indica turnover celular aumentado, hemólise, lesão muscular ou hepática. Marcador de necrose tecidual, hemólise intravascular e atividade tumoral. Elevações agudas sugerem IAM (LDH1/LDH2 >1), hemólise (LDH + bilirrubina indireta), rabdomiólise (LDH + CK). Elevações crônicas podem indicar linfoma, leucemia ou anemia megaloblástica. Isoenzimas (LDH1-5) permitem localização tecidual.$cr$,
  patient_explanation = $pe$LDH é uma enzima presente dentro das células de todo o corpo - músculos, fígado, coração, sangue. Quando células são danificadas ou morrem, LDH "vaza" para o sangue. Níveis elevados são um sinal inespecífico de lesão celular. Pode indicar desde destruição de glóbulos vermelhos (hemólise), lesão muscular intensa, problemas no fígado, até situações mais graves como infarto ou tumores. Valores levemente elevados precisam ser investigados junto com outros exames. Valores muito altos (>500) requerem investigação urgente de hemólise ou malignidade.$pe$,
  conduct = $cd$Valores ótimos MFI: LDH 140-200 U/L. Interpretação: LDH < 200 U/L: Normal → acompanhamento de rotina; LDH 200-300 U/L: Elevação leve → investigar hemólise (haptoglobina, bilirrubinas, reticulócitos, esfregaço), lesão muscular (CK, aldolase), hepatopatia (AST, ALT, bilirrubinas); LDH > 300 U/L: Elevação moderada-severa → solicitar isoenzimas LDH1-5, avaliar esfregaço periférico, considerar malignidade hematológica; LDH >500 U/L + anemia: Hemólise aguda → Coombs direto, haptoglobina <25, reticulócitos >2%, bilirrubina indireta elevada; se anemia megaloblástica → B12, folato, homocisteína.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '8a9c0674-535f-41c6-8f98-34289a05ce75';

-- ============================================================================
-- 5. Dihidrotestosterona (DHT)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Metabólito ativo da testosterona via 5α-redutase. DHT tem afinidade 2-3x maior pelo receptor androgênico e não aromatiza para estradiol. Essencial para desenvolvimento de caracteres sexuais masculinos, libido e função erétil. DHT elevado associa-se a alopecia androgenética e hiperplasia prostática benigna (HPB). Avaliar balanço androgênico em casos de alopecia, acne severa, HPB ou hipogonadismo com testosterona normal. DHT baixo com testosterona normal sugere deficiência de 5α-redutase. DHT/T >10 indica conversão excessiva (risco HPB e calvície).$cr$,
  patient_explanation = $pe$DHT é a forma "super ativa" da testosterona, com potência 2-3x maior. É produzida a partir da testosterona pela enzima 5-alfa-redutase, principalmente na próstata, pele e folículos capilares. DHT é essencial para libido, função erétil e masculinização. Porém, DHT muito alto acelera calvície masculina e crescimento benigno da próstata (HPB). Homens com testosterona normal mas DHT baixo podem ter deficiência da enzima conversora. Relação DHT/testosterona acima de 10 indica conversão excessiva. Finasterida e dutasterida bloqueiam essa conversão (usados para calvície e próstata).$pe$,
  conduct = $cd$Valores ótimos MFI: Homens 30-85 ng/dL. Interpretação: DHT < 30 ng/dL: Déficit de conversão → otimizar zinco 30mg/dia, magnésio 400mg, avaliar uso de inibidores (finasterida, dutasterida, saw palmetto); DHT 30-85 ng/dL: Ótimo → manutenção com exercício resistido, sono 7-8h; DHT > 85 ng/dL: Conversão excessiva → avaliar HPB (PSA, IPSS, ultrassom prostático), alopecia androgenética; considerar saw palmetto 320mg/dia ou finasterida 1-5mg/dia se sintomático. DHT alto + acne: reduzir lácteos, açúcares, suplementar zinco 30mg + vitamina A 5.000 UI, retinoides tópicos.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '53726baa-6b52-4e90-9569-4d9bf2b33e26';

-- ============================================================================
-- 6. Eletroforese de proteínas
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Separação eletroforética de proteínas séricas em 5 frações: albumina (55-65%), α1-globulina (2-5%), α2-globulina (7-13%), β-globulina (8-14%), γ-globulina (11-21%). Padrão normal reflete síntese hepática, resposta imune e balanço proteico. Alterações indicam inflamação, doença hepática, gamopatias ou malabsorção. Triagem de gamopatias monoclonais (mieloma múltiplo), deficiências imunológicas, inflamação crônica e doenças hepáticas. Pico monoclonal (banda M) requer investigação com imunofixação e dosagem de cadeias leves. Hipoalbuminemia sugere desnutrição ou síndrome nefrótica.$cr$,
  patient_explanation = $pe$Este exame separa as proteínas do sangue em 5 grupos principais usando corrente elétrica. Cada grupo tem funções específicas: albumina (60%, produzida pelo fígado, mantém líquido nos vasos), alfa e beta globulinas (transporte e inflamação) e gama-globulinas (anticorpos do sistema imune). O padrão gráfico mostra a "saúde" do fígado, sistema imune e nutrição. Um "pico" anormal (banda M) pode indicar mieloma múltiplo (câncer de células plasmáticas). Albumina baixa sugere desnutrição ou perda renal. Gama-globulina alta indica infecção crônica ou autoimunidade.$pe$,
  conduct = $cd$Valores ótimos MFI: Albumina 4.0-5.0 g/dL, α1-globulina 0.1-0.3 g/dL, α2-globulina 0.6-1.0 g/dL, β-globulina 0.7-1.2 g/dL, γ-globulina 0.7-1.6 g/dL. Interpretação: Albumina <3.5: avaliar desnutrição (PCR, pré-albumina), síndrome nefrótica (proteinúria 24h), hepatopatia (INR, bilirrubinas); α1 baixa: investigar deficiência α1-antitripsina (fenótipo, espirometria); α2 elevada: inflamação aguda (PCR, VHS); Pico monoclonal (banda M): imunofixação + urinária, cadeias leves livres κ/λ, biópsia medula óssea; γ-globulina elevada: avaliar infecção crônica (HIV, hepatites), autoimunidade (ANA, FR), cirrose.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '688ab963-09d8-47e6-ab35-8e82a19a7a81';

-- ============================================================================
-- 7. Alfa-1 Globulina
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Fração composta principalmente por α1-antitripsina (AAT), proteína de fase aguda produzida pelo fígado. AAT inibe elastase neutrofílica, protegendo tecido pulmonar e hepático. Deficiência congênita (fenótipos PiZZ, PiSZ) predispõe a enfisema precoce (<45 anos) e cirrose. Elevações ocorrem em inflamação aguda. Triagem de deficiência de AAT, especialmente em enfisema precoce, DPOC em não-fumantes ou cirrose sem etiologia. α1-globulina <0.1 g/dL requer dosagem sérica de AAT e fenótipo genético.$cr$,
  patient_explanation = $pe$Alfa-1 globulina é composta principalmente por alfa-1 antitripsina (AAT), uma proteína protetora produzida pelo fígado que inibe enzimas destrutivas nos pulmões. Deficiência genética de AAT (herança dos pais) predispõe a enfisema pulmonar precoce (antes dos 45 anos) e cirrose hepática, especialmente em fumantes. Valores abaixo de 0.1 g/dL são suspeitos e requerem teste genético. Valores elevados indicam inflamação aguda (infecção, trauma). Pacientes com deficiência devem evitar rigorosamente tabaco, álcool e poluentes, e podem necessitar reposição IV da proteína.$pe$,
  conduct = $cd$Valores ótimos MFI: α1-globulina 0.1-0.3 g/dL (2-5% das proteínas totais). Interpretação: α1 < 0.1 g/dL: Suspeita deficiência AAT → dosagem sérica AAT (<80 mg/dL confirma), fenótipo Pi (PiZZ, PiSZ), espirometria, TC tórax; α1 0.1-0.3 g/dL: Normal → sem intervenção; α1 > 0.3 g/dL: Resposta fase aguda → investigar infecção, trauma ou inflamação (PCR, VHS). Deficiência confirmada: evitar tabagismo, poluentes, álcool; terapia de reposição AAT (Prolastin 60 mg/kg IV semanal) se FEV1 <65% predito; vacinação anual influenza e pneumococo.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '88081d50-7089-4f41-b463-c23347afedbc';

-- ============================================================================
-- 8. Beta Globulina
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Fração contendo transferrina (transporte de ferro), complemento (C3, C4), LDL, β2-microglobulina e hemopexina. β-globulina elevada sugere deficiência de ferro (↑transferrina), hipercolesterolemia ou gamopatia IgA. Redução indica malnutrição ou consumo de complemento (lúpus, vasculites). Avaliar status de ferro (transferrina é principal componente), hiperlipidemia e ativação imune. Elevações com hipoalbuminemia sugerem síndrome nefrótica. Reduções com inflamação indicam consumo de complemento.$cr$,
  patient_explanation = $pe$Beta-globulina inclui várias proteínas importantes: transferrina (transporta ferro no sangue), complemento (sistema imune), LDL colesterol e outras. Quando beta-globulina está elevada, geralmente indica deficiência de ferro (o corpo aumenta transferrina para captar mais ferro) ou colesterol alto. Quando está baixa, pode indicar desnutrição ou consumo de complemento em doenças autoimunes como lúpus. Elevação junto com albumina baixa sugere síndrome nefrótica (perda de proteína pela urina). Um "pico" na fração beta pode indicar gamopatia IgA.$pe$,
  conduct = $cd$Valores ótimos MFI: β-globulina 0.7-1.2 g/dL (8-14% das proteínas). Interpretação: β < 0.7 g/dL: investigar desnutrição (albumina, pré-albumina), consumo de complemento (C3, C4, CH50, ANA); β 0.7-1.2 g/dL: Normal → manutenção; β > 1.2 g/dL: avaliar deficiência de ferro (ferritina, CTLF, saturação transferrina), dislipidemia (LDL, colesterol total). β alta + hipoalbuminemia: síndrome nefrótica → proteinúria 24h, função renal, ultrassom renal. Pico na fração β: considerar gamopatia IgA → imunofixação, dosagem de imunoglobulinas quantitativas.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '6dd766d6-85cd-4916-a8cc-eb1b2f761450';

-- ============================================================================
-- 9. Estradiol - Homens
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Estradiol (E2) em homens é produzido por aromatização periférica de testosterona. Essencial para saúde óssea, função cognitiva, libido e regulação de lipídios. E2 <10 pg/mL associa-se a osteoporose, disfunção sexual e risco cardiovascular. E2 >40 pg/mL indica aromatização excessiva (obesidade, uso de testosterona exógena) e risco de ginecomastia. Avaliar balanço estrogênico em homens com ginecomastia, osteopenia, disfunção erétil ou em terapia de reposição de testosterona (TRT). E2 baixo em TRT requer ajuste de dose.$cr$,
  patient_explanation = $pe$Sim, homens produzem estradiol! Ele é produzido a partir da testosterona pela enzima aromatase, principalmente no tecido gorduroso. Estradiol é essencial para ossos fortes, memória, libido e saúde cardiovascular em homens. Valores muito baixos (<10 pg/mL) aumentam risco de osteoporose e disfunção sexual. Valores muito altos (>40 pg/mL) causam ginecomastia (crescimento de mama), retenção de líquidos e queda de libido. Homens obesos convertem mais testosterona em estradiol. Em terapia de reposição de testosterona, monitorar estradiol é crucial.$pe$,
  conduct = $cd$Valores ótimos MFI: Homens 20-30 pg/mL. Interpretação: E2 < 10 pg/mL: Déficit estrogênico → aumentar testosterona (se baixa), suplementar boro 6mg/dia, exercício resistido 3x/semana; E2 20-30 pg/mL: Ótimo → manutenção com exercício, sono adequado, evitar xenoestrogênios; E2 30-40 pg/mL: Limítrofe superior → reduzir adiposidade visceral, evitar álcool, suplementar DIM 200mg/dia ou crisina 500mg/dia; E2 > 40 pg/mL: Aromatização excessiva → inibidor de aromatase (anastrozol 0.25mg 2x/semana se em TRT), reduzir dose de testosterona, perda de peso 7-10%.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = 'b35c53c1-a97f-43de-9eae-c0e7ef1741f7';

-- ============================================================================
-- 10. Estradiol - Mulheres Pós-Menopausa (Com TRH)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Em mulheres pós-menopáusicas em terapia de reposição hormonal (TRH), estradiol deve mimetizar níveis de fase folicular precoce. E2 40-100 pg/mL proporciona benefícios sobre sintomas vasomotores, saúde óssea, cognitiva e cardiovascular. E2 <30 pg/mL indica dose insuficiente. E2 >150 pg/mL aumenta risco de trombose e câncer de mama. Monitorar adequação de TRH. Níveis terapêuticos reduzem fogachos, atrofia vaginal, perda óssea e risco de fraturas. Supradosagem aumenta risco de eventos tromboembólicos e neoplasias hormônio-dependentes.$cr$,
  patient_explanation = $pe$Mulheres na menopausa em terapia de reposição hormonal (TRH) devem ter estradiol entre 40-100 pg/mL para obter benefícios sem riscos excessivos. Nessa faixa, você tem alívio de fogachos, suores noturnos, atrofia vaginal, melhora de humor e proteção óssea. Valores abaixo de 30 pg/mL indicam dose insuficiente (sintomas persistem). Valores acima de 150 pg/mL indicam dose excessiva, aumentando risco de trombose (coágulos) e câncer de mama. Estradiol transdérmico (gel, adesivo) é mais seguro que oral. Sempre usar progesterona se você tem útero.$pe$,
  conduct = $cd$Valores ótimos MFI: Pós-menopausa com TRH 40-100 pg/mL. Interpretação: E2 < 30 pg/mL: Dose insuficiente → aumentar estradiol transdérmico (0.05-0.1 mg/dia) ou oral (1-2 mg/dia); E2 40-100 pg/mL: Ótimo → manutenção, reavaliar sintomas e densitometria óssea anualmente; E2 100-150 pg/mL: Limítrofe superior → considerar redução de dose se assintomática, avaliar risco trombótico (D-dímero, proteína C e S); E2 > 150 pg/mL: Supradosagem → reduzir dose, avaliar trombose (ultrassom venoso se sintomas), mamografia anual. Ajustar progesterona micronizada 100-200mg/noite se útero presente.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '89114370-2ab3-48fb-a542-b61a7e1fd979';

-- ============================================================================
-- 11. FAN (Fator Antinuclear)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Anticorpos contra componentes nucleares (DNA, histonas, nucleolos). Triagem para doenças autoimunes sistêmicas. FAN ≥1:160 com padrão específico (homogêneo, pontilhado, nucleolar) sugere lúpus, esclerodermia, Sjögren ou doença mista do tecido conjuntivo. FAN baixo (1:80) pode ocorrer em 5-10% da população saudável. FAN é altamente sensível (>95%) para lúpus eritematoso sistêmico (LES), mas pouco específico. FAN negativo praticamente exclui LES ativo. FAN positivo requer correlação clínica e autoanticorpos específicos (anti-dsDNA, anti-Sm, anti-Ro, anti-La).$cr$,
  patient_explanation = $pe$FAN é um teste de triagem para doenças autoimunes sistêmicas como lúpus, esclerodermia e Sjögren. Detecta anticorpos que atacam o núcleo das células. Resultado negativo ou abaixo de 1:80 é normal. Títulos de 1:160 ou mais com sintomas (artrite, manchas na pele, fotossensibilidade) requerem investigação com testes mais específicos. Importante: 5-10% de pessoas saudáveis têm FAN baixo positivo (1:80) sem doença. O "padrão" (homogêneo, pontilhado, nucleolar) orienta quais doenças investigar. FAN positivo isolado sem sintomas geralmente não requer tratamento.$pe$,
  conduct = $cd$Valores ótimos MFI: FAN não reagente ou título < 1:80. Interpretação: FAN não reagente: baixa probabilidade de autoimunidade → sem intervenção; FAN 1:80-1:160 + assintomático: monitorar anualmente, avaliar inflamação (PCR, VHS), otimizar vitamina D >40 ng/mL; FAN ≥1:160 + sintomas (artrite, fotossensibilidade, rash): solicitar anti-dsDNA, anti-Sm, ENA (anti-Ro, anti-La), complemento (C3, C4), hemograma, função renal. FAN padrão nucleolar: investigar esclerodermia → capilaroscopia, prova de função pulmonar, anti-Scl-70. FAN padrão citoplasmático: hepatite autoimune → transaminases, IgG, biópsia hepática.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '423066a6-80f6-45af-b6bc-d4882077f541';

-- ============================================================================
-- 12. Ferritina - Homens
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Proteína de armazenamento de ferro, reflete estoques corporais. Em homens, ferritina <70 ng/mL indica depleção de reservas, mesmo com hemoglobina normal. Ferritina >300 ng/mL sugere sobrecarga (hemocromatose), inflamação ou síndrome metabólica. Ferritina é proteína de fase aguda, podendo elevar-se falsamente em inflamação. Avaliar estoques de ferro e diferenciar anemia ferropriva de anemia de doença crônica. Ferritina baixa é específica para deficiência de ferro. Ferritina alta requer investigação de hemocromatose (saturação de transferrina, genotipagem HFE) ou inflamação (PCR).$cr$,
  patient_explanation = $pe$Ferritina é a principal proteína que armazena ferro no corpo. Em homens, valores abaixo de 70 ng/mL indicam estoques baixos de ferro, causando fadiga, queda de cabelo e redução de performance física, mesmo que você não tenha anemia ainda. Valores acima de 300 ng/mL podem indicar hemocromatose (sobrecarga genética de ferro), inflamação crônica ou síndrome metabólica. Hemocromatose não tratada pode causar cirrose, diabetes e problemas cardíacos. Ferritina também sobe com inflamação (infecções, artrite), então precisa ser avaliada junto com PCR.$pe$,
  conduct = $cd$Valores ótimos MFI: Homens 70-150 ng/mL. Interpretação: Ferritina < 30 ng/mL: Deficiência severa → ferro quelato 30mg/dia elementar em jejum + vitamina C 500mg, investigar sangramento oculto (sangue oculto, colonoscopia se >50 anos); Ferritina 30-70 ng/mL: Reserva subótima → suplementar ferro 15-30mg/dia, otimizar absorção (carnes vermelhas, evitar chá/café nas refeições); Ferritina 70-150 ng/mL: Ótimo → manutenção; Ferritina > 300 ng/mL: avaliar saturação transferrina (se >45%: investigar hemocromatose HFE C282Y/H63D), PCR (inflamação), elastografia hepática, glicemia.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = 'bf4a95b2-2c34-4a1c-96bb-07ca58662932';

-- ============================================================================
-- 13. Ferritina - Mulheres Pré-Menopausa
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Mulheres menstruadas têm perdas mensais de ferro (15-30 mg/mês), predispondo a ferropenia. Ferritina <50 ng/mL é comum e associa-se a fadiga, queda de cabelo, síndrome das pernas inquietas e disfunção tireoidiana. Otimizar ferritina melhora energia, cognição e performance física. Triagem de deficiência de ferro, causa mais comum de anemia em mulheres pré-menopáusicas. Ferritina <30 ng/mL indica depleção de estoques e requer reposição imediata. Ferritina 30-50 ng/mL associa-se a sintomas clínicos mesmo sem anemia.$cr$,
  patient_explanation = $pe$Mulheres que menstruam perdem ferro mensalmente e frequentemente têm estoques baixos. Ferritina abaixo de 50 ng/mL causa cansaço crônico, queda de cabelo, unhas quebradiças, pernas inquietas à noite e dificuldade de concentração, mesmo que você não tenha anemia. Abaixo de 15 ng/mL é deficiência severa. Otimizar ferritina para 50-100 ng/mL melhora dramaticamente energia, cabelo e performance física. Ferro quelato ou bisglicinato são melhores tolerados que sulfato ferroso. Tomar em jejum com vitamina C aumenta absorção. Se fluxo menstrual intenso, pode precisar suplementar continuamente.$pe$,
  conduct = $cd$Valores ótimos MFI: Mulheres pré-menopausa 50-100 ng/mL. Interpretação: Ferritina < 15 ng/mL: Deficiência severa → ferro IV (carboximaltose férrica 500-1000mg dose única) se intolerância oral ou má absorção; Ferritina 15-50 ng/mL: Deficiência leve-moderada → ferro quelato 30mg/dia + vitamina C 500mg, bisglicinato de ferro 25mg/dia (melhor tolerância); Ferritina 50-100 ng/mL: Ótimo → manutenção, considerar suplementação periódica se fluxo menstrual intenso; Ferritina > 200 ng/mL: investigar inflamação (PCR), resistência insulínica (glicemia, insulina), hemocromatose se saturação transferrina >45%.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '70b77e5e-0d09-44be-92bb-3aaad9c448a7';

-- ============================================================================
-- 14. Fibrinogênio
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Glicoproteína de fase aguda produzida pelo fígado, essencial para formação do coágulo (conversão em fibrina). Fibrinogênio elevado (>400 mg/dL) é marcador independente de risco cardiovascular, associado a aterosclerose, trombose e síndrome metabólica. Deficiência (<100 mg/dL) causa diátese hemorrágica. Avaliar risco trombótico e inflamatório. Fibrinogênio >400 mg/dL aumenta risco de IAM, AVE e TEV em 2-3x. Elevações persistentes sugerem inflamação crônica subclínica. Reduções indicam consumo (CIVD) ou deficiência congênita.$cr$,
  patient_explanation = $pe$Fibrinogênio é uma proteína produzida pelo fígado essencial para coagulação do sangue. Quando há um corte, fibrinogênio se transforma em fibrina formando o coágulo. Valores acima de 400 mg/dL indicam inflamação crônica e aumentam 2-3x o risco de infarto, derrame e trombose. É um marcador independente de risco cardiovascular. Valores muito baixos (<100) causam sangramentos. Fibrinogênio alto responde bem a estilo de vida: perder peso, parar de fumar, exercício aeróbico e suplementos anti-inflamatórios como ômega-3 e nattokinase.$pe$,
  conduct = $cd$Valores ótimos MFI: Fibrinogênio 200-300 mg/dL. Interpretação: Fibrinogênio < 150 mg/dL: investigar CIVD (D-dímero, plaquetas, TAP/PTT), deficiência congênita (dosagem funcional), hepatopatia severa (INR, albumina); Fibrinogênio 200-300 mg/dL: Ótimo → manutenção com dieta anti-inflamatória, exercício aeróbico 150min/semana; Fibrinogênio 300-400 mg/dL: Limítrofe → otimizar estilo de vida (reduzir tabagismo, obesidade), suplementar ômega-3 2-3g/dia, nattokinase 100mg/dia; Fibrinogênio > 400 mg/dL: Risco CV aumentado → investigar inflamação (PCR-us, VHS), síndrome metabólica, considerar AAS 100mg/dia se alto risco CV.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '2e962724-dbaa-47bb-8ea4-73a17991b15c';

-- ============================================================================
-- 15. Fosfatase alcalina
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Enzima presente em osso, fígado, intestino e placenta. FA elevada sugere doença hepatobiliar (colestase), doença óssea (Paget, osteomalacia) ou crescimento ósseo (adolescência, fraturas). FA baixa (<30 U/L) pode indicar deficiência de zinco, magnésio ou hipofosfatasia. Avaliar colestase (FA elevada + GGT elevada) versus doença óssea (FA elevada + GGT normal). FA >150 U/L requer investigação hepatobiliar (ultrassom, CPRM) ou óssea (cálcio, PTH, vitamina D, cintilografia óssea). Isoenzimas diferenciam origem tecidual.$cr$,
  patient_explanation = $pe$Fosfatase alcalina (FA) é uma enzima presente nos ossos, fígado, intestino e outros tecidos. Valores elevados geralmente indicam problemas no fígado (bile presa - colestase) ou nos ossos (doença de Paget, osteomalacia, fraturas cicatrizando). Para diferenciar, medimos GGT: se GGT também está alto, é fígado; se GGT está normal, é osso. Valores baixos (<30) podem indicar deficiências nutricionais (zinco, magnésio) ou doença genética rara (hipofosfatasia). Adolescentes têm FA naturalmente elevado pelo crescimento ósseo.$pe$,
  conduct = $cd$Valores ótimos MFI: Fosfatase alcalina 40-100 U/L. Interpretação: FA < 40 U/L: deficiência nutricional → otimizar zinco 30mg/dia, magnésio 400mg, vitamina B6 50mg; FA 40-100 U/L: Ótimo → manutenção; FA 100-150 U/L: elevação leve → investigar GGT (se alto: hepatopatia; se normal: doença óssea), vitamina D, cálcio, PTH; FA > 150 U/L + GGT alto: colestase → ultrassom abdominal, bilirrubinas, considerar CPRM, biópsia hepática; FA > 150 U/L + GGT normal: doença óssea → cintilografia óssea (Paget, metástases), PTH (hiperparatireoidismo), vitamina D <20 (osteomalacia).$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '741a272e-4674-4864-9602-dd92148c7287';

-- ============================================================================
-- 16. FSH - Mulheres Pós-menopausa
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Hormônio folículo-estimulante elevado reflete falência ovariana por depleção folicular. FSH >30-40 mUI/mL com amenorreia >12 meses confirma menopausa. FSH persistentemente alto indica ausência de feedback estrogênico negativo. FSH >100 mUI/mL é típico de menopausa estabelecida. Confirmar transição menopáusica em mulheres com sintomas (fogachos, irregularidade menstrual) e idade adequada (45-55 anos). FSH <30 mUI/mL em mulher >55 anos com amenorreia sugere insuficiência ovariana prematura secundária ou hipopituitarismo.$cr$,
  patient_explanation = $pe$FSH é o hormônio que estimula os ovários a produzir óvulos. Após a menopausa, os ovários param de responder, então o cérebro aumenta FSH tentando "forçar" os ovários a funcionar. FSH acima de 30-40 mUI/mL junto com ausência de menstruação por 12 meses confirma menopausa. Valores acima de 100 mUI/mL são típicos. FSH baixo (<30) com amenorreia prolongada em mulher acima de 55 anos é incomum e pode indicar problema na hipófise. FSH alto confirma que sintomas como fogachos, suores e irregularidade menstrual são realmente da menopausa.$pe$,
  conduct = $cd$Valores ótimos MFI: Pós-menopausa FSH > 30 mUI/mL. Interpretação: FSH < 30 mUI/mL + amenorreia >12 meses + idade >50: menopausa atípica → avaliar estradiol (<20 pg/mL confirma), considerar causas secundárias (prolactina, TSH, LH); FSH 30-100 mUI/mL: menopausa confirmada → otimizar TRH se indicada (estradiol transdérmico 0.05mg + progesterona micronizada 100mg); FSH > 100 mUI/mL: menopausa estabelecida → foco em saúde óssea (densitometria, vitamina D >40 ng/mL, cálcio 1200mg/dia, exercício resistido 3x/semana). Sintomas vasomotores: TRH bioidêntica, venlafaxina 37.5-75mg/dia, gabapentina 300mg 3x/dia.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '922dd181-3945-4a99-8304-6839b7907185';

-- ============================================================================
-- 17. pH Venoso
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Medida de íons hidrogênio no sangue venoso. pH <7.32 indica acidose (metabólica ou respiratória), pH >7.38 indica alcalose. pH venoso é ~0.03-0.05 unidades menor que arterial. Útil para triagem de distúrbios ácido-base sem necessidade de gasometria arterial dolorosa. Avaliar equilíbrio ácido-base em pacientes ambulatoriais. Acidose metabólica (pH baixo + HCO3 baixo) sugere cetoacidose, acidose láctica ou insuficiência renal. Alcalose metabólica (pH alto + HCO3 alto) sugere vômitos ou uso de diuréticos.$cr$,
  patient_explanation = $pe$pH mede o equilíbrio entre ácido e base (alcalino) no sangue. O corpo mantém pH rigorosamente entre 7.32-7.38 (sangue venoso) para funcionamento celular adequado. pH abaixo de 7.32 indica acidose (muito ácido), causada por diabetes descompensado, insuficiência renal, diarreia ou acúmulo de ácido lático. pH acima de 7.38 indica alcalose (muito alcalino), geralmente por vômitos ou uso excessivo de diuréticos. pH venoso é uma forma menos invasiva de avaliar equilíbrio ácido-base, evitando punção arterial dolorosa.$pe$,
  conduct = $cd$Valores ótimos MFI: pH venoso 7.32-7.38. Interpretação: pH < 7.30: acidose significativa → gasometria arterial completa, dosagem de lactato, cetona sérica, ânion gap, função renal (ureia, creatinina); pH 7.30-7.32: acidose leve → investigar causas (diabetes, diarreia, IRC), otimizar hidratação, considerar bicarbonato oral se acidose metabólica crônica; pH 7.32-7.38: Ótimo → sem intervenção; pH > 7.40: alcalose → investigar vômitos, uso de diuréticos, hiperventilação, dosar eletrólitos (K, Cl), corrigir causa base. Acidose severa <7.20: emergência médica.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '2c0c344e-f932-4e6b-9ece-88f6882efa51';

-- ============================================================================
-- 18. pO2 Venoso
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Pressão parcial de oxigênio no sangue venoso, reflete oxigenação tecidual e extração de O2. pO2 venoso <35 mmHg indica hipoxemia venosa, sugerindo aumento de extração periférica (hipoperfusão, anemia, insuficiência cardíaca) ou hipoxemia arterial. pO2 venoso >50 mmHg pode indicar shunt ou baixa extração (sepse, cianeto). Avaliar oxigenação tecidual e adequação de perfusão. Útil em pacientes com insuficiência cardíaca, choque ou sepse. pO2 venoso misto (veia cava superior) <28 mmHg indica choque grave.$cr$,
  patient_explanation = $pe$pO2 venoso mede quanto oxigênio ainda resta no sangue depois que passou pelos tecidos. Valores normais (35-45 mmHg) indicam que os tecidos estão recebendo oxigênio adequadamente. Valores abaixo de 35 mmHg indicam que os tecidos estão "famintos" de oxigênio - pode ser por insuficiência cardíaca, anemia, ou falta de oxigênio no sangue arterial. Valores acima de 50 mmHg podem indicar que os tecidos não estão conseguindo extrair oxigênio (sepse, intoxicação por cianeto) ou que há um curto-circuito (shunt) no coração.$pe$,
  conduct = $cd$Valores ótimos MFI: pO2 venoso 35-45 mmHg. Interpretação: pO2 < 30 mmHg: hipoxemia venosa severa → gasometria arterial, avaliar perfusão (lactato, tempo enchimento capilar), considerar choque ou hipoxemia arterial grave; pO2 30-35 mmHg: hipoxemia leve → otimizar oxigenação (suplementação O2 se SatO2 <92%), avaliar hemoglobina (meta >10 g/dL), perfusão cardíaca (BNP, ecocardiograma); pO2 35-45 mmHg: Ótimo → manutenção; pO2 > 50 mmHg: baixa extração de O2 → investigar sepse (lactato, procalcitonina, hemoculturas), shunt intracardíaco (ecocardiograma com contraste), intoxicação por cianeto.$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '0899a7f5-e9d5-4adf-a825-4374ecc3292f';

-- ============================================================================
-- 19. Gliadina Deaminada IgA (DGP-IgA)
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Anticorpo contra gliadina deaminada, marcador de doença celíaca com especificidade >95%. DGP-IgA complementa anti-transglutaminase tecidual (tTG-IgA) e tem melhor sensibilidade em crianças <2 anos. Positivo indica resposta imune contra glúten, requerendo biópsia duodenal para confirmação. Triagem de doença celíaca em pacientes com diarreia crônica, perda de peso, anemia ferropriva refratária ou osteoporose precoce. DGP-IgA >20 U/mL com tTG-IgA positivo tem VPP >95% para doença celíaca. Falsos-negativos ocorrem em deficiência de IgA.$cr$,
  patient_explanation = $pe$DGP-IgA detecta anticorpos contra uma forma modificada de gliadina (proteína do glúten). É um marcador altamente específico (>95%) de doença celíaca. Teste positivo (>20 U/mL) indica que seu sistema imune está atacando o glúten, mas diagnóstico definitivo requer biópsia do intestino delgado (endoscopia). Doença celíaca causa má absorção intestinal, diarreia crônica, perda de peso, anemia por deficiência de ferro, osteoporose e fadiga. Tratamento é dieta totalmente sem glúten para a vida toda. Teste funciona melhor quando feito enquanto você ainda come glúten.$pe$,
  conduct = $cd$Valores ótimos MFI: DGP-IgA < 20 U/mL (negativo). Interpretação: DGP-IgA < 20 U/mL: negativo → doença celíaca improvável (se tTG-IgA também negativo), considerar outras causas de sintomas; DGP-IgA 20-50 U/mL: positivo leve → dosar tTG-IgA, IgA total (excluir deficiência de IgA), considerar biópsia duodenal; DGP-IgA > 50 U/mL: positivo forte → alta probabilidade doença celíaca → endoscopia digestiva alta com biópsias duodenais (4-6 fragmentos de 2ª e 3ª porções duodenais). Se confirmada: dieta isenta de glúten vitalícia, suplementar ferro, cálcio, vitamina D, B12, zinco; reavaliar DGP-IgA em 6-12 meses (deve negativar).$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '05b8bdbe-ce1b-44c6-9cbf-178320c23119';

-- ============================================================================
-- 20. Hemograma completo
-- ============================================================================
UPDATE score_items
SET
  clinical_relevance = $cr$Avaliação quantitativa e qualitativa de células sanguíneas. Hemoglobina baixa indica anemia (microcítica: ferro; macrocítica: B12/folato; normocítica: doença crônica, hemólise). Leucopenia (<4.000) sugere imunossupressão ou autoimunidade. Leucocitose (>11.000) indica infecção ou inflamação. Plaquetopenia (<150.000) requer investigação de consumo ou produção deficiente. Triagem de anemias, infecções, neoplasias hematológicas e distúrbios de coagulação. Hemograma alterado orienta investigações específicas (VCM, RDW, esfregaço periférico, ferritina, B12, mielograma).$cr$,
  patient_explanation = $pe$Hemograma completo avalia três tipos de células do sangue: glóbulos vermelhos (transportam oxigênio - anemia se baixo), glóbulos brancos (combatem infecções - baixo indica imunidade fraca, alto indica infecção ou inflamação) e plaquetas (coagulação - baixo aumenta risco de sangramentos). VCM (tamanho das hemácias) orienta causa da anemia: VCM baixo (<80) sugere falta de ferro, VCM alto (>100) sugere falta de B12 ou folato. Hemograma é exame básico essencial que detecta desde deficiências nutricionais até leucemias e linfomas.$pe$,
  conduct = $cd$Valores ótimos MFI: Hemoglobina H 14-16 g/dL, M 13-15 g/dL | Leucócitos 5.000-7.500/µL | Plaquetas 200.000-300.000/µL | VCM 85-95 fL. Interpretação: Anemia microcítica (VCM <80): investigar ferro (ferritina, saturação transferrina), talassemia (eletroforese Hb se ferritina normal), doença crônica (PCR); Anemia macrocítica (VCM >100): dosar B12, folato, TSH; avaliar álcool, medicamentos (metotrexato), SMD (esfregaço periférico); Leucopenia <3.500: investigar autoimunidade (ANA, anti-DNA), deficiências (B12, folato, cobre), medicamentos; Leucocitose >11.000: avaliar infecção (PCR, procalcitonina), leucemia (esfregaço, se blastos → mielograma); Plaquetopenia <100.000: investigar PTI, consumo (D-dímero), produção (mielograma).$cd$,
  last_review = CURRENT_TIMESTAMP
WHERE id = '36f3af60-3772-413f-9338-73a98cef11f1';

-- ============================================================================
-- COMMIT E VERIFICAÇÃO
-- ============================================================================

-- Verifica quantos registros foram atualizados
DO $$
DECLARE
  updated_count INTEGER;
BEGIN
  SELECT COUNT(*) INTO updated_count
  FROM score_items
  WHERE id IN (
    '47222b99-19cd-466f-9443-2ec818a30625',
    '09a6852e-bf3c-44ae-b307-8cb78e736f5f',
    '87b7f94c-4ef6-4451-9a12-b0d99b16817f',
    '8a9c0674-535f-41c6-8f98-34289a05ce75',
    '53726baa-6b52-4e90-9569-4d9bf2b33e26',
    '688ab963-09d8-47e6-ab35-8e82a19a7a81',
    '88081d50-7089-4f41-b463-c23347afedbc',
    '6dd766d6-85cd-4916-a8cc-eb1b2f761450',
    'b35c53c1-a97f-43de-9eae-c0e7ef1741f7',
    '89114370-2ab3-48fb-a542-b61a7e1fd979',
    '423066a6-80f6-45af-b6bc-d4882077f541',
    'bf4a95b2-2c34-4a1c-96bb-07ca58662932',
    '70b77e5e-0d09-44be-92bb-3aaad9c448a7',
    '2e962724-dbaa-47bb-8ea4-73a17991b15c',
    '741a272e-4674-4864-9602-dd92148c7287',
    '922dd181-3945-4a99-8304-6839b7907185',
    '2c0c344e-f932-4e6b-9ece-88f6882efa51',
    '0899a7f5-e9d5-4adf-a825-4374ecc3292f',
    '05b8bdbe-ce1b-44c6-9cbf-178320c23119',
    '36f3af60-3772-413f-9338-73a98cef11f1'
  )
  AND clinical_relevance IS NOT NULL
  AND last_review >= CURRENT_TIMESTAMP - INTERVAL '1 minute';

  RAISE NOTICE '============================================================================';
  RAISE NOTICE 'BATCH 5.4 - ENRIQUECIMENTO COMPLETO';
  RAISE NOTICE '============================================================================';
  RAISE NOTICE 'Total de items enriquecidos: %', updated_count;
  RAISE NOTICE 'Esperado: 20 items';

  IF updated_count = 20 THEN
    RAISE NOTICE 'STATUS: SUCESSO - Todos os 20 items foram enriquecidos com sucesso!';
  ELSE
    RAISE WARNING 'STATUS: ATENCAO - Apenas % de 20 items foram atualizados', updated_count;
  END IF;
  RAISE NOTICE '============================================================================';
END $$;

COMMIT;
