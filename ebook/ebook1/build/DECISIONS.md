# Decisões de build do EPUB

Decisões tomadas em conversa com o autor sobre como o EPUB deve ser gerado.
Aplicar no próximo rebuild.

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
