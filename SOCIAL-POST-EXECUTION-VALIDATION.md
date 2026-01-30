# Validação Pós-Execução - Batch SOCIAL

Checklist completo para validar o enriquecimento dos 30 items SOCIAL após execução.

---

## 1. Validação Técnica

### 1.1 Relatório de Execução

```bash
# Verificar relatório gerado
cat /home/user/plenya/SOCIAL-BATCH-REPORT.json
```

**Esperado**:
```json
{
  "success": [
    "c84412f7-393f-41d0-8bd7-0a28824dbeb0",
    "91e450db-29df-4a78-8741-441f89630ff7",
    ...
  ],
  "failed": []
}
```

**Checklist**:
- [ ] Arquivo `SOCIAL-BATCH-REPORT.json` existe
- [ ] Array `success` tem 30 items
- [ ] Array `failed` está vazio
- [ ] Todos os 30 UUIDs estão em `success`

---

### 1.2 Verificação no Banco de Dados

```sql
-- Conectar ao banco
docker compose exec db psql -U plenya_user plenya_db

-- Contar items SOCIAL enriquecidos
SELECT
    COUNT(*) as total_enriquecidos
FROM score_items
WHERE
    group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL
    AND clinical_relevance != ''
    AND interpretation_guidelines IS NOT NULL
    AND interpretation_guidelines != ''
    AND actionable_insights IS NOT NULL
    AND actionable_insights != ''
    AND red_flags IS NOT NULL
    AND red_flags != '';
```

**Esperado**: `total_enriquecidos = 30`

**Checklist**:
- [ ] 30 items retornados
- [ ] Todos os 4 campos clínicos preenchidos
- [ ] Nenhum campo vazio ou NULL

---

### 1.3 Validação de Tamanho de Conteúdo

```sql
-- Verificar tamanho médio dos campos clínicos
SELECT
    AVG(LENGTH(clinical_relevance)) as avg_clinical,
    AVG(LENGTH(interpretation_guidelines)) as avg_interpretation,
    AVG(LENGTH(actionable_insights)) as avg_actionable,
    AVG(LENGTH(red_flags)) as avg_red_flags,
    MIN(LENGTH(clinical_relevance)) as min_clinical,
    MAX(LENGTH(clinical_relevance)) as max_clinical
FROM score_items
WHERE
    group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL;
```

**Esperado**:
- `avg_clinical`: 800-1500 caracteres
- `avg_interpretation`: 1000-2000 caracteres
- `avg_actionable`: 1500-2500 caracteres
- `avg_red_flags`: 600-1200 caracteres

**Checklist**:
- [ ] Médias dentro das faixas esperadas
- [ ] Mínimos > 500 caracteres (não muito curto)
- [ ] Máximos < 4000 caracteres (não muito longo)

---

### 1.4 Amostragem de Conteúdo

```sql
-- Ver 3 exemplos aleatórios
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as len_clinical,
    LENGTH(interpretation_guidelines) as len_interpretation,
    LENGTH(actionable_insights) as len_actionable,
    LENGTH(red_flags) as len_red_flags,
    SUBSTRING(clinical_relevance, 1, 150) as preview_clinical
FROM score_items
WHERE group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL
ORDER BY RANDOM()
LIMIT 3;
```

**Checklist**:
- [ ] Preview mostra conteúdo clínico relevante
- [ ] Linguagem técnica mas acessível
- [ ] Menção a mecanismos fisiopatológicos

---

## 2. Validação de Qualidade de Conteúdo

### 2.1 Checklist por Campo

Para cada um dos 30 items, validar:

#### Clinical Relevance
- [ ] Menciona mecanismos fisiopatológicos específicos
- [ ] Cita evidências epidemiológicas (estudos, estatísticas)
- [ ] Explica impacto em sistemas (cardiovascular, endócrino, etc.)
- [ ] Conecta com doenças crônicas (DM2, HTA, câncer, autoimunidade)
- [ ] Linguagem técnica mas compreensível
- [ ] 2-3 parágrafos bem estruturados

#### Interpretation Guidelines
- [ ] Formato "PADRÃO X:" para diferentes respostas
- [ ] Significado clínico de cada padrão
- [ ] Sistemas potencialmente comprometidos
- [ ] Investigações complementares sugeridas
- [ ] Diagnósticos diferenciais listados
- [ ] Guia prático e acionável

#### Actionable Insights
- [ ] Formato "SE [condição]: [ação]"
- [ ] 5-8 intervenções listadas
- [ ] Mudanças ambientais prioritárias
- [ ] Suplementação quando aplicável
- [ ] Cronograma de reavaliação
- [ ] Intervenções viáveis e práticas

