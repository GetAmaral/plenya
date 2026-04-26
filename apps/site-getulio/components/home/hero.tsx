import Image from 'next/image';
import { Link } from '@/lib/i18n/navigation';

export function Hero() {
  return (
    <section className="editorial-container pt-12 md:pt-20 pb-16 md:pb-28">
      <div className="grid md:grid-cols-[1.4fr_1fr] gap-10 md:gap-16 items-end">
        <div>
          <p className="label-meta mb-8">Médico · Professor · Autor · Palestrante</p>

          <h1 className="heading-display text-[clamp(2.5rem,7vw,5.5rem)]">
            Dr. Getúlio<br />Amaral Filho
          </h1>

          <div className="mt-12 prose-body max-w-xl">
            <p>
              Vinte anos dentro do hospital me ensinaram que muita doença grave começa
              anos antes — em silêncio, em exames "normais".
            </p>
            <p>
              A medicina que pratico hoje busca esse intervalo. É sobre o que vem antes.
            </p>
          </div>

          <div className="mt-10 flex items-center gap-8 font-sans text-sm">
            <Link href="/sobre" className="link-text">Sobre mim</Link>
            <Link href="/livro" className="link-text">Leia o livro</Link>
          </div>
        </div>

        <div className="relative aspect-[3/4] w-full max-w-md justify-self-end">
          <Image
            src="/images/getulio-hero-bw.jpg"
            alt="Dr. Getúlio Amaral Filho"
            fill
            priority
            className="object-cover grayscale"
            sizes="(min-width: 768px) 480px, 100vw"
          />
        </div>
      </div>
    </section>
  );
}
