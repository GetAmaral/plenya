-- Fechamento de extremos abertos + re-base por literatura em 16 ScoreItems.
-- Auditoria estrutural (gaps/overlaps/extremos): 12 itens "Tipo 1" (estende banda da ponta /
-- re-base com âncora clínica), Hora de dormir (mantido, só piso) e 3 itens "Tipo 2" com ponta
-- patológica viraram U com níveis distintos (Magnésio Sérico, USG Volume Ovariano, pO2 Venoso).
-- Decisões clínicas item-a-item validadas pelo usuário; fontes (NHANES BRI, SES-CD, consenso
-- breath test, ref. funcionais de testosterona/magnésio, UK Biobank hora de dormir).
-- Idempotente por id (UPDATE) e por id+ON CONFLICT (os 2 níveis novos).
-- Itens: IgE, Alfa-2 Globulina, BRI, SES-CD, H₂/CH₄ basal, FSH Lútea, Magnésio RBC,
-- Progesterona Gestação 1º/2º/3º, Testosterona Total Homens, Hora de dormir, Magnésio Sérico,
-- USG Volume Ovariano, pO2 Venoso.

-- +goose Up

-- Alfa-2 Globulina (re-base 0,8/1,0/1,4)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='0.8', name='≤0,8', updated_at=NOW() WHERE id='019bf31d-2ef0-77b2-a9d0-c88174357041';
UPDATE score_levels SET level=3, operator='between', lower_limit='0.8', upper_limit='1.0', name='0,8-1,0', updated_at=NOW() WHERE id='019bf31d-2ef0-784c-8e84-13bb84b16a6d';
UPDATE score_levels SET level=1, operator='between', lower_limit='1.0', upper_limit='1.4', name='1,0-1,4', updated_at=NOW() WHERE id='019bf31d-2ef0-7ef4-ae2f-3f3e00fa8093';
UPDATE score_levels SET level=0, operator='>', lower_limit='1.4', upper_limit=NULL, name='>1,4', updated_at=NOW() WHERE id='019bf31d-2ef0-7e7a-bffb-535dd8a02420';

-- BRI (U assimétrico, quintis NHANES; level 1 -> 2)
UPDATE score_levels SET level=5, operator='between', lower_limit='4.5', upper_limit='5.5', name='4,5-5,5', updated_at=NOW() WHERE id='019cbe95-820f-787a-9d95-b4417d6aaf9a';
UPDATE score_levels SET level=4, operator='between', lower_limit='3.4', upper_limit='4.5', name='3,4-4,5', updated_at=NOW() WHERE id='019cbe95-820f-7480-b304-5ae51ec1d4b5';
UPDATE score_levels SET level=3, operator='between', lower_limit='5.5', upper_limit='6.9', name='5,5-6,9', updated_at=NOW() WHERE id='019cbe95-820f-7ec9-9ffd-2c9c54165289';
UPDATE score_levels SET level=2, operator='<', lower_limit=NULL, upper_limit='3.4', name='<3,4', updated_at=NOW() WHERE id='019cbe95-820f-7547-b500-dced9053b379';
UPDATE score_levels SET level=0, operator='>=', lower_limit='6.9', upper_limit=NULL, name='≥6,9', updated_at=NOW() WHERE id='019cbe95-820f-7bc3-b081-15ea4ef234f9';

-- Colonoscopia - SES-CD Crohn (realinhado 0-2/3-6/7-15/≥16)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='2', name='0-2 (Remissão)', updated_at=NOW() WHERE id='019bf7af-ba01-7040-b370-d7a053da3dff';
UPDATE score_levels SET level=4, operator='between', lower_limit='2', upper_limit='6', name='3-6 (Leve)', updated_at=NOW() WHERE id='019bf7af-ba01-7cd1-aae7-44f8c22119ae';
UPDATE score_levels SET level=3, operator='between', lower_limit='6', upper_limit='11', name='7-11 (Moderado baixo)', updated_at=NOW() WHERE id='019bf7af-ba01-795d-af59-a934ea01fbae';
UPDATE score_levels SET level=2, operator='between', lower_limit='11', upper_limit='15', name='12-15 (Moderado alto)', updated_at=NOW() WHERE id='019bf7af-ba01-7d48-b220-c78308fe2c80';
UPDATE score_levels SET level=1, operator='between', lower_limit='15', upper_limit='30', name='16-30 (Severo)', updated_at=NOW() WHERE id='019bf7af-ba01-71dd-b723-642e4f02b527';
UPDATE score_levels SET level=0, operator='>', lower_limit='30', upper_limit=NULL, name='>30 (Muito severo)', updated_at=NOW() WHERE id='019bf7af-ba01-77b5-985d-dd1e6244c22d';

