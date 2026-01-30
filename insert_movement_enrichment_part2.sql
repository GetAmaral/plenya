-- SQL Parte 2: Items 4-9 (Esportes e Atividades ao Ar Livre)

-- ITEMS 4-6: Esportes praticados (frequência e intensidade)
-- ID: d33e6d45-936f-4e8a-ad02-789abdf15ae6
UPDATE score_items
SET
  clinical_relevance = 'A frequência e intensidade do treinamento esportivo são determinantes primários das adaptações metabólicas, neuromusculares e cardiovasculares. Estudos meta-analíticos demonstram que o volume cumulativo de treinamento - produto de intensidade, duração e frequência - governa a magnitude das adaptações mitocondriais. Volumes maiores de treinamento (maior frequência semanal ou maior duração da intervenção) levam a aumentos mais pronunciados tanto no conteúdo mitocondrial quanto no VO2 máx.

As adaptações temporais seguem cronologia previsível: o conteúdo mitocondrial começa a melhorar já após 2 semanas de intervenção, com aumentos contínuos observados entre 2-6 semanas e novamente de 6-10 semanas para protocolos de resistência aeróbica e alta intensidade. A biogênese mitocondrial é mediada pela ativação de PGC-1α (peroxisome proliferator-activated receptor gamma coactivator 1-alpha), AMPK (AMP-activated protein kinase) e p38 MAPK.

Tanto intervalos de alta intensidade quanto sessões contínuas de resistência aeróbica promovem melhorias comparáveis, sugerindo que a carga cumulativa de treinamento - não apenas a intensidade - determina quão robustamente as mitocôndrias se adaptam. Dados recentes de 2025 desafiam a ampla recomendação de treinamento em "Zona 2" para população geral, pois contradizem evidências substanciais que apoiam o uso de exercício de alta intensidade para melhorar capacidade mitocondrial e saúde cardiometabólica.

Para adultos jovens (19-25 anos) e atletas experientes, frequências baixas de treinamento (≤2 sessões/semana) podem produzir adaptações favoráveis, especialmente em protocolos de treinamento resistido metabólico (MRT). Frequências mais altas (3-5x/semana) são necessárias para maximizar hipertrofia muscular e força máxima.

O treinamento intervalado combinado com restrição de fluxo sanguíneo (BFR) traz benefícios adicionais para adaptações metabólicas (capacidades aeróbica e anaeróbica) e musculares (força e resistência), potencialmente otimizando desempenho de resistência mesmo com cargas reduzidas.

A intensidade adequada é contexto-dependente: atletas de endurance beneficiam-se de modelos polarizados (80% volume baixa intensidade + 20% alta intensidade), enquanto população geral e indivíduos com tempo limitado podem obter benefícios cardiometabólicos superiores com protocolos de HIIT (4-10 intervalos de 30s-4min a 85-95% FCmax).',

  patient_explanation = 'A frequência e intensidade com que você pratica esportes determinam os resultados que obterá. Não é necessário treinar todos os dias em alta intensidade para ter benefícios - na verdade, para a maioria das pessoas, 2-3 sessões por semana de exercício intenso combinado com movimento regular de baixa intensidade já produz excelentes resultados.

Seu corpo começa a se adaptar ao exercício muito rapidamente - já nas primeiras 2 semanas suas células começam a produzir mais "usinas de energia" (mitocôndrias), e esses benefícios continuam se acumulando nas primeiras 10 semanas. Depois disso, manter a prática regular preserva e continua aprimorando essas adaptações.

O importante é encontrar um equilíbrio: exercícios muito fáceis não estimulam adaptações suficientes, mas exercícios excessivamente intensos todos os dias podem levar a fadiga crônica e lesões. A chave é variar entre dias mais intensos e dias mais leves, sempre respeitando a recuperação adequada.

Se você tem pouco tempo, saiba que sessões curtas de alta intensidade (20-30 minutos) podem ser tão efetivas quanto sessões longas de intensidade moderada, especialmente para saúde cardiovascular e metabólica.',

  conduct = 'Avaliação inicial:
- Histórico esportivo detalhado: modalidades, frequência atual, duração, intensidade percebida
- Nível de condicionamento: sedentário, iniciante, intermediário, avançado, atleta
- Objetivos: saúde geral, perda de peso, performance esportiva, composição corporal
- Limitações: lesões prévias, condições médicas, tempo disponível
- Preferências: individual vs. coletivo, indoor vs. outdoor, estruturado vs. livre

Prescrição estratificada por objetivo e nível:

