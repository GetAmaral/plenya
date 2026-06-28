-- Remove o resíduo de formulário de papel "____/NN" do NOME das escalas (Tier A).
-- Motivo: o widget de escala já calcula e mostra o total no cabeçalho do item (total/max,
-- com max vindo do SCALE_REGISTRY), e o cabeçalho agora exibe todos os níveis com o nível
-- certo auto-marcado. O "____/NN" no nome era uma duplicata manual e, no caso do ASEX, estava
-- inconsistente (prod dizia /25, dev /30). A literatura do ASEX é 1–6, máx 30 (McGahuey 2000),
-- então o max correto vem do registry (=30); tirar o número do nome elimina a divergência.
-- Afeta PHQ-9, GAD-7, Epworth, IIEF-5, Dubois (imediato/tardio), Span (direto/inverso), ASEX.
-- Idempotente: o WHERE não casa mais depois de remover o sufixo.

-- +goose Up
UPDATE score_items
SET name = regexp_replace(name, '\s*:?\s*_{2,}\s*/\s*[0-9]+\s*$', ''),
    updated_at = now()
WHERE name ~ '_{2,}\s*/\s*[0-9]+\s*$' AND deleted_at IS NULL;

-- +goose Down
-- Cosmético e não reversível com fidelidade (o formato e o número originais por escala não
-- foram snapshotados aqui; o max canônico vive no SCALE_REGISTRY). Restaurar do backup se
-- realmente necessário. No-op proposital.
SELECT 1;
