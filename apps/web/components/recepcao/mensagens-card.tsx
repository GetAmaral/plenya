"use client";

import Link from "next/link";
import { MessageSquare, Inbox, ArrowRight } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  useConversations,
  useConversationsUnreadCount,
} from "@/lib/api/conversations-api";
import { EmptyState } from "./empty-state";
import { ErrorState } from "./error-state";
import { CHANNEL_LABELS, formatRelativeDayTime } from "./recepcao-helpers";

/**
 * "Mensagens": conversas nao lidas. Mostra a contagem total em destaque e uma
 * lista curta das threads mais recentes (nome, canal, trecho, horario), com
 * link para a Central de Conversas.
 *
 * Sem endpoint dedicado de contagem: o hook useConversationsUnreadCount soma
 * unreadCount client-side a partir da lista (?unread_only=true). A lista curta
 * usa o mesmo endpoint via useConversations({ unreadOnly, limit }).
 */
export function MensagensCard() {
  const { data: unreadTotal, isLoading: countLoading } =
    useConversationsUnreadCount();
  const {
    data: list,
    isLoading: listLoading,
    isError: listError,
    refetch,
  } = useConversations({
    unreadOnly: true,
    limit: 5,
  });

  const items = list?.items ?? [];
  const total = unreadTotal ?? 0;

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <MessageSquare className="h-5 w-5 text-ocean-700" />
            Mensagens
          </span>
          {!countLoading && total > 0 && (
            <Badge variant="critical">{total} nao lidas</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {listError ? (
          <ErrorState
            title="Nao foi possivel carregar as mensagens"
            onRetry={() => refetch()}
          />
        ) : listLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-14 w-full" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <EmptyState
            icon={Inbox}
            title="Caixa de entrada em dia"
            hint="Nenhuma conversa aguardando resposta no momento."
          />
        ) : (
          <div className="space-y-2">
            {items.map((c) => (
              <Link
                key={`${c.ownerType}-${c.ownerId}`}
                href="/conversas"
                className="flex items-start justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors"
              >
                <div className="min-w-0">
                  <div className="flex items-center gap-2">
                    <span className="font-medium truncate">{c.name}</span>
                    <Badge variant="outline" className="shrink-0">
                      {CHANNEL_LABELS[c.lastChannel] ?? c.lastChannel}
                    </Badge>
                  </div>
                  {c.lastSnippet && (
                    <p className="text-sm text-muted-foreground truncate">
                      {c.lastSnippet}
                    </p>
                  )}
                </div>
                <div className="flex flex-col items-end gap-1 shrink-0">
                  {c.unreadCount > 0 && (
                    <Badge variant="critical">{c.unreadCount}</Badge>
                  )}
                  {c.lastAt && (
                    <span className="text-xs text-muted-foreground whitespace-nowrap">
                      {formatRelativeDayTime(c.lastAt)}
                    </span>
                  )}
                </div>
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
