import {
  Body,
  Button,
  Container,
  Head,
  Hr,
  Html,
  Link,
  Preview,
  Section,
  Text,
} from '@react-email/components';

/**
 * Template `magic_link` — enviado quando paciente solicita salvar resultado do Escore Light.
 * Variáveis substituídas pelo Go via {{LINK}} e {{SITE_URL}}.
 */
export default function MagicLink() {
  // Os placeholders {{LINK}} e {{SITE_URL}} são strings literais que o backend Go
  // substitui via strings.ReplaceAll após carregar o template via go:embed.
  // Não usar JSX expressions aqui — precisa virar HTML estático com os placeholders intactos.
  const link = '{{LINK}}';
  const siteURL = '{{SITE_URL}}';

  return (
    <Html lang="pt-BR">
      <Head />
      <Preview>Acesse seu Escore Plenya Light</Preview>
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
          <Text style={{ fontSize: 16, marginBottom: 16 }}>Olá,</Text>
          <Text style={{ fontSize: 16, marginBottom: 24 }}>
            Você solicitou guardar seu resultado do <strong>Escore Plenya Light</strong>. Clique
            no botão abaixo para acessar e salvar seu radar.
          </Text>

          <Section style={{ textAlign: 'center', margin: '32px 0' }}>
            <Button
              href={link}
              style={{
                backgroundColor: '#c19a4a',
                color: '#fff8eb',
                padding: '14px 28px',
                textDecoration: 'none',
                display: 'inline-block',
                letterSpacing: '0.5px',
                fontSize: 15,
                fontFamily: 'Helvetica, Arial, sans-serif',
              }}
            >
              Acessar meu resultado
            </Button>
          </Section>

          <Text style={{ fontSize: 13, color: '#4a6478', marginBottom: 8 }}>
            Ou copie este link:
          </Text>
          <Text style={{ fontSize: 13, marginBottom: 24, wordBreak: 'break-all' }}>
            <Link href={link} style={{ color: '#c19a4a' }}>
              {link}
            </Link>
          </Text>

          <Text style={{ fontSize: 13, color: '#4a6478' }}>
            O link expira em <strong>7 dias</strong>. Se você não solicitou, ignore este email.
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
            <strong>LGPD · Seus direitos</strong>
            <br />
            Acesso, correção, portabilidade ou exclusão dos seus dados em{' '}
            <Link href={`${siteURL}/lgpd/direitos`} style={{ color: '#4a6478' }}>
              {siteURL}/lgpd/direitos
            </Link>
            <br />
            Encarregado de Proteção de Dados (DPO):{' '}
            <Link href="mailto:dpo@plenyasaude.com.br" style={{ color: '#4a6478' }}>
              dpo@plenyasaude.com.br
            </Link>
          </Text>
        </Container>
      </Body>
    </Html>
  );
}
