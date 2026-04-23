'use client';

import Link from 'next/link';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { track } from '@/lib/plausible';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';

const schema = z.object({
  name: z.string().min(2, 'Nome obrigatório'),
  phone: z.string().min(8, 'Telefone obrigatório'),
  email: z.string().email('Email inválido'),
  reason: z.enum(['cansaco', 'performance', 'cronica', 'prevencao', 'checkup', 'outro']),
  window: z.string().min(2, 'Janela preferida obrigatória'),
});

type FormData = z.infer<typeof schema>;

export function ContactForm() {
  const [status, setStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>('idle');
  const [consentAccepted, setConsentAccepted] = useState(false);
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>();

  async function onSubmit(data: FormData) {
    const parsed = schema.safeParse(data);
    if (!parsed.success) return;
    if (!consentAccepted) return;
    setStatus('sending');
    try {
      const res = await fetch('/api/leads', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          ...parsed.data,
          source: 'contact-form',
          consentVersion: PRIVACY_POLICY_VERSION,
        }),
      });
      if (!res.ok) throw new Error('Falha no envio');
      track('form_contato_enviado', { reason: parsed.data.reason });
      setStatus('sent');
    } catch {
      setStatus('error');
    }
  }

  if (status === 'sent') {
    return (
      <div className="border-t border-gold pt-8 space-y-3">
        <p className="label-upper text-gold">Recebido</p>
        <h3 className="heading-section text-petrol text-2xl">Obrigado por entrar em contato.</h3>
        <p className="text-petrol/70">Nossa equipe responde em até 2 horas em dias úteis.</p>
      </div>
    );
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      <Field label="Nome" error={errors.name?.message}>
        <input {...register('name')} className="form-input" />
      </Field>
      <Field label="Telefone (WhatsApp)" error={errors.phone?.message}>
        <input {...register('phone')} inputMode="tel" className="form-input" />
      </Field>
      <Field label="Email" error={errors.email?.message}>
        <input {...register('email')} type="email" className="form-input" />
      </Field>
      <Field label="Motivo principal" error={errors.reason?.message}>
        <select {...register('reason')} className="form-input" defaultValue="">
          <option value="" disabled>Selecione</option>
          <option value="cansaco">Cansaço / queda de energia</option>
          <option value="performance">Performance / longevidade</option>
          <option value="cronica">Condição crônica</option>
          <option value="prevencao">Prevenção</option>
          <option value="checkup">Check-up</option>
          <option value="outro">Outro</option>
        </select>
      </Field>
      <Field label="Melhor horário para contato" error={errors.window?.message}>
        <input {...register('window')} placeholder="Ex: manhã, tarde, fins de semana" className="form-input" />
      </Field>
      <label className="flex items-start gap-3 cursor-pointer">
        <input
          type="checkbox"
          checked={consentAccepted}
          onChange={(e) => setConsentAccepted(e.target.checked)}
          className="mt-1 h-4 w-4 accent-gold cursor-pointer"
        />
        <span className="text-petrol/80 text-sm leading-relaxed">
          Li e concordo com a{' '}
          <Link href="/privacidade" className="underline underline-offset-4">Política de Privacidade</Link>
          {' '}e os{' '}
          <Link href="/termos" className="underline underline-offset-4">Termos de Uso</Link>. Autorizo a Plenya a entrar em contato pelos canais informados.
        </span>
      </label>
      <button
        type="submit"
        disabled={status === 'sending' || !consentAccepted}
        className={`btn-gold w-full ${status === 'sending' || !consentAccepted ? 'opacity-40 cursor-not-allowed' : ''}`}
      >
        {status === 'sending' ? 'Enviando…' : 'Enviar'}
      </button>
      {status === 'error' && (
        <p className="text-sm text-red-700">Erro ao enviar. Tente novamente ou nos chame no WhatsApp.</p>
      )}
    </form>
  );
}

function Field({ label, error, children }: { label: string; error?: string; children: React.ReactNode }) {
  return (
    <label className="block space-y-2">
      <span className="label-upper text-petrol/60">{label}</span>
      {children}
      {error && <span className="text-xs text-red-700">{error}</span>}
    </label>
  );
}
