-- Indicações e posologia em TÓPICOS, ao lado do texto corrido.
--
-- PORQUÊ: o texto extraído do material é bom para conferir, e ruim para decidir. Num painel de
-- prescrição o médico quer bater o olho em "para quem" e "quanto", não ler três frases. Os dois
-- formatos convivem: tópico em cima, texto completo a um clique.
--
-- FORMATO: uma linha por tópico, texto puro. NÃO é jsonb de propósito — a migration 00060
-- documenta uma linha jsonb malformada derrubando a leitura de uma tabela inteira pelo
-- serializer do GORM, e aqui não há nada que justifique correr esse risco: o campo é editado à
-- mão num textarea, onde uma linha por item é exatamente o que o médico já escreveria.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD COLUMN IF NOT EXISTS indication_bullets text,
    ADD COLUMN IF NOT EXISTS dose_bullets       text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    DROP COLUMN IF EXISTS indication_bullets,
    DROP COLUMN IF EXISTS dose_bullets;
-- +goose StatementEnd
