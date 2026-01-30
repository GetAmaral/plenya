-- Enriquecimento do item FSH - Mulheres Ovulação
-- ID: 915bc591-b263-4fd2-a64d-4a04b38c97f7
-- Data: 2026-01-29

BEGIN;

-- Inserir artigos científicos relacionados
INSERT INTO articles (id, title, authors, journal, publish_date, article_type, original_link, doi, pm_id, abstract, language, created_at, updated_at) VALUES

-- Artigo 1: StatPearls - Physiology, Follicle Stimulating Hormone
('a1f5b001-0c11-4001-a001-a4e1f5b00001'::uuid,
 'Physiology, Follicle Stimulating Hormone',
 'Orlowski M, Sarao MS',
 'StatPearls [Internet]',
 '2023-05-01',
 'review',
 'https://www.ncbi.nlm.nih.gov/books/NBK535442/',
 NULL,
 '30571063',
 'Comprehensive review of FSH physiology detailing its role throughout the menstrual cycle. The article explains that FSH reaches peak levels simultaneously with the LH surge that triggers ovulation. During the follicular phase, when a dominant follicle produces sufficient estradiol (200-300 pg/ml for 48 hours), the hypothalamus responds with a GnRH surge stimulating FSH secretion. The review covers clinical applications including fertility assessment, PCOS diagnosis, ovarian reserve evaluation, and therapeutic uses in assisted reproduction.',
 'en',
 CURRENT_TIMESTAMP,
 CURRENT_TIMESTAMP),

-- Artigo 2: Journal of Ovarian Research - LH Surge Definition
('a1f5b002-0c11-4002-a002-a4e2f5b00002'::uuid,
 'Defining the LH surge in natural cycle frozen-thawed embryo transfer: the role of LH, estradiol, and progesterone',
 'Orvieto R, Morag N, Rubin E, Nahum R',
 'Journal of Ovarian Research',
 '2025-04-14',
 'research_article',
 'https://pmc.ncbi.nlm.nih.gov/articles/PMC11995576/',
 '10.1186/s13048-025-01658-7',
 NULL,
 'Study of 668 natural cycle FET examining precise hormonal changes during the LH surge. Findings showed that successful ovulation involves simultaneous LH increase (>180% baseline), estradiol decrease, and progesterone rise (threefold increase). Patients who conceived showed significantly higher progesterone increases. The research emphasizes the importance of monitoring multiple hormones, including FSH context, for optimal fertility timing and embryo transfer success (52-56.8% pregnancy rates).',
 'en',
 CURRENT_TIMESTAMP,
 CURRENT_TIMESTAMP),

-- Artigo 3: Endotext - The Normal Menstrual Cycle
('a1f5b003-0c11-4003-a003-a4e3f5b00003'::uuid,
 'The Normal Menstrual Cycle and the Control of Ovulation',
 'Reed BG, Carr BR',
 'Endotext [Internet]',
 '2018-08-05',
 'review',
 'https://www.ncbi.nlm.nih.gov/books/NBK279054/',
 NULL,
 NULL,
 'Comprehensive endocrinology textbook chapter detailing menstrual cycle physiology and ovulation control. The authors explain that FSH is elevated during early follicular phase and declines until ovulation, while LH remains low initially but rises during mid-follicular phase due to positive feedback from rising estrogen. The LH surge, initiated by elevated estradiol from the preovulatory follicle, triggers ovulation approximately 10-12 hours after the LH peak.',
 'en',
 CURRENT_TIMESTAMP,
 CURRENT_TIMESTAMP),

