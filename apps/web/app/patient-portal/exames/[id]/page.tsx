"use client";

/**
 * /exames/[id] — detalhe do batch de resultados.
 */
import { use } from "react";
import Link from "next/link";
import { Loader2, ArrowLeft, FlaskConical } from "lucide-react";

import { formatDate } from "@/lib/format-date";
import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { useMyLabBatch } from "@/lib/api/patient-portal-api";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

export default function MyLabBatchDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyLabBatch(ready ? id : undefined);

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }
  if (!data) {
    return (
      <div className="mx-auto max-w-2xl py-10">
        <Link href="/exames" className="text-sm underline-offset-4 hover:underline">
          ← Voltar
        </Link>
        <p className="mt-6 text-muted-foreground">Resultado não encontrado.</p>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <Link href="/exames" className="inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground">
        <ArrowLeft className="h-3.5 w-3.5" /> Meus exames
      </Link>

      <header className="space-y-2">
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <FlaskConical className="h-4 w-4" />
          {formatDate(data.collectionDate, "EEEE, d 'de' MMMM")}
        </div>
        <h1 className="text-2xl font-light">{data.laboratoryName}</h1>
        <div className="flex flex-wrap gap-1.5">
          <Badge variant="outline">{data.status}</Badge>
          {data.requestingDoctor && <Badge variant="outline">por {data.requestingDoctor}</Badge>}
        </div>
      </header>

      {data.observations && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Observações</CardTitle>
          </CardHeader>
          <CardContent className="text-sm text-muted-foreground whitespace-pre-wrap">
            {data.observations}
          </CardContent>
        </Card>
      )}

      <Card>
        <CardHeader>
          <CardTitle className="text-base">Valores ({data.values.length})</CardTitle>
        </CardHeader>
        <CardContent className="p-0">
          {data.values.length === 0 ? (
            <p className="px-6 py-8 text-center text-sm text-muted-foreground">
              Sem valores estruturados. Veja anexos no PDF.
            </p>
          ) : (
            <table className="w-full text-sm">
              <thead className="border-b bg-muted/40 text-xs uppercase tracking-wider text-muted-foreground">
                <tr>
                  <th className="px-4 py-2 text-left">Exame</th>
                  <th className="px-4 py-2 text-right">Valor</th>
                  <th className="px-4 py-2 text-left">Unidade</th>
                </tr>
              </thead>
              <tbody>
                {data.values.map((v) => (
                  <tr key={v.id} className="border-b last:border-0">
                    <td className="px-4 py-2 font-medium">{v.testName}</td>
                    <td className="px-4 py-2 text-right tabular-nums">
                      {v.value != null ? v.value : v.valueText || "—"}
                    </td>
                    <td className="px-4 py-2 text-muted-foreground">{v.unit ?? ""}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
