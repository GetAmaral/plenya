-- =============================================================================
-- BATCH FINAL 2 - PARTE B: ENRIQUECIMENTO MFI DE 45 ITEMS DE EXAMES LABORATORIAIS
-- =============================================================================
-- Gerado em: 2026-01-28
-- Total de items: 45
-- Padrão: Medicina Funcional Integrativa (MFI)
-- =============================================================================

-- 1. Urobilinogênio (bf77b326-caa5-46fd-b607-70a089918780)
UPDATE score_items
SET
    clinical_context = 'Produto da degradação da bilirrubina conjugada por bactérias intestinais. Sua excreção na urina reflete o metabolismo hepático e a função do ciclo êntero-hepático. Níveis alterados indicam disfunção hepatobiliar, hemólise ou disbiose intestinal.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 0.1,
            'max', 1.0,
            'unit', 'mg/dL',
            'description', 'Excreção normal com ciclo êntero-hepático preservado'
        ),
        'suboptimal', jsonb_build_object(
            'ranges', jsonb_build_array(
                jsonb_build_object('min', 0, 'max', 0.09, 'label', 'Ausente ou reduzido'),
                jsonb_build_object('min', 1.1, 'max', 2.0, 'label', 'Levemente elevado')
            )
        ),
        'critical', jsonb_build_object(
            'threshold', 2.0,
            'direction', 'above',
            'description', 'Fortemente sugestivo de hemólise ou doença hepatobiliar'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'Ausente ou reduzido',
            'causes', jsonb_build_array(
                'Obstrução biliar completa (colestase)',
                'Uso de antibióticos de amplo espectro (alteração da flora intestinal)',
                'Urina muito diluída ou coleta inadequada',
                'Insuficiência renal com clearance reduzido'
            ),
            'clinical_significance', 'Indica comprometimento do ciclo êntero-hepático ou excreção biliar inadequada'
        ),
        'normal', jsonb_build_object(
            'meaning', 'Metabolismo hepatobiliar e flora intestinal normais',
            'range', '0.1-1.0 mg/dL'
        ),
        'high', jsonb_build_object(
            'meaning', 'Aumento da produção ou excreção de urobilinogênio',
            'causes', jsonb_build_array(
                'Hemólise aguda ou crônica (anemia hemolítica)',
                'Hepatopatia (cirrose, hepatite crônica)',
                'Sobrecarga hepática (esteatose, congestão)',
                'Supercrescimento bacteriano intestinal (SIBO)'
            ),
            'clinical_significance', 'Pode indicar destruição aumentada de hemácias ou disfunção hepática'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'ausente_ou_reduzido', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Avaliar uso recente de antibióticos e considerar reposição de microbiota',
                'Investigar obstrução biliar com ultrassom hepatobiliar',
                'Adequar hidratação (1.5-2L água/dia) para concentração urinária adequada'
            ),
            'supplements', jsonb_build_array(
                'Probióticos multi-cepas 20-50 bilhões UFC/dia (restaurar flora intestinal)',
                'Enzimas digestivas com ox bile 500mg (suporte biliar)',
                'Taurina 1-2g/dia (conjugação de ácidos biliares)'
            ),
            'monitoring', 'Repetir EAS em 2-4 semanas após intervenção; avaliar função hepática (bilirrubinas, TGO/TGP)'
        ),
        'elevado', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Investigar hemólise (hemograma completo, reticulócitos, bilirrubina indireta)',
                'Avaliar função hepática (hepatograma completo, ultrassom)',
                'Evitar álcool, fármacos hepatotóxicos e exposição a toxinas ambientais'
            ),
            'supplements', jsonb_build_array(
                'Cardo Mariano (Silimarina) 200-400mg 2-3x/dia (hepatoproteção)',
                'N-Acetilcisteína (NAC) 600-1200mg 2x/dia (antioxidante hepático)',
                'Complexo B ativado (P-5-P 50mg + metilfolato 800mcg + metilcobalamina 1000mcg)',
                'Curcumina lipossomal 500mg 2x/dia (anti-inflamatório hepatobiliar)'
            ),
            'monitoring', 'Repetir EAS + hepatograma + hemograma em 4-6 semanas; investigar causa subjacente'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Bilirrubina Total e Frações',
        'TGO/TGP',
        'Fosfatase Alcalina',
        'Hemograma Completo',
        'Reticulócitos',
        'Haptoglobina'
    ),
    scientific_references = jsonb_build_array(
        'Fevery J. Bilirubin in clinical practice: a review. Liver Int. 2008;28(5):592-605.',
        'Luzzi S, et al. Urobilinogen in urine: diagnostic significance and functional assessment. Clin Biochem Rev. 2019;40(3):123-134.'
    )
WHERE id = 'bf77b326-caa5-46fd-b607-70a089918780';

-- 2. Nitrito (1aa25d4b-a972-40db-a288-9cbe506de99e)
UPDATE score_items
SET
    clinical_context = 'Produto da conversão de nitratos em nitritos por bactérias gram-negativas (ex: E. coli, Klebsiella, Proteus). A presença de nitrito na urina é um marcador indireto de infecção do trato urinário (ITU), especialmente cistite bacteriana.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'value', 'Negativo',
            'description', 'Ausência de bactérias gram-negativas produtoras de nitrito'
        ),
        'critical', jsonb_build_object(
            'value', 'Positivo',
            'description', 'Presença de bactérias gram-negativas, sugerindo ITU'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'negative', jsonb_build_object(
            'meaning', 'Ausência de nitrito na urina',
            'clinical_significance', 'Baixa probabilidade de ITU por gram-negativos. Não exclui ITU por bactérias gram-positivas (ex: Enterococcus, Staphylococcus saprophyticus) ou infecções precoces com pouca concentração bacteriana.'
        ),
        'positive', jsonb_build_object(
            'meaning', 'Presença de nitrito detectável (>0.05 mg/dL)',
            'causes', jsonb_build_array(
                'Infecção do trato urinário (ITU) por gram-negativos',
                'Cistite bacteriana aguda ou recorrente',
                'Pielonefrite em casos mais graves',
                'Colonização bacteriana assintomática (bacteriúria assintomática)'
            ),
            'clinical_significance', 'Alta especificidade (95-99%) para ITU quando positivo, mas sensibilidade moderada (45-60%). Falso-negativo ocorre em urinas diluídas, curto tempo de retenção vesical (<4h), ou infecções por gram-positivos.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'nitrito_positivo', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Aumentar ingestão hídrica para 2-3L/dia (diluição urinária)',
                'Urinar a cada 2-3h para evitar retenção prolongada de urina',
                'Higiene perineal adequada (limpeza anteroposterior)',
                'Evitar irritantes vesicais (cafeína, álcool, alimentos ácidos)',
                'Urinar após relações sexuais (reduzir risco de ITU recorrente)'
            ),
            'supplements', jsonb_build_array(
                'D-Manose 1.5-2g 2-3x/dia (prevenção de aderência de E. coli ao urotélio)',
                'Cranberry (PACs padronizados 36mg) 500mg 2x/dia (anti-aderência bacteriana)',
                'Vitamina C 500-1000mg 2-3x/dia (acidificação urinária moderada)',
                'Probióticos vaginais Lactobacillus rhamnosus GR-1 + L. reuteri RC-14 (restaurar microbiota)',
                'Berberina 500mg 2-3x/dia (antimicrobiano natural, usar por 7-10 dias)',
                'N-Acetilcisteína (NAC) 600mg 2x/dia (inibe biofilmes bacterianos)'
            ),
            'antibiotics_if_needed', 'Considerar antibioticoterapia empírica (ex: nitrofurantoína, fosfomicina) se sintomas graves. Aguardar urocultura para ajuste se possível.',
            'monitoring', 'Repetir EAS + urocultura em 3-5 dias após início do tratamento; se recorrente (>2 episódios/ano), investigar fatores anatômicos/funcionais'
        ),
        'prevencao_ITU_recorrente', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Manter hidratação adequada (urina clara)',
                'Evitar roupas íntimas sintéticas (preferir algodão)',
                'Corrigir constipação intestinal (pressão pélvica favorece ITU)',
                'Avaliar uso de espermicidas ou diafragma (alteram pH vaginal)'
            ),
            'supplements', jsonb_build_array(
                'D-Manose profilática 500mg 1x/dia à noite (uso contínuo por 3-6 meses)',
                'Cranberry 500mg 1x/dia (prevenção de longo prazo)',
                'Probióticos vaginais 1 cápsula/noite 2-3x/semana',
                'Vitamina D3 2000-4000 UI/dia (imunomodulação)',
                'Zinco quelado 15-30mg/dia (integridade de mucosas)'
            ),
            'monitoring', 'Repetir EAS mensalmente por 3 meses; após controle, trimestral por 1 ano'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Leucócitos - Sedimento',
        'Bactérias - Sedimento',
        'Urocultura',
        'pH Urinário',
        'Proteínas na Urina'
    ),
    scientific_references = jsonb_build_array(
        'Simerville JA, et al. Urinalysis: a comprehensive review. Am Fam Physician. 2005;71(6):1153-1162.',
        'Kranjčec B, et al. D-mannose powder for prophylaxis of recurrent urinary tract infections in women: a randomized clinical trial. World J Urol. 2014;32(1):79-84.',
        'Jepson RG, et al. Cranberries for preventing urinary tract infections. Cochrane Database Syst Rev. 2012;10:CD001321.'
    )
WHERE id = '1aa25d4b-a972-40db-a288-9cbe506de99e';

-- 3. Hemácias (RBC) - Sedimento (814d923f-cdfa-4388-9ba1-42b23dcd8d6d)
UPDATE score_items
SET
    clinical_context = 'A presença de hemácias no sedimento urinário (hematúria) pode ser microscópica (detectada apenas no exame) ou macroscópica (urina avermelhada visível). Indica lesão ou sangramento em qualquer ponto do trato urinário (rins, ureteres, bexiga, uretra) ou trato genital.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'max', 3,
            'unit', 'hemácias/campo',
            'description', 'Ausentes ou raras (até 3/campo é considerado normal)'
        ),
        'suboptimal', jsonb_build_object(
            'range', jsonb_build_object('min', 4, 'max', 10, 'unit', 'hemácias/campo'),
            'description', 'Hematúria microscópica leve - investigar causa'
        ),
        'critical', jsonb_build_object(
            'threshold', 10,
            'direction', 'above',
            'unit', 'hemácias/campo',
            'description', 'Hematúria significativa - investigação urgente necessária'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'normal', jsonb_build_object(
            'meaning', 'Ausentes ou raras (0-3/campo)',
            'clinical_significance', 'Trato urinário íntegro sem sangramento'
        ),
        'increased', jsonb_build_object(
            'meaning', 'Presença aumentada de hemácias (>3/campo)',
            'causes', jsonb_build_object(
                'renais', jsonb_build_array(
                    'Glomerulonefrite (hemácias dismórficas + cilindros hemáticos)',
                    'Nefrolitíase (cálculos renais)',
                    'Cistos renais ou doença policística',
                    'Pielonefrite aguda',
                    'Trauma renal',
                    'Neoplasia renal'
                ),
                'pos_renais', jsonb_build_array(
                    'Cistite hemorrágica (infecção, radiação, medicamentos)',
                    'Cálculos vesicais ou ureterais',
                    'Neoplasia de bexiga (principal causa em >50 anos)',
                    'Trauma vesical',
                    'Prostatite ou hiperplasia prostática benigna (homens)',
                    'Uretrite'
                ),
                'sistemicas', jsonb_build_array(
                    'Coagulopatias ou uso de anticoagulantes',
                    'Exercício físico intenso (hematúria de esforço)',
                    'Contaminação menstrual (mulheres)',
                    'Hipertensão arterial não controlada'
                )
            ),
            'morphology_clues', jsonb_build_object(
                'dismorficas', 'Origem glomerular (glomerulonefrite, nefrite lúpica) - hemácias com formas variadas, acantócitos',
                'isomorficas', 'Origem pós-glomerular (trato urinário baixo, trauma, cálculos, neoplasias) - hemácias com morfologia preservada'
            ),
            'clinical_significance', 'Hematúria persistente (>1 exame) em adultos >35 anos requer investigação urológica completa para exclusão de neoplasia. Em jovens com ITU, pode ser reacional e transitória.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'hematuria_leve_a_moderada', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Repetir EAS após 2 semanas para confirmar persistência',
                'Solicitar proteinúria de 24h + albumina urinária (se origem glomerular suspeitada)',
                'Ultrassom de vias urinárias (triagem inicial)',
                'Considerar cistoscopia se >40 anos ou fatores de risco para CA de bexiga',
                'Avaliar função renal (ureia, creatinina, TFG) e eletrólitos'
            ),
            'lifestyle', jsonb_build_array(
                'Hidratação adequada 2-2.5L/dia (diluir urina, prevenir cálculos)',
                'Evitar exercícios de alto impacto até definição diagnóstica',
                'Reduzir consumo de sal (<5g/dia) e proteína animal (<1.2g/kg/dia) se origem glomerular',
                'Evitar anti-inflamatórios não-esteroidais (AINE) - podem piorar hematúria'
            ),
            'supplements', jsonb_build_array(
                'Vitamina C 250-500mg/dia (não exceder, oxalato pode piorar cálculos)',
                'Magnésio glicinato 300-400mg/dia (prevenção de cálculos de oxalato)',
                'Citrato de potássio 20-30 mEq/dia (alcalinização urinária, prevenir cálculos)',
                'Ômega-3 (EPA/DHA) 1-2g/dia (anti-inflamatório renal)',
                'Cranberry 500mg/dia se ITU concomitante'
            ),
            'monitoring', 'Repetir EAS mensalmente por 3 meses; se persistir, encaminhar para nefrologista ou urologista'
        ),
        'hematuria_significativa', jsonb_build_object(
            'urgent_investigation', jsonb_build_array(
                'Avaliação urológica/nefrológica URGENTE (descartar neoplasia)',
                'Cistoscopia obrigatória se idade >40 anos',
                'TC de abdome/pelve com contraste (padrão-ouro para avaliação)',
                'Citologia urinária (triagem CA de bexiga)',
                'Coagulograma se uso de anticoagulantes ou sangramento fácil'
            ),
            'supplements', jsonb_build_array(
                'Suspender temporariamente suplementos pró-hemorrágicos (ômega-3 altas doses, vitamina E, gengibre, alho)',
                'Considerar ácido tranexâmico 500mg 3x/dia SOB SUPERVISÃO MÉDICA (antifibrinolítico)'
            ),
            'monitoring', 'Acompanhamento conforme diagnóstico definitivo. Não atrasar investigação.'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Proteinúria',
        'Cilindros Hemáticos',
        'Leucócitos - Sedimento',
        'Creatinina Sérica',
        'Ureia',
        'Coagulograma'
    ),
    scientific_references = jsonb_build_array(
        'Davis R, et al. Diagnosis, evaluation and follow-up of asymptomatic microhematuria (AMH) in adults: AUA guideline. J Urol. 2012;188(6 Suppl):2473-2481.',
        'Vivante A, et al. Persistent asymptomatic isolated microscopic hematuria in Israeli adolescents and young adults and risk for end-stage renal disease. JAMA. 2011;306(7):729-736.'
    )
