-- =====================================================
-- BATCH 6.2: HISTÓRICO DE DOENÇAS PARTE 2
-- 39 items: Cirurgias + Medicamentos + Vícios + Saúde Bucal
-- Medicina Funcional Integrativa
-- =====================================================

-- Item 1: Derivação ventrículo-peritoneal
UPDATE score_items
SET
  clinical_significance = 'Procedimento neurocirúrgico que cria uma via de drenagem do líquido cefalorraquidiano (LCR) dos ventrículos cerebrais para a cavidade peritoneal. Indicado principalmente em hidrocefalia de diversas etiologias. A presença dessa derivação tem implicações sistêmicas importantes na MFI.',

  functional_aspects = 'Fisiopatologia: Hidrocefalia ocorre por desequilíbrio entre produção (plexos coroides) e drenagem de LCR, levando a hipertensão intracraniana. Na MFI, investigamos causas raiz como infecções prévias (meningites), hemorragias, malformações congênitas, inflamação crônica de baixo grau que pode afetar circulação liquórica. Impacto sistêmico: Válvula representa corpo estranho com risco de colonização bacteriana, inflamação crônica de baixo grau, alteração da pressão intracraniana (mesmo controlada), possível impacto em barreira hematoencefálica. Pode haver disfunção cognitiva residual, alterações neuroendócrinas secundárias.',

  risk_factors = 'Causas raiz MFI: (1) Inflamação sistêmica crônica que pode afetar produção/drenagem LCR; (2) Desequilíbrios nutricionais (vitamina A, ômega-3) que afetam função coroide; (3) Disbiose intestinal com eixo intestino-cérebro comprometido; (4) Estresse oxidativo cerebral; (5) Disfunção mitocondrial neuronal; (6) Metais pesados neurotóxicos; (7) Alterações na metilação afetando desenvolvimento neural; (8) Deficiências de folato/B12 em casos congênitos.',

  recommendations = 'Abordagem MFI: (1) Suporte anti-inflamatório: ômega-3 (EPA 2-3g/dia), curcumina lipossomal (500-1000mg), resveratrol; (2) Neuroproteção: magnésio L-treonato (2000mg), NAC (1200-1800mg), coenzima Q10 ubiquinol (200-400mg); (3) Suporte mitocondrial: L-carnitina, ácido alfa-lipóico, PQQ; (4) Modulação intestino-cérebro: probióticos psicotrópicos, prebióticos, dieta anti-inflamatória; (5) Otimização de micronutrientes: vitamina D >50ng/mL, B12, folato metilado, zinco; (6) Avaliação metais pesados e quelação se necessário; (7) Manejo de estresse: mindfulness, técnicas respiratórias; (8) Monitoramento neurológico regular; (9) Avaliação cognitiva funcional periódica; (10) Protocolos de neuroplasticidade.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'c9d8939a-e776-43a2-8d2d-a11801e9f19d';

-- Item 2: Esplenectomia
UPDATE score_items
SET
  clinical_significance = 'Remoção cirúrgica do baço, órgão linfoide crucial para imunidade, filtração sanguínea e reserva de plaquetas. Indicada em traumas esplênicos graves, doenças hematológicas (PTI, esferocitose), linfomas ou hiperesplenismo. A ausência do baço cria vulnerabilidade imunológica permanente.',

  functional_aspects = 'Fisiopatologia: Baço filtra patógenos encapsulados (pneumococo, meningococo, H. influenzae), remove eritrócitos senescentes, armazena plaquetas (30% pool corporal), produz IgM e opsoninas. Pós-esplenectomia: (1) Risco aumentado de sepse fulminante (OPSI - overwhelming post-splenectomy infection) com mortalidade 50-70%; (2) Trombocitose reativa crônica; (3) Defeito na resposta IgM; (4) Acúmulo de corpos de Howell-Jolly; (5) Microbiota alterada por mudança imunológica. Na MFI: Foco em compensação imunológica hepática/medular, modulação inflamatória intestinal, suporte de mucosas (primeira barreira).',

  risk_factors = 'Causas raiz MFI: (1) Disfunção imune de mucosas (intestino, respiratória) como nova primeira linha; (2) Disbiose com translocação bacteriana aumentada; (3) Inflamação crônica de baixo grau não controlada; (4) Deficiências de IgA secretora; (5) Estresse oxidativo prejudicando células NK residuais; (6) Depleção de micronutrientes imunomoduladores (vitamina D, zinco, selênio); (7) Disfunção mitocondrial em células imunes; (8) Desregulação do eixo HPA com cortisol afetando imunidade.',

  recommendations = 'Protocolo MFI pós-esplenectomia: (1) VACINAÇÃO obrigatória: pneumocócica (13-valente + 23-valente), meningocócica ACWY + B, H. influenzae tipo b, influenza anual; (2) Antibiótico profilático em viagens/procedimentos (amoxicilina 500mg); (3) Imunomodulação: vitamina D 5000-10000UI/dia (alvo >60ng/mL), zinco 30-50mg, selênio 200mcg, vitamina C 2-3g/dia; (4) Suporte de mucosas: IgA colostral, L-glutamina 10-15g, probióticos Lactobacillus rhamnosus GG + Saccharomyces boulardii; (5) Anti-inflamatórios naturais: ômega-3, curcumina, quercetina; (6) Modulação intestinal: dieta baixa FODMAP se necessário, prebióticos; (7) Suporte antioxidante: NAC, glutationa lipossomal; (8) Evitar álcool excessivo (sobrecarga hepática); (9) Higiene rigorosa, evitar exposição a doentes; (10) Cartão de emergência indicando asplenia.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'e355a9e4-2b73-4461-9f21-b414e3f4c425';

