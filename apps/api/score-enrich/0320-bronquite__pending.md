# ScoreItem: Bronquite

**ID:** `c77cedd3-2800-7707-8985-9beb91d0768c`
**FullName:** Bronquite (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 11 artigos
- Avg Similarity: 0.588

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7707-8985-9beb91d0768c`.**

```json
{
  "score_item_id": "c77cedd3-2800-7707-8985-9beb91d0768c",
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

**ScoreItem:** Bronquite (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual):)

**30 chunks de 11 artigos (avg similarity: 0.588)**

### Chunk 1/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.635

eções.
    *   **Fitoterápicos:** **Pelargonium sidoides** (Caloba, Imunoflan) diminui a replicação viral, a duração e a intensidade da doença.
    *   **Homeopatias:** **Corizalia** para coriza inicial e **Oscillococcinum** para quadros gripais.
    *   **Suplementação na Fase Aguda:** N-acetilcisteína (NAC), própolis verde, e uso curto (3-5 dias) de zinco, vitamina D e A (Ad-til) se os níveis não forem conhecidos.
### 4. Saúde Intestinal e Estratégias de Modulação
*   **Investigação Laboratorial**
    *   Solicitar: Vitamina D, A, Zinco (eritrocitário), perfil de ferro, hemograma, B12. Considerar dosagem de imunoglobulinas e prick test para ácaros.
*   **Lisados Bacterianos (Broncho-Vaxom)**
    *   Estimula o sistema imunológico contra as principais bactérias respiratórias. O tratamento padrão é de 10 dias/mês por 3 meses.
*   **Zinco para Infecções e Diarreia**
    *   O uso rotineiro (10-15 mg/dia) reduz a recorrência de infecções respiratórias.

---

### Chunk 2/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.628

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

### Chunk 3/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.616

esposta a tratamento; testes de desafio.
* Acompanhamento do controle
  - ACT (5 itens; 5–25 pontos) nas versões pediátrica e adulta.
  - Critérios GINA (4 itens/4 semanas): 0 = controlada; 1–2 = parcialmente; 3–4 = não controlada.
### 3. Risco de remodelamento e progressão
* Inflamação subclínica persistente + broncoespasmo levam a destruição epitelial e remodelamento brônquico, com irreversibilidade e evolução para DPOC.
### 4. Terapêutica tradicional por faixa etária (steps GINA) e adesão
* Princípios
  - ICS e broncodilatadores de curta/longa ação conforme steps; doses baixa/média/alta por tabelas; LABA e eventual LAMA.
  - <5 anos: preferência por baixa dose; se necessário, dobrar (alta dose).
* Adesão em adolescentes
  - Dificuldade elevada; responsabilidade deve ser compartilhada e conduzida pelos pais.
### 5.

---

### Chunk 4/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.612

s/condutas gerais:
     - Analgésico/antitérmico: Dipirona (novalgina) preferida sobre paracetamol.
     - Mucolítico: N-acetilcisteína (Fluimucil) em doses pediátricas (300–400 mg) em quadros com secreção.
     - Soluções nasais: Soro fisiológico e soro hipertônico 3% (jatos nasais 3–4 vezes/dia) para congestão.
   - Propostas terapêuticas adicionais discutidas: pelargonium sidoides (Caloba, Imunoflã/Imunoflan), homeopáticos (Corizalha; Ocilococcinum/anas barbariae), própolis verde, zinco, vitaminas D e A (cursos curtos 3–5 dias quando níveis desconhecidos), homotoxicologia (Ingestol) e homeopatia (Erizidoro) para modulação de febre; Broncho-Vaxom (lisado bacteriano). Probióticos (Saccharomyces boulardii e simbióticos) e smectite para diarreia; evitar loperamida.

---

### Chunk 5/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.607

o de vida.
- Mitocôndria:
  - Estratégias já ensinadas no curso (suporte energético, antioxidantes, cofatores).

## Tosse Pós-COVID: Abordagem Específica
- Mecanismo:
  - Modulação do nervo vago frequentemente necessária; interleucina-12 aumentada em tosse irritativa.
- Técnicas e recursos:
  - Microfisioterapia.
  - Dispositivo Miltapod para modular o trajeto do nervo vago.
- Suplementação e fármacos:
  - Quercetina:
    - Doses elevadas: ≥500 mg/dia; preferir forma fitossômica pela biodisponibilidade.
    - Observação: muitas vezes insuficiente isoladamente.
  - Considerar hiperativação mastocitária/intolerância à histamina:
    - Ebastel (ebastina) 10 mg manhã e 10 mg noite por 1 mês; depois reduzir dose conforme evolução.
    - Solicitar exames quando suspeito:
      - Metil-histamina urinária de 24 horas.
      - Atividade de DAO (diaminoxidase) sanguínea.

---

### Chunk 6/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

lis e zinco** por 3 meses foi eficaz na redução da recorrência de otites.
*   **Bronquiolite**
    *   O tratamento padrão é inalação com soro fisiológico. O uso de corticoides e broncodilatadores deve ser evitado na maioria dos casos, pois podem atrapalhar o sistema imunológico.
*   **Refluxo, Cólica e Constipação**
    *   Quadros exacerbados devem levantar a suspeita de Alergia à Proteína do Leite de Vaca (APLV).
    *   O guideline de gastroenterologia indica a dieta de restrição de leite (na mãe ou troca da fórmula) antes de iniciar medicamentos para refluxo. A constipação em menores de 1 ano também é um forte indicativo.
*   **Tratamento para Quadros Agudos (Estratégias Integrativas)**
    *   **Medidas Iniciais:** Lavagem nasal e inalação para mobilizar secreções.
    *   **Fitoterápicos:** **Pelargonium sidoides** (Caloba, Imunoflan) diminui a replicação viral, a duração e a intensidade da doença.

---

### Chunk 7/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.599

cional e Integrativa
*   **Princípio:** Usar a menor dose efetiva de medicação para controle da doença, focando na redução gradual ("step-down").
*   **Intervenções:**
    *   **Remoção de Gatilhos:** Além de alérgenos, inclui produtos químicos (amaciantes), perfumaria e metais pesados (arsênico).
    *   **Dieta e Nutrição:** Dieta anti-inflamatória, livre de alérgenos e contaminantes.
    *   **Atividade Física:** Recomendada, com uso preventivo de SABA se necessário para broncoespasmo induzido por exercício.
    *   **Técnicas Mente-Corpo:** Mindfulness e exercícios respiratórios.
    *   **Controle de Comorbidades:** Manejo de anemia, carências nutricionais, obesidade e efeitos colaterais dos corticoides.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Implementar uma dieta anti-inflamatória, livre de alérgenos, contaminantes e defensivos agrícolas.
- [ ] 2.

---

### Chunk 8/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.597

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

### Chunk 9/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.596

ão controlados, especialmente na asma neutrofílica (comum em obesos), e o ômega 3 pode ajudar a reduzir essa inflamação.
### 6. Fitoterápicos: Quercetina
- **Mecanismo:** Inibe a liberação de citocinas inflamatórias e de histamina pelos mastócitos (ação similar ao cromoglicato), além de regular a atividade da musculatura lisa.
- **Evidências e Segurança:** Estudos mostraram que a quercetina diminui sintomas e aumenta o peak flow. Doses seguras em adultos são de 500mg por até 12 semanas. Faltam estudos de segurança e dose em crianças.
### 7. Fitoterápicos: Cúrcuma na Asma e Rinite
- **Mecanismo:** A cúrcuma é segura e demonstrou diminuir marcadores inflamatórios (IL-4, TNF-alfa) e aumentar os anti-inflamatórios (IL-10).
- **Evidências:** Um estudo brasileiro com crianças mostrou melhora nos sintomas e redução no uso de medicação de resgate. Como 90-95% dos asmáticos têm rinite, tratar a rinite é fundamental.

---

### Chunk 10/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.596

e (6-12 meses até normalização) e dose de estresse em infecções/cirurgias.
    *   **Suplementação/Prevenção:** Imunoestimulação, Vitamina D, Ômega 3, Carotenoides (foco em antioxidação e inflamação).
    *   **Exacerbações:** Corticoides orais (ex: Prednisolona).
*   **Próximos Passos e Exames:**
    *   **Monitoramento Respiratório:** Espirometria e aplicação do ACT (Teste de Controle da Asma) a cada consulta. Avaliação da técnica inalatória.
    *   **Monitoramento Endócrino/Crescimento:** Acompanhamento da estatura a cada 6 meses (crianças em uso de CI). Dosagem de cortisol às 8h (rastreio) e Teste de ACTH se sintomático com cortisol normal.
    *   **Investigação:** Monitorar interações com inibidores de CYP3A4 e comorbidades (refluxo, apneia).
*   **Plano de Tratamento de Acompanhamento:**
    *   **Controle Ambiental:** Redução de mofo, poeira, pelos, produtos químicos.

---

### Chunk 11/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.594

x.: Oxberry 30%, 160 mg 2x/dia; total 320 mg/dia por até 24 semanas).
- [ ] 9. Evitar probióticos em fases de fermentação/gases excessivos; introduzir posteriormente conforme melhora; monitorar sintomas.
- [ ] 10. Estabelecer atuação integrada com nutricionista qualificado para desenho, acompanhamento e ajuste das estratégias nutricionais.
- [ ] 11. Revisar/executar plano de gerenciamento de estresse para elevar tônus parassimpático (sono, respiração, mindfulness, rotinas).
- [ ] 12. Prescrever atividade física com foco em aumento de massa muscular como proteção contra infecções e desfechos pós-inflamatórios.
- [ ] 13. Orientar padrão alimentar evitando ultraprocessados/farináceos; não remover gorduras de forma indiscriminada, limitando gordura trans e priorizando qualidade.
- [ ] 14. Integrar polifenóis e micronutrientes com evidência (quercetina, resveratrol, EGCG, licopeno, curcumina, luteolina, magnésio) conforme caso e referências do material.
- [ ] 15.

