═══════════════════════════════════════════════════════════════
  FILES CREATED/MODIFIED DURING MIGRATION
═══════════════════════════════════════════════════════════════

NEW FILES CREATED:
─────────────────────────────────────────────────────────────

1. apps/api/database/init.sql (3.4KB)
   - UUID v7 function definitions
   - Auto-executed on new DB initialization
   - Ensures UUID v7 support in all environments

2. apps/api/database/migrations/20260202_create_uuid_v7_function.sql (3.0KB)
   - Creates uuid_generate_v7() function
   - Creates uuid_v7_from_timestamp() helper
   - Includes test validation

3. apps/api/database/migrations/20260202_migrate_all_uuids_to_v7_v2.sql (9.0KB)
   - Main data migration script
   - Migrates 5,369 UUIDs from v4 to v7
   - Creates uuid_migration_map table for rollback
   - Executed in 0.58 seconds

4. apps/api/database/migrations/20260202_update_defaults_to_uuid_v7.sql (2.3KB)
   - Updates DEFAULT values for all table ID columns
   - Changes from gen_random_uuid() to uuid_generate_v7()

5. backups/backup_pre_uuid_v7_20260201_124233.sql (11MB)
   - Complete database backup before migration
   - 15,872 lines
   - Tested and validated for restore

MODIFIED FILES:
─────────────────────────────────────────────────────────────

Configuration:
──────────────
1. docker-compose.yml
   - Added init.sql volume mount to db service

Go Models (7 files):
─────────────────────
2. apps/api/internal/models/user.go
   - Added BeforeCreate hook for UUID v7
   - Removed default:gen_random_uuid() from tag

3. apps/api/internal/models/patient.go
   - Added BeforeCreate hook
   - Removed default from tag

4. apps/api/internal/models/article.go
   - Fixed uuid.New() → uuid.NewV7()
   - Removed default from tag

5. apps/api/internal/models/lab_result.go
   - Added BeforeCreate hook
   - Added gorm import
   - Removed default from tag

6. apps/api/internal/models/audit_log.go
   - Added BeforeCreate hook
   - Added gorm import
   - Removed default from tag

7. apps/api/internal/models/appointment.go
   - Added BeforeCreate hook
   - Removed default from tag

8. apps/api/internal/models/prescription.go
   - Added BeforeCreate hook
   - Removed default from tag

Additional Models (tags updated, already had BeforeCreate):
────────────────────────────────────────────────────────────
9-19. All other models in apps/api/internal/models/*.go
   - Removed default:gen_random_uuid() from GORM tags
   - BeforeCreate hooks already existed

Documentation:
──────────────
20. CLAUDE.md
   - Added comprehensive UUID v7 section
   - Updated Go model example
   - Added BeforeCreate hook template
   - Added validation queries
   - Total additions: ~100 lines

═══════════════════════════════════════════════════════════════
  COMMANDS EXECUTED
═══════════════════════════════════════════════════════════════

1. Database Backup:
   docker compose exec -T db pg_dump -U plenya_user -d plenya_db \
     --clean --if-exists --create > backup_pre_uuid_v7_20260201_124233.sql

2. Apply init.sql to existing DB:
   docker compose exec -T db psql -U plenya_user -d plenya_db \
     < apps/api/database/init.sql

3. Execute Migrations (in order):
   - 20260202_create_uuid_v7_function.sql
   - 20260202_migrate_all_uuids_to_v7_v2.sql
   - 20260202_update_defaults_to_uuid_v7.sql

4. Global replace in Go models:
   find apps/api/internal/models -name "*.go" -type f \
     -exec sed -i 's/;default:gen_random_uuid()//' {} \;

5. Rebuild and restart services:
   docker compose up -d --build api
   docker compose start web

═══════════════════════════════════════════════════════════════
  DATABASE OBJECTS CREATED
═══════════════════════════════════════════════════════════════

Functions:
──────────
1. uuid_generate_v7() → uuid
   - Generates time-ordered UUID v7 with current timestamp
   
2. uuid_v7_from_timestamp(ts TIMESTAMP, old_uuid UUID) → uuid
   - Converts UUID v4 to v7 using provided timestamp
   - Used during migration
   
3. uuid_v7_timestamp(u UUID) → TIMESTAMP
   - Extracts timestamp from UUID v7
   - Useful for debugging

Tables:
───────
4. uuid_migration_map
   - Stores v4→v7 mapping for rollback
   - 5,369 records
   - Keep for 7 days, then DROP

═══════════════════════════════════════════════════════════════
  VERIFICATION QUERIES
═══════════════════════════════════════════════════════════════

-- Check UUID version
SELECT (get_byte(uuid_send(id), 6) >> 4) AS version FROM users LIMIT 1;
-- Should return: 7

-- Check all tables
SELECT table_name, column_default 
FROM information_schema.columns
WHERE column_name = 'id' AND table_schema = 'public'
ORDER BY table_name;
-- Most should show: uuid_generate_v7()

-- Check migration count
SELECT COUNT(*) FROM uuid_migration_map;
-- Should return: 5369

-- Test new UUID generation
SELECT uuid_generate_v7();
-- Should return UUID v7 format

═══════════════════════════════════════════════════════════════
