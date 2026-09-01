-- +goose Up
-- Plano de paciente (a devolutiva de resultados) como entidade do prontuário.
--
-- Até aqui o plano vivia FORA do EMR: um `build.py` por paciente, copiado do paciente anterior,
-- gerando HTML+CSS renderizado por Playwright. Três pacientes depois o custo apareceu — os scripts
-- de render tinham o nome de saída hardcoded e gravaram por cima do deck de outro paciente duas
-- vezes, e nada disso ficava registrado nem chegava ao portal.
--
-- O CONTEÚDO fica em JSONB e não em tabela filha de propósito: os slides são heterogêneos (capa,
-- resumo, réguas, dois-caminhos, plano, sequência, para levar, fecho), a ORDEM é o que importa, e
-- ninguém consulta um slide isolado. Mesmo raciocínio das escalas pergunta-a-pergunta (mig 00041).

CREATE TABLE IF NOT EXISTS public.patient_plans (
    id                  uuid PRIMARY KEY,
    patient_id          uuid NOT NULL REFERENCES public.patients(id) ON DELETE CASCADE,

    title               varchar(300) NOT NULL,
    status              varchar(20)  NOT NULL DEFAULT 'draft'
                          CHECK (status IN ('draft', 'published')),
    -- Versão publicada. Um plano é reeditado a cada consulta; cada publicação vira um documento
    -- novo no portal, e a versão é o que dá idempotência à materialização (source_ref).
    version             integer NOT NULL DEFAULT 1 CHECK (version >= 1),

    -- Os slides. Contrato em pdfdoc.DeckSlide.
    content             jsonb NOT NULL DEFAULT '[]'::jsonb,

    -- Escore que motivou o plano, para rastrear o "porquê" — igual a care_plan_items.
    source_snapshot_id  uuid REFERENCES public.patient_score_snapshots(id) ON DELETE SET NULL,
    author_user_id      uuid NOT NULL REFERENCES public.users(id) ON DELETE RESTRICT,

    published_at        timestamptz,
    -- PDF 16:9 e A4 publicados no portal (patient_documents).
    document_16x9_id    uuid REFERENCES public.patient_documents(id) ON DELETE SET NULL,
    document_a4_id      uuid REFERENCES public.patient_documents(id) ON DELETE SET NULL,

    created_at          timestamptz NOT NULL DEFAULT now(),
    updated_at          timestamptz NOT NULL DEFAULT now(),
    deleted_at          timestamptz
);

-- O conteúdo é uma LISTA de slides. Guardar um objeto aqui já quebrou o catálogo de templates do
-- WhatsApp uma vez (mig 00060): o GORM gravou slice de 1 elemento como objeto jsonb e o Find
-- abortava. O CHECK impede a mesma classe de erro.
-- Guardado como o CREATE acima: se a tabela já existir, ADD CONSTRAINT cru falharia com
-- duplicate_object em vez de ser no-op.
-- +goose StatementBegin
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'patient_plans_content_is_array') THEN
    ALTER TABLE public.patient_plans
      ADD CONSTRAINT patient_plans_content_is_array
      CHECK (jsonb_typeof(content) = 'array');
  END IF;
END $$;
-- +goose StatementEnd

CREATE INDEX IF NOT EXISTS idx_patient_plans_patient   ON public.patient_plans (patient_id);
CREATE INDEX IF NOT EXISTS idx_patient_plans_status    ON public.patient_plans (status);
CREATE INDEX IF NOT EXISTS idx_patient_plans_deleted   ON public.patient_plans (deleted_at);

COMMENT ON TABLE  public.patient_plans IS
  'Devolutiva de resultados do paciente. Uma fonte de conteúdo, três saídas: tela do portal, PDF 16:9 e PDF A4 paisagem.';
COMMENT ON COLUMN public.patient_plans.content IS
  'Lista ordenada de slides (pdfdoc.DeckSlide). Heterogênea de propósito; nunca consultada slide a slide.';
COMMENT ON COLUMN public.patient_plans.version IS
  'Publicação. Compõe o source_ref "patient_plan:<id>:v<n>" que deduplica o documento no portal.';

-- +goose Down
DROP TABLE IF EXISTS public.patient_plans;