---

### Chunk 12/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.592

ção de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.
- **Recomendação Prática:** Usar doses diárias, por longos períodos, buscando níveis acima de 60 ng/ml, com cautela em crianças.
### 5. Ômega 3, PCR e Asma
- **Mecanismo:** O EPA inibe a via do ácido araquidônico (diminuindo leucotrienos) e o DHA diminui a infiltração eosinofílica.
- **Evidências:** Em gestantes, doses altas reduziram a incidência de asma nas crianças. Pacientes com ômega-3 index > 8% necessitam de menos corticoide. A falha de meta-análises pode ser explicada por polimorfismos (ex: FADES) que determinam a resposta à suplementação.
- **Contexto Clínico:** O PCR está aumentado em asmáticos não controlados, especialmente na asma neutrofílica (comum em obesos), e o ômega 3 pode ajudar a reduzir essa inflamação.
### 6.

---

### Chunk 13/30
**Article:** Emagrecimento - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.588

elevância clínica.
- Boswellia padronizada entrega mesma eficácia com menos cápsulas, favorecendo adesão.
- Suplementos lipídicos devem ser tomados com refeições para melhor absorção e conforto gástrico.
### Alavancas clínicas complementares
Protocolos simples e personalizados maximizam resultados em dor, inflamação e emagrecimento.
- Inalação direta supera difusão ambiental para efeitos terapêuticos de óleos essenciais.
- Beta-cariofileno da copaíba ativa CB2 e favorece analgesia e modulação inflamatória.
- Otimizar vitamina D melhora resistência insulínica e marcadores inflamatórios, com doses individualizadas por polimorfismos GC/VDR.

