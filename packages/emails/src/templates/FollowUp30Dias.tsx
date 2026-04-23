import {
  Body,
  Container,
  Head,
  Hr,
  Html,
  Link,
  Preview,
  Text,
} from '@react-email/components';

/**
 * Template `follow_up_30_dias` — disparo manual via admin (botão no detail page do Patient)
 * pra quem fez Light há 30+ dias e ainda não converteu em Continuum.
 * Variáveis: {{NAME}}, {{SITE_URL}}.
 */
export default function FollowUp30Dias() {
  const name = '{{NAME}}';
  const siteURL = '{{SITE_URL}}';

  return (
    <Html lang="pt-BR">
      <Head />
      <Preview>Refaça seu Escore — veja sua evolução em 30 dias</Preview>
      <Body
        style={{
          backgroundColor: '#fff8eb',
          fontFamily: '"Cormorant Garamond", Georgia, serif',
          margin: 0,
          padding: '32px 16px',
          color: '#1f3640',
          lineHeight: 1.6,
        }}
      >
        <Container
          style={{
            maxWidth: 560,
            margin: '0 auto',
            backgroundColor: '#ffffff',
            padding: '40px 32px',
            border: '1px solid #e6dfd1',
          }}
        >
          <Text style={{ fontSize: 16, marginBottom: 16 }}>Olá, {name}.</Text>

          <Text style={{ fontSize: 16, marginBottom: 16 }}>
            Faz 30 dias que você fez seu Escore Plenya Light. Pequenas mudanças no dia-a-dia
            já podem aparecer no radar — sono, alimentação, atividade física são os primeiros a
            mexer.
          </Text>

          <Text style={{ fontSize: 16, marginBottom: 24 }}>
            Vale refazer a avaliação?{' '}
            <Link href={`${siteURL}/escore-plenya/avaliar`} style={{ color: '#c19a4a' }}>
              É grátis e leva 7 minutos.
            </Link>
          </Text>

          <Hr style={{ borderColor: '#e6dfd1', margin: '32px 0' }} />

          <Text style={{ fontSize: 14, color: '#4a6478' }}>
            Quer ir além? O <strong>Continuum</strong> é o programa completo da Plenya — exames,
            anamnese aprofundada e plano clínico individualizado pela equipe multidisciplinar. Se
            faz sentido pra você, responda este email ou nos chame no WhatsApp.
          </Text>

          <Text style={{ fontSize: 14, marginTop: 40 }}>
            — Equipe Plenya
            <br />
            <Link href={siteURL} style={{ color: '#4a6478' }}>
              {siteURL}
            </Link>
          </Text>

          <Hr style={{ borderColor: '#e6dfd1', margin: '32px 0' }} />

          <Text style={{ fontSize: 11, color: '#6b7c8a', lineHeight: 1.5 }}>
            Você está recebendo este email porque fez o Escore Plenya Light. Para parar de receber,
            escreva pro DPO em{' '}
            <Link href="mailto:dpo@plenyasaude.com.br" style={{ color: '#4a6478' }}>
              dpo@plenyasaude.com.br
            </Link>
            .
          </Text>
        </Container>
      </Body>
    </Html>
  );
}
