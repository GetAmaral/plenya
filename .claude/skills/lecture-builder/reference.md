# Reference: Marp Templates + DALL-E Prompts

## Marp Slide Templates

### Slide de Capa (Slide 1)
```markdown
---
marp: true
theme: medical
paginate: true
header: 'Título da Aula'
footer: 'Apresentador | Instituição | 2026'
---

<!-- SLIDE:1 "Título" -->
# Título Completo da Aula

**Subtítulo ou tema específico**

---
*Nome do Apresentador*
*Cargo / Especialidade*
*Data | Local*

---
```

### Slide de Objetivos (Slide 2)
```markdown
<!-- SLIDE:2 "Objetivos" -->
## 2. Objetivos de Aprendizagem

Ao final desta aula, você será capaz de:

1. **Identificar** os critérios diagnósticos de...
2. **Interpretar** os exames laboratoriais relevantes
3. **Selecionar** a conduta terapêutica adequada
4. **Orientar** o paciente sobre...

---
```

### Slide de Conteúdo — Texto + Conceito
```markdown
<!-- SLIDE:N "Fisiopatologia" -->
## N. Fisiopatologia

- Mecanismo principal → consequência direta
- Cascata inflamatória: **IL-6 ↑ → PCR ↑ → risco CV ↑**
- Resistência à insulina → hiperinsulinemia compensatória
- Dislipidemia aterogênica: ↑TG, ↓HDL-c

<div class="concept">
**Mecanismo Central:** Resistência à insulina é o ponto comum
entre todos os componentes da síndrome
</div>

![](images/section-N.png)

---
```

### Slide de Tabela — Critérios Diagnósticos
```markdown
<!-- SLIDE:N "Critérios Diagnósticos" -->
## N. Critérios Diagnósticos

| Organização | Critério Principal | Componentes |
|-------------|-------------------|-------------|
| IDF 2006 | Obesidade central obrigatória | + 2 de 4 |
| NCEP ATP III | Sem critério obrigatório | 3 de 5 |
| Harmonizado 2009 | Sem critério obrigatório | 3 de 5 |

**Critérios componentes (valores harmonizados 2009):**

| Componente | Ponto de Corte |
|------------|----------------|
| Cintura abdominal | ≥90 cm (H) / ≥80 cm (M) — Latino-americanos |
| Triglicerídeos | ≥150 mg/dL ou tratamento |
| HDL-c | <40 mg/dL (H) / <50 mg/dL (M) ou tratamento |
| Pressão arterial | ≥130/85 mmHg ou tratamento |
| Glicemia jejum | ≥100 mg/dL ou DM2 tratado |

---
```

### Slide com Alerta Clínico
```markdown
<!-- SLIDE:N "Complicações" -->
## N. Complicações e Alertas

- Risco cardiovascular aumentado em **2-3x**
- Progressão para DM2: **5x maior** que população geral
- Esteatose hepática não-alcoólica presente em **70%**

<div class="alert">
⚠️ **Atenção:** Pacientes com SM e glicemia 100-125 mg/dL têm
alto risco de progressão para DM2 — intervenção imediata obrigatória
</div>

<div class="evidence">
📊 **Evidência:** Metanálise (Mottillo et al., JACC 2010, n=172.573):
SM associada a RR=2,35 para eventos CV fatais (IC95% 2,02-2,73)
</div>

---
```

### Slide de Tratamento — Dupla Coluna
```markdown
<!-- SLIDE:N "Tratamento" -->
## N. Abordagem Terapêutica

<div class="columns">
<div>

**Não Farmacológico (1ª linha)**
- Perda ponderal ≥7-10% do peso
- Dieta mediterrânea ou DASH
- Atividade física: 150-300 min/sem aeróbica
- Cessação tabágica
- Moderação de álcool

</div>
<div>

**Farmacológico (por componente)**
- HAS: preferencialmente IECAs/BRAs
- Dislipidemia: estatinas (↓LDL) + fibratos (↑HDL)
- Glicemia: metformina (prevenção DM2)
- Obesidade: orlistat, semaglutida

</div>
</div>

---
```

