import Image from 'next/image';
import { Link } from '@/lib/i18n/navigation';

export function BookStrip() {
  return (
    <section className="bg-navy text-paper">
      <div className="editorial-container py-20 md:py-28">
        <div className="grid md:grid-cols-[280px_1fr] gap-10 md:gap-16 items-center">
          <div className="relative aspect-[2/3] w-full max-w-[280px] mx-auto md:mx-0 shadow-2xl">
            <Image
              src="/images/livro-capa.jpg"
              alt="ANTES — A Janela Silenciosa entre o Normal e o Ótimo"
              fill
              className="object-cover"
              sizes="(min-width: 768px) 280px, 70vw"
            />
          </div>

          <div className="space-y-6">
            <div className="flex items-center gap-3">
              <span className="filete-gold" aria-hidden="true" />
              <p className="label-meta text-gold">Livro · 2026</p>
            </div>
            <h2 className="heading-section text-paper text-3xl md:text-5xl">
              Antes — A Janela Silenciosa<br />entre o Normal e o Ótimo
            </h2>
            <blockquote className="font-serif text-paper/85 text-xl md:text-2xl italic leading-relaxed border-l-2 border-gold/60 pl-6">
              Entre o que o laboratório chama de normal e o que o corpo consegue como ótimo,
              existem dez a vinte anos de janela silenciosa.
            </blockquote>
            <p className="font-sans text-sm text-paper/60">ISBN 978-65-02-06742-0</p>
            <Link href="/livro" className="link-text-light inline-block font-sans text-sm">
              Ler o livro →
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
