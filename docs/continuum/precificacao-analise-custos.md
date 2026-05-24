# Precificação Continuum Plenya — Análise de Custos

> **Status:** análise pausada na fase de inventário. A quantificação depende de inputs do Getúlio
> (ver seção final). Origem da discussão: sessão de 18/05/2026.
> **Preços atuais no deck (slide 21):** Semestral R$ 25.000 (ou 12× R$ 2.500) · Anual R$ 40.000
> (ou 12× R$ 3.750). Foram definidos antes desta análise — esta análise existe pra validá-los/ajustá-los.

Mapeamento exaustivo das categorias de custo. **65 itens em 10 categorias.** Sem números ainda —
é o inventário do que precisa entrar na conta antes de fixar o preço.

---

## A. Custos diretos com profissionais (variável — o maior custo)

A mão de obra é o maior custo do Continuum por natureza. Quatro profissionais trabalham o mesmo paciente.

**Tempo clínico por paciente, ciclo completo:**
1. Quatro consultas iniciais (uma com cada profissional)
2. Reunião de equipe pra construir plano integrado (4 profissionais sentados juntos)
3. Encontros semanais em rotação (24 no semestral, 48 no anual)
4. Reavaliações trimestrais do Escore (2 ou 4)
5. Avaliação final estruturada
6. WhatsApp em horário comercial (tempo difuso, mas contabilizável)

**Tempo invisível por trás de cada consulta:**
7. Preparação pré-consulta (leitura do prontuário, exames novos)
8. Documentação pós-consulta (registro no EMR, atualização do plano)
9. Discussão entre profissionais sobre casos complexos
10. Pensar sobre o paciente entre encontros (overhead cognitivo do médico-gestor)
11. Revisão de exames externos chegando ao longo do ciclo

**Profissionais envolvidos** (cada um com custo-hora distinto, varia por seniority):
- Médico nefrologista (Getúlio, ativo central)
- Nutricionista
- Psicóloga
- Educador físico

## B. Box Plenya (material físico)

12. Caixa/embalagem (wenge wood, papel, tecido)
13. Suplementos (compra direta ou margem na revenda)
14. Manipulados (parceria com farmácia de manipulação)
15. Mimos personalizados (livro, chá, vela, cartão escrito à mão)
16. Frete (qualquer cidade do Brasil, eventualmente exterior)
17. Personalização (etiqueta, cartão, envelope)
18. Frequência: abertura + sempre que houver indicação (estimar média)

## C. Infraestrutura tecnológica (mista fixa + variável)

**Fixos (rateados por paciente):**
19. EMR Plenya (amortização do desenvolvimento + manutenção)
20. Hosting (Coolify, Hetzner VPS)
21. Domínios e certificados
22. Mailserver Stalwart self-hosted
23. CRM (Resend outbound, Meta WhatsApp inbound)
24. Sentry, observability, backups

**Variáveis por paciente:**
25. Daily.co (minutos de vídeo, scale com 24-48 sessões/paciente)
26. WhatsApp Business API (conversations pricing Meta)
27. Resend (volume email transacional)
28. Storage (exames, fotos, prontuário longo)

## D. Operacional e administrativo

29. Contabilidade (CNPJ Plenya)
30. Impostos (Simples Nacional ou Lucro Presumido, depende da receita projetada)
31. Folha (CLT) ou contratos PJ pros profissionais
32. Pró-labore Getúlio
33. Encargos sociais (INSS, FGTS se CLT)
34. Seguros (responsabilidade civil médica, multi-profissional)
35. Tarifas de gateway de pagamento (Stripe, PagBank ou Asaas)
36. Conciliação financeira / cobrança recorrente
37. Material de escritório, papelaria personalizada

## E. Aquisição de cliente (CAC)

