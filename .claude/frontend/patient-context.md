# Frontend - Patient Context (OBRIGATÓRIO)

## Conceito Fundamental

O sistema EMR opera no contexto de **UM paciente por vez**. O médico seleciona um paciente e então todas as telas mostram apenas dados desse paciente.

## Entidades que Pertencem a Paciente

- Anamnesis
- Lab Results
- Prescriptions
- Appointments
- Health Scores

## Implementação Obrigatória

**TODAS as páginas relacionadas a dados de paciente DEVEM usar:**

```tsx
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'

export default function MyPatientDataPage() {
  const { selectedPatient, isLoading } = useRequireSelectedPatient()

  // Se não tiver selectedPatient, hook redireciona automaticamente
  // para /patients/select e salva URL atual para voltar depois

  return (
    <div>
      <SelectedPatientHeader /> {/* Mostra qual paciente selecionado */}

      {/* Seus dados aqui - já filtrados pelo selectedPatient */}
    </div>
  )
}
```

## Fluxo de Uso

1. Usuário acessa `/lab-requests`
2. Sistema verifica: Tem `selectedPatient`?
   - ❌ Não → Redireciona para `/patients/select`
   - ✅ Sim → Mostra página com dados filtrados
3. Usuário seleciona paciente em `/patients/select`
4. Sistema salva `selectedPatient` no Zustand + localStorage
5. Sistema redireciona de volta para `/lab-requests`
6. Página carrega dados apenas do paciente selecionado

## Backend Automático

O backend usa middleware `WithSelectedPatient()` que:
- Lê `selectedPatient` do token JWT
- Preenche automaticamente `patientId` nas criações
- Filtra automaticamente queries pelo `patientId`

## Formulários NÃO Devem Ter Select de Paciente

❌ **ERRADO:**
```tsx
<Select value={patientId} onChange={setPatientId}>
  <SelectItem value="123">João Silva</SelectItem>
</Select>
```

✅ **CORRETO:**
```tsx
// Paciente já está selecionado no contexto global
// Formulário só precisa dos campos específicos
<Input name="chiefComplaint" placeholder="Queixa principal" />
```

## Exceções (Não Precisam de selectedPatient)

- `/patients` - Lista de todos os pacientes
- `/patients/select` - Seleção de paciente
- `/articles` - Artigos científicos
- `/lab-test-definitions` - Templates de exames
- `/score-groups` - Sistema de escores (não é de paciente)

Ver: `apps/web/lib/use-require-selected-patient.ts`
