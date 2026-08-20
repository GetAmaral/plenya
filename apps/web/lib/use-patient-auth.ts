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
  // Espera a leitura do localStorage antes de concluir que não há sessão. Sem isto, o primeiro
  // render (store ainda vazia) mandava para o login em todo boot frio do PWA.
  const hasHydrated = useAuthStore((s) => s.hasHydrated);

  useEffect(() => {
    if (!hasHydrated) return;
    if (!accessToken || !user) {
      router.replace("/login");
      return;
    }
    if (!isGranted(user, "patient")) {
      router.replace("/login");
      return;
    }
  }, [hasHydrated, accessToken, user, router]);

  return { user, accessToken, ready: !!user && isGranted(user, "patient") };
}
