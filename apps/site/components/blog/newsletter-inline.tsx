'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { track } from '@/lib/plausible';

const schema = z.object({ email: z.string().email() });
type FormData = z.infer<typeof schema>;

export function NewsletterInline({ source = 'blog-post' }: { source?: string }) {
  const [status, setStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>('idle');
  const { register, handleSubmit, reset, formState: { errors } } = useForm<FormData>();

  async function onSubmit(data: FormData) {
    const parsed = schema.safeParse(data);
    if (!parsed.success) return;
    setStatus('sending');
    try {
      const res = await fetch('/api/newsletter', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ ...parsed.data, source }),
      });
      if (!res.ok) throw new Error('Falha');
      track('newsletter_inscreveu', { source });
      reset();
      setStatus('sent');
    } catch {
      setStatus('error');
    }
  }

  if (status === 'sent') {
    return (
      <aside className="my-16 border-l-2 border-gold pl-8 py-6 space-y-2">
        <p className="label-upper text-gold">Boletim Plenya</p>
        <p className="heading-section text-2xl text-petrol">Inscrição confirmada.</p>
        <p className="text-petrol/70 text-sm">
          Em instantes você receberá um email para confirmar o cadastro. Verifique a caixa de entrada.
        </p>
      </aside>
    );
  }

  return (
    <aside className="my-16 border-l-2 border-gold pl-8 py-6 space-y-4">
      <div className="space-y-1">
        <p className="label-upper text-gold">Boletim Plenya</p>
        <p className="heading-section text-2xl text-petrol max-w-xl">
          Receba um insight clínico exclusivo por semana.
        </p>
        <p className="text-petrol/70 text-sm max-w-prose">
          Conteúdo curto e direto, escrito pelos médicos da Plenya. Sem spam — você pode descadastrar a qualquer momento.
        </p>
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col sm:flex-row gap-3 max-w-xl">
        <input
          {...register('email')}
          type="email"
          placeholder="seu@email.com"
          className="form-input flex-1"
        />
        <button type="submit" disabled={status === 'sending'} className="btn-gold">
          {status === 'sending' ? 'Inscrevendo…' : 'Inscrever-me'}
        </button>
      </form>
      {errors.email && <p className="text-xs text-red-700">Email inválido</p>}
      {status === 'error' && (
        <p className="text-xs text-red-700">Erro ao inscrever. Tente novamente em instantes.</p>
      )}
    </aside>
  );
}
