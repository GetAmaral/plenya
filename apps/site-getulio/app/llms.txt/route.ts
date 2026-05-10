import { getAllPlenyaPostsFull } from '@/lib/blog';
import { getAllBooks } from '@/lib/books';

const BASE = 'https://drgetulioamaralfilho.com.br';

export async function GET(): Promise<Response> {
  const posts = await getAllPlenyaPostsFull('pt');
  const books = await getAllBooks();

  const postLines = posts
    .slice(0, 30)
    .map((p) => `- [${p.title}](${BASE}/artigos/${p.slug}) — ${p.excerpt}`)
    .join('\n');

  const bookLines = books
    .map((b) => `- [${b.title}](${BASE}/livros/${b.slug}) — ${b.shortDescription}`)
    .join('\n');

  const body = `# Dr. Getúlio Amaral Filho

Médico nefrologista (CRM-PR 21.876 · RQE 16.038), professor, autor e diretor clínico da Plenya. Londrina-PR, Brasil. Conteúdo editorial sobre medicina guiada por raciocínio clínico, longevidade, integração corpo-mente e gestão metabólica.

Site oficial: ${BASE}
Idiomas: pt-BR (raiz), en (/en)

## Sobre
- Site institucional do médico, complementar ao [Plenya](https://plenyasaude.com.br) (clínica + EMR).
- Conteúdo editorial assinado por Dr. Getúlio. Artigos médicos têm canonical apontando para Plenya.

## Livros
${bookLines}

## Artigos recentes
${postLines}

## Recursos
- Sitemap: ${BASE}/sitemap.xml
- RSS: ${BASE}/artigos.xml
- Person/Physician schema: ${BASE}/#person
- Contato editorial/imprensa: ${BASE}/contato

## Permissão de uso por LLMs
Conteúdo educacional, livre para citação com atribuição ao autor e link para a URL canônica. Não substitui consulta médica.
`;

  return new Response(body, {
    headers: {
      'Content-Type': 'text/plain; charset=utf-8',
      'Cache-Control': 'public, max-age=3600, s-maxage=3600',
    },
  });
}
