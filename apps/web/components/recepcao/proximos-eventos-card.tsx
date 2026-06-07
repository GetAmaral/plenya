"use client";

import Link from "next/link";
import { CalendarHeart, Cake, ArrowRight } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  DOSSIER_EVENT_TYPE_LABELS,
  useUpcomingEvents,
} from "@/lib/api/conversations-api";
import { EmptyState } from "./empty-state";
import { ErrorState } from "./error-state";

/**
 * "Proximos": aniversarios e eventos de relacionamento (todos os contatos) nos
 * proximos dias, para o time interagir proativamente (Fase C da Livia). Cada item
 * leva para a conversa do contato. So dado social.
 */
function whenLabel(days: number): string {
  if (days <= 0) return "hoje";
  if (days === 1) return "amanha";
  return `em ${days} dias`;
}

export function ProximosEventosCard() {
  const { data, isLoading, isError, refetch } = useUpcomingEvents(14);
  const items = data?.items ?? [];

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <CalendarHeart className="h-5 w-5 text-rose-500" />
            Proximos
          </span>
          {items.length > 0 && <Badge variant="outline">{items.length}</Badge>}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {isError ? (
          <ErrorState
            title="Nao foi possivel carregar os proximos eventos"
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
            icon={Cake}
            title="Nada nos proximos dias"
            hint="Aniversarios e datas importantes vao aparecer aqui."
          />
        ) : (
          <div className="space-y-2">
            {items.map((e) => (
              <Link
                key={e.id}
                href={`/conversas?owner=${e.ownerType}&id=${e.ownerId}`}
                className="flex items-start justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors"
              >
                <div className="min-w-0">
                  <div className="flex items-center gap-2">
                    <span className="font-medium truncate">{e.ownerName || "Contato"}</span>
                    <Badge variant="outline" className="shrink-0">
                      {DOSSIER_EVENT_TYPE_LABELS[e.type] ?? e.type}
                    </Badge>
                  </div>
                  <p className="text-sm text-muted-foreground truncate">{e.title}</p>
                </div>
                <span className="text-xs text-muted-foreground whitespace-nowrap shrink-0">
                  {whenLabel(e.daysUntil)}
                </span>
              </Link>
            ))}
          </div>
        )}

        <Button asChild variant="outline" className="w-full">
          <Link href="/conversas">
            Abrir Central de Conversas
            <ArrowRight className="h-4 w-4" />
          </Link>
        </Button>
      </CardContent>
    </Card>
  );
}
