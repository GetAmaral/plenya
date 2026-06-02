'use client';

/**
 * Plano de cuidado AGIR (P3 frente 3) — recomendações ancoradas por pilar (letra A/G/I/R)
 * + biomarcador opcional. Path-scoped por paciente. Usado na capa e no workspace.
 */

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type AgirLetter = 'A' | 'G' | 'I' | 'R';
export type CarePlanPriority = 'high' | 'medium' | 'low';
export type CarePlanItemStatus = 'active' | 'achieved' | 'suspended';

export interface CarePlanItem {
  id: string;
  patientId: string;
  letterCode: AgirLetter;
  recommendation: string;
  rationale?: string;
  target?: string;
  labTestCode?: string;
  scoreItemId?: string;
  sourceSnapshotId?: string;
  priority: CarePlanPriority;
  status: CarePlanItemStatus;
  authorUserId: string;
  authorRole: string;
  authorName?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateCarePlanItemPayload {
  letterCode: AgirLetter;
  recommendation: string;
  rationale?: string;
  target?: string;
  labTestCode?: string;
  priority?: CarePlanPriority;
}

export type UpdateCarePlanItemPayload = Partial<{
  letterCode: AgirLetter;
  recommendation: string;
  rationale: string;
  target: string;
  labTestCode: string;
  priority: CarePlanPriority;
  status: CarePlanItemStatus;
}>;

export const carePlanKeys = {
  byPatient: (patientId: string) => ['care-plan', patientId] as const,
};

export function useCarePlan(patientId: string | undefined, includeInactive = false) {
  return useQuery({
    queryKey: [...carePlanKeys.byPatient(patientId ?? ''), includeInactive],
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<CarePlanItem[]>(
        `/api/v1/patients/${patientId}/care-plan-items${includeInactive ? '?includeInactive=true' : ''}`,
      ),
  });
}

export function useCreateCarePlanItem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateCarePlanItemPayload) =>
      apiClient.post<CarePlanItem>(`/api/v1/patients/${patientId}/care-plan-items`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: carePlanKeys.byPatient(patientId) }),
  });
}

export function useUpdateCarePlanItem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateCarePlanItemPayload }) =>
      apiClient.put<CarePlanItem>(`/api/v1/patients/${patientId}/care-plan-items/${id}`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: carePlanKeys.byPatient(patientId) }),
  });
}

export function useDeleteCarePlanItem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/patients/${patientId}/care-plan-items/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: carePlanKeys.byPatient(patientId) }),
  });
}

/** Gera + assina + publica o relatório longitudinal AGIR no portal (frente 3b). */
export function useGenerateCarePlanReport(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () =>
      apiClient.post<{ issuedDocumentId: string }>(`/api/v1/patients/${patientId}/care-plan-report`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ['issued-documents', patientId] });
    },
  });
}
