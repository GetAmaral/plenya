"use client";

/**
 * /login — entrada do portal (minha.plenyasaude.com.br/login)
 * Híbrido: email + senha OU magic link, mesmo formulário.
 */
import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { toast } from "sonner";
import { Loader2 } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuthStore } from "@/lib/auth-store";
import { patientAuthApi } from "@/lib/api/patient-portal-api";

export default function PatientLoginPage() {
  const router = useRouter();
  const setAuth = useAuthStore((s) => s.setAuth);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [magicSent, setMagicSent] = useState(false);

  const handlePasswordLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email || !password) return;
    setLoading(true);
    try {
      const resp = await patientAuthApi.loginPassword(email, password);
      setAuth(resp.user, resp.accessToken, resp.refreshToken);
      toast.success("Bem-vindo!");
      router.replace("/");
    } catch (err: any) {
      toast.error(err?.message ?? "Falha ao entrar");
    } finally {
      setLoading(false);
    }
  };

  const handleMagicLink = async () => {
    if (!email) {
      toast.error("Informe seu email primeiro");
      return;
    }
    setLoading(true);
    try {
      await patientAuthApi.requestMagicLink(email);
      setMagicSent(true);
    } catch (err: any) {
      toast.error(err?.message ?? "Falha ao enviar link");
    } finally {
      setLoading(false);
    }
  };

  if (magicSent) {
    return (
      <div className="mx-auto max-w-md space-y-6 py-10">
        <h1 className="text-3xl font-light">Verifique seu email</h1>
        <p className="text-muted-foreground">
          Se houver uma conta vinculada a <strong>{email}</strong>, enviamos um link
          de acesso. O link expira em 15 minutos.
        </p>
        <Button variant="outline" onClick={() => setMagicSent(false)}>
          Voltar
        </Button>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-md space-y-8 py-10">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Acesso
        </p>
        <h1 className="text-3xl font-light">Entrar na sua área Plenya</h1>
        <p className="text-muted-foreground">
          Use email e senha — ou peça um link mágico se preferir.
        </p>
      </header>

      <form onSubmit={handlePasswordLogin} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            autoComplete="email"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="password">Senha</Label>
          <Input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            autoComplete="current-password"
          />
        </div>

        <div className="flex flex-col gap-2 pt-2">
          <Button type="submit" disabled={loading || !email || !password}>
            {loading ? <Loader2 className="h-4 w-4 animate-spin" /> : "Entrar"}
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={handleMagicLink}
            disabled={loading}
          >
            Receber link mágico no email
          </Button>
        </div>

        <div className="pt-3 text-center text-sm text-muted-foreground">
          <Link href="/esqueci-senha" className="underline-offset-4 hover:underline">
            Esqueci minha senha
          </Link>
        </div>
      </form>
    </div>
  );
}
