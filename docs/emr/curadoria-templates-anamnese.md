# Curadoria de conteúdo — Templates de Anamnese

Registro vivo das decisões de quais score items entram em cada template. Catálogo macro e regras
em [plano-templates-anamnese.md](plano-templates-anamnese.md).

## ✅ APLICADO NO DEV (2026-06-06) — seed `seed-anamnese-templates.sql`
13 templates criados (idempotente, UUID fixo `…111111101`–`…113`), placeholders "Teste"/"Medica
Inicial" aposentados. Counts finais: A1 **437** · B1 137 (cross-AGIR: A/G/I/R todos cobertos) ·
B2 300 · A2=B3 204 · B4 251 · A3 33 · B5 136 · B6 85 · B7 135 · B8 77 · B9 141 · B10 87.
Re-rodar: `sed 's/ROLLBACK/COMMIT/' seed-anamnese-templates.sql | psql`. **Prod: pendente**
(método dry-run-em-prod c/ rollback, quando o Dr. validar no builder do dev).

> ⚠️ A revisar com o Dr. no builder: (1) A2/B3 com 204 itens é pesado p/ acompanhamento "mensal leve" —
> talvez enxugar; (2) rastreio Movimento usou "Estratégia macro atual" + "Divisão das atividades" como
> proxy de atividade atual (o item "freq/intensidade" do banco é ambíguo, duplicado por fase de vida).

## Fronteira (confirmada 2026-06-06)
- Universo de anamnese = **538 itens de pergunta** (`lab_test_code IS NULL`) em 11 grupos.
- **Exames (312) e Genética (361) ficam fora** dos templates de anamnese — entram pelo fluxo de
  resultados de laboratório (results-inbox), não como pergunta.
- Vitais/antropometria medidos na consulta estão em **Composição corporal** (é anamnese, entra).

## Legenda de status por (sub)grupo
✅ incluir tudo · �ðŸ”¶ parcial (itens marcados) · ⬜ delegar a outra disciplina/template · ⏳ proposto (aguarda Dr.)

---

## A1 · Avaliação Médica Inicial (Plenya avulsa, Area=Medicina) — conjunto-mãe do track médico

> Status: **✅ FECHADA (2026-06-06).** A1 ampla (avulsa = médico único). Núcleo + Mental integrais +
> os dois rastreios (Alimentação ~13, Movimento ~6) aprovados pelo Dr. Total ≈ **436 itens de anamnese**
> (universo 538 menos Alimentação-fora-do-rastreio ~45 e Movimento-fora-do-rastreio ~57 incl. Testes
> práticos 32 — contagem exata sai no seed).

| Grupo | Subgrupo | Itens | Proposta | Decisão |
|---|---|---:|---|---|
| Objetivos | Objetivos iniciais | 15 | ✅ | |
| Objetivos | Percepção de futuro | 3 | ✅ | |
| Objetivos | Adesão / perfil comportamental | 1 | ✅ | |
| Histórico de doenças | Doenças crônicas | 26 | ✅ | |
| Histórico de doenças | Histórico de saúde | 66 | ✅ | |
| Histórico de doenças | Cirurgias já realizadas | 22 | ✅ | |
| Histórico de doenças | Medicamentos | 33 | ✅ | |
| Histórico de doenças | Hábitos e vícios nocivos | 6 | ✅ | |
| Histórico de doenças | Saúde bucal | 8 | ✅ | |
| Histórico de doenças | Especialistas médicos externos | 11 | ✅ | |
| Histórico de doenças | Equipe multiprofissional externa | 6 | ✅ | |
| Histórico Familiar | Parentes próximos | 11 | ✅ | |
| Histórico Familiar | Parentes distantes | 1 | ✅ | |
| Histórico Familiar | Hábitos/vícios dos parentes | 12 | ✅ | |
| Sono | Histórico | 10 | ✅ (R é médico-led) | |
| Sono | Atual | 39 | ✅ | |
| Vida Sexual | Histórico | 7 | ✅ (eixo hormonal) | |
| Vida Sexual | Atual | 14 | ✅ | |
| Composição corporal | Medidas Objetivas (antropometria/vitais) | 47 | ✅ | |
| Composição corporal | Histórico | 9 | ✅ | |
| Composição corporal | Atual | 3 | ✅ | |
| Cognição | Histórico | 18 | ✅ | |
| Cognição | Atual | 20 | ✅ (médico faz o screen) | |
| Stress | Histórico + Atual | 6 | ✅ | |
| Social | Histórico + Atual | 23 | ✅ | |
| Alimentação | rastreio (subset ~12 de 58) | ~12 | 🔶 proposto (resto → Nutri) | |
| Movimento e atividade física | rastreio (subset ~6 de 31) | ~6 | 🔶 proposto (resto + Testes práticos → Ed. física) | |

