# 00 — Estratégia da Spinoff

> Documento de decisão. Registra **o quê** e **por quê**, não o como (ver [01](01-stack-e-infra.md)
> e [02](02-dominio-e-mvp.md)). Data: 2026-06-29.

## 1. O que é o Gymnai

Destacamento do módulo de treinamento físico do EMR Plenya, virando produto próprio:

- **PWA scan-first**: o usuário lê um QR code afixado em cada aparelho de academia e acessa o
  **vídeo explicativo** daquele aparelho. Funciona para qualquer um, sem login (isca de aquisição).
- **Treino do dia por aparelho**: se o usuário está logado com assinatura ativa, além do vídeo
  vê **o treino dele para aquele aparelho hoje**.
- **B2B2C**: **academias e condomínios** são organizações/tenants (lugares com aparelhos);
  **personal trainers** entram como o ator `Professional` (pessoa, não tenant). O aluno é o
  usuário final.
- **Assinatura paga, pagamento por duas fontes**: **pessoal** (o aluno paga) ou **patrocinado**
  (academia/condomínio **ou** professor paga pelo aluno, por conciliação de roster). **Escanear o
  QR nunca concede assinatura** — é isca pública (ver §8).

A tag importa em dois níveis: **por aparelho** (qual vídeo/arquétipo) e **por organização** (qual
parceria/tenant, para co-branding e atribuição — não para liberar acesso). Ver desenho de domínio
em [02](02-dominio-e-mvp.md).

## 2. Decisão: totalmente separado da Plenya

O Gymnai é um **produto separado**, não um módulo da Plenya. Repo próprio, infra própria, marca
própria, banco próprio, auth própria. O único elo é a **base de treinos/vídeos como fonte de
verdade compartilhada** — e mesmo essa é compartilhada como **dado**, não como código (ver §4).

Consequências:
- Identidades separadas: um aluno do Gymnai **não é** `Patient` do EMR. Mesmo que uma pessoa seja
  cliente das duas, são cadastros distintos, ligados só por consentimento explícito no futuro.
- Some o risco LGPD de misturar dado de academia com prontuário — porque nada é compartilhado por
  padrão.
- A separação preserva a opcionalidade de funding/venda do Gymnai e isola acesso (time/contractor
  do Gymnai não toca no código clínico do EMR).

## 3. Decisão: repo próprio em `/home/user/gymnai` (não subprojeto do monorepo)

Para efeito de Claude Code, o Gymnai é um **projeto separado em pasta própria**, não um
`apps/gymnai` dentro do monorepo Plenya.

Por quê (mecânica do Claude Code):
- **Memória é indexada por diretório.** Pasta separada (`/home/user/gymnai`) ganha um banco de
  memória próprio e limpo; dentro do monorepo, herdaria o banco inteiro da Plenya (~90%
  irrelevante para um SaaS de academia → ruído de recall).
- **CLAUDE.md é cascata.** No monorepo, o Gymnai herdaria as Regras de Ouro da Plenya (deploy
  por-app, single-tenancy, voz editorial, regras clínicas) — boa parte **erradas** para o Gymnai,
  contaminando o contexto.
- **Isolamento de acesso e venda**: histórico git próprio; time do Gymnai não recebe o EMR clínico.

Como aproveitar a base de conhecimento mesmo separado: **seed único e curado** (não herança).
Ver [CLAUDE-seed.md](CLAUDE-seed.md).

> Conveniência: manter `/home/user/plenya` e `/home/user/gymnai` lado a lado é ok; o que importa é
> serem **raízes de projeto distintas**. Não colocar o Gymnai *dentro* de `/home/user/plenya`.

## 4. Decisão: greenfield TypeScript (não copy-first Go)

O Gymnai é **construído do zero em TypeScript**, não uma cópia do motor de treino Go do EMR.

