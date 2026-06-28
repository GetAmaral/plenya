import { useQuery } from '@tanstack/react-query'
import { apiClient } from '../api-client'

// Origem de uma entrada do histórico clínico.
export type ClinicalTimelineSource = 'escore_light' | 'pre_consulta' | 'anamnese'

export interface ClinicalTimelineEntry {
  date: string
  source: ClinicalTimelineSource
  sourceLabel: string
  selectedLevel?: number
  numericValue?: number
  textValue?: string
  scaleResponses?: { answers?: Record<string, number>; total?: number; words?: string[] }
}

export interface ClinicalItemHistoryResponse {
  scoreItemId: string
  entries: ClinicalTimelineEntry[]
}

export function getItemHistory(patientId: string, scoreItemId: string) {
  return apiClient.get<ClinicalItemHistoryResponse>(
    `/api/v1/patients/${patientId}/score-items/${scoreItemId}/history`
  )
}

// Histórico longitudinal de um item clínico para um paciente (Escore Light + pré-consulta +
// anamnese). Lazy: só dispara quando habilitado (item expandido) e com paciente conhecido.
export function useItemHistory(
  patientId: string | null | undefined,
  scoreItemId: string,
  enabled: boolean
) {
  return useQuery({
    queryKey: ['clinical-timeline', patientId, scoreItemId],
    queryFn: () => getItemHistory(patientId!, scoreItemId),
    enabled: !!patientId && !!scoreItemId && enabled,
    staleTime: 60_000,
  })
}
