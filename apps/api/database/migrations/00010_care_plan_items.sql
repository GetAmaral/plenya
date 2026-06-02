-- Plano de cuidado AGIR (P3 frente 3): recomendações ancoradas a um pilar (letra A/G/I/R)
-- e, opcionalmente, a um biomarcador/score item. Multidisciplinar. Aditivo.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.care_plan_items (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    letter_code character varying(1) NOT NULL,
    recommendation text NOT NULL,
    rationale text,
    target character varying(300),
    lab_test_code character varying(20),
    score_item_id uuid,
    source_snapshot_id uuid,
    priority character varying(10) NOT NULL DEFAULT 'medium',
    status character varying(10) NOT NULL DEFAULT 'active',
    author_user_id uuid NOT NULL,
    author_role character varying(30),
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT care_plan_items_pkey PRIMARY KEY (id),
    CONSTRAINT chk_care_plan_items_letter CHECK (letter_code IN ('A','G','I','R')),
    CONSTRAINT chk_care_plan_items_priority CHECK (priority IN ('high','medium','low')),
    CONSTRAINT chk_care_plan_items_status CHECK (status IN ('active','achieved','suspended'))
);
CREATE INDEX IF NOT EXISTS idx_care_plan_items_patient_id ON public.care_plan_items (patient_id);
CREATE INDEX IF NOT EXISTS idx_care_plan_items_active ON public.care_plan_items (patient_id) WHERE status = 'active' AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_care_plan_items_deleted_at ON public.care_plan_items (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.care_plan_items;
-- +goose StatementEnd
