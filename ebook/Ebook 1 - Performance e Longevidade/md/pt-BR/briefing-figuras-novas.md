# Briefing para produção das 9 figuras novas

Arquivo destinado ao gerador de IA (ou designer) que vai produzir as 9 imagens aprovadas. Este documento define o padrão visual comum do livro (derivado das ~30 figuras já prontas) e entrega um prompt detalhado para cada nova figura.

---

## Padrão visual comum (aplicar a TODAS as figuras)

### Moldura e identidade

- **Caixa "FIGURA N" no canto superior esquerdo:** retângulo vermelho-magenta sólido (~`#C6264B`), texto branco em fonte sans-serif bold, caixa-alta ("FIGURA 1", "FIGURA 2", …). Numeração: a figura inteira do livro é renumerada por capítulo (ex.: "FIGURA 1" para a primeira do Cap 6, ainda que seja `Cap06 Fig04`).
- **Proporção:** landscape, aproximadamente 4:3 ou 16:10. Compatível com leitura em EPUB/KDP e impressão 6×9".
- **Margens internas generosas.** Estilo "data journalism" (referência: *The Economist*, *Our World in Data*, *NYT Graphics*). Minimalista, sem sombras, sem bordas grossas, sem ornamentos decorativos.

### Tipografia

- **Título principal da figura** (dentro da imagem, logo abaixo da caixa "FIGURA"): sans-serif bold, preto (`#111`), tamanho ~24–30pt no artwork. Texto afirmativo, curto, 1–2 linhas. Funciona como manchete.
- **Subtítulo / descrição**: sans-serif regular, cinza médio (`#5A5A5A`), ~14–16pt. Uma linha.
- **Rótulos de eixos e dados**: sans-serif regular, cinza escuro (`#2E2E2E`), ~11–13pt.
- **Rodapé com fonte do estudo**: sans-serif regular, cinza claro (`#8A8A8A`), ~9–10pt, itálico em "Fonte:". Formato: *"Fonte: [autor et al., ano, periódico]. [nota complementar se cabível]."*
- **Fontes sugeridas:** Inter (corpo) e Fraunces (se quiser elegância em títulos). Fallback seguro: Helvetica Neue / Arial.

### Paleta de cores

| Uso | Cor | Hex aproximado |
|---|---|---|
| Preto (títulos, eixos) | preto suave | `#111` |
| Cinza escuro (texto corpo) | cinza escuro | `#2E2E2E` |
| Cinza médio (subtítulo, dados secundários) | cinza médio | `#5A5A5A` |
| Cinza claro (rodapé, grid) | cinza claro | `#8A8A8A` / `#E5E5E5` (grid) |
| **Vermelho/magenta de destaque** (FIGURA, dado-chave, alerta) | vermelho-magenta | `#C6264B` |
| Verde de "ótimo" / "protetor" | verde floresta | `#2E7D4F` |
| Amarelo/ocre de "subótimo" / "atenção" | âmbar | `#D89C3D` |
| Azul (linha temporal, dado secundário) | azul escuro | `#2C5F8A` |
| Off-white de fundo | bege muito claro | `#FAF7F2` (se usar fundo) |

**Regra:** no máximo 4 cores semânticas além dos tons de cinza em cada figura. Mais cores = ruído.

### Elementos recorrentes

- **Destaques de pacientes** (Ricardo, Fernanda, André, Marcos, Paulo, Ana, Marina): ponto circular colorido + label com nome + número pequeno em caixa arredondada. Cor: vermelho-magenta para o "protagonista" da figura; cinza para pacientes secundários.
- **Zona ótima / zona subótima**: faixa horizontal colorida (verde / amarelo / vermelho) atrás do gráfico, com texto rotulado dentro da faixa.
- **Setas de movimento**: setas finas vermelhas/verdes indicando evolução (ex.: "de X para Y") com valor antes → depois.
- **Grid**: cinza muito claro (`#E5E5E5`), linhas finas, apenas horizontal em gráficos de barra, apenas onde absolutamente necessário. Zero grid vertical desnecessário.
- **Anotações**: balões curtos fora do gráfico com linha-guia fina, nunca ocupando mais de 3 linhas de texto.
- **Rodapé de fonte**: sempre abaixo da imagem, sempre com referência precisa ao estudo citado.

