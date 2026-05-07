import Image from 'next/image';
import { Link } from '@/lib/i18n/navigation';
import { Wordmark } from '@/components/brand/wordmark';

/**
 * Hero da home — réplica em escala maior do wordmark institucional sobre
 * fundo cream. Camadas: filete dourado · "DR." · GETÚLIO AMARAL · tagline.
 * Abaixo, identificação clínica e dois CTAs editoriais. Retrato em uma
 * faixa estreita ladeada por filetes dourados.
 */
export function Hero() {
  return (
    <section className="border-b border-rule">
      <div className="editorial-container pt-20 md:pt-32 pb-16 md:pb-24 text-center">
        <Wordmark size="xl" />

        <p className="label-meta mt-12 text-ink-muted">
          Médico nefrologista · CRM-PR 21.876 · RQE 16.038
        </p>

        <div className="mt-12 mx-auto max-w-xl space-y-5 prose-body text-ink-soft">
          <p>
            Vinte anos dentro do hospital me ensinaram que muita doença grave começa anos
            antes — em silêncio, em exames "normais".
          </p>
          <p>
            A medicina que pratico hoje busca esse intervalo. É sobre o que vem antes — e
            sobre conduzir, no longo prazo, a saúde inteira da pessoa.
          </p>
        </div>

        <div className="mt-12 flex items-center justify-center gap-5 flex-wrap">
          <Link href="/livro" className="btn-gold">
            Leia o livro
          </Link>
          <Link href="/sobre" className="btn-outline">
            Sobre mim
          </Link>
        </div>
      </div>

      {/* Retrato editorial — abaixo do hero, em faixa estreita */}
      <div className="editorial-container pb-16 md:pb-20 grid md:grid-cols-[1fr_320px_1fr] gap-10 items-center">
        <div className="hidden md:block h-px bg-gold/40" aria-hidden="true" />
        <div className="relative aspect-[3/4] w-full max-w-[320px] mx-auto">
          <Image
            src="/images/getulio-hero-bw.jpg"
            alt="Dr. Getúlio Amaral Filho"
            fill
            priority
            className="object-cover grayscale"
            sizes="(min-width: 768px) 320px, 70vw"
          />
        </div>
        <div className="hidden md:block h-px bg-gold/40" aria-hidden="true" />
      </div>
    </section>
  );
}
