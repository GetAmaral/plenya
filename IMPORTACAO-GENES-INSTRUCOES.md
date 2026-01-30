# Instruções de Importação - Enriquecimento Genético Batch 1

**Arquivo:** `/home/user/plenya/genetic_enrichment_batch1_CONSOLIDATED.json`
**Genes processados:** 9 de 20 (45%)
**Data:** 2026-01-26

---

## Resumo dos Genes

1. FTO rs9939609 (Obesidade) - da4da973-821f-4859-bc4f-c913a8b0b8d3
2. MC4R rs17782313 (Obesidade) - de659fee-1ab3-437f-b971-a3b4e251a599
3. TCF7L2 rs7903146 (Diabetes) - a979d8f3-7e94-4ad9-bc86-d120765eef2f
4. PPARG Pro12Ala rs1801282 (Diabetes) - 944287b2-de63-4b77-8260-551db06da9bb
5. MTHFR C677T rs1801133 (Homocisteína) - 93ffb2ea-0a41-428c-b481-13f55a321e85
6. APOE (Alzheimer e Lipídios) - 160b5986-1045-4ae7-85c7-8dfe058c4df7
7. FADS1 rs174546 (Ômega-3) - b841503a-0d6c-466d-a210-70198aebb7e5
8. FADS2 rs174575 (Ômega-3/DHA) - 23ce1010-b72c-475d-b9c0-8d9585a04bf3
9. ACE I/D rs4646994 (Hipertensão) - d00ac0c9-a38c-4f36-aab6-e796dee450d2

---

## Método 1: Importação via Script Python (Recomendado)

### Criar script de importação

```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Script de importação de enriquecimento genético
Atualiza score_groups com clinical_relevance, patient_explanation e conduct
"""

import json
import requests
import sys

API_URL = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

def login():
    """Faz login e retorna token"""
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
    )
    response.raise_for_status()
    return response.json()["token"]

def update_gene(token, gene_data):
    """Atualiza um gene com textos enriquecidos"""
    headers = {"Authorization": f"Bearer {token}"}

    payload = {
        "clinical_relevance": gene_data["clinical_relevance"],
        "patient_explanation": gene_data["patient_explanation"],
        "conduct": gene_data["conduct"]
    }

    # Atualizar grupo (genes foram importados como grupos)
    response = requests.put(
        f"{API_URL}/score-groups/{gene_data['gene_id']}",
        headers=headers,
        json=payload
    )

    if response.status_code == 200:
        print(f"✅ {gene_data['gene_name']}")
    else:
        print(f"❌ {gene_data['gene_name']}: {response.text}")

    return response.status_code == 200

def main():
    # Ler JSON
    with open('genetic_enrichment_batch1_CONSOLIDATED.json', 'r', encoding='utf-8') as f:
        genes = json.load(f)

    print(f"📂 {len(genes)} genes para importar\n")

    # Login
    print("🔐 Fazendo login...")
    token = login()
    print("✅ Login bem-sucedido\n")

    # Atualizar cada gene
    success_count = 0
    for gene in genes:
        if update_gene(token, gene):
            success_count += 1

    print(f"\n✅ {success_count}/{len(genes)} genes importados com sucesso")

if __name__ == "__main__":
    main()
```

### Executar

```bash
cd /home/user/plenya
pip3 install requests
python3 import_genetic_enrichment.py
```

---

## Método 2: Importação Manual via curl

### Obter token de autenticação

```bash
TOKEN=$(curl -s -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"import@plenya.com","password":"Import@123456"}' \
  | jq -r '.token')

echo "Token: $TOKEN"
```

### Atualizar cada gene (exemplo para FTO)

```bash
curl -X PUT http://localhost:3001/api/v1/score-groups/da4da973-821f-4859-bc4f-c913a8b0b8d3 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d @- <<'EOF'
{
  "clinical_relevance": "O gene FTO (Fat Mass and Obesity-associated) contém o polimorfismo rs9939609, um dos SNPs mais estudados na genética da obesidade...",
  "patient_explanation": "O gene FTO contém uma variação (rs9939609) que influencia como seu corpo regula o peso...",
  "conduct": "**Quando testar:** Considerar teste em pacientes com obesidade familiar..."
}
EOF
```

Repetir para cada um dos 9 genes, substituindo ID e campos.

---

## Método 3: Importação via Interface Web (se disponível)

1. Acessar http://localhost:3000/scores
2. Buscar grupo "Genética" ou genes individuais
3. Editar cada gene manualmente
4. Copiar e colar textos do JSON nos campos apropriados
5. Salvar

---

## Validação Após Importação

### Verificar no banco

