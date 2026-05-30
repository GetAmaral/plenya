-- Baseline do schema do EMR Plenya (pg_dump --schema-only do dev, 2026-05-30).
-- Schema vivo construído historicamente por GORM AutoMigrate + blocos DO.
-- A partir daqui, mudanças de schema entram como NOVAS migrations goose (00002_*, ...).
-- Bancos JÁ existentes (dev/prod) recebem `migrate stamp` (marca aplicada, não reexecuta).
-- Só roda inteira em bancos vazios (CI, setup novo). search_path-reset removido de propósito.

-- +goose Up
-- +goose StatementBegin
--
-- PostgreSQL database dump
--


-- Dumped from database version 17.7 (Debian 17.7-3.pgdg12+1)
-- Dumped by pg_dump version 17.7 (Debian 17.7-3.pgdg12+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: btree_gist; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gist WITH SCHEMA public;


--
-- Name: EXTENSION btree_gist; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION btree_gist IS 'support for indexing common datatypes in GiST';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: unaccent; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS unaccent WITH SCHEMA public;


--
-- Name: EXTENSION unaccent; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION unaccent IS 'text search dictionary that removes accents';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: vector; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS vector WITH SCHEMA public;


--
-- Name: EXTENSION vector; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION vector IS 'vector data type and ivfflat and hnsw access methods';


--
-- Name: complete_auto_link_run(uuid, character varying, text); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.complete_auto_link_run(p_run_id uuid, p_status character varying, p_error_message text DEFAULT NULL::text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE auto_link_processing_state
    SET
        status = p_status,
        completed_at = NOW(),
        error_message = p_error_message
    WHERE run_id = p_run_id;
END;
$$;


--
-- Name: FUNCTION complete_auto_link_run(p_run_id uuid, p_status character varying, p_error_message text); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.complete_auto_link_run(p_run_id uuid, p_status character varying, p_error_message text) IS 'Marks a run as completed or failed';


--
-- Name: create_auto_link_run(integer, integer); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.create_auto_link_run(p_total_items integer, p_batch_size integer DEFAULT 50) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
    v_run_id UUID;
BEGIN
    v_run_id := uuid_generate_v7();

    INSERT INTO auto_link_processing_state (run_id, total_items, batch_size, status)
    VALUES (v_run_id, p_total_items, p_batch_size, 'running');

    RETURN v_run_id;
END;
$$;


--
-- Name: FUNCTION create_auto_link_run(p_total_items integer, p_batch_size integer); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.create_auto_link_run(p_total_items integer, p_batch_size integer) IS 'Creates a new auto-link processing run with initial state';


--
-- Name: immutable_unaccent(text); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.immutable_unaccent(text) RETURNS text
    LANGUAGE sql IMMUTABLE STRICT PARALLEL SAFE
    AS $_$
  SELECT public.unaccent('public.unaccent', $1)
$_$;


--
-- Name: invalidate_embedding(character varying, uuid, text, integer); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.invalidate_embedding(p_entity_type character varying, p_entity_id uuid, p_reason text DEFAULT NULL::text, p_priority integer DEFAULT 0) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Mark embedding as stale
    IF p_entity_type = 'score_item' THEN
        UPDATE score_item_embeddings
        SET is_stale = true
        WHERE score_item_id = p_entity_id;
    ELSIF p_entity_type = 'article' THEN
        UPDATE article_embeddings
        SET is_stale = true
        WHERE article_id = p_entity_id;
    END IF;

    -- Queue for regeneration (upsert)
    INSERT INTO embedding_queue (entity_type, entity_id, status, priority, metadata)
    VALUES (
        p_entity_type,
        p_entity_id,
        'pending',
        p_priority,
        jsonb_build_object('reason', p_reason, 'invalidated_at', NOW())
    )
    ON CONFLICT (entity_type, entity_id)
    DO UPDATE SET
        status = 'pending',
        priority = GREATEST(embedding_queue.priority, p_priority),
        created_at = NOW(),
        retry_count = 0,
        metadata = embedding_queue.metadata || jsonb_build_object('reason', p_reason, 'invalidated_at', NOW());

    -- Log to audit trail
    INSERT INTO embedding_audit_log (entity_type, entity_id, action, reason, triggered_by)
    VALUES (p_entity_type, p_entity_id, 'invalidated', p_reason, 'invalidate_embedding function');
END;
$$;


--
-- Name: FUNCTION invalidate_embedding(p_entity_type character varying, p_entity_id uuid, p_reason text, p_priority integer); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.invalidate_embedding(p_entity_type character varying, p_entity_id uuid, p_reason text, p_priority integer) IS 'Helper function to mark embedding as stale and queue for regeneration';


--
-- Name: invalidate_preparation(uuid, text); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.invalidate_preparation(p_score_item_id uuid, p_reason text DEFAULT NULL::text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Mark as stale or delete (currently deleting for simplicity)
    DELETE FROM score_item_enrichment_preparation
    WHERE score_item_id = p_score_item_id;

    -- Could alternatively mark as stale:
    -- UPDATE score_item_enrichment_preparation
    -- SET status = 'stale'
    -- WHERE score_item_id = p_score_item_id;
END;
$$;


--
-- Name: FUNCTION invalidate_preparation(p_score_item_id uuid, p_reason text); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.invalidate_preparation(p_score_item_id uuid, p_reason text) IS 'Invalidate preparation when ScoreItem changes (deletes for re-preparation)';


--
-- Name: migrate_fk_column(text, text, text); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.migrate_fk_column(p_table_name text, p_column_name text, p_referenced_table text) RETURNS void
    LANGUAGE plpgsql
    AS $$
DECLARE
    v_sql TEXT;
    v_updated_count BIGINT;
BEGIN
    v_sql := format(
        'UPDATE %I SET %I = m.new_id FROM uuid_migration_map m ' ||
        'WHERE %I.%I = m.old_id AND m.table_name = %L',
        p_table_name, p_column_name,
        p_table_name, p_column_name, p_referenced_table
    );

    EXECUTE v_sql;
    GET DIAGNOSTICS v_updated_count = ROW_COUNT;

    RAISE NOTICE 'Updated %.%: % rows', p_table_name, p_column_name, v_updated_count;
END;
$$;


--
-- Name: normalize_alt_names(jsonb); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.normalize_alt_names(alt_names_input jsonb) RETURNS jsonb
    LANGUAGE plpgsql IMMUTABLE
    AS $$
DECLARE
  normalized_array jsonb := '[]'::jsonb;
  item text;
BEGIN
  -- Se for NULL ou array vazio, retornar array vazio
  IF alt_names_input IS NULL OR jsonb_array_length(alt_names_input) = 0 THEN
    RETURN '[]'::jsonb;
  END IF;

  -- Iterar sobre cada item do array
  FOR item IN SELECT jsonb_array_elements_text(alt_names_input)
  LOOP
    -- Aplicar unaccent + lower e adicionar ao array normalizado
    normalized_array := normalized_array || to_jsonb(unaccent(lower(trim(item))));
  END LOOP;

  RETURN normalized_array;
END;
$$;


--
-- Name: save_batch_checkpoint(uuid, integer, integer, integer, integer, integer, integer, integer, double precision, integer, uuid); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.save_batch_checkpoint(p_run_id uuid, p_batch_number integer, p_start_index integer, p_end_index integer, p_items_processed integer, p_links_created integer, p_items_skipped integer, p_items_failed integer, p_avg_confidence double precision, p_processing_time_ms integer, p_last_item_id uuid) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Insert batch checkpoint
    INSERT INTO auto_link_batch_checkpoints (
        run_id, batch_number, batch_start_index, batch_end_index,
        items_processed, links_created, items_skipped, items_failed,
        avg_confidence, processing_time_ms
    ) VALUES (
        p_run_id, p_batch_number, p_start_index, p_end_index,
        p_items_processed, p_links_created, p_items_skipped, p_items_failed,
        p_avg_confidence, p_processing_time_ms
    );

    -- Update run state
    UPDATE auto_link_processing_state
    SET
        last_processed_item_id = p_last_item_id,
        last_processed_index = p_end_index,
        total_processed = total_processed + p_items_processed,
        total_linked = total_linked + p_links_created,
        total_skipped = total_skipped + p_items_skipped,
        last_checkpoint_at = NOW()
    WHERE run_id = p_run_id;
END;
$$;


--
-- Name: FUNCTION save_batch_checkpoint(p_run_id uuid, p_batch_number integer, p_start_index integer, p_end_index integer, p_items_processed integer, p_links_created integer, p_items_skipped integer, p_items_failed integer, p_avg_confidence double precision, p_processing_time_ms integer, p_last_item_id uuid); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.save_batch_checkpoint(p_run_id uuid, p_batch_number integer, p_start_index integer, p_end_index integer, p_items_processed integer, p_links_created integer, p_items_skipped integer, p_items_failed integer, p_avg_confidence double precision, p_processing_time_ms integer, p_last_item_id uuid) IS 'Saves checkpoint after each batch completes';


--
-- Name: submit_article_link_feedback(uuid, uuid, character varying, uuid); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.submit_article_link_feedback(p_score_item_id uuid, p_article_id uuid, p_feedback character varying, p_user_id uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
DECLARE
    v_exists BOOLEAN;
BEGIN
    -- Validate feedback value
    IF p_feedback NOT IN ('approved', 'rejected', 'irrelevant') THEN
        RAISE EXCEPTION 'Invalid feedback value: %. Must be: approved, rejected, or irrelevant', p_feedback;
    END IF;

    -- Check if link exists
    SELECT EXISTS(
        SELECT 1 FROM article_score_items
        WHERE score_item_id = p_score_item_id AND article_id = p_article_id
    ) INTO v_exists;

    IF NOT v_exists THEN
        RAISE EXCEPTION 'Article-ScoreItem link not found';
    END IF;

    -- Update feedback
    UPDATE article_score_items
    SET
        user_feedback = p_feedback,
        feedback_at = NOW(),
        feedback_by = p_user_id
    WHERE score_item_id = p_score_item_id AND article_id = p_article_id;

    RETURN TRUE;
END;
$$;


--
-- Name: FUNCTION submit_article_link_feedback(p_score_item_id uuid, p_article_id uuid, p_feedback character varying, p_user_id uuid); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.submit_article_link_feedback(p_score_item_id uuid, p_article_id uuid, p_feedback character varying, p_user_id uuid) IS 'Submit user feedback for an article-scoreitem link';


--
-- Name: update_enrichment_prep_timestamp(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_enrichment_prep_timestamp() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;


--
-- Name: update_lab_result_values_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_lab_result_values_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


--
-- Name: update_lab_test_definitions_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_lab_test_definitions_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;


--
-- Name: uuid_generate_v7(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.uuid_generate_v7() RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
    unix_ts_ms BIGINT;
    uuid_bytes BYTEA;
BEGIN
    -- Get current Unix timestamp in milliseconds (48 bits)
    unix_ts_ms := FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000)::BIGINT;

    -- Construct UUID v7: timestamp (6 bytes) + random (10 bytes)
    uuid_bytes :=
        substring(int8send(unix_ts_ms) FROM 3 FOR 6) ||
        gen_random_bytes(10);

    -- Set version bits: 0111 (v7) at byte 6
    uuid_bytes := set_byte(uuid_bytes, 6, (get_byte(uuid_bytes, 6) & 15) | 112);

    -- Set variant bits: 10xx (RFC 4122) at byte 8
    uuid_bytes := set_byte(uuid_bytes, 8, (get_byte(uuid_bytes, 8) & 63) | 128);

    RETURN encode(uuid_bytes, 'hex')::uuid;
END;
$$;


--
-- Name: FUNCTION uuid_generate_v7(); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.uuid_generate_v7() IS 'Generates time-ordered UUID v7 (RFC 9562). First 48 bits = Unix timestamp ms. Used as DEFAULT in all Plenya tables.';


--
-- Name: uuid_v7_from_timestamp(timestamp without time zone, uuid); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.uuid_v7_from_timestamp(ts timestamp without time zone, old_uuid uuid) RETURNS uuid
    LANGUAGE plpgsql IMMUTABLE
    AS $$
DECLARE
    unix_ts_ms BIGINT;
    uuid_bytes BYTEA;
    random_part BYTEA;
BEGIN
    -- Extract timestamp in milliseconds
    unix_ts_ms := FLOOR(EXTRACT(EPOCH FROM ts) * 1000)::BIGINT;

    -- Get random part from old UUID (preserve uniqueness)
    uuid_bytes := uuid_send(old_uuid);
    random_part := substring(uuid_bytes FROM 7 FOR 10);

    -- Construct new UUID v7
    uuid_bytes :=
        substring(int8send(unix_ts_ms) FROM 3 FOR 6) ||
        random_part;

    -- Set version and variant bits
    uuid_bytes := set_byte(uuid_bytes, 6, (get_byte(uuid_bytes, 6) & 15) | 112);
    uuid_bytes := set_byte(uuid_bytes, 8, (get_byte(uuid_bytes, 8) & 63) | 128);

    RETURN encode(uuid_bytes, 'hex')::uuid;
END;
$$;


--
-- Name: FUNCTION uuid_v7_from_timestamp(ts timestamp without time zone, old_uuid uuid); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.uuid_v7_from_timestamp(ts timestamp without time zone, old_uuid uuid) IS 'Converts UUID v4 to v7 using provided timestamp. Used for data migration only.';


--
-- Name: uuid_v7_timestamp(uuid); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.uuid_v7_timestamp(u uuid) RETURNS timestamp without time zone
    LANGUAGE sql IMMUTABLE
    AS $$
SELECT to_timestamp(
    ('x' || substring(u::TEXT, 1, 8) || substring(u::TEXT, 10, 4))::BIT(48)::BIGINT / 1000.0
);
$$;


--
-- Name: FUNCTION uuid_v7_timestamp(u uuid); Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON FUNCTION public.uuid_v7_timestamp(u uuid) IS 'Extracts timestamp from UUID v7. Useful for debugging and analysis.';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: anamnesis; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anamnesis (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    patient_id uuid NOT NULL,
    author_id uuid NOT NULL,
    consultation_date timestamp with time zone NOT NULL,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    content text,
    summary text,
    visibility character varying(20) DEFAULT 'all'::character varying NOT NULL,
    anamnesis_template_id uuid,
    content_html text,
    summary_html text,
    CONSTRAINT chk_anamnesis_visibility CHECK (((visibility)::text = ANY ((ARRAY['all'::character varying, 'medicalOnly'::character varying, 'psychOnly'::character varying])::text[])))
);


--
-- Name: COLUMN anamnesis.author_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.anamnesis.author_id IS 'ID do profissional que realizou a anamnese (médico, psicólogo, etc.)';


--
-- Name: COLUMN anamnesis.content; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.anamnesis.content IS 'Plain text content for search and indexing';


--
-- Name: COLUMN anamnesis.summary; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.anamnesis.summary IS 'Plain text summary for search and indexing';


--
-- Name: COLUMN anamnesis.content_html; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.anamnesis.content_html IS 'HTML formatted content for display';


--
-- Name: COLUMN anamnesis.summary_html; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.anamnesis.summary_html IS 'HTML formatted summary for display';


--
-- Name: anamnesis_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anamnesis_items (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    anamnesis_id uuid NOT NULL,
    score_item_id uuid NOT NULL,
    text_value text,
    numeric_value double precision,
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    selected_level integer
);


--
-- Name: anamnesis_template_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anamnesis_template_items (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    anamnesis_template_id uuid NOT NULL,
    score_item_id uuid NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: anamnesis_templates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anamnesis_templates (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    name character varying(200) NOT NULL,
    area character varying(100) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: anonymous_score_group_results; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anonymous_score_group_results (
    id uuid NOT NULL,
    snapshot_id uuid NOT NULL,
    group_id uuid NOT NULL,
    actual_points double precision DEFAULT 0 NOT NULL,
    possible_points double precision DEFAULT 0 NOT NULL,
    score_percentage double precision DEFAULT 0 NOT NULL,
    items_evaluated_count integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: anonymous_score_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anonymous_score_items (
    id uuid NOT NULL,
    session_id uuid NOT NULL,
    score_item_id uuid NOT NULL,
    numeric_value double precision,
    text_value text,
    selected_level integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_anonymous_score_items_selected_level CHECK (((selected_level >= 0) AND (selected_level <= 6)))
);


--
-- Name: anonymous_score_sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anonymous_score_sessions (
    id uuid NOT NULL,
    public_code character varying(12) NOT NULL,
    age integer NOT NULL,
    gender character varying(20) NOT NULL,
    post_menopause boolean,
    height double precision,
    weight double precision,
    claimed_by_patient_id uuid,
    claimed_at timestamp with time zone,
    email character varying(255),
    expires_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    consent_version character varying(20),
    consent_timestamp timestamp with time zone,
    consent_ip_hash character varying(64),
    phone character varying(20),
    email_opt_in boolean DEFAULT false,
    whats_app_opt_in boolean DEFAULT false,
    utm_source character varying(80),
    utm_medium character varying(80),
    utm_campaign character varying(120),
    utm_term character varying(120),
    CONSTRAINT chk_anonymous_score_sessions_age CHECK (((age >= 0) AND (age <= 150))),
    CONSTRAINT chk_anonymous_score_sessions_gender CHECK (((gender)::text = ANY ((ARRAY['male'::character varying, 'female'::character varying, 'other'::character varying])::text[])))
);


--
-- Name: anonymous_score_snapshots; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.anonymous_score_snapshots (
    id uuid NOT NULL,
    session_id uuid NOT NULL,
    total_actual_points double precision DEFAULT 0 NOT NULL,
    total_possible_points double precision DEFAULT 0 NOT NULL,
    total_score_percentage double precision DEFAULT 0 NOT NULL,
    items_evaluated_count integer DEFAULT 0 NOT NULL,
    items_not_evaluated_count integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: api_usage_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.api_usage_logs (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    provider character varying(50) NOT NULL,
    model character varying(100) NOT NULL,
    endpoint character varying(100) NOT NULL,
    input_tokens integer,
    output_tokens integer,
    total_tokens integer,
    cost_usd numeric(10,6),
    user_id uuid,
    metadata jsonb DEFAULT '{}'::jsonb,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT api_usage_logs_provider_check CHECK (((provider)::text = ANY ((ARRAY['openai'::character varying, 'anthropic'::character varying, 'voyage'::character varying])::text[]))),
    CONSTRAINT api_usage_logs_total_tokens_check CHECK ((total_tokens >= 0))
);


--
-- Name: TABLE api_usage_logs; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.api_usage_logs IS 'Tracks API usage for cost monitoring and debugging. OpenAI embeddings: $0.13/1M tokens.';


--
-- Name: appointment_payments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.appointment_payments (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    appointment_id uuid,
    amount_cents bigint NOT NULL,
    method character varying(20) NOT NULL,
    status character varying(20) DEFAULT 'paid'::character varying NOT NULL,
    paid_at timestamp with time zone NOT NULL,
    receipt_number character varying(20) NOT NULL,
    notes text,
    refunded_at timestamp with time zone,
    refund_reason text,
    created_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_appointment_payments_method CHECK (((method)::text = ANY ((ARRAY['cash'::character varying, 'pix'::character varying, 'card'::character varying, 'transfer'::character varying, 'other'::character varying])::text[]))),
    CONSTRAINT chk_appointment_payments_status CHECK (((status)::text = ANY ((ARRAY['paid'::character varying, 'refunded'::character varying])::text[])))
);


--
-- Name: appointment_resources; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.appointment_resources (
    id uuid NOT NULL,
    appointment_id uuid NOT NULL,
    resource_type character varying(20) NOT NULL,
    resource_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_appointment_resources_resource_type CHECK (((resource_type)::text = ANY ((ARRAY['doctor'::character varying, 'room'::character varying, 'equipment'::character varying])::text[])))
);


--
-- Name: appointments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.appointments (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    patient_id uuid NOT NULL,
    doctor_id uuid NOT NULL,
    scheduled_at timestamp with time zone NOT NULL,
    duration_minutes bigint DEFAULT 30 NOT NULL,
    type character varying(20) NOT NULL,
    status character varying(20) DEFAULT 'scheduled'::character varying NOT NULL,
    reason text NOT NULL,
    patient_notes text,
    doctor_notes text,
    diagnosis text,
    anamnesis_id uuid,
    confirmed_at timestamp with time zone,
    completed_at timestamp with time zone,
    cancelled_at timestamp with time zone,
    cancellation_reason text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    external_calendar_event_id character varying(255),
    daily_room_url character varying(255),
    daily_room_name character varying(255),
    confirmation_sent_at timestamp with time zone,
    reminder_sent_at timestamp with time zone,
    continuum_item_id uuid,
    patient_confirmed_at timestamp with time zone,
    push_reminder1h_sent_at timestamp with time zone,
    end_at timestamp with time zone NOT NULL,
    checked_in_at timestamp with time zone,
    started_at timestamp with time zone,
    CONSTRAINT chk_appointments_status CHECK (((status)::text = ANY ((ARRAY['scheduled'::character varying, 'confirmed'::character varying, 'checked_in'::character varying, 'in_progress'::character varying, 'completed'::character varying, 'cancelled'::character varying, 'no_show'::character varying])::text[]))),
    CONSTRAINT chk_appointments_type CHECK (((type)::text = ANY ((ARRAY['initial_assessment'::character varying, 'follow_up'::character varying, 'telemedicine'::character varying, 'procedure'::character varying, 'results_review'::character varying])::text[])))
);


--
-- Name: article_embeddings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.article_embeddings (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    article_id uuid NOT NULL,
    chunk_index integer NOT NULL,
    chunk_text text NOT NULL,
    chunk_metadata jsonb DEFAULT '{}'::jsonb,
    embedding public.vector(1024),
    created_at timestamp with time zone,
    is_stale boolean DEFAULT false NOT NULL
);


--
-- Name: TABLE article_embeddings; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.article_embeddings IS 'Stores vector embeddings for article chunks. Each article is split into semantic chunks (abstract + sliding window). Used for RAG semantic search.';


--
-- Name: COLUMN article_embeddings.embedding; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_embeddings.embedding IS '1024-dimensional vector from OpenAI text-embedding-3-large. Cosine similarity used for semantic search.';


--
-- Name: article_score_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.article_score_items (
    score_item_id uuid DEFAULT gen_random_uuid() NOT NULL,
    article_id uuid DEFAULT gen_random_uuid() NOT NULL,
    confidence_score double precision DEFAULT 1.0,
    auto_linked boolean DEFAULT false,
    linked_at timestamp with time zone DEFAULT now(),
    linked_by uuid,
    user_feedback character varying(20),
    feedback_at timestamp with time zone,
    feedback_by uuid,
    CONSTRAINT article_score_items_user_feedback_check CHECK (((user_feedback)::text = ANY ((ARRAY['approved'::character varying, 'rejected'::character varying, 'irrelevant'::character varying])::text[]))),
    CONSTRAINT check_article_score_items_confidence CHECK (((confidence_score >= (0.0)::double precision) AND (confidence_score <= (1.0)::double precision)))
);


--
-- Name: COLUMN article_score_items.confidence_score; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.confidence_score IS 'Cosine similarity score (0.0-1.0) if auto-linked by RAG, or 1.0 if manually linked. Threshold: 0.7 for auto-suggestions.';


--
-- Name: COLUMN article_score_items.auto_linked; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.auto_linked IS 'True if link was suggested by RAG semantic search, false if manually linked by user.';


--
-- Name: COLUMN article_score_items.linked_by; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.linked_by IS 'User ID who created the link (NULL for system auto-links).';


--
-- Name: COLUMN article_score_items.user_feedback; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.user_feedback IS 'User feedback on link quality: approved (good match), rejected (bad match), irrelevant (not related)';


--
-- Name: COLUMN article_score_items.feedback_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.feedback_at IS 'Timestamp when feedback was provided';


--
-- Name: COLUMN article_score_items.feedback_by; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article_score_items.feedback_by IS 'User ID who provided feedback';


--
-- Name: article_score_items_backup_20260217; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.article_score_items_backup_20260217 (
    score_item_id uuid,
    article_id uuid,
    confidence_score double precision,
    auto_linked boolean,
    linked_at timestamp with time zone,
    linked_by uuid,
    user_feedback character varying(20),
    feedback_at timestamp with time zone,
    feedback_by uuid
);


--
-- Name: articles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.articles (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    title character varying(1000) NOT NULL,
    authors character varying(2000) NOT NULL,
    journal character varying(500) NOT NULL,
    publish_date date NOT NULL,
    language character varying(10) DEFAULT 'en'::character varying NOT NULL,
    doi character varying(255),
    pm_id character varying(50),
    issn character varying(20),
    abstract text,
    full_content text,
    notes text,
    original_link character varying(2048),
    internal_link character varying(2048),
    article_type character varying(50) DEFAULT 'research_article'::character varying NOT NULL,
    keywords text,
    mesh_terms text,
    specialty character varying(200),
    favorite boolean DEFAULT false,
    rating smallint,
    file_hash character varying(64),
    file_size bigint,
    indexed_at timestamp with time zone,
    last_accessed_at timestamp with time zone,
    created_by uuid,
    updated_by uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    embedding_status character varying(20) DEFAULT 'pending'::character varying,
    chunk_count integer DEFAULT 0,
    last_embedded_at timestamp with time zone,
    source_type character varying(20) DEFAULT 'article'::character varying NOT NULL,
    parent_article_id uuid,
    chapter_number smallint,
    chapter_title character varying(500),
    total_chapters smallint,
    CONSTRAINT check_articles_embedding_status CHECK (((embedding_status)::text = ANY (ARRAY[('pending'::character varying)::text, ('processing'::character varying)::text, ('completed'::character varying)::text, ('failed'::character varying)::text]))),
    CONSTRAINT chk_articles_article_type CHECK (((article_type)::text = ANY ((ARRAY['research_article'::character varying, 'review'::character varying, 'meta_analysis'::character varying, 'case_study'::character varying, 'clinical_trial'::character varying, 'editorial'::character varying, 'letter'::character varying, 'protocol'::character varying, 'lecture'::character varying])::text[]))),
    CONSTRAINT chk_articles_embedding_status CHECK (((embedding_status)::text = ANY (ARRAY[('pending'::character varying)::text, ('processing'::character varying)::text, ('completed'::character varying)::text, ('failed'::character varying)::text]))),
    CONSTRAINT chk_articles_rating CHECK (((rating >= 0) AND (rating <= 5)))
);


--
-- Name: COLUMN articles.embedding_status; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.articles.embedding_status IS 'Status of embedding generation: pending (not started), processing (in progress), completed (done), failed (error).';


--
-- Name: COLUMN articles.chunk_count; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.articles.chunk_count IS 'Number of chunks generated for this article (typically 1 abstract + N content chunks).';


--
-- Name: audit_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.audit_logs (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    user_id uuid NOT NULL,
    action character varying(20) NOT NULL,
    resource character varying(50) NOT NULL,
    resource_id uuid,
    ip_address character varying(45) NOT NULL,
    user_agent text,
    success boolean DEFAULT true NOT NULL,
    error_message text,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_audit_logs_action CHECK (((action)::text = ANY ((ARRAY['view'::character varying, 'create'::character varying, 'update'::character varying, 'delete'::character varying])::text[])))
);


--
-- Name: auto_link_batch_checkpoints; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auto_link_batch_checkpoints (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    run_id uuid NOT NULL,
    batch_number integer NOT NULL,
    batch_start_index integer NOT NULL,
    batch_end_index integer NOT NULL,
    items_processed integer NOT NULL,
    links_created integer NOT NULL,
    items_skipped integer NOT NULL,
    items_failed integer NOT NULL,
    avg_confidence double precision,
    processing_time_ms integer,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: TABLE auto_link_batch_checkpoints; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.auto_link_batch_checkpoints IS 'Detailed checkpoint data for each batch processed';


--
-- Name: auto_link_item_log; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auto_link_item_log (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    run_id uuid NOT NULL,
    score_item_id uuid NOT NULL,
    score_item_name character varying(300),
    status character varying(20) NOT NULL,
    links_created integer DEFAULT 0,
    avg_confidence double precision,
    error_message text,
    processing_time_ms integer,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT auto_link_item_log_status_check CHECK (((status)::text = ANY ((ARRAY['success'::character varying, 'skipped'::character varying, 'failed'::character varying])::text[])))
);


--
-- Name: TABLE auto_link_item_log; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.auto_link_item_log IS 'Item-level processing log for debugging and retry logic';


--
-- Name: auto_link_processing_state; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auto_link_processing_state (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    run_id uuid NOT NULL,
    batch_size integer DEFAULT 50 NOT NULL,
    last_processed_item_id uuid,
    last_processed_index integer DEFAULT 0,
    total_items integer DEFAULT 0,
    total_processed integer DEFAULT 0,
    total_linked integer DEFAULT 0,
    total_skipped integer DEFAULT 0,
    failed_items jsonb DEFAULT '[]'::jsonb,
    started_at timestamp with time zone DEFAULT now() NOT NULL,
    last_checkpoint_at timestamp with time zone,
    completed_at timestamp with time zone,
    status character varying(20) DEFAULT 'running'::character varying NOT NULL,
    error_message text,
    metadata jsonb DEFAULT '{}'::jsonb,
    CONSTRAINT auto_link_processing_state_status_check CHECK (((status)::text = ANY ((ARRAY['running'::character varying, 'completed'::character varying, 'failed'::character varying, 'cancelled'::character varying])::text[])))
);


--
-- Name: TABLE auto_link_processing_state; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.auto_link_processing_state IS 'Tracks state of auto-link batch processing runs for crash recovery';


--
-- Name: COLUMN auto_link_processing_state.run_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auto_link_processing_state.run_id IS 'Unique identifier for this processing run';


--
-- Name: COLUMN auto_link_processing_state.last_processed_item_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auto_link_processing_state.last_processed_item_id IS 'UUID of last successfully processed ScoreItem';


--
-- Name: COLUMN auto_link_processing_state.failed_items; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auto_link_processing_state.failed_items IS 'JSONB array of failed item IDs for retry';


--
-- Name: calendar_credentials; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.calendar_credentials (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    provider character varying(20) NOT NULL,
    google_account_email character varying(255) NOT NULL,
    encrypted_refresh_token text NOT NULL,
    encrypted_access_token text NOT NULL,
    access_token_expires_at timestamp with time zone NOT NULL,
    dedicated_calendar_id character varying(255) NOT NULL,
    last_sync_at timestamp with time zone,
    revoked_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);


--
-- Name: campaigns; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.campaigns (
    id uuid NOT NULL,
    name character varying(160) NOT NULL,
    slug character varying(120) NOT NULL,
    description text,
    landing_path character varying(255) DEFAULT '/escore-plenya/painel'::character varying NOT NULL,
    utm_source character varying(80) NOT NULL,
    utm_medium character varying(80) NOT NULL,
    utm_campaign character varying(120) NOT NULL,
    utm_term character varying(120),
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    created_by_user_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: consultation_prices; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.consultation_prices (
    id uuid NOT NULL,
    type character varying(20) NOT NULL,
    amount_cents bigint DEFAULT 0 NOT NULL,
    active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: continuum_box_templates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.continuum_box_templates (
    id uuid NOT NULL,
    name character varying(120) NOT NULL,
    description text,
    contents text,
    notes text,
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    created_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_continuum_box_templates_status CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'archived'::character varying])::text[])))
);


--
-- Name: continuum_template_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.continuum_template_items (
    id uuid NOT NULL,
    template_id uuid NOT NULL,
    type character varying(30) NOT NULL,
    specialty character varying(30),
    title character varying(160) NOT NULL,
    description text,
    week_offset bigint NOT NULL,
    expected_offset_days bigint DEFAULT 0 NOT NULL,
    late_after_days bigint DEFAULT 7 NOT NULL,
    box_template_id uuid,
    "position" bigint DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_continuum_template_items_expected_offset_days CHECK (((expected_offset_days >= 0) AND (expected_offset_days <= 6))),
    CONSTRAINT chk_continuum_template_items_late_after_days CHECK ((late_after_days >= 0)),
    CONSTRAINT chk_continuum_template_items_type CHECK (((type)::text = ANY ((ARRAY['appointment'::character varying, 'box'::character varying, 'reassessment'::character varying, 'milestone'::character varying, 'custom'::character varying])::text[]))),
    CONSTRAINT chk_continuum_template_items_week_offset CHECK ((week_offset >= 0))
);


--
-- Name: continuum_templates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.continuum_templates (
    id uuid NOT NULL,
    name character varying(120) NOT NULL,
    description text,
    duration_weeks bigint NOT NULL,
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    created_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_continuum_templates_duration_weeks CHECK ((duration_weeks > 0)),
    CONSTRAINT chk_continuum_templates_status CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'archived'::character varying])::text[])))
);


--
-- Name: conversation_reads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.conversation_reads (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    owner_type character varying(20) NOT NULL,
    owner_id uuid NOT NULL,
    last_read_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    CONSTRAINT conversation_reads_owner_type_check CHECK (((owner_type)::text = ANY ((ARRAY['lead'::character varying, 'patient'::character varying])::text[])))
);


--
-- Name: device_tokens; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.device_tokens (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    platform character varying(10) NOT NULL,
    app_variant character varying(10) NOT NULL,
    token text NOT NULL,
    device_label character varying(200),
    app_version character varying(40),
    last_seen_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: doctor_absences; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.doctor_absences (
    id uuid NOT NULL,
    doctor_id uuid NOT NULL,
    start_at timestamp with time zone NOT NULL,
    end_at timestamp with time zone NOT NULL,
    reason character varying(100),
    created_at timestamp with time zone NOT NULL
);


--
-- Name: email_ingest_states; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.email_ingest_states (
    id uuid NOT NULL,
    account character varying(255) NOT NULL,
    folder character varying(100) NOT NULL,
    last_uid bigint DEFAULT 0 NOT NULL,
    last_seen_at timestamp with time zone,
    created_at timestamp with time zone
);


--
-- Name: embedding_audit_log; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.embedding_audit_log (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    entity_type character varying(50) NOT NULL,
    entity_id uuid NOT NULL,
    action character varying(50) NOT NULL,
    reason text,
    old_text_source text,
    new_text_source text,
    triggered_by character varying(100),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT embedding_audit_log_action_check CHECK (((action)::text = ANY ((ARRAY['invalidated'::character varying, 'regenerated'::character varying, 'failed'::character varying])::text[])))
);


--
-- Name: TABLE embedding_audit_log; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.embedding_audit_log IS 'Audit trail for all embedding invalidations and regenerations';


--
-- Name: embedding_queue; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.embedding_queue (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    entity_type character varying(50) NOT NULL,
    entity_id uuid NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    retry_count integer DEFAULT 0,
    error_message text,
    created_at timestamp with time zone,
    processed_at timestamp with time zone,
    priority integer DEFAULT 0,
    max_retries integer DEFAULT 3,
    metadata jsonb DEFAULT '{}'::jsonb,
    CONSTRAINT chk_embedding_queue_entity_type CHECK (((entity_type)::text = ANY ((ARRAY['article'::character varying, 'score_item'::character varying])::text[]))),
    CONSTRAINT chk_embedding_queue_retry_count CHECK (((retry_count >= 0) AND (retry_count <= 5))),
    CONSTRAINT chk_embedding_queue_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'processing'::character varying, 'completed'::character varying, 'failed'::character varying])::text[]))),
    CONSTRAINT embedding_queue_entity_type_check CHECK (((entity_type)::text = ANY ((ARRAY['article'::character varying, 'score_item'::character varying])::text[]))),
    CONSTRAINT embedding_queue_retry_count_check CHECK (((retry_count >= 0) AND (retry_count <= 5))),
    CONSTRAINT embedding_queue_status_check CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'processing'::character varying, 'completed'::character varying, 'failed'::character varying])::text[])))
);


