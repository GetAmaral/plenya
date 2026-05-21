# GUIA DE USO - SÍNTESE DE NUTRIÇÃO INFANTIL

**Data:** 2026-02-16
**ScoreItem:** Infância (ID: 019c500a-c35b-7f35-85a5-d935b36b2970)

---

## VISÃO GERAL

Foram criados **3 documentos principais** para fundamentar a revisão científica do ScoreItem "Infância" (histórico alimentar na infância):

```
📄 SINTESE_NUTRICAO_INFANTIL_11_ARTIGOS.md (25 KB)
   └─> Síntese científica completa de 11 artigos (2021-2024)

📄 PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md (12 KB)
   └─> Textos prontos para os 3 campos clínicos

📄 RESUMO_EXECUTIVO_SINTESE_INFANCIA.md (11 KB)
   └─> Visão consolidada e top 10 achados
```

---

## COMO USAR CADA DOCUMENTO

### 1️⃣ SINTESE_NUTRICAO_INFANTIL_11_ARTIGOS.md

**PARA QUÊ:** Documento de referência científica completo

**QUANDO USAR:**
- Quando precisar de fundamentação científica detalhada
- Para validar dados quantitativos (ORs, RRs, HRs)
- Para entender mecanismos fisiopatológicos
- Para consultar referências dos artigos originais

**ESTRUTURA DO DOCUMENTO:**
1. Sumário Executivo
2. Principais Achados Científicos (por artigo)
   - Artigo 1: Breastfeeding Beyond Six Months
   - Artigo 2: Role of breastfeeding in disease prevention
   - Artigo 4: Obesity Prevention
   - Artigo 7: Adequate Nutrition in Early Childhood
   - Artigo 9: Infancy Dietary Patterns
   - Artigo 11: Early Nutritional Education
3. Mecanismos Fisiopatológicos Consolidados
4. Dados Epidemiológicos Relevantes
5. Recomendações Práticas (Tier 1, 2, 3)
6. Intervenções Eficazes Comprovadas
7. Pontos de Atenção para Avaliação Clínica
8. Conclusões e Implicações Clínicas
9. Referências dos 11 Artigos

**DICA:** Use Ctrl+F para buscar por termos específicos (ex: "RR", "OR", "p <", "obesity", "breastfeeding")

---

### 2️⃣ PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md

**PARA QUÊ:** Implementação prática dos campos clínicos

**QUANDO USAR:**
- **AGORA:** Para revisar os textos propostos
- **EM SEGUIDA:** Para implementar no banco de dados
- **DEPOIS:** Como referência do que foi implementado

**ESTRUTURA DO DOCUMENTO:**
- CAMPO: clinical_relevance (versão científica para profissionais)
- CAMPO: patient_explanation (linguagem acessível para pacientes)
- CAMPO: conduct (protocolos clínicos estruturados)
- Justificativa das propostas
- Próximos passos

**COMO USAR NA PRÁTICA:**

#### PASSO 1: REVISÃO CLÍNICA
```
☐ Ler os 3 textos propostos
☐ Validar precisão científica (com médico/nutricionista se possível)
☐ Ajustar linguagem se necessário (sem perder rigor)
☐ Aprovar versão final
```

#### PASSO 2: IMPLEMENTAÇÃO NO BANCO
```sql
-- Copiar texto de PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md
-- e inserir nos campos abaixo:

UPDATE score_items
SET
  clinical_relevance = 'COLAR AQUI O TEXTO DO CAMPO clinical_relevance',
  patient_explanation = 'COLAR AQUI O TEXTO DO CAMPO patient_explanation',
  conduct = 'COLAR AQUI O TEXTO DO CAMPO conduct',
  last_review = NOW()
WHERE id = '019c500a-c35b-7f35-85a5-d935b36b2970';
```

#### PASSO 3: EXECUÇÃO DO UPDATE
```bash
# Via Docker (RECOMENDADO para desenvolvimento)
docker compose exec -T db psql -U plenya_user -d plenya_db << 'EOF'
[COLAR COMANDO SQL COMPLETO AQUI]
EOF
```

