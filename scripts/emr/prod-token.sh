#!/usr/bin/env bash
# Devolve um accessToken da API de PRODUÇÃO, para `emr.py` e `plano.py` escreverem prontuário.
#
#   export EMR_API=https://api.plenyasaude.com.br
#   export EMR_TOKEN=$(scripts/emr/prod-token.sh)
#
# Lê a credencial da CONTA DE SERVIÇO em ~/.plenya-vps-secrets/emr-prod-api.env, que é um arquivo
# separado de propósito: `emr-prod.env` guarda os segredos do servidor (chave de criptografia,
# senha do banco, senha de e-mail) e não deve ser aberto para isto.
#
# Por que existe: a porta do prontuário (`emr.py`) sempre soube falar com prod, e nunca houve token.
# Ver docs/emr/dados-de-paciente-em-producao.md para o porquê de prontuário não ir por psql.
#
# NÃO imprime nada além do token no stdout — o resto vai para stderr, para poder ser usado em
# substituição de comando sem sujar a variável.
set -euo pipefail

ENVFILE="${EMR_PROD_ENV:-$HOME/.plenya-vps-secrets/emr-prod-api.env}"

# Reserva: a conta do médico, que é a que a auditoria de produção já mostra nas cargas anteriores
# ("Getulio Amaral Filho · curl/8.5.0"). Decisão dele, em 2026-09-06: escrever em nome próprio e
# reusar esta credencial, em vez de provisionar conta de serviço.
#
# A conta de serviço continua sendo melhor e o motivo está em
# docs/emr/dados-de-paciente-em-producao.md: com ela a trilha diz que foi um script, e a assinatura
# do médico continua sendo o que transforma rascunho em ato. Enquanto não existir, a trilha atribui
# ao médico o que o script escreveu — o que é aceitável porque é escolha dele, e não um efeito
# colateral silencioso.
FALLBACK="$HOME/.plenya-vps-secrets/emr-prod.env"
if [ ! -f "$ENVFILE" ] && [ -f "$FALLBACK" ]; then
  echo "usando a conta do médico de $FALLBACK (não há conta de serviço)" >&2
  # `|| true` em cada um: sob `set -euo pipefail`, um grep sem match derruba o script inteiro, e
  # este arquivo não tem EMR_API (é o .env do servidor, não do cliente).
  API=$(grep -m1 '^EMR_API=' "$FALLBACK" | cut -d= -f2- || true)
  USER=$(grep -m1 '^ADMIN_INITIAL_EMAIL=' "$FALLBACK" | cut -d= -f2- || true)
  PASS=$(grep -m1 '^ADMIN_INITIAL_PASSWORD=' "$FALLBACK" | cut -d= -f2- || true)
  ENVFILE=""
fi

if [ -z "${USER:-}" ] && [ ! -f "$ENVFILE" ]; then
  cat >&2 <<MSG
não achei $ENVFILE

Provisione a conta de serviço uma vez (ver docs/emr/dados-de-paciente-em-producao.md) e crie o
arquivo com:

  EMR_API=https://api.plenyasaude.com.br
  EMR_USER=servico@plenyasaude.com.br
  EMR_PASSWORD=<senha só desta conta>

A conta precisa do papel 'doctor' para criar receita, ou 'nurse' se não precisar. Deixe o 2FA
DESLIGADO nela: com 2FA não existe login não-interativo.
MSG
  exit 1
fi

# `grep`+`cut` em vez de `source`: senha com caracteres de shell (|, $, espaço) quebra o source, e
# foi exatamente o que aconteceu com o token do Coolify (ver docs/infra/vps-runbook.md).
if [ -n "$ENVFILE" ]; then
  val() { grep -m1 "^$1=" "$ENVFILE" | cut -d= -f2- ; }
  API=$(val EMR_API); USER=$(val EMR_USER); PASS=$(val EMR_PASSWORD)
fi
API=${API:-https://api.plenyasaude.com.br}
[ -n "${USER:-}" ] && [ -n "${PASS:-}" ] || { echo "credencial ausente (nem conta de serviço, nem reserva)" >&2; exit 1; }

RESP=$(USER="$USER" PASS="$PASS" python3 - "$API" <<'PY'
import json, os, sys, urllib.request, urllib.error
api = sys.argv[1].rstrip('/')
corpo = json.dumps({'email': os.environ['USER'], 'password': os.environ['PASS']}).encode()
req = urllib.request.Request(api + '/api/v1/auth/login', data=corpo,
                             headers={'Content-Type': 'application/json'})
try:
    d = json.load(urllib.request.urlopen(req, timeout=30))
except urllib.error.HTTPError as e:
    sys.exit(f'login falhou ({e.code}): {e.read().decode()[:200]}')
if d.get('requires2FA') or d.get('mfaRequired'):
    sys.exit('a conta de serviço está com 2FA ligado; desligue para permitir login não-interativo')
tok = d.get('accessToken')
if not tok:
    sys.exit('resposta sem accessToken: ' + json.dumps(list(d))[:120])
print(tok)
PY
)

echo "token obtido para $USER em $API" >&2
printf '%s' "$RESP"
