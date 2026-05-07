'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useTranslations } from 'next-intl';
import { z } from 'zod';
import { track } from '@/lib/plausible';

const schema = z.object({ email: z.string().email() });
type FormData = z.infer<typeof schema>;

export function NewsletterInline({ source = 'blog-post' }: { source?: string }) {
  const t = useTranslations('newsletter');
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
      if (!res.ok) throw new Error('Failed');
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
        <p className="label-upper text-gold">{t('inlineLabel')}</p>
        <p className="heading-section text-2xl text-petrol">{t('inlineTitleConfirmed')}</p>
        <p className="text-petrol/70 text-sm">
          {t('inlineConfirmedDesc')}
        </p>
      </aside>
    );
  }

  return (
    <aside className="my-16 border-l-2 border-gold pl-8 py-6 space-y-4">
      <div className="space-y-1">
        <p className="label-upper text-gold">{t('inlineLabel')}</p>
        <p className="heading-section text-2xl text-petrol max-w-xl">
          {t('inlineTitle')}
        </p>
        <p className="text-petrol/70 text-sm max-w-prose">
          {t('inlineDesc')}
        </p>
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col sm:flex-row gap-3 max-w-xl">
        <input
          {...register('email')}
          type="email"
          placeholder={t('inlinePlaceholder')}
          className="form-input flex-1"
        />
        <button type="submit" disabled={status === 'sending'} className="btn-gold">
          {status === 'sending' ? t('inlineSubscribing') : t('inlineSubscribe')}
        </button>
      </form>
      {errors.email && <p className="text-xs text-red-700">{t('inlineEmailInvalid')}</p>}
      {status === 'error' && (
        <p className="text-xs text-red-700">{t('inlineError')}</p>
      )}
    </aside>
  );
}
