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
  const { user, accessToken, refreshToken } = useAuthStore();
  // Espera a leitura do localStorage antes de concluir que não há sessão. Sem isto, o primeiro
  // render (store ainda vazia) mandava para o login em todo boot frio do PWA.
  const hasHydrated = useAuthStore((s) => s.hasHydrated);

  useEffect(() => {
    if (!hasHydrated) return;
    // Sessão é o refresh token: o access dura 30min e está sempre vencido ao reabrir o app.
    if ((!accessToken && !refreshToken) || !user) {
      router.replace("/login");
      return;
    }
    if (!isGranted(user, "patient")) {
      router.replace("/login");
      return;
    }
  }, [hasHydrated, accessToken, refreshToken, user, router]);

  return { user, accessToken, ready: !!user && isGranted(user, "patient") };
}
