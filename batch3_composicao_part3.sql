-- ============================================================
-- BATCH 3 PARTE 3: Items finais (ângulo de fase e TMB)
-- ============================================================

-- Items de ângulo de fase (homem e mulher)
UPDATE score_items
SET
  clinical_relevance = 'O ângulo de fase (PhA, phase angle) medido por bioimpedância é marcador de integridade celular, funcionalidade de membrana e prognóstico. Reflete relação entre resistência e reactância: células saudáveis com membranas íntegras têm maior capacitância (maior reactância, maior ângulo de fase). Valores de referência: homens saudáveis ~6-8°, mulheres ~5-6.5° (varia com idade, método). Ângulo de fase reduzido (<5° homens, <4.5° mulheres) indica comprometimento celular e associa-se a: desnutrição, sarcopenia, senescência celular, doença crítica, câncer, HIV/AIDS, insuficiência cardíaca, renal ou hepática, inflamação crônica, estresse oxidativo, maior risco de complicações e mortalidade. Ângulo de fase elevado indica células íntegras, boa massa celular, vitalidade. É marcador prognóstico independente em várias condições clínicas. Na medicina funcional integrativa, PhA baixo alerta para investigar: desnutrição proteica, deficiências nutricionais, inflamação crônica, estresse oxidativo, disfunção mitocondrial, sobrecarga tóxica. Intervenção: nutrição adequada (proteína, antioxidantes, ômega-3), exercício (especialmente resistido - aumenta PhA), controle de inflamação, suplementação direcionada (glutamina, arginina, HMB em desnutrição grave), tratamento de doenças de base. Ângulo de fase melhora com ganho de massa celular e redução de inflamação.',

  patient_explanation = 'O ângulo de fase é um marcador avançado medido pela bioimpedância que avalia a QUALIDADE e VITALIDADE das suas células. Células saudáveis têm membranas íntegras que seguram água e nutrientes dentro delas, gerando um ângulo de fase maior. Ângulo de fase baixo (<5° homens, <4.5° mulheres) indica que suas células estão "fracas", desnutridas, inflamadas ou envelhecidas - isso pode acontecer em desnutrição, doenças crônicas, inflamação persistente, ou envelhecimento acelerado. Ângulo de fase alto indica células vigorosas, boa massa muscular, boa nutrição. É um dos melhores marcadores de prognóstico de saúde: pessoas com ângulo de fase maior têm melhor resposta a tratamentos, recuperação mais rápida de doenças, e maior longevidade. Podemos MELHORAR seu ângulo de fase através de nutrição adequada (especialmente proteínas e antioxidantes), exercícios de musculação (que constroem células musculares saudáveis), controle de inflamação, e suplementação direcionada quando necessário. É um marcador que literalmente reflete a vitalidade celular do seu corpo.',

  conduct = 'Medir ângulo de fase através de bioimpedância multifrequencial (equipamentos que fornecem PhA: InBody, Tanita profissional, analisadores clínicos). Protocolo padronizado: jejum mínimo 4h, hidratação normal, evitar exercício 12h antes, bexiga vazia. Registrar PhA em graus. Interpretar segundo sexo e idade: valores de referência aproximados: homens 6-8°, mulheres 5-6.5°. PhA <5° (H) ou <4.5° (M) é preocupante e requer investigação. Solicitar exames complementares se PhA baixo: albumina, pré-albumina, proteínas totais, hemograma completo, marcadores inflamatórios (PCR-us, VHS, ferritina), perfil nutricional (vitaminas, minerais, ômega-3), função hepática e renal, avaliação de massa muscular. Investigar causas: desnutrição (ingestão inadequada, má absorção), inflamação crônica (autoimune, infecciosa, neoplásica), doenças consumptivas, estresse oxidativo. Intervenção para aumentar PhA: nutrição hiperproteica (1.5-2.0g/kg/dia), suplementação de aminoácidos essenciais (leucina, HMB), ômega-3 (anti-inflamatório, melhora membranas celulares), antioxidantes (vitamina C, E, selênio, glutationa), exercício resistido progressivo 3-4x/semana (estimula síntese proteica e melhora integridade celular), controle de inflamação, tratamento de doenças de base. Reavaliação mensal durante fase intensiva, trimestral em manutenção. Objetivo: PhA >5.5° (H), >5° (M); idealmente >6° (H), >5.5° (M).',

  updated_at = NOW()
