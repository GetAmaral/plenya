# Negócio de verdade
### Estudo — agosto/2026

Briefing revisado: não é mais "renda extra de R$ 5 a 10 mil com 30 min/dia". É um negócio, com
ativo, receita recorrente e valor de venda. Saúde está na mesa. O mundo está na mesa. Não há
barreira de idioma. Capital disponível abaixo de R$ 50 mil.

O arquivo das rotas descartadas e da alternativa de caixa rápido está em
[alternativas-avaliadas-2026-08.md](alternativas-avaliadas-2026-08.md).

---

## 1. A tese, em um parágrafo

Em 26 de agosto de 2026, daqui a **21 dias**, entra em vigor a Resolução CFM 2.454/2026, que
obriga toda instituição médica brasileira que use inteligência artificial a ter governança,
classificação de risco, auditoria, registro em prontuário, treinamento formal e relatório de
transparência. **Não há período de graça para sistemas já em uso.** O mercado que precisa se
adequar tem centenas de milhares de estabelecimentos e quase duas mil healthtechs, e quem vende
adequação hoje são dois fundadores não médicos com um produto ainda em desenvolvimento,
escritórios de advocacia vendendo parecer e integradores de TI vendendo infraestrutura.
**Ninguém nesse mercado é médico com registro ativo que também escreve o software e opera um
prontuário eletrônico em produção.** Você é a única pessoa que conheço que está simultaneamente
dos dois lados dessa norma.

---

## 2. Por que agora, e por que isso não é mais um "nicho regulatório" genérico

Na Parte 2 do estudo anterior eu matei três ideias regulatórias porque o comprador não tinha dor
real, não tinha orçamento, ou o mercado já estava servido. Esta passa nos três filtros, e é
importante ver por quê.

**A dor tem data e tem fiscal.** A norma foi publicada em 27/02/2026 e entra em vigor em
26/08/2026. A fiscalização é dos Conselhos Regionais de Medicina, e o responsável final é o
médico, pessoa física, com registro em jogo. Isso é qualitativamente diferente de multa
administrativa: coloca o CRM do sujeito na linha. Medo de processo ético-profissional move
orçamento muito mais rápido que medo de multa.

**Não há período de adaptação.** Diferente da reforma tributária, que deu ano de teste e
suspendeu penalidade, a resolução vale para sistemas em desenvolvimento **e já em uso** na data
de vigência. Um hospital que usa triagem algorítmica desde 2023 precisa estar conforme no dia
26, independentemente de quando comprou.

**O mercado já está exposto.** Sete em cada dez hospitais de médio porte já têm pelo menos um
algoritmo de apoio à decisão clínica em produção, segundo a Associação Brasileira de
Telemedicina (2025), e 17% dos médicos brasileiros já usam IA generativa na rotina, segundo a
TIC Saúde 2024. Ou seja, a maioria dos obrigados já está em não conformidade hoje, agora, sem
saber.

**A obrigação é recorrente, não um projeto único.** Monitoramento contínuo, auditoria periódica,
relatório de transparência, registro de treinamento e reavaliação a cada mudança de modelo. Isso
é assinatura, não serviço avulso. É a diferença entre consultoria e negócio.

### O que a norma exige, na prática

Consolidado das análises jurídicas e do checklist de implementação publicados:

1. **Inventário** de todas as soluções de IA em uso, incluindo as embutidas em sistemas
   contratados, integrações e módulos automatizados. Muita gente não sabe onde a IA está rodando.
2. **Classificação de risco** de cada sistema em baixo, médio, alto ou inaceitável, considerando
   impacto em direitos fundamentais, complexidade do modelo, grau de autonomia e sensibilidade
   dos dados.
3. **Avaliação de impacto algorítmico** para sistemas de risco médio e alto.
4. **Supervisão médica obrigatória.** Nenhuma decisão clínica relevante pode ser delegada sem
   médico avaliando, validando e registrando.
5. **Registro em prontuário**: identificação da ferramenta, finalidade de uso, revisão humana da
   sugestão, informação prestada ao paciente e eventual recusa.
6. **Direito do paciente** de ser informado e de recusar, com termo de consentimento revisado
   onde a IA for apoio relevante.
7. **Revisão de contratos** com fornecedores de IA, incluindo cláusulas de conformidade.
8. **Treinamento formal** dos médicos, com registro de participação.
9. **Relatório de transparência** de governança de IA, publicado.
10. **Auditoria especializada e monitoramento contínuo**, proporcionais ao grau de impacto.
11. **Alinhamento com a LGPD**, já que tudo isso roda sobre dado sensível.