### Idioma

Todas as figuras em **português do Brasil**. Sem anglicismos desnecessários. Siglas médicas internacionais (HbA1c, ApoB, CAC, PCR, TSH, HDL, LDL) são mantidas em inglês por convenção médica — já listadas em `glossario.yaml > never_translate`.

---

## FIGURAS NOVAS (9 no total)

---

### 1. `Cap06 Fig04.PNG` — Kraft curve do André

**Título principal:** *"O TOTG de André: quando o jejum mente."*

**Subtítulo:** *"Padrão Kraft de hiperinsulinemia compensatória em triatleta de 45 anos com glicemia de jejum 'normal'."*

**Tipo de gráfico:** gráfico de linha dupla (dois eixos Y) ao longo de tempo (0, 30, 60, 90, 120 min).

**Dados a plotar:**

| Tempo (min) | Glicose (mg/dL) | Insulina (µIU/mL) |
|---|---|---|
| 0 | 92 | 8,5 |
| 30 | 148 | 78 |
| 60 | 162 (pico) | 124 (pico) |
| 90 | 154 | 118 |
| 120 | 131 | 89 |

**Eixos:**
- X: Tempo (min), rótulos nos 5 pontos.
- Y-esquerdo: Glicose (mg/dL), escala 80–180.
- Y-direito: Insulina (µIU/mL), escala 0–140.

**Cores e estilo:**
- Linha de glicose: azul escuro (`#2C5F8A`), espessura 2,5pt.
- Linha de insulina: vermelho-magenta (`#C6264B`), espessura 2,5pt, linha tracejada para distinguir.
- Faixa verde horizontal no eixo esquerdo marcando faixa "normal pós-prandial" (≤140 mg/dL) — para mostrar que glicose ficou ligeiramente acima mas não dramática.
- Pico de insulina em 60 min destacado com círculo vermelho vazado + anotação: *"Pico de insulina 124 µIU/mL aos 60 min — Kraft II: hiperinsulinemia compensatória."*
- Anotação no valor de jejum: *"Insulina jejum 8,5 — 'normal' pelo laboratório, mas resposta pós-prandial disfuncional."*

**Rodapé:** *"Fonte: dados reconstruídos do caso-tipo apresentado no Capítulo 6. Padrão de referência: Kraft JR (*Detection of Diabetes Mellitus In Situ*, 2008)."*

---

### 2. `Cap08 Fig02.PNG` — Algoritmo de alopecia androgenética com posicionamento sobre PFS

**Título principal:** *"Finasterida: quando sim, quando não."*

**Subtítulo:** *"O posicionamento clínico do autor — indicação regulatória vs. uso cosmético, à luz da síndrome pós-finasterida (PFS)."*

**Tipo de gráfico:** fluxograma vertical + timeline regulatória abaixo.

**Estrutura do fluxograma (topo → base):**

1. Caixa de topo (cinza escuro, texto branco): **"Paciente adulto com queda de cabelo e/ou sintomas prostáticos"**
2. Pergunta-diamante (laranja-ocre, `#D89C3D`): **"Tem HPB sintomática (IPSS ≥ 8)?"**
3. Ramo SIM (seta verde `#2E7D4F` para a esquerda):
   - Caixa: **"Finasterida 5 mg/dia OU dutasterida 0,5 mg/dia — indicação clínica consistente, décadas de evidência."**
   - Subcaixa menor: *"Consentimento inclui conversa sobre efeitos sexuais possíveis."*
