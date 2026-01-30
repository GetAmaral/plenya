# Batch SOCIAL - Sumário Executivo

## Visão Geral

**Missão**: Enriquecimento de 30 items do grupo SOCIAL com conteúdo clínico profundo baseado em determinantes sociais da saúde na medicina funcional.

**Metodologia**: Pesquisa científica avançada focada em mecanismos fisiopatológicos, evidências epidemiológicas e intervenções práticas.

**Tecnologia**: Claude Sonnet 4 com prompts especializados para geração de conteúdo médico de alta qualidade.

---

## Categorias dos 30 Items

| Categoria | Items | Foco Clínico |
|-----------|-------|--------------|
| **Ambiente Sonoro** | 4-5 | Poluição sonora, cortisol, sono, HTA |
| **Condições de Moradia** | 5-6 | Mofo, micotoxinas, SIRS, toxinas domésticas |
| **Espaço para Movimento** | 3-4 | Sedentarismo ambiental, NEAT, design urbano |
| **Exposição Ambiental Externa** | 5-6 | PM2.5, metais pesados, pesticidas |
| **Hobbies e Lazer** | 4-5 | Longevidade, Blue Zones, isolamento social |
| **Luminosidade Natural** | 3-4 | Ritmo circadiano, vitamina D, melatonina |
| **Profissões** | 4-5 | Trabalho noturno, burnout, exposições ocupacionais |

---

## Estrutura de Enriquecimento

Cada item recebe **4 campos clínicos**:

### 1. Clinical Relevance (800-1500 chars)
- Mecanismos fisiopatológicos específicos
- Evidências epidemiológicas
- Impacto em sistemas (cardiovascular, endócrino, imune, neurológico)
- Conexão com doenças crônicas

### 2. Interpretation Guidelines (1000-2000 chars)
- Padrões de resposta e significado clínico
- Sistemas potencialmente comprometidos
- Investigações complementares sugeridas
- Diagnósticos diferenciais

### 3. Actionable Insights (1500-2500 chars)
- 5-8 intervenções práticas
- Mudanças ambientais prioritárias
- Suplementação baseada em evidências
- Cronograma de reavaliação

### 4. Red Flags (600-1200 chars)
- 3-5 sinais de alerta críticos
- Situações que exigem ação imediata
- Riscos de mortalidade/morbidade grave

---

## Base Científica

### Fontes de Evidência

**Guidelines Internacionais**:
- WHO Environmental Noise Guidelines (2018)
- WHO Air Quality Guidelines (2021)
- IARC Monographs Vol 124: Night Shift Work (2019)
- EPA National Ambient Air Quality Standards (2023)

**Journals de Referência**:
- The Lancet (Landrigan Commission on Pollution, 2018)
- PNAS (Circadian rhythm, 2013)
- European Heart Journal (Noise & CVD, 2021)
- JAMA (Night shift & mortality, 2016)

**Estudos Chave**:
- Blue Zones Research (Buettner, 2008)
- Framingham Heart Study
- Nurses' Health Study (74,862 participantes, 22 anos)
- WHO Dampness and Mould Guidelines (2009)

### Mecanismos Fisiopatológicos

| Exposição | Mecanismo | Desfecho |
|-----------|-----------|----------|
| Ruído >70dB | Eixo HPA → Cortisol ↑ | HTA, DM2, síndrome metabólica |
| Mofo/Micotoxinas | Inflamação → TGF-β1 ↑ | SIRS, autoimunidade, fadiga crônica |
| PM2.5 | Estresse oxidativo | Aterosclerose, IAM, AVE |
| Trabalho noturno | Dessincronização circadiana | Câncer, DM2, obesidade |
| Isolamento social | Hipercortisolemia crônica | ↑ 23% mortalidade |
| Déficit luz natural | Melatonina ↓, Vitamina D ↓ | Depressão, insônia, osteoporose |

---

## Arquivos do Sistema

