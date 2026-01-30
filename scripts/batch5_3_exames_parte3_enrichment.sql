-- =====================================================
-- BATCH 5.3: EXAMES LABORATORIAIS PARTE 3
-- 44 itens enriquecidos com padrão MFI
-- Executar: cat batch5_3_exames_parte3_enrichment.sql | docker compose exec -T db psql -U plenya_user -d plenya_db
-- =====================================================

BEGIN;

-- 1. VDRL
UPDATE score_items
SET
    interpretation = 'VDRL (Venereal Disease Research Laboratory) é teste não treponêmico usado no rastreamento e monitoramento de sífilis. Detecta anticorpos contra cardiolipina liberada por células danificadas pela infecção. Em medicina funcional, VDRL positivo pode indicar infecção ativa por Treponema pallidum, exigindo confirmação com testes treponêmicos (FTA-ABS, TPHA). Falsos positivos podem ocorrer em doenças autoimunes (LES), gravidez, vacinações recentes e infecções virais. Títulos VDRL também são usados para monitorar resposta ao tratamento antibiótico.',

    clinical_relevance = 'VDRL é fundamental no diagnóstico de sífilis primária, secundária e latente. A abordagem funcional enfatiza: interpretação combinada com história clínica e exames treponêmicos; diferenciação entre infecção ativa e cicatriz sorológica; monitoramento de títulos (queda de 4x após tratamento indica sucesso); investigação de falsos positivos em contextos autoimunes; rastreamento em gestantes (sífilis congênita) e parceiros sexuais. VDRL positivo sem confirmação treponêmica pode refletir reatividade cruzada com cardiolipinas endógenas, vista em LES e síndrome antifosfolípide.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'VDRL ótimo: NÃO REAGENTE (negativo) em indivíduos sem história de sífilis. Em pacientes tratados, títulos devem reduzir 4x (2 diluições) em 6-12 meses após terapia adequada. Persistência de títulos baixos (≤1:4) pode indicar cicatriz sorológica (serofast), comum em sífilis tardia tratada. Títulos crescentes ou ≥1:16 sugerem reinfecção ou falha terapêutica. Em gestantes, VDRL deve ser negativo ou com títulos estáveis baixos se tratadas previamente.',

    health_recommendations = 'VDRL reagente:
• Confirmar com testes treponêmicos (FTA-ABS, TPHA, TP-PA)
• Penicilina benzatina 2.4 milhões UI IM: dose única (sífilis primária/secundária) ou 3 doses semanais (latente tardia)
• Avaliar parceiros sexuais e contatos dos últimos 90 dias
• Repetir VDRL em 3, 6, 12 meses para monitorar declínio de títulos

Falsos positivos (VDRL+ e treponêmicos negativos):
• Investigar LES, síndrome antifosfolípide, hepatite viral, mononucleose
• Repetir após 3 meses para confirmar transitoriedade

Prevenção:
• Uso consistente de preservativos
• Testagem regular em populações de risco (HSH, múltiplos parceiros)
• Rastreamento universal em gestantes no 1º trimestre e parto',

    last_review_date = CURRENT_DATE
WHERE id = '3348de39-de90-4792-832b-80ec942db408';

-- 2. Piridoxal-5-fosfato (PLP) por HPLC
UPDATE score_items
SET
    interpretation = 'Piridoxal-5-fosfato (PLP) é a forma bioativa da vitamina B6, atuando como coenzima em >140 reações metabólicas, incluindo síntese de neurotransmissores (serotonina, dopamina, GABA), metabolismo de homocisteína (via transulfuração), gliconeogênese e síntese de hemoglobina. PLP sanguíneo reflete status funcional de B6 melhor que piridoxina total. Deficiência pode resultar de ingestão inadequada, má absorção, uso de medicamentos (isoniazida, hidralazina, L-dopa), alcoolismo crônico ou polimorfismos genéticos em enzimas de conversão (NBPF3). Níveis elevados são raros, mas podem indicar suplementação excessiva.',

    clinical_relevance = 'PLP é essencial para saúde neurológica, cardiovascular e metabólica. A medicina funcional valoriza dosagem de PLP por HPLC em: sintomas neuropsiquiátricos (depressão, ansiedade, neuropatia periférica); elevação de homocisteína refratária a B12/folato; anemia microcítica não responsiva a ferro; dermatite seborreica e queilite angular; síndrome pré-menstrual e náuseas gravídicas; usuários de medicamentos depletores (contraceptivos orais, terapia hormonal). Deficiência subclínica (<30 nmol/L) é comum em idosos, gestantes e portadores de MTHFR C677T.',

    optimal_value_male_min = 35.0,
    optimal_value_male_max = 110.0,
    optimal_value_female_min = 35.0,
    optimal_value_female_max = 110.0,
    optimal_value_description = 'PLP ótimo: 35-110 nmol/L (padrão MFI). Valores <30 nmol/L indicam deficiência funcional, associada a comprometimento na síntese de neurotransmissores e metilação. Faixa 30-35 nmol/L é insuficiente para máxima função enzimática. Níveis >110 nmol/L sugerem suplementação ativa (geralmente >50 mg/dia piridoxina). Não há toxicidade aguda, mas doses crônicas >200 mg/dia podem causar neuropatia sensorial reversível. Faixa ideal otimiza metabolismo de homocisteína e síntese de GABA.',

    health_recommendations = 'PLP < 35 nmol/L:
• Piridoxal-5-fosfato 25-50 mg/dia (forma ativa, biodisponibilidade superior à piridoxina)
• Ou piridoxina HCl 50-100 mg/dia (necessita conversão hepática)
• Combinar com magnésio 300-400 mg/dia (cofator para conversão B6 → PLP)
• Fontes alimentares: salmão (0.9 mg/100g), frango (0.5 mg), banana (0.4 mg), batata (0.3 mg)

Deficiência com homocisteína elevada:
• Protocolo trimetilglicina: B6 (50 mg) + B12 metilcobalamina (1000 mcg) + folato 5-MTHF (800 mcg)

Protocolo neuropatia:
• PLP 100 mg 2x/dia + ácido alfa-lipóico 600 mg/dia

Reavaliar PLP em 8-12 semanas
Evitar doses crônicas >200 mg/dia piridoxina',

    last_review_date = CURRENT_DATE
WHERE id = 'ee1942f1-25f7-4b6e-8b43-fed81edd14af';

-- 3. Zinco
UPDATE score_items
SET
    interpretation = 'Zinco é oligoelemento essencial, cofator para >300 enzimas envolvidas em síntese proteica, replicação de DNA, função imune, cicatrização e apoptose. Zinco sérico reflete apenas 0.1% do zinco corporal total (maioria intracelular), mas é marcador prático de status. Deficiência é comum em vegetarianos (fitatos limitam absorção), idosos, alcoolistas, portadores de doença inflamatória intestinal e usuários de inibidores de bomba de prótons (omeprazol). Zinco baixo compromete imunidade celular (linfócitos T), função barreira intestinal, síntese de testosterona e conversão T4→T3.',

    clinical_relevance = 'Medicina funcional valoriza zinco em: infecções recorrentes (resfriados, candidíase); cicatrização lenta e lesões cutâneas; queda de cabelo e alopecia areata; hipogonadismo e disfunção erétil; disgeusia (alteração de paladar) e hiposmia; acne e dermatite; hipotireoidismo subclínico (zinco ativa desiodase tipo 1); resistência à insulina e diabetes tipo 2. Relação zinco/cobre deve ser 0.7-1.0; desequilíbrio (cobre alto/zinco baixo) aumenta estresse oxidativo e inflamação crônica.',

    optimal_value_male_min = 90.0,
    optimal_value_male_max = 120.0,
    optimal_value_female_min = 90.0,
    optimal_value_female_max = 120.0,
    optimal_value_description = 'Zinco ótimo: 90-120 mcg/dL (padrão MFI). Valores <70 mcg/dL indicam deficiência clínica, com comprometimento imune e reprodutivo. Faixa 70-90 mcg/dL é insuficiente para máxima função enzimática, especialmente em contextos de infecção ou cicatrização. Níveis >130 mcg/dL sugerem suplementação excessiva, podendo inibir absorção de cobre e ferro (efeito inibitório competitivo no DMT1 intestinal). Ideal: 100-110 mcg/dL para otimizar imunidade, síntese hormonal e metabolismo tireoidiano.',

    health_recommendations = 'Zinco < 90 mcg/dL:
• Zinco quelado (picolinato, bisglicinato) 30-50 mg/dia (elemento) em jejum
• Dosagem terapêutica aguda (resfriado): 75 mg/dia por 5-7 dias
• Fontes alimentares: ostras (74 mg/100g), carne vermelha (7 mg), sementes de abóbora (7 mg), castanhas (3 mg)

Absorção otimizada:
• Tomar longe de cálcio, ferro e fitatos (cereais integrais)
• Combinar com vitamina A (retinol aumenta mobilização de zinco hepático)

Protocolo acne/alopecia:
• Zinco 50 mg/dia + vitamina A 10,000 UI/dia por 3 meses

Monitorar relação zinco/cobre (dosar cobre sérico)
Reavaliar em 12 semanas
Evitar doses crônicas >100 mg/dia (risco de deficiência de cobre)',

    last_review_date = CURRENT_DATE
WHERE id = 'd1f27505-2fe7-48aa-a54e-67723b8c035d';

-- 4. Vitamina E (Alfa-Tocoferol)
UPDATE score_items
SET
    interpretation = 'Vitamina E (alfa-tocoferol) é principal antioxidante lipossolúvel, protegendo membranas celulares da peroxidação lipídica. Atua em sinergia com selênio, glutationa e vitamina C. Alfa-tocoferol é a forma mais biodisponível das 8 formas de vitamina E (4 tocoferóis, 4 tocotrienóis). Deficiência é rara em adultos saudáveis, mas pode ocorrer em má absorção de gorduras (colestase, doença celíaca, fibrose cística), deficiência de lipoproteínas (abetalipoproteinemia) e prematuridade. Níveis baixos aumentam fragilidade eritrocitária e neuropatia periférica.',

    clinical_relevance = 'Medicina funcional valoriza vitamina E em: esteatose hepática não alcoólica (EHNA) - tocoferóis melhoram esteatose e inflamação; dislipidemias e prevenção de oxidação LDL (reduz aterogênese); neuropatia periférica em diabéticos; degeneração macular relacionada à idade (DMRI); claudicação intermitente e doença arterial periférica; síndrome pré-menstrual (mastalgia). Importante: suplementação com alfa-tocoferol isolado depleta gama-tocoferol, que tem ação anti-inflamatória única (inibe COX-2). Fórmulas balanceadas (mixed tocopherols) são preferíveis.',

    optimal_value_male_min = 12.0,
    optimal_value_male_max = 20.0,
    optimal_value_female_min = 12.0,
    optimal_value_female_max = 20.0,
    optimal_value_description = 'Vitamina E ótima: 12-20 mg/L (padrão MFI). Valores <5.5 mg/L indicam deficiência clínica, com risco de anemia hemolítica e neuropatia. Faixa 5.5-12 mg/L é insuficiente para proteção antioxidante máxima, especialmente em contextos de estresse oxidativo elevado (tabagismo, poluição, exercício intenso). Níveis >20 mg/L sugerem suplementação ativa. Não há toxicidade significativa até 1000 mg/dia, mas doses >400 UI/dia podem aumentar risco de sangramento (inibição de agregação plaquetária).',

    health_recommendations = 'Vitamina E < 12 mg/L:
• Vitamina E natural (d-alfa-tocoferol) 400 UI/dia com refeição contendo gordura
• Preferir mixed tocopherols (alfa + beta + gama + delta) para evitar depleção de gama-tocoferol
• Fontes alimentares: óleo de gérmen de trigo (149 mg/100g), amêndoas (26 mg), avelãs (15 mg), espinafre (2 mg)

Protocolo EHNA:
• Vitamina E 800 UI/dia + pioglitazona 30 mg/dia (protocolo PIVENS - melhora esteatose e fibrose)

Protocolo cardiovascular:
• Vitamina E 400 UI/dia + CoQ10 100 mg/dia + ômega-3 2g/dia

DMRI:
• Fórmula AREDS2: vitamina E 400 UI + luteína 10 mg + zeaxantina 2 mg + zinco 80 mg + cobre 2 mg

Reavaliar em 12 semanas
Evitar doses >800 UI/dia sem supervisão',

    last_review_date = CURRENT_DATE
WHERE id = '80ffc2b2-545b-4389-891f-b6aba1f7865c';

-- 5. Triglicerídeos
UPDATE score_items
SET
    interpretation = 'Triglicerídeos são principais lipídeos de armazenamento energético, transportados em VLDL e quilomícrons. Elevação reflete excesso calórico (especialmente carboidratos refinados e álcool), resistência à insulina, síntese hepática aumentada e clearance deficiente (lipoproteína lipase reduzida). Hipertrigliceridemia moderada (150-500 mg/dL) aumenta risco cardiovascular via aterosclerose residual (partículas remanescentes ricas em triglicerídeos). Níveis muito elevados (>500 mg/dL) aumentam risco de pancreatite aguda. Triglicerídeos baixos (<70 mg/dL) podem indicar desnutrição, má absorção ou hipertireoidismo.',

    clinical_relevance = 'Medicina funcional enfatiza triglicerídeos como marcador de: resistência à insulina (TG/HDL ≥3 prediz síndrome metabólica melhor que glicemia isolada); excesso de frutose e açúcares (conversão hepática preferencial frutose→triglicerídeos); inflamação crônica de baixo grau; esteatose hepática (TG >150 correlaciona com acúmulo hepático); hipotireoidismo subclínico (T3 regula lipoproteína lipase). Triglicerídeos pós-prandiais (2h após refeição) ≥220 mg/dL indicam clearance deficiente e maior risco cardiovascular que jejum isolado.',

    optimal_value_male_min = 50.0,
    optimal_value_male_max = 100.0,
    optimal_value_female_min = 50.0,
    optimal_value_female_max = 100.0,
    optimal_value_description = 'Triglicerídeos ótimos: 50-100 mg/dL (padrão MFI). Valores <50 mg/dL são raros, mas podem sugerir hipertireoidismo ou desnutrição. Faixa 100-150 mg/dL ainda é "normal" por critérios convencionais, mas indica início de disfunção metabólica (resistência insulínica incipiente). Níveis >150 mg/dL definem síndrome metabólica. Ideal: 70-90 mg/dL, associado a melhor sensibilidade insulínica e menor risco cardiovascular. Meta agressiva em diabéticos: <70 mg/dL para reduzir aterosclerose residual.',

    health_recommendations = 'TG > 100 mg/dL:
• Restrição rigorosa de carboidratos refinados (<100g/dia carboidratos líquidos)
• Eliminar açúcar, frutose adicionada, sucos, refrigerantes
• Reduzir/eliminar álcool (etanol bloqueia oxidação lipídica hepática)
• Aumentar fibras solúveis 25-35g/dia (β-glucana aveia, psyllium)

TG 150-500 mg/dL (adicionar):
• Ômega-3 EPA+DHA 2-4g/dia (reduz síntese hepática VLDL em 25-30%)
• Niacina liberação lenta 1000-2000 mg/dia (inibe lipólise adiposa - precaução com rubor)
• Berberina 500 mg 3x/dia (ativa AMPK, mimetiza metformina)

TG > 500 mg/dL (risco pancreatite):
• Fenofibrato 160 mg/dia (ativa PPAR-α, aumenta lipoproteína lipase)
• Dieta cetogênica <20g carboidratos/dia
• Evitar gorduras saturadas temporariamente até TG <200 mg/dL

Reavaliar em 6-8 semanas
Meta agressiva: <80 mg/dL',

    last_review_date = CURRENT_DATE
WHERE id = 'dfa58a95-7d7a-491f-b7c5-d2fe8c694daa';

