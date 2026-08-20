-- Correções na importação do formulário das parceiras.
--
-- Duas naturezas distintas, e vale separar:
--   1. ERRO DE UNIDADE na fonte — cada um confirmado pela própria fonte, porque a MESMA
--      substância aparece na unidade certa em outras fórmulas do mesmo documento;
--   2. FORMA PREFERIDA do prescritor, aplicada como ele pediu.
--
-- Nada aqui é adivinhação de dose: onde não havia evidência dentro do documento, ficou como está
-- e o painel aponta.

BEGIN;

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

COMMIT;

-- Dose escrita como faixa na fonte ("PEG 4000  5G a 10G"): o parser leu o nome grudado no
-- primeiro número. Fica o valor de baixo da faixa, e a faixa inteira vai para a observação.
BEGIN;
UPDATE magistral_formula_template_components c SET substance='Polietilenoglicol 4000', quantity=5, unit='g',
       note='Fonte traz faixa de 5 a 10 g/dia; ficou o piso. Ajustar pela resposta.'
  FROM magistral_formula_templates t
 WHERE t.id=c.template_id AND t.name ILIKE 'Melhora constipação intestinal infantil%' AND c.substance ILIKE 'Peg 4000%';
UPDATE magistral_formula_template_components c SET substance='Polietilenoglicol 4000', quantity=10, unit='g',
       note='Fonte traz faixa de 10 a 20 g/dia; ficou o piso. Uma colher de sopa equivale a 10 g.'
  FROM magistral_formula_templates t
 WHERE t.id=c.template_id AND t.name ILIKE 'Melhora constipação intestinal adulto%' AND c.substance ILIKE 'Peg 4000%';
UPDATE magistral_formula_template_components SET unit='g' WHERE unit='G';
COMMIT;

-- Fórmula sem posologia: a receita sairia sem dizer como tomar.
BEGIN;
UPDATE magistral_formula_templates SET posology = '1 dose ao dia'
 WHERE deleted_at IS NULL AND coalesce(trim(posology), '') = '';
COMMIT;