#### PASSO 4: VALIDAÇÃO
```bash
# Verificar se campos foram atualizados
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
  name,
  LENGTH(clinical_relevance) as len_clinical,
  LENGTH(patient_explanation) as len_patient,
  LENGTH(conduct) as len_conduct,
  last_review
FROM score_items
WHERE id = '019c500a-c35b-7f35-85a5-d935b36b2970';
"
```

---

### 3️⃣ RESUMO_EXECUTIVO_SINTESE_INFANCIA.md

**PARA QUÊ:** Visão rápida e consolidada

**QUANDO USAR:**
- Para apresentações rápidas
- Para comunicação com stakeholders
- Para entender impacto esperado
- Para acompanhar próximos passos

**DESTAQUES DO DOCUMENTO:**
- Top 10 achados científicos (versão resumida)
- Qualidade da síntese (pontos fortes e limitações)
- Diferenciais em relação a conteúdos genéricos
- Impacto esperado (qualidade clínica, diferenciação de mercado)
- Checklist de próximos passos

**DICA:** Ideal para compartilhar com equipe médica/gestão antes da implementação completa

---

## FLUXO RECOMENDADO DE TRABALHO

### FASE 1: REVISÃO (VOCÊ AGORA) ⬅ ESTÁ AQUI

```
1. Ler RESUMO_EXECUTIVO_SINTESE_INFANCIA.md (5-10 min)
   └─> Entender visão geral e top 10 achados

2. Ler PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md (15-20 min)
   └─> Revisar textos propostos para os 3 campos

3. Consultar SINTESE_NUTRICAO_INFANTIL_11_ARTIGOS.md conforme necessário
   └─> Validar dados específicos ou entender melhor algum ponto
```

### FASE 2: VALIDAÇÃO CLÍNICA (OPCIONAL MAS RECOMENDADO)

```
4. Compartilhar PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md com médico/nutricionista
   └─> Validar precisão científica e aplicabilidade prática

5. Ajustar textos se necessário
   └─> Manter rigor científico, ajustar apenas clareza/aplicabilidade
```

### FASE 3: IMPLEMENTAÇÃO TÉCNICA

```
6. Preparar comando SQL UPDATE com textos aprovados
   └─> Copiar de PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md

7. Executar UPDATE no banco via Docker
   └─> Seguir comandos da seção "PASSO 2" acima

8. Validar atualização
   └─> Verificar lengths dos campos e last_review
```

### FASE 4: TESTE E FEEDBACK

```
9. Testar renderização no frontend web
   └─> Verificar formatação, quebras de linha, listas

10. Coletar feedback inicial de profissionais
    └─> Após 1 semana de uso, perguntar sobre aplicabilidade

11. Ajustar se necessário
    └─> Pequenos refinamentos com base no uso real
```

---

## PERGUNTAS FREQUENTES (FAQ)

### Q1: Preciso ler os 11 artigos originais?

**R:** Não. A síntese já extraiu e consolidou as informações mais relevantes. Mas se quiser validar algum dado específico, os artigos estão no banco:

```sql
-- Para acessar artigos originais
SELECT title, journal, full_content
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '019c500a-c35b-7f35-85a5-d935b36b2970';
```

---

### Q2: Posso modificar os textos propostos?

**R:** SIM, desde que mantenha:
- Rigor científico (dados quantitativos corretos)
- Estrutura lógica (seções organizadas)
- Referências aos estudos (implícitas ou explícitas)

**NÃO modifique:**
- Números (ORs, RRs, HRs) sem validar na fonte original
- Conclusões científicas (sem re-análise dos artigos)

---

### Q3: Os textos estão muito longos?

**R:** São detalhados INTENCIONALMENTE por 3 razões:

1. **clinical_relevance:** Profissionais de saúde VALORIZAM detalhamento científico
2. **patient_explanation:** Pacientes engajados querem entender PROFUNDAMENTE
3. **conduct:** Protocolos precisam ser ESPECÍFICOS para serem aplicáveis

Se necessário, você pode criar versões "resumidas" adicionais, mas mantenha as versões completas como referência.

---

### Q4: Quanto tempo vai demorar a implementação?

**R:** Estimativa conservadora:

