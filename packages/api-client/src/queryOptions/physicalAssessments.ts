import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface PhysicalAssessmentSummary {
  id: string;
  patientId: string;
  performedAt: string;
  weightKg?: number;
  heightCm?: number;
  bmi?: number;
  bloodPressureSystolic?: number;
  bloodPressureDiastolic?: number;
}

export interface ACSMTag {
  label: string;
  category: string;
  color?: string;
}

export interface PhysicalAssessmentDetail extends PhysicalAssessmentSummary {
  notes?: string;
  acsmTagsStructured?: ACSMTag[];
  htmlContent?: string;
}

const physicalAssessmentsKeys = {
  all: () => [...queryKeys.all, 'physical-assessments'] as const,
  byPatient: (patientId: string) =>
    [...physicalAssessmentsKeys.all(), 'patient', patientId] as const,
  detail: (id: string) => [...physicalAssessmentsKeys.all(), 'detail', id] as const,
};

export const physicalAssessmentsKeysFor = physicalAssessmentsKeys;

/**
 * Lista avaliações físicas do paciente atualmente selecionado.
 * Backend escopa via User.SelectedPatientID; patientId compõe queryKey.
 */
export const patientPhysicalAssessmentsOptions = (patientId: string) =>
  queryOptions({
    queryKey: physicalAssessmentsKeys.byPatient(patientId),
    queryFn: ({ signal }) =>
      api.get<PhysicalAssessmentSummary[]>('/api/v1/physical-assessments', { signal }),
    enabled: Boolean(patientId),
  });

export const physicalAssessmentDetailOptions = (id: string) =>
  queryOptions({
    queryKey: physicalAssessmentsKeys.detail(id),
    queryFn: ({ signal }) =>
      api.get<PhysicalAssessmentDetail>(`/api/v1/physical-assessments/${id}`, { signal }),
    enabled: Boolean(id),
  });

export interface CreatePhysicalAssessmentInput {
  weightKg?: number;
  heightCm?: number;
  bmi?: number;
  bloodPressureSystolic?: number;
  bloodPressureDiastolic?: number;
  notes?: string;
  /** URLs/paths de fotos já uploadadas via /api/v1/uploads */
  photoUrls?: string[];
}

export const physicalAssessmentMutations = {
  create: (body: CreatePhysicalAssessmentInput) =>
    api.post<PhysicalAssessmentDetail>('/api/v1/physical-assessments', body),
};
