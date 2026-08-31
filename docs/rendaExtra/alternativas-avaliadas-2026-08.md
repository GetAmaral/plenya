# Renda extra de R$ 5.000 a R$ 10.000/mês
### Estudo de opções — agosto/2026

> **STATUS: arquivo de alternativas avaliadas e não escolhidas.**
> A rota da Parte 4 (trabalho de especialista médico para laboratórios de IA, US$ 130-200/h)
> fica documentada aqui como **plano B de caixa rápido**. Veredito do usuário em 05/08/2026:
> "não foge muito de um emprego". Correto. Ela troca hora por dinheiro, não constrói ativo e
> não tem valor de venda. Serve para financiar o negócio real, não para ser o negócio.
> O estudo do negócio de verdade está em [negocio-real-2026-08.md](negocio-real-2026-08.md).
Restrições do briefing: baixo risco, remoto, automatizável por IA/LLM, teto de 30 min/dia de
atenção humana, sem relação com a atividade médica, qualquer setor, legal no Brasil.

---

## 1. O veredito, antes dos detalhes

Existem exatamente três máquinas que produzem R$ 5-10 mil/mês com 30 minutos de atenção
diária. Todas as outras respostas que aparecem na internet são disfarce de emprego.

1. **Capital rendendo.** Zero minuto por dia, risco realmente baixo, mas exige de R$ 490 mil a
   R$ 1,06 milhão parados.
2. **Ativo digital com receita já existente, comprado pronto.** Mesmo resultado por
   R$ 150-250 mil, com 20-40 min/dia de operação. É a rota com melhor relação capital/renda
   disponível hoje.
3. **Produto de software próprio, self-serve, em nicho regulatório brasileiro.** Capital
   baixo (R$ 300-800/mês de infra), mas leva de 9 a 18 meses e depende de você ter uma
   vantagem injusta. Você tem.

Serviço (agência, automação, consultoria) entrega dinheiro rápido e é a rota mais recomendada
por aí, mas viola frontalmente o teto de 30 min/dia enquanto você for a pessoa que vende e que
atende. Só entra no portfólio com um sócio comercial.

---

## 2. A régua: quanto custa comprar R$ 5.000/mês

Com Selic 14,25% e CDI 14,15% a.a. (agosto/2026), o CDI rende **1,11% ao mês bruto**,
**0,94% líquido** depois de IR de 15% (prazo acima de 2 anos). LCI/LCA seguem isentas de IR
(a MP 1.303/2025 caducou e a tributação de 5% não passou); a 92% do CDI isento, dão cerca de
**1,02% líquido ao mês**.

| Rota | Capital para R$ 5k/mês | Capital para R$ 10k/mês | Yield anual implícito | Tempo/dia | Risco |
|---|---|---|---|---|---|
| CDB/Tesouro pós-fixado | R$ 530 mil | R$ 1,06 milhão | ~11,3% líq. | 0 min | muito baixo |
| LCI/LCA (isentas) | R$ 490 mil | R$ 980 mil | ~12,3% líq. | 0 min | baixo (FGC) |
| FIIs de papel/tijolo | R$ 550-650 mil | R$ 1,1-1,3 mi | 9-11% isento | 5 min | médio (cota oscila) |
| **Micro-SaaS comprado (2,5-4x lucro anual)** | **R$ 150-240 mil** | **R$ 300-480 mil** | **25-40%** | 20-40 min | médio |
| Site de conteúdo comprado (2,6x lucro) | R$ 130-190 mil | R$ 260-380 mil | 30-45% | 15 min | alto (busca por IA) |
| Lavanderia self-service | R$ 120-300 mil | não escala assim | 15-30% | 30-60 min presencial | médio-alto |
| Micro-SaaS construído | R$ 5-10 mil + 9-18 meses | idem | — | 30 min | alto na largada |

A leitura importante da tabela: **o mercado paga 3 a 4 vezes menos capital por um real de
lucro operacional do que por um real de juros**. Esse desconto é o preço do risco operacional.
Para quem sabe auditar um sistema (você sabe), esse é o arbitramento mais eficiente disponível.

---

## 3. As rotas que sobrevivem à restrição

### Rota A — Capital para cashflow
**Renda:** proporcional ao capital. **Prazo:** imediato. **Tempo:** zero.

Não é resposta empreendedora, mas é a única resposta honestamente de "risco baixo". Se você
já tem patrimônio parado ou mal alocado (conta corrente, poupança, fundo caro, previdência
ruim), converter isso em LCI/LCA e Tesouro Selic já produz uma fração relevante da meta sem
nenhum trabalho. R$ 300 mil realocados = cerca de R$ 3 mil/mês.

Serve como **piso** do portfólio, não como plano.

### Rota B — Comprar um ativo digital que já fatura
**Renda:** R$ 5-10 mil/mês desde o mês 1. **Prazo:** 60-90 dias até fechar a compra.
**Tempo:** 20-40 min/dia. **Capital:** R$ 150-250 mil por R$ 5 mil/mês.

Dados de mercado 2026: micro-SaaS abaixo de US$ 1M ARR negocia a **2,85x lucro anual na
média** (topo de quartil 6,1x); em deals pequenos a faixa praticada é **2,5x a 4,5x SDE** no
Flippa e mediana de 3,9x no Acquire.com. Sites de conteúdo: 2,6x. Com dólar a R$ 5,13, um
SaaS de US$ 1.000 MRR sai por US$ 30-45 mil, ou seja R$ 155-230 mil.

Por que isso encaixa em você especificamente: a parte cara de comprar software é a auditoria
técnica, e você faz isso sozinho com Claude Code em um fim de semana (ler o repo, medir dívida
técnica, checar se o custo de infra declarado bate, verificar se dá para operar sem o
fundador). Quem compra sem saber ler código paga consultor ou compra gato por lebre.

**Checklist de due diligence (não negociável):**
- Acesso *read-only* ao Stripe/Paddle/Asaas, não a planilha nem a print. Receita dos últimos
  24 meses, mês a mês.
