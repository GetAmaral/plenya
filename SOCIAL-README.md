# Batch SOCIAL - Enriquecimento de 30 Items

Sistema completo de enriquecimento de conteúdo clínico para items do grupo SOCIAL, focado em determinantes sociais da saúde na medicina funcional.

---

## Navegação Rápida

| Documento | Descrição | Quando Usar |
|-----------|-----------|-------------|
| **[SOCIAL-BATCH-EXECUTIVE-SUMMARY.md](./SOCIAL-BATCH-EXECUTIVE-SUMMARY.md)** | Visão geral completa do projeto | Entender o contexto geral |
| **[SOCIAL-QUICK-START.md](./SOCIAL-QUICK-START.md)** | Guia rápido de execução (3 passos) | **EXECUTAR AGORA** |
| **[SOCIAL-ENRICHMENT-METHODOLOGY.md](./SOCIAL-ENRICHMENT-METHODOLOGY.md)** | Metodologia detalhada e prompts | Entender a abordagem |
| **[SOCIAL-SCIENTIFIC-REFERENCES.md](./SOCIAL-SCIENTIFIC-REFERENCES.md)** | 50+ referências científicas | Validar evidências |
| **[SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md](./SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md)** | Exemplos de output esperado | Referência de qualidade |
| **[SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md)** | Checklist de validação completo | Após execução |

---

## Execução em 3 Passos

### 1. Pré-requisitos
```bash
# API rodando
docker compose up -d

# ANTHROPIC_API_KEY configurada
export ANTHROPIC_API_KEY='sk-ant-api03-...'

# Dependências Python
pip install anthropic requests
```

### 2. Executar
```bash
chmod +x execute_social_batch.sh
./execute_social_batch.sh
```

### 3. Validar
```bash
# Verificar relatório
cat SOCIAL-BATCH-REPORT.json

# Verificar banco
docker compose exec db psql -U plenya_user plenya_db -c "
SELECT COUNT(*) FROM score_items
WHERE group_name = 'SOCIAL'
AND clinical_relevance IS NOT NULL;"
```

**Esperado**: 30 items enriquecidos

---

## Estrutura do Projeto

```
/home/user/plenya/
│
├── scripts/
│   └── batch_social_enrichment.py    # Script principal Python
│
├── execute_social_batch.sh            # Executor bash
│
├── SOCIAL-README.md                   # Este arquivo
├── SOCIAL-BATCH-EXECUTIVE-SUMMARY.md  # Sumário executivo
├── SOCIAL-QUICK-START.md              # Guia rápido
├── SOCIAL-ENRICHMENT-METHODOLOGY.md   # Metodologia detalhada
├── SOCIAL-SCIENTIFIC-REFERENCES.md    # Referências científicas
├── SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md # Exemplos de output
├── SOCIAL-POST-EXECUTION-VALIDATION.md # Validação pós-execução
│
└── SOCIAL-BATCH-REPORT.json           # Relatório (gerado após execução)
```

---

## O Que Este Sistema Faz

### Input: 30 Items SOCIAL Vazios
```json
{
  "id": "c84412f7-393f-41d0-8bd7-0a28824dbeb0",
  "name": "Nível de ruído no ambiente de trabalho/casa",
  "description": "Avalia exposição a poluição sonora",
  "group_name": "SOCIAL",
  "clinical_relevance": null,
  "interpretation_guidelines": null,
  "actionable_insights": null,
  "red_flags": null
}
```

### Output: Items Enriquecidos com Conteúdo Clínico
```json
{
  "id": "c84412f7-393f-41d0-8bd7-0a28824dbeb0",
  "name": "Nível de ruído no ambiente de trabalho/casa",
  "description": "Avalia exposição a poluição sonora",
  "group_name": "SOCIAL",
  "clinical_relevance": "A exposição crônica a ruído ambiental representa um estressor fisiológico significativo com impactos mensuráveis na saúde cardiometabólica. A ativação repetida do eixo hipotálamo-pituitária-adrenal (HPA) em resposta a ruído >70dB resulta em hipercortisolemia crônica...",
  "interpretation_guidelines": "PADRÃO 1: Exposição a ruído >70dB durante >8h/dia - Significado clínico: Ativação crônica do eixo HPA, dominância simpática - Sistemas comprometidos: Cardiovascular (HTA, arritmias)...",
  "actionable_insights": "1. SE exposição ocupacional >85dB: IMEDIATO - Protetores auriculares (plug + concha = atenuação -30 a -35dB)...",
  "red_flags": "🚩 RED FLAG 1: Exposição ocupacional >85dB sem EPI + perda auditiva progressiva - Risco: PAIR irreversível..."
}
```

