import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";

export type RecallStatus = "pending" | "scheduled" | "dismissed";

export interface Recall {
  id: string;
  patientId: string;
  doctorId?: string;
  sourceAppointmentId?: string;
  dueDate: string; // YYYY-MM-DD
  reason?: string;
  notes?: string;
  status: RecallStatus;
  scheduledAppointmentId?: string;
  createdByUserId: string;
  createdAt: string;
  updatedAt: string;
  patient?: { id: string; name: string; email?: string; phone?: string };
  doctor?: { id: string; name: string };
}

export interface CreateRecallInput {
  patientId: string;
  doctorId?: string;
  sourceAppointmentId?: string;
  dueDate: string; // YYYY-MM-DD
  reason?: string;
  notes?: string;
}

export const recallKeys = {
  all: ["recalls"] as const,
  list: (status: string, dueBefore?: string) =>
    [...recallKeys.all, status, dueBefore] as const,
};

export function useRecalls(status = "pending", dueBefore?: string) {
  return useQuery({
    queryKey: recallKeys.list(status, dueBefore),
    queryFn: () => {
      const p = new URLSearchParams();
      if (status) p.set("status", status);
      if (dueBefore) p.set("dueBefore", dueBefore);
      return apiClient.get<Recall[]>(`/api/v1/recalls?${p.toString()}`);
    },
    staleTime: 30_000,
    refetchInterval: 30_000,
  });
}

export function useCreateRecall() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: CreateRecallInput) =>
      apiClient.post<Recall>("/api/v1/recalls", data),
    onSuccess: () => qc.invalidateQueries({ queryKey: recallKeys.all }),
  });
}

export function useUpdateRecall(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (data: {
      status?: RecallStatus;
      dueDate?: string;
      reason?: string;
      notes?: string;
      scheduledAppointmentId?: string;
    }) => apiClient.put<Recall>(`/api/v1/recalls/${id}`, data),
    onSuccess: () => qc.invalidateQueries({ queryKey: recallKeys.all }),
  });
}

export function useDeleteRecall() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => apiClient.delete<void>(`/api/v1/recalls/${id}`),
    onSuccess: () => qc.invalidateQueries({ queryKey: recallKeys.all }),
  });
}
