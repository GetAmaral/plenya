"use client";

/**
 * /esqueci-senha — atalho UX pro fluxo de magic link.
 * Tecnicamente reusa /api/v1/auth/patient/forgot (alias de magic-link).
 */
import { useState } from "react";
import Link from "next/link";
import { Loader2 } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { patientAuthApi } from "@/lib/api/patient-portal-api";

export default function ForgotPasswordPage() {
  const [email, setEmail] = useState("");
  const [loading, setLoading] = useState(false);
  const [sent, setSent] = useState(false);

  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      await patientAuthApi.forgotPassword(email);
      setSent(true);
    } finally {
      setLoading(false);
    }
  };

  if (sent) {
    return (
      <div className="mx-auto max-w-md space-y-4 py-10">
        <h1 className="text-3xl font-light">Verifique seu email</h1>
        <p className="text-muted-foreground">
          Se houver uma conta vinculada a <strong>{email}</strong>, enviamos um link
          de acesso. O link expira em 15 minutos. Use-o pra entrar e definir uma nova senha.
        </p>
        <Link href="/login" className="inline-block underline-offset-4 hover:underline">
          Voltar
        </Link>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-md space-y-6 py-10">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Recuperação
        </p>
        <h1 className="text-3xl font-light">Esqueci minha senha</h1>
        <p className="text-muted-foreground">
          Vamos enviar um link de acesso para o seu email. Lá dentro você pode
          redefinir sua senha (ou seguir entrando só com link mágico).
        </p>
      </header>

      <form onSubmit={onSubmit} className="space-y-4">
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
        <Button type="submit" disabled={loading || !email}>
          {loading ? <Loader2 className="h-4 w-4 animate-spin" /> : "Enviar link"}
        </Button>
      </form>
    </div>
  );
}
