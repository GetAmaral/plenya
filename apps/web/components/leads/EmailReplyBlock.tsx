'use client';

import { useMemo, useState } from 'react';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { Mail, Plus, Reply } from 'lucide-react';
import { toast } from 'sonner';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  useSendLeadEmailReply,
  type Lead,
  type LeadActivity,
} from '@/lib/api/leads-api';

type Props = {
  lead: Lead;
  /** Email FROM exibido no hint (default contato@plenyasaude.com.br). */
  fromAddress?: string;
};

type EmailMeta = {
  subject?: string;
  message_id?: string;
  in_reply_to?: string;
  references?: string[];
  recipient?: string;
  source?: string;
  resend_id?: string;
  mirrored?: boolean;
};

type ThreadMessage = {
  id: string;
  direction: 'in' | 'out';
  from: string;
  subject: string;
  content: string;
  createdAt: string;
  messageId?: string;
};

type Thread = {
  key: string;
  subject: string;
  messages: ThreadMessage[];
};

const DEFAULT_FROM = 'contato@plenyasaude.com.br';

/**
 * Normaliza subject pra agrupar threads: strip "Re:"/"Res:"/"Fwd:" repetidos + lowercase.
 * Threads que compartilham o mesmo subject normalizado entram juntas.
 */
function normalizeSubject(s: string | undefined): string {
  if (!s) return '(sem assunto)';
  let cur = s.trim();
  // strip prefixes encadeados ("Re: Re: Fwd: foo" → "foo")
  let changed = true;
  while (changed) {
    changed = false;
    const lower = cur.toLowerCase();
    for (const p of ['re:', 'res:', 'fwd:', 'fw:', 'enc:']) {
      if (lower.startsWith(p)) {
        cur = cur.slice(p.length).trim();
        changed = true;
        break;
      }
    }
  }
  return cur.toLowerCase() || '(sem assunto)';
}

function asMeta(m: LeadActivity['metadata']): EmailMeta {
  return (m as EmailMeta | undefined) ?? {};
}

function buildThreads(activities: LeadActivity[] | undefined, leadEmail?: string): Thread[] {
  if (!activities || activities.length === 0) return [];
  const emailActs = activities.filter(
    (a) => a.channel === 'email' && (a.type === 'message_sent' || a.type === 'message_received')
  );
  if (emailActs.length === 0) return [];

  const buckets = new Map<string, ThreadMessage[]>();
  const subjectByKey = new Map<string, string>();

  for (const a of emailActs) {
    const meta = asMeta(a.metadata);
    const subject = meta.subject ?? '(sem assunto)';
    const key = normalizeSubject(subject);

    const direction: 'in' | 'out' = a.type === 'message_received' ? 'in' : 'out';
    const from =
      direction === 'in'
        ? leadEmail ?? '(lead)'
        : a.actor?.email ?? a.actor?.name ?? 'Plenya';

    const msg: ThreadMessage = {
      id: a.id,
      direction,
      from,
      subject,
      content: a.content ?? '',
      createdAt: a.createdAt,
      messageId: meta.message_id,
    };
    const arr = buckets.get(key);
    if (arr) {
      arr.push(msg);
    } else {
      buckets.set(key, [msg]);
      subjectByKey.set(key, subject);
    }
  }

  // Ordena: dentro da thread cronológico ascendente. Threads ordenadas por última msg desc.
  const threads: Thread[] = Array.from(buckets.entries()).map(([key, msgs]) => {
    const sorted = [...msgs].sort(
      (a, b) => new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime()
    );
    return { key, subject: subjectByKey.get(key) ?? key, messages: sorted };
  });
  threads.sort((a, b) => {
    const lastA = new Date(a.messages[a.messages.length - 1].createdAt).getTime();
    const lastB = new Date(b.messages[b.messages.length - 1].createdAt).getTime();
    return lastB - lastA;
  });
  return threads;
}

function snippet(content: string, max = 100): string {
  const trimmed = content.replace(/\s+/g, ' ').trim();
  if (trimmed.length <= max) return trimmed;
  return trimmed.slice(0, max - 1) + '…';
}

