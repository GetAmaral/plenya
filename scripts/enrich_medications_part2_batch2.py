#!/usr/bin/env python3
"""
MEDICAMENTOS PARTE 2 - BATCH 2 (11 items restantes)
"""

import requests
import json
import time
from typing import Dict

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

MEDICATION_CONTENT = {
    "2799ebd7-4c10-4362-9119-d0238a5c4890": {  # Anticonvulsivantes
        "name": "Anticonvulsivantes",
        "clinical_relevance": """O histórico de uso de anticonvulsivantes indica presença de epilepsia, distúrbios convulsivos ou uso off-label para dor neuropática, transtorno bipolar, enxaqueca profilática. Na medicina funcional integrativa, anticonvulsivantes são reconhecidos como terapias essenciais para controle de crises epilépticas e prevenção de dano neurológico, mas a abordagem enfatiza compreensão dos efeitos metabólicos, neurocognitivos e nutricionais significativos.

PRINCIPAIS ANTICONVULSIVANTES:

FENOBARBITAL, FENITOÍNA (anticonvulsivantes clássicos): Induzem enzimas hepáticas CYP450, causando múltiplas interações medicamentosas e depleção de vitaminas D, K, ácido fólico, biotina. Fenitoína tem janela terapêutica estreita (10-20mcg/mL), efeitos adversos: hiperplasia gengival, hirsutismo, ataxia, nistagmo, diplopia, osteoporose, déficit cognitivo. Fenobarbital: sedação, comprometimento cognitivo, dependência.

CARBAMAZEPINA (Tegretol): Bloqueador de canais de sódio. Primeira linha em epilepsia focal. Também usado em neuralgia do trigêmeo, transtorno bipolar. Induz CYP450 (depleta vitamina D, ácido fólico, biotina, CoQ10). Efeitos adversos: hiponatremia (SIADH), leucopenia, aplasia medular (raro mas grave - monitorar hemograma), exantema cutâneo (10% - pode evoluir para Stevens-Johnson), tontura, diplopia, ataxia. Teratogênico (defeitos do tubo neural).

ÁCIDO VALPROICO/VALPROATO (Depakote): Múltiplos mecanismos. Amplo espectro (generalizada e focal). Também usado em bipolar, enxaqueca. Efeitos adversos: ganho de peso significativo (especialmente mulheres), tremor, queda de cabelo, hepatotoxicidade (monitorar enzimas hepáticas), pancreatite, hiperamonemia, SOP (ovários policísticos) em mulheres jovens, trombocitopenia. TERATOGÊNICO grave (>10% defeitos tubo neural, comprometimento cognitivo fetal) - EVITAR em mulheres em idade fértil. Depleta carnitina, selênio, zinco.

LAMOTRIGINA (Lamictal): Bloqueador de canais de sódio. Eficaz em generalizada e focal, depressão bipolar. Geralmente bem tolerada. Principal efeito adverso: rash cutâneo (10%, pode evoluir para Stevens-Johnson ou Lyell - GRAVE). Risco maior se titulação rápida ou uso concomitante de valproato. Requer titulação lenta. Tontura, cefaleia, diplopia.

LEVETIRACETAM (Keppra): Mecanismo único (proteína SV2A). Amplo espectro, bem tolerado, sem interações medicamentosas. Efeitos adversos: alterações comportamentais (irritabilidade, agressividade, depressão) em 10-15%, fadiga, tontura.

TOPIRAMATO (Topamax): Múltiplos mecanismos. Usado em epilepsia, enxaqueca profilática, transtorno bipolar, obesidade. Efeitos adversos: parestesias (formigamento), comprometimento cognitivo ("lentificação mental"), dificuldade de memória e linguagem, perda de peso, litíase renal (acidose metabólica hiperclorêmica), redução de sudorese (risco de hipertermia), glaucoma de ângulo fechado agudo. Teratogênico (fenda palatina).

GABAPENTINA, PREGABALINA: Ligam-se a canais de cálcio. Uso principal: dor neuropática (diabética, pós-herpética), fibromialgia (pregabalina). Pregabalina também ansiolítico. Efeitos adversos: sedação, tontura, ganho de peso, edema periférico. Potencial de abuso/dependência (pregabalina).

OXCARBAZEPINA (Trileptal): Derivado da carbamazepina com menos interações. Similar em eficácia. Hiponatremia mais comum (2,5%).

ZONISAMIDA, LACOSAMIDA, PERAMPANEL: Anticonvulsivantes mais novos.

DEFICIÊNCIAS NUTRICIONAIS INDUZIDAS: Fenitoína, fenobarbital, carbamazepina, ácido valproico DEPLETEM vitamina D (risco de osteoporose, osteomalacia), ácido fólico (anemia megaloblástica, hiperhomocisteinemia, defeitos tubo neural em fetos), vitamina K (risco hemorrágico em recém-nascidos), biotina, L-carnitina (valproato), CoQ10, selênio, zinco.

TERATOGENICIDADE: TODOS os anticonvulsivantes têm risco teratogênico aumentado (2-3x). Risco maior: ácido valproico (10% defeitos), topiramato, fenobarbital, fenitoína. Menor risco: lamotrigina, levetiracetam. Planejamento pré-concepcional essencial, suplementação de ácido fólico 4-5mg/dia (dose alta).

A medicina funcional investiga causas de convulsões: lesões estruturais cerebrais, distúrbios metabólicos (hipoglicemia, hiponatremia, hipocalcemia, hipomagnesemia), deficiências nutricionais (magnésio, B6), intoxicações, abstinência (álcool, benzodiazepínicos), privação de sono, estresse, epilepsia reflexa (fotossensível). Suporte nutricional pode reduzir frequência de crises: dieta cetogênica (eficaz em epilepsia refratária), magnésio, taurina, ômega-3.""",

        "patient_explanation": """Anticonvulsivantes são medicamentos usados principalmente para prevenir crises epilépticas (convulsões), mas também para tratar dor neuropática (formigamento, queimação nos nervos), transtorno bipolar e prevenir enxaqueca.

COMO FUNCIONAM: Estabilizam a atividade elétrica do cérebro, evitando que os neurônios disparem de forma descontrolada (o que causa convulsões).

Os anticonvulsivantes mais comuns incluem:
- Ácido valproico (Depakote): eficaz mas pode causar ganho de peso e problemas sérios em gestantes
- Lamotrigina (Lamictal): bem tolerada, mas requer aumento lento da dose para evitar reações na pele
- Levetiracetam (Keppra): moderno, bem tolerado, mas pode causar irritabilidade
- Carbamazepina (Tegretol): eficaz, mas requer exames de sangue regulares
- Topiramato (Topamax): pode causar perda de peso e "lentidão mental"
- Gabapentina/Pregabalina: usadas principalmente para dor nos nervos

EFEITOS COLATERAIS COMUNS: tontura, sonolência, falta de coordenação, ganho ou perda de peso (conforme medicamento), alterações de humor.

EFEITOS GRAVES (raros): reações alérgicas na pele (especialmente lamotrigina se dose subir rápido), problemas no fígado ou sangue, pensamentos suicidas.

DEFICIÊNCIAS NUTRICIONAIS: Muitos anticonvulsivantes reduzem vitaminas D, ácido fólico e biotina, aumentando risco de ossos fracos e anemia. Suplementação é essencial.

GRAVIDEZ: Todos anticonvulsivantes têm algum risco para o bebê, mas NÃO pare o medicamento sem falar com seu médico - convulsões não controladas são mais perigosas para o bebê. Ácido valproico é o mais arriscado. Se planeja engravidar, converse com antecedência para ajustar tratamento e tomar ácido fólico em dose alta.

É FUNDAMENTAL:
- Tomar exatamente como prescrito, no mesmo horário
- NUNCA pular doses ou parar abruptamente (risco de crise grave)
- Fazer exames de sangue conforme solicitado
- Evitar álcool (reduz eficácia e aumenta sedação)
- Manter sono regular (privação de sono pode causar crises)
- Usar carteira de identificação de epilepsia""",

        "conduct": """1. Monitoramento de níveis séricos (quando aplicável):
   - Fenitoína: 10-20mcg/mL
   - Carbamazepina: 4-12mcg/mL
   - Ácido valproico: 50-100mcg/mL (epilepsia)
   - Coleta em vale (antes da próxima dose)

2. Monitoramento laboratorial:
   - Hemograma: baseline, 3 meses, depois semestral (carbamazepina, valproato)
   - Enzimas hepáticas: baseline, mensal (3 meses), depois trimestral (valproato, carbamazepina, fenitoína)
   - Função renal: anual
   - Sódio: se carbamazepina ou oxcarbazepina
   - Amônia: se valproato + sintomas neurológicos

3. Monitoramento de efeitos adversos específicos:
   LAMOTRIGINA:
   - Orientar sobre rash cutâneo - SUSPENDER se exantema
   - Titulação lenta conforme protocolo

   VALPROATO:
   - Peso, IMC regular
   - Rastreamento SOP (mulheres jovens)
   - Amônia se encefalopatia

   TOPIRAMATO:
   - Cognição, memória
   - Peso
   - Hidratação (prevenir litíase)

4. Suplementação OBRIGATÓRIA para prevenir deficiências:
   TODOS os anticonvulsivantes indutores enzimáticos (fenitoína, carbamazepina, fenobarbital):
   - Vitamina D3 (2.000-4.000UI/dia) + monitorar níveis (alvo >40ng/mL)
   - Ácido fólico (1-5mg/dia)
   - Complexo B
   - Biotina (5-10mg/dia)
   - Cálcio (1000-1200mg/dia) se risco de osteoporose
   - Vitamina K (especialmente em gestantes no 3º trimestre)

   VALPROATO:
   - L-carnitina (500-1000mg/dia)
   - Selênio (100-200mcg/dia)
   - Zinco (15-30mg/dia)
   - CoQ10 (100-200mg/dia)

5. Densidade óssea:
   - DEXA baseline e anual se uso prolongado de indutores enzimáticos
   - Suplementação proativa vitamina D + cálcio

6. Suplementação funcional adjuvante (redução de crises):
   - Magnésio (400-600mg/dia) - ESSENCIAL, cofator GABA
   - Taurina (1-3g/dia) - neuroinibitória
   - Ômega-3 (2-3g/dia) - anti-inflamatório neuronal
   - Vitamina B6 (50-100mg/dia) - cofator síntese GABA
   - Vitamina E (400-800UI/dia)
   - Melatonina (se distúrbios do sono)

7. Dieta cetogênica (considerar em epilepsia refratária):
   - 70-80% gorduras, 10-15% proteínas, 5-10% carboidratos
   - Induz cetose (efeito anticonvulsivante)
   - Eficácia comprovada (redução >50% crises em 30-40%)
   - Requer acompanhamento nutricional especializado

8. Modulação do estilo de vida:
   - Sono regular (7-9h/noite) - privação de sono é gatilho
   - Evitar álcool (baixa limiar convulsivo)
   - Manejo de estresse
   - Evitar luzes intermitentes (se epilepsia fotossensível)
   - Hidratação adequada

9. Planejamento pré-concepcional (mulheres em idade fértil):
   - Ácido fólico 4-5mg/dia (dose ALTA) mínimo 3 meses antes
   - Trocar para anticonvulsivante de menor risco se possível (lamotrigina, levetiracetam)
   - EVITAR ácido valproico
   - Monitorar níveis durante gravidez (podem precisar ajuste)
   - Vitamina K 10mg/dia no último mês de gestação (prevenir hemorragia neonatal)

10. Monitoramento de interações medicamentosas:
    - Indutores enzimáticos (fenito ína, carbamazepina) reduzem eficácia de: anticoncepcionais, varfarina, estatinas, antidepressivos
    - Valproato potencializa lamotrigina (ajustar dose)

11. Descontinuação (apenas sob supervisão médica):
    - Considerar após 2-5 anos livre de crises (conforme tipo epilepsia)
    - Redução gradual lenta (meses)
    - Nunca parar abruptamente (risco de status epilepticus)

12. Reavaliação neurológica periódica (6-12 meses)
    - Frequência de crises
    - Efeitos adversos
    - Adesão ao tratamento
    - Cognição"""
    },

    "c77ebbe3-b387-4427-842d-04ab6960c01d": {  # Antidepressivos
        "name": "Antidepressivos",
        "clinical_relevance": """O histórico de uso de antidepressivos indica presença de transtornos de humor (depressão maior, distimia), transtornos de ansiedade (TAG, pânico, fobia social, TOC, TEPT), dor neuropática, fibromialgia ou outras condições. Na medicina funcional integrativa, antidepressivos são reconhecidos como intervenções válidas para casos moderados a graves, mas a abordagem enfatiza investigação de causas raiz de disfunção neuroquímica (inflamação, deficiências nutricionais, disfunção mitocondrial, disbiose intestinal, hipotireoidismo, desequilíbrio hormonal) e terapias complementares.

PRINCIPAIS CLASSES:

INIBIDORES SELETIVOS DE RECAPTAÇÃO DE SEROTONINA (ISRS - fluoxetina, sertralina, paroxetina, citalopram, escitalopram): Primeira linha em depressão e ansiedade. Bloqueiam recaptação de serotonina na fenda sináptica. Latência terapêutica 2-6 semanas. Efeitos adversos: náusea (primeiras semanas), disfunção sexual (40-60% - redução libido, anorgasmia, disfunção erétil), ganho de peso (especialmente paroxetina), síndrome serotoninérgica (se combinação com outros serotoninérgicos), sangramento aumentado (inibem agregação plaquetária), hiponatremia (SIADH - especialmente idosos), síndrome de descontinuação (tonturas, parestesias, "choques cerebrais" - paroxetina é pior). Fluoxetina tem meia-vida longa (menos descontinuação). Citalopram/escitalopram: prolongam QT em doses altas.

INIBIDORES DE RECAPTAÇÃO DE SEROTONINA E NORADRENALINA (IRSN - venlafaxina, desvenlafaxina, duloxetina): Dupla ação serotonina + noradrenalina. Usados em depressão resistente, dor neuropática, fibromialgia. Efeitos adversos: similares a ISRS + hipertensão (doses altas), sudorese, síndrome de descontinuação intensa (venlafaxina).

ANTIDEPRESSIVOS ATÍPICOS:
- BUPROPIONA (Wellbutrin): Inibe recaptação dopamina e noradrenalina. VANTAGENS: sem disfunção sexual, auxilia cessação tabágica, pode dar energia. DESVANTAGENS: ansiedade/agitação, insônia, reduz limiar convulsivo (contraindicado em epilepsia, bulimia), hipertensão.
- MIRTAZAPINA (Remeron): Antagonista alfa-2 (aumenta serotonina e noradrenalina). VANTAGENS: sedativo (útil em insônia), antiemético, estimula apetite. DESVANTAGENS: ganho de peso significativo, sedação intensa (doses baixas são mais sedativas).
- TRAZODONA (Donaren): Antagonista serotonina. Uso principal: insônia (50-100mg). Priapismo (raro mas urgência urológica).
- VORTIOXETINA (Brintellix): Modulador serotonérgico. Menor disfunção sexual.

ANTIDEPRESSIVOS TRICÍCLICOS (ADT - amitriptilina, nortriptilina, imipramina, clomipramina): Antigos mas ainda usados em dor neuropática, enxaqueca profilática, insônia (amitriptilina), TOC (clomipramina). Bloqueiam recaptação serotonina + noradrenalina + efeitos anticolinérgicos, anti-histamínicos, alfa-bloqueio. Efeitos adversos: boca seca, constipação, retenção urinária, visão turva, sedação, ganho de peso, hipotensão postural, prolongamento QT, arritmias (TÓXICOS em overdose - risco suicídio). Monitorar ECG.

INIBIDORES DA MONOAMINA OXIDASE (IMAO - fenelzina, tranilcipromina, selegilina): Raramente usados (2ª/3ª linha). Eficazes em depressão atípica. CRISE HIPERTENSIVA com tiramina (queijos envelhecidos, vinhos, embutidos, fermentados). Múltiplas interações medicamentosas (síndrome serotoninérgica com ISRS/IRSN).

DEFICIÊNCIAS NUTRICIONAIS ASSOCIADAS A DEPRESSÃO: Vitaminas B6, B12, folato (cofatores síntese neurotransmissores), vitamina D (<20ng/mL associada a depressão), magnésio, zinco, ferro, ômega-3 (EPA especialmente), aminoácidos (triptofano→serotonina, tirosina→dopamina/noradrenalina).

CAUSAS RAIZ DE DEPRESSÃO (medicina funcional): Inflamação crônica (citocinas pró-inflamatórias IL-6, TNF-α, CRP elevada), disfunção mitocondrial, estresse oxidativo, disbiose intestinal (eixo intestino-cérebro), hiperpermeabilidade intestinal, hipotireoidismo subclínico, desequilíbrio hormonal (cortisol alto, estrogênio baixo, testosterona baixa), resistência à insulina, toxicidade por metais pesados, deficiências nutricionais, privação de sono, estresse crônico (eixo HPA desregulado), traumas não processados.

SÍNDROME SEROTONINÉRGICA: Emergência médica. Tríade: alterações mentais (agitação, confusão), hiperatividade autonômica (taquicardia, hipertensão, sudorese, hipertermia), anormalidades neuromusculares (tremor, hiperreflexia, clônus, rigidez). Causada por excesso serotoninérgico (combinação ISRS+IMAO, ISRS+tramadol, etc). Tratamento: suspensão drogas, suporte.""",

        "patient_explanation": """Antidepressivos são medicamentos que corrigem desequilíbrios químicos no cérebro, particularmente de serotonina, noradrenalina e dopamina - neurotransmissores que regulam humor, ansiedade, sono e energia. São usados para tratar depressão, ansiedade, pânico, TOC, dor crônica e outras condições.

TIPOS PRINCIPAIS:
- ISRS (fluoxetina/Prozac, sertralina/Zoloft, escitalopram/Lexapro): mais prescritos, geralmente bem tolerados
- IRSN (venlafaxina/Efexor, duloxetina/Cymbalta): dupla ação, úteis em dor
- Bupropiona (Wellbutrin/Zyban): não causa disfunção sexual, ajuda parar de fumar
- Mirtazapina (Remeron): sedativo, aumenta apetite
- Tricíclicos (amitriptilina): antigos, mais efeitos colaterais, usados em dor

COMO FUNCIONAM: Aumentam disponibilidade de neurotransmissores no cérebro. NÃO fazem efeito imediato - levam 2-6 semanas para melhorar humor.

EFEITOS COLATERAIS COMUNS (primeiras semanas):
- Náusea, boca seca
- Sonolência ou insônia (conforme medicamento)
- Tontura
- Ganho de peso (alguns)
- Disfunção sexual (ISRS - redução de desejo, dificuldade orgasmo)

É FUNDAMENTAL:
- Ter paciência - melhora leva semanas
- NÃO parar abruptamente (causa "síndrome de descontinuação": tonturas, sensação de choques elétricos, irritabilidade)
- Avisar médico se piora súbita ou pensamentos suicidas (paradoxalmente pode ocorrer no início)
- Evitar álcool (potencializa sedação, piora depressão)
- Informar TODOS os médicos (interações medicamentosas)

Procure médico se: piora súbita do humor, pensamentos suicidas, agitação extrema, confusão, febre alta + rigidez muscular (síndrome serotoninérgica).""",

        "conduct": """1. Avaliação pré-tratamento:
   - Triagem de bipolaridade (risco de mania/hipomania)
   - Ideação/risco suicida
   - Uso de substâncias
   - Comorbidades médicas

2. Investigação de causas subjacentes (medicina funcional):
   - Função tireoidiana completa (TSH, T4L, T3L, anti-TPO)
   - Vitamina D (alvo >40ng/mL)
   - Vitamina B12, folato
   - Ferritina (alvo >50ng/mL)
   - Magnésio
   - PCR-us (inflamação)
   - Cortisol (ritmo circadiano se disponível)
   - Glicemia/insulina (resistência insulínica)
   - Hormônios sexuais (se apropriado)

3. Monitoramento inicial:
   - Reavaliação semanal (primeiras 4 semanas) - risco suicídio
   - Escalas: PHQ-9 (depressão), GAD-7 (ansiedade)
   - Efeitos adversos

4. Monitoramento de longo prazo:
   - Peso, IMC
   - Pressão arterial (IRSN)
   - Sódio (se idoso - risco SIADH)
   - ECG (se tricíclicos ou citalopram dose alta)

5. Tratamento de deficiências nutricionais:
   - Vitamina D3 (2.000-5.000UI/dia) se <40ng/mL
   - Vitamina B12 (500-1000mcg/dia) + complexo B
   - Metilfolato (L-5-MTHF 400-1000mcg/dia) - forma ativa
   - Magnésio (300-600mg/dia)
   - Ferro (se deficiente)

6. Suplementação adjuvante baseada em evidências:
   ÔMEGA-3 (EPA>DHA):
   - 1-2g EPA/dia - evidência forte (meta-análises)
   - Efeito anti-inflamatório cerebral
   - Pode potencializar efeito antidepressivos

   SAMe (S-adenosil metionina):
   - 400-1600mg/dia
   - Eficácia similar a antidepressivos leves
   - Doador de metil (síntese neurotransmissores)
   - CUIDADO em bipolaridade (risco mania)

   5-HTP (5-hidroxitriptofano):
   - 50-300mg/dia
   - Precursor serotonina
   - NÃO combinar com ISRS/IRSN (risco síndrome serotoninérgica)

   L-TRIPTOFANO:
   - 1-3g/dia (à noite)
   - Aminoácido precursor serotonina
   - Auxilia sono

   N-ACETILCISTEÍNA (NAC):
   - 1-2g/dia
   - Antioxidante, anti-inflamatório
   - Evidência em depressão bipolar

   ZINCO:
   - 15-30mg/dia
   - Cofator síntese neurotransmissores

   PROBIÓTICOS (eixo intestino-cérebro):
   - Lactobacilos e bifidobactérias
   - Evidência emergente em depressão

7. Terapias não-farmacológicas (essenciais):
   - Psicoterapia (TCC, terapia interpessoal)
   - Exercícios aeróbicos (30-45min, 3-5x/semana) - eficácia similar antidepressivo leve
   - Fototerapia (depressão sazonal)
   - Mindfulness, meditação
   - Higiene do sono
   - Manejo de estresse

8. Dieta anti-inflamatória:
   - Padrão mediterrâneo
   - Redução açúcares, ultraprocessados
   - Aumento vegetais, frutas, peixes
   - Probióticos naturais (fermentados)

9. Manejo de disfunção sexual (ISRS):
   - Redução de dose (se possível)
   - Troca para bupropiona ou mirtazapina
   - "Holiday" de medicamento (skip fins de semana - discutível)
   - Adição de bupropiona (potencializa libido)
   - Sildenafil (homens)

10. Descontinuação (quando apropriado - após 6-24 meses de remissão):
    - Redução GRADUAL lenta (10-25% da dose a cada 2-4 semanas)
    - Meses de desmame
    - Monitoramento de sintomas
    - Suporte psicoterapêutico
    - Manter suplementação (ômega-3, vit D, magnésio)
    - Nunca parar abruptamente

11. Educação do paciente:
    - Latência terapêutica (2-6 semanas)
    - Importância da adesão
    - Efeitos colaterais esperados vs preocupantes
    - Risco de síndrome descontinuação
    - Quando procurar emergência (ideação suicida)"""
    },

    "ae0c5229-76f9-4c8a-8b48-761739ac2268": {  # Antidiabéticos orais
        "name": "Antidiabéticos orais",
        "clinical_relevance": """O histórico de uso de antidiabéticos orais indica presença de diabetes mellitus tipo 2 ou pré-diabetes avançado necessitando intervenção farmacológica. Na medicina funcional integrativa, os antidiabéticos orais são reconhecidos como ferramentas importantes para controle glicêmico e prevenção de complicações, mas a abordagem enfatiza otimização metabólica através de intervenções no estilo de vida, correção de resistência à insulina e tratamento de disfunções subjacentes.

PRINCIPAIS CLASSES:

BIGUANIDAS (METFORMINA): Primeira linha em DM2. Mecanismo: reduz gliconeogênese hepática, melhora sensibilidade à insulina periférica. BENEFÍCIOS: não causa hipoglicemia, não causa ganho de peso (pode ajudar perda), reduz eventos cardiovasculares, baixo custo. EFEITOS ADVERSOS: gastrointestinais (diarreia, náusea, desconforto - 20-30%, melhoram com tempo), deficiência de VITAMINA B12 (10-30% após anos de uso - crucial monitorar), acidose láctica (raro mas grave - contraindicado se insuficiência renal grave ClCr <30, insuficiência hepática, insuficiência cardíaca descompensada, alcoolismo). Suspender 48h antes de contraste iodado. Dose: 500-2550mg/dia, tomar com refeições.

SULFONILURÉIAS (glibenclamida, gliclazida, glimepirida): Estimulam secreção de insulina pelas células beta pancreáticas. BENEFÍCIOS: eficazes, baixo custo. DESVANTAGENS: hipoglicemia (especialmente glibenclamida - evitar em idosos), ganho de peso (2-5kg), exaustão de células beta (perda eficácia ao longo do tempo). Interações: álcool potencializa hipoglicemia.

INIBIDORES DPP-4 (gliptinas - sitagliptina, vildagliptina, linagliptina, saxagliptina): Inibem degradação de incretinas (GLP-1), prolongando sua ação. BENEFÍCIOS: neutro no peso, baixo risco de hipoglicemia, bem tolerados. DESVANTAGENS: eficácia modesta (redução HbA1c 0,5-0,8%), custo elevado, possível risco de pancreatite e insuficiência cardíaca (saxagliptina).

INIBIDORES SGLT2 (gliflozinas - empagliflozina, dapagliflozina, canagliflozina): Inibem reabsorção de glicose nos túbulos renais proximais → glicosúria. BENEFÍCIOS: perda de peso (2-3kg), redução de pressão arterial, PROTEÇÃO CARDIOVASCULAR (redução IAM, IC, morte CV), PROTEÇÃO RENAL (retardam DRC), não causam hipoglicemia. DESVANTAGENS: custo elevado, infecções genitais (candidíase - 10%), infecções urinárias, poliúria, cetoacidose euglicêmica (raro mas grave), risco de amputação (canagliflozina - controverso), depleção de volume. INDICAÇÃO preferencial: DM2 + doença cardiovascular ou IC ou DRC.

TIAZOLIDINEDIONAS (glitazonas - pioglitazona, rosiglitazona): Ativam PPAR-gama, melhoram sensibilidade à insulina. BENEFÍCIOS: efeito duradouro, reduzem esteatose hepática. DESVANTAGENS: ganho de peso significativo (edema, adiposidade), retenção hídrica, risco de insuficiência cardíaca (contraindicado em IC NYHA III-IV), aumento de fraturas (especialmente mulheres), possível risco de câncer de bexiga (pioglitazona - controverso). Uso limitado atualmente.

INIBIDORES ALFA-GLICOSIDASE (acarbose, miglitol): Inibem digestão de carboidratos no intestino delgado, reduzindo glicemia pós-prandial. BENEFÍCIOS: não causam hipoglicemia, neutro no peso. DESVANTAGENS: flatulência intensa, diarreia (50%), eficácia modesta, necessidade de múltiplas doses diárias. Uso limitado.

COMBINAÇÕES: Frequentemente necessário combinar 2-3 agentes. Metformina + iDPP4, metformina + iSGLT2, metformina + sulfoniluréia são combinações comuns.

DEFICIÊNCIAS INDUZIDAS: METFORMINA depleta VITAMINA B12 (mecanismo: interfere absorção ileal dependente de cálcio). Deficiência B12 causa: neuropatia periférica (formigamento, dormência), anemia megaloblástica, déficit cognitivo, fadiga. Monitorar B12 anualmente, suplementar se <400pg/mL.

A medicina funcional enfatiza: Reversão de DM2 através de dieta low-carb/cetogênica (redução HbA1c >1%), jejum intermitente, perda de peso (5-15%), exercícios (resistido + aeróbico), correção de deficiências (vitamina D, magnésio, cromo), redução de inflamação, otimização do sono, manejo de estresse.""",

        "patient_explanation": """Antidiabéticos orais são medicamentos em comprimidos usados para controlar o açúcar no sangue em pessoas com diabetes tipo 2. Cada tipo funciona de forma diferente.

METFORMINA (Glifage, Glucoformina): O mais usado, primeira escolha. Ajuda o fígado a produzir menos açúcar e melhora a ação da insulina. Efeitos colaterais: diarreia/desconforto na barriga (melhora com tempo), deficiência de vitamina B12 (se usar por anos).

GLIBENCLAMIDA, GLICLAZIDA: Fazem o pâncreas produzir mais insulina. Risco de hipoglicemia (açúcar muito baixo) e ganho de peso.

INIBIDORES SGLT2 (Jardiance, Forxiga, Invokana): Fazem o rim eliminar açúcar pela urina. Benefícios extras: perdem peso, protegem coração e rins. Efeitos colaterais: infecções urinárias/genitais, urinar muito.

INIBIDORES DPP-4 (Januvia, Galvus, Trajenta): Aumentam hormônios que controlam açúcar. Bem tolerados, neutros no peso.

É FUNDAMENTAL:
- Tomar exatamente como prescrito
- Não parar sem orientação médica
- Manter dieta saudável e exercícios (medicamentos NÃO substituem hábitos saudáveis)
- Se usar metformina há anos: suplementar vitamina B12
- Reconhecer hipoglicemia: tremor, suor frio, fraqueza, confusão - comer açúcar rapidamente
- Fazer exames regulares (HbA1c, função renal)

ATENÇÃO: Diabetes tipo 2 pode ser REVERSÍVEL em muitos casos através de mudanças intensivas no estilo de vida (perda de peso, dieta low-carb, exercícios).""",

        "conduct": """1. Monitoramento glicêmico:
   - HbA1c: trimestral até meta, depois semestral (meta <7% ou individualizada)
   - Glicemia de jejum
   - Glicemia pós-prandial
   - Automonitoramento (frequência conforme tratamento)

2. Monitoramento de complicações:
   - Função renal (creatinina, TFG) semestral
   - Microalbuminúria anual
   - Perfil lipídico
   - Exame oftalmológico anual (retinopatia)
   - Monofilamento (neuropatia) anual

3. Monitoramento específico por medicamento:

   METFORMINA:
   - Vitamina B12 ANUAL (ou a cada 2 anos)
   - Suplementar se <400pg/mL
   - Função renal (ajustar dose se TFG 30-45, suspender se <30)

   iSGLT2:
   - Função renal baseline e regular
   - Orientar sobre sinais de ITU/candidíase
   - Hidratação adequada
   - Suspender em doenças agudas graves (risco cetoacidose)

   SULFONILURÉIAS:
   - Educação sobre hipoglicemia
   - Evitar em idosos (preferir alternativas)
   - Função renal (ajustar dose)

4. Correção de deficiências nutricionais:
   METFORMINA:
   - Vitamina B12: 500-1000mcg/dia VO ou 1000mcg IM mensal
   - Metilfolato (acompanha deficiência B12)

   GERAL (DM2):
   - Vitamina D (manter >40ng/mL): 2.000-5.000UI/dia
   - Magnésio: 300-600mg/dia (30-50% diabéticos deficientes)
   - Cromo: 200-400mcg/dia
   - Zinco: 15-30mg/dia
   - Ômega-3: 2-3g/dia (anti-inflamatório, proteção CV)

5. Suplementação funcional adjuvante:
   - Berberina: 500mg 3x/dia (eficácia similar metformina)
   - Ácido alfa-lipóico: 300-600mg/dia (neuropatia, sensibilidade insulínica)
   - Canela: 1-6g/dia
   - Inositol: 2-4g/dia
   - Carnitina: 1-2g/dia

6. Intervenção nutricional INTENSIVA:
   - Dieta low-carb (<130g carb/dia) ou cetogênica (<50g/dia)
   - Controle de carga glicêmica
   - Jejum intermitente (16:8 ou 18:6)
   - Educação nutricional contínua
   - Meta: redução HbA1c >1% e eventual descontinuação medicamentos

7. Programa de exercícios:
   - Aeróbico: 150min/semana
   - Resistido: 2-3x/semana (crucial - melhora sensibilidade insulínica)
   - HIIT quando apropriado

8. Perda de peso (se sobrepeso/obesidade):
   - Meta: 5-15% do peso corporal
   - Pode reverter DM2 em muitos casos

9. Educação do paciente:
   - Reconhecimento de hipoglicemia
   - Automonitoramento de glicemia
   - Cuidados com os pés
   - Quando ajustar medicamentos em doenças agudas

10. Escalada terapêutica (se HbA1c não na meta):
    - Adicionar 2º agente (preferir iSGLT2 se DCV/DRC/IC)
    - Considerar GLP-1 agonistas injetáveis
    - Eventualmente insulina

11. Desescalada/descontinuação (se reversão metabólica):
    - Reduzir gradualmente conforme melhora glicêmica
    - Monitoramento frequente
    - Muitos pacientes podem descontinuar com mudanças intensivas estilo vida"""
    }
}

