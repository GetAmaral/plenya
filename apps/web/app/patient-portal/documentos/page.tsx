"use client";

/**
 * /documentos — atestados, declarações, relatórios e outros docs clínicos
 * que a equipe Plenya disponibiliza pra você.
 */
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Loader2, FileText, Download } from "lucide-react";
import { toast } from "sonner";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyDocuments,
  patientDocumentsApi,
  type DocumentType,
  type PatientDocumentView,
} from "@/lib/api/patient-portal-api";
import { useAuthStore } from "@/lib/auth-store";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";

const TYPE_LABEL: Record<DocumentType, string> = {
  certificate: "Atestado",
  report: "Relatório",
  referral: "Encaminhamento",
  declaration: "Declaração",
  other: "Documento",
};

export default function MyDocumentsPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useMyDocuments();

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  const handleDownload = (doc: PatientDocumentView) => {
    const token = useAuthStore.getState().accessToken;
    if (!token) {
      toast.error("Sessão expirou");
      return;
    }
    fetch(patientDocumentsApi.downloadURL(doc.id), {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then(async (r) => {
        if (!r.ok) throw new Error("Falha ao baixar");
        const blob = await r.blob();
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.href = url;
        link.download = doc.fileName;
        link.click();
        URL.revokeObjectURL(url);
      })
      .catch((e) => toast.error(e?.message ?? "Falha"));
  };

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Documentos</p>
        <h1 className="text-3xl font-light">Meus documentos</h1>
        <p className="text-muted-foreground">
          Atestados, declarações e relatórios disponibilizados pela equipe Plenya.
        </p>
      </header>

      {!data?.length ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            <FileText className="mx-auto mb-3 h-8 w-8" />
            Nenhum documento por aqui ainda.
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-3">
          {data.map((d) => (
            <Card key={d.id}>
              <CardContent className="flex items-center justify-between gap-3 py-4">
                <div className="min-w-0 flex-1 space-y-1">
                  <div className="flex items-center gap-2 text-xs text-muted-foreground">
                    <FileText className="h-3 w-3" />
                    {format(new Date(d.issuedAt), "dd 'de' MMM 'de' yyyy", { locale: ptBR })}
                  </div>
                  <p className="truncate font-medium">{d.title}</p>
                  {d.description && (
                    <p className="line-clamp-2 text-sm text-muted-foreground">
                      {d.description}
                    </p>
                  )}
                  <div className="flex flex-wrap gap-1.5">
                    <Badge variant="secondary">{TYPE_LABEL[d.type] ?? d.type}</Badge>
                    {d.uploadedByUser?.name && (
                      <Badge variant="outline">por {d.uploadedByUser.name}</Badge>
                    )}
                  </div>
                </div>
                <Button variant="outline" size="sm" onClick={() => handleDownload(d)}>
                  <Download className="mr-1 h-4 w-4" /> Baixar
                </Button>
              </CardContent>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
