-- ============================================================
-- ENRIQUECIMENTO: ECG - QTc (Bazett) - Homens
-- ID: 3a0d7e6b-c53d-47de-a50c-d7774e542835
-- Grupo: Exames > Imagem
-- Data: 2026-01-28
-- ============================================================

-- 1. ATUALIZAR SCORE ITEM COM CONTEÚDO CLÍNICO
UPDATE score_items
SET
    clinical_relevance = 'O intervalo QT corrigido (QTc) pela fórmula de Bazett representa a duração da despolarização e repolarização ventricular ajustada pela frequência cardíaca. Em homens adultos, valores normais situam-se abaixo de 450 ms. QTc entre 431-450 ms é considerado limítrofe, enquanto valores acima de 450 ms indicam prolongamento anormal. O prolongamento do QTc (>450 ms em homens) está associado a risco aumentado de arritmias ventriculares graves, especialmente Torsades de Pointes (TdP), que pode evoluir para fibrilação ventricular e morte súbita cardíaca. A cada 10 ms de aumento no QTc, o risco de TdP aumenta 5-7%. Valores ≥500 ms representam alto risco de eventos arrítmicos fatais, com TdP sendo incomum apenas quando QTc <500 ms. O prolongamento pode ser congênito (Síndrome do QT Longo) ou adquirido (medicamentos, distúrbios eletrolíticos, bradiarritmias). Monitoramento rigoroso é essencial em pacientes com QTc prolongado, especialmente na presença de fatores de risco como hipocalemia, hipomagnesemia, sexo feminino, idade avançada, insuficiência cardíaca e uso de múltiplas drogas que prolongam o QT.',

    patient_explanation = 'O QTc (intervalo QT corrigido) é uma medida do tempo que seu coração leva para se preparar eletricamente para o próximo batimento, ajustada pela sua frequência cardíaca. Valores normais para homens ficam abaixo de 450 milissegundos (ms). Quando o QTc está prolongado (acima de 450 ms), indica que o coração está levando mais tempo que o normal para se "recarregar" eletricamente, o que aumenta o risco de desenvolver arritmias perigosas – batimentos cardíacos irregulares que podem causar tonturas, desmaios ou, em casos graves, morte súbita. Valores muito elevados (acima de 500 ms) representam alto risco e exigem atenção médica imediata. O prolongamento pode ser causado por medicamentos, baixos níveis de potássio ou magnésio no sangue, ou condições genéticas. Seu médico monitorará seu QTc e ajustará tratamentos conforme necessário para manter você seguro.',

    conduct = 'QTc <430 ms: Normal. Nenhuma conduta específica necessária. Manter seguimento de rotina. QTc 431-450 ms (limítrofe): Revisar medicações em uso (antiarrítmicos, antipsicóticos, antibióticos macrolídeos, antifúngicos azólicos); corrigir distúrbios eletrolíticos (K+, Mg2+, Ca2+); avaliar função tireoidiana; considerar monitoramento periódico com ECG; orientar paciente sobre sintomas de alerta (síncope, palpitações); evitar novos fármacos que prolonguem QT. QTc 451-470 ms (prolongado leve): Reduzir ou suspender drogas que prolonguem QT quando possível; correção agressiva de eletrólitos (K+ >4.5 mEq/L, Mg2+ >2.0 mg/dL); ECG de controle em 2-4 semanas; avaliação cardiológica para investigação de causas secundárias; considerar teste genético se história familiar sugestiva; orientar evitar esforços extremos. QTc 471-499 ms (prolongado moderado): Suspender imediatamente todas drogas que prolonguem QT; internação hospitalar para monitoramento contínuo se sintomas (síncope, pré-síncope); telemetria cardíaca contínua por 24-48h; reposição IV de K+ e Mg2+ se deficiência; ecocardiograma transtorácico; teste ergométrico (se assintomático); encaminhamento urgente ao cardiologista/eletrofisiologista; considerar marca-passo temporário se bradicardia associada. QTc ≥500 ms (alto risco de TdP): EMERGÊNCIA MÉDICA. Internação imediata em UTI/UCO com monitorização contínua; suspender TODAS as drogas não essenciais; reposição agressiva de eletrólitos (K+ alvo >4.5 mEq/L, Mg2+ >2.0 mg/dL); magnésio IV empírico mesmo se níveis normais (2g em 15 min); isoproterenol ou marca-passo temporário se bradicardia ou pausas; desfibrilador externo disponível à beira do leito; evitar estimulação simpática (dor, ansiedade); cardioversão elétrica se TdP sustentada; encaminhamento para eletrofisiologista para considerar CDI (cardioversor-desfibrilador implantável); teste genético para SQTL (Síndrome do QT Longo congênita); orientação familiar e rastreamento de parentes de primeiro grau. Em todos os casos: Revisar lista completa de medicamentos (incluindo OTC e suplementos); avaliar interações medicamentosas; educar paciente sobre sintomas de alerta (palpitações, tonturas, síncope); evitar álcool e drogas recreacionais; manter hidratação adequada; evitar exercícios extenuantes até avaliação completa.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '3a0d7e6b-c53d-47de-a50c-d7774e542835';