38. Custo de mídia paga (Meta Ads, Google Ads se houver)
39. Conteúdo (tempo produzindo blog, posts LinkedIn/Instagram)
40. Site e manutenção (já existe, mas tem custo contínuo)
41. Materiais comerciais (este deck, brochuras, vídeos)
42. Tempo de Getúlio em conversas comerciais pré-fechamento
43. Eventos, palestras, networking (quando contribuem pro Continuum)
44. Tempo da equipe atendendo leads/dúvidas

## F. Compliance e jurídico

45. Adequação LGPD contínua
46. Termo de consentimento (revisão jurídica)
47. Contratos de prestação de serviço
48. Pareceres jurídicos pontuais (telemedicina CFM 2.314/2022, regulamentação)
49. Auditoria periódica de prontuário e segurança

## G. Educação continuada da equipe

50. Cursos, congressos, certificações pros 4 profissionais
51. Assinatura de bases de dados clínicas (UpToDate, etc)
52. Bibliografia

## H. Custos de oportunidade

53. Cada hora de Getúlio no Continuum é hora que NÃO está em consulta avulsa Plenya, Nefroclínica ou hemodiálise hospitalar
54. Cada vaga ocupada (12/trimestre) é uma vaga indisponível
55. Profissionais com tempo dedicado ao Continuum não estão disponíveis pra outros pacientes

## I. Custos de risco e contingência

56. Inadimplência (parcelamento 12× significa exposição)
57. Reserva pra reembolsos (tema acoplado — política de reembolso depende desta análise)
58. Substituição de profissional doente/em férias
59. Casos clínicos que exigem MAIS tempo que o previsto (paciente complexo)
60. Cliente que pede mais Box / mais sessões fora do escopo (escopo creep)

## J. Custos de onboarding e fechamento (não-recorrentes por ciclo)

61. Avaliação ampliada inicial (custo pesado upfront, 4 profissionais + reunião)
62. Setup do paciente no EMR
63. Primeiro Box (frequentemente o mais elaborado)
64. Relatório final estruturado
65. Encaminhamentos pós-programa

---

---

# QUANTIFICAÇÃO

## Parâmetros editáveis (base de cálculo)

> Altere aqui e os valores das tabelas abaixo seguem essa base. Os honorários por sessão são
> derivados destas taxas-hora.

Os honorários **não** derivam de uma hora-base única. São três taxas por intensidade de sessão,
e o médico é sempre **2,0×** os demais. Editar a base (ou o multiplicador) reajusta tudo.

**Base — profissional não-médico** (nutri, psico, educador físico):

| Tipo de sessão | Valor | Taxa-hora implícita |
|---|---|---|
| Consulta 60min (inicial / apresentação / feedback) | R$ 450 | R$ 450/h |
| Call follow-up 45min | R$ 300 | R$ 400/h |
| Reunião dos 4, 45min | R$ 225 | R$ 300/h |

**Multiplicador do médico:** `2,0×` → consulta 60min R$ 900 · call 45min R$ 600 · reunião R$ 450
(taxas-hora R$ 900 / R$ 800 / R$ 600).

> Lógica das três taxas: a **consulta** (entrega clínica plena) é a mais cara por hora; o **call**
> de follow-up vem abaixo; a **reunião dos 4** é a mais barata por ser coordenação interna, não
> entrega ao cliente.

**Outros parâmetros:**

| Parâmetro | Valor |
|---|---|
| WhatsApp — disponibilidade/mês (médico) | R$ 600/mês (cenário médio) |
| WhatsApp — disponibilidade/mês (cada outro) | R$ 300/mês (cenário médio) |
| Duração do ciclo semestral | 26 semanas (6 meses) |
| Câmbio US$ → R$ | R$ 5,40 |
| Volume de rateio | 24 pacientes/ano |
| Fatia Continuum dos custos compartilhados | 80% |
| Fração EMR do gasto de dev | 60% |
| Horizonte de amortização do EMR | 3 anos (72 pacientes) |
| Licença por paciente | R$ 50/mês |
| Regime tributário | Lucro Presumido (presunção 32%, conservador) |
| ISS Londrina | 3% |
| Carga tributária efetiva s/ receita | 14,33% |
| Gateway Asaas (cartão parcelado) | 2,99% + R$ 0,49 |
| Contabilidade | R$ 800/mês (fixo, rateado 80% ÷ 24) |
| Seguro RC equipe (estimativa) | ~R$ 6.600/ano (rateado 80% ÷ 24) |
| Pró-labore Getúlio (CEO) | R$ 2.000/paciente |
| Fração marketing atribuível ao Continuum | 80% |
| Mentoria Pedro — horizonte de amortização | 2 anos (48 pacientes) |
| Sites — fração da base de dev (não-EMR) | 40% · amort. 3 anos |
| Conteúdo IG + tráfego | R$ 5.800/mês |

