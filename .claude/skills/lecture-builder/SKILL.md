# Medical Lecture Builder — Instruções do Skill /aula

> Este arquivo é referência interna usada por `.claude/commands/aula.md`.

---

## Modo de Operação

Analise `$ARGUMENTS`:

- "editar slide N" / "modificar slide N" → **MODO EDIT**
- "regenerar imagem slide N" / "trocar imagem slide N" → **MODO IMAGE**
- "exportar" / "gerar pdf" sem novo tema → **MODO EXPORT**
- Qualquer outra coisa → **MODO CREATE**

---

## MODO CREATE — Nova Aula

### Passo 0: Coletar briefing completo via AskUserQuestion

**OBRIGATÓRIO:** Use a ferramenta `AskUserQuestion` com as 4 perguntas abaixo em uma única chamada. Não pergunte em texto livre.

```
AskUserQuestion(questions=[
  {
    "question": "Qual o nível do público?",
    "header": "Público-alvo",
    "multiSelect": false,
    "options": [
      {"label": "Residência médica", "description": "R1–R3, conhecimento clínico intermediário"},
      {"label": "Graduação", "description": "Estudantes de medicina, base ainda em formação"},
      {"label": "Pós-graduação / Fellows", "description": "Especialistas em formação avançada"},
      {"label": "Congresso / Atualização", "description": "Plateia mista, foco em novidades e cases"}
    ]
  },
  {
    "question": "Qual a duração da aula?",
    "header": "Duração",
    "multiSelect": false,
    "options": [
      {"label": "30 minutos", "description": "~8–10 slides, foco em 1 tema central"},
      {"label": "45 minutos", "description": "~12–14 slides, tema central + caso clínico"},
      {"label": "60 minutos", "description": "~16–18 slides, cobertura completa do tema"},
      {"label": "90 minutos", "description": "~22–26 slides, aprofundamento + múltiplos casos"}
    ]
  },
  {
    "question": "Quem está apresentando?",
    "header": "Apresentador",
    "multiSelect": false,
    "options": [
      {"label": "Dr. Getúlio Amaral Filho — Nefrologista", "description": "Professor de Nefrologia, aula para residência de nefrologia (padrão)"},
      {"label": "Outro apresentador", "description": "Informe nome e especialidade na descrição abaixo"}
    ]
  },
  {
    "question": "Descreva o que você quer na aula",
    "header": "Briefing",
    "multiSelect": false,
    "options": [
      {"label": "Foco diagnóstico", "description": "Ênfase em como reconhecer, investigar e classificar o problema"},
      {"label": "Foco terapêutico", "description": "Ênfase em condutas, protocolos e manejo prático"},
      {"label": "Fisiopatologia + clínica", "description": "Explicar o mecanismo e conectar com apresentação clínica"},
      {"label": "Descrição livre (use o campo de notas)", "description": "Digite sua instrução específica no campo 'Notas' desta opção"}
    ]
  }
])
```

Aguarde as respostas antes de prosseguir. Use as escolhas para calibrar tom, profundidade e estrutura da aula.

```bash
mkdir -p ~/lecture-output/[slug]/images
```

---

### Passo 1: Buscar literatura ANTES de qualquer outro passo

Execute **ambos** os scripts e aguarde as respostas completas.

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-rag.sh "TEMA_AQUI"
```

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-pubmed.sh "TEMA_AQUI" 10
```

**O que fazer com os resultados:**
- Monte uma lista numerada de referências reais com: autores (sobrenome et al.), título, journal, ano, PMID
- Anote quais referências suportam cada ponto factual
- Se os scripts falharem, use WebSearch com termos em inglês para obter PMIDs reais antes de continuar

**NUNCA invente ou alucine referências. Dado sem PMID real = não citar como artigo.**

---

### Passo 2: Definir o fio condutor (storytelling)

Com base no briefing do usuário e na literatura encontrada, defina **antes de escrever qualquer slide**:

- **Gancho de abertura:** uma situação clínica concreta, um dado surpreendente ou uma pergunta que cria tensão ("Você já atendeu um paciente que fez tudo certo e mesmo assim...?")
- **Caso clínico central:** um paciente fictício mas verossímil que aparece no início (apresentação), retorna no meio (complicação ou achado) e fecha no final (desfecho, lição). O caso deve ter nome, idade, queixa principal e dados objetivos realistas.
- **Arco narrativo:** introdução do problema → aprofundamento fisiopatológico → momento de virada clínica → resolução e mensagem
- **Frase de fechamento:** a única coisa que o público deve levar embora

Esses elementos devem estar no outline.json e costurar explicitamente os slides.

---

### Passo 3: Montar outline.json

```json
{
  "title": "Título da Aula",
  "audience": "residência",
  "duration_min": 60,
  "output_dir": "/home/user/lecture-output/slug",
  "hook": "Dr. X atende uma mulher de 48 anos...",
  "clinical_case": {
    "name": "Maria, 48 anos",
    "presentation": "HAS + obesidade central + glicemia 118 mg/dL em check-up",
    "twist": "Slide 7: ECG com alterações — o que mudou?",
    "resolution": "Slide 12: desfecho após 6 meses de manejo"
  },
  "closing_message": "Uma frase. A mensagem que fica.",
  "sections": [
    {
      "slide_number": 3,
      "title": "Fisiopatologia",
      "duration_min": 10,
      "key_points": [
        "Resistência à insulina: mecanismo central",
        "Hiperinsulinemia compensatória → disfunção endotelial",
        "IL-6, TNF-α, PCR elevados"
      ],
      "storytelling_link": "Aqui voltamos à Maria: seus exames mostram exatamente esse perfil",
      "refs": [1, 3],
      "image_prompt": "Diagram of insulin resistance cascade..."
    }
  ],
  "references": [
    {
      "index": 1,
      "authors": "Grundy SM et al.",
      "title": "Diagnosis and Management of the Metabolic Syndrome",
      "journal": "Circulation",
      "year": 2005,
      "pmid": "16157765"
    }
  ]
}
```

