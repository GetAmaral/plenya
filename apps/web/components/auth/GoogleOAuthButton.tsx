"use client";

import { GoogleOAuthProvider, GoogleLogin } from "@react-oauth/google";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { useAuthStore } from "@/lib/auth-store";
import { loginWithGoogle } from "@/lib/api/auth";

const GOOGLE_CLIENT_ID =
  process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID ||
  "your_client_id.apps.googleusercontent.com";

export function GoogleOAuthButton() {
  const router = useRouter();
  const setAuth = useAuthStore((state) => state.setAuth);

  const handleSuccess = async (credentialResponse: any) => {
    try {
      const idToken = credentialResponse.credential;

      if (!idToken) {
        toast.error("Erro ao fazer login", {
          description: "Token não recebido do Google",
        });
        return;
      }

      // Fazer login no backend
      const response = await loginWithGoogle(idToken);

      // Salvar autenticação
      setAuth(response.user, response.accessToken, response.refreshToken);

      toast.success("Login realizado com sucesso!", {
        description: `Bem-vindo, ${response.user.email}`,
      });

      router.push("/dashboard");
    } catch (err) {
      toast.error("Erro ao fazer login com Google", {
        description: err instanceof Error ? err.message : "Erro desconhecido",
      });
    }
  };

  const handleError = () => {
    toast.error("Erro ao fazer login com Google", {
      description: "Não foi possível autenticar com o Google",
    });
  };

  // Não renderizar se o client ID não estiver configurado
  if (!GOOGLE_CLIENT_ID || GOOGLE_CLIENT_ID === "your_client_id.apps.googleusercontent.com") {
    return null;
  }

  return (
    <GoogleOAuthProvider clientId={GOOGLE_CLIENT_ID}>
      <div className="w-full">
        <GoogleLogin
          onSuccess={handleSuccess}
          onError={handleError}
          theme="outline"
          size="large"
          text="signin_with"
          shape="rectangular"
        />
      </div>
    </GoogleOAuthProvider>
  );
}