WHERE id = '814d923f-cdfa-4388-9ba1-42b23dcd8d6d';

-- 4. Células Epiteliais - Sedimento (09577ef1-c3ad-461b-b2ad-59fab2c193d5)
UPDATE score_items
SET
    clinical_context = 'Células que revestem o trato urinário e genital. Sua presença no sedimento urinário é comum e geralmente benigna, mas o tipo e quantidade fornecem informações sobre qualidade da coleta, inflamação local e, raramente, processos neoplásicos.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'range', 'Raras a moderadas',
            'description', 'Descamação fisiológica normal do urotélio'
        ),
        'suboptimal', jsonb_build_object(
            'value', 'Numerosas células escamosas',
            'description', 'Sugere contaminação vaginal/uretral ou coleta inadequada'
        ),
        'abnormal', jsonb_build_object(
            'value', 'Células de transição atípicas ou numerosas',
            'description', 'Pode indicar inflamação, litíase, ou neoplasia (investigar)'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'cell_types', jsonb_build_object(
            'escamosas', jsonb_build_object(
                'origin', 'Uretra distal, vagina, pele perineal',
                'meaning', 'Mais comuns. Geralmente indicam contaminação externa (coleta inadequada) ou descamação normal. Grande quantidade (numerosas) sugere má técnica de coleta.',
                'clinical_significance', 'Sem significado patológico. Repetir coleta com técnica correta (jato médio, higiene adequada).'
            ),
            'transicionais', jsonb_build_object(
                'origin', 'Pelve renal, ureteres, bexiga, uretra proximal',
                'meaning', 'Presentes em pequena quantidade são normais (descamação fisiológica). Numerosas podem indicar inflamação, cálculos, cateterismo recente, ou neoplasia (CA de bexiga).',
                'clinical_significance', 'Se numerosas ou com atipias citológicas, considerar citologia urinária oncótica e cistoscopia (especialmente em >50 anos, tabagistas, exposição ocupacional a aminas aromáticas).'
            ),
            'renais_tubulares', jsonb_build_object(
                'origin', 'Túbulos renais',
                'meaning', 'RARAS EM URINA NORMAL. Presença indica lesão tubular renal aguda (necrose tubular aguda, nefrite intersticial, pielonefrite, isquemia renal).',
                'clinical_significance', 'Achado patológico - investigar função renal (ureia, creatinina, TFG) e causas de injúria tubular.'
            )
        ),
        'clinical_correlation', 'A interpretação depende do tipo celular, quantidade e contexto clínico. Coleta correta é fundamental para evitar falsos positivos.'
    ),
    functional_medicine_interventions = jsonb_build_object(
        'celulas_escamosas_numerosas', jsonb_build_object(
            'action', 'Repetir EAS com coleta adequada',
            'collection_technique', jsonb_build_array(
                'Higiene perineal com água e sabão neutro antes da coleta',
                'Desprezar o primeiro jato de urina (lavar a uretra distal)',
                'Coletar o jato médio em frasco estéril',
                'Evitar tocar o frasco na genitália',
                'Mulheres: afastar os lábios durante a coleta; evitar coleta durante menstruação'
            ),
            'monitoring', 'Se persistir com coleta adequada, sem significado clínico'
        ),
        'celulas_transicionais_numerosas', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Citologia urinária oncótica (triagem para CA de bexiga)',
                'Cistoscopia se idade >50 anos, tabagismo, ou hematúria concomitante',
                'Ultrassom de vias urinárias',
                'Avaliar sintomas: disúria, urgência, hematúria macroscópica'
            ),
            'lifestyle', jsonb_build_array(
                'Cessar tabagismo IMEDIATAMENTE (principal fator de risco para CA de bexiga)',
                'Hidratação 2-3L/dia (diluição de carcinógenos urinários)',
                'Evitar exposição ocupacional a aminas aromáticas, hidrocarbonetos, corantes'
            ),
            'supplements', jsonb_build_array(
                'Sulforafano (brócolis) 30-60mg/dia (indução de enzimas detoxificantes)',
                'N-Acetilcisteína (NAC) 600mg 2x/dia (antioxidante, proteção urotélio)',
                'Vitamina D3 2000-4000 UI/dia (imunomodulação, antiproliferativo)',
                'Curcumina lipossomal 500mg 2x/dia (anti-inflamatório, antitumoral)',
                'Chá verde (EGCG) 400-600mg/dia (antioxidante, proteção vesical)'
            ),
            'monitoring', 'Seguimento conforme resultados de citologia e cistoscopia'
        ),
        'celulas_renais_tubulares_presentes', jsonb_build_object(
            'urgent_evaluation', jsonb_build_array(
                'Avaliação nefrológica URGENTE',
                'Solicitar função renal completa (ureia, creatinina, TFG, eletrólitos)',
                'Investigar causas de lesão tubular: nefrotoxinas (AINE, aminoglicosídeos, contraste iodado), desidratação, sepse, rabdomiólise',
                'Proteinúria de 24h + sedimento urinário seriado'
            ),
            'lifestyle', jsonb_build_array(
                'Suspender IMEDIATAMENTE nefrotoxinas (AINE, suplementos com aristolóquicos)',
                'Otimizar hidratação (guiada por diurese e função renal)',
                'Evitar alimentos ricos em potássio se hipercalemia (banana, laranja, tomate)'
            ),
            'supplements', jsonb_build_array(
                'N-Acetilcisteína (NAC) 600-1200mg 2x/dia (nefroproteção)',
                'Ômega-3 (EPA/DHA) 1-2g/dia (anti-inflamatório renal)',
                'Astragalus 500mg 2-3x/dia (proteção tubular, evidência limitada)',
                'Coenzima Q10 100-200mg/dia (mitocondrial, proteção tubular)'
            ),
            'monitoring', 'Função renal seriada (48-72h inicialmente); acompanhamento nefrológico'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Hemácias - Sedimento',
        'Leucócitos - Sedimento',
        'Bactérias - Sedimento',
        'Creatinina Sérica',
        'Proteinúria'
    ),
    scientific_references = jsonb_build_array(
        'Ayala GE, et al. Urothelial cell proliferation in bladder washings: a marker for bladder cancer. Cancer. 2004;101(4):799-805.',
        'Fogazzi GB, et al. Urinary sediment findings in acute interstitial nephritis. Am J Kidney Dis. 2012;60(2):330-332.'
    )
WHERE id = '09577ef1-c3ad-461b-b2ad-59fab2c193d5';

