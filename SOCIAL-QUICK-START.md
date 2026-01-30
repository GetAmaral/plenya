# Quick Start Guide - Batch SOCIAL Enrichment

Guia rápido para executar o enriquecimento de 30 items SOCIAL.

---

## Pré-requisitos

### 1. API Rodando
```bash
docker compose up -d
```

Verifique se está respondendo:
```bash
curl http://localhost:3001/health
```

### 2. ANTHROPIC_API_KEY
```bash
export ANTHROPIC_API_KEY='sk-ant-api03-...'
```

Verifique se foi exportada:
```bash
echo $ANTHROPIC_API_KEY
```

### 3. Dependências Python
```bash
pip install anthropic requests
```

---

## Execução em 3 Passos

### Passo 1: Tornar script executável
```bash
chmod +x /home/user/plenya/execute_social_batch.sh
```

### Passo 2: Executar enriquecimento
```bash
/home/user/plenya/execute_social_batch.sh
```

### Passo 3: Revisar relatório
```bash
cat /home/user/plenya/SOCIAL-BATCH-REPORT.json
```

---

## O Que Esperar

### Durante a Execução
- Processamento de 30 items sequencialmente
- Tempo estimado: 20-25 minutos
- Progresso exibido no terminal: `[1/30] Processing item...`

### Output por Item
```
================================================================================
[1/30] Processing item c84412f7-393f-41d0-8bd7-0a28824dbeb0
================================================================================
✓ Item: Ambiente Sonoro
  Group: SOCIAL
  Description: Avalia exposição a ruído ambiental...
🤖 Gerando conteúdo clínico com Claude...
✅ Item atualizado com sucesso!
   Clinical Relevance: 1247 chars
   Interpretation: 1589 chars
   Actionable Insights: 1834 chars
   Red Flags: 892 chars
```

### Relatório Final
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

---

## Verificação de Sucesso

### No Terminal
```
==============================================================================
RELATÓRIO FINAL
==============================================================================
✅ Sucesso: 30/30
❌ Falhas: 0/30

==============================================================================
```

### No Banco de Dados
```sql
-- Verificar quantos items foram enriquecidos
SELECT
    COUNT(*) as total_enriquecidos
FROM score_items
WHERE
    group_name = 'SOCIAL'
    AND clinical_relevance IS NOT NULL
    AND interpretation_guidelines IS NOT NULL
    AND actionable_insights IS NOT NULL
    AND red_flags IS NOT NULL;

-- Esperado: 30
```

### Via API
```bash
TOKEN=$(curl -s -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"import@plenya.com","password":"Import@123456"}' | \
  python3 -c "import sys, json; print(json.load(sys.stdin)['accessToken'])")

curl -s -X GET "http://localhost:3001/api/v1/score-items/c84412f7-393f-41d0-8bd7-0a28824dbeb0" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool
```

---

## Troubleshooting

### Erro: "ANTHROPIC_API_KEY not found"
**Solução**:
```bash
export ANTHROPIC_API_KEY='sua-chave-aqui'
./execute_social_batch.sh
```

### Erro: "API não está respondendo"
**Solução**:
```bash
docker compose up -d
# Aguarde 10s para API inicializar
curl http://localhost:3001/health
./execute_social_batch.sh
```

### Erro: "Module 'anthropic' not found"
**Solução**:
```bash
pip install anthropic requests
./execute_social_batch.sh
```

### Erro: "401 Unauthorized"
**Causa**: Token de login expirado ou credenciais inválidas

**Solução**:
```bash
# Verificar credenciais no script
cat scripts/batch_social_enrichment.py | grep -A 2 "LOGIN"
# Deve mostrar:
# LOGIN_EMAIL = "import@plenya.com"
# LOGIN_PASSWORD = "Import@123456"
```

### Erro: "JSON decode error"
**Causa**: Claude retornou resposta malformada

**Solução**: O script tenta automaticamente remover markdown. Se persistir, abra issue com o item ID que falhou.

---

## Análise Pós-Execução

