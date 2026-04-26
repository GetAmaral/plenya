import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export type PeriodizationFramework = 'bompa' | 'linear' | 'undulating' | 'block';

export interface PeriodizationMesocycle {
  id: string;
  name: string;
  weekStart: number;
  weekEnd: number;
  focus?: string;
  intensityPct?: number;
  volumePct?: number;
  notes?: string;
}

export interface PeriodizationSummary {
  id: string;
  patientId: string;
  framework: PeriodizationFramework;
  startDate: string;
  totalWeeks: number;
  goal?: string;
  status?: string;
}

export interface PeriodizationDetail extends PeriodizationSummary {
  mesocycles: PeriodizationMesocycle[];
  notes?: string;
}

const periodizationKeys = {
  all: () => [...queryKeys.all, 'periodizations'] as const,
  byPatient: (patientId: string) =>
    [...periodizationKeys.all(), 'patient', patientId] as const,
  detail: (id: string) => [...periodizationKeys.all(), 'detail', id] as const,
};

export const periodizationKeysFor = periodizationKeys;

export const patientPeriodizationsOptions = (patientId: string) =>
  queryOptions({
    queryKey: periodizationKeys.byPatient(patientId),
    queryFn: ({ signal }) =>
      api.get<PeriodizationSummary[]>('/api/v1/periodizations', { signal }),
    enabled: Boolean(patientId),
  });

export const periodizationDetailOptions = (id: string) =>
  queryOptions({
    queryKey: periodizationKeys.detail(id),
    queryFn: ({ signal }) =>
      api.get<PeriodizationDetail>(`/api/v1/periodizations/${id}`, { signal }),
    enabled: Boolean(id),
  });
