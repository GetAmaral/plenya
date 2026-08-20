-- Semente inicial do catálogo magistral (dev e prod).
--
-- REGRA QUE VALE MAIS QUE O CONTEÚDO: aqui só entra o que tem fonte. Densidade aparente NÃO é
-- semeada — não existe tabela pública confiável, e ela varia por lote e por compactação da
-- farmácia (Anfarmag; Quallitá/Renylab). Enquanto a densidade não for cadastrada à mão, a
-- calculadora de cápsula se cala em vez de chutar, que é o comportamento correto.
--
-- O que entra são SINALIZADORES com fonte:
--   · formadores de mistura eutética: mentol, cânfora, timol, fenol, resorcina (sólido + sólido
--     vira líquido por abaixamento do ponto de fusão; contorna-se com pó adsorvente).
--   · talco adsorve cianocobalamina.
--
-- Fontes: Anfarmag, "Anfarmag dá dicas de como solucionar as incompatibilidades na manipulação
-- magistral" (anfarmag.org.br); Memento Terapêutico da Farmácia Universitária da UFRJ, 4ª ed.
-- 2021; Acofarma, "Instabilidade na formulação".
--
-- Idempotente: pode rodar de novo sem duplicar.

INSERT INTO public.magistral_components
    (id, name, synonyms, default_unit, eutectic_former, source, notes, last_review)
VALUES
    (uuidv7(), 'Mentol',    'levomentol, mentol racêmico', 'mg', true,  'seed',
     'Forma mistura eutética com cânfora, timol, fenol e resorcina. Contorna-se interpondo pó adsorvente (óxido/carbonato de magnésio, amido, sílica).', now()),
    (uuidv7(), 'Cânfora',   'canfora',                     'mg', true,  'seed',
     'Forma mistura eutética com mentol e timol.', now()),
    (uuidv7(), 'Timol',     '',                            'mg', true,  'seed',
     'Forma mistura eutética com mentol e cânfora.', now()),
    (uuidv7(), 'Fenol',     '',                            'mg', true,  'seed',
     'Forma mistura eutética com mentol.', now()),
    (uuidv7(), 'Resorcina', 'resorcinol',                  'mg', true,  'seed',
     'Forma mistura eutética com mentol.', now()),
    (uuidv7(), 'Talco',     '',                            'mg', false, 'seed',
     'Adsorve cianocobalamina; evitar como deslizante em fórmula que a contenha.', now()),
    (uuidv7(), 'Cianocobalamina', 'vitamina B12',           'mcg', false, 'seed',
     'Adsorvida pelo talco.', now())
ON CONFLICT DO NOTHING;

-- Par curado: talco × cianocobalamina.
INSERT INTO public.magistral_incompatibilities
    (id, component_a_id, component_b_id, severity, mechanism, note, source, last_review)
SELECT
    uuidv7(),
    LEAST(a.id::text, b.id::text)::uuid,
    GREATEST(a.id::text, b.id::text)::uuid,
    'warn',
    'o talco adsorve a cianocobalamina e reduz a dose disponível',
    'Trocar o deslizante (ex.: dióxido de silício) ou separar em fórmulas distintas.',
    'Anfarmag — incompatibilidades na manipulação magistral',
    now()
FROM public.magistral_components a, public.magistral_components b
WHERE lower(a.name) = 'talco' AND lower(b.name) = 'cianocobalamina'
ON CONFLICT DO NOTHING;
