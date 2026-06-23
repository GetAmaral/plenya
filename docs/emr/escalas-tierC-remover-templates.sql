-- Tier C fora por ora: remove FSS, FSFI e PSQI dos templates de anamnese (não serão usadas).
-- ASEX permanece. Os score_items continuam existindo no escore; só saem do preenchimento.
-- Idempotente. Aplicado em DEV; rodar em PROD quando autorizado.

DELETE FROM anamnesis_template_items ati
USING score_items si
WHERE ati.score_item_id = si.id
  AND si.anamnese_item_code IN ('FSS_9', 'FSFI_36', 'PSQI_21');
