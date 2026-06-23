-- F0 — Higiene de dados das escalas clínicas (plano-escalas-pergunta-a-pergunta.md)
-- Aplicado em DEV via psql. Para PROD: rodar este mesmo arquivo (idempotente) quando autorizado.
-- NÃO é goose (é dado, não schema), seguindo o padrão do bundle AGIR/Bioma.

BEGIN;

-- 1) Epworth: níveis 1 e 2 estavam com operator '=' e SEM limites (nomes "De 15 a 17"/"De 11
--    a 14"), o que faz a classificação por total falhar em 11–17. Corrigir para 'between'.
UPDATE score_levels sl
SET operator = 'between', lower_limit = '15', upper_limit = '17'
FROM score_items si
WHERE sl.item_id = si.id
  AND si.anamnese_item_code = 'ESCALA_DE_SONOLENCIA_DE_EPWORTH_24'
  AND sl.level = 1;

UPDATE score_levels sl
SET operator = 'between', lower_limit = '11', upper_limit = '14'
FROM score_items si
WHERE sl.item_id = si.id
  AND si.anamnese_item_code = 'ESCALA_DE_SONOLENCIA_DE_EPWORTH_24'
  AND sl.level = 2;

-- 2) Atribuir anamnese_item_code estável às 2 escalas que estavam sem código (chave do
--    SCALE_REGISTRY). Só seta se ainda estiver NULL (idempotente).
UPDATE score_items
SET anamnese_item_code = 'PSQI_21'
WHERE id = 'c77cedd3-2800-722d-bb14-e3d528ce05ed' AND anamnese_item_code IS NULL;

UPDATE score_items
SET anamnese_item_code = 'FSS_9'
WHERE id = 'c77cedd3-2800-7a2e-99ee-0fba00ab9037' AND anamnese_item_code IS NULL;

-- 3) IIEF-5: ao incluir a opção "0 = sem atividade sexual" em Q2/Q3/Q5, o total pode cair
--    abaixo de 5. A banda mais baixa era 'between 5 e 7' e deixaria 2–4 sem classificar.
--    Estende para '<= 7' (total baixo = disfunção grave). As demais bandas seguem cobrindo 8–25.
UPDATE score_levels sl
SET operator = '<=', lower_limit = NULL, upper_limit = '7'
FROM score_items si
WHERE sl.item_id = si.id
  AND si.anamnese_item_code = 'IIEF_5_25'
  AND sl.level = 0;

COMMIT;

-- NOTA (Tier C, não corrigido aqui — depende da fórmula que o Dr. Getúlio validar):
--   FSS (FSS_9): nível 2 está como operator '=' lower=upper=4 (casa só média exatamente 4.0).
--     A média FSS é contínua (1–7); a faixa "borderline" precisa virar um range (ex.: 4 a <5)
--     quando definirmos a fórmula/cutoff oficial.
--   ASEX (ASEX_25): níveis classificam por regra de item ("1 item >=5 pts"), não pelo total —
--     será classify() custom no registro; os níveis ficam como estão.
