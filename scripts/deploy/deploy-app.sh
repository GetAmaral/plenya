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

# Higiene da fila, restrita a ESTE app.
#
# Era uma limpeza geral, e isso matava build de outro app que estivesse rodando. Aconteceu em
# 2026-09-09: o push acordou o webhook do site-getulio, o deploy do api entrou quatro segundos
# depois, e a higiene marcou o build do site-getulio como `failed` com ele no meio do caminho. O
# processo continua rodando mesmo marcado, então os dois builds disputaram a VPS de 8GB, e o clone
# do api travou em 226MB e ficou parado até o script desistir.
#
# O `WHERE application_id` resolve: cada deploy limpa a própria fila, e ninguém encosta na do
# vizinho. Um deploy alheio em andamento é motivo para ESPERAR, não para derrubar.
echo "==> higiene: marca deploys presos DESTE app como failed"
ssh plenya "sudo docker exec coolify-db psql -U coolify -d coolify -c \"UPDATE application_deployment_queues SET status='failed', finished_at=now() WHERE status IN ('queued','in_progress') AND application_id = (SELECT id::text FROM applications WHERE uuid='${UUID}');\"" >/dev/null

# Build de OUTRO app em andamento: esperar é mais barato que competir por RAM e disco. O clone do
# repositório sozinho passa de 200MB, e dois ao mesmo tempo foi o que travou o deploy acima.
OUTRO=$(ssh plenya "sudo docker exec coolify-db psql -U coolify -d coolify -At -c \"SELECT count(*) FROM application_deployment_queues q JOIN applications a ON a.id::text = q.application_id WHERE q.status IN ('queued','in_progress') AND a.uuid <> '${UUID}';\"" 2>/dev/null || echo 0)
if [ "${OUTRO:-0}" != "0" ]; then
  echo "==> há $OUTRO deploy(s) de outro app em andamento; esperando até 10min para não competir"
  for _ in $(seq 1 60); do
    sleep 10
    OUTRO=$(ssh plenya "sudo docker exec coolify-db psql -U coolify -d coolify -At -c \"SELECT count(*) FROM application_deployment_queues q JOIN applications a ON a.id::text = q.application_id WHERE q.status IN ('queued','in_progress') AND a.uuid <> '${UUID}';\"" 2>/dev/null || echo 0)
    [ "${OUTRO:-0}" = "0" ] && break
  done
  [ "${OUTRO:-0}" != "0" ] && echo "==> ATENÇÃO: outro deploy segue rodando; seguindo mesmo assim" >&2
fi

OLD=$(ssh plenya "sudo docker ps --format '{{.Names}}' | grep '^${PREFIX}' || true")
echo "==> container atual: ${OLD:-(nenhum)}"

# Guard de RAM: o build do web (Next/Turbopack) é faminto de memória; num box de 8GB com os
# containers rodando, começar sem folga leva a OOM/thrash que derruba o HOST inteiro (ssh+HTTP
# fora). Abortar antes é mais barato que recuperar. Override: DEPLOY_SKIP_RAM_CHECK=1.
MIN_MB=1500; [ "$APP" = "web" ] && MIN_MB=2800
AVAIL_MB=$(ssh plenya "awk '/MemAvailable/{print int(\$2/1024)}' /proc/meminfo" 2>/dev/null || echo 0)
echo "==> RAM disponível na VPS: ${AVAIL_MB}MB (mínimo p/ ${APP}: ${MIN_MB}MB)"
if [ "${DEPLOY_SKIP_RAM_CHECK:-0}" != "1" ] && [ "$AVAIL_MB" -lt "$MIN_MB" ]; then
  echo "ABORTADO: pouca RAM livre na VPS — risco de OOM no build derrubar o host." >&2
  echo "          Libere memória (ou espere) e tente de novo; ou force com DEPLOY_SKIP_RAM_CHECK=1." >&2
  exit 3
fi

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