-- Artigo 4: Nature Cell Biology - Spatiotemporal Control of Ovulation
('a1f5b004-0c11-4004-a004-a4e4f5b00004'::uuid,
 'Ex vivo imaging reveals the spatiotemporal control of ovulation',
 'Multiple Authors',
 'Nature Cell Biology',
 '2024-10-01',
 'research_article',
 'https://www.nature.com/articles/s41556-024-01524-6',
 '10.1038/s41556-024-01524-6',
 NULL,
 'Advanced imaging study revealing the spatiotemporal dynamics of ovulation. Research demonstrates how preovulatory follicles respond to mid-cycle LH surge, which triggers oocyte maturation and ovulation. In response to FSH, follicles undergo immense growth and remodeling to become preovulatory follicles with fluid-filled antrum. Study shows that FSH drop stops further follicle maturation once estrogen peaks, and ovulation occurs 29-39 hours after LH peak.',
 'en',
 CURRENT_TIMESTAMP,
 CURRENT_TIMESTAMP);

-- Vincular artigos ao score_item FSH - Mulheres Ovulação
INSERT INTO article_score_items (score_item_id, article_id) VALUES
('915bc591-b263-4fd2-a64d-4a04b38c97f7'::uuid, 'a1f5b001-0c11-4001-a001-a4e1f5b00001'::uuid),
('915bc591-b263-4fd2-a64d-4a04b38c97f7'::uuid, 'a1f5b002-0c11-4002-a002-a4e2f5b00002'::uuid),
('915bc591-b263-4fd2-a64d-4a04b38c97f7'::uuid, 'a1f5b003-0c11-4003-a003-a4e3f5b00003'::uuid),
('915bc591-b263-4fd2-a64d-4a04b38c97f7'::uuid, 'a1f5b004-0c11-4004-a004-a4e4f5b00004'::uuid);

