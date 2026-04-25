"use client";

/**
 * /exames — tabs Resultados / Pedidos / Prescrições / Avaliações.
 */
import { useState } from "react";
import Link from "next/link";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Loader2,
  FlaskConical,
  ClipboardList,
  Pill,
  Activity,
  FileDown,
  ExternalLink,
  CheckCircle2,
  Clock,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyLabBatches,
  useMyLabRequests,
  useMyPrescriptions,
  useMyPhysicalAssessments,
} from "@/lib/api/patient-portal-api";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";

export default function MyExamsPage() {
  const { ready } = useRequirePatientAuth();
  const [tab, setTab] = useState("results");

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
          Saúde
        </p>
        <h1 className="text-3xl font-light">Meus exames e laudos</h1>
      </header>

      <Tabs value={tab} onValueChange={setTab}>
        <TabsList className="grid w-full grid-cols-2 md:grid-cols-4">
          <TabsTrigger value="results">Resultados</TabsTrigger>
          <TabsTrigger value="requests">Pedidos</TabsTrigger>
          <TabsTrigger value="prescriptions">Prescrições</TabsTrigger>
          <TabsTrigger value="assessments">Avaliações</TabsTrigger>
        </TabsList>

        <TabsContent value="results" className="mt-4">
          <ResultsTab />
        </TabsContent>
        <TabsContent value="requests" className="mt-4">
          <RequestsTab />
        </TabsContent>
        <TabsContent value="prescriptions" className="mt-4">
          <PrescriptionsTab />
        </TabsContent>
        <TabsContent value="assessments" className="mt-4">
          <AssessmentsTab />
        </TabsContent>
      </Tabs>
    </div>
  );
}

function ResultsTab() {
  const { data, isLoading } = useMyLabBatches();
  if (isLoading) return <LoadingCard />;
  if (!data?.length) return <EmptyCard text="Nenhum resultado de exame ainda." />;
  return (
    <div className="space-y-3">
      {data.map((b) => (
        <Link key={b.id} href={`/exames/${b.id}`}>
          <Card className="transition-colors hover:border-primary/40">
            <CardContent className="flex items-center justify-between gap-3 py-4">
              <div className="min-w-0 flex-1 space-y-1">
                <div className="flex items-center gap-2 text-xs text-muted-foreground">
                  <FlaskConical className="h-3 w-3" />
                  {format(new Date(b.collectionDate), "dd 'de' MMM 'de' yyyy", { locale: ptBR })}
                </div>
                <p className="truncate font-medium">{b.laboratoryName}</p>
                <div className="flex flex-wrap gap-1.5">
                  <Badge variant="outline" className="gap-1">
                    {b.status === "completed" ? (
                      <CheckCircle2 className="h-3 w-3 text-emerald-600" />
                    ) : (
                      <Clock className="h-3 w-3" />
                    )}
                    {statusLabel(b.status)}
                  </Badge>
                  <Badge variant="secondary">{b.resultsCount} resultados</Badge>
                  {b.requestingDoctor && <Badge variant="outline">por {b.requestingDoctor}</Badge>}
                </div>
              </div>
            </CardContent>
          </Card>
        </Link>
      ))}
    </div>
  );
}

function statusLabel(s: string) {
  switch (s) {
    case "completed":
      return "Completo";
    case "partial":
      return "Parcial";
    default:
      return "Pendente";
  }
}

function RequestsTab() {
  const { data, isLoading } = useMyLabRequests();
  if (isLoading) return <LoadingCard />;
  if (!data?.length) return <EmptyCard text="Nenhum pedido de exame pendente." />;
  return (
    <div className="space-y-3">
      {data.map((r) => (
        <Card key={r.id}>
          <CardContent className="space-y-2 py-4">
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              <ClipboardList className="h-3 w-3" />
              {format(new Date(r.date), "dd 'de' MMM 'de' yyyy", { locale: ptBR })}
              {r.doctorName && <> · por {r.doctorName}</>}
            </div>
            <p className="whitespace-pre-wrap text-sm">{r.exams}</p>
            {r.pdfUrl && (
              <a
                href={r.pdfUrl}
                target="_blank"
                rel="noreferrer"
                className="inline-flex items-center gap-1 text-xs underline-offset-4 hover:underline"
              >
                <FileDown className="h-3 w-3" /> Baixar PDF
              </a>
            )}
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

function PrescriptionsTab() {
  const { data, isLoading } = useMyPrescriptions();
  if (isLoading) return <LoadingCard />;
  if (!data?.length) return <EmptyCard text="Nenhuma prescrição." />;
  return (
    <div className="space-y-3">
      {data.map((p) => (
        <Card key={p.id}>
          <CardContent className="space-y-2 py-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2 text-xs text-muted-foreground">
                <Pill className="h-3 w-3" />
                {format(new Date(p.issuedAt), "dd 'de' MMM 'de' yyyy", { locale: ptBR })}
                {p.doctorName && <> · por {p.doctorName}</>}
              </div>
              <Badge variant={p.status === "active" ? "default" : "outline"}>{statusLabelPx(p.status)}</Badge>
            </div>
            <p className="text-sm">{p.medsCount} medicamento{p.medsCount !== 1 ? "s" : ""}</p>
            <p className="text-xs text-muted-foreground">
              Válida até {format(new Date(p.validUntil), "dd/MM/yyyy")}
            </p>
            {p.pdfPath && (
              <a
                href={p.pdfPath}
                target="_blank"
                rel="noreferrer"
                className="inline-flex items-center gap-1 text-xs underline-offset-4 hover:underline"
              >
                <FileDown className="h-3 w-3" /> Baixar receita
              </a>
            )}
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

function statusLabelPx(s: string) {
  switch (s) {
    case "active":
      return "Ativa";
    case "completed":
      return "Concluída";
    case "cancelled":
      return "Cancelada";
    case "expired":
      return "Vencida";
    default:
      return s;
  }
}

function AssessmentsTab() {
  const { data, isLoading } = useMyPhysicalAssessments();
  if (isLoading) return <LoadingCard />;
  if (!data?.length) return <EmptyCard text="Nenhuma avaliação física registrada." />;
  return (
    <div className="space-y-3">
      {data.map((a) => (
        <Card key={a.id}>
          <CardContent className="space-y-2 py-4">
            <div className="flex items-center gap-2 text-xs text-muted-foreground">
              <Activity className="h-3 w-3" />
              {format(new Date(a.assessmentDate), "dd 'de' MMM 'de' yyyy", { locale: ptBR })}
              {a.assessorName && <> · por {a.assessorName}</>}
            </div>
            <div className="flex flex-wrap gap-3 text-sm">
              {a.weight != null && <span>Peso: {a.weight} kg</span>}
              {a.height != null && <span>Altura: {a.height} cm</span>}
              {a.bmi != null && <span>IMC: {a.bmi.toFixed(1)}</span>}
            </div>
            <a
              href={`${process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001"}/api/v1/patient/me/physical-assessments/${a.id}/html`}
              target="_blank"
              rel="noreferrer"
              className="inline-flex items-center gap-1 text-xs underline-offset-4 hover:underline"
            >
              <ExternalLink className="h-3 w-3" /> Abrir relatório completo
            </a>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}

function LoadingCard() {
  return (
    <Card className="flex h-32 items-center justify-center">
      <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
    </Card>
  );
}

function EmptyCard({ text }: { text: string }) {
  return (
    <Card>
      <CardContent className="py-10 text-center text-sm text-muted-foreground">{text}</CardContent>
    </Card>
  );
}
