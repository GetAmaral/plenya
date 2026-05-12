# Wikidata — draft de criação (Dr. Getúlio + Livro Antes)

> **Como usar:** logue em wikidata.org (crie conta se preciso), clique "Create a new item" e preencha campo a campo seguindo as tabelas abaixo. Cada **statement** (P-code) precisa de pelo menos uma **referência** (URL externa) — quanto mais, mais robusto e menos chance de remoção pela comunidade.

---

## ⚙️ Procedimento operacional (resumo)

1. **Cria conta:** wikidata.org → "Create account" → confirma email
2. **Cria Q-item Pessoa primeiro** (livro depende dele)
3. **Cria Q-item Livro** referenciando o Pessoa via P50 (author)
4. Aguarda ~5 minutos pra Wikidata indexar
5. **Verifica:** busca "Getúlio Amaral Filho" no buscador Wikidata — deve aparecer
6. **Avisa o site Plenya/getulio:** vou adicionar Q-id no Person Schema (`identifier` PropertyValue) — sinal pro Google ler Knowledge Graph

---

# Q-ITEM 1 — PESSOA (Dr. Getúlio Amaral Filho)

## Labels (multilingual)

| Idioma | Texto |
|---|---|
| English (en) | `Getúlio Amaral Filho` |
| Português (pt-BR) | `Getúlio José Mattos do Amaral Filho` |
| Português (pt) | `Getúlio Amaral Filho` |

## Aliases

**en aliases:**
- `Getulio Amaral Filho`
- `Dr. Getúlio Amaral Filho`
- `Getulio J. M. Amaral Filho`
- `AMARAL FILHO, G. J. M.`
- `Getúlio José Mattos do Amaral Filho`

**pt-BR aliases:**
- `Getulio Amaral Filho`
- `Dr. Getúlio Amaral Filho`
- `Getulio José Mattos do Amaral Filho`
- `AMARAL FILHO, G. J. M.`

## Description

| Idioma | Texto |
|---|---|
| English (en) | `Brazilian nephrologist, internal medicine physician and author` |
| Português (pt-BR) | `nefrologista brasileiro, clínico geral e autor` |

---

## Statements (propriedades)

> **Como adicionar cada statement:** clique "+ add statement" → digite o nome da propriedade (Wikidata sugere) → escolhe o P-code → escolhe o valor (Q-item para itens conhecidos, ou texto livre para identificadores).
> **Como adicionar referência:** clique no statement → "+ add reference" → propriedade `reference URL (P854)` ou `stated in (P248)` + URL.

### Identidade base

| Propriedade | P-code | Valor | Q-code (se Wikidata item) |
|---|---|---|---|
| instance of | **P31** | human | **Q5** |
| sex or gender | **P21** | male | **Q6581097** |
| country of citizenship | **P27** | Brazil | **Q155** |
| date of birth | **P569** | 17 August 1981 | precisão "day" |
| place of birth | **P19** | Londrina | **Q170354** (Londrina, PR) |
| native language | **P103** | Portuguese | **Q5146** |
| languages spoken/written | **P1412** | Portuguese, English | **Q5146**, **Q1860** |
| father | **P22** | Getulio Jose Mattos do Amaral | string (ou Q-item se ele tiver Wikidata) |
| mother | **P25** | Ana Maria Calcavara Amaral | string |

### Nome

| Propriedade | P-code | Valor |
|---|---|---|
| given name | **P735** | `Getúlio` (busca Q-item; se não existir, cria string) |
| family name | **P734** | `Amaral Filho` (string) |

### Ocupação (P106) — adicionar 3 statements separados:

| Valor | Q-code |
|---|---|
| physician | **Q39631** |
| nephrologist | **Q19595170** |
| author | **Q482980** |

### Educated at (P69) — 3 statements separados:

Cada um com qualifier P512 (academic degree) + P580 (start time) + P582 (end time):

| Instituição | Q-code | Diploma | Período |
|---|---|---|---|
| Universidade Estadual de Londrina | **Q3550114** | Graduação em Medicina | 1999–2004 |
| Santa Casa de Londrina (ISCAL) | **buscar** — se não existir, usa string | Especialização Clínica Médica + Nefrologia | 2005–2008 |
| Associação Brasileira de Medicina Funcional Integrativa (ABMFI) | string | Pós-graduação Medicina Funcional Integrativa | 2026 |

