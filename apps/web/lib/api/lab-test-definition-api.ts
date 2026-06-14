import { apiClient } from '../api-client'
// LabTestDefinition vem do GERADO (Go model → swag → openapi-typescript), via @plenya/types.
import type { LabTestDefinition } from '@plenya/types'

export type { LabTestDefinition }

// Opção de exame qualitativo (ex.: genótipo) — nível + rótulo, vindo dos score_levels.
export interface QualitativeOption {
  id: string
  level: number
  name: string
}

export const labTestDefinitionApi = {
  // Opções (níveis) de um exame qualitativo por lab_test_code — ex.: genótipos de um gene.
  // Retorna [] quando o exame não tem score item vinculado (não é qualitativo pontuável).
  getQualitativeOptions: async (code: string) =>
    apiClient.get<QualitativeOption[]>(
      `/api/v1/score-items/by-lab-code/${encodeURIComponent(code)}/levels`
    ),

  // Busca TODOS os exames (para edição de resultados)
  getAll: async () =>
    apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/definitions'),

  // Busca apenas exames que podem ser solicitados (para formulários)
  getRequestable: async () =>
    apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/requestable'),

  // Busca por termo (nome, código, categoria)
  search: async (query: string) =>
    apiClient.get<LabTestDefinition[]>(
      `/api/v1/lab-tests/definitions/search?q=${encodeURIComponent(query)}`
    ),

  // Busca por ID
  getById: async (id: string) =>
    apiClient.get<LabTestDefinition>(`/api/v1/lab-tests/definitions/${id}`),
}
