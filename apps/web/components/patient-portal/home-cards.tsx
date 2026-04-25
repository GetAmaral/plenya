"use client";

/**
 * Cards da home do portal do paciente.
 * Cada card recebe os dados já filtrados; renderiza estado vazio elegante quando ausente.
 */
import Link from "next/link";
import { format, formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Calendar,
  Workflow,
  Activity,
  Package,
  MessageSquare,
  ArrowUp,
  ArrowDown,
  Video,
  CheckCircle2,
  Clock,
  AlertTriangle,
} from "lucide-react";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import type { PatientDashboard } from "@/lib/api/patient-portal-api";

const TYPE_LABEL: Record<string, string> = {
  initial_assessment: "Avaliação inicial",
  follow_up: "Retorno",
  telemedicine: "Teleconsulta",
  procedure: "Procedimento",
  results_review: "Revisão de exames",
};

export function NextAppointmentCard({ data }: { data: PatientDashboard["nextAppointment"] }) {
  if (!data) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <Calendar className="h-4 w-4" /> Próxima consulta
          </CardTitle>
        </CardHeader>
        <CardContent className="text-sm text-muted-foreground">
          Nenhuma consulta agendada. A equipe entra em contato pra marcar.
        </CardContent>
      </Card>
    );
  }
  const when = new Date(data.scheduledAt);
  const isUpcoming = data.minutesUntilStart > 0;
  const canEnterTelemed =
    data.isTelemedicine && data.minutesUntilStart < 30 && data.minutesUntilStart > -120;

  return (
    <Card className="border-primary/30">
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <Calendar className="h-4 w-4" /> Próxima consulta
        </CardTitle>
        <CardDescription>
          {isUpcoming
            ? `${formatDistanceToNow(when, { addSuffix: true, locale: ptBR })}`
            : "Agora"}
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-3">
        <div>
          <p className="text-2xl font-light">
            {format(when, "EEEE, d 'de' MMMM", { locale: ptBR })}
          </p>
          <p className="text-3xl font-light tabular-nums">
            {format(when, "HH:mm")}
          </p>
        </div>
        <div className="flex flex-wrap gap-2">
          <Badge variant="secondary">{TYPE_LABEL[data.type] ?? data.type}</Badge>
          {data.isTelemedicine && (
            <Badge variant="outline" className="gap-1">
              <Video className="h-3 w-3" /> Online
            </Badge>
          )}
          <Badge variant="outline">com {data.doctorName}</Badge>
        </div>
        <div className="flex flex-wrap gap-2 pt-1">
          {canEnterTelemed && (
            <Link href={`/consultas/${data.id}`}>
              <Button size="sm" className="gap-1">
                <Video className="h-3.5 w-3.5" /> Entrar na sala
              </Button>
            </Link>
          )}
          <Link href={`/consultas/${data.id}`}>
            <Button size="sm" variant="outline">
              Ver detalhes
            </Button>
          </Link>
        </div>
      </CardContent>
    </Card>
  );
}

export function ContinuumProgressCard({ data }: { data: PatientDashboard["continuum"] }) {
  if (!data) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <Workflow className="h-4 w-4" /> Meu Continuum
          </CardTitle>
        </CardHeader>
        <CardContent className="text-sm text-muted-foreground">
          Você ainda não está em um programa Continuum. Fale com a equipe.
        </CardContent>
      </Card>
    );
  }

  const pct = data.itemsTotal > 0 ? Math.round((data.itemsCompleted / data.itemsTotal) * 100) : 0;

  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <Workflow className="h-4 w-4" /> {data.templateName || "Meu Continuum"}
        </CardTitle>
        <CardDescription>
          Semana {data.currentWeek} de {data.totalWeeks}
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div>
          <div className="mb-1 flex justify-between text-xs text-muted-foreground">
            <span>{data.itemsCompleted} concluídos</span>
            <span>{pct}%</span>
          </div>
          <Progress value={pct} className="h-2" />
        </div>
        {data.itemsLate > 0 && (
          <div className="flex items-center gap-2 text-xs text-amber-600">
            <AlertTriangle className="h-3.5 w-3.5" />
            {data.itemsLate} marco{data.itemsLate > 1 ? "s" : ""} atrasado{data.itemsLate > 1 ? "s" : ""}
          </div>
        )}
        {data.nextItemTitle && data.nextItemExpectedDate && (
          <div className="rounded-md bg-muted/40 p-3">
            <p className="text-xs uppercase tracking-wider text-muted-foreground">Próximo marco</p>
            <p className="mt-0.5 text-sm font-medium">{data.nextItemTitle}</p>
            <p className="text-xs text-muted-foreground">
              {format(new Date(data.nextItemExpectedDate), "d 'de' MMMM", { locale: ptBR })}
            </p>
          </div>
        )}
        <Link href="/continuum">
          <Button size="sm" variant="outline">
            Abrir programa
          </Button>
        </Link>
      </CardContent>
    </Card>
  );
}

