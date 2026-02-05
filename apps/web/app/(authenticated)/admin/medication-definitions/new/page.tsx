'use client'

import { useMutation } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

import { MedicationDefinitionForm } from '@/components/admin/MedicationDefinitionForm'
import { createMedicationDefinition } from '@/lib/api/medication-definitions'
import type { MedicationDefinitionInput } from '@/lib/validations/medication-definition'

export default function NewMedicationDefinitionPage() {
  const router = useRouter()

  const createMutation = useMutation({
    mutationFn: createMedicationDefinition,
    onSuccess: () => {
      toast.success('Medicamento criado com sucesso')
      router.push('/admin/medication-definitions')
    },
    onError: () => {
      toast.error('Erro ao criar medicamento')
    },
  })

  return (
    <div className="container mx-auto py-8 max-w-4xl">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Novo Medicamento</h1>
        <p className="text-muted-foreground mt-2">
          Cadastre um novo medicamento no sistema
        </p>
      </div>

      <MedicationDefinitionForm
        onSubmit={(data) => createMutation.mutate(data)}
        isSubmitting={createMutation.isPending}
        submitLabel="Criar Medicamento"
      />
    </div>
  )
}
