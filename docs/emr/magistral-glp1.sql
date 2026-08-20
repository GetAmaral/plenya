-- Suporte durante o uso de análogos de GLP-1 — material da Arboretum.
--
-- É o melhor documento do lote: cada substância vem com FAIXA DE DOSE e REFERÊNCIA, o que resolve
-- o problema que o formulário criou (faixa numérica sem base). As faixas abaixo são as do
-- material, com a citação guardada no texto de posologia.
--
-- Três ressalvas minhas, que o material não traz, entram como observação da substância.

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- 1. Substâncias novas, com a faixa e a referência do material
-- ---------------------------------------------------------------------------------------------
INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, sachet_ok, bitterness, source, evidence_status, indications, dose_reference,
   is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Psyllium', 'PLANTAGO OVATE, PSILLYUM, Plantago ovata, psilium', 'g',
  5, 3, 20, 'por_dia', 0.55, 'classe', true, 1, 'parceiro', 'suggested',
  'Fibra que absorve água no intestino, aumenta o volume do bolo fecal e estimula o trânsito. Nos análogos de GLP-1 entra pela constipação por retardo do esvaziamento gástrico.',
  'Até 20 g/dia (McRorie & McKeown, 2017). Tomar com bastante líquido; sem água suficiente a fibra piora a obstipação.', true, now(), now()),

 (uuid_generate_v7(), 'Saccharomyces boulardii', 'S. boulardii, levedura probiótica', 'mg',
  250, 150, 500, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Levedura probiótica que reduz episódios e gravidade de diarreia e ajuda no equilíbrio da microbiota.',
  'De 150 a 500 mg/dia (Pal et al., 2020).', true, now(), now()),

 (uuid_generate_v7(), 'Bacillus clausii', 'B. clausii', 'bilhões UFC',
  2, 0.1, 10, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Probiótico esporulado, com efeito na diarreia aguda e na diversidade da microbiota.',
  'De 0,1 a 10 bilhões de UFC/dia (Ianiro et al., 2018).', true, now(), now()),

 (uuid_generate_v7(), 'Simeticone', 'simeticona, dimeticona ativada', 'mg',
  80, 20, 500, 'por_dia', 0.60, 'classe', true, 0, 'parceiro', 'suggested',
  'Antiflatulento: reduz a tensão superficial das bolhas de gás e alivia distensão e desconforto abdominal.',
  'De 20 a 500 mg/dia (Ingold & Akhondi, 2025).', true, now(), now()),

 (uuid_generate_v7(), 'Carvão ativado', 'carvão vegetal ativado', 'mg',
  300, 200, 500, 'por_dia', 0.35, 'classe', true, 1, 'parceiro', 'suggested',
  'Adsorvente usado no alívio de flatulência e distensão por gases.',
  'De 200 a 500 mg/dia (Silberman, Galuska, Taylor, 2023). ATENÇÃO: adsorve fármacos e nutrientes de forma inespecífica. Afastar pelo menos duas horas de qualquer medicamento e das demais fórmulas, ou o carvão leva junto o que deveria ser absorvido.', true, now(), now()),

 (uuid_generate_v7(), 'HMB', 'beta-hidroxi-beta-metilbutirato, HMB CALCIO', 'mg',
  3000, 1000, 3000, 'por_dia', 0.55, 'classe', true, 2, 'parceiro', 'suggested',
  'Metabólito da leucina que reduz o catabolismo muscular e estimula a síntese proteica. Nos análogos de GLP-1 entra pela perda de massa magra que acompanha a perda rápida de peso.',
  'De 1 a 3 g/dia (Kaczka et al., 2019).', true, now(), now()),

 (uuid_generate_v7(), 'PeptiStrong', 'peptídeos de fava, Vicia faba', 'g',
  2.4, 2.4, 20, 'por_dia', 0.50, 'classe', true, 2, 'parceiro', 'suggested',
  'Peptídeos bioativos de fava com efeito sobre preservação de massa magra e síntese muscular.',
  'Dose usual de 2,4 g/dia; o material cita 10 g pela manhã e 10 g à noite em atrofia muscular (Kerr et al., 2023). A distância entre as duas indicações é grande: conferir a fonte antes de usar a dose alta.', true, now(), now()),

 (uuid_generate_v7(), 'Nutricolin', 'silício orgânico, NUTRICOLIN, ácido ortossilícico', 'mg',
  100, 50, 300, 'por_dia', 0.60, 'classe', true, 0, 'parceiro', 'suggested',
  'Silício orgânico biodisponível; estimula síntese de colágeno e queratina, com efeito sobre elasticidade da pele e resistência capilar.',
  'De 50 a 300 mg/dia (Araújo et al., 2016).', true, now(), now()),

 (uuid_generate_v7(), 'Akkermansia muciniphila', 'Bio MAMPs, akkermansia', 'mg',
  75, 50, 100, 'por_dia', 0.45, 'classe', true, 0, 'parceiro', 'suggested',
  'Postbiótico com ação sobre a barreira intestinal e a imunomodulação, com dados em disbiose, inflamação e controle de peso.',
  'De 50 a 100 mg/dia (Cani & Knauf, 2021).', true, now(), now()),

 (uuid_generate_v7(), 'Slendesta', 'extrato de batata, inibidor de protease PI2', 'mg',
  100, 50, 300, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'suggested',
  'Extrato de batata que estimula colecistoquinina e reduz o apetite.',
  'Até 300 mg/dia (Peters et al., 2010).', true, now(), now()),

 (uuid_generate_v7(), 'Motility', 'MOTILITY, blend para motilidade', 'mg',
  300, 150, 500, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'pending',
  'Blend do fornecedor para hidratação do bolo fecal e proteção da mucosa intestinal.',
  'De 150 a 500 mg/dia. O próprio material dá como fonte "literatura do fornecedor", sem citação externa — fica como pendente até haver referência independente.', true, now(), now()),

 (uuid_generate_v7(), 'Fibregum B', 'goma acácia, fibra de acácia', 'g',
  2, 0.1, 10, 'por_dia', 0.50, 'classe', true, 0, 'parceiro', 'suggested',
  'Fibra prebiótica de goma acácia; equilibra a microbiota e alivia gases e inchaço.',
  'De 0,1 a 10 g/dia (Singh et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Cistina', 'L-cistina', 'mg',
  150, 100, 200, 'por_dia', 0.65, 'classe', true, 2, 'parceiro', 'suggested',
  'Aminoácido sulfurado, matéria-prima da queratina; usado em queda capilar.',
  'De 100 a 200 mg/dia (Riegel et al., 2020).', true, now(), now()),

 (uuid_generate_v7(), 'Metionina', 'L-metionina', 'mg',
  200, 100, 500, 'por_dia', 0.65, 'classe', true, 2, 'parceiro', 'suggested',
  'Aminoácido essencial sulfurado que participa da produção de queratina.',
  'De 100 a 500 mg/dia (Milani et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Saw palmetto', 'SAW PALMETO, Serenoa repens', 'mg',
  320, 160, 360, 'por_dia', 0.45, 'classe', true, 1, 'parceiro', 'suggested',
  'Inibidor de 5-alfa-redutase de origem vegetal, usado em queda capilar androgenética.',
  'Até 360 mg/dia (Sudeep et al., 2023).', true, now(), now()),

 (uuid_generate_v7(), 'Verisol', 'peptídeos bioativos de colágeno, VERISOL', 'g',
  2.5, 2.5, 2.5, 'por_dia', 0.50, 'classe', true, 1, 'parceiro', 'suggested',
  'Peptídeos bioativos de colágeno com efeito em firmeza da pele e redução de rugas.',
  'Dose única de 2,5 g/dia nos ensaios (Proksch et al., 2014).', true, now(), now())
