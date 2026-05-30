-- Migration: UUID v7 Generation Function (RFC 9562)
-- Autor: Claude Sonnet 4.5
-- Data: 2026-02-02

-- Drop existing function (idempotent)
DROP FUNCTION IF EXISTS uuid_generate_v7() CASCADE;
DROP FUNCTION IF EXISTS uuid_v7_from_timestamp(TIMESTAMP, UUID) CASCADE;

-- Function 1: Generate fresh UUID v7 (current timestamp)
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

-- Function 2: Convert UUID v4 to v7 using timestamp (for migration)
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

-- Add comments
COMMENT ON FUNCTION uuid_generate_v7() IS
'Generates time-ordered UUID v7 (RFC 9562). First 48 bits = Unix timestamp ms.';

COMMENT ON FUNCTION uuid_v7_from_timestamp(TIMESTAMP, UUID) IS
'Converts UUID v4 to v7 using provided timestamp. Used for data migration.';

-- Test functions (optional, remove in production)
DO $$
DECLARE
    v_uuid1 UUID;
    v_uuid2 UUID;
    v_migrated UUID;
BEGIN
    -- Test uuid_generate_v7()
    v_uuid1 := uuid_generate_v7();
    PERFORM pg_sleep(0.01);
    v_uuid2 := uuid_generate_v7();

    ASSERT v_uuid1 < v_uuid2, 'UUID v7 should be time-ordered';
    ASSERT (get_byte(uuid_send(v_uuid1), 6) >> 4) = 7, 'Should be version 7';

    -- Test uuid_v7_from_timestamp()
    v_migrated := uuid_v7_from_timestamp('2026-01-01'::TIMESTAMP, gen_random_uuid());
    ASSERT (get_byte(uuid_send(v_migrated), 6) >> 4) = 7, 'Migrated should be v7';

    RAISE NOTICE '✅ UUID v7 functions created successfully';
END;
$$;