#### Red Flags
- [ ] Formato "🚩 RED FLAG X:"
- [ ] 3-5 sinais de alerta
- [ ] Risco claramente descrito
- [ ] Ação imediata especificada
- [ ] Situações críticas identificadas
- [ ] Senso de urgência apropriado

---

### 2.2 Validação de Referências Científicas

Verificar se o conteúdo menciona:

**Estudos Esperados por Categoria**:

#### Ambiente Sonoro
- [ ] WHO Environmental Noise Guidelines
- [ ] Münzel et al. (European Heart Journal)
- [ ] Mecanismo: Eixo HPA → cortisol
- [ ] Estatística: 10dB = 8% aumento risco

#### Condições de Moradia (Mofo)
- [ ] SIRS (Síndrome Resposta Inflamatória Sistêmica)
- [ ] Shoemaker/Brewer (micotoxinas)
- [ ] Mecanismo: TGF-β1, C4a, MSH
- [ ] WHO Dampness and Mould Guidelines

#### Exposição Ambiental
- [ ] Landrigan et al. (Lancet Commission on Pollution)
- [ ] Pope & Dockery (PM2.5 cardiovascular)
- [ ] Mecanismo: Estresse oxidativo
- [ ] Estatística: 10µg/m³ PM2.5 = 6% mortalidade

#### Hobbies e Lazer
- [ ] Blue Zones (Buettner)
- [ ] Pressman et al. (cortisol, IL-6)
- [ ] Holt-Lunstad (isolamento = 15 cigarros/dia)
- [ ] Mecanismo: Redução cortisol, BDNF

#### Luminosidade Natural
- [ ] Wright et al. (circadian entrainment)
- [ ] Holick (vitamina D deficiency)
- [ ] Rosenthal (SAD - Seasonal Affective Disorder)
- [ ] Mecanismo: Melatonina, SCN

#### Profissões
- [ ] IARC Grupo 2A (trabalho noturno)
- [ ] Kecklund & Axelsson (shift work)
- [ ] Salvagioni (burnout cardiovascular)
- [ ] Mecanismo: Dessincronização clock genes

---

### 2.3 Validação de Medicina Funcional

Verificar se o conteúdo reflete princípios de medicina funcional:

- [ ] Foco em **causas raiz**, não apenas sintomas
- [ ] Avaliação de **sistemas interconectados**
- [ ] Ênfase em **intervenções de estilo de vida** (alimentação, movimento, sono, estresse)
- [ ] Suplementação baseada em **mecanismos fisiopatológicos**
- [ ] Consideração de **fatores ambientais**
- [ ] Abordagem **personalizada** (diferentes padrões de resposta)
- [ ] **Cronograma de reavaliação** (não apenas prescrever e esquecer)

---

## 3. Validação via API

### 3.1 Teste de Item Individual

```bash
# Login e obter token
TOKEN=$(curl -s -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"import@plenya.com","password":"Import@123456"}' | \
  python3 -c "import sys, json; print(json.load(sys.stdin)['accessToken'])")

# Buscar item específico (exemplo: Ambiente Sonoro)
curl -s -X GET "http://localhost:3001/api/v1/score-items/c84412f7-393f-41d0-8bd7-0a28824dbeb0" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool > item_sample.json

# Visualizar
cat item_sample.json
```

**Checklist**:
- [ ] Status 200 OK
- [ ] JSON válido retornado
- [ ] Campos `clinical_relevance`, `interpretation_guidelines`, `actionable_insights`, `red_flags` preenchidos
- [ ] Conteúdo em português
- [ ] Sem erros de encoding

---

### 3.2 Teste de Listagem SOCIAL

```bash
# Listar todos os items SOCIAL
curl -s -X GET "http://localhost:3001/api/v1/score-items?group=SOCIAL&limit=100" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool > social_items_full.json

# Contar items retornados
python3 -c "import json; data=json.load(open('social_items_full.json')); print(f'Total items: {len(data[\"items\"])}')"
```

**Esperado**: `Total items: 30`

**Checklist**:
- [ ] 30 items retornados
- [ ] Todos têm `group_name: "SOCIAL"`
- [ ] Todos têm campos clínicos preenchidos

---

## 4. Validação no Frontend

### 4.1 Acessar Interface Web

```bash
# Verificar se web está rodando
curl -s http://localhost:3000 | head -n 1
```

**Checklist**:
- [ ] Frontend acessível em http://localhost:3000
- [ ] Login funcional
- [ ] Navegação para página de scores

