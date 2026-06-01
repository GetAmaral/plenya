'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { cn } from '@/lib/cn';

const REASON_CODES = [
  'consulta-plenya',
  'consulta-nefroclinica',
  'palestra',
  'imprensa',
  'outro',
] as const;

type ReasonCode = (typeof REASON_CODES)[number];

const REASON_LABEL_KEY: Record<ReasonCode, string> = {
  'consulta-plenya': 'reasonOptionPlenya',
  'consulta-nefroclinica': 'reasonOptionNefroclinica',
  palestra: 'reasonOptionPalestra',
  imprensa: 'reasonOptionImprensa',
  outro: 'reasonOptionOutro',
};

type Status = 'idle' | 'sending' | 'sent' | 'error';

export function ContactForm() {
  const t = useTranslations('contactForm');
  const [reason, setReason] = useState<ReasonCode | null>(null);
  const [status, setStatus] = useState<Status>('idle');
  const [error, setError] = useState<string>('');

  async function onSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!reason) {
      setStatus('error');
      setError(t('errorReasonRequired'));
      return;
    }
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
      setReason(null);
    } catch (err) {
      setStatus('error');
      setError(err instanceof Error ? err.message : t('errorGeneric'));
    }
  }

  if (status === 'sent') {
    return (
      <div className="border border-bordo bg-paper p-8 space-y-3">
        <p className="label-meta text-bordo">{t('successKicker')}</p>
        <p className="font-serif text-lg text-ink">{t('successBody')}</p>
      </div>
    );
  }

  return (
    <form onSubmit={onSubmit} className="space-y-8">
      <div className="space-y-3">
        <p className="label-meta">
          {t('reasonLabel')} <span aria-hidden="true" className="text-bordo">*</span>
        </p>
        <div className="space-y-2" role="radiogroup" aria-required="true">
          {REASON_CODES.map((code) => (
            <label key={code} className="flex items-start gap-3 cursor-pointer group">
              <input
                type="radio"
                name="reason"
                value={code}
                checked={reason === code}
                onChange={() => setReason(code)}
                required
                className="mt-1.5 accent-bordo"
              />
              <span className="font-serif text-ink-soft group-hover:text-ink transition-colors">
                {t(REASON_LABEL_KEY[code])}
              </span>
            </label>
          ))}
        </div>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <FieldText name="name" label={t('fieldName')} required autoComplete="name" />
        <FieldText name="phone" label={t('fieldPhone')} required autoComplete="tel" inputMode="tel" />
      </div>
      <FieldText name="email" label={t('fieldEmail')} required type="email" autoComplete="email" />

      <div>
        <label className="block label-meta mb-2">{t('fieldMessage')}</label>
        <textarea
          name="message"
          rows={5}
          className="w-full bg-transparent border-b border-rule focus:border-bordo outline-hidden py-2 font-serif text-ink resize-none transition-colors"
        />
      </div>

      <div className="flex items-center gap-6 flex-wrap">
        <button
          type="submit"
          disabled={status === 'sending'}
          className={cn(
            'btn-gold',
            status === 'sending' && 'opacity-50 cursor-wait',
          )}
        >
          {status === 'sending' ? t('submitting') : t('submit')}
        </button>
        {status === 'error' && (
          <p className="font-sans text-xs text-bordo" role="alert">{error}</p>
        )}
      </div>

      <p className="font-sans text-xs text-ink-muted/70 max-w-md">{t('privacy')}</p>
    </form>
  );
}

function FieldText({
  name,
  label,
  type = 'text',
  required,
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
      <label className="block label-meta mb-2" htmlFor={name}>
        {label}
        {required && <span aria-hidden="true" className="text-bordo"> *</span>}
      </label>
      <input
        id={name}
        name={name}
        type={type}
        required={required}
        className="w-full bg-transparent border-b border-rule focus:border-bordo outline-hidden py-2 font-serif text-ink transition-colors"
        {...rest}
      />
    </div>
  );
}
