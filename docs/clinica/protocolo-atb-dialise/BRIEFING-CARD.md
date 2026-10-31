# Briefing do card: protocolo de antibióticos na hemodiálise

Fonte única do conteúdo do card. Tudo aqui foi decidido pela equipe (registro e fontes em `PLANO.md`).
O card só pode conter o que está neste arquivo. Mudou aqui, muda no card.

## Diretrizes de forma (o que a equipe pediu)

- **Unidades (equipe 2026-10-30):** DaVita Londrina | Bandeirantes | Lago Parque | Rolândia, no cabeçalho de cada folha
  e na abertura do texto. O texto que acompanha o card se chama **"Fundamentação científica"** (não "documento teórico";
  o arquivo continua `TEORICO.md`).

- **Prático, não slide de aula.** Pouco texto, nada de parágrafo. Consulta em segundos, na beira da cadeira.
- **Mesma lógica sempre.** Toda infecção segue o mesmo molde; toda droga segue o mesmo molde.
- **Molde da infecção:** Situação → Antibiótico → Duração → Atenção (no máximo uma linha).
- **Molde da droga EV:** Dose → Diluição → Via → Tempo → Momento na sessão.
- **Molde da droga VO:** Dose → Frequência → Dias de hemodiálise.
- Doses ficam só na ficha de drogas; nas infecções aparece só o nome da droga.
- **Nomes comerciais de material são mantidos** quando a equipe usa (Hipafix, Tegaderm): é protocolo interno, não conteúdo público. Sem travessão, sem fonte no card (fontes ficam no documento teórico).
- Ilustração só como ícone de seção (gpt-image-2.5); número e texto sempre no HTML.
- **Sem legenda** das formas (decidido 2026-10-29).
- **Formato: A4 vertical** (decidido 2026-10-29; começando pela folha 1).
- **Listas nunca começam por "outros/demais"**: primeiro os nomeados, o genérico por último.
- **Rótulo nunca encosta em caixa ou seta** (conferir com imagem ampliada).
- **Toda resposta tem caixa (equipe 2026-10-30: "as respostas soltas ficam muito ruins, tudo tem que ter box").** Padrão de
  algoritmo clínico (Society for Medical Decision Making 1992: situação clínica, decisão, ação, fluxo): o resultado da
  hemocultura vira caixa de situação (branca, borda arredondada), "Sim"/"Não" ficam em caixinha com borda sobre a seta.
- **Durações sempre em dias (equipe 2026-10-30: "tudo dias")**: 14, 21, 28, 42 dias, nunca semanas.
- **Trocar por fio-guia não é alerta** (equipe 2026-10-30): caixa de ação comum; alerta (coral) só para retirar o cateter.

---

## 1. Infecção de cateter

### Empírico
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Bacteremia por cateter | Vancomicina + ceftazidima | 14 dias após hemocultura negativa | Lock após cada sessão |

### Ajuste pela cultura (toda resposta tem ramo)
| Resultado da hemocultura | Antibiótico | Atenção |
|---|---|---|
| **Negativa** | Manter vancomicina + ceftazidima (empírico) | Discussão original: "Vanco / ceftazidima empírica - 14 dias". **Também passa pela avaliação de 72 h** (decidido 2026-10-29) |
| ***S. aureus* oxa-sensível** | **Só cefazolina** (suspender vanco e cefta; decidido 2026-10-29) | IDSA 2009. Segue a faixa própria do *S. aureus* (abaixo) |
| ***S. aureus* oxa-resistente** | Manter vancomicina (suspender cefta) | Segue a faixa própria do *S. aureus* (abaixo) |
| **Coagulase-negativo** | Cefazolina ou vanco pela oxacilina (suspender cefta) | Oxa-sensível → cefazolina; oxa-resistente → vanco. Uma caixa só (equipe 2026-10-30) |
| Gram-negativo | Pelo antibiograma (suspender vanco) | |
| **Outro germe ou polimicrobiana** | Ajustar pelo antibiograma | Enterococo: regra da IDSA só no texto teórico |
| *Candida* | **Antifúngico pela cultura; suspender vanco e cefta** (decidido 2026-10-29) | **Retirar o cateter**. Candida + bactéria = ramo "outro germe ou polimicrobiana" |

