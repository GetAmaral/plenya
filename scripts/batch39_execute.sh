#!/bin/bash

# BATCH 39: Composição Corporal Parte 2 - 100 items
# Execução via psql com blocos IN()

DBCMD="docker compose exec -T db psql -U plenya_user -d plenya_db -c"

echo "🚀 Iniciando Batch 39 - Composição Corporal Parte 2..."

# 1. Nascimento (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Peso ao nascer e idade gestacional são preditores de composição corporal futura. Baixo peso (< 2500g) associa-se a sarcopenia e diabetes tipo 2. Alto peso (> 4000g) predispõe a obesidade e síndrome metabólica.',
  patient_explanation = 'O peso ao nascer e se você nasceu prematuro influenciam sua composição corporal. Bebês com baixo peso têm maior risco de menos massa muscular.',
  conduct = 'Documentar peso ao nascer, idade gestacional. Baixo peso: rastreamento de sarcopenia. Alto peso: monitoramento de obesidade.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('74dfa65f-df48-4331-9a47-72fa2d52b5e2', '4ba81974-b66a-42c5-92cf-9093d3084e38', '5e02310c-08f1-4f88-b22d-4d89447fe565');"

# 2. Infância (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Trajetória de peso/altura na infância prediz composição corporal adulta. Ganho rápido nos primeiros 2 anos associa-se a adiposidade visceral. Obesidade infantil predispõe a obesidade adulta e sarcopenia por inflamação crônica.',
  patient_explanation = 'Como você cresceu e ganhou peso na infância afeta sua composição corporal hoje. Obesidade infantil aumenta risco de problemas metabólicos.',
  conduct = 'Revisar curvas de crescimento (OMS), IMC percentilado, histórico de obesidade/desnutrição. Obesidade infantil: monitoramento de síndrome metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('1bdd6114-82c5-4b57-8352-ea32449169ba', '64614d75-3b62-48ac-a6e8-294fa906ae2e', '84baa5e6-db01-4a90-af2b-801ce9f0670b');"

# 3. Adolescência (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Adolescência é período crítico para desenvolvimento muscular e ósseo (pico aos 18-25 anos). Ganho excessivo na puberdade associa-se a obesidade adulta. Anorexia/bulimia comprometem saúde óssea irreversivelmente.',
  patient_explanation = 'A adolescência é quando você constrói maior parte da massa muscular e óssea. Problemas de peso nessa fase afetam saúde metabólica para sempre.',
  conduct = 'Documentar estirão puberal, IMC adolescente, transtornos alimentares. Obesidade: rastreamento de sarcopenia prematura. Anorexia: DEXA ósseo.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('8a932b13-17c6-42d1-b428-dfe296accd4b', '657dcaae-f159-4415-8531-5a2908c0d095');"

# 4. Vida adulta (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Histórico de composição corporal (DEXA, bioimpedância) identifica trajetórias de ganho/perda muscular. Perda > 5-10% em 6-12 meses indica sarcopenia secundária. Variações cíclicas associam-se a perda muscular e ganho visceral.',
  patient_explanation = 'Exames prévios de composição corporal mostram como seu corpo mudou. Perdas musculares ou ganhos de gordura visceral indicam riscos metabólicos.',
  conduct = 'Revisar avaliações prévias (DEXA, bioimpedância). Calcular variações de ASMI, FMI, gordura visceral. Perda > 5-10%: investigar sarcopenia secundária.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('ac04ab06-bdc7-4959-a33d-32f3646c1a87', '3c9b2568-a4ea-4132-9c3e-f327ad1b0ce7', '852be041-945c-4ced-91a8-4f85b55e9651');"

# 5. Picos de peso (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Picos de peso identificam períodos de risco metabólico e gatilhos ambientais. Ganho rápido (> 10% em 6 meses) associa-se a gordura visceral, resistência insulínica e inflamação. Contextos: menopausa, corticoides, antipsicóticos.',
  patient_explanation = 'Momentos de ganho rápido de peso podem indicar fatores desencadeantes (estresse, medicamentos, hormônios). Compreender esses padrões ajuda a prevenir novos ganhos.',
  conduct = 'Identificar gatilhos de ganho rápido. Ganho > 10% em 6 meses: HOMA-IR, gordura visceral (DEXA/TC). Ajustar medicações se necessário.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('3b9f7a1e-527e-4a4c-bc23-87c68c41ee89', 'ef3181a6-135a-4110-9275-68eae1c3890a', '1af195d6-77c0-495b-9833-9ad38362e6fb');"

