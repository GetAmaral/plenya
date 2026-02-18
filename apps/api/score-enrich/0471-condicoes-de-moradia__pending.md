# ScoreItem: Condições de moradia

**ID:** `019bf31d-2ef0-7801-bb58-5cf2e4209288`
**FullName:** Condições de moradia (Social - Atual)

**Preparation Metadata:**
- Quality Grade: **FAIR**
- Total Chunks: 30 de 18 artigos
- Avg Similarity: 0.470

---

## Contexto

Você é um especialista em medicina funcional integrativa e está contribuindo com o **Escore Plenya** — um escore completo de análise de saúde que avalia todos os aspectos da saúde, performance e longevidade humana. Cada ScoreItem representa um parâmetro clínico, laboratorial, genético, comportamental ou histórico que compõe esse escore.

Seu papel é gerar conteúdo clínico de alta qualidade para enriquecer cada parâmetro do escore com relevância clínica, orientação ao paciente e conduta prática.

**Regras inegociáveis:**
- Use **apenas** o conhecimento médico real consolidado e os dados presentes nos chunks científicos abaixo
- **Não alucine, não invente** dados, estudos, estatísticas ou referências que não estejam nos chunks ou no seu conhecimento médico estabelecido
- Se um dado específico não constar nos chunks e não for do seu conhecimento consolidado, **não o inclua**
- Seja preciso: prefira omitir a inventar

## Instrução

Com base nos chunks científicos abaixo, gere as respostas em formato JSON.

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7801-bb58-5cf2e4209288`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7801-bb58-5cf2e4209288",
  "clinical_relevance": "Texto técnico para médicos (1000-5000 chars): definição fisiológica precisa, valores de referência e interpretação, fisiopatologia resumida, dados epidemiológicos com números concretos, estratificação de risco baseada em evidências.",
  "points": 0,
  "patient_explanation": "Texto simples para pacientes (500-1000 chars): o que é este parâmetro sem jargões, por que é importante para a saúde, o que valores alterados podem significar. Tom tranquilizador e educativo.",
  "conduct": "Conduta clínica em Markdown (1000-5000 chars): investigação complementar necessária, critérios de encaminhamento a especialistas, intervenções baseadas em evidências. Use bullet points, seções e negrito."
}
```

**Nota:** `points` deve ser `0` neste item — não calcule pontuação.

---

### Contexto Científico

**ScoreItem:** Condições de moradia (Social - Atual)

**30 chunks de 18 artigos (avg similarity: 0.470)**

### Chunk 1/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.503

desfavorável (poluição, mofo, poeira, pelos, químicos) e interno inflamado (alergias, intoxicações); comorbidades (alergia alimentar, refluxo, obesidade, anemia, rinite).
### 2. Conceitos diagnósticos e fisiopatologia segundo GINA
* Definição e sintomas
  - Asma: inflamação crônica das vias aéreas inferiores com obstrução reversível; sintomas: tosse, sibilância, aperto torácico, dispneia; padrão persistente (leve a grave) ou intermitente.
  - Desencadeantes: IgE (frequente: viroses) vs fenótipo neutrofílico (sem desencadeante claro).
* Confirmação diagnóstica
  - Limitação do fluxo (VEF1 e VEF1/CVF: <0,8 adulto; <0,9 criança).
  - Variabilidade: broncodilatador (≥12% e ≥200 ml adulto; ≥12% criança); PFE 2x/dia por 2 semanas (>10% adulto; >13% criança); resposta a tratamento; testes de desafio.
* Acompanhamento do controle
  - ACT (5 itens; 5–25 pontos) nas versões pediátrica e adulta.

---

### Chunk 2/30
**Article:** MFI - Psiquiatria 13 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.494

