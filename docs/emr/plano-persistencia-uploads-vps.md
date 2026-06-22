# Persistência de uploads do EMR (VPS / Coolify) — diagnóstico e solução definitiva

**Data:** 2026-06-22 · **Severidade:** crítica (perda silenciosa de dados clínicos)

## Sintoma que levou ao bug
Áudio da lead "Lucia" (`/conversas`) parou de tocar após um redeploy do `plenya-api`: o
endpoint de mídia passou a devolver 404 e o front escondeu até a transcrição.

## Causa raiz
`apps/api/Dockerfile:51` declara `VOLUME ["/app/uploads", "/app/enrichment-batch"]`. O Coolify
**não tem persistência configurada** para o `plenya-api` (o `docker-compose.yaml` gerado não tem
seção `volumes:`). Sem um mount explícito, o Docker cria um **volume anônimo novo a cada deploy**.

Consequência: a cada redeploy do api, **todo `/app/uploads` é órfãnado** — o banco mantém a
`media_storage_key` / `file_path`, mas o arquivo fica num volume abandonado. Atinge:
- mídia de conversa do WhatsApp (`whatsapp-media/...`) — áudios, fotos
- **documentos de prontuário (`patient-docs/<patientID>/...`) — PDFs/exames dos pacientes**

Encontrados 3 volumes anônimos órfãos de deploys anteriores, cada um com mídia de um período.
O volume montado no deploy atual estava praticamente vazio.

## Resgate imediato (já feito em 2026-06-22)
Consolidei os 4 arquivos sobreviventes dos volumes órfãos no volume montado atual, restaurando
o áudio da Lucia. **É paliativo:** o próximo redeploy órfã tudo de novo.

## Solução definitiva (a executar)

Princípio: dado clínico não pode depender de volume anônimo do Docker. Três camadas.

### 1. Tirar a armadilha do Dockerfile
Remover a linha `VOLUME [...]` de `apps/api/Dockerfile` (mantendo o `mkdir -p`). Sem o `VOLUME`,
o Docker nunca mais cria volume anônimo silencioso para esses paths — defense in depth, caso a
config do Coolify um dia derive.

### 2. Persistência via bind mount no Coolify (não volume nomeado)
Configurar persistent storage do `plenya-api` como **bind mount** para o caminho-padrão do
Coolify, inserindo na tabela `local_persistent_volumes` (fonte de verdade que o Coolify lê ao
gerar o compose — mesma coisa que a UI faz):

| mount_path (container)  | host_path (bind)                                                   |
|-------------------------|--------------------------------------------------------------------|
| `/app/uploads`          | `/data/coolify/applications/kgcuxgvmnbx6yya35e3ca2v0/uploads`       |
| `/app/enrichment-batch` | `/data/coolify/applications/kgcuxgvmnbx6yya35e3ca2v0/enrichment-batch` |

`resource_id=2`, `resource_type='App\Models\Application'`, `host_path` preenchido (= bind),
`is_preview_suffix_enabled=false`.

**Por que bind mount e não volume nomeado** (como o Postgres usa): para dado que dependemos muito,
o bind ganha em quase tudo —
- imune a `docker volume prune` / `system prune --volumes` / "delete volumes" do Coolify;
- backup trivial (rsync/tar de um diretório, sem docker);
- inspecionável (`ls` direto no host) e portável (copiar a pasta para outro servidor).
O volume nomeado sobreviveria a redeploys, mas é mais frágil a operações de limpeza e menos
transparente. O caminho `/data/coolify/applications/<uuid>/` é a convenção do próprio Coolify e
o diretório já existe desde o setup inicial.

### 3. Migração dos dados + cutover
1. Copiar o conteúdo consolidado (volume montado atual) para o host_path do bind.
2. Disparar **um** deploy do api (limpando fila/órfãos antes — ver [[coolify_deploy_orphan_lock_procedure]]).
3. Confirmar no `docker inspect` que `/app/uploads` agora é `bind` para o host_path (não mais volume).
4. Conferir que a mídia da Lucia volta 200 e toca.
5. Só depois: remover os volumes anônimos órfãos (`docker volume rm`).

### 4. Backup do uploads (rede de segurança real)
O bind protege contra churn de container, mas não contra falha de disco. O DB já é dumpado para
`~/.plenya-vps-secrets/backups/`. Adicionar backup periódico do `uploads` (tar/rsync do host_path
para o mesmo destino de backup, via cron). Sem isto, "persistente" ainda é single-point-of-failure.

## Fix de frontend acoplado
`apps/web/components/conversations/conversation-viewer.tsx` (`WhatsAppMediaView`): quando a mídia
falha (`err`), o componente faz early-return e **esconde a transcrição**. A transcrição deve
aparecer sempre — é o fallback quando o áudio não carrega/toca. Reordenar para mostrar o aviso de
falha E a transcrição.

## Escopo
Só o `plenya-api` tem `VOLUME`/uploads. `plenya-web` e `plenya-site` são stateless (sem mounts).