---

### Chunk 14/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.588

o de magnésio inalatório para crianças de 2-6 anos com exacerbação grave. Para maiores de 6 anos e adultos, pode ser usado 2g EV em caso de falha no tratamento inicial para evitar internação.
- **Uso Preventivo:** Um estudo com 330mg de magnésio por 6 meses mostrou melhora na qualidade de vida e controle da doença, mas sem alteração no VEF1 ou nos níveis séricos de magnésio.
### 4. Vitamina D e Asma
- **Mecanismo:** Níveis baixos (< 30 ng/ml) pioram o controle da asma. A Vitamina D melhora a ação do corticoide e modula a resposta imune (diminui citocinas inflamatórias e aumenta a anti-inflamatória IL-10).
- **Evidências:** Apesar da forte plausibilidade, meta-análises falham em demonstrar que a suplementação reduz exacerbações em adultos. Em crianças, há uma redução de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.

---

### Chunk 15/30
**Article:** MFI PÓS RACHEL GAIGER AULA 01 - OXIGÊNIO HIPERBÁRICO (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.583

cofatores mitocondriais (CoQ10, L-carnitina, ácido alfa-lipoico, glutationa, glicina, taurina, tirosina, MCT, vitaminas B1/B2/B3).
- Integração com terapias padrão e medicina funcional: redução de sintomas para permitir abordagem da causa raiz; sinergia com antibióticos/antifúngicos, esteroides em DII e reabilitação motora/cognitiva em casos neurológicos.
### Segurança, contraindicações e manejo de efeitos adversos
- Contraindicações relativas: doenças pulmonares crônicas (enfisema), pneumotórax recente, cirurgia cardíaca/trauma recente, neurite óptica, uso atual de bleomicina/doxorrubicina, otite média/dificuldade de equalização, implante coclear, marcapasso (avaliação individual), hipercapnia, hipertensão não controlada, transtornos convulsivos.
- Absolutas: pneumotórax, broncoespasmo agudo não resolvido, próteses orbitais específicas.
- Drogas incompatíveis: cisplatina, doxorrubicina; cautela com amiodarona e antiangiogênicos.

---

### Chunk 16/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.580

ra fenótipo de sibilância.
**Corticosteroides inalatórios: efetivos, mas com riscos hormonais, de crescimento e ósseos que exigem vigilância e individualização.**
- Supressão do eixo HPA: 10% sintomática e até 40% bioquímica; risco aumenta 6x em crianças e 4x em adultos com alta dose por 3–6 meses.
- Supressão com corticoide oral: cursos >2 semanas consecutivas ou >3 semanas em 6 meses elevam risco.
- Eixos de monitoramento: cortisol às 8h da manhã; se normal, reavaliar em 6 meses; no teste com ACTH, resposta deve subir 18 µg/dL; preocupação com valores de cortisol tão baixos quanto 3 mg/dL.
- Tratamento de supressão: hidrocortisona base por 6–12 meses; atrofia suprarrenal pode persistir até um ano após suspensão de inalatórios.
- ICS e crescimento: perda final de ~1 cm; diferença anual de ~0,2 cm; achados em revisão com quase 3.400 crianças, por 12–52 semanas; contínuo vs cromoglicato: ~1 cm a menos.

---

### Chunk 17/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.579

io inferior/ano.
- Achados relatados:
  - Radiografia com descrição leiga de “catarro no pulmão” (sem laudo formal).
  - Otites predominantemente virais; antibiótico apenas em bilateral grave, dor intensa 2–3 dias sem controle, ou supuração.
- Condutas objetivas em IVR/otites:
  - Lavagem nasal com soro fisiológico (preferir baixa pressão); soro hipertônico 3% 3–4x/dia em congestão.
  - Inalação para fluidificação.
  - N-acetilcisteína 300–400 mg conforme bula.
  - Própolis como adjuvante.
  - Analgésicos: Dipirona; anti-inflamatórios curto prazo para dor em casos selecionados.
- Febre: Evitar antitérmicos indiscriminados; tratar pela clínica (prostração/dor) mais que pelo número; antitérmico não previne convulsão febril.
- Bronquiolite: Inalação com soro fisiológico; evitar corticoide e broncodilatador na maioria sem desconforto respiratório significativo.

---

### Chunk 18/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.579

 mega-índex e suplementar ômega-3 para atingir níveis > 8%, especialmente em pacientes obesos, para reduzir a inflamação.
- [ ] 8. Utilizar sulfato de magnésio como terapia coadjuvante em exacerbações graves e considerar seu papel na prevenção.
- [ ] 9. Incorporar fitoterápicos como cúrcuma, quercetina e Boswellia serrata como coadjuvantes, respeitando as doses seguras.
- [ ] 10. Focar na manutenção de uma microbiota saudável, especialmente em crianças na "janela de oportunidade", promovendo parto normal, amamentação e evitando o uso excessivo de antibióticos.
- [ ] 11. Investigar a presença de inflamação sistêmica (ex: medir PCR-US) em pacientes com asma de difícil controle.
- [ ] 12. Incorporar práticas de mindfulness e exercícios respiratórios para melhorar o controle da asma.

---

### Chunk 19/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.578

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 20/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.577

oro fisiológico; evitar corticoide e broncodilatador na maioria sem desconforto respiratório significativo.
- APLV (alergia à proteína do leite de vaca) como diferencial em refluxo/cólicas/constipação 0–12 meses; considerar dieta de exclusão antes de medicar.
- Exames sugeridos para avaliação imunológica e nutricional:
  - 25-OH vitamina D, vitamina A.
  - Zinco (idealmente eritrocitário).
  - Perfil de ferro (ferritina, ferro sérico, transferrina/TSAT).
  - Hemograma completo; vitamina B12 opcional.
  - Imunoglobulinas (perfil imunológico) devido a infecções de repetição e múltiplos antibióticos.
  - Prick test para aeroalérgenos (ex.: ácaros).
- Observação clínica em fase aguda (“vir ao consultório quando estiver doente”) para confirmação diagnóstica.

---

### Chunk 21/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.575

atórios. Além disso, a discussão aprofunda-se na importância do microbioma pulmonar e do eixo intestino-pulmão, a hipótese da higiene e a janela de oportunidade na primeira infância para modular o sistema imune. A apresentação diferencia os fenótipos e endótipos da asma (TH2 e não-TH2), suas respostas distintas aos tratamentos e como a medicina funcional pode oferecer uma abordagem mais eficaz, especialmente para casos de asma grave, focando na remoção de gatilhos, dieta, controle de comorbidades e modulação imunológica.
## 🔖 Pontos de Conhecimento
### 1. Abordagem Funcional no Tratamento da Asma
*   **Vitamina K2 e Saúde Óssea**
    *   Os corticoides reduzem a massa óssea por meio do bloqueio da osteoprotegerina, o que leva a um aumento dos osteoclastos e, consequentemente, à perda de massa óssea.

---

### Chunk 22/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

; incluir fontes de ômega-3 (peixes); reduzir farináceos/ultraprocessados.
  - Higiene nasal diária e inalação durante quadros respiratórios.
  - Evitar corticoides e broncodilatadores em bronquiolite não complicada; usar apenas com indicação específica.
  - Otimizar hidratação; reduzir mamadeiras noturnas gradualmente; melhorar higiene do sono.
  - Considerar redução de lactose em diarreia persistente (>14 dias); abordagem de FODMAPs em fermentação/desconforto pós-infecção se necessário.
  - Probióticos (Bifidobacterium/Lactobacillus) para reduzir IVR recorrentes, com cautela em intestino muito inflamado; glutamina pode ser considerada em plano nutricional.
  - Educação familiar para manejo de febre/dor, natureza viral das otites, e redução de idas desnecessárias ao pronto-socorro e de prescrições inadequadas.
  - Manter calendário vacinal atualizado; reforçar medidas de controle de exposição em creche e ambiente domiciliar.

---

### Chunk 23/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.572

/ml, com cautela para evitar hipercalcemia.
*   **Ômega-3**
    *   Possui plausibilidade bioquímica, inibindo o NF-κB e a via do ácido araquidônico (reduzindo leucotrienos).
    *   Em gestantes, a suplementação reduziu a incidência de asma na criança.
    *   Ensaios clínicos falham em mostrar diferenças, levando a American Thoracic Society a recomendar contra seu uso.
    *   Polimorfismos no gene FADS podem explicar resultados conflitantes. Pacientes com ômega-índex > 8% necessitam de menos corticoide inalatório, inclusive obesos, contradizendo a recomendação da ATS.
### 3. Fitoterápicos no Tratamento da Asma
*   **Quercetina**
    *   Inibe a liberação de citocinas inflamatórias, a diferenciação de linfócitos TH2 e a liberação de histamina pelos mastócitos.
    *   Em humanos, a forma lipossomada (250-300mg) melhorou sintomas e peak flow. Doses de 500mg são consideradas seguras.
*   **Cúrcuma (Curcuminoides)**
    *   É segura em altas doses.

---

### Chunk 24/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.571

*   A retirada do leite pode diminuir as infecções, não necessariamente por alergia, mas por reduzir um processo inflamatório crônico sistêmico.
    *   Uma quantidade exagerada de proteína (como a caseína) pode causar disbiose e aumento da permeabilidade intestinal, tornando o corpo mais suscetível a infecções.
### 3. Abordagem de Condições Específicas e Tratamentos
*   **Otite Média Aguda**
    *   Mais de 80% são virais. Sinais de complicação bacteriana incluem otite bilateral, dor intensa não controlada ou supuração.
    *   **Tratamento Clínico:** Analgesia (Novalgina preferível ao paracetamol), lavagem nasal, inalação com soro, fluidificantes (N-acetilcisteína), soro hipertônico (3%) e própolis.
    *   Um estudo mostrou que a associação de **própolis e zinco** por 3 meses foi eficaz na redução da recorrência de otites.
*   **Bronquiolite**
    *   O tratamento padrão é inalação com soro fisiológico.

---

### Chunk 25/30
**Article:** Base Metabólica das Doenças Crônicas e Gerenciamento - Oxidação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.571

a, delta). Doses de ataque (estudos) de 800 UI/dia por 2 meses, depois reduzindo para 200-400 UI.
    - **Funções:** Neuroprotetora, previne câncer, catarata, auxilia no uso da vitamina A e é adicionada a suplementos (ex: ômega 3) para evitar oxidação.
