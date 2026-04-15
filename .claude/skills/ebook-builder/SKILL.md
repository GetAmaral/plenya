# eBook Builder — Instruções do Skill `/ebook`

> Referência interna usada por `.claude/commands/ebook.md`.
> Gera eBooks para público **leigo** (ajustável), com referências **reais** (sem alucinar), capítulo-a-capítulo com aprovação humana em cada etapa. Toda saída em **Português**.

---

## Modos de operação

Analise `$ARGUMENTS`:

- vazio, "novo", "criar", ou tema livre → **MODO BRIEFING**
- "capítulo N" / "próximo capítulo" / "escrever cap N" → **MODO CAPITULO**
- "revisar capítulo N" / "editar cap N" → **MODO REVISAR**
- "exportar" / "compilar" / "gerar pdf" → **MODO EXPORTAR**

Ao iniciar qualquer modo que não seja BRIEFING, localize o livro ativo:

```bash
ls ~/ebook-output/*/00-briefing.md 2>/dev/null
```

Se houver vários, pergunte qual. O `output_dir` é o diretório do briefing.

---

## MODO BRIEFING — passo 0 a 2

### Passo 0 — Coletar parâmetros via AskUserQuestion

**OBRIGATÓRIO:** use `AskUserQuestion` com uma única chamada contendo as perguntas abaixo. Não pergunte em texto livre.

```
AskUserQuestion(questions=[
  {
    "question": "Qual o público-alvo?",
    "header": "Público",
    "multiSelect": false,
    "options": [
      {"label": "Leigo geral", "description": "Sem conhecimento prévio; linguagem simples, analogias, zero jargão"},
      {"label": "Leigo interessado", "description": "Já leu sobre o tema; tolera termos técnicos bem explicados"},
      {"label": "Paciente/cuidador", "description": "Tem proximidade pessoal com o tema; tom empático e prático"},
      {"label": "Estudante iniciante", "description": "Base acadêmica inicial; aceita mais profundidade"},
      {"label": "Profissional da área", "description": "Rigor técnico, jargão permitido, referências densas"}
    ]
  },
  {
    "question": "Qual a extensão desejada?",
    "header": "Extensão",
    "multiSelect": false,
    "options": [
      {"label": "Curto (~30 pág)", "description": "5–7 capítulos de 4–6 páginas"},
      {"label": "Médio (~80 pág)", "description": "8–12 capítulos de 6–10 páginas"},
      {"label": "Longo (~150 pág)", "description": "12–18 capítulos de 8–12 páginas"},
      {"label": "Livre", "description": "Eu sugiro a extensão a partir do tema"}
    ]
  },
  {
    "question": "Qual o tom e linha editorial?",
    "header": "Tom",
    "multiSelect": false,
    "options": [
      {"label": "Narrativo-jornalístico", "description": "Histórias reais conduzindo o conteúdo (estilo reportagem)"},
      {"label": "Didático caloroso", "description": "Professor conversando; analogias e exemplos do cotidiano"},
      {"label": "Ensaístico-reflexivo", "description": "Tom autoral, reflexivo, com voz marcada"},
      {"label": "Prático / manual", "description": "Direto ao ponto; passos, checklists, quadros"}
    ]
  },
  {
    "question": "Descreva o tema e o ângulo desejado",
    "header": "Tema",
    "multiSelect": false,
    "options": [
      {"label": "Descrição livre", "description": "Use o campo 'Notas' para descrever tema, ângulo, objetivo, o que já foi decidido"}
    ]
  }
])
```

Aguarde as respostas antes de prosseguir.

### Passo 1 — Buscar literatura e fontes reais

Execute as três buscas em paralelo e aguarde os resultados:

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-rag.sh "TEMA"
```

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-pubmed.sh "TEMA" 20
```

Também use `WebSearch` e, se necessário, o skill `firecrawl` para livros/sites/matérias de referência. Para cada fonte relevante registre: autor, título, veículo, ano, identificador (PMID / DOI / URL). **Sem identificador = não cita.**

### Passo 2 — Montar e apresentar o briefing

Crie o diretório e gere um slug a partir do título:

```bash
mkdir -p ~/ebook-output/<slug>/figuras
```

Escreva `~/ebook-output/<slug>/00-briefing.md` com esta estrutura:

```markdown
# Briefing — <Título do eBook>

## Identidade
- **Título de trabalho:** ...
- **Subtítulo:** ...
- **Público-alvo:** <resposta do passo 0>
- **Extensão estimada:** <N capítulos, ~X páginas>
- **Tom / linha editorial:** ...
- **Objetivo do livro (o que o leitor ganha):** ...
- **Promessa central (uma frase):** ...

## Fio condutor
- **Gancho de abertura:** ...
- **Arco narrativo:** ...
- **Personagem/fio recorrente:** (opcional — pessoa real/fictícia/metáfora que atravessa o livro)
- **Mensagem final:** ...

## Linha editorial visual (para as figuras)
- **Estilo:** (ex.: ilustração editorial minimalista, linhas finas, paleta terrosa)
- **Paleta:** (3–5 cores fixas, citadas em todo prompt)
- **Tipografia conceitual / mood:** ...
- **O que evitar:** (ex.: fotorrealismo, texto nas imagens, clichês)

## Capítulos propostos
1. **<Título do capítulo 1>** — promessa de 1 linha; o que o leitor entende ao terminar
2. **<Título do capítulo 2>** — ...
...
N. **<Título do capítulo N>** — ...

## Bibliografia-base (fontes reais verificadas)
1. Autor. *Título*. Veículo, ano. [PMID/DOI/URL]
2. ...
```

**Mostre o briefing ao usuário e pergunte explicitamente:**
> "Aprova este briefing? Posso ajustar título, capítulos, tom, ordem, público — diga o que mudar ou responda **'aprovado'** para começar o capítulo 1."

Após aprovação, crie também `_context.md` (uso interno, não vai pro livro final):

```markdown
# Contexto Narrativo (interno)

## Terminologia fixa
- <termo>: sempre escrito assim; não alternar com sinônimos
- ...

## Personagens / fios recorrentes
- <Nome>: descrição, função no livro, em quais capítulos aparece

## Linha editorial (resumo de 1 parágrafo)
...

## Estilo visual das figuras (resumo para colar em todo prompt)
Estilo: ... | Paleta: ... | Proibido: ...

## Capítulos já escritos
(vazio inicialmente; atualizar ao final de cada capítulo com: título, síntese de 3 linhas, ganchos deixados para capítulos futuros)
```

E `referencias.md` vazio com cabeçalho:

```markdown
# Referências — <Título>

> Lista cumulativa. Cada entrada aparece uma única vez; capítulos citam pelo índice.

```

---

## MODO CAPITULO — escrever um capítulo por vez

### Passo A — Carregar contexto

1. Ler `00-briefing.md`, `_context.md`, `referencias.md`.
2. Ler o capítulo anterior (se houver) — últimas 2 páginas para captar o gancho de saída.
3. Identificar o capítulo N alvo no briefing.

### Passo B — Discutir o capítulo antes de escrever

**Antes de redigir, apresente ao usuário um mini-plano do capítulo** e aguarde ajustes:

```markdown
## Plano — Capítulo N: <Título>

- **Promessa ao leitor:** ...
- **Abertura (gancho):** ...
- **Estrutura em seções:** 1) ... 2) ... 3) ...
- **Ponto de virada / insight central:** ...
- **Fechamento + gancho para o próximo capítulo:** ...
- **Fontes que pretendo usar:** [IDs do referencias.md + novas buscas a fazer]
- **Figuras previstas:** 2–4 ideias em 1 linha cada
```

Pergunte: "Ajusto algo ou sigo para escrever?"

### Passo C — Buscar fontes específicas do capítulo

Execute buscas dirigidas ao tema do capítulo (PubMed / RAG / WebSearch / Firecrawl). Adicione novas referências ao `referencias.md` **antes** de escrever, com índice incremental global.

### Passo D — Escrever o capítulo

Arquivo: `NN-<slug-do-capitulo>.md` (ex.: `03-o-que-acontece-no-cerebro.md`). Número com 2 dígitos.

**Regras de escrita:**

