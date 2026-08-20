-- +goose Up
-- Substâncias cuja dose se prescreve em ELEMENTO, não em massa do insumo.
--
-- Mineral se prescreve assim: "cobre 1 mg" quer dizer 1 mg de cobre elementar, e a farmácia pesa
-- a massa de quelato que entrega isso. Até aqui o formulário só marcava a dose como elementar
-- quando o catálogo tinha `elemental_percent` preenchido — e justamente os quelatos de cobre,
-- boro, manganês e vanádio não têm, porque o percentual muda de fornecedor. Resultado: mineral
-- sem percentual entrava na receita como se a dose fosse do pó.
--
-- O fato "esta substância se prescreve em elemento" é da substância e não depende de sabermos o
-- percentual. Por isso vira coluna própria, em vez de continuar deduzido de `elemental_percent`.

ALTER TABLE public.magistral_components
  ADD COLUMN IF NOT EXISTS dose_as_elemental boolean NOT NULL DEFAULT false;

COMMENT ON COLUMN public.magistral_components.dose_as_elemental IS
  'Dose desta substância se prescreve em elemento (minerais). A conversão para massa de insumo usa elemental_percent quando existir.';

-- Minerais do catálogo. A lista é por nutriente da IN 28, que já classifica o que é mineral, mais
-- os nomes que ainda não têm mapeamento de teto.
UPDATE public.magistral_components SET dose_as_elemental = true
 WHERE deleted_at IS NULL
   AND (in28_nutrient IN ('Cálcio','Cobre','Cromo','Ferro','Fósforo','Iodo','Magnésio','Manganês',
                          'Molibdênio','Potássio','Selênio','Zinco','Boro','Vanádio')
        OR name ~* '^(cobre|cálcio|ferro|manganês|molibdênio|potássio|boro|vanádio|zinco|magnésio|selênio|cromo|iodo)( |$)'
        OR name ~* '(quelato|quelado|bisglicinato)$'
        OR name IN ('Iodeto de potássio','Selenometionina','Picolinato de cromo','Zinco carnosina'));

-- +goose Down
ALTER TABLE public.magistral_components DROP COLUMN IF EXISTS dose_as_elemental;