-- Item 3: Herniorrafia
UPDATE score_items
SET
  clinical_significance = 'Correção cirúrgica de hérnias da parede abdominal (inguinal, umbilical, incisional, femoral). Procedimento comum que envolve reforço da parede com ou sem tela sintética. Na MFI, o histórico de hérnia indica fragilidade do tecido conectivo e possíveis desequilíbrios sistêmicos na síntese de colágeno.',

  functional_aspects = 'Fisiopatologia: Hérnia resulta de fraqueza na fáscia/aponeurose por: (1) Defeitos congênitos na síntese de colágeno; (2) Degradação aumentada da matriz extracelular (MMP/TIMP desbalanceados); (3) Pressão intra-abdominal elevada crônica (obesidade, tosse crônica, constipação); (4) Falha na cicatrização. Na MFI: Investigar deficiências nutricionais (vitamina C, prolina, lisina, cobre, zinco) que prejudicam síntese de colágeno tipo I/III, estresse oxidativo que degrada matriz, inflamação crônica que ativa metaloproteinases. Pós-operatório: tela sintética pode gerar reação inflamatória crônica de baixo grau, dor crônica (neuralgia), disbiose local.',

  risk_factors = 'Causas raiz MFI para fragilidade conectiva: (1) Deficiência de vitamina C (cofator hidroxilação prolina/lisina); (2) Depleção de aminoácidos glicina, prolina, lisina; (3) Deficiências de cobre, zinco, manganês (cofatores lisil oxidase); (4) Estresse oxidativo excessivo degradando colágeno; (5) Inflamação crônica com ativação de MMPs; (6) Alterações genéticas (polimorfismos COL1A1, COL3A1); (7) Disbiose com endotoxemia afetando síntese hepática; (8) Hipercortisolismo crônico (catabólico); (9) Hipotireoidismo (síntese proteica reduzida); (10) Tabagismo (isquemia microvascular).',

  recommendations = 'Protocolo MFI para integridade conectiva: (1) Suporte colágeno: peptídeos de colágeno hidrolisado 10-15g/dia, vitamina C 2-3g/dia, glicina 5-10g, L-prolina 2-3g, L-lisina 1-2g; (2) Cofatores: cobre 2mg, zinco 30mg, manganês 5mg, silício 10-20mg; (3) Antioxidantes: NAC 1200mg, vitamina E 400UI, selênio 200mcg; (4) Anti-inflamatórios: ômega-3 EPA/DHA 2-3g, curcumina, boswellia; (5) Moduladores MMP: EGCG (chá verde), quercetina; (6) Otimização hormonal: tireoide, cortisol, testosterona/estrogênio; (7) Controle pressão abdominal: tratar constipação (fibras, magnésio, probióticos), reduzir obesidade; (8) Fisioterapia core: fortalecer musculatura abdominal profunda; (9) Evitar tabagismo; (10) Avaliação genética se hérnias recorrentes.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '38323529-f631-4b62-a444-ddf312915ea0';

-- Item 4: Cistectomia
UPDATE score_items
SET
  clinical_significance = 'Remoção cirúrgica da vesícula biliar, geralmente por colelitíase ou colecistite. Procedimento comum (>1 milhão/ano nos EUA) que altera permanentemente a fisiologia biliar. Na MFI, ausência da vesícula tem implicações em digestão de gorduras, microbiota intestinal e metabolismo lipídico.',

  functional_aspects = 'Fisiopatologia: Vesícula concentra bile (10x) e libera pulsatilmente em resposta a CCK durante refeições gordurosas. Pós-colecistectomia: (1) Fluxo biliar contínuo e diluído no intestino; (2) Deficiência relativa de sais biliares em refeições gordurosas (má absorção); (3) Excesso biliar entre refeições (diarreia colerrética, dano mucosa); (4) Alteração microbiota (bile é antimicrobiana seletiva); (5) SIBO secundário (50% dos casos); (6) Má absorção de vitaminas lipossolúveis (A, D, E, K); (7) Aumento colesterol circulante (feedback hepático); (8) Refluxo biliar gastroduodenal. Na MFI: Foco em suporte digestivo, modulação biliar, proteção mucosa, reequilíbrio microbiota.',

  risk_factors = 'Causas raiz MFI pós-colecistectomia: (1) Depleção pool de sais biliares com reciclagem êntero-hepática ineficiente; (2) SIBO/SIFO por redução efeito antimicrobiano biliar; (3) Disbiose com aumento Clostridium, redução Lactobacillus; (4) Inflamação intestinal crônica por bile não conjugada; (5) Esteatorreia com perda calórica e deficiências lipossolúveis; (6) Estresse oxidativo por má absorção antioxidantes; (7) Disfunção mitocondrial por deficiência coenzima Q10; (8) Litíase intra-hepática residual; (9) Síndrome pós-colecistectomia (20-40%): dor, dispepsia, diarreia.',

  recommendations = 'Protocolo MFI pós-colecistectomia: (1) Suporte biliar: taurina 1-2g/dia (conjugação), colina 500mg, betaína, alcachofra, cardo mariano; (2) Enzimas digestivas: lipase 10000-20000UI com refeições gordurosas; (3) Sais biliares suplementares: ácido ursodesoxicólico 300-600mg ou ox bile 100-500mg; (4) Vitaminas lipossolúveis: D3 5000UI, K2 MK-7 200mcg, E 400UI, A 10000UI (monitorar); (5) Modulação microbiota: probióticos (Lactobacillus plantarum, Bifidobacterium), prebióticos, dieta baixa FODMAP; (6) Tratamento SIBO se presente: rifaximina + herbal antimicrobials; (7) Anti-inflamatórios mucosa: L-glutamina 10g, zinco-carnosina, omega-3; (8) Dieta fracionada: 5-6 refeições pequenas, evitar grandes volumes de gordura; (9) Fibras solúveis: psyllium (liga bile excessiva); (10) Hidratação adequada.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '949f1e95-239e-4b9c-a1eb-81ef1f2ec646';

-- Item 5: Paratireoidectomia
UPDATE score_items
SET
  clinical_significance = 'Remoção cirúrgica de uma ou mais glândulas paratireoides, geralmente por adenoma causando hiperparatireoidismo primário. As paratireoides regulam cálcio sérico via PTH (paratormônio). Pós-operatório pode cursar com hipoparatireoidismo permanente, exigindo manejo vitalício do cálcio.',

  functional_aspects = 'Fisiopatologia: PTH regula cálcio por: (1) Reabsorção óssea (osteoclastos); (2) Reabsorção renal de cálcio; (3) Ativação vitamina D (1,25-OH). Hipoparatireoidismo pós-cirúrgico: (1) Hipocalcemia (parestesias, tetania, convulsões, QT longo); (2) Hiperfosfatemia; (3) Deficiência de calcitriol ativo; (4) Calcificações ectópicas (gânglios da base, catarata); (5) Desmineralização óssea paradoxal (sem PTH para remodelação). Na MFI: Otimização de cofatores para absorção cálcio, suporte ósseo sem PTH, modulação paratireoides residuais se parcial, prevenção calcificações, suporte mitocondrial (cálcio é essencial).',

  risk_factors = 'Causas raiz MFI pós-paratireoidectomia: (1) Deficiência absoluta de PTH com desregulação cálcio; (2) Má absorção intestinal de cálcio (baixo calcitriol); (3) Resistência óssea sem PTH para remodelação; (4) Deficiências de magnésio (cofator PTH residual); (5) Vitamina K2 baixa (calcificações vasculares); (6) Acidose metabólica crônica (perda óssea); (7) Estresse oxidativo por hipocalcemia; (8) Disfunção mitocondrial (cálcio regula ciclo Krebs); (9) Alterações neuropsiquiátricas (cálcio neurotransmissão); (10) Permeabilidade intestinal (inflamação afeta absorção).',

  recommendations = 'Protocolo MFI hipoparatireoidismo: (1) Reposição cálcio: carbonato/citrato 1500-3000mg/dia fracionado (com refeições); (2) Calcitriol ativo: 0,25-2mcg/dia (não vitamina D3 simples!); (3) Magnésio: 400-800mg/dia (glicinato, malato) - essencial para função PTH residual; (4) Vitamina K2 MK-7: 200-400mcg (previne calcificações); (5) Vitamina D3: manter 40-60ng/mL (não mega doses); (6) Boro: 3-6mg (metabolismo cálcio); (7) Suporte ósseo: colágeno tipo I, silício, estrôncio, vitamina C; (8) Evitar hiperfosfatemia: limitar refrigerantes, processados; (9) Dieta alcalinizante: vegetais, reduzir proteína animal excessiva; (10) Monitoramento: cálcio sérico/urinário, fósforo, PTH (se parcial), função renal, ECG (QT); (11) Exercício resistido (osteogênico); (12) Probióticos (absorção intestinal); (13) Emergência médica se tetania (cálcio IV).',

  article_ids = ARRAY[]::uuid[]
