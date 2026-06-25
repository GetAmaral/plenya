# Plano — corrigir redundância e profissional dos 5 templates iniciais do Continuum

**Status:** proposta (aguarda decisões clínicas do Dr. Getúlio) · **Data:** 2026-06-23

## Os 5 templates iniciais
| Template | Área | Itens |
|---|---|---|
| Avaliação Médica de Entrada (…104) | Medicina | 137 |
| Complemento da Avaliação Médica (…105) | Medicina | 299 |
| Avaliação Nutricional Inicial (…108) | Nutrição | 136 |
| Avaliação Psicológica Inicial (…110) | Psicologia | 135 |
| Avaliação Física Inicial (…112) | Ed. Física | 141 |

## Diagnóstico
**214 itens aparecem em ≥2 dos 5** (19 em 4 templates, 59 em 3, 136 em 2). Por domínio:

| Grupo (domínio) | Itens redundantes | Em até N templates |
|---|---|---|
| Composição corporal | 59 | 3 |
| Sono | 49 | 2 |
| Cognição | 38 | 2 |
| Social | 23 | 2 |
| Objetivos | 19 | 4 |
| Alimentação | 14 | 2 |
| Movimento e atividade física | 6 | 2 |
| Stress | 6 | 2 |

**Profissional errado:** os dois templates médicos carregam domínios de outras áreas:
- *Alimentação* (Água, Álcool, Açúcar, Frutas/Verduras, Intolerâncias, Suplementação, Padrão
  alimentar) → **Nutrição**.
- *Movimento e atividade física* → **Ed. Física**.
- *Composição corporal* → **Nutrição/Ed. Física**.
- *Sono, Cognição, Social, Stress* (sobretudo no Complemento) → **Psicologia**.
- *Objetivos* (19) repetidos em todos os iniciais → coletar **uma vez**.

Distribuição atual por grupo (resumo): Médica/Complemento têm Composição (13+46), Sono (10+39),
Cognição (7+31), Social (5+18), Alimentação (8+6), Movimento (2+4), Stress (3+3), Vida Sexual (20).
Psicológica, Nutricional e Física estão "limpas" (só seus domínios + Objetivos).

## Princípio de correção
**Cada domínio (grupo do score) pertence a UM template inicial.** O profissional dono coleta;
os outros não repetem. Isso resolve redundância e "profissional errado" de uma vez.

## Matriz de propriedade proposta (a confirmar)
| Domínio | Dono proposto | Remover de |
|---|---|---|
| Histórico de doenças | Médica (Entrada/Complemento) | — |
| Histórico Familiar de Doenças | Médica | — |
| Genética / Exames | Médica | — |
| Vida Sexual | Médica (Complemento) | — *(ou dividir c/ Psico?)* |
| Alimentação | **Nutrição** | Médica, Complemento |
| Composição corporal | **Nutrição** (medições) | Médica, Complemento, *Física?* |
| Movimento e atividade física | **Ed. Física** | Médica, Complemento |
| Cognição | **Psicologia** | Médica, Complemento |
| Social | **Psicologia** | Médica, Complemento |
| Stress | **Psicologia** | Médica, Complemento |
| Sono | **Psicologia** (comportamental) | Médica/Complemento *(ver decisão)* |
| Objetivos | **1 template só** (entrada médica?) | os outros 4 |

## Decisões que preciso de você (clínicas)
1. **Composição corporal** (peso, altura, IMC, circunferências, bioimpedância): dono único
   (Nutrição?) ou medido por Nutrição **e** Ed. Física (mantém nos dois, tira só da Médica)?
2. **Sono**: 100% Psicologia, ou o médico mantém um subconjunto clínico (apneia/ronco/SAOS)
   e a Psicologia fica com o comportamental (insônia/higiene)?
3. **Vida Sexual**: fica só na Médica (Complemento) ou divide com Psicologia?
4. **Objetivos**: coletar uma única vez — em qual template? (sugiro Avaliação Médica de Entrada,
   por ser o 1º contato.)
5. **Complemento Médico (299)**: depois de tirar o que é de outras áreas, ele encolhe bastante.
   Faz sentido manter Entrada + Complemento separados, ou fundir num só médico?

