-- Verificar os 30 itens de Composição Corporal

SELECT
    id,
    question,
    description,
    CASE
        WHEN LENGTH(clinical_context) > 50 THEN LEFT(clinical_context, 50) || '...'
        ELSE COALESCE(clinical_context, '(vazio)')
    END as clinical_context_preview,
    LENGTH(clinical_context) as context_length
FROM score_items
WHERE id IN (
    '299e22bd-184e-4513-8f21-26f26d91d737',
    '9ff7a160-8ad1-4c2a-857a-53677355a631',
    '35f8405e-5bd5-4803-bade-c50e6d615582',
    '29ec8df2-5b0f-442d-aa13-1647a9759a0d',
    '119de191-7399-47e4-8c4a-e3e5f21623ff',
    'a2688922-5b23-4f7c-a842-ffff6acf081e',
    'b73bd6c3-f529-41c2-b3de-7b322663f22e',
    '924e147d-eb99-4539-a1c3-424d79f577b7',
    '9a97c090-ce0b-4dbd-a030-e652542afc6c',
    '08c992b6-3bc9-425a-b848-a9fa0bc0c0f2',
    '8fef997c-73d3-44fa-9aec-bc15f625068c',
    'f6a6515f-5488-455d-9459-8c606a13029e',
    '3be0bbdb-15bb-4899-ae09-35a811ba6c62',
    '48b082bf-3697-4c8a-a183-b2fc4396d270',
    '40a6fbbc-c6c1-4086-9622-b66d4cb67d17',
    '405b0018-088c-4a35-8688-11de766fc557',
    '01e84baf-d377-4cc4-ab36-3661b3868f1a',
    '9cbf71b3-83de-4b71-b987-120e525c790e',
    '779d1bf5-d183-4607-9c86-672badfb78e0',
    '7fe0b34c-151d-407f-9de8-6ff2492dde4b',
    '30fa255c-83d9-4cd4-8b06-75f70e2fb3eb',
    '371f12db-5a68-4092-a4c3-1430ec21f18a',
    '9e815787-7948-4030-a470-6f050a91b2f8',
    '8e0efcc9-8281-48dc-856e-fa86a845e97d',
    'ac2da1dc-80aa-4035-8703-5e669caa37a6',
    '088b4d4d-1873-45bb-8a3a-19e4463de7a5',
    '0b58623e-63d3-4685-b864-c41a0058ed56',
    '731d8307-6295-469e-95b6-361373449dcf',
    '8062469b-618a-4cfc-9239-046fd1f882e2',
    'f72f4d3f-3af8-408a-9b45-4a2988bc778a'
)
ORDER BY question;
