-- Substâncias do formulário das parceiras que faltavam no catálogo.
-- Nome canônico, unidade e densidade por classe. NENHUMA faixa de dose entra aqui: dose
-- vinda da própria fórmula é o que fazia o catálogo conferir a fórmula contra ela mesma.
BEGIN;
INSERT INTO magistral_components (id, name, synonyms, default_unit, bulk_density, density_source,
  eutectic_former, hygroscopic, oxidizing, oxidation_sensitive, photosensitive, sachet_ok,
  source, evidence_status, notes, is_active, created_at, updated_at) VALUES
  (uuid_generate_v7(), 'AAKG', 'arginina alfa-cetoglutarato', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Abacateiro', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Alfa-amilase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Altilix', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ásiaticosídeo', 'ASIATICOSIDE, asiaticoside', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Astragalus', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'BCAA', 'aminoácidos de cadeia ramificada', 'g', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Beta-alanina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Betacaroteno', 'BETACAROTENO', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como vitamina; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Betaína anidra', 'BETAÍNA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium bifidum', 'BIFIDOB BIFIDUM', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium breve', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bifidobacterium longum', 'BIFIDOB LONGUM', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Boro', 'BORO QUELATO', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como mineral; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Boswellia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Bromelina', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Buclizina', 'BUCLISINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Capsiate', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cássia angustifólia', 'sene', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Castanha-da-índia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cavalinha', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Centella asiatica', 'CENTELA ASIATICA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Chia', 'CHIA', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Chlorella', 'CLORELLA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ciproeptadina', 'CIPROHEPTADINE', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cissus quadrangularis', 'CISSUS QUADRANGULARIS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cisteína', 'CISTÉINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Citrus aurantium', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Coleus forskohlii', 'COLEUS FORSKOHLLI', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Colágeno tipo II', 'COLAGENO TIPO II, UC II', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cranberry', 'CRAMBERRY EXT SECO 25%', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Cyanotis vagas', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'DMAE bitartarato', 'DMAE BITARTARATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Dimpless', 'SOD DIMPLESS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Diosmina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Dong quai', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Enzimas pancreáticas', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Epicor', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Equinácea', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Extrato de semente de uva', 'EXT SEMENTE UVA, VITIS VINIFERA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Faseolamina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Folha de oliveira', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Fosfolipídeos de caviar', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Gama-orizanol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Gengibre', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ginseng', 'EXT SECO DE GINSENG', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Glisodim', 'GLISODIN', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Glutationa reduzida', 'GLUTATIONA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Griffonia', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'HMB', 'HMB CALCIO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Hesperidina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Hibisco', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'ID-alG', 'ID ALG, id alg', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Kava-kava', 'KAWA-KAWA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-arginina', 'L ARGININA, L-ARGININA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-citrulina malato', 'L CITRULINA MALATO', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-fenilalanina', 'L FENILALANINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-isoleucina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-leucina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-lisina', 'L-LISINA, LISINA', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-ornitina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-prolina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'L-valina', '', 'mg', 0.65, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como aminoácido; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus acidophilus', 'LACTOB ACIDOPHILLUS', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus bulgaricus', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus casei', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus delbrueckii', 'LACTOB DELBRUECKII', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus gasseri', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus plantarum', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus reuteri', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus rhamnosus', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lactobacillus salivarius', '', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Streptococcus thermophilus', 'LACTOB THERMOPHILUS', 'bilhões UFC', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lecitina', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lipase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Lowat', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Luteína', 'LUTEÍNA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Maca peruana', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Manganês', 'MANGANÊS QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Melissa', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Meratrim', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Molibdênio', 'MOLIBEDÊNIO QUELATO', 'mcg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Mulungu', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Nattoquinase', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Nutricolin', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Oli-Ola', 'OLI OLA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Palatinose', '', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Papaína', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Phosfator', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Picnogenol', 'PICNOGENOOL', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Polidextrose', '', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Polypodium leucotomos', 'POLYPODIUM LEUCOTOMOS', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Romã', 'POMEGRANATE, pomegranate, Punica granatum', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Potássio', 'CLORETO DE POTÁSSIO, POTÁSSIO QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Pregnenolona', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Protease', '', 'mg', 0.5, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como enzima; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Psyllium', 'PLANTAGO OVATE, PSILLYUM', 'g', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Pygeum africanum', 'PYGEUM AFRICANUM', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Trevo-vermelho', 'RED CLOVER, red clover, Trifolium pratense', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Rutina', 'RUTINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Sinetrol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Sucupira', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de alcachofra', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de alecrim', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de espinheira-santa', 'TINT ESPINHEIRA SANTA', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de funcho', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tintura de hortelã', '', 'ml', 1.0, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como tintura; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Tribulus terrestris', 'TRIBULUS TERRESTRE', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vanádio', 'VANADIO QUELADO, VANÁDIO QUELATO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Veiculo oleoso qsp', '', 'Gotas', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Verisol', '', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vinpocetina', 'VIMPOCETINA', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Vitamina B1', 'TIAMINA, VIT B1', 'mg', 0.75, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como vitamina; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ácido D-aspártico', 'AC D ASPARTICO', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
  (uuid_generate_v7(), 'Ácido pantotênico', 'PANTOTENATO DE CÁLCIO, VIT B5', 'mg', 0.45, 'classe', false, false, false, false, false, true, 'parceiro', 'pending', 'Classificada como fitoterápico ou ativo patenteado; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now())
ON CONFLICT DO NOTHING;

UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AMORA VERDE EXT SECO')
 WHERE name = 'Amora';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ASWAGANDHA')
 WHERE name = 'Ashwagandha';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', COBRE QUELADO, COBRE QUELATO')
 WHERE name = 'Cobre';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', COE Q10, COENZIMA Q0')
 WHERE name = 'Coenzima Q10';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', CALCIO QUELADO, CÁLCIO CITRATO, CÁLCIO QUELATO')
 WHERE name = 'Cálcio';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', D CHIRO INOSITOL')
 WHERE name = 'D-quiro-inositol';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', D RIBOSE')
 WHERE name = 'D-ribose';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ACTIVE EGCG, CHA VERDE')
 WHERE name = 'Extrato de chá verde';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', FOSFATIDIL COLINA')
 WHERE name = 'Fosfatidilcolina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L-GLICINA')
 WHERE name = 'Glicina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GREEN COFFE')
 WHERE name = 'Green coffee';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GINMENA SILVESTRE, GYMENA SILVESTRE')
 WHERE name = 'Gymnema silvestre';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L GLUTAMINA')
 WHERE name = 'L-glutamina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L THEANINA, L THEANINE, THEANINA')
 WHERE name = 'L-teanina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', L TIROSINA')
 WHERE name = 'L-tirosina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', MAGNESIO QUELADO, MAGNÉSIO')
 WHERE name = 'Magnésio quelato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT B12')
 WHERE name = 'Metilcobalamina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AC FOLICO, LEVOFOLIC ACID, ÁC FOLICO, ÁC FÓLICO, ÁC. FOLICO')
 WHERE name = 'Metilfolato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', MIO INOSITOL')
 WHERE name = 'Mio-inositol';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', N ACETIL CISTEINA')
 WHERE name = 'N-acetilcisteína';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', NIACINA, VIT B3')
 WHERE name = 'Nicotinamida';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT C')
 WHERE name = 'Palmitato de ascorbila';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', CROMO QUELATO, PICOLINATO CROMO')
 WHERE name = 'Picolinato de cromo';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', PIRIDOXAL 5 FOSFATO, PIRIDOXAL 5-FOSFATO, PIRIDOXAL FOSFATO, VIT B6, VITB6')
 WHERE name = 'Piridoxal-5-fosfato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT B2')
 WHERE name = 'Riboflavina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', EXSELEN, SELENIO QUELADO, SELÊNIO QUELATO')
 WHERE name = 'Selenometionina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', TEACRINE')
 WHERE name = 'Teacrina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT A')
 WHERE name = 'Vitamina A';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT D, VIT D 3, VIT D3')
 WHERE name = 'Vitamina D3';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT E')
 WHERE name = 'Vitamina E';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', VIT K, VIT K2 MK7')
 WHERE name = 'Vitamina K2 MK-7';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', ZINCO, ZINCO QUELADO')
 WHERE name = 'Zinco quelato';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', AC ALFA LIPOICO, ALFA LIPOICO, ÁC ALFA LIPOICO')
 WHERE name = 'Ácido alfa-lipoico';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', GARCINEA CAMBOJA, GARCINIA CAMBOJA')
 WHERE name = 'Ácido hidroxicítrico';
