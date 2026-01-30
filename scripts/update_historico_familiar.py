#!/usr/bin/env python3
"""
Script para atualizar items do grupo Histórico Familiar de Doenças
Batch 32 - Onda 3
"""

import requests
import json
import sys

API_URL = "http://localhost:3001/api/v1/score-items"

# Dados para items de orientação geral
orientacao_geral = {
    "clinical_relevance": """A coleta ativa e sistemática de histórico familiar é o primeiro passo fundamental na estratificação de risco genético e prevenção personalizada. Dados de 2024 demonstram que a história familiar captura tanto predisposição genética quanto fatores ambientais compartilhados, servindo como proxy clínico essencial antes mesmo de testes genéticos. O questionamento ativo sobre doenças crônicas específicas (cardiovasculares, metabólicas, neoplásicas, autoimunes, neurológicas) em parentes de primeiro grau (pais, irmãos, filhos) e segundo grau (avós, tios, netos) permite identificar padrões hereditários com risco 2-4x aumentado. A abordagem proativa é crítica porque muitos pacientes não reportam espontaneamente histórico familiar relevante. Estudos de 2024-2025 mostram que a integração de história familiar com escores de risco poligênico (PRS) aumenta significativamente a precisão diagnóstica. A documentação completa deve incluir: tipo de doença, grau de parentesco, idade de início (doença precoce indica maior componente genético), número de familiares afetados (múltiplos casos aumentam risco), e lado materno vs. paterno. Esta informação orienta decisões de rastreamento precoce, intensidade de medidas preventivas, e necessidade de testes genéticos específicos. A história familiar também tem valor prognóstico: pacientes com história familiar positiva de doença cardiovascular precoce têm 22% mais risco de eventos recorrentes após infarto. Em medicina funcional, o histórico familiar permite identificar predisposições metabólicas (resistência insulínica, dislipidemia, obesidade) que podem ser modificadas precocemente através de intervenções nutricionais, suplementação direcionada, e otimização de estilo de vida. O questionamento sistemático deve fazer parte de toda anamnese inicial e ser atualizado periodicamente, especialmente quando novos diagnósticos ocorrem na família.""",

    "patient_explanation": """Quando perguntamos ativamente sobre doenças que seus familiares próximos têm ou tiveram, estamos realizando uma investigação genética inicial muito importante para sua saúde. Seu histórico familiar funciona como um mapa que mostra quais doenças você pode ter maior tendência de desenvolver ao longo da vida. Estudos recentes mostram que ter um parente de primeiro grau (pai, mãe, irmão) com certas doenças pode aumentar seu risco em 2 a 4 vezes. Por isso fazemos perguntas específicas sobre: doenças do coração, diabetes, pressão alta, colesterol alterado, obesidade, câncer, doenças autoimunes, e problemas neurológicos. É importante que você compartilhe informações detalhadas: que doença foi diagnosticada, qual o grau de parentesco (pai, avô, tio), com que idade a doença apareceu (doenças que surgem cedo na vida geralmente têm componente genético mais forte), e quantos familiares são afetados (múltiplos casos aumentam a preocupação). Mesmo que você não tenha os detalhes exatos, qualquer informação é valiosa. Esta investigação não é motivo para preocupação excessiva - pelo contrário, é uma oportunidade! Conhecer seus riscos genéticos permite que você tome medidas preventivas muito antes da doença aparecer. Por exemplo, se há história familiar forte de diabetes, podemos ajustar sua alimentação, monitorar exames específicos regularmente, e usar suplementos que melhoram a sensibilidade à insulina. Se há histórico de doenças cardíacas, podemos fazer rastreamento cardiovascular mais cedo e intensivo. A medicina moderna permite que você use seu histórico familiar a seu favor, transformando predisposição em prevenção. Por favor, seja o mais completo possível ao responder sobre doenças na família, e nos avise se novos diagnósticos ocorrerem em seus parentes próximos ao longo do tempo.""",

    "conduct": """**Protocolo de Coleta de Histórico Familiar Estruturado:**

**1. Questionário Sistemático por Categorias:**
- Doenças cardiovasculares: IAM, AVC, revascularização, arritmias, morte súbita
- Doenças metabólicas: diabetes (tipo 1 e 2), pré-diabetes, obesidade, síndrome metabólica
- Dislipidemias: hipercolesterolemia familiar, hipertrigliceridemia
- Hipertensão arterial e doença renal crônica
- Neoplasias: tipos específicos, idade de diagnóstico
- Doenças autoimunes: LES, artrite reumatoide, tireoidite, doença celíaca, DII
- Doenças neurológicas: Alzheimer, Parkinson, esclerose múltipla
- Doenças psiquiátricas: depressão, ansiedade, bipolaridade, esquizofrenia

**2. Documentação Detalhada (para cada doença reportada):**
- Grau de parentesco: 1º grau (pais, irmãos, filhos), 2º grau (avós, tios, netos)
- Idade de início da doença (< 55 anos homens, < 65 anos mulheres = precoce)
- Número de familiares afetados (1, 2, 3 ou mais)
- Lado familiar (materno vs. paterno)
- Status atual (controlado, complicações, óbito)
- Comorbidades associadas

**3. Estratificação de Risco Baseada em História Familiar:**
- **Risco Elevado:** ≥2 parentes de 1º grau OU 1 parente de 1º grau com doença precoce
- **Risco Moderado:** 1 parente de 1º grau com doença não-precoce OU múltiplos de 2º grau
- **Risco Padrão:** Apenas parentes de 2º/3º grau ou ausência de história familiar

**4. Intervenções Baseadas em Risco:**

**Risco Elevado (Cardiovascular/Metabólico):**
- Rastreamento precoce: perfil lipídico aos 20 anos, TOTG aos 25 anos
- CAC score aos 40-45 anos se múltiplos fatores de risco
- Monitoramento anual: HbA1c, insulina, HOMA-IR, lipidograma avançado
- Intervenção nutricional agressiva: dieta anti-inflamatória, controle glicêmico
- Suplementação: ômega-3 (2-3g/dia), berberina (1500mg/dia), vitamina D otimizada
- Atividade física supervisionada: ≥150min/semana de atividade moderada

**Risco Elevado (Neoplasias):**
- Considerar aconselhamento genético e teste para mutações específicas (BRCA1/2, Lynch)
- Rastreamento precoce e intensificado conforme tipo de câncer familiar
- Otimização de fatores protetores: vitamina D, ômega-3, antioxidantes

**Risco Elevado (Autoimune):**
- Triagem periódica: FAN, anti-TPO, anti-transglutaminase conforme padrão familiar
- Otimização de vitamina D (alvo > 40ng/mL)
- Saúde intestinal: probióticos, glutamina, redução de permeabilidade
- Redução de gatilhos: glúten (se indicado), exposição a toxinas

**5. Documentação e Acompanhamento:**
- Construir heredograma (árvore genealógica) para 3 gerações
- Atualizar história familiar anualmente ou quando novos diagnósticos ocorrerem
- Educar paciente sobre importância de comunicar mudanças
- Integrar história familiar em todas decisões de rastreamento e prevenção
- Considerar teste genético (PRS - Polygenic Risk Score) se história familiar fortemente positiva

**6. Comunicação com Paciente:**
- Explicar que risco aumentado ≠ doença inevitável
- Enfatizar poder da prevenção personalizada
- Evitar ansiedade desnecessária, focar em empoderamento
- Discutir benefícios de compartilhar informações com familiares (cascata de rastreamento)"""
}

