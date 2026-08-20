-- Ajuste dos nomes gerados a partir do formulário: o formulário escreve tudo em caixa alta e
-- abreviado, e a capitalização automática produz "Bcaa" e "Lactob plantarum". Aqui os nomes ficam
-- como se escreve de verdade, e as siglas voltam a ser siglas.
--
-- O veículo sai do catálogo de ativos: "veículo oleoso qsp" é base, não substância.

BEGIN;

UPDATE magistral_components SET name = 'AAKG', synonyms = 'arginina alfa-cetoglutarato' WHERE name = 'Aakg';
UPDATE magistral_components SET name = 'BCAA', synonyms = 'aminoácidos de cadeia ramificada' WHERE name = 'Bcaa';
UPDATE magistral_components SET name = 'ID-alG', synonyms = 'ID ALG, id alg' WHERE name = 'Id alg';
UPDATE magistral_components SET name = 'Oli-Ola', synonyms = 'OLI OLA' WHERE name = 'Oli ola';
UPDATE magistral_components SET name = 'Beta-alanina' WHERE name = 'Beta alanina';
UPDATE magistral_components SET name = 'Alfa-amilase' WHERE name = 'Alfa amilase';
UPDATE magistral_components SET name = 'Gama-orizanol' WHERE name = 'Gama orizanol';
UPDATE magistral_components SET name = 'Castanha-da-índia' WHERE name = 'Castanha da india';
UPDATE magistral_components SET name = 'Fosfolipídeos de caviar' WHERE name = 'Fosfolipideos de caviar';
UPDATE magistral_components SET name = 'Papaína' WHERE name = 'Papaina';
UPDATE magistral_components SET name = 'Equinácea' WHERE name = 'Equinacea';
UPDATE magistral_components SET name = 'Ásiaticosídeo', synonyms = 'ASIATICOSIDE, asiaticoside' WHERE name = 'Asiaticoside';
UPDATE magistral_components SET name = 'Cássia angustifólia', synonyms = 'sene' WHERE name = 'Cassia angustifolia';
UPDATE magistral_components SET name = 'Romã', synonyms = 'POMEGRANATE, pomegranate, Punica granatum' WHERE name = 'Pomegranate';
UPDATE magistral_components SET name = 'Trevo-vermelho', synonyms = 'RED CLOVER, red clover, Trifolium pratense' WHERE name = 'Red clover';

-- probióticos com o gênero por extenso
UPDATE magistral_components SET name = 'Lactobacillus acidophilus', synonyms = 'LACTOB ACIDOPHILLUS' WHERE name = 'Lactob acidophillus';
UPDATE magistral_components SET name = 'Lactobacillus bulgaricus',  synonyms = 'LACTOB BULGARICUS'  WHERE name = 'Lactob bulgaricus';
UPDATE magistral_components SET name = 'Lactobacillus casei',       synonyms = 'LACTOB CASEI'       WHERE name = 'Lactob casei';
UPDATE magistral_components SET name = 'Lactobacillus delbrueckii', synonyms = 'LACTOB DELBRUECKII' WHERE name = 'Lactob delbrueckii';
UPDATE magistral_components SET name = 'Lactobacillus gasseri',     synonyms = 'LACTOB GASSERI'     WHERE name = 'Lactob gasseri';
UPDATE magistral_components SET name = 'Lactobacillus plantarum',   synonyms = 'LACTOB PLANTARUM'   WHERE name = 'Lactob plantarum';
UPDATE magistral_components SET name = 'Lactobacillus reuteri',     synonyms = 'LACTOB REUTERI'     WHERE name = 'Lactob reuteri';
UPDATE magistral_components SET name = 'Lactobacillus rhamnosus',   synonyms = 'LACTOB RHAMNOSUS'   WHERE name = 'Lactob rhamnosus';
UPDATE magistral_components SET name = 'Lactobacillus salivarius',  synonyms = 'LACTOB SALIVARUS'   WHERE name = 'Lactob salivarus';
UPDATE magistral_components SET name = 'Streptococcus thermophilus', synonyms = 'LACTOB THERMOPHILUS' WHERE name = 'Lactob thermophilus';
UPDATE magistral_components SET name = 'Bifidobacterium bifidum', synonyms = 'BIFIDOB BIFIDUM' WHERE name = 'Bifidob bifidum';
UPDATE magistral_components SET name = 'Bifidobacterium breve',  synonyms = 'BIFIDOB BREVE'  WHERE name = 'Bifidob breve';
UPDATE magistral_components SET name = 'Bifidobacterium longum', synonyms = 'BIFIDOB LONGUM' WHERE name = 'Bifidob longum';

-- tinturas
UPDATE magistral_components SET name = 'Tintura de alcachofra',       synonyms = 'TINT ALCACHOFRA'       WHERE name = 'Tint alcachofra';
UPDATE magistral_components SET name = 'Tintura de alecrim',          synonyms = 'TINT ALECRIM'          WHERE name = 'Tint alecrim';
UPDATE magistral_components SET name = 'Tintura de espinheira-santa', synonyms = 'TINT ESPINHEIRA SANTA' WHERE name = 'Tint espinheira santa';
UPDATE magistral_components SET name = 'Tintura de funcho',           synonyms = 'TINT FUNCHO'           WHERE name = 'Tint funcho';
UPDATE magistral_components SET name = 'Tintura de hortelã',          synonyms = 'TINT HORTELA'          WHERE name = 'Tint hortela';

-- veículo não é ativo
DELETE FROM magistral_components WHERE name = 'Veiculo oleoso qsp';

COMMIT;

-- Os nomes acima mudaram depois que as fórmulas já apontavam para a grafia antiga: os sinônimos
-- mantêm a ligação, senão a fórmula perde densidade e o cálculo de cápsula para de sair.
BEGIN;
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', alfa amilase')       WHERE name = 'Alfa-amilase';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', beta alanina')       WHERE name = 'Beta-alanina';
UPDATE magistral_components SET synonyms = trim(both ', ' FROM coalesce(synonyms,'') || ', castanha da india') WHERE name = 'Castanha-da-índia';
COMMIT;
