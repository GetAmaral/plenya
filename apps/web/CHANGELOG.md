# Changelog - Plenya EMR Web App

## [Janeiro 2026] - Fase 4 Completa

### ✨ Fase 4 Parte 1 - Fundação do Frontend

#### Design System Moderno
- **Design Tokens completos** (`lib/design-tokens.ts`)
  - Spacing scale baseado em 4px
  - Typography scale fluida
  - Border radius modernos
  - Shadows suaves
  - Animation timings otimizados
  - Cores semânticas médicas

- **Tailwind CSS Avançado** (`tailwind.config.ts`)
  - 10+ animações customizadas (fade, slide, scale)
  - Utilities para glassmorphism
  - Utilities para Bento Grid
  - Cores médicas semânticas
  - Scrollbar customizado

- **Global CSS** (`app/globals.css`)
  - Classes utilitárias: `.glass`, `.bento-grid`, `.skeleton`
  - Gradients médicos
  - Focus states modernos
  - Selection styling

#### Componentes UI Modernos
- **Badge** - Variantes médicas (stable, observation, critical, urgent)
- **Avatar** - Com fallback de iniciais
- **Skeleton** - Loading states suaves
- **Toaster (Sonner)** - Notificações elegantes
- **Button** - Com variantes e animações
- **Input** - Estilizado com focus states
- **Card** - Container com shadow e border
- **Label** - Labels acessíveis

#### Páginas
- **Login Page**
  - Glassmorphism com backdrop blur
  - Animated gradient orbs (Framer Motion)
  - Micro-interactions em todos elementos
  - Spring animation no logo
  - Toast notifications
  - Loading states

- **Dashboard (Bento Grid)**
  - 4 cards de estatísticas com icons
  - Lista de pacientes recentes com avatares
  - Timeline de próximas consultas
  - Card de ações rápidas
  - Stagger animations
  - Hover effects

#### Infraestrutura
- **QueryProvider** - TanStack Query configurado
- **API Client** - HTTP client com autenticação
- **Auth Store** - Zustand com persistência localStorage
- **Design Tokens** - Single source of truth

### ✨ Fase 4 Parte 2 - Dashboard e CRUD

#### Autenticação e Navegação
- **Middleware** (`middleware.ts`)
  - Proteção de rotas server-side
  - Redirect para login se não autenticado
  - Matcher configuration

- **Auth Hooks** (`lib/use-auth.ts`)
  - `useRequireAuth()` - Hook de proteção
  - `useAuth()` - Hook de auth state
  - Logout functionality

- **Sidebar Moderna** (`components/dashboard/sidebar.tsx`)
  - Navegação animada com Framer Motion
  - 6 links principais (Dashboard, Pacientes, Consultas, etc)
  - User section com avatar e badge de role
  - Logout button
  - Active state highlighting
  - Hover animations

#### Layouts
- **Dashboard Layout** (`app/dashboard/layout.tsx`)
  - Sidebar fixa à esquerda
  - Content area responsivo
  - Auth protection

- **Patients Layout** (`app/patients/layout.tsx`)
  - Mesmo padrão do dashboard
  - Reutilização da sidebar

#### CRUD de Pacientes
- **Patients List Page** (`app/patients/page.tsx`)
  - **TanStack React Table** v8
    - Sorting (colunas clicáveis)
    - Global filtering (busca em tempo real)
    - Pagination (next/previous)
    - Column visibility

  - **Features**
    - Skeleton loading states (5 rows)
    - Empty state elegante
    - Search bar com ícone
    - Action buttons (Ver, Editar)
    - Badge de gênero
    - Formatação de datas (date-fns pt-BR)
    - Paginação com contadores
    - Hover effects nas rows

  - **API Integration**
    - React Query para fetching
    - Cache automático
    - Loading states
    - Error handling

### ✨ Fase 4 Parte 3 - CRUD Completo e Features Avançadas

#### Páginas CRUD Médicas
- **Appointments Page** (`app/appointments/page.tsx`)
  - TanStack Table com 6 colunas
  - Status tracking (scheduled, confirmed, completed, cancelled)
  - Stats cards por status com icons
  - Date/time formatting (dd/MM/yyyy HH:mm)
  - Action buttons (Ver, Confirmar)
  - Skeleton loading states
  - Search e pagination

- **Prescriptions Page** (`app/prescriptions/page.tsx`)
  - TanStack Table com detalhes de medicação
  - Status tracking (active, completed, cancelled)
  - Colunas: Medicamento, Dosagem, Frequência, Duração
  - Stats cards por status
  - Action buttons (Ver, Imprimir)
  - Search e pagination

- **Anamnesis Page** (`app/anamnesis/page.tsx`)
  - TanStack Table para histórico médico
  - Colunas: Queixa Principal, Alergias, Medicamentos
  - Stats: Total de anamneses, Com alergias, Em medicação
  - Badge para alergias com contador
  - Date formatting pt-BR
  - Action buttons (Ver, Editar)

- **Lab Results Page** (`app/lab-results/page.tsx`)
  - TanStack Table para exames laboratoriais
  - Status tracking (pending, completed, reviewed)
  - Colunas: Tipo de Exame, Data, Resultado, Valores de Referência
  - Stats cards por status
  - Download button para resultados
  - TestTube icon (Lucide React)

#### Features Avançadas
- **Command Palette** (`components/command-palette.tsx`)
  - Atalho global ⌘K / Ctrl+K
  - Navegação rápida para todas as páginas
  - Busca com filtro em tempo real
  - Ações rápidas (Logout)
  - Dialog com cmdk
  - Keyboard shortcuts

