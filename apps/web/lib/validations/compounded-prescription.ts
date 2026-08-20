import { z } from 'zod'

import type { CreatePrescriptionRequest, Prescription } from '@/lib/api/prescriptions'
import { medicationCategoryEnum, type MedicationCategory } from '@/lib/validations/prescription'

/**
 * Receita de manipulado (fórmula magistral).
 *
 * Estrutura diferente da receita de industrializado: N fórmulas, cada uma com M substâncias,
 * forma farmacêutica, veículo, quantidade a aviar e posologia próprias. Por isso schema próprio
 * em vez de esticar o comercial — as duas telas compartilham a casca, não os campos.
 */

export const formulaComponentSchema = z.object({
  magistralComponentId: z.string().optional(),
  /** O que a regra da fórmula-base sugeriu, quando houve sugestão aceita. Vai ao lado da dose. */
  suggestedQuantity: z.coerce.number().optional(),
  /** A dose escrita é do elemento (magnésio elementar), não do insumo (bisglicinato). */
  asElemental: z.boolean().optional(),
  medicationDefinitionId: z.string().optional(),
  substance: z.string().min(2, 'Informe a substância'),
  quantity: z.coerce.number().gt(0, 'Quantidade deve ser maior que zero'),
  unit: z.string().min(1, 'Informe a unidade'),
  category: medicationCategoryEnum,
  note: z.string().optional(),
})

export type FormulaComponentFormData = z.infer<typeof formulaComponentSchema>

export const formulaSchema = z.object({
  templateId: z.string().optional(),
  name: z.string().optional(),
  pharmaceuticalForm: z.string().min(1, 'Escolha a forma farmacêutica'),
  usageType: z.enum(['internal', 'external']),
  route: z.string().optional(),
  vehicle: z.string().optional(),
  quantityToDispense: z.coerce.number().gt(0, 'Informe quanto aviar'),
  quantityUnit: z.string().min(1, 'Informe a unidade do aviamento'),
  quantityInWords: z.string().optional(),
  posology: z.string().min(1, 'Informe a posologia'),
  duration: z.coerce.number().min(0).max(365).optional(),
  instructions: z.string().optional(),
  components: z
    .array(formulaComponentSchema)
    .min(1, 'A fórmula precisa de pelo menos um componente')
    // 20 é limite de LAYOUT do receituário, não clínico: o paginador do PDF não quebra o bloco
    // de uma fórmula, então acima disso a receita sairia cortada. O backend recusa igual.
    .max(20, 'Máximo 20 componentes por fórmula'),
})

export type FormulaFormData = z.infer<typeof formulaSchema>

export const compoundedPrescriptionSchema = z
  .object({
    formulas: z
      .array(formulaSchema)
      .min(1, 'Adicione pelo menos uma fórmula')
      .max(10, 'Máximo 10 fórmulas'),
    generalInstructions: z.string().optional(),
  })
  .refine((data) => distinctControlledInFormulas(data.formulas).length <= 3, {
    message: 'Receita de controle especial aceita no máximo 3 substâncias diferentes',
    path: ['formulas'],
  })
  .refine(
    (data) =>
      data.formulas.every(
        (f) => !formulaHasControlled(f) || (f.quantityInWords || '').trim().length > 0
      ),
    {
      message: 'Fórmula com substância controlada exige a quantidade a aviar por extenso',
      path: ['formulas'],
    }
  )

export type CompoundedPrescriptionFormData = z.infer<typeof compoundedPrescriptionSchema>

const CONTROLLED: MedicationCategory[] = ['c1', 'c5']

export function formulaHasControlled(formula: Pick<FormulaFormData, 'components'>): boolean {
  return formula.components.some((c) => CONTROLLED.includes(c.category))
}