WHERE id = '8749a67e-51db-4fc3-b4ed-3e9c8cfd5b36';

-- Item 6: Cirurgia de descolamento de retina ou glaucoma
UPDATE score_items
SET
  clinical_significance = 'Cirurgias oftalmológicas complexas para patologias que ameaçam visão. Descolamento de retina: emergência que requer retinopexia (laser, crioterapia, vitrectomia). Glaucoma: trabeculectomia ou implantes de drenagem para controlar pressão intraocular (PIO). Histórico indica vulnerabilidade neuro-oftalmológica e vascular.',

  functional_aspects = 'Fisiopatologia: Descolamento retina: separação retina neurossensorial do EPR por tração vítrea, rotura, ou exsudação. Fatores: miopia, trauma, diabetes, inflamação. Glaucoma: neuropatia óptica por PIO elevada (produção/drenagem humor aquoso) com perda campo visual irreversível. Na MFI: (1) Estresse oxidativo retiniano massivo (alta demanda metabólica); (2) Disfunção mitocondrial células ganglionares; (3) Neuroinflamação; (4) Isquemia/reperfusão; (5) Glicação AGEs em diabetes; (6) Permeabilidade vascular alterada; (7) Defeitos matriz vítrea; (8) Desregulação metaloproteinases. Cirurgia gera inflamação adicional, fibrose, catarata secundária.',

  risk_factors = 'Causas raiz MFI neuro-oftalmológicas: (1) Estresse oxidativo: baixos níveis luteína, zeaxantina, glutationa; (2) Inflamação crônica: citocinas IL-6, TNF-alfa; (3) Disfunção mitocondrial retiniana: deficiência coenzima Q10, carnitina; (4) Glicação excessiva: hiperglicemia, frutose; (5) Isquemia microvascular: aterosclerose, hipertensão; (6) Deficiências nutricionais: vitamina A, zinco, ômega-3; (7) Fototoxicidade: exposição UV/luz azul sem proteção; (8) Metais pesados neurotóxicos; (9) Homocisteína elevada (vascular); (10) Tabagismo (isquemia).',

  recommendations = 'Protocolo MFI neuroproteção ocular: (1) Antioxidantes retinianos: luteína 20mg, zeaxantina 4mg, astaxantina 6-12mg, zinco 25mg (formulação AREDS2); (2) Suporte mitocondrial: coenzima Q10 ubiquinol 200-400mg, PQQ 20mg, L-carnitina 2g, ácido alfa-lipóico 600mg; (3) Ômega-3: EPA/DHA 2-3g/dia (reduz inflamação, suporta EPR); (4) Glutationa: NAC 1200-1800mg, glutationa lipossomal; (5) Vitamina A: 5000-10000UI (rodopsina), beta-caroteno natural; (6) Controle glicêmico: dieta baixo IG, berberina 1500mg, cromo; (7) Controle pressão: magnésio, CoQ10, alho; (8) Ginkgo biloba: 120-240mg (fluxo microvascular) - cautela se anticoagulantes; (9) Curcumina: 500-1000mg (anti-inflamatório); (10) Vitaminas B: B6, B12, folato (homocisteína); (11) Proteção luz azul: óculos filtro, reduzir telas noturnas; (12) Evitar tabagismo; (13) Monitoramento oftalmológico regular; (14) Controle fatores vasculares: PA, lipídios, glicemia.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'de362b24-46a5-4998-b857-e28f32fe8a4e';

-- Item 7: Histórico de medicamentos utilizados, reações adversas, respostas terapêuticas
UPDATE score_items
SET
  clinical_significance = 'Registro sistemático de todos os medicamentos já utilizados pelo paciente, incluindo reações adversas (RAM), alergias medicamentosas e variabilidade na resposta terapêutica. Na MFI, este histórico é fundamental para identificar polimorfismos farmacogenéticos, sobrecarga de detoxificação hepática, disbiose medicamentosa e depleções nutricionais secundárias.',

  functional_aspects = 'Fisiopatologia: Resposta medicamentosa varia por: (1) Polimorfismos genéticos CYP450 (metabolizadores lentos/rápidos/ultra-rápidos); (2) Polimorfismos transportadores (ABCB1, SLCO1B1); (3) Polimorfismos alvos farmacológicos (receptores, enzimas); (4) Microbiota intestinal (metaboliza 30+ classes); (5) Função hepática/renal; (6) Status nutricional (cofatores). Reações adversas: hipersensibilidade tipo I-IV, toxicidade dose-dependente, idiossincrasias. Impacto MFI crônico: (1) Sobrecarga fase I/II detox; (2) Depleção glutationa, SAMe, cofatores; (3) Disbiose por antibióticos; (4) Permeabilidade intestinal; (5) Estresse oxidativo; (6) Disfunção mitocondrial; (7) Depleções nutricionais específicas (estatinas→CoQ10, metformina→B12, IBP→B12/magnésio).',

  risk_factors = 'Causas raiz MFI de variabilidade/reações: (1) Polimorfismos CYP2D6, CYP2C19, CYP3A4 (40% variação); (2) Polimorfismos MTHFR afetando folato/B12 (homocisteína); (3) Polimorfismos GST (glutationa-S-transferase) - detox comprometida; (4) COMT lento (catecol-O-metiltransferase) - sensibilidade estimulantes; (5) NAT2 lento (N-acetiltransferase) - toxicidade isoniazida, sulfas; (6) Disbiose severa alterando metabolismo fármacos; (7) Permeabilidade intestinal (absorção errática); (8) Disfunção hepática subclínica; (9) Depleções de cofatores (B6, B12, magnésio, zinco); (10) Carga tóxica acumulada (metais, pesticidas).',

  recommendations = 'Protocolo MFI farmacogenética e suporte: (1) Teste farmacogenético: painel CYP450, MTHFR, GST, COMT, NAT2, TPMT, ABCB1; (2) Suporte detoxificação: NAC 1200-1800mg, glutationa lipossomal, ácido alfa-lipóico, SAMe 400-800mg; (3) Cofatores fase I: riboflavina, niacina, magnésio; (4) Cofatores fase II: glicina 3-5g, taurina 1-2g, glutamina, sulfato (MSM); (5) Suporte hepático: cardo mariano (silimarina 300mg), colina, betaína; (6) Modulação microbiota: probióticos pós-antibióticos (Saccharomyces boulardii, Lactobacillus rhamnosus GG), prebióticos; (7) Restauração mucosa: L-glutamina 10-15g, zinco-carnosina; (8) Antioxidantes: vitamina C 2g, vitamina E 400UI, selênio 200mcg; (9) Reposição depletados: CoQ10 (estatinas), B12 metilada (metformina, IBP), magnésio (IBP, diuréticos), cálcio/vitamina D (corticoides), folato (metotrexato); (10) Estratégias não-farmacológicas quando possível; (11) Ajuste doses por genótipo; (12) Monitoramento hepático/renal; (13) Evitar polifarmácia desnecessária.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '53d4949a-2b9c-45aa-838b-0a0928acc4b1';