ico de uso de medicamentos psiquiátricos (o que usaram, como, por quanto tempo) e sintomas relacionados para mapear sistemas de neurotransmissores (glutamato, GABA, dopamina, etc.).
- [ ] Indicar terapias comportamentais aos pacientes com ansiedade, sugerindo troca de abordagem caso a terapia atual não esteja gerando resultados práticos.
- [ ] Recomendar uso de óleo de lavanda (ex.: 5 gotas na palma da mão para inalação, 3 vezes ao dia) como intervenção simples e eficaz para pacientes com ansiedade.
- [ ] Considerar associação de suplementos como magnésio, zinco, L-teanina, probióticos e adaptógenos no plano de tratamento da ansiedade.
- [ ] Investigar e abordar saúde do eixo intestino-cérebro em todos os pacientes com transtornos de humor, considerando neuroinflamação como fator causal.

---

### Chunk 3/30
**Article:** Remissão do Lúpus Através da Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.492

e sono propício,
- [ ] Pedir para o paciente terminar de comer, pelo menos duas a três horas antes de dormir
- [ ] Pedir para o paciente exercitar-se regularmente,
- [ ] Pedir para o paciente evitar cafeína, nicotina e álcool, principalmente perto do horário de dormir
- [ ] Pedir para o paciente manter um diário de sono,
- [ ] Avaliar os aplicativos e gadgets, que podem trazer informações de qualidade do sono
- [ ] Pedir para o paciente fazer uso de chás calmantes e relaxantes,
- [ ] Pedir para o paciente fazer uso de óleos essenciais,
- [ ] Revisar a dieta anti-inflamatória, em todas as consultas para ter o melhor resultado possível
- [ ] Revisar a realização de atividade física, em todas as consultas para ter o melhor resultado possível
- [ ] Rever a qualidade do sono, em todas as consultas para ter o melhor resultado possível
- [ ] Rever as ações que o paciente está fazendo para gerir o seu estresse, em todas as consultas para ter o melhor resultado possível
- [

---

### Chunk 4/30
**Article:** TDAH - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.488

ma) frequentemente atribuem problemas de atenção a TDAH quando o sono é um fator-chave a corrigir.
* Prioridade de intervenções
   - Antes de suplementos ou medicações, abordar rotinas de sono, tempo de tela, comunicação familiar e atividades físicas; corrigir ferro e outros fatores sem ajustar comportamento/sono não gera os resultados esperados na vida real.
### 6. Fatores sociais e risco de TDAH
* Renda familiar
   - Baixa renda durante o final da infância aumenta risco de TDAH em até 83%; renda média aumenta em 42% em comparação à linha de base.
   - Possíveis mediadores: menor tempo dos pais, maior carga laboral, mais pessoas em mesmo quarto, conflitos domésticos, alcoolismo, organização difícil e sono comprometido.
* Escolaridade materna
   - Baixa escolaridade materna aumenta a probabilidade de TDAH no final da infância em até 113%; escolaridade média aumenta em 42%.

---

### Chunk 5/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.488

3 min) ou HIIT (4 min/dia).
- [ ] 7. Melhorar a saúde bucal com raspagem de língua (cobre) e probióticos.
- [ ] 8. Avaliar e tratar fontes de inflamação crônica: infecções silenciosas (nasais, bucais), exposição a mofo e metais tóxicos. Investigar CIRS quando aplicável.
- [ ] 9. Para quem vai passar por cirurgia, utilizar o pool de suplementos sugerido para mitigar a neurotoxicidade da anestesia.
- [ ] 10. Discutir com um profissional de saúde a suplementação direcionada com base nos resultados da cognoscopia.

---

## SOAP

> Data e Hora: 2025-11-18 14:44:23
> Paciente:
> Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico:
2. Histórico de Medicação: Inserir mais aqui
## Subjetivo:
- Conteúdo educacional/apresentação sobre prevenção e manejo de risco para doença de Alzheimer, sem relato direto de queixas de um paciente específico.

---

### Chunk 6/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.485

