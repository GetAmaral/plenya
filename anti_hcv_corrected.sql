BEGIN;

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'A 2024 global report on national policies and hepatitis C elimination',
  'Lemoine M, et al.',
  'The Lancet Gastroenterology & Hepatology',
  '2025-05-01'::date,
  'en',
  'research_article',
  '10.1016/S2468-1253(24)00350-9',
  '39396793',
  '33-country analysis showing only 30% met WHO HCV diagnosis targets, highlighting global gaps in elimination progress.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'Hepatitis C Virus: Epidemiological Challenges and Global Strategies',
  'Petruzziello A',
  'Viruses',
  '2025-07-01'::date,
  'en',
  'review',
  '10.3390/v16071000',
  '39066180',
  'Review of global HCV epidemiology: 1M annual infections, 242K deaths, prison prevalence 17.7%, and elimination strategies.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'SASLT guidelines: Update in HCV treatment 2024',
  'Saudi Association for the Study of Liver diseases and Transplantation',
  'Saudi Journal of Gastroenterology',
  '2024-01-01'::date,
  'en',
  'protocol',
  '10.4103/sjg.sjg_306_23',
  '38409782',
  'Updated guidelines recommending pangenotypic DAAs as first-line therapy with universal treatment approach.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'Cure of chronic HCV infection by direct-acting antivirals',
  'Schnyder A, et al.',
  'Frontiers in Immunology',
  '2025-01-01'::date,
  'en',
  'review',
  '10.3389/fimmu.2025.1546915',
  '39850066',
  'Review of DAA treatment achieving >95% cure rates with partial immune restoration and need for post-SVR surveillance.'
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '73411e66-c180-46ad-8f5b-7d67b26ef877'::uuid
FROM articles WHERE doi = '10.1016/S2468-1253(24)00350-9'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '73411e66-c180-46ad-8f5b-7d67b26ef877'::uuid
FROM articles WHERE doi = '10.3390/v16071000'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '73411e66-c180-46ad-8f5b-7d67b26ef877'::uuid
FROM articles WHERE doi = '10.4103/sjg.sjg_306_23'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '73411e66-c180-46ad-8f5b-7d67b26ef877'::uuid
FROM articles WHERE doi = '10.3389/fimmu.2025.1546915'
ON CONFLICT DO NOTHING;

