-- Reconciliação de drift dev↔prod antes do cutover pro goose (2026-06-01).
-- O AutoMigrate de prod logava "✅ concluído" mas não aplicava ALTERs; prod ficou
-- sem ~26 colunas que os models exigem (a tabela appointment_recalls é coberta pela
-- 00003). Migration ADITIVA e IDEMPOTENTE: no dev (que já tem tudo) é no-op; em prod
-- aplica o delta. Definições copiadas do schema do dev (fiel aos models).
-- Tabelas afetadas estavam VAZIAS em prod (verificado), então NOT NULL é seguro.

-- +goose Up
-- +goose StatementBegin

-- users — colunas OAuth + check + índices
ALTER TABLE public.users ADD COLUMN IF NOT EXISTS oauth_provider character varying(20);
ALTER TABLE public.users ADD COLUMN IF NOT EXISTS oauth_provider_id character varying(255);
ALTER TABLE public.users ADD COLUMN IF NOT EXISTS oauth_picture_url text;
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'chk_oauth_provider') THEN
        ALTER TABLE public.users ADD CONSTRAINT chk_oauth_provider CHECK (((oauth_provider IS NULL) OR ((oauth_provider)::text = ANY (ARRAY[('google'::character varying)::text, ('apple'::character varying)::text]))));
    END IF;
END $$;
CREATE INDEX IF NOT EXISTS idx_users_oauth_provider ON public.users USING btree (oauth_provider) WHERE (oauth_provider IS NOT NULL);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_oauth_unique ON public.users USING btree (oauth_provider, oauth_provider_id) WHERE ((oauth_provider IS NOT NULL) AND (oauth_provider_id IS NOT NULL));

-- prescriptions — colunas do medicamento (tabela vazia em prod) + check categoria + índice
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS medication_definition_id uuid;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS medication_name character varying(200) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS active_ingredient character varying(200) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS category character varying(20) DEFAULT 'simple'::character varying NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS concentration character varying(100) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS dosage character varying(100) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS frequency character varying(100) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS route character varying(50) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS duration integer NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS quantity integer NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS quantity_in_words character varying(200) NOT NULL;
ALTER TABLE public.prescriptions ADD COLUMN IF NOT EXISTS instructions text;
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'chk_prescriptions_category') THEN
        ALTER TABLE public.prescriptions ADD CONSTRAINT chk_prescriptions_category CHECK (((category)::text = ANY (ARRAY[('simple'::character varying)::text, ('c1'::character varying)::text, ('c5'::character varying)::text, ('antibiotic'::character varying)::text, ('glp1'::character varying)::text])));
    END IF;
END $$;
CREATE INDEX IF NOT EXISTS idx_prescriptions_medication_definition_id ON public.prescriptions USING btree (medication_definition_id);

-- article_score_items — linking RAG + feedback (tabela vazia) + checks + índices
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS confidence_score double precision DEFAULT 1.0;
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS auto_linked boolean DEFAULT false;
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS linked_at timestamp with time zone DEFAULT now();
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS linked_by uuid;
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS user_feedback character varying(20);
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS feedback_at timestamp with time zone;
ALTER TABLE public.article_score_items ADD COLUMN IF NOT EXISTS feedback_by uuid;
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'article_score_items_user_feedback_check') THEN
        ALTER TABLE public.article_score_items ADD CONSTRAINT article_score_items_user_feedback_check CHECK (((user_feedback)::text = ANY (ARRAY[('approved'::character varying)::text, ('rejected'::character varying)::text, ('irrelevant'::character varying)::text])));
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'check_article_score_items_confidence') THEN
        ALTER TABLE public.article_score_items ADD CONSTRAINT check_article_score_items_confidence CHECK (((confidence_score >= (0.0)::double precision) AND (confidence_score <= (1.0)::double precision)));
    END IF;
END $$;
CREATE INDEX IF NOT EXISTS idx_article_score_items_auto_linked ON public.article_score_items USING btree (auto_linked);
CREATE INDEX IF NOT EXISTS idx_article_score_items_auto_linked_feedback ON public.article_score_items USING btree (auto_linked, user_feedback) WHERE (auto_linked = true);
CREATE INDEX IF NOT EXISTS idx_article_score_items_feedback ON public.article_score_items USING btree (user_feedback) WHERE (user_feedback IS NOT NULL);

-- embedding_queue — priority/max_retries/metadata + índice
ALTER TABLE public.embedding_queue ADD COLUMN IF NOT EXISTS priority integer DEFAULT 0;
ALTER TABLE public.embedding_queue ADD COLUMN IF NOT EXISTS max_retries integer DEFAULT 3;
ALTER TABLE public.embedding_queue ADD COLUMN IF NOT EXISTS metadata jsonb DEFAULT '{}'::jsonb;
CREATE INDEX IF NOT EXISTS idx_embedding_queue_priority ON public.embedding_queue USING btree (priority DESC, created_at) WHERE ((status)::text = 'pending'::text);

-- lab_request_template_tests.created_at (tabela vazia em prod)
ALTER TABLE public.lab_request_template_tests ADD COLUMN IF NOT EXISTS created_at timestamp with time zone DEFAULT now() NOT NULL;

-- patients.birth_date — model exige NOT NULL; prod tem 1 linha sem null (verificado)
ALTER TABLE public.patients ALTER COLUMN birth_date SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Forward-only (reconciliação aditiva; reverter colunas perderia dados futuros).
SELECT 1;
-- +goose StatementEnd