```sql
-- Verificar quantos genes têm textos enriquecidos
SELECT
  COUNT(*) as total,
  COUNT(clinical_relevance) as with_clinical,
  COUNT(patient_explanation) as with_patient,
  COUNT(conduct) as with_conduct
FROM score_groups
WHERE name LIKE '%rs%' OR name LIKE '%FTO%' OR name LIKE '%APOE%';

-- Ver amostra
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars
FROM score_groups
WHERE clinical_relevance IS NOT NULL
LIMIT 5;
```

### Verificar via API

```bash
# Ver detalhes de um gene
curl http://localhost:3001/api/v1/score-groups/da4da973-821f-4859-bc4f-c913a8b0b8d3 \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

---

## Estrutura Esperada no Banco

### Tabela score_groups (após ajuste de schema)

Assumindo que campos foram adicionados:

```sql
ALTER TABLE score_groups
ADD COLUMN clinical_relevance TEXT,
ADD COLUMN patient_explanation TEXT,
ADD COLUMN conduct TEXT;
```

**Nota:** Se os campos não existirem, criar migration primeiro:

```go
// apps/api/internal/models/score_group.go
type ScoreGroup struct {
    ID                   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Name                 string    `gorm:"type:varchar(500);not null"`
    Order                int       `gorm:"not null"`
    ClinicalRelevance    *string   `gorm:"type:text"`
    PatientExplanation   *string   `gorm:"type:text"`
    Conduct              *string   `gorm:"type:text"`
    CreatedAt            time.Time
    UpdatedAt            time.Time
    DeletedAt            gorm.DeletedAt `gorm:"index"`
}
```

Depois gerar migration:

```bash
cd apps/api
pnpm generate:migrations
atlas migrate apply --env dev
```

---

## Próximos Passos

### Curto Prazo

1. **Executar importação dos 9 genes**
   - Escolher método (Python, curl ou manual)
   - Validar no banco após importação

2. **Completar Batch 1 (genes 10-20)**
   - 11 genes restantes para processar
   - Tempo estimado: 22-33 horas

3. **Criar interface de visualização**
   - Frontend para exibir textos enriquecidos
   - Componente ScoreGeneDetail.tsx
   - Tabs: Relevância Clínica | Explicação Paciente | Conduta

### Médio Prazo

4. **Batches 2-4 (genes 21-80)**
   - 60 genes restantes
   - Dividir em 3 batches de 20 genes

5. **Vincular Articles**
   - Importar PDFs relevantes quando disponíveis
   - POST /api/v1/articles/{article_id}/score-items

6. **Revisão médica**
   - Validar textos com especialista
   - Ajustar recomendações conforme necessário

### Longo Prazo

7. **Sistema de relatórios**
   - Gerar PDFs personalizados por perfil genético
   - Combinar múltiplos genes em score composto

8. **Painel nutrigenômico**
   - Recomendações dietéticas agregadas
   - Interações gene-nutriente destacadas
   - Suplementação personalizada

---

## Troubleshooting

### Erro: Column does not exist

**Problema:** Campos clinical_relevance, patient_explanation, conduct não existem na tabela

**Solução:** Adicionar migration conforme descrito acima

### Erro: Authentication failed

**Problema:** Token expirado ou inválido

**Solução:** Gerar novo token via login

### Erro: JSON parsing error

**Problema:** Encoding UTF-8 não preservado

**Solução:** Garantir `Content-Type: application/json; charset=utf-8` e `ensure_ascii=False` no Python

### Genes não aparecem na busca

**Problema:** Estrutura de importação criou genes como "grupos" ao invés de items

**Solução:** Trabalhar com score_groups por ora, refatorar hierarquia depois se necessário

---

## Métricas de Sucesso

- ✅ 9 genes com textos completos (clinical_relevance + patient_explanation + conduct)
- ✅ Média de 1166 palavras por gene
- ✅ 15+ fontes científicas citadas
- ✅ 100% dos genes com intervenções nutrigenômicas específicas
- ✅ Textos em português-BR fluente
- ✅ Abordagem não-determinística (gene não é destino)
- ✅ Farmacogenômica incluída quando relevante

---

## Contato e Suporte

Para dúvidas sobre importação ou ajustes nos textos:
- Revisar `/home/user/plenya/GENETIC-ENRICHMENT-FINAL-REPORT.md`
- Verificar logs da API
- Consultar documentação do Plenya EMR em `/home/user/plenya/CLAUDE.md`

---

**Última atualização:** 2026-01-26
**Processador:** Claude Sonnet 4.5
**Próximo batch:** Genes 10-20 (pendente)