-- Item 8: Uso atual de medicamentos
UPDATE score_items
SET
  clinical_significance = 'Listagem completa e atualizada de todos os medicamentos em uso (prescritos, OTC, fitoterápicos, suplementos). Essencial para identificar interações medicamentosas, duplicações terapêuticas, cascata prescritiva (medicamento para tratar efeito de outro) e depleções nutricionais iatrogênicas. Na MFI, cada medicamento é avaliado quanto à real necessidade, alternativas não-farmacológicas e suporte para minimizar efeitos adversos.',

  functional_aspects = 'Fisiopatologia polifarmácia: (1) Interações farmacocinéticas (indução/inibição CYP450, competição transportadores); (2) Interações farmacodinâmicas (sinergismo, antagonismo); (3) Cascata prescritiva (ex: AINE→HAS→anti-hipertensivo→depleção potássio→arritmia); (4) Sobrecarga metabólica hepática/renal; (5) Depleções nutricionais múltiplas; (6) Disbiose cumulativa; (7) Estresse oxidativo; (8) Disfunção mitocondrial; (9) Risco quedas (sedativos, anti-hipertensivos); (10) Declínio cognitivo (anticolinérgicos). MFI foco: desprescrição criteriosa, otimização doses, suporte detox, reposição cofatores, alternativas fitoterapêuticas.',

  risk_factors = 'Riscos MFI polifarmácia: (1) >5 medicamentos = risco exponencial RAM; (2) Idosos: farmacocinética alterada (clearance reduzido); (3) Anticoagulantes + antiplaquetários + AINEs = sangramento; (4) Sedativos múltiplos = quedas, fraturas; (5) Anticolinérgicos (antidepressivos, anti-histamínicos, antimuscarínicos) = demência; (6) IBP crônico = deficiência B12, magnésio, cálcio, SIBO, pneumonia; (7) Estatinas = depleção CoQ10, miopatia, diabetes; (8) Metformina = deficiência B12, acidose láctica (raro); (9) Corticoides = síndrome Cushing iatrogênica, osteoporose, hiperglicemia; (10) Benzodiazepínicos crônicos = dependência, cognição, quedas.',

  recommendations = 'Protocolo MFI revisão medicamentosa: (1) Desprescrição: critérios STOPP/START, Beers Criteria (idosos), avaliar risco-benefício; (2) Otimização horários: cronoterapia (estatinas noite, corticoides manhã); (3) Suporte específico: CoQ10 100-200mg (estatinas), B12 1000mcg (metformina/IBP), magnésio 400mg (IBP/diuréticos), probióticos (antibióticos/IBP), vitamina D+K2+cálcio (corticoides); (4) Detoxificação: NAC, glutationa, suporte hepático (cardo mariano); (5) Alternativas fitoterapêuticas: berberina (metformina), arroz vermelho (estatinas leves), valeriana/passiflora (benzodiazepínicos), cúrcuma (AINEs); (6) Intervenções estilo de vida: dieta anti-inflamatória mediterrânea, exercício (HAS, diabetes), higiene do sono (sedativos), mindfulness (ansiolíticos); (7) Monitoramento: função hepática/renal, eletrólitos, glicemia, lipídios, B12, vitamina D; (8) Educação paciente: adesão, horários, sinais alerta; (9) Reconciliação medicamentosa periódica; (10) Evitar cascata prescritiva.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '682decb0-659a-48da-9d77-4e8258d3cf8e';

-- Item 9: Broncodilatadores / Corticosteroides inalatórios
UPDATE score_items
SET
  clinical_significance = 'Medicamentos fundamentais para asma e DPOC. Broncodilatadores (beta-2 agonistas, anticolinérgicos) relaxam músculo liso brônquico. Corticosteroides inalatórios (CI) reduzem inflamação das vias aéreas. Uso prolongado tem implicações sistêmicas apesar da via inalatória.',

  functional_aspects = 'Fisiopatologia: Asma/DPOC envolvem inflamação crônica Th2 (asma) ou Th1/Th17 (DPOC), remodelamento brônquico, hiperreatividade. Beta-2 agonistas: estimulam AMPc, relaxamento. Anticolinérgicos: bloqueiam acetilcolina, broncodilatação. CI: suprimem citocinas inflamatórias, estabilizam mastócitos. Efeitos sistêmicos CI: (1) Absorção 10-20% (orofaringe, pulmonar); (2) Supressão eixo HPA (doses altas/crônicas); (3) Osteoporose (osteoblastos); (4) Candidíase oral; (5) Disfonia; (6) Hiperglicemia leve; (7) Catarata/glaucoma (raro). MFI: reduzir inflamação pulmonar de base, otimizar função mitocondrial, suporte adrenal.',

  risk_factors = 'Causas raiz MFI inflamação pulmonar: (1) Disbiose intestinal (eixo intestino-pulmão); (2) Alergenos alimentares (glúten, lácteos, ovos); (3) Sensibilidades químicas (VOCs, fragrâncias); (4) Deficiências antioxidantes (glutationa pulmonar); (5) Vitamina D baixa (imunomodulação); (6) Magnésio baixo (broncodilatação endógena); (7) Ômega-3 baixo (resolvinas anti-inflamatórias); (8) Estresse oxidativo; (9) Disfunção mitocondrial; (10) Permeabilidade intestinal aumentada.',

  recommendations = 'Protocolo MFI asma/DPOC: (1) Anti-inflamatórios naturais: ômega-3 EPA 2-3g/dia, curcumina lipossomal 500-1000mg, boswellia 400mg, quercetina 500-1000mg; (2) Suporte respiratório: NAC 600-1200mg (mucolítico, antioxidante), vitamina C 2-3g; (3) Broncodilatação natural: magnésio 400-600mg, ginkgo biloba 120mg; (4) Imunomodulação: vitamina D3 5000-10000UI (alvo >50ng/mL), probióticos (Lactobacillus rhamnosus, Bifidobacterium); (5) Antioxidantes pulmonares: glutationa lipossomal, SOD; (6) Suporte adrenal (se CI): vitamina C 2g, B5 500mg, adaptógenos (ashwagandha, rhodiola); (7) Prevenção osteoporose: cálcio 1000mg, vitamina D3, K2 200mcg, exercício resistido; (8) Dieta anti-inflamatória: eliminar alergenos comuns, aumentar vegetais coloridos; (9) Modulação microbiota: prebióticos, probióticos, evitar antibióticos desnecessários; (10) Controle ambiental: purificadores HEPA, evitar VOCs; (11) Técnicas respiratórias: Buteyko, yoga; (12) Higiene bucal: enxágue pós-CI (prevenir candidíase); (13) Monitorar: função pulmonar, cortisol se sintomas, densitometria óssea.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'd8736b07-be1f-4f5b-93c1-4590ed06b91f';