WHERE id IN ('f694e6f7-8d03-4cbd-9813-3635f0bee1d5', 'ad60ffd2-8d47-4d9f-9924-47e7a3b5830b');

-- ============================================================
-- Item: Taxa metabólica basal (kcal)
-- ============================================================

UPDATE score_items
SET
  clinical_relevance = 'A taxa metabólica basal (TMB ou REE - Resting Energy Expenditure) é o gasto energético mínimo necessário para manter funções vitais em repouso (respiração, circulação, síntese proteica, função neuronal, temperatura corporal). Representa ~60-70% do gasto energético total diário (GET) em indivíduos sedentários. Determinantes principais da TMB: massa muscular (músculo é metabolicamente ativo, ~13kcal/kg/dia), massa gorda (menos ativa, ~4.5kcal/kg/dia), idade (declínio ~2-3%/década após 30 anos devido a sarcopenia), sexo (homens têm TMB ~5-10% maior devido a maior massa muscular), genética, hormônios (tireoide, testosterona, GH aumentam TMB; hipotireoidismo, hipogonadismo reduzem), temperatura ambiente, estado nutricional. Métodos de avaliação: calorimetria indireta (padrão-ouro, mede VO2 e VCO2), equações preditivas (Harris-Benedict, Mifflin-St Jeor, Cunningham), bioimpedância (estima com base em composição corporal). Limitações de equações: não consideram variações individuais, metabolismo adaptativo, disfunções hormonais. Metabolismo adaptativo: após perda ponderal ou restrição calórica crônica, TMB pode reduzir além do esperado pela mudança de composição (down-regulation metabólico compensatório, mediado por leptina, tireoide, sistema nervoso simpático) - fenômeno que dificulta manutenção de perda ponderal. Na medicina funcional integrativa, investigamos causas de TMB reduzida: hipotireoidismo (mesmo subclínico), déficit de testosterona/GH, sarcopenia, metabolismo adaptativo pós-dietas restritivas, deficiências nutricionais (ferro, iodo, selênio, zinco), privação de sono, estresse crônico (paradoxalmente pode reduzir TMB via alterações hormonais). Estratégias para aumentar TMB: ganho de massa muscular (exercício resistido), termogênese alimentar (proteína tem maior TEF - Thermic Effect of Food ~25-30%), exercício (EPOC - Excess Post-exercise Oxygen Consumption após treino intenso), otimização hormonal, correção de deficiências, sono adequado, exposição ao frio (termogênese adaptativa via gordura marrom).',

  patient_explanation = 'Sua taxa metabólica basal (TMB) é a quantidade de calorias que seu corpo gasta por dia apenas para manter funções vitais em repouso - respirar, circular sangue, fazer o coração bater, manter temperatura corporal, regenerar células. É como o consumo de gasolina de um carro ligado mas parado. A TMB representa cerca de 60-70% de todo o gasto calórico diário em pessoas sedentárias. O que determina sua TMB? Principalmente a quantidade de MÚSCULO que você tem - músculo queima calorias mesmo em repouso, cerca de 13 calorias por quilo por dia. Gordura queima muito menos, apenas 4-5 calorias por quilo. Por isso, duas pessoas de mesmo peso podem ter TMBs muito diferentes: quem tem mais músculo gasta mais calorias parado. Outros fatores que influenciam TMB: idade (metabolismo diminui ~2-3% a cada década após os 30 anos, principalmente por perda muscular), sexo (homens têm TMB maior porque têm mais músculo), hormônios (tireoide e testosterona aumentam metabolismo), genética. Um problema comum: após várias dietas restritivas, a TMB pode CAIR além do esperado - seu corpo "aprende" a economizar energia, fenômeno chamado metabolismo adaptativo, que explica por que fica cada vez mais difícil emagrecer e fácil reganhar peso. Como AUMENTAR sua TMB? Ganhar massa muscular é a melhor estratégia (musculação), comer proteína suficiente (proteína aumenta metabolismo temporariamente), fazer exercícios intensos (metabolismo fica elevado por horas após treino), otimizar hormônios (corrigir tireoide e testosterona se baixos), dormir bem, e corrigir deficiências de nutrientes que afetam metabolismo (ferro, iodo, selênio, zinco).',

  conduct = 'Avaliar TMB preferencialmente por calorimetria indireta (gold standard, se disponível) ou estimar por equações validadas ajustadas para composição corporal. Equação de Mifflin-St Jeor (mais precisa em populações gerais): Homens: TMB = (10 × peso_kg) + (6.25 × altura_cm) - (5 × idade_anos) + 5; Mulheres: TMB = (10 × peso_kg) + (6.25 × altura_cm) - (5 × idade_anos) - 161. Equação de Cunningham (mais precisa se massa magra conhecida): TMB = 500 + (22 × massa_magra_kg). Bioimpedâncias avançadas fornecem estimativa de TMB baseada em composição corporal. Comparar TMB medida/estimada com valores preditos: TMB significativamente menor que o esperado (>10-15% abaixo) sugere metabolismo adaptativo (histórico de dietas restritivas), hipotireoidismo, hipogonadismo, sarcopenia, ou outras disfunções metabólicas. Investigar: perfil tiroideano completo (TSH, T3 livre, T4 livre, T3 reverso - T3r elevado indica down-regulation metabólico), testosterona (homens), cortisol (estresse crônico pode paradoxalmente reduzir TMB via alterações tireoideanas), leptina (se disponível - níveis baixos indicam privação energética), avaliação de composição corporal (sarcopenia reduz TMB), histórico de dietas (restrição calórica crônica causa adaptação metabólica), qualidade de sono (privação reduz TMB), deficiências nutricionais (ferro, iodo, selênio, zinco, vitamina D). Estratégias para aumentar TMB: treinamento resistido progressivo para ganho de massa muscular (efeito mais sustentado), HIIT (aumenta EPOC - metabolismo elevado pós-exercício), aporte proteico adequado (1.6-2.2g/kg/dia, distribuído em 4-5 refeições - TEF de proteína é 25-30% vs 5-10% carboidratos e 0-3% gordura), reverse dieting se metabolismo adaptativo (aumento gradual e progressivo de calorias para restaurar metabolismo, 50-100kcal/semana até TMB normalizar), otimização hormonal (reposição tireoidiana se hipotireoidismo, testosterona se hipogonadismo sintomático), correção de deficiências nutricionais, sono reparador (7-9h/noite, boa qualidade), redução de estresse crônico, considerar exposição ao frio controlada (ativa gordura marrom - brown adipose tissue - termogênese adaptativa). Calcular GET (Gasto Energético Total) = TMB × fator de atividade (sedentário 1.2, levemente ativo 1.375, moderadamente ativo 1.55, muito ativo 1.725, extremamente ativo 1.9). Prescrição calórica individualizada: para perda de gordura, déficit moderado de 300-500kcal/dia (não mais para evitar metabolismo adaptativo), para ganho muscular, superávit de 200-400kcal/dia, para manutenção, ingerir próximo ao GET. Reavaliar TMB e composição corporal trimestralmente, ajustar prescrição conforme resposta.',

  updated_at = NOW()
