BEGIN;

-- Insert articles
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
  'Current ovulation and luteal phase tracking methods and technologies for fertility and family planning: a review',
  'Wegrzynowicz AK, Eyvazzadeh A, Beckley A',
  'Seminars in Reproductive Medicine',
  '2024-09-20'::date,
  'en',
  'review',
  '10.1055/s-0044-1791190',
  '39302838',
  'Review of current methods and emerging technologies for tracking ovulation and luteal phase including apps hormonal monitoring and wearables.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW();

INSERT INTO articles (title, authors, journal, publish_date, language, article_type, pm_id, abstract)
VALUES (
  'Physiology, Follicle Stimulating Hormone',
  'Stamler R, Kapoor N, Ioachimescu AG',
  'StatPearls',
  '2024-08-12'::date,
  'en',
  'review',
  'NBK535442',
  'Comprehensive review of FSH physiology across the menstrual cycle including luteal phase FSH dynamics.'
);

-- Link articles
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '0726b373-4b78-4cb8-91a9-0916681aef61'::uuid
FROM articles WHERE doi = '10.1055/s-0044-1791190'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '0726b373-4b78-4cb8-91a9-0916681aef61'::uuid
FROM articles WHERE pm_id = 'NBK535442'
ON CONFLICT DO NOTHING;

-- Update score item
UPDATE score_items SET
  clinical_relevance = 'O FSH (hormônio folículo-estimulante) durante a fase lútea apresenta valores característicamente baixos (1,1-9,2 mUI/mL), refletindo a retroalimentação negativa exercida pelo corpo lúteo através da progesterona e estradiol. Esta supressão fisiológica do FSH é essencial para prevenir o recrutamento de novos folículos, permitindo que o corpo lúteo mantenha o endométrio preparado para potencial implantação. A medição do FSH na fase lútea possui aplicações clínicas limitadas quando comparada à fase folicular, porém valores elevados podem indicar disfunção do corpo lúteo, insuficiência ovariana prematura (IOP) ou transição menopausal. Segundo diretrizes ASRM 2025, FSH>25 mUI/mL em duas ocasiões (intervalo ≥4 semanas) em mulheres <40 anos sugere IOP. Valores persistentemente baixos (<1 mUI/mL) podem indicar hipopituitarismo. A interpretação deve sempre considerar os níveis concomitantes de LH, estradiol e progesterona para caracterizar adequadamente a função do corpo lúteo e eixo hipotálamo-hipófise-ovário.',
  
  patient_explanation = 'O FSH (hormônio folículo-estimulante) é um hormônio produzido pela hipófise que ajuda a controlar seu ciclo menstrual. Durante a fase lútea (período após a ovulação até a próxima menstruação, aproximadamente 14 dias), os níveis de FSH ficam naturalmente mais baixos (1,1-9,2 mUI/mL) porque o corpo lúteo (estrutura que se forma após a ovulação) está produzindo progesterona, que "avisa" sua hipófise para diminuir o FSH. Isso é importante para que novos folículos não comecem a se desenvolver nesta fase. Valores muito elevados de FSH na fase lútea podem indicar que seus ovários não estão funcionando adequadamente ou que você está entrando na menopausa. Valores muito baixos podem sugerir problema na hipófise. É importante saber que a avaliação da fertilidade é feita preferencialmente medindo FSH no início do ciclo (dias 2-5), não na fase lútea. Se você precisa investigar sua função ovariana ou fertilidade, converse com seu médico sobre o melhor momento do ciclo para fazer os exames.',
  
  conduct = 'COLETA E TIMING: Coletar idealmente entre os dias 19-23 de um ciclo regular de 28 dias (fase lútea média). Para ciclos irregulares, solicitar dosagem seriada ou combinar com monitoramento ultrassonográfico do corpo lúteo. Coleta em jejum de 3-4h, pela manhã (7-10h) para minimizar variabilidade circadiana. INTERPRETAÇÃO POR FAIXAS: Normal (1,1-9,2 mUI/mL): Supressão fisiológica adequada pelo corpo lúteo funcionante. Elevado (>10 mUI/mL): Investigar insuficiência do corpo lúteo, transição menopausal, ou IOP. Solicitar FSH e estradiol em fase folicular para confirmar. Muito elevado (>25 mUI/mL): Forte suspeita de IOP se <40 anos. Confirmar com segunda dosagem após ≥4 semanas (critério ASRM 2025). Baixo (<1 mUI/mL): Considerar hipopituitarismo, uso de contraceptivos hormonais, ou hiperprolactinemia. INVESTIGAÇÃO COMPLEMENTAR OBRIGATÓRIA: LH (razão FSH/LH útil), Estradiol (corpo lúteo produz 15-30 pg/mL na fase lútea), Progesterona (>10 ng/mL confirma ovulação e fase lútea adequada; <3 ng/mL sugere fase lútea inadequada), TSH e Prolactina (descartar causas tireoidianas/hiperprolactinemia). CENÁRIOS CLÍNICOS: 1) FSH lúteo >10 mUI/mL + Progesterona <3 ng/mL: Insuficiência do corpo lúteo. Repetir em ciclo subsequente. Considerar suplementação com progesterona se ciclos recorrentes. 2) FSH lúteo >25 mUI/mL em mulher <40 anos: Investigação completa para IOP (cariótipo, FMR1, anticorpos anti-ovário). 3) FSH lúteo <1 mUI/mL: RNM hipofisária, perfil hormonal completo (cortisol, GH, outros eixos). LIMITAÇÕES: FSH fase lútea tem menor valor diagnóstico que FSH folicular para avaliação de reserva ovariana. Resultados devem ser interpretados no contexto clínico completo e outros hormônios. SEGUIMENTO: Repetir dosagem em fase folicular precoce (dias 2-5) para melhor avaliação da função ovariana e reserva.',
  
  last_review = NOW(),
  updated_at = NOW()
WHERE id = '0726b373-4b78-4cb8-91a9-0916681aef61';

COMMIT;

-- Verification
SELECT 
  id, name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = '0726b373-4b78-4cb8-91a9-0916681aef61') as articles
FROM score_items
WHERE id = '0726b373-4b78-4cb8-91a9-0916681aef61';