Saúde geral / Prevenção (sedentário a iniciante):
- Frequência: 3-5 dias/semana
- Duração: 30-60 min/sessão
- Intensidade: 60-75% FCmax (zona 2-3)
- Modalidades: caminhada rápida, ciclismo leve, natação, dança
- Progressão: aumentar 10% volume/semana

Condicionamento intermediário:
- Frequência: 4-6 dias/semana
- Modelo polarizado: 80% volume baixa-moderada intensidade + 20% alta intensidade
- Incluir 2-3 sessões de fortalecimento muscular
- Exemplo semana: 3 dias aeróbico moderado (45-60min) + 1-2 HIIT (20-30min) + 2-3 força

Performance esportiva / Atletas:
- Periodização: ciclos de base (volume alto, intensidade moderada), intensificação (volume moderado, intensidade alta), pico (volume baixo, intensidade muito alta)
- Frequência: 5-7 dias/semana com microciclos de recuperação
- Monitorar variabilidade de frequência cardíaca (HRV) para ajustar carga
- Considerar BFR training para otimizar adaptações com menor carga articular

Protocolos de alta eficiência temporal (< 30min disponível):
- HIIT clássico: 10 x 1min a 90% FCmax com 1min recuperação ativa
- Tabata: 8 x 20s máximo esforço com 10s repouso
- MRT (Metabolic Resistance Training): circuitos de exercícios resistidos com intervalos curtos

Monitoramento:
- Sinais de overtraining: fadiga persistente, insônia, irritabilidade, platô ou declínio de performance
- Marcadores objetivos: frequência cardíaca em repouso aumentada, HRV reduzida
- Ajustar volume/intensidade se 2+ sinais presentes
- Reavaliar prescrição a cada 8-12 semanas',

  last_review = CURRENT_TIMESTAMP
WHERE id = 'd33e6d45-936f-4e8a-ad02-789abdf15ae6';

-- ID: bc23dde1-1bc3-4b99-8f31-e862e74e14c6
UPDATE score_items
SET
  clinical_relevance = 'A frequência e intensidade do treinamento esportivo são determinantes primários das adaptações metabólicas, neuromusculares e cardiovasculares. Estudos meta-analíticos demonstram que o volume cumulativo de treinamento - produto de intensidade, duração e frequência - governa a magnitude das adaptações mitocondriais. Volumes maiores de treinamento (maior frequência semanal ou maior duração da intervenção) levam a aumentos mais pronunciados tanto no conteúdo mitocondrial quanto no VO2 máx.

As adaptações temporais seguem cronologia previsível: o conteúdo mitocondrial começa a melhorar já após 2 semanas de intervenção, com aumentos contínuos observados entre 2-6 semanas e novamente de 6-10 semanas para protocolos de resistência aeróbica e alta intensidade. A biogênese mitocondrial é mediada pela ativação de PGC-1α (peroxisome proliferator-activated receptor gamma coactivator 1-alpha), AMPK (AMP-activated protein kinase) e p38 MAPK.

Tanto intervalos de alta intensidade quanto sessões contínuas de resistência aeróbica promovem melhorias comparáveis, sugerindo que a carga cumulativa de treinamento - não apenas a intensidade - determina quão robustamente as mitocôndrias se adaptam. Dados recentes de 2025 desafiam a ampla recomendação de treinamento em "Zona 2" para população geral, pois contradizem evidências substanciais que apoiam o uso de exercício de alta intensidade para melhorar capacidade mitocondrial e saúde cardiometabólica.

Para adultos jovens (19-25 anos) e atletas experientes, frequências baixas de treinamento (≤2 sessões/semana) podem produzir adaptações favoráveis, especialmente em protocolos de treinamento resistido metabólico (MRT). Frequências mais altas (3-5x/semana) são necessárias para maximizar hipertrofia muscular e força máxima.

O treinamento intervalado combinado com restrição de fluxo sanguíneo (BFR) traz benefícios adicionais para adaptações metabólicas (capacidades aeróbica e anaeróbica) e musculares (força e resistência), potencialmente otimizando desempenho de resistência mesmo com cargas reduzidas.

A intensidade adequada é contexto-dependente: atletas de endurance beneficiam-se de modelos polarizados (80% volume baixa intensidade + 20% alta intensidade), enquanto população geral e indivíduos com tempo limitado podem obter benefícios cardiometabólicos superiores com protocolos de HIIT (4-10 intervalos de 30s-4min a 85-95% FCmax).',

  patient_explanation = 'A frequência e intensidade com que você pratica esportes determinam os resultados que obterá. Não é necessário treinar todos os dias em alta intensidade para ter benefícios - na verdade, para a maioria das pessoas, 2-3 sessões por semana de exercício intenso combinado com movimento regular de baixa intensidade já produz excelentes resultados.

