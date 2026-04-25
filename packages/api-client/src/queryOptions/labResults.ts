import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface LabResultSummary {
  id: string;
  patientId: string;
  collectedAt: string;
  labName?: string;
  status: 'pending' | 'completed';
}

export interface LabResultValue {
  testCode: string;
  testName: string;
  value: number | string;
  unit?: string;
  referenceMin?: number;
  referenceMax?: number;
  flag?: 'low' | 'normal' | 'high' | 'critical';
}

export interface LabResultDetail extends LabResultSummary {
  values: LabResultValue[];
  attachments?: Array<{ id: string; url: string; type: string }>;
  notes?: string;
}

/**
 * Lista resultados de exames do paciente atualmente selecionado.
 * Backend escopa via User.SelectedPatientID; patientId compõe queryKey.
 */
export const patientLabResultsOptions = (patientId: string) =>
  queryOptions({
    queryKey: queryKeys.patients.labResults(patientId),
    queryFn: ({ signal }) => api.get<LabResultSummary[]>('/api/v1/lab-results', { signal }),
    enabled: Boolean(patientId),
  });

export const labResultOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.labResults.detail(id),
    queryFn: ({ signal }) => api.get<LabResultDetail>(`/api/v1/lab-results/${id}`, { signal }),
    enabled: Boolean(id),
  });
