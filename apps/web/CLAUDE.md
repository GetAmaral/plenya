# apps/web — Frontend EMR + Portal do Paciente (Next.js)

Um único app Next.js que serve **dois públicos**: o EMR profissional e o portal do paciente.
Veja as Regras de Ouro no [CLAUDE.md raiz](../../CLAUDE.md).

## Stack
Next 16.2 (Turbopack) · React 19.2 · TanStack Query v5.99 · Zustand 5 · shadcn/Radix ·
Tailwind 3.4 (+ `@plenya/brand`) · react-hook-form + Zod · Recharts/Tremor · Tiptap.
`next.config.ts` tem `ignoreBuildErrors: true` (dívida TS pré-existente).

## Roteamento (dois portais, um codebase)
- **`app/(authenticated)/`** — EMR para staff (admin, doctor, nurse, nutritionist, psychologist,
  physicalEducator, secretary, manager). Grupos: dashboard, patients, anamnesis(+templates),
  appointments/calendario, health-scores, scores (admin: tree/mindmap/poster), lab-results
  (+requests, views, templates), prescriptions, articles, methods, leads, campaigns, conversas,
  continuum, training, configuracoes, admin, profile.
- **`app/patient-portal/`** — paciente (role `patient`): escores, escore-light, exames, consultas,
  documentos, continuum, boxes, mensagens, perfil, auth (login/magic/invite/esqueci-senha).
- **`app/(unauthenticated)/`** + `app/sala/[token]/` — workout público e sala de telemedicina.
- **`middleware.ts`** reescreve `minha.*`/`portal.*` → `/patient-portal` (URL fica limpa).
  Exceção: `/sala/[token]` não é reescrito.

## API e estado
- `lib/api-client.ts` — classe `APIClient` (refresh de token automático, blob/PDF, FormData).
- `lib/api/*.ts` — 40 módulos de hooks TanStack Query (um por domínio). Mutations invalidam via
  `queryClient.invalidateQueries`. Portal usa `patient-portal-api.ts` (auth pública própria).
- `lib/auth-store.ts` — Zustand persistido (`plenya-auth`). `NEXT_PUBLIC_DEV_BYPASS_AUTH=true`
  injeta admin fake em dev (ver [.claude/workflows/dev-bypass-auth.md](../../.claude/workflows/dev-bypass-auth.md)).

## Hooks obrigatórios
- **Todo formulário:** `useFormNavigation({ formRef })` — [.claude/frontend/form-navigation.md](../../.claude/frontend/form-navigation.md).
- **Toda página com dados de paciente:** `useRequireSelectedPatient()` — [.claude/frontend/patient-context.md](../../.claude/frontend/patient-context.md).
- Auth: `useRequireAuth()` (staff) · `useRequirePatientAuth()` (portal).

## Componentes
`components/` por domínio: layout (collapsible-sidebar ~linha 52, top-bar, patient-context-bar),
ui (design system Radix), health-scores (radar/evolution), scores (tree + dialogs CRUD),
training, lab-results, continuum, patient-portal, conversations, campaigns, leads, anamnesis,
prescriptions, admin.

Patterns de query: [.claude/frontend/tanstack-query.md](../../.claude/frontend/tanstack-query.md).
