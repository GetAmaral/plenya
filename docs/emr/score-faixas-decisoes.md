# Decisões de faixas de score (aprovadas, a aplicar no conserto limpo)

> Semântica: extremo baixo `<=`, alto `>`; `between` = `(>lower e <=upper]`.
> Aplicar junto, via conserto definitivo (prod só com ordem explícita).

## RDW (aprovado)
Estrutura nova — **4 níveis** (gradiente contínuo de mortalidade; sem base p/ 6).
| nível | faixa | base |
|---|---|---|
| 5 (melhor) | `<= 13,0` | RDW baixo-normal = menor mortalidade (Perlstein, NHANES III) |
| 3 | `13,0 – 14,5` | normal, risco subindo |
| 1 | `14,5 – 15,6` | 14,5% = "elevado" na literatura; até o teto de ref. do lab (15,6%) |
| 0 (pior) | `> 15,6` | acima da referência = elevado |

Operação: setar L0,L1,L3,L5 conforme acima e **remover L2 e L4** atuais (10,5–11,6 e 13,0–14,5).

## Tempo de sono (aprovado)
| nível | faixa |
|---|---|
| 0 (pior) | `> 10` h |
| 1 | `<= 4` h |
| 3 | `4 – 6` h |
| 5 (melhor) | `6 – 8` h |
| 4 | `8 – 10` h |

Operação: L0 → `>10` (remover upper=5 espúrio); L1 → `<=4` (operator `<`→`<=`, upper=4, remover lower=4); L3/L4/L5 mantidos. Cobertura sem vão.

## Homocisteína (aprovado) — µmol/L, valores de medicina funcional
Níveis ÚNICOS (regra: sem `level` repetido). U-shaped assimétrico (alto mais grave).
| nível | faixa |
|---|---|
| 5 (melhor) | `6 – 8` |
| 4 | `8 – 11` |
| 3 | `<= 6` (baixa) |
| 2 | `11 – 15` |
| 1 | `15 – 30` (hiper leve/moderada) |
| 0 (pior) | `> 30` (intermediária/grave) |


## Consumo de Calorias (aprovado) — % do alvo. Abaixo = pior. Níveis únicos.
| nível | faixa |
|---|---|
| 0 (pior) | `<= 70%` |
| 1 | `> 150%` |
| 2 | `70 – 90%` |
| 3 | `110 – 150%` |
| 5 (melhor) | `90 – 110%` |

## Consumo de Verduras e Legumes (aprovado) — alinhar pelo rótulo
| nível | operator | lower | upper |
|---|---|---|---|
| 0 | `=` | 0 | — |
| 1 | `=` | 0,5 | — |
| 3 | `=` | 1 | — (remover 1,5) |
| 5 | `>=` | 2 | — (era 3; remover upper=2) |

## Hora de acordar (aprovado) — manhã = melhor; operadores limpos
| nível | faixa |
|---|---|
| 5 (melhor) | `5 – 7h` |
| 4 | `7 – 9h` |
| 3 | `9 – 11h` |
| 2 | `4 – 5h` |
| 1 | `<= 4h` |
| 0 (pior) | `> 11h` |

## Hora de dormir (aprovado) — pós-meia-noite codificado como +24; ótimo 22–23h (UK Biobank)
| nível | faixa (pós-meia-noite=+24) | hora |
|---|---|---|
| 5 (melhor) | `22 – 23` | 22–23h |
| 4 | `21 – 22` | 21–22h |
| 3 | `23 – 25` | 23h–1h |
| 2 | `19 – 21` | 19–21h |
| 1 | `25 – 27` | 1h–3h |
| 0 (pior) | `> 27` | após 3h |
Dormir de dia (5h–19h): IGNORAR (não classifica).

## ACTN3 rs1815739 R577X (aprovado) — longevidade/anti-fragilidade. Níveis únicos.
| nível | genótipo |
|---|---|
| 5 (melhor) | `RR` (poder; menos fragilidade) |
| 3 | `RX` (misto) |
| 0 (pior) | `XX` (+fragilidade no envelhecimento) |

## ADH1B rs1229984 Arg48His (aprovado) — His = protetor. Níveis únicos.
| nível | genótipo |
|---|---|
| 5 (melhor) | `His/His` (rápido — protetor) |
| 3 | `Arg/His` (intermediário) |
| 0 (pior) | `Arg/Arg` (lento — maior risco) |

## Colonoscopia - Número Total Adenomas (aprovado) — `=` mortos: pôr valor no limite
| nível | operator | lower |
|---|---|---|
| 5 | `=` | 0 |
| 4 | `=` | 1 |
| 3 | `=` | 2 |
(L0 >=10, L1 5-9, L2 3-4 mantidos)