Seu corpo começa a se adaptar ao exercício muito rapidamente - já nas primeiras 2 semanas suas células começam a produzir mais "usinas de energia" (mitocôndrias), e esses benefícios continuam se acumulando nas primeiras 10 semanas. Depois disso, manter a prática regular preserva e continua aprimorando essas adaptações.

O importante é encontrar um equilíbrio: exercícios muito fáceis não estimulam adaptações suficientes, mas exercícios excessivamente intensos todos os dias podem levar a fadiga crônica e lesões. A chave é variar entre dias mais intensos e dias mais leves, sempre respeitando a recuperação adequada.

Se você tem pouco tempo, saiba que sessões curtas de alta intensidade (20-30 minutos) podem ser tão efetivas quanto sessões longas de intensidade moderada, especialmente para saúde cardiovascular e metabólica.',

  conduct = 'Avaliação inicial:
- Histórico esportivo detalhado: modalidades, frequência atual, duração, intensidade percebida
- Nível de condicionamento: sedentário, iniciante, intermediário, avançado, atleta
- Objetivos: saúde geral, perda de peso, performance esportiva, composição corporal
- Limitações: lesões prévias, condições médicas, tempo disponível
- Preferências: individual vs. coletivo, indoor vs. outdoor, estruturado vs. livre

Prescrição estratificada por objetivo e nível:

Saúde geral / Prevenção (sedentário a iniciante):
- Frequência: 3-5 dias/semana
- Duração: 30-60 min/sessão
- Intensidade: 60-75% FCmax (zona 2-3)
- Modalidades: caminhada rápida, ciclismo leve, natação, dança
- Progressão: aumentar 10% volume/semana

Condicionamento intermediário:
- Frequência: 4-6 dias/semana
- Modelo polarizado: 80% volume baixa-moderada intensidade + 20% alta intensidade
- Incluir 2-3 sessões de fortalecimento muscular
- Exemplo semana: 3 dias aeróbico moderado (45-60min) + 1-2 HIIT (20-30min) + 2-3 força

Performance esportiva / Atletas:
- Periodização: ciclos de base (volume alto, intensidade moderada), intensificação (volume moderado, intensidade alta), pico (volume baixo, intensidade muito alta)
- Frequência: 5-7 dias/semana com microciclos de recuperação
- Monitorar variabilidade de frequência cardíaca (HRV) para ajustar carga
- Considerar BFR training para otimizar adaptações com menor carga articular

Protocolos de alta eficiência temporal (< 30min disponível):
- HIIT clássico: 10 x 1min a 90% FCmax com 1min recuperação ativa
- Tabata: 8 x 20s máximo esforço com 10s repouso
- MRT (Metabolic Resistance Training): circuitos de exercícios resistidos com intervalos curtos

Monitoramento:
- Sinais de overtraining: fadiga persistente, insônia, irritabilidade, platô ou declínio de performance
- Marcadores objetivos: frequência cardíaca em repouso aumentada, HRV reduzida
- Ajustar volume/intensidade se 2+ sinais presentes
- Reavaliar prescrição a cada 8-12 semanas',

  last_review = CURRENT_TIMESTAMP
WHERE id = 'bc23dde1-1bc3-4b99-8f31-e862e74e14c6';

-- ID: fd055f4d-13fc-4d78-9988-31336e100104
UPDATE score_items
SET
  clinical_relevance = 'A frequência e intensidade do treinamento esportivo são determinantes primários das adaptações metabólicas, neuromusculares e cardiovasculares. Estudos meta-analíticos demonstram que o volume cumulativo de treinamento - produto de intensidade, duração e frequência - governa a magnitude das adaptações mitocondriais. Volumes maiores de treinamento (maior frequência semanal ou maior duração da intervenção) levam a aumentos mais pronunciados tanto no conteúdo mitocondrial quanto no VO2 máx.

As adaptações temporais seguem cronologia previsível: o conteúdo mitocondrial começa a melhorar já após 2 semanas de intervenção, com aumentos contínuos observados entre 2-6 semanas e novamente de 6-10 semanas para protocolos de resistência aeróbica e alta intensidade. A biogênese mitocondrial é mediada pela ativação de PGC-1α (peroxisome proliferator-activated receptor gamma coactivator 1-alpha), AMPK (AMP-activated protein kinase) e p38 MAPK.

Tanto intervalos de alta intensidade quanto sessões contínuas de resistência aeróbica promovem melhorias comparáveis, sugerindo que a carga cumulativa de treinamento - não apenas a intensidade - determina quão robustamente as mitocôndrias se adaptam. Dados recentes de 2025 desafiam a ampla recomendação de treinamento em "Zona 2" para população geral, pois contradizem evidências substanciais que apoiam o uso de exercício de alta intensidade para melhorar capacidade mitocondrial e saúde cardiometabólica.

