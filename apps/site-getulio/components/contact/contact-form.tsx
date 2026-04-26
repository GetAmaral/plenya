'use client';

import { useState } from 'react';
import { cn } from '@/lib/cn';

const reasons = [
  { value: 'consulta-plenya', label: 'Consulta — Plenya (medicina funcional, longevidade)' },
  { value: 'consulta-nefroclinica', label: 'Consulta — Nefroclínica (nefrologia)' },
  { value: 'palestra', label: 'Convite para palestra' },
  { value: 'imprensa', label: 'Imprensa' },
  { value: 'outro', label: 'Outro' },
] as const;

type Status = 'idle' | 'sending' | 'sent' | 'error';

export function ContactForm() {
  const [reason, setReason] = useState<string>('consulta-plenya');
  const [status, setStatus] = useState<Status>('idle');
  const [error, setError] = useState<string>('');

  async function onSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setStatus('sending');
    setError('');
    const fd = new FormData(e.currentTarget);
    const payload = {
      name: String(fd.get('name') ?? ''),
      email: String(fd.get('email') ?? ''),
      phone: String(fd.get('phone') ?? ''),
      reason,
      message: String(fd.get('message') ?? ''),
    };
    try {
      const res = await fetch('/api/contact', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
      if (!res.ok) {
        const j = await res.json().catch(() => ({}));
        throw new Error(j?.error ?? `HTTP ${res.status}`);
      }
      setStatus('sent');
      e.currentTarget.reset();
      setReason('consulta-plenya');
    } catch (err) {
      setStatus('error');
      setError(err instanceof Error ? err.message : 'Erro ao enviar.');
    }
  }

  if (status === 'sent') {
    return (
      <div className="border border-bordo bg-paper p-8 space-y-3">
        <p className="label-meta text-bordo">Recebido</p>
        <p className="font-serif text-lg text-ink">
          Sua mensagem chegou. Respondo em até 48 horas úteis.
        </p>
      </div>
    );
  }

  return (
    <form onSubmit={onSubmit} className="space-y-8">
      <div className="space-y-3">
        <p className="label-meta">Motivo</p>
        <div className="space-y-2">
          {reasons.map((r) => (
            <label key={r.value} className="flex items-start gap-3 cursor-pointer group">
              <input
                type="radio"
                name="reason"
                value={r.value}
                checked={reason === r.value}
                onChange={() => setReason(r.value)}
                className="mt-1.5 accent-bordo"
              />
              <span className="font-serif text-ink-soft group-hover:text-ink transition-colors">
                {r.label}
              </span>
            </label>
          ))}
        </div>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <FieldText name="name" label="Nome" required autoComplete="name" />
        <FieldText name="phone" label="Telefone" required autoComplete="tel" inputMode="tel" />
      </div>
      <FieldText name="email" label="Email" required type="email" autoComplete="email" />

      <div>
        <label className="block label-meta mb-2">Mensagem (opcional)</label>
        <textarea
          name="message"
          rows={5}
          className="w-full bg-transparent border-b border-rule focus:border-bordo outline-none py-2 font-serif text-ink resize-none transition-colors"
        />
      </div>

      <div className="flex items-center gap-6">
        <button
          type="submit"
          disabled={status === 'sending'}
          className={cn(
            'font-sans text-sm tracking-wide border-b-2 border-bordo pb-1 text-ink hover:text-bordo transition-colors',
            status === 'sending' && 'opacity-50 cursor-wait',
          )}
        >
          {status === 'sending' ? 'Enviando…' : 'Enviar mensagem →'}
        </button>
        {status === 'error' && (
          <p className="font-sans text-xs text-bordo">{error}</p>
        )}
      </div>

      <p className="font-sans text-xs text-ink-muted/70 max-w-md">
        Seus dados são usados exclusivamente para responder a este contato. Não compartilhamos
        com terceiros.
      </p>
    </form>
  );
}

function FieldText({
  name,
  label,
  type = 'text',
  ...rest
}: {
  name: string;
  label: string;
  type?: string;
  required?: boolean;
  autoComplete?: string;
  inputMode?: 'text' | 'tel' | 'email' | 'numeric';
}) {
  return (
    <div>
      <label className="block label-meta mb-2">{label}</label>
      <input
        name={name}
        type={type}
        className="w-full bg-transparent border-b border-rule focus:border-bordo outline-none py-2 font-serif text-ink transition-colors"
        {...rest}
      />
    </div>
  );
}
