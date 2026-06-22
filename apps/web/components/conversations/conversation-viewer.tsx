'use client';

import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import Link from 'next/link';
import { isSameDay, isToday, isYesterday } from 'date-fns';
import { ExternalLink, FileText, Image as ImageIcon, Loader2, Mail, MessageSquare, ArrowLeft, CalendarPlus, RefreshCw, Sparkles, MoreVertical, Bot, Hand, SendHorizonal, Check, CheckCheck, IdCard } from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '@/lib/utils';
import { formatDate } from '@/lib/format-date';
import { safeUrl } from '@/lib/security';
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
  fetchConversationMediaStreamUrl,
  initials,
  useConversationAISummary,
  useConversationMessages,
  useInterpretExam,
  useReceptionSettings,
  useSaveToProntuario,
  useSendConversationWhatsApp,
  useSetConversationAutomation,
  useSuggestedReply,
  type SuggestedReply,
} from '@/lib/api/conversations-api';
import { ConversationComposer } from './conversation-composer';
import { AutomationToggle } from './automation-toggle';
import { DossierPanel } from './dossier-panel';

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
  return formatDate(iso, "d 'de' MMMM 'de' yyyy");
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
    if (isAudio) {
      // Áudio (e vídeo) precisa de URL direta com Range: o Safari/iOS não toca
      // mídia de blob: URL. Busca a URL de streaming assinada e aponta no <audio>.
      fetchConversationMediaStreamUrl(ownerType, ownerId, msg.id)
        .then((streamUrl) => {
          if (!cancelled) setUrl(streamUrl);
        })
        .catch(() => {
          if (!cancelled) setErr(true);
        });
    } else {
      // Imagem/documento: blob autenticado (funciona em todo navegador).
      fetchConversationMedia(ownerType, ownerId, msg.id)
        .then((blob) => {
          if (cancelled) return;
          objectUrl = URL.createObjectURL(blob);
          setUrl(objectUrl);
        })
        .catch(() => {
          if (!cancelled) setErr(true);
        });
    }
    return () => {
      cancelled = true;
      if (objectUrl) URL.revokeObjectURL(objectUrl);
    };
  }, [ownerType, ownerId, msg.id, isAudio]);

  const sizeLabel = formatBytes(msg.mediaSizeBytes);
  const filename = msg.mediaFilename || mediaKindLabel(kind);

  return (
    <div className="mt-2">
      {err ? (
        // Mídia não carregou (ex: arquivo ausente no storage). Ainda assim a
        // transcrição abaixo deve aparecer — é o fallback de leitura do áudio.
        <p className="text-[11px] italic text-rose-600">Falha ao carregar mídia.</p>
      ) : isImage ? (
        url ? (
          <a href={safeUrl(url)} target="_blank" rel="noopener noreferrer">
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img src={safeUrl(url)} alt={filename} className="max-h-64 rounded-md border" />
          </a>
        ) : (
          <Skeleton className="h-40 w-40 rounded-md" />
        )
      ) : isAudio ? (
        url ? (
          <audio controls preload="metadata" src={safeUrl(url)} className="w-full max-w-[280px]" />
        ) : (
          <Skeleton className="h-10 w-[260px] rounded-md" />
        )
      ) : (
        <a
          href={safeUrl(url)}
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

/** Ticks de entrega (estilo WhatsApp) — dobram os eventos sent/delivered/read num só ícone. */
function DeliveryTicks({ status }: { status: string }) {
  if (status === 'failed') {
    return <span className="text-rose-600" title="Falha no envio">· falhou</span>;
  }
  if (status === 'read') {
    return <CheckCheck className="h-3 w-3 text-sky-500" aria-label="Lida" />;
  }
  if (status === 'delivered') {
    return <CheckCheck className="h-3 w-3" aria-label="Entregue" />;
  }
  if (status === 'sent') {
    return <Check className="h-3 w-3" aria-label="Enviada" />;
  }
  return null;
}

function MessageBubble({
  msg,
  ownerName,
  ownerType,
  ownerId,
  deliveryStatus,
}: {
  msg: ConversationMessage;
  ownerName: string;
  ownerType: ConversationOwnerType;
  ownerId: string;
  deliveryStatus?: string;
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
          {formatDate(msg.createdAt, 'HH:mm:ss')} · {status || msg.type}
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
          <span>{formatDate(msg.createdAt, "dd/MM 'às' HH:mm")}</span>
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
          <span className="inline-flex shrink-0 items-center gap-1">
            {formatDate(msg.createdAt, "dd/MM 'às' HH:mm")}
            {isOutbound && msg.channel === 'whatsapp' && deliveryStatus && (
              <DeliveryTicks status={deliveryStatus} />
            )}
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

/**
 * AIDraftBar — barra acima do compositor quando há rascunho do servidor (Atendimento IA).
 * No modo Auto mostra o countdown X/3 (a IA envia ao zerar) + ações "Assumir" e "Enviar agora".
 * No Copiloto é só um aviso de que o rascunho foi preenchido para revisão.
 */
function AIDraftBar({
  draft,
  previewSeconds,
  onAssume,
  onSendNow,
  assuming,
  sending,
}: {
  draft: SuggestedReply;
  previewSeconds: number;
  onAssume: () => void;
  onSendNow: () => void;
  assuming: boolean;
  sending: boolean;
}) {
  const isAuto = draft.effectiveMode === 'auto';
  const isHandoff = draft.action === 'handoff';

  // Auto envia X2 (janela de preview) segundos depois do rascunho ficar pronto. O rascunho foi
  // salvo ~em updatedAt, logo o envio acontece ~em updatedAt + X2. Countdown ao vivo até lá.
  const deadline = useMemo(
    () => new Date(draft.updatedAt).getTime() + previewSeconds * 1000,
    [draft.updatedAt, previewSeconds]
  );
  const [now, setNow] = useState(() => Date.now());
  useEffect(() => {
    if (!isAuto) return;
    const t = setInterval(() => setNow(Date.now()), 1000);
    return () => clearInterval(t);
  }, [isAuto, draft.updatedAt]);

  const remainingSec = Math.max(0, Math.ceil((deadline - now) / 1000));

  if (isHandoff) {
    return (
      <div className="flex items-center gap-2 border-t border-amber-200 bg-amber-50 px-3 py-2 text-xs text-amber-900">
        <Hand className="h-3.5 w-3.5 shrink-0" />
        <span>A IA sugere passar este atendimento para um humano. O rascunho foi preenchido como cortesia.</span>
      </div>
    );
  }

  if (!isAuto) {
    // Copiloto (ou qualquer outro modo com rascunho): só aviso, sem auto-envio.
    return (
      <div className="flex items-center gap-2 border-t border-violet-200 bg-violet-50 px-3 py-2 text-xs text-violet-900">
        <Bot className="h-3.5 w-3.5 shrink-0" />
        <span>Rascunho da IA preenchido no compositor. Revise e envie.</span>
      </div>
    );
  }

  return (
    <div className="flex flex-wrap items-center gap-2 border-t border-violet-200 bg-violet-50 px-3 py-2 text-xs text-violet-900">
      <Bot className="h-3.5 w-3.5 shrink-0" />
      <span className="font-medium">
        {remainingSec > 0
          ? `A IA envia em ${remainingSec}s`
          : 'A IA está enviando…'}
      </span>
      <span className="text-violet-700">Revise o texto no compositor ou intervenha.</span>
      <div className="ml-auto flex items-center gap-1.5">
        <Button
          type="button"
          size="sm"
          variant="outline"
          className="h-7 gap-1 px-2 text-[11px]"
          disabled={assuming}
          onClick={onAssume}
          title="Assumir a conversa (a IA para de enviar sozinha)"
        >
          {assuming ? <Loader2 className="h-3 w-3 animate-spin" /> : <Hand className="h-3 w-3" />}
          Assumir
        </Button>
        <Button
          type="button"
          size="sm"
          className="h-7 gap-1 px-2 text-[11px]"
          disabled={sending}
          onClick={onSendNow}
          title="Enviar o rascunho agora"
        >
          {sending ? <Loader2 className="h-3 w-3 animate-spin" /> : <SendHorizonal className="h-3 w-3" />}
          Enviar agora
        </Button>
      </div>
    </div>
  );
}

export function ConversationViewer({ item, onBack, channel, menuControls, compact }: Props) {
  const { data: messages, isLoading } = useConversationMessages(item.ownerType, item.ownerId, channel);
  const scrollRef = useRef<HTMLDivElement | null>(null);
  const lastScrollKeyRef = useRef<string | null>(null);

  // ===== IA =====
  const [summaryDialog, setSummaryDialog] = useState<SummaryDialogState>({ open: false });
  const [dossierOpen, setDossierOpen] = useState(false);
  /** Texto sugerido pela IA (controlado externo do composer pra prefill). */
  const [draftBody, setDraftBody] = useState<string | null>(null);
  /** Token incrementado a cada nova sugestão pra forçar o composer a aceitar o prefill
   * mesmo que o vendedor já tenha digitado algo (sobrescreve apenas no momento do "set"). */
  const [draftToken, setDraftToken] = useState(0);

  const summary = useConversationAISummary(item.ownerType, item.ownerId);

  // ===== Rascunho do servidor (Atendimento IA) — só WhatsApp =====
  const isWhatsApp = channel !== 'email';
  const suggested = useSuggestedReply(
    isWhatsApp ? item.ownerType : null,
    isWhatsApp ? item.ownerId : null
  );
  const receptionSettings = useReceptionSettings();
  const setAutomation = useSetConversationAutomation(item.ownerType, item.ownerId);
  const sendWa = useSendConversationWhatsApp(item.ownerType, item.ownerId);
  // updatedAt do último rascunho já injetado no compositor (evita reinjetar a cada poll).
  const lastInjectedDraftRef = useRef<string | null>(null);
  // Marca se havia um rascunho do servidor visível — quando ele some (enviado pelo Auto), limpamos
  // o compositor pra não correr o risco de o atendente reenviar e duplicar.
  const hadServerDraftRef = useRef(false);

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
    setDossierOpen(false);
    lastInjectedDraftRef.current = null;
    hadServerDraftRef.current = false;
  }, [item.ownerType, item.ownerId]);

  // Sincroniza o compositor com o rascunho do servidor:
  //  - aparece/muda (novo inbound) → injeta no campo (preview).
  //  - some (o Auto enviou e limpou o rascunho) → ESVAZIA o campo, pra não reenviar/duplicar.
  useEffect(() => {
    const d = suggested.data;
    if (d?.reply) {
      hadServerDraftRef.current = true;
      if (lastInjectedDraftRef.current !== d.updatedAt) {
        lastInjectedDraftRef.current = d.updatedAt;
        setDraftBody(d.reply);
        setDraftToken((t) => t + 1);
      }
    } else if (hadServerDraftRef.current) {
      hadServerDraftRef.current = false;
      lastInjectedDraftRef.current = null;
      setDraftBody('');
      setDraftToken((t) => t + 1);
    }
  }, [suggested.data]);

  const handleAssume = useCallback(() => {
    setAutomation.mutate(
      { mode: 'copilot' },
      {
        onSuccess: () => toast.info('Você assumiu a conversa. A IA não envia sozinha; revise e envie.'),
        onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao assumir'),
      }
    );
  }, [setAutomation]);

  const handleSendNow = useCallback(() => {
    const d = suggested.data;
    if (!d?.reply) return;
    sendWa.mutate(
      { bodyText: d.reply },
      {
        onSuccess: () => {
          toast.success('Resposta da IA enviada.');
          setDraftBody('');
          setDraftToken((t) => t + 1);
          lastInjectedDraftRef.current = null;
        },
        onError: (e) => toast.error(e instanceof Error ? e.message : 'Falha ao enviar'),
      }
    );
  }, [suggested.data, sendWa]);

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

  // Eventos de status de entrega (sent/delivered/read) viram ticks na bolha, não linhas
  // soltas. Mapeia por id da mensagem original (metadata.original_activity) e por wa_message_id,
  // guardando o status mais avançado (failed > read > delivered > sent).
  const STATUS_RANK: Record<string, number> = { sent: 1, delivered: 2, read: 3, failed: 4 };
  const { statusById, statusByWamid } = useMemo(() => {
    const byId: Record<string, string> = {};
    const byWamid: Record<string, string> = {};
    for (const m of messages ?? []) {
      if (m.type !== 'message_status_changed') continue;
      const st = m.metadata?.status;
      if (!st) continue;
      const origId = m.metadata?.original_activity as string | undefined;
      const wamid = m.metadata?.wa_message_id;
      const better = (cur?: string) => !cur || (STATUS_RANK[st] ?? 0) > (STATUS_RANK[cur] ?? 0);
      if (origId && better(byId[origId])) byId[origId] = st;
      if (wamid && better(byWamid[wamid])) byWamid[wamid] = st;
    }
    return { statusById: byId, statusByWamid: byWamid };
  }, [messages]);

  const deliveryStatusFor = useCallback(
    (m: ConversationMessage): string | undefined =>
      // Prioriza o casamento por wa_message_id (agrega o status mais avançado da mensagem):
      // os eventos delivered/read podem ter original_activity errado, mas todos compartilham
      // o wa_message_id da mensagem enviada. Cai pro id da atividade como reforço.
      (m.metadata?.wa_message_id ? statusByWamid[m.metadata.wa_message_id] : undefined) ?? statusById[m.id],
    [statusById, statusByWamid]
  );

  const sortedMessages = useMemo(() => {
    if (!messages) return [];
    // Esconde os eventos de status (viram ticks) e ordena ASC defensivamente.
    return [...messages]
      .filter((m) => m.type !== 'message_status_changed')
      .sort((a, b) => new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime());
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
              {channel !== 'email' && (
                <div className="px-2 py-1.5">
                  <p className="mb-1 text-[10px] uppercase tracking-wide text-muted-foreground">
                    Recepção IA
                  </p>
                  <AutomationToggle ownerType={item.ownerType} ownerId={item.ownerId} />
                </div>
              )}
              <DropdownMenuItem
                onSelect={(e) => {
                  e.preventDefault();
                  setDossierOpen(true);
                }}
              >
                <IdCard className="mr-2 h-4 w-4" /> Dossiê
              </DropdownMenuItem>
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
            {channel !== 'email' && (
              <AutomationToggle ownerType={item.ownerType} ownerId={item.ownerId} />
            )}
            <Button
              type="button"
              variant="ghost"
              size="sm"
              onClick={() => setDossierOpen(true)}
              className="hidden shrink-0 sm:inline-flex"
              aria-label="Ver dossiê do contato"
              title="Dossiê — o que sabemos sobre o contato"
            >
              <IdCard className="mr-1 h-3.5 w-3.5" /> Dossiê
            </Button>
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
                    deliveryStatus={deliveryStatusFor(msg)}
                  />
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* Barra do rascunho do servidor (Atendimento IA) — preview X/3 no Auto */}
      {isWhatsApp && suggested.data?.reply && (
        <AIDraftBar
          draft={suggested.data}
          previewSeconds={receptionSettings.data?.previewSeconds ?? 10}
          onAssume={handleAssume}
          onSendNow={handleSendNow}
          assuming={setAutomation.isPending}
          sending={sendWa.isPending}
        />
      )}

      {/* Composer */}
      <ConversationComposer
        item={item}
        draftBody={draftBody}
        draftToken={draftToken}
        onSuggestionApplied={handleSuggestionInjected}
        lockChannel={channel}
        compact={compact}
      />

      {/* Painel: Dossiê 360 (social) */}
      <DossierPanel
        ownerType={item.ownerType}
        ownerId={item.ownerId}
        name={item.name}
        open={dossierOpen}
        onOpenChange={setDossierOpen}
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
                {formatDate(summaryDialog.generatedAt, "dd/MM 'às' HH:mm")}
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