-- FSH - Mulheres Fase Lútea (só piso fechado)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='5', name='1-5', updated_at=NOW() WHERE id='019bf31d-2ef0-763d-8021-0fb43eeedfec';
UPDATE score_levels SET level=3, operator='between', lower_limit='5', upper_limit='9', name='5-9', updated_at=NOW() WHERE id='019bf31d-2ef0-7747-8ae5-8e74f5f622ef';
UPDATE score_levels SET level=0, operator='>', lower_limit='9', upper_limit=NULL, name='>9', updated_at=NOW() WHERE id='019bf31d-2ef0-76ca-95b3-c3a05a738bd6';

-- Hidrogênio Basal (H₂ Jejum) (âncoras 16/20 consenso)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='8', name='≤8', updated_at=NOW() WHERE id='019bf7af-ba01-7b8d-a42c-94ac5f992e7b';
UPDATE score_levels SET level=4, operator='between', lower_limit='8', upper_limit='12', name='8-12', updated_at=NOW() WHERE id='019bf7af-ba01-7286-b58b-47e03aa6421c';
UPDATE score_levels SET level=3, operator='between', lower_limit='12', upper_limit='16', name='12-16', updated_at=NOW() WHERE id='019bf7af-ba01-737c-8bb7-9499e1f4603f';
UPDATE score_levels SET level=2, operator='between', lower_limit='16', upper_limit='20', name='16-20', updated_at=NOW() WHERE id='019bf7af-ba01-70bb-88d2-2bbb67e7866b';
UPDATE score_levels SET level=1, operator='between', lower_limit='20', upper_limit='35', name='20-35', updated_at=NOW() WHERE id='019bf7af-ba01-76a3-a7f7-5715f7b933e2';
UPDATE score_levels SET level=0, operator='>', lower_limit='35', upper_limit=NULL, name='>35', updated_at=NOW() WHERE id='019bf7af-ba01-7936-b711-347cf2c2b389';

-- Hora de dormir (mantido; piso <19h coberto via L2 <=21)
UPDATE score_levels SET level=5, operator='between', lower_limit='22', upper_limit='23', name='22 às 23h', updated_at=NOW() WHERE id='019c5395-0c3b-7d4c-86f4-c26395a3738c';
UPDATE score_levels SET level=4, operator='between', lower_limit='21', upper_limit='22', name='21 às 22h', updated_at=NOW() WHERE id='019c5395-493c-757d-8cac-dd86f0499b0e';
UPDATE score_levels SET level=3, operator='between', lower_limit='23', upper_limit='25', name='23h à 1h', updated_at=NOW() WHERE id='019c5395-88e1-7249-ba0e-7286a4b99584';
UPDATE score_levels SET level=2, operator='<=', lower_limit=NULL, upper_limit='21', name='19 às 21h', updated_at=NOW() WHERE id='019c5395-cb5a-7c55-9da6-97fbe98db43e';
UPDATE score_levels SET level=1, operator='between', lower_limit='25', upper_limit='27', name='1h às 3h', updated_at=NOW() WHERE id='019c53a2-5130-79af-9be9-0a074df4142d';
UPDATE score_levels SET level=0, operator='>', lower_limit='27', upper_limit=NULL, name='Após 3h', updated_at=NOW() WHERE id='019c5396-0eb0-72b7-88b9-5002918767f8';

