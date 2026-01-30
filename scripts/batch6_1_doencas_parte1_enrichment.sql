-- ============================================
-- BATCH 6.1: HISTÓRICO DE DOENÇAS - PARTE 1
-- Total: 40 items
-- Autor: Claude Sonnet 4.5
-- Data: 2026-01-28
-- ============================================

-- Item 1: LES
UPDATE score_items
SET
  clinical_relevance = 'Lúpus Eritematoso Sistêmico (LES) é doença autoimune crônica multissistêmica caracterizada por produção de autoanticorpos (anti-DNA, anti-Sm, antifosfolípides) e deposição de imunocomplexos. Prevalência: 50-100/100.000, 90% mulheres, pico 15-45 anos. Fisiopatologia: perda tolerância imunológica, ativação células B/T, produção citocinas pró-inflamatórias (IL-6, TNF-α, interferon tipo I). Manifestações: cutâneas (rash malar, fotossensibilidade), articulares (90%), renais (nefrite lúpica 50%), hematológicas (anemia, leucopenia), neuropsiquiátricas (30%), cardiopulmonares. Mortalidade: doença cardiovascular, infecções, nefrite. Na MFI: investigar gatilhos (radiação UV, disbiose, estresse oxidativo, metais pesados, deficiência vitamina D, permeabilidade intestinal aumentada). Eixos terapêuticos: modulação microbiota (probióticos lactobacillus), anti-inflamatórios naturais (curcumina 1g, ômega-3 EPA/DHA 2-4g), antioxidantes (NAC, glutationa, resveratrol), vitamina D3 (50-80ng/ml), manejo estresse (cortisol), dieta anti-inflamatória (eliminação glúten, laticínios, solanáceas). Suplementação: DHEA 50-200mg (mulheres), vitamina D3 5.000-10.000UI, ômega-3 2-4g, probióticos, NAC 1.200mg, curcumina lipossomal 1g.',

  patient_explanation = 'Lúpus (LES) é uma doença onde seu sistema imunológico fica confuso e começa a atacar tecidos saudáveis do próprio corpo - pele, articulações, rins, coração, cérebro. É como se suas defesas travassem uma guerra contra você mesmo. Afeta principalmente mulheres jovens. Os sintomas variam: manchas no rosto (formato de borboleta), dor nas juntas, cansaço extremo, queda de cabelo, sensibilidade ao sol. Pode afetar rins (nefrite lúpica) e aumentar risco de doenças cardíacas. Na abordagem funcional, investigamos gatilhos: intestino permeável, disbiose, deficiências (vitamina D muito baixa é comum), toxinas ambientais, estresse crônico. O tratamento combina medicações tradicionais com estratégias naturais: dieta anti-inflamatória rigorosa (sem glúten, laticínios, açúcar), suplementação (vitamina D alta dose, ômega-3, antioxidantes), probióticos, manejo estresse, proteção solar. Muitos pacientes conseguem longos períodos de remissão com abordagem integrativa.',

  conduct = 'Investigação laboratorial: FAN (título, padrão), anti-DNA dupla hélice, anti-Sm, anti-Ro/SSA, anti-La/SSB, anticorpos antifosfolípides (anticoagulante lúpico, anticardiolipina, anti-beta2-glicoproteína). Atividade: complemento C3/C4 baixos, VHS, PCR, proteinúria 24h, sedimento urinário. Avaliar: hemograma (citopenias), função renal/hepática, perfil lipídico, homocisteína, vitamina D (alvo >50ng/ml), DHEA-S, cortisol salivar 4 pontos, micronutrientes (zinco, selênio, magnésio), LDH, ferritina. Investigar disbiose (calprotectina, zonulina), metais pesados (mercúrio, chumbo), estresse oxidativo (8-OHdG). Protocolo MFI: dieta AIP (autoimmune protocol) 90 dias, eliminar glúten/laticínios/solanáceas/açúcar. Suplementação: vitamina D3 10.000UI, ômega-3 EPA/DHA 3g, curcumina lipossomal 1g, DHEA 50-100mg (mulheres), probióticos multi-cepas, NAC 1.200mg, glutationa lipossomal 500mg, resveratrol 500mg. Monitorar atividade trimestral (C3/C4, anti-DNA, proteinúria). Hidroquinoproteção solar rigorosa.',

  updated_at = NOW()