COMMIT;

-- Segunda passada do parser: componentes que só apareceram depois de aceitar a dose separada por
-- um espaço só ("STREPTOCOCCUS THERMOPHILUS 1 BLH") e o prefixo "POSOLOGIA:" na linha de uso.
BEGIN;
INSERT INTO magistral_components (id, name, synonyms, default_unit, bulk_density, density_source,
  sachet_ok, source, evidence_status, notes, is_active, created_at, updated_at) VALUES
 (uuid_generate_v7(), 'Bifidobacterium lactis', 'BIFIDOBACTERIUM LACTIS', 'bilhões UFC', 0.45, 'classe', true,
  'parceiro', 'pending', 'Probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Lactobacillus paracasei', 'LACTOBACILLUS PARACASEI', 'bilhões UFC', 0.45, 'classe', true,
  'parceiro', 'pending', 'Probiótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Polietilenoglicol 4000', 'PEG 4000, macrogol 4000', 'g', 0.60, 'classe', true,
  'parceiro', 'pending', 'Laxante osmótico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now()),
 (uuid_generate_v7(), 'Ganoderma lucidum', 'WEG LEM 70, reishi, cogumelo do sol', 'mg', 0.45, 'classe', true,
  'parceiro', 'pending', 'Fitoterápico; densidade aproximada por classe. Sem faixa de dose até conferência.', true, now(), now())
ON CONFLICT DO NOTHING;
COMMIT;