---

## A — Honorários dos profissionais · Plano 6 meses (26 semanas)

> Input do Getúlio, 24/05/2026. **Os valores por sessão já englobam o tempo invisível**
> (preparação, documentação no EMR, revisão de exames, discussão de casos) — não há custo
> adicional sobre os itens A.7–A.11. WhatsApp disponível é tratado como item separado (ver abaixo).

### Cronograma (26 semanas)

| Sem | Quem | O quê | Duração |
|----|------|-------|---------|
| 1 | Todos | 4 consultas iniciais + reunião dos 4 (montagem do plano) | 60min cada + 45min |
| 2 | Médico | Apresentação do plano | 60min |
| 3 | Nutri | Call rotação | 45min |
| 4 | Educador | Call rotação | 45min |
| 5 | Psico | Call rotação | 45min |
| 6 | Médico | Call rotação | 45min |
| 7 | Nutri | Call rotação | 45min |
| 8 | Educador | Call rotação | 45min |
| 9 | Psico | Call rotação | 45min |
| 10 | Médico | Call rotação | 45min |
| 11 | Nutri | Call rotação | 45min |
| 12 | Educador | Call rotação | 45min |
| 13 | Psico + **reunião dos 4** | Call + reavaliação meio do plano | 45min + 45min |
| 14 | Médico | Feedback meio | 60min |
| 15 | Nutri | Call rotação | 45min |
| 16 | Educador | Call rotação | 45min |
| 17 | Psico | Call rotação | 45min |
| 18 | Médico | Call rotação | 45min |
| 19 | Nutri | Call rotação | 45min |
| 20 | Educador | Call rotação | 45min |
| 21 | Psico | Call rotação | 45min |
| 22 | Médico | Call rotação | 45min |
| 23 | Nutri | Call rotação | 45min |
| 24 | Educador | Call rotação | 45min |
| 25 | Psico + **reunião dos 4** | Call + reavaliação final | 45min + 45min |
| 26 | Médico | Feedback final | 60min |

### Custo por profissional / ciclo

**Médico**
| Item | Qtd | Valor | Subtotal |
|---|---|---|---|
| Consulta inicial 60min | 1 | 900 | 900 |
| Apresentação do plano 60min (S2) | 1 | 900 | 900 |
| Reavaliação 60min (S14, S26) | 2 | 900 | 1.800 |
| Call rotação 45min (S6, 10, 18, 22) | 4 | 600 | 2.400 |
| Reunião dos 4 45min (S1, 13, 25) | 3 | 450 | 1.350 |
| **Total médico** | | | **R$ 7.350** |

**Nutricionista / Psicóloga / Educador físico** (estrutura idêntica entre os três)
| Item | Qtd | Valor | Subtotal |
|---|---|---|---|
| Consulta inicial 60min | 1 | 450 | 450 |
| Call rotação 45min | 6 | 300 | 1.800 |
| Reunião dos 4 45min (S1, 13, 25) | 3 | 225 | 675 |
| **Total cada** | | | **R$ 2.925** |

### Totais do ciclo (6 meses)

| | Valor |
|---|---|
| Médico | R$ 7.350 |
| Nutricionista | R$ 2.925 |
| Psicóloga | R$ 2.925 |
| Educador físico | R$ 2.925 |
| **Subtotal honorários de sessões / ciclo** | **R$ 16.125** |

