/**
 * Hook for accessing and managing the selected patient
 */

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { useAuthStore, Patient } from "@/lib/auth-store";
import { userApi } from "@/lib/api/user-api";

export function useSelectedPatient() {
  const { user, updateUser } = useAuthStore();
  const queryClient = useQueryClient();

  // Query to fetch current user (with selectedPatient)
  const { data: currentUser, isLoading } = useQuery({
    queryKey: ["user", "me"],
    queryFn: () => userApi.getMe(),
    enabled: !!user, // Only fetch if user is authenticated
    staleTime: 1000 * 60 * 5, // 5 minutes
    refetchOnMount: "always",
  });

  // Mutation to update selected patient
  const updateSelectedPatientMutation = useMutation({
    mutationFn: (patientId: string) => userApi.updateSelectedPatient(patientId),
    onSuccess: (updatedUser) => {
      // Update auth store
      updateUser(updatedUser);

      // Update query cache
      queryClient.setQueryData(["user", "me"], updatedUser);

      // Invalidate all patient-related queries
      queryClient.invalidateQueries({ queryKey: ["lab-results"] });
      queryClient.invalidateQueries({ queryKey: ["anamnesis"] });
      queryClient.invalidateQueries({ queryKey: ["prescriptions"] });
      queryClient.invalidateQueries({ queryKey: ["appointments"] });

      toast.success("Paciente selecionado", {
        description: `Agora trabalhando com: ${updatedUser.selectedPatient?.name}`,
      });
    },
    onError: (error: Error) => {
      toast.error("Erro ao selecionar paciente", {
        description: error.message,
      });
    },
  });

  // Use currentUser from query if available, fallback to store
  const effectiveUser = currentUser || user;
  const selectedPatient: Patient | null = effectiveUser?.selectedPatient || null;
  const selectedPatientId: string | null = effectiveUser?.selectedPatientId || null;

  return {
    selectedPatient,
    selectedPatientId,
    setSelectedPatient: (patientId: string) =>
      updateSelectedPatientMutation.mutate(patientId),
    isLoading: isLoading || updateSelectedPatientMutation.isPending,
    refreshSelectedPatient: () => queryClient.invalidateQueries({ queryKey: ["user", "me"] }),
  };
}
