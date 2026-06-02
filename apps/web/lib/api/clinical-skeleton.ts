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

// ============================================================
// P2b — Problem list + Medicações em uso + CID-10
// ============================================================

export type ProblemStatus = 'active' | 'resolved' | 'inactive';

export interface PatientProblem {
  id: string;
  patientId: string;
  description: string;
  cidCode?: string;
  cidVersion: string;
  status: ProblemStatus;
  onsetDate?: string;
  resolvedDate?: string;
  onsetAppointmentId?: string;
  notes?: string;
  recordedByUserId: string;
  recordedByName?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateProblemPayload {
  description: string;
  cidCode?: string;
  onsetDate?: string;
  onsetAppointmentId?: string;
  notes?: string;
}
export interface UpdateProblemPayload {
  description?: string;
  cidCode?: string;
  status?: ProblemStatus;
  onsetDate?: string;
  resolvedDate?: string;
  notes?: string;
}

export type MedicationSource = 'prescribed_here' | 'external' | 'patient_reported';
export type MedicationInUseStatus = 'active' | 'suspended' | 'stopped';

export interface MedicationInUse {
  id: string;
  patientId: string;
  medicationName: string;
  activeIngredient: string;
  dosage: string;
  frequency: string;
  route: string;
  source: MedicationSource;
  status: MedicationInUseStatus;
  sourcePrescriptionId?: string;
  reconciledAppointmentId?: string;
  startDate?: string;
  endDate?: string;
  notes?: string;
  recordedByUserId: string;
  recordedByName?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateMedicationPayload {
  medicationName: string;
  activeIngredient?: string;
  dosage?: string;
  frequency?: string;
  route?: string;
  source?: MedicationSource;
  startDate?: string;
  notes?: string;
}
export interface UpdateMedicationPayload {
  medicationName?: string;
  activeIngredient?: string;
  dosage?: string;
  frequency?: string;
  route?: string;
  status?: MedicationInUseStatus;
  endDate?: string;
  notes?: string;
}

export interface CIDCode {
  code: string;
  description: string;
  chapter?: string;
  version: string;
}

export const problemKeys = { byPatient: (p: string) => ['problems', p] as const };
export const medicationKeys = { byPatient: (p: string) => ['medications-in-use', p] as const };

// ---- Problemas ----
export function useProblems(patientId: string | undefined, includeResolved = false) {
  return useQuery({
    queryKey: [...problemKeys.byPatient(patientId ?? ''), includeResolved],
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<PatientProblem[]>(
        `/api/v1/patients/${patientId}/problems${includeResolved ? '?includeResolved=true' : ''}`,
      ),
  });
}
export function useCreateProblem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateProblemPayload) =>
      apiClient.post<PatientProblem>(`/api/v1/patients/${patientId}/problems`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: problemKeys.byPatient(patientId) }),
  });
}
export function useUpdateProblem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateProblemPayload }) =>
      apiClient.put<PatientProblem>(`/api/v1/patients/${patientId}/problems/${id}`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: problemKeys.byPatient(patientId) }),
  });
}
export function useDeleteProblem(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete<void>(`/api/v1/patients/${patientId}/problems/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: problemKeys.byPatient(patientId) }),
  });
}

// ---- Medicações em uso ----
export function useMedicationsInUse(patientId: string | undefined, includeInactive = false) {
  return useQuery({
    queryKey: [...medicationKeys.byPatient(patientId ?? ''), includeInactive],
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<MedicationInUse[]>(
        `/api/v1/patients/${patientId}/medications${includeInactive ? '?includeInactive=true' : ''}`,
      ),
  });
}
export function useCreateMedication(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateMedicationPayload) =>
      apiClient.post<MedicationInUse>(`/api/v1/patients/${patientId}/medications`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: medicationKeys.byPatient(patientId) }),
  });
}
export function useUpdateMedication(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, payload }: { id: string; payload: UpdateMedicationPayload }) =>
      apiClient.put<MedicationInUse>(`/api/v1/patients/${patientId}/medications/${id}`, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: medicationKeys.byPatient(patientId) }),
  });
}
export function useDeleteMedication(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete<void>(`/api/v1/patients/${patientId}/medications/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: medicationKeys.byPatient(patientId) }),
  });
}

// ---- CID-10 autocomplete ----
export function useCIDSearch(q: string, enabled = true) {
  return useQuery({
    queryKey: ['cid-codes', q],
    enabled: enabled && q.trim().length >= 2,
    queryFn: () => apiClient.get<CIDCode[]>(`/api/v1/cid-codes?q=${encodeURIComponent(q.trim())}`),
  });
}
