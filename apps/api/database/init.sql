-- PostgreSQL Initialization Script
-- Auto-executed when database is first created
-- Ensures UUID v7 support is available immediately
-- Author: Claude Sonnet 4.5
-- Date: 2026-02-02

-- Enable required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "vector"; -- pgvector for RAG semantic search

-- Create UUID v7 generation function (RFC 9562)
CREATE OR REPLACE FUNCTION uuid_generate_v7()
RETURNS uuid
LANGUAGE plpgsql
VOLATILE
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

-- Create conversion function (for migrations)
CREATE OR REPLACE FUNCTION uuid_v7_from_timestamp(ts TIMESTAMP, old_uuid UUID)
RETURNS uuid
LANGUAGE plpgsql
IMMUTABLE
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

-- Create helper function to extract timestamp from UUID v7
CREATE OR REPLACE FUNCTION uuid_v7_timestamp(u UUID)
RETURNS TIMESTAMP
LANGUAGE SQL
IMMUTABLE
AS $$
SELECT to_timestamp(
    ('x' || substring(u::TEXT, 1, 8) || substring(u::TEXT, 10, 4))::BIT(48)::BIGINT / 1000.0
);
$$;

-- Add comments
COMMENT ON FUNCTION uuid_generate_v7() IS
'Generates time-ordered UUID v7 (RFC 9562). First 48 bits = Unix timestamp ms. Used as DEFAULT in all Plenya tables.';

COMMENT ON FUNCTION uuid_v7_from_timestamp(TIMESTAMP, UUID) IS
'Converts UUID v4 to v7 using provided timestamp. Used for data migration only.';

COMMENT ON FUNCTION uuid_v7_timestamp(UUID) IS
'Extracts timestamp from UUID v7. Useful for debugging and analysis.';

-- Test the functions
DO $$
DECLARE
    v_uuid1 UUID;
    v_uuid2 UUID;
    v_ts TIMESTAMP;
BEGIN
    -- Test uuid_generate_v7()
    v_uuid1 := uuid_generate_v7();
    PERFORM pg_sleep(0.01);
    v_uuid2 := uuid_generate_v7();

    ASSERT v_uuid1 < v_uuid2, 'UUID v7 should be time-ordered';
    ASSERT (get_byte(uuid_send(v_uuid1), 6) >> 4) = 7, 'Should be version 7';

    -- Test uuid_v7_timestamp()
    v_ts := uuid_v7_timestamp(v_uuid1);
    ASSERT v_ts IS NOT NULL, 'Should extract timestamp';

    RAISE NOTICE '✅ UUID v7 functions initialized successfully';
    RAISE NOTICE '   Sample UUID v7: %', v_uuid1;
    RAISE NOTICE '   Extracted timestamp: %', v_ts;
END;
$$;
