# ScoreItem: Asma

**ID:** `c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b`
**FullName:** Asma (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente))

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 16 artigos
- Avg Similarity: 0.625

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b`.**

```json
{
  "score_item_id": "c77cedd3-2800-7ef5-b90a-7b4b6cc19d5b",
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

**ScoreItem:** Asma (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente))

**30 chunks de 16 artigos (avg similarity: 0.625)**

### Chunk 1/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.663

; em descendência asiática, ≥4.
- Exames laboratoriais para diferenciais:
  - Prolactina (hiperprolactinemia).
  - 17-OHP (HAC não clássica).
  - TSH, T4 (± T3) para disfunção tireoidiana.
  - Testosterona total/livre, DHEA-S (tumores secretores/uso exógeno).
  - USG pélvica; RM/TC se suspeita de tumores.
  - Síndrome de Cushing: cortisol salivar noturno ou teste de supressão com dexametasona 1 mg (se suspeita clínica).
- Achados clínicos gerais:
  - Irregularidade menstrual frequente; ciclos <21 dias, oligomenorreia >35 dias, amenorreia ≥3 meses ou <8 menstruações/ano.
  - Sangramento uterino anormal de causa ovulatória (não estrutural) pode ocorrer.
  - Fenótipo A (três critérios presentes) com maior risco de complicações metabólicas.

---

### Chunk 2/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.657

sidade e inflamação é fundamental.
*   **Modulação dos metabólitos do estrogênio (estronas)**
   - Crucíferas (brócolis, couve-flor, couve) ajudam a tornar estronas menos proliferativas; consumo moderado (≤3–4x/semana) por serem goitrogênicas.
   - Suplementação:
     - **Indol-3-carbinol (I3C):** 200–400 mg/dia; mais fraco e mais barato.
     - **Di-indolilmetano (DIM):** 100–200 mg/dia; estrutura dupla, mais potente.
*   **Acompanhamento avançado com o DUTCH Test**
   - Ideal para acompanhamento assertivo: metabolômica dos hormônios esteroides via DUTCH Test (D-U-T-C-H).
   - Permite visualizar todos os metabólitos hormonais.
   - Exame caro, pouco acessível e complexo; requer estudo prévio do profissional antes de discutir resultados com o paciente.

---

### Chunk 3/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.656

escrever prasterona (DHEA) 25–50 mg por curto período se S-DHEA estiver baixo, monitorando.
    - Confirmada a gravidez, interromper prasterona/DHEA e testosterona.
    - Progesterona (ex.: Utrogestan 100 mg vaginal) do 15º ao 24º dia do ciclo pode aumentar chances de gravidez em mulheres >30 anos ou com pico inadequado; se houver gravidez, o obstetra pode recomendar continuar nos primeiros meses.
### 2. Ômega-3 e Programação Metabólica
*   **Importância do ômega-3 na gestação**
    - A dieta moderna é desproporcionalmente rica em ômega-6 (óleos de soja, canola, milho), tornando o ômega-3 (especialmente DHA) crucial.
    - DHA é fundamental para o desenvolvimento neurológico fetal, com alta concentração no cérebro do feto.
    - Suplementação de ômega-3 na gestação pode reduzir risco de parto prematuro e alergias na infância, além de aumentar modestamente o peso ao nascer.

---

### Chunk 4/30
**Article:** Terapia de Reposição Hormonal Feminina II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.649

tes e fraturas. Recomenda-se o uso de hormônios bioidênticos (17-beta-estradiol e progesterona natural micronizada) e a via de administração transdérmica para otimizar os resultados e minimizar riscos como tromboembolismo.
## 🔖 Pontos de Conhecimento
### 1. Ciclo Menstrual, Climatério e Deficiência Hormonal
*   **Ciclo Menstrual e Produção Hormonal:** O ciclo é dividido nas fases folicular (predominância de estrogênio) e lútea (predominância de progesterona). Todos os hormônios esteroides (estrogênios, progesterona, testosterona) derivam do colesterol. A produção ovariana é explicada pela "teoria das duas células", onde as células da teca produzem androgênios que são convertidos em estrogênios nas células da granulosa.
*   **Queda Hormonal e Menopausa:** A partir dos 25-30 anos, os níveis hormonais declinam. O climatério é o período de transição, com ciclos irregulares e anovulatórios, sendo o momento ideal para iniciar a TRH.

