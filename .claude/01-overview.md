# 01 - Visão Geral do Projeto

## Plenya EMR (Electronic Medical Records)

Sistema de prontuário médico eletrônico completo, self-hosted, com foco em baixo custo e compliance LGPD.

## Objetivo

Fornecer solução EMR acessível para clínicas pequenas/médias no Brasil, com:
- Custo operacional < R$100/mês
- Compliance LGPD total
- Sem dependência de fornecedores internacionais
- Self-hosted para controle total dos dados

## Escala

- **Pacientes:** Centenas (não milhares)
- **Profissionais:** 5-20 usuários simultâneos
- **Performance:** Latência < 500ms (p95) para operações comuns
- **Disponibilidade:** 99.5% uptime

## Plataformas

### Web App (Next.js)
- Dashboard principal para profissionais
- CRUD completo de pacientes, anamneses, prescrições
- Visualizações de exames e escores
- Relatórios e PDFs

### Mobile Apps (React Native + Expo)
- iOS e Android nativo
- Acesso para profissionais em atendimento
- Visualização de prontuários
- Upload de documentos

### Backend (Go)
- API REST única para web e mobile
- Autenticação JWT
- RBAC multi-role
- Audit logging

## Diferenciais

1. **Sistema de Escores Integrado**
   - Estratificação de risco automática
   - Baseado em guidelines clínicos
   - Filtros demográficos (idade, gênero, menopausa)

2. **Enriquecimento Clínico**
   - Explicações técnicas (profissionais)
   - Linguagem simples (pacientes)
   - Conduta recomendada
   - Artigos científicos de referência

3. **LGPD by Design**
   - Criptografia de dados sensíveis (CPF, RG)
   - Audit logs imutáveis
   - Rastreabilidade completa
   - Exportação de dados (direito do paciente)

4. **Baixo Custo**
   - VPS Hetzner: ~R$50/mês
   - Self-hosted (sem mensalidades SaaS)
   - PostgreSQL (sem custos de licença)
   - Coolify (alternativa open-source ao Vercel)

## Usuários e Roles

### Profissionais de Saúde
- **Admin:** Configurações, usuários, score groups
- **Doctor:** Anamnese, prescrições, laudos
- **Nurse:** Sinais vitais, administração de medicamentos
- **Psychologist:** Avaliações psicológicas
- **Nutritionist:** Planos alimentares
- **Secretary:** Agendamentos, cadastros básicos

### Pacientes
- **Patient:** Visualização próprio prontuário, download PDFs

## Arquitetura de Alto Nível

```
┌─────────────┐     ┌─────────────┐
│  Web App    │     │ Mobile Apps │
│  Next.js    │     │ React Native│
└──────┬──────┘     └──────┬──────┘
       │                   │
       └───────┬───────────┘
               │
         ┌─────▼─────┐
         │  API Go   │
         │  Fiber    │
         └─────┬─────┘
               │
         ┌─────▼─────┐
         │PostgreSQL │
         │    17     │
         └───────────┘
```

## Contexto Brasileiro

- **LGPD:** Lei Geral de Proteção de Dados (2020)
- **CFM:** Resolução CFM nº 1.821/2007 (prontuário eletrônico)
- **ICP-Brasil:** Assinatura digital (opcional, futuro)
- **Idioma:** Português brasileiro (padrão)
- **Moeda:** Real (R$)
- **Timezone:** America/Sao_Paulo (padrão)

## Roadmap

### Fase 1: Fundação ✅
- Monorepo setup
- Docker development
- Go + PostgreSQL
- Atlas migrations

### Fase 2: Backend Core ✅
- Auth (JWT)
- RBAC
- Models principais
- CRUD básico

### Fase 3: Features Médicas ✅
- Patients
- Anamnesis
- Lab Results
- Prescriptions
- Appointments

### Fase 4: Frontend Web ✅
- Next.js 16.1 setup
- Dashboard
- CRUD completo
- Sistema de Escores

### Fase 5: Mobile Apps (Em andamento)
- Expo setup
- Autenticação
- Visualização prontuários

### Fase 6: Hardening LGPD
- Relatórios de acesso
- Exportação de dados
- Anonimização
- Retention policies

### Fase 7: Deploy Produção
- Hetzner VPS
- Coolify setup
- SSL/TLS
- Backups automatizados

## Custos Mensais Estimados

| Item | Custo |
|------|-------|
| VPS Hetzner CPX31 | €9 (~R$50) |
| Backup Snapshots | €2 (~R$11) |
| S3 Glacier (backups) | ~R$5 |
| Domínio .com.br | ~R$3 |
| **Total Base** | **~R$69/mês** |

Opcionais:
- Expo EAS Updates: $99 USD
- App Stores: ~R$50/mês

## Métricas de Sucesso

- **Performance:** p95 < 500ms (rotas comuns), < 2s (PDFs)
- **Disponibilidade:** > 99.5% uptime
- **Segurança:** 0 breaches, 100% compliance LGPD
- **Custo:** < R$100/mês operacional
- **Satisfação:** NPS > 8/10

## Próximos Passos

1. Completar mobile apps
2. Testes de carga (100 usuários simultâneos)
3. Hardening LGPD (exportação, anonimização)
4. Deploy staging
5. Certificação ICP-Brasil (assinatura digital)
