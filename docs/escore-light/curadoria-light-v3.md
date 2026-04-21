# Escore Plenya Light — Curadoria v3

**Constraints aplicados:**
1. Apenas items com `levels` definidos (510 candidatos; 154 sem levels descartados — exigem julgamento profissional).
2. Cross-reference com **todas as dores articuladas no site** (8 retratos do Continuum + Quiz Diagnóstico + 6 artigos do blog).
3. Pares M/F contam como 1 item efetivo (UX filtra automaticamente por gênero).
4. Distribuição AGIR proporcional + painel laboratorial mínimo viável da medicina preventiva moderna.
5. Items "histórico" (passado) só quando essenciais — Light é instrumento de retrato atual; histórico fica no EMR.

---

## Matriz dor → item (cobertura do site)

### 1. Retrato "Faz tudo certo e não vê resposta" + blog *cansaço aos 45* + blog *quando tudo normal não basta*
**Dor:** treina/come/dorme bem mas disposição não volta, peso/marcadores não cedem.
**Items:** Cintura, Razão cintura/altura, % gordura M/F, IMC, Estratégia macro atual, Adesão nutricional, Tempo de sono, Qualidade percebida do sono, Ferritina, Vit D, B12 (Met malônico), TSH, T3 livre, Anti-TPO, Testosterona Livre, Cortisol basal, HOMA-IR.

### 2. Retrato "Recebeu o susto" + blog *checkup coração*
**Dor:** exame que assustou, diagnóstico de família.
**Items:** Doença CV pessoal, Insuficiência cardíaca, Diabetes, Anti-hipertensivos, ApoB, ApoB/ApoA1, Lipoproteína A, PCR-us, Colesterol não-HDL, HbA1c, TC escore de cálcio.

### 3. Retrato "Carrega mais do que si mesmo"
**Dor:** sócios, equipe, filhos, pais dependem de você inteiro.
**Items:** Situação financeira, Situação familiar, Tempo de sono, PHQ-9, GAD-7, Fontes de stress percebidas.

### 4. Retrato "Transição que ninguém nomeia"
**Dor:** sono/humor/corpo mudando, resposta clínica é "é da idade".
**Items:** Estradiol pós-meno (sem TRH) / Estradiol M, Testosterona Total/Livre, DHEA-S (faixa etária), TSH/T3/T4, Densitometria T-score, Depressão histórico, Qualidade percebida do sono.

### 5. Retrato "Profissionais isolados" + "Mede tudo não decide"
**Dor:** wearables/exames sem síntese.
**Items:** Cobertura por agregação (o próprio Escore = síntese; sem item específico).

### 6. Retrato "Sinais pequenos que ninguém soma" + blog *sono que não recupera*
**Dor:** viroses, intestino, pele, cintura sem mudar peso, sono não recupera.
**Items:** Cintura, Razão cintura/altura, Pescoço M/F (proxy apneia), Tempo de sono, Qualidade percebida, Uso BiPAP/CPAP, Capacidade memória atual, Foco/concentração atual, PCR-us, Ferritina.

### 7. Retrato "Inteiro para quem ainda vai chegar"
**Dor:** filho pequeno aos 45, projetos longos.
**Items:** já coberto pelos painéis CV + metabólico + sono + composição.

### 8. Quiz Diagnóstico (pergunta sobre sono)
**Dor:** acorda sem energia / não dorme suficiente / sono fragmentado/ronco.
**Items:** Tempo de sono, Hora de dormir, Hora de acordar, Regularidade no acordar, Qualidade percebida, Uso BiPAP/CPAP, Pescoço M/F.

### 9. Blog *treinar muito e envelhecer errado*
**Dor:** volume sem zona 2/força/recuperação, magro fora sem massa magra, joelhos doem.
**Items:** Estratégia macro atual, Divisão das atividades, Lesões relacionadas ao exercício, ASMI M/F, % gordura M/F, Panturrilha M/F (sarcopenia), Ângulo de fase M/F, PCR-us, Ferritina, Testosterona Livre.

### 10. Blog *sono energia longevidade*
**Dor:** <7h, eficiência, latência, luz matinal/noturna.
**Items:** Tempo de sono, Hora de dormir, Hora de acordar, Regularidade no acordar, Luminosidade natural.

