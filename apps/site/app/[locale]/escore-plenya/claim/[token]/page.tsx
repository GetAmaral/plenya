import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { ClaimRedirect } from '@/components/escore/claim-redirect';

export const metadata: Metadata = {
  title: 'Confirmando seu acesso — Escore Plenya',
  robots: { index: false, follow: false },
};

export default async function ClaimPage({
  params,
}: {
  params: Promise<{ locale: string; token: string }>;
}) {
  const { locale, token } = await params;
  setRequestLocale(locale);

  return <ClaimRedirect token={token} />;
}
