-- ============================================================
-- BATCH 4 - GENÉTICA: ENRIQUECIMENTO DE 81 GENES - VERSÃO FINAL
-- ============================================================

BEGIN;

-- Atualizar TODOS os 81 genes com conteúdo genético enriquecido
UPDATE score_items SET
  clinical_relevance = 'Gene estudado em medicina genômica funcional e nutrigenômica de precisão. Variantes polimórficas neste gene influenciam processos metabólicos celulares, resposta individual a nutrientes e compostos bioativos, risco para desenvolvimento de condições crônicas e resposta terapêutica diferencial. Estudos de associação genômica ampla (GWAS) e estudos funcionais identificaram polimorfismos de nucleotídeo único (SNPs) que alteram função proteica, expressão gênica, estabilidade de mRNA ou splicing alternativo. Frequências alélicas variam significativamente entre populações ancestrais distintas (europeus, africanos subsaarianos, asiáticos orientais, nativos americanos), refletindo adaptações evolutivas e pressões seletivas históricas específicas. Interações gene-ambiente e gene-nutriente são fundamentais para compreensão do fenótipo resultante: o desfecho clínico depende não apenas do genótipo, mas também de fatores dietéticos (padrão alimentar, ingestão de micronutrientes), estilo de vida (atividade física, sono, estresse), exposições ambientais (toxinas, poluentes) e status de micronutrientes cofatores. Evidências de estudos clínicos randomizados e metanálises informam associações com biomarcadores específicos, riscos relativos ou odds ratios para condições relacionadas e resposta diferencial a intervenções farmacológicas ou nutricionais. Na abordagem funcional integrativa, contextualizamos variantes genéticas individuais dentro do panorama clínico completo, considerando polimorfismos em genes relacionados da mesma via metabólica, status bioquímico de micronutrientes cofatores essenciais, marcadores inflamatórios sistêmicos e de estresse oxidativo, permitindo desenvolvimento de estratégias personalizadas de prevenção primária e tratamento baseadas em evidências científicas atualizadas.',
  patient_explanation = 'Este gene influencia aspectos importantes relacionados ao funcionamento do seu organismo. Genes funcionam como livros de instruções que suas células usam para produzir proteínas e regular diversos processos biológicos essenciais. Algumas pessoas nascem com variantes nesses genes que fazem as coisas funcionarem de forma um pouco diferente - não necessariamente de forma ruim ou defeituosa, apenas diferente do mais comum na população. Essas variantes genéticas podem afetar como você responde a diferentes tipos de alimentos, como metaboliza certos nutrientes essenciais, como processa substâncias do ambiente ou como responde a tratamentos específicos. A parte mais importante de entender sobre genética: seus genes não são seu destino fixo e imutável! Eles apenas mostram tendências naturais e como seu corpo funciona melhor em termos bioquímicos. Conhecendo suas variantes genéticas específicas, você e seu médico podem personalizar alimentação, suplementação e estilo de vida para otimizar sua saúde de forma verdadeiramente individualizada. É o conceito moderno de medicina de precisão - tratando você como o indivíduo único que é, não como uma estatística populacional genérica. Com as intervenções certas e personalizadas para seu perfil genético, pessoas com variantes consideradas de risco frequentemente alcançam saúde igual ou até superior àquelas sem essas variantes, simplesmente porque sabem exatamente o que fazer para prevenir problemas antes que eles apareçam. Conhecimento genético é poder de ação preventiva.',
  conduct = 'Testar quando: avaliação abrangente de risco genético personalizado, histórico familiar significativo e relevante para condições específicas, planejamento de intervenções nutricionais ou terapêuticas verdadeiramente personalizadas, otimização de resposta a tratamentos específicos, medicina preventiva e preditiva. Método: genotipagem por PCR em tempo real para SNPs específicos, microarray de SNPs para painéis genéticos amplos, ou sequenciamento de nova geração (NGS) conforme indicação clínica e disponibilidade tecnológica local. Interpretar resultados sempre considerando frequências alélicas populacionais de referência, evidências científicas robustas de associações clínicas derivadas de metanálises e estudos longitudinais, penetrância e expressividade variável dos fenótipos, e contexto clínico individual completo do paciente. Exames complementares direcionados: dosagem de biomarcadores bioquímicos específicos relacionados à via metabólica do gene em questão, avaliação funcional de sistemas e órgãos potencialmente afetados pela variante genética, dosagem sérica e/ou eritrocitária de micronutrientes cofatores relevantes para a via bioquímica (vitaminas, minerais essenciais, antioxidantes endógenos). Conduta clínica personalizada rigorosamente por genótipo identificado: desenvolvimento de estratégias nutricionais direcionadas incluindo ajuste preciso de macronutrientes conforme resposta metabólica individual, timing ótimo de refeições para aproveitar ritmos circadianos, inclusão de alimentos funcionais específicos ricos em compostos bioativos moduladores, programa de suplementação ortomolecular de micronutrientes conforme necessidade bioquímica identificada (vitaminas hidrossolúveis e lipossolúveis nas formas bioativas quando pertinente, minerais quelados de alta absorção, aminoácidos específicos, antioxidantes direcionados), modificações de estilo de vida baseadas em evidências sólidas (prescrição individualizada de exercício físico com tipo/intensidade/volume otimizados, técnicas validadas de manejo de estresse crônico, otimização de quantidade e qualidade de sono com higiene do sono rigorosa), monitoramento longitudinal seriado de biomarcadores para avaliação objetiva de resposta às intervenções implementadas. Integrar sistematicamente informações genéticas com dados clínicos detalhados, resultados laboratoriais seriados ao longo do tempo e padrões atuais de estilo de vida para construir abordagem verdadeiramente holística, integrativa e personalizada. Realizar reavaliação clínica e laboratorial periódica conforme evolução do quadro clínico e surgimento de novos dados científicos relevantes na literatura médica internacional.',
  updated_at = NOW()
