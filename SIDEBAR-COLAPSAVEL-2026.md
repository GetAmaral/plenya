# 🎨 Sidebar Colapsável - Implementação 2026

## 📋 Resumo

Sistema de navegação moderno com sidebar colapsável aplicado em **TODAS** as páginas autenticadas, seguindo as melhores práticas de UX/UI para 2026.

---

## ✨ Features Implementadas

### 1. Sidebar Colapsável
- **Estado Expandido:** 256px de largura
- **Estado Colapsado:** 80px de largura
- **Transições suaves:** 300ms ease-in-out
- **Ícones aumentados** quando colapsado para melhor usabilidade

### 2. Persistência de Estado
- Salvo em **localStorage** como `sidebar-collapsed`
- Sincronizado entre abas do navegador
- Estado restaurado ao recarregar a página

### 3. Atalhos de Teclado
- **⌘ + B** (Mac) ou **Ctrl + B** (Windows/Linux) para alternar
- Dica visual no rodapé quando expandido

### 4. Responsividade
- **Desktop (≥1024px):** Sidebar sempre visível, pode colapsar
- **Mobile (<1024px):**
  - Sidebar oculta por padrão
  - Menu hamburguer no canto superior esquerdo
  - Overlay escuro quando aberto
  - Fecha automaticamente ao navegar

### 5. Tooltips
- Quando **colapsado**, todos os itens mostram tooltip ao passar mouse
- Posicionados à direita (side="right")
- Delay: 0ms para resposta instantânea

### 6. Indicadores Visuais
- **Rota ativa:** Destaque azul com sombra
- **Hover:** Animação de deslocamento (4px)
- **Avatar do usuário:** Iniciais do email
- **Badge de role:** Cor diferente por função (admin, doctor, nurse, patient)

---

## 🏗️ Arquitetura

### Route Group: (authenticated)

```
apps/web/app/
├── (authenticated)/           ← Route Group (não afeta URL)
│   ├── layout.tsx            ← Layout com sidebar aplicado aqui
│   ├── dashboard/
│   ├── patients/
│   ├── appointments/
│   ├── anamnesis/
│   ├── prescriptions/
│   ├── lab-results/
│   ├── lab-requests/
│   ├── lab-request-templates/
│   ├── scores/
│   └── articles/
└── login/                     ← Fora do route group (sem sidebar)
```

**Vantagens:**
- Route group `(authenticated)` não afeta as URLs (transparente)
- Layout aplicado automaticamente em todas as páginas do grupo
- Proteção de rota centralizada (`useRequireAuth()`)
- Fácil adicionar novas páginas autenticadas

### Componentes

```
components/
├── layout/
│   └── collapsible-sidebar.tsx    ← Sidebar reutilizável
└── dashboard/
    └── sidebar.tsx                 ← Antigo (não usado mais)
```

---

## 💻 Uso

### Para o Usuário Final

1. **Expandir/Colapsar:**
   - Clique no botão `<` no topo da sidebar (desktop)
   - Ou use o atalho **⌘/Ctrl + B**

2. **Mobile:**
   - Toque no ícone ☰ no canto superior esquerdo
   - Toque fora da sidebar para fechar

3. **Navegação:**
   - Clique em qualquer item do menu
   - Item ativo fica destacado em azul

### Para Desenvolvedores

#### Adicionar Nova Página Autenticada

```tsx
// Basta criar em apps/web/app/(authenticated)/nova-pagina/page.tsx
export default function NovaPaginaPage() {
  return (
    <div>
      {/* A sidebar já está disponível automaticamente */}
      <h1>Minha Nova Página</h1>
    </div>
  );
}
```

Não precisa:
- ❌ Importar sidebar manualmente
- ❌ Adicionar layout específico
- ❌ Configurar proteção de rota
- ✅ Tudo já está no layout do route group!

#### Customizar Sidebar

Editar: `apps/web/components/layout/collapsible-sidebar.tsx`

```tsx
const navigation = [
  { name: "Dashboard", href: "/dashboard", icon: Home },
  { name: "Pacientes", href: "/patients", icon: Users },
  // Adicionar novo item aqui
  { name: "Nova Feature", href: "/nova-feature", icon: Star },
];
```