/** Limite da 344/98 é por SUBSTÂNCIA e vale para a receita inteira, não por fórmula. */
export function distinctControlledInFormulas(formulas: FormulaFormData[]): string[] {
  const seen = new Set<string>()
  for (const f of formulas) {
    for (const c of f.components) {
      if (c.category !== 'c1') continue
      const key = (c.substance || '')
        .normalize('NFD')
        .replace(/[̀-ͯ]/g, '')
        .toLowerCase()
        .trim()
      if (key) seen.add(key)
    }
  }
  return [...seen]
}

/** Categoria mais restritiva da fórmula — é ela que decide o rótulo e o modo de assinatura. */
const CATEGORY_RANK: Record<MedicationCategory, number> = {
  simple: 0,
  glp1: 1,
  antibiotic: 2,
  c5: 3,
  c1: 4,
}

export function highestCategory(formula: Pick<FormulaFormData, 'components'>): MedicationCategory {
  let highest: MedicationCategory = 'simple'
  for (const c of formula.components) {
    if (CATEGORY_RANK[c.category] > CATEGORY_RANK[highest]) highest = c.category
  }
  return highest
}

/**
 * Formas farmacêuticas com os defaults que evitam digitar o óbvio: escolher "cápsula" já
 * preenche unidade de aviamento, uso, via e veículo. Tudo continua editável.
 */
export interface PharmaceuticalFormOption {
  value: string
  label: string
  quantityUnit: string
  usageType: 'internal' | 'external'
  route: string
  vehicle: string
  /** Unidade sugerida para os componentes desta forma. */
  componentUnit: string
}

export const PHARMACEUTICAL_FORMS: PharmaceuticalFormOption[] = [
  {
    value: 'cápsula',
    label: 'Cápsula',
    quantityUnit: 'cápsulas',
    usageType: 'internal',
    route: 'oral',
    vehicle: 'Excipiente qsp 1 cápsula',
    componentUnit: 'mg',
  },
  {
    value: 'sachê',
    label: 'Sachê',
    quantityUnit: 'sachês',
    usageType: 'internal',
    route: 'oral',
    vehicle: 'Veículo qsp 1 sachê',
    componentUnit: 'mg',
  },
  {
    value: 'solução oral',
    label: 'Solução oral',
    quantityUnit: 'mL',
    usageType: 'internal',
    route: 'oral',
    vehicle: 'Veículo qsp 100 mL',
    componentUnit: 'mg',
  },
  {
    value: 'creme',
    label: 'Creme',
    quantityUnit: 'g',
    usageType: 'external',
    route: 'tópica',
    vehicle: 'Creme não iônico qsp 30 g',
    componentUnit: '%',
  },
  {
    value: 'gel',
    label: 'Gel',
    quantityUnit: 'g',
    usageType: 'external',
    route: 'tópica',
    vehicle: 'Gel de natrosol qsp 30 g',
    componentUnit: '%',
  },
  {
    value: 'pomada',
    label: 'Pomada',
    quantityUnit: 'g',
    usageType: 'external',
    route: 'tópica',
    vehicle: 'Vaselina sólida qsp 30 g',
    componentUnit: '%',
  },
  {
    value: 'loção',
    label: 'Loção',
    quantityUnit: 'mL',
    usageType: 'external',
    route: 'tópica',
    vehicle: 'Loção cremosa qsp 100 mL',
    componentUnit: '%',
  },
  {
    // Transdérmico é via, não uso externo: o ativo atravessa a pele e cai na circulação. Por isso
    // a dose vem em mg por dose aplicada, e não em porcentagem como nos tópicos de ação local.
    value: 'transdérmico',
    label: 'Transdérmico (Pentravan)',
    quantityUnit: 'mL',
    usageType: 'external',
    route: 'transdérmica',
    vehicle: 'Pentravan qsp 1 mL',
    componentUnit: 'mg',
  },
  {
    value: 'vaginal',
    label: 'Vaginal (Pentravan)',
    quantityUnit: 'g',
    usageType: 'external',
    route: 'vaginal',
    vehicle: 'Pentravan qsp 1 g',
    componentUnit: 'mg',
  },
  {
    value: 'xampu',
    label: 'Xampu',
    quantityUnit: 'mL',
    usageType: 'external',
    route: 'tópica',
    vehicle: 'Base de xampu qsp 200 mL',
    componentUnit: '%',
  },
]

