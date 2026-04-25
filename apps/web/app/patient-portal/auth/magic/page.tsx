"use client";

/**
 * /auth/magic?token= — consome um magic link e cria sessão.
 */
import { useEffect, useRef, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Loader2 } from "lucide-react";

import { useAuthStore } from "@/lib/auth-store";
import { patientAuthApi } from "@/lib/api/patient-portal-api";

export default function ConsumeMagicLinkPage() {
  const router = useRouter();
  const params = useSearchParams();
  const setAuth = useAuthStore((s) => s.setAuth);
  const [error, setError] = useState<string | null>(null);
  const ranRef = useRef(false);

  useEffect(() => {
    if (ranRef.current) return;
    ranRef.current = true;
    const token = params?.get("token");
    if (!token) {
      setError("Link inválido — token ausente.");
      return;
    }
    (async () => {
      try {
        const resp = await patientAuthApi.consumeMagicLink(token);
        setAuth(resp.user, resp.accessToken, resp.refreshToken);
        router.replace("/");
      } catch (err: any) {
        setError(err?.message ?? "Não foi possível validar o link");
      }
    })();
  }, [params, router, setAuth]);

  if (error) {
    return (
      <div className="mx-auto max-w-md space-y-4 py-10">
        <h1 className="text-3xl font-light">Link expirado</h1>
        <p className="text-muted-foreground">{error}</p>
        <a href="/login" className="inline-block underline-offset-4 hover:underline">
          Voltar pro login
        </a>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-md space-y-4 py-10">
      <p className="text-sm uppercase tracking-wider text-muted-foreground">
        Validando seu acesso
      </p>
      <h1 className="flex items-center gap-3 text-3xl font-light">
        <Loader2 className="h-6 w-6 animate-spin" /> Um instante…
      </h1>
    </div>
  );
}