---

### 4.2 Visualizar Items SOCIAL

Na interface web:

1. **Login**: import@plenya.com / Import@123456
2. **Navegar**: Dashboard → Scores ou Items
3. **Filtrar**: Grupo = SOCIAL
4. **Selecionar**: Qualquer item SOCIAL
5. **Verificar**: Campos clínicos exibidos corretamente

**Checklist**:
- [ ] Items SOCIAL listados corretamente
- [ ] Campos clínicos visíveis (Clinical Relevance, etc.)
- [ ] Formatação adequada (parágrafos, listas)
- [ ] Sem erros de encoding (caracteres especiais corretos)
- [ ] Responsividade OK (mobile/desktop)

---

## 5. Validação de Integridade

### 5.1 Verificar Timestamps

```sql
-- Verificar quando items foram atualizados
SELECT
    id,
    name,
    updated_at,
    EXTRACT(EPOCH FROM (NOW() - updated_at)) / 60 as minutes_ago
FROM score_items
WHERE group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL
ORDER BY updated_at DESC;
```

**Checklist**:
- [ ] Todos os 30 items têm `updated_at` recente (< 1h se acabou de rodar)
- [ ] Timestamps consistentes (todos próximos entre si)
- [ ] Nenhum item com data muito antiga (indicaria falha silenciosa)

---

### 5.2 Verificar Caracteres Especiais

```sql
-- Verificar se há problemas de encoding
SELECT
    id,
    name,
    clinical_relevance
FROM score_items
WHERE
    group_name = 'SOCIAL'
    AND (
        clinical_relevance LIKE '%�%'  -- Caractere de erro encoding
        OR clinical_relevance LIKE '%\\u%'  -- Unicode mal-formado
    );
```

**Esperado**: 0 resultados

**Checklist**:
- [ ] Nenhum caractere de erro encontrado
- [ ] Acentuação correta (português)
- [ ] Símbolos especiais corretos (🚩, →, °)

---

## 6. Validação Clínica (Revisão Manual)

### 6.1 Amostra Aleatória (5 items)

Selecionar 5 items aleatórios e revisar **manualmente**:

```sql
SELECT id, name
FROM score_items
WHERE group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL
ORDER BY RANDOM()
LIMIT 5;
```

Para cada item:

**Acurácia Científica**:
- [ ] Mecanismos fisiopatológicos corretos
- [ ] Estatísticas plausíveis e bem citadas
- [ ] Estudos mencionados são reais (verificar PubMed se necessário)
- [ ] Sem afirmações exageradas ou não-comprovadas

**Relevância Clínica**:
- [ ] Informação útil para tomada de decisão
- [ ] Intervenções viáveis na prática
- [ ] Red flags genuinamente críticos
- [ ] Linguagem apropriada para médico

**Medicina Funcional**:
- [ ] Abordagem sistêmica (não reducionista)
- [ ] Foco em prevenção e otimização
- [ ] Considera contexto do paciente
- [ ] Intervenções além de farmacoterapia

---

### 6.2 Revisão por Especialista (Opcional mas Recomendado)

Enviar amostra de 10 items para médico especialista em medicina funcional revisar:

**Perguntas-chave**:
1. O conteúdo é cientificamente acurado?
2. As intervenções são práticas e seguras?
3. Os red flags estão apropriados?
4. O tom e linguagem são adequados?
5. Falta alguma informação crítica?
6. Há algo que deveria ser removido/modificado?

**Checklist**:
- [ ] Revisão médica realizada
- [ ] Feedback documentado
- [ ] Correções implementadas (se necessário)
- [ ] Aprovação final recebida

---

## 7. Testes de Regressão

### 7.1 Verificar Items Não-SOCIAL

```sql
-- Garantir que items de outros grupos não foram afetados
SELECT
    group_name,
    COUNT(*) as total,
    SUM(CASE WHEN updated_at > NOW() - INTERVAL '1 hour' THEN 1 ELSE 0 END) as updated_last_hour
FROM score_items
WHERE group_name != 'SOCIAL'
GROUP BY group_name;
```

**Checklist**:
- [ ] Items de outros grupos não têm `updated_last_hour` > 0
- [ ] Apenas grupo SOCIAL foi modificado
- [ ] Totais de outros grupos inalterados

---

### 7.2 Verificar Outras Tabelas

```sql
-- Garantir que nenhuma outra tabela foi afetada acidentalmente
SELECT COUNT(*) FROM score_levels;
SELECT COUNT(*) FROM patients;
SELECT COUNT(*) FROM users;
```