/**
 * Bloco de email no detail do Lead — lista threads agrupadas por subject e permite responder.
 *
 * Renderiza null se o lead não tem email ou não tem opt-in (UX: bloco invisível em vez de
 * aparecer "desabilitado" — admin já vê os badges no card de identificação acima).
 */
export function EmailReplyBlock({ lead, fromAddress = DEFAULT_FROM }: Props) {
  const sendEmail = useSendLeadEmailReply(lead.id);
  const [openReplyFor, setOpenReplyFor] = useState<Thread | null>(null);
  const [composeOpen, setComposeOpen] = useState(false);
  const [replySubject, setReplySubject] = useState('');
  const [replyBody, setReplyBody] = useState('');
  const isComposing = composeOpen && openReplyFor === null;

  const threads = useMemo(() => buildThreads(lead.activities, lead.email), [lead.activities, lead.email]);
  const totalMessages = threads.reduce((acc, t) => acc + t.messages.length, 0);

  if (!lead.email || !lead.emailOptIn) {
    return null;
  }

  const openReply = (thread: Thread | null) => {
    const baseSubject = thread?.subject ?? '';
    const prefilled =
      baseSubject && !/^re:/i.test(baseSubject) ? `Re: ${baseSubject}` : baseSubject;
    setReplySubject(prefilled);
    setReplyBody('');
    setOpenReplyFor(thread);
  };

  const openCompose = () => {
    setReplySubject('');
    setReplyBody('');
    setOpenReplyFor(null);
    setComposeOpen(true);
  };

  const closeReply = () => {
    setOpenReplyFor(null);
    setComposeOpen(false);
    setReplySubject('');
    setReplyBody('');
  };

  const handleSend = () => {
    const body = replyBody.trim();
    if (!body) {
      toast.error('Escreva uma mensagem antes de enviar.');
      return;
    }
    if (isComposing && !replySubject.trim()) {
      toast.error('Defina um assunto para o novo email.');
      return;
    }
    // Pega a última mensagem da thread pra threading (In-Reply-To + References).
    // Em modo compose (novo email), não envia headers de threading.
    let inReplyTo: string | undefined;
    let references: string[] | undefined;
    if (openReplyFor && openReplyFor.messages.length > 0) {
      const last = openReplyFor.messages[openReplyFor.messages.length - 1];
      if (last.messageId) {
        inReplyTo = last.messageId;
        references = openReplyFor.messages
          .map((m) => m.messageId)
          .filter((v): v is string => !!v);
      }
    }

    sendEmail.mutate(
      {
        subject: replySubject.trim() || undefined,
        bodyText: body,
        inReplyTo,
        references,
      },
      {
        onSuccess: () => {
          toast.success('Email enviado.');
          closeReply();
        },
        onError: (err: unknown) => {
          const msg = err instanceof Error ? err.message : 'Falha ao enviar email';
          toast.error(msg);
        },
      }
    );
  };

  return (
    <>
      <Card>
        <CardHeader>
          <div className="flex flex-wrap items-center justify-between gap-2">
            <CardTitle className="flex min-w-0 items-center gap-2 text-base">
              <Mail className="h-4 w-4 shrink-0" /> Email
              <Badge variant="secondary" className="ml-2 shrink-0 text-xs">
                {totalMessages} {totalMessages === 1 ? 'mensagem' : 'mensagens'}
              </Badge>
            </CardTitle>
            <div className="flex flex-wrap gap-2">
              {threads.length > 0 && (
                <Button size="sm" variant="outline" onClick={() => openReply(threads[0])}>
                  <Reply className="mr-1 h-4 w-4" /> Responder
                </Button>
              )}
              <Button size="sm" variant="default" onClick={openCompose}>
                <Plus className="mr-1 h-4 w-4" /> Novo email
              </Button>
            </div>
          </div>
          <p className="text-sm text-muted-foreground">
            Enviado de <strong>{fromAddress}</strong>. Histórico aparece também no Mail (iPhone/Outlook).
          </p>
        </CardHeader>
        <CardContent>
          {threads.length === 0 ? (
            <p className="py-4 text-center text-sm text-muted-foreground">
              Nenhuma troca de email ainda. Clique em <strong>Novo email</strong> para iniciar.
            </p>
          ) : (
            <ol className="space-y-4">
              {threads.map((thread) => (
                <ThreadItem
                  key={thread.key}
                  thread={thread}
                  onReply={() => openReply(thread)}
                />
              ))}
            </ol>
          )}
        </CardContent>
      </Card>

      <Dialog open={openReplyFor !== null || composeOpen} onOpenChange={(open) => !open && closeReply()}>
        <DialogContent className="sm:max-w-2xl">
          <DialogHeader>
            <DialogTitle>{isComposing ? 'Novo email' : 'Responder email'}</DialogTitle>
            <DialogDescription>
              Para: <strong>{lead.email}</strong> · De <strong>{fromAddress}</strong>
            </DialogDescription>
          </DialogHeader>

          <div className="space-y-3">
            <label className="block">
              <span className="text-xs uppercase tracking-wide text-muted-foreground">
                Assunto {isComposing && <span className="text-red-500">*</span>}
              </span>
              <Input
                value={replySubject}
                onChange={(e) => setReplySubject(e.target.value)}
                placeholder={isComposing ? 'Assunto do email' : "(opcional — backend monta 'Re: ...' automaticamente)"}
                className="mt-1"
              />
            </label>
            <label className="block">
              <span className="text-xs uppercase tracking-wide text-muted-foreground">Mensagem</span>
              <textarea
                value={replyBody}
                onChange={(e) => setReplyBody(e.target.value)}
                placeholder="Escreva sua resposta…"
                rows={10}
                className="mt-1 w-full rounded-md border border-input bg-transparent p-3 text-sm focus:outline-none focus:ring-1 focus:ring-ring"
              />
              <p className="mt-1 text-xs text-muted-foreground">
                Texto puro — quebras de linha são preservadas. HTML é gerado automaticamente.
              </p>
            </label>
          </div>

          <DialogFooter>
            <Button variant="outline" onClick={closeReply} disabled={sendEmail.isPending}>
              Cancelar
            </Button>
            <Button onClick={handleSend} disabled={sendEmail.isPending || !replyBody.trim()}>
              {sendEmail.isPending ? 'Enviando…' : 'Enviar email'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  );
}

function ThreadItem({ thread, onReply }: { thread: Thread; onReply: () => void }) {
  const [expanded, setExpanded] = useState(false);
  const last = thread.messages[thread.messages.length - 1];

  return (
    <li className="rounded-md border">
      <button
        type="button"
        onClick={() => setExpanded((v) => !v)}
        className="flex w-full items-start justify-between gap-3 p-3 text-left transition hover:bg-muted/40"
      >
        <div className="min-w-0 flex-1">
          <div className="flex items-center gap-2 text-sm font-medium">
            <span className="truncate">{thread.subject}</span>
            <Badge variant="outline" className="shrink-0 text-xs">
              {thread.messages.length}
            </Badge>
          </div>
          <p className="mt-0.5 truncate text-xs text-muted-foreground">
            <span className="font-medium">{last.from}</span> · {snippet(last.content)}
          </p>
        </div>
        <span className="shrink-0 text-xs text-muted-foreground">
          {format(new Date(last.createdAt), "dd/MM HH:mm", { locale: ptBR })}
        </span>
      </button>
      {expanded && (
        <div className="space-y-2 border-t p-3">
          {thread.messages.map((m) => (
            <div
              key={m.id}
              className={`rounded-md p-3 text-sm ${
                m.direction === 'in'
                  ? 'bg-stone-50 border border-stone-200'
                  : 'bg-sky-50 border border-sky-200'
              }`}
            >
              <div className="mb-1 flex flex-wrap items-center justify-between gap-x-2 gap-y-1 text-xs text-muted-foreground">
                <span className="min-w-0 break-all">
                  <strong>{m.from}</strong> · {m.direction === 'in' ? 'recebido' : 'enviado'}
                </span>
                <span className="shrink-0">{format(new Date(m.createdAt), "dd/MM 'às' HH:mm", { locale: ptBR })}</span>
              </div>
              <p className="whitespace-pre-wrap wrap-break-word">{m.content || '(corpo vazio)'}</p>
            </div>
          ))}
          <div className="flex justify-end">
            <Button size="sm" variant="outline" onClick={onReply}>
              <Reply className="mr-1 h-4 w-4" /> Responder esta thread
            </Button>
          </div>
        </div>
      )}
    </li>
  );
}