# Continuar com os 7 medicamentos restantes no próximo arquivo...

class MedicationEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()

    def login(self) -> bool:
        try:
            response = self.session.post(
                f"{API_BASE_URL}/auth/login",
                json={"email": EMAIL, "password": PASSWORD}
            )
            response.raise_for_status()
            data = response.json()
            self.token = data["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login realizado!\n")
            return True
        except Exception as e:
            print(f"✗ Erro login: {e}")
            return False

    def update_item(self, item_id: str, content: Dict) -> bool:
        try:
            payload = {
                "clinicalRelevance": content["clinical_relevance"],
                "patientExplanation": content["patient_explanation"],
                "conduct": content["conduct"]
            }
            response = self.session.put(f"{API_BASE_URL}/score-items/{item_id}", json=payload)
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"  ✗ Erro: {e}")
            return False

    def process_all(self):
        results = {"success": [], "failed": [], "total": 0}
        print("="*80)
        print("BATCH 2 - 4 MEDICAMENTOS")
        print("="*80 + "\n")

        for item_id, content in MEDICATION_CONTENT.items():
            results["total"] += 1
            print(f"[{results['total']}/4] {content['name']}")
            if self.update_item(item_id, content):
                print(f"    ✓ Sucesso\n")
                results["success"].append(content['name'])
            else:
                print(f"    ✗ Falhou\n")
                results["failed"].append(content['name'])
            time.sleep(0.5)

        print("="*80)
        print(f"SUCESSO: {len(results['success'])}/{results['total']}")
        print("="*80)
        return results

def main():
    enricher = MedicationEnricher()
    if not enricher.login():
        return
    enricher.process_all()

if __name__ == "__main__":
    main()