---

### Chunk 5/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.647

riais específicas:
  - Progesterona (dia 21–24 do ciclo; qualquer dia em amenorreia): >3 ng/mL sugere ovulação; >10 ng/mL para manutenção da gravidez; alvo 20–25 ng/mL em contextos selecionados.
- Doses/terapias comentadas:
  - Metformina: prática clínica até 1.500 mg/dia (XR 500 mg; estudos até 1.700 mg/dia); atenção à B12/B6/metilfolato e hiper-homocisteinemia.
  - Inositol: mio-inositol 2 g + ácido fólico 200 mcg, 2x/dia; associar alfa-lactoalbumina 50 mg, 2x/dia em resistentes.
  - Vitamina D: meta sérica 45–60 ng/mL.
  - Melatonina: 3–6 mg/noite; benefícios metabólicos e antioxidantes.
  - NAC: antioxidante, melhora sensibilidade à insulina, proteção hepática, reduz resistência a indutores de ovulação.
  - Ômega-3 (EPA+DHA): 2–4 g/dia.
  - Curcumina: 1.000–2.000 mg/dia com enhancers de biodisponibilidade.
  - Coenzima Q10: suporte à foliculogênese e resposta a indutores.

---

### Chunk 6/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.641

rante a gestação, sem monitorar exames.
    - O palestrante sugere individualização, mantendo níveis sanguíneos acima de 60 ng/mL, idealmente entre 60 e 80 ng/mL, considerando a experiência do professor Coimbra.
    - Importante acompanhar o PTH, conforme aula da Dra. Jéssica.
    - Altas doses (ex.: 50.000 UI em um estudo) mostraram melhorar o hormônio antimülleriano e a fertilidade.
    - Revisões maiores sugerem ponto de corte de 40 a 60 ng/mL para mulheres com dificuldade de engravidar.
*   **Mio-inositol e melatonina**
    - Ambos podem melhorar a qualidade dos oócitos.
    - Mio-inositol melhora a fertilidade principalmente por combater a resistência periférica à insulina; dose de 1 a 2 g/dia.
    - Melatonina é benéfica dada a alta prevalência de sono de má qualidade e maior estresse oxidativo. Ajustar primeiro a rotina de sono; se necessário, suplementar em doses baixas (sublingual), aumentando progressivamente.

---

### Chunk 7/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.637

m: obesidade, sedentarismo, má alimentação, dislipidemia, esteatose hepática, hiperandrogenismo, resistência à insulina, inflamação crônica, disbiose intestinal, estresse oxidativo, disfunção mitocondrial e exposição a desreguladores endócrinos.
2.  **Histórico de Medicação:** Inserir mais aqui
## Subjetivo:
O conteúdo é uma palestra informativa sobre a Síndrome dos Ovários Policísticos (SOP) e não um registro de um paciente específico. O texto não contém queixas subjetivas de um paciente.
## Objetivo:
O conteúdo é uma palestra informativa sobre a Síndrome dos Ovários Policísticos (SOP) e não contém achados de exames de um paciente específico.

---

### Chunk 8/30
**Article:** Effects of micronized progesterone added to non-oral estradiol on lipids and cardiovascular risk factors in early postmenopause (1)
**Journal:** Climacteric
**Section:** discussion | **Similarity:** 0.636