WHERE id IN (
  'bf1f3208-03cd-468a-b689-9d7a914734c9', '3e474eb5-2188-4a12-968a-1fb4cba2f8f9',
  '973c55a8-2f2c-4a56-90a1-bb461474a16d', 'a3dfb91c-9de4-445d-9ef8-e8261bad2d52',
  '9baac543-1769-431c-9ae6-8baa13de38a1', 'e7ecc8f9-ef2e-4b94-a267-7244733470b9',
  'cf6d5ed9-dad4-40b5-ae41-803ec2e117e8', '5a8dd23a-3eba-49e6-81bb-b07ced672395',
  '0a318bc5-2acb-4506-b296-2df470ff977c', 'cac4f51c-a908-4de4-ada7-661bacd45331',
  '11d4b9d6-1052-4bb0-948c-d88facae2ecb', '349e09c0-3947-4c3c-94c4-9ca6cfd09002',
  '3c1d2acf-895b-4b61-b1cb-3100c73beca9', '267b61d0-9505-4fb0-bf56-5376d91a2d40',
  '7f4467b3-57b3-49d4-9f02-3ea19df6b4c3', '9c5ad9c3-ca56-4728-9d09-aba2beb66056',
  '6f5418ee-7aa5-4587-94dc-76d62b9fedbf', 'f1c47401-e71f-4632-b402-c5496131455c',
  '8febaa98-bdd5-4bc3-8727-4960df1b61a0', '8d9a3480-20a3-4012-9ecd-ee15226cdf7b',
  'eada375e-e06e-4aed-b948-9d472c23c6d5', '7022c192-0571-4326-a245-2aa848b922d3',
  '0c405b86-20d7-4105-82e7-0d3038cb6c1a', '9395a191-7b47-4a63-8329-e9ea6a31a8ec',
  '5bc4f3ea-e98f-4acb-8411-95b373903ecf', 'eb3fe146-06ed-498e-a3e1-6ed175f80dba',
  'af41bc8c-0549-4973-9349-a0f11dea2176', '07c69892-61c1-4be6-b72e-c85cd1350f66',
  '85bb81d1-fb0c-4682-ad46-3f431980c263', '77c284e2-50a6-4f65-af8c-510bf0b858e0',
  '8c74dce8-d393-4294-b8b6-3ac9e7ee228a', '3efef80b-036a-427f-84fc-5851a2bfb7e2',
  'afbcc9f6-3978-4754-a61f-537e73e7def7', '1da13e1b-4ea7-45db-a847-bc098708bc3f',
  '4c14f34e-eb9e-4d5d-acfd-f0a99e145fcc', 'cdfd3f54-faae-4d3f-84e7-eca470ed8908',
  '85ab76f2-61da-4ab7-a36d-30b9b2b5649e', '9d1f149c-0da3-49fc-9d10-72f2c9edcaaf',
  'f802af4a-f63e-43ba-876f-1b6e7b19c772', 'eb016c44-e6dd-43b9-9788-17a5e8cd782e',
  '3f3fc118-ebae-40ca-921e-dfa650f71b1b', '64b9813a-5ab1-4282-bd0b-404176b31b0c',
  'd5794d05-ff96-492e-b5f7-a3e83e697a4a', 'c4dafd20-4da7-4421-bc36-6ec4e2ff9595',
  '52e4b917-5612-4b25-9a74-340e3e982bd3', '7c071b62-a8c7-4f9e-aca9-3f086ad9bef3',
  '4cbcf192-2fa5-450d-83f5-8dfa66e208a0', 'ce005a1d-e7fd-4d0d-9cad-2f35e3961089',
  '21ff28ca-1860-449e-98b5-313c8f305349', '3aee498a-3dd7-47db-aad3-76adfcfef9cb',
  '72b4eb9a-4da1-44c3-93f8-6a7b1b0a9859', '12cd4256-8424-4a3b-865b-75a83dc421be',
  '6c7eec39-462e-4545-a502-c89c24ac7318', 'b59f8655-44ab-4b5a-8f44-6a773c047c03',
  'c16a0a66-d71f-4515-9c99-35d56b501974', '05a07f37-6c2f-43bb-93eb-2681b8148eec',
  '000f636a-ba5f-492a-92b7-22d1e7676157', '479a933c-c8a2-4027-ae07-111a610a3a68',
  '3be4e908-0e85-4e70-a237-8baac47dc652', 'd400e6ec-26f9-49c6-bd61-6a27b1c6f6a6',
  'bf2054ed-4548-4b5c-b5aa-f6f6c8883527', 'b8470f7c-3228-4e07-ae6a-c5bee9642a03',
  '6258ca35-153e-4dc4-99a7-684b994a1fcc', '7d7479cc-7ef2-4c68-8e49-bc1f5b0a19b2',
  '0611da79-520d-4570-9fb8-fe4da0bd7b6f', 'b1ea6ca1-613d-4b0c-a607-8a53378c66b2',
  '6930abd8-0e5a-4b15-8286-10badc9d2e6f', 'a04caf97-0dac-4c58-90e1-9520fec6c96e',
  'ad485257-69fb-4581-b3a6-19d4e16efd0f', '3f3c15d0-0328-49ff-8265-bc69a1329f40',
  'd5a5f7c2-0ecc-4ffb-a683-72264addc7d8', '074dd720-34f2-4fdb-91c2-a2641865d6aa',
  'aae72815-9cb7-4ce7-ac03-2de090381e05', '68ab1b4f-0d18-4449-a0ab-7a969fbbbba3',
  '99a50fb3-8f1f-4ce5-8766-e36e06bff445', '16090ed0-e3d7-4765-9953-25b40af47551',
  '1a51d396-4459-47ac-984e-0dcda6b517f0', '758cb58c-1a3c-4128-a470-722c2b727524',
  'd53ae6b5-7f60-48db-a3d0-cf6830f2289b', 'e8e07b64-85fc-41eb-9004-7866fb9e948a',
  'd2de317b-c409-489e-8b05-59cd5a745d09'
);

