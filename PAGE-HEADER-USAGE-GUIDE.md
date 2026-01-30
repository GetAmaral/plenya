# Guia de Uso: PageHeader Component

## Importação

```tsx
import { PageHeader } from '@/components/layout/page-header'
```

## Props

```tsx
interface PageHeaderProps {
  title: string                    // Título da página (obrigatório)
  description?: string             // Subtítulo/descrição (opcional)
  actions?: PageHeaderAction[]     // Array de ações (opcional)
  children?: ReactNode             // Conteúdo customizado (opcional)
}

interface PageHeaderAction {
  label: string                    // Label do botão
  icon: ReactNode                  // Ícone (lucide-react)
  onClick: () => void              // Handler do click
  variant?: 'default' | 'outline' | 'ghost'  // Estilo do botão
  disabled?: boolean               // Estado desabilitado
  tooltip?: string                 // Texto do tooltip (se diferente do label)
}
```

---

## Exemplos de Uso

### 1. Básico (Apenas Título)

```tsx
export default function PatientsPage() {
  return (
    <div>
      <PageHeader
        title="Pacientes"
        description="Gerencie os pacientes do sistema"
      />

      {/* Conteúdo da página */}
    </div>
  )
}
```

**Resultado:**
```
┌─────────────────────────┐
│ Pacientes               │
│ Gerencie os pacientes...│
├─────────────────────────┤
```

---

### 2. Com Ação Primária

```tsx
import { Plus } from 'lucide-react'

export default function PatientsPage() {
  return (
    <div>
      <PageHeader
        title="Pacientes"
        description="Lista de todos os pacientes cadastrados"
        actions={[
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setDialogOpen(true),
            variant: 'default', // Destaque
          }
        ]}
      />
    </div>
  )
}
```

**Resultado:**
```
┌──────────────────────────────┐
│ Pacientes                    │
│ Lista de todos os pacientes  │
├──────────────────────────────┤
│ [+ Novo]  ← azul, destaque   │
└──────────────────────────────┘
```

---

### 3. Múltiplas Ações (Padrão Recomendado)

```tsx
import { Plus, Search, Filter, Download, Upload } from 'lucide-react'

export default function PatientsPage() {
  return (
    <div>
      <PageHeader
        title="Pacientes"
        actions={[
          {
            label: 'Filtrar',
            icon: <Filter className="h-4 w-4" />,
            onClick: handleFilter,
            variant: 'ghost',
          },
          {
            label: 'Buscar',
            icon: <Search className="h-4 w-4" />,
            onClick: handleSearch,
            tooltip: 'Buscar paciente (Ctrl+F)',
          },
          {
            label: 'Exportar',
            icon: <Download className="h-4 w-4" />,
            onClick: handleExport,
            tooltip: 'Exportar lista em CSV',
          },
          {
            label: 'Importar',
            icon: <Upload className="h-4 w-4" />,
            onClick: handleImport,
          },
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: handleCreate,
            variant: 'default', // Ação primária sempre por último
          }
        ]}
      />
    </div>
  )
}
```

**Resultado Desktop:**
```
┌────────────────────────────────────────────────────────┐
│ Pacientes                                              │
├────────────────────────────────────────────────────────┤
│ [🔽 Filtrar] [🔍 Buscar] [⬇ Exportar] [⬆ Importar] [+ Novo] │
└────────────────────────────────────────────────────────┘
```

**Resultado Mobile:**
```
┌──────────────────────┐
│ Pacientes            │
├──────────────────────┤
│ [🔽][🔍][⬇][⬆][+ Novo]│
│  ↑ Tooltips no touch │
└──────────────────────┘
```

---

### 4. Com Children Customizados

Quando precisar de controles mais complexos (ex: tabs, select, etc):

```tsx
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { Select } from '@/components/ui/select'

export default function AppointmentsPage() {
  return (
    <div>
      <PageHeader
        title="Consultas"
        actions={[
          {
            label: 'Nova',
            icon: <Plus className="h-4 w-4" />,
            onClick: handleCreate,
            variant: 'default',
          }
        ]}
      >
        {/* Filtros customizados */}
        <Tabs defaultValue="upcoming">
          <TabsList>
            <TabsTrigger value="upcoming">Próximas</TabsTrigger>
            <TabsTrigger value="past">Passadas</TabsTrigger>
            <TabsTrigger value="cancelled">Canceladas</TabsTrigger>
          </TabsList>
        </Tabs>

        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Médico" />
          </SelectTrigger>
          {/* ... */}
        </Select>
      </PageHeader>
    </div>
  )
}
```