-- Item 10: Corticosteroides sistêmicos
UPDATE score_items
SET
  clinical_significance = 'Corticoides sistêmicos (prednisona, metilprednisolona, hidrocortisona, dexametasona) são anti-inflamatórios e imunossupressores potentes usados em doenças autoimunes, inflamatórias, alergias graves. Uso crônico (>3 meses) ou repetido tem efeitos adversos múltiplos e graves, criando "síndrome de Cushing iatrogênica".',

  functional_aspects = 'Fisiopatologia: Corticoides ligam-se a receptores citoplasmáticos, modulam transcrição gênica: (1) Supressão NFkB (anti-inflamatório); (2) Redução citocinas IL-1, IL-6, TNF-alfa; (3) Inibição fosfolipase A2 (reduz prostaglandinas/leucotrienos); (4) Imunossupressão (linfócitos T/B). Efeitos adversos crônicos: (1) Supressão eixo HPA (insuficiência adrenal); (2) Síndrome Cushing: face em lua, obesidade central, estrias, hirsutismo; (3) Osteoporose (osteoblastos inibidos, osteoclastos ativados); (4) Sarcopenia (catabolismo proteico); (5) Hiperglicemia/diabetes (resistência insulina, gliconeogênese); (6) HAS (retenção sódio); (7) Dislipidemia; (8) Úlceras gástricas; (9) Catarata/glaucoma; (10) Infecções oportunistas; (11) Osteonecrose; (12) Alterações psiquiátricas (ansiedade, depressão, psicose).',

  risk_factors = 'Riscos MFI corticoterapia crônica: (1) Depleção vitamina D, cálcio (ósseo); (2) Deficiência vitaminas B (metabolismo glicose); (3) Depleção magnésio, potássio (arritmias); (4) Permeabilidade intestinal aumentada; (5) Disbiose com Candida overgrowth; (6) Estresse oxidativo massivo; (7) Disfunção mitocondrial; (8) Resistência insulínica/síndrome metabólica; (9) Atrofia muscular e adiposidade visceral; (10) Imunossupressão com infecções recorrentes.',

  recommendations = 'Protocolo MFI proteção corticoterapia: (1) Prevenção osteoporose: cálcio 1200-1500mg, vitamina D3 2000-5000UI, vitamina K2 MK-7 200mcg, magnésio 400mg, exercício resistido; considerar bifosfonatos se >7,5mg prednisona >3 meses; (2) Controle glicêmico: dieta baixo IG, cromo 200-400mcg, ácido alfa-lipóico 600mg, berberina 1500mg, canela; (3) Proteção muscular: proteína 1,5-2g/kg, leucina 3-5g, creatina 5g, HMB 3g, exercício resistido; (4) Suporte adrenal (desmame): vitamina C 2-3g, B5 500mg, DHEA 25-50mg (se baixo), adaptógenos (ashwagandha 600mg, rhodiola 400mg); (5) Proteção GI: probióticos, L-glutamina 10g, zinco-carnosina, evitar AINEs; (6) Antioxidantes: NAC 1200mg, glutationa, vitamina E 400UI, selênio 200mcg; (7) Anti-fúngico: Saccharomyces boulardii, óleo orégano, alho; (8) Controle PA: magnésio, CoQ10, reduzir sódio; (9) Lipídios: ômega-3 2-3g, niacina, fibras; (10) Saúde ocular: luteína, zeaxantina, exames regulares; (11) Monitoramento: glicemia, PA, lipídios, densitometria óssea, cortisol (desmame), eletrólitos; (12) Desmame gradual (não parar abruptamente!); (13) Cartão de emergência indicando corticoterapia.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '7383daa3-1138-4f37-99af-bba1e5d57177';

-- Item 11: Digitálicos
UPDATE score_items
SET
  clinical_significance = 'Digoxina (digitálico) é inotrópico positivo e cronotrópico negativo usado em insuficiência cardíaca (IC) e fibrilação atrial. Mecanismo: inibe bomba Na+/K+ ATPase, aumenta cálcio intracelular, aumenta contratilidade. Margem terapêutica estreita (0,5-2ng/mL), risco toxicidade. Uso indica doença cardíaca estrutural significativa.',

  functional_aspects = 'Fisiopatologia IC: disfunção sistólica/diastólica com débito reduzido, congestão pulmonar, ativação neuro-hormonal (SRAA, SNS). Digoxina: aumenta contratilidade, reduz frequência (bloqueia nó AV), efeito neurohormonal (reduz norepinefrina). Toxicidade digitálica: (1) Arritmias (BAV, taquicardia atrial, bigeminismo ventricular); (2) Sintomas GI (náusea, vômito, anorexia); (3) Visuais (visão amarelada, halos); (4) Neurológicos (confusão, fadiga). Fatores predisponentes: hipocalemia, hipomagnesemia, hipercalcemia, hipotireoidismo, função renal reduzida, idade avançada. MFI: otimizar função mitocondrial cardíaca, suporte iônico, reduzir carga.',

  risk_factors = 'Causas raiz MFI IC e risco digitálico: (1) Disfunção mitocondrial cardiomiócitos (IC de base); (2) Deficiência coenzima Q10, L-carnitina, D-ribose; (3) Depleção magnésio (fundamental para bomba Na+/K+); (4) Depleção potássio (diuréticos concomitantes); (5) Estresse oxidativo cardíaco; (6) Inflamação crônica com remodelamento; (7) Deficiência tiamina (IC alto débito); (8) Taurina baixa (estabilizador membrana); (9) Isquemia miocárdica; (10) Sobrecarga volume/pressão.',

  recommendations = 'Protocolo MFI IC + digoxina: (1) Suporte energético cardíaco: coenzima Q10 ubiquinol 200-400mg/dia, L-carnitina 2-3g, D-ribose 5g 3x/dia, ácido alfa-lipóico 600mg; (2) Eletrólitos críticos: magnésio 400-600mg (glicinato, taurato) - ESSENCIAL, potássio alimentar abundante (vegetais, frutas), evitar hipercalcemia; (3) Taurina: 3-6g/dia (antiarrítmico natural, estabiliza membrana); (4) Hawthorn (Crataegus): 900-1800mg (inotrópico natural, reduz pré/pós-carga) - CAUTELA com digoxina (sinergismo); (5) Tiamina: 100-300mg se IC; (6) Ômega-3: EPA/DHA 2-3g (anti-inflamatório, antiarrítmico); (7) Antioxidantes: NAC, vitamina E, selênio; (8) Controle PA: magnésio, CoQ10, alho, reduzir sódio <2g/dia; (9) Dieta: mediterrânea, DASH, restrição sódio/fluidos se congestão; (10) Monitoramento rigoroso: digoxinemia (0,5-0,9 ideal), função renal, eletrólitos (magnésio, potássio, cálcio), ECG, sintomas toxicidade; (11) Interações: evitar amiodarona, verapamil, claritromicina (aumentam digoxinemia); (12) Exercício supervisionado (reabilitação cardíaca); (13) Alerta para sinais toxicidade: náusea, visão alterada, bradicardia, arritmias.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '594b827b-e718-4982-a41f-ea8bb0eb9da4';

