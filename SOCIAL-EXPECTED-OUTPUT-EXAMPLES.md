# Exemplos de Output Esperado - Batch SOCIAL

Exemplos de conteúdo clínico esperado para diferentes categorias de items SOCIAL.

---

## Exemplo 1: AMBIENTE SONORO

### Item
```json
{
  "id": "c84412f7-393f-41d0-8bd7-0a28824dbeb0",
  "name": "Nível de ruído no ambiente de trabalho/casa",
  "description": "Avalia exposição a poluição sonora",
  "input_type": "select",
  "options": ["Silencioso (<40dB)", "Moderado (40-70dB)", "Alto (>70dB)", "Muito Alto (>85dB)"]
}
```

### Output Esperado

**Clinical Relevance** (2-3 parágrafos):
```
A exposição crônica a ruído ambiental representa um estressor fisiológico significativo
com impactos mensuráveis na saúde cardiometabólica. A ativação repetida do eixo
hipotálamo-pituitária-adrenal (HPA) em resposta a ruído >70dB resulta em hipercortisolemia
crônica, aumento de catecolaminas circulantes e ativação do sistema nervoso simpático.
Este estado de "alerta" constante promove resistência insulínica, disfunção endotelial
e inflamação sistêmica com elevação de marcadores como IL-6, TNF-α e proteína C reativa.

Estudos epidemiológicos demonstram que cada 10dB de aumento na exposição a ruído de
tráfego correlaciona com aumento de 8% no risco de hipertensão arterial sistêmica e
14% no risco de diabetes tipo 2. O mecanismo envolve vasoconstrição mediada por
catecolaminas, redução da variabilidade da frequência cardíaca (indicando domínio
simpático) e estresse oxidativo que compromete a função endotelial. A exposição noturna
é particularmente deletéria, pois fragmenta o sono profundo (estágios 3-4 NREM),
impedindo a recuperação parassimpática necessária para regeneração celular.

Na medicina funcional, avaliar exposição sonora é essencial para pacientes com síndrome
metabólica, hipertensão resistente, insônia crônica e fadiga adrenal. O ruído atua como
disruptor endócrino silencioso, elevando cortisol vespertino e suprimindo DHEA, criando
um perfil hormonal pró-inflamatório e catabólico. Profissionais expostos a ruído
ocupacional >85dB sem proteção adequada apresentam risco 40% maior de doença
cardiovascular prematura.
```

**Interpretation Guidelines** (Guia prático):
```
PADRÃO 1: Exposição a ruído >70dB durante >8h/dia (ex: tráfego intenso, construção)
- Significado clínico: Ativação crônica do eixo HPA, dominância simpática
- Sistemas comprometidos: Cardiovascular (HTA, arritmias), endócrino (hipercortisolemia),
  neurológico (declínio cognitivo), metabólico (resistência insulínica)
- Investigar: Cortisol salivar 4 pontos (incluir ponto noturno), variabilidade FC (Holter),
  marcadores inflamatórios (PCR-us, IL-6), MAPA 24h
- Diagnósticos diferenciais: HTA secundária, síndrome fadiga adrenal, síndrome metabólica

PADRÃO 2: Ruído noturno >45dB (ex: tráfego, vizinhos, parceiro roncando)
- Significado clínico: Fragmentação do sono, impossibilidade de atingir sono profundo
- Sistemas comprometidos: Sono (arquitetura anormal), cognitivo (consolidação memória),
  imune (redução NK cells), metabólico (resistência leptina)
- Investigar: Polissonografia, cortisol noturno (deveria ser <1µg/dL às 23h),
  questionário qualidade sono (Pittsburgh), biomarcadores recuperação (DHEA/cortisol ratio)
- Diagnósticos diferenciais: Insônia primária, apneia do sono, síndrome pernas inquietas

PADRÃO 3: Trabalho em ambiente com ruído >85dB sem proteção (ex: fábrica, aeroporto)
- Significado clínico: Exposição ocupacional de alto risco, perda auditiva + efeitos
  sistêmicos
- Sistemas comprometidos: Auditivo (PAIR - Perda Auditiva Induzida por Ruído),
  cardiovascular, psicológico (irritabilidade, depressão)
- Investigar: Audiometria tonal, MAPA 24h, cortisol e catecolaminas urinárias 24h,
  avaliação psicológica (burnout, ansiedade)
- Ação imediata: Exigir EPIs (protetores auriculares tipo plug + concha = -30dB),
  pausas em ambiente silencioso, considerar mudança função

PADRÃO 4: Sensibilidade aumentada a ruído (misofonia)
- Significado clínico: Possível sensibilização central, hiperatividade amígdala
- Sistemas comprometidos: Neurológico (processamento sensorial), psicológico
- Investigar: Histórico de trauma, ansiedade generalizada, sensibilidade a outros
  estímulos (luz, cheiros), déficits nutricionais (magnésio, B12, ômega-3)
```

