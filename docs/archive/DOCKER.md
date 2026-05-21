# Docker Development Setup - Plenya EMR

## Sincronização Automática de Arquivos

O Docker está configurado para **sincronização automática** entre o filesystem do host e os containers.

### Como Funciona

#### Container Web (Frontend)

**Volume principal:**
```yaml
- .:/app:cached  # Todo o monorepo sincronizado
```

**Exclusões (mantidas no container):**
```yaml
- /app/node_modules                      # Root
- /app/apps/web/node_modules             # Web app
- /app/apps/api/node_modules             # API (se existir)
- /app/packages/*/node_modules           # Packages
- /app/apps/web/.next                    # Build cache
- /app/.turbo                            # Turborepo cache
- /app/.git                              # Git history
```

**Variáveis de ambiente:**
```yaml
- WATCHPACK_POLLING=true      # Webpack/Turbopack file watching
- CHOKIDAR_USEPOLLING=true    # Next.js hot reload
```

#### Container API (Backend Go)

**Volume principal:**
```yaml
- ./apps/api:/app:cached  # Código Go sincronizado
```

**Exclusões:**
```yaml
- /app/bin  # Binários compilados
```

---

## Comportamento de Sincronização

### ✅ Sincronizado Automaticamente (Hot Reload)

**Frontend:**
- ✅ `apps/web/**/*.tsx` → Next.js detecta e recarrega
- ✅ `apps/web/**/*.ts` → TypeScript recompila
- ✅ `apps/web/app/**/*` → Rotas atualizadas
- ✅ `apps/web/components/**/*` → Componentes recarregados
- ✅ `packages/**/*` → Shared code atualizado
- ✅ `apps/web/tailwind.config.ts` → CSS regenerado
- ✅ `apps/web/next.config.ts` → Requer restart manual

**Backend:**
- ✅ `apps/api/**/*.go` → Air detecta e recompila (hot reload)
- ✅ `apps/api/internal/**/*.go` → Recompilação automática
- ✅ `apps/api/.env` → Requer restart manual

### ⚠️ Requer Ação Manual

**Frontend:**
- `package.json` → Rodar: `docker exec plenya-web pnpm install`
- `pnpm-lock.yaml` → Rodar: `docker exec plenya-web pnpm install`
- `next.config.ts` → Restart: `docker restart plenya-web`

**Backend:**
- `go.mod` / `go.sum` → Restart: `docker restart plenya-api`
- Novas migrations → Aplicar manualmente

**Ambos:**
- `docker-compose.yml` → Recrear: `docker compose up -d --force-recreate`
- `Dockerfile.dev` → Rebuild: `docker compose up -d --build`

---

## Comandos Úteis

### Desenvolvimento Normal

```bash
# Iniciar todos os serviços
docker compose up -d

# Ver logs em tempo real
docker compose logs -f web    # Frontend
docker compose logs -f api    # Backend
docker compose logs -f db     # Database

# Verificar status
docker compose ps

# Parar todos
docker compose down
```

### Instalar Dependências

```bash
# Frontend (após adicionar package no package.json)
docker exec plenya-web pnpm install

# Ou instalar um pacote específico
docker exec plenya-web pnpm add <package>

# Backend Go (após mudar go.mod)
docker exec plenya-api go mod tidy
```

### Reiniciar Serviços

```bash
# Reiniciar um serviço específico
docker restart plenya-web
docker restart plenya-api

# Reiniciar todos
docker compose restart

# Rebuild completo (após mudanças em Dockerfile)
docker compose up -d --build
```

### Acessar Shells

```bash
# Shell no container web
docker exec -it plenya-web sh

# Shell no container API
docker exec -it plenya-api sh

# Shell no PostgreSQL
docker exec -it plenya-db psql -U plenya_user -d plenya_db
```

### Limpeza

```bash
# Remover containers e networks (mantém volumes)
docker compose down

# Remover TUDO (incluindo volumes/database)
docker compose down -v

# Limpar cache do Docker
docker system prune -a
```

---

## Performance de Sincronização

### Modo `:cached`

Usamos o flag `:cached` nos volumes:
```yaml
- .:/app:cached
```

**Significado:**
- **Consistência eventual** entre host e container
- **Writes no host** são propagados para o container (pode ter delay de ~100ms)
- **Reads no container** podem usar cache
- **Performance:** ~30-50% mais rápido que sync padrão

**Alternativas:**
- `:consistent` → Sync bidirecional imediato (mais lento)
- `:delegated` → Container escreve, host lê (inverso do cached)

### File Watching

**Polling habilitado:**
```yaml
WATCHPACK_POLLING=true
CHOKIDAR_USEPOLLING=true
```

**Por quê?**
- Inotify (file watching nativo) não funciona através do Docker volume mount
- Polling verifica mudanças a cada ~500ms
- Usa ~5-10% mais CPU, mas garante hot reload

