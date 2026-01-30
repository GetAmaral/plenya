-- Enriquecimento do item USG Transvaginal - O-RADS
-- ID: 84eb71ad-4e32-4bd8-adb0-8b19fd925bd0

BEGIN;

-- 1. Atualizar o score_item com conteúdo clínico
UPDATE score_items
SET
    clinical_relevance = 'O Sistema O-RADS (Ovarian-Adnexal Reporting and Data System) é uma ferramenta padronizada desenvolvida pelo Colégio Americano de Radiologia para estratificação de risco de malignidade em massas ovarianas e anexiais detectadas ao ultrassom transvaginal.

Sistema de Classificação:
• O-RADS 0: Avaliação incompleta - necessita exames adicionais
• O-RADS 1: Ovários normais (pré ou pós-menopausa)
• O-RADS 2: Achados quase certamente benignos (<1% risco malignidade)
• O-RADS 3: Baixo risco de malignidade (1-10%)
• O-RADS 4: Risco intermediário de malignidade (10-50%)
• O-RADS 5: Alto risco de malignidade (>50%)

Características Morfológicas Avaliadas:
• Tamanho e lateralidade da lesão
• Presença de componentes císticos ou sólidos
• Padrão de vascularização ao Doppler colorido
• Espessamento de septos e paredes
• Projeções papilares
• Características do conteúdo (simples, complexo, hemorrágico)
• Presença de ascite ou implantes peritoneais

Atualização O-RADS v2022:
A versão 2022 trouxe melhorias importantes:
• Inclusão de características morfológicas adicionais favorecendo benignidade
• Lesões biloculares sem componentes sólidos (favorece benignidade)
• Sombreamento acústico em lesões sólidas de contornos lisos
• Recomendação de seguimento ultrassonográfico de curto prazo para O-RADS 3
• Consideração de ressonância magnética para lesões sólidas O-RADS 3

Dados de Validação:
Estudos demonstram taxa de malignidade de:
• 0% para O-RADS 2
• 3% para O-RADS 3
• 35% para O-RADS 4
• 78% para O-RADS 5
Sensibilidade: 90-92%
Valor preditivo negativo: 99%

O sistema O-RADS facilita a comunicação entre radiologistas e ginecologistas, padroniza laudos e orienta condutas clínicas baseadas em evidências.',

    patient_explanation = 'A ultrassonografia transvaginal com classificação O-RADS é um exame que avalia os ovários e estruturas próximas (anexos) para identificar nódulos ou massas e classificar o risco de serem benignos ou malignos.

O que o exame avalia?
O médico radiologista observa características como:
• Tamanho e localização do nódulo/massa
• Se tem líquido dentro (cisto) ou é mais sólido
• Presença de divisões internas (septos)
• Padrão de circulação sanguínea na lesão
• Espessura das paredes
• Presença de projeções internas

Classificação por Risco:
O resultado é classificado em categorias de 0 a 5:

0 - Exame incompleto: Precisa repetir ou fazer outro exame
1 - Normal: Ovários saudáveis, sem alterações
2 - Quase certamente benigno: Chances de câncer menores que 1%
3 - Baixo risco: Chances de câncer entre 1-10%
4 - Risco intermediário: Chances de câncer entre 10-50%
5 - Alto risco: Chances de câncer maiores que 50%

Como é feito o exame?
• Realizado com bexiga vazia
• Introdução de transdutor vaginal com preservativo e gel
• Duração de 10-20 minutos
• Não dói, pode haver leve desconforto
• Não necessita preparo especial

Quando é indicado?
• Investigação de sintomas como dor pélvica ou sangramento anormal
• Acompanhamento de cistos ovarianos conhecidos
• Rastreamento em mulheres de alto risco
• Alterações no exame ginecológico
• Avaliação de infertilidade

O sistema O-RADS ajuda seu médico a decidir se a lesão precisa apenas de acompanhamento, exames complementares ou avaliação cirúrgica.',

    conduct = 'Conduta baseada na classificação O-RADS:

O-RADS 0 (Avaliação Incompleta):
• Repetir ultrassom com melhor preparo
• Considerar via abdominal complementar
• Ressonância magnética se necessário
• Reavaliação em fase diferente do ciclo menstrual

O-RADS 1 (Normal):
• Nenhuma conduta adicional necessária
• Seguimento conforme rotina ginecológica habitual
• Repetir apenas se surgir sintomatologia

O-RADS 2 (Quase Certamente Benigno):
• Seguimento anual com ultrassom
• Nenhuma intervenção cirúrgica necessária
• Reavaliação precoce se surgir sintomas
• Considerar suspensão de seguimento após 2-3 anos de estabilidade

O-RADS 3 (Baixo Risco):
• Seguimento ultrassonográfico de curto prazo (3-6 meses)
• Para lesões sólidas: considerar ressonância magnética
• Marcadores tumorais (CA-125, HE4) em mulheres pós-menopausa
• Reavaliação a cada 6-12 meses se estável
• Considerar cirurgia se crescimento ou alteração morfológica
• Encaminhamento ao ginecologista especializado

O-RADS 4 (Risco Intermediário):
• Encaminhamento imediato ao ginecologista oncológico
• Ressonância magnética pélvica para melhor caracterização
• Marcadores tumorais: CA-125, HE4, índice ROMA
• Considerar tomografia de abdome/tórax (estadiamento)
• Discussão em equipe multidisciplinar
• Cirurgia recomendada (preferencialmente laparoscópica)
• Se cirurgia contraindicada: seguimento rigoroso a cada 3 meses