ON CONFLICT DO NOTHING;

-- ---------------------------------------------------------------------------------------------
-- 2. Interferência em exame — a ressalva que o material não faz
-- ---------------------------------------------------------------------------------------------

-- A fórmula capilar do material leva biotina 10 mg. A partir de 5 mg/dia a biotina interfere em
-- imunoensaio biotinilado e devolve TSH, troponina, hormônios tireoidianos e hCG falsamente altos
-- ou baixos conforme o formato do ensaio (alerta do FDA de 2017 e 2019; orientação da AACC).
-- Neste sistema isso fecha um ciclo ruim: a receita sai daqui e o exame corrompido volta para cá,
-- alimentando as regras de dose dinâmica.
UPDATE magistral_components
   SET assay_interference = 'acima de 5 mg/dia interfere em imunoensaio biotinilado (TSH, T4 livre, troponina, hCG, hormônios), com resultado falsamente alto ou baixo conforme o ensaio. Suspender 3 dias antes da coleta e avisar o laboratório',
       assay_interference_dose = 5
 WHERE name = 'Biotina';

-- ---------------------------------------------------------------------------------------------
-- 3. Par curado: o carvão ativado leva junto o que deveria ser absorvido
-- ---------------------------------------------------------------------------------------------
WITH carvao AS (SELECT id FROM magistral_components WHERE name = 'Carvão ativado'),
     alvos AS (SELECT id, name FROM magistral_components
                WHERE name IN ('Lipase', 'Protease', 'Alfa-amilase', 'Metilcobalamina', 'Ferro',
                               'Zinco quelato', 'Palmitato de ascorbila'))
