# Instruções: Geração de JSON de Enrichment — Escore Plenya

## O que é isso?

Cada arquivo `xxxx-nome-do-item__pending.md` nesta pasta contém um ScoreItem do Escore Plenya com chunks científicos já selecionados via RAG. Seu trabalho é ler o arquivo, seguir o prompt contido nele e gerar o JSON de resposta.

---

## Passo a passo

### 1. Leia o arquivo `.md` na íntegra

O arquivo contém:
- **Metadados do ScoreItem** (ID, FullName, unidade, gênero, faixa etária)
- **Metadados da preparation** (quality grade, número de chunks, similaridade média)
- **Contexto e regras** (quem você é, o que o Escore Plenya é, regras inegociáveis)
- **Template JSON** com os campos que você deve preencher
- **Regras de pontuação** (apenas se o item tiver `points > 0`)
- **Chunks científicos** — os artigos que você deve usar como base

### 2. Siga as regras inegociáveis do arquivo

- Use **apenas** conhecimento médico real consolidado e os dados dos chunks
- **Não alucine, não invente** dados, estudos ou estatísticas
- Se um dado não estiver nos chunks nem for conhecimento médico estabelecido, **omita**

### 3. Gere o JSON

O JSON deve seguir **exatamente** o template do arquivo. Campos obrigatórios:

```json
{
  "score_item_id": "<UUID extraído do campo **ID:** do arquivo>",
  "clinical_relevance": "<texto técnico para médicos, 1000-5000 chars>",
  "points": 0,
  "patient_explanation": "<texto simples para pacientes, 500-1000 chars>",
  "conduct": "<conduta clínica em Markdown, 1000-5000 chars>"
}
```

- Se o arquivo indicar `"points": 0` e a nota **"não calcule pontuação"** → mantenha `0`
- Se o arquivo tiver regras de pontuação → determine o valor entre 1-50 conforme as regras

### 4. Salve o arquivo JSON

- Pasta: `score-enrich/json/`
- Nome: mesmo nome do `.md`, trocando `__pending.md` por `__processed.json`
- Exemplo: `0001-controle-de-diabetes-mellitus__pending.md` → `0001-controle-de-diabetes-mellitus__processed.json`
- O arquivo deve conter **apenas JSON puro**, sem nenhum texto fora do JSON

---

## Exemplo de instrução para o agente

```
Leia o arquivo `/path/score-enrich/0001-controle-de-diabetes-mellitus__pending.md`
na íntegra e siga exatamente as instruções contidas nele.

O arquivo contém o contexto, as regras e o template JSON que você deve preencher.
Siga todas as regras rigorosamente, especialmente as regras inegociáveis sobre
não alucinar.

Salve o resultado em `/path/score-enrich/json/0001-controle-de-diabetes-mellitus__processed.json`
— JSON puro, sem nenhum texto fora do JSON.
```

---

## Nomenclatura dos arquivos

| Status | Sufixo |
|---|---|
| Aguardando processamento | `__pending.md` |
| JSON gerado | `__processed.json` (em `json/`) |

---

## Campos do JSON e correspondência com o banco

| Campo JSON | Coluna no banco (`score_items`) |
|---|---|
| `score_item_id` | `id` |
| `clinical_relevance` | `clinical_relevance` |
| `points` | `points` |
| `patient_explanation` | `patient_explanation` |
| `conduct` | `conduct` |
