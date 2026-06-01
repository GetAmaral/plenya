'use client';

import { Mail, MessageSquare, StickyNote, ArrowDownLeft, ArrowUpRight } from 'lucide-react';
import { formatDistanceToNow } from 'date-fns';
import { ptBR } from 'date-fns/locale';

import { cn } from '@/lib/utils';
import { Badge } from '@/components/ui/badge';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import {
  type ConversationItem,
  avatarColorClass,
  initials,
} from '@/lib/api/conversations-api';

type Props = {
  item: ConversationItem;
  active: boolean;
  onSelect: () => void;
};

function channelIcon(channel: string) {
  if (channel === 'email') return Mail;
  if (channel === 'whatsapp') return MessageSquare;
  return StickyNote;
}

function relativeTime(iso: string): string {
  try {
    return formatDistanceToNow(new Date(iso), { locale: ptBR, addSuffix: false });
  } catch {
    return '';
  }
}

export function ConversationListRow({ item, active, onSelect }: Props) {
  const ChannelIcon = channelIcon(item.lastChannel);
  const DirectionIcon = item.lastDirection === 'in' ? ArrowDownLeft : ArrowUpRight;
  const isUnread = item.unreadCount > 0;
  // SLA: se a última mensagem foi do contato (inbound), estamos devendo resposta.
  const waitingSince =
    item.lastDirection === 'in' ? item.lastInboundAt ?? item.lastAt : null;

  return (
    <button
      type="button"
      onClick={onSelect}
      aria-current={active ? 'true' : undefined}
      aria-label={`Abrir conversa com ${item.name}, ${item.unreadCount} não lidas`}
      className={cn(
        'flex w-full items-start gap-3 border-b border-border px-3 py-3 text-left transition-colors',
        'hover:bg-muted/60 focus:outline-none focus-visible:bg-muted',
        active && 'bg-muted'
      )}
    >
      <Avatar className="h-10 w-10 shrink-0">
        <AvatarFallback className={cn('text-sm font-semibold', avatarColorClass(item.name))}>
          {initials(item.name)}
        </AvatarFallback>
      </Avatar>

      <div className="min-w-0 flex-1">
        <div className="flex items-center gap-2">
          <span
            className={cn(
              'min-w-0 flex-1 truncate text-sm',
              isUnread ? 'font-semibold text-foreground' : 'font-medium text-foreground/90'
            )}
            title={item.name}
          >
            {item.name}
          </span>
          <span className="shrink-0 text-[11px] text-muted-foreground" title={item.lastAt}>
            {relativeTime(item.lastAt)}
          </span>
        </div>

        <div className="mt-0.5 flex items-center gap-1.5 text-xs text-muted-foreground">
          <ChannelIcon className="h-3 w-3 shrink-0" aria-hidden />
          <DirectionIcon
            className="h-3 w-3 shrink-0"
            aria-label={item.lastDirection === 'in' ? 'recebida' : 'enviada'}
          />
          <span
            className={cn(
              'min-w-0 flex-1 truncate',
              isUnread && 'text-foreground/80'
            )}
            title={item.lastSnippet}
          >
            {item.lastSnippet || '(sem conteúdo)'}
          </span>
        </div>

        <div className="mt-1 flex flex-wrap items-center gap-1">
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
          {waitingSince && (
            <Badge
              variant="outline"
              className="text-[10px] border-amber-300 bg-amber-50 text-amber-900"
              title={`Aguardando resposta desde ${waitingSince}`}
            >
              Aguardando {relativeTime(waitingSince)}
            </Badge>
          )}
          {item.channels.includes('email') && (
            <Badge variant="secondary" className="text-[10px]">Email</Badge>
          )}
          {item.channels.includes('whatsapp') && (
            <Badge variant="secondary" className="text-[10px]">WhatsApp</Badge>
          )}
        </div>
      </div>

      {isUnread && (
        <span
          className="mt-1 inline-flex h-5 min-w-[20px] shrink-0 items-center justify-center rounded-full bg-rose-600 px-1.5 text-[11px] font-semibold text-white"
          aria-label={`${item.unreadCount} mensagens não lidas`}
        >
          {item.unreadCount > 99 ? '99+' : item.unreadCount}
        </span>
      )}
    </button>
  );
}