4. Ramo NÃO (seta vermelha `#C6264B` para a direita):
   - Caixa: **"Queixa estritamente cosmética (alopecia androgenética)?"**
   - Caixa final (vermelha, texto branco): **"NÃO prescrever finasterida nem dutasterida. Risco de síndrome pós-finasterida (PFS) desproporcional ao desfecho estético."**
   - Caixa-alternativa abaixo (verde claro): **"Alternativas com evidência: minoxidil tópico 5% ou oral baixa dose; PRP capilar; microneedling; correção de ferritina ≥ 40; otimização de vit. D, B12, zinco, função tireoidiana."**

**Timeline regulatória (abaixo do fluxograma):**
Linha horizontal do tempo, 3 marcos:
- **2011 — FDA:** advertência sobre depressão em bula
- **2022 — FDA:** atualização da bula alertando sobre **ideação suicida**
- **2025 — EMA:** reconhecimento formal da **síndrome pós-finasterida** (disfunção sexual persistente + sintomas neuropsiquiátricos)

Cada marco em caixa pequena cinza com a data destacada em negrito.

**Rodapé:** *"Fonte: FDA Drug Safety Communications (2011, 2022); EMA PRAC Recommendation (2025). Posicionamento clínico do autor expresso no Capítulo 8."*

---

### 3. `Cap09 Fig01.PNG` — Eixo Cardio-Reno-Metabólico

**Título principal:** *"Três sistemas, uma medicina só."*

**Subtítulo:** *"Por que três classes de medicamentos deixaram de ser 'remédio de diabético' nos últimos cinco anos."*

**Tipo de gráfico:** diagrama de Venn com três círculos sobrepostos (cardiovascular, renal, metabólico) + drogas posicionadas na interseção central.

**Estrutura:**

Três círculos grandes, translúcidos, sobrepostos no centro formando uma trinca:
- **Coração** (círculo superior esquerdo, azul-escuro translúcido): ícone de coração estilizado + rótulo "Insuficiência cardíaca · Infarto · AVC".
- **Rim** (círculo superior direito, verde floresta translúcido): ícone de rim estilizado + rótulo "DRC · Proteinúria · Progressão renal".
- **Metabolismo** (círculo inferior central, âmbar translúcido): ícone de gordura visceral/fígado + rótulo "Obesidade · Esteatose · Diabetes tipo 2".

**Na interseção central dos 3 círculos:** caixa branca com fundo sólido, borda vermelho-magenta, listando as 3 classes:

1. **SGLT2** (dapagliflozina / empagliflozina)
2. **Finerenona** (Firialta)
3. **GLP-1 / GIP-GLP-1** (semaglutida / tirzepatida)

Cada classe com 1–2 ensaios-referência em texto pequeno embaixo:
- SGLT2: *DAPA-HF, EMPA-KIDNEY, DELIVER*
- Finerenona: *FIDELIO-DKD, FIGARO-DKD, FIND-CKD*
- GLP-1: *SELECT (2023), retatrutida fase 2*

**Anotação lateral** (balão com linha-guia apontando para o centro): *"Antigamente: 'remédio de diabético'. Hoje: ferramenta de proteção cardiorrenal e metabólica em paciente sem diabetes."*

**Rodapé:** *"Fonte: ensaios DAPA-HF (McMurray, NEJM 2019), EMPA-KIDNEY (Herrington, NEJM 2023), DELIVER (Solomon, NEJM 2022), FIDELIO-DKD (Bakris, NEJM 2020), FIGARO-DKD (Pitt, NEJM 2021), FIND-CKD (2025), SELECT (Lincoff, NEJM 2023). Síntese clínica do Capítulo 9."*

---

### 4. `Cap10 Fig03.PNG` — Hipótese da janela para TRH

**Título principal:** *"A janela dura cerca de 10 anos."*

**Subtítulo:** *"Começar a reposição hormonal dentro dela protege. Começar fora pode fazer o oposto."*

**Tipo de gráfico:** timeline horizontal com faixas coloridas + pontos de ensaios-marco.