### WhatsApp disponível (retainer mensal de disponibilidade)

> **Lógica adotada:** retainer fixo mensal por profissional, **não** por mensagem nem por minuto —
> espelha o modelo concierge real (a mensagem assíncrona está embutida na disponibilidade vendida).
> Cenário **médio** travado. Trava operacional: horário comercial + SLA de resposta definido
> (urgência clínica não é canal de WhatsApp).

| Profissional | Valor/mês | × 6 meses |
|---|---|---|
| Médico | R$ 600 | R$ 3.600 |
| Nutricionista | R$ 300 | R$ 1.800 |
| Psicóloga | R$ 300 | R$ 1.800 |
| Educador físico | R$ 300 | R$ 1.800 |
| **Subtotal WhatsApp / ciclo** | | **R$ 9.000** |

### Total do bloco A / ciclo

| | Valor |
|---|---|
| Honorários de sessões | R$ 16.125 |
| WhatsApp disponível (médio) | R$ 9.000 |
| **TOTAL BLOCO A / ciclo** | **R$ 25.125** |
| Tempo clínico faturado em sessões | ~32,5h (médico 9,25h · cada outro 7,75h) |
| Toques com o cliente | 29 (4 iniciais + 25 semanais) |
| Reuniões internas dos 4 | 3 (S1, S13, S25) |

---

## B — Box Plenya · 4 boxes por ciclo (6 meses)

> Input do Getúlio, 24/05/2026. **4 boxes/ciclo:** abertura · mês 1 · mês 3 (pós-reavaliação) ·
> mês 6 (final). Envio Brasil apenas (sem internacional por enquanto). Mimos ainda em curadoria.

| Item | Cálculo | Custo/ciclo |
|---|---|---|
| Embalagem/caixa | R$ 100 × 4 boxes | R$ 400 |
| Suplemento magnésio inositol — True Source (1 lata/box) | R$ 180 × 4 | R$ 720 |
| Manipulado 3 componentes (R$ 180 + 27 + 38 = R$ 245/mês) | R$ 245 × 6 meses | R$ 1.470 |
| Frete nacional *(estimativa)* | ~R$ 60 × 4 | R$ 240 |
| Mimos *(a curar — cenário médio)* | progressivo (ver abaixo) | R$ 530 |
| Personalização (etiqueta, cartão, envelope) | ~R$ 15 × 4 | R$ 60 |
| **TOTAL BLOCO B / ciclo** | | **~R$ 3.420** |

### Mimos — curadoria progressiva (cenário médio R$ 530)

> Princípio: poucos itens, bem feitos, anti-"drip bar". Trunfo proprietário de custo quase zero:
> **livro assinado do Dr. Getúlio** (custo de impressão, valor percebido alto). Mimos progressivos,
> não iguais. Faixa de calibração: enxuto ~R$350 · médio ~R$530 · generoso ~R$750.

| Box | Mimos sugeridos | Custo est. |
|---|---|---|
| Abertura (o "wow") | Livro assinado do Dr. Getúlio + chá funcional artesanal + cartão escrito à mão | ~R$ 200 |
| Mês 1 | Item único curado (vela de soja / máscara de dormir em linho) | ~R$ 80 |
| Mês 3 | Item de ritual (infusor / garrafa de vidro premium) | ~R$ 100 |
| Mês 6 (celebração) | Item de fechamento marcante (peça em linho + nota de progresso) | ~R$ 150 |

### Assunções a revisitar (aprovadas como base em 24/05, mas editáveis)

- **Manipulado:** suprimento mensal ininterrupto, R$ 245/mês × 6 meses = R$ 1.470.
- **Magnésio:** literal "1 lata por box" = 4 latas. Caveat: o vão mês3→mês6 (3 meses, 1 box) pode
  exigir mais latas se 1 lata = 1 mês de consumo. Revisar quando definir posologia.