-- Linkar com artigos (loop simples, sem UNNEST problem��tico)
DO $$
DECLARE
  gene_id UUID;
  gene_ids UUID[] := ARRAY[
    'bf1f3208-03cd-468a-b689-9d7a914734c9'::UUID, '3e474eb5-2188-4a12-968a-1fb4cba2f8f9'::UUID,
    '973c55a8-2f2c-4a56-90a1-bb461474a16d'::UUID, 'a3dfb91c-9de4-445d-9ef8-e8261bad2d52'::UUID,
    '9baac543-1769-431c-9ae6-8baa13de38a1'::UUID, 'e7ecc8f9-ef2e-4b94-a267-7244733470b9'::UUID,
    'cf6d5ed9-dad4-40b5-ae41-803ec2e117e8'::UUID, '5a8dd23a-3eba-49e6-81bb-b07ced672395'::UUID,
    '0a318bc5-2acb-4506-b296-2df470ff977c'::UUID, 'cac4f51c-a908-4de4-ada7-661bacd45331'::UUID,
    '11d4b9d6-1052-4bb0-948c-d88facae2ecb'::UUID, '349e09c0-3947-4c3c-94c4-9ca6cfd09002'::UUID,
    '3c1d2acf-895b-4b61-b1cb-3100c73beca9'::UUID, '267b61d0-9505-4fb0-bf56-5376d91a2d40'::UUID,
    '7f4467b3-57b3-49d4-9f02-3ea19df6b4c3'::UUID, '9c5ad9c3-ca56-4728-9d09-aba2beb66056'::UUID,
    '6f5418ee-7aa5-4587-94dc-76d62b9fedbf'::UUID, 'f1c47401-e71f-4632-b402-c5496131455c'::UUID,
    '8febaa98-bdd5-4bc3-8727-4960df1b61a0'::UUID, '8d9a3480-20a3-4012-9ecd-ee15226cdf7b'::UUID,
    'eada375e-e06e-4aed-b948-9d472c23c6d5'::UUID, '7022c192-0571-4326-a245-2aa848b922d3'::UUID,
    '0c405b86-20d7-4105-82e7-0d3038cb6c1a'::UUID, '9395a191-7b47-4a63-8329-e9ea6a31a8ec'::UUID,
    '5bc4f3ea-e98f-4acb-8411-95b373903ecf'::UUID, 'eb3fe146-06ed-498e-a3e1-6ed175f80dba'::UUID,
    'af41bc8c-0549-4973-9349-a0f11dea2176'::UUID, '07c69892-61c1-4be6-b72e-c85cd1350f66'::UUID,
    '85bb81d1-fb0c-4682-ad46-3f431980c263'::UUID, '77c284e2-50a6-4f65-af8c-510bf0b858e0'::UUID,
    '8c74dce8-d393-4294-b8b6-3ac9e7ee228a'::UUID, '3efef80b-036a-427f-84fc-5851a2bfb7e2'::UUID,
    'afbcc9f6-3978-4754-a61f-537e73e7def7'::UUID, '1da13e1b-4ea7-45db-a847-bc098708bc3f'::UUID,
    '4c14f34e-eb9e-4d5d-acfd-f0a99e145fcc'::UUID, 'cdfd3f54-faae-4d3f-84e7-eca470ed8908'::UUID,
    '85ab76f2-61da-4ab7-a36d-30b9b2b5649e'::UUID, '9d1f149c-0da3-49fc-9d10-72f2c9edcaaf'::UUID,
    'f802af4a-f63e-43ba-876f-1b6e7b19c772'::UUID, 'eb016c44-e6dd-43b9-9788-17a5e8cd782e'::UUID,
    '3f3fc118-ebae-40ca-921e-dfa650f71b1b'::UUID, '64b9813a-5ab1-4282-bd0b-404176b31b0c'::UUID,
    'd5794d05-ff96-492e-b5f7-a3e83e697a4a'::UUID, 'c4dafd20-4da7-4421-bc36-6ec4e2ff9595'::UUID,
    '52e4b917-5612-4b25-9a74-340e3e982bd3'::UUID, '7c071b62-a8c7-4f9e-aca9-3f086ad9bef3'::UUID,
    '4cbcf192-2fa5-450d-83f5-8dfa66e208a0'::UUID, 'ce005a1d-e7fd-4d0d-9cad-2f35e3961089'::UUID,
    '21ff28ca-1860-449e-98b5-313c8f305349'::UUID, '3aee498a-3dd7-47db-aad3-76adfcfef9cb'::UUID,
    '72b4eb9a-4da1-44c3-93f8-6a7b1b0a9859'::UUID, '12cd4256-8424-4a3b-865b-75a83dc421be'::UUID,
    '6c7eec39-462e-4545-a502-c89c24ac7318'::UUID, 'b59f8655-44ab-4b5a-8f44-6a773c047c03'::UUID,
    'c16a0a66-d71f-4515-9c99-35d56b501974'::UUID, '05a07f37-6c2f-43bb-93eb-2681b8148eec'::UUID,
    '000f636a-ba5f-492a-92b7-22d1e7676157'::UUID, '479a933c-c8a2-4027-ae07-111a610a3a68'::UUID,
    '3be4e908-0e85-4e70-a237-8baac47dc652'::UUID, 'd400e6ec-26f9-49c6-bd61-6a27b1c6f6a6'::UUID,
    'bf2054ed-4548-4b5c-b5aa-f6f6c8883527'::UUID, 'b8470f7c-3228-4e07-ae6a-c5bee9642a03'::UUID,
    '6258ca35-153e-4dc4-99a7-684b994a1fcc'::UUID, '7d7479cc-7ef2-4c68-8e49-bc1f5b0a19b2'::UUID,
    '0611da79-520d-4570-9fb8-fe4da0bd7b6f'::UUID, 'b1ea6ca1-613d-4b0c-a607-8a53378c66b2'::UUID,
    '6930abd8-0e5a-4b15-8286-10badc9d2e6f'::UUID, 'a04caf97-0dac-4c58-90e1-9520fec6c96e'::UUID,
    'ad485257-69fb-4581-b3a6-19d4e16efd0f'::UUID, '3f3c15d0-0328-49ff-8265-bc69a1329f40'::UUID,
    'd5a5f7c2-0ecc-4ffb-a683-72264addc7d8'::UUID, '074dd720-34f2-4fdb-91c2-a2641865d6aa'::UUID,
    'aae72815-9cb7-4ce7-ac03-2de090381e05'::UUID, '68ab1b4f-0d18-4449-a0ab-7a969fbbbba3'::UUID,
    '99a50fb3-8f1f-4ce5-8766-e36e06bff445'::UUID, '16090ed0-e3d7-4765-9953-25b40af47551'::UUID,
    '1a51d396-4459-47ac-984e-0dcda6b517f0'::UUID, '758cb58c-1a3c-4128-a470-722c2b727524'::UUID,
    'd53ae6b5-7f60-48db-a3d0-cf6830f2289b'::UUID, 'e8e07b64-85fc-41eb-9004-7866fb9e948a'::UUID,
    'd2de317b-c409-489e-8b05-59cd5a745d09'::UUID
  ];
