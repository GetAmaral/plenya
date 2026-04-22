# DPIA — Relatório de Impacto à Proteção de Dados Pessoais
## Escore Plenya Light (autoavaliação anônima pública)

| Campo | Valor |
|-------|-------|
| **Data de elaboração** | 2026-04-22 |
| **Versão** | 1.0 |
| **Controlador** | Plenya Saúde Ltda. — Londrina/PR, Brasil |
| **Encarregado (DPO)** | dpo@plenyasaude.com.br |
| **Responsável pela elaboração** | Equipe de Engenharia + Direção Médica |
| **Próxima revisão** | 2026-10-22 (semestral) ou na próxima alteração material do produto |
| **Base legal aplicável** | LGPD art. 11 §1º a (consentimento) + §2º f (tutela da saúde) |

---

## 1. Descrição da operação

### 1.1 O que é
O **Escore Plenya Light** é uma autoavaliação anônima online que gera um radar AGIR (Atividade física, Gestão clínica, Integração mente-corpo, Ritmo circadiano) baseado em respostas que o próprio usuário fornece sobre hábitos, sintomas, histórico de saúde pessoal/familiar, medicamentos em uso e — opcionalmente — valores de exames laboratoriais (digitados ou extraídos de PDF via IA).

### 1.2 Por que existe
- Permitir que pessoas tenham uma fotografia clara da sua saúde antes de uma consulta clínica
- Reduzir fricção de entrada no Continuum Plenya (programa principal)
- Educar o público sobre os 4 pilares do Método AGIR

### 1.3 Quem opera
Plataforma totalmente automatizada. Não há intervenção humana na geração do radar. A equipe Plenya só vê dados se o usuário fizer **claim** (vincular a uma conta via magic link) — e ainda assim só dentro do contexto de consulta clínica.

---

## 2. Dados pessoais tratados

### 2.1 Categorias

| Categoria | Dados específicos | Sensível? | Origem |
|-----------|-------------------|-----------|--------|
| Demográficos | idade, sexo biológico, peso, altura | Não | Auto-declarado |
| **Saúde — autorrelato** | hábitos, sintomas, doenças pessoais, histórico familiar, medicamentos, sono, estresse, alimentação | **SIM** (art. 5º II) | Auto-declarado |
| **Saúde — laboratoriais** | valores de exames (ApoB, LDL, HDL, TSH, Vit D, etc.) | **SIM** | Digitado ou extraído de PDF |
| **Saúde — autotriagem PHQ-9** | respostas individuais sobre humor (depressão) | **SIM** | Auto-declarado |
| Identificadores | email (opcional, se claim) | Não | Voluntário pós-resultado |
| Metadados de consentimento | versão da Política aceita, timestamp UTC | Não | Sistema |

### 2.2 Volume estimado
- Início: ~10-50 sessões/dia
- Cenário 12 meses: até ~200 sessões/dia
- 90 dias de retenção (não-claimed) → população ativa máxima estimada: ~18.000 sessões anônimas + ~5.000 sessões claimed

### 2.3 Crianças e adolescentes
**Não tratamos** dados de menores de 18 anos. Termos de Uso restringem o serviço a maiores de idade.

---

## 3. Finalidades e bases legais

| Finalidade | Base legal LGPD | Justificativa |
|------------|-----------------|---------------|
| Calcular pontuação e radar | Art. 11 §1º a (consentimento específico e destacado) + §2º f (tutela da saúde) | Usuário aceita explicitamente antes de iniciar; checkbox bloqueia avanço; versão da política registrada |
| Persistir sessão por 90 dias | Art. 11 §1º a (consentimento) | Permite acesso futuro via URL |
| Vincular sessão a conta (claim) | Art. 11 §1º a (consentimento) | Usuário fornece email voluntariamente após ver o resultado |
| Enviar magic link por email | Art. 7º V (execução de contrato) | Necessário para entregar o serviço solicitado |
| Métricas agregadas (Plausible) | Art. 7º IX (legítimo interesse — análise de uso) | Sem identificação individual; cookieless |
| Extração automática de PDF (Anthropic) | Art. 11 §1º a (consentimento) + §2º f | Usuário escolhe ativamente fazer upload; aviso destacado |

---

## 4. Princípios LGPD — autoavaliação

| Princípio | Como atendemos | Evidência |
|-----------|----------------|-----------|
| **Finalidade** (art. 6º I) | Cada finalidade é específica e informada na Política | `/privacidade` seção 3 |
| **Adequação** | Dados coletados são compatíveis com o serviço | Lista de itens documentada |
| **Necessidade** | Coletamos só o mínimo necessário; perguntas são opcionais | Form não exige nenhuma resposta específica |
| **Livre acesso** | Sessão é acessível via URL; download por portabilidade sob demanda | Página `/lgpd/direitos` |
| **Qualidade dos dados** | Usuário pode editar a qualquer momento antes de enviar | Toggle/limpar em cada item |
| **Transparência** | Política clara, em português, sem jargão | `/privacidade` revisada |
| **Segurança** | TLS, senhas hashed, tokens curtos, logs sem dados sensíveis | Sec. 7 da Política |
| **Prevenção** | DPIA + revisão periódica + monitoring | Este documento |
| **Não-discriminação** | Não há decisão automatizada com efeito jurídico | N/A |
| **Responsabilização** | Logs de consentimento, versionamento de política, DPO nomeado | Modelo `AnonymousScoreSession` campos `consent_*` |

