-- Peptídeos biomiméticos (material de dermatologia).
--
-- Deste documento entram as SUBSTÂNCIAS, não as fórmulas. O PDF é diagramado em duas colunas e a
-- extração mistura componentes de fórmulas vizinhas: montar receita a partir de uma atribuição
-- que eu não consigo garantir seria pior do que não montar. As concentrações abaixo são as que o
-- material usa, e cada uma aparece de forma inequívoca ao lado do próprio nome.
--
-- Todas em uso tópico, unidade em porcentagem — é como se escreve fórmula dermatológica.

BEGIN;

INSERT INTO magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis, bulk_density,
   density_source, source, evidence_status, indications, dose_reference, is_active, created_at, updated_at) VALUES

 (uuid_generate_v7(), 'Argireline', 'acetil hexapeptídeo-8, ARGIRELINE', '%', 5, 3, 8, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo inibidor de neurotransmissor: reduz a contração muscular mimética e suaviza linhas de expressão.',
  'De 3 a 8% no material, em sérum de uso facial. Também usado como suporte pós-toxina botulínica.', true, now(), now()),

 (uuid_generate_v7(), 'Syn-Ake', 'dipeptídeo diaminobutiroil benzilamida, SYN-AKE', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo miorrelaxante tópico, análogo sintético de peptídeo do veneno de Tropidolaemus wagleri.',
  '2% no material, associado a Argireline em sérum facial.', true, now(), now()),

 (uuid_generate_v7(), 'Munapsys', 'acetil octapeptídeo-3, MUNAPSYS', '%', 3, 1, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com ação sobre a contração muscular mimética, associado aos demais em rugas dinâmicas.',
  'De 1 a 3% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Matrixyl Synthe-6', 'palmitoil tripeptídeo-38, MATRIXYL SYNTHE 6', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo sinalizador: estimula fibroblasto e a síntese de matriz extracelular.',
  '2% no material, em fórmula de firmeza para pele madura.', true, now(), now()),

 (uuid_generate_v7(), 'Idealift', 'peptídeo tensor, IDEALIFT', '%', 2, 2, 4, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com efeito tensor e de firmeza cutânea.',
  '2% no material.', true, now(), now()),

 (uuid_generate_v7(), 'GHK-Cu', 'peptídeo de cobre, tripeptídeo-1 de cobre, copper peptide', '%', 1, 0.5, 2, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo de cobre com ação em reparo tecidual, remodelação de matriz e cicatrização.',
  '1% no material. É a molécula com literatura própria mais antiga deste grupo.', true, now(), now()),

 (uuid_generate_v7(), 'Haloxyl', 'HALOXYL, quimiotripsina peptídica para olheira', '%', 3, 2, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Blend peptídico para olheiras: atua sobre a hemossiderina depositada na região periorbital.',
  '3% no material, em fórmula periorbital.', true, now(), now()),

 (uuid_generate_v7(), 'TGP-2', 'TGP-2 peptídeo, Nano TGP-2', '%', 2, 1, 2, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo para região periorbital, associado ao Haloxyl no material.',
  'De 1 a 2% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Nopigmerin', 'NOPIGMERIN', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo com ação sobre a pigmentação cutânea.',
  '3% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Procapil', 'PROCAPIL, biotinil tripeptídeo-1', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Complexo com biotinil tripeptídeo para couro cabeludo, voltado à queda capilar.',
  '3% no material, em loção capilar.', true, now(), now()),

 (uuid_generate_v7(), 'Prohairin', 'PROHAIRIN, peptídeo capilar', '%', 3, 3, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo para estímulo do crescimento capilar.',
  '3% no material, associado ao Procapil.', true, now(), now()),

 (uuid_generate_v7(), 'EGF Nanofactor', 'fator de crescimento epidérmico, EGF', '%', 0.5, 0.5, 0.5, 'por_dose', 1.0, 'classe',
  'parceiro', 'pending',
  'Fator de crescimento epidérmico em veiculação nanossomal, usado em regeneração e em protocolo capilar.',
  '0,5% no material. Fator de crescimento tópico é assunto com discussão regulatória e de segurança própria: fica como pendente até conferência.', true, now(), now()),

 (uuid_generate_v7(), 'FGF Nanofactor', 'fator de crescimento de fibroblasto, FGF', '%', 0.5, 0.5, 0.5, 'por_dose', 1.0, 'classe',
  'parceiro', 'pending',
  'Fator de crescimento de fibroblasto em veiculação nanossomal.',
  '0,5% no material. Mesma ressalva do EGF.', true, now(), now()),

 (uuid_generate_v7(), 'Nano IDP-2', 'NANO IDP-2', '%', 1, 1, 1, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Peptídeo em veiculação nanossomal para reparo e rejuvenescimento.',
  '1% no material.', true, now(), now()),

 (uuid_generate_v7(), 'Vc-IP', 'tetraisopalmitato de ascorbila, VC-IP', '%', 3, 2, 3, 'por_dose', 1.0, 'classe',
  'parceiro', 'suggested',
  'Derivado lipossolúvel de vitamina C para clareamento e antioxidação cutânea.',
  '3% no material, em sérum facial. É a forma lipossolúvel da vitamina C, mais estável que o ácido ascórbico livre.', true, now(), now())
ON CONFLICT DO NOTHING;

COMMIT;