-- Imunoglobulina E (IgE) (piso fechado <=31)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='31', name='0-31', updated_at=NOW() WHERE id='019bf7af-d85e-71b5-aec2-4cd266ae9b0e';
UPDATE score_levels SET level=4, operator='between', lower_limit='31', upper_limit='101', name='31-101', updated_at=NOW() WHERE id='019bf7af-d85e-7f9b-b8cc-69dac59f4584';
UPDATE score_levels SET level=2, operator='between', lower_limit='101', upper_limit='201', name='101-201', updated_at=NOW() WHERE id='019bf7af-d85e-7b4c-b487-690f76223ee4';
UPDATE score_levels SET level=1, operator='between', lower_limit='201', upper_limit='1000', name='201-1000', updated_at=NOW() WHERE id='019bf7af-d85e-7192-ac9c-90df6b3e038f';
UPDATE score_levels SET level=0, operator='>', lower_limit='1000', upper_limit=NULL, name='>1000', updated_at=NOW() WHERE id='019bf7af-d85e-7680-b27a-e233c815ca95';

-- Magnésio RBC (piso 4,2 / ótimo ≥6,0)
UPDATE score_levels SET level=5, operator='>=', lower_limit='6.0', upper_limit=NULL, name='≥6,0', updated_at=NOW() WHERE id='019bf31d-2ef0-78b6-a287-c6a1787c21cd';
UPDATE score_levels SET level=4, operator='between', lower_limit='5.6', upper_limit='6.0', name='5,6-6,0', updated_at=NOW() WHERE id='019bf31d-2ef0-76ce-acfa-dc58576b92e4';
UPDATE score_levels SET level=3, operator='between', lower_limit='5.2', upper_limit='5.6', name='5,2-5,6', updated_at=NOW() WHERE id='019bf31d-2ef0-77ca-809c-8c3193057086';
UPDATE score_levels SET level=2, operator='between', lower_limit='4.8', upper_limit='5.2', name='4,8-5,2', updated_at=NOW() WHERE id='019bf31d-2ef0-77c8-93b2-48aa9e364c48';
UPDATE score_levels SET level=1, operator='between', lower_limit='4.2', upper_limit='4.8', name='4,2-4,8', updated_at=NOW() WHERE id='019bf31d-2ef0-76fe-bbd3-437b986a51cb';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='4.2', name='≤4,2', updated_at=NOW() WHERE id='019bf31d-2ef0-7dcb-a579-e58dac93db48';

-- Magnésio Sérico (U invertido, ideal 2,36-2,5)
UPDATE score_levels SET level=5, operator='between', lower_limit='2.36', upper_limit='2.5', name='2,36-2,5', updated_at=NOW() WHERE id='019bf31d-2ef0-76ec-b227-3ee4785a9233';
UPDATE score_levels SET level=4, operator='between', lower_limit='2.1', upper_limit='2.36', name='2,1-2,36', updated_at=NOW() WHERE id='019bf31d-2ef0-7d90-9ca3-e3f9138951ef';
UPDATE score_levels SET level=3, operator='between', lower_limit='2.5', upper_limit='3.5', name='2,5-3,5', updated_at=NOW() WHERE id='019bf31d-2ef0-7105-a54a-c0142128a51a';
UPDATE score_levels SET level=2, operator='between', lower_limit='1.5', upper_limit='2.1', name='1,5-2,1', updated_at=NOW() WHERE id='019bf31d-2ef0-72d8-8676-ffd504574ce5';
UPDATE score_levels SET level=1, operator='>', lower_limit='3.5', upper_limit=NULL, name='>3,5', updated_at=NOW() WHERE id='019bf31d-2ef0-7d9a-b915-7ead6dc30240';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='1.5', name='≤1,5', updated_at=NOW() WHERE id='019bf31d-2ef0-792d-b972-8d57d6abefe2';

