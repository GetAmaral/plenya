# ScoreItem: DHEA-S - Mulheres (40-49 anos)

**ID:** `c77cedd3-2800-7329-9692-6d7cd8bb92c6`
**FullName:** DHEA-S - Mulheres (40-49 anos) (Exames - Laboratoriais)
**Unit:** µg/dL
**Gender:** female
**Age Min:** 40 anos
**Age Max:** 49 anos

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.676

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7329-9692-6d7cd8bb92c6`.**

```json
{
  "score_item_id": "c77cedd3-2800-7329-9692-6d7cd8bb92c6",
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

**ScoreItem:** DHEA-S - Mulheres (40-49 anos) (Exames - Laboratoriais)
**Unidade:** µg/dL
**Gênero:** female
**Faixa Etária:** 40-49 anos

**30 chunks de 10 artigos (avg similarity: 0.676)**

### Chunk 1/30
**Article:** Society for Endocrinology Clinical Practice Guideline for the Evaluation of Androgen Excess in Women (2025)
**Journal:** Clinical Endocrinology
**Section:** abstract | **Similarity:** 0.743

Diretriz clínica da Society for Endocrinology para avaliação de excesso androgênico em mulheres. DHEA-S sérico é a medida mais confiável de produção adrenal de andrógenos, particularmente útil na transição perimenopáusica. Aproximadamente 25% da produção androgênica ocorre nas adrenais em mulheres pré-menopáusicas, aumentando para 100% na pós-menopausa. Declínio de DHEA-S inicia aos 30 anos devido à involução da zona reticular adrenal.

---

### Chunk 2/30
**Article:** The Utilization of Dehydroepiandrosterone as a Sexual Hormone Precursor in Premenopausal and Postmenopausal Women: An Overview (2022)
**Journal:** Endocrine Reviews
**Section:** abstract | **Similarity:** 0.723

Revisão sobre DHEA como precursor hormonal sexual em mulheres pré e pós-menopáusicas. DHEA-S é produzido quase exclusivamente nas adrenais e representa 75% dos andrógenos em mulheres pré-menopáusicas e 100% na pós-menopausa. Mulheres perimenopáusicas têm apenas 50% dos níveis de pico de DHEA. DHEA intravaginal (Prasterona) é o único produto aprovado pela FDA para dispareunia moderada a severa na síndrome geniturinária da menopausa.

---

### Chunk 3/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** results | **Similarity:** 0.718

dose dehydroepiandrosterone affects
behavior in hypopituitary androgen-deficient women: a placebocontrolled trial. J Clin
Endocrinol Metab 2002;87:2046
144. van Thiel SW, Romijn JA, Pereira AM, et al. Effects of dehydroepiandrostenedione,
superimposed on growth hormone substitution, on quality of life and insulin-like growth factor
I in patients with secondary adrenal insufficiency: a randomized, placebo-controlled, cross-
over trial. J Clin Endocrinol Metab 2005;90:3295–303
145. Lovas K, Gebre-Medhin G, Trovik T, et al. Replacement of dehydroepiandrosterone in
adrenal failure: no benefit for subjective health status and sexuality in a 9-month randomized
parallel group clinical trial. J Clin Endocrinol Metab 2003;88:1112–8
146. Avis NE, Brockwell S, Randolph Jr JF, et al. Longitudinal changes in sexual functioning
as women transition through menopause: results from the Study of Women’s Health Across
the Nation. Menopause 2009;16:442–52
147.

---

### Chunk 4/30
**Article:** MFI - Reposição Hormonal - AULA 01 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.709

80 anos com testosterona de 500-600, sugerindo que seus níveis de pico na juventude poderiam ter sido muito mais altos (ex: 1.500).
*   **Declínio Hormonal Feminino e DHEA**
    - Mulheres também experimentam um declínio na testosterona, que é produzida pelos ovários e adrenais. A queda pode ser menos linear que nos homens, pois a produção adrenal pode ser preservada.
    - O DHEA (deidroepiandrosterona), um pró-hormônio precursor, diminui nitidamente com a idade, indicando um "desgaste" da glândula adrenal.
    - A prescrição de DHEA (prasterona no Brasil) é feita em comprimidos (forma mais estudada) ou creme (forma mais fisiológica). Doses para mulheres são menores (10-15mg) para evitar conversões indesejadas. Para homens, sua função como estimulador de outros hormônios é menos estabelecida.
### 4.

---

### Chunk 5/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** results | **Similarity:** 0.707

Stanczyk FZ, Bairey Merz CN. DHEA-S levels and
cardiovascular disease mortality in postmenopausal women: Results from the National
Institutes of Health–National Heart, Lung, and Blood Institute (NHLBI)-sponsored Women’s
Ischemia Syndrome Evaluation (WISE). J Clin Endocrinol Metab 2010;95:4985–4992
204. Christiansen JJ, Andersen NH, Sørensen KE, Pedersen EM, Bennett P, Andersen M,
Christiansen JS, Jørgensen JO, Gravholt CH. Dehydroepiandrosterone substitution in
female adrenal failure: no impact on endothelial function and cardiovascular parameters
despite normalization of androgen status. Clin Endocrinol 2007;66:426–433
205. Panjari M, Bell RJ, Jane F, Adams J, Morrow C, Davis SR. The safety of 52 weeks of
oral DHEA therapy for postmenopausal women. Maturitas 2009;63:240–245
206. Cleary MP, Zisk JF. Anti-obesity effect of two different levels of
dehydroepiandrosterone in lean and obese middle-aged female Zucker rats. Int J Obes
1986;10:193–204
207.

---

### Chunk 6/30
**Article:** Ritmo Circadiano Eixo HPA - Parte XI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.704

omens, a dose média de DHEA varia entre 25 e 50 miligramas.
- Notavelmente, uma dose de 50 mg de DHEA, geralmente associada a homens, também é indicada para mulheres com o objetivo de melhorar o humor, o bem-estar e a função sexual.
**A Síndrome da Fadiga Crônica (SFC) afeta predominantemente adultos jovens entre 20 e 40 anos, com uma prevalência duas a três vezes maior em mulheres do que em homens.**
- A SFC é mais comum em adultos na faixa etária de 20 a 40 anos.
- As mulheres são diagnosticadas com SFC de duas a três vezes mais frequentemente do que os homens, indicando uma clara disparidade de gênero na prevalência da condição.
### Achados Adicionais Chave
- A ideia de que as glândulas adrenais são as primeiras a mostrar sinais de cansaço foi introduzida em uma publicação de 1921.

---

### Chunk 7/30
**Article:** MFI - Reposição Hormonal - AULA 01 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.699

reavaliar a cada 6–8 semanas, checar SHBG, estradiol, DHT se pertinente) para padronizar seguimento.
### 8. DHEA (prasterona): papel, prescrição e diferenças entre sexos
- DHEA/DHEA-S: pró-hormônio, não considerado hormônio pleno; declínio com a idade, especialmente adrenal.
- Em mulheres, pode estimular hormônios subsequentes; em homens, não costuma ter função robusta para estimular outros hormônios.
- No Brasil, prescrição como prasterona; estudos principalmente com comprimido; creme seria mais fisiológico, porém menos estudado; ambas formas possíveis.
- Doses mulheres: baixas (10–15 mg) e monitorar porque pode elevar hormônios indesejados; utilidade em infertilidade/autoimunes com dose básica.
- Conceito de “desgaste” glandular ao longo da vida aplicável às adrenais.
> **Sugestões de IA**
> - Organização: Você delimitou bem diferenças por sexo; inclua critérios claros de quando considerar prasterona (indicações vs quando evitar).

