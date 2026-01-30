# Batch ALIMENTAÇÃO Parte 2 - 20 Items

## Status
**Script criado**: `scripts/batch_alimentacao_parte2.py`
**Status**: Pronto para execução (requer configuração da API key)

## Items a Processar (20)

1. **Qualidade alimentação parentes** - `b781e715-372d-4f05-b01e-4db68c05d8db`
2. **Infância** - `2aa102b7-9d48-4870-a1f1-53f53b29a3e6`
3. **Intolerâncias** - `fb918f6e-32c6-494f-b7a9-cc6ee738ce32`
4. **Introdução alimentar** - `39d2c6c6-7b8b-41a8-a90e-05b396ad61dc`
5. **Lactose** - `7ba24074-1ffa-444d-82d7-0decb8db039d`
6. **Alimentação parental pré/durante gestação** - `ce6181cc-827b-4ad8-89ce-a555884c71c6`
7. **Recordatório** - `d8696390-b619-4752-b366-5c63f4730811`
8. **Onde e como come** - `71cc4a7b-5ebd-437a-bbee-6f85105ac8dc`
9. **Ordem alimentos** - `5dc99862-e9c2-4ddb-b0b3-0631a32e3dc4`
10. **Outros** - `3dcaa794-d1f1-4551-ad62-8bae902abaf8`
11. **Perfil metabólico parental** - `f7a0aa7a-cfea-4c61-8391-52d82b236429`
12. **Piores períodos** - `5cac2737-c0c5-47d9-aaa8-2ca4fad39f47`
13. **Preferências** - `c3b9b8b0-2e94-40f8-b3d6-143a7d8224f1`
14. **Proteína do leite** - `a7fa71ab-72c6-430d-b6f6-ca6c95611197`
15. **Quanto come** - `d8de21dc-24bd-489c-93ed-aa094428f292`
16. **Quem cozinha** - `a35a3012-60ea-4659-a742-fbbcb741c6db`
17. **Regras alimentares** - `27582ef6-b5e1-4d6f-9a7a-368c27fdf42a`
18. **Relação com comida** - `6998408b-341d-487c-9c99-7e8bc72de3fb`
19. **Restrições pessoais** - `8f24a93b-8cd7-4157-adf4-3c9a7985e368`
20. **Satisfação pós-refeição** - `01498f3d-551a-4285-bc84-86d015569d31`

## Configuração da API Key

### Opção 1: Variável de Ambiente (Recomendado)

```bash
export ANTHROPIC_API_KEY='sk-ant-...'
python3 scripts/batch_alimentacao_parte2.py
```

### Opção 2: Arquivo de Configuração

```bash
echo 'sk-ant-...' > ~/.anthropic_key
python3 scripts/batch_alimentacao_parte2.py
```

## Estrutura do Enriquecimento

Cada item será enriquecido com:

### 1. Question (String)
Pergunta clínica melhorada:
- Clara e específica
- Amigável ao paciente
- Inclui timeframes quando relevante

### 2. Clinical Relevance (String)
Por que isso importa clinicamente (2-3 sentenças):
- Mecanismos fisiopatológicos específicos
- Impacto na saúde metabólica, longevidade, prevenção de doenças

### 3. Interpretation Guide (String)
Como interpretar as respostas (3-4 sentenças):
- O que constitui padrões ótimos, subótimos e preocupantes
- Red flags ou indicadores positivos chave
- Fatores contextuais a considerar

### 4. Health Implications (Array)
4-6 implicações de saúde específicas:
- Concretas e baseadas em evidências
- Cobrem efeitos imediatos e de longo prazo
- Incluem impactos metabólicos, inflamatórios e sistêmicos

### 5. Followup Questions (Array)
3-5 perguntas de follow-up contextuais:
- Aprofundam compreensão baseada na resposta inicial
- Ajudam a quantificar ou qualificar a informação
- Exploram fatores relacionados

### 6. Red Flags (Array)
3-4 padrões preocupantes que requerem atenção:
- Sinais clínicos de alerta claros
- Padrões sugerindo necessidade de intervenção

### 7. Recommendations (Array)
4-6 recomendações baseadas em evidências:
- Orientações acionáveis e específicas
- Aplicáveis a diferentes cenários/tipos de pacientes
- Baseadas na ciência nutricional atual

### 8. Scientific Background (String)
Contexto científico breve (2-3 sentenças):
- Principais achados de pesquisa
- Mecanismos relevantes
- Consenso científico atual

## Modelo Utilizado

**Claude Opus 4.5** (`claude-opus-4-5-20251101`)
- Modelo mais avançado da Anthropic
- Extended thinking habilitado (10.000 tokens)
- Temperature: 1.0 (criatividade balanceada)
- Max output: 16.000 tokens

## Outputs Gerados

### 1. JSON Completo
**Arquivo**: `batch_alimentacao_parte2_results.json`

Contém:
- Todos os 20 items processados
- Conteúdo clínico completo
- Métricas de tokens utilizados
- Timestamp de processamento

### 2. SQL Migration
**Arquivo**: `batch_alimentacao_parte2.sql`

Contém:
- 20 statements UPDATE
- Pronto para execução no PostgreSQL
- Escape de caracteres especiais
- Comentários identificando cada item

### 3. Relatório de Execução
Console output com:
- Progress em tempo real
- Contagem de tokens
- Estatísticas de sucesso/falha
- Tempo de execução

## Tempo Estimado

- **Por item**: ~20-30 segundos (com extended thinking)
- **Total**: ~10-12 minutos
- **Rate limiting**: 2 segundos entre items

## Custo Estimado

Baseado em Claude Opus 4.5 pricing:
- Input: $15 / 1M tokens
- Output: $75 / 1M tokens

Estimativa para 20 items:
- Thinking tokens: ~120k-150k
- Output tokens: ~40k-60k
- **Custo total**: ~$5-7 USD

## Como Executar

```bash
# 1. Configurar API key
export ANTHROPIC_API_KEY='sk-ant-...'

# 2. Executar script
python3 scripts/batch_alimentacao_parte2.py

# 3. Aguardar conclusão (~10 minutos)

# 4. Revisar outputs
cat batch_alimentacao_parte2_results.json
cat batch_alimentacao_parte2.sql

# 5. Aplicar no banco (após revisão)
docker compose exec db psql -U plenya_user -d plenya_db -f /path/to/batch_alimentacao_parte2.sql
```

## Validação Pós-Execução

```bash
# Verificar items atualizados
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  id,
  title,
  LENGTH(clinical_relevance) as relevance_len,
  ARRAY_LENGTH(health_implications, 1) as implications_count,
  ARRAY_LENGTH(recommendations, 1) as recommendations_count
FROM score_items
WHERE id IN (
  'b781e715-372d-4f05-b01e-4db68c05d8db',
  '2aa102b7-9d48-4870-a1f1-53f53b29a3e6',
  -- ... outros IDs
)
ORDER BY title;
"
```

## Próximos Passos

Após esta execução:
1. Revisar conteúdo gerado
2. Validar qualidade clínica
3. Aplicar SQL no banco
4. Testar no frontend
5. Documentar items restantes do grupo ALIMENTAÇÃO

## Notas Importantes

- **Todos os campos em Português**: Conteúdo clínico em PT-BR
- **Baseado em Evidências**: Referências científicas atuais
- **Orientação Clínica**: Foco em medicina funcional e preventiva
- **Patient-Friendly**: Linguagem acessível mas tecnicamente precisa

---

**Criado**: 2026-01-27
**Grupo**: ALIMENTAÇÃO - Parte 2
**Total Items**: 20
**Status**: Aguardando configuração de API key