### Rastreio Alimentação (proposto p/ A1 — resto fica no template Nutri)
Padrão alimentar atual · Água (30ml/kg/dia) · Líquidos no dia · Álcool (discriminar tipo) · Consumo de
Açúcar · Refrigerantes e energéticos · Consumo de Frutas · Consumo de Verduras e Legumes · Café e chás
(cafeína) · Intolerâncias · Restrições alimentares · Suplementações utilizadas (marcas e doses) · Tem
suplementação prescrita? está usando adequadamente?

### Rastreio Movimento (proposto p/ A1 — resto + Testes práticos fica no template Ed. física)
Estratégia macro atual · Atividade física atual (frequência e intensidade) · Restrições de atividades ·
Lesões relacionadas ao exercício · Cirurgias relacionadas ao exercício · Piores fases de sedentarismo

---

---

## Derivados — proposta (⏳ aguarda redline do Dr.)

Base: Continuum médico = **amplo (= A1)** (decisão do Dr.). Logo B1+B2 particionam a A1; acompanhamento
e reavaliação são subsets dinâmicos. Disciplinas puxam dos mesmos grupos.

### B1 · Avaliação Médica de Entrada (Continuum, 45min — encantamento) — varredura CROSS-AGIR
> Decisão do Dr.: **tocar cada aspecto do AGIR** em 45min (não o núcleo G-pesado). Fatia fina por letra
> + escalas validadas (rápidas, impressionam) + medidas objetivas → **radar parcial dos 4 eixos** já na
> entrada, montando o gancho "normal × ótimo".

**Abertura — Objetivos (transversal):** todos (19) — metas + percepção de futuro + adesão.
**A · Atividade/Alimentação/Suplementação (~8):** Padrão alimentar atual · Álcool · Açúcar · Refrigerantes
e energéticos · Água/Líquidos · Suplementações utilizadas · Estratégia macro de movimento atual ·
Atividade física atual (frequência/intensidade).
**G · Gestão Clínica e Metabólica (~70):** Doenças crônicas (26, checklist) · Medicamentos (33) · Histórico
Familiar › Parentes próximos (11) · Medidas Objetivas core: Peso · Altura · IMC · Cintura · Razão
cintura/altura · % Gordura corporal · Gordura visceral · Massa muscular esquelética · Ângulo de fase.
**I · Integração Mente-Corpo (~13):** PHQ-9 (humor) · GAD-7 (ansiedade) · Epworth (sonolência) · Disposição/
energia · Memória percebida · Foco/concentração · Fontes de stress atuais · Sintomas atribuídos ao stress ·
Manejo do stress · Situação conjugal · Situação familiar · Profissões · Lazer/hobbies.
**R · Ritmo Circadiano e Repouso (~9):** Qualidade percebida do sono · Tempo de sono · Hora de dormir/
acordar · Regularidade · Roncos · Apneias · Insônia · Uso de medicamentos p/ dormir.

