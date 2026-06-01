import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { AppointmentType } from "@/lib/api/calendar-api";

export type WaitlistStatus = "waiting" | "scheduled" | "cancelled";

export interface WaitlistEntry {
  id: string;
  patientId: string;
  doctorId?: string;
  preferredType?: AppointmentType;
  notes?: string;
  status: WaitlistStatus;
  scheduledAppointmentId?: string;
  createdByUserId: string;
  createdAt: string;
  updatedAt: string;
  patient?: { id: string; name: string; email?: string; phone?: string };
  doctor?: { id: string; name: string };
}

export interface CreateWaitlistInput {
  patientId: string;
  doctorId?: string;
  preferredType?: AppointmentType;
  notes?: string;
}

export const waitlistKeys = {
  all: ["waitlist"] as const,
  list: (status: string) => [...waitlistKeys.all, status] as const,
};

export function useWaitlist(status = "waiting") {
  return useQuery({
    queryKey: waitlistKeys.list(status),
    queryFn: () =>
      apiClient.get<WaitlistEntry[]>(`/api/v1/waitlist?status=${status}`),
    staleTime: 30_000,
    // Polling: encaixe novo aparece sozinho no cockpit (igual a agenda/leads).
    refetchInterval: 30_000,
  });
}

export function useCreateWaitlistEntry() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: CreateWaitlistInput) =>
      apiClient.post<WaitlistEntry>("/api/v1/waitlist", data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: waitlistKeys.all });
    },
  });
}

export function useUpdateWaitlistEntry(id: string) {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: {
      status?: WaitlistStatus;
      notes?: string;
      scheduledAppointmentId?: string;
    }) => apiClient.put<WaitlistEntry>(`/api/v1/waitlist/${id}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: waitlistKeys.all });
    },
  });
}

export function useDeleteWaitlistEntry() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/waitlist/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: waitlistKeys.all });
    },
  });
}
