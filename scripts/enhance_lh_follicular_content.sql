-- Enhanced content for LH - Mulheres Fase Folicular
-- Expanding clinical_relevance and patient_explanation to meet character targets

BEGIN;

UPDATE score_items
SET
  clinical_relevance = 'O hormônio luteinizante (LH) na fase folicular é um marcador fundamental da função ovariana e do eixo hipotálamo-hipófise-gonadal. Durante a fase folicular inicial (dias 1-7 do ciclo), os níveis de LH permanecem relativamente baixos (1,9-12,5 mIU/mL), aumentando gradualmente em resposta à elevação do estradiol produzido pelos folículos em desenvolvimento. A avaliação do LH na fase folicular inicial é essencial para o diagnóstico de diversas condições endócrinas: na síndrome dos ovários policísticos (SOP), frequentemente observa-se elevação do LH com razão LH/FSH >2:1 em cerca de 60% das pacientes, embora este critério não seja diagnóstico isoladamente devido à variabilidade individual e metodológica. Na amenorreia hipotalâmica funcional, a razão LH/FSH é tipicamente ≤1,0 em mais de 80% dos casos, refletindo supressão do eixo hipotálamo-hipofisário por estresse, exercício excessivo ou restrição calórica. A avaliação combinada do LH com FSH permite estimar a reserva ovariana: quando o FSH está desproporcionalmente elevado em relação ao LH na fase folicular inicial, sugere-se diminuição da reserva ovariana com recrutamento de folículos primordiais remanescentes. O LH também é crucial na avaliação pré-tratamento de reprodução assistida, onde níveis profundamente suprimidos durante estimulação ovariana controlada (<0,5 mIU/mL) podem comprometer a qualidade oocitária e as taxas de nascidos vivos. Estudos recentes demonstram que o ambiente intrafolicular com concentrações reduzidas de estradiol secundárias à deficiência iatrogênica de LH pode levar à maturação oocitária subótima durante a fase folicular, destacando a importância do equilíbrio hormonal adequado. A interpretação do LH deve considerar a idade da paciente, IMC, uso de medicações (anticoncepcionais orais, metformina, corticoides), fase exata do ciclo menstrual e método analítico utilizado pelo laboratório.',

  patient_explanation = 'O hormônio luteinizante (LH) é produzido pela glândula hipófise no cérebro e tem um papel essencial no funcionamento dos seus ovários. Na primeira metade do ciclo menstrual (fase folicular), que vai do primeiro dia da menstruação até a ovulação, o LH trabalha junto com outro hormônio chamado FSH para fazer os folículos (pequenas estruturas que contêm os óvulos) crescerem e amadurecerem. Nesta fase, o LH normalmente fica em níveis baixos a moderados (entre 1,9 e 12,5 mIU/mL). Este exame é especialmente importante para investigar irregularidades menstruais, dificuldade para engravidar, síndrome dos ovários policísticos (SOP) e avaliar se os ovários estão respondendo adequadamente aos sinais do cérebro. Quando o LH está muito elevado em relação ao FSH, pode indicar SOP, uma condição hormonal comum que afeta cerca de 10% das mulheres em idade reprodutiva. Quando está muito baixo, pode sugerir que o cérebro não está enviando sinais adequados para os ovários, situação frequentemente associada a estresse extremo, dietas muito restritivas ou exercícios físicos excessivos. O exame deve ser feito entre o 2º e 5º dia do ciclo menstrual para resultados mais precisos. É importante fazer o exame junto com outros hormônios (FSH, estradiol, prolactina, TSH, testosterona) para uma avaliação completa da sua saúde reprodutiva e hormonal. Medicamentos como anticoncepcionais, corticoides e alguns tratamentos para diabetes podem afetar os resultados, assim como variações de peso, estresse e exercícios muito intensos.',

  last_review = CURRENT_DATE
WHERE id = 'c8d0d878-43a0-4788-834e-62578a4c5657';

-- Verify enhanced character counts
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = 'c8d0d878-43a0-4788-834e-62578a4c5657';

COMMIT;
