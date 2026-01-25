# Performance Guide - Plenya EMR

## Otimizações Implementadas

### 1. Loading Global (NProgress)

**Localização:** `/components/navigation-progress.tsx`

- Barra de progresso azul no topo durante navegação
- Cursor `wait` automático durante carregamento
- Feedback visual imediato ao usuário

### 2. Next.js Config Otimizado

**Arquivo:** `next.config.ts`

#### Otimizações Ativas:

```typescript
// Imagens otimizadas (AVIF + WebP)
images: {
  formats: ['image/avif', 'image/webp']
}

// Tree-shaking de imports grandes
experimental: {
  optimizePackageImports: [
    'lucide-react',
    'recharts',
    '@radix-ui/*'
  ]
}

// Remove console.log em produção
compiler: {
  removeConsole: process.env.NODE_ENV === 'production'
}
```

### 3. Prevenção de Duplo Submit

**CSS Global:** Botões de submit ficam automaticamente `opacity: 0.6` e `cursor: not-allowed` quando disabled.

**Padrão de código:**
```tsx
<Button
  type="submit"
  disabled={mutation.isPending}
>
  {mutation.isPending ? "Salvando..." : "Salvar"}
</Button>
```

### 4. React Query Otimizado

**Localização:** `/lib/query-provider.tsx`

Configurações:
- Stale time: 5 minutos (reduz re-fetches)
- Cache time: 10 minutos
- Retry: 2 tentativas com backoff exponencial

---

## Métricas de Performance

### Targets (Dev Mode)

- **First Load:** < 2s
- **Navigation:** < 1s
- **Form Submit:** < 500ms (response do backend)
- **Interactive:** < 3s

### Targets (Produção)

- **First Load:** < 1s
- **Navigation:** < 500ms
- **Form Submit:** < 300ms
- **Interactive:** < 2s
- **Lighthouse Score:** > 90

---

## Causas Comuns de Lentidão

### Dev Mode vs Produção

**Dev mode é 3-5x mais lento que produção.** Isso é normal:

- ✅ **Dev:** Hot reload, source maps, sem minificação
- ✅ **Prod:** Minificado, bundled, cached

**Para testar performance real:**
```bash
cd apps/web
pnpm build
pnpm start
```

### Turbopack (Experimental)

Atualmente usando `next dev --turbopack`. Se estiver muito lento:

```bash
# Desabilitar Turbopack (voltar para Webpack)
# Em package.json:
"dev": "next dev"  # sem --turbopack
```

### Docker Overhead

Containers Docker adicionam ~10-20% de overhead. Para desenvolvimento mais rápido:

```bash
# Desenvolvimento local (sem Docker)
cd apps/web
pnpm install
pnpm dev
```

---

## Checklist de Performance

### Frontend

- [x] NProgress instalado e configurado
- [x] Cursor wait durante navegação
- [x] Botões disabled durante submit
- [x] Next.js config otimizado
- [x] React Query com cache inteligente
- [ ] Code splitting por rota (Next.js faz automaticamente)
- [ ] Lazy loading de componentes pesados
- [ ] Memoização de cálculos caros (useMemo/useCallback)

### Backend

- [ ] Índices em colunas de busca frequente
- [ ] Paginação em todas as listas
- [ ] Eager loading de relações (evitar N+1)
- [ ] Response compression (gzip/brotli)
- [ ] CDN para assets estáticos

### Database

- [x] Índices em foreign keys
- [x] Soft delete com índices
- [ ] EXPLAIN ANALYZE em queries lentas
- [ ] Vacuum/Analyze periódico
- [ ] Connection pooling

---

## Debugging de Performance

### 1. Chrome DevTools

```bash
# Abrir DevTools
F12 → Performance tab → Record

# Métricas importantes:
- LCP (Largest Contentful Paint): < 2.5s
- FID (First Input Delay): < 100ms
- CLS (Cumulative Layout Shift): < 0.1
```

### 2. React DevTools Profiler

```bash
# Instalar extensão React DevTools
# Aba Profiler → Record → Interagir → Stop

# Procurar por:
- Componentes com render > 100ms
- Re-renders desnecessários
- Updates em cascata
```

### 3. Next.js Bundle Analyzer

```bash
cd apps/web
pnpm add @next/bundle-analyzer

# Em next.config.ts:
import bundleAnalyzer from '@next/bundle-analyzer'

const withBundleAnalyzer = bundleAnalyzer({
  enabled: process.env.ANALYZE === 'true'
})

export default withBundleAnalyzer(nextConfig)

# Rodar análise:
ANALYZE=true pnpm build
```

### 4. Backend Profiling

```bash
# Logs de queries lentas (GORM)
# Em config/database.go:
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
})

# Procurar queries com > 100ms nos logs
```

---

## Quick Wins

### Frontend

```tsx
// 1. Lazy load componentes pesados
const HeavyChart = lazy(() => import('./HeavyChart'))

// 2. Memoizar cálculos
const expensiveValue = useMemo(() =>
  data.filter(/* ... */).map(/* ... */),
  [data]
)

// 3. Debounce de inputs
const debouncedSearch = useDeferredValue(searchTerm)

// 4. Virtual scrolling para listas grandes
import { useVirtualizer } from '@tanstack/react-virtual'
```

### Backend

```go
// 1. Eager loading
db.Preload("Patient").Preload("Doctor").Find(&appointments)

// 2. Paginação
db.Limit(limit).Offset(offset).Find(&patients)

// 3. Índices compostos
type Patient struct {
    UserID uuid.UUID `gorm:"index:idx_user_created,priority:1"`
    CreatedAt time.Time `gorm:"index:idx_user_created,priority:2"`
}

// 4. Cache de queries frequentes
// Usar Redis ou in-memory cache
```

---

## Monitoramento Contínuo

### Ferramentas Recomendadas (Futuro)

1. **Sentry** - Error tracking + Performance monitoring
2. **Vercel Analytics** - Web vitals em produção
3. **Grafana + Prometheus** - Backend metrics
4. **New Relic** - APM completo

### Alertas

Configure alertas para:
- Response time > 2s (p95)
- Error rate > 1%
- Database queries > 100ms
- Memory usage > 80%

---

## FAQ

**Q: Por que o dev mode está lento?**
A: Normal. Dev mode é 3-5x mais lento. Teste com `pnpm build && pnpm start` para ver performance real.

**Q: Turbopack está ajudando?**
A: Em teoria sim, mas ainda é experimental. Se estiver com problemas, volte para Webpack removendo `--turbopack`.

**Q: Posso usar Docker em desenvolvimento?**
A: Sim, mas adiciona overhead. Para máxima velocidade, rode localmente sem Docker.

**Q: Como otimizar framer-motion?**
A: Use `layout` prop com cuidado, desabilite animações em listas grandes, e considere `useReducedMotion`.

---

**Última atualização:** Janeiro 2026