--
-- Name: TABLE embedding_queue; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.embedding_queue IS 'Queue for asynchronous regeneration of stale embeddings';


--
-- Name: COLUMN embedding_queue.retry_count; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.embedding_queue.retry_count IS 'Number of times this item has failed regeneration';


--
-- Name: exercises; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.exercises (
    id uuid NOT NULL,
    external_id character varying(20) NOT NULL,
    name character varying(200) NOT NULL,
    name_pt character varying(200),
    body_parts jsonb,
    body_parts_pt jsonb,
    target_muscles jsonb,
    target_muscles_pt jsonb,
    equipments jsonb,
    equipments_pt jsonb,
    secondary_muscles jsonb,
    instructions jsonb,
    instructions_pt jsonb,
    gif_url text,
    gif_url_fallback text,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    biomechanics_data jsonb,
    program_design jsonb,
    nsca_references jsonb
);


--
-- Name: fitness_test_results; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.fitness_test_results (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    created_by_id uuid NOT NULL,
    assessment_date date NOT NULL,
    abdominal_reps bigint,
    pushup_reps bigint,
    plank_seconds bigint,
    burpee_cycles bigint,
    frt_reps bigint,
    abdominal_level character varying(20),
    pushup_level character varying(20),
    plank_level character varying(20),
    burpee_level character varying(20),
    frt_level character varying(20),
    overall_score bigint DEFAULT 0 NOT NULL,
    overall_classification character varying(30) DEFAULT ''::character varying NOT NULL,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: health_check_ins; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.health_check_ins (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    energy bigint NOT NULL,
    pain bigint NOT NULL,
    pain_location character varying(200),
    mood bigint NOT NULL,
    sleep_hours numeric(4,1) NOT NULL,
    sleep_quality bigint NOT NULL,
    stress bigint NOT NULL,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_health_check_ins_energy CHECK (((energy >= 1) AND (energy <= 5))),
    CONSTRAINT chk_health_check_ins_mood CHECK (((mood >= 1) AND (mood <= 5))),
    CONSTRAINT chk_health_check_ins_pain CHECK (((pain >= 0) AND (pain <= 5))),
    CONSTRAINT chk_health_check_ins_sleep_hours CHECK (((sleep_hours >= (0)::numeric) AND (sleep_hours <= (24)::numeric))),
    CONSTRAINT chk_health_check_ins_sleep_quality CHECK (((sleep_quality >= 1) AND (sleep_quality <= 5))),
    CONSTRAINT chk_health_check_ins_stress CHECK (((stress >= 1) AND (stress <= 5)))
);


--
-- Name: integrated_plan_revisions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.integrated_plan_revisions (
    id uuid NOT NULL,
    continuum_id uuid NOT NULL,
    content text NOT NULL,
    updated_by_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL
);


--
-- Name: lab_request_template_tests; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_request_template_tests (
    lab_request_template_id uuid NOT NULL,
    lab_test_definition_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: TABLE lab_request_template_tests; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_request_template_tests IS 'Relação many-to-many entre templates e exames';


--
-- Name: lab_request_templates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_request_templates (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    name character varying(200) NOT NULL,
    description text,
    display_order integer DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: TABLE lab_request_templates; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_request_templates IS 'Templates pré-configurados de pedidos de exames';


--
-- Name: lab_requests; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_requests (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    patient_id uuid NOT NULL,
    date date NOT NULL,
    exams text NOT NULL,
    notes text,
    doctor_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    lab_request_template_id uuid,
    pdf_url text,
    signed_pdf_path character varying(500),
    signed_pdf_hash character varying(64),
    qr_code_data text,
    signed_at timestamp with time zone,
    certificate_serial character varying(100)
);


--
-- Name: TABLE lab_requests; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_requests IS 'Pedidos de exames laboratoriais';


--
-- Name: COLUMN lab_requests.lab_request_template_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_requests.lab_request_template_id IS 'ID do template utilizado para criar este pedido (opcional)';


--
-- Name: COLUMN lab_requests.pdf_url; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_requests.pdf_url IS 'URL do PDF gerado. Quando preenchido, o pedido fica bloqueado para edição';


--
-- Name: COLUMN lab_requests.signed_pdf_path; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_requests.signed_pdf_path IS 'Caminho do PDF de pedido de exames assinado digitalmente';


--
-- Name: COLUMN lab_requests.qr_code_data; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_requests.qr_code_data IS 'QR Code para validação pública do pedido de exames';


--
-- Name: lab_result_batches; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_result_batches (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    lab_request_id uuid,
    requesting_doctor_id uuid,
    laboratory_name character varying(200) NOT NULL,
    collection_date timestamp with time zone NOT NULL,
    result_date timestamp with time zone,
    status character varying(20) DEFAULT 'pending'::character varying NOT NULL,
    observations text,
    attachments text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    pdf_content_full text,
    pdf_content_short text,
    pdf_content_need_ai text,
    pdf_content_json text,
    CONSTRAINT chk_lab_result_batches_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'partial'::character varying, 'completed'::character varying])::text[]))),
    CONSTRAINT lab_result_batches_status_check CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'partial'::character varying, 'completed'::character varying])::text[])))
);


--
-- Name: COLUMN lab_result_batches.pdf_content_json; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_result_batches.pdf_content_json IS 'Conteúdo processado pela IA em formato JSON estruturado';


