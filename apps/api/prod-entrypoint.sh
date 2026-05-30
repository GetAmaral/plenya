#!/bin/sh
set -e

# Entrypoint de produção. As migrations goose são aplicadas como passo do deploy,
# ANTES de a aplicação subir (nunca via AutoMigrate no boot).
#
# Gate de segurança: só roda migrations quando RUN_MIGRATIONS=true estiver setado
# no ambiente (Coolify). Default (variável ausente) = NÃO roda — comportamento
# idêntico ao anterior (só inicia o server). Isso permite preparar a imagem sem
# disparar migrations em prod automaticamente; ative o gate no cutover, depois de
# conferir drift do schema e tirar backup. Ver docs/emr/migrations-decisao.md.
if [ "${RUN_MIGRATIONS:-false}" = "true" ]; then
  echo "🗄️  RUN_MIGRATIONS=true — aplicando migrations (goose)..."
  ./migrate up
else
  echo "⏭️  RUN_MIGRATIONS!=true — migrations não aplicadas neste boot."
fi

echo "🚀 Starting server..."
exec ./server
