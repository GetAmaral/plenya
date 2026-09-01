'use client';

/**
 * Plano de devolutiva do paciente (o "deck") — o mesmo conteúdo com três saídas: a tela do portal,
 * o PDF 16:9 (apresentar e mandar) e o PDF A4 paisagem (imprimir).
 *
 * O dossiê é o insumo DERIVADO do prontuário (réguas por exame, achados classificados e ordenados
 * por peso); o plano é o que alguém escreve em cima dele. Path-scoped por paciente, sempre.
 */

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type PatientPlanStatus = 'draft' | 'published';

/** Um slide. Espelha pdfdoc.DeckSlide — o contrato é o mesmo do JSONB. */
export interface DeckSlide {
  kind:
    | 'cover'
    | 'summary'
    | 'rulers'
    | 'two-cards'
    | 'plan-step'
    | 'sequence'
    | 'takeaway'
    | 'closing'
    | 'table';
  variant?: '' | 'dark' | 'deep';
  eyebrow?: string;
  title?: string;
  lede?: string;
  kicker?: string;
  source?: string;
  punch?: string;
  legend?: boolean;
  // Os blocos específicos são livres no front: quem os escreve é a skill, e o servidor valida.
  rulers?: unknown[];
  summary?: unknown;
  cards?: unknown[];
  steps?: unknown[];
  takeaway?: unknown;
  table?: unknown;
}

export interface PatientPlan {
  id: string;
  patientId: string;
  title: string;
  status: PatientPlanStatus;
  version: number;
  content: DeckSlide[];
  sourceSnapshotId?: string;
  authorUserId: string;
  publishedAt?: string;
  document16x9Id?: string;
  documentA4Id?: string;
  createdAt: string;
  updatedAt: string;
}

/** Um slide cujo conteúdo passou da moldura de 1920×1080. */
export interface DeckOverflow {
  slide: number;
  title: string;
  right: number;
  bottom: number;
}

export interface SavePatientPlanPayload {
  title?: string;
  content?: DeckSlide[];
  sourceSnapshotId?: string;
}

export const patientPlanKeys = {
  list: (patientId: string) => ['patient-plans', patientId] as const,
  one: (patientId: string, planId: string) => ['patient-plans', patientId, planId] as const,
};

const base = (patientId: string) => `/api/v1/patients/${patientId}/plans`;

export function usePatientPlans(patientId: string | undefined) {
  return useQuery({
    queryKey: patientPlanKeys.list(patientId ?? ''),
    enabled: !!patientId,
    queryFn: () => apiClient.get<PatientPlan[]>(base(patientId!)),
  });
}

export function usePatientPlan(patientId: string | undefined, planId: string | undefined) {
  return useQuery({
    queryKey: patientPlanKeys.one(patientId ?? '', planId ?? ''),
    enabled: !!patientId && !!planId,
    queryFn: () => apiClient.get<PatientPlan>(`${base(patientId!)}/${planId}`),
  });
}

export function useCreatePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: SavePatientPlanPayload) =>
      apiClient.post<PatientPlan>(base(patientId), payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

export function useUpdatePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: SavePatientPlanPayload }) =>
      apiClient.put<PatientPlan>(`${base(patientId)}/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) });
    },
  });
}

export function useDeletePatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete<void>(`${base(patientId)}/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

/** Mede quais slides transbordam. Lista vazia = pode publicar. */
export function useCheckPlanOverflow(patientId: string) {
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.get<{ slides: DeckOverflow[] }>(`${base(patientId)}/${id}/overflow`),
  });
}

/**
 * Publica no portal. Se algum slide transbordar o servidor responde 422 com a lista — conteúdo que
 * não cabe some do PDF em silêncio, então é bloqueio, não aviso.
 */
export function usePublishPatientPlan(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.post<PatientPlan>(`${base(patientId)}/${id}/publish`, {}),
    onSuccess: () => qc.invalidateQueries({ queryKey: patientPlanKeys.list(patientId) }),
  });
}

export const patientPlansApi = {
  previewURL: (patientId: string, planId: string) => `${base(patientId)}/${planId}/preview`,
};
