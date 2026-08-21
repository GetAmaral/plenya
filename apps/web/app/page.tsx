"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Loader2 } from "lucide-react";
import { PlenyaMark } from "@/components/layout/plenya-mark";
import { apiClient } from "@/lib/api-client";
import { useAuthStore, waitForAuthHydration } from "@/lib/auth-store";
import { homeFor } from "@/lib/auth-routes";

/**
 * Raiz do EMR. É a `start_url` do PWA, ou seja: TODA abertura pelo ícone do celular passa aqui.
 *
 * Antes esta página redirecionava para `/login` sem olhar nada, então abrir o app pelo ícone era
 * sinônimo de digitar a senha de novo — mesmo com a sessão viva no aparelho e no servidor. Agora
 * ela espera a leitura do storage, renova a sessão (deslizando os 7 dias) e só manda para o login
 * quem realmente não tem sessão.
 */
export default function Home() {
  const router = useRouter();

  useEffect(() => {
    let cancelled = false;
    (async () => {
      await waitForAuthHydration();
      const alive = await apiClient.ensureFreshSession();
      if (cancelled) return;
      router.replace(alive ? homeFor(useAuthStore.getState().user) : "/login");
    })();
    return () => {
      cancelled = true;
    };
  }, [router]);

  return (
    <div className="flex min-h-screen items-center justify-center bg-linear-to-br from-cream via-paper to-sage-100 dark:from-petrol-800 dark:via-petrol dark:to-petrol-700">
      <div className="flex flex-col items-center gap-4">
        <PlenyaMark className="h-14 w-14" />
        <Loader2 className="h-5 w-5 animate-spin text-petrol/60 dark:text-cream/60" />
      </div>
    </div>
  );
}