WHERE id = '3f26c7c9-88e6-4fcc-9a3c-886a6fb709a6';

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '3f26c7c9-88e6-4fcc-9a3c-886a6fb709a6'
FROM articles a
WHERE a.title ILIKE '%lúpus%' OR a.title ILIKE '%autoimun%' OR a.title ILIKE '%reumatolog%'
LIMIT 5
ON CONFLICT DO NOTHING;

-- Item 2: Esclerodermia
UPDATE score_items
SET
  clinical_relevance = 'Esclerodermia (Esclerose Sistêmica) é doença autoimune crônica caracterizada por fibrose cutânea/visceral, vasculopatia e autoimunidade. Prevalência: 15-30/100.000, 3-4:1 mulheres. Classificação: limitada (CREST: calcinose, Raynaud, dismotilidade esofágica, esclerodactilia, telangiectasias) e difusa (fibrose extensa, acometimento visceral precoce). Fisiopatologia: lesão endotelial → ativação fibroblastos → excesso colágeno tipo I/III, citocinas fibrogênicas (TGF-β, CTGF), radicais livres. Manifestações: fenômeno Raynaud (95%), espessamento cutâneo, úlceras digitais, refluxo gastroesofágico, hipertensão pulmonar (30%), crise renal esclerodérmica, fibrose pulmonar intersticial. Autoanticorpos: anti-Scl70 (difusa), anticentrômero (limitada), anti-RNA polimerase III. Mortalidade: hipertensão pulmonar, fibrose pulmonar, crise renal. Na MFI: investigar estresse oxidativo (ROSs), disfunção mitocondrial, deficiências (vitamina D, CoQ10, L-carnitina), exposição sílica/solventes, permeabilidade intestinal. Modulação: antioxidantes (NAC, vitamina E, selênio), anti-fibróticos (curcumina, resveratrol, EGCG chá verde), vasodilatadores naturais (L-arginina, L-citrulina, magnésio), ômega-3, probióticos.',

  patient_explanation = 'Esclerodermia significa "pele dura" - doença onde tecido conjuntivo produz colágeno em excesso, causando endurecimento e espessamento da pele e órgãos internos. Existem dois tipos: limitada (mais lenta, afeta principalmente mãos/face) e difusa (progressão mais rápida, atinge órgãos). Quase todos têm fenômeno de Raynaud: dedos ficam brancos/roxos no frio devido espasmo vascular. Outros sintomas: pele brilhante e esticada, dificuldade engolir, refluxo, falta de ar, fadiga. Pode afetar pulmões (fibrose), coração e rins. É condição séria mas tratável. Na medicina funcional investigamos: por que sistema imunológico desregulou? Há toxinas ambientais (sílica, solventes)? Deficiências nutricionais? Estresse oxidativo? Disbiose? O tratamento combina medicações com abordagem natural: antioxidantes potentes (NAC, vitamina E, selênio), anti-inflamatórios (ômega-3, curcumina), suporte vascular (L-arginina, magnésio), fisioterapia para pele, proteção contra frio.',

  conduct = 'Autoanticorpos: FAN padrão, anti-Scl70/topoisomerase I (difusa, pior prognóstico), anticentrômero (limitada), anti-RNA polimerase III (crise renal). Extensão: capilaroscopia periungueal (padrão SD), TC tórax alta resolução (fibrose pulmonar), ecocardiograma + cateterismo (hipertensão pulmonar), manometria esofágica, função pulmonar (CVF, DLCO). Marcadores: aldolase, CK (miopatia), NT-proBNP (hipertensão pulmonar). Avaliar MFI: vitamina D (alvo >60ng/ml), CoQ10, L-carnitina, selênio, zinco, vitamina E, estresse oxidativo (malondialdeído, 8-isoprostano), homocisteína, perfil tireoidiano, função renal. Protocolo: dieta anti-inflamatória, eliminar glúten/laticínios. Suplementação: NAC 1.800mg (antifibrótico), vitamina E 800UI, selênio 200mcg, ômega-3 3g, curcumina lipossomal 1g, resveratrol 500mg, EGCG 400mg, L-arginina 3-6g (Raynaud), L-citrulina 3g, magnésio 400mg, probióticos, CoQ10 200mg, vitamina D3 10.000UI. Fisioterapia: mobilização cutânea, exercícios respiratórios. Monitorar CVF/DLCO semestral, ecocardiograma anual.',

  updated_at = NOW()
