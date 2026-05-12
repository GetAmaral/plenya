# RAG — Como buscar evidência ao gerar drafts

> Procedimento de busca científica para respostas que precisam de embasamento.
> Não usar em elogios, agendamentos, spam — só em perguntas clínicas (Tipos 1, 2 do PLAYBOOK).

---

## Camadas de busca (use nesta ordem)

### Camada 1 — RAG Plenya (sempre primeiro)

A base científica do Plenya tem ~1.000+ artigos curados, com embeddings + busca semântica. É a fonte mais alinhada com o que o Dr Getúlio já validou.

```bash
bash /home/user/plenya/.claude/skills/responder-insta/scripts/search-rag.sh "<query>" 5
```

Retorna até 5 artigos relevantes com:
- Título e autores
- Trecho mais relevante (chunk)
- DOI/PMID se disponível
- Score de similaridade

**Como formar a query:**
- Use termos clínicos em português (a base é mista PT/EN, mas headers são PT)
- Combine o conceito principal + contexto (ex.: "creatina suplementação pós-menopausa" em vez de só "creatina")
- Se a primeira query der ruim, refine com termos mais específicos

**Se o backend `localhost:3001` não responder:** a base não está rodando. Pular pra Camada 2 e avisar no draft que "fontes citadas vêm de literatura aberta, não da curadoria interna".

---

### Camada 2 — WebSearch / Literatura aberta

Use quando:
- RAG retornou pouco/nada relevante
- Tema exige evidência **muito recente** (≥2024)
- Tema fora do escopo da base Plenya (ex.: ortopedia, dermatologia)

**Estratégia:**

1. `WebSearch` com termos **em inglês** (literatura é majoritariamente EN):
   - "Tirzepatide chronic kidney disease 2024 trial"
   - "GLP-1 receptor agonist eGFR decline meta-analysis"
2. Priorize: **PubMed, NEJM, Lancet, BMJ, JAMA, Circulation, KDIGO/ESC/ACSM/ADA guidelines**
3. Evite: blogs, "Verywell Health", Healthline (citações secundárias)

Se WebSearch falhar, pode usar **firecrawl-search** (skill instalada) — busca com extração de markdown completo.

---

### Camada 3 — PubMed direto (mais específico)

Se precisa de um estudo muito específico (PMID, ano, autor):

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-pubmed.sh "<query>" 10
```

(Script do skill `/aula` — pode ser reusado aqui)

---

## Como citar no draft

**Regra de ouro:** cita evidência em **prosa**, não em formato acadêmico.

✅ Bom:
> "O FLOW (estudo com semaglutida, classe próxima) reduziu progressão da DRC em 24%, e a tirzepatida vem mostrando achados na mesma direção."

❌ Ruim:
> "Conforme demonstrado por Perkovic et al. (NEJM 2024;390:1745-1756) (FLOW Trial), o uso de semaglutida 1.0 mg semanal..."

**Padrões aceitáveis:**

- Nome do estudo: **OK e desejável** (FLOW, SURMOUNT-3, SURPASS-4, ATLAS-ACS, JUPITER)
- Sociedade médica: **OK** (KDIGO 2024, ESC 2023, ACSM, ADA)
- Ano da publicação se relevante: **OK** ("um trial de 2024 mostrou...")
- Revista/journal: **evitar** (paciente não sabe diferença entre NEJM e Healthline)
- PMID/DOI: **nunca** em comentário/DM
- Autor: **nunca** (não tem peso pra leigo)

---

## Quanto citar?

| Tipo de pergunta | Quantas fontes mencionar |
|---|---|
| Pergunta clínica simples | 1-2 estudos/diretrizes |
| Pergunta clínica complexa | 2-4 estudos/diretrizes |
| Pergunta de mecanismo | 1 estudo/livro-texto + 1 sociedade |
| Pergunta de "isso é verdade?" | 1-2 fontes que apoiam OU desafiam a afirmação |

**Mais de 4 = paper, não comentário.** Cortar.

---

## Quando NÃO buscar evidência

- **Elogios** (Tipo 4)
- **Agendamentos** (Tipo 7)
- **Spam** (Tipo 6)
- **Posts pessoais** (Dia das Mães, aniversário, etc.)
- **Pergunta de tom emocional** ("doutor, estou com medo de tomar...") — aqui o foco é acolhimento, não dados

---

## Erro comum a evitar

**Alucinar PMID, DOI ou autor.** Se a busca não retornou a fonte exata, é melhor escrever menos do que inventar. Em caso de dúvida:

❌ "Há estudos que mostram..." (vago, evitar)
✅ "A literatura aponta..." (também vago, mas honesto)
✅ "No FLOW (2024) com semaglutida, a redução foi de 24%" (específico, só se foi visto)

Validação: **se cita estudo por nome, o nome do estudo deve aparecer nos resultados de busca da Camada 1 ou 2.** Não cita estudo só por intuição.
