'use client';

import { useCallback, useEffect, useMemo, useState } from 'react';
import { useDropzone } from 'react-dropzone';
import {
  FileText,
  Image as ImageIcon,
  Loader2,
  Mail,
  MessageSquare,
  Paperclip,
  Send,
  Sparkles,
  X,
  Zap,
} from 'lucide-react';
import { toast } from 'sonner';
import { CANNED_REPLIES, applyCannedReply } from '@/lib/canned-replies';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  type ConversationItem,
  type SendConversationEmailAttachment,
  useConversationAISuggestion,
  useConversationReceptionReply,
  useSendConversationEmail,
  useSendConversationWhatsApp,
  useSendConversationWhatsAppTemplate,
  useUploadConversationAttachment,
} from '@/lib/api/conversations-api';

type Channel = 'email' | 'whatsapp';

type Props = {
  item: ConversationItem;
  /**
   * Texto de prefill (ex: sugestão IA). Aplicado quando draftToken muda — permite
   * o pai forçar overwrite mesmo se vendedor já digitou. Null = sem prefill.
   */
  draftBody?: string | null;
  /** Token incrementado pelo pai a cada nova sugestão. Dispara o effect de prefill. */
  draftToken?: number;
  /** Callback quando IA gera sugestão e ela é injetada no campo. */
  onSuggestionApplied?: (text: string) => void;
  /**
   * Trava o canal (esconde as abas Email/WhatsApp). Usado nas superfícies dedicadas
   * (página/dock WhatsApp e caixa de e-mail). Default: deixa o usuário escolher.
   */
  lockChannel?: Channel;
  /** Modo enxuto (dock): textarea menor e padding reduzido. */
  compact?: boolean;
};

// Limites: por arquivo 10MB (alinha com handler), total 40MB (limite Resend).
const MAX_PER_FILE = 10 * 1024 * 1024;
const MAX_TOTAL = 40 * 1024 * 1024;

const ACCEPT = {
  'application/pdf': ['.pdf'],
  'image/jpeg': ['.jpg', '.jpeg'],
  'image/png': ['.png'],
  'image/webp': ['.webp'],
  'application/msword': ['.doc'],
  'application/vnd.openxmlformats-officedocument.wordprocessingml.document': ['.docx'],
  'application/vnd.ms-excel': ['.xls'],
  'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet': ['.xlsx'],
  'audio/ogg': ['.ogg'],
  'audio/mpeg': ['.mp3'],
  'audio/mp4': ['.m4a'],
  'audio/aac': ['.aac'],
};

type AttachmentState = {
  /** ID local pra rastreio na UI (não enviado pro backend). */
  localId: string;
  file: File;
  status: 'uploading' | 'done' | 'error';
  /** Path retornado pelo upload endpoint (preenchido em status=done). */
  path?: string;
  /** Mensagem de erro pra UI quando status=error. */
  error?: string;
};

/** Define o canal default baseado no que o owner suporta + último canal usado. */
function defaultChannel(item: ConversationItem): Channel {
  if (item.lastChannel === 'whatsapp' && item.phone) return 'whatsapp';
  if (item.lastChannel === 'email' && item.email) return 'email';
  if (item.email) return 'email';
  if (item.phone) return 'whatsapp';
  return 'email';
}

type Validation = { ok: true } | { ok: false; reason: string };

function validateEmail(item: ConversationItem): Validation {
  if (!item.email) return { ok: false, reason: 'Sem email cadastrado.' };
  if (item.ownerType === 'lead' && item.emailOptIn === false) {
    return { ok: false, reason: 'Lead sem opt-in de email.' };
  }
  return { ok: true };
}

function validateWhatsApp(item: ConversationItem): Validation {
  if (!item.phone) return { ok: false, reason: 'Sem WhatsApp cadastrado.' };
  if (item.ownerType === 'lead' && item.whatsAppOptIn === false) {
    return { ok: false, reason: 'Lead sem opt-in de WhatsApp.' };
  }
  if (item.ownerType === 'lead') {
    if (!item.lastInboundAt) {
      return {
        ok: false,
        reason: 'Sem inbound recente — janela 24h fechada. Use template aprovado.',
      };
    }
    const elapsed = Date.now() - new Date(item.lastInboundAt).getTime();
    if (elapsed >= 24 * 36e5) {
      return { ok: false, reason: 'Janela 24h expirou. Use template aprovado.' };
    }
  }
  return { ok: true };
}