---

### Chunk 8/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** discussion | **Similarity:** 0.693

asures [
27]
.
Since DHEA is the main source of androgens in women, its age-related decline leads to a
corresponding decrease in the total androgen pool. Although there is actually no defined
level of androgen below which women can be said to be deficient, the decline of DHEA in
postmenopausal women would 
suggest
 they are “deficient” in both estrogens and androgens
[16]. The declining circulating levels of adrenal androgens with advancing age have been
related to clinical symptoms and disorders (see below).
In the last few years, the concept that adrenal androgen production gradually declines with
advancing age has changed following the analysis of the longitudinal data collected in the
Study of Women’s Health Across the Nation (SWAN) [28].

---

### Chunk 9/30
**Article:** Testosterone in women—the clinical significance (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.692

ndrostenedione and dehydroepiandrosterone (DHEA), 
which are pre-androgens synthesised by the ovaries and 
adrenal glands. The pre-androgens contribute about 
50% of circulating testosterone in premenopausal 
women.5 Concentrations of testosterone ﬁ rst begin to increase in girls at about the age of 6–8 years, when 
maturation of the adrenal zona reticularis results in 
increased production of DHEA and its sulphate, DHEAS,5 heralding the onset of adrenarche.

---

### Chunk 10/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** other | **Similarity:** 0.688

late postmenopause. Gynecolog Endocrinol 2000;14:342–363
139. Panjari M, Davis SR. DHEA therapy for women: effect on sexual function andwellbeing.
Human Reprod Update 2007;13:239–48 
140. Arlt W, Callies F, van Vlijmen JC, et al.. Dehydroepiandrosterone replacement in
women with adrenal insufficiency. N Engl J Med 1999; 341: 1013–1020
141. Mortola JF, Yen SS.The effects of oral dehydroepiandrosterone on endocrine-metabolic
parameters in postmenopausal women. J Clin Endocrinol Metab 1990; 71: 696–704
142. Wolf OT, Neumann O, Hellhammer DH, et al..Effects of a two-week physiological
dehydroepiandrosterone substitution on cognitive performance and well-being in healthy
elderly women and men. J Clin Endocrinol Metab 1997; 82: 2363–2367
143. Johannsson G, Burman P, Wiren L, et al. Low dose dehydroepiandrosterone affects
behavior in hypopituitary androgen-deficient women: a placebocontrolled trial. J Clin
Endocrinol Metab 2002;87:2046
144.

---

### Chunk 11/30
**Article:** Testosterone in women—the clinical significance (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.685

ual function 
is not related to changes in circulating concentrations of 
androgens at menopause or early postmenopause, as 
blood concentrations of androgens do not change during 
the menopausal transition;8,9 instead, associations exist between speciﬁ c androgens and self-reported measures of sexual function in premenopausal and post-
menopausal women.Results of a community-based study14 of 1021 randomly recruited healthy women showed a direct association 
between an endogenous level of DHEAS below the tenth 
percentile and low sexual responsiveness in women aged 
45 years or older. In women aged 18–44 years, 
concentrations of DHEAS below the tenth percentile 
were directly associated with low sexual desire, arousal, 
and responsiveness.14 No associations with andro-stenedione or total and free testosterone were seen.

---

### Chunk 12/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** discussion | **Similarity:** 0.678

Clin Endocrinol Metab 2009;94:3676–81
78. Labrie F, Archer D, Bouchard C, et al. 
Effect of intravaginal dehydroepiandrosterone
(Prasterone) on libido and sexual dysfunction in postmenopausal women. Menopause (New
York, NY) 2009;16:923–31
25

79. 
van den Brandt PA
,
 
Spiegelman D
,
 
Yaun SS
,
 
Adami HO
,
 
Beeson L
,
 
Folsom AR
,
 
Fraser
G
,
 
Goldbohm RA
,
Graham S
,
 
Kushi L
,
 
Marshall JR
,
 
Miller AB
,
 
Rohan T
,
 
Smith-Warner
SA
,
 
Speizer FE
,
 
Willett WC
,
 
Wolk A
,
Hunter DJ
. Pooled analysis of prospective cohort
studies on height, weight, and breast cancer risk. 
Am J Epidemiol.
 
2000 ;152:514-27.
80. Arlt W. The approach to the adult with newly diagnosed adrenal insufficiency. J Clin
Endocrinol Metab 2009;94:1059–1067
81. Panjari M, Davis SR. DHEA for postmenopausal women: A review of the evidence.
Maturitas 2010;66:172–179
82.  Labrie F, Belanger A, Belanger P, et al. Metabolism of DHEA in postmenopausal
women following percutaneous administration.

---

### Chunk 13/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** results | **Similarity:** 0.667

y oﬀer estro-genic beneﬁts, such as improved bone health and pro-tection against bone loss [47]. ese results are further supported by other meta-analyses that reported similar hormonal increases following DHEA supplementation [40, 41].Our review also observed a signiﬁcant increase in tes-tosterone levels with DHEA supplementation admin-istered for  26 weeks. In line with this, the review conducted by Wierman etal. evaluated the eﬀect of DHEA over a long period in older women and found that both estradiol and testosterone levels signiﬁcantly increased [48]. A prospective clinical trial found that 50 mg/day of DHEA led to a higher estradiol concentra-tion and testosterone levels after 6 months of treatment. ese ﬁndings highlight the dynamic, time-dependent hormonal response to DHEA supplementation.Biochemically, DHEA serves as a precursor to both estrogens and androgens.

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.666

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

### Chunk 15/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** results | **Similarity:** 0.664

was further supported by the results of Eggers regression test.Discussionis meta-analysis of randomized controlled trials (RCTs) evaluated the eﬀects of dehydroepiandrosterone (DHEA) supplementation on testosterone and estradiol levels in postmenopausal women. Our analysis of 21 RCTs revealed that DHEA supplementation signiﬁcantly increased both testosterone and estradiol concentrations, although substantial heterogeneity was observed across the included trials.Our ﬁndings indicate that DHEA is particularly eﬀec-tive in elevating testosterone levels when administered at dosages  50 mg/day. Additionally, a marked increase in estradiol levels was noted among participants aged  60 years receiving the same dosage. ese ﬁndings 
Fig. 2 Forest plot of the randomized controlled trials investigating the eﬀect of dehydroepiandrosterone (DHEA) supplementation on serum estradiol concentrations

Page 9 of 13
Heetal. Diabetology & Metabolic Syndrome          (2025) 17:258 
Fig.

---

### Chunk 16/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** discussion | **Similarity:** 0.664

randomized clinical trials. Exp Gerontol. 2020;141: 111110. 41. Zhu Y, Qiu L, Jiang F, Găman M-A, Abudoraehem OS, Okunade KS, et al. The eﬀect of dehydroepiandrosterone (DHEA) supplementation on estradiol levels in women: a dose-response and meta-analysis of randomized clinical trials. Steroids. 2021;173:108889. 42. Jamka K, Adamczuk P, Skowronska A, Bojar I, Raszewski G. Assess-ment of the eﬀect of estradiol on biochemical bone turnover markers among postmenopausal women. Ann Agric Environ Med. 2021;28(2):7896. 43. Clark BJ, Prough RA, Klinge CM. Mechanisms of action of dehydroepi-androsterone. Vitam Horm. 2018;108:2973. 44. Brzozowska M, Lewiński A. Changes of androgens levels in meno-pausal women. Menopause Review/Przegląd Menopauzalny. 2020;19(4):1514. 45. Nair KS, Rizza RA, OBrien P, Dhatariya K, Short KR, Nehra A, et al. DHEA in elderly women and DHEA or testosterone in elderly men. N Engl J Med. 2006;355(16):164759. 46. Villareal DT, Holloszy JO.

---

### Chunk 17/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** results | **Similarity:** 0.663

changes in sexual functioning
as women transition through menopause: results from the Study of Women’s Health Across
the Nation. Menopause 2009;16:442–52
147. Labrie F, Cusan L, Gomez JL, et al.Effect of intravaginal DHEA on serum DHEA and
eleven of its metabolites in postmenopausal women. J Steroid Biochem Mol Biol 2008; 111:
178–194
148. Labrie F, Archer D, Bouchard C, et al.High internal consistency and efficacy of
intravaginal DHEA for vaginal atrophy. Gynecol Endocrinol 2010; 26: 524–532
149. Ibe C, Simon JA. Vulvovaginal atrophy: current and future therapies (CME). J Sex Med
2010; 7: 1042–1050; quiz 51
150. Reiter WJ, Pycha A, Schatzl G, et al. Dehydroepiandrosterone in the treatment of
erectile dysfunction: a prospective double-blind, randomized, placebo-controlled
study.
 
Urology 
1999;53:590-595
151. Steffens DC. A multiplicity of approaches to characterize geriatric depression and its
outcomes. Curr Opin Psychiatry 2009;22:522–526
152. Robichaud M, Debonnel G.

---

### Chunk 18/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** other | **Similarity:** 0.663

verted into its sulfated form, DHEA sulfate (DHEA-S), which circulates at much higher concentrations and can be further metabolized into active sex hormones as needed [5].DHEA levels decline signiﬁcantly with age, decreasing by approximately 70% between the ages of 20 and 60, pri-marily due to reduced adrenal production. DHEA sup-plementation has been explored as an adjunct therapy for postmenopausal women, alongside hormone replace-ment therapy (HRT), and has demonstrated potential beneﬁts in improving metabolic and endocrine function, as well as sexual health [6]. DHEA exerts androgenic, estrogenic, and steroidal eﬀects and has been associated with various biological functions, including antioxidant properties, neuroprotection, and anticancer activity [7].

---

### Chunk 19/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** results | **Similarity:** 0.663

A may
improve the postmenopausal vaginal atrophy-related sexual dysfunction [
147
] without
increasing the circulating levels of estrogen above the postmenopausal range [
81
,
147
-
149
].
Despite initial promising, beneficial effects on sexual function, again, even with intravaginally
administered DHEA, a study failed to show significant benefits [
78
]. 
Apart from women, lower circulating levels of DHEA were also related to erectile dysfunction
in men. A double-blind, placebo-controlled study that enrolled men with
 
erectile dysfunction
treated with 
per os
 with DHEA 50 mg daily has shown some promise for improving
 
sexual
performance
 
in men who had low DHEA blood levels 
[
150
]
. However, high-quality studies
have demonstrated inconsistent results regarding DHEA supplementation for improving
sexual function, libido, and erectile dysfunction. Although research in this area is promising,
additional well-designed studies are required.
3c.

---

### Chunk 20/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** other | **Similarity:** 0.662

arch on Menopause in the 1990s. WHO Technical Report Series
1996;no. 866: 4
30. Lasley B, Crawford S, Laughlin G, et al. Circulating Dehydroepiandrosterone Levels in
Women with Bilateral Salpingo-Oophorectomy during the Menopausal Transition.
Menopause. 2011; 18:494–498
31. Crawford S, Santoro N, Laughlin GA, et al. Circulating Dehydroepiandrosterone Sulfate
Concentrations during the Menopausal Transition. J Clin Endocrinol Metab. 2009; 94:2945–
2951
32. Randolph JF Jr, Sowers M, Gold EB, et al. Reproductive hormones in the early
menopausal transition: relationship to ethnicity, body size, and menopausal status. J Clin
Endocrinol Metab. 2003; 88:1516–1522
22

33. B. L. Lasley, S. Crawford, and D.S. McConnell.Adrenal Androgens and the Menopausal
Transition. Obstet Gynecol Clin North Am. 2011 ; 38: 467–475
34. Parker, LN (1991). "Adrenarche".
 
Endocrinology and metabolism clinics of North
America
 
20
 
: 71–83
35. Liu D, Dillon JS.

---

### Chunk 21/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** discussion | **Similarity:** 0.658

ffect of Dehydroepiandrosterone
on Muscle Strength and Physical Function in Older Adults: A Systematic Review. J Am
Geriatr Soc 2011;59:997–1002
75. Kritz-Silverstein D, von Mühlen D, Gail A. Laughlin GA, Bettencourt R. Effects of
dehydroepiandrosterone supplementation on cognitive function and quality of life: the DHEA
and well-ness (DAWN) trial. J Am Geriat Soc 2008;56:1292–8
76. Panjari M, Bell RJ, Jane F, et al..A randomized trial of oral DHEA treatment for sexual
function, well-being, and menopausal symptoms in postmenopausal women with low libido. J
Sex Med 2009; 6: 2579–2590
77. Alkatib AA, Cosma M, Elamin MB, et al. A systematic review and metaanalysis of
randomized placebo-controlled trials of DHEA treatment effects on quality of life in women
with adrenal insufficiency.
J Clin Endocrinol Metab 2009;94:3676–81
78. Labrie F, Archer D, Bouchard C, et al.

---

### Chunk 22/30
**Article:** Serum concentrations of oxytocin, DHEA and follistatin are associated with osteoporosis or sarcopenia in community-dwelling postmenopausal women (2021)
**Journal:** BMC Geriatrics
**Section:** results | **Similarity:** 0.658

andgaitspeed.Furtherresultsshowedthatpostmeno-pausalwomenwithsarcopeniaweremorelikelytohavehigherDHEAlevels.However,T2andE2werenotre-latedtomusclemass,gripstrengthorgaitspeed.TheseresultssuggestedthatthecirculatingsexhormonesDHEAwasabetterbiomarkerthanT2andE2formusclestrengthandgaitspeed,whichlikelyoccurredbecauseDHEAisaprecursorforsexsteroids.Threerea-sonsmaysupportthishypothesis.First,DHEAdoesnotbindtosexhormone-bindingglobulin(SHBG),whichsuggeststhatDHEAhasfreeaccesstotargetorganscomparedtoSHBG-boundE2andT2[24].Second,DHEAisanandrogenprehormoneproducedinthezonareticularisofthefemaleadrenalglandandovarianthecacells,anditactsasanupstreamprecursorofT2andE2inpostmenopausalwomen[25].DHEAinpostmeno-pausalwomenbecomesthepredominantsexhormoneinsteadofE2[26].CirculatingDHEAprovidessubstratesthatarerequiredforconversionintopotentandrogensandestrogensinperipheraltissues.SkeletalmusclesarecapableofsynthesizingandrogensandestrogensfromDHEA[27].Therefore,TandE2levelsinserumdonotrepresenttheconcen

---

### Chunk 23/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** other | **Similarity:** 0.657

].
Given that a) the adrenal steroids are the most abundant sex steroids in post-menopausal
women and provide a large reservoir of precursors for the intracellular production of
androgens and estrogens in non-reproductive tissues, b) DHEA levels decline with age, c)
pre- and post-menopausal women with lower sexual responsiveness have lower levels of
serum DHEAS [
132
] and d) treatment of postmenopausal women with estrogen and
testosterone have shown some improvement in sexual function, it was proposed that
restoring the circulating levels of DHEA to those found in young women may improve sexual
function and well-being in postmenopausal women [
81
]. Some early 
randomized trials
 that