---

## 5. Compartilhamento e sub-processadores

| Operador | Finalidade | Local | Dados | Salvaguardas |
|----------|------------|-------|-------|--------------|
| **Anthropic (Claude)** | Extração de exames de PDF | EUA | Texto OCR do PDF (efêmero — Anthropic não retém) | Cláusulas-padrão de proteção; zero data retention contratual |
| **OpenAI (DALL-E)** | Geração de ilustrações do site | EUA | Apenas prompts genéricos (sem dados pessoais) | N/A — não recebe dados pessoais |
| **Plausible Analytics** | Métricas agregadas | UE | Hash de IP + URL (sem cookie) | GDPR-compliant |
| **Provedor SMTP** | Envio de magic link | A definir | Email + link único | TLS obrigatório |
| **KingHost / Coolify** | Hospedagem | Brasil | Banco e arquivos | Backups criptografados |

---

## 6. Avaliação de riscos

### 6.1 Matriz

| # | Risco | Probabilidade | Impacto | Risco bruto | Controles | Risco residual |
|---|-------|---------------|---------|-------------|-----------|----------------|
| R1 | Vazamento de banco | Baixa | Alto | Médio | Backup criptografado, acesso restrito, logs de query, ENV de senha não em git | Baixo |
| R2 | URL pública vazada (compartilhamento acidental) | Média | Médio | Médio | Sessão expira em 90d; usuário pode excluir self-service | Baixo |
| R3 | PDF persiste em cache/storage | Baixa | Alto | Médio | `defer os.Remove`; `/tmp/light-lab-uploads/` limpo periodicamente; sem persistência em DB | Baixo |
| R4 | Anthropic retém conteúdo do PDF | Baixa | Alto | Médio | Cláusula contratual; uso de tool_use estruturado (não generativo livre); revisar trimestralmente | Baixo |
| R5 | Logs de aplicação contêm dados sensíveis | Média | Médio | Médio | Auditoria realizada (P2.12); apenas tamanho/contagem nos logs | Baixo |
| R6 | Magic link interceptado em trânsito | Baixa | Médio | Baixo | TLS, link 15min, single-use | Mínimo |
| R7 | Acesso indevido a sessão por adivinhação de publicCode | Mínima | Médio | Baixo | publicCode 12 chars alfanuméricos case-sensitive (entropia ~71 bits) | Mínimo |
| R8 | Reidentificação por cruzamento de demográficos | Média | Médio | Médio | Demográficos limitados (idade, sexo, peso, altura); sem CEP, sem nome | Baixo |
| R9 | Usuário <18 anos usa o serviço | Média | Médio | Médio | Termos exigem 18+; validação de idade no form (mínima 18) | Baixo |
| R10 | Não-resposta a pedido de direito (art. 18) | Baixa | Alto (multa ANPD) | Médio | Formulário público + email DPO + SLA de 15 dias úteis | Baixo |

### 6.2 Riscos residuais aceitos
Após controles, todos os riscos estão em nível **baixo ou mínimo**. A operação é **proporcional ao benefício** (educação em saúde, redução de fricção de entrada no programa principal) e respeita os princípios da LGPD.

---

## 7. Direitos do titular — operacionalização

| Direito | Como atende | Tempo de resposta |
|---------|-------------|-------------------|
| Confirmação | Email para DPO ou formulário `/lgpd/direitos` | 15 dias úteis |
| Acesso | Email com cópia dos dados; ou exportação JSON sob demanda | 15 dias úteis |
| Correção | Manual via DPO (sessão pode ser refeita) | 15 dias úteis |
| Anonimização/eliminação | **Self-service**: botão "Excluir esta sessão" no resultado | Imediato |
| Portabilidade | Email com JSON estruturado | 15 dias úteis |
| Informação sobre compartilhamento | `/privacidade` lista todos os sub-processadores | Imediato (público) |
| Revogação de consentimento | Equivale à exclusão (sem dados, sem tratamento) | Imediato |

---

## 8. Medidas de segurança (art. 46)

### Técnicas
- HTTPS/TLS 1.2+ em todo o tráfego
- Senhas/tokens com bcrypt (Patient/User no EMR)
- JWT curtos (access 15min, refresh 7d)
- Magic link single-use com TTL 15min
- PostgreSQL com backup criptografado
- Logs sem dados sensíveis (auditoria realizada P2.12)
- Rate limiting na API (a configurar para `/extract-labs`)

### Administrativas
- DPO nomeado e canal público
- Política de Privacidade versionada
- Termos de Uso versionados
- Documento de Resposta a Incidentes (P2.11)
- Revisão semestral deste DPIA

### Físicas
- Hospedagem em datacenter Tier III (KingHost)
- Sem cópias locais não controladas

---

## 9. Conclusão

O tratamento de dados realizado pelo Escore Plenya Light é **conforme com a LGPD**. As salvaguardas implementadas reduzem os riscos identificados a níveis aceitáveis, e o serviço atende ao princípio da proporcionalidade (benefício clínico-educacional × dado tratado).

Recomendações para próxima revisão:
1. Configurar rate limiting por IP em `/extract-labs` (5 uploads/hora)
2. Implementar download de portabilidade self-service
3. Revisar contratos com sub-processadores (Anthropic, SMTP) — anexar a este DPIA
4. Integrar consentimento granular (separado para PDF de exames) na próxima iteração

---

**Aprovado por:** Encarregado de Proteção de Dados
**Data:** 2026-04-22
