# Batch ALIMENTAÇÃO Parte 2 - Executive Summary

## Status Atual
**Sistema pronto para execução** - Aguardando configuração da API key Anthropic

## O Que Foi Criado

### 1. Script de Enriquecimento Principal
**Arquivo**: `/home/user/plenya/scripts/batch_alimentacao_parte2.py`

**Características**:
- Processa 20 items do grupo ALIMENTAÇÃO
- Utiliza Claude Opus 4.5 (modelo mais avançado)
- Extended thinking habilitado (10k tokens)
- Rate limiting automático (2s entre requests)
- Tratamento robusto de erros
- Geração de JSON e SQL

### 2. Script de Execução Facilitado
**Arquivo**: `/home/user/plenya/scripts/execute_batch_alimentacao_p2.sh`

**Funcionalidades**:
- Validação automática de dependências
- Checagem de API key
- Output formatado e informativo
- Instruções de próximos passos

### 3. Documentação Completa
**Arquivo**: `/home/user/plenya/BATCH-ALIMENTACAO-PARTE2-README.md`

**Conteúdo**:
- Lista completa dos 20 items
- Instruções de configuração
- Estrutura detalhada do enriquecimento
- Estimativas de tempo e custo
- Comandos de validação

### 4. Exemplo de Output Esperado
**Arquivo**: `/home/user/plenya/BATCH-ALIMENTACAO-PARTE2-EXAMPLE.json`

**Demonstra**:
- Qualidade do conteúdo clínico
- Estrutura completa dos dados
- Nível de detalhamento esperado

## Items a Processar (20)

### Grupo 1: Histórico e Contexto Familiar (6 items)
1. Qualidade alimentação parentes
2. Infância
3. Introdução alimentar
4. Alimentação parental pré/durante gestação
5. Perfil metabólico parental
6. Piores períodos

### Grupo 2: Intolerâncias e Restrições (4 items)
7. Intolerâncias
8. Lactose
9. Proteína do leite
10. Restrições pessoais

### Grupo 3: Padrões e Comportamentos (7 items)
11. Recordatório (24h)
12. Onde e como come
13. Ordem alimentos
14. Quanto come
15. Quem cozinha
16. Regras alimentares
17. Preferências

### Grupo 4: Aspectos Emocionais e Satisfação (3 items)
18. Relação com comida
19. Satisfação pós-refeição
20. Outros

## Como Executar

### Passo 1: Configurar API Key

**Opção A - Variável de Ambiente**:
```bash
export ANTHROPIC_API_KEY='sk-ant-api03-...'
```

**Opção B - Arquivo de Configuração**:
```bash
echo 'sk-ant-api03-...' > ~/.anthropic_key
```

### Passo 2: Instalar Dependência (se necessário)

```bash
pip install anthropic
```

### Passo 3: Executar

**Método Simplificado**:
```bash
./scripts/execute_batch_alimentacao_p2.sh
```

**Método Direto**:
```bash
python3 scripts/batch_alimentacao_parte2.py
```

### Passo 4: Revisar Outputs

```bash
# Ver resultados JSON
cat batch_alimentacao_parte2_results.json | jq '.[] | {id, title}'

# Ver SQL gerado
head -n 50 batch_alimentacao_parte2.sql
```

### Passo 5: Aplicar no Banco

```bash
# Copiar SQL para container
docker compose cp batch_alimentacao_parte2.sql db:/tmp/

# Executar no PostgreSQL
docker compose exec db psql -U plenya_user -d plenya_db -f /tmp/batch_alimentacao_parte2.sql
```

### Passo 6: Validar

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  title,
  LENGTH(clinical_relevance) as rel_len,
  ARRAY_LENGTH(health_implications, 1) as impl_count,
  ARRAY_LENGTH(recommendations, 1) as rec_count
FROM score_items
WHERE group_name = 'Alimentação'
  AND clinical_relevance IS NOT NULL
