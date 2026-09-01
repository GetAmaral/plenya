-- +goose Up
-- Marca no resultado o que aconteceu com a unidade na ingestão.
--
-- O conversor de unidades existe desde fevereiro e funciona: no caminho de OCR converteu 738 de
-- 738 resultados. Mas quando o par (exame, unidade) não estava na tabela curada, ele devolvia o
-- valor original SEM erro, sem log e sem marca — e o resultado seguia gravado na unidade do
-- laudo. Lá na frente o motor do escore comparava esse número contra uma escala em outra
-- grandeza e produzia um nível que ninguém tinha calculado: um `%` de linfócitos caindo numa
-- faixa em `k/µL`, um `0,5/µL` de sedimento caindo em "≤10 células/campo" e saindo como ÓTIMO.
--
-- Sem esta coluna não há como distinguir "chegou na unidade certa" de "não deu para converter",
-- e a fila de curadoria tem que ser garimpada à mão a cada vez.

ALTER TABLE public.lab_results
  ADD COLUMN IF NOT EXISTS unit_conversion_status varchar(20),
  ADD COLUMN IF NOT EXISTS unit_conversion_note   text;

COMMENT ON COLUMN public.lab_results.unit_conversion_status IS
  'ok = laudo já veio na unidade do exame; convertido = valor foi convertido (regra curada ou aritmética de prefixo); revisar = unidades diferem e não há conversão segura, valor mantido como veio.';
COMMENT ON COLUMN public.lab_results.unit_conversion_note IS
  'Por que não converteu, em português. Preenchido só quando status = revisar.';

CREATE INDEX IF NOT EXISTS idx_lab_results_unit_revisar
  ON public.lab_results (unit_conversion_status)
  WHERE unit_conversion_status = 'revisar' AND deleted_at IS NULL;

-- Duas conversões curadas estão erradas por 10x. Achadas conferindo a tabela inteira contra a
-- aritmética de prefixo: das 145 linhas, 74 batem exatamente, 69 estão fora do alcance da
-- aritmética (molar para massa exige peso molecular) e estas 2 divergem.
--
-- `ConvertToMain` divide pelo fator, então o fator é a razão principal -> secundária.
-- T3 Reverso: principal ng/dL, secundária ng/mL. 1 ng/mL = 100 ng/dL, logo o fator é 0,01.
-- Alfa-1 Globulina: principal g/dL, secundária mg/dL. 1 mg/dL = 0,001 g/dL, logo o fator é 1000.
-- Nenhum resultado chegou a passar por elas (verificado em dev e em prod), então isto é
-- preventivo, não corretivo.

UPDATE public.lab_test_unit_conversions c
   SET conversion_factor = 0.01
  FROM public.lab_test_definitions d
 WHERE d.id = c.lab_test_definition_id AND d.name = 'T3 Reverso'
   AND c.main_unit = 'ng/dL' AND c.secondary_unit = 'ng/mL' AND c.conversion_factor = 0.1;

UPDATE public.lab_test_unit_conversions c
   SET conversion_factor = 1000
  FROM public.lab_test_definitions d
 WHERE d.id = c.lab_test_definition_id AND d.name = 'Alfa-1 Globulina'
   AND c.main_unit = 'g/dL' AND c.secondary_unit = 'mg/dL' AND c.conversion_factor = 100;

-- Densidade urinária: `specific gravity` é adimensional e numericamente igual a g/mL. Não é
-- aritmética de prefixo (nenhum dos dois é grandeza com prefixo SI), então precisa da linha.
INSERT INTO public.lab_test_unit_conversions (id, lab_test_definition_id, main_unit, secondary_unit, conversion_factor, created_at, updated_at)
SELECT gen_random_uuid(), d.id, 'specific gravity', 'g/mL', 1, now(), now()
  FROM public.lab_test_definitions d
 WHERE d.name = 'Densidade Urinária (USG)' AND d.deleted_at IS NULL
   AND NOT EXISTS (SELECT 1 FROM public.lab_test_unit_conversions c
                    WHERE c.lab_test_definition_id = d.id AND c.secondary_unit = 'g/mL' AND c.deleted_at IS NULL);

-- +goose Down
DROP INDEX IF EXISTS idx_lab_results_unit_revisar;
ALTER TABLE public.lab_results
  DROP COLUMN IF EXISTS unit_conversion_status,
  DROP COLUMN IF EXISTS unit_conversion_note;
