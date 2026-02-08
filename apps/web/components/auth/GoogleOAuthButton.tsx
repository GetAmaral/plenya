"use client";

import { useEffect, useRef } from "react";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { useAuthStore } from "@/lib/auth-store";
import { loginWithGoogle } from "@/lib/api/auth";

const GOOGLE_CLIENT_ID =
  process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID ||
  "your_client_id.apps.googleusercontent.com";

declare global {
  interface Window {
    google?: any;
  }
}

export function GoogleOAuthButton() {
  const router = useRouter();
  const setAuth = useAuthStore((state) => state.setAuth);
  const buttonRef = useRef<HTMLDivElement>(null);
  const scriptLoaded = useRef(false);

  const handleCredentialResponse = async (response: any) => {
    try {
      const idToken = response.credential;

      if (!idToken) {
        toast.error("Erro ao fazer login", {
          description: "Token não recebido do Google",
        });
        return;
      }

      // Fazer login no backend
      const authResponse = await loginWithGoogle(idToken);

      // Salvar autenticação
      setAuth(authResponse.user, authResponse.accessToken, authResponse.refreshToken);

      toast.success("Login realizado com sucesso!", {
        description: `Bem-vindo, ${authResponse.user.email}`,
      });

      router.push("/dashboard");
    } catch (err) {
      toast.error("Erro ao fazer login com Google", {
        description: err instanceof Error ? err.message : "Erro desconhecido",
      });
    }
  };

  useEffect(() => {
    // Não carregar se já foi carregado ou se não tiver client ID
    if (scriptLoaded.current || !GOOGLE_CLIENT_ID || GOOGLE_CLIENT_ID === "your_client_id.apps.googleusercontent.com") {
      return;
    }

    // Carregar script do Google Identity Services
    const script = document.createElement("script");
    script.src = "https://accounts.google.com/gsi/client";
    script.async = true;
    script.defer = true;

    script.onload = () => {
      scriptLoaded.current = true;

      if (window.google && buttonRef.current) {
        // Inicializar Google Identity Services
        window.google.accounts.id.initialize({
          client_id: GOOGLE_CLIENT_ID,
          callback: handleCredentialResponse,
        });

        // Renderizar botão
        window.google.accounts.id.renderButton(
          buttonRef.current,
          {
            theme: "outline",
            size: "large",
            text: "signin_with",
            shape: "rectangular",
            width: buttonRef.current.offsetWidth,
          }
        );
      }
    };

    document.body.appendChild(script);

    return () => {
      // Cleanup não remove o script pois pode ser usado em outras partes
    };
  }, []);

  // Não renderizar se o client ID não estiver configurado
  if (!GOOGLE_CLIENT_ID || GOOGLE_CLIENT_ID === "your_client_id.apps.googleusercontent.com") {
    return null;
  }

  return <div ref={buttonRef} className="w-full" />;
}
