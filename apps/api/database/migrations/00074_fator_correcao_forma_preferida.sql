-- Fator de correção e forma preferida — as duas lições das fichas técnicas das farmácias parceiras.
--
-- FATOR DE CORREÇÃO. Toda ficha técnica de insumo traz o campo "Fator de Correção", e ele existe
-- porque prescrever "Magnésio quelato 300 mg" é AMBÍGUO: 300 mg de magnésio elementar equivalem a
-- 1 g de bisglicinato 30%, e 300 mg do quelato equivalem a 90 mg de elementar. A diferença é de
-- mais de três vezes, e hoje o sistema não sabia distinguir. Três consequências:
--   · a receita passa a dizer se a dose é do ELEMENTO ou do INSUMO;
--   · a tela mostra a conversão, com o percentual do laudo;
--   · a calculadora de cápsula passa a usar a massa do INSUMO, que é o pó que ocupa volume. Sem
--     isto ela subestimava a cápsula em fórmula com quelato ou diluído a 1%.
-- Selenometionina 1% e bisglicinato 30% são os exemplos documentados nas fichas.
--
-- FORMA PREFERIDA. O médico usa sempre metilcobalamina, metilfolato, piridoxal-5-fosfato,
-- palmitato de ascorbila e selenometionina quando precisa dessas vitaminas e minerais, e CavaQ10
-- para coenzima Q10. O catálogo passa a saber disso: a substância genérica aponta para a forma
-- que ele prescreve, e a tela sugere a troca em vez de depender da memória.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD COLUMN IF NOT EXISTS elemental_percent        numeric(7,3),
    ADD COLUMN IF NOT EXISTS correction_note          varchar(300),
    ADD COLUMN IF NOT EXISTS preferred_alternative_id uuid,
    ADD COLUMN IF NOT EXISTS preference_note          varchar(300),
    ADD COLUMN IF NOT EXISTS brand                    varchar(60);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS fk_magistral_components_preferred;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD CONSTRAINT fk_magistral_components_preferred
    FOREIGN KEY (preferred_alternative_id) REFERENCES public.magistral_components(id);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS chk_magistral_components_elemental;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD CONSTRAINT chk_magistral_components_elemental
    CHECK (elemental_percent IS NULL OR (elemental_percent > 0 AND elemental_percent <= 100));
-- +goose StatementEnd

-- A dose escrita é do elemento ou do insumo? Sem esta resposta, a fórmula é ambígua para a
-- farmácia e para a calculadora.
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components
    ADD COLUMN IF NOT EXISTS as_elemental boolean NOT NULL DEFAULT false;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_formula_template_components
    ADD COLUMN IF NOT EXISTS as_elemental boolean NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.magistral_formula_template_components DROP COLUMN IF EXISTS as_elemental;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components DROP COLUMN IF EXISTS as_elemental;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS fk_magistral_components_preferred;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS chk_magistral_components_elemental;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    DROP COLUMN IF EXISTS elemental_percent,
    DROP COLUMN IF EXISTS correction_note,
    DROP COLUMN IF EXISTS preferred_alternative_id,
    DROP COLUMN IF EXISTS preference_note,
    DROP COLUMN IF EXISTS brand;
-- +goose StatementEnd
