# Prescrição v2 — catálogo, ergonomia e manipulados

**Status:** itens 1 a 3 + frente magistral (fase 3) implementados em DEV (nada deployado). Ver "Estado" no fim do arquivo.
**Data:** 2026-08-19 · **Atualizado:** 2026-08-19

## 1. Por que não salvava (corrigido, commit `1b76c821`)

`apps/web/lib/api/prescriptions.ts` chamava `POST /prescriptions` em vez de
`POST /api/v1/prescriptions`. Toda operação de prescrição batia 404 — listar, abrir, salvar,
assinar, excluir. O log da API em produção mostrava a batida errada repetida
(`GET /prescriptions 404`). O mesmo defeito existia em
`apps/web/lib/api/medication-definitions.ts`, que é por que a busca de medicamento nunca
achava nada. Eram os dois únicos módulos de `lib/api` sem o prefixo; o resto está correto.

## 2. Catálogo de medicamentos: nunca foi importado

`medication_definitions` está com **0 linhas em produção** e **não existe importador no
repositório** — nenhum `cmd/import-medications`, nenhum seed, nenhum CSV. A tela sempre
dependeu de cadastro manual, que nunca aconteceu. Corrigir a URL faz a busca funcionar, mas
ela continua sem resultados até popular.

### Fontes oficiais

