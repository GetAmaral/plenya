#!/bin/sh
# Puxa os backups da VPS de produção (banco + uploads) pra máquina de dev.
# RODA NO DEV em /home/user/backups/ (cron local 15 9 * * *). Cópia versionada no repo.
#
# A VPS já gera/rotaciona em /home/deploy/.plenya-vps-secrets/backups/. Este script
# ESPELHA essa pasta em /home/user/backups/prod/vps/ (cópia off-host do prod).
# Best-effort: WSL não fica 24/7; rodar manual quando precisar.
set -eu
DEST=/home/user/backups/prod/vps
mkdir -p "$DEST"
# --delete: o espelho acompanha a rotação da VPS (tamanho limitado). Dumps pra guardar
# pra sempre: mover de prod/vps/ para prod/ (fora do espelho).
rsync -az --delete plenya:/home/deploy/.plenya-vps-secrets/backups/ "$DEST/"
echo "$(date '+%Y-%m-%d %H:%M') pull-prod-backups ok ($(ls -1 "$DEST" | wc -l) arquivos)" >> "$DEST/../pull.log"