-- Item 12: Disfunção erétil (medicamentos)
UPDATE score_items
SET
  clinical_significance = 'Inibidores PDE5 (sildenafil, tadalafil, vardenafil) são primeira linha para disfunção erétil (DE). Mecanismo: inibem fosfodiesterase-5, aumentam GMPc, relaxam músculo liso vascular peniano. Uso crônico indica DE orgânica (vascular, neurogênica, hormonal) ou psicogênica. Na MFI, DE é marcador precoce de disfunção endotelial sistêmica e risco cardiovascular.',

  functional_aspects = 'Fisiopatologia DE: ereção requer (1) NO endotelial→GMPc→vasodilatação; (2) fluxo arterial adequado; (3) oclusão venosa; (4) inervação intacta; (5) testosterona adequada. Causas: (1) Disfunção endotelial (aterosclerose, diabetes, HAS, tabagismo) - 80% dos casos; (2) Neuropatia (diabetes, cirurgia pélvica); (3) Hipogonadismo; (4) Psicogênica (ansiedade, depressão); (5) Medicamentosa (anti-hipertensivos, antidepressivos, finasterida). MFI: DE como "canário na mina de carvão" - aterosclerose peniana precede coronária em 2-5 anos. Foco: otimizar saúde vascular, NO, testosterona, saúde metabólica.',

  risk_factors = 'Causas raiz MFI DE: (1) Disfunção endotelial: estresse oxidativo, inflamação, glicação; (2) Deficiência NO: arginina baixa, BH4 baixo (cofator NOS), estresse oxidativo (NO→peroxinitrito); (3) Hipogonadismo: obesidade, diabetes, metais pesados, disruptores endócrinos, estresse crônico; (4) Resistência insulínica/síndrome metabólica; (5) Deficiências: vitamina D, zinco (testosterona), folato/B12 (NO); (6) Inflamação crônica de baixo grau; (7) Disfunção mitocondrial; (8) Homocisteína elevada (endotélio); (9) Disbiose com endotoxemia; (10) Tabagismo, álcool excessivo.',

  recommendations = 'Protocolo MFI DE (além de iPDE5): (1) Suporte NO: L-arginina 3-5g + L-citrulina 3-6g/dia (precursores NO), BH4 via folato metilado 5mg + B12 metilada 1mg; (2) Saúde endotelial: ômega-3 EPA/DHA 2-3g, vitamina C 2g, vitamina E 400UI, resveratrol 500mg, cacau rico flavonoides; (3) Otimização testosterona: zinco 30-50mg, vitamina D3 5000UI (alvo >50ng/mL), magnésio 400mg, Tribulus terrestris 750mg, Tongkat ali 200mg, ashwagandha 600mg; (4) Controle glicêmico: berberina 1500mg, cromo, ácido alfa-lipóico, canela; (5) Suporte mitocondrial: coenzima Q10 200mg, L-carnitina 2g, PQQ; (6) Homocisteína: folato, B6, B12, TMG; (7) Anti-inflamatórios: curcumina, boswellia, quercetina; (8) Estilo de vida: exercício aeróbico + resistido 150min/semana, perda peso se IMC>25, dieta mediterrânea, reduzir álcool, parar tabagismo; (9) Suporte psicológico: terapia cognitivo-comportamental, mindfulness, redução estresse; (10) Otimizar sono: 7-9h (testosterona); (11) Evitar disruptores endócrinos: plásticos, pesticidas; (12) Monitorar: testosterona total/livre, SHBG, estradiol, lipídios, glicemia/HbA1c, homocisteína, vitamina D; (13) Cautela iPDE5: contraindicado com nitratos, cautela se doença CV instável.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'ffcb5134-09e5-437c-a061-b81dddab34cd';

-- Item 13: Estatinas / Hipolipemiantes
UPDATE score_items
SET
  clinical_significance = 'Estatinas (atorvastatina, rosuvastatina, sinvastatina) inibem HMG-CoA redutase, reduzindo síntese hepática de colesterol. São os hipolipemiantes mais prescritos para prevenção cardiovascular. Uso indica dislipidemia e/ou risco CV elevado. Efeitos adversos significativos, especialmente depleção de CoQ10.',

  functional_aspects = 'Fisiopatologia: Estatinas bloqueiam via mevalonato, reduzindo: (1) Colesterol (LDL -30-50%); (2) Coenzima Q10 (ubiquinona - mesma via!); (3) Dolicol (glicosilação proteínas); (4) Heme A (cadeia respiratória); (5) Selenoproteínas. Efeitos pleiotrópicos: anti-inflamatório, estabilização placa. Efeitos adversos: (1) Miopatia/mialgia (10-15%): depleção CoQ10 mitocondrial, disfunção membrana; (2) Rabdomiólise (rara, grave); (3) Hepatotoxicidade (transaminases); (4) Diabetes novo (10-20% aumento risco); (5) Declínio cognitivo (colesterol essencial para sinapse); (6) Depleção vitamina K2 (calcificação arterial paradoxal); (7) Neuropatia periférica; (8) Disfunção sexual.',

  risk_factors = 'Causas raiz MFI dislipidemia: (1) Resistência insulínica (VLDL aumentado, HDL reduzido); (2) Inflamação crônica (LDL oxidado); (3) Estresse oxidativo; (4) Disfunção tireoidiana (TSH>2,5); (5) Disbiose com endotoxemia (LPS aumenta LDL); (6) Deficiências: magnésio, ômega-3, niacina, fibras; (7) Dieta pró-inflamatória (trans, açúcares, ômega-6); (8) Sedentarismo; (9) Genética (FH, ApoE4); (10) Metais pesados.',

  recommendations = 'Protocolo MFI com estatinas: (1) Reposição CoQ10 OBRIGATÓRIA: ubiquinol 100-300mg/dia (forma reduzida, biodisponível) - reduz mialgia 40%; (2) Vitamina K2 MK-7: 200-400mcg (previne calcificação arterial); (3) Vitamina D3: 5000UI (músculo, inflamação); (4) Magnésio: 400-600mg (músculo, insulina); (5) Ômega-3: EPA/DHA 2-4g (triglicérides, inflamação); (6) Niacina: 500-2000mg (HDL, Lp(a)) - cautela com estatina (hepatotoxicidade); (7) Bergamota: 500-1000mg (hipolipemiante natural); (8) Berberina: 1500mg (LDL, glicemia); (9) Fibras solúveis: psyllium 10g, beta-glucana; (10) Antioxidantes: NAC, vitamina E, resveratrol; (11) Dieta: mediterrânea, aumentar vegetais/fibras, reduzir carboidratos refinados; (12) Exercício: 150min aeróbico + resistido; (13) Monitorar: lipidograma avançado (LDL-P, ApoB, Lp(a)), transaminases, CK, glicemia/HbA1c, CoQ10 sérico, TSH; (14) Avaliar necessidade: calcular risco ASCVD, considerar escore cálcio coronário; (15) Alternativas se intolerância: ezetimibe, PCSK9i, bergamota, berberina, arroz vermelho (contém lovastatina natural); (16) Tomar à noite (síntese colesterol noturna).',

  article_ids = ARRAY[]::uuid[]
