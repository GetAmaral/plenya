# Dev Bypass Auth Mode

## Visão Geral

Modo de desenvolvimento que bypassa completamente autenticação JWT, permitindo acesso direto ao sistema como `admin@plenya.com`. Útil para Claude Code e desenvolvimento rápido sem necessidade de login.

## ⚠️ Segurança

**CRÍTICO:**
- ✅ Usar APENAS em ambiente de desenvolvimento
- ❌ NUNCA ativar em produção
- ❌ NUNCA commitar `.env` com bypass=true
- ⚠️ Backend loga avisos visíveis quando ativo

## Como Funciona

### Backend (Go)

1. **Config** carrega `DEV_BYPASS_AUTH` do `.env`
2. **Main** após DB init, carrega dados reais do `admin@plenya.com`:
   - UUID do usuário
   - Email
   - Roles (admin, doctor, manager, etc.)
3. **Auth Middleware** injeta dados do admin no contexto **antes** de validar JWT:
   ```go
   if cfg.Dev.BypassAuth {
       c.Locals("userID", cfg.Dev.AdminUserID)
       c.Locals("userEmail", cfg.Dev.AdminEmail)
       c.Locals("userRoles", cfg.Dev.AdminRoles)
       return c.Next()
   }
   ```

### Frontend (Next.js)

1. **Auth Store** inicializa com placeholder admin se `NEXT_PUBLIC_DEV_BYPASS_AUTH=true`:
   ```typescript
   user: {
     id: 'dev-admin-placeholder',
     email: 'admin@plenya.com',
     name: 'Dev Admin',
     roles: ['admin'],
   }
   ```
2. **useRequireAuth** pula verificação e redirecionamento para login
3. Chamadas API enviam token dummy `dev-bypass-token`
4. Backend ignora token e injeta admin real

## Ativação

### 1. Backend

```bash
# apps/api/.env
DEV_BYPASS_AUTH=true
```

### 2. Frontend

```bash
# apps/web/.env.local
NEXT_PUBLIC_DEV_BYPASS_AUTH=true
```

### 3. Restart

```bash
docker compose restart api web
```

### Logs Esperados

Backend ao iniciar:
```
⚠️  DEV_BYPASS_AUTH habilitado — autenticação desativada!
⚠️  DEV bypass ativo como: admin@plenya.com (019bf2f3-c528-70cc-8f81-6741cf3fad27)
```

## Desativação

```bash
# apps/api/.env
DEV_BYPASS_AUTH=false

# apps/web/.env.local
NEXT_PUBLIC_DEV_BYPASS_AUTH=false

# Restart
docker compose restart api web
```

## Testes

### Verificar Ativação

```bash
# 1. Backend logs
docker compose logs api | grep DEV_BYPASS_AUTH

# 2. Endpoint protegido SEM token
curl http://localhost:3001/api/v1/users/me
# ✅ Deve retornar dados do admin@plenya.com

# 3. Frontend direto
open http://localhost:3000
# ✅ Deve entrar sem pedir login
```

### Verificar Desativação

```bash
# Endpoint protegido SEM token
curl http://localhost:3001/api/v1/users/me
# ❌ Deve retornar 401 Unauthorized
```

## Arquivos Modificados

### Backend
- `apps/api/internal/config/config.go` - DevConfig struct
- `apps/api/internal/middleware/auth.go` - Bypass antes de validar JWT
- `apps/api/cmd/server/main.go` - Inicialização após DB

### Frontend
- `apps/web/lib/auth-store.ts` - Placeholder admin no estado inicial
- `apps/web/lib/use-auth.ts` - Skip de verificação

### Environment
- `apps/api/.env` - DEV_BYPASS_AUTH
- `apps/web/.env.local` - NEXT_PUBLIC_DEV_BYPASS_AUTH

## Casos de Uso

### ✅ Quando Usar

- Claude Code fazendo operações automáticas
- Desenvolvimento rápido sem login manual
- Testes de features sem setup de autenticação
- Debug de problemas não relacionados a auth