INSERT INTO magistral_incompatibilities (id, component_a_id, component_b_id, severity, mechanism, note)
SELECT uuid_generate_v7(), carvao.id, alvos.id, 'warn',
       'o carvão ativado adsorve de forma inespecífica e reduz a absorção do que estiver junto',
       'Não impede a associação, mas na mesma cápsula o carvão tira parte do efeito do outro ativo. Afastar pelo menos duas horas.'
  FROM carvao, alvos
 WHERE NOT EXISTS (SELECT 1 FROM magistral_incompatibilities i
                    WHERE i.component_a_id = carvao.id AND i.component_b_id = alvos.id);

COMMIT;

-- ---------------------------------------------------------------------------------------------
-- 4. As que já existiam do formulário ganham a faixa e a referência deste material
--
-- O formulário só dava o nome; este documento dá faixa com citação. Quando as duas fontes falam
-- da mesma substância, vale a que tem referência.
-- ---------------------------------------------------------------------------------------------
BEGIN;

UPDATE magistral_components SET usual_dose=5, min_dose=3, max_dose=20, dose_basis='por_dia', default_unit='g',
  dose_reference='Até 20 g/dia (McRorie & McKeown, 2017). Tomar com bastante líquido; sem água suficiente a fibra piora a obstipação.',
  indications='Fibra que absorve água no intestino, aumenta o volume do bolo fecal e estimula o trânsito. Nos análogos de GLP-1 entra pela constipação por retardo do esvaziamento gástrico.',
  evidence_status='suggested' WHERE name='Psyllium';

UPDATE magistral_components SET usual_dose=3000, min_dose=1000, max_dose=3000, dose_basis='por_dia',
  dose_reference='De 1 a 3 g/dia (Kaczka et al., 2019).',
  indications='Metabólito da leucina que reduz o catabolismo muscular e estimula a síntese proteica. Nos análogos de GLP-1 entra pela perda de massa magra que acompanha a perda rápida de peso.',
  evidence_status='suggested' WHERE name='HMB';

UPDATE magistral_components SET usual_dose=100, min_dose=50, max_dose=300, dose_basis='por_dia',
  dose_reference='De 50 a 300 mg/dia (Araújo et al., 2016).',
  indications='Silício orgânico biodisponível; estimula síntese de colágeno e queratina, com efeito sobre elasticidade da pele e resistência capilar.',
  evidence_status='suggested' WHERE name='Nutricolin';

UPDATE magistral_components SET usual_dose=2.5, min_dose=2.5, max_dose=2.5, dose_basis='por_dia', default_unit='g',
  dose_reference='Dose única de 2,5 g/dia nos ensaios (Proksch et al., 2014).',
  indications='Peptídeos bioativos de colágeno com efeito em firmeza da pele e redução de rugas.',
  evidence_status='suggested' WHERE name='Verisol';

UPDATE magistral_components SET usual_dose=600, min_dose=300, max_dose=900, dose_basis='por_dia',
  dose_reference='De 300 a 900 mg/dia (Pirahanchi & Sharma, 2025).',
  indications='Enzima que auxilia na digestão de lipídios. Nos análogos de GLP-1 entra pela alteração da secreção biliar.',
  evidence_status='suggested' WHERE name='Lipase';

UPDATE magistral_components SET usual_dose=200, min_dose=100, max_dose=300, dose_basis='por_dia',
  dose_reference='De 100 a 300 mg/dia (Liu et al., 2024).',
  indications='Enzima que favorece a degradação de proteínas e a absorção.',
  evidence_status='suggested' WHERE name='Protease';

UPDATE magistral_components SET usual_dose=100, min_dose=30, max_dose=300, dose_basis='por_dia',
  dose_reference='De 30 a 300 mg/dia (Ianiro et al., 2016).',
  indications='Enzima que facilita a degradação de carboidratos e reduz fermentação e desconforto gastrointestinal.',
  evidence_status='suggested' WHERE name='Alfa-amilase';

COMMIT;

-- O limiar de interferência precisa estar na unidade em que a substância é cadastrada. Biotina
-- está em mcg no catálogo, e o limiar tinha sido gravado em mg: 500 mcg de rotina disparavam
-- alarme (500 > 5) e 10 mg passavam batido no teto. 5 mg = 5.000 mcg.
BEGIN;
UPDATE magistral_components SET assay_interference_dose = 5000
 WHERE name = 'Biotina' AND default_unit = 'mcg' AND assay_interference_dose = 5;
COMMIT;
