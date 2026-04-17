import type { Author } from '@/lib/authors';
import type { Post } from '@/lib/blog';
import { brand } from '@plenya/brand';

export function ArticleSchema({ post, author, url }: { post: Post; author: Author; url: string }) {
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalScholarlyArticle',
    headline: post.title,
    description: post.excerpt,
    datePublished: post.date,
    dateModified: post.updated ?? post.date,
    inLanguage: post.locale === 'pt' ? 'pt-BR' : post.locale,
    mainEntityOfPage: url,
    author: {
      '@type': 'Person',
      name: author.name,
      url: `${brand.url}/equipe/${author.slug}`,
    },
    publisher: {
      '@type': 'MedicalOrganization',
      name: brand.legalName,
      url: brand.url,
      logo: {
        '@type': 'ImageObject',
        url: `${brand.url}/logo.svg`,
      },
    },
    keywords: post.tags.join(', '),
  };
  return <script type="application/ld+json" dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }} />;
}
