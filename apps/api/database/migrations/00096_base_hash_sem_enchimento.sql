-- +goose Up
-- `char(64)` enche com espaço, e isso quebrava as sugestões estruturais.
--
-- Uma sugestão de `add` não tem slide alvo para congelar, então grava `base_hash` vazio. Em
-- `char(64)` o Postgres devolve 64 ESPAÇOS, que no Go não é string vazia: o guarda de "o slide
-- mudou desde que esta sugestão nasceu" disparava, procurava um slide de id vazio, não achava, e
-- recusava a sugestão com "o slide não existe mais". Nenhum slide novo jamais entrou por aceite.
--
-- `varchar(64)` guarda o que foi escrito. Os hashes reais têm exatamente 64 caracteres e não mudam.
ALTER TABLE public.patient_plan_suggestions
  ALTER COLUMN base_hash TYPE varchar(64);

UPDATE public.patient_plan_suggestions SET base_hash = '' WHERE btrim(base_hash) = '';

-- +goose Down
ALTER TABLE public.patient_plan_suggestions
  ALTER COLUMN base_hash TYPE char(64);
