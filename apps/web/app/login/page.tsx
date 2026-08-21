"use client";

import { useState, useRef, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent } from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { touchActivity } from "@/components/auth/inactivity-lock";
import { apiClient } from "@/lib/api-client";
import { useAuthStore, waitForAuthHydration, type UserRole } from "@/lib/auth-store";
import { homeFor } from "@/lib/auth-routes";
import { useFormNavigation } from "@/lib/use-form-navigation";
import { ArrowRight, Loader2 } from "lucide-react";
import { PlenyaMark } from "@/components/layout/plenya-mark";
import { GoogleOAuthButton } from "@/components/auth/GoogleOAuthButton";

const loginSchema = z.object({
  email: z.string().email("Email inválido"),
  password: z.string().min(6, "Senha deve ter no mínimo 6 caracteres"),
});

type LoginForm = z.infer<typeof loginSchema>;

interface LoginResponse {
  accessToken: string;
  refreshToken: string;
  user: {
    id: string;
    email: string;
    roles: UserRole[];
    twoFactorEnabled: boolean;
    createdAt: string;
  };
}

export default function LoginPage() {
  const router = useRouter();
  const setAuth = useAuthStore((state) => state.setAuth);
  /**
   * Enquanto true, esta tela ainda não sabe se já existe sessão neste aparelho — e por isso não
   * mostra o formulário.
   *
   * Era ESTE o "logout" do iPhone: o PWA abre sempre na start_url `/`, que mandava direto para
   * cá, e o login não olhava se já havia sessão. A sessão continuava viva no servidor (os
   * tokens antigos estão lá, válidos e nunca usados) — o app é que pedia a senha de novo a cada
   * abertura, e digitar a senha abria mais uma sessão. Ver docs/emr/estudo-sessao-login-persistente.md.
   */
  const [checkingSession, setCheckingSession] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  // "Manter conectado": sessão longa deslizante (30 dias) neste aparelho. Marcado por padrão
  // pra resolver o logout rápido no PWA; o aviso pede pra desmarcar em aparelho compartilhado.
  const [remember, setRemember] = useState(true);
  const formRef = useRef<HTMLFormElement>(null);

  // Navegação por Enter nos campos do formulário
  useFormNavigation({ formRef });

  // Já logado? Renova a sessão (desliza os 7 dias) e entra direto, sem pedir senha.
  useEffect(() => {
    let cancelled = false;
    (async () => {
      await waitForAuthHydration();
      const { user, refreshToken } = useAuthStore.getState();
      if (!user || !refreshToken) {
        if (!cancelled) setCheckingSession(false);
        return;
      }
      const alive = await apiClient.ensureFreshSession();
      if (cancelled) return;
      if (alive) {
        touchActivity();
        router.replace(homeFor(useAuthStore.getState().user));
        return;
      }
      setCheckingSession(false);
    })();
    return () => {
      cancelled = true;
    };
  }, [router]);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginForm>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginForm) => {
    try {
      setIsLoading(true);

      const response = await apiClient.post<LoginResponse>(
        "/api/v1/auth/login",
        { ...data, rememberDevice: remember }
      );

      setAuth(response.user, response.accessToken, response.refreshToken);
      touchActivity(); // zera o relógio de inatividade pra não travar logo após o login

      toast.success("Login realizado com sucesso!", {
        description: `Bem-vindo de volta, ${response.user.email}`,
      });

      // Secretaria (sem papel clinico) cai direto na Recepcao. Demais papeis
      // mantem o destino padrao /dashboard.
      router.push(homeFor(response.user));
    } catch (err) {
      toast.error("Erro ao fazer login", {
        description: err instanceof Error ? err.message : "Credenciais inválidas",
      });
    } finally {
      setIsLoading(false);
    }
  };

  if (checkingSession) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-linear-to-br from-cream via-paper to-sage-100 dark:from-petrol-800 dark:via-petrol dark:to-petrol-700">
        <div className="flex flex-col items-center gap-4">
          <PlenyaMark className="h-14 w-14" />
          <Loader2 className="h-5 w-5 animate-spin text-petrol/60 dark:text-cream/60" />
        </div>
      </div>
    );
  }

  return (
    <div className="relative flex min-h-screen items-center justify-center overflow-hidden bg-linear-to-br from-cream via-paper to-sage-100 dark:from-petrol-800 dark:via-petrol dark:to-petrol-700">
      {/* Animated background gradient orbs */}
      <div className="absolute inset-0 overflow-hidden">
        <motion.div
          className="absolute -top-40 -right-40 h-80 w-80 rounded-full bg-linear-to-br from-gold-300 to-gold-500 opacity-20 blur-3xl"
          animate={{
            scale: [1, 1.2, 1],
            rotate: [0, 90, 0],
          }}
          transition={{
            duration: 20,
            repeat: Infinity,
            ease: "linear",
          }}
        />
        <motion.div
          className="absolute -bottom-40 -left-40 h-80 w-80 rounded-full bg-linear-to-br from-ocean-300 to-petrol opacity-20 blur-3xl"
          animate={{
            scale: [1.2, 1, 1.2],
            rotate: [90, 0, 90],
          }}
          transition={{
            duration: 25,
            repeat: Infinity,
            ease: "linear",
          }}
        />
      </div>

      {/* Login card */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="relative z-10 w-full max-w-md px-4"
      >
        <Card className="border-0 shadow-2xl backdrop-blur-xl bg-white/80 dark:bg-gray-900/80">
          <CardContent className="p-8">
            {/* Logo and title */}
            <div className="mb-8 text-center">
              <motion.div
                initial={{ scale: 0 }}
                animate={{ scale: 1 }}
                transition={{ delay: 0.2, type: "spring", stiffness: 200 }}
                className="mb-4 flex justify-center"
              >
                <PlenyaMark className="h-16 w-16 shadow-lg" />
              </motion.div>

              <motion.h1
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.3 }}
                className="font-heading text-4xl tracking-tight text-petrol dark:text-cream"
              >
                Plenya
              </motion.h1>

              <motion.p
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.4 }}
                className="label-upper mt-2 text-muted-foreground"
              >
                Prontuário Eletrônico
              </motion.p>
            </div>

            {/* OAuth Login Buttons */}
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.5 }}
              className="space-y-3 mb-6"
            >
              <GoogleOAuthButton />

              {/* Divider */}
              <div className="relative my-6">
                <div className="absolute inset-0 flex items-center">
                  <span className="w-full border-t border-border" />
                </div>
                <div className="relative flex justify-center text-xs uppercase">
                  <span className="bg-card px-2 text-muted-foreground">
                    Ou continue com email
                  </span>
                </div>
              </div>
            </motion.div>

            {/* Login form */}
            <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-5">
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.6 }}
                className="space-y-2"
              >
                <Label htmlFor="email" className="text-gray-700 dark:text-gray-300">
                  Email
                </Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="seu@email.com"
                  className="h-11"
                  {...register("email")}
                />
                {errors.email && (
                  <p className="text-sm text-destructive">
                    {errors.email.message}
                  </p>
                )}
              </motion.div>

              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 0.7 }}
                className="space-y-2"
              >
                <Label htmlFor="password" className="text-gray-700 dark:text-gray-300">
                  Senha
                </Label>
                <Input
                  id="password"
                  type="password"
                  placeholder="••••••••"
                  className="h-11"
                  {...register("password")}
                />
                {errors.password && (
                  <p className="text-sm text-destructive">
                    {errors.password.message}
                  </p>
                )}
              </motion.div>

              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.7 }}
                className="flex items-start gap-2"
              >
                <Checkbox
                  id="remember"
                  checked={remember}
                  onCheckedChange={(v) => setRemember(v === true)}
                  className="mt-0.5"
                />
                <Label htmlFor="remember" className="text-sm font-normal leading-snug text-gray-600 dark:text-gray-400 cursor-pointer">
                  Manter conectado neste aparelho
                  <span className="block text-xs text-gray-400">
                    Use só em aparelho pessoal. Em computador compartilhado, desmarque.
                  </span>
                </Label>
              </motion.div>

              <motion.div
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.7 }}
              >
                <Button
                  type="submit"
                  className="group relative h-11 w-full overflow-hidden bg-linear-to-r from-gold-500 to-gold-600 text-petrol shadow-lg transition-all hover:shadow-xl hover:scale-[1.02]"
                  disabled={isLoading}
                >
                  {isLoading ? (
                    <>
                      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                      Entrando...
                    </>
                  ) : (
                    <>
                      Entrar
                      <ArrowRight className="ml-2 h-4 w-4 transition-transform group-hover:translate-x-1" />
                    </>
                  )}
                </Button>
              </motion.div>
            </form>

            {/* Footer */}
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 0.8 }}
              className="mt-6 text-center"
            >
              <p className="text-xs text-gray-500 dark:text-gray-400">
                Acesso restrito a profissionais autorizados
              </p>
            </motion.div>
          </CardContent>
        </Card>

        {/* Version badge */}
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 1 }}
          className="mt-4 text-center text-xs text-gray-500 dark:text-gray-400"
        >
          v1.0.0 - Janeiro 2026
        </motion.div>
      </motion.div>
    </div>
  );
}
