-- Suplementação na arquitetura hormonal (material da Arboretum): eixos androgênico, DHEA e HPA.
--
-- Três fórmulas, extraídas com fatia por coluna — o PDF é diagramado em duas colunas e a leitura
-- ingênua misturava componentes de fórmulas vizinhas.
--
-- Uma ressalva de literatura entra junto: o material apresenta Tribulus, Testofen e Eurycoma como
-- otimizadores de testosterona livre. Para o Tribulus a evidência não sustenta: em revisão
-- sistemática de 2025, oito de dez ensaios não mostraram mudança no perfil androgênico, e os dois
-- que mostraram tiveram magnitude pequena (60 a 70 ng/dL) em homens com hipogonadismo. Isso fica
-- escrito na substância, não escondido.

BEGIN;

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, source, evidence_status, indications, dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Testofen', 'TESTOFEN, extrato de feno-grego padronizado', 'mg',
  300, 50, 600, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Extrato padronizado de feno-grego usado para suporte androgênico e libido.',
  '50 mg na fórmula do material. Os ensaios do próprio ingrediente usam 300 a 600 mg/dia — a dose da fórmula fica bem abaixo disso.', true, now(), now()),

 (uuid_generate_v7(), 'Eurycoma longifolia', 'long jack, tongkat ali, EURYCOMA', 'mg',
  200, 100, 400, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Fitoterápico usado em libido, estresse e suporte androgênico.',
  '200 mg no material. Ensaios costumam usar 200 a 400 mg/dia de extrato padronizado.', true, now(), now()),

 (uuid_generate_v7(), 'Turkesterone', 'turkesterona, Ajuga turkestanica', 'mg',
  300, 250, 500, 'por_dia', 0.45, 'classe', 'parceiro', 'pending',
  'Ecdisteroide vegetal apresentado como anabolizante não hormonal.',
  '300 mg no material. Fica como pendente de propósito: os dados humanos são escassos e a promessa anabólica repete a linguagem que a Resolução CFM 2.333/2023 restringe para hormônio — em fitoterápico não é vedação, mas a expectativa criada no paciente é a mesma.', true, now(), now()),

 (uuid_generate_v7(), 'Robuvit', 'ROBUVIT, extrato de carvalho francês, Quercus robur', 'mg',
  200, 100, 300, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Extrato de carvalho francês (roburinas) com dados em fadiga e recuperação.',
  '120 mg no material; os ensaios do ingrediente usam 200 a 300 mg/dia.', true, now(), now()),

 (uuid_generate_v7(), 'UbiQsome', 'UBIQSOME, coenzima Q10 fitossoma', 'mg',
  100, 50, 200, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Coenzima Q10 em veiculação fitossomal, com biodisponibilidade maior que a ubiquinona pura.',
  '100 mg no material. Como é CoQ10 veiculada, o teto de 200 mg/dia do Anexo IV da IN 28 para coenzima Q10 se aplica.', true, now(), now()),

 (uuid_generate_v7(), 'Panax ginseng', 'ginseng coreano, ginseng vermelho', 'mg',
  200, 100, 400, 'por_dia', 0.45, 'classe', 'parceiro', 'suggested',
  'Adaptógeno com dados em fadiga, cognição e resposta adrenal.',
  '150 mg no material. Ensaios usam 200 a 400 mg/dia de extrato padronizado em ginsenosídeos.', true, now(), now())
ON CONFLICT DO NOTHING;

-- A ressalva sobre o Tribulus vai na substância que já existe no catálogo.
UPDATE magistral_components
   SET dose_reference = coalesce(dose_reference || ' ', '') ||
       'Revisão sistemática de 2025: oito de dez ensaios sem mudança no perfil androgênico; os dois positivos tiveram magnitude pequena (60 a 70 ng/dL) em homens com hipogonadismo. Metanálise de 2023 encontrou aumento não significativo de testosterona e LH. Prescrever com essa expectativa, e não com a de reposição.',
       evidence_status = 'suggested'
 WHERE name = 'Tribulus terrestris' AND coalesce(dose_reference,'') NOT LIKE '%oito de dez%';

COMMIT;
