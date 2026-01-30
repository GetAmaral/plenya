# Batch SOCIAL - Índice de Documentação

Navegação completa de toda a documentação do batch SOCIAL.

---

## 📚 Documentação Principal

### 🚀 [SOCIAL-README.md](./SOCIAL-README.md)
**Ponto de partida principal**
- Visão geral completa do sistema
- Como executar (3 passos)
- O que o sistema faz
- Estrutura do projeto
- FAQ e troubleshooting
- Status atual

👉 **Comece aqui se for sua primeira vez**

---

### ⚡ [SOCIAL-QUICK-START.md](./SOCIAL-QUICK-START.md)
**Guia rápido de execução**
- Pré-requisitos (3 comandos)
- Execução (1 comando)
- Verificação de sucesso
- Troubleshooting comum
- Comandos úteis

👉 **Use este se quiser executar AGORA**

---

### 📊 [SOCIAL-BATCH-EXECUTIVE-SUMMARY.md](./SOCIAL-BATCH-EXECUTIVE-SUMMARY.md)
**Sumário executivo completo**
- Visão geral da missão
- Categorias dos 30 items
- Estrutura de enriquecimento
- Base científica
- Arquivos do sistema
- Métricas esperadas
- Impacto esperado
- Próximos passos

👉 **Use este para apresentações ou overview**

---

## 🔬 Documentação Técnica

### 📖 [SOCIAL-ENRICHMENT-METHODOLOGY.md](./SOCIAL-ENRICHMENT-METHODOLOGY.md)
**Metodologia detalhada**
- Categorias dos items SOCIAL detalhadas
- Estrutura de conteúdo clínico (4 campos)
- Convenções de escrita
- Modelo de prompt Claude
- Workflow de execução
- Validação de qualidade

👉 **Use este para entender a metodologia**

---

### 📚 [SOCIAL-SCIENTIFIC-REFERENCES.md](./SOCIAL-SCIENTIFIC-REFERENCES.md)
**50+ referências científicas**
- Estudos chave por categoria:
  - Ambiente Sonoro (WHO, Münzel et al.)
  - Condições de Moradia (Shoemaker, WHO Mould)
  - Exposição Ambiental (Landrigan, Pope & Dockery)
  - Hobbies e Lazer (Blue Zones, Pressman)
  - Luminosidade Natural (Wright, Holick)
  - Profissões (IARC, Kecklund)
- Mecanismos fisiopatológicos
- Guidelines internacionais
- Databases de referência

👉 **Use este para validar evidências científicas**

---

### 📝 [SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md](./SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md)
**Exemplos detalhados de output**
- Exemplo 1: Ambiente Sonoro (completo)
- Exemplo 2: Condições de Moradia (resumido)
- Exemplo 3: Hobbies e Lazer (resumido)
- Padrões comuns em todos os items
- Métricas de qualidade
- Tom e estrutura esperados

👉 **Use este como referência de qualidade**

---

## ✅ Validação e Qualidade

### 🔍 [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md)
**Checklist completo de validação**
1. Validação Técnica
   - Relatório de execução
   - Verificação no banco
   - Validação de tamanho
   - Amostragem de conteúdo
2. Validação de Qualidade
   - Checklist por campo
   - Validação de referências
   - Validação de medicina funcional
3. Validação via API
4. Validação no Frontend
5. Validação de Integridade
6. Validação Clínica (revisão manual)
7. Testes de Regressão
8. Performance e Otimização
9. Checklist Final de Aprovação
10. Ações Pós-Validação

👉 **Use este APÓS executar o batch**

---

## 🛠️ Arquivos Técnicos

### Python
```
/home/user/plenya/scripts/batch_social_enrichment.py
```
**Script principal de enriquecimento**
- Classe `SocialItemEnricher`
- Métodos: login, get_item, update_item, generate_clinical_content
- Lista de 30 UUIDs
- Lógica de retry e error handling
- Geração de relatório JSON

---

### Bash
```
/home/user/plenya/execute_social_batch.sh
```
**Executor com verificações**
- Verifica ANTHROPIC_API_KEY
- Verifica se API está rodando
- Verifica dependências Python
- Executa script Python
- Exibe relatório de sucesso/falha

---

### Output
```
/home/user/plenya/SOCIAL-BATCH-REPORT.json
```
**Relatório de execução** (gerado após rodar)
- Array `success`: UUIDs processados com sucesso
- Array `failed`: UUIDs que falharam
- Timestamps de execução

---

## 📋 Checklists Rápidos

### Pré-Execução
```
[ ] API rodando (docker compose up -d)
[ ] ANTHROPIC_API_KEY exportada
[ ] Dependências Python instaladas
[ ] Backup do banco (opcional)
[ ] Tempo disponível (25+ min)
[ ] Documentação revisada
```