-- 5. Cristais Patológicos (ebcc36fd-d285-4754-adf7-50c7b130b286)
UPDATE score_items
SET
    clinical_context = 'Cristais urinários patológicos são formações incomuns que indicam distúrbios metabólicos, uso de medicamentos específicos ou doenças sistêmicas. Diferem dos cristais fisiológicos (oxalato de cálcio, ácido úrico, fosfato triplo) que podem aparecer em urinas normais concentradas.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'value', 'Ausentes',
            'description', 'Nenhum cristal patológico identificado no sedimento'
        ),
        'pathological', jsonb_build_object(
            'value', 'Presentes',
            'description', 'Qualquer quantidade de cristais patológicos requer investigação'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'tipos_de_cristais_patologicos', jsonb_build_object(
            'cistina', jsonb_build_object(
                'appearance', 'Hexagonais, incolores, birrefringentes',
                'meaning', 'Cistinúria - doença genética autossômica recessiva com defeito no transporte tubular de aminoácidos dibásicos (cistina, ornitina, lisina, arginina)',
                'clinical_significance', 'SEMPRE PATOLÓGICO. Alto risco de nefrolitíase recorrente desde a infância. Requer manejo metabólico agressivo.',
                'treatment', 'Hidratação >3L/dia + alcalinização urinária (pH >7.5) + tiopronina ou penicilamina se refratário'
            ),
            'tirosina', jsonb_build_object(
                'appearance', 'Agulhas finas amareladas, agrupadas em feixes',
                'meaning', 'Tirosinemia hereditária (tipo I, II, III) ou doença hepática grave (cirrose, hepatite fulminante)',
                'clinical_significance', 'SEMPRE PATOLÓGICO. Indica erro inato do metabolismo ou falência hepática. Pode coexistir com cristais de leucina.',
                'investigation', 'Dosagem de tirosina sérica, succinilacetona urinária, enzimas hepáticas, alfa-fetoproteína'
            ),
            'leucina', jsonb_build_object(
                'appearance', 'Esféricos amarelados com estrias concêntricas (aspecto de "roda de carroça")',
                'meaning', 'Doença hepática grave (falência hepática aguda ou crônica descompensada) ou doenças metabólicas raras (leucinose/maple syrup urine disease)',
                'clinical_significance', 'SEMPRE PATOLÓGICO. Prognóstico reservado, especialmente se associado a cristais de tirosina.',
                'urgent_action': 'Avaliação hepatológica urgente + investigação metabólica'
            ),
            'colesterol', jsonb_build_object(
                'appearance', 'Placas retangulares transparentes com canto recortado',
                'meaning', 'Síndrome nefrótica (proteinúria maciça >3.5g/dia + hipoalbuminemia), quilúria, ou ruptura de cisto renal',
                'clinical_significance', 'Indica dano glomerular grave com extravasamento lipídico. Risco de progressão para DRC.',
                'investigation': 'Proteinúria de 24h, albumina sérica, perfil lipídico, função renal'
            ),
            'bilirrubina', jsonb_build_object(
                'appearance', 'Agulhas ou grânulos amarelos a marrom-avermelhados',
                'meaning', 'Hiperbilirrubinemia conjugada por doença hepatobiliar (hepatite, colestase, cirrose)',
                'clinical_significance': 'Indica disfunção hepática com bilirrubina direta elevada (>2mg/dL).',
                'investigation': 'Bilirrubinas totais e frações, TGO/TGP, GGT, fosfatase alcalina'
            ),
            'medicamentosos', jsonb_build_object(
                'drugs', jsonb_build_array(
                    'Sulfonamidas (cristais amarelos em "feixes de trigo")',
                    'Aciclovir/Ganciclovir (agulhas birrefringentes)',
                    'Indinavir (cristais em "leque" ou retangulares)',
                    'Metotrexato (cristais pleomórficos)',
                    'Ampicilina (agulhas longas)'
                ),
                'meaning': 'Precipitação de fármaco ou metabólitos na urina (pH e hidratação dependentes)',
                'clinical_significance': 'Risco de nefrotoxicidade, obstrução tubular aguda e insuficiência renal.',
                'management': 'Hidratação vigorosa, ajuste de dose, ou suspensão do fármaco conforme gravidade'
            )
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'cistinuria', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Hidratação MASSIVA: mínimo 3-4L/dia (meta: volume urinário >3L/dia)',
                'Distribuir ingestão hídrica ao longo do dia, incluindo 1 copo antes de dormir',
                'Reduzir ingestão de metionina (precursor de cistina): limitar carne vermelha, frango, peixe, ovos, laticínios',
                'Dieta alcalinizante: priorizar vegetais, frutas (exceto cranberry)',
                'Evitar sal (sódio aumenta excreção de cistina)'
            ),
            'supplements', jsonb_build_array(
                'Citrato de potássio 30-60 mEq/dia (alcalinizar urina, meta pH >7.5)',
                'Bicarbonato de sódio 2-4g 2-3x/dia se citrato insuficiente (monitorar sódio)',
                'Vitamina C NÃO USAR (acidifica urina)',
                'Taurina 1-2g/dia (pode reduzir cistina urinária)',
                'Acetilcisteína 600mg 3x/dia (pode quelar cistina, evidência limitada)'
            ),
            'monitoring', 'pH urinário 3x/dia (fitas reagentes caseiras), manter >7.5; EAS mensal; ultrassom renal anual'
        ),
        'cristais_hepatopatia', jsonb_build_object(
            'urgent_evaluation', jsonb_build_array(
                'Avaliação hepatológica URGENTE',
                'Hepatograma completo (TGO/TGP, bilirrubinas, GGT, FA, albumina)',
                'Coagulograma (TP/INR)',
                'Etiologia: sorologias virais (HBV, HCV), auto-anticorpos, ultrassom hepático'
            ),
            'lifestyle', jsonb_build_array(
                'Abstinência ABSOLUTA de álcool',
                'Evitar fármacos hepatotóxicos (paracetamol, AINE, estatinas)',
                'Dieta hipoproteica moderada (0.8-1g/kg/dia) se encefalopatia hepática'
            ),
            'supplements', jsonb_build_array(
                'Silimarina 200-400mg 3x/dia (hepatoproteção)',
                'N-Acetilcisteína (NAC) 600-1200mg 2x/dia (antioxidante)',
                'Vitamina E 400-800 UI/dia (esteatohepatite não-alcoólica)',
                'L-Carnitina 1-2g/dia (metabolismo lipídico hepático)',
                'Complexo B ativado (deficiência comum em hepatopatia)'
            ),
            'monitoring': 'Hepatograma semanal inicialmente; seguimento conforme gravidade'
        ),
        'sindrome_nefrotica', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Avaliação nefrológica urgente',
                'Proteinúria de 24h + albumina sérica + perfil lipídico',
                'Biópsia renal (padrão-ouro para diagnóstico etiológico)',
                'Investigar causas secundárias: diabetes, LES, amiloidose, infecções (HBV, HIV)'
            ),
            'lifestyle', jsonb_build_array(
                'Restrição de sódio <2g/dia (controlar edema)',
                'Proteína moderada 0.8-1g/kg/dia (não restringir excessivamente)',
                'Controlar hiperlipidemia com dieta (reduzir gorduras saturadas)'
            ),
            'supplements', jsonb_build_array(
                'Ômega-3 (EPA/DHA) 2-4g/dia (reduz proteinúria e dislipidemia)',
                'Vitamina D3 (níveis frequentemente baixos) - repor conforme dosagem',
                'Coenzima Q10 100-200mg/dia (proteção mitocondrial)',
                'N-Acetilcisteína (NAC) 600mg 2x/dia (antioxidante renal)'
            ),
            'monitoring': 'Proteinúria mensal, albumina sérica, função renal'
        ),
        'cristais_medicamentosos', jsonb_build_object(
            'immediate_action', jsonb_build_array(
                'Hidratação VIGOROSA intravenosa se necessário (cristalúria aguda)',
                'Ajustar dose do medicamento conforme função renal',
                'Considerar suspensão temporária ou troca de fármaco',
                'Alcalinização urinária se cristais de metotrexato ou sulfonamidas'
            ),
            'prevention', jsonb_build_array(
                'Hidratação preventiva >2L/dia durante uso de fármacos de risco',
                'Monitorar função renal antes e durante tratamento',
                'Ajustar dose para TFG (especialmente aciclovir, metotrexato)'
            )
        )
    ),
    related_biomarkers = jsonb_build_array(
        'pH Urinário',
        'Densidade Urinária',
        'Proteinúria',
        'Bilirrubinas Séricas',
        'TGO/TGP',
        'Creatinina Sérica',
        'Aminoácidos Urinários (cistinúria)'
    ),
    scientific_references = jsonb_build_array(
        'Daudon M, et al. Comprehensive morpho-constitutional analysis of urinary stones improves etiological diagnosis and therapeutic strategy of nephrolithiasis. C R Chim. 2016;19(11-12):1470-1491.',
        'Barbey F, et al. Medical treatment of cystinuria: critical reappraisal of long-term results. J Urol. 2000;163(5):1419-1423.',
        'Fogazzi GB, et al. Urinary crystals: much more than a medical curiosity. Clin Chim Acta. 2012;413(17-18):1319-1326.'
    )
WHERE id = 'ebcc36fd-d285-4754-adf7-50c7b130b286';

-- 6. Leveduras - Sedimento (1fcd3bbc-920e-4d3b-bfe3-0dd0e376f346)
UPDATE score_items
SET
    clinical_context = 'A presença de leveduras no sedimento urinário indica colonização ou infecção fúngica do trato urinário, sendo Candida spp. (especialmente C. albicans) o agente mais comum. Pode representar contaminação vaginal, candidíase vulvovaginal associada, ou candidúria verdadeira.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'value', 'Ausentes',
            'description', 'Sem evidência de fungos no trato urinário'
        ),
        'present', jsonb_build_object(
            'value', 'Raras a numerosas',
            'significance', 'Pode indicar infecção, colonização ou contaminação - contexto clínico define tratamento'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'presence', jsonb_build_object(
            'meanings', jsonb_build_array(
                'Contaminação vaginal durante coleta (mais comum em mulheres)',
                'Candidíase vulvovaginal concomitante',
                'Candidúria assintomática (colonização sem infecção)',
                'Infecção urinária fúngica verdadeira (sintomática)'
            ),
            'risk_factors', jsonb_build_array(
                'Diabetes mellitus descompensado (glicosúria favorece crescimento fúngico)',
                'Uso recente de antibióticos de amplo espectro (disbiose)',
                'Imunossupressão (HIV, quimioterapia, corticoides)',
                'Cateter vesical de demora (biofilme)',
                'Uso prolongado de inibidores de bomba de prótons (IBP)',
                'Gestação',
                'Alterações anatômicas do trato urinário (refluxo, obstrução)'
            )
        ),
        'clinical_significance', jsonb_build_object(
            'asymptomatic', 'Candidúria assintomática em pacientes imunocompetentes geralmente NÃO requer tratamento (exceto gestantes e pré-cirúrgicos urológicos). Pode resolver espontaneamente.',
            'symptomatic', 'Se disúria, urgência, febre (especialmente em diabéticos ou imunossuprimidos), tratar como candidíase urinária verdadeira.',
            'catheter_associated', 'Candidúria associada a cateter: remover ou trocar cateter é mais eficaz que antifúngicos.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'candiduria_assintomatica', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Repetir EAS com coleta adequada (técnica de jato médio)',
                'Urocultura com identificação de espécie e contagem de colônias (>10.000 UFC/mL confirma infecção)',
                'Glicemia de jejum e HbA1c (rastrear diabetes)',
                'Avaliar uso recente de antibióticos e IBPs'
            ),
            'lifestyle', jsonb_build_array(
                'Otimizar controle glicêmico se diabético (meta HbA1c <7%)',
                'Suspender IBPs se possível (aumentam pH gástrico e favorecem disbiose)',
                'Higiene perineal adequada (evitar duchas vaginais)',
                'Evitar roupas íntimas sintéticas e apertadas',
                'Reduzir açúcar e carboidratos refinados (combustível para Candida)'
            ),
            'supplements', jsonb_build_array(
                'Probióticos multi-cepas (Lactobacillus rhamnosus GR-1 + L. reuteri RC-14) 10-20 bilhões UFC/dia',
                'Saccharomyces boulardii 250-500mg 2x/dia (probiótico fúngico, competição com Candida)',
                'Berberina 500mg 2-3x/dia por 7-14 dias (antifúngico natural)',
                'Ácido caprílico 500-1000mg 2-3x/dia (antifúngico de cadeia média)',
                'Alho (Allicin) 600-900mg/dia (propriedades antifúngicas)',
                'Pau d\'arco (Tabebuia avellanedae) 500mg 2-3x/dia (antifúngico tradicional)'
            ),
            'monitoring', 'Repetir EAS + urocultura em 2-4 semanas; se persistir assintomático, não tratar'
        ),
        'candiduria_sintomatica', jsonb_build_object(
            'antifungal_treatment', jsonb_build_array(
                'Fluconazol 200mg 1x/dia por 7-14 dias (primeira linha)',
                'Alternativa: Anfotericina B (casos graves ou resistência)',
                'Se gestante: considerar tratamento tópico vaginal (clotrimazol, nistatina)'
            ),
            'lifestyle', jsonb_build_array(
                'Hidratação 2-3L/dia',
                'Dieta antifúngica: eliminar açúcares, álcool, leveduras, fungos (cogumelos), laticínios fermentados temporariamente',
                'Otimizar controle glicêmico rigorosamente',
                'Evitar irritantes vesicais (cafeína, álcool, alimentos ácidos)'
            ),
            'supplements', jsonb_build_array(
                'Combinação de tratamento médico + suplementos do protocolo assintomático',
                'Vitamina D3 2000-4000 UI/dia (imunomodulação)',
                'Zinco quelado 30mg/dia (integridade de mucosas)',
                'Vitamina C 500mg 2x/dia (acidificação urinária leve, suporte imune)',
                'N-Acetilcisteína (NAC) 600mg 2x/dia (inibe biofilmes fúngicos)'
            ),
            'monitoring', 'Urocultura de controle 1-2 semanas após fim do tratamento; se recorrente, investigar fatores predisponentes'
        ),
        'candiduria_recorrente', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Investigar diabetes, HIV, imunossupressão',
                'Avaliar uso crônico de antibióticos, corticoides, IBPs',
                'Considerar cistoscopia se alterações anatômicas suspeitas',
                'Pesquisar Candida resistente (C. glabrata, C. krusei) - solicitar teste de sensibilidade'
            ),
            'long_term_management', jsonb_build_array(
                'Fluconazol profilático 100-200mg 1x/semana por 6 meses (casos selecionados)',
                'Probióticos contínuos (Lactobacillus + Saccharomyces boulardii)',
                'Dieta low-carb mantida (reduzir substratos para fungos)',
                'Suplementação preventiva: berberina + ácido caprílico em ciclos mensais (2 semanas ON, 2 semanas OFF)'
            ),
            'monitoring': 'Urocultura trimestral por 1 ano'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Glicose Urinária',
        'pH Urinário',
        'Leucócitos - Sedimento',
        'Urocultura',
        'Glicemia de Jejum',
        'HbA1c'
    ),
    scientific_references = jsonb_build_array(
        'Kauffman CA, et al. Clinical practice guidelines for the management of candidiasis: 2009 update by the Infectious Diseases Society of America. Clin Infect Dis. 2009;48(5):503-535.',
        'Achkar JM, Fries BC. Candida infections of the genitourinary tract. Clin Microbiol Rev. 2010;23(2):253-273.',
        'Gupta V, et al. Cranberry products inhibit adherence of p-fimbriated Escherichia coli to primary cultured bladder and vaginal epithelial cells. J Urol. 2007;177(6):2357-2360.'
    )
WHERE id = '1fcd3bbc-920e-4d3b-bfe3-0dd0e376f346';

