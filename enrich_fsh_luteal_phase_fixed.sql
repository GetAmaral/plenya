-- Enrichment for FSH - Mulheres Fase Lútea
-- Score Item ID: 0726b373-4b78-4cb8-91a9-0916681aef61
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, doi, article_type, created_at, updated_at)
VALUES
  (
    gen_random_uuid(),
    'Current ovulation and luteal phase tracking methods and technologies for fertility and family planning: a review',
    'Wegrzynowicz AK, Eyvazzadeh A, Beckley A',
    'Seminars in Reproductive Medicine',
    '2024-09-20',
    '10.1055/s-0044-1791190',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11837971/',
    'review',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Physiology, Follicle Stimulating Hormone',
    'Stamler R, Kapoor N, Ioachimescu AG',
    'StatPearls',
    '2024-08-12',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK535442/',
    'review',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'The Normal Menstrual Cycle and the Control of Ovulation',
    'Reed BG, Carr BR',
    'Endotext',
    '2024-06-15',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK279054/',
    'review',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Evidence-based guideline: Premature Ovarian Insufficiency',
    'American Society for Reproductive Medicine',
    'Fertility and Sterility',
    '2025-01-10',
    NULL,
    'https://www.asrm.org/practice-guidance/practice-committee-documents/evidence-based-guideline-premature-ovarian-insufficiency--2024/',
    'clinical_guideline',
    NOW(),
    NOW()
  );

-- Link articles to the score item
INSERT INTO article_score_items (score_item_id, article_id, created_at, updated_at)
SELECT
  '0726b373-4b78-4cb8-91a9-0916681aef61'::uuid,
  id,
  NOW(),
  NOW()
FROM articles
WHERE title IN (
  'Current ovulation and luteal phase tracking methods and technologies for fertility and family planning: a review',
  'Physiology, Follicle Stimulating Hormone',
  'The Normal Menstrual Cycle and the Control of Ovulation',
  'Evidence-based guideline: Premature Ovarian Insufficiency'
)
AND created_at >= NOW() - INTERVAL '1 minute';

