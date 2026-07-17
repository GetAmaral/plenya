-- Conserto de 2 itens do catálogo cujas faixas não são aplicáveis a nenhum paciente real.
--
-- 1) Alfa-1 Globulina (g/dL) — faixas ~5x fora de escala. A alfa-1 sérica fisiológica é
--    0,2-0,4 g/dL (fração eletroforética; majoritariamente alfa-1-antitripsina + orosomucoide).
--    Com N5=1,00-2,20 g/dL, NENHUM paciente pode atingir N5 — todo mundo cai em N0 (<=0,80).
--    A estrutura já estava certa (N0=extremo baixo = deficiência de A1AT, o desfecho grave;
--    N1=extremo alto = resposta de fase aguda); só os números estavam fora de escala.
--
-- 2) Relação T3 Livre:T3 Reverso (ratio) — faixas incompatíveis com a própria fórmula do item
--    (unit_conversion = "FT3 (pg/mL) ÷ rT3 (ng/dL)"). Com FT3 em pg/mL, um adulto saudável
--    (FT3~3,0 pg/mL, rT3~15 ng/dL) dá razão ~0,20 — que cai em N0 (<=6) na faixa atual.
--    As faixas só fazem sentido com FT3 em pg/dL (x100). Optei por preservar a fórmula
--    (que está no unit_conversion e no pôster, e é a convenção publicada) e reescalar as
--    faixas /100, alinhando o item à literatura: razão ~0,20 = ótimo.
--
-- Ambos os itens têm lab_test_code (PLN930C1388 / PLN777A5542) e portanto seriam
-- auto-classificados a partir de resultado de exame. Raio de impacto na data desta migration:
-- ZERO resultados lançados para qualquer um dos dois — nenhum paciente foi classificado por
-- eles até aqui, e nenhum snapshot existente muda.
--
-- Convenção de fronteira: 'between' é meio-aberto (> lower e <= upper) e '<'/'<=' guardam o
-- limite em upper_limit — ver models.ScoreLevel.EvaluatesTrue. As faixas abaixo particionam
-- o domínio sem gap nem sobreposição.
--
-- Idempotente por id. Down restaura exatamente os valores que estavam em prod em 17/07/2026.

-- +goose Up

-- Alfa-1 Globulina — partição: <=0,15 | (0,15-0,20] | (0,20-0,35] | (0,35-0,45] | >0,45
UPDATE score_levels SET operator='<=',      lower_limit=NULL,  upper_limit='0.15', name='≤0,15 (deficiência de alfa-1-antitripsina)', updated_at=NOW() WHERE id='019bf31d-2ef0-7bd3-9b2f-61de03b4b9e8';
UPDATE score_levels SET operator='between', lower_limit='0.15', upper_limit='0.20', name='0,15-0,20 (baixo)',                          updated_at=NOW() WHERE id='019bf31d-2ef0-7a51-a09a-40f6535d9c2b';
UPDATE score_levels SET operator='between', lower_limit='0.20', upper_limit='0.35', name='0,20-0,35',                                  updated_at=NOW() WHERE id='019bf31d-2ef0-7ba1-b1cd-6bd8fc03f3c9';
UPDATE score_levels SET operator='between', lower_limit='0.35', upper_limit='0.45', name='0,35-0,45 (limítrofe alto)',                 updated_at=NOW() WHERE id='019bf31d-2ef0-7a17-91fa-7400176b0667';
UPDATE score_levels SET operator='>',       lower_limit='0.45', upper_limit=NULL,   name='>0,45 (resposta de fase aguda)',             updated_at=NOW() WHERE id='019bf31d-2ef0-7180-b467-583f1bb80e45';

