"use client";

import { CalendarClock, CalendarCheck } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import type { Appointment } from "@/lib/api/calendar-api";
import { AppointmentRow } from "./appointment-row";
import { EmptyState } from "./empty-state";

/**
 * "Agenda de hoje": todas as consultas do dia, todos os medicos, ordenadas por
 * horario. A ordenacao e feita aqui pra manter a fonte (hook) generica.
 */
export function AgendaHojeCard({
  appointments,
  isLoading,
}: {
  appointments: Appointment[];
  isLoading: boolean;
}) {
  const sorted = [...appointments].sort(
    (a, b) =>
      new Date(a.scheduledAt).getTime() - new Date(b.scheduledAt).getTime(),
  );

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <CalendarClock className="h-5 w-5 text-blue-600" />
            Agenda de hoje
          </span>
          {!isLoading && sorted.length > 0 && (
            <Badge variant="secondary">{sorted.length}</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 4 }).map((_, i) => (
              <Skeleton key={i} className="h-16 w-full" />
            ))}
          </div>
        ) : sorted.length === 0 ? (
          <EmptyState
            icon={CalendarCheck}
            title="Nenhuma consulta para hoje"
            hint="A agenda do dia esta livre. Aproveite para revisar pendencias."
          />
        ) : (
          <div className="space-y-3">
            {sorted.map((a) => (
              <AppointmentRow key={a.id} appointment={a} />
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