-- 2. VERIFICAR ATUALIZAÇÃO
SELECT
    id,
    name,
    unit,
    SUBSTRING(clinical_relevance, 1, 100) || '...' as clinical_relevance_preview,
    SUBSTRING(patient_explanation, 1, 100) || '...' as patient_explanation_preview,
    SUBSTRING(conduct, 1, 100) || '...' as conduct_preview,
    last_review
FROM score_items
WHERE id = '3a0d7e6b-c53d-47de-a50c-d7774e542835';

-- 3. ATUALIZAR SCORE LEVELS COM RELEVÂNCIA CLÍNICA ESPECÍFICA

-- Level 0: QTc ≥500 ms (Alto risco)
UPDATE score_levels
SET
    clinical_relevance = 'QTc ≥500 ms representa ALTO RISCO de Torsades de Pointes (TdP) e morte súbita cardíaca. Cada 10 ms de aumento acima de 500 eleva exponencialmente o risco de arritmias ventriculares fatais. Esta faixa exige internação imediata em UTI com monitorização contínua, suspensão de todas drogas que prolonguem QT, reposição agressiva de eletrólitos (K+ >4.5 mEq/L, Mg2+ >2.0 mg/dL), e disponibilidade de desfibrilador. Pode indicar Síndrome do QT Longo congênita (canalopatias genéticas) ou causas adquiridas graves (intoxicação medicamentosa, distúrbios eletrolíticos severos, isquemia miocárdica). Necessita avaliação por eletrofisiologista para considerar implante de CDI (cardioversor-desfibrilador implantável) e teste genético. Prognóstico depende de identificação e tratamento imediato da causa subjacente.',

    patient_explanation = 'Este valor do QTc está em uma faixa muito perigosa. Indica que seu coração está levando tempo excessivo para se "recarregar" eletricamente, colocando você em alto risco de arritmias graves que podem causar morte súbita. É uma EMERGÊNCIA MÉDICA que exige internação hospitalar imediata para monitoramento contínuo do coração, correção de eletrólitos no sangue, suspensão de medicamentos perigosos, e disponibilidade de equipamentos de ressuscitação. Pode ser causado por medicamentos, problemas nos níveis de potássio/magnésio, ou uma condição genética. Requer avaliação urgente por especialista em ritmo cardíaco (eletrofisiologista) e pode necessitar dispositivo implantável para prevenir parada cardíaca.',

    conduct = 'PROTOCOLO DE EMERGÊNCIA: 1) Internação imediata em UTI/UCO com telemetria contínua; 2) Suspender TODAS as medicações não vitais, especialmente: antiarrítmicos classe IA/III, antipsicóticos (haloperidol, quetiapina), antibióticos macrolídeos (azitromicina, eritromicina), quinolonas (levofloxacino, moxifloxacino), antifúngicos azólicos, antidepressivos tricíclicos, metadona, ondansetrona; 3) Reposição agressiva de eletrólitos: K+ IV para alvo >4.5 mEq/L, Mg2+ IV 2g em 15 min seguido de infusão contínua para alvo >2.0 mg/dL, Ca2+ se hipocalcemia; 4) Sulfato de magnésio IV empírico (2g em 15 min) mesmo se níveis séricos normais (efeito estabilizador de membrana); 5) Marca-passo temporário transvenoso ou isoproterenol IV se bradicardia (<60 bpm) ou pausas sinusais; 6) Desfibrilador externo à beira do leito; 7) Cardioversão elétrica imediata se TdP sustentada; 8) ECG 12 derivações a cada 2-4h até QTc <500 ms; 9) Ecocardiograma urgente; 10) Coleta para níveis séricos de drogas suspeitas; 11) Encaminhamento urgente para eletrofisiologista; 12) Teste genético para SQTL (painéis: KCNQ1, KCNH2, SCN5A, KCNE1, KCNE2); 13) Avaliar indicação de CDI (cardioversor-desfibrilador implantável); 14) Evitar estímulos simpáticos (controlar dor, ansiedade, náuseas); 15) Rastreamento familiar de primeiro grau com ECG; 16) Documentar detalhadamente todas medicações e suplementos em uso; 17) Evitar exercícios até QTc normalizado e causa definida.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835' AND level = 0;

