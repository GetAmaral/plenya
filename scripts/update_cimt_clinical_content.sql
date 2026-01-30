-- Atualização do conteúdo clínico do score_item CIMT Carótidas
-- ID: 9064aa05-9f21-402f-9c36-49887d2b8ec9
-- Data: 2026-01-28

UPDATE score_items
SET
    clinical_relevance = 'A espessura íntima-média carotídea (CIMT) é um marcador subclínico de aterosclerose bem estabelecido, medido por ultrassonografia de alta resolução modo-B. A medida representa a distância entre a interface lúmen-íntima e a interface média-adventícia da parede arterial carotídea. Valores elevados de CIMT (>800 µm) estão significativamente associados com aumento do risco de doença arterial coronariana (HR: 2,15; IC 95%: 1,07-4,31) e infarto do miocárdio (HR: 2,46; IC 95%: 0,93-6,53), conforme demonstrado no estudo UK Biobank com 29.292 participantes. Meta-análises recentes confirmam relação direta entre CIMT e severidade da doença coronariana (p < 0,001), com valores acima de 1,54 mm fortemente associados à doença arterial coronariana grave. A medida combinada de múltiplos segmentos carotídeos (artéria carótida comum, bulbo e carótida interna) demonstra performance preditiva superior comparada à medida isolada da carótida comum. A presença de placas ateroscleróticas associada ao aumento da CIMT confere poder preditivo ainda maior para eventos cardiovasculares. O exame é especialmente valioso na reclassificação de risco em pacientes de risco intermediário, permitindo decisões terapêuticas mais assertivas. Protocolos ultrassonográficos estritos são essenciais para garantir reprodutibilidade, recomendando-se dynamic range próximo a 60 decibéis conforme consenso de Mannheim atualizado. A CIMT também serve como marcador substituto em ensaios clínicos para avaliar eficácia de terapias farmacológicas na redução da aterosclerose.',

    patient_explanation = 'A CIMT (espessura íntima-média carotídea) é um exame de ultrassom que mede a espessura da parede das artérias carótidas, localizadas no pescoço. Essas artérias levam sangue para o cérebro. O exame identifica precocemente o endurecimento e estreitamento das artérias (aterosclerose), mesmo antes de você ter sintomas. Quando a parede da artéria está mais espessa que o normal, indica acúmulo de gordura e inflamação, o que aumenta o risco de infarto, derrame e outros problemas cardiovasculares. O exame é indolor, não invasivo e dura cerca de 15-20 minutos. Valores normais geralmente são abaixo de 0,8 mm (800 micrômetros). Valores acima de 1,5 mm indicam alto risco cardiovascular. Este exame ajuda seu médico a decidir se você precisa iniciar medicações preventivas, fazer mudanças mais intensivas no estilo de vida, ou realizar acompanhamento mais frequente. É especialmente útil para pessoas com risco intermediário de doenças do coração.',

    conduct = 'Para pacientes com CIMT aumentada, recomenda-se abordagem multifatorial: 1) Estratificação de risco cardiovascular completa incluindo cálculo de escores de risco (Framingham, ASCVD), perfil lipídico avançado, glicemia, hemoglobina glicada e proteína C reativa ultrassensível. 2) Otimização agressiva de fatores de risco modificáveis: controle pressórico rigoroso (meta <130/80 mmHg), controle glicêmico em diabéticos (HbA1c <7%), cessação tabágica obrigatória. 3) Terapia hipolipemiante intensiva com estatinas de alta potência (atorvastatina 40-80mg ou rosuvastatina 20-40mg), visando LDL <70 mg/dL ou redução >50% dos valores basais; considerar ezetimibe ou inibidores PCSK9 se meta não atingida. 4) Antiagregação plaquetária com AAS 100mg/dia deve ser individualizada conforme risco-benefício. 5) Modificações intensivas do estilo de vida: dieta mediterrânea, atividade física aeróbica regular (150 min/semana), controle de peso (IMC <25 kg/m²). 6) Reavaliação da CIMT a cada 2-3 anos para monitorar progressão; progressão >50 µm/ano indica necessidade de intensificação terapêutica. 7) Investigação complementar com escore de cálcio coronário pode ser considerada para melhor estratificação. 8) Encaminhamento ao cardiologista está indicado se CIMT >1,5 mm, presença de placas obstrutivas ou múltiplas, ou progressão rápida documentada.',

    last_review = NOW()
WHERE id = '9064aa05-9f21-402f-9c36-49887d2b8ec9';

-- Verificar atualização
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_length,
    LENGTH(patient_explanation) as patient_length,
    LENGTH(conduct) as conduct_length,
    last_review
FROM score_items
WHERE id = '9064aa05-9f21-402f-9c36-49887d2b8ec9';