# 6. Melhor composição (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Períodos de melhor composição corporal revelam condições ideais (atividade física, sono, alimentação, estresse). Identificar contextos favoráveis (musculação, low-carb, jejum) permite replicá-los. Melhora associa-se a redução de PCR, IL-6.',
  patient_explanation = 'Momentos de melhor forma física indicam o que funciona para seu corpo. Identificar estratégias que deram certo ajuda a criar plano personalizado e sustentável.',
  conduct = 'Documentar condições da melhor composição: tipo de exercício, padrão alimentar, rotina de sono, manejo de estresse. Replicar estratégias bem-sucedidas.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('7b4975a7-a24d-4bbc-8d27-2dd89f9a99ad', '8d587ace-8023-41ab-85b2-334a2a00f41b', 'e05d895d-36fc-48ff-8201-0098148ca9ac');"

# 7. Tratamentos prévios (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Histórico de tratamentos identifica estratégias eficazes e efeitos adversos. Dietas restritivas repetidas associam-se a perda muscular e ganho visceral. Medicamentos (GLP-1, orlistat) têm eficácia variável. Cirurgia bariátrica requer suplementação vitalícia.',
  patient_explanation = 'Conhecer tratamentos que você tentou ajuda a evitar repetir estratégias que não funcionaram e identificar o que foi eficaz. Isso personaliza seu plano.',
  conduct = 'Documentar tratamentos prévios: dietas, medicamentos (orlistat, liraglutida, semaglutida), hormônios, cirurgias. Avaliar eficácia, efeitos adversos, motivos de descontinuação.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('c1c9eb73-b2ac-456f-b854-ed707c817a3d', 'e965a941-475a-4c19-9407-22b964af9ca3', '70dae5d1-7c26-42cd-a746-36b21a49e9cd');"

# 8. Histórico familiar (3x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Histórico familiar identifica predisposição genética a obesidade, sarcopenia e distribuição de gordura. Obesidade em parentes de 1º grau aumenta risco 2-8x. Variantes genéticas (FTO, MC4R) afetam apetite e adipogênese. Sarcopenia familiar sugere miopatias hereditárias.',
  patient_explanation = 'A composição corporal de pais, irmãos e avós indica sua tendência genética a ganhar peso, perder músculos ou acumular gordura em certas regiões. Compreender sua genética ajuda a personalizar estratégias.',
  conduct = 'Documentar composição corporal de familiares de 1º e 2º grau. Obesidade familiar: rastreamento intensivo de síndrome metabólica. Sarcopenia familiar: investigar miopatias, vitamina D.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('03438548-88b3-44a7-8449-5441c26ea829', '85a74aaf-8c04-4552-9c8f-636cfe1c9133', '5d7f01cb-f4bb-404a-8d89-d80066384f5a');"

# 9. Avaliação atual (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Avaliação atual (DEXA, bioimpedância, pregas, fotos) fornece dados objetivos para estratificação de risco metabólico. DEXA é padrão-ouro para ASMI (sarcopenia), gordura visceral e densidade óssea. Comparação com EWGSOP2/AWGS 2019 orienta diagnóstico.',
  patient_explanation = 'Exames de composição corporal medem sua massa muscular, gordura e água. Esses dados mostram riscos metabólicos e permitem criar plano personalizado para melhorar sua saúde.',
  conduct = 'Solicitar DEXA (padrão-ouro) ou bioimpedância (acessível). Avaliar: ASMI (sarcopenia), FMI, gordura visceral, ângulo de fase, DMO. Comparar com referências (EWGSOP2, AWGS 2019).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('088b4d4d-1873-45bb-8a3a-19e4463de7a5', '0b58623e-63d3-4685-b864-c41a0058ed56');"

# 10. Tratamentos atuais (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Tratamentos em uso devem ser avaliados quanto a eficácia, segurança e adesão. GLP-1 agonistas, suplementos (proteínas, creatina, vitamina D), hormônios e dietas têm impactos diferentes em massa muscular e gordura visceral. Monitorar efeitos adversos e interações.',
  patient_explanation = 'Se você está fazendo tratamentos para composição corporal (remédios, suplementos, dietas), é importante monitorar se estão funcionando e se há efeitos colaterais. Ajustes personalizados aumentam eficácia.',
  conduct = 'Documentar tratamentos atuais: medicamentos (GLP-1, orlistat), suplementos (whey, creatina, vitamina D), hormônios, dietas (cetogênica, jejum). Monitorar eficácia, efeitos adversos, adesão.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('2579198d-65b8-4cff-9970-7c437076adc7', 'a6222101-4fee-41a4-aca2-fd466d1c789e');"

