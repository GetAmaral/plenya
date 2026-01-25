# Plano de Importação CSV → Banco de Dados Score System

**Data:** 24 de Janeiro de 2026
**Arquivo fonte:** `Escore Plenya SPL v1.3.csv` (817 linhas)
**Objetivo:** Popular tabelas score_groups, score_subgroups, score_items, score_levels

---

## 1. Análise da Estrutura do CSV

### Colunas do CSV (separador `;`)

| Posição | Nome Header | Descrição | Exemplo |
|---------|-------------|-----------|---------|
| **1** | Nível 1 | **Grupo** (sticky - se vazio, usar último preenchido) | "Alimentação", "Exames" |
| **2** | Nível 2 | **Subgrupo** (sticky - se vazio, usar último preenchido) | "Histórico", "Laboratoriais" |
| **3** | Nível 3 | **Item** ou **Nome do Exame** | "Pré-natal", "Ácido úrico - homem" |
| **4** | Nível 4 | **Subitem** OU **Unit \| Conversion** | "mg/dL \| mg/dL × 59,48 = μmol/L" |
| **5** | Pontos EPSPL | **Points** (0-100) | "20", "5", "10" |
| **6-11** | Nivel 0-5 | **Definições dos níveis** | "≥7,0", "6,0–6,9", etc. |

### Regras de Interpretação

#### A) Grupo e Subgrupo (Colunas 1 e 2)
- **Sticky behavior**: Só muda quando aparece novo valor
- Quando vazio, continua do anterior
- Exemplo:
  ```
  Alimentação;Histórico;...    <- Grupo="Alimentação", Subgrupo="Histórico"
  ;;Pré-natal;...               <- Continua: Grupo="Alimentação", Subgrupo="Histórico"
  ;;;Item filho;...             <- Continua: Grupo="Alimentação", Subgrupo="Histórico"
  ```

#### B) Item (Coluna 3) e Subitem/Unit (Coluna 4)

**CASO 1: Col3 preenchida + Col4 vazia**
- Item standalone (sem unit/conversion)
- Pode ter ou não points
```csv
;;Divisão e Horários;;10;...  <- Item simples com points
```

**CASO 2: Col3 preenchida + Col4 com "|"**
- Item é um **exame laboratorial**
- Col3 = Nome do item
- Col4 = "unit | conversion"
- Parsear Col4: split por " | "
  - array[0] = unit
  - array[1] = conversion (se existir)
```csv
;;Ácido úrico - homem;mg/dL | mg/dL × 59,48 = μmol/L;20;≥7,0;...
```

**CASO 3: Col3 vazia + Col4 preenchida**
- Item hierárquico (filho do último item da Col3 que tinha points > 0)
- Col4 = Nome do subitem
- parent_item_id = último item criado da Col3
```csv
;;Curva insulina;;0;...       <- Item pai descritivo (NÃO criar, points=0)
;;;GLICOSE 0 MIN;...;20;...   <- Filho independente? OU usar "Curva..." como pai?
```

#### C) Points (Coluna 5)

- Se vazio ou "0": Item descritivo/organizador
  - **NÃO criar score_item**
  - Apenas trackearcomo possível parent para hierarquia
- Se > 0: Item de pontuação
  - **Criar score_item** com points
  - Criar 6 níveis (0-5)

#### D) Níveis 0-5 (Colunas 6-11)

Para cada coluna preenchida, criar ScoreLevel:
- **level** = 0, 1, 2, 3, 4, 5 (conforme posição)
- **name** = texto exato da coluna
- **definition** = texto exato
- **lower_limit** e **upper_limit** = parseados do texto

**Padrões de parsing:**

| Padrão | Exemplo | lower_limit | upper_limit |
|--------|---------|-------------|-------------|
| `X a Y` | "6,0–6,9" | "6.0" | "6.9" |
| `X–Y` | "6,0–6,9" | "6.0" | "6.9" |
| `X - Y` | "6.0 - 6.9" | "6.0" | "6.9" |
| `<X` ou `<=X` | "<10" | NULL | "10" |
| `>X` ou `>=X` | "≥7,0" | "7.0" | NULL |
| `Texto puro` | "Remissão" | NULL | NULL |

**Normalização:**
- Vírgula → ponto: "6,9" → "6.9"
- Espaços removidos
- `≥` → mantenha como string "≥7.0"
- `≤` → mantenha como string "≤10"

**Casos com "ou":**
```
"<0.5 ou 1.21 a 1.5"  <- Pegar apenas PRIMEIRA parte: "<0.5"
">100 ou <10"         <- Pegar apenas PRIMEIRA parte: ">100"
```

---

