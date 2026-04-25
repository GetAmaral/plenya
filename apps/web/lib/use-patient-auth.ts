"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore, isGranted } from "@/lib/auth-store";

/**
 * useRequirePatientAuth — redireciona pra /login se o usuário não está logado
 * ou não tem role=patient. Use no topo de cada página autenticada do portal.
 *
 * Reusa useAuthStore (mesmo localStorage do staff) — o backend valida role.
 */
export function useRequirePatientAuth() {
  const router = useRouter();
  const { user, accessToken } = useAuthStore();

  useEffect(() => {
    if (!accessToken || !user) {
      router.replace("/login");
      return;
    }
    if (!isGranted(user, "patient")) {
      router.replace("/login");
      return;
    }
  }, [accessToken, user, router]);

  return { user, accessToken, ready: !!user && isGranted(user, "patient") };
}
