'use client';

import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import Link from 'next/link';
import { format, isSameDay, isToday, isYesterday } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { ExternalLink, FileText, Image as ImageIcon, Loader2, Mail, MessageSquare, ArrowLeft, CalendarPlus, RefreshCw, Sparkles, MoreVertical } from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '@/lib/utils';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Skeleton } from '@/components/ui/skeleton';
import {
  type ConversationItem,
  type ConversationMessage,
  type ConversationMessageAttachment,
  type ConversationOwnerType,
  avatarColorClass,
  fetchConversationAttachment,
  fetchConversationMedia,
  initials,
  useConversationAISummary,
  useConversationMessages,
  useInterpretExam,
  useSaveToProntuario,
} from '@/lib/api/conversations-api';
import { ConversationComposer } from './conversation-composer';
import { AutomationToggle } from './automation-toggle';

type Props = {
  item: ConversationItem;
  /** Mostra botão "voltar" no header — usado em mobile drawer e no dock. */
  onBack?: () => void;
  /** Trava o canal do thread + compositor (ex.: 'whatsapp' na página/dock dedicados). */
  channel?: 'whatsapp' | 'email';
  /** Agrupa os controles avançados (Resumir/Agendar/Ver ficha/IA) num menu "•••". */
  menuControls?: boolean;
  /** Modo enxuto (dock): paddings menores no compositor. */
  compact?: boolean;
};

function detailHref(item: ConversationItem): string {
  if (item.ownerType === 'patient') return `/patients/${item.ownerId}`;
  return `/leads/${item.ownerId}`;
}

