import { useEffect, useState } from "react";
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
  const { user, accessToken } = useAuthStore();
  const [isHydrated, setIsHydrated] = useState(false);

  // Wait for zustand to hydrate from localStorage
  useEffect(() => {
    setIsHydrated(true);
  }, []);

  useEffect(() => {
    // Only redirect after hydration is complete
    if (isHydrated && (!user || !accessToken)) {
      router.push("/login");
    }
  }, [isHydrated, user, accessToken, router]);

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