, registro de medicamentos/estressores; repetição padronizada (3–5).
- Evidências: revisões sistemáticas/meta-análises e colaborações institucionais sustentam interpretação.
- Educação: bibliografia em medicina autonômica; acesso a abstracts via Academia Brasileira; capacitação em teoria polivagal e vias neuroendócrinas/neuroimunes.
## Exemplos e correlações clínicas
- Caso familiar com diabetes gestacional e componente emocional: necessidade de acompanhamento prolongado.
- Exemplo pós-COVID: broncoespasmo e deambulação difícil; proposta de fotobiomodulação em gânglio simpático da 1ª costela com broncodilatação e menor risco cardíaco.
- Perfis com baixa VFC e baixa reserva fisiológica: suspender exercício vigoroso até recuperar alostase.
## 📅 Next Arrangements
- [ ] Implementar protocolo de VFC com repetição padronizada (3–5 medições) em condições controladas.

---

### Chunk 7/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.481

Tarefas
- [ ] 1. Implementar uma dieta anti-inflamatória, livre de alérgenos, contaminantes e defensivos agrícolas.
- [ ] 2. Reduzir a exposição a gatilhos ambientais, incluindo poluentes, produtos químicos domésticos (ex: amaciantes), perfumaria e mofo.
- [ ] 3. Investigar e tratar possíveis intoxicações por metais pesados, como o arsênico.
- [ ] 4. Avaliar e corrigir os níveis de ferro, evitando tanto a deficiência (que mimetiza sintomas de asma) quanto o excesso (que é pró-inflamatório).
- [ ] 5. Considerar a suplementação de Vitamina K2 em pacientes em uso crônico de corticoides para prevenir a perda de massa óssea.
- [ ] 6. Manter os níveis de Vitamina D acima de 60 ng/ml através de suplementação diária, com atenção especial a crianças.
- [ ] 7. Avaliar o ômega-índex e suplementar ômega-3 para atingir níveis > 8%, especialmente em pacientes obesos, para reduzir a inflamação.
- [ ] 8.

---

### Chunk 8/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.480

g/kg) e prática de exercícios de resistência para preservar massa muscular.
- [ ] 3. Todos os profissionais: Em doenças crônicas sem causa orgânica clara ou com má resposta ao tratamento, investigar ativamente traumas de infância, estresse crônico e questões emocionais não resolvidas como possível "causa primeira".
- [ ] 4. Terapeutas e psicólogos: Adotar "terapia de precisão", utilizando múltiplas ferramentas e combinando diferentes abordagens terapêuticas para personalizar o tratamento e focar em resultados mensuráveis, em vez de seguir uma única linha teórica por longos períodos.
- [ ] 5. Estudo pessoal: Pesquisar o conceito de "causa primeira" de Aristóteles para aprofundar a lógica de buscar a origem dos problemas.
- [ ] 6. Estudo pessoal: Ler o livro de Bruce Lipton sobre a conexão entre mente e doenças físicas.

---

## SOAP

> Data e Hora: 2025-11-17 16:33:53
> Paciente: 
> Diagnóstico:

## Histórico do Diagnóstico:
1. Histórico Médico: 
2.

---

### Chunk 9/30
**Article:** TDAH - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.479

aliação bioquímica e nutricional antes de fechar diagnósticos de TDAH e comorbidades.
   - Considerar que “problemas de aprendizado” podem derivar de dieta rica em açúcar e deficiências vitamínicas/minerais.
### 8. Sono e arquitetura do sono
* Impacto do sono no comportamento
   - Sono insuficiente ou de má qualidade provoca desatenção, irritabilidade e impulsividade sem implicar TDAH.
   - Fatores: apneia do sono, respiração oral, deficiência de melatonina, exposição noturna à luz azul.
* Avaliação recomendada
   - Polissonografia ou monitoramento domiciliar (dispositivos de consumo) para parâmetros básicos (agitação, movimentos, respiração).
   - Melhorar o sono antes de confirmar diagnóstico pode alterar o quadro comportamental.
