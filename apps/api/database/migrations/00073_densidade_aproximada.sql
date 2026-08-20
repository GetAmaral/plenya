-- Procedência da densidade aparente.
--
-- PORQUÊ: a calculadora de cápsula ficou muda porque densidade aparente não existe em base
-- pública e varia por lote. A decisão agora é outra: usar valores APROXIMADOS para a calculadora
-- opinar, desde que a tela diga que são aproximados. Estimativa declarada é útil; estimativa
-- disfarçada de medida é que seria o problema.
--
-- `density_source` guarda de onde veio o número:
--   · 'medida'  — valor publicado para aquela substância específica
--   · 'classe'  — aproximação pela classe do pó (mineral quelado, extrato seco, aminoácido…)
--   · texto livre — ficha técnica do insumo, quando a farmácia informar
-- Sem valor em `bulk_density`, o campo fica NULL e a calculadora continua se calando.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD COLUMN IF NOT EXISTS density_source varchar(60);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP COLUMN IF EXISTS density_source;
-- +goose StatementEnd
