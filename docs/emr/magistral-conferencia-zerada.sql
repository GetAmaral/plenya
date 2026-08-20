-- O carimbo de conferência das fórmulas-base estava mentindo.
--
-- Os seeds gravaram last_review junto com a fórmula, então as 24 apareciam como revisadas sem que
-- ninguém tivesse olhado. Marca de conferência que nasce preenchida não é marca de conferência: é
-- ruído esperando para enganar alguém no dia em que a conferência de verdade começar.
--
-- Fica NULO até um humano salvar a fórmula pela tela, que é o único caminho que carimba.

BEGIN;

UPDATE magistral_formula_templates
   SET last_review = NULL, reviewed_by = NULL
 WHERE deleted_at IS NULL;

COMMIT;
