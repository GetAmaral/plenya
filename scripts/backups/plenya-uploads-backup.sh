#!/bin/sh
# Backup diário do uploads do EMR (mídia WhatsApp + patient-docs). O bind mount
# protege contra churn de container; ISTO protege contra falha de disco.
# RODA NA VPS em /usr/local/bin/ (cron root, 30 3 * * *). Cópia versionada no repo.
set -eu
SRC=/data/coolify/applications/kgcuxgvmnbx6yya35e3ca2v0/uploads   # bind dos uploads (UUID do app api)
DEST=/home/deploy/.plenya-vps-secrets/backups
TS=$(date +%Y%m%d_%H%M)
mkdir -p "$DEST"
tar -czf "$DEST/uploads_${TS}.tar.gz" -C "$SRC" . 2>/dev/null
# rotação: mantém os 7 backups mais recentes
ls -1t "$DEST"/uploads_*.tar.gz 2>/dev/null | tail -n +8 | xargs -r rm -f
