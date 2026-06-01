"use client";

import Link from "next/link";
import { useState } from "react";
import {
  ExternalLink,
  Check,
  Video,
  Loader2,
  LogIn,
  Play,
  CheckCheck,
  DollarSign,
  UserX,
} from "lucide-react";
import { toast } from "sonner";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { RegisterPaymentDialog } from "@/components/payments/register-payment-dialog";
import { cn } from "@/lib/utils";
import {
  type Appointment,
  APPOINTMENT_STATUS_LABELS,
  APPOINTMENT_STATUS_COLORS,
  APPOINTMENT_TYPE_LABELS,
  useConfirmAppointment,
  useCheckInAppointment,
  useStartAppointment,
  useUpdateAppointment,
  useTelemedToken,
  waitMinutes,
} from "@/lib/api/calendar-api";
import { useWaitlist } from "@/lib/api/waitlist";
import { formatBRL, type Payment } from "@/lib/api/payments";
import {
  STATUS_BADGE_VARIANT,
  STATUS_DOT_CLASS,
  formatTimeBRT,
  appointmentTypeIsTelemed,
} from "./recepcao-helpers";

/**
 * Linha da "Agenda de hoje". Hora, paciente, medico, tipo, status, e acoes
 * rapidas inline.
 *
 * Fluxo de balcao por status:
 *   - "Chegou" (agendada/confirmada -> aguardando)
 *   - "Iniciar" (confirmada/aguardando -> em atendimento)
 *   - "Concluir" (em atendimento -> concluida)
 * Mais Confirmar (so quando agendada), Teleconsulta (so telemedicina),
 * Pagamento (abre o dialogo de registro) e Abrir (detalhe).
 *
 * Cada acao e uma mutation propria por linha — cada linha monta seus hooks
 * com o id da consulta. As mutations ja invalidam as queries do calendario.
 */
