import { Link } from '@/lib/i18n/navigation';
import { pillarLabels, type Post } from '@/lib/blog';

export function PostCard({ post, featured = false }: { post: Post; featured?: boolean }) {
  return (
    <Link
      href={`/blog/${post.slug}`}
      className={`group block ${featured ? 'post-card-featured' : 'post-card'}`}
    >
      <div className="flex items-center gap-3 label-upper mb-4">
        <span className="text-gold">{pillarLabels[post.pillar]}</span>
        <span className="text-petrol/30">·</span>
        <time className="text-petrol/50" dateTime={post.date}>
          {new Date(post.date).toLocaleDateString('pt-BR', { day: '2-digit', month: 'short', year: 'numeric' })}
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