-- Level 1: QTc 470-499 ms (Prolongado moderado)
UPDATE score_levels
SET
    clinical_relevance = 'QTc entre 470-499 ms em homens representa prolongamento moderado com risco significativo de arritmias ventriculares, especialmente se associado a outros fatores (hipocalemia, hipomagnesemia, bradiarritmia, uso de múltiplas drogas que prolongam QT). O risco de Torsades de Pointes aumenta progressivamente nesta faixa. Necessita suspensão imediata de drogas que prolonguem o QT, correção de distúrbios eletrolíticos, e monitoramento hospitalar se sintomático (síncope, pré-síncope, palpitações). Pode indicar Síndrome do QT Longo adquirida ou forma branda de SQTL congênita. Requer avaliação cardiológica urgente para estratificação de risco e investigação etiológica. ECG seriados e monitoramento ambulatorial prolongado (Holter) são essenciais para detectar arritmias intermitentes.',

    patient_explanation = 'Seu QTc está moderadamente elevado, indicando que o coração está levando mais tempo que o normal para se recarregar eletricamente. Isso aumenta o risco de desenvolver batimentos cardíacos irregulares perigosos, especialmente se você tiver sintomas como tonturas, desmaios ou palpitações. É importante revisar urgentemente todos os seus medicamentos (alguns podem estar causando isso), corrigir níveis de potássio e magnésio no sangue, e fazer avaliação com cardiologista. Evite esforços físicos intensos até ser avaliado. Se tiver desmaios ou palpitações fortes, procure emergência imediatamente.',

    conduct = 'CONDUTA URGENTE: 1) Suspender imediatamente todas drogas que prolonguem QT (antiarrítmicos, antipsicóticos, antibióticos macrolídeos/quinolonas, antifúngicos, antidepressivos tricíclicos, metadona, ondansetrona); 2) Avaliar necessidade de internação hospitalar se: sintomas (síncope, pré-síncope, palpitações), bradicardia (<50 bpm), distúrbios eletrolíticos graves, impossibilidade de suspender drogas essenciais; 3) Se internação: monitorização contínua (telemetria) por 24-48h, ECG a cada 4-6h até estabilização; 4) Correção agressiva de eletrólitos: K+ alvo >4.5 mEq/L (reposição VO ou IV), Mg2+ alvo >2.0 mg/dL (reposição VO ou IV); 5) Se ambulatorial: ECG de controle em 1-2 semanas, dosagens de K+, Mg2+, Ca2+, TSH, Holter 24h; 6) Ecocardiograma transtorácico para avaliar função ventricular e cardiopatia estrutural; 7) Teste ergométrico (se assintomático) para avaliar comportamento do QTc ao esforço; 8) Encaminhamento para cardiologista/eletrofisiologista em 1-2 semanas; 9) Considerar teste genético se: história familiar de morte súbita, síncope inexplicada, idade jovem sem causas adquiridas evidentes; 10) Orientar paciente sobre sintomas de alerta (palpitações, tonturas, síncope) e retorno imediato se ocorrerem; 11) Evitar: exercícios extenuantes, álcool, drogas recreacionais, natação sozinho, desidratação; 12) Revisar interações medicamentosas no site CredibleMeds (www.crediblemeds.org); 13) Evitar novos fármacos que prolonguem QT; 14) Avaliar função tireoidiana (hipotireoidismo pode prolongar QT); 15) Seguimento ambulatorial quinzenal até QTc <450 ms.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835' AND level = 1;

