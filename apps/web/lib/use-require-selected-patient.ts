/**
 * Hook that ensures a patient is selected before accessing patient-related pages
 * Redirects to /patients/select if no patient is selected
 */

import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useSelectedPatient } from "./use-selected-patient";
import { useAuthStore } from "./auth-store";

export function useRequireSelectedPatient() {
  const router = useRouter();
  const pathname = usePathname();
  const { selectedPatient, isLoading } = useSelectedPatient();
  const { user } = useAuthStore();
  const [hasMounted, setHasMounted] = useState(false);

  // Wait for component to mount before checking (prevents hydration issues)
  useEffect(() => {
    setHasMounted(true);
  }, []);

  useEffect(() => {
    // Don't redirect if:
    // 1. Component hasn't mounted yet (hydration)
    // 2. Still loading mutation
    // 3. User is not authenticated yet
    if (!hasMounted || isLoading || !user) {
      return;
    }

    // Check store directly (it's synchronous and always up-to-date)
    const hasSelectedPatient = !!user.selectedPatient || !!user.selectedPatientId;

    // Only redirect if we're absolutely sure there's no selected patient
    if (!hasSelectedPatient) {
      // Save the intended destination in sessionStorage
      if (typeof window !== "undefined") {
        sessionStorage.setItem("plenya-redirect-after-patient-select", pathname);
      }

      router.push("/patients/select");
    }
  }, [hasMounted, selectedPatient, isLoading, user, router, pathname]);

  return {
    selectedPatient,
    isLoading: !hasMounted || isLoading,
    hasSelectedPatient: !!selectedPatient,
  };
}
