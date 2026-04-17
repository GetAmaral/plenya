import { NextResponse } from 'next/server';
import { z } from 'zod';
import { subscribeToNewsletter } from '@/lib/beehiiv';

export const runtime = 'nodejs';

const schema = z.object({
  email: z.string().email(),
  source: z.string().optional(),
});

const RATE_WINDOW_MS = 60_000;
const RATE_MAX = 3;
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
  if (!rateLimit(ip)) return NextResponse.json({ error: 'Too many requests' }, { status: 429 });

  let payload: unknown;
  try {
    payload = await request.json();
  } catch {
    return NextResponse.json({ error: 'Invalid JSON' }, { status: 400 });
  }
  const parsed = schema.safeParse(payload);
  if (!parsed.success) return NextResponse.json({ error: 'Validation failed' }, { status: 422 });

  const result = await subscribeToNewsletter(parsed.data);
  return NextResponse.json({ ok: true, ...result });
}
