# Estrutura do Livro ANTES — Multilíngue

Fonte de verdade do livro, organizada para suportar publicação simultânea em múltiplos idiomas.

## Estrutura de pastas

```
Ebook 1 - Performance e Longevidade/
├── README.md                    ← este arquivo
├── glossario.yaml               ← termos-chave e traduções oficiais (FONTE DE VERDADE TERMINOLÓGICA)
├── folha-de-rosto.pdf           ← folha de rosto para cadastro CBL (pt-BR)
├── folha-de-rosto.png
│
├── md/                          ← CONTEÚDO TEXTUAL — um .md por capítulo, um idioma por subpasta
│   ├── pt-BR/                   ← fonte (Português do Brasil)
│   │   ├── 00-indice.md
│   │   ├── 00a-creditos.md
│   │   ├── 00b-introducao.md
│   │   ├── 01-homem-que-morreu-saudavel.md
│   │   ├── 02-quatro-assassinos-silenciosos.md
│   │   ├── 03-seu-corpo-esta-envelhecendo.md
│   │   ├── 04-mapa-dos-biomarcadores.md
│   │   ├── 05-seu-coracao-esta-falando.md
│   │   ├── 06-saude-metabolica.md
│   │   ├── 06b-parte-iii-intro.md
│   │   ├── 07-atividade-fisica.md              (A1)
│   │   ├── 08-alimentacao-suplementacao.md     (A2 — inclui pele/cabelo/unhas + PFS)
│   │   ├── 09-sistemas-cardio-renal-hepatico.md (G1)
│   │   ├── 10-paineis-bioquimicos.md           (G2)
│   │   ├── 11-genomica-exposicoes.md           (G3)
│   │   ├── 12-integracao-mente-corpo.md        (I1 — trabalho interno)
│   │   ├── 13-conexao-proposito-sentido.md     (I2 — dimensão relacional)
│   │   ├── 14-ritmo-circadiano-e-repouso.md    (R)
│   │   ├── 15-seu-placar-de-longevidade.md
│   │   ├── 16-quando-procurar-especialista.md
│   │   ├── 17-manifesto-agir.md
│   │   ├── 18-referencias-recursos.md
│   │   ├── agradecimentos.md
│   │   ├── sobre-o-autor.md
│   │   └── frontmatter.md
│   ├── en/                      ← tradução inglês (vazio até tradução)
│   ├── es/                      ← tradução espanhol
│   ├── fr/                      ← tradução francês
│   └── de/                      ← tradução alemão
│
├── figuras/                     ← FIGURAS DOS CAPÍTULOS — regeradas por idioma
│   ├── pt-BR/                   ← PNGs por capítulo (Cap01 Fig01.PNG etc.); ver md/pt-BR/figuras-a-atualizar.md para mapa de rename pós-reestruturação AGIR
│   ├── en/
│   ├── es/
│   ├── fr/
│   └── de/
│
├── capas/                       ← CAPAS — uma por idioma
│   ├── pt-BR/capa.jpg           ← capaGrande04.jpg (1920x3072, KDP-compliant)
│   ├── en/                      ← Before cover (futuro)
│   ├── es/
│   ├── fr/
│   └── de/
│
├── fotos/                       ← FOTOS DO AUTOR — não traduzem, compartilhadas
│   ├── getulio_bw_halfbody_1000.jpg     (EPUB/Kindle — todas as línguas)
│   ├── getulio_bw_halfbody_fullres.jpg  (impresso)
│   ├── getulio_color_halfbody_1200.jpg  (Author Central Amazon)
│   ├── getulio_color_square_1200.jpg    (avatar social)
│   └── getulio_color_fullbody.jpg       (press kit)
│
├── briefings/                   ← REFERÊNCIA INTERNA (não publicado)
│   ├── Briefing Ebook 1 AGIR Corrigido.pdf
│   ├── Briefing Figuras Capitulo 01-12.pdf
│   └── Briefing Protocolo Longevidade.pdf
│
└── build/                       ← OUTPUT — EPUBs e PDFs gerados por idioma
    ├── Antes-pt-BR.epub         (quando montado)
    ├── Before-en.epub           (futuro)
    └── ...
```

## Regras de ouro

### 1. `pt-BR` é a fonte

Toda edição começa em `md/pt-BR/`. Mudanças em outras línguas **nunca** retroalimentam o pt-BR — são traduções. Se um erro ou melhoria é identificado em uma tradução, ele deve ser corrigido primeiro em pt-BR e depois repropagado.

### 2. Glossário é lei

`glossario.yaml` é a fonte de verdade dos termos-chave e suas traduções. Todo tradutor (humano ou máquina) recebe o glossário antes de começar. Novos termos importantes descobertos em capítulos devem ser adicionados ao glossário antes de serem traduzidos.

### 3. Frontmatter YAML em cada capítulo

Cada arquivo `.md` começa com YAML frontmatter contendo:
- `id` — identificador curto do capítulo (01, 04, 06b, bm1...)
- `language` — código ISO (pt-BR, en, es, fr, de)
- `title` — título no idioma do arquivo
- `title_en_suggested` — tradução sugerida para inglês (referência)
- `translation_status` — `source` | `in_progress` | `review` | `done`
- `last_reviewed` — data ISO
- `figures` — lista de IDs das figuras usadas no capítulo

