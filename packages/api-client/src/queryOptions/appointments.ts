import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface Appointment {
  id: string;
  patientId?: string;
  patientName: string;
  professionalId: string;
  startAt: string;
  endAt: string;
  status: 'scheduled' | 'confirmed' | 'completed' | 'cancelled' | 'no_show';
  kind?: string;
  notes?: string;
}

export const appointmentsByRangeOptions = (from: string, to: string) =>
  queryOptions({
    queryKey: queryKeys.appointments.byRange(from, to),
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams({ from, to });
      return api.get<Appointment[]>(`/api/v1/appointments?${qs.toString()}`, { signal });
    },
  });
