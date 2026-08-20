-- +goose Up
-- +goose StatementBegin

-- CONTEUDO do catalogo magistral: 290 substancias, 132 formulas-base, 54 regras de dose, os
-- tetos do Anexo IV da IN 28, os pares de incompatibilidade e as regras de base.
--
-- POR QUE ISTO E UMA MIGRATION E NAO UM SCRIPT: o deploy roda migrations e mais nada. Qualquer
-- carga fora daqui vira um passo que alguem precisa lembrar de rodar — e producao nasceria com a
-- tela do magistral vazia. Os seeds antigos do repo (docs/emr/*.sql) sao carga manual porque
-- corrigem dado que ja existe; este e conteudo SEM O QUAL a funcionalidade nao existe.
--
-- Gerado da concatenacao de docs/emr/magistral-carga-prod.sh na ordem de dependencia: substancias
-- primeiro, depois norma e incompatibilidades, depois formulas, e as regras por ultimo, porque
-- cada camada referencia a anterior. Todo bloco e idempotente (WHERE NOT EXISTS / ON CONFLICT),
-- entao rodar de novo nao duplica.
--
-- Conferido de duas formas: carga do zero num banco vazio batendo com o desenvolvimento
-- (290 substancias, 132 formulas, 653 componentes, 54 regras, 167 faixas, 161 tetos, 12 pares,
-- 8 regras de base), e TRES passadas seguidas no mesmo banco devolvendo sempre os mesmos numeros.
--
-- A segunda conferencia achou um defeito real: um seed inseria "Aakg" e um arquivo seguinte
-- renomeava para "AAKG"; na segunda passada o insert recriava o nome antigo e a renomeacao
-- colidia. Os seeds passaram a inserir ja com o nome canonico — carga declarativa, sem etapa de
-- migracao de nome no meio.
--
-- BEGIN/COMMIT dos arquivos de origem foram removidos: goose ja envolve a migration inteira numa
-- transacao, e transacao aninhada quebraria a carga no meio.

-- ═══ magistral-seed-inicial.sql ═══
-- Semente inicial do catálogo magistral (dev e prod).
--
-- REGRA QUE VALE MAIS QUE O CONTEÚDO: aqui só entra o que tem fonte. Densidade aparente NÃO é
-- semeada — não existe tabela pública confiável, e ela varia por lote e por compactação da
-- farmácia (Anfarmag; Quallitá/Renylab). Enquanto a densidade não for cadastrada à mão, a
-- calculadora de cápsula se cala em vez de chutar, que é o comportamento correto.
--
-- O que entra são SINALIZADORES com fonte:
--   · formadores de mistura eutética: mentol, cânfora, timol, fenol, resorcina (sólido + sólido
--     vira líquido por abaixamento do ponto de fusão; contorna-se com pó adsorvente).
--   · talco adsorve cianocobalamina.
--
-- Fontes: Anfarmag, "Anfarmag dá dicas de como solucionar as incompatibilidades na manipulação
-- magistral" (anfarmag.org.br); Memento Terapêutico da Farmácia Universitária da UFRJ, 4ª ed.
-- 2021; Acofarma, "Instabilidade na formulação".
--
-- Idempotente: pode rodar de novo sem duplicar.

INSERT INTO public.magistral_components
    (id, name, synonyms, default_unit, eutectic_former, source, notes, last_review)
VALUES
    (uuidv7(), 'Mentol',    'levomentol, mentol racêmico', 'mg', true,  'seed',
     'Forma mistura eutética com cânfora, timol, fenol e resorcina. Contorna-se interpondo pó adsorvente (óxido/carbonato de magnésio, amido, sílica).', now()),
    (uuidv7(), 'Cânfora',   'canfora',                     'mg', true,  'seed',
     'Forma mistura eutética com mentol e timol.', now()),
    (uuidv7(), 'Timol',     '',                            'mg', true,  'seed',
     'Forma mistura eutética com mentol e cânfora.', now()),
    (uuidv7(), 'Fenol',     '',                            'mg', true,  'seed',
     'Forma mistura eutética com mentol.', now()),
    (uuidv7(), 'Resorcina', 'resorcinol',                  'mg', true,  'seed',
     'Forma mistura eutética com mentol.', now()),
    (uuidv7(), 'Talco',     '',                            'mg', false, 'seed',
     'Adsorve cianocobalamina; evitar como deslizante em fórmula que a contenha.', now()),
    (uuidv7(), 'Cianocobalamina', 'vitamina B12',           'mcg', false, 'seed',
     'Adsorvida pelo talco.', now())
ON CONFLICT DO NOTHING;

-- Par curado: talco × cianocobalamina.
INSERT INTO public.magistral_incompatibilities
    (id, component_a_id, component_b_id, severity, mechanism, note, source, last_review)
SELECT
    uuidv7(),
    LEAST(a.id::text, b.id::text)::uuid,
    GREATEST(a.id::text, b.id::text)::uuid,
    'warn',
    'o talco adsorve a cianocobalamina e reduz a dose disponível',
    'Trocar o deslizante (ex.: dióxido de silício) ou separar em fórmulas distintas.',
    'Anfarmag — incompatibilidades na manipulação magistral',
    now()
FROM public.magistral_components a, public.magistral_components b
WHERE lower(a.name) = 'talco' AND lower(b.name) = 'cianocobalamina'
ON CONFLICT DO NOTHING;

-- ═══ magistral-seed-funcional.sql ═══
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

-- ═══ magistral-indicacoes-externas.sql ═══
-- Indicações das substâncias que o RAG do consultório NÃO cobre.
--
-- O enriquecimento pelo RAG (cmd/magistral-enrich) deu indicação para 72 das 95 substâncias, com
-- trecho de aula guardado ao lado. As 23 restantes ficaram sem — ou porque são de uso
-- farmacotécnico (mentol, cânfora, timol, fenol, resorcina, talco, maltodextrina), ou porque o
-- material da pós não as discute. Estas vêm de pesquisa externa e entram como `suggested`: o
-- médico confere na tela de curadoria e o status vira `confirmed`.
--
-- Fontes: Tua Saúde, Nutritotal, Manuais MSD, informativos técnicos de insumo (Florien, Infinity,
-- Farmacam), Revista de Ciências Farmacêuticas Básica e Aplicada (Cimicifuga/Vitex),
-- Formulário Fitoterápico (Vitex agnus-castus), páginas de farmácias magistrais.

UPDATE public.magistral_components AS c SET
    indications = v.indications,
    dose_reference = COALESCE(c.dose_reference, v.dose_reference),
    evidence_status = CASE WHEN c.evidence_status = 'confirmed' THEN 'confirmed' ELSE 'suggested' END,
    last_review = now()
FROM (VALUES
    ('Ácido málico',
     'Intermediário do ciclo de Krebs, usado como sal do magnésio (dimalato) por participar da produção de energia celular. Associado a quadros de fadiga e dor muscular.',
     NULL),
    ('Amora',
     'Fonte de flavonoides e antocianinas com ação antioxidante, usada em sintomas do climatério, sobretudo fogachos, e como coadjuvante no controle lipídico.',
     NULL),
    ('Cimicifuga racemosa',
     'Fitoterápico de uso consagrado nos sintomas vasomotores do climatério, em especial fogachos e sudorese.',
     NULL),
    ('Vitex agnus-castus',
     'Fitoterápico com ação sobre a regulação hormonal feminina, usado em tensão pré-menstrual e irregularidade do ciclo.',
     NULL),
    ('Yam mexicano',
     'Fonte de diosgenina, sapogenina esteroidal precursora na síntese de esteroides, usada em queixas do climatério, dismenorreia e tensão pré-menstrual. A conversão endógena da diosgenina em progesterona é pequena, o que limita o que se pode esperar dela.',
     NULL),
    ('Isoflavona',
     'Fitoestrógeno da soja usado nos sintomas do climatério, principalmente fogachos.',
     '40 a 150 mg/dia, podendo ser fracionada em duas tomadas a critério médico.'),
    ('Valeriana',
     'Fitoterápico usado em insônia e ansiedade, geralmente à noite e associado a outros calmantes.',
     NULL),
    ('Passiflora',
     'Fitoterápico ansiolítico leve, usado em ansiedade e como coadjuvante do sono em associações.',
     NULL),
    ('FOS',
     'Fibra prebiótica não digerida no intestino delgado que serve de substrato para a microbiota do cólon; a fermentação gera ácidos graxos de cadeia curta, entre eles o butirato, ligado à saúde da mucosa.',
     NULL),
    ('Inulina',
     'Fibra prebiótica de comportamento semelhante ao FOS: chega íntegra ao cólon, alimenta bactérias benéficas e gera butirato na fermentação.',
     NULL),
    ('XOS',
     'Xilo-oligossacarídeo, prebiótico eficaz em quantidades menores que FOS e inulina, usado em suporte de microbiota.',
     NULL),
    ('MSM',
     'Enxofre orgânico usado em dor e rigidez articular (osteoartrite), recuperação muscular e suporte de tecido conjuntivo, pele, unhas e cabelo. Descrito com efeito anti-inflamatório por inibição de citocinas e aumento de glutationa.',
     NULL),
    ('Silício orgânico',
     'Forma biodisponível do silício, usada em pele, cabelos e unhas por favorecer a síntese de colágeno, elastina e ácido hialurônico; também citada em saúde óssea.',
     NULL),
    ('Sulfato de condroitina',
     'Usado em osteoartrite e dor articular, em geral associado à glicosamina.',
     NULL),
    ('Licopeno',
     'Carotenoide antioxidante do tomate, usado em saúde cardiovascular, próstata e fotoproteção da pele.',
     NULL),
    ('Magnésio taurato',
     'Magnésio ligado à taurina, com perfil voltado à saúde cardiovascular, incluindo regulação da pressão arterial.',
     NULL),
    ('Cloreto de magnésio',
     'Sal de magnésio de reposição, também citado em constipação. Higroscópico: em cápsula gelatinosa pede excipiente adsorvente ou cápsula vegetal.',
     NULL),
    -- Farmacotécnicos: a "indicação" aqui é o papel na fórmula, não uma indicação clínica.
    ('Mentol',
     'Uso farmacotécnico e tópico em associações. Forma mistura eutética com cânfora, timol, fenol e resorcina: juntas liquefazem na trituração, o que se contorna com pó adsorvente.',
     NULL),
    ('Cânfora',
     'Uso farmacotécnico e tópico em associações. Forma mistura eutética com mentol e timol.',
     NULL),
    ('Timol',
     'Uso farmacotécnico e tópico. Forma mistura eutética com mentol e cânfora.',
     NULL),
    ('Fenol',
     'Uso farmacotécnico e tópico. Forma mistura eutética com mentol.',
     NULL),
    ('Resorcina',
     'Uso farmacotécnico e tópico (dermatologia). Forma mistura eutética com mentol.',
     NULL),
    ('Talco',
     'Deslizante e adsorvente da manipulação. Adsorve cianocobalamina, o que desaconselha o uso em fórmula que a contenha.',
     NULL),
    ('Maltodextrina',
     'Diluente de sachês e cápsulas. Densidade aparente publicada de 0,70 g/mL, usada no cálculo volumétrico.',
     NULL)
) AS v(name, indications, dose_reference)
WHERE lower(c.name) = lower(v.name);

-- ═══ magistral-densidades-aproximadas.sql ═══
-- Densidades aparentes APROXIMADAS, para a calculadora de cápsula sair do silêncio.
--
-- A DECISÃO E O QUE ELA CUSTA: não existe tabela pública de densidade aparente por ativo, e o
-- valor real muda com lote, granulometria e compactação da farmácia (Anfarmag; Quallitá). Em vez
-- de esperar o número exato de cada insumo, entram aproximações POR CLASSE DE PÓ, declaradas como
-- tal em `density_source`. A calculadora já devolve faixa de ±25% e diz que é auxílio de
-- consultório; com estes valores ela passa a opinar, e o resultado continua sendo estimativa.
--
-- ÂNCORAS PUBLICADAS que calibram as classes:
--   · cápsula 0 (0,68 mL) comporta 400 a 450 mg de pó comum  → ~0,59 a 0,66 g/mL
--   · cápsula 00 (0,95 mL) comporta 500 a 600 mg             → ~0,53 a 0,63 g/mL
--   · a 0,70 g/mL, a cápsula 00 comporta ~665 mg             → confirma a ordem de grandeza
--   · creatina monoidratada ~0,55 g/mL · maltodextrina 0,70 · sulfato de condroitina 0,60
--
-- CLASSES USADAS (g/mL):
--   0,90  sal mineral inorgânico denso (óxido, carbonato, cloreto)
--   0,75  vitamina hidrossolúvel cristalina (C, complexo B)
--   0,65  mineral quelado, citrato, dimalato, treonato
--   0,60  aminoácido cristalino e derivado
--   0,55  fibra e polissacarídeo
--   0,50  probiótico veiculado
--   0,45  extrato seco vegetal, ativo lipossolúvel ou oleoso em pó, cristais aromáticos
--   0,35  colágeno e proteína hidrolisada
--
-- Reversível: `UPDATE magistral_components SET bulk_density=NULL, density_source=NULL
--             WHERE density_source = 'classe'` desfaz tudo o que este arquivo fez.
-- Idempotente: só preenche quem está sem densidade.

UPDATE public.magistral_components AS c
SET bulk_density = v.d, density_source = v.origem
FROM (VALUES
  -- medidas publicadas
  ('Creatina',0.55,'medida'), ('Maltodextrina',0.70,'medida'), ('Sulfato de condroitina',0.60,'medida'),

  -- minerais inorgânicos densos
  ('Cloreto de magnésio',0.90,'classe'), ('Cálcio',0.90,'classe'), ('Iodo',0.90,'classe'),

  -- minerais quelados e sais orgânicos
  ('Magnésio quelato',0.65,'classe'), ('Magnésio dimalato',0.65,'classe'),
  ('Magnésio L-treonato',0.65,'classe'), ('Magnésio taurato',0.65,'classe'),
  ('Zinco quelato',0.65,'classe'), ('Zinco carnosina',0.65,'classe'),
  ('Picolinato de cromo',0.65,'classe'), ('Selênio',0.65,'classe'),
  ('Cobre',0.65,'classe'), ('Ferro',0.65,'classe'), ('Silício orgânico',0.65,'classe'),

  -- vitaminas hidrossolúveis cristalinas
  ('Vitamina C',0.75,'classe'), ('Nicotinamida',0.75,'classe'), ('Riboflavina',0.75,'classe'),
  ('Piridoxal-5-fosfato',0.75,'classe'), ('Metilcobalamina',0.75,'classe'),
  ('Cianocobalamina',0.75,'classe'), ('Metilfolato',0.75,'classe'), ('Biotina',0.75,'classe'),
  ('Colina',0.75,'classe'),

  -- aminoácidos e derivados
  ('L-glutamina',0.60,'classe'), ('Glicina',0.60,'classe'), ('L-teanina',0.60,'classe'),
  ('L-taurina',0.60,'classe'), ('L-triptofano',0.60,'classe'), ('5-HTP',0.60,'classe'),
  ('L-carnitina',0.60,'classe'), ('Acetil-L-carnitina',0.60,'classe'),
  ('N-acetilcisteína',0.60,'classe'), ('GABA',0.60,'classe'), ('Mio-inositol',0.60,'classe'),
  ('D-quiro-inositol',0.60,'classe'), ('MSM',0.60,'classe'), ('Ácido málico',0.60,'classe'),
  ('Ácido alfa-lipoico',0.60,'classe'), ('Ácido R-alfa-lipoico',0.60,'classe'),
  ('Butirato de sódio',0.60,'classe'), ('Ácido hialurônico',0.60,'classe'),
  ('Alfa-GPC',0.60,'classe'), ('Citicolina',0.60,'classe'), ('NADH',0.60,'classe'),
  ('PQQ',0.60,'classe'), ('Melatonina',0.60,'classe'),

  -- fibras e polissacarídeos
  ('FOS',0.55,'classe'), ('Inulina',0.55,'classe'), ('XOS',0.55,'classe'),

  -- probióticos
  ('Probiótico multicepas',0.50,'classe'),

  -- extratos secos, lipossolúveis e cristais aromáticos
  ('Coenzima Q10',0.45,'classe'), ('Ubiquinol',0.45,'classe'), ('Vitamina D3',0.45,'classe'),
  ('Vitamina K2 MK-7',0.45,'classe'), ('Vitamina E',0.45,'classe'), ('Vitamina A',0.45,'classe'),
  ('Ômega-3',0.45,'classe'), ('Astaxantina',0.45,'classe'), ('Licopeno',0.45,'classe'),
  ('Curcumina',0.45,'classe'), ('Trans-resveratrol',0.45,'classe'), ('Quercetina',0.45,'classe'),
  ('Berberina',0.45,'classe'), ('Silimarina',0.45,'classe'), ('Extrato de chá verde',0.45,'classe'),
  ('Gymnema silvestre',0.45,'classe'), ('Ginseng brasileiro',0.45,'classe'),
  ('Valeriana',0.45,'classe'), ('Passiflora',0.45,'classe'), ('Amora',0.45,'classe'),
  ('Cimicifuga racemosa',0.45,'classe'), ('Vitex agnus-castus',0.45,'classe'),
  ('Ashwagandha',0.45,'classe'), ('Mucuna pruriens',0.45,'classe'), ('Yam mexicano',0.45,'classe'),
  ('Relora',0.45,'classe'), ('Tongkat ali',0.45,'classe'), ('Isoflavona',0.45,'classe'),
  ('Bacopa monnieri',0.45,'classe'), ('Ginkgo biloba',0.45,'classe'), ('Huperzina A',0.45,'classe'),
  ('Fosfatidilserina',0.45,'classe'), ('Fosfatidilcolina',0.45,'classe'),
  ('Teacrina',0.45,'classe'), ('Mentol',0.45,'classe'), ('Cânfora',0.45,'classe'),
  ('Timol',0.45,'classe'), ('Fenol',0.45,'classe'), ('Resorcina',0.45,'classe'),
  ('Talco',0.45,'classe'),

  -- proteínas hidrolisadas
  ('Colágeno Verisol',0.35,'classe'), ('Colágeno hidrolisado',0.35,'classe')
) AS v(nome, d, origem)
WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.nome))
  AND c.bulk_density IS NULL;

-- ═══ magistral-magnesio-formas.sql ═══
-- Formas de magnésio — tabela comparativa da Arboretum (lida por OCR) + checagem na literatura.
--
-- O QUE A TABELA RESOLVE: ela publica a EQUIVALÊNCIA ELEMENTAR de cada forma, que é o dado que
-- faltava para o fator de correção funcionar no magnésio, a família mais prescrita do catálogo.
-- Dois valores vêm explícitos da tabela:
--   · treonato: 1.500 a 2.000 mg do sal equivalem a 117 a 156 mg de magnésio elementar (~7,8%);
--   · inositol: 500 a 1.000 mg do sal equivalem a 150 a 300 mg de elementar (~30%).
-- Para as demais formas a tabela expressa a dose JÁ em magnésio elementar ("50 a 500 mg de
-- magnésio como malato"), que é a convenção de prescrição que ela adota.
--
-- ACHADO QUE MUDA DOSE: o catálogo tinha magnésio L-treonato com dose usual de 300 mg. Os ensaios
-- clínicos de cognição e sono usam 2 g/dia do SAL (Magtein), e a tabela da Arboretum diz o mesmo:
-- 1.500 a 2.000 mg. Trezentos miligramas do sal entregam cerca de 23 mg de magnésio elementar —
-- muito abaixo do que foi testado. A faixa foi corrigida, e as fórmulas que usam 300 mg passam a
-- ser sinalizadas pela própria tela como abaixo da faixa.
--
-- Percentuais estequiométricos (óxido, citrato, cloreto, sulfato, carbonato) entram anotados como
-- tal: são cálculo de fórmula molecular, não medida de lote.

-- ---------- formas que faltavam no catálogo ----------
INSERT INTO public.magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
   elemental_percent, correction_note, indications, indication_bullets, dose_reference, notes,
   source, evidence_status, last_review)
SELECT * FROM (VALUES
  (uuidv7(),'Magnésio citrato','citrato de magnésio','mg',300::numeric,50::numeric,500::numeric,
   0.65::numeric,'classe',16.2::numeric,
   'Citrato tem cerca de 16% de magnésio elementar (estequiometria). A tabela da Arboretum prescreve a dose já em elementar.',
   'Alta biodisponibilidade oral, com efeito laxativo em parte dos pacientes. Usado em quadros cardiovasculares, neuromusculares e nervosos.',
   E'reposição de magnésio\nconstipação, pelo efeito laxativo',
   'Suplementação: 50 a 500 mg de magnésio (como citrato) ao dia. Como laxante: 10 a 30 g do sal.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio aspartato','aspartato de magnésio','mg',300::numeric,200::numeric,400::numeric,
   0.65::numeric,'classe',NULL,NULL,
   'Forma associada ao ácido aspártico, indicada em cardiopatias, fadiga crônica e estresse físico e mental.',
   E'fadiga crônica\ncardiopatias\nestresse físico e mental',
   '200 a 400 mg de magnésio (como aspartato) ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio ascorbato','ascorbato de magnésio','mg',300::numeric,50::numeric,500::numeric,
   0.65::numeric,'classe',NULL,NULL,
   'Une magnésio e vitamina C na mesma molécula, com uso voltado a suporte imune e cardíaco.',
   E'suporte imune\nsuporte cardiovascular',
   '50 a 500 mg de magnésio (como ascorbato) ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio inositol','magnésio com inositol','mg',750::numeric,500::numeric,1000::numeric,
   0.60::numeric,'classe',30::numeric,
   'A tabela publica a equivalência: 500 a 1.000 mg do sal correspondem a 150 a 300 mg de magnésio elementar.',
   'Une o magnésio quelado ao inositol, com uso em relaxamento mental e corporal, qualidade do sono, memória e dores musculares.',
   E'sono e relaxamento\nmemória\ndores musculares e ósseas',
   '500 a 1.000 mg do sal ao dia (150 a 300 mg de magnésio elementar).',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio óxido','óxido de magnésio','mg',600::numeric,400::numeric,800::numeric,
   0.90::numeric,'classe',60.3::numeric,
   'Óxido tem cerca de 60% de magnésio elementar (estequiometria), mas é a forma menos absorvida.',
   'Forma menos absorvida, usada sobretudo pelo efeito laxativo.',
   E'constipação\nreposição de baixo custo, com absorção limitada',
   '400 a 800 mg do sal ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio carbonato','carbonato de magnésio','mg',500::numeric,250::numeric,1000::numeric,
   0.90::numeric,'classe',28.8::numeric,
   'Carbonato tem cerca de 29% de magnésio elementar (estequiometria).',
   'Propriedade antiácida, também usado pelo efeito laxativo em doses maiores.',
   E'antiácido\nlaxativo em dose maior',
   'Antiácido: 250 a 1.000 mg do sal. Laxante: 2 a 5 g do sal ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio sulfato','sulfato de magnésio, sal amargo','mg',NULL::numeric,NULL::numeric,NULL::numeric,
   0.90::numeric,'classe',9.9::numeric,
   'Sulfato heptaidratado tem cerca de 10% de magnésio elementar (estequiometria).',
   'Boa absorção tópica; por via oral tem efeito laxativo.',
   E'uso tópico transdérmico\nlaxativo por via oral',
   'Oral: 5 a 10 g do sal ao dia. Tópico: 50 a 200 mg do sal em creme transdérmico.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now())
) AS v(id,name,synonyms,default_unit,usual_dose,min_dose,max_dose,bulk_density,density_source,
       elemental_percent,correction_note,indications,indication_bullets,dose_reference,notes,
       source,evidence_status,last_review)
WHERE NOT EXISTS (
  SELECT 1 FROM public.magistral_components c
  WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.name))
);

-- ---------- correções nas formas que já existiam ----------
-- Treonato: a faixa estava uma ordem de grandeza abaixo do que os ensaios usam.
UPDATE public.magistral_components SET
  usual_dose = 2000, min_dose = 1500, max_dose = 2000,
  elemental_percent = 7.8,
  correction_note = 'Treonato tem cerca de 7,8% de magnésio elementar: 2 g do sal equivalem a cerca de 156 mg de elementar (tabela Arboretum).',
  dose_reference = '1.500 a 2.000 mg do sal ao dia, em duas tomadas. Os ensaios de cognição e sono com Magtein usam 2 g/dia; a tabela da Arboretum diz o mesmo. Doses de 200 a 500 mg do sal, comuns em formulários, entregam menos de 40 mg de magnésio elementar.',
  notes = coalesce(notes,'') || ' Faixa corrigida a partir da tabela comparativa Arboretum e dos ensaios com Magtein (2 g/dia).',
  last_review = now()
WHERE lower(name) = 'magnésio l-treonato';

UPDATE public.magistral_components SET
  elemental_percent = 30,
  correction_note = 'Bisglicinato com cerca de 30% de magnésio elementar (confirmar no laudo do lote). A tabela da Arboretum prescreve a dose já em elementar: 50 a 500 mg/dia.'
WHERE lower(name) = 'magnésio quelato';

UPDATE public.magistral_components SET
  correction_note = 'A tabela da Arboretum prescreve a dose já em magnésio elementar: 50 a 500 mg/dia.',
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: 50 a 500 mg de magnésio (como malato) ao dia.'
WHERE lower(name) = 'magnésio dimalato' AND correction_note IS NULL;

UPDATE public.magistral_components SET
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: 50 a 500 mg de magnésio (como taurato) ao dia, com uso voltado à saúde cardiovascular e ao diabetes tipo 2.',
  usual_dose = coalesce(usual_dose, 300), min_dose = coalesce(min_dose, 50), max_dose = coalesce(max_dose, 500)
WHERE lower(name) = 'magnésio taurato';

UPDATE public.magistral_components SET
  elemental_percent = 11.9,
  correction_note = 'Cloreto hexaidratado tem cerca de 12% de magnésio elementar (estequiometria). A tabela indica sobretudo uso tópico, a 15%.',
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: melhor absorção tópica, 15% em formulação de uso externo.'
WHERE lower(name) = 'cloreto de magnésio';

-- ═══ magistral-formas-preferidas-e-correcao.sql ═══
-- Formas preferidas do Dr. Getúlio, ativos de marca das farmácias parceiras e fatores de correção.
--
-- FORMA PREFERIDA: quando ele precisa dessas vitaminas e minerais, usa sempre a forma ativa. O
-- catálogo passa a saber disso e sugere a troca ao montar a fórmula. A nota de cada par traz o
-- que a literatura sustenta — inclusive quando ela NÃO sustenta, que é o caso do palmitato de
-- ascorbila por via oral.
--
-- FATOR DE CORREÇÃO: das fichas técnicas. Selenometionina é diluída a 1% em maltodextrina, e o
-- magnésio bisglicinato tem 30% de mineral elementar; prescrever "300 mg" sem dizer se é elemento
-- ou insumo muda a fórmula em três vezes.

-- ---------- 1. novas substâncias (formas preferidas e ativos de marca) ----------
INSERT INTO public.magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
   brand, elemental_percent, correction_note, indications, indication_bullets, dose_reference,
   notes, source, evidence_status, last_review)
SELECT * FROM (VALUES
  (uuidv7(), 'Palmitato de ascorbila', 'ascorbil palmitato, vitamina C lipossolúvel', 'mg',
   500::numeric, 100::numeric, 1000::numeric, 0.55::numeric, 'classe', NULL, NULL, NULL,
   'Forma lipossolúvel da vitamina C. Por via oral é hidrolisada no trato digestivo antes da absorção, e o ascorbato liberado tem a mesma biodisponibilidade do ácido ascórbico; a vantagem real está na estabilidade da fórmula e no uso tópico.',
   E'fonte de vitamina C em fórmula com lipossolúveis\nuso tópico e estabilidade da formulação',
   'Dose equivalente à de vitamina C. Cerca de 43% do peso é ascorbato.',
   'Forma preferida do prescritor. Ressalva de literatura: Linus Pauling Institute descreve hidrólise pré-absortiva, sem ganho de biodisponibilidade oral sobre o ácido ascórbico.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Selenometionina', 'selenometionina 1%, L-selenometionina', 'mcg',
   100::numeric, 50::numeric, 200::numeric, 0.60::numeric, 'classe', NULL, 1::numeric,
   'Insumo diluído a 1% em maltodextrina: 100 mcg de selênio elementar equivalem a 10 mg do insumo. Aplicar fator de correção quando a dose prescrita for do elemento.',
   'Forma orgânica do selênio, mais biodisponível e mais retida em tecidos que o selenito de sódio. Antioxidante via glutationa peroxidase, com sinergia clássica com vitamina E.',
   E'reposição de selênio\nsuporte antioxidante e tireoidiano',
   '50 a 200 mcg/dia de selênio elementar (ficha técnica: até 319,75 mcg).',
   'Forma preferida do prescritor. Literatura sustenta: selenometionina é cerca de duas vezes mais biodisponível que o selenito.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'CavaQ10', 'coenzima Q10 gama-ciclodextrina, ubiquinona ciclodextrina', 'mg',
   100::numeric, 30::numeric, 250::numeric, 0.50::numeric, 'classe', 'CavaQ10®', NULL, NULL,
   'Coenzima Q10 (ubiquinona) complexada em gama-ciclodextrina, dispersível em água e mais estável que a CoQ10 comum.',
   E'suporte mitocondrial\nuso de estatina\nsaúde cardiovascular',
   '30 a 250 mg/dia (ficha técnica).',
   'Forma preferida do prescritor. Literatura: estudo humano cruzado com 22 voluntários mostra AUC e Cmax maiores para o complexo com gama-ciclodextrina; o ganho publicado é da ordem de 35% sobre CoQ10 com celulose. O "1800% mais biodisponível" da lâmina é claim do fabricante, sem correspondência na literatura indexada.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Morosil', 'extrato de laranja moro, Citrus sinensis Moro', 'mg',
   400::numeric, 400::numeric, 400::numeric, 0.45::numeric, 'classe', 'Morosil®', NULL, NULL,
   'Extrato padronizado de laranja moro em antocianinas, usado em redução de gordura corporal e circunferência abdominal, associado a dieta e exercício.',
   E'gordura visceral e abdominal\ncoadjuvante do emagrecimento',
   '400 mg/dia, dose usada nos ensaios.',
   'Literatura sustenta: RCT de 24 semanas com 180 participantes e meta-análise de 3 RCTs (2025) com redução de peso, IMC, circunferências e massa gorda.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Açafrão padronizado', 'Saffrin, Satiereal, Crocus sativus', 'mg',
   176::numeric, 88::numeric, 200::numeric, 0.45::numeric, 'classe', 'Saffrin®', NULL, NULL,
   'Extrato padronizado de Crocus sativus usado para reduzir beliscos e desejo por doces entre refeições.',
   E'compulsão alimentar e beliscos\ncontrole de apetite',
   '176 a 200 mg/dia nos ensaios; formulário usa 160 mg.',
   'Literatura sustenta parcialmente: RCT de 8 semanas com 60 mulheres com sobrepeso leve mostra menos beliscos e maior saciedade; efeito modesto e em população específica.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Cactinea', 'Opuntia ficus-indica, figo-da-índia', 'mg',
   1000::numeric, 500::numeric, 1000::numeric, 0.45::numeric, 'classe', 'Cactinea™', NULL, NULL,
   'Extrato de figo-da-índia usado com proposta diurética e de redução de retenção hídrica.',
   E'retenção hídrica\ncoadjuvante em celulite',
   '500 a 1.000 mg/dia (formulário).',
   'Evidência fraca: as alegações diuréticas vêm sobretudo de material do fabricante; a literatura indexada é escassa para o desfecho proposto.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Glycoxil', 'carcinina, peptidomimético de carnosina', 'mg',
   200::numeric, 50::numeric, 300::numeric, 0.45::numeric, 'classe', 'Glycoxil®', NULL, NULL,
   'Peptidomimético da carnosina (carcinina) com ação antiglicante e antiglicoxidante, usado em fórmulas de pele e antienvelhecimento.',
   E'glicação cutânea\nfórmulas antienvelhecimento',
   '50 a 300 mg/dia (formulário e material técnico).',
   'Evidência majoritariamente do fabricante (Exsymol), in vitro e in vivo de pequeno porte.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Bio Arct', 'Kappaphycus alvarezii, alga vermelha', 'mg',
   100::numeric, 100::numeric, 200::numeric, 0.45::numeric, 'classe', 'Bio Arct®', NULL, NULL,
   'Extrato de alga vermelha usado em fórmulas de pele com proposta de suporte energético celular e firmeza.',
   E'firmeza e qualidade de pele\nassociado a silício e colágeno',
   '100 a 200 mg/dia (formulário).',
   'Evidência escassa para uso oral: o que a busca encontra é sobretudo aplicação cosmética tópica e uso agrícola da alga.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Koubo', 'Cereus sylvestrii', 'mg',
   400::numeric, 300::numeric, 500::numeric, 0.45::numeric, 'classe', 'Koubo®', NULL, NULL,
   'Extrato de cacto usado com proposta de saciedade e controle de compulsão.',
   E'compulsão alimentar\nsaciedade',
   '300 a 500 mg/dia (formulário).',
   'Evidência escassa na literatura indexada; sustentação principal é material do fabricante.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Green coffee', 'café verde, ácido clorogênico', 'mg',
   30::numeric, 30::numeric, 400::numeric, 0.45::numeric, 'classe', NULL, NULL, NULL,
   'Extrato de café verde, fonte de ácido clorogênico, usado em fórmulas de pele e de metabolismo.',
   E'antioxidante\ncoadjuvante metabólico',
   '30 mg em fórmulas de pele; doses metabólicas chegam a 400 mg.',
   'Evidência modesta e heterogênea para desfecho de peso.',
   'parceiro', 'suggested', now())
) AS v(id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
       brand, elemental_percent, correction_note, indications, indication_bullets, dose_reference,
       notes, source, evidence_status, last_review)
WHERE NOT EXISTS (
  SELECT 1 FROM public.magistral_components c
  WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.name))
);

-- ---------- 2. fator de correção nas substâncias que já existiam ----------
UPDATE public.magistral_components SET
  elemental_percent = 30,
  correction_note = 'Bisglicinato com cerca de 30% de magnésio elementar (confirmar no laudo do lote). 300 mg de magnésio elementar equivalem a cerca de 1 g do quelato.'
WHERE lower(name) = 'magnésio quelato' AND elemental_percent IS NULL;

UPDATE public.magistral_components SET
  correction_note = 'Mineral quelado: se a dose prescrita for do elemento, a farmácia aplica o fator do laudo do lote. Diga na fórmula se a dose é do elemento ou do insumo.'
WHERE lower(name) IN ('zinco quelato','ferro','cobre','picolinato de cromo','cálcio')
  AND correction_note IS NULL;

-- ---------- 3. formas preferidas ----------
UPDATE public.magistral_components AS c SET
  preferred_alternative_id = p.id,
  preference_note = v.nota
FROM (VALUES
  ('Vitamina C','Palmitato de ascorbila',
   'A literatura não mostra ganho de biodisponibilidade oral sobre o ácido ascórbico, mas ele estabiliza a fórmula com lipossolúveis.'),
  ('Selênio','Selenometionina',
   'A literatura sustenta: cerca de duas vezes mais biodisponível que o selenito.'),
  ('Coenzima Q10','CavaQ10',
   'O complexo com gama-ciclodextrina tem estudo humano com AUC maior; o ganho publicado é bem menor que o do material promocional.'),
  ('Cianocobalamina','Metilcobalamina',
   'A evidência favorece, sem ser conclusiva para reposição oral simples; o caso mais forte é neuropatia diabética.')
) AS v(generico, preferida, nota)
JOIN public.magistral_components p
  ON lower(public.immutable_unaccent(p.name)) = lower(public.immutable_unaccent(v.preferida))
WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.generico));

-- ═══ magistral-substancias-formulario.sql ═══
-- Substâncias do formulário das parceiras que faltavam no catálogo.
-- Nome canônico, unidade e densidade por classe. NENHUMA faixa de dose entra aqui: dose
-- vinda da própria fórmula é o que fazia o catálogo conferir a fórmula contra ela mesma.

INSERT INTO magistral_components (id, name, synonyms, default_unit, bulk_density, density_source,
  eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, sachet_ok,
  source, evidence_status, notes, is_active, created_at, updated_at) VALUES
  (uuid_generate_v7(), 'AAKG', 'arginina alfa-cetoglutarato', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Abacateiro', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Alfa-amilase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Altilix', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ásiaticosídeo', 'ASIATICOSIDE, asiaticoside', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Astragalus', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'BCAA', 'aminoácidos de cadeia ramificada', 'g', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Beta-alanina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Betacaroteno', 'BETACAROTENO', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como vitamina; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Betaína anidra', 'BETAÍNA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium bifidum', 'BIFIDOB BIFIDUM', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium breve', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium longum', 'BIFIDOB LONGUM', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Boro', 'BORO QUELATO', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como mineral; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Boswellia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bromelina', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Buclizina', 'BUCLISINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Capsiate', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cássia angustifólia', 'sene', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Castanha-da-índia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cavalinha', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Centella asiatica', 'CENTELA ASIATICA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Chia', 'CHIA', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Chlorella', 'CLORELLA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ciproeptadina', 'CIPROHEPTADINE', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cissus quadrangularis', 'CISSUS QUADRANGULARIS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cisteína', 'CISTÉINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Citrus aurantium', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Coleus forskohlii', 'COLEUS FORSKOHLLI', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Colágeno tipo II', 'COLAGENO TIPO II, UC II', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cranberry', 'CRAMBERRY EXT SECO 25%', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cyanotis vagas', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'DMAE bitartarato', 'DMAE BITARTARATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Dimpless', 'SOD DIMPLESS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Diosmina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Dong quai', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Enzimas pancreáticas', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Epicor', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Equinácea', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Extrato de semente de uva', 'EXT SEMENTE UVA, VITIS VINIFERA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Faseolamina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Folha de oliveira', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Fosfolipídeos de caviar', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Gama-orizanol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Gengibre', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ginseng', 'EXT SECO DE GINSENG', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Glisodim', 'GLISODIN', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Glutationa reduzida', 'GLUTATIONA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Griffonia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'HMB', 'HMB CALCIO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Hesperidina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Hibisco', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'ID-alG', 'ID ALG, id alg', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Kava-kava', 'KAWA-KAWA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-arginina', 'L ARGININA, L-ARGININA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-citrulina malato', 'L CITRULINA MALATO', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-fenilalanina', 'L FENILALANINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-isoleucina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-leucina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-lisina', 'L-LISINA, LISINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-ornitina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-prolina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-valina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus acidophilus', 'LACTOB ACIDOPHILLUS', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus bulgaricus', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus casei', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus delbrueckii', 'LACTOB DELBRUECKII', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus gasseri', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus plantarum', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus reuteri', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus rhamnosus', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus salivarius', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Streptococcus thermophilus', 'LACTOB THERMOPHILUS', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lecitina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lipase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lowat', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Luteína', 'LUTEÍNA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Maca peruana', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Manganês', 'MANGANÊS QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Melissa', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Meratrim', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Molibdênio', 'MOLIBEDÊNIO QUELATO', 'mcg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Mulungu', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Nattoquinase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Nutricolin', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Oli-Ola', 'OLI OLA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Palatinose', '', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Papaína', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Phosfator', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Picnogenol', 'PICNOGENOOL', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Polidextrose', '', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Polypodium leucotomos', 'POLYPODIUM LEUCOTOMOS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Romã', 'POMEGRANATE, pomegranate, Punica granatum', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Potássio', 'CLORETO DE POTÁSSIO, POTÁSSIO QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Pregnenolona', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Protease', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Psyllium', 'PLANTAGO OVATE, PSILLYUM', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Pygeum africanum', 'PYGEUM AFRICANUM', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Trevo-vermelho', 'RED CLOVER, red clover, Trifolium pratense', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Rutina', 'RUTINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Sinetrol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Sucupira', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de alcachofra', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de alecrim', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de espinheira-santa', 'TINT ESPINHEIRA SANTA', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de funcho', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de hortelã', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tribulus terrestris', 'TRIBULUS TERRESTRE', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vanádio', 'VANADIO QUELADO, VANÁDIO QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Veiculo oleoso qsp', '', 'Gotas', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Verisol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vinpocetina', 'VIMPOCETINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vitamina B1', 'TIAMINA, VIT B1', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como vitamina; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ácido D-aspártico', 'AC D ASPARTICO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ácido pantotênico', 'PANTOTENATO DE CÁLCIO, VIT B5', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now())
ON CONFLICT DO NOTHING;

UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AMORA VERDE EXT SECO')
 WHERE name = 'Amora';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ASWAGANDHA')
 WHERE name = 'Ashwagandha';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', COBRE QUELADO, COBRE QUELATO')
 WHERE name = 'Cobre';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', COE Q10, COENZIMA Q0')
 WHERE name = 'Coenzima Q10';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', CALCIO QUELADO, CÁLCIO CITRATO, CÁLCIO QUELATO')
 WHERE name = 'Cálcio';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', D CHIRO INOSITOL')
 WHERE name = 'D-quiro-inositol';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', D RIBOSE')
 WHERE name = 'D-ribose';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ACTIVE EGCG, CHA VERDE')
 WHERE name = 'Extrato de chá verde';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', FOSFATIDIL COLINA')
 WHERE name = 'Fosfatidilcolina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L-GLICINA')
 WHERE name = 'Glicina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GREEN COFFE')
 WHERE name = 'Green coffee';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GINMENA SILVESTRE, GYMENA SILVESTRE')
 WHERE name = 'Gymnema silvestre';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L GLUTAMINA')
 WHERE name = 'L-glutamina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L THEANINA, L THEANINE, THEANINA')
 WHERE name = 'L-teanina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L TIROSINA')
 WHERE name = 'L-tirosina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', MAGNESIO QUELADO, MAGNÉSIO')
 WHERE name = 'Magnésio quelato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT B12')
 WHERE name = 'Metilcobalamina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AC FOLICO, LEVOFOLIC ACID, ÁC FOLICO, ÁC FÓLICO, ÁC. FOLICO')
 WHERE name = 'Metilfolato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', MIO INOSITOL')
 WHERE name = 'Mio-inositol';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', N ACETIL CISTEINA')
 WHERE name = 'N-acetilcisteína';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', NIACINA, VIT B3')
 WHERE name = 'Nicotinamida';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT C')
 WHERE name = 'Palmitato de ascorbila';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', CROMO QUELATO, PICOLINATO CROMO')
 WHERE name = 'Picolinato de cromo';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', PIRIDOXAL 5 FOSFATO, PIRIDOXAL 5-FOSFATO, PIRIDOXAL FOSFATO, VIT B6, VITB6')
 WHERE name = 'Piridoxal-5-fosfato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT B2')
 WHERE name = 'Riboflavina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', EXSELEN, SELENIO QUELADO, SELÊNIO QUELATO')
 WHERE name = 'Selenometionina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', TEACRINE')
 WHERE name = 'Teacrina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT A')
 WHERE name = 'Vitamina A';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT D, VIT D 3, VIT D3')
 WHERE name = 'Vitamina D3';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT E')
 WHERE name = 'Vitamina E';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT K, VIT K2 MK7')
 WHERE name = 'Vitamina K2 MK-7';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ZINCO, ZINCO QUELADO')
 WHERE name = 'Zinco quelato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AC ALFA LIPOICO, ALFA LIPOICO, ÁC ALFA LIPOICO')
 WHERE name = 'Ácido alfa-lipoico';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GARCINEA CAMBOJA, GARCINIA CAMBOJA')
 WHERE name = 'Ácido hidroxicítrico';

-- Segunda passada do parser: componentes que só apareceram depois de aceitar a dose separada por
-- um espaço só ("STREPTOCOCCUS THERMOPHILUS 1 BLH") e o prefixo "POSOLOGIA:" na linha de uso.

INSERT INTO magistral_components (id, name, synonyms, default_unit, bulk_density, density_source,
  sachet_ok, source, evidence_status, notes, is_active, created_at, updated_at) VALUES
 (uuid_generate_v7(), 'Bifidobacterium lactis', 'BIFIDOBACTERIUM LACTIS', 'bilhões UFC', 0.45, 'classe', true,
  'parceiro', 'pending', 'Probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Lactobacillus paracasei', 'LACTOBACILLUS PARACASEI', 'bilhões UFC', 0.45, 'classe', true,
  'parceiro', 'pending', 'Probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Polietilenoglicol 4000', 'PEG 4000, macrogol 4000', 'g', 0.60, 'classe', true,
  'parceiro', 'pending', 'Laxante osmótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Ganoderma lucidum', 'WEG LEM 70, reishi, cogumelo do sol', 'mg', 0.45, 'classe', true,
  'parceiro', 'pending', 'Fitoterápico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now())
ON CONFLICT DO NOTHING;

-- ═══ magistral-substancias-nomes.sql ═══
-- A guarda "AND NOT EXISTS" em cada renomeação existe porque este arquivo roda depois de um seed
-- que insere os nomes ANTIGOS: numa segunda passada o insert recria "Aakg" e a renomeação para
-- "AAKG" colidia com a linha que já estava lá. Com a guarda, a segunda passada não faz nada.
--
-- Ajuste dos nomes gerados a partir do formulário: o formulário escreve tudo em caixa alta e
-- abreviado, e a capitalização automática produz "Bcaa" e "Lactob plantarum". Aqui os nomes ficam
-- como se escreve de verdade, e as siglas voltam a ser siglas.
--
-- O veículo sai do catálogo de ativos: "veículo oleoso qsp" é base, não substância.

UPDATE magistral_components SET name = 'AAKG', synonyms = 'arginina alfa-cetoglutarato' WHERE name = 'Aakg'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'AAKG');
UPDATE magistral_components SET name = 'BCAA', synonyms = 'aminoácidos de cadeia ramificada' WHERE name = 'Bcaa'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'BCAA');
UPDATE magistral_components SET name = 'ID-alG', synonyms = 'ID ALG, id alg' WHERE name = 'Id alg'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'ID-alG');
UPDATE magistral_components SET name = 'Oli-Ola', synonyms = 'OLI OLA' WHERE name = 'Oli ola'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Oli-Ola');
UPDATE magistral_components SET name = 'Beta-alanina' WHERE name = 'Beta alanina'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Beta-alanina');
UPDATE magistral_components SET name = 'Alfa-amilase' WHERE name = 'Alfa amilase'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Alfa-amilase');
UPDATE magistral_components SET name = 'Gama-orizanol' WHERE name = 'Gama orizanol'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Gama-orizanol');
UPDATE magistral_components SET name = 'Castanha-da-índia' WHERE name = 'Castanha da india'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Castanha-da-índia');
UPDATE magistral_components SET name = 'Fosfolipídeos de caviar' WHERE name = 'Fosfolipideos de caviar'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Fosfolipídeos de caviar');
UPDATE magistral_components SET name = 'Papaína' WHERE name = 'Papaina'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Papaína');
UPDATE magistral_components SET name = 'Equinácea' WHERE name = 'Equinacea'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Equinácea');
UPDATE magistral_components SET name = 'Ásiaticosídeo', synonyms = 'ASIATICOSIDE, asiaticoside' WHERE name = 'Asiaticoside'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Ásiaticosídeo');
UPDATE magistral_components SET name = 'Cássia angustifólia', synonyms = 'sene' WHERE name = 'Cassia angustifolia'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Cássia angustifólia');
UPDATE magistral_components SET name = 'Romã', synonyms = 'POMEGRANATE, pomegranate, Punica granatum' WHERE name = 'Pomegranate'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Romã');
UPDATE magistral_components SET name = 'Trevo-vermelho', synonyms = 'RED CLOVER, red clover, Trifolium pratense' WHERE name = 'Red clover'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Trevo-vermelho');

-- probióticos com o gênero por extenso
UPDATE magistral_components SET name = 'Lactobacillus acidophilus', synonyms = 'LACTOB ACIDOPHILLUS' WHERE name = 'Lactob acidophillus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus acidophilus');
UPDATE magistral_components SET name = 'Lactobacillus bulgaricus',  synonyms = 'LACTOB BULGARICUS'  WHERE name = 'Lactob bulgaricus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus bulgaricus');
UPDATE magistral_components SET name = 'Lactobacillus casei',       synonyms = 'LACTOB CASEI'       WHERE name = 'Lactob casei'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus casei');
UPDATE magistral_components SET name = 'Lactobacillus delbrueckii', synonyms = 'LACTOB DELBRUECKII' WHERE name = 'Lactob delbrueckii'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus delbrueckii');
UPDATE magistral_components SET name = 'Lactobacillus gasseri',     synonyms = 'LACTOB GASSERI'     WHERE name = 'Lactob gasseri'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus gasseri');
UPDATE magistral_components SET name = 'Lactobacillus plantarum',   synonyms = 'LACTOB PLANTARUM'   WHERE name = 'Lactob plantarum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus plantarum');
UPDATE magistral_components SET name = 'Lactobacillus reuteri',     synonyms = 'LACTOB REUTERI'     WHERE name = 'Lactob reuteri'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus reuteri');
UPDATE magistral_components SET name = 'Lactobacillus rhamnosus',   synonyms = 'LACTOB RHAMNOSUS'   WHERE name = 'Lactob rhamnosus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus rhamnosus');
UPDATE magistral_components SET name = 'Lactobacillus salivarius',  synonyms = 'LACTOB SALIVARUS'   WHERE name = 'Lactob salivarus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus salivarius');
UPDATE magistral_components SET name = 'Streptococcus thermophilus', synonyms = 'LACTOB THERMOPHILUS' WHERE name = 'Lactob thermophilus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Streptococcus thermophilus');
UPDATE magistral_components SET name = 'Bifidobacterium bifidum', synonyms = 'BIFIDOB BIFIDUM' WHERE name = 'Bifidob bifidum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium bifidum');
UPDATE magistral_components SET name = 'Bifidobacterium breve',  synonyms = 'BIFIDOB BREVE'  WHERE name = 'Bifidob breve'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium breve');
UPDATE magistral_components SET name = 'Bifidobacterium longum', synonyms = 'BIFIDOB LONGUM' WHERE name = 'Bifidob longum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium longum');

-- tinturas
UPDATE magistral_components SET name = 'Tintura de alcachofra',       synonyms = 'TINT ALCACHOFRA'       WHERE name = 'Tint alcachofra'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de alcachofra');
UPDATE magistral_components SET name = 'Tintura de alecrim',          synonyms = 'TINT ALECRIM'          WHERE name = 'Tint alecrim'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de alecrim');
UPDATE magistral_components SET name = 'Tintura de espinheira-santa', synonyms = 'TINT ESPINHEIRA SANTA' WHERE name = 'Tint espinheira santa'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de espinheira-santa');
UPDATE magistral_components SET name = 'Tintura de funcho',           synonyms = 'TINT FUNCHO'           WHERE name = 'Tint funcho'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de funcho');
UPDATE magistral_components SET name = 'Tintura de hortelã',          synonyms = 'TINT HORTELA'          WHERE name = 'Tint hortela'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de hortelã');

-- veículo não é ativo
DELETE FROM magistral_components WHERE name = 'Veiculo oleoso qsp';

-- Os nomes acima mudaram depois que as fórmulas já apontavam para a grafia antiga: os sinônimos
-- mantêm a ligação, senão a fórmula perde densidade e o cálculo de cápsula para de sair.

UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', alfa amilase')       WHERE name = 'Alfa-amilase';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', beta alanina')       WHERE name = 'Beta-alanina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', castanha da india') WHERE name = 'Castanha-da-índia';

-- ═══ magistral-glp1.sql ═══
-- Suporte durante o uso de análogos de GLP-1 — material da Arboretum.
--
-- É o melhor documento do lote: cada substância vem com FAIXA DE DOSE e REFERÊNCIA, o que resolve
-- o problema que o formulário criou (faixa numérica sem base). As faixas abaixo são as do
-- material, com a citação guardada no texto de posologia.
--
-- Três ressalvas minhas, que o material não traz, entram como observação da substância.

-- ---------------------------------------------------------------------------------------------
-- 1. Substâncias novas, com a faixa e a referência do material
-- ---------------------------------------------------------------------------------------------
INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, sachet_ok, bitterness, source, evidence_status, indications, dose_reference,
   is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Psyllium', 'PLANTAGO OVATE, PSILLYUM, Plantago ovata, psilium', 'g',
  5, 3, 20, 'por_dia', 0.55, 'classe', true, 1, 'parceiro', 'suggested',
  'Fibra que absorve água no intestino, aumenta o volume do bolo fecal e estimula o trânsito. Nos análogos de GLP-1 entra pela constipação por retardo do esvaziamento gástrico.',
  'Até 20 g/dia (McRorie & McKeown, 2017). Tomar com bastante líquido; sem água suficiente a fibra piora a obstipação.', true, now(), now()),

 (uuid_generate_v7(), 'Saccharomyces boulardii', 'S. boulardii, levedura probiótica', 'mg',
  250, 150, 500, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Levedura probiótica que reduz episódios e gravidade de diarreia e ajuda no equilíbrio da microbiota.',
  'De 150 a 500 mg/dia (Pal et al., 2020).', true, now(), now()),

 (uuid_generate_v7(), 'Bacillus clausii', 'B. clausii', 'bilhões UFC',
  2, 0.1, 10, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Probiótico esporulado, com efeito na diarreia aguda e na diversidade da microbiota.',
  'De 0,1 a 10 bilhões de UFC/dia (Ianiro et al., 2018).', true, now(), now()),

 (uuid_generate_v7(), 'Simeticone', 'simeticona, dimeticona ativada', 'mg',
  80, 20, 500, 'por_dia', 0.60, 'classe', true, 0, 'parceiro', 'suggested',
  'Antiflatulento: reduz a tensão superficial das bolhas de gás e alivia distensão e desconforto abdominal.',
  'De 20 a 500 mg/dia (Ingold & Akhondi, 2025).', true, now(), now()),

 (uuid_generate_v7(), 'Carvão ativado', 'carvão vegetal ativado', 'mg',
  300, 200, 500, 'por_dia', 0.35, 'classe', true, 1, 'parceiro', 'suggested',
  'Adsorvente usado no alívio de flatulência e distensão por gases.',
  'De 200 a 500 mg/dia (Silberman, Galuska, Taylor, 2023). ATENÇÃO: adsorve fármacos e nutrientes de forma inespecífica. Afastar pelo menos duas horas de qualquer medicamento e das demais fórmulas, ou o carvão leva junto o que deveria ser absorvido.', true, now(), now()),

 (uuid_generate_v7(), 'HMB', 'beta-hidroxi-beta-metilbutirato, HMB CALCIO', 'mg',
  3000, 1000, 3000, 'por_dia', 0.55, 'classe', true, 2, 'parceiro', 'suggested',
  'Metabólito da leucina que reduz o catabolismo muscular e estimula a síntese proteica. Nos análogos de GLP-1 entra pela perda de massa magra que acompanha a perda rápida de peso.',
  'De 1 a 3 g/dia (Kaczka et al., 2019).', true, now(), now()),

 (uuid_generate_v7(), 'PeptiStrong', 'peptídeos de fava, Vicia faba', 'g',
  2.4, 2.4, 20, 'por_dia', 0.50, 'classe', true, 2, 'parceiro', 'suggested',
  'Peptídeos bioativos de fava com efeito sobre preservação de massa magra e síntese muscular.',
  'Dose usual de 2,4 g/dia; o material cita 10 g pela manhã e 10 g à noite em atrofia muscular (Kerr et al., 2023). A distância entre as duas indicações é grande: conferir a fonte antes de usar a dose alta.', true, now(), now()),

 (uuid_generate_v7(), 'Nutricolin', 'silício orgânico, NUTRICOLIN, ácido ortossilícico', 'mg',
  100, 50, 300, 'por_dia', 0.60, 'classe', true, 0, 'parceiro', 'suggested',
  'Silício orgânico biodisponível; estimula síntese de colágeno e queratina, com efeito sobre elasticidade da pele e resistência capilar.',
  'De 50 a 300 mg/dia (Araújo et al., 2016).', true, now(), now()),

 (uuid_generate_v7(), 'Akkermansia muciniphila', 'Bio MAMPs, akkermansia', 'mg',
  75, 50, 100, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Postbiótico com ação sobre a barreira intestinal e a imunomodulação, com dados em disbiose, inflamação e controle de peso.',
  'De 50 a 100 mg/dia (Cani & Knauf, 2021).', true, now(), now()),

 (uuid_generate_v7(), 'Slendesta', 'extrato de batata, inibidor de protease PI2', 'mg',
  100, 50, 300, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'suggested',
  'Extrato de batata que estimula colecistoquinina e reduz o apetite.',
  'Até 300 mg/dia (Peters et al., 2010).', true, now(), now()),

 (uuid_generate_v7(), 'Motility', 'MOTILITY, blend para motilidade', 'mg',
  300, 150, 500, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'pending',
  'Blend do fornecedor para hidratação do bolo fecal e proteção da mucosa intestinal.',
  'De 150 a 500 mg/dia. O próprio material dá como fonte "literatura do fornecedor", sem citação externa — fica como pendente até haver referência independente.', true, now(), now()),

 (uuid_generate_v7(), 'Fibregum B', 'goma acácia, fibra de acácia', 'g',
  2, 0.1, 10, 'por_dia', 0.50, 'classe', true, 0, 'parceiro', 'suggested',
  'Fibra prebiótica de goma acácia; equilibra a microbiota e alivia gases e inchaço.',
  'De 0,1 a 10 g/dia (Singh et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Cistina', 'L-cistina', 'mg',
  150, 100, 200, 'por_dia', 0.65, 'classe', true, 2, 'parceiro', 'suggested',
  'Aminoácido sulfurado, matéria-prima da queratina; usado em queda capilar.',
  'De 100 a 200 mg/dia (Riegel et al., 2020).', true, now(), now()),

 (uuid_generate_v7(), 'Metionina', 'L-metionina', 'mg',
  200, 100, 500, 'por_dia', 0.65, 'classe', true, 2, 'parceiro', 'suggested',
  'Aminoácido essencial sulfurado que participa da produção de queratina.',
  'De 100 a 500 mg/dia (Milani et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Saw palmetto', 'SAW PALMETO, Serenoa repens', 'mg',
  320, 160, 360, 'por_dia', 0.45, 'classe', true, 1, 'parceiro', 'suggested',
  'Inibidor de 5-alfa-redutase de origem vegetal, usado em queda capilar androgenética.',
  'Até 360 mg/dia (Sudeep et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Verisol', 'peptídeos bioativos de colágeno, VERISOL', 'g',
  2.5, 2.5, 2.5, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'suggested',
  'Peptídeos bioativos de colágeno com efeito em firmeza da pele e redução de rugas.',
  'Dose única de 2,5 g/dia nos ensaios (Proksch et al., 2014).', true, now(), now())
ON CONFLICT DO NOTHING;

-- ---------------------------------------------------------------------------------------------
-- 2. Interferência em exame — a ressalva que o material não faz
-- ---------------------------------------------------------------------------------------------

-- A fórmula capilar do material leva biotina 10 mg. A partir de 5 mg/dia a biotina interfere em
-- imunoensaio biotinilado e devolve TSH, troponina, hormônios tireoidianos e hCG falsamente altos
-- ou baixos conforme o formato do ensaio (alerta do FDA de 2017 e 2019; orientação da AACC).
-- Neste sistema isso fecha um ciclo ruim: a receita sai daqui e o exame corrompido volta para cá,
-- alimentando as regras de dose dinâmica.
UPDATE magistral_components
   SET assay_interference = 'acima de 5 mg/dia interfere em imunoensaio biotinilado (TSH, T4 livre, troponina, hCG, hormônios), com resultado falsamente alto ou baixo conforme o ensaio. Suspender 3 dias antes da coleta e avisar o laboratório',
       assay_interference_dose = 5
 WHERE name = 'Biotina';

-- ---------------------------------------------------------------------------------------------
-- 3. Par curado: o carvão ativado leva junto o que deveria ser absorvido
-- ---------------------------------------------------------------------------------------------
WITH carvao AS (SELECT id FROM magistral_components WHERE name = 'Carvão ativado'),
     alvos AS (SELECT id, name FROM magistral_components
                WHERE name IN ('Lipase', 'Protease', 'Alfa-amilase', 'Metilcobalamina', 'Ferro',
                               'Zinco quelato', 'Palmitato de ascorbila'))
INSERT INTO magistral_incompatibilities (id, component_a_id, component_b_id, severity, mechanism, note)
SELECT uuid_generate_v7(), carvao.id, alvos.id, 'warn',
       'o carvão ativado adsorve de forma inespecífica e reduz a absorção do que estiver junto',
       'Não impede a associação, mas na mesma cápsula o carvão tira parte do efeito do outro ativo. Afastar pelo menos duas horas.'
  FROM carvao, alvos
 WHERE NOT EXISTS (SELECT 1 FROM magistral_incompatibilities i
                    WHERE i.component_a_id = carvao.id AND i.component_b_id = alvos.id);

-- ---------------------------------------------------------------------------------------------
-- 4. As que já existiam do formulário ganham a faixa e a referência deste material
--
-- O formulário só dava o nome; este documento dá faixa com citação. Quando as duas fontes falam
-- da mesma substância, vale a que tem referência.
-- ---------------------------------------------------------------------------------------------

UPDATE magistral_components SET usual_dose=5, min_dose=3, max_dose=20, dose_basis='por_dia', default_unit='g',
  dose_reference='Até 20 g/dia (McRorie & McKeown, 2017). Tomar com bastante líquido; sem água suficiente a fibra piora a obstipação.',
  indications='Fibra que absorve água no intestino, aumenta o volume do bolo fecal e estimula o trânsito. Nos análogos de GLP-1 entra pela constipação por retardo do esvaziamento gástrico.',
  evidence_status='suggested' WHERE name='Psyllium';

UPDATE magistral_components SET usual_dose=3000, min_dose=1000, max_dose=3000, dose_basis='por_dia',
  dose_reference='De 1 a 3 g/dia (Kaczka et al., 2019).',
  indications='Metabólito da leucina que reduz o catabolismo muscular e estimula a síntese proteica. Nos análogos de GLP-1 entra pela perda de massa magra que acompanha a perda rápida de peso.',
  evidence_status='suggested' WHERE name='HMB';

UPDATE magistral_components SET usual_dose=100, min_dose=50, max_dose=300, dose_basis='por_dia',
  dose_reference='De 50 a 300 mg/dia (Araújo et al., 2016).',
  indications='Silício orgânico biodisponível; estimula síntese de colágeno e queratina, com efeito sobre elasticidade da pele e resistência capilar.',
  evidence_status='suggested' WHERE name='Nutricolin';

UPDATE magistral_components SET usual_dose=2.5, min_dose=2.5, max_dose=2.5, dose_basis='por_dia', default_unit='g',
  dose_reference='Dose única de 2,5 g/dia nos ensaios (Proksch et al., 2014).',
  indications='Peptídeos bioativos de colágeno com efeito em firmeza da pele e redução de rugas.',
  evidence_status='suggested' WHERE name='Verisol';

UPDATE magistral_components SET usual_dose=600, min_dose=300, max_dose=900, dose_basis='por_dia',
  dose_reference='De 300 a 900 mg/dia (Pirahanchi & Sharma, 2025).',
  indications='Enzima que auxilia na digestão de lipídios. Nos análogos de GLP-1 entra pela alteração da secreção biliar.',
  evidence_status='suggested' WHERE name='Lipase';

UPDATE magistral_components SET usual_dose=200, min_dose=100, max_dose=300, dose_basis='por_dia',
  dose_reference='De 100 a 300 mg/dia (Liu et al., 2024).',
  indications='Enzima que favorece a degradação de proteínas e a absorção.',
  evidence_status='suggested' WHERE name='Protease';

UPDATE magistral_components SET usual_dose=100, min_dose=30, max_dose=300, dose_basis='por_dia',
  dose_reference='De 30 a 300 mg/dia (Ianiro et al., 2016).',
  indications='Enzima que facilita a degradação de carboidratos e reduz fermentação e desconforto gastrointestinal.',
  evidence_status='suggested' WHERE name='Alfa-amilase';

-- O limiar de interferência precisa estar na unidade em que a substância é cadastrada. Biotina
-- está em mcg no catálogo, e o limiar tinha sido gravado em mg: 500 mcg de rotina disparavam
-- alarme (500 > 5) e 10 mg passavam batido no teto. 5 mg = 5.000 mcg.

UPDATE magistral_components SET assay_interference_dose = 5000
 WHERE name = 'Biotina' AND default_unit = 'mcg' AND assay_interference_dose = 5;

-- ═══ magistral-pentravan.sql ═══
-- Pentravan — veículo transdérmico (material da Fagron).
--
-- O que este documento acrescenta ao sistema é de outra natureza: não é dose oral, é VIA. Entram
-- as substâncias que o material formula em Pentravan, cada uma com a categoria de receita que
-- carrega — e é isso que faz "testosterona 50 mg" sair como Controle Especial em vez de receita
-- simples.
--
-- A tabela de permeação do material (ativo, concentração, tecido, percentual em 24 ou 48 h) vai
-- na posologia de cada substância: é o dado que justifica a via.

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, default_category, regulatory_note, source, evidence_status, indications,
   dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Testosterona micronizada', 'testosterona, testosterona Fagron Micro', 'mg',
  50, 0.5, 90, 'por_dia', 0.60, 'classe', 'c5',
  'Esteroide androgênico e anabolizante. A Resolução CFM 2.333/2023 veda a prescrição com finalidade estética, ganho de massa muscular ou desempenho esportivo; a reposição exige deficiência comprovada com nexo causal. Lista C5 da Portaria 344/98: Receituário de Controle Especial em duas vias.',
  'parceiro', 'suggested',
  'Reposição androgênica por via transdérmica ou vulvar. No material, deficiência androgênica feminina usa 0,5 a 5 mg/mL e declínio androgênico masculino 40 a 90 mg/mL.',
  'Em Pentravan a 1% (10 mg/g), permeação de 68,3% em 24 h em pele humana; a 5% (50 mg/g), 68,31% em 24 h e 76,8% em 48 h (Polonini et al., 2014).', true, now(), now()),

 (uuid_generate_v7(), 'Oxandrolona', 'oxandrolona Fagron', 'mg',
  10, 5, 20, 'por_dia', 0.60, 'classe', 'c5',
  'Esteroide anabolizante. A Resolução CFM 2.333/2023 veda a prescrição com finalidade estética, ganho de massa muscular ou desempenho esportivo — e a própria fórmula do material se intitula "sarcopenia e ganho de peso". Sarcopenia com deficiência documentada é outra conversa, e precisa estar escrita no prontuário. Lista C5.',
  'parceiro', 'suggested',
  'Anabolizante usado por via transdérmica em sarcopenia no material do fornecedor.',
  'Em Pentravan a 2% (20 mg/g), permeação de 25,9% em 24 h em pele humana (Polonini et al., 2017).', true, now(), now()),

 (uuid_generate_v7(), '17-beta-estradiol', 'estradiol, E2', 'mg',
  1, 0.25, 2, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estrogênio para sintomas climatéricos por via transdérmica.',
  'Em Pentravan a 0,1% (1 mg/g), permeação de 86,33% em 24 h e 99,9% em 48 h em pele humana (Polonini et al., 2014). No material, 0,25 a 2 mg/mL.', true, now(), now()),

 (uuid_generate_v7(), 'Estriol', 'E3', 'mg',
  4, 2, 8, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estrogênio de ação local, usado em climatério e em prevenção de flacidez cutânea.',
  'Em Pentravan a 0,4% (4 mg/g), permeação de 43,67% em 24 h (Polonini et al., 2014). No material, 2 a 8 mg/mL por via transdérmica e 0,3% no uso facial.', true, now(), now()),

 (uuid_generate_v7(), 'Progesterona micronizada', 'progesterona', 'mg',
  50, 20, 80, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Proteção endometrial por via vaginal.',
  'De 20 a 80 mg por grama de Pentravan, aplicado por via vaginal nos últimos 13 a 15 dias do mês (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Gestrinona', 'gestrinona Fagron', 'mg',
  5, 2.5, 5, 'por_dia', 0.60, 'classe', 'simple',
  'Esteroide com ação androgênica. Uso em endometriose é off-label no Brasil e ganhou escrutínio depois dos implantes hormonais: registrar indicação e consentimento.',
  'parceiro', 'suggested',
  'Endometriose por via vaginal.',
  'Em Pentravan a 0,5% (5 mg/g), permeação de 61,4% em 24 h em mucosa vaginal suína. Estudos clínicos de Maia Jr. et al. (2015 a 2019) com gestrinona vaginal em Pentravan.', true, now(), now()),

 (uuid_generate_v7(), 'Danazol', 'danazol', 'mg',
  50, 50, 100, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Mastalgia cíclica por via transdérmica na mama.',
  '50 mg por grama de Pentravan, uma aplicação ao dia (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Citrato de sildenafila', 'sildenafila', '%',
  0.25, 0.25, 0.25, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estimulante sexual feminino por uso vulvar.',
  '0,25% em Pentravan, 1 mL na região dos lábios vaginais 30 minutos antes da relação (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Tadalafila', 'tadalafila', 'mg',
  5, 5, 5, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Inibidor de fosfodiesterase-5 associado à modulação de testosterona no material.',
  'Em Pentravan a 0,5% (5 mg/g), permeação de 89,07% em 12 h em pele humana (Calixto, 2015).', true, now(), now()),

 (uuid_generate_v7(), 'Alprostadil', 'alprostadil Fagron, PGE1', 'mcg',
  100, 100, 1000, 'por_dose', 0.60, 'classe', 'simple',
  'Prostaglandina E1. Uso tópico peniano e vulvar é off-label; registrar indicação.',
  'parceiro', 'suggested',
  'Disfunção erétil e transtorno da excitação sexual feminina, por uso tópico sob demanda.',
  'Cada pump com 100 mcg; 5 a 10 pumps por aplicação, no mínimo três vezes por semana (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Mesilato de fentolamina', 'fentolamina', 'mg',
  4, 4, 4, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Alfabloqueador associado ao alprostadil na disfunção erétil tópica.',
  '4 mg por pump, associado a 100 mcg de alprostadil (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Baclofeno', 'baclofen', 'mg',
  50, 25, 50, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Relaxante muscular de ação central, usado por via tópica na vulvodínia.',
  '50 mg por mL de Pentravan, associado a PEA, 1 a 2 aplicações ao dia (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Palmitoiletanolamida', 'PEA, PEA BioActive', 'mg',
  10, 10, 600, 'por_dia', 0.55, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Amida lipídica endógena com ação analgésica e anti-inflamatória; na vulvodínia entra por via tópica.',
  '10 mg por mL de Pentravan no material. O teto da IN 28 para uso oral é de 600 mg/dia.', true, now(), now()),

 (uuid_generate_v7(), 'Metformina', 'metformina HCl, cloridrato de metformina', 'mg',
  75, 50, 100, 'por_dia', 0.65, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Modulação de AMPK para longevidade, por via transdérmica no material.',
  'Em Pentravan a 10% (100 mg/g), permeação de 46,7% em 24 h em pele humana (Polonini et al., 2019). No material, 50 a 100 mg/mL duas vezes ao dia.', true, now(), now()),

 (uuid_generate_v7(), 'Miodesin', 'MIODESIN', 'mg',
  170, 170, 170, 'por_dia', 0.50, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Blend anti-inflamatório do fornecedor, usado em miomatose e endometriose por via vaginal.',
  '170 mg por grama de Pentravan, uma aplicação à noite por até dois meses. Estudos de Maia Jr. et al. (2018, 2019).', true, now(), now()),

 (uuid_generate_v7(), 'SiliciuMax', 'silício líquido, SILICIUMAX', '%',
  5, 5, 30, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Silício em forma líquida para uso tópico, em flacidez e envelhecimento cutâneo.',
  'Em Pentravan a 30% (300 mg/g), permeação de 60% em 24 h em pele humana. No material facial, 5%.', true, now(), now())
ON CONFLICT DO NOTHING;

-- Pentravan como veículo. Entra no catálogo para ter densidade e ficar disponível na busca, mas
-- sem faixa de dose: é base, não ativo.
INSERT INTO magistral_components
  (id, name, synonyms, default_unit, bulk_density, density_source, source, evidence_status,
   indications, dose_reference, is_active, created_at, updated_at)
VALUES (uuid_generate_v7(), 'Pentravan', 'PENTRAVAN, veículo transdérmico lipossomal', 'g',
  1.0, 'classe', 'parceiro', 'suggested',
  'Veículo transdérmico em matriz fosfolipídica lamelar com partículas nanossomais, para permeação de ativos em pele íntegra e mucosa.',
  'Sem faixa de dose de propósito: é veículo, entra como qsp. A concentração do ativo é que define a fórmula.', true, now(), now())
ON CONFLICT DO NOTHING;

-- ═══ magistral-peptideos.sql ═══
-- Peptídeos biomiméticos (material de dermatologia).
--
-- Deste documento entram as SUBSTÂNCIAS, não as fórmulas. O PDF é diagramado em duas colunas e a
-- extração mistura componentes de fórmulas vizinhas: montar receita a partir de uma atribuição
-- que eu não consigo garantir seria pior do que não montar. As concentrações abaixo são as que o
-- material usa, e cada uma aparece de forma inequívoca ao lado do próprio nome.
--
-- Todas em uso tópico, unidade em porcentagem — é como se escreve fórmula dermatológica.

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, source, evidence_status, indications, dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Argireline', 'acetil hexapeptídeo-8, ARGIRELINE', '%', 5, 3, 8, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo inibidor de neurotransmissor: reduz a contração muscular mimética e suaviza linhas de expressão.',
  'De 3 a 8% no material, em sérum de uso facial. Também usado como suporte pós-toxina botulínica.', true, now(), now()),

 (uuid_generate_v7(), 'Syn-Ake', 'dipeptídeo diaminobutiroil benzilamida, SYN-AKE', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo miorrelaxante tópico, análogo sintético de peptídeo do veneno de Tropidolaemus wagleri.',
  '2% no material, associado a Argireline em sérum facial.', true, now(), now()),

 (uuid_generate_v7(), 'Munapsys', 'acetil octapeptídeo-3, MUNAPSYS', '%', 3, 1, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com ação sobre a contração muscular mimética, associado aos demais em rugas dinâmicas.',
  'De 1 a 3% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Matrixyl Synthe-6', 'palmitoil tripeptídeo-38, MATRIXYL SYNTHE 6', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo sinalizador: estimula fibroblasto e a síntese de matriz extracelular.',
  '2% no material, em fórmula de firmeza para pele madura.', true, now(), now()),

 (uuid_generate_v7(), 'Idealift', 'peptídeo tensor, IDEALIFT', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com efeito tensor e de firmeza cutânea.',
  '2% no material.', true, now(), now()),

 (uuid_generate_v7(), 'GHK-Cu', 'peptídeo de cobre, tripeptídeo-1 de cobre, copper peptide', '%', 1, 0.5, 2, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo de cobre com ação em reparo tecidual, remodelação de matriz e cicatrização.',
  '1% no material. É a molécula com literatura própria mais antiga deste grupo.', true, now(), now()),

 (uuid_generate_v7(), 'Haloxyl', 'HALOXYL, quimiotripsina peptídica para olheira', '%', 3, 2, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Blend peptídico para olheiras: atua sobre a hemossiderina depositada na região periorbital.',
  '3% no material, em fórmula periorbital.', true, now(), now()),

 (uuid_generate_v7(), 'TGP-2', 'TGP-2 peptídeo, Nano TGP-2', '%', 2, 1, 2, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo para região periorbital, associado ao Haloxyl no material.',
  'De 1 a 2% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Nopigmerin', 'NOPIGMERIN', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com ação sobre a pigmentação cutânea.',
  '3% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Procapil', 'PROCAPIL, biotinil tripeptídeo-1', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Complexo com biotinil tripeptídeo para couro cabeludo, voltado à queda capilar.',
  '3% no material, em loção capilar.', true, now(), now()),

 (uuid_generate_v7(), 'Prohairin', 'PROHAIRIN, peptídeo capilar', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo para estímulo do crescimento capilar.',
  '3% no material, associado ao Procapil.', true, now(), now()),

 (uuid_generate_v7(), 'EGF Nanofactor', 'fator de crescimento epidérmico, EGF', '%', 0.5, 0.5, 0.5, 'por_dose', 1.0, 'classe',
  'parceiro', 'pending',
  'Fator de crescimento epidérmico em veiculação nanossomal, usado em regeneração e em protocolo capilar.',
  '0,5% no material. Fator de crescimento tópico é assunto com discussão regulatória e de segurança própria: fica como pendente até conferência.', true, now(), now()),

 (uuid_generate_v7(), 'FGF Nanofactor', 'fator de crescimento de fibroblasto, FGF', '%', 0.5, 0.5, 0.5, 'por_dose', 1.0, 'classe',
  'parceiro', 'pending',
  'Fator de crescimento de fibroblasto em veiculação nanossomal.',
  '0,5% no material. Mesma ressalva do EGF.', true, now(), now()),

 (uuid_generate_v7(), 'Nano IDP-2', 'NANO IDP-2', '%', 1, 1, 1, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo em veiculação nanossomal para reparo e rejuvenescimento.',
  '1% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Vc-IP', 'tetraisopalmitato de ascorbila, VC-IP', '%', 3, 2, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Derivado lipossolúvel de vitamina C para clareamento e antioxidação cutânea.',
  '3% no material, em sérum facial. É a forma lipossolúvel da vitamina C, mais estável que o ácido ascórbico livre.', true, now(), now())
ON CONFLICT DO NOTHING;

-- ═══ magistral-arquitetura-hormonal.sql ═══
-- Suplementação na arquitetura hormonal (material da Arboretum): eixos androgênico, DHEA e HPA.
--
-- Três fórmulas, extraídas com fatia por coluna — o PDF é diagramado em duas colunas e a leitura
-- ingênua misturava componentes de fórmulas vizinhas.
--
-- Uma ressalva de literatura entra junto: o material apresenta Tribulus, Testofen e Eurycoma como
-- otimizadores de testosterona livre. Para o Tribulus a evidência não sustenta: em revisão
-- sistemática de 2025, oito de dez ensaios não mostraram mudança no perfil androgênico, e os dois
-- que mostraram tiveram magnitude pequena (60 a 70 ng/dL) em homens com hipogonadismo. Isso fica
-- escrito na substância, não escondido.

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, source, evidence_status, indications, dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Testofen', 'TESTOFEN, extrato de feno-grego padronizado', 'mg',
  300, 50, 600, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Extrato padronizado de feno-grego usado para suporte androgênico e libido.',
  '50 mg na fórmula do material. Os ensaios do próprio ingrediente usam 300 a 600 mg/dia — a dose da fórmula fica bem abaixo disso.', true, now(), now()),

 (uuid_generate_v7(), 'Eurycoma longifolia', 'long jack, tongkat ali, EURYCOMA', 'mg',
  200, 100, 400, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Fitoterápico usado em libido, estresse e suporte androgênico.',
  '200 mg no material. Ensaios costumam usar 200 a 400 mg/dia de extrato padronizado.', true, now(), now()),

 (uuid_generate_v7(), 'Turkesterone', 'turkesterona, Ajuga turkestanica', 'mg',
  300, 250, 500, 'por_dia', 0.45, 'classe', 'parceiro', 'pending',
  'Ecdisteroide vegetal apresentado como anabolizante não hormonal.',
  '300 mg no material. Fica como pendente de propósito: os dados humanos são escassos e a promessa anabólica repete a linguagem que a Resolução CFM 2.333/2023 restringe para hormônio — em fitoterápico não é vedação, mas a expectativa criada no paciente é a mesma.', true, now(), now()),

 (uuid_generate_v7(), 'Robuvit', 'ROBUVIT, extrato de carvalho francês, Quercus robur', 'mg',
  200, 100, 300, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Extrato de carvalho francês (roburinas) com dados em fadiga e recuperação.',
  '120 mg no material; os ensaios do ingrediente usam 200 a 300 mg/dia.', true, now(), now()),

 (uuid_generate_v7(), 'UbiQsome', 'UBIQSOME, coenzima Q10 fitossoma', 'mg',
  100, 50, 200, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Coenzima Q10 em veiculação fitossomal, com biodisponibilidade maior que a ubiquinona pura.',
  '100 mg no material. Como é CoQ10 veiculada, o teto de 200 mg/dia do Anexo IV da IN 28 para coenzima Q10 se aplica.', true, now(), now()),

 (uuid_generate_v7(), 'Panax ginseng', 'ginseng coreano, ginseng vermelho', 'mg',
  200, 100, 400, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Adaptógeno com dados em fadiga, cognição e resposta adrenal.',
  '150 mg no material. Ensaios usam 200 a 400 mg/dia de extrato padronizado em ginsenosídeos.', true, now(), now())
ON CONFLICT DO NOTHING;

-- A ressalva sobre o Tribulus vai na substância que já existe no catálogo.
UPDATE magistral_components
   SET dose_reference = coalesce(dose_reference || ' ', '') ||
       'Revisão sistemática de 2025: oito de dez ensaios sem mudança no perfil androgênico; os dois positivos tiveram magnitude pequena (60 a 70 ng/dL) em homens com hipogonadismo. Metanálise de 2023 encontrou aumento não significativo de testosterona e LH. Prescrever com essa expectativa, e não com a de reposição.',
       evidence_status = 'suggested'
 WHERE name = 'Tribulus terrestris' AND coalesce(dose_reference,'') NOT LIKE '%oito de dez%';

-- ═══ magistral-avulsos-capturados.sql ═══
-- Substâncias e fórmulas que existiam só no banco de desenvolvimento.
--
-- POR QUE ESTE ARQUIVO EXISTE: carregar todos os seeds num banco vazio devolveu 281 substâncias e
-- 122 fórmulas, contra 290 e 132 no dev. A diferença veio de comandos avulsos rodados direto no
-- psql durante o trabalho, que nunca foram capturados em arquivo — deploy de produção nasceria
-- diferente do que foi conferido aqui.
--
-- O teste que achou isso vale mais que o conserto: carga limpa num banco vazio, e comparação de
-- contagem com o dev. Idempotente.

-- ---------------------------------------------------------------------------------------------
-- Substâncias
-- ---------------------------------------------------------------------------------------------

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Rhodiola rosea', 'rodiola, raiz-do-ártico', NULL, NULL, 'mg', 150.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Adaptógeno usado em fadiga associada a estresse e em suporte de foco.', NULL, 'fadiga por estresse
foco e desempenho mental', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Rhodiola rosea');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Piperina', 'extrato de pimenta preta, BioPerine', NULL, NULL, 'mg', 10.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Usada para aumentar a biodisponibilidade da curcumina e de outros ativos.', NULL, 'potencializa absorção da curcumina', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Piperina');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Policosanol', '', NULL, NULL, 'mg', 30.0, NULL, NULL, 'por_dia', 0.45, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Álcool graxo de cana usado em fórmulas de perfil lipídico; evidência heterogênea, com resultados positivos concentrados em estudos cubanos.', NULL, 'coadjuvante lipídico', NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Policosanol');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'L-tirosina', 'tirosina, L TIROSINA', NULL, NULL, 'mg', 500.0, NULL, NULL, 'por_dia', 0.6, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Aminoácido precursor de dopamina, noradrenalina e hormônio tireoidiano.', NULL, 'precursor de catecolaminas
cofator da síntese tireoidiana', NULL, NULL, NULL, NULL, 'Tirosina', 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'L-tirosina');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'D-ribose', 'ribose, D RIBOSE', NULL, NULL, 'g', 5.0, NULL, NULL, 'por_dia', 0.55, 'classe', false, false, false, false, false, NULL, NULL, NULL, 'parceiro', 'suggested', 'Açúcar de cinco carbonos usado em suporte energético e recuperação.', NULL, 'fadiga e recuperação
suporte de ATP', NULL, NULL, NULL, NULL, 'D-ribose', 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'D-ribose');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Beta-hidroxibutirato', 'BHB, cetona exógena, sais de BHB', NULL, NULL, 'g', 10.0, 3.0, 12.0, 'por_dia', 0.8, 'classe', false, true, false, false, false, 3, true, NULL, 'pesquisa', 'suggested', 'Corpo cetônico usado como fonte energética alternativa em protocolos cetogênicos e de desempenho cognitivo. Vem como sal de cálcio, magnésio, sódio ou potássio, e a carga desses minerais conta na fórmula.', 'De 3 a 12 g/dia do sal, geralmente fracionados. Higroscópico e salgado: sachê é a forma prática. A carga de sódio e cálcio do sal precisa entrar na conta do dia.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Beta-hidroxibutirato');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Gynostemma pentaphyllum', 'jiaogulan, ginostema, ginseng do sul', NULL, NULL, 'mg', 300.0, 200.0, 450.0, 'por_dia', 0.5, 'classe', false, false, false, false, false, 3, true, NULL, 'pesquisa', 'suggested', 'Adaptógeno com ação sobre a AMPK, usado em resistência insulínica e esteatose. Amargo característico.', 'De 200 a 450 mg/dia do extrato padronizado em gipenosídeos. Os ensaios em esteatose e resistência insulínica usam essa faixa por 8 a 12 semanas.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Gynostemma pentaphyllum');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Goma cássia', 'cassia gum, goma de cássia', NULL, NULL, 'mg', NULL, NULL, NULL, 'por_dia', 0.7, 'classe', false, true, false, false, false, 0, true, NULL, 'pesquisa', 'suggested', 'Fibra solúvel usada como espessante e agente de corpo em sachês. Não tem dose terapêutica própria: a quantidade vem da textura pretendida.', 'Sem faixa cadastrada de propósito: é excipiente de textura, não ativo com posologia estabelecida.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Goma cássia');

INSERT INTO magistral_components (id, name, synonyms, cas, dcb_code, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density, density_source, eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok, notes, source, evidence_status, indications, dose_reference, indication_bullets, dose_bullets, elemental_percent, correction_note, brand, in28_nutrient, in28_factor, default_category, regulatory_note, assay_interference, assay_interference_dose, is_active, created_at, updated_at)
SELECT uuid_generate_v7(), 'Ácido hidroxicítrico', 'HCA, Citrimax, Garcinia cambogia, GARCINEA CAMBOJA, GARCINIA CAMBOJA', NULL, NULL, 'mg', 1500.0, 1500.0, 3000.0, 'por_dia', 0.6, 'classe', false, false, false, false, false, 2, true, NULL, 'pesquisa', 'suggested', 'Extrato de Garcinia cambogia padronizado em ácido hidroxicítrico, usado para saciedade e controle de peso. As metanálises mostram efeito pequeno e inconsistente sobre o peso.', 'De 500 a 1.000 mg três vezes ao dia, antes das refeições, padronizado a 50 a 60% de HCA. O efeito sobre o peso nas metanálises é pequeno e de qualidade baixa.', NULL, NULL, NULL, NULL, NULL, NULL, 1.0, 'simple', NULL, NULL, NULL, true, now(), now()
 WHERE NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Ácido hidroxicítrico');


-- ---------------------------------------------------------------------------------------------
-- Fórmulas-base
-- ---------------------------------------------------------------------------------------------

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Fórmula sublingual do sono', 'Indução do sono por via mais fisiológica, antes de recorrer à melatonina. Atua como precursor de serotonina e melatonina com efeito gabaérgico.', 'Insônia inicial
Precursores de melatonina
Efeito gabaérgico noturno', 'sublingual', 'internal', '', '', 60.0, 'cápsulas', 'Uso sublingual antes de dormir', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Fórmula sublingual do sono' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg', 'simple', '', false),
    (1, 'L-teanina', 20.0::numeric, 'mg', 'simple', '', false),
    (2, 'Piridoxal-5-fosfato (P5P)', 10.0::numeric, 'mg', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Precursores de melatonina sublingual', 'Estímulo à via triptofano, 5-HTP, serotonina e melatonina. Cautela em usuários de ISRS.', 'Insônia
Suporte à via serotoninérgica', 'sublingual', 'internal', '', '', 60.0, 'cápsulas', 'À noite, por via sublingual', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Precursores de melatonina sublingual' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg', 'simple', 'sublingual à noite', false),
    (1, 'Piridoxal-5-fosfato (P5P)', 10.0::numeric, 'mg', 'simple', 'sublingual', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê matinal mitocondrial', 'Suporte à função mitocondrial e ao metabolismo energético.', 'Otimização da função mitocondrial
Baixa energia em condições crônicas
Pacientes acima de 50 anos', 'sachê', 'internal', '', '', 60.0, 'cápsulas', '1 sachê pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê matinal mitocondrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'L-carnitina', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'D-ribose', 5.0::numeric, 'g', 'simple', 'cautela em diabéticos', false),
    (2, 'Magnésio glicina', 150.0::numeric, 'mg', 'simple', '', true)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Fórmula pré-refeição para controle glicêmico', 'Suporte ao controle glicêmico e à resistência insulínica em contexto de emagrecimento.', 'Resistência insulínica
Apoio ao emagrecimento
Uso pré-refeição', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', 'Antes das refeições', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Fórmula pré-refeição para controle glicêmico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Citrimax (ácido hidroxicítrico - HCA)', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'Gimnema silvestre', 300.0::numeric, 'mg', 'simple', 'faixa de 200 a 300 mg', false),
    (2, 'Ginostema pentaphyllum (Gynostemma pentaphyllum)', 200.0::numeric, 'mg', 'simple', '150 a 200 mg se padronizado a 80% de gipenosídeos; cerca de 500 mg sem padronização, para 100 a 180 mg de gipenosídeos', false),
    (3, 'Cromo (GTF ou picolinato)', 300.0::numeric, 'mcg', 'simple', 'faixa de 200 a 300 mcg', true)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Combinação colinérgica fosfatidilcolina + alfa-GPC', 'Suporte à síntese de acetilcolina por múltiplas vias de absorção, voltado a cognição e foco.', 'Precursores de acetilcolina
Suporte cognitivo
Diversificação de vias de absorção', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', 'Dividir em até 4 doses conforme tolerabilidade', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Combinação colinérgica fosfatidilcolina + alfa-GPC' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Fosfatidilcolina', 250.0::numeric, 'mg', 'simple', '', false),
    (1, 'Alfa-GPC (L-alfa-glicerofosfocolina)', 500.0::numeric, 'mg', 'simple', 'faixa de 250 a 500 mg', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Curcumina com piperina', 'Suporte anti-inflamatório e antioxidante com melhora da biodisponibilidade da curcumina.', 'Modulação da inflamação crônica
Aumento da absorção da curcumina
Cautela em uso de anticoagulantes', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', '500 mg a 2 g/dia de curcuminoides, respeitando 5 mg de piperina a cada 500 mg', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Curcumina com piperina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Curcumina padronizada (95% curcuminoides)', 500.0::numeric, 'mg', 'simple', 'faixa de 500 mg a 2 g/dia conforme tolerância; anticoagulados limitar a ≤500 mg/dia ou 250 mg lipossomada', false),
    (1, 'Piperina', 5.0::numeric, 'mg', 'simple', '5 mg para cada 500 mg de curcumina; avaliar alergia', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê matinal de energia', 'Exemplo prático de prescrição em sachê para uso matinal.', 'Uso matinal
Suporte energético', 'sachê', 'internal', '', '', 60.0, 'cápsulas', 'Diluir e tomar pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê matinal de energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Glutamina', 4.0::numeric, 'g', 'simple', '', false),
    (1, 'Magnésio glicina', 500.0::numeric, 'mg', 'simple', '', false),
    (2, 'L-tirosina', 500.0::numeric, 'mg', 'simple', '', false),
    (3, 'Goma cássia', 2500.0::numeric, 'mg', 'simple', '', false),
    (4, 'BHB', 3.0::numeric, 'g', 'simple', '', false),
    (5, 'D-ribose', 5.0::numeric, 'g', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Sachê noturno de relaxamento', 'Combinação em pó para promover relaxamento e foco no período noturno. Formato em sachê evita ingestão de muitas cápsulas.', 'Relaxamento noturno
Suporte ao foco', 'sachê', 'internal', '', '', 60.0, 'cápsulas', '1 sachê ao deitar', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Sachê noturno de relaxamento' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Magnésio glicina', 500.0::numeric, 'mg', 'simple', '', false),
    (1, 'Magnésio treonato', 500.0::numeric, 'mg', 'simple', '', false),
    (2, 'Inositol', 1.0::numeric, 'g', 'simple', '', false),
    (3, 'Taurina', 1.0::numeric, 'g', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Vitamina D conforme exame', 'Reposição de vitamina D com a dose ajustada pela 25-hidroxivitamina D mais recente do paciente.', 'Reposição de vitamina D guiada pelo exame
Faixa-alvo de 40 a 60 ng/mL
Reavaliar 25-OH-D em 90 dias', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60.0, 'cápsulas', '1 cápsula pela manhã', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Vitamina D conforme exame' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Vitamina D3', 2000.0::numeric, 'UI', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

WITH nova AS (
  INSERT INTO magistral_formula_templates (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle, quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Cápsulas de suporte mitocondrial', 'Suplementação oral para biogênese e desempenho mitocondrial em condições crônicas degenerativas, neurológicas ou metabólicas.', 'Disfunção mitocondrial
Condições degenerativas crônicas
Fadiga e baixa energia
Pacientes acima de 50 anos', 'cápsula', 'internal', '', '', 60.0, 'cápsulas', '1 dose ao dia', 0, NULL, NULL
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t
                      WHERE t.name = 'Cápsulas de suporte mitocondrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, c.e FROM nova, (VALUES
    (0, 'Acetil-L-carnitina', 500.0::numeric, 'mg', 'simple', 'em jejum, manhã ou tarde', false),
    (1, 'Coenzima Q10 (ubiquinona) ou ubiquinol', 100.0::numeric, 'mg', 'simple', 'preferir com refeição gordurosa', false),
    (2, 'Vitamina B2 (riboflavina)', 25.0::numeric, 'mg', 'simple', '', false),
    (3, 'Vitamina B3 (nicotinamida)', 100.0::numeric, 'mg', 'simple', '', false),
    (4, 'Vitamina B6 (piridoxal-5-fosfato)', 10.0::numeric, 'mg', 'simple', '', false),
    (5, 'Magnésio dimalato', 500.0::numeric, 'mg', 'simple', '', false),
    (6, 'Ácido alfa-lipoico', 600.0::numeric, 'mg', 'simple', '300 a 600 mg, final da tarde em jejum, pode exigir cápsula gastrorresistente', false),
    (7, 'PQQ (pirroloquinolina quinona)', 20.0::numeric, 'mg', 'simple', '', false)
) AS c(ord, s, q, u, cat, n, e);

-- ═══ magistral-faixas-corrigidas.sql ═══
-- Correção das faixas de dose que contradiziam o próprio texto de posologia da linha.
--
-- Como apareceram: a faixa numérica foi semeada a partir das fórmulas das parceiras (dose de UMA
-- cápsula) e o texto veio da literatura (dose do DIA). O painel comparava uma contra a outra,
-- então fórmula com dose baixa passava sem alerta e fórmula correta era acusada.
--
-- Aqui a faixa numérica passa a ser DIÁRIA em todas, com o valor da literatura. Idempotente.

-- Gimnema: extrato padronizado, 200 a 400 mg/dia antes das refeições. Estava 12,5 a 25 mg, que é
-- a dose por cápsula da fórmula "Resistência insulínica" — o catálogo conferia a fórmula contra
-- ela mesma. Ganha também a grafia com i, que é como as fórmulas escrevem.
UPDATE magistral_components
   SET min_dose = 200, max_dose = 400, usual_dose = 300, dose_basis = 'por_dia',
       synonyms = 'Gymnema sylvestre, Gimnema silvestre, gurmar'
 WHERE name = 'Gymnema silvestre';

-- N-acetilcisteína: 600 a 1.800 mg/dia.
UPDATE magistral_components
   SET min_dose = 600, max_dose = 1800, usual_dose = 600, dose_basis = 'por_dia'
 WHERE name = 'N-acetilcisteína';

-- Picolinato de cromo: 200 a 1.000 mcg/dia de cromo elementar.
UPDATE magistral_components
   SET min_dose = 200, max_dose = 1000, usual_dose = 400, dose_basis = 'por_dia'
 WHERE name = 'Picolinato de cromo';

-- K2 MK-7: 90 a 360 mcg/dia; 200 mcg é a dose comum.
UPDATE magistral_components
   SET min_dose = 90, max_dose = 360, usual_dose = 200, dose_basis = 'por_dia'
 WHERE name = 'Vitamina K2 MK-7';

-- Acetil-L-carnitina: 500 mg a 2 g/dia.
UPDATE magistral_components
   SET min_dose = 500, max_dose = 2000, usual_dose = 1000, dose_basis = 'por_dia'
 WHERE name = 'Acetil-L-carnitina';

-- PQQ: 10 a 20 mg/dia.
UPDATE magistral_components
   SET min_dose = 10, max_dose = 20, usual_dose = 20, dose_basis = 'por_dia'
 WHERE name = 'PQQ';

-- Ácido alfa-lipoico: 300 a 600 mg/dia, chegando a 1,3 g. A faixa de 25 a 200 mg deixava passar
-- sem alerta qualquer fórmula com dose sub-terapêutica.
UPDATE magistral_components
   SET min_dose = 300, max_dose = 1300, usual_dose = 600, dose_basis = 'por_dia'
 WHERE name = 'Ácido alfa-lipoico';

-- Iodo: a faixa começava em 200 mcg, acima da própria RDA (150 mcg). Em tireoidite autoimune o
-- excesso de iodo é gatilho conhecido, então piso alto aqui empurra para o lado errado.
UPDATE magistral_components
   SET min_dose = 75, max_dose = 600, usual_dose = 150, dose_basis = 'por_dia'
 WHERE name = 'Iodo';

-- ---------------------------------------------------------------------------------------------
-- Segunda leva: faixas cujo extremo é exatamente a dose de uma cápsula das fórmulas parceiras
-- (impressão digital de faixa semeada da própria fórmula, não da literatura).
-- ---------------------------------------------------------------------------------------------

-- 5-HTP: 50 a 300 mg/dia por via oral. O teto de 100 mg era a dose da cápsula da "Ansiedade
-- diurna" e transformava 2 cápsulas ao dia em alerta de dose alta.
UPDATE magistral_components
   SET min_dose = 50, max_dose = 300, usual_dose = 100, dose_basis = 'por_dia'
 WHERE name = '5-HTP';

-- Valeriana: extrato de 300 a 600 mg antes de dormir, chegando a 900 mg nos estudos de insônia.
UPDATE magistral_components
   SET min_dose = 200, max_dose = 900, usual_dose = 400, dose_basis = 'por_dia'
 WHERE name = 'Valeriana';

-- Ginseng brasileiro (fáfia): o próprio registro diz que os trechos não informam posologia. A
-- faixa de 30 a 60 mg saiu da fórmula, não da literatura — faixa inventada é pior que faixa
-- ausente, porque vira alerta com cara de fundamento. Volta a ser nula.
UPDATE magistral_components
   SET min_dose = NULL, max_dose = NULL, usual_dose = NULL,
       dose_reference = coalesce(dose_reference, '') ||
         ' Sem faixa cadastrada de propósito: não há posologia estabelecida na literatura consultada.'
 WHERE name = 'Ginseng brasileiro';

-- Piridoxal-5-fosfato: teto de 50 mg/dia. Fica abaixo do limite da IN 28 (98,6 mg) de propósito,
-- por causa da neuropatia sensitiva descrita em uso prolongado de B6 em dose alta.
UPDATE magistral_components
   SET min_dose = 5, max_dose = 50, usual_dose = 10, dose_basis = 'por_dia'
 WHERE name = 'Piridoxal-5-fosfato';

-- Isoflavona: os ensaios de climatério usam de 40 a 160 mg/dia.
UPDATE magistral_components
   SET min_dose = 40, max_dose = 160, usual_dose = 80, dose_basis = 'por_dia'
 WHERE name = 'Isoflavona';

-- ---------------------------------------------------------------------------------------------
-- Terceira leva: o catálogo discordando das regras de dose dinâmica
-- ---------------------------------------------------------------------------------------------

-- Metilcobalamina: a regra por faixa de B12 prescreve 100 mcg de manutenção para quem está dentro
-- do alvo, e o catálogo dizia que 100 mcg é dose baixa. Dois subsistemas discordando da mesma
-- substância. O piso desce para 100 mcg, que é a manutenção, e o texto guarda a faixa de reposição.
UPDATE magistral_components
   SET min_dose = 100, max_dose = 2000, usual_dose = 500, dose_basis = 'por_dia'
 WHERE name = 'Metilcobalamina';

-- ═══ in28-anexo-iv.sql ═══
-- Anexo IV da IN 28/2018 (texto consolidado). Gerado do texto da norma, não digitado à mão.
-- Fonte: LegisWeb, IN 28 de 26/07/2018 consolidada até a IN 373/2025.

INSERT INTO in28_limits (nutrient, unit, max_adult, kind) VALUES
  ('2''-fucosil-lactose', 'g', 3.0, 'valor'),
  ('Proteínas', 'g', NULL, 'NE'),
  ('Carboidratos', 'g', NULL, 'NE'),
  ('D-ribose', 'g', 2.5, 'valor'),
  ('Fibras alimentares', 'g', NULL, 'NE'),
  ('Lipídeos totais', 'g', NULL, 'NE'),
  ('EPA e DHA', 'mg', 2000.0, 'valor'),
  ('Teanina', 'mg', 250.0, 'valor'),
  ('DHA e EPA obtido de lisinato de ômega 3', 'mg', 2000.0, 'valor'),
  ('Ácido linoleico n-6', 'G', 25.5, 'valor'),
  ('Ácido alfa-linolênico n-3', 'G', 2.4, 'valor'),
  ('Colina', 'mg', 3235.15, 'valor'),
  ('Vitamina Ai', 'µg', 2623.61, 'valor'),
  ('Vitamina B6', 'mg', 98.6, 'valor'),
  ('Vitamina C', 'mg', 1916.02, 'valor'),
  ('Vitamina Dii', 'µg', 50.0, 'valor'),
  ('Niacina', 'mg', 35.0, 'valor'),
  ('Vitamina Eiii', 'mg', 1000.0, 'valor'),
  ('Ácido fólico iv', 'mcg', 1281.5, 'valor'),
  ('Ácido pantotênico', 'mg', 5.64, 'valor'),
  ('Biotina', 'µg', 45.0, 'valor'),
  ('Riboflavina', 'mg', 2.74, 'valor'),
  ('Tiamina', 'mg', 2.02, 'valor'),
  ('Vitamina B12', 'µg', 9.94, 'valor'),
  ('Vitamina K', 'µg', 149.06, 'valor'),
  ('Cálciov', 'mg', 1534.67, 'valor'),
  ('Cobre', 'µg', 8975.52, 'valor'),
  ('Manganês', 'mg', 1.66, 'valor'),
  ('Molibdênio', 'µg', 1955.0, 'valor'),
  ('Fósforov', 'mg', 2083.89, 'valor'),
  ('Selênio', 'µg', 319.75, 'valor'),
  ('Zinco', 'mg', 29.59, 'valor'),
  ('Iodo', 'µg', 919.02, 'valor'),
  ('Ferro', 'mg', 34.31, 'valor'),
  ('Magnésio', 'mg', 350.0, 'valor'),
  ('Cromo', 'µg', 250.0, 'valor'),
  ('Leucina', 'mg', 5660.0, 'valor'),
  ('Lisina', 'mg', 4940.0, 'valor'),
  ('Valina', 'mg', 3600.0, 'valor'),
  ('Isoleucina', 'mg', 3240.0, 'valor'),
  ('Treonina', 'mg', 2720.0, 'valor'),
  ('Fenilalanina', 'mg', 2820.0, 'valor'),
  ('Tirosina', 'mg', 2750.0, 'valor'),
  ('Metionina', 'mg', 1530.0, 'valor'),
  ('Cisteína', 'mg', 830.0, 'valor'),
  ('Histidina', 'mg', 2120.0, 'valor'),
  ('Triptofano', 'mg', 860.0, 'valor'),
  ('Arginina', 'mg', 3810.0, 'valor'),
  ('Aspartato', 'mg', 5320.0, 'valor'),
  ('Glicina', 'mg', 2980.0, 'valor'),
  ('Serina', 'mg', 3151.0, 'valor'),
  ('Ácido glutâmico', 'mg', 15880.0, 'valor'),
  ('Prolina', 'mg', 5360.0, 'valor'),
  ('Alanina', 'mg', 3320.0, 'valor'),
  ('Beta-alanina', 'g', 2.0, 'valor'),
  ('Glutamina', 'mg', 5000.0, 'valor'),
  ('Taurina', 'mg', 2000.0, 'valor'),
  ('L-Carnitina', 'mg', 2000.0, 'valor'),
  ('Creatina', 'mg', 3000.0, 'valor'),
  ('Adenosina', 'mg', 1.2, 'valor'),
  ('Inositol', 'g', 2.0, 'valor'),
  ('Ácido estearidônico (SDA) do óleo da semente de Buglossoides arvensis (L.)', 'mg', 420.0, 'valor'),
  ('Ácido gama aminobutírico (GABA)', 'mg', 300.0, 'valor'),
  ('Ácido hialurônico', 'mg', 157.7, 'valor'),
  ('Boro', 'mg', 8.866, 'valor'),
  ('Cafeína', 'mg', 200.0, 'valor'),
  ('Coenzima Q10', 'mg', 200.0, 'valor'),
  ('Colágeno', 'mg', NULL, 'NE'),
  ('Colágeno tipo II não desnaturado', 'mg', NULL, 'NE'),
  ('Compostos fenólicos deOpuntia ficus-indica', 'mg', 54.6, 'valor'),
  ('Fitoesterois e fitoestanois', 'g', 3.0, 'valor'),
  ('Licopeno', 'mg', 8.0, 'valor'),
  ('Luteína', 'mg', 20.0, 'valor'),
  ('Metilsulfonilmetano', 'mg', 900.0, 'valor'),
  ('Palmitoiletanolamida', 'mg', 600.0, 'valor'),
  ('Zeaxantina', 'mg', 3.0, 'valor'),
  ('Astaxantina', 'mg', 6.0, 'valor'),
  ('Alicina', 'mg', 3.0, 'valor'),
  ('Compostos fenólicos totais', 'mg', NULL, 'NE'),
  ('Compostos fenólicos de Extrato de polpa de oliva (Olea europaea L.)', 'mg', 23.2, 'valor'),
  ('Condroitina', 'mg', 600.0, 'valor'),
  ('Curcumina', 'mg', NULL, 'NA'),
  ('Dihidrocapsiato (DHC)', 'mg', 9.0, 'valor'),
  ('D-Limoneno', 'mg', 150.0, 'valor'),
  ('Extrato de cacauviii', 'mg', 300.0, 'valor'),
  ('10-HDA (ácido hidroxidecenóico)', 'mg', 25.0, 'valor'),
  ('Rutina', 'mg', 0.6, 'valor'),
  ('Saponinas de glicosídeos de furostanol', 'mg', 450.0, 'valor'),
  ('Silício', 'mg', 15.67, 'valor'),
  ('Ácido clorogênico', 'mg', 400.0, 'valor'),
  ('Ácido rosmarínico', 'mg', 160.0, 'valor'),
  ('Alfa-casozepina', 'mg', 21.0, 'valor'),
  ('Antocianinas da laranja moro', 'mg', 5.0, 'valor'),
  ('Antocianinas da casca de jabuticaba', 'mg', 28.0, 'valor'),
  ('Arabinogalactana', 'g', 20.0, 'valor'),
  ('Beta-glucana (Euglena gracilis)', 'mg', 187.5, 'valor'),
  ('Beta-glucana de levedura (Saccharomyces cerevisiae)', 'mg', 1275.0, 'valor'),
  ('Colostro bovino desnatado viii', 'g', 4.0, 'valor'),
  ('Colostro bovino integral viii', 'g', 4.0, 'valor'),
  ('Proantociadininas Proantocianidinas de cranberry', 'mg', NULL, 'NA'),
  ('Fosfatidilserina', 'mg', 400.0, 'valor'),
  ('Glucosamina', 'mg', 750.0, 'valor'),
  ('Glicosaminoglicanos', 'mg', 19.2, 'valor'),
  ('Glutationa', 'mg', 500.0, 'valor'),
  ('Hidroximetilbutirato', 'g', 3.0, 'valor'),
  ('Hidroxitirosol e derivados', 'mg', NULL, 'NA'),
  ('Lactoferrina', 'mg', 1200.0, 'valor'),
  ('L-citrulinilarginina', 'mg', 52.8, 'valor'),
  ('Melatonina', 'mg', 0.21, 'valor'),
  ('Peptídeos bioativos de colágeno hidrolisado com peso molecular médio de 2kDa', 'g', NULL, 'NE'),
  ('Peptídeos bioativos de soro de leite', 'mg', NULL, 'NE'),
  ('Proantocianidinas de semente de uva', 'mg', 116.0, 'valor'),
  ('Proantocianidinas do tipo A', 'mg', 25.0, 'valor'),
  ('Procianidinas', 'mg', 262.5, 'valor'),
  ('Tetraidrocurcuminoides', 'mg', 120.0, 'valor'),
  ('Trans-resveratrol', 'mg', 165.0, 'valor'),
  ('Verbascosídeo', 'mg', 7.68, 'valor'),
  ('Xilo-oligossacarídeos', 'mg', 2.0, 'valor'),
  ('Alfa-galactosidase de Aspergillus niger', 'GaIU', 4800.0, 'valor'),
  ('Protease de Aspergillus niger expressa em Aspergillus niger', 'PPI', 2000000.0, 'valor'),
  ('Fitase', 'FTU', NULL, 'NE'),
  ('Lactase', 'U.FCC', NULL, 'NE'),
  ('Associação deLactobacillus plantarum(CECT 7484),Lactobacillus plantarum(CECT 7485) ePediococcus acidilactici(CECT 7483)', 'UFC', NULL, 'NE'),
  ('Bacillus clausii O/C (CNCM-I-276), SIN (CNCM-I-275), N/R (CNCM-I-274), T (CNCM-I-273)', 'Esporos/ porção', NULL, 'NE'),
  ('Bacillus clausii UBBC-07', 'UFC', NULL, 'NE'),
  ('Bacillus coagulansSNZ 1969 (MTCC 5724; BCCM LMG S-27484; ATCC 3560)', 'Esporos/ porção', NULL, 'NE'),
  ('Bacillus subtilis HU58 (NCIMB 30283)', 'UFC', NULL, 'NE'),
  ('Bifidobacterium animalis subsp. lactis HN019 (ATCC SD5674)', 'UFC', NULL, 'NE'),
  ('Bifidobacterium animalis subsp. lactis BB12 (DSM 15954)', 'UFC', NULL, 'NE'),
  ('Bifidobacterium animalis subsp. lactis LAFTI® B94', 'UFC', NULL, 'NE'),
  ('Bifidobacterium lactis NCC 2818', 'UFC', NULL, 'NA'),
  ('Lactobacillus acidophillusDDS-1® (NCIMB 30333)', 'UFC', NULL, 'NE'),
  ('Lactobacillus acidophilus NCFM (ATCC SD5221)', 'UFC', NULL, 'NE'),
  ('Lactobacillus coryniformis CECT 5711', 'UFC', NULL, 'NE'),
  ('Lactobacillus gasseri BNR17 (KCTC 10902BP)', 'UFC', NULL, 'NE'),
  ('Lactobacillus rhamnosus GG (DSM 33156)', 'UFC', NULL, 'NE'),
  ('Lactobacillus rhamnosus GG (ATCC 53103)', 'UFC', NULL, 'NE'),
  ('Lactobacillus rhamnosus HN001 (ATCC SD5675)', 'UFC', NULL, 'NA'),
  ('Lactobacillus rhamnosus NCC 4007', 'UFC', NULL, 'NE'),
  ('Limosilactobacillus reuteri DSM 17938', 'UFC', NULL, 'NE'),
  ('Associação de Lactobacillus rhamnosus R0011 (CNCM I-1720) e de Lactobacillus helveticus R0052 (CNCM I-1722)', 'UFC', NULL, 'NE'),
  ('Associação de Lactobacillus helveticus R0052 (CNCM I-1722) e de Bifidobacterium longum R0175 (CNCM I-3470)', 'UFC', NULL, 'NE'),
  ('Associação de Lactobacillus helveticus R0052, Bifidobacterium longum ssp. Infantis R0033 e Bifidobacterium bifidum R0071', 'UFC', NULL, 'NA'),
  ('Associação de Lactobacillus plantarum CECT 7527, Lactobacillus plantarum CECT 7528 e Lactobacillus plantarum CECT 7529', 'UFC', NULL, 'NE'),
  ('Associação de Bifidobacterium lactis BI- 07 (ATCC SD5220), de Lactobacillus acidophilus NCFM (ATCC SD5221), de Bifidobacterium lactis BI-04 (ATCC SD5219) e de Lactobacillus paracasei Lpc-37 (ATCC SD5275)', 'UFC', NULL, 'NE'),
  ('Associação de Bifidobacterium longum subsp. longum BB536, Bifidobacterium longum subsp. infantis M-63 e Bifidobacterium breve M-16V', 'UFC', NULL, 'NA'),
  ('Associação de Pediococcus pentosaceus (CECT 8330) e Bifidobacterium longum (CECT 7894)', 'UFC', NULL, 'NA'),
  ('Associação de Lactobacillus acidophilus La-14 (ATCC SD5212) e Lactobacillus rhamnosus HN001 (ATCC SD5675)', 'UFC', NULL, 'NE'),
  ('Associação de Lactobacillus rhamnosus GR-1 (DSM33426) e Limosillactobacillus reuteri RC-14 (DSM33016)', 'UFC', NULL, 'NE'),
  ('Niacina ix', 'mg', 35.0, 'valor'),
  ('Beta-glucana (Euglena gracilis) viii', 'mg', 244.0, 'valor'),
  ('Extrato Aquoso de Cogumelo Shiitake (Lentinula edodes) viii', 'mg', 1800.0, 'valor'),
  ('Hidrolisado de pectina de cenoura (Daucus carotasativus) rico em ramnogalacturonano (cRG-I) viii', 'mg', 4140.0, 'valor'),
  ('Oleuropeína', 'mg', 161.0, 'valor'),
  ('Oligopeptídeos de colágeno', 'g', NULL, 'NE'),
  ('Quercetina', 'mg', 280.0, 'valor'),
  ('Bacillus subtilissubsp.inaquosorumDE111', 'UFC', NULL, 'NE'),
  ('Bifidobacterium lactisNCC 2818', 'UFC', NULL, 'NE'),
  ('Bifidobacterium longumsubsp.longumBB536', 'UFC', NULL, 'NE'),
  ('Lactobacillus plantarumDR-7', 'UFC', NULL, 'NE'),
  ('Associação deLactobacillus rhamnosus19070-2 (DSM 26357) eLactobacillus reuteriDSM 12246 (CBS 145621)', 'UFC', NULL, 'NA')
ON CONFLICT (nutrient) DO UPDATE
  SET unit = EXCLUDED.unit, max_adult = EXCLUDED.max_adult, kind = EXCLUDED.kind, updated_at = now();

-- ═══ magistral-in28-mapa.sql ═══
-- Liga cada substância do catálogo ao nutriente correspondente do Anexo IV da IN 28.
--
-- O fator diz quantas unidades do Anexo IV valem UMA unidade da substância. Onde o fator seria
-- chute, a substância fica SEM mapa: teto conferido com fator inventado é pior que teto ausente.
-- Por isso acetil-L-carnitina não entra (a fração de carnitina muda com o sal) e a vitamina C do
-- catálogo, cadastrada em %, também não.
--
-- Para minerais, o motor converte a dose para elemento pelo elemental_percent antes de comparar:
-- o fator aqui é só conversão de unidade.

UPDATE magistral_components mc SET in28_nutrient = m.nutriente, in28_factor = m.fator
  FROM (VALUES
    -- minerais (a conversão para elemento é do motor)
    ('Cálcio',                'Cálciov',        1),
    ('Cobre',                 'Cobre',          1000),   -- catálogo em mg, norma em µg
    ('Ferro',                 'Ferro',          1),
    ('Iodo',                  'Iodo',           1),
    ('Selênio',               'Selênio',        1),
    ('Selenometionina',       'Selênio',        1),
    ('Zinco quelato',         'Zinco',          1),
    ('Zinco carnosina',       'Zinco',          1),
    ('Picolinato de cromo',   'Cromo',          1),
    ('Magnésio quelato',      'Magnésio',       1),
    ('Magnésio L-treonato',   'Magnésio',       1),
    ('Magnésio dimalato',     'Magnésio',       1),
    ('Magnésio citrato',      'Magnésio',       1),
    ('Magnésio taurato',      'Magnésio',       1),
    ('Magnésio aspartato',    'Magnésio',       1),
    ('Magnésio ascorbato',    'Magnésio',       1),
    ('Magnésio carbonato',    'Magnésio',       1),
    ('Magnésio óxido',        'Magnésio',       1),
    ('Magnésio sulfato',      'Magnésio',       1),
    ('Magnésio inositol',     'Magnésio',       1),
    ('Cloreto de magnésio',   'Magnésio',       1),

    -- vitaminas
    ('Biotina',               'Biotina',        1),
    ('Riboflavina',           'Riboflavina',    1),
    ('Nicotinamida',          'Niacina',        1),
    ('Piridoxal-5-fosfato',   'Vitamina B6',    1),      -- conta a massa do P5P, superestima ~30%
    ('Metilcobalamina',       'Vitamina B12',   1),
    ('Cianocobalamina',       'Vitamina B12',   1),
    ('Metilfolato',           'Ácido fólico iv', 1),
    ('Vitamina A',            'Vitamina Ai',    0.3),     -- 1 UI de retinol = 0,3 µg RE
    ('Vitamina D3',           'Vitamina Dii',   0.025),   -- 1 µg = 40 UI
    ('Vitamina E',            'Vitamina Eiii',  1),
    ('Vitamina K2 MK-7',      'Vitamina K',     1),
    ('Palmitato de ascorbila','Vitamina C',     0.43),    -- 43% do peso é ácido ascórbico
    ('Colina',                'Colina',         1),
    ('Alfa-GPC',              'Colina',         0.4),     -- alfa-GPC tem 40% de colina

    -- aminoácidos e bioativos
    ('L-teanina',             'Teanina',        1),
    ('L-taurina',             'Taurina',        1),
    ('L-carnitina',           'L-Carnitina',    1),
    ('Creatina',              'Creatina',       1),
    ('Glicina',               'Glicina',        1),
    ('L-glutamina',           'Glutamina',      1),
    ('L-tirosina',            'Tirosina',       1),
    ('L-triptofano',          'Triptofano',     1),
    ('Mio-inositol',          'Inositol',       0.001),   -- catálogo em mg, norma em g
    ('D-quiro-inositol',      'Inositol',       0.001),
    ('Coenzima Q10',          'Coenzima Q10',   1),
    ('Ubiquinol',             'Coenzima Q10',   1),
    ('CavaQ10',               'Coenzima Q10',   1),
    ('Licopeno',              'Licopeno',       1),
    ('Ácido hialurônico',     'Ácido hialurônico', 1),
    ('GABA',                  'Ácido gama aminobutírico (GABA)', 1),
    ('MSM',                   'Metilsulfonilmetano', 1),
    ('D-ribose',              'D-ribose',       1),
    ('Ômega-3',               'EPA e DHA',      1)
  ) AS m(nome, nutriente, fator)
 WHERE mc.name = m.nome;

-- Zinco carnosina tem 23% de zinco: sem isso, 100 mg do insumo virariam 100 mg de zinco.
UPDATE magistral_components
   SET elemental_percent = 23,
       correction_note = 'A carnosina de zinco (polaprezinco) tem 23% de zinco elementar.'
 WHERE name = 'Zinco carnosina' AND elemental_percent IS NULL;

-- ═══ magistral-incompat-base.sql ═══
-- Incompatibilidades com a base, todas com mecanismo e fonte.
--
-- Fonte principal: Alves FC, Passos MMB, Melo ASP, Monteiro MSSB. Perfil dos erros de prescrições
-- de medicamentos manipulados em uma farmácia-escola. Vigil. sanit. debate 2019;7(1):5-13
-- (400 prescrições, farmácia-escola da UFRJ) — que descreve caso a caso o mecanismo de cada uma.
-- Complemento: Formulário Nacional da Farmacopeia Brasileira, 2ª ed.
--
-- Idempotente pelo par (base, substância, percentual).

DELETE FROM magistral_base_incompatibilities
 WHERE source LIKE 'Vigil. sanit. debate%' OR source LIKE 'Formulário Nacional%';

INSERT INTO magistral_base_incompatibilities
    (base_pattern, substance_pattern, min_percent, severity, mechanism, recommendation, source)
VALUES
 ('lanette', 'ácido', 10, 'warn',
  'O creme Lanette é uma emulsão aniônica: acima de mais ou menos 10%, o ativo ácido neutraliza as cargas do tensoativo e a emulsão se separa.',
  'Creme não iônico comporta ativo com carga em concentração alta.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'ureia', 30, 'warn',
  'A ureia é muito hidrossolúvel: em concentração alta ela migra para a fase aquosa e muda a proporção entre as fases, desestabilizando a emulsão.',
  'Creme não iônico, ou dividir a ureia em outra preparação.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'lactato de amônio', NULL, 'warn',
  'Sal ionizável: interage com o tensoativo aniônico do Lanette em qualquer concentração.',
  'Creme não iônico.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'PCA', NULL, 'warn',
  'O PCA de sódio é sal do ácido pirrolidona carboxílico: ioniza em água e interage com o tensoativo aniônico em qualquer concentração.',
  'Creme não iônico.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('vaselina', 'LCD', NULL, 'avoid',
  'O LCD (licor carbonis detergens) é polar e a vaselina sólida é apolar: o ativo não incorpora nem homogeneíza.',
  'Pomada base (vaselina sólida com lanolina 7:3), que tem parte hidrofílica.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('diadermina', 'ácido', NULL, 'warn',
  'A base de diadermina é parcialmente saponificada com agente alcalino: o ativo ácido neutraliza o emulsionante, a viscosidade cai e a emulsão separa.',
  'Base não iônica.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('não iônico', 'hidroquinona', NULL, 'warn',
  'A hidroquinona oxida com facilidade e a associação com creme não iônico desestabiliza na prática. O artigo registra a observação e diz que não há estudo explicando o mecanismo.',
  'Creme Lanette para a hidroquinona; se a fórmula também tiver ácidos, separar em duas preparações.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('xarope', 'cloreto de potássio', 6, 'avoid',
  'Em xarope o açúcar já está perto da saturação e sobra pouca água livre: acima de 6% o cloreto de potássio não solubiliza e a dose por tomada passa a variar.',
  'Manter em 6%, que é o limite do Formulário Nacional.',
  'Formulário Nacional da Farmacopeia Brasileira, 2ª ed.');

-- Par ativo x ativo documentado na mesma fonte.

WITH a AS (SELECT id FROM magistral_components WHERE name = 'Cetoconazol'),
     b AS (SELECT id FROM magistral_components WHERE name = 'Ureia')
INSERT INTO magistral_incompatibilities (id, component_a_id, component_b_id, severity, mechanism, note)
SELECT uuid_generate_v7(), a.id, b.id, 'warn',
       'a associação muda de cor na prática, o que sugere queda de teor ou formação de outro composto',
       'A farmácia-escola da UFRJ deixou de manipular os dois juntos por causa das reclamações de cor. O artigo registra que a literatura não descreve o mecanismo.'
  FROM a, b
 WHERE NOT EXISTS (SELECT 1 FROM magistral_incompatibilities i WHERE i.component_a_id=a.id AND i.component_b_id=b.id);

-- ═══ magistral-formulas-base-seed.sql ═══
-- Fórmulas-base vindas dos formulários magistrais publicados (pesquisa externa).
--
-- POR QUE ESTE ARQUIVO EXISTE: o material da pós descreve fórmula completa em algumas frentes
-- (sono, mitocôndria, glicemia, cognição) e em outras discute a substância isolada sem fechar a
-- associação. Onde o RAG não fechou fórmula, estas vêm de formulários publicados por farmácias de
-- manipulação brasileiras, com a fonte anotada em `notes`.
--
-- Todas nascem PARA CONFERÊNCIA. São ponto de partida do repertório, não protocolo: as doses são
-- as do formulário consultado, e quem decide o que fica é o médico.
--
-- Nenhuma regra de dose é criada aqui. Regra tem trava de piso e teto e é decisão clínica —
-- cadastra-se na tela de fórmulas-base.
--
-- Idempotente: não duplica se rodar de novo.



WITH nova AS (
  INSERT INTO public.magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT * FROM (VALUES
    (uuidv7(), 'Sono completo',
     'Associação noturna que cobre precursor de serotonina, relaxamento muscular e modulação gabaérgica.',
     E'insônia de indução e manutenção\nquando o precursor isolado não basta',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula ao deitar', 60,
     'Formulário de farmácia magistral (Farmácia João Falcão). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Ansiedade diurna',
     'Associação para ansiedade com componente gabaérgico e fitoterápico, sem sedação intensa.',
     E'ansiedade diurna\ncoadjuvante do sono',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula 1 a 2 vezes ao dia', 60,
     'Formulário de farmácia magistral (BioVittare). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Resistência insulínica',
     'Associação de berberina, cromo e antioxidante para suporte do controle glicêmico.',
     E'resistência insulínica\nsíndrome metabólica\nSOP com hiperinsulinemia',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 dose após o almoço e 1 após o jantar', 90,
     'Formulário de farmácia magistral (Farmacam). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Antioxidante e imunidade',
     'Associação antioxidante ampla, com cofatores e vitaminas lipossolúveis.',
     E'suporte antioxidante\nimunidade\nestresse oxidativo aumentado',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário de farmácia magistral (Invictus / Farmabotânica). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Suporte intestinal em sachê',
     'Prebióticos com glutamina para suporte de mucosa e microbiota.',
     E'permeabilidade intestinal\ndisbiose\nconstipação',
     'sachê', 'internal', 'oral', 'Veículo qsp 1 sachê', 30::numeric, 'sachês',
     '1 sachê ao dia, dissolvido em água', 60,
     'Formulário de farmácia magistral (Beleza Saúde). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Pele, cabelo e unhas',
     'Silício orgânico com colágeno, biotina e vitamina C para tecido conjuntivo e anexos.',
     E'queda de cabelo e unhas frágeis\nqualidade de pele\nsuporte de colágeno',
     'sachê', 'internal', 'oral', 'Veículo qsp 1 sachê', 30::numeric, 'sachês',
     '1 sachê ao dia', 90,
     'Formulário de farmácia magistral (Biostevi / Natuformulas). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Climatério',
     'Associação fitoterápica para sintomas vasomotores e humor no climatério.',
     E'fogachos\nirritabilidade e sono no climatério',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula 1 a 2 vezes ao dia', 90,
     'Formulário de farmácia magistral (Pharmac / BioVittare). Doses do formulário; conferir antes de adotar.')
  ) AS v(id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
         quantity_to_dispense, quantity_unit, posology, duration, notes)
  WHERE NOT EXISTS (
    SELECT 1 FROM public.magistral_formula_templates t
    WHERE lower(t.name) = lower(v.name) AND t.deleted_at IS NULL
  )
  RETURNING id, name
)
INSERT INTO public.magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category)
SELECT uuidv7(), nova.id, c.ord, c.substance, c.qty, c.unit, 'simple'
FROM nova
JOIN (VALUES
  ('Sono completo', 0, 'Melatonina',            5::numeric,   'mg'),
  ('Sono completo', 1, '5-HTP',                 100::numeric, 'mg'),
  ('Sono completo', 2, 'Magnésio L-treonato',   300::numeric, 'mg'),
  ('Sono completo', 3, 'Glicina',               75::numeric,  'mg'),
  ('Sono completo', 4, 'L-teanina',             100::numeric, 'mg'),
  ('Sono completo', 5, 'Metilcobalamina',       100::numeric, 'mcg'),

  ('Ansiedade diurna', 0, '5-HTP',              100::numeric, 'mg'),
  ('Ansiedade diurna', 1, 'L-teanina',          100::numeric, 'mg'),
  ('Ansiedade diurna', 2, 'Magnésio quelato',   260::numeric, 'mg'),
  ('Ansiedade diurna', 3, 'Piridoxal-5-fosfato', 20::numeric, 'mg'),
  ('Ansiedade diurna', 4, 'Valeriana',          250::numeric, 'mg'),

  ('Resistência insulínica', 0, 'Berberina',            250::numeric, 'mg'),
  ('Resistência insulínica', 1, 'Ácido alfa-lipoico',   150::numeric, 'mg'),
  ('Resistência insulínica', 2, 'Picolinato de cromo',  100::numeric, 'mcg'),
  ('Resistência insulínica', 3, 'Ginseng brasileiro',   60::numeric,  'mg'),
  ('Resistência insulínica', 4, 'Gymnema silvestre',    25::numeric,  'mg'),
  ('Resistência insulínica', 5, 'Metilcobalamina',      12.5::numeric,'mcg'),

  ('Antioxidante e imunidade', 0, 'Coenzima Q10',       100::numeric, 'mg'),
  ('Antioxidante e imunidade', 1, 'Vitamina C',         100::numeric, 'mg'),
  ('Antioxidante e imunidade', 2, 'Curcumina',          100::numeric, 'mg'),
  ('Antioxidante e imunidade', 3, 'Trans-resveratrol',  100::numeric, 'mg'),
  ('Antioxidante e imunidade', 4, 'N-acetilcisteína',   200::numeric, 'mg'),
  ('Antioxidante e imunidade', 5, 'Ácido alfa-lipoico', 100::numeric, 'mg'),
  ('Antioxidante e imunidade', 6, 'Extrato de chá verde', 50::numeric,'mg'),
  ('Antioxidante e imunidade', 7, 'Vitamina D3',        2000::numeric,'UI'),

  ('Suporte intestinal em sachê', 0, 'FOS',         2::numeric, 'g'),
  ('Suporte intestinal em sachê', 1, 'Inulina',     2::numeric, 'g'),
  ('Suporte intestinal em sachê', 2, 'XOS',         2::numeric, 'g'),
  ('Suporte intestinal em sachê', 3, 'L-glutamina', 1::numeric, 'g'),

  ('Pele, cabelo e unhas', 0, 'Silício orgânico',   300::numeric, 'mg'),
  ('Pele, cabelo e unhas', 1, 'Biotina',            500::numeric, 'mcg'),
  ('Pele, cabelo e unhas', 2, 'Vitamina C',         300::numeric, 'mg'),
  ('Pele, cabelo e unhas', 3, 'Ácido hialurônico',  100::numeric, 'mg'),
  ('Pele, cabelo e unhas', 4, 'Colágeno Verisol',   2.5::numeric, 'g'),

  ('Climatério', 0, 'Isoflavona',          80::numeric,  'mg'),
  ('Climatério', 1, 'Amora',               500::numeric, 'mg'),
  ('Climatério', 2, 'Cimicifuga racemosa', 40::numeric,  'mg'),
  ('Climatério', 3, 'Vitex agnus-castus',  20::numeric,  'mg'),
  ('Climatério', 4, 'Ashwagandha',         80::numeric,  'mg'),
  ('Climatério', 5, '5-HTP',               50::numeric,  'mg')
) AS c(formula, ord, substance, qty, unit) ON c.formula = nova.name;

-- ═══ magistral-formulas-parceiras-curadas.sql ═══
-- Fórmulas do formulário das farmácias parceiras, CURADAS antes de entrar.
--
-- O formulário traz 80 fórmulas. Não entram todas: entram as que cobrem frentes que o Dr. Getúlio
-- atende e que sobrevivem à conferência. Três coisas foram feitas em cada uma:
--
--   1. FORMAS TROCADAS pelas que ele usa: vitamina C → palmitato de ascorbila, B12 →
--      metilcobalamina, B6 → piridoxal-5-fosfato, folato → metilfolato, selênio → selenometionina,
--      coenzima Q10 → CavaQ10.
--   2. ERROS DE UNIDADE CORRIGIDOS. O formulário traz "SELÊNIO 30mg" na Anti-Ox Maxi Ultra (é
--      mcg: 30 mg seriam ~100 vezes o teto de suplemento e dose tóxica) e "VIT A 50mg" na de
--      cãimbras (50 mg de retinol seriam ~166.000 UI).
--   3. DUPLICIDADE E TETO DE B6. A fórmula "Energizante" soma piridoxal-5-fosfato 7 mg com
--      vitamina B6 100 mg — a mesma vitamina duas vezes, ~107 mg/dia, acima do teto da IN 28
--      (98,6 mg) e na faixa associada a neuropatia sensitiva em uso crônico. Aqui vai só P5P 25 mg.
--
-- Cada fórmula guarda em `notes` o que foi alterado e por quê. Todas entram para conferência.

WITH nova AS (
  INSERT INTO public.magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT * FROM (VALUES
    (uuidv7(), 'Cortisol elevado e estresse',
     'Adaptógenos com L-teanina para modulação do eixo HPA em estresse com cortisol elevado.',
     E'estresse com cortisol alto\nansiedade diurna\ncoadjuvante do sono',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',60::numeric,'cápsulas',
     '1 cápsula 2 vezes ao dia', 60,
     'Formulário das farmácias parceiras, sem alteração: as três substâncias e as doses são as do formulário.'),

    (uuidv7(), 'Concentração e memória',
     'Suporte cognitivo com adaptógeno, colina e coenzima Q10 para foco e memória.',
     E'queixa de foco e memória\nfadiga mental',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula pela manhã, após o café', 60,
     'Formulário das parceiras, com coenzima Q10 trocada por CavaQ10.'),

    (uuidv7(), 'Esteatose e resistência insulínica',
     'Morosil com ácido alfa-lipoico e cofatores para esteatose hepática, resistência insulínica e dislipidemia.',
     E'esteatose hepática\nresistência insulínica\ndislipidemia',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',60::numeric,'cápsulas',
     '1 cápsula 2 vezes ao dia', 90,
     'Formulário das parceiras. Vitamina C trocada por palmitato de ascorbila e B6 por P5P. Morosil mantido em 200 mg 2x (400 mg/dia, que é a dose dos ensaios).'),

    (uuidv7(), 'Antioxidante amplo',
     'Associação antioxidante ampla com cofatores minerais, polifenóis e suporte mitocondrial.',
     E'estresse oxidativo\nsuporte antienvelhecimento',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário das parceiras (Anti-Ox Maxi Ultra), com três correções: selênio de 30 mg para 30 mcg de selenometionina (o formulário traz mg, que seria dose tóxica), vitamina C para palmitato de ascorbila e CoQ10 para CavaQ10.'),

    (uuidv7(), 'Hipotireoidismo, suporte de cofatores',
     'Cofatores da síntese e conversão de hormônio tireoidiano: iodo, selênio, zinco, tirosina e vitamina A.',
     E'hipotireoidismo, suporte nutricional\nconversão de T4 em T3',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário das parceiras, com selênio quelato trocado por selenometionina e vitamina C por palmitato de ascorbila. Iodo mantido em 100 mcg: em tireoidite autoimune, avaliar antes de repor.'),

    (uuidv7(), 'Insônia fitoterápica',
     'Associação fitoterápica para indução do sono, sem melatonina.',
     E'insônia de indução\nquando se prefere evitar melatonina',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 dose antes de dormir', 30,
     'Formulário das parceiras, sem alteração.'),

    (uuidv7(), 'Fadiga pós atividade física',
     'Suporte mitocondrial e de recuperação com carnitina, D-ribose, aminoácidos e cofatores.',
     E'fadiga após esforço\nrecuperação muscular',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 60,
     'Formulário das parceiras, enxugada: das 19 substâncias originais ficaram as 10 com dose relevante. B6 virou P5P, B12 virou metilcobalamina, selênio virou selenometionina, CoQ10 virou CavaQ10 e a dose de CoQ10 subiu de 10 para 50 mg, porque 10 mg não tem efeito descrito.')
  ) AS v(id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
         quantity_to_dispense, quantity_unit, posology, duration, notes)
  WHERE NOT EXISTS (
    SELECT 1 FROM public.magistral_formula_templates t
    WHERE lower(t.name) = lower(v.name) AND t.deleted_at IS NULL
  )
  RETURNING id, name
)
INSERT INTO public.magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuidv7(), nova.id, c.ord, c.substance, c.qty, c.unit, 'simple', c.nota, c.elem
FROM nova JOIN (VALUES
  ('Cortisol elevado e estresse',0,'Relora',150::numeric,'mg','',false),
  ('Cortisol elevado e estresse',1,'L-teanina',200::numeric,'mg','',false),
  ('Cortisol elevado e estresse',2,'Ashwagandha',200::numeric,'mg','',false),

  ('Concentração e memória',0,'Teacrina',150::numeric,'mg','',false),
  ('Concentração e memória',1,'Rhodiola rosea',150::numeric,'mg','',false),
  ('Concentração e memória',2,'Fosfatidilcolina',100::numeric,'mg','',false),
  ('Concentração e memória',3,'CavaQ10',50::numeric,'mg','troca do prescritor',false),

  ('Esteatose e resistência insulínica',0,'Morosil',200::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',1,'Ácido alfa-lipoico',75::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',2,'Piridoxal-5-fosfato',3::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',3,'Palmitato de ascorbila',60::numeric,'mg','troca do prescritor',false),

  ('Antioxidante amplo',0,'Ácido alfa-lipoico',50::numeric,'mg','',false),
  ('Antioxidante amplo',1,'CavaQ10',40::numeric,'mg','troca do prescritor',false),
  ('Antioxidante amplo',2,'N-acetilcisteína',100::numeric,'mg','no lugar da cisteína do formulário',false),
  ('Antioxidante amplo',3,'Trans-resveratrol',5::numeric,'mg','',false),
  ('Antioxidante amplo',4,'Licopeno',5::numeric,'mg','',false),
  ('Antioxidante amplo',5,'PQQ',10::numeric,'mg','',false),
  ('Antioxidante amplo',6,'Selenometionina',30::numeric,'mcg','corrigido: o formulário trazia 30 mg',true),
  ('Antioxidante amplo',7,'Zinco quelato',15::numeric,'mg','',true),
  ('Antioxidante amplo',8,'Cobre',1::numeric,'mg','',true),
  ('Antioxidante amplo',9,'Riboflavina',10::numeric,'mg','',false),
  ('Antioxidante amplo',10,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false),
  ('Antioxidante amplo',11,'Vitamina E',45::numeric,'UI','',false),

  ('Hipotireoidismo, suporte de cofatores',0,'Iodo',100::numeric,'mcg','avaliar antes em tireoidite autoimune',true),
  ('Hipotireoidismo, suporte de cofatores',1,'Selenometionina',100::numeric,'mcg','',true),
  ('Hipotireoidismo, suporte de cofatores',2,'Zinco quelato',15::numeric,'mg','',true),
  ('Hipotireoidismo, suporte de cofatores',3,'L-tirosina',100::numeric,'mg','',false),
  ('Hipotireoidismo, suporte de cofatores',4,'Vitamina A',1000::numeric,'UI','',false),
  ('Hipotireoidismo, suporte de cofatores',5,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false),
  ('Hipotireoidismo, suporte de cofatores',6,'Vitamina D3',50::numeric,'UI','dose simbólica no formulário; ajustar pelo exame',false),

  ('Insônia fitoterápica',0,'Passiflora',120::numeric,'mg','',false),
  ('Insônia fitoterápica',1,'Valeriana',50::numeric,'mg','',false),
  ('Insônia fitoterápica',2,'5-HTP',75::numeric,'mg','griffonia do formulário',false),

  ('Fadiga pós atividade física',0,'L-carnitina',500::numeric,'mg','',false),
  ('Fadiga pós atividade física',1,'CavaQ10',50::numeric,'mg','formulário trazia 10 mg de CoQ10',false),
  ('Fadiga pós atividade física',2,'NADH',5::numeric,'mg','',false),
  ('Fadiga pós atividade física',3,'L-taurina',100::numeric,'mg','',false),
  ('Fadiga pós atividade física',4,'Magnésio quelato',100::numeric,'mg','',true),
  ('Fadiga pós atividade física',5,'Zinco quelato',15::numeric,'mg','',true),
  ('Fadiga pós atividade física',6,'Selenometionina',30::numeric,'mcg','',true),
  ('Fadiga pós atividade física',7,'Piridoxal-5-fosfato',10::numeric,'mg','no lugar de B6',false),
  ('Fadiga pós atividade física',8,'Metilcobalamina',100::numeric,'mcg','no lugar de B12',false),
  ('Fadiga pós atividade física',9,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false)
) AS c(formula, ord, substance, qty, unit, nota, elem) ON c.formula = nova.name;

-- ═══ magistral-formulario-parceiras-completo.sql ═══
-- As fórmulas do formulário das parceiras. Parser conferido contra o sumário do documento;
-- nomes canonizados contra o catálogo; formas preferidas do prescritor aplicadas depois, em
-- magistral-formulario-correcoes.sql. last_review NULO: nenhuma foi conferida.

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 30+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 30+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Extrato de semente de uva', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 40+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 40+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Glycoxil', 50.0::numeric, 'mg'),
    (3, 'Green coffee', 30.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Formula da beleza 50+ anos', 'Fórmula do formulário das parceiras, seção Beleza.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 sache ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Formula da beleza 50+ anos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Verisol', 2.5::numeric, 'g'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'Silício orgânico', 150.0::numeric, 'mg'),
    (3, 'Green coffee', 30.0::numeric, 'mg'),
    (4, 'Glycoxil', 50.0::numeric, 'mg'),
    (5, 'Licopeno', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Peeling oral', 'Fórmula do formulário das parceiras, seção Beleza.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Peeling oral' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Oli-Ola', 300.0::numeric, 'mg'),
    (1, 'Nutricolin', 100.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Fosfolipídeos de caviar', 100.0::numeric, 'mg'),
    (4, 'Licopeno', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Celulite I', 'Fórmula do formulário das parceiras, seção Celulite.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia, após café da manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Celulite I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Silício orgânico', 100.0::numeric, 'mg'),
    (1, 'Bio Arct', 100.0::numeric, 'mg'),
    (2, 'DMAE bitartarato', 100.0::numeric, 'mg'),
    (3, 'Castanha-da-índia', 150.0::numeric, 'mg'),
    (4, 'Centella asiatica', 150.0::numeric, 'mg'),
    (5, 'Rutina', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Celulite II', 'Fórmula do formulário das parceiras, seção Celulite.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose pela manhã após o café', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Celulite II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Dimpless', 40.0::numeric, 'mg'),
    (1, 'Ásiaticosídeo', 30.0::numeric, 'mg'),
    (2, 'Ácido alfa-lipoico', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Aumento do metabolismo', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Aumento do metabolismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Glisodim', 100.0::numeric, 'mg'),
    (1, 'Extrato de chá verde', 75.0::numeric, 'mg'),
    (2, 'Lactobacillus plantarum', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Compulsao por doces e carboidratos', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula pela manhã e à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Compulsao por doces e carboidratos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Açafrão padronizado', 90.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (3, 'Rhodiola rosea', 100.0::numeric, 'mg'),
    (4, 'L-teanina', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Redução de medidas', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula pela manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Redução de medidas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 400.0::numeric, 'mg'),
    (1, 'Ácido alfa-lipoico', 75.0::numeric, 'mg'),
    (2, 'Cactinea', 1.0::numeric, 'g'),
    (3, 'Vitamina C', 100.0::numeric, 'mg'),
    (4, 'Magnésio quelato', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Redução gordura, apetite e ação diurética', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Redução gordura, apetite e ação diurética' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Koubo', 200.0::numeric, 'mg'),
    (1, 'Morosil', 250.0::numeric, 'mg'),
    (2, 'Açafrão padronizado', 90.0::numeric, 'mg'),
    (3, 'Cactinea', 500.0::numeric, 'mg'),
    (4, 'Hibisco', 150.0::numeric, 'mg'),
    (5, 'Abacateiro', 250.0::numeric, 'mg'),
    (6, 'Cavalinha', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogênica acelera metabolismo', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogênica acelera metabolismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 250.0::numeric, 'mcg'),
    (1, 'Teacrina', 50.0::numeric, 'mg'),
    (2, 'Ácido hidroxicítrico', 150.0::numeric, 'mg'),
    (3, 'Gymnema silvestre', 100.0::numeric, 'mg'),
    (4, 'Extrato de chá verde', 250.0::numeric, 'mg'),
    (5, 'Cactinea', 500.0::numeric, 'mg'),
    (6, 'Gengibre', 100.0::numeric, 'mg'),
    (7, 'Capsiate', 3.0::numeric, 'mg'),
    (8, 'Cássia angustifólia', 150.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogênica queima gordura', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogênica queima gordura' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'ID-alG', 100.0::numeric, 'mg'),
    (1, 'Meratrim', 200.0::numeric, 'mg'),
    (2, 'Lowat', 100.0::numeric, 'mg'),
    (3, 'Ácido hidroxicítrico', 200.0::numeric, 'mg'),
    (4, 'Gymnema silvestre', 100.0::numeric, 'mg'),
    (5, 'Green coffee', 100.0::numeric, 'mg'),
    (6, 'Cavalinha', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Queima gordura e ganho massa', 'Fórmula do formulário das parceiras, seção Emagrecimento.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Queima gordura e ganho massa' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Coleus forskohlii', 300.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Gripes, resfriados e herpes', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Gripes, resfriados e herpes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Equinácea', 150.0::numeric, 'mg'),
    (1, 'L-lisina', 300.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Chlorella', 100.0::numeric, 'mg'),
    (4, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Imunomoduladora', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Imunomoduladora' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Epicor', 50.0::numeric, 'mg'),
    (1, 'L-lisina', 300.0::numeric, 'mg'),
    (2, 'Vitamina C', 100.0::numeric, 'mg'),
    (3, 'Romã', 100.0::numeric, 'mg'),
    (4, 'Zinco quelato', 15.0::numeric, 'mg'),
    (5, 'Extrato de semente de uva', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Fortalecimento sistema imunológico', 'Fórmula do formulário das parceiras, seção Imunomodulação.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Fortalecimento sistema imunológico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Astragalus', 100.0::numeric, 'mg'),
    (1, 'Curcumina', 200.0::numeric, 'mg'),
    (2, 'Cyanotis vagas', 200.0::numeric, 'mg'),
    (3, 'Extrato de chá verde', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Detox', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose pela manhã após o café', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Detox' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Chlorella', 250.0::numeric, 'mg'),
    (1, 'Folha de oliveira', 200.0::numeric, 'mg'),
    (2, 'Gengibre', 100.0::numeric, 'mg'),
    (3, 'Green coffee', 50.0::numeric, 'mg'),
    (4, 'Altilix', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Controle do apetite', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Controle do apetite' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Koubo', 200.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Citrus aurantium', 400.0::numeric, 'mg'),
    (3, 'Ácido hidroxicítrico', 300.0::numeric, 'mg'),
    (4, 'Gymnema silvestre', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Termogenico – queima gordura', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 1h antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Termogenico – queima gordura' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 250.0::numeric, 'mcg'),
    (1, 'Sinetrol', 500.0::numeric, 'mg'),
    (2, 'Citrus aurantium', 200.0::numeric, 'mg'),
    (3, 'Morosil', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pré treino i– fornecedor de energia', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 a 2 cáps antes do treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pré treino i– fornecedor de energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Teacrina', 100.0::numeric, 'mg'),
    (1, 'L-taurina', 200.0::numeric, 'mg'),
    (2, 'AAKG', 500.0::numeric, 'mg'),
    (3, 'L-citrulina malato', 300.0::numeric, 'mg'),
    (4, 'Beta-alanina', 500.0::numeric, 'mg'),
    (5, 'Piridoxal-5-fosfato', 5.0::numeric, 'mg'),
    (6, 'Selenometionina', 40.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pré treino II – força e explosão', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em líquido e tomar 30min antes do treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pré treino II – força e explosão' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'BCAA', 5.0::numeric, 'g'),
    (1, 'L-tirosina', 400.0::numeric, 'mg'),
    (2, 'Piperina', 5.0::numeric, 'mg'),
    (3, 'L-arginina', 2000.0::numeric, 'mg'),
    (4, 'Beta-alanina', 1000.0::numeric, 'mg'),
    (5, 'HMB', 1000.0::numeric, 'mg'),
    (6, 'D-ribose', 2.0::numeric, 'g'),
    (7, 'Palatinose', 10.0::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pós treino - recuperação muscular', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em líquido e tomar após o treino', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pós treino - recuperação muscular' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'BCAA', 5.0::numeric, 'g'),
    (1, 'Phosfator', 500.0::numeric, 'mg'),
    (2, 'Gama-orizanol', 300.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pos treino – aumento de testosterona', 'Fórmula do formulário das parceiras, seção Nutrição esportiva.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ào dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pos treino – aumento de testosterona' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido D-aspártico', 500.0::numeric, 'mg'),
    (1, 'Tribulus terrestris', 500.0::numeric, 'mg'),
    (2, 'Mucuna pruriens', 200.0::numeric, 'mg'),
    (3, 'Maca peruana', 200.0::numeric, 'mg'),
    (4, 'Pygeum africanum', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Equilibrio da flora - biodisponibilidade de nutrientes', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Equilibrio da flora - biodisponibilidade de nutrientes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus acidophilus', 1.0::numeric, 'bilhões UFC'),
    (1, 'Bifidobacterium bifidum', 1.0::numeric, 'bilhões UFC'),
    (2, 'Lactobacillus bulgaricus', 1.0::numeric, 'bilhões UFC'),
    (3, 'Lactobacillus rhamnosus', 1.0::numeric, 'bilhões UFC'),
    (4, 'Lactobacillus plantarum', 1.0::numeric, 'bilhões UFC'),
    (5, 'Lactobacillus salivarius', 1.0::numeric, 'bilhões UFC'),
    (6, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Intolerância a lactose-melhora digestão e reduz sintomas', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Intolerância a lactose-melhora digestão e reduz sintomas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus delbrueckii', 4.0::numeric, 'bilhões UFC'),
    (1, 'Lactobacillus bulgaricus', 2.0::numeric, 'bilhões UFC'),
    (2, 'Streptococcus thermophilus', 2.0::numeric, 'bilhões UFC'),
    (3, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Gerenciamento de peso', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 dose ao deitar, não ingerir com alimentos quentes', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Gerenciamento de peso' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Lactobacillus gasseri', 1.0::numeric, 'bilhões UFC'),
    (1, 'Lactobacillus reuteri', 1.0::numeric, 'bilhões UFC'),
    (2, 'Lactobacillus casei', 1.0::numeric, 'bilhões UFC'),
    (3, 'Bifidobacterium breve', 1.0::numeric, 'bilhões UFC'),
    (4, 'Bifidobacterium longum', 1.0::numeric, 'bilhões UFC'),
    (5, 'Inulina', 200.0::numeric, 'MG')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Restaura mucosa intestinal', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         7::numeric, 'sachês', 'Dissolver 1 sache em liquido e por 7 dias', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Restaura mucosa intestinal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-glutamina', 5.0::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Melhora constipação intestinal infantil', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Dissolver 1 sache em liquido e tomar diariamente', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Melhora constipação intestinal infantil' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Peg 4000 5g a', 10.0::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Melhora constipação intestinal adulto gestantes e idosos', 'Fórmula do formulário das parceiras, seção Prebiótiocos e próbioticos.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Dissolver 1 a 2 colheres de sopa em liquido e tomar diariamente', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Melhora constipação intestinal adulto gestantes e idosos' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Peg 4000 10g a', 20.0::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anemia gestacional', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula após o almoço', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anemia gestacional' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Ferro', 30.0::numeric, 'mg'),
    (3, 'Piridoxal-5-fosfato', 50.0::numeric, 'mg'),
    (4, 'Metilcobalamina', 50.0::numeric, 'mcg'),
    (5, 'Vitamina C', 100.0::numeric, 'mg'),
    (6, 'Vitamina E', 60.0::numeric, 'ui'),
    (7, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antidiabetes funcional', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia, não ingerir com líquidos', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antidiabetes funcional' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Weg lem 70(ganoderma lucidum)', 200.0::numeric, 'mg'),
    (1, 'Glisodim', 100.0::numeric, 'mg'),
    (2, 'Bifidobacterium lactis', 1.0::numeric, 'bilhões UFC'),
    (3, 'Streptococcus thermophilus', 1.0::numeric, 'bilhões UFC'),
    (4, 'Lactobacillus delbrueckii', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antiobesidade', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 cápsula ao dia, não ingerir com líquidos', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antiobesidade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Weg lem 70(ganoderma lucidum)', 200.0::numeric, 'mg'),
    (1, 'Glisodim', 200.0::numeric, 'mg'),
    (2, 'Lactobacillus gasseri', 1.0::numeric, 'bilhões UFC'),
    (3, 'Lactobacillus paracasei', 1.0::numeric, 'bilhões UFC')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anorexia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anorexia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Betaína anidra', 50.0::numeric, 'mg'),
    (2, 'Cianocobalamina', 100.0::numeric, 'mcg'),
    (3, 'Cobre', 0.5::numeric, 'mg'),
    (4, 'L-lisina', 100.0::numeric, 'mg'),
    (5, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (6, 'Nicotinamida', 5.0::numeric, 'mg'),
    (7, 'Vitamina B1', 5.0::numeric, 'mg'),
    (8, 'Riboflavina', 5.0::numeric, 'mg'),
    (9, 'Vitamina C', 100.0::numeric, 'mg'),
    (10, 'Zinco quelato', 8.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anorexia e inapetência II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anorexia e inapetência II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-carnitina', 200.0::numeric, 'mg'),
    (1, 'Buclizina', 50.0::numeric, 'mg'),
    (2, 'Ciproeptadina', 6.0::numeric, 'mg'),
    (3, 'Metilcobalamina', 0.5::numeric, 'mg'),
    (4, 'L-lisina', 100.0::numeric, 'mg'),
    (5, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (6, 'Metilfolato', 10.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ansiedade generalizada', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula de 12/12 horas', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ansiedade generalizada' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, '5-HTP', 25.0::numeric, 'mg'),
    (1, 'Ashwagandha', 200.0::numeric, 'mg'),
    (2, 'Kava-kava', 50.0::numeric, 'mg'),
    (3, 'Cálcio', 25.0::numeric, 'mg'),
    (4, 'Valeriana', 25.0::numeric, 'mg'),
    (5, 'L-glutamina', 100.0::numeric, 'mg'),
    (6, 'L-taurina', 75.0::numeric, 'mg'),
    (7, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (8, 'Melissa', 100.0::numeric, 'mg'),
    (9, 'L-teanina', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidante básico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula de 12/12 horas', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidante básico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Betacaroteno', 50.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Manganês', 1.0::numeric, 'mg'),
    (3, 'Selenometionina', 30.0::numeric, 'mcg'),
    (4, 'Vitamina C', 100.0::numeric, 'mg'),
    (5, 'Vitamina E', 30.0::numeric, 'ui'),
    (6, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidante maxi', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidante maxi' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Coenzima Q10', 40.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Zinco quelato', 15.0::numeric, 'mg'),
    (4, 'Cisteína', 100.0::numeric, 'mg'),
    (5, 'Manganês', 1.0::numeric, 'mg'),
    (6, 'NADH', 1.0::numeric, 'mg'),
    (7, 'Picnogenol', 35.0::numeric, 'mg'),
    (8, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (9, 'Selênio', 30.0::numeric, 'mcg'),
    (10, 'Dimpless', 40.0::numeric, 'mg'),
    (11, 'Riboflavina', 10.0::numeric, 'mg'),
    (12, 'Vitamina C', 100.0::numeric, 'mg'),
    (13, 'Vitamina E', 45.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anti-ox maxi ultra', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anti-ox maxi ultra' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 40.0::numeric, 'mg'),
    (3, 'Cisteína', 100.0::numeric, 'mg'),
    (4, 'Licopeno', 5.0::numeric, 'mg'),
    (5, 'Luteína', 5.0::numeric, 'mg'),
    (6, 'Manganês', 1.0::numeric, 'mg'),
    (7, 'NADH', 1.0::numeric, 'mg'),
    (8, 'Picnogenol', 20.0::numeric, 'mg'),
    (9, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (10, 'PQQ', 10.0::numeric, 'mg'),
    (11, 'Selênio', 30.0::numeric, 'mg'),
    (12, 'Riboflavina', 10.0::numeric, 'mg'),
    (13, 'Vitamina C', 100.0::numeric, 'mg'),
    (14, 'Vitamina E', 45.0::numeric, 'ui'),
    (15, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Antioxidantes para fumantes', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Antioxidantes para fumantes' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 0.5::numeric, 'mg'),
    (1, 'Cisteína', 100.0::numeric, 'mg'),
    (2, 'Luteína', 1.0::numeric, 'mg'),
    (3, 'Nicotinamida', 50.0::numeric, 'mg'),
    (4, 'Selênio', 30.0::numeric, 'mcg'),
    (5, 'L-taurina', 75.0::numeric, 'mg'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina E', 90.0::numeric, 'ui'),
    (8, 'Zinco quelato', 8.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Artrite artrose I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Artrite artrose I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Colágeno tipo II', 40.0::numeric, 'mg'),
    (1, 'Boswellia', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Artrite artrose I I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Artrite artrose I I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Sucupira', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cãimbras', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cãimbras' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 50.0::numeric, 'mg'),
    (1, 'Potássio', 5.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 30.0::numeric, 'mg'),
    (3, 'Creatina', 250.0::numeric, 'mg'),
    (4, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (5, 'MSM', 100.0::numeric, 'mg'),
    (6, 'Vitamina A', 50.0::numeric, 'mg'),
    (7, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (8, 'Vitamina C', 100.0::numeric, 'mg'),
    (9, 'Vitamina E', 90.0::numeric, 'ui'),
    (10, 'Vitamina K2 MK-7', 25.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Circulacao periférica - varizes , varicoses e hemorroidas', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Circulacao periférica - varizes , varicoses e hemorroidas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Diosmina', 250.0::numeric, 'mg'),
    (1, 'Hesperidina', 100.0::numeric, 'mg'),
    (2, 'Castanha-da-índia', 150.0::numeric, 'mg'),
    (3, 'Centella asiatica', 150.0::numeric, 'mg'),
    (4, 'Rutina', 100.0::numeric, 'mg'),
    (5, 'Ásiaticosídeo', 20.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cistite infecções urinárias', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cistite infecções urinárias' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cranberry', 500.0::numeric, 'mg'),
    (1, 'Vitamina C', 250.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Colesterol I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Colesterol I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Policosanol', 30.0::numeric, 'mg'),
    (1, 'Cissus quadrangularis', 250.0::numeric, 'mg'),
    (2, 'Trans-resveratrol', 30.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Colesterol II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose à noite', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Colesterol II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Policosanol', 30.0::numeric, 'mg'),
    (1, 'Trans-resveratrol', 30.0::numeric, 'mg'),
    (2, 'Ácido alfa-lipoico', 150.0::numeric, 'mg'),
    (3, 'Coenzima Q10', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Cortisol elevado – modulação estresse', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cáps 2x/dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Cortisol elevado – modulação estresse' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Relora', 150.0::numeric, 'mg'),
    (1, 'L-teanina', 200.0::numeric, 'mg'),
    (2, 'Ashwagandha', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes mellitus', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes mellitus' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Beta-alanina', 500.0::numeric, 'mg'),
    (2, 'Potássio', 100.0::numeric, 'mg'),
    (3, 'Cobre', 1.0::numeric, 'mg'),
    (4, 'Coenzima Q10', 25.0::numeric, 'mg'),
    (5, 'Picolinato de cromo', 25.0::numeric, 'mg'),
    (6, 'Cisteína', 50.0::numeric, 'mg'),
    (7, 'Licopeno', 3.0::numeric, 'mg'),
    (8, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (9, 'Selenometionina', 30.0::numeric, 'mcg'),
    (10, 'Vanádio', 50.0::numeric, 'mcg'),
    (11, 'Piridoxal-5-fosfato', 50.0::numeric, 'mg'),
    (12, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (13, 'Vitamina C', 100.0::numeric, 'mg'),
    (14, 'Vitamina E', 90.0::numeric, 'ui'),
    (15, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes tipo I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes tipo I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Picolinato de cromo', 50.0::numeric, 'mcg'),
    (3, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (4, 'Manganês', 1.0::numeric, 'mg'),
    (5, 'Potássio', 50.0::numeric, 'mg'),
    (6, 'Selênio', 30.0::numeric, 'mcg'),
    (7, 'Vanádio', 25.0::numeric, 'mcg'),
    (8, 'Vitamina C', 100.0::numeric, 'mg'),
    (9, 'Vitamina E', 50.0::numeric, 'ui'),
    (10, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diabetes tipo II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diabetes tipo II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 10.0::numeric, 'mg'),
    (1, 'Metilfolato', 10.0::numeric, 'mg'),
    (2, 'Magnésio aspartato', 100.0::numeric, 'mg'),
    (3, 'Faseolamina', 100.0::numeric, 'mg'),
    (4, 'Glutationa reduzida', 50.0::numeric, 'mg'),
    (5, 'L-arginina', 100.0::numeric, 'mg'),
    (6, 'Psyllium', 1.0::numeric, 'g'),
    (7, 'Selênio', 30.0::numeric, 'mcg'),
    (8, 'Vanádio', 100.0::numeric, 'mcg'),
    (9, 'Vitamina C', 100.0::numeric, 'mg'),
    (10, 'Vitamina E', 90.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Doença cardiovascular', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Doença cardiovascular' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Nattoquinase', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Doença celíaca (sachês)', 'Fórmula do formulário das parceiras, seção Patologias.', 'sachê', 'internal', 'oral', 'Veículo para sachê qsp',
         30::numeric, 'sachês', 'Tomar 1 sachê 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Doença celíaca (sachês)' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Enzimas pancreáticas', 250.0::numeric, 'mg'),
    (3, 'FOS', 250.0::numeric, 'mg'),
    (4, 'L-glutamina', 100.0::numeric, 'mg'),
    (5, 'Lecitina', 100.0::numeric, 'mg'),
    (6, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (7, 'Psyllium', 1.0::numeric, 'g'),
    (8, 'Polidextrose', 1.0::numeric, 'g'),
    (9, 'Vitamina A', 1000.0::numeric, 'ui'),
    (10, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (11, 'Vitamina C', 100.0::numeric, 'mg'),
    (12, 'Vitamina D3', 100.0::numeric, 'ui'),
    (13, 'Vitamina E', 100.0::numeric, 'ui'),
    (14, 'Zinco quelato', 100.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energético mineral (catalisadores)', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energético mineral (catalisadores)' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 100.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (4, 'Manganês', 10.0::numeric, 'mg'),
    (5, 'Picolinato de cromo', 25.0::numeric, 'mcg'),
    (6, 'Molibdênio', 25.0::numeric, 'mcg'),
    (7, 'Potássio', 50.0::numeric, 'mg'),
    (8, 'Selenometionina', 30.0::numeric, 'mcg'),
    (9, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energético vitamínico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energético vitamínico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Metilfolato', 100.0::numeric, 'mcg'),
    (1, 'Vitamina B1', 1.0::numeric, 'mg'),
    (2, 'Nicotinamida', 10.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 50.0::numeric, 'mg'),
    (4, 'Piridoxal-5-fosfato', 10.0::numeric, 'mg'),
    (5, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina E', 100.0::numeric, 'ui'),
    (8, 'Vitamina D3', 400.0::numeric, 'ui')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Energizante – combate a fadiga', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose ao dia, após café da manhã', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Energizante – combate a fadiga' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'L-taurina', 200.0::numeric, 'mg'),
    (1, 'Piridoxal-5-fosfato', 7.0::numeric, 'mg'),
    (2, 'Magnésio quelato', 100.0::numeric, 'mg'),
    (3, 'Cálcio', 200.0::numeric, 'mg'),
    (4, 'Zinco quelato', 20.0::numeric, 'mg'),
    (5, 'Piridoxal-5-fosfato', 100.0::numeric, 'mg'),
    (6, 'Griffonia', 100.0::numeric, 'mg'),
    (7, 'Coenzima Q10', 50.0::numeric, 'mg'),
    (8, 'Selenometionina', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Esteatose hepática, resistencia insulinica e dislipidemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Esteatose hepática, resistencia insulinica e dislipidemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 200.0::numeric, 'mg'),
    (1, 'Ácido alfa-lipoico', 75.0::numeric, 'mg'),
    (2, 'Piridoxal-5-fosfato', 3.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 30.0::numeric, 'mg'),
    (4, 'Vitamina C', 60.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Enzimas digestivas MIX', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia, 30 min antes do almoço e jantar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Enzimas digestivas MIX' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Alfa-amilase', 100.0::numeric, 'mg'),
    (1, 'Bromelina', 150.0::numeric, 'mg'),
    (2, 'Lactase', 85.0::numeric, 'mg'),
    (3, 'Lipase', 50.0::numeric, 'mg'),
    (4, 'Papaína', 50.0::numeric, 'mg'),
    (5, 'Protease', 120.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Flacidez', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Flacidez' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 100.0::numeric, 'mcg'),
    (1, 'L-carnitina', 100.0::numeric, 'mg'),
    (2, 'Glicina', 100.0::numeric, 'mg'),
    (3, 'L-lisina', 100.0::numeric, 'mg'),
    (4, 'L-prolina', 100.0::numeric, 'mg'),
    (5, 'Manganês', 10.0::numeric, 'mg'),
    (6, 'Silício orgânico', 100.0::numeric, 'mg'),
    (7, 'Vitamina C', 150.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Fotoproteção oral', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Fotoproteção oral' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Polypodium leucotomos', 250.0::numeric, 'mg'),
    (1, 'Betacaroteno', 15.0::numeric, 'mg'),
    (2, 'Picnogenol', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hipoglicemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hipoglicemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cobre', 1.0::numeric, 'mg'),
    (1, 'Picolinato de cromo', 50.0::numeric, 'mg'),
    (2, 'D-ribose', 250.0::numeric, 'mg'),
    (3, 'Manganês', 1.0::numeric, 'mg'),
    (4, 'Vanádio', 25.0::numeric, 'mg'),
    (5, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hiperglicemia', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hiperglicemia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 20.0::numeric, 'mg'),
    (1, 'Betaína anidra', 50.0::numeric, 'mg'),
    (2, 'Cobre', 0.5::numeric, 'mg'),
    (3, 'Picolinato de cromo', 50.0::numeric, 'mcg'),
    (4, 'Faseolamina', 100.0::numeric, 'mg'),
    (5, 'FOS', 250.0::numeric, 'mg'),
    (6, 'Gymnema silvestre', 200.0::numeric, 'mg'),
    (7, 'Cisteína', 100.0::numeric, 'mg'),
    (8, 'Manganês', 1.0::numeric, 'mg'),
    (9, 'Nicotinamida', 10.0::numeric, 'mg'),
    (10, 'Potássio', 50.0::numeric, 'mg'),
    (11, 'Vanádio', 50.0::numeric, 'mcg'),
    (12, 'Metilcobalamina', 100.0::numeric, 'mcg'),
    (13, 'Riboflavina', 25.0::numeric, 'mg'),
    (14, 'Piridoxal-5-fosfato', 25.0::numeric, 'mg'),
    (15, 'Vitamina C', 100.0::numeric, 'mg'),
    (16, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Hipotireoidismo', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Hipotireoidismo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Betacaroteno', 2.5::numeric, 'mg'),
    (1, 'Cobre', 1.0::numeric, 'mg'),
    (2, 'Iodo', 100.0::numeric, 'mcg'),
    (3, 'L-tirosina', 100.0::numeric, 'mg'),
    (4, 'Selenometionina', 30.0::numeric, 'mcg'),
    (5, 'Vitamina A', 1000.0::numeric, 'ui'),
    (6, 'Vitamina C', 100.0::numeric, 'mg'),
    (7, 'Vitamina D3', 50.0::numeric, 'ui'),
    (8, 'Zinco quelato', 15.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Insonia I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Insonia I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Passiflora', 120.0::numeric, 'mg'),
    (1, 'Melissa', 120.0::numeric, 'mg'),
    (2, 'Mulungu', 80.0::numeric, 'mg'),
    (3, 'Valeriana', 50.0::numeric, 'mg'),
    (4, 'Griffonia', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Insonia II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 dose via oral antes de dormir', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Insonia II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Melatonina', 3.0::numeric, 'mg'),
    (1, 'Magnésio quelato', 50.0::numeric, 'mg'),
    (2, 'Piridoxal-5-fosfato', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Menopausa', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Menopausa' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Amora', 500.0::numeric, 'mg'),
    (1, 'Cimicifuga racemosa', 100.0::numeric, 'mg'),
    (2, 'Isoflavona', 50.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Menopausa II', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 dose via oral 2 vezes ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Menopausa II' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Isoflavona', 80.0::numeric, 'mg'),
    (1, 'Trevo-vermelho', 40.0::numeric, 'mg'),
    (2, 'Dong quai', 80.0::numeric, 'mg'),
    (3, 'Yam mexicano', 100.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Modulação nutricional em infecções', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Modulação nutricional em infecções' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Equinácea', 25.0::numeric, 'mg'),
    (1, 'Epicor', 250.0::numeric, 'mg'),
    (2, 'Cobre', 1.0::numeric, 'mg'),
    (3, 'Quercetina', 75.0::numeric, 'mg'),
    (4, 'L-taurina', 500.0::numeric, 'mg'),
    (5, 'Vitamina C', 500.0::numeric, 'mg'),
    (6, 'Zinco quelato', 10.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose I', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula ao deitar', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose I' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 500.0::numeric, 'mg'),
    (2, 'Magnésio dimalato', 200.0::numeric, 'mg'),
    (3, 'Ácido pantotênico', 100.0::numeric, 'mg'),
    (4, 'Vitamina K2 MK-7', 50.0::numeric, 'mcg'),
    (5, 'Vitamina D3', 400.0::numeric, 'ui'),
    (6, 'Colágeno tipo II', 40.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose II hormonal', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose II hormonal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Boro', 1.0::numeric, 'mg'),
    (1, 'Cálcio', 500.0::numeric, 'mg'),
    (2, 'Isoflavona', 120.0::numeric, 'mg'),
    (3, 'Vitamina D3', 400.0::numeric, 'ui'),
    (4, 'Magnésio dimalato', 200.0::numeric, 'mg'),
    (5, 'Vitamina K2 MK-7', 50.0::numeric, 'mcg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Osteoporose – MIX vit lipolíticas', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar conforme orientação', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Osteoporose – MIX vit lipolíticas' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Vitamina D3', 1000.0::numeric, 'UI'),
    (1, 'Vitamina A', 1000.0::numeric, 'UI'),
    (2, 'Vitamina K2 MK-7', 100.0::numeric, 'mcg'),
    (3, 'Veiculo oleoso qsp', 5.0::numeric, 'Gotas')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ovário policistico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         60::numeric, 'cápsulas', 'Tomar 1 cápsula 2x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ovário policistico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'D-quiro-inositol', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Suplemento geriátrico', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Suplemento geriátrico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Ácido alfa-lipoico', 20.0::numeric, 'mg'),
    (1, 'Ashwagandha', 200.0::numeric, 'mg'),
    (2, 'Coenzima Q10', 10.0::numeric, 'mg'),
    (3, 'Ginseng', 50.0::numeric, 'mg'),
    (4, 'Fosfatidilserina', 100.0::numeric, 'mg'),
    (5, 'L-carnitina', 250.0::numeric, 'mg'),
    (6, 'L-isoleucina', 100.0::numeric, 'mg'),
    (7, 'L-leucina', 100.0::numeric, 'mg'),
    (8, 'L-valina', 100.0::numeric, 'mg'),
    (9, 'NADH', 10.0::numeric, 'mg'),
    (10, 'Picnogenol', 15.0::numeric, 'mg'),
    (11, 'Pregnenolona', 10.0::numeric, 'mg'),
    (12, 'Trans-resveratrol', 5.0::numeric, 'mg'),
    (13, 'Vinpocetina', 5.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'TPM control', 'Fórmula do formulário das parceiras, seção Patologias.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', 'Tomar 1 cápsula 1x ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'TPM control' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Açafrão padronizado', 100.0::numeric, 'mg'),
    (1, 'Griffonia', 50.0::numeric, 'mg'),
    (2, 'Metilfolato', 400.0::numeric, 'mcg'),
    (3, 'Rhodiola rosea', 200.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura anti refluxo', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura anti refluxo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de alecrim', 60.0::numeric, 'mL')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura anti flatulência', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura anti flatulência' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de hortelã', 60.0::numeric, '%'),
    (1, 'Tintura de alcachofra', 40.0::numeric, '%')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Tintura proteção gastrica MIX', 'Fórmula do formulário das parceiras, seção Patologias.', 'solução', 'internal', 'oral', 'Veículo hidroalcoólico qsp',
         30::numeric, 'ml', 'Tomar 20 gotas em água antes das principais refeições', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Tintura proteção gastrica MIX' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Tintura de funcho', 30.0::numeric, '%'),
    (1, 'Tintura de espinheira-santa', 30.0::numeric, '%'),
    (2, 'Tintura de alecrim', 40.0::numeric, '%')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Ansiedade', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Ansiedade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Griffonia', 75.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Anti –aging', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Anti –aging' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Verisol', 2.5::numeric, 'G')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Compulsão alimentar', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Compulsão alimentar' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Picolinato de cromo', 350.0::numeric, 'mcg'),
    (1, 'Açafrão padronizado', 160.0::numeric, 'mg'),
    (2, 'Koubo', 400.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Diurético', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Diurético' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Cactinea', 1000.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Queima barriga', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Queima barriga' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Morosil', 500.0::numeric, 'mg')
  ) AS c(ord, subst, qtd, un);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Saciedade', 'Fórmula do formulário das parceiras, seção Formas farmacêuticas.', 'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula',
         30::numeric, 'cápsulas', '1 dose ao dia', 30,
         'Formulário de farmácia magistral (parceiras). Doses do formulário; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Saciedade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.subst, c.qtd, c.un, 'simple', '', false
  FROM nova, (VALUES
    (0, 'Chia', 1.5::numeric, 'g')
  ) AS c(ord, subst, qtd, un);

-- ═══ magistral-glp1-formulas.sql ═══
-- As 7 fórmulas do material de suporte a análogos de GLP-1.
--
-- Cada uma vem com o objetivo clínico que o material declara. last_review NULO: são sugestões do
-- fornecedor, não conduta conferida. Idempotente pelo nome.

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · constipação', 
    'Estimular o peristaltismo, hidratar o bolo fecal e modular a microbiota. Os análogos de GLP-1 retardam o esvaziamento gástrico e a motilidade, e a constipação é o efeito mais persistente.',
    E'constipação em uso de análogo de GLP-1\nretardo de esvaziamento gástrico\ntomar com bastante líquido',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em água, preferencialmente pela manhã', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · constipação' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Psyllium',1::numeric,'g','',false),
  (1,'Magnésio quelato',200::numeric,'mg','Dose do elemento, como o material escreve (faixa citada de 150 a 600 mg).',true),
  (2,'Bifidobacterium lactis',5::numeric,'bilhões UFC','',false),
  (3,'Motility',150::numeric,'mg','Blend do fornecedor; a única substância do material sem referência externa.',false)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · diarreia',
    'Equilibrar a diversidade da microbiota, reduzir inflamação e restaurar a mucosa intestinal na diarreia funcional que acompanha o uso de análogos de GLP-1.',
    E'diarreia em uso de análogo de GLP-1\nalteração de microbiota e permeabilidade',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia, em jejum', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · diarreia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', '', false FROM nova, (VALUES
  (0,'Saccharomyces boulardii',150::numeric,'mg'),
  (1,'Bacillus clausii',2::numeric,'bilhões UFC'),
  (2,'Lactobacillus rhamnosus',2::numeric,'bilhões UFC'),
  (3,'Fibregum B',200::numeric,'mg')
) AS c(ord,s,q,u);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · digestão e gases',
    'Reduzir formação de gases, otimizar a digestão e aliviar o desconforto abdominal. A digestão incompleta e a fermentação aumentada vêm da menor motilidade.',
    E'distensão abdominal e flatulência\ndigestão incompleta por menor motilidade',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 90::numeric, 'cápsulas',
    '1 dose após as refeições principais', 30,
    'O carvão ativado adsorve de forma inespecífica: afastar pelo menos duas horas de qualquer medicamento e das outras fórmulas.',
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · digestão e gases' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'Simeticone',80::numeric,'mg',''),
  (1,'Carvão ativado',200::numeric,'mg','Adsorve fármaco e nutriente: separar por duas horas do resto.'),
  (2,'Alfa-amilase',30::numeric,'mg',''),
  (3,'Protease',100::numeric,'mg',''),
  (4,'Lipase',300::numeric,'mg','')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · saciedade e manutenção do peso',
    'Estímulo endógeno de GLP-1 e PYY para manter a regulação do apetite e a perda de peso alcançada.',
    E'manutenção do peso após o análogo\nsaciedade por estímulo de receptores intestinais',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 90::numeric, 'cápsulas',
    '1 dose 30 minutos antes das principais refeições', 30,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · saciedade e manutenção do peso' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', '', false FROM nova, (VALUES
  (0,'Berberina',300::numeric,'mg'),
  (1,'Akkermansia muciniphila',50::numeric,'mg'),
  (2,'Slendesta',100::numeric,'mg')
) AS c(ord,s,q,u);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · suporte dérmico pós-emagrecimento',
    'Estimular a biossíntese de colágeno e elastina e restaurar a firmeza cutânea. A perda rápida de gordura subcutânea leva à flacidez, inclusive facial.',
    E'flacidez após perda rápida de peso\nperda de elasticidade cutânea',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em 100 mL de água pela manhã', 90,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · suporte dérmico pós-emagrecimento' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Verisol',2.5::numeric,'g','',false),
  (1,'Palmitato de ascorbila',200::numeric,'mg','Material pede vitamina C 200 mg; a dose é do ativo e o insumo sai pelo fator de correção.',true),
  (2,'Nutricolin',100::numeric,'mg','',false)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · preservação de massa magra',
    'Preservar massa muscular e estimular a síntese proteica durante a perda rápida de peso, sobretudo em quem tem baixa ingestão proteica, sedentarismo ou idade avançada.',
    E'perda de massa magra durante o análogo\nbaixa ingestão proteica ou idade avançada',
    'sachê', 'internal', 'oral', 'Veículo para sachê qsp', 30::numeric, 'sachês',
    '1 sachê ao dia, diluído em 100 mL de líquido, preferencialmente após o treino', 90,
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · preservação de massa magra' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'PeptiStrong',1.2::numeric,'g','O material dá 2,4 g/dia como dose usual: esta fórmula entrega metade em um sachê.'),
  (1,'HMB',1.5::numeric,'g','')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'GLP-1 · ciclo capilar',
    'Reduzir eflúvio, regenerar a matriz folicular e dar resistência ao fio. A queda parece vir de deficiência nutricional, inflamação e estresse metabólico da perda rápida.',
    E'eflúvio após perda rápida de peso\nfragilidade capilar',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia', 90,
    'Contém biotina em dose alta: suspender 3 dias antes de qualquer coleta de sangue e avisar o laboratório. Acima de 5 mg/dia a biotina falseia TSH, T4 livre, troponina e hormônios em imunoensaio.',
    'Material de suporte a análogos de GLP-1 (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'GLP-1 · ciclo capilar' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Cistina',100::numeric,'mg','',false),
  (1,'Metionina',100::numeric,'mg','',false),
  (2,'Nutricolin',100::numeric,'mg','',false),
  (3,'Biotina',10::numeric,'mg','Dose do material. Interfere em imunoensaio acima de 5 mg/dia: suspender antes da coleta.',false),
  (4,'Ferro',10::numeric,'mg','Dose do elemento. Repor ferro sem ferritina medida é decisão a conferir.',true),
  (5,'Cisteína',100::numeric,'mg','',false),
  (6,'Ácido pantotênico',100::numeric,'mg','',false),
  (7,'Saw palmetto',150::numeric,'mg','',false)
) AS c(ord,s,q,u,n,e);

-- ═══ magistral-pentravan-formulas.sql ═══
-- Fórmulas do material do Pentravan (Fagron), por via transdérmica e vaginal.
--
-- Cada componente já entra com a CATEGORIA de receita que carrega: as fórmulas com
-- testosterona e oxandrolona saem como Controle Especial (lista C5), que é o que a Portaria
-- 344/98 pede e o que o sistema não fazia sozinho. last_review NULO.

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · miomatose e endometriose vaginal', 'Miodesin e Pentravan por via vaginal na miomatose uterina e na endometriose.', E'endometriose\nmiomatose uterina',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal, à noite, por até 2 meses', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · miomatose e endometriose vaginal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Miodesin', 170::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · proteção endometrial', 'Progesterona por via vaginal para proteção do endométrio.', E'proteção endometrial\núltimos 13 a 15 dias do mês',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal à noite, nos últimos 13 a 15 dias do mês', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · proteção endometrial' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Progesterona micronizada', 50::numeric, 'mg', 'simple', 'Material dá faixa de 20 a 80 mg por grama.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · endometriose com gestrinona', 'Gestrinona por via vaginal na endometriose, com os estudos de Maia Jr. em Pentravan.', E'dor de endometriose\naplicação três vezes por semana',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 g por via vaginal, 3 vezes por semana', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · endometriose com gestrinona' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Gestrinona', 5::numeric, 'mg', 'simple', 'Material dá 2,5 mg ou 5 mg por grama.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · deficiência androgênica feminina', 'Reposição androgênica feminina por via transdérmica em região de pouco pelo e pouco tecido adiposo.', E'deficiência androgênica feminina comprovada\naplicar em pulsos ou antebraços',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador, 1 vez ao dia ou em dias alternados', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · deficiência androgênica feminina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 3::numeric, 'mg', 'c5', 'Material dá faixa de 0,5 a 5 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · climatério transdérmico', 'Estradiol e estriol por via transdérmica no alívio de sintomas climatéricos.', E'sintomas climatéricos\n25 dias de uso com 5 de intervalo',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador ao dia por 25 dias, com intervalo de 5 dias', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · climatério transdérmico' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, '17-beta-estradiol', 1::numeric, 'mg', 'simple', 'Material dá 0,25 a 2 mg por mL.'),
  (1, 'Estriol', 4::numeric, 'mg', 'simple', 'Material dá 2 a 8 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · estimulante sexual feminino', 'Sildenafila por uso vulvar como estimulante sexual feminino.', E'uso sob demanda\naplicar 30 minutos antes',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL na região dos lábios vaginais 30 minutos antes da relação sexual', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · estimulante sexual feminino' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Citrato de sildenafila', 0.25::numeric, '%', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · vulvodínia', 'PEA e baclofeno por via tópica no alívio da vulvodínia.', E'vulvodínia\naplicar nas áreas afetadas',
         'vaginal', 'external', 'vaginal', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL nas áreas afetadas, 1 a 2 vezes ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · vulvodínia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Palmitoiletanolamida', 10::numeric, 'mg', 'simple', ''),
  (1, 'Baclofeno', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · mastalgia cíclica', 'Danazol por via transdérmica na mama, na mastalgia cíclica.', E'mastalgia cíclica\naplicação na mama',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 g', 60::numeric, 'g',
         'Aplicar 1 mL (1 pump) na mama, 1 vez ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · mastalgia cíclica' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Danazol', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · AMPK e longevidade', 'Metformina por via transdérmica para modulação de AMPK.', E'modulação de AMPK\nduas aplicações ao dia',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 2 vezes ao dia, em região com poucos pelos', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · AMPK e longevidade' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Metformina', 75::numeric, 'mg', 'simple', 'Material dá 50 a 100 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · flacidez e envelhecimento cutâneo', 'Silício, estriol e resveratrol por via tópica facial, na prevenção de flacidez.', E'flacidez facial\nuso em rosto e pescoço',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL com pump dosador no rosto e pescoço, 1 vez ao dia', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · flacidez e envelhecimento cutâneo' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'SiliciuMax', 5::numeric, '%', 'simple', ''),
  (1, 'Estriol', 0.3::numeric, '%', 'simple', ''),
  (2, 'Trans-resveratrol', 3::numeric, '%', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · modulação de testosterona com resveratrol', 'Testosterona com trans-resveratrol como inibidor de aromatase, por via transdérmica.', E'declínio androgênico masculino\nassociação com inibidor de aromatase',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · modulação de testosterona com resveratrol' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 50::numeric, 'mg', 'c5', ''),
  (1, 'Trans-resveratrol', 50::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · declínio androgênico masculino', 'Testosterona transdérmica no declínio androgênico masculino com deficiência comprovada.', E'declínio androgênico com deficiência documentada\naplicar em região de pouco pelo',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · declínio androgênico masculino' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 60::numeric, 'mg', 'c5', 'Material dá faixa de 40 a 90 mg por mL.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · testosterona com tadalafila', 'Testosterona associada a inibidor de fosfodiesterase-5, uso diário.', E'declínio androgênico com disfunção erétil\nuso diário',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em pulsos, antebraços ou ombros', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · testosterona com tadalafila' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Testosterona micronizada', 50::numeric, 'mg', 'c5', ''),
  (1, 'Tadalafila', 5::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · disfunção erétil sob demanda', 'Alprostadil tópico, uso sob demanda.', E'disfunção erétil\n5 a 10 pumps por aplicação',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar de 5 a 10 pumps, no mínimo 3 vezes por semana, de 5 a 30 minutos antes da atividade sexual', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · disfunção erétil sob demanda' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Alprostadil', 100::numeric, 'mcg', 'simple', 'Cada pump contém 100 mcg.')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · disfunção erétil com fentolamina', 'Alprostadil com mesilato de fentolamina, uso tópico sob demanda.', E'disfunção erétil\nintervalo mínimo de 24 horas entre aplicações',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar de 5 a 10 pumps, no mínimo 3 vezes por semana, mantendo 24 horas entre aplicações', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · disfunção erétil com fentolamina' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Alprostadil', 100::numeric, 'mcg', 'simple', ''),
  (1, 'Mesilato de fentolamina', 4::numeric, 'mg', 'simple', '')
) AS c(ord,s,q,u,cat,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Pentravan · oxandrolona em sarcopenia', 'Oxandrolona transdérmica em sarcopenia, no material do fornecedor.', E'sarcopenia com deficiência documentada\nregistrar indicação no prontuário',
         'transdérmico', 'external', 'transdérmica', 'Pentravan qsp 1 mL', 60::numeric, 'mL',
         'Aplicar 1 mL (1 pump), 1 vez ao dia, em região com poucos pelos', 60,
         'Material do Pentravan (Fagron). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Pentravan · oxandrolona em sarcopenia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, c.cat, c.n, false FROM nova, (VALUES
  (0, 'Oxandrolona', 10::numeric, 'mg', 'c5', 'A fórmula original do material se intitula ''sarcopenia e ganho de peso''; ganho de massa é finalidade vedada pela Resolução CFM 2.333/2023.')
) AS c(ord,s,q,u,cat,n);

-- ═══ magistral-arquitetura-hormonal-formulas.sql ═══
-- As três fórmulas do material de arquitetura hormonal.

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, instructions, notes)
  SELECT uuid_generate_v7(), 'Eixo androgênico · manutenção',
    'Equilíbrio de testosterona total e livre, regulação de SHBG e redução de cortisol, com suporte de zinco, magnésio e boro.',
    E'suporte androgênico fisiológico\ndisposição e desempenho',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose pela manhã, diariamente', 90,
    'Fitoterápico não é reposição: a expectativa a combinar com o paciente é de suporte, não de elevação de testosterona comparável a hormônio.',
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo androgênico · manutenção' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Testofen',50::numeric,'mg','Ensaios do ingrediente usam 300 a 600 mg/dia: a dose da fórmula fica bem abaixo.',false),
  (1,'Tribulus terrestris',300::numeric,'mg','Revisão sistemática de 2025 não sustenta elevação de testosterona.',false),
  (2,'Eurycoma longifolia',200::numeric,'mg','',false),
  (3,'Boro',3::numeric,'mg','Dose do elemento.',true),
  (4,'Zinco quelato',20::numeric,'mg','Dose do elemento.',true),
  (5,'Magnésio quelato',50::numeric,'mg','Dose do elemento.',true)
) AS c(ord,s,q,u,n,e);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Eixo adrenal · cortisol e energia',
    'Equilíbrio do eixo hipotálamo-hipófise-adrenal com adaptógenos, para regulação de cortisol, disposição e recuperação.',
    E'estresse crônico\ncortisol alto com fadiga',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose ao dia, preferencialmente à tarde', 90,
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo adrenal · cortisol e energia' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, false FROM nova, (VALUES
  (0,'Robuvit',120::numeric,'mg','Ensaios do ingrediente usam 200 a 300 mg/dia.'),
  (1,'Rhodiola rosea',150::numeric,'mg',''),
  (2,'Maca peruana',300::numeric,'mg',''),
  (3,'Teacrina',75::numeric,'mg',''),
  (4,'Ashwagandha',300::numeric,'mg','KSM-66 no material.')
) AS c(ord,s,q,u,n);

WITH nova AS (
  INSERT INTO magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT uuid_generate_v7(), 'Eixo DHEA · função adrenal',
    'Estímulo fisiológico da produção de DHEA e suporte adrenal, com ginseng, yam mexicano e cofatores.',
    E'queda de DHEA\nfadiga e baixa capacidade adaptativa',
    'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 30::numeric, 'cápsulas',
    '1 dose pela manhã, diariamente', 90,
    'Material de arquitetura hormonal (Arboretum). Sugestão do fornecedor; conferir antes de adotar.'
   WHERE NOT EXISTS (SELECT 1 FROM magistral_formula_templates t WHERE t.name = 'Eixo DHEA · função adrenal' AND t.deleted_at IS NULL)
  RETURNING id)
INSERT INTO magistral_formula_template_components (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuid_generate_v7(), nova.id, c.ord, c.s, c.q, c.u, 'simple', c.n, c.e FROM nova, (VALUES
  (0,'Panax ginseng',150::numeric,'mg','',false),
  (1,'UbiQsome',100::numeric,'mg','',false),
  (2,'Magnésio quelato',150::numeric,'mg','Dose do elemento.',true),
  (3,'Turkesterone',300::numeric,'mg','Dados humanos escassos; ver a observação da substância.',false),
  (4,'Yam mexicano',200::numeric,'mg','',false),
  (5,'Boro',3::numeric,'mg','Dose do elemento.',true)
) AS c(ord,s,q,u,n,e);

-- ═══ magistral-formulario-correcoes.sql ═══
-- Correções na importação do formulário das parceiras.
--
-- Duas naturezas distintas, e vale separar:
--   1. ERRO DE UNIDADE na fonte — cada um confirmado pela própria fonte, porque a MESMA
--      substância aparece na unidade certa em outras fórmulas do mesmo documento;
--   2. FORMA PREFERIDA do prescritor, aplicada como ele pediu.
--
-- Nada aqui é adivinhação de dose: onde não havia evidência dentro do documento, ficou como está
-- e o painel aponta.

-- ---------------------------------------------------------------------------------------------
-- 0. Catálogo
-- ---------------------------------------------------------------------------------------------

-- A vitamina C estava cadastrada em "%", herança de fórmula tópica. Em fórmula oral a unidade é mg.
UPDATE magistral_components SET default_unit = 'mg' WHERE name = 'Vitamina C' AND default_unit = '%';

-- Palmitato de ascorbila tem 43% do peso em ácido ascórbico. Sem isso, trocar "vitamina C 100 mg"
-- por "palmitato 100 mg" entregaria 43 mg de vitamina C — a troca de forma viraria corte de dose.
UPDATE magistral_components
   SET elemental_percent = 43,
       correction_note = 'Cerca de 43% do peso é ácido ascórbico: 100 mg de vitamina C equivalem a 233 mg de palmitato.'
 WHERE name = 'Palmitato de ascorbila' AND elemental_percent IS NULL;

-- ---------------------------------------------------------------------------------------------
-- 1. Erros de unidade na fonte
-- ---------------------------------------------------------------------------------------------

-- Picolinato de cromo em mg: o mesmo formulário traz cromo em mcg em sete fórmulas, com os mesmos
-- números (25 e 50). Em mg seriam 100 a 200 vezes o teto do Anexo IV (250 µg).
UPDATE magistral_formula_template_components c SET unit = 'mcg',
       note = 'Unidade corrigida de mg para mcg: o formulário traz cromo em mcg nas demais fórmulas.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Picolinato de cromo' AND c.unit = 'mg';

-- Selenometionina 50 mg: as outras oito fórmulas do documento usam 30 a 100 mcg.
UPDATE magistral_formula_template_components c SET unit = 'mcg',
       note = 'Unidade corrigida de mg para mcg: as demais fórmulas do formulário usam 30 a 100 mcg.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Selenometionina' AND c.unit = 'mg';

-- Vanádio 25 mg: o mesmo valor aparece em mcg em outra fórmula do documento.
UPDATE magistral_formula_template_components c SET unit = 'mcg',
       note = 'Unidade corrigida de mg para mcg: o mesmo valor aparece em mcg em outra fórmula.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Vanádio' AND c.unit = 'mg';

-- Vitamina A 50 mg equivaleria a cerca de 166.000 UI, dose tóxica. As outras fórmulas do mesmo
-- formulário usam 1.000 UI, que é o valor adotado aqui.
UPDATE magistral_formula_template_components c SET unit = 'UI', quantity = 1000,
       note = 'Fonte trazia 50 mg (cerca de 166.000 UI, dose tóxica). Ajustado para 1.000 UI, que é o que as demais fórmulas do formulário usam.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Vitamina A' AND c.unit = 'mg';

-- Unidade escrita em caixa baixa.
UPDATE magistral_formula_template_components SET unit = 'UI' WHERE unit = 'ui';

-- Piridoxal-5-fosfato duas vezes na "Energizante": 7 mg mais 100 mg somam 107 mg/dia de B6, acima
-- do teto de 98,6 mg e na faixa associada a neuropatia sensitiva em uso prolongado. Fica uma linha.
DELETE FROM magistral_formula_template_components c
 USING magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name ILIKE 'Energizante%'
   AND c.substance = 'Piridoxal-5-fosfato' AND c.quantity = 7;

-- ---------------------------------------------------------------------------------------------
-- 2. Formas preferidas do prescritor
-- ---------------------------------------------------------------------------------------------

UPDATE magistral_formula_template_components c SET substance = 'Metilcobalamina',
       note = 'Forma trocada de cianocobalamina para metilcobalamina, como o prescritor usa.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Cianocobalamina';

-- Vitamina C vira palmitato de ascorbila, e a dose passa a ser lida como do ATIVO: o fator de
-- correção converte para a massa do insumo, senão a troca de forma cortaria a dose para 43%.
UPDATE magistral_formula_template_components c
   SET substance = 'Palmitato de ascorbila', as_elemental = true,
       note = 'Forma trocada para palmitato de ascorbila, como o prescritor usa. A dose continua sendo de vitamina C; o insumo sai pelo fator de correção.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Vitamina C';

-- Dose escrita como faixa na fonte ("PEG 4000  5G a 10G"): o parser leu o nome grudado no
-- primeiro número. Fica o valor de baixo da faixa, e a faixa inteira vai para a observação.

UPDATE magistral_formula_template_components c SET substance='Polietilenoglicol 4000', quantity=5, unit='g',
       note='Fonte traz faixa de 5 a 10 g/dia; ficou o piso. Ajustar pela resposta.'
  FROM magistral_formula_templates t
 WHERE t.id=c.template_id AND t.name ILIKE 'Melhora constipação intestinal infantil%' AND c.substance ILIKE 'Peg 4000%';
UPDATE magistral_formula_template_components c SET substance='Polietilenoglicol 4000', quantity=10, unit='g',
       note='Fonte traz faixa de 10 a 20 g/dia; ficou o piso. Uma colher de sopa equivale a 10 g.'
  FROM magistral_formula_templates t
 WHERE t.id=c.template_id AND t.name ILIKE 'Melhora constipação intestinal adulto%' AND c.substance ILIKE 'Peg 4000%';
UPDATE magistral_formula_template_components SET unit='g' WHERE unit='G';

-- Fórmula sem posologia: a receita sairia sem dizer como tomar.

UPDATE magistral_formula_templates SET posology = '1 dose ao dia'
 WHERE deleted_at IS NULL AND coalesce(trim(posology), '') = '';

-- ═══ magistral-conferencia-zerada.sql ═══
-- O carimbo de conferência das fórmulas-base estava mentindo.
--
-- Os seeds gravaram last_review junto com a fórmula, então as 24 apareciam como revisadas sem que
-- ninguém tivesse olhado. Marca de conferência que nasce preenchida não é marca de conferência: é
-- ruído esperando para enganar alguém no dia em que a conferência de verdade começar.
--
-- Fica NULO até um humano salvar a fórmula pela tela, que é o único caminho que carimba.

UPDATE magistral_formula_templates
   SET last_review = NULL, reviewed_by = NULL
 WHERE deleted_at IS NULL;

-- ═══ magistral-regras-dose-dinamica.sql ═══
-- Regras de dose dinâmica das fórmulas-base.
--
-- Cada regra tem procedência declarada na própria nota: o que vem das aulas da pós (RAG) e o que
-- vem da literatura, inclusive quando as duas DIVERGEM. Onde diverge, a faixa segue a conduta da
-- casa e a nota diz qual é a diretriz contrária — quem prescreve precisa ver as duas.
--
-- Piso e teto são obrigatórios em toda regra; a sugestão nunca escreve na receita.
-- Idempotente: apaga a regra do componente antes de recriar.

-- ---------------------------------------------------------------------------------------------
-- 0. Correções de dado encontradas ao escrever as regras
-- ---------------------------------------------------------------------------------------------

-- Vitamina D3 de 50 UI na fórmula de hipotireoidismo: 50 UI não tem efeito biológico nenhum,
-- é erro de transcrição do formulário. A base vira 2.000 UI e a regra ajusta pelo exame.
UPDATE magistral_formula_template_components c
   SET quantity = 2000
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Hipotireoidismo, suporte de cofatores'
   AND c.substance = 'Vitamina D3' AND c.quantity = 50;

-- Magnésio quelato estava marcado como dose do ELEMENTO numa fórmula e como dose do INSUMO em
-- outra. Mesma substância, duas leituras, três vezes de diferença na cápsula. Padroniza em
-- elemento (260 mg de quelato a 30% ≈ 80 mg de magnésio).
UPDATE magistral_formula_template_components c
   SET quantity = 80, as_elemental = true
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Ansiedade diurna'
   AND c.substance = 'Magnésio quelato' AND c.as_elemental = false;

-- Magnésio glicina do sachê mitocondrial passa a ser escrito em elemento, que é a unidade em que
-- a regra por peso raciocina. 500 mg de bisglicinato a 30% ≈ 150 mg de magnésio.
UPDATE magistral_formula_template_components c
   SET quantity = 150, as_elemental = true
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND t.name = 'Sachê matinal mitocondrial'
   AND c.substance = 'Magnésio glicina' AND c.as_elemental = false;

-- Fórmula sem posologia nenhuma: a receita sairia sem dizer como tomar.
UPDATE magistral_formula_templates
   SET posology = '1 sachê ao deitar'
 WHERE name = 'Sachê noturno de relaxamento' AND coalesce(trim(posology), '') = '';

-- O "exemplo de regra" era andaime meu, não fórmula. Vira a fórmula de vitamina D de verdade.
UPDATE magistral_formula_templates
   SET name = 'Vitamina D conforme exame',
       indication = 'Reposição de vitamina D com a dose ajustada pela 25-hidroxivitamina D mais recente do paciente.',
       indication_bullets = 'Reposição de vitamina D guiada pelo exame'||chr(10)||
                            'Faixa-alvo de 40 a 60 ng/mL'||chr(10)||
                            'Reavaliar 25-OH-D em 90 dias'
 WHERE name = 'Vitamina D conforme exame (exemplo de regra)';

-- ---------------------------------------------------------------------------------------------
-- 0.1 Limpeza das regras que este arquivo recria
--
-- Statement PRÓPRIO, e não um CTE junto do INSERT: há UNIQUE em template_component_id, e o
-- INSERT dentro do mesmo statement enxerga o snapshot anterior ao DELETE — a limpeza não teria
-- acontecido ainda quando a chave fosse conferida.
-- ---------------------------------------------------------------------------------------------
DELETE FROM magistral_formula_template_rules r
 USING magistral_formula_template_components c, magistral_formula_templates t
 WHERE c.id = r.template_component_id AND t.id = c.template_id
   AND (t.name, c.substance) IN (
        ('Vitamina D conforme exame',                'Vitamina D3'),
        ('Antioxidante e imunidade',                 'Vitamina D3'),
        ('Hipotireoidismo, suporte de cofatores',    'Vitamina D3'),
        ('Sono completo',                            'Metilcobalamina'),
        ('Fadiga pós atividade física',              'Metilcobalamina'),
        ('Antioxidante amplo',                       'Zinco quelato'),
        ('Fadiga pós atividade física',              'Zinco quelato'),
        ('Hipotireoidismo, suporte de cofatores',    'Zinco quelato'),
        ('Hipotireoidismo, suporte de cofatores',    'Selenometionina'),
        ('Sachê matinal mitocondrial',               'Magnésio glicina')
   );

-- ---------------------------------------------------------------------------------------------
-- 1. Vitamina D3 por faixa de 25-hidroxivitamina D
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Vitamina D3' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Vitamina D conforme exame', 'Antioxidante e imunidade',
                      'Hipotireoidismo, suporte de cofatores')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN1BF562ED', 'ng/mL', 500, 1000, 7000, 365,
           'Alvo de 40 a 60 ng/mL, como nas aulas da pós. A diretriz da Endocrine Society de 2024 não fixa alvo e desaconselha rastreio: a divergência é deliberada. Acima de 4.000 UI/dia passa do teto de suplemento da IN 28 e vira decisão prescritiva.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 20::numeric,   7000::numeric, 'deficiência'),
      (1, 20::numeric,   30::numeric,   5000::numeric, 'insuficiência'),
      (2, 30::numeric,   40::numeric,   3000::numeric, 'abaixo do alvo'),
      (3, 40::numeric,   60::numeric,   2000::numeric, 'dentro do alvo'),
      (4, 60::numeric,   NULL::numeric, 1000::numeric, 'acima do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 2. Metilcobalamina por faixa de vitamina B12
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Metilcobalamina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Sono completo', 'Fadiga pós atividade física')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN9B054BBD', 'pg/mL', 50, 100, 1000, 365,
           'Alvo acima de 550 pg/mL, como nas aulas da pós. Na literatura, abaixo de 200 é deficiência e de 200 a 500 há deficiência funcional (a EFNS usa 500 como corte). Oral de 1.000 mcg equivale à via intramuscular. Confirmar com homocisteína.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 300::numeric,  1000::numeric, 'deficiência'),
      (1, 300::numeric,  550::numeric,  500::numeric,  'abaixo do alvo'),
      (2, 550::numeric,  NULL::numeric, 100::numeric,  'dentro do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 3. Zinco quelato por faixa de zinco sérico (dose em zinco elementar)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Zinco quelato' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name IN ('Antioxidante amplo', 'Fadiga pós atividade física',
                      'Hipotireoidismo, suporte de cofatores')
), ins AS (
    INSERT INTO magistral_formula_template_rules
        (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN7B3753DD', 'µg/dL', 5, 10, 40, 365,
           'Alvo de pelo menos 100 µg/dL, como nas aulas da pós. Teto de 40 mg/dia de zinco elementar: acima disso, por meses, compete com o cobre. A dose é do elemento; o insumo sai pelo fator de correção.'
      FROM comp
    RETURNING id
)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot
  FROM ins, (VALUES
      (0, NULL::numeric, 70::numeric,   30::numeric, 'deficiência'),
      (1, 70::numeric,   100::numeric,  20::numeric, 'abaixo do alvo'),
      (2, 100::numeric,  NULL::numeric, 10::numeric, 'dentro do alvo')
  ) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 4. Selenometionina por anti-TPO (limiar, não faixa: a conduta aqui é binária)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Selenometionina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name = 'Hipotireoidismo, suporte de cofatores'
)
INSERT INTO magistral_formula_template_rules
    (id, template_component_id, kind, lab_code, lab_unit, lab_operator, lab_threshold,
     dose_if_true, dose_if_false, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'lab_threshold', 'PLNF479B8FF', 'IU/mL', 'gt', 35,
       200, 100, 10, 50, 200, 730,
       'Anti-TPO alterado: 200 mcg de selênio elementar. Nas metanálises, a selenometionina é a única forma com efeito nos anticorpos; a queda de TSH é pequena (0,2 mIU/L) e só um terço dos braços mostrou queda de anti-TPO. Toxicidade a partir de 300 a 400 mcg/dia.'
  FROM comp;

-- ---------------------------------------------------------------------------------------------
-- 5. Magnésio por peso (sachê — a única forma que comporta a dose)
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id
      FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id
     WHERE c.substance = 'Magnésio glicina' AND c.deleted_at IS NULL AND t.deleted_at IS NULL
       AND t.name = 'Sachê matinal mitocondrial'
)
INSERT INTO magistral_formula_template_rules
    (id, template_component_id, kind, per_kg, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'per_kg', 5, 50, 200, 350, 180,
       'Cinco miligramas de magnésio elementar por quilo. Teto de 350 mg/dia de magnésio suplementar (UL do IOM); acima disso o efeito laxativo domina. A dose é do elemento; o bisglicinato sai pelo fator de correção.'
  FROM comp;

-- ═══ magistral-regras-dose-expansao.sql ═══
-- Expansão das regras de dose dinâmica: de 10 para o conjunto que cobre o repertório.
--
-- Toda regra é escrita em dose DIÁRIA. A tela divide pelas tomadas que lê da posologia e mostra a
-- conta — sem isso, uma regra de 5.000 UI/dia numa fórmula tomada duas vezes ao dia entregaria
-- 10.000. A trava também é diária, e corta antes de dividir.
--
-- Idempotente: apaga a regra do componente antes de recriar.

-- ---------------------------------------------------------------------------------------------
-- 0. Duas doses da fonte que não dá para inferir — ficam como estão, anotadas
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_formula_template_components c
   SET note = 'Fonte traz 100 mcg. Zinco terapêutico é de 10 a 30 mg: a unidade parece errada, mas o documento não dá o número certo em nenhum outro lugar. Conferir com a farmácia.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Zinco quelato' AND c.unit = 'mcg';

UPDATE magistral_formula_template_components c
   SET note = 'Fonte traz 100 UI. As demais fórmulas do documento usam 1.000 UI ou mais; a regra por exame ajusta.'
  FROM magistral_formula_templates t
 WHERE t.id = c.template_id AND c.substance = 'Vitamina D3' AND c.unit IN ('UI','ui') AND c.quantity <= 100;

-- ---------------------------------------------------------------------------------------------
-- 1. Limpeza das regras que este arquivo cria
-- ---------------------------------------------------------------------------------------------
DELETE FROM magistral_formula_template_rules r
 USING magistral_formula_template_components c, magistral_formula_templates t
 WHERE c.id = r.template_component_id AND t.id = c.template_id AND t.deleted_at IS NULL
   AND (
     (c.substance = 'Vitamina D3'              AND c.unit IN ('UI','ui')) OR
     (c.substance = 'Metilcobalamina'          AND c.unit = 'mcg')        OR
     (c.substance = 'Zinco quelato'            AND c.unit = 'mg')         OR
     (c.substance = 'Metilfolato'              AND c.unit = 'mcg')        OR
     (c.substance = 'Ferro'                    AND c.unit = 'mg')         OR
     (c.substance = 'Berberina'                AND c.unit = 'mg')         OR
     (c.substance = 'Testosterona micronizada' AND c.unit = 'mg')         OR
     (c.substance = 'Selenometionina'          AND c.unit = 'mcg' AND t.name ~* 'tireoid|hipotireo')
   );

-- ---------------------------------------------------------------------------------------------
-- 2. Vitamina D3 por faixa de 25-OH-D — todas as fórmulas que a contêm
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Vitamina D3' AND c.unit IN ('UI','ui') AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN1BF562ED', 'ng/mL', 250, 1000, 7000, 365,
      'Alvo de 40 a 60 ng/mL, como nas aulas da pós. A diretriz da Endocrine Society de 2024 não fixa alvo e desaconselha rastreio: a divergência é deliberada. Acima de 4.000 UI/dia passa do teto de suplemento da IN 28.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 20::numeric,   7000::numeric, 'deficiência'),
  (1, 20::numeric,   30::numeric,   5000::numeric, 'insuficiência'),
  (2, 30::numeric,   40::numeric,   3000::numeric, 'abaixo do alvo'),
  (3, 40::numeric,   60::numeric,   2000::numeric, 'dentro do alvo'),
  (4, 60::numeric,   NULL::numeric, 1000::numeric, 'acima do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 3. Metilcobalamina por faixa de B12
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Metilcobalamina' AND c.unit = 'mcg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN9B054BBD', 'pg/mL', 50, 100, 2000, 365,
      'Alvo acima de 550 pg/mL, como nas aulas da pós. Na literatura, abaixo de 200 é deficiência e de 200 a 500 há deficiência funcional (a EFNS usa 500). Oral de 1.000 mcg equivale à via intramuscular.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 300::numeric,  1000::numeric, 'deficiência'),
  (1, 300::numeric,  550::numeric,  500::numeric,  'abaixo do alvo'),
  (2, 550::numeric,  NULL::numeric, 100::numeric,  'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 4. Zinco por faixa de zinco sérico
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Zinco quelato' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN7B3753DD', 'µg/dL', 5, 10, 29, 365,
      'Alvo de pelo menos 100 µg/dL, como nas aulas da pós. O teto de 29,59 mg/dia é o do Anexo IV da IN 28; acima disso, por meses, o zinco compete com o cobre. A dose é do elemento.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 70::numeric,   29::numeric, 'deficiência'),
  (1, 70::numeric,   100::numeric,  20::numeric, 'abaixo do alvo'),
  (2, 100::numeric,  NULL::numeric, 10::numeric, 'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 5. Metilfolato por homocisteína
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Metilfolato' AND c.unit = 'mcg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNC01E9624', 'µmol/L', 100, 200, 1000, 365,
      'Homocisteína acima de 15 µmol/L é hiper-homocisteinemia; o alvo é abaixo de 10. Repor B12 antes ou junto: folato isolado corrige o hemograma e deixa a lesão neurológica da falta de B12 avançar.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, 15::numeric,   NULL::numeric, 1000::numeric, 'hiper-homocisteinemia'),
  (1, 10::numeric,   15::numeric,   800::numeric,  'acima do alvo'),
  (2, NULL::numeric, 10::numeric,   400::numeric,  'dentro do alvo')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 6. Ferro por ferritina — com um buraco de propósito acima do alvo
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Ferro' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNCEFB97FD', 'ng/mL', 5, 20, 60, 180,
      'Ferritina acima de 70 ng/mL não tem faixa cadastrada de propósito: repor ferro sem falta comprovada é risco, não conveniência. Em queda capilar o alvo costuma ser acima de 40 a 70. Dose em ferro elementar, longe de cálcio e zinco; dia sim, dia não absorve proporcionalmente mais.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 15::numeric, 60::numeric, 'deficiência absoluta'),
  (1, 15::numeric,   30::numeric, 45::numeric, 'deficiência'),
  (2, 30::numeric,   70::numeric, 30::numeric, 'abaixo do alvo capilar')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 7. Berberina por hemoglobina glicada
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Berberina' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLN3FC5EDA6', '%', 50, 500, 1500, 365,
      'Os ensaios usam 1 a 1,5 g/dia fracionados. A berberina inibe CYP3A4 e P-glicoproteína: conferir interação com o que o paciente já toma, sobretudo estatina, ciclosporina e anticoagulante.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, 6.4::numeric,  NULL::numeric, 1500::numeric, 'faixa de diabetes'),
  (1, 5.7::numeric,  6.4::numeric,  1000::numeric, 'pré-diabetes'),
  (2, NULL::numeric, 5.7::numeric,  500::numeric,  'normal')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 8. Testosterona por testosterona total — o buraco acima de 350 é a regra do CFM
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Testosterona micronizada' AND c.unit = 'mg' AND c.deleted_at IS NULL
), ins AS (
    INSERT INTO magistral_formula_template_rules
      (id, template_component_id, kind, lab_code, lab_unit, round_to, min_dose, max_dose, max_data_age_days, note)
    SELECT uuid_generate_v7(), comp.id, 'lab_band', 'PLNDE1A5575', 'ng/dL', 10, 20, 90, 180,
      'Acima de 350 ng/dL não há faixa cadastrada de propósito: a Resolução CFM 2.333/2023 só admite reposição com deficiência comprovada e nexo causal. O sistema não sugere dose para quem não tem indicação.'
      FROM comp RETURNING id)
INSERT INTO magistral_formula_template_rule_bands (id, rule_id, display_order, lower_bound, upper_bound, dose, label)
SELECT uuid_generate_v7(), ins.id, b.ord, b.lo, b.hi, b.dose, b.rot FROM ins, (VALUES
  (0, NULL::numeric, 250::numeric, 60::numeric, 'deficiência'),
  (1, 250::numeric,  350::numeric, 40::numeric, 'limítrofe')
) AS b(ord, lo, hi, dose, rot);

-- ---------------------------------------------------------------------------------------------
-- 9. Selenometionina por anti-TPO — só nas fórmulas de tireoide
-- ---------------------------------------------------------------------------------------------
WITH comp AS (
    SELECT c.id FROM magistral_formula_template_components c
      JOIN magistral_formula_templates t ON t.id = c.template_id AND t.deleted_at IS NULL
     WHERE c.substance = 'Selenometionina' AND c.unit = 'mcg' AND c.deleted_at IS NULL
       AND t.name ~* 'tireoid|hipotireo'
)
INSERT INTO magistral_formula_template_rules
  (id, template_component_id, kind, lab_code, lab_unit, lab_operator, lab_threshold,
   dose_if_true, dose_if_false, round_to, min_dose, max_dose, max_data_age_days, note)
SELECT uuid_generate_v7(), comp.id, 'lab_threshold', 'PLNF479B8FF', 'IU/mL', 'gt', 35,
       200, 100, 10, 50, 200, 730,
       'Anti-TPO alterado: 200 mcg de selênio elementar. Nas metanálises, a selenometionina é a única forma com efeito nos anticorpos; a queda de TSH é pequena (0,2 mIU/L) e só um terço dos braços mostrou queda de anti-TPO. Toxicidade a partir de 300 a 400 mcg/dia.'
  FROM comp;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Rollback de dado de catalogo e destrutivo de um jeito que rollback de schema nao e: depois que
-- o medico conferir uma formula, ajustar uma dose ou criar a formula dele, apagar por nome leva
-- o trabalho dele junto. E nao da para distinguir o que veio daqui do que ele mexeu.
--
-- Por isso o Down nao apaga nada. Para desfazer a carga inteira, reverta ate a 00069, que derruba
-- as tabelas — e ai a intencao de perder o conteudo esta explicita.

SELECT 1;

-- +goose StatementEnd
