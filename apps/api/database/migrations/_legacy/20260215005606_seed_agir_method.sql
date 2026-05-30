-- Seed AGIR Method with Letters and Pillars
-- Migration created: 2026-02-15

-- Insert AGIR Method
DO $$
DECLARE
    agir_method_id UUID;
    letter_a_id UUID;
    letter_g_id UUID;
    letter_i_id UUID;
    letter_r_id UUID;
BEGIN
    -- Create AGIR method if it doesn't exist
    IF NOT EXISTS (SELECT 1 FROM methods WHERE short_name = 'AGIR') THEN
        INSERT INTO methods (id, name, short_name, description, color, "order", created_at, updated_at)
        VALUES (
            gen_random_uuid(),
            'AGIR - Protocolo de Saúde Integrativa',
            'AGIR',
            'Alimentação e Atividade Física, Gestão metabólica, Integração mente-corpo, Ritmo Circadiano',
            '#6366F1',
            1,
            NOW(),
            NOW()
        )
        RETURNING id INTO agir_method_id;

        -- Insert Letters
        INSERT INTO method_letters (id, code, name, description, color, icon, "order", method_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'A', 'Alimentação e Atividade Física', 'Pilar de nutrição e exercício físico', '#10B981', '🥗', 1, agir_method_id, NOW(), NOW())
        RETURNING id INTO letter_a_id;

        INSERT INTO method_letters (id, code, name, description, color, icon, "order", method_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'G', 'Gestão Metabólica', 'Pilar de controle metabólico e hormonal', '#3B82F6', '⚡', 2, agir_method_id, NOW(), NOW())
        RETURNING id INTO letter_g_id;

        INSERT INTO method_letters (id, code, name, description, color, icon, "order", method_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'I', 'Integração Mente-Corpo', 'Pilar de saúde mental e integração psicossomática', '#8B5CF6', '🧠', 3, agir_method_id, NOW(), NOW())
        RETURNING id INTO letter_i_id;

        INSERT INTO method_letters (id, code, name, description, color, icon, "order", method_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'R', 'Ritmo Circadiano', 'Pilar de cronobiologia e ciclos circadianos', '#F59E0B', '🌙', 4, agir_method_id, NOW(), NOW())
        RETURNING id INTO letter_r_id;

        -- Insert example pillars for each letter

        -- A - Alimentação e Atividade Física
        INSERT INTO method_pillars (id, name, description, "order", letter_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'Avaliação Nutricional', 'Avaliação do estado nutricional e padrão alimentar', 1, letter_a_id, NOW(), NOW()),
            (gen_random_uuid(), 'Prescrição de Exercícios', 'Planejamento de atividade física personalizada', 2, letter_a_id, NOW(), NOW()),
            (gen_random_uuid(), 'Composição Corporal', 'Análise de massa muscular, gordura e hidratação', 3, letter_a_id, NOW(), NOW());

        -- G - Gestão Metabólica
        INSERT INTO method_pillars (id, name, description, "order", letter_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'Controle Glicêmico', 'Monitoramento e controle da glicemia', 1, letter_g_id, NOW(), NOW()),
            (gen_random_uuid(), 'Perfil Lipídico', 'Avaliação e manejo dos lipídeos', 2, letter_g_id, NOW(), NOW()),
            (gen_random_uuid(), 'Função Hepática', 'Avaliação da saúde hepática e metabólica', 3, letter_g_id, NOW(), NOW()),
            (gen_random_uuid(), 'Função Renal', 'Monitoramento da função renal', 4, letter_g_id, NOW(), NOW());

        -- I - Integração Mente-Corpo
        INSERT INTO method_pillars (id, name, description, "order", letter_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'Avaliação Psicológica', 'Avaliação de saúde mental e estresse', 1, letter_i_id, NOW(), NOW()),
            (gen_random_uuid(), 'Técnicas de Relaxamento', 'Práticas de mindfulness e relaxamento', 2, letter_i_id, NOW(), NOW()),
            (gen_random_uuid(), 'Função Cognitiva', 'Avaliação de memória e função executiva', 3, letter_i_id, NOW(), NOW());

        -- R - Ritmo Circadiano
        INSERT INTO method_pillars (id, name, description, "order", letter_id, created_at, updated_at)
        VALUES
            (gen_random_uuid(), 'Qualidade do Sono', 'Avaliação e otimização do sono', 1, letter_r_id, NOW(), NOW()),
            (gen_random_uuid(), 'Cronobiologia', 'Sincronização de ritmos circadianos', 2, letter_r_id, NOW(), NOW()),
            (gen_random_uuid(), 'Exposição à Luz', 'Otimização de exposição solar e luz artificial', 3, letter_r_id, NOW(), NOW());

        RAISE NOTICE 'AGIR method created successfully with ID: %', agir_method_id;
    ELSE
        RAISE NOTICE 'AGIR method already exists, skipping seed';
    END IF;
END $$;
