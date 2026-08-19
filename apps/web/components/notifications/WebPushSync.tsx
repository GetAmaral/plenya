"use client";

import { useEffect } from "react";
import { apiClient } from "@/lib/api-client";
import { useAuthStore } from "@/lib/auth-store";

/**
 * Mantém a inscrição de avisos deste aparelho registrada no servidor.
 *
 * O navegador guarda a inscrição por conta própria e continua recebendo push mesmo com o
 * EMR fechado ou deslogado — mas quem manda o aviso é o servidor, e ele precisa saber que
 * ESTE endpoint pertence a ESTE usuário. Dois jeitos de essa ligação sumir sem ninguém
 * perceber: o servidor apagar a inscrição ao receber 404/410 de um push antigo, e o
 * usuário deslogar e entrar de novo (talvez em outra conta). Nos dois casos os avisos
 * param em silêncio, com o botão do sino ainda dizendo "ativado".
 *
 * Este componente só RE-REGISTRA o que o navegador já tem: não pede permissão, não cria
 * inscrição nova, não mostra nada na tela. Se o aparelho nunca ativou avisos, não faz nada.
 */
export function WebPushSync() {
  const accessToken = useAuthStore((state) => state.accessToken);

  useEffect(() => {
    if (!accessToken) return;
    if (typeof window === "undefined") return;
    if (!("serviceWorker" in navigator) || !("PushManager" in window)) return;
    if (typeof Notification === "undefined" || Notification.permission !== "granted") return;

    let cancelled = false;

    (async () => {
      try {
        const reg =
          (await navigator.serviceWorker.getRegistration("/sw.js")) ??
          (await navigator.serviceWorker.register("/sw.js"));
        const sub = await reg.pushManager.getSubscription();
        if (!sub || cancelled) return;

        await apiClient.post("/api/v1/web-push/subscribe", {
          ...sub.toJSON(),
          deviceLabel: deviceLabel(),
        });
      } catch {
        // Sem rede ou sem permissão: tenta de novo no próximo boot. Nada a avisar.
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [accessToken]);

  return null;
}

function deviceLabel(): string {
  if (typeof navigator === "undefined") return "";
  const ua = navigator.userAgent;
  const browser = /Edg/.test(ua)
    ? "Edge"
    : /Chrome/.test(ua)
      ? "Chrome"
      : /Firefox/.test(ua)
        ? "Firefox"
        : /Safari/.test(ua)
          ? "Safari"
          : "Navegador";
  // iPhone/iPad se declaram antes do "Mac" no UA — checar iOS primeiro evita rotular o
  // celular como macOS (era o que aparecia em "Aparelhos conectados").
  const os = /iPhone|iPad|iPod/.test(ua)
    ? "iOS"
    : /Android/.test(ua)
      ? "Android"
      : /Windows/.test(ua)
        ? "Windows"
        : /Mac/.test(ua)
          ? "macOS"
          : /Linux/.test(ua)
            ? "Linux"
            : "";
  return os ? `${browser} · ${os}` : browser;
}
