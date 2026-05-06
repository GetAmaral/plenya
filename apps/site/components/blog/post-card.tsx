import Image from 'next/image';
import { Link } from '@/lib/i18n/navigation';
import { formatPostDate, getPillarLabels, type Post } from '@/lib/blog';
import { defaultLocale, type Locale } from '@/lib/i18n/config';

export function PostCard({ post, featured = false, locale = defaultLocale }: { post: Post; featured?: boolean; locale?: Locale }) {
  const labels = getPillarLabels(locale);
  return (
    <Link
      href={{ pathname: '/blog/[slug]', params: { slug: post.slug } }}
      className={`group block ${featured ? 'post-card-featured' : 'post-card'}`}
    >
      {post.cover && (
        <div className={`relative overflow-hidden mb-6 bg-petrol/5 ${featured ? 'aspect-[3/2]' : 'aspect-[16/10]'}`}>
          <Image
            src={post.cover}
            alt={post.title}
            fill
            sizes={featured ? '(min-width: 768px) 720px, 100vw' : '(min-width: 768px) 360px, 100vw'}
            className="object-cover transition duration-500 group-hover:scale-[1.02]"
          />
        </div>
      )}
      <div className="flex items-center gap-3 label-upper mb-4">
        <span className="text-gold">{labels[post.pillar]}</span>
        <span className="text-petrol/30">·</span>
        <time className="text-petrol/50" dateTime={post.date}>
          {formatPostDate(post.date, locale)}
        </time>
        <span className="text-petrol/30">·</span>
        <span className="text-petrol/50">{post.readingMinutes} min</span>
      </div>
      <h3
        className={`heading-section text-petrol group-hover:text-gold transition ${
          featured ? 'text-2xl md:text-4xl' : 'text-xl md:text-2xl'
        }`}
      >
        {post.title}
      </h3>
      <p className="text-petrol/65 mt-3 max-w-2xl leading-relaxed">{post.excerpt}</p>
    </Link>
  );
}