### 5. N-acetilcisteína (NAC)
- **Definição:** Forma estável do aminoácido cisteína, precursor da glutationa.
- **Ação:** Efeito antioxidante, reduz citocinas pró-inflamatórias. Atua tanto na via antioxidante não enzimática quanto na enzimática.
- **Usos clínicos:** Expectorante, redutor de muco, e estudos para depressão, transtorno bipolar, esquizofrenia, TDAH e prevenção de diabetes.
- **Formas e dosagem:** Idealmente em comprimido (devido ao gosto ruim). Doses de 600 a 1.800 mg/dia.
### 6. Gestão do Estresse Oxidativo e Suplementação Avançada
- **Avaliação:** Pode ser feita por testes genéticos ou análise clínica (histórico de infarto, LDL oxidada, envelhecimento precoce).

---

### Chunk 26/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.570

ianças mostrou melhora nos sintomas e redução no uso de medicação de resgate. Como 90-95% dos asmáticos têm rinite, tratar a rinite é fundamental. Um estudo em adultos com rinite alérgica mostrou que 500mg de cúrcuma reduziu drasticamente os sintomas e aumentou o fluxo nasal.
### 8. Fitoterápicos: Boswellia Serrata
- **Mecanismo:** Os ácidos bosvélicos inibem a síntese de leucotrienos, um mecanismo relevante para o controle da asma.
- **Evidências:** Um estudo de 1998 (300mg, 3x/dia) mostrou melhora nos sintomas, VEF1 e marcadores inflamatórios. No entanto, a evidência científica geral é limitada, e o uso se baseia principalmente na plausibilidade bioquímica.
### 9. Microbioma, Hipótese da Higiene e Asma
- **Eixo Intestino-Pulmão:** O pulmão não é estéril. Existe um eixo bidirecional onde a microbiota intestinal e pulmonar se influenciam, modulando a imunidade local e sistêmica. A disbiose pulmonar (aumento de proteobactérias) está associada à asma.

