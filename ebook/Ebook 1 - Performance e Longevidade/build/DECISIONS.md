# Decisões de build do EPUB

Decisões tomadas em conversa com o autor sobre como o EPUB deve ser gerado.
Aplicar no próximo rebuild.

---

## Pendentes (ainda não aplicadas)

### 1. Remover sumário inline visível
**Status:** aprovado, aguarda rebuild.
**Onde:** `build-epub.py` — remover flag `--toc` do comando Pandoc (manter `nav.xhtml` automático para navegação nativa do Kindle).
**Racional:** padrão dos best-sellers atuais do gênero (Outlive, Lifespan, Being Mortal). Abre o livro direto: créditos → dedicatória → epígrafe → introdução. Leitor encontra estrutura via menu nativo do Kindle.
**Data:** 2026-04-20.

### 2. Converter PNG → JPG (qualidade 85)
**Status:** aprovado, aguarda rebuild.
**Onde:** `build-epub.py` — na função `prepare_images()`, converter cada PNG para JPG q=85 ao copiar.
**Racional:** reduz EPUB de ~20 MB para ~5-7 MB. Menor tempo de download Kindle, melhor experiência em conexões lentas.
**Data:** 2026-04-20.

### 3. Remover headers de Parte duplicados
**Status:** aprovado, aguarda edição dos capítulos.
**Onde:** chapters 02, 03 (Parte I), 05, 06 (Parte II), 07, 08, 09, 10 (Parte III), 12 (Parte IV), 14 (Parte V) — remover `# PARTE X — ...`.
**Manter header de Parte apenas em:** 01, 04, 06b, 11, 13 (primeiro capítulo de cada parte).
**Racional:** leitura contínua não precisa repetir o nome da parte 3 vezes. Estrutura fica no sumário nativo do Kindle.
**Data:** 2026-04-20.

### 4. Manter drop caps e hifenização
**Status:** confirmado, sem mudança necessária.
**Onde:** CSS atual já tem `h1 + p::first-letter` e `hyphens: auto`.
**Racional:** autor gosta da estética. Dispositivos Kindle modernos (2018+) renderizam bem.
**Data:** 2026-04-20.

### 6. Contatos ampliados no "Sobre o Autor"
**Status:** aplicado na fonte (md/pt-BR/sobre-o-autor.md), aguarda rebuild.
**Decisão:** adicionar hierarquia de 2 níveis no bloco final da bio.
- **Principal (pessoal):** drgetulioamaralfilho.com.br + @drGetulioAmaralFilho
- **Secundário (onde atende):** Plenya Saúde — Londrina-PR · plenyasaude.com.br · @plenyaSaude
- Sem verbo ("Atende medicina preventiva na..."), estrutura de afiliação institucional.
**Racional:** conteúdo lidera (bio) → destaque pessoal (autor vai a palestras, escreve, fala) → afiliação profissional (onde atende). Precedente: Outlive/Attia cita peterattiamd.com e Early Medical; Lifespan/Sinclair cita Sinclair Lab. Afiliação clínica é autoridade aplicada, não brochura.
**Data:** 2026-04-20.

### 5. Foto do autor — manter no FIM
**Status:** confirmado, sem mudança necessária.
**Onde:** `sobre-o-autor.md` — foto já está como primeira linha da página final.
**Racional:** padrão trade nonfiction (Outlive, Lifespan, etc.) e medicina brasileira. Conteúdo lidera, credencial fecha. Marketing da foto se resolve no Author Central da Amazon (perfil separado).
**Data:** 2026-04-20.

---

## Aplicadas

### ✅ Nomenclatura AGIR alinhada ao site Plenya (2026-04-20)
Definição canônica da marca (site `plenya-continuum.md` §5) aplicada em todo o livro:
- **A** — Atividade Física, Alimentação e Suplementação **Inteligente** (ordem invertida + "Inteligente")
- **G** — Gestão **Clínica e** Metabólica
- **I** — Integração **Mente-Corpo** (ordem invertida)
- **R** — Ritmo Circadiano e Repouso (sem mudança)

Atingiu: `00-indice.md`, `00b-introducao.md`, `06-saude-metabolica.md`, `06b-parte-iii-intro.md`, capítulos 7-11, 14, `briefing-branding.md`, `kdp-description.md`, `personagens.md`, `glossario.yaml`, `build-epub.py`.

Arquivos renomeados (git mv): `07-alimentacao-atividade-suplementacao.md` → `07-atividade-alimentacao-suplementacao.md`; `09-integracao-corpo-mente.md` → `09-integracao-mente-corpo.md`. ACTS (EN) ajustado em paralelo: A — Activity, Alimentation & Smart Adjuncts; T — Tending Mind & Body.

### ✅ Marcadores `<!-- EPUB-START -->` e `<!-- EPUB-END -->`
Delimitam o conteúdo publicável dentro de cada arquivo `.md`. Tudo fora dos marcadores vira nota editorial interna — fica no arquivo para referência mas não sai no livro.

### ✅ `{.unnumbered .unlisted}` nos headings de frontmatter
Aplicado a: Dedicatória, Epígrafe, Créditos. Páginas existem no livro mas não aparecem no sumário nem recebem numeração.

### ✅ Instruções editoriais removidas
Frontmatter, créditos, agradecimentos e bio do autor — todos limpos de frases tipo "Posicionamento: página iv...", "Texto pronto para diagramação...", "Diagramação sugerida...", "Versão final".

### ✅ Foto do autor embutida na página "Sobre o Autor"
Usa `getulio_bw_halfbody_1000.jpg` (P&B, meio corpo, 1000×1339). Copiada pelo build script para `images/autor.jpg` dentro do EPUB.

### ✅ Estrutura multilíngue
`md/{pt-BR,en,es,fr,de}/`, `figuras/{...}/`, `capas/{...}/`. Build script recebe idioma como argumento: `python3 build-epub.py pt-BR`.

---

## Notas de possíveis ajustes futuros (não confirmados)

- [ ] **Part headers repetidos no sumário** — `# PARTE I — O DESPERTAR` aparece 3× (uma vez por capítulo de Part I), mesmo com `--toc` removido ainda fica no texto. Solução seria manter o header só no primeiro capítulo de cada parte.
- [ ] **Otimização de imagens** — 28 PNGs totalizam ~18 MB no EPUB. Conversão para JPG qualidade 85 reduziria para ~4 MB, mas PNGs de figuras técnicas com texto às vezes ficam borradas em JPG. Avaliar caso a caso.
- [ ] **Drop caps** — CSS atual tem `h1 + p::first-letter` com drop cap marrom. Suporte Kindle inconsistente — pode ser melhor desabilitar para consistência entre dispositivos.
- [ ] **Hyphenation** — atualmente `hyphens: auto` no CSS. Alguns dispositivos Kindle não suportam. Manter ou forçar `hyphens: manual`?
- [ ] **Foto do autor — posicionamento** — está como primeira linha da página "Sobre o Autor". Alternativas: como inset lateral, como rodapé, ou em página separada antes da bio.