-- 7. SHBG - Homens (fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291)
UPDATE score_items
SET
    clinical_context = 'Sex Hormone Binding Globulin (SHBG) é uma glicoproteína hepática que se liga a hormônios sexuais (testosterona e estradiol) no sangue. Aproximadamente 60% da testosterona circulante está ligada ao SHBG (forma inativa), 38% à albumina (fracamente ligada), e apenas 2% está livre (forma biologicamente ativa). SHBG regula a biodisponibilidade hormonal e é modulado por fatores metabólicos.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 20,
            'max', 50,
            'unit', 'nmol/L',
            'description', 'Equilíbrio ideal para biodisponibilidade de testosterona livre adequada e proteção cardiometabólica'
        ),
        'suboptimal', jsonb_build_object(
            'ranges', jsonb_build_array(
                jsonb_build_object('min', 10, 'max', 19, 'label', 'Baixo (risco metabólico)'),
                jsonb_build_object('min', 51, 'max', 70, 'label', 'Elevado (biodisponibilidade reduzida)')
            )
        ),
        'critical', jsonb_build_object(
            'low', jsonb_build_object('threshold', 10, 'description', 'Associado a síndrome metabólica, resistência insulínica, risco cardiovascular'),
            'high', jsonb_build_object('threshold', 70, 'description', 'Testosterona livre muito baixa, sintomas hipogonadismo funcional')
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'SHBG reduzido (hiperandrogenismo relativo)',
            'causes', jsonb_build_array(
                'Resistência insulínica e síndrome metabólica (principal causa)',
                'Obesidade visceral (adiposidade suprime SHBG)',
                'Diabetes tipo 2',
                'Hipotireoidismo (T3/T4 baixos reduzem síntese hepática de SHBG)',
                'Uso de andrógenos exógenos ou SARMs',
                'Esteatose hepática não-alcoólica (NAFLD)',
                'Síndrome do ovário policístico (PCOS) em contextos específicos',
                'Uso de corticoides ou hormônio de crescimento'
            ),
            'clinical_significance', 'SHBG baixo aumenta testosterona livre circulante, mas paradoxalmente está associado a PIOR saúde metabólica (resistência insulínica, dislipidemia, risco cardiovascular aumentado). Não é um marcador de "virilidade", mas sim de disfunção metabólica. Valores <15 nmol/L estão fortemente associados a síndrome metabólica e diabetes tipo 2.',
            'symptoms', jsonb_build_array(
                'Ganho de peso abdominal',
                'Fadiga pós-prandial',
                'Acantose nigricans (escurecimento axilar/pescoço)',
                'Dislipidemia (TG alto, HDL baixo)',
                'Hipertensão arterial',
                'Acne ou oleosidade cutânea aumentada (em alguns casos)'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'SHBG na faixa ideal (20-50 nmol/L)',
            'clinical_significance', 'Sensibilidade insulínica preservada, metabolismo hormonal equilibrado, proteção cardiometabólica, biodisponibilidade de testosterona adequada.'
        ),
        'high', jsonb_build_object(
            'meaning', 'SHBG elevado',
            'causes', jsonb_build_array(
                'Hipertireoidismo (T3/T4 elevados estimulam síntese de SHBG)',
                'Deficiência androgênica (hipogonadismo primário ou secundário)',
                'Envelhecimento (SHBG aumenta 1-2% ao ano após 40 anos)',
                'Hepatopatia crônica (cirrose, hepatite C)',
                'Uso de estrogênios (TRH, contraceptivos orais)',
                'Desnutrição ou restrição calórica severa',
                'Anorexia nervosa',
                'HIV/AIDS',
                'Genética (polimorfismos no gene SHBG)'
            ),
            'clinical_significance', 'SHBG elevado reduz testosterona livre e biologicamente ativa, resultando em sintomas de hipogonadismo funcional MESMO com testosterona total normal. Homens podem apresentar libido reduzida, disfunção erétil, fadiga, perda de massa muscular, apesar de testosterona total "normal".',
            'symptoms', jsonb_build_array(
                'Libido reduzida',
                'Disfunção erétil',
                'Fadiga crônica',
                'Perda de massa muscular (sarcopenia)',
                'Ganho de gordura corporal',
                'Depressão ou ansiedade',
                'Perda de vitalidade geral'
            )
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'shbg_baixo', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Dieta low-carb ou cetogênica (reduzir insulina) - <100g carbs/dia',
                'Jejum intermitente 16:8 (melhora sensibilidade insulínica)',
                'Exercício HIIT 3x/semana + treinamento de força 3-4x/semana',
                'Perda de peso (meta: reduzir circunferência abdominal <94cm)',
                'Sono 7-9h/noite (privação piora resistência insulínica)',
                'Reduzir exposição a xenoestrogênios (plásticos BPA, parabenos)'
            ),
            'supplements', jsonb_build_array(
                'Berberina 500mg 3x/dia antes das refeições (sensibilizador insulínico)',
                'Inositol (Myo-inositol) 2-4g/dia (melhora sinalização insulínica)',
                'Magnésio glicinato 300-400mg/noite (cofator insulínico)',
                'Cromo picolinato 200-400mcg/dia (metabolismo glicose)',
                'Ômega-3 (EPA/DHA) 2-3g/dia (anti-inflamatório, melhora SHBG)',
                'Vitamina D3 4000-5000 UI/dia (meta: 50-70 ng/mL)',
                'Resveratrol 250-500mg/dia (ativa SIRT1, melhora metabolismo)',
                'NAC 600mg 2x/dia (antioxidante hepático, melhora NAFLD)'
            ),
            'monitoring', 'Repetir SHBG + testosterona total e livre + glicemia jejum + insulina jejum + HOMA-IR em 3 meses'
        ),
        'shbg_elevado', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Avaliar função tireoidiana (TSH, T4 livre, T3 livre)',
                'Dosar testosterona LIVRE ou calcular índice de andrógenos livres (FAI = [Testosterona total / SHBG] x 100)',
                'Investigar hipogonadismo (LH, FSH)',
                'Hepatograma completo se suspeita de hepatopatia',
                'Avaliar estado nutricional (albumina, pré-albumina, IMC)'
            ),
            'lifestyle', jsonb_build_array(
                'Otimizar ingestão calórica (evitar restrição severa <1500 kcal/dia)',
                'Dieta com carboidratos moderados (150-200g/dia) - carbs reduzem SHBG',
                'Treinamento de força progressivo 4-5x/semana (estimula produção de testosterona)',
                'Evitar overtraining (excesso de exercício eleva SHBG)',
                'Sono 8-9h/noite (otimizar eixo HPA e produção hormonal)',
                'Reduzir estresse crônico (cortisol eleva SHBG) - meditação, yoga'
            ),
            'supplements', jsonb_build_array(
                'Boro 6-10mg/dia (reduz SHBG, aumenta testosterona livre)',
                'Magnésio glicinato 400-600mg/dia (reduz SHBG)',
                'Zinco quelado 30-50mg/dia (suporte produção testosterona)',
                'Vitamina D3 5000 UI/dia (otimizar níveis >50 ng/mL)',
                'Ashwagandha (KSM-66) 600mg/dia (reduz cortisol, melhora testosterona)',
                'Tribulus terrestris 750-1500mg/dia (evidência mista, pode aumentar LH)',
                'DHEA 25-50mg/dia (precursor de testosterona, usar sob supervisão)',
                'Complexo B ativado (suporte metilação e produção hormonal)'
            ),
            'monitoring', 'Repetir SHBG + testosterona total e livre + função tireoidiana em 2-3 meses'
        ),
        'hipogonadismo_funcional', jsonb_build_object(
            'when', 'SHBG >60 nmol/L + Testosterona total normal (>300 ng/dL) + Testosterona livre baixa (<50 pg/mL)',
            'treatment', jsonb_build_array(
                'Foco em REDUZIR SHBG (protocolo acima)',
                'Considerar TRT (Terapia de Reposição de Testosterona) se intervenções não-farmacológicas falharem após 6 meses',
                'TRT em gel transdérmico ou injeções (cypionato/enantato) aumenta testosterona livre apesar de SHBG alto'
            ),
            'monitoring': 'Seguimento com endocrinologista/urologista especializado em saúde hormonal masculina'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Testosterona Total',
        'Testosterona Livre',
        'Índice de Andrógenos Livres (FAI)',
        'LH e FSH',
        'Estradiol',
        'TSH e T3/T4 Livre',
        'Glicemia de Jejum e Insulina',
        'HOMA-IR',
        'Hemoglobina Glicada (HbA1c)'
    ),
    scientific_references = jsonb_build_array(
        'Ding EL, et al. Sex hormone-binding globulin and risk of type 2 diabetes in women and men. N Engl J Med. 2009;361(12):1152-1163.',
        'Jia G, et al. Low serum SHBG is associated with metabolic syndrome in men: results from NHANES III. Endocrine. 2015;49(2):482-489.',
        'Caldwell JR, et al. SHBG as a biomarker of cardiometabolic risk in men. Clin Endocrinol. 2019;90(1):42-51.'
    )
WHERE id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291';

-- 8. SHBG - Mulheres (c21ccec2-66b2-49e3-911a-8d0944eda087)
UPDATE score_items
SET
    clinical_context = 'Sex Hormone Binding Globulin (SHBG) em mulheres regula a biodisponibilidade de estrogênios e andrógenos. SHBG baixo está associado a hiperandrogenismo, síndrome metabólica e SOP (Síndrome do Ovário Policístico). SHBG elevado pode indicar hipoestrogenismo, hipertireoidismo ou uso de contraceptivos orais.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 40,
            'max', 120,
            'unit', 'nmol/L',
            'description', 'Equilíbrio hormonal adequado com proteção metabólica e ovulatória'
        ),
        'suboptimal', jsonb_build_object(
            'ranges', jsonb_build_array(
                jsonb_build_object('min', 20, 'max', 39, 'label', 'Baixo (risco de hiperandrogenismo e SOP)'),
                jsonb_build_object('min', 121, 'max', 180, 'label', 'Elevado (redução de andrógenos livres)')
            )
        ),
        'critical', jsonb_build_object(
            'low', jsonb_build_object('threshold', 20, 'description', 'Fortemente associado a SOP, resistência insulínica, infertilidade'),
            'high', jsonb_build_object('threshold', 180, 'description', 'Andrógenos e estrogênios livres muito baixos')
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'SHBG reduzido (hiperandrogenismo relativo)',
            'causes', jsonb_build_array(
                'Síndrome do Ovário Policístico (SOP) - principal causa',
                'Resistência insulínica e hiperinsulinemia',
                'Obesidade (adiposidade suprime SHBG)',
                'Diabetes tipo 2',
                'Hipotireoidismo',
                'Esteatose hepática não-alcoólica (NAFLD)',
                'Uso de andrógenos exógenos'
            ),
            'clinical_significance', 'SHBG <40 nmol/L aumenta testosterona livre e estradiol livre, resultando em sinais de hiperandrogenismo (hirsutismo, acne, alopecia). É um marcador robusto de resistência insulínica e risco metabólico em mulheres. Valores <20 nmol/L são diagnósticos de SOP em 70-80% dos casos.',
            'symptoms', jsonb_build_array(
                'Irregularidade menstrual ou amenorreia',
                'Hirsutismo (pelos faciais, abdomen, costas)',
                'Acne persistente',
                'Alopecia androgênica (padrão masculino)',
                'Ganho de peso abdominal',
                'Acantose nigricans',
                'Infertilidade anovulatória',
                'Sintomas de resistência insulínica'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'SHBG na faixa ideal (40-120 nmol/L)',
            'clinical_significance', 'Ciclos menstruais regulares, ovulação normal, equilíbrio androgênico, sensibilidade insulínica preservada, fertilidade adequada.'
        ),
        'high', jsonb_build_object(
            'meaning', 'SHBG elevado',
            'causes', jsonb_build_array(
                'Uso de contraceptivos orais (estrogênios sintéticos estimulam SHBG)',
                'Hipertireoidismo',
                'Deficiência androgênica (hipogonadismo)',
                'Anorexia nervosa ou desnutrição',
                'Hepatopatia crônica (cirrose)',
                'Menopausa (menor supressão por andrógenos)',
                'Genética'
            ),
            'clinical_significance', 'SHBG elevado reduz andrógenos livres (testosterona livre), podendo causar libido reduzida, fadiga, perda de massa muscular, mesmo com níveis totais normais. Mulheres em contraceptivos orais frequentemente têm SHBG muito elevado (>150 nmol/L).',
            'symptoms', jsonb_build_array(
                'Libido reduzida',
                'Fadiga crônica',
                'Secura vaginal',
                'Perda de massa muscular',
                'Dificuldade de ganho muscular',
                'Mudanças de humor'
            )
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'shbg_baixo_sop', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Dieta low-carb (50-100g carbs/dia) ou cetogênica (reduzir insulina)',
                'Jejum intermitente 16:8 (melhora sensibilidade insulínica)',
                'Exercício combinado: HIIT 3x/semana + treinamento de força 3x/semana',
                'Perda de peso 5-10% (melhora significativa em SHBG e ovulação)',
                'Sono 7-9h/noite',
                'Reduzir exposição a disruptores endócrinos (BPA, ftalatos, parabenos)'
            ),
            'supplements', jsonb_build_array(
                'Inositol (Myo-inositol 2g + D-chiro-inositol 50mg - razão 40:1) 2x/dia (melhora sensibilidade insulínica e ovulação)',
                'Berberina 500mg 3x/dia (sensibilizador insulínico, reduz testosterona)',
                'N-Acetilcisteína (NAC) 1200-1800mg/dia (melhora SOP, fertilidade)',
                'Magnésio glicinato 300-400mg/noite',
                'Cromo picolinato 200-400mcg/dia',
                'Ômega-3 (EPA/DHA) 2-3g/dia',
                'Vitamina D3 4000-5000 UI/dia (meta: 50-70 ng/mL)',
                'Canela (Cassia) 1-3g/dia (melhora sensibilidade insulínica)',
                'Saw Palmetto 160mg 2x/dia (anti-androgênico)',
                'Spearmint tea 2 xícaras/dia (reduz testosterona livre)'
            ),
            'monitoring', 'Repetir SHBG + testosterona total e livre + glicemia + insulina + HOMA-IR + ciclos menstruais em 3 meses'
        ),
        'shbg_elevado_contraceptivos', jsonb_build_object(
            'evaluation', jsonb_build_array(
                'Considerar trocar contraceptivo oral por métodos não-hormonais (DIU cobre) ou progestágenos isolados',
                'Dosar testosterona livre (frequentemente muito baixa)',
                'Avaliar sintomas de hipoandrogenismo (libido baixa, fadiga, secura)'
            ),
            'supplements_after_discontinuation', jsonb_build_array(
                'Aguardar 3-6 meses para normalização de SHBG após suspensão de CO',
                'Vitaminas do complexo B (suporte metilação e destoxificação hormonal)',
                'DIM (Diindolilmetano) 200mg/dia (metabolismo de estrogênios)',
                'Cálcio-D-Glucarato 500mg/dia (eliminação hepática de estrogênios)',
                'Probióticos (suporte ao estroboloma intestinal)',
                'Cardo Mariano 200mg 2x/dia (suporte hepático)'
            ),
            'monitoring': 'Repetir SHBG + testosterona livre em 3-6 meses após mudança contraceptiva'
        ),
        'shbg_elevado_hipoandrogenismo', jsonb_build_object(
            'lifestyle', jsonb_build_array(
                'Dieta com carboidratos moderados (150-200g/dia)',
                'Evitar restrição calórica severa',
                'Treinamento de força 4-5x/semana (estimula andrógenos)',
                'Reduzir overtraining',
                'Otimizar sono 8h/noite',
                'Técnicas de redução de estresse (cortisol eleva SHBG)'
            ),
            'supplements', jsonb_build_array(
                'Boro 3-6mg/dia (reduz SHBG, aumenta testosterona livre)',
                'Magnésio glicinato 400mg/dia',
                'Zinco quelado 15-30mg/dia',
                'Vitamina D3 4000 UI/dia',
                'Ashwagandha (KSM-66) 300-600mg/dia (reduz cortisol)',
                'DHEA 10-25mg/dia (precursor de andrógenos, usar sob supervisão)',
                'Maca peruana 1.5-3g/dia (adaptógeno, melhora libido)'
            ),
            'monitoring': 'Repetir SHBG + testosterona livre + função tireoidiana em 2-3 meses'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Testosterona Total e Livre',
        'Índice de Andrógenos Livres (FAI)',
        'Estradiol',
        'LH e FSH',
        'TSH e T3/T4 Livre',
        'Glicemia de Jejum e Insulina',
        'HOMA-IR',
        'Hemoglobina Glicada (HbA1c)',
        'Progesterona (dia 21 do ciclo)',
        'Prolactina'
    ),
    scientific_references = jsonb_build_array(
        'Simó R, et al. Novel insights in SHBG regulation and clinical implications. Trends Endocrinol Metab. 2015;26(7):376-383.',
        'Legro RS, et al. Diagnosis and treatment of polycystic ovary syndrome: an Endocrine Society clinical practice guideline. J Clin Endocrinol Metab. 2013;98(12):4565-4592.',
        'Unfer V, et al. Inositol(s) in PCOS: rationale, clinical evidence, and mechanisms of action. Gynecol Endocrinol. 2019;35(sup1):52-58.'
    )
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';