---

### Chunk 27/30
**Article:** TDAH - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.570

- **Suplementação Nutricional**:
        - Ácidos graxos ômega-3 (doses elevadas, ex: 2,5g/dia com ênfase em EPA).
        - Vitamina D (ex: 1.000 UI/dia a 50.000 UI/semana, com monitoramento).
        - Magnésio bisglicinato (ex: 200mg a 1g) combinado com Glicina (ex: 1-2g) à noite para melhorar o sono.
        - Considerar suplementação de Zinco, Ferro e Vitaminas do complexo B (B6, folato) se houver deficiência.
    - **Fitoterápicos e Outros**:
        - Curcuminoides para reduzir a inflamação.
        - Açafrão, Ashwagandha, L-teanina para ansiedade e sono.
        - Rhodiola rosea ou Ginseng siberiano para energia e foco.
- **Próximos Passos/Exames**:
    - Avaliar e corrigir deficiências nutricionais (Ômega-3, Magnésio, Zinco, Ferro, Vitamina D, Complexo B).
    - Avaliar marcadores inflamatórios (Proteína C-Reativa, TNF-alfa, IL-6).

---

### Chunk 28/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.570

de indicação em <2 anos sem rinite confirmada.
  - Pelargonium sidoides (Caloba/Imunoflã) conforme idade e bula nos primeiros dias de IVR.
  - Homeopáticos em situações específicas: Corizalha para rinorreia aquosa inicial; Ocilococcinum em quadro sugestivo de influenza (evitar uso preventivo semanal).
  - Suplementação dirigida:
    - Zinco: 10–15 mg/dia por 4–7 meses para profilaxia de IVR; em diarreia aguda, <6m 10 mg/dia; ≥6m 20 mg/dia.
    - Vitaminas D e A: se níveis desconhecidos, curso curto de 3–5 dias durante fase mais intensa da infecção; não suplementar se níveis previamente adequados.
  - Modulação de febre: Dipirona; considerar Ingestol (homotoxicologia) e Erizidoro (homeopatia) conforme bula.
  - Broncho-Vaxom (lisado bacteriano): esquema de 3 meses (10 dias de uso em jejum + 20 dias de pausa), podendo ampliar em casos mais graves.
  - Probióticos e adjuvantes em diarreia: Saccharomyces boulardii; smectite; simbióticos; evitar loperamida.