/** Rótulo do separador de dia na timeline (Hoje / Ontem / data). */
function dateDividerLabel(iso: string): string {
  const d = new Date(iso);
  if (isToday(d)) return 'Hoje';
  if (isYesterday(d)) return 'Ontem';
  return format(d, "d 'de' MMMM 'de' yyyy", { locale: ptBR });
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

/** Anexo é candidato a exame (interpretável): PDF ou imagem. */
function isExamCandidate(ct?: string): boolean {
  return ct === 'application/pdf' || ct === 'image/jpeg' || ct === 'image/png';
}

/** Lista de anexos (e-mail) com download autenticado + ações de prontuário (paciente). */
function AttachmentChips({
  attachments,
  tone,
  ownerType,
  ownerId,
  activityId,
}: {
  attachments: ConversationMessageAttachment[];
  tone: 'inbound' | 'outbound';
  ownerType: ConversationOwnerType;
  ownerId: string;
  activityId: string;
}) {
  const interpret = useInterpretExam(ownerType, ownerId);
  const save = useSaveToProntuario(ownerType, ownerId);
  const [busyIdx, setBusyIdx] = useState<number | null>(null);

  const openAttachment = useCallback(
    async (idx: number, filename: string) => {
      try {
        const blob = await fetchConversationAttachment(ownerType, ownerId, activityId, idx);
        const url = URL.createObjectURL(blob);
        window.open(url, '_blank', 'noopener,noreferrer');
        setTimeout(() => URL.revokeObjectURL(url), 60_000);
      } catch {
        toast.error(`Falha ao abrir ${filename}`);
      }
    },
    [ownerType, ownerId, activityId]
  );

  if (!attachments.length) return null;
  // Só conversa de paciente, anexo recebido e arquivo interpretável vão pro prontuário.
  const canProntuario = ownerType === 'patient' && tone === 'inbound';

  return (
    <ul className="mt-2 flex flex-col gap-1.5" aria-label="Anexos">
      {attachments.map((att, i) => {
        const Icon = attachmentIcon(att);
        const sizeLabel = formatBytes(att.size_bytes);
        const examable = canProntuario && isExamCandidate(att.content_type);
        const busy = busyIdx === i && (save.isPending || interpret.isPending);
        return (
          <li key={`${att.path}-${i}`} className="flex flex-wrap items-center gap-1.5">
            <button
              type="button"
              onClick={() => openAttachment(i, att.filename)}
              title={`Abrir ${att.filename}`}
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
            </button>
            {examable && (
              <>
                <Button
                  type="button"
                  size="sm"
                  variant="ghost"
                  className="h-6 gap-1 px-2 text-[10px]"
                  disabled={busy}
                  onClick={() => {
                    setBusyIdx(i);
                    save.mutate(
                      { activityId, attachmentIndex: i },
                      {
                        onSuccess: () => toast.success('Salvo no prontuário.'),
                        onError: (e: unknown) =>
                          toast.error(e instanceof Error ? e.message : 'Falha ao salvar'),
                        onSettled: () => setBusyIdx(null),
                      }
                    );
                  }}
                >
                  <FileText className="h-3 w-3" /> Salvar no prontuário
                </Button>
                <Button
                  type="button"
                  size="sm"
                  variant="ghost"
                  className="h-6 gap-1 px-2 text-[10px]"
                  disabled={busy}
                  onClick={() => {
                    setBusyIdx(i);
                    interpret.mutate(
                      { activityId, attachmentIndex: i },
                      {
                        onSuccess: () =>
                          toast.success('Enviado para interpretação. Acompanhe em Exames.'),
                        onError: (e: unknown) =>
                          toast.error(e instanceof Error ? e.message : 'Falha ao interpretar'),
                        onSettled: () => setBusyIdx(null),
                      }
                    );
                  }}
                >
                  {busy ? <Loader2 className="h-3 w-3 animate-spin" /> : <Sparkles className="h-3 w-3" />}{' '}
                  Interpretar como exame
                </Button>
              </>
            )}
          </li>
        );
      })}
    </ul>
  );
}

function mediaKindLabel(kind: string): string {
  switch (kind) {
    case 'image':
      return 'imagem.jpg';
    case 'sticker':
      return 'figurinha.webp';
    case 'audio':
    case 'voice':
      return 'audio.ogg';
    case 'video':
      return 'video.mp4';
    case 'document':
      return 'documento';
    default:
      return 'arquivo';
  }
}

/** Renderiza a mídia (WhatsApp) de uma mensagem, baixando o blob autenticado. */
function WhatsAppMediaView({
  msg,
  ownerType,
  ownerId,
}: {
  msg: ConversationMessage;
  ownerType: ConversationOwnerType;
  ownerId: string;
}) {
  const [url, setUrl] = useState<string | null>(null);
  const [err, setErr] = useState(false);
  const kind = msg.mediaType ?? '';
  const isImage = kind === 'image' || kind === 'sticker';
  const isAudio = kind === 'audio' || kind === 'voice';

  useEffect(() => {
    let cancelled = false;
    let objectUrl: string | null = null;
    setUrl(null);
    setErr(false);
    fetchConversationMedia(ownerType, ownerId, msg.id)
      .then((blob) => {
        if (cancelled) return;
        objectUrl = URL.createObjectURL(blob);
        setUrl(objectUrl);
      })
      .catch(() => {
        if (!cancelled) setErr(true);
      });
    return () => {
      cancelled = true;
      if (objectUrl) URL.revokeObjectURL(objectUrl);
    };
  }, [ownerType, ownerId, msg.id]);

  const sizeLabel = formatBytes(msg.mediaSizeBytes);
  const filename = msg.mediaFilename || mediaKindLabel(kind);

  if (err) {
    return <p className="mt-2 text-[11px] italic text-rose-600">Falha ao carregar mídia.</p>;
  }

  return (
    <div className="mt-2">
      {isImage ? (
        url ? (
          <a href={url} target="_blank" rel="noopener noreferrer">
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img src={url} alt={filename} className="max-h-64 rounded-md border" />
          </a>
        ) : (
          <Skeleton className="h-40 w-40 rounded-md" />
        )
      ) : isAudio ? (
        url ? (
          <audio controls src={url} className="w-full max-w-[280px]" />
        ) : (
          <Skeleton className="h-10 w-[260px] rounded-md" />
        )
      ) : (
        <a
          href={url ?? undefined}
          download={filename}
          target="_blank"
          rel="noopener noreferrer"
          className="inline-flex max-w-[260px] items-center gap-1.5 rounded-full border border-stone-300 bg-white px-2.5 py-1 text-[11px] font-medium text-stone-800 hover:bg-stone-100"
        >
          {url ? (
            <FileText className="h-3 w-3 shrink-0" aria-hidden />
          ) : (
            <Loader2 className="h-3 w-3 shrink-0 animate-spin" aria-hidden />
          )}
          <span className="truncate">{filename}</span>
          {sizeLabel && <span className="shrink-0 text-muted-foreground">· {sizeLabel}</span>}
        </a>
      )}
      {msg.transcription && (
        <p className="mt-1 rounded bg-white/60 p-1.5 text-[11px] text-stone-700">
          <span className="font-medium">Transcrição:</span> {msg.transcription}
        </p>
      )}
    </div>
  );
}

/** Botão (staff) que envia o documento de prontuário ao interpretador de exames. */
function InterpretExamButton({
  ownerType,
  ownerId,
  activityId,
}: {
  ownerType: ConversationOwnerType;
  ownerId: string;
  activityId: string;
}) {
  const interpret = useInterpretExam(ownerType, ownerId);
  return (
    <Button
      type="button"
      size="sm"
      variant="outline"
      className="mt-2 h-7 gap-1 text-[11px]"
      disabled={interpret.isPending}
      onClick={() =>
        interpret.mutate(
          { activityId },
          {
            onSuccess: () =>
              toast.success('Enviado para interpretação. Acompanhe em Exames do paciente.'),
            onError: (e: unknown) =>
              toast.error(e instanceof Error ? e.message : 'Falha ao interpretar como exame'),
          }
        )
      }
    >
      {interpret.isPending ? (
        <Loader2 className="h-3 w-3 animate-spin" />
      ) : (
        <Sparkles className="h-3 w-3" />
      )}
      Interpretar como exame
    </Button>
  );
}

function MessageBubble({
  msg,
  ownerName,
  ownerType,
  ownerId,
}: {
  msg: ConversationMessage;
  ownerName: string;
  ownerType: ConversationOwnerType;
  ownerId: string;
}) {
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
        {msg.content && <p className="mt-1 whitespace-pre-wrap wrap-break-word">{msg.content}</p>}
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
            {msg.metadata?.origin === 'phone_app' && (
              <span className="shrink-0 rounded bg-emerald-100 px-1 text-[10px] text-emerald-800">
                via app
              </span>
            )}
            {Boolean(msg.metadata?.ai_generated) && (
              <span className="shrink-0 rounded bg-violet-100 px-1 text-[10px] font-medium text-violet-800">
                respondido pela IA
              </span>
            )}
          </span>
          <span className="shrink-0">
            {format(new Date(msg.createdAt), "dd/MM 'às' HH:mm", { locale: ptBR })}
          </span>
        </div>
        {subject && msg.channel === 'email' && (
          <p className="mb-1 text-xs font-semibold text-foreground/90 wrap-break-word">{subject}</p>
        )}
        {msg.content?.trim() ? (
          <p className="whitespace-pre-wrap wrap-break-word">{msg.content}</p>
        ) : !msg.mediaType ? (
          <p className="whitespace-pre-wrap wrap-break-word">
            <span className="italic text-muted-foreground">(corpo vazio)</span>
          </p>
        ) : null}
        {msg.mediaType && (
          <WhatsAppMediaView msg={msg} ownerType={ownerType} ownerId={ownerId} />
        )}
        {msg.patientDocumentId && (
          <div className="mt-2 flex flex-wrap items-center gap-2">
            <Badge variant="secondary" className="gap-1 text-[10px]">
              <FileText className="h-3 w-3" aria-hidden /> Salvo no prontuário
            </Badge>
            {ownerType === 'patient' && (
              <InterpretExamButton ownerType={ownerType} ownerId={ownerId} activityId={msg.id} />
            )}
          </div>
        )}
        {msg.metadata?.attachments && msg.metadata.attachments.length > 0 && (
          <AttachmentChips
            attachments={msg.metadata.attachments}
            tone={isInbound ? 'inbound' : 'outbound'}
            ownerType={ownerType}
            ownerId={ownerId}
            activityId={msg.id}
          />
        )}
      </div>
    </div>
  );
}

