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
 * Template `boas_vindas` — disparado quando Lead vira Patient (claim do Light confirmado
 * ou conversão manual no admin). Variáveis: {{NAME}}, {{SITE_URL}}.
 */
export default function BoasVindas() {
  const name = '{{NAME}}';
  const siteURL = '{{SITE_URL}}';

  return (
    <Html lang="pt-BR">
      <Head />
      <Preview>Bem-vindo à Plenya — sua conta está pronta</Preview>
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
          <Text style={{ fontSize: 22, marginBottom: 24, color: '#1f3640' }}>
            Olá, {name}.
          </Text>

          <Text style={{ fontSize: 16, marginBottom: 16 }}>
            Sua conta na Plenya está pronta. Agora você pode acompanhar seu Escore Plenya Light,
            refazer a avaliação a cada 3 meses e ver sua evolução ao longo do tempo.
          </Text>

          <Text style={{ fontSize: 16, marginBottom: 24 }}>
            Quando quiser conversar com a equipe — sobre uma avaliação completa no Continuum, ou
            sobre os pontos do seu radar que mais chamaram atenção — estamos no WhatsApp ou no
            email <Link href="mailto:contato@plenyasaude.com.br">contato@plenyasaude.com.br</Link>.
          </Text>

          <Text style={{ fontSize: 16, marginBottom: 8 }}>
            Sem pressa, sem cobrança. Quando fizer sentido pra você.
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
            Você está recebendo este email porque criou conta na Plenya. Pra parar de receber
            comunicações, escreva pro DPO em{' '}
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
