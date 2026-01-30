# Como Usar o Sistema de Enriquecimento: TC Tórax - Nódulo Pulmonar

Este guia explica como executar o enriquecimento clínico do item **TC Tórax - Maior Nódulo Sólido** e adaptar para outros itens.

---

## Pré-requisitos

### Software Necessário

```bash
# Python 3.8+
python3 --version

# PostgreSQL (via Docker)
docker compose ps | grep db

# Bibliotecas Python
pip install psycopg2-binary anthropic
```

### Acesso ao Banco de Dados

```bash
# Credenciais (já configuradas no docker-compose.yml)
Host: localhost
Port: 5432
Database: plenya_db
User: plenya_user
Password: plenya_dev_password
```

### API Anthropic (Opcional)

Para usar o script com busca real de artigos:

```bash
export ANTHROPIC_API_KEY='sua-chave-aqui'
```

---

## Método 1: Execução Mock (Recomendado para Testes)

Usa dados pré-gerados baseados em literatura científica real.

### Vantagens

- Não requer API key
- Execução rápida (~3 segundos)
- Conteúdo já validado
- Ideal para desenvolvimento e testes

### Execução

```bash
# A partir do diretório raiz do projeto
cd /home/user/plenya

# Executar script
python3 scripts/enrich_tc_torax_mock.py
```

### Saída Esperada

```
======================================================================
🔬 ENRIQUECIMENTO MOCK: TC Tórax - Maior Nódulo Sólido
📋 ID: dd6e920c-b203-4d40-b230-55f2074ac613
======================================================================
⚠️  Usando dados pré-gerados (sem API Anthropic)
======================================================================
✅ Conectado ao banco de dados

📊 ETAPA 1: Buscando item no banco...
   Item: TC Tórax - Maior Nódulo Sólido
   Subgrupo: d80afd62-a3e6-491d-b3c3-558b25e27fe5
   Conteúdo atual: Vazio

📚 ETAPA 2: Usando artigos científicos pré-selecionados...
   Total: 4 artigos
   - Guidelines for Management of Incidental Pulmonary Nodules...
   - Lung Cancer Screening with Low-Dose Computed Tomography...
   - The British Thoracic Society Guidelines...
   - Risk Stratification of Pulmonary Nodules...

✍️  ETAPA 3: Usando conteúdo clínico pré-gerado...
   Clinical relevance: 957 chars
   Patient explanation: 814 chars
   Conduct: 844 chars

💾 ETAPA 4: Salvando artigos no banco...
   ✅ Artigo salvo: Guidelines for Management...
   ✅ Artigo salvo: Lung Cancer Screening...
   ✅ Artigo salvo: The British Thoracic Society...
   ✅ Artigo salvo: Risk Stratification...
   Total: 4 artigos processados

🔗 ETAPA 5: Vinculando artigos ao item...
   ✅ 4 artigos vinculados ao item

💾 ETAPA 6: Atualizando conteúdo do item...
✅ Item atualizado no banco

======================================================================
✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!
======================================================================
📋 Item ID: dd6e920c-b203-4d40-b230-55f2074ac613
📚 Artigos vinculados: 4
📝 Clinical relevance: 957 chars
👤 Patient explanation: 814 chars
🏥 Conduct: 844 chars
📅 Last review: 2026-01-28 16:08:59
======================================================================
```

---

## Método 2: Execução com API Real

Busca artigos científicos em tempo real e gera conteúdo via Claude.

### Vantagens

- Conteúdo sempre atualizado
- Busca artigos mais recentes
- Flexível para qualquer item
- Maior controle sobre o conteúdo

### Desvantagens

- Requer API key da Anthropic
- Mais lento (~30-60 segundos)
- Consome créditos da API
- Pode variar entre execuções

### Execução

```bash
# Definir API key
export ANTHROPIC_API_KEY='sk-ant-api03-...'

# Executar script
python3 scripts/enrich_tc_torax_nodulo.py
```

### Customização

Edite o arquivo `scripts/enrich_tc_torax_nodulo.py` para ajustar:

```python
# Modificar queries de busca
search_queries = [
    "pulmonary nodule management Fleischner guidelines 2023",
    "solitary pulmonary nodule malignancy risk",
    "lung cancer screening CT nodule follow-up"
]

# Adicionar mais queries conforme necessário
```

---

## Método 3: Execução via Docker

Wrapper que instala dependências automaticamente.

### Execução

```bash
# Definir API key
export ANTHROPIC_API_KEY='sua-chave-aqui'

# Tornar script executável
chmod +x scripts/run_enrich_tc_torax.sh

# Executar
./scripts/run_enrich_tc_torax.sh
```

---

## Verificação dos Resultados

### Verificar no Banco de Dados

```bash
# Verificar conteúdo do item
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_len,
  LENGTH(patient_explanation) as patient_len,
  LENGTH(conduct) as conduct_len,
  last_review
FROM score_items
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';
"

# Verificar artigos vinculados
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT a.title, a.journal, a.publish_date
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'dd6e920c-b203-4d40-b230-55f2074ac613'
ORDER BY a.publish_date DESC;
"
```

### Ver Conteúdo Completo