**Actionable Insights** (5-8 intervenções):
```
1. SE exposição ocupacional >85dB:
   - IMEDIATO: Protetores auriculares (plug + concha = atenuação -30 a -35dB)
   - MÉDIO PRAZO: Negociar com empregador pausas em ambiente silencioso (15min a cada 2h)
   - LONGO PRAZO: Considerar mudança de função ou setor menos ruidoso
   - Monitorar: Audiometria anual, marcadores cardiovasculares

2. SE ruído residencial >70dB (tráfego, construção):
   - AMBIENTAL: Janelas anti-ruído (vidro duplo/triplo = -30dB), cortinas acústicas
     (-10dB), vedação portas/janelas
   - QUARTO: Priorizar isolamento acústico no quarto (sono reparador é crítico)
   - ALTERNATIVA: Ruído branco ou rosa (máquina de som) para mascarar ruídos irregulares
   - INVESTIMENTO: Isolamento acústico profissional (custo-benefício excelente para saúde)

3. SE ruído noturno >45dB:
   - IMEDIATO: Protetores auriculares macios para dormir (reduzem -20 a -30dB)
   - ALTERNATIVA: Fones com cancelamento ativo de ruído + ruído branco/rosa
   - PARCEIRO: Se ronco, avaliar apneia do sono (polissonografia), perda de peso,
     CPAP se indicado
   - SUPLEMENTAÇÃO: Magnésio glicinato 400mg (relaxamento muscular, melhora sono profundo),
     L-teanina 200mg (reduz ativação simpática)

4. SE cortisol elevado por ruído crônico:
   - NUTRACÊUTICOS: Ashwagandha KSM-66 600mg/dia (reduz cortisol 27% em 60 dias),
     Fosfatidilserina 300mg noite (reduz cortisol pós-estresse)
   - PRÁTICAS: Meditação 10-20min/dia (reduz reatividade amígdala), respiração
     vagal (4-7-8) antes dormir
   - MONITORAR: Cortisol salivar 4 pontos a cada 3 meses até normalização

5. SE impossibilidade de mudança ambiental (limitações financeiras/contratuais):
   - OTIMIZAR RECUPERAÇÃO: Criar "santuário silencioso" em casa (1 cômodo isolado
     acusticamente)
   - PAUSAS ESTRATÉGICAS: 5min silêncio absoluto a cada hora (reduz carga alostática)
   - FINS DE SEMANA: Buscar ambientes naturais silenciosos (floresta, praia) - 2h
     mínimo para recuperação do sistema nervoso
   - SUPORTE SISTÊMICO: Antioxidantes (NAC 600mg, vitamina C 1g) para combater
     estresse oxidativo induzido por ruído

6. SE sintomas cardiovasculares (HTA, palpitações):
   - AVALIAÇÃO: MAPA 24h (hipertensão pode ser pior em horários de pico de ruído),
     Holter (detectar arritmias induzidas por estresse)
   - SUPLEMENTAÇÃO: Magnésio 400-600mg, CoQ10 100-200mg, ômega-3 2g/dia
   - FARMACOLÓGICO: Considerar beta-bloqueador se HTA + taquicardia (reduz resposta
     adrenérgica ao ruído)

7. SE declínio cognitivo/memória em expostos a ruído crônico:
   - NEUROPROTEÇÃO: Ômega-3 DHA 1g/dia, bacopa monnieri 300mg, lions mane 500mg
   - AVALIAÇÃO: Testes cognitivos (MoCA, MMSE), RNM crânio se >55 anos + exposição
     prolongada
   - INTERVENÇÃO: Treinamento cognitivo, quebra-cabeças, aprendizado novo idioma
     (neuroplasticidade compensatória)

8. Reavaliação: 3 meses
   - Métricas: Decibéis medidos (app de smartphone), qualidade sono (Pittsburgh),
     cortisol salivar, pressão arterial, questionário estresse percebido
```

