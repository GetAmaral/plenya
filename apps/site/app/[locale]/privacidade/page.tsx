import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { PRIVACY_POLICY_VERSION, LEGAL_CONTACT } from '@/lib/legal';
import { LegalEnNotice } from '@/components/legal/legal-en-notice';

export const metadata: Metadata = {
  title: 'Política de Privacidade — Plenya',
  description:
    'Política de Privacidade da Plenya em conformidade com a LGPD (Lei Geral de Proteção de Dados — Lei 13.709/2018).',
  alternates: {
    canonical: '/privacidade',
    languages: {
      'pt-BR': '/privacidade', pt: '/privacidade',
      en: '/en/privacy',
      'x-default': '/privacidade',
    },
  },
};

export default async function PrivacyPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <LegalEnNotice />
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">LGPD</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Política de Privacidade</h1>
          <p className="text-cream/70 mt-6 max-w-2xl">
            Como a Plenya coleta, usa, compartilha e protege seus dados pessoais — em
            conformidade com a Lei Geral de Proteção de Dados (Lei 13.709/2018).
          </p>
          <p className="text-cream/50 text-sm mt-4 font-mono">
            Versão {PRIVACY_POLICY_VERSION}
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-narrow section space-y-10 text-petrol/85 text-base leading-relaxed">

          <Section n="1" title="Quem somos">
            <p>
              <strong>{LEGAL_CONTACT.controllerName}</strong> ({LEGAL_CONTACT.controllerAddress}) é a
              controladora dos dados pessoais tratados nos serviços online disponíveis em{' '}
              <a href="https://plenyasaude.com.br" className="text-gold underline underline-offset-4">plenyasaude.com.br</a>,
              incluindo o <strong>Escore Plenya Light</strong> (autoavaliação anônima em
              {' '}<a href="https://plenyasaude.com.br/escore-plenya" className="text-gold underline underline-offset-4">plenyasaude.com.br/escore-plenya</a>)
              {' '}e a <strong>área do paciente</strong> em{' '}
              <a href="https://minha.plenyasaude.com.br" className="text-gold underline underline-offset-4">minha.plenyasaude.com.br</a>
              {' '}(acesso restrito a pacientes cadastrados pela equipe).
            </p>
          </Section>

          <Section n="2" title="Que dados coletamos">
            <p className="mb-3"><strong>No site público</strong> (navegação institucional):</p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Métricas agregadas de uso via <strong>Plausible Analytics</strong> (cookieless, sem identificação individual; processa endereço IP de forma efêmera para distinguir visitantes únicos por dia, sem persistir).</li>
              <li>Dados informados em formulários de contato (nome, email, mensagem) — apenas para responder.</li>
            </ul>

            <p className="mt-6 mb-3"><strong>No Escore Plenya Light</strong> (autoavaliação anônima):</p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li><strong>Demográficos:</strong> idade, sexo biológico, peso e altura (opcionais).</li>
              <li><strong>Respostas sobre saúde:</strong> hábitos, sintomas, histórico de doenças pessoais e familiares, uso de medicamentos, valores de exames laboratoriais (todas opcionais — você responde só o que souber ou quiser).</li>
              <li><strong>PDF de exames (opcional):</strong> se você optar por usar a função de upload, o arquivo é enviado para extração automática dos valores laboratoriais e <strong>imediatamente descartado</strong> após o processamento (em segundos). Apenas os números extraídos ficam guardados.</li>
              <li><strong>Email (opcional):</strong> apenas se você optar por salvar o resultado para comparar evoluções futuras (claim via magic link).</li>
            </ul>

            <p className="mt-6 mb-3"><strong>Na área do paciente</strong> (minha.plenyasaude.com.br):</p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li><strong>Dados de identificação e contato:</strong> nome, CPF, data de nascimento, gênero, telefone, email, endereço (informados pela equipe Plenya no momento do cadastro do prontuário).</li>
              <li><strong>Histórico clínico:</strong> anamnese, prescrições, resultados de exames laboratoriais, avaliações físicas, escores de saúde, programa Continuum (marcos, plano integrado da equipe, boxes recebidos).</li>
              <li><strong>Consultas:</strong> agendamentos passados e futuros, telemedicina (sala Daily.co), confirmações de presença, solicitações de reagendamento.</li>
              <li><strong>Mensagens com a equipe:</strong> conversa via portal entre paciente e clínica.</li>
              <li><strong>Documentos clínicos:</strong> atestados, declarações, encaminhamentos disponibilizados pela equipe.</li>
              <li><strong>Logs de acesso:</strong> data, hora e ação realizada (audit trail) — necessários para conformidade LGPD e detecção de incidentes.</li>
            </ul>

            <p className="mt-6 text-petrol/70 text-sm bg-paper border-l-2 border-gold pl-4 py-2">
              <strong>Atenção:</strong> dados sobre saúde são <strong>dados pessoais sensíveis</strong> nos termos do art. 5º, II, da LGPD. Tratamos esses dados com base em seu <strong>consentimento específico e destacado</strong> (art. 11, §1º, a) e na hipótese de tutela da saúde por procedimento realizado em serviço/profissional de saúde (art. 11, §2º, f).
            </p>
          </Section>

          <Section n="3" title="Para que usamos seus dados">
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Gerar sua <strong>pontuação e radar AGIR</strong> da autoavaliação Light.</li>
              <li>Permitir que você <strong>refaça a avaliação no futuro</strong> e compare a evolução, caso opte por salvar o resultado via email (claim).</li>
              <li>Responder mensagens enviadas via formulário de contato.</li>
              <li>Melhorar o serviço com base em <strong>métricas agregadas anônimas</strong> (Plausible).</li>
            </ul>
            <p className="mt-4">Não usamos seus dados para publicidade, profiling comercial, decisões automatizadas com efeito jurídico, nem vendemos qualquer dado a terceiros.</p>
          </Section>

          <Section n="4" title="Por quanto tempo guardamos">
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li><strong>Sessões Light não-identificadas:</strong> apagadas automaticamente <strong>90 dias</strong> após a criação.</li>
              <li><strong>Sessões Light identificadas (claimed):</strong> mantidas enquanto sua conta existir, ou até você solicitar exclusão.</li>
              <li><strong>PDF de exames enviado:</strong> apagado <strong>em segundos</strong> após a extração — não persiste em nenhum sistema.</li>
              <li><strong>Prontuário (área do paciente):</strong> mantido conforme exigência regulamentar — Resolução CFM 1.821/2007 obriga retenção mínima de <strong>20 anos</strong> a partir do último registro. Após esse prazo (ou mediante solicitação de exclusão analisada pelo DPO), os dados são anonimizados ou eliminados.</li>
              <li><strong>Logs de acesso ao portal (audit trail):</strong> mantidos por <strong>5 anos</strong> para conformidade LGPD e investigação de incidentes.</li>
              <li><strong>Mensagens de contato:</strong> mantidas pelo tempo necessário para resposta e arquivamento (até 24 meses).</li>
              <li><strong>Métricas Plausible:</strong> agregadas, sem identificação individual; mantidas indefinidamente para análise histórica anônima.</li>
            </ul>
          </Section>

          <Section n="5" title="Sub-processadores e transferência internacional">
            <p className="mb-3">Para oferecer o serviço, contratamos parceiros que processam dados em nosso nome, sob contrato e com cláusulas-padrão de proteção:</p>
            <ul className="space-y-3 ml-5 list-disc marker:text-gold">
              <li><strong>Anthropic</strong> (Claude AI — Estados Unidos): processa o texto extraído do PDF de exames para identificar valores laboratoriais. <strong>O conteúdo é descartado após a chamada</strong> (zero retenção contratual).</li>
              <li><strong>OpenAI</strong> (Estados Unidos): usado apenas para gerar ilustrações genéricas do site — <strong>não recebe dados pessoais de usuários</strong>.</li>
              <li><strong>Plausible Analytics</strong> (União Europeia, sob GDPR): métricas agregadas, sem cookies, sem identificação individual.</li>
              <li><strong>Resend</strong> (Estados Unidos): envio de emails transacionais (magic links de acesso, confirmações de consulta, convites de portal). Recebe apenas o email do destinatário e o conteúdo da mensagem; não usa os dados para outros fins.</li>
              <li><strong>Meta Platforms (WhatsApp Business)</strong> (Estados Unidos): quando você opta por receber comunicações por WhatsApp, seu telefone e o conteúdo das mensagens são processados pela infraestrutura do WhatsApp da Meta, sob cláusulas-padrão de proteção e sob a Política de Privacidade do WhatsApp. Você pode revogar esse consentimento a qualquer momento.</li>
              <li><strong>Daily.co</strong> (Estados Unidos): provedor de salas de teleconsulta usadas no portal do paciente. Processa áudio e vídeo durante a sessão; <strong>não grava por padrão</strong>. Sala expira após o atendimento.</li>
              <li><strong>Google Calendar</strong> (Estados Unidos): quando seu médico conecta o Google Calendar, sincronizamos eventos com título opaco (apenas iniciais do paciente, sem PHI) para evitar conflitos de agenda. Você não é convidado como participante; o evento existe apenas no calendário do médico.</li>
              <li><strong>Hospedagem (KingHost / Coolify):</strong> infraestrutura de servidor no Brasil.</li>
            </ul>
            <p className="mt-4 text-sm text-petrol/70">Transferências internacionais ocorrem nas hipóteses do art. 33 da LGPD, com cláusulas contratuais que garantem nível de proteção adequado.</p>
          </Section>

          <Section n="6" title="Seus direitos como titular">
            <p className="mb-3">Você pode exercer a qualquer tempo, gratuitamente, os direitos previstos no art. 18 da LGPD:</p>
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Confirmação da existência de tratamento de seus dados.</li>
              <li>Acesso aos seus dados.</li>
              <li>Correção de dados incompletos, inexatos ou desatualizados.</li>
              <li>Anonimização, bloqueio ou eliminação de dados desnecessários ou tratados em desconformidade.</li>
              <li>Portabilidade a outro fornecedor de serviço.</li>
              <li>Eliminação dos dados pessoais tratados com seu consentimento.</li>
              <li>Informação sobre as entidades com quem compartilhamos seus dados.</li>
              <li>Revogação do consentimento.</li>
            </ul>
            <p className="mt-4">
              Para exercer qualquer desses direitos, use o formulário em{' '}
              <Link href="/lgpd/direitos" className="text-gold underline underline-offset-4">/lgpd/direitos</Link>{' '}
              ou escreva para <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a>.
              Respondemos em até <strong>15 dias úteis</strong>.
            </p>
            <p className="mt-3 text-sm text-petrol/70">
              Pacientes com acesso à área <a href="https://minha.plenyasaude.com.br" className="text-gold underline underline-offset-4">minha.plenyasaude.com.br</a> podem
              exercer portabilidade (<em>baixar JSON com todos os seus dados</em>) e solicitar exclusão diretamente no menu <strong>Perfil → LGPD</strong>.
              Solicitações de exclusão de dados clínicos são analisadas pelo DPO em virtude da retenção mínima exigida pela Resolução CFM 1.821/2007 (20 anos).
            </p>
          </Section>

          <Section n="7" title="Segurança">
            <ul className="space-y-2 ml-5 list-disc marker:text-gold">
              <li>Conexão HTTPS (TLS) em todo o tráfego.</li>
              <li>Banco de dados em ambiente controlado, com backups criptografados.</li>
              <li>Acesso a dados pessoais restrito por papel (RBAC) e auditado.</li>
              <li>Tokens de acesso (magic link) curtos e de uso único.</li>
              <li>Logs de aplicação não contêm respostas de saúde em texto plano.</li>
            </ul>
            <p className="mt-4 text-sm text-petrol/70">Em caso de incidente de segurança que possa acarretar risco ou dano relevante, notificaremos a ANPD e os titulares afetados em prazo razoável (art. 48 LGPD).</p>
          </Section>

          <Section n="8" title="Encarregado de Proteção de Dados (DPO)">
            <p>{LEGAL_CONTACT.dpoName}</p>
            <p>Email: <a href={`mailto:${LEGAL_CONTACT.dpoEmail}`} className="text-gold underline underline-offset-4">{LEGAL_CONTACT.dpoEmail}</a></p>
            <p className="mt-4">
              Mais informações em{' '}
              <Link href="/lgpd/encarregado" className="text-gold underline underline-offset-4">/lgpd/encarregado</Link>.
            </p>
          </Section>

          <Section n="9" title="Cookies e armazenamento local">
            <p>O site da Plenya <strong>não usa cookies de rastreamento publicitário</strong>. Plausible Analytics opera sem cookies. A área do paciente em <a href="https://minha.plenyasaude.com.br" className="text-gold underline underline-offset-4">minha.plenyasaude.com.br</a> usa <strong>localStorage estritamente necessário</strong> para autenticação (JWT de sessão) e preferência de UI (sidebar colapsada/expandida) — sem isso a área restrita não funciona. Você pode limpar a qualquer momento nas configurações do navegador (isso encerra sua sessão).</p>
          </Section>

          <Section n="10" title="Alterações a esta Política">
            <p>Esta política pode ser atualizada para refletir mudanças no serviço, na legislação ou em sub-processadores. A versão vigente é exibida no topo. Mudanças materiais serão comunicadas por aviso destacado no site e/ou por email aos usuários identificados.</p>
          </Section>

          <Section n="11" title="Reclamação à autoridade">
            <p>Você pode apresentar reclamação à <strong>Autoridade Nacional de Proteção de Dados (ANPD)</strong> em <a href="https://www.gov.br/anpd" target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">gov.br/anpd</a>.</p>
          </Section>

          <p className="text-petrol/50 text-sm pt-8 border-t border-petrol/10">
            Última atualização: {new Date(PRIVACY_POLICY_VERSION.split('.')[0]).toLocaleDateString('pt-BR')} · Versão {PRIVACY_POLICY_VERSION}
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
