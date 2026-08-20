-- Pentravan — veículo transdérmico (material da Fagron).
--
-- O que este documento acrescenta ao sistema é de outra natureza: não é dose oral, é VIA. Entram
-- as substâncias que o material formula em Pentravan, cada uma com a categoria de receita que
-- carrega — e é isso que faz "testosterona 50 mg" sair como Controle Especial em vez de receita
-- simples.
--
-- A tabela de permeação do material (ativo, concentração, tecido, percentual em 24 ou 48 h) vai
-- na posologia de cada substância: é o dado que justifica a via.

BEGIN;

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, default_category, regulatory_note, source, evidence_status, indications,
   dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Testosterona micronizada', 'testosterona, testosterona Fagron Micro', 'mg',
  50, 0.5, 90, 'por_dia', 0.60, 'classe', 'c5',
  'Esteroide androgênico e anabolizante. A Resolução CFM 2.333/2023 veda a prescrição com finalidade estética, ganho de massa muscular ou desempenho esportivo; a reposição exige deficiência comprovada com nexo causal. Lista C5 da Portaria 344/98: Receituário de Controle Especial em duas vias.',
  'parceiro', 'suggested',
  'Reposição androgênica por via transdérmica ou vulvar. No material, deficiência androgênica feminina usa 0,5 a 5 mg/mL e declínio androgênico masculino 40 a 90 mg/mL.',
  'Em Pentravan a 1% (10 mg/g), permeação de 68,3% em 24 h em pele humana; a 5% (50 mg/g), 68,31% em 24 h e 76,8% em 48 h (Polonini et al., 2014).', true, now(), now()),

 (uuid_generate_v7(), 'Oxandrolona', 'oxandrolona Fagron', 'mg',
  10, 5, 20, 'por_dia', 0.60, 'classe', 'c5',
  'Esteroide anabolizante. A Resolução CFM 2.333/2023 veda a prescrição com finalidade estética, ganho de massa muscular ou desempenho esportivo — e a própria fórmula do material se intitula "sarcopenia e ganho de peso". Sarcopenia com deficiência documentada é outra conversa, e precisa estar escrita no prontuário. Lista C5.',
  'parceiro', 'suggested',
  'Anabolizante usado por via transdérmica em sarcopenia no material do fornecedor.',
  'Em Pentravan a 2% (20 mg/g), permeação de 25,9% em 24 h em pele humana (Polonini et al., 2017).', true, now(), now()),

 (uuid_generate_v7(), '17-beta-estradiol', 'estradiol, E2', 'mg',
  1, 0.25, 2, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estrogênio para sintomas climatéricos por via transdérmica.',
  'Em Pentravan a 0,1% (1 mg/g), permeação de 86,33% em 24 h e 99,9% em 48 h em pele humana (Polonini et al., 2014). No material, 0,25 a 2 mg/mL.', true, now(), now()),

 (uuid_generate_v7(), 'Estriol', 'E3', 'mg',
  4, 2, 8, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estrogênio de ação local, usado em climatério e em prevenção de flacidez cutânea.',
  'Em Pentravan a 0,4% (4 mg/g), permeação de 43,67% em 24 h (Polonini et al., 2014). No material, 2 a 8 mg/mL por via transdérmica e 0,3% no uso facial.', true, now(), now()),

 (uuid_generate_v7(), 'Progesterona micronizada', 'progesterona', 'mg',
  50, 20, 80, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Proteção endometrial por via vaginal.',
  'De 20 a 80 mg por grama de Pentravan, aplicado por via vaginal nos últimos 13 a 15 dias do mês (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Gestrinona', 'gestrinona Fagron', 'mg',
  5, 2.5, 5, 'por_dia', 0.60, 'classe', 'simple',
  'Esteroide com ação androgênica. Uso em endometriose é off-label no Brasil e ganhou escrutínio depois dos implantes hormonais: registrar indicação e consentimento.',
  'parceiro', 'suggested',
  'Endometriose por via vaginal.',
  'Em Pentravan a 0,5% (5 mg/g), permeação de 61,4% em 24 h em mucosa vaginal suína. Estudos clínicos de Maia Jr. et al. (2015 a 2019) com gestrinona vaginal em Pentravan.', true, now(), now()),

 (uuid_generate_v7(), 'Danazol', 'danazol', 'mg',
  50, 50, 100, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Mastalgia cíclica por via transdérmica na mama.',
  '50 mg por grama de Pentravan, uma aplicação ao dia (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Citrato de sildenafila', 'sildenafila', '%',
  0.25, 0.25, 0.25, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Estimulante sexual feminino por uso vulvar.',
  '0,25% em Pentravan, 1 mL na região dos lábios vaginais 30 minutos antes da relação (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Tadalafila', 'tadalafila', 'mg',
  5, 5, 5, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Inibidor de fosfodiesterase-5 associado à modulação de testosterona no material.',
  'Em Pentravan a 0,5% (5 mg/g), permeação de 89,07% em 12 h em pele humana (Calixto, 2015).', true, now(), now()),

 (uuid_generate_v7(), 'Alprostadil', 'alprostadil Fagron, PGE1', 'mcg',
  100, 100, 1000, 'por_dose', 0.60, 'classe', 'simple',
  'Prostaglandina E1. Uso tópico peniano e vulvar é off-label; registrar indicação.',
  'parceiro', 'suggested',
  'Disfunção erétil e transtorno da excitação sexual feminina, por uso tópico sob demanda.',
  'Cada pump com 100 mcg; 5 a 10 pumps por aplicação, no mínimo três vezes por semana (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Mesilato de fentolamina', 'fentolamina', 'mg',
  4, 4, 4, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Alfabloqueador associado ao alprostadil na disfunção erétil tópica.',
  '4 mg por pump, associado a 100 mcg de alprostadil (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Baclofeno', 'baclofen', 'mg',
  50, 25, 50, 'por_dia', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Relaxante muscular de ação central, usado por via tópica na vulvodínia.',
  '50 mg por mL de Pentravan, associado a PEA, 1 a 2 aplicações ao dia (material da Fagron).', true, now(), now()),

 (uuid_generate_v7(), 'Palmitoiletanolamida', 'PEA, PEA BioActive', 'mg',
  10, 10, 600, 'por_dia', 0.55, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Amida lipídica endógena com ação analgésica e anti-inflamatória; na vulvodínia entra por via tópica.',
  '10 mg por mL de Pentravan no material. O teto da IN 28 para uso oral é de 600 mg/dia.', true, now(), now()),

 (uuid_generate_v7(), 'Metformina', 'metformina HCl, cloridrato de metformina', 'mg',
  75, 50, 100, 'por_dia', 0.65, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Modulação de AMPK para longevidade, por via transdérmica no material.',
  'Em Pentravan a 10% (100 mg/g), permeação de 46,7% em 24 h em pele humana (Polonini et al., 2019). No material, 50 a 100 mg/mL duas vezes ao dia.', true, now(), now()),

 (uuid_generate_v7(), 'Miodesin', 'MIODESIN', 'mg',
  170, 170, 170, 'por_dia', 0.50, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Blend anti-inflamatório do fornecedor, usado em miomatose e endometriose por via vaginal.',
  '170 mg por grama de Pentravan, uma aplicação à noite por até dois meses. Estudos de Maia Jr. et al. (2018, 2019).', true, now(), now()),

 (uuid_generate_v7(), 'SiliciuMax', 'silício líquido, SILICIUMAX', '%',
  5, 5, 30, 'por_dose', 0.60, 'classe', 'simple', NULL, 'parceiro', 'suggested',
  'Silício em forma líquida para uso tópico, em flacidez e envelhecimento cutâneo.',
  'Em Pentravan a 30% (300 mg/g), permeação de 60% em 24 h em pele humana. No material facial, 5%.', true, now(), now())
ON CONFLICT DO NOTHING;

-- Pentravan como veículo. Entra no catálogo para ter densidade e ficar disponível na busca, mas
-- sem faixa de dose: é base, não ativo.
INSERT INTO magistral_components
  (id, name, synonyms, default_unit, bulk_density, density_source, source, evidence_status,
   indications, dose_reference, is_active, created_at, updated_at)
VALUES (uuid_generate_v7(), 'Pentravan', 'PENTRAVAN, veículo transdérmico lipossomal', 'g',
  1.0, 'classe', 'parceiro', 'suggested',
  'Veículo transdérmico em matriz fosfolipídica lamelar com partículas nanossomais, para permeação de ativos em pele íntegra e mucosa.',
  'Sem faixa de dose de propósito: é veículo, entra como qsp. A concentração do ativo é que define a fórmula.', true, now(), now())
ON CONFLICT DO NOTHING;

COMMIT;