- Churn mensal abaixo de 5%. Acima de 8%, o negócio é uma esteira de vendas disfarçada de SaaS.
- Idade mínima de 18 meses e nenhuma concentração acima de 20% da receita em um cliente.
- Origem do tráfego: se mais de 60% vem de busca orgânica do Google, desconte pesado. As
  respostas geradas por IA derrubaram tráfego de conteúdo e esse risco ainda está em curso.
- Nada de dependência de uma única API que pode mudar preço ou política (aprendi isso caro com
  Composio/Meta na Plenya).
- Escrow obrigatório, 30 dias de transição contratados, preço no máximo 3,5x lucro anual.
- Prefira B2B com cartão recorrente e produto "chato" (conformidade, faturamento, backup,
  monitoramento). Produto chato tem churn baixo e cliente que não some.

**Onde procurar:** Acquire.com (melhor curadoria, US$ 500M+ transacionados), Flippa (volume
maior, curadoria pior, transações de SaaS subiram 73,5% em 2025), Empire Flippers (ticket
maior, diligência já feita, preço maior). No Brasil não existe marketplace equivalente maduro;
o mercado local acontece em comunidades fechadas e no boca a boca, com múltiplos menores e
risco documental maior.

**Risco real:** métrica inflada, dependência do fundador, plataforma que muda regra. Mitiga-se
com o checklist acima e comprando dois ativos de R$ 100 mil em vez de um de R$ 200 mil.

### Rota C — Construir um micro-SaaS regulatório brasileiro
**Renda:** R$ 5-10 mil/mês no mês 9-18. **Capital:** desprezível. **Tempo:** 30 min/dia é
viável de verdade, porque o trabalho pesado é do agente.

Estatística de base, para calibrar expectativa: a mediana até US$ 10k MRR é de 12 a 18 meses
depois do primeiro cliente pagante, e só cerca de 6% dos micro-SaaS chegam lá. A faixa
realista de primeiro alvo é US$ 1-5k MRR. Média geral: US$ 1.735 MRR com 64% de margem.

O que muda o jogo a favor de um produto brasileiro: **regulação nova cria dor com data
marcada, e solução gringa não resolve.** A janela aberta agora:

- Desde **03/08/2026**, empresas do regime regular não conseguem emitir documento fiscal
  eletrônico sem preencher os campos de IBS e CBS (alíquota teste de 1%: 0,1% IBS + 0,9% CBS).
- **Empresas do Simples Nacional só entram em 2027.** São mais de 20 milhões de CNPJs que vão
  bater na parede em janeiro, com contador sobrecarregado e ERP caro demais.
- Isso te dá 12 meses exatos para construir e distribuir antes da onda.

Ideias que se encaixam no perfil "chato, recorrente, regulatório, suporte baixo":
- Validador/conferidor de conformidade de NF-e para CBS e IBS, com alerta de erro de apuração.
- Calendário fiscal inteligente do Simples (DAS, PGDAS, DEFIS, eSocial) com semáforo e aviso
  por WhatsApp antes do vencimento. Multa por atraso chega a R$ 500/mês, então o produto se
  paga sozinho a R$ 39/mês.
- Simulador de transição MEI para ME (14 milhões de MEIs, teto de R$ 81 mil, ninguém sabe se
  compensa Simples, Presumido ou Real).
- Gerador de política de privacidade e registro de tratamento de dados (LGPD vale para MEI e
  não tem teto de faturamento) por R$ 15-30/mês.

**Sua vantagem injusta, que quase ninguém tem:** você já tem VPS + Coolify rodando, WhatsApp
Cloud API homologada em produção com número verificado e templates aprovados pela Meta (a
parte que trava 90% dos concorrentes por semanas), Go + Next + Postgres em produção, e um
fluxo de trabalho com agente que produz software sem consumir suas horas. Custo marginal de
lançar mais um produto no seu stack é próximo de zero: VPS de R$ 28-55/mês e R$ 30-80/mês de
API de LLM.

**Nota lateral, você decide se conta como "médico":** o EMR que você já construiu contém, de
graça, um micro-SaaS pronto para 547 mil psicólogos autônomos (agenda + recibo digital do
Receita Saúde + Carnê-Leão + IRPF, obrigatório desde 2025 e sem produto integrado no mercado).
Não é exercício da medicina, é licenciamento de software, e o custo marginal para você é
próximo de zero. É provavelmente a maior assimetria da lista inteira, mesmo violando a
preferência de "nada de saúde".

### Rota D — Serviço produtizado de automação (só com sócio comercial)
**Renda:** R$ 5 mil/mês em 60-90 dias. **Tempo:** 2h/dia se você fizer sozinho, 20 min/dia se
alguém vender e atender.

Preços praticados no Brasil em 2026: automação simples R$ 400-900, intermediária
R$ 900-2.500, e a margem de verdade está na **manutenção mensal**. Casos publicados mostram
agentes de IA recorrentes em torno de R$ 2.450/mês por cliente. Custo: VPS R$ 28-55/mês e
R$ 30-80/mês de API por cliente de pequeno porte. Margem bruta acima de 90%.

Aritmética da meta: 6 a 10 clientes a R$ 500-900/mês de retainer, ou 3 clientes a R$ 1.500.
Tecnicamente é trivial para você e a demanda existe (menos de 5% das 24 milhões de empresas
ativas no Brasil usam qualquer SaaS; 1,2% têm qualquer otimização para IA generativa).

O problema não é técnico, é humano: prospecção, reunião, escopo, cliente que liga no domingo.
**Só entra no portfólio se você achar um sócio comercial** que venda e atenda por 40-50% da
receita, com você entregando a parte técnica em blocos assíncronos. Nessa configuração é a
rota mais rápida da lista.

---

## 4. O que descartar, com o motivo numérico

