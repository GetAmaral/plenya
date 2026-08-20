import { apiClient } from '../api-client'

/**
 * Catálogo magistral: busca de componentes, curadoria oportunista e a checagem de fórmula
 * (compatibilidade, palatabilidade e tamanho de cápsula) que a tela de manipulado chama enquanto
 * o médico monta a receita.
 */

export interface MagistralComponent {
  id: string
  name: string
  synonyms?: string
  defaultUnit: string
  usualDose?: number
  minDose?: number
  maxDose?: number
  bulkDensity?: number
  /** 'medida' · 'classe' (aproximação) · texto da ficha técnica. */
  densitySource?: string
  brand?: string
  /** Percentual de ativo no insumo (bisglicinato 30%, selenometionina 1%). */
  elementalPercent?: number
  correctionNote?: string
  preferredAlternativeId?: string
  preferenceNote?: string
  eutecticFormer: boolean
  hygroscopic: boolean
  oxidizing: boolean
  oxidationSensitive: boolean
  photosensitive: boolean
  bitterness?: number
  sachetOk?: boolean
  notes?: string
  /** Para que serve, em prosa curta. Vem da curadoria ou do RAG (ver evidenceStatus). */
  indications?: string
  /** Posologia em texto, com a origem embutida. */
  doseReference?: string
  /** Os mesmos campos em tópicos, uma linha por item. É o que a tela mostra por default. */
  indicationBullets?: string
  doseBullets?: string
  /** pending · suggested (veio do RAG, aguarda conferência) · confirmed (o médico conferiu) */
  evidenceStatus?: 'pending' | 'suggested' | 'confirmed'
  enrichedAt?: string
  source: string
  usageCount: number
}

/** Trecho de aula/artigo do RAG ligado ao componente. Material de leitura, não de cálculo. */
export interface MagistralEvidence {
  id: string
  componentId: string
  articleId: string
  chunkIndex?: number
  similarity?: number
  excerpt: string
  pinned: boolean
  article?: { id: string; title: string; journal?: string; articleType?: string }
}

export interface MagistralAlert {
  level: 'info' | 'warn' | 'avoid'
  kind:
    | 'eutectic'
    | 'oxidation'
    | 'hygroscopic'
    | 'palatability'
    | 'pair'
    | 'dose'
    | 'correction'
    | 'in28'
    | 'base'
  substances: string[]
  message: string
}

export interface CapsuleAdvice {
  /** Falso = falta densidade cadastrada; `missing` diz de quem. A tela não inventa número. */
  decided: boolean
  missing?: string[]
  volumeMl?: number
  volumeMinMl?: number
  volumeMaxMl?: number
  size?: string
  capsulesPerDose?: number
  sizeIfCompacted?: string
  sachetRecommended: boolean
  explanation: string
  /** Algum componente usou densidade aproximada por classe. */
  approximate?: boolean
}

export interface MagistralComponentMatch {
  id?: string
  substance: string
  known: boolean
  hasDose: boolean
  hasDensity: boolean
  usualDose?: number
  defaultUnit?: string
  indications?: string
  doseReference?: string
  indicationBullets?: string
  doseBullets?: string
  evidenceStatus?: 'pending' | 'suggested' | 'confirmed'
  elementalPercent?: number
  /** Massa do insumo correspondente à dose do elemento, em mg. */
  rawMaterialMg?: number
  correctionNote?: string
  /** Forma que o prescritor usa no lugar desta. */
  preferredName?: string
  preferenceNote?: string
  /** Categoria de receita que a substância carrega (c5 em anabolizante, por exemplo). */
  defaultCategory?: string
  /** Restrição de finalidade ou exigência regulatória da substância. */
  regulatoryNote?: string
}

/** Uma linha por tópico. Linha vazia não vira item. */
export function toBullets(value?: string): string[] {
  return (value ?? '')
    .split('\n')
    .map((l) => l.replace(/^[-•·]\s*/, '').trim())
    .filter(Boolean)
}

export interface MagistralCheckResponse {
  alerts: MagistralAlert[]
  capsule?: CapsuleAdvice
  components: MagistralComponentMatch[]
}

export interface MagistralCheckRequest {
  pharmaceuticalForm: string
  components: {
    magistralComponentId?: string
    substance: string
    quantity: number
    unit: string
    asElemental?: boolean
  }[]
  /** Posologia da fórmula: o backend deduz dela quantas tomadas o dia tem. */
  posology?: string
  /** Veículo/base — é contra ele que rodam as regras de incompatibilidade de base. */
  vehicle?: string
}

export async function searchMagistralComponents(
  query: string,
  { allowEmpty = true }: { allowEmpty?: boolean } = {}
): Promise<MagistralComponent[]> {
  const q = query.trim()
  // Vazio lista o repertório (tela de curadoria); 1 caractere não busca nada em lugar nenhum.
  if (q.length === 0 && !allowEmpty) return []
  if (q.length === 1) return []
  return apiClient.get<MagistralComponent[]>(
    `/api/v1/magistral-components/search?q=${encodeURIComponent(query)}`
  )
}

export async function checkMagistralFormula(
  body: MagistralCheckRequest
): Promise<MagistralCheckResponse> {
  return apiClient.post<MagistralCheckResponse>('/api/v1/magistral-components/check', body)
}

export async function saveMagistralDefaultDose(body: {
  substance: string
  dose: number
  unit: string
}): Promise<MagistralComponent> {
  return apiClient.post<MagistralComponent>('/api/v1/magistral-components/default-dose', body)
}

export async function updateMagistralComponent(
  id: string,
  body: Partial<MagistralComponent> & { name: string }
): Promise<MagistralComponent> {
  return apiClient.put<MagistralComponent>(`/api/v1/magistral-components/${id}`, body)
}

export async function getMagistralEvidence(componentId: string): Promise<MagistralEvidence[]> {
  return apiClient.get<MagistralEvidence[]>(`/api/v1/magistral-components/${componentId}/evidence`)
}

/** Aceite do médico: o conteúdo extraído do RAG deixa de ser sugestão. */
export async function confirmMagistralComponent(componentId: string): Promise<void> {
  await apiClient.post(`/api/v1/magistral-components/${componentId}/confirm`)
}

export async function pinMagistralEvidence(evidenceId: string, pinned: boolean): Promise<void> {
  await apiClient.post(`/api/v1/magistral-components/evidence/${evidenceId}/pin?pinned=${pinned}`)
}
