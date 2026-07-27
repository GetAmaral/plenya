-- Recalibra as faixas de "Ferritina - Homens" (14pt, lab_test_code PLNCEFB97FD).
--
-- PROBLEMA
-- As faixas atuais (N5=45-80, N4=80-150, N0=>150) tratam ferritina >150 ng/mL em homem como o
-- PIOR nível possível, zerando 14 pontos. Isso não se sustenta:
--
-- 1) Não há desfecho duro. Nenhum ensaio mostra dano de ferritina 150-300 em homem sem
--    inflamação. O FeAST (Zacharski, JAMA 2007; n=1.277), único RCT de flebotomia com desfecho
--    duro, foi NEGATIVO (mortalidade HR 0,85, IC 0,67-1,08, p=0,17) — e sua ferritina basal era
--    ~122 ng/mL, ou seja, o ensaio nunca testou "baixar de 240 para 100". A randomização
--    mendeliana aponta na direção OPOSTA: ferro mais alto foi protetor para doença coronariana
--    (Gill, ATVB 2017: OR 0,94 por DP; Liu, JAHA 2024, corrobora), com sinal adverso restrito a
--    DM2. Mesmo na hemocromatose confirmada, o EASL 2022 registra "absence of high-level
--    evidence" para depleção de ferro.
--
-- 2) O limiar de investigação é 300, não 150. EASL 2022, literal: genotipar HFE em HOMENS só
--    com "TSAT >50% and ferritin >300 µg/L" (o corte de 45% que circula é o FEMININO).
--    Ferritina 150-300 com PCR normal não abre investigação nenhuma.
--
-- 3) O corte não discrimina. Em prod nesta data há 6 resultados masculinos de ferritina:
--    98,4 / 172 / 172 / 189 / 207 / 240 ng/mL. Pelas faixas atuais, QUATRO caem em N0 (172, 172,
--    189, 207) e o quinto (240) cairia. Um corte que põe 2/3 dos pacientes no pior nível não está
--    medindo o paciente, está medindo o corte.
--
-- 4) INCOERÊNCIA INTERNA — o argumento decisivo. O próprio catálogo já tem a versão correta:
--       Ferritina - Mulheres Pós-Menopausa : N5=30-200 · N3=200-300 · N0=>300
--       Ferritina - Homens (este item)     : N5=45-80  · N4=80-150  · N0=>150
--    Está invertido em relação à fisiologia. Homem tem estoque de ferro MAIOR que mulher, e a
--    mulher pós-menopausa converge para o padrão masculino exatamente por cessar a perda
--    menstrual. Não há leitura em que o teto masculino deva ser metade do feminino pós-menopausa.
--    O item de homens é o outlier; o de pós-menopausa é a referência.
--
-- 5) Viés contra atleta. Corredor de fundo tende a ferritina MAIS BAIXA (hemólise de impacto,
--    hepcidina pós-exercício via IL-6, hematúria de esforço, perda por suor). Penalizar 150-300
--    empurra flebotomia/doação num grupo cujo risco real é ferropenia — que degrada VO2máx.
--
-- DECISÃO
-- Alinhar com a estrutura do item de pós-menopausa, com granularidade maior no lado BAIXO, que é
-- onde a evidência de dano é sólida (deficiência tem consequência funcional certa; excesso
-- moderado, não). Mantém penalidade real para os dois extremos verdadeiros.
--
--   N0  ≤15        deficiência absoluta de ferro (pior nível: patologia certa)
--   N2  15-30      depleção de estoques
--   N3  30-50      baixo-normal; em atleta pode limitar performance
--   N5  50-200     faixa funcional; sem evidência de dano em desfecho duro
--   N4  200-300    alto-normal; vigiar com IST + PCR-us (OMS 2020 sinaliza >200 em homens)
--   N1  >300       hiperferritinemia: investigar (limiar EASL) — inflamação, álcool, síndrome
--                  metabólica, hepatopatia, neoplasia; HFE só se IST >50%
--
-- Inversão deliberada N0/N1 em relação ao desenho anterior: a deficiência absoluta (≤15) passa a
-- ser o pior nível, e a hiperferritinemia (>300) passa a N1. Motivo: ≤15 é patologia estabelecida
-- com consequência funcional; >300 é um SINAL PARA INVESTIGAR, não um diagnóstico, e a MR não
-- sustenta dano cardiovascular do ferro alto.
--
-- Convenção de fronteira: 'between' é meio-aberto (> lower e <= upper); '<'/'<=' guardam o limite
-- em upper_limit — ver models.ScoreLevel.EvaluatesTrue. Empate resolve pelo primeiro match com
-- níveis ordenados ASC. A partição abaixo cobre o domínio sem gap nem sobreposição:
--   ≤15 | (15,30] | (30,50] | (50,200] | (200,300] | >300
--
-- RAIO DE IMPACTO nesta data: 6 resultados masculinos lançados; 5 snapshots existentes no banco
-- inteiro. Reclassificação esperada: 98,4 → N4→N5 · 172/172/189 → N0→N5 · 207 → N0→N4 ·
-- 240 → (não classificado)→N4. Nenhum paciente piora. Snapshots já calculados NÃO são
-- recalculados por esta migration — quem quiser o valor novo precisa gerar snapshot novo.
--
-- PENDENTE, FORA DO ESCOPO DESTA MIGRATION: "Ferritina - Mulheres Pré-Menopausa" (15pt) tem o
-- mesmo defeito (N5=40-81, N0=>150). Em pré-menopausa o teto baixo é menos absurdo (perda
-- menstrual real), mas N0 em >150 continua sem respaldo. Decisão editorial separada.
--
-- Idempotente por id. Down restaura exatamente os valores que estavam em prod em 17/07/2026.