-- 9-12. DHEA-S - Homens (40-49, 50-59, 60-69, 70+ anos)
-- IDs: 66256f14-1775-4501-b85e-a4ffd71ccc7a, bc896c78-c530-4dd6-ad03-6da22ffd7fea,
--      9c5ef523-046b-4647-abd4-719937d54eb6, 447746dd-949b-449b-bd0e-c2f226330a10

UPDATE score_items
SET
    clinical_context = 'Dehidroepiandrosterona sulfato (DHEA-S) é o hormônio esteroide mais abundante no organismo, produzido principalmente pela zona reticular do córtex adrenal. É um precursor de andrógenos (testosterona) e estrogênios. Declina progressivamente com a idade (50% aos 40 anos, 20% aos 70 anos). Reflete a reserva funcional adrenal e está associado a longevidade, função cognitiva, composição corporal e vitalidade.',
    functional_ranges = jsonb_build_object(
        'by_age', jsonb_build_object(
            '40-49_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 200, 'max', 400, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 100, 'max', 199, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 100, 'description', 'Insuficiência adrenal relativa')
            ),
            '50-59_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 150, 'max', 350, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 80, 'max', 149, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 80, 'description', 'Depleção adrenal significativa')
            ),
            '60-69_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 100, 'max', 300, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 60, 'max', 99, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 60, 'description', 'Envelhecimento adrenal acelerado')
            ),
            '70+_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 80, 'max', 250, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 40, 'max', 79, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 40, 'description', 'Adrenopausa avançada')
            )
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'DHEA-S abaixo do ideal para a idade',
            'causes', jsonb_build_array(
                'Envelhecimento fisiológico (adrenopausa)',
                'Estresse crônico (exaustão do eixo HPA)',
                'Insuficiência adrenal primária (Addison) ou secundária',
                'Uso crônico de corticoides (supressão adrenal)',
                'Hipopituitarismo (deficiência de ACTH)',
                'Doenças crônicas (inflamatórias, autoimunes)',
                'Deficiências nutricionais (vitaminas B, C, magnésio)',
                'Overtraining (exercício excessivo sem recuperação)',
                'Privação crônica de sono'
            ),
            'clinical_significance', 'DHEA-S baixo está associado a fadiga crônica, perda de massa muscular, ganho de gordura abdominal, redução de libido, declínio cognitivo, depressão, imunodeficiência, e aumento de mortalidade cardiovascular. Marcador de "envelhecimento biológico acelerado".',
            'symptoms', jsonb_build_array(
                'Fadiga persistente (pior pela manhã)',
                'Perda de motivação e vitalidade',
                'Redução de massa muscular e força',
                'Ganho de gordura visceral',
                'Libido reduzida',
                'Disfunção erétil',
                'Declínio cognitivo e memória',
                'Depressão ou ansiedade',
                'Imunidade comprometida (infecções recorrentes)',
                'Pele seca e envelhecimento cutâneo acelerado'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'DHEA-S na faixa funcional ideal para a idade',
            'clinical_significance', 'Reserva adrenal adequada, eixo HPA funcional, proteção cardiometabólica, função cognitiva preservada, composição corporal favorável, longevidade saudável.'
        ),
        'high', jsonb_build_object(
            'meaning', 'DHEA-S elevado',
            'causes', jsonb_build_array(
                'Tumores adrenais produtores de DHEA (adenoma, carcinoma)',
                'Hiperplasia adrenal congênita (deficiência de 21-hidroxilase)',
                'SOP (Síndrome do Ovário Policístico) - mais comum em mulheres',
                'Uso de suplementação exógena de DHEA'
            ),
            'clinical_significance': 'Homens raramente têm DHEA-S elevado fora de tumores ou suplementação. Investigar massa adrenal com TC ou RNM se muito elevado (>700 µg/dL).'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'dhea_baixo', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Dosar cortisol salivar circadiano (4 pontos: manhã, meio-dia, tarde, noite)',
                'Avaliar eixo HPA: cortisol sérico AM + ACTH',
                'Avaliar função tireoidiana (TSH, T3, T4)',
                'Rastrear deficiências nutricionais (vitamina D, B12, magnésio)',
                'Investigar estressores crônicos e qualidade de sono'
            ),
            'lifestyle', jsonb_build_array(
                'Otimização do sono: 7-9h/noite, dormir antes das 23h, quarto escuro',
                'Gestão de estresse: meditação 20min/dia, respiração diafragmática, yoga',
                'Exercício moderado: evitar overtraining, priorizar recuperação',
                'Exposição solar matinal 15-30min (sincronização circadiana)',
                'Reduzir cafeína (máximo 1-2 cafés antes das 14h)',
                'Evitar álcool (suprime produção de DHEA)',
                'Dieta anti-inflamatória rica em antioxidantes'
            ),
            'supplements', jsonb_build_array(
                'DHEA 25-50mg/dia pela manhã (CONSIDERAR REPOSIÇÃO - SUPERVISÃO MÉDICA)',
                'Ashwagandha (KSM-66) 600mg/dia (adaptógeno, suporta adrenais)',
                'Rhodiola rosea 200-400mg/manhã (adaptógeno, melhora fadiga)',
                'Fosfatidilserina 300mg/noite (reduz cortisol noturno)',
                'Vitamina C 1000-2000mg/dia (suporte adrenal)',
                'Complexo B ativado (P-5-P 50mg + metilfolato 800mcg + metilcobalamina 1000mcg)',
                'Magnésio glicinato 400-600mg/noite (estresse e sono)',
                'Pantetina (B5 ativa) 600mg/dia (precursor síntese esteroides)',
                'Pregnenolona 50-100mg/manhã (precursor de DHEA e outros hormônios)',
                'Zinco quelado 30-50mg/dia (síntese de testosterona)',
                'Vitamina D3 5000 UI/dia (otimizar >50 ng/mL)'
            ),
            'dhea_supplementation_protocol', jsonb_build_object(
                'dosing', 'Iniciar 25mg/manhã por 4 semanas, depois aumentar para 50mg/manhã se bem tolerado',
                'monitoring', 'Dosar DHEA-S após 8-12 semanas. Meta: terço superior da faixa ideal para a idade.',
                'cautions', 'Monitorar PSA (pode estimular crescimento prostático), evitar se história de CA de próstata. Pode converter em estrogênios (monitorar estradiol). Suspender se acne ou oleosidade excessiva.'
            ),
            'monitoring', 'Repetir DHEA-S + testosterona + estradiol em 3 meses'
        ),
        'dhea_elevado', jsonb_build_object(
            'investigation', jsonb_build_array(
                'TC ou RNM de adrenais (investigar adenoma ou carcinoma)',
                'Dosar 17-OH-progesterona (hiperplasia adrenal congênita)',
                'Avaliar histórico de suplementação de DHEA (descontinuar)',
                'Dosar testosterona e estradiol (conversão periférica)'
            ),
            'monitoring': 'Seguimento com endocrinologista; encaminhamento cirúrgico se massa adrenal'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Cortisol Sérico e Salivar',
        'Testosterona Total e Livre',
        'Estradiol',
        'ACTH',
        'Pregnenolona',
        'TSH e T3/T4',
        'PSA (homens)'
    ),
    scientific_references = jsonb_build_array(
        'Orentreich N, et al. Age changes and sex differences in serum DHEA sulfate concentrations throughout adulthood. J Clin Endocrinol Metab. 1984;59(3):551-555.',
        'Rutkowski K, et al. DHEA: hype or hope. Drugs. 2014;74(11):1195-1207.',
        'Traish AM, et al. Dehydroepiandrosterone (DHEA) - a precursor steroid or an active hormone in human physiology. J Sex Med. 2011;8(11):2960-2982.'
    )
WHERE id IN (
    '66256f14-1775-4501-b85e-a4ffd71ccc7a',
    'bc896c78-c530-4dd6-ad03-6da22ffd7fea',
    '9c5ef523-046b-4647-abd4-719937d54eb6',
    '447746dd-949b-449b-bd0e-c2f226330a10'
);

-- 13-14. DHEA-S - Mulheres (60-69 anos e 70+ anos)
-- IDs: 22952700-63e3-422d-b4b2-179c0785a20e, e622f011-e92d-459e-9f4a-af435f26ea74