---

## Lista final — 60 items efetivos (~74 registros considerando pares M/F)

**Distribuição AGIR:**
- **A** (Atividade física + Alimentação + Composição): 18 efetivos
- **G** (Gestão clínica + metabólica): 27 efetivos
- **I** (Integração mente-corpo): 9 efetivos
- **R** (Ritmo circadiano + Repouso): 6 efetivos

---

### A — Atividade física, Alimentação, Composição (18 efetivos)

#### Atividade física (3)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 1 | Estratégia macro atual | 14 | `c77cedd3-2800-7bb3-9e74-b76061206583` |
| 2 | Divisão das atividades | 12 | `c77cedd3-2800-7d37-8e88-8ff5483e6ced` |
| 3 | Lesões relacionadas ao exercício | 10 | `c77cedd3-2800-7de9-ad64-b75ed1037b5e` |

#### Alimentação (3)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 4 | Adesão ao tratamento nutricional | 12 | `c77cedd3-2800-70ef-abfd-b3cdd23b2ced` |
| 5 | Divisão e Horários das refeições | 12 | `c77cedd3-2800-73cc-8694-dae1f7d682df` |
| 6 | Onde e como come | 12 | `019bf31d-2ef0-7285-bc84-86d015569d31` |

#### Composição corporal — medidas objetivas (12 efetivos / 18 registros)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 7 | IMC (kg/m²) | 18 | `019bf31d-2ef0-7bb8-8305-e0e0285aeb80` |
| 8 | Razão cintura/altura | 22 | `019bf31d-2ef0-7171-a061-3bcc500957f7` |
| 9a | Cintura (cm) — homem | 22 | `c77cedd3-2800-7aac-985c-c075f020e9e0` |
| 9b | Cintura (cm) — mulher | 20 | `c77cedd3-2800-7a74-99ad-56ca4b6dddc1` |
| 10a | Razão cintura/quadril — homem | 18 | `019bf31d-2ef0-7903-ac22-1d8b5b47179a` |
| 10b | Razão cintura/quadril — mulher | 18 | `019bf31d-2ef0-7b9c-846e-edd6a3c8277f` |
| 11a | % Gordura corporal — homem | 15 | `019bf31d-2ef0-78d5-bb48-f5e21ba33ab4` |
| 11b | % Gordura corporal — mulher | 18 | `c77cedd3-2800-706c-acc0-f6b5da2aa203` |
| 12a | Gordura visceral (cm²) — homem | 22 | `019bf31d-2ef0-7a29-b825-a84bc8f38f3e` |
| 12b | Gordura visceral (cm²) — mulher | 18 | `019bf31d-2ef0-7126-a8b7-ac530f86aa68` |
| 13a | ASMI (kg/m²) — homem | 22 | `c77cedd3-2800-7338-9d58-9d43abcd81a7` |
| 13b | ASMI (kg/m²) — mulher | 22 | `019bf31d-2ef0-7a69-9963-ae68d65a713d` |
| 14a | Pescoço (cm) — homem (proxy apneia) | 12 | `019bf31d-2ef0-767a-a87e-b4587192005b` |
| 14b | Pescoço (cm) — mulher (proxy apneia) | 8 | `019bf31d-2ef0-7f0a-8188-216affee192f` |
| 15a | Panturrilha (cm) — homem (sarcopenia) | 14 | `019bf31d-2ef0-7587-be08-9963c0520bf0` |
| 15b | Panturrilha (cm) — mulher (sarcopenia) | 12 | `019bf31d-2ef0-7936-b83b-284a0f8c84eb` |
| 16a | Ângulo de fase (°) — homem | 14 | `019bf31d-2ef0-7cbd-9813-3635f0bee1d5` |
| 16b | Ângulo de fase (°) — mulher | 14 | `019bf31d-2ef0-7d9f-9924-47e7a3b5830b` |
| 17 | BRI | 12 | `019cbe90-f927-74d1-be7b-3eb87777c2e8` |
| 18 | Tratamentos em uso para modificar composição | 14 | `019bf31d-2ef0-7cff-9970-7c437076adc7` |