---

## Categorias dos 30 Items

| Categoria | Foco Clínico | Items (aprox) |
|-----------|--------------|---------------|
| **Ambiente Sonoro** | Poluição sonora, cortisol, HTA, sono | 4-5 |
| **Condições de Moradia** | Mofo, micotoxinas, SIRS, toxinas domésticas | 5-6 |
| **Espaço para Movimento** | Sedentarismo ambiental, NEAT, design urbano | 3-4 |
| **Exposição Ambiental Externa** | PM2.5, metais pesados, pesticidas | 5-6 |
| **Hobbies e Lazer** | Longevidade, Blue Zones, isolamento social | 4-5 |
| **Luminosidade Natural** | Ritmo circadiano, vitamina D, melatonina | 3-4 |
| **Profissões** | Trabalho noturno, burnout, exposições ocupacionais | 4-5 |

---

## Conteúdo Clínico Gerado

Cada item recebe **4 campos enriquecidos**:

### 1. Clinical Relevance (800-1500 chars)
**O QUE É**: Por que este aspecto social/ambiental importa na medicina funcional

**CONTÉM**:
- Mecanismos fisiopatológicos específicos
- Evidências epidemiológicas (estudos, estatísticas)
- Impacto em sistemas (cardiovascular, endócrino, imune, neurológico)
- Conexão com doenças crônicas

**EXEMPLO**: "A exposição crônica a ruído ambiental >70dB ativa o eixo HPA resultando em hipercortisolemia. Estudos demonstram que cada 10dB aumento correlaciona com aumento de 8% no risco de hipertensão..."

---

### 2. Interpretation Guidelines (1000-2000 chars)
**O QUE É**: Como interpretar as respostas do paciente

**CONTÉM**:
- Padrões de resposta específicos (PADRÃO 1, PADRÃO 2, etc.)
- Significado clínico de cada padrão
- Sistemas potencialmente comprometidos
- Investigações complementares sugeridas
- Diagnósticos diferenciais

**EXEMPLO**: "PADRÃO 1: Mofo visível + sintomas respiratórios - Investigar: IgG/IgE fungos, eosinofilia - DD: Asma alérgica, aspergilose..."

---

### 3. Actionable Insights (1500-2500 chars)
**O QUE É**: Intervenções práticas baseadas nas respostas

**CONTÉM**:
- 5-8 intervenções concretas
- Formato "SE [condição]: [ação]"
- Mudanças ambientais prioritárias
- Suplementação baseada em mecanismos
- Cronograma de reavaliação

**EXEMPLO**: "1. SE ruído >85dB: Protetores auriculares (plug + concha = -30dB). 2. SE cortisol elevado: Ashwagandha 600mg/dia (reduz cortisol 27%)..."

---

### 4. Red Flags (600-1200 chars)
**O QUE É**: Sinais de alerta que exigem ação imediata

**CONTÉM**:
- 3-5 situações críticas
- Formato "🚩 RED FLAG X:"
- Risco claramente descrito
- Ação imediata especificada
- Contexto de urgência

**EXEMPLO**: "🚩 RED FLAG 1: Trabalho noturno >5 anos - Risco: Câncer (IARC Grupo 2A) - Ação: Rastreamento oncológico urgente..."

---

## Base Científica

### Guidelines Internacionais
- WHO Environmental Noise Guidelines (2018)
- WHO Air Quality Guidelines (2021)
- IARC Monographs Vol 124: Night Shift Work (2019)
- EPA National Ambient Air Quality Standards (2023)

