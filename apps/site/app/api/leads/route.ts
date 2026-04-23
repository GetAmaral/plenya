import { NextResponse } from 'next/server';
import { z } from 'zod';
import { sendLeadEmail } from '@/lib/email';
import { PRIVACY_POLICY_VERSION } from '@/lib/legal';

export const runtime = 'nodejs';

const leadSchema = z.object({
  name: z.string().min(2),
  phone: z.string().min(8),
  email: z.string().email(),
  reason: z.string().optional(),
  window: z.string().optional(),
  source: z.string().optional(),
  message: z.string().optional(),
  consentVersion: z.string().optional(),
  newsletterOptIn: z.boolean().optional(),
});

// Converte telefone BR livre (ex: "(11) 99999-9999") para E.164 (+5511999998888).
// Aceita 10 ou 11 dígitos. Retorna null se inválido.
function toBRE164(raw: string): string | null {
  const d = raw.replace(/\D/g, '');
  if (d.length === 10 || d.length === 11) return `+55${d}`;
  if (d.length === 12 || d.length === 13) {
    if (d.startsWith('55')) return `+${d}`;
  }
  return null;
}

const EMR_API_URL =
  process.env.INTERNAL_API_URL?.replace(/\/$/, '') ??
  process.env.NEXT_PUBLIC_API_URL?.replace(/\/$/, '') ??
  'http://api:3001';

// Origens autorizadas a postar no proxy. Bloqueia CSRF do navegador da vítima
// vinda de outros domínios. Em dev permitimos localhost com qualquer porta.
const ALLOWED_ORIGINS = [
  'https://plenyasaude.com.br',
  'https://www.plenyasaude.com.br',
];

function isAllowedOrigin(origin: string | null): boolean {
  if (!origin) return false;
  if (ALLOWED_ORIGINS.includes(origin)) return true;
  if (process.env.NODE_ENV !== 'production') {
    try {
      const url = new URL(origin);
      if (url.hostname === 'localhost' || url.hostname === '127.0.0.1') return true;
    } catch {
      return false;
    }
  }
  return false;
}

const RATE_WINDOW_MS = 60_000;
const RATE_MAX = 5;
const buckets = new Map<string, { count: number; resetAt: number }>();

function rateLimit(ip: string) {
  const now = Date.now();
  const b = buckets.get(ip);
  if (!b || b.resetAt < now) {
    buckets.set(ip, { count: 1, resetAt: now + RATE_WINDOW_MS });
    return true;
  }
  if (b.count >= RATE_MAX) return false;
  b.count++;
  return true;
}

export async function POST(request: Request) {
  const ip =
    request.headers.get('cf-connecting-ip') ??
    request.headers.get('x-forwarded-for')?.split(',')[0]?.trim() ??
    'anon';

  if (!rateLimit(ip)) {
    return NextResponse.json({ error: 'Too many requests' }, { status: 429 });
  }

  // CSRF: aceita só requests vindos do próprio site. Origin é setado pelo navegador
  // automaticamente em POSTs e não pode ser forjado por JS de outro domínio.
  const origin = request.headers.get('origin');
  if (!isAllowedOrigin(origin)) {
    return NextResponse.json({ error: 'Forbidden origin' }, { status: 403 });
  }

  let payload: unknown;
  try {
    payload = await request.json();
  } catch {
    return NextResponse.json({ error: 'Invalid JSON' }, { status: 400 });
  }

  const parsed = leadSchema.safeParse(payload);
  if (!parsed.success) {
    return NextResponse.json(
      { error: 'Validation failed', issues: parsed.error.issues },
      { status: 422 },
    );
  }

  const phoneE164 = toBRE164(parsed.data.phone);
  const consentVersion = parsed.data.consentVersion ?? PRIVACY_POLICY_VERSION;
  const messageParts = [
    parsed.data.message,
    parsed.data.reason ? `Motivo: ${parsed.data.reason}` : null,
    parsed.data.window ? `Janela: ${parsed.data.window}` : null,
  ].filter(Boolean);

  const emrPayload = {
    name: parsed.data.name,
    email: parsed.data.email,
    phone: phoneE164 ?? parsed.data.phone,
    message: messageParts.join(' · ') || undefined,
    metadata: {
      reason: parsed.data.reason,
      window: parsed.data.window,
      source: parsed.data.source ?? 'contact-form',
    },
    consentVersion,
    newsletterOptIn: parsed.data.newsletterOptIn ?? false,
  };

  const [emrResult, emailResult] = await Promise.allSettled([
    fetch(`${EMR_API_URL}/api/v1/leads`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(emrPayload),
    }).then(async (r) => {
      if (!r.ok) throw new Error(`EMR ${r.status}: ${await r.text().catch(() => '')}`);
      return r.json();
    }),
    sendLeadEmail(parsed.data),
  ]);

  if (emrResult.status === 'rejected') {
    console.error('[leads] EMR failed', emrResult.reason);
  }
  if (emailResult.status === 'rejected') {
    console.error('[leads] email failed', emailResult.reason);
  }

  const webhook = process.env.LEADS_WEBHOOK_URL;
  if (webhook) {
    try {
      await fetch(webhook, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          ...(process.env.LEADS_WEBHOOK_SECRET
            ? { Authorization: `Bearer ${process.env.LEADS_WEBHOOK_SECRET}` }
            : {}),
        },
        body: JSON.stringify({ ...parsed.data, receivedAt: new Date().toISOString() }),
      });
    } catch (error) {
      console.error('[leads] webhook failed', error);
    }
  }

  return NextResponse.json({ ok: true });
}