suffered from methodological issues, such as small number of participants, short treatment
duration and supraphysiological doses, demonstrated positive effects of DHEA replacement
on sexual function and well-being [
70
,
133
-
136
], as well as on relief of menopausal
symptoms [
136
,
138
].

---

### Chunk 24/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** other | **Similarity:** 0.654

ement
in healthy men with age-related decline of DHEA-S: Effects on fat distribution, insulin
sensitivity and lipid metabolism. Aging Male 2003; 6:151–156
27

108. Callies F, Fassnacht M, van Vlijmen JC, Koehler I, Huebler D, Seibel MJ, Arlt W, Allolio
B. Dehydroepiandrosterone replacement in women with adrenal insufficiency: Effects on
body composition, serum leptin, bone turnover, and exercise capacity. J Clin Endocrinol
Metab 2001;86:1968– 1972
109. Flynn MA, Weaver-Osterholtz D, Sharpe-Timms KL, Allen S, Krause G.
Dehydroepiandrosterone replacement in aging humans. J Clin Endocrinol Metab
1999;84:1527–1533
110. Percheron G, Hogrel JY, Denot-Ledunois S, Fayet G, Forette F, Baulieu EE, Fardeau
M, Marini JF. Effect of 1-year oral administration of dehydroepiandrosterone to 60- to 80-
year-old individuals on muscle function and cross-sectional area: A double-blind placebo-
controlled trial. Arch Int Med 2003;163:720–727
111.

