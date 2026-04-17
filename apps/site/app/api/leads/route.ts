import { NextResponse } from 'next/server';
import { z } from 'zod';
import { sendLeadEmail } from '@/lib/email';
import { sendToRdStation } from '@/lib/rdstation';

export const runtime = 'nodejs';

const leadSchema = z.object({
  name: z.string().min(2),
  phone: z.string().min(8),
  email: z.string().email(),
  reason: z.string().optional(),
  window: z.string().optional(),
  source: z.string().optional(),
  message: z.string().optional(),
});

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

  const [emailResult, rdResult] = await Promise.allSettled([
    sendLeadEmail(parsed.data),
    sendToRdStation(parsed.data),
  ]);

  if (emailResult.status === 'rejected') {
    console.error('[leads] email failed', emailResult.reason);
  }
  if (rdResult.status === 'rejected') {
    console.error('[leads] rd failed', rdResult.reason);
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