- Tom conforme briefing; respeitar terminologia fixa do `_context.md`.
- Para público leigo: analogias concretas, frases curtas, jargão só quando explicado na primeira aparição.
- **Nunca invente fatos, dados ou referências.** Todo dado factual deve estar ancorado em uma entrada do `referencias.md`.
- Citações inline no corpo: `[12]` remetendo ao índice em `referencias.md`.
- Abrir com o gancho planejado; fechar com ponte explícita para o próximo capítulo (respeitar o arco do briefing).
- Em capítulos que continuam um fio narrativo (personagem, caso, metáfora), retomá-lo explicitamente na abertura.
- Marcadores de figura no corpo do texto: `<!-- FIGURA:NN-fig-K "legenda curta" -->` onde NN = capítulo, K = ordem.

**Estrutura do arquivo:**

```markdown
# Capítulo N — <Título>

> <Epígrafe curta ou promessa em 1 linha>

<corpo do capítulo...>

<!-- FIGURA:03-fig-1 "Como o estresse altera o sono" -->

...

---

## Para saber mais (neste capítulo)
- [12] referência usada
- [15] referência usada
```

### Passo E — Gerar o arquivo de figuras

Arquivo paralelo: `NN-<slug>-figuras.md`.

```markdown
# Figuras — Capítulo N: <Título>

> Linha editorial visual fixa (do briefing):
> **Estilo:** ... | **Paleta:** ... | **Proibido:** ...

## Figura 03-fig-1 — <Legenda>

- **Aparece em:** seção "<subtítulo>" do capítulo
- **Função narrativa:** o que essa imagem precisa fazer pelo leitor
- **Conteúdo a representar:** ...
- **Composição:** enquadramento, elementos, hierarquia visual
- **Texto na imagem?** não (ou: sim, com as palavras "...")
- **Prompt (PT):**

> Ilustração editorial, <estilo fixo>, paleta <cores fixas>. Cena: <descrição concreta>. Composição: <enquadramento>. Atmosfera: <tom>. Sem texto na imagem. Sem <lista de proibições do briefing>.

## Figura 03-fig-2 — ...
```

**Todo prompt repete estilo + paleta + proibições** para garantir consistência visual ao longo do livro.

### Passo F — Atualizar `_context.md`

Adicione ao final da seção "Capítulos já escritos":

```markdown
### Capítulo N — <Título>
- Síntese (3 linhas): ...
- Gancho deixado para o capítulo N+1: ...
- Termos/conceitos introduzidos: ...
```

### Passo G — Apresentar ao usuário

Mostre:
- caminho do `.md` do capítulo
- caminho do `.md` de figuras
- 1 parágrafo de síntese
- pergunta: **"Aprova este capítulo? Posso ajustar trechos, trocar referências, reescrever seções — ou digo 'aprovado' para seguir ao capítulo N+1."**

Não avance para o próximo capítulo sem aprovação.

---

## MODO REVISAR

1. Ler o capítulo alvo + `_context.md` + `referencias.md`.
2. Aplicar as mudanças pedidas via `Edit` (preservar estrutura, marcadores de figura e índices de referência).
3. Se uma referência for removida, checar se ainda é usada em outros capítulos antes de tirar do `referencias.md`.
4. Re-apresentar a síntese e pedir nova aprovação.

---

## MODO EXPORTAR

1. Listar capítulos do diretório em ordem (`NN-*.md`, excluindo `-figuras.md` e `_*`).
2. Concatenar em `<slug>.md` precedido por capa (do `00-briefing.md`) e sumário auto-gerado dos títulos de capítulo.
3. Adicionar `referencias.md` ao final.
4. Se o usuário pediu PDF, reutilizar:

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/export-pdf.sh \
  "<output_dir>/<slug>.md" "<output_dir>/<slug>.pdf"
```

5. Reportar tamanho final, nº de capítulos, nº de figuras previstas, nº de referências.

---

## Regras invioláveis

1. **Zero alucinação de fontes.** Sem PMID/DOI/URL verificável ⇒ não cita.
2. **Aprovação humana entre etapas:** briefing → cap 1 → cap 2 → ... → export. Nunca pular.
3. **Coerência:** reler `_context.md` antes de cada capítulo; nomes, termos e personagens não mudam.
4. **Linha visual fixa:** todo prompt de figura repete estilo + paleta + proibições do briefing.
5. **Português** em todo o output (texto do livro, legendas e prompts de figura).
6. **Um capítulo = um `.md`**, com um `.md` de figuras irmão.
