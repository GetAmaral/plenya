'use client';

import { useCallback, useMemo, useState } from 'react';
import { useDropzone } from 'react-dropzone';
import {
  FileText,
  Image as ImageIcon,
  Loader2,
  Mail,
  Paperclip,
  Send,
  X,
} from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
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
  COMPOSE_EMAIL_REGEX,
  type ComposeConversationResult,
  type ConversationOwnerType,
  type SendConversationEmailAttachment,
  useConversationCompose,
  useUploadConversationAttachment,
} from '@/lib/api/conversations-api';

// Limites alinhados com o composer interno (Bloco 4): 10MB por arquivo, 40MB total (Resend).
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
};

type AttachmentState = {
  localId: string;
  file: File;
  status: 'uploading' | 'done' | 'error';
  path?: string;
  error?: string;
};

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
  return `compose-att-${Date.now()}-${localIdCounter}`;
}

type Props = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  /**
   * Chamado após envio bem-sucedido. Recebe owner+id criado/reusado pra UI
   * navegar/selecionar a conversa imediatamente.
   */
  onSent?: (result: { ownerType: ConversationOwnerType; ownerId: string }) => void;
};

/**
 * NewEmailDialog — botão "Novo email" global da Central de Conversas.
 *
 * UX:
 *   - Usuário digita destinatário (validação regex local + server)
 *   - Backend resolve: Patient existente → reusa; Lead ativo → reusa;
 *     senão cria Lead novo (source=manual) e usa como owner
 *   - Toast diferencia "Email enviado" (reuso) de "Lead novo: X" (criação)
 *   - Anexos compartilham infra de upload do composer interno
 */