Repare no item 5 e no item 8: os dois exigem que alguém mexa no prontuário eletrônico e no
processo clínico. Não é trabalho de advogado nem de integrador de rede.

---

## 3. Quem já está nesse mercado (checagem de concorrência)

| Quem | O que vende | Lacuna |
|---|---|---|
| **Certifq** | Plataforma de governança de IA em saúde: políticas, avaliação de risco, due diligence de fornecedor. Único concorrente direto encontrado. | **Dois fundadores: uma psicóloga com mestrado em saúde pública e um especialista em GRC com formação em filosofia. Produto ainda em desenvolvimento, recrutando beta testers.** Nenhum médico, nenhum engenheiro. |
| Escritórios de advocacia (M3BS, Mattos Filho, NPLaw) | Parecer jurídico, guia prático, revisão contratual | Não implementam. Entregam PDF, não conformidade operante. |
| Integradores de TI (Prolinx, Selbetti) | Infraestrutura, cibersegurança, governança de dados | Não têm leitura clínica. Não sabem classificar risco assistencial. |
| Consultorias de LGPD e DPO as a Service | Encarregado terceirizado, adequação de dados | Sabem privacidade, não sabem IA nem medicina. |

**Conclusão:** o mercado tem oferta jurídica e oferta de TI, e não tem oferta clínica-técnica. A
norma exige exatamente a interseção. Esse buraco não fica aberto muito tempo, mas está aberto
agora e ninguém com o seu perfil está nele.

---

## 4. Por que você, especificamente

Não é discurso motivacional, é lista de barreiras de entrada que você já pagou:

- **CRM ativo.** A norma responsabiliza o médico. Quem assina a classificação de risco
  assistencial precisa poder responder por ela. Consultor sem CRM não assina.
- **Você escreve o software.** Auditar sistema de IA exige ler comportamento de modelo, entender
  prompt, contexto, alucinação, versionamento e log. Você faz isso todo dia dirigindo agentes.
- **Você opera um prontuário eletrônico em produção.** O item mais chato da norma, que é
  registrar em prontuário o uso da IA com finalidade, revisão humana e ciência do paciente, é
  uma feature que você sabe especificar e construir porque já mantém as tabelas onde ela mora.
- **Você já vive LGPD, telemedicina, CFM e dado sensível** na operação da Plenya.
- **A Plenya é obrigada pela norma.** Sua primeira implementação é interna, sem risco comercial,
  e vira o estudo de caso e a demonstração.
- **Você fala com médicos** e tem audiência de médicos. O comprador desse produto é o diretor
  técnico, que é médico, e desconfia de vendedor que não é.

---

## 5. As quatro formas de monetizar, e qual eu recomendo

### Forma A — Adequação como projeto (serviço)
Inventário, classificação de risco, avaliação de impacto, pacote de políticas, revisão de termo
de consentimento, treinamento e relatório de transparência. Entrega em 2 a 3 semanas com o
trabalho pesado feito por agentes sobre um template seu.

Ticket: R$ 8.000 a R$ 30.000 por instituição, conforme porte. Bom para caixa imediato e para
aprender o problema real. Ruim como negócio: vende hora, não escala, morre quando você para.

### Forma B — Encarregado de IA Clínica como serviço (assinatura)
O modelo já validado ao lado: DPO as a Service no Brasil cobra R$ 1.500 a 3.000/mês para
microempresa, R$ 3.000 a 5.000 para pequena, R$ 5.000 a 8.000 para média e R$ 8.000 a 15.000
para grande. Um DPO interno custa R$ 20 a 35 mil por mês com encargos. O mercado já aceita esse
formato, o preço está ancorado e a norma acabou de criar a versão de IA dele.

Ticket: R$ 1.500 a 4.000/mês por clínica, R$ 6.000 a 15.000/mês por hospital. **Quatro clínicas
médias já batem a meta original de R$ 10 mil, com receita recorrente.**

### Forma C — Módulo de conformidade vendido a quem já tem os clientes (recomendado)
Em vez de vender para 50 mil clínicas uma a uma, vender para quem já as atende. Todo fornecedor
de prontuário eletrônico e todo escriba de IA no Brasil (Voa, Doclin, Escriba Médico, NAIA,
iClinic, Ninsaúde, Feegow, Amplimed e dezenas de outros) tem dois problemas a partir do dia 26:

1. **O produto deles é a IA que precisa ser classificada e auditada.** Eles precisam de
   documentação de conformidade para continuar vendendo, porque o cliente vai passar a exigir.
