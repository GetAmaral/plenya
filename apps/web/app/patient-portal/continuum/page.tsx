"use client";

/**
 * /continuum — Meu Continuum (read-only). Agrupa items por mês,
 * mostra plano integrado em markdown, destaca próximo marco.
 */
import { useMemo } from "react";
import Link from "next/link";
import { formatDate } from "@/lib/format-date";
import {
  Loader2,
  CheckCircle2,
  Clock,
  AlertTriangle,
  XCircle,
  Calendar,
  Package,
  Activity,
  Workflow,
  Stethoscope,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyContinuum,
  type PatientContinuumItemView,
  type PatientContinuumItemStatus,
} from "@/lib/api/patient-portal-api";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { MarkdownContent } from "@/components/ui/markdown-content";
import { cn } from "@/lib/utils";

const STATUS_LABEL: Record<PatientContinuumItemStatus, string> = {
  pending: "Programado",
  scheduled: "Agendado",
  completed: "Concluído",
  missed: "Atrasado",
  cancelled: "Cancelado",
  skipped: "Pulado",
};

const TYPE_ICON = {
  appointment: Stethoscope,
  box: Package,
  reassessment: Activity,
  milestone: Workflow,
  custom: Calendar,
} as const;

function statusBadgeVariant(s: PatientContinuumItemStatus) {
  switch (s) {
    case "completed":
      return "default" as const;
    case "missed":
      return "destructive" as const;
    case "scheduled":
      return "secondary" as const;
    default:
      return "outline" as const;
  }
}

export default function MyContinuumPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyContinuum(ready);

  const groupedByMonth = useMemo(() => {
    if (!data?.continuum?.items) return [];
    const groups = new Map<string, PatientContinuumItemView[]>();
    for (const item of data.continuum.items) {
      const key = formatDate(item.expectedDate, "yyyy-MM");
      const arr = groups.get(key) ?? [];
      arr.push(item);
      groups.set(key, arr);
    }
    return Array.from(groups.entries()).sort(([a], [b]) => a.localeCompare(b));
  }, [data]);

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  const c = data?.continuum;
  if (!c) {
    return (
      <div className="mx-auto max-w-3xl space-y-6 py-10">
        <h1 className="text-3xl font-light">Meu Continuum</h1>
        <Card>
          <CardContent className="py-10 text-center text-muted-foreground">
            Você ainda não está em um programa Continuum. Fale com a equipe pra entender qual é o melhor pra você.
          </CardContent>
        </Card>
      </div>
    );
  }

  const itemsCompleted = c.items.filter((i) => i.status === "completed").length;
  const pct = c.items.length ? Math.round((itemsCompleted / c.items.length) * 100) : 0;
  const itemsLate = c.items.filter((i) => i.status === "missed").length;
  const nextItem = c.items
    .filter((i) => i.status === "pending" || i.status === "scheduled")
    .sort(
      (a, b) =>
        new Date(a.expectedDate).getTime() - new Date(b.expectedDate).getTime(),
    )[0];

  return (
    <div className="mx-auto max-w-4xl space-y-8">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Meu Continuum
        </p>
        <h1 className="text-3xl font-light md:text-4xl">{c.templateName || "Programa Plenya"}</h1>
        <p className="text-muted-foreground">
          {formatDate(c.startDate, "d 'de' MMMM 'de' yyyy")} →{" "}
          {formatDate(c.endDate, "d 'de' MMMM 'de' yyyy")}
        </p>
      </header>

      {/* Progresso */}
      <Card>
        <CardContent className="py-5">
          <div className="mb-1 flex justify-between text-sm">
            <span className="font-medium">{itemsCompleted} de {c.items.length} marcos concluídos</span>
            <span className="text-muted-foreground tabular-nums">{pct}%</span>
          </div>
          <Progress value={pct} className="h-2" />
          {itemsLate > 0 && (
            <p className="mt-2 flex items-center gap-1 text-xs text-amber-600">
              <AlertTriangle className="h-3.5 w-3.5" /> {itemsLate} marco{itemsLate > 1 ? "s" : ""} atrasado{itemsLate > 1 ? "s" : ""} — fale com a equipe.
            </p>
          )}
        </CardContent>
      </Card>

      {/* Próximo marco em destaque */}
      {nextItem && (
        <Card className="border-primary/30">
          <CardHeader className="pb-2">
            <CardTitle className="text-sm uppercase tracking-wider text-muted-foreground">
              Próximo marco
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-xl font-medium">{nextItem.title}</p>
            <p className="text-sm text-muted-foreground">
              {formatDate(nextItem.expectedDate, "EEEE, d 'de' MMMM")}
            </p>
            {nextItem.appointmentId && (
              <Link
                href={`/consultas/${nextItem.appointmentId}`}
                className="mt-2 inline-block text-sm underline-offset-4 hover:underline"
              >
                Ver consulta →
              </Link>
            )}
          </CardContent>
        </Card>
      )}

      {/* Plano integrado */}
      {c.integratedPlanMarkdown && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Plano da equipe</CardTitle>
            {c.integratedPlanUpdatedAt && (
              <p className="text-xs text-muted-foreground">
                Atualizado em{" "}
                {formatDate(c.integratedPlanUpdatedAt, "d 'de' MMMM")}
              </p>
            )}
          </CardHeader>
          <CardContent>
            <MarkdownContent content={c.integratedPlanMarkdown} className="text-sm" />
          </CardContent>
        </Card>
      )}

      {/* Timeline */}
      <section className="space-y-6">
        <h2 className="text-lg font-medium">Sua trajetória</h2>
        {groupedByMonth.map(([key, items]) => (
          <div key={key} className="space-y-2">
            <h3 className="text-xs font-semibold uppercase tracking-wider text-muted-foreground">
              {formatDate(key + "-01", "MMMM 'de' yyyy")}
            </h3>
            <div className="space-y-2">
              {items.map((item) => (
                <ContinuumItemRow key={item.id} item={item} />
              ))}
            </div>
          </div>
        ))}
      </section>
    </div>
  );
}