- **Frete R$ 60/box:** estimativa para caixa premium ~1–2kg via Sedex/transportadora. Trocar por valor real.
- **Mimos R$ 530:** placeholder até a curadoria fechar.

---

## C — Infraestrutura tecnológica

> Input do Getúlio, 24/05/2026. Volume de rateio: **24 pacientes/ano**. Continuum carrega **80%**
> dos fixos compartilhados (o EMR serve toda a Plenya, Continuum banca a maior fatia).

### Fixos compartilhados (todo o EMR Plenya) — anual

| Item | Custo/ano |
|---|---|
| VPS KingHost (R$ 200/mês) | R$ 2.400 |
| Domínios (2 × R$ 80/ano) | R$ 160 |
| Resend / Daily.co / Sentry / backups | R$ 0 (free tier) |
| **Total fixo compartilhado/ano** | **R$ 2.560** |

**Rateio:** R$ 2.560 × 80% = R$ 2.048 ÷ 24 pacientes = **~R$ 85/paciente/ciclo**

### Variáveis por paciente

| Item | Estimativa/ciclo |
|---|---|
| Daily.co (vídeo, ~29 sessões) | ~R$ 30 |
| WhatsApp Business API (maioria "service", grátis) | ~R$ 30 |
| Resend (transacional) | ~R$ 5 |
| Storage | ~R$ 5 |
| **Licença (R$ 50/paciente/mês × 6 meses)** | **R$ 300** |
| **Total variável/paciente** | **~R$ 370** |

### Amortização do desenvolvimento do EMR

> O EMR foi construído via assinatura Claude Code + GPT + tempo de dev. Amortizado sobre 3 anos /
> 72 pacientes, fração EMR 60%, fatia Continuum 80%.

| Item | Cálculo | Valor/mês |
|---|---|---|
| Claude Code | US$ 200 × 5,40 | R$ 1.080 |
| GPT | US$ 20 × 5,40 | R$ 108 |
| Dev (tempo) | 60h × R$ 60 | R$ 3.600 |
| **Total/mês** | | **R$ 4.788** |

| Etapa | Valor |
|---|---|
| Gasto total de dev (× 8 meses) | R$ 38.304 |
| × 60% (fração EMR) | R$ 22.982 |
| × 80% (fatia Continuum) | R$ 18.386 |
| ÷ 72 pacientes (3 anos × 24/ano) | **~R$ 255/paciente/ciclo** |

### TOTAL BLOCO C

| Componente | Por paciente/ciclo |
|---|---|
| Fixos compartilhados (rateados) | R$ 85 |
| Variáveis (Daily, WhatsApp, Resend, storage) | R$ 70 |
| Licença (R$ 50/mês × 6) | R$ 300 |
| Amortização do EMR | R$ 255 |
| **TOTAL BLOCO C / ciclo** | **~R$ 710** |

---

## D — Operacional e administrativo

> Input do Getúlio, 24/05/2026. Regime: **Lucro Presumido**, cenário conservador (presunção 32%).

### Impostos (sobre a receita)

| Tributo | Base | Alíquota efetiva s/ receita |
|---|---|---|
| IRPJ | 15% sobre presunção de 32% | 4,80% |
| CSLL | 9% sobre presunção de 32% | 2,88% |
| PIS (cumulativo) | receita | 0,65% |
| COFINS (cumulativo) | receita | 3,00% |
| ISS (Londrina) | receita | 3,00% |
| **Carga tributária efetiva** | | **14,33%** |

**Imposto = 14,33% da receita** (escala com o preço final). Sobre o preço atual de R$ 25.000 =
**R$ 3.583/ciclo**.

Notas:
- (+) Adicional de IRPJ: 10% sobre o lucro presumido que exceder R$ 20.000/mês. Com 24 pacientes/ano
  e ticket atual o lucro presumido fica perto do teto — pode entrar uma fração. Sinalizado, não embutido.