```bash
# Clinical relevance
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT clinical_relevance
FROM score_items
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';
"

# Patient explanation
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT patient_explanation
FROM score_items
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';
"

# Conduct
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT conduct
FROM score_items
WHERE id = 'dd6e920c-b203-4d40-b230-55f2074ac613';
"
```

---

## Adaptar para Outros Itens

### Passo 1: Identificar o Item

```bash
# Buscar itens de TC Tórax sem enriquecimento
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name, clinical_relevance IS NULL as needs_enrichment
FROM score_items
WHERE name LIKE 'TC Tórax%'
AND clinical_relevance IS NULL
ORDER BY name;
"
```

### Passo 2: Copiar e Adaptar o Script

```bash
# Copiar script mock
cp scripts/enrich_tc_torax_mock.py scripts/enrich_outro_item.py

# Editar constantes
vim scripts/enrich_outro_item.py
```

Modificar:

```python
# Mudar ID do item
ITEM_ID = "novo-id-aqui"

# Adaptar conteúdo clínico em MOCK_CONTENT
MOCK_CONTENT = {
    "clinical_relevance": "...",
    "patient_explanation": "...",
    "conduct": "...",
    "articles": [...]
}
```

### Passo 3: Executar

```bash
python3 scripts/enrich_outro_item.py
```

---

## Enriquecimento em Lote (Batch)

Para enriquecer múltiplos itens de uma vez:

### Criar Script de Batch

```python
#!/usr/bin/env python3
# scripts/enrich_tc_torax_batch.py

import sys
from enrich_tc_torax_mock import enrich_item  # Função extraída

# Lista de IDs de itens
ITEMS = [
    'dd6e920c-b203-4d40-b230-55f2074ac613',  # Maior Nódulo Sólido
    'outro-id-aqui',  # Número de Nódulos
    'mais-um-id',  # Nódulo em Vidro Fosco
    # ...
]

def main():
    success = 0
    failed = 0

    for item_id in ITEMS:
        print(f"\n{'='*70}")
        print(f"Enriquecendo item: {item_id}")
        print('='*70)

        try:
            if enrich_item(item_id):
                success += 1
            else:
                failed += 1
        except Exception as e:
            print(f"ERRO: {str(e)}")
            failed += 1

    print(f"\n{'='*70}")
    print(f"RESUMO: {success} sucesso, {failed} falhas")
    print('='*70)

if __name__ == "__main__":
    sys.exit(main())
```

### Executar Batch

```bash
python3 scripts/enrich_tc_torax_batch.py
```

---

## Solução de Problemas

### Erro: Item não encontrado

```
❌ Item não encontrado: dd6e920c-b203-4d40-b230-55f2074ac613
```

**Solução:** Verificar se o ID está correto no banco:

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT id, name FROM score_items WHERE name LIKE '%Nódulo%';
"
```

### Erro: Conexão recusada

```
❌ Erro ao conectar ao banco: Connection refused
```

**Solução:** Verificar se o Docker está rodando:

```bash
docker compose ps
docker compose up -d db
```

### Erro: Artigo duplicado

```
❌ Erro ao salvar artigo: duplicate key value violates unique constraint
```

**Solução:** Normal! O script detecta artigos existentes e reutiliza. Não é um erro real.

### Aviso: Item já possui conteúdo

```
⚠️  ATENÇÃO: Item já possui conteúdo!
   Deseja sobrescrever? (s/N):
```

**Solução:** Responda `s` para sobrescrever ou `N` para cancelar.

---

## Melhores Práticas

### 1. Sempre Verificar Antes de Enriquecer

```bash
# Ver conteúdo atual
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT name, clinical_relevance IS NOT NULL as has_content
FROM score_items
WHERE id = 'item-id-aqui';
"
```

### 2. Backup Antes de Modificações em Massa

```bash
# Backup da tabela
docker compose exec db pg_dump -U plenya_user -d plenya_db -t score_items > backup_score_items.sql
```

### 3. Testar com Mock Primeiro

Sempre teste com o script mock antes de usar API real:

1. Execute `enrich_tc_torax_mock.py`
2. Verifique resultados no banco
3. Se OK, adapte para API real

### 4. Documentar Mudanças

Após cada enriquecimento, documente:

- Item enriquecido
- Data/hora
- Artigos adicionados
- Fonte do conteúdo (mock ou API)

---

## Arquivos de Referência

### Scripts

- `/home/user/plenya/scripts/enrich_tc_torax_mock.py` - Mock com dados pré-gerados
- `/home/user/plenya/scripts/enrich_tc_torax_nodulo.py` - API real
- `/home/user/plenya/scripts/run_enrich_tc_torax.sh` - Wrapper Docker

### Documentação

- `/home/user/plenya/TC-TORAX-NODULO-ENRICHMENT-REPORT.md` - Relatório detalhado
- `/home/user/plenya/TC-TORAX-NODULO-SUMMARY.json` - Sumário JSON
- `/home/user/plenya/COMO-USAR-ENRIQUECIMENTO-TC-TORAX.md` - Este guia

---

## Contato e Suporte

Para dúvidas ou problemas:

1. Verificar logs de execução
2. Consultar documentação técnica
3. Revisar código dos scripts
4. Testar com dados mock primeiro

---

**Última atualização:** 2026-01-28
**Versão:** 1.0
**Autor:** Sistema de Enriquecimento Clínico Plenya
