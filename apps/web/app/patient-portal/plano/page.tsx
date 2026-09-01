"use client";

/**
 * /plano — a devolutiva dos seus exames, para reler em casa.
 *
 * É a MESMA leitura que o médico fez na consulta, no mesmo desenho, mas montada para a tela: o PDF
 * existe para imprimir e para guardar, e fica ao lado como download. Quem abre isto no celular
 * precisa de texto que se lê sem beliscar a tela, não de um slide de projetor reduzido.
 */
import { useState } from "react";
import { Loader2, FileText, Printer, ChevronLeft } from "lucide-react";
import { formatDate } from "@/lib/format-date";
import { toast } from "sonner";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { useMyPlans, useMyPlan, patientDocumentsApi } from "@/lib/api/patient-portal-api";
import { useAuthStore } from "@/lib/auth-store";
import { PlanDeck } from "@/components/patient-portal/plan-deck";
import type { DeckSlide } from "@/lib/api/patient-plans";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";

export default function MyPlanPage() {
  const { ready } = useRequirePatientAuth();
  const { data: plans = [], isLoading } = useMyPlans();
  const [openId, setOpenId] = useState<string | null>(null);
  const { data: plan, isLoading: loadingPlan } = useMyPlan(openId ?? undefined);

  const baixar = (docId: string, nome: string) => {
    const token = useAuthStore.getState().accessToken;
    if (!token) {
      toast.error("Sessão expirou");
      return;
    }
    fetch(patientDocumentsApi.downloadURL(docId), { headers: { Authorization: `Bearer ${token}` } })
      .then(async (r) => {
        if (!r.ok) throw new Error("Falha ao baixar");
        const url = URL.createObjectURL(await r.blob());
        const link = document.createElement("a");
        link.href = url;
        link.download = nome;
        link.click();
        URL.revokeObjectURL(url);
      })
      .catch((e) => toast.error(e instanceof Error ? e.message : "Falha ao baixar"));
  };

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  // Um plano aberto: a leitura inteira.
  if (openId) {
    return (
      <div className="mx-auto max-w-4xl space-y-4 px-3 py-6 sm:px-6">
        <div className="flex flex-wrap items-center gap-2">
          <Button variant="ghost" size="sm" onClick={() => setOpenId(null)}>
            <ChevronLeft className="mr-1 h-4 w-4" />
            Meus planos
          </Button>
          {plan?.document16x9Id && (
            <Button
              variant="outline"
              size="sm"
              className="ml-auto"
              onClick={() => baixar(plan.document16x9Id!, `${plan.title}.pdf`)}
            >
              <FileText className="mr-2 h-4 w-4" />
              PDF
            </Button>
          )}
          {plan?.documentA4Id && (
            <Button
              variant="outline"
              size="sm"
              onClick={() => baixar(plan.documentA4Id!, `${plan.title} (impressão).pdf`)}
            >
              <Printer className="mr-2 h-4 w-4" />
              Imprimir
            </Button>
          )}
        </div>

        {loadingPlan ? (
          <div className="flex h-[40vh] items-center justify-center">
            <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
          </div>
        ) : plan ? (
          <PlanDeck slides={(plan.content ?? []) as DeckSlide[]} />
        ) : (
          <p className="text-sm text-muted-foreground">Plano não encontrado.</p>
        )}
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-4 px-3 py-6 sm:px-6">
      <div>
        <h1 className="text-xl font-medium">Meu plano</h1>
        <p className="text-sm text-muted-foreground">
          A leitura dos seus exames e o que combinamos fazer.
        </p>
      </div>

      {plans.length === 0 && (
        <Card>
          <CardContent className="py-12 text-center text-sm text-muted-foreground">
            Você ainda não tem um plano publicado. Ele aparece aqui depois da consulta de devolutiva.
          </CardContent>
        </Card>
      )}

      {plans.map((p) => (
        <Card key={p.id} className="cursor-pointer transition-colors hover:bg-accent/40">
          <CardContent className="flex items-center justify-between gap-4 py-5">
            <button className="min-w-0 flex-1 text-left" onClick={() => setOpenId(p.id)}>
              <div className="flex items-center gap-2">
                <span className="truncate font-medium">{p.title}</span>
                {p.version > 1 && <Badge variant="secondary">v{p.version}</Badge>}
              </div>
              {p.publishedAt && (
                <div className="mt-1 text-xs text-muted-foreground">
                  de {formatDate(p.publishedAt)}
                </div>
              )}
            </button>
            <Button variant="ghost" size="sm" onClick={() => setOpenId(p.id)}>
              Abrir
            </Button>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