### Estudos Chave
- Landrigan et al., "Lancet Commission on Pollution" (2018)
- Münzel et al., "Noise and Cardiovascular Disease" (2021)
- Buettner, "Blue Zones" (2008)
- IARC, "Night Shift Work and Cancer" (2019)

### Mecanismos Fisiopatológicos
- **Ruído** → Eixo HPA → Cortisol ↑ → HTA/DM2
- **Mofo** → Micotoxinas → TGF-β1 ↑ → SIRS
- **PM2.5** → Estresse oxidativo → Aterosclerose
- **Trabalho noturno** → Dessincronização circadiana → Câncer
- **Isolamento social** → Hipercortisolemia → ↑23% mortalidade
- **Déficit luz** → Melatonina ↓ → Depressão/Insônia

Ver todas as referências em: **[SOCIAL-SCIENTIFIC-REFERENCES.md](./SOCIAL-SCIENTIFIC-REFERENCES.md)**

---

## Tecnologia

### Stack
- **LLM**: Claude Sonnet 4 (Anthropic)
- **Linguagem**: Python 3
- **API**: REST (Plenya API)
- **Banco**: PostgreSQL 17

### Workflow
```
1. Login → Obter access token
2. Para cada item (30x):
   a. Fetch item atual
   b. Gerar conteúdo com Claude
   c. Parse JSON
   d. Update item via API
3. Gerar relatório JSON
```

### Características
- ✅ Retry automático em falhas
- ✅ JSON validation
- ✅ Progresso em tempo real
- ✅ Relatório detalhado
- ✅ Tratamento de erros robusto

---

## Métricas Esperadas

| Métrica | Valor Esperado |
|---------|----------------|
| **Items processados** | 30/30 (100%) |
| **Tempo total** | 20-25 minutos |
| **Tempo por item** | 45-60 segundos |
| **Taxa de sucesso** | 100% |
| **Clinical Relevance** | 800-1500 chars |
| **Interpretation Guidelines** | 1000-2000 chars |
| **Actionable Insights** | 1500-2500 chars |
| **Red Flags** | 600-1200 chars |
| **Total por item** | ~4,000-6,500 chars |
| **Total batch** | ~120,000-195,000 chars |

---

## Troubleshooting

### Problema: "ANTHROPIC_API_KEY not found"
**Solução**:
```bash
export ANTHROPIC_API_KEY='sua-chave'
```

### Problema: "API não está respondendo"
**Solução**:
```bash
docker compose up -d
curl http://localhost:3001/health
```

### Problema: "Module 'anthropic' not found"
**Solução**:
```bash
pip install anthropic requests
```

### Problema: "JSON decode error"
**Causa**: Claude retornou resposta malformada
**Solução**: Script tenta remover markdown automaticamente. Se persistir, verificar logs.

### Problema: Alguns items falharam
**Solução**:
```bash
# Ver quais falharam
cat SOCIAL-BATCH-REPORT.json

# Re-executar apenas items falhados
# (editar script para usar apenas IDs de 'failed')
```

---

## Fluxo de Validação Pós-Execução

```
1. Verificar relatório JSON
   ├─ 30 items em "success"?
   └─ 0 items em "failed"?

2. Verificar banco de dados
   ├─ 30 items enriquecidos?
   └─ Campos não-null?

3. Verificar tamanho de conteúdo
   ├─ Médias dentro da faixa?
   └─ Sem conteúdo muito curto/longo?

4. Verificar qualidade de conteúdo
   ├─ Mecanismos fisiopatológicos?
   ├─ Referências científicas?
   ├─ Intervenções práticas?
   └─ Red flags apropriados?

5. Teste via API
   ├─ GET items SOCIAL funciona?
   └─ JSON retornado correto?

6. Teste no frontend
   ├─ Items exibidos?
   └─ Formatação OK?

7. Revisão médica
   ├─ Acurácia científica?
   └─ Relevância clínica?
```

Ver checklist completo em: **[SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md)**

---

## Próximos Passos

### Após Execução Bem-Sucedida
1. ✅ Validar conteúdo (checklist pós-execução)
2. ✅ Revisar com especialista médico
3. ✅ Testar no frontend
4. ✅ Coletar feedback de usuários
5. ✅ Iterar e melhorar
6. ✅ Documentar lições aprendidas
7. ✅ Preparar próximo batch