- Revisão dos textos: **30-60 min** (você)
- Validação clínica: **1-2 horas** (se envolver médico/nutricionista)
- Implementação técnica: **15-30 min** (SQL + validação)
- Teste no frontend: **30 min**

**TOTAL:** 2-4 horas (distribuídas em 1-2 dias)

---

### Q5: E se encontrar erros depois de implementar?

**R:** Sem problemas! É só fazer um novo UPDATE:

```sql
-- Atualizar campo específico
UPDATE score_items
SET clinical_relevance = '[TEXTO CORRIGIDO]',
    last_review = NOW()
WHERE id = '019c500a-c35b-7f35-85a5-d935b36b2970';
```

A arquitetura do sistema permite atualizações iterativas. Não há risco de "quebrar" nada.

---

### Q6: Posso usar essa metodologia para outros ScoreItems?

**R:** SIM! Esta abordagem é replicável:

1. Identificar artigos linkados ao ScoreItem (tabela article_score_items)
2. Ler full_content dos artigos
3. Sintetizar em documento estruturado (use a estrutura desta síntese como template)
4. Propor textos para os 3 campos clínicos
5. Implementar e validar

**DICA:** Crie um script/template para padronizar o processo e ganhar velocidade.

---

## CHECKLIST FINAL

Use este checklist para acompanhar o progresso:

### REVISÃO
- [ ] Li o RESUMO_EXECUTIVO_SINTESE_INFANCIA.md
- [ ] Li o PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md
- [ ] Consultei a SINTESE_NUTRICAO_INFANTIL_11_ARTIGOS.md (se necessário)
- [ ] Entendi a estrutura e conteúdo dos 3 campos propostos

### VALIDAÇÃO (OPCIONAL)
- [ ] Compartilhei com profissional de saúde para revisão
- [ ] Recebi feedback e ajustei se necessário
- [ ] Versão final dos textos aprovada

### IMPLEMENTAÇÃO
- [ ] Preparei comando SQL UPDATE com textos finais
- [ ] Executei UPDATE no banco de dados
- [ ] Validei que campos foram atualizados (verificar lengths)
- [ ] Testei renderização no frontend web

### FOLLOW-UP
- [ ] Documentei o que foi implementado
- [ ] Configurei coleta de feedback de profissionais (após 1 semana)
- [ ] Agendei revisão de impacto (após 1 mês)

---

## SUPORTE E DÚVIDAS

### Documentação de Referência:
- **SINTESE_NUTRICAO_INFANTIL_11_ARTIGOS.md** - Detalhamento científico
- **PROPOSTA_CAMPOS_CLINICOS_INFANCIA.md** - Textos prontos
- **RESUMO_EXECUTIVO_SINTESE_INFANCIA.md** - Visão consolidada

### Artigos Originais (no banco de dados):
```sql
SELECT id, title, journal, EXTRACT(YEAR FROM publish_date) as year
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '019c500a-c35b-7f35-85a5-d935b36b2970'
ORDER BY publish_date DESC;
```

### Contato com Claude (AI Assistant):
- Para esclarecimentos adicionais, consulte novamente com referência a este guia
- Para novas sínteses de ScoreItems, utilize a mesma metodologia aplicada aqui

---

## PRÓXIMOS SCOREIDS A SEREM REVISADOS (SUGESTÃO)

Após implementar e validar "Infância", considere aplicar mesma metodologia para:

1. ScoreItems com muitos artigos linkados (>5 artigos)
2. ScoreItems sem campos clínicos preenchidos
3. ScoreItems com conteúdo genérico/superficial atual
4. ScoreItems de alta relevância clínica

**Como identificar candidatos:**
```sql
SELECT
  si.id,
  si.name,
  COUNT(asi.article_id) as num_articles,
  LENGTH(si.clinical_relevance) as len_clinical
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
GROUP BY si.id, si.name
HAVING COUNT(asi.article_id) > 5
ORDER BY COUNT(asi.article_id) DESC;
```

---

**BOA IMPLEMENTAÇÃO! 🚀**

Este guia contém tudo que você precisa para revisar, validar e implementar a síntese científica no ScoreItem "Infância" com máxima qualidade e confiança.

Em caso de dúvidas, consulte sempre os 3 documentos principais como referência.