**Estrutura:**

Eixo horizontal em **"anos desde a última menstruação"**, de 0 a 20 anos. Marcações em 0, 5, 10, 15, 20.

Três faixas horizontais sobrepostas:

1. **Faixa VERDE (`#2E7D4F`)** de 0 a ~10 anos — rotulada **"JANELA DE PROTEÇÃO"** em caixa-alta: *"Estradiol transdérmico reduz progressão de aterosclerose, protege osso, preserva função cognitiva."*
2. **Faixa AMARELA/OCRE** de 10 a 15 anos — rotulada **"ZONA DE DECISÃO INDIVIDUAL"**: *"Evidência menos clara. Decisão caso a caso com estratificação de risco."*
3. **Faixa VERMELHA/MAGENTA (`#C6264B`)** de 15 a 20+ anos — rotulada **"FORA DA JANELA"**: *"Aterosclerose estabelecida torna o estradiol potencialmente deletério. Considerar alternativas não-hormonais."*

**Marcos de ensaios (pontos circulares com caixa de texto):**

- **WHI (2002)** em ~12–13 anos da linha do tempo, caixa cinza: *"Média de 63 anos, 10+ anos pós-menopausa. Assustou uma geração inteira."*
- **ELITE (2016)** em ~4 anos (grupo precoce) e ~12 anos (grupo tardio), em cores contrastantes: verde para o precoce (*"Reduziu progressão de aterosclerose"*), vermelho para o tardio (*"Não reduziu; em alguns recortes, piorou"*).
- **KEEPS (2024 follow-up)** em ~5 anos, caixa verde: *"Confirmou a hipótese da janela em seguimento longo."*

**Destaque de paciente:**
- **Fernanda (44 anos)** posicionada dentro da zona verde, em ~2 anos da última menstruação, com caixa de destaque: *"Fernanda, 44 — FSH 38, estradiol 28 pg/mL. Na janela."*

**Rodapé:** *"Fonte: Hodis et al., ELITE, NEJM 2016; Harman et al., KEEPS primary, Annals 2014; Miller et al., KEEPS continuation, PLOS Medicine 2024; Rossouw et al., WHI primary, JAMA 2002."*

---

### 5. `Cap11 Fig02.PNG` — Polimorfismos que mudam a conduta

**Título principal:** *"Oito genes. Oito decisões clínicas diferentes."*

**Subtítulo:** *"Quando o teste genético muda o plano — e quando apenas adiciona ansiedade."*

**Tipo de gráfico:** grid de 8 cards organizados em 4 colunas × 2 linhas (ou 2 × 4), cada card representando um polimorfismo com estrutura padronizada.

**Estrutura de cada card (modelo):**

```
┌─────────────────────────┐
│  GENE (em bold)         │
│  Variante               │
├─────────────────────────┤
│  O que significa        │
│  (ícone + frase curta)  │
├─────────────────────────┤
│  AÇÃO CLÍNICA           │
│  (destaque em vermelho  │
│  ou verde conforme o    │
│  caso)                  │
└─────────────────────────┘
```

**Os 8 cards (em ordem):**

1. **MTHFR** (C677T, A1298C) — *"Metaboliza folato mal."* → **"Suplementar L-metilfolato + metilcobalamina. Monitorar homocisteína."**
2. **APOE4** (ε4/ε4) — *"Risco aumentado de Alzheimer tardio."* → **"Prevenção multidomínio intensiva (FINGER). Benefício maior em portadores."**
3. **CYP1A2** (rs762551) — *"Metabolizador lento ou rápido da cafeína."* → **"Lento (CC): parar cafeína antes das 12h. Rápido (AA): tolera até 16h."**
4. **FADS1/2** — *"Conversão ruim de ALA → EPA/DHA."* → **"Peixe ou suplemento direto. Linhaça não resolve."**
5. **FTO** (rs9939609) — *"Maior apetite e preferência calórica."* → **"Exercício de intensidade moderada-alta responde acima da média. Torna treino inegociável."**
6. **ALDH2** (rs671) — *"Flush + risco de câncer esofágico com álcool."* → **"Abstinência ou consumo muito ocasional, não 'com moderação'."**
7. **VDR FokI** (rs2228570) — *"Menor responsividade à vitamina D."* → **"Dose maior (7.000–10.000 UI/dia) sob monitoramento até atingir 40 ng/mL."**
8. **ESR1 / COL1A1** — *"Maior perda óssea na pós-menopausa."* → **"Pesar no cálculo da decisão sobre TRH em mulher com história familiar de osteoporose."**

