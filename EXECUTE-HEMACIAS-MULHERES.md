# Enriquecimento: Hemácias - Mulheres

## Item
- **Nome:** Hemácias - Mulheres
- **ID:** `501fd84a-a440-4c13-9b11-35e2f69017d1`
- **Grupo:** Exames > Laboratoriais

## Contexto Clínico
- Contagem de eritrócitos em mulheres: **4.0-5.2 milhões/µL**
- Valores menores que homens (4.5-5.5) devido a:
  - Menstruação (perda de ferro)
  - Níveis hormonais diferentes
  - Menor massa muscular
- Alterações importantes:
  - **Baixo (<4.0):** Anemia (ferro, B12, folato)
  - **Alto (>5.2):** Policitemia, desidratação

## Como Executar

### Opção 1: Via Docker (Recomendado)

```bash
# 1. Configurar sua API key da Anthropic
export ANTHROPIC_API_KEY="sk-ant-..."

# 2. Executar via Docker
docker compose exec \
  -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" \
  -e NODE_PATH=/app/apps/web/node_modules \
  web node /app/apps/web/enrich.mjs
```

### Opção 2: Diretamente com Node.js

```bash
# 1. Instalar dependências (se necessário)
cd /home/user/plenya/apps/web
npm install @anthropic-ai/sdk pg

# 2. Configurar API key
export ANTHROPIC_API_KEY="sk-ant-..."

# 3. Executar
node /home/user/plenya/scripts/enrich_hemacias_mulheres.mjs
```

### Opção 3: Script Shell Simplificado

```bash
# Criar script wrapper
cat > /tmp/run_enrich.sh << 'EOF'
#!/bin/bash
read -sp "Cole sua ANTHROPIC_API_KEY: " API_KEY
echo ""

docker compose exec \
  -e ANTHROPIC_API_KEY="$API_KEY" \
  -e NODE_PATH=/app/apps/web/node_modules \
  web node /app/apps/web/enrich.mjs
EOF

chmod +x /tmp/run_enrich.sh
/tmp/run_enrich.sh
```

## O Que o Script Faz

1. **Busca artigos científicos** via Claude API:
   - Red blood cell count in women
   - Sex differences in hemoglobin
   - Iron deficiency and menstruation
   - Hormonal effects on RBC production

2. **Gera conteúdo clínico em PT-BR**:
   - `clinical_relevance`: Explicação médica detalhada
   - `patient_explanation`: Texto para pacientes
   - `conduct`: Orientações de conduta médica

3. **Salva no banco de dados**:
   - Insere artigos na tabela `articles`
   - Cria relações em `score_item_articles`
   - Atualiza `score_items` (ID: 501fd84a-a440-4c13-9b11-35e2f69017d1)

## Verificação Pós-Execução

```sql
-- Ver item atualizado
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars
FROM score_items
WHERE id = '501fd84a-a440-4c13-9b11-35e2f69017d1';

-- Ver artigos vinculados
SELECT
  a.title,
  a.url,
  a.article_type
FROM score_item_articles sia
JOIN articles a ON a.id = sia.article_id
WHERE sia.score_item_id = '501fd84a-a440-4c13-9b11-35e2f69017d1';
```

## Arquivos Criados

- `/home/user/plenya/scripts/enrich_hemacias_mulheres.mjs` - Script principal
- `/home/user/plenya/apps/web/enrich.mjs` - Cópia no container web
- Este README

## Próximos Passos

Após executar com sucesso:

1. Verificar dados no banco
2. Testar renderização no frontend
3. Validar conteúdo com profissional médico
4. Repetir processo para outros itens laboratoriais

## Troubleshooting

**Erro: API key não configurada**
```bash
export ANTHROPIC_API_KEY="sua-chave-aqui"
```

**Erro: Cannot find package '@anthropic-ai/sdk'**
```bash
docker compose exec web pnpm add @anthropic-ai/sdk pg --filter web
docker compose restart web
```

**Erro: Connection refused ao banco**
```bash
docker compose ps  # Verificar se db está running
docker compose up -d db
```

## Tempo Estimado
- Execução: 2-3 minutos
- Custo API: ~$0.05-0.10 (Claude Sonnet)