### 9.

---

### Chunk 10/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.478

Plano de Tratamento de Acompanhamento:**
    *   **Controle Ambiental:** Redução de mofo, poeira, pelos, produtos químicos.
    *   **Estratégia Terapêutica:** Preferir uso intermitente ou doses baixas de corticoide para preservar massa óssea e estatura.
    *   **Saúde Geral:** Tratamento de comorbidades (rinite, refluxo, obesidade) e vigilância para síndrome metabólica. Abordagem de medicina funcional (inflamação/metilação).
    *   **Educação:** Orientação sobre a doença, dispositivos e prevenção primária.

---

## Quantitative Data

### Narrativa Quantitativa
A asma é altamente prevalente globalmente e no Brasil, com grande impacto em crianças, gerando alto volume de atendimentos e internações, enquanto o controle clínico permanece insuficiente.

---

### Chunk 11/30
**Article:** Ritmo Circadiano Eixo HPA - Parte X (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.477

udo sobre Trabalho Noturno e Câncer de Mama**
    - Meta-análise de 61 estudos (≈115 mil mulheres): trabalho em regime noturno aumenta o risco de câncer de mama em 32% na população geral.
    - Em enfermeiras, o risco sobe a 58%, possivelmente por alto consumo de café, alimentação inadequada (pizza, hambúrguer, doces) e estresse elevado do ambiente noturno.
*   **Higiene do Sono e Rotinas Matinais**
    - Orientação de higiene do sono é fundamental para todos os pacientes, mesmo sem queixas, pois muitos não percebem a má qualidade do descanso.
    - Evitar eletrônicos perto da cama à noite (celulares — especialmente carregando — e relógios eletrônicos).
    - Exposição à luz natural logo ao acordar é essencial para regular o ritmo circadiano, pois as células são fotossensíveis.
    - Rotina matinal sugerida: abrir a janela para luz natural, orar/conectar-se com uma força maior, agradecer e pedir por um dia iluminado antes de olhar o celular.
### 2.

---

### Chunk 12/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.469

uso de filtros de água, idealmente por osmose reversa, é recomendado.
*   **Infecções Silenciosas e "Leaky Brain"**
    - A beta-amiloide possui ação antibiótica. O conceito de "leaky brain" (cérebro permeável) explica como toxinas e bactérias podem atravessar a barreira hematoencefálica, especialmente em pessoas com "leaky gut".
    - Infecções próximas ao cérebro (boca, seios nasais) devem ser tratadas agressivamente.
*   **Mofo e Toxinas Ambientais**
    - O mofo gera uma reação da proteína beta-amiloide. Recomenda-se o uso de filtros HEPA e a função "dry" do ar-condicionado.
*   **Saúde Nasal**
    - O nariz abriga um microbioma complexo. Probióticos nasais e soluções caseiras (líquido de fermentados) podem ser usados. Terapias como EDTA nasal, glutationa e NAC auxiliam na desintoxicação.
*   **Saúde Bucal**
    - Bactérias como a *Porphyromonas gingivalis* estão implicadas no Alzheimer.

---

### Chunk 13/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.467

.
## Subjetivo:
- Queixa principal: Infecções respiratórias recorrentes; secreção nasal diária há 4 meses; otalgia/otites em resfriados; constipação crônica com gases; despertares noturnos para mamadeira.
- Sintomas associados: Febre recorrente em alguns episódios; broncoespasmo em bronquiolite prévia; rinorreia persistente; irritabilidade em febre; dor de ouvido em otite.
- Alimentação inadequada com excesso de lácteos e farináceos e pouca variedade de vegetais, sem peixes/ômega-3, sugerindo disbiose, inflamação de baixo grau e possíveis carências nutricionais (vitaminas A, D, zinco, ferro).
- Exposição elevada em creche e por irmão mais velho.
## Objetivo:
- Critérios de infecção respiratória de repetição: >6 infecções/ano; >1/mês; >3 do trato respiratório inferior/ano.
- Achados relatados:
  - Radiografia com descrição leiga de “catarro no pulmão” (sem laudo formal).

---

### Chunk 14/30
**Article:** MFI Psiquiatria 04 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.466

# MFI Psiquiatria 04

**Source:** https://web.plaud.ai/share/91421764035206153::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-18 15:41:55
Local: [Inserir Local]
Instrutor: Vitor
## 📝 Resumo
Esta palestra, ministrada por Vitor, enfatiza a importância fundamental de uma abordagem funcional e integrativa em todas as áreas da medicina, com um foco especial na psiquiatria. O instrutor argumenta que a prática médica moderna não pode mais ignorar a análise de exames laboratoriais, a bioquímica, a nutrição e a profunda conexão intestino-cérebro. Através do prisma da psiquiatria, a palestra prova a necessidade de avaliar o sistema gastrointestinal, destacando como a microbiota, a inflamação e a dieta influenciam diretamente o humor e distúrbios como depressão, ansiedade e TDAH. Além disso, são abordados fatores cruciais como o eixo HPA, a função mitocondrial, a poluição eletromagnética e as toxinas ambientais.

---

### Chunk 15/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.466

ermeabilidade intestinal/BHE, disbiose e alterações metabólicas influenciam cognição; encefalopatias metabólicas são exemplos clínicos.
- Toxinas e íons metálicos
  - Acúmulo de ferro e exposição ao alumínio aumentam estresse oxidativo e disfunção neuronal.
- Citotoxicidade de neurotransmissores e estresse
  - Excesso de glutamato e cortisol relaciona-se a menor volume hipocampal; manejo de estresse é essencial.
- Metabolismo da glicose/insulina
  - Alzheimer como possível “diabetes tipo 3”; resistência à insulina e controle glicêmico são pilares.
- Infecções e gatilhos
  - COVID longo e infecções orais latentes podem desencadear/piorar queixas cognitivas.
### 3. Dimensões Psicossociais e Rotina
- Aspectos sociais e sentido de utilidade
  - Perdas e mudanças afetam cognição; utilidade e pertencimento (zonas azuis) são protetores.

---

### Chunk 16/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.466

ogamaglobulinemia transitória da infância.
    *   **Exposição:** Frequência em creche/escola, presença de irmãos mais velhos, moradia com muitas pessoas.
    *   **Ambientais:** Exposição à fumaça de cigarro e poluição.
    *   **Histórico:** Desmame precoce, menor nível socioeconômico.
*   **Diagnósticos Diferenciais**
    *   É crucial considerar outras condições além da imunodeficiência, como: sintomas alérgicos (rinite, asma), doença do refluxo gastroesofágico, e doenças de base como fibrose cística.
*   **Relação entre Alimentação, Inflamação e Infecções**
    *   O consumo excessivo de laticínios, industrializados e glúten pode estar relacionado a sintomas gastrointestinais (cólica, refluxo, diarreia, constipação) e infecções de repetição.
    *   A retirada do leite pode diminuir as infecções, não necessariamente por alergia, mas por reduzir um processo inflamatório crônico sistêmico.

---

### Chunk 17/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Oxidação 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.465

sfunção endotelial.
  - Melhoria: Orientar repetição de exame (intervalo e condições pré-analíticas).
### 5. Fatores de estilo de vida e ambiente que elevam ROS
- Causadores: cigarro, álcool, dieta pobre em nutrientes, sedentarismo, pesticidas, metais tóxicos, medicações, infecções; varicocele pode aumentar ROS.
- Leucocitose por inflamação crônica como sinal de processo ativo.
- Estresse oxidativo amplamente estudado em cardiologia e fertilidade (feminina e masculina).
- Sugestões de IA:
  - Organização: Dividir em “comportamentais”, “ambientais” e “clínicos”.
  - Métodos: Checklist de triagem de estilo de vida para uso ambulatorial.
  - Clareza: Micro-caso (varicocele + ROS alto).
  - Melhoria: Metas acionáveis (150 min/sem de exercício, cessação tabágica, dieta rica em antioxidantes).
### 6.

---

### Chunk 18/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.464

l Antes de dormir e ao acordar
- [ ] Furar o bloqueio dos 5 mil passos, para sair do sedentarismo Obrigação diária
- [ ] Certificar-se que seus sistemas de desintoxicação estão otimizados
- [ ] Usar glutationa lipossomal 250mg duas vezes ao dia, para otimizar os sistemas de desintoxicação
- [ ] Usar sulforafano 400 microgramas no dia a dia, para otimizar os sistemas de desintoxicação
- [ ] Utilizar o protocolo de EDTA, para retirar qualquer metal, especialmente o mercúrio
- [ ] Ter um filtro por osmose reversa, para filtrar metais e toxinas
- [ ] Tratar agressivamente qualquer infecção que ocorra no perímetro perto do cérebro
- [ ] Usar o filtro EPA e a função dry no ar-condicionado, para remover as condições da umidade e as toxinas ambientais
- [ ] Pedir pela internet os produtos Probiomax e o Restore, para a saúde do microbioma nasal
- [ ] Passar no nariz o líquido que fica de um fermentado, para melhorar o microbioma nasal e combater patógenos
- [ ] Fazer gluta

---

### Chunk 19/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.463

lina.
  - Observação: oscilações em glicemia de jejum/hemoglobina glicada pós-infecção; correlacionar com quadro clínico e evitar alarmismos.

## Ritmo Circadiano, Sono e Humor
- Sono e hábitos noturnos impactam eixo HPA e sintomas de humor/fadiga:
  - Vinho noturno, telas/tarde e deprivação de sono desregulam o ritmo circadiano.
- Diferenciar:
  - Depressão por neuroinflamação/eixo intestino-cérebro/dano mitocondrial versus desregulação circadiana primária.

## Neuroinflamação, Neurotransmissores e Mitocôndria
- Consequências da neuroinflamação:
  - Disrupção HPA, alteração do SNA, citocinas elevadas.
- Vias afetadas:
  - Quinureninas: aumento da via → menor serotonina; sintomas de irritabilidade/desânimo.
  - Receptores NMDA: excitotoxicidade glutamatérgica → dano neuronal e mitocondrial.
- Efeitos cognitivos e neurodegenerativos:
  - Diminuição do BDNF → piora de memória; agravamento de Alzheimer/Parkinson em vulneráveis.

---

### Chunk 20/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.462

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 21/30
**Article:** MFI Psiquiatria 04 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.461

istamina e inibidores da DAO.
### 4. Fatores Essenciais na Saúde Funcional e Integrativa
*   **Eixo HPA e Ritmo Circadiano**
    *   É impossível praticar a medicina moderna sem um conhecimento profundo sobre o eixo HPA e o ritmo circadiano.
    *   Um estudo mostrou que uma pequena redução na amplitude do ritmo circadiano associou-se a um risco aumentado de transtorno depressivo, bipolaridade e instabilidade de humor.
*   **Função Mitocondrial**
    *   As mitocôndrias são cruciais para a liberação de neurotransmissores, pois fornecem o ATP necessário para o recrutamento de vesículas sinápticas e para o tamponamento de cálcio.
    *   Parte da produção de neurotransmissores e hormônios ocorre nas mitocôndrias.
*   **Poluição Eletromagnética**
    *   É uma agressão física celular que precisa ser avaliada.

---

### Chunk 22/30
**Article:** Ritmo Circadiano Eixo HPA - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.460

m cápsula com óleo de coco fracionado) melhora qualidade do sono, principalmente em mulheres.
* Exercício físico
   - Melhora o sono; paciente deve se comprometer com prática regular.
   - Aeróbio é o mais eficaz para modular sono; melhor horário sugerido é 06:00, mas pode ser individualizado (alguns toleram treinos vespertinos sem prejuízo do sono).
