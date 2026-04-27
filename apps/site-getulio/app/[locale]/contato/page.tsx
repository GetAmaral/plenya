import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { ContactForm } from '@/components/contact/contact-form';

export const metadata: Metadata = {
  title: 'Contato',
  description: 'Entre em contato com Dr. Getúlio Amaral Filho — consultas, palestras, imprensa.',
  alternates: { canonical: '/contato' },
};

export default async function ContatoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Contato</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Como posso ajudar.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Para consultas, escolha a clínica de referência. Para palestras e imprensa, use o
          formulário ou os emails dedicados abaixo.
        </p>
      </header>

      <section className="editorial-container pb-24 grid lg:grid-cols-[2fr_1fr] gap-16">
        <ContactForm />

        <aside className="space-y-10 lg:border-l lg:border-rule lg:pl-12">
          <div>
            <p className="label-meta mb-4">Canais diretos</p>
            <ul className="space-y-3 font-serif text-ink-soft text-sm">
              <li>
                <p className="text-ink-muted text-xs mb-1">Geral</p>
                <a href="mailto:contato@drgetulioamaralfilho.com.br" className="link-text">contato@drgetulioamaralfilho.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">Palestras</p>
                <a href="mailto:palestras@drgetulioamaralfilho.com.br" className="link-text">palestras@drgetulioamaralfilho.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">Imprensa</p>
                <a href="mailto:imprensa@drgetulioamaralfilho.com.br" className="link-text">imprensa@drgetulioamaralfilho.com.br</a>
              </li>
            </ul>
          </div>

          <div>
            <p className="label-meta mb-4">Atendimento clínico</p>
            <ul className="space-y-3 font-serif text-ink-soft text-sm">
              <li>
                <p className="text-ink-muted text-xs mb-1">Plenya · longevidade</p>
                <a href="https://plenyasaude.com.br" target="_blank" rel="noreferrer" className="link-text">plenyasaude.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">Nefroclínica · nefrologia</p>
                <a href="https://nefroclinica.com" target="_blank" rel="noreferrer" className="link-text">nefroclinica.com</a>
              </li>
            </ul>
          </div>

          <div>
            <p className="label-meta mb-4">Local</p>
            <p className="font-serif text-ink-soft text-sm">Londrina · Paraná · Brasil</p>
          </div>
        </aside>
      </section>
    </article>
  );
}