WHERE id = '697adbc6-3a8b-4039-889b-d9b18025392e';

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '697adbc6-3a8b-4039-889b-d9b18025392e'
FROM articles a
WHERE a.title ILIKE '%esclerodermia%' OR a.title ILIKE '%autoimun%' OR a.title ILIKE '%reumatolog%' OR a.title ILIKE '%fibrose%'
LIMIT 5
ON CONFLICT DO NOTHING;

-- Item 3: Doença de Crohn
UPDATE score_items
SET
  clinical_relevance = 'Doença de Crohn (DC) é Doença Inflamatória Intestinal (DII) crônica transmural que pode acometer qualquer segmento do TGI (boca ao ânus), mais comum íleo terminal e cólon. Prevalência BR: 15-20/100.000, pico 15-35 anos. Fisiopatologia: desregulação imune (Th1/Th17), perda tolerância à microbiota comensal, barreira intestinal comprometida, disbiose (redução Faecalibacterium prausnitzii, aumento E.coli aderente-invasiva). Genética: NOD2/CARD15 (30% caucasianos), ATG16L1, IRGM (autofagia). Manifestações: diarreia crônica sanguinolenta/não, dor abdominal, perda peso, febre. Complicações: estenoses (fibrose), fístulas (entero-entéricas, perianais), abscessos, má absorção (B12, ferro, zinco, vitamina D). Extra-intestinais: artrite (20%), eritema nodoso, pioderma gangrenoso, uveíte, colangite esclerosante. Risco câncer colorretal aumentado 2-3x após 8-10 anos. Na MFI: investigar causas raiz (disbiose, SIBO, fungos, parasitas), permeabilidade intestinal (zonulina), deficiências nutricionais severas, estresse oxidativo, disfunção mitocondrial, sensibilidades alimentares (glúten, laticínios, FODMAPs). Modulação: dieta específica (SCD, low-FODMAP, AIP), probióticos estratégicos, prebióticos, L-glutamina, butirato, curcumina, ômega-3 EPA.',

  patient_explanation = 'Doença de Crohn é inflamação crônica que pode atingir qualquer parte do sistema digestivo, mas principalmente a parte final do intestino delgado (íleo) e início do intestino grosso. A inflamação atravessa todas camadas da parede intestinal (não é superficial). Sintomas principais: diarreia frequente (às vezes com sangue), dor abdominal tipo cólica, perda de peso, cansaço, febre baixa. Com o tempo, inflamação repetida causa cicatrizes e estreitamentos (estenoses) que podem obstruir o intestino, ou formar fístulas (túneis anormais entre órgãos). Também causa problemas fora do intestino: dor nas juntas, lesões de pele, inflamação nos olhos. Na abordagem funcional investigamos o que desregulou o sistema: desequilíbrio das bactérias intestinais (disbiose), intestino permeável, deficiências nutricionais (B12, ferro, zinco, vitamina D ficam muito baixos), sensibilidades alimentares. Tratamento combina medicações com dieta personalizada (eliminar gatilhos), probióticos específicos, suplementação intensiva, cicatrização intestinal (L-glutamina, butirato, zinco).',

  conduct = 'Diagnóstico: colonoscopia com biópsias (inflamação transmural, granulomas não-caseosos), cápsula endoscópica ou enteroressonância (íleo). Atividade: calprotectina fecal (>250 ativa), PCR, VHS, albumina (desnutrição). Avaliar extensão: colonoscopia completa, EDA (acometimento alto). Investigar MFI: microbioma (PCR fezes ou sequenciamento 16S), SIBO (teste hidrogênio/metano), fungos (Candida), parasitas (Blastocystis, Giardia, Entamoeba), permeabilidade (zonulina), micronutrientes (B12, folato, ferro, ferritina, zinco, selênio, magnésio, vitamina D, vitaminas lipossolúveis A/E/K), perfil lipídico (má absorção), homocisteína, marcadores ósseos (osteoporose). Protocolo: dieta SCD (Specific Carbohydrate Diet) ou AIP 90 dias, eliminar glúten/laticínios/açúcar/grãos processados. Suplementação: L-glutamina 15-30g/dia (cicatrização mucosa), butirato 1.200mg, curcumina 3g (remissão), ômega-3 EPA 2-4g, probióticos VSL#3 ou multi-cepas, zinco carnosina 75mg, vitamina D3 10.000UI, B12 injetável 1.000mcg semanal, ferro IV se anemia severa. Monitorar calprotectina trimestral (alvo <50), colonoscopia 1-2 anos.',

  updated_at = NOW()