- **Oportunidade — equiparação hospitalar** (presunção 8% IRPJ / 12% CSLL): derrubaria o efetivo
  pra ~8,9% (~R$ 2.225 sobre R$ 25k, economia ~R$ 1.350/paciente). Exige sociedade empresária +
  requisitos Anvisa; historicamente vale pra procedimentos/exames, não consulta pura. **Validar
  com o contador.**

### Gateway de pagamento — Asaas

Cartão parcelado = **2,99% + R$ 0,49** por cobrança (taxa fixa sobre o total, independente do nº de
parcelas). Escala com o preço. Sobre R$ 25.000 = **~R$ 748/ciclo**.

> Decisão de produto associada: **semestral em 6× / anual em 12×** (não mais 6m em 12×). Ajuste do
> deck em andamento.

### Contabilidade

R$ 800/mês × 12 = R$ 9.600/ano × 80% (fatia Continuum) ÷ 24 pacientes = **R$ 320/paciente/ciclo**.

### Seguro RC profissional (estimativa — sem apólice ainda)

Base de mercado R$ 1.000–5.000/ano por profissional. Estimativa equipe: médico ~R$ 3.000 + 3 demais
~R$ 1.200 = **~R$ 6.600/ano** × 80% ÷ 24 = **~R$ 220/paciente/ciclo**. Cotar apólice real.

### Pró-labore Getúlio (CEO)

Custo de sócio separado dos honorários clínicos do bloco A: **R$ 2.000/paciente/ciclo**.

### TOTAL BLOCO D (sobre preço de referência R$ 25.000)

| Componente | Por paciente/ciclo |
|---|---|
| Impostos (14,33%) | R$ 3.583 |
| Gateway Asaas (2,99% + R$ 0,49) | R$ 748 |
| Contabilidade | R$ 320 |
| Seguro RC (estimativa) | R$ 220 |
| Pró-labore CEO | R$ 2.000 |
| **TOTAL BLOCO D / ciclo** | **~R$ 6.871** |

---

## E — Aquisição de cliente (CAC)

> Input do Getúlio, 24/05/2026. Estratégia completa em `plano-aquisicao-marketing.md`.
> **Estágio MVP, sem pacientes/prova social** — o gasto de hoje é majoritariamente investimento de
> autoridade que beneficia toda a Plenya, não CAC transacional. Fração Continuum: 80%.

| Componente | Cálculo | Por paciente/ciclo |
|---|---|---|
| Conteúdo IG + tráfego (R$ 5.800/mês × 12) | × 80% ÷ 24 | R$ 2.320 |
| Mentoria Pedro Quintanilha (R$ 60.000) | × 80% ÷ 48 (2 anos) | R$ 1.000 |
| Sites Plenya + Getúlio (40% da base de dev R$ 38.304 = R$ 15.322) | × 80% ÷ 72 (3 anos) | R$ 170 |
| **TOTAL BLOCO E / ciclo** | | **R$ 3.490** |

Notas:
- **CAC-MVP é alto por baixo volume** (24 pacientes). Conforme a base cresce e o funil
  (consulta-porta-de-entrada + webinar evergreen + aplicação) amadurece, o CAC por paciente cai.
  Definir um **CAC-alvo de regime** (ex.: 10–15% do ticket) na Fase 2.
- VPS compartilhada já contabilizada no Bloco C (sem duplicação).
- Funil recomendado e faseamento: ver `plano-aquisicao-marketing.md`.

---

## F — Compliance e jurídico

> Input do Getúlio, 24/05/2026 (estimativas de mercado — nada contratado ainda). Fração Continuum 80%.

| Componente | Cálculo | Por paciente/ciclo |
|---|---|---|
| Setup jurídico inicial (contratos + termo de consentimento + parecer telemedicina/CFM + abatimento da consulta) — R$ 12.000 | × 80% ÷ 72 (amort. 3 anos) | R$ 133 |
| LGPD/jurídico recorrente **sob demanda** (~R$ 4.800/ano, pareceres pontuais) | × 80% ÷ 24 | R$ 160 |
| **TOTAL BLOCO F / ciclo** | | **~R$ 293** |