**Mostre o outline ao usuário. Aguarde confirmação antes de continuar.**

---

### Passo 4: Gerar slides.md

**Tom absoluto:** médico apresentando para colegas médicos. Direto, clínico, factual.

❌ NUNCA:
- "O residente deve saber que..."
- "É importante que o aluno compreenda..."
- "Ao final desta aula, o participante será capaz de..."
- "Vamos entender o conceito de..."
- "Como todos sabemos..."
- Qualquer meta-linguagem pedagógica ou corporativa

✅ SEMPRE:
- Dados concretos com números reais: "30% da população adulta brasileira"
- Verbos clínicos diretos: "Rastreamos com...", "A conduta é...", "Evitamos..."
- Tensão narrativa: o caso clínico reaparece nos slides de maior impacto
- Frases curtas. Sem enrolação.

**Estrutura obrigatória (adapte quantidade de slides à duração):**

```markdown
---
marp: true
theme: medical
paginate: true
header: 'Título da Aula'
footer: 'Apresentador | Instituição | Ano'
---

<!-- SLIDE:1 "Título" -->
# Título Completo

**Subtítulo**

*Apresentador — Serviço — Data*

---

<!-- SLIDE:2 "Caso Clínico" -->
## O Caso

**Maria, 48 anos.** Vem ao check-up sem queixas. Traz exames.

| Dado | Valor |
|------|-------|
| Cintura | 94 cm |
| PA | 138/88 mmHg |
| Glicemia jejum | 118 mg/dL |
| TG | 198 mg/dL |
| HDL | 38 mg/dL |

*O que você faz agora?*

---

<!-- SLIDE:3 "Epidemiologia" -->
## 3. Por que isso importa

- **30%** dos adultos brasileiros têm síndrome metabólica *(Pititto et al., 2020)*
- Risco cardiovascular **2–3×** maior *(Mottillo et al., 2010)*
- 70% evoluem para DM2 em 10 anos sem intervenção

<div class="alert">
Ela é a principal causa modificável de infarto em mulheres na perimenopausa
</div>

*[1, 2]*

---

<!-- SLIDE:N "Caso Clínico — Desdobramento" -->
## Voltando à Maria...

*[Slide de caso clínico intermediário — complicação ou achado inesperado]*

---

<!-- SLIDE:N "Caso Clínico — Desfecho" -->
## Maria — 6 meses depois

*[Slide de fechamento do caso — resultado do manejo]*

---

<!-- SLIDE:N "Mensagem Final" -->
## Uma frase para levar embora

> "Frase de fechamento impactante"

---

<!-- SLIDE:N "Referências" -->
## Referências

1. Grundy SM et al. Circulation. 2005. PMID: 16157765
2. Mottillo S et al. JACC. 2010. PMID: 20863953
```

**Regras de citação:**
- Inline quando o dado é específico de um artigo: `*(Autor et al., Ano)*`
- Ao final do slide com múltiplas referências: linha `*[N, M]*`
- Slide final de referências: formato `N. Autor et al. Journal. Ano. PMID: XXXXX`
- Apenas PMIDs reais obtidos no Passo 1

**Regras de formatação:**
- Máx. 6 bullets por slide
- Tabelas para critérios, DD, comparações de regimes
- `<div class="concept">` — conceito-chave (box azul)
- `<div class="alert">` — alerta clínico (box vermelho)
- `<div class="evidence">` — evidência/guideline (box verde)
- Placeholder de imagem: `<!-- IMAGE:section-N.png -->`

---

### Passo 5: Imagens DALL-E 3

Para cada seção (exceto capa, casos clínicos e referências):

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/generate-image.sh \
  "PROMPT_DA_SECAO" "section-N.png" "OUTPUT_DIR"
```

Substitua `<!-- IMAGE:section-N.png -->` por `![](images/section-N.png)`.

---

### Passo 6: Exportar PDF

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/export-pdf.sh \
  "[output_dir]/slides.md" "[output_dir]/slides.pdf"
```

---

### Passo 7: Resumo

```
✅ [output_dir]/
   slides.md | slides.pdf | outline.json | images/

/aula editar slide N "instrução"
/aula regenerar imagem slide N
/aula exportar
```

---

## MODO EDIT

```bash
ls ~/lecture-output/*/outline.json 2>/dev/null
```

1. Leia outline.json → output_dir
2. Leia slides.md completo
3. Localize `<!-- SLIDE:N "..." -->`
4. **Edit** o bloco de `---` a `---` com novo conteúdo — mantendo tom clínico e citações reais
5. Re-exporte:
```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/export-pdf.sh \
  "[output_dir]/slides.md" "[output_dir]/slides.pdf"
```

---

## MODO IMAGE

1. Leia outline.json → `image_prompt` da seção N
2. Gere:
```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/generate-image.sh \
  "PROMPT" "section-N.png" "OUTPUT_DIR"
```
3. Atualize referência no slide se necessário
4. Re-exporte PDF

---

## MODO EXPORT

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/export-pdf.sh \
  "[output_dir]/slides.md" "[output_dir]/slides.pdf"
```