/** Formata bytes pra "1.2 MB" / "245 KB" / "12 B". */
function formatBytes(bytes: number): string {
  if (bytes < 1024) return `${bytes} B`;
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(0)} KB`;
  return `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
}

function fileIcon(file: File) {
  if (file.type.startsWith('image/')) return ImageIcon;
  return FileText;
}

let localIdCounter = 0;
function nextLocalId(): string {
  localIdCounter += 1;
  return `att-${Date.now()}-${localIdCounter}`;
}

export function ConversationComposer({
  item,
  draftBody,
  draftToken,
  onSuggestionApplied,
  lockChannel,
  compact,
}: Props) {
  const [channel, setChannel] = useState<Channel>(() => lockChannel ?? defaultChannel(item));
  const [subject, setSubject] = useState('');
  const [body, setBody] = useState('');
  const [showCanned, setShowCanned] = useState(false);
  const [attachments, setAttachments] = useState<AttachmentState[]>([]);
  // Template (reabre conversa fora da janela 24h).
  const [showTemplate, setShowTemplate] = useState(false);
  const [tplName, setTplName] = useState('');
  const [tplLang, setTplLang] = useState('pt_BR');
  const [tplParams, setTplParams] = useState('');

  // Reset campos ao trocar de conversa
  useEffect(() => {
    setSubject('');
    setBody('');
    setAttachments([]);
    setChannel(lockChannel ?? defaultChannel(item));
  }, [item.ownerType, item.ownerId, lockChannel]);

  // Prefill via prop quando o pai (viewer) injeta uma sugestão IA. Effect dispara
  // só quando draftToken muda — permite reaplicar mesmo se vendedor digitou no meio.
  useEffect(() => {
    if (!draftToken || draftBody == null) return;
    setBody(draftBody);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [draftToken]);

  const sendEmail = useSendConversationEmail(item.ownerType, item.ownerId);
  const sendWa = useSendConversationWhatsApp(item.ownerType, item.ownerId);
  const sendTpl = useSendConversationWhatsAppTemplate(item.ownerType, item.ownerId);
  const upload = useUploadConversationAttachment();

  const handleSendTemplate = () => {
    if (!tplName.trim()) {
      toast.error('Informe o nome do template aprovado.');
      return;
    }
    const params = tplParams
      .split('\n')
      .map((p) => p.trim())
      .filter(Boolean);
    sendTpl.mutate(
      { name: tplName.trim(), language: tplLang.trim() || 'pt_BR', params: params.length ? params : undefined },
      {
        onSuccess: () => {
          toast.success('Template enviado.');
          setShowTemplate(false);
          setTplName('');
          setTplParams('');
        },
        onError: (err: unknown) =>
          toast.error(err instanceof Error ? err.message : 'Falha ao enviar template'),
      }
    );
  };
  const suggest = useConversationAISuggestion(item.ownerType, item.ownerId);
  const reception = useConversationReceptionReply(item.ownerType, item.ownerId);

  const emailValidation = useMemo(() => validateEmail(item), [item]);
  const waValidation = useMemo(() => validateWhatsApp(item), [item]);
  const validation: Validation = channel === 'email' ? emailValidation : waValidation;

  const totalBytes = useMemo(
    () => attachments.reduce((s, a) => s + a.file.size, 0),
    [attachments]
  );
  const hasUploading = attachments.some((a) => a.status === 'uploading');
  const hasErrored = attachments.some((a) => a.status === 'error');

  const isPending = sendEmail.isPending || sendWa.isPending;
  const trimmedBody = body.trim();
  const readyAttachments = attachments.filter((a) => a.status === 'done' && a.path);
  const canSend =
    validation.ok &&
    (!!trimmedBody || (channel === 'whatsapp' && readyAttachments.length > 0)) &&
    !isPending &&
    !hasUploading &&
    !hasErrored;

  // Upload de uma lista de arquivos. Faz validação local + dispara POST individual.
  const startUploads = useCallback(
    (files: File[]) => {
      // WhatsApp envia uma mídia por mensagem.
      if (channel === 'whatsapp' && attachments.length + files.length > 1) {
        toast.error('WhatsApp: envie uma mídia por mensagem.');
        return;
      }
      // Validação local: per-file e somatório.
      let projectedTotal = totalBytes;
      const accepted: File[] = [];
      for (const file of files) {
        if (file.size > MAX_PER_FILE) {
          toast.error(`${file.name} excede 10MB.`);
          continue;
        }
        if (projectedTotal + file.size > MAX_TOTAL) {
          toast.error(
            `Soma de anexos passaria de ${formatBytes(MAX_TOTAL)} — ${file.name} ignorado.`
          );
          continue;
        }
        projectedTotal += file.size;
        accepted.push(file);
      }
      if (accepted.length === 0) return;

      const newStates: AttachmentState[] = accepted.map((file) => ({
        localId: nextLocalId(),
        file,
        status: 'uploading',
      }));
      setAttachments((prev) => [...prev, ...newStates]);

      newStates.forEach((s) => {
        upload.mutate(s.file, {
          onSuccess: (resp) => {
            setAttachments((prev) =>
              prev.map((a) =>
                a.localId === s.localId
                  ? { ...a, status: 'done', path: resp.path }
                  : a
              )
            );
          },
          onError: (err: unknown) => {
            const msg = err instanceof Error ? err.message : 'Falha no upload';
            setAttachments((prev) =>
              prev.map((a) =>
                a.localId === s.localId ? { ...a, status: 'error', error: msg } : a
              )
            );
            toast.error(`${s.file.name}: ${msg}`);
          },
        });
      });
    },
    [channel, totalBytes, upload, attachments.length]
  );

  const onDrop = useCallback(
    (acceptedFiles: File[]) => startUploads(acceptedFiles),
    [startUploads]
  );

  const { getRootProps, getInputProps, isDragActive, open } = useDropzone({
    onDrop,
    accept: ACCEPT,
    noClick: true, // só drop + botão dedicado pra evitar abrir picker ao clicar no textarea
    noKeyboard: true,
    disabled: !validation.ok || isPending,
  });

  const removeAttachment = (localId: string) => {
    setAttachments((prev) => prev.filter((a) => a.localId !== localId));
  };

  const handleSuggest = () => {
    if (channel !== 'email') {
      toast.error('Sugestão IA disponível apenas no canal email.');
      return;
    }
    suggest.mutate(
      {},
      {
        onSuccess: (data) => {
          setBody(data.suggestion);
          onSuggestionApplied?.(data.suggestion);
          toast.info('Sugestão da IA inserida no campo. Edite antes de enviar.');
        },
        onError: (err: unknown) => {
          const msg = err instanceof Error ? err.message : 'Falha ao gerar sugestão';
          toast.error(msg);
        },
      }
    );
  };

  // Recepcionista virtual (Copiloto): gera a resposta ancorada no script + objeções e
  // insere no campo pro atendente revisar e enviar. Avisa quando o caso pede handoff.
  const handleReceptionReply = () => {
    reception.mutate(undefined, {
      onSuccess: (data) => {
        setBody(data.reply);
        onSuggestionApplied?.(data.reply);
        if (data.action === 'handoff') {
          toast.warning(
            data.handoffReason
              ? `IA sugere passar para um humano: ${data.handoffReason}`
              : 'IA sugere passar este atendimento para um humano.'
          );
        } else {
          toast.info('Resposta da recepção inserida. Revise antes de enviar.');
        }
      },
      onError: (err: unknown) => {
        toast.error(err instanceof Error ? err.message : 'Falha ao gerar resposta da recepção');
      },
    });
  };

  // Insere uma resposta rápida no corpo (anexa se já houver rascunho), com {nome}
  // resolvido pelo primeiro nome do contato.
  const insertCanned = (text: string) => {
    const applied = applyCannedReply(text, item.name);
    setBody((prev) => (prev.trim() ? `${prev.trim()}\n\n${applied}` : applied));
    setShowCanned(false);
  };

  const handleSend = () => {
    if (!validation.ok) {
      toast.error(validation.reason);
      return;
    }
    if (!trimmedBody && !(channel === 'whatsapp' && readyAttachments.length > 0)) {
      toast.error('Escreva uma mensagem antes de enviar.');
      return;
    }

    if (channel === 'email') {
      const readyAtts: SendConversationEmailAttachment[] = attachments
        .filter((a) => a.status === 'done' && a.path)
        .map((a) => ({ path: a.path as string, filename: a.file.name }));

      sendEmail.mutate(
        {
          subject: subject.trim() || undefined,
          bodyText: trimmedBody,
          attachments: readyAtts.length > 0 ? readyAtts : undefined,
        },
        {
          onSuccess: () => {
            toast.success('Email enviado.');
            setSubject('');
            setBody('');
            setAttachments([]);
          },
          onError: (err: unknown) => {
            const msg = err instanceof Error ? err.message : 'Falha ao enviar email';
            toast.error(msg);
          },
        }
      );
      return;
    }

    const waAtts: SendConversationEmailAttachment[] = readyAttachments.map((a) => ({
      path: a.path as string,
      filename: a.file.name,
    }));
    sendWa.mutate(
      {
        bodyText: trimmedBody || undefined,
        attachments: waAtts.length > 0 ? waAtts : undefined,
      },
      {
        onSuccess: () => {
          toast.success('WhatsApp enviado.');
          setBody('');
          setAttachments([]);
        },
        onError: (err: unknown) => {
          const msg = err instanceof Error ? err.message : 'Falha ao enviar WhatsApp';
          toast.error(msg);
        },
      }
    );
  };

  const showAttachUI = true;

  return (
    <div className={cn('border-t border-border bg-background', compact ? 'p-2' : 'p-3 sm:p-4')}>
      {/* Channel switcher */}
      <div className={cn('flex flex-wrap items-center gap-2', compact ? 'mb-2' : 'mb-3')} role="tablist" aria-label="Canal de envio">
        {!lockChannel && (
          <>
            <button
              type="button"
              role="tab"
              aria-selected={channel === 'email'}
              onClick={() => setChannel('email')}
              disabled={emailValidation.ok === false && emailValidation.reason === 'Sem email cadastrado.'}
              title={emailValidation.ok ? '' : emailValidation.reason}
              className={cn(
                'inline-flex items-center gap-1.5 rounded-full border px-3 py-1.5 text-xs font-medium transition-colors',
                channel === 'email'
                  ? 'border-sky-300 bg-sky-100 text-sky-900'
                  : 'border-border bg-background text-muted-foreground hover:bg-muted',
                !item.email && 'cursor-not-allowed opacity-50'
              )}
            >
              <Mail className="h-3.5 w-3.5" /> Email
            </button>
            <button
              type="button"
              role="tab"
              aria-selected={channel === 'whatsapp'}
              onClick={() => setChannel('whatsapp')}
              disabled={waValidation.ok === false && waValidation.reason === 'Sem WhatsApp cadastrado.'}
              title={waValidation.ok ? '' : waValidation.reason}
              className={cn(
                'inline-flex items-center gap-1.5 rounded-full border px-3 py-1.5 text-xs font-medium transition-colors',
                channel === 'whatsapp'
                  ? 'border-emerald-300 bg-emerald-100 text-emerald-900'
                  : 'border-border bg-background text-muted-foreground hover:bg-muted',
                !item.phone && 'cursor-not-allowed opacity-50'
              )}
            >
              <MessageSquare className="h-3.5 w-3.5" /> WhatsApp
            </button>
          </>
        )}

        {/* Recepcionista virtual (Copiloto) — ancorado no script + objeções; ambos canais. */}
        <button
          type="button"
          onClick={handleReceptionReply}
          disabled={!validation.ok || isPending || reception.isPending}
          title="Gerar resposta da recepção (script + objeções)"
          className={cn(
            'ml-auto inline-flex items-center gap-1.5 rounded-full border border-emerald-300 bg-white px-3 py-1.5 text-xs font-medium text-emerald-900 transition-colors hover:bg-emerald-50',
            'disabled:cursor-not-allowed disabled:opacity-50'
          )}
        >
          {reception.isPending ? (
            <Loader2 className="h-3.5 w-3.5 animate-spin" />
          ) : (
            <Sparkles className="h-3.5 w-3.5" />
          )}
          {reception.isPending ? 'Pensando…' : 'Recepção IA'}
        </button>

        {/* Sugestão IA genérica — só email no MVP. */}
        {channel === 'email' && (
          <button
            type="button"
            onClick={handleSuggest}
            disabled={!validation.ok || isPending || suggest.isPending}
            title="Gerar sugestão de resposta com IA"
            className={cn(
              'inline-flex items-center gap-1.5 rounded-full border border-violet-300 bg-white px-3 py-1.5 text-xs font-medium text-violet-900 transition-colors hover:bg-violet-50',
              'disabled:cursor-not-allowed disabled:opacity-50'
            )}
          >
            {suggest.isPending ? (
              <Loader2 className="h-3.5 w-3.5 animate-spin" />
            ) : (
              <Sparkles className="h-3.5 w-3.5" />
            )}
            {suggest.isPending ? 'Pensando…' : 'Sugerir resposta'}
          </button>
        )}

        {/* Template — só WhatsApp; reabre conversa fora da janela 24h. */}
        {channel === 'whatsapp' && (
          <button
            type="button"
            onClick={() => setShowTemplate((v) => !v)}
            title="Enviar template aprovado (fora da janela de 24h)"
            className={cn(
              'ml-auto inline-flex items-center gap-1.5 rounded-full border px-3 py-1.5 text-xs font-medium transition-colors',
              showTemplate
                ? 'border-emerald-400 bg-emerald-100 text-emerald-900'
                : 'border-emerald-300 bg-white text-emerald-900 hover:bg-emerald-50'
            )}
          >
            <MessageSquare className="h-3.5 w-3.5" /> Template
          </button>
        )}
      </div>

      {/* Painel de template WhatsApp */}
      {channel === 'whatsapp' && showTemplate && (
        <div className="mb-2 space-y-2 rounded-md border border-emerald-200 bg-emerald-50/40 p-3">
          <p className="text-xs text-muted-foreground">
            Use um template aprovado pela Meta pra reabrir a conversa fora da janela de 24h.
          </p>
          <div className="flex gap-2">
            <Input
              value={tplName}
              onChange={(e) => setTplName(e.target.value)}
              placeholder="Nome do template (ex: reengajamento)"
              className="text-sm"
            />
            <Input
              value={tplLang}
              onChange={(e) => setTplLang(e.target.value)}
              placeholder="pt_BR"
              className="w-28 text-sm"
            />
          </div>
          <textarea
            value={tplParams}
            onChange={(e) => setTplParams(e.target.value)}
            rows={2}
            placeholder="Parâmetros do corpo ({{1}}, {{2}}…), um por linha (opcional)"
            className="w-full resize-y rounded-md border border-input bg-transparent p-2 text-sm focus:outline-none focus:ring-1 focus:ring-ring"
          />
          <div className="flex justify-end">
            <Button
              type="button"
              size="sm"
              disabled={sendTpl.isPending || !tplName.trim()}
              onClick={handleSendTemplate}
            >
              {sendTpl.isPending && <Loader2 className="h-3.5 w-3.5 animate-spin" />}
              Enviar template
            </Button>
          </div>
        </div>
      )}

      {/* Subject só pra email */}
      {channel === 'email' && (
        <Input
          value={subject}
          onChange={(e) => setSubject(e.target.value)}
          placeholder="Assunto (opcional — backend monta 'Re: ...' do histórico)"
          className="mb-2 text-sm"
          disabled={!validation.ok || isPending}
        />
      )}

      {/* Textarea com dropzone só no canal email */}
      <div
        {...(showAttachUI ? getRootProps() : {})}
        className={cn(
          'relative rounded-md border border-input transition-colors',
          isDragActive && 'border-sky-400 bg-sky-50/40 ring-2 ring-sky-200'
        )}
      >
        {showAttachUI && <input {...getInputProps()} />}
        <textarea
          value={body}
          onChange={(e) => setBody(e.target.value)}
          rows={compact ? 2 : 4}
          placeholder={
            channel === 'email'
              ? 'Escreva o email… (arraste arquivos pra anexar)'
              : 'Escreva a mensagem WhatsApp…'
          }
          disabled={!validation.ok || isPending}
          className={cn(
            'w-full resize-y rounded-md border-0 bg-transparent p-3 text-sm',
            'focus:outline-none focus:ring-1 focus:ring-ring',
            'disabled:cursor-not-allowed disabled:opacity-50'
          )}
          aria-label={channel === 'email' ? 'Corpo do email' : 'Corpo do WhatsApp'}
        />
        {showAttachUI && isDragActive && (
          <div className="pointer-events-none absolute inset-0 flex items-center justify-center rounded-md bg-sky-100/60 text-sm font-medium text-sky-900">
            Solte os arquivos pra anexar
          </div>
        )}
      </div>

      {/* Lista de anexos */}
      {showAttachUI && attachments.length > 0 && (
        <ul className="mt-2 space-y-1.5" aria-label="Anexos">
          {attachments.map((a) => {
            const Icon = fileIcon(a.file);
            return (
              <li
                key={a.localId}
                className={cn(
                  'flex items-center gap-2 rounded-md border px-2 py-1.5 text-xs',
                  a.status === 'error'
                    ? 'border-rose-300 bg-rose-50 text-rose-900'
                    : 'border-border bg-muted/40'
                )}
              >
                <Icon className="h-4 w-4 shrink-0 text-muted-foreground" aria-hidden />
                <span className="min-w-0 flex-1 truncate" title={a.file.name}>
                  {a.file.name}
                </span>
                <span className="shrink-0 text-muted-foreground">
                  {formatBytes(a.file.size)}
                </span>
                {a.status === 'uploading' && (
                  <Loader2 className="h-3.5 w-3.5 shrink-0 animate-spin text-muted-foreground" aria-label="Enviando" />
                )}
                {a.status === 'error' && (
                  <span className="shrink-0 text-[11px]" title={a.error}>
                    erro
                  </span>
                )}
                <button
                  type="button"
                  onClick={() => removeAttachment(a.localId)}
                  className="shrink-0 rounded p-0.5 hover:bg-background"
                  aria-label={`Remover ${a.file.name}`}
                >
                  <X className="h-3.5 w-3.5" />
                </button>
              </li>
            );
          })}
        </ul>
      )}

      {/* Respostas rápidas (canned text) */}
      {showCanned && (
        <div className="mt-2 flex flex-wrap gap-1.5 rounded-md border border-border bg-muted/30 p-2">
          {CANNED_REPLIES.map((r) => (
            <button
              key={r.id}
              type="button"
              onClick={() => insertCanned(r.text)}
              title={applyCannedReply(r.text, item.name)}
              className="rounded-full border border-border bg-background px-2.5 py-1 text-xs font-medium hover:bg-muted"
            >
              {r.label}
            </button>
          ))}
        </div>
      )}

      <div className="mt-2 flex flex-wrap items-center justify-between gap-2">
        <div className="flex min-w-0 flex-1 flex-wrap items-center gap-x-3 gap-y-1 text-xs">
          {showAttachUI && (
            <button
              type="button"
              onClick={open}
              disabled={!validation.ok || isPending}
              className="inline-flex items-center gap-1 rounded-md border border-border bg-background px-2 py-1 text-xs font-medium hover:bg-muted disabled:cursor-not-allowed disabled:opacity-50"
            >
              <Paperclip className="h-3.5 w-3.5" /> Anexar
            </button>
          )}
          <button
            type="button"
            onClick={() => setShowCanned((v) => !v)}
            className={cn(
              'inline-flex items-center gap-1 rounded-md border px-2 py-1 text-xs font-medium hover:bg-muted',
              showCanned
                ? 'border-primary/40 bg-primary/5'
                : 'border-border bg-background'
            )}
          >
            <Zap className="h-3.5 w-3.5" /> Respostas rápidas
          </button>
          {showAttachUI && attachments.length > 0 && (
            <span
              className={cn(
                'text-muted-foreground',
                totalBytes > MAX_TOTAL && 'text-rose-700'
              )}
            >
              {formatBytes(totalBytes)} / {formatBytes(MAX_TOTAL)}
            </span>
          )}
          {!validation.ok ? (
            <p className="wrap-break-word text-amber-700" role="alert">
              {validation.reason}
            </p>
          ) : (
            <p className="text-muted-foreground">
              {channel === 'email'
                ? `Enviando para ${item.email}`
                : `Enviando para ${item.phone}`}
            </p>
          )}
        </div>
        <Button onClick={handleSend} disabled={!canSend} size="sm">
          <Send className="mr-1 h-3.5 w-3.5" />
          {isPending ? 'Enviando…' : hasUploading ? 'Aguardando upload…' : 'Enviar'}
        </Button>
      </div>
    </div>
  );
}