| Opção | Por que não |
|---|---|
| Canal "dark"/faceless no YouTube com IA | Política de conteúdo inautêntico desde 15/07/2025 desmonetiza produção em massa e templada. Só 3% dos canais chegam a monetizar. Não é renda, é loteria com custo de produção. |
| KDP low-content / ebooks gerados por IA | Mercado maduro, publicação em massa com variação mínima gera remoção e restrição de conta inteira. Ticket baixíssimo, precisaria de centenas de títulos vivos. |
| Afiliados e sites de nicho novos por SEO | Respostas geradas por IA no topo da busca cortaram o clique. Construir tráfego orgânico do zero em 2026 é a definição oposta de risco baixo. Comprar site já ranqueado com desconto é outra história. |
| Dropshipping / e-commerce | SAC, logística reversa, gestão de anúncios. Consome 3-4h/dia e margem líquida de 5-15% com capital de giro travado. |
| Lavanderia self-service, vending | R$ 120-300 mil de CAPEX para faturar R$ 15-25 mil e lucrar talvez R$ 4-7 mil, payback de 18-36 meses. Não é remoto e não é 30 min/dia (máquina quebra, ponto, vandalismo, inadimplência de aluguel). Pelo mesmo dinheiro, um SaaS comprado rende mais e não tem endereço. |
| Co-host de Airbnb | R$ 1,2-3,6 mil/mês por muita mensagem humana em horário ruim. Não chega na meta e não cabe no tempo. |
| Trading, cripto, arbitragem, "robô" | Não é renda de risco baixo, é exposição alavancada com nome bonito. |
| Mentoria/consultoria genérica de IA | Funciona, mas é venda do seu tempo. Mesmo problema da Rota D, sem o ativo no fim. |

---

## 5. Portfólio recomendado

Não escolha uma rota. Monte três camadas, porque elas têm curvas de risco e de prazo
diferentes e se financiam entre si.

**Camada 1 — piso (mês 1, 0 min/dia).**
Auditar o que já existe de capital e realocar para LCI/LCA e Tesouro Selic. Cada R$ 100 mil
bem alocados são R$ 1.000/mês. Isso não é o plano, é a linha de base contra a qual todo o
resto precisa se justificar.

**Camada 2 — núcleo (mês 1-3, 20-40 min/dia).**
Comprar um ativo digital com receita comprovada, na faixa de R$ 120-250 mil, seguindo o
checklist da Rota B. Meta: R$ 4-6 mil/mês de lucro operando com o produto no piloto
automático. Se o capital permitir, dois ativos menores em vez de um grande.

**Camada 3 — assimetria (mês 1-12, 30 min/dia).**
Um micro-SaaS regulatório próprio mirando a virada do Simples Nacional em janeiro/2027,
construído no stack que você já opera. Investimento: R$ 300-800/mês de infra e meia hora
diária de direção de agente. Se der certo, é o ativo que vale 3-4x o lucro anual daqui a dois
anos, ou seja, vira patrimônio, não só renda.

**Cronograma realista da meta:**
- Mês 1: camada 1 no ar. Renda dependendo do capital realocado.
- Mês 3: camada 2 fechada. R$ 4-6 mil/mês.
- Mês 12: camada 3 entre R$ 1-4 mil/mês. Total na faixa de R$ 6-10 mil.
- Mês 24: camada 3 madura pode sozinha passar a meta, e as camadas 1 e 2 seguem.

Sem capital disponível, o cronograma muda: Rota D com sócio comercial no mês 2-3 para gerar o
caixa, Rota C em paralelo, e a Rota B só quando houver R$ 120 mil acumulados.

---

## 6. Estrutura jurídica e tributária

- Abrir **PJ separada da PJ médica**. CNAE de desenvolvimento de software (62.01-5) ou de
  licenciamento (62.03-1). Simples Nacional Anexo III (com fator R acima de 28%) fica em torno
  de 6% a 11,2% de carga na faixa relevante, contra até 27,5% de IRPF como pessoa física.
- Receita recorrente de software é o melhor caso possível para o fator R, porque pró-labore já
  é despesa que você teria.
- Se comprar ativo internacional que fatura em dólar via Stripe/Paddle, avaliar recebimento
  como exportação de serviço (isenta de ISS, e há regime tributário mais leve). Isso precisa
  de contador com experiência em software, não o contador da clínica.
- Não misturar caixa com a Plenya. Ativo pessoal, CNPJ próprio, conta própria.
- Reforma tributária: em 2026 o Simples não tem alteração, a mudança entra em 2027. Vale
  revisar o enquadramento com o contador ao longo de 2026.

---

## 7. Primeiros passos (semana 1)

1. Definir o capital efetivamente disponível e sem uso, porque essa única variável decide entre
   Rota B (comprar) e Rota D (vender serviço) como motor principal.
2. Criar conta em Acquire.com e Flippa, aplicar filtros (SaaS, B2B, MRR entre US$ 800 e
   US$ 2.000, idade acima de 18 meses, múltiplo abaixo de 4x) e acompanhar o fluxo de listagens
   por 30 dias antes de dar qualquer lance. Ver 50 negócios antes de avaliar 1 a sério.
3. Escolher **um** nicho regulatório da Rota C e validar em uma semana: falar com 5 contadores
   sobre o que mais dói na virada de 2027 para os clientes do Simples.
4. Marcar conversa com contador de software para desenhar a PJ.

---

---

# PARTE 2 — Validação dos nichos (05/08/2026)

Contexto: capital disponível abaixo de R$ 50 mil, então a Rota B (comprar ativo pronto) sai de
cena por ora e o motor vira construção. Antes de escolher o nicho, fui checar concorrência real
de cada um. **Os três candidatos da Parte 1 morreram na checagem.** Vale mais matar três ideias
em duas horas de pesquisa do que descobrir isso no mês 8.

## 2.1 Psicólogos autônomos: oceano vermelho, e minha premissa estava errada

A premissa da Parte 1 era "nenhum app integra agenda + recibo do Receita Saúde + Carnê-Leão".
**É falso.** Preços públicos hoje:

| Produto | Preço | Observação |
|---|---|---|
| PsiSync | R$ 49/mês (R$ 39,20 no anual) | agenda, prontuário, NFS-e automática, Pix |
| Hotina | R$ 89,90 e R$ 149,90/mês | **emissão via Receita Saúde já no plano superior**, lembrete de pagamento por WhatsApp, Google Calendar |
| PsicoManager | plano individual para autônomo | player estabelecido, app próprio nas lojas |
| Sinappsy, Menta | — | Menta já vende automação com IA |
| iClinic, Ninsaúde, Doctoralia | preço sob consulta | 33 mil profissionais no iClinic |