CLIA:Electrochemiluminescenceimmunoassay;FSH:Follicle-stimulatinghormone;HDL-c:High-densitylipoproteincholesterol;hsCRP:High-sensitivityC-reactiveproteintest;
HT:Hormonetherapy;LDL-c:Low-densitylipoproteincholesterol;
MPA:Medroxyprogesterone;NETA:Noretisterone;P:Progesterone;SBP:Systolicbloodpressure;SD:Standarddeviation;SPSS:StatisticalPackagefortheSocialSciences;usCRP:Ultra-sensitiveC-reactiveprotein;WC:Waist
circumference;WHR:Waist-to-hipratio.CompetinginterestsTheauthorsdeclarethattheyhavenocompetinginterests.Authors'contributionsGCandPMScontributedtoacquisitionofdata,analysisandinterpretationof
dataandmanuscriptreview.PMSconceivedanddesignedthestudy.Both
authorscontributedtotheanalysisandinterpretationofdata,draftingmanuscriptandfinalreview,andbothapprovedthefinalversionofthemanuscript.AcknowledgementsThisworkwassupportedbygrantsfromConselhoNacionaldeDesenvolvimentoCientíficoeTecnológico(CNPqINCT573747/2008-3)andFundodeApoioàPesquisadoHospitaldeClínicasdePortoAlegre(FIPE-
HCPA

---

### Chunk 9/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.635

iora a resistência à insulina e a infertilidade. Níveis ideais (acima de 45-60 ng/ml) melhoram a ovulação e a qualidade do óvulo.
*   **Melatonina**
    *   Sua deficiência está ligada à resistência à insulina. A suplementação (3 a 6 mg) melhora o metabolismo energético, a qualidade do óvulo e a função lútea.
*   **N-acetilcisteína (NAC)**
    *   Poderoso antioxidante que aumenta a glutationa, melhora a sensibilidade à insulina, protege o fígado e reduz a resistência a indutores de ovulação.
*   **Ômega 3**
    *   Reduz o estresse oxidativo, a inflamação e os triglicerídeos. Melhora a sensibilidade à insulina e o perfil lipídico. Dose recomendada: 2 a 4 gramas de EPA+DHA.
*   **Curcumina**
    *   Possui potente ação anti-inflamatória, melhora a função mitocondrial e a sinalização da insulina. Dose usual: 1.500 mg a 2g/dia, associada a piperina ou TCM para melhor absorção.

---

### Chunk 10/30
**Article:** The inadequate corpus luteum (2021)
**Journal:** Reproduction and Fertility
**Section:** abstract | **Similarity:** 0.629

Estudo revisional sobre a função do corpo lúteo demonstrando que a produção de progesterona lútea é absolutamente dependente da estimulação do receptor de LH/hCG. O LH é essencial em três momentos críticos: formação do corpo lúteo durante a luteinização folicular, manutenção da produção de progesterona durante a fase lútea, e suporte da gravidez inicial até que o hCG placentário assuma esta função. Identificou que corpos lúteos inadequados em ciclos naturais refletem desenvolvimento oocitário subótimo ao invés de defeito intrínseco do corpo lúteo.

---

### Chunk 11/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.625

melhorar qualidade dos oócitos, especialmente com resistência à insulina e sono inadequado.
- Folato (metilfolato), vitaminas B6 e B12 (metilcobalamina sublingual quando necessário) para reduzir homocisteína em perdas recorrentes e mutação MTHFR; evitar prescrever enoxaparina (“Clexane”) exclusivamente por polimorfismo MTHFR.
- Evitar fitoterápicos e hormônios durante a gestação; exceção: suporte com progesterona micronizada (utrogestan 100 mg via vaginal do 15º ao 24º dia em ciclos regulares, manter no 1º trimestre conforme obstetra).
- DHEA/prasterona 25–50 mg por curto período para fertilidade antes da gestação, com monitorização de testosterona/estradiol; suspender ao engravidar.

---

### Chunk 12/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.623

ação intestinal; óvulos vaginais com óleo de coco e óleos essenciais por 2 semanas, seguidos de probióticos vaginais (lactobacilos) por 2–3 semanas; considerar via oral óleo de orégano 200 mg 2x/dia por 1 mês (não ideal na gestação).
- Ômega-3 (DHA/EPA): suplementação na gestação pode reduzir parto prematuro precoce e aumentar peso ao nascer; atenção especial a veganos/vegetarianos (usar DHA de algas); conversão de ALA (linhaça) é limitada; corrigir desequilíbrio ômega-6/ômega-3.
- Obesidade/SOP: forte associação com infertilidade, complicações gestacionais (DMG, hipertensão, pré-eclâmpsia, parto prematuro, abortamento) e riscos ao bebê; abordar resistência à insulina, inflamação e hábitos de vida; SOP afeta 15–20% das mulheres em idade fértil; 50–70% têm RI; maior risco de desfechos adversos materno-infantis.

---

### Chunk 13/30
**Article:** Follicle Stimulating Hormone (LH:FSH) Ratio in Polycystic Ovary Syndrome (PCOS) - Obese vs. Non-Obese Women (2020)
**Journal:** Med Arch
**Section:** results | **Similarity:** 0.622

rmality of the hypothalamic-pituitary-ovarian or adrenal axis has been imposed in the pathophysiology of polycystic ovarian disease. A dis-turbance in the secretion pattern of the gonadotrophin-releasing hormone (GnRH) results in the relative increase in LH to FSH release (6). Ovarian estrogen is responsible for causing an abnormal feedback mechanism that caused an increase in LH release (7). Usually, in healthy women, the ratio between LH and FSH usually lies between 1 and 2. In polycystic ovary disease women, this ratio becomes reversed, and it might reach as high as 2 or 3 (8).ORIGINAL PAPERdoi: 10.5455/medarh.2020.74.289-293MED ARCH. 2020 AUG; 74(4): 289-293RECEIVED: JUN 10, 2020 | ACCEPTED: AUG 07, 2020 Department of Obstetrics and Gynecology, College of Medicine, Buraidah, Saudi ArabiaCorresponding author: Zaheera Saadia, MD, PhD. Department of Obstetrics and Gynecology, College of Medicine, Buraidah, Saudi Arabia. E-mail: zaheerasaadia@hotmail.com.

---

### Chunk 14/30
**Article:** Identification of the LH surge by measuring intact and total immunoreactivity in urine for prediction of ovulation time (2022)
**Journal:** Hormones (Athens)
**Section:** other | **Similarity:** 0.621

urinary pregnanediol 3-glucuro-
nide to conﬁrm ovulation. Steroids 78(10):10351040
 
18.
 
Ecochard R, Marret H, Rabilloud M, Bradai R, Boehringer H, Girotto S, Barbato M (2000) Sensitivity and speciﬁcity of ultra-
sound indices of ovulation in spontaneous cycles. Eur J Obstet 
Gynecol Reprod Biol 91(1):5964
 
19.
 
Cahill DJ, Wardle PG, Harlow CR, Hull MG (1998) Onset of the preovulatory luteinizing hormone surge: diurnal timing and criti-
cal follicular prerequisites. Fertil Steril 70(1):5659
 
20.
 
Hoﬀ JD, Quigley ME, Yen SS (1983) Hormonal dynamics at mid-cycle: a reevaluation. J Clin Endocrinol Metab 57(4):792796
 
21.
 
Demir A, Hero M, Alfthan H, Passioni A, Tapanainen JS, and Stenman UH (2022) Intact luteinizing hormone (LH), LHbeta, and 
LHbeta core fragment in urine of menstruating women Minerva 
Endocrinol (Torino). https://
 
doi.
 
org/
 
10.
 
23736/
 
S2724-
 
6507.
 
22.
 
03565-5. Online ahead of print.
 
22.

---

### Chunk 15/30
**Article:** Diagnosis and treatment of luteal phase deficiency: a committee opinion (2021)
**Journal:** Fertility and Sterility
**Section:** abstract | **Similarity:** 0.621

Opinião de comitê da ASRM estabelecendo que progesterona é secretada em pulsos sob controle do LH, com produção pulsátil pelo corpo lúteo em resposta aos pulsos de LH. Os pulsos de progesterona são mais pronunciados nas fases média e tardia da fase lútea, podendo flutuar até 8 vezes dentro de 90 minutos. Valores de progesterona oscilam entre 5 e 40 ng/mL em curtos períodos em mulheres ovulatórias normais. Nenhum teste diagnóstico para deficiência de fase lútea provou ser confiável em diferenciar mulheres férteis de inférteis.

---

### Chunk 16/30
**Article:** MFI - Reposição Hormonal - AULA 09 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.619

suplemento mais prescrito, mas com resultados inconsistentes para o aumento de testosterona.
     - Revisões sistemáticas são categóricas em afirmar que não funciona para elevar a testosterona.
     - No entanto, pode melhorar a libido e a disposição, especialmente em mulheres, independentemente dos níveis de testosterona.
     - O palestrante adverte que seu uso pode desregular outros hormônios, como o aumento excessivo de DHT em algumas mulheres, tornando seu efeito imprevisível.
* **Terapias Médicas e Estratégia de Tratamento**
   - **Abordagens não medicamentosas (Prioridade 1):** Dieta, exercício, perda de peso, melhora do sono, redução do estresse e reparo de varicocele (se presente) são fundamentais e devem ser sempre orientadas.
   - **Indicações Médicas (Abordagens Sequenciais):**
     - **HCG (Gonadotrofina Coriônica Humana):** É um análogo do LH que estimula diretamente os testículos a produzirem testosterona.

---

### Chunk 17/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

 mico/insulinêmico (curva glicose/insulina) para resistência à insulina; perfil lipídico e inflamatório; composição corporal; avaliar SOP (clínico, laboratorial e ultrassom); revisar uso de cafeína; investigar candidíase e microbiota vaginal/intestinal conforme sintomas; avaliar elegibilidade para DHEA antes da gestação (não durante).
- Plano de Acompanhamento/Terapêutica:
    - Estilo de vida: dieta equilibrada reduzindo ômega-6 industrial (óleos vegetais refinados) e aumentando ômega-3; ingestão adequada de DHA (peixes ou DHA de algas para veganos); higiene do sono para otimizar melatonina endógena; manejo de peso/obesidade com suporte multidisciplinar.

---

### Chunk 18/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.616

 mica.
    *   **Dosagem e Efeitos Colaterais:** A dose usual é de até 1.500-1.700 mg/dia. Efeitos gastrointestinais são comuns. O uso crônico pode causar deficiência de vitamina B12, exigindo monitoramento. Contraindicada em insuficiência renal.
*   **Progesterona Micronizada**
    *   Essencial para mulheres com SOP que não produzem progesterona adequadamente.
    *   **Função:** Restabelece a regularidade menstrual, protege contra o câncer de endométrio e é crucial para a fertilidade.
    *   **Protocolo:** Usada de forma cíclica (ex: 10-14 dias por mês), na dose de 200 mg. A via oral pode melhorar o humor e o sono devido à produção de alopregnenolona.
    *   Pode causar acne, que pode ser manejada com antiandrogênicos como espironolactona e saw palmetto.
### 4. Suplementação e Terapias Adjuvantes
*   **Inositol (Mio-inositol e D-chiro-inositol)**
    *   Considerado um tratamento de ponta, atua como segundo mensageiro da insulina.

---

### Chunk 19/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 19 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.616

ado em aula posterior com a Dra. Juliana Bica.
### 8. Avaliação e manejo prático antecipados
* Testosterona salivar e cortisol em mulheres
   - Para maior acurácia, preferir avaliação salivar de testosterona e cortisol versus sanguínea, especialmente em mulheres.
* Retirada de contraceptivos orais
   - O processo pode desencadear queda de cabelo e acne; é necessário preparo, aviso e suporte para manejar os efeitos de retirada (“tem que estar bem amparado”).
* Avaliação hormonal ampla
   - Em homens e mulheres, além de testosterona, observar estradiol e diidrotestosterona; entender metabólitos ativos e a “dança” hormonal para decisões terapêuticas equilibradas.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ler a revisão sistemática e meta-análise recente citada pelo instrutor sobre contraceptivos orais e suas implicações (incluindo recomendação nível A para L-metilfolato).
- [ ] 2.

---

### Chunk 20/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.616

1.000–2.000 mg/dia com enhancers de biodisponibilidade.
  - Coenzima Q10: suporte à foliculogênese e resposta a indutores.
## Diagnóstico Principal:
- Avaliação: SOP em discussão, com foco em resistência à insulina, hiperandrogenismo, irregularidade menstrual, riscos e impactos de COCs, e estratégias de manejo metabólico e reprodutivo. Conteúdo educacional; sem confirmação em paciente específico.
- Diagnóstico Suspeito: Nenhum no momento.
## Plano:
- Prescrição:
  - Inserir de acordo com o caso individual.
  - Opções discutidas:
    - Metformina XR até 1.500 mg/dia (avaliar renal, B12/B6/folato).
    - Mio-inositol 2 g + ácido fólico 200 mcg, 2x/dia; considerar alfa-lactoalbumina 50 mg, 2x/dia.
    - Progesterona micronizada oral em esquema cíclico: 200 mg/dia por 10–14 dias do ciclo (ex.: dias 15–24); monitorar humor/sono.
    - Antiandrogênios em sintomas: espironolactona, finasterida; considerar Serenoa repens 400 mg em acne/hirsutismo.

---

### Chunk 21/30
**Article:** Follicle Stimulating Hormone (LH:FSH) Ratio in Polycystic Ovary Syndrome (PCOS) - Obese vs. Non-Obese Women (2020)
**Journal:** Med Arch
**Section:** discussion | **Similarity:** 0.613

ormones (LH &amp; FSH) Associate With Clinical Symptoms Among Women With Polycystic Ovary Syndrome. Glob J Health Sci [Internet]. 2014 Sep 28;7(2). Available from: http://www.ccsenet.org/journal/index.php/gjhs/article/view/3933222.  Yuan C, Liu X, Mao Y, Diao F, Cui Y, Liu J. Polycystic ovary syndrome patients with high BMI tend to have functional dis-orders of androgen excess: a prospective study. J Biomed Res [Internet]. 2016 May 30;30(3):197202. Available from: http://www.jbr-pub.org.cn/ch/reader/view_abstract.aspx?le_no=-jbr160306&ag=123.  Cai J, Zhang Y, Wang Y, Li S, Wang L, Zheng J, et al. High yroid Stimulating Hormone Level Is Associated With Hy-perandrogenism in Euthyroid Polycystic Ovary Syndrome (PCOS) Women, Independent of Age, BMI, and yroid Au-toimmunity: A Cross-Sectional Analysis. Front Endocrinol (Lausanne) [Internet]. 2019 Apr 10;10. Available from: https://www.frontiersin.org/article/10.3389/fendo.2019.00222/full24.  Yu Q, Wang J-B.

---

### Chunk 22/30
**Article:** Progesterone and the Luteal Phase: A Requisite to Reproduction (2015)
**Journal:** Obstetrics and Gynecology Clinics of North America
**Section:** abstract | **Similarity:** 0.612

Estudo abrangente sobre progesterona e fisiologia da fase lútea. Demonstra que o corpo lúteo desenvolve-se com neovascularização imediata resultando em fluxo sanguíneo excepcionalmente elevado. Células lúteas diferenciam-se em dois tipos morfológicos complementares: células pequenas contendo receptores de LH regulando captação de colesterol, e células grandes com maior capacidade esteroidogênica mas sem receptores de LH, conectadas por gap junctions para síntese coordenada de progesterona. A fase lútea normal dura 11-17 dias (média 14,2 dias).

---

### Chunk 23/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

ternativo: Progesterona
- Mulheres com SOP frequentemente têm deficiência de progesterona. A reposição com progesterona micronizada (dose de 200mg, de forma cíclica) pode restabelecer a regularidade menstrual e a fertilidade, proteger o endométrio e melhorar o humor e o sono. Diferente dos COCs, ela permite a ocorrência de uma menstruação fisiológica.
### 7. Suplementação no Tratamento da SOP
- **Inositol (Mio-inositol e D-chiro-inositol):** Melhora a sensibilidade à insulina e a função ovariana. Na SOP, o "paradoxo do ovário" leva a uma conversão inadequada desses componentes, prejudicando a ovulação. A suplementação (dose de 2g de mio-inositol, 2x/dia) ajuda a restaurar o equilíbrio.
    - **Potencialização com Alfa-Lactoalbumina:** Para mulheres resistentes ao inositol, a associação com alfa-lactoalbumina (50mg, 2x/dia) aumenta sua absorção e eficácia.
- **Vitamina D:** A deficiência é comum na SOP e piora a resistência à insulina e a fertilidade.

---

### Chunk 24/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

10–14 dias do ciclo (ex.: dias 15–24); monitorar humor/sono.
    - Antiandrogênios em sintomas: espironolactona, finasterida; considerar Serenoa repens 400 mg em acne/hirsutismo.
    - Melatonina 3–6 mg à noite.
    - NAC, ômega-3 (2–4 g/dia), curcumina (1.000–2.000 mg/dia com piperina/MCT/nanotecnologia), coenzima Q10.
    - COCs (drospirenona, clormadinona, acetato de ciproterona) para controle de acne/hirsutismo/irregularidade menstrual, ponderando riscos metabólicos e trombóticos; estratégia de transição ao suspender COC mantendo por 1–2 meses o tratamento de base de resistência à insulina e estilo de vida para reduzir rebote de acne/queda de cabelo/hirsutismo.
    - Indução de ovulação: letrozol (preferencial) ou citrato de clomifeno; gonadotrofinas conforme caso.
    - Procedimentos: drilling ovariano por videolaparoscopia apenas em refratários após falha clínica e reprodução assistida.

---

### Chunk 25/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.610

e do desequilíbrio hormonal (aumento de LH) e hiperandrogenismo.
## Diagnóstico Primário:
-   **Avaliação:** Síndrome dos Ovários Policísticos (SOP), caracterizada por uma complexa interação bidirecional entre resistência à insulina, hiperandrogenismo, inflamação crônica, estresse oxidativo e disbiose intestinal.
-   **Diagnóstico Suspeito:** Nenhum no momento
## Plano:
-   **Prescrição:** Inserir mais aqui
-   **Próximos Passos/Exame:**
    -   A abordagem terapêutica deve ser multifatorial, atacando cada um dos componentes fisiopatológicos (resistência à insulina, inflamação, disbiose, etc.), e não apenas os sintomas.
-   **Plano de Tratamento de Acompanhamento:**
    -   O tratamento deve focar na raiz do problema, em vez de ser puramente sintomático (como o uso de pílulas anticoncepcionais).
    -   A base do tratamento é a modificação do estilo de vida.
    -   É crucial tratar a disbiose intestinal para melhorar a fertilidade e controlar a SOP.

---

### Chunk 26/30
**Article:** MFI - Reposição Hormonal - AULA 08 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.610

licular e miniaturização irreversível.
* Inflamação, cortisol e DHT
   - Inflamação recorrente aumenta cortisol; sua metabolização exige maior atividade de 5-alfa e 5-beta-redutases.
   - Elevação da 5-alfa-redutase aumenta conversão de testosterona em DHT sistêmica/tecidual; fases inflamatórias elevam DHT e, associadas à inflamação, favorecem queda.
* Nutrientes e metabolismo
   - Inflamação consome complexo B e ferro; mulheres frequentemente têm baixa ferritina.
   - Metilação é essencial para proliferação celular do folículo; folato baixo dificulta crescimento; suplementos capilares incluem complexo B.
* Avaliação sistêmica recomendada
   - Avaliar alimentação, intestino, micronutrientes (zinco), homocisteína, folato, B12, ferro (ferritina, saturação de transferrina), DHT (sangue e saliva), 3-alfa-diol, cortisol.

---

### Chunk 27/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.610

do ao bloqueio da enzima 5-alfa-redutase.
    *   16-hidroxiestrona e 4-hidroxiestrona: Elevadas, indicando desvio do metabolismo hormonal.
    *   Beta-pregnanediol e Alfa-pregnanediol: Níveis baixos, indicando depleção e estresse.
*   **Exames de Sangue Anteriores:** A testosterona sérica estava em um nível normal-baixo.
**Achados Gerais e de Estudos (Apresentação Médica):**
*   **Minoxidil:** Eficaz em cerca de 33% dos casos para queda de cabelo. A eficácia depende do gene SULT1A1; um polimorfismo comum neste gene leva à falta de resposta.
*   **Finasterida e Dutasterida:**
    *   **Mecanismo:** Inibem a enzima 5-alfa-redutase, que converte testosterona em DHT. A dutasterida é mais potente, inibindo os tipos 1 e 2 da enzima.
    *   **Síndrome Pós-Finasterida:** Associação de sintomas sexuais, físicos e psicológicos que se desenvolvem durante ou após o uso e persistem após a descontinuação.

---

### Chunk 28/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 03 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.609

is associa-se a parto prematuro, espermatogênese prejudicada e menor sucesso gestacional.
    - Questão de saúde pública que precisa ser combatida, não romantizada.
*   **Síndrome dos Ovários Policísticos (SOP)**
    - Afeta 15–20% das mulheres em idade fértil; maior causa de infertilidade anovulatória (90–95%).
    - 60–70% têm sobrepeso/obesidade; 50–70% têm resistência à insulina, independentemente do peso.
    - Crucial tratar a causa raiz (resistência à insulina, inflamação), não apenas os sintomas.
*   **Riscos associados à SOP na gestação**
    - **Para a mãe:** Maior risco de diabetes gestacional (3x), hipertensão gestacional (3–4x), pré-eclâmpsia (2–4x), parto prematuro e abortamento.
    - **Para o filho:** Risco aumentado de hipoglicemia, morte perinatal, prematuridade, anomalias congênitas (cardiovasculares, urogenitais), desordens metabólicas, asma e TDAH. Meninas têm maior risco de desenvolver SOP na vida adulta.
### 4.

---

### Chunk 29/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.607

caso.
    - Procedimentos: drilling ovariano por videolaparoscopia apenas em refratários após falha clínica e reprodução assistida.
- Próximos Passos/Exames:
  - Aplicar critérios de Roterdã (confirmar 2 de 3).
  - Excluir diferenciais: prolactina, 17-OHP, TSH/T4(±T3), perfil androgênico (testosterona total/livre, DHEA-S).
  - USG pélvica; RM/TC se suspeita de tumores secretores.
  - Em suspeita de Cushing: cortisol salivar noturno ou teste de supressão com dexametasona 1 mg.
  - Avaliação para uso de COCs: critérios de elegibilidade médica (OMS), triagem de trombofilia quando indicada, PA e doença vascular, investigação de enxaqueca com aura.
  - Avaliação metabólica: glicemia, insulina, HOMA-IR, perfil lipídico, função renal (creatinina) antes de metformina.
  - Dosar progesterona (dia 21–24; qualquer dia em amenorreia).
  - Monitorar vitamina D sérica; ajustar para 45–60 ng/mL.

---

### Chunk 30/30
**Article:** Terapia de Reposição Hormonal Feminina I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

to discreto nos níveis de HDL e LDL.
**Progesterona:**
*   A progesterona transdérmica não tem comprovação científica de proteção endometrial ou mamária.
*   A progesterona oral é mandatória na TRH, mesmo em mulheres sem útero, devido aos seus benefícios na massa óssea e no sistema nervoso central (melhora do sono através da conversão em alopregnanolona).
**Metabolização de Estrogênios e Estroboloma:**
*   A metabolização hepática ocorre em fases (Fase 1, 2 e 3).
*   **Fase 1 (Detoxicação):** Depende do citocromo P450 e de cofatores como vitaminas do complexo B, ácido fólico e ferro. Pacientes com deficiência de ferro podem ter a metabolização prejudicada.
*   **Fase 2 (Conjugação):** Reações como a glucuronidação marcam os estrogênios para eliminação. Depende de nutrientes como magnésio, metionina, cisteína e vitaminas (B5, B12, C).
*   **Fase 3 (Eliminação/Intestinal):** A microbiota intestinal (estroboloma) desempenha um papel crucial.

---