function ContinuumItemRow({ item }: { item: PatientContinuumItemView }) {
  const Icon = TYPE_ICON[item.type] ?? Calendar;
  const completed = item.status === "completed";
  const late = item.status === "missed";
  const cancelled = item.status === "cancelled" || item.status === "skipped";

  return (
    <div
      className={cn(
        "flex items-start gap-3 rounded-lg border bg-card p-3",
        completed && "border-emerald-200 bg-emerald-50/40 dark:bg-emerald-950/10",
        late && "border-amber-300 bg-amber-50/40 dark:bg-amber-950/10",
        cancelled && "opacity-60",
      )}
    >
      <div className="mt-1 shrink-0">
        {completed ? (
          <CheckCircle2 className="h-5 w-5 text-emerald-600" />
        ) : late ? (
          <AlertTriangle className="h-5 w-5 text-amber-600" />
        ) : cancelled ? (
          <XCircle className="h-5 w-5 text-muted-foreground" />
        ) : (
          <Clock className="h-5 w-5 text-muted-foreground" />
        )}
      </div>
      <div className="min-w-0 flex-1">
        <div className="flex items-center gap-2">
          <Icon className="h-3.5 w-3.5 text-muted-foreground" />
          <p className="truncate font-medium">{item.title}</p>
          <Badge variant={statusBadgeVariant(item.status)} className="shrink-0">
            {STATUS_LABEL[item.status]}
          </Badge>
        </div>
        <p className="mt-0.5 text-xs text-muted-foreground">
          {formatDate(item.expectedDate, "d 'de' MMMM")}
          {item.specialty && ` · ${item.specialty}`}
        </p>
        {item.description && (
          <p className="mt-1 text-sm text-muted-foreground">{item.description}</p>
        )}
        {item.appointmentId && (
          <Link
            href={`/consultas/${item.appointmentId}`}
            className="mt-1 inline-block text-xs underline-offset-4 hover:underline"
          >
            Abrir consulta →
          </Link>
        )}
      </div>
    </div>
  );
}