**Red Flags** (3-5 sinais de alerta):
```
🚩 RED FLAG 1: Exposição ocupacional >85dB sem EPI + perda auditiva progressiva
   - Risco: PAIR (Perda Auditiva Induzida por Ruído) irreversível, zumbido crônico
   - Ação: IMEDIATA - Audiometria urgente, afastamento temporário, exigir EPIs,
     considerar mudança função. PAIR é 100% previsível e 100% prevenível.

🚩 RED FLAG 2: Ruído crônico + HTA resistente (>3 anti-hipertensivos sem controle)
   - Risco: Evento cardiovascular agudo (IAM, AVE), emergência hipertensiva
   - Ação: MAPA 24h urgente, investigar causas secundárias HTA, avaliar mudança
     ambiental imediata, considerar internação para controle pressórico se PA >180/110mmHg

🚩 RED FLAG 3: Ruído noturno + sonolência diurna excessiva + acidentes/quase-acidentes
   - Risco: Acidente de trânsito/trabalho, morte súbita (privação sono crônica)
   - Ação: Afastamento atividades de risco (dirigir, operar máquinas), solução
     ambiental URGENTE (mudança residencial se necessário), polissonografia

🚩 RED FLAG 4: Exposição a ruído + sintomas psicóticos (paranoia, alucinações auditivas)
   - Risco: Psicose induzida por estresse, decomposação psiquiátrica
   - Ação: Avaliação psiquiátrica URGENTE, considerar internação, remover estressor
     imediatamente, descartar uso substâncias

🚩 RED FLAG 5: Trabalhador com ruído ocupacional + IAM/AVE prematuro (<50 anos)
   - Risco: Nexo causal ocupacional, recorrência eventos cardiovasculares
   - Ação: Notificação SINAN (vigilância epidemiológica), perícia INSS para
     reconhecimento doença ocupacional, mudança função permanente, reabilitação
     cardiovascular agressiva
```

---

## Exemplo 2: CONDIÇÕES DE MORADIA (Mofo)

### Item
```json
{
  "id": "91e450db-29df-4a78-8741-441f89630ff7",
  "name": "Presença de mofo ou umidade visível na residência",
  "input_type": "select",
  "options": ["Não", "Sim, em 1 cômodo", "Sim, em múltiplos cômodos", "Sim, com odor característico"]
}
```

### Output Esperado (Resumido)

**Clinical Relevance**:
```
A exposição a fungos de interior (Aspergillus, Penicillium, Stachybotrys) e suas
micotoxinas representa um gatilho significativo para doença inflamatória sistêmica
e disfunção imune. Micotoxinas como aflatoxinas, ocratoxina A e tricotecenos são
moléculas pequenas (<1000 Da) que atravessam barreiras biológicas, acumulando-se
em tecidos adiposos e cérebro. [...]

A Síndrome da Resposta Inflamatória Sistêmica (SIRS) relacionada a biotoxinas
de mofo é caracterizada por ativação imune desregulada com elevação de TGF-β1,
C4a, MMP-9 e redução de MSH (hormônio estimulador de melanócitos). Clinicamente,
manifesta-se como fadiga crônica refratária, "brain fog", dores musculares migrantes,
sintomas respiratórios (tosse, dispneia) e sensibilidade química múltipla. [...]
```