**Resultado:**
```
┌──────────────────────────────────────────────────┐
│ Consultas                                        │
├──────────────────────────────────────────────────┤
│ [Próximas|Passadas|Canceladas] [Médico ▼] | [+ Nova] │
│  ↑ Children customizados        ↑ Separador   ↑ Actions │
└──────────────────────────────────────────────────┘
```

---

### 5. Ações Condicionais

```tsx
export default function LabResultsPage() {
  const hasSelection = selectedItems.length > 0

  return (
    <div>
      <PageHeader
        title="Resultados"
        actions={[
          {
            label: 'Excluir',
            icon: <Trash className="h-4 w-4" />,
            onClick: handleDelete,
            variant: 'outline',
            disabled: !hasSelection, // Só ativo com seleção
            tooltip: hasSelection
              ? `Excluir ${selectedItems.length} itens`
              : 'Selecione itens para excluir',
          },
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: handleCreate,
            variant: 'default',
          }
        ]}
      />
    </div>
  )
}
```

---

## Boas Práticas

### ✅ DO (Faça)

**1. Use labels curtos (1-2 palavras)**
```tsx
label: 'Novo'        // ✅
label: 'Expandir'    // ✅
label: 'Buscar'      // ✅
```

**2. Coloque ação primária por último**
```tsx
actions={[
  { label: 'Buscar', variant: 'ghost' },
  { label: 'Filtrar', variant: 'ghost' },
  { label: 'Novo', variant: 'default' }, // ✅ Por último
]}
```

**3. Use tooltips para contexto adicional**
```tsx
{
  label: 'Exportar',
  tooltip: 'Exportar lista completa em formato CSV',
}
```

**4. Agrupe ações relacionadas por variant**
```tsx
// Controles de UI (ghost)
{ label: 'Expandir', variant: 'ghost' },
{ label: 'Recolher', variant: 'ghost' },

// Ações frequentes (outline)
{ label: 'Buscar', variant: 'outline' },
{ label: 'Filtrar', variant: 'outline' },

// Ação primária (default)
{ label: 'Novo', variant: 'default' },
```

### ❌ DON'T (Não faça)

**1. Labels prolixos**
```tsx
label: 'Clique aqui para criar um novo paciente' // ❌
label: 'Novo Paciente'                           // ❌ (redundante)
label: 'Novo'                                    // ✅
```

**2. Muitas ações primárias**
```tsx
{ label: 'Novo', variant: 'default' },    // ❌ Múltiplos primários
{ label: 'Salvar', variant: 'default' },  // ❌ confundem usuário
{ label: 'Editar', variant: 'default' },  // ❌
```

**3. Ícones sem significado**
```tsx
<Star className="h-4 w-4" /> // ❌ O que significa?
<Circle className="h-4 w-4" /> // ❌ Não é intuitivo
```

**4. Mais de 8 ações**
```tsx
actions={[...10 botões]} // ❌ Muito poluído
// Use tabs, command palette ou agrupe em dropdown
```

---

## Hierarquia de Variantes

### `variant="default"` (Primária)
- **Quando usar:** Ação principal da página (1 por página)
- **Exemplos:** Criar, Salvar, Enviar
- **Visual:** Azul sólido, destaque máximo

### `variant="outline"` (Secundária)
- **Quando usar:** Ações frequentes mas não primárias
- **Exemplos:** Buscar, Filtrar, Exportar
- **Visual:** Borda, background branco

### `variant="ghost"` (Terciária)
- **Quando usar:** Controles de UI, ações menos frequentes
- **Exemplos:** Expandir, Recolher, Mostrar/Ocultar
- **Visual:** Transparente, hover colorido

---

## Ícones Recomendados (lucide-react)

### Ações Comuns
```tsx
import {
  Plus,          // Criar novo
  Search,        // Buscar
  Filter,        // Filtrar
  Download,      // Exportar/Download
  Upload,        // Importar/Upload
  Printer,       // Imprimir
  Mail,          // Enviar email
  Share,         // Compartilhar
  Edit,          // Editar
  Trash,         // Excluir
  Copy,          // Copiar/Duplicar
  Archive,       // Arquivar
  Eye,           // Visualizar
  EyeOff,        // Ocultar
  Settings,      // Configurações
  MoreVertical,  // Menu de opções
} from 'lucide-react'
```

### Controles de Navegação
```tsx
import {
  ChevronLeft,   // Voltar
  ChevronRight,  // Avançar
  ChevronsDown,  // Expandir
  ChevronsUp,    // Recolher
  Minimize2,     // Minimizar
  Maximize2,     // Maximizar
  ZoomIn,        // Aumentar
  ZoomOut,       // Diminuir
} from 'lucide-react'
```

