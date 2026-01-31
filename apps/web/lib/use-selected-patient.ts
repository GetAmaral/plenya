/**
 * Hook for accessing and managing the selected patient
 */

import { useEffect } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { useAuthStore, Patient } from "@/lib/auth-store";
import { userApi } from "@/lib/api/user-api";

export function useSelectedPatient() {
  const { user, updateUser } = useAuthStore();
  const queryClient = useQueryClient();

  // Query to fetch current user (with selectedPatient) in background only
  const { data: currentUser } = useQuery({
    queryKey: ["user", "me"],
    queryFn: () => userApi.getMe(),
    enabled: !!user, // Only fetch if user is authenticated
    staleTime: 1000 * 60 * 5, // 5 minutes - cache for 5 min
    gcTime: 1000 * 60 * 30, // Keep in cache for 30 min
    refetchOnMount: false, // Never refetch on mount
    refetchOnWindowFocus: false, // Never refetch on focus
    refetchOnReconnect: false, // Never refetch on reconnect
  });

  // Sync currentUser from query to auth store when it loads (background only)
  useEffect(() => {
    if (currentUser) {
      // Only update store if there's actually new data from server
      if (JSON.stringify(user) !== JSON.stringify(currentUser)) {
        updateUser(currentUser);
      }
    }
  }, [currentUser, user, updateUser]);

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
      queryClient.invalidateQueries({ queryKey: ["lab-requests"] });

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

  // ALWAYS use store data as primary source (from localStorage)
  // Query only syncs in background, never blocks or clears data
  const selectedPatient: Patient | null = user?.selectedPatient || null;
  const selectedPatientId: string | null = user?.selectedPatientId || null;

  return {
    selectedPatient,
    selectedPatientId,
    setSelectedPatient: (patientId: string) =>
      updateSelectedPatientMutation.mutate(patientId),
    isLoading: updateSelectedPatientMutation.isPending,
    refreshSelectedPatient: () => queryClient.invalidateQueries({ queryKey: ["user", "me"] }),
  };
}
