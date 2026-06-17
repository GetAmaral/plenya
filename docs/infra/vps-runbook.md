# VPS — Runbook de operação e tuning (produção)

> Host KingHost 8GB (Ubuntu 24.04), Docker + Coolify. Acesso: `ssh plenya` (usuário `deploy`,
> sudo NOPASSWD). **Credenciais/tokens NÃO ficam aqui** — vivem em `~/.plenya-vps-secrets/` na
> máquina do operador e nas memórias `plenya_vps*`. Este doc é procedimento, não segredo.

## Leitura de memória (o "free baixo" engana)

`free -h` mostrando ~1GB "free" **não** é escassez: o Linux usa RAM ociosa como cache de disco
reclaimável. O número que importa é **`available`** (`free` coluna 7) — saudável em ~5GB de 7,8GB.
Sob pico (ex.: vários Chromium do gerador de PDF nascendo juntos), o gargalo é alocação anônima
rápida competindo com swap quase cheia, não falta de RAM em repouso.

## Tuning aplicado (2026-06-17)

### swappiness 60 → 10
A swap vivia ~76% cheia com swappiness padrão (60), empurrando páginas quentes pra disco lento.
```bash
echo 'vm.swappiness=10' | sudo tee /etc/sysctl.d/99-plenya-swappiness.conf
sudo sysctl -w vm.swappiness=10
sudo swapoff -a && sudo swapon -a   # de-swap imediato (precisa de `available` > swap usada)
```
Faixa de produção é 1–10 (Red Hat recomenda 10 p/ hosts de banco). Persiste em sysctl.d.

### daemon.json: live-restore + GC do BuildKit
`/etc/docker/daemon.json` (backup automático em `daemon.json.bak-*` antes de editar):
```json
{
  "log-driver": "json-file",
  "log-opts": { "max-size": "10m", "max-file": "3" },
  "default-address-pools": [ { "base": "10.0.0.0/8", "size": 24 } ],
  "live-restore": true,
  "builder": { "gc": { "enabled": true, "defaultKeepStorage": "10GB" } }
}
```
- **`live-restore`**: containers seguem rodando se o daemon cair/reiniciar. Sem isso, restart do
  docker derruba tudo.
- **`builder.gc`**: o GC do BuildKit vem **desligado** por default — foi por isso que o cache de
  build chegou a 16GB. Com `defaultKeepStorage: 10GB` ele se auto-poda e não volta a inchar.

## Por que o dockerd incha (700–800MB RSS) e como resetar sem downtime

O RSS do dockerd cresce por design do runtime Go: goroutines de `docker exec` frequente
(healthchecks + `coolify-sentinel`) alocam memória que o Go nunca devolve ao SO. **`prune` não
encolhe** isso — só um restart do daemon zera (e volta a crescer com o tempo).

🔑 **Reset sem derrubar containers** (provado: 764MB → 122MB, zero downtime):
```bash
# 1) garantir "live-restore": true no daemon.json (ver acima)
sudo systemctl reload docker                              # SIGHUP ativa live-restore SEM bounce
docker info --format '{{.LiveRestoreEnabled}}'            # precisa imprimir: true
sudo systemctl restart docker                             # com live-restore ON, containers seguem; daemon reseta
# validar que NÃO caíram (StartedAt não muda):
docker inspect <container> --format '{{.State.StartedAt}}'
```
Sem o `reload` antes, o `restart` bounce os containers. Depois que live-restore está ON, todo
restart futuro é indolor — dá pra resetar o dockerd periodicamente sem janela.

> Caveat live-restore: só vira problema se o daemon ficar **minutos** fora (buffer de log FIFO
> enche e bloqueia escrita). Restart é segundos → não se aplica.

## Limpeza de disco (Docker)
```bash
docker builder prune -af      # cache de build (recuperou ~16GB)
docker image prune -af        # imagens órfãs (não usadas por container)
docker system df              # panorama (Images/Containers/Build Cache)
```
Com o GC do BuildKit ligado (acima), o `builder prune` manual vira raro.

## Apps Next — pendência de RAM (Frente B)
Os 3 apps Next (`web`, `site`, `site-getulio`) sobem com `CMD ["pnpm","start"]`, que mantém um
processo `node /usr/local/bin/pnpm start` parasita de ~98MB **além** do `next-server` (~200–290MB
desperdiçados no total). O padrão de produção é `output: 'standalone'` no `next.config` → roda
`node server.js` direto (sem wrapper) e encolhe a imagem. Cuidados no monorepo:
- `outputFileTracingRoot` apontando pra raiz do repo (senão o tracing de deps quebra);
- copiar `.next/static` e `public` pra dentro da árvore `standalone/`;
- conteúdo lido em runtime via `process.cwd()` (ex.: `apps/site-getulio/lib/blog.ts`) precisa ser
  copiado pro caminho certo relativo ao novo cwd;
- **buildar a imagem local + smoke test (`/`, listagens, slug, imagem) ANTES de deployar.**
O `api` (Go) já é exemplar (~15MB) — não se aplica.

## Saúde rápida
```bash
curl -s -o /dev/null -w '%{http_code} %{time_total}s\n' https://api.plenyasaude.com.br/health
curl -s -o /dev/null -w '%{http_code} %{time_total}s\n' https://plenyasaude.com.br
ssh plenya "docker ps --format '{{.Names}} {{.Status}}'; free -h; docker system df"
```

## Coolify — API (operação)
Token Sanctum (`7|…`) em `~/.plenya-vps-secrets/coolify.env` var `COOLIFY_API_TOKEN`.
🚨 **Nunca `source` o arquivo** — o `|` do token é pipe pro bash e zera a variável. Extrair cru:
```bash
TOKEN=$(grep '^COOLIFY_API_TOKEN=' ~/.plenya-vps-secrets/coolify.env | cut -d= -f2-)
BASE=http://191.252.159.48:8000   # ou https://coolify.plenyasaude.com.br
curl -s -H "Authorization: Bearer $TOKEN" $BASE/api/v1/version          # sanity → 200
```
- `GET /api/v1/deployments` lista **só os em andamento** (vazio = nenhum, não histórico).
- Status/commit por app: `GET /api/v1/applications/{uuid}`. Disparar deploy: `GET /api/v1/deploy?uuid={uuid}`.
- `concurrent_builds=1` no server_settings (fix de raiz do OOM: builds serializados, um por vez).
- App UUIDs = prefixo do nome do container: api `kgcuxgvmnbx6yya35e3ca2v0` · web `nwbhak0fscs2th13gz5g9zjm`
  · site `ycpklto5n1qjkmelhdp0pvhf` · site-getulio `qkdzqaauicc001qfkghfur0s` · postgres `mb511beqjtgd7nsjlnngh3m6`.