Notas:
- Setup baseado na Tabela OAB 2026 (~R$ 3.502/instrumento de análise de privacidade).
- Recorrente modelado **sob demanda** (sem mensalidade fixa) por estar em MVP. Níveis de mercado
  se escalar: mentoria/DPO-light R$ 1.200/mês · consultoria adequação R$ 1.900/mês · DPO dedicado
  R$ 3.000–6.000/mês.

---

## G — Educação continuada da equipe

> Decisão do Getúlio, 24/05/2026: **a cargo de cada profissional** (já embutida no que cobram).
> **Não é custo do Continuum.**

| **TOTAL BLOCO G / ciclo** | **R$ 0** |
|---|---|

---

## H — Custo de oportunidade (régua de decisão, NÃO soma ao custo)

> Decisão do Getúlio, 24/05/2026. **Soma R$ 0** ao custo — somar seria contar em dobro (o tempo
> dos profissionais já está pago no Bloco A). Serve como **piso de decisão da margem**.

**Benchmark:** consulta avulsa do Dr. Getúlio = **R$ 800 / 60min = R$ 800/h**. Detalhe relevante:
isso é *menor* que a hora-médico no Continuum (consulta R$ 900/h, call R$ 800/h) — o programa já
remunera a hora do Getúlio pelo menos tão bem quanto a avulsa.

**Como usar no final:** a margem do Continuum **por hora de Getúlio** (~9,25h clínicas/ciclo, ver
Bloco A) tem que superar o que essa hora renderia em consulta avulsa. Se não superar, o Continuum
não se justifica como uso do tempo dele, mesmo "dando lucro" no papel. Calcular na etapa de
definição do preço final.

| **TOTAL BLOCO H / ciclo (somado ao custo)** | **R$ 0** |
|---|---|

---

## I — Risco e contingência

> Decisão do Getúlio, 24/05/2026. Provisão **agregada** (não item a item), cobrindo inadimplência +
> gap de cancelamento precoce + substituição de profissional + caso complexo + escopo creep.

**Provisão = 5% sobre o custo de caixa (A–F + J)** (R$ 39.909 + R$ 4.000 = R$ 43.909) =
**~R$ 2.195/paciente/ciclo**. (G e H não entram: G = 0, H é não-caixa.)

Base de mercado: high-ticket é categoria de risco elevado (chargeback alvo <1%); reembolso concierge
é pró-rata com aviso de 30 dias (devolve só a parte não usada). Risco maior é cancelamento precoce
(onboarding front-loaded nas semanas 1-2).

Notas:
- **Reserva de reembolso** propriamente dita fica acoplada à **política de reembolso** (a definir).
- O modelo de **consulta-porta-de-entrada** (ver `plano-aquisicao-marketing.md`) já mitiga muito o
  risco: o cliente conhece o Getúlio e o programa antes de comprometer R$ 25–40k.

| **TOTAL BLOCO I / ciclo** | **~R$ 2.195** |
|---|---|

---

## J — Onboarding, fechamento e extras

> Input do Getúlio, 24/05/2026. A maior parte do onboarding/fechamento já está capturada em outros
> blocos; aqui ficam os **extras identificados** (painel genético + interpretação + coordenação).

| Componente | Por paciente/ciclo |
|---|---|
| Painel genético LIFECODE Nutri (Bioma) — custo Plenya | R$ 1.900 |
| Interpretação do painel pelo Dr. Getúlio (+1h) | R$ 900 |
| Coordenação/secretária-concierge (R$ 3.000/mês loaded × 12 × 80% ÷ 24) | R$ 1.200 |
| Avaliação inicial, 1º box, setup EMR, relatório final | R$ 0 (já em A/B/C) |
| **TOTAL BLOCO J / ciclo** | **R$ 4.000** |

