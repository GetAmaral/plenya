-- BATCH 39 PARTE 1: Composição Corporal - Histórico (16 items)

-- Pré-natal (3x)
UPDATE score_items SET
  clinical_relevance = 'História pré-natal dos pais influencia programação metabólica fetal, predispondo a obesidade/sarcopenia na vida adulta. Estado nutricional materno (desnutrição ou excesso), diabetes gestacional, tabagismo e estresse crônico alteram epigenética do feto, afetando adipogênese, miogênese e regulação hormonal do apetite (leptina, grelina). Filhos de mães com diabetes/obesidade têm maior risco de adiposidade visceral e resistência insulínica. Avaliação do IMC pré-gestacional, ganho de peso gestacional e complicações metabólicas maternas é fundamental.',
  patient_explanation = 'A saúde dos seus pais antes e durante a gestação influenciou sua composição corporal atual. Estado nutricional da mãe, diabetes ou obesidade na gravidez afetam como seu corpo armazena gordura e constrói músculos desde o nascimento.',
  conduct = 'Investigar: IMC materno pré-gestacional, ganho de peso gestacional, diabetes gestacional, complicações perinatais. Considerar risco aumentado de obesidade/sarcopenia para estratégias preventivas.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '6c0651a0-f8a4-47be-95b8-a42a36673de1',
  'b898d743-cca9-48e4-aab9-50c3e04d5b1b',
  '917bc528-e4a1-4a30-9890-d704517e2669'
);

-- Nascimento (3x)
UPDATE score_items SET
  clinical_relevance = 'Peso ao nascer e idade gestacional são preditores de composição corporal futura. Baixo peso (< 2500g) associa-se a sarcopenia, baixa massa muscular e risco aumentado de diabetes tipo 2 ("thrifty phenotype"). Alto peso (> 4000g) predispõe a obesidade e síndrome metabólica. Prematuridade compromete desenvolvimento muscular e ósseo. Complicações perinatais (hipóxia, hipoglicemia) afetam programação metabólica. Avaliar peso/idade gestacional contextualiza vulnerabilidades atuais.',
  patient_explanation = 'O peso ao nascer e se você nasceu prematuro influenciam sua composição corporal. Bebês com baixo peso têm maior risco de ter menos massa muscular. Bebês grandes têm mais chance de obesidade. Conhecer essas informações ajuda a entender seu metabolismo.',
  conduct = 'Documentar: peso ao nascer, idade gestacional, complicações perinatais. Baixo peso: rastreamento precoce de sarcopenia e diabetes. Alto peso: monitoramento de obesidade e síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '74dfa65f-df48-4331-9a47-72fa2d52b5e2',
  '4ba81974-b66a-42c5-92cf-9093d3084e38',
  '5e02310c-08f1-4f88-b22d-4d89447fe565'
);

-- Infância (3x)
UPDATE score_items SET
  clinical_relevance = 'Trajetória de peso/altura na infância prediz composição corporal adulta. Ganho de peso acelerado (catch-up growth) nos primeiros 2 anos associa-se a maior adiposidade visceral e risco metabólico. Obesidade infantil (IMC > p95) predispõe a obesidade adulta e sarcopenia por inflamação crônica. Desnutrição infantil compromete desenvolvimento muscular e ósseo irreversivelmente. Estirões inadequados podem indicar disfunções hormonais. Análise de percentis de crescimento (OMS/CDC) é essencial.',
  patient_explanation = 'Como você cresceu e ganhou peso na infância afeta sua composição corporal hoje. Crianças que ganharam peso muito rápido ou que tiveram obesidade infantil têm maior risco de problemas metabólicos. Compreender sua história de crescimento ajuda a identificar padrões.',
  conduct = 'Revisar: curvas de crescimento infantil (OMS), IMC percentilado, histórico de obesidade/desnutrição. Ganho rápido ou obesidade infantil: monitoramento intensivo de síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '1bdd6114-82c5-4b57-8352-ea32449169ba',
  '64614d75-3b62-48ac-a6e8-294fa906ae2e',
  '84baa5e6-db01-4a90-af2b-801ce9f0670b'
);

