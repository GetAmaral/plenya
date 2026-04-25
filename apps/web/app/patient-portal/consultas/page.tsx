"use client";

/**
 * /consultas — lista de consultas do paciente.
 * Tabs: Próximas / Histórico.
 */
import { useState } from "react";
import Link from "next/link";
import { format, formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Loader2, Calendar, Video, CheckCircle2, Clock, XCircle } from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { useMyAppointments, type MyAppointment } from "@/lib/api/patient-portal-api";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";

const TYPE_LABEL: Record<MyAppointment["type"], string> = {
  initial_assessment: "Avaliação inicial",
  follow_up: "Retorno",
  telemedicine: "Teleconsulta",
  procedure: "Procedimento",
  results_review: "Revisão de exames",
};

export default function MyAppointmentsPage() {
  const { ready } = useRequirePatientAuth();
  const [tab, setTab] = useState<"upcoming" | "past">("upcoming");
  const { data, isLoading } = useMyAppointments(tab);

  if (!ready) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Consultas
        </p>
        <h1 className="text-3xl font-light">Minhas consultas</h1>
      </header>

      <Tabs value={tab} onValueChange={(v) => setTab(v as any)}>
        <TabsList>
          <TabsTrigger value="upcoming">Próximas</TabsTrigger>
          <TabsTrigger value="past">Histórico</TabsTrigger>
        </TabsList>

        <TabsContent value={tab} className="mt-4">
          {isLoading ? (
            <Card className="flex h-32 items-center justify-center">
              <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            </Card>
          ) : !data || data.length === 0 ? (
            <Card>
              <CardContent className="py-10 text-center text-sm text-muted-foreground">
                {tab === "upcoming"
                  ? "Nenhuma consulta agendada."
                  : "Nenhuma consulta no histórico."}
              </CardContent>
            </Card>
          ) : (
            <div className="space-y-3">
              {data.map((a) => (
                <AppointmentRow key={a.id} a={a} />
              ))}
            </div>
          )}
        </TabsContent>
      </Tabs>
    </div>
  );
}

function statusIcon(s: MyAppointment["status"]) {
  switch (s) {
    case "completed":
      return <CheckCircle2 className="h-4 w-4 text-emerald-600" />;
    case "cancelled":
    case "no_show":
      return <XCircle className="h-4 w-4 text-muted-foreground" />;
    case "confirmed":
      return <CheckCircle2 className="h-4 w-4 text-primary" />;
    default:
      return <Clock className="h-4 w-4 text-muted-foreground" />;
  }
}

function statusLabel(s: MyAppointment["status"]) {
  switch (s) {
    case "scheduled":
      return "Agendada";
    case "confirmed":
      return "Confirmada";
    case "completed":
      return "Realizada";
    case "cancelled":
      return "Cancelada";
    case "no_show":
      return "Não compareceu";
  }
}

function AppointmentRow({ a }: { a: MyAppointment }) {
  const when = new Date(a.scheduledAt);
  const isUpcoming = a.minutesUntilStart > -120;

  return (
    <Link href={`/consultas/${a.id}`}>
      <Card className="transition-colors hover:border-primary/40">
        <CardContent className="flex items-start justify-between gap-3 py-4">
          <div className="min-w-0 flex-1 space-y-1">
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              <Calendar className="h-3 w-3" />
              {format(when, "EEEE, d 'de' MMMM", { locale: ptBR })} ·{" "}
              {format(when, "HH:mm")}
            </div>
            <p className="font-medium">{TYPE_LABEL[a.type] ?? a.type}</p>
            <p className="text-sm text-muted-foreground">com {a.doctorName}</p>
            <div className="flex flex-wrap gap-1.5 pt-1">
              <Badge variant="outline" className="gap-1">
                {statusIcon(a.status)}
                {statusLabel(a.status)}
              </Badge>
              {a.isTelemedicine && (
                <Badge variant="outline" className="gap-1">
                  <Video className="h-3 w-3" /> Online
                </Badge>
              )}
              {isUpcoming && (
                <Badge variant="secondary">
                  {formatDistanceToNow(when, { addSuffix: true, locale: ptBR })}
                </Badge>
              )}
            </div>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}
