import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface AnamnesisSummary {
  id: string;
  patientId: string;
  title: string;
  createdAt: string;
  updatedAt: string;
}

export interface AnamnesisDetail extends AnamnesisSummary {
  structuredJson?: string;
  freeText?: string;
  templateId?: string;
}

/**
 * Lista anamneses do paciente atualmente selecionado pelo profissional logado.
 * Backend escopa via User.SelectedPatientID — passe o patientId apenas para
 * compor o queryKey (cache-busting quando troca de paciente).
 */
export const patientAnamnesisOptions = (patientId: string) =>
  queryOptions({
    queryKey: queryKeys.patients.anamnesis(patientId),
    queryFn: ({ signal }) => api.get<AnamnesisSummary[]>('/api/v1/anamnesis', { signal }),
    enabled: Boolean(patientId),
  });

export const anamnesisDetailOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.anamnesis.detail(id),
    queryFn: ({ signal }) => api.get<AnamnesisDetail>(`/api/v1/anamnesis/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const anamnesisMutations = {
  create: (body: Partial<AnamnesisDetail>) =>
    api.post<AnamnesisDetail>('/api/v1/anamnesis', body),
  update: (id: string, body: Partial<AnamnesisDetail>) =>
    api.put<AnamnesisDetail>(`/api/v1/anamnesis/${id}`, body),
  remove: (id: string) => api.delete<void>(`/api/v1/anamnesis/${id}`),
};
