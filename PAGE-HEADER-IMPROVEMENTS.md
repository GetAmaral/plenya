# Melhorias no PageHeader - EMR Moderno

## O Que Foi Implementado

### 1. **Novo Componente PageHeader Redesenhado**
**Arquivo:** `apps/web/components/layout/page-header.tsx`

**Melhorias:**
- ✅ **Ícones com tooltips** - Em mobile mostra apenas ícones, em desktop mostra label
- ✅ **Labels concisos** - Removida prolixidade ("Expandir Tudo" → "Expandir")
- ✅ **Tooltips descritivos** - Hover mostra explicação detalhada
- ✅ **Layout responsivo** - `sm:inline` para mostrar texto apenas em telas maiores
- ✅ **Separador visual** - Divide ações customizadas das ações do array
- ✅ **Flexbox com wrap** - Botões quebram linha automaticamente
- ✅ **Tamanho consistente** - Todos os botões com `size="sm"`

**Código Exemplo:**
```tsx
<PageHeader
  title="Escores"
  description="Gestão de critérios"
  actions={[
    {
      label: 'Expandir',
      icon: <ChevronsDown className="h-4 w-4" />,
      onClick: handleExpand,
      tooltip: 'Expandir tudo (com textos clínicos)',
    },
    // ...
  ]}
/>
```

### 2. **Página de Escores Otimizada**
**Arquivo:** `apps/web/app/(authenticated)/scores/page.tsx`

**Antes:**
```tsx
primaryAction: 'Novo Grupo'
secondaryActions: [
  'Procurar',
  'Visualizar Mindmap',
  'Versão Impressão',
  'Pôster 60x300cm'
]
+ 3 botões de expansão no children
= 8 botões total (prolixo!)
```

**Depois:**
```tsx
actions: [
  'Expandir',           // + tooltip: "com textos clínicos"
  'Expandir Rápido',    // + tooltip: "sem textos"
  'Recolher',
  'Buscar',             // + tooltip: "Ctrl+F"
  'Mindmap',
  'Imprimir',
  'Pôster',
  'Novo'                // variant: 'default' (destaque)
]
= 8 botões organizados em linha única, responsivos
```

### 3. **Layout Mobile Melhorado**
**Arquivo:** `apps/web/app/(authenticated)/layout.tsx`

**Ajustes:**
```tsx
// Antes
pt-16 sm:p-6 lg:pt-8

// Depois
pt-20 sm:pt-8  // +25% padding mobile, normaliza desktop
```

**Motivo:** Botão de menu mobile tem `z-50` e ocupa ~56px. Com `pt-20` (80px) garantimos 24px de margem.

---

## Hierarquia Z-Index (Corrigida)

```
z-50: Botão menu mobile
z-40: Sidebar overlay + Sidebar
z-30: (disponível para modals)
z-20: (disponível para dropdowns)
z-10: (disponível para tooltips)
```

---

## Padrões de UX EMR 2026 Aplicados

### ✅ **Progressive Disclosure**
- Mobile: Apenas ícones (economia de espaço)
- Desktop: Ícone + Label (clareza)
- Tooltip: Contexto completo (ajuda)

### ✅ **Visual Hierarchy**
- Botão primário (`variant="default"`) → Novo Grupo
- Botões secundários (`variant="outline"`) → Ações frequentes
- Botões terciários (`variant="ghost"`) → Controles de UI

### ✅ **Touch-Friendly**
- Botões `size="sm"` = 32px altura (mínimo WCAG)
- Gap de 8px entre botões
- Área clicável adequada para dedos

### ✅ **Keyboard Navigation**
- Tooltips acessíveis via teclado
- Atalhos preservados (Ctrl+F, Cmd+B)
- Focus ring visível (shadcn/ui padrão)

---

## Comparação Visual

### Antes:
```
┌──────────────────────────────────────────┐
│ Gestão de Escores                        │
│ Gerencie os critérios de...             │
├──────────────────────────────────────────┤
│ [Novo Grupo]                             │
│ [Procurar] [Visualizar Mindmap] [▼ Mais]│
│ [Expandir Tudo] [Expandir (sem textos)] │
│ [Recolher Tudo]                          │
└──────────────────────────────────────────┘
  ❌ 8 botões em 3 linhas (bagunça)
  ❌ Labels prolixos
  ❌ Dropdown esconde opções
```