Ou seja: 6 ou mais concorrentes, o diferencial que eu tinha imaginado já é feature paga de um
deles, e a âncora de preço do mercado é R$ 49 a R$ 150. Para tirar R$ 5.000/mês seriam de
**50 a 100 assinantes ativos**, conquistados um a um contra empresas com máquina de conteúdo
rodando. Isso é marketing em tempo integral, não 30 minutos por dia. **Descartado.**

## 2.2 Calendário fiscal do Simples e app para MEI: comprador errado

O produto seria bom, mas quem sente a dor não é quem paga:

- O **App MEI oficial e o PGMEI emitem o DAS de graça**, e as contas digitais PJ (Jota,
  InfinitePay, Nubank e afins) já entregam DAS pelo app como isca de aquisição.
- Contabilidade online completa custa **R$ 137/mês** (Contabilizei, com 50 a 70 mil clientes,
  Contajá, Agilize). Quem tem dor de prazo já resolve pagando isso.
- São 15 a 16 milhões de MEIs, mas o MEI é o pior pagador de SaaS do país. Ele não paga
  R$ 39/mês por um lembrete que o calendário do celular dá de graça.

E a versão B2B (vender para o escritório contábil, que é quem tem orçamento) já tem fila de
incumbentes: Acessórias, Pasta Contábil, Digiliza, Jettax, Qive, e-Auditoria, além de
SocialHub e BeeMessage no recorte de WhatsApp. Mercado atendido, venda consultiva, ciclo longo.
**Descartado para operação solo com 30 min/dia.**

## 2.3 LGPD para MEI: a regulação isenta justamente esse público

A Resolução CD/ANPD nº 2/2022 **dispensa ME, EPP, MEI e startups de nomear DPO**. E a
fiscalização de 2026, agora com a ANPD virada em agência com autonomia e 200 fiscais de
carreira, tem alvo declarado: empresas que usam **IA generativa, dados biométricos ou dados de
menores**. A menor multa já aplicada a microempresa foi de cerca de R$ 14 mil, em caso isolado.

Tradução: o MEI não tem obrigação prática, não tem medo e não vai pagar. Quem tem medo de
verdade é empresa média que roda IA ou biometria, e para essa o produto não é um wizard de
R$ 29, é assessoria jurídica. **Descartado.**

## 2.4 O critério que sobrou da autópsia

As três morreram pelo mesmo motivo, e isso vira regra para a próxima escolha:

> **Ticket abaixo de R$ 150/mês exige uma máquina de marketing. Você não tem tempo para operar
> uma. Logo, só serve produto de R$ 300+/mês (10 a 17 clientes bate a meta) ou venda self-serve
> em dólar (20 clientes a US$ 50 = R$ 5.100).**

Segundo critério, igualmente duro: **suporte em WhatsApp às 21h mata o teto de 30 minutos.**
Público brasileiro de PME espera resposta imediata em canal síncrono. Público global paga por
cartão e abre ticket por e-mail, que você responde quando quiser.

## 2.5 Reshortlist (candidatos que passam nos dois critérios)

Nenhum destes está validado ainda, é o próximo lote a passar pela mesma autópsia:

1. **Utilitário B2B global, self-serve, em dólar.** Nichos com margem e churn conhecidos:
   recuperação de pagamento e dunning (70-90% de margem, receita atrelada à do cliente),
   integrações entre ferramentas de nicho, conformidade vertical. Meta: 20 clientes a US$ 50.
   Suporte assíncrono em inglês, zero venda por telefone. É o que melhor cabe em 30 min/dia.
2. **Agente de WhatsApp homologado, vertical, ticket alto.** Sua infra Meta já verificada com
   templates aprovados é a barreira que trava o concorrente por semanas. R$ 500-900/mês x 10
   clientes. Só funciona com sócio comercial vendendo e atendendo.
3. **Licenciamento de módulo que você já construiu.** Custo marginal próximo de zero, mas
   esbarra na sua restrição de não misturar com saúde.

---

# PARTE 3 — E-commerce sem estoque e infoprodutos

Duas rotas que você levantou e que eu tinha tratado rápido demais na Parte 1.

## 3.1 E-commerce sem estoque (dropshipping e print on demand)

**Reprovado nas duas restrições ao mesmo tempo, e por margem larga.**

A aritmética é implacável. Margem líquida real, depois de anúncios, taxas de plataforma,
gateway e devolução, fica entre **5% e 25%** (fornecedor nacional puxa para 10-30% de margem
bruta, o que é mais saudável operacionalmente e pior no bolso). Print on demand fica em
**10% a 30%**, e fornecedor internacional como Printful ainda come o resultado com custo
alfandegário.

Para tirar **R$ 5.000 de lucro líquido**, você precisa faturar entre **R$ 25 mil e R$ 50 mil por
mês**. Isso são centenas de pedidos, e cada pedido traz rastreio, atraso, troca, chargeback,
reclamação em marketplace e SAC. Não existe versão de 30 minutos por dia disso.

Piorou em 2026, especificamente:
- O CAC de e-commerce subiu de 40% a 60% em três anos, com Mercado Livre e Shopee comprando
  mídia agressivamente e empurrando o CPM para cima.
- Desde janeiro de 2026, Google e Meta repassam na fatura do anunciante cerca de **12,15% de
  tributo** (PIS/COFINS 9,25% + ISS ~2,9%). Uma campanha que dava ROAS 4:1 precisa agora de
  4,48:1 só para empatar com o ano passado. O ROAS mediano de mercado está em 1,93x.
- O padrão relatado é **prejuízo nos primeiros 2 a 3 meses** na fase de validação de produto e
  criativo, ou seja, queima de capital antes de qualquer receita, justamente o que você não tem.

Resumo: capital de giro travado, margem fina, atenção humana diária alta, risco de banimento de
conta de anúncio e dependência total de plataforma de mídia. É o oposto exato do briefing.

## 3.2 Infoprodutos