### Slide de Referências (último slide)
```markdown
<!-- SLIDE:N "Referências" -->
## N. Referências Principais

1. Alberti KGMM et al. *Harmonizing the Metabolic Syndrome*. Circulation. 2009;120:1640-5. PMID: 19805654

2. Grundy SM et al. *Diagnosis and Management of the Metabolic Syndrome*. Circulation. 2005;112:2735-52. PMID: 16157765

3. Mottillo S et al. *The Metabolic Syndrome and Cardiovascular Risk*. JACC. 2010;56:1113-32. PMID: 20863953

4. Diretriz Brasileira de Síndrome Metabólica. *Arq Bras Cardiol*. 2005;84(Supl I):1-28.

---
*Obrigado!*
*Dúvidas?*
```

---

## DALL-E 3 — Prompts por Tipo de Slide

### Fisiopatologia / Mecanismo
```
Medical educational diagram illustrating [mechanism name] pathophysiology.
Show a clear flowchart with boxes and arrows depicting the biological cascade.
Clinical illustration style, clean white background, professional medical textbook quality.
Labels in Portuguese. High detail, no photograph elements.
```

### Epidemiologia / Prevalência
```
Medical infographic showing global prevalence of [disease name].
World map with color-coded regions showing prevalence rates.
Include bar charts or pie charts with statistics.
Clean white background, professional medical style, data visualization.
Labels in Portuguese.
```

### Critérios Diagnósticos / Classificação
```
Medical educational diagram showing diagnostic criteria for [condition].
Organized flowchart or decision tree format.
Clean boxes with text areas, arrows showing decision paths.
Professional clinical illustration style, white background.
```

### Anatomia / Fisiologia Normal
```
Detailed anatomical medical illustration of [organ/system].
Cross-section view showing relevant structures.
Professional medical textbook style, clean white background.
Color-coded regions for different anatomical components.
High detail suitable for medical education.
```

### Tratamento / Manejo
```
Medical educational treatment algorithm diagram for [condition].
Hierarchical flowchart showing treatment steps and decision points.
Color-coded levels: green for first-line, yellow for second-line, red for rescue.
Professional clinical style, clean white background, organized layout.
```

### Comparação / Tabela Visual
```
Medical educational comparative diagram showing [comparison topic].
Side-by-side visual comparison with clear categories.
Professional infographic style, clean white background.
Use medical iconography, color coding for different options.
```

### Exames / Diagnóstico por Imagem
```
Medical illustration showing [exam type] findings in [condition].
Schematic representation of key diagnostic features.
Professional medical educational style, annotated diagram.
Clean white background, clear labels, clinical accuracy.
```

---

## Marp Frontmatter Completo

```yaml
---
marp: true
theme: medical
paginate: true
size: 16:9
header: 'NOME DA AULA'
footer: 'Apresentador | Instituição | Ano'
style: |
  section {
    font-size: 1.2em;
  }
  h1 {
    font-size: 2em;
    color: #1a3a5c;
  }
  h2 {
    font-size: 1.5em;
    color: #1a3a5c;
    border-bottom: 2px solid #2ecc71;
    padding-bottom: 0.2em;
  }
---
```

---

## Estrutura de Seções Recomendada

### Aula de 30 min (6-8 slides)
1. Capa + Objetivos
2. Epidemiologia
3. Fisiopatologia
4. Diagnóstico
5. Tratamento
6. Casos Clínicos (1-2)
7. Conclusão + Referências

### Aula de 60 min (12-15 slides)
1. Capa
2. Objetivos
3. Introdução / Definição
4. Epidemiologia
5. Fisiopatologia
6. Apresentação Clínica
7. Diagnóstico
8. Exames Complementares
9. Diagnóstico Diferencial
10. Tratamento não-farmacológico
11. Tratamento farmacológico
12. Complicações
13. Casos Clínicos (2-3)
14. Mensagens-chave
15. Referências

### Aula de 90 min (18-22 slides)
Adicionar às de 60 min:
- Epidemiologia brasileira detalhada
- Genética / Fatores de risco
- Populações especiais (gestantes, idosos, pediátrico)
- Guidelines internacionais comparados
- Prevenção / Rastreamento
- Follow-up / Monitoramento
