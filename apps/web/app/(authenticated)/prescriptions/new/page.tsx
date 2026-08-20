'use client'

import { useEffect, useState } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import { useMutation, useQuery } from '@tanstack/react-query'
import { toast } from 'sonner'
import { FileCheck, Loader2 } from 'lucide-react'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { CommercialPrescriptionForm } from '@/components/prescriptions/commercial-prescription-form'
import { CompoundedPrescriptionForm } from '@/components/prescriptions/compounded-prescription-form'
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import {
  createPrescription,
  getPrescription,
  openPrescriptionPdf,
  signPrescription,
  type SignPrescriptionResponse,
} from '@/lib/api/prescriptions'
import {
  emptyCommercialPrescription,
  toCreatePayload,
  fromPrescription,
  type CommercialPrescriptionFormData,
} from '@/lib/validations/prescription'
import {
  emptyCompoundedPrescription,
  fromCompoundedPrescription,
  toCompoundedPayload,
  type CompoundedPrescriptionFormData,
} from '@/lib/validations/compounded-prescription'

export default function NewPrescriptionPage() {
  const router = useRouter()
  const searchParams = useSearchParams()
  const [signedPrescriptionId, setSignedPrescriptionId] = useState<string | null>(null)
  const [signResult, setSignResult] = useState<SignPrescriptionResponse | null>(null)

  // OBRIGATÓRIO: verificar paciente selecionado
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()

  // Repetir receita: `?from=<id>` carrega uma prescrição anterior como ponto de partida. Paciente
  // de retorno costuma sair com a mesma receita; isso resolve o caso mais frequente em 2 cliques.
  const fromId = searchParams.get('from')
  const { data: fonte, isLoading: loadingSource } = useQuery({
    // A chave leva o paciente junto: sem isso, trocar de paciente com a tela aberta deixava a
    // receita do anterior no formulário, e o submit a carimbava com o id do novo.
    queryKey: ['prescription', fromId, selectedPatient?.id],
    queryFn: () => getPrescription(fromId as string),
    enabled: !!fromId && !!selectedPatient?.id,
  })
  // Toda tela com dado de paciente escopa pelo paciente selecionado.
  const source = fonte && fonte.patientId === selectedPatient?.id ? fonte : undefined

  // O modo vive na URL: a casca (guarda de paciente, mutations, tela de sucesso) é a mesma nos
  // dois, então rota própria para manipulado seria a quarta cópia do mesmo arquivo. Repetir uma
  // receita manipulada também entra no modo certo, mesmo sem o parâmetro na URL.
  const isCompounded = searchParams.get('tipo') === 'manipulado' || source?.type === 'compounded'

  const [defaults, setDefaults] = useState<CommercialPrescriptionFormData | null>(null)
  const [compoundedDefaults, setCompoundedDefaults] =
    useState<CompoundedPrescriptionFormData | null>(null)
  // defaultValues do react-hook-form só valem na montagem; trocar de origem (nova × repetir ×
  // depois de assinar) precisa remontar o formulário.
  const [formKey, setFormKey] = useState(0)

  useEffect(() => {
    if (!source) return
    if (source.type === 'compounded') {
      setCompoundedDefaults(fromCompoundedPrescription(source))
    } else {
      setDefaults(fromPrescription(source))
    }
    setFormKey((n) => n + 1)
    toast.info('Receita anterior carregada', {
      description: 'Revise os itens antes de assinar. A data passa a ser a de hoje.',
    })
  }, [source])

  const buildPayload = (data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData) => {
    if (!selectedPatient?.id) throw new Error('Nenhum paciente selecionado')
    return 'formulas' in data
      ? toCompoundedPayload(data, selectedPatient.id)
      : toCreatePayload(data, selectedPatient.id)
  }

  const saveDraftMutation = useMutation({
    mutationFn: async (data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData) =>
      createPrescription(buildPayload(data)),
    onSuccess: (prescription) => {
      toast.success('Rascunho salvo', { description: 'Você pode editar e assinar depois.' })
      router.push(`/prescriptions/${prescription.id}/edit`)
    },
    onError: (error: any) => {
      toast.error('Erro ao salvar rascunho', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  const createAndSignMutation = useMutation({
    mutationFn: async (data: CommercialPrescriptionFormData | CompoundedPrescriptionFormData) => {
      const prescription = await createPrescription(buildPayload(data))
      const result = await signPrescription(prescription.id)
      return { prescription, result }
    },
    onSuccess: ({ prescription, result }) => {
      setSignedPrescriptionId(prescription.id)
      setSignResult(result)
      toast.success(
        result.signatureMode === 'manual'
          ? 'Receita gerada para impressão'
          : 'Prescrição criada e assinada',
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
      toast.error('Erro ao criar prescrição', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  if (loadingPatient || (fromId && loadingSource)) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
      </div>
    )
  }

  if (signedPrescriptionId) {
    const isManual = signResult?.signatureMode === 'manual'
    return (
      <div className="container mx-auto max-w-3xl py-8">
        <SelectedPatientHeader />

        <Alert className="mb-6 border-green-500 bg-green-50 dark:bg-green-950">
          <FileCheck className="h-4 w-4 text-green-600" />
          <AlertTitle className="text-green-600">
            {isManual ? 'Receita gerada para impressão' : 'Prescrição criada com sucesso'}
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
          <CardContent className="space-y-4">
            <div className="flex flex-col gap-3 sm:flex-row sm:gap-4">
              <Button
                onClick={() =>
                  openPrescriptionPdf(signedPrescriptionId).catch((e) =>
                    toast.error(e?.message ?? 'Falha ao baixar o PDF')
                  )
                }
                className="flex-1"
              >
                {isManual ? 'Baixar para imprimir' : 'Baixar PDF assinado'}
              </Button>
              <Button variant="outline" onClick={() => router.push('/prescriptions')}>
                Ver todas as prescrições
              </Button>
            </div>
            <Button
              variant="ghost"
              className="w-full"
              onClick={() => {
                setSignedPrescriptionId(null)
                setSignResult(null)
                setDefaults(emptyCommercialPrescription())
                setCompoundedDefaults(emptyCompoundedPrescription())
                setFormKey((n) => n + 1)
              }}
            >
              Criar nova prescrição
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
        <h1 className="text-3xl font-bold">Nova prescrição</h1>
        <p className="mt-2 text-muted-foreground">
          {isCompounded
            ? 'Fórmulas magistrais com componentes, veículo e quantidade a aviar.'
            : 'Medicamentos industrializados, do catálogo ou digitados à mão.'}{' '}
          Você pode salvar como rascunho ou assinar digitalmente com certificado ICP-Brasil.
        </p>
      </div>

      {/* O tipo muda o escopo inteiro da receita e é imutável depois de criada; por isso a
          escolha fica aqui, antes de preencher, e some na edição. */}
      <Tabs
        value={isCompounded ? 'manipulado' : 'industrializado'}
        onValueChange={(v) =>
          router.replace(v === 'manipulado' ? '/prescriptions/new?tipo=manipulado' : '/prescriptions/new')
        }
        className="mb-6"
      >
        <TabsList>
          <TabsTrigger value="industrializado">Industrializado</TabsTrigger>
          <TabsTrigger value="manipulado">Manipulado</TabsTrigger>
        </TabsList>
      </Tabs>

      {isCompounded ? (
        <CompoundedPrescriptionForm
          key={`c-${formKey}`}
          patientId={selectedPatient?.id}
          defaultValues={compoundedDefaults ?? undefined}
          onSubmit={(data) => createAndSignMutation.mutate(data)}
          onSaveDraft={(data) => saveDraftMutation.mutate(data)}
          isSigning={createAndSignMutation.isPending}
          isSavingDraft={saveDraftMutation.isPending}
        />
      ) : (
        <CommercialPrescriptionForm
          key={formKey}
          patientId={selectedPatient?.id}
          defaultValues={defaults ?? undefined}
          onSubmit={(data) => createAndSignMutation.mutate(data)}
          onSaveDraft={(data) => saveDraftMutation.mutate(data)}
          isSigning={createAndSignMutation.isPending}
          isSavingDraft={saveDraftMutation.isPending}
        />
      )}
    </div>
  )
}