### 6. Hábitos que interferem no sono e controle de estímulos
* Itens a avaliar com o paciente
   - Cafeína (café, chimarrão, tereré): horários e última dose.
   - Netflix/telas: duração, ajuste para luz amarelada/escura à noite.
   - Jantar: tipo de alimento e horário.
   - Álcool: evitar; apesar de sensação de melhora, piora fases do sono e reduz percepção de reparo.
   - Sons: reduzir volume/ruído à noite.
   - Rotina: após ~20:00, idealmente apenas higiene, banho, relaxamento.

---

### Chunk 23/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.458

sitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8. Alinhar expectativas com a família: foco em reduzir progressão, tentar estagnar e recuperar funcionalidade; definir cuidador(es) e dividir funções.
- [ ] 9. Revisar medicações que pioram sono e cognição (antipsicóticos, hipnóticos, benzodiazepínicos), buscando alternativas menos deletérias.
- [ ] 10. Modulação de fatores de risco: plano de atividade física, higiene do sono, manejo de estresse, melhora da ingestão proteica e redução de açúcares simples.
- [ ] 11. Avaliar e tratar disbiose intestinal e investigar infecções latentes (especialmente cavidade oral) que possam atuar como gatilhos.
- [ ] 12. Considerar fitoterápicos com titulação lenta e monitoramento de efeitos, especialmente os com ação anticolinesterásica; evitar polifarmácia e iniciar um por vez.