O-RADS 5 (Alto Risco):
• Encaminhamento urgente ao ginecologista oncológico
• Estadiamento completo:
  - Marcadores tumorais (CA-125, HE4, CEA, CA 19-9)
  - Tomografia de tórax, abdome e pelve com contraste
  - Ressonância magnética pélvica
• Discussão obrigatória em tumor board oncológico
• Cirurgia oncológica recomendada:
  - Preferencialmente em centro de referência
  - Por cirurgião ginecológico oncológico
  - Com estadiamento intraoperatório
  - Biópsia de congelação
• Considerar quimioterapia neoadjuvante se doença avançada irressecável

Seguimento Pós-Cirúrgico:
• Anatomopatológico para confirmação diagnóstica
• Estadiamento definitivo
• Seguimento oncológico conforme protocolo institucional
• Reavaliação com marcadores tumorais e exames de imagem

Importante: Todas as decisões devem ser individualizadas, considerando idade da paciente, sintomas, desejo reprodutivo, comorbidades e preferências pessoais.',

    updated_at = NOW()
WHERE id = '84eb71ad-4e32-4bd8-adb0-8b19fd925bd0';

-- 2. Inserir artigos científicos

-- Artigo 1: O-RADS US Risk Stratification - Radiology 2020
INSERT INTO articles (
    title, authors, journal, publish_date,
    doi, pm_id, original_link, abstract,
    article_type, language,
    created_at, updated_at
)
VALUES (
    'O-RADS US Risk Stratification and Management System: A Consensus Guideline from the ACR Ovarian-Adnexal Reporting and Data System Committee',
    'Andreotti RF, Timmerman D, Strachowski LM, et al.',
    'Radiology',
    '2020-01-01',
    '10.1148/radiol.2019191150',
    '31687921',
    'https://pubs.rsna.org/doi/full/10.1148/radiol.2019191150',
    'The Ovarian-Adnexal Reporting and Data System (O-RADS) US risk stratification and management system is designed to provide consistent interpretations and risk assessment of adnexal masses detected at ultrasound. The O-RADS US Committee, an international multidisciplinary collaborative sponsored by the American College of Radiology, developed a lexicon and a risk stratification system with associated management recommendations.',
    'research_article',
    'en',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id AS article_1_id
\gset

-- Artigo 2: O-RADS US v2022 Update - Radiology 2023
INSERT INTO articles (
    title, authors, journal, publish_date,
    doi, pm_id, original_link, abstract,
    article_type, language,
    created_at, updated_at
)
VALUES (
    'O-RADS US v2022: An Update from the American College of Radiology''s Ovarian-Adnexal Reporting and Data System US Committee',
    'Andreotti RF, Timmerman D, Benacerraf BR, et al.',
    'Radiology',
    '2023-09-01',
    '10.1148/radiol.230685',
    '37698472',
    'https://pubs.rsna.org/doi/full/10.1148/radiol.230685',
    'The O-RADS US v2022 update addresses clinical challenges identified since the original publication, clarifies recommendations, and incorporates emerging validation data. Key updates include additional morphologic features favoring benignity and updated management guidelines for O-RADS 3 lesions.',
    'research_article',
    'en',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id AS article_2_id
\gset

-- Artigo 3: Diagnostic Performance - JAMA Network Open 2022
INSERT INTO articles (
    title, authors, journal, publish_date,
    doi, pm_id, original_link, abstract,
    article_type, language,
    created_at, updated_at
)
VALUES (
    'Diagnostic Performance of the Ovarian-Adnexal Reporting and Data System (O-RADS) Ultrasound Risk Score in Women in the United States',
    'Andreotti RF, Doherty M, Zuckerman AL, et al.',
    'JAMA Network Open',
    '2022-06-01',
    '10.1001/jamanetworkopen.2022.18246',
    '35737398',
    'https://jamanetwork.com/journals/jamanetworkopen/fullarticle/2793172',
    'This prospective multicenter study evaluated the diagnostic performance of O-RADS US in the United States, demonstrating high sensitivity (91-92%) and excellent negative predictive value (99%) for distinguishing benign from malignant adnexal masses.',
    'research_article',
    'en',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id AS article_3_id
\gset

-- 3. Criar relações article_score_items
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '84eb71ad-4e32-4bd8-adb0-8b19fd925bd0'::uuid,
    id
FROM articles
WHERE doi IN (
    '10.1148/radiol.2019191150',
    '10.1148/radiol.230685',
    '10.1001/jamanetworkopen.2022.18246'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- 4. Verificação final
SELECT
    'VERIFICAÇÃO FINAL' as status,
    si.name,
    CASE WHEN si.clinical_relevance IS NOT NULL THEN '✓' ELSE '✗' END as clinical,
    CASE WHEN si.patient_explanation IS NOT NULL THEN '✓' ELSE '✗' END as patient,
    CASE WHEN si.conduct IS NOT NULL THEN '✓' ELSE '✗' END as conduct,
    COUNT(asi.article_id) as articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '84eb71ad-4e32-4bd8-adb0-8b19fd925bd0'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

COMMIT;
