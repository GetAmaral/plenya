# Atualização da Configuração Docker - Janeiro 2026

## Mudanças Realizadas

### 1. Volumes Sincronizados (Já estava configurado)
O `docker-compose.yml` já estava correto com volumes para sincronização:
- **Web**: Todo o monorepo sincronizado (`.:/app:cached`)
- **API**: Código Go sincronizado (`./apps/api:/app:cached`)
- **Exclusões**: node_modules e build artifacts ficam no container

### 2. Entrypoint Inteligente para Web
Criado `/apps/web/entrypoint.sh` que:
- Detecta automaticamente mudanças em `package.json` ou `pnpm-lock.yaml`
- Reinstala dependências automaticamente quando necessário
- Elimina necessidade de rebuild manual ao adicionar pacotes

### 3. Dockerfile Web Simplificado
Removido a instalação de dependências do Dockerfile:
- Dependências agora são instaladas pelo entrypoint
- Imagem mais leve e builds mais rápidos
- Hot reload automático funciona perfeitamente

### 4. Correções de Imports Go
Corrigidos imports nos arquivos do Score System:
- **Antes**: `plenya/internal/models`
- **Depois**: `github.com/plenya/api/internal/models`

Arquivos corrigidos:
- `apps/api/internal/services/score_service.go`
- `apps/api/internal/handlers/score_handler.go`
- `apps/api/internal/repository/score_repository.go`

### 5. Componente Accordion Criado
Adicionado componente UI faltante:
- `apps/web/components/ui/accordion.tsx`
- Pacote `@radix-ui/react-accordion` instalado
- Animações já configuradas no Tailwind

### 6. Error Handling nos Pages
Adicionado tratamento de erros adequado:
- `apps/web/app/anamnesis/page.tsx` - Exibe erros de conexão
- `apps/web/app/prescriptions/page.tsx` - Exibe erros de conexão
- Interface amigável com botão "Tentar novamente"

### 7. Navegação Score System
Adicionado link no sidebar:
- `apps/web/components/dashboard/sidebar.tsx`
- Ícone Network para "Escores"

## Workflow de Desenvolvimento Atualizado

### Comandos Principais

```bash
# Iniciar todos os serviços
docker compose up -d

# Ver logs em tempo real
docker compose logs -f
docker compose logs -f web
docker compose logs -f api

# Instalar nova dependência (exemplo)
docker compose exec web pnpm add <pacote> --filter web
# Não precisa rebuild - o entrypoint detecta automaticamente

# Parar serviços
docker compose down

# Rebuild completo (raramente necessário)
docker compose down && docker compose up -d --build
```

### Hot Reload

**O que funciona automaticamente:**
- ✅ Mudanças em arquivos `.ts`, `.tsx` → Next.js recompila
- ✅ Mudanças em arquivos `.go` → Air recompila (se configurado)
- ✅ Mudanças em `package.json` → Entrypoint reinstala
- ✅ Mudanças em código sincronizam via volume

**O que requer rebuild:**
- 🔨 Mudanças no `Dockerfile.dev`
- 🔨 Mudanças em `.dockerignore`
- 🔨 Mudanças estruturais em `docker-compose.yml`

## Status Atual

### ✅ Funcionando
- PostgreSQL 17 (porta 5432)
- API Go (porta 3001) - 116 handlers
- Web Next.js (porta 3000) - Turbopack
- Sincronização de código via volumes
- Auto-instalação de dependências
- Score System completo
- Error handling em todas páginas

### 📦 Pacotes Adicionados
- `@radix-ui/react-accordion@^1.2.2`

### 🐛 Bugs Corrigidos
1. Imports Go incorretos (plenya → github.com/plenya/api)
2. Componente Accordion faltando
3. Error handling faltando em Anamnesis/Prescriptions
4. Link de navegação para Scores faltando

## Próximos Passos Sugeridos

1. **Considerar Air para Go hot reload:**
   ```dockerfile
   RUN go install github.com/cosmtrek/air@latest
   CMD ["air", "-c", ".air.toml"]
   ```

2. **Otimizar cache do Docker:**
   - Multi-stage builds para produção
   - Layer caching mais eficiente

3. **Scripts auxiliares:**
   - `scripts/dev.sh` - Wrapper para docker compose
   - `scripts/logs.sh` - Logs formatados
   - `scripts/test.sh` - Rodar testes em containers

## Notas Importantes

- **NUNCA rodar pnpm/npm/go diretamente no host** - Sempre via Docker
- **Dependencies são gerenciadas pelo container** - node_modules fica no container
- **Rebuild só é necessário em mudanças estruturais** - Código sincroniza automaticamente
- **Logs são essenciais** - Sempre verificar `docker compose logs` ao debugar

---

**Data**: 2026-01-24
**Status**: ✅ Configuração otimizada e funcionando