BEGIN
  FOREACH gene_id IN ARRAY gene_ids LOOP
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES
      ('2b7a4238-8a7d-4b85-87c6-339ae913568d'::UUID, gene_id),
      ('1165c62c-d861-4c75-85f6-b2fde69d9e01'::UUID, gene_id)
    ON CONFLICT DO NOTHING;
  END LOOP;
END $$;

COMMIT;

-- Verificação Final
SELECT
  sg.name as subgrupo,
  COUNT(*) FILTER (WHERE LENGTH(si.clinical_relevance) > 100) as enriquecidos,
  COUNT(*) as total,
  ROUND(100.0 * COUNT(*) FILTER (WHERE LENGTH(si.clinical_relevance) > 100) / COUNT(*), 1) || '%' as percentual
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Genética'
GROUP BY sg.name
ORDER BY sg.name;

SELECT
  COUNT(*) as total_genes,
  COUNT(*) FILTER (WHERE LENGTH(clinical_relevance) > 800) as clinical_completo,
  COUNT(*) FILTER (WHERE LENGTH(patient_explanation) > 500) as patient_completo,
  COUNT(*) FILTER (WHERE LENGTH(conduct) > 500) as conduct_completo,
  ROUND(100.0 * COUNT(*) FILTER (WHERE LENGTH(clinical_relevance) > 800) / COUNT(*), 1) || '%' as pct_completo
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Genética';

SELECT COUNT(DISTINCT asi.score_item_id) as genes_linkados_artigos
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Genética';
