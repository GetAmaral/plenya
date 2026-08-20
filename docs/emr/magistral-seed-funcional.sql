-- Catálogo magistral: substâncias mais usadas em medicina integrativa/funcional e nutrologia.
--
-- COMO ISTO FOI MONTADO: levantamento online de formulários e fórmulas publicadas por farmácias
-- de manipulação brasileiras (aFormula, Pharmac, Beleza Saúde, BioVittare, João Falcão, Farmacam,
-- Farmabotânica, Invictus, Alquimia, Manipula, MaxPharma, entre outras) mais os limites do
-- ANEXO IV da IN 28/2018 da Anvisa para suplementos alimentares.
--
-- REGRA MANTIDA: cada dose gravada tem fonte. Substância cuja dose não apareceu em fonte fica com
-- `usual_dose` NULL e o que se sabe registrado em `notes` — a tela trata NULL como silêncio, não
-- como zero, e o médico preenche pelo seu próprio protocolo (o botão "salvar como padrão" grava).
--
-- DUAS ADVERTÊNCIAS QUE VALEM MAIS QUE OS NÚMEROS:
--
-- 1. O teto da IN 28/2018 é de SUPLEMENTO ALIMENTAR INDUSTRIALIZADO, não de prescrição magistral.
--    Vitamina D (2.000 UI), biotina (45 mcg) e B12 (9,94 mcg) são tetos de rótulo que a prática
--    médica supera deliberadamente. Por isso esses tetos foram para `notes`, e NÃO para
--    `max_dose` — colocá-los em max_dose faria a tela alertar em toda receita legítima, e alerta
--    que grita à toa é alerta que o médico aprende a ignorar.
--
-- 2. DENSIDADE APARENTE continua fora, salvo dois valores publicados (sulfato de condroitina
--    0,60 g/mL e maltodextrina 0,70 g/mL). Não existe tabela pública por ativo e o valor muda com
--    o lote e a compactação da farmácia. Enquanto não for cadastrada, a calculadora de cápsula
--    diz que não sabe.
--
-- HIGROSCOPICIDADE: marcada só onde a fonte nomeia a substância (cloreto de magnésio, carnitina,
-- acetil-carnitina, fosfolipídios, minerais quelados). As fontes citam "aminoácidos" e "extratos
-- secos" como classes sensíveis à umidade, mas marcar a classe inteira faria a tela avisar em
-- quase toda cápsula — fica registrado aqui e o sinalizador é seu para ligar caso queira.
--
-- Idempotente: pode rodar de novo sem duplicar.

INSERT INTO public.magistral_components
    (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, hygroscopic, source, notes, last_review)