**Aprovado com ressalva, e é a rota mais rápida da lista para quem tem menos de R$ 50 mil.**

Números do mercado brasileiro em 2026: produtor iniciante fica em R$ 1-5 mil/mês, produtor
médio em R$ 10-50 mil/mês, e **90% nunca passa de R$ 5 mil/mês**. Taxas: Kiwify 4,99% + R$ 0,50
por venda, Hotmart 9,9% + R$ 1. A faixa de R$ 5-30 mil/mês em 12 meses é descrita como realista
para produto sólido, e essa estatística é honesta: metade do jogo é ter público.

A conta com ticket de R$ 497 e taxa da Kiwify: sobram cerca de R$ 472 por venda, então
**11 vendas por mês batem R$ 5.000**. Onze. Compare com as 100 assinaturas de R$ 49 do SaaS de
psicólogo ou com os 400 pedidos do dropshipping. É a rota com a menor quantidade de eventos
necessários para atingir a meta, e é por isso que ela merece atenção apesar da fama.

**A diferença entre os 90% que travam em R$ 5 mil e os que passam é uma só: audiência.** Quem
depende de tráfego pago compra atenção a CAC inflado (mesmo repasse de 12,15%, mesmo ROAS
mediano de 1,93x) e trabalha para o Meta. Quem tem público próprio vende com margem de 95%.

**Onde isso te encaixa, e onde não encaixa.** Você tem duas audiências, e elas são de saúde. Um
infoproduto para elas seria, na prática, extensão da sua atividade médica, que é justamente o
que o briefing exclui. Só que você tem um segundo ativo, esse sim fora da medicina e bem mais
raro: **você construiu sozinho, dirigindo agentes de IA, um EMR em produção com 74 models,
WhatsApp API homologada, migrations, deploy e apps mobile.** Prova social verificável, que quase
nenhum criador de conteúdo de IA no Brasil consegue mostrar, num mercado onde vendem curso de
"vibe coding" com muito menos lastro (Hora de Codar, Scopphu, FindSkill e outros já disputam
essa faixa).

O produto natural é o profissional liberal ou dono de PME que quer sistema próprio em vez de
alugar SaaS caro para sempre. Ticket de R$ 500 a R$ 2.000, 3 a 10 vendas por mês, distribuição
por LinkedIn e YouTube com o material que você já produz.

**A ressalva, que é séria:** infoproduto não é 30 min/dia no começo. Gravar, editar, montar
página, checkout, suporte de aluno e atualização quando a ferramenta muda (e o Claude Code muda
todo mês) custam de 1 a 2 horas por dia durante 60 a 90 dias. Depois disso cai para perto de
30 minutos, se o produto for perene e não por turma. É um sprint inicial, não uma maratona, mas
é um sprint que fura o teto que você definiu.

## 3.3 Comparação final das rotas viáveis com capital abaixo de R$ 50 mil

| Rota | Eventos para R$ 5k/mês | Prazo | Esforço no pico | Depois de pronto | Risco |
|---|---|---|---|---|---|
| Infoproduto técnico (fora de saúde) | 11 vendas/mês a R$ 497 | 90-120 dias | 1-2 h/dia por 3 meses | 30 min/dia | médio (depende de você aparecer) |
| Micro-SaaS B2B em dólar | 20 clientes a US$ 50 | 9-18 meses | 30 min/dia | 20 min/dia | alto na largada, vira patrimônio |
| Agente WhatsApp vertical | 10 clientes a R$ 600 | 60-90 dias | alto (venda) | 20 min/dia se houver sócio | médio |
| Dropshipping / POD | 300-500 pedidos/mês | 6+ meses | 3-4 h/dia sempre | 3-4 h/dia sempre | alto, queima capital |
| SaaS para psicólogo | 100 assinaturas a R$ 49 | 12-24 meses | marketing integral | marketing integral | alto, mercado saturado |

**Recomendação:** infoproduto técnico como motor de caixa nos primeiros 120 dias, aceitando o
sprint inicial acima do teto, e micro-SaaS B2B em dólar rodando em paralelo a 30 min/dia como
ativo de longo prazo. As duas usam a mesma matéria-prima (você dirigindo agentes) e uma financia
a outra. Dropshipping e os três nichos da Parte 1 saem do mapa.

---

# PARTE 4 — Mercado global, com saúde de volta à mesa (05/08/2026)

Correção de escopo: saúde nunca esteve excluída, só não podia ser a única lente. E não existe
barreira de idioma. Isso reabre a maior assimetria do seu perfil, que as Partes 1 a 3 estavam
proibidas de usar.

**O que você é, em termos de mercado global:** nefrologista com registro ativo, autor publicado,
que dirige agentes de IA para construir e operar software em produção, fala português nativo e
opera em qualquer idioma. Cada uma dessas coisas isolada é comum. A interseção das três é rara o
suficiente para ter preço próprio em dólar, e em 2026 existe um mercado comprando exatamente
isso.

## 4.1 A descoberta central: o gargalo do mundo é gente como você

Os laboratórios de IA estão limitados por **oferta de especialistas verificados**, não por
dinheiro. Preços praticados em 2026 para médico com credencial ativa:

| Plataforma | Faixa para médico | Observações |
|---|---|---|
| **Mercor** | US$ 110-250/h (redes de médicos); US$ 130-180/h em clínica geral e emergência | tem vaga específica **"Healthcare Expert — International Physicians"**; pagamento semanal por Stripe ou **Wise, e o Brasil está na lista de países suportados** |
| Handshake AI | até US$ 125-300/h | **exige M.D., D.O. ou PhD de instituição dos EUA. Você está fora.** |
| Surge AI | US$ 120-200/h | credencial ativa |
| Outlier (Scale) | US$ 60-95/h em domínio médico e jurídico | faixa mais baixa |
| DataAnnotation | a partir de US$ 50/h | porta de entrada |

Racional econômico declarado pelo próprio mercado: há cerca de 1,4 milhão de médicos em
atividade nos EUA e talvez 1% tem interesse em trabalho paralelo. A escassez define o preço, e
a segurança clínica obriga os laboratórios a investir pesado em avaliação por especialista
verificado.

