import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface StaffUser {
  id: string;
  name: string;
  email: string;
  role: string;
  avatarUrl?: string;
}

export const staffListOptions = () =>
  queryOptions({
    queryKey: queryKeys.users.staff(),
    queryFn: ({ signal }) =>
      api.get<StaffUser[]>('/api/v1/users/staff', { signal }),
    staleTime: 5 * 60_000,
  });

export const doctorsListOptions = () =>
  queryOptions({
    queryKey: queryKeys.users.doctors(),
    queryFn: ({ signal }) =>
      api.get<StaffUser[]>('/api/v1/users/doctors', { signal }),
    staleTime: 5 * 60_000,
  });