### Position held (P39) — 6 statements separados:

| Cargo | Período | Empregador |
|---|---|---|
| coordinator of nephrology medical residency | desde 2009 | Santa Casa de Londrina |
| founder of internal medicine medical residency | 2014 | Santa Casa de Londrina |
| clinical director | desde 2026 | Plenya |
| medical director (technical) | atual | DaVita Intra Hospitalar Londrina |
| professor (lecturer) | 2013–2014 | PUC Londrina |
| committee member (Cardiorenal Committee) | atual | Sociedade Brasileira de Nefrologia |

### Employer (P108) — 5 statements:

- Plenya (string)
- Nefroclínica Londrina (string)
- Santa Casa de Londrina / ISCAL (string ou Q-item)
- Hospital Araucária (string)
- Hospital Unimed Londrina (string)

### Member of (P463) — 3 statements:

- Sociedade Brasileira de Nefrologia (string ou Q-item)
- Sociedade Paranaense de Nefrologia (string)
- Associação Brasileira de Medicina Funcional Integrativa (string)

### Notable work (P800)

- Antes — A Janela Silenciosa entre o Normal e o Ótimo → **Q-id do Q-item 2 (criar primeiro)**

### Award received (P166)

| Prêmio | Ano | Outorgante |
|---|---|---|
| Specialist Certification in Nephrology | 2008 | SBN + AMB |
| Member of Cardiorenal Committee | 2025 | Sociedade Brasileira de Nefrologia |

### Notable work (P800) — extras além do livro

| Obra | Tipo | Identificador |
|---|---|---|
| Antes (livro) | Q-id do Q-item 2 | criar primeiro |
| NAI — Nephrology Artificial Intelligence | Software (Programa de Computador) | INPI 102025001407-6 |

### Work location (P937)

| Valor | Q-code |
|---|---|
| Londrina | **buscar** |

### Identificadores externos (cada um statement com P-code dedicado)

| Plataforma | P-code | Valor |
|---|---|---|
| ORCID iD | **P496** | `0009-0009-2506-2455` |
| Lattes ID | **P6829** | `2492350974849886` |
| LinkedIn personal profile ID | **P6634** | `getulio-amaral-filho-951981404` |
| Instagram username | **P2003** | `drGetulioAmaralFilho` |
| Doctoralia (Brazil) ID | **P10632** | `getulio-amaral-filho` |

### Site oficial (P856)

`https://drgetulioamaralfilho.com.br`

### Image (P18)

Upload primeiro a foto profissional pra Commons (commons.wikimedia.org → "Upload file" → use `drgetulioamaralfilho.com.br/images/getulio-square.jpg`) → depois referencia o nome do arquivo.

> **Sem foto Commons agora?** Pula essa propriedade — pode adicionar depois.

---

## Referências (cole em CADA statement quando possível)

Use sempre **P854 (reference URL)**:

| Statement | URL de referência |
|---|---|
| ORCID, occupation physician | https://orcid.org/0009-0009-2506-2455 |
| Lattes, educated at, occupation | http://lattes.cnpq.br/2492350974849886 |
| Position held (coordinator residency) | https://iepi.iscal.com.br/br/ensino/residencia-medica |
| Member of (SBN Comitê Cardiorrenal) | https://sbn.org.br/medicos/a-sbn/comites/ |
| Notable work (livro) | https://drgetulioamaralfilho.com.br/livros/antes |
| Site oficial | https://drgetulioamaralfilho.com.br |
| Publicação BJN 2024 | https://www.bjnephrology.org/en/article/assessment-of-dialysis-practices-for-acute-patients-among-nephrologists-in-hospital-services-across-brazil-a-cross-sectional-survey/ |

---

# Q-ITEM 2 — LIVRO (Antes — A Janela Silenciosa entre o Normal e o Ótimo)

## Labels (multilingual)

| Idioma | Texto |
|---|---|
| English (en) | `BEFORE — The silent window between normal and optimal — a decade where health is decided` |
| Português (pt-BR) | `Antes — A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida` |