**Paleta de ação clínica:**
- Verde (`#2E7D4F`) se a ação é preventiva/suplementar simples
- Âmbar (`#D89C3D`) se exige decisão clínica complexa
- Vermelho-magenta (`#C6264B`) se é "evitar" / "não consumir"

**Rodapé:** *"Fonte: variantes de impacto clínico estabelecido referenciadas em: MTHFR (Frosst 1995); APOE (Fortea, Nature Medicine 2024); FTO (Frayling 2007); ALDH2 (Brooks 2009); CYP1A2 (Cornelis 2006). Síntese do Capítulo 11."*

---

### 6. `Cap12 Fig02.PNG` — Ana, 6 meses depois

**Título principal:** *"Quando o pilar psicológico entra, a biologia responde."*

**Subtítulo:** *"Ana, 44 anos: 18 meses de otimização biológica sem mover dois marcadores. Seis meses de trabalho na mente e os quatro saíram do lugar."*

**Tipo de gráfico:** painel em 2 linhas horizontais com 4 marcadores cada linha, setas de movimento (antes → depois).

**Linha 1 — Marcadores biológicos:**

| Marcador | Antes | Depois | Meta |
|---|---|---|---|
| hs-CRP (mg/L) | 1,8 | 0,7 | < 1,0 |
| Cortisol matinal (µg/dL) | 22 | 14 | 10–18 |

**Linha 2 — Escalas psicológicas:**

| Instrumento | Antes | Depois | Corte-alerta |
|---|---|---|---|
| PHQ-9 (depressão) | 14 | 6 | ≥ 10 |
| GAD-7 (ansiedade) | 16 | 5 | ≥ 10 |

**Estilo visual:**
- Cada marcador em "dot plot" horizontal: ponto vermelho no valor "antes" + ponto verde no valor "depois" + seta fina conectando.
- Faixa verde atrás indicando zona-meta.
- Rótulos: "ANTES" (cinza), "DEPOIS" (verde).
- Linha divisória horizontal fina entre as duas linhas, com rótulos: **"BIOLÓGICO"** (cinza escuro) e **"PSICOLÓGICO"** (azul escuro).

**Anotação à direita** (balão grande): *"Os pilares biológicos não mudaram no período. O que mudou foi TCC + MBSR + antidepressivo em dose baixa por 6 meses. Idade epigenética desacelerou 2 anos."*

**Rodapé:** *"Fonte: caso-tipo do Capítulo 12. PHQ-9 (Kroenke et al., JGIM 2001); GAD-7 (Spitzer et al., Arch Intern Med 2006)."*

---

### 7. `Cap12 Fig03.PNG` — Instrumentos de triagem psicológica

**Título principal:** *"Cinco instrumentos, cinco minutos, cinco perguntas diferentes."*

**Subtítulo:** *"O que o consultório não-psiquiátrico precisa ter à mão para identificar quando a camada psicológica virou prioridade do plano."*

**Tipo de gráfico:** tabela visual com 5 linhas (um instrumento por linha), 4 colunas (instrumento, o que mede, pontuação/cortes, quando vira prioridade).

**Linhas:**