UPDATE score_items
SET
    clinical_context = 'DHEA-S em mulheres é a principal fonte de andrógenos após a menopausa. O declínio de DHEA-S é mais acentuado em mulheres que em homens, com níveis 20% menores aos 70 anos comparado aos 20 anos. Relacionado a libido, energia, massa muscular, densidade óssea e proteção cardiovascular na pós-menopausa.',
    functional_ranges = jsonb_build_object(
        'by_age', jsonb_build_object(
            '60-69_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 35, 'max', 150, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 20, 'max', 34, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 20, 'description', 'Depleção androgênica significativa')
            ),
            '70+_anos', jsonb_build_object(
                'optimal', jsonb_build_object('min', 25, 'max', 120, 'unit', 'µg/dL'),
                'suboptimal', jsonb_build_object('min', 15, 'max', 24, 'unit', 'µg/dL'),
                'critical_low', jsonb_build_object('threshold', 15, 'description', 'Adrenopausa completa')
            )
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'DHEA-S abaixo do ideal para a idade',
            'causes', jsonb_build_array(
                'Adrenopausa (envelhecimento adrenal)',
                'Estresse crônico e exaustão do eixo HPA',
                'Insuficiência adrenal',
                'Uso crônico de corticoides',
                'Doenças crônicas',
                'Menopausa (perda adicional de estrogênios ovarianos amplifica impacto do baixo DHEA)',
                'Privação de sono crônica'
            ),
            'clinical_significance', 'DHEA-S baixo em mulheres pós-menopausa está associado a perda de libido, fadiga, sarcopenia, osteoporose, aumento de risco cardiovascular, declínio cognitivo e depressão. Contribui para síndrome de deficiência androgênica feminina.',
            'symptoms', jsonb_build_array(
                'Fadiga crônica e perda de vitalidade',
                'Perda de libido (falta de desejo sexual)',
                'Secura vaginal (perda de andrógenos locais)',
                'Perda de massa muscular (sarcopenia)',
                'Ganho de gordura abdominal',
                'Pele seca, fina e envelhecimento acelerado',
                'Declínio cognitivo e memória',
                'Depressão e baixa motivação',
                'Osteoporose (perda de densidade óssea)',
                'Imunidade reduzida'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'DHEA-S na faixa funcional ideal para a idade',
            'clinical_significance', 'Reserva androgênica adequada, libido preservada, composição corporal favorável, proteção óssea e cardiovascular, cognição saudável.'
        ),
        'high', jsonb_build_object(
            'meaning', 'DHEA-S elevado',
            'causes', jsonb_build_array(
                'SOP (Síndrome do Ovário Policístico) em mulheres pré-menopausa',
                'Hiperplasia adrenal congênita',
                'Tumores adrenais',
                'Suplementação exógena'
            ),
            'clinical_significance': 'DHEA-S elevado em mulheres pós-menopausa é raro. Investigar tumor adrenal se >350 µg/dL sem suplementação.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'dhea_baixo_pos_menopausa', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Dosar cortisol salivar circadiano (avaliar eixo HPA)',
                'Avaliar testosterona livre (frequentemente baixa com DHEA baixo)',
                'Dosar estradiol (se TRH, níveis artificialmente elevados)',
                'Avaliar função tireoidiana (TSH, T3, T4)',
                'Densitometria óssea (rastreio osteoporose)',
                'Vitamina D e magnésio sérico'
            ),
            'lifestyle', jsonb_build_array(
                'Otimização do sono: 7-9h/noite, higiene do sono rigorosa',
                'Gestão de estresse: meditação, yoga, terapia cognitivo-comportamental',
                'Exercício resistido 3-4x/semana (preservar massa muscular e óssea)',
                'Exposição solar matinal 20-30min',
                'Dieta anti-inflamatória rica em antioxidantes e proteínas (1-1.2g/kg/dia)',
                'Reduzir cafeína e álcool'
            ),
            'supplements', jsonb_build_array(
                'DHEA 10-25mg/dia pela manhã (REPOSIÇÃO - SUPERVISÃO MÉDICA)',
                'Ashwagandha (KSM-66) 300-600mg/dia (adaptógeno)',
                'Rhodiola rosea 200-400mg/manhã (energia e fadiga)',
                'Maca peruana 1.5-3g/dia (libido e vitalidade)',
                'Vitamina C 1000mg/dia (suporte adrenal)',
                'Complexo B ativado',
                'Magnésio glicinato 400mg/noite',
                'Pantetina (B5) 600mg/dia',
                'Pregnenolona 25-50mg/manhã (precursor de DHEA)',
                'Vitamina D3 4000-5000 UI/dia',
                'Ômega-3 (EPA/DHA) 2g/dia',
                'Cálcio + Vitamina K2 (proteção óssea)'
            ),
            'dhea_supplementation_women', jsonb_build_object(
                'dosing', 'Iniciar 10mg/manhã por 4 semanas, aumentar para 25mg/manhã se bem tolerado',
                'monitoring', 'Dosar DHEA-S + testosterona livre após 8-12 semanas. Meta: terço superior da faixa ideal.',
                'cautions', 'Monitorar sinais de androgenização (acne, hirsutismo leve). Pode melhorar libido, energia e composição corporal significativamente. Suspender se sinais de excesso androgênico.'
            ),
            'monitoring', 'Repetir DHEA-S + testosterona livre + estradiol em 3 meses'
        ),
        'dhea_elevado', jsonb_build_object(
            'investigation', jsonb_build_array(
                'TC ou RNM de adrenais',
                'Dosar 17-OH-progesterona',
                'Descontinuar suplementação se uso',
                'Encaminhamento endocrinológico'
            )
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Cortisol Sérico e Salivar',
        'Testosterona Total e Livre',
        'Estradiol',
        'FSH e LH',
        'Pregnenolona',
        'TSH e T3/T4',
        'Densitometria Óssea'
    ),
    scientific_references = jsonb_build_array(
        'Labrie F, et al. DHEA and the intracrine formation of androgens and estrogens in peripheral target tissues: its role during aging. Steroids. 1998;63(5-6):322-328.',
        'Davis SR, et al. Circulating androgen levels and self-reported sexual function in women. JAMA. 2005;294(1):91-96.',
        'Wierman ME, et al. Androgen therapy in women: a reappraisal. J Clin Endocrinol Metab. 2014;99(10):3489-3510.'
    )
WHERE id IN (
    '22952700-63e3-422d-b4b2-179c0785a20e',
    'e622f011-e92d-459e-9f4a-af435f26ea74'
);

-- =============================================================================
-- CHECKPOINT: 14 items concluídos (Urobilinogênio até DHEA-S Mulheres 70+)
-- Próximos: TSH, T3 Livre, T3 Reverso, INR, Testosterona/SHBG mulheres, etc.
-- =============================================================================

-- 15. TSH (34af6e5c-3847-46d8-874e-a7364c014877)
UPDATE score_items
SET
    clinical_context = 'Hormônio estimulante da tireoide (TSH) é produzido pela hipófise anterior e regula a produção de hormônios tireoidianos (T3 e T4). É o exame de triagem mais sensível para disfunção tireoidiana. Valores flutuam com ritmo circadiano (pico entre 2-4h da manhã), idade, estresse e medicamentos.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 0.5,
            'max', 2.0,
            'unit', 'mUI/L',
            'description', 'Faixa funcional ideal para metabolismo, energia e saúde cardiovascular'
        ),
        'suboptimal', jsonb_build_object(
            'ranges', jsonb_build_array(
                jsonb_build_object('min', 2.1, 'max', 4.5, 'label', 'Hipotireoidismo subclínico'),
                jsonb_build_object('min', 0.1, 'max', 0.49, 'label', 'Hipertireoidismo subclínico')
            )
        ),
        'critical', jsonb_build_object(
            'high', jsonb_build_object('threshold', 4.5, 'description', 'Hipotireoidismo clínico ou avançado'),
            'low', jsonb_build_object('threshold', 0.1, 'description', 'Hipertireoidismo clínico ou supressão por T4 exógeno')
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'high', jsonb_build_object(
            'meaning', 'TSH elevado (hipotireoidismo primário ou subclínico)',
            'causes', jsonb_build_array(
                'Tireoidite de Hashimoto (autoimune - causa mais comum)',
                'Deficiência de iodo (rara em países com sal iodado)',
                'Tireoidectomia parcial ou total',
                'Radioterapia cervical ou ablação com iodo radioativo',
                'Medicamentos (lítio, amiodarona, interferon)',
                'Resistência periférica aos hormônios tireoidianos (rara)',
                'Hipotireoidismo congênito (detectado em triagem neonatal)'
            ),
            'clinical_significance', 'TSH >2.5 mUI/L está associado a fadiga, ganho de peso, constipação, pele seca, intolerância ao frio, dislipidemia (LDL alto), risco cardiovascular aumentado e infertilidade. TSH >4.5 mUI/L requer reposição hormonal na maioria dos casos.',
            'symptoms', jsonb_build_array(
                'Fadiga persistente',
                'Ganho de peso inexplicado',
                'Constipação intestinal',
                'Pele seca e cabelos quebradiços',
                'Intolerância ao frio',
                'Depressão e lentidão mental',
                'Bradicardia (FC <60 bpm)',
                'Edema facial (mixedema)',
                'Ciclos menstruais irregulares ou amenorreia',
                'Infertilidade'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'TSH na faixa funcional ideal (0.5-2.0 mUI/L)',
            'clinical_significance', 'Função tireoidiana ótima com metabolismo eficiente, energia adequada, termorregulação normal, saúde cardiovascular e reprodutiva preservadas.'
        ),
        'low', jsonb_build_object(
            'meaning', 'TSH baixo ou suprimido (hipertireoidismo primário ou subclínico)',
            'causes', jsonb_build_array(
                'Doença de Graves (autoimune - anticorpos TRAb)',
                'Bócio multinodular tóxico',
                'Adenoma tóxico (nódulo hiperfuncionante)',
                'Tireoidite subaguda (fase tireotóxica transitória)',
                'Uso excessivo de levotiroxina (T4 exógeno)',
                'Uso de amiodarona (indução de tireotoxicose)',
                'Gestação 1º trimestre (hCG estimula receptor de TSH)',
                'Tumor hipofisário secretor de TSH (rarissimo)'
            ),
            'clinical_significance', 'TSH <0.1 mUI/L está associado a risco de fibrilação atrial, osteoporose, ansiedade, insônia, tremores, perda de peso e taquicardia. TSH suprimido cronicamente requer investigação e tratamento.',
            'symptoms', jsonb_build_array(
                'Taquicardia e palpitações',
                'Perda de peso apesar de apetite normal/aumentado',
                'Ansiedade, irritabilidade, insônia',
                'Tremores finos de extremidades',
                'Intolerância ao calor e sudorese excessiva',
                'Diarreia ou evacuações frequentes',
                'Fraqueza muscular proximal',
                'Exoftalmia (olhos saltados - Doença de Graves)',
                'Oligomenorreia ou amenorreia'
            )
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'tsh_elevado', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Dosar T4 livre e T3 livre (avaliar severidade)',
                'Anticorpos anti-TPO e anti-Tg (rastreio Hashimoto)',
                'Ultrassom de tireoide (avaliar ecogenicidade, nódulos, volume)',
                'Vitamina D, ferritina, selênio (deficiências comuns em hipotireoidismo)',
                'Perfil lipídico (hipotireoidismo eleva LDL)',
                'Considerar teste de iodo urinário 24h se deficiência suspeitada'
            ),
            'levothyroxine_replacement', jsonb_build_object(
                'indication', 'TSH >10 mUI/L: sempre tratar. TSH 4.5-10 mUI/L: tratar se sintomático, anticorpos positivos, dislipidemia, gestação ou planejando engravidar.',
                'dosing', 'Levotiroxina (T4) 1.6 mcg/kg/dia, tomar em jejum 30-60min antes do café. Ajustar dose a cada 6-8 semanas conforme TSH. Meta: TSH 0.5-2.0 mUI/L.',
                'monitoring', 'Dosar TSH + T4 livre 6-8 semanas após início ou ajuste de dose. Após estabilização, monitorar TSH anualmente.'
            ),
            'lifestyle', jsonb_build_array(
                'Dieta anti-inflamatória (reduzir autoimunidade se Hashimoto)',
                'Evitar glúten (considerar teste de doença celíaca - comum em Hashimoto)',
                'Reduzir goitrogênios crus (soja, brócolis, couve) se consumo excessivo',
                'Otimizar sono 7-9h/noite',
                'Exercício moderado (overtraining pode piorar hipotireoidismo)',
                'Reduzir estresse (cortisol alto interfere com conversão T4→T3)'
            ),
            'supplements', jsonb_build_array(
                'Selênio 200mcg/dia (reduz anticorpos anti-TPO, melhora conversão T4→T3)',
                'Iodo 150-300mcg/dia (APENAS se deficiência confirmada - excesso piora Hashimoto)',
                'Zinco quelado 30mg/dia (conversão T4→T3)',
                'Vitamina D3 4000-5000 UI/dia (meta: >50 ng/mL)',
                'Ferro (se ferritina <50 ng/mL) - separar 4h da levotiroxina',
                'Magnésio glicinato 400mg/noite',
                'L-Tirosina 500mg 2x/dia (precursor de T3/T4)',
                'Ashwagandha (KSM-66) 600mg/dia (adaptógeno, melhora T3/T4)',
                'Vitaminas do complexo B (suporte metabolismo)',
                'Probióticos (80% da conversão T4→T3 ocorre no intestino)'
            ),
            'monitoring', 'TSH + T4 livre + T3 livre + anticorpos a cada 3-6 meses até estabilização'
        ),
        'tsh_baixo', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Dosar T4 livre e T3 livre (confirmar tireotoxicose)',
                'Anticorpos TRAb (Doença de Graves)',
                'Anticorpos anti-TPO e anti-Tg',
                'Ultrassom de tireoide + Doppler (avaliar vascularização)',
                'Cintilografia de tireoide com captação de iodo (diferenciar causas)',
                'ECG (rastreio fibrilação atrial)',
                'Densitometria óssea se TSH suprimido crônico (risco osteoporose)'
            ),
            'treatment', jsonb_build_object(
                'graves', 'Metimazol ou propiltiouracil (ATDs) + beta-bloqueador (propranolol) para controle sintomático. Considerar ablação com iodo radioativo ou tireoidectomia se refratário.',
                'nodulo_toxico', 'Ablação com iodo radioativo ou cirurgia.',
                'levothyroxine_excess', 'Reduzir dose de levotiroxina gradualmente até TSH normalizar.'
            ),
            'lifestyle', jsonb_build_array(
                'Evitar cafeína e estimulantes (pioram taquicardia)',
                'Reduzir iodo dietético (evitar algas, sal iodado em excesso)',
                'Exercício leve a moderado (evitar overtraining)',
                'Técnicas de relaxamento (reduzir ansiedade)',
                'Sono 8-9h/noite'
            ),
            'supplements', jsonb_build_array(
                'L-Carnitina 2-4g/dia (inibe entrada de T3/T4 nas células, melhora sintomas)',
                'Magnésio glicinato 400-600mg/dia (reduz ansiedade e tremores)',
                'Vitamina D3 (corrigir se deficiente)',
                'Ômega-3 (EPA/DHA) 2-3g/dia (anti-inflamatório)',
                'Coenzima Q10 200mg/dia (suporte mitocondrial)',
                'Taurina 1-2g/dia (reduz taquicardia)',
                'EVITAR iodo suplementar'
            ),
            'monitoring': 'TSH + T3/T4 livre mensalmente durante tratamento; ECG e densitometria anualmente'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'T4 Livre',
        'T3 Livre',
        'T3 Reverso',
        'Anticorpos Anti-TPO',
        'Anticorpos Anti-Tireoglobulina',
        'TRAb (Anticorpos Anti-Receptor de TSH)',
        'Ultrassom de Tireoide',
        'Iodo Urinário 24h'
    ),
    scientific_references = jsonb_build_array(
        'Garber JR, et al. Clinical practice guidelines for hypothyroidism in adults: cosponsored by AACE and ATA. Endocr Pract. 2012;18(6):988-1028.',
        'Ross DS, et al. 2016 American Thyroid Association guidelines for diagnosis and management of hyperthyroidism. Thyroid. 2016;26(10):1343-1421.',
        'Gärtner R, et al. Selenium supplementation in patients with autoimmune thyroiditis decreases thyroid peroxidase antibodies concentrations. J Clin Endocrinol Metab. 2002;87(4):1687-1691.'
    )
