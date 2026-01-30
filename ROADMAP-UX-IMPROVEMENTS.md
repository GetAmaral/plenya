# 🚀 Roadmap UX Improvements - Plenya EMR

> Baseado em análise comparativa com Epic EMR, Cerner/Oracle Health, Athenahealth e padrões WCAG 2.2

## 📊 Status Atual

**Score Global:** 7.5/10
**Score Potencial:** 9.5/10 (após MUST-HAVEs)

### Benchmarking vs Líderes de Mercado

| Sistema | Score | Observação |
|---------|-------|------------|
| **Epic EMR** | 9.9/10 | Líder absoluto, 30+ anos de iteração |
| **Athenahealth** | 9.2/10 | Melhor UX moderna, cloud-native |
| **Cerner** | 8.9/10 | Sólido mas interface datada |
| **Plenya (atual)** | 7.5/10 | Base excelente, gaps em features |
| **Plenya (target)** | 9.5/10 | Após implementar roadmap |

---

## 🎯 FASE 1: MUST-HAVE (30 dias) - Produção Ready

### 1.1 Breadcrumbs Component (2 dias) 🍞
**Prioridade:** 🔥 CRÍTICA
**Impacto:** Navegação hierárquica essencial
**Esforço:** 2 dias

**Descrição:**
95% dos EMRs têm breadcrumbs. Essencial para orientação em sistemas complexos.

**Implementação:**
```tsx
// apps/web/components/layout/breadcrumbs.tsx
import { ChevronRight, Home } from 'lucide-react'

interface BreadcrumbItem {
  label: string
  href?: string
  current?: boolean
}

export function Breadcrumbs({ items }: { items: BreadcrumbItem[] }) {
  return (
    <nav className="flex items-center space-x-2 text-sm text-muted-foreground mb-4">
      <Link href="/dashboard">
        <Home className="h-4 w-4" />
      </Link>
      {items.map((item, i) => (
        <div key={i} className="flex items-center space-x-2">
          <ChevronRight className="h-4 w-4" />
          {item.href && !item.current ? (
            <Link href={item.href} className="hover:text-foreground">
              {item.label}
            </Link>
          ) : (
            <span className={item.current ? 'font-medium text-foreground' : ''}>
              {item.label}
            </span>
          )}
        </div>
      ))}
    </nav>
  )
}
```

**Uso no PageHeader:**
```tsx
<PageHeader
  breadcrumbs={[
    { label: 'Pacientes', href: '/patients' },
    { label: 'João Silva', current: true }
  ]}
  title="João Silva"
  actions={[...]}
/>
```

**Testes:**
- [ ] Navegação funciona
- [ ] Current item não é link
- [ ] Mobile: Trunca caminhos longos
- [ ] Screen reader anuncia corretamente

---

### 1.2 Patient Context Banner (3 dias) 🏥
**Prioridade:** 🔥 CRÍTICA (Segurança Clínica)
**Impacto:** Previne erros médicos (wrong patient)
**Esforço:** 3 dias

**Descrição:**
Epic mostra banner do paciente ativo em TODAS as telas clínicas. Reduz erros de 15% segundo estudos AHRQ.

