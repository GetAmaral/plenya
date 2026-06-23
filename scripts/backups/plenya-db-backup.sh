#!/bin/sh
# Backup diário do Postgres do EMR (plenya_db), formato custom (-Fc) restaurável via
# pg_restore. Self-contained: usa as próprias envs do container (sem secret externo).
# RODA NA VPS em /usr/local/bin/ (cron root, 0 3 * * *). Cópia versionada no repo.
set -eu
DEST=/home/deploy/.plenya-vps-secrets/backups
PGC=mb511beqjtgd7nsjlnngh3m6   # container Postgres (Coolify); ver memória plenya_vps_emr
TS=$(date +%Y%m%d_%H%M)
OUT="$DEST/db_plenya_${TS}.dump"
mkdir -p "$DEST"
if ! docker exec "$PGC" sh -c 'PGPASSWORD="$POSTGRES_PASSWORD" pg_dump -U "$POSTGRES_USER" -d "$POSTGRES_DB" -Fc' > "$OUT.tmp" 2>/tmp/plenya-db-backup.err; then
  rm -f "$OUT.tmp"
  echo "pg_dump falhou: $(cat /tmp/plenya-db-backup.err 2>/dev/null)" >&2
  exit 1
fi
# sanidade: dump custom começa com a assinatura "PGDMP"
if ! head -c 5 "$OUT.tmp" | grep -q 'PGDMP'; then
  rm -f "$OUT.tmp"; echo "dump sem assinatura PGDMP — abortado" >&2; exit 1
fi
mv "$OUT.tmp" "$OUT"
# rotação: mantém os 14 dumps mais recentes
ls -1t "$DEST"/db_plenya_*.dump 2>/dev/null | tail -n +15 | xargs -r rm -f