# 11. Satisfação e objetivos (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Grau de satisfação e objetivos devem ser realistas, considerando genética, idade, comorbidades e adesão. Insatisfação excessiva pode indicar dismorfia ou transtornos alimentares. Objetivos irrealistas (IMC < 18, %GC < 5%) são contraproducentes. Foco: saúde metabólica.',
  patient_explanation = 'Seus objetivos de composição corporal devem ser realistas e focados em saúde, não apenas estética. Juntos, definimos metas alcançáveis que consideram sua genética e estilo de vida.',
  conduct = 'Avaliar grau de satisfação corporal (escala 1-10), objetivos específicos. Insatisfação severa: rastreamento de dismorfia, transtornos alimentares. Educar sobre expectativas realistas e saúde metabólica.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('f836753e-9e45-48f5-a54c-298195ff0588', 'a3707b3b-d45e-4235-93b4-1b011deb5028');"

# 12. Medidas objetivas (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Medidas objetivas (antropometria, bioimpedância, DEXA) são fundamentais para diagnóstico de obesidade, sarcopenia e risco metabólico. Circunferência abdominal e razão cintura/quadril predizem risco cardiovascular melhor que IMC. ASMI < 7,0 kg/m² (H) ou < 5,5 kg/m² (M) indica sarcopenia.',
  patient_explanation = 'Medidas objetivas (peso, circunferências, exames) fornecem dados precisos sobre sua saúde metabólica e nutricional. Esses números guiam decisões e permitem acompanhar progresso cientificamente.',
  conduct = 'Realizar: peso, altura, IMC, circunferências (abdominal, quadril, pescoço, panturrilha, braço, coxa). Idealmente: bioimpedância ou DEXA. Calcular: razões cintura/quadril, cintura/altura.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('662902a2-d3f7-4d08-b3ee-4fde9e633cf8', '7c77b924-d8fb-478d-81fc-10ba491d51a0');"

# 13. Peso (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Peso corporal é medida básica para cálculo de IMC, monitoramento de ganho/perda e ajuste de medicamentos. Variações > 5% em 1 mês exigem investigação (retenção hídrica, doença). Peso isolado não diferencia músculo de gordura. Perda deve ser lenta (0,5-1 kg/semana) para preservar músculo.',
  patient_explanation = 'Seu peso é importante, mas não conta a história completa. Variações rápidas podem indicar retenção de líquidos ou mudanças metabólicas. Acompanhar regularmente ajuda a detectar tendências.',
  conduct = 'Monitorar peso semanalmente (mesmas condições: jejum, após micção). Variação > 5% em 1 mês: investigar causas (retenção hídrica, doença, mudança composição). Combinar com bioimpedância.',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('400295c1-25e5-44c5-8d99-0f355f1aa7cc', 'c2718674-8ee8-4c54-b63c-86784cbd903c');"

# 14. Altura (1x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'Altura é fundamental para cálculo de IMC, ASMI (massa muscular apendicular/altura²) e razão cintura/altura. Redução de altura em idosos indica compressões vertebrais por osteoporose. Medir anualmente em > 50 anos.',
  patient_explanation = 'Sua altura é usada para calcular vários índices de saúde, como IMC e ASMI. Em idosos, redução de altura pode indicar osteoporose.',
  conduct = 'Medir altura com estadiômetro, sem calçados, calcanhares juntos, olhar no horizonte. Em > 50 anos: medir anualmente. Redução > 2 cm: investigar osteoporose (DEXA, vitamina D, cálcio).',
  last_review = CURRENT_TIMESTAMP
WHERE id = '48b082bf-3697-4c8a-a183-b2fc4396d270';"

# 15. IMC (2x)
$DBCMD "UPDATE score_items SET
  clinical_relevance = 'IMC (kg/m²) é triagem inicial: < 18,5 baixo peso, 18,5-24,9 normal, 25-29,9 sobrepeso, 30-34,9 obesidade I, 35-39,9 obesidade II, ≥ 40 obesidade III. Limitações: não diferencia músculo de gordura, não avalia distribuição. Circunferência abdominal e bioimpedância são complementares.',
  patient_explanation = 'O IMC relaciona seu peso e altura. Embora útil como triagem inicial, não diferencia músculos de gordura. Por isso, usamos outras medidas complementares para avaliar melhor sua composição corporal.',
  conduct = 'Calcular IMC = peso (kg) / altura² (m²). < 18,5: investigar desnutrição, sarcopenia. 25-29,9: rastreamento de síndrome metabólica. ≥ 30: avaliação completa de composição corporal (DEXA, bioimpedância).',
  last_review = CURRENT_TIMESTAMP
WHERE id IN ('bf875f17-1a90-4bb8-8305-e0e0285aeb80', '7efe0605-a423-4560-91aa-05c285a93c4f');"

echo "✅ Batch 39 Parte 1 (até IMC): 44 items processados!"
echo "Continuando com medidas antropométricas..."
