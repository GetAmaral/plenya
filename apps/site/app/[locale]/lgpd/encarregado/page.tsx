import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { LEGAL_CONTACT } from '@/lib/legal';
import { LegalEnNotice } from '@/components/legal/legal-en-notice';

export const metadata: Metadata = {
  title: 'Encarregado de Proteção de Dados (DPO) — Plenya',
  description:
    'Identidade e contato do Encarregado de Proteção de Dados da Plenya, conforme art. 41 da LGPD.',
};

export default async function DPOPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <LegalEnNotice />
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">LGPD · Art. 41</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Encarregado de Proteção de Dados</h1>
          <p className="text-cream/70 mt-6 max-w-2xl">
            Pessoa indicada pela Plenya para atuar como canal de comunicação entre o controlador,
            os titulares e a Autoridade Nacional de Proteção de Dados (ANPD).
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-narrow section space-y-10 text-petrol/85 text-base leading-relaxed">

          <div className="border border-petrol/15 bg-paper p-6 rounded-md">
            <p className="label-upper text-gold mb-3 text-[10px]">{LEGAL_CONTACT.dpoTitle}</p>
            <p className="heading-section text-petrol text-2xl mb-1">{LEGAL_CONTACT.dpoName}</p>
            <p className="text-petrol/60 text-sm mb-3">Médico responsável e Encarregado de Proteção de Dados da Plenya</p>
            <p>
              Email: <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a>
            </p>
            <p className="text-petrol/60 text-sm mt-3">
              Tempo médio de resposta: até 15 dias úteis.
            </p>
          </div>

          <section className="space-y-3">
            <h2 className="heading-section text-petrol text-2xl">Quando entrar em contato com o Encarregado</h2>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Para exercer qualquer direito previsto no art. 18 da LGPD (acesso, exclusão, correção, portabilidade, revogação de consentimento, etc).</li>
              <li>Para esclarecer dúvidas sobre como a Plenya trata seus dados pessoais.</li>
              <li>Para solicitar informações sobre as bases legais utilizadas, sub-processadores ou transferências internacionais.</li>
              <li>Para reportar incidentes de segurança ou suspeitas de uso indevido.</li>
              <li>Para registrar reclamações antes de recorrer à ANPD.</li>
            </ul>
          </section>

          <section className="space-y-3">
            <h2 className="heading-section text-petrol text-2xl">Atribuições (art. 41 §2º)</h2>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Receber comunicações da ANPD.</li>
              <li>Receber reclamações e comunicações dos titulares e prestar esclarecimentos.</li>
              <li>Orientar funcionários e contratados sobre as práticas de proteção de dados.</li>
              <li>Executar demais atribuições conforme orientação da ANPD.</li>
            </ul>
          </section>

          <section className="space-y-3">
            <h2 className="heading-section text-petrol text-2xl">Identidade do Controlador</h2>
            <p>
              <strong>{LEGAL_CONTACT.controllerName}</strong><br />
              {LEGAL_CONTACT.controllerAddress}<br />
              Email: <a href={`mailto:${LEGAL_CONTACT.controllerEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.controllerEmail}</a>
            </p>
          </section>

          <div className="bg-paper border-l-2 border-gold pl-4 py-3 text-petrol/75 text-sm">
            <p>
              Para exercer seus direitos online, use o formulário em{' '}
              <Link href="/lgpd/direitos" className="text-gold underline underline-offset-4">/lgpd/direitos</Link>.
            </p>
            <p className="mt-2">
              Você também pode apresentar reclamação à <strong>Autoridade Nacional de Proteção de Dados (ANPD)</strong>:{' '}
              <a href="https://www.gov.br/anpd" target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">gov.br/anpd</a>.
            </p>
          </div>
        </div>
      </section>
    </>
  );
}