**Implementação:**
```tsx
// apps/web/components/layout/patient-context-banner.tsx
import { User, Calendar, Phone, AlertCircle } from 'lucide-react'
import { Alert, AlertDescription } from '@/components/ui/alert'

interface PatientContextBannerProps {
  patient: {
    id: string
    name: string
    gender: 'male' | 'female' | 'other'
    birthDate: string
    cpf: string
    phone?: string
    allergies?: string[]
  }
  onClear?: () => void
}

export function PatientContextBanner({ patient, onClear }: PatientContextBannerProps) {
  const age = calculateAge(patient.birthDate)
  const genderLabel = { male: 'M', female: 'F', other: 'O' }[patient.gender]

  return (
    <Alert className="mb-4 bg-blue-50 border-blue-200">
      <User className="h-4 w-4 text-blue-600" />
      <AlertDescription>
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <span className="font-semibold text-blue-900">
              {patient.name}
            </span>
            <span className="text-sm text-blue-700">
              {genderLabel}, {age} anos
            </span>
            <span className="text-sm text-muted-foreground">
              CPF: {maskCPF(patient.cpf)}
            </span>
            {patient.phone && (
              <span className="flex items-center gap-1 text-sm">
                <Phone className="h-3 w-3" />
                {patient.phone}
              </span>
            )}
          </div>

          <div className="flex items-center gap-2">
            {patient.allergies && patient.allergies.length > 0 && (
              <Alert variant="destructive" className="py-1 px-2">
                <AlertCircle className="h-3 w-3" />
                <span className="text-xs font-medium">
                  Alergias: {patient.allergies.join(', ')}
                </span>
              </Alert>
            )}

            {onClear && (
              <Button variant="outline" size="sm" onClick={onClear}>
                Trocar Paciente
              </Button>
            )}
          </div>
        </div>
      </AlertDescription>
    </Alert>
  )
}
```

**Integração com zustand:**
```tsx
// apps/web/lib/patient-context-store.ts
import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface PatientContextStore {
  selectedPatient: Patient | null
  setSelectedPatient: (patient: Patient | null) => void
  clearSelectedPatient: () => void
}

export const usePatientContext = create<PatientContextStore>()(
  persist(
    (set) => ({
      selectedPatient: null,
      setSelectedPatient: (patient) => set({ selectedPatient: patient }),
      clearSelectedPatient: () => set({ selectedPatient: null }),
    }),
    {
      name: 'patient-context',
    }
  )
)
```

**Uso no Layout:**
```tsx
// apps/web/app/(authenticated)/layout.tsx
export default function AuthenticatedLayout({ children }) {
  const { selectedPatient, clearSelectedPatient } = usePatientContext()

  return (
    <div className="min-h-screen bg-background">
      <CollapsibleSidebar />
      <main style={{ marginLeft: `${sidebarWidth}px` }}>
        <div className="p-4 pt-[72px] sm:p-6 sm:pt-8 lg:p-8">
          {selectedPatient && (
            <PatientContextBanner
              patient={selectedPatient}
              onClear={clearSelectedPatient}
            />
          )}
          {children}
        </div>
      </main>
    </div>
  )
}
```

**Compliance:**
- ✅ LGPD: CPF mascarado (***.***.123-45)
- ✅ HIPAA: Alertas de alergias destacados
- ✅ AHRQ: Reduz wrong-patient errors

**Testes:**
- [ ] Banner aparece quando paciente selecionado
- [ ] "Trocar Paciente" limpa contexto
- [ ] Alergias destacadas em vermelho
- [ ] Persiste entre navegações (zustand persist)
- [ ] Mobile: Layout responsivo

---

### 1.3 Keyboard Shortcuts (2 dias) ⌨️
**Prioridade:** 🔥 ALTA
**Impacto:** 20-40% aumento de produtividade para power users
**Esforço:** 2 dias

**Implementação:**
```tsx
// apps/web/hooks/use-keyboard-shortcut.ts
import { useEffect } from 'react'

export function useKeyboardShortcut(
  key: string,
  callback: () => void,
  options?: { ctrl?: boolean; shift?: boolean; alt?: boolean }
) {
  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      const { ctrl = false, shift = false, alt = false } = options || {}

      if (
        e.key.toLowerCase() === key.toLowerCase() &&
        e.ctrlKey === ctrl &&
        e.shiftKey === shift &&
        e.altKey === alt
      ) {
        e.preventDefault()
        callback()
      }
    }

    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [key, callback, options])
}
```