| Instrumento | O que mede | Pontuação e cortes | Vira prioridade se… |
|---|---|---|---|
| **PHQ-9** | Depressão | 0–27. 5 leve · 10 moderada · 15 mod-severa · 20 severa | ≥ 10; qualquer resposta positiva na pergunta 9 (ideação) — **conversa no mesmo dia** |
| **GAD-7** | Ansiedade generalizada | 0–21. 5 leve · 10 moderada · 15 severa | ≥ 10 |
| **AUDIT** | Consumo problemático de álcool | 0–40. 8 risco · 16 uso nocivo/dependência | ≥ 8 — sobretudo em paciente que diz "só bebo socialmente" |
| **PCL-5** | Transtorno de estresse pós-traumático | 0–80. 33 = encaminhar | ≥ 33 — aplicar se a pergunta de trauma na anamnese foi positiva |
| **UCLA-3** | Solidão percebida | 0–9 (3 itens de 0–3) | ≥ 6 — mesmo em paciente que diz ter "muitos amigos" |

**Estilo visual:**
- Cada linha como uma tira horizontal com sua própria cor lateral (barra vertical à esquerda):
  - PHQ-9: cinza escuro
  - GAD-7: azul escuro
  - AUDIT: âmbar
  - PCL-5: vermelho-magenta
  - UCLA-3: verde floresta
- Cortes numéricos em caixas quadradas pequenas dentro de cada linha.
- A pergunta 9 do PHQ-9 destacada em caixa vermelha separada com texto branco: **"Qualquer resposta positiva = conversa direta agora, antes do paciente sair."**

**Box lateral (callout):**

> **"Quando a camada psicológica vira prioridade do plano:"**
> 1. PCR > 1,5 persistente apesar de pilares biológicos otimizados.
> 2. Pontuação ≥ corte em qualquer escala acima.
> 3. Ideação suicida ou trauma reportados em qualquer intensidade.

**Rodapé:** *"Fontes: PHQ-9 (Kroenke et al., JGIM 2001); GAD-7 (Spitzer et al., Arch Intern Med 2006); AUDIT (Saunders et al., Addiction 1993); PCL-5 (Weathers et al., NCPTSD 2013); UCLA-3 (Hughes et al., Research on Aging 2004). Todos validados em português brasileiro."*

---

### 8. `Cap13 Fig02.PNG` — Ricardo em três tempos

**Título principal:** *"Os exames estão bons. Eu não estou."*

**Subtítulo:** *"Ricardo, 18 meses após o infarto: painel biológico em ordem, dimensão relacional em colapso. O trabalho em conexão, propósito e sentido restaurou os dois lados."*

**Tipo de gráfico:** timeline horizontal com 3 momentos e 5 marcadores não-bioquímicos + 2 bioquímicos de controle.

**Eixo X (tempo):** 3 colunas igualmente espaçadas — **T+18m**, **T+21m**, **T+30m** (tempo pós-infarto).

**Linhas de marcadores (de cima para baixo):**

| Marcador | T+18m | T+21m | T+30m |
|---|---|---|---|
| **Vida sexual com Marina** | vermelho: "8 meses sem" | âmbar: "retomada incipiente + tadalafila 5 mg/dia" | verde: "plena, qualidade > pré-IAM" |
| **Amigos próximos (nº)** | 0 | 1 | 3 |
| **Celular no quarto** | sim, checagem 3h–5h | retirado | 10 meses sem |
| **Propósito profissional** | esvaziado | mentoria voluntária iniciada | ritual semanal firme |
| **Ritual laico com Marina** | — | caminhada silenciosa dominical | 52 domingos/ano |
| **PCR ultrassensível (mg/L)** | 1,4 | 1,0 | **0,6** |
| **Cortisol matinal (µg/dL)** | 19 | 17 | **15** |