| Fonte | O que tem | Serve para |
|---|---|---|
| [Dados Abertos ANVISA — registros válidos](https://dados.anvisa.gov.br) | registro, produto, princípio ativo, empresa, classe terapêutica | nome comercial ↔ princípio ativo, código ANVISA |
| [CMED — listas de preços](https://www.gov.br/anvisa/pt-br/assuntos/medicamentos/cmed/precos) | substância, produto, apresentação, laboratório, EAN, **tarja/lista de controle**, preços | apresentação/concentração e a categoria regulatória |

A **tarja** da CMED é o que permite preencher `category` (`simple` / `c1` / `c5` /
`antibiotic` / `glp1`) automaticamente — e é dela que saem as regras que o modelo já tem
(validade da receita, máximo de substâncias, exigência de assinatura e de SNCR).

### Importador proposto (`apps/api/cmd/import-medications`)

1. Baixar o arquivo oficial e **inspecionar o layout real antes de mapear** (os nomes de
   coluna mudam entre edições; nada de assumir).
2. Normalizar: `commonName` = produto + apresentação, `activeIngredient` = substância,
   `anvisaCode` = registro, `category` = derivada da tarja.
3. Campos novos que a tela precisa e o modelo ainda não tem: **forma farmacêutica**,
   **concentração**, **laboratório**, **apresentação**, **genérico/similar/referência**.
4. Idempotente por (registro + apresentação), rodável de novo a cada atualização da lista.
5. **Dedupe para a busca**: 25 mil apresentações poluem o autocomplete. Duas camadas —
   busca por **princípio ativo** (prescrever genérico pela DCB) e por **produto comercial**.

## 3. Ergonomia da prescrição

Hoje cada medicamento pede **11 campos**, quase todos texto livre: Nome Comercial, Princípio
Ativo, Categoria, Concentração, Via, Dosagem, Frequência, Duração, Quantidade, Quantidade por
Extenso (essa sim automática) e Instruções. Escolher do catálogo preenche só três (nome,
princípio, categoria) — o resto é digitação, toda vez, inclusive para o medicamento que o
médico prescreve todo dia.

| Problema | Proposta |
|---|---|
| Catálogo preenche pouco | Selecionar traz concentração, forma, via padrão, categoria e validade |
| Dose/frequência/duração em texto livre | Posologia estruturada com atalhos ("1 cp", "8/8h", "7 dias") que compõem a frase |
| Quantidade digitada à mão | Calculada de dose × frequência × duração (8/8h por 7 dias = 21), com override |
| "Comprimidos" fixo no extenso | Unidade vem da forma farmacêutica (gotas, mL, cápsulas) |
| Tudo do zero a cada receita | **Favoritos e repetir receita anterior** — o maior ganho real de tempo |
| Regras regulatórias invisíveis | Avisar no ato: C1 limita 3 substâncias/60 dias e exige assinatura |
| Princípio ativo e categoria editáveis | Read-only quando vêm do catálogo; texto livre só no modo manual |

Ordem de maior retorno: **repetir/favoritos > catálogo preenchendo > quantidade calculada >
posologia estruturada**.

## 4. Manipulados — hoje não existe

O modelo atual (`prescription_medications`) é uma linha por medicamento industrializado. Uma
fórmula manipulada é outra coisa: **um item composto de N substâncias**, cada uma com sua
concentração, mais forma farmacêutica, quantidade a aviar e uma posologia única — e vai em
receituário próprio, endereçado à farmácia de manipulação.

### Proposta

- `prescriptions.type`: `commercial` | `compounded`. O switch troca o formulário inteiro,
  como pedido.
- Item manipulado: tabela `prescription_formulas` (forma farmacêutica, quantidade a aviar,
  posologia, uso interno/externo) + `prescription_formula_components` (substância,
  quantidade, unidade). Tabela filha explícita em vez de JSONB: dá para consultar, validar e
  reaproveitar fórmula.
- **Fórmulas favoritas** reaproveitáveis entre pacientes — é o padrão de quem manipula.
- PDF em layout de manipulação: componentes listados, "aviar N cápsulas", posologia, via.
- Regras regulatórias continuam valendo por substância (uma fórmula com substância C1 é
  receituário de controle especial).

## 5. Ordem sugerida

1. ~~Desbloquear produção (URL)~~ — feito.
2. Catálogo ANVISA/CMED + busca que preenche os campos.
3. Manipulados (switch, modelo, formulário, PDF).
4. Ergonomia: repetir receita, favoritos, quantidade calculada, posologia estruturada.

O item 2 é pré-requisito real do 4: sem catálogo, não há o que autopreencher.


---

## Estado (2026-08-19, tudo em DEV — nenhum deploy)

### Feito

**Desbloqueio de produção.** Além da URL sem `/api/v1` (item 1), o INSERT batia numa causa mais
funda: `prescriptions` guardava, do desenho antigo (1 receita = 1 medicamento), nove colunas
`NOT NULL` sem default que o model Go não preenche mais. Nenhuma receita jamais foi gravada em
produção (`prescriptions` = 0 linhas). Migration `00067_prescricao_colunas_legadas.sql` derruba as
colunas órfãs.

**Catálogo ANVISA/CMED (item 2).** `internal/cmed` + `cmd/import-cmed`, migration `00066`,
26.001 apresentações importadas em dev, busca em dois níveis e rotina de conferência das
classificações deduzidas (`/admin/medication-definitions/revisao`).

**Manipulados (item 3).** Migration `00068`: `prescriptions.type` + `prescription_formulas` +
`prescription_formula_components`. Models, DTOs, validações (10 fórmulas × 20 componentes, 3
substâncias C1 por receita, extenso obrigatório em controlado), receituário magistral no `pdfdoc`
(substância à esquerda, quantidade à direita, pontilhado, veículo, "Aviar N", uso interno/externo)
e formulário com switch `?tipo=manipulado`. Emitido e assinado ponta a ponta em dev.

**Ergonomia (item 4, parte).** Repetir receita (`?from=<id>`), frequência estruturada com
quantidade calculada, extenso pela forma farmacêutica, menos campos obrigatórios, defaults de
veículo/unidade/uso por forma farmacêutica no manipulado.

**Dívidas pagas no caminho.** Schema zod e formulário extraídos (as duas telas eram 1.500 linhas
duplicadas); `/prescriptions/[id]` e `/duplicate` criadas (davam 404); Download pelo endpoint
autenticado; Excluir rascunho funcionando; `Create`/`Update` em transação; guarda de receita
assinada (409); código morto de gofpdf removido.

### Decisões tomadas

- **Limite de C1 conta substâncias distintas**, não itens — é a regra da 344/98, e vale para os
  dois tipos de receita.
- **Receita magistral pode misturar fórmula controlada e comum.** A 344/98 quer receituário
  exclusivo, mas quem separa é o médico na impressão: a tela avisa, não bloqueia.
- **Tipo é imutável** depois de criada a receita (não há switch na edição).

**Frente magistral (fase 3).** Migration `00069`: `magistral_components` + `magistral_incompatibilities`.

- **Curadoria oportunista, não projeto de cadastro.** Substância prescrita que ainda não está no
  catálogo entra sozinha como esboço (só nome e unidade); a que já está ganha +1 no contador de
  uso, e a busca ordena por isso — em algumas semanas o autocomplete fica igual ao repertório
  real. Quando o médico dosa uma substância sem faixa cadastrada, aparece "salvar como padrão".
- **Calculadora de cápsula** (`magistral_capsule.go`, função pura com teste de mesa): volume =
  Σ(massa/densidade), escolhe o MENOR tamanho que resolve em uma cápsula, cai para duas só se nem
  a 000 couber, e recomenda sachê quando nem isso resolve. Resultado é FAIXA (±25%), não número
  seco. **Sem densidade cadastrada de qualquer componente, ela não opina** e diz de quem falta.
- **Compatibilidade e palatabilidade** (`magistral_compat.go`): regras derivadas de sinalizador
  (eutético × eutético, oxidante × oxidável, higroscópico em cápsula gelatinosa, amargor alto em
  sachê) + pares curados com mecanismo. Avisa, não bloqueia — mesmo padrão do alerta de alergia.
- **Semente com fonte** (`docs/emr/magistral-seed-inicial.sql`): mentol, cânfora, timol, fenol e
  resorcina como formadores de eutético; talco × cianocobalamina como par curado.
- **Semente do repertório funcional/nutrológico** (`docs/emr/magistral-seed-funcional.sql`): 85
  substâncias levantadas em formulários e fórmulas publicadas por farmácias magistrais brasileiras,
  organizadas por frente (sono/humor, magnésios, mitocôndria, resistência insulínica, antioxidante/
  imunidade, intestino, cognição, pele/cabelo, menopausa/hormonal). 61 com dose usual de fonte; as
  24 restantes entram só com nome, sinônimos e o que se sabe em `notes`.
  - O teto da **IN 28/2018 (Anexo IV)** foi para `notes`, NUNCA para `max_dose`: é teto de
    suplemento industrializado, e vitamina D (2.000 UI), biotina (45 mcg) e B12 (9,94 mcg) são
    superados de propósito na prática magistral. Em `max_dose` isso alertaria em toda receita
    legítima.
  - **Densidade aparente**: só os dois valores publicados que a pesquisa achou (sulfato de
    condroitina 0,60 g/mL, maltodextrina 0,70 g/mL). Não existe tabela pública por ativo e o valor
    muda por lote (Anfarmag; Quallitá). Enquanto não vier da farmácia, a calculadora se cala.
  - **Higroscopicidade** marcada só onde a fonte nomeia a substância (cloreto de magnésio,
    carnitina, acetil-carnitina, fosfolipídios, minerais quelados). As fontes citam "aminoácidos" e
    "extratos secos" como classes sensíveis — marcar a classe faria a tela avisar em quase toda
    cápsula, então ficou registrado sem ligar o sinalizador.
  - Pares de **competição de absorção** (zinco × cobre, zinco × ferro, cálcio × ferro,
    vitamina C × B12) entram como `info` e o texto diz que são competição de absorção, não
    incompatibilidade de manipulação.
- Tela de curadoria em `/admin/magistral-components` (densidade, faixa de dose, sinalizadores,
  pares) e painel de conferência dentro do formulário de manipulado.

**Indicações e evidência do RAG (fase 4, primeira metade — feita).** Migration `00070`:
`indications`, `dose_reference`, `evidence_status` e a tabela `magistral_component_articles`.

- **O passo 0 que o plano exigia foi feito e o resultado é bom**: o RAG tem 240 aulas da pós de
  Medicina Funcional Integrativa entre os 1.183 artigos, e elas cobrem de fato o repertório —
  72 das 95 substâncias acharam material (740 dos 760 trechos vieram de AULA, não de artigo).
  A fase não é projeto de conteúdo; o conteúdo já estava lá.
- `cmd/magistral-enrich` busca os trechos por similaridade (embedding da substância + sinônimos),
  guarda cada um como evidência anexada e pede ao Claude que **extraia** indicação e posologia do
  que está escrito. Salvaguardas: dose numérica só entra em campo VAZIO e só quando a unidade
  bate com a do catálogo; nada sobrescreve curadoria manual; tudo entra como `suggested` e o
  médico confirma na tela, com o trecho ao lado.
- As 23 substâncias que o material não cobre receberam indicação de pesquisa externa
  (`docs/emr/magistral-indicacoes-externas.sql`) — inclusive as farmacotécnicas, cuja "indicação"
  é o papel na fórmula.
- **Painel lateral do catálogo** (`magistral-substance-sheet.tsx`) no lugar do dropdown de texto:
  o campo de substância tem dois caminhos — a lista curta de uma linha para quem já sabe o que
  quer, e a lupa que abre o painel para quem está decidindo. O painel busca por substância ou
  pelo que se quer tratar, compara candidatas em cartões, e no detalhe mostra dose usual com a
  faixa desenhada, sinalizadores, indicação e posologia **em tópicos** (texto completo a um
  clique), o aviso de densidade faltando e os trechos das aulas com a proximidade, cada um
  expansível. A dose é editável no rodapé antes de entrar na fórmula.
  - A primeira versão despejava a indicação inteira no dropdown: uma substância ocupava a lista
    e escondia as outras. Texto longo agora só no painel.
  - Os tópicos ficam em `indication_bullets`/`dose_bullets` (migration `00071`), gerados por
    `cmd/magistral-bullets` como COMPRESSÃO FIEL do texto já gravado — sem busca nova, sem fonte
    nova. Uma linha por tópico, texto puro e não jsonb, porque o campo é editado à mão.

**Fórmulas-base e dose dinâmica (fase 4, segunda metade — feita).** Migration `00072`:
`magistral_formula_templates`, `..._components`, `..._rules`, `..._articles`, FK de
`prescription_formulas.template_id` e `prescription_formula_components.suggested_quantity`.

- **Fórmula-base**: a fórmula pronta reusada entre pacientes, com indicação própria (texto e
  tópicos) e busca por nome OU pelo que trata. Salva a partir da fórmula montada na receita
  ("Salvar como base") ou direto em `/admin/formulas-base`. A receita emitida registra de qual
  base saiu, e a base sobe na lista conforme vira receita.
- **Dose dinâmica**: regra por componente — `fixed`, `per_kg` ou `lab_threshold`. O exame é uma
  referência ao `code` do catálogo, escolhida na tela; nenhum código clínico entra hardcoded.
- **As três fronteiras que sustentam isso** (`magistral_dose_rules.go`, função pura com teste):
  1. a sugestão NUNCA escreve na receita — preenche um campo que o médico ainda vai olhar, e o
     payload de criação de prescrição não tem noção de regra, então um bug no motor é incapaz de
     produzir receita assinada;
  2. `min_dose`/`max_dose` são NOT NULL: peso errado no prontuário não vira dose absurda, a trava
     corta e a resposta DIZ que cortou;
  3. dado ausente ou mais velho que o prazo da regra não sugere nada, e devolve o motivo em
     português.
- **A base é sempre escrita**: "peso 82,5 kg (peso da consulta) × 4 mg/kg" ou "25-hidroxivitamina
  D = 37 ng/mL de 09/02 · regra: não atinge < 30 ⇒ 2.000 UI · reavaliar em 90 dias". O ramo falso
  diz "não atinge" — a primeira versão escrevia "regra: < 30" com valor 37, o que sugeria o
  contrário do que houve.
- **"Resposta do paciente" não foi inventada**: o EMR não tem escala estruturada, então a tela
  mostra a última dose que o próprio médico ASSINOU daquela substância naquela base.
- `suggested_quantity` fica gravada ao lado de `quantity`: é o que revela, depois, quando o médico
  discordou da sugestão.

**Biblioteca de fórmulas-base (conteúdo).** O mecanismo sem conteúdo não serve para nada, então a
biblioteca foi montada em dev: **17 fórmulas**.

- `cmd/magistral-formulas` minera as AULAS por frente clínica e extrai as fórmulas COMPLETAS
  descritas ali (componentes com dose). Regras: componente sem dose não entra, fórmula com menos
  de dois componentes com dose não entra, e trecho que não descreve associação não vira fórmula —
  o comando devolve lista vazia sem reclamar. Rendeu 10 fórmulas, com os trechos anexados.
- Onde o material discute a substância isolada sem fechar associação (intestino, climatério,
  pele, ansiedade), as fórmulas vêm de formulários magistrais publicados
  (`docs/emr/magistral-formulas-base-seed.sql`), com a farmácia-fonte anotada em `notes`.
- **Nenhuma regra de dose foi criada pela máquina.** Regra tem trava de piso e teto e é decisão
  clínica: fica uma só, marcada como exemplo (vitamina D conforme 25-OH), para mostrar o
  comportamento. As demais o médico cadastra.
- Foram descartadas na conferência: duas que eram produto industrializado citado em aula
  (Omega-COR, Omega Joint), uma de aromaterapia e uma com nome de substância provavelmente mal
  transcrito.

**Densidades aproximadas (calculadora destravada).** Migration `00073` acrescenta
`density_source`, e `docs/emr/magistral-densidades-aproximadas.sql` preenche as 95 substâncias.

- A decisão mudou por pedido explícito: em vez de esperar o número exato de cada insumo, entram
  **aproximações por classe de pó** (mineral inorgânico 0,90 · vitamina hidrossolúvel 0,75 ·
  mineral quelado 0,65 · aminoácido 0,60 · fibra 0,55 · probiótico 0,50 · extrato seco e
  lipossolúvel 0,45 · colágeno 0,35), calibradas por âncoras publicadas: cápsula 0 comporta 400 a
  450 mg de pó comum, a 00 comporta 500 a 600 mg, creatina 0,55, maltodextrina 0,70, condroitina
  0,60.
- `density_source` distingue `medida` de `classe`, e a calculadora ACRESCENTA a frase "densidade
  aproximada por classe do pó" quando qualquer componente usa aproximação. Estimativa declarada é
  útil; estimativa disfarçada de medida é que seria o problema. Editar o valor à mão troca o
  rótulo para "informada pelo médico".
- Reverter é uma linha: `UPDATE magistral_components SET bulk_density=NULL, density_source=NULL
  WHERE density_source='classe'`.

**Material das farmácias parceiras (25 documentos).** Fichas técnicas de insumo, formulário com 80
fórmulas e lâminas de ativos de marca. O que saiu de lá:

- **Fator de correção** (migration `00074`, `elemental_percent` + `as_elemental` no componente).
  Toda ficha técnica traz esse campo, e ele resolve uma ambiguidade real: "Magnésio quelato
  300 mg" pode ser 300 mg de elemento (1 g de bisglicinato 30%) ou 300 mg do quelato (90 mg de
  elemento). A tela pergunta, a conversão aparece, e **a calculadora de cápsula passa a usar a
  massa do insumo** — no exemplo, o volume salta de 1,13 mL para 2,21 mL e a recomendação muda de
  duas cápsulas 00 para sachê. Sem isso ela subestimava toda fórmula com quelato ou diluído.
- **Formas preferidas do prescritor**: metilcobalamina, metilfolato, piridoxal-5-fosfato,
  palmitato de ascorbila, selenometionina e CavaQ10. A substância genérica aponta para a forma que
  ele usa, e a tela sugere a troca com a nota do que a literatura sustenta — inclusive quando não
  sustenta.
- **10 ativos novos** no catálogo, com dose e evidência anotadas (CavaQ10, Morosil, açafrão
  padronizado, Cactinea, Glycoxil, Bio Arct, Koubo, green coffee, palmitato de ascorbila,
  selenometionina) e mais 5 vindos das fórmulas (rhodiola, L-tirosina, D-ribose, piperina,
  policosanol).
- **7 fórmulas do formulário**, curadas: formas trocadas pelas dele, erros de unidade corrigidos
  (o formulário traz "selênio 30 mg" onde é mcg, e "vitamina A 50 mg") e duplicidade de B6
  removida (uma fórmula somava P5P 7 mg com B6 100 mg, acima do teto da IN 28).

### Próximo

Conferência clínica: as 110 substâncias e as 24 fórmulas estão marcadas como sugestão.

Pendências que dependem de você: veículos padrão por forma farmacêutica (os atuais são palpite
razoável); se o tamanho da cápsula deve sair impresso no receituário (recomendo só no EMR — a
farmácia é a autoridade técnica); e as densidades aparentes da sua farmácia de confiança, que são
o que liga a calculadora.

---

## Fase 7 — Qualidade da dose dinâmica (migration 00075)

O motor de dose dinâmica existia desde a Fase 4 e estava sendo usado por **uma regra em 24
fórmulas** — construído e quase inerte. Três buracos explicavam boa parte disso, e todos foram
medidos no dado real antes de virar código.

### 1. A unidade do exame não era conferida (o mais grave)

O motor comparava `lab_results.result_numeric` com o limiar da regra **sem nunca olhar em que
unidade o resultado foi gravado**. Medido no banco: **390 dos 1.243 resultados numéricos (31%)
estão numa unidade diferente da definição do exame** — cortisol em nmol/L sobre definição em
µg/dL (fator 27,6), contagens absolutas gravadas em %, e uma 25-hidroxivitamina D gravada em
pg/mL. Pior: o motor lia a unidade da *definição* e o valor do *resultado*, então a frase que o
médico lê podia dizer "37 ng/mL" sobre um número que o laboratório reportou em outra escala.

Agora a regra guarda a unidade em que foi escrita (`lab_unit`, preenchida pela tela a partir do
exame escolhido), o service lê a unidade do resultado com a da definição só como fallback, e
resultado em unidade diferente **não vira sugestão** — devolve o motivo nomeando as duas
unidades. A guarda só é útil se não for barulhenta: variantes cosméticas da mesma unidade
(`mcg/L` ≡ `µg/L` ≡ `ug/L`, `µUI/mL` ≡ `µU/mL` ≡ `mUI/L`, `IU` ≡ `UI`) normalizam para igual e
passam sem alarde. Só entram no mapa de equivalência pares com valor numérico idêntico.

### 2. Limiar binário onde a conduta é por faixa

Vitamina D, B12 e zinco não têm um corte, têm três. Com uma regra por componente e só
`dose_if_true`/`dose_if_false`, essa conduta era impossível de escrever. Entra o tipo `lab_band`
com tabela filha de faixas, na **mesma convenção meio-aberta `(lower, upper]` das faixas do
escore** — duas convenções de faixa no mesmo sistema seria pedir para alguém ler a errada.

Faixas sobrepostas são recusadas no cadastro (com sobreposição, a dose passaria a depender da
ordem de inserção); buraco entre faixas é permitido de propósito, e o motor responde dizendo que
não há conduta cadastrada em vez de escolher a faixa mais próxima.

A tela do médico mostra **a escada inteira com a faixa dele marcada**, não só o degrau que pegou:
vendo a régua ele julga a conduta, não apenas o número.

### 3. Dose sem arredondamento prático

Regra por peso devolvia 1.234,5678 mg. `round_to` é o passo da substância (500 UI, 50 mg, 5 mg).
Arredonda **antes** de travar, senão o arredondamento poderia devolver dose acima do teto — e a
frase diz que arredondou e diz que travou.

### As regras que entraram

De 1 para **10 regras em 7 fórmulas**, cada uma com procedência declarada na própria nota, e
dizendo quando a casa e a diretriz divergem:

| Substância | Exame | Conduta |
|---|---|---|
| Vitamina D3 (3 fórmulas) | 25-OH-D | 5 faixas, alvo 40 a 60 ng/mL |
| Metilcobalamina (2) | B12 | 3 faixas, alvo > 550 pg/mL |
| Zinco quelato (3) | zinco sérico | 3 faixas, alvo ≥ 100 µg/dL |
| Selenometionina (1) | anti-TPO | > 35 UI/mL ⇒ 200 mcg |
| Magnésio glicina (1) | peso | 5 mg/kg de elemento, teto 350 mg |

Os alvos vêm das aulas da pós (RAG); as notas trazem a literatura ao lado, **inclusive quando ela
diverge**: a Endocrine Society de 2024 não fixa alvo de vitamina D e desaconselha rastreio; nas
metanálises de Hashimoto a selenometionina é a única forma com efeito nos anticorpos, mas a queda
de TSH é pequena e só um terço dos braços mostrou queda de anti-TPO.

### Correções de dado achadas ao escrever as regras

- **Vitamina D3 de 50 UI** na fórmula de hipotireoidismo — 50 UI não tem efeito biológico nenhum,
  é erro de transcrição do formulário. Base corrigida para 2.000 UI.
- **Magnésio quelato lido de dois jeitos**: dose do elemento numa fórmula, dose do insumo em
  outra. Mesma substância, três vezes de diferença na cápsula. Padronizado em elemento.
- **Sachê noturno sem posologia nenhuma** — a receita sairia sem dizer como tomar.
- O template "Vitamina D conforme exame (exemplo de regra)" era andaime; virou fórmula de verdade,
  com indicação própria.

## Fase 8 — Conferência das 24 fórmulas pelo próprio painel (migration 00076)

As 24 fórmulas-base foram passadas pelo endpoint de conferência, uma a uma, e o que ele acusou
virou auditoria. Três defeitos do sistema apareceram antes de qualquer julgamento clínico.

### 1. Nome com parêntese não casava com o catálogo

Quatro fórmulas ficavam sem cálculo de cápsula por "falta de densidade" de substâncias que
**estão** no catálogo, com densidade: as fórmulas escrevem "Vitamina B6 (piridoxal-5-fosfato)",
"PQQ (pirroloquinolina quinona)", "Coenzima Q10 (ubiquinona) ou ubiquinol", e a busca era só por
nome exato. `findByName` virou uma escada de tentativas — nome, sinônimo, sem parêntese, conteúdo
do parêntese, sem qualificador de insumo — **sem nenhum degrau aproximado**: casar por semelhança
traria faixa de dose e densidade de outra substância para dentro de uma receita. Qualificador que
muda o insumo ("anidra", "quelato") nunca é descartado.

### 2. A faixa de dose não dizia se era por dia ou por tomada

Medido: em **7 das 71 substâncias** com faixa numérica e texto de posologia, a faixa numérica
**não encosta no que o próprio texto da linha diz**. Gimnema cadastrada em 12,5 a 25 mg com o
texto dizendo 300 mg. Alfa-lipoico em 25 a 200 mg com o texto dizendo 300 a 600.

A origem: a faixa numérica foi semeada das fórmulas das parceiras (dose de UMA cápsula) e o texto
veio da literatura (dose do DIA). O painel comparava uma contra a outra — **fórmula com dose baixa
passava sem alerta e fórmula correta era acusada**. Pior: com a faixa vinda da própria fórmula, o
catálogo conferia a fórmula contra ela mesma.

Entra `dose_basis` (`por_dia` default, `por_dose`), a tela deduz as tomadas/dia da posologia e o
painel compara na base certa, dizendo qual base usou e quantas tomadas contou.

### 3. Faixa oral aplicada a sublingual

Sublingual pula a primeira passagem e usa dose menor de propósito — o próprio registro do 5-HTP
diz "sublingual a partir de 20 a 35 mg". Acusar dose baixa ali é acusar a via, não a dose. O lado
de cima continua valendo: dose alta por qualquer via merece conferência.

### Faixas corrigidas (`docs/emr/magistral-faixas-corrigidas.sql`)

Gimnema, NAC, picolinato de cromo, K2 MK-7, acetil-L-carnitina, PQQ, ácido alfa-lipoico, iodo
(o piso estava acima da própria RDA, e em tireoidite autoimune o excesso de iodo é gatilho),
5-HTP, valeriana, P5P, isoflavona e metilcobalamina.

Duas anotações de método que valem mais que os números:

- **Ginseng brasileiro voltou a não ter faixa.** O próprio registro dizia "os trechos não informam
  posologia" e mesmo assim havia 30 a 60 mg cadastrados, que era a dose da fórmula. Faixa
  inventada é pior que faixa ausente: vira alerta com cara de fundamento.
- **Metilcobalamina tinha o catálogo brigando com a regra de dose.** A regra por faixa de B12
  prescreve 100 mcg de manutenção para quem está no alvo e o catálogo dizia que 100 mcg é dose
  baixa. Dois subsistemas discordando da mesma substância.

Resultado: **26 alertas de dose → 17**, com os falsos positivos eliminados e o que sobrou sendo
achado clínico de verdade.

### O que sobrou é decisão clínica

As fórmulas de muitos ativos dosam cada um bem abaixo do estudado. Ordenado pela distância:
glicina 6,7× abaixo (75 mg contra 3 g dos ensaios de sono), NAC e alfa-lipoico e resveratrol 6×,
treonato 5×, curcumina 5×, gimnema e valeriana 4×. Não é erro de digitação: é o desenho das
fórmulas "guarda-chuva" das parceiras. A decisão de subir a dose ou tirar o ativo é sua.

### Segunda passada da conferência

**Um erro meu, achado pelo próprio painel.** A regra por peso do "Sachê matinal mitocondrial" foi
escrita sobre "Magnésio glicina" marcado como dose do elemento — e "Magnésio glicina" não casava
com "Magnésio quelato" no catálogo. Sem casar, não há fator de correção: a farmácia receberia
150 mg do bisglicinato (45 mg de magnésio) no lugar dos 500 mg pretendidos. Resolvido por
sinônimo, e conferido: a dose do elemento agora converte de volta para os 500 mg do insumo.

Junto entraram os sinônimos que faltavam (magnésio treonato, gimnema, cromo), os percentuais
elementares das formas de magnésio que estavam sem (dimalato 15,5%, taurato 8,9%, aspartato 8,4%,
ascorbato 6,5%), o do picolinato de cromo (12,4%) e as quatro substâncias que as fórmulas usavam
e não existiam no catálogo: beta-hidroxibutirato, ácido hidroxicítrico, *Gynostemma pentaphyllum*
e goma cássia — esta última **sem faixa de dose de propósito**, porque é excipiente de textura.

Estado depois das duas passadas: **nenhuma substância das 24 fórmulas deixa de casar com o
catálogo** (eram 19 formas de escrever que não casavam) e **nenhuma fórmula fica sem cálculo de
cápsula** (eram 4).

### Por que as fórmulas guarda-chuva são subdosadas

A pergunta óbvia depois de 19 alertas de dose baixa é "por que não corrigir e pronto". Passei cada
fórmula pelo cálculo de cápsula com todos os ativos na dose usual da literatura:

| Fórmula | Hoje | Na dose estudada |
|---|---|---|
| Antioxidante amplo | 0,59 mL · 1 cápsula 00 | 2,55 mL · **sachê** |
| Antioxidante e imunidade | 1,41 mL · 2 cápsulas 00 | 2,91 mL · **sachê** |
| Sono completo | 0,93 mL · 1 cápsula 000 | 3,54 mL · **sachê** |
| Pré-refeição glicêmico | 1,90 mL · 2 cápsulas 000 | 3,57 mL · **sachê** |
| Insônia fitoterápica | 0,50 mL · 1 cápsula 0 | 1,28 mL · 2 cápsulas 00 |
| Concentração e memória | 0,99 mL · 1 cápsula 000 | 1,32 mL · 2 cápsulas 00 |

A dose baixa não é descuido: **doze ativos na dose estudada não cabem numa cápsula.** A escolha
real não é "subir a dose", é uma decisão de desenho — menos ativos por fórmula, migrar para sachê,
ou dividir em duas. Que é uma conversa clínica, não um conserto de dado.

## Fase 9 — Carimbo de conferência, tetos da IN 28 e base (migrations 00077-00078)

### O carimbo que mentia

Os seeds gravavam `last_review` junto com a fórmula, então as 24 apareciam como revisadas sem que
ninguém tivesse olhado. Marca de conferência que nasce preenchida não é marca de conferência.
Zerado nas 24, removido dos seeds, e a lista de fórmulas passa a mostrar **"a conferir"** quando o
campo está nulo — o carimbo só existe quando um humano salva pela tela.

### Tetos do Anexo IV da IN 28

A tabela inteira entrou como **tabela de referência** (`in28_limits`, 161 nutrientes, 114 com teto
para adulto), extraída do texto consolidado da norma — não digitada à mão. Cada substância do
catálogo aponta para o nutriente correspondente (54 mapeadas) com um fator de conversão de
unidade; onde o fator seria chute, a substância fica sem mapa, porque teto conferido com fator
inventado é pior que teto ausente.

**A conferência soma por NUTRIENTE, não por substância.** É exatamente assim que o formulário das
parceiras acumulou 107 mg/dia de vitamina B6 somando piridoxal-5-fosfato com "vitamina B6" na
mesma cápsula, cada linha parecendo comportada sozinha. Aquele achado, que veio de leitura manual
documento a documento, agora é automático e permanente.

Mineral prescrito como insumo vira elemento antes de encostar na norma: 1 g de bisglicinato são
300 mg de magnésio, abaixo do teto de 350 — comparar o peso do insumo acusaria errado.

**O alerta é informativo de propósito.** O teto é de suplemento alimentar, não de prescrição:
B12 de 1.000 mcg fica 100× acima e é conduta corriqueira. Passar do teto é justamente o que separa
um suplemento de uma receita, e é isso que a frase diz. Tratar como erro faria o médico parar de
ler o painel.

Resultado nas 24 fórmulas: **9 passam de algum teto** — riboflavina, niacina, teanina, B12,
biotina e D-ribose. Nenhuma perigosa; todas informação real sobre o que a fórmula é.

### Incompatibilidade com a base

A tabela de pares cobria ativo × ativo, que é a **minoria** do problema. No levantamento de 400
prescrições magistrais da farmácia-escola da UFRJ (Vigil. sanit. debate 2019;7(1):5-13), 63% dos
erros farmacotécnicos são ativo × formulação e 23% ativo × base, contra 13% ativo × ativo. A maior
fatia não estava sendo representada.

Entra `magistral_base_incompatibilities`, com 8 regras documentadas com mecanismo e fonte: ácido
acima de 10% e ureia acima de 30% em creme Lanette (emulsão aniônica), lactato de amônio e PCA-Na
em Lanette (ionizam em qualquer concentração), LCD em vaselina sólida (polar em base apolar),
ácidos em diadermina (base saponificada), hidroquinona em creme não iônico, e cloreto de potássio
acima de 6% em xarope (limite do Formulário Nacional).

Duas decisões de projeto: a regra casa por **texto** porque o veículo é campo livre; e o percentual
mínimo só dispara quando a quantidade está escrita **em porcentagem** — comparar 15 mg com "15%"
seria inventar.

Sobre o que ficou de fora: **não existe tabela de par-a-par para cápsula oral na literatura
acessível**. O que ela documenta é ativo × base, que é o que entrou. Uma tabela grande de pares
sai do seu repertório de consultório, não de artigo — e vale mais um conjunto pequeno que nunca dá
falso positivo do que um grande e barulhento.

## Fase 10 — As 80 fórmulas do formulário, e a normalização de nomes de paciente

### O formulário inteiro entrou

O PDF do formulário das parceiras foi lido por parser, não a olho: 86 fórmulas e 502 linhas de
componente, conferidas contra o próprio sumário do documento (78 entradas). Duas armadilhas
custaram uma releitura inteira: a linha "POSOLOGIA: TOMAR..." é toda em caixa alta e virava
**nome de fórmula nova**, partindo em duas a fórmula que estava aberta (13 fórmulas ficaram com um
componente só); e a dose às vezes vem separada por um espaço só ("STREPTOCOCCUS THERMOPHILUS
1 BLH"), o que derrubava os probióticos. Saiu de fora uma "fórmula" que era a tabela nutricional
de um chocolate.

**266 grafias distintas de substância** para 188 substâncias reais. O formulário escreve em caixa
alta e abreviado: "VIT D 3", "ÁC FOLICO", "ALFA LIPOICO", "PIRIDOXAL 5 FOSFATO", "COE Q10". Foram
111 apelidos mapeados à mão para o nome canônico — cada um vira sinônimo no catálogo, então a
próxima vez que alguém digitar assim, casa. E 116 substâncias novas entraram no catálogo com nome
próprio, unidade e densidade por classe, **sem nenhuma faixa de dose**: faixa vinda da própria
fórmula é o que fazia o catálogo conferir a fórmula contra ela mesma.

Catálogo: 121 → **240 substâncias**. Fórmulas-base: 24 → **106**, todas marcadas "a conferir" — 83 em
cápsula, 17 em sachê, 6 em chocolate, 4 em solução e 2 sublinguais.

### O que a conferência automática pegou

As três correções que eu tinha achado lendo documento a documento apareceram sozinhas, e mais duas:

| Fórmula | Fonte diz | Evidência dentro do próprio documento |
|---|---|---|
| Energizante | Selenometionina 50 **mg** | as outras oito fórmulas usam 30 a 100 mcg |
| Diabetes mellitus, Hipoglicemia | Cromo 25 e 50 **mg** | sete fórmulas usam cromo em mcg, com os mesmos números |
| Hipoglicemia | Vanádio 25 **mg** | o mesmo valor aparece em mcg em outra fórmula |
| Cãimbras | Vitamina A 50 **mg** (~166.000 UI) | as demais usam 1.000 UI |
| Energizante | P5P 7 mg + vitamina B6 100 mg | somam 107 mg/dia, acima do teto de 98,6 |

Nenhuma foi adivinhada: cada correção se apoia na mesma substância aparecendo na unidade certa em
outra fórmula do mesmo documento. Onde não havia essa evidência, ficou como está e o painel aponta.

**Formas preferidas aplicadas**: 5 cianocobalaminas viraram metilcobalamina e 24 vitaminas C
viraram palmitato de ascorbila — estas marcadas como dose do ativo, porque o palmitato tem 43% de
ácido ascórbico e a troca de forma, feita sem o fator, viraria corte de dose para 43%.

Estado da biblioteca depois da importação: **65 das 106 fórmulas têm pelo menos um alerta** — 96 de
dose, 32 de elemento-ou-insumo, 31 de higroscopia, 27 de teto da IN 28 (B12 e riboflavina lideram),
16 de par e 1 de palatabilidade. Nove **não cabem em cápsula** na dose escrita.

### Nome de paciente no padrão

`NormalizePersonName` põe qualquer nome no padrão "João da Silva", venha em caixa alta, minúsculas
ou misturado, no `BeforeSave` do model — vale para recepção, portal, importação e lead de WhatsApp
sem depender de tela.

A regra pedida era "palavra de uma ou duas letras fica minúscula". Ela quebra nomes reais: **Sá**,
**Pó** e **Ré** são sobrenomes portugueses de duas letras, e **Li**, **Xu** e **Ye** são chineses
comuns aqui. O que ficou minúsculo é uma **lista de partículas** (de, da, do, das, dos, e, di, du,
del, van, von, la, le…), não um tamanho. Assim "maria de sá" vira "Maria de Sá".

Também trata: acento e cedilha na caixa certa, partícula que abre o nome (fica maiúscula), hífen e
apóstrofo como fronteira interna ("ana-maria d'ávila" → "Ana-Maria D'Ávila"), iniciais com ponto
("j.p." → "J.P."), sufixo dinástico até XII, e espaço sobrando.

O ponto **não** é fronteira de palavra fora de iniciais: sem essa ressalva, um e-mail digitado no
campo do nome — que existe no banco — virava "Getfilho@yahoo.Com.Br".

`cmd/normalize-patient-names` arruma o que já estava lá, com `--dry-run` por padrão.

### Nome do PDF de pedido de exames

`compactName` unia as partes em CamelCase ("LuizGustavoJoséCarvalho"). Agora sai
**"Luiz-Gustavo-José-Carvalho"** — o arquivo vai parar no WhatsApp e na pasta de downloads de quem
recebe, e CamelCase só se lê com esforço.

## Fase 11 — Os quatro documentos das parceiras (migrations 00079-00080)

### GLP-1: o melhor documento do lote, e um problema que ele não vê

Sete fórmulas com objetivo clínico, posologia e — o que faltava em tudo o mais — **tabela de
substâncias com faixa de dose E referência**. Entraram 16 substâncias com a citação guardada no
próprio campo de posologia, e sete que já existiam ganharam a faixa que não tinham.

**A fórmula capilar leva biotina 10 mg.** A partir de 5 mg/dia a biotina interfere em imunoensaio
biotinilado e devolve TSH, T4 livre, troponina, hCG e hormônios falsamente altos ou baixos
conforme o formato do ensaio (alertas do FDA de 2017 e 2019; documento de orientação da AACC). Num
sistema que **prescreve e lê exame**, isso fecha um ciclo ruim: a receita sai daqui e o exame
corrompido volta para cá alimentar as regras de dose dinâmica.

Virou mecanismo: `assay_interference` no catálogo e um alerta que dispara na dose diária somada, com
o que fazer antes da coleta. Migration 00079.

Outra ressalva minha: a mesma fórmula de digestão mistura **carvão ativado com enzimas** — o carvão
adsorve de forma inespecífica e leva junto o que deveria ser absorvido. Entraram 7 pares curados e
a instrução de afastar duas horas.

### Pentravan: o catálogo passou a saber o que é controlado

Este documento não traz dose oral, traz **via**. Entraram 16 substâncias transdérmicas com a tabela
de permeação do material (ativo, concentração, tecido, percentual em 24 ou 48 h) e 16 fórmulas em
duas vias novas: transdérmica e vaginal.

E expôs um buraco: a categoria de receita nascia sempre `simple`. Quem digitasse "testosterona
50 mg" numa fórmula magistral emitia **receita simples**, quando a Portaria 344/98 pede Receituário
de Controle Especial em duas vias. Agora o catálogo carrega `default_category` e a tela preenche a
partir dele — cinco fórmulas do material saem corretamente como C5. Migration 00080.

Junto entrou `regulatory_note`: a **Resolução CFM 2.333/2023** veda a prescrição de esteroides
androgênicos e anabolizantes com finalidade estética, ganho de massa muscular ou desempenho
esportivo. A fórmula do material se intitula "oxandrolona em sarcopenia **e ganho de peso**" — a
nota fica na substância e na fórmula, e é ele quem decide o que escrever no prontuário.

### Peptídeos: as substâncias sim, as fórmulas não

O PDF é diagramado em duas colunas e a extração mistura componentes de fórmulas vizinhas. Montar
receita a partir de uma atribuição que eu não consigo garantir seria pior do que não montar:
entraram as **15 substâncias** com a concentração que cada uma usa (Argireline, Syn-Ake, GHK-Cu,
Matrixyl, Haloxyl, Procapil e companhia), e as fórmulas ficaram de fora, declaradamente.

EGF e FGF entraram como `pending`: fator de crescimento tópico tem discussão regulatória própria.

### Arquitetura hormonal: três fórmulas e uma correção de expectativa

A fatia por coluna resolveu o mesmo problema de diagramação e devolveu três fórmulas limpas: eixo
androgênico, eixo adrenal e eixo DHEA.

O material apresenta Tribulus, Testofen e Eurycoma como otimizadores de testosterona livre. Para o
Tribulus a literatura não sustenta: em revisão sistemática de 2025, **oito de dez ensaios não
mostraram mudança no perfil androgênico**, e os dois positivos tiveram magnitude pequena (60 a
70 ng/dL) em homens com hipogonadismo. A metanálise de 2023 encontrou aumento não significativo.
Isso ficou escrito na substância e na fórmula — fitoterápico ali é suporte, não reposição.

Também entrou um achado de dose: o material usa **Testofen 50 mg** e **Robuvit 120 mg**, enquanto
os ensaios dos próprios ingredientes usam 300 a 600 mg e 200 a 300 mg.

### Estado da frente magistral

| | |
|---|---|
| Fórmulas-base | **132**, todas "a conferir" |
| Substâncias no catálogo | **290**, 149 com referência de posologia |
| Fórmulas com pelo menos um alerta | 75 |
| Alertas | 107 de dose · 32 de elemento-ou-insumo · 34 de higroscopia · 27 de teto da IN 28 · 19 de par · 2 de interferência em exame · 2 de palatabilidade |

## Fase 12 — As regras de dose saem de 10 para 54

O motor de dose dinâmica estava bom e quase parado: 10 regras em 7 fórmulas de 132. Espalhar as
regras exigiu antes corrigir uma premissa errada.

### A regra é diária; o campo é por cápsula

Toda regra é escrita em dose do DIA, e o campo que a tela preenche é a dose de UMA cápsula. Uma
regra de 5.000 UI/dia numa fórmula tomada duas vezes ao dia entregaria 10.000. A sugestão agora
divide pelas tomadas que lê da posologia da própria fórmula-base, **e mostra a conta**:

> Hemoglobina glicada = 5,5% de 09/02 · faixa ≤ 5,7 (normal) ⇒ 500 mg · **500 mg ao dia ÷ 2
> tomadas** · Os ensaios usam 1 a 1,5 g/dia fracionados.

A ordem importa e ficou explícita no código: a trava é diária e corta **antes** de dividir; o
arredondamento vem por último, porque o que a farmácia pesa é a dose da cápsula, não a do dia.

A leitura da posologia virou uma função só, no backend (`DosesPorDia`). Antes existia em TypeScript
para o painel de conferência e teria de existir em Go para a sugestão — duas implementações da
mesma regra divergem com o tempo. A tela manda a posologia; o servidor faz a conta.

### As regras que entraram

| Substância | Exame | Conduta | Fórmulas |
|---|---|---|---|
| Zinco quelato | zinco sérico | 3 faixas, alvo ≥ 100 µg/dL, teto 29 mg (IN 28) | 20 |
| Vitamina D3 | 25-OH-D | 5 faixas, alvo 40 a 60 ng/mL | 9 |
| Metilcobalamina | B12 | 3 faixas, alvo > 550 pg/mL | 8 |
| Metilfolato | homocisteína | 3 faixas, alvo < 10 µmol/L | 6 |
| Testosterona micronizada | testosterona total | 2 faixas **e um buraco acima de 350** | 4 |
| Berberina | HbA1c | 3 faixas, 500 a 1.500 mg/dia | 2 |
| Ferro | ferritina | 3 faixas **e um buraco acima de 70** | 2 |
| Selenometionina | anti-TPO | > 35 UI/mL ⇒ 200 mcg | 2 |
| Magnésio glicina | peso | 5 mg/kg de elemento, teto 350 mg | 1 |

**Os dois buracos são a parte mais importante.** Faixa que não cobre todo o eixo faz o motor
responder "não há conduta cadastrada" em vez de sugerir dose — e é assim que duas regras clínicas
ficam escritas na estrutura, não num comentário:

- acima de **350 ng/dL de testosterona total** o sistema não sugere reposição, porque a Resolução
  CFM 2.333/2023 só admite reposição com deficiência comprovada e nexo causal;
- acima de **70 ng/mL de ferritina** não sugere ferro, porque repor ferro sem falta comprovada é
  risco e não conveniência.

Nas duas, o motor devolve o motivo em português na tela.

54 regras em 38 das 132 fórmulas. As 94 restantes são majoritariamente fitoterápicos e blends, para
os quais não existe exame que guie dose — e regra sem exame que a sustente seria enfeite.