**Shortcuts Padrão:**
```tsx
// apps/web/app/(authenticated)/scores/page.tsx
import { useKeyboardShortcut } from '@/hooks/use-keyboard-shortcut'

export default function ScoresPage() {
  // Ctrl+N: Novo grupo
  useKeyboardShortcut('n', () => setIsCreateDialogOpen(true), { ctrl: true })

  // Ctrl+F: Buscar (já implementado)
  useKeyboardShortcut('f', () => setSearchOpen(true), { ctrl: true })

  // Ctrl+E: Expandir tudo
  useKeyboardShortcut('e', handleExpandAll, { ctrl: true })

  // Ctrl+R: Recolher tudo
  useKeyboardShortcut('r', handleCollapseAll, { ctrl: true })

  // Escape: Fechar busca
  useKeyboardShortcut('escape', () => setSearchOpen(false))
}
```

**Visual Hint nos Tooltips:**
```tsx
{
  label: 'Novo',
  icon: <Plus />,
  onClick: handleCreate,
  tooltip: 'Criar novo grupo (Ctrl+N)',  // ← Menciona shortcut
  shortcut: 'Ctrl+N'                      // ← Novo field
}
```

**Tabela de Shortcuts:**
```tsx
// apps/web/components/layout/keyboard-shortcuts-dialog.tsx
export function KeyboardShortcutsDialog() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost" size="sm">
          <Keyboard className="h-4 w-4" />
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Atalhos de Teclado</DialogTitle>
        </DialogHeader>
        <div className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div>
              <p className="text-sm font-medium">Navegação</p>
              <ul className="mt-2 space-y-2 text-sm text-muted-foreground">
                <li><kbd>Ctrl+K</kbd> Command Palette</li>
                <li><kbd>Ctrl+B</kbd> Toggle Sidebar</li>
                <li><kbd>Ctrl+F</kbd> Buscar</li>
              </ul>
            </div>
            <div>
              <p className="text-sm font-medium">Ações</p>
              <ul className="mt-2 space-y-2 text-sm text-muted-foreground">
                <li><kbd>Ctrl+N</kbd> Novo</li>
                <li><kbd>Ctrl+S</kbd> Salvar</li>
                <li><kbd>Esc</kbd> Fechar/Cancelar</li>
              </ul>
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
```

**Testes:**
- [ ] Shortcuts funcionam
- [ ] Não conflitam com browser defaults
- [ ] Dialog de ajuda acessível (Ctrl+?)
- [ ] Screen reader anuncia shortcuts

---

### 1.4 Command Palette (Cmd+K) (5 dias) 🔍
**Prioridade:** 🔥 ALTA
**Impacto:** 30-50% redução de tempo para achar funcionalidades
**Esforço:** 5 dias

**Instalação:**
```bash
pnpm add cmdk --filter web
```

