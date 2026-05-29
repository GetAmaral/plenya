import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { TERMS_VERSION, LEGAL_CONTACT } from '@/lib/legal';
import { LegalEnNotice } from '@/components/legal/legal-en-notice';

export const metadata: Metadata = {
  title: 'Termos de Uso — Plenya',
  description: 'Termos de Uso do site e dos serviços online da Plenya.',
  alternates: {
    canonical: '/termos',
    languages: {
      'pt-BR': '/termos', pt: '/termos',
      en: '/en/terms',
      'x-default': '/termos',
    },
  },
};

export default async function TermsPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <LegalEnNotice />
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">Legal</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Termos de Uso</h1>
          <p className="text-cream/70 mt-6 max-w-2xl">
            Condições para uso do site da Plenya, do Escore Plenya Light e demais serviços online.
          </p>
          <p className="text-cream/50 text-sm mt-4 font-mono">Versão {TERMS_VERSION}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-narrow section space-y-10 text-petrol/85 text-base leading-relaxed">

          <Section n="1" title="Aceitação">
            <p>
              Ao acessar ou usar qualquer recurso do site da Plenya — incluindo o <strong>Escore Plenya Light</strong>, formulários de contato e o portal do paciente — você concorda com estes Termos. Se não concordar, não use o serviço.
            </p>
          </Section>

          <Section n="2" title="Idade mínima">
            <p>
              Os serviços online da Plenya são destinados a pessoas com <strong>18 anos ou mais</strong>. Não solicitamos nem coletamos dados de menores de idade. Se você é responsável por um menor, entre em contato direto com nossa equipe para orientações.
            </p>
          </Section>

          <Section n="3" title="O Escore Plenya Light não é diagnóstico médico">
            <div className="bg-paper border-l-2 border-gold pl-4 py-3">
              <p>
                <strong>Importante:</strong> o Escore Plenya Light é uma <strong>autoavaliação informativa</strong>, baseada em respostas que você mesmo fornece. <strong>Não é diagnóstico, não substitui consulta médica</strong> e não deve ser usado para tomar decisões clínicas isoladamente.
              </p>
              <p className="mt-3">
                Se você tem sintomas, dúvidas ou suspeitas sobre sua saúde, procure um profissional de saúde qualificado. Em emergências, busque pronto-atendimento ou ligue para o SAMU (192).
              </p>
            </div>
            <p className="mt-4">
              A pontuação e o radar gerados pela ferramenta servem como ponto de partida para uma conversa clínica, não como conclusão diagnóstica. Resultados podem ter limitações por:
            </p>
            <ul className="space-y-1 ml-5 list-disc marker:text-gold">
              <li>Imprecisão das respostas autorrelatadas;</li>
              <li>Subjetividade de percepções (ex.: sintomas, qualidade de sono);</li>
              <li>Erros de extração automática de PDFs de exames;</li>
              <li>Faixas de referência que podem não refletir sua condição clínica individual.</li>
            </ul>
          </Section>

          <Section n="4" title="Uso adequado">
            <p>Você concorda em:</p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Fornecer informações verdadeiras ao usar a autoavaliação.</li>
              <li>Não usar o serviço para fins ilegais, enganosos ou abusivos.</li>
              <li>Não tentar acessar dados de outros usuários, burlar mecanismos de segurança ou prejudicar a infraestrutura.</li>
              <li>Não fazer upload de PDFs que não sejam seus exames laboratoriais pessoais.</li>
            </ul>
          </Section>

          <Section n="5" title="Propriedade intelectual">
            <p>
              Todo o conteúdo do site (textos, imagens, marca, layout, código, ilustrações) é de propriedade da Plenya ou de seus licenciantes. Você pode visualizar, compartilhar links e citar trechos com atribuição. Reprodução comercial sem autorização é vedada.
            </p>
            <p className="mt-3">
              Os <strong>seus dados pessoais</strong> (respostas, valores de exames) permanecem seus — usamos apenas para entregar o serviço, conforme nossa{' '}
              <Link href="/privacidade" className="text-gold underline underline-offset-4">Política de Privacidade</Link>.
            </p>
          </Section>

          <Section n="6" title="Disponibilidade do serviço">
            <p>
              Empenhamo-nos em manter o site funcionando, mas <strong>não garantimos disponibilidade ininterrupta</strong>. Pode haver indisponibilidade por manutenção, atualizações ou falhas técnicas. Não nos responsabilizamos por perdas decorrentes de indisponibilidade temporária.
            </p>
          </Section>

          <Section n="7" title="Limitação de responsabilidade">
            <p>
              Na máxima extensão permitida em lei, a Plenya não responde por:
            </p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Decisões de saúde tomadas com base apenas na autoavaliação Light;</li>
              <li>Consequências do uso indevido do serviço por terceiros que tenham obtido acesso a sua sessão (proteja seu link de acesso);</li>
              <li>Falhas de extração automática de PDFs (revise sempre os valores antes de enviar);</li>
              <li>Indisponibilidade momentânea da plataforma.</li>
            </ul>
            <p className="mt-3">
              Esta limitação não exclui obrigações legais imperativas (ex.: Código de Defesa do Consumidor, LGPD).
            </p>
          </Section>

          <Section n="8" title="Lei aplicável e foro">
            <p>
              Estes Termos são regidos pelas leis da República Federativa do Brasil. Fica eleito o <strong>foro da Comarca de Londrina/PR</strong> para dirimir quaisquer controvérsias, com exceção das hipóteses em que o consumidor possa optar pelo foro de seu domicílio (art. 101 CDC).
            </p>
          </Section>

          <Section n="9" title="Contato">
            <p>
              Dúvidas sobre estes Termos: <a href={`mailto:${LEGAL_CONTACT.controllerEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.controllerEmail}</a>.
            </p>
            <p>
              Dúvidas sobre proteção de dados: <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a>.
            </p>
          </Section>

          <Section n="10" title="Alterações">
            <p>
              Podemos atualizar estes Termos periodicamente. A versão vigente está sempre disponível nesta página, com a data e número da versão no topo. Continuar usando o serviço após uma atualização indica aceitação dos novos Termos.
            </p>
          </Section>

          <p className="text-petrol/50 text-sm pt-8 border-t border-petrol/10">
            Versão {TERMS_VERSION} · {LEGAL_CONTACT.controllerName} · CNPJ {LEGAL_CONTACT.controllerCnpj} · {LEGAL_CONTACT.controllerAddress}
          </p>
        </div>
      </section>
    </>
  );
}

function Section({ n, title, children }: { n: string; title: string; children: React.ReactNode }) {
  return (
    <section className="space-y-3">
      <h2 className="heading-section text-petrol text-2xl flex items-baseline gap-3">
        <span className="text-gold/60 text-base font-mono">{n}.</span>
        {title}
      </h2>
      <div className="space-y-3">{children}</div>
    </section>
  );
}
