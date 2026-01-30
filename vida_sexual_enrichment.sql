-- ============================================================================
-- ENRIQUECIMENTO DE SCORE ITEMS - GRUPO VIDA SEXUAL
-- Data: 2026-01-27
-- Items: 2 items do grupo "Vida Sexual"
-- ============================================================================

-- ----------------------------------------------------------------------------
-- Item 1: Épocas de pior libido/desempenho
-- ID: 49588a39-e4cf-4f4d-b8a7-fc83c68d28ee
-- Subgrupo: Histórico
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A anamnese temporal da função sexual é fundamental na medicina integrativa, pois permite identificar gatilhos e fatores moduláveis associados a períodos de disfunção sexual. Estudos demonstram que a libido e o desempenho sexual masculino são regulados primariamente por testosterona, dopamina e óxido nítrico, enquanto nas mulheres envolvem estrogênio, testosterona e serotonina de forma mais complexa. A investigação de épocas específicas de declínio pode revelar: períodos de estresse crônico com hipercortisolemia (que suprime eixo gonadal), uso de medicamentos serotonérgicos (SSRIs causam disfunção em até 70% dos usuários), alterações metabólicas (resistência insulínica e síndrome metabólica afetam função endotelial), distúrbios do sono (reduzem testosterona em 10-15% por noite mal dormida) e mudanças hormonais fisiológicas (perimenopausa, andropausa). A correlação temporal com eventos de vida, mudanças ocupacionais, início/término de medicamentos e alterações de peso corporal permite traçar plano terapêutico personalizado focando em fatores reversíveis através de intervenções em estilo de vida, suporte nutricional e modulação hormonal quando apropriado.',
  patient_explanation = 'Entender quando você teve mais dificuldades com sua vida sexual ajuda a identificar o que pode estar afetando seu desejo e desempenho. Muitos fatores podem influenciar: períodos de muito estresse no trabalho ou vida pessoal, noites mal dormidas, ganho de peso, uso de certos remédios (especialmente antidepressivos), mudanças hormonais naturais da idade, ou problemas de relacionamento. Não é "só psicológico" - sua saúde sexual está intimamente ligada a hormônios, circulação sanguínea, qualidade do sono, alimentação e níveis de estresse. Ao mapear os períodos mais difíceis, conseguimos investigar o que mudou naquela época e trabalhar para melhorar esses aspectos através de mudanças no estilo de vida, nutrição adequada, manejo do estresse e, quando necessário, suporte hormonal natural. O objetivo é recuperar sua vitalidade sexual de forma sustentável.',
  conduct = 'Investigar detalhadamente períodos específicos: correlacionar com eventos de vida (mudanças profissionais, relacionais), início de medicamentos (especialmente anti-hipertensivos, antidepressivos, finasterida), alterações ponderais, padrões de sono e níveis de atividade física. Solicitar avaliação hormonal basal (testosterona total e livre, SHBG, estradiol, prolactina, TSH, cortisol). Avaliar marcadores metabólicos (glicemia, insulina, perfil lipídico, HbA1c) e inflamatórios (PCR-us). Rastrear apneia do sono se roncos/sonolência diurna. Considerar trial de suplementação com vitamina D, zinco, magnésio se deficiências. Enfatizar intervenções em estilo de vida: exercício resistido, manejo de estresse, higiene do sono.',
  updated_at = NOW()
WHERE id = '49588a39-e4cf-4f4d-b8a7-fc83c68d28ee';

-- Artigos científicos já linkados ao Item 1 (15 artigos):
-- - Hormonal Regulation of Men's Sexual Desire and Arousal: ICSM 2024 Recommendations
-- - Understanding the Role of Serotonin in Female Hypoactive Sexual Desire Disorder and Treatment Options
-- - Dopamine, Erectile Function and Male Sexual Behavior: A Review
-- - Dehydroepiandrosterone and Cortisol as Markers of HPA Axis Dysregulation in Women with Low Sexual Desire
-- - Hormone Therapy for Sexual Function in Perimenopausal and Postmenopausal Women
-- - Integrative Medicine for Erectile Dysfunction: A Holistic Approach
-- - DISFUNÇÃO ERÉTIL (Equipe MFI)
-- - Fisiologia Endócrina Feminina (Equipe MFI)
-- - Terapia de Reposição Hormonal com Testosterona XII (Equipe MFI)
-- - Ritmo Circadiano Eixo HPA Partes I, II, III (Equipe MFI)
-- - Neurologia Funcional Integrativa 1, 2, 3 (Equipe MFI)