-- Level 2: QTc 451-469 ms (Prolongado leve)
UPDATE score_levels
SET
    clinical_relevance = 'QTc entre 451-469 ms em homens caracteriza prolongamento leve, situando-se acima do limite superior da normalidade (>450 ms) mas abaixo do limiar de alto risco. Embora o risco absoluto de Torsades de Pointes seja menor que em faixas mais elevadas, ainda representa aumento de risco cardiovascular, especialmente na presença de cofatores (distúrbios eletrolíticos, drogas que prolonguem QT, cardiopatia estrutural, idade avançada). Necessita investigação etiológica para identificar causas adquiridas reversíveis (medicamentos, hipocalemia, hipomagnesemia, hipocalcemia, hipotireoidismo) versus possível variante de Síndrome do QT Longo. Requer monitoramento periódico e otimização do perfil de risco cardiovascular.',

    patient_explanation = 'Seu QTc está levemente elevado acima do normal para homens (que é até 450 ms). Isso significa que o coração está levando um pouco mais de tempo para se recarregar eletricamente após cada batimento. Embora não seja uma emergência, indica risco aumentado de desenvolver problemas no ritmo do coração, especialmente se você usar certos medicamentos ou tiver baixos níveis de potássio/magnésio. É importante revisar seus medicamentos com o médico, fazer exames de sangue para verificar eletrólitos, e evitar novos remédios que possam piorar o problema. Também deve ficar atento a sintomas como tonturas ou palpitações e informar seu médico se ocorrerem.',

    conduct = '1) Revisão completa de medicações: Identificar e reduzir/suspender drogas que prolonguem QT quando possível (substituir por alternativas seguras). Principais classes: antiarrítmicos (amiodarona, sotalol), antipsicóticos (haloperidol, quetiapina, ziprasidona), antibióticos macrolídeos (azitromicina, eritromicina), quinolonas (levofloxacino, moxifloxacino), antifúngicos azólicos (fluconazol, voriconazol), antidepressivos tricíclicos, metadona, ondansetrona; 2) Avaliação laboratorial: K+, Mg2+, Ca2+, função renal, TSH, T4 livre, glicemia, troponina (se dor torácica); 3) Correção de eletrólitos: K+ alvo >4.0 mEq/L (suplementação VO com KCl 20-40 mEq/dia), Mg2+ alvo >2.0 mg/dL (suplementação VO com óxido de magnésio 400-800 mg/dia); 4) ECG de controle em 2-4 semanas para reavaliar após ajustes terapêuticos; 5) Considerar Holter 24h se sintomas (palpitações, tonturas) ou suspeita de arritmias intermitentes; 6) Avaliar função tireoidiana: Hipotireoidismo pode prolongar QT (tratamento com levotiroxina normaliza QTc); 7) Ecocardiograma se não realizado nos últimos 2 anos ou se suspeita de cardiopatia estrutural; 8) Encaminhamento para cardiologista se: QTc persistentemente >460 ms após correção de causas reversíveis, história familiar de morte súbita, sintomas sugestivos de arritmia; 9) Considerar teste genético se: idade <40 anos, história familiar positiva, QTc persistentemente elevado sem causas adquiridas; 10) Orientações ao paciente: Evitar automedicação, consultar médico antes de iniciar novos fármacos, manter hidratação adequada, evitar consumo excessivo de álcool, evitar exercícios extenuantes em ambientes muito quentes; 11) Evitar novos medicamentos que prolonguem QT (consultar www.crediblemeds.org); 12) Seguimento ambulatorial em 1-2 meses com ECG de controle; 13) Orientar reconhecimento de sintomas de alerta (síncope, palpitações intensas, tonturas severas) e busca de atendimento emergencial se ocorrerem.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835' AND level = 2;

