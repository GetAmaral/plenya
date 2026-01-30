/**
 * Hook that ensures a patient is selected before accessing patient-related pages
 * Redirects to /patients/select if no patient is selected
 */

import { useEffect } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useSelectedPatient } from "./use-selected-patient";

export function useRequireSelectedPatient() {
  const router = useRouter();
  const pathname = usePathname();
  const { selectedPatient, isLoading } = useSelectedPatient();

  useEffect(() => {
    // Don't redirect while loading
    if (isLoading) return;

    // If no selected patient, save current path and redirect to selection
    if (!selectedPatient) {
      // Save the intended destination in sessionStorage
      if (typeof window !== "undefined") {
        sessionStorage.setItem("plenya-redirect-after-patient-select", pathname);
      }

      router.push("/patients/select");
    }
  }, [selectedPatient, isLoading, router, pathname]);

  return {
    selectedPatient,
    isLoading,
    hasSelectedPatient: !!selectedPatient,
  };
}