WHERE id = '1fcbc463-f274-4c2d-b263-5ba6bc208d28';

-- Item 14: Estimulantes (metilfenidato, anfetaminas)
UPDATE score_items
SET
  clinical_significance = 'Psicoestimulantes (metilfenidato, lisdexanfetamina, anfetaminas) são primeira linha para TDAH. Mecanismo: bloqueiam recaptação/aumentam liberação de dopamina e norepinefrina. Uso crônico indica TDAH (déficit executivo, atencional) ou narcolepsia. Na MFI, avaliamos causas raiz de sintomas tipo TDAH (sono, nutrição, inflamação, metais pesados) e efeitos adversos a longo prazo.',

  functional_aspects = 'Fisiopatologia TDAH: (1) Disfunção dopaminérgica córtex pré-frontal; (2) Déficit norepinefrina (atenção); (3) Hipoativação circuitos fronto-estriatais; (4) Polimorfismos (DAT1, DRD4, COMT). Estimulantes: aumentam dopamina/norepinefrina, melhoram função executiva, atenção, controle impulsos. Efeitos adversos: (1) Cardiovasculares: taquicardia, HAS; (2) Apetite reduzido, perda peso; (3) Insônia; (4) Ansiedade, irritabilidade; (5) Crescimento reduzido (crianças); (6) Tiques; (7) Abuso/dependência; (8) Depleção neurotransmissores (uso crônico); (9) Estresse oxidativo neuronal. MFI: considerar causas tipo TDAH tratáveis (apneia sono, anemia ferropriva, intoxicação chumbo, sensibilidades alimentares, disbiose).',

  risk_factors = 'Causas raiz MFI sintomas tipo TDAH: (1) Deficiência ferro (cofator tirosina hidroxilase→dopamina); (2) Deficiência magnésio (NMDA, ansiedade); (3) Deficiência zinco (neurotransmissão); (4) Deficiência ômega-3 (membrana neuronal, inflamação); (5) Deficiências B6, B12, folato (metilação, neurotransmissores); (6) Intoxicação chumbo, mercúrio (neurotóxicos); (7) Sensibilidades alimentares (corantes, glúten, lácteos); (8) Disbiose intestinal (eixo intestino-cérebro); (9) Privação sono/apneia; (10) Estresse crônico (cortisol afeta PFC); (11) Tela excessiva; (12) Polimorfismos MTHFR, COMT.',

  recommendations = 'Protocolo MFI TDAH: (1) Suporte dopaminérgico: L-tirosina 1-3g/dia, mucuna pruriens (L-dopa natural) 300mg, SAMe 400-800mg; (2) Ferro: 30-60mg se ferritina <50ng/mL (ESSENCIAL); (3) Magnésio: 400-600mg (glicinato, treonato); (4) Zinco: 30mg; (5) Ômega-3: EPA/DHA 2-3g (evidência forte); (6) Vitaminas B: B6 50mg, B12 metilada 1mg, folato metilado 5mg (metilação); (7) Fosfatidilserina: 200-300mg (membranas neuronais); (8) Ginkgo biloba + Panax ginseng: 120mg + 200mg (atenção); (9) Rhodiola rosea: 400mg (fadiga, foco); (10) Probióticos psicotrópicos: Lactobacillus rhamnosus, Bifidobacterium longum; (11) Dieta: eliminar corantes artificiais, reduzir açúcares/refinados, aumentar proteína (aminoácidos), dieta Feingold; (12) Sono: higiene rigorosa, 8-10h (crianças); (13) Exercício: 60min diário (dopamina, norepinefrina); (14) Mindfulness: meditação, neurofeedback; (15) Limitar telas: <2h/dia; (16) Avaliação metais pesados: chumbo, mercúrio (quelação se necessário); (17) Monitorar com estimulantes: PA, FC, peso, altura (crianças), apetite, sono, humor; (18) Férias de medicação (fins de semana, férias) - avaliar com médico.',

  article_ids = ARRAY[]::uuid[]
WHERE id = '2319cba0-9ba5-4b72-94e7-367771f837bd';