### Visualizações
```tsx
import {
  List,          // Lista
  Grid,          // Grade
  Network,       // Mindmap/Grafo
  LayoutGrid,    // Dashboard
  Table,         // Tabela
  FileImage,     // Pôster/Imagem
} from 'lucide-react'
```

---

## Responsividade

O componente é **100% responsivo** por padrão:

### Mobile (<640px)
- Apenas ícones visíveis
- Tooltips no touch
- Layout em linha única com wrap
- Padding reduzido

### Tablet (640-1024px)
- Labels visíveis
- Tooltips no hover
- Layout pode quebrar linha
- Espaçamento normal

### Desktop (≥1024px)
- Labels completos
- Tooltips no hover
- Layout em linha única
- Espaçamento confortável

**Não precisa de media queries manuais!** O componente usa `hidden sm:inline` do Tailwind.

---

## Acessibilidade

O PageHeader segue WCAG 2.1 AA:

✅ **Keyboard Navigation**
- Tab navega entre botões
- Enter/Space ativa ação
- Esc fecha tooltips

✅ **Screen Readers**
- Labels semânticos
- ARIA labels em ícones
- Tooltips descritivos

✅ **Visual**
- Contraste adequado (4.5:1)
- Focus ring visível
- Touch targets ≥32px

✅ **Motor**
- Espaçamento generoso
- Área clicável grande
- Sem hover obrigatório

---

## Troubleshooting

### Problema: Botões sobrepostos em mobile
**Solução:** Reduza número de ações ou use variant="ghost" para controles secundários

### Problema: Labels muito grandes em desktop
**Solução:** Use labels curtos (1-2 palavras) e coloque detalhes no tooltip

### Problema: Muitas ações (>8)
**Solução:**
1. Agrupe em tabs (ex: visualizações)
2. Use dropdown para ações raras
3. Command Palette (Cmd+K) para power users

### Problema: Título sobreposto pelo menu mobile
**Solução:** O layout já tem `pt-[72px]` em mobile. Se ainda sobrepõe, verifique z-index customizados.

---

## Migração de Código Antigo

### Antes (Padrão Antigo)
```tsx
<div className="flex justify-between items-center mb-6">
  <div>
    <h1 className="text-2xl font-bold">Pacientes</h1>
    <p className="text-muted-foreground">Gerencie os pacientes</p>
  </div>
  <div className="flex gap-2">
    <Button onClick={handleSearch} variant="outline">
      <Search className="mr-2 h-4 w-4" />
      Procurar Pacientes
    </Button>
    <Button onClick={handleCreate}>
      <Plus className="mr-2 h-4 w-4" />
      Criar Novo Paciente
    </Button>
  </div>
</div>
```

### Depois (PageHeader)
```tsx
<PageHeader
  title="Pacientes"
  description="Gerencie os pacientes"
  actions={[
    {
      label: 'Buscar',
      icon: <Search className="h-4 w-4" />,
      onClick: handleSearch,
    },
    {
      label: 'Novo',
      icon: <Plus className="h-4 w-4" />,
      onClick: handleCreate,
      variant: 'default',
    }
  ]}
/>
```

**Benefícios:**
- ✅ 30% menos código
- ✅ Responsivo automaticamente
- ✅ Tooltips gratuitos
- ✅ Padrão consistente
- ✅ Acessibilidade embutida

---

## Componentes Relacionados

Para layouts mais complexos, combine com:

- **Breadcrumbs** - Navegação hierárquica
- **Tabs** - Múltiplas visões da mesma entidade
- **Command Palette** - Ações via teclado (Cmd+K)
- **Floating Action Button** - Ação primária fixa em mobile

Exemplo completo:
```tsx
<div>
  <Breadcrumbs items={[
    { label: 'Dashboard', href: '/' },
    { label: 'Pacientes', href: '/patients' },
  ]} />

  <PageHeader
    title="Pacientes"
    actions={[...]}
  >
    <Tabs>
      <TabsList>
        <TabsTrigger value="all">Todos</TabsTrigger>
        <TabsTrigger value="active">Ativos</TabsTrigger>
      </TabsList>
    </Tabs>
  </PageHeader>

  {/* Conteúdo */}
</div>
```

---

## Referências

- **Componente:** `apps/web/components/layout/page-header.tsx`
- **Exemplo Real:** `apps/web/app/(authenticated)/scores/page.tsx`
- **Documentação:** `PAGE-HEADER-IMPROVEMENTS.md`
- **shadcn/ui:** https://ui.shadcn.com/docs/components/button
- **Lucide Icons:** https://lucide.dev/icons
- **Tailwind:** https://tailwindcss.com/docs/responsive-design
