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

// ============= Fitness Tests + Postural Assessments =============

export type FitnessTestKind = 'abdominal' | 'pushup' | 'plank' | 'burpee' | 'frt';

export const fitnessTestLabels: Record<FitnessTestKind, string> = {
  abdominal: 'Abdominal (1 min)',
  pushup: 'Flexão (1 min)',
  plank: 'Prancha (s)',
  burpee: 'Burpee (1 min)',
  frt: 'Functional Reach (cm)',
};

export interface FitnessTestResult {
  id: string;
  patientId: string;
  performedAt: string;
  kind: FitnessTestKind;
  value: number;
  unit?: string;
  classification?: string;
  notes?: string;
}

export interface CreateFitnessTestInput {
  kind: FitnessTestKind;
  value: number;
  unit?: string;
  notes?: string;
}

export interface PosturalMeasurement {
  region: string;
  angleDeg: number;
  withinNormal?: boolean;
}

export interface PosturalAssessment {
  id: string;
  patientId: string;
  performedAt: string;
  measurements: PosturalMeasurement[];
  totalPenalty?: number;
  notes?: string;
}

export interface CreatePosturalAssessmentInput {
  measurements: PosturalMeasurement[];
  notes?: string;
}

const fitnessKeys = {
  all: () => [...queryKeys.all, 'fitness-tests'] as const,
  byPatient: (patientId: string) => [...fitnessKeys.all(), 'patient', patientId] as const,
};

const posturalKeys = {
  all: () => [...queryKeys.all, 'postural-assessments'] as const,
  byPatient: (patientId: string) => [...posturalKeys.all(), 'patient', patientId] as const,
};

export const fitnessKeysFor = fitnessKeys;
export const posturalKeysFor = posturalKeys;

export const patientFitnessTestsOptions = (patientId: string) =>
  queryOptions({
    queryKey: fitnessKeys.byPatient(patientId),
    queryFn: ({ signal }) =>
      api.get<FitnessTestResult[]>('/api/v1/fitness-tests', { signal }),
    enabled: Boolean(patientId),
  });

export const patientPosturalAssessmentsOptions = (patientId: string) =>
  queryOptions({
    queryKey: posturalKeys.byPatient(patientId),
    queryFn: ({ signal }) =>
      api.get<PosturalAssessment[]>('/api/v1/postural-assessments', { signal }),
    enabled: Boolean(patientId),
  });

export const fitnessTestMutations = {
  create: (body: CreateFitnessTestInput) =>
    api.post<FitnessTestResult>('/api/v1/fitness-tests', body),
};

export const posturalMutations = {
  create: (body: CreatePosturalAssessmentInput) =>
    api.post<PosturalAssessment>('/api/v1/postural-assessments', body),
};
