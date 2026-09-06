-- +goose Up
-- Duas correções de catálogo que vieram do mesmo lugar: um exame lançado certo pontuou errado.
--
-- 1) EXAME SEM UNIDADE NÃO CONVERTE, E O ESCORE COMPARA NA UNIDADE ERRADA, EM SILÊNCIO.
--
-- `LabTestDefinition.ConverteParaUnidadePrincipal` precisa de um alvo. Com
-- `lab_test_definitions.unit` vazia não há alvo, então nenhuma das quatro camadas roda: o valor
-- entra CRU e `unit_conversion_status` fica NULO, ou seja, nem a marca de "não deu" aparece. O
-- motor do escore então casa o número contra faixas que estão noutra unidade.
--
-- Aconteceu com a pregnenolona. O laudo brasileiro reporta em ng/mL, o valor entrou como 1,19, e
-- as faixas do score item estão em ng/dL (33 a 248), então 1,19 caiu em "≤33 deficiência" e
-- pontuou 0 de 4. O correto eram 119 ng/dL, que é o nível 4. O buraco é a divergência entre as
-- duas pontas: `score_items.unit` estava preenchida, `lab_test_definitions.unit` não, e nada no
-- código compara as duas.
--
-- A correção preenche a unidade da definição a partir da unidade que o score item já usa, que é a
-- que as faixas assumem. Duas exclusões deliberadas:
--
--   * As descritivas ("Qualitativo", "TI-RADS", "grau", "0-100") ficam de fora, porque unidade
--     preenchida nelas faria a aritmética de prefixo SI tentar converter rótulo.
--   * `PLND7C2752F` ("Capacidade de fixação de ferro - IST") fica de fora porque guarda DUAS
--     grandezas sob o mesmo código: há resultados em µg/dL (capacidade de ligação) e em %
--     (saturação da transferrina) apontando para ele. Declarar qualquer uma das duas como a
--     unidade faria a outra virar conversão impossível. O conserto ali é separar em dois exames
--     de catálogo, que é curadoria, não migration de unidade.
UPDATE public.lab_test_definitions d SET unit = v.unidade
FROM (VALUES
  ('PLNPREGNEN1', 'ng/dL'),      -- pregnenolona: o caso que revelou o problema
  ('PLNB153CC32', 'ng/mL'),      -- IGF-1, que também é reportado em nmol/L por alguns laboratórios
  ('PLNANTIDNA1', 'IU/mL'),
  ('PLNCA15301',  'U/mL'),
  ('PLNCA19901',  'U/mL'),
  ('PLNOMEGA301', '%'),
  ('PLNSLEEPEFF', '%'),
  ('PLNRXCTR01',  '%'),
  ('PLNSLEEPIAH', 'eventos/h'),
  ('PLNSLEEPTST', 'h'),
  ('PLNRXTNOD01', 'mm')
) AS v(codigo, unidade)
WHERE d.code = v.codigo AND coalesce(btrim(d.unit), '') = '' AND d.deleted_at IS NULL;

-- As conversões da pregnenolona. O fator é a razão principal→secundária e `ConvertToMain` DIVIDE
-- por ele, como já fazem T3 reverso e dihidrotestosterona, que têm a mesma unidade principal.
--   ng/mL:  1 ng/mL = 100 ng/dL          → 0,01
--   nmol/L: massa molar 316,48 g/mol, então 1 nmol/L = 31,648 ng/dL → 1/31,648 = 0,0316
-- ng/mL a aritmética de prefixo SI resolveria sozinha; fica curada mesmo assim porque é a forma em
-- que os laudos brasileiros chegam, e curada ela não depende da rede de plausibilidade.
INSERT INTO public.lab_test_unit_conversions (id, lab_test_definition_id, main_unit, secondary_unit, conversion_factor, created_at, updated_at)
SELECT uuid_generate_v7(), d.id, 'ng/dL', v.secundaria, v.fator, now(), now()
FROM public.lab_test_definitions d
CROSS JOIN (VALUES ('ng/mL', 0.01), ('nmol/L', 0.0316)) AS v(secundaria, fator)
WHERE d.code = 'PLNPREGNEN1' AND d.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM public.lab_test_unit_conversions c
    WHERE c.lab_test_definition_id = d.id AND c.secondary_unit = v.secundaria AND c.deleted_at IS NULL
  );

