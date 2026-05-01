import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link, type Href } from '@/lib/i18n/navigation';
import { LEGAL_CONTACT, PRIVACY_POLICY_VERSION, TERMS_VERSION } from '@/lib/legal';

export const metadata: Metadata = {
  title: 'LGPD — Plenya',
  description:
    'Hub de privacidade e proteção de dados da Plenya. Política de Privacidade, Termos, Encarregado e direitos do titular.',
};

export default async function LGPDHubPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  const cards: Array<{ href: Href; title: string; description: string; meta?: string }> = [
    {
      href: '/privacidade',
      title: 'Política de Privacidade',
      description: 'Que dados coletamos, para que usamos, com quem compartilhamos e por quanto tempo guardamos.',
      meta: `Versão ${PRIVACY_POLICY_VERSION}`,
    },
    {
      href: '/termos',
      title: 'Termos de Uso',
      description: 'Condições para uso do site e do Escore Plenya Light.',
      meta: `Versão ${TERMS_VERSION}`,
    },
    {
      href: '/lgpd/direitos',
      title: 'Exerça seus direitos',
      description: 'Solicite acesso, correção, exclusão, portabilidade ou revogação do consentimento.',
      meta: 'LGPD art. 18',
    },
    {
      href: '/lgpd/encarregado',
      title: 'Encarregado (DPO)',
      description: 'Identidade e contato do Encarregado de Proteção de Dados.',
      meta: 'LGPD art. 41',
    },
  ];

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">LGPD · Lei 13.709/2018</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Privacidade & Proteção de Dados</h1>
          <p className="text-cream/70 mt-6 max-w-2xl">
            A Plenya trata dados pessoais sensíveis (de saúde) com seriedade. Aqui
            você encontra toda a documentação legal e os canais para exercer seus
            direitos.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-narrow section">
          <div className="grid sm:grid-cols-2 gap-px bg-petrol/15 border-y border-petrol/15">
            {cards.map((c) => (
              <Link
                key={c.title}
                href={c.href}
                className="bg-paper p-8 hover:bg-paper/60 transition group"
              >
                {c.meta && <p className="label-upper text-gold mb-3 text-[10px]">{c.meta}</p>}
                <h2 className="heading-section text-petrol text-xl group-hover:text-gold transition">
                  {c.title} →
                </h2>
                <p className="text-petrol/70 text-sm mt-3 leading-relaxed">{c.description}</p>
              </Link>
            ))}
          </div>

          <div className="mt-16 space-y-3 text-petrol/75 text-sm">
            <p>
              <strong>Encarregado:</strong> {LEGAL_CONTACT.dpoName} ·{' '}
              <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a>
            </p>
            <p>
              <strong>Controlador:</strong> {LEGAL_CONTACT.controllerName} · {LEGAL_CONTACT.controllerAddress}
            </p>
            <p>
              <strong>Reclamação à autoridade:</strong>{' '}
              <a href="https://www.gov.br/anpd" target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">
                ANPD — Autoridade Nacional de Proteção de Dados
              </a>
            </p>
          </div>
        </div>
      </section>
    </>
  );
}