### 1. Estatísticas de Conteúdo
```bash
python3 << 'EOF'
import requests
import json

# Login
response = requests.post(
    "http://localhost:3001/api/v1/auth/login",
    json={"email": "import@plenya.com", "password": "Import@123456"}
)
token = response.json()["accessToken"]

# Buscar items SOCIAL
headers = {"Authorization": f"Bearer {token}"}
response = requests.get(
    "http://localhost:3001/api/v1/score-items?group=SOCIAL&limit=100",
    headers=headers
)
items = response.json()["items"]

# Estatísticas
enriched = [i for i in items if i.get("clinical_relevance")]
print(f"Items enriquecidos: {len(enriched)}/{len(items)}")

if enriched:
    avg_clinical = sum(len(i["clinical_relevance"]) for i in enriched) / len(enriched)
    avg_interpretation = sum(len(i["interpretation_guidelines"]) for i in enriched) / len(enriched)
    avg_actionable = sum(len(i["actionable_insights"]) for i in enriched) / len(enriched)
    avg_red_flags = sum(len(i["red_flags"]) for i in enriched) / len(enriched)

    print(f"\nMédia de caracteres:")
    print(f"  Clinical Relevance: {avg_clinical:.0f}")
    print(f"  Interpretation: {avg_interpretation:.0f}")
    print(f"  Actionable Insights: {avg_actionable:.0f}")
    print(f"  Red Flags: {avg_red_flags:.0f}")
EOF
```

### 2. Exemplo de Item Enriquecido
```bash
python3 << 'EOF'
import requests
import json

# Login
response = requests.post(
    "http://localhost:3001/api/v1/auth/login",
    json={"email": "import@plenya.com", "password": "Import@123456"}
)
token = response.json()["accessToken"]

# Buscar primeiro item SOCIAL
headers = {"Authorization": f"Bearer {token}"}
response = requests.get(
    "http://localhost:3001/api/v1/score-items/c84412f7-393f-41d0-8bd7-0a28824dbeb0",
    headers=headers
)
item = response.json()

print(f"Item: {item['name']}")
print(f"\nClinical Relevance:")
print(item.get('clinical_relevance', 'N/A')[:200] + "...")
print(f"\nInterpretation Guidelines:")
print(item.get('interpretation_guidelines', 'N/A')[:200] + "...")
EOF
```

---

## Próximos Passos

### 1. Revisão Médica
- [ ] Validar conteúdo clínico com especialista
- [ ] Verificar acurácia de mecanismos fisiopatológicos
- [ ] Confirmar intervenções práticas

### 2. Teste no Frontend
- [ ] Acessar interface web
- [ ] Navegar para items SOCIAL
- [ ] Verificar exibição dos campos enriquecidos
- [ ] Testar responsividade

### 3. Documentação
- [ ] Atualizar guia do médico com insights SOCIAL
- [ ] Criar exemplos de casos clínicos
- [ ] Documentar fluxo de avaliação

### 4. Iteração
- [ ] Coletar feedback de médicos
- [ ] Ajustar conteúdo baseado em uso real
- [ ] Adicionar mais referências científicas

---

## Comandos Úteis

### Listar todos os items SOCIAL
```bash
TOKEN=$(curl -s -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"import@plenya.com","password":"Import@123456"}' | \
  python3 -c "import sys, json; print(json.load(sys.stdin)['accessToken'])")

curl -s "http://localhost:3001/api/v1/score-items?group=SOCIAL&limit=100" \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool
```

### Re-executar apenas items que falharam
```bash
# Editar script para usar apenas IDs de results["failed"]
nano scripts/batch_social_enrichment.py
# Modificar SOCIAL_ITEM_IDS = results["failed"]
./execute_social_batch.sh
```

### Backup do banco antes de executar
```bash
docker compose exec db pg_dump -U plenya_user plenya_db > backup_before_social.sql
```

### Restaurar backup se necessário
```bash
docker compose exec -T db psql -U plenya_user plenya_db < backup_before_social.sql
```

---

## Suporte

**Documentação completa**:
- Metodologia: `SOCIAL-ENRICHMENT-METHODOLOGY.md`
- Referências: `SOCIAL-SCIENTIFIC-REFERENCES.md`
- Arquitetura: `CLAUDE.md` e `ARQUITETURA.md`

**Logs**:
- Relatório: `SOCIAL-BATCH-REPORT.json`
- Terminal output (redirecionar para arquivo):
  ```bash
  ./execute_social_batch.sh 2>&1 | tee social_execution.log
  ```

**Contato**:
- Issues GitHub: Para bugs técnicos
- Pull Requests: Para melhorias no conteúdo clínico

---

**Última atualização**: 2026-01-27
