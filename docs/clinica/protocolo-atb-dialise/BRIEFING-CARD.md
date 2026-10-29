# Briefing do card: protocolo de antibióticos na hemodiálise

Fonte única do conteúdo do card. Tudo aqui foi decidido pela equipe (registro e fontes em `PLANO.md`).
O card só pode conter o que está neste arquivo. Mudou aqui, muda no card.

## Diretrizes de forma (o que a equipe pediu)

- **Prático, não slide de aula.** Pouco texto, nada de parágrafo. Consulta em segundos, na beira da cadeira.
- **Mesma lógica sempre.** Toda infecção segue o mesmo molde; toda droga segue o mesmo molde.
- **Molde da infecção:** Situação → Antibiótico → Duração → Atenção (no máximo uma linha).
- **Molde da droga EV:** Dose → Diluição → Via → Tempo → Momento na sessão.
- **Molde da droga VO:** Dose → Frequência → Dias de hemodiálise.
- Doses ficam só na ficha de drogas; nas infecções aparece só o nome da droga.
- Sem marca comercial, sem travessão, sem fonte no card (fontes ficam no documento teórico).
- Ilustração só como ícone de seção (gpt-image-2.5); número e texto sempre no HTML.
- **Sem legenda** das formas (decidido 2026-10-29).

---

## 1. Infecção de cateter

### Empírico
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Bacteremia por cateter | Vancomicina + ceftazidima | 14 dias após hemocultura negativa | Lock se o cateter fica |

### Ajuste pela cultura (toda resposta tem ramo)
| Resultado da hemocultura | Antibiótico | Atenção |
|---|---|---|
| **Negativa** | Manter vancomicina + ceftazidima (empírico) | Discussão original: "Vanco / ceftazidima empírica - 14 dias" |
| Estafilococo **sensível** à oxacilina (*S. aureus* ou coagulase-negativo) | **Só cefazolina** (suspender vanco e cefta; decidido 2026-10-29) | IDSA 2009; extensão ao coagulase-negativo aprovada |
| Estafilococo **resistente** à oxacilina (MRSA ou coagulase-negativo) | Vancomicina (suspender cefta) | |
| Gram-negativo | Pelo antibiograma (suspender vanco) | |
| **Outro germe ou polimicrobiana** | Ajustar pelo antibiograma | Enterococo: regra da IDSA só no texto teórico |
| *Candida* | Pela cultura | **Retirar o cateter** |

### Duração (depende só do germe; conta SEMPRE da 1ª hemocultura negativa, nunca da troca)
| Situação | Duração |
|---|---|
| Hemocultura negativa (empírico) | 14 dias do início do antibiótico |
| Outros germes | 14 dias após hemocultura negativa |
| *S. aureus* (com lock ou após troca) | **3 semanas** após hemocultura negativa |
| *Candida* | 14 dias após hemocultura negativa (IDSA candidíase 2016, Pappas) · confirmado pela equipe |
| Endocardite | 4 semanas (valva protética: 6) |

Ecocardiograma em todo *S. aureus*. A troca do cateter é decisão separada (abaixo) e **não muda o prazo**.

### Manter, trocar ou retirar o cateter (avaliação em 72 h, qualquer germe)
1. **Instável, sepse ou infecção à distância?** Sim → **retirar o cateter**. (IDSA 2009 rec. 32, 59)
2. Se não: **febre persistente ou hemocultura de 72 h positiva?** Sim → **trocar por fio-guia** (regra
   conservadora da equipe, 2026-10-29; IDSA rec. 33, 57, 66 apoiam o fio-guia no assintomático; na febre a IDSA
   retiraria). Não → **manter com lock**.
- *Candida*: retirar sempre (no ramo da cultura).
- A troca não muda a duração (ver Duração).

### Hemoculturas
- 2 amostras: cateter + linha venosa.
- **D0** antes do antibiótico → **72 h** sempre (positiva: repetir a cada 72 h até negativar) → **fim do tratamento**.

### Troca e novo acesso
- Troca por falha: só após 72 h de antibiótico.
- Novo permcath: 72 h após hemocultura negativa.
- Retirado por hemocultura positiva: tentar ficar 72 h sem acesso.
- Procurar focos: acessos antigos, stent venoso, marcapasso, prótese ortopédica.

### Lock (sempre que o cateter fica, após cada sessão, todo o tratamento)
| Componente | Concentração | Volume |
|---|---|---|
| Vancomicina | 5 mg/ml | 1 ml |
| Ceftazidima | 5 mg/ml | 0,5 ml |
| Heparina | 5.000 UI/ml | 0,5 ml |
| SF 0,9% | | até o volume de cada via (impresso no ramo do cateter) |

