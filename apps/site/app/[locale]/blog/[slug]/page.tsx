import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { brand } from '@plenya/brand';
import { defaultLocale, isLocale, locales } from '@/lib/i18n/config';
import { getAllPosts, getPost, pillarLabels } from '@/lib/blog';
import { getAuthor } from '@/lib/authors';
import { MdxContent } from '@/components/blog/mdx-content';
import { AuthorBox } from '@/components/blog/author-box';
import { WhatsAppShare } from '@/components/blog/whatsapp-share';
import { BlogCTA } from '@/components/blog/blog-cta';
import { BlogCTARecognition } from '@/components/blog/blog-cta-recognition';
import { PostTrackReader } from '@/components/blog/post-track-reader';
import { ArticleSchema } from '@/components/blog/article-schema';
import { NewsletterInline } from '@/components/blog/newsletter-inline';
import { Link } from '@/lib/i18n/navigation';

export async function generateStaticParams() {
  const all = await Promise.all(locales.map((l) => getAllPosts(l)));
  return all.flatMap((posts, i) => posts.map((p) => ({ locale: locales[i], slug: p.slug })));
}

export async function generateMetadata({ params }: { params: Promise<{ locale: string; slug: string }> }): Promise<Metadata> {
  const { locale: rawLocale, slug } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  const post = await getPost(locale, slug);
  if (!post) return {};
  return {
    title: post.title,
    description: post.excerpt,
    openGraph: { type: 'article', title: post.title, description: post.excerpt, publishedTime: post.date, modifiedTime: post.updated ?? post.date, authors: [post.author], tags: post.tags },
  };
}

export default async function BlogPostPage({ params }: { params: Promise<{ locale: string; slug: string }> }) {
  const { locale: rawLocale, slug } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);
  const post = await getPost(locale, slug);
  if (!post) notFound();

  const author = await getAuthor(post.author);
  if (!author) notFound();
  const reviewedBy = post.reviewedBy ? await getAuthor(post.reviewedBy) : null;
  const url = `${brand.url}${locale === 'pt' ? '' : `/${locale}`}/blog/${post.slug}`;

  return (
    <article className="bg-cream">
      <ArticleSchema post={post} author={author} url={url} />
      <PostTrackReader slug={post.slug} pillar={post.pillar} />

      {/* Post header */}
      <div className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <Link href={`/blog/categoria/${post.pillar}`} className="label-upper text-gold hover:text-gold-300 transition">
            ← {pillarLabels[post.pillar]}
          </Link>
          <h1 className="heading-hero text-[clamp(2.2rem,5vw,4rem)] text-cream mt-6">{post.title}</h1>
          <div className="flex flex-wrap items-center gap-3 label-upper text-cream/50 mt-6">
            <span>Por {author.name}</span>
            <span>·</span>
            <time dateTime={post.date}>
              {new Date(post.date).toLocaleDateString('pt-BR', { day: '2-digit', month: 'long', year: 'numeric' })}
            </time>
            {post.updated && (
              <><span>·</span><span>Atualizado {new Date(post.updated).toLocaleDateString('pt-BR')}</span></>
            )}
            <span>·</span>
            <span>{post.readingMinutes} min</span>
          </div>
        </div>
      </div>

      {/* TL;DR */}
      <div className="site-narrow pt-16">
        <div className="border-l-2 border-gold pl-6 py-3">
          <p className="label-upper text-gold mb-2">TL;DR</p>
          <p className="text-petrol/85 text-base leading-relaxed">{post.excerpt}</p>
        </div>
      </div>

      {/* Body */}
      <div className="site-narrow section prose-plenya">
        <MdxContent source={post.content} />
      </div>

      {/* References */}
      {post.references.length > 0 && (
        <div className="site-narrow border-t border-petrol/10 pt-8 mb-16">
          <p className="label-upper text-gold mb-4">Referências</p>
          <ol className="list-decimal pl-6 space-y-2 text-petrol/80 text-sm">
            {post.references.map((ref) => (
              <li key={ref.url}>
                <a href={ref.url} target="_blank" rel="noreferrer" className="underline decoration-gold underline-offset-4 hover:text-gold">{ref.label}</a>
              </li>
            ))}
          </ol>
        </div>
      )}

      <div className="site-narrow">
        <AuthorBox author={author} reviewedBy={reviewedBy} />
        <NewsletterInline source={`blog-post-${post.slug}`} />
        {post.cta === 'recognition' ? <BlogCTARecognition /> : <BlogCTA />}
        <div className="flex justify-center pt-8 pb-16">
          <WhatsAppShare title={post.title} url={url} />
        </div>
      </div>
    </article>
  );
}
