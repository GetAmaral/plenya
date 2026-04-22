import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { LEGAL_CONTACT } from '@/lib/legal';
import { LGPDRightsForm } from '@/components/legal/lgpd-rights-form';

export const metadata: Metadata = {
  title: 'Exerça seus direitos LGPD — Plenya',
  description:
    'Solicite acesso, correção, exclusão ou portabilidade dos seus dados pessoais. Direitos do titular conforme art. 18 da LGPD.',
};

export default async function LGPDRightsPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">LGPD · Art. 18</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Seus direitos como titular</h1>
          <p className="text-cream/70 mt-6 max-w-2xl">
            Você pode solicitar a qualquer momento, gratuitamente, o exercício dos direitos
            previstos no art. 18 da Lei Geral de Proteção de Dados.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-narrow section space-y-10">

          <div className="text-petrol/85 text-base leading-relaxed space-y-4">
            <p>
              Use o formulário abaixo para enviar sua solicitação. Ela é encaminhada
              diretamente ao nosso <strong>Encarregado de Proteção de Dados</strong>, que responde em
              até <strong>15 dias úteis</strong>.
            </p>
            <p>
              Para mais informações sobre como tratamos seus dados, consulte a{' '}
              <Link href="/privacidade" className="text-gold underline underline-offset-4">Política de Privacidade</Link>{' '}
              e a página do{' '}
              <Link href="/lgpd/encarregado" className="text-gold underline underline-offset-4">Encarregado</Link>.
            </p>
          </div>

          <LGPDRightsForm dpoEmail={LEGAL_CONTACT.dpoEmail} />

          <details className="text-sm text-petrol/70 border-t border-petrol/10 pt-6">
            <summary className="cursor-pointer hover:text-petrol">Outras formas de contato</summary>
            <p className="mt-3">
              Você também pode escrever diretamente para{' '}
              <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a>{' '}
              informando: (1) seu nome, (2) email/identificador, (3) qual direito deseja exercer
              e (4) detalhes sobre o pedido.
            </p>
            <p className="mt-3">
              Se preferir, registre reclamação direta à <strong>Autoridade Nacional de Proteção de Dados (ANPD)</strong>{' '}
              em <a href="https://www.gov.br/anpd" target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">gov.br/anpd</a>.
            </p>
          </details>
        </div>
      </section>
    </>
  );
}