2. **O cliente deles precisa registrar em prontuário o uso da IA**, e isso é feature do sistema
   deles, não do cliente. Eles vão ter que construir. Ou comprar pronto.

Você vende o módulo: registro em prontuário no padrão da norma, classificação de risco,
gerador de relatório de transparência, trilha de auditoria, registro de treinamento. White
label, integrado por API.

Ticket: R$ 3.000 a 10.000/mês por fornecedor. **Dez fornecedores são R$ 30 a 100 mil por mês
com dez clientes para atender**, o que é a única forma desse negócio caber em pouco tempo seu.
E cria efeito de rede: quanto mais fornecedores adotam, mais o padrão vira o seu.

### Forma D — Selo de conformidade (o ativo de longo prazo)
Certificação anual auditada por médico, exibida pelo fornecedor de IA em saúde para vender a
hospitais. O hospital passa a exigir o selo dos fornecedores, os fornecedores passam a comprar o
selo, e a auditoria vira receita recorrente com barreira reputacional. É exatamente o modelo que
o CHAI está montando nos Estados Unidos com laboratórios de garantia certificados sob ISO 17025,
e que ainda não existe no Brasil.

Esse é o negócio que tem valor de venda de verdade, mas nasce das formas B e C, não antes.

**Recomendação: A para caixa nos primeiros 60 dias, C como produto principal, B em paralelo com
poucos clientes selecionados, D a partir do mês 12.**

---

## 6. Aritmética

| Cenário | Composição | Receita mensal |
|---|---|---|
| Mês 1-2 | 2 projetos de adequação (Forma A) | R$ 16-40 mil (não recorrente) |
| Mês 3-6 | 3 assinaturas B + 2 fornecedores C | R$ 12-25 mil recorrente |
| Mês 6-12 | 5 assinaturas B + 6 fornecedores C | R$ 30-70 mil recorrente |
| Mês 12-24 | + selo D com 15 certificados/ano | R$ 60-150 mil recorrente |

Custo de operação: infraestrutura que você já paga, mais contador. Margem acima de 85%. Capital
necessário: praticamente zero, o que resolve a restrição de menos de R$ 50 mil.

Valor de venda: negócio de compliance recorrente com contratos anuais negocia entre 3x e 5x
lucro anual. Em dois anos, uma operação a R$ 50 mil por mês de lucro é um ativo de R$ 1,8 a 3
milhões. Isso é o que diferencia negócio de emprego.

---

## 7. Expansão internacional (a razão de isso não ser um negócio de janela curta)

A objeção óbvia é que a corrida de agosto passa. Ela passa, e o que fica é maior:

- **Obrigação recorrente.** Monitoramento contínuo, relatório anual, novo modelo exige nova
  classificação. Cada atualização de um escriba de IA reabre o processo.
- **Marco Legal da IA (PL 2338)** aprovado no Senado em dezembro de 2024, em votação na Câmara em
  2026, com modelo de risco no estilo europeu e sanções de até R$ 50 milhões. Quando sair,
  amplia a mesma demanda para fora da saúde.
- **EU AI Act:** obrigações de alto risco valendo desde agosto de 2026, e sistemas regulados como
  dispositivo médico com prazo até agosto de 2027. **Nenhum organismo notificado havia sido
  designado até o início de 2026**, com custo de conformidade estimado em US$ 2 a 15 milhões por
  empresa. Mercado de red teaming e avaliação de IA em US$ 2,26 bilhões em 2026, com preços
  praticados de US$ 8 a 25 mil por auditoria pontual, US$ 50 a 150 mil por engajamento completo e
  a partir de US$ 5 mil por mês em teste contínuo.
- **Estados Unidos:** o CHAI está certificando laboratórios de garantia independentes com base na
  ISO 17025, e a Nature Medicine publicou em junho de 2026 um estudo mostrando que modelos
  genéricos superam IA clínica com liberação da FDA, expondo uma lacuna de validação que os
  reguladores não fecharam.

Ou seja: o método que você construir para atender a norma brasileira é o mesmo produto, com
outra régua, para um mercado global que paga em dólar de 10 a 30 vezes mais por engajamento. A
resolução do CFM é a cunha de entrada, não o teto.

---

## 8. Riscos, sem maquiagem

1. **A fiscalização pode demorar.** CRM não tem tradição de auditar tecnologia. Se ninguém for
   autuado, o medo esfria e a venda fica mais dura. **Mitigação:** priorizar quem tem exposição
   contratual e reputacional real (healthtech que vende para hospital, hospital com acreditação,
   operadora), não a clínica pequena que só teme o que já aconteceu com o vizinho.
