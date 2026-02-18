# ScoreItem: Drogas ilícitas

**ID:** `019bf31d-2ef0-7295-958a-97935e4329ca`
**FullName:** Drogas ilícitas (Histórico de doenças - Hábitos e vícios nocivos (Questionar ativamente sobre uso passado ou atual):)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 19 artigos
- Avg Similarity: 0.519

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7295-958a-97935e4329ca`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7295-958a-97935e4329ca",
  "clinical_relevance": "Texto técnico para médicos (1000-5000 chars): definição fisiológica precisa, valores de referência e interpretação, fisiopatologia resumida, dados epidemiológicos com números concretos, estratificação de risco baseada em evidências.",
  "points": 1,
  "patient_explanation": "Texto simples para pacientes (500-1000 chars): o que é este parâmetro sem jargões, por que é importante para a saúde, o que valores alterados podem significar. Tom tranquilizador e educativo.",
  "conduct": "Conduta clínica em Markdown (1000-5000 chars): investigação complementar necessária, critérios de encaminhamento a especialistas, intervenções baseadas em evidências. Use bullet points, seções e negrito."
}
```

**Regras para `points` (1-50):**
- Baixo impacto clínico: 1-9 pts
- Alto impacto clínico: 10-19 pts
- Alto impacto em mortalidade: 20-50 pts
- Critérios: gravidade/mortalidade (40%), prevalência (30%), intervencionabilidade (30%)

---

### Contexto Científico

**ScoreItem:** Drogas ilícitas (Histórico de doenças - Hábitos e vícios nocivos (Questionar ativamente sobre uso passado ou atual):)

**30 chunks de 19 artigos (avg similarity: 0.519)**

### Chunk 1/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 09 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.594

prescrever a pacientes em antidepressivos ou ansiolíticos devido a possíveis interações desconhecidas.
*   **Mucuna Pruriens**
    - Fitoterápico ayurvédico com L-Dopa (levodopa), precursor direto da dopamina que atravessa a barreira hematoencefálica.
    - L-Dopa é convertida em dopamina pela Dopa descarboxilase.
    - Estudos focam em doença de Parkinson; também investigada em Alzheimer, ELA e AVC por ação neuroprotetora.
    - O instrutor relata ausência de grandes resultados em uso pessoal.
*   **Selegilina**
    - Fármaco antigo, inibidor de MAO, usado em Parkinson e considerado nootrópico.
    - Inibe degradação de dopamina; combinação com fenilalanina melhorou escores de depressão em estudo.
    - Doses baixas (2–2,5 mg) podem auxiliar memória, foco e atenção, sem os efeitos colaterais ou restrições alimentares (queijos, cerveja) típicos de doses altas de IMAO.

---

### Chunk 2/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 09 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.591

 lica.
    - Enfoque na sinergia: combinar suplementos para melhores resultados e evitar adaptação. O instrutor mescla fenilpiracetam, teacrina, N-acetil L-tirosina, neuroavena, KSM-66, rhodiola, ginseng e feniletilamina.
    - A eficácia depende do contexto individual: saúde mitocondrial, alimentação e qualidade do sono.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida]
## 📚 Próximos Arranjos
- [ ] Ler o livro "Dopamination" para aprofundar o entendimento sobre a dopamina e seus efeitos.
- [ ] Investigar, na avaliação de pacientes, histórico de exposição a estímulos dopaminérgicos (redes sociais, videogames, etc.) para mapear possíveis dependências e compulsões.
- [ ] Considerar, como primeira abordagem à performance cognitiva, nootrópicos (família racetam) ou suplementos como feniletilamina e mucuna pruriens, antes de recorrer a medicamentos como Ritalina ou Venvanse.

---

### Chunk 3/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.560

-teanina, Huperzine, Ginseng.
- Adaptação individual
  - Ajustar doses e frequência conforme resposta; introduzir um composto por vez.