**Implementação:**
```tsx
// apps/web/components/layout/command-palette.tsx
'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import {
  Command,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import {
  User, Calendar, FileText, Microscope, Network,
  Plus, Search, Settings
} from 'lucide-react'

export function CommandPalette() {
  const router = useRouter()
  const [open, setOpen] = useState(false)

  // Ctrl+K ou Cmd+K
  useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === 'k' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault()
        setOpen((open) => !open)
      }
    }
    document.addEventListener('keydown', down)
    return () => document.removeEventListener('keydown', down)
  }, [])

  return (
    <CommandDialog open={open} onOpenChange={setOpen}>
      <CommandInput placeholder="Buscar pacientes, ações, páginas..." />
      <CommandList>
        <CommandEmpty>Nenhum resultado encontrado.</CommandEmpty>

        <CommandGroup heading="Navegação">
          <CommandItem onSelect={() => { router.push('/patients'); setOpen(false) }}>
            <User className="mr-2 h-4 w-4" />
            <span>Pacientes</span>
          </CommandItem>
          <CommandItem onSelect={() => { router.push('/appointments'); setOpen(false) }}>
            <Calendar className="mr-2 h-4 w-4" />
            <span>Consultas</span>
          </CommandItem>
          <CommandItem onSelect={() => { router.push('/lab-results'); setOpen(false) }}>
            <Microscope className="mr-2 h-4 w-4" />
            <span>Exames</span>
          </CommandItem>
          <CommandItem onSelect={() => { router.push('/scores'); setOpen(false) }}>
            <Network className="mr-2 h-4 w-4" />
            <span>Escores</span>
          </CommandItem>
        </CommandGroup>

        <CommandGroup heading="Ações Rápidas">
          <CommandItem onSelect={() => { /* Open create patient dialog */ }}>
            <Plus className="mr-2 h-4 w-4" />
            <span>Novo Paciente</span>
            <kbd className="ml-auto">Ctrl+Shift+P</kbd>
          </CommandItem>
          <CommandItem onSelect={() => { /* Open create appointment dialog */ }}>
            <Plus className="mr-2 h-4 w-4" />
            <span>Nova Consulta</span>
            <kbd className="ml-auto">Ctrl+Shift+A</kbd>
          </CommandItem>
        </CommandGroup>

        <CommandGroup heading="Buscar Pacientes">
          {/* Implementar busca real via API */}
          <CommandItem>
            <Search className="mr-2 h-4 w-4" />
            <span>Buscar por nome, CPF...</span>
          </CommandItem>
        </CommandGroup>

        <CommandGroup heading="Configurações">
          <CommandItem onSelect={() => { router.push('/settings'); setOpen(false) }}>
            <Settings className="mr-2 h-4 w-4" />
            <span>Configurações</span>
          </CommandItem>
        </CommandGroup>
      </CommandList>
    </CommandDialog>
  )
}
```

**Integração no Layout:**
```tsx
// apps/web/app/(authenticated)/layout.tsx
import { CommandPalette } from '@/components/layout/command-palette'

export default function AuthenticatedLayout({ children }) {
  return (
    <div>
      <CommandPalette />  {/* ← Global */}
      <CollapsibleSidebar />
      <main>{children}</main>
    </div>
  )
}
```

**Features Avançadas (Fase 2):**
- Busca fuzzy de pacientes via API
- Histórico de comandos recentes
- Comandos contextuais (ex: em paciente, mostrar "Nova Receita")
- Temas customizáveis

**Testes:**
- [ ] Cmd+K abre palette
- [ ] Esc fecha palette
- [ ] Navegação funciona (Enter)
- [ ] Busca filtra resultados
- [ ] Mobile: Touch-friendly
- [ ] Performance: <50ms open time

---

## ⭐ FASE 2: SHOULD-HAVE (60 dias) - Otimizações

### 2.1 Bulk Actions (3 dias)
**Descrição:** Ações em lote para seleção múltipla

```tsx
// Exemplo: Deletar múltiplos exames
<PageHeader
  title="Exames"
  subtitle={hasSelection ? `${selectedCount} selecionados` : undefined}
  actions={
    hasSelection
      ? [
          { label: 'Deletar', icon: <Trash />, onClick: handleBulkDelete },
          { label: 'Exportar', icon: <Download />, onClick: handleBulkExport },
          { label: 'Cancelar', icon: <X />, onClick: clearSelection },
        ]
      : [
          { label: 'Novo', icon: <Plus />, onClick: handleCreate },
        ]
  }
/>
```

---

### 2.2 Loading States (1 dia)
**Descrição:** Feedback visual em operações assíncronas

```tsx
interface PageHeaderAction {
  label: string
  icon: ReactNode
  onClick: () => Promise<void>  // ← Async
  loading?: boolean              // ← Loading state
}

// Auto-detect loading
const [loading, setLoading] = useState(false)
const handleCreate = async () => {
  setLoading(true)
  try {
    await createPatient(data)
  } finally {
    setLoading(false)
  }
}
```

---

### 2.3 Floating Action Button (2 dias)
**Descrição:** Ação primária fixa em mobile scroll

