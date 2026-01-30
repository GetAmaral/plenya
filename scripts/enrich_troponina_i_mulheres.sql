-- Enrichment for Score Item: Troponina I Ultrassensível - Mulheres
-- ID: 247a9b99-ec59-4dd7-b7a3-b5482b1dd553
-- Generated: 2026-01-29

BEGIN;

-- Update clinical content fields
UPDATE score_items
SET
  clinical_relevance = 'A troponina I ultrassensível (hs-TnI) é o biomarcador mais sensível e específico para detecção de lesão miocárdica, sendo fundamental no diagnóstico de síndrome coronariana aguda (SCA) em mulheres. Estudos demonstram que mulheres apresentam concentrações basais de troponina significativamente menores que homens, com percentil 99 específico por sexo de aproximadamente 16 ng/L (versus 34 ng/L em homens). A utilização de limiares específicos para o sexo feminino aumenta em 42% a identificação de lesão miocárdica em mulheres, corrigindo o histórico subdiagnóstico de infarto agudo do miocárdio nesta população. As diretrizes da Sociedade Europeia de Cardiologia (ESC 2023) recomendam algoritmos acelerados 0h/1h com hs-TnI para rule-in e rule-out de infarto, permitindo decisões clínicas mais rápidas e seguras. A especificidade por sexo é particularmente relevante considerando diferenças fisiopatológicas: mulheres têm menor massa ventricular esquerda, menor densidade de receptores beta-adrenérgicos e apresentam mais frequentemente SCA com artérias coronárias não obstrutivas (MINOCA). Valores abaixo do limite de detecção (<2 ng/L) na apresentação têm excelente valor preditivo negativo, enquanto elevações acima do percentil 99 específico (>16 ng/L) com delta >50% em 1 hora indicam lesão miocárdica aguda. A interpretação deve sempre considerar contexto clínico, cinética temporal (elevação/queda), e excluir causas não isquêmicas de elevação como sepse, insuficiência renal crônica, embolia pulmonar e miocardite. A implementação de cut-offs específicos por sexo representa avanço significativo na medicina de precisão cardiovascular, embora estudos demonstrem que mulheres ainda recebem metade das intervenções coronárias comparadas aos homens mesmo após diagnóstico adequado, indicando necessidade de abordagem mais equitativa no tratamento.',

  patient_explanation = 'A troponina I ultrassensível é uma proteína presente nas células do músculo cardíaco que, quando aumentada no sangue, indica lesão ou sofrimento do coração. Mulheres naturalmente têm níveis mais baixos desta proteína comparadas aos homens - o valor normal máximo para mulheres é cerca de 16 ng/L, metade do valor masculino (34 ng/L). Esta diferença é importante porque usar o mesmo valor de referência para ambos os sexos pode resultar em diagnóstico tardio de infarto em mulheres. O exame é fundamental quando há suspeita de infarto agudo do miocárdio (ataque cardíaco), especialmente se você apresenta sintomas como dor no peito, falta de ar, náuseas ou desconforto torácico atípico. Em mulheres, os sintomas de infarto podem ser diferentes dos homens, incluindo cansaço extremo, dor nas costas ou mandíbula, sem dor clássica no peito. O protocolo atual recomenda duas coletas: uma na chegada à emergência e outra após 1 hora. Valores muito baixos (<2 ng/L) na primeira coleta praticamente descartam infarto. Valores acima de 16 ng/L ou aumento significativo (mais de 50%) entre as duas coletas indicam lesão cardíaca aguda e necessidade de tratamento urgente. Outras condições podem elevar a troponina sem ser infarto, como insuficiência renal, infecções graves, embolia pulmonar ou inflamação do coração, por isso o médico avaliará junto com seu quadro clínico completo, eletrocardiograma e histórico. A detecção precoce e precisa de infarto em mulheres salva vidas, pois permite tratamento imediato com medicamentos anticoagulantes, cateterismo cardíaco ou cirurgia quando necessário.',

  conduct = 'Diante de resultado alterado de troponina I ultrassensível em mulheres, seguir protocolo ESC 0h/1h: (1) RULE-OUT: Valores <2 ng/L na apresentação (0h) associados a baixo risco clínico e ECG normal permitem exclusão segura de SCA com alta precoce e seguimento ambulatorial em 30 dias. Se 0h entre 2-16 ng/L, coletar 1h - se permanecer <16 ng/L sem delta significativo e baixo risco, considerar rule-out. (2) RULE-IN: Valores ≥16 ng/L (percentil 99 específico) na apresentação OU delta absoluto ≥5 ng/L entre 0h-1h indicam lesão miocárdica aguda - proceder: monitorização contínua em unidade coronariana, AAS 300mg dose ataque seguido 100mg/dia, clopidogrel 300-600mg ou ticagrelor 180mg dose ataque, anticoagulação com fondaparinux 2,5mg SC ou enoxaparina 1mg/kg 12/12h, estratificação de risco com escores GRACE/HEART, discussão com hemodinâmica para estratégia invasiva precoce (<24h se alto risco ou <72h se risco intermediário), considerar betabloqueador, IECA/BRA e estatina de alta intensidade. (3) OBSERVE-ZONE: Valores entre 5-16 ng/L sem delta diagnóstico requerem terceira coleta em 3h, troponina seriada até 6-12h, investigação complementar com ecocardiograma, TC coronárias ou teste provocativo, e exclusão de causas não-SCA (disfunção renal com TFG <30, sepse, TEP com D-dímero e angioTC se indicado, miocardite se quadro viral recente). Atenção especial: em mulheres avaliar MINOCA (25% das SCA femininas) - considerar RNM cardíaca se cineangiocoronariografia sem lesões obstrutivas. (4) INTERPRETAÇÃO CONTEXTO-DEPENDENTE: Troponina elevada crônica em IRC avaliar delta >20% para agudização, em ICC descompensada ajustar terapia otimizada, em sepse/choque tratar causa base. Sempre correlacionar com apresentação clínica (dor torácica típica aumenta probabilidade pré-teste), alterações ECG (supra de ST ativa código infarto imediatamente, infra/inversão T sugerem isquemia), cinética enzimática (curva ascendente-descendente típica de SCA versus platô de lesão crônica). Documentar valores absolutos e deltas em prontuário. Compartilhar decisões com paciente considerando disparidades de tratamento documentadas - mulheres recebem menos intervenções mesmo com diagnóstico adequado. Encaminhar para reabilitação cardíaca e prevenção secundária rigorosa independente da estratégia inicial escolhida.',

  last_review = CURRENT_DATE