**Estilo visual:**
- Primeiras 5 linhas: ícones coloridos (vermelho → âmbar → verde) + texto curto.
- 2 linhas inferiores separadas por linha divisória cinza: dot plot com valores numéricos, setas.
- Cabeçalho de cada coluna com subtítulo curto em caixa cinza:
  - T+18m: "Colapso silencioso"
  - T+21m: "Intervenção multi-frente"
  - T+30m: "Inteireza restaurada"

**Anotação no final** (fora do gráfico, embaixo): *"'Doutor, eu achei que ia sair dessa história mais frágil. Eu acho que saí mais inteiro.' — Ricardo, T+30m."*

**Rodapé:** *"Fonte: caso-tipo do Capítulo 13. Princeton Consensus III (Nehra et al., Mayo Clinic Proc 2012); AHA Sexual Activity and CVD (Levine et al., Circulation 2012)."*

---

### 9. `Cap13 Fig03.PNG` — Diagrama de Ikigai

**Título principal:** *"Ikigai: o que te faz levantar da cama."*

**Subtítulo:** *"Quatro círculos que se cruzam num ponto. O estudo Ohsaki (2008) associou a resposta afirmativa à pergunta a 30–50% menos mortalidade em 7 anos."*

**Tipo de gráfico:** diagrama clássico de Ikigai — 4 círculos sobrepostos em mandala, com interseções nomeadas.

**Estrutura (canônica):**

Quatro círculos translúcidos dispostos em cruz:

- **Topo — "O que você ama"** (cor: vermelho-magenta translúcido)
- **Direita — "Pelo que pode ser pago"** (verde floresta translúcido)
- **Base — "O que o mundo precisa"** (âmbar translúcido)
- **Esquerda — "O que faz bem"** (azul escuro translúcido)

**Interseções nomeadas:**

- Ama + Faz bem = **PAIXÃO**
- Faz bem + Pago = **PROFISSÃO**
- Pago + Precisa = **VOCAÇÃO**
- Precisa + Ama = **MISSÃO**

**Centro (interseção dos 4):** caixa destacada com fundo branco e borda vermelho-magenta — **"IKIGAI"** em caixa-alta, tipografia maior + subtítulo: *"A razão de acordar."*

**Balão lateral (à direita do diagrama):**

> **Pergunta clínica:**
> *"Qual é o motivo que te faz levantar da cama de manhã?"*
>
> Se não houver resposta clara em um adulto pós-40 anos, há uma camada do plano que está faltando.

**Rodapé:** *"Fonte: conceito japonês tradicional. Evidência: Sone et al. ('Sense of Life Worth Living and Mortality in Japan: Ohsaki Study'), Psychosomatic Medicine 2008, 70(6):709–715. Coorte de 43.391 adultos, 7 anos de seguimento."*

---

## Checklist antes de aprovar cada figura

- [ ] Caixa "FIGURA N" no canto superior esquerdo, vermelho-magenta, texto branco.
- [ ] Título principal em 1–2 linhas, sans-serif bold preto.
- [ ] Subtítulo em cinza médio, 1 linha.
- [ ] Máximo de 4 cores semânticas além dos cinzas.
- [ ] Destaque vermelho-magenta apenas para o dado mais importante.
- [ ] Nenhuma borda grossa, sombra ou ornamento.
- [ ] Rodapé com fonte/referência do estudo.
- [ ] Proporção landscape (~4:3 ou 16:10).
- [ ] Todo texto em português do Brasil, exceto siglas médicas internacionais.
- [ ] Nome do arquivo exato (ex.: `Cap09 Fig01.PNG`, com espaço antes de "Fig").
- [ ] Tipografia: Inter (corpo) e Fraunces (se desejar) — fallback Helvetica/Arial.

---

*Briefing consolidado em 2026-04-21, pronto para execução. As 9 figuras cobrem as lacunas visuais do livro em 18 capítulos e completam os pilares A (Cap 8), G (Caps 9, 10, 11), I (Caps 12, 13). Nenhuma figura proposta para o Cap 17 (Manifesto) ou Cap 18 (Referências) — conforme decisão editorial.*