### 9. Caso Prático e Abordagem Integrativa
- Aromaterapia e dieta
  - Óleo de gergelim com óleos essenciais, caldo enriquecido com colágeno; respeitar paladar e otimizar dieta para funcionalidade.
- Continuidade terapêutica
  - Uso de fitoterápicos, suplementos e, em próxima sessão, óleo de cannabis para otimização neurológica.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Implementar rastreio precoce em pacientes com queixas sutis (humor, sono, preferência por doces), incluindo PET-CT/FDG PET, ressonância funcional e biomarcadores no líquor quando indicado.
- [ ] 2. Solicitar exames laboratoriais para avaliar magnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3.

---

### Chunk 4/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.540

- [ ] Revisar medicações: 5-ARIs, psicotrópicos, estatinas, anti-hipertensivos, lisdexanfetamina; discutir alternativas e risco/benefício.
- [ ] Intervenção comportamental: reduzir/cessar tabagismo, álcool, maconha e outras drogas; educação sobre pornografia; técnicas de relaxamento para reduzir predominância simpática.
- [ ] Implementar plano alimentar centrado em proteínas e gorduras de qualidade, vegetais variados; reduzir ultraprocessados, farináceos, refinados e óleos de sementes ricos em ômega-6.
- [ ] Solicitar e corrigir deficiências: vitamina D (visar >40 ng/mL), folato; considerar suporte antioxidante (NAC, glicina, ácido glutâmico; vitamina C, AAL, selênio, vitamina E; riboflavina 100–200 mg/dia).
- [ ] Considerar arginina e L‑carnitina como adjuvantes; avaliar hipogonadismo e iniciar reposição de testosterona quando indicado e seguro.

---

### Chunk 5/30
**Article:** MFI - Psiquiatria 13 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.534

ico de uso de medicamentos psiquiátricos (o que usaram, como, por quanto tempo) e sintomas relacionados para mapear sistemas de neurotransmissores (glutamato, GABA, dopamina, etc.).
- [ ] Indicar terapias comportamentais aos pacientes com ansiedade, sugerindo troca de abordagem caso a terapia atual não esteja gerando resultados práticos.
- [ ] Recomendar uso de óleo de lavanda (ex.: 5 gotas na palma da mão para inalação, 3 vezes ao dia) como intervenção simples e eficaz para pacientes com ansiedade.
- [ ] Considerar associação de suplementos como magnésio, zinco, L-teanina, probióticos e adaptógenos no plano de tratamento da ansiedade.
- [ ] Investigar e abordar saúde do eixo intestino-cérebro em todos os pacientes com transtornos de humor, considerando neuroinflamação como fator causal.

---

### Chunk 6/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 09 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.530

o, sem os efeitos colaterais ou restrições alimentares (queijos, cerveja) típicos de doses altas de IMAO.
*   **Família Racetam (Nootrópicos)**
    - Suplementos voltados à performance neurológica; na prática clínica, resultados aquém do prometido.
    - **Piracetam:** Primeiro desenvolvido; dose de 1 g.
    - **Aniracetam:** Fornece energia com baixa estimulação; útil para falta de disposição com ansiedade.
    - **Fenilpiracetam:** Mais estimulante; o instrutor usa 150 mg duas vezes ao dia, 3–4 vezes por semana, em combinação com outras substâncias.
    - **Fasoracetam:** Indicado para TDAH e ansiedade.
*   **Outros Nutrientes e Sinergia**
    - **N-acetil L-tirosina:** Forma acetilada da tirosina, precursora da dopamina, com melhor passagem pela barreira hematoencefálica.
    - Enfoque na sinergia: combinar suplementos para melhores resultados e evitar adaptação.

---

### Chunk 7/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 17 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.527