WHERE id = '76985c39-f90a-45f2-8485-f9d26e4e1369';

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '76985c39-f90a-45f2-8485-f9d26e4e1369'
FROM articles a
WHERE a.title ILIKE '%intestin%' OR a.title ILIKE '%Crohn%' OR a.title ILIKE '%inflamatória%' OR a.title ILIKE '%gastrointestinal%'
LIMIT 5
ON CONFLICT DO NOTHING;

-- Item 4: RCU (Retocolite Ulcerativa)
UPDATE score_items
SET
  clinical_relevance = 'Retocolite Ulcerativa Crônica (RCU) é DII caracterizada por inflamação superficial (mucosa/submucosa) contínua do cólon, sempre acometendo reto, progredindo proximalmente. Prevalência BR: 10-15/100.000. Fisiopatologia: resposta Th2 aberrante, defeito barreira mucosa, depleção camada mucina, disbiose (redução Akkermansia muciniphila, Faecalibacterium, aumento Proteobacteria). Classificação: proctite (15cm distais), colite esquerda (até flexura esplênica), pancolite (todo cólon). Manifestações: diarreia sanguinolenta, tenesmo, urgência fecal, dor abdominal baixa, muco/pus nas fezes. Complicações: megacólon tóxico (3-5%, emergência), hemorragia maciça, perfuração, displasia/câncer colorretal (risco acumulado 2% 10anos, 8% 20anos, 18% 30anos). Extra-intestinais: colangite esclerosante primária (5%, associação forte), artrite, sacroileíte, uveíte, pioderma, eritema nodoso. Na MFI: avaliar disbiose específica (baixo butirato, Akkermansia), permeabilidade intestinal, estresse oxidativo mucosa, deficiências (ferro, folato, zinco, selênio, vitamina D), sensibilidades (sulfitos aumentam sintomas). Modulação: enemas butirato, probióticos (E.coli Nissle 1917, VSL#3), curcumina, Boswellia, ômega-3, dieta low-sulfur se sulfito-sensível.',

  patient_explanation = 'Retocolite (RCU) é inflamação crônica que começa sempre no reto (parte final intestino grosso) e pode subir pelo cólon. Diferente da Crohn, a inflamação é superficial (só camadas internas) e contínua (sem áreas poupadas). Principal sintoma: diarreia com sangue vivo, urgência intensa (precisa correr ao banheiro), evacuar 10-20x/dia nas crises, dor abdominal baixa, muco/pus nas fezes. Pode ter sangramento importante. Complicações: nas crises graves, cólon pode dilatar perigosamente (megacólon tóxico - emergência). Após muitos anos, risco aumentado de câncer colorretal, por isso colonoscopias regulares são essenciais. Também afeta fígado (colangite esclerosante), juntas e pele. Na medicina funcional investigamos: quais bactérias protetoras estão faltando? (Akkermansia e produtoras de butirato geralmente baixas). Há permeabilidade intestinal? Deficiências nutricionais? Sensibilidade a sulfitos? Tratamento: medicações + probióticos específicos (E.coli Nissle, VSL#3), enemas com butirato (cicatrizam mucosa), curcumina, ômega-3, dieta anti-inflamatória personalizada.',

  conduct = 'Diagnóstico: colonoscopia (inflamação contínua superficial mucosa, pseudopólipos, friabilidade, ulcerações), biópsias (distorção criptas, abscessos criptas, depleção células caliciformes). Atividade: calprotectina fecal (>250 ativa), PCR, VHS, albumina, hemograma (anemia). Rastreio CEP: fosfatase alcalina, GGT, elastase fecal. Extensão: classificar Montreal (E1 proctite, E2 esquerda, E3 pancolite). Investigar MFI: microbioma fezes (butirato baixo, Akkermansia, Faecalibacterium), permeabilidade (zonulina), estresse oxidativo (malondialdeído), micronutrientes (ferro, ferritina, folato, B12, zinco, selênio, vitamina D), perfil tireoidiano (associação tireoidite). Protocolo: dieta low-FODMAP ou SCD, reduzir sulfitos se sensível (vinho, conservantes, crucíferas). Suplementação oral: curcumina 3-4g (eficácia equivalente mesalazina leve-moderada), Boswellia serrata 900mg, ômega-3 EPA 2-4g, probióticos VSL#3 900bilhões UFC ou E.coli Nissle, zinco carnosina 75mg, vitamina D3 10.000UI. Enemas: butirato 2g (cicatrização mucosa distal), curcumina, Aloe vera. Monitorar calprotectina mensal em atividade, colonoscopia 1-2 anos, rastreio displasia após 8 anos doença.',

  updated_at = NOW()