-- Item 15: Hormônios tireoidianos / Antitireoidianos
UPDATE score_items
SET
  clinical_significance = 'Levotiroxina (T4) para hipotireoidismo, liotironina (T3) para casos específicos, metimazol/propiltiouracil para hipertireoidismo. Tireoide regula metabolismo basal, termogênese, função cardiovascular, cognitiva, reprodutiva. Uso indica disfunção tireoidiana, autoimune (Hashimoto, Graves) ou não.',

  functional_aspects = 'Fisiopatologia: Hipotireoidismo (TSH alto): (1) Hashimoto (autoimune, 90% dos casos); (2) Pós-tireoidectomia/radioiodo; (3) Deficiência iodo; (4) Medicamentoso (lítio, amiodarona). Hipertireoidismo (TSH baixo): (1) Graves (autoimune); (2) Nódulos tóxicos; (3) Tireoidite. T4→T3 conversão: desiodinases D1/D2 (selenoproteínas), inibidas por estresse, inflamação, metais pesados, fluoreto. MFI foco: causas raiz autoimunidade (permeabilidade intestinal, disbiose, glúten, deficiências), otimizar conversão T4→T3, suporte mitocondrial (hormônios tireoidianos regulam biogênese mitocondrial).',

  risk_factors = 'Causas raiz MFI disfunção tireoidiana: (1) Autoimunidade: permeabilidade intestinal, disbiose, glúten (mimetismo molecular), estresse; (2) Deficiências: iodo (150-300mcg), selênio (200mcg - desiodinases, reduz TPO-Ab), zinco, ferro, vitamina D, vitamina A; (3) Toxinas: fluoreto, brometo, perclorato (competem iodo); mercúrio, cádmio (autoimunidade); (4) Estresse crônico: cortisol inibe TSH, conversão T4→T3; (5) Inflamação: citocinas inibem desiodinases; (6) Resistência insulínica; (7) Disbiose: 20% T4→T3 no intestino; (8) Polimorfismos: DIO1, DIO2 (conversão); (9) Medicamentos: lítio, amiodarona, IFN-alfa.',

  recommendations = 'Protocolo MFI tireoide: (1) Iodo: 150-300mcg/dia (algas, sal iodado) - CAUTELA em Hashimoto ativo (pode piorar); (2) Selênio: 200mcg (selenometionina) - reduz TPO-Ab 30-40%, suporta desiodinases; (3) Zinco: 30mg, ferro: tratar se ferritina <50ng/mL; (4) Vitamina D3: >50ng/mL (imunomodulação); (5) Vitamina A: 5000UI (receptores tireoidianos); (6) Magnésio: 400mg (ATP, centenas de reações); (7) Ômega-3: 2g (anti-inflamatório); (8) Suporte adrenal: adapógenos (ashwagandha 600mg - CAUTELA, pode aumentar T4), rhodiola; (9) Suporte conversão T4→T3: reduzir estresse, inflamação, otimizar cortisol; (10) Modulação autoimune: (a) Eliminar glúten 100% (mimetismo), (b) Permeabilidade intestinal: L-glutamina 10-15g, zinco-carnosina, probióticos, (c) Disbiose: protocolo 5R, (d) Vitamina D >50ng/mL; (11) Evitar disruptores: fluoreto (água filtrada), brometos, soja não-fermentada (goitrogênicos); (12) Detox metais pesados se indicado; (13) Suporte hepático: cardo mariano, NAC (conversão T4→T3); (14) Dieta: nutrientes densos, anti-inflamatória, considerar AIP se Hashimoto; (15) Monitorar: TSH (alvo 1-2), T4 livre, T3 livre, T3 reverso, TPO-Ab, Tg-Ab, vitamina D, ferritina, selênio; (16) Timing levotiroxina: jejum, 30-60min antes café/alimentos; (17) Considerar T3 (liotironina) se conversão prejudicada (T3 baixo, rT3 alto).',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'd70c5667-7201-4d60-aa28-180a7e672899';

-- Item 16: HPB (alfa-bloqueadores, inibidores 5-alfa-redutase)
UPDATE score_items
SET
  clinical_significance = 'Medicamentos para hiperplasia prostática benigna (HPB). Alfa-bloqueadores (tansulosina, doxazosina) relaxam músculo liso prostático/vesical. Inibidores 5-alfa-redutase (finasterida, dutasterida) bloqueiam conversão testosterona→DHT, reduzem volume prostático. Uso indica sintomas urinários obstrutivos (LUTS) em homens >50 anos.',

  functional_aspects = 'Fisiopatologia HPB: crescimento benigno zona transicional prostática por: (1) DHT (di-hidrotestosterona) estimula proliferação; (2) Desbalanço testosterona/estrogênio (estrogênio aumenta com idade); (3) Inflamação crônica; (4) Resistência insulínica. Sintomas: hesitação, jato fraco, noctúria, urgência, retenção. Alfa-bloqueadores: efeitos adversos - hipotensão ortostática, tontura, ejaculação retrógrada, fadiga. Inibidores 5AR: efeitos adversos - disfunção sexual (DE, libido reduzida, ejaculação anormal) em 5-10%, "síndrome pós-finasterida" persistente em raros casos, redução PSA 50% (ajustar rastreamento câncer). MFI: reduzir DHT naturalmente, anti-inflamatório, modular estrogênio.',

  risk_factors = 'Causas raiz MFI HPB: (1) Excesso DHT: alta atividade 5-alfa-redutase; (2) Dominância estrogênica relativa: aromatase aumentada (obesidade, idade), xenoestrogênios; (3) Inflamação prostática crônica: citocinas (IL-6, IL-8), estresse oxidativo; (4) Resistência insulínica/síndrome metabólica; (5) Deficiência zinco (inibe 5AR naturalmente); (6) Deficiência vitamina D (proliferação); (7) Disbiose intestinal (metabolismo hormonal); (8) Estresse oxidativo; (9) Sedentarismo; (10) Dieta pró-inflamatória.',

  recommendations = 'Protocolo MFI HPB: (1) Saw palmetto (Serenoa repens): 320mg/dia (inibe 5AR naturalmente, anti-inflamatório) - eficácia similar finasterida sem efeitos sexuais; (2) Beta-sitosterol: 60-130mg (melhora fluxo urinário); (3) Pygeum africanum: 100mg (anti-inflamatório prostático); (4) Licopeno: 15-30mg (antioxidante, reduz crescimento); (5) Zinco: 30-50mg (inibe 5AR, anti-inflamatório); (6) Selênio: 200mcg (antioxidante); (7) Ômega-3: EPA/DHA 2-3g (anti-inflamatório); (8) Vitamina D3: >50ng/mL (antiproliferativo); (9) Modulação estrogênio: (a) DIM 200-400mg ou I3C (metabolismo estrogênio), (b) Fibras: 30-40g (excreção estrogênio), (c) Lignanas (linhaça), (d) Cálcio-D-glucarato; (10) Chá verde EGCG: 400mg (inibe 5AR, antioxidante); (11) Quercetina: 500-1000mg (anti-inflamatório); (12) Urtica dioica (urtiga): 300mg (liga SHBG, reduz DHT livre); (13) Dieta: (a) Mediterrânea, (b) Tomate cozido (licopeno), (c) Crucíferas (DIM/I3C), (d) Reduzir lácteos (IGF-1); (14) Exercício: 150min/semana (reduz risco progressão); (15) Perda peso se IMC>25 (reduz estrogênio); (16) Evitar: álcool excessivo, cafeína excessiva (irritam bexiga), descongestionantes (constritores alfa); (17) Monitorar: IPSS (score sintomas), PSA (ajustar se 5ARi - multiplicar x2), toque retal, urofluxometria, testosterona total/livre, estradiol, função sexual; (18) Urinar antes dormir, duplo esvaziamento; (19) Considerar cirurgia (RTU, laser) se refratário ou complicações.',

  article_ids = ARRAY[]::uuid[]
WHERE id = 'c528f41c-011a-460f-baec-2680dc817e47';

-- Continua no próximo bloco...
