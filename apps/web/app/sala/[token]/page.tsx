"use client";

/**
 * /sala/[token] — lobby standalone de telemedicina (sem login).
 *
 * Fluxo:
 *  - Carrega LobbyView pública via API (sem auth)
 *  - Se janela ainda não abriu: countdown
 *  - Se janela aberta + dailyJoinUrl: botão "Entrar" → embed Daily inline
 *  - Se expirado/erro: mensagem clara
 */
import { use, useEffect, useState } from "react";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Loader2, Video, Clock, Calendar, AlertCircle } from "lucide-react";

import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { DailyCoEmbed } from "@/components/calendar/daily-co-embed";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

interface LobbyView {
  appointmentId: string;
  patientFirstName: string;
  doctorName: string;
  scheduledAt: string;
  durationMinutes: number;
  opensAt: string;
  closesAt: string;
  // HIGH H9 — URL completa com meeting_token (?t=...) escopado ao paciente.
  // Sala Daily.co é privacy=private; URL crua não funciona.
  dailyJoinUrl?: string | null;
  isOpen: boolean;
}

type LoadState =
  | { kind: "loading" }
  | { kind: "ok"; view: LobbyView }
  | { kind: "expired" }
  | { kind: "not_found" }
  | { kind: "error"; message: string };

export default function LobbyPage({ params }: { params: Promise<{ token: string }> }) {
  const { token } = use(params);
  const [state, setState] = useState<LoadState>({ kind: "loading" });
  const [entered, setEntered] = useState(false);
  const [now, setNow] = useState(() => new Date());

  // Polling pra que página atualize quando janela abrir
  useEffect(() => {
    const fetchView = async () => {
      try {
        const res = await fetch(`${API_URL}/api/v1/sala/${token}`);
        if (res.status === 404) {
          setState({ kind: "not_found" });
          return;
        }
        if (res.status === 410) {
          setState({ kind: "expired" });
          return;
        }
        if (!res.ok) {
          setState({ kind: "error", message: `${res.status}` });
          return;
        }
        const data: LobbyView = await res.json();
        setState({ kind: "ok", view: data });
      } catch (err: any) {
        setState({ kind: "error", message: err?.message ?? "falha de rede" });
      }
    };
    fetchView();
    const interval = setInterval(fetchView, 30_000);
    return () => clearInterval(interval);
  }, [token]);

  // Tick pra countdown
  useEffect(() => {
    const tick = setInterval(() => setNow(new Date()), 1000);
    return () => clearInterval(tick);
  }, []);

  const handleEnter = async () => {
    setEntered(true);
    // Best-effort: marca primeiro acesso
    fetch(`${API_URL}/api/v1/sala/${token}/used`, { method: "POST" }).catch(() => {});
  };

  if (state.kind === "loading") {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  if (state.kind === "not_found") {
    return (
      <div className="mx-auto max-w-md py-10 text-center">
        <AlertCircle className="mx-auto mb-3 h-8 w-8 text-muted-foreground" />
        <h1 className="text-2xl font-light">Link não encontrado</h1>
        <p className="mt-2 text-muted-foreground">
          Verifique se o link no email/WhatsApp está completo. Se persistir, fale com a equipe.
        </p>
      </div>
    );
  }

  if (state.kind === "expired") {
    return (
      <div className="mx-auto max-w-md py-10 text-center">
        <AlertCircle className="mx-auto mb-3 h-8 w-8 text-amber-600" />
        <h1 className="text-2xl font-light">Link expirado</h1>
        <p className="mt-2 text-muted-foreground">
          Este acesso à sala já expirou. Se ainda precisa entrar, fale com a equipe Plenya.
        </p>
      </div>
    );
  }

  if (state.kind === "error") {
    return (
      <div className="mx-auto max-w-md py-10 text-center">
        <AlertCircle className="mx-auto mb-3 h-8 w-8 text-destructive" />
        <h1 className="text-2xl font-light">Não foi possível carregar</h1>
        <p className="mt-2 text-muted-foreground">{state.message}</p>
      </div>
    );
  }

  const { view } = state;
  const when = new Date(view.scheduledAt);
  const opens = new Date(view.opensAt);
  const closes = new Date(view.closesAt);
  const beforeOpen = now < opens;
  const afterClose = now > closes;

  // Já entrou + tem URL → mostra Daily inline em fullscreen-ish
  if (entered && view.dailyJoinUrl) {
    return (
      <div className="mx-auto max-w-5xl">
        <DailyCoEmbed roomURL={view.dailyJoinUrl} onLeave={() => setEntered(false)} />
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-md space-y-6 py-6">
      <div className="space-y-1 text-center">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Teleconsulta</p>
        <h1 className="text-3xl font-light">Olá, {view.patientFirstName}</h1>
        <p className="text-muted-foreground">com {view.doctorName}</p>
      </div>

      <Card>
        <CardContent className="space-y-4 py-6 text-center">
          <div className="flex items-center justify-center gap-2 text-lg">
            <Calendar className="h-4 w-4 text-muted-foreground" />
            {format(when, "EEEE, d 'de' MMM 'às' HH:mm", { locale: ptBR })}
          </div>
          <p className="text-sm text-muted-foreground">{view.durationMinutes} minutos</p>

          {beforeOpen && (
            <div className="space-y-2 pt-2">
              <Clock className="mx-auto h-6 w-6 text-muted-foreground" />
              <p className="text-sm text-muted-foreground">
                A sala abre {format(opens, "'às' HH:mm")}.
              </p>
              <p className="text-xs text-muted-foreground">
                Pode deixar essa página aberta — o botão aparece automaticamente.
              </p>
            </div>
          )}

          {afterClose && (
            <p className="pt-2 text-sm text-muted-foreground">
              Esta sessão já encerrou. Fale com a equipe se precisar reagendar.
            </p>
          )}

          {view.isOpen && view.dailyJoinUrl && (
            <Button className="w-full" size="lg" onClick={handleEnter}>
              <Video className="mr-2 h-4 w-4" />
              Entrar na sala
            </Button>
          )}
          {view.isOpen && !view.dailyJoinUrl && (
            <p className="pt-2 text-sm text-muted-foreground">
              Sala temporariamente indisponível — atualize esta página em alguns segundos.
            </p>
          )}
        </CardContent>
      </Card>

      <p className="text-center text-xs text-muted-foreground">
        Permita acesso à câmera e ao microfone quando o navegador pedir.
      </p>
    </div>
  );
}
