-- A guarda "AND NOT EXISTS" em cada renomeação existe porque este arquivo roda depois de um seed
-- que insere os nomes ANTIGOS: numa segunda passada o insert recria "Aakg" e a renomeação para
-- "AAKG" colidia com a linha que já estava lá. Com a guarda, a segunda passada não faz nada.
--
-- Ajuste dos nomes gerados a partir do formulário: o formulário escreve tudo em caixa alta e
-- abreviado, e a capitalização automática produz "Bcaa" e "Lactob plantarum". Aqui os nomes ficam
-- como se escreve de verdade, e as siglas voltam a ser siglas.
--
-- O veículo sai do catálogo de ativos: "veículo oleoso qsp" é base, não substância.

BEGIN;

UPDATE magistral_components SET name = 'AAKG', synonyms = 'arginina alfa-cetoglutarato' WHERE name = 'Aakg'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'AAKG');
UPDATE magistral_components SET name = 'BCAA', synonyms = 'aminoácidos de cadeia ramificada' WHERE name = 'Bcaa'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'BCAA');
UPDATE magistral_components SET name = 'ID-alG', synonyms = 'ID ALG, id alg' WHERE name = 'Id alg'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'ID-alG');
UPDATE magistral_components SET name = 'Oli-Ola', synonyms = 'OLI OLA' WHERE name = 'Oli ola'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Oli-Ola');
UPDATE magistral_components SET name = 'Beta-alanina' WHERE name = 'Beta alanina'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Beta-alanina');
UPDATE magistral_components SET name = 'Alfa-amilase' WHERE name = 'Alfa amilase'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Alfa-amilase');
UPDATE magistral_components SET name = 'Gama-orizanol' WHERE name = 'Gama orizanol'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Gama-orizanol');
UPDATE magistral_components SET name = 'Castanha-da-índia' WHERE name = 'Castanha da india'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Castanha-da-índia');
UPDATE magistral_components SET name = 'Fosfolipídeos de caviar' WHERE name = 'Fosfolipideos de caviar'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Fosfolipídeos de caviar');
UPDATE magistral_components SET name = 'Papaína' WHERE name = 'Papaina'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Papaína');
UPDATE magistral_components SET name = 'Equinácea' WHERE name = 'Equinacea'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Equinácea');
UPDATE magistral_components SET name = 'Ásiaticosídeo', synonyms = 'ASIATICOSIDE, asiaticoside' WHERE name = 'Asiaticoside'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Ásiaticosídeo');
UPDATE magistral_components SET name = 'Cássia angustifólia', synonyms = 'sene' WHERE name = 'Cassia angustifolia'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Cássia angustifólia');
UPDATE magistral_components SET name = 'Romã', synonyms = 'POMEGRANATE, pomegranate, Punica granatum' WHERE name = 'Pomegranate'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Romã');
UPDATE magistral_components SET name = 'Trevo-vermelho', synonyms = 'RED CLOVER, red clover, Trifolium pratense' WHERE name = 'Red clover'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Trevo-vermelho');

-- probióticos com o gênero por extenso
UPDATE magistral_components SET name = 'Lactobacillus acidophilus', synonyms = 'LACTOB ACIDOPHILLUS' WHERE name = 'Lactob acidophillus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus acidophilus');
UPDATE magistral_components SET name = 'Lactobacillus bulgaricus',  synonyms = 'LACTOB BULGARICUS'  WHERE name = 'Lactob bulgaricus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus bulgaricus');
UPDATE magistral_components SET name = 'Lactobacillus casei',       synonyms = 'LACTOB CASEI'       WHERE name = 'Lactob casei'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus casei');
UPDATE magistral_components SET name = 'Lactobacillus delbrueckii', synonyms = 'LACTOB DELBRUECKII' WHERE name = 'Lactob delbrueckii'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus delbrueckii');
UPDATE magistral_components SET name = 'Lactobacillus gasseri',     synonyms = 'LACTOB GASSERI'     WHERE name = 'Lactob gasseri'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus gasseri');
UPDATE magistral_components SET name = 'Lactobacillus plantarum',   synonyms = 'LACTOB PLANTARUM'   WHERE name = 'Lactob plantarum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus plantarum');
UPDATE magistral_components SET name = 'Lactobacillus reuteri',     synonyms = 'LACTOB REUTERI'     WHERE name = 'Lactob reuteri'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus reuteri');
UPDATE magistral_components SET name = 'Lactobacillus rhamnosus',   synonyms = 'LACTOB RHAMNOSUS'   WHERE name = 'Lactob rhamnosus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus rhamnosus');
UPDATE magistral_components SET name = 'Lactobacillus salivarius',  synonyms = 'LACTOB SALIVARUS'   WHERE name = 'Lactob salivarus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Lactobacillus salivarius');
UPDATE magistral_components SET name = 'Streptococcus thermophilus', synonyms = 'LACTOB THERMOPHILUS' WHERE name = 'Lactob thermophilus'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Streptococcus thermophilus');
UPDATE magistral_components SET name = 'Bifidobacterium bifidum', synonyms = 'BIFIDOB BIFIDUM' WHERE name = 'Bifidob bifidum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium bifidum');
UPDATE magistral_components SET name = 'Bifidobacterium breve',  synonyms = 'BIFIDOB BREVE'  WHERE name = 'Bifidob breve'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium breve');
UPDATE magistral_components SET name = 'Bifidobacterium longum', synonyms = 'BIFIDOB LONGUM' WHERE name = 'Bifidob longum'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Bifidobacterium longum');

-- tinturas
UPDATE magistral_components SET name = 'Tintura de alcachofra',       synonyms = 'TINT ALCACHOFRA'       WHERE name = 'Tint alcachofra'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de alcachofra');
UPDATE magistral_components SET name = 'Tintura de alecrim',          synonyms = 'TINT ALECRIM'          WHERE name = 'Tint alecrim'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de alecrim');
UPDATE magistral_components SET name = 'Tintura de espinheira-santa', synonyms = 'TINT ESPINHEIRA SANTA' WHERE name = 'Tint espinheira santa'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de espinheira-santa');
UPDATE magistral_components SET name = 'Tintura de funcho',           synonyms = 'TINT FUNCHO'           WHERE name = 'Tint funcho'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de funcho');
UPDATE magistral_components SET name = 'Tintura de hortelã',          synonyms = 'TINT HORTELA'          WHERE name = 'Tint hortela'
   AND NOT EXISTS (SELECT 1 FROM magistral_components m WHERE m.name = 'Tintura de hortelã');

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