**A aritmética contra a sua meta, que é o ponto:**

| Meta | Em dólar (câmbio 5,13) | A US$ 130/h | A US$ 180/h |
|---|---|---|---|
| R$ 5.000/mês | US$ 975 | 7,5 h/mês = **22 min/dia** | 5,4 h/mês = 16 min/dia |
| R$ 10.000/mês | US$ 1.950 | 15 h/mês = **30 min/dia** | 11 h/mês = 22 min/dia |

Isso não é uma aproximação do briefing, é o briefing atendido ao pé da letra: remoto,
assíncrono, sem hora fixa, sem capital, sem risco patrimonial, pagamento semanal em moeda forte,
totalmente legal. O risco existente é de disponibilidade de projeto, não de perder dinheiro.

**Seu diferencial dentro desse mercado, que empurra você para o topo da faixa e não para o piso:**

1. **Idioma.** A demanda de IA saiu da era só-inglês, e **português lidera a demanda não inglesa
   com 16,5%** em serviços de transcrição e anotação. Modelos multilíngues reconhecidamente
   perdem para modelos monolíngues em nuance clínica e terminológica em português. Médico
   brasileiro com credencial ativa é exatamente o insumo escasso.
2. **Código.** Avaliação de agentes, revisão de trajetória e red teaming clínico exigem alguém
   que entenda o raciocínio médico e o comportamento do sistema. Você faz as duas leituras.
   Priorize vagas que pedem as duas, porque é onde a faixa vai a US$ 200+.
3. **Nefrologia.** Especialidade com literatura densa e alto risco clínico, que é onde os
   laboratórios mais precisam de avaliação fina.

**Complementos assíncronos, para dias de fila vazia:** M3 Global Research (2 milhões de médicos,
US$ 20-150 por questionário e US$ 150-500 por entrevista em vídeo, atuação internacional), Sermo
(1,3 milhão de profissionais, 96 especialidades, US$ 25-200+ por questionário). ZoomRx é só
EUA. Redes de especialistas como Guidepoint pagam US$ 500-700/h para especialista clínico, mas
o fluxo para médico fora dos EUA é menor e irregular, então trate como bônus, não como base.

## 4.2 O salto: deixar de vender a própria hora

Vender a própria hora resolve o caixa e não constrói patrimônio. Só que a mesma pesquisa aponta
o caminho de saída, e ele é bom.

**O mercado está com escassez de oferta, e você tem acesso direto à oferta.** Você não conhece
um médico brasileiro, você tem audiência de médicos brasileiros no LinkedIn e no Instagram, mais
rede profissional real. E existe arbitragem legítima: a hora de um médico brasileiro custa uma
fração da hora de um médico americano, enquanto o laboratório paga pela credencial e pela
qualidade da avaliação, não pelo CEP.

**Modelo:** montar um pod de especialistas médicos lusófonos e hispanofalantes, com curadoria e
controle de qualidade seus, vendido a laboratórios e a fornecedores de dados que precisam de
cobertura não inglesa. Você recruta, verifica credencial, treina no padrão de avaliação, revisa
amostra e entrega. Margem sobre trabalho que você não executa.

Ordem de grandeza: 5 médicos entregando 20 h/mês cada, com margem de US$ 40/h, são US$ 4.000/mês
(R$ 20 mil) sem você tocar a tarefa. Mesmo uma versão mínima com 3 médicos e 15 h/mês já passa a
meta superior.

**Sequência obrigatória, sem pular etapa:**
1. **Fase 1 (mês 1-2):** entrar como especialista individual na Mercor pela vaga internacional.
   Objetivo duplo: caixa imediato e aprender por dentro o padrão de qualidade, o formato de
   tarefa e quem são os compradores. Ninguém vende um pod sem ter sido a pessoa do pod.
2. **Fase 2 (mês 3-6):** recrutar de 3 a 5 médicos da sua rede, começando por quem já demonstrou
   interesse em IA. Rodar como coletivo dentro das plataformas que permitem, e em paralelo
   prospectar contrato direto com fornecedores de dados que precisam de cobertura em português.
3. **Fase 3 (mês 6-12):** o ferramental que você construir para operar o pod (triagem de
   candidato, verificação de CRM, distribuição de tarefa, amostragem de QA, medição de
   concordância entre avaliadores) é software proprietário com valor de venda. Aí sim vira ativo,
   e nasce do seu problema real, não de um palpite de nicho.

**Riscos, honestamente:** as plataformas podem desintermediar, os laboratórios preferem
fornecedor conhecido e com processo, e contrato direto exige volume e SLA. Por isso a Fase 1 não
é opcional: ela produz caixa enquanto você compra informação sobre o comprador.

## 4.3 A segunda via de moeda forte: sua assinatura clínica sobre trabalho da máquina

Mesmo formato, outro mercado. **Tradução e validação médica para farmacêutica e dispositivo
médico paga US$ 0,20 a 0,35 por palavra**, justamente porque exige revisor com credencial. A
pós-edição de tradução automática paga US$ 0,03 a 0,15, porque assume revisor comum.

Com LLM fazendo a primeira passada e você assinando a revisão como médico, o esforço é de
pós-edição e o preço é de tradução médica. 5.000 palavras por mês a US$ 0,20 são US$ 1.000, ou
seja R$ 5.100, com algo entre 3 e 5 horas de revisão. O canal são os grandes provedores de
localização (RWS Life Sciences, Lionbridge, TransPerfect, Welocalize), que mantêm cadastro
permanente de revisores médicos por par de idiomas, além de validação linguística de
questionários de desfecho reportado por paciente, que exige médico por protocolo.

## 4.4 O que eu checei e NÃO recomendo, para você não repetir meu erro da Parte 2

**Escriba de IA em português está saturado.** Antes de sugerir produto, fui olhar: Voa (feita por
médicos brasileiros), Doclin, Escriba Médico, NAIA e AppHealth no mercado nacional, mais
Abridge, Nabla, Corti e Heidi no internacional, com Commure suportando 90 idiomas e DeepCura 34.
Não há lacuna, há guerra de preço.