VALUES
    (uuidv7(), 'Melatonina', '', 'mg', 5, 0.5, 10, false, 'pesquisa', 'Fórmulas magistrais BR: 1 a 5 mg por dose ao deitar.', now()),
    (uuidv7(), '5-HTP', '5-hidroxitriptofano, Griffonia simplicifolia', 'mg', 100, 50, 100, false, 'pesquisa', 'Fórmulas BR: 50, 75 e 100 mg por dose.', now()),
    (uuidv7(), 'L-teanina', 'teanina', 'mg', 100, 50, 500, false, 'pesquisa', 'Fórmulas BR: 50, 100, 250 e 500 mg por dose.', now()),
    (uuidv7(), 'L-triptofano', 'triptofano', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada usava 70 mg em associação; a faixa isolada varia muito. Sem dose padrão definida.', now()),
    (uuidv7(), 'GABA', 'ácido gama-aminobutírico', 'mg', 80, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 80 mg por dose.', now()),
    (uuidv7(), 'Valeriana', 'Valeriana officinalis', 'mg', 250, 50, 250, false, 'pesquisa', 'Fórmulas BR: 50 a 250 mg por dose.', now()),
    (uuidv7(), 'Passiflora', 'Passiflora incarnata, maracujá', 'mg', 40, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 40 mg por dose.', now()),
    (uuidv7(), 'Glicina', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Em fórmula de sono BR aparece como 75 mg; para sono isolada a literatura usa 3 g. Contextos diferentes — defina o seu.', now()),
    (uuidv7(), 'Magnésio dimalato', 'magnésio malato', 'mg', 550, NULL, 1500, false, 'pesquisa', 'Dose usual até 1.500 mg/dia (formulário magistral). Teto IN 28/2018 para suplemento é 350 mg de magnésio elementar/dia — atenção: o rótulo do sal não é magnésio elementar.', now()),
    (uuidv7(), 'Magnésio quelato', 'magnésio bisglicinato, magnésio glicinato', 'mg', 100, 50, 500, true, 'pesquisa', 'Dose usual 50 a 500 mg/dia (formulário magistral). Mineral quelado é citado como sensível à umidade (cápsula vegetal recomendada).', now()),
    (uuidv7(), 'Magnésio L-treonato', 'magtein, treonato de magnésio', 'mg', 300, NULL, 500, false, 'pesquisa', 'Dose usual até 500 mg (formulário magistral). Fórmulas BR: 150, 200, 300 e 500 mg.', now()),
    (uuidv7(), 'Cloreto de magnésio', '', 'mg', NULL, NULL, NULL, true, 'pesquisa', 'HIGROSCÓPICO — absorve umidade do ar e amolece cápsula gelatinosa. Excipiente adsorvente ou cápsula vegetal.', now()),
    (uuidv7(), 'Magnésio taurato', 'taurato de magnésio', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Forma citada em prescrição magistral; sem dose padrão levantada.', now()),
    (uuidv7(), 'Coenzima Q10', 'ubiquinona, CoQ10', 'mg', 100, 50, 200, false, 'pesquisa', 'Fórmulas BR: 50, 100 e 150 mg. Teto IN 28/2018 para suplemento: 200 mg/dia.', now()),
    (uuidv7(), 'Ubiquinol', 'coenzima Q10 reduzida', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Forma reduzida da CoQ10, usada em suporte mitocondrial; sem dose padrão levantada.', now()),
    (uuidv7(), 'PQQ', 'pirroloquinolina quinona', 'mg', 10, 5, 11, false, 'pesquisa', 'Fórmulas BR: 5, 10 e 11 mg por dose.', now()),
    (uuidv7(), 'Ácido alfa-lipoico', 'ácido tióctico, ALA', 'mg', 100, 25, 200, false, 'pesquisa', 'Fórmulas BR: 25 a 200 mg por dose.', now()),
    (uuidv7(), 'Ácido R-alfa-lipoico', 'R-ALA', 'mg', 25, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 25 mg por dose.', now()),
    (uuidv7(), 'L-carnitina', 'carnitina', 'mg', 500, 250, 500, true, 'pesquisa', 'Fórmulas BR: 250 e 500 mg. Carnitina é citada como sensível à umidade (cápsula vegetal recomendada).', now()),
    (uuidv7(), 'Acetil-L-carnitina', 'ALCAR', 'mg', 250, 50, 250, true, 'pesquisa', 'Fórmulas BR: 50 a 250 mg. Sensível à umidade, como a carnitina.', now()),
    (uuidv7(), 'NADH', 'nicotinamida adenina dinucleotídeo reduzida', 'mg', 10, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 10 mg por dose.', now()),
    (uuidv7(), 'Nicotinamida', 'niacinamida, vitamina B3', 'mg', 25, NULL, NULL, false, 'pesquisa', 'Fórmula BR: 25 mg. Teto IN 28/2018 para niacina em suplemento: 35 mg/dia.', now()),
    (uuidv7(), 'Riboflavina', 'vitamina B2', 'mg', 5, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 5 mg por dose.', now()),
    (uuidv7(), 'Creatina', 'creatina monoidratada', 'mg', 3000, NULL, 3000, false, 'pesquisa', 'Teto IN 28/2018 para suplemento: 3 g/dia. Sabor levemente amargo quando pura, relevante em sachê.', now()),
    (uuidv7(), 'Ácido málico', '', 'mg', 25, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 25 mg por dose.', now()),
    (uuidv7(), 'Teacrina', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Citada em fórmulas de energia/foco; sem dose padrão levantada.', now()),
    (uuidv7(), 'Berberina', 'cloridrato de berberina', 'mg', 250, 125, 500, false, 'pesquisa', 'Fórmulas BR: 125, 200, 250 e 500 mg. Sabor acentuadamente amargo — avaliar antes de sachê.', now()),
    (uuidv7(), 'Picolinato de cromo', 'cromo picolinato', 'mcg', 200, 25, 250, false, 'pesquisa', 'Fórmulas BR: 25 a 200 mcg. Teto IN 28/2018 para cromo em suplemento: 250 mcg/dia.', now()),
    (uuidv7(), 'Mio-inositol', 'inositol', 'mg', NULL, NULL, 2000, false, 'pesquisa', 'Fórmula BR em associação: 100 a 250 mg. Teto IN 28/2018 para inositol em suplemento: 2 g/dia; protocolos de SOP usam 2 a 4 g/dia.', now()),
    (uuidv7(), 'D-quiro-inositol', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Usado com mio-inositol em protocolos de SOP; sem dose padrão levantada.', now()),
    (uuidv7(), 'Gymnema silvestre', 'Gymnema sylvestre', 'mg', 25, 12.5, 25, false, 'pesquisa', 'Fórmulas BR: 12,5 e 25 mg por dose.', now()),
    (uuidv7(), 'Silimarina', 'Silybum marianum, cardo mariano', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Presente em fórmula BR de resistência insulínica; dose não especificada na fonte.', now()),
    (uuidv7(), 'Ginseng brasileiro', 'Pfaffia paniculata, suma', 'mg', 60, 30, 60, false, 'pesquisa', 'Fórmulas BR: 30 e 60 mg por dose.', now()),
    (uuidv7(), 'Vitamina C', 'ácido ascórbico', 'mg', 300, 100, 1916, false, 'pesquisa', 'Fórmulas BR: 100 a 300 mg. Teto IN 28/2018 para suplemento: 1.916,02 mg/dia.', now()),
    (uuidv7(), 'Quercetina', '', 'mg', 200, 200, 2000, false, 'pesquisa', 'Fórmulas BR: 200 mg. Faixa clínica descrita: 500 a 2.000 mg/dia.', now()),
    (uuidv7(), 'N-acetilcisteína', 'NAC, acetilcisteína', 'mg', 600, 200, 600, false, 'pesquisa', 'Fórmulas BR: 200 a 600 mg por dose.', now()),
    (uuidv7(), 'Curcumina', 'cúrcuma longa, açafrão-da-terra', 'mg', 100, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 100 mg por dose (extrato).', now()),
    (uuidv7(), 'Trans-resveratrol', 'resveratrol', 'mg', 100, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 100 mg por dose.', now()),
    (uuidv7(), 'Astaxantina', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Citada em fórmulas antioxidantes BR; dose não especificada na fonte.', now()),
    (uuidv7(), 'Extrato de chá verde', 'Camellia sinensis, EGCG', 'mg', 50, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 50 mg por dose.', now()),
    (uuidv7(), 'Selênio', 'selenometionina', 'mcg', NULL, NULL, 319, false, 'pesquisa', 'Fórmula BR pesquisada: 30 mcg. Teto IN 28/2018 para suplemento: 319,75 mcg/dia.', now()),
    (uuidv7(), 'Zinco quelato', 'zinco bisglicinato', 'mg', 25, 5, 29, true, 'pesquisa', 'Fórmulas BR: 5 a 35 mg. Teto IN 28/2018 para zinco em suplemento: 29,59 mg/dia. Mineral quelado é citado como sensível à umidade.', now()),
    (uuidv7(), 'Vitamina D3', 'colecalciferol', 'UI', 2000, NULL, NULL, false, 'pesquisa', 'Fórmulas BR: 1.000 a 5.000 UI. Teto IN 28/2018 para SUPLEMENTO: 50 mcg (2.000 UI)/dia — prescrição magistral médica costuma superar isso de forma deliberada.', now()),
    (uuidv7(), 'Vitamina K2 MK-7', 'menaquinona-7', 'mcg', 100, 65, 100, false, 'pesquisa', 'Fórmulas BR: 65 e 100 mcg por dose.', now()),
    (uuidv7(), 'Vitamina E', 'tocoferol', 'mg', 400, NULL, 1000, false, 'pesquisa', 'Fórmula BR: 400 mg. Teto IN 28/2018 para suplemento: 1.000 mg/dia.', now()),
    (uuidv7(), 'Vitamina A', 'retinol', 'UI', 2000, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 2.000 UI por dose.', now()),
    (uuidv7(), 'Licopeno', '', 'mg', 10, NULL, 8, false, 'pesquisa', 'Fórmula BR pesquisada: 10 mg. Teto IN 28/2018 para suplemento: 8 mg/dia — a fórmula magistral pesquisada supera o teto de suplemento.', now()),
    (uuidv7(), 'Cobre', '', 'mg', 1, NULL, 8.9, false, 'pesquisa', 'Fórmula BR pesquisada: 1 mg. Teto IN 28/2018 para suplemento: 8.975,52 mcg/dia.', now()),
    (uuidv7(), 'L-glutamina', 'glutamina', 'mg', 1000, 150, NULL, false, 'pesquisa', 'Fórmulas BR: 150 mg em cápsula associada; 1 g em sachê de suporte intestinal.', now()),
    (uuidv7(), 'Zinco carnosina', 'carnosina de zinco', 'mg', 75, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 75 mg por dose.', now()),
    (uuidv7(), 'FOS', 'fruto-oligossacarídeo', 'g', 2, NULL, NULL, false, 'pesquisa', 'Sachê BR pesquisado: 2 g por dose.', now()),
    (uuidv7(), 'Inulina', '', 'g', 2, NULL, NULL, false, 'pesquisa', 'Sachê BR pesquisado: 2 g por dose.', now()),
    (uuidv7(), 'XOS', 'xilo-oligossacarídeo', 'g', 2, NULL, NULL, false, 'pesquisa', 'Sachê BR pesquisado: 2 g por dose.', now()),
    (uuidv7(), 'Butirato de sódio', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Usado em suporte de mucosa intestinal; dose não levantada nas fontes consultadas.', now()),
    (uuidv7(), 'Probiótico multicepas', 'Lactobacillus, Bifidobacterium', 'UFC', 5000000000, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: pool de 5 cepas, 5 bilhões de UFC por sachê. Confirme a unidade com a farmácia (UFC por dose).', now()),
    (uuidv7(), 'Alfa-GPC', 'alfa-glicerilfosforilcolina', 'mg', 300, NULL, NULL, false, 'pesquisa', 'Fórmulas de suporte cognitivo: 300 mg por dose.', now()),
    (uuidv7(), 'Citicolina', 'CDP-colina', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Usada em suporte cognitivo; dose não especificada nas fontes consultadas.', now()),
    (uuidv7(), 'Fosfatidilserina', '', 'mg', NULL, NULL, NULL, true, 'pesquisa', 'Fosfolipídio citado como sensível à umidade (cápsula vegetal recomendada).', now()),
    (uuidv7(), 'Fosfatidilcolina', '', 'mg', NULL, NULL, NULL, true, 'pesquisa', 'Citada explicitamente entre os ativos que pedem cápsula vegetal por sensibilidade à umidade.', now()),
    (uuidv7(), 'Bacopa monnieri', 'bacopa', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Extrato usado em suporte cognitivo; dose não especificada nas fontes consultadas.', now()),
    (uuidv7(), 'Ginkgo biloba', '', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Extrato usado em suporte cognitivo; dose não especificada nas fontes consultadas.', now()),
    (uuidv7(), 'Huperzina A', 'huperzine A', 'mcg', NULL, NULL, NULL, false, 'pesquisa', 'Usada em suporte cognitivo; dose não especificada nas fontes consultadas.', now()),
    (uuidv7(), 'Silício orgânico', 'Exsynutriment, ácido ortossilícico', 'mg', 100, 100, 300, false, 'pesquisa', 'Fórmulas BR: 100 e 300 mg por dose.', now()),
    (uuidv7(), 'Biotina', 'vitamina B7, vitamina H', 'mcg', 500, NULL, NULL, false, 'pesquisa', 'Fórmulas BR: 500 mcg a 10 mg. Teto IN 28/2018 para SUPLEMENTO: 45 mcg/dia — a prática magistral vai muito além, de forma deliberada.', now()),
    (uuidv7(), 'Ácido hialurônico', '', 'mg', 100, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 100 mg por dose.', now()),
    (uuidv7(), 'Colágeno Verisol', 'peptídeo bioativo de colágeno', 'g', 2.5, NULL, NULL, false, 'pesquisa', 'Dose de referência do ativo em pó; confirmar com a farmácia.', now()),
    (uuidv7(), 'Colágeno hidrolisado', '', 'g', 10, NULL, NULL, false, 'pesquisa', 'Fórmulas em pó BR: cerca de 10 g por dose.', now()),
    (uuidv7(), 'MSM', 'metilsulfonilmetano', 'mg', NULL, NULL, NULL, false, 'pesquisa', 'Usado em fórmulas de pele/articulação; dose não especificada nas fontes consultadas.', now()),
    (uuidv7(), 'Isoflavona', 'Glycine max, isoflavona de soja', 'mg', 80, 40, 150, false, 'pesquisa', 'Dose para sintomas do climatério: 40 a 150 mg/dia.', now()),
    (uuidv7(), 'Cimicifuga racemosa', 'black cohosh', 'mg', 200, 40, 200, false, 'pesquisa', 'Fórmulas BR: 40 e 200 mg por dose.', now()),
    (uuidv7(), 'Vitex agnus-castus', 'agnus castus', 'mg', 200, 20, 200, false, 'pesquisa', 'Fórmulas BR: 20 e 200 mg por dose.', now()),
    (uuidv7(), 'Amora', 'Morus nigra', 'mg', 500, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 500 mg por dose.', now()),
    (uuidv7(), 'Ashwagandha', 'Withania somnifera, KSM-66, Sensoril', 'mg', 500, 500, 1000, false, 'pesquisa', 'Extrato padronizado: 500 mg a 1 g/dia.', now()),
    (uuidv7(), 'Mucuna pruriens', 'mucuna', 'mg', 150, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 150 mg por dose.', now()),
    (uuidv7(), 'Yam mexicano', 'Dioscorea villosa', 'mg', 200, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 200 mg por dose.', now()),
    (uuidv7(), 'Relora', 'Magnolia officinalis + Phellodendron', 'mg', 20, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 20 mg por dose.', now()),
    (uuidv7(), 'Tongkat ali', 'Long Jack, Eurycoma longifolia', 'mg', 50, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 50 mg por dose.', now()),
    (uuidv7(), 'L-taurina', 'taurina', 'mg', 100, NULL, NULL, false, 'pesquisa', 'Fórmula BR pesquisada: 100 mg por dose.', now()),
    (uuidv7(), 'Piridoxal-5-fosfato', 'P5P, vitamina B6 ativa', 'mg', 25, NULL, NULL, false, 'pesquisa', 'Fórmula BR: 25 mg. Teto IN 28/2018 para vitamina B6 em suplemento: 98,60 mg/dia.', now()),
    (uuidv7(), 'Metilcobalamina', 'vitamina B12 metilada', 'mcg', 100, NULL, NULL, false, 'pesquisa', 'Fórmulas BR: 62,5 mcg a 1 mg. Teto IN 28/2018 para SUPLEMENTO: 9,94 mcg/dia — a prática magistral vai muito além, de forma deliberada.', now()),
    (uuidv7(), 'Metilfolato', '5-MTHF, folato metilado', 'mcg', NULL, NULL, NULL, false, 'pesquisa', 'Citado em fórmulas BR de metilação; dose não especificada na fonte.', now()),
    (uuidv7(), 'Ômega-3', 'EPA, DHA, óleo de peixe', 'mg', 1000, NULL, 2000, false, 'pesquisa', 'Teto IN 28/2018 para EPA+DHA em suplemento: 2.000 mg/dia.', now()),
    (uuidv7(), 'Cálcio', 'citrato de cálcio, CMG', 'mg', 100, NULL, 1534, false, 'pesquisa', 'Fórmula BR: 100 mg. Teto IN 28/2018 para suplemento: 1.534,67 mg/dia.', now()),
    (uuidv7(), 'Colina', 'bitartarato de colina', 'mg', NULL, NULL, 3235, false, 'pesquisa', 'Teto IN 28/2018 para suplemento: 3.235,15 mg/dia.', now()),
    (uuidv7(), 'Iodo', 'iodeto de potássio', 'mcg', NULL, NULL, 919, false, 'pesquisa', 'Teto IN 28/2018 para suplemento: 919,02 mcg/dia.', now()),
    (uuidv7(), 'Ferro', 'ferro quelato, bisglicinato ferroso', 'mg', NULL, NULL, 34, true, 'pesquisa', 'Teto IN 28/2018 para suplemento: 34,31 mg/dia. Mineral quelado é citado como sensível à umidade.', now()),
    (uuidv7(), 'Sulfato de condroitina', 'condroitina', 'mg', 1000, NULL, NULL, false, 'pesquisa', 'Densidade aparente 0,60 g/mL (exemplo publicado de cálculo de sachê). Dose de referência do exemplo: 1 g/sachê.', now()),
    (uuidv7(), 'Maltodextrina', '', 'g', NULL, NULL, NULL, false, 'pesquisa', 'Diluente. Densidade aparente 0,70 g/mL (fonte: e-book de manipulação de sachês).', now())
ON CONFLICT DO NOTHING;


-- Pares que a literatura de suplementação descreve como COMPETIÇÃO DE ABSORÇÃO — não são
-- incompatibilidades de manipulação. Entram como 'info' e o mecanismo diz isso com todas as
-- letras, para não virar alarme de fórmula que a farmácia manipula sem problema nenhum.
INSERT INTO public.magistral_incompatibilities
    (id, component_a_id, component_b_id, severity, mechanism, note, source, last_review)
SELECT uuidv7(),
       LEAST(a.id::text, b.id::text)::uuid,
       GREATEST(a.id::text, b.id::text)::uuid,
       p.severity, p.mechanism, p.note, p.source, now()
FROM (VALUES
    ('Zinco quelato','Cobre','info',
     'competição de absorção: zinco em uso continuado acima de 40 mg/dia tende a reduzir o cobre',
     'Não impede a associação; é o motivo de as fórmulas com zinco alto levarem cobre junto.',
     'Literatura de suplementação (competição de absorção)'),
    ('Zinco quelato','Ferro','info',
     'competição de absorção pela mesma via intestinal',
     'Se a intenção for repor os dois, considere doses ou horários separados.',
     'Literatura de suplementação (competição de absorção)'),
    ('Cálcio','Ferro','info',
     'o cálcio reduz a absorção do ferro',
     'Separar as tomadas resolve; a manipulação em si não é impedida.',
     'Literatura de suplementação (competição de absorção)'),
    ('Vitamina C','Metilcobalamina','info',
     'a vitamina C em dose alta pode degradar a cobalamina quando ficam juntas',
     'Descrito para as duas na mesma tomada; separar em fórmulas ou horários evita a dúvida.',
     'Literatura de suplementação')
) AS p(nome_a, nome_b, severity, mechanism, note, source)
JOIN public.magistral_components a ON lower(a.name) = lower(p.nome_a)
JOIN public.magistral_components b ON lower(b.name) = lower(p.nome_b)
ON CONFLICT DO NOTHING;

-- Densidades aparentes PUBLICADAS (as duas únicas que a pesquisa encontrou com fonte).
-- Fonte: "Facilitando a manipulação de sachês em farmácias" (Ideal Equipamentos), exemplo de
-- cálculo volumétrico. Continua valendo o aviso: densidade varia por lote e por farmácia.
UPDATE public.magistral_components SET bulk_density = 0.60
 WHERE lower(name) = 'sulfato de condroitina' AND bulk_density IS NULL;
UPDATE public.magistral_components SET bulk_density = 0.70
 WHERE lower(name) = 'maltodextrina' AND bulk_density IS NULL;

-- Substâncias que já tinham entrado no catálogo por USO (esboço criado ao prescrever) recebem o
-- que a pesquisa trouxe, sem sobrescrever o que já foi curado à mão.
UPDATE public.magistral_components c SET
    synonyms   = CASE WHEN c.synonyms = '' THEN v.synonyms ELSE c.synonyms END,
    usual_dose = COALESCE(c.usual_dose, v.usual_dose),
    min_dose   = COALESCE(c.min_dose, v.min_dose),
    max_dose   = COALESCE(c.max_dose, v.max_dose),
    notes      = COALESCE(c.notes, v.notes),
    last_review = now()
FROM (VALUES
    ('Magnésio dimalato','magnésio malato', 550::numeric, NULL::numeric, 1500::numeric,
     'Dose usual até 1.500 mg/dia (formulário magistral).'),
    ('Vitamina C','ácido ascórbico', 300::numeric, 100::numeric, 1916::numeric,
     'Fórmulas BR: 100 a 300 mg. Teto IN 28/2018 para suplemento: 1.916,02 mg/dia.'),
    ('Melatonina','', 5::numeric, 0.5::numeric, 10::numeric,
     'Fórmulas magistrais BR: 1 a 5 mg por dose ao deitar.')
) AS v(name, synonyms, usual_dose, min_dose, max_dose, notes)
WHERE lower(c.name) = lower(v.name);
