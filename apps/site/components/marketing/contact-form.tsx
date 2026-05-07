'use client';

import Link from 'next/link';
import Script from 'next/script';
import { useEffect, useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { useTranslations } from 'next-intl';
import { z } from 'zod';
import { ChevronDown } from 'lucide-react';
import { track } from '@/lib/plausible';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';
import { cn } from '@/lib/cn';

// M10 — Cloudflare Turnstile site key (pública). Quando vazia, widget some
// e backend pula validação automaticamente (dev local).
const TURNSTILE_SITE_KEY = process.env.NEXT_PUBLIC_TURNSTILE_SITE_KEY ?? '';

// Tipos do Turnstile global.
declare global {
  interface Window {
    turnstile?: {
      render: (selector: string | HTMLElement, opts: { sitekey: string; callback: (token: string) => void; 'error-callback'?: () => void; 'expired-callback'?: () => void; theme?: 'light' | 'dark' | 'auto' }) => string;
      reset: (id?: string) => void;
    };
  }
}

const schema = z.object({
  name: z.string().min(2),
  email: z.string().email(),
  phone: z.string().min(8),
  reason: z.enum(['cansaco', 'performance', 'cronica', 'prevencao', 'checkup', 'outro']).optional(),
  window: z.string().optional(),
  message: z.string().optional(),
});

type FormData = z.infer<typeof schema>;

export function ContactForm() {
  const t = useTranslations('contact');
  const [status, setStatus] = useState<'idle' | 'sending' | 'sent' | 'error'>('idle');
  const [consentAccepted, setConsentAccepted] = useState(false);
  const [showMore, setShowMore] = useState(false);
  const [turnstileToken, setTurnstileToken] = useState('');
  const turnstileRef = useRef<HTMLDivElement | null>(null);
  const turnstileWidgetId = useRef<string | null>(null);
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>();

  // M10 — render Turnstile widget quando script carrega + container monta.
  useEffect(() => {
    if (!TURNSTILE_SITE_KEY) return;
    const tryRender = () => {
      if (!window.turnstile || !turnstileRef.current || turnstileWidgetId.current) return;
      turnstileWidgetId.current = window.turnstile.render(turnstileRef.current, {
        sitekey: TURNSTILE_SITE_KEY,
        callback: (token) => setTurnstileToken(token),
        'error-callback': () => setTurnstileToken(''),
        'expired-callback': () => setTurnstileToken(''),
        theme: 'light',
      });
    };
    tryRender();
    const interval = setInterval(tryRender, 500);
    return () => clearInterval(interval);
  }, []);

  async function onSubmit(data: FormData) {
    const parsed = schema.safeParse(data);
    if (!parsed.success) return;
    if (!consentAccepted) return;
    if (TURNSTILE_SITE_KEY && !turnstileToken) {
      setStatus('error');
      return;
    }
    setStatus('sending');
    try {
      const res = await fetch('/api/leads', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          ...parsed.data,
          source: 'contact-form',
          consentVersion: PRIVACY_POLICY_VERSION,
          turnstileToken: turnstileToken || undefined,
        }),
      });
      if (!res.ok) throw new Error('Failed');
      track('form_contato_enviado', { reason: parsed.data.reason ?? 'nao_informado' });
      setStatus('sent');
    } catch {
      setStatus('error');
      if (window.turnstile && turnstileWidgetId.current) {
        window.turnstile.reset(turnstileWidgetId.current);
        setTurnstileToken('');
      }
    }
  }

  if (status === 'sent') {
    return (
      <div className="border-t border-gold pt-8 space-y-3">
        <p className="label-upper text-gold">{t('formSentLabel')}</p>
        <h3 className="heading-section text-petrol text-2xl">{t('formSentTitle')}</h3>
        <p className="text-petrol/70">{t('formSentDesc')}</p>
      </div>
    );
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      <Field label={t('formNameLabel')} error={errors.name && t('formNameError')}>
        <input {...register('name')} className="form-input" autoComplete="name" />
      </Field>
      <Field label={t('formEmailLabel')} error={errors.email && t('formEmailError')}>
        <input {...register('email')} type="email" className="form-input" autoComplete="email" />
      </Field>
      <Field label={t('formPhoneLabel')} error={errors.phone && t('formPhoneError')}>
        <input {...register('phone')} inputMode="tel" className="form-input" autoComplete="tel" />
      </Field>

      <button
        type="button"
        onClick={() => setShowMore((s) => !s)}
        aria-expanded={showMore}
        className="flex items-center gap-2 label-upper text-gold hover:text-petrol transition"
      >
        <span>{showMore ? t('formMoreHide') : t('formMoreShow')}</span>
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
            <Field label={t('formReasonLabel')}>
              <select {...register('reason')} className="form-input" defaultValue="">
                <option value="">{t('formReasonPlaceholder')}</option>
                <option value="cansaco">{t('formReasonCansaco')}</option>
                <option value="performance">{t('formReasonPerformance')}</option>
                <option value="cronica">{t('formReasonCronica')}</option>
                <option value="prevencao">{t('formReasonPrevencao')}</option>
                <option value="checkup">{t('formReasonCheckup')}</option>
                <option value="outro">{t('formReasonOutro')}</option>
              </select>
            </Field>
            <Field label={t('formWindowLabel')}>
              <input
                {...register('window')}
                placeholder={t('formWindowPlaceholder')}
                className="form-input"
              />
            </Field>
            <Field label={t('formMessageLabel')}>
              <textarea
                {...register('message')}
                rows={3}
                placeholder={t('formMessagePlaceholder')}
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
          {t('formConsentPart1')}
          <Link href="/privacidade" className="underline underline-offset-4">{t('formConsentPrivacyLink')}</Link>
          {t('formConsentPart2')}
          <Link href="/termos" className="underline underline-offset-4">{t('formConsentTermsLink')}</Link>
          {t('formConsentPart3')}
        </span>
      </label>

      {/* M10 — Cloudflare Turnstile (CAPTCHA invisível/leve). Só renderiza se SITE_KEY setado. */}
      {TURNSTILE_SITE_KEY && (
        <>
          <Script
            src="https://challenges.cloudflare.com/turnstile/v0/api.js"
            strategy="afterInteractive"
            async
            defer
          />
          <div ref={turnstileRef} className="cf-turnstile" />
        </>
      )}

      <button
        type="submit"
        disabled={status === 'sending' || !consentAccepted || (Boolean(TURNSTILE_SITE_KEY) && !turnstileToken)}
        className={`btn-gold w-full ${status === 'sending' || !consentAccepted || (Boolean(TURNSTILE_SITE_KEY) && !turnstileToken) ? 'opacity-40 cursor-not-allowed' : ''}`}
      >
        {status === 'sending' ? t('formSubmitting') : t('formSubmit')}
      </button>
      {status === 'error' && (
        <p className="text-sm text-red-700">{t('formError')}</p>
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
