'use client'

import { useRef } from 'react'
import { useFieldArray, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { AlertCircle, FileCheck, Loader2, Plus, Save, Trash2 } from 'lucide-react'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Separator } from '@/components/ui/separator'
import { Textarea } from '@/components/ui/textarea'
import { MedicationSearch } from '@/components/prescriptions/MedicationSearch'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { matchAllergies, useAllergies } from '@/lib/api/clinical-skeleton'
import type { MedicationDefinition } from '@/lib/api/medication-definitions'
import { isValidPrescriptionQuantity, numberToWordsWithUnit } from '@/lib/utils/number-to-words'
import {
  CATEGORY_OPTIONS,
  commercialPrescriptionSchema,
  computeQuantity,
  distinctControlledSubstances,
  emptyMedication,
  emptyCommercialPrescription,
  FREQUENCY_OPTIONS,
  ROUTE_FROM_CATALOG,
  ROUTE_OPTIONS,
  unitFromPharmaceuticalForm,
  type CommercialPrescriptionFormData,
} from '@/lib/validations/prescription'

interface CommercialPrescriptionFormProps {
  patientId?: string
  defaultValues?: CommercialPrescriptionFormData
  onSubmit: (data: CommercialPrescriptionFormData) => void
  onSaveDraft: (data: CommercialPrescriptionFormData) => void
  isSigning?: boolean
  isSavingDraft?: boolean
  submitLabel?: string
  draftLabel?: string
  readOnly?: boolean
}

/**
 * Formulário de receita de medicamento industrializado.
 *
 * Estava embutido em `prescriptions/new/page.tsx` e copiado em `[id]/edit/page.tsx` — as duas
 * telas somavam ~1.500 linhas com o mesmo corpo e o mesmo schema. Extraído aqui, as páginas
 * viram casca (guarda de paciente, mutations, tela de sucesso) e a regra de validação passa a
 * ter um dono só.
 */
