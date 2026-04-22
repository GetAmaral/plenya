# Plano de Resposta a Incidentes de Segurança da Informação
## Plenya — site público + Escore Plenya Light

| Campo | Valor |
|-------|-------|
| **Versão** | 1.0 — 2026-04-22 |
| **Próxima revisão** | 2026-10-22 |
| **Responsável** | Encarregado de Proteção de Dados (DPO) — dpo@plenyasaude.com.br |
| **Aplicabilidade** | Sistemas públicos: site, Escore Plenya Light, integrações (Anthropic, SMTP, Plausible, KingHost) |
| **Referência legal** | LGPD art. 48 (notificação de incidente) |

---

## 1. Definições

**Incidente de segurança** — qualquer evento adverso confirmado, relacionado à segurança de dados pessoais, que possa acarretar **risco ou dano relevante** aos titulares.

**Exemplos típicos:**
- Acesso não autorizado a banco de dados
- Vazamento de credenciais
- Exposição pública de dados sensíveis
- Comprometimento de PDF de exames durante upload
- Uso indevido por sub-processador (Anthropic, SMTP, etc.)
- Ransomware ou indisponibilidade prolongada com risco de vazamento

**Não é incidente** (mas vale registrar):
- Falha temporária de serviço sem exposição de dados
- Tentativas de invasão bloqueadas pela infraestrutura
- Bug funcional sem impacto em dados pessoais

---

## 2. Severidade

| Nível | Critério | Exemplo |
|-------|----------|---------|
| **Crítico** | Vazamento confirmado de dados sensíveis (saúde) afetando ≥10 titulares | DB exposto, PDFs vazados |
| **Alto** | Acesso não autorizado pontual a dados sensíveis | Acesso a 1-9 sessões via brute force de URL |
| **Médio** | Exposição de dados não-sensíveis ou risco potencial não confirmado | Email ou nome em log público |
| **Baixo** | Violação de controle interno sem impacto direto | Senha fraca de admin |

---

## 3. Fluxo de resposta (5 fases)

### Fase 1 — Detecção e triagem (T+0h a T+1h)

**Quem detecta:**
- Equipe técnica via monitoring/alertas
- Reporte de usuário (formulário ou email DPO)
- Sub-processador notifica (Anthropic, KingHost, etc.)
- Pesquisador externo / bug bounty

**Ação imediata:**
1. Registrar timestamp de detecção
2. Acionar **canal de incidente**: Slack/WhatsApp do DPO + CTO/equipe técnica
3. Definir severidade preliminar (Crítico / Alto / Médio / Baixo)
4. Designar **Coordenador do incidente** (default: DPO)

### Fase 2 — Contenção (T+1h a T+4h)

**Objetivo:** parar a sangria.

Ações por tipo:
- **Vazamento de DB:** revogar credenciais, rotacionar senhas, isolar serviço afetado
- **Token comprometido:** invalidar todos os tokens da sessão/usuário
- **Vulnerabilidade explorada:** publicar patch, bloquear IP atacante
- **Sub-processador:** acionar contato de incidente do contrato (Anthropic, KingHost)

Snapshot forense:
- Logs do período relevante (`docker compose logs --since`)
- Estado do banco (backup do momento do incidente)
- Evidências de origem (IPs, user-agents, payloads)

### Fase 3 — Avaliação de impacto (T+4h a T+24h)

Determinar:
- **Quantos titulares afetados** (precisa? estimativa?)
- **Que tipo de dado** vazou (sensível ou não? art. 5º II?)
- **Em poder de quem** está o dado agora
- **Risco real** ao titular (financeiro, discriminatório, danoso à honra/saúde mental)

Documentar tudo em formulário interno: `incident-YYYY-MM-DD-NNN.md` no `docs/lgpd/incidentes/`.

### Fase 4 — Notificação (T+24h a T+72h)

#### 4.1 Notificar a ANPD
**Critério:** se o incidente puder acarretar risco/dano relevante (art. 48).

**Prazo:** sem prazo fixo na LGPD, mas a ANPD tem orientado **2 dias úteis** desde a tomada de conhecimento.