## 2. Algoritmo Proposto

```python
# Estado global (sticky)
current_group_name = None
current_group_id = None
current_subgroup_name = None
current_subgroup_id = None
last_item_col3_id = None  # Para parent_item_id

# Contadores de ordem
group_order = 0
subgroup_order = {}  # dict[group_id] = int
item_order = {}      # dict[subgroup_id] = int

for row in csv_rows[1:]:  # Pular header
    col1 = row[0].strip()  # Grupo
    col2 = row[1].strip()  # Subgrupo
    col3 = row[2].strip()  # Item ou Nome Exame
    col4 = row[3].strip()  # Subitem ou Unit|Conversion
    col5 = parse_float(row[4])  # Points
    niveis = [row[5], row[6], row[7], row[8], row[9], row[10]]

    # 1. Processar Grupo
    if col1:
        current_group_name = col1
        group_order += 1
        current_group_id = insert_or_get_group(col1, group_order)
        subgroup_order[current_group_id] = 0

    # 2. Processar Subgrupo
    if col2:
        current_subgroup_name = col2
        subgroup_order[current_group_id] += 1
        current_subgroup_id = insert_or_get_subgroup(
            col2,
            current_group_id,
            subgroup_order[current_group_id]
        )
        item_order[current_subgroup_id] = 0

    # 3. Pular linhas sem conteúdo
    if not col3 and not col4:
        continue

    # 4. Se points == 0, não criar item (apenas tracker)
    if col5 == 0:
        if col3:
            # Atualizar tracker para possível parent
            last_item_col3_id = None  # OU criar item fake?
        continue

    # 5. Processar Item
    item_order[current_subgroup_id] += 1
    parent_id = None
    unit = None
    unit_conversion = None

    if col3:
        # Col3 preenchida
        if col4 and " | " in col4:
            # EXAME: col4 = unit | conversion
            parts = col4.split(" | ", 1)
            unit = parts[0].strip()
            unit_conversion = parts[1].strip() if len(parts) > 1 else None
            item_name = col3
        else:
            # Item simples
            item_name = col3

        last_item_col3_id = insert_score_item(
            name=item_name,
            unit=unit,
            unit_conversion=unit_conversion,
            points=col5,
            order=item_order[current_subgroup_id],
            subgroup_id=current_subgroup_id,
            parent_item_id=None
        )
        item_id = last_item_col3_id

    elif col4:
        # Col3 vazia, col4 preenchida = subitem hierárquico
        item_id = insert_score_item(
            name=col4,
            unit=None,
            unit_conversion=None,
            points=col5,
            order=item_order[current_subgroup_id],
            subgroup_id=current_subgroup_id,
            parent_item_id=last_item_col3_id
        )

    # 6. Processar Níveis (0-5)
    for level_idx, nivel_text in enumerate(niveis):
        if not nivel_text or nivel_text.strip() == "":
            continue

        nivel_text = nivel_text.strip()
        lower, upper = parse_range(nivel_text)

        insert_score_level(
            level=level_idx,
            name=nivel_text,
            definition=nivel_text,
            lower_limit=lower,
            upper_limit=upper,
            item_id=item_id
        )
```

---

## 3. Função de Parsing de Ranges

```python
import re

def parse_range(text):
    """
    Retorna (lower_limit, upper_limit) como strings ou None
    """
    # Normalizar
    text = text.replace(",", ".")
    text = text.strip()

    # Se tem "ou", pegar apenas primeira parte
    if " ou " in text:
        text = text.split(" ou ")[0].strip()

    # Padrão: >=X ou >X
    match = re.match(r'^[>≥]=?\s*([0-9.]+)', text)
    if match:
        return (match.group(1), None)

    # Padrão: <=X ou <X
    match = re.match(r'^[<≤]=?\s*([0-9.]+)', text)
    if match:
        return (None, match.group(1))

    # Padrão: X a Y, X–Y, X - Y
    match = re.match(r'^([0-9.]+)\s*(?:a|–|-)\s*([0-9.]+)', text)
    if match:
        return (match.group(1), match.group(2))

    # Se não tem número, é texto puro
    if not re.search(r'[0-9]', text):
        return (None, None)

    # Número único exato
    match = re.match(r'^([0-9.]+)$', text)
    if match:
        return (match.group(1), match.group(1))

    # Não conseguiu parsear
    return (None, None)
```

---

## 4. Implementação: Script Python

### Estrutura do Script

```
/home/user/plenya/scripts/import_score_csv.py
```

**Dependências:**
- `psycopg2-binary` (PostgreSQL)
- `csv` (stdlib)
- `re` (stdlib)
- `argparse` (stdlib)