Contexto da decisão: chegamos a considerar copy-first (copiar o módulo Go e divergir, "menos
trabalho"). Mas, analisando o produto do zero, a parte difícil do Gymnai não é computação — é
**multi-tenancy, auth, pagamento e funil consumer**, e essa tríade vive melhor num stack TS único
da borda ao banco. Sem o motor Go para herdar, Go vira escolha por gosto, não por necessidade.

O que **não** muda com essa decisão:
- O elo com a Plenya continua sendo **o catálogo como fonte de verdade** — só que compartilhado
  como **dado** (exercícios/vídeos exportados/semeados no banco do Gymnai), não como código Go.
- A **fronteira prontuário-vs-catálogo** (§5) vale igual.

Trade-off aceito: mais trabalho inicial (reescrever o domínio de treino em TS) em troca de melhor
encaixe de longo prazo e velocidade de time pequeno. Racional de stack detalhado em [01](01-stack-e-infra.md).

> Nota de coerência: Go foi e continua sendo a escolha certa **para o EMR** (sistema clínico
> regulado, compute-heavy, integridade > velocidade). O Gymnai é um problema diferente
> (consumer, integração, funil) → outra resposta certa. Mesmo critério, produtos diferentes.

## 5. Fronteira: prontuário fica no EMR, catálogo/motor vai para o Gymnai

A divisão é **pela natureza do dado, não pelo módulo**:

| Natureza | Exemplos | Onde mora |
|----------|----------|-----------|
| Conteúdo / motor genérico | catálogo `Exercise`, GIFs, **vídeos**, mapeamento aparelho→exercício, lógica de montar/recomendar plano, periodização, IA de treino | **Gymnai** (fonte de verdade) |
| Dado clínico ligado ao paciente | `PhysicalAssessment`, `PosturalAssessment`, `FitnessTestResult`, o `WorkoutPlan` **prescrito** a um paciente, execução/adesão (`WorkoutSession`/logs) | **EMR** (prontuário) |

Quem prescreve no EMR é o **educador físico** (`RolePhysicalEducator`, papel clínico dentro de
`RequireClinician()`) — a prescrição continua sendo ato profissional ligado a um paciente, logo
prontuário. Nunca sai do EMR.

Integração futura (opcional, não garantida): quando/se o EMR consumir o catálogo do Gymnai, ele
guarda o resultado localmente como prontuário com **snapshot denormalizado** dos itens do catálogo
(mesmo padrão `PlanSnapshot` já usado em subscriptions), para a ficha sobreviver a mudanças no
catálogo.

## 6. Lado EMR: congelar o módulo de treino + dono do catálogo

Mesmo com greenfield (sem copiar código), duas medidas evitam retrabalho e divergência de catálogo:

1. **Dono futuro do catálogo = Gymnai.** É lá que o time grande vai mexer em conteúdo e vídeo; o
   catálogo do Gymnai vira o mais rico. O catálogo do EMR é snapshot a ser substituído na
   integração futura.
2. **Congelar o módulo de treino do EMR para correções apenas.** Durante a janela de
   desenvolvimento do Gymnai, **nenhum conteúdo/feature novo de treino entra no EMR** — só bug
   fix. Todo investimento em exercício/vídeo/IA vai para o Gymnai. Assim não se constrói a mesma
   coisa duas vezes.

> "Integrar de volta" é tratado como **bônus, não garantia.** Se nunca acontecer, ninguém fica
> pior: o EMR segue com seu módulo de treino congelado e o Gymnai segue sozinho. Se acontecer, é
> plugar o catálogo, não desmontar nada.

## 7. Riscos a vigiar

- **Divergência do catálogo** entre EMR e Gymnai → mitigado por "dono = Gymnai" + freeze do EMR.
- **Multi-tenancy retrofitada** → o Gymnai nasce multi-tenant (`organization_id` + RLS) desde a
  migration 00001 (ver [02](02-dominio-e-mvp.md)). Nunca adiar.
- **Confusão aluno-vs-paciente (LGPD)** → bancos e identidades separados; integração só com
  consentimento.
- **Co-hospedar na VPS da Plenya** → **não fazer** (ver [01](01-stack-e-infra.md) §infra).
- **MVP depende do motor do treinador** → o produto pago (Básico) é avaliação+treino+periodização;
  esse motor está na VPS treinador, parcialmente no EMR. Risco de cronograma se o motor não portar
  fácil para TS. Inventariar via acesso à VPS (ver [04](04-questoes-abertas.md) I1/I2).

## 8. Modelo de negócio (funil e planos)

O QR e os vídeos são **isca**, não o produto. O produto que monetiza é a assinatura.

```
Grátis (isca)   → QR no aparelho → vídeo educativo (uso correto, variações, erros) + ad/CTA
                  Sem custo para usuário nem academia. Pública, sem login.
Plano Básico    → avaliação física + treino montado + periodização (motor do treinador)
   (pago)         Conversão a partir do CTA. B2C via Asaas/Pix.
Plano Plus      → treinos/periodizações avançados + contato com time de especialistas.
   (pago)
```

**Patrocínio (academia/condomínio OU professor/personal):** um **patrocinador paga o Básico** para
um grupo de usuários. Concessão por **conciliação de roster** — cruzamos nossa base com a lista do
patrocinador (**planilha** no MVP, **API** depois). **Escanear o QR não concede assinatura** — só
pagar (pessoal) ou estar no roster pago (patrocinado). "Facility" (org paga) é um caso de patrocínio;
o **professor patrocinando seus alunos** é outro — mesmo mecanismo (`Sponsorship`).

**Os atores e por que entram:**
- **Aluno (consumer):** vê de graça como usar o aparelho; assina se quiser treino/avaliação.
- **Academia/condomínio:** aceita o QR grátis sem fricção; pode patrocinar o Básico dos alunos como
  perk de retenção.
- **Professor / personal trainer:** atua em várias academias e tem vários alunos; pode **patrocinar**
  a assinatura dos próprios alunos e **montar o treino manualmente ou revisar/ajustar o gerado pela
  IA** (autoria do `WorkoutAssignment` = IA + professor, human-in-the-loop).

## 9. LGPD e dado sensível (I5)

**Decisão (2026-06-29): avaliação física COMPLETA (ACSM) + idade mínima 18+.**

O Plano Básico coleta avaliação completa (antropometria, **pressão arterial, exames laboratoriais,
condições/risco cardiovascular**). Isso é **dado sensível de saúde** (LGPD art. 11) → requisitos:

- **Consentimento específico e destacado** para dado sensível (separado do consentimento geral).
- **Base legal:** execução de contrato (a assinatura) + consentimento para o sensível.
- **Segurança reforçada:** criptografia at-rest dos campos clínicos + controle de acesso + **audit log**.
- **Retenção** explícita (definir prazos por categoria).
- **18+ no cadastro** (sem fluxo de menor no MVP; menores ficam para fase posterior, art. 14).
- **Responsável por privacidade** designado (não DPO formal no MVP).
- **Fronteira:** a avaliação do Gymnai é **dado próprio**, separada do prontuário EMR. Nenhum dado
  clínico do EMR cruza para o Gymnai.
- **Responsabilidade profissional (DECIDIDO):** avaliação e treino são **100% gerados por IA**. Um
  **educador físico do Gymnai (CREF) é o responsável técnico** que **assina/responde** pela avaliação;
  **revisão é opcional**, não obrigatória por caso (sem gargalo humano). Cobertura = responsabilidade
  técnica + **disclaimer "screening/orientação de atividade física, não diagnóstico médico"**.
  Manter o enquadramento como **screening/orientação** (escopo do educador físico) — se virar
  "interpretação de exames/diagnóstico", sobe para ato médico. (Esse RT é papel interno do Gymnai,
  distinto do ator `Professional`, que são os personais externos com seus alunos.)

**Em aberto (decisão de preço — I2):** valores de Básico e Plus, preço/estrutura da parceria
facility (por usuário conciliado? pacote de seats?), e se há trial.
