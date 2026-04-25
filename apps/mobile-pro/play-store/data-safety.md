# Plenya Pro — Google Play Data Safety form

Estes valores devem ser preenchidos no formulário **App content → Data safety**
do Play Console. Não há um arquivo manifesto auto-importável (Google ainda
não publicou um); este arquivo é a fonte da verdade do que declaramos.

## Visão geral

- O app **coleta** dados pessoais (necessário para funcionamento como EMR).
- O app **NÃO compartilha** dados com terceiros para advertising/marketing.
- O app **NÃO vende** dados pessoais.
- **Toda** a coleta acontece sob consentimento explícito (termo LGPD aceito
  no primeiro login do device, registrado em `/api/v1/me/consent/lgpd`).
- Dados são criptografados em trânsito (HTTPS + cert pinning).
- Dados sensíveis em repouso no device são criptografados (MMKV com chave
  derivada do Keychain/Keystore).
- Usuário pode solicitar exclusão de dados via `/api/v1/me/sessions` e
  pelo perfil web (LGPD Art. 18, V).

## Tipos de dados coletados

### Personal info
- **Name** — coletado, vinculado ao usuário, não para tracking.
  Propósito: App functionality.
- **Email address** — coletado, vinculado, não tracking.
  Propósito: App functionality, Account management.
- **User IDs** — coletado, vinculado, não tracking.
  Propósito: App functionality.

### Health and fitness
- **Health info** — coletado, vinculado, não tracking.
  Inclui: anamnese clínica, escores de saúde, prescrições, resultados
  de exames, avaliações físicas, planos de treino prescritos pelo
  profissional.
  Propósito: App functionality.

### Photos and videos
- **Photos** — coletado, vinculado, não tracking.
  Inclui: fotos de avaliação postural, composição corporal e anexos
  de prontuário tiradas pelo profissional.
  Propósito: App functionality.

### App activity
- **App interactions** — coletado, NÃO vinculado a identidade,
  não tracking. Inclui: telas visitadas, cliques em features.
  Propósito: Analytics.
- **In-app search history** — não coletado.

### App info and performance
- **Crash logs** — coletado, NÃO vinculado, não tracking.
  PHI é stripped no `beforeSend` do Sentry antes do envio.
  Propósito: App functionality, Analytics.
- **Diagnostics** — coletado, NÃO vinculado, não tracking.
  Propósito: App functionality, Analytics.

### Device or other IDs
- **Device or other IDs** — coletado, vinculado, não tracking.
  Inclui: Expo push token (registrado em `/me/device-tokens`).
  Propósito: App functionality.

## Práticas de segurança

- Dados em trânsito: criptografados (HTTPS obrigatório, cert pinning com
  backup pin).
- Dados em repouso (servidor): CPF/RG criptografados AES-256-GCM, audit
  logs imutáveis com retenção de 5 anos.
- Dados em repouso (device): MMKV criptografado, tokens em Android Keystore.
- Dados podem ser deletados a pedido: sim, via perfil web e via solicitação
  para suporte@plenyasaude.com.br (LGPD Art. 18, V; resposta em até 15 dias).
- Compromisso com guidelines da família: não aplicável (app profissional).

## Categoria do Play Store

- **Categoria primária:** Medical
- **Audiência:** apenas adultos profissionais de saúde (autenticação obrigatória)

## Ground truth

Sempre que adicionar um novo `data type` ao backend, atualizar:
1. Este arquivo
2. Apple `PrivacyInfo.xcprivacy`
3. Termo LGPD em `app/(auth)/lgpd-consent.tsx`
4. README do projeto se for mudança de escopo