-- Update the score item with clinical content
UPDATE score_items
SET
  clinical_relevance = 'O FSH (Hormônio Folículo Estimulante) na fase lútea apresenta importância clínica significativa na avaliação da função ovariana e fertilidade feminina. Durante esta fase do ciclo menstrual, que ocorre após a ovulação e dura aproximadamente 14 dias, os níveis de FSH permanecem relativamente baixos e suprimidos devido ao feedback negativo exercido pelo estrogênio e progesterona produzidos pelo corpo lúteo. Os valores de referência para FSH na fase lútea em mulheres em idade reprodutiva situam-se tipicamente entre 1,1 a 9,2 mUI/mL, embora diferentes laboratórios possam apresentar pequenas variações metodológicas (1,2 a 9,0 mUI/mL ou 1,5 a 12,0 UI/L). A manutenção de níveis baixos de FSH durante a fase lútea é essencial para prevenir o desenvolvimento de novos folículos enquanto o corpo prepara-se para uma potencial gravidez. Clinicamente, a mensuração do FSH na fase lútea tem aplicações limitadas em comparação com a fase folicular, pois a avaliação da reserva ovariana é mais fidedigna quando realizada entre os dias 2-7 do ciclo. No entanto, a elevação anormal do FSH durante a fase lútea pode indicar insuficiência do corpo lúteo, comprometimento da função ovariana ou transição para insuficiência ovariana prematura. Ao final da fase lútea, quando não ocorre fecundação, a queda na produção de esteroides pelo corpo lúteo e a redução dramática de inibina A permitem a elevação do FSH, preparando o início de um novo ciclo menstrual. Este padrão dinâmico hormonal é fundamental para a compreensão da fisiologia reprodutiva e para o diagnóstico diferencial de distúrbios da fertilidade. Valores persistentemente elevados de FSH podem sugerir diminuição da reserva ovariana, mesmo em mulheres jovens, enquanto valores inapropriadamente baixos podem indicar disfunção hipotalâmica ou hipofisária. A interpretação deve sempre considerar o contexto clínico completo, incluindo idade, história menstrual, sintomas associados e outros marcadores hormonais como LH, estradiol e progesterona. Diretrizes recentes da ASRM (2025) estabelecem que uma única elevação de FSH >25 UI/L é suficiente para diagnóstico de insuficiência ovariana prematura quando combinada com quadro clínico característico.',

  patient_explanation = 'O FSH (Hormônio Folículo Estimulante) é um hormônio produzido pela glândula hipófise no cérebro que desempenha papel fundamental na regulação do seu ciclo menstrual e fertilidade. Durante a fase lútea do ciclo menstrual, que é o período de aproximadamente 14 dias após a ovulação e antes da próxima menstruação, os níveis de FSH naturalmente ficam mais baixos. Isso acontece porque, após você ovular, forma-se no ovário uma estrutura chamada corpo lúteo, que produz progesterona e estrogênio. Esses hormônios "avisam" ao cérebro que não é necessário produzir muito FSH neste momento, pois o corpo está aguardando para ver se ocorrerá gravidez. Os valores normais de FSH na fase lútea geralmente ficam entre 1,1 a 9,2 mUI/mL, dependendo do laboratório. Níveis baixos de FSH nesta fase são esperados e saudáveis, indicando que seu sistema reprodutivo está funcionando adequadamente. No entanto, se o FSH estiver muito elevado durante esta fase, pode sinalizar que seus ovários não estão produzindo hormônios suficientes ou que a função ovariana está comprometida. Ao final da fase lútea, quando não há gravidez, os níveis de progesterona e estrogênio caem, e o FSH naturalmente começa a subir novamente, preparando o corpo para recrutar novos folículos e iniciar um novo ciclo. Este padrão hormonal rítmico é essencial para a saúde reprodutiva. É importante saber que a dosagem de FSH para avaliar fertilidade geralmente é feita no início do ciclo (dias 2-7), não na fase lútea, pois fornece informações mais precisas sobre a reserva de óvulos nos ovários. Seu médico interpretará seus resultados considerando sua idade, sintomas, histórico menstrual e outros exames hormonais para fornecer orientações personalizadas sobre sua saúde reprodutiva.',

  conduct = 'PROTOCOLO CLÍNICO PARA INTERPRETAÇÃO DO FSH NA FASE LÚTEA:\n\n1. COLETA E TIMING:\n- Confirmar que a paciente está de fato na fase lútea (dia 15-28 do ciclo, ou 1-14 dias pós-ovulação)\n- Idealmente confirmar ovulação prévia através de: ultrassonografia seriada, teste de LH urinário positivo prévio, ou temperatura basal corporal bifásica\n- Coleta matinal preferencial (8h-10h) para reduzir variabilidade circadiana\n- Jejum não obrigatório, mas recomendado de 3-4 horas\n- Evitar coleta se paciente usar anticoncepcional hormonal ou terapia de reposição hormonal\n\n2. VALORES DE REFERÊNCIA E INTERPRETAÇÃO:\n- NORMAL: 1,1 a 9,2 mUI/mL (pode variar conforme laboratório: 1,2-9,0 ou 1,5-12,0 UI/L)\n- FSH BAIXO (<1,0 mUI/mL): Investigar disfunção hipotalâmica/hipofisária, considerar prolactinoma, síndrome de Kallmann, ou uso de medicamentos supressores\n- FSH NORMAL (1,1-9,2 mUI/mL): Função lútea adequada, corpo lúteo funcionante, prognóstico favorável para fertilidade\n- FSH LIMÍTROFE (9,3-12,0 mUI/mL): Zona cinzenta, repetir exame em fase folicular precoce (dia 2-5), considerar AMH e contagem de folículos antrais\n- FSH ELEVADO (>12,0 mUI/mL): Sugestivo de insuficiência lútea, diminuição de reserva ovariana ou transição perimenopáusica; solicitar FSH em fase folicular\n- FSH MUITO ELEVADO (>25,0 UI/L): Conforme ASRM 2025, uma única elevação é suficiente para diagnóstico de Insuficiência Ovariana Prematura (IOP) quando associada a quadro clínico\n\n3. INVESTIGAÇÃO COMPLEMENTAR OBRIGATÓRIA:\n- LH sérico (avaliar relação FSH/LH)\n- Estradiol sérico (descartar supressão inadequada)\n- Progesterona sérica (confirmar ovulação e adequação lútea; valores >3 ng/mL confirmam ovulação, >10 ng/mL indicam fase lútea robusta)\n- TSH e prolactina (descartar causas secundárias de disfunção ovariana)\n- Para pacientes com FSH elevado: repetir FSH + estradiol em fase folicular precoce (dia 2-5), solicitar AMH e ultrassonografia transvaginal para contagem de folículos antrais (CFA)\n\n4. CENÁRIOS CLÍNICOS ESPECÍFICOS:\n\nA) FSH ELEVADO NA FASE LÚTEA + CICLOS IRREGULARES:\n- Investigar transição perimenopáusica (idade >35 anos)\n- Solicitar AMH (<1,0 ng/mL sugere reserva ovariana diminuída)\n- Avaliar sintomas vasomotores, alterações humor, secura vaginal\n- Considerar IOP se idade <40 anos\n\nB) FSH NORMAL + DEFICIÊNCIA LÚTEA SUSPEITADA:\n- Medir progesterona no meio da fase lútea (dia 21 ciclo de 28 dias, ou 7 dias pós-ovulação)\n- Progesterona <10 ng/mL sugere deficiência lútea\n- Ultrassonografia para avaliar espessura endometrial\n- Considerar suplementação com progesterona se história de perdas gestacionais precoces\n\nC) FSH BAIXO + AMENORREIA/OLIGOMENORREIA:\n- Investigar eixo hipotálamo-hipófise: RM de sela túrcica\n- Avaliar histórico de exercício físico intenso, baixo peso, estresse\n- Descartar hiperprolactinemia, hipotireoidismo\n- Considerar hipogonadismo hipogonadotrófico\n\n5. LIMITAÇÕES E CONSIDERAÇÕES:\n- FSH na fase lútea tem MENOR valor preditivo para fertilidade comparado à fase folicular\n- Avaliação de reserva ovariana deve priorizar: FSH basal (dia 2-5), AMH e CFA\n- FSH isolado não prediz resposta ovariana em tratamentos de reprodução assistida\n- Variabilidade inter e intralaboratorial pode atingir 10-15%\n- Sempre interpretar no contexto clínico: idade, IMC, história familiar, exposições ambientais\n\n6. SEGUIMENTO:\n- FSH limítrofe/elevado: repetir em 1-2 meses na fase folicular\n- FSH normal com sintomas persistentes: considerar outros marcadores (AMH, inibina B)\n- Documentar tendência ao longo do tempo em mulheres >35 anos\n- Encaminhar para especialista em reprodução humana se: FSH >12 mUI/mL em mulher <35 anos desejando gravidez, ou FSH >25 UI/L em qualquer idade',

  updated_at = NOW()
WHERE id = '0726b373-4b78-4cb8-91a9-0916681aef61';

COMMIT;

-- Verification query
SELECT
  si.id,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  COUNT(sia.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items sia ON si.id = sia.score_item_id
WHERE si.id = '0726b373-4b78-4cb8-91a9-0916681aef61'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;
