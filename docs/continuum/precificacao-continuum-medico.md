# Precificação Continuum Médico — versão enxuta (LOCK 08/09/2026)

> **Produto novo. O completo continua existindo, mas fica em stand-by.** O **Continuum Plenya**
> (4 profissionais, semestral R$ 37.000 / anual R$ 65.000, `deck-v5.html`) não é retirado nem
> reprecificado; apenas sai de venda por ora. O que vai a mercado é o **Continuum Médico**: o mesmo
> método, conduzido só por médico, sem nutricionista, psicóloga nem educador físico.
>
> **Fonte da estrutura de custo:** [precificacao-mvp.md](precificacao-mvp.md) (planilha travada em
> 24-25/05/2026, Trims #1-4). Toda tarifa, alíquota e rateio abaixo vem de lá; o que muda aqui são
> as quantidades e os blocos removidos.

## Decisões de produto (LOCK 08/09/2026)

| Decisão | Valor | Consequência |
|---|---|---|
| Equipe | **só médico** | some o Bloco A dos outros 3 e as 3 reuniões de equipe |
| Cadência semestral | **6 encontros** | inicial, plano, 2 calls, 2 reavaliações |
| Cadência anual | **12 encontros** | inicial, plano, 6 acompanhamentos, 4 reavaliações trimestrais |
| Box Plenya | **não existe** | Bloco B zera; manipulado e magnésio viram prescrição que o paciente compra |
| WhatsApp | **retainer R$ 200/mês** | metade do Continuum completo, com SLA mais estreito |
| Modalidades | **semestral e anual** | mantém a narrativa dos dois horizontes |
| Formato dos encontros | **online ou presencial em Londrina** | escolha do paciente (LOCK 08/09) |
| Preço | **R$ 12.000 (6× R$ 2.000) · R$ 20.000 (12× R$ 1.667)** | âncora fixada primeiro, planilha ajustada até caber |
| Único programa à venda | **sim** | o Continuum completo sai de venda; a fatia de aquisição do programa passa toda para o Médico |
| Painel genético | **add-on à parte** (Trim #2) | fora do preço |
| Coordenação/concierge | **absorvida** (Trim #4) | Bloco J zera |
| Pró-labore CEO | **fora** (Trim #3) | volta na versão ideal |

## Tarifas e parâmetros herdados

Tarifa MVP (Trim #1): **médico R$ 600 por sessão**, qualquer tipo. Sem reunião de equipe porque não
há equipe multidisciplinar no ciclo. WhatsApp cortado de R$ 400 para **R$ 200/mês** (ver
[Por que o WhatsApp foi cortado](#por-que-o-whatsapp-foi-cortado)).

Alíquotas: carga tributária efetiva **14,33%** (Lucro Presumido, presunção 32%) + gateway Asaas
**2,99%**. Fator de escala `f = 1 − 0,1433 − 0,0299 = **0,8268**`.

Margem efetiva de uma âncora: `margem = (P × 0,8268 − C_fixo) / P`.

## Cronogramas

### Semestral — 26 semanas, 6 encontros

| Sem | O quê |
|---|---|
| 1 | Consulta inicial (anamnese completa + Escore) |
| 2 | Apresentação do plano |
| 9 | Call de acompanhamento |
| 14 | Reavaliação de meio |
| 20 | Call de acompanhamento |
| 26 | Reavaliação final |

### Anual — 52 semanas, 12 encontros

Inicial + apresentação do plano + 6 encontros de acompanhamento + 4 reavaliações trimestrais.
Dá aproximadamente um encontro por mês. Razão anual÷semestral = 2,0× em encontros e 1,62× em
C_fixo, praticamente igual ao 1,64× do Continuum completo.

## C_fixo — cenário base (24 pacientes/ano)

| Bloco | Semestral | Anual | Origem |
|---|---:|---:|---|
| A — Encontros (6 × 600 / 12 × 600) | 3.600 | 7.200 | tarifa MVP |
| A — WhatsApp (R$ 200/mês) | 1.200 | 2.400 | cortado pela metade |
| **A total** | **4.800** | **9.600** | |
| B — Box | **0** | **0** | box removido |
| C — Infra + amortização EMR | 410 | 820 | idem MVP (sem licença) |
| D — Contábil + seguro RC só do médico | 420 | 840 | seguro 3.000/ano × 80% ÷ 24 |
| E — Aquisição (CAC-MVP) | 3.490 | 3.490 | idem MVP (ver nota sobre os 80% abaixo) |
| F — Jurídico | 293 | 453 | idem MVP |
| J — Extras | **0** | **0** | Trims #2 e #4 |
| I — Risco (5% sobre A-F+J) | 471 | 760 | idem MVP |
| **C_fixo** | **R$ 9.884** | **R$ 15.963** | |

Comparação: o Continuum completo tem C_fixo MVP de R$ 28.878 (semestral) e R$ 50.571 (anual).
O Médico custa **34% do semestral** e **32% do anual**.

> **Sobre os 80% do Bloco E.** Os R$ 3.490 herdados do MVP já são a *fatia do programa* dentro do
> gasto de marketing: 80% vai para o programa, 20% fica com o resto da prática (Consulta Plenya,
> presencial em Londrina), que continua existindo e continua sendo alimentada pelo mesmo tráfego.
> Tirar o Continuum completo de venda não muda esse 80/20; muda apenas qual programa carrega os
> 80%. Por isso o número não é rerrateado aqui. Se um dia a Consulta Plenya sair do funil, aí sim
> E sobe para R$ 4.363 (3.490 ÷ 0,8), o C_fixo semestral vai a ~R$ 10.800 e a âncora de R$ 12.000
> passa a dar prejuízo de ~7%: com margem de R$ 38 por ciclo, esse rateio é o parâmetro mais
> sensível do modelo inteiro.

## Margem nas âncoras travadas

| Volume | Semestral R$ 12.000 | Anual R$ 20.000 |
|---|---:|---:|
| 24 pacientes/ano | C_fixo 9.884 · lucro R$ 38 · **0,3%** | C_fixo 15.963 · lucro R$ 573 · **2,9%** |
| 32 pacientes/ano | C_fixo 8.693 · lucro R$ 1.229 · **10,2%** | C_fixo 14.532 · lucro R$ 2.004 · **10,0%** |
| 40 pacientes/ano | C_fixo 7.976 · lucro R$ 1.946 · **16,2%** | C_fixo 13.669 · lucro R$ 2.867 · **14,3%** |

**Break-even: 24 pacientes/ano no semestral, 22 no anual.** Com a cadência de 12 encontros no
anual (decisão de 08/09), **os dois planos passam a nascer perto do break-even**: o anual, que antes
carregava a rentabilidade do produto com 9,2%, caiu para 2,9%. A margem do negócio deixou de estar
na modalidade e passou a estar inteiramente no volume: aos 32 pacientes/ano os dois voltam a ~10%.

## Por que o WhatsApp foi cortado

O retainer de R$ 400/mês custava R$ 2.400 no ciclo semestral. No Continuum completo isso era 6,5%
de um preço de R$ 37.000. Num produto de R$ 12.000 vira **20% do preço**, e sozinho colocava a
âncora em prejuízo de −10,2% no volume de hoje. Cortar para R$ 200/mês devolve R$ 1.200 ao ciclo
e é o que faz R$ 12.000 fechar sem depender de premissa de volume.

| Cenário, 24 pac/ano | Semestral R$ 12.000 | Anual R$ 20.000 |
|---|---:|---:|
| WhatsApp R$ 400/mês | −10,2% | −9,7% |
| **WhatsApp R$ 200/mês (adotado)** | **0,3%** | **2,9%** |
| Sem retainer | 10,8% | 15,5% |

Contrapartida operacional: SLA mais estreito que o do completo (horário comercial, resposta em até
24h úteis). Urgência clínica nunca foi canal de WhatsApp, e continua não sendo.

## O presencial em Londrina e o custo de sala

O programa deixou de ser 100% online em 08/09/2026: o paciente escolhe encontro online ou
presencial na clínica de Londrina. **A planilha não ganhou linha nova de custo**, e a razão é a
mesma convenção que já rege o Bloco E: a clínica física existe, é usada pela Consulta Plenya e pelo
presencial avulso, e é carregada pelo resto da prática. Um paciente do Continuum que escolhe
presencial ocupa uma sala que já está paga. O honorário do médico (R$ 600) é o mesmo nas duas
modalidades, então o custo marginal real é a hora-sala, não um novo bloco.

**Mas essa é a premissa mais frágil do modelo, e vale saber o tamanho da corda.** Com folga de
R$ 38 por ciclo, o semestral absorve **R$ 6 por encontro** de custo de sala antes de zerar. O anual,
com folga de R$ 573 sobre 12 encontros, absorve **R$ 45 por encontro** (eram R$ 175 quando o anual
tinha 10 encontros).

| Se a sala for cobrada do programa (100% presencial) | Semestral R$ 12.000 | Anual R$ 20.000 |
|---|---:|---:|
| R$ 50/hora | −2,3% | −0,3% |
| R$ 100/hora | −4,9% | −3,4% |

Leitura: **qualquer rateio de sala, por menor que seja, põe os dois planos em prejuízo.** Com 10
encontros o anual tinha colchão; com 12, não tem mais. Se a clínica passar a cobrar hora-sala dos
programas, ou se o presencial virar a escolha da maioria e a capacidade tiver que crescer, os dois
preços têm que ser revistos. Não é problema hoje, é o gatilho.

## Narrativa comercial

- **Dois semestrais custam R$ 24.000; o anual custa R$ 20.000.** O cliente economiza R$ 4.000, ou
  16,7%, e o anual é mais lucrativo para o negócio (aquisição uma vez). No Continuum completo o
  mesmo desconto é de 12,2% (R$ 74.000 contra R$ 65.000), então o anual do Médico é
  proporcionalmente mais vantajoso para quem compra.
- **O Médico custa cerca de um terço do completo** (32% do semestral, 31% do anual). É a porta de
  entrada acessível ao mesmo método, não uma versão diluída dele.

## Riscos assumidos, explicitamente

- **Os dois planos nascem perto do break-even.** A 24 pacientes/ano o semestral dá R$ 38 de lucro
  por ciclo e o anual dá R$ 573. Ambos são ruído. Com a cadência de 12 encontros o anual deixou de
  ser o plano que sustenta o produto, e a margem passou a depender só de volume: aos 32
  pacientes/ano os dois voltam a ~10%.
- **O CAC é o segundo maior custo do produto**, 35% do C_fixo semestral, e continua sendo um
  artefato de volume baixo. Cada paciente a mais no ano derruba o custo de todos os outros.
- **O presencial entra sem custo por premissa, não por medida.** A sala é do resto da prática. Se
  isso mudar, o semestral não tem folga nenhuma para absorver (R$ 6 por encontro).
- **Payback e caixa**: aquisição e onboarding são pagos adiantados, as parcelas entram ao longo de
  6 a 12 meses. Com margem de 0,3% no semestral, não há folga para antecipar recebível (a
  antecipação Asaas custa a partir de 1,25%/mês e comeria a margem inteira).

## Pendências

- Definir se há upgrade do Médico para o completo no meio do ciclo, e como é precificado.
- Cotar apólice de RC real só do médico (hoje é estimativa de R$ 3.000/ano).
- Revisar o rateio de 80% do marketing sempre que a composição do funil mudar, não só se o
  completo voltar. Com margem de R$ 38 no semestral, qualquer rerrateio decide o resultado.