e (polifenóis): faixa de 100 a 150 miligramas como modulador antioxidante/metabólico.
- Capsaicina/capsiate: 5 a 10 miligramas como agente termogênico e modulador do apetite/metabolismo.
**Estratégias comportamentais e nutricionais complementares podem modular apetite e sintomas, mas exigem uso criterioso.**
- Ácido hidroxicítrico (Citrimax/Garcinia cambogia): 500 mg antes das refeições, especialmente meia hora antes do jantar para controle de fome no final da tarde.
- Óleos essenciais cítricos por inalação: três a cinco gotinhas, com instrução de inalação profunda; não ingeríveis e de custo elevado.
**Achados epidemiológicos sugerem papel das vitaminas do complexo B em comportamento, reforçando a importância da qualidade dietética.**
- Análise transversal (2012) relacionou baixa ingestão de B1, B2, B3, B5, B6 e folato a maiores escores de comportamento externalizante.

---

### Chunk 8/30
**Article:** Psiquiatria Metabólica Funcional Integrativa 23 - Frederico Porto (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.524

tocôndrias.
*   **DHEA:** Neuroesteroide importante. A reposição em mulheres com deficiência pode melhorar o cansaço.
*   **Orotato de Lítio:** Suplemento que gera neurogênese. Pode potencializar antidepressivos, mas não deve substituir o carbonato de lítio no tratamento do transtorno bipolar.
### 7. Novas Fronteiras em Tratamentos Psiquiátricos
*   **Óleo de Cannabis (Sistema Endocanabinoide):** Representa uma nova fronteira na medicina. É importante que médicos aprendam a utilizá-lo, pois não possui altas doses de THC.
*   **Ketamina:** Considerada a maior novidade na psiquiatria dos últimos 30 anos. Atua no glutamato e BDNF. O uso é clínico (injetável) e, em baixas doses, pode ser usado em terapia assistida.
*   **Psicodélicos em Baixa Dose:** Drogas "egolíticas" que diminuem a rigidez do ego. Aumentam a capacidade de resolução de problemas ao mudar a perspectiva sobre eles, permitindo intervenção terapêutica eficaz.
### 8.

---

### Chunk 9/30
**Article:** TDAH - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.522

ca.
### 14. Mucuna pruriens (levodopa)
- Adjuvante com resultados limitados em TDAH; evidências mais robustas em Parkinson. Usar com cautela em casos selecionados.
### 15. Resistência insulínica, overnutrição e neurofunção
- Excesso calórico de baixa qualidade, sedentarismo e resistência insulínica afetam neurotransmissão, atenção, humor e sono; integrar manejo metabólico ao cuidado do TDAH.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida]
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Considerar avaliação nutricional completa: dieta; exames de ferro, ferritina, saturação de transferrina, zinco; vitaminas do complexo B (incluindo B12); homocisteína; e, se possível, metabolômica e microbioma intestinal.
- [ ] 2. Implementar rotina de refeições familiares: aumentar o jantar em pelo menos 10 minutos, retirar telas, incentivar mastigação lenta e degustação para melhorar saciedade e consumo de frutas/vegetais.
- [ ] 3.

---

### Chunk 10/30
**Article:** TDAH - Parte XXIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.520

strializado, evitando que alunos demonizem frutas indevidamente na orientação nutricional.
### 5. Abordagem Prática e Medicação no TDAH
- Reflexão crítica sobre "febre de diagnósticos" versus necessidade real de medicação.
- Reconhecimento de que a medicação é benéfica e necessária em alguns casos, especialmente para dar o "impulso" inicial para mudanças de estilo de vida.
- Estudo de caso (filho de amigo): perfil "COMT rápida" (busca por novidade/dopamina; dificuldade de longo prazo).
- Efeitos colaterais: disfunção erétil/perda de libido com metilfenidato em jovens, levando ao abandono do tratamento. Ajuste para Venvanse (lisdexanfetamina) em doses seguras.
- Importância de gerenciar expectativas (paciente que abandonou o tratamento por buscar resultados irreais apenas com medicação).
> **[Sugestões da IA]**
> Casos clínicos reais foram o ponto alto, aproximando teoria e prática.

