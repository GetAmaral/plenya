# Auditoria URL × Query — site Plenya (maio/2026)

Verificação direta no código (sem suposição). Para cada query-alvo listada em `07-pos-onda3-monitoramento.md`, mapeia a URL canônica candidata e os ajustes específicos.

Convenção: ✅ = metaTitle on-point · ⚠ = parcial · ❌ = gap.

---

## Queries-alvo e estado atual

| Query | URL canônica | metaTitle atual | Estado |
|---|---|---|---|
| `healthspan` | `/healthspan` | "Healthspan — viver bem por mais tempo, com evidência" | ✅ |
| `medicina funcional integrativa Londrina` | `/medicina-funcional-integrativa` | "Medicina funcional integrativa — abordagem clínica em Londrina" | ✅ |
| `escore plenya` (branded) | `/escore-plenya` | "Escore Plenya — instrumento de medida do Método AGIR" | ✅ |
| `método agir` (branded) | `/metodo-agir` | "Método AGIR" | ⚠ minimalista — sem cauda longa, perde CTR |
| `nefrologia preventiva` | `/avaliacao-renal-preventiva` | "Avaliação renal preventiva — nefrologia antes do sintoma" | ⚠ falta a frase exata "nefrologia preventiva" no título |
| `checkup longevidade Londrina` | `/checkup-longevidade` | "Checkup de longevidade — avaliação clínica baseada em evidência" | ⚠ falta geo "Londrina" |
| `clínica longevidade Londrina` | `/` (home) | "Plenya — Saúde, Performance & Longevidade" (template) | ❌ sem "clínica" + "Londrina" |
| `clínica medicina funcional Londrina` | `/medicina-funcional-integrativa` ou `/a-plenya` | duplo candidato → escolha + canonical | ⚠ disputa interna entre duas páginas |

---

## Ajustes propostos (por URL)

### `/checkup-longevidade`

**De:** "Checkup de longevidade — avaliação clínica baseada em evidência"
**Para:** "Check-up de Longevidade em Londrina — avaliação clínica com VO₂ máx, ApoB e composição corporal | Plenya"

Razão: adiciona geo + 3 exames-âncora (cauda longa). Mantém credibilidade.

### `/avaliacao-renal-preventiva`

**De:** "Avaliação renal preventiva — nefrologia antes do sintoma"
**Para:** "Nefrologia Preventiva em Londrina — avaliação renal antes do sintoma | Dr. Getúlio Amaral Filho · Plenya"

Razão: frase-chave "nefrologia preventiva" canônica no início + médico signatário + geo.

### `/` (home)

**De:** template "Plenya — Saúde, Performance & Longevidade"
**Para:** "Plenya — Clínica de Longevidade em Londrina · Medicina Funcional Integrativa"

Razão: capta "clínica longevidade Londrina" e "medicina funcional integrativa Londrina" ao mesmo tempo. Ajustar `metadata.title.default` em `apps/site/app/[locale]/layout.tsx`.

### `/metodo-agir`

**De:** "Método AGIR"
**Para:** "Método AGIR — os quatro pilares da medicina de longevidade Plenya | Atividade, Gestão, Integração, Ritmo"

Razão: brand-term + descritor + os 4 pilares ajudam descoberta de cauda longa.

### `/a-plenya`

Status: hoje conflita com `/medicina-funcional-integrativa` para a query "clínica MFI Londrina".

**Decisão:** manter `/medicina-funcional-integrativa` como página canônica de serviço (mais transacional). Em `/a-plenya`, ajustar metaTitle para foco institucional: "Plenya — Manifesto e Método de uma Clínica de Longevidade".

---

## Gaps de conteúdo (não-meta)

1. **/healthspan** — não menciona o livro ANTES. Como o livro é a obra-fundadora do conceito de "janela silenciosa", deve ter um bloco de aprofundamento linkando para `drgetulioamaralfilho.com.br/livros/antes` com `BookSchema` (criado em `apps/site/components/seo/book-schema.tsx`).
2. **/escore-plenya** — não tem FAQ. FAQPage schema já existe em `/a-plenya` e `/onde-atendo` (task #222 completed) mas escore é a "porta de entrada" e merece 5-8 perguntas próprias ("o escore é gratuito?", "quanto tempo leva?", "vale para quem já tem doença?", etc.).
3. **/checkup-longevidade** — bom title, mas conteúdo deve listar **explicitamente** os 12 exames do post `12-exames-que-um-checkup-de-longevidade-pede.mdx`. Hoje a página fala em conceito; a query "checkup longevidade Londrina" é transacional — o usuário quer ver a lista.
4. **/avaliacao-renal-preventiva** — falta seção "para quem é" (diabetes, hipertensão, histórico familiar, uso de AINEs, idade >40). Marca intenção clínica.

---

## Link-interno / cluster

Já implementado nesta sessão em `apps/site/app/[locale]/blog/categoria/[pilar]/page.tsx`:

- **CollectionPage + ItemList** schema em cada hub de pilar
- Bloco "Aprofunde no método" linkando cada hub para 2-3 páginas institucionais correspondentes

Próximo (pendente): reverse-link — adicionar bloco "Artigos do pilar" em `/metodo-agir`, `/healthspan`, `/escore-plenya` puxando 3 posts mais recentes do pilar correspondente.

---

## Prioridade de execução

1. Ajustar metaTitle home `/` (maior impacto — capta a query mais transacional)
2. Ajustar metaTitle `/checkup-longevidade` e `/avaliacao-renal-preventiva` (geo + frase-chave)
3. Reverse-link: posts do pilar nas páginas institucionais
4. FAQ em `/escore-plenya`
5. Lista explícita dos 12 exames em `/checkup-longevidade`

Itens 1-2 = 30 min. Itens 3-5 = 2-3h.

---

**Fonte primária verificada:**
- metaTitles: `apps/site/messages/pt.json` (e equivalente EN em `messages/en.json`)
- queries: `docs/seo/07-pos-onda3-monitoramento.md`
- pillar hubs: `apps/site/app/[locale]/blog/categoria/[pilar]/page.tsx` (já enriquecido)