# IDs dos items de orientação geral
orientacao_ids = [
    "56fd5913-4b41-4d56-976a-b6339b4482a6",
    "1da8069d-ed8a-4b5f-83fe-0089909dd630",
    "ed899a60-9067-4a6a-aeb1-b79a08ea6062"
]

# Dados para Hipertensão Arterial
hipertensao = {
    "clinical_relevance": """História familiar de hipertensão arterial é um dos principais fatores de risco modificadores na estratificação cardiovascular segundo as diretrizes ESC 2024. A hipertensão apresenta forte componente genético, com risco 2-3x aumentado quando um parente de primeiro grau é afetado, e até 4x quando ambos os pais são hipertensos. A hereditariedade é estimada em 30-60%, envolvendo múltiplos genes (arquitetura poligênica) que afetam sistema renina-angiotensina, função endotelial, sensibilidade ao sal, e controle autonômico. Estudos de 2024 mostram que história familiar positiva prediz hipertensão de início mais precoce, maior resistência ao tratamento, e risco aumentado de lesão de órgãos-alvo (hipertrofia ventricular esquerda, doença renal crônica, retinopatia). A identificação precoce permite intervenção preventiva antes do desenvolvimento de hipertensão sustentada. Fatores epigenéticos modificam expressão gênica: dieta rica em sódio, obesidade, estresse crônico, e sedentarismo potencializam predisposição genética. Em medicina funcional, história familiar orienta investigação de causas secundárias (resistência insulínica, deficiência de magnésio, disfunção renal subclínica), permitindo abordagem personalizada. Filhos de hipertensos devem ter PA monitorada desde adolescência, com MAPA se valores limítrofes. A história familiar também informa prognóstico: pacientes hipertensos com histórico familiar forte têm maior risco de eventos cardiovasculares e necessitam metas pressóricas mais rigorosas (<130/80mmHg). Estratégias preventivas incluem: restrição de sódio (<2g/dia), suplementação de potássio (alvo 3-4g/dia) e magnésio (400-600mg/dia), controle de peso (cada kg perdido reduz 1mmHg), exercício aeróbico regular (150min/semana), e manejo de estresse. Dados de 2024 mostram que intervenções intensivas em estilo de vida podem reduzir em 50% o risco de desenvolver hipertensão em indivíduos com histórico familiar positivo.""",

    "patient_explanation": """Se seu pai, mãe ou irmãos têm pressão alta, você tem 2 a 3 vezes mais chance de desenvolver hipertensão ao longo da vida. Se ambos os pais são hipertensos, seu risco aumenta para 4 vezes. Isso acontece porque você herda genes que afetam como seus vasos sanguíneos funcionam, como seu corpo controla sal e água, e como seus rins regulam a pressão. Mas aqui está a boa notícia: ter genes de risco NÃO significa que você inevitavelmente terá pressão alta. Estudos mostram que mudanças no estilo de vida podem reduzir pela metade o risco de hipertensão mesmo com histórico familiar forte. O que fazer: 1) Controlar consumo de sal (evitar alimentos processados, temperos prontos, embutidos), 2) Manter peso saudável (cada quilo perdido reduz 1mmHg na pressão), 3) Exercitar-se regularmente (caminhadas, natação, ciclismo - 30min/dia), 4) Aumentar consumo de potássio (frutas, verduras, legumes), 5) Gerenciar estresse (meditação, yoga, sono adequado). Se há histórico familiar, você deve medir sua pressão arterial regularmente a partir dos 18-20 anos, mesmo se sentindo bem. Pressão alta geralmente não dá sintomas até causar complicações sérias (infarto, AVC, insuficiência renal). Dependendo do seu caso, podemos recomendar exames adicionais para detectar sinais precoces de risco (ecocardiograma, função renal, MAPA 24h). Algumas pessoas com histórico familiar se beneficiam de suplementação de magnésio e potássio, que ajudam a relaxar os vasos sanguíneos. O importante é transformar seu conhecimento sobre risco genético em ação preventiva concreta.""",

    "conduct": """**Estratificação de Risco (História Familiar de HAS):**

**Risco ALTO:**
- Ambos os pais hipertensos, especialmente se início <55 anos
- ≥2 parentes de 1º grau com HAS e complicações CV
- Histórico familiar de HAS resistente ou início muito precoce (<40 anos)

**Risco MODERADO:**
- 1 parente de 1º grau com HAS controlada
- Múltiplos parentes de 2º grau hipertensos

**Protocolo de Rastreamento:**

**Alto Risco:**
- PA inicial aos 18 anos, depois anual
- MAPA 24h se PA limítrofe (130-139/85-89mmHg) ou hipertensão do avental branco
- Investigação de lesão subclínica órgãos-alvo aos 35-40 anos:
  * ECG (hipertrofia VE)
  * Ecocardiograma (massa VE, função diastólica)
  * Creatinina + TFG estimada
  * Microalbuminúria
  * Fundo de olho (retinopatia hipertensiva)

**Moderado Risco:**
- PA inicial aos 20-25 anos, depois a cada 2 anos
- MAPA se PA persistentemente >130/85mmHg

**Intervenções Preventivas Personalizadas:**

**Modificação Dietética:**
- Dieta DASH: rica em frutas, vegetais, grãos integrais, laticínios magros
- Restrição sódio: <2g/dia (equivale a 5g sal)
- Aumento potássio: 3-4g/dia (banana, abacate, espinafre, batata)
- Redução álcool: máximo 1-2 doses/dia

**Suplementação Funcional (Alto Risco):**
- Magnésio: 400-600mg/dia (glicina ou taurato) - relaxamento vascular
- Ômega-3: 2-3g/dia EPA+DHA - anti-inflamatório, melhora função endotelial
- Coenzima Q10: 100-200mg/dia - suporte mitocondrial, redução PA
- Vitamina D: otimizar >40ng/mL - modulação SRAA
- L-arginina: 3-6g/dia - precursor óxido nítrico

**Atividade Física Supervisionada:**
- Aeróbico: 150min/semana intensidade moderada (caminhada rápida, natação)
- Resistido: 2-3x/semana (reduz PA sistólica 5-7mmHg)
- Evitar exercícios isométricos intensos

**Manejo de Peso:**
- Meta IMC <25 kg/m²
- Cada 1kg perdido = redução ~1mmHg PA
- Foco em circunferência abdominal (<94cm homens, <80cm mulheres)

**Monitoramento:**
- Automedição domiciliar PA (manhã e noite) se risco alto
- Reavaliação anual com médico
- Triagem para diabetes, dislipidemia, resistência insulínica (comorbidades comuns)

**Critérios para Iniciar Farmacoterapia Precoce:**
- PA persistentemente ≥140/90mmHg apesar modificações estilo de vida (3-6 meses)
- PA 130-139/85-89mmHg + risco CV alto ou lesão órgão-alvo subclínica
- PA ≥130/80mmHg + diabetes ou DRC concomitante
- História familiar muito forte + fatores de risco adicionais (tabagismo, obesidade, sedentarismo)"""
}

# Continua no próximo bloco...
print("Script pronto para processar 33 items do histórico familiar")