### Potenciais Próximos Batches
- **Nutrição**: Items sobre padrão alimentar
- **Suplementação**: Items sobre uso de suplementos
- **Relações**: Items sobre relacionamentos interpessoais
- **Espiritualidade**: Items sobre propósito e significado

---

## FAQ

**P: Quanto tempo leva?**
R: 20-25 minutos para os 30 items.

**P: Posso interromper e continuar depois?**
R: Sim, mas precisa modificar script para pular items já processados.

**P: O que fazer se alguns items falharem?**
R: Verificar `SOCIAL-BATCH-REPORT.json` para IDs falhados e re-executar apenas esses.

**P: Preciso de conhecimento médico para executar?**
R: Não, é automatizado. Mas revisão médica pós-execução é recomendada.

**P: Posso customizar os prompts?**
R: Sim, edite `scripts/batch_social_enrichment.py`, método `generate_clinical_content()`.

**P: Como garantir qualidade científica?**
R: (1) Prompts baseados em literatura, (2) Claude Sonnet 4 (modelo avançado), (3) Revisão médica pós-execução.

**P: Quanto custa (API Claude)?**
R: ~$0.50-1.00 para os 30 items (Claude Sonnet 4, ~150K tokens total).

**P: Posso usar outro LLM?**
R: Sim, mas precisaria adaptar código. Claude Sonnet 4 é recomendado pela qualidade médica.

---

## Suporte

### Documentação
- **Projeto geral**: `/home/user/plenya/CLAUDE.md`
- **Arquitetura**: `/home/user/plenya/ARQUITETURA.md`
- **Sistema de exames**: `/home/user/plenya/LAB-TEST-SYSTEM-COMPLETE.md`

### Contato
- **Issues técnicas**: GitHub Issues
- **Melhorias de conteúdo**: Pull Requests
- **Dúvidas**: Consultar documentação ou abrir issue

---

## Contribuindo

### Melhorar Conteúdo Clínico
1. Identificar item a melhorar
2. Pesquisar literatura atualizada
3. Propor mudanças específicas
4. Submeter PR com referências

### Melhorar Sistema
1. Identificar bug ou limitação
2. Propor solução
3. Implementar (se técnico)
4. Submeter PR com testes

### Reportar Problemas
1. Verificar se já foi reportado
2. Incluir logs e contexto
3. Passos para reproduzir
4. Comportamento esperado vs. atual

---

## Licença e Uso

Este sistema faz parte do projeto **Plenya EMR**.

- **Código**: Licença do projeto
- **Conteúdo clínico gerado**: Para uso interno do Plenya
- **Referências científicas**: Citadas apropriadamente, domínio público ou fair use

---

## Agradecimentos

### Literatura Científica
- WHO (World Health Organization)
- IARC (International Agency for Research on Cancer)
- Lancet Commission on Pollution
- Blue Zones Research

### Tecnologia
- Anthropic (Claude Sonnet 4)
- PostgreSQL
- Python ecosystem

---

## Changelog

### v1.0.0 (2026-01-27)
- ✅ Sistema completo de enriquecimento implementado
- ✅ 30 items SOCIAL identificados
- ✅ Metodologia documentada
- ✅ 50+ referências científicas compiladas
- ✅ Exemplos de output criados
- ✅ Validação pós-execução detalhada
- ✅ Quick start guide
- ✅ README completo

---

## Status Atual

```
[█████████░] 90% Pronto para Execução

Completo:
✅ Script Python implementado
✅ Executor bash criado
✅ Documentação completa
✅ Metodologia definida
✅ Referências compiladas
✅ Exemplos de output
✅ Validação pós-execução

Pendente:
⏳ Execução do batch (aguardando comando)
⏳ Validação de resultados
⏳ Revisão médica
```

---

## Comando para Iniciar

```bash
./execute_social_batch.sh
```

**Próxima etapa**: Executar e validar os 30 items SOCIAL.

---

**Criado**: 2026-01-27
**Última atualização**: 2026-01-27
**Versão**: 1.0.0
**Status**: ✅ Pronto para Execução