**Interpretation Guidelines**:
```
PADRÃO 1: Mofo visível em 1 cômodo + sintomas respiratórios
- Investigar: IgG/IgE para fungos comuns (painel: Aspergillus, Penicillium,
  Cladosporium, Alternaria), eosinofilia, IgE total
- DD: Asma alérgica, rinosinusite fúngica, aspergilose broncopulmonar alérgica

PADRÃO 2: Mofo em múltiplos cômodos + fadiga crônica + brain fog
- Investigar: Micotoxinas urinárias (ocratoxina A, aflatoxinas, tricotecenos),
  TGF-β1, C4a, MSH, MMP-9, painel SIRS
- DD: Síndrome fadiga crônica, SIRS por biotoxinas, sensibilidade química múltipla

[...]
```

**Actionable Insights**:
```
1. SE mofo visível em qualquer extensão:
   - REMEDIAÇÃO AMBIENTAL (prioridade #1): Correção fonte umidade (infiltração,
     vazamento), remoção material contaminado, não apenas limpar superficialmente
   - PROFISSIONAL: Contratar remediador certificado, não DIY se >10m² afetados
   - DURANTE REMEDIAÇÃO: Paciente deve sair da residência (micotoxinas aerossolizam)

2. SE sintomas SIRS (fadiga + brain fog + dor):
   - BINDERS: Colestiramina 4g 2x/dia OU carvão ativado 500mg 3x/dia (liga micotoxinas
     intestinais, impede recirculação êntero-hepática)
   - SUPORTE HEPÁTICO: NAC 600mg 2x/dia, glutationa lipossomal 500mg/dia,
     silimarina 200mg 3x/dia
   - ANTIINFLAMATÓRIOS NATURAIS: Curcumina (Meriva®) 1g 2x/dia, ômega-3 2g/dia

[...]

6. SE impossibilidade de mudança residencial (restrições financeiras):
   - ISOLAMENTO: Criar "quarto limpo" com filtro HEPA (mínimo H13, >99.97% partículas
     >0.3µm), vedação portas
   - DESUMIDIFICADOR: Manter umidade <50% (ideal 35-45%), medir com higrômetro
   - PURIFICADOR AR: HEPA + carvão ativado em cômodos principais (troca 5-6 vezes/hora)
   - MONITORAR: Sintomas, micotoxinas urinárias a cada 3 meses até negativação
```

**Red Flags**:
```
🚩 RED FLAG 1: Mofo + sintomas respiratórios graves (dispneia, hemoptise, febre)
   - Risco: Aspergilose invasiva (se imunossuprimido), pneumonite hipersensibilidade
   - Ação: RX tórax URGENTE, broncoscopia se imunossuprimido, remoção imediata
     da residência

🚩 RED FLAG 2: Mofo + múltiplos moradores sintomáticos (cluster)
   - Risco: Exposição tóxica significativa, possível Stachybotrys (mofo negro)
   - Ação: Inspeção profissional URGENTE, teste ar (contagem esporos), mudança
     temporária toda família, notificar vigilância sanitária

[...]
```

---

## Exemplo 3: HOBBIES E LAZER

### Item
```json
{
  "id": "ac2c2df9-e65f-4156-8c39-d296c613f85c",
  "name": "Horas por semana dedicadas a hobbies ou atividades de lazer",
  "input_type": "number",
  "unit": "horas/semana"
}
```

### Output Esperado (Resumido)

