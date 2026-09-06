import { z } from 'zod'

import type { CreatePrescriptionRequest, Prescription } from '@/lib/api/prescriptions'

/**
 * Schemas da prescrição — fonte única.
 *
 * Este arquivo existe porque o schema estava LITERALMENTE duplicado entre
 * `prescriptions/new/page.tsx` e `prescriptions/[id]/edit/page.tsx`, e a receita de
 * manipulados ia criar a terceira e a quarta cópia. Qualquer regra de validação da receita
 * mora aqui; as telas só consomem.
 */

export const medicationCategoryEnum = z.enum(['simple', 'c1', 'c5', 'antibiotic', 'glp1'])
export type MedicationCategory = z.infer<typeof medicationCategoryEnum>

export const medicationSchema = z.object({
  medicationDefinitionId: z.string().optional(),
  category: medicationCategoryEnum,
  medicationName: z.string().min(3, 'Nome do medicamento é obrigatório'),
  // Princípio ativo e concentração deixaram de ser obrigatórios: quando o médico escreve
  // "Losartana 50mg" no nome, exigir os dois de novo é digitar a mesma coisa três vezes. O
  // backend cai para o nome quando o princípio ativo vem vazio.
  activeIngredient: z.string().optional(),
  concentration: z.string().optional(),
  dosage: z.string().min(1, 'Dosagem é obrigatória'),
  frequency: z.string().min(1, 'Frequência é obrigatória'),
  route: z.string().min(1, 'Via é obrigatória'),
  // OPCIONAIS: em branco significa USO CONTÍNUO, que é a maior parte da prescrição de longo prazo.
  // Campo vazio vira 0 (o backend trata 0 como "sem prazo"); preenchido, o teto espelha o do
  // servidor (max=365), senão vem 400 sem erro de campo na tela.
  duration: z.preprocess(
    (v) => (v === '' || v === null || v === undefined ? 0 : v),
    z.coerce
      .number()
      .min(0)
      .max(365, 'Duração máxima de 365 dias')
  ),
  quantity: z.preprocess(
    (v) => (v === '' || v === null || v === undefined ? 0 : v),
    z.coerce.number().min(0)
  ),
  // Derivado da quantidade; some da tela e é recalculado a cada mudança.
  quantityInWords: z.string().optional(),
  instructions: z.string().optional(),
  // Só do cliente: guarda a forma farmacêutica que veio do catálogo para o "por extenso"
  // saber se são comprimidos, gotas ou frascos. Não vai no payload da prescrição.
  pharmaceuticalForm: z.string().optional(),
})

export type MedicationFormData = z.infer<typeof medicationSchema>

export const commercialPrescriptionSchema = z
  .object({
    medications: z
      .array(medicationSchema)
      .min(1, 'Adicione pelo menos um medicamento')
      .max(10, 'Máximo 10 medicamentos'),
    generalInstructions: z.string().optional(),
  })
  .refine((data) => distinctControlledSubstances(data.medications).length <= 3, {
    message: 'Receita de controle especial aceita no máximo 3 substâncias diferentes',
    path: ['medications'],
  })

export type CommercialPrescriptionFormData = z.infer<typeof commercialPrescriptionSchema>

/**
 * A Portaria 344/98 limita a 3 SUBSTÂNCIAS por receituário de controle especial — não a 3
 * itens. Duas apresentações da mesma substância contam uma vez; a contagem antiga por item
 * recusava receita legítima e aceitava receita irregular com nomes comerciais diferentes.
 */
export function distinctControlledSubstances(
  medications: Pick<MedicationFormData, 'category' | 'activeIngredient' | 'medicationName'>[]
): string[] {
  const seen = new Set<string>()
  for (const m of medications) {
    if (m.category !== 'c1') continue
    const key = normalizeSubstance(m.activeIngredient || m.medicationName || '')
    if (key) seen.add(key)
  }
  return [...seen]
}

function normalizeSubstance(value: string): string {
  return value
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .toLowerCase()
    .trim()
}