---

### Chunk 25/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** results | **Similarity:** 0.654

sely associated with glucose concentration in the medium. J Steroid Biochem Mol Biol. 2000;75(23):15966. 8. De Menezes KJ, Peixoto C, Nardi AE, Carta MG, Machado S, Veras AB. Dehydroepiandrosterone, its sulfate and cognitive functions. Clin Pract Epidemiol Ment Health CP EMH. 2016;12:24. 9. Ma Y, Fan X, Han J, Cheng Y, Zhao J, Fang W, et al. Critical illness and sex hormones: response and impact of the hypothalamicpituitarygonadal axis. Therap Adv Endocrinol Metab. 2025;16:20420188251328190. 10. Kawakita T, Yasui T, Yoshida K, Matsui S, Iwasa T. Correlations of andros-tenediol with reproductive hormones and cortisol according to stages during the menopausal transition in Japanese women. J Steroid Biochem Mol Biol. 2021;214: 106009. 11. Colldén H, Nilsson ME, Norlén A-K, Landin A, Windahl SH, Wu J, et al. Dehydroepiandrosterone supplementation results in varying tissue-speciﬁc levels of dihydrotestosterone in male mice. Endocrinology. 2022;163(12):163. 12.

---

### Chunk 26/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** discussion | **Similarity:** 0.653