---

### Chunk 11/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.520

berberina, cromo e canela ajudam. A cetose pode ser alcançada com a redução de carboidratos (25-50g/dia).
*   **Uso de Canabinoides (CBD e THC)**
    - O CBD é indicado para ansiedade e o THC para agitação, insônia e inapetência. Ambos reduzem estresse oxidativo, inflamação e formação de beta-amiloide.
*   **Suporte Neuronal e Cognitivo**
    - **Sinalização Neurotrófica (BDNF, NGF):** Cogumelo juba de leão, magnésio treonato, zinco.
    - **Sinalização Colinérgica:** Citicolina, alfa-GPC, huperzina A.
    - **Memória:** Colinas, Bacopa monnieri, Ginkgo biloba, maca.
    - **Foco e Atenção:** L-teanina, cafeína, fosfatidilserina.
*   **Saúde Mitocondrial e Circulação**
    - **Mitocôndrias:** Coenzima Q10, PQQ, L-carnitina, ácido alfalipoico.
    - **Circulação:** Ginkgo biloba, Vinpocetina.
*   **Vitamina D e Reposição Hormonal**
    - **Vitamina D:** Níveis devem ser mantidos acima de 50.

---

### Chunk 12/30
**Article:** TDAH - Parte XV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.520

e educação do paciente.
- [ ] 6. Preparar síntese para a próxima aula: identificar estudos relevantes sobre uso de fitoterápicos e adaptógenos no TDAH, com qualidade metodológica e resultados principais.

---

## Teaching Note

> Data e Hora: 2025-12-09 05:00:40
> Local: [Inserir Local]
> Aula: [Inserir Nome da Aula]
## Visão Geral
A aula abordou o impacto histórico do Relatório Flexner na educação médica e na adoção do modelo biomédico, discutindo a marginalização de práticas consideradas não científicas. Em seguida, explorou-se a história e o uso de plantas medicinais em diversas culturas (Egito, Suméria/Babilônia, Índia/Ayurveda, China) e conectou-se esse legado a exemplos modernos de fitoterápicos/adaptógenos como ginseng, rodiola e própolis. Também houve exibição/comentário de um vídeo crítico sobre a evolução dos psicofármacos e a psiquiatria.
## Conteúdo Não Coberto
1.

---

