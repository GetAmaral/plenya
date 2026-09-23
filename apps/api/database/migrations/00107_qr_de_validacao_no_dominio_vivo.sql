-- +goose Up
-- O endereço de validação gravado nas linhas antigas aponta para plenya.com.br, um domínio que
-- não responde. O host correto é app.plenyasaude.com.br, onde as três páginas públicas de
-- validação são servidas. O PDF já impresso não tem conserto (o QR está no papel), mas a coluna
-- tem: ela é devolvida pela API e vira resposta errada no instante em que alguma tela a renderizar.
UPDATE issued_documents
   SET qr_code_data = replace(qr_code_data, 'https://plenya.com.br', 'https://app.plenyasaude.com.br')
 WHERE qr_code_data LIKE 'https://plenya.com.br/%';

UPDATE prescriptions
   SET qr_code_data = replace(qr_code_data, 'https://plenya.com.br', 'https://app.plenyasaude.com.br')
 WHERE qr_code_data LIKE 'https://plenya.com.br/%';

UPDATE lab_requests
   SET qr_code_data = replace(qr_code_data, 'https://plenya.com.br', 'https://app.plenyasaude.com.br')
 WHERE qr_code_data LIKE 'https://plenya.com.br/%';

-- +goose Down
-- Sem volta: reverter recolocaria um domínio morto na coluna.
SELECT 1;
