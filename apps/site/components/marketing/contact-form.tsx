'use client';

import Link from 'next/link';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { ChevronDown } from 'lucide-react';
import { track } from '@/lib/plausible';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';
import { cn } from '@/lib/cn';

const schema = z.object({
  name: z.string().min(2, 'Nome obrigatório'),
  email: z.string().email('Email inválido'),
  phone: z.string().min(8, 'Telefone obrigatório'),
  reason: z.enum(['cansaco', 'performance', 'cronica', 'prevencao', 'checkup', 'outro']).optional(),
  window: z.string().optional(),
  message: z.string().optional(),
});

type FormData = z.infer<typeof schema>;

export function ContactForm() {
  const [status, setStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>('idle');
  const [consentAccepted, setConsentAccepted] = useState(false);
  const [showMore, setShowMore] = useState(false);
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
      track('form_contato_enviado', { reason: parsed.data.reason ?? 'nao_informado' });
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
        <input {...register('name')} className="form-input" autoComplete="name" />
      </Field>
      <Field label="Email" error={errors.email?.message}>
        <input {...register('email')} type="email" className="form-input" autoComplete="email" />
      </Field>
      <Field label="Telefone (WhatsApp)" error={errors.phone?.message}>
        <input {...register('phone')} inputMode="tel" className="form-input" autoComplete="tel" />
      </Field>

      <button
        type="button"
        onClick={() => setShowMore((s) => !s)}
        aria-expanded={showMore}
        className="flex items-center gap-2 label-upper text-gold hover:text-petrol transition"
      >
        <span>{showMore ? 'Esconder detalhes' : 'Conte mais (opcional)'}</span>
        <ChevronDown size={14} className={cn('transition-transform', showMore && 'rotate-180')} />
      </button>

      <div
        className={cn(
          'grid transition-all duration-300 ease-out',
          showMore ? 'grid-rows-[1fr] opacity-100' : 'grid-rows-[0fr] opacity-0',
        )}
      >
        <div className="overflow-hidden">
          <div className="space-y-6 pt-2">
            <Field label="Motivo principal (opcional)">
              <select {...register('reason')} className="form-input" defaultValue="">
                <option value="">Selecione (opcional)</option>
                <option value="cansaco">Cansaço / queda de energia</option>
                <option value="performance">Performance / longevidade</option>
                <option value="cronica">Condição crônica</option>
                <option value="prevencao">Prevenção</option>
                <option value="checkup">Check-up</option>
                <option value="outro">Outro</option>
              </select>
            </Field>
            <Field label="Melhor horário para contato (opcional)">
              <input
                {...register('window')}
                placeholder="Ex: manhã, tarde, fins de semana"
                className="form-input"
              />
            </Field>
            <Field label="Mensagem (opcional)">
              <textarea
                {...register('message')}
                rows={3}
                placeholder="Algo que ajude a equipe a entender o seu momento"
                className="form-input"
              />
            </Field>
          </div>
        </div>
      </div>

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