Notas:
- LIFECODE: varejo R$ 2.290; usado R$ 1.900 (entre parceiro e varejo). **Confirmar tabela B2B com
  a Bioma** — preço-parceiro pode chegar a ~R$ 1.400.
- A interpretação (R$ 900) soma +1h às horas clínicas do Getúlio → ~10,25h/ciclo (relevante pro
  teste de margem/hora do Bloco H).

---

## CUSTO TOTAL CONSOLIDADO — Plano Semestral (6 meses)

| Bloco | Descrição | Por ciclo |
|---|---|---|
| A | Honorários dos 4 profissionais + WhatsApp | R$ 25.125 |
| B | Box Plenya (4 boxes) | R$ 3.420 |
| C | Infra + licença + amortização EMR | R$ 710 |
| D | Impostos (14,33%) + gateway + contábil + seguro + pró-labore CEO | R$ 6.871 |
| E | Aquisição (CAC-MVP) | R$ 3.490 |
| F | Compliance e jurídico | R$ 293 |
| G | Educação continuada (a cargo dos profissionais) | R$ 0 |
| H | Custo de oportunidade (régua, não-caixa) | R$ 0 |
| I | Risco e contingência (5%) | R$ 2.195 |
| J | Painel genético + interpretação + coordenação | R$ 4.000 |
| **TOTAL** | **Custo por ciclo semestral** | **~R$ 46.104** |

> ⚠️ **Custo (R$ 46.104) > preço atual do deck (R$ 25.000).** O semestral, como precificado hoje,
> dá prejuízo de ~R$ 21.000/paciente. Esta é a conclusão central da análise. Próximo passo:
> definir o preço final (ver seção de definição de preço, a construir).
>
> Observações importantes pra definição:
> - **Impostos e gateway escalam com o preço** (14,33% + 2,99%): ao subir o preço, ~17% do aumento
>   volta como imposto/taxa. Recalcular ao fixar o valor.
> - **CAC-MVP é inflado por baixo volume**; cai com escala (Fase 2 do plano de aquisição).
> - **Amortizações (EMR, sites, Pedro) são temporárias** — somem após o horizonte, baixando o custo.
> - **Teste de margem (H):** a margem por hora de Getúlio (~10,25h/ciclo) deve superar R$ 800/h da avulsa.

---

## Categorização estratégica pra precificação

| Tipo | Categorias | O que define |
|---|---|---|
| **Variável** | A, B, parte de C, J | preço-piso por paciente |
| **Fixo** | D, parte de C, E, F, G | break-even em volume |
| **Oportunidade** | H | preço-teto sustentável |
| **Risco** | I | margem de segurança |

### Onde os custos pesam mais (intuição inicial)

1. **Tempo profissional dos 4** — maior contribuinte ao custo variável
2. **Box e logística** — segundo maior custo material
3. **CAC** — precisa ser amortizado em 1 Continuum ou dividido entre os serviços Plenya?
4. **Onboarding upfront** — pesado nos primeiros 30-45 dias do ciclo, crítico pro cálculo de reembolso
5. **Fixos administrativos** — diluem com volume; 12 vagas/trimestre = 48/ano é o teto

---

## Inputs do Getúlio necessários antes de quantificar

- [ ] Quantas horas por paciente/mês CADA profissional dedica
- [ ] Modelo contratual dos profissionais (PJ, CLT, comissão por paciente)
- [ ] Custo médio de um Box (abertura + reposições)
- [ ] Custo da consulta avulsa Plenya (referência de custo-oportunidade)
- [ ] Margem de break-even desejada
- [ ] CAC estimado de 1 cliente Continuum
- [ ] Capacity utilization atual (as 12 vagas/trimestre se preenchem? quanto?)

## Próximo passo (quando retomar)

Quantificar uma categoria de cada vez. Duas ordens possíveis:
- **Fixos primeiro** → achar o break-even em volume
- **Variáveis primeiro** → achar a margem por paciente

Tema acoplado: **política de reembolso** (slide 16/Compromisso do deck) depende desta análise.
