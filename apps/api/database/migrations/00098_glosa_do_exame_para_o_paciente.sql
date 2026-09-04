-- +goose Up
-- A glosa do exame: o que ele MEDE, em até cinco palavras, na língua do paciente.
--
-- É o `sub` da régua nos decks aprovados ("estoque de ferro", "o rim filtrando", "o esforço do
-- pâncreas"), presente em 66% delas. Não existia em campo nenhum: era escrito à mão, deck a deck,
-- e por isso se perdia entre um paciente e outro.
--
-- Curada uma vez por exame, vale para sempre e para todo paciente, com custo zero de modelo. É o
-- oposto de pedir ao LLM a cada geração a mesma frase de quatro palavras.
--
-- NÃO é o nome do exame (`name` já é amigável: "Hemoglobina", "PCR ultrassensível") nem a
-- abreviação (`short_name` é "Hb", "PCR-us", que não explicam nada). É a explicação.
ALTER TABLE public.lab_test_definitions
  ADD COLUMN IF NOT EXISTS patient_gloss varchar(40) NOT NULL DEFAULT '';

COMMENT ON COLUMN public.lab_test_definitions.patient_gloss IS
  'O que o exame mede, em até cinco palavras, para o paciente ler sob o nome na régua ("estoque de ferro"). Minúsculo, sem ponto final. Vazio quando o nome já se explica.';

-- Um primeiro punhado, tirado dos dois decks aprovados: são as glosas que o médico já escreveu à
-- mão e aprovou. O resto é curadoria, e a régua funciona sem elas.
UPDATE public.lab_test_definitions SET patient_gloss = v.glosa
FROM (VALUES
  ('PLN543993C6', 'quantas partículas circulam'),
  ('PLNBB0DDBD4', 'inflamação'),
  ('PLN86387B73', 'o esforço do pâncreas'),
  ('PLN24EA4ACE', 'o esforço do pâncreas'),
  ('PLN9AF0BCF5', 'açúcar no sangue em jejum'),
  ('PLN3FC5EDA6', 'média do açúcar em 3 meses'),
  ('PLN585CE3E3', 'transporte de oxigênio'),
  ('PLN2ADF9420', 'as células de defesa'),
  ('PLN1BF562ED', 'vitamina do sol e do osso'),
  ('PLN20FCAA74', 'proteína que o rim deixa passar'),
  ('PLNBBB9A8C7', 'equilíbrio de água no corpo'),
  ('PLN74FF1276', 'sal que regula o coração'),
  ('PLN6885D35A', 'proteína que carrega e segura água')
) AS v(code, glosa)
WHERE lab_test_definitions.code = v.code;

-- +goose Down
ALTER TABLE public.lab_test_definitions DROP COLUMN IF EXISTS patient_gloss;
