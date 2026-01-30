-- Ajustar clinical_relevance para atingir contagem mínima de caracteres

UPDATE score_items
SET clinical_relevance = 'A Globulina Ligadora de Hormônios Sexuais (SHBG) é uma proteína produzida pelo fígado que se liga com alta afinidade à testosterona e estradiol, regulando a biodisponibilidade desses hormônios. Em mulheres, níveis diminuídos de SHBG (<40 nmol/L) estão fortemente associados à Síndrome dos Ovários Policísticos (SOP), resistência à insulina, síndrome metabólica e doença hepática gordurosa não alcoólica. A produção hepática de SHBG é suprimida pela insulina e por andrógenos, tornando-a um marcador indireto sensível de resistência insulínica e hiperandrogenismo. Valores baixos aumentam a fração livre de testosterona, intensificando os sinais clínicos de hiperandrogenismo (hirsutismo, acne, alopecia androgênica, irregularidades menstruais). O Índice de Andrógenos Livres (FAI), calculado como (Testosterona Total/SHBG) × 100, é utilizado para avaliar o status androgênico, com valores >5 indicando hiperandrogenismo. Em mulheres com SOP, a SHBG serve como biomarcador precoce de distúrbios metabólicos e pode orientar decisões terapêuticas. Contraceptivos hormonais combinados aumentam significativamente os níveis de SHBG através do componente estrogênico, reduzindo andrógenos livres e melhorando manifestações clínicas de hiperandrogenismo. Valores elevados de SHBG (>114 nmol/L) podem ocorrer em uso de estrogênios, hipertireoidismo, anorexia nervosa ou hepatopatias, reduzindo a biodisponibilidade hormonal. A dosagem de SHBG é essencial na avaliação de irregularidades menstruais, infertilidade, hiperandrogenismo e risco cardiometabólico em mulheres.'
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';

-- Verificar contagem final
SELECT
  'clinical_relevance' as field,
  LENGTH(clinical_relevance) as char_count,
  CASE
    WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
    ELSE '✗ FORA DO RANGE'
  END as status
FROM score_items
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';