o DHEA supplementation.Biochemically, DHEA serves as a precursor to both estrogens and androgens. Once converted to DHEA sulfate in the liver, it undergoes further enzymatic trans-formation in peripheral tissues, where it contributes to the synthesis of active estrogens and androgens, includ-ing estradiol and testosterone [6].It is important to note that elevated levels of estradiol can have adverse eﬀects, such as stimulating the endo-metrium [27]. Consequently, prolonged high estradiol levels may increase the risk of developing endometrial carcinoma [49]. erefore, it is crucial to use low doses of DHEA to raise estradiol levels without overstimulating the endometrium [50]. In our review, subgroup analysis based on DHEA dosage showed that doses of 50 mg/day or higher signiﬁcantly increased estradiol levels.

---

### Chunk 27/30
**Article:** Adrenal Androgens and Aging (2023)
**Journal:** Endotext
**Section:** results | **Similarity:** 0.652

rett-Connor E, von Muhlen D, Laughlin GA, Kripke A. Endogenous levels of
dehydroepiandrosterone sulfate, but not other sex hormones, are associated with depressed
mood in older women: The Rancho Bernardo Study. J Am Geriatr Soc 1999;47:685–691
156. Michael A, Jenaway A, Paykel ES, Herbert J. Altered salivary dehydroepiandrosterone
levels in major depression in adults. Biol Psychiatry 2000;48:989–995
157. Morrison MF, Freeman EW, Lin H, and Sammel MD. Higher DHEA-S
(Dehydroepiandrosterone Sulfate) Levels are Associated with Depressive Symptoms during
the Menopausal Transition: Results from the PENN Ovarian Aging Study. Arch Womens
Ment Health  2011; 14: 375–382
158. Takebayashi M, Kagaya A, Uchitomi Y, Kugaya A, Muraoka M, Yokota N, Horiguchi J,
Yamawaki S. Plasma dehydroepiandrosterone sulfate in unipolar major depression. Short
communication.J Neural Transm  1998; 105:537-42
159. Assies J, Visser I, Nicolson NA, Eggelte TA, Wekking EM, Huyser J, Lieverse R,
Schene AH.