### Chunk 13/30
**Article:** MFI PÓS RACHEL GAIGER AULA 01 - OXIGÊNIO HIPERBÁRICO (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.518

cofatores mitocondriais (CoQ10, L-carnitina, ácido alfa-lipoico, glutationa, glicina, taurina, tirosina, MCT, vitaminas B1/B2/B3).
- Integração com terapias padrão e medicina funcional: redução de sintomas para permitir abordagem da causa raiz; sinergia com antibióticos/antifúngicos, esteroides em DII e reabilitação motora/cognitiva em casos neurológicos.
### Segurança, contraindicações e manejo de efeitos adversos
- Contraindicações relativas: doenças pulmonares crônicas (enfisema), pneumotórax recente, cirurgia cardíaca/trauma recente, neurite óptica, uso atual de bleomicina/doxorrubicina, otite média/dificuldade de equalização, implante coclear, marcapasso (avaliação individual), hipercapnia, hipertensão não controlada, transtornos convulsivos.
- Absolutas: pneumotórax, broncoespasmo agudo não resolvido, próteses orbitais específicas.
- Drogas incompatíveis: cisplatina, doxorrubicina; cautela com amiodarona e antiangiogênicos.

---

### Chunk 14/30
**Article:** TDAH - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.517

dade do sono.
- [ ] 6. Introduzir exercícios físicos regulares para modular HPA e inflamação de baixo grau.
- [ ] 7. Avaliar e corrigir deficiências nutricionais: vitamina D (com dosagem e reposição), ferro, complexo B (especialmente B6/P5P), suporte mitocondrial.
- [ ] 8. Priorizar hierarquia terapêutica: intestino, HPA, sono, nutrientes; só então considerar fitoterápicos (ex.: bacopa) e nootrópicos.
- [ ] 9. Avaliar polimorfismos genéticos relevantes (SLC6A3/DAT1, ALDH2, MAO-B; futuramente COMT) para personalização em sintomas dopaminérgicos.
- [ ] 10. Considerar uso de selegilina (MAO-B) em candidatos com polimorfismos/sinais de acúmulo de aldeídos catecólicos; iniciar em 1 mg e titular até 5 mg conforme resposta.
- [ ] 11. Solicitar painel de metabolômica urinária com DOPAC e HVA (ou líquor quando aplicável), interpretando níveis no contexto enzimático.
- [ ] 12.

---

### Chunk 15/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 17 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.517

, e educar sobre o mito do “sugar rush”.
- [ ] 2. Implementar rotina de exercícios físicos conforme perfil: iniciar com aeróbicos para ansiosos; considerar HIIT para biogênese mitocondrial; ajustar intensidade gradualmente.
- [ ] 3. Avaliar sinais de resistência insulínica e métricas do eixo HPA; integrar monitoramento e intervenção clínica.
- [ ] 4. Introduzir suplementação de complexo B quando indicado, priorizando B1, B2, B3, B5, B6 e folato com base em necessidades e histórico nutricional.
- [ ] 5. Prescrever creatina monoidratada (preferencialmente Creapure): 2–3 g/dia para saúde cerebral; 5 g/dia para vegetarianos/veganos; orientar consumo imediato após preparo e uso diário.
- [ ] 6. Considerar acetil-L-carnitina em transtornos depressivos como adjuvante, com base na evidência de meta-análise.
- [ ] 7.

---

### Chunk 16/30
**Article:** TDAH - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.515

ar testes genéticos para COMT (Val/Val vs. Met/Met), MAO, tirosina hidroxilase, DBH, ALDH2, HCRT1/2 e HCRTR1/2.
- [ ] 2. Realizar análise de neurotransmissores/metabólitos urinários: 3-MT, DOPAC, HVA; considerar 3-MT em LCR e sangue se aplicável.
- [ ] 3. Avaliar sono noturno (qualidade, REM e profundo) antes de considerar modafinil; corrigir distúrbios de sono primariamente.
- [ ] 4. Considerar metilfenidato quando predomina desatenção e o perfil sugere benefício.
- [ ] 5. Testar modafinil em fadiga diurna/hipoalerta com suspeita de baixa tonicidade de orexinas, após excluir causas de sono ruim.
- [ ] 6. Avaliar bupropiona em TDAH com apatia/anedonia e baixa dopamina, reconhecendo resultados modestos.
- [ ] 7. Implementar L-tirosina (500–1.000/1.500 mg) e P5P (5–30 mg), monitorando homocisteína para evitar excesso de metiladores.
- [ ] 8. Otimizar nutrientes metiladores (B12, B9, magnésio, colina, P5P) e considerar SAM conforme perfil genético/metabólico.
- [ ] 9.

---

### Chunk 17/30
**Article:** Emagrecimento XVII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.512

no (ex.: nidra), dentro da psiquiatria metabólica; avaliar custo-benefício.
### 16. Protocolos injetáveis auxiliares (venosos e intramusculares)
- Princípios éticos: possibilidade, não necessidade; transparência e benefício.
- Logística: 1–3x/semana; viabilidade, custo e tempo.
- Componentes: metilfolato (cautela), 5-HTP (distância de 4h de SSRIs; risco de síndrome serotoninérgica; preferência IM), glutationa, ácido alfa-lipóico, glicina, carnitina; cromo/vanádio e niacinamida com menor impacto semanal.
- Estrutura sugerida: ~8 sessões, idealmente 2x/semana.
### 17. Ciclagem de estratégias dietéticas e individualização
- Proposta de ciclo: low-carb → mediterrânea → jejum intermitente → cetogênica → mimicking diet → plant-based, conforme meses e realidade do paciente.
- Doses, medicações e protocolos individualizados; encontrar o modelo que “cabe na vida”.
### 18.

---

### Chunk 18/30
**Article:** TDAH - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.509

quência e a lógica da prática clínica. A forma final do conceito não é apenas uma lista, mas um sistema de dependências: a eficácia de uma intervenção na "copa" da árvore (ex: um fitoterápico) depende inteiramente da saúde das "raízes" (os fundamentos metabólicos). Isto explica a falha de muitos tratamentos e "abre a porta" para uma prática mais rigorosa, sequencial e personalizada, onde a otimização da base fisiológica, guiada por biomarcadores, precede e potencializa qualquer tratamento sintomático.
**Rasto de Evidência:**
> Melhor? Quem disse que a copa vai ser a melhor para a TDAH? Se você não estiver hierarquicamente controlado... Modulação intestinal, eixo HPA, o sono, nutrientes, mitocôndrias. Você não vai ter função, você não vai ter resultados.

---

### Chunk 19/30
**Article:** TDAH - Parte XV (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.508

 ão clínica e reaprendizado de metodologias integradas ao estilo de vida.
### 6. Encaminhamentos para a próxima aula
* Foco
   - Retomar estudos e potenciais terapêuticos de plantas e abordagens para condições crônicas.
   - Contextualização específica: potencial no TDAH (Transtorno de Déficit de Atenção e Hiperatividade).
   - Transição da história para aplicação prática com evidências e protocolos.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Assistir ao vídeo de 8 minutos recomendado sobre psiquiatria, psicofármacos e regulação (conforme indicado pelo instrutor) e refletir sobre impactos na prática clínica.
- [ ] 2. Pesquisar e revisar fontes históricas mencionadas: Código de Hamurabi, Papiro de Ébers, Sushruta Samhita, Charaka Samhita, Ben Cao/Pengzao e lenda de Shen Nong, destacando usos de plantas específicas (ópio, gálbano, meimendro, efedra).
- [ ] 3.

---

### Chunk 20/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.508

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 21/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 08 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.507

tar tolerância e perda do “efeito surpresa”; utilizar estrategicamente em momentos de necessidade.
   - Vício e dependência: relatado que “o café vicia mais do que a maconha” no contexto da fala—prevenir uso diário compulsivo; cefaleia de privação pode ocorrer.
   - Doses e formas: prescrever cafeína anidra 100–200 mg para quem não consome café regularmente; preferir cápsulas oleosas de liberação mais lenta (ex.: com óleo de cártamo).
   - Integração em fórmulas manipuladas: incluir cafeína anidra 100–200 mg e ajustar conforme resposta; para quem já consome café, suplementar cafeína pode ser desnecessário; considerar mistura com outros componentes para modulação de efeitos.
* Benefícios e riscos sistêmicos
   - Positivos: efeitos neurocognitivos, possíveis benefícios metabólicos e sistêmicos (pulmões, fígado, rins—com moderação).

---

### Chunk 22/30
**Article:** TDAH - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.506

res de bomba de prótons e diuréticos) e sintomas comportamentais devem ser avaliados.
   - **Combinações Eficazes:** Bons resultados com magnésio, ômega-3, vitamina D e zinco.
* **Fisiopatologia Multifatorial do TDAH**
   - Fatores de risco: baixo peso ao nascer, exposição pré-natal a toxinas (álcool, nicotina, chumbo, pesticidas), questões educacionais e familiares.
   - Fatores maternos e genéticos: dieta materna, toxinas, polimorfismos genéticos e hereditariedade.
   - Fatores individuais e neurológicos: idade, gênero, status socioeconômico, predisposições de neurotransmissores, inflamação e desordens associadas (distúrbios do sono, depressão, ansiedade).
* **Crítica à Abordagem Médica Convencional**
   - O instrutor critica veementemente os médicos que negam a influência da alimentação no TDAH e em outras condições de saúde, classificando tal atitude como irresponsável e antiética.

---

### Chunk 23/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.502

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

### Chunk 24/30
**Article:** TDAH - Parte XXV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.500

sco de acúmulo) ou rápidos.
- **Fatores Externos:** pH urinário (acidez aumenta eliminação), disfunção hepática (esteatose) e uso crônico (depleção de glutationa) afetam o metabolismo.
- **Lacunas em Estudos:** Estudos não costumam avaliar o impacto de polimorfismos, carga de metais pesados, inflamação crônica, déficits nutricionais (vitaminas B, zinco, magnésio) ou microbioma intestinal.
- **Interações Medicamentosas:** Anfetaminas interagem com anticoncepcionais orais, inibidores de bomba de prótons e antidepressivos (inibidores de CYP2D6), alterando seus níveis plasmáticos.
> **Sugestões da IA**
> Esta seção foi extremamente rica e crítica, conectando a farmacologia com a nutrição e a genética, o que é um grande diferencial do curso. Você fez um excelente trabalho ao apontar as lacunas na pesquisa tradicional.

