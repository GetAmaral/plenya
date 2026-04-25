"use client";

/**
 * /auth/invite?token= — primeira entrada via convite enviado pela equipe.
 *
 * Fluxo:
 *  1. Consome o token (cria sessão JWT)
 *  2. Mostra "Olá [nome], crie uma senha (opcional)"
 *  3. Paciente define senha OU pula (continua entrando só por magic link)
 *  4. Vai pra home
 */
import { useEffect, useRef, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { toast } from "sonner";
import { Loader2 } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useAuthStore } from "@/lib/auth-store";
import { patientAuthApi, patientMeApi } from "@/lib/api/patient-portal-api";

type Stage = "consuming" | "set-password" | "error";

export default function ConsumeInvitePage() {
  const router = useRouter();
  const params = useSearchParams();
  const setAuth = useAuthStore((s) => s.setAuth);
  const user = useAuthStore((s) => s.user);
  const [stage, setStage] = useState<Stage>("consuming");
  const [error, setError] = useState<string | null>(null);
  const [password, setPassword] = useState("");
  const [saving, setSaving] = useState(false);
  const ranRef = useRef(false);

  useEffect(() => {
    if (ranRef.current) return;
    ranRef.current = true;
    const token = params?.get("token");
    if (!token) {
      setStage("error");
      setError("Link inválido — token ausente.");
      return;
    }
    (async () => {
      try {
        const resp = await patientAuthApi.consumeInvite(token);
        setAuth(resp.user, resp.accessToken, resp.refreshToken);
        setStage("set-password");
      } catch (err: any) {
        setError(err?.message ?? "Convite inválido ou expirado");
        setStage("error");
      }
    })();
  }, [params, setAuth]);

  const handleSetPassword = async () => {
    if (password.length < 8) {
      toast.error("Senha deve ter ao menos 8 caracteres");
      return;
    }
    setSaving(true);
    try {
      await patientMeApi.setPassword(password);
      toast.success("Senha definida");
      router.replace("/");
    } catch (err: any) {
      toast.error(err?.message ?? "Falha ao definir senha");
    } finally {
      setSaving(false);
    }
  };

  if (stage === "error") {
    return (
      <div className="mx-auto max-w-md space-y-4 py-10">
        <h1 className="text-3xl font-light">Convite inválido</h1>
        <p className="text-muted-foreground">{error}</p>
        <p className="text-sm text-muted-foreground">
          Fale com a equipe Plenya para receber um novo convite.
        </p>
      </div>
    );
  }

  if (stage === "consuming") {
    return (
      <div className="mx-auto max-w-md space-y-4 py-10">
        <h1 className="flex items-center gap-3 text-3xl font-light">
          <Loader2 className="h-6 w-6 animate-spin" />
          Validando seu convite…
        </h1>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-md space-y-6 py-10">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Bem-vindo
        </p>
        <h1 className="text-3xl font-light">
          Olá, {user?.name?.split(" ")[0] ?? "paciente"}.
        </h1>
        <p className="text-muted-foreground">
          Sua área Plenya está pronta. Você pode definir uma senha agora ou pular —
          dá pra entrar sempre por link mágico no email.
        </p>
      </header>

      <div className="space-y-3">
        <Label htmlFor="pwd">Definir senha (opcional)</Label>
        <Input
          id="pwd"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="mínimo 8 caracteres"
          autoComplete="new-password"
        />
        <div className="flex gap-2 pt-2">
          <Button
            onClick={handleSetPassword}
            disabled={saving || password.length < 8}
          >
            {saving ? <Loader2 className="h-4 w-4 animate-spin" /> : "Salvar e entrar"}
          </Button>
          <Button variant="outline" onClick={() => router.replace("/")}>
            Pular
          </Button>
        </div>
      </div>
    </div>
  );
}
