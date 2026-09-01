-- +goose Up
-- Revisões do plano de devolutiva, identidade do slide e concorrência otimista.
--
-- `patient_plan_service.Update` faz `plan.Content = req.Content` e salva. `patient_plans.version`
-- conta PUBLICAÇÕES, não edições. Na sessão que motivou isto, o deck de uma paciente passou por
-- quatro versões em algumas horas (10 slides, 18, 19, 20) e as três primeiras não existem mais:
-- cada PUT apagou a anterior. Enquanto só o médico escrevia, "desfazer" era não salvar. Com um
-- assistente escrevendo no MESMO campo, desfazer virou requisito, e requisito que não pode
-- depender do estado do navegador.
--
-- De quebra conserta um defeito que já existe hoje: `Publish` sobrescreve `published_content`.
-- Republicar destrói os bytes exatos da v1 para sempre — o PDF continua no portal, mas o banco
-- deixa de saber o que estava escrito nele. Com `is_publication`, "o que o paciente leu em
-- janeiro" volta a ser uma consulta.

-- ---------------------------------------------------------------------------
-- 1 · Identidade do slide.
--
-- Até aqui um slide só era endereçável pelo índice no array. Índice muda quando alguém reordena,
-- então uma sugestão criada sobre "o slide 6" e aceita cinco minutos depois escreveria no slide
-- errado, sem erro e sem log, num documento que o paciente lê. O id é o alvo estável de toda
-- operação, sugestão e diff.
--
-- O backfill preenche `content` e `published_content`. `published_content` pode ser NULL ou não
-- ser array em linhas antigas, daí o CASE.
-- +goose StatementBegin
DO $$
DECLARE r record;
BEGIN
  FOR r IN SELECT id, content, published_content FROM public.patient_plans LOOP
    UPDATE public.patient_plans SET
      content = COALESCE((
        SELECT jsonb_agg(
                 CASE WHEN e ? 'id' THEN e
                      ELSE e || jsonb_build_object('id', gen_random_uuid()::text) END
                 ORDER BY ord)
        FROM jsonb_array_elements(r.content) WITH ORDINALITY AS t(e, ord)
      ), '[]'::jsonb),
      published_content = CASE
        WHEN r.published_content IS NULL
          OR jsonb_typeof(r.published_content) <> 'array' THEN r.published_content
        ELSE (
          SELECT jsonb_agg(
                   CASE WHEN e ? 'id' THEN e
                        ELSE e || jsonb_build_object('id', gen_random_uuid()::text) END
                   ORDER BY ord)
          FROM jsonb_array_elements(r.published_content) WITH ORDINALITY AS t(e, ord)
        )
      END
    WHERE id = r.id;
  END LOOP;
END $$;
-- +goose StatementEnd

-- ---------------------------------------------------------------------------
-- 2 · As revisões.
--
-- Guarda o estado RESULTANTE, não o delta. Restaurar passa a ser copiar uma linha, e uma cadeia
-- de patches corrompida no meio mataria todo o histórico a partir dali. Mesmo formato do
-- precedente `integrated_plan_revisions` (patient_continuum.go), com as colunas que aquele não
-- precisava: quem escreveu, por quê, com qual modelo, contra qual dossiê.
CREATE TABLE IF NOT EXISTS public.patient_plan_revisions (
    id             uuid PRIMARY KEY,
    plan_id        uuid NOT NULL REFERENCES public.patient_plans(id) ON DELETE CASCADE,

    -- Conta EDIÇÕES, 1..N, monotônico, atravessando publicações.
    seq            integer NOT NULL CHECK (seq >= 1),
    -- A publicação a que esta edição pertence (o valor de `patient_plans.version` no momento).
    -- Separado de `seq` de propósito: confundir "edição 47" com "versão 47" é o erro que a
    -- coluna existe para evitar. Na tela: "v2 no portal · rascunho, edição 47".
    plan_version   integer NOT NULL,

    title          varchar(300) NOT NULL,
    content        jsonb NOT NULL,
    -- sha256 do conteúdo canônico. Save que não mudou nada não vira linha.
    content_hash   char(64) NOT NULL,

    author_kind    varchar(12) NOT NULL CHECK (author_kind IN ('human','assistant','system')),
    -- SEMPRE o clínico logado, inclusive quando author_kind = 'assistant'. A ferramenta escreve
    -- sob a conta de alguém, e alguém responde pelo que o paciente leu.
    created_by_id  uuid NOT NULL REFERENCES public.users(id) ON DELETE RESTRICT,
    reason         varchar(24) NOT NULL
                     CHECK (reason IN ('edit','ai_apply','suggestion_accept','restore','publish')),

    -- As operações aplicadas, para responder "o que mudou" sem diffar dois JSONB de 60 KB na tela.
    ops            jsonb,
    message_id     uuid,
    dossier_id     uuid,
    ai_model          varchar(60),
    ai_prompt_version varchar(20),

    is_publication boolean NOT NULL DEFAULT false,
    created_at     timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT patient_plan_revisions_plan_seq UNIQUE (plan_id, seq)
);