-- 6. Alfa-2 Globulina
UPDATE score_items
SET
    interpretation = 'Alfa-2 globulinas incluem proteínas de fase aguda (haptoglobina, alfa-2 macroglobulina, ceruloplasmina) e transportadoras (alfa-2 antiplasmina). Elevação ocorre em inflamação aguda e crônica, infecções, neoplasias, síndrome nefrótica (perda seletiva de albumina preserva globulinas) e uso de corticoides. Alfa-2 macroglobulina inibe proteases (tripsina, plasmina) e pode estar elevada em cirrose, diabetes e queimaduras. Redução é menos comum, vista em hemólise intravascular (depleção de haptoglobina), desnutrição proteica e doença de Wilson (ceruloplasmina baixa).',

    clinical_relevance = 'Medicina funcional valoriza alfa-2 globulinas em: monitoramento de inflamação crônica de baixo grau (complementa PCR ultrassensível); diagnóstico diferencial de proteinúria (alfa-2 elevada + albumina baixa sugere perda glomerular); rastreamento de síndrome nefrótica; investigação de hemólise (haptoglobina baixa); suspeita de doença de Wilson (ceruloplasmina <20 mg/dL). Alfa-2 globulinas elevadas isoladamente, sem infecção aparente, podem refletir inflamação metabólica (obesidade, resistência insulínica, disbiose intestinal).',

    optimal_value_male_min = 0.6,
    optimal_value_male_max = 1.0,
    optimal_value_female_min = 0.6,
    optimal_value_female_max = 1.0,
    optimal_value_description = 'Alfa-2 globulinas ótimas: 0.6-1.0 g/dL (padrão MFI). Valores <0.6 g/dL podem indicar hemólise, desnutrição proteica ou doença de Wilson (dosar ceruloplasmina e cobre sérico/urinário). Faixa >1.0 g/dL sugere inflamação ativa, infecção ou síndrome nefrótica. Ideal: 0.7-0.9 g/dL, refletindo ausência de inflamação aguda e função hepática/renal preservada. Alfa-2 persistentemente elevada (>1.2 g/dL) sem causa infecciosa requer investigação de neoplasia oculta ou doença autoimune.',

    health_recommendations = 'Alfa-2 globulinas > 1.0 g/dL:
• Investigar infecção ativa (hemograma, PCR, VHS)
• Avaliar função renal (ureia, creatinina, proteinúria 24h)
• Descartar síndrome nefrótica (albumina <3.5 g/dL + proteinúria >3.5g/24h)
• Rastrear neoplasias se persistentemente elevada (TC tórax/abdome, marcadores tumorais)

Protocolo anti-inflamatório (se inflamação crônica):
• Ômega-3 2-3g/dia (reduz citocinas pró-inflamatórias)
• Curcumina lipossomal 1000 mg/dia (inibe NF-κB)
• Dieta anti-inflamatória (eliminar glúten, laticínios, açúcar por 30 dias)
• Probióticos 50 bilhões UFC/dia (modula inflamação intestinal)

Alfa-2 < 0.6 g/dL:
• Investigar hemólise (reticulócitos, LDH, bilirrubina indireta, haptoglobina)
• Doença de Wilson: ceruloplasmina + cobre sérico + cobre urinário 24h + exame de lâmpada de fenda (anel de Kayser-Fleischer)

Reavaliar eletroforese de proteínas em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e';

-- 7. VCM (MCV)
UPDATE score_items
SET
    interpretation = 'Volume Corpuscular Médio (VCM) mede tamanho médio das hemácias. VCM alto (macrocitose >100 fL) sugere deficiência de B12/folato, hipotireoidismo, alcoolismo, mielodisplasia ou uso de antifolatos (metotrexato). VCM baixo (microcitose <80 fL) indica anemia ferropriva, talassemia ou anemia de doença crônica. VCM normal (80-100 fL) não exclui anemia, podendo haver deficiência mista (ferro + B12) ou anemia normocrômica-normocítica (doença renal, inflamação). VCM é primeiro passo para classificação morfológica de anemias.',

    clinical_relevance = 'Medicina funcional valoriza VCM em: rastreamento de deficiências nutricionais subclínicas (B12 no limite superior 90-100 fL pode indicar deficiência incipiente); monitoramento de vegetarianos/veganos (risco B12); detecção de alcoolismo oculto (VCM >95 fL mesmo sem anemia); diagnóstico de hipotireoidismo subclínico (VCM >94 fL + TSH >3.0 mUI/L); diferenciação entre deficiência de ferro (VCM <80 fL + ferritina <30 ng/mL) e talassemia minor (VCM <75 fL + ferritina normal). VCM ideal: 85-92 fL para otimizar oxigenação tecidual.',

    optimal_value_male_min = 85.0,
    optimal_value_male_max = 92.0,
    optimal_value_female_min = 85.0,
    optimal_value_female_max = 92.0,
    optimal_value_description = 'VCM ótimo: 85-92 fL (padrão MFI). Valores <80 fL indicam microcitose, sugerindo deficiência de ferro ou talassemia. Faixa 80-85 fL é normocítica inferior, mas pode refletir depleção de ferro incipiente (checar ferritina <50 ng/mL). Valores >100 fL indicam macrocitose, associada a deficiência de B12/folato ou hipotireoidismo. Faixa 92-100 fL é normocítica superior, mas pode sinalizar deficiência subclínica de B12 (dosar metilmalônico e homocisteína). Ideal: 88-90 fL para hemácias de tamanho ótimo e capacidade carreadora de oxigênio máxima.',

    health_recommendations = 'VCM < 85 fL (microcitose):
• Investigar com: ferritina (<30 ng/mL sugere deficiência), ferro sérico, saturação de transferrina (<20% confirma deficiência)
• Ferro quelado 100-200 mg/dia (elemento) em jejum com vitamina C 500 mg
• Se ferritina normal: considerar talassemia minor (hemoglobina A2 >3.5%)
• Fontes alimentares: carne vermelha (2.5 mg/100g), fígado bovino (6 mg), lentilhas (3 mg)

VCM > 92 fL (tendência macrocitose):
• Dosar B12 sérica (<400 pg/mL), metilmalônico (>300 nmol/L confirma deficiência), homocisteína (>12 μmol/L)
• B12 metilcobalamina 1000 mcg sublingual/dia ou 1000 mcg IM semanal por 4 semanas
• Folato 5-MTHF 800-1000 mcg/dia (sempre combinar com B12)
• Avaliar função tireoidiana (TSH, T4 livre)
• Investigar alcoolismo oculto (GGT, AST, história)

VCM > 100 fL:
• Descartar hipotireoidismo (TSH >3.0 mUI/L) e mielodisplasia (em idosos com citopenia)

Reavaliar hemograma em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21';

-- 8. Progesterona - Homens
UPDATE score_items
SET
    interpretation = 'Progesterona em homens é produzida principalmente em adrenais e testículos, servindo como precursor para síntese de corticosteroides e andrógenos. Níveis são naturalmente baixos (<1.0 ng/mL), mas podem estar elevados em hiperplasia adrenal congênita (déficit de 21-hidroxilase ou 17-hidroxilase), tumores adrenais produtores de progesterona ou uso exógeno de progestágenos. Progesterona elevada em homens pode causar ginecomastia, disfunção erétil e redução de libido, pois compete com andrógenos em receptores e pode inibir 5-alfa-redutase.',

    clinical_relevance = 'Medicina funcional valoriza progesterona em homens em: investigação de ginecomastia inexplicada; avaliação de hiperplasia adrenal congênita de início tardio; monitoramento de reposição hormonal bioidêntica (alguns protocolos incluem progesterona para modular conversão testosterona→DHT); rastreamento de tumores adrenais (progesterona >2.0 ng/mL em homem requer imagem de adrenais). Progesterona pode ter efeito neuroprotetor (promove mielinização, reduz neuroinflamação), sendo investigada em TBI e doenças neurodegenerativas.',

    optimal_value_male_min = 0.1,
    optimal_value_male_max = 0.8,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Progesterona ótima em homens: 0.1-0.8 ng/mL (padrão MFI). Valores <0.1 ng/mL são raros e sem significado clínico. Faixa >1.0 ng/mL requer investigação de hiperplasia adrenal congênita (dosar 17-hidroxiprogesterona, cortisol, ACTH). Níveis >2.0 ng/mL sugerem tumor adrenal produtor de progesterona (realizar TC ou RM de adrenais). Ideal: 0.2-0.6 ng/mL, refletindo produção adrenal basal sem hipersecreção patológica.',

    health_recommendations = 'Progesterona > 1.0 ng/mL em homens:
• Investigar hiperplasia adrenal congênita:
  - 17-hidroxiprogesterona basal e pós-estímulo com ACTH
  - Cortisol, DHEA-S, androstenediona
  - Teste genético se 17-OHP >10 ng/mL (mutações CYP21A2)

• Se ginecomastia + progesterona elevada:
  - Suspender anabolizantes, finasterida, espironolactona
  - Anastrozol 0.5 mg 2x/semana (inibe aromatização)
  - Cirurgia se mama >2 cm (mastectomia subcutânea)

Progesterona > 2.0 ng/mL:
• TC adrenais com contraste (descartar adenoma/carcinoma)
• Se tumor: excisão cirúrgica + análise histopatológica

Progesterona normal alta (0.8-1.0 ng/mL) com sintomas:
• Otimizar testosterona (meta 600-900 ng/dL)
• Zinco 30 mg/dia + vitamina E 400 UI/dia (modulam receptores hormonais)

Reavaliar em 3 meses',

    last_review_date = CURRENT_DATE
WHERE id = '5e465089-1492-44c4-9e7b-6696731a56a3';

-- 9. Gama GT
UPDATE score_items
SET
    interpretation = 'Gama-glutamiltransferase (GGT) é enzima de membrana presente em fígado, rins, pâncreas e bile. Elevação reflete colestase (obstrução biliar intra ou extra-hepática), lesão hepatocelular, consumo excessivo de álcool (induz síntese de GGT), esteatose hepática, uso de medicamentos hepatotóxicos (fenitoína, rifampicina) e síndrome metabólica. GGT é marcador mais sensível que fosfatase alcalina para doenças hepatobiliares. Valores elevados sem outros marcadores hepáticos alterados sugerem consumo de álcool ou síndrome metabólica. GGT também é marcador de estresse oxidativo sistêmico.',

    clinical_relevance = 'Medicina funcional valoriza GGT como: marcador de consumo alcoólico oculto (GGT >50 U/L em homens, >30 U/L em mulheres); indicador de esteatose hepática não alcoólica (EHNA) - GGT >40 U/L correlaciona com fibrose hepática; preditor de risco cardiovascular e metabólico (GGT >30 U/L associa-se a resistência insulínica, mesmo com transaminases normais); monitoramento de abstinência alcoólica (GGT normaliza em 4-6 semanas após parar álcool). Relação AST/GGT <1 sugere doença hepática alcoólica; >1 favorece outras etiologias.',

    optimal_value_male_min = 10.0,
    optimal_value_male_max = 30.0,
    optimal_value_female_min = 7.0,
    optimal_value_female_max = 25.0,
    optimal_value_description = 'GGT ótima: 10-30 U/L (homens), 7-25 U/L (mulheres) - padrão MFI. Valores >50 U/L (homens) ou >35 U/L (mulheres) indicam disfunção hepatobiliar ou consumo alcoólico. Faixa 30-50 U/L (homens) ou 25-35 U/L (mulheres) já sinaliza estresse hepático incipiente ou síndrome metabólica. Ideal: 15-25 U/L (homens), 10-20 U/L (mulheres), refletindo função hepática ótima e baixo estresse oxidativo. GGT <10 U/L é rara, mas sem implicação clínica negativa.',

    health_recommendations = 'GGT > 30 U/L (mulheres) ou > 40 U/L (homens):
• Abstinência total de álcool por 8 semanas
• Investigar esteatose hepática: ultrassom abdome, TG/HDL (>3 sugere EHNA)
• Revisar medicamentos hepatotóxicos (paracetamol, estatinas, AINEs)

Protocolo EHNA:
• Perda de peso 7-10% (melhora esteatose e inflamação)
• Dieta low-carb <100g/dia ou mediterrânea
• Silimarina (cardo-mariano) 420 mg/dia dividida 3x (estabiliza membranas hepatócitos)
• Vitamina E 800 UI/dia (reduz inflamação hepática)
• N-acetilcisteína (NAC) 600 mg 2x/dia (precursor glutationa, antioxidante)

GGT > 100 U/L:
• Descartar colestase: fosfatase alcalina, bilirrubinas, ultrassom vias biliares
• Sorologias virais: hepatite B (HBsAg), hepatite C (anti-HCV)
• Investigar hepatotoxicidade medicamentosa ou autoimune

Protocolo abstinência alcoólica:
• Tiamina (B1) 100 mg/dia (previne encefalopatia de Wernicke)
• Ácido fólico 1 mg/dia + complexo B
• Milk thistle 300 mg 3x/dia

Reavaliar GGT em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = 'ccfde47c-b3ca-4465-91d2-cab643ae08d2';

-- 10. Progesterona - Gestação 2º Trimestre
UPDATE score_items
SET
    interpretation = 'Progesterona no 2º trimestre gestacional (13-27 semanas) é produzida principalmente pela placenta (corpo lúteo involui após 10-12 semanas). Progesterona mantém endométrio decidualizado, suprime contrações uterinas (relaxamento miometrial), modula resposta imune materna (tolerância fetal), estimula desenvolvimento das glândulas mamárias e previne rejeição fetal. Níveis baixos podem indicar insuficiência placentária, risco de trabalho de parto prematuro ou restrição de crescimento intrauterino (RCIU). Progesterona elevada é comum e fisiológica na gestação.',

    clinical_relevance = 'Medicina funcional valoriza progesterona no 2º trimestre em: avaliação de risco de parto prematuro (níveis <40 ng/mL em 20-24 semanas associam-se a maior risco); monitoramento de gestações de alto risco (história de abortos recorrentes, insuficiência cervical); investigação de RCIU (progesterona baixa + doppler umbilical anormal); protocolo de prevenção de parto prematuro (progesterona vaginal 200 mg/dia reduz risco em mulheres com colo curto <25 mm). Progesterona >100 ng/mL é esperada no 2º trimestre, refletindo função placentária adequada.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 60.0,
    optimal_value_female_max = 150.0,
    optimal_value_description = 'Progesterona ótima no 2º trimestre: 60-150 ng/mL (padrão MFI). Valores <40 ng/mL sugerem insuficiência placentária e risco aumentado de parto prematuro ou RCIU. Faixa 40-60 ng/mL é limite inferior, exigindo monitoramento fetal intensivo (ultrassom morfológico, doppler). Níveis >150 ng/mL são comuns no final do 2º trimestre e indicam função placentária robusta. Ideal: 80-120 ng/mL, associado a crescimento fetal adequado e baixo risco de complicações.',

    health_recommendations = 'Progesterona < 60 ng/mL no 2º trimestre:
• Ultrassom morfológico detalhado + doppler artéria umbilical e uterina
• Avaliar comprimento cervical (colo <25 mm indica risco de parto prematuro)
• Progesterona vaginal micronizada 200 mg/dia (reduz risco de parto prematuro em 45%)
• Repouso relativo se colo curto ou contrações frequentes

Prevenção parto prematuro (progesterona <70 ng/mL + colo <30 mm):
• Progesterona vaginal 200 mg ao deitar até 36 semanas
• 17-alfa-hidroxiprogesterona caproato 250 mg IM semanal (se história de parto prematuro prévio)
• Cerclagem cervical se colo <15 mm antes de 24 semanas

Suporte nutricional gestacional:
• Ácido fólico 1 mg/dia (reduz defeitos tubo neural)
• Ferro 60 mg/dia + vitamina C 500 mg (previne anemia)
• Ômega-3 DHA 300 mg/dia (neurodesenvolvimento fetal)
• Vitamina D 2000-4000 UI/dia (meta 25-OH >40 ng/mL)
• Magnésio 400 mg/dia (reduz contrações uterinas prematuras)

Reavaliar progesterona + ultrassom em 4 semanas
Monitoramento fetal intensivo no 3º trimestre',

    last_review_date = CURRENT_DATE
WHERE id = '60c5b79e-7a16-4043-b25f-bf00c43a928a';

-- 11. Ferritina - Pós-Menopausa
UPDATE score_items
SET
    interpretation = 'Ferritina em mulheres pós-menopausa reflete estoques de ferro sem perdas menstruais. Níveis podem elevar-se progressivamente após menopausa devido à ausência de sangramento mensal. Ferritina >200 ng/mL pode indicar sobrecarga de ferro, aumentando estresse oxidativo (reação de Fenton: Fe²⁺ + H₂O₂ → radicais hidroxila), inflamação crônica ou hemocromatose hereditária. Ferritina <30 ng/mL indica deficiência de ferro, mas níveis <50 ng/mL já comprometem função tireoidiana (ferro é cofator para peroxidase tireoideana) e síntese de neurotransmissores.',

    clinical_relevance = 'Medicina funcional valoriza ferritina em pós-menopausa para: diagnóstico de deficiência de ferro (ferritina <50 ng/mL causa fadiga, queda de cabelo, síndrome das pernas inquietas mesmo sem anemia); rastreamento de sobrecarga de ferro (ferritina >200 ng/mL aumenta risco de diabetes, doença hepática, aterosclerose); investigação de hemocromatose em mulheres (manifestações surgem após menopausa, quando ferro acumula-se rapidamente); monitoramento de inflamação crônica (ferritina é reagente de fase aguda - valores >300 ng/mL podem refletir inflamação, não sobrecarga).',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 50.0,
    optimal_value_female_max = 100.0,
    optimal_value_description = 'Ferritina ótima pós-menopausa: 50-100 ng/mL (padrão MFI). Valores <30 ng/mL indicam deficiência de ferro absoluta. Faixa 30-50 ng/mL é insuficiente para função enzimática ótima (tireoide, mitocôndrias). Níveis >150 ng/mL podem indicar sobrecarga incipiente ou inflamação crônica (dosar PCR para diferenciar). Valores >300 ng/mL requerem investigação de hemocromatose (saturação de transferrina >45%, genotipagem HFE C282Y/H63D). Ideal: 70-90 ng/mL para otimizar energia e função cognitiva sem risco oxidativo.',

    health_recommendations = 'Ferritina < 50 ng/mL:
