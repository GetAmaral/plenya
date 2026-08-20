-- Volta o palmitato para dose do ativo, e tira as notas que viraram redundância no receituário.
--
-- ERRO CORRIGIDO AQUI: desmarquei os 25 componentes de palmitato de ascorbila achando que a dose
-- era do insumo. Não é. A fonte das parceiras escreve "VITAMINA C 100 mg" e a curadoria
-- (magistral-formulario-correcoes.sql) trocou a FORMA para palmitato mantendo a dose em vitamina
-- C, justamente para a troca não cortar a dose a 43%. Desmarcar cortava a dose de todas as 25.
--
-- O sinal do catálogo acompanha: `dose_as_elemental` é "a dose se escreve em ativo/elemento e a
-- farmácia converte", que é exatamente o caso — mineral por convenção, palmitato por decisão de
-- forma preferida.
--
-- E as notas "Dose em cobre elementar." que eu tinha posto em cada mineral saem: o PDF passou a
-- imprimir "(do elemento)" ao lado da quantidade, então a nota repetia a mesma informação no meio
-- da linha do receituário.

BEGIN;

UPDATE magistral_components SET dose_as_elemental = true WHERE name = 'Palmitato de ascorbila';

UPDATE magistral_formula_template_components c SET as_elemental = true
 WHERE c.substance = 'Palmitato de ascorbila' AND c.deleted_at IS NULL
   AND c.note LIKE 'Forma trocada para palmitato%';

-- Nota mais curta: o "(do elemento)" já diz que a dose é do ativo.
UPDATE magistral_formula_template_components SET note = 'Dose de vitamina C; o palmitato é a forma preferida do prescritor.'
 WHERE substance = 'Palmitato de ascorbila' AND note LIKE 'Forma trocada para palmitato%';

UPDATE magistral_formula_template_components SET note = ''
 WHERE note ~ '^Dose em ([a-zà-ú]+ )?element(o|ar)\.$';

COMMIT;

-- As 5 restantes vieram da mesma troca, por passes de curadoria diferentes ("troca do
-- prescritor"). Uma delas dizia na nota que a dose era do ativo e estava marcada como insumo —
-- contradição dentro da própria linha. Com `dose_as_elemental` no catálogo, palmitato é elementar
-- em todas.
UPDATE magistral_formula_template_components
   SET as_elemental = true,
       note = 'Dose de vitamina C; o palmitato é a forma preferida do prescritor.'
 WHERE substance = 'Palmitato de ascorbila' AND deleted_at IS NULL AND NOT as_elemental;
