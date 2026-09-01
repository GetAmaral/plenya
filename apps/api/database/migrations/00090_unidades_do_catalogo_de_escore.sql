-- +goose Up
-- Unidades do catálogo do escore: uma faltando e três em grandeza diferente da do exame.
--
-- Achado ao derivar as réguas da devolutiva. A unidade não é enfeite: é ela que diz se a escala do
-- item e o resultado do laboratório falam da mesma coisa.

-- 1 · Saturação de transferrina estava SEM unidade, e a régua saía com o número solto.
UPDATE public.score_items
   SET unit = '%'
 WHERE lab_test_code = 'PLND7C2752F' AND (unit IS NULL OR btrim(unit) = '');

-- 2 · Sedimento urinário: três itens têm a escala em `células/campo` (contagem por campo do
-- microscópio) enquanto o exame reporta `/µL` (concentração). São grandezas diferentes, então a
-- classificação compara números que não se comparam: um resultado de 0,5/µL cai na faixa "≤10
-- células/campo" e sai como nível ÓTIMO sem que ninguém tenha olhado nada.
--
-- Reescrever as faixas em /µL é curadoria clínica e depende do Getúlio, então NÃO se inventa
-- ponto de corte aqui. O que esta migration faz é registrar o problema onde ele é visível, e o
-- código passa a se recusar a classificar quando as unidades divergem (ScoreItem.UnitMatches).

COMMENT ON COLUMN public.score_items.unit IS
  'Unidade da escala. Tem que casar com lab_test_definitions.unit: se divergir, o item não é classificado (grandezas diferentes não se comparam). Pendente: os 3 itens do sedimento urinário estão em células/campo e o exame vem em /µL.';

-- +goose Down
UPDATE public.score_items SET unit = NULL WHERE lab_test_code = 'PLND7C2752F' AND unit = '%';