• Ferro quelado (bisglicinato) 100-200 mg/dia em jejum com vitamina C 500 mg
• Evitar chá, café, cálcio junto ao ferro (reduzem absorção)
• Fontes alimentares: carne vermelha (2.5 mg/100g), fígado bovino (6 mg), espinafre cozido (3 mg)
• Investigar causas: sangramento gastrointestinal oculto (pesquisa sangue oculto, endoscopia/colonoscopia), má absorção (doença celíaca, gastrite atrófica)

Ferritina 150-300 ng/mL:
• Dosar PCR ultrassensível (se elevada, ferritina reflete inflamação)
• Saturação de transferrina (<45% descarta sobrecarga)
• Se sobrecarga confirmada: flebotomia 1 unidade (450 mL) a cada 2-4 semanas até ferritina 50-100 ng/mL

Ferritina > 300 ng/mL:
• Rastrear hemocromatose: saturação transferrina, genotipagem HFE (C282Y/H63D)
• Se hemocromatose: flebotomia semanal até ferritina <50 ng/mL, depois manutenção trimestral
• Evitar suplementos de ferro, vitamina C em altas doses, álcool

Protocolo antioxidante (se ferritina >200 ng/mL):
• Quercetina 500 mg 2x/dia (quelante de ferro)
• Curcumina 1000 mg/dia (reduz estresse oxidativo)
• Chá verde 3 xícaras/dia (polifenóis quelam ferro)

Reavaliar ferritina em 12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = '9d9f1270-d24d-4236-8694-5e28d8748475';

-- 12. DHEA-S - Homens 20-29 anos
UPDATE score_items
SET
    interpretation = 'DHEA-S (sulfato de dehidroepiandrosterona) é andrógeno adrenal, marcador de função da zona reticular das glândulas adrenais. Pico ocorre aos 20-30 anos, declinando ~2% ao ano após 30 anos (adrenopausa). DHEA-S serve como precursor para síntese periférica de testosterona e estrógenos. Níveis baixos podem indicar insuficiência adrenal primária (Addison), hipopituitarismo, uso de corticoides ou envelhecimento acelerado. DHEA-S elevado sugere hiperplasia adrenal congênita (déficit de 21-hidroxilase), tumores adrenais ou síndrome de Cushing.',

    clinical_relevance = 'Medicina funcional valoriza DHEA-S em jovens (20-29 anos) para: avaliar resposta ao estresse crônico (DHEA-S <250 mcg/dL sugere exaustão adrenal); investigar fadiga persistente, baixa libido ou dificuldade de ganho muscular; diagnóstico de insuficiência adrenal primária ou secundária; rastreamento de hiperplasia adrenal congênita de início tardio (DHEA-S >700 mcg/dL). DHEA-S tem efeitos neuroprotetores (melhora memória, humor), imunomoduladores (aumenta IgG, células NK) e anabólicos (síntese proteica muscular).',

    optimal_value_male_min = 350.0,
    optimal_value_male_max = 500.0,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'DHEA-S ótimo em homens 20-29 anos: 350-500 mcg/dL (padrão MFI). Valores <250 mcg/dL sugerem insuficiência adrenal ou exaustão crônica. Faixa 250-350 mcg/dL é subótima, associada a menor reserva anabólica e resiliência ao estresse. Níveis >700 mcg/dL requerem investigação de hiperplasia adrenal (dosar 17-hidroxiprogesterona, cortisol) ou tumor adrenal. Ideal: 400-450 mcg/dL, associado a melhor composição corporal, libido e função imune.',

    health_recommendations = 'DHEA-S < 350 mcg/dL:
• Suplementação DHEA 25-50 mg/dia pela manhã (mimetiza ritmo circadiano)
• Monitorar testosterona (DHEA converte-se perifericamente em andrógenos)
• Reavaliar em 8 semanas (meta: 400-500 mcg/dL)

Protocolo suporte adrenal:
• Vitamina C 2000 mg/dia dividida (adrenais concentram vitamina C)
• Pantetina (vitamina B5) 500 mg 2x/dia (precursor coenzima A, síntese hormonal)
• Ashwagandha 600 mg/dia (adaptógeno, reduz cortisol, eleva DHEA)
• Rhodiola rosea 400 mg/dia (melhora fadiga, modula HPA)

DHEA-S < 250 mcg/dL:
• Investigar insuficiência adrenal: cortisol basal 8h (<10 mcg/dL suspeito), ACTH, teste estímulo ACTH
• Se Addison confirmado: hidrocortisona 15-25 mg/dia + fludrocortisona 0.1 mg/dia

DHEA-S > 700 mcg/dL:
• Dosar 17-hidroxiprogesterona (>10 ng/mL sugere CAH), cortisol, testosterona
• TC adrenais se suspeita de tumor
• Genotipagem CYP21A2 se CAH confirmada

Otimização estilo de vida:
• Sono 7-9h/noite (DHEA-S cai com privação de sono)
• Exercício resistido 3-4x/semana (estimula síntese DHEA)
• Reduzir estresse crônico (cortisol elevado inibe DHEA)

Reavaliar DHEA-S em 12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = 'a6742c97-a610-4bd4-b9de-87e91ecc8d34';

-- 13. FSH - Fase Lútea
UPDATE score_items
SET
    interpretation = 'FSH (hormônio folículo-estimulante) na fase lútea (dias 21-28 do ciclo) normalmente é baixo, pois corpo lúteo produz progesterona e inibina B, que suprimem FSH por feedback negativo. FSH lúteo >10 mUI/mL pode indicar insuficiência de corpo lúteo, reserva ovariana diminuída (início de perimenopausa) ou fase folicular tardia (coleta fora da janela ideal). FSH lúteo muito baixo (<1 mUI/mL) pode sugerir hipopituitarismo ou supressão por contraceptivo oral. FSH lúteo é menos usado clinicamente que FSH folicular (dia 3), mas pode sinalizar disfunção ovariana incipiente.',

    clinical_relevance = 'Medicina funcional valoriza FSH lúteo em: avaliação de insuficiência de corpo lúteo (fase lútea curta <10 dias, progesterona baixa); rastreamento precoce de diminuição de reserva ovariana em mulheres <40 anos (FSH lúteo >8 mUI/mL pode preceder elevação de FSH folicular); investigação de infertilidade com fase lútea inadequada (defeito de luteinização); monitoramento de resposta a tratamentos de fertilidade. FSH lúteo elevado sugere que ovário não responde adequadamente ao pico de LH, resultando em corpo lúteo deficiente.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 1.0,
    optimal_value_female_max = 8.0,
    optimal_value_description = 'FSH ótimo na fase lútea: 1.0-8.0 mUI/mL (padrão MFI). Valores >10 mUI/mL sugerem insuficiência de corpo lúteo ou diminuição de reserva ovariana (considerar dosar FSH no dia 3 do ciclo seguinte + HAM). Faixa 8-10 mUI/mL é limítrofe, indicando luteinização subótima. Níveis <1.0 mUI/mL podem refletir hipopituitarismo (dosar LH, TSH, prolactina) ou uso de contraceptivos. Ideal: 2-6 mUI/mL, refletindo feedback negativo adequado por progesterona/inibina do corpo lúteo.',

    health_recommendations = 'FSH lúteo > 8 mUI/mL:
• Confirmar fase do ciclo (idealmente dias 21-23 em ciclo de 28 dias)
• Dosar progesterona lútea (dia 21) - meta >10 ng/mL; <5 ng/mL confirma insuficiência
• Avaliar reserva ovariana: FSH dia 3 (<10 mUI/mL normal), HAM (>1.0 ng/mL normal), contagem de folículos antrais (>10 total)

Insuficiência corpo lúteo confirmada:
• Progesterona micronizada 200 mg vaginal do dia 14 ao dia 28 do ciclo
• Ou progesterona oral 200-300 mg ao deitar (fase lútea)
• Clomifeno 50-100 mg dias 5-9 (estimula ovulação e forma corpo lúteo mais robusto)

Suporte nutricional:
• Vitamina C 1000 mg 2x/dia (aumenta progesterona em 50-75%)
• Vitamina E 400 UI/dia (melhora vascularização corpo lúteo)
• Magnésio 400 mg/dia (modula receptores de LH)
• Maca peruana 3000 mg/dia (melhora função ovariana)

FSH lúteo > 10 mUI/mL + FSH dia 3 > 10 mUI/mL:
• Considerar técnicas reprodução assistida (reserva ovariana baixa)
• DHEA 75 mg/dia + CoQ10 600 mg/dia (podem melhorar qualidade oocitária)

FSH < 1.0 mUI/mL:
• Investigar hipopituitarismo (RM sela túrcica, outros hormônios hipofisários)

Reavaliar FSH + progesterona no próximo ciclo',

    last_review_date = CURRENT_DATE
WHERE id = '0726b373-4b78-4cb8-91a9-0916681aef61';

-- 14. Magnésio RBC
UPDATE score_items
SET
    interpretation = 'Magnésio intracelular (RBC) reflete estoques teciduais melhor que magnésio sérico, que representa <1% do magnésio corporal total. Magnésio RBC é cofator em >300 reações enzimáticas, incluindo síntese de ATP, estabilização de DNA/RNA, função neuromuscular e regulação de canais iônicos (cálcio, potássio). Deficiência é comum em dietas ocidentais (alimentos processados, solo empobrecido), uso de diuréticos, inibidores de bomba de prótons, diabetes e alcoolismo. Magnésio baixo causa cãibras, arritmias, insônia, ansiedade e resistência à insulina.',

    clinical_relevance = 'Medicina funcional valoriza magnésio RBC em: síndrome metabólica e diabetes tipo 2 (magnésio melhora sensibilidade à insulina); hipertensão arterial (magnésio relaxa músculo liso vascular); arritmias cardíacas (previne taquicardia ventricular, fibrilação atrial); enxaqueca crônica (magnésio RBC <4.0 mg/dL presente em 50% dos casos); fibromialgia e fadiga crônica; síndrome das pernas inquietas; cãibras musculares recorrentes; ansiedade e insônia. Magnésio RBC <4.2 mg/dL já compromete função mitocondrial e síntese de ATP.',

    optimal_value_male_min = 5.0,
    optimal_value_male_max = 6.5,
    optimal_value_female_min = 5.0,
    optimal_value_female_max = 6.5,
    optimal_value_description = 'Magnésio RBC ótimo: 5.0-6.5 mg/dL (padrão MFI). Valores <4.0 mg/dL indicam deficiência severa, associada a risco cardiovascular e metabólico. Faixa 4.0-5.0 mg/dL é insuficiente para função enzimática ótima, especialmente em contextos de estresse (exercício intenso, doença aguda). Níveis >6.5 mg/dL são raros e geralmente refletem suplementação agressiva (>800 mg/dia) ou insuficiência renal. Ideal: 5.5-6.0 mg/dL para otimizar ATP mitocondrial, função neuromuscular e sensibilidade insulínica.',

    health_recommendations = 'Magnésio RBC < 5.0 mg/dL:
• Magnésio quelado (glicinato, treonato, taurato) 400-600 mg/dia (elemento) dividido em 2-3 doses
• Magnésio treonato 2000 mg/dia (atravessa barreira hematoencefálica, ideal para enxaqueca/ansiedade)
• Magnésio taurato 400 mg/dia (ideal para hipertensão, arritmias)
• Fontes alimentares: espinafre cozido (157 mg/xícara), amêndoas (80 mg/30g), abacate (58 mg), chocolate amargo (64 mg/30g)

Absorção otimizada:
• Tomar com refeição contendo gordura (aumenta absorção)
• Evitar cálcio simultâneo (compete por absorção)
• Vitamina D adequada (25-OH >40 ng/mL) melhora absorção de magnésio

Protocolos específicos:
• Enxaqueca: magnésio treonato 2000 mg/dia + riboflavina 400 mg/dia + CoQ10 300 mg/dia
• Hipertensão: magnésio taurato 400 mg/dia + potássio 3000 mg/dia (alimentar) + reduzir sódio <2300 mg/dia
• Diabetes: magnésio glicinato 400 mg/dia + cromo picolinato 200 mcg/dia + ácido alfa-lipóico 600 mg/dia
• Arritmias: magnésio taurato 600 mg/dia IV se agudo, oral manutenção
• Insônia: magnésio glicinato 400-600 mg 1h antes de dormir + glicina 3g

Magnésio RBC < 4.0 mg/dL:
• Considerar magnésio IV 2g em 500 mL SF 0.9% semanal por 4 semanas (reposição agressiva)
• Investigar causas: diabetes (poliúria), diuréticos, IBP, doença celíaca, alcoolismo

Reavaliar magnésio RBC em 12-16 semanas
Monitorar função renal (creatinina) se doses >800 mg/dia',

    last_review_date = CURRENT_DATE
WHERE id = 'c76becd2-8a0c-40b7-bb09-7ce24db94bb1';

-- 15. Sódio
UPDATE score_items
SET
    interpretation = 'Sódio sérico reflete balanço hídrico mais que status total de sódio corporal. Hiponatremia (<135 mEq/L) pode resultar de retenção hídrica (SIADH, insuficiência cardíaca, cirrose), perdas renais (diuréticos, insuficiência adrenal), perdas gastrointestinais ou polidipsia primária. Hipernatremia (>145 mEq/L) indica desidratação (perdas hídricas excedem perdas de sódio) ou diabetes insipidus. Sódio é fundamental para potencial de ação neuronal, contração muscular e regulação de volume extracelular. Distúrbios agudos podem causar encefalopatia, convulsões e coma.',

    clinical_relevance = 'Medicina funcional valoriza sódio em: avaliação de fadiga inexplicada e hipotensão (hiponatremia <130 mEq/L pode indicar insuficiência adrenal); investigação de SIADH (secreção inadequada de ADH) em contextos oncológicos ou neurológicos; monitoramento de atletas de endurance (hiponatremia dilucional por ingestão excessiva de água); diagnóstico diferencial de confusão mental em idosos (hiponatremia é causa comum de delirium). Sódio ideal: 138-142 mEq/L, refletindo euvolemia e função renal/endócrina normal.',

    optimal_value_male_min = 138.0,
    optimal_value_male_max = 142.0,
    optimal_value_female_min = 138.0,
    optimal_value_female_max = 142.0,
    optimal_value_description = 'Sódio ótimo: 138-142 mEq/L (padrão MFI). Valores <135 mEq/L indicam hiponatremia, exigindo investigação de causa (osmolalidade sérica/urinária, sódio urinário). Faixa 135-138 mEq/L é limite inferior, podendo associar-se a fadiga e hipotensão ortostática. Níveis >145 mEq/L indicam hipernatremia, geralmente por desidratação. Ideal: 139-141 mEq/L, refletindo equilíbrio hidroeletrolítico ótimo e função renal preservada.',

    health_recommendations = 'Sódio < 135 mEq/L (hiponatremia):
• URGENTE se <125 mEq/L (risco de edema cerebral, convulsões)
• Investigar causa:
  - SIADH: osmolalidade urinária >100 mOsm/kg + sódio urinário >20 mEq/L
  - Insuficiência adrenal: cortisol <10 mcg/dL, ACTH, teste estímulo
  - Diuréticos: revisar hidroclorotiazida, furosemida
  - Hipotireoidismo: TSH >5 mUI/mL

Hiponatremia leve-moderada (125-134 mEq/L):
• Restrição hídrica <1000-1500 mL/dia se SIADH
• Suplementação sal 3-5g/dia (1/2 colher chá) em refeições se perdas renais
• Fludrocortisona 0.1 mg/dia se insuficiência adrenal

Hiponatremia grave (<125 mEq/L):
• Internação hospitalar
• Salina hipertônica 3% 100 mL bolus (eleva sódio 2-3 mEq/L)
• Correção máxima: 6-8 mEq/L em 24h (evitar mielinólise pontina)

Sódio > 145 mEq/L (hipernatremia):
• Reidratação oral: água 2-3L/dia + eletrólitos
• Se diabetes insipidus (poliúria + osmolalidade urinária <300 mOsm/kg):
  - Desmopressina 0.1-0.2 mg nasal 2x/dia (DI central)
  - Hidroclorotiazida 25 mg/dia (paradoxalmente reduz poliúria no DI nefrogênico)

Prevenção hiponatremia em atletas:
• Bebidas isotônicas durante exercício >2h (não só água)
• Reposição sódio 500-700 mg/L líquido ingerido

Reavaliar sódio em 24-72h se anormal',

    last_review_date = CURRENT_DATE
