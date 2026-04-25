import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface PrescriptionSummary {
  id: string;
  patientId: string;
  issuedAt: string;
  status: 'draft' | 'signed' | 'cancelled';
  signedAt?: string;
}

export interface PrescriptionItem {
  medicationCode?: string;
  name: string;
  dosage: string;
  frequency: string;
  duration?: string;
  notes?: string;
}

export interface PrescriptionDetail extends PrescriptionSummary {
  items: PrescriptionItem[];
  signatureUrl?: string;
  pdfUrl?: string;
  notes?: string;
}

export const patientPrescriptionsOptions = (patientId: string) =>
  queryOptions({
    queryKey: queryKeys.patients.prescriptions(patientId),
    queryFn: ({ signal }) =>
      api.get<PrescriptionSummary[]>(`/api/v1/patients/${patientId}/prescriptions`, { signal }),
    enabled: Boolean(patientId),
  });

export const prescriptionOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.prescriptions.detail(id),
    queryFn: ({ signal }) =>
      api.get<PrescriptionDetail>(`/api/v1/prescriptions/${id}`, { signal }),
    enabled: Boolean(id),
  });
