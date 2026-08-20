'use client'

import { useParams, useRouter } from 'next/navigation'
import { useQuery } from '@tanstack/react-query'
import { AlertCircle, Loader2 } from 'lucide-react'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { PrescriptionDetail } from '@/components/prescriptions/prescription-detail'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { getPrescription } from '@/lib/api/prescriptions'

export default function PrescriptionViewPage() {
  const router = useRouter()
  const params = useParams()
  const prescriptionId = params.id as string

  // OBRIGATÓRIO: verificar paciente selecionado
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()

  const { data: prescription, isLoading } = useQuery({
    queryKey: ['prescription', prescriptionId],
    queryFn: () => getPrescription(prescriptionId),
    enabled: !!prescriptionId,
  })

  if (loadingPatient || isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
      </div>
    )
  }

  // Prontuário é sempre isolado pelo paciente selecionado: abrir por URL a receita de outro
  // paciente não pode mostrar conteúdo clínico.
  const belongsToSelected = prescription && prescription.patientId === selectedPatient?.id

  if (!prescription || !belongsToSelected) {
    return (
      <div className="container mx-auto max-w-4xl py-8">
        <SelectedPatientHeader />
        <Alert variant="destructive" className="mt-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Prescrição não encontrada</AlertTitle>
          <AlertDescription>
            A prescrição não existe, foi removida, ou pertence a outro paciente.
          </AlertDescription>
        </Alert>
        <Button onClick={() => router.push('/prescriptions')} className="mt-4">
          Voltar para prescrições
        </Button>
      </div>
    )
  }

  return (
    <div className="container mx-auto max-w-5xl py-8">
      <SelectedPatientHeader />

      <div className="mb-6">
        <h1 className="text-3xl font-bold">Prescrição</h1>
        <p className="mt-2 text-muted-foreground">
          Detalhes da prescrição, download do PDF e envio ao paciente.
        </p>
      </div>

      <PrescriptionDetail prescription={prescription} />
    </div>
  )
}