-- Metano Basal (CH₄ Jejum) (âncora IMO 10)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='3', name='≤3', updated_at=NOW() WHERE id='019bf31d-2ef0-7c79-b7e5-c01c544268e1';
UPDATE score_levels SET level=4, operator='between', lower_limit='3', upper_limit='7', name='3-7', updated_at=NOW() WHERE id='019bf31d-2ef0-7331-87d5-0c92e8352131';
UPDATE score_levels SET level=3, operator='between', lower_limit='7', upper_limit='10', name='7-10', updated_at=NOW() WHERE id='019bf31d-2ef0-7bad-b022-860ec2ab08b3';
UPDATE score_levels SET level=2, operator='between', lower_limit='10', upper_limit='16', name='10-16', updated_at=NOW() WHERE id='019bf31d-2ef0-71cf-8666-ee6c1b626f53';
UPDATE score_levels SET level=1, operator='between', lower_limit='16', upper_limit='25', name='16-25', updated_at=NOW() WHERE id='019bf31d-2ef0-74d4-87f9-7a6b04a8b817';
UPDATE score_levels SET level=0, operator='>', lower_limit='25', upper_limit=NULL, name='>25', updated_at=NOW() WHERE id='019bf31d-2ef0-79b9-a009-3d4fb6e73bb1';

-- pO2 Venoso (U: >50 = L2 extração diminuída; L2 é nível NOVO -> INSERT mais abaixo)
UPDATE score_levels SET level=5, operator='between', lower_limit='40', upper_limit='50', name='40-50', updated_at=NOW() WHERE id='019bf31d-2ef0-7540-a28f-10b271979473';
UPDATE score_levels SET level=4, operator='between', lower_limit='35', upper_limit='40', name='35-40', updated_at=NOW() WHERE id='019bf31d-2ef0-7596-a3f7-59641999dfbf';
UPDATE score_levels SET level=3, operator='between', lower_limit='30', upper_limit='35', name='30-35', updated_at=NOW() WHERE id='019bf31d-2ef0-7786-b047-e85da5131f52';
UPDATE score_levels SET level=1, operator='between', lower_limit='25', upper_limit='30', name='25-30', updated_at=NOW() WHERE id='019bf31d-2ef0-7647-99a5-0e049927c020';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='25', name='≤25', updated_at=NOW() WHERE id='019bf31d-2ef0-7d66-898b-bef6ddb899f1';

-- Progesterona - Mulheres Gestação 1º Trimestre (melhor ≥25; antigo L1 5-10 -> L2 10-15; novo L1 -> INSERT)
UPDATE score_levels SET level=5, operator='>', lower_limit='25', upper_limit=NULL, name='>25', updated_at=NOW() WHERE id='019bf31d-2ef0-767c-b1f9-e8ad7e210895';
UPDATE score_levels SET level=4, operator='between', lower_limit='15', upper_limit='25', name='15-25', updated_at=NOW() WHERE id='019bf31d-2ef0-7431-aa0d-a32aaa88137b';
UPDATE score_levels SET level=2, operator='between', lower_limit='10', upper_limit='15', name='10-15', updated_at=NOW() WHERE id='019bf31d-2ef0-7925-b8d7-decf90b02ead';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='5.0', name='≤5', updated_at=NOW() WHERE id='019bf31d-2ef0-7df5-a1c2-341a29872cd2';

-- Progesterona - Mulheres Gestação 2º Trimestre (só piso; rótulo limpo)
UPDATE score_levels SET level=5, operator='>', lower_limit='25', upper_limit=NULL, name='>25', updated_at=NOW() WHERE id='019bf31d-2ef0-7750-893c-0d102a17763b';
UPDATE score_levels SET level=2, operator='between', lower_limit='17.0', upper_limit='25.0', name='17-25', updated_at=NOW() WHERE id='019bf31d-2ef0-79ea-82fc-ac89d8d9eb43';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='17.0', name='≤17', updated_at=NOW() WHERE id='019bf31d-2ef0-72a8-9fab-ffb5110f897d';

-- Progesterona - Mulheres Gestação 3º Trimestre (só piso; rótulo limpo)
UPDATE score_levels SET level=5, operator='>', lower_limit='80', upper_limit=NULL, name='>80', updated_at=NOW() WHERE id='019bf7af-ba01-7fd4-9e00-303cd84c7fc3';
UPDATE score_levels SET level=2, operator='between', lower_limit='55.0', upper_limit='80.0', name='55-80', updated_at=NOW() WHERE id='019bf7af-ba01-7e20-b847-90ee5fd88ee2';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='55.0', name='≤55', updated_at=NOW() WHERE id='019bf7af-ba01-71c0-9441-fbd2936212c3';