-- Level 4: QTc 431-450 ms (Limítrofe)
UPDATE score_levels
SET
    clinical_relevance = 'QTc entre 431-450 ms em homens representa faixa limítrofe, situada no limite superior da normalidade segundo diretrizes europeias. Embora tecnicamente dentro da variação aceitável para alguns guidelines (que consideram <450 ms como normal), valores nesta faixa merecem atenção, especialmente se houver fatores de risco adicionais (uso de drogas que prolonguem QT, distúrbios eletrolíticos, cardiopatia estrutural, história familiar de morte súbita precoce). O risco de arritmias ventriculares é baixo, mas não desprezível. Recomenda-se vigilância quanto a fatores que possam elevar ainda mais o QTc (novos medicamentos, distúrbios metabólicos) e otimização do perfil cardiovascular. Monitoramento periódico é prudente, especialmente em contexto de polifarmácia ou comorbidades.',

    patient_explanation = 'Seu QTc está na faixa limítrofe, ou seja, no limite entre o normal e o levemente elevado. Valores até 430 ms são claramente normais, e você está um pouco acima disso. Embora não represente risco significativo neste momento, é importante ter cautela com medicamentos que possam aumentar ainda mais este valor e manter bons níveis de potássio e magnésio no sangue. Seu médico deve revisar periodicamente seus medicamentos e exames de sangue. Mantenha hábitos saudáveis e informe seu médico se tiver palpitações ou tonturas.',

    conduct = '1) Revisão periódica de medicações: Identificar drogas em uso que prolonguem QT e avaliar necessidade de continuação. Se possível, substituir por alternativas mais seguras. Principais classes: antiarrítmicos, antipsicóticos, alguns antibióticos (macrolídeos, quinolonas), antifúngicos azólicos, antidepressivos, metadona; 2) Avaliação laboratorial básica: K+, Mg2+, Ca2+, TSH (anuais ou se sintomas); 3) Manter eletrólitos em faixas ótimas: K+ >4.0 mEq/L, Mg2+ >2.0 mg/dL através de dieta rica em vegetais, frutas, oleaginosas; 4) ECG de controle anual ou antes de iniciar novos fármacos que prolonguem QT; 5) Cautela ao prescrever novos medicamentos: Consultar lista de drogas que prolongam QT (www.crediblemeds.org) antes de iniciar novos tratamentos; 6) Orientações ao paciente: Informar todos os prescritores sobre o QTc limítrofe, evitar automedicação, consultar médico antes de usar medicamentos isentos de prescrição (alguns antigripais e antialérgicos prolongam QT); 7) Manter estilo de vida saudável: Dieta balanceada rica em potássio (banana, laranja, abacate) e magnésio (oleaginosas, folhas verdes), hidratação adequada, evitar consumo excessivo de álcool, exercícios regulares moderados; 8) Não necessita restrição de atividades físicas em indivíduos assintomáticos; 9) Considerar Holter 24h se sintomas (palpitações, tonturas, síncope) ou história familiar de morte súbita precoce; 10) Encaminhamento para cardiologista apenas se: QTc progredir para >450 ms em controles futuros, sintomas sugestivos de arritmia, história familiar positiva para morte súbita ou SQTL; 11) Seguimento em atenção primária com ECG anual; 12) Orientar paciente a reconhecer sintomas de alerta (palpitações intensas, desmaios, tonturas severas) e buscar avaliação médica se ocorrerem, embora sejam improváveis nesta faixa de QTc.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835' AND level = 4;

