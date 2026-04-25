"use client";

/**
 * /boxes — boxes Plenya (mimos + manipulados) com tracking.
 * Read-only pra paciente — equipe gerencia o fluxo logístico no EMR.
 */
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Loader2,
  Package,
  Clock,
  Truck,
  CheckCircle2,
  XCircle,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { useMyBoxes, type PatientBoxView } from "@/lib/api/patient-portal-api";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

const STATUS_LABEL: Record<PatientBoxView["status"], string> = {
  planned: "Planejado",
  preparing: "Em preparo",
  shipped: "A caminho",
  delivered: "Entregue",
  cancelled: "Cancelado",
};

const STATUS_ICON = {
  planned: Clock,
  preparing: Package,
  shipped: Truck,
  delivered: CheckCircle2,
  cancelled: XCircle,
} as const;

export default function MyBoxesPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyBoxes();

  if (!ready || isLoading) {
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
          Logística
        </p>
        <h1 className="text-3xl font-light">Meus boxes Plenya</h1>
        <p className="text-muted-foreground">
          Cada box leva mimos selecionados pela equipe e os manipulados/suplementos do seu protocolo.
        </p>
      </header>

      {!data?.length ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            Nenhum box programado ainda.
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-3">
          {data.map((b) => (
            <BoxRow key={b.id} b={b} />
          ))}
        </div>
      )}
    </div>
  );
}

function BoxRow({ b }: { b: PatientBoxView }) {
  const Icon = STATUS_ICON[b.status];
  const isFinal = b.status === "delivered";
  const trackingUrl =
    b.trackingCode && b.carrier?.toLowerCase().includes("correios")
      ? `https://rastreamento.correios.com.br/app/index.php?objetos=${b.trackingCode}`
      : null;

  return (
    <Card>
      <CardContent className="space-y-3 py-4">
        <div className="flex items-start justify-between gap-3">
          <div className="min-w-0 flex-1">
            <p className="font-medium">{b.name}</p>
            {b.contents && (
              <p className="mt-1 text-sm text-muted-foreground line-clamp-2">{b.contents}</p>
            )}
          </div>
          <Badge variant={isFinal ? "default" : "outline"} className="gap-1 shrink-0">
            <Icon className="h-3 w-3" />
            {STATUS_LABEL[b.status]}
          </Badge>
        </div>

        <div className="flex flex-wrap gap-x-4 gap-y-1 text-xs text-muted-foreground">
          {b.expectedDate && (
            <span>
              Previsto: {format(new Date(b.expectedDate), "dd 'de' MMM", { locale: ptBR })}
            </span>
          )}
          {b.shippedAt && (
            <span>
              Enviado em {format(new Date(b.shippedAt), "dd/MM/yyyy")}
            </span>
          )}
          {b.deliveredAt && (
            <span>
              Entregue em {format(new Date(b.deliveredAt), "dd/MM/yyyy")}
            </span>
          )}
          {b.carrier && <span>via {b.carrier}</span>}
        </div>

        {b.trackingCode && (
          <div className="flex items-center gap-2 text-sm">
            <span className="text-muted-foreground">Rastreio:</span>
            <code className="rounded bg-muted px-2 py-0.5 text-xs">{b.trackingCode}</code>
            {trackingUrl && (
              <a
                href={trackingUrl}
                target="_blank"
                rel="noreferrer"
                className="text-xs underline-offset-4 hover:underline"
              >
                Acompanhar →
              </a>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