**Como:** formulário oficial em https://www.gov.br/anpd/pt-br/canais_atendimento/agente-de-tratamento/comunicado-de-incidente-de-seguranca-cis

**Conteúdo mínimo (art. 48 §1º):**
1. Descrição da natureza dos dados afetados
2. Informações sobre os titulares envolvidos (categorias, número aproximado)
3. Indicação das medidas técnicas e de segurança utilizadas
4. Riscos relacionados ao incidente
5. Motivos da demora, se a comunicação não for imediata
6. Medidas adotadas para reverter ou mitigar efeitos

#### 4.2 Notificar titulares afetados
Se houver risco/dano relevante:
- Email direto a cada titular afetado (em linguagem clara, sem jargão técnico)
- Aviso destacado no site (banner) se afetar volume significativo
- Conteúdo: o que aconteceu, que dados, o que estamos fazendo, o que o titular pode/deve fazer

**Template básico:**
> Identificamos em [data] um incidente de segurança que pode ter afetado seus dados pessoais. Os dados envolvidos foram [tipo]. Já tomamos as seguintes medidas: [lista]. Recomendamos que você [ação]. Estamos à disposição em dpo@plenyasaude.com.br.

### Fase 5 — Pós-incidente (T+72h em diante)

1. **Post-mortem** documentado: causa raiz, linha do tempo, decisões tomadas, lições
2. **Ações corretivas** com prazo: patch, novo controle, revisão de processo
3. **Atualização do DPIA** se o incidente revelar risco não previsto
4. **Atualização deste plano** se identificar lacuna no fluxo
5. **Comunicação interna** para a equipe sobre o que mudou

---

## 4. Contatos críticos

| Função | Contato |
|--------|---------|
| Encarregado (DPO) | dpo@plenyasaude.com.br |
| Diretor responsável | (preencher) |
| CTO / Equipe técnica | (preencher) |
| Hospedagem (KingHost) | (preencher canal de suporte 24h) |
| Anthropic security | security@anthropic.com |
| ANPD | https://www.gov.br/anpd |

---

## 5. Cenários ensaiados (tabletop)

### Cenário A — Vazamento de URL pública de sessão
**Trigger:** usuário relata que compartilhou link e agora teme que terceiros tenham visto.
- Severidade: Baixa-Média (1 titular)
- Resposta: orientar usuário a usar self-service de exclusão; oferecer email de confirmação
- Notificação ANPD: não obrigatória (1 titular, sem risco coletivo)

### Cenário B — Banco de dados exposto
**Trigger:** alerta de pgAudit, query suspeita ou tentativa de dump
- Severidade: Crítica
- Resposta: isolar API, revogar credenciais DB, snapshot forense, notificar ANPD em 2d, notificar titulares
- Comunicação: banner no site, email para todos os usuários claimed

### Cenário C — Anthropic notifica vazamento de prompts
**Trigger:** comunicado oficial da Anthropic sobre incidente que afetou dados em trânsito.
- Severidade: Alta-Crítica
- Resposta: avaliar quantos PDFs Light passaram pelo período; notificar titulares afetados; reavaliar contrato
- Notificação ANPD: provável

### Cenário D — Empregado mal-intencionado
**Trigger:** acesso indevido detectado em logs de auditoria do EMR.
- Severidade: Alta-Crítica
- Resposta: bloquear acesso, ação disciplinar, snapshot, notificação ANPD se houver vazamento
- Comunicação: titulares afetados

---

## 6. Métricas e revisão

- **Frequência de revisão:** semestral ou após qualquer incidente
- **Indicadores acompanhados:** tempo de detecção, tempo de contenção, tempo de notificação, número de titulares afetados, recorrência de causa raiz
- **Treinamento:** simulação tabletop anual com a equipe técnica + DPO

---

## 7. Histórico de incidentes

| Data | ID | Severidade | Resumo | Resolvido |
|------|----|------------|--------|-----------|
| (nenhum até a publicação deste plano) | | | | |

---

**Próxima revisão:** 2026-10-22
**Aprovado por:** DPO
