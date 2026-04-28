"use client";

/**
 * /consultas/[id] — detalhe da consulta. Se telemed dentro da janela,
 * mostra o Daily.co embed inline. Botões Confirmar / Solicitar reagendamento /
 * Solicitar cancelamento.
 */
import { use, useState } from "react";
import Link from "next/link";
import { format, formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import { toast } from "sonner";
import {
  Loader2,
  Calendar,
  Video,
  CheckCircle2,
  ArrowLeft,
  CalendarX,
  CalendarClock,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyAppointment,
  useConfirmAppointment,
  useRequestAppointmentAction,
} from "@/lib/api/patient-portal-api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Textarea } from "@/components/ui/textarea";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { DailyCoEmbed } from "@/components/calendar/daily-co-embed";

const TYPE_LABEL: Record<string, string> = {
  initial_assessment: "Avaliação inicial",
  follow_up: "Retorno",
  telemedicine: "Teleconsulta",
  procedure: "Procedimento",
  results_review: "Revisão de exames",
};

export default function MyAppointmentDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const { ready } = useRequirePatientAuth();
  const { data: a, isLoading } = useMyAppointment(ready ? id : undefined);
  const confirm = useConfirmAppointment();
  const requestAction = useRequestAppointmentAction();
  const [actionDialog, setActionDialog] = useState<"cancel" | "reschedule" | null>(null);
  const [reason, setReason] = useState("");

  const handleConfirm = async () => {
    try {
      await confirm.mutateAsync(id);
      toast.success("Consulta confirmada");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao confirmar");
    }
  };

  const handleSendAction = async () => {
    if (!actionDialog) return;
    try {
      await requestAction.mutateAsync({ id, action: actionDialog, reason });
      toast.success(
        actionDialog === "cancel"
          ? "Solicitação de cancelamento enviada"
          : "Solicitação de reagendamento enviada",
      );
      setActionDialog(null);
      setReason("");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao enviar solicitação");
    }
  };

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }
  if (!a) {
    return (
      <div className="mx-auto max-w-2xl py-10">
        <Link href="/consultas" className="text-sm underline-offset-4 hover:underline">
          ← Voltar
        </Link>
        <p className="mt-6 text-muted-foreground">Consulta não encontrada.</p>
      </div>
    );
  }

  const when = new Date(a.scheduledAt);
  const inPast = a.status === "completed" || a.status === "cancelled" || a.status === "no_show";

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <Link href="/consultas" className="inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground">
        <ArrowLeft className="h-3.5 w-3.5" /> Minhas consultas
      </Link>

      <header className="space-y-3">
        <div className="flex flex-wrap items-center gap-2">
          <Badge variant="secondary">{TYPE_LABEL[a.type] ?? a.type}</Badge>
          {a.isTelemedicine && (
            <Badge variant="outline" className="gap-1">
              <Video className="h-3 w-3" /> Online
            </Badge>
          )}
          {a.patientConfirmedAt && (
            <Badge variant="default" className="gap-1">
              <CheckCircle2 className="h-3 w-3" /> Confirmada por você
            </Badge>
          )}
        </div>
        <h1 className="text-3xl font-light">
          {format(when, "EEEE, d 'de' MMMM", { locale: ptBR })}
        </h1>
        <p className="flex items-center gap-2 text-xl text-muted-foreground">
          <Calendar className="h-4 w-4" />
          {format(when, "HH:mm")} · {a.durationMinutes} min · com {a.doctorName}
        </p>
        {!inPast && (
          <p className="text-sm text-muted-foreground">
            {formatDistanceToNow(when, { addSuffix: true, locale: ptBR })}
          </p>
        )}
      </header>

      {/* Telemed inline */}
      {a.isTelemedicine && a.dailyJoinUrl && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-base">
              <Video className="h-4 w-4" /> Sala da consulta
            </CardTitle>
          </CardHeader>
          <CardContent>
            <DailyCoEmbed roomURL={a.dailyJoinUrl} />
          </CardContent>
        </Card>
      )}
      {a.isTelemedicine && !a.dailyJoinUrl && !inPast && (
        <Card className="border-dashed">
          <CardContent className="space-y-2 py-6 text-center">
            <Video className="mx-auto h-6 w-6 text-muted-foreground" />
            <p className="text-sm text-muted-foreground">
              A sala abre 30 minutos antes do horário agendado.
            </p>
          </CardContent>
        </Card>
      )}

      {/* Notas */}
      {a.notes && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Observações</CardTitle>
          </CardHeader>
          <CardContent className="text-sm text-muted-foreground whitespace-pre-wrap">
            {a.notes}
          </CardContent>
        </Card>
      )}

      {/* Ações */}
      {!inPast && (
        <div className="flex flex-wrap gap-2">
          {!a.patientConfirmedAt && (
            <Button onClick={handleConfirm} disabled={confirm.isPending}>
              {confirm.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : <CheckCircle2 className="mr-1 h-4 w-4" />}
              Confirmar presença
            </Button>
          )}
          <Button variant="outline" onClick={() => setActionDialog("reschedule")}>
            <CalendarClock className="mr-1 h-4 w-4" /> Solicitar reagendamento
          </Button>
          <Button variant="outline" onClick={() => setActionDialog("cancel")}>
            <CalendarX className="mr-1 h-4 w-4" /> Solicitar cancelamento
          </Button>
        </div>
      )}

      <Dialog open={!!actionDialog} onOpenChange={(v) => !v && setActionDialog(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>
              {actionDialog === "cancel"
                ? "Solicitar cancelamento"
                : "Solicitar reagendamento"}
            </DialogTitle>
            <DialogDescription>
              A equipe recebe sua solicitação e entra em contato pra confirmar. Conta pra gente o motivo (opcional).
            </DialogDescription>
          </DialogHeader>
          <Textarea
            value={reason}
            onChange={(e) => setReason(e.target.value)}
            placeholder="Ex: imprevisto pessoal, prefiro um horário no fim do dia, etc"
            rows={4}
          />
          <DialogFooter>
            <Button variant="outline" onClick={() => setActionDialog(null)}>
              Cancelar
            </Button>
            <Button onClick={handleSendAction} disabled={requestAction.isPending}>
              {requestAction.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : "Enviar solicitação"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