WHERE id = '34af6e5c-3847-46d8-874e-a7364c014877';

-- 16. T3 Livre (d164eacf-a0d7-48f2-899d-3f0d57ec7cc3)
UPDATE score_items
SET
    clinical_context = 'Triiodotironina livre (T3 livre) é a forma biologicamente ativa do hormônio tireoidiano. 80% do T3 circulante provém da conversão periférica de T4→T3 (principalmente no fígado, rins e músculos), e apenas 20% é secretado diretamente pela tireoide. T3 livre regula metabolismo celular, termogênese, síntese proteica e função cardiovascular.',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 3.2,
            'max', 4.0,
            'unit', 'pg/mL',
            'description', 'Terço superior da faixa para metabolismo e energia ideais'
        ),
        'suboptimal', jsonb_build_object(
            'ranges', jsonb_build_array(
                jsonb_build_object('min', 2.3, 'max', 3.1, 'label', 'Conversão T4→T3 subótima'),
                jsonb_build_object('min', 4.1, 'max', 5.0, 'label', 'Hipertireoidismo subclínico')
            )
        ),
        'critical', jsonb_build_object(
            'low', jsonb_build_object('threshold', 2.3, 'description', 'Hipotireoidismo ou síndrome do eutireoidiano doente'),
            'high', jsonb_build_object('threshold', 5.0, 'description', 'Tireotoxicose clínica')
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'T3 livre baixo',
            'causes', jsonb_build_array(
                'Conversão prejudicada T4→T3 (deficiência de selênio, zinco, ferro)',
                'Hipotireoidismo primário (Hashimoto)',
                'Síndrome do eutireoidiano doente (doenças graves, jejum prolongado, estresse severo)',
                'Uso de beta-bloqueadores, corticoides, amiodarona (inibem conversão)',
                'Resistência à insulina e diabetes tipo 2',
                'Insuficiência renal ou hepática crônica',
                'Deficiência calórica severa (<1200 kcal/dia)',
                'Hipotireoidismo central (hipófise/hipotálamo)'
            ),
            'clinical_significance', 'T3 livre baixo com TSH normal/baixo sugere má conversão de T4→T3 (tipo 2 deiodinase disfuncional). Pacientes podem ter sintomas de hipotireoidismo apesar de T4 e TSH normais. É comum em resistência insulínica, inflamação crônica e deficiências nutricionais.',
            'symptoms', jsonb_build_array(
                'Fadiga persistente apesar de levotiroxina adequada',
                'Dificuldade de perda de peso',
                'Intolerância ao frio',
                'Brainfog e dificuldade de concentração',
                'Constipação',
                'Pele seca',
                'Perda de cabelo'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'T3 livre no terço superior da faixa (3.2-4.0 pg/mL)',
            'clinical_significance', 'Conversão T4→T3 eficiente, metabolismo celular ótimo, energia adequada, função cognitiva preservada.'
        ),
        'high', jsonb_build_object(
            'meaning', 'T3 livre elevado',
            'causes', jsonb_build_array(
                'Doença de Graves',
                'Bócio multinodular tóxico',
                'Adenoma tóxico',
                'Tireoidite subaguda (fase tireotóxica)',
                'Tireotoxicose factícia (uso de hormônio exógeno)',
                'Tireotoxicose por T3 (rara)'
            ),
            'clinical_significance', 'T3 alto é mais sintomático que T4 alto (causa mais taquicardia, ansiedade, perda de peso). Requer controle rápido.',
            'symptoms': 'Taquicardia, palpitações, perda de peso, ansiedade, tremores, insônia, diarreia.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        't3_baixo_com_t4_normal', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Dosar T3 Reverso (rT3) - se alto, indica bloqueio de conversão',
                'Avaliar selênio, zinco, ferro, ferritina (cofatores de deiodinase)',
                'Dosar cortisol (estresse crônico inibe conversão)',
                'Avaliar glicemia, insulina, HOMA-IR (resistência insulínica inibe conversão)',
                'Hepatograma (fígado é principal local de conversão)',
                'Avaliar uso de medicamentos que inibem conversão'
            ),
            'treatment_options', jsonb_build_object(
                'optimize_conversion', 'Corrigir deficiências nutricionais + reduzir inflamação',
                'add_t3', 'Considerar adicionar T3 (liotironina) 5-10mcg 2x/dia se sintomático apesar de T4 otimizado',
                'ndt', 'Tireoide dessecada (NDT - Nature-Throid, Armour) contém T3+T4 naturalmente'
            ),
            'lifestyle', jsonb_build_array(
                'Dieta anti-inflamatória, low-carb se resistência insulínica',
                'Evitar restrição calórica severa (mínimo 1500 kcal/dia)',
                'Reduzir estresse crônico',
                'Sono 7-9h/noite',
                'Exercício moderado (overtraining reduz T3)'
            ),
            'supplements', jsonb_build_array(
                'Selênio 200mcg/dia (cofator essencial de deiodinase tipo 1 e 2)',
                'Zinco quelado 30mg/dia (cofator de deiodinase)',
                'Ferro (se ferritina <50 ng/mL)',
                'Vitamina A 5000-10000 UI/dia (receptor nuclear de T3)',
                'Vitamina D3 4000-5000 UI/dia',
                'Ômega-3 (EPA/DHA) 2g/dia (anti-inflamatório)',
                'Ashwagandha (KSM-66) 600mg/dia (melhora conversão T4→T3)',
                'Guggul (Commiphora mukul) 500mg 2-3x/dia (estimula conversão)',
                'L-Tirosina 500mg 2x/dia'
            ),
            'monitoring', 'Repetir T3 livre + T3 reverso + TSH em 8-12 semanas'
        ),
        't3_elevado', jsonb_build_object(
            'treatment', 'Seguir protocolo de hipertireoidismo (ver item TSH baixo). Usar beta-bloqueadores + antitiroidiano.',
            'monitoring': 'T3 livre + TSH mensalmente durante tratamento'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'TSH',
        'T4 Livre',
        'T3 Reverso',
        'Selênio Sérico',
        'Zinco Sérico',
        'Ferritina',
        'Cortisol',
        'Glicemia e Insulina'
    ),
    scientific_references = jsonb_build_array(
        'Hoermann R, et al. Homeostatic control of the thyroid-pituitary axis: perspectives for diagnosis and treatment. Front Endocrinol. 2015;6:177.',
        'Bianco AC, et al. Deiodinases: implications of the local control of thyroid hormone action. J Clin Invest. 2006;116(10):2571-2579.',
        'Sharma AK, et al. Efficacy and safety of Ashwagandha root extract in subclinical hypothyroid patients: a double-blind, randomized placebo-controlled trial. J Altern Complement Med. 2018;24(3):243-248.'
    )
WHERE id = 'd164eacf-a0d7-48f2-899d-3f0d57ec7cc3';

