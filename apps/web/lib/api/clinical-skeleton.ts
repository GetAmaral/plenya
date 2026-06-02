'use client';

/**
 * Esqueleto clínico P2a — alergias + sinais vitais por consulta.
 * Escopado por patientId (path), como score-snapshots. Tipos locais.
 */
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

export type AllergySubstanceType = 'drug' | 'food' | 'environmental' | 'other';
export type AllergySeverity = 'mild' | 'moderate' | 'severe' | 'anaphylaxis';
export type AllergyStatus = 'active' | 'inactive';

export interface PatientAllergy {
  id: string;
  patientId: string;
  substance: string;
  substanceType: AllergySubstanceType;
  reaction?: string;
  severity: AllergySeverity;
  status: AllergyStatus;
  noKnownAllergies: boolean;
  notes?: string;
  recordedByUserId: string;
  recordedByName?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateAllergyPayload {
  substance?: string;
  substanceType?: AllergySubstanceType;
  reaction?: string;
  severity?: AllergySeverity;
  noKnownAllergies?: boolean;
  notes?: string;
}

export type UpdateAllergyPayload = {
  substance?: string;
  substanceType?: AllergySubstanceType;
  reaction?: string;
  severity?: AllergySeverity;
  status?: AllergyStatus;
  notes?: string;
};

export interface ConsultationVitals {
  id: string;
  appointmentId?: string;
  patientId: string;
  systolicBp?: number;
  diastolicBp?: number;
  heartRate?: number;
  respRate?: number;
  temperature?: number;
  spo2?: number;
  weight?: number;
  height?: number;
  waistCircumference?: number;
  bmi?: number;
  measuredByUserId: string;
  measuredByName?: string;
  measuredAt: string;
  createdAt: string;
}

export interface CreateVitalsPayload {
  appointmentId?: string;
  systolicBp?: number;
  diastolicBp?: number;
  heartRate?: number;
  respRate?: number;
  temperature?: number;
  spo2?: number;
  weight?: number;
  height?: number;
  waistCircumference?: number;
  measuredAt?: string;
}

export const allergyKeys = {
  byPatient: (patientId: string) => ['allergies', patientId] as const,
};
export const vitalsKeys = {
  byPatient: (patientId: string, appointmentId?: string) =>
    ['vitals', patientId, appointmentId ?? null] as const,
};

// ===== Alergias =====
export function useAllergies(patientId: string | undefined, includeInactive = false) {
  return useQuery({
    queryKey: [...allergyKeys.byPatient(patientId ?? ''), includeInactive],
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<PatientAllergy[]>(
        `/api/v1/patients/${patientId}/allergies${includeInactive ? '?includeInactive=true' : ''}`,
      ),
  });
}

export function useCreateAllergy(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateAllergyPayload) =>
      apiClient.post<PatientAllergy>(`/api/v1/patients/${patientId}/allergies`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: allergyKeys.byPatient(patientId) }),
  });
}

export function useUpdateAllergy(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateAllergyPayload }) =>
      apiClient.put<PatientAllergy>(`/api/v1/patients/${patientId}/allergies/${id}`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: allergyKeys.byPatient(patientId) }),
  });
}

export function useDeleteAllergy(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/patients/${patientId}/allergies/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: allergyKeys.byPatient(patientId) }),
  });
}

// ===== Sinais vitais =====
export function useVitals(patientId: string | undefined, appointmentId?: string) {
  return useQuery({
    queryKey: vitalsKeys.byPatient(patientId ?? '', appointmentId),
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<ConsultationVitals[]>(
        `/api/v1/patients/${patientId}/vitals${appointmentId ? `?appointmentId=${appointmentId}` : ''}`,
      ),
  });
}

export function useCreateVitals(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateVitalsPayload) =>
      apiClient.post<ConsultationVitals>(`/api/v1/patients/${patientId}/vitals`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: ['vitals', patientId] }),
  });
}

// Cross-check de alergia ativa por princípio ativo (CDS não-bloqueante).
// Normaliza acento/caixa e casa por substring nos dois sentidos.
export function normalize(s: string): string {
  return (s || '')
    .normalize('NFD')
    .replace(/[̀-ͯ]/g, '')
    .toLowerCase()
    .trim();
}

export function matchAllergies(
  allergies: PatientAllergy[],
  ingredient: string,
): PatientAllergy[] {
  const ing = normalize(ingredient);
  if (!ing) return [];
  return allergies.filter((a) => {
    if (a.status !== 'active' || a.noKnownAllergies || !a.substance) return false;
    const sub = normalize(a.substance);
    return sub.length > 2 && (ing.includes(sub) || sub.includes(ing));
  });
}
