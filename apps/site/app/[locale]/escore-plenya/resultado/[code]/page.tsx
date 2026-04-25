import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { EscoreLightResultado } from '@/components/escore/escore-light-resultado';
import type { PublicSession } from '@/lib/score-light/types';

export const dynamic = 'force-dynamic';

export const metadata: Metadata = {
  title: 'Seu resultado — Escore Plenya',
  description: 'Resultado da sua autoavaliação Escore Plenya.',
  robots: { index: false, follow: false },
};

async function fetchSession(code: string): Promise<PublicSession | null> {
  // SSR roda dentro do container do site. Em dev (docker compose) o serviço
  // EMR é "api". Em produção (Coolify) é a URL pública. Sem env, prefere
  // a pública — bate sempre com o que existe em prod.
  const apiBase = (
    process.env.INTERNAL_API_URL ??
    process.env.NEXT_PUBLIC_API_URL ??
    'https://api.plenyasaude.com.br'
  ).replace(/\/$/, '');
  const url = `${apiBase}/api/v1/score-light/sessions/${encodeURIComponent(code)}`;
  try {
    const res = await fetch(url, { cache: 'no-store' });
    if (res.status === 404) return null;
    if (!res.ok) throw new Error(`API ${res.status}`);
    return (await res.json()) as PublicSession;
  } catch (err) {
    console.error('[escore-light/resultado] fetch failed for', url, err);
    return null;
  }
}

export default async function ResultadoPage({
  params,
}: {
  params: Promise<{ locale: string; code: string }>;
}) {
  const { locale, code } = await params;
  setRequestLocale(locale);

  const session = await fetchSession(code);
  if (!session) notFound();

  return <EscoreLightResultado session={session} locale={locale} />;
}
