'use client';

/**
 * Notas de evolução clínica (SOAP/APSO) por consulta.
 *
 * Distinta da anamnese (intake one-shot): é a evolução de cada visita,
 * ancorada ao appointment. Imutável após assinada (correção via adendo).
 *
 * Tipos definidos localmente (mesmo padrão de calendar-api.ts).
 */
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type ClinicalNoteLayout = 'soap' | 'apso';
export type ClinicalNoteStatus = 'draft' | 'signed';
export type ClinicalNoteVisibility = 'all' | 'medicalOnly' | 'psychOnly' | 'authorOnly';

export interface ClinicalNote {
  id: string;
  appointmentId?: string;
  patientId: string;
  authorId: string;
  layout: ClinicalNoteLayout;
  subjective?: string;
  subjectiveHtml?: string;
  objective?: string;
  objectiveHtml?: string;
  assessment?: string;
  assessmentHtml?: string;
  plan?: string;
  planHtml?: string;
  status: ClinicalNoteStatus;
  signedAt?: string;
  amendmentOfId?: string;
  visibility: ClinicalNoteVisibility;
  displayTitle: string;
  author?: { id: string; name: string; email: string; role: string };
  createdAt: string;
  updatedAt: string;
}

export interface CreateClinicalNotePayload {
  appointmentId?: string;
  patientId?: string;
  layout?: ClinicalNoteLayout;
  subjective?: string;
  subjectiveHtml?: string;
  objective?: string;
  objectiveHtml?: string;
  assessment?: string;
  assessmentHtml?: string;
  plan?: string;
  planHtml?: string;
  visibility?: ClinicalNoteVisibility;
  sign?: boolean;
}

export type UpdateClinicalNotePayload = Omit<CreateClinicalNotePayload, 'appointmentId' | 'patientId' | 'sign'>;

export const clinicalNoteKeys = {
  all: ['clinical-notes'] as const,
  byAppointment: (appointmentId: string) =>
    [...clinicalNoteKeys.all, 'appointment', appointmentId] as const,
  detail: (id: string) => [...clinicalNoteKeys.all, 'detail', id] as const,
};

/**
 * Nota da consulta (ou null se ainda não existe).
 * O backend retorna 404 quando não há nota — tratamos como null (estado normal).
 */
export function useClinicalNoteByAppointment(appointmentId: string | undefined) {
  return useQuery({
    queryKey: clinicalNoteKeys.byAppointment(appointmentId ?? ''),
    enabled: !!appointmentId,
    queryFn: async (): Promise<ClinicalNote | null> => {
      try {
        return await apiClient.get<ClinicalNote>(
          `/api/v1/clinical-notes?appointmentId=${appointmentId}`,
        );
      } catch (err) {
        if ((err as { status?: number }).status === 404) return null;
        throw err;
      }
    },
  });
}

export function useCreateClinicalNote() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateClinicalNotePayload) =>
      apiClient.post<ClinicalNote>('/api/v1/clinical-notes', payload),
    onSuccess: (note) => {
      if (note.appointmentId) {
        qc.invalidateQueries({ queryKey: clinicalNoteKeys.byAppointment(note.appointmentId) });
      }
      qc.invalidateQueries({ queryKey: clinicalNoteKeys.all });
    },
  });
}

export function useUpdateClinicalNote() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateClinicalNotePayload }) =>
      apiClient.put<ClinicalNote>(`/api/v1/clinical-notes/${id}`, payload),
    onSuccess: (note) => {
      if (note.appointmentId) {
        qc.invalidateQueries({ queryKey: clinicalNoteKeys.byAppointment(note.appointmentId) });
      }
      qc.invalidateQueries({ queryKey: clinicalNoteKeys.all });
    },
  });
}

export function useSignClinicalNote() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.post<ClinicalNote>(`/api/v1/clinical-notes/${id}/sign`),
    onSuccess: (note) => {
      if (note.appointmentId) {
        qc.invalidateQueries({ queryKey: clinicalNoteKeys.byAppointment(note.appointmentId) });
      }
      qc.invalidateQueries({ queryKey: clinicalNoteKeys.all });
    },
  });
}