Para adultos jovens (19-25 anos) e atletas experientes, frequências baixas de treinamento (≤2 sessões/semana) podem produzir adaptações favoráveis, especialmente em protocolos de treinamento resistido metabólico (MRT). Frequências mais altas (3-5x/semana) são necessárias para maximizar hipertrofia muscular e força máxima.

O treinamento intervalado combinado com restrição de fluxo sanguíneo (BFR) traz benefícios adicionais para adaptações metabólicas (capacidades aeróbica e anaeróbica) e musculares (força e resistência), potencialmente otimizando desempenho de resistência mesmo com cargas reduzidas.

A intensidade adequada é contexto-dependente: atletas de endurance beneficiam-se de modelos polarizados (80% volume baixa intensidade + 20% alta intensidade), enquanto população geral e indivíduos com tempo limitado podem obter benefícios cardiometabólicos superiores com protocolos de HIIT (4-10 intervalos de 30s-4min a 85-95% FCmax).',

  patient_explanation = 'A frequência e intensidade com que você pratica esportes determinam os resultados que obterá. Não é necessário treinar todos os dias em alta intensidade para ter benefícios - na verdade, para a maioria das pessoas, 2-3 sessões por semana de exercício intenso combinado com movimento regular de baixa intensidade já produz excelentes resultados.

Seu corpo começa a se adaptar ao exercício muito rapidamente - já nas primeiras 2 semanas suas células começam a produzir mais "usinas de energia" (mitocôndrias), e esses benefícios continuam se acumulando nas primeiras 10 semanas. Depois disso, manter a prática regular preserva e continua aprimorando essas adaptações.

O importante é encontrar um equilíbrio: exercícios muito fáceis não estimulam adaptações suficientes, mas exercícios excessivamente intensos todos os dias podem levar a fadiga crônica e lesões. A chave é variar entre dias mais intensos e dias mais leves, sempre respeitando a recuperação adequada.

Se você tem pouco tempo, saiba que sessões curtas de alta intensidade (20-30 minutos) podem ser tão efetivas quanto sessões longas de intensidade moderada, especialmente para saúde cardiovascular e metabólica.',

  conduct = 'Avaliação inicial:
- Histórico esportivo detalhado: modalidades, frequência atual, duração, intensidade percebida
- Nível de condicionamento: sedentário, iniciante, intermediário, avançado, atleta
- Objetivos: saúde geral, perda de peso, performance esportiva, composição corporal
- Limitações: lesões prévias, condições médicas, tempo disponível
- Preferências: individual vs. coletivo, indoor vs. outdoor, estruturado vs. livre

Prescrição estratificada por objetivo e nível:

Saúde geral / Prevenção (sedentário a iniciante):
- Frequência: 3-5 dias/semana
- Duração: 30-60 min/sessão
- Intensidade: 60-75% FCmax (zona 2-3)
- Modalidades: caminhada rápida, ciclismo leve, natação, dança
- Progressão: aumentar 10% volume/semana

Condicionamento intermediário:
- Frequência: 4-6 dias/semana
- Modelo polarizado: 80% volume baixa-moderada intensidade + 20% alta intensidade
- Incluir 2-3 sessões de fortalecimento muscular
- Exemplo semana: 3 dias aeróbico moderado (45-60min) + 1-2 HIIT (20-30min) + 2-3 força

Performance esportiva / Atletas:
- Periodização: ciclos de base (volume alto, intensidade moderada), intensificação (volume moderado, intensidade alta), pico (volume baixo, intensidade muito alta)
- Frequência: 5-7 dias/semana com microciclos de recuperação
- Monitorar variabilidade de frequência cardíaca (HRV) para ajustar carga
- Considerar BFR training para otimizar adaptações com menor carga articular

Protocolos de alta eficiência temporal (< 30min disponível):
- HIIT clássico: 10 x 1min a 90% FCmax com 1min recuperação ativa
- Tabata: 8 x 20s máximo esforço com 10s repouso
- MRT (Metabolic Resistance Training): circuitos de exercícios resistidos com intervalos curtos

Monitoramento:
- Sinais de overtraining: fadiga persistente, insônia, irritabilidade, platô ou declínio de performance
- Marcadores objetivos: frequência cardíaca em repouso aumentada, HRV reduzida
- Ajustar volume/intensidade se 2+ sinais presentes
- Reavaliar prescrição a cada 8-12 semanas',

  last_review = CURRENT_TIMESTAMP
WHERE id = 'fd055f4d-13fc-4d78-9988-31336e100104';

-- Continua com Atividades ao Ar Livre...
