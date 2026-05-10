import { brand } from '@plenya/brand';
import { getAllPosts, getPillarLabels } from '@/lib/blog';

const BASE = brand.url;

function escapeXml(v: string): string {
  return v
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&apos;');
}

export async function GET(): Promise<Response> {
  const posts = await getAllPosts('pt');
  const labels = getPillarLabels('pt');
  const lastBuild = new Date().toUTCString();

  const items = posts
    .map((p) => {
      const url = `${BASE}/blog/${p.slug}`;
      const pubDate = new Date(p.updated ?? p.date).toUTCString();
      const category = labels[p.pillar];
      return `    <item>
      <title>${escapeXml(p.title)}</title>
      <link>${url}</link>
      <guid isPermaLink="true">${url}</guid>
      <description>${escapeXml(p.excerpt)}</description>
      <pubDate>${pubDate}</pubDate>
      <category>${escapeXml(category)}</category>
      <author>noreply@plenyasaude.com.br (${escapeXml(p.author)})</author>
    </item>`;
    })
    .join('\n');

  const xml = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>Plenya · Blog</title>
    <link>${BASE}/blog</link>
    <atom:link href="${BASE}/blog/rss.xml" rel="self" type="application/rss+xml"/>
    <description>Medicina funcional integrativa, longevidade e saúde metabólica — artigos da equipe Plenya.</description>
    <language>pt-BR</language>
    <lastBuildDate>${lastBuild}</lastBuildDate>
${items}
  </channel>
</rss>`;

  return new Response(xml, {
    headers: {
      'Content-Type': 'application/rss+xml; charset=utf-8',
      'Cache-Control': 'public, max-age=3600, s-maxage=3600',
    },
  });
}