### Duração (sempre em dias, "após HC negativa"; conta da 1ª hemocultura negativa, nunca da troca)
| Prazo | Para quem |
|---|---|
| **14 dias após HC negativa** | Coagulase-negativo · Gram-negativo · Outro germe · Hemocultura negativa (a 1ª negativa é a do início) · *Candida* (IDSA 2016, confirmado) |
| **28 dias após HC negativa** | *S. aureus* (após troca por fio-guia ou retirada). Equipe 2026-10-30, feedback item 4 ("28"). Base: KDOQI Fig. 25.1 (4 semanas), IDSA rec. 78 (4-6 semanas, B-II). Divergência: consenso IDSA/ESCMID 2026 (14 dias no baixo risco) |
| **28 dias** | Endocardite (valva protética: 42 dias). No card fica em "Sempre", abaixo do ecocardiograma |

Ecocardiograma: *S. aureus*, *Candida* ou HC positiva em 72 h (feedback item 5). A troca do cateter é decisão separada (abaixo) e **não muda o prazo**.
No card (equipe 2026-10-30): a duração é caixa no pé do fluxo, ligada por seta à conduta final de cada caminho
(*S. aureus* → 28 dias; lock ou fio-guia → 14 dias; *Candida* → 14 dias). A faixa "Duração" saiu; com ela saiu o
subtítulo "com ou sem troca do cateter".

### Manter, trocar ou retirar o cateter
- **A qualquer momento** (alarme ligado à caixa de entrada, à direita; texto na ordem "Se instável, sepse ou infecção à distância" → "Retirar o cateter" → "a qualquer momento", decidido 2026-10-29): **instável, sepse ou infecção à distância → retirar o
  cateter**. (IDSA 2009 rec. 32, 59.) O alarme não espera as 72 h (confirmado pela equipe 2026-10-29).
- **Em 72 h, demais germes** (coagulase-negativo, gram-negativo, outro germe, hemocultura negativa): **febre persistente ou
  hemocultura positiva?** Não → **manter com lock**. Sim → **Acesso difícil?** Sim → **trocar por fio-guia**; Não → **retirar
  o cateter** (equipe 2026-10-30, feedback item 3; IDSA rec. 32, 33, 57, 59; KDOQI texto b-c).
- **Em 72 h, *S. aureus*** (equipe 2026-10-30, feedback item 2; faixa própria no card, à esquerda, fundo coral):
  **febre ou HC positiva?** Não → **trocar por fio-guia** (sem tentar manter com lock). Sim → **Acesso difícil?** Sim →
  trocar por fio-guia (mesma caixa, seta lateral); Não → **retirar o cateter** (feedback item 3).
  Base: KDOQI 2019 Fig. 25.1 (no *S. aureus*, fio-guia abaixo de "resolução em 2-3 dias"); IDSA 2009 rec. 59 (B-II: sintomas
  resolvidos em 2-3 dias → fio-guia; persistentes → retirar); Aslam 2014 (fio-guia cura mais que lock no *S. aureus*, OR 4,72).
  **Divergência:** IDSA 2009 rec. 57 (A-II) manda retirar sempre no *S. aureus*; a única referência da rec. 57 é Maya 2007,
  que mostra a falha do lock e não compara retirar com fio-guia.
- *Candida*: retirar sempre (no ramo da cultura).
- A troca não muda a duração (ver Duração).

### Hemoculturas
- 2 amostras: cateter + linha venosa.
- **D0** antes do antibiótico → **72 h** sempre (positiva: repetir a cada 72 h até negativar) → **fim do tratamento**.

### Troca e novo acesso
- Troca por falha: só após 72 h de antibiótico.
- Novo permcath: 72 h após hemocultura negativa.
- Retirado por hemocultura positiva: tentar ficar 72 h sem acesso.
- **Acesso difícil (equipe 2026-10-30, "mais concreto"):** jugulares internas e femorais estenosadas ou ocluídas. Base: KDOQI
  2019 (ordem de sítios: jugular interna → jugular externa → femoral → subclávia → lombar; "the internal jugular and femoral
  veins may become inaccessible, either due to stenosis or chronic total occlusion"); IDSA rec. 57 (fio-guia só sem sítio
  alternativo). Sem número de cateteres prévios (nenhuma fonte fixa).