### Scripts
```
/home/user/plenya/scripts/batch_social_enrichment.py
└── Script principal Python com lógica de enriquecimento

/home/user/plenya/execute_social_batch.sh
└── Executor bash com verificações de ambiente
```

### Documentação
```
/home/user/plenya/SOCIAL-ENRICHMENT-METHODOLOGY.md
└── Metodologia completa, prompts, workflow

/home/user/plenya/SOCIAL-SCIENTIFIC-REFERENCES.md
└── 50+ referências científicas por categoria

/home/user/plenya/SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md
└── Exemplos detalhados de output esperado

/home/user/plenya/SOCIAL-QUICK-START.md
└── Guia rápido de execução (3 passos)

/home/user/plenya/SOCIAL-BATCH-EXECUTIVE-SUMMARY.md
└── Este documento
```

### Output
```
/home/user/plenya/SOCIAL-BATCH-REPORT.json
└── Relatório de execução (success/failed)
```

---

## Execução

### Pré-requisitos
1. API rodando: `docker compose up -d`
2. ANTHROPIC_API_KEY: `export ANTHROPIC_API_KEY='sk-ant-...'`
3. Dependências Python: `pip install anthropic requests`

### Comando
```bash
chmod +x execute_social_batch.sh
./execute_social_batch.sh
```

### Métricas Esperadas
- **Tempo total**: 20-25 minutos
- **Taxa de sucesso**: 100% (30/30 items)
- **Tempo por item**: ~45-60 segundos
- **Total caracteres/item**: ~4,000-6,500

---

## Diferenciais da Metodologia

### 1. Medicina Funcional
- Foco em **causas raiz**, não apenas sintomas
- Avaliação de **sistemas interconectados**
- Ênfase em **intervenções de estilo de vida**
- Suplementação **baseada em mecanismos fisiopatológicos**

### 2. Determinantes Sociais
- Reconhecimento que **ambiente molda saúde**
- Avaliação de fatores frequentemente negligenciados
- Intervenções ambientais como **primeira linha**
- Consideração de **barreiras socioeconômicas**

### 3. Evidências Científicas
- Cita estudos específicos, não generalidades
- Menciona **mecanismos bioquímicos precisos**
- Quantifica riscos (ex: "aumento 8% risco HTA")
- Base em **guidelines internacionais**

### 4. Acionabilidade
- Médico sabe **exatamente o que fazer**
- Intervenções priorizadas (imediato/médio/longo prazo)
- Cronogramas de reavaliação definidos
- Métricas de sucesso claras

---

## Impacto Esperado

### Para Médicos
- **Avaliação holística** de determinantes sociais
- **Guidance clínico** para interpretação de respostas
- **Intervenções práticas** além de farmacoterapia
- **Red flags** para identificar situações críticas

### Para Pacientes
- **Validação** de que ambiente importa
- **Empoderamento** com mudanças práticas
- **Educação** sobre impactos ambientais na saúde
- **Esperança** de que mudanças são possíveis

### Para o Sistema
- **Dados estruturados** para pesquisa
- **Padrões epidemiológicos** de exposições
- **Base** para políticas de saúde pública
- **Correlações** entre ambiente e desfechos clínicos

---

## Métricas de Qualidade

### Conteúdo
- [ ] Mecanismos fisiopatológicos específicos mencionados
- [ ] Estudos científicos citados (autor, ano, journal)
- [ ] Intervenções práticas e acionáveis
- [ ] Red flags identificados corretamente
- [ ] Tom técnico mas acessível

### Técnico
- [ ] JSON válido (sem erros de parsing)
- [ ] 4 campos preenchidos por item
- [ ] Caracteres dentro da faixa esperada
- [ ] Sem erros de API ou autenticação
- [ ] 100% items processados com sucesso

### Clínico
- [ ] Conteúdo revisado por médico especialista
- [ ] Coerência com guidelines internacionais
- [ ] Intervenções alinhadas com medicina funcional
- [ ] Consideração de contraindicações e riscos

---

## Próximos Passos