export function CommercialPrescriptionForm({
  patientId,
  defaultValues,
  onSubmit,
  onSaveDraft,
  isSigning = false,
  isSavingDraft = false,
  submitLabel = 'Salvar e assinar',
  draftLabel = 'Salvar rascunho',
  readOnly = false,
}: CommercialPrescriptionFormProps) {
  const formRef = useRef<HTMLFormElement>(null)

  const form = useForm<CommercialPrescriptionFormData>({
    // schema usa .default() (input≠output); destipa só o resolver p/ alinhar com useForm<output>
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    resolver: zodResolver(commercialPrescriptionSchema as any),
    defaultValues: defaultValues ?? emptyCommercialPrescription(),
  })

  const { fields, append, remove } = useFieldArray({ control: form.control, name: 'medications' })

  // OBRIGATÓRIO no repo: Enter navega entre campos em vez de submeter.
  useFormNavigation({ formRef })

  const medications = form.watch('medications')
  const controlledSubstances = distinctControlledSubstances(medications)
  const hasControlled = medications.some((m) => m.category !== 'simple')

  // CDS não-bloqueante: cruza o princípio ativo (ou o nome) de cada item com as alergias
  // ATIVAS do paciente.
  const { data: patientAllergies = [] } = useAllergies(patientId)
  const allergyConflicts = (() => {
    const seen = new Set<string>()
    const out: { med: string; substance: string }[] = []
    for (const m of medications) {
      const term = m.activeIngredient || m.medicationName || ''
      for (const a of matchAllergies(patientAllergies, term)) {
        const key = `${term}|${a.id}`
        if (seen.has(key)) continue
        seen.add(key)
        out.push({ med: m.medicationName || term, substance: a.substance })
      }
    }
    return out
  })()

  /** Recalcula o "por extenso" a partir da quantidade e da forma farmacêutica do item. */
  const refreshQuantityInWords = (index: number, quantity: number) => {
    if (!Number.isFinite(quantity) || quantity <= 0 || quantity > 9999) return
    const form_ = form.getValues(`medications.${index}.pharmaceuticalForm`)
    form.setValue(
      `medications.${index}.quantityInWords`,
      numberToWordsWithUnit(quantity, unitFromPharmaceuticalForm(form_))
    )
  }

  /**
   * Quantidade derivada de dose × frequência × duração. Só escreve quando dá para afirmar e
   * quando o médico ainda não mexeu no campo à mão — palpite não sobrescreve decisão.
   */
  const recomputeQuantity = (index: number) => {
    if (form.getFieldState(`medications.${index}.quantity`).isDirty) return
    const m = form.getValues(`medications.${index}`)
    const qty = computeQuantity(m.dosage || '', m.frequency || '', Number(m.duration) || 0)
    if (qty === null) return
    form.setValue(`medications.${index}.quantity`, qty)
    refreshQuantityInWords(index, qty)
  }

  /**
   * Receitas antigas (e as importadas por "repetir receita") têm frequência em texto livre, que
   * pode não estar na lista. Sem incluir o valor atual, o Select mostraria o placeholder e daria
   * a impressão de campo vazio com dado gravado por baixo.
   */
  const frequencyOptionsFor = (current?: string) =>
    !current || FREQUENCY_OPTIONS.some((f) => f.label === current)
      ? FREQUENCY_OPTIONS
      : [...FREQUENCY_OPTIONS, { label: current, timesPerDay: null }]

  const handleMedicationSelect = (index: number, medication: MedicationDefinition) => {
    form.setValue(`medications.${index}.medicationDefinitionId`, medication.id)
    form.setValue(`medications.${index}.medicationName`, medication.commonName)
    form.setValue(`medications.${index}.activeIngredient`, medication.activeIngredient)
    form.setValue(`medications.${index}.category`, medication.category as never)

    if (medication.concentration) {
      form.setValue(`medications.${index}.concentration`, medication.concentration)
    }
    if (medication.pharmaceuticalForm) {
      form.setValue(`medications.${index}.pharmaceuticalForm`, medication.pharmaceuticalForm)
    }
    // Vias que a tela não oferece (ex.: "injetavel" genérico) ficam para o médico escolher —
    // melhor campo vazio que via errada numa receita.
    const route = medication.route ? ROUTE_FROM_CATALOG[medication.route] : undefined
    if (route) form.setValue(`medications.${index}.route`, route)

    if (medication.packageQuantity && !form.getFieldState(`medications.${index}.quantity`).isDirty) {
      form.setValue(`medications.${index}.quantity`, medication.packageQuantity)
      refreshQuantityInWords(index, medication.packageQuantity)
    }
  }

  const submitting = isSigning || isSavingDraft

  return (
    <>
      {hasControlled && (
        <Alert className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Medicamentos Controlados</AlertTitle>
          <AlertDescription>
            {controlledSubstances.length > 0 && (
              <p className={controlledSubstances.length > 3 ? 'font-semibold text-destructive' : ''}>
                {controlledSubstances.length} substância(s) de controle especial — o limite do
                receituário é 3.
              </p>
            )}
            <p className="mt-2 text-sm">
              Controle Especial C1: máximo 3 substâncias, 30 dias de validade.
              <br />
              Antimicrobiano: 10 dias de validade (RDC 471/2021).
              <br />
              GLP-1 Agonista: 90 dias de validade.
            </p>
          </AlertDescription>
        </Alert>
      )}

      {allergyConflicts.length > 0 && (
        <Alert variant="destructive" className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Alerta de alergia</AlertTitle>
          <AlertDescription>
            <p className="text-sm">
              O paciente tem alergia registrada que pode conflitar com esta prescrição:
            </p>
            <ul className="mt-1 list-disc pl-5 text-sm">
              {allergyConflicts.map((c, i) => (
                <li key={i}>
                  <strong>{c.med}</strong> — alergia a <strong>{c.substance}</strong>
                </li>
              ))}
            </ul>
            <p className="mt-2 text-xs">
              Revise antes de assinar. Este alerta não bloqueia a prescrição.
            </p>
          </AlertDescription>
        </Alert>
      )}

      <form ref={formRef} onSubmit={form.handleSubmit(onSubmit)}>
        {/* min-w-0: fieldset tem `min-inline-size: min-content` por default do navegador e,
            diferente de uma div, não encolhe abaixo do conteúdo — no iPhone isso empurrava a
            página 5px além da viewport e criava rolagem horizontal. */}
        <fieldset disabled={readOnly} className="min-w-0 space-y-6">
          {fields.map((field, index) => (
            <Card key={field.id}>
              <CardHeader>
                <div className="flex items-center justify-between">
                  <CardTitle>Medicamento {index + 1}</CardTitle>
                  {fields.length > 1 && !readOnly && (
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      onClick={() => remove(index)}
                      className="text-destructive hover:text-destructive"
                    >
                      <Trash2 className="mr-2 h-4 w-4" />
                      Remover
                    </Button>
                  )}
                </div>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <Label>Medicamento *</Label>
                  <MedicationSearch
                    valueLabel={form.watch(`medications.${index}.medicationName`)}
                    onSelect={(med) => handleMedicationSelect(index, med)}
                    placeholder="Buscar medicamento cadastrado..."
                  />
                  <p className="text-xs text-muted-foreground">
                    Busque por nome comercial ou princípio ativo. Se não encontrar, preencha
                    manualmente abaixo.
                  </p>
                </div>

                <Separator />

                <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                  <div className="space-y-2">
                    <Label>Nome Comercial *</Label>
                    <Input
                      placeholder="Ex: Losartana Potássica 50mg"
                      {...form.register(`medications.${index}.medicationName`)}
                    />
                    {form.formState.errors.medications?.[index]?.medicationName && (
                      <p className="text-sm text-destructive">
                        {form.formState.errors.medications[index]?.medicationName?.message}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label>Princípio Ativo (DCB)</Label>
                    <Input
                      placeholder="Ex: Losartana"
                      {...form.register(`medications.${index}.activeIngredient`)}
                    />
                  </div>
                </div>

                <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
                  <div className="space-y-2">
                    <Label>Categoria *</Label>
                    <Select
                      value={form.watch(`medications.${index}.category`)}
                      onValueChange={(value) =>
                        form.setValue(`medications.${index}.category`, value as never)
                      }
                    >
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        {CATEGORY_OPTIONS.map((c) => (
                          <SelectItem key={c.value} value={c.value}>
                            {c.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>

                  <div className="space-y-2">
                    <Label>Concentração</Label>
                    <Input
                      placeholder="Ex: 50mg"
                      {...form.register(`medications.${index}.concentration`)}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Via *</Label>
                    <Select
                      value={form.watch(`medications.${index}.route`)}
                      onValueChange={(value) => form.setValue(`medications.${index}.route`, value)}
                    >
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        {ROUTE_OPTIONS.map((r) => (
                          <SelectItem key={r.value} value={r.value}>
                            {r.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>
                </div>

                <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
                  <div className="space-y-2">
                    <Label>Dosagem *</Label>
                    <Input
                      placeholder="Ex: 1 comprimido"
                      {...form.register(`medications.${index}.dosage`, {
                        onBlur: () => recomputeQuantity(index),
                      })}
                    />
                  </div>

                  <div className="space-y-2">
                    <Label>Frequência *</Label>
                    {/* Select em vez de texto livre: as opções carregam quantas vezes ao dia,
                        que é o que permite calcular a quantidade sem inventar campo novo. */}
                    <Select
                      value={form.watch(`medications.${index}.frequency`)}
                      onValueChange={(value) => {
                        form.setValue(`medications.${index}.frequency`, value)
                        recomputeQuantity(index)
                      }}
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Escolha" />
                      </SelectTrigger>
                      <SelectContent>
                        {frequencyOptionsFor(
                          form.watch(`medications.${index}.frequency`)
                        ).map((f) => (
                          <SelectItem key={f.label} value={f.label}>
                            {f.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>

                  <div className="space-y-2">
                    <Label>Duração (dias)</Label>
                    <Input
                      type="number"
                      min="0"
                      placeholder="em branco = uso contínuo"
                      {...form.register(`medications.${index}.duration`, {
                        onBlur: () => recomputeQuantity(index),
                      })}
                    />
                    {/* Deixar em branco não é esquecimento: é a forma de dizer que não há prazo,
                        e a receita sai escrito "uso contínuo". */}
                    {!Number(form.watch(`medications.${index}.duration`)) && (
                      <p className="text-xs text-muted-foreground">
                        Sem prazo: a receita sai como uso contínuo.
                      </p>
                    )}
                  </div>
                </div>

                <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                  <div className="space-y-2">
                    <Label>Quantidade</Label>
                    <Input
                      type="number"
                      min="0"
                      placeholder="em branco = não especificar"
                      {...form.register(`medications.${index}.quantity`, {
                        onChange: (e) => refreshQuantityInWords(index, Number(e.target.value)),
                      })}
                    />
                    <p className="text-xs text-muted-foreground">
                      {form.watch(`medications.${index}.quantityInWords`) ||
                        'Calculada de dose × frequência × duração; em branco, não sai na receita.'}
                    </p>
                    {/* Aviso, não bloqueio: quem decide a quantidade é o médico. A função já
                        existia no repo e ninguém chamava. */}
                    {Number(form.watch(`medications.${index}.quantity`)) > 0 &&
                      !isValidPrescriptionQuantity(
                        Number(form.watch(`medications.${index}.quantity`)),
                        form.watch(`medications.${index}.category`)
                      ) && (
                        <p className="text-xs text-amber-600">
                          Quantidade acima do usual para esta categoria. Confira antes de assinar.
                        </p>
                      )}
                  </div>

                  <div className="space-y-2">
                    <Label>Instruções Específicas (opcional)</Label>
                    <Textarea
                      placeholder="Orientações específicas para este medicamento..."
                      rows={2}
                      {...form.register(`medications.${index}.instructions`)}
                    />
                  </div>
                </div>
              </CardContent>
            </Card>
          ))}

          {fields.length < 10 && !readOnly && (
            <Button
              type="button"
              variant="outline"
              onClick={() => append(emptyMedication())}
              className="w-full"
            >
              <Plus className="mr-2 h-4 w-4" />
              Adicionar Medicamento ({fields.length}/10)
            </Button>
          )}

          <Card>
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

          {form.formState.errors.medications &&
            typeof form.formState.errors.medications.message === 'string' && (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertTitle>Erro de Validação</AlertTitle>
                <AlertDescription>{form.formState.errors.medications.message}</AlertDescription>
              </Alert>
            )}

          {!readOnly && (
            <div className="flex flex-col gap-3 sm:flex-row sm:gap-4">
              <Button
                type="button"
                variant="outline"
                size="lg"
                // Passa pelo handleSubmit para o rascunho também ser validado: antes chamava
                // getValues() direto e mandava payload inválido, que voltava como erro cru do Go.
                onClick={form.handleSubmit(onSaveDraft)}
                disabled={submitting}
                className="flex-1"
              >
                {isSavingDraft ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Salvando...
                  </>
                ) : (
                  <>
                    <Save className="mr-2 h-4 w-4" />
                    {draftLabel}
                  </>
                )}
              </Button>

              <Button type="submit" size="lg" disabled={submitting} className="flex-1">
                {isSigning ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Salvando e assinando...
                  </>
                ) : (
                  <>
                    <FileCheck className="mr-2 h-4 w-4" />
                    {submitLabel}
                  </>
                )}
              </Button>
            </div>
          )}
        </fieldset>
      </form>
    </>
  )
}
