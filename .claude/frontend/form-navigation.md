# Frontend - Form Navigation (OBRIGATÓRIO)

## Regra de Ouro

**TODOS os formulários DEVEM usar `useFormNavigation`.**

## Comportamento

- **Auto-focus:** Foca automaticamente no primeiro campo quando formulário carrega
- **Enter em input:** Foca próximo campo
- **Enter no último campo:** Submete formulário
- **Tab:** Continua funcionando normalmente
- **Textareas:** Permitem quebra de linha (Enter não navega)

## Uso Básico

```tsx
import { useFormNavigation } from '@/lib/use-form-navigation'

function MyForm() {
  const formRef = useRef<HTMLFormElement>(null)

  useFormNavigation({ formRef })

  return (
    <form ref={formRef} onSubmit={handleSubmit}>
      <Input name="field1" />
      <Input name="field2" />
      <button type="submit">Salvar</button>
    </form>
  )
}
```

## Uso em Dialog/Sheet

```tsx
function MyDialog({ isOpen }: { isOpen: boolean }) {
  const formRef = useRef<HTMLFormElement>(null)

  // Desabilitar quando dialog fechado
  useFormNavigation({ formRef, disabled: !isOpen })

  return (
    <Dialog open={isOpen}>
      <form ref={formRef} onSubmit={handleSubmit}>
        {/* ... */}
      </form>
    </Dialog>
  )
}
```

## Opções Avançadas

```tsx
// Desabilitar auto-focus (raro)
useFormNavigation({ formRef, autoFocus: false })

// Desabilitar completamente
useFormNavigation({ formRef, disabled: true })
```

## Checklist

- [ ] Criar `formRef` com `useRef<HTMLFormElement>(null)`
- [ ] Passar `ref={formRef}` no `<form>`
- [ ] Chamar `useFormNavigation({ formRef })`
- [ ] Se Dialog/Sheet: `disabled: !isOpen`

Ver: `apps/web/lib/use-form-navigation.ts`