2. **Virar consultoria disfarçada.** Se você aceitar todo projeto que aparecer, em três meses
   estará vendendo hora com outro nome. **Mitigação:** limite de projetos por mês desde o
   primeiro dia, e tudo que for feito duas vezes vira template ou código.
3. **Concorrência acorda.** As grandes consultorias e os escritórios vão entrar. **Mitigação:** a
   janela é de 12 a 18 meses para fincar padrão, e a defesa é a Forma C, porque quem virou
   infraestrutura de dez fornecedores não é trocado por um PowerPoint.
4. **Conflito com a Plenya.** Você vai vender para empresas que competem com a Plenya em algum
   grau. **Mitigação:** entidade separada, e cláusula de confidencialidade levada a sério. O
   próprio CHAI trata conflito de interesse entre laboratório e desenvolvedor como requisito de
   certificação. Trate igual desde o começo.
5. **Reputação.** Você vende para colegas e o assunto é medo regulatório. Se o tom for alarmista,
   queima. **Mitigação:** conteúdo técnico e substantivo, nada de "adeque-se ou perca o CRM".
6. **Tempo.** Isso não cabe em 30 minutos por dia nos primeiros meses. Cabe em 1 a 2 horas, e
   volta para 30 minutos quando a Forma C estiver rodando. É a troca que transforma renda em
   patrimônio, e ela precisa ser aceita conscientemente.

---

## 9. Plano de 90 dias

**Dias 1-7.** Ler a resolução inteira, artigo por artigo, e transformá-la em matriz de requisito
com evidência exigida. Esse documento é o ativo inicial, tudo nasce dele. Implementar na Plenya:
inventário de IA, classificação, registro em prontuário, relatório. Você vira o primeiro caso.

**Dias 8-21.** Publicar a matriz como conteúdo técnico gratuito, em português, antes do dia 26.
Quem publicar a referência prática na semana da vigência ganha a posição. Você já tem canal e
sabe escrever. Sem alarmismo, só utilidade.

**Dias 22-45.** Fechar 2 projetos de adequação (Forma A) com quem responder ao conteúdo. Preço
de entrada, escopo fechado, prazo de 3 semanas. Objetivo não é o dinheiro, é a matéria-prima.

**Dias 46-75.** Converter o que foi feito duas vezes em produto. Especificar o módulo da Forma C.
Abordar 10 fornecedores de prontuário e escriba de IA com a demonstração rodando na Plenya.

**Dias 76-90.** Fechar os 2 primeiros contratos de fornecedor. Estruturar PJ separada. Definir se
o selo (Forma D) entra no ano 1 ou no ano 2.

---

## 10. As outras teses que estudei e por que ficaram atrás

| Tese | Por que não é a primeira |
|---|---|
| Laboratório de garantia de IA clínica global (modelo CHAI, EU AI Act) | Mercado maior e ticket muito maior, mas exige credencial institucional, acreditação ISO 17025 e presença no mercado alvo. É a expansão natural do negócio brasileiro no ano 2, não a largada. |
| Roll-up de serviços automatizado por IA | Tese quente, com General Catalyst alocando US$ 1,5 bilhão e Thrive US$ 1 bilhão, e resultados reais (Long Lake com US$ 100 milhões de EBITDA em dois anos). Mas o modelo é comprar empresa de serviço, e isso exige capital que você não tem hoje. |
| Rede de evidência do mundo real na América Latina | Mercado de RWE global saindo de US$ 20 bilhões em 2026, LatAm ainda pequena em US$ 0,88 bilhão, e o Brasil tem lacuna regulatória documentada além de escassez de profissionais. É negócio de dado e de volume, precisa de rede de clínicas antes de valer algo. Guardar como opção do ano 3, aproveitando a base da Plenya. |
| Micro-SaaS global genérico em dólar | Cabe em 30 min/dia, mas não tem vantagem sua nenhuma. Mediana de 12 a 18 meses até US$ 10 mil e só 6% chegam lá. Perde feio para uma janela regulatória onde você é a pessoa certa. |

---

## Fontes

