import { apiClient } from '../api-client'
import type { MedicationCategory } from '@/lib/validations/prescription'

/**
 * Fórmulas-base: a fórmula pronta que se reusa entre pacientes, e as regras que sugerem dose a
 * partir do prontuário.
 *
 * A sugestão vem de um endpoint SEPARADO da criação de receita, de propósito: o que a tela
 * recebe é um número com a base escrita ao lado, para o médico aceitar ou ignorar. Nada aqui
 * escreve na prescrição.
 */

export type DoseRuleKind = 'fixed' | 'per_kg' | 'lab_threshold' | 'lab_band'
export type LabOperator = 'lt' | 'lte' | 'gt' | 'gte'

/**
 * Uma faixa do exame. Intervalo meio-aberto (lowerBound, upperBound], a mesma convenção das
 * faixas do escore: nulo em lowerBound é -infinito, nulo em upperBound é +infinito.
 */
export interface DoseBand {
  id?: string
  lowerBound?: number | null
  upperBound?: number | null
  dose: number
  label?: string
}

export interface DoseRule {
  id?: string
  kind: DoseRuleKind
  perKg?: number
  labCode?: string
  labOperator?: LabOperator
  labThreshold?: number
  doseIfTrue?: number
  doseIfFalse?: number
  fixedDose?: number
  /** Conduta por faixa — como se dosa vitamina D, B12 e zinco de verdade. */
  bands?: DoseBand[]
  /** Unidade em que o limiar/faixa foi escrito. Resultado em outra unidade não sugere nada. */
  labUnit?: string
  /** Passo prático de arredondamento (100 UI, 50 mg). */
  roundTo?: number
  /** Piso e teto são obrigatórios: é a trava que impede dado errado de virar dose absurda. */
  minDose: number
  maxDose: number
  maxDataAgeDays?: number
  note?: string
}

export interface FormulaTemplateComponent {
  id?: string
  magistralComponentId?: string
  substance: string
  quantity: number
  unit: string
  category?: MedicationCategory
  note?: string
  rule?: DoseRule | null
}

export interface FormulaTemplate {
  id: string
  name: string
  indication?: string
  indicationBullets?: string
  pharmaceuticalForm: string
  usageType: 'internal' | 'external'
  route?: string
  vehicle?: string
  quantityToDispense: number
  quantityUnit: string
  posology?: string
  duration?: number
  instructions?: string
  notes?: string
  usageCount: number
  isActive: boolean
  /** Nulo = ninguém conferiu ainda. Carimba só quando um humano salva pela tela. */
  lastReview?: string
  components?: FormulaTemplateComponent[]
}

export interface DoseSuggestionItem {
  templateComponentId: string
  substance: string
  unit: string
  baseDose: number
  /** Ausente quando não houve como sugerir; `reason` diz por quê. */
  suggested?: number
  basis?: string
  clamped: boolean
  reason?: string
  lastPrescribed?: number
  /** A escada inteira da regra, com a faixa deste paciente marcada. */
  bands?: DoseBandView[]
}

export interface DoseBandView {
  lowerBound?: number | null
  upperBound?: number | null
  dose: number
  label?: string
  active: boolean
}

export interface DoseSuggestion {
  templateId: string
  templateName: string
  weightKg?: number
  weightSource?: string
  weightDate?: string
  items: DoseSuggestionItem[]
}

export async function listFormulaTemplates(query = ''): Promise<FormulaTemplate[]> {
  const qs = query.trim() ? `?q=${encodeURIComponent(query.trim())}` : ''
  return apiClient.get<FormulaTemplate[]>(`/api/v1/magistral-templates${qs}`)
}

export async function getFormulaTemplate(id: string): Promise<FormulaTemplate> {
  return apiClient.get<FormulaTemplate>(`/api/v1/magistral-templates/${id}`)
}

export async function createFormulaTemplate(
  body: Omit<FormulaTemplate, 'id' | 'usageCount' | 'isActive'>
): Promise<FormulaTemplate> {
  return apiClient.post<FormulaTemplate>('/api/v1/magistral-templates', body)
}

export async function updateFormulaTemplate(
  id: string,
  body: Omit<FormulaTemplate, 'id' | 'usageCount' | 'isActive'>
): Promise<FormulaTemplate> {
  return apiClient.put<FormulaTemplate>(`/api/v1/magistral-templates/${id}`, body)
}

export async function deleteFormulaTemplate(id: string): Promise<void> {
  await apiClient.delete(`/api/v1/magistral-templates/${id}`)
}

/** Sugestões da fórmula-base para um paciente. Não escreve nada. */
export async function suggestDoses(templateId: string, patientId: string): Promise<DoseSuggestion> {
  return apiClient.get<DoseSuggestion>(
    `/api/v1/magistral-templates/${templateId}/suggest?patientId=${patientId}`
  )
}
