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
  burpee: 'Burpee (3 min)',
  frt: 'FRT (90s)',
};

/**
 * Reflete `models.FitnessTestResult` Go.
 * Campos `*Reps`/`*Seconds` são input; `*Level`, `overallScore`,
 * `overallClassification` são computados pelo serviço.
 */
export interface FitnessTestResult {
  id: string;
  patientId: string;
  createdById: string;
  assessmentDate: string;
  abdominalReps?: number;
  pushupReps?: number;
  plankSeconds?: number;
  burpeeCycles?: number;
  frtReps?: number;
  abdominalLevel?: string;
  pushupLevel?: string;
  plankLevel?: string;
  burpeeLevel?: string;
  frtLevel?: string;
  overallScore: number;
  overallClassification: string;
  notes?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateFitnessTestInput {
  patientId: string;
  assessmentDate: string;
  abdominalReps?: number;
  pushupReps?: number;
  plankSeconds?: number;
  burpeeCycles?: number;
  frtReps?: number;
  notes?: string;
}

export type PosturalViewType = 'front' | 'side_left' | 'side_right' | 'back';

export const posturalViewTypeLabels: Record<PosturalViewType, string> = {
  front: 'Vista frontal',
  side_left: 'Lateral esquerda',
  side_right: 'Lateral direita',
  back: 'Vista posterior',
};

/**
 * Reflete `models.PosturalAssessment` Go.
 * Mede ângulos em graus; `posturalScore`/`posturalClassification`/
 * `severeDeviations` são computados.
 */
export interface PosturalAssessment {
  id: string;
  patientId: string;
  createdById: string;
  assessmentDate: string;
  physicalAssessmentId?: string;
  viewType: PosturalViewType;
  shoulderDeviation?: number;
  hipDeviation?: number;
  headLateralDeviation?: number;
  fhp?: number;
  thoracicKyphosis?: number;
  lumbarLordosis?: number;
  kneeAngle?: number;
  photoUrl?: string;
  posturalScore: number;
  posturalClassification: string;
  severeDeviations: number;
  notes?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreatePosturalAssessmentInput {
  patientId: string;
  assessmentDate: string;
  viewType: PosturalViewType;
  shoulderDeviation?: number;
  hipDeviation?: number;
  headLateralDeviation?: number;
  fhp?: number;
  thoracicKyphosis?: number;
  lumbarLordosis?: number;
  kneeAngle?: number;
  photoUrl?: string;
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
