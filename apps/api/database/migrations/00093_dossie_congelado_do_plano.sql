-- +goose Up
-- O dossiê que o plano usou, congelado no momento em que foi montado.
--
-- Hoje `GET /patients/:id/plan-dossier` recalcula tudo a cada chamada: ~28 queries, dominadas por
-- `ScoreSnapshotRepository.GetLatestByPatientID` com ~15 Preloads que servem outra tela. Isso era
-- tolerável quando o dossiê era lido uma vez por abertura de tela. Deixa de ser por dois motivos.
--
-- O primeiro é que os números MUDAM debaixo do rascunho. No caso real que motivou isto, o escore
-- da paciente mudou três vezes enquanto o deck estava sendo escrito (correção de menopausa,
-- correção de unidade, complemento de anamnese). Um slide escrito contra ferritina 48 e revisado
-- num mundo onde ela é 96 é uma armadilha silenciosa.
--
-- O segundo é auditoria: sem o congelamento não há como responder "o que a máquina sabia quando
-- esta frase foi escrita".
--
-- É tabela e não coluna em `patient_plans` de propósito. `PatientPlanService.load` faz `SELECT *`
-- e é chamado por preview, overflow, publish e get; um jsonb de centenas de KB naquela linha seria
-- detoastado em toda leitura do plano, inclusive no caminho que hoje custa 5 ms.

CREATE TABLE IF NOT EXISTS public.patient_plan_dossiers (
    id       uuid PRIMARY KEY,
    plan_id  uuid NOT NULL REFERENCES public.patient_plans(id) ON DELETE CASCADE,
    -- 1..N: um plano tem mais de um congelamento quando o médico pede para atualizar. O anterior
    -- fica, senão não dá para explicar por que o slide 6 diz o que diz.
    seq      integer NOT NULL CHECK (seq >= 1),

    -- `dto.PlanDossierResponse` inteiro, como estava no momento.
    payload  jsonb NOT NULL,

    -- Marcas d'água do prontuário no momento do congelamento. Servem para detectar
    -- envelhecimento com UMA query em vez de remontar o dossiê inteiro só para comparar.
    source_snapshot_id uuid REFERENCES public.patient_score_snapshots(id) ON DELETE SET NULL,
    latest_lab_at      timestamptz,
    latest_vitals_at   timestamptz,
    latest_snapshot_at timestamptz,

    built_at    timestamptz NOT NULL DEFAULT now(),
    built_by_id uuid REFERENCES public.users(id) ON DELETE SET NULL,

    CONSTRAINT patient_plan_dossiers_plan_seq UNIQUE (plan_id, seq)
);

CREATE INDEX IF NOT EXISTS idx_plan_dossiers_plan ON public.patient_plan_dossiers (plan_id, seq DESC);

COMMENT ON TABLE public.patient_plan_dossiers IS
  'O prontuário compilado que o plano usou, congelado. Nunca é atualizado sozinho: refrescar troca números debaixo do autor e invalida a base contra a qual o conteúdo foi conferido.';
COMMENT ON COLUMN public.patient_plan_dossiers.latest_lab_at IS
  'Marca d''água do último resultado de exame no congelamento. Comparar com o valor atual detecta envelhecimento sem remontar o dossiê.';

ALTER TABLE public.patient_plans
  ADD COLUMN IF NOT EXISTS current_dossier_id uuid
    REFERENCES public.patient_plan_dossiers(id) ON DELETE SET NULL;

COMMENT ON COLUMN public.patient_plans.current_dossier_id IS
  'O congelamento em vigor. A tela de autoria lê deste, nunca do dossiê vivo do paciente.';

-- As revisões já referenciavam o dossiê por id; agora a FK existe.
-- +goose StatementBegin
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'patient_plan_revisions_dossier_fk') THEN
    ALTER TABLE public.patient_plan_revisions
      ADD CONSTRAINT patient_plan_revisions_dossier_fk
      FOREIGN KEY (dossier_id) REFERENCES public.patient_plan_dossiers(id) ON DELETE SET NULL;
  END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
ALTER TABLE public.patient_plan_revisions DROP CONSTRAINT IF EXISTS patient_plan_revisions_dossier_fk;
ALTER TABLE public.patient_plans DROP COLUMN IF EXISTS current_dossier_id;
DROP TABLE IF EXISTS public.patient_plan_dossiers;