-- Testosterona Total - Homens (651->650; topo >900; U funcional intacto)
UPDATE score_levels SET level=5, operator='between', lower_limit='650', upper_limit='900', name='650-900', updated_at=NOW() WHERE id='019bf7af-d85e-710e-aacd-fac7d0b36c6e';
UPDATE score_levels SET level=4, operator='between', lower_limit='550', upper_limit='650', name='550-650', updated_at=NOW() WHERE id='019bf7af-d85e-7d69-82d5-2eba32adece0';
UPDATE score_levels SET level=3, operator='>', lower_limit='900', upper_limit=NULL, name='>900', updated_at=NOW() WHERE id='019bf7af-d85e-701e-8434-8dbabae8e0ee';
UPDATE score_levels SET level=2, operator='between', lower_limit='300', upper_limit='550', name='300-550', updated_at=NOW() WHERE id='019bf7af-d85e-7293-958a-b60b82ad1454';
UPDATE score_levels SET level=1, operator='between', lower_limit='200', upper_limit='300', name='200-300', updated_at=NOW() WHERE id='019bf7af-d85e-739d-b3b3-9ba1109a84ba';
UPDATE score_levels SET level=0, operator='<=', lower_limit=NULL, upper_limit='200', name='≤200', updated_at=NOW() WHERE id='019bf7af-d85e-77bd-8f6e-9852c3a5c7eb';

-- USG Transvaginal - Volume Ovariano (piso benigno <=8; sem split menopausa)
UPDATE score_levels SET level=5, operator='<=', lower_limit=NULL, upper_limit='8', name='≤8', updated_at=NOW() WHERE id='019bf31d-2ef0-7c5c-982f-a46490610673';
UPDATE score_levels SET level=3, operator='between', lower_limit='8', upper_limit='10', name='8-10', updated_at=NOW() WHERE id='019bf31d-2ef0-7c93-b6bc-c2cedbf9daa4';
UPDATE score_levels SET level=1, operator='between', lower_limit='10', upper_limit='15', name='10-15', updated_at=NOW() WHERE id='019bf31d-2ef0-7494-92a4-351c7764ab4b';
UPDATE score_levels SET level=0, operator='>', lower_limit='15', upper_limit=NULL, name='>15', updated_at=NOW() WHERE id='019bf31d-2ef0-742e-986d-62dff8abbe38';

-- Níveis NOVOS (não existem em prod): inseridos com o mesmo id gerado no dev (dev≡prod).
INSERT INTO score_levels (id, level, name, operator, lower_limit, upper_limit, item_id, created_at, updated_at)
SELECT '019f0c3e-6861-7976-b9c9-fe37252da8e8', 1, '5-10', 'between', '5', '10', si.id, NOW(), NOW()
FROM score_items si WHERE si.name='Progesterona - Mulheres Gestação 1º Trimestre' AND si.deleted_at IS NULL
ON CONFLICT (id) DO NOTHING;

INSERT INTO score_levels (id, level, name, operator, lower_limit, upper_limit, item_id, created_at, updated_at)
SELECT '019f0c5d-86f5-79e5-9e40-6bf1f0cba14c', 2, '>50', '>', '50', NULL, si.id, NOW(), NOW()
FROM score_items si WHERE si.name='pO2 Venoso' AND si.deleted_at IS NULL
ON CONFLICT (id) DO NOTHING;

-- +goose Down
-- Reverte os 2 níveis inseridos (novo L1 da Progesterona 1º Trim e novo L2 do pO2 Venoso).
DELETE FROM score_levels WHERE id='019f0c3e-6861-7976-b9c9-fe37252da8e8';
DELETE FROM score_levels WHERE id='019f0c5d-86f5-79e5-9e40-6bf1f0cba14c';
-- Os 78 UPDATEs de faixa NÃO são restaurados automaticamente (os valores anteriores de prod não
-- foram snapshotados nesta migration). Para rollback completo das faixas, restaurar do backup
-- pré-deploy do banco. As mudanças são apenas de configuração de score (sem perda de dados de
-- paciente: os ids dos níveis são preservados, então os links históricos permanecem).