--
-- Name: lab_result_values; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_result_values (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    lab_result_id uuid NOT NULL,
    lab_test_definition_id uuid NOT NULL,
    numeric_value double precision,
    text_value text,
    boolean_value boolean,
    unit character varying(50),
    notes text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: TABLE lab_result_values; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_result_values IS 'Valores estruturados de resultados de exames laboratoriais. Faixas de referência e classificação (anormal/crítico) são determinadas pelo sistema de Score.';


--
-- Name: COLUMN lab_result_values.unit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_result_values.unit IS 'Unidade de medida do valor (pode sobrescrever a unidade padrão do exame)';


--
-- Name: lab_result_view_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_result_view_items (
    id uuid NOT NULL,
    lab_result_view_id uuid NOT NULL,
    lab_test_definition_id uuid NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: lab_result_views; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_result_views (
    id uuid NOT NULL,
    name character varying(200) NOT NULL,
    description text,
    is_active boolean DEFAULT true NOT NULL,
    display_order integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: lab_results; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_results (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    test_name character varying(200) NOT NULL,
    test_type character varying(100) NOT NULL,
    unit character varying(50),
    interpretation text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    lab_result_batch_id uuid NOT NULL,
    lab_test_definition_id uuid,
    result_text text,
    result_numeric numeric(12,4),
    level integer,
    matched boolean DEFAULT true NOT NULL,
    result_numeric_original numeric(12,4),
    unit_original character varying(50)
);


--
-- Name: COLUMN lab_results.unit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_results.unit IS 'Unidade CONVERTIDA para unidade padrão (ou original se sem conversão)';


--
-- Name: COLUMN lab_results.result_numeric; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_results.result_numeric IS 'Valor numérico CONVERTIDO para unidade padrão (ou original se sem conversão)';


--
-- Name: COLUMN lab_results.result_numeric_original; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_results.result_numeric_original IS 'Valor numérico ORIGINAL (antes da conversão)';


--
-- Name: COLUMN lab_results.unit_original; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_results.unit_original IS 'Unidade ORIGINAL (antes da conversão)';


--
-- Name: lab_test_definitions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_test_definitions (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    code character varying(100) NOT NULL,
    name character varying(300) NOT NULL,
    short_name character varying(50),
    tuss_code character varying(20),
    loinc_code character varying(20),
    category character varying(30) NOT NULL,
    is_requestable boolean DEFAULT true NOT NULL,
    parent_test_id uuid,
    unit character varying(50),
    unit_conversion text,
    result_type character varying(20) DEFAULT 'numeric'::character varying NOT NULL,
    collection_method text,
    fasting_hours integer,
    specimen_type character varying(100),
    description text,
    clinical_indications text,
    display_order integer DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    clinical_significance text,
    longevity_context text,
    clinical_recommendations text,
    alt_names jsonb,
    CONSTRAINT chk_category CHECK (((category)::text = ANY ((ARRAY['hematology'::character varying, 'biochemistry'::character varying, 'hormones'::character varying, 'immunology'::character varying, 'microbiology'::character varying, 'urine'::character varying, 'imaging'::character varying, 'functional'::character varying, 'genetics'::character varying, 'other'::character varying])::text[]))),
    CONSTRAINT chk_result_type CHECK (((result_type)::text = ANY ((ARRAY['numeric'::character varying, 'text'::character varying, 'boolean'::character varying, 'categorical'::character varying])::text[])))
);


--
-- Name: TABLE lab_test_definitions; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_test_definitions IS 'Definições estruturadas de exames laboratoriais e seus parâmetros';


--
-- Name: COLUMN lab_test_definitions.is_requestable; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_definitions.is_requestable IS 'Indica se o exame pode ser solicitado individualmente (true) ou só aparece no resultado (false)';


--
-- Name: COLUMN lab_test_definitions.parent_test_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_definitions.parent_test_id IS 'ID do exame pai para criar hierarquia (ex: Hemoglobina tem parent=Hemograma)';


--
-- Name: COLUMN lab_test_definitions.clinical_significance; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_definitions.clinical_significance IS 'Significância clínica detalhada (200-400 palavras): mecanismos fisiológicos, aplicações clínicas, interpretação de valores alterados';


--
-- Name: COLUMN lab_test_definitions.longevity_context; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_definitions.longevity_context IS 'Contexto de longevidade (100-200 palavras): relação com envelhecimento saudável, marcadores de longevidade, implicações preventivas';


--
-- Name: COLUMN lab_test_definitions.clinical_recommendations; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_definitions.clinical_recommendations IS 'Recomendações clínicas (150-300 palavras): quando solicitar, interpretação de resultados, fatores que afetam valores, intervenções baseadas em resultados';


--
-- Name: lab_test_definitions_backup_consolidation; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_test_definitions_backup_consolidation (
    id uuid,
    code character varying(100),
    name character varying(300),
    short_name character varying(50),
    tuss_code character varying(20),
    loinc_code character varying(20),
    category character varying(30),
    is_requestable boolean,
    parent_test_id uuid,
    unit character varying(50),
    unit_conversion text,
    result_type character varying(20),
    collection_method text,
    fasting_hours integer,
    specimen_type character varying(100),
    description text,
    clinical_indications text,
    display_order integer,
    is_active boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    clinical_significance text,
    longevity_context text,
    clinical_recommendations text,
    alt_names jsonb
);


--
-- Name: lab_test_unit_conversions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lab_test_unit_conversions (
    id uuid NOT NULL,
    lab_test_definition_id uuid NOT NULL,
    main_unit character varying(50) NOT NULL,
    secondary_unit character varying(50) NOT NULL,
    conversion_factor double precision NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_lab_test_unit_conversions_conversion_factor CHECK ((conversion_factor > (0)::double precision)),
    CONSTRAINT lab_test_unit_conversions_conversion_factor_check CHECK ((conversion_factor > (0)::double precision))
);


--
-- Name: TABLE lab_test_unit_conversions; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.lab_test_unit_conversions IS 'Conversões de unidades para exames laboratoriais';


--
-- Name: COLUMN lab_test_unit_conversions.main_unit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_unit_conversions.main_unit IS 'Unidade principal/padrão (ex: g/dL)';


--
-- Name: COLUMN lab_test_unit_conversions.secondary_unit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_unit_conversions.secondary_unit IS 'Unidade alternativa/secundária (ex: g/L)';


--
-- Name: COLUMN lab_test_unit_conversions.conversion_factor; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.lab_test_unit_conversions.conversion_factor IS 'Fator: secondaryValue = mainValue * factor';


--
-- Name: lead_activities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.lead_activities (
    id uuid NOT NULL,
    lead_id uuid,
    type character varying(40) NOT NULL,
    channel character varying(20),
    content text,
    metadata jsonb,
    actor_user_id uuid,
    created_at timestamp with time zone,
    patient_id uuid,
    media_type character varying(20),
    media_storage_key character varying(500),
    media_mime character varying(120),
    media_filename character varying(255),
    media_size_bytes bigint,
    patient_document_id uuid,
    transcription text,
    CONSTRAINT lead_activities_owner_check CHECK ((((lead_id IS NOT NULL) AND (patient_id IS NULL)) OR ((lead_id IS NULL) AND (patient_id IS NOT NULL))))
);


--
-- Name: leads; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.leads (
    id uuid NOT NULL,
    source character varying(40) NOT NULL,
    status character varying(40) DEFAULT 'new'::character varying NOT NULL,
    name character varying(255),
    email character varying(255),
    phone character varying(20),
    message text,
    metadata jsonb,
    email_opt_in boolean DEFAULT false,
    whats_app_opt_in boolean DEFAULT false,
    newsletter_opt_in boolean DEFAULT false,
    consent_version character varying(20),
    consent_timestamp timestamp with time zone,
    consent_ip_hash character varying(64),
    anonymous_score_session_id uuid,
    converted_patient_id uuid,
    converted_at timestamp with time zone,
    converted_by_user_id uuid,
    assigned_to_user_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    last_inbound_at timestamp with time zone,
    utm_source character varying(80),
    utm_medium character varying(80),
    utm_campaign character varying(120),
    utm_term character varying(120)
);


--
-- Name: medication_definitions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.medication_definitions (
    id uuid NOT NULL,
    common_name character varying(500) NOT NULL,
    active_ingredient character varying(500) NOT NULL,
    category character varying(20) NOT NULL,
    validity_days integer DEFAULT 30 NOT NULL,
    max_per_prescription integer DEFAULT 10 NOT NULL,
    max_treatment_days integer DEFAULT 60 NOT NULL,
    requires_digital_signature boolean DEFAULT false NOT NULL,
    requires_sncr boolean DEFAULT false NOT NULL,
    anvisa_code character varying(50),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_medication_definitions_category CHECK (((category)::text = ANY (ARRAY[('simple'::character varying)::text, ('c1'::character varying)::text, ('c5'::character varying)::text, ('antibiotic'::character varying)::text, ('glp1'::character varying)::text])))
);


--
-- Name: TABLE medication_definitions; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.medication_definitions IS 'Catálogo de medicamentos com regras regulatórias (ANVISA/CFM)';


--
-- Name: COLUMN medication_definitions.category; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.medication_definitions.category IS 'Categoria regulatória: simple (receita simples), c1 (controle especial), c5 (psicotrópicos), antibiotic (antimicrobianos), glp1 (GLP-1 agonistas)';


--
-- Name: COLUMN medication_definitions.requires_sncr; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.medication_definitions.requires_sncr IS 'Se requer registro no SNCR (Sistema Nacional de Controle de Receitas) da ANVISA';


--
-- Name: method_letters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.method_letters (
    id uuid NOT NULL,
    code character varying(10) NOT NULL,
    name character varying(300) NOT NULL,
    description text,
    clinical_relevance text,
    patient_explanation text,
    conduct text,
    last_review timestamp with time zone,
    color character varying(7),
    icon character varying(50),
    "order" integer DEFAULT 0 NOT NULL,
    method_id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: method_pillars; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.method_pillars (
    id uuid NOT NULL,
    name character varying(300) NOT NULL,
    description text,
    clinical_relevance text,
    patient_explanation text,
    conduct text,
    last_review timestamp with time zone,
    "order" integer DEFAULT 0 NOT NULL,
    letter_id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: methods; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.methods (
    id uuid NOT NULL,
    name character varying(200) NOT NULL,
    short_name character varying(20) NOT NULL,
    description text,
    version character varying(20),
    color character varying(7),
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_default boolean DEFAULT false NOT NULL
);


--
-- Name: COLUMN methods.is_default; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.methods.is_default IS 'Indicates if this is the default method (only one can be default at a time)';


--
-- Name: notification_preferences; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notification_preferences (
    user_id uuid NOT NULL,
    appointment_reminder boolean DEFAULT true NOT NULL,
    message_alert boolean DEFAULT true NOT NULL,
    workout_reminder boolean DEFAULT true NOT NULL,
    workout_reminder_time character varying(5) DEFAULT '07:00'::character varying NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: notifications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notifications (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    patient_id uuid,
    subscription_id uuid,
    type character varying(50) NOT NULL,
    title character varying(200) NOT NULL,
    message text NOT NULL,
    action_url character varying(500),
    action_text character varying(50),
    read boolean DEFAULT false NOT NULL,
    read_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    lead_id uuid,
    CONSTRAINT chk_notifications_type CHECK (((type)::text = ANY ((ARRAY['trial_expiring'::character varying, 'renewal_upcoming'::character varying, 'subscription_expired'::character varying, 'payment_pending'::character varying, 'general'::character varying, 'lead_new'::character varying, 'lead_whatsapp_inbound'::character varying, 'lead_email_inbound'::character varying, 'lead_assigned'::character varying])::text[]))),
    CONSTRAINT notifications_type_check CHECK (((type)::text = ANY ((ARRAY['trial_expiring'::character varying, 'renewal_upcoming'::character varying, 'subscription_expired'::character varying, 'payment_pending'::character varying, 'general'::character varying])::text[])))
);


--
-- Name: patient_continuum_boxes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_continuum_boxes (
    id uuid NOT NULL,
    continuum_item_id uuid NOT NULL,
    box_template_id uuid,
    name character varying(120) NOT NULL,
    contents text,
    status character varying(20) DEFAULT 'planned'::character varying NOT NULL,
    prepared_at timestamp with time zone,
    shipped_at timestamp with time zone,
    delivered_at timestamp with time zone,
    tracking_code character varying(60),
    carrier character varying(40),
    address_snapshot text,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_patient_continuum_boxes_status CHECK (((status)::text = ANY ((ARRAY['planned'::character varying, 'preparing'::character varying, 'shipped'::character varying, 'delivered'::character varying, 'cancelled'::character varying])::text[])))
);


--
-- Name: patient_continuum_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_continuum_items (
    id uuid NOT NULL,
    continuum_id uuid NOT NULL,
    type character varying(30) NOT NULL,
    specialty character varying(30),
    title character varying(160) NOT NULL,
    description text,
    week_offset bigint NOT NULL,
    expected_date timestamp with time zone NOT NULL,
    late_after_date timestamp with time zone NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying NOT NULL,
    appointment_id uuid,
    box_id uuid,
    completed_at timestamp with time zone,
    completed_ref_type character varying(40),
    completed_ref_id uuid,
    "position" bigint DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_patient_continuum_items_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'scheduled'::character varying, 'completed'::character varying, 'missed'::character varying, 'cancelled'::character varying, 'skipped'::character varying])::text[]))),
    CONSTRAINT chk_patient_continuum_items_type CHECK (((type)::text = ANY ((ARRAY['appointment'::character varying, 'box'::character varying, 'reassessment'::character varying, 'milestone'::character varying, 'custom'::character varying])::text[])))
);


--
-- Name: patient_continuums; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_continuums (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    template_id uuid NOT NULL,
    template_snapshot jsonb NOT NULL,
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL,
    coordinator_doctor_id uuid,
    integrated_plan_markdown text,
    integrated_plan_updated_at timestamp with time zone,
    integrated_plan_updated_by uuid,
    whatsapp_group_name character varying(160),
    whatsapp_group_invite_link character varying(255),
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_patient_continuums_status CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'paused'::character varying, 'completed'::character varying, 'cancelled'::character varying])::text[])))
);


--
-- Name: patient_documents; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_documents (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    uploaded_by uuid,
    type character varying(30) NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    file_path character varying(500) NOT NULL,
    file_name character varying(200) NOT NULL,
    content_type character varying(100) NOT NULL,
    size_bytes bigint NOT NULL,
    issued_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    source character varying(20) DEFAULT 'staff_upload'::character varying NOT NULL,
    origin_wa_message_id character varying(120),
    CONSTRAINT chk_patient_documents_type CHECK (((type)::text = ANY ((ARRAY['certificate'::character varying, 'report'::character varying, 'referral'::character varying, 'declaration'::character varying, 'other'::character varying])::text[])))
);


--
-- Name: patient_magic_links; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_magic_links (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    token character varying(128),
    expires_at timestamp with time zone NOT NULL,
    used_at timestamp with time zone,
    ip_hash character varying(64),
    created_at timestamp with time zone,
    token_hash character varying(64)
);


--
-- Name: patient_portal_invites; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_portal_invites (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    invited_by uuid NOT NULL,
    token character varying(128),
    expires_at timestamp with time zone NOT NULL,
    accepted_at timestamp with time zone,
    channel_email boolean DEFAULT true,
    channel_wa boolean DEFAULT false,
    created_at timestamp with time zone,
    token_hash character varying(64)
);


--
-- Name: patient_score_group_results; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_score_group_results (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    snapshot_id uuid NOT NULL,
    group_id uuid NOT NULL,
    actual_points double precision DEFAULT 0 NOT NULL,
    possible_points double precision DEFAULT 0 NOT NULL,
    score_percentage double precision DEFAULT 0 NOT NULL,
    items_evaluated_count integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_group_results_actual_points CHECK ((actual_points >= (0)::double precision)),
    CONSTRAINT chk_group_results_items_evaluated CHECK ((items_evaluated_count >= 0)),
    CONSTRAINT chk_group_results_possible_points CHECK ((possible_points >= (0)::double precision)),
    CONSTRAINT chk_group_results_score_percentage CHECK (((score_percentage >= (0)::double precision) AND (score_percentage <= (100)::double precision)))
);


--
-- Name: TABLE patient_score_group_results; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.patient_score_group_results IS 'Resultados agregados por grupo clínico dentro de um snapshot';


--
-- Name: patient_score_item_results; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_score_item_results (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    snapshot_id uuid NOT NULL,
    item_id uuid NOT NULL,
    group_id uuid NOT NULL,
    status character varying(30) NOT NULL,
    data_source character varying(20),
    lab_result_id uuid,
    anamnesis_item_id uuid,
    value_used double precision,
    level_matched_id uuid,
    level_number integer,
    max_points double precision DEFAULT 0 NOT NULL,
    actual_points double precision DEFAULT 0 NOT NULL,
    not_evaluated_reason text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_item_results_actual_points CHECK (((actual_points >= (0)::double precision) AND (actual_points <= (100)::double precision))),
    CONSTRAINT chk_item_results_data_source CHECK ((((data_source)::text = ANY ((ARRAY['lab_result'::character varying, 'anamnesis_item'::character varying])::text[])) OR (data_source IS NULL))),
    CONSTRAINT chk_item_results_level_number CHECK ((((level_number >= 0) AND (level_number <= 6)) OR (level_number IS NULL))),
    CONSTRAINT chk_item_results_max_points CHECK (((max_points >= (0)::double precision) AND (max_points <= (100)::double precision))),
    CONSTRAINT chk_item_results_status CHECK (((status)::text = ANY ((ARRAY['evaluated'::character varying, 'not_applicable'::character varying, 'no_data_available'::character varying])::text[])))
);


--
-- Name: TABLE patient_score_item_results; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.patient_score_item_results IS 'Avaliação individual de cada item de score para o paciente';


--
-- Name: COLUMN patient_score_item_results.status; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_item_results.status IS 'Status da avaliação: evaluated, not_applicable, no_data_available';


--
-- Name: COLUMN patient_score_item_results.data_source; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_item_results.data_source IS 'Fonte dos dados: lab_result ou anamnesis_item';


--
-- Name: COLUMN patient_score_item_results.level_number; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_item_results.level_number IS 'Número do nível atingido (0-6), denormalizado';


--
-- Name: COLUMN patient_score_item_results.actual_points; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_item_results.actual_points IS 'Pontos reais obtidos (proporcional ao nível)';


--
-- Name: patient_score_snapshots; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_score_snapshots (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    patient_id uuid NOT NULL,
    calculated_by_user_id uuid NOT NULL,
    calculated_at timestamp with time zone NOT NULL,
    total_actual_points double precision DEFAULT 0 NOT NULL,
    total_possible_points double precision DEFAULT 0 NOT NULL,
    total_score_percentage double precision DEFAULT 0 NOT NULL,
    items_evaluated_count integer DEFAULT 0 NOT NULL,
    items_not_evaluated_count integer DEFAULT 0 NOT NULL,
    notes text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_snapshots_items_evaluated CHECK ((items_evaluated_count >= 0)),
    CONSTRAINT chk_snapshots_items_not_evaluated CHECK ((items_not_evaluated_count >= 0)),
    CONSTRAINT chk_snapshots_total_actual_points CHECK ((total_actual_points >= (0)::double precision)),
    CONSTRAINT chk_snapshots_total_possible_points CHECK ((total_possible_points >= (0)::double precision)),
    CONSTRAINT chk_snapshots_total_score_percentage CHECK (((total_score_percentage >= (0)::double precision) AND (total_score_percentage <= (100)::double precision)))
);


--
-- Name: TABLE patient_score_snapshots; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.patient_score_snapshots IS 'Snapshots completos de escores de saúde do paciente em pontos específicos no tempo';


--
-- Name: COLUMN patient_score_snapshots.calculated_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_snapshots.calculated_at IS 'Data e hora do cálculo do snapshot';


--
-- Name: COLUMN patient_score_snapshots.total_score_percentage; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_snapshots.total_score_percentage IS 'Score final em percentual (0-100)';


--
-- Name: COLUMN patient_score_snapshots.items_evaluated_count; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_snapshots.items_evaluated_count IS 'Quantidade de itens que foram avaliados';


--
-- Name: COLUMN patient_score_snapshots.items_not_evaluated_count; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patient_score_snapshots.items_not_evaluated_count IS 'Quantidade de itens que não foram avaliados';


--
-- Name: patient_subscriptions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patient_subscriptions (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    subscription_plan_id uuid NOT NULL,
    plan_snapshot jsonb NOT NULL,
    status character varying(20) NOT NULL,
    auto_renew boolean DEFAULT true NOT NULL,
    start_date date NOT NULL,
    end_date date,
    trial_end_date date,
    next_billing_date date,
    cancelled_at timestamp with time zone,
    discount_percent numeric(5,2) DEFAULT 0 NOT NULL,
    discount_reason text,
    custom_price numeric(10,2),
    custom_trial_days integer,
    cancellation_reason text,
    notes text,
    renewal_count integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_patient_subscriptions_custom_price CHECK ((custom_price >= (0)::numeric)),
    CONSTRAINT chk_patient_subscriptions_custom_trial_days CHECK ((custom_trial_days >= 0)),
    CONSTRAINT chk_patient_subscriptions_discount_percent CHECK (((discount_percent >= (0)::numeric) AND (discount_percent <= (100)::numeric))),
    CONSTRAINT chk_patient_subscriptions_status CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'inactive'::character varying, 'cancelled'::character varying, 'expired'::character varying, 'suspended'::character varying, 'trial'::character varying])::text[]))),
    CONSTRAINT patient_subscriptions_custom_price_check CHECK ((custom_price >= (0)::numeric)),
    CONSTRAINT patient_subscriptions_custom_trial_days_check CHECK ((custom_trial_days >= 0)),
    CONSTRAINT patient_subscriptions_discount_percent_check CHECK (((discount_percent >= (0)::numeric) AND (discount_percent <= (100)::numeric))),
    CONSTRAINT patient_subscriptions_status_check CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'inactive'::character varying, 'cancelled'::character varying, 'expired'::character varying, 'suspended'::character varying, 'trial'::character varying])::text[])))
);


--
-- Name: patients; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patients (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    user_id uuid NOT NULL,
    name character varying(200) NOT NULL,
    cpf text,
    birth_date date NOT NULL,
    gender character varying(10) NOT NULL,
    phone character varying(20),
    address character varying(500),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    mother_name character varying(200),
    father_name character varying(200),
    height numeric(5,2),
    weight numeric(5,2),
    municipality character varying(100),
    state character varying(2),
    social_gender character varying(20),
    age bigint DEFAULT 0 NOT NULL,
    age_text character varying(20) DEFAULT ''::character varying NOT NULL,
    rg text,
    email character varying(200),
    blood_type character varying(3),
    marital_status character varying(10),
    occupation character varying(200),
    emergency_contact character varying(200),
    emergency_phone character varying(20),
    menopause boolean,
    source character varying(40) DEFAULT 'direct'::character varying NOT NULL,
    cpf_blind_index character varying(64),
    CONSTRAINT chk_patients_blood_type CHECK (((blood_type)::text = ANY ((ARRAY['A+'::character varying, 'A-'::character varying, 'B+'::character varying, 'B-'::character varying, 'AB+'::character varying, 'AB-'::character varying, 'O+'::character varying, 'O-'::character varying])::text[]))),
    CONSTRAINT chk_patients_gender CHECK (((gender)::text = ANY ((ARRAY['male'::character varying, 'female'::character varying, 'other'::character varying])::text[]))),
    CONSTRAINT chk_patients_marital_status CHECK (((marital_status)::text = ANY ((ARRAY['single'::character varying, 'married'::character varying, 'divorced'::character varying, 'widowed'::character varying, 'other'::character varying])::text[]))),
    CONSTRAINT chk_patients_social_gender CHECK (((social_gender)::text = ANY ((ARRAY['male'::character varying, 'female'::character varying, 'non_binary'::character varying, 'trans_male'::character varying, 'trans_female'::character varying, 'other'::character varying, 'prefer_not_to_say'::character varying])::text[])))
);


--
-- Name: COLUMN patients.municipality; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patients.municipality IS 'City/municipality where the patient resides';


--
-- Name: COLUMN patients.state; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patients.state IS 'State/UF (2-letter code) where the patient resides';


--
-- Name: COLUMN patients.menopause; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.patients.menopause IS 'Indica se a paciente está na menopausa (apenas para gender=female)';


--
-- Name: payment_receipt_counters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.payment_receipt_counters (
    year bigint NOT NULL,
    last_seq bigint DEFAULT 0 NOT NULL
);


--
-- Name: payment_receipt_counters_year_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.payment_receipt_counters_year_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: payment_receipt_counters_year_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.payment_receipt_counters_year_seq OWNED BY public.payment_receipt_counters.year;


--
-- Name: physical_assessments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.physical_assessments (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    created_by_id uuid NOT NULL,
    assessment_date date NOT NULL,
    acsm_risk_level character varying(10),
    acsm_risk_factors_count bigint DEFAULT 0 NOT NULL,
    acsm_risk_factors jsonb,
    acsm_protective_factors jsonb,
    acsm_recommendation text,
    acsm_tags jsonb,
    front_photo_url text,
    side_photo_url text,
    ai_recommendation text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight numeric(6,2),
    height numeric(5,1),
    waist_circumference numeric(5,1),
    bmi numeric(5,2),
    bri numeric(5,2),
    body_fat_percent numeric(5,2),
    lean_mass numeric(6,2),
    systolic_bp bigint,
    diastolic_bp bigint,
    resting_heart_rate bigint,
    ldl numeric(6,2),
    hdl numeric(6,2),
    total_cholesterol numeric(6,2),
    triglycerides numeric(6,2),
    fasting_glucose numeric(6,2),
    hb_a1c numeric(5,2),
    family_history boolean,
    smoking_status character varying(10),
    physical_activity_level character varying(15),
    cardiovascular_disease boolean,
    diabetes_type character varying(20),
    symptoms text,
    clinical_alert boolean,
    acsm_tags_structured jsonb,
    html_content text,
    CONSTRAINT chk_physical_assessments_acsm_risk_level CHECK (((acsm_risk_level)::text = ANY ((ARRAY['low'::character varying, 'moderate'::character varying, 'high'::character varying])::text[]))),
    CONSTRAINT chk_physical_assessments_physical_activity_level CHECK (((physical_activity_level)::text = ANY ((ARRAY['sedentary'::character varying, 'insufficient'::character varying, 'moderate'::character varying, 'active'::character varying])::text[]))),
    CONSTRAINT chk_physical_assessments_smoking_status CHECK (((smoking_status)::text = ANY ((ARRAY['never'::character varying, 'former'::character varying, 'current'::character varying])::text[])))
);