-- ----------------------------------------------------------------------------
-- Item 2: Uso recente outros medicamentos/suplementos para libido/desempenho sexual
-- ID: d35c6ad4-cd19-4839-96b6-08260315e87c
-- Subgrupo: Atual
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A documentação de uso prévio ou atual de medicamentos e suplementos para função sexual é essencial para avaliação de risco-benefício e planejamento terapêutico. Inibidores de PDE5 (sildenafila, tadalafila, vardenafila) são eficazes em 60-80% dos casos de disfunção erétil, mas exigem cautela em cardiopatias e uso de nitratos. Afrodisíacos naturais têm evidências variáveis: ginseng coreano, L-arginina, maca peruana e tribulus terrestris apresentam estudos promissores mas heterogêneos; DHEA pode beneficiar mulheres pós-menopáusicas com baixa libido mas riscos de conversão androgênica devem ser monitorados. Fitoterápicos como ginkgo biloba, açafrão (Crocus sativus) e feno-grego demonstram efeitos modestos. É crucial avaliar interações medicamentosas (yohimbina com antidepressivos pode causar crises hipertensivas), qualidade/pureza de suplementos (contaminação com PDE5 sintéticos é comum), e expectativas realistas. A abordagem integrativa prioriza correção de deficiências nutricionais (vitamina D, zinco, magnésio), otimização hormonal fisiológica, e terapias comprovadas antes de afrodisíacos controversos. Histórico de falhas terapêuticas orienta investigação mais profunda de causas vasculares, neurológicas ou psicogênicas.',
  patient_explanation = 'Saber quais remédios ou suplementos você já tentou para melhorar sua vida sexual ajuda a evitar repetir tratamentos ineficazes e identificar o melhor caminho. Medicamentos como Viagra e Cialis funcionam bem para ereção mas não aumentam desejo sexual - se o problema é libido baixa, eles podem não ajudar. Muitos suplementos naturais vendidos prometem milagres mas têm pouca comprovação científica, e alguns são adulterados com substâncias sintéticas perigosas. Alguns suplementos têm evidências razoáveis: ginseng coreano, maca peruana, L-arginina podem ajudar modestamente; DHEA pode beneficiar mulheres após menopausa. Porém, antes de usar qualquer suplemento, é mais eficaz corrigir deficiências básicas (vitamina D, zinco, magnésio são essenciais para hormônios sexuais), melhorar circulação com exercícios, reduzir estresse e dormir bem. Se já tentou várias coisas sem sucesso, pode indicar necessidade de investigação mais profunda.',
  conduct = 'Documentar medicamentos/suplementos utilizados: nome, dose, duração, resposta percebida e efeitos adversos. Investigar especificamente: inibidores de PDE5 (dose adequada? uso correto 30-60min antes?), testosterona exógena (gel, injeções - verificar níveis séricos e efeitos colaterais), DHEA, tribulus, maca, ginseng, L-arginina. Alertar sobre riscos de suplementos não regulamentados (contaminação com análogos sintéticos de PDE5). Avaliar interações com medicamentos de uso contínuo (anti-hipertensivos, anticoagulantes, antidepressivos). Revisar evidências científicas de cada substância utilizada. Priorizar otimização de fatores basais antes de terapias avançadas: corrigir deficiências nutricionais (solicitar 25-OH vitamina D, zinco, magnésio), melhorar composição corporal, exercício regular. Considerar encaminhamento para urologia/ginecologia se falhas terapêuticas múltiplas.',
  updated_at = NOW()
WHERE id = 'd35c6ad4-cd19-4839-96b6-08260315e87c';

-- Artigos científicos já linkados ao Item 2 (15 artigos):
-- - Efficacy and Safety of Common Ingredients in Aphrodisiacs Used for Erectile Dysfunction: A Review
-- - Integrative Medicine for Erectile Dysfunction: A Holistic Approach
-- - Hormonal Regulation of Men's Sexual Desire and Arousal: ICSM 2024 Recommendations
-- - Hormone Therapy for Sexual Function in Perimenopausal and Postmenopausal Women
-- - Dopamine, Erectile Function and Male Sexual Behavior: A Review
-- - Understanding the Role of Serotonin in Female Hypoactive Sexual Desire Disorder and Treatment Options
-- - DISFUNÇÃO ERÉTIL (Equipe MFI)
-- - Fisiologia Endócrina Feminina (Equipe MFI)
-- - Terapia de Reposição Hormonal com Testosterona XII (Equipe MFI)
-- - Ritmo Circadiano Eixo HPA Partes I, II, III (Equipe MFI)
-- - Neurologia Funcional Integrativa 1, 2, 3 (Equipe MFI)

-- ============================================================================
-- VERIFICAÇÃO
-- ============================================================================

-- Verificar atualização dos textos
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_length,
  LENGTH(patient_explanation) as patient_length,
  LENGTH(conduct) as conduct_length,
  updated_at
FROM score_items
WHERE id IN ('49588a39-e4cf-4f4d-b8a7-fc83c68d28ee', 'd35c6ad4-cd19-4839-96b6-08260315e87c');

-- Verificar artigos linkados
SELECT
  si.name as item_name,
  COUNT(asi.article_id) as num_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id IN ('49588a39-e4cf-4f4d-b8a7-fc83c68d28ee', 'd35c6ad4-cd19-4839-96b6-08260315e87c')
GROUP BY si.name;

-- ============================================================================
-- ESTATÍSTICAS FINAIS
-- ============================================================================
-- Items enriquecidos: 2/2 (100%)
-- Artigos científicos linkados: 15 por item
-- Texto Clinical Relevance: ~1.200 caracteres
-- Texto Patient Explanation: ~900 caracteres
-- Texto Conduct: ~800 caracteres
-- ============================================================================