## Decisões travadas (Getúlio, 2026-06-23)
1. **Composição corporal: dividida Nutrição × Ed. Física, sem repetir** — itens focados em
   exercício/desempenho ficam na Física; medidas de composição (peso, IMC, % gordura,
   circunferências, bioimpedância) ficam na Nutrição.
2. **Sono: parte clínica fica no médico** (apneia/ronco/SAOS/sonolência) e o comportamental
   (insônia, higiene, cronotipo) vai pra Psicologia.
3. **Vida Sexual: dividida** médico × Psicologia (orgânico no médico, vínculo/desejo na Psico).
4. **Objetivos: só na Avaliação Médica de Entrada** (1º contato). Remover dos outros 4.
5. **Complemento Médico permanece.** As 5 preenchidas = escore completo (800+).

## Reformulação da Avaliação Médica de Entrada (camada de "encantamento")
**Conceito:** consulta médica inicial de ~40min que precisa **encantar** e gerar um **score
parcial de alto impacto** (radar AGIR já com os 4 eixos preenchidos, mesmo que parcial) pra
mostrar ao cliente — ferramenta de marketing/conversão. NÃO repete o que as outras 4 coletam:
os itens da Inicial são exclusivos dela; Complemento + Nutri + Psico + Física completam o resto.

**Princípios de seleção (alto valor de impacto):**
- **Cobrir os 4 eixos AGIR** pra o radar parcial não ficar "torto" (mesmo sem exames, que só
  saem depois). Fonte do radar = itens **com níveis** (escore vem de `selected_level`); priorizar
  itens pontuáveis, rápidos de coletar (sintoma/história/antropometria), **sem depender de labs**.
- **Ancorar nas dores principais** (site/deck): fadiga e baixa disposição, sono não recuperador,
  peso/composição, risco cardiometabólico, estresse/humor/ansiedade, libido/vida sexual,
  cognição/memória, intestino/GI, inflamação/imunidade recorrente, transições hormonais
  (peri/andropausa). Mensagem: "Normal ≠ Ótimo" / "achados em silêncio".
- **Inclui Objetivos** (a expectativa do cliente, base do plano e do "antes/depois").

**Cobertura proposta da Inicial por eixo AGIR (alto impacto, sem lab):**
| Eixo | Dor principal | Itens-âncora candidatos (com nível, rápidos) |
|---|---|---|
| **A** Atividade/Alimentação | peso, sedentarismo, padrão alimentar | cintura, IMC, nível de atividade, padrão alimentar geral, ultraprocessados/açúcar |
| **G** Gestão clínica/metabólica | risco cardiometabólico, GI, imune, hormonal | PA/antropometria, sintomas cardiometabólicos, queixa GI, infecções recorrentes, sintomas hormonais/menopausa-androp. |
| **I** Mente-corpo | fadiga, estresse, humor, libido, cognição | disposição/vitalidade, estresse percebido, humor (1-2 itens), libido, memória/foco |
| **R** Ritmo/Sono | sono não recuperador, ronco | qualidade do sono, ronco/apneia (clínico), sonolência diurna |

Tamanho-alvo: enxuta o suficiente pra 40min (estimativa ~40–70 itens pontuáveis, a calibrar),
todos contribuindo ao radar. Os ~137 atuais serão recurados (tirar texto/processo de baixo
valor, manter os de impacto, garantir ≥1 item por eixo/dor).

**Próximo passo (após você aprovar o conceito):** eu gero a **lista curada concreta** da Inicial
(item a item, por eixo/dor, marcando quais são novos/mantidos/saem) pra sua revisão linha a linha,
e só então aplico em dev.

## Execução (depois de aprovar a matriz + a curadoria da Inicial)
- Tudo é **dado** (`anamnesis_template_items`): DELETE dos vínculos do não-dono. Idempotente,
  reversível. **Dev primeiro, depois prod** (mesmo padrão da reorg de pilares).
- Eu gero o SQL + um relatório "antes/depois" (contagem por template) pra você conferir.
- Não mexe em score_items nem no motor de score — só em quais itens cada template puxa.
