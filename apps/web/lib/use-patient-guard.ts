"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "./use-auth";
import { toast } from "sonner";

/**
 * Redireciona patients para /dashboard se tentarem acessar páginas staff-only
 *
 * Use este hook em páginas que devem ser restritas a profissionais de saúde.
 * Patients (usuários com apenas o role "patient") serão redirecionados automaticamente
 * para o dashboard com uma mensagem de erro.
 *
 * @example
 * ```tsx
 * export default function StaffOnlyPage() {
 *   usePatientGuard(); // Adiciona proteção contra acesso de patients
 *
 *   // Resto da página...
 * }
 * ```
 */
export function usePatientGuard() {
  const { user } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (!user) return;

    // Check if user is patient-only (only has patient role, no other roles)
    const isPatientOnly =
      user.roles.includes("patient") && user.roles.length === 1;

    if (isPatientOnly) {
      toast.error("Acesso negado", {
        description: "Esta página é restrita a profissionais de saúde",
      });
      router.replace("/dashboard");
    }
  }, [user, router]);
}
