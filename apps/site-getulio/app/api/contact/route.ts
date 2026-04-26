import { NextResponse } from 'next/server';
import { z } from 'zod';
import nodemailer from 'nodemailer';

export const runtime = 'nodejs';

const schema = z.object({
  name: z.string().min(2),
  email: z.string().email(),
  phone: z.string().min(8),
  reason: z.enum(['consulta-plenya', 'consulta-nefroclinica', 'palestra', 'imprensa', 'outro']),
  message: z.string().optional(),
});

const reasonLabels = {
  'consulta-plenya': 'Consulta — Plenya',
  'consulta-nefroclinica': 'Consulta — Nefroclínica',
  palestra: 'Palestra',
  imprensa: 'Imprensa',
  outro: 'Outro',
} as const;

// Destinatários por motivo. Em produção, sobrescritíveis por env var
// CONTACT_RECIPIENT_DEFAULT (catch-all) ou CONTACT_RECIPIENT_<MOTIVO>.
// Enquanto Stalwart não recebe @drgetulioamaralfilho, default cai em CONTACT_RECIPIENT_DEFAULT.
const reasonDefaultRouting = {
  'consulta-plenya': 'contato@drgetulioamaralfilho.com.br',
  'consulta-nefroclinica': 'contato@drgetulioamaralfilho.com.br',
  palestra: 'palestras@drgetulioamaralfilho.com.br',
  imprensa: 'imprensa@drgetulioamaralfilho.com.br',
  outro: 'contato@drgetulioamaralfilho.com.br',
} as const;

function recipientFor(reason: keyof typeof reasonDefaultRouting): string {
  const key = `CONTACT_RECIPIENT_${reason.replace(/-/g, '_').toUpperCase()}`;
  return process.env[key] ?? process.env.CONTACT_RECIPIENT_DEFAULT ?? reasonDefaultRouting[reason];
}

// Rate limit em memória — suficiente para uso pessoal
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
    return NextResponse.json({ error: 'Muitas tentativas. Tente em 1 minuto.' }, { status: 429 });
  }

  let payload: unknown;
  try {
    payload = await request.json();
  } catch {
    return NextResponse.json({ error: 'JSON inválido' }, { status: 400 });
  }

  const parsed = schema.safeParse(payload);
  if (!parsed.success) {
    return NextResponse.json({ error: 'Validação falhou' }, { status: 422 });
  }

  const { name, email, phone, reason, message } = parsed.data;
  const to = recipientFor(reason);
  const subject = `[${reasonLabels[reason]}] ${name} · site Getúlio`;
  const body = [
    `Motivo: ${reasonLabels[reason]}`,
    `Nome: ${name}`,
    `Email: ${email}`,
    `Telefone: ${phone}`,
    '',
    'Mensagem:',
    message?.trim() || '(sem mensagem)',
  ].join('\n');

  const host = process.env.SMTP_HOST;
  const port = Number(process.env.SMTP_PORT ?? 587);
  const user = process.env.SMTP_USER;
  const pass = process.env.SMTP_PASS;
  const from = process.env.SMTP_FROM ?? 'no-reply@drgetulioamaralfilho.com.br';

  if (!host || !user || !pass) {
    // Dev fallback — só loga; não falha
    console.log('[contact] SMTP não configurado. Mensagem recebida:');
    console.log({ to, subject, body });
    return NextResponse.json({ ok: true, dev: true });
  }

  try {
    const transport = nodemailer.createTransport({
      host,
      port,
      secure: port === 465,
      auth: { user, pass },
    });
    await transport.sendMail({
      from,
      to,
      replyTo: email,
      subject,
      text: body,
    });
  } catch (err) {
    console.error('[contact] envio falhou', err);
    return NextResponse.json({ error: 'Falha no envio. Tente novamente em instantes.' }, { status: 502 });
  }

  return NextResponse.json({ ok: true });
}