**Desabilitar (se CPU for problema):**
```yaml
# Em docker-compose.yml, remover:
# - WATCHPACK_POLLING=true
# - CHOKIDAR_USEPOLLING=true
```

---

## Troubleshooting

### Hot Reload Não Funciona

**Sintoma:** Arquivo muda, mas Next.js não recarrega.

**Soluções:**
```bash
# 1. Verificar se polling está ativado
docker exec plenya-web env | grep POLLING

# 2. Verificar permissões de arquivo
ls -la apps/web/

# 3. Restart container
docker restart plenya-web

# 4. Rebuild completo
docker compose up -d --build web
```

### Mudanças em package.json Não Aplicadas

**Sintoma:** Instalou pacote mas não funciona.

**Solução:**
```bash
# Reinstalar dependências no container
docker exec plenya-web pnpm install

# Restart para garantir
docker restart plenya-web
```

### Container Fica Reiniciando

**Sintoma:** `docker ps` mostra "Restarting".

**Investigar:**
```bash
# Ver logs de erro
docker logs plenya-web

# Possíveis causas:
# - Erro de sintaxe em código
# - Porta já em uso
# - Dependência faltando
```

### Performance Muito Lenta

**Sintomas:** Edição de arquivo trava, hot reload demora 10s+.

**Soluções:**
```bash
# 1. Verificar uso de disco (Docker Desktop)
# Docker → Settings → Resources → Disk

# 2. Limpar cache do Next.js
rm -rf apps/web/.next
docker restart plenya-web

# 3. Reduzir watchers (se muitos arquivos)
# Adicionar em .gitignore arquivos temporários

# 4. Usar desenvolvimento local (sem Docker)
cd apps/web
pnpm dev  # ~3-5x mais rápido
```

### Volumes Ficam Dessincronizados

**Sintoma:** Mudança no host não aparece no container.

**Soluções:**
```bash
# 1. Recrear containers
docker compose up -d --force-recreate

# 2. Verificar volume mount
docker inspect plenya-web | grep Mounts -A 20

# 3. Último recurso: Rebuild
docker compose down
docker compose up -d --build
```

---

## Arquitetura de Volumes

```
Host (seu computador)
  /home/user/plenya/
    ├── apps/web/          → Container: /app/apps/web/
    ├── apps/api/          → Container: /app/
    ├── packages/          → Container: /app/packages/
    ├── node_modules/      ✗ NÃO sincronizado (container mantém o próprio)
    └── .git/              ✗ NÃO sincronizado (ignorado)

Container (plenya-web)
  /app/
    ├── apps/web/          ← Seu código (sincronizado)
    ├── packages/          ← Seu código (sincronizado)
    ├── node_modules/      ← Dependências (isoladas)
    └── .next/             ← Build cache (isolado)
```

**Vantagem:** Dependências corretas no container, sem conflitos com host.

---

## Desenvolvimento Local (Sem Docker)

Para **máxima velocidade** (não recomendado para produção):

```bash
# Terminal 1: Database
docker compose up -d db

# Terminal 2: Backend
cd apps/api
go run cmd/server/main.go

# Terminal 3: Frontend
cd apps/web
pnpm dev

# ~3-5x mais rápido que Docker
# Mas perde isolamento e reprodutibilidade
```

---

## Best Practices

### ✅ Fazer

1. **Sempre commitar** antes de testar em Docker
2. **Usar .dockerignore** para excluir arquivos desnecessários
3. **Manter node_modules isolados** (não compartilhar com host)
4. **Rodar pnpm install** no container após mudanças em package.json
5. **Usar docker compose logs** para debug

### ❌ Não Fazer

1. **Não editar node_modules manualmente** (usar pnpm)
2. **Não compartilhar .next entre host/container** (causa conflitos)
3. **Não usar Docker Desktop em máquinas lentas** (<8GB RAM)
4. **Não ignorar avisos de rebuild** (causa comportamento estranho)
5. **Não fazer build de produção no container de dev**

---

## Monitoramento de Mudanças

### Script Helper (Opcional)

Criado em: `/scripts/watch-deps.sh`

**Uso:**
```bash
# Monitorar mudanças em package.json automaticamente
docker exec plenya-web /app/scripts/watch-deps.sh

# Reinstala dependências automaticamente quando package.json muda
```

**Não recomendado para uso contínuo** (usa CPU extra).

---

## FAQ

**Q: Posso editar arquivos direto no container?**
A: Sim, mas mudanças serão perdidas no restart. Sempre edite no host.

**Q: Por que node_modules não sincroniza?**
A: Para evitar conflitos. Node modules são específicos do OS (Linux no container).

**Q: Hot reload parou de funcionar.**
A: Restart o container: `docker restart plenya-web`

**Q: Posso usar WSL2 com Docker?**
A: Sim! Recomendado no Windows. Muito mais rápido que Docker Desktop tradicional.

**Q: Como testar performance de produção?**
A: `cd apps/web && pnpm build && pnpm start` (fora do Docker)

---

**Última atualização:** Janeiro 2026
