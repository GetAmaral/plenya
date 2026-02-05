'use client'

import { useRef, useState } from 'react'
import { useRouter } from 'next/navigation'
import { useForm, useFieldArray } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useMutation } from '@tanstack/react-query'
import { toast } from 'sonner'
import { z } from 'zod'
import { Loader2, FileCheck, AlertCircle, Plus, Trash2, Save } from 'lucide-react'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Separator } from '@/components/ui/separator'

import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { createPrescription, signPrescription } from '@/lib/api/prescriptions'
import { numberToWordsWithUnit } from '@/lib/utils/number-to-words'
import { MedicationSearch } from '@/components/prescriptions/MedicationSearch'
import type { MedicationDefinition } from '@/lib/api/medication-definitions'

// Medication schema
const medicationSchema = z.object({
  medicationDefinitionId: z.string().optional(),
  category: z.enum(['simple', 'c1', 'c5', 'antibiotic', 'glp1']),
  medicationName: z.string().min(3, 'Nome do medicamento é obrigatório'),
  activeIngredient: z.string().min(3, 'Princípio ativo é obrigatório'),
  concentration: z.string().min(1, 'Concentração é obrigatória'),
  dosage: z.string().min(1, 'Dosagem é obrigatória'),
  frequency: z.string().min(1, 'Frequência é obrigatória'),
  route: z.string().min(1, 'Via é obrigatória'),
  duration: z.coerce.number().min(1, 'Duração deve ser pelo menos 1 dia'),
  quantity: z.coerce.number().min(1, 'Quantidade deve ser pelo menos 1'),
  quantityInWords: z.string().min(1, 'Quantidade por extenso é obrigatória'),
  instructions: z.string().optional(),
})

// Prescription schema with multiple medications
const prescriptionSchema = z.object({
  medications: z.array(medicationSchema).min(1, 'Adicione pelo menos um medicamento').max(10, 'Máximo 10 medicamentos'),
  generalInstructions: z.string().optional(),
}).refine(
  (data) => {
    // Validar max 3 medicamentos C1
    const c1Count = data.medications.filter(m => m.category === 'c1').length
    return c1Count <= 3
  },
  {
    message: 'Prescrição não pode ter mais de 3 medicamentos controlados (C1)',
    path: ['medications'],
  }
)

type PrescriptionFormData = z.infer<typeof prescriptionSchema>

