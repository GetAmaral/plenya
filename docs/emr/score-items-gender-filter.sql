-- Backfill do filtro de sexo (gender) em ScoreItems anatomicamente exclusivos.
--
-- O "motor" de filtragem por sexo já existe e está ligado:
--   - Backend:  ScoreItem.AppliesToPatient(patient)  (internal/models/score_item.go:199)
--   - Frontend: itemAppliesToPatient()  (components/anamnesis/AnamnesisTemplateItemsRenderer.tsx:93)
--               aplicado em organizeTemplateItems() ao carregar um template de anamnese.
-- O motor só age quando score_items.gender está preenchido ('male'|'female'). Estes itens
-- estavam como 'not_applicable' e por isso apareciam para os dois sexos (ex.: paciente homem
-- via "Histerectomia"; paciente mulher via "Prostatectomia").
--
-- Critério: SÓ itens anatomicamente/fisiologicamente exclusivos de um sexo. Itens que se
-- aplicam aos dois (libido, TRH, DHT, "Mamas" no exame físico, história gestacional dos PAIS,
-- SNPs genéticos) ficam como 'not_applicable' de propósito.
-- Idempotente (filtra por gender = 'not_applicable'). Aplicado em DEV; rodar em PROD quando autorizado.

BEGIN;

-- ===== Femininos (estruturas/fisiologia exclusivas) =====
UPDATE score_items SET gender = 'female', updated_at = now()
WHERE gender = 'not_applicable' AND id IN (
  '019bf31d-2ef0-77e9-9045-5c994cfcbf94', -- Histerectomia
  'c77cedd3-2800-7bd8-adb0-8b19fd925bd0', -- USG Transvaginal - O-RADS
  'c77cedd3-2800-712f-be04-315636a9e3ec', -- USG transvaginal (útero, ovário, anexos e vagina)
  '019bf31d-2ef0-7e6e-a77d-0e75475ca72d', -- USG Transvaginal - Volume Ovariano
  '019bf31d-2ef0-7f0c-b76b-806c1d6ff1fd', -- Vagina e Colo Uterino
  '019bf31d-2ef0-7078-b942-860ee56136a4', -- Vulva e Estruturas Externas
  'c77cedd3-2800-7166-80cd-65e5a7e956fc', -- Ciclo menstrual (duração, volume, sintomas, anticoncepcional)
  '019c4fad-ba4a-7294-80a2-2bc7434ee2db', -- Melhora de sintomas de peri-menopausa
  'c77cedd3-2800-7f39-a8fd-1332fa906fa5', -- Anticoncepcionais
  'c77cedd3-2800-77eb-885c-476883a4208f', -- Com uso de anticoncepcional
  'c77cedd3-2800-784b-b3bc-516e0c8887b4', -- Sem uso de anticoncepcional
  'c77cedd3-2800-75f3-a287-89c9b36f13b5'  -- Histórico reprodutivo (gestações, abortos, partos)
);

-- Feminino + pós-menopausa (o exame é definido para mulher pós-menopausa).
-- post_menopause só filtra mulher quando o status de menopausa está preenchido; mulher sem
-- esse dato continua vendo o item.
UPDATE score_items SET gender = 'female', post_menopause = true, updated_at = now()
WHERE gender = 'not_applicable'
  AND id = '019bf31d-2ef0-77e6-b363-8279c7bd3a69'; -- USG Transvaginal - Espessura Endometrial Pós-Menopausa

-- ===== Masculinos (estruturas/fisiologia exclusivas) =====
UPDATE score_items SET gender = 'male', updated_at = now()
WHERE gender = 'not_applicable' AND id IN (
  '019bf31d-2ef0-73a8-b1ae-3e27305e25d8', -- Prostatectomia
  '019bf31d-2ef0-7f97-8e14-799354166f5e', -- USG Próstata - Densidade PSA (PSAD)
  '019bf31d-2ef0-7016-96ad-2fa95ae09f99', -- PSA Total
  '019bf31d-2ef0-76de-bba3-93a62b6950e5', -- PSA Livre/Total (%Free PSA)
  '019bf31d-2ef0-7b42-9d37-487e91411a18', -- Escroto / epidídimos
  '019bf31d-2ef0-7e35-94d7-0d232cc258ce', -- Testículos
  '019bf31d-2ef0-737c-a061-b81dddab34cd', -- Disfunção erétil
  'c77cedd3-2800-7e5e-b7c0-bd312f6a7ed4'  -- IIEF-5 (Índice Internacional de Função Erétil)
);

-- ===== Imagem de mama (rastreio predominantemente feminino) =====
-- Decisão clínica do Dr. Getúlio: exames de imagem da mama são femininos; a seção "Mamas" do
-- exame físico (id 019bf31d-2ef0-7ad3-9340-a962ffb32c58) fica para ambos (ginecomastia / ca de
-- mama masculino) e por isso NÃO entra aqui.
UPDATE score_items SET gender = 'female', updated_at = now()
WHERE gender = 'not_applicable' AND id IN (
  'c77cedd3-2800-78bc-b316-71e29954eedd', -- Mamografia - Densidade Mamária
  'c77cedd3-2800-71a2-b317-07fb2e53253e'  -- USG mamas
);

COMMIT;