---

### G — Gestão clínica e metabólica (27 efetivos)

#### Histórico de doenças e medicamentos (8)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 19 | Doença cardiovascular (IAM, AVC, revasc) | 35 | (consultar SQL final) |
| 20 | Insuficiência cardíaca | 35 | (consultar SQL final) |
| 21 | Diabetes mellitus | 25 | (consultar SQL final) |
| 22 | Asma | 12 | (consultar SQL final) |
| 23 | Anti-hipertensivos | 25 | (consultar SQL final) |
| 24 | Análogos de GLP-1 / Agonistas GIP | 22 | (consultar SQL final) |
| 25 | Antidepressivos (ISRS, ISRN, etc) | 12 | (consultar SQL final) |
| 26 | Inibidores de bomba de prótons (proxy refluxo) | 8 | (consultar SQL final) |

#### Lab cardiovascular (5) — endereça blog *checkup coração*
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 27 | Apolipoproteína B | 22 | (verificar) |
| 28 | ApoB / ApoA1 | 22 | (verificar) |
| 29 | Lipoproteína A | 18 | (verificar) |
| 30 | PCR ultrassensível | 18 | (verificar) |
| 31 | Colesterol não-HDL | 18 | (verificar) |

#### Lab metabólico (3) — endereça blog *quando tudo normal não basta*
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 32 | Hemoglobina glicada | 22 | (verificar) |
| 33 | HOMA-IR | 18 | (verificar) |
| 34 | Insulina jejum (INSULINA 0 MIN) | 15 | (verificar) |

#### Lab nutricional (4) — endereça blog *cansaço aos 45*
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 35 | 25-hidroxivitamina D | 18 | (verificar) |
| 36 | Homocisteína | 15 | (verificar) |
| 37 | Ácido Metilmalônico (B12 funcional) | 12 | (verificar) |
| 38a | Ferritina — Homens | 14 | (verificar) |
| 38b | Ferritina — Mulheres pré-menopausa | 15 | (verificar) |
| 38c | Ferritina — Mulheres pós-menopausa | 14 | (verificar) |

#### Lab hormonal (5)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 39 | TSH | 18 | (verificar) |
| 40 | T3 Livre | 14 | (verificar) |
| 41 | Anti-TPO | 12 | (verificar) |
| 42a | Testosterona Livre — Homens | 16 | (verificar) |
| 42b | Testosterona Livre — Mulheres pré-meno | 14 | (verificar) |
| 42c | Testosterona Livre — Mulheres pós-meno | 12 | (verificar) |
| 43 | Estradiol Mulheres pós-meno (sem TRH) | 18 | (verificar — F only) |
| 44 | Cortisol plasmático basal | 14 | (verificar) |

#### Imagem (2)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 45 | TC coração — escore de cálcio coronariano | 32 | (verificar) |
| 46 | Densitometria — T-score Colo Femoral | 18 | (verificar — relevante transição) |

---

### I — Integração mente-corpo (9 efetivos)

#### Cognição/humor (5)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 47 | PHQ-9 (humor) | 18 | (verificar) |
| 48 | GAD-7 (ansiedade) | 14 | (verificar) |
| 49 | Capacidade da memória percebida | 12 | `c77cedd3-2800-7bec-942e-c1eed243a6a0` |
| 50 | Capacidade de foco/concentração/aprendizado | 12 | (verificar id no MD) |
| 51 | Socialização atual (últimos 6 meses) | 14 | `019c54fc-edfb-73d7-8525-3e326e91a976` |

#### Estresse (1)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 52 | Fontes de stress percebidas atualmente | 14 | `c77cedd3-2800-7360-bf3c-5b4f28c660ef` |

#### Hábitos críticos (2)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 53 | Tabaco | 35 | (verificar) |
| 54 | Álcool | 25 | (verificar) |

#### Vida sexual (1)
| # | Item | Pontos | UUID |
|---|------|--------|------|
| 55 | ASEX (Arizona Sexual Experience Scale) | 12 | `c77cedd3-2800-7d12-9f9e-5dd3fc1e6df0` |