**Checklist**:
- [ ] Contagens iguais a antes da execução
- [ ] Nenhuma tabela relacionada foi modificada

---

## 8. Performance e Otimização

### 8.1 Tempo de Query

```sql
-- Testar velocidade de busca de items enriquecidos
EXPLAIN ANALYZE
SELECT *
FROM score_items
WHERE group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL;
```

**Checklist**:
- [ ] Query executa em <100ms
- [ ] Índices sendo utilizados corretamente
- [ ] Sem seq scans desnecessários

---

### 8.2 Tamanho de Dados

```sql
-- Verificar tamanho dos dados clínicos adicionados
SELECT
    pg_size_pretty(pg_total_relation_size('score_items')) as table_size,
    pg_size_pretty(SUM(LENGTH(clinical_relevance) +
                       LENGTH(interpretation_guidelines) +
                       LENGTH(actionable_insights) +
                       LENGTH(red_flags))::bigint) as clinical_content_size
FROM score_items
WHERE group_name = 'SOCIAL';
```

**Checklist**:
- [ ] Tamanho de conteúdo clínico: ~120-200KB (30 items × 4-6KB)
- [ ] Tamanho de tabela não explodiu (crescimento proporcional)

---

## 9. Checklist Final de Aprovação

### Técnico
- [ ] 30/30 items processados com sucesso
- [ ] Relatório JSON gerado corretamente
- [ ] Banco de dados atualizado (30 items com 4 campos cada)
- [ ] API retornando dados corretamente
- [ ] Frontend exibindo campos enriquecidos
- [ ] Sem erros de encoding ou JSON malformado
- [ ] Performance de queries OK

### Conteúdo
- [ ] Clinical Relevance: mecanismos + evidências (800-1500 chars)
- [ ] Interpretation Guidelines: padrões práticos (1000-2000 chars)
- [ ] Actionable Insights: 5-8 intervenções (1500-2500 chars)
- [ ] Red Flags: 3-5 alertas críticos (600-1200 chars)
- [ ] Referências científicas mencionadas
- [ ] Linguagem técnica mas acessível
- [ ] Foco em medicina funcional

### Qualidade Clínica
- [ ] Acurácia científica verificada
- [ ] Intervenções práticas e seguras
- [ ] Red flags apropriados
- [ ] Tom e linguagem adequados
- [ ] Revisão médica realizada (ou agendada)
- [ ] Feedback incorporado

### Documentação
- [ ] SOCIAL-BATCH-REPORT.json salvo
- [ ] Logs de execução arquivados
- [ ] Documentação atualizada (se necessário)
- [ ] Próximos passos definidos

---

## 10. Ações Pós-Validação

### Se Tudo OK (100% Aprovado)
1. ✅ **Commit**: Commitar mudanças (se aplicável)
2. ✅ **Deploy**: Preparar para produção
3. ✅ **Comunicação**: Notificar equipe clínica
4. ✅ **Treinamento**: Agendar sessão de overview do conteúdo SOCIAL
5. ✅ **Monitoramento**: Acompanhar uso nos primeiros dias

### Se Falhas Detectadas (<100% Sucesso)
1. ⚠️ **Identificar**: Quais items falharam?
2. ⚠️ **Diagnosticar**: Por que falharam? (API, JSON, conteúdo?)
3. ⚠️ **Corrigir**: Re-executar apenas items falhados
4. ⚠️ **Re-validar**: Repetir checklist
5. ⚠️ **Documentar**: Adicionar lições aprendidas

### Se Problemas de Qualidade
1. 🔧 **Revisar**: Identificar padrões de problemas
2. 🔧 **Ajustar**: Melhorar prompts Claude
3. 🔧 **Re-gerar**: Re-executar items problemáticos
4. 🔧 **Validar**: Verificar melhoria
5. 🔧 **Iterar**: Até atingir qualidade desejada

---

## Conclusão

Esta validação garante que o enriquecimento dos 30 items SOCIAL foi:
- ✅ **Tecnicamente bem-sucedido** (todos processados, sem erros)
- ✅ **Cientificamente acurado** (baseado em evidências)
- ✅ **Clinicamente útil** (acionável para médicos)
- ✅ **Alinhado com medicina funcional** (foco em causas raiz)

**Próximo passo**: Se validação 100% OK, marcar batch SOCIAL como COMPLETO e avançar para próximo grupo.

---

**Data de criação**: 2026-01-27
**Última atualização**: 2026-01-27
**Status**: Aguardando execução para validação