/** Vias oferecidas na tela. O catálogo grava sem acento; ver ROUTE_FROM_CATALOG. */
export const ROUTE_OPTIONS = [
  { value: 'oral', label: 'Oral' },
  { value: 'sublingual', label: 'Sublingual' },
  { value: 'intravenosa', label: 'Intravenosa' },
  { value: 'intramuscular', label: 'Intramuscular' },
  { value: 'subcutânea', label: 'Subcutânea' },
  { value: 'tópica', label: 'Tópica' },
  { value: 'oftálmica', label: 'Oftálmica' },
  { value: 'nasal', label: 'Nasal' },
] as const

/**
 * O catálogo da ANVISA grava a via sem acento (derivada da forma farmacêutica); o select usa
 * os rótulos acentuados. Sem esta tradução o preenchimento automático falharia calado
 * justamente em tópica, subcutânea e oftálmica.
 */
export const ROUTE_FROM_CATALOG: Record<string, string> = {
  oral: 'oral',
  sublingual: 'sublingual',
  intravenosa: 'intravenosa',
  intramuscular: 'intramuscular',
  subcutanea: 'subcutânea',
  topica: 'tópica',
  oftalmica: 'oftálmica',
  nasal: 'nasal',
}

/** Rótulos curtos, para os selects estreitos (linha de componente da fórmula). */
export const CATEGORY_OPTIONS_SHORT = [
  { value: 'simple', label: 'Simples' },
  { value: 'c1', label: 'C1' },
  { value: 'c5', label: 'C5' },
  { value: 'antibiotic', label: 'Antimicrob.' },
  { value: 'glp1', label: 'GLP-1' },
] as const

export const CATEGORY_OPTIONS = [
  { value: 'simple', label: 'Receita Simples' },
  { value: 'c1', label: 'Controle Especial (C1)' },
  { value: 'c5', label: 'Psicotrópico (C5)' },
  { value: 'antibiotic', label: 'Antimicrobiano' },
  { value: 'glp1', label: 'GLP-1 Agonista' },
] as const

/**
 * Frequências com o multiplicador embutido. A coluna do banco continua guardando o TEXTO
 * (nada muda no schema nem no PDF); o número só vive no cliente, para calcular a quantidade.
 * `timesPerDay: null` = não dá para calcular ("se necessário").
 */
export const FREQUENCY_OPTIONS: { label: string; timesPerDay: number | null }[] = [
  { label: '1x ao dia', timesPerDay: 1 },
  { label: '2x ao dia', timesPerDay: 2 },
  { label: '3x ao dia', timesPerDay: 3 },
  { label: 'de 12/12h', timesPerDay: 2 },
  { label: 'de 8/8h', timesPerDay: 3 },
  { label: 'de 6/6h', timesPerDay: 4 },
  { label: 'ao deitar', timesPerDay: 1 },
  { label: 'em jejum', timesPerDay: 1 },
  { label: 'pela manhã', timesPerDay: 1 },
  { label: '1x por semana', timesPerDay: 1 / 7 },
  { label: 'se necessário', timesPerDay: null },
]

/**
 * Quantidade a dispensar = dose por tomada × tomadas/dia × dias. Devolve null quando não dá
 * para afirmar (frequência sem multiplicador, dosagem sem número) — nesse caso é melhor não
 * mexer no que o médico digitou do que sobrescrever com um palpite.
 */
export function computeQuantity(
  dosage: string,
  frequencyLabel: string,
  durationDays: number
): number | null {
  const perDay = FREQUENCY_OPTIONS.find((f) => f.label === frequencyLabel)?.timesPerDay
  if (!perDay || !durationDays || durationDays < 1) return null

  const dose = parseLeadingNumber(dosage)
  if (dose === null) return null

  const total = Math.ceil(dose * perDay * durationDays)
  return total > 0 && total <= 9999 ? total : null
}

/** "1 comprimido" → 1 · "1/2 comprimido" → 0.5 · "conforme orientação" → null */
export function parseLeadingNumber(value: string): number | null {
  const text = (value || '').trim()
  const fraction = text.match(/^(\d+)\s*\/\s*(\d+)/)
  if (fraction) {
    const denominator = Number(fraction[2])
    return denominator ? Number(fraction[1]) / denominator : null
  }
  const match = text.match(/^(\d+(?:[.,]\d+)?)/)
  if (!match) return null
  const n = Number(match[1].replace(',', '.'))
  return Number.isFinite(n) && n > 0 ? n : null
}

