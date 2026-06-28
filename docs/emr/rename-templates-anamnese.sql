-- ============================================================================
-- Renomeação dos 13 templates de anamnese — padrão "[Continuum |] Prof | Momento"
-- ============================================================================
-- Aplica só os novos nomes (keyed por UUID fixo) SEM re-rodar o seed inteiro
-- (re-rodar o seed reverteria a curadoria manual de itens). Idempotente.
-- Padrão: labels curtas (Médico/Nutri/Psico/Ed. Física), separador " | ".
-- ============================================================================
BEGIN;
UPDATE anamnesis_templates SET name='Médico | Inicial',                            updated_at=now() WHERE id='11111111-1111-7111-8111-111111111101';
UPDATE anamnesis_templates SET name='Médico | Acompanhamento',                     updated_at=now() WHERE id='11111111-1111-7111-8111-111111111102';
UPDATE anamnesis_templates SET name='Médico | Revisão de Exames',                  updated_at=now() WHERE id='11111111-1111-7111-8111-111111111103';
UPDATE anamnesis_templates SET name='Continuum | Médico | Inicial',               updated_at=now() WHERE id='11111111-1111-7111-8111-111111111104';
UPDATE anamnesis_templates SET name='Continuum | Médico | Complemento',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111105';
UPDATE anamnesis_templates SET name='Continuum | Médico | Acompanhamento',        updated_at=now() WHERE id='11111111-1111-7111-8111-111111111106';
UPDATE anamnesis_templates SET name='Continuum | Médico | Reavaliação Trimestral',updated_at=now() WHERE id='11111111-1111-7111-8111-111111111107';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111108';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111109';
UPDATE anamnesis_templates SET name='Continuum | Psico | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111110';
UPDATE anamnesis_templates SET name='Continuum | Psico | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111111';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Inicial',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111112';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Acompanhamento',    updated_at=now() WHERE id='11111111-1111-7111-8111-111111111113';
SELECT name, area FROM anamnesis_templates WHERE deleted_at IS NULL ORDER BY area, name;
COMMIT;
