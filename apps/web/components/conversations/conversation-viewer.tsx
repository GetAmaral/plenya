'use client';

import { useEffect, useMemo, useRef } from 'react';
import Link from 'next/link';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { ExternalLink, FileText, Image as ImageIcon, Mail, MessageSquare, ArrowLeft } from 'lucide-react';

import { cn } from '@/lib/utils';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import {
  type ConversationItem,
  type ConversationMessage,
  type ConversationMessageAttachment,
  attachmentDownloadUrl,
  avatarColorClass,
  initials,
  useConversationMessages,
} from '@/lib/api/conversations-api';
import { ConversationComposer } from './conversation-composer';

type Props = {
  item: ConversationItem;
  /** Mostra botão "voltar" no header — usado em mobile drawer. */
  onBack?: () => void;
};

function detailHref(item: ConversationItem): string {
  if (item.ownerType === 'patient') return `/patients/${item.ownerId}`;
  return `/leads/${item.ownerId}`;
}

/** Formata bytes pra rótulo curto. */
function formatBytes(n?: number): string | null {
  if (!n || n <= 0) return null;
  if (n < 1024) return `${n} B`;
  if (n < 1024 * 1024) return `${(n / 1024).toFixed(0)} KB`;
  return `${(n / (1024 * 1024)).toFixed(1)} MB`;
}

function attachmentIcon(att: ConversationMessageAttachment) {
  if (att.content_type?.startsWith('image/')) return ImageIcon;
  return FileText;
}

/** Lista clicável de anexos abaixo do corpo da mensagem. Usado em inbound + outbound. */
function AttachmentChips({
  attachments,
  tone,
}: {
  attachments: ConversationMessageAttachment[];
  tone: 'inbound' | 'outbound';
}) {
  if (!attachments.length) return null;
  return (
    <ul className="mt-2 flex flex-wrap gap-1.5" aria-label="Anexos">
      {attachments.map((att, i) => {
        const Icon = attachmentIcon(att);
        const url = attachmentDownloadUrl(att.path);
        const sizeLabel = formatBytes(att.size_bytes);
        return (
          <li key={`${att.path}-${i}`}>
            <a
              href={url}
              target="_blank"
              rel="noopener noreferrer"
              title={`Baixar ${att.filename}`}
              className={cn(
                'inline-flex max-w-[260px] items-center gap-1.5 rounded-full border px-2.5 py-1 text-[11px] font-medium transition-colors',
                tone === 'inbound'
                  ? 'border-stone-300 bg-white text-stone-800 hover:bg-stone-100'
                  : 'border-sky-300 bg-white text-sky-900 hover:bg-sky-100'
              )}
            >
              <Icon className="h-3 w-3 shrink-0" aria-hidden />
              <span className="truncate">{att.filename}</span>
              {sizeLabel && (
                <span className="shrink-0 text-muted-foreground">· {sizeLabel}</span>
              )}
            </a>
          </li>
        );
      })}
    </ul>
  );
}

function MessageBubble({ msg, ownerName }: { msg: ConversationMessage; ownerName: string }) {
  const isStatus = msg.type === 'message_status_changed';
  const isInbound = msg.type === 'message_received';
  const isOutbound = msg.type === 'message_sent';

  if (isStatus) {
    const status = msg.metadata?.status ?? '';
    const failed = status === 'failed';
    return (
      <div className="px-2 py-1 text-center text-[11px] text-muted-foreground">
        <span className={cn(failed && 'text-rose-600')}>
          {format(new Date(msg.createdAt), 'HH:mm:ss')} · {status || msg.type}
        </span>
      </div>
    );
  }

  // Mensagens de outros tipos (note_added, status_changed, created, etc.) — render genérico cinza
  if (!isInbound && !isOutbound) {
    return (
      <div className="rounded-md bg-muted/50 p-2 text-xs text-muted-foreground">
        <div className="flex flex-wrap items-center justify-between gap-2">
          <span className="capitalize">{msg.type.replace(/_/g, ' ')}</span>
          <span>{format(new Date(msg.createdAt), "dd/MM 'às' HH:mm", { locale: ptBR })}</span>
        </div>
        {msg.content && <p className="mt-1 whitespace-pre-wrap break-words">{msg.content}</p>}
      </div>
    );
  }

  const ChannelIcon = msg.channel === 'whatsapp' ? MessageSquare : Mail;
  const subject = msg.metadata?.subject;
  const fromLabel = isInbound
    ? msg.metadata?.from || ownerName
    : msg.actor?.name || msg.metadata?.from || 'Plenya';

  return (
    <div className={cn('flex w-full', isOutbound ? 'justify-end' : 'justify-start')}>
      <div
        className={cn(
          'max-w-[85%] rounded-lg border p-3 text-sm shadow-sm md:max-w-[75%]',
          isInbound
            ? 'border-stone-200 bg-stone-50 text-stone-900'
            : 'border-sky-200 bg-sky-50 text-sky-900'
        )}
      >
        <div className="mb-1 flex flex-wrap items-center justify-between gap-x-2 gap-y-1 text-[11px] text-muted-foreground">
          <span className="flex min-w-0 items-center gap-1">
            <ChannelIcon className="h-3 w-3 shrink-0" aria-hidden />
            <strong className="truncate">{fromLabel}</strong>
            <span className="shrink-0">· {isInbound ? 'recebido' : 'enviado'}</span>
          </span>
          <span className="shrink-0">
            {format(new Date(msg.createdAt), "dd/MM 'às' HH:mm", { locale: ptBR })}
          </span>
        </div>
        {subject && msg.channel === 'email' && (
          <p className="mb-1 text-xs font-semibold text-foreground/90 break-words">{subject}</p>
        )}
        <p className="whitespace-pre-wrap break-words">
          {msg.content?.trim() || <span className="italic text-muted-foreground">(corpo vazio)</span>}
        </p>
        {msg.metadata?.attachments && msg.metadata.attachments.length > 0 && (
          <AttachmentChips
            attachments={msg.metadata.attachments}
            tone={isInbound ? 'inbound' : 'outbound'}
          />
        )}
      </div>
    </div>
  );
}