-- Level 5: QTc <430 ms (Normal)
UPDATE score_levels
SET
    clinical_relevance = 'QTc <430 ms em homens adultos está dentro da faixa de normalidade, indicando função adequada dos canais iônicos cardíacos (potássio, sódio, cálcio) envolvidos na repolarização ventricular. Valores nesta faixa associam-se a baixo risco de arritmias ventriculares relacionadas a prolongamento do QT. Representa eletrofisiologia cardíaca normal, sem necessidade de investigação adicional ou medidas terapêuticas específicas em relação ao intervalo QT. No entanto, deve-se manter vigilância em contextos de introdução futura de medicamentos que prolonguem QT, desenvolvimento de distúrbios eletrolíticos, ou aparecimento de cardiopatia estrutural, situações que poderiam alterar o QTc subsequentemente.',

    patient_explanation = 'Seu QTc está completamente normal. O coração está se recarregando eletricamente no tempo adequado, sem risco de arritmias relacionadas ao QT. Você não precisa de nenhum tratamento específico nem preocupações quanto a este parâmetro. Mantenha hábitos de vida saudáveis e siga as orientações médicas gerais para saúde cardiovascular. Se no futuro você precisar tomar medicamentos que possam afetar o ritmo do coração, seu médico fará novo monitoramento, mas por enquanto está tudo bem.',

    conduct = '1) Nenhuma conduta específica necessária em relação ao QTc; 2) Manter seguimento cardiovascular de rotina conforme idade e fatores de risco; 3) ECG de rotina conforme indicação clínica geral (não específica para QTc); 4) Em caso de introdução futura de medicamentos que prolonguem QT (antiarrítmicos, antipsicóticos, certos antibióticos/antifúngicos), considerar ECG de controle após início do tratamento; 5) Se desenvolver sintomas sugestivos de arritmia no futuro (palpitações, síncope, pré-síncope), reavaliar com ECG e Holter independentemente do QTc basal normal; 6) Manter estilo de vida saudável: Dieta balanceada, exercícios regulares, controle de fatores de risco cardiovascular (hipertensão, diabetes, dislipidemia, tabagismo); 7) Orientar paciente que QTc normal não exclui outras arritmias (fibrilação atrial, extrassístoles, taquicardias supraventriculares) caso desenvolva sintomas; 8) Não necessita restrições de atividades físicas; 9) Não necessita suplementação profilática de eletrólitos (K+, Mg2+) além da dieta habitual; 10) Seguimento conforme protocolo de atenção primária/especializada para outras condições clínicas; 11) Orientar que QTc pode mudar ao longo do tempo (envelhecimento, medicações, doenças) e monitoramento futuro será realizado conforme necessidade clínica.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835' AND level = 5;

-- 4. VERIFICAR ATUALIZAÇÃO DOS NÍVEIS
SELECT
    level,
    name,
    lower_limit,
    upper_limit,
    operator,
    SUBSTRING(clinical_relevance, 1, 80) || '...' as clinical_preview,
    last_review
FROM score_levels
WHERE item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835'
ORDER BY level;

-- 5. CRIAR RELAÇÕES COM ARTICLES (article_score_items)

-- Artigo 1: Genetics Insights in the Relationship Between Type 2 Diabetes and Coronary Heart Disease
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
VALUES (
    '3957a6ce-dfd4-4027-b5a2-a46391cd6827',
    '3a0d7e6b-c53d-47de-a50c-d7774e542835',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Artigo 2: Heart Rate Variability in Cardiovascular Disease Diagnosis, Prognosis and Management
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
VALUES (
    'eddc9921-0f50-406b-aea4-b2b37594385c',
    '3a0d7e6b-c53d-47de-a50c-d7774e542835',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Artigo 3: ACC/AHA/HRS Guideline on the Evaluation and Management of Patients With Bradycardia and Cardiac Conduction Delay
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
VALUES (
    '5f6a3374-d88d-4f9f-9abd-97906a74919d',
    '3a0d7e6b-c53d-47de-a50c-d7774e542835',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- 6. VERIFICAR RELAÇÕES CRIADAS
SELECT
    a.title,
    a.doi,
    a.journal,
    si.name as score_item_name
FROM article_score_items asi
JOIN articles a ON a.id = asi.article_id
JOIN score_items si ON si.id = asi.score_item_id
WHERE asi.score_item_id = '3a0d7e6b-c53d-47de-a50c-d7774e542835';

-- 7. RESUMO FINAL
SELECT
    'ITEM ATUALIZADO' as status,
    COUNT(DISTINCT sl.id) as total_levels_updated,
    COUNT(DISTINCT asi.article_id) as total_articles_linked
FROM score_items si
LEFT JOIN score_levels sl ON sl.item_id = si.id
LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
WHERE si.id = '3a0d7e6b-c53d-47de-a50c-d7774e542835';

-- ============================================================
-- FIM DO SCRIPT
-- ============================================================