---

### Chunk 24/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.458

Mãe com asma e rinite alérgica persistente; pai ansioso; necessidade de higiene ambiental; fatores de risco gerais mencionados (idade <5 anos, creche, irmãos, múltiplas pessoas em casa, possível nível socioeconômico mais baixo).
2. Histórico de Medicação:
   - Passado: Inibidor de bomba de prótons (IBP); domperidona (Motilium) por ~2 meses; antibióticos múltiplos (4–6 cursos, muitos possivelmente desnecessários); anti-histamínicos; corticoides; anti-inflamatórios; produto transcrito como “Label” (não confirmado).
   - Atual: “Aditivo” 2 gotas (suplemento não especificado); Monteler 4 mg (provável montelucaste) para “imunidade” — não indicado pela médica devido a efeitos adversos neurológicos e gastrointestinais; recomenda suspensão.
   - Preferências/condutas gerais:
     - Analgésico/antitérmico: Dipirona (novalgina) preferida sobre paracetamol.

---

### Chunk 25/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.456

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 26/30
**Article:** Ritmo Circadiano Eixo HPA - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.454

l, inflamação de baixo grau mantém doenças crônicas.
- Consequências: redução de células NK, IgA secretora baixa, diminuição de bifidobactérias/lactobacilos (Gram-positivos), aumento de E. coli; maior risco de COVID grave, ITU, intestino irritável, gases, má digestão, infarto.
> Sugestões de IA
> - Separar “causas” vs “efeitos” para clareza; gráfico de barreira mucosa; exemplos de intervenções iniciais (sono, estresse, dieta anti-inflamatória, probióticos específicos); lista de sinais clínicos de hiperpermeabilidade.
### 8. Fases do estresse e interpretação da curva de cortisol
- Síndrome de adaptação geral:
  - Estresse agudo: elevação de cortisol que retorna em horas/dias.
  - Fase adaptativa: pico matinal reduzido, curva mais alta ao longo do dia.
  - Estresse crônico fases 1 e 2: pico matinal no limite inferior, vespertino/noturno elevado; sono prejudicado.
  - Fase 3: curva “flat” com níveis baixos todo o dia.

