-- Gravação + transcrição da teleconsulta (Daily.co cloud recording + Deepgram nova-3).
-- Uma linha por consulta (upsert por appointment_id). Artefatos chegam via webhook
-- do Daily (POST /webhooks/daily), mapeados por daily_room_name. O MP4 NÃO é
-- armazenado aqui (referência + link sob demanda); a transcrição (WebVTT) é baixada,
-- parseada e guardada em transcript_text. Aditiva.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.telemed_recordings (
    id                          uuid PRIMARY KEY,
    appointment_id              uuid,
    patient_id                  uuid,
    daily_room_name             varchar(255) NOT NULL DEFAULT '',

    recording_id                varchar(64),
    recording_status            varchar(16) NOT NULL DEFAULT 'pending',
    recording_started_at        timestamp with time zone,
    recording_ready_at          timestamp with time zone,
    recording_duration_seconds  integer,
    recording_s3_key            varchar(500),
    recording_error             text,

    transcript_id               varchar(64),
    transcript_status           varchar(16) NOT NULL DEFAULT 'none',
    transcript_ready_at         timestamp with time zone,
    transcript_vtt_path         varchar(500),
    transcript_text             text,
    transcript_error            text,

    created_at                  timestamp with time zone NOT NULL DEFAULT now(),
    updated_at                  timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at                  timestamp with time zone
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS idx_telemed_recordings_appointment
    ON public.telemed_recordings (appointment_id)
    WHERE deleted_at IS NULL AND appointment_id IS NOT NULL;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_telemed_recordings_room
    ON public.telemed_recordings (daily_room_name)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_telemed_recordings_patient
    ON public.telemed_recordings (patient_id);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_telemed_recordings_deleted_at
    ON public.telemed_recordings (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.telemed_recordings;
-- +goose StatementEnd