- Procurar focos: acessos antigos, stent venoso, marcapasso, prótese ortopédica.

### Lock (sempre que o cateter fica, após cada sessão, todo o tratamento)
| Componente | Concentração | Volume |
|---|---|---|
| Vancomicina | 5 mg/ml | 1 ml |
| Ceftazidima | 5 mg/ml | 0,5 ml |
| Heparina | 5.000 UI/ml | 0,5 ml |
| SF 0,9% | | até o volume de cada via. Nota "(impresso no ramo do cateter)" **removida** (equipe 2026-10-30: "tira") |

### Óstio e túnel
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Óstio | Mupirocina tópica (sem mupirocina: cefazolina) | 14 dias | Curativo com gaze e Hipafix (não Tegaderm), como na discussão original. **Sem cultura de secreção** (não se usa no serviço; removido 2026-10-29) |
| Tunelite | Cefazolina → MRSA prévio: vanco → *Pseudomonas* ou falha: associar ceftazidima | 14 dias | **Colher hemocultura antes do antibiótico** (discussão original), **cateter + periférica**; no card vira caixa de ação antes da cefazolina (equipe 2026-10-29). Na caixa da cefazolina, ***Pseudomonas* ou falha: associar ceftazidima** logo abaixo de "MRSA prévio: vanco" (equipe 2026-10-29); as duas notas saem do rodapé. **Cateter fica** (discussão original). Hemocultura positiva ou bacteremia: seguir a folha 1. **Sem troca por fio-guia com novo túnel** (não se faz mais no serviço; removido 2026-10-29) |

## 2. FAV e prótese
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| FAV nativa | Vanco + cefta empírico → ajustar pela cultura | 14 dias; **bacteremia por *S. aureus*: 28 dias** (equipe 2026-10-30, feedback item 4) | Doppler inicial |
| Prótese | Vanco + cefta empírico → ajustar pela cultura | **28 dias em qualquer infecção** (inclusive *S. aureus* e após cirurgia) | Doppler semanal. Pergunta no ramo: "Infecção extensa, abscesso, exposta ou sangrando?" Sim → **cirurgia vascular para retirada** → 28 dias; Não → 28 dias. **Todo ramo termina em duração** |