-- Atualizar score_item com conteúdo clínico enriquecido
UPDATE score_items SET
  clinical_relevance = 'O FSH (hormônio folículo-estimulante) durante a ovulação desempenha papel fundamental na regulação do ciclo menstrual e na fertilidade feminina. Durante a fase folicular, o FSH estimula o crescimento e maturação dos folículos ovarianos, preparando-os para a liberação do óvulo. Quando um folículo dominante produz níveis suficientes de estradiol (200-300 pg/ml mantidos por 48 horas), ocorre feedback positivo no hipotálamo que desencadeia pico simultâneo de FSH e LH (hormônio luteinizante), conhecido como surge ovulatório. O pico de FSH na ovulação tem funções específicas: libera o oócito das ligações foliculares, estimula ativador de plasminogênio e aumenta receptores de LH nas células da granulosa. A mensuração adequada do FSH durante esta fase permite predição precisa da janela fértil, essencial para planejamento reprodutivo natural ou técnicas de reprodução assistida. Valores normais de FSH durante o surge ovulatório variam entre 5-20 mUI/ml, com pico ocorrendo simultaneamente ao LH. A relação LH:FSH é igualmente importante, mantendo-se normalmente entre 1:1 durante fase folicular e aumentando durante o surge. Alterações nesta dinâmica hormonal podem indicar condições como síndrome dos ovários policísticos (PCOS), insuficiência ovariana primária ou disfunção hipotalâmica. Estudos recentes demonstram que o monitoramento integrado de FSH, LH, estradiol e progesterona oferece predição mais precisa da ovulação, com taxas de concepção melhoradas quando há sincronização adequada destes hormônios.',

  patient_explanation = 'O FSH (hormônio folículo-estimulante) é um hormônio essencial para sua fertilidade e ciclo menstrual. Imagine-o como um "maestro" que coordena o amadurecimento dos óvulos nos seus ovários. Durante a primeira metade do seu ciclo (fase folicular), o FSH estimula os folículos ovarianos a crescerem, preparando óvulos para eventual liberação. Quando um folículo está suficientemente maduro e produz bastante estrogênio, seu corpo responde com um "pico hormonal" - tanto FSH quanto LH aumentam rapidamente juntos. Este pico acontece cerca de 24-36 horas antes da ovulação e marca sua janela mais fértil do mês. Após o pico de FSH, acontecem três coisas importantes: o óvulo se solta do folículo, o folículo se prepara para romper e liberar o óvulo, e seu corpo fica pronto para possível gravidez. Valores normais de FSH durante a ovulação ficam entre 5-20 mUI/ml. Se seu FSH está sempre muito alto ou muito baixo, pode dificultar a gravidez. Por isso, entender e acompanhar seu FSH durante a ovulação ajuda você e seu médico a identificar o melhor momento para tentar engravidar, seja naturalmente ou com auxílio de tratamentos de fertilidade. Testes modernos permitem detectar este pico hormonal em casa, ajudando a identificar seus dias mais férteis do ciclo.',

  conduct = 'AVALIAÇÃO CLÍNICA DO FSH NA OVULAÇÃO:\n\n1. INDICAÇÕES PARA MENSURAÇÃO:\n- Investigação de infertilidade (após 12 meses de tentativas sem sucesso)\n- Planejamento de ciclos de reprodução assistida (FIV, IA)\n- Identificação precisa da janela fértil\n- Monitoramento de indução de ovulação\n- Avaliação de disfunções ovulatórias\n- Diagnóstico diferencial de amenorreia secundária\n\n2. TIMING DA COLETA:\n- FSH basal: dias 2-4 do ciclo menstrual\n- FSH ovulatório: monitoramento seriado a partir do dia 10-12 do ciclo\n- Coletas seriadas com intervalos de 24-48h para detectar surge\n- Considerar ultrassonografia transvaginal concomitante para confirmar crescimento folicular\n\n3. INTERPRETAÇÃO DOS VALORES:\n- FSH ovulatório normal: 5-20 mUI/ml\n- Surge de FSH: aumento significativo em relação ao basal\n- Relação LH:FSH durante surge: LH predomina (LH aumenta >180% do basal)\n- Avaliar simultaneamente: estradiol (pico 200-300 pg/ml), progesterona (aumento ~3x)\n\n4. CONDUTA CONFORME ACHADOS:\n\nFSH ELEVADO NA FASE FOLICULAR (>10-15 mUI/ml):\n- Sugestivo de diminuição da reserva ovariana\n- Solicitar contagem de folículos antrais (CFA) e hormônio antimülleriano (AMH)\n- Considerar encaminhamento precoce para reprodução assistida\n- Discutir prognóstico reprodutivo com a paciente\n\nFSH BAIXO (<2-3 mUI/ml):\n- Investigar disfunção hipotalâmica-hipofisária\n- Avaliar histórico de perda ponderal, exercício excessivo, estresse\n- Solicitar prolactina, TSH, hormônios tireoidianos\n- Considerar teste de estímulo com GnRH\n\nSURGE OVULATÓRIO INADEQUADO:\n- Ausência de pico hormonal detectável: anovulação\n- Considerar citrato de clomifeno ou letrozol para indução\n- Monitorar resposta com ultrassonografia seriada\n- Programar coito ou inseminação baseado no surge detectado\n\n5. ORIENTAÇÕES À PACIENTE:\n- Explicar importância do timing das coletas\n- Orientar sobre janela fértil (2 dias antes até dia da ovulação)\n- Ensinar uso de testes de LH urinário domiciliares como complemento\n- Discutir fatores que afetam FSH: idade, IMC, tabagismo, reserva ovariana\n\n6. SEGUIMENTO:\n- Reavaliação após 3-6 ciclos de tentativas programadas\n- Considerar reprodução assistida se idade >35 anos e FSH alterado\n- Monitoramento anual em mulheres com reserva ovariana diminuída',

  last_review = CURRENT_TIMESTAMP,
  updated_at = CURRENT_TIMESTAMP

WHERE id = '915bc591-b263-4fd2-a64d-4a04b38c97f7';

COMMIT;

-- Verificação da atualização
SELECT
  si.id,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '915bc591-b263-4fd2-a64d-4a04b38c97f7'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
