import { infiniteQueryOptions, queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface PatientListParams {
  search?: string;
  limit?: number;
  cursor?: string;
}

export interface PatientSummary {
  id: string;
  name: string;
  email?: string;
  phone?: string;
  birthDate?: string;
  gender?: string;
  avatarUrl?: string;
  cpfMasked?: string;
  lastVisitAt?: string;
}

export interface PatientDetail extends PatientSummary {
  addressJson?: string;
  rgMasked?: string;
  notes?: string;
  createdAt: string;
  updatedAt: string;
}

export interface PaginatedResponse<T> {
  items: T[];
  nextCursor?: string;
  total?: number;
}

export const patientsListOptions = (params: PatientListParams = {}) =>
  infiniteQueryOptions({
    queryKey: queryKeys.patients.list(params as Record<string, unknown>),
    queryFn: ({ pageParam, signal }) => {
      const qs = new URLSearchParams();
      if (params.search) qs.set('search', params.search);
      if (params.limit) qs.set('limit', String(params.limit));
      if (pageParam) qs.set('cursor', String(pageParam));
      const suffix = qs.toString() ? `?${qs.toString()}` : '';
      return api.get<PaginatedResponse<PatientSummary>>(`/api/v1/patients${suffix}`, { signal });
    },
    initialPageParam: undefined as string | undefined,
    getNextPageParam: (last) => last.nextCursor,
  });

export const patientOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.patients.detail(id),
    queryFn: ({ signal }) => api.get<PatientDetail>(`/api/v1/patients/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const patientMutations = {
  create: (body: Partial<PatientDetail>) => api.post<PatientDetail>('/api/v1/patients', body),
  update: (id: string, body: Partial<PatientDetail>) =>
    api.put<PatientDetail>(`/api/v1/patients/${id}`, body),
  remove: (id: string) => api.delete<void>(`/api/v1/patients/${id}`),
  setSelected: (id: string) => api.post<void>(`/api/v1/me/selected-patient`, { patientId: id }),
};