Existe **uma** lacuna documentada e ela é técnica, não comercial: um benchmark de julho de 2026
mostra que nenhum dos principais modelos de reconhecimento de fala foi otimizado para fala
médica em espanhol latino-americano. É oportunidade real, mas exige trabalho de modelo e dados,
o que contradiz o teto de 30 minutos. Fica anotado, não recomendado agora.

## 4.5 Estrutura para receber em dólar

Diferença material, não detalhe burocrático:

- **Pessoa física:** rendimento do exterior entra na tabela progressiva do IRPF, 7,5% a 27,5%,
  recolhido mensalmente por Carnê-Leão.
- **PJ com exportação de serviços:** em 2026 há **imunidade de CBS e isenção de IBS** quando o
  serviço é usufruído no exterior, e no Simples a receita declarada como exportação já sai da
  composição de PIS, COFINS e ISS. Teto de R$ 4,8 milhões por ano contando a exportação
  integralmente.
- Entrada de recursos com IOF de 0,38%, e operação de exportação de serviço pode ter tratamento
  específico. Wise, Stripe e similares precisam ficar documentados.
- Portanto: **PJ própria, separada da PJ médica**, com contador que entenda exportação de
  serviço. A diferença de carga paga o contador várias vezes no primeiro ano.

## 4.6 Plano recomendado

| Fase | O que | Renda | Tempo |
|---|---|---|---|
| Semana 1-2 | Candidatura na Mercor (vaga internacional de médico) e cadastro em M3 e Sermo. Abrir conta Wise. | — | 3 h no total |
| Mês 1-2 | Executar como especialista. Subir de faixa buscando tarefa que exija médico **e** leitura de código. | R$ 4-10 mil | 20-30 min/dia |
| Mês 2-3 | Abrir PJ de exportação. Cadastro como revisor médico em 3 provedores de localização. | +R$ 3-5 mil | 3-5 h/mês |
| Mês 3-6 | Recrutar 3-5 médicos da sua rede e montar o pod. | R$ 10-20 mil | 30-45 min/dia |
| Mês 6-12 | Produtizar o ferramental do pod. Ativo com valor de venda. | — | 30 min/dia |

A meta original de R$ 5.000 a R$ 10.000 é atingida na Fase 1, com 22 a 30 minutos por dia, sem
capital e sem risco. Tudo depois disso é construção de patrimônio, não de renda.

---

## Fontes

