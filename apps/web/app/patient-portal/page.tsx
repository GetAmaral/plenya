"use client";

/**
 * /  — Home da área do paciente.
 * Dashboard rico: próxima consulta, status Continuum, último escore, boxes,
 * mensagens, e banner de pendências quando aplicável.
 */
import { Loader2 } from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { usePatientDashboard } from "@/lib/api/patient-portal-api";
import {
  NextAppointmentCard,
  ContinuumProgressCard,
  LastScoreCard,
  BoxesCard,
  MessagesCard,
  PendingActionsBanner,
} from "@/components/patient-portal/home-cards";

export default function PatientHomePage() {
  const { ready, user } = useRequirePatientAuth();
  const { data, isLoading } = usePatientDashboard(ready);

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  const firstName = user?.name?.split(" ")[0] ?? "paciente";

  return (
    <div className="mx-auto max-w-5xl space-y-8">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Início</p>
        <h1 className="text-3xl font-light md:text-4xl">Olá, {firstName}.</h1>
        <p className="text-muted-foreground">
          Aqui você acompanha seu programa, consultas, exames e fala com a equipe.
        </p>
      </header>

      {data && <PendingActionsBanner actions={data.pendingActions} />}

      <div className="grid gap-4 md:grid-cols-2">
        <NextAppointmentCard data={data?.nextAppointment} />
        <ContinuumProgressCard data={data?.continuum} />
        <LastScoreCard data={data?.lastScore} />
        <BoxesCard data={data?.boxesCount ?? { preparing: 0, shipped: 0, delivered: 0 }} />
        <MessagesCard unread={data?.unreadMessages ?? 0} />
      </div>
    </div>
  );
}