≈ **120 itens** (muitos são checklist/escala de marcação rápida). *Tudo o que faz o prospect se sentir visto
nas 4 dimensões e alimenta o pitch.*

### B2 · Complemento da Avaliação Médica (Continuum) = A1 − B1 (pós-fechamento)
Histórico de doenças › Histórico de saúde (66) + Cirurgias (22) + Saúde bucal (8) + Especialistas (11) +
Equipe multiprofissional (6) · Histórico Familiar › distantes (1) + hábitos dos parentes (12) · Sono (49) ·
Vida Sexual (21) · Cognição (38) · Stress (6) · Social (23) · Composição › Histórico (9) + Atual (3) ·
rastreio Alimentação (~13) · rastreio Movimento (~6). ≈ **294 itens.** (B1+B2 ≡ A1.)

### A2 · Acompanhamento Médico (avulsa)  ≡  B3 · Acompanhamento Médico (Continuum) — subset dinâmico
Objetivos (revisão) · Doenças crônicas (controle) · Medicamentos (reconciliação) · Medidas Objetivas
(re-medir) · subgrupos **"Atual"**: Sono-Atual (39) + Cognição-Atual (20) + Stress-Atual (3) + Vida
Sexual-Atual (14) + Social-Atual (15) + rastreio Alimentação + rastreio Movimento (adesão). *Sem
histórico imutável (infância, cirurgias, familiar).*

### B4 · Reavaliação Médica Trimestral (Continuum) — RECALCULAR O ESCORE (decisão do Dr.)
Re-administra **todos os itens dinâmicos que pontuam** (os imutáveis — infância, cirurgias, familiar — não
voltam; exames re-entram pelo results-inbox). Inclui: subgrupos **"Atual"** de Sono/Cognição/Stress/Vida
Sexual/Social + Doenças crônicas (controle) + Medicamentos + Medidas Objetivas (re-medir) + rastreios
Alimentação/Movimento + Objetivos (revisão de metas). Mais largo que o acompanhamento mensal (A2/B3).

### A3 · Revisão de Exames (avulsa) — template MÍNIMO (decisão do Dr.)
Objetivos › queixa/foco atual · Medicamentos (reconciliação). O conteúdo clínico vem do results-inbox,
não da anamnese.

### B5 · Avaliação Nutricional Inicial (Continuum) — Area=Nutricao
Alimentação (58, integral) · Composição corporal (59) · Objetivos (19, metas). ≈ **136 itens.**

### B6 · Acompanhamento Nutricional (Continuum)
Alimentação › Atual (35) + Adesão · Composição corporal › Medidas Objetivas + Atual.

### B7 · Avaliação Psicológica Inicial (Continuum) — Area=Psicologia
Cognição (38) · Stress (6) · Social (23) · Sono (49, sono↔saúde mental) · Objetivos (metas). ≈ **135 itens.**

### B8 · Acompanhamento Psicológico (Continuum)
"Atual" de Cognição (20) + Stress (3) + Social (15) + Sono-Atual (39).

### B9 · Avaliação Física Inicial (Continuum) — Area=Educacao Fisica
Movimento e atividade física (63, integral, **inclui Testes práticos 32**) · Composição corporal (59) ·
Objetivos (metas). ≈ **141 itens.**

### B10 · Acompanhamento (Educação Física) (Continuum)
Movimento › Atual (8) + Testes práticos (32, re-teste) · Composição corporal › Medidas Objetivas.

---

## Sobreposições intencionais (mesmo item em vários templates — suportado por FK)
- **Composição corporal / Medidas Objetivas** → A1, B1, B2, A2/B3, B4, B5, B6, B7?, B9, B10 (antropometria é comum).
- **Objetivos** → todas as iniciais.
- **rastreio Alimentação/Movimento** (A1/médico) vs **Alimentação/Movimento integral** (Nutri/Ed.física): o médico vê o screen, o especialista vê o todo.
- **Sono** → médico (B2) + psico (B7/B8).