WHERE id = '161dcbd1-6694-4175-958b-2b260ae48a40';

-- 16. Hematócrito - Mulheres
UPDATE score_items
SET
    interpretation = 'Hematócrito é percentual do volume sanguíneo ocupado por hemácias. Em mulheres, valores são ~3-4% menores que homens devido a perdas menstruais e efeitos do estrogênio. Hematócrito baixo (<36%) indica anemia, podendo resultar de deficiência de ferro, B12, folato, hemorragia crônica ou doença renal crônica (deficiência de eritropoetina). Hematócrito elevado (>48%) sugere policitemia, vista em desidratação, tabagismo, apneia do sono, DPOC ou policitemia vera. Hematócrito ideal: 39-44%, otimizando transporte de oxigênio sem aumentar viscosidade sanguínea.',

    clinical_relevance = 'Medicina funcional valoriza hematócrito em mulheres para: diagnóstico de anemia ferropriva (causa mais comum em mulheres pré-menopausais); avaliação de fadiga, dispneia e intolerância ao exercício; monitoramento de menstruações abundantes (menorragia) - hematócrito <36% sugere perda sanguínea significativa; rastreamento de desidratação crônica (hematócrito >46%); investigação de policitemia secundária em fumantes ou portadoras de apneia do sono. Hematócrito 36-39% ainda é subótimo, podendo causar fadiga e desempenho físico reduzido.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 39.0,
    optimal_value_female_max = 44.0,
    optimal_value_description = 'Hematócrito ótimo em mulheres: 39-44% (padrão MFI). Valores <36% indicam anemia, exigindo investigação (ferritina, B12, folato, função renal). Faixa 36-39% é limítrofe, associada a fadiga e redução de VO2 máx. Níveis >48% sugerem policitemia, aumentando risco de trombose (viscosidade sanguínea elevada). Ideal: 41-43%, otimizando capacidade de transporte de O₂ e desempenho físico sem comprometer fluxo sanguíneo.',

    health_recommendations = 'Hematócrito < 39%:
• Investigar anemia: hemoglobina, VCM, ferritina (<30 ng/mL deficiência), B12 (<400 pg/mL), folato
• Ferro quelado 100-200 mg/dia + vitamina C 500 mg em jejum
• Se VCM >100 fL: B12 metilcobalamina 1000 mcg sublingual/dia + folato 5-MTHF 800 mcg/dia
• Fontes alimentares: carne vermelha (ferro heme, alta absorção), lentilhas, espinafre

Hematócrito < 36% (anemia moderada):
• Avaliar perdas: menstruação (fluxo/duração), sangue oculto nas fezes (descartar sangramento GI)
• Se menorragia: ácido tranexâmico 1000 mg 3x/dia durante menstruação (reduz perda em 40%)
• Ferro IV se má absorção oral ou anemia grave (Hb <10 g/dL)

Protocolo anemia ferropriva:
• Ferro 200 mg/dia + vitamina C 500 mg + ácido fólico 1 mg
• Reavaliar hemograma em 8 semanas (esperar aumento Hb 1-2 g/dL)

Hematócrito > 46%:
• Descartar desidratação: reidratar 2-3L água/dia, reavaliar em 1 semana
• Se persistente: investigar policitemia secundária (gasometria, espirometria, polissonografia)
• Se policitemia vera (JAK2 V617F positivo): flebotomia 450 mL semanal até Hct <42%

Otimização performance:
• Meta atletas: 42-44% (maximiza VO2 máx)
• Ferro otimizado (ferritina 70-100 ng/mL)
• Altitudes >2500m estimulam eritropoiese natural (treino em altitude)

Reavaliar hematócrito em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = '32037968-f263-4e7d-ab85-ea83efd61c7b';

-- 17. Interleucina-6 (IL-6)
UPDATE score_items
SET
    interpretation = 'IL-6 é citocina pró-inflamatória produzida por macrófagos, células T e fibroblastos em resposta a infecções, trauma e estresse oxidativo. IL-6 induz síntese de proteínas de fase aguda (PCR, fibrinogênio) no fígado e estimula resposta imune. Elevação crônica de IL-6 caracteriza inflammaging (inflamação crônica de baixo grau associada ao envelhecimento) e está presente em obesidade, diabetes, aterosclerose, Alzheimer, sarcopenia e câncer. IL-6 >5 pg/mL correlaciona-se com maior risco cardiovascular, declínio cognitivo e fragilidade.',

    clinical_relevance = 'Medicina funcional valoriza IL-6 como: marcador de inflammaging e envelhecimento biológico acelerado; preditor de risco de doenças crônicas (diabetes, demência, câncer); indicador de resposta a intervenções anti-inflamatórias (dieta, exercício, sono); biomarcador em condições autoimunes e neurodegenerativas. IL-6 >3 pg/mL já indica ativação inflamatória crônica, mesmo sem doença manifesta. Redução de IL-6 com perda de peso, exercício e dieta anti-inflamatória correlaciona-se com melhora de outcomes metabólicos e cognitivos.',

    optimal_value_male_min = 0.5,
    optimal_value_male_max = 2.0,
    optimal_value_female_min = 0.5,
    optimal_value_female_max = 2.0,
    optimal_value_description = 'IL-6 ótima: 0.5-2.0 pg/mL (padrão MFI). Valores <0.5 pg/mL são raros e sem significado clínico negativo. Faixa 2-5 pg/mL indica inflamação crônica de baixo grau, associada a maior risco de síndrome metabólica. Níveis >5 pg/mL sugerem inflamação ativa (infecção, autoimunidade, trauma) ou inflammaging significativo. Valores >10 pg/mL requerem investigação de infecção aguda, doença autoimune ou neoplasia. Ideal: 0.8-1.5 pg/mL, refletindo homeostase imune sem ativação inflamatória crônica.',

    health_recommendations = 'IL-6 > 2.0 pg/mL:
• Investigar inflamação crônica: PCR ultrassensível (<1.0 mg/L ideal), VHS, hemograma
• Descartar infecção ativa, doença autoimune ou neoplasia se IL-6 >10 pg/mL

Protocolo anti-inflamatório:
• Ômega-3 EPA+DHA 2-3g/dia (reduz IL-6, TNF-α)
• Curcumina lipossomal 1000 mg/dia (inibe NF-κB, principal via IL-6)
• Resveratrol 500 mg/dia (ativa sirtuínas, reduz inflammaging)
• Vitamina D 4000 UI/dia (meta 25-OH 50-70 ng/mL) - modula resposta imune

Dieta anti-inflamatória (mediterrânea):
• Eliminar: açúcar, alimentos ultraprocessados, gorduras trans, excesso de ômega-6
• Priorizar: peixes gordos 3x/semana, azeite extra-virgem, vegetais coloridos, frutas vermelhas, nozes

Exercício anti-inflamatório:
• Aeróbico moderado 150 min/semana (reduz IL-6 em 20-30%)
• Evitar overtraining (exercício excessivo eleva IL-6)

Perda de peso (se IMC >25):
• 5-10% de redução de peso diminui IL-6 em 30-50%
• Jejum intermitente 16:8 (reduz citocinas pró-inflamatórias)

Sono:
• 7-9h/noite (privação de sono eleva IL-6)
• Tratar apneia do sono se presente (CPAP reduz IL-6)

Suplementos adicionais:
• N-acetilcisteína (NAC) 600 mg 2x/dia (precursor glutationa)
• Quercetina 500 mg 2x/dia (senolítico, elimina células senescentes produtoras de IL-6)
• Probióticos 50 bilhões UFC/dia (modula IL-6 intestinal)

IL-6 > 10 pg/mL:
• Investigação abrangente: TC tórax/abdome, marcadores tumorais, painel autoimune (FAN, anti-CCP)

Reavaliar IL-6 em 12-16 semanas',

    last_review_date = CURRENT_DATE
WHERE id = '053644b3-09b9-48cd-a31c-51ae7fe31131';

-- 18. FSH - Ovulação
UPDATE score_items
SET
    interpretation = 'FSH na ovulação (meio do ciclo, dia 14 em ciclo de 28 dias) apresenta pico transitório 24-36h antes da liberação do óvulo, geralmente atingindo 2-3x o nível basal folicular. Esse pico de FSH, sincronizado com pico de LH (mais pronunciado), é essencial para maturação final do oócito e ruptura folicular. FSH ovulatório muito baixo (<10 mUI/mL) pode indicar insuficiência hipofisária ou falha em resposta ao feedback positivo de estradiol. FSH excessivamente alto (>25 mUI/mL) é incomum na ovulação, sugerindo coleta em momento errado ou disfunção ovariana.',

    clinical_relevance = 'Medicina funcional valoriza FSH ovulatório em: confirmação de ovulação (pico FSH + pico LH >20 mUI/mL); avaliação de infertilidade com ciclos anovulatórios (ausência de pico de FSH/LH); monitoramento de indução de ovulação em tratamentos de fertilidade; investigação de síndrome do folículo luteinizado não roto (LUF - pico hormonal sem liberação do óvulo). FSH ovulatório <15 mUI/mL pode indicar pico subótimo, comprometendo fertilidade. Relação LH/FSH no pico >1.5 é típica da ovulação normal.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 10.0,
    optimal_value_female_max = 25.0,
    optimal_value_description = 'FSH ótimo na ovulação: 10-25 mUI/mL (padrão MFI). Valores <10 mUI/mL sugerem pico ovulatório subótimo ou coleta fora do momento ideal (dosar diariamente dias 12-16 para capturar pico). Faixa >25 mUI/mL é incomum e pode indicar disfunção ovariana ou coleta tardia (pós-ovulação). Ideal: 12-20 mUI/mL, com pico de LH simultâneo >20 mUI/mL (relação LH/FSH ~2:1), garantindo ruptura folicular e liberação oocitária efetiva.',

    health_recommendations = 'FSH ovulatório < 10 mUI/mL (pico ausente/insuficiente):
• Confirmar momento da coleta: usar kits de LH urinário (detecta pico 24h antes) ou dosagens seriadas FSH/LH dias 12-16
• Ultrassom transvaginal seriado para confirmar crescimento folicular (folículo dominante >18 mm)
• Progesterona lútea (dia 21): <5 ng/mL confirma anovulação

Indução de ovulação:
• Clomifeno 50-100 mg dias 5-9 do ciclo (estimula pico de FSH/LH)
• Letrozol 2.5-5 mg dias 3-7 (inibidor aromatase, mais efetivo em SOP)
• Gonadotrofinas exógenas (FSH recombinante) 75-150 UI/dia se refratária a clomifeno

Suporte natural ovulação:
• Inositol (mio-inositol 2g + d-chiro-inositol 50 mg) 2x/dia (melhora qualidade oocitária, especialmente SOP)
• Vitamina E 400 UI/dia + selênio 200 mcg/dia (antioxidantes foliculares)
• CoQ10 ubiquinol 300-600 mg/dia (melhora função mitocondrial oocitária)
• Ômega-3 2g/dia (melhora ambiente folicular)

FSH ovulatório > 25 mUI/mL:
• Repetir dosagem com timing correto (dia do pico de LH)
• Se persistentemente elevado: considerar diminuição de reserva ovariana (dosar FSH dia 3 + HAM)

Protocolo fertilidade:
• Rastrear ovulação com LH urinário diário (dias 10-16)
• Relações sexuais dia do pico LH e dia seguinte
• Progesterona vaginal 200 mg/dia do dia 15 ao dia 28 (suporte fase lútea)

Reavaliar no próximo ciclo com ultrassom seriado',

    last_review_date = CURRENT_DATE
WHERE id = '915bc591-b263-4fd2-a64d-4a04b38c97f7';

-- 19. Urocultura com Antibiograma
UPDATE score_items
SET
    interpretation = 'Urocultura identifica microorganismos em infecção do trato urinário (ITU), com contagem de colônias quantificando carga bacteriana. ≥100,000 UFC/mL (10⁵) de único patógeno em amostra de jato médio define bacteriúria significativa. Patógenos comuns: E. coli (80-85%), Klebsiella, Proteus, Enterococcus, Staphylococcus saprophyticus. Antibiograma automatizado testa sensibilidade a antibióticos (S=sensível, I=intermediário, R=resistente), guiando terapia direcionada. Culturas com <10,000 UFC/mL podem ser contaminação, exceto em punção suprapúbica ou cateterismo (qualquer crescimento é significativo).',

    clinical_relevance = 'Medicina funcional valoriza urocultura em: ITU recorrente (>3 episódios/ano em mulheres) - investigar fatores predisponentes (relação sexual, diafragma, hipoestrogenismo pós-menopausa, anormalidades urinárias); bacteriúria assintomática (não tratar exceto em gestantes, pré-cirurgia urológica); resistência antibiótica crescente (E. coli resistente a fluoroquinolonas >20% em muitas regiões); prevenção de pielonefrite e sepse. ITU complicada (homens, gestantes, imunossupressão, anormalidades estruturais) requer tratamento guiado por cultura.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Urocultura ótima: NEGATIVA ou <10,000 UFC/mL (sem crescimento bacteriano significativo). Crescimento ≥100,000 UFC/mL de patógeno único com sintomas (disúria, urgência, polaciúria) confirma ITU e requer tratamento. Contagem 10,000-100,000 UFC/mL é indeterminada, podendo ser ITU incipiente ou contaminação (repetir cultura se sintomática). Crescimento polimicrobiano (>3 organismos) sugere contaminação da coleta. Bacteriúria assintomática (≥100,000 UFC/mL sem sintomas) não requer tratamento em mulheres não-gestantes.',

    health_recommendations = 'Urocultura positiva (≥100,000 UFC/mL):
• Iniciar antibiótico conforme antibiograma:
  - Nitrofurantoína 100 mg 12/12h por 5-7 dias (primeira linha, não usar se TFG <30 mL/min)
  - Fosfomicina 3g dose única (alternativa, menos resistência)
  - Sulfametoxazol-trimetoprima 800/160 mg 12/12h por 3 dias (evitar se resistência local >20%)
  - Fluoroquinolonas (ciprofloxacino 500 mg 12/12h) reservar para ITU complicada

ITU complicada (homens, pielonefrite, sepse):
• Ceftriaxona 1-2g IV 24/24h + aminoglicosídeo (gentamicina 5 mg/kg/dia) até antibiograma
• Tratamento prolongado 10-14 dias

Prevenção ITU recorrente:
• D-manose 2g/dia (impede adesão E. coli ao urotélio, eficácia ~60%)
• Cranberry (proantocianidinas PAC-A 36 mg) 2x/dia (reduz recorrência em 30-40%)
• Probióticos vaginais (Lactobacillus rhamnosus, L. reuteri) - restauram microbiota
• Estrógeno vaginal 0.5 mg 2x/semana (pós-menopausa, reduz pH vaginal, restaura lactobacilos)

Higiene:
• Urinar após relações sexuais (reduz risco em 50%)
• Hidratação 2-3L água/dia (dilui bactérias)
• Evitar duchas vaginais, espermicidas

Bacteriúria assintomática:
• NÃO tratar exceto: gestantes, pré-cirurgia urológica, transplante renal
• Tratamento desnecessário aumenta resistência antibiótica

ITU recorrente refratária:
• Investigar anormalidades estruturais: ultrassom rins/vias urinárias, cistoscopia
• Profilaxia antibiótica: nitrofurantoína 50 mg ao deitar (long-term suppression)

Reavaliar cultura 1 semana após antibiótico se sintomas persistirem',

    last_review_date = CURRENT_DATE
WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';

-- 20. Muco - Sedimento Urinário
UPDATE score_items
SET
    interpretation = 'Muco no sedimento urinário é secreção de células epiteliais e glândulas das vias urinárias. Presença ocasional é normal, especialmente em mulheres (contaminação vaginal). Muco aumentado pode indicar irritação/inflamação do trato urinário (cistite, uretrite), contaminação vaginal (leucorreia) ou cateterismo prolongado. Muco associado a leucócitos, bactérias e eritrócitos sugere processo inflamatório/infeccioso ativo. Muco isolado em grande quantidade pode ser contaminação ou, raramente, produção excessiva em estados alérgicos ou irritativos.',

    clinical_relevance = 'Medicina funcional valoriza muco urinário em: diferenciação entre ITU e contaminação vaginal (muco + células epiteliais escamosas numerosas sugere contaminação); monitoramento de cistite intersticial/síndrome da bexiga dolorosa (muco pode estar presente); avaliação de uretrite não-gonocócica (Chlamydia, Ureaplasma); investigação de cálculos renais (irritação ureteral pode aumentar produção de muco). Muco isolado, sem leucócitos ou bactérias, geralmente não tem significado patológico.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Muco ótimo no sedimento: AUSENTE ou ESCASSO (traços ocasionais são normais, especialmente em mulheres). Muco MODERADO a ABUNDANTE isolado (sem leucócitos/bactérias) sugere contaminação vaginal - repetir EAS com coleta de jato médio cuidadosa, higienizando genitália. Muco ABUNDANTE + leucócitos + bactérias indica inflamação/infecção urinária. Muco ABUNDANTE + eritrócitos pode sugerir trauma, cálculo ou irritação mecânica.',

    health_recommendations = 'Muco MODERADO a ABUNDANTE isolado:
