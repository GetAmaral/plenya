import { NextResponse } from 'next/server';
import { getAllPlenyaPostsFull } from '@/lib/plenya-blog';
import { getAllLectures } from '@/lib/lectures';
import { getAllBooks } from '@/lib/books';

const HOST = 'drgetulioamaralfilho.com.br';
const KEY = '65955bc7265423c52bc95e6a940427d5';
const KEY_LOCATION = `https://${HOST}/${KEY}.txt`;
const ENDPOINTS = ['https://api.indexnow.org/indexnow', 'https://www.bing.com/indexnow'];

const STATIC_PATHS = ['', '/sobre', '/livros', '/palestras', '/artigos', '/onde-atendo', '/ensino', '/contato'];

async function collectUrls(): Promise<string[]> {
  const urls: string[] = [];

  for (const p of STATIC_PATHS) {
    urls.push(`https://${HOST}${p || '/'}`);
    urls.push(`https://${HOST}/en${p}`);
  }

  const posts = await getAllPlenyaPostsFull('pt');
  for (const post of posts) {
    urls.push(`https://${HOST}/escritos/${post.slug}`);
    urls.push(`https://${HOST}/en/escritos/${post.slug}`);
  }

  const books = await getAllBooks();
  for (const b of books) {
    urls.push(`https://${HOST}/livros/${b.slug}`);
    urls.push(`https://${HOST}/en/livros/${b.slug}`);
  }

  const lectures = await getAllLectures();
  for (const l of lectures) {
    urls.push(`https://${HOST}/palestras/${l.slug}`);
    urls.push(`https://${HOST}/en/palestras/${l.slug}`);
  }

  return urls;
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
