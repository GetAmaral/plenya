# Plenya EMR - Design System

## 🎨 Overview

Design System moderno implementado em Janeiro 2026 para o Plenya EMR, seguindo os padrões mais avançados de UI/UX para aplicações médicas.

## 📐 Princípios de Design

1. **Clareza Visual** - Interface limpa e focada em conteúdo médico crítico
2. **Hierarquia Clara** - Informações importantes destacadas
3. **Feedback Imediato** - Micro-interactions e animações suaves
4. **Acessibilidade** - WCAG 2.1 AA compliant
5. **Performance** - Otimizado para carregamento rápido

## 🎯 Design Tokens

### Spacing Scale
Base de 4px para consistência:
- `xs`: 4px
- `sm`: 8px
- `md`: 16px
- `lg`: 24px
- `xl`: 32px

### Typography
- **Font Family**: Inter (sans-serif)
- **Scale**: Fluida com clamp() para responsividade
- **Weights**: 400 (normal), 500 (medium), 600 (semibold), 700 (bold)

### Border Radius
Moderna e arredondada:
- `sm`: 4px
- `DEFAULT`: 8px
- `md`: 12px
- `lg`: 16px
- `xl`: 20px
- `2xl`: 24px

## 🎨 Color System

### Base Colors
```css
--primary: 210 100% 50%        /* Trust Blue */
--secondary: 210 40% 96.1%     /* Subtle Gray */
--destructive: 0 84.2% 60.2%   /* Critical Red */
--radius: 0.75rem              /* Modern rounded */
```

### Medical Semantic Colors

#### Patient Status
- **Stable** (Verde): `hsl(142 76% 36%)` - Paciente estável
- **Observation** (Amarelo): `hsl(38 92% 50%)` - Requer observação
- **Critical** (Vermelho): `hsl(0 84% 60%)` - Estado crítico
- **Unknown** (Cinza): `hsl(220 13% 46%)` - Status desconhecido

#### Priority Levels
- **Urgent**: `hsl(0 84% 60%)` - Urgente
- **High**: `hsl(25 95% 53%)` - Alta prioridade
- **Normal**: `hsl(210 100% 50%)` - Prioridade normal
- **Low**: `hsl(220 13% 46%)` - Baixa prioridade

## 🧩 Components

### Badge (Medical)
Variantes específicas para contexto médico:
```tsx
<Badge variant="stable">Estável</Badge>
<Badge variant="observation">Observação</Badge>
<Badge variant="critical">Crítico</Badge>
<Badge variant="urgent">Urgente</Badge>
```

### Avatar
Com fallback de iniciais e suporte a imagem:
```tsx
<Avatar>
  <AvatarImage src="/user.jpg" />
  <AvatarFallback>JS</AvatarFallback>
</Avatar>
```

### Skeleton
Loading states modernos:
```tsx
<Skeleton className="h-12 w-full" />
```

## ✨ Animations

### Timing
- **Fast**: 150ms - Hover states
- **Normal**: 250ms - Transitions padrão
- **Slow**: 350ms - Page transitions
- **Slower**: 500ms - Complex animations

### Easing
- `default`: cubic-bezier(0.4, 0, 0.2, 1)
- Suave e natural

### Framer Motion Variants
```tsx
const container = {
  hidden: { opacity: 0 },
  show: {
    opacity: 1,
    transition: { staggerChildren: 0.1 }
  }
};

const item = {
  hidden: { opacity: 0, y: 20 },
  show: { opacity: 1, y: 0 }
};
```

## 🔧 Utilities

### Glassmorphism
```css
.glass {
  background: bg-white/70 dark:bg-gray-950/70
  backdrop-blur: xl
  border: white/20 dark:gray-800/50
}
```

### Gradients
```css
.gradient-primary - Blue gradient
.gradient-medical - Medical blue gradient
```

### Bento Grid
```css
.bento-grid - Responsive grid layout
.bento-item - Card com hover effects
.bento-item-large - 2 colunas
.bento-item-tall - 2 linhas
```

## 📱 Responsive Breakpoints

```tsx
sm: 640px   // Mobile landscape
md: 768px   // Tablet
lg: 1024px  // Desktop
xl: 1280px  // Large desktop
2xl: 1400px // Max container width
```

## 🎭 Patterns

### Bento Grid Layout
Layout moderno com cards de diferentes tamanhos:
- Cards pequenos: 1x1 (estatísticas)
- Cards médios: 2x1 (listas)
- Cards grandes: 2x2 (gráficos)

### Loading States
Skeleton screens ao invés de spinners genéricos

### Empty States
Ilustrações + mensagens + CTAs claros

### Toasts (Sonner)
```tsx
toast.success("Sucesso!", {
  description: "Operação realizada"
});

toast.error("Erro!", {
  description: "Algo deu errado"
});
```

## 🔐 Accessibility

- Focus states visíveis (ring-2 ring-ring)
- Contraste WCAG AA compliant
- Semantic HTML
- ARIA labels onde necessário
- Keyboard navigation

## 📚 References

- **Design Tokens**: `/lib/design-tokens.ts`
- **Tailwind Config**: `/tailwind.config.ts`
- **Global Styles**: `/app/globals.css`
- **Components**: `/components/ui/`

## 🚀 Usage

### Importar Componentes
```tsx
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
```

### Usar Design Tokens
```tsx
import { designTokens, token } from "@/lib/design-tokens";

const spacing = token("spacing");
// spacing.md = "1rem"
```

### Aplicar Utilities
```tsx
<div className="glass-card bento-item">
  <div className="gradient-medical text-gradient">
    Modern Design
  </div>
</div>
```

## 📝 Best Practices

1. **Sempre use design tokens** - Não hardcode valores
2. **Prefira utilities** - Use classes Tailwind ao invés de CSS custom
3. **Animações suaves** - 250ms padrão
4. **Mobile-first** - Sempre responsive
5. **Semantic HTML** - Acessibilidade primeiro
6. **Loading states** - Skeleton > Spinners
7. **Feedback visual** - Toast para ações importantes

---

**Última atualização**: Janeiro 2026
**Versão**: 1.0.0