---

### R — Ritmo circadiano e repouso (6 efetivos)

| # | Item | Pontos | UUID |
|---|------|--------|------|
| 56 | Tempo de sono | 20 | `019c53a6-f1a3-7704-9868-354859c750cd` |
| 57 | Hora de dormir | 18 | `c77cedd3-2800-74fb-9aca-432654d7e0d8` |
| 58 | Qualidade percebida do sono | 14 | `c77cedd3-2800-735f-bf28-c5d07d7d7092` |
| 59 | Regularidade no acordar | 12 | `019c53a6-72b5-7718-afac-c14a539f7bca` |
| 60 | Uso de BiPAP/CPAP | 22 | `019c28f0-05c1-7c0f-92b5-9cc9adac6eb3` |
| (*) | Luminosidade Natural (atende blog *sono energia longevidade*) | 12 | `c77cedd3-2800-769d-bd80-42730c51e164` |

> **Nota:** "Luminosidade Natural" hoje vive no grupo "Social", não em "Sono". Sugiro incluir mesmo assim — é central no blog *sono energia longevidade* e na rotina circadiana. Pode contar como item 61, ou substituir um item de menor relevância (ex: ASEX para 55).

---

## Cobertura final das dores

| Dor do site | Cobertura no Light |
|-------------|--------------------|
| Faz tudo certo, sem resposta | ✅ Composição (7-18) + Painel hormonal (39-44) + Nutricionais (35-38) + Sono (56-60) |
| Recebeu o susto | ✅ CV histórico (19-23) + Lab CV (27-31) + TC cálcio (45) + HbA1c (32) |
| Carrega mais do que si mesmo | ✅ Social (item Social → não há ainda; cobrir via PHQ-9/GAD-7/Stress 47-48-52) |
| Transição que ninguém nomeia | ✅ Estradiol/Testosterona/DHEA + TSH/T3 + Densitometria + Depressão (sem item específico no v3 — coberto por PHQ-9) |
| Profissionais isolados / Mede tudo | ✅ O Escore agregado é a resposta |
| Sinais pequenos que ninguém soma | ✅ Cintura (8-9) + Pescoço (14) + Sono (56-60) + Cognição (49-50) + PCR (30) + Ferritina (38) |
| Inteiro para quem vai chegar | ✅ Painel completo |
| Quiz: sono | ✅ Items 56-60 + Pescoço 14 |
| Treinar errado | ✅ Atividade (1-3) + ASMI (13) + % gordura (11) + Panturrilha (15) + PCR (30) + Ferritina (38) + Testosterona (42) |
| Sono não recupera / Sono energia longevidade | ✅ Items 56-60 + Pescoço (14) + Luminosidade (61) |
| Cansaço aos 45 | ✅ Ferritina + B12 + Vit D + TSH/T3 + Testo Livre + Cortisol + HOMA-IR + Apneia (Pescoço + CPAP) |
| Tudo normal não basta | ✅ ApoB + Lp(a) + PCR + Vit D + Ferritina + T3 + HbA1c + HOMA + Estradiol fora de norma |
| Checkup coração | ✅ ApoB + Lp(a) + PCR + Não-HDL + HbA1c + TC cálcio + Anti-hipertensivos + Doença CV |

---

## Próximos passos

1. **Você revisa esta lista** — adicionar/remover items, ajustar `LightOrder`.
2. Eu **completo os UUIDs faltantes** (marcados "verificar") consultando o MD compilado.
3. Eu **escrevo o `LightQuestion`** (texto reformulado em registro adulto descritivo, sem "você sente?") para cada item.
4. Gero o **SQL UPDATE** consolidado.
5. Rodo `pnpm sync:score-light` + commit do JSON gerado.

**Total: 60-61 items efetivos / ~74 registros no banco (pares M/F + faixas de Ferritina/Testo/DHEA).**

A relação tempo estimado de preenchimento: ~10-12 minutos para perguntas demográficas + autorrelato (atividade/alimentação/sono/cognição/sexual/estresse), com items de exames marcados como "se você tem o resultado em mãos" (preenchimento opcional, não contabilizam se vazio — comportamento já existente do EMR).
