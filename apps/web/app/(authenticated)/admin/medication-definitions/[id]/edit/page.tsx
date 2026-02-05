'use client'

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter, useParams } from 'next/navigation'
import { toast } from 'sonner'

import { MedicationDefinitionForm } from '@/components/admin/MedicationDefinitionForm'
import {
  getMedicationDefinition,
  updateMedicationDefinition,
} from '@/lib/api/medication-definitions'
import type { MedicationDefinitionInput } from '@/lib/validations/medication-definition'

export default function EditMedicationDefinitionPage() {
  const router = useRouter()
  const params = useParams()
  const queryClient = useQueryClient()
  const id = params.id as string

  const { data: medication, isLoading } = useQuery({
    queryKey: ['medication-definition', id],
    queryFn: () => getMedicationDefinition(id),
  })

  const updateMutation = useMutation({
    mutationFn: (data: MedicationDefinitionInput) =>
      updateMedicationDefinition(id, data),
    onSuccess: () => {
      toast.success('Medicamento atualizado com sucesso')
      queryClient.invalidateQueries({ queryKey: ['medication-definitions'] })
      router.push('/admin/medication-definitions')
    },
    onError: () => {
      toast.error('Erro ao atualizar medicamento')
    },
  })

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 max-w-4xl">
        <div className="text-center text-muted-foreground">Carregando...</div>
      </div>
    )
  }

  if (!medication) {
    return (
      <div className="container mx-auto py-8 max-w-4xl">
        <div className="text-center text-destructive">Medicamento não encontrado</div>
      </div>
    )
  }

  return (
    <div className="container mx-auto py-8 max-w-4xl">
      <div className="mb-6">
        <h1 className="text-3xl font-bold">Editar Medicamento</h1>
        <p className="text-muted-foreground mt-2">
          Atualize as informações do medicamento
        </p>
      </div>

      <MedicationDefinitionForm
        defaultValues={{
          commonName: medication.commonName,
          activeIngredient: medication.activeIngredient,
          category: medication.category,
          validityDays: medication.validityDays,
          maxPerPrescription: medication.maxPerPrescription,
          maxTreatmentDays: medication.maxTreatmentDays,
          requiresDigitalSignature: medication.requiresDigitalSignature,
          requiresSNCR: medication.requiresSNCR,
          anvisaCode: medication.anvisaCode,
        }}
        onSubmit={(data) => updateMutation.mutate(data)}
        isSubmitting={updateMutation.isPending}
        submitLabel="Atualizar Medicamento"
      />
    </div>
  )
}