-- +goose Up

-- Ferritina - Homens — partição: ≤15 | (15,30] | (30,50] | (50,200] | (200,300] | >300
UPDATE score_levels SET operator='<=',      lower_limit=NULL,  upper_limit='15',  name='≤15 (deficiência absoluta de ferro)',        updated_at=NOW() WHERE id='019bf31d-2ef0-7940-8243-e49a62ad93d3';
UPDATE score_levels SET operator='between', lower_limit='15',  upper_limit='30',  name='15-30 (depleção de estoques)',               updated_at=NOW() WHERE id='019bf31d-2ef0-740e-9157-62ac9623beed';
UPDATE score_levels SET operator='between', lower_limit='30',  upper_limit='50',  name='30-50 (baixo-normal)',                       updated_at=NOW() WHERE id='019bf31d-2ef0-77aa-98b0-df4ab4657c1f';
UPDATE score_levels SET operator='between', lower_limit='50',  upper_limit='200', name='50-200 (faixa funcional)',                   updated_at=NOW() WHERE id='019bf31d-2ef0-75c8-8e2c-a7fbe9d07155';
UPDATE score_levels SET operator='between', lower_limit='200', upper_limit='300', name='200-300 (alto-normal, vigiar com IST e PCR)', updated_at=NOW() WHERE id='019bf31d-2ef0-740d-a1d4-c4cdd009a79f';
UPDATE score_levels SET operator='>',       lower_limit='300', upper_limit=NULL,  name='>300 (hiperferritinemia, investigar)',       updated_at=NOW() WHERE id='019bf31d-2ef0-71c9-92e1-87fafba5a65f';

-- Troca de papel dos níveis: o extremo BAIXO passa a ser o pior (N0) e o extremo ALTO vai para N1.
UPDATE score_levels SET level=0 WHERE id='019bf31d-2ef0-7940-8243-e49a62ad93d3'; -- ≤15  : N1 -> N0
UPDATE score_levels SET level=1 WHERE id='019bf31d-2ef0-71c9-92e1-87fafba5a65f'; -- >300 : N0 -> N1
UPDATE score_levels SET level=2 WHERE id='019bf31d-2ef0-740e-9157-62ac9623beed'; -- 15-30 : mantém N2
UPDATE score_levels SET level=3 WHERE id='019bf31d-2ef0-77aa-98b0-df4ab4657c1f'; -- 30-50 : mantém N3
UPDATE score_levels SET level=4 WHERE id='019bf31d-2ef0-740d-a1d4-c4cdd009a79f'; -- 200-300 : mantém N4
UPDATE score_levels SET level=5 WHERE id='019bf31d-2ef0-75c8-8e2c-a7fbe9d07155'; -- 50-200 : mantém N5

-- +goose Down

-- Restaura prod 17/07/2026 (níveis e faixas)
UPDATE score_levels SET level=1, operator='<=',      lower_limit=NULL,  upper_limit='15',  name='≤15',     updated_at=NOW() WHERE id='019bf31d-2ef0-7940-8243-e49a62ad93d3';
UPDATE score_levels SET level=0, operator='>',       lower_limit='150', upper_limit=NULL,  name='>150',    updated_at=NOW() WHERE id='019bf31d-2ef0-71c9-92e1-87fafba5a65f';
UPDATE score_levels SET level=2, operator='between', lower_limit='15',  upper_limit='30',  name='15-30',   updated_at=NOW() WHERE id='019bf31d-2ef0-740e-9157-62ac9623beed';
UPDATE score_levels SET level=3, operator='between', lower_limit='30',  upper_limit='45',  name='30-45',   updated_at=NOW() WHERE id='019bf31d-2ef0-77aa-98b0-df4ab4657c1f';
UPDATE score_levels SET level=4, operator='between', lower_limit='80',  upper_limit='150', name='80-150',  updated_at=NOW() WHERE id='019bf31d-2ef0-740d-a1d4-c4cdd009a79f';
UPDATE score_levels SET level=5, operator='between', lower_limit='45',  upper_limit='80',  name='45-80',   updated_at=NOW() WHERE id='019bf31d-2ef0-75c8-8e2c-a7fbe9d07155';