export function ConversationViewer({ item, onBack }: Props) {
  const { data: messages, isLoading } = useConversationMessages(item.ownerType, item.ownerId);
  const scrollRef = useRef<HTMLDivElement | null>(null);
  const lastScrollKeyRef = useRef<string | null>(null);

  // Auto-scroll pro fim:
  //   1. ao trocar de conversa (ownerId mudou)
  //   2. quando chega mensagem nova (count cresce)
  // Não scrolla se usuário está scrollado pra cima lendo histórico antigo? V1: sempre scrolla. Aceitável.
  const messagesCount = messages?.length ?? 0;
  const scrollKey = `${item.ownerType}:${item.ownerId}:${messagesCount}`;

  useEffect(() => {
    if (lastScrollKeyRef.current === scrollKey) return;
    lastScrollKeyRef.current = scrollKey;
    const el = scrollRef.current;
    if (el) {
      // setTimeout pra esperar o paint dos novos itens
      requestAnimationFrame(() => {
        el.scrollTop = el.scrollHeight;
      });
    }
  }, [scrollKey]);

  const sortedMessages = useMemo(() => {
    if (!messages) return [];
    // Backend retorna ASC, mas garantimos defensivamente
    return [...messages].sort(
      (a, b) => new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime()
    );
  }, [messages]);

  return (
    <div className="flex h-full min-h-0 flex-col bg-background">
      {/* Header */}
      <div className="flex items-center gap-3 border-b border-border px-4 py-3">
        {onBack && (
          <button
            type="button"
            onClick={onBack}
            className="flex h-8 w-8 shrink-0 items-center justify-center rounded-md hover:bg-muted md:hidden"
            aria-label="Voltar para a lista"
          >
            <ArrowLeft className="h-4 w-4" />
          </button>
        )}
        <Avatar className="h-10 w-10 shrink-0">
          <AvatarFallback className={cn('text-sm font-semibold', avatarColorClass(item.name))}>
            {initials(item.name)}
          </AvatarFallback>
        </Avatar>
        <div className="min-w-0 flex-1">
          <div className="flex flex-wrap items-center gap-2">
            <h2 className="min-w-0 truncate text-base font-semibold">{item.name}</h2>
            <Badge
              variant="outline"
              className={cn(
                'text-[10px] uppercase tracking-wide',
                item.ownerType === 'patient'
                  ? 'border-emerald-300 bg-emerald-50 text-emerald-900'
                  : 'border-sky-300 bg-sky-50 text-sky-900'
              )}
            >
              {item.ownerType === 'patient' ? 'Paciente' : 'Lead'}
            </Badge>
          </div>
          <div className="mt-0.5 flex flex-wrap gap-x-3 gap-y-1 text-xs text-muted-foreground">
            {item.email && (
              <a
                href={`mailto:${item.email}`}
                className="inline-flex items-center gap-1 hover:text-foreground hover:underline"
              >
                <Mail className="h-3 w-3" /> {item.email}
              </a>
            )}
            {item.phone && (
              <a
                href={`https://wa.me/${item.phone.replace(/\D/g, '')}`}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-1 hover:text-foreground hover:underline"
              >
                <MessageSquare className="h-3 w-3" /> {item.phone}
              </a>
            )}
          </div>
        </div>
        <Link
          href={detailHref(item)}
          className="hidden shrink-0 items-center gap-1 text-xs text-blue-600 hover:underline sm:inline-flex"
        >
          Ver detalhe <ExternalLink className="h-3 w-3" />
        </Link>
      </div>

      {/* Timeline */}
      <div
        ref={scrollRef}
        className="min-h-0 flex-1 overflow-y-auto bg-muted/20 px-4 py-4"
        role="log"
        aria-live="polite"
        aria-label="Mensagens da conversa"
      >
        {isLoading ? (
          <div className="space-y-3">
            <Skeleton className="h-16 w-2/3" />
            <Skeleton className="h-16 w-3/4 ml-auto" />
            <Skeleton className="h-16 w-1/2" />
          </div>
        ) : sortedMessages.length === 0 ? (
          <div className="flex h-full items-center justify-center text-sm text-muted-foreground">
            Nenhuma mensagem ainda.
          </div>
        ) : (
          <div className="space-y-3">
            {sortedMessages.map((msg) => (
              <MessageBubble key={msg.id} msg={msg} ownerName={item.name} />
            ))}
          </div>
        )}
      </div>

      {/* Composer */}
      <ConversationComposer item={item} />
    </div>
  );
}