-- +goose StatementBegin
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'patient_plan_revisions_content_is_array') THEN
    ALTER TABLE public.patient_plan_revisions
      ADD CONSTRAINT patient_plan_revisions_content_is_array
      CHECK (jsonb_typeof(content) = 'array');
  END IF;
END $$;
-- +goose StatementEnd

CREATE INDEX IF NOT EXISTS idx_plan_revisions_plan ON public.patient_plan_revisions (plan_id, seq DESC);
CREATE INDEX IF NOT EXISTS idx_plan_revisions_pub  ON public.patient_plan_revisions (plan_id) WHERE is_publication;

COMMENT ON TABLE public.patient_plan_revisions IS
  'Histórico de edições do rascunho. Guarda o conteúdo resultante, não o delta. Uma revisão por gravação, com coalescência de 2 min para o mesmo autor humano (senão o autosave geraria ~1.800 linhas de 60 KB por sessão de autoria).';
COMMENT ON COLUMN public.patient_plan_revisions.seq IS
  'Conta EDIÇÕES. patient_plans.version conta PUBLICAÇÕES. Nunca somar os dois.';
COMMENT ON COLUMN public.patient_plan_revisions.created_by_id IS
  'O clínico logado, sempre — inclusive quando author_kind = assistant.';

-- ---------------------------------------------------------------------------
-- 3 · Concorrência otimista e ponteiros.
--
-- `revision_seq` é o token: o PUT manda o que acha que é a revisão corrente e leva 409 se não
-- bater. Sem isso, com um segundo escritor no mesmo `content`, um autosave em voo (carregado
-- antes) apaga a edição que acabou de entrar, sem ninguém ver.
ALTER TABLE public.patient_plans
  ADD COLUMN IF NOT EXISTS revision_seq          integer NOT NULL DEFAULT 0,
  ADD COLUMN IF NOT EXISTS published_revision_id uuid REFERENCES public.patient_plan_revisions(id) ON DELETE SET NULL;

COMMENT ON COLUMN public.patient_plans.revision_seq IS
  'Número da última revisão. Token de concorrência otimista: PUT e turno do assistente mandam expectedRevision e levam 409 quando estão velhos.';
COMMENT ON COLUMN public.patient_plans.version IS
  'Conta PUBLICAÇÕES, não edições. Edições são patient_plan_revisions.seq.';

-- ---------------------------------------------------------------------------
-- 4 · Semente do histórico.
--
-- Todo plano que já existe ganha a revisão 1 com o conteúdo atual, para o histórico não começar
-- vazio e o `revision_seq` não nascer mentindo. Autor: o autor do plano.
-- +goose StatementBegin
DO $$
DECLARE p record;
BEGIN
  FOR p IN SELECT id, title, content, version, author_user_id, created_at
             FROM public.patient_plans WHERE deleted_at IS NULL LOOP
    IF NOT EXISTS (SELECT 1 FROM public.patient_plan_revisions WHERE plan_id = p.id) THEN
      INSERT INTO public.patient_plan_revisions
        (id, plan_id, seq, plan_version, title, content, content_hash,
         author_kind, created_by_id, reason, is_publication, created_at)
      VALUES
        (gen_random_uuid(), p.id, 1, p.version, p.title, p.content,
         encode(sha256(p.content::text::bytea), 'hex'),
         'human', p.author_user_id, 'edit', false, p.created_at);
      UPDATE public.patient_plans SET revision_seq = 1 WHERE id = p.id;
    END IF;
  END LOOP;
END $$;
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.patient_plans
  DROP COLUMN IF EXISTS published_revision_id,
  DROP COLUMN IF EXISTS revision_seq;
DROP TABLE IF EXISTS public.patient_plan_revisions;
-- Os ids dos slides ficam: são inertes para o render, e removê-los quebraria qualquer referência
-- que já tenha sido criada contra eles.