Empírico vanco + cefta decidido 2026-10-29 (KDOQI 2019: S. aureus/epidermidis 70-90%).
Regra da equipe (2026-10-29): **3 semanas no *S. aureus* só na FAV nativa; na prótese qualquer infecção é 4 semanas.**
Ecocardiograma: *S. aureus*, *Candida* ou HC positiva em 72 h (regra da seção 1).
Texto do alerta: item 7 do PLANO. Forma (pergunta + alerta ligado ao ramo): proposta minha, aprovada 2026-10-29.
Correção 2026-10-29: a primeira versão deixou o ramo da cirurgia sem duração; a equipe apontou.
Pedido da equipe (2026-10-29): **"*S. aureus*: 3 semanas" dentro da própria pílula de duração da FAV nativa**
(subtítulo de "2 semanas"), não em nota. Quadro atual reprovado ("horrível"): verticalmente espremido e
"ajustar pela cultura" sem dizer quando colher e quando repetir. **Em estudo:** a FAV/prótese usa o mesmo tronco
da folha 1 (hemoculturas, vanco + cefta, ajuste pelo germe); na discussão original ela está dentro do bloco
"Infecção de cateter". A opção A (juntar na folha 1) foi escolhida e **revertida pela equipe no mesmo dia: "melhor deixar separado"**.
**Estrutura decidida 2026-10-29:** **Óstio e Túnel vai para a folha 1** (é infecção do cateter); **FAV e Prótese
ocupa uma coluna inteira da folha 2** (altura toda), com Urina, Celulite e Pneumonia na outra coluna.
Feito no card (2026-10-29): Óstio e Túnel embaixo do fluxo da folha 1, ao lado de Hemoculturas / Novo Acesso /
Sempre; nota do túnel passou de "seguir a folha 1" para "seguir o fluxo acima" (só a referência mudou). FAV e
Prótese na coluna esquerda inteira da folha 2; "*S. aureus*: 3 semanas" dentro da pílula de 2 semanas.
**Hemoculturas na FAV/prótese (decidido 2026-10-29, "seguem"):** mesmo esquema do cateter: D0 antes do
antibiótico, 72 h sempre, a cada 72 h se positiva até negativar, fim do tratamento. No card: a caixa inicial
passa a "Colher 2 hemoculturas e iniciar vancomicina + ceftazidima" (mesma frase da folha 1) e a linha do tempo
das hemoculturas entra na coluna da FAV. "Ajustar pelo germe (folha 1)" como subtítulo: **proposta minha**.
**Amostras na FAV/prótese (decidido 2026-10-29): linha venosa + acesso periférico** (subtítulo da caixa inicial).
**Ajuste pela hemocultura repetido no próprio quadro** (equipe 2026-10-29: "ajustar pelo germe é preguiçoso"),
mesmos ramos da folha 1: Negativa → manter vanco + cefta · Estafilo oxa-sensível → só cefazolina (suspender vanco e
cefta) · Estafilo oxa-resistente → manter vanco (suspender cefta) · Gram-negativo → suspender vanco, seguir
antibiograma · Outro germe → ajustar pelo antibiograma · *Candida* → antifúngico pela cultura, suspender vanco e
cefta. ***Candida* na prótese: retirar a prótese** (equipe 2026-10-29). No card: caixa de alerta (coral) como na folha 1,
"Antifúngico / pela cultura · suspender vanco e cefta / Prótese: retirar". *Candida* na FAV nativa: **sem conduta extra** (só antifúngico e suspender vanco e cefta; decidido 2026-10-29).
***S. aureus* na prótese sem infecção extensa: deixa como está** (só antibiótico; retirada só pelo ramo de gravidade). Decidido pela equipe 2026-10-30 após ver KDOQI 2019 G16 ("AVG infections can rarely be treated with antibiotics alone"); registrar a divergência no texto teórico.
**Forma (equipe 2026-10-29):** a grade de germes "ficou confusa"; **FAV e Prótese em largura inteira, no molde da
folha 1** (leque "Qual germe?" com os ramos lado a lado). Não cabe com as outras infecções: **FAV e Prótese vira a
folha 2 própria**; Urina, Celulite e Pneumonia vão para a folha 3 (revisão depois); drogas + lock viram folha 4.
**Gatilho (equipe 2026-10-30: "iniciar com sintomas"):** caixa de situação no topo, mesmos sintomas da folha 1:
"Febre ou calafrio na sessão / paciente com FAV ou prótese" (cópia da folha 1, **reprovado**: "você só copiou de
cateter"). **Aprovado 2026-10-30 ("aprova"), a partir do estudo (PLANO):** gatilho "Febre ou sinais no
acesso", com subtítulo "vermelhidão · dor · secreção · pele rompida ou prótese exposta"; nota "Sinal local basta,
mesmo sem febre · prótese antiga ou trombosada com febre sem foco: suspeitar".
**FAV nativa e Prótese viram caixas de situação** (equipe 2026-10-30: "tá só um textinho pequeno"), como Óstio/Túnel.
**Ordem corrigida (2026-10-29, erro meu apontado pela equipe):** coleta + vanco/cefta → "Qual acesso?" → Doppler
(e, na prótese, a pergunta de gravidade com alerta de cirurgia) → "Hemocultura positiva?" → germes → faixa de
Duração (FAV nativa 2 semanas, *S. aureus* 3; prótese 4 semanas, qualquer germe, também após a cirurgia). Antes o
Doppler e a gravidade vinham depois da cultura e a *Candida* mandava retirar a prótese e depois seguia para o
Doppler semanal dela.

## 3. Infecção urinária
Urocultura + antibiograma para todos. Sem urocultura de controle.

| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Adere ao oral | Ciprofloxacino VO | 7 dias | |
| Não adere ao oral | Ceftazidima EV (alternativa: amicacina EV) | 7 dias | Amicacina é ototóxica. **Card mantido como está** ("Ceftazidima EV / ou amicacina" + nota ototóxica; equipe 2026-10-30: "mantém como está") |

## 4. Pneumonia
Raio-X nos casos arrastados ou graves. Pergunta no card: "Adere ao oral?" (mesma da urina e da celulite; confirmado pela equipe 2026-10-29).
Proposta de pergunta única ("Doença pulmonar prévia ou não adere ao oral?") **reprovada** 2026-10-30 ("ficou mais confuso"); volta às duas perguntas.

| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Em geral | Amoxicilina-clavulanato + azitromicina VO | 5 dias (azitro 3 dias) | |
| Doença pulmonar prévia | Levofloxacino VO | **3 doses** | Equipe 2026-10-30 (revisto): "levo 750 + 500 + 500 nos dias de diálise mesmo, mesmo oral". Substitui "oral = 5 dias" e a dose cheia 750 mg/dia |
| Não adere ao oral | Levofloxacino EV | 3 doses | Discussão original: "levofloxacino EV por 3 doses". O "ou VO na sessão" foi acréscimo meu na consolidação, sem decisão da equipe; removido 2026-10-30 (equipe: "não adere ao oral - VO??") |

## 5. Celulite
| Situação | Antibiótico | Duração | Atenção |
|---|---|---|---|
| Em geral, adere ao oral | Cefalexina VO | 7 dias corridos | |
| Em geral, não adere ao oral | Cefazolina EV na sessão (decidido 2026-10-29) | 7 dias corridos | |
| MRSA prévio (cultura anterior com MRSA; decidido 2026-10-29) | Clindamicina VO | 7 dias corridos | |
| Grave | Vancomicina | 14 dias corridos | |

**"Dias corridos" sai do card** (equipe 2026-10-30: "não precisa dias corridos"): pílulas "7 dias" e "14 dias", sem nota.

No card: perguntas "Grave?" e depois "MRSA prévio?"; a adesão vai como subtítulo da cefalexina ("ou cefazolina EV se não adere"), não como pergunta (aprovado 2026-10-29).

---

## 6. Ficha de drogas
Textos padronizados em 2026-10-30 (equipe: "padronize os textos"): diluição sempre "SF: …"; momento "Fim de toda sessão"
/ "Após toda sessão"; intervalo "48/48 h"; duração só nos fluxos (sai da ficha); "Sem ajuste" para o que não muda na HD.

### Endovenosas
| Droga | Dose | Diluição | Via | Tempo | Momento |
|---|---|---|---|---|---|
| Vancomicina, ataque | 35 mg/kg (mín. 2 g · máx. 3 g) | **SF 250 ml; 3 g: 300 ml** (equipe 2026-10-30, feedback item 6, opção A: sempre ≤ 10 mg/ml) | Linha da HD | 90 min | 90 min finais da sessão |
| Vancomicina, ataque fora da HD | 20 mg/kg (mín. 1 g · máx. 2 g) (equipe 2026-10-31, feedback item 7) | SF 250 ml (equipe 2026-10-31) | EV | 60 min | Fora da sessão |
| Vancomicina, manutenção | 15 mg/kg (mín. 1 g · máx. 1,5 g) | SF 250 ml (equipe 2026-10-31) | Linha da HD | 60 min | Fim de toda sessão |
| Ceftazidima | 1 g (2 g antes do intervalo de 72 h) | SF: 100 ml | Linha da HD | 30 min | Fim de toda sessão |
| Cefazolina | 20 mg/kg (mín. 1 g · máx. 2 g) | SF: 10 ml por g | Linha venosa | 3-5 min (bolus) | Fim de toda sessão |
| Amicacina | 5-7,5 mg/kg | SF: 100-200 ml | EV | 30-60 min | Após toda sessão |
| Levofloxacino | 750 mg → 500 mg → 500 mg (3 doses, nos dias de HD) | Bolsa pronta: 5 mg/ml | EV | 60-90 min | Qualquer horário |

### Orais
| Droga | Dose | Frequência | Dias de hemodiálise |
|---|---|---|---|
| Ciprofloxacino | 500 mg | 1x/dia | Após a sessão |
| Amoxicilina-clavulanato | 500/125 mg | 1x/dia | Após a sessão (equipe 2026-10-30, opção "a": sem dose extra; substitui a bula americana) |
| Azitromicina | 500 mg | 1x/dia | Sem ajuste (duração 3 dias no fluxo) |
| Levofloxacino | 750 mg → 500 mg → 500 mg | 3 doses | Uma dose por dia de HD (igual ao EV; equipe 2026-10-30) |
| Cefalexina | 500 mg | 1x/dia | Após a sessão |
| Clindamicina | 300 mg | 8/8 h | Sem ajuste |

### Lock (na folha 4)
| Componente | Concentração | Volume |
|---|---|---|
| Vancomicina | 5 mg/ml | 1 ml |
| Ceftazidima | 5 mg/ml | 0,5 ml |
| Heparina | 5.000 UI/ml | 0,5 ml |
| SF | 0,9% | Até completar a via |

---

## Divergências da busca de 2026-10-30 (decisões, uma por vez)
1. **Hemocultura negativa: MANTER vanco + cefta 14 dias** (equipe: "vamos manter"). IDSA 2009 rec 55 e KDOQI Fig. 25.1
   mandam suspender; registrar como divergência no texto.
2. **Hemocultura de vigilância: MANTER no fim do tratamento** (equipe: "mantém no fim"). IDSA rec 67 pede 1 semana após o
   fim com cateter mantido; registrar como divergência.
3. **Coagulase-negativo com cateter retirado: MANTER 14 dias** (equipe: "mantém 14"). IDSA rec 75 pede 5-7 dias se retirado;
   KDOQI Fig. 25.1 dá 14 dias; registrar como divergência da IDSA.
4. **Tunelite (equipe 2026-10-30, "confirma"): empírico = CEFTAZIDIMA; MRSA prévio = vancomicina + ceftazidima; 14 dias;
   cateter fica; cefazolina SAI da tunelite.** Base: bula FDA ceftazidima (pele/partes moles por MSSA e gram-negativos) e
   KDOQI (empírico gram+ e gram-). Divergências a registrar: IDSA rec 44 e Miller 2016 (retirar o cateter); ISPD 2023
   (cefalosporina de 1ª geração); ceftazidima menos potente contra S. aureus in vitro (J Antimicrob Chemother 1981).
   Substitui a "tunelite revisada" de 2026-10-29 (cefazolina escalonada). **Card e texto a atualizar.**
5. **Celulite grave: MANTER vancomicina isolada 14 dias** (equipe: "mantém"). IDSA 2014: grave = vanco + pip-tazo ou
   carbapenêmico, 5 dias (estender se não melhora); registrar como divergência.
6. **ITU: MANTER como está** (equipe: "mantém assim"). Registrar divergências: homem 14 dias (EAU 2026), amicacina não
   preferida e evitar FQ com uso nos últimos 12 meses (IDSA 2025).

## Referências no card (decidido 2026-10-30, opção "a": número elevado junto do item)
**Regra da equipe (2026-10-30): toda decisão do card tem pelo menos uma referência.** Se a fonte diverge, a referência
entra assim mesmo e o texto teórico registra a divergência. Item sem fonte encontrada: dizer isso à equipe, nunca inventar.
Numeração = `TEORICO.md` aprovado. Só leva número o item que o texto sustenta com aquela fonte; decisão da equipe sem
fonte fica sem número. Rodapés: "Números: referências do documento teórico".
- **Mapa item → referência: `REFERENCIAS-CARD.md`** (regerado 2026-10-30 após a busca "uma referência por decisão";
  numeração do TEORICO com 40 referências). O mapa anterior desta seção ficou obsoleto com a renumeração.

## Fica fora do card (vai só no documento teórico)
- Nível de vancomicina pré-diálise (15-20 mg/l), quando a unidade tiver.
- Pré-medicação com anti-histamínico antes da vancomicina.
- Hemocultura de controle em 48 h (diretriz) vs 72 h (unidade).
- Diretrizes pedem imagem para confirmar pneumonia (RX inviável na unidade).
- Amoxicilina-clavulanato: card simplificado para 500/125 mg 1x/dia após a sessão (equipe 2026-10-30). Bula americana pede
  dose extra durante e ao fim da diálise; bula brasileira não recomenda o comprimido com ClCr < 10; não usar 875 mg.
- Celulite: IDSA 2014 sugere 5 dias; unidade usa 7.
- Pneumonia: ATS 2025 sugere menos de 5 dias (mínimo 3) no não grave; unidade usa 5, como a ATS/IDSA 2019 (mínimo 5). **Mantidos 5 dias** (equipe 2026-10-30: "mantém os 5").
- IDSA/ESCMID 2026, ESC 2023 e ATS 2025 conferidas no texto primário em 2026-10-30 (PLANO).
- Sulfametoxazol-trimetoprima excluído (hipercalemia; bula: ClCr < 15 não recomendado).

## Feedback da equipe de 2026-10-30 (8 sugestões, decididas uma por vez)
1. **Rótulos "Estafilo oxa-sensível" e "Estafilo oxa-resistente"** (sugestão da equipe; aprovado 2026-10-30). Só nome,
   conduta igual (oxa-sensível → cefazolina). Vale nas setas das folhas 1 e 2. Na folha 1 o leque foi refeito no item 2.
2. **Retirar cateter em *S. aureus*, *Pseudomonas* e *Candida*.** Decidido germe a germe.
   - ***S. aureus* (aprovado 2026-10-30, "gostei"):** conservador. Faixa própria; em 72 h, melhorou → trocar por fio-guia;
     não melhorou → retirar. Ver "Manter, trocar ou retirar o cateter".
   - ***Pseudomonas* (decidido 2026-10-30, "mantém como está"):** segue o fluxo dos gram-negativos (72 h: lock ou fio-guia).
     Divergência com IDSA rec. 57 e KDOQI (texto, item e); apoio: Poole 2004 [12], lock cura 87% dos gram-negativos.
   - *Candida*: já retirava.
   Folha 1 refeita (equipe 2026-10-30): sem hexágono "Qual germe/Resultado da hemocultura"; do "Colher 2 hemoculturas"
   saem 7 caixas de situação: *S. aureus* oxa-sensível, *S. aureus* oxa-resistente, Coagulase-negativo, Gram-negativo,
   Outro germe, Hemocultura negativa (penúltimo), *Candida* (último).
3. **Falha em 72 h → retirar; fio-guia só com acesso difícil** (sugestão da equipe; aprovado 2026-10-30, mock 3 "aprovado").
   Vale nas duas faixas. Hexágono "Acesso difícil?" depois do "Sim" das 72 h. Definição em "Novo Acesso" (acima).
   Forma: faixa do *S. aureus* mais larga; Óstio e Túnel e a coluna da direita compactados para dar espaço ao fluxo;
   endocardite continua em "Sempre" (equipe: "pode continuar ali no sempre").
4. ***S. aureus* 28 dias** (equipe 2026-10-30, "28", após tabela de força de evidência de 14/28/42 dias; nenhum ensaio
   compara durações). Cateter: 21 → 28 dias. FAV nativa: 21 → 28 dias (equipe: "28").
5. **Ecocardiograma: *S. aureus*, *Candida* ou HC positiva em 72 h** (equipe 2026-10-30). "Mais de 2 amostras" = mais de duas
   coletas com 72 h de intervalo; a equipe escolheu antecipar para a HC positiva em 72 h (opção A), que também cobre "bacteremia
   persistente". *S. aureus*: consenso 2026 afirmação 3 [4], ESC 2023 IIa B [5]. HC positiva em 72 h: consenso 2026 afirmação 1
   (HC ≥ 48 h = fator-chave) [4], KDOQI Fig. 25.1 [2]. *Candida*: "eco pelo KDOQI - procurar metastáticas" [2]; Pappas 2016 não
   pede eco (só fundo de olho) [3] → divergência na seção 8. Folhas 1 e 2.
6. **Vancomicina de ataque em SF 250 ml; 3 g em 300 ml** (aprovado 2026-10-31, opção A). 250 ml fixo daria 12 mg/ml com 3 g,
   acima do teto de 10 mg/ml das bulas (FDA e brasileira, restrição de volume). Pendente no item 6: HDF.
7. **Vancomicina de ataque fora da HD: 20 mg/kg em 60 min** (equipe 2026-10-31, "adicione essa linha"). Divergências: consenso
   de vancomicina [9] dá 25 mg/kg após a diálise; IDSA 2009 tabela 7 [1] dá 20 mg/kg, mas na última hora da sessão; bula ≤ 10
   mg/min (20 mg/kg × 70 kg = 1,4 g em 60 min = 23 mg/min). Mín. 1 g, máx. 2 g, SF 250 ml (equipe). Manutenção também
   passa a SF 250 ml (equipe). Concentrações: 4-8 mg/ml fora da HD, 4-6 mg/ml na manutenção, dentro do teto da bula.
   **BIC ou microgotas** no tempo de infusão das 3 linhas da vancomicina (equipe 2026-10-31; resolve a BIC do item 6).

## Status do card (2026-10-29)
- **CARD FECHADO 2026-10-30: as 4 folhas aprovadas** (1 cateter + óstio/túnel; 2 FAV e prótese; 3 urina, celulite,
  pneumonia; 4 ficha de drogas + lock, A4 vertical). Próximo: documento teórico com referências e as divergências
  da lista "Fica fora do card".
- **Documento teórico `TEORICO.md`: APROVADO 2026-10-30** (rascunho 3 + levofloxacino; revisão independente aplicada).
  37 referências numeradas. Próximo: números de referência no card (equipe pediu; substitui a regra "sem fonte no card").
- **Folha 2 (FAV e Prótese, folha própria): APROVADA 2026-10-30** ("ok, 2 aprovada"), incluindo quadro "Atenção", faixa
  "qualquer germe, também após a cirurgia" e título "Infecção de FAV e Prótese". Não reabrir.
- **Folha 3 (Urina, Celulite, Pneumonia): APROVADA 2026-10-30** ("Aprovado 3"). Não reabrir. Histórico: Sem caixa de sintomas (decidido 2026-10-30:
  "não precisa de sintomas aqui"). "Grave?" da celulite fica sem critério (decidido 2026-10-30: "não precisa"). Forma: largura inteira, 10 pt
  (proposta minha, em execução).
- **Folha 1 (bacteremia por cateter): APROVADA.** Reaberta só para receber o quadro Óstio e Túnel (pedido da
  equipe 2026-10-29); o resto da folha não muda.
- **Folha 2 (outras infecções), A4 vertical:** quadro 1 **Óstio e Túnel discutido e corrigido** (cateter fica
  na tunelite; sem fio-guia com novo túnel; sem cultura de secreção; Hipafix/Tegaderm; hemocultura no túnel).
  **Quadro 2, FAV e Prótese, em revisão:** (1) alerta da prótese ligado ao ramo, com duração também no ramo da
  cirurgia; (2) *S. aureus* 3 semanas só na FAV nativa, prótese sempre 4 semanas: decididos. Abertos, um por vez: (3) colher hemocultura antes do antibiótico? (não está na discussão); (4) "Doppler
  inicial" fica **só "Doppler inicial"**, sem "pelo menos" (decidido 2026-10-29); (5) "2 semanas" × "14 dias" (só forma).
  Depois: Infecção Urinária, Celulite, Pneumonia.
- **Folha 4 (ficha de drogas + lock): APROVADA 2026-10-30** ("Aprovada 4"), A4 vertical.
  Nota do lock "(impresso no ramo do cateter)" removida 2026-10-30.

## Regras de processo (pedidas pela equipe)
- **Nada entra sem "sim" explícito da equipe para aquele item.** Resposta sobre outro assunto não é aprovação.
  Ao propor, dizer de onde veio (fonte ou "proposta minha").
- **Briefing primeiro, card depois.** Toda mudança de conteúdo entra aqui antes de mexer no card.
- **Conferir contra a discussão original** (`PLANO.md`, seção "Proposta da equipe") antes de cada versão:
  nada da discussão pode sumir do card sem decisão registrada.

## Pendências de conteúdo
- Nenhuma.