---

## 🎨 Design System

### Cores e Estados

| Estado | Classe CSS | Cor |
|--------|-----------|-----|
| **Ativo** | `bg-primary text-primary-foreground` | Azul primário |
| **Hover** | `hover:bg-accent hover:text-accent-foreground` | Cinza claro |
| **Normal** | `text-muted-foreground` | Cinza médio |

### Animações

```tsx
// Hover nos itens
whileHover={{ x: 4 }}

// Click feedback
whileTap={{ scale: 0.98 }}

// Expansão/colapso
animate={{ width: isCollapsed ? 80 : 256 }}
transition={{ duration: 0.3, ease: "easeInOut" }}
```

### Ícones

- **Lucide React** para todos os ícones
- Tamanho padrão: `h-5 w-5`
- Tamanho colapsado: `h-6 w-6` (maior para compensar falta de texto)

---

## 📱 Comportamento Responsivo

### Breakpoints

| Tela | Comportamento |
|------|---------------|
| **≥1024px** (Desktop) | Sidebar sempre visível, pode colapsar |
| **<1024px** (Tablet/Mobile) | Sidebar em overlay, botão hamburguer |

### Layout Dinâmico

```tsx
// Desktop: Margin adapta ao estado da sidebar
marginLeft: isCollapsed ? "80px" : "256px"

// Mobile: Sem margin, overlay sobre conteúdo
marginLeft: "0px"
```

### Padding do Conteúdo

```tsx
// Mobile: Padding-top maior para não cobrir menu hamburguer
className="p-4 pt-16 lg:pt-6 lg:p-8"
```

---

## 🔍 Detalhes Técnicos

### Eventos Customizados

```tsx
// Sidebar emite evento quando alterna
window.dispatchEvent(new Event("sidebar-toggle"));

// Layout escuta e atualiza margin
window.addEventListener("sidebar-toggle", handleSidebarToggle);
```

**Benefícios:**
- Margin atualiza **instantaneamente** sem reload
- Sincroniza estado entre sidebar e layout
- Funciona mesmo se usuário muda localStorage diretamente

### Prevenção de Hydration Mismatch

```tsx
const [isMounted, setIsMounted] = useState(false);

useEffect(() => {
  setIsMounted(true);
}, []);

// Só acessa localStorage após montar
marginLeft: isMounted && window.innerWidth >= 1024 ? ... : "0px"
```

### Proteção de Rota

```tsx
export default function AuthenticatedLayout({ children }) {
  useRequireAuth(); // Redireciona para /login se não autenticado

  return (
    <div>
      <CollapsibleSidebar />
      <main>{children}</main>
    </div>
  );
}
```

---

## 🚀 Melhores Práticas 2026

### ✅ O que fizemos bem

1. **Route Groups** - Organização clara sem afetar URLs
2. **Persistência de Estado** - UX melhorada
3. **Atalhos de Teclado** - Acessibilidade e produtividade
4. **Responsividade First** - Funciona em todos os dispositivos
5. **Tooltips** - UI clean quando colapsado
6. **Animações Suaves** - Framer Motion para performance
7. **Radix UI** - Componentes acessíveis (WCAG 2.1)
8. **localStorage** - Não precisa backend para preferências
9. **Event-driven** - Comunicação desacoplada entre componentes

### 🎯 Padrões Seguidos

- **Don't Repeat Yourself (DRY):** Layout centralizado
- **Single Responsibility:** Sidebar só cuida da navegação
- **Separation of Concerns:** Layout vs. Conteúdo
- **Progressive Enhancement:** Funciona sem JS (SSR)
- **Mobile First:** Design responsivo desde o início

---

## 📊 Comparação: Antes vs. Depois

| Aspecto | ❌ Antes | ✅ Depois |
|---------|----------|-----------|
| **Sidebar** | Só em dashboard | Todas as páginas |
| **Colapsável** | Não | Sim, com animações |
| **Atalho Teclado** | Não | ⌘/Ctrl + B |
| **Responsivo** | Parcial | Completo |
| **Estado Persistente** | Não | localStorage |
| **Tooltips** | Não | Sim (quando colapsado) |
| **Route Organization** | Flat | Route Groups |
| **Proteção Centralizada** | Não | useRequireAuth() |