-- Adolescência (2x)
UPDATE score_items SET
  clinical_relevance = 'Adolescência é período crítico para desenvolvimento de massa muscular e óssea (pico aos 18-25 anos). Ganho de peso excessivo na puberdade associa-se a obesidade adulta persistente. Estirões inadequados podem indicar deficiências nutricionais ou disfunções hormonais (GH, IGF-1, tireoidianos, esteroides gonadais). Anorexia/bulimia comprometem irreversivelmente saúde óssea e muscular. Obesidade adolescente predispõe a sarcopenia prematura por lipotoxicidade.',
  patient_explanation = 'A adolescência é quando você constrói a maior parte da sua massa muscular e óssea. Problemas de peso nessa fase, como obesidade ou desnutrição, afetam sua saúde metabólica para o resto da vida. Compreender como foi seu crescimento ajuda a personalizar estratégias.',
  conduct = 'Documentar: estirão puberal, IMC na adolescência, transtornos alimentares. Obesidade adolescente: rastreamento de sarcopenia prematura. Anorexia/bulimia: avaliação de densidade óssea (DEXA).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '8a932b13-17c6-42d1-b428-dfe296accd4b',
  '657dcaae-f159-4415-8531-5a2908c0d095'
);

-- Vida adulta (3x)
UPDATE score_items SET
  clinical_relevance = 'Histórico de avaliações de composição corporal (DEXA, bioimpedância, pregas) permite identificar trajetórias de ganho/perda muscular e gordura. Perda muscular > 5-10% em 6-12 meses indica sarcopenia secundária (doença, inatividade, desnutrição). Ganho de gordura visceral progressivo prediz risco metabólico. Variações cíclicas (efeito sanfona) associam-se a perda muscular preferencial e ganho visceral. Comparação longitudinal de ASMI, FMI e gordura visceral orienta intervenções.',
  patient_explanation = 'Se você já fez exames de composição corporal (DEXA, bioimpedância), esses resultados mostram como seu corpo mudou ao longo do tempo. Perdas de massa muscular ou ganhos de gordura visceral indicam riscos metabólicos. Acompanhar essas mudanças permite intervenções mais precisas.',
  conduct = 'Revisar: avaliações prévias de composição corporal (DEXA, bioimpedância). Calcular variações de ASMI, FMI, gordura visceral. Perda muscular > 5-10%: investigar sarcopenia secundária.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  'ac04ab06-bdc7-4959-a33d-32f3646c1a87',
  '3c9b2568-a4ea-4132-9c3e-f327ad1b0ce7',
  '852be041-945c-4ced-91a8-4f85b55e9651'
);

-- Picos de peso (3x)
UPDATE score_items SET
  clinical_relevance = 'Picos de peso/gordura identificam períodos de risco metabólico aumentado e gatilhos ambientais (estresse, rotina, medicamentos). Ganho rápido (> 10% em 6 meses) associa-se a acúmulo preferencial de gordura visceral, resistência insulínica e inflamação crônica. Identificar contextos (gravidez, menopausa, corticoterapia, antipsicóticos, antidepressivos tricíclicos) orienta manejo. Obesidade prévia predispõe a reganho rápido (set point metabólico elevado).',
  patient_explanation = 'Momentos em que você ganhou muito peso rapidamente podem indicar fatores desencadeantes (estresse, medicamentos, mudanças hormonais). Compreender esses padrões ajuda a prevenir novos ganhos de peso e a personalizar estratégias de emagrecimento sustentável.',
  conduct = 'Identificar: gatilhos de ganho rápido (medicamentos, estresse, mudanças hormonais). Ganho > 10% em 6 meses: investigar resistência insulínica (HOMA-IR), gordura visceral (DEXA/TC). Ajustar medicações se necessário.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN (
  '3b9f7a1e-527e-4a4c-bc23-87c68c41ee89',
  'ef3181a6-135a-4110-9275-68eae1c3890a',
  '1af195d6-77c0-495b-9833-9ad38362e6fb'
);

SELECT COUNT(*) as items_atualizados FROM score_items WHERE last_review > CURRENT_TIMESTAMP - INTERVAL '1 minute';