UPDATE score_items SET
  clinical_relevance = 'O anti-HCV (anticorpo contra vírus da hepatite C) é um teste sorológico de triagem que detecta exposição prévia ou atual ao HCV, permanecendo positivo por toda vida após infecção independente de clearance viral. A infecção crônica por HCV afeta aproximadamente 58 milhões de pessoas globalmente (2022), causando 242.000 mortes anuais por cirrose e carcinoma hepatocelular. A revolução terapêutica com antivirais de ação direta (DAAs) pangenotípicos alcança taxas de cura >95% em 8-12 semanas, tornando a eliminação da hepatite C uma meta realista (WHO 2030). Contudo, análise de 33 países (2024) revela que apenas 30% atingem metas de diagnóstico, evidenciando lacunas críticas na triagem. Populações de alto risco incluem: pessoas privadas de liberdade (prevalência 17,7%), usuários de drogas injetáveis (39-67%), receptores de hemoderivados pré-1993, pacientes hemodiálise, profissionais saúde expostos, HIV+, e gestantes. O anti-HCV positivo isolado não diferencia infecção ativa vs resolvida, requerendo confirmação com HCV RNA quantitativo. Diretrizes atuais (SASLT 2024, WHO) recomendam tratamento universal independente de fibrose hepática dada a alta eficácia e segurança dos DAAs. Pós-cura (resposta virológica sustentada - RVS12), vigilância para carcinoma hepatocelular permanece mandatória em cirróticos, refletindo risco residual mesmo após eliminação viral.',
  
  patient_explanation = 'O anti-HCV é um exame de sangue que detecta anticorpos contra o vírus da hepatite C, indicando que você teve contato com esse vírus em algum momento da vida. Importante: um resultado positivo não significa necessariamente que você está com infecção ativa hoje, pois cerca de 15-25% das pessoas eliminam o vírus espontaneamente, mas os anticorpos permanecem positivos para sempre. Se seu anti-HCV for positivo, será necessário fazer outro exame chamado "HCV RNA" (carga viral) para confirmar se o vírus ainda está ativo no seu organismo. A hepatite C é transmitida principalmente por contato com sangue contaminado (compartilhamento de agulhas, transfusões antigas, procedimentos sem esterilização adequada) e pode causar dano progressivo ao fígado ao longo de décadas, levando à cirrose ou câncer hepático. A excelente notícia é que hoje a hepatite C tem CURA: tratamentos modernos com comprimidos (antivirais de ação direta) curam mais de 95% dos casos em apenas 8-12 semanas, sem efeitos colaterais graves. É fundamental confirmar o diagnóstico e iniciar tratamento precocemente para prevenir complicações hepáticas. Após a cura, embora o anti-HCV permaneça sempre positivo, você não transmitirá mais o vírus.',
  
  conduct = 'INTERPRETAÇÃO ANTI-HCV: Positivo: Solicitar HCV RNA qualitativo/quantitativo imediatamente para confirmar viremia ativa. Negativo com exposição recente (<6 meses): Repetir anti-HCV após 6 meses (janela imunológica) ou solicitar HCV RNA se alta suspeição clínica. HCV RNA CONFIRMAÇÃO: Detectável: Infecção crônica ativa - prosseguir workup completo. Indetectável + anti-HCV positivo: Infecção resolvida (clearance espontâneo ou tratamento prévio bem-sucedido). WORKUP INICIAL COMPLETO: Perfil hepático (ALT, AST, GGT, fosfatase alcalina, bilirrubinas), função hepática (albumina, TAP/INR, plaquetas), genotipagem HCV (pangenotípicos dispensam), quantificação carga viral basal, avaliação fibrose (elastografia hepática ou FIB-4/APRI), rastreamento coinfecções (HBsAg, anti-HBc, anti-HIV), hemograma completo, creatinina/TFG. TRATAMENTO - REGIMES PANGENOTÍPICOS PRIMEIRA LINHA (SASLT 2024): Sofosbuvir/Velpatasvir 400/100mg 1x/dia por 12 semanas (todos genótipos, com/sem cirrose). Glecaprevir/Pibrentasvir 100/40mg 3cp 1x/dia por 8-12 semanas (8 semanas se não cirrótico, 12 se cirrose compensada). Tratamento universal recomendado independente de grau de fibrose. MONITORAMENTO DURANTE TRATAMENTO: Adesão terapêutica, efeitos adversos, interações medicamentosas. Evitar amiodarona, carbamazepina, rifampicina, erva de São João (interações significativas). HCV RNA na semana 4 (opcional) e obrigatório 12 semanas pós-tratamento (RVS12). DEFINIÇÃO DE CURA: RVS12 (HCV RNA indetectável 12 semanas após fim tratamento) = cura virológica >99% durável. VIGILÂNCIA PÓS-CURA: Cirrose (F4): rastreamento semestral carcinoma hepatocelular (USG abdominal + alfa-fetoproteína) por toda vida, mesmo após RVS. Fibrose avançada (F3): considerar vigilância anual. Ausência fibrose significativa: seguimento hepático anual suficiente.',
  
  last_review = NOW(),
  updated_at = NOW()
WHERE id = '73411e66-c180-46ad-8f5b-7d67b26ef877';

COMMIT;

SELECT 
  id, name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = '73411e66-c180-46ad-8f5b-7d67b26ef877') as articles
FROM score_items
WHERE id = '73411e66-c180-46ad-8f5b-7d67b26ef877';