• Repetir EAS com coleta de jato médio adequada:
  - Higienizar genitália com água e sabão neutro
  - Desprezar primeiro jato, coletar jato médio
  - Mulheres: afastar lábios vaginais durante coleta
• Se persistente e assintomática: sem significado clínico, apenas contaminação

Muco + leucócitos + bactérias (infecção):
• Solicitar urocultura com antibiograma
• Tratamento antibiótico conforme cultura (ver protocolo urocultura)

Muco + eritrócitos (hematúria):
• Investigar: ultrassom rins/vias urinárias, cistoscopia
• Descartar cálculos, tumores, traumatismos

Muco recorrente com sintomas urinários (disúria, urgência):
• Investigar cistite intersticial: cistoscopia + biópsia, teste de sensibilidade potássio
• Protocolo cistite intersticial:
  - Eliminar alimentos irritantes: cafeína, álcool, picantes, cítricos, adoçantes artificiais
  - Pentosan polissulfato 100 mg 3x/dia (repara camada GAG da bexiga)
  - Amitriptilina 25-75 mg ao deitar (reduz dor neuropática)
  - Instilação vesical com DMSO, heparina, lidocaína

Muco + disúria em homens jovens (uretrite):
• Investigar ISTs: PCR Chlamydia, gonorreia, Ureaplasma, Mycoplasma
• Tratamento empírico se PCR não disponível: azitromicina 1g dose única + ceftriaxona 500 mg IM

Prevenção contaminação:
• Instrução adequada de coleta
• Primeira urina da manhã (maior concentração, menos contaminação)

Reavaliar EAS em 1 semana se sintomas persistirem',

    last_review_date = CURRENT_DATE
WHERE id = 'c0269b3c-2098-4f71-a999-d20fc4ce7cdd';

-- 21. Hepatite B - HbsAg
UPDATE score_items
SET
    interpretation = 'HBsAg (antígeno de superfície do vírus da hepatite B) é marcador de infecção ativa por HBV. Detectável 1-10 semanas após exposição, persiste em infecção aguda (clearance em 6 meses) ou crônica (persistência >6 meses). HBsAg positivo indica viremia e infectividade. Infecção crônica por HBV afeta ~250 milhões de pessoas globalmente, aumentando risco de cirrose e carcinoma hepatocelular em 100x. HBsAg negativo com anti-HBs positivo indica imunidade (vacinação ou infecção resolvida). Anti-HBc IgM positivo diferencia infecção aguda de crônica.',

    clinical_relevance = 'Medicina funcional valoriza HBsAg em: rastreamento universal de gestantes (transmissão vertical prevenível com imunoglobulina e vacina neonatal); triagem pré-vacinação em grupos de risco (profissionais saúde, parceiros de portadores); investigação de elevação de transaminases; monitoramento de portadores crônicos (carga viral, DNA-HBV, fibrose hepática). HBsAg positivo requer confirmação com HBeAg (infectividade alta se positivo), DNA-HBV quantitativo e avaliação de fibrose (elastografia/FibroScan). Meta: supressão viral para prevenir progressão para cirrose/CHC.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'HBsAg ótimo: NÃO REAGENTE (negativo), indicando ausência de infecção ativa por HBV. HBsAg REAGENTE (positivo) confirma infecção aguda ou crônica, exigindo avaliação adicional (anti-HBc IgM, HBeAg, DNA-HBV). Imunidade ideal: HBsAg negativo + anti-HBs >10 mUI/mL (proteção por vacinação ou infecção resolvida). Anti-HBs <10 mUI/mL após vacinação indica não-resposta (5-10% da população), necessitando dose de reforço ou esquema alternativo.',

    health_recommendations = 'HBsAg POSITIVO (infecção aguda ou crônica):
• Confirmar com testes adicionais:
  - Anti-HBc IgM: positivo em aguda, negativo em crônica
  - HBeAg: se positivo, alta infectividade (DNA-HBV >200,000 UI/mL)
  - DNA-HBV quantitativo (carga viral)
  - ALT, AST, bilirrubinas, albumina, TAP (função hepática)
  - FibroScan ou elastografia (avaliar fibrose - F0-F4)

Hepatite B crônica (HBsAg >6 meses):
• Indicações tratamento antiviral:
  - DNA-HBV >2000 UI/mL + ALT elevada (>2x limite superior)
  - Fibrose ≥F2 (FibroScan >7.0 kPa)
  - HBeAg positivo com DNA-HBV >20,000 UI/mL
  - Cirrose compensada (qualquer DNA-HBV detectável)

Tratamento antiviral:
• Tenofovir alafenamida (TAF) 25 mg/dia (primeira linha, menos toxicidade renal/óssea que TDF)
• Ou entecavir 0.5 mg/dia (alternativa)
• Duração: prolongada (anos), até soroconversão HBsAg→anti-HBs ou supressão viral persistente

Monitoramento portador crônico:
• DNA-HBV, ALT, AST a cada 3-6 meses
• Alfa-fetoproteína (AFP) + ultrassom abdome a cada 6 meses (rastreamento CHC)
• FibroScan anual

Prevenção transmissão:
• Parceiros sexuais: vacinar (3 doses 0-1-6 meses)
• Gestantes HBsAg+: imunoglobulina anti-HBV + vacina ao neonato nas primeiras 12h de vida (reduz transmissão >95%)
• Uso de preservativo, não compartilhar lâminas/escovas

Hepatoprotecão:
• Abstinência total de álcool (acelera fibrose)
• Silimarina 420 mg/dia (estabiliza hepatócitos)
• Ácido fólico 1 mg/dia + complexo B (apoio metilação hepática)
• Perda de peso se esteatose (MRI ou ultrassom)

HBsAg NEGATIVO + anti-HBs <10 mUI/mL (não-imune):
• Vacinar: Engerix-B 20 mcg IM (0-1-6 meses)
• Dosar anti-HBs 1-2 meses após 3ª dose
• Se anti-HBs <10 mUI/mL: repetir 3 doses (não-respondedor primário) ou dose dobrada (40 mcg)

Reavaliar HBsAg a cada 6-12 meses se crônico',

    last_review_date = CURRENT_DATE
WHERE id = 'eab8daae-3a2c-463b-bd03-d98434f27605';

-- 22. Proteínas Totais
UPDATE score_items
SET
    interpretation = 'Proteínas totais somam albumina (60%) e globulinas (40%). Refletem estado nutricional proteico, função hepática (síntese) e renal (perda). Hipoproteinemia (<6.0 g/dL) pode indicar desnutrição, má absorção, perda renal (síndrome nefrótica), perda gastrointestinal (enteropatia perdedora de proteína) ou insuficiência hepática. Hiperproteinemia (>8.5 g/dL) sugere desidratação, mieloma múltiplo (gamopatia monoclonal) ou inflamação crônica (globulinas elevadas). Fracionamento em albumina e globulinas (eletroforese) diferencia causas.',

    clinical_relevance = 'Medicina funcional valoriza proteínas totais em: avaliação de sarcopenia e desnutrição em idosos (proteínas <6.5 g/dL indicam depleção); investigação de edema e ascite (hipoalbuminemia <3.5 g/dL reduz pressão oncótica); monitoramento de doença hepática crônica (cirrose reduz albumina, eleva globulinas); diagnóstico de gamopatias (proteínas >9.0 g/dL requer eletroforese, imunofixação). Proteínas totais isoladas são inespecíficas; fracionamento (albumina, globulinas) é essencial para diagnóstico.',

    optimal_value_male_min = 6.8,
    optimal_value_male_max = 8.0,
    optimal_value_female_min = 6.8,
    optimal_value_female_max = 8.0,
    optimal_value_description = 'Proteínas totais ótimas: 6.8-8.0 g/dL (padrão MFI). Valores <6.5 g/dL sugerem hipoproteinemia (desnutrição, má absorção, perda renal/GI, insuficiência hepática). Faixa 6.5-6.8 g/dL é limítrofe, podendo refletir ingestão proteica inadequada ou inflamação crônica. Níveis >8.5 g/dL indicam hiperproteinemia (desidratação, mieloma, inflamação). Ideal: 7.2-7.6 g/dL, refletindo estado nutricional proteico adequado e síntese/perda equilibradas.',

    health_recommendations = 'Proteínas totais < 6.8 g/dL:
• Fracionar: albumina (normal >3.5 g/dL) + globulinas (normal 2.0-3.5 g/dL)
• Investigar hipoalbuminemia:
  - Função hepática: ALT, AST, bilirrubinas, TAP (cirrose reduz síntese)
  - Função renal: ureia, creatinina, proteinúria 24h (>3.5g/dia = síndrome nefrótica)
  - Desnutrição: história alimentar, peso, IMC, pré-albumina (<15 mg/dL deficiência aguda)

Protocolo reposição proteica:
• Aumentar ingestão: 1.2-1.6 g/kg/dia proteína (carne, ovos, leguminosas, whey)
• Whey protein isolado 30-40g/dia (aminoácidos essenciais)
• Glutamina 10g/dia (preserva massa muscular, reduz catabolismo)
• Leucina 3g 3x/dia (estimula síntese proteica muscular - mTOR)

Proteínas < 6.0 g/dL (grave):
• Descartar síndrome nefrótica: proteinúria >3.5g/24h + albumina <3.5 g/dL + edema + dislipidemia
• Tratar causa de base: corticoides se glomerulonefrite, bloqueador RAAS
• Enteropatia perdedora de proteína: alfa-1-antitripsina fecal elevada (Crohn, linfangiectasia)

Proteínas > 8.5 g/dL:
• Descartar desidratação: reidratar 2-3L água/dia, repetir
• Se persistente: eletroforese de proteínas (identificar pico monoclonal)
• Mieloma múltiplo: proteínas >9.0 g/dL + pico monoclonal + anemia + lesões líticas (radiografia esqueleto)

Gamopatia monoclonal:
• Imunofixação sérica e urinária (identificar tipo Ig)
• Proteinúria Bence-Jones (cadeias leves urinárias)
• Mielograma (plasmócitos >10% confirma mieloma)

Otimização nutricional:
• Meta: albumina >4.0 g/dL (associada a melhor mortalidade)
• Proteína alta qualidade: ovos, carne vermelha, aves, peixes, whey
• Vitamina C 1000 mg/dia (síntese de colágeno)
• Zinco 30 mg/dia (síntese proteica)

Reavaliar proteínas totais + fracionamento em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845';

-- 23. USG Próstata - Volume Prostático
UPDATE score_items
SET
    interpretation = 'Volume prostático é calculado por ultrassom transretal (USTR) ou suprapúbico usando fórmula elipsoide (comprimento × largura × altura × 0.52). Próstata normal em adultos jovens é 20-25 mL, aumentando com idade (hiperplasia prostática benigna - HPB). Volume >30 mL é considerado aumentado; >50 mL indica HPB moderada-grave, associada a maior risco de sintomas urinários (LUTS - lower urinary tract symptoms), retenção urinária aguda e necessidade de tratamento cirúrgico. Volume prostático correlaciona-se com PSA (densidade do PSA = PSA/volume).',

    clinical_relevance = 'Medicina funcional valoriza volume prostático em: diagnóstico e estratificação de HPB (sintomas correlacionam melhor com volume que PSA); cálculo de densidade do PSA (PSAD = PSA/volume) - PSAD >0.15 ng/mL/mL aumenta suspeita de câncer; decisão terapêutica em HPB (próstata >50 mL responde melhor a inibidores 5-alfa-redutase como finasterida/dutasterida); planejamento cirúrgico (ressecção transuretral em próstatas <80 mL, prostatectomia aberta se >80 mL). Volume também guia biópsia prostática (próstatas grandes requerem mais fragmentos).',

    optimal_value_male_min = 20.0,
    optimal_value_male_max = 30.0,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Volume prostático ótimo: 20-30 mL (padrão MFI, adultos 40-60 anos). Valores <20 mL são normais em jovens, mas incomuns após 50 anos. Faixa 30-50 mL indica HPB leve-moderada, geralmente com sintomas urinários iniciais (noctúria, jato fraco). Volume >50 mL define HPB moderada-grave, com maior risco de retenção urinária e necessidade de intervenção. Volumes >80-100 mL requerem cirurgia. Ideal: manter <35 mL após 50 anos, prevenindo progressão sintomática.',

    health_recommendations = 'Volume prostático 30-50 mL (HPB leve-moderada):
• Alfa-bloqueadores (relaxam músculo liso prostático/vesical):
  - Tansulosina 0.4 mg/dia ao deitar (seletivo, menos hipotensão)
  - Alfuzosina 10 mg/dia (liberação prolongada)
  - Doxazosina 4-8 mg/dia (menos seletivo, mais hipotensão)

Fitoterapia:
• Saw palmetto (Serenoa repens) 320 mg/dia (inibe 5-alfa-redutase, reduz volume ~6%)
• Beta-sitosterol 60-130 mg/dia (melhora fluxo urinário)
• Pygeum africanum 100 mg 2x/dia (anti-inflamatório prostático)

Volume prostático > 50 mL (HPB moderada-grave):
• Inibidores 5-alfa-redutase (reduzem DHT, volume prostático -20-30%):
  - Finasterida 5 mg/dia (inibe tipo 2, reduz PSA 50%)
  - Dutasterida 0.5 mg/dia (inibe tipo 1 e 2, maior redução volume)
  - Efeito após 6-12 meses, uso contínuo necessário

Terapia combinada (próstata >50 mL + IPSS >15):
• Alfa-bloqueador + inibidor 5-alfa-redutase (superior a monoterapia)
• Exemplo: tansulosina 0.4 mg + dutasterida 0.5 mg/dia

Volume > 80 mL ou refratário a medicamentos:
• Ressecção transuretral da próstata (RTUP) - padrão-ouro
• Laser Holmium (HoLEP) - menos sangramento, internação curta
• Embolização arterial prostática (menos invasivo)

Prevenção progressão HPB:
• Licopeno 10-30 mg/dia (antioxidante, reduz crescimento prostático)
• Zinco 30 mg/dia + selênio 200 mcg/dia (modulam 5-alfa-redutase)
• Reduzir álcool e cafeína (irritantes vesicais)
• Exercício regular (30 min/dia reduz risco HPB em 25%)
• Evitar descongestionantes sistêmicos (pseudoefedrina piora retenção)

Monitoramento:
• IPSS (International Prostate Symptom Score) trimestral (meta <8 leve)
• Urofluxometria anual (fluxo máximo >15 mL/s normal)
• Resíduo pós-miccional <50 mL (ultrassom)
• PSA anual (finasterida/dutasterida reduzem PSA 50% - ajustar interpretação)

Reavaliar volume prostático (USG) em 12-24 meses',

    last_review_date = CURRENT_DATE
WHERE id = '23b012f9-ce8b-4a1d-95f4-cfe9183295e0';

-- 24. USG Próstata - Densidade PSA (PSAD)
UPDATE score_items
SET
    interpretation = 'Densidade do PSA (PSAD) é razão entre PSA sérico total e volume prostático (PSAD = PSA ng/mL / volume mL). PSAD diferencia HPB (densidade baixa) de câncer de próstata (densidade alta), pois HPB aumenta volume sem elevar muito PSA, enquanto câncer produz muito PSA em pequeno volume. PSAD >0.15 ng/mL/mL aumenta risco de câncer e é indicação de biópsia prostática, mesmo com PSA <4.0 ng/mL. PSAD <0.10 ng/mL/mL sugere HPB como causa de PSA elevado, podendo adiar biópsia em casos selecionados.',

    clinical_relevance = 'Medicina funcional valoriza PSAD em: refinar indicação de biópsia prostática (PSAD >0.15 tem VPP ~30-50% para câncer, enquanto PSAD <0.10 tem VPN >90%); estratificar pacientes com PSA zona cinza (4-10 ng/mL) - PSAD evita biópsias desnecessárias; monitoramento de vigilância ativa em câncer de baixo risco (PSAD >0.20 sugere upgrade patológico); decisão em ressonância multiparamétrica (RM próstata) - PSAD >0.15 + PI-RADS ≥3 indica biópsia. PSAD é superior ao PSA isolado para predição de câncer clinicamente significativo.',

    optimal_value_male_min = 0.00,
    optimal_value_male_max = 0.10,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'PSAD ótima: <0.10 ng/mL/mL (padrão MFI). Valores 0.10-0.15 ng/mL/mL são intermediários, sugerindo HPB ou prostatite, mas exigindo acompanhamento rigoroso (PSA a cada 6-12 meses, toque retal, considerar RM próstata). PSAD >0.15 ng/mL/mL indica risco elevado de câncer prostático (biópsia recomendada). PSAD >0.20 ng/mL/mL tem risco muito alto (>50% probabilidade câncer). Ideal: manter <0.08 ng/mL/mL, refletindo próstata normal sem neoformação.',

    health_recommendations = 'PSAD 0.10-0.15 ng/mL/mL (risco intermediário):