--
-- Name: postural_assessments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.postural_assessments (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    created_by_id uuid NOT NULL,
    assessment_date date NOT NULL,
    physical_assessment_id uuid,
    view_type character varying(20) DEFAULT 'front'::character varying NOT NULL,
    shoulder_deviation numeric(5,2),
    hip_deviation numeric(5,2),
    head_lateral_deviation numeric(5,2),
    fhp numeric(5,2),
    thoracic_kyphosis numeric(5,2),
    lumbar_lordosis numeric(5,2),
    knee_angle numeric(5,2),
    photo_url text,
    postural_score bigint DEFAULT 0 NOT NULL,
    postural_classification character varying(30) DEFAULT ''::character varying NOT NULL,
    severe_deviations bigint DEFAULT 0 NOT NULL,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: prescription_medications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prescription_medications (
    id uuid NOT NULL,
    prescription_id uuid NOT NULL,
    medication_definition_id uuid,
    medication_name character varying(200) NOT NULL,
    active_ingredient character varying(200) NOT NULL,
    category character varying(20) DEFAULT 'simple'::character varying NOT NULL,
    concentration character varying(100) NOT NULL,
    dosage character varying(100) NOT NULL,
    frequency character varying(100) NOT NULL,
    route character varying(50) NOT NULL,
    duration integer NOT NULL,
    quantity integer NOT NULL,
    quantity_in_words character varying(200) NOT NULL,
    instructions text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: prescriptions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prescriptions (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    doctor_id uuid NOT NULL,
    medication_definition_id uuid,
    medication_name character varying(200) NOT NULL,
    active_ingredient character varying(200) NOT NULL,
    category character varying(20) DEFAULT 'simple'::character varying NOT NULL,
    concentration character varying(100) NOT NULL,
    dosage character varying(100) NOT NULL,
    frequency character varying(100) NOT NULL,
    route character varying(50) NOT NULL,
    duration integer NOT NULL,
    quantity integer NOT NULL,
    quantity_in_words character varying(200) NOT NULL,
    instructions text,
    prescription_date timestamp with time zone NOT NULL,
    valid_until date NOT NULL,
    sncr_number character varying(50),
    sncr_status character varying(20),
    signed_pdf_path character varying(500),
    signed_pdf_hash character varying(64),
    qr_code_data text,
    signed_at timestamp with time zone,
    certificate_serial character varying(100),
    status character varying(20) DEFAULT 'active'::character varying NOT NULL,
    is_used boolean DEFAULT false NOT NULL,
    dispensed_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    general_instructions text,
    CONSTRAINT chk_prescriptions_category CHECK (((category)::text = ANY ((ARRAY['simple'::character varying, 'c1'::character varying, 'c5'::character varying, 'antibiotic'::character varying, 'glp1'::character varying])::text[]))),
    CONSTRAINT chk_prescriptions_status CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'completed'::character varying, 'cancelled'::character varying, 'expired'::character varying])::text[]))),
    CONSTRAINT prescriptions_category_check CHECK (((category)::text = ANY ((ARRAY['simple'::character varying, 'c1'::character varying, 'c5'::character varying, 'antibiotic'::character varying, 'glp1'::character varying])::text[]))),
    CONSTRAINT prescriptions_status_check CHECK (((status)::text = ANY ((ARRAY['active'::character varying, 'completed'::character varying, 'cancelled'::character varying, 'expired'::character varying])::text[])))
);


--
-- Name: TABLE prescriptions; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.prescriptions IS 'Prescrições médicas digitais com assinatura ICP-Brasil';


--
-- Name: COLUMN prescriptions.sncr_number; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.prescriptions.sncr_number IS 'Número SNCR gerado pela ANVISA (ou stub)';


--
-- Name: COLUMN prescriptions.signed_pdf_hash; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.prescriptions.signed_pdf_hash IS 'Hash SHA-256 do PDF assinado (integridade)';


--
-- Name: COLUMN prescriptions.qr_code_data; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.prescriptions.qr_code_data IS 'Dados do QR Code para validação pública';


--
-- Name: COLUMN prescriptions.certificate_serial; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.prescriptions.certificate_serial IS 'Número de série do certificado ICP-Brasil usado na assinatura';


--
-- Name: processing_jobs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.processing_jobs (
    id uuid NOT NULL,
    lab_result_batch_id uuid NOT NULL,
    type character varying(50) NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying NOT NULL,
    pdf_path text NOT NULL,
    extracted_text text,
    ai_response text,
    error_message text,
    attempts integer DEFAULT 0 NOT NULL,
    max_attempts integer DEFAULT 3 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    started_at timestamp with time zone,
    completed_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    progress_step integer,
    progress_message text,
    CONSTRAINT chk_processing_jobs_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'processing'::character varying, 'completed'::character varying, 'failed'::character varying])::text[]))),
    CONSTRAINT chk_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'processing'::character varying, 'completed'::character varying, 'failed'::character varying])::text[])))
);


--
-- Name: COLUMN processing_jobs.progress_step; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.processing_jobs.progress_step IS 'Etapa atual do processamento (1-6)';


--
-- Name: COLUMN processing_jobs.progress_message; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.processing_jobs.progress_message IS 'Mensagem descritiva da etapa atual';


--
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.refresh_tokens (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    token_hash character varying(64) NOT NULL,
    type character varying(20) DEFAULT 'refresh'::character varying NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    revoked_at timestamp with time zone,
    used_at timestamp with time zone,
    user_agent text,
    ip_address character varying(45),
    created_at timestamp with time zone NOT NULL
);