### Pós-Execução
```
[ ] Verificar SOCIAL-BATCH-REPORT.json
[ ] 30/30 items em "success"
[ ] Query banco: 30 items enriquecidos
[ ] Campos não-null verificados
[ ] Teste via API
[ ] Teste no frontend
[ ] Revisão médica agendada
```

---

## 🎯 Navegação por Objetivo

### "Quero executar o batch AGORA"
1. [SOCIAL-QUICK-START.md](./SOCIAL-QUICK-START.md) ← Comece aqui
2. Executar: `./execute_social_batch.sh`
3. [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md) ← Validar

### "Quero entender o que o sistema faz"
1. [SOCIAL-README.md](./SOCIAL-README.md) ← Overview
2. [SOCIAL-BATCH-EXECUTIVE-SUMMARY.md](./SOCIAL-BATCH-EXECUTIVE-SUMMARY.md) ← Detalhes

### "Quero entender a metodologia"
1. [SOCIAL-ENRICHMENT-METHODOLOGY.md](./SOCIAL-ENRICHMENT-METHODOLOGY.md) ← Metodologia
2. [SOCIAL-SCIENTIFIC-REFERENCES.md](./SOCIAL-SCIENTIFIC-REFERENCES.md) ← Evidências

### "Quero ver exemplos de output"
1. [SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md](./SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md) ← Exemplos

### "Já executei, como valido?"
1. [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md) ← Checklist

### "Quero modificar/customizar"
1. Ler: [SOCIAL-ENRICHMENT-METHODOLOGY.md](./SOCIAL-ENRICHMENT-METHODOLOGY.md)
2. Editar: `/home/user/plenya/scripts/batch_social_enrichment.py`
3. Testar: `./execute_social_batch.sh`
4. Validar: [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md)

---

## 📊 Estatísticas da Documentação

| Documento | Linhas | Palavras | Tópicos |
|-----------|--------|----------|---------|
| SOCIAL-README.md | ~600 | ~4,500 | 15+ |
| SOCIAL-QUICK-START.md | ~400 | ~2,500 | 10+ |
| SOCIAL-BATCH-EXECUTIVE-SUMMARY.md | ~500 | ~3,500 | 12+ |
| SOCIAL-ENRICHMENT-METHODOLOGY.md | ~800 | ~6,000 | 20+ |
| SOCIAL-SCIENTIFIC-REFERENCES.md | ~600 | ~4,000 | 30+ |
| SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md | ~700 | ~5,500 | 8+ |
| SOCIAL-POST-EXECUTION-VALIDATION.md | ~900 | ~7,000 | 35+ |
| **TOTAL** | **~4,500** | **~33,000** | **130+** |

---

## 🔄 Fluxo de Trabalho Recomendado

```
1. Planejamento
   ├─ Ler SOCIAL-README.md (10 min)
   └─ Ler SOCIAL-BATCH-EXECUTIVE-SUMMARY.md (15 min)

2. Preparação
   ├─ Verificar pré-requisitos (5 min)
   ├─ Revisar SOCIAL-QUICK-START.md (5 min)
   └─ Configurar ambiente (5 min)

3. Execução
   ├─ Rodar ./execute_social_batch.sh (25 min)
   └─ Monitorar progresso

4. Validação
   ├─ SOCIAL-POST-EXECUTION-VALIDATION.md (30 min)
   ├─ Checklist técnico
   └─ Checklist de qualidade

5. Revisão
   ├─ Amostra aleatória (20 min)
   ├─ Comparar com SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md
   └─ Validar referências em SOCIAL-SCIENTIFIC-REFERENCES.md

6. Aprovação
   ├─ Revisão médica (especialista)
   ├─ Ajustes se necessário
   └─ Deploy para produção

Total estimado: ~2-3 horas (incluindo execução)
```

---

## 🎓 Glossário de Termos

**Batch**: Processamento em lote de múltiplos items (30 neste caso)

**Enriquecimento**: Adição de conteúdo clínico aos 4 campos: clinical_relevance, interpretation_guidelines, actionable_insights, red_flags

**SOCIAL**: Grupo de items relacionados a determinantes sociais da saúde

**Claude Sonnet 4**: Modelo de linguagem avançado da Anthropic usado para gerar conteúdo

**Medicina Funcional**: Abordagem médica focada em causas raiz e sistemas interconectados

**Determinantes Sociais**: Fatores ambientais, ocupacionais e de estilo de vida que impactam saúde

**Red Flags**: Sinais de alerta que exigem ação médica imediata

