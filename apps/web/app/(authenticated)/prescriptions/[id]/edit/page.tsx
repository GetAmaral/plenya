'use client'

import { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import { AlertCircle, FileCheck, Loader2 } from 'lucide-react'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { CommercialPrescriptionForm } from '@/components/prescriptions/commercial-prescription-form'
import { CompoundedPrescriptionForm } from '@/components/prescriptions/compounded-prescription-form'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import {
  getPrescription,
  openPrescriptionPdf,
  signPrescription,
  updatePrescription,
  type SignPrescriptionResponse,
} from '@/lib/api/prescriptions'
import {
  fromPrescription,
  toCreatePayload,
  type CommercialPrescriptionFormData,
} from '@/lib/validations/prescription'
import {
  fromCompoundedPrescription,
  toCompoundedPayload,
  type CompoundedPrescriptionFormData,
} from '@/lib/validations/compounded-prescription'

export default function EditPrescriptionPage() {
  const router = useRouter()
  const params = useParams()
  const queryClient = useQueryClient()
  const prescriptionId = params.id as string
  const [signResult, setSignResult] = useState<SignPrescriptionResponse | null>(null)

  // OBRIGATÓRIO: verificar paciente selecionado
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()

  const { data: prescription, isLoading: loadingPrescription } = useQuery({
    queryKey: ['prescription', prescriptionId],
    queryFn: () => getPrescription(prescriptionId),
    enabled: !!prescriptionId,
  })

  // Receita assinada não se edita. A visualização é a tela de leitura, com download, envio e
  // "repetir receita" — antes esse modo somente-leitura morava aqui dentro.
  const isSigned = !!prescription?.signedPdfPath
  useEffect(() => {
    if (isSigned && !signResult) router.replace(`/prescriptions/${prescriptionId}`)
  }, [isSigned, signResult, router, prescriptionId])

  const invalidate = () => {
    queryClient.invalidateQueries({ queryKey: ['prescriptions'] })
    queryClient.invalidateQueries({ queryKey: ['prescription', prescriptionId] })
  }

  // Payload de edição: mesmo builder da criação, sem os campos que não se atualizam (paciente e
  // data da prescrição são do momento em que ela nasceu).
  const buildUpdatePayload = (
    data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData
  ) => {
    if (!selectedPatient?.id) throw new Error('Nenhum paciente selecionado')
    const full =
      'formulas' in data
        ? toCompoundedPayload(data, selectedPatient.id)
        : toCreatePayload(data, selectedPatient.id)
    const { patientId: _ignored, prescriptionDate: _date, type: _type, ...payload } = full
    // Apagar as instruções gerais tem que apagar de verdade. O payload de criação manda `undefined`
    // quando o campo está vazio, a chave some do JSON e o Update só atribui ponteiro não nulo:
    // o médico limpava o texto, recebia "salvo" e o texto antigo continuava no banco e no PDF.
    return { ...payload, generalInstructions: data.generalInstructions?.trim() ?? '' }
  }

  const updateMutation = useMutation({
    mutationFn: async (data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData) =>
      updatePrescription(prescriptionId, buildUpdatePayload(data)),
    onSuccess: () => {
      toast.success('Prescrição atualizada')
      invalidate()
      router.push('/prescriptions')
    },
    onError: (error: any) => {
      toast.error('Erro ao atualizar prescrição', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  // Salvar e assinar: grava as alterações do rascunho ANTES de assinar. Assinar direto
  // congelaria o que estava no banco, não o que está na tela.
  const saveAndSignMutation = useMutation({
    mutationFn: async (data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData) => {
      await updatePrescription(prescriptionId, buildUpdatePayload(data))
      return signPrescription(prescriptionId)
    },
    onSuccess: (result) => {
      setSignResult(result)
      invalidate()
      toast.success(
        result.signatureMode === 'manual'
          ? 'Receita gerada para impressão'
          : 'Prescrição assinada',
        {
          description:
            result.message ||
            (result.sncrNumber
              ? `SNCR: ${result.sncrNumber}`
              : 'PDF assinado digitalmente com certificado ICP-Brasil'),
        }
      )
    },
    onError: (error: any) => {
      toast.error('Erro ao assinar prescrição', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  if (loadingPatient || loadingPrescription) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
      </div>
    )
  }

  const belongsToSelected = prescription && prescription.patientId === selectedPatient?.id
  const isCompoundedPrescription = prescription?.type === 'compounded'

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

  if (signResult) {
    const isManual = signResult.signatureMode === 'manual'
    return (
      <div className="container mx-auto max-w-3xl py-8">
        <SelectedPatientHeader />

        <Alert className="mb-6 border-green-500 bg-green-50 dark:bg-green-950">
          <FileCheck className="h-4 w-4 text-green-600" />
          <AlertTitle className="text-green-600">
            {isManual ? 'Receita gerada para impressão' : 'Prescrição assinada'}
          </AlertTitle>
          <AlertDescription className="text-green-600">
            {isManual
              ? 'Imprima, carimbe e assine à mão. Receita de medicamento controlado segue o receituário físico.'
              : 'A prescrição foi assinada digitalmente e está pronta para download.'}
          </AlertDescription>
        </Alert>

        <Card>
          <CardHeader>
            <CardTitle>
              {isManual ? 'Receita para impressão' : 'PDF assinado digitalmente'}
            </CardTitle>
            <CardDescription>
              {isManual
                ? 'Documento para assinatura e carimbo do médico (sem assinatura digital).'
                : 'Prescrição com assinatura ICP-Brasil e QR Code de validação.'}
            </CardDescription>
          </CardHeader>
          <CardContent className="flex flex-col gap-3 sm:flex-row sm:gap-4">
            <Button
              className="flex-1"
              onClick={() =>
                openPrescriptionPdf(prescriptionId).catch((e) =>
                  toast.error(e?.message ?? 'Falha ao baixar o PDF')
                )
              }
            >
              {isManual ? 'Baixar para imprimir' : 'Baixar PDF assinado'}
            </Button>
            <Button variant="outline" onClick={() => router.push('/prescriptions')}>
              Ver todas as prescrições
            </Button>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="container mx-auto max-w-5xl py-8">
      <SelectedPatientHeader />

      <div className="mb-6">
        <h1 className="text-3xl font-bold">
          Editar {isCompoundedPrescription ? 'receita de manipulado' : 'prescrição'} (rascunho)
        </h1>
        <p className="mt-2 text-muted-foreground">
          Atualize os dados da prescrição. Você pode salvar as alterações ou assinar digitalmente.
        </p>
      </div>

      {/* Não há switch aqui: o tipo é imutável. Trocar o tipo de uma receita existente
          misturaria os dois receituários; para mudar, crie outra. */}
      {isCompoundedPrescription ? (
        <CompoundedPrescriptionForm
          patientId={selectedPatient?.id}
          defaultValues={fromCompoundedPrescription(prescription)}
          onSubmit={(data) => saveAndSignMutation.mutate(data)}
          onSaveDraft={(data) => updateMutation.mutate(data)}
          isSigning={saveAndSignMutation.isPending}
          isSavingDraft={updateMutation.isPending}
          draftLabel="Salvar alterações"
        />
      ) : (
        <CommercialPrescriptionForm
          patientId={selectedPatient?.id}
          defaultValues={fromPrescription(prescription)}
          onSubmit={(data) => saveAndSignMutation.mutate(data)}
          onSaveDraft={(data) => updateMutation.mutate(data)}
          isSigning={saveAndSignMutation.isPending}
          isSavingDraft={updateMutation.isPending}
          draftLabel="Salvar alterações"
        />
      )}
    </div>
  )
}
