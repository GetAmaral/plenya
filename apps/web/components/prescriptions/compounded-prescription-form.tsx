'use client'

import { useRef, useState } from 'react'
import { useFieldArray, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { AlertCircle, BookMarked, FileCheck, Loader2, Plus, Save } from 'lucide-react'
import { useMutation } from '@tanstack/react-query'
import { toast } from 'sonner'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Textarea } from '@/components/ui/textarea'
import { FormulaCard } from '@/components/prescriptions/formula-card'
import { FormulaTemplateSheet } from '@/components/prescriptions/formula-template-sheet'
import { createFormulaTemplate, type FormulaTemplate } from '@/lib/api/magistral-templates'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { matchAllergies, useAllergies } from '@/lib/api/clinical-skeleton'
import {
  compoundedPrescriptionSchema,
  distinctControlledInFormulas,
  emptyCompoundedPrescription,
  emptyFormula,
  formulaHasControlled,
  type CompoundedPrescriptionFormData,
  type FormulaFormData,
} from '@/lib/validations/compounded-prescription'

interface CompoundedPrescriptionFormProps {
  patientId?: string
  defaultValues?: CompoundedPrescriptionFormData
  onSubmit: (data: CompoundedPrescriptionFormData) => void
  onSaveDraft: (data: CompoundedPrescriptionFormData) => void
  isSigning?: boolean
  isSavingDraft?: boolean
  submitLabel?: string
  draftLabel?: string
}

/**
 * Receita de manipulado: N fórmulas, cada uma com M substâncias.
 *
 * Compartilha a casca com a receita comercial (guarda de paciente, mutations, tela de sucesso)
 * e nada mais — os campos são outros, e esticar o formulário comercial para caber os dois
 * produziria um componente com dois modos e nenhum legível.
 */
