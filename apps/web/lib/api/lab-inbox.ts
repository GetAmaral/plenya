'use client';

/**
 * Results inbox (P3 frente 1) — fila cross-patient de "Exames a revisar".
 * Distinta do lab-result-batch-api (scoped ao selectedPatient): aqui a fila
 * cruza TODOS os pacientes, críticos no topo, com ciência clínica auditada.
 * Endpoints: GET /lab-result-inbox, /count, /:id, POST /:id/acknowledge.
 */

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export interface LabInboxItem {
  id: string;
  patientId: string;
  patientName: string;
  laboratoryName: string;
  collectionDate: string;
  resultDate?: string;
  status: string;
  isCritical: boolean;
  worstLevel?: number;
  abnormalCount: number;
  totalResults: number;
  reviewedAt?: string;
}

export interface LabInboxCount {
  total: number;
  critical: number;
}

// Subconjunto do LabResultBatchDetailResponse que a tela de revisão consome.
export interface LabInboxResultRow {
  id: string;
  testName: string;
  testType: string;
  resultText?: string;
  resultNumeric?: number;
  unit?: string;
  interpretation?: string;
  level?: number;
  labTestDefinition?: { name: string; code: string; unit?: string };
}

export interface LabInboxDetail {
  id: string;
  patientId: string;
  laboratoryName: string;
  collectionDate: string;
  resultDate?: string;
  status: string;
  isCritical: boolean;
  worstLevel?: number;
  reviewedAt?: string;
  resultCount: number;
  labResults: LabInboxResultRow[];
}

export const labInboxKeys = {
  list: ['lab-inbox'] as const,
  count: ['lab-inbox', 'count'] as const,
  detail: (id: string) => ['lab-inbox', 'detail', id] as const,
};

export function useLabInbox() {
  return useQuery({
    queryKey: labInboxKeys.list,
    queryFn: () => apiClient.get<LabInboxItem[]>('/api/v1/lab-result-inbox?limit=100'),
    staleTime: 15_000,
    refetchInterval: 30_000,
    refetchOnWindowFocus: true,
  });
}

/** Contador para o badge do sidebar. Gateado por `enabled` p/ rodar só na linha "Exames a revisar". */
export function useLabReviewCount(opts?: { enabled?: boolean }) {
  return useQuery({
    queryKey: labInboxKeys.count,
    queryFn: () => apiClient.get<LabInboxCount>('/api/v1/lab-result-inbox/count'),
    enabled: opts?.enabled ?? true,
    staleTime: 15_000,
    refetchInterval: 30_000,
    refetchOnWindowFocus: true,
  });
}

export function useLabInboxDetail(id: string | undefined) {
  return useQuery({
    queryKey: labInboxKeys.detail(id ?? ''),
    queryFn: () => apiClient.get<LabInboxDetail>(`/api/v1/lab-result-inbox/${id}`),
    enabled: !!id,
  });
}

export function useAcknowledgeBatch() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (batchId: string) =>
      apiClient.post<LabInboxDetail>(`/api/v1/lab-result-inbox/${batchId}/acknowledge`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: labInboxKeys.list });
      qc.invalidateQueries({ queryKey: labInboxKeys.count });
    },
  });
}