WHERE id = '247a9b99-ec59-4dd7-b7a3-b5482b1dd553';

-- Insert scientific articles (checking for DOI duplicates)
-- Article 1: Systematic Review 2024
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  -- Check if article exists by DOI
  SELECT id INTO v_article_id FROM articles WHERE doi = '10.1016/j.clinthera.2024.09.025';

  IF v_article_id IS NULL THEN
    -- Insert new article
    INSERT INTO articles (title, authors, journal, publish_date, pm_id, doi, article_type)
    VALUES (
      'Systematic Review of Sex-specific High Sensitivity Cardiac Troponin I and T Thresholds',
      'Cao M, Pierce AE, Norman MS, Thakur B, Diercks K, Hale C, Issioui Y, Diercks DB',
      'Clinical Therapeutics',
      '2024-12-01'::date,
      '39505672',
      '10.1016/j.clinthera.2024.09.025',
      'review'
    )
    RETURNING id INTO v_article_id;
  END IF;

  -- Link to score item (ignore if already linked)
  INSERT INTO article_score_items (score_item_id, article_id)
  VALUES ('247a9b99-ec59-4dd7-b7a3-b5482b1dd553', v_article_id)
  ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- Article 2: High-STEACS Trial (RCT) 2019
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  SELECT id INTO v_article_id FROM articles WHERE doi = '10.1016/j.jacc.2019.07.082';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (title, authors, journal, publish_date, pm_id, doi, article_type)
    VALUES (
      'Sex-Specific Thresholds of High-Sensitivity Troponin in Patients With Suspected Acute Coronary Syndrome',
      'Lee KK, Ferry AV, Anand A, Strachan FE, Chapman AR, Kimenai DM, Meex SJR, Berry C, Findlay I, Reid A, Gray A, Collinson PO, Apple FS, McAllister DA, Maguire D, Fox KAA, Newby DE, Tuck C, Keerie C, Weir CJ, Shah ASV, Mills NL',
      'Journal of the American College of Cardiology',
      '2019-10-22'::date,
      '31623760',
      '10.1016/j.jacc.2019.07.082',
      'clinical_trial'
    )
    RETURNING id INTO v_article_id;
  END IF;

  INSERT INTO article_score_items (score_item_id, article_id)
  VALUES ('247a9b99-ec59-4dd7-b7a3-b5482b1dd553', v_article_id)
  ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- Article 3: Diagnostic Impact Study 2016
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  SELECT id INTO v_article_id FROM articles WHERE doi = '10.1373/clinchem.2015.252569';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (title, authors, journal, publish_date, pm_id, doi, article_type)
    VALUES (
      'Impact of High-Sensitivity Troponin I Testing with Sex-Specific Cutoffs on the Diagnosis of Acute Myocardial Infarction',
      'Trambas C, Pickering JW, Than M, Bain C, Nie L, Paul E, Dart A, Broughton A, Schneider HG',
      'Clinical Chemistry',
      '2016-06-01'::date,
      '27117468',
      '10.1373/clinchem.2015.252569',
      'research_article'
    )
    RETURNING id INTO v_article_id;
  END IF;

  INSERT INTO article_score_items (score_item_id, article_id)
  VALUES ('247a9b99-ec59-4dd7-b7a3-b5482b1dd553', v_article_id)
  ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- Article 4: Diagnostic and Prognostic Value Study 2024
DO $$
DECLARE
  v_article_id uuid;
BEGIN
  SELECT id INTO v_article_id FROM articles WHERE doi = '10.1093/ehjacc/zuad131';

  IF v_article_id IS NULL THEN
    INSERT INTO articles (title, authors, journal, publish_date, pm_id, doi, article_type)
    VALUES (
      'Diagnostic and prognostic value of the sex-specific 99th percentile of four high-sensitivity cardiac troponin assays in patients with suspected myocardial infarction',
      'Lehmacher J, Sörensen NA, Twerenbold R, Goßling A, Haller PM, Hartikainen TS, Schock A, Toprak B, Zeller T, Westermann D, Neumann JT',
      'European Heart Journal - Acute Cardiovascular Care',
      '2024-02-09'::date,
      '37890108',
      '10.1093/ehjacc/zuad131',
      'research_article'
    )
    RETURNING id INTO v_article_id;
  END IF;

  INSERT INTO article_score_items (score_item_id, article_id)
  VALUES ('247a9b99-ec59-4dd7-b7a3-b5482b1dd553', v_article_id)
  ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

COMMIT;

-- Verification queries
SELECT
  si.id,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_relevance_length,
  LENGTH(si.patient_explanation) as patient_explanation_length,
  LENGTH(si.conduct) as conduct_length,
  si.last_review
FROM score_items si
WHERE si.id = '247a9b99-ec59-4dd7-b7a3-b5482b1dd553';

SELECT
  a.pm_id,
  a.title,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '247a9b99-ec59-4dd7-b7a3-b5482b1dd553'
ORDER BY a.publish_date DESC;