-- 2) DIHIDROTESTOSTERONA SÓ PONTUA PARA HOMEM.
--
-- As faixas do item são masculinas (30 a 85 ng/dL no ótimo), mas ele está como
-- `gender = 'not_applicable'`, então vale para todo mundo. Mulher com DHT normal para mulher cai
-- em "≤30" e pontua 0. Aconteceu numa paciente de 45 anos com 9,14 ng/dL, valor sem nada de
-- errado. É o mesmo gap de dado que o filtro por sexo já resolvia noutros itens: o motor
-- (`AppliesToPatient`) sempre soube filtrar, faltava marcar o item.
--
-- Fica 'male' por ora. Faixas femininas de DHT são outra curadoria, e item sem faixa certa é pior
-- que item ausente: ausente vira lacuna declarada, errado vira ponto perdido silencioso.
--
-- EFEITO COLATERAL CONHECIDO: `MotivoDeNaoAplicar` compara a letra do item com `patient.Gender`
-- literalmente, e o cadastro admite paciente sem sexo informado (vira 'other'). Para esses, o DHT
-- passa a sair como "não aplicável: é de paciente do sexo masculino" em vez de indeterminado por
-- falta de dado, que é como o mesmo arquivo já trata a idade ausente. É melhor que a situação
-- anterior, em que toda mulher perdia ponto em silêncio, mas o motivo mente para quem não tem sexo
-- no cadastro. O conserto é no motor, não aqui.
UPDATE public.score_items
   SET gender = 'male'
 WHERE lab_test_code = 'PLN00994110'
   AND gender = 'not_applicable'
   AND deleted_at IS NULL;

-- 3) ESTA MIGRATION NÃO É RETROATIVA, E O DEPLOY PRECISA DE DOIS COMANDOS DEPOIS.
--
-- Preencher a unidade só muda o que for importado DAQUI PARA A FRENTE: `applyUnitConversion` roda
-- na ingestão, não sobre o que já está gravado. As pregnenolonas já lançadas em ng/mL continuam
-- com o valor cru e a pontuação errada, e os snapshots já calculados continuam descontando pontos
-- de DHT das mulheres. Sem os dois comandos abaixo o deploy PARECE ter corrigido e os pacientes
-- afetados seguem com escore errado:
--
--   docker exec <api> /app/reconvert-lab-units -aplicar   # idempotente: parte do valor ORIGINAL
--   docker exec <api> /app/recalc-scores                  # refaz os snapshots
--
-- Em dev os dois já rodaram junto com esta migration.

-- +goose Down
UPDATE public.score_items
   SET gender = 'not_applicable'
 WHERE lab_test_code = 'PLN00994110' AND gender = 'male' AND deleted_at IS NULL;

-- O Down desfaz só o que o Up fez. O Up é guardado (`NOT EXISTS`), então pode não ter inserido
-- nada; apagar por código e unidade removeria uma conversão curada antes, que o guarda preservou
-- de propósito. O fator identifica as duas linhas desta migration.
DELETE FROM public.lab_test_unit_conversions c
USING public.lab_test_definitions d
WHERE c.lab_test_definition_id = d.id
  AND d.code = 'PLNPREGNEN1'
  AND (   (c.secondary_unit = 'ng/mL'  AND c.conversion_factor = 0.01)
       OR (c.secondary_unit = 'nmol/L' AND c.conversion_factor = 0.0316));

-- Idem: o Up só preencheu onde estava vazio, então o Down só esvazia o que ficou com o valor que
-- o Up escreveu. Uma unidade mudada depois por outra migration ou por curadoria sobrevive.
UPDATE public.lab_test_definitions d SET unit = ''
FROM (VALUES
  ('PLNPREGNEN1', 'ng/dL'), ('PLNB153CC32', 'ng/mL'), ('PLNANTIDNA1', 'IU/mL'),
  ('PLNCA15301', 'U/mL'), ('PLNCA19901', 'U/mL'), ('PLNOMEGA301', '%'),
  ('PLNSLEEPEFF', '%'), ('PLNRXCTR01', '%'), ('PLNSLEEPIAH', 'eventos/h'),
  ('PLNSLEEPTST', 'h'), ('PLNRXTNOD01', 'mm')
) AS v(codigo, unidade)
WHERE d.code = v.codigo AND d.unit = v.unidade;