export function NewEmailDialog({ open, onOpenChange, onSent }: Props) {
  const [to, setTo] = useState('');
  const [name, setName] = useState('');
  const [subject, setSubject] = useState('');
  const [body, setBody] = useState('');
  const [attachments, setAttachments] = useState<AttachmentState[]>([]);

  const compose = useConversationCompose();
  const upload = useUploadConversationAttachment();

  const trimmedTo = to.trim();
  const trimmedSubject = subject.trim();
  const trimmedBody = body.trim();

  const emailValid = useMemo(
    () => COMPOSE_EMAIL_REGEX.test(trimmedTo),
    [trimmedTo]
  );
  const totalBytes = useMemo(
    () => attachments.reduce((s, a) => s + a.file.size, 0),
    [attachments]
  );
  const hasUploading = attachments.some((a) => a.status === 'uploading');
  const hasErrored = attachments.some((a) => a.status === 'error');

  const isPending = compose.isPending;
  const canSend =
    emailValid &&
    !!trimmedSubject &&
    !!trimmedBody &&
    !isPending &&
    !hasUploading &&
    !hasErrored;

  const reset = useCallback(() => {
    setTo('');
    setName('');
    setSubject('');
    setBody('');
    setAttachments([]);
  }, []);

  const handleClose = (next: boolean) => {
    if (!next && isPending) return; // não fecha durante envio
    onOpenChange(next);
    if (!next) reset();
  };

  const startUploads = useCallback(
    (files: File[]) => {
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
    [totalBytes, upload]
  );

  const onDrop = useCallback(
    (acceptedFiles: File[]) => startUploads(acceptedFiles),
    [startUploads]
  );

  const { getRootProps, getInputProps, isDragActive, open: openPicker } = useDropzone({
    onDrop,
    accept: ACCEPT,
    noClick: true,
    noKeyboard: true,
    disabled: isPending,
  });

  const removeAttachment = (localId: string) => {
    setAttachments((prev) => prev.filter((a) => a.localId !== localId));
  };

  const handleSend = () => {
    if (!emailValid) {
      toast.error('Email inválido.');
      return;
    }
    if (!trimmedSubject || !trimmedBody) {
      toast.error('Preencha assunto e corpo.');
      return;
    }

    const readyAtts: SendConversationEmailAttachment[] = attachments
      .filter((a) => a.status === 'done' && a.path)
      .map((a) => ({ path: a.path as string, filename: a.file.name }));

    compose.mutate(
      {
        to: trimmedTo,
        name: name.trim() || undefined,
        subject: trimmedSubject,
        bodyText: trimmedBody,
        attachments: readyAtts.length > 0 ? readyAtts : undefined,
      },
      {
        onSuccess: (res: ComposeConversationResult) => {
          if (res.leadCreated) {
            const label = name.trim() || trimmedTo;
            toast.success(`Email enviado · Lead novo criado: ${label}`);
          } else {
            const target =
              res.ownerType === 'patient' ? 'paciente existente' : 'lead existente';
            toast.success(`Email enviado · anexado a ${target}`);
          }
          onSent?.({ ownerType: res.ownerType, ownerId: res.ownerId });
          onOpenChange(false);
          reset();
        },
        onError: (err: unknown) => {
          const msg = err instanceof Error ? err.message : 'Falha ao enviar email';
          toast.error(msg);
        },
      }
    );
  };

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2">
            <Mail className="h-5 w-5" /> Novo email
          </DialogTitle>
          <DialogDescription>
            Envia email pra qualquer endereço. Se o destinatário já é Lead ou Paciente,
            anexa à conversa existente; senão, cria um Lead novo automaticamente.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-3">
          <div className="grid gap-2 sm:grid-cols-2">
            <div className="space-y-1">
              <label htmlFor="compose-to" className="text-xs font-medium">
                Para <span className="text-rose-600">*</span>
              </label>
              <Input
                id="compose-to"
                type="email"
                value={to}
                onChange={(e) => setTo(e.target.value)}
                placeholder="email@exemplo.com"
                disabled={isPending}
                className={cn(
                  trimmedTo && !emailValid && 'border-rose-400 focus-visible:ring-rose-400'
                )}
                aria-invalid={!!trimmedTo && !emailValid}
              />
              {trimmedTo && !emailValid && (
                <p className="text-[11px] text-rose-700">Formato de email inválido.</p>
              )}
            </div>
            <div className="space-y-1">
              <label htmlFor="compose-name" className="text-xs font-medium">
                Nome (opcional)
              </label>
              <Input
                id="compose-name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="João Silva"
                disabled={isPending}
              />
            </div>
          </div>

          <div className="space-y-1">
            <label htmlFor="compose-subject" className="text-xs font-medium">
              Assunto <span className="text-rose-600">*</span>
            </label>
            <Input
              id="compose-subject"
              value={subject}
              onChange={(e) => setSubject(e.target.value)}
              placeholder="Sobre nosso bate-papo"
              disabled={isPending}
              maxLength={200}
            />
          </div>

          <div className="space-y-1">
            <label htmlFor="compose-body" className="text-xs font-medium">
              Mensagem <span className="text-rose-600">*</span>
            </label>
            <div
              {...getRootProps()}
              className={cn(
                'relative rounded-md border border-input transition-colors',
                isDragActive && 'border-ocean-400 bg-ocean-50/40 ring-2 ring-ocean-200'
              )}
            >
              <input {...getInputProps()} />
              <textarea
                id="compose-body"
                value={body}
                onChange={(e) => setBody(e.target.value)}
                rows={6}
                placeholder="Escreva o email… (arraste arquivos pra anexar)"
                disabled={isPending}
                className={cn(
                  'w-full resize-y rounded-md border-0 bg-transparent p-3 text-sm',
                  'focus:outline-none focus:ring-1 focus:ring-ring',
                  'disabled:cursor-not-allowed disabled:opacity-50'
                )}
              />
              {isDragActive && (
                <div className="pointer-events-none absolute inset-0 flex items-center justify-center rounded-md bg-ocean-100/60 text-sm font-medium text-ocean-800">
                  Solte os arquivos pra anexar
                </div>
              )}
            </div>
          </div>

          {attachments.length > 0 && (
            <ul className="space-y-1.5" aria-label="Anexos">
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
                      <Loader2
                        className="h-3.5 w-3.5 shrink-0 animate-spin text-muted-foreground"
                        aria-label="Enviando"
                      />
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
                      disabled={isPending}
                    >
                      <X className="h-3.5 w-3.5" />
                    </button>
                  </li>
                );
              })}
            </ul>
          )}

          <div className="flex items-center gap-3 text-xs text-muted-foreground">
            <button
              type="button"
              onClick={openPicker}
              disabled={isPending}
              className="inline-flex items-center gap-1 rounded-md border border-border bg-background px-2 py-1 text-xs font-medium hover:bg-muted disabled:cursor-not-allowed disabled:opacity-50"
            >
              <Paperclip className="h-3.5 w-3.5" /> Anexar
            </button>
            {attachments.length > 0 && (
              <span className={cn(totalBytes > MAX_TOTAL && 'text-rose-700')}>
                {formatBytes(totalBytes)} / {formatBytes(MAX_TOTAL)}
              </span>
            )}
          </div>
        </div>

        <DialogFooter>
          <Button variant="ghost" onClick={() => handleClose(false)} disabled={isPending}>
            Cancelar
          </Button>
          <Button onClick={handleSend} disabled={!canSend}>
            <Send className="mr-1 h-3.5 w-3.5" />
            {isPending ? 'Enviando…' : hasUploading ? 'Aguardando upload…' : 'Enviar'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