export default function NewPrescriptionPage() {
  const router = useRouter()
  const formRef = useRef<HTMLFormElement>(null)
  const [signedPdfUrl, setSignedPdfUrl] = useState<string | null>(null)

  // OBRIGATÓRIO: Verificar paciente selecionado
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()

  // React Hook Form with field array for medications
  const form = useForm<PrescriptionFormData>({
    resolver: zodResolver(prescriptionSchema),
    defaultValues: {
      medications: [
        {
          category: 'simple',
          medicationName: '',
          activeIngredient: '',
          concentration: '',
          dosage: '',
          frequency: '',
          route: 'oral',
          duration: 30,
          quantity: 30,
          quantityInWords: '',
          instructions: '',
        },
      ],
      generalInstructions: '',
    },
  })

  const { fields, append, remove } = useFieldArray({
    control: form.control,
    name: 'medications',
  })

  // OBRIGATÓRIO: Form navigation com Enter
  useFormNavigation({ formRef })

  // Mutation: Save Draft (without signing)
  const saveDraftMutation = useMutation({
    mutationFn: async (data: PrescriptionFormData) => {
      if (!selectedPatient?.id) {
        throw new Error('Nenhum paciente selecionado')
      }

      const prescription = await createPrescription({
        patientId: selectedPatient.id,
        medications: data.medications.map(m => ({
          medicationDefinitionId: m.medicationDefinitionId,
          medicationName: m.medicationName,
          activeIngredient: m.activeIngredient,
          category: m.category as any,
          concentration: m.concentration,
          dosage: m.dosage,
          frequency: m.frequency,
          route: m.route,
          duration: m.duration,
          quantity: m.quantity,
          quantityInWords: m.quantityInWords,
          instructions: m.instructions || undefined,
        })),
        generalInstructions: data.generalInstructions || undefined,
        prescriptionDate: new Date().toISOString(),
      })

      return prescription
    },
    onSuccess: (prescription) => {
      toast.success('Rascunho salvo com sucesso!', {
        description: 'Você pode editar e assinar depois.',
      })
      router.push(`/prescriptions/${prescription.id}/edit`)
    },
    onError: (error: any) => {
      toast.error('Erro ao salvar rascunho', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  // Mutation: Create + Sign
  const createAndSignMutation = useMutation({
    mutationFn: async (data: PrescriptionFormData) => {
      if (!selectedPatient?.id) {
        throw new Error('Nenhum paciente selecionado')
      }

      // 1. Criar prescrição
      const prescription = await createPrescription({
        patientId: selectedPatient.id,
        medications: data.medications.map(m => ({
          medicationDefinitionId: m.medicationDefinitionId,
          medicationName: m.medicationName,
          activeIngredient: m.activeIngredient,
          category: m.category as any,
          concentration: m.concentration,
          dosage: m.dosage,
          frequency: m.frequency,
          route: m.route,
          duration: m.duration,
          quantity: m.quantity,
          quantityInWords: m.quantityInWords,
          instructions: m.instructions || undefined,
        })),
        generalInstructions: data.generalInstructions || undefined,
        prescriptionDate: new Date().toISOString(),
      })

      // 2. Assinar e gerar PDF
      const signResult = await signPrescription(prescription.id)

      return { prescription, signResult }
    },
    onSuccess: ({ prescription, signResult }) => {
      setSignedPdfUrl(signResult.signedPdfUrl)
      toast.success('Prescrição criada e assinada com sucesso!', {
        description: signResult.sncrNumber
          ? `SNCR: ${signResult.sncrNumber}`
          : 'PDF assinado digitalmente com certificado ICP-Brasil',
      })
    },
    onError: (error: any) => {
      toast.error('Erro ao criar prescrição', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  const onSubmit = (data: PrescriptionFormData) => {
    createAndSignMutation.mutate(data)
  }

  const onSaveDraft = () => {
    const data = form.getValues()
    saveDraftMutation.mutate(data)
  }

  // Auto-preencher quantityInWords
  const handleQuantityChange = (index: number, value: string) => {
    const quantity = parseInt(value)
    if (!isNaN(quantity) && quantity > 0 && quantity <= 999) {
      const unit = 'comprimidos' // Pode ser dinâmico baseado no medicamento
      const words = numberToWordsWithUnit(quantity, unit)
      form.setValue(`medications.${index}.quantityInWords`, words)
    }
  }

  // Handler when medication is selected from search
  const handleMedicationSelect = (index: number, medication: MedicationDefinition) => {
    form.setValue(`medications.${index}.medicationDefinitionId`, medication.id)
    form.setValue(`medications.${index}.medicationName`, medication.commonName)
    form.setValue(`medications.${index}.activeIngredient`, medication.activeIngredient)
    form.setValue(`medications.${index}.category`, medication.category as any)
  }

  // Add new medication
  const addMedication = () => {
    append({
      category: 'simple',
      medicationName: '',
      activeIngredient: '',
      concentration: '',
      dosage: '',
      frequency: '',
      route: 'oral',
      duration: 30,
      quantity: 30,
      quantityInWords: '',
      instructions: '',
    })
  }

  // Count C1 medications
  const c1Count = form.watch('medications').filter(m => m.category === 'c1').length
  const hasControlled = form.watch('medications').some(m => m.category !== 'simple')

  if (loadingPatient) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
      </div>
    )
  }

  // Se já foi assinada com sucesso, mostrar resultado
  if (signedPdfUrl) {
    return (
      <div className="container mx-auto py-8 max-w-3xl">
        <SelectedPatientHeader />

        <Alert className="border-green-500 bg-green-50 dark:bg-green-950 mb-6">
          <FileCheck className="h-4 w-4 text-green-600" />
          <AlertTitle className="text-green-600">Prescrição Criada com Sucesso!</AlertTitle>
          <AlertDescription className="text-green-600">
            A prescrição foi assinada digitalmente e está pronta para download.
          </AlertDescription>
        </Alert>

        <Card>
          <CardHeader>
            <CardTitle>PDF Assinado Digitalmente</CardTitle>
            <CardDescription>
              Prescrição com assinatura ICP-Brasil e QR Code de validação.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="flex gap-4">
              <Button onClick={() => window.open(signedPdfUrl, '_blank')} className="flex-1">
                Baixar PDF Assinado
              </Button>
              <Button variant="outline" onClick={() => router.push('/prescriptions')}>
                Ver Todas Prescrições
              </Button>
            </div>
            <Button
              variant="ghost"
              onClick={() => {
                setSignedPdfUrl(null)
                form.reset()
              }}
              className="w-full"
            >
              Criar Nova Prescrição
            </Button>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="container mx-auto py-8 max-w-5xl">
      <SelectedPatientHeader />

      <div className="mb-6">
        <h1 className="text-3xl font-bold">Nova Prescrição Digital</h1>
        <p className="text-muted-foreground mt-2">
          Preencha os dados da prescrição. Você pode salvar como rascunho ou assinar digitalmente
          com certificado ICP-Brasil.
        </p>
      </div>

      {/* Warnings */}
      {hasControlled && (
        <Alert className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Medicamentos Controlados</AlertTitle>
          <AlertDescription>
            {c1Count > 0 && (
              <p className={c1Count > 3 ? 'text-destructive font-semibold' : ''}>
                {c1Count > 3 ? '❌' : '✅'} {c1Count} medicamento(s) C1 (máximo 3)
              </p>
            )}
            <p className="mt-2 text-sm">
              Controle Especial C1: Máximo 3 substâncias, 30 dias de validade.
              <br />
              Antimicrobiano: 10 dias de validade (RDC 471/2021).
              <br />
              GLP-1 Agonista: 90 dias de validade.
            </p>
          </AlertDescription>
        </Alert>
      )}

      <form ref={formRef} onSubmit={form.handleSubmit(onSubmit)}>
        {/* Medications */}
        <div className="space-y-6">
          {fields.map((field, index) => (
            <Card key={field.id}>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle>Medicamento {index + 1}</CardTitle>
                  {fields.length > 1 && (
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      onClick={() => remove(index)}
                      className="text-destructive hover:text-destructive"
                    >
                      <Trash2 className="h-4 w-4 mr-2" />
                      Remover
                    </Button>
                  )}
                </div>
              </CardHeader>
              <CardContent className="space-y-4">
                {/* Medication Search */}
                <div className="space-y-2">
                  <Label>Medicamento *</Label>
                  <MedicationSearch
                    value={undefined}
                    onSelect={(med) => handleMedicationSelect(index, med)}
                    placeholder="Buscar medicamento cadastrado..."
                  />
                  <p className="text-xs text-muted-foreground">
                    Busque por nome comercial ou princípio ativo. Se não encontrar, preencha
                    manualmente abaixo.
                  </p>
                </div>

                <Separator />

                {/* Manual inputs */}
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label>Nome Comercial *</Label>
                    <Input
                      placeholder="Ex: Losartana Potássica"
                      {...form.register(`medications.${index}.medicationName`)}
                    />
                    {form.formState.errors.medications?.[index]?.medicationName && (
                      <p className="text-sm text-destructive">
                        {form.formState.errors.medications[index]?.medicationName?.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label>Princípio Ativo (DCB) *</Label>
                    <Input
                      placeholder="Ex: Losartana"
                      {...form.register(`medications.${index}.activeIngredient`)}
                    />
                    {form.formState.errors.medications?.[index]?.activeIngredient && (
                      <p className="text-sm text-destructive">
                        {form.formState.errors.medications[index]?.activeIngredient?.message}
                      </p>
                    )}
                  </div>
                </div>

                <div className="grid grid-cols-3 gap-4">
                  <div className="space-y-2">
                    <Label>Categoria *</Label>
                    <Select
                      value={form.watch(`medications.${index}.category`)}
                      onValueChange={(value) =>
                        form.setValue(`medications.${index}.category`, value as any)
                      }
                    >
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="simple">Receita Simples</SelectItem>
                        <SelectItem value="c1">Controle Especial (C1)</SelectItem>
                        <SelectItem value="c5">Psicotrópico (C5)</SelectItem>
                        <SelectItem value="antibiotic">Antimicrobiano</SelectItem>
                        <SelectItem value="glp1">GLP-1 Agonista</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>

                  <div className="space-y-2">
                    <Label>Concentração *</Label>
                    <Input
                      placeholder="Ex: 50mg"
                      {...form.register(`medications.${index}.concentration`)}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Via *</Label>
                    <Select
                      value={form.watch(`medications.${index}.route`)}
                      onValueChange={(value) =>
                        form.setValue(`medications.${index}.route`, value)
                      }
                    >
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="oral">Oral</SelectItem>
                        <SelectItem value="sublingual">Sublingual</SelectItem>
                        <SelectItem value="intravenosa">Intravenosa</SelectItem>
                        <SelectItem value="intramuscular">Intramuscular</SelectItem>
                        <SelectItem value="subcutânea">Subcutânea</SelectItem>
                        <SelectItem value="tópica">Tópica</SelectItem>
                        <SelectItem value="oftálmica">Oftálmica</SelectItem>
                        <SelectItem value="nasal">Nasal</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label>Dosagem *</Label>
                    <Input
                      placeholder="Ex: 1 comprimido"
                      {...form.register(`medications.${index}.dosage`)}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Frequência *</Label>
                    <Input
                      placeholder="Ex: 1x ao dia"
                      {...form.register(`medications.${index}.frequency`)}
                    />
                  </div>
                </div>

                <div className="grid grid-cols-3 gap-4">
                  <div className="space-y-2">
                    <Label>Duração (dias) *</Label>
                    <Input
                      type="number"
                      min="1"
                      placeholder="30"
                      {...form.register(`medications.${index}.duration`)}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Quantidade *</Label>
                    <Input
                      type="number"
                      min="1"
                      placeholder="30"
                      {...form.register(`medications.${index}.quantity`, {
                        onChange: (e) => handleQuantityChange(index, e.target.value),
                      })}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Quantidade por Extenso</Label>
                    <Input
                      {...form.register(`medications.${index}.quantityInWords`)}
                      disabled
                      className="bg-muted"
                    />
                  </div>
                </div>

                <div className="space-y-2">
                  <Label>Instruções Específicas (opcional)</Label>
                  <Textarea
                    placeholder="Orientações específicas para este medicamento..."
                    rows={2}
                    {...form.register(`medications.${index}.instructions`)}
                  />
                </div>
              </CardContent>
            </Card>
          ))}

          {/* Add Medication Button */}
          {fields.length < 10 && (
            <Button
              type="button"
              variant="outline"
              onClick={addMedication}
              className="w-full"
              disabled={fields.length >= 10}
            >
              <Plus className="h-4 w-4 mr-2" />
              Adicionar Medicamento ({fields.length}/10)
            </Button>
          )}
        </div>

        {/* General Instructions */}
        <Card className="mt-6">
          <CardHeader>
            <CardTitle>Instruções Gerais</CardTitle>
            <CardDescription>
              Instruções que se aplicam a todos os medicamentos da prescrição.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Textarea
              placeholder="Orientações gerais para o paciente..."
              rows={3}
              {...form.register('generalInstructions')}
            />
          </CardContent>
        </Card>

        {/* Form validation error */}
        {form.formState.errors.medications && typeof form.formState.errors.medications.message === 'string' && (
          <Alert variant="destructive" className="mt-6">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Erro de Validação</AlertTitle>
            <AlertDescription>
              {form.formState.errors.medications.message}
            </AlertDescription>
          </Alert>
        )}

        {/* Action Buttons */}
        <div className="flex gap-4 mt-6">
          <Button
            type="button"
            variant="outline"
            size="lg"
            onClick={onSaveDraft}
            disabled={saveDraftMutation.isPending || createAndSignMutation.isPending}
            className="flex-1"
          >
            {saveDraftMutation.isPending ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Salvando...
              </>
            ) : (
              <>
                <Save className="mr-2 h-4 w-4" />
                Salvar Rascunho
              </>
            )}
          </Button>

          <Button
            type="submit"
            size="lg"
            disabled={createAndSignMutation.isPending || saveDraftMutation.isPending}
            className="flex-1"
          >
            {createAndSignMutation.isPending ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Criando e Assinando...
              </>
            ) : (
              <>
                <FileCheck className="mr-2 h-4 w-4" />
                Salvar e Assinar
              </>
            )}
          </Button>
        </div>
      </form>
    </div>
  )
}
