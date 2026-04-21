'use client';

import { useEffect, useState } from 'react';
import { Link } from '@/lib/i18n/navigation';
import { confirmClaim } from '@/lib/score-light/api';

const EMR_BASE_URL =
  process.env.NEXT_PUBLIC_EMR_URL?.replace(/\/$/, '') ?? 'http://localhost:3000';

export function ClaimRedirect({ token }: { token: string }) {
  const [status, setStatus] = useState<'pending' | 'success' | 'error'>('pending');
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    (async () => {
      try {
        const result = await confirmClaim(token);
        if (cancelled) return;

        // Persiste tokens no localStorage para o EMR consumir após redirect.
        // O EMR (apps/web) já tem o padrão de ler accessToken/refreshToken aí.
        try {
          localStorage.setItem('plenya_access_token', result.accessToken);
          localStorage.setItem('plenya_refresh_token', result.refreshToken);
        } catch {
          // localStorage indisponível (modo privado, ssr) — segue para o redirect
        }

        setStatus('success');

        // Redireciona para a área autenticada do paciente.
        // O EMR pode aproveitar querystring para rehydratar tokens caso localStorage
        // não tenha sido populado a tempo (cross-origin).
        const params = new URLSearchParams({
          accessToken: result.accessToken,
          refreshToken: result.refreshToken,
          sessionCode: result.sessionCode,
        });
        window.location.href = `${EMR_BASE_URL}/patient-portal/escore-light?${params.toString()}`;
      } catch (err) {
        if (cancelled) return;
        setErrorMsg(err instanceof Error ? err.message : 'Falha ao validar o link');
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
            <p className="label-upper text-gold">Verificando</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              Validando seu link...
            </h1>
            <p className="text-petrol/70 mt-6">
              Em instantes você será redirecionado para sua área pessoal.
            </p>
          </>
        )}

        {status === 'success' && (
          <>
            <p className="label-upper text-gold">Confirmado</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              Tudo certo. Redirecionando...
            </h1>
          </>
        )}

        {status === 'error' && (
          <>
            <p className="label-upper text-gold">Link inválido ou expirado</p>
            <h1 className="heading-section text-petrol text-3xl mt-6">
              Não foi possível validar seu link.
            </h1>
            <p className="text-petrol/70 mt-6 leading-relaxed">
              {errorMsg ?? 'O link expirou (15 minutos) ou já foi usado. Refaça a solicitação no resultado da sua avaliação.'}
            </p>
            <div className="pt-8">
              <Link href="/escore-plenya/avaliar" className="btn-gold">
                Refazer minha avaliação
              </Link>
            </div>
          </>
        )}
      </div>
    </section>
  );
}