---

## 🔧 Troubleshooting

### Problema: Sidebar não aparece

**Solução:**
1. Verificar se página está em `app/(authenticated)/`
2. Verificar console do navegador para erros
3. Limpar localStorage: `localStorage.clear()`

### Problema: Margin do conteúdo incorreta

**Solução:**
1. Verificar se `isMounted` está true
2. Abrir DevTools → Application → Local Storage → `sidebar-collapsed`
3. Forçar recalculo: Alternar sidebar 2x

### Problema: Atalho ⌘+B não funciona

**Solução:**
1. Verificar se há outros atalhos conflitantes (extensões do browser)
2. Testar em janela anônima
3. Verificar console: `addEventListener("keydown")` registrado?

---

## 📚 Referências

**Documentação:**
- [Next.js Route Groups](https://nextjs.org/docs/app/building-your-application/routing/route-groups)
- [Framer Motion](https://www.framer.com/motion/)
- [Radix UI Tooltip](https://www.radix-ui.com/primitives/docs/components/tooltip)
- [Tailwind CSS Responsive Design](https://tailwindcss.com/docs/responsive-design)

**Arquivos Modificados:**
- `apps/web/components/layout/collapsible-sidebar.tsx` (criado)
- `apps/web/app/(authenticated)/layout.tsx` (criado)
- `apps/web/app/(authenticated)/dashboard/page.tsx` (ajustado padding)

**Arquivos Movidos:**
- Todas as páginas autenticadas para `app/(authenticated)/`

---

## 🎓 Lições Aprendidas

### 1. Route Groups são Poderosos
- Permitem organização sem afetar URLs
- Layouts podem ser aplicados em grupos de páginas
- Facilita manutenção e escalabilidade

### 2. Estado Local vs. Global
- Preferências visuais: **localStorage** (não precisa backend)
- Dados de negócio: **React Query** (cache + sync)

### 3. Animações Performáticas
- Framer Motion usa GPU acceleration
- `transform` e `opacity` são cheap
- `width` pode causar reflow, mas 300ms é aceitável

### 4. Responsividade é Complexa
- Mobile != Desktop colapsado
- Considerar touch targets (min 44x44px)
- Overlay vs. Push content (escolhemos overlay mobile)

---

## 📈 Próximos Passos (Opcional)

**Melhorias Futuras:**

1. **Temas Customizáveis**
   - Light/Dark mode toggle na sidebar
   - Cores customizáveis por usuário

2. **Favoritos**
   - Pin items mais usados no topo
   - Salvar no backend (user_preferences table)

3. **Breadcrumbs**
   - Navegação hierárquica
   - Ex: Dashboard > Pacientes > João Silva

4. **Busca Global**
   - Cmd+K para busca
   - Navegar por páginas, pacientes, artigos

5. **Notificações**
   - Badge com contadores
   - Ex: "3 novos exames"

6. **Analytics**
   - Track quais páginas mais acessadas
   - Otimizar ordem dos itens

---

## ✅ Checklist de Implementação

- [x] Criar componente CollapsibleSidebar
- [x] Criar route group (authenticated)
- [x] Criar layout autenticado
- [x] Mover todas as páginas para route group
- [x] Implementar persistência (localStorage)
- [x] Adicionar atalho de teclado
- [x] Implementar tooltips
- [x] Responsividade mobile
- [x] Animações suaves
- [x] Indicadores visuais (active, hover)
- [x] Avatar e badge de role
- [x] Botão logout
- [x] Testar em todos os breakpoints
- [x] Documentar implementação
- [x] Commit com mensagem descritiva

---

**Status:** ✅ Completo
**Data:** 30 de Janeiro de 2026
**Tempo de Implementação:** ~1 hora
**Arquivos Modificados:** 27
**Linhas Adicionadas:** +2102
**Linhas Removidas:** -40

---

**Desenvolvido com:** Claude Sonnet 4.5 ✨