---

### Chunk 25/30
**Article:** TDAH - Parte XXV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.499

ão e funções executivas.
- **Perfil de Segurança:** Apresenta menor potencial de abuso e dependência em comparação com os psicoestimulantes.
> **Sugestões da IA**
> A explicação da atomoxetina foi concisa e clara, destacando bem sua principal diferença em relação aos psicoestimulantes. A organização foi lógica, apresentando-a como uma alternativa com um mecanismo mais simples. A explanação foi muito boa e não requer grandes ajustes.
### 4. Metabolismo Hepático, Fatores de Influência e Interações
- **Vias Metabólicas (Anfetaminas):** Metabolizadas pelo fígado (via CYP2D6, CYP2C) e flavina monooxigenase, com excreção renal dependente do pH urinário.
- **Variações Genéticas:** O gene CYP2D6 é altamente polimórfico, o que pode levar a metabolizadores lentos (risco de acúmulo) ou rápidos.
- **Fatores Externos:** pH urinário (acidez aumenta eliminação), disfunção hepática (esteatose) e uso crônico (depleção de glutationa) afetam o metabolismo.

---

### Chunk 26/30
**Article:** Genética e Epigenética II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.499

nas que impactem acetilação de histonas, metilação e reparo/dano do DNA.
- [ ] 2. Implementar estratégias para aumentar AGCC (fibras fermentáveis, modulação da microbiota) com protocolos de prescrição e monitoramento.
- [ ] 3. Avaliar status mitocondrial (sinais clínicos, exames indiretos) e intervir em cofatores (NAD/B3, FAD, alfa-cetoglutarato) conforme necessidade e segurança.
- [ ] 4. Em oncologia (p.ex., quimioterapia), monitorar homocisteína e manter doadores de metil em níveis normais; documentar racional e acompanhamento.
- [ ] 5. Para depressão refratária, considerar metilfolato em doses altas (200–1.000 mcg, podendo 2.000 mcg; em casos específicos, titulação até 15 mg), com monitoramento clínico e laboratorial.
- [ ] 6. Elaborar planos de exercício individualizados: definir faixas de FC, escolher modalidades (força/resistência) que promovam remodelamento muscular e biogênese mitocondrial; ajustar conforme nutrição e status hormonal.
- [ ] 7.

