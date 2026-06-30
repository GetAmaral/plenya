# Gymnai — Planejamento da Spinoff

Gymnai é uma **spinoff do módulo de treinamento físico do EMR Plenya**, virando produto
próprio: um PWA scan-first onde o usuário lê um QR code no aparelho de academia e vê o vídeo
explicativo daquele aparelho — e, se tiver conta ativa, o treino dele para aquele aparelho no
dia. Modelo B2B2C com parcerias de **academias, condomínios e personal trainers**.

> **Status:** planejamento (pré-código). Decisões estratégicas fechadas; código ainda não
> iniciado. O projeto vai nascer em **`/home/user/gymnai`** (repo próprio, separado do monorepo
> Plenya). Estes documentos vivem em `plenya/docs/gymnai/` porque é onde a decisão foi tomada e
> porque parte das decisões toca o lado EMR (fronteira prontuário-vs-catálogo).

## Decisões já fechadas (resumo)

| Tema | Decisão | Onde detalha |
|------|---------|--------------|
| Independência | **Totalmente separado** da Plenya. Compartilha só o catálogo de treinos/vídeos como fonte de verdade. | [00](00-estrategia-spinoff.md) |
| Localização do código | **Repo próprio em `/home/user/gymnai`** (não subprojeto do monorepo Plenya). Contexto de Claude próprio, semeado. | [00](00-estrategia-spinoff.md) |
| Base do código | **Greenfield em TypeScript** (não copy-first Go). Compartilha com a Plenya só o **dado** do catálogo, não código. | [00](00-estrategia-spinoff.md) · [01](01-stack-e-infra.md) |
| Fronteira de dados | **Prontuário fica no EMR.** Gymnai é dono do catálogo + motor; dado ligado a paciente nunca sai do EMR. | [00](00-estrategia-spinoff.md) |
| Stack | TS full-stack: Next.js PWA + Postgres/Drizzle + RLS + Better Auth. | [01](01-stack-e-infra.md) |
| Infra MVP | **VPS dedicada própria** (não a VPS da Plenya), Cloudflare grátis na frente. Vídeo na VPS no MVP, com gatilho de migração. | [01](01-stack-e-infra.md) |
| Pagamento | **Asaas** (Pix Automático + boleto + split nativo p/ parcerias). | [01](01-stack-e-infra.md) |
| Vídeo | Biblioteca **por arquétipo, IA/biomecânica, já existente na VPS treinador** (importar). VPS no MVP; migração = **Bunny Stream**. | [01](01-stack-e-infra.md) · [04](04-questoes-abertas.md) |
| Modelo de negócio | QR+vídeo = **isca grátis**; assinatura paga (**Básico/Plus**); pagamento **pessoal ou patrocinado** (roster). Escanear **nunca** concede acesso. | [00 §8](00-estrategia-spinoff.md) |
| LGPD | Avaliação completa = **dado sensível** (art. 11); **18+**; **educador físico CREF** = responsável técnico (revisão opcional). | [00 §9](00-estrategia-spinoff.md) |
| Marca | Identidade **fechada**: ouro/petróleo, Trajan→**Cinzel** + Montserrat, "Intelligence in Motion". Vetores + ícones gerados. | [03](03-marca.md) · [`identidade/`](identidade/) |
| Wedge | **Academia** primeiro (QR por aparelho; patrocínio por **roster**, não por scan). | [02](02-dominio-e-mvp.md) |

**Pendências externas (não bloqueiam o scaffold de código):**
- 🔑 **Acesso à VPS treinador** (`72.62.108.11`) — inventariar os vídeos (formato/qtd) e ler o
  `gerador_treinos.py` (motor de treino). Bloqueia o *conteúdo* do catálogo, não a estrutura.
- 💰 **Preço** de Básico/Plus + estrutura do patrocínio — decisão do usuário; só vira código na fase 2.
- ❄️ **Congelar o módulo de treino do EMR** (só bug-fix) — ação do lado Plenya ao iniciar o Gymnai.

## Índice

- **[00 — Estratégia da spinoff](00-estrategia-spinoff.md)** — o que é, separação, localização,
  greenfield-vs-copy-first, fronteira prontuário-vs-catálogo, congelamento do treino no EMR.
- **[01 — Stack e infraestrutura](01-stack-e-infra.md)** — stack greenfield TS, as divergências
  vs Plenya, infra MVP (VPS dedicada), estratégia de vídeo na VPS com plano de migração.
- **[02 — Domínio e MVP](02-dominio-e-mvp.md)** — modelo de domínio (Organization/Equipment/QR/
  Entitlement), multi-tenancy com RLS, página de scan contextual (vídeo público + treino autenticado), escopo do MVP e roadmap,
  decisões em aberto.
- **[03 — Marca](03-marca.md)** — identidade visual fechada (ouro/petróleo, Trajan + Montserrat,
  "Intelligence in Motion"). Assets em [`identidade/`](identidade/).
- **[04 — Questões abertas](04-questoes-abertas.md)** — tracker da revisão detalhada do plano
  (itens críticos/altos/médios), fechados um a um. **Ler antes de iniciar o código.**
- **[CLAUDE-seed.md](CLAUDE-seed.md)** — rascunho do `CLAUDE.md` a copiar para `/home/user/gymnai`
  quando o repo for criado, + o que semear de memória.

## Como usar

Ao iniciar o projeto: criar `/home/user/gymnai`, `git init`, copiar `CLAUDE-seed.md` para
`/home/user/gymnai/CLAUDE.md`, semear a memória conforme indicado, e seguir o roadmap do [02](02-dominio-e-mvp.md).
