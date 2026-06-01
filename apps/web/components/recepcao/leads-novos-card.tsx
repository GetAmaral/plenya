"use client";

import Link from "next/link";
import { UserPlus, Sparkles, ArrowRight } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { useLeads, SOURCE_LABELS } from "@/lib/api/leads-api";
import { EmptyState } from "./empty-state";
import { ErrorState } from "./error-state";
import { formatRelativeDayTime } from "./recepcao-helpers";

/**
 * "Leads novos": leads mais recentes com status=new (cap 5), nome, origem e
 * horario, com link para a lista de Leads. A contagem do cabecalho usa o total
 * do endpoint (todos os leads novos, nao so os 5 exibidos).
 */
export function LeadsNovosCard() {
  // page=0, pageSize=5 — backend ja ordena do mais recente pro mais antigo.
  const { data, isLoading, isError, refetch } = useLeads({ status: "new" }, 0, 5);

  const items = data?.items ?? [];
  const total = data?.total ?? 0;

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <Sparkles className="h-5 w-5 text-amber-600" />
            Leads novos
          </span>
          {!isLoading && total > 0 && (
            <Badge variant="observation">{total}</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {isError ? (
          <ErrorState
            title="Nao foi possivel carregar os leads"
            onRetry={() => refetch()}
          />
        ) : isLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-14 w-full" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <EmptyState
            icon={UserPlus}
            title="Nenhum lead novo"
            hint="Assim que um contato chegar pelo site ou redes, ele aparece aqui."
          />
        ) : (
          <div className="space-y-2">
            {items.map((lead) => (
              <Link
                key={lead.id}
                href={`/leads/${lead.id}`}
                className="flex items-start justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors"
              >
                <div className="min-w-0">
                  <p className="font-medium truncate">
                    {lead.name || lead.email || lead.phone || "Lead sem nome"}
                  </p>
                  <p className="text-sm text-muted-foreground truncate">
                    {SOURCE_LABELS[lead.source] ?? lead.source}
                  </p>
                </div>
                <span className="text-xs text-muted-foreground whitespace-nowrap shrink-0">
                  {formatRelativeDayTime(lead.createdAt)}
                </span>
              </Link>
            ))}
          </div>
        )}

        <Button asChild variant="outline" className="w-full">
          <Link href="/leads">
            Ver todos os leads
            <ArrowRight className="h-4 w-4" />
          </Link>
        </Button>
      </CardContent>
    </Card>
  );
}