• Repetir PSA total, PSA livre/total em 3-6 meses
• Toque retal (nódulo palpável indica biópsia independente de PSAD)
• Considerar RM multiparamétrica da próstata (PI-RADS):
  - PI-RADS 1-2: acompanhar com PSA semestral
  - PI-RADS ≥3: biópsia dirigida por fusão RM-US

PSAD > 0.15 ng/mL/mL (indicação biópsia):
• Biópsia transretal guiada por ultrassom (TRUS) com 12-14 fragmentos
• Ou biópsia por fusão RM-ultrassom (maior detecção de câncer significativo)
• Avaliar relação PSA livre/total (<15% aumenta risco câncer)

PSAD > 0.20 ng/mL/mL (risco muito alto):
• Biópsia obrigatória
• Estadiamento se câncer confirmado: Gleason score, TNM, cintilografia óssea, TC/RM

Se biópsia negativa mas PSAD >0.15:
• Repetir biópsia em 12 meses (risco câncer anterior não detectado ~20%)
• RM próstata anual

PSAD < 0.10 ng/mL/mL com PSA 4-10 ng/mL:
• Provável HPB, não câncer
• Tratamento HPB (ver protocolo volume prostático)
• PSA + toque retal anual

Prevenção câncer prostático:
• Licopeno 10-30 mg/dia (reduz risco câncer em 20%, concentra-se na próstata)
• Selênio 200 mcg/dia + vitamina E 400 UI/dia (controverso após SELECT trial, benefício em deficientes)
• Chá verde 3 xícaras/dia (catequinas inibem 5-alfa-redutase e crescimento tumoral)
• Reduzir consumo laticínios ricos em cálcio (>1200 mg/dia aumenta risco)
• Ômega-3 2g/dia (reduz inflamação prostática)

Vigilância ativa (câncer Gleason 6, T1c-T2a, PSA <10):
• Critério: PSAD <0.15 ng/mL/mL
• Se PSAD aumenta >0.20: considerar tratamento definitivo (prostatectomia, radioterapia)

Reavaliar PSAD (PSA + volume por USG) em 6-12 meses',

    last_review_date = CURRENT_DATE
WHERE id = '317acc85-3ce9-4f97-8e14-799354166f5e';

-- 25. TC Tórax - Maior Nódulo Sólido
UPDATE score_items
SET
    interpretation = 'Nódulos pulmonares sólidos são opacidades arredondadas <3 cm, detectadas em TC de tórax. Tamanho é principal preditor de malignidade: nódulos <6 mm têm <1% risco de câncer; 6-8 mm ~2-5%; >8 mm ~15%; >20 mm ~50%. Outras características suspeitas: margens espiculadas, crescimento em TCs seriadas (tempo de duplicação <400 dias), SUV alto em PET-TC (>2.5), localização em lobo superior. Nódulos benignos comuns: granulomas (calcificação central/pipoca), hamartomas (calcificação popcorn/fat), cicatrizes fibróticas.',

    clinical_relevance = 'Medicina funcional valoriza manejo de nódulos por: estratificação de risco com calculadoras validadas (Brock, Mayo, Herder models); protocolo Lung-RADS para seguimento padronizado; decisão entre vigilância (nódulos <6 mm), TC seriadas (6-20 mm) ou biópsia/cirurgia (>20 mm ou crescimento); rastreamento em grupos de risco (fumantes/ex-fumantes 55-80 anos, carga tabágica >30 maços-ano). Nódulos >8 mm com características suspeitas requerem PET-TC ou biópsia (TTNA, broncoscopia). Adenocarcinoma pulmonar é tipo mais comum (40% dos nódulos malignos).',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 4.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 4.0,
    optimal_value_description = 'Nódulo pulmonar ótimo: AUSENTE ou <4 mm (considerado negativo, sem seguimento necessário). Faixa 4-6 mm: Lung-RADS 2 (seguimento opcional ou 12 meses se alto risco). Nódulos 6-8 mm: Lung-RADS 3 (TC controle 6-12 meses). Nódulos >8 mm: Lung-RADS 4 (PET-TC, biópsia ou ressecção). Nódulos >20 mm têm >50% risco malignidade, exigindo abordagem cirúrgica urgente. Ideal: ausência de nódulos ou nódulos calcificados benignos.',

    health_recommendations = 'Nódulo 4-6 mm (baixo risco):
• TC tórax de controle em 12 meses (estável = benigno)
• Se novo nódulo ou múltiplos: repetir em 6 meses
• Sem necessidade de outras investigações se estável

Nódulo 6-8 mm (risco intermediário):
• TC tórax de controle em 6-12 meses
• Se crescimento (aumento de 2 mm ou volume >25%): PET-TC ou biópsia
• Calcular risco com modelo Brock/Mayo (considera idade, tabagismo, tamanho, margem, localização)

Nódulo 8-20 mm (risco intermediário-alto):
• PET-TC (SUV >2.5 sugere malignidade)
• Se PET+ ou características suspeitas: biópsia transtorácica (TTNA) ou broncoscopia
• Ou TC controle em 3 meses (avaliar crescimento)

Nódulo > 20 mm (alto risco ~50% malignidade):
• Biópsia ou ressecção cirúrgica sem demora
• Videotoracoscopia (VATS) com biópsia de congelação + lobectomia se maligno
• Estadiamento completo: TC tórax/abdome/pelve, PET-TC, RM crânio

Características suspeitas (qualquer tamanho):
• Margens espiculadas, sinal de halo, derrame pleural
• PET-TC + biópsia mesmo se <8 mm

Prevenção câncer de pulmão:
• Cessação tabagismo (reduz risco 50% após 10 anos)
• Rastreamento TC baixa dose anual (fumantes 55-80 anos, >30 maços-ano)
• N-acetilcisteína (NAC) 600 mg 2x/dia (antioxidante pulmonar, controverso)
• Vitamina D 4000 UI/dia (meta 25-OH >50 ng/mL, reduz risco)
• Curcumina 1000 mg/dia (anti-inflamatório, pró-apoptótico)
• Licopeno 10-30 mg/dia (antioxidante)

Nódulo benigno confirmado (calcificação típica, estável >2 anos):
• Sem seguimento necessário
• TC apenas se novos sintomas (hemoptise, dor torácica)

Reavaliar conforme protocolo Lung-RADS
Biópsia se crescimento ou SUV >2.5',

    last_review_date = CURRENT_DATE
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';

-- 26. Endoscopia Alta - Esofagite (LA Classification)
UPDATE score_items
SET
    interpretation = 'Classificação de Los Angeles (LA) estratifica esofagite de refluxo endoscopicamente: Grau A (erosões <5 mm, confinadas a pregas mucosas), Grau B (erosões >5 mm em pregas, não confluentes), Grau C (erosões confluentes <75% da circunferência), Grau D (erosões ≥75% da circunferência). Gravidade correlaciona-se com exposição ácida esofágica e risco de complicações (estenose, Barrett). Esofagite Grau C-D indica DRGE grave, exigindo IBP em dose alta e investigação de hérnia hiatal.',

    clinical_relevance = 'Medicina funcional valoriza LA em: diagnóstico de DRGE erosiva (40% dos pacientes com sintomas têm endoscopia normal - DRGE não-erosiva); decisão terapêutica (Grau A-B: IBP dose padrão; Grau C-D: dose dobrada); rastreamento de Barrett (esofagite crônica metaplasia colunar); monitoramento de cicatrização (repetir endoscopia em 8-12 semanas pós-tratamento). Esofagite refratária a IBP requer pHmetria 24h para confirmar supressão ácida inadequada ou refluxo não-ácido.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'LA ótima: AUSÊNCIA DE ESOFAGITE (mucosa esofágica íntegra). LA Grau A-B: esofagite leve-moderada, geralmente responsiva a IBP padrão. LA Grau C-D: esofagite grave, alto risco de complicações (estenose 5-10%, Barrett 10-15%). Objetivo: cicatrização completa (mucosa normal) após 8-12 semanas de IBP. Persistência de erosões indica refratariedade ou não-aderência.',

    health_recommendations = 'LA Grau A-B (leve-moderada):
• Omeprazol 20 mg ou esomeprazol 40 mg em jejum por 8 semanas
• Medidas comportamentais:
  - Elevar cabeceira 15-20 cm
  - Evitar refeições 3h antes de deitar
  - Perda de peso se IMC >25 (reduz pressão intra-abdominal)
  - Eliminar: chocolate, hortelã, álcool, cafeína, alimentos gordurosos

LA Grau C-D (grave):
• Esomeprazol 40 mg 2x/dia ou pantoprazol 40 mg 2x/dia por 12 semanas
• Adicionar pró-cinético: domperidona 10 mg 3x/dia antes das refeições
• Endoscopia controle em 8-12 semanas (confirmar cicatrização)

Esofagite refratária (persistência após 8 semanas IBP):
• pHmetria 24h com impedância (avaliar supressão ácida)
• Se pH anormal: aumentar IBP ou trocar por rabeprazol 20 mg 2x/dia
• Se pH normal: considerar refluxo fracamente ácido/não-ácido (baclofen 10 mg 3x/dia)

Cirurgia anti-refluxo (fundoplicatura):
• Indicações: esofagite grave recorrente, refratariedade a IBP, intolerância/desejo de descontinuar IBP
• Fundoplicatura de Nissen (360°) ou Toupet (270°)

Manutenção após cicatrização:
• IBP mínima dose efetiva (omeprazol 10-20 mg/dia)
• Ou terapia sob-demanda (IBP apenas com sintomas)
• Probióticos 50 bilhões UFC/dia (contrabalançar disbiose por IBP)

Redução de IBP longo prazo (riscos: deficiência B12, magnésio, fraturas, infecções):
• Substituir gradualmente por antiácidos (ranitidina 150 mg 2x/dia, famotidina 20 mg 2x/dia)
• Alginato de sódio 10 mL após refeições (barreira mecânica)
• Zinco-L-carnosina 75 mg 2x/dia (cicatrização mucosa)
• Glutamina 5g 2x/dia (reparo epitelial)

Reavaliar endoscopia em 8-12 semanas',

    last_review_date = CURRENT_DATE
WHERE id = '4f6e007b-dcc2-4e51-aaad-b7359717f297';

-- 27. Endoscopia Alta - Barrett Esophagus (Prague C)
UPDATE score_items
SET
    interpretation = 'Classificação de Praga para Barrett: C (circunferencial, extensão circunferencial de metaplasia colunar acima da junção esôfago-gástrica) e M (máximo, maior extensão incluindo línguas). Exemplo: C3M5 = 3 cm circunferencial + línguas até 5 cm. Barrett é metaplasia intestinal (células caliciformes) em esôfago distal por refluxo crônico, aumentando risco de adenocarcinoma esofágico (0.5%/ano em Barrett não-displásico, >6%/ano em displasia de alto grau). Prague C >3 cm tem maior risco neoplásico.',

    clinical_relevance = 'Medicina funcional valoriza Prague C em: estratificação de risco de câncer (C ≥3 cm requer vigilância intensiva); protocolo de Seattle para biópsia (4 quadrantes a cada 2 cm + lesões visíveis); decisão entre vigilância (não-displásico, displasia baixo grau) e ablação (displasia alto grau). Barrett sem displasia: endoscopia a cada 3-5 anos; com displasia baixo grau: 6-12 meses; alto grau: ablação por radiofrequência ou ressecção endoscópica. Supressão ácida com IBP não regride Barrett, mas previne progressão.',

    optimal_value_male_min = 0.0,
    optimal_value_male_max = 0.0,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 0.0,
    optimal_value_description = 'Prague C ótimo: C0M0 (AUSÊNCIA DE BARRETT). Qualquer extensão de Barrett (C >0 ou M >0) requer vigilância endoscópica. C <1 cm (Barrett curto): risco menor, vigilância 3-5 anos. C 1-3 cm (Barrett médio): risco intermediário, vigilância 2-3 anos. C >3 cm (Barrett longo): risco alto, vigilância anual ou 6/6 meses se displasia. Objetivo: prevenir progressão displasia→adenocarcinoma através de supressão ácida rigorosa e ablação precoce de displasia.',

    health_recommendations = 'Barrett sem displasia (C <3 cm):
• Esomeprazol 40 mg/dia (supressão ácida otimizada)
• pHmetria 24h: confirmar pH esofágico <4 por <4% do tempo (meta)
• Endoscopia vigilância a cada 3-5 anos com protocolo Seattle (biópsias 4 quadrantes a cada 2 cm)
• Medidas anti-refluxo rigorosas (ver protocolo esofagite)

Barrett C ≥3 cm (Barrett longo):
• Endoscopia anual com biópsias extensivas
• Considerar cromoscopia (NBI, FICE) para identificar displasia
• Esomeprazol 40 mg 2x/dia

Barrett com displasia baixo grau:
• Confirmar displasia com 2º patologista experiente (alta variabilidade interobservador)
• Opções:
  1. Vigilância intensiva: endoscopia a cada 6 meses por 1 ano, depois anual
  2. Ablação por radiofrequência (RFA): erradica Barrett em 70-90%, reduz progressão câncer >80%
• Esomeprazol 40 mg 2x/dia (essencial pré e pós-ablação)

Barrett com displasia alto grau ou carcinoma intramucoso:
• Ressecção endoscópica mucosa (EMR) de lesões visíveis
• Ablação por radiofrequência do Barrett residual
• Ou esofagectomia se invasão submucosa

Pós-ablação:
• IBP dose alta permanentemente (esomeprazol 40 mg 2x/dia)
• Endoscopia vigilância: 3, 6, 12 meses, depois anual
• Recorrência Barrett em 10-15% (reablação se necessário)

Quimioprevenção (controversa):
• Aspirina 81 mg/dia (reduz risco adenocarcinoma ~30% em estudos observacionais)
• Estatina (atorvastatina 20 mg/dia, efeito anti-inflamatório)
• Metformina 500 mg 2x/dia (se diabético ou resistência insulínica)

Otimização estilo de vida:
• Perda de peso agressiva se IMC >25 (reduz pressão intra-abdominal)
• Eliminar álcool e tabaco (dobram risco de progressão)
• Dieta rica em antioxidantes: vegetais crucíferos, frutas vermelhas, chá verde

Reavaliar endoscopia conforme protocolo de vigilância',

    last_review_date = CURRENT_DATE
WHERE id = '66a4571d-f9e2-4f94-96cf-15145ef62499';

-- 28. Fundoscopia - Retinopatia Diabética
UPDATE score_items
SET
    interpretation = 'Retinopatia diabética (RD) é microangiopatia retiniana progressiva, classificada em: não-proliferativa leve (microaneurismas, hemorragias puntiformes), moderada (exsudatos algodonosos, hemorragias/microaneurismas moderados), grave (hemorragias extensas em 4 quadrantes, ou rosários venosos em 2 quadrantes, ou AMIR em 1 quadrante - regra 4-2-1); proliferativa (neovascularização de disco/retina, hemorragia vítrea). Edema macular diabético (EMD) pode ocorrer em qualquer estágio, comprometendo visão central. RD é principal causa de cegueira em adultos 20-74 anos.',

    clinical_relevance = 'Medicina funcional valoriza RD em: rastreamento anual em diabéticos (iniciar no diagnóstico em tipo 2, 5 anos após em tipo 1); controle glicêmico rigoroso (HbA1c <7% reduz progressão RD em 76% - DCCT trial); manejo de pressão arterial (meta <130/80 mmHg) e dislipidemia; fotocoagulação a laser em RD proliferativa ou EMD; anti-VEGF intravítreo (ranibizumab, aflibercept) para EMD com perda visual. Melhora rápida de HbA1c (>2% em 3-6 meses) pode paradoxalmente piorar RD temporariamente.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'Fundoscopia ótima: AUSÊNCIA DE RETINOPATIA (retina normal sem microaneurismas, hemorragias ou exsudatos). RD não-proliferativa leve: vigilância anual. RD moderada: vigilância 6-12 meses. RD grave: vigilância 3-4 meses ou fotocoagulação panretiniana (PRP). RD proliferativa: PRP urgente (risco hemorragia vítrea/descolamento retina). EMD: OCT macular + anti-VEGF se espessura central >250 microns. Meta: prevenir progressão para proliferativa através de controle glicêmico/PA rigoroso.',

    health_recommendations = 'RD ausente (prevenção):
• Controle glicêmico rigoroso: HbA1c <7.0% (<6.5% se jovem sem hipoglicemia)
• Controle pressão arterial: meta <130/80 mmHg
• Controle lipídico: LDL <70 mg/dL, TG <150 mg/dL
• Fundoscopia anual com retinografia