---

### Chunk 28/30
**Article:** Should Dehydroepiandrosterone Be Administered to Women? (2022)
**Journal:** The Journal of Clinical Endocrinology and Metabolism
**Section:** abstract | **Similarity:** 0.649

Revisão crítica sobre administração de DHEA em mulheres. Evidências suportam pequenos benefícios na qualidade de vida e humor em mulheres com insuficiência adrenal primária/secundária ou anorexia, mas não para ansiedade ou função sexual. O efeito ósseo de DHEA é significativamente menor que terapia estrogênica ou medicações aprovadas para osteoporose. DHEA vaginal local mostra promessa para atrofia vulvovaginal.

---

### Chunk 29/30
**Article:** MFI - Reposição Hormonal - AULA 01 (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.648

adrenais.
> **Sugestões de IA**
> - Organização: Você delimitou bem diferenças por sexo; inclua critérios claros de quando considerar prasterona (indicações vs quando evitar).
> - Métodos: Uma tabela de “dose inicial, forma, monitorização, possíveis efeitos” tornaria a aplicação clínica mais direta.
> - Clareza: Bom contraste entre comprimido vs creme; reforce com um protocolo de acompanhamento (ex.: revisar DHEA-S e sintomas após 8–12 semanas).
> - Melhoria: Apresente um caso clínico feminino (p. ex., fadiga, baixa libido, infertilidade) para mostrar decisão e ajuste de DHEA.
## Perguntas dos Alunos
Nenhuma pergunta foi levantada pelos alunos.

---

### Chunk 30/30
**Article:** Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects (2025)
**Journal:** Diabetology & Metabolic Syndrome
**Section:** results | **Similarity:** 0.647

Windahl SH, Wu J, et al. Dehydroepiandrosterone supplementation results in varying tissue-speciﬁc levels of dihydrotestosterone in male mice. Endocrinology. 2022;163(12):163. 12. Jankowski CM, Wolfe P, Schmiege SJ, Nair KS, Khosla S, Jensen M, et al. Sex-speciﬁc eﬀects of dehydroepiandrosterone (DHEA) on bone mineral density and body composition: a pooled analysis of four clinical trials. Clin Endocrinol. 2019;90(2):293300.

Page 12 of 13Heetal. Diabetology & Metabolic Syndrome          (2025) 17:258 
 13. Yuan H, Ou Q, Ma Y, Chen R, Wu H, Li F, et al. Eﬃciency of DHEA on diminished ovarian reserve patients undergoing IVF-ET based on ultrasonography. Int J Clin Exp Med. 2016;9:1433744. 14. Scheﬀers CS, Armstrong S, Cantineau AE, Farquhar C, Jordan V. Dehy-droepiandrosterone for women in the perior postmenopausal phase. Cochrane Datab Syst Rev. 2015(1). 15. McInnes MDF, Moher D, Thombs BD, McGrath TA, Bossuyt PM.

---