---

### Chunk 27/30
**Article:** MFI - Psiquiatria 13 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.497

como Silexan) é apresentado como alternativa eficaz e segura aos benzodiazepínicos (como Lorazepam) para transtorno de ansiedade generalizada, conforme estudo clínico de 2010.
    - Outras estratégias: magnésio, probióticos, zinco, adaptógenos e L-teanina.
*   **Uso Correto de Medicamentos**
    - Medicamentos são úteis e devem ser prescritos quando necessário, especialmente em casos graves de depressão.
    - A medicação isolada raramente resolve a causa raiz; em casos graves pode “zumbificar” o paciente.
    - “Remédio bom é aquele que entra e sai”: remedeia a situação e depois é descontinuado.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Próximas Providências
- [ ] A partir de 19 de novembro de 2025, começar a perguntar aos pacientes sobre histórico de uso de medicamentos psiquiátricos (o que usaram, como, por quanto tempo) e sintomas relacionados para mapear sistemas de neurotransmissores (glutamato, GABA, dopamina, etc.).

---

### Chunk 28/30
**Article:** TDAH - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.496

A suplementação nutricional demonstra uma eficácia notável no tratamento do TDAH, com uma resposta ao tratamento com multinutrientes (54%) sendo três vezes superior à do placebo (18%).**
- Um estudo com crianças a partir de 9 anos mostrou melhoras significativas nos sintomas do TDAH após 10 semanas de tratamento com uma mistura de ácidos graxos poliinsaturados, como o ômega-3.
- A deficiência de ômega-3 é frequentemente observada em pacientes com TDAH, e sua suplementação tem mostrado melhoras significativas.
- A enzima glutamato descarboxilase, crucial para a função neurológica, é dependente da vitamina B6 (piridoxal 5-fosfato), cuja suplementação é sugerida em doses de até 30 mg.
- A suplementação com Mucuna pruriens, em doses de 500 mg, é indicada para obter resultados, podendo ser usada até duas vezes ao dia.

