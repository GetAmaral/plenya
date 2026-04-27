# Pós-Onda 3 — monitoramento e iteração

O SEO é um sinal lento. As Ondas 1-2-3 plantam — a colheita começa em 30-90 dias e madura em 6-12 meses.

Este doc define o que olhar **mensalmente** pra saber se a estratégia está funcionando.

---

## Painel mensal (15 min/mês)

### 1. Plausible (já no ar nos 2 sites)

| Métrica | Onde | O que olhar |
|---|---|---|
| Visitantes únicos | Dashboard | Crescendo MoM? |
| Páginas mais vistas | Top Pages | Pages-serviço estão entrando no top 10? |
| Sources | Sources > Search | Google é a maior fonte? Crescendo? |
| Time on page | Engagement | Páginas-serviço >2 min = bom |
| Bounce rate | Engagement | Blog posts <70% = bom |

### 2. Google Search Console (após cadastro — Onda 3 doc 05)

| Relatório | Frequência | Sinal |
|---|---|---|
| **Performance** | Mensal | Cliques + impressões crescentes |
| **Top queries** | Mensal | Aparecendo pra "healthspan", "nefrologia preventiva", "checkup longevidade Londrina"? |
| **Top pages** | Mensal | Pages-serviço entrando no ranking? |
| **Coverage / Indexing** | Mensal | Erros = 0; "Submitted, indexed" = todos os URLs do sitemap |
| **Core Web Vitals** | Trimestral | LCP < 2.5s, FID < 100ms, CLS < 0.1 |
| **Manual actions** | Mensal | Tem que estar VAZIO. Se aparecer algo, é urgente |

### 3. Google Business Profile

| Métrica | Onde | O que olhar |
|---|---|---|
| **Buscas por nome** | GBP Insights | Crescendo? Significa awareness |
| **Buscas categóricas** | GBP Insights | "clínica longevidade Londrina" — Plenya aparece? |
| **Ações** (chamadas, direções, site) | GBP Insights | Conversão real |
| **Reviews** | GBP | +1 a +3/mês orgânico = saudável |
| **Mensagens** | GBP | Resposta em <24h sempre |

---

## Alertas vermelhos (agir imediato)

🚨 **Queda de tráfego >30% num mês** sem mudança no site → checar GSC > Manual actions, depois Coverage por erros novos.

🚨 **Posição caindo numa query importante** → checar concorrência (alguém publicou conteúdo melhor) e revisar copy/profundidade do nosso.

🚨 **Plausible sem dados há 7 dias** → script de tracking quebrou (raro, mas verificar).

🚨 **Review negativa** → responder em <24h, profissional, sem revelar dados clínicos do paciente (sigilo CFM). Resposta padrão:

> Agradeço o feedback. Vou entrar em contato em particular pra entender melhor sua experiência. Att, Equipe Plenya.

---

## Sinais verdes (continuar fazendo igual)

✅ Posição melhorando numa query principal → publicar 1 artigo do blog **reforçando** o tema.

✅ Backlink novo descoberto (GSC > Links) → linkar de volta se relevante (não obrigatório).

✅ Página entrando no top 10 → otimizar title+description pra **subir** posição (CTR-based reranking).

---

## Cadência de novos artigos (manter motor rodando)

A Plenya tem 28 posts hoje. Para sustentar autoridade tópica:

- **Mínimo**: 2 artigos/mês, ~1500 palavras, com 5+ refs científicas
- **Bom**: 4 artigos/mês
- **Ótimo**: 1 artigo/semana (= 52/ano = top 10% dos blogs médicos brasileiros)

Pipeline existente em `scripts/blog-generator/` já facilita. Quando bater 50 posts, vamos pra próxima onda (cluster topic + linkagem interna automatizada).

---

## Quando reavaliar a estratégia inteira

A cada **6 meses**:
- Rodar nova auditoria competitiva (queries top, quem ranka)
- Atualizar lista de podcasts (alguns morrem, novos surgem)
- Revisar lista de páginas-serviço — alguma está performando pouco e vale repensar?
- Ler GSC com o ChatGPT/Claude pra extrair queries não óbvias

A cada **12 meses**:
- Rodar Lighthouse completo nos 2 sites
- Revisar JSON-LD (Google muda specs)
- Confirmar GBP ainda ativo, fotos atualizadas, posts publicando

---

## Resumo: 3 perguntas que importam

1. Mais gente está entrando no site mês contra mês?
2. As pessoas certas estão entrando (queries de intenção comercial / próximas da consulta)?
3. Estão fazendo o que a gente quer (clicar em /contato, baixar livro, agendar)?

Se as 3 estiverem em verde, a engrenagem está girando. Se uma travar, mergulhe nela primeiro.
