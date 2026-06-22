'use client';

import { useState } from 'react';
import { BellOff, Check, Mail } from 'lucide-react';

import { cn } from '@/lib/utils';
import { formatDate } from '@/lib/format-date';
import { Skeleton } from '@/components/ui/skeleton';
import {
  useMarkNotificationRead,
  useNotificationEmails,
} from '@/lib/api/conversations-api';

/**
 * Lista (somente leitura) do bucket "Notificações": e-mails automáticos que não viram Lead.
 * Clicar expande o corpo e marca como lido. Tira o ruído da caixa principal.
 */
export function NotificationsList() {
  const { data, isLoading, isError } = useNotificationEmails();
  const markRead = useMarkNotificationRead();
  const [expanded, setExpanded] = useState<string | null>(null);

  const items = data?.items ?? [];

  const toggle = (id: string, isRead: boolean) => {
    setExpanded((cur) => (cur === id ? null : id));
    if (!isRead) markRead.mutate(id);
  };

  return (
    <div className="flex min-h-0 flex-1 flex-col overflow-hidden rounded-lg border border-border bg-card">
      <div className="flex-1 overflow-y-auto">
        {isLoading ? (
          <div className="space-y-0">
            {Array.from({ length: 4 }).map((_, i) => (
              <div key={i} className="border-b border-border p-3">
                <Skeleton className="mb-2 h-4 w-1/3" />
                <Skeleton className="h-3 w-2/3" />
              </div>
            ))}
          </div>
        ) : isError ? (
          <div className="flex h-full items-center justify-center p-6 text-center text-sm text-rose-700">
            Falha ao carregar notificações.
          </div>
        ) : items.length === 0 ? (
          <div className="flex h-full flex-col items-center justify-center gap-3 p-6 text-center text-sm text-muted-foreground">
            <BellOff className="h-10 w-10 text-muted-foreground/50" aria-hidden />
            <p>Nenhuma notificação automática.</p>
          </div>
        ) : (
          <ul role="list" aria-label="Notificações automáticas">
            {items.map((n) => {
              const open = expanded === n.id;
              return (
                <li key={n.id} className="border-b border-border">
                  <button
                    type="button"
                    onClick={() => toggle(n.id, n.isRead)}
                    className={cn(
                      'flex w-full items-start gap-3 p-3 text-left transition-colors hover:bg-muted/60',
                      !n.isRead && 'bg-amber-50/50'
                    )}
                  >
                    <Mail
                      className={cn(
                        'mt-0.5 h-4 w-4 shrink-0',
                        n.isRead ? 'text-muted-foreground' : 'text-amber-600'
                      )}
                      aria-hidden
                    />
                    <div className="min-w-0 flex-1">
                      <div className="flex items-center justify-between gap-2">
                        <span
                          className={cn(
                            'truncate text-sm',
                            n.isRead ? 'text-muted-foreground' : 'font-semibold text-foreground'
                          )}
                        >
                          {n.fromName || n.fromEmail}
                        </span>
                        <span className="shrink-0 text-[11px] text-muted-foreground">
                          {formatDate(n.receivedAt, "dd/MM HH:mm")}
                        </span>
                      </div>
                      <p className={cn('truncate text-sm', !n.isRead && 'text-foreground/90')}>
                        {n.subject || '(sem assunto)'}
                      </p>
                      {!open && n.bodyText && (
                        <p className="mt-0.5 truncate text-xs text-muted-foreground">{n.bodyText}</p>
                      )}
                      {open && (
                        <div className="mt-2 space-y-1 rounded-md bg-muted/40 p-2 text-xs text-foreground/90">
                          <p className="text-[11px] text-muted-foreground">{n.fromEmail}</p>
                          <p className="whitespace-pre-wrap wrap-break-word">
                            {n.bodyText || '(sem corpo)'}
                          </p>
                        </div>
                      )}
                    </div>
                    {n.isRead && <Check className="mt-0.5 h-3.5 w-3.5 shrink-0 text-muted-foreground" aria-hidden />}
                  </button>
                </li>
              );
            })}
          </ul>
        )}
      </div>
    </div>
  );
}