RD não-proliferativa leve-moderada:
• Otimizar HbA1c <6.5% (reduzir progressão gradualmente, <2% em 6 meses)
• Fenofibrato 160 mg/dia (reduz progressão RD em 30-40% mesmo sem dislipidemia - FIELD/ACCORD trials)
• Ômega-3 2-3g/dia (reduz inflamação retiniana)
• Fundoscopia a cada 6-12 meses

RD grave não-proliferativa:
• Fotocoagulação panretiniana (PRP) preventiva ou vigilância 3-4 meses
• Anti-VEGF intravítreo (alternativa ao laser, menos perda campo visual)
• HbA1c <7%, PA <130/80 mmHg

RD proliferativa:
• PRP urgente (1500-2000 queimaduras laser, 2-3 sessões)
• Se hemorragia vítrea: vitrectomia pars plana
• Anti-VEGF adjuvante (ranibizumab 0.5 mg intravítreo mensal)

Edema macular diabético (EMD):
• OCT macular (espessura central >250 microns indica tratamento)
• Anti-VEGF intravítreo mensal por 3-6 meses, depois PRN (treat-and-extend)
• Laser focal/grid se EMD sem centro foveal
• Implante intravítreo dexametasona (Ozurdex) se refratário a anti-VEGF

Suplementação neuroprotetora:
• Luteína 10 mg + zeaxantina 2 mg/dia (concentram-se na mácula, filtram luz azul)
• Vitamina D 4000 UI/dia (meta 25-OH 50-70 ng/mL, reduz neovascularização)
• Ácido alfa-lipóico 600 mg/dia (antioxidante, previne peroxidação lipídica retiniana)
• Benfotiamina 300 mg/dia (forma lipossolúvel B1, bloqueia vias AGE)

Gestantes diabéticas:
• Fundoscopia trimestral (gravidez acelera RD)
• PRP se RD proliferativa, diferir anti-VEGF se possível

Reavaliar fundoscopia conforme gravidade
Meta: reverter ou estabilizar RD',

    last_review_date = CURRENT_DATE
WHERE id = 'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52';

-- 29. USG Transvaginal - Espessura Endometrial Pós-Menopausa
UPDATE score_items
SET
    interpretation = 'Espessura endometrial pós-menopausa normal é <5 mm (atrofia fisiológica por hipoestrogenismo). Espessura >5 mm em mulher com sangramento pós-menopausa requer investigação (risco hiperplasia/câncer endometrial). Espessura >11 mm mesmo sem sangramento tem risco aumentado. Causas de espessamento: hiperplasia endometrial (simples/complexa, com/sem atipia), pólipo endometrial, câncer endometrial, tamoxifeno, terapia hormonal. Histeroscopia com biópsia é padrão-ouro para diagnóstico histológico.',

    clinical_relevance = 'Medicina funcional valoriza espessura endometrial pós-menopausa em: investigação de sangramento (qualquer sangramento pós-menopausa é câncer até prova em contrário); rastreamento em usuárias de tamoxifeno (risco câncer 2-3x); monitoramento de terapia hormonal (estrogênio sem progesterona espessa endométrio); diagnóstico de hiperplasia com atipia (lesão pré-neoplásica, risco câncer 30% se não tratada). Endométrio <4 mm tem VPN >99% para câncer em mulher com sangramento.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 0.0,
    optimal_value_female_max = 5.0,
    optimal_value_description = 'Espessura endometrial pós-menopausa ótima: <5 mm (linha endometrial fina, atrófica). Valores 5-11 mm: zona cinza, biópsia se sangramento presente. Espessura >11 mm: alto risco, biópsia obrigatória mesmo sem sangramento. Usuárias de tamoxifeno: espessamento até 8 mm é comum (efeito estrogênico parcial), biópsia se >8 mm ou sangramento. Terapia hormonal (estrogênio+progesterona): <5 mm esperado; >5 mm requer ajuste progesterona.',

    health_recommendations = 'Espessura 5-11 mm COM sangramento:
• Histeroscopia com biópsia endometrial (padrão-ouro)
• Ou curetagem uterina se histeroscopia indisponível
• Histopatologia: descartar hiperplasia com atipia, câncer endometrial

Espessura >11 mm (com ou sem sangramento):
• Biópsia endometrial obrigatória
• Se câncer confirmado: estadiamento (RM pelve, TC tórax-abdome, CA-125)
• Histerectomia total com salpingo-ooforectomia bilateral (tratamento padrão)

Hiperplasia sem atipia:
• Progestágeno cíclico ou contínuo:
  - Acetato de medroxiprogesterona 10 mg/dia por 14 dias/mês (cíclico)
  - Ou progesterona micronizada 200 mg/dia contínuo
• DIU liberador de levonorgestrel (Mirena) - opção
• Biópsia controle em 3-6 meses (regressão em 80-90%)

Hiperplasia com atipia (lesão pré-maligna):
• Histerectomia total (padrão-ouro, risco câncer oculto 30%)
• Ou progestágeno alto dose se desejo fertilidade: megestrol 160 mg/dia por 6 meses + biópsia trimestral

Prevenção câncer endometrial:
• Terapia hormonal balanceada: estrogênio sempre com progesterona (se útero intacto)
• Perda de peso se IMC >30 (obesidade aumenta estrogênio periférico, risco câncer 2-4x)
• Metformina 500 mg 2x/dia se resistência insulínica (reduz risco câncer endometrial ~30%)
• Exercício 150 min/semana (reduz risco)

Tamoxifeno:
• Ultrassom transvaginal anual (espessura >8 mm ou sangramento → biópsia)
• Considerar substituir por inibidor aromatase se pós-menopausa (anastrozol, letrozol - sem risco endometrial)

Espessura <5 mm SEM sangramento:
• Tranquilizar paciente, sem necessidade investigação
• Reavaliar apenas se sangramento surgir

Reavaliar ultrassom transvaginal em 6-12 meses se espessamento tratado',

    last_review_date = CURRENT_DATE
WHERE id = '30af9809-7316-47e6-b363-8279c7bd3a69';

-- 30. USG Transvaginal - Volume Ovariano
UPDATE score_items
SET
    interpretation = 'Volume ovariano é calculado por fórmula elipsoide (comprimento × largura × altura × 0.523). Ovário normal em menacme: 3-10 mL (média 6 mL). Após menopausa, ovários atrofiam para 1.5-3 mL. Volume >10 mL em menacme ou >5 mL pós-menopausa é anormal, sugerindo síndrome dos ovários policísticos (SOP - volume >10 mL + ≥12 folículos antrais 2-9 mm), cistoadenoma, tecoma ou malignidade. Volume <3 mL em menacme pode indicar insuficiência ovariana prematura (IOP). Volume ovariano correlaciona-se com reserva ovariana.',

    clinical_relevance = 'Medicina funcional valoriza volume ovariano em: diagnóstico de SOP (Rotterdam criteria: 2 de 3 - hiperandrogenismo, oligo-anovulação, ovários policísticos >10 mL); avaliação de massa anexial (volume >20 mL + CA-125 elevado sugere malignidade); monitoramento de insuficiência ovariana (volume <3 mL + FSH >25 mUI/mL); protocolo fertilidade (volume 7-10 mL com 10-20 folículos antrais indica boa reserva). Índice ROMA (CA-125 + HE4 + status menopausal + volume ovariano) estratifica risco de câncer.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 4.0,
    optimal_value_female_max = 10.0,
    optimal_value_description = 'Volume ovariano ótimo em menacme: 4-10 mL (média 6-8 mL). Volumes <3 mL sugerem insuficiência ovariana (dosar FSH, HAM, estradiol). Volume >10 mL indica ovários policísticos (considerar SOP se critérios clínicos presentes). Pós-menopausa: 1.5-3 mL (atrofia fisiológica). Volume >5 mL pós-menopausa requer investigação de massa ovariana (CA-125, doppler, RM pelve). Ideal menacme: 6-8 mL, refletindo reserva ovariana adequada e foliculogênese normal.',

    health_recommendations = 'Volume ovariano > 10 mL em menacme (SOP suspeita):
• Critérios Rotterdam (2 de 3): hiperandrogenismo clínico/laboratorial, oligo/anovulação, USG policístico
• Investigação: testosterona total (>60 ng/dL), androstenediona, DHEA-S, LH/FSH (>2-3), glicemia jejum, insulina, HOMA-IR
• Tratamento SOP:
  - Perda de peso 5-10% se IMC >25 (restaura ovulação em 50%)
  - Metformina 500 mg 3x/dia ou 850 mg 2x/dia (melhora sensibilidade insulínica, regulariza ciclos)
  - Inositol (mio-inositol 2g + d-chiro-inositol 50 mg) 2x/dia
  - Espironolactona 100 mg/dia (anti-androgênico, reduz hirsutismo)
  - Anticoncepcional oral (etinilestradiol 30 mcg + drospirenona 3 mg) se não deseja gravidez

Volume ovariano < 3 mL (insuficiência ovariana):
• Dosar FSH (>25 mUI/mL em 2 ocasiões confirma IOP), HAM (<0.5 ng/mL), estradiol (<50 pg/mL)
• Investigar causas: cariótipo (síndrome Turner 45,X0), FMR1 pré-mutação (X frágil), autoimunidade (anti-ovário, anti-21-hidroxilase)
• Terapia hormonal: estrogênio + progesterona até idade menopausa natural (previne osteoporose, doença cardiovascular)
• DHEA 75 mg/dia + CoQ10 600 mg/dia (podem melhorar reserva ovariana residual)

Volume > 5 mL pós-menopausa:
• CA-125 (>35 U/mL suspeito, mas inespecífico)
• Doppler colorido (fluxo arterial central com IR <0.4 sugere malignidade)
• RM pelve com contraste (diferencia benigno/maligno)
• Se suspeita malignidade: cirurgia exploradora + estadiamento

Massa ovariana complexa (volume >20 mL + septações/componente sólido):
• Índice ROMA (combina CA-125, HE4, idade, menopausa)
• Se alto risco: encaminhar oncologia ginecológica (cirurgia citorredutora + quimioterapia)

Prevenção câncer ovariano:
• Contraceptivo oral >5 anos (reduz risco 50%)
• Salpingectomia bilateral profilática pós-fertilidade (remove trompas, origem câncer seroso)
• Rastreamento genético se história familiar: BRCA1/BRCA2 (ooforectomia profilática aos 35-40 anos se mutação)

Otimização reserva ovariana:
• CoQ10 ubiquinol 300-600 mg/dia (melhora função mitocondrial oocitária)
• DHEA 75 mg/dia (em mulheres >35 anos com reserva baixa)
• Vitamina D 4000 UI/dia (meta 25-OH 50-70 ng/mL)
• Ômega-3 2g/dia (reduz inflamação ovariana)
• Evitar tabaco (acelera atresia folicular)

Reavaliar volume ovariano + marcadores em 6-12 meses',

    last_review_date = CURRENT_DATE
WHERE id = 'afdd9d67-3f42-4e6e-a77d-0e75475ca72d';

-- 31. ECG - QTc (Bazett) - Homens
UPDATE score_items
SET
    interpretation = 'QTc (intervalo QT corrigido pela frequência cardíaca, fórmula de Bazett: QTc = QT/√RR) mede duração da repolarização ventricular. QTc prolongado (>450 ms homens, >460 ms mulheres) aumenta risco de torsades de pointes (taquicardia ventricular polimórfica) e morte súbita. Causas: congênitas (síndrome do QT longo - mutações canais iônicos KCNQ1, KCNH2, SCN5A), medicamentos (antiarrítmicos classe IA/III, antipsicóticos, antibióticos macrolídeos/fluoroquinolonas), distúrbios eletrolíticos (hipocalemia, hipomagnesemia, hipocalcemia), hipotireoidismo. QTc <350 ms (curto) também é arritmogênico.',

    clinical_relevance = 'Medicina funcional valoriza QTc em: rastreamento pré-medicação com drogas de risco (verificar sempre antes de antidepressivos tricíclicos, antipsicóticos, metadona); investigação de síncope inexplicada ou história familiar de morte súbita (QTc >500 ms tem risco alto); monitoramento de eletrólitos (potássio 4.0-5.0 mEq/L, magnésio >2.0 mg/dL); ajuste de medicações em idosos (acúmulo de drogas prolonga QTc). QTc 450-470 ms é zona cinza, exigindo correção de fatores reversíveis antes de adicionar drogas de risco.',

    optimal_value_male_min = 350.0,
    optimal_value_male_max = 440.0,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'QTc ótimo em homens: 350-440 ms (padrão MFI). Valores <350 ms indicam síndrome do QT curto (risco fibrilação ventricular). Faixa 440-450 ms é limítrofe, exigindo vigilância. QTc >450 ms define prolongamento patológico (risco torsades de pointes). QTc >500 ms indica risco muito alto de arritmia maligna (considerar CDI - cardiodesfibrilador implantável). Ideal: 380-420 ms, refletindo repolarização ventricular homogênea e estável.',

    health_recommendations = 'QTc 440-460 ms (limítrofe):
• Revisar medicações: suspender/substituir drogas que prolongam QT (lista em crediblemeds.org)
• Corrigir eletrólitos:
  - Potássio 4.5-5.0 mEq/L (suplementar KCl 20-40 mEq/dia se <4.0)
  - Magnésio >2.0 mg/dL (magnésio oral 400-600 mg/dia)
  - Cálcio 8.5-10.5 mg/dL (vitamina D 4000 UI/dia se baixo)
• Avaliar função tireoidiana (TSH, T4 livre - hipotireoidismo prolonga QT)

QTc > 460 ms (prolongado):
• ECG de confirmação (minimizar erro de medição)
• Investigar causas secundárias:
  - Eletrólitos (K, Mg, Ca)
  - Medicamentos (suspender imediatamente se possível)
  - TSH (hipotireoidismo)
  - História familiar (síndrome QT longo congênita)
• Teste genético se suspeita congênita (KCNQ1, KCNH2, SCN5A)

QTc > 500 ms (alto risco):
• EVITAR: antiarrítmicos classe IA (quinidina, procainamida) e III (sotalol, amiodarona)
• Beta-bloqueador: nadolol 40-160 mg/dia ou propranolol 80-320 mg/dia (reduz eventos em QT longo congênito em 60%)
• Magnésio IV se torsades de pointes: 2g bolus em 2 min
• Considerar CDI se: síncope recorrente, torsades documentada, QTc >550 ms

Prevenção torsades de pointes:
• Manter K >4.0 mEq/L, Mg >2.0 mg/dL (especialmente em uso de diuréticos)
• Evitar bradicardia (aumenta QT) - marca-passo se FC <50 bpm
• Evitar polifarmácia com drogas de risco

Medicamentos de ALTO RISCO que prolongam QT:
• Antiarrítmicos: sotalol, dofetilide, ibutilide, quinidina
• Antipsicóticos: haloperidol, ziprasidona, quetiapina
• Antidepressivos: citalopram (max 20 mg se >60 anos), escitalopram
• Antibióticos: azitromicina, eritromicina, levofloxacino, moxifloxacino
• Antifúngicos: fluconazol, voriconazol
• Outros: metadona, domperidona, ondansetrona

QTc < 350 ms (QT curto):
• Investigar síndrome QT curto congênita (teste genético)
• Risco fibrilação atrial e ventricular
• Considerar CDI se síncope ou FV documentada

Reavaliar ECG após correção de causas reversíveis (4-8 semanas)
ECG basal antes de iniciar qualquer droga de risco',

    last_review_date = CURRENT_DATE
WHERE id = '3a0d7e6b-c53d-47de-a50c-d7774e542835';

-- 32. ECG - QTc (Bazett) - Mulheres
UPDATE score_items
SET
    interpretation = 'QTc em mulheres é fisiologicamente ~10-20 ms maior que homens (efeito hormonal - estrogênio prolonga repolarização). QTc >460 ms em mulheres define prolongamento patológico. Mulheres têm maior risco de torsades de pointes induzida por drogas (60-70% dos casos) devido a: QTc basal maior, maior densidade de canais IKr (bloqueados por drogas), flutuações hormonais (fase lútea e gravidez prolongam QTc). Síndrome do QT longo congênita também é mais sintomática em mulheres (eventos arrítmicos em menarca, pós-parto).',

    clinical_relevance = 'Medicina funcional valoriza QTc em mulheres para: rastreamento rigoroso antes de drogas de risco (especialmente antidepressivos, antieméticos); monitoramento durante gravidez e puerpério (QTc pode aumentar 10-15 ms); investigação de síncope em fase lútea ou menarca (sugestivo de QT longo congênito); ajuste de medicações em pós-menopausa (menor risco que pré-menopausa). QTc >500 ms em mulher jovem com síncope requer avaliação cardiológica urgente e teste genético.',

    optimal_value_male_min = NULL,
    optimal_value_male_max = NULL,
    optimal_value_female_min = 360.0,
    optimal_value_female_max = 450.0,
    optimal_value_description = 'QTc ótimo em mulheres: 360-450 ms (padrão MFI). Valores <360 ms indicam síndrome QT curto. Faixa 450-460 ms é limítrofe (vigilância). QTc >460 ms define prolongamento patológico. QTc >500 ms indica risco muito alto de torsades de pointes (considerar CDI). Ideal: 390-430 ms. Mulheres têm QTc 10-20 ms maior que homens (fisiológico), mas limiares de risco são ajustados (>460 ms vs >450 ms em homens).',

    health_recommendations = 'QTc 450-470 ms (limítrofe):
