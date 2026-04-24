'use client';

import { useEffect, useMemo, useState } from 'react';
import { Mail, MessageSquare, Send } from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  type ConversationItem,
  useSendConversationEmail,
  useSendConversationWhatsApp,
} from '@/lib/api/conversations-api';

type Channel = 'email' | 'whatsapp';

type Props = {
  item: ConversationItem;
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
  // emailOptIn só vem pra Lead — pra Patient backend autoriza por padrão
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
    // Janela de 24h: só vale pra Lead (Patient não precisa)
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

export function ConversationComposer({ item }: Props) {
  const [channel, setChannel] = useState<Channel>(() => defaultChannel(item));
  const [subject, setSubject] = useState('');
  const [body, setBody] = useState('');

  // Reset campos ao trocar de conversa
  useEffect(() => {
    setSubject('');
    setBody('');
    setChannel(defaultChannel(item));
  }, [item.ownerType, item.ownerId]);

  const sendEmail = useSendConversationEmail(item.ownerType, item.ownerId);
  const sendWa = useSendConversationWhatsApp(item.ownerType, item.ownerId);

  const emailValidation = useMemo(() => validateEmail(item), [item]);
  const waValidation = useMemo(() => validateWhatsApp(item), [item]);
  const validation: Validation = channel === 'email' ? emailValidation : waValidation;

  const isPending = sendEmail.isPending || sendWa.isPending;
  const trimmedBody = body.trim();
  const canSend = validation.ok && !!trimmedBody && !isPending;

  const handleSend = () => {
    if (!validation.ok) {
      toast.error(validation.reason);
      return;
    }
    if (!trimmedBody) {
      toast.error('Escreva uma mensagem antes de enviar.');
      return;
    }

    if (channel === 'email') {
      sendEmail.mutate(
        {
          subject: subject.trim() || undefined,
          bodyText: trimmedBody,
        },
        {
          onSuccess: () => {
            toast.success('Email enviado.');
            setSubject('');
            setBody('');
          },
          onError: (err: unknown) => {
            const msg = err instanceof Error ? err.message : 'Falha ao enviar email';
            toast.error(msg);
          },
        }
      );
      return;
    }

    sendWa.mutate(
      { bodyText: trimmedBody },
      {
        onSuccess: () => {
          toast.success('WhatsApp enviado.');
          setBody('');
        },
        onError: (err: unknown) => {
          const msg = err instanceof Error ? err.message : 'Falha ao enviar WhatsApp';
          toast.error(msg);
        },
      }
    );
  };

  return (
    <div className="border-t border-border bg-background p-3 sm:p-4">
      {/* Channel switcher */}
      <div className="mb-3 flex flex-wrap items-center gap-2" role="tablist" aria-label="Canal de envio">
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
      </div>

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

      <textarea
        value={body}
        onChange={(e) => setBody(e.target.value)}
        rows={4}
        placeholder={
          channel === 'email' ? 'Escreva o email…' : 'Escreva a mensagem WhatsApp…'
        }
        disabled={!validation.ok || isPending}
        className={cn(
          'w-full resize-y rounded-md border border-input bg-transparent p-3 text-sm',
          'focus:outline-none focus:ring-1 focus:ring-ring',
          'disabled:cursor-not-allowed disabled:opacity-50'
        )}
        aria-label={channel === 'email' ? 'Corpo do email' : 'Corpo do WhatsApp'}
      />

      <div className="mt-2 flex flex-wrap items-center justify-between gap-2">
        <div className="min-w-0 flex-1 text-xs">
          {!validation.ok ? (
            <p className="break-words text-amber-700" role="alert">
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
          {isPending ? 'Enviando…' : 'Enviar'}
        </Button>
      </div>
    </div>
  );
}