### 4. Figuras são regeradas, não editadas

Cada idioma tem sua própria pasta `figuras/<lang>/`. Quando o conteúdo muda no pt-BR, a versão regenerada no pt-BR deve ser propagada para as outras línguas.

### 5. Nomes e marcas

Ver `glossario.yaml` seção `keep_original` e `never_translate`. Regra geral:
- Marcas não traduzem: **Plenya**, **Método AGIR** (exceto ACTS em EN), **ANTES** (título localiza)
- Personagens dos casos clínicos mantêm nome brasileiro (Ricardo, Fernanda, André…)
- Siglas médicas internacionais preservam (HbA1c, ApoB, CAC score, NT-proBNP…)

## Workflow de tradução

1. **Finalizar pt-BR** — `translation_status: source` e `last_reviewed` atualizado.
2. **Congelar a versão** — git tag `pt-BR-v1.0`.
3. **Enviar para tradutor** o pacote `md/pt-BR/` + `glossario.yaml`.
4. **Tradutor entrega** `md/<lang>/` com mesmos nomes de arquivo, frontmatter atualizado com `language: <lang>` e `translation_status: review`.
5. **Revisão médica bilíngue** — revisor confere terminologia clínica contra o glossário.
6. **Regerar figuras** em `figuras/<lang>/` com textos traduzidos.
7. **Traduzir capa** — arte em `capas/<lang>/capa.jpg`.
8. **Montar EPUB** para aquele idioma em `build/Before-en.epub` etc.
9. **Validar com `epubcheck`**.
10. **Publicar** com ISBN próprio do idioma.

## Comandos úteis (a implementar em build)

```bash
# Montar EPUB para um idioma
./build.sh pt-BR
./build.sh en

# Validar EPUB gerado
epubcheck build/Antes-pt-BR.epub

# Contar palavras de um idioma
wc -w md/pt-BR/*.md | tail -1
```

---

## Estrutura AGIR em 18 capítulos

O livro foi reestruturado para alinhar com a arquitetura pública do Método AGIR no site (`apps/site/lib/agir-structure.ts`). Os quatro pilares se distribuem agora em oito capítulos:

| Pilar | Capítulos | Conteúdo |
|---|---|---|
| **A** — Atividade Física, Alimentação e Suplementação Inteligente | 7, 8 | Treino; alimentação + suplementação + camada visível (pele/cabelo/unhas com posicionamento sobre PFS) |
| **G** — Gestão Clínica e Metabólica | 9, 10, 11 | Sistemas cardio-reno-hepático-metabólicos; painéis bioquímicos e hormônios; genômica, epigenética e exposições |
| **I** — Integração Mente-Corpo | 12, 13 | I1: mente individual (avaliação, regulação, função cognitiva); I2: conexão, propósito e sentido |
| **R** — Ritmo Circadiano e Repouso | 14 | Sono, cronobiologia, apneia |

As Partes IV (Placar + Especialista) e V (Manifesto + Referências) ocupam os Caps 15–18.

## Status de tradução por idioma

As traduções para EN/ES/FR/DE ainda não foram iniciadas. Quando começarem, o pacote a entregar para o tradutor deve incluir todo o conteúdo atualizado — especialmente os capítulos novos e as seções adicionadas durante a reestruturação AGIR:

**Capítulos novos ou substancialmente reformulados (prioridade alta na tradução):**
- Cap 7 (A1) — novo, dedicado à atividade física
- Cap 8 (A2) — reformulado + **Seção 3 nova** (Camada Visível: pele/cabelo/unhas, posicionamento sobre síndrome pós-finasterida / PFS)
- Cap 9 (G1) — novo, inclui Fronteira da Geromedicina com nota sobre IV therapy
- Cap 10 (G2) — reformulado
- Cap 11 (G3) — novo
- Cap 12 (I1) — reescrito com seções novas: **Avaliação Psicológica** (PHQ-9, GAD-7, AUDIT, PCL-5, UCLA, burnout CID-11), **Função Cognitiva** (reserva cognitiva, MoCA, SCD, p-tau217 plasmático, FINGER/MAPT/POINTER, ACTIVE trial)
- Cap 13 (I2) — novo inteiro: Vida Sexual (incl. AHA 2012 pós-evento CV), Vínculos Sociais, Telas/Atenção, Propósito/Ikigai, Espiritualidade, bidirecionalidade, Encaminhamento
- Cap 14 (R) — mantido, renumerado
- Cap 18 (Referências) — expandido com novas referências: PHQ-9, GAD-7, Stern, US POINTER 2025, p-tau217, Princeton Consensus, AHA 2012 pós-CV, Mark, Haidt, Sone 2008, Frankl, Koenig, Meijer pós-IAM

**Novos termos adicionados ao glossário** (ver `glossario.yaml` — seções expandidas em 2026-04-21): síndrome pós-finasterida (PFS), camada visível, ikigai, eudaimonia, reserva cognitiva, p-tau217, PDE5, IV therapy.

---

**Última atualização:** 2026-04-21