export function AppointmentRow({
  appointment,
  payment,
}: {
  appointment: Appointment;
  payment?: Payment;
}) {
  const confirmMutation = useConfirmAppointment(appointment.id);
  const checkInMutation = useCheckInAppointment(appointment.id);
  const startMutation = useStartAppointment(appointment.id);
  const completeMutation = useUpdateAppointment(appointment.id);
  const telemedMutation = useTelemedToken(appointment.id);
  // Encaixe ativo: ao liberar vaga (falta), avisamos quantos aguardam encaixe.
  const { data: waitlist = [] } = useWaitlist("waiting");

  const [payOpen, setPayOpen] = useState(false);

  const patientName = appointment.patient?.name ?? "Paciente";
  const patientId = appointment.patient?.id;
  const doctorName = appointment.doctor?.name ?? "Profissional";

  async function handleConfirm() {
    try {
      await confirmMutation.mutateAsync();
      toast.success("Consulta confirmada", {
        description: `${patientName} as ${formatTimeBRT(appointment.scheduledAt)}.`,
      });
    } catch (err) {
      toast.error("Nao foi possivel confirmar", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  async function handleCheckIn() {
    try {
      await checkInMutation.mutateAsync();
      toast.success("Chegada registrada", {
        description: `${patientName} esta aguardando atendimento.`,
      });
    } catch (err) {
      toast.error("Nao foi possivel registrar a chegada", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  async function handleStart() {
    try {
      await startMutation.mutateAsync();
      toast.success("Atendimento iniciado", {
        description: `${patientName} em atendimento.`,
      });
    } catch (err) {
      toast.error("Nao foi possivel iniciar o atendimento", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  async function handleComplete() {
    try {
      await completeMutation.mutateAsync({ status: "completed" });
      toast.success("Atendimento concluido", {
        description: `${patientName}.`,
      });
    } catch (err) {
      toast.error("Nao foi possivel concluir o atendimento", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  async function handleNoShow() {
    try {
      await completeMutation.mutateAsync({ status: "no_show" });
      toast.success("Falta registrada", {
        description:
          waitlist.length > 0
            ? `Vaga livre. ${waitlist.length} na lista de espera para encaixar.`
            : `${patientName} nao compareceu.`,
      });
    } catch (err) {
      toast.error("Nao foi possivel registrar a falta", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  async function handleStartTelemed() {
    try {
      const { joinUrl } = await telemedMutation.mutateAsync();
      window.open(joinUrl, "_blank", "noopener,noreferrer");
    } catch (err) {
      toast.error("Nao foi possivel abrir a sala", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  const status = appointment.status;
  const isScheduled = status === "scheduled";
  const isTelemed = appointmentTypeIsTelemed(appointment.type);

  const canCheckIn = status === "scheduled" || status === "confirmed";
  const canStart = status === "confirmed" || status === "checked_in";
  const canComplete = status === "in_progress";
  const canNoShow = status === "scheduled" || status === "confirmed";

  // Tempo de espera, mostrado discreto ao lado do status enquanto aguarda ou
  // ja esta em atendimento.
  const wait =
    status === "checked_in" || status === "in_progress"
      ? waitMinutes(appointment)
      : null;
  const waitLabel =
    wait != null
      ? status === "in_progress"
        ? `espera ${wait} min`
        : `esperando ${wait} min`
      : null;

  // STATUS_BADGE_VARIANT/STATUS_DOT_CLASS nao cobrem checked_in/in_progress;
  // nesses casos caimos no mapa de cores completo do calendar-api.
  const badgeVariant = STATUS_BADGE_VARIANT[status];
  const dotClass = STATUS_DOT_CLASS[status];

  return (
    <div className="flex items-center justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors">
      <div className="flex items-center gap-3 min-w-0">
        <div className="flex flex-col items-center justify-center w-14 shrink-0">
          <span className="text-base font-semibold tabular-nums">
            {formatTimeBRT(appointment.scheduledAt)}
          </span>
          <span
            className={cn(
              "mt-1 h-1.5 w-1.5 rounded-full",
              dotClass ?? "bg-violet-500",
            )}
            aria-hidden
          />
        </div>
        <div className="min-w-0">
          <p className="font-medium truncate">{patientName}</p>
          <p className="text-sm text-muted-foreground truncate">
            {doctorName} · {APPOINTMENT_TYPE_LABELS[appointment.type]}
          </p>
        </div>
      </div>

      <div className="flex items-center gap-2 shrink-0">
        <div className="flex flex-col items-end gap-0.5">
          {badgeVariant ? (
            <Badge variant={badgeVariant}>
              {APPOINTMENT_STATUS_LABELS[status]}
            </Badge>
          ) : (
            <Badge variant="outline" className={APPOINTMENT_STATUS_COLORS[status]}>
              {APPOINTMENT_STATUS_LABELS[status]}
            </Badge>
          )}
          {waitLabel && (
            <span className="text-xs text-muted-foreground tabular-nums">
              {waitLabel}
            </span>
          )}
        </div>

        {isScheduled && (
          <Button
            size="sm"
            variant="outline"
            onClick={handleConfirm}
            disabled={confirmMutation.isPending}
          >
            {confirmMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Check className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Confirmar</span>
          </Button>
        )}

        {canCheckIn && (
          <Button
            size="sm"
            variant="outline"
            onClick={handleCheckIn}
            disabled={checkInMutation.isPending}
          >
            {checkInMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <LogIn className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Chegou</span>
          </Button>
        )}

        {canStart && (
          <Button
            size="sm"
            onClick={handleStart}
            disabled={startMutation.isPending}
          >
            {startMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Play className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Iniciar</span>
          </Button>
        )}

        {canComplete && (
          <Button
            size="sm"
            variant="outline"
            onClick={handleComplete}
            disabled={completeMutation.isPending}
          >
            {completeMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <CheckCheck className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Concluir</span>
          </Button>
        )}

        {canNoShow && (
          <Button
            size="sm"
            variant="ghost"
            onClick={handleNoShow}
            disabled={completeMutation.isPending}
            title="Marcar como nao compareceu"
            className="text-muted-foreground hover:text-rose-700"
          >
            <UserX className="h-4 w-4" />
            <span className="hidden sm:inline">Faltou</span>
          </Button>
        )}

        {isTelemed && (
          <Button
            size="sm"
            variant="ghost"
            onClick={handleStartTelemed}
            disabled={telemedMutation.isPending}
          >
            {telemedMutation.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Video className="h-4 w-4" />
            )}
            <span className="hidden sm:inline">Teleconsulta</span>
          </Button>
        )}

        {payment ? (
          <Badge
            variant="outline"
            className="border-emerald-200 bg-emerald-100 text-emerald-900"
            title={`Recibo ${payment.receiptNumber}`}
          >
            Pago {formatBRL(payment.amountCents)}
          </Badge>
        ) : (
          patientId && (
            <Button
              size="sm"
              variant="ghost"
              onClick={() => setPayOpen(true)}
              title="Registrar pagamento"
            >
              <DollarSign className="h-4 w-4" />
              <span className="hidden sm:inline">Pagamento</span>
            </Button>
          )
        )}

        <Button asChild size="sm" variant="ghost">
          <Link href={`/appointments/${appointment.id}`}>
            <ExternalLink className="h-4 w-4" />
            <span className="hidden sm:inline">Abrir</span>
          </Link>
        </Button>
      </div>

      {patientId && (
        <RegisterPaymentDialog
          open={payOpen}
          onOpenChange={setPayOpen}
          patientId={patientId}
          patientName={appointment.patient?.name}
          appointmentId={appointment.id}
          appointmentType={appointment.type}
        />
      )}
    </div>
  );
}