**SIRS**: Síndrome da Resposta Inflamatória Sistêmica

**Blue Zones**: Regiões com maior concentração de centenários

**NEAT**: Non-Exercise Activity Thermogenesis (termogênese por atividade não-exercício)

---

## 📞 Contato e Suporte

### Documentação Adicional
- Projeto: `/home/user/plenya/CLAUDE.md`
- Arquitetura: `/home/user/plenya/ARQUITETURA.md`
- Sistema de exames: `/home/user/plenya/LAB-TEST-SYSTEM-COMPLETE.md`

### Issues e Melhorias
- **Issues técnicas**: GitHub Issues
- **Melhorias de conteúdo**: Pull Requests
- **Dúvidas**: Consultar documentação ou abrir issue

---

## ✨ Destaques

### Diferenciais do Sistema
✅ **Automatizado**: 30 items enriquecidos em 25 min
✅ **Baseado em Evidências**: 50+ referências científicas
✅ **Medicina Funcional**: Foco em causas raiz
✅ **Acionável**: Intervenções práticas para médicos
✅ **Validação Robusta**: Checklist de 35+ pontos

### Qualidade do Conteúdo
✅ **Mecanismos fisiopatológicos específicos**
✅ **Estudos citados com autor, ano, journal**
✅ **Intervenções práticas e seguras**
✅ **Red flags genuinamente críticos**
✅ **Linguagem técnica mas acessível**

---

## 🚀 Status Atual

```
Sistema: PRONTO PARA EXECUÇÃO ✅

Completado:
✅ 7 documentos de suporte criados
✅ Scripts Python e Bash implementados
✅ 50+ referências científicas compiladas
✅ Exemplos de output detalhados
✅ Checklist de validação completo
✅ Metodologia documentada

Próximo passo:
⏳ Executar ./execute_social_batch.sh
⏳ Validar resultados
⏳ Revisão médica
```

---

## 📅 Timeline

| Fase | Duração | Status |
|------|---------|--------|
| **Planejamento** | 2h | ✅ Completo |
| **Documentação** | 4h | ✅ Completo |
| **Implementação** | 3h | ✅ Completo |
| **Execução** | 25min | ⏳ Pendente |
| **Validação** | 1h | ⏳ Pendente |
| **Revisão Médica** | 2-4h | ⏳ Pendente |
| **Deploy** | 30min | ⏳ Pendente |

---

## 🎯 Objetivos

### Curto Prazo (Hoje)
- [x] Documentação completa
- [ ] Executar batch
- [ ] Validação técnica
- [ ] Relatório de execução

### Médio Prazo (Semana)
- [ ] Revisão médica
- [ ] Ajustes de conteúdo
- [ ] Teste no frontend
- [ ] Feedback de usuários

### Longo Prazo (Mês)
- [ ] Deploy para produção
- [ ] Monitoramento de uso
- [ ] Iteração baseada em feedback
- [ ] Próximo batch (Nutrição?)

---

## 📖 Leitura Recomendada por Perfil

### Desenvolvedor
1. [SOCIAL-README.md](./SOCIAL-README.md) - Overview técnico
2. Scripts: `batch_social_enrichment.py` + `execute_social_batch.sh`
3. [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md) - Testes

### Médico/Clínico
1. [SOCIAL-BATCH-EXECUTIVE-SUMMARY.md](./SOCIAL-BATCH-EXECUTIVE-SUMMARY.md) - Contexto
2. [SOCIAL-SCIENTIFIC-REFERENCES.md](./SOCIAL-SCIENTIFIC-REFERENCES.md) - Evidências
3. [SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md](./SOCIAL-EXPECTED-OUTPUT-EXAMPLES.md) - Qualidade

### Gestor de Projeto
1. [SOCIAL-BATCH-EXECUTIVE-SUMMARY.md](./SOCIAL-BATCH-EXECUTIVE-SUMMARY.md) - Overview
2. [SOCIAL-README.md](./SOCIAL-README.md) - Status e métricas
3. [SOCIAL-POST-EXECUTION-VALIDATION.md](./SOCIAL-POST-EXECUTION-VALIDATION.md) - Critérios de sucesso

### Novo no Projeto
1. [SOCIAL-README.md](./SOCIAL-README.md) - Comece aqui
2. [SOCIAL-QUICK-START.md](./SOCIAL-QUICK-START.md) - Execução rápida
3. [SOCIAL-ENRICHMENT-METHODOLOGY.md](./SOCIAL-ENRICHMENT-METHODOLOGY.md) - Entender metodologia

---

**Última atualização**: 2026-01-27
**Versão da documentação**: 1.0.0
**Status**: ✅ Completo e pronto para uso