export function LastScoreCard({ data }: { data: PatientDashboard["lastScore"] }) {
  if (!data) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-base">
            <Activity className="h-4 w-4" /> Seu último Escore
          </CardTitle>
        </CardHeader>
        <CardContent className="text-sm text-muted-foreground">
          Você ainda não tem nenhum Escore. Faça uma autoavaliação Light no site.
        </CardContent>
      </Card>
    );
  }
  const total = Math.round(data.totalScorePercentage);
  const delta = data.deltaVsPrevious;

  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <Activity className="h-4 w-4" /> Seu último Escore
        </CardTitle>
        <CardDescription>
          {data.source === "light" ? "Autoavaliação (Light)" : "Avaliação completa"} ·{" "}
          {format(new Date(data.createdAt), "d 'de' MMM", { locale: ptBR })}
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-3">
        <div className="flex items-baseline gap-3">
          <span className="text-5xl font-light tabular-nums">{total}</span>
          <span className="text-muted-foreground">/100</span>
          {typeof delta === "number" && (
            <span
              className={`flex items-center gap-0.5 text-sm ${delta >= 0 ? "text-emerald-600" : "text-amber-600"}`}
            >
              {delta >= 0 ? <ArrowUp className="h-3.5 w-3.5" /> : <ArrowDown className="h-3.5 w-3.5" />}
              {Math.abs(Math.round(delta))} pts
            </span>
          )}
        </div>
        <Link href="/escores">
          <Button size="sm" variant="outline">
            Ver evolução
          </Button>
        </Link>
      </CardContent>
    </Card>
  );
}

export function BoxesCard({ data }: { data: PatientDashboard["boxesCount"] }) {
  const total = data.preparing + data.shipped;
  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <Package className="h-4 w-4" /> Boxes Plenya
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        {total === 0 && data.delivered === 0 && (
          <p className="text-sm text-muted-foreground">Nenhum box em andamento.</p>
        )}
        {data.preparing > 0 && (
          <div className="flex items-center justify-between text-sm">
            <span className="flex items-center gap-2 text-muted-foreground">
              <Clock className="h-3.5 w-3.5" /> Em preparo
            </span>
            <span className="font-medium tabular-nums">{data.preparing}</span>
          </div>
        )}
        {data.shipped > 0 && (
          <div className="flex items-center justify-between text-sm">
            <span className="flex items-center gap-2 text-muted-foreground">
              <Package className="h-3.5 w-3.5" /> A caminho
            </span>
            <span className="font-medium tabular-nums">{data.shipped}</span>
          </div>
        )}
        {data.delivered > 0 && (
          <div className="flex items-center justify-between text-sm">
            <span className="flex items-center gap-2 text-muted-foreground">
              <CheckCircle2 className="h-3.5 w-3.5" /> Entregues
            </span>
            <span className="font-medium tabular-nums">{data.delivered}</span>
          </div>
        )}
      </CardContent>
    </Card>
  );
}

export function MessagesCard({ unread }: { unread: number }) {
  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <MessageSquare className="h-4 w-4" /> Mensagens
          {unread > 0 && <Badge>{unread}</Badge>}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2 text-sm text-muted-foreground">
        {unread === 0 ? (
          <p>Nenhuma mensagem nova.</p>
        ) : (
          <p>Você tem {unread} mensagem{unread > 1 ? "s" : ""} nova{unread > 1 ? "s" : ""}.</p>
        )}
        <Link href="/mensagens">
          <Button size="sm" variant="outline">
            Abrir caixa
          </Button>
        </Link>
      </CardContent>
    </Card>
  );
}

export function PendingActionsBanner({ actions }: { actions: PatientDashboard["pendingActions"] }) {
  if (!actions.length) return null;
  return (
    <div className="space-y-2">
      {actions.map((a) => (
        <Link
          key={a.kind + a.link}
          href={a.link}
          className="flex items-center justify-between rounded-lg border border-amber-300/50 bg-amber-50 px-4 py-3 text-sm text-amber-900 transition-colors hover:bg-amber-100 dark:bg-amber-950/30 dark:text-amber-200"
        >
          <span className="flex items-center gap-2">
            <AlertTriangle className="h-4 w-4" /> {a.description}
          </span>
          <span className="text-xs underline-offset-4 hover:underline">abrir →</span>
        </Link>
      ))}
    </div>
  );
}
