"use client";

import AppleSignin from "react-apple-signin-auth";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { Apple } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/lib/auth-store";
import { loginWithApple } from "@/lib/api/auth";

const APPLE_CLIENT_ID =
  process.env.NEXT_PUBLIC_APPLE_CLIENT_ID || "com.plenya.emr";
const APPLE_REDIRECT_URI =
  process.env.NEXT_PUBLIC_APPLE_REDIRECT_URI ||
  "http://localhost:3000/auth/apple/callback";

export function AppleOAuthButton() {
  const router = useRouter();
  const setAuth = useAuthStore((state) => state.setAuth);

  const handleResponse = async (response: any) => {
    try {
      const idToken = response.authorization.id_token;
      const user = response.user
        ? {
            name: {
              firstName: response.user.name.firstName,
              lastName: response.user.name.lastName,
            },
          }
        : undefined;

      // Fazer login no backend
      const authResponse = await loginWithApple(idToken, user);

      // Salvar autenticação
      setAuth(
        authResponse.user,
        authResponse.accessToken,
        authResponse.refreshToken
      );

      toast.success("Login realizado com sucesso!", {
        description: `Bem-vindo, ${authResponse.user.email}`,
      });

      router.push("/dashboard");
    } catch (err) {
      toast.error("Erro ao fazer login com Apple", {
        description: err instanceof Error ? err.message : "Erro desconhecido",
      });
    }
  };

  const handleError = (error: any) => {
    console.error("Apple Sign In error:", error);
    toast.error("Erro ao fazer login com Apple", {
      description: "Não foi possível autenticar com a Apple",
    });
  };

  return (
    <AppleSignin
      authOptions={{
        clientId: APPLE_CLIENT_ID,
        scope: "email name",
        redirectURI: APPLE_REDIRECT_URI,
        usePopup: true,
      }}
      onSuccess={handleResponse}
      onError={handleError}
      render={(props) => (
        <Button
          type="button"
          variant="outline"
          className="w-full h-11 bg-black text-white hover:bg-gray-900 hover:text-white border-black"
          onClick={props.onClick}
        >
          <Apple className="mr-2 h-5 w-5" />
          Entrar com Apple
        </Button>
      )}
    />
  );
}