export function CompoundedPrescriptionForm({
  patientId,
  defaultValues,
  onSubmit,
  onSaveDraft,
  isSigning = false,
  isSavingDraft = false,
  submitLabel = 'Salvar e assinar',
  draftLabel = 'Salvar rascunho',
}: CompoundedPrescriptionFormProps) {
  const formRef = useRef<HTMLFormElement>(null)
  const [templateSheet, setTemplateSheet] = useState(false)

  const form = useForm<CompoundedPrescriptionFormData>({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    resolver: zodResolver(compoundedPrescriptionSchema as any),
    defaultValues: defaultValues ?? emptyCompoundedPrescription(),
  })

  const { fields, append, remove, replace } = useFieldArray({ control: form.control, name: 'formulas' })

  // OBRIGATÓRIO no repo: Enter navega entre campos em vez de submeter.
  useFormNavigation({ formRef })

  const formulas = form.watch('formulas')
  const controlledSubstances = distinctControlledInFormulas(formulas ?? [])
  const hasControlled = (formulas ?? []).some((f) => formulaHasControlled(f))

  // CDS não-bloqueante: cruza cada componente com as alergias ATIVAS do paciente.
  const { data: patientAllergies = [] } = useAllergies(patientId)
  const allergyConflicts = (() => {
    const seen = new Set<string>()
    const out: { formula: string; substance: string }[] = []
    ;(formulas ?? []).forEach((f, i) => {
      for (const c of f.components ?? []) {
        for (const a of matchAllergies(patientAllergies, c.substance || '')) {
          const key = `${c.substance}|${a.id}`
          if (seen.has(key)) continue
          seen.add(key)
          out.push({ formula: f.name || `Fórmula ${i + 1}`, substance: a.substance })
        }
      }
    })
    return out
  })()

  /**
   * Aplica a fórmula-base no primeiro cartão vazio (ou acrescenta um). As doses vêm da tela do
   * painel, onde o médico já viu o que a regra sugeriu e decidiu; o que a regra sugeriu fica
   * gravado ao lado, para revelar depois quando ele discordou.
   */
  const applyTemplate = (
    template: FormulaTemplate,
    doses: Record<string, number>,
    suggested: Record<string, number>
  ) => {
    const nova: FormulaFormData = {
      templateId: template.id,
      name: template.name,
      pharmaceuticalForm: template.pharmaceuticalForm,
      usageType: template.usageType,
      route: template.route ?? '',
      vehicle: template.vehicle ?? '',
      quantityToDispense: template.quantityToDispense,
      quantityUnit: template.quantityUnit,
      quantityInWords: '',
      posology: template.posology ?? '',
      duration: template.duration ?? 0,
      instructions: template.instructions ?? '',
      components: (template.components ?? []).map((c) => ({
        magistralComponentId: c.magistralComponentId,
        substance: c.substance,
        quantity: doses[c.substance] ?? c.quantity,
        unit: c.unit,
        category: (c.category ?? 'simple') as FormulaFormData['components'][number]['category'],
        // A base já sabe que a dose do mineral é do elemento; sem carregar a marca, a receita
        // saía dizendo que a quantidade era do insumo.
        asElemental: c.asElemental,
        note: c.note ?? '',
        suggestedQuantity: suggested[c.substance],
      })),
    }

    const atuais = form.getValues('formulas') ?? []
    const vazia = atuais.findIndex(
      (f) => (f.components ?? []).every((c) => !c.substance?.trim()) && !f.name?.trim()
    )
    if (vazia >= 0) {
      // `setValue('formulas.0', ...)` não avisa o useFieldArray aninhado do card, que continua
      // com as linhas antigas: a base de 4 substâncias mostrava 1 linha e mandava as 4 no submit.
      // Trocar o array inteiro com `replace` re-renderiza os cards com o que será enviado.
      const proximas = [...atuais]
      proximas[vazia] = nova
      replace(proximas)
    } else {
      append(nova)
    }
    toast.success(`${template.name} aplicada`)
  }

  const saveAsTemplate = useMutation({
    mutationFn: async (index: number) => {
      const f = form.getValues(`formulas.${index}`)
      if (!f?.components?.length || !f.components.some((c) => c.substance?.trim())) {
        throw new Error('Preencha os componentes antes de salvar como base')
      }
      return createFormulaTemplate({
        name: f.name?.trim() || `Fórmula ${index + 1}`,
        pharmaceuticalForm: f.pharmaceuticalForm,
        usageType: f.usageType,
        route: f.route,
        vehicle: f.vehicle,
        quantityToDispense: f.quantityToDispense,
        quantityUnit: f.quantityUnit,
        posology: f.posology,
        duration: f.duration || undefined,
        instructions: f.instructions || undefined,
        components: f.components
          .filter((c) => c.substance?.trim())
          .map((c) => ({
            magistralComponentId: c.magistralComponentId,
            substance: c.substance,
            quantity: c.quantity,
            unit: c.unit,
            category: c.category,
            note: c.note,
          })),
      })
    },
    onSuccess: (t) =>
      toast.success(`"${t.name}" salva como fórmula-base`, {
        description: 'As regras de dose se cadastram na tela de fórmulas-base.',
      }),
    onError: (e: any) => toast.error('Não deu para salvar', { description: e?.message }),
  })

  const submitting = isSigning || isSavingDraft

  return (
    <>
      {hasControlled && (
        <Alert className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Fórmula com substância controlada</AlertTitle>
          <AlertDescription>
            <p className={controlledSubstances.length > 3 ? 'font-semibold text-destructive' : ''}>
              {controlledSubstances.length} substância(s) de controle especial — o limite do
              receituário é 3.
            </p>
            <p className="mt-2 text-sm">
              A receita sai como receituário de controle especial, para impressão e assinatura à
              mão. A 344/98 pede receituário exclusivo para o controlado: se a farmácia exigir,
              emita a fórmula controlada numa receita separada.
            </p>
          </AlertDescription>
        </Alert>
      )}

      {allergyConflicts.length > 0 && (
        <Alert variant="destructive" className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Alerta de alergia</AlertTitle>
          <AlertDescription>
            <ul className="mt-1 list-disc pl-5 text-sm">
              {allergyConflicts.map((c, i) => (
                <li key={i}>
                  <strong>{c.formula}</strong> — alergia a <strong>{c.substance}</strong>
                </li>
              ))}
            </ul>
            <p className="mt-2 text-xs">
              Revise antes de assinar. Este alerta não bloqueia a prescrição.
            </p>
          </AlertDescription>
        </Alert>
      )}

      {/* A fórmula-base é o caminho normal de quem já tem repertório: escolhe a base, confere o
          que as regras sugerem para este paciente, e a fórmula nasce montada. */}
      <div className="mb-4 flex flex-wrap gap-2">
        <Button type="button" variant="outline" onClick={() => setTemplateSheet(true)}>
          <BookMarked className="mr-2 h-4 w-4" />
          Usar fórmula-base
        </Button>
      </div>

      <FormulaTemplateSheet
        open={templateSheet}
        onOpenChange={setTemplateSheet}
        patientId={patientId}
        onApply={applyTemplate}
      />

      <form ref={formRef} onSubmit={form.handleSubmit(onSubmit)}>
        <fieldset className="min-w-0 space-y-6">
          {fields.map((field, index) => (
            <FormulaCard
              key={field.id}
              form={form}
              control={form.control}
              index={index}
              canRemove={fields.length > 1}
              onRemove={() => remove(index)}
              onSaveAsTemplate={() => saveAsTemplate.mutate(index)}
              savingTemplate={saveAsTemplate.isPending}
            />
          ))}

          {fields.length < 10 && (
            <Button
              type="button"
              variant="outline"
              className="w-full"
              onClick={() => append(emptyFormula())}
            >
              <Plus className="mr-2 h-4 w-4" />
              Adicionar fórmula ({fields.length}/10)
            </Button>
          )}

          <Card>
            <CardHeader>
              <CardTitle>Instruções gerais</CardTitle>
              <CardDescription>
                Instruções que se aplicam a todas as fórmulas da receita.
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

          {form.formState.errors.formulas &&
            typeof form.formState.errors.formulas.message === 'string' && (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertTitle>Erro de validação</AlertTitle>
                <AlertDescription>{form.formState.errors.formulas.message}</AlertDescription>
              </Alert>
            )}

          <div className="flex flex-col gap-3 sm:flex-row sm:gap-4">
            <Button
              type="button"
              variant="outline"
              size="lg"
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
        </fieldset>
      </form>
    </>
  )
}
