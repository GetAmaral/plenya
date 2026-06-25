"use client";

/**
 * Home do médico — agenda do dia + atendimentos em curso + próximas consultas.
 *
 * Dados reais via useAppointments (mesma query da Recepção, polling 15s).
 * Substitui o dashboard mock anterior (dados hardcoded que não navegavam).
 */
import { useMemo } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { formatDate } from "@/lib/format-date";
import {
  AlertTriangle,
  ArrowRight,
  Calendar,
  CheckCircle2,
  Clock,
  Loader2,
  Plus,
  Stethoscope,
  Video,
} from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { PageHeader } from "@/components/layout/page-header";
import { useRequireAuth } from "@/lib/use-auth";
import {
  APPOINTMENT_STATUS_COLORS,
  APPOINTMENT_STATUS_LABELS,
  APPOINTMENT_TYPE_LABELS,
  useAppointments,
  type Appointment,
} from "@/lib/api/calendar-api";

export default function DashboardPage() {
  useRequireAuth();
  const router = useRouter();

  const { from, to } = useMemo(() => {
    const now = new Date();
    const start = new Date(now.getFullYear(), now.getMonth(), now.getDate());
    const end = new Date(now.getFullYear(), now.getMonth(), now.getDate() + 1);
    return { from: start.toISOString(), to: end.toISOString() };
  }, []);

  const { data: appts = [], isLoading, isError } = useAppointments({
    dateFrom: from,
    dateTo: to,
    limit: 200,
  });

  const sorted = useMemo(
    () =>
      [...appts].sort(
        (a, b) => new Date(a.scheduledAt).getTime() - new Date(b.scheduledAt).getTime(),
      ),
    [appts],
  );

  const emAtendimento = sorted.filter(
    (a) => a.status === "in_progress" || a.status === "checked_in",
  );
  const concluidas = sorted.filter((a) => a.status === "completed");
  const ativas = sorted.filter((a) => a.status !== "cancelled" && a.status !== "no_show");
  const proximas = useMemo(() => {
    const nowMs = Date.now();
    return sorted.filter(
      (a) =>
        (a.status === "scheduled" || a.status === "confirmed") &&
        new Date(a.scheduledAt).getTime() >= nowMs - 60 * 60 * 1000,
    );
  }, [sorted]);
  const proximaTime = proximas[0] ? formatDate(proximas[0].scheduledAt, "HH:mm") : "—";

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        title="Início"
        description={format(new Date(), "EEEE, dd 'de' MMMM yyyy", { locale: ptBR })}
        actions={[
          {
            label: "Calendário",
            icon: <Calendar className="h-4 w-4" />,
            variant: "outline",
            onClick: () => router.push("/calendario"),
          },
          {
            label: "Nova consulta",
            icon: <Plus className="h-4 w-4" />,
            variant: "default",
            onClick: () => router.push("/appointments/new"),
          },
        ]}
      />

      {/* KPIs */}
      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        <StatCard icon={Calendar} label="Consultas hoje" value={ativas.length} tone="blue" />
        <StatCard icon={Stethoscope} label="Em atendimento" value={emAtendimento.length} tone="amber" />
        <StatCard icon={CheckCircle2} label="Concluídas hoje" value={concluidas.length} tone="emerald" />
        <StatCard icon={Clock} label="Próxima" value={proximaTime} tone="violet" />
      </div>

      {isError ? (
        <Card>
          <CardContent className="flex items-center gap-3 py-8 text-sm text-destructive">
            <AlertTriangle className="h-5 w-5 shrink-0" />
            Não foi possível carregar a agenda de hoje. Verifique a conexão e tente novamente.
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-6 lg:grid-cols-3">
          {/* Agenda de hoje */}
          <Card className="lg:col-span-2">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <Calendar className="h-5 w-5 text-ocean-600" />
                Agenda de hoje
              </CardTitle>
            </CardHeader>
            <CardContent>
              {isLoading ? (
                <div className="flex h-32 items-center justify-center">
                  <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
                </div>
              ) : ativas.length === 0 ? (
                <p className="py-8 text-center text-sm text-muted-foreground">
                  Nenhuma consulta agendada para hoje.
                </p>
              ) : (
                <div className="space-y-2">
                  {ativas.map((a) => (
                    <AppointmentRow key={a.id} appt={a} onClick={() => router.push(`/appointments/${a.id}`)} />
                  ))}
                </div>
              )}
            </CardContent>
          </Card>

          {/* Em atendimento agora */}
          <Card className="h-fit">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <Stethoscope className="h-5 w-5 text-amber-600" />
                Em atendimento agora
              </CardTitle>
            </CardHeader>
            <CardContent>
              {emAtendimento.length === 0 ? (
                <p className="py-6 text-center text-sm text-muted-foreground">
                  Nenhum atendimento em curso.
                </p>
              ) : (
                <div className="space-y-2">
                  {emAtendimento.map((a) => (
                    <Link
                      key={a.id}
                      href={`/appointments/${a.id}`}
                      className="block rounded-lg border border-amber-200 bg-amber-50/60 p-3 transition hover:border-amber-300"
                    >
                      <div className="flex items-center justify-between gap-2">
                        <p className="truncate text-sm font-medium">{a.patient?.name ?? "Paciente"}</p>
                        <Badge variant="outline" className={APPOINTMENT_STATUS_COLORS[a.status]}>
                          {APPOINTMENT_STATUS_LABELS[a.status]}
                        </Badge>
                      </div>
                      <div className="mt-1 flex items-center gap-2 text-xs text-muted-foreground">
                        <span>{APPOINTMENT_TYPE_LABELS[a.type]}</span>
                        <span>·</span>
                        <span>{formatDate(a.scheduledAt, "HH:mm")}</span>
                        <ArrowRight className="ml-auto h-3.5 w-3.5" />
                      </div>
                    </Link>
                  ))}
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      )}
    </div>
  );
}

function StatCard({
  icon: Icon,
  label,
  value,
  tone,
}: {
  icon: React.ComponentType<{ className?: string }>;
  label: string;
  value: number | string;
  tone: "blue" | "amber" | "emerald" | "violet";
}) {
  const tones: Record<string, string> = {
    blue: "bg-ocean-50 text-ocean-600 dark:bg-petrol/40",
    amber: "bg-amber-50 text-amber-600 dark:bg-amber-950/30",
    emerald: "bg-emerald-50 text-emerald-600 dark:bg-emerald-950/30",
    violet: "bg-gold-50 text-gold-600 dark:bg-gold-900/20",
  };
  return (
    <Card>
      <CardContent className="flex items-center justify-between p-5">
        <div>
          <p className="text-sm text-muted-foreground">{label}</p>
          <p className="mt-1 text-2xl font-bold">{value}</p>
        </div>
        <div className={`rounded-xl p-3 ${tones[tone]}`}>
          <Icon className="h-5 w-5" />
        </div>
      </CardContent>
    </Card>
  );
}

function AppointmentRow({ appt, onClick }: { appt: Appointment; onClick: () => void }) {
  const isTelemed = appt.type === "telemedicine";
  return (
    <button
      type="button"
      onClick={onClick}
      className="flex w-full items-center gap-3 rounded-lg border border-border p-3 text-left transition hover:bg-accent/50"
    >
      <div className="flex w-14 shrink-0 flex-col items-center">
        <span className="text-sm font-semibold">{formatDate(appt.scheduledAt, "HH:mm")}</span>
        <span className="text-[10px] text-muted-foreground">{appt.durationMinutes}min</span>
      </div>
      <div className="min-w-0 flex-1">
        <p className="truncate text-sm font-medium">{appt.patient?.name ?? "Paciente"}</p>
        <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
          {isTelemed && <Video className="h-3 w-3" />}
          <span>{APPOINTMENT_TYPE_LABELS[appt.type]}</span>
        </div>
      </div>
      <Badge variant="outline" className={APPOINTMENT_STATUS_COLORS[appt.status]}>
        {APPOINTMENT_STATUS_LABELS[appt.status]}
      </Badge>
    </button>
  );
}