ORDER BY title
LIMIT 10;
"
```

## Estimativas

### Tempo de Execução
- **Por item**: 20-30 segundos
- **Total**: 10-12 minutos
- **Com rate limiting**: +40 segundos

**Total estimado**: ~12-15 minutos

### Custo (Claude Opus 4.5)
- **Input tokens**: ~5k por item × 20 = 100k tokens
- **Thinking tokens**: ~6k por item × 20 = 120k tokens
- **Output tokens**: ~2.5k por item × 20 = 50k tokens

**Custos**:
- Input: 100k × $15/1M = $1.50
- Thinking: 120k × $15/1M = $1.80
- Output: 50k × $75/1M = $3.75

**Total estimado**: ~$7 USD

### Token Budget
- **Disponível na sessão**: 200k tokens
- **Necessário**: ~20-25k tokens (orquestração)
- **Margem**: Confortável

## Estrutura do Enriquecimento

Cada item receberá 8 campos completos:

1. **question** - Pergunta clínica melhorada
2. **clinical_relevance** - Relevância clínica (2-3 sentenças)
3. **interpretation_guide** - Guia de interpretação (3-4 sentenças)
4. **health_implications** - 4-6 implicações de saúde
5. **followup_questions** - 3-5 perguntas de follow-up
6. **red_flags** - 3-4 sinais de alerta
7. **recommendations** - 4-6 recomendações baseadas em evidências
8. **scientific_background** - Contexto científico (2-3 sentenças)

## Qualidade do Conteúdo

### Características Garantidas
- Todo conteúdo em Português (PT-BR)
- Baseado em evidências científicas atuais
- Terminologia médica precisa mas acessível
- Foco em medicina funcional e preventiva
- Orientações práticas e acionáveis
- Contexto fisiopatológico quando relevante

### Validação de Qualidade
- Revisão manual recomendada antes de aplicar SQL
- Verificação de consistência clínica
- Checagem de completude dos campos
- Validação de encoding (UTF-8)

## Outputs Gerados

### 1. batch_alimentacao_parte2_results.json
Estrutura completa:
```json
[
  {
    "id": "uuid",
    "title": "Nome do item",
    "enrichment": {
      "question": "...",
      "clinical_relevance": "...",
      ...
    },
    "tokens": {
      "thinking": 6543,
      "output": 2891
    }
  }
]
```

### 2. batch_alimentacao_parte2.sql
Estrutura:
```sql
-- BATCH ALIMENTAÇÃO PARTE 2 - 20 ITEMS
-- Generated: 2026-01-27 ...

-- Item: Intolerâncias
UPDATE score_items SET
  question = '...',
  clinical_relevance = '...',
  interpretation_guide = '...',
  health_implications = ARRAY[...]::text[],
  followup_questions = ARRAY[...]::text[],
  red_flags = ARRAY[...]::text[],
  recommendations = ARRAY[...]::text[],
  scientific_background = '...'
WHERE id = 'fb918f6e-32c6-494f-b7a9-cc6ee738ce32';

-- [Repetido para os 20 items]
```

## Troubleshooting

### Problema: API Key não encontrada
**Solução**:
```bash
export ANTHROPIC_API_KEY='sk-ant-...'
# ou
echo 'sk-ant-...' > ~/.anthropic_key
```

### Problema: anthropic package não instalado
**Solução**:
```bash
pip install anthropic
# ou
pip3 install anthropic
```

### Problema: Rate limit exceeded
**Solução**: O script já inclui delay de 2s entre requests. Se persistir, aumentar o delay no código.

### Problema: JSON parse error
**Solução**: Revisar manualmente o output do Claude e corrigir formatação.

### Problema: SQL escape error
**Solução**: O script já faz escape de aspas simples. Verificar caracteres especiais no output.

## Próximos Passos

### Após Esta Execução
1. Processar items restantes do grupo ALIMENTAÇÃO (se houver)
2. Revisar qualidade do conteúdo gerado
3. Testar no frontend com dados reais
4. Coletar feedback clínico
5. Iterar melhorias se necessário

### Outros Grupos para Enriquecer
- **Movimento**: ~15-20 items
- **Sono**: ~15-20 items
- **Stress**: ~10-15 items
- **Social**: ~10-15 items
- **Propósito**: ~5-10 items

## Checklist de Execução

- [ ] Configurar ANTHROPIC_API_KEY
- [ ] Instalar dependência anthropic
- [ ] Executar script de enriquecimento
- [ ] Aguardar conclusão (~12 minutos)
- [ ] Revisar JSON gerado
- [ ] Revisar SQL gerado
- [ ] Aplicar SQL no banco de dados
- [ ] Validar dados no PostgreSQL
- [ ] Testar no frontend
- [ ] Documentar resultados

## Contato e Suporte

Para dúvidas ou problemas:
1. Revisar logs de execução
2. Verificar arquivos de documentação
3. Consultar exemplo de output esperado
4. Verificar status da API Anthropic

---

**Criado**: 2026-01-27
**Versão**: 1.0
**Status**: Pronto para execução
**Tempo estimado**: 12-15 minutos
**Custo estimado**: ~$7 USD
