import { NextResponse } from 'next/server';
import sitemap from '@/app/sitemap';

const HOST = 'plenyasaude.com.br';
const KEY = 'c0f02a593783c0e35e95f834bfe5828f';
const KEY_LOCATION = `https://${HOST}/${KEY}.txt`;
const ENDPOINTS = ['https://api.indexnow.org/indexnow'];

async function collectUrls(): Promise<string[]> {
  const entries = await sitemap();
  return entries.map((e) => e.url);
}

export async function POST(req: Request): Promise<NextResponse> {
  const auth = req.headers.get('x-indexnow-auth');
  if (!process.env.INDEXNOW_TRIGGER_SECRET || auth !== process.env.INDEXNOW_TRIGGER_SECRET) {
    return NextResponse.json({ error: 'unauthorized' }, { status: 401 });
  }

  const body = (await req.json().catch(() => ({}))) as { urlList?: string[] };
  const urlList = Array.isArray(body.urlList) && body.urlList.length > 0 ? body.urlList : await collectUrls();

  const payload = {
    host: HOST,
    key: KEY,
    keyLocation: KEY_LOCATION,
    urlList,
  };

  const results = await Promise.all(
    ENDPOINTS.map(async (endpoint) => {
      try {
        const res = await fetch(endpoint, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json; charset=utf-8' },
          body: JSON.stringify(payload),
        });
        return { endpoint, status: res.status };
      } catch (error) {
        return { endpoint, error: String(error) };
      }
    }),
  );

  return NextResponse.json({ submitted: urlList.length, results });
}

export async function GET(): Promise<NextResponse> {
  return NextResponse.json({
    info: 'IndexNow endpoint. POST with header x-indexnow-auth: <INDEXNOW_TRIGGER_SECRET>. Optional body { urlList: [...] }; default = full sitemap.',
    keyLocation: KEY_LOCATION,
  });
}