-- Relação T3 Livre:T3 Reverso — reescala /100 p/ casar com FT3 (pg/mL) ÷ rT3 (ng/dL)
-- partição: <=0,06 | (0,06-0,10] | (0,10-0,15] | (0,15-0,20] | (0,20-0,30] | >0,30
UPDATE score_levels SET operator='<=',      lower_limit=NULL,   upper_limit='0.06', name='≤0,06', updated_at=NOW() WHERE id='019bf7af-ba01-7eec-9a45-00cc800268b3';
UPDATE score_levels SET operator='between', lower_limit='0.06', upper_limit='0.10', name='0,06-0,10', updated_at=NOW() WHERE id='019bf7af-ba01-74fb-9ea3-8004b2f5eb7e';
UPDATE score_levels SET operator='between', lower_limit='0.10', upper_limit='0.15', name='0,10-0,15', updated_at=NOW() WHERE id='019bf7af-ba01-787f-9861-1b45f8de98af';
UPDATE score_levels SET operator='between', lower_limit='0.15', upper_limit='0.20', name='0,15-0,20', updated_at=NOW() WHERE id='019bf7af-ba01-72c3-aeff-3bf9b9adf526';
UPDATE score_levels SET operator='between', lower_limit='0.20', upper_limit='0.30', name='0,20-0,30', updated_at=NOW() WHERE id='019bf7af-ba01-7a9b-9bc7-c9a4d1f4fad9';
UPDATE score_levels SET operator='>',       lower_limit='0.30', upper_limit=NULL,   name='>0,30', updated_at=NOW() WHERE id='019bf7af-ba01-723d-b86a-d021d001d72b';

-- +goose Down

-- Alfa-1 Globulina — restaura prod 17/07/2026
UPDATE score_levels SET operator='<=',      lower_limit=NULL,   upper_limit='0.80', name='≤0,8',    updated_at=NOW() WHERE id='019bf31d-2ef0-7bd3-9b2f-61de03b4b9e8';
UPDATE score_levels SET operator='between', lower_limit='0.80', upper_limit='1.00', name='0,8-1',   updated_at=NOW() WHERE id='019bf31d-2ef0-7a51-a09a-40f6535d9c2b';
UPDATE score_levels SET operator='between', lower_limit='1.00', upper_limit='2.20', name='1-2,2',   updated_at=NOW() WHERE id='019bf31d-2ef0-7ba1-b1cd-6bd8fc03f3c9';
UPDATE score_levels SET operator='between', lower_limit='2.20', upper_limit='2.50', name='2,2-2,5', updated_at=NOW() WHERE id='019bf31d-2ef0-7a17-91fa-7400176b0667';
UPDATE score_levels SET operator='>',       lower_limit='2.50', upper_limit=NULL,   name='>2,5',    updated_at=NOW() WHERE id='019bf31d-2ef0-7180-b467-583f1bb80e45';

-- Relação T3 Livre:T3 Reverso — restaura prod 17/07/2026
UPDATE score_levels SET operator='<=',      lower_limit=NULL,   upper_limit='6.0',  name='≤6',      updated_at=NOW() WHERE id='019bf7af-ba01-7eec-9a45-00cc800268b3';
UPDATE score_levels SET operator='between', lower_limit='6.0',  upper_limit='10',   name='6-10',    updated_at=NOW() WHERE id='019bf7af-ba01-74fb-9ea3-8004b2f5eb7e';
UPDATE score_levels SET operator='between', lower_limit='10.0', upper_limit='15',   name='10-15',   updated_at=NOW() WHERE id='019bf7af-ba01-787f-9861-1b45f8de98af';
UPDATE score_levels SET operator='between', lower_limit='15.0', upper_limit='20.1', name='15-20,1', updated_at=NOW() WHERE id='019bf7af-ba01-72c3-aeff-3bf9b9adf526';
UPDATE score_levels SET operator='between', lower_limit='20.1', upper_limit='30',   name='20,1-30', updated_at=NOW() WHERE id='019bf7af-ba01-7a9b-9bc7-c9a4d1f4fad9';
UPDATE score_levels SET operator='>',       lower_limit='30',   upper_limit=NULL,   name='>30',     updated_at=NOW() WHERE id='019bf7af-ba01-723d-b86a-d021d001d72b';