### Depois:
```
┌──────────────────────────────────────────┐
│ Escores                                  │
│ Gestão de critérios de...               │
├──────────────────────────────────────────┤
│ [↓] [⇅] [⊟] | [🔍] [🗺] [🖨] [📄] [+ Novo]│
│                                          │
│  Desktop: Labels visíveis                │
│  Mobile: Apenas ícones (com tooltips)    │
└──────────────────────────────────────────┘
  ✅ 8 botões em 1 linha (clean)
  ✅ Labels curtos
  ✅ Todas opções visíveis
  ✅ Responsivo
```

---

## Próximos Passos Sugeridos

### 1. **Command Palette (Opcional)**
Se quiser levar UX ao próximo nível:
```tsx
// Cmd+K para abrir
<CommandPalette>
  <Command.Item>Novo Grupo (Ctrl+N)</Command.Item>
  <Command.Item>Buscar (Ctrl+F)</Command.Item>
  <Command.Item>Expandir Tudo (Ctrl+E)</Command.Item>
  ...
</CommandPalette>
```

### 2. **Breadcrumbs (Recomendado)**
Para navegação hierárquica:
```tsx
<Breadcrumbs>
  <Breadcrumb.Item href="/dashboard">Dashboard</Breadcrumb.Item>
  <Breadcrumb.Item href="/scores">Escores</Breadcrumb.Item>
  <Breadcrumb.Item active>Grupo XYZ</Breadcrumb.Item>
</Breadcrumbs>
```

### 3. **Tabs para Visualizações (Recomendado)**
Em vez de botões separados para Mindmap/Print/Poster:
```tsx
<Tabs defaultValue="list">
  <TabsList>
    <TabsTrigger value="list">Lista</TabsTrigger>
    <TabsTrigger value="mindmap">Mindmap</TabsTrigger>
    <TabsTrigger value="print">Impressão</TabsTrigger>
  </TabsList>
</Tabs>
```

### 4. **Floating Action Button (Mobile)**
Ação primária sempre acessível:
```tsx
{isMobile && (
  <Button className="fixed bottom-6 right-6 z-50 h-14 w-14 rounded-full">
    <Plus className="h-6 w-6" />
  </Button>
)}
```

---

## Testes Necessários

### Desktop (≥1024px)
- [ ] Título visível sem overlap com sidebar
- [ ] Todos os 8 botões em linha única
- [ ] Labels de botões visíveis
- [ ] Tooltips funcionam no hover
- [ ] Sidebar collapse funciona (Cmd+B)

### Tablet (768-1023px)
- [ ] Layout se ajusta corretamente
- [ ] Botões quebram linha se necessário
- [ ] Touch targets adequados (≥32px)

### Mobile (<768px)
- [ ] Título com pt-20 (não sobrepõe menu)
- [ ] Apenas ícones visíveis
- [ ] Tooltips funcionam no touch longo
- [ ] Scroll horizontal não aparece
- [ ] Sidebar overlay fecha ao clicar fora

### Acessibilidade
- [ ] Navegação por teclado (Tab)
- [ ] Screen reader lê labels
- [ ] Focus ring visível
- [ ] Contraste adequado (WCAG AA)

---

## Código de Referência

### Estrutura do PageHeader
```tsx
interface PageHeaderAction {
  label: string          // Label curto (mobile: tooltip, desktop: visível)
  icon: ReactNode       // Sempre visível
  onClick: () => void
  variant?: 'default' | 'outline' | 'ghost'
  disabled?: boolean
  tooltip?: string      // Override do label para tooltip
}

// Mobile
<Button>
  {icon}
  <span className="hidden sm:inline">{label}</span>
</Button>

// Desktop
<Button>
  {icon}
  <span>{label}</span>
</Button>
```

### Hierarquia de Importância
1. **Primary Action** → `variant="default"` (azul, destaque)
2. **Frequent Actions** → `variant="outline"` (borda)
3. **UI Controls** → `variant="ghost"` (transparente)

---

## Performance

- ✅ Tooltips lazy-loaded (Radix UI)
- ✅ Icons tree-shaken (lucide-react)
- ✅ CSS-in-JS evitado (Tailwind)
- ✅ Re-renders minimizados (React.memo se necessário)

---

## Conclusão

O novo PageHeader segue as melhores práticas de EMR moderno:
- **Clean** - Apenas 1 linha de ações
- **Conciso** - Labels curtos e diretos
- **Responsivo** - Adapta mobile/desktop
- **Acessível** - Tooltips e keyboard nav
- **Escalável** - Fácil adicionar novas ações

**Referências de UX:**
- Epic EMR (2026) - Command palette + tooltips
- Athenahealth - Icon-first navigation
- Cerner Millennium - Progressive disclosure
- Material Design 3 - Touch targets
- shadcn/ui - Component patterns
