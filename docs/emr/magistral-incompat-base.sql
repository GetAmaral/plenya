-- Incompatibilidades com a base, todas com mecanismo e fonte.
--
-- Fonte principal: Alves FC, Passos MMB, Melo ASP, Monteiro MSSB. Perfil dos erros de prescrições
-- de medicamentos manipulados em uma farmácia-escola. Vigil. sanit. debate 2019;7(1):5-13
-- (400 prescrições, farmácia-escola da UFRJ) — que descreve caso a caso o mecanismo de cada uma.
-- Complemento: Formulário Nacional da Farmacopeia Brasileira, 2ª ed.
--
-- Idempotente pelo par (base, substância, percentual).

BEGIN;

DELETE FROM magistral_base_incompatibilities
 WHERE source LIKE 'Vigil. sanit. debate%' OR source LIKE 'Formulário Nacional%';

INSERT INTO magistral_base_incompatibilities
    (base_pattern, substance_pattern, min_percent, severity, mechanism, recommendation, source)
VALUES
 ('lanette', 'ácido', 10, 'warn',
  'O creme Lanette é uma emulsão aniônica: acima de mais ou menos 10%, o ativo ácido neutraliza as cargas do tensoativo e a emulsão se separa.',
  'Creme não iônico comporta ativo com carga em concentração alta.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'ureia', 30, 'warn',
  'A ureia é muito hidrossolúvel: em concentração alta ela migra para a fase aquosa e muda a proporção entre as fases, desestabilizando a emulsão.',
  'Creme não iônico, ou dividir a ureia em outra preparação.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'lactato de amônio', NULL, 'warn',
  'Sal ionizável: interage com o tensoativo aniônico do Lanette em qualquer concentração.',
  'Creme não iônico.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('lanette', 'PCA', NULL, 'warn',
  'O PCA de sódio é sal do ácido pirrolidona carboxílico: ioniza em água e interage com o tensoativo aniônico em qualquer concentração.',
  'Creme não iônico.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('vaselina', 'LCD', NULL, 'avoid',
  'O LCD (licor carbonis detergens) é polar e a vaselina sólida é apolar: o ativo não incorpora nem homogeneíza.',
  'Pomada base (vaselina sólida com lanolina 7:3), que tem parte hidrofílica.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('diadermina', 'ácido', NULL, 'warn',
  'A base de diadermina é parcialmente saponificada com agente alcalino: o ativo ácido neutraliza o emulsionante, a viscosidade cai e a emulsão separa.',
  'Base não iônica.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('não iônico', 'hidroquinona', NULL, 'warn',
  'A hidroquinona oxida com facilidade e a associação com creme não iônico desestabiliza na prática. O artigo registra a observação e diz que não há estudo explicando o mecanismo.',
  'Creme Lanette para a hidroquinona; se a fórmula também tiver ácidos, separar em duas preparações.',
  'Vigil. sanit. debate 2019;7(1):5-13'),

 ('xarope', 'cloreto de potássio', 6, 'avoid',
  'Em xarope o açúcar já está perto da saturação e sobra pouca água livre: acima de 6% o cloreto de potássio não solubiliza e a dose por tomada passa a variar.',
  'Manter em 6%, que é o limite do Formulário Nacional.',
  'Formulário Nacional da Farmacopeia Brasileira, 2ª ed.');

COMMIT;

-- Par ativo x ativo documentado na mesma fonte.
BEGIN;
WITH a AS (SELECT id FROM magistral_components WHERE name = 'Cetoconazol'),
     b AS (SELECT id FROM magistral_components WHERE name = 'Ureia')
INSERT INTO magistral_incompatibilities (id, component_a_id, component_b_id, severity, mechanism, note)
SELECT uuid_generate_v7(), a.id, b.id, 'warn',
       'a associação muda de cor na prática, o que sugere queda de teor ou formação de outro composto',
       'A farmácia-escola da UFRJ deixou de manipular os dois juntos por causa das reclamações de cor. O artigo registra que a literatura não descreve o mecanismo.'
  FROM a, b
 WHERE NOT EXISTS (SELECT 1 FROM magistral_incompatibilities i WHERE i.component_a_id=a.id AND i.component_b_id=b.id);
COMMIT;