---

### Chunk 29/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

entes de corticoide oral).
2. **Histórico de Medicamentos:**
    *   Corticoides inalatórios (Budesonida, Fluticasona, Beclometasona - incluindo apresentação em nanopartículas).
    *   Corticoides orais.
    *   Cromoglicato de sódio e Nedocromil.
    *   Antifúngicos e Antivirais (inibidores de CYP3A4, ex: Ritonavir, antifúngicos azólicos) que causam interação medicamentosa.
## [Subjetivo:]
O quadro clínico aborda tanto a sintomatologia respiratória quanto os efeitos adversos do tratamento:
*   **Respiratório:** Episódios recorrentes de tosse (pode ser variante tosse), sibilância, aperto no peito e dispneia. Sintomas podem ser persistentes (leve, moderado, grave) ou intermitentes, com despertar noturno e limitação física.

---

### Chunk 30/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

reavaliar em 24-36 horas antes de prescrever antibióticos.
- [ ] 6. Evitar a prescrição de corticoides e broncodilatadores para casos leves de bronquiolite, focando na inalação com soro fisiológico.
- [ ] 7. Considerar um ciclo de tratamento com lisados bacterianos (Broncho-Vaxom) e/ou a suplementação de zinco para reduzir a recorrência de infecções.
- [ ] 8. Questionar e suspender o uso de montelucaste de sódio prescrito para "melhorar a imunidade" devido ao seu perfil de efeitos colaterais.
- [ ] 9. Educar as famílias sobre a função benéfica da febre, orientando a medicar com base no estado geral da criança e não apenas na temperatura.

---

## SOAP

> Data e Hora: 2025-12-09 04:52:05
> Paciente: [Speaker 1]
> Diagnóstico:
## Histórico do Diagnóstico:
1. Histórico Médico:
   - Criança do sexo feminino, 1 ano e 10 meses.
   - Gestação/Parto: Nasceu de parto normal.
   - Aleitamento: Mamou ao peito até 3 meses (desmame precoce).

---

