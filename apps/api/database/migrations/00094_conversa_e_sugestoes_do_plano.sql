-- +goose Up
-- A conversa que edita o rascunho, e as sugestões que ela produz.
--
-- A regra combinada: alteração de TEXTO entra direto e reversível; alteração que toque número,
-- unidade, dose ou régua vira SUGESTÃO, aceita slide a slide, com a origem do número à vista.
--
-- A régua vem de evidência, não de gosto. Em mensagens de portal redigidas por IA com erro
-- plantado, mais da metade dos clínicos não pegou todos os erros e 35% a 45% enviaram sem editar,
-- com 90% deles declarando confiar na ferramenta (npj Digital Medicine, 2025). Ou seja: não dá
-- para apostar a segurança na releitura. O que dá para verificar por código é verificado, e o que
-- não dá vira sugestão com a origem anexada, para o médico julgar uma afirmação específica em vez
-- de reler um parágrafo.

-- ---------------------------------------------------------------------------
-- As mensagens.
CREATE TABLE IF NOT EXISTS public.patient_plan_messages (
    id      uuid PRIMARY KEY,
    plan_id uuid NOT NULL REFERENCES public.patient_plans(id) ON DELETE CASCADE,
    seq     integer NOT NULL CHECK (seq >= 1),
    role    varchar(12) NOT NULL CHECK (role IN ('user','assistant')),
    body    text NOT NULL DEFAULT '',
    user_id uuid REFERENCES public.users(id) ON DELETE SET NULL,

    -- Idempotência. A chamada é síncrona e leva de dez a vinte segundos; fechar a aba depois de o
    -- modelo ter respondido não pode virar um turno duplicado quando o médico reenvia.
    client_message_id varchar(64),

    status        varchar(12) NOT NULL DEFAULT 'ok' CHECK (status IN ('ok','failed','refused')),
    error_message text,

    -- Metadados da chamada. Nunca o prompt nem a resposta: conteúdo clínico não vai para log.
    ai_model          varchar(60),
    ai_prompt_version varchar(20),
    ai_input_tokens   integer,
    ai_cache_read_tokens integer,
    ai_output_tokens  integer,
    ai_stop_reason    varchar(30),
    latency_ms        integer,

    dossier_id uuid REFERENCES public.patient_plan_dossiers(id) ON DELETE SET NULL,
    created_at timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT patient_plan_messages_plan_seq UNIQUE (plan_id, seq)
);

-- O predicado exclui a string VAZIA, e não só NULL. O Go grava `''` num campo string sem
-- ponteiro, então a mensagem do assistente (que não tem id de cliente) entraria no índice e
-- colidiria com a do turno anterior — quebrando o segundo turno de toda conversa. Apareceu no
-- primeiro teste de ponta a ponta.
CREATE UNIQUE INDEX IF NOT EXISTS uq_plan_messages_client
  ON public.patient_plan_messages (plan_id, client_message_id)
  WHERE client_message_id IS NOT NULL AND client_message_id <> '';

COMMENT ON COLUMN public.patient_plan_messages.ai_cache_read_tokens IS
  'Tokens lidos do cache de prompt. A partir do segundo turno tem que ser > 0: o dossiê congelado é byte-idêntico entre turnos, e é justamente por isso que ele é congelado. Zero aqui significa que algo volátil entrou antes do ponto de cache.';

-- ---------------------------------------------------------------------------
-- As sugestões pendentes.
--
-- Tabela própria, e não um campo dentro da revisão: a revisão é imutável e a sugestão muda de
-- estado; "aceitar a do slide 6" dentro de um JSONB seria ler-modificar-escrever sem trava de
-- linha; e "quantas sugestões numéricas foram recusadas" — a métrica que diz se a ferramenta é
-- confiável — ficaria inconsultável.
CREATE TABLE IF NOT EXISTS public.patient_plan_suggestions (
    id         uuid PRIMARY KEY,
    plan_id    uuid NOT NULL REFERENCES public.patient_plans(id) ON DELETE CASCADE,
    message_id uuid NOT NULL REFERENCES public.patient_plan_messages(id) ON DELETE CASCADE,

    op             varchar(12) NOT NULL CHECK (op IN ('add','edit','remove','reorder')),
    slide_id       varchar(64),
    after_slide_id varchar(64),
    field_path     text,
    old_value      jsonb,
    new_value      jsonb,

    -- Hash do slide alvo quando a sugestão nasceu. Se o médico editou aquele slide à mão depois,
    -- aceitar sobrescreveria o trabalho dele em silêncio; o hash faz a sugestão virar obsoleta.
    base_hash char(64),

    class     varchar(12) NOT NULL CHECK (class IN ('numeric','structural')),
    rationale text NOT NULL DEFAULT '',
    -- De onde cada número desta sugestão veio no dossiê congelado. É o que o médico lê ao lado do
    -- botão para decidir; sem isto ele estaria aceitando prosa.
    provenance jsonb NOT NULL DEFAULT '[]'::jsonb,

    status         varchar(12) NOT NULL DEFAULT 'pending'
                     CHECK (status IN ('pending','accepted','rejected','stale','superseded')),
    resolved_by_id uuid REFERENCES public.users(id) ON DELETE SET NULL,
    resolved_at    timestamptz,
    revision_id    uuid REFERENCES public.patient_plan_revisions(id) ON DELETE SET NULL,
    created_at     timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_plan_suggestions_pendentes
  ON public.patient_plan_suggestions (plan_id, slide_id) WHERE status = 'pending';

-- A revisão já referenciava a mensagem por id; agora a FK existe.
--
-- Limpa referência órfã antes de criar: se esta migration já foi revertida alguma vez, sobraram
-- `message_id` apontando para mensagens que o `down` apagou, e o ALTER falharia. Apareceu ao
-- testar a reversão — sem isto, um rollback deixaria a tabela impossível de re-migrar.
UPDATE public.patient_plan_revisions SET message_id = NULL
 WHERE message_id IS NOT NULL
   AND NOT EXISTS (SELECT 1 FROM public.patient_plan_messages m WHERE m.id = message_id);

-- +goose StatementBegin
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'patient_plan_revisions_message_fk') THEN
    ALTER TABLE public.patient_plan_revisions
      ADD CONSTRAINT patient_plan_revisions_message_fk
      FOREIGN KEY (message_id) REFERENCES public.patient_plan_messages(id) ON DELETE SET NULL;
  END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.patient_plan_revisions DROP CONSTRAINT IF EXISTS patient_plan_revisions_message_fk;
-- Solta as referências antes de apagar as mensagens, senão a re-aplicação encontra órfãos.
UPDATE public.patient_plan_revisions SET message_id = NULL WHERE message_id IS NOT NULL;
DROP TABLE IF EXISTS public.patient_plan_suggestions;
DROP TABLE IF EXISTS public.patient_plan_messages;