## Aliases

**en:**
- `BEFORE`
- `The Silent Window`
- `A Decade Between Normal and Optimal`
- `Antes`

**pt-BR:**
- `Antes`
- `A Janela Silenciosa entre o Normal e o Ótimo`

## Description

| Idioma | Texto |
|---|---|
| English (en) | `2026 book by Getúlio Amaral Filho on preventive medicine, longevity and the silent window between normal and optimal health` |
| Português (pt-BR) | `livro de 2026 de Getúlio Amaral Filho sobre medicina preventiva, longevidade e a janela silenciosa entre o normal e o ótimo` |

---

## Statements

### Identidade da obra

| Propriedade | P-code | Valor | Q-code |
|---|---|---|---|
| instance of | **P31** | written work | **Q47461344** (ou Q571 — book) |
| author | **P50** | Getúlio Amaral Filho | **Q-id do Q-item 1** |
| publisher | **P123** | Plenya | string |
| publication date | **P577** | 2026 | precisão year |
| country of origin | **P495** | Brazil | **Q155** |

### Idiomas (P407) — 2 statements:

- Brazilian Portuguese — **Q750553**
- English — **Q1860**

### ISBNs (P212 — ISBN-13) — 4 statements separados, cada um com qualifiers:

| ISBN | Qualifier P437 (distribution format) | Qualifier P407 (language) |
|---|---|---|
| **978-65-02-06742-0** | EPUB / e-book | Brazilian Portuguese |
| **978-65-02-07691-0** | paperback / softcover | Brazilian Portuguese |
| **978-65-975814-0-5** | EPUB / e-book | English |
| **978-65-975814-1-2** | paperback / softcover | English |

### Conteúdo (P921 — main subject) — múltiplos statements:

- preventive medicine
- nephrology
- longevity
- healthspan
- biomarkers
- metabolic health
- preventive cardiology

### Web

| Propriedade | P-code | Valor |
|---|---|---|
| full work available at URL | **P953** | `https://drgetulioamaralfilho.com.br/livros/antes` |
| official website | **P856** | `https://drgetulioamaralfilho.com.br/livros/antes` |

### Referências para o livro

- ISBN registry brasileira (BN/CBL): https://www.bn.gov.br/explore/acervos/isbn (busca pelo ISBN)
- Amazon edição PT: https://a.co/d/0fxsmomI
- Amazon edição EN: https://a.co/d/00Jgudq4
- Site do autor: https://drgetulioamaralfilho.com.br/livros/antes

---

## ✅ Checklist final

Após criar os 2 Q-items:

- [ ] Anotar o **Q-id de cada um** (formato `Q12345678`)
- [ ] Me passar os 2 Q-ids pra eu adicionar no Person Schema do site Getúlio
- [ ] Verificar no Google após 24-72h: busca "Dr. Getúlio Amaral Filho" e veja se aparece painel lateral começando a se formar
- [ ] (Opcional) editar próximas vezes pra adicionar mais statements (palestras, capítulos do livro como obras independentes)

---

## ⚠️ Notas operacionais

- **Wikidata é colaborativo** — outros editores podem revisar/refinar/desafiar statements. Por isso referências externas verificáveis são essenciais.
- **Notabilidade** — seu perfil atende: ORCID + Lattes + livro publicado com 4 ISBNs + cargo CNRM/MEC + comitê SBN + publicação científica. Não deve ter problema.
- **Privacidade da data de nascimento** — Wikidata tradicionalmente publica datas de nascimento de pessoas notáveis. Se preferir privacidade, pode usar precisão "year" só (1981) em vez de "day" (17/08/1981).
- **Plenya** — não tem Q-item ainda; criar quando houver matéria de imprensa publicada (ex: cobertura no jornal local sobre lançamento do livro). Pula por enquanto.

---

**Próximo passo após criação:** me passa os 2 Q-ids → eu adiciono em `apps/site-getulio/components/seo/person-schema.tsx` como `identifier` PropertyValue (`P-code: Wikidata QID`) e em `sameAs` (`https://www.wikidata.org/wiki/Q...`). Isso fecha o ciclo Site ↔ Wikidata ↔ Knowledge Graph do Google.
