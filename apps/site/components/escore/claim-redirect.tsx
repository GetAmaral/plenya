'use client';

import { useEffect, useState } from 'react';
import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { confirmClaim } from '@/lib/score-light/api';

const EMR_BASE_URL =
  process.env.NEXT_PUBLIC_EMR_URL?.replace(/\/$/, '') ?? 'http://localhost:3000';

export function ClaimRedirect({ token }: { token: string }) {
  const t = useTranslations('escoreLight');
  const [status, setStatus] = useState<'pending' | 'success' | 'error'>('pending');
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const result = await confirmClaim(token);
        if (cancelled) return;

        try {
          localStorage.setItem('plenya_access_token', result.accessToken);
          localStorage.setItem('plenya_refresh_token', result.refreshToken);
        } catch {
          // localStorage indisponível (modo privado, ssr) — segue para o redirect
        }

        setStatus('success');

        const params = new URLSearchParams({
          accessToken: result.accessToken,
          refreshToken: result.refreshToken,
          sessionCode: result.sessionCode,
        });
        window.location.href = `${EMR_BASE_URL}/patient-portal/escore-light?${params.toString()}`;
      } catch (err) {
        if (cancelled) return;
        setErrorMsg(err instanceof Error ? err.message : null);
        setStatus('error');
      }
    })();
    return () => {
      cancelled = true;
    };
  }, [token]);

  return (
    <section className="bg-cream min-h-[60vh]">
      <div className="site-container pt-32 pb-32 max-w-xl">
        {status === 'pending' && (
          <>
            <p className="label-upper text-gold">{t('claimVerifyingLabel')}</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              {t('claimVerifyingTitle')}
            </h1>
            <p className="text-petrol/70 mt-6">
              {t('claimVerifyingDesc')}
            </p>
          </>
        )}

        {status === 'success' && (
          <>
            <p className="label-upper text-gold">{t('claimSuccessLabel')}</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              {t('claimSuccessTitle')}
            </h1>
          </>
        )}

        {status === 'error' && (
          <>
            <p className="label-upper text-gold">{t('claimErrorLabel')}</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              {t('claimErrorTitle')}
            </h1>
            <p className="text-petrol/70 mt-6 leading-relaxed">
              {errorMsg ?? t('claimErrorDescDefault')}
            </p>
            <div className="pt-8">
              <Link href="/escore-plenya/avaliar" className="btn-gold">
                {t('claimErrorCta')}
              </Link>
            </div>
          </>
        )}
      </div>
    </section>
  );
}