---

### Chunk 29/30
**Article:** TDAH - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.496

ória universal: não existe; é preciso medir e personalizar.
- Biomarcadores: cortisol salivar/urinário, IL-6, TNF-α, PCR-us, cofatores nutricionais guiam condutas.
### 5. Revisão das Monoaminas
- Serotonina: regulação de humor, sono, ansiedade, dor; desbalanços em transtornos afetivos e dor crônica.
- Histamina: vigília, motivação, apetite; papel imunológico e alérgico.
- Dopamina: recompensa, foco, motivação; disfunções em TDAH, Parkinson, vícios; ação de psicoestimulantes.
- Epinefrina/norepinefrina: luta-fuga, atenção e foco; alvo de ISRSN e estimulantes (ex.: Atenta).
### 6. Síntese, armazenamento e receptores da dopamina
- Substratos e cofatores: fenilalanina/tirosina; TH (ferro, BH4, O2), AADC (P5P), sensibilidade da BH4 à oxidação; vitamina D aumenta expressão da TH.
- Armazenamento/liberação: VMAT2, liberação dependente de cálcio.

---

### Chunk 30/30
**Article:** TDAH - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.495

hemoglobina em crianças.
### 12. Metabolismo de B6/PLP, GABA, glutamato e vias do triptofano
- GAD e dopa descarboxilase dependem de PLP; disbiose desvia triptofano para indóis (ativação AhR), aumentando excitotoxicidade e “leaky gut”.
- Via das quinureninas: dependência crítica de PLP/zinco; deficiência aumenta radicais livres e neurotoxicidade (ácido quinolínico).
- B6 sanguínea não é fidedigna; preferir inferências por metabolômica, enzimas, homocisteína e sinais clínicos.
### 13. Genética, barreiras e resposta ao tratamento
- Polimorfismos em LPHN3 (dopamina, glutamato) e CDH13 (neuroplasticidade, barreiras) influenciam suscetibilidade e resposta.
- Estratégias: proteger barreiras intestinal/hematoencefálica; nutricional e estilo de vida modulam expressão gênica.
### 14. Mucuna pruriens (levodopa)
- Adjuvante com resultados limitados em TDAH; evidências mais robustas em Parkinson. Usar com cautela em casos selecionados.
### 15.

---

