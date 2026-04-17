import Image from 'next/image';
import type { Author } from '@/lib/authors';

export function AuthorBox({ author, reviewedBy }: { author: Author; reviewedBy?: Author | null }) {
  return (
    <aside className="border-y border-petrol/10 py-8 my-12 grid gap-6 sm:grid-cols-[80px_1fr] items-start">
      <div className="relative h-20 w-20 rounded-full overflow-hidden bg-cream-200 border border-petrol/10">
        {author.photo ? (
          <Image src={author.photo} alt={author.name} fill className="object-cover" sizes="80px" />
        ) : (
          <div className="absolute inset-0 flex items-center justify-center label-upper text-petrol/40">FOTO</div>
        )}
      </div>
      <div className="space-y-2">
        <p className="label-upper text-gold">Autor</p>
        <p className="heading-section text-petrol text-xl">{author.name}</p>
        <p className="label-upper text-petrol/55">{author.credentials}</p>
        <p className="text-petrol/75 text-sm leading-relaxed max-w-prose">{author.bio}</p>
        {reviewedBy && (
          <p className="label-upper text-petrol/45 pt-2">
            Revisado por {reviewedBy.name} · {reviewedBy.credentials}
          </p>
        )}
      </div>
    </aside>
  );
}