--
-- Name: score_groups; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_groups (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    name character varying(200) NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


--
-- Name: TABLE score_groups; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_groups IS 'Grupos principais de escores clínicos (ex: Hemograma, Lipídeos)';


--
-- Name: COLUMN score_groups.name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_groups.name IS 'Nome do grupo de escores';


--
-- Name: COLUMN score_groups."order"; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_groups."order" IS 'Ordem de exibição do grupo';


--
-- Name: score_item_embeddings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_item_embeddings (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    score_item_id uuid NOT NULL,
    embedding public.vector(1024),
    text_source text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    is_stale boolean DEFAULT false NOT NULL
);


--
-- Name: TABLE score_item_embeddings; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_item_embeddings IS 'Stores vector embeddings for ScoreItems. Used to recommend articles for clinical parameters and discover which parameters an article covers.';


--
-- Name: COLUMN score_item_embeddings.is_stale; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_embeddings.is_stale IS 'True when source data changed and embedding needs regeneration';


--
-- Name: score_item_enrichment_preparation; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_item_enrichment_preparation (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    score_item_id uuid NOT NULL,
    selected_chunks jsonb DEFAULT '[]'::jsonb NOT NULL,
    metadata jsonb DEFAULT '{}'::jsonb,
    status character varying(20) DEFAULT 'ready'::character varying NOT NULL,
    error_message text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    prompt_clinical_relevance text,
    prompt_patient_explanation text,
    prompt_conduct text,
    prompt_max_points text,
    CONSTRAINT score_item_enrichment_preparation_status_check CHECK (((status)::text = ANY ((ARRAY['ready'::character varying, 'processing'::character varying, 'completed'::character varying, 'failed'::character varying, 'stale'::character varying])::text[])))
);


--
-- Name: TABLE score_item_enrichment_preparation; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_item_enrichment_preparation IS 'Preparação de chunks científicos para enrichment via Claude (Etapa 1 - sem IA)';


--
-- Name: COLUMN score_item_enrichment_preparation.selected_chunks; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.selected_chunks IS 'Array JSONB com 15-20 chunks dos artigos mais relevantes via RAG';


--
-- Name: COLUMN score_item_enrichment_preparation.metadata; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.metadata IS 'Estatísticas da seleção: total_chunks, articles_count, avg_similarity, sections_distribution';


--
-- Name: COLUMN score_item_enrichment_preparation.status; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.status IS 'Status: ready (awaiting enrichment), processing (being enriched), completed (enriched), failed (enrichment failed), stale (needs re-preparation)';


--
-- Name: COLUMN score_item_enrichment_preparation.prompt_clinical_relevance; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.prompt_clinical_relevance IS 'Prompt pronto para Claude gerar clinical_relevance (1200-1800 chars)';


--
-- Name: COLUMN score_item_enrichment_preparation.prompt_patient_explanation; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.prompt_patient_explanation IS 'Prompt pronto para Claude gerar patient_explanation (600-900 chars)';


--
-- Name: COLUMN score_item_enrichment_preparation.prompt_conduct; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.prompt_conduct IS 'Prompt pronto para Claude gerar conduct (1000-1500 chars)';


--
-- Name: COLUMN score_item_enrichment_preparation.prompt_max_points; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_item_enrichment_preparation.prompt_max_points IS 'Prompt pronto para Claude gerar max_points (0-50)';


--
-- Name: score_item_method_pillars; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_item_method_pillars (
    score_item_id uuid NOT NULL,
    method_pillar_id uuid NOT NULL
);


--
-- Name: score_item_review_history; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_item_review_history (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    score_item_id uuid NOT NULL,
    review_type character varying(50) DEFAULT 'llm_enrichment'::character varying NOT NULL,
    before_snapshot jsonb,
    after_snapshot jsonb,
    tier character varying(20),
    confidence_score double precision,
    model_used character varying(100),
    reviewed_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


--
-- Name: score_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_items (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    name character varying(300) NOT NULL,
    unit character varying(50),
    unit_conversion text,
    points double precision,
    "order" integer DEFAULT 0 NOT NULL,
    subgroup_id uuid NOT NULL,
    parent_item_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    lab_test_code character varying(100),
    clinical_relevance text,
    patient_explanation text,
    conduct text,
    last_review timestamp with time zone,
    gender character varying(20) DEFAULT 'not_applicable'::character varying,
    age_range_min integer,
    age_range_max integer,
    post_menopause boolean,
    anamnese_item_code character varying(200),
    is_light_version boolean DEFAULT false NOT NULL,
    light_order integer,
    light_question text,
    CONSTRAINT chk_score_items_age_range_max CHECK (((age_range_max >= 0) AND (age_range_max <= 150))),
    CONSTRAINT chk_score_items_age_range_min CHECK (((age_range_min >= 0) AND (age_range_min <= 150))),
    CONSTRAINT chk_score_items_gender CHECK (((gender)::text = ANY (ARRAY[('not_applicable'::character varying)::text, ('male'::character varying)::text, ('female'::character varying)::text]))),
    CONSTRAINT chk_score_items_points CHECK (((points >= (0)::double precision) AND (points <= (100)::double precision))),
    CONSTRAINT score_items_age_range_max_check CHECK (((age_range_max >= 0) AND (age_range_max <= 150))),
    CONSTRAINT score_items_age_range_min_check CHECK (((age_range_min >= 0) AND (age_range_min <= 150))),
    CONSTRAINT score_items_gender_check CHECK (((gender)::text = ANY (ARRAY[('not_applicable'::character varying)::text, ('male'::character varying)::text, ('female'::character varying)::text])))
);


--
-- Name: TABLE score_items; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_items IS 'Itens individuais de escores (ex: Hemoglobina - Homens, FEVE)';


--
-- Name: COLUMN score_items.name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.name IS 'Nome do parâmetro clínico';


--
-- Name: COLUMN score_items.unit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.unit IS 'Unidade de medida (ex: g/dL, %, mm)';


--
-- Name: COLUMN score_items.unit_conversion; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.unit_conversion IS 'Fórmula de conversão de unidades (ex: 1 g/dL = 10 g/L)';


--
-- Name: COLUMN score_items.points; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.points IS 'Pontos máximos para este item no escore (0-100)';


--
-- Name: COLUMN score_items."order"; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items."order" IS 'Ordem de exibição do item dentro do subgrupo';


--
-- Name: COLUMN score_items.subgroup_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.subgroup_id IS 'Referência ao subgrupo pai';


--
-- Name: COLUMN score_items.parent_item_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.parent_item_id IS 'Referência a item pai para hierarquia (opcional)';


--
-- Name: COLUMN score_items.gender; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.gender IS 'Gênero aplicável: not_applicable (padrão), male, female';


--
-- Name: COLUMN score_items.age_range_min; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.age_range_min IS 'Idade mínima aplicável em anos (0-150)';


--
-- Name: COLUMN score_items.age_range_max; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.age_range_max IS 'Idade máxima aplicável em anos (0-150)';


--
-- Name: COLUMN score_items.post_menopause; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_items.post_menopause IS 'Indica se o score_item é aplicável apenas para mulheres pós-menopausa';


--
-- Name: score_levels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_levels (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    level integer NOT NULL,
    name character varying(500) NOT NULL,
    lower_limit character varying(50),
    upper_limit character varying(50),
    item_id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    operator character varying(10) DEFAULT 'between'::character varying NOT NULL,
    clinical_relevance text,
    patient_explanation text,
    conduct text,
    last_review timestamp with time zone,
    CONSTRAINT chk_score_levels_level CHECK (((level >= 0) AND (level <= 6))),
    CONSTRAINT chk_score_levels_operator CHECK (((operator)::text = ANY (ARRAY[('='::character varying)::text, ('>'::character varying)::text, ('>='::character varying)::text, ('<'::character varying)::text, ('<='::character varying)::text, ('between'::character varying)::text])))
);


--
-- Name: TABLE score_levels; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_levels IS 'Níveis de estratificação de risco para cada item (0-6)';


--
-- Name: COLUMN score_levels.level; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.level IS 'Nível de risco (0=crítico, 1-4=intermediários, 5=ótimo, 6=reservado)';


--
-- Name: COLUMN score_levels.name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.name IS 'Descrição do nível (ex: 55 a 70 (Ótimo))';


--
-- Name: COLUMN score_levels.lower_limit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.lower_limit IS 'Limite inferior do intervalo';


--
-- Name: COLUMN score_levels.upper_limit; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.upper_limit IS 'Limite superior do intervalo';


--
-- Name: COLUMN score_levels.item_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.item_id IS 'Referência ao item de escore';


--
-- Name: COLUMN score_levels.operator; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_levels.operator IS 'Comparison operator: = (equals), > (greater than), >= (greater or equal), < (less than), <= (less or equal), between (range)';


--
-- Name: score_subgroups; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.score_subgroups (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    name character varying(200) NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    group_id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    max_select integer DEFAULT 0 NOT NULL
);


--
-- Name: TABLE score_subgroups; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.score_subgroups IS 'Subgrupos dentro de um grupo de escores (ex: Série Vermelha, Série Branca)';


--
-- Name: COLUMN score_subgroups.name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_subgroups.name IS 'Nome do subgrupo';


--
-- Name: COLUMN score_subgroups."order"; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_subgroups."order" IS 'Ordem de exibição do subgrupo dentro do grupo';


--
-- Name: COLUMN score_subgroups.group_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.score_subgroups.group_id IS 'Referência ao grupo pai';


--
-- Name: subscription_plans; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.subscription_plans (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    features jsonb,
    price numeric(10,2) NOT NULL,
    currency character varying(3) DEFAULT 'BRL'::character varying NOT NULL,
    billing_cycle character varying(20) NOT NULL,
    method_id uuid,
    is_active boolean DEFAULT true NOT NULL,
    trial_period_days integer DEFAULT 0 NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_subscription_plans_billing_cycle CHECK (((billing_cycle)::text = ANY ((ARRAY['monthly'::character varying, 'quarterly'::character varying, 'yearly'::character varying, 'one_time'::character varying])::text[]))),
    CONSTRAINT chk_subscription_plans_price CHECK ((price >= (0)::numeric)),
    CONSTRAINT chk_subscription_plans_trial_period_days CHECK ((trial_period_days >= 0)),
    CONSTRAINT subscription_plans_billing_cycle_check CHECK (((billing_cycle)::text = ANY ((ARRAY['monthly'::character varying, 'quarterly'::character varying, 'yearly'::character varying, 'one_time'::character varying])::text[]))),
    CONSTRAINT subscription_plans_price_check CHECK ((price >= (0)::numeric)),
    CONSTRAINT subscription_plans_trial_period_days_check CHECK ((trial_period_days >= 0))
);


--
-- Name: telemed_lobby_tokens; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.telemed_lobby_tokens (
    id uuid NOT NULL,
    appointment_id uuid NOT NULL,
    token character varying(128) NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    used_at timestamp with time zone,
    created_at timestamp with time zone
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v7() NOT NULL,
    email character varying(255) NOT NULL,
    password_hash text,
    two_factor_enabled boolean DEFAULT false,
    two_factor_secret text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    selected_patient_id uuid,
    preferences jsonb DEFAULT '{}'::jsonb,
    name character varying(200) NOT NULL,
    roles jsonb DEFAULT '["patient"]'::jsonb NOT NULL,
    crm character varying(20),
    crmuf character varying(2),
    rqe character varying(20),
    specialty character varying(100),
    professional_address text,
    professional_phone character varying(20),
    certificate_pfx text,
    certificate_password text,
    certificate_expiry date,
    certificate_serial character varying(100),
    oauth_provider character varying(20),
    oauth_provider_id character varying(255),
    oauth_picture_url text,
    certificate_cpf character varying(14),
    certificate_name character varying(200),
    cpf text,
    certificate_active boolean DEFAULT false,
    o_auth_provider character varying(20),
    o_auth_provider_id character varying(255),
    o_auth_picture_url text,
    lgpd_consented_at timestamp with time zone,
    gender character varying(10),
    treatment character varying(10),
    CONSTRAINT chk_oauth_provider CHECK (((oauth_provider IS NULL) OR ((oauth_provider)::text = ANY ((ARRAY['google'::character varying, 'apple'::character varying])::text[]))))
);


--
-- Name: COLUMN users.selected_patient_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.users.selected_patient_id IS 'ID of the currently selected patient for this user context';


--
-- Name: COLUMN users.preferences; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.users.preferences IS 'Preferências do usuário em formato JSON (viewport do mindmap, configurações de interface, etc.)';


--
-- Name: COLUMN users.crm; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.users.crm IS 'Número do CRM (Conselho Regional de Medicina)';


--
-- Name: COLUMN users.certificate_pfx; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.users.certificate_pfx IS 'Certificado digital A1 (criptografado em Base64)';


--
-- Name: COLUMN users.certificate_password; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.users.certificate_password IS 'Senha do certificado A1 (criptografada)';


--
-- Name: waitlist_entries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.waitlist_entries (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    doctor_id uuid,
    preferred_type character varying(20),
    notes text,
    status character varying(20) DEFAULT 'waiting'::character varying NOT NULL,
    scheduled_appointment_id uuid,
    created_by_user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_waitlist_entries_status CHECK (((status)::text = ANY ((ARRAY['waiting'::character varying, 'scheduled'::character varying, 'cancelled'::character varying])::text[])))
);


--
-- Name: working_hours; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.working_hours (
    id uuid NOT NULL,
    doctor_id uuid NOT NULL,
    weekday bigint NOT NULL,
    start_minute bigint NOT NULL,
    end_minute bigint NOT NULL,
    slot_duration bigint DEFAULT 30 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_working_hours_end_minute CHECK (((end_minute > 0) AND (end_minute <= 1440))),
    CONSTRAINT chk_working_hours_slot_duration CHECK (((slot_duration > 0) AND (slot_duration <= 480))),
    CONSTRAINT chk_working_hours_start_minute CHECK (((start_minute >= 0) AND (start_minute < 1440))),
    CONSTRAINT chk_working_hours_weekday CHECK (((weekday >= 0) AND (weekday <= 6)))
);


--
-- Name: workout_mesocycles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_mesocycles (
    id uuid NOT NULL,
    periodization_id uuid NOT NULL,
    "order" bigint NOT NULL,
    phase character varying(20) NOT NULL,
    duration_weeks bigint NOT NULL,
    volume_percent bigint NOT NULL,
    intensity_percent bigint NOT NULL,
    physiological_focus character varying(100) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_workout_mesocycles_phase CHECK (((phase)::text = ANY ((ARRAY['accumulation'::character varying, 'transformation'::character varying, 'realization'::character varying, 'hypertrophy'::character varying, 'strength'::character varying, 'endurance'::character varying, 'power'::character varying])::text[])))
);


--
-- Name: workout_periodizations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_periodizations (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    created_by_id uuid NOT NULL,
    framework character varying(20) NOT NULL,
    start_date date NOT NULL,
    total_weeks bigint NOT NULL,
    objective character varying(200) NOT NULL,
    scientific_justification text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_workout_periodizations_framework CHECK (((framework)::text = ANY ((ARRAY['bompa'::character varying, 'linear'::character varying, 'undulating'::character varying, 'block'::character varying])::text[])))
);


--
-- Name: workout_plan_sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_plan_sessions (
    id uuid NOT NULL,
    plan_id uuid NOT NULL,
    name character varying(100) NOT NULL,
    "order" bigint DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: workout_plans; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_plans (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    created_by_id uuid NOT NULL,
    name character varying(200) NOT NULL,
    objective character varying(30) NOT NULL,
    intensity character varying(20) NOT NULL,
    duration_minutes bigint DEFAULT 60 NOT NULL,
    weekly_frequency bigint DEFAULT 3 NOT NULL,
    public_code character varying(8) NOT NULL,
    html_content text,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_workout_plans_intensity CHECK (((intensity)::text = ANY ((ARRAY['very_light'::character varying, 'light'::character varying, 'moderate'::character varying, 'high'::character varying, 'very_high'::character varying])::text[]))),
    CONSTRAINT chk_workout_plans_objective CHECK (((objective)::text = ANY ((ARRAY['hypertrophy'::character varying, 'strength'::character varying, 'endurance'::character varying, 'weight_loss'::character varying, 'conditioning'::character varying, 'rehabilitation'::character varying])::text[])))
);


--
-- Name: workout_session_exercise_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_session_exercise_logs (
    id uuid NOT NULL,
    session_id uuid NOT NULL,
    plan_exercise_id uuid NOT NULL,
    exercise_id uuid NOT NULL,
    set_number bigint NOT NULL,
    reps bigint,
    weight numeric(6,2),
    duration_sec bigint,
    rpe bigint,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_workout_session_exercise_logs_rpe CHECK (((rpe IS NULL) OR ((rpe >= 1) AND (rpe <= 10)))),
    CONSTRAINT chk_workout_session_exercise_logs_set_number CHECK (((set_number >= 1) AND (set_number <= 20)))
);


--
-- Name: workout_session_exercises; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_session_exercises (
    id uuid NOT NULL,
    session_id uuid NOT NULL,
    exercise_id uuid NOT NULL,
    phase character varying(10) NOT NULL,
    "order" bigint DEFAULT 0 NOT NULL,
    sets bigint DEFAULT 3 NOT NULL,
    reps character varying(20) DEFAULT '10'::character varying NOT NULL,
    cadence character varying(15) DEFAULT 'normal'::character varying NOT NULL,
    rest_between_sets_sec bigint DEFAULT 60 NOT NULL,
    rest_between_exercises_sec bigint DEFAULT 90 NOT NULL,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_workout_session_exercises_cadence CHECK (((cadence)::text = ANY ((ARRAY['normal'::character varying, 'slow'::character varying, 'paused'::character varying, 'explosive'::character varying, 'free'::character varying])::text[]))),
    CONSTRAINT chk_workout_session_exercises_phase CHECK (((phase)::text = ANY ((ARRAY['warmup'::character varying, 'main'::character varying, 'cooldown'::character varying])::text[])))
);


--
-- Name: workout_sessions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.workout_sessions (
    id uuid NOT NULL,
    patient_id uuid NOT NULL,
    plan_id uuid NOT NULL,
    plan_session_id uuid NOT NULL,
    scheduled_date date NOT NULL,
    completed_at timestamp with time zone,
    notes text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: payment_receipt_counters year; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment_receipt_counters ALTER COLUMN year SET DEFAULT nextval('public.payment_receipt_counters_year_seq'::regclass);


--
-- Name: anamnesis_items anamnesis_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_items
    ADD CONSTRAINT anamnesis_items_pkey PRIMARY KEY (id);


--
-- Name: anamnesis anamnesis_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_pkey PRIMARY KEY (id);


--
-- Name: anamnesis_template_items anamnesis_template_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_template_items
    ADD CONSTRAINT anamnesis_template_items_pkey PRIMARY KEY (id);


--
-- Name: anamnesis_templates anamnesis_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_templates
    ADD CONSTRAINT anamnesis_templates_pkey PRIMARY KEY (id);


--
-- Name: anonymous_score_group_results anonymous_score_group_results_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_group_results
    ADD CONSTRAINT anonymous_score_group_results_pkey PRIMARY KEY (id);


--
-- Name: anonymous_score_items anonymous_score_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_items
    ADD CONSTRAINT anonymous_score_items_pkey PRIMARY KEY (id);


--
-- Name: anonymous_score_sessions anonymous_score_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_sessions
    ADD CONSTRAINT anonymous_score_sessions_pkey PRIMARY KEY (id);


--
-- Name: anonymous_score_snapshots anonymous_score_snapshots_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_snapshots
    ADD CONSTRAINT anonymous_score_snapshots_pkey PRIMARY KEY (id);


--
-- Name: api_usage_logs api_usage_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.api_usage_logs
    ADD CONSTRAINT api_usage_logs_pkey PRIMARY KEY (id);


--
-- Name: appointment_payments appointment_payments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointment_payments
    ADD CONSTRAINT appointment_payments_pkey PRIMARY KEY (id);


--
-- Name: appointment_resources appointment_resources_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointment_resources
    ADD CONSTRAINT appointment_resources_pkey PRIMARY KEY (id);


--
-- Name: appointments appointments_no_overlap; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointments_no_overlap EXCLUDE USING gist (doctor_id WITH =, tstzrange(scheduled_at, end_at) WITH &&) WHERE ((((status)::text <> ALL ((ARRAY['cancelled'::character varying, 'no_show'::character varying])::text[])) AND (deleted_at IS NULL)));


--
-- Name: appointments appointments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointments_pkey PRIMARY KEY (id);


--
-- Name: article_embeddings article_embeddings_article_id_chunk_index_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_embeddings
    ADD CONSTRAINT article_embeddings_article_id_chunk_index_key UNIQUE (article_id, chunk_index);


--
-- Name: article_embeddings article_embeddings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_embeddings
    ADD CONSTRAINT article_embeddings_pkey PRIMARY KEY (id);


--
-- Name: article_score_items article_score_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_score_items
    ADD CONSTRAINT article_score_items_pkey PRIMARY KEY (score_item_id, article_id);


--
-- Name: articles articles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);


--
-- Name: audit_logs audit_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT audit_logs_pkey PRIMARY KEY (id);


--
-- Name: auto_link_batch_checkpoints auto_link_batch_checkpoints_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_batch_checkpoints
    ADD CONSTRAINT auto_link_batch_checkpoints_pkey PRIMARY KEY (id);


--
-- Name: auto_link_batch_checkpoints auto_link_batch_checkpoints_run_id_batch_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_batch_checkpoints
    ADD CONSTRAINT auto_link_batch_checkpoints_run_id_batch_number_key UNIQUE (run_id, batch_number);


--
-- Name: auto_link_item_log auto_link_item_log_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_item_log
    ADD CONSTRAINT auto_link_item_log_pkey PRIMARY KEY (id);


--
-- Name: auto_link_processing_state auto_link_processing_state_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_processing_state
    ADD CONSTRAINT auto_link_processing_state_pkey PRIMARY KEY (id);


--
-- Name: auto_link_processing_state auto_link_processing_state_run_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_processing_state
    ADD CONSTRAINT auto_link_processing_state_run_id_key UNIQUE (run_id);


--
-- Name: calendar_credentials calendar_credentials_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.calendar_credentials
    ADD CONSTRAINT calendar_credentials_pkey PRIMARY KEY (id);


--
-- Name: campaigns campaigns_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);


--
-- Name: consultation_prices consultation_prices_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.consultation_prices
    ADD CONSTRAINT consultation_prices_pkey PRIMARY KEY (id);


--
-- Name: continuum_box_templates continuum_box_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.continuum_box_templates
    ADD CONSTRAINT continuum_box_templates_pkey PRIMARY KEY (id);


--
-- Name: continuum_template_items continuum_template_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.continuum_template_items
    ADD CONSTRAINT continuum_template_items_pkey PRIMARY KEY (id);


--
-- Name: continuum_templates continuum_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.continuum_templates
    ADD CONSTRAINT continuum_templates_pkey PRIMARY KEY (id);


--
-- Name: conversation_reads conversation_reads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.conversation_reads
    ADD CONSTRAINT conversation_reads_pkey PRIMARY KEY (id);


--
-- Name: device_tokens device_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.device_tokens
    ADD CONSTRAINT device_tokens_pkey PRIMARY KEY (id);


--
-- Name: doctor_absences doctor_absences_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.doctor_absences
    ADD CONSTRAINT doctor_absences_pkey PRIMARY KEY (id);


--
-- Name: email_ingest_states email_ingest_states_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.email_ingest_states
    ADD CONSTRAINT email_ingest_states_pkey PRIMARY KEY (id);


--
-- Name: embedding_audit_log embedding_audit_log_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.embedding_audit_log
    ADD CONSTRAINT embedding_audit_log_pkey PRIMARY KEY (id);


--
-- Name: embedding_queue embedding_queue_entity_type_entity_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.embedding_queue
    ADD CONSTRAINT embedding_queue_entity_type_entity_id_key UNIQUE (entity_type, entity_id);


--
-- Name: embedding_queue embedding_queue_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.embedding_queue
    ADD CONSTRAINT embedding_queue_pkey PRIMARY KEY (id);


--
-- Name: exercises exercises_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.exercises
    ADD CONSTRAINT exercises_pkey PRIMARY KEY (id);


--
-- Name: fitness_test_results fitness_test_results_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fitness_test_results
    ADD CONSTRAINT fitness_test_results_pkey PRIMARY KEY (id);


--
-- Name: health_check_ins health_check_ins_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.health_check_ins
    ADD CONSTRAINT health_check_ins_pkey PRIMARY KEY (id);


--
-- Name: integrated_plan_revisions integrated_plan_revisions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.integrated_plan_revisions
    ADD CONSTRAINT integrated_plan_revisions_pkey PRIMARY KEY (id);


--
-- Name: lab_request_template_tests lab_request_template_tests_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_template_tests
    ADD CONSTRAINT lab_request_template_tests_pkey PRIMARY KEY (lab_request_template_id, lab_test_definition_id);


--
-- Name: lab_request_templates lab_request_templates_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_templates
    ADD CONSTRAINT lab_request_templates_name_key UNIQUE (name);


--
-- Name: lab_request_templates lab_request_templates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_templates
    ADD CONSTRAINT lab_request_templates_pkey PRIMARY KEY (id);


--
-- Name: lab_requests lab_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT lab_requests_pkey PRIMARY KEY (id);


--
-- Name: lab_result_batches lab_result_batches_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT lab_result_batches_pkey PRIMARY KEY (id);


--
-- Name: lab_result_values lab_result_values_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_values
    ADD CONSTRAINT lab_result_values_pkey PRIMARY KEY (id);


--
-- Name: lab_result_view_items lab_result_view_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_view_items
    ADD CONSTRAINT lab_result_view_items_pkey PRIMARY KEY (id);


--
-- Name: lab_result_views lab_result_views_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_views
    ADD CONSTRAINT lab_result_views_pkey PRIMARY KEY (id);


--
-- Name: lab_results lab_results_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_results
    ADD CONSTRAINT lab_results_pkey PRIMARY KEY (id);


--
-- Name: lab_test_definitions lab_test_definitions_code_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_definitions
    ADD CONSTRAINT lab_test_definitions_code_key UNIQUE (code);


--
-- Name: lab_test_definitions lab_test_definitions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_definitions
    ADD CONSTRAINT lab_test_definitions_pkey PRIMARY KEY (id);


--
-- Name: lab_test_unit_conversions lab_test_unit_conversions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_unit_conversions
    ADD CONSTRAINT lab_test_unit_conversions_pkey PRIMARY KEY (id);


--
-- Name: lead_activities lead_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lead_activities
    ADD CONSTRAINT lead_activities_pkey PRIMARY KEY (id);


--
-- Name: leads leads_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.leads
    ADD CONSTRAINT leads_pkey PRIMARY KEY (id);


--
-- Name: medication_definitions medication_definitions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.medication_definitions
    ADD CONSTRAINT medication_definitions_pkey PRIMARY KEY (id);


--
-- Name: method_letters method_letters_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_letters
    ADD CONSTRAINT method_letters_pkey PRIMARY KEY (id);


--
-- Name: method_pillars method_pillars_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_pillars
    ADD CONSTRAINT method_pillars_pkey PRIMARY KEY (id);


--
-- Name: methods methods_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.methods
    ADD CONSTRAINT methods_pkey PRIMARY KEY (id);


--
-- Name: notification_preferences notification_preferences_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notification_preferences
    ADD CONSTRAINT notification_preferences_pkey PRIMARY KEY (user_id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: patient_continuum_boxes patient_continuum_boxes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuum_boxes
    ADD CONSTRAINT patient_continuum_boxes_pkey PRIMARY KEY (id);


--
-- Name: patient_continuum_items patient_continuum_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuum_items
    ADD CONSTRAINT patient_continuum_items_pkey PRIMARY KEY (id);


--
-- Name: patient_continuums patient_continuums_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuums
    ADD CONSTRAINT patient_continuums_pkey PRIMARY KEY (id);


--
-- Name: patient_documents patient_documents_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_documents
    ADD CONSTRAINT patient_documents_pkey PRIMARY KEY (id);


--
-- Name: patient_magic_links patient_magic_links_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_magic_links
    ADD CONSTRAINT patient_magic_links_pkey PRIMARY KEY (id);


--
-- Name: patient_portal_invites patient_portal_invites_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_portal_invites
    ADD CONSTRAINT patient_portal_invites_pkey PRIMARY KEY (id);


--
-- Name: patient_score_group_results patient_score_group_results_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_group_results
    ADD CONSTRAINT patient_score_group_results_pkey PRIMARY KEY (id);


--
-- Name: patient_score_item_results patient_score_item_results_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT patient_score_item_results_pkey PRIMARY KEY (id);


--
-- Name: patient_score_snapshots patient_score_snapshots_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_snapshots
    ADD CONSTRAINT patient_score_snapshots_pkey PRIMARY KEY (id);


--
-- Name: patient_subscriptions patient_subscriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_subscriptions
    ADD CONSTRAINT patient_subscriptions_pkey PRIMARY KEY (id);


--
-- Name: patients patients_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT patients_pkey PRIMARY KEY (id);


--
-- Name: payment_receipt_counters payment_receipt_counters_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.payment_receipt_counters
    ADD CONSTRAINT payment_receipt_counters_pkey PRIMARY KEY (year);


--
-- Name: physical_assessments physical_assessments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.physical_assessments
    ADD CONSTRAINT physical_assessments_pkey PRIMARY KEY (id);


--
-- Name: postural_assessments postural_assessments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postural_assessments
    ADD CONSTRAINT postural_assessments_pkey PRIMARY KEY (id);


--
-- Name: prescription_medications prescription_medications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_medications
    ADD CONSTRAINT prescription_medications_pkey PRIMARY KEY (id);


--
-- Name: prescriptions prescriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_pkey PRIMARY KEY (id);


--
-- Name: prescriptions prescriptions_sncr_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_sncr_number_key UNIQUE (sncr_number);


--
-- Name: processing_jobs processing_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.processing_jobs
    ADD CONSTRAINT processing_jobs_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- Name: score_groups score_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_groups
    ADD CONSTRAINT score_groups_pkey PRIMARY KEY (id);


--
-- Name: score_item_embeddings score_item_embeddings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_embeddings
    ADD CONSTRAINT score_item_embeddings_pkey PRIMARY KEY (id);


--
-- Name: score_item_enrichment_preparation score_item_enrichment_preparation_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_enrichment_preparation
    ADD CONSTRAINT score_item_enrichment_preparation_pkey PRIMARY KEY (id);


--
-- Name: score_item_enrichment_preparation score_item_enrichment_preparation_score_item_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_enrichment_preparation
    ADD CONSTRAINT score_item_enrichment_preparation_score_item_id_key UNIQUE (score_item_id);


--
-- Name: score_item_method_pillars score_item_method_pillars_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_method_pillars
    ADD CONSTRAINT score_item_method_pillars_pkey PRIMARY KEY (score_item_id, method_pillar_id);


--
-- Name: score_item_review_history score_item_review_history_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_review_history
    ADD CONSTRAINT score_item_review_history_pkey PRIMARY KEY (id);


--
-- Name: score_items score_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT score_items_pkey PRIMARY KEY (id);


--
-- Name: score_levels score_levels_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_levels
    ADD CONSTRAINT score_levels_pkey PRIMARY KEY (id);


--
-- Name: score_subgroups score_subgroups_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_subgroups
    ADD CONSTRAINT score_subgroups_pkey PRIMARY KEY (id);


--
-- Name: subscription_plans subscription_plans_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.subscription_plans
    ADD CONSTRAINT subscription_plans_pkey PRIMARY KEY (id);


--
-- Name: telemed_lobby_tokens telemed_lobby_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.telemed_lobby_tokens
    ADD CONSTRAINT telemed_lobby_tokens_pkey PRIMARY KEY (id);


--
-- Name: patients uni_patients_cpf; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT uni_patients_cpf UNIQUE (cpf);


--
-- Name: patients uni_patients_user_id; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT uni_patients_user_id UNIQUE (user_id);


--
-- Name: users uni_users_cpf; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_cpf UNIQUE (cpf);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: waitlist_entries waitlist_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.waitlist_entries
    ADD CONSTRAINT waitlist_entries_pkey PRIMARY KEY (id);


--
-- Name: working_hours working_hours_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.working_hours
    ADD CONSTRAINT working_hours_pkey PRIMARY KEY (id);


--
-- Name: workout_mesocycles workout_mesocycles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_mesocycles
    ADD CONSTRAINT workout_mesocycles_pkey PRIMARY KEY (id);


--
-- Name: workout_periodizations workout_periodizations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_periodizations
    ADD CONSTRAINT workout_periodizations_pkey PRIMARY KEY (id);


--
-- Name: workout_plan_sessions workout_plan_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_plan_sessions
    ADD CONSTRAINT workout_plan_sessions_pkey PRIMARY KEY (id);


--
-- Name: workout_plans workout_plans_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_plans
    ADD CONSTRAINT workout_plans_pkey PRIMARY KEY (id);


--
-- Name: workout_session_exercise_logs workout_session_exercise_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercise_logs
    ADD CONSTRAINT workout_session_exercise_logs_pkey PRIMARY KEY (id);


--
-- Name: workout_session_exercises workout_session_exercises_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercises
    ADD CONSTRAINT workout_session_exercises_pkey PRIMARY KEY (id);


--
-- Name: workout_sessions workout_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_sessions
    ADD CONSTRAINT workout_sessions_pkey PRIMARY KEY (id);


--
-- Name: idx_absence_doctor; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_absence_doctor ON public.doctor_absences USING btree (doctor_id);


--
-- Name: idx_absence_range; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_absence_range ON public.doctor_absences USING btree (start_at, end_at);


--
-- Name: idx_anamnesis_author; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_author ON public.anamnesis USING btree (author_id);


--
-- Name: idx_anamnesis_consultation_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_consultation_date ON public.anamnesis USING btree (consultation_date);


--
-- Name: idx_anamnesis_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_deleted_at ON public.anamnesis USING btree (deleted_at);


--
-- Name: idx_anamnesis_item_anamnesis; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_item_anamnesis ON public.anamnesis_items USING btree (anamnesis_id);


--
-- Name: idx_anamnesis_item_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_item_order ON public.anamnesis_items USING btree ("order");


--
-- Name: idx_anamnesis_item_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_item_score_item ON public.anamnesis_items USING btree (score_item_id);


--
-- Name: idx_anamnesis_items_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_items_deleted_at ON public.anamnesis_items USING btree (deleted_at);


--
-- Name: idx_anamnesis_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_patient ON public.anamnesis USING btree (patient_id);


--
-- Name: idx_anamnesis_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_patient_id ON public.anamnesis USING btree (patient_id);


--
-- Name: idx_anamnesis_template; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template ON public.anamnesis USING btree (anamnesis_template_id);


--
-- Name: idx_anamnesis_template_area; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template_area ON public.anamnesis_templates USING btree (area);


--
-- Name: idx_anamnesis_template_item_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template_item_order ON public.anamnesis_template_items USING btree ("order");


--
-- Name: idx_anamnesis_template_item_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template_item_score_item ON public.anamnesis_template_items USING btree (score_item_id);


--
-- Name: idx_anamnesis_template_item_template; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template_item_template ON public.anamnesis_template_items USING btree (anamnesis_template_id);


--
-- Name: idx_anamnesis_template_items_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_template_items_deleted_at ON public.anamnesis_template_items USING btree (deleted_at);


--
-- Name: idx_anamnesis_templates_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anamnesis_templates_deleted_at ON public.anamnesis_templates USING btree (deleted_at);


--
-- Name: idx_anon_group_group; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_group_group ON public.anonymous_score_group_results USING btree (group_id);


--
-- Name: idx_anon_group_snapshot; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_group_snapshot ON public.anonymous_score_group_results USING btree (snapshot_id);


--
-- Name: idx_anon_item_score; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_item_score ON public.anonymous_score_items USING btree (score_item_id);


--
-- Name: idx_anon_item_session; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_item_session ON public.anonymous_score_items USING btree (session_id);


--
-- Name: idx_anon_session_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_email ON public.anonymous_score_sessions USING btree (email);


--
-- Name: idx_anon_session_expires; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_expires ON public.anonymous_score_sessions USING btree (expires_at);


--
-- Name: idx_anon_session_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_patient ON public.anonymous_score_sessions USING btree (claimed_by_patient_id);


--
-- Name: idx_anon_session_phone; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_phone ON public.anonymous_score_sessions USING btree (phone);


--
-- Name: idx_anon_session_utm_campaign; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_utm_campaign ON public.anonymous_score_sessions USING btree (utm_campaign);


--
-- Name: idx_anon_session_utm_source; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anon_session_utm_source ON public.anonymous_score_sessions USING btree (utm_source);


--
-- Name: idx_anon_snapshot_session; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_anon_snapshot_session ON public.anonymous_score_snapshots USING btree (session_id);


--
-- Name: idx_anonymous_score_group_results_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anonymous_score_group_results_deleted_at ON public.anonymous_score_group_results USING btree (deleted_at);


--
-- Name: idx_anonymous_score_items_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anonymous_score_items_deleted_at ON public.anonymous_score_items USING btree (deleted_at);


--
-- Name: idx_anonymous_score_sessions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anonymous_score_sessions_deleted_at ON public.anonymous_score_sessions USING btree (deleted_at);


--
-- Name: idx_anonymous_score_sessions_public_code; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_anonymous_score_sessions_public_code ON public.anonymous_score_sessions USING btree (public_code);


--
-- Name: idx_anonymous_score_snapshots_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_anonymous_score_snapshots_deleted_at ON public.anonymous_score_snapshots USING btree (deleted_at);


--
-- Name: idx_api_usage_logs_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_api_usage_logs_created_at ON public.api_usage_logs USING btree (created_at);


--
-- Name: idx_api_usage_logs_provider_model; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_api_usage_logs_provider_model ON public.api_usage_logs USING btree (provider, model);


--
-- Name: idx_api_usage_logs_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_api_usage_logs_user_id ON public.api_usage_logs USING btree (user_id);


--
-- Name: idx_appointment_payments_appointment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_payments_appointment_id ON public.appointment_payments USING btree (appointment_id);


--
-- Name: idx_appointment_payments_created_by_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_payments_created_by_user_id ON public.appointment_payments USING btree (created_by_user_id);


--
-- Name: idx_appointment_payments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_payments_deleted_at ON public.appointment_payments USING btree (deleted_at);


--
-- Name: idx_appointment_payments_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_payments_patient_id ON public.appointment_payments USING btree (patient_id);


--
-- Name: idx_appointment_payments_receipt_number; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_appointment_payments_receipt_number ON public.appointment_payments USING btree (receipt_number);


--
-- Name: idx_appointment_resources_appointment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_resources_appointment_id ON public.appointment_resources USING btree (appointment_id);


--
-- Name: idx_appointment_resources_resource_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointment_resources_resource_id ON public.appointment_resources USING btree (resource_id);


--
-- Name: idx_appointments_anamnesis_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_anamnesis_id ON public.appointments USING btree (anamnesis_id);


--
-- Name: idx_appointments_continuum_item_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_continuum_item_id ON public.appointments USING btree (continuum_item_id);


--
-- Name: idx_appointments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_deleted_at ON public.appointments USING btree (deleted_at);


--
-- Name: idx_appointments_doctor_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_doctor_id ON public.appointments USING btree (doctor_id);


--
-- Name: idx_appointments_end_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_end_at ON public.appointments USING btree (end_at);


--
-- Name: idx_appointments_external_calendar_event_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_external_calendar_event_id ON public.appointments USING btree (external_calendar_event_id);


--
-- Name: idx_appointments_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_patient_id ON public.appointments USING btree (patient_id);


--
-- Name: idx_appointments_scheduled_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_appointments_scheduled_at ON public.appointments USING btree (scheduled_at);


--
-- Name: idx_article_embeddings_article_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_embeddings_article_id ON public.article_embeddings USING btree (article_id);


--
-- Name: idx_article_embeddings_embedding; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_embeddings_embedding ON public.article_embeddings USING ivfflat (embedding public.vector_cosine_ops) WITH (lists='100');


--
-- Name: idx_article_embeddings_stale; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_embeddings_stale ON public.article_embeddings USING btree (is_stale) WHERE (is_stale = true);


--
-- Name: idx_article_score_items_auto_linked; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_score_items_auto_linked ON public.article_score_items USING btree (auto_linked);


--
-- Name: idx_article_score_items_auto_linked_feedback; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_score_items_auto_linked_feedback ON public.article_score_items USING btree (auto_linked, user_feedback) WHERE (auto_linked = true);


--
-- Name: idx_article_score_items_feedback; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_article_score_items_feedback ON public.article_score_items USING btree (user_feedback) WHERE (user_feedback IS NOT NULL);


--
-- Name: idx_articles_abstract_trgm; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_abstract_trgm ON public.articles USING gin (lower(public.immutable_unaccent(abstract)) public.gin_trgm_ops);


--
-- Name: idx_articles_authors_trgm; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_authors_trgm ON public.articles USING gin (lower(public.immutable_unaccent((authors)::text)) public.gin_trgm_ops);


--
-- Name: idx_articles_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_deleted_at ON public.articles USING btree (deleted_at);


--
-- Name: idx_articles_doi; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_articles_doi ON public.articles USING btree (doi);


--
-- Name: idx_articles_embedding_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_embedding_status ON public.articles USING btree (embedding_status);


--
-- Name: idx_articles_favorite; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_favorite ON public.articles USING btree (favorite);


--
-- Name: idx_articles_file_hash; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_file_hash ON public.articles USING btree (file_hash);


--
-- Name: idx_articles_journal; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_journal ON public.articles USING btree (journal);


--
-- Name: idx_articles_journal_trgm; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_journal_trgm ON public.articles USING gin (lower(public.immutable_unaccent((journal)::text)) public.gin_trgm_ops);


--
-- Name: idx_articles_parent_article_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_parent_article_id ON public.articles USING btree (parent_article_id);


--
-- Name: idx_articles_pm_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_pm_id ON public.articles USING btree (pm_id);


--
-- Name: idx_articles_publish_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_publish_date ON public.articles USING btree (publish_date);


--
-- Name: idx_articles_specialty; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_specialty ON public.articles USING btree (specialty);


--
-- Name: idx_articles_title; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_title ON public.articles USING btree (title);


--
-- Name: idx_articles_title_trgm; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_articles_title_trgm ON public.articles USING gin (lower(public.immutable_unaccent((title)::text)) public.gin_trgm_ops);


--
-- Name: idx_audit_logs_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_created_at ON public.audit_logs USING btree (created_at);


--
-- Name: idx_audit_logs_resource_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_resource_id ON public.audit_logs USING btree (resource_id);


--
-- Name: idx_audit_logs_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_user_id ON public.audit_logs USING btree (user_id);


--
-- Name: idx_auto_link_state_run_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_auto_link_state_run_id ON public.auto_link_processing_state USING btree (run_id);


--
-- Name: idx_auto_link_state_started_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_auto_link_state_started_at ON public.auto_link_processing_state USING btree (started_at DESC);


--
-- Name: idx_auto_link_state_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_auto_link_state_status ON public.auto_link_processing_state USING btree (status);


--
-- Name: idx_batch_checkpoints_run_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_checkpoints_run_id ON public.auto_link_batch_checkpoints USING btree (run_id);


--
-- Name: idx_batch_collection_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_collection_date ON public.lab_result_batches USING btree (collection_date);


--
-- Name: idx_batch_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_deleted_at ON public.lab_result_batches USING btree (deleted_at);


--
-- Name: idx_batch_doctor; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_doctor ON public.lab_result_batches USING btree (requesting_doctor_id);


--
-- Name: idx_batch_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_patient ON public.lab_result_batches USING btree (patient_id);


--
-- Name: idx_batch_request; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_request ON public.lab_result_batches USING btree (lab_request_id);


--
-- Name: idx_batch_result_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_batch_result_date ON public.lab_result_batches USING btree (result_date);


--
-- Name: idx_campaigns_created_by; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_campaigns_created_by ON public.campaigns USING btree (created_by_user_id);


--
-- Name: idx_campaigns_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_campaigns_deleted_at ON public.campaigns USING btree (deleted_at);


--
-- Name: idx_campaigns_slug; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_campaigns_slug ON public.campaigns USING btree (slug);


--
-- Name: idx_campaigns_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_campaigns_status ON public.campaigns USING btree (status);


--
-- Name: idx_campaigns_utm_source; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_campaigns_utm_source ON public.campaigns USING btree (utm_source);


--
-- Name: idx_consultation_prices_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_consultation_prices_deleted_at ON public.consultation_prices USING btree (deleted_at);


--
-- Name: idx_consultation_prices_type; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_consultation_prices_type ON public.consultation_prices USING btree (type);


--
-- Name: idx_continuum_box_templates_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_continuum_box_templates_deleted_at ON public.continuum_box_templates USING btree (deleted_at);


--
-- Name: idx_continuum_item_expected; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_continuum_item_expected ON public.patient_continuum_items USING btree (expected_date);


--
-- Name: idx_continuum_template_items_template_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_continuum_template_items_template_id ON public.continuum_template_items USING btree (template_id);


--
-- Name: idx_continuum_templates_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_continuum_templates_deleted_at ON public.continuum_templates USING btree (deleted_at);


--
-- Name: idx_conv_reads_user_owner; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_conv_reads_user_owner ON public.conversation_reads USING btree (user_id, owner_type, owner_id);


--
-- Name: idx_credential_user_provider; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_credential_user_provider ON public.calendar_credentials USING btree (user_id, provider);


--
-- Name: idx_device_token_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_device_token_unique ON public.device_tokens USING btree (token);


--
-- Name: idx_device_tokens_app_variant; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_device_tokens_app_variant ON public.device_tokens USING btree (app_variant);


--
-- Name: idx_device_tokens_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_device_tokens_deleted_at ON public.device_tokens USING btree (deleted_at);


--
-- Name: idx_device_tokens_last_seen_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_device_tokens_last_seen_at ON public.device_tokens USING btree (last_seen_at);


--
-- Name: idx_device_tokens_platform; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_device_tokens_platform ON public.device_tokens USING btree (platform);


--
-- Name: idx_device_tokens_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_device_tokens_user_id ON public.device_tokens USING btree (user_id);


--
-- Name: idx_eis_account_folder; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_eis_account_folder ON public.email_ingest_states USING btree (account, folder);


--
-- Name: idx_embedding_audit_log_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_embedding_audit_log_created_at ON public.embedding_audit_log USING btree (created_at DESC);


--
-- Name: idx_embedding_audit_log_entity; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_embedding_audit_log_entity ON public.embedding_audit_log USING btree (entity_type, entity_id);


--
-- Name: idx_embedding_queue_entity; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_embedding_queue_entity ON public.embedding_queue USING btree (entity_type, entity_id);


--
-- Name: idx_embedding_queue_priority; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_embedding_queue_priority ON public.embedding_queue USING btree (priority DESC, created_at) WHERE ((status)::text = 'pending'::text);


--
-- Name: idx_embedding_queue_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_embedding_queue_status ON public.embedding_queue USING btree (status, created_at);


--
-- Name: idx_enrichment_prep_chunks; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_enrichment_prep_chunks ON public.score_item_enrichment_preparation USING gin (selected_chunks);


--
-- Name: idx_enrichment_prep_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_enrichment_prep_created ON public.score_item_enrichment_preparation USING btree (created_at DESC);


--
-- Name: idx_enrichment_prep_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_enrichment_prep_score_item ON public.score_item_enrichment_preparation USING btree (score_item_id);


--
-- Name: idx_enrichment_prep_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_enrichment_prep_status ON public.score_item_enrichment_preparation USING btree (status);


--
-- Name: idx_exercises_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_exercises_deleted_at ON public.exercises USING btree (deleted_at);


--
-- Name: idx_exercises_external_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_exercises_external_id ON public.exercises USING btree (external_id);


--
-- Name: idx_fitness_test_results_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fitness_test_results_deleted_at ON public.fitness_test_results USING btree (deleted_at);


--
-- Name: idx_fitness_test_results_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_fitness_test_results_patient_id ON public.fitness_test_results USING btree (patient_id);


--
-- Name: idx_group_result_group; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_group_result_group ON public.patient_score_group_results USING btree (group_id);


--
-- Name: idx_group_result_snapshot; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_group_result_snapshot ON public.patient_score_group_results USING btree (snapshot_id);


--
-- Name: idx_health_check_ins_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_health_check_ins_created_at ON public.health_check_ins USING btree (created_at);


--
-- Name: idx_health_check_ins_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_health_check_ins_deleted_at ON public.health_check_ins USING btree (deleted_at);


--
-- Name: idx_health_check_ins_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_health_check_ins_patient_id ON public.health_check_ins USING btree (patient_id);


--
-- Name: idx_integrated_plan_revisions_continuum_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_integrated_plan_revisions_continuum_id ON public.integrated_plan_revisions USING btree (continuum_id);


--
-- Name: idx_integrated_plan_revisions_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_integrated_plan_revisions_created_at ON public.integrated_plan_revisions USING btree (created_at);


--
-- Name: idx_item_log_run_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_log_run_id ON public.auto_link_item_log USING btree (run_id);


--
-- Name: idx_item_log_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_log_score_item ON public.auto_link_item_log USING btree (score_item_id);


--
-- Name: idx_item_log_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_log_status ON public.auto_link_item_log USING btree (status);


--
-- Name: idx_item_result_anamnesis_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_anamnesis_item ON public.patient_score_item_results USING btree (anamnesis_item_id);


--
-- Name: idx_item_result_group; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_group ON public.patient_score_item_results USING btree (group_id);


--
-- Name: idx_item_result_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_item ON public.patient_score_item_results USING btree (item_id);


--
-- Name: idx_item_result_lab_result; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_lab_result ON public.patient_score_item_results USING btree (lab_result_id);


--
-- Name: idx_item_result_level; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_level ON public.patient_score_item_results USING btree (level_matched_id);


--
-- Name: idx_item_result_snapshot; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_item_result_snapshot ON public.patient_score_item_results USING btree (snapshot_id);


--
-- Name: idx_job_batch; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_job_batch ON public.processing_jobs USING btree (lab_result_batch_id);


--
-- Name: idx_job_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_job_created ON public.processing_jobs USING btree (created_at);


--
-- Name: idx_job_deleted; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_job_deleted ON public.processing_jobs USING btree (deleted_at);


--
-- Name: idx_job_poll; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_job_poll ON public.processing_jobs USING btree (status, created_at) WHERE ((deleted_at IS NULL) AND ((status)::text = 'pending'::text));


--
-- Name: idx_job_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_job_status ON public.processing_jobs USING btree (status);


--
-- Name: idx_junction_pillar; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_junction_pillar ON public.score_item_method_pillars USING btree (method_pillar_id);


--
-- Name: idx_junction_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_junction_score_item ON public.score_item_method_pillars USING btree (score_item_id);


--
-- Name: idx_lab_request_template_tests_template; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_request_template_tests_template ON public.lab_request_template_tests USING btree (lab_request_template_id);


--
-- Name: idx_lab_request_template_tests_test; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_request_template_tests_test ON public.lab_request_template_tests USING btree (lab_test_definition_id);


--
-- Name: idx_lab_request_templates_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_request_templates_deleted_at ON public.lab_request_templates USING btree (deleted_at);


--
-- Name: idx_lab_request_templates_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_request_templates_is_active ON public.lab_request_templates USING btree (is_active);


--
-- Name: idx_lab_request_templates_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_request_templates_name ON public.lab_request_templates USING btree (name);


--
-- Name: idx_lab_requests_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_date ON public.lab_requests USING btree (date);


--
-- Name: idx_lab_requests_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_deleted_at ON public.lab_requests USING btree (deleted_at);


--
-- Name: idx_lab_requests_doctor_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_doctor_id ON public.lab_requests USING btree (doctor_id);


--
-- Name: idx_lab_requests_lab_request_template_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_lab_request_template_id ON public.lab_requests USING btree (lab_request_template_id);


--
-- Name: idx_lab_requests_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_patient_id ON public.lab_requests USING btree (patient_id);


--
-- Name: idx_lab_requests_pdf_url; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_pdf_url ON public.lab_requests USING btree (pdf_url) WHERE (pdf_url IS NOT NULL);


--
-- Name: idx_lab_requests_template_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_requests_template_id ON public.lab_requests USING btree (lab_request_template_id);


--
-- Name: idx_lab_result_batches_collection_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_batches_collection_date ON public.lab_result_batches USING btree (collection_date);


--
-- Name: idx_lab_result_batches_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_batches_deleted_at ON public.lab_result_batches USING btree (deleted_at);


--
-- Name: idx_lab_result_batches_result_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_batches_result_date ON public.lab_result_batches USING btree (result_date);


--
-- Name: idx_lab_result_matched; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_matched ON public.lab_results USING btree (matched);


--
-- Name: idx_lab_result_value_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_value_created ON public.lab_result_values USING btree (created_at);


--
-- Name: idx_lab_result_value_deleted; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_value_deleted ON public.lab_result_values USING btree (deleted_at);


--
-- Name: idx_lab_result_value_result; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_value_result ON public.lab_result_values USING btree (lab_result_id);


--
-- Name: idx_lab_result_value_test; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_value_test ON public.lab_result_values USING btree (lab_test_definition_id);


--
-- Name: idx_lab_result_values_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_values_deleted_at ON public.lab_result_values USING btree (deleted_at);


--
-- Name: idx_lab_result_view_items_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_view_items_deleted_at ON public.lab_result_view_items USING btree (deleted_at);


--
-- Name: idx_lab_result_view_items_lab_result_view_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_view_items_lab_result_view_id ON public.lab_result_view_items USING btree (lab_result_view_id);


--
-- Name: idx_lab_result_view_items_lab_test_definition_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_view_items_lab_test_definition_id ON public.lab_result_view_items USING btree (lab_test_definition_id);


--
-- Name: idx_lab_result_view_items_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_view_items_order ON public.lab_result_view_items USING btree ("order");


--
-- Name: idx_lab_result_views_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_views_deleted_at ON public.lab_result_views USING btree (deleted_at);


--
-- Name: idx_lab_result_views_display_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_views_display_order ON public.lab_result_views USING btree (display_order);


--
-- Name: idx_lab_result_views_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_result_views_is_active ON public.lab_result_views USING btree (is_active);


--
-- Name: idx_lab_results_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_results_deleted_at ON public.lab_results USING btree (deleted_at);


--
-- Name: idx_lab_results_matched; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_results_matched ON public.lab_results USING btree (matched);


--
-- Name: idx_lab_results_test_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_results_test_type ON public.lab_results USING btree (test_type);


--
-- Name: idx_lab_test_def_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_active ON public.lab_test_definitions USING btree (is_active);


--
-- Name: idx_lab_test_def_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_category ON public.lab_test_definitions USING btree (category);


--
-- Name: idx_lab_test_def_clinical_rec; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_clinical_rec ON public.lab_test_definitions USING gin (to_tsvector('portuguese'::regconfig, COALESCE(clinical_recommendations, ''::text)));


--
-- Name: idx_lab_test_def_clinical_sig; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_clinical_sig ON public.lab_test_definitions USING gin (to_tsvector('portuguese'::regconfig, COALESCE(clinical_significance, ''::text)));


--
-- Name: idx_lab_test_def_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_code ON public.lab_test_definitions USING btree (code);


--
-- Name: idx_lab_test_def_deleted; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_deleted ON public.lab_test_definitions USING btree (deleted_at);


--
-- Name: idx_lab_test_def_loinc; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_loinc ON public.lab_test_definitions USING btree (loinc_code);


--
-- Name: idx_lab_test_def_longevity; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_longevity ON public.lab_test_definitions USING gin (to_tsvector('portuguese'::regconfig, COALESCE(longevity_context, ''::text)));


--
-- Name: idx_lab_test_def_parent; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_parent ON public.lab_test_definitions USING btree (parent_test_id);


--
-- Name: idx_lab_test_def_requestable; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_requestable ON public.lab_test_definitions USING btree (is_requestable);


--
-- Name: idx_lab_test_def_tuss; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_def_tuss ON public.lab_test_definitions USING btree (tuss_code);


--
-- Name: idx_lab_test_definitions_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_category ON public.lab_test_definitions USING btree (category);


--
-- Name: idx_lab_test_definitions_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_code ON public.lab_test_definitions USING btree (code);


--
-- Name: idx_lab_test_definitions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_deleted_at ON public.lab_test_definitions USING btree (deleted_at);


--
-- Name: idx_lab_test_definitions_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_is_active ON public.lab_test_definitions USING btree (is_active);


--
-- Name: idx_lab_test_definitions_is_requestable; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_is_requestable ON public.lab_test_definitions USING btree (is_requestable);


--
-- Name: idx_lab_test_definitions_loinc_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_loinc_code ON public.lab_test_definitions USING btree (loinc_code);


--
-- Name: idx_lab_test_definitions_parent_test_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_parent_test_id ON public.lab_test_definitions USING btree (parent_test_id);


--
-- Name: idx_lab_test_definitions_tuss_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_definitions_tuss_code ON public.lab_test_definitions USING btree (tuss_code);


--
-- Name: idx_lab_test_unit_conversions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lab_test_unit_conversions_deleted_at ON public.lab_test_unit_conversions USING btree (deleted_at);


--
-- Name: idx_lead_activities_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_created ON public.lead_activities USING btree (created_at);


--
-- Name: idx_lead_activities_email_message_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_email_message_id ON public.lead_activities USING btree (((metadata ->> 'message_id'::text)));


--
-- Name: idx_lead_activities_lead; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_lead ON public.lead_activities USING btree (lead_id);


--
-- Name: idx_lead_activities_lead_type_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_lead_type_created ON public.lead_activities USING btree (lead_id, type, created_at);


--
-- Name: idx_lead_activities_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_patient ON public.lead_activities USING btree (patient_id);


--
-- Name: idx_lead_activities_patient_doc; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_patient_doc ON public.lead_activities USING btree (patient_document_id);


--
-- Name: idx_lead_activities_patient_type_created; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_patient_type_created ON public.lead_activities USING btree (patient_id, type, created_at);


--
-- Name: idx_lead_activities_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_type ON public.lead_activities USING btree (type);


--
-- Name: idx_lead_activities_wa_message_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_lead_activities_wa_message_id ON public.lead_activities USING btree (((metadata ->> 'wa_message_id'::text)));


--
-- Name: idx_leads_assigned; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_assigned ON public.leads USING btree (assigned_to_user_id);


--
-- Name: idx_leads_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_deleted_at ON public.leads USING btree (deleted_at);


--
-- Name: idx_leads_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_email ON public.leads USING btree (email);


--
-- Name: idx_leads_last_inbound; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_last_inbound ON public.leads USING btree (last_inbound_at);


--
-- Name: idx_leads_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_patient ON public.leads USING btree (converted_patient_id);


--
-- Name: idx_leads_phone; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_phone ON public.leads USING btree (phone);


--
-- Name: idx_leads_session; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_session ON public.leads USING btree (anonymous_score_session_id);


--
-- Name: idx_leads_session_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_leads_session_unique ON public.leads USING btree (anonymous_score_session_id);


--
-- Name: idx_leads_source; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_source ON public.leads USING btree (source);


--
-- Name: idx_leads_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_status ON public.leads USING btree (status);


--
-- Name: idx_leads_utm_campaign; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_utm_campaign ON public.leads USING btree (utm_campaign);


--
-- Name: idx_leads_utm_source; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_leads_utm_source ON public.leads USING btree (utm_source);


--
-- Name: idx_medication_definitions_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_medication_definitions_category ON public.medication_definitions USING btree (category);


--
-- Name: idx_medication_definitions_common_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_medication_definitions_common_name ON public.medication_definitions USING btree (common_name);


--
-- Name: idx_medication_definitions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_medication_definitions_deleted_at ON public.medication_definitions USING btree (deleted_at);


--
-- Name: idx_method_is_default; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_is_default ON public.methods USING btree (is_default);


--
-- Name: idx_method_letter_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_letter_code ON public.method_letters USING btree (code);


--
-- Name: idx_method_letter_method; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_letter_method ON public.method_letters USING btree (method_id);


--
-- Name: idx_method_letter_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_letter_order ON public.method_letters USING btree ("order");


--
-- Name: idx_method_letters_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_letters_deleted_at ON public.method_letters USING btree (deleted_at);


--
-- Name: idx_method_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_method_name ON public.methods USING btree (name) WHERE (deleted_at IS NULL);


--
-- Name: idx_method_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_order ON public.methods USING btree ("order");


--
-- Name: idx_method_pillar_letter; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_pillar_letter ON public.method_pillars USING btree (letter_id);


--
-- Name: idx_method_pillar_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_pillar_order ON public.method_pillars USING btree ("order");


--
-- Name: idx_method_pillars_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_method_pillars_deleted_at ON public.method_pillars USING btree (deleted_at);


--
-- Name: idx_method_short_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_method_short_name ON public.methods USING btree (short_name) WHERE (deleted_at IS NULL);


--
-- Name: idx_methods_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_methods_deleted_at ON public.methods USING btree (deleted_at);


--
-- Name: idx_notification_preferences_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notification_preferences_deleted_at ON public.notification_preferences USING btree (deleted_at);


--
-- Name: idx_notifications_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_created_at ON public.notifications USING btree (created_at DESC);


--
-- Name: idx_notifications_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_deleted_at ON public.notifications USING btree (deleted_at);


--
-- Name: idx_notifications_lead; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_lead ON public.notifications USING btree (lead_id);


--
-- Name: idx_notifications_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_patient ON public.notifications USING btree (patient_id);


--
-- Name: idx_notifications_read; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_read ON public.notifications USING btree (read);


--
-- Name: idx_notifications_subscription; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_subscription ON public.notifications USING btree (subscription_id);


--
-- Name: idx_notifications_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_type ON public.notifications USING btree (type);


--
-- Name: idx_notifications_user; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_user ON public.notifications USING btree (user_id);


--
-- Name: idx_patient_continuum_boxes_continuum_item_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuum_boxes_continuum_item_id ON public.patient_continuum_boxes USING btree (continuum_item_id);


--
-- Name: idx_patient_continuum_items_appointment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuum_items_appointment_id ON public.patient_continuum_items USING btree (appointment_id);


--
-- Name: idx_patient_continuum_items_continuum_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuum_items_continuum_id ON public.patient_continuum_items USING btree (continuum_id);


--
-- Name: idx_patient_continuums_coordinator_doctor_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuums_coordinator_doctor_id ON public.patient_continuums USING btree (coordinator_doctor_id);


--
-- Name: idx_patient_continuums_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuums_deleted_at ON public.patient_continuums USING btree (deleted_at);


--
-- Name: idx_patient_continuums_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_continuums_patient_id ON public.patient_continuums USING btree (patient_id);


--
-- Name: idx_patient_documents_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_documents_deleted_at ON public.patient_documents USING btree (deleted_at);


--
-- Name: idx_patient_documents_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_documents_patient_id ON public.patient_documents USING btree (patient_id);


--
-- Name: idx_patient_documents_wa_msg; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_documents_wa_msg ON public.patient_documents USING btree (origin_wa_message_id);


--
-- Name: idx_patient_magic_links_hash; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_patient_magic_links_hash ON public.patient_magic_links USING btree (token_hash);


--
-- Name: idx_patient_magic_links_token; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_patient_magic_links_token ON public.patient_magic_links USING btree (token);


--
-- Name: idx_patient_magic_links_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_magic_links_user_id ON public.patient_magic_links USING btree (user_id);


--
-- Name: idx_patient_portal_invites_hash; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_patient_portal_invites_hash ON public.patient_portal_invites USING btree (token_hash);


--
-- Name: idx_patient_portal_invites_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_portal_invites_patient_id ON public.patient_portal_invites USING btree (patient_id);


--
-- Name: idx_patient_portal_invites_token; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_patient_portal_invites_token ON public.patient_portal_invites USING btree (token);


--
-- Name: idx_patient_score_group_results_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_score_group_results_deleted_at ON public.patient_score_group_results USING btree (deleted_at);


--
-- Name: idx_patient_score_item_results_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_score_item_results_deleted_at ON public.patient_score_item_results USING btree (deleted_at);


--
-- Name: idx_patient_score_snapshots_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_score_snapshots_deleted_at ON public.patient_score_snapshots USING btree (deleted_at);


--
-- Name: idx_patient_subscriptions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_subscriptions_deleted_at ON public.patient_subscriptions USING btree (deleted_at);


--
-- Name: idx_patient_subscriptions_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_subscriptions_patient ON public.patient_subscriptions USING btree (patient_id);


--
-- Name: idx_patient_subscriptions_plan; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_subscriptions_plan ON public.patient_subscriptions USING btree (subscription_plan_id);


--
-- Name: idx_patient_subscriptions_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patient_subscriptions_status ON public.patient_subscriptions USING btree (status);


--
-- Name: idx_patients_cpf_blind; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_cpf_blind ON public.patients USING btree (cpf_blind_index);


--
-- Name: idx_patients_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_deleted_at ON public.patients USING btree (deleted_at);


--
-- Name: idx_patients_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_email ON public.patients USING btree (email);


--
-- Name: idx_patients_email_lower; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_email_lower ON public.patients USING btree (lower((email)::text));


--
-- Name: idx_patients_municipality; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_municipality ON public.patients USING btree (municipality);


--
-- Name: idx_patients_phone; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_phone ON public.patients USING btree (phone);


--
-- Name: idx_patients_source; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_source ON public.patients USING btree (source);


--
-- Name: idx_patients_state; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_state ON public.patients USING btree (state);


--
-- Name: idx_patients_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_patients_user_id ON public.patients USING btree (user_id);


--
-- Name: idx_physical_assessments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_physical_assessments_deleted_at ON public.physical_assessments USING btree (deleted_at);


--
-- Name: idx_physical_assessments_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_physical_assessments_patient_id ON public.physical_assessments USING btree (patient_id);


--
-- Name: idx_plan_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plan_deleted_at ON public.subscription_plans USING btree (deleted_at);


--
-- Name: idx_plan_is_active; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plan_is_active ON public.subscription_plans USING btree (is_active);


--
-- Name: idx_plan_method; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plan_method ON public.subscription_plans USING btree (method_id);


--
-- Name: idx_plan_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plan_name ON public.subscription_plans USING btree (name);


--
-- Name: idx_plan_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plan_order ON public.subscription_plans USING btree ("order");


--
-- Name: idx_postural_assessments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postural_assessments_deleted_at ON public.postural_assessments USING btree (deleted_at);


--
-- Name: idx_postural_assessments_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postural_assessments_patient_id ON public.postural_assessments USING btree (patient_id);


--
-- Name: idx_postural_assessments_physical_assessment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postural_assessments_physical_assessment_id ON public.postural_assessments USING btree (physical_assessment_id);


--
-- Name: idx_preparation_stale; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_preparation_stale ON public.score_item_enrichment_preparation USING btree (status) WHERE ((status)::text = 'stale'::text);


--
-- Name: idx_prescription_medications_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescription_medications_category ON public.prescription_medications USING btree (category);


--
-- Name: idx_prescription_medications_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescription_medications_deleted_at ON public.prescription_medications USING btree (deleted_at);


--
-- Name: idx_prescription_medications_medication_definition_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescription_medications_medication_definition_id ON public.prescription_medications USING btree (medication_definition_id);


--
-- Name: idx_prescription_medications_prescription_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescription_medications_prescription_id ON public.prescription_medications USING btree (prescription_id);


--
-- Name: idx_prescriptions_category; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_category ON public.prescriptions USING btree (category);


--
-- Name: idx_prescriptions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_deleted_at ON public.prescriptions USING btree (deleted_at);


--
-- Name: idx_prescriptions_doctor_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_doctor_id ON public.prescriptions USING btree (doctor_id);


--
-- Name: idx_prescriptions_is_used; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_is_used ON public.prescriptions USING btree (is_used);


--
-- Name: idx_prescriptions_medication_definition_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_medication_definition_id ON public.prescriptions USING btree (medication_definition_id);


--
-- Name: idx_prescriptions_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_patient_id ON public.prescriptions USING btree (patient_id);


--
-- Name: idx_prescriptions_prescription_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_prescription_date ON public.prescriptions USING btree (prescription_date);


--
-- Name: idx_prescriptions_sncr_number; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_sncr_number ON public.prescriptions USING btree (sncr_number);


--
-- Name: idx_prescriptions_valid_until; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prescriptions_valid_until ON public.prescriptions USING btree (valid_until);


--
-- Name: idx_processing_jobs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_processing_jobs_deleted_at ON public.processing_jobs USING btree (deleted_at);


--
-- Name: idx_refresh_tokens_expires_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_refresh_tokens_expires_at ON public.refresh_tokens USING btree (expires_at);


--
-- Name: idx_refresh_tokens_revoked_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_refresh_tokens_revoked_at ON public.refresh_tokens USING btree (revoked_at);


--
-- Name: idx_refresh_tokens_token_hash; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_refresh_tokens_token_hash ON public.refresh_tokens USING btree (token_hash);


--
-- Name: idx_refresh_tokens_type; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_refresh_tokens_type ON public.refresh_tokens USING btree (type);


--
-- Name: idx_refresh_tokens_used_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_refresh_tokens_used_at ON public.refresh_tokens USING btree (used_at);


--
-- Name: idx_refresh_tokens_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);


--
-- Name: idx_result_batch; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_result_batch ON public.lab_results USING btree (lab_result_batch_id);


--
-- Name: idx_result_test_def; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_result_test_def ON public.lab_results USING btree (lab_test_definition_id);


--
-- Name: idx_review_history_reviewed_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_review_history_reviewed_at ON public.score_item_review_history USING btree (reviewed_at);


--
-- Name: idx_review_history_score_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_review_history_score_item ON public.score_item_review_history USING btree (score_item_id);


--
-- Name: idx_score_group_name; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_score_group_name ON public.score_groups USING btree (name) WHERE (deleted_at IS NULL);


--
-- Name: idx_score_group_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_group_order ON public.score_groups USING btree ("order");


--
-- Name: idx_score_groups_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_groups_deleted_at ON public.score_groups USING btree (deleted_at);


--
-- Name: idx_score_item_embeddings_embedding; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_embeddings_embedding ON public.score_item_embeddings USING ivfflat (embedding public.vector_cosine_ops) WITH (lists='50');


--
-- Name: idx_score_item_embeddings_score_item_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_embeddings_score_item_id ON public.score_item_embeddings USING btree (score_item_id);


--
-- Name: idx_score_item_embeddings_stale; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_embeddings_stale ON public.score_item_embeddings USING btree (is_stale) WHERE (is_stale = true);


--
-- Name: idx_score_item_light; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_light ON public.score_items USING btree (is_light_version);


--
-- Name: idx_score_item_light_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_light_order ON public.score_items USING btree (light_order);


--
-- Name: idx_score_item_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_order ON public.score_items USING btree ("order");


--
-- Name: idx_score_item_parent; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_parent ON public.score_items USING btree (parent_item_id);


--
-- Name: idx_score_item_subgroup; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_item_subgroup ON public.score_items USING btree (subgroup_id);


--
-- Name: idx_score_items_anamnese_item_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_items_anamnese_item_code ON public.score_items USING btree (anamnese_item_code) WHERE (anamnese_item_code IS NOT NULL);


--
-- Name: idx_score_items_code_non_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_items_code_non_unique ON public.score_items USING btree (lab_test_code) WHERE (deleted_at IS NULL);


--
-- Name: idx_score_items_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_items_deleted_at ON public.score_items USING btree (deleted_at);


--
-- Name: idx_score_items_lab_test_code; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_items_lab_test_code ON public.score_items USING btree (lab_test_code);


--
-- Name: idx_score_level_item; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_level_item ON public.score_levels USING btree (item_id);


--
-- Name: idx_score_level_level; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_level_level ON public.score_levels USING btree (level);


--
-- Name: idx_score_levels_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_levels_deleted_at ON public.score_levels USING btree (deleted_at);


--
-- Name: idx_score_levels_operator; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_levels_operator ON public.score_levels USING btree (operator);


--
-- Name: idx_score_subgroup_group; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_subgroup_group ON public.score_subgroups USING btree (group_id);


--
-- Name: idx_score_subgroup_order; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_subgroup_order ON public.score_subgroups USING btree ("order");


--
-- Name: idx_score_subgroups_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_score_subgroups_deleted_at ON public.score_subgroups USING btree (deleted_at);


--
-- Name: idx_snapshot_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_snapshot_date ON public.patient_score_snapshots USING btree (calculated_at DESC);


--
-- Name: idx_snapshot_patient; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_snapshot_patient ON public.patient_score_snapshots USING btree (patient_id);


--
-- Name: idx_snapshot_user; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_snapshot_user ON public.patient_score_snapshots USING btree (calculated_by_user_id);


--
-- Name: idx_subscription_plans_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_subscription_plans_deleted_at ON public.subscription_plans USING btree (deleted_at);


--
-- Name: idx_telemed_lobby_tokens_appointment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_telemed_lobby_tokens_appointment_id ON public.telemed_lobby_tokens USING btree (appointment_id);


--
-- Name: idx_telemed_lobby_tokens_token; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_telemed_lobby_tokens_token ON public.telemed_lobby_tokens USING btree (token);


--
-- Name: idx_unique_test_secondary_unit; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_unique_test_secondary_unit ON public.lab_test_unit_conversions USING btree (lab_test_definition_id, secondary_unit) WHERE (deleted_at IS NULL);


--
-- Name: idx_unit_conv_test; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_unit_conv_test ON public.lab_test_unit_conversions USING btree (lab_test_definition_id, deleted_at);


--
-- Name: idx_users_crm; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_crm ON public.users USING btree (crm);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_o_auth_provider; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_o_auth_provider ON public.users USING btree (o_auth_provider);


--
-- Name: idx_users_o_auth_provider_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_o_auth_provider_id ON public.users USING btree (o_auth_provider_id);


--
-- Name: idx_users_oauth_provider; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_oauth_provider ON public.users USING btree (oauth_provider) WHERE (oauth_provider IS NOT NULL);


--
-- Name: idx_users_oauth_unique; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users_oauth_unique ON public.users USING btree (oauth_provider, oauth_provider_id) WHERE ((oauth_provider IS NOT NULL) AND (oauth_provider_id IS NOT NULL));


--
-- Name: idx_users_preferences; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_preferences ON public.users USING gin (preferences);


--
-- Name: idx_users_selected_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_selected_patient_id ON public.users USING btree (selected_patient_id);


--
-- Name: idx_waitlist_entries_created_by_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_waitlist_entries_created_by_user_id ON public.waitlist_entries USING btree (created_by_user_id);


--
-- Name: idx_waitlist_entries_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_waitlist_entries_deleted_at ON public.waitlist_entries USING btree (deleted_at);


--
-- Name: idx_waitlist_entries_doctor_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_waitlist_entries_doctor_id ON public.waitlist_entries USING btree (doctor_id);


--
-- Name: idx_waitlist_entries_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_waitlist_entries_patient_id ON public.waitlist_entries USING btree (patient_id);


--
-- Name: idx_waitlist_entries_scheduled_appointment_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_waitlist_entries_scheduled_appointment_id ON public.waitlist_entries USING btree (scheduled_appointment_id);


--
-- Name: idx_wh_doctor; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_wh_doctor ON public.working_hours USING btree (doctor_id);


--
-- Name: idx_workout_mesocycles_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_mesocycles_deleted_at ON public.workout_mesocycles USING btree (deleted_at);


--
-- Name: idx_workout_mesocycles_periodization_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_mesocycles_periodization_id ON public.workout_mesocycles USING btree (periodization_id);


--
-- Name: idx_workout_periodizations_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_periodizations_deleted_at ON public.workout_periodizations USING btree (deleted_at);


--
-- Name: idx_workout_periodizations_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_periodizations_patient_id ON public.workout_periodizations USING btree (patient_id);


--
-- Name: idx_workout_plan_sessions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_plan_sessions_deleted_at ON public.workout_plan_sessions USING btree (deleted_at);


--
-- Name: idx_workout_plan_sessions_plan_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_plan_sessions_plan_id ON public.workout_plan_sessions USING btree (plan_id);


--
-- Name: idx_workout_plans_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_plans_deleted_at ON public.workout_plans USING btree (deleted_at);


--
-- Name: idx_workout_plans_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_plans_patient_id ON public.workout_plans USING btree (patient_id);


--
-- Name: idx_workout_plans_public_code; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_workout_plans_public_code ON public.workout_plans USING btree (public_code);


--
-- Name: idx_workout_session_exercise_logs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercise_logs_deleted_at ON public.workout_session_exercise_logs USING btree (deleted_at);


--
-- Name: idx_workout_session_exercise_logs_exercise_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercise_logs_exercise_id ON public.workout_session_exercise_logs USING btree (exercise_id);


--
-- Name: idx_workout_session_exercise_logs_plan_exercise_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercise_logs_plan_exercise_id ON public.workout_session_exercise_logs USING btree (plan_exercise_id);


--
-- Name: idx_workout_session_exercise_logs_session_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercise_logs_session_id ON public.workout_session_exercise_logs USING btree (session_id);


--
-- Name: idx_workout_session_exercises_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercises_deleted_at ON public.workout_session_exercises USING btree (deleted_at);


--
-- Name: idx_workout_session_exercises_session_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_session_exercises_session_id ON public.workout_session_exercises USING btree (session_id);


--
-- Name: idx_workout_sessions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_sessions_deleted_at ON public.workout_sessions USING btree (deleted_at);


--
-- Name: idx_workout_sessions_patient_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_sessions_patient_id ON public.workout_sessions USING btree (patient_id);


--
-- Name: idx_workout_sessions_plan_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_sessions_plan_id ON public.workout_sessions USING btree (plan_id);


--
-- Name: idx_workout_sessions_plan_session_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_sessions_plan_session_id ON public.workout_sessions USING btree (plan_session_id);


--
-- Name: idx_workout_sessions_scheduled_date; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_workout_sessions_scheduled_date ON public.workout_sessions USING btree (scheduled_date);


--
-- Name: method_letters_deleted_at_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX method_letters_deleted_at_idx ON public.method_letters USING btree (deleted_at);


--
-- Name: method_pillars_deleted_at_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX method_pillars_deleted_at_idx ON public.method_pillars USING btree (deleted_at);


--
-- Name: methods_deleted_at_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX methods_deleted_at_idx ON public.methods USING btree (deleted_at);


--
-- Name: uniq_session_exercise_set; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX uniq_session_exercise_set ON public.workout_session_exercise_logs USING btree (session_id, plan_exercise_id, set_number);


--
-- Name: uq_patient_documents_wa_msg; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX uq_patient_documents_wa_msg ON public.patient_documents USING btree (origin_wa_message_id);


--
-- Name: lab_result_values trg_lab_result_values_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trg_lab_result_values_updated_at BEFORE UPDATE ON public.lab_result_values FOR EACH ROW EXECUTE FUNCTION public.update_lab_result_values_updated_at();


--
-- Name: lab_test_definitions trg_lab_test_definitions_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trg_lab_test_definitions_updated_at BEFORE UPDATE ON public.lab_test_definitions FOR EACH ROW EXECUTE FUNCTION public.update_lab_test_definitions_updated_at();


--
-- Name: score_item_enrichment_preparation trigger_enrichment_prep_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_enrichment_prep_updated_at BEFORE UPDATE ON public.score_item_enrichment_preparation FOR EACH ROW EXECUTE FUNCTION public.update_enrichment_prep_timestamp();


--
-- Name: api_usage_logs api_usage_logs_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.api_usage_logs
    ADD CONSTRAINT api_usage_logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: article_embeddings article_embeddings_article_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_embeddings
    ADD CONSTRAINT article_embeddings_article_id_fkey FOREIGN KEY (article_id) REFERENCES public.articles(id) ON DELETE CASCADE;


--
-- Name: article_score_items article_score_items_linked_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_score_items
    ADD CONSTRAINT article_score_items_linked_by_fkey FOREIGN KEY (linked_by) REFERENCES public.users(id);


--
-- Name: auto_link_batch_checkpoints auto_link_batch_checkpoints_run_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_batch_checkpoints
    ADD CONSTRAINT auto_link_batch_checkpoints_run_id_fkey FOREIGN KEY (run_id) REFERENCES public.auto_link_processing_state(run_id) ON DELETE CASCADE;


--
-- Name: auto_link_item_log auto_link_item_log_run_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auto_link_item_log
    ADD CONSTRAINT auto_link_item_log_run_id_fkey FOREIGN KEY (run_id) REFERENCES public.auto_link_processing_state(run_id) ON DELETE CASCADE;


--
-- Name: anamnesis fk_anamnesis_anamnesis_template; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT fk_anamnesis_anamnesis_template FOREIGN KEY (anamnesis_template_id) REFERENCES public.anamnesis_templates(id) ON DELETE SET NULL;


--
-- Name: anamnesis fk_anamnesis_author; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT fk_anamnesis_author FOREIGN KEY (author_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: anamnesis_items fk_anamnesis_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_items
    ADD CONSTRAINT fk_anamnesis_items FOREIGN KEY (anamnesis_id) REFERENCES public.anamnesis(id) ON DELETE CASCADE;


--
-- Name: anamnesis_items fk_anamnesis_items_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_items
    ADD CONSTRAINT fk_anamnesis_items_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE RESTRICT;


--
-- Name: anamnesis fk_anamnesis_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT fk_anamnesis_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: anamnesis_template_items fk_anamnesis_template_items_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_template_items
    ADD CONSTRAINT fk_anamnesis_template_items_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE RESTRICT;


--
-- Name: anamnesis_template_items fk_anamnesis_templates_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anamnesis_template_items
    ADD CONSTRAINT fk_anamnesis_templates_items FOREIGN KEY (anamnesis_template_id) REFERENCES public.anamnesis_templates(id) ON DELETE CASCADE;


--
-- Name: anonymous_score_group_results fk_anonymous_score_group_results_group; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_group_results
    ADD CONSTRAINT fk_anonymous_score_group_results_group FOREIGN KEY (group_id) REFERENCES public.score_groups(id) ON DELETE RESTRICT;


--
-- Name: anonymous_score_items fk_anonymous_score_items_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_items
    ADD CONSTRAINT fk_anonymous_score_items_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE RESTRICT;


--
-- Name: anonymous_score_sessions fk_anonymous_score_sessions_claimed_by_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_sessions
    ADD CONSTRAINT fk_anonymous_score_sessions_claimed_by_patient FOREIGN KEY (claimed_by_patient_id) REFERENCES public.patients(id) ON DELETE SET NULL;


--
-- Name: anonymous_score_items fk_anonymous_score_sessions_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_items
    ADD CONSTRAINT fk_anonymous_score_sessions_items FOREIGN KEY (session_id) REFERENCES public.anonymous_score_sessions(id) ON DELETE CASCADE;


--
-- Name: anonymous_score_snapshots fk_anonymous_score_sessions_snapshot; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_snapshots
    ADD CONSTRAINT fk_anonymous_score_sessions_snapshot FOREIGN KEY (session_id) REFERENCES public.anonymous_score_sessions(id) ON DELETE CASCADE;


--
-- Name: anonymous_score_group_results fk_anonymous_score_snapshots_group_results; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.anonymous_score_group_results
    ADD CONSTRAINT fk_anonymous_score_snapshots_group_results FOREIGN KEY (snapshot_id) REFERENCES public.anonymous_score_snapshots(id) ON DELETE CASCADE;


--
-- Name: appointment_payments fk_appointment_payments_appointment; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointment_payments
    ADD CONSTRAINT fk_appointment_payments_appointment FOREIGN KEY (appointment_id) REFERENCES public.appointments(id) ON DELETE SET NULL;


--
-- Name: appointment_payments fk_appointment_payments_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointment_payments
    ADD CONSTRAINT fk_appointment_payments_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: appointments fk_appointments_anamnesis; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT fk_appointments_anamnesis FOREIGN KEY (anamnesis_id) REFERENCES public.anamnesis(id);


--
-- Name: appointments fk_appointments_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT fk_appointments_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: appointments fk_appointments_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT fk_appointments_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: article_embeddings fk_article_embeddings_article; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_embeddings
    ADD CONSTRAINT fk_article_embeddings_article FOREIGN KEY (article_id) REFERENCES public.articles(id) ON DELETE CASCADE;


--
-- Name: article_score_items fk_article_score_items_article; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_score_items
    ADD CONSTRAINT fk_article_score_items_article FOREIGN KEY (article_id) REFERENCES public.articles(id);


--
-- Name: article_score_items fk_article_score_items_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article_score_items
    ADD CONSTRAINT fk_article_score_items_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id);


--
-- Name: articles fk_articles_chapters; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.articles
    ADD CONSTRAINT fk_articles_chapters FOREIGN KEY (parent_article_id) REFERENCES public.articles(id);


--
-- Name: audit_logs fk_audit_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT fk_audit_logs_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: calendar_credentials fk_calendar_credentials_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.calendar_credentials
    ADD CONSTRAINT fk_calendar_credentials_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: campaigns fk_campaigns_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT fk_campaigns_created_by FOREIGN KEY (created_by_user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: continuum_template_items fk_continuum_templates_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.continuum_template_items
    ADD CONSTRAINT fk_continuum_templates_items FOREIGN KEY (template_id) REFERENCES public.continuum_templates(id) ON DELETE CASCADE;


--
-- Name: doctor_absences fk_doctor_absences_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.doctor_absences
    ADD CONSTRAINT fk_doctor_absences_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: fitness_test_results fk_fitness_test_results_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fitness_test_results
    ADD CONSTRAINT fk_fitness_test_results_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: fitness_test_results fk_fitness_test_results_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.fitness_test_results
    ADD CONSTRAINT fk_fitness_test_results_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: health_check_ins fk_health_check_ins_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.health_check_ins
    ADD CONSTRAINT fk_health_check_ins_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: integrated_plan_revisions fk_integrated_plan_revisions_updated_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.integrated_plan_revisions
    ADD CONSTRAINT fk_integrated_plan_revisions_updated_by FOREIGN KEY (updated_by_id) REFERENCES public.users(id);


--
-- Name: lab_request_template_tests fk_lab_request_template_tests_lab_request_template; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_template_tests
    ADD CONSTRAINT fk_lab_request_template_tests_lab_request_template FOREIGN KEY (lab_request_template_id) REFERENCES public.lab_request_templates(id) ON DELETE CASCADE;


--
-- Name: lab_request_template_tests fk_lab_request_template_tests_lab_test_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_template_tests
    ADD CONSTRAINT fk_lab_request_template_tests_lab_test_definition FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE CASCADE;


--
-- Name: lab_requests fk_lab_requests_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT fk_lab_requests_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: lab_requests fk_lab_requests_lab_request_template; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT fk_lab_requests_lab_request_template FOREIGN KEY (lab_request_template_id) REFERENCES public.lab_request_templates(id) ON DELETE SET NULL;


--
-- Name: lab_requests fk_lab_requests_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT fk_lab_requests_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: processing_jobs fk_lab_result_batch; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.processing_jobs
    ADD CONSTRAINT fk_lab_result_batch FOREIGN KEY (lab_result_batch_id) REFERENCES public.lab_result_batches(id) ON DELETE CASCADE;


--
-- Name: lab_result_batches fk_lab_result_batches_lab_request; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT fk_lab_result_batches_lab_request FOREIGN KEY (lab_request_id) REFERENCES public.lab_requests(id) ON DELETE SET NULL;


--
-- Name: lab_results fk_lab_result_batches_lab_results; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_results
    ADD CONSTRAINT fk_lab_result_batches_lab_results FOREIGN KEY (lab_result_batch_id) REFERENCES public.lab_result_batches(id) ON DELETE CASCADE;


--
-- Name: lab_result_batches fk_lab_result_batches_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT fk_lab_result_batches_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: lab_result_batches fk_lab_result_batches_requesting_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT fk_lab_result_batches_requesting_doctor FOREIGN KEY (requesting_doctor_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: lab_result_values fk_lab_result_values_lab_test_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_values
    ADD CONSTRAINT fk_lab_result_values_lab_test_definition FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE RESTRICT;


--
-- Name: lab_result_view_items fk_lab_result_view_items_lab_test_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_view_items
    ADD CONSTRAINT fk_lab_result_view_items_lab_test_definition FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE RESTRICT;


--
-- Name: lab_result_view_items fk_lab_result_views_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_view_items
    ADD CONSTRAINT fk_lab_result_views_items FOREIGN KEY (lab_result_view_id) REFERENCES public.lab_result_views(id) ON DELETE CASCADE;


--
-- Name: lab_results fk_lab_results_batch; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_results
    ADD CONSTRAINT fk_lab_results_batch FOREIGN KEY (lab_result_batch_id) REFERENCES public.lab_result_batches(id) ON DELETE CASCADE;


--
-- Name: lab_result_values fk_lab_results_lab_result_values; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_values
    ADD CONSTRAINT fk_lab_results_lab_result_values FOREIGN KEY (lab_result_id) REFERENCES public.lab_results(id) ON DELETE CASCADE;


--
-- Name: lab_results fk_lab_results_lab_test_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_results
    ADD CONSTRAINT fk_lab_results_lab_test_definition FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE SET NULL;


--
-- Name: score_items fk_lab_test_definitions_score_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT fk_lab_test_definitions_score_items FOREIGN KEY (lab_test_code) REFERENCES public.lab_test_definitions(code);


--
-- Name: lab_test_definitions fk_lab_test_definitions_sub_tests; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_definitions
    ADD CONSTRAINT fk_lab_test_definitions_sub_tests FOREIGN KEY (parent_test_id) REFERENCES public.lab_test_definitions(id) ON DELETE SET NULL;


--
-- Name: lab_test_unit_conversions fk_lab_test_unit_conversions_lab_test_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_unit_conversions
    ADD CONSTRAINT fk_lab_test_unit_conversions_lab_test_definition FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE CASCADE;


--
-- Name: lead_activities fk_lead_activities_actor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lead_activities
    ADD CONSTRAINT fk_lead_activities_actor FOREIGN KEY (actor_user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: lead_activities fk_leads_activities; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lead_activities
    ADD CONSTRAINT fk_leads_activities FOREIGN KEY (lead_id) REFERENCES public.leads(id) ON DELETE CASCADE;


--
-- Name: leads fk_leads_assigned_to; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.leads
    ADD CONSTRAINT fk_leads_assigned_to FOREIGN KEY (assigned_to_user_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: leads fk_leads_converted_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.leads
    ADD CONSTRAINT fk_leads_converted_patient FOREIGN KEY (converted_patient_id) REFERENCES public.patients(id) ON DELETE SET NULL;


--
-- Name: leads fk_leads_session; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.leads
    ADD CONSTRAINT fk_leads_session FOREIGN KEY (anonymous_score_session_id) REFERENCES public.anonymous_score_sessions(id) ON DELETE SET NULL;


--
-- Name: method_letters fk_method_letters_method; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_letters
    ADD CONSTRAINT fk_method_letters_method FOREIGN KEY (method_id) REFERENCES public.methods(id) ON DELETE CASCADE;


--
-- Name: method_pillars fk_method_letters_pillars; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_pillars
    ADD CONSTRAINT fk_method_letters_pillars FOREIGN KEY (letter_id) REFERENCES public.method_letters(id) ON DELETE CASCADE;


--
-- Name: method_pillars fk_method_pillars_letter; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_pillars
    ADD CONSTRAINT fk_method_pillars_letter FOREIGN KEY (letter_id) REFERENCES public.method_letters(id) ON DELETE CASCADE;


--
-- Name: method_letters fk_methods_letters; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.method_letters
    ADD CONSTRAINT fk_methods_letters FOREIGN KEY (method_id) REFERENCES public.methods(id) ON DELETE CASCADE;


--
-- Name: notification_preferences fk_notification_preferences_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notification_preferences
    ADD CONSTRAINT fk_notification_preferences_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notifications_lead; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_lead FOREIGN KEY (lead_id) REFERENCES public.leads(id) ON DELETE SET NULL;


--
-- Name: notifications fk_notifications_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notifications_subscription; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_subscription FOREIGN KEY (subscription_id) REFERENCES public.patient_subscriptions(id) ON DELETE CASCADE;


--
-- Name: notifications fk_notifications_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: lab_test_definitions fk_parent_test; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_definitions
    ADD CONSTRAINT fk_parent_test FOREIGN KEY (parent_test_id) REFERENCES public.lab_test_definitions(id) ON DELETE SET NULL;


--
-- Name: patient_continuums fk_patient_continuums_coordinator_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuums
    ADD CONSTRAINT fk_patient_continuums_coordinator_doctor FOREIGN KEY (coordinator_doctor_id) REFERENCES public.users(id);


--
-- Name: patient_continuum_items fk_patient_continuums_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuum_items
    ADD CONSTRAINT fk_patient_continuums_items FOREIGN KEY (continuum_id) REFERENCES public.patient_continuums(id);


--
-- Name: patient_continuums fk_patient_continuums_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_continuums
    ADD CONSTRAINT fk_patient_continuums_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: patient_documents fk_patient_documents_uploaded_by_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_documents
    ADD CONSTRAINT fk_patient_documents_uploaded_by_user FOREIGN KEY (uploaded_by) REFERENCES public.users(id);


--
-- Name: patient_score_group_results fk_patient_score_group_results_group; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_group_results
    ADD CONSTRAINT fk_patient_score_group_results_group FOREIGN KEY (group_id) REFERENCES public.score_groups(id) ON DELETE RESTRICT;


--
-- Name: patient_score_group_results fk_patient_score_group_results_snapshot; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_group_results
    ADD CONSTRAINT fk_patient_score_group_results_snapshot FOREIGN KEY (snapshot_id) REFERENCES public.patient_score_snapshots(id) ON DELETE CASCADE;


--
-- Name: patient_score_item_results fk_patient_score_item_results_anamnesis_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_anamnesis_item FOREIGN KEY (anamnesis_item_id) REFERENCES public.anamnesis_items(id) ON DELETE SET NULL;


--
-- Name: patient_score_item_results fk_patient_score_item_results_group; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_group FOREIGN KEY (group_id) REFERENCES public.score_groups(id) ON DELETE RESTRICT;


--
-- Name: patient_score_item_results fk_patient_score_item_results_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_item FOREIGN KEY (item_id) REFERENCES public.score_items(id) ON DELETE RESTRICT;


--
-- Name: patient_score_item_results fk_patient_score_item_results_lab_result; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_lab_result FOREIGN KEY (lab_result_id) REFERENCES public.lab_results(id) ON DELETE SET NULL;


--
-- Name: patient_score_item_results fk_patient_score_item_results_level; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_level FOREIGN KEY (level_matched_id) REFERENCES public.score_levels(id) ON DELETE SET NULL;


--
-- Name: patient_score_item_results fk_patient_score_item_results_snapshot; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_item_results
    ADD CONSTRAINT fk_patient_score_item_results_snapshot FOREIGN KEY (snapshot_id) REFERENCES public.patient_score_snapshots(id) ON DELETE CASCADE;


--
-- Name: patient_score_snapshots fk_patient_score_snapshots_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_snapshots
    ADD CONSTRAINT fk_patient_score_snapshots_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: patient_score_snapshots fk_patient_score_snapshots_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_score_snapshots
    ADD CONSTRAINT fk_patient_score_snapshots_user FOREIGN KEY (calculated_by_user_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: patient_subscriptions fk_patient_subscriptions_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_subscriptions
    ADD CONSTRAINT fk_patient_subscriptions_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: patient_subscriptions fk_patient_subscriptions_subscription_plan; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_subscriptions
    ADD CONSTRAINT fk_patient_subscriptions_subscription_plan FOREIGN KEY (subscription_plan_id) REFERENCES public.subscription_plans(id) ON DELETE RESTRICT;


--
-- Name: lead_activities fk_patients_activities; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lead_activities
    ADD CONSTRAINT fk_patients_activities FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: patients fk_patients_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT fk_patients_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: physical_assessments fk_physical_assessments_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.physical_assessments
    ADD CONSTRAINT fk_physical_assessments_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: physical_assessments fk_physical_assessments_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.physical_assessments
    ADD CONSTRAINT fk_physical_assessments_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: postural_assessments fk_postural_assessments_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postural_assessments
    ADD CONSTRAINT fk_postural_assessments_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: postural_assessments fk_postural_assessments_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postural_assessments
    ADD CONSTRAINT fk_postural_assessments_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: prescription_medications fk_prescription_medications_medication_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_medications
    ADD CONSTRAINT fk_prescription_medications_medication_definition FOREIGN KEY (medication_definition_id) REFERENCES public.medication_definitions(id);


--
-- Name: prescriptions fk_prescriptions_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT fk_prescriptions_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: prescriptions fk_prescriptions_medication_definition; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT fk_prescriptions_medication_definition FOREIGN KEY (medication_definition_id) REFERENCES public.medication_definitions(id) ON DELETE SET NULL;


--
-- Name: prescription_medications fk_prescriptions_medications; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_medications
    ADD CONSTRAINT fk_prescriptions_medications FOREIGN KEY (prescription_id) REFERENCES public.prescriptions(id) ON DELETE CASCADE;


--
-- Name: prescriptions fk_prescriptions_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT fk_prescriptions_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: processing_jobs fk_processing_jobs_lab_result_batch; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.processing_jobs
    ADD CONSTRAINT fk_processing_jobs_lab_result_batch FOREIGN KEY (lab_result_batch_id) REFERENCES public.lab_result_batches(id) ON DELETE CASCADE;


--
-- Name: refresh_tokens fk_refresh_tokens_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT fk_refresh_tokens_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: score_subgroups fk_score_groups_subgroups; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_subgroups
    ADD CONSTRAINT fk_score_groups_subgroups FOREIGN KEY (group_id) REFERENCES public.score_groups(id) ON DELETE CASCADE;


--
-- Name: score_item_embeddings fk_score_item_embeddings_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_embeddings
    ADD CONSTRAINT fk_score_item_embeddings_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_item_method_pillars fk_score_item_method_pillars_method_pillar; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_method_pillars
    ADD CONSTRAINT fk_score_item_method_pillars_method_pillar FOREIGN KEY (method_pillar_id) REFERENCES public.method_pillars(id);


--
-- Name: score_item_method_pillars fk_score_item_method_pillars_pillar; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_method_pillars
    ADD CONSTRAINT fk_score_item_method_pillars_pillar FOREIGN KEY (method_pillar_id) REFERENCES public.method_pillars(id) ON DELETE CASCADE;


--
-- Name: score_item_method_pillars fk_score_item_method_pillars_score_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_method_pillars
    ADD CONSTRAINT fk_score_item_method_pillars_score_item FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_items fk_score_items_child_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT fk_score_items_child_items FOREIGN KEY (parent_item_id) REFERENCES public.score_items(id) ON DELETE SET NULL;


--
-- Name: score_levels fk_score_items_levels; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_levels
    ADD CONSTRAINT fk_score_items_levels FOREIGN KEY (item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_items fk_score_items_parent; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT fk_score_items_parent FOREIGN KEY (parent_item_id) REFERENCES public.score_items(id) ON DELETE SET NULL;


--
-- Name: score_items fk_score_items_subgroup; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT fk_score_items_subgroup FOREIGN KEY (subgroup_id) REFERENCES public.score_subgroups(id) ON DELETE CASCADE;


--
-- Name: score_levels fk_score_levels_item; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_levels
    ADD CONSTRAINT fk_score_levels_item FOREIGN KEY (item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_subgroups fk_score_subgroups_group; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_subgroups
    ADD CONSTRAINT fk_score_subgroups_group FOREIGN KEY (group_id) REFERENCES public.score_groups(id) ON DELETE CASCADE;


--
-- Name: score_items fk_score_subgroups_items; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_items
    ADD CONSTRAINT fk_score_subgroups_items FOREIGN KEY (subgroup_id) REFERENCES public.score_subgroups(id) ON DELETE CASCADE;


--
-- Name: subscription_plans fk_subscription_plans_method; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.subscription_plans
    ADD CONSTRAINT fk_subscription_plans_method FOREIGN KEY (method_id) REFERENCES public.methods(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: lab_test_unit_conversions fk_unit_conv_test; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_test_unit_conversions
    ADD CONSTRAINT fk_unit_conv_test FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE CASCADE;


--
-- Name: device_tokens fk_users_device_tokens; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.device_tokens
    ADD CONSTRAINT fk_users_device_tokens FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users fk_users_selected_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_selected_patient FOREIGN KEY (selected_patient_id) REFERENCES public.patients(id) ON DELETE SET NULL;


--
-- Name: lab_result_values fk_value_lab_result; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_values
    ADD CONSTRAINT fk_value_lab_result FOREIGN KEY (lab_result_id) REFERENCES public.lab_results(id) ON DELETE CASCADE;


--
-- Name: lab_result_values fk_value_lab_test_def; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_values
    ADD CONSTRAINT fk_value_lab_test_def FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE RESTRICT;


--
-- Name: waitlist_entries fk_waitlist_entries_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.waitlist_entries
    ADD CONSTRAINT fk_waitlist_entries_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: waitlist_entries fk_waitlist_entries_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.waitlist_entries
    ADD CONSTRAINT fk_waitlist_entries_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: working_hours fk_working_hours_doctor; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.working_hours
    ADD CONSTRAINT fk_working_hours_doctor FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: workout_periodizations fk_workout_periodizations_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_periodizations
    ADD CONSTRAINT fk_workout_periodizations_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: workout_mesocycles fk_workout_periodizations_mesocycles; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_mesocycles
    ADD CONSTRAINT fk_workout_periodizations_mesocycles FOREIGN KEY (periodization_id) REFERENCES public.workout_periodizations(id);


--
-- Name: workout_periodizations fk_workout_periodizations_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_periodizations
    ADD CONSTRAINT fk_workout_periodizations_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: workout_session_exercises fk_workout_plan_sessions_exercises; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercises
    ADD CONSTRAINT fk_workout_plan_sessions_exercises FOREIGN KEY (session_id) REFERENCES public.workout_plan_sessions(id);


--
-- Name: workout_plans fk_workout_plans_created_by; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_plans
    ADD CONSTRAINT fk_workout_plans_created_by FOREIGN KEY (created_by_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: workout_plans fk_workout_plans_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_plans
    ADD CONSTRAINT fk_workout_plans_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: workout_plan_sessions fk_workout_plans_sessions; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_plan_sessions
    ADD CONSTRAINT fk_workout_plans_sessions FOREIGN KEY (plan_id) REFERENCES public.workout_plans(id);


--
-- Name: workout_session_exercise_logs fk_workout_session_exercise_logs_exercise; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercise_logs
    ADD CONSTRAINT fk_workout_session_exercise_logs_exercise FOREIGN KEY (exercise_id) REFERENCES public.exercises(id) ON DELETE RESTRICT;


--
-- Name: workout_session_exercise_logs fk_workout_session_exercise_logs_plan_exercise; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercise_logs
    ADD CONSTRAINT fk_workout_session_exercise_logs_plan_exercise FOREIGN KEY (plan_exercise_id) REFERENCES public.workout_session_exercises(id) ON DELETE CASCADE;


--
-- Name: workout_session_exercises fk_workout_session_exercises_exercise; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercises
    ADD CONSTRAINT fk_workout_session_exercises_exercise FOREIGN KEY (exercise_id) REFERENCES public.exercises(id) ON DELETE RESTRICT;


--
-- Name: workout_session_exercise_logs fk_workout_sessions_logs; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_session_exercise_logs
    ADD CONSTRAINT fk_workout_sessions_logs FOREIGN KEY (session_id) REFERENCES public.workout_sessions(id);


--
-- Name: workout_sessions fk_workout_sessions_patient; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_sessions
    ADD CONSTRAINT fk_workout_sessions_patient FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: workout_sessions fk_workout_sessions_plan; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_sessions
    ADD CONSTRAINT fk_workout_sessions_plan FOREIGN KEY (plan_id) REFERENCES public.workout_plans(id) ON DELETE CASCADE;


--
-- Name: workout_sessions fk_workout_sessions_plan_session; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.workout_sessions
    ADD CONSTRAINT fk_workout_sessions_plan_session FOREIGN KEY (plan_session_id) REFERENCES public.workout_plan_sessions(id) ON DELETE CASCADE;


--
-- Name: lab_request_template_tests lab_request_template_tests_lab_request_template_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_template_tests
    ADD CONSTRAINT lab_request_template_tests_lab_request_template_id_fkey FOREIGN KEY (lab_request_template_id) REFERENCES public.lab_request_templates(id) ON DELETE CASCADE;


--
-- Name: lab_request_template_tests lab_request_template_tests_lab_test_definition_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_request_template_tests
    ADD CONSTRAINT lab_request_template_tests_lab_test_definition_id_fkey FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE CASCADE;


--
-- Name: lab_requests lab_requests_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT lab_requests_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: lab_requests lab_requests_lab_request_template_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT lab_requests_lab_request_template_id_fkey FOREIGN KEY (lab_request_template_id) REFERENCES public.lab_request_templates(id) ON DELETE SET NULL;


--
-- Name: lab_requests lab_requests_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_requests
    ADD CONSTRAINT lab_requests_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: lab_result_batches lab_result_batches_lab_request_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT lab_result_batches_lab_request_id_fkey FOREIGN KEY (lab_request_id) REFERENCES public.lab_requests(id) ON DELETE SET NULL;


--
-- Name: lab_result_batches lab_result_batches_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT lab_result_batches_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: lab_result_batches lab_result_batches_requesting_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_result_batches
    ADD CONSTRAINT lab_result_batches_requesting_doctor_id_fkey FOREIGN KEY (requesting_doctor_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: lab_results lab_results_lab_test_definition_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.lab_results
    ADD CONSTRAINT lab_results_lab_test_definition_id_fkey FOREIGN KEY (lab_test_definition_id) REFERENCES public.lab_test_definitions(id) ON DELETE SET NULL;


--
-- Name: notifications notifications_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: notifications notifications_subscription_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_subscription_id_fkey FOREIGN KEY (subscription_id) REFERENCES public.patient_subscriptions(id) ON DELETE CASCADE;


--
-- Name: notifications notifications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: patient_subscriptions patient_subscriptions_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_subscriptions
    ADD CONSTRAINT patient_subscriptions_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: patient_subscriptions patient_subscriptions_subscription_plan_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patient_subscriptions
    ADD CONSTRAINT patient_subscriptions_subscription_plan_id_fkey FOREIGN KEY (subscription_plan_id) REFERENCES public.subscription_plans(id) ON DELETE RESTRICT;


--
-- Name: prescriptions prescriptions_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE RESTRICT;


--
-- Name: prescriptions prescriptions_medication_definition_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_medication_definition_id_fkey FOREIGN KEY (medication_definition_id) REFERENCES public.medication_definitions(id) ON DELETE SET NULL;


--
-- Name: prescriptions prescriptions_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(id) ON DELETE CASCADE;


--
-- Name: score_item_embeddings score_item_embeddings_score_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_embeddings
    ADD CONSTRAINT score_item_embeddings_score_item_id_fkey FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_item_enrichment_preparation score_item_enrichment_preparation_score_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_enrichment_preparation
    ADD CONSTRAINT score_item_enrichment_preparation_score_item_id_fkey FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: score_item_review_history score_item_review_history_score_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.score_item_review_history
    ADD CONSTRAINT score_item_review_history_score_item_id_fkey FOREIGN KEY (score_item_id) REFERENCES public.score_items(id) ON DELETE CASCADE;


--
-- Name: subscription_plans subscription_plans_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.subscription_plans
    ADD CONSTRAINT subscription_plans_method_id_fkey FOREIGN KEY (method_id) REFERENCES public.methods(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Baseline forward-only: sem down (dados de saúde; reverter = restaurar backup).
SELECT 1;
-- +goose StatementEnd