WHERE id = '94b650ab-27b5-429d-a9df-4b53569231dc';

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '94b650ab-27b5-429d-a9df-4b53569231dc'
FROM articles a
WHERE a.title ILIKE '%intestin%' OR a.title ILIKE '%colite%' OR a.title ILIKE '%inflamatória%' OR a.title ILIKE '%gastrointestinal%'
LIMIT 5
ON CONFLICT DO NOTHING;

-- Item 5: Doença Celíaca
UPDATE score_items
SET
  clinical_relevance = 'Doença Celíaca (DC) é enteropatia autoimune desencadeada por glúten (gliadina do trigo, secalina centeio, hordeína cevada) em geneticamente suscetíveis (HLA-DQ2 95%, HLA-DQ8 5%). Prevalência: 1% população mundial, subdiagnosticada. Fisiopatologia: gliadina atravessa barreira, desamidada transglutaminase tecidual (tTG), apresentada APCs via HLA-DQ2/8, resposta Th1 → atrofia vilositária, hiperplasia criptas, linfócitos intraepiteliais >25/100 enterócitos. Consequências: má absorção ferro, cálcio, zinco, vitaminas B9/B12/D/A/E/K, proteínas, gorduras. Manifestações clássicas (30%): diarreia crônica, esteatorreia, distensão, perda peso, desnutrição. Não-clássicas (70%): anemia ferropriva refratária, osteoporose, infertilidade, abortos recorrentes, hipoplasia esmalte dentário, dermatite herpetiforme, ataxia cerebelar, neuropatia periférica, estatura baixa, transaminases elevadas. Associações: DM1 (8%), tireoidite Hashimoto (15%), síndrome Sjögren, déficit IgA seletivo. Complicações: linfoma intestinal células T (risco 2-3x se não tratada), adenocarcinoma intestino delgado, osteoporose (50%), infertilidade. Na MFI: rastreio familiares 1º grau (10-15% risco), investigar deficiências múltiplas, permeabilidade persistente, contaminação cruzada, sensibilidade aveia.',

  patient_explanation = 'Doença celíaca é reação autoimune ao glúten (proteína do trigo, centeio, cevada). Quando você come glúten, seu sistema imunológico ataca a parede do intestino delgado, destruindo as vilosidades (dobrinhas que absorvem nutrientes). Resultado: você não absorve bem vitaminas, minerais, proteínas. Sintomas variam muito: algumas pessoas têm diarreia intensa, barriga inchada, desnutrição (forma clássica). Mas maioria tem sintomas "silenciosos": anemia que não melhora com ferro, osteoporose jovem, dificuldade engravidar, abortos repetidos, formigamento nas mãos/pés, lesões na pele (dermatite herpetiforme), alterações hepáticas. Há associação forte com diabetes tipo 1 e tireoidite de Hashimoto. Diagnóstico: exame de sangue (anti-transglutaminase, anti-endomísio) + biópsia intestinal mostrando destruição das vilosidades. ÚNICO TRATAMENTO: dieta 100% sem glúten para vida toda. Nada de "um pouquinho não faz mal" - mesmo traços microscópicos causam inflamação. Com dieta rigorosa, intestino cicatriza em 6-12 meses e deficiências melhoram. Precisa suplementar vitaminas/minerais inicialmente.',

  conduct = 'Sorologia (EM USO GLÚTEN): IgA total (excluir deficiência), anti-transglutaminase IgA (tTG-IgA, sensibilidade 98%), anti-endomísio IgA (EMA-IgA, especificidade 99%), anti-gliadina desamidada IgG (DGP-IgG se déficit IgA). Confirmação: endoscopia digestiva alta com biópsias duodeno distal (mínimo 4-6 fragmentos): Marsh 3 (atrofia vilositária), aumento linfócitos intraepiteliais, hiperplasia criptas. Genotipagem HLA-DQ2/DQ8 se dúvida (valor preditivo negativo 99%). Avaliar deficiências: hemograma (anemia microcítica), ferro, ferritina, zinco, cobre, selênio, vitamina D, B12, folato, vitaminas A/E/K, cálcio, magnésio, albumina, pré-albumina, perfil lipídico (hipolipoproteinemia), transaminases, densitometria óssea, TSH/anti-TPO (rastreio Hashimoto). Protocolo: dieta RIGOROSAMENTE sem glúten (trigo, centeio, cevada, malte, contaminação cruzada <20ppm). Suplementação inicial: multivitamínico, ferro 100-200mg (anemia), vitamina D3 10.000UI, cálcio 1.200mg, zinco 30mg, B12 1.000mcg, folato 5mg, probióticos (restaurar microbiota). Reavaliação: tTG-IgA 6-12 meses (normalização), biópsia controle 12-24 meses (cicatrização mucosa). Rastreio familiares 1º grau. Monitorar adesão, vigilância osteoporose, colonoscopia se >50 anos.',

  updated_at = NOW()
WHERE id = '4fbb0fe3-a93b-402c-9375-02d892ce7440';

INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '4fbb0fe3-a93b-402c-9375-02d892ce7440'
FROM articles a
WHERE a.title ILIKE '%celíac%' OR a.title ILIKE '%glúten%' OR a.title ILIKE '%intestin%' OR a.title ILIKE '%autoimun%'
LIMIT 5
ON CONFLICT DO NOTHING;