### ❌ Quando NÃO Usar

- Testes de autenticação/autorização
- Validação de fluxos de login/logout
- Verificação de permissões por role
- Qualquer ambiente que não seja local dev

## Troubleshooting

### Bypass não funciona

**Sintomas:**
- Backend não loga avisos de bypass
- Frontend redireciona para login
- API retorna 401

**Verificar:**
```bash
# 1. Variáveis de ambiente
grep DEV_BYPASS_AUTH apps/api/.env
grep NEXT_PUBLIC_DEV_BYPASS_AUTH apps/web/.env.local

# 2. Containers reiniciados após mudar .env
docker compose ps

# 3. Logs de inicialização
docker compose logs api | head -50
```

### Admin não encontrado

**Erro:**
```
DEV_BYPASS_AUTH requer usuário admin@plenya.com no banco
```

**Solução:**
```bash
# Verificar se admin existe
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT id, email, roles FROM users WHERE email = 'admin@plenya.com';"

# Se não existir, criar via seed/migration
```

### Frontend ainda pede login

**Verificar:**
1. `.env.local` tem `NEXT_PUBLIC_DEV_BYPASS_AUTH=true`
2. Variável é `NEXT_PUBLIC_*` (exposta ao browser)
3. Container web foi reiniciado após alterar `.env.local`
4. Limpar localStorage do browser: `localStorage.clear()`

## Notas de Implementação

### Por que não usar sempre?

- Mascara bugs de autenticação/autorização
- Pode criar dependência de bypass em testes
- Não valida fluxos reais de login
- Pode deixar código sem proteção adequada

### Alternativas

Para testes automatizados, considerar:
- Mock de JWT em testes unitários
- Fixtures de usuários em testes E2E
- Helper de login automatizado em Cypress/Playwright

### Extensões Futuras

Possíveis melhorias:
- Permitir escolher usuário via `DEV_BYPASS_USER_EMAIL`
- Modo debug que loga todas as verificações de auth
- Endpoint `/dev/impersonate/:userId` (apenas dev)

---

## 🚨 A família de bug que o bypass ESCONDE, medida em produção

"Mascara bugs de autenticação" acima é abstrato demais. Em 2026-09-06 isso custou um defeito ativo
em produção, e o padrão é sempre o mesmo:

> **Baixar arquivo autenticado com `window.open(url)` ou `<a href={url} download>`.**

Nenhum dos dois manda cabeçalho, e `middleware.Auth` lê **só** `Authorization: Bearer` (não há
cookie de sessão). Em produção: **401**. Em dev: **funciona**, porque o bypass não exige token.

Foi o caso do download de PDF de artigo, em três pontos de chamada, incluindo um `FileViewer` que
repassava a mesma URL crua para o visualizador de PDF, para o de EPUB e para um link de download —
três caminhos quebrados por uma prop só. Ninguém percebeu porque o ambiente onde se testa é
justamente o que esconde.

**A forma certa** já existia e estava comentada em `openPrescriptionPdf`:

```ts
const token = useAuthStore.getState().accessToken
const res = await fetch(url, { headers: token ? { Authorization: `Bearer ${token}` } : {} })
if (!res.ok) throw new Error('…')
const blobUrl = URL.createObjectURL(await res.blob())
window.open(blobUrl, '_blank')          // ou passar blobUrl ao visualizador
setTimeout(() => URL.revokeObjectURL(blobUrl), 60_000)
```

**Ao mexer em qualquer download autenticado:**

1. Procure `window.open(` e `href={...download}` apontando para `/api/v1/` — são suspeitos por
   construção.
2. **Teste contra produção**, não contra a dev: `curl -o /dev/null -w '%{http_code}' <url>` sem
   token. 401 é o comportamento correto da rota e a prova de que `window.open` não serve.
3. Rotas que servem arquivo mas ninguém chama, e ids de artefato no DTO que a tela ignora, são a
   mesma família: veja a varredura no commit `d20e6e37`.
