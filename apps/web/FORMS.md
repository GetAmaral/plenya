# Padrão de Formulários - Plenya EMR

## Navegação com Enter

Todos os formulários da aplicação seguem um padrão consistente de navegação:

- **Enter** em qualquer campo move para o próximo campo (como Tab)
- **Enter** no último campo faz submit do formulário
- **Tab** continua funcionando normalmente
- **Textarea** mantém comportamento padrão (Enter quebra linha)

## Como Implementar

### 1. Importar o Hook

```tsx
import { useFormNavigation } from "@/lib/use-form-navigation";
```

### 2. Criar Ref do Formulário

```tsx
import { useRef } from "react";

export default function MyForm() {
  const formRef = useRef<HTMLFormElement>(null);

  // ... resto do código
}
```

### 3. Aplicar o Hook

```tsx
export default function MyForm() {
  const formRef = useRef<HTMLFormElement>(null);

  // Navegação automática por Enter
  useFormNavigation({ formRef });

  // ... resto do código
}
```

### 4. Adicionar Ref ao Form

```tsx
<form ref={formRef} onSubmit={handleSubmit(onSubmit)}>
  {/* campos do formulário */}
</form>
```

## Exemplo Completo

```tsx
"use client";

import { useRef } from "react";
import { useForm } from "react-hook-form";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

interface MyFormData {
  name: string;
  email: string;
  phone: string;
}

export default function MyForm() {
  const formRef = useRef<HTMLFormElement>(null);

  // Navegação automática por Enter
  useFormNavigation({ formRef });

  const { register, handleSubmit } = useForm<MyFormData>();

  const onSubmit = (data: MyFormData) => {
    console.log(data);
  };

  return (
    <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <Label htmlFor="name">Nome</Label>
        <Input id="name" {...register("name")} />
      </div>

      <div>
        <Label htmlFor="email">Email</Label>
        <Input id="email" type="email" {...register("email")} />
      </div>

      <div>
        <Label htmlFor="phone">Telefone</Label>
        <Input id="phone" {...register("phone")} />
      </div>

      <Button type="submit">Salvar</Button>
    </form>
  );
}
```

## Opções Avançadas

### Desabilitar Navegação

```tsx
useFormNavigation({
  formRef,
  disabled: isLoading // Desabilita durante loading
});
```

### Callback Customizado no Último Campo

```tsx
useFormNavigation({
  formRef,
  onLastFieldEnter: () => {
    // Ação customizada ao invés de submit automático
    console.log("Último campo atingido!");
  }
});
```

## Comportamento Esperado

### ✅ Campos que Navegam com Enter
- `<input type="text">`
- `<input type="email">`
- `<input type="number">`
- `<input type="tel">`
- `<input type="date">`
- `<select>`
- Componentes shadcn/ui `<Select>` (role="combobox")

### ❌ Campos que Mantêm Enter Padrão
- `<textarea>` - Enter quebra linha
- `<button type="submit">` - Enter faz submit
- Campos desabilitados
- Campos hidden

## Formulários Já Implementados

- ✅ Login (`/login`)
- ✅ Edição de Paciente (`/patients/[id]/edit`)

## Próximos Formulários a Implementar

- [ ] Cadastro de Paciente
- [ ] Cadastro de Consulta
- [ ] Cadastro de Prescrição
- [ ] Cadastro de Anamnese
- [ ] Cadastro de Exame Laboratorial

## Notas Importantes

1. **Sempre adicione o ref ao form**: Sem o ref, o hook não funciona
2. **Use com React Hook Form**: Compatível e recomendado
3. **Funciona com shadcn/ui**: Testado com todos os componentes
4. **Performance**: Zero impacto, apenas um event listener por formulário
5. **Acessibilidade**: Não interfere com leitores de tela ou navegação por teclado

## Troubleshooting

### Enter não está navegando

- ✅ Verifique se o ref foi adicionado ao `<form>`
- ✅ Confirme que o hook foi chamado: `useFormNavigation({ formRef })`
- ✅ Certifique-se de que os campos não estão desabilitados

### Submit acontecendo duas vezes

- ✅ Remova `type="submit"` dos campos de input
- ✅ Certifique-se de ter apenas um botão submit por formulário

### Select do shadcn não funciona

- ✅ O componente deve ter `role="combobox"` (padrão do shadcn)
- ✅ Use o componente `<Select>` de `/components/ui/select`