---

### Chunk 27/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.454

ças e 30% das mulheres adultas.
- Antagonistas de leucotrienos, usados no tratamento da asma, podem causar sintomas psiquiátricos em até 20% das crianças.
- Pacientes asmáticos em CTI apresentam uma alta taxa de colonização fúngica na pele (54%).

---

## Teaching Note

Data e Hora: 2025-12-09 04:55:32
Local: [Inserir Local]: [Inserir Local]
Aula: [Inserir Nome da Aula]: [Inserir Nome da Aula]
## Visão Geral
A aula abordou a abordagem funcional e integrativa no tratamento da asma, focando em suplementos, fitoterápicos e na modulação do sistema imunológico. Foram discutidos os papéis e evidências da Vitamina K2, Ferro, Magnésio, Vitamina D, Ômega 3, Quercetina, Cúrcuma e Boswellia Serrata, contrastando a plausibilidade bioquímica com os resultados de ensaios clínicos.

---

### Chunk 28/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.452

via SNA: base para superar dicotomia físico–mental; anamnese ampliada com timeline e matriz da Medicina Funcional; comunicação empática para engajamento.
- VFC como ferramenta central: indicador de alostase, resiliência e saúde global; orienta diagnóstico diferencial e decisões terapêuticas.
- Reprogramação do SNA: necessária em raízes da early life; abordagens multimodais neuroendócrinas/neuroimunes; hierarquia embriológica prioriza equilibrar SNA antes de ajustes dietéticos profundos.
- Pós-COVID: desautonomias correlacionadas a sequelas; intervenções focadas em SNA e VFC beneficiam sobreviventes; atenção a POTS/hipotensão neurogênica e digestão (ajuste de fibras).
## Boas práticas e padrões de qualidade
- Medição: ambiente controlado, consistência temporal, registro de medicamentos/estressores; repetição padronizada (3–5).
- Evidências: revisões sistemáticas/meta-análises e colaborações institucionais sustentam interpretação.