/** Estado simples pra o dialog de resumo IA. */
type SummaryDialogState =
  | { open: false }
  | { open: true; summary?: string; generatedAt?: string; cached?: boolean };

/**
 * Renderiza bullets de markdown simples ("- foo"). Mantemos parser minimalista
 * porque o prompt instrui formato bullet — sem necessidade de markdown completo.
 */
function SummaryBullets({ summary }: { summary: string }) {
  const lines = useMemo(
    () =>
      summary
        .split(/\r?\n/)
        .map((l) => l.trim())
        .filter(Boolean),
    [summary]
  );
  const isBulletList = lines.length > 0 && lines.every((l) => /^[-*•]\s+/.test(l));

  if (isBulletList) {
    return (
      <ul className="list-disc space-y-2 pl-5 text-sm leading-relaxed text-foreground">
        {lines.map((l, i) => (
          <li key={i}>{l.replace(/^[-*•]\s+/, '')}</li>
        ))}
      </ul>
    );
  }
  return <p className="whitespace-pre-wrap text-sm leading-relaxed">{summary}</p>;
}

export function ConversationViewer({ item, onBack, channel, menuControls, compact }: Props) {
  const { data: messages, isLoading } = useConversationMessages(item.ownerType, item.ownerId, channel);
  const scrollRef = useRef<HTMLDivElement | null>(null);
  const lastScrollKeyRef = useRef<string | null>(null);

  // ===== IA =====
  const [summaryDialog, setSummaryDialog] = useState<SummaryDialogState>({ open: false });
  /** Texto sugerido pela IA (controlado externo do composer pra prefill). */
  const [draftBody, setDraftBody] = useState<string | null>(null);
  /** Token incrementado a cada nova sugestão pra forçar o composer a aceitar o prefill
   * mesmo que o vendedor já tenha digitado algo (sobrescreve apenas no momento do "set"). */
  const [draftToken, setDraftToken] = useState(0);

  const summary = useConversationAISummary(item.ownerType, item.ownerId);

  const handleSummarize = useCallback(
    (force = false) => {
      setSummaryDialog({ open: true });
      summary.mutate(
        { force },
        {
          onSuccess: (data) => {
            setSummaryDialog({
              open: true,
              summary: data.summary,
              generatedAt: data.generatedAt,
              cached: data.cached,
            });
          },
          onError: (err: unknown) => {
            const msg = err instanceof Error ? err.message : 'Falha ao gerar resumo';
            toast.error(msg);
            setSummaryDialog({ open: false });
          },
        }
      );
    },
    [summary]
  );

  const handleSuggestionInjected = useCallback((text: string) => {
    setDraftBody(text);
    setDraftToken((t) => t + 1);
  }, []);

  // Reset draft ao trocar de conversa.
  useEffect(() => {
    setDraftBody(null);
    setDraftToken(0);
    setSummaryDialog({ open: false });
  }, [item.ownerType, item.ownerId]);

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

  const agendarHref =
    item.ownerType === 'patient'
      ? `/appointments/new?patientId=${item.ownerId}`
      : `/leads/${item.ownerId}`;

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
        {menuControls ? (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                type="button"
                variant="ghost"
                size="icon"
                className="shrink-0"
                aria-label="Mais ações da conversa"
                title="Mais ações"
              >
                <MoreVertical className="h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-60">
              <div className="px-2 py-1.5">
                <p className="mb-1 text-[10px] uppercase tracking-wide text-muted-foreground">
                  Recepção IA
                </p>
                <AutomationToggle ownerType={item.ownerType} ownerId={item.ownerId} />
              </div>
              <DropdownMenuItem
                onSelect={(e) => {
                  e.preventDefault();
                  handleSummarize(false);
                }}
                disabled={summary.isPending}
              >
                <Sparkles className="mr-2 h-4 w-4" /> Resumir conversa
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href={agendarHref}>
                  <CalendarPlus className="mr-2 h-4 w-4" /> Agendar consulta
                </Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link href={detailHref(item)}>
                  <ExternalLink className="mr-2 h-4 w-4" /> Ver ficha completa
                </Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <>
            <AutomationToggle ownerType={item.ownerType} ownerId={item.ownerId} />
            <Button
              type="button"
              variant="ghost"
              size="sm"
              onClick={() => handleSummarize(false)}
              disabled={summary.isPending}
              className="hidden shrink-0 sm:inline-flex"
              aria-label="Resumir conversa com IA"
              title="Resumir conversa com IA"
            >
              {summary.isPending ? (
                <Loader2 className="mr-1 h-3.5 w-3.5 animate-spin" />
              ) : (
                <Sparkles className="mr-1 h-3.5 w-3.5" />
              )}
              Resumir
            </Button>
            <Button asChild variant="outline" size="sm" className="hidden shrink-0 sm:inline-flex">
              <Link href={agendarHref} title="Agendar consulta para este contato">
                <CalendarPlus className="mr-1 h-3.5 w-3.5" /> Agendar
              </Link>
            </Button>
            <Link
              href={detailHref(item)}
              className="hidden shrink-0 items-center gap-1 text-xs text-blue-600 hover:underline sm:inline-flex"
            >
              Ver detalhe <ExternalLink className="h-3 w-3" />
            </Link>
          </>
        )}
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
            {sortedMessages.map((msg, i) => {
              const prev = i > 0 ? sortedMessages[i - 1] : null;
              const showDivider =
                !prev || !isSameDay(new Date(prev.createdAt), new Date(msg.createdAt));
              return (
                <div key={msg.id} className="space-y-3">
                  {showDivider && (
                    <div className="flex justify-center py-1">
                      <span className="rounded-full bg-background/80 px-3 py-0.5 text-[11px] font-medium text-muted-foreground shadow-sm ring-1 ring-border">
                        {dateDividerLabel(msg.createdAt)}
                      </span>
                    </div>
                  )}
                  <MessageBubble
                    msg={msg}
                    ownerName={item.name}
                    ownerType={item.ownerType}
                    ownerId={item.ownerId}
                  />
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* Composer */}
      <ConversationComposer
        item={item}
        draftBody={draftBody}
        draftToken={draftToken}
        onSuggestionApplied={handleSuggestionInjected}
        lockChannel={channel}
        compact={compact}
      />

      {/* Dialog: Resumo IA */}
      <Dialog
        open={summaryDialog.open}
        onOpenChange={(o) => {
          if (!o) setSummaryDialog({ open: false });
        }}
      >
        <DialogContent className="max-w-xl">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Sparkles className="h-4 w-4 text-violet-600" />
              Resumo da conversa
            </DialogTitle>
            <DialogDescription>
              Análise gerada por IA — revise antes de tomar decisões.
            </DialogDescription>
          </DialogHeader>

          <div className="min-h-[120px] rounded-md border border-border bg-muted/30 p-4">
            {(() => {
              const hasSummary = summaryDialog.open && summaryDialog.summary;
              if (summary.isPending && !hasSummary) {
                return (
                  <div className="space-y-2" role="status" aria-live="polite">
                    <p className="text-xs text-muted-foreground">Analisando conversa…</p>
                    <Skeleton className="h-3 w-full" />
                    <Skeleton className="h-3 w-11/12" />
                    <Skeleton className="h-3 w-9/12" />
                    <Skeleton className="h-3 w-10/12" />
                  </div>
                );
              }
              if (summaryDialog.open && summaryDialog.summary) {
                return <SummaryBullets summary={summaryDialog.summary} />;
              }
              return <p className="text-sm text-muted-foreground">Sem conteúdo.</p>;
            })()}
          </div>

          {summaryDialog.open && summaryDialog.generatedAt && (
            <div className="flex flex-wrap items-center gap-2 text-[11px] text-muted-foreground">
              <Badge variant="outline" className="border-violet-300 bg-violet-50 text-violet-900">
                IA · {summaryDialog.cached ? 'cache' : 'novo'}
              </Badge>
              <span>
                gerado às{' '}
                {format(new Date(summaryDialog.generatedAt), "dd/MM 'às' HH:mm", { locale: ptBR })}
              </span>
            </div>
          )}

          <DialogFooter className="gap-2 sm:gap-2">
            <Button
              type="button"
              variant="outline"
              size="sm"
              onClick={() => handleSummarize(true)}
              disabled={summary.isPending}
            >
              <RefreshCw className={cn('mr-1 h-3.5 w-3.5', summary.isPending && 'animate-spin')} />
              Regenerar
            </Button>
            <Button type="button" size="sm" onClick={() => setSummaryDialog({ open: false })}>
              Fechar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
