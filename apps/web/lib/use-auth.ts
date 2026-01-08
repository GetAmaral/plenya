import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "./auth-store";

export function useRequireAuth() {
  const router = useRouter();
  const { user, accessToken } = useAuthStore();

  useEffect(() => {
    // Se não houver usuário ou token, redirecionar para login
    if (!user || !accessToken) {
      router.push("/login");
    }
  }, [user, accessToken, router]);

  return { user, accessToken, isAuthenticated: !!user && !!accessToken };
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
    logout,
  };
}