## Perdas Dentárias (aprovado) — `=` mortos: pôr valor no limite
| nível | operator | valor |
|---|---|---|
| 5 | `=` | 0 (Dentição completa) |
| 4 | `=` | 1 |
(L0 >10, L1 7-10, L2 4-6, L3 2-3 mantidos)

## Razão AEC/ACT (%) (aprovado) — L2 morto vira faixa + fronteiras encostam
| nível | faixa |
|---|---|
| 5 (melhor) | `<= 36` |
| 4 | `36 – 37,5` |
| 3 | `37,5 – 39` |
| 2 | `39 – 40` |
| 1 | `40 – 42` |
| 0 (pior) | `> 42` |

## Doppler Carótidas - Estenose Carotídea (aprovado) — só conserta o morto
| nível | fix |
|---|---|
| 5 | `= 0` (era `=` sem limite) |
(L0 70-99, L1 50-69, L3 1-49 mantidos como estão)

## Endoscopia Alta - Barrett (Prague C) (aprovado) — 0 e 1 separados
| nível | fix |
|---|---|
| 5 | `= 0` (Ausente) |
| 4 | `= 1` (era `<1`) |
(L2 1-2,9, L1 3-4,9, L0 ≥5 mantidos)

## Radiografia panorâmica mandíbula/maxila (aprovado) — 0 separado do <20
| nível | operator | lower | upper |
|---|---|---|---|
| 5 | `=` | 0 | — |
| 4 | `between` | 0 | 20 |
| 2 | `between` | 20 | 40 |
| 1 | `between` | 40 | 60 |
| 0 | `>` | 60 | — |

## Urocultura (aprovado) — potências viram números reais; Negativa = ≤10; vãos FECHADOS
| nível | operator | lower | upper |
|---|---|---|---|
| 5 (melhor) | `<=` | — | 10 (Negativa) |
| 2 | `between` | 10 | 100000 |
| 1 | `between` | 100000 | 10000000 |
| 0 (pior) | `>` | 10000000 | — |
(vãos fechados pelo tool — cobertura contínua, sem buracos; 10⁵ = corte clássico de ITU)

## Itens sem L0/L5 obrigatórios (aprovado) — regra: L0=pior e L5=melhor obrigatórios
Conserto = **remapeamento de nível** (rótulos descrevem a opção/faixa, não mudam).

**17 "profissionais que acompanha"** (Cardiologista, Nefrologista, Dentista, Psicólogo, etc.) — hoje L0/L1/L2:
| era | vira | rótulo |
|---|---|---|
| L0 | L0 | Não acompanha (pior) |
| L2 | L3 | Acompanha — atrasado |
| L1 | L5 | Acompanha — em dia (melhor) |

**Uso de BiPAP/CPAP** — hoje L1/L3/L5:
| era | vira | rótulo |
|---|---|---|
| L1 | L0 | Tem indicação mas não usa (pior) |
| L3 | L3 | Usa com indicação |
| L5 | L5 | Não tem indicação (melhor) |

**Ferritina - Mulheres Pós-Menopausa** (numérico em U) — hoje L0–L3:
| era | vira | faixa | rótulo |
|---|---|---|---|
| L3 | L0 | >300 | Sobrecarga (pior) |
| L0 | L1 | <30 | Deficiência |
| L2 | L3 | 200–300 | Limítrofe-alta |
| L1 | L5 | 30–200 | Normal (melhor) |

## Itens com faixa INVERTIDA no prod (lower>upper) — reestruturados à mão (aprovado)
Estavam mortos no prod (between vazio sob `>lower e <=upper`). Excluídos do tool (manualSkip).

**Ecodopplercardiograma - GLS** (mais negativo = melhor):
| nível | faixa |
|---|---|
| 5 | `<= -22` (ótimo) |
| 4 | `-22 a -20` |
| 3 | `-20 a -18` |
| 2 | `-18 a -17` |
| 1 | `-17 a -16` |
| 0 | `> -16` (anormal) |

**ECG - Eixo Cardíaco** (bidirecional; normal no meio; fronteira -30 textbook p/ sem vão):
| nível | faixa | leitura |
|---|---|---|
| 5 | `-30 a +90` | normal |
| 4 | `-60 a -30` | desvio esq. leve |
| 3 | `+90 a +119` | desvio dir. leve |
| 2 | `-90 a -60` | desvio esq. |
| 1 | `> +119` | desvio dir. |
| 0 | `<= -90` | desvio esq. extremo (pior) |