WHERE id = 'b6f334f1-e4dc-4c30-b627-a7ab9c14f5fd';

-- ============================================================
-- Linkar artigos científicos a TODOS os items do batch
-- ============================================================

-- Artigos de emagrecimento e metabolismo são relevantes para todo o batch
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, si.id
FROM articles a
CROSS JOIN score_items si
WHERE si.id IN (
  -- IDs de todos os 46 items do batch
  'ac04ab06-bdc7-4959-a33d-32f3646c1a87', '3b9f7a1e-527e-4a4c-bc23-87c68c41ee89',
  '7b4975a7-a24d-4bbc-8d27-2dd89f9a99ad', 'c1c9eb73-b2ac-456f-b854-ed707c817a3d',
  '03438548-88b3-44a7-8449-5441c26ea829', '088b4d4d-1873-45bb-8a3a-19e4463de7a5',
  '2579198d-65b8-4cff-9970-7c437076adc7', 'f836753e-9e45-48f5-a54c-298195ff0588',
  '662902a2-d3f7-4d08-b3ee-4fde9e633cf8', '400295c1-25e5-44c5-8d99-0f355f1aa7cc',
  'bf875f17-1a90-4bb8-8305-e0e0285aeb80', '1a9be52d-aa0b-4a53-a71a-4eff3e0cc6ba',
  'c9348fbd-df44-4903-ac22-1d8b5b47179a', 'b2414f5e-8ad7-4b9c-846e-edd6a3c8277f',
  '936855d1-e715-4171-a061-3bcc500957f7', 'c3ee7527-104b-467a-a87e-b4587192005b',
  '421db9cd-d75d-4f0a-8188-216affee192f', 'ff5472cf-e148-48c1-be7a-245df89c200c',
  '472969f8-7ea9-4532-8c73-c4a96f31e2fc', '094b4b55-e7a4-4587-be08-9963c0520bf0',
  '660ec23e-8957-4936-b83b-284a0f8c84eb', '30fa255c-83d9-4cd4-8b06-75f70e2fb3eb',
  '371f12db-5a68-4092-a4c3-1430ec21f18a', '8062469b-618a-4cfc-9239-046fd1f882e2',
  '25822830-7d95-4e20-9b32-bb5cd7d4a3d1', '2ff747e7-e48d-4207-85ee-168b1c14e35e',
  '725d0bc7-8f25-4c44-b268-53b72f75adff', 'cd592cb8-4c34-4964-ac29-63bf9c224bef',
  'bdf03bb7-ef91-4e12-9247-4e29ce1e7185', 'f0e7e520-6e01-437e-9274-5a5c90255d77',
  '1442a990-a9f0-4cd3-85b1-f3d22705e469', 'e0dbadd4-b307-43ff-8438-c24fa0876d10',
  'd3f0556e-53cb-4a29-b825-a84bc8f38f3e', 'a4e0dc5b-e931-4126-a8b7-ac530f86aa68',
  '2265fbe7-fc73-48e1-ac8d-c8e239d29e3c', 'bc698d3d-c9da-4276-b0d1-677ebd1fdf95',
  'e5b6d407-4744-48eb-8efb-bdae72dc34b9', '8bf6396f-b07c-4c93-82c4-b69ec8e87aa2',
  '2eafc97e-7134-49d7-b34c-6903192383d5', 'f28c8a35-fa77-49a0-a69f-ba649fcffef2',
  'e79394a0-3d1f-4d18-8c55-d6a20179017a', '7857982a-8b2b-401f-bf42-8db8a3af264b',
  '806cb0c4-1954-4567-b727-55bf4560553e', 'f694e6f7-8d03-4cbd-9813-3635f0bee1d5',
  'ad60ffd2-8d47-4d9f-9924-47e7a3b5830b', 'b6f334f1-e4dc-4c30-b627-a7ab9c14f5fd'
)
AND (
  a.title ILIKE '%emagrecimento%'
  OR a.title ILIKE '%dieta%'
  OR a.title ILIKE '%metabol%'
  OR a.title ILIKE '%exercício%'
  OR a.title ILIKE '%bioquímica%'
  OR a.title ILIKE '%hormônio%'
  OR a.title ILIKE '%nutrição%'
  OR a.title ILIKE '%bases metabólicas%'
  OR a.title ILIKE '%inflamação%'
  OR a.title ILIKE '%oxidação%'
)
LIMIT 300
ON CONFLICT DO NOTHING;

-- ============================================================
-- FIM DO BATCH 3 - COMPOSIÇÃO CORPORAL
-- Total: 46 items enriquecidos
-- ============================================================
