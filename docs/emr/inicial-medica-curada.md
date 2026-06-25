# Avaliação Médica de Entrada (Continuum) — curadoria de encantamento

**Status:** proposta para revisão do Dr. Getúlio · **Data:** 2026-06-23
Consulta de ~40min que gera **radar AGIR parcial de alto impacto**. Itens **pontuáveis, rápidos,
sem depender de exame de sangue e sem equipamento** (bioimpedância → Nutri; testes físicos →
Ed. Física; ambiente de sono → Psico). **Cada item é exclusivo da Inicial** (sem redundância);
os demais templates completam a profundidade. Inclui **Objetivos** (só aqui).

Legenda: ⭐ = escala-widget (score instantâneo). Itens marcam por grupo do score.

## Eixo A — Atividade, Alimentação, Composição
**Composição corporal (antropometria rápida de risco cardiometabólico):**
- IMC (kg/m²)
- Cintura (cm) — homem / mulher
- Razão cintura/altura
- Razão cintura/quadril — homem / mulher
- % Gordura corporal — homem / mulher  *(se houver medida rápida; senão fica na Nutri)*
- Gordura visceral — homem / mulher  *(idem)*

**Alimentação (padrão geral, rápido):**
- Padrão alimentar atual
- Consumo de açúcar
- Refrigerantes e energéticos / ultraprocessados
- Álcool (tipo/quantidade)
- Água (30 ml/kg/dia)

**Movimento (hábito, sem teste físico):**
- Atividade física (frequência e intensidade)
- Exercício físico (frequência e intensidade)

## Eixo G — Gestão Clínica e Metabólica  *(núcleo do médico)*
- **Histórico de doenças**: manter o núcleo atual (~53 itens) de condições/sintomas crônicos
  de alta prevalência (cardiometabólico, tireoide, GI, respiratório, etc.). O Complemento fica
  com o aprofundamento e as condições raras (os outros ~105 do pool).
- **Histórico Familiar de Doenças**: manter os ~9 (HAS, DM, obesidade, dislipidemia, câncer,
  cardiovascular, renal, autoimune, virais) — já vinculados ao pilar Histórico Familiar.
- **Vida Sexual (orgânico/clínico):**
  - Libido/desempenho sexual atual
  - ⭐ IIEF-5 (função erétil, homens) · ⭐ FSFI (função sexual feminina) *(gênero-específico)*
  - Uso recente de hormônios/inibidores de fosfodiesterase; pós-menopausa

## Eixo I — Integração Mente-Corpo  *(grande gerador de "wow")*
- **Vitalidade/energia (fadiga — dor nº1):**
  - Disposição/energia ao longo da vida + disposição atual (trabalho, social, exercício)
  - ⭐ FSS (severidade de fadiga)
- **Humor / ansiedade:**
  - ⭐ PHQ-9 (humor) · ⭐ GAD-7 (ansiedade)
- **Cognição:**
  - Capacidade de foco/concentração + memória percebida
  - *(opcional, se o tempo permitir: ⭐ 5 palavras de Dubois imediato/tardio, ⭐ Span de dígitos)*
- **Estresse:**
  - Fontes de stress percebidas atualmente
  - Sintomas atuais que atribui ao stress
  - Formas/ferramentas de manejo do estresse
- **Social/contexto (determinantes rápidos):**
  - Situação conjugal · familiar · financeira · lazer/hobbies

## Eixo R — Ritmo e Sono  *(parte clínica fica no médico)*
- Qualidade percebida do sono
- ⭐ Escala de Sonolência de Epworth (sonolência diurna)
- Roncos · Apnéias · Insônia (dificuldade de iniciar) · Bruxismo · Sudorese noturna
- *(ambiente/higiene do sono — colchão, telas, luz, temperatura — vão para Psico/profundo)*

## Objetivos  *(só na Inicial)*
- Os 19 itens de Objetivos (metas do cliente) — base do plano e do "antes/depois".

---
## Resumo / tamanho estimado
~**110–140 itens pontuáveis** cobrindo os 4 eixos e as principais dores (fadiga, sono, peso/
composição, cardiometabólico, humor/ansiedade, estresse, libido, cognição) — viável em 40min,
sobretudo com pré-preenchimento do paciente. As escalas-widget dão score imediato e impacto visual.

## Implicações de dedup (sem redundância)
- O que entra na Inicial **sai** dos owners (Complemento/Nutri/Psico/Física) — fica exclusivo.
- Owners ficam com a profundidade: Histórico de doenças raro/aprofundado (Complemento),
  bioimpedância completa (Nutri), testes físicos (Ed. Física), ambiente/higiene do sono e
  cognição profunda (Psico), composição de exercício (Ed. Física).
- Objetivos removidos dos outros 4.

## Próximo passo
Você revisa/poda esta lista (tirar/marcar o que não quer na 1ª consulta). Aí eu mapeio nome→id,
gero o SQL (realocação de vínculos em `anamnesis_template_items`, idempotente) e aplico em **dev**
com relatório antes/depois, depois prod.
