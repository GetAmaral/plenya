-- Hardening do usuário-bot "Assistente Plenya": troca a role 'secretary' (que é um grant
-- de autorização e o fazia aparecer em listas/notificações de staff) por uma role própria
-- sem privilégio ('bot'). O bot só é usado como ActorUserID das mensagens da IA; não loga
-- (sem password_hash) e não precisa de nenhuma permissão. Idempotente.

-- +goose Up
-- +goose StatementBegin
UPDATE public.users SET roles = '["bot"]'::jsonb
WHERE id = '019e9301-0000-7000-a000-0000000b0b01';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
UPDATE public.users SET roles = '["secretary"]'::jsonb
WHERE id = '019e9301-0000-7000-a000-0000000b0b01';
-- +goose StatementEnd
