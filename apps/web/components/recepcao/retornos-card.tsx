"use client";

import { useRouter } from "next/navigation";
import { CalendarClock, CalendarPlus, Loader2, X } from "lucide-react";
import { toast } from "sonner";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { useRecalls, useUpdateRecall, type Recall } from "@/lib/api/recalls";
import { EmptyState } from "./empty-state";
import { ErrorState } from "./error-state";

function fmtDate(due: string): string {
  const [y, m, d] = due.split("-");
  return `${d}/${m}/${y}`;
}

function RecallRow({ recall }: { recall: Recall }) {
  const router = useRouter();
  const update = useUpdateRecall(recall.id);
  const name = recall.patient?.name ?? "Paciente";
  const overdue = recall.dueDate <= new Date().toISOString().slice(0, 10);

  function handleSchedule() {
    // Leva pra Nova consulta com o paciente em contexto; a secretaria agenda lá.
    const params = new URLSearchParams({ patientId: recall.patientId });
    router.push(`/appointments/new?${params.toString()}`);
  }

  function handleDismiss() {
    update.mutate(
      { status: "dismissed" },
      {
        onSuccess: () => toast.success("Retorno dispensado", { description: name }),
        onError: (e: unknown) =>
          toast.error(e instanceof Error ? e.message : "Falha ao dispensar"),
      },
    );
  }

  return (
    <div className="flex items-center justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors">
      <div className="min-w-0">
        <p className="font-medium truncate">{name}</p>
        <p className="text-sm text-muted-foreground truncate">
          {recall.reason || "Retorno"}
        </p>
      </div>
      <div className="flex items-center gap-2 shrink-0">
        <Badge
          variant="outline"
          className={overdue ? "border-amber-300 bg-amber-50 text-amber-900" : ""}
          title={overdue ? "Vencido" : "Previsto"}
        >
          {fmtDate(recall.dueDate)}
        </Badge>
        <Button
          size="sm"
          variant="outline"
          onClick={handleSchedule}
          title="Agendar retorno"
        >
          <CalendarPlus className="h-4 w-4" />
          <span className="hidden sm:inline">Agendar</span>
        </Button>
        <Button
          size="sm"
          variant="ghost"
          onClick={handleDismiss}
          disabled={update.isPending}
          title="Dispensar"
        >
          {update.isPending ? (
            <Loader2 className="h-4 w-4 animate-spin" />
          ) : (
            <X className="h-4 w-4" />
          )}
        </Button>
      </div>
    </div>
  );
}

/**
 * "Retornos previstos": fila de recalls pendentes (mais antigo/vencido no topo).
 * Cada linha leva a secretária a agendar (Nova consulta com o paciente) ou a
 * dispensar. Sem aviso automático — contato é manual.
 */
export function RetornosCard() {
  const { data: recalls = [], isLoading, isError, refetch } =
    useRecalls("pending");

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <CalendarClock className="h-5 w-5 text-teal-600" />
            Retornos previstos
          </span>
          {!isLoading && !isError && recalls.length > 0 && (
            <Badge variant="observation">{recalls.length}</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {isError ? (
          <ErrorState
            title="Nao foi possivel carregar os retornos"
            onRetry={() => refetch()}
          />
        ) : isLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-16 w-full" />
            ))}
          </div>
        ) : recalls.length === 0 ? (
          <EmptyState
            icon={CalendarClock}
            title="Nenhum retorno previsto"
            hint="Ao concluir uma consulta, registre o retorno para acompanhar aqui."
          />
        ) : (
          <div className="space-y-3">
            {recalls.map((r) => (
              <RecallRow key={r.id} recall={r} />
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