### Óstio e túnel
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Óstio | Mupirocina tópica (sem mupirocina: cefazolina) | 14 dias | Gaze + fita, sem filme transparente. Com secreção: cultura + hemocultura |
| Tunelite | Cefazolina → MRSA prévio: vanco → *Pseudomonas* ou falha: + cefta → bacteremia: vanco + cefta | 14 dias | Lock não trata o túnel: trocar por fio-guia com túnel novo |

## 2. FAV e prótese
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| FAV nativa | Pela cultura | 2 semanas | Doppler inicial |
| Prótese | Pela cultura | 4 semanas | Doppler semanal. Extensa, abscesso, exposta ou sangrando → cirurgia vascular |

## 3. Infecção urinária
Urocultura + antibiograma para todos. Sem urocultura de controle.

| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Adere ao oral | Ciprofloxacino VO | 7 dias | |
| Não adere ao oral | Ceftazidima EV (alternativa: amicacina EV) | 7 dias | Amicacina é ototóxica |

## 4. Pneumonia
Raio-X nos casos arrastados ou graves.

| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Em geral | Amoxicilina-clavulanato + azitromicina VO | 5 dias (azitro 3 dias) | |
| Doença pulmonar prévia | Levofloxacino VO | 3 doses | |
| Sem condição social | Levofloxacino EV ou VO na sessão | 3 doses | Mesma dose por qualquer via |

## 5. Celulite
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Em geral | Cefalexina VO ou cefazolina EV | 7 dias corridos | |
| MRSA | Clindamicina VO | 7 dias corridos | |
| Grave | Vancomicina | 14 dias corridos | |

---

## 6. Ficha de drogas

### Endovenosas
| Droga | Dose | Diluição | Via | Tempo | Momento |
|---|---|---|---|---|---|
| Vancomicina, ataque | 35 mg/kg (mín. 2 g, máx. 3 g) | 10 mg/ml em SF | Linha da HD | 90 min | 90 min finais |
| Vancomicina, manutenção | 15 mg/kg (mín. 1 g, máx. 1,5 g) | 10 mg/ml em SF | Linha da HD | 60 min | Final de toda sessão |
| Ceftazidima | 1 g; 2 g antes do intervalo de 72 h | 100 ml de SF | Linha da HD | 30 min | Fim da sessão |
| Cefazolina | 20 mg/kg (mín. 1 g, máx. 2 g) | SF, ~10 ml por grama | Linha venosa | Bolus 3-5 min | Fim da sessão, toda sessão |
| Amicacina | 5-7,5 mg/kg | 100-200 ml de SF | EV | 30-60 min | Após a sessão, toda sessão |
| Levofloxacino EV | 750 mg, depois 500 mg a cada 48 h (3 doses) | Bolsa pronta 5 mg/ml | EV | 60-90 min | Qualquer horário |

### Orais
| Droga | Dose | Frequência | Dias de hemodiálise |
|---|---|---|---|
| Ciprofloxacino | 500 mg | 1x/dia | Tomar após a sessão |
| Amoxicilina-clavulanato | 500/125 mg | 1x/dia | + 1 durante e + 1 ao fim da sessão. Não usar 875 mg |
| Azitromicina | 500 mg | 1x/dia, 3 dias | Sem ajuste |
| Levofloxacino VO | 750 mg, depois 500 mg | A cada 48 h, 3 doses | Qualquer horário (não dialisa) |
| Cefalexina | 500 mg | 1x/dia | Tomar após a sessão |
| Clindamicina | 300 mg | 8/8 h | Sem ajuste |

---

## Fica fora do card (vai só no documento teórico)
- Nível de vancomicina pré-diálise (15-20 mg/l), quando a unidade tiver.
- Pré-medicação com anti-histamínico antes da vancomicina.
- Hemocultura de controle em 48 h (diretriz) vs 72 h (unidade).
- Diretrizes pedem imagem para confirmar pneumonia (RX inviável na unidade).
- Amoxicilina-clavulanato: bula brasileira não recomenda com ClCr < 10; dose da bula americana.
- Celulite: IDSA 2014 sugere 5 dias; unidade usa 7.
- Sulfametoxazol-trimetoprima excluído (hipercalemia; bula: ClCr < 15 não recomendado).

## Regras de processo (pedidas pela equipe)
- **Briefing primeiro, card depois.** Toda mudança de conteúdo entra aqui antes de mexer no card.
- **Conferir contra a discussão original** (`PLANO.md`, seção "Proposta da equipe") antes de cada versão:
  nada da discussão pode sumir do card sem decisão registrada.

## Pendências de conteúdo
- Nenhuma.