Resolução e obrigações: [CFM - Resolução 2.454/2026, texto oficial](https://sistemas.cfm.org.br/normas/arquivos/resolucoes/BR/2026/2454_2026.pdf) ·
[CFM - normatiza uso da IA na medicina](https://portal.cfm.org.br/noticias/cfm-normatiza-uso-da-ia-na-medicina/) ·
[Conjur - o que hospitais e clínicas precisam fazer até agosto](https://www.conjur.com.br/2026-mar-30/ia-na-saude-o-que-hospitais-e-clinicas-precisam-fazer-ate-agosto-com-a-resolucao-cfm-2-454-2026/) ·
[Saúde Digital News - agenda de conformidade](https://saudedigitalnews.com.br/31/03/2026/ia-na-saude-o-que-hospitais-e-clinicas-precisam-fazer-ate-agosto-com-a-resolucao-cfm-2-454-2026/) ·
[Selbetti - checklist de adequação](https://selbetti.com.br/blog/resolucao-cfm/) ·
[M3BS Advogados - o que muda para hospitais, operadoras e healthtechs](https://m3bs.com.br/resolucao-cfm-no-2-454-2026-o-que-muda-para-hospitais-operadoras-e-empresas-de-tecnologia-em-saude/) ·
[Conjur - responsabilidade civil](https://conjur.com.br/2026-ago-01/inteligencia-artificial-na-medicina-a-resolucao-cfm-2-454-2026-e-os-novos-contornos-da-responsabilidade-civil/) ·
[Futuro da Saúde - implementação desafia hospitais e clínicas](https://futurodasaude.com.br/cfm-resolucao-de-ia-na-medicina/) ·
[NPLaw - guia prático IA e saúde (PDF)](https://www.nplaw.com.br/wp-content/uploads/2026/04/NPLaw_IAeSaude_2026.pdf)

Mercado e concorrência: [Certifq](https://www.certifq.com.br/) ·
[Engenharia Biomédica - mapa de 1.900+ healthtechs brasileiras](https://engenhariabiomedica.com/artigos/healthtechs-mapa-brasil) ·
[CNES - base de estabelecimentos de saúde](https://cnes.datasus.gov.br/) ·
[Futuro da Saúde - investimento em healthtechs 2026](https://futurodasaude.com.br/investimentos-healthtechs-2026/) ·
[The Big Insights - preços de DPO as a Service](https://thebiginsights.com/custo-dpo-as-a-service-guia-de-precos/) ·
[Implementando a LGPD - adequação a partir de R$ 1.900/mês](https://www.implementandoalgpd.com.br/blog/adequacao-lgpd/)

Contexto regulatório ampliado: [ANVISA e SaMD, RDC 657/2022](https://engenhariabiomedica.com/artigos/samd-software-dispositivo-medico) ·
[Exame - Marco Legal da IA (PL 2338)](https://exame.com/inteligencia-artificial/marco-legal-da-inteligencia-artificial-pl-2338-o-que-muda-para-empresas-com-a-nova-lei/) ·
[Senado - PL 2338/2023](https://www25.senado.leg.br/web/atividade/materias/-/materia/157233) ·
[Trilateral Research - linha do tempo do EU AI Act](https://trilateralresearch.com/responsible-ai/eu-ai-act-implementation-timeline-mapping-your-models-to-the-new-risk-tiers) ·
[Cloud Security Alliance - lacuna de prontidão para agosto de 2026](https://labs.cloudsecurityalliance.org/research/csa-research-note-eu-ai-act-high-risk-compliance-deadline-20/) ·
[CHAI - certificação de laboratórios de garantia](https://www.chai.org/blog/chai-advances-assurance-lab-certification-and-nutrition-label-for-health-ai) ·
[Market.us - mercado de red teaming de IA](https://market.us/report/ai-red-teaming-services-market/) ·
[AI Vyuh - preços de red teaming 2026](https://security.aivyuh.com/blog/ai-red-teaming-pricing-2026/) ·
[Clinical Trial Vanguard - lacuna de validação, estudo Nature Medicine jun/2026](https://www.clinicaltrialvanguard.com/opinion/nature-medicines-june-2026-benchmark-study-reveals-general-purpose-llms-outperform-fda-cleared-clinical-ai-and-exposes-a-validation-gap-regulators-have-not-closed/)

Teses comparadas: [Capital and Clarity - roll-ups de IA da General Catalyst](https://capitalandclarity.substack.com/p/the-general-catalyst-behind-15-billion) ·
[L40 - AI roll-ups em 2026](https://www.l40.com/insights/ai-rollups) ·
[Fortune Business Insights - mercado de RWE](https://www.fortunebusinessinsights.com/real-world-evidence-solutions-market-107676) ·
[The Lancet Regional Health Americas - lacunas regulatórias de RWE no Brasil](https://www.thelancet.com/journals/lanam/article/PIIS2667-193X(25)00355-2/fulltext?rss=yes)
