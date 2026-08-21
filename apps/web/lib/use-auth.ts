import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "./auth-store";

const DEV_BYPASS_AUTH = process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH === 'true';

export function useRequireAuth() {
  // DEV BYPASS: pula toda verificação
  if (DEV_BYPASS_AUTH) {
    const { user, accessToken } = useAuthStore();
    return { user, accessToken, isAuthenticated: true };
  }

  const router = useRouter();
  const { user, accessToken, refreshToken } = useAuthStore();
  // hasHydrated vem da própria store: é o sinal de que o localStorage FOI LIDO. O efeito de
  // montagem que existia aqui só dizia "estamos no cliente", que é outra coisa — no boot lento do
  // PWA ele marcava "hidratado" antes da leitura e mandava para o login com sessão válida guardada.
  const hasHydrated = useAuthStore((s) => s.hasHydrated);

  // Ter sessão é ter o refresh token — o access token dura 30min e VAI estar vencido sempre que
  // o app é reaberto depois de um tempo. Exigir access aqui mandava para o login quem estava
  // perfeitamente logado, bastando o aparelho ter ficado algumas horas fechado.
  const hasSession = !!user && (!!accessToken || !!refreshToken);

  useEffect(() => {
    if (hasHydrated && !hasSession) {
      router.push("/login");
    }
  }, [hasHydrated, hasSession, router]);

  return { user, accessToken, isAuthenticated: hasSession };
}

export function useAuth() {
  const { user, accessToken, clearAuth } = useAuthStore();

  const logout = () => {
    clearAuth();
  };

  return {
    user,
    accessToken,
    isAuthenticated: !!user && !!accessToken,
    isLoading: false,
    logout,
  };
}
