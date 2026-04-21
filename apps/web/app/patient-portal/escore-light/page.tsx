"use client";

import { useEffect, useState } from "react";
import { useSearchParams, useRouter } from "next/navigation";
import { useAuthStore } from "@/lib/auth-store";
import { userApi } from "@/lib/api/user-api";
import { scoreLightApi, type LightSession } from "@/lib/api/score-light-api";

type LoadState = "bootstrap" | "loading" | "ready" | "error";

export default function PatientPortalEscoreLightPage() {
  const params = useSearchParams();
  const router = useRouter();
  const { user, accessToken, setAuth } = useAuthStore();

  const [state, setState] = useState<LoadState>("bootstrap");
  const [error, setError] = useState<string | null>(null);
  const [sessions, setSessions] = useState<LightSession[]>([]);

  // Bootstrap: se vieram tokens via URL (vindo do magic link), persistir no auth store
  useEffect(() => {
    const urlAccess = params?.get("accessToken");
    const urlRefresh = params?.get("refreshToken");

    async function bootstrap() {
      // Se já temos sessão ativa, segue direto
      if (accessToken && user) {
        setState("loading");
        return;
      }

      // Se não temos sessão e nem tokens via URL, redireciona para login
      if (!urlAccess || !urlRefresh) {
        router.replace("/login");
        return;
      }

      try {
        // Persiste tokens primeiro (apiClient lê do store)
        setAuth(
          {
            id: "bootstrap-placeholder",
            email: "",
            roles: ["patient"],
            twoFactorEnabled: false,
            createdAt: new Date().toISOString(),
          },
          urlAccess,
          urlRefresh
        );

        // Limpa querystring
        if (typeof window !== "undefined") {
          const cleanUrl = window.location.pathname;
          window.history.replaceState({}, "", cleanUrl);
        }

        // Busca dados reais do usuário
        const me = await userApi.getMe();
        setAuth(me, urlAccess, urlRefresh);
        setState("loading");
      } catch (err) {
        setError(err instanceof Error ? err.message : "Falha ao recuperar sua conta");
        setState("error");
      }
    }
    void bootstrap();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Quando autenticado, carrega sessões
  useEffect(() => {
    if (state !== "loading") return;
    (async () => {
      try {
        const list = await scoreLightApi.mySessions();
        setSessions(list);
        setState("ready");
      } catch (err) {
        setError(err instanceof Error ? err.message : "Falha ao carregar sessões");
        setState("error");
      }
    })();
  }, [state]);

  if (state === "bootstrap" || state === "loading") {
    return (
      <div className="max-w-2xl">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Carregando seus dados
        </p>
        <h1 className="text-3xl font-light mt-3">Um instante...</h1>
      </div>
    );
  }

  if (state === "error") {
    return (
      <div className="max-w-2xl">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Erro
        </p>
        <h1 className="text-3xl font-light mt-3">Não foi possível carregar.</h1>
        <p className="text-muted-foreground mt-4">{error}</p>
      </div>
    );
  }

  return (
    <div className="max-w-3xl space-y-12">
      <header className="space-y-3">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Escore Plenya Light
        </p>
        <h1 className="text-3xl md:text-4xl font-light">
          Olá, {user?.name?.split(" ")[0] || user?.email || "paciente"}.
        </h1>
        <p className="text-muted-foreground leading-relaxed">
          Aqui ficam guardadas todas as autoavaliações que você realizou. Refazer a
          cada 3 meses ajuda a visualizar a evolução.
        </p>
      </header>

      {sessions.length === 0 ? (
        <div className="border rounded-md p-8 text-center space-y-4">
          <p className="text-muted-foreground">
            Nenhuma sessão salva ainda.
          </p>
          <a
            href={
              (process.env.NEXT_PUBLIC_SITE_URL ?? "http://localhost:3002") +
              "/pt/escore-plenya/avaliar"
            }
            className="inline-block px-6 py-3 bg-primary text-primary-foreground rounded-md hover:opacity-90 transition"
          >
            Fazer minha primeira avaliação
          </a>
        </div>
      ) : (
        <section>
          <h2 className="text-lg font-medium mb-6">Suas avaliações</h2>
          <div className="border-t">
            {sessions.map((s) => (
              <div
                key={s.publicCode}
                className="grid grid-cols-[1fr_auto_auto] gap-6 py-5 border-b items-center"
              >
                <div>
                  <p className="font-medium">
                    {new Date(s.createdAt).toLocaleDateString("pt-BR", {
                      day: "2-digit",
                      month: "long",
                      year: "numeric",
                    })}
                  </p>
                  <p className="text-sm text-muted-foreground mt-1">
                    {s.snapshot?.itemsEvaluatedCount ?? 0} itens avaliados
                  </p>
                </div>
                <div className="text-right">
                  {s.snapshot ? (
                    <>
                      <span className="text-3xl font-light tabular-nums">
                        {Math.round(s.snapshot.totalScorePercentage)}
                      </span>
                      <span className="text-muted-foreground">/100</span>
                    </>
                  ) : (
                    <span className="text-muted-foreground text-sm">—</span>
                  )}
                </div>
                <a
                  href={
                    (process.env.NEXT_PUBLIC_SITE_URL ?? "http://localhost:3002") +
                    "/pt/escore-plenya/resultado/" +
                    s.publicCode
                  }
                  className="text-sm underline-offset-4 hover:underline"
                >
                  ver radar →
                </a>
              </div>
            ))}
          </div>

          {sessions.length >= 1 && (
            <div className="mt-10 p-6 border rounded-md bg-muted/30">
              <p className="font-medium">Faça uma nova avaliação</p>
              <p className="text-sm text-muted-foreground mt-2 mb-4">
                Refazer o Escore Light em 3-6 meses permite acompanhar como cada pilar
                evoluiu.
              </p>
              <a
                href={
                  (process.env.NEXT_PUBLIC_SITE_URL ?? "http://localhost:3002") +
                  "/pt/escore-plenya/avaliar"
                }
                className="inline-block px-5 py-2.5 bg-primary text-primary-foreground rounded-md hover:opacity-90 transition text-sm"
              >
                Fazer nova avaliação
              </a>
            </div>
          )}
        </section>
      )}
    </div>
  );
}