**Clinical Relevance**:
```
Hobbies e atividades de lazer representam um determinante social de saúde frequentemente
negligenciado, porém com impacto mensurável em marcadores de longevidade e qualidade
de vida. Estudos prospectivos das Blue Zones (regiões com maior concentração de
centenários - Okinawa, Sardenha, Loma Linda, Ikaria) identificam "ikigai" (propósito/
razão de viver) e engajamento regular em atividades prazerosas como fatores comuns
entre os mais longevos. [...]

O mecanismo biológico envolve redução do cortisol basal e da reatividade do eixo HPA
ao estresse. Pressman et al. (2009) demonstraram que maior frequência de atividades
prazerosas correlaciona com níveis reduzidos de cortisol salivar, IL-6 e circunferência
abdominal. Hobbies criativos (pintura, música, jardinagem) estimulam neuroplasticidade
com aumento de BDNF (fator neurotrófico cerebral), enquanto hobbies sociais (grupos
comunitários) combatem isolamento - um fator de risco equivalente a fumar 15 cigarros/
dia (Holt-Lunstad, 2010). [...]
```

**Actionable Insights**:
```
1. SE <2h lazer/semana (privação recreacional):
   - PRESCREVER hobbies como tratamento médico ("receita verde")
   - SUGESTÕES: Atividades criativas não-digitais (pintura, música, jardinagem,
     artesanato), meta: 30min/dia, 3-4x/semana
   - BENEFÍCIO: Redução cortisol 25%, melhora qualidade sono, redução sintomas
     depressivos/ansiosos
   - MONITORAR: Adesão, escala estresse percebido, qualidade sono (Pittsburgh) a cada mês

2. SE isolamento social (sem hobbies em grupo):
   - RISCO: Mortalidade 23% maior, equivalente fumar 15 cigarros/dia
   - ENCAMINHAR: Grupos comunitários (caminhada, dança, coral, voluntariado),
     centros convivência idosos (se >60 anos)
   - RASTREIO: Depressão (PHQ-9), solidão (UCLA Loneliness Scale)
   - FOLLOW-UP: Reavaliação sintomas depressivos em 6 semanas

[...]

5. SE hobbies ao ar livre regulares (otimização):
   - REFORÇAR benefício: Vitamina D (exposição solar), grounding (contato terra =
     redução inflamação), biofilia (contato natureza = redução cortisol)
   - OTIMIZAR TIMING: Manhã (sincronização circadiana), 30min mínimo exposição solar
   - QUANTIFICAR: App para tracking vitamina D (ex: dminder), medir nível sérico
     25-OH vitamina D (meta: 40-60 ng/mL)

[...]
```

---

## Padrões Comuns em Todos os Items

### Estrutura
1. **Clinical Relevance**: Sempre inicia com mecanismo fisiopatológico, depois evidências epidemiológicas
2. **Interpretation Guidelines**: Sempre formato "PADRÃO X: [descrição resposta]" seguido de significado clínico, sistemas comprometidos, investigações, DD
3. **Actionable Insights**: Sempre formato "SE [condição]: [ação]" com priorização (IMEDIATO, MÉDIO PRAZO, LONGO PRAZO quando aplicável)
4. **Red Flags**: Sempre formato "🚩 RED FLAG X: [situação] - Risco: [...] - Ação: [...]"

### Tom
- Técnico mas acessível
- Base em evidências (cita estudos, mecanismos, estatísticas)
- Acionável (médico sabe exatamente o que fazer)
- Foco em medicina funcional (causa raiz, sistemas interconectados)

### Métricas de Qualidade
- Clinical Relevance: 800-1500 caracteres
- Interpretation Guidelines: 1000-2000 caracteres
- Actionable Insights: 1500-2500 caracteres
- Red Flags: 600-1200 caracteres

---

**Nota**: Estes exemplos representam o padrão de qualidade esperado. O Claude deve gerar conteúdo similar em profundidade, especificidade e utilidade clínica.

**Última atualização**: 2026-01-27