• Revisar medicações: suspender/substituir drogas que prolongam QT
• Corrigir eletrólitos (K >4.0 mEq/L, Mg >2.0 mg/dL, Ca normal)
• Avaliar TSH (hipotireoidismo prolonga QT)
• Considerar fase do ciclo menstrual (QTc pode variar 5-10 ms)

QTc > 470 ms (prolongado):
• ECG de confirmação
• Investigar causas secundárias (eletrólitos, medicamentos, tireoide)
• História familiar de morte súbita ou síncope
• Teste genético se suspeita congênita

QTc > 500 ms (alto risco):
• Beta-bloqueador: nadolol 40-160 mg/dia (primeira linha em QT longo congênito)
• Magnésio IV se torsades: 2g bolus
• CDI se síncope recorrente ou torsades documentada

Gravidez e QTc:
• QTc aumenta 10-15 ms no 2º-3º trimestre (fisiológico)
• ECG basal no 1º trimestre se história de QT longo
• Pós-parto: maior risco de eventos arrítmicos em portadoras de QT longo (9 meses após parto)
• Manter beta-bloqueador se QT longo congênito

Contraceptivos orais e QTc:
• Estrogênio pode prolongar QTc em 5-10 ms
• Evitar se QTc basal >470 ms ou história QT longo congênito
• Preferir progestágenos isolados (não prolongam QT)

Terapia hormonal menopausa:
• Estrogênio oral prolonga QTc mais que transdérmico
• Monitorar ECG se QTc basal >450 ms

Medicações ALTO RISCO em mulheres:
• Antieméticos: ondansetrona (risco 3x maior em mulheres)
• Antidepressivos: citalopram, escitalopram (limitar dose mulheres >60 anos)
• Procinéticos: domperidona (evitar se QTc >450 ms)

QTc < 360 ms (QT curto):
• Investigar síndrome QT curto congênita
• Considerar CDI se síncope ou FV

Reavaliar ECG após correção (4-8 semanas)
ECG basal obrigatório antes de drogas de risco',

    last_review_date = CURRENT_DATE
WHERE id = '2e3c06ce-bcb6-4649-984e-8c30a92e26f4';

-- 33. ECG - Intervalo PR
UPDATE score_items
SET
    interpretation = 'Intervalo PR mede tempo de condução atrioventricular (nó AV + His-Purkinje), do início da onda P ao início do QRS. Normal: 120-200 ms. PR >200 ms define bloqueio AV de 1º grau (condução lenta, mas cada P conduz QRS). PR curto (<120 ms) pode indicar pré-excitação (Wolff-Parkinson-White - via acessória), ou variante fisiológica. Bloqueio AV 2º grau: algumas ondas P não conduzem (Mobitz I/Wenckebach - progressão gradual PR; Mobitz II - bloqueio súbito sem progressão). Bloqueio 3º grau: dissociação AV completa.',

    clinical_relevance = 'Medicina funcional valoriza PR em: diagnóstico de bloqueios AV (1º grau benigno, 2º grau Mobitz II e 3º grau requerem marca-passo); identificação de WPW (PR <120 ms + onda delta, risco taquicardia supraventricular e morte súbita); monitoramento de medicações cronotrópicas negativas (beta-bloqueadores, bloqueadores canais cálcio não-diidropiridínicos, digoxina); rastreamento de cardiopatia de Lyme (PR >300 ms é manifestação cardíaca comum). Atletas podem ter PR até 220 ms (hipertonia vagal) sem significado patológico.',

    optimal_value_male_min = 120.0,
    optimal_value_male_max = 200.0,
    optimal_value_female_min = 120.0,
    optimal_value_female_max = 200.0,
    optimal_value_description = 'Intervalo PR ótimo: 120-200 ms (padrão MFI). Valores <120 ms sugerem pré-excitação (WPW) ou variante fisiológica (investigar se taquicardias). PR >200 ms define bloqueio AV 1º grau (geralmente benigno, mas pode progredir). PR >300 ms indica bloqueio severo (investigar cardiopatia de Lyme, doença nó AV). Ideal: 140-180 ms, refletindo condução AV normal sem atraso significativo.',

    health_recommendations = 'PR < 120 ms (condução AV acelerada):
• Investigar síndrome de Wolff-Parkinson-White (WPW):
  - Onda delta (empastamento inicial QRS)
  - QRS alargado (>110 ms)
  - História de taquicardia paroxística
• Se WPW confirmado + sintomas: ablação por radiofrequência da via acessória (cura >95%)
• Se WPW assintomático: teste ergométrico (perda pré-excitação com exercício indica baixo risco)

PR 200-240 ms (bloqueio AV 1º grau leve):
• Geralmente benigno, sem necessidade de intervenção
• Revisar medicações: reduzir/suspender beta-bloqueadores, verapamil, diltiazem, digoxina
• Monitorar progressão: ECG anual ou se sintomas (fadiga, síncope)

PR > 240 ms (bloqueio AV 1º grau acentuado):
• Investigar doença nó AV (idiopática, fibrose, isquemia)
• Holter 24h (avaliar bloqueios intermitentes de 2º/3º grau)
• Se sintomático (fadiga, intolerância exercício): considerar marca-passo

PR > 300 ms:
• Descartar cardiopatia de Lyme: sorologia Borrelia burgdorferi (IgG/IgM), Western blot
• Se Lyme+: ceftriaxona 2g IV/dia por 14-21 dias (bloqueio AV reverte em 80-90%)
• Ecocardiograma (avaliar função ventricular)

Bloqueio AV 2º grau Mobitz I (Wenckebach):
• PR alonga progressivamente até P bloqueado
• Geralmente benigno (nível nó AV)
• Monitorar, marca-passo apenas se sintomático

Bloqueio AV 2º grau Mobitz II:
• P bloqueado sem prolongamento PR prévio
• Alto risco progressão para bloqueio 3º grau
• Marca-passo definitivo indicado (mesmo se assintomático)

Bloqueio AV 3º grau (completo):
• Dissociação AV, ritmo de escape (ventricular 20-40 bpm ou juncional 40-60 bpm)
• Marca-passo definitivo URGENTE
• Temporário se adquirido (IAM inferior, Lyme, medicamentos)

Prevenção progressão bloqueio AV:
• Evitar medicações cronotrópicas negativas se PR >200 ms
• Controlar fatores risco: hipertensão, diabetes, doença coronariana
• CoQ10 100-300 mg/dia (suporte função mitocondrial nó AV)

Reavaliar ECG anualmente se PR >200 ms
Holter se sintomas (síncope, fadiga, dispneia)',

    last_review_date = CURRENT_DATE
WHERE id = 'b2dd0c76-7bce-4beb-a8e2-52d70d467241';

-- 34. ECG - Sokolow-Lyon (S V1 + R V5/V6)
UPDATE score_items
SET
    interpretation = 'Critério Sokolow-Lyon (SV1 + RV5 ou RV6 >35 mm) é usado para diagnóstico eletrocardiográfico de hipertrofia ventricular esquerda (HVE). HVE resulta de sobrecarga pressórica crônica (hipertensão arterial, estenose aórtica) ou volumétrica (insuficiência aórtica, miocardiopatia dilatada). HVE aumenta risco de insuficiência cardíaca, arritmias ventriculares, morte súbita e doença coronariana. Sensibilidade do Sokolow-Lyon é baixa (~20-30%), mas especificidade é alta (~90%). Cornell voltage (RaVL + SV3 >28 mm homens, >20 mm mulheres) é alternativa mais sensível.',

    clinical_relevance = 'Medicina funcional valoriza Sokolow-Lyon em: rastreamento de HVE em hipertensos (indica lesão de órgão-alvo, requer tratamento agressivo); monitoramento de regressão HVE com tratamento anti-hipertensivo (bloqueadores SRAA - IECA/BRA - são mais efetivos); diagnóstico diferencial de sobrecarga ventricular (HVE concêntrica vs excêntrica); estratificação de risco cardiovascular. Ecocardiograma com cálculo de massa VE é padrão-ouro (índice massa VE >115 g/m² homens, >95 g/m² mulheres define HVE).',

    optimal_value_male_min = 10.0,
    optimal_value_male_max = 35.0,
    optimal_value_female_min = 10.0,
    optimal_value_female_max = 35.0,
    optimal_value_description = 'Sokolow-Lyon ótimo: 10-35 mm (SV1 + RV5/RV6). Valores <10 mm podem indicar baixa voltagem (obesidade, DPOC, derrame pericárdico). Faixa >35 mm define HVE eletrocardiográfica. Valores >45 mm indicam HVE severa. Ideal: 20-30 mm, refletindo massa ventricular normal sem hipertrofia. Importante: jovens/atletas podem ter >35 mm sem HVE patológica (hipertrofia fisiológica).',

    health_recommendations = 'Sokolow-Lyon > 35 mm (HVE):
• Confirmar com ecocardiograma transtorácico:
  - Massa VE indexada >115 g/m² (homens) ou >95 g/m² (mulheres)
  - Espessura septo/parede posterior >11 mm (concêntrica) ou >12 mm (excêntrica)
  - Avaliar função sistólica (FEVE), diastólica (E/e'), valvas

HVE confirmada - tratamento hipertensão:
• Meta PA <130/80 mmHg (regressão HVE em 60-70% dos casos)
• IECA ou BRA (primeira linha, maior regressão HVE):
  - Enalapril 10-40 mg/dia ou losartana 50-100 mg/dia
• Beta-bloqueador se taquicardia ou DAC associada
• Bloqueador canal cálcio (anlodipino 5-10 mg/dia) adicional se necessário

Regressão HVE:
• Perda de peso 5-10% (reduz massa VE ~10 g/m²)
• Restrição sódio <2300 mg/dia (<1500 mg se HAS resistente)
• Exercício aeróbico 150 min/semana (reduz HVE em 15-20%)
• DASH diet (rico em K, Mg, Ca - frutas, vegetais, grãos integrais, laticínios magros)

HVE + DAC ou angina:
• Evitar hidralazina, minoxidil (vasodilatadores diretos aumentam HVE)
• Preferir beta-bloqueadores + IECA/BRA

HVE + arritmias:
• Investigar fibrilação atrial (Holter 24h)
• Tratar FA se presente (anticoagulação + controle FC)
• Considerar ablação se FA paroxística recorrente

HVE atlética (jovem, atleta, eco normal):
• Diferenciar de miocardiopatia hipertrófica:
  - Eco: espessura septo <13 mm (atlética) vs >15 mm (patológica)
  - RM cardíaca se dúvida (realce tardio gadolíneo ausente em atlética)
  - História familiar negativa para morte súbita

Suplementação cardioprotectora:
• CoQ10 ubiquinol 200-300 mg/dia (melhora função mitocondrial, reduz HVE)
• Magnésio 400 mg/dia (vasodilatação, reduz PA)
• Ômega-3 2-3g/dia (anti-inflamatório, reduz remodelamento)
• Vitamina D 4000 UI/dia (meta 25-OH >50 ng/mL, modula SRAA)

Reavaliar ECG + ecocardiograma em 12-24 meses
Meta: regressão HVE (massa VE normalizada)',

    last_review_date = CURRENT_DATE
WHERE id = '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb';

-- 35. Ecodopplercardiograma - FEVE Homens
UPDATE score_items
SET
    interpretation = 'Fração de ejeção do ventrículo esquerdo (FEVE) é percentual do volume diastólico final ejetado na sístole, medida pelo método Simpson biplanar. FEVE normal: ≥55%. FEVE 41-54% define disfunção leve, 30-40% moderada, <30% severa. FEVE reduzida (insuficiência cardíaca com fração ejeção reduzida - ICFEr) resulta de DAC (IAM prévio), miocardiopatia dilatada, valvopatias, hipertensão não-controlada ou tóxicos (álcool, quimioterapia). FEVE preservada (≥50%) com sintomas IC define insuficiência cardíaca com fração ejeção preservada (ICFEp), comum em idosos e hipertensos.',

    clinical_relevance = 'Medicina funcional valoriza FEVE em: estratificação de risco IC (FEVE <35% indica CDI para prevenção morte súbita); monitoramento de cardiotoxicidade (quimioterapia com antraciclinas - doxorrubicina reduz FEVE); avaliação pré-operatória (FEVE <40% aumenta risco cirúrgico); titulação de terapia IC (meta medicações: IECA/BRA, beta-bloqueador, espironolactona, SGLT2-i). Melhora FEVE >10% após tratamento otimizado indica bom prognóstico (reversibilidade miocárdica).',

    optimal_value_male_min = 55.0,
    optimal_value_male_max = 70.0,
    optimal_value_female_min = NULL,
    optimal_value_female_max = NULL,
    optimal_value_description = 'FEVE ótima em homens: 55-70% (padrão MFI). Valores <55% indicam disfunção sistólica (investigar causa: DAC, miocardiopatia, valvopatia). FEVE 41-54% é disfunção leve (otimizar tratamento IC). FEVE <40% define ICFEr (tratamento farmacológico triplo obrigatório + considerar dispositivos). FEVE >70% pode ser hiperdinámica (hipertireoidismo, anemia, choque séptico compensado). Ideal: 60-65%, refletindo contratilidade miocárdica robusta e reserva cardíaca adequada.',

    health_recommendations = 'FEVE 41-54% (disfunção leve):
• Investigar etiologia: cineangiocoronariografia (DAC), RM cardíaca (fibrose, viabilidade)
• Tratamento IC:
  - IECA ou BRA: enalapril 10-20 mg 2x/dia ou sacubitril-valsartana 49/51 mg 2x/dia
  - Beta-bloqueador: carvedilol 25 mg 2x/dia ou metoprolol succinato 200 mg/dia
  - Espironolactona 25-50 mg/dia (se FEVE <45%)
• Controle fatores risco: PA <130/80 mmHg, LDL <70 mg/dL, HbA1c <7%

FEVE < 40% (ICFEr):
• Terapia quádrupla (GDMT - guideline-directed medical therapy):
  1. IECA/ARNI: sacubitril-valsartana 97/103 mg 2x/dia (superior a IECA)
  2. Beta-bloqueador: carvedilol 25 mg 2x/dia
  3. ARM: espironolactona 25-50 mg/dia (reduz mortalidade 30%)
  4. SGLT2-i: dapagliflozina 10 mg/dia ou empagliflozina 10 mg/dia (reduz hospitalizações 30%)

FEVE < 35% + QRS ≥ 130 ms (BRE):
• Terapia ressincronização cardíaca (TRC - biventricular pacing, melhora FEVE ~5-10%)

FEVE < 35% sintomático (NYHA II-III):
• Cardiodesfibrilador implantável (CDI) para prevenção primária morte súbita
• Considerar TRC-D (ressincronizador + desfibrilador)

Reversão miocardiopatia:
• DAC: revascularização (angioplastia ou bypass) se viabilidade presente (RM com dobutamina)
• Álcool: abstinência total (FEVE pode melhorar 10-15% em 6 meses)
• Tireotoxicose: tratamento hipertireoidismo (metimazol, betabloqueador)
• Taquicardiomiopatia: controle FC <80 bpm (ablação FA se refratária)

Suplementação cardioprotectora:
• CoQ10 ubiquinol 300-600 mg/dia (melhora FEVE ~3-5%, reduz sintomas)
• L-carnitina 2g/dia (melhora metabolismo energético miocárdico)
• D-ribose 5g 3x/dia (substrato ATP, controverso)
• Ômega-3 4g/dia (reduz hospitaliz. IC em REDUCE-IT)
• Tiamina (B1) 200 mg/dia (deficiência comum em IC, diurese)

Exercício supervisionado:
• Reabilitação cardíaca: aeróbico 30 min 5x/semana (melhora FEVE, VO2 máx, sintomas)
• Evitar exercício isométrico intenso (aumenta pós-carga)

Dieta IC:
• Restrição sódio <2g/dia (<1.5g se grave)
• Restrição hídrica <1.5L/dia se hiponatremia ou congestão
• DASH diet adaptada

Monitoramento:
• Ecocardiograma a cada 6-12 meses
• BNP/NT-proBNP (redução >30% indica melhora)
• Peso diário (ganho >2 kg em 3 dias = descompensação)

Reavaliar FEVE em 6-12 meses
Meta: FEVE >40% ou melhora >10%',

    last_review_date = CURRENT_DATE
WHERE id = '6e542cb0-6982-42e8-93dc-40f139652223';

-- (Continuação com os itens restantes no próximo bloco...)

COMMIT;