### Fase 1: Execução (Agora)
1. Rodar `./execute_social_batch.sh`
2. Monitorar progresso no terminal
3. Verificar relatório `SOCIAL-BATCH-REPORT.json`

### Fase 2: Validação (Semana 1)
1. Revisão médica de conteúdo
2. Teste no frontend (exibição campos)
3. Correções pontuais se necessário

### Fase 3: Deploy (Semana 2)
1. Commit das mudanças no banco
2. Atualização documentação médica
3. Treinamento equipe clínica

### Fase 4: Iteração (Contínuo)
1. Coleta de feedback de médicos
2. Ajustes baseados em uso real
3. Adição de novos insights científicos
4. Expansão para outros grupos (próximos batches)

---

## Lições Aprendidas (Batches Anteriores)

### O Que Funcionou Bem
- ✅ Prompts estruturados com exemplos claros
- ✅ Validação JSON automática
- ✅ Retry logic para falhas de API
- ✅ Relatórios detalhados de progresso
- ✅ Documentação extensa antecipada

### O Que Melhoramos
- 🔄 Prompts mais específicos para medicina funcional
- 🔄 Ênfase maior em determinantes sociais
- 🔄 Red flags mais práticos e críticos
- 🔄 Referências científicas mais robustas
- 🔄 Exemplos de output esperado

### Riscos Mitigados
- ⚠️ Claude retornar JSON malformado → Parsing com cleanup de markdown
- ⚠️ API timeout → Retry com backoff exponencial
- ⚠️ Token expirado → Re-login automático
- ⚠️ Conteúdo genérico → Prompts com especificidade máxima

---

## Recursos Adicionais

### Documentação Projeto
- `/home/user/plenya/CLAUDE.md` - Diretrizes gerais do projeto
- `/home/user/plenya/ARQUITETURA.md` - Arquitetura técnica
- `/home/user/plenya/LAB-TEST-SYSTEM-COMPLETE.md` - Sistema de exames

### Batches Anteriores (Referência)
- `BATCH-50-EXAMES-FINAL-REPORT.md` - 50 items exames laboratoriais
- `COGNICAO-BATCH-SUMMARY.md` - Items de cognição
- `STRESS-GROUP-ENRICHMENT-COMPLETE.md` - Items de estresse
- `SONO-ENRICHMENT-REPORT.md` - Items de sono

### Contato e Suporte
- **Issues**: Para bugs técnicos
- **Pull Requests**: Para melhorias no conteúdo
- **Documentação**: Sempre atualizada no repositório

---

## Checklist Final

Antes de executar, confirme:

- [ ] API está rodando (`curl http://localhost:3001/health`)
- [ ] ANTHROPIC_API_KEY está configurada (`echo $ANTHROPIC_API_KEY`)
- [ ] Dependências Python instaladas (`python3 -c "import anthropic, requests"`)
- [ ] Backup do banco realizado (opcional mas recomendado)
- [ ] Documentação revisada (metodologia + referências + exemplos)
- [ ] Tempo disponível (25+ minutos sem interrupções)

---

## Conclusão

Este batch SOCIAL representa um marco na integração de **determinantes sociais da saúde** com **medicina funcional baseada em evidências**. Ao enriquecer 30 items com conteúdo clínico profundo, estamos capacitando médicos a:

1. **Reconhecer** que ambiente, trabalho e lazer impactam profundamente a saúde
2. **Avaliar** sistematicamente exposições e fatores de risco socioambientais
3. **Intervir** de forma prática e baseada em ciência
4. **Prevenir** doenças crônicas abordando causas raiz

O diferencial está na **acionabilidade**: cada item não apenas educa, mas **equipa o médico com ferramentas concretas** para transformar insights em mudanças reais na vida dos pacientes.

---

**Status**: Pronto para execução
**Última atualização**: 2026-01-27
**Autor**: Sistema de enriquecimento automatizado Plenya
**Revisão**: Pendente (pós-execução)

---

**COMANDO PARA INICIAR**:
```bash
./execute_social_batch.sh
```
