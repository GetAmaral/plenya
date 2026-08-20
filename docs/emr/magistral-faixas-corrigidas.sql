-- Correção das faixas de dose que contradiziam o próprio texto de posologia da linha.
--
-- Como apareceram: a faixa numérica foi semeada a partir das fórmulas das parceiras (dose de UMA
-- cápsula) e o texto veio da literatura (dose do DIA). O painel comparava uma contra a outra,
-- então fórmula com dose baixa passava sem alerta e fórmula correta era acusada.
--
-- Aqui a faixa numérica passa a ser DIÁRIA em todas, com o valor da literatura. Idempotente.

BEGIN;

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

COMMIT;

-- ---------------------------------------------------------------------------------------------
-- Segunda leva: faixas cujo extremo é exatamente a dose de uma cápsula das fórmulas parceiras
-- (impressão digital de faixa semeada da própria fórmula, não da literatura).
-- ---------------------------------------------------------------------------------------------

BEGIN;

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

COMMIT;

-- ---------------------------------------------------------------------------------------------
-- Terceira leva: o catálogo discordando das regras de dose dinâmica
-- ---------------------------------------------------------------------------------------------

BEGIN;

-- Metilcobalamina: a regra por faixa de B12 prescreve 100 mcg de manutenção para quem está dentro
-- do alvo, e o catálogo dizia que 100 mcg é dose baixa. Dois subsistemas discordando da mesma
-- substância. O piso desce para 100 mcg, que é a manutenção, e o texto guarda a faixa de reposição.
UPDATE magistral_components
   SET min_dose = 100, max_dose = 2000, usual_dose = 500, dose_basis = 'por_dia'
 WHERE name = 'Metilcobalamina';

COMMIT;
