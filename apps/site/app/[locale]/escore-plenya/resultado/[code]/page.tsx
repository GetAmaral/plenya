import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { EscoreLightResultado } from '@/components/escore/escore-light-resultado';
import type { PublicSession } from '@/lib/score-light/types';

export const dynamic = 'force-dynamic';

export const metadata: Metadata = {
  title: 'Seu resultado — Escore Plenya Light',
  description: 'Resultado da sua autoavaliação Plenya Light.',
  robots: { index: false, follow: false },
};

async function fetchSession(code: string): Promise<PublicSession | null> {
  // SSR roda dentro do container do site → precisa usar nome do serviço (api:3001),
  // não localhost (que apontaria pro próprio container).
  // Browser usa NEXT_PUBLIC_API_URL (localhost:3001 mapeado pelo Docker).
  const apiBase = (
    process.env.INTERNAL_API_URL ??
    process.env.NEXT_PUBLIC_API_URL ??
    'http://api:3001'
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
