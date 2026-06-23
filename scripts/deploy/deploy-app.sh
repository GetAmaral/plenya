#!/usr/bin/env bash
# Deploy MANUAL de UM app do Plenya no Coolify.
#
# Auto-deploy está DESLIGADO de propósito (push no master NÃO deploya nada) — a VPS
# tem 8GB e rebuildar os 3 apps juntos causa OOM. Deploy é deliberado e por-app:
# rode este script só pro(s) app(s) cujo código mudou.
#
#   Uso:  scripts/deploy/deploy-app.sh <api|web|site>
#
# Roda da máquina de dev. Lê o token do Coolify de ~/.plenya-vps-secrets/coolify.env.
# Faz a higiene (limpa fila presa), dispara UM deploy e espera o container novo subir
# + healthcheck. Ver memórias coolify_deploy_orphan_lock_procedure e plenya_deploy_manual.
set -euo pipefail

APP="${1:-}"
case "$APP" in
  api)  UUID=kgcuxgvmnbx6yya35e3ca2v0; PREFIX=kgcuxgvmnbx6; HEALTH=https://api.plenyasaude.com.br/health ;;
  web)  UUID=nwbhak0fscs2th13gz5g9zjm; PREFIX=nwbhak0fsc;   HEALTH=https://app.plenyasaude.com.br/login ;;
  site) UUID=ycpklto5n1qjkmelhdp0pvhf; PREFIX=ycpklto5n1;   HEALTH=https://plenyasaude.com.br/ ;;
  *) echo "uso: $0 <api|web|site>"; exit 2 ;;
esac

TOKEN=$(grep '^COOLIFY_API_TOKEN=' "$HOME/.plenya-vps-secrets/coolify.env" | cut -d= -f2-)
[ -n "$TOKEN" ] || { echo "token Coolify não encontrado"; exit 1; }

echo "==> higiene: marca deploys presos como failed"
ssh plenya "sudo docker exec coolify-db psql -U coolify -d coolify -c \"UPDATE application_deployment_queues SET status='failed' WHERE status IN ('queued','in_progress');\"" >/dev/null

OLD=$(ssh plenya "sudo docker ps --format '{{.Names}}' | grep '^${PREFIX}' || true")
echo "==> container atual: ${OLD:-(nenhum)}"

echo "==> dispara deploy de plenya-${APP} ($UUID)"
curl -sS "https://coolify.plenyasaude.com.br/api/v1/deploy?uuid=${UUID}&force=true" \
  -H "Authorization: Bearer $TOKEN"; echo

echo "==> aguardando container novo + healthcheck (até ~13min)"
# O status da API do Coolify costuma ficar em "running:unknown" mesmo no sucesso;
# por isso confiamos no sinal real: container NOVO de pé + health 200.
for i in $(seq 1 40); do
  sleep 20
  NEW=$(ssh plenya "sudo docker ps --format '{{.Names}}' | grep '^${PREFIX}' || true")
  CODE=$(curl -s -o /dev/null -w '%{http_code}' "$HEALTH" || echo 000)
  echo "  [$i] container=${NEW:-(nenhum)} health=$CODE"
  if [ -n "$NEW" ] && [ "$NEW" != "$OLD" ] && [ "$CODE" = "200" ]; then
    echo "==> OK: plenya-${APP} no ar ($NEW, health $CODE)"
    exit 0
  fi
done
echo "==> ATENÇÃO: não confirmou em ~13min. Cheque os logs do build no Coolify."
exit 1
