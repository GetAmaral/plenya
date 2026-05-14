import Image from 'next/image';
import { getTranslations } from 'next-intl/server';

const BOOK_URL_PT = 'https://drgetulioamaralfilho.com.br/livros/antes';
const BOOK_URL_EN = 'https://drgetulioamaralfilho.com.br/en/books/antes';

type Tone = 'paper' | 'cream' | 'petrol';

type Props = {
  locale: string;
  /** Bloco contextual antes do título (ex.: "Obra de referência do método"). */
  contextLabel?: string;
  /** Linha contextual abaixo do título (ex.: "Por que o livro acompanha o método AGIR"). */
  contextLine?: string;
  /** Cor de fundo. `paper` (padrão) e `cream` usam tipografia escura; `petrol` inverte. */
  tone?: Tone;
};

/**
 * Bloco visual + JSON-LD textual de menção ao livro ANTES.
 *
 * Reutiliza as chaves de tradução `drGetulio.publication*` (já existentes
 * em `messages/{locale}.json`) — fonte única do título, descrição e CTA do
 * livro. Schema JSON-LD do livro fica em `components/seo/book-schema.tsx`
 * (emitido apenas em /dr-getulio; este componente é só superfície visual
 * + link canônico, sem schema duplicado).
 */
export async function BookReferenceBlock({
  locale,
  contextLabel,
  contextLine,
  tone = 'paper',
}: Props) {
  const t = await getTranslations({ locale, namespace: 'drGetulio' });
  const url = locale === 'en' ? BOOK_URL_EN : BOOK_URL_PT;

  const isPetrol = tone === 'petrol';
  const sectionBg = tone === 'petrol' ? 'bg-petrol' : tone === 'cream' ? 'bg-cream' : 'bg-paper';
  const textBody = isPetrol ? 'text-cream/80' : 'text-petrol/75';
  const textTitle = isPetrol ? 'text-cream' : 'text-petrol';
  const ctaLink = isPetrol
    ? 'text-gold hover:text-cream'
    : 'text-gold hover:text-petrol';
  const contextLineColor = isPetrol ? 'text-cream/60' : 'text-petrol/60';

  return (
    <section className={sectionBg}>
      <div className="site-container section">
        <div className="grid md:grid-cols-[180px_1fr] gap-10 md:gap-14 items-center max-w-3xl">
          <a
            href={url}
            target="_blank"
            rel="noreferrer"
            className="relative aspect-[2/3] w-full max-w-[180px] mx-auto md:mx-0 shadow-xl block"
          >
            <Image
              src="/images/livro-capa.jpg"
              alt={t('publicationCoverAlt')}
              fill
              className="object-cover"
              sizes="180px"
            />
          </a>
          <div className="space-y-4">
            {contextLabel ? (
              <p className="label-upper text-gold">{contextLabel}</p>
            ) : (
              <p className="label-upper text-gold">{t('publicationLabel')}</p>
            )}
            <p className={`font-serif text-xl md:text-2xl ${textTitle} leading-snug`}>
              {t('publicationTitle')}
            </p>
            {contextLine ? (
              <p className={`${contextLineColor} text-sm leading-relaxed`}>{contextLine}</p>
            ) : null}
            <p className={`${textBody} leading-relaxed`}>{t('publicationLine')}</p>
            <p>
              <a
                href={url}
                target="_blank"
                rel="noreferrer"
                className={`font-sans text-sm ${ctaLink} transition underline-offset-4 hover:underline`}
              >
                {t('publicationCta')}
              </a>
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}