/**
 * Unidade do "por extenso", vinda da forma farmacêutica do catálogo. Antes era a string fixa
 * "comprimidos", que escrevia "trinta comprimidos" numa receita de gotas.
 */
export function unitFromPharmaceuticalForm(form?: string | null): string {
  const f = (form || '').toLowerCase()
  if (!f) return 'unidades'
  if (f.includes('capsula')) return 'cápsulas'
  if (f.includes('comprimido') || f.includes('drage')) return 'comprimidos'
  if (f.includes('gotas')) return 'frascos'
  if (f.includes('sache')) return 'sachês'
  if (f.includes('supositorio')) return 'supositórios'
  if (f.includes('pastilha')) return 'pastilhas'
  if (f.includes('adesivo')) return 'adesivos'
  if (f.includes('ampola') || f.includes('injetavel')) return 'ampolas'
  if (f.includes('creme') || f.includes('pomada') || f.includes('gel') || f.includes('locao')) {
    return 'bisnagas'
  }
  if (f.includes('solucao') || f.includes('suspensao') || f.includes('xarope') || f.includes('spray')) {
    return 'frascos'
  }
  return 'unidades'
}

/** Um item novo do formulário, com os defaults que evitam campo vazio à toa. */
export function emptyMedication(): MedicationFormData {
  return {
    category: 'simple',
    medicationName: '',
    activeIngredient: '',
    concentration: '',
    dosage: '',
    frequency: '',
    route: 'oral',
    duration: 0,
    quantity: 0,
    quantityInWords: '',
    instructions: '',
    pharmaceuticalForm: '',
  }
}

export function emptyCommercialPrescription(): CommercialPrescriptionFormData {
  return { medications: [emptyMedication()], generalInstructions: '' }
}

/**
 * Formulário → payload da API. Vive aqui porque `new` e `edit` montavam o mesmo objeto de dois
 * jeitos ligeiramente diferentes, e o campo que divergia (`activeIngredient` vazio) era
 * exatamente o que o backend recusava.
 */
type MedicationRequestPayload = NonNullable<CreatePrescriptionRequest['medications']>[number]

export function toCreatePayload(
  data: CommercialPrescriptionFormData,
  patientId: string
): CreatePrescriptionRequest {
  return {
    patientId,
    type: 'commercial',
    medications: data.medications.map((m) => ({
      medicationDefinitionId: m.medicationDefinitionId,
      medicationName: m.medicationName,
      // Princípio ativo deixou de ser obrigatório na tela. Quando vem vazio, o nome comercial é
      // a melhor informação disponível — e é o que sai impresso.
      activeIngredient: m.activeIngredient || m.medicationName,
      category: m.category as MedicationRequestPayload['category'],
      concentration: m.concentration || '',
      dosage: m.dosage,
      frequency: m.frequency,
      route: m.route,
      duration: m.duration,
      quantity: m.quantity,
      quantityInWords: m.quantityInWords || '',
      instructions: m.instructions || undefined,
    })),
    generalInstructions: data.generalInstructions || undefined,
    prescriptionDate: new Date().toISOString(),
  }
}

/** Prescrição da API → valores do formulário (edição e "repetir receita"). */
export function fromPrescription(prescription: Prescription): CommercialPrescriptionFormData {
  return {
    medications: prescription.medications.map((m) => ({
      medicationDefinitionId: m.medicationDefinitionId,
      category: (m.category ?? 'simple') as MedicationCategory,
      medicationName: m.medicationName,
      activeIngredient: m.activeIngredient ?? '',
      concentration: m.concentration ?? '',
      dosage: m.dosage,
      frequency: m.frequency,
      route: m.route,
      duration: m.duration,
      quantity: m.quantity,
      quantityInWords: m.quantityInWords ?? '',
      instructions: m.instructions ?? '',
      pharmaceuticalForm: '',
    })),
    generalInstructions: prescription.generalInstructions ?? '',
  }
}