export const COMPONENT_UNITS = ['mg', 'g', 'mcg', 'UI', 'mL', '%'] as const

export const USAGE_OPTIONS = [
  { value: 'internal', label: 'Uso interno' },
  { value: 'external', label: 'Uso externo' },
] as const

export function emptyFormulaComponent(unit = 'mg'): FormulaComponentFormData {
  return { substance: '', quantity: 0, unit, category: 'simple', note: '' }
}

export function emptyFormula(): FormulaFormData {
  const capsule = PHARMACEUTICAL_FORMS[0]
  return {
    name: '',
    pharmaceuticalForm: capsule.value,
    usageType: capsule.usageType,
    route: capsule.route,
    vehicle: capsule.vehicle,
    quantityToDispense: 60,
    quantityUnit: capsule.quantityUnit,
    quantityInWords: '',
    posology: '',
    duration: 30,
    instructions: '',
    components: [emptyFormulaComponent(capsule.componentUnit)],
  }
}

export function emptyCompoundedPrescription(): CompoundedPrescriptionFormData {
  return { formulas: [emptyFormula()], generalInstructions: '' }
}

type FormulaPayload = NonNullable<CreatePrescriptionRequest['formulas']>[number]

export function toCompoundedPayload(
  data: CompoundedPrescriptionFormData,
  patientId: string
): CreatePrescriptionRequest {
  return {
    patientId,
    type: 'compounded',
    formulas: data.formulas.map((f) => ({
      templateId: f.templateId || undefined,
      name: f.name || undefined,
      pharmaceuticalForm: f.pharmaceuticalForm,
      usageType: f.usageType as FormulaPayload['usageType'],
      route: f.route || undefined,
      vehicle: f.vehicle || undefined,
      quantityToDispense: f.quantityToDispense,
      quantityUnit: f.quantityUnit,
      quantityInWords: f.quantityInWords || undefined,
      posology: f.posology,
      duration: f.duration || undefined,
      instructions: f.instructions || undefined,
      components: f.components.map((c) => ({
        magistralComponentId: c.magistralComponentId || undefined,
        suggestedQuantity: c.suggestedQuantity,
        asElemental: c.asElemental,
        medicationDefinitionId: c.medicationDefinitionId || undefined,
        substance: c.substance,
        quantity: c.quantity,
        unit: c.unit,
        category: c.category as FormulaPayload['components'][number]['category'],
        note: c.note || undefined,
      })),
    })),
    generalInstructions: data.generalInstructions || undefined,
    prescriptionDate: new Date().toISOString(),
  }
}

/** Prescrição da API → valores do formulário (edição e "repetir receita"). */
export function fromCompoundedPrescription(
  prescription: Prescription
): CompoundedPrescriptionFormData {
  return {
    formulas: (prescription.formulas ?? []).map((f) => ({
      templateId: f.templateId,
      name: f.name ?? '',
      pharmaceuticalForm: f.pharmaceuticalForm,
      usageType: (f.usageType ?? 'internal') as FormulaFormData['usageType'],
      route: f.route ?? '',
      vehicle: f.vehicle ?? '',
      quantityToDispense: f.quantityToDispense,
      quantityUnit: f.quantityUnit,
      quantityInWords: f.quantityInWords ?? '',
      posology: f.posology,
      duration: f.duration ?? 0,
      instructions: f.instructions ?? '',
      components: (f.components ?? []).map((c) => ({
        magistralComponentId: c.magistralComponentId,
        suggestedQuantity: c.suggestedQuantity,
        asElemental: c.asElemental,
        medicationDefinitionId: c.medicationDefinitionId,
        substance: c.substance,
        quantity: c.quantity,
        unit: c.unit,
        category: (c.category ?? 'simple') as MedicationCategory,
        note: c.note ?? '',
      })),
    })),
    generalInstructions: prescription.generalInstructions ?? '',
  }
}
