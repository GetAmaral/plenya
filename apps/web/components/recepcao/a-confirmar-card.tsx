"use client";

import Link from "next/link";
import { CheckCircle2, Check, Loader2, ExternalLink } from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  type Appointment,
  APPOINTMENT_TYPE_LABELS,
  useConfirmAppointment,
} from "@/lib/api/calendar-api";
import { EmptyState } from "./empty-state";
import { formatTimeBRT, formatRelativeDayTime } from "./recepcao-helpers";

/**
 * Linha de confirmacao com botao de um clique. Cada linha monta sua mutation
 * proprio com o id (useConfirmAppointment invalida calendarKeys.all, entao a
 * agenda e a contagem reagem juntas).
 */
function ConfirmRow({ appointment }: { appointment: Appointment }) {
  const confirmMutation = useConfirmAppointment(appointment.id);
  const patientName = appointment.patient?.name ?? "Paciente";

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

  return (
    <div className="flex items-center justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors">
      <div className="min-w-0">
        <p className="font-medium truncate">{patientName}</p>
        <p className="text-sm text-muted-foreground truncate">
          {formatRelativeDayTime(appointment.scheduledAt)} ·{" "}
          {APPOINTMENT_TYPE_LABELS[appointment.type]}
        </p>
      </div>
      <div className="flex items-center gap-2 shrink-0">
        <Button
          size="sm"
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
        <Button asChild size="sm" variant="ghost">
          <Link href={`/appointments/${appointment.id}`}>
            <ExternalLink className="h-4 w-4" />
          </Link>
        </Button>
      </div>
    </div>
  );
}

/**
 * "A confirmar": consultas de hoje e amanha ainda agendadas (status=scheduled),
 * ordenadas por horario. Mostra a contagem no cabecalho.
 */
export function AConfirmarCard({
  appointments,
  isLoading,
}: {
  appointments: Appointment[];
  isLoading: boolean;
}) {
  const pending = [...appointments]
    .filter((a) => a.status === "scheduled")
    .sort(
      (a, b) =>
        new Date(a.scheduledAt).getTime() - new Date(b.scheduledAt).getTime(),
    );

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <CheckCircle2 className="h-5 w-5 text-emerald-600" />
            A confirmar
          </span>
          {!isLoading && pending.length > 0 && (
            <Badge variant="observation">{pending.length}</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-14 w-full" />
            ))}
          </div>
        ) : pending.length === 0 ? (
          <EmptyState
            icon={CheckCircle2}
            title="Tudo confirmado"
            hint="Nao ha consultas pendentes de confirmacao para hoje e amanha."
          />
        ) : (
          <div className="space-y-3">
            {pending.map((a) => (
              <ConfirmRow key={a.id} appointment={a} />
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
