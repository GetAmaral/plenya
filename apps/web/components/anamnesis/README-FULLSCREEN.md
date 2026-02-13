# Modo Fullscreen para Anamnese - Otimizado para Tablets

## Visão Geral

Versão otimizada do formulário de anamnese para uso durante consultas em tablets. Design limpo, touch-friendly e com foco em reduzir carga cognitiva do profissional.

## Características

### 🎯 Design Principles

Baseado em pesquisas sobre UX médico em tablets (2026):

1. **Clean & Calm Design** - Muito white space, sem poluição visual
2. **Touch-First** - Botões e áreas clicáveis grandes (min 44x44px)
3. **Cognitive Load Reduzido** - Seções expansíveis, foco em uma área por vez
4. **Quick Data Entry** - Campos otimizados para entrada rápida
5. **AI-Powered Templates** - Templates inteligentes pré-configurados
6. **Persistent Feedback** - Status visual constante do progresso

### 📱 Layout

- **Fullscreen opcional** - Botão para entrar/sair de tela cheia
- **Single column** - Otimizado para tablet landscape (1024px+)
- **Seções collapsible** - Accordion para organizar informação
- **Template picker modal** - Não ocupa espaço fixo na tela
- **Quick info bar** - Resumo visual no topo (data, template, visibilidade)

### 🎨 Diferenças vs Formulário Normal

| Feature | Normal | Fullscreen |
|---------|--------|------------|
| Layout | 3 colunas (2+1 sidebar) | 1 coluna wide |
| Template Picker | Sidebar fixa | Modal on-demand |
| Seções | Sempre visíveis | Collapsible accordion |
| Touch Targets | Padrão (36px) | Large (48px+) |
| Level Buttons | Small pills | Large cards (80px height) |
| Typography | Standard | Larger (text-base) |
| Spacing | Normal | Extra spacious |
| Navigation | Scroll livre | Seção por seção |

### 🧩 Componentes

#### 1. `AnamnesisFullscreenForm.tsx`

Formulário principal otimizado:

- **Collapsible sections:**
  - Data e Template
  - Resumo (com RichTextEditor)
  - Conteúdo Detalhado
  - Detalhes Adicionais (visibilidade + observações)
  - Items do Template

- **Features:**
  - Template picker modal (Dialog)
  - Quick info bar no topo
  - Status badges para progresso
  - Rich text toggle único
  - Botões de ação grandes

#### 2. `AnamnesisFullscreenTemplateItems.tsx`

Items de template otimizados para touch:

- **Level buttons:** Grid 2-3 colunas, 80px height
- **Color coding:** Visual claro por nível (red, orange, yellow, blue, green)
- **Check indicators:** Visual feedback de preenchimento
- **Progress badges:** Contador de items preenchidos
- **Larger textareas:** 3 rows padrão
- **Card-based layout:** Cada item em card destacado

#### 3. `/anamnesis/fullscreen/page.tsx`

Página wrapper:

- Header minimalista
- Botão fullscreen toggle
- Info do paciente selecionado
- Navegação de volta

### 🚀 Como Usar

#### Para o Usuário

1. **Acessar:** Clicar em "Modo Tablet" na lista de anamneses
2. **Fullscreen (opcional):** Clicar em "Tela Cheia" no header
3. **Preencher:** Expandir seções conforme necessário
4. **Template:** Clicar para abrir modal de templates
5. **Salvar:** Botão grande no final

#### Para Desenvolvedores

```tsx
// Página fullscreen
import { AnamnesisFullscreenForm } from '@/components/anamnesis/AnamnesisFullscreenForm'

<AnamnesisFullscreenForm
  onSuccess={() => router.push('/anamnesis')}
  onCancel={() => router.push('/anamnesis')}
/>
```

### 📐 Breakpoints

- **Desktop:** max-w-4xl (768px) centered
- **Tablet Landscape:** 1024px+ ideal
- **Touch targets:** min 48px (WCAG AAA)
- **Font sizes:** text-base (16px) minimum

### ♿ Acessibilidade

- [x] Touch targets 48px+ (WCAG AAA)
- [x] Keyboard navigation (useFormNavigation)
- [x] Focus indicators claros
- [x] Semantic HTML (sections, headers)
- [x] ARIA labels em dialogs
- [x] Color contrast > 4.5:1

### 🔄 Fluxo de Uso

```
1. Lista Anamneses → Clicar "Modo Tablet"
   ↓
2. Fullscreen Page → (opcional) Tela cheia
   ↓
3. Seção Data/Template → Expandida por padrão
   ↓
4. Seção Resumo → Expandir para preencher
   ↓
5. Seção Conteúdo → Expandir se necessário
   ↓
6. Seção Detalhes → Visibilidade + Observações
   ↓
7. Seção Items → Se template selecionado
   ↓
8. Salvar → Volta para lista
```

### 🎯 Performance

- **Lazy loading:** RichTextEditor apenas quando ativo
- **Optimistic updates:** Feedback imediato em level selection
- **Debounced onChange:** Template items atualizam em RAF
- **Memoization:** Organized data calculado uma vez

### 🔮 Melhorias Futuras

- [ ] Auto-save a cada 30s
- [ ] Voice-to-text para observações
- [ ] Shortcuts de teclado (Ctrl+1-5 para seções)
- [ ] Swipe gestures para navegação
- [ ] Modo offline (PWA)
- [ ] Template suggestions baseado em IA
- [ ] Histórico de rascunhos
- [ ] Split-screen com prontuário do paciente

### 📚 Referências

- [Healthcare UI Design 2026: Best Practices](https://www.eleken.co/blog-posts/user-interface-design-for-healthcare-applications)
- [Form UX Design Best Practices 2026](https://www.designstudiouiux.com/blog/form-ux-design-best-practices/)
- [Top 10 UX trends shaping digital healthcare in 2026](https://www.uxstudioteam.com/ux-blog/healthcare-ux)

---

**Última atualização:** 2026-02-08
**Autor:** Claude Sonnet 4.5 (com pesquisa de UX médico 2026)