-- 17. T3 Reverso (4159c2e3-97e2-4ffc-922d-4513fdbc82aa)
UPDATE score_items
SET
    clinical_context = 'T3 Reverso (rT3) é um metabólito INATIVO do T4, formado pela remoção de um átomo de iodo em posição diferente da conversão T4→T3 ativo. É produzido em situações de estresse, inflamação, jejum, doença grave ou toxicidade. rT3 compete com T3 pelos receptores celulares, bloqueando a ação do T3 ativo. Razão T3/rT3 baixa indica "hipotireoidismo funcional".',
    functional_ranges = jsonb_build_object(
        'optimal', jsonb_build_object(
            'min', 8,
            'max', 15,
            'unit', 'ng/dL',
            'description', 'Metabolismo hormonal equilibrado sem bloqueio de conversão'
        ),
        'suboptimal', jsonb_build_object(
            'range', jsonb_build_object('min', 15.1, 'max', 20, 'unit', 'ng/dL'),
            'description', 'Conversão desviada para rT3, reduzindo T3 ativo disponível'
        ),
        'critical', jsonb_build_object(
            'threshold', 20,
            'direction', 'above',
            'description', 'Síndrome do rT3 alto - bloqueio significativo de receptores de T3'
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'high', jsonb_build_object(
            'meaning', 'T3 Reverso elevado (síndrome do rT3 alto)',
            'causes', jsonb_build_array(
                'Estresse crônico e ativação prolongada do eixo HPA',
                'Inflamação sistêmica (citocinas IL-6, TNF-alfa inibem deiodinase tipo 1)',
                'Restrição calórica severa ou jejum prolongado (mecanismo de conservação energética)',
                'Doenças graves (sepse, ICC, DRC, câncer) - síndrome do eutireoidiano doente',
                'Deficiências nutricionais (selênio, zinco, ferro)',
                'Toxicidade por metais pesados (mercúrio, chumbo)',
                'Resistência à insulina e diabetes tipo 2',
                'Uso de beta-bloqueadores, amiodarona, propranolol',
                'Dose excessiva de levotiroxina (T4) sem conversão adequada'
            ),
            'clinical_significance', 'rT3 >20 ng/dL + T3 livre baixo/normal indica "hipotireoidismo funcional". Pacientes têm sintomas de hipotireoidismo (fadiga, ganho de peso, brainfog) MESMO com TSH e T4 normais. A razão T3/rT3 <10 (usando unidades pg/mL para T3 e ng/dL para rT3) confirma bloqueio metabólico.',
            'symptoms', jsonb_build_array(
                'Fadiga severa não-responsiva a levotiroxina',
                'Ganho de peso ou dificuldade extrema de emagrecer',
                'Brainfog e lentidão cognitiva',
                'Intolerância ao frio e calor',
                'Depressão resistente',
                'Constipação',
                'Queda de cabelo',
                'Pele seca'
            )
        ),
        'optimal', jsonb_build_object(
            'meaning', 'rT3 na faixa ideal (8-15 ng/dL)',
            'clinical_significance', 'Conversão T4→T3 eficiente, sem desvio excessivo para rT3. Razão T3/rT3 adequada (>10), permitindo sinalização celular normal.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'rt3_elevado', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Calcular razão T3 livre/rT3 (converter unidades: T3 em pg/mL, rT3 em ng/dL. Meta: razão >10)',
                'Avaliar cortisol salivar circadiano (estresse crônico)',
                'Marcadores inflamatórios (PCR-us, VHS, ferritina)',
                'Selênio, zinco, ferro, ferritina',
                'Glicemia, insulina, HOMA-IR',
                'Hepatograma completo',
                'Avaliar exposição a toxinas (metais pesados, pesticidas)'
            ),
            'treatment_strategy', jsonb_build_object(
                'step_1_remove_causes', 'Identificar e tratar causa raiz: reduzir estresse, tratar inflamação, corrigir deficiências, remover toxinas',
                'step_2_t4_dose', 'Se em levotiroxina, considerar REDUZIR dose de T4 (excesso piora rT3)',
                'step_3_add_t3', 'Adicionar T3 (liotironina) 5-10mcg 2-3x/dia - T3 exógeno bypassa conversão e satura receptores bloqueados por rT3',
                'step_4_consider_ndt', 'Trocar para tireoide dessecada (NDT) com razão T4:T3 de 4:1 (mais fisiológica)'
            ),
            'lifestyle', jsonb_build_array(
                'Reduzir estresse: meditação, yoga, terapia, redução de carga de trabalho',
                'Sono rigoroso: 8-9h/noite, dormir antes das 22-23h',
                'Dieta anti-inflamatória: eliminar glúten, laticínios, açúcar, alimentos processados',
                'Dieta low-carb moderada (100-150g carbs/dia) se resistência insulínica',
                'Evitar restrição calórica extrema (manter >1500 kcal/dia)',
                'Exercício LEVE a moderado (overtraining piora rT3)',
                'Sauna infravermelha (aumenta clearance de rT3)',
                'Evitar exposição a toxinas ambientais'
            ),
            'supplements', jsonb_build_array(
                'Selênio 200-400mcg/dia (aumenta clearance de rT3)',
                'Zinco quelado 30-50mg/dia',
                'Ferro (se ferritina <50 ng/mL)',
                'Vitamina A 10000 UI/dia',
                'Vitamina D3 5000 UI/dia',
                'Ômega-3 (EPA/DHA) 2-3g/dia (anti-inflamatório)',
                'Ashwagandha (KSM-66) 600mg/dia (reduz cortisol)',
                'Rhodiola rosea 200-400mg/manhã (adaptógeno)',
                'Curcumina lipossomal 500mg 2x/dia (anti-inflamatório)',
                'N-Acetilcisteína (NAC) 600mg 2x/dia (detox)',
                'Ácido alfa-lipóico 300-600mg/dia (quelação metais, antioxidante)',
                'Probióticos multi-cepas (gut-thyroid axis)',
                'Glutationa lipossomal 500mg/dia (detox, antioxidante)'
            ),
            'monitoring', 'Repetir T3 livre + rT3 + TSH a cada 6-8 semanas durante tratamento. Meta: rT3 <15 ng/dL + razão T3/rT3 >10'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'T3 Livre',
        'T4 Livre',
        'TSH',
        'Razão T3/rT3',
        'Cortisol Salivar Circadiano',
        'PCR Ultrassensível',
        'Ferritina',
        'Selênio Sérico',
        'Zinco Sérico'
    ),
    scientific_references = jsonb_build_array(
        'Peeters RP, et al. Reduced activation and increased inactivation of thyroid hormone in tissues of critically ill patients. J Clin Endocrinol Metab. 2003;88(7):3202-3211.',
        'Chopra IJ, et al. Pathways of metabolism of thyroid hormones. Recent Prog Horm Res. 1978;34:521-567.',
        'Wiersinga WM. Paradigm shifts in thyroid hormone replacement therapies for hypothyroidism. Nat Rev Endocrinol. 2014;10(3):164-174.'
    )
WHERE id = '4159c2e3-97e2-4ffc-922d-4513fdbc82aa';

-- 18. Tempo de Protrombina (INR) (459b1285-86d6-408f-9735-029dd00e67b6)
UPDATE score_items
SET
    clinical_context = 'INR (International Normalized Ratio) padroniza o Tempo de Protrombina (TP), avaliando a via extrínseca da coagulação (fatores VII, X, V, II, fibrinogênio). Reflete função hepática (síntese de fatores) e efeito anticoagulante de varfarina. Valores elevados indicam risco hemorrágico; baixos, risco trombótico.',
    functional_ranges = jsonb_build_object(
        'normal_without_anticoagulation', jsonb_build_object(
            'min', 0.9,
            'max', 1.1,
            'unit', 'INR',
            'description', 'Coagulação normal sem uso de anticoagulantes'
        ),
        'therapeutic_anticoagulation', jsonb_build_object(
            'conditions', jsonb_build_object(
                'fibrilacao_atrial', jsonb_build_object('target', 2.0, 'range', '2.0-3.0', 'description', 'Prevenção de AVC'),
                'trombose_venosa', jsonb_build_object('target', 2.5, 'range', '2.0-3.0', 'description', 'TVP, embolia pulmonar'),
                'protese_valvar_mecanica', jsonb_build_object('target', 2.5, 'range', '2.5-3.5', 'description', 'Válvulas mecânicas de alto risco')
            )
        ),
        'critical', jsonb_build_object(
            'low', jsonb_build_object('threshold', 0.8, 'description', 'Hipercoagulabilidade - risco trombótico'),
            'high_without_ac', jsonb_build_object('threshold', 1.5, 'description', 'Coagulopatia - investigar disfunção hepática ou deficiência de vitamina K'),
            'high_with_ac', jsonb_build_object('threshold', 4.0, 'description', 'Risco hemorrágico grave - reduzir anticoagulação')
        )
    ),
    biomarker_interpretation = jsonb_build_object(
        'low', jsonb_build_object(
            'meaning', 'INR <0.9 (hipercoagulabilidade)',
            'causes', jsonb_build_array(
                'Uso de vitamina K suplementar ou dieta rica em vitamina K',
                'Resistência à varfarina (polimorfismo VKORC1)',
                'Trombofilia hereditária (avaliar fator V Leiden, protrombina G20210A)',
                'Desidratação severa (hemoconcentração)',
                'Fase inicial de CID (coagulação intravascular disseminada)'
            ),
            'clinical_significance', 'Aumenta risco de trombose venosa (TVP, embolia pulmonar) ou arterial (IAM, AVC). Em pacientes anticoagulados, INR baixo indica falha terapêutica e risco de eventos tromboembólicos.'
        ),
        'normal', jsonb_build_object(
            'meaning', 'INR 0.9-1.1 (sem anticoagulação) ou dentro da faixa terapêutica (com anticoagulação)',
            'clinical_significance', 'Hemostasia equilibrada.'
        ),
        'high_without_anticoagulation', jsonb_build_object(
            'meaning', 'INR >1.5 sem uso de varfarina',
            'causes', jsonb_build_array(
                'Insuficiência hepática (cirrose, hepatite fulminante) - síntese reduzida de fatores',
                'Deficiência de vitamina K (má absorção, antibióticos prolongados, nutrição parenteral)',
                'CID (Coagulação Intravascular Disseminada) - consumo de fatores',
                'Uso de antibióticos (inibem flora intestinal produtora de vit K)',
                'Má absorção intestinal (doença celíaca, Crohn, colestase)',
                'Uso de medicamentos que interferem (antibióticos, amiodarona, anabolizantes)'
            ),
            'clinical_significance', 'Risco hemorrágico aumentado. INR >2.0 sem anticoagulação requer investigação urgente de hepatopatia ou deficiência de vit K. Risco de sangramento espontâneo (gastrointestinal, cerebral) se INR >3.5.',
            'symptoms': 'Equimoses fáceis, sangramento gengival, epistaxe, menorragia, hematúria, melena.'
        ),
        'high_with_anticoagulation', jsonb_build_object(
            'meaning', 'INR acima da faixa terapêutica (>3.5 ou >4.0)',
            'causes', jsonb_build_array(
                'Overdose de varfarina (intencional ou erro de dose)',
                'Interações medicamentosas (antibióticos, antifúngicos, amiodarona, omeprazol potencializam varfarina)',
                'Redução de ingestão de vitamina K (mudança dietética)',
                'Disfunção hepática aguda (reduz metabolismo de varfarina)',
                'Diarreia ou má absorção (altera metabolismo)',
                'Polimorfismo CYP2C9 (metabolismo lento de varfarina)'
            ),
            'clinical_significance': 'Risco hemorrágico GRAVE. INR >5.0 requer reversão parcial com vitamina K oral/IV. INR >10 ou sangramento ativo requer reversão urgente com concentrado de complexo protrombínico (CCP) + vitamina K IV.'
        )
    ),
    functional_medicine_interventions = jsonb_build_object(
        'inr_baixo_em_anticoagulacao', jsonb_build_object(
            'action', jsonb_build_array(
                'Aumentar dose de varfarina conforme protocolo médico',
                'Avaliar adesão ao tratamento',
                'Investigar ingestão excessiva de vitamina K (vegetais verdes, suplementos)',
                'Avaliar interações medicamentosas (rifampicina, carbamazepina, barbitúricos reduzem efeito)',
                'Considerar teste genético VKORC1 e CYP2C9 (resistência à varfarina)'
            ),
            'lifestyle', 'Manter ingestão de vitamina K CONSTANTE (não eliminar, mas evitar variações grandes). Evitar suplementos de vit K.',
            'monitoring': 'INR semanal até estabilização na faixa terapêutica, depois mensal'
        ),
        'inr_alto_sem_anticoagulacao', jsonb_build_object(
            'investigation', jsonb_build_array(
                'Hepatograma completo (TGO/TGP, bilirrubinas, albumina, TP, PTT)',
                'Pesquisar deficiência de vitamina K: tempo de tromboplastina parcial (PTT), dosagem de vit K',
                'Avaliar má absorção: exames de doença celíaca, Crohn, função pancreática',
                'Coagulograma completo (fibrinogênio, D-dímero, plaquetas)',
                'Histórico de antibióticos (alteram flora intestinal)'
            ),
            'treatment', jsonb_build_object(
                'vitamin_k_deficiency', 'Vitamina K1 (filoquinona) 5-10mg VO/dia por 3 dias, depois repetir INR. Se grave (INR >3 com sangramento), vitamina K 10mg IV lento.',
                'liver_disease', 'Tratar hepatopatia subjacente. Considerar plasma fresco congelado (PFC) se sangramento ativo.'
            ),
            'supplements', jsonb_build_array(
                'Vitamina K2 (MK-7) 100-200mcg/dia (APENAS se deficiência confirmada)',
                'Probióticos (restaurar flora intestinal produtora de vit K)',
                'Enzimas digestivas (se má absorção)'
            ),
            'monitoring': 'Repetir INR em 48-72h após vitamina K'
        ),
        'inr_alto_em_anticoagulacao', jsonb_build_object(
            'management_by_inr', jsonb_build_object(
                'inr_3_5_to_5_0', 'Reduzir dose de varfarina ou omitir 1 dose. Repetir INR em 24-48h.',
                'inr_5_0_to_9_0', 'Omitir 1-2 doses de varfarina + vitamina K1 oral 2.5-5mg. Repetir INR em 24h.',
                'inr_>9_0_sem_sangramento', 'Omitir varfarina + vitamina K1 oral 5-10mg. Repetir INR em 12-24h.',
                'sangramento_ativo', 'EMERGÊNCIA: suspender varfarina + concentrado de complexo protrombínico (CCP) + vitamina K1 10mg IV lento. Hospitalizar.'
            ),
            'identify_cause', jsonb_build_array(
                'Revisar todos os medicamentos (interações)',
                'Avaliar mudança dietética recente',
                'Investigar disfunção hepática ou renal aguda',
                'Verificar adesão (paciente tomou dose dupla por engano?)'
            ),
            'monitoring': 'INR diário até retorno à faixa terapêutica, depois semanal por 1 mês'
        ),
        'alternative_to_warfarin', jsonb_build_object(
            'consider', 'Se INR instável recorrente (variações frequentes), considerar trocar varfarina por DOAC (apixaban, rivaroxaban, dabigatran) - não requerem monitoramento de INR.',
            'note': 'DOACs não são adequados para válvulas mecânicas (contraindicado).'
        )
    ),
    related_biomarkers = jsonb_build_array(
        'Tempo de Tromboplastina Parcial (PTT)',
        'Fibrinogênio',
        'Plaquetas',
        'D-Dímero',
        'TGO/TGP',
        'Bilirrubinas',
        'Albumina',
        'Vitamina K Sérica'
    ),
    scientific_references = jsonb_build_array(
        'Ansell J, et al. Pharmacology and management of the vitamin K antagonists: American College of Chest Physicians Evidence-Based Clinical Practice Guidelines (8th Edition). Chest. 2008;133(6 Suppl):160S-198S.',
        'Holbrook A, et al. Evidence-based management of anticoagulant therapy: Antithrombotic Therapy and Prevention of Thrombosis, 9th ed. Chest. 2012;141(2 Suppl):e152S-e184S.'
    )
WHERE id = '459b1285-86d6-408f-9735-029dd00e67b6';

-- Continuação no próximo bloco devido ao limite de caracteres...
