-- +goose Up
-- Ultrassonografia com Doppler colorido de tireoide, ao lado da USG nos templates de pedido.
--
-- Acrescenta, não substitui: no TUSS são dois procedimentos distintos, e a operadora cobra assim.
--
-- CÓDIGO TUSS: 40901386 — "Doppler colorido de órgão ou estrutura isolada"
--   (Tabela 22, grupo Métodos diagnósticos por imagem, subgrupo Ultrassonografia Diagnóstica),
--   conferido em duas fontes independentes. Três coisas que enganam aqui:
--     - NÃO existe código TUSS específico para "Doppler de tireoide". O que existe é o genérico
--       de órgão isolado, acima.
--     - 40901378 parece candidato pela vizinhança numérica e não é: é "Doppler colorido de vasos
--       cervicais VENOSOS bilateral (subclávias e jugulares)". É pescoço, não tireoide.
--     - 40901203, que o catálogo chama de "Ultrassonografia de tireoide", é oficialmente
--       "Órgãos superficiais (tireóide ou escroto ou pênis ou crânio)" e NÃO inclui o Doppler.
--
-- Idempotente de propósito: em dev o dado já existe (foi aplicado à mão antes desta migration
-- existir), e sem as guardas o `up` deslocaria a ordem dos itens seguintes um degrau de novo.

-- ---------------------------------------------------------------------------
-- 1. A definição do exame. `id` fixo para dev e prod terem o MESMO id.
INSERT INTO lab_test_definitions (
    id, code, name, short_name, tuss_code, category, result_type,
    is_requestable, is_active, description, sex_applicability,
    display_order, created_at, updated_at
)
SELECT '01a066ce-a5d4-7c16-bbc7-fe08246515d8'::uuid,
       'PLNUSGTIRDOP01',
       'Ultrassonografia com Doppler colorido de tireoide',
       'USG Doppler Tireoide',
       '40901386',
       'imaging',
       'text',
       true, true,
       'Doppler colorido de órgão isolado aplicado à tireoide. Pedido junto da USG de tireoide (TUSS 40901203), não no lugar dela: no TUSS são dois procedimentos. Acrescenta a leitura da vascularização, que é o que separa tireoidite com hiperfluxo de nódulo com padrão vascular suspeito.',
       'all',
       0, NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM lab_test_definitions d WHERE d.code = 'PLNUSGTIRDOP01');

-- ---------------------------------------------------------------------------
-- 2. Abre a vaga logo depois da USG de tireoide, em cada template que a pede.
--
-- Roda ANTES da inserção e só nos templates que ainda não têm o Doppler — é o que impede o
-- deslocamento de acontecer duas vezes.
WITH alvo AS (
    SELECT tt.lab_request_template_id AS tpl, tt.display_order AS ord
    FROM lab_request_template_tests tt
    JOIN lab_test_definitions d ON d.id = tt.lab_test_definition_id
    WHERE d.code = 'PLNUSGTIR01'
      AND NOT EXISTS (
          SELECT 1 FROM lab_request_template_tests x
          JOIN lab_test_definitions xd ON xd.id = x.lab_test_definition_id
          WHERE x.lab_request_template_id = tt.lab_request_template_id
            AND xd.code = 'PLNUSGTIRDOP01'
      )
)
UPDATE lab_request_template_tests t
   SET display_order = t.display_order + 1
  FROM alvo a
 WHERE t.lab_request_template_id = a.tpl
   AND t.display_order > a.ord;

-- ---------------------------------------------------------------------------
-- 3. O Doppler na vaga aberta.
--
-- `page_break_before = false`: quem ABRE o bloco de imagem é a USG de tireoide (é ela que carrega
-- a quebra), e o Doppler pertence à mesma página do pedido. A ordem e a paginação do pedido vêm
-- destes dois campos, não de regra no código — ver o comentário em
-- `apps/web/lib/lab-request-apply.ts`.
INSERT INTO lab_request_template_tests (
    lab_request_template_id, lab_test_definition_id, display_order, page_break_before, created_at
)
SELECT tt.lab_request_template_id,
       '01a066ce-a5d4-7c16-bbc7-fe08246515d8'::uuid,
       tt.display_order + 1,
       false,
       NOW()
FROM lab_request_template_tests tt
JOIN lab_test_definitions d ON d.id = tt.lab_test_definition_id
WHERE d.code = 'PLNUSGTIR01'
ON CONFLICT (lab_request_template_id, lab_test_definition_id) DO NOTHING;

-- +goose Down
-- Desfaz na ordem inversa: tira o Doppler dos templates, fecha a vaga que ele ocupava e só então
-- apaga a definição (a FK dos templates aponta para ela).
WITH alvo AS (
    SELECT tt.lab_request_template_id AS tpl, tt.display_order AS ord
    FROM lab_request_template_tests tt
    JOIN lab_test_definitions d ON d.id = tt.lab_test_definition_id
    WHERE d.code = 'PLNUSGTIRDOP01'
)
DELETE FROM lab_request_template_tests t
 USING alvo a
 WHERE t.lab_request_template_id = a.tpl
   AND t.lab_test_definition_id = '01a066ce-a5d4-7c16-bbc7-fe08246515d8'::uuid;

WITH alvo AS (
    SELECT tt.lab_request_template_id AS tpl, tt.display_order AS ord
    FROM lab_request_template_tests tt
    JOIN lab_test_definitions d ON d.id = tt.lab_test_definition_id
    WHERE d.code = 'PLNUSGTIR01'
)
UPDATE lab_request_template_tests t
   SET display_order = t.display_order - 1
  FROM alvo a
 WHERE t.lab_request_template_id = a.tpl
   AND t.display_order > a.ord + 1;

DELETE FROM lab_test_definitions WHERE code = 'PLNUSGTIRDOP01';
