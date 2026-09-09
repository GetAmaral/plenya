# Continuum Médico · deck-medico.html — plano e estado

> Origem: discussão de 08/09/2026. Derivado de `deck-v5.html` (13 slides, Continuum completo).
> **A v5 fica congelada e continua válida.** Editar só `deck-medico.html`.

## O produto

Versão enxuta do Continuum, conduzida só por médico. Sem nutricionista, psicóloga nem educador
físico, e sem Box Plenya. Planilha e preço em
[docs/continuum/precificacao-continuum-medico.md](../../../docs/continuum/precificacao-continuum-medico.md).

| | Continuum (v5, em stand-by) | Continuum Médico |
|---|---|---|
| Equipe no ciclo | 4 profissionais | só médico (núcleo de 4 nefrologistas) |
| Encontros (semestral) | 29 toques, 1 por semana | 6 |
| Encontros (anual) | 52 toques | 10 |
| Box | 4 boxes/ciclo | não existe |
| Formato dos encontros | 100% online | online ou presencial em Londrina |
| WhatsApp | R$ 400/mês | R$ 200/mês |
| Semestral | R$ 37.000 (6× R$ 6.167) | **R$ 12.000 (6× R$ 2.000)** |
| Anual | R$ 65.000 (12× R$ 5.417) | **R$ 20.000 (12× R$ 1.667)** |

## Sequência (12 slides)

| # | Slide | Vem da v5 | O que mudou |
|---|---|---|---|
| 1 | Capa | s01 | «Continuum» + «MÉDICO» em caps gold; subtítulo cita condução por médico |
| 2 | Para quem não é | s02 | «a equipe propõe» virou «o médico propõe»; WhatsApp sem «grupo» |
| 3 | Para quem é | s03 | só o nome do produto |
| 4 | O médico-gestor | s04 | intacto (virou o coração do produto) |
| 5 | Corpo como um só sistema | s05 | intacto |
| 6 | Método AGIR | s06 | intacto |
| 7 | A jornada | s07 | **reescrito**: 6 encontros por semana datada, sem rotação e sem box |
| 8 | Núcleo médico | **novo** | grid dos 4 retratos de estúdio com CRM e RQE |
| 9 | O escopo | s10 | **reescrito**: coluna «O Box» virou «O registro» (portal do paciente) |
| 10 | Ancoragem de valor | s11 | intacto |
| 11 | Investimento | s12 | **preços novos**; «à vista» virou «total de R$ X no ciclo» |
| 12 | Fechamento | s13 | intacto |

**Cortado da v5:** s08 «Equipe Plenya» (foto dos seis, com nutri e psicóloga) e s09 «Box Plenya».

## Decisões visuais desta rodada

- **Slide 8** usa os retratos individuais, não recorte da foto de grupo: os quatro médicos não são
  adjacentes na `EquipePlenyaEmPe.jpg`, e os retratos já vêm no mesmo fundo petrol da paleta do deck.
  Fontes: `apps/site/public/images/dr-getulio.jpg` e `apps/site/public/images/team/{danilo-ramos-cunha,
  michel-moro-alves,taynara-fratoni}.jpg`. Nomes, CRM e RQE vêm de `apps/site/content/doctors/*.mdx`.
- **Título do slide 8** é canônico do site (`team.doctorsTitle` = «Núcleo médico», caption
  «a mesma escola, a mesma língua»), não copy inventada.
- **Slide 9**: a `Slide18.PNG` mostra um Box Plenya na mesa. Como o box saiu do produto, a imagem é
  ampliada para 152% e ancorada à esquerda, o que corta o box e deixa o notebook com o radar do
  Escore como herói.

## Regras que continuam valendo

Sem preços fora do slide 11, sem marcas comerciais, sem casos clínicos identificáveis, sem
travessão, sem «não é X. é Y.» empilhado, aspas «francesas». Ver `EDITORIAL.md` e as memórias
editoriais.

## Correções da revisão (08/09/2026)

Rodada de `/code-review` sobre o bloco. Aplicadas:

- Slide 7: a linha do tempo estava fora de ordem (semanas 9 e 20, depois «todo o ciclo», depois 14
  e 26), num eixo que o slide apresenta como escada ascendente. Reordenada para 1 · 2 · 9 · 14 ·
  20 e 26, com o WhatsApp movido para o rodapé, onde já estava a nota do anual.
- Slide 9: «Reavaliação do Escore no meio e no fim do ciclo» contradizia a cadência trimestral do
  anual, numa lista que é explicitamente bimodal. Agora diz «no meio e no fim do semestral, a cada
  trimestre no anual».
- Slide 11: `12 × R$ 1.667 = R$ 20.004`, não R$ 20.000. A parcela virou **R$ 1.666,67**, que fecha
  exato. O semestral já fechava (6 × R$ 2.000).

Uma constatação da revisão foi **rejeitada**: a de que o Bloco E deveria subir para R$ 4.363 por
ser «100% deste produto». Os 80% do MVP são a fatia do *programa* contra o resto da prática
(Consulta Plenya, presencial), não um rateio entre produtos do Continuum. O resto da prática segue
existindo, e o próprio slide 2 vende a Consulta Plenya. O número fica; a redação ambígua da
planilha foi corrigida.

## Rodada 2 (08/09/2026) — presencial e fechamento das pendências

- **Encontros passam a ser online ou presenciais em Londrina**, à escolha do paciente. Mexeu em
  quatro pontos: a capa (que dizia «100% online», agora em duas linhas para caber na parede escura
  da imagem), a jornada («Consulta inicial online» virou «Consulta inicial»), o escopo (ganhou a
  nota «online ou presenciais em Londrina, como preferir») e, principalmente, o «Fica de fora», de
  onde **«Atendimento presencial em Londrina» foi removido**, porque agora está dentro.
- Aproveitando o aperto de espaço no escopo: «Reavaliação do Escore no meio e no fim do semestral,
  a cada trimestre no anual» virou **«Reavaliação trimestral do Escore»**, que é mais curto e igual
  de correto: as reavaliações do semestral caem no 3º e no 6º mês, ou seja, já são trimestrais.
- **«Doze vagas a cada janela trimestral» confirmado.** Pendência fechada.
- **Escada de ancoragem mantida** em R$ 100/200/300 mil. Pendência fechada.

**Consequência de custo do presencial, registrada na planilha:** nenhuma linha nova, porque a
clínica já é carregada pelo resto da prática. Mas o semestral absorve só **R$ 6 por encontro** de
hora-sala antes de zerar (o anual absorve R$ 175). Se a sala um dia for cobrada do programa, o
semestral precisa ser reprecificado.

## Pendências

- Confirmar com o Getúlio a cadência exata dos 10 encontros do anual.
- Definir se o presencial tem limite (quantos encontros do ciclo podem ser presenciais) ou se é
  livre. Hoje o deck diz «como preferir», sem teto.
- Deploy do PDF pro VPS (`decks.plenyasaude.com.br`) só sob ordem.

## Como gerar

```bash
node render-pngs.js --file=deck-medico.html --out=previews-medico [--slide=NN]
node render.js --file=deck-medico.html --out=../../../docs/decks/continuum-medico-AAAAMMDD.pdf
```