Juros e tributação: [Meelion - Selic 14,25% em 05/08/2026](https://www.meelion.com/indicadores-financeiros/selic/) ·
[Renova Invest - CDI 14,15%](https://renovainvest.com.br/blog/cdi-hoje/) ·
[B3 Bora Investir - MP 1.303 e LCI/LCA](https://borainvestir.b3.com.br/noticias/lcis-e-lcas-tributacao-da-renda-fixa-jcp-como-a-mp-1303-afeta-seus-investimentos/) ·
[Adriano Freire - LCI e LCA isentas em 2026](https://www.adrianofreire.com.br/blog/reforma-tributaria-lci-lca-cdb-2026)

Múltiplos e compra de ativos: [Flippa - SaaS Valuation Multiples 2026](https://flippa.com/blog/saas-multiples/) ·
[BigIdeasDB - SaaS multiples 2026](https://bigideasdb.com/saas-valuation-multiples-2026) ·
[Flippa - Online Business M&A 2025/2026](https://flippa.com/blog/2025-online-business-ma-insights-from-flippa/) ·
[Mediafast - How to sell your micro-SaaS 2026](https://www.mediafa.st/how-to-sell-your-micro-saas)

Micro-SaaS e prazos: [Flowjam - 27 micro SaaS examples](https://www.flowjam.com/blog/27-micro-saas-examples-that-actually-print-money-in-2025) ·
[SoftwareSeni - Solo founder SaaS metrics](https://www.softwareseni.com/solo-founder-saas-metrics-from-0-to-10k-mrr-in-6-months-with-realistic-timelines/) ·
[Superframeworks - micro SaaS para solopreneurs](https://superframeworks.com/articles/best-micro-saas-ideas-solopreneurs)

Reforma tributária: [CGIBS - marco de 03/08/2026](https://www.cgibs.gov.br/novo-marco-da-reforma-tributaria-inicia-em-03-de-agosto-com-preenchimento-obrigatorio-dos-campos-relativos-ao-ibs-e-a-cbs) ·
[Receita Federal - orientações 2026](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/reforma-tributaria-do-consumo/orientacoes-2026) ·
[Contábeis - NF-e exige IBS e CBS a partir de agosto](https://www.contabeis.com.br/noticias/77889/reforma-tributaria-nf-e-passa-a-exigir-ibs-e-cbs-a-partir-de-agosto/) ·
[Comsefaz - período de adaptação](https://comsefaz.org.br/novo/reforma-tributaria-comeca-em-2026-com-periodo-de-adaptacao-destaque-informativo-dos-novos-tributos-e-dispensa-de-penalidades/)

Automação como serviço: [Hora de Codar - quanto cobrar automação n8n](https://www.horadecodar.com.br/quanto-cobrar-automacao-n8n/) ·
[Hora de Codar - custo de agente de IA n8n 2026](https://www.horadecodar.com.br/quanto-custa-agente-ia-n8n-2026/) ·
[IBE - caso de agente recorrente R$ 2.450/mês](https://blog.ibe.ia.br/blog/como-daniel-ganha-r-2-450-por-mes-de-1-cliente-com-agente/) ·
[Message Central - WhatsApp API pricing Brasil 2026](https://www.messagecentral.com/blog/whatsapp-business-api-pricing-brazil)

Nichos validados na Parte 2: [PsiSync](https://psisync.com.br/) ·
[Hotina - preços](https://hotina.app/precos/) ·
[PsicoManager - planos](https://www.psicomanager.com.br/planos) ·
[Agilize - melhores softwares para psicólogos 2026](https://agilize.com.br/artigos/melhor-software-gestao-psicologos-2026/) ·
[Receita Federal - manual Receita Saúde 2.1](https://www.gov.br/receitafederal/pt-br/centrais-de-conteudo/publicacoes/manuais/orientacao-tributaria/receita-saude-2.1.pdf) ·
[Sebrae - MEIs lideram aberturas em 2026](https://agenciasebrae.com.br/dados/meis-lideram-abertura-de-empresas-no-pais-e-ja-representam-78-dos-novos-negocios-em-2026/) ·
[Contajá - comparativo de contabilidades online](https://contaja.com.br/blog/melhor-contabilidade-online/) ·
[Pasta Contábil - automação de cobrança de documentos](https://pastacontabil.com.br/blog/automacao-cobranca-documentos-clientes-contabilidade) ·
[Jettax - automação fiscal para escritórios](https://www.jettax.com.br/) ·
[DPOnet - ANPD 2026, fiscalização e sanções](https://dponet.com.br/blog/anpd-2026-fiscalizacao-lgpd-empresas-sancoes/) ·
[Confidata - mapa de sanções da ANPD](https://confidata.com.br/blog/mapa-sancoes-anpd-todos-casos-2026)

Mercado global (Parte 4): [Mercor - AI training para profissionais de saúde](https://www.mercor.com/experts/healthcare/) ·
[Mercor - vaga Healthcare Expert, International Physicians](https://work.mercor.com/jobs/list_AAABn4b_udoVxl0h5FFGKJex/healthcare-expert-international-physicians) ·
[Mercor - países suportados](https://www.mercor.com/supported-countries/) ·
[Handshake AI - Medicine Expert (exige instituição dos EUA)](https://joinhandshake.com/ai/opportunities/medicine-expert-ai-trainer/) ·
[HireFeed - Outlier vs Mercor vs Surge, faixas de 2026](https://hirefeed.co.in/blog/outlier-vs-mercor-vs-surge-2026) ·
[AI Gig Jobs - médicos ganhando US$150+/h treinando modelos](https://www.aigigjobs.com/blog/medical-professionals-ai-earnings) ·
[SF Standard - médicos assumindo trabalho paralelo de tutoria de IA](https://sfstandard.com/2026/04/17/sf-doctors-taking-side-hustles-tutoring-ai/) ·
[Outsource Accelerator - anotação multilíngue em 2026](https://www.outsourceaccelerator.com/articles/multilingual-data-annotation/) ·
[BrassTranscripts - português lidera demanda não inglesa](https://brasstranscripts.com/blog/global-ai-transcription-trends-2026) ·
[M3 Global Research](https://www.m3globalresearch.com/) ·
[SalaryDr - melhores questionários médicos pagos 2026](https://www.salarydr.com/blog/best-paid-medical-surveys-physicians-2026) ·
[CleverX - comparativo de redes de especialistas 2026](https://cleverx.com/blog/best-expert-network-platforms-complete-comparison-guide/) ·
[GTE Localize - tarifas de tradução médica por palavra](https://gtelocalize.net/medical-translation-rates-per-word/) ·
[myJuno - custo de tradução em 2026](https://myjuno.io/translation-services-costs-2026-pricing-guide/) ·
[medRxiv - reconhecimento de fala médica em espanhol latino-americano](https://www.medrxiv.org/content/10.64898/2026.07.14.26358062v1.full) ·
[Voa Health - comparativo de escribas de IA em português](https://blog.voa.health/blog/ia-e-inovacao-4/ias-transcricao-consultas-medicas-2025-50) ·
[Glass Health - escribas multilíngues e preços](https://glass.health/resources/best-multilingual-ai-medical-scribe) ·
[Contabilidade Zen - receber em dólar, PJ ou PF](https://www.contabilidadezen.com.br/blog/receber-dolar-pj-ou-pf-2026/) ·
[Moura - tributação em exportação de serviços 2026](https://mouraservicoscontabeis.com.br/tributacao-em-exportacao-de-servicos/)

E-commerce sem estoque e infoprodutos: [Ganhe Recompensa - dropshipping 2026, realidade vs ilusão](https://ganherecompensa.com.br/blog/dropshipping-2026-vale-a-pena-realidade-vs-ilusao) ·
[Many Store - custos do dropshipping em 2026](https://manystore.com.br/dropshipping-em-2026-vale-a-pena-custos-fornecedores-e-como-comecar) ·
[Goodds - mitos e lucros de print on demand 2026](https://www.bloggoodds.com.br/print-on-demand-mitos-lucros-modelos-2026/) ·
[Nexus Growth - Meta Ads 12,15% mais caro em 2026](https://blog.nexusgrowth.com.br/blog/meta-ads/meta-ads-custo-2026) ·
[Vini Ensina - 48 estatísticas de custo de Meta Ads Brasil 2026](https://viniensina.com.br/estatisticas-meta-ads-brasil-2026/) ·
[Ganhe Recompensa - vender infoprodutos na Hotmart e Kiwify 2026](https://ganherecompensa.com.br/blog/como-vender-infoprodutos-hotmart-kiwify-2026) ·
[Bit4learn - Kiwify, taxas e análise 2026](https://bit4learn.com/pt/lms/kiwify-venda-cursos-online/)

Descartes iniciais: [Faceless.my - YouTube inauthentic content policy](https://faceless.my/youtube/youtube-inauthentic-content-policy-faceless-ai-2026/) ·
[AITuber - faceless demonetizados](https://aituber.app/blog/faceless-youtube-channels-demonetized-2026/) ·
[PubNook - KDP low-content policy 2026](https://pubnook.com/article/amazon-kdp-low-content-book-policy-2026-whats-allowed-and-what-gets-removed) ·
[Univers - KDP AI content policy 2026](https://www.univers.studio/blog/kdp-ai-content-policy-2026/) ·
[Correio do Estado - custo de lavanderia self-service 2026](https://correiodoestado.com.br/mix/quanto-custa-abrir-uma-lavanderia-self-service-no-brasil-em-2026/) ·
[HostnJoy - quanto ganha co-anfitrião](https://hostnjoy.com.br/blog/co-anfitriao-airbnb) ·
[GeoStack - mapa de agências GEO no Brasil](https://geostack.com.br/mapa-das-agencias-de-geo/)
