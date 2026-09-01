-- +goose Up
-- O que o paciente vê tem que ser o que foi PUBLICADO, não o rascunho em andamento.
--
-- O plano guardava um `content` só. Editar um plano já publicado — corrigir um typo que fosse —
-- devolvia o registro para rascunho, e como a consulta do portal filtra `status = 'published'`, o
-- plano SUMIA da tela do paciente até alguém republicar. Pior: a instrução da skill promete o
-- contrário ("o que está no portal continua sendo a versão publicada até alguém publicar outra").
--
-- Com a coluna abaixo, publicar tira uma cópia do conteúdo e é essa cópia que o portal lê. O
-- médico edita o rascunho à vontade; o paciente continua com a versão que recebeu.

ALTER TABLE public.patient_plans
  ADD COLUMN IF NOT EXISTS published_content jsonb;

COMMENT ON COLUMN public.patient_plans.published_content IS
  'Cópia do content no momento da publicação. É o que o portal do paciente lê; o content é o rascunho vivo do médico.';

-- Planos já publicados antes desta migration: o content vigente É o que foi publicado.
UPDATE public.patient_plans
   SET published_content = content
 WHERE status = 'published' AND published_content IS NULL;

-- +goose Down
ALTER TABLE public.patient_plans DROP COLUMN IF EXISTS published_content;