**Funções principais:**
- `main()` - Entry point
- `parse_csv(filepath)` - Lê e processa CSV
- `parse_range(text)` - Parse limites
- `insert_group(name, order)` - Insert grupo
- `insert_subgroup(name, group_id, order)` - Insert subgrupo
- `insert_item(...)` - Insert item
- `insert_level(...)` - Insert nível
- `validate_import()` - Queries de validação

### Execução

```bash
# Instalar dependência
pip3 install psycopg2-binary

# Executar import
python3 scripts/import_score_csv.py \
  --csv "Escore Plenya SPL v1.3.csv" \
  --host localhost \
  --port 5432 \
  --user plenya_user \
  --password plenya_dev_password \
  --database plenya_db
```

---

## 5. Validações Pós-Import

### A) Count de Registros

```sql
SELECT
    'Grupos' as tabela, COUNT(*) FROM score_groups
UNION ALL SELECT 'Subgrupos', COUNT(*) FROM score_subgroups
UNION ALL SELECT 'Itens', COUNT(*) FROM score_items
UNION ALL SELECT 'Níveis', COUNT(*) FROM score_levels;
```

### B) Itens sem 6 níveis

```sql
SELECT i.id, i.name, COUNT(l.id) as nivel_count
FROM score_items i
LEFT JOIN score_levels l ON l.item_id = i.id
WHERE i.points > 0
GROUP BY i.id, i.name
HAVING COUNT(l.id) != 6;
```

### C) Níveis duplicados

```sql
SELECT item_id, level, COUNT(*)
FROM score_levels
GROUP BY item_id, level
HAVING COUNT(*) > 1;
```

### D) Hierarquia

```sql
-- Itens com parent
SELECT i.name, p.name as parent_name
FROM score_items i
JOIN score_items p ON i.parent_item_id = p.id
LIMIT 20;
```

---

## 6. Casos Especiais a Testar

### A) Exames Laboratoriais

**Linha esperada:**
```csv
;;Ácido úrico - homem;mg/dL | mg/dL × 59,48 = μmol/L;20;≥7,0;6,0–6,9;5,5–5,9;<3,0;5,0–5,4;3,5–4,9;
```

**Resultado esperado:**
- name: "Ácido úrico - homem"
- unit: "mg/dL"
- unit_conversion: "mg/dL × 59,48 = μmol/L"
- points: 20
- 6 níveis com ranges parseados

### B) Itens Hierárquicos

**Linhas esperadas:**
```csv
;;Eletroforese de proteínas;;0;...     <- NÃO criar (points=0)
;;;Proteínas Totais;...;20;...        <- Criar independente? OU filho?
```

**DÚVIDA:** Se item anterior tem points=0, os próximos ";;;" são filhos dele?

### C) Ranges Complexos

```csv
Nível: "<0.5 ou 1.21 a 1.5"
```
**Esperado:** lower=NULL, upper="0.5"

### D) Caracteres Especiais

```csv
"≥7,0" -> lower="7.0", upper=NULL
"6,0–6,9" -> lower="6.0", upper="6.9"
```

---

## 7. Perguntas para Aprovação

### 1. Itens com points=0
**Pergunta:** Quando item tem points=0, devo:
- A) NÃO criar score_item (apenas skip)
- B) Criar score_item sem níveis (para organização)
- C) Usar como parent_item_id para próximos ";;;" subitems

**Recomendação:** **Opção A** (não criar)

### 2. Hierarquia ambígua
**Exemplo:**
```csv
;;Curva insulina;;0;...
;;;GLICOSE 0 MIN;...;20;...
;;;GLICOSE 30 MIN;...;20;...
```

**Pergunta:** "GLICOSE 0 MIN" deve ter parent_item_id = "Curva insulina"?
- A) Sim, usar último item da col3 (mesmo se points=0)
- B) Não, criar independente

**Recomendação:** **Opção B** (independente), pois "Curva" não foi criado no DB

### 3. Script Python vs Go
**Pergunta:** Prosseguir com Python?
- ✅ Mais rápido
- ✅ Mais fácil debug
- ✅ Pode executar agora

---

## 8. Próximos Passos (após aprovação)

1. ✅ Criar arquivo `scripts/import_score_csv.py`
2. ✅ Implementar funções de parsing
3. ✅ Implementar inserções no DB
4. ✅ Testar com primeiras 100 linhas
5. ✅ Executar import completo
6. ✅ Validar resultados
7. ✅ Corrigir erros se necessário

---

**Status:** 📋 PLANO COMPLETO - AGUARDANDO SUA APROVAÇÃO

**Você aprova este plano? Alguma modificação necessária?**