---

### Chunk 29/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.451

de Montelucaste (Montelair) para "melhorar a imunidade" é inadequado e possui muitos efeitos colaterais neurológicos (insônia, irritabilidade).
*   **A Função da Febre e sua Modulação**
    *   A febre é um mecanismo de defesa que ativa o sistema imunológico. Medicar a febre a temperaturas baixas (ex: 37,5°C) pode interromper esse processo benéfico.
    *   A recomendação é medicar o estado da criança (prostração, dor), não apenas o número no termômetro.
    *   Além da Dipirona, medicações como **Ingystol** (homotoxicologia) e **Erizidoro** (antroposofia) podem ser usadas para modular a febre sem suprimi-la, estimulando a autorregulação.
### 2. Fatores de Risco e Diagnósticos Diferenciais
*   **Fatores que Aumentam a Chance de Infecções**
    *   **Idade:** Hipogamaglobulinemia transitória da infância.
    *   **Exposição:** Frequência em creche/escola, presença de irmãos mais velhos, moradia com muitas pessoas.

---

### Chunk 30/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.450

2022 com 83 mil internações e 574 óbitos.
  - Observação crítica: possível viés em óbitos pós-2020 (pandemia/isolamento/CID), sem mudanças de política além do impacto da Farmácia Popular desde 2015 (queda de 7 para 3 óbitos/dia).
* Impacto pediátrico e controle
  - Principal causa de visitas à emergência; doença crônica mais comum na infância após rinite (“vias aéreas unidas”).
  - ISAAC: até 30% das crianças; PENSE: 23% dos alunos do 9º ano com sintomas.
  - Entre 1–4 anos: 1/3 das internações por asma.
  - Controle insuficiente: Brasil ~10% controlados (dados SBPT/2014); UE/Austrália ~50%.
  - Causas: baixa adesão, medo de medicação/efeitos, falta de orientação, percepção equivocada de “não necessidade” com poucos sintomas; ambiente externo desfavorável (poluição, mofo, poeira, pelos, químicos) e interno inflamado (alergias, intoxicações); comorbidades (alergia alimentar, refluxo, obesidade, anemia, rinite).
### 2.

---