- **Charts com Tremor** (Dashboard aprimorado)
  - **AreaChart**: Crescimento de Pacientes (6 meses)
  - **DonutChart**: Status dos Pacientes (Estável/Observação/Crítico)
  - **BarChart**: Consultas por Tipo (Consulta/Retorno/Exame/Emergência)
  - Formatação pt-BR (Intl.NumberFormat)
  - Animações suaves
  - Cores semânticas

#### Novos Componentes UI
- **Command** (`components/ui/command.tsx`)
  - CommandDialog, CommandInput, CommandList
  - CommandEmpty, CommandGroup, CommandItem
  - CommandSeparator, CommandShortcut
  - Baseado em cmdk + Radix Dialog

- **Dialog** (`components/ui/dialog.tsx`)
  - DialogOverlay, DialogContent
  - DialogHeader, DialogFooter, DialogTitle
  - Radix UI Dialog primitives
  - Animations (fade, zoom, slide)

#### Layouts Adicionais
- `app/appointments/layout.tsx` - Appointments layout com sidebar
- `app/prescriptions/layout.tsx` - Prescriptions layout com sidebar
- `app/anamnesis/layout.tsx` - Anamnesis layout com sidebar
- `app/lab-results/layout.tsx` - Lab Results layout com sidebar

#### Bibliotecas Adicionadas
```json
{
  "framer-motion": "^11.15.0",
  "@tremor/react": "^3.18.5",
  "recharts": "^2.15.0",
  "sonner": "^1.7.1",
  "cmdk": "^1.0.4",
  "vaul": "^1.1.1",
  "date-fns": "^4.1.0",
  "@tanstack/react-table": "^8.20.5",
  "@radix-ui/*": "Multiple packages"
}
```

## Build Status

```
Route (app)                    Size    First Load JS
┌ ○ /                         122 B        103 kB
├ ○ /_not-found              998 B        103 kB
├ ○ /anamnesis              4.41 kB       181 kB  ✅ NEW!
├ ○ /appointments           4.37 kB       181 kB  ✅ NEW!
├ ○ /dashboard               135 kB       281 kB  ✅ (com Charts!)
├ ○ /lab-results            4.40 kB       181 kB  ✅ NEW!
├ ○ /login                  26.3 kB       182 kB  ✅
├ ○ /patients               4.11 kB       181 kB  ✅
└ ○ /prescriptions          4.30 kB       181 kB  ✅ NEW!

ƒ Middleware                 33.9 kB      ✅
```

**Total**: 9 routes, Command Palette global, Sidebar em todas as páginas protegidas

## Estrutura de Arquivos

```
apps/web/
├── app/
│   ├── anamnesis/
│   │   ├── layout.tsx       ✅ NEW! Anamnesis layout
│   │   └── page.tsx         ✅ NEW! Histórico médico
│   ├── appointments/
│   │   ├── layout.tsx       ✅ NEW! Appointments layout
│   │   └── page.tsx         ✅ NEW! Agendamento de consultas
│   ├── dashboard/
│   │   ├── layout.tsx       ✅ Dashboard layout com sidebar
│   │   └── page.tsx         ✅ Dashboard com charts Tremor
│   ├── lab-results/
│   │   ├── layout.tsx       ✅ NEW! Lab results layout
│   │   └── page.tsx         ✅ NEW! Exames laboratoriais
│   ├── login/
│   │   └── page.tsx         ✅ Login com glassmorphism
│   ├── patients/
│   │   ├── layout.tsx       ✅ Patients layout
│   │   └── page.tsx         ✅ TanStack Table listing
│   ├── prescriptions/
│   │   ├── layout.tsx       ✅ NEW! Prescriptions layout
│   │   └── page.tsx         ✅ NEW! Prescrições médicas
│   ├── layout.tsx           ✅ Root layout (CommandPalette global!)
│   ├── page.tsx             ✅ Redirect to login
│   └── globals.css          ✅ Design system CSS
├── components/
│   ├── command-palette.tsx  ✅ NEW! ⌘K quick navigation
│   ├── dashboard/
│   │   └── sidebar.tsx      ✅ Modern sidebar navigation
│   └── ui/
│       ├── avatar.tsx       ✅ Radix Avatar
│       ├── badge.tsx        ✅ Medical badges
│       ├── button.tsx       ✅ Button component
│       ├── card.tsx         ✅ Card container
│       ├── command.tsx      ✅ NEW! Command palette UI
│       ├── dialog.tsx       ✅ NEW! Dialog primitives
│       ├── input.tsx        ✅ Form input
│       ├── label.tsx        ✅ Form label
│       ├── skeleton.tsx     ✅ Loading skeleton
│       └── toaster.tsx      ✅ Sonner toasts
├── lib/
│   ├── api-client.ts        ✅ HTTP client
│   ├── auth-store.ts        ✅ Zustand auth store
│   ├── design-tokens.ts     ✅ Design system tokens
│   ├── query-provider.tsx   ✅ React Query provider
│   ├── use-auth.ts          ✅ Auth hooks
│   └── utils.ts             ✅ cn() helper
├── middleware.ts            ✅ Route protection
├── turbo.json               ✅ Updated to Turborepo 2.0
├── tailwind.config.ts       ✅ Tailwind advanced config
├── DESIGN_SYSTEM.md         ✅ Design documentation
└── CHANGELOG.md             ✅ This file
```

## Próximos Passos

### Fase 5: Mobile Apps (Futuro)
- React Native + Expo SDK 56
- Compartilhar types com backend
- Offline-first architecture

### Fase 6: Hardening LGPD (Futuro)
- Relatórios de auditoria
- Exportação de dados do paciente
- Políticas de retenção

### Fase 7: Deploy (Futuro)
- Hetzner VPS setup
- Coolify deployment
- SSL + Domain configuration

---

**Última atualização**: Janeiro 2026
**Versão**: 1.0.0