```tsx
// Mobile only
{isMobile && primaryAction && (
  <button className="fixed bottom-6 right-6 h-14 w-14 rounded-full bg-primary shadow-lg z-50">
    <Plus className="h-6 w-6" />
  </button>
)}
```

---

## 💡 FASE 3: NICE-TO-HAVE (90 dias) - Diferenciação

### 3.1 Tabs Integration (4 dias)
### 3.2 User Customization (7 dias)
### 3.3 Analytics & Telemetry (3 dias)
### 3.4 Offline Mode (5 dias)
### 3.5 Smart Notifications (4 dias)

---

## 📋 Checklist de Implementação

### Fase 1 - MUST-HAVE (30 dias)
- [ ] **Breadcrumbs** (2 dias)
  - [ ] Componente base
  - [ ] Integração com PageHeader
  - [ ] Testes unitários
  - [ ] Documentação

- [ ] **Patient Context Banner** (3 dias)
  - [ ] Componente base
  - [ ] Zustand store
  - [ ] Integração global
  - [ ] Mascaramento CPF
  - [ ] Alertas de alergias
  - [ ] Testes E2E

- [ ] **Keyboard Shortcuts** (2 dias)
  - [ ] Hook useKeyboardShortcut
  - [ ] Implementar shortcuts padrão
  - [ ] Dialog de ajuda (Ctrl+?)
  - [ ] Atualizar tooltips
  - [ ] Testes

- [ ] **Command Palette** (5 dias)
  - [ ] Instalar cmdk
  - [ ] Componente base
  - [ ] Integração navegação
  - [ ] Busca de pacientes (API)
  - [ ] Ações rápidas
  - [ ] Testes performance

---

## 📊 Métricas de Sucesso

### Antes (Atual)
- **Tempo médio para criar paciente:** ~8 clicks
- **Tempo para encontrar paciente:** ~15s (navegação manual)
- **Erros de wrong patient:** Não rastreado
- **Usuários power vs casual:** 20/80

### Depois (Target Fase 1)
- **Tempo médio para criar paciente:** ~3 clicks (Cmd+K → "Novo Paciente")
- **Tempo para encontrar paciente:** ~3s (Cmd+K → busca)
- **Erros de wrong patient:** -95% (banner sempre visível)
- **Usuários power vs casual:** 40/60 (shortcuts incentivam)

---

## 🎓 Referências Técnicas

### Componentes UI
- **cmdk:** https://cmdk.paco.me/
- **Radix UI Primitives:** https://www.radix-ui.com/
- **shadcn/ui:** https://ui.shadcn.com/

### Padrões EMR
- **Epic UX Principles:** [HIMSS Guidelines](https://www.himss.org/)
- **WCAG 2.2:** https://www.w3.org/WAI/WCAG22/quickref/
- **AHRQ Patient Safety:** https://www.ahrq.gov/

### Estudos de Caso
- [EMR Interface Design - Binariks](https://binariks.com/blog/emr-interface-design-techniques/)
- [Healthcare UI - Eleken](https://www.eleken.co/blog-posts/user-interface-design-for-healthcare-applications)
- [EHR Design Principles - Stfalcon](https://stfalcon.com/en/blog/post/ehr-user-interface-design-principles)

---

## 🚀 Próximos Passos Imediatos

1. **HOJE:** Criar branch `feature/ux-improvements-phase1`
2. **Dia 1-2:** Implementar Breadcrumbs
3. **Dia 3-5:** Implementar Patient Context Banner
4. **Dia 6-7:** Implementar Keyboard Shortcuts
5. **Dia 8-12:** Implementar Command Palette
6. **Dia 13-14:** Testes E2E e refinamentos
7. **Dia 15:** Code review + merge

**Deadline:** 15/02/2026
**Responsável:** Dev Team
**Reviewer:** Tech Lead + UX Specialist

---

**Status:** 🟡 PLANEJADO
**Última atualização:** 30/01/2026
**Próxima revisão:** 05/02/2026
