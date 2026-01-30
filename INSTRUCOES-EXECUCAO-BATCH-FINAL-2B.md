# INSTRUÇÕES DE EXECUÇÃO - BATCH FINAL 2B

## Missão
Enriquecer 45 items de exames laboratoriais com conteúdo MFI.

## Arquivos Prontos

```
/home/user/plenya/
├── scripts/enrichment_data/
│   ├── batch_final_2_exames_B.json          # JSON fonte (45 items)
│   ├── batch_final_2_exames_B.sql           # SQL Parte 1 (18 items detalhados)
│   ├── batch_final_2_exames_B_part2.sql     # SQL Parte 2 (7 items)
│   └── batch_final_2_exames_B_COMPLETE.sql  # SQL Parte 3 (20 items otimizados)
├── EXECUTE_BATCH_FINAL_2B.sh                # Script automatizado
├── BATCH-FINAL-2B-REPORT.md                 # Relatório completo
└── INSTRUCOES-EXECUCAO-BATCH-FINAL-2B.md    # Este arquivo
```

## EXECUÇÃO (Escolha uma opção)

### OPÇÃO 1: Script Automatizado (RECOMENDADO)

```bash
cd /home/user/plenya
./EXECUTE_BATCH_FINAL_2B.sh
```

**Saída esperada:**
```
==========================================
BATCH FINAL 2 - PARTE B: 45 ITEMS
Enrichment MFI de Exames Laboratoriais
==========================================

✅ Container PostgreSQL está rodando

📋 Executando Parte 1 (Items 1-18 - detalhados)...
UPDATE 18
✅ Parte 1 concluída

📋 Executando Parte 2 (Items 19-25 - complementares)...
UPDATE 7
✅ Parte 2 concluída

📋 Executando Parte 3 (Items 26-45 - otimizados)...
UPDATE 20
✅ Parte 3 concluída

🔍 Verificando enrichment...

==========================================
✅ ENRICHMENT CONCLUÍDO
==========================================
📊 Items enriquecidos verificados: 45
🎯 Meta: 45 items
```

---

### OPÇÃO 2: Execução Manual (3 comandos independentes)

#### Passo 1: Verificar container do PostgreSQL
```bash
docker compose ps
```
Se não estiver rodando:
```bash
docker compose up -d
```

#### Passo 2: Executar os 3 SQLs em sequência

```bash
# Parte 1: Items 1-18 (detalhados - Urobilinogênio, Nitrito, SHBG, DHEA-S, TSH, T3, INR)
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B.sql

# Parte 2: Items 19-25 (complementares - Testosterona, TRAb, AST, Troponina, Ureia)
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_part2.sql

# Parte 3: Items 26-45 (otimizados - Vitamina E, VCM, Ferritina, FSH, Sódio, etc)
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/enrichment_data/batch_final_2_exames_B_COMPLETE.sql
```

#### Passo 3: Verificar resultado
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) as total_enriquecidos
FROM score_items
WHERE clinical_context IS NOT NULL
  AND clinical_context != ''
  AND id IN (
    'bf77b326-caa5-46fd-b607-70a089918780',
    '1aa25d4b-a972-40db-a288-9cbe506de99e',
    '814d923f-cdfa-4388-9ba1-42b23dcd8d6d',
    '09577ef1-c3ad-461b-b2ad-59fab2c193d5',
    'ebcc36fd-d285-4754-adf7-50c7b130b286'
  );
"
```

**Resultado esperado:**
```
 total_enriquecidos
--------------------
                  5
```

---

### OPÇÃO 3: Visualizar exemplo de enrichment

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    name,
    clinical_context,
    functional_ranges->'optimal' as optimal_range,
    jsonb_array_length(related_biomarkers) as num_related_biomarkers
FROM score_items
WHERE id = 'bf77b326-caa5-46fd-b607-70a089918780';
"
```

**Saída esperada:**
```
        name        |                     clinical_context                      |            optimal_range             | num_related_biomarkers
--------------------+-----------------------------------------------------------+--------------------------------------+------------------------
 Urobilinogênio     | Produto da degradação da bilirrubina conjugada...         | {"min": 0.1, "max": 1.0, "unit"...   |                      6
```

---

## Verificação Completa

### Contar TODOS os items enriquecidos:
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    COUNT(*) FILTER (WHERE clinical_context IS NOT NULL) as com_clinical_context,
    COUNT(*) FILTER (WHERE functional_ranges IS NOT NULL) as com_functional_ranges,
    COUNT(*) FILTER (WHERE biomarker_interpretation IS NOT NULL) as com_interpretation,
    COUNT(*) as total
FROM score_items
WHERE id IN (
    SELECT jsonb_array_elements_text(
        '[\"bf77b326-caa5-46fd-b607-70a089918780\",
          \"1aa25d4b-a972-40db-a288-9cbe506de99e\",
          \"814d923f-cdfa-4388-9ba1-42b23dcd8d6d\",
          \"34af6e5c-3847-46d8-874e-a7364c014877\",
          \"d164eacf-a0d7-48f2-899d-3f0d57ec7cc3\"]'::jsonb
    )
);
"
```

### Visualizar item completo (exemplo TSH):
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
    name,
    jsonb_pretty(functional_ranges) as ranges,
    jsonb_pretty(functional_medicine_interventions) as interventions
FROM score_items
WHERE id = '34af6e5c-3847-46d8-874e-a7364c014877';
" | less
```

---

## Estrutura de Dados Enriquecidos

Cada item dos 45 recebeu:

### 1. `clinical_context` (TEXT)
Contexto clínico, fisiologia e significado do biomarcador.

### 2. `functional_ranges` (JSONB)
```json
{
  "optimal": {
    "min": 0.5,
    "max": 2.0,
    "unit": "mUI/L",
    "description": "Faixa funcional ideal para metabolismo"
  },
  "suboptimal": {...},
  "critical": {...}
}
```

### 3. `biomarker_interpretation` (JSONB)
```json
{
  "low": {
    "meaning": "Significado clínico",
    "causes": ["Causa 1", "Causa 2"],
    "clinical_significance": "Implicações",
    "symptoms": ["Sintoma 1", "Sintoma 2"]
  },
  "optimal": {...},
  "high": {...}
}
```

### 4. `functional_medicine_interventions` (JSONB)
```json
{
  "tsh_elevado": {
    "investigation": ["Exame 1", "Exame 2"],
    "lifestyle": ["Mudança 1", "Mudança 2"],
    "supplements": [
      "Selênio 200mcg/dia",
      "Ashwagandha 600mg/dia"
    ],
    "monitoring": "Repetir em 3 meses"
  }
}
```

### 5. `related_biomarkers` (JSONB Array)
```json
["T4 Livre", "T3 Livre", "Anti-TPO", "Ultrassom Tireoide"]
```

### 6. `scientific_references` (JSONB Array)
```json
[
  "Garber JR, et al. Clinical practice guidelines for hypothyroidism. Endocr Pract. 2012.",
  "Ross DS, et al. 2016 ATA guidelines. Thyroid. 2016."
]
```

---

## Diferenciais MFI

| Aspecto | Convencional | MFI (Este Batch) |
|---------|--------------|------------------|
| **Valores** | Laboratoriais genéricos | Funcionais otimizados por idade/sexo |
| **Interpretação** | Normal/Patológico | Subótimo/Ótimo/Crítico + Causas raiz |
| **Tratamento** | Farmacológico | Lifestyle + Nutraceuticals (DOSES) + Fármacos |
| **Monitoramento** | Genérico | Específico com prazos |
| **Integração** | Isolado | Biomarcadores relacionados |
| **Evidências** | Raras | Referências científicas atualizadas |

---

## Exemplos de Qualidade

### TSH (Item 15)
- **Range MFI:** 0.5-2.0 mUI/L (ótimo) vs 0.4-4.5 (laboratorial)
- **Interpretação:** TSH >2.5 = hipotireoidismo subclínico com sintomas reais
- **Condutas:**
  - Selênio 200mcg/dia (reduz anti-TPO)
  - Ashwagandha 600mg/dia (melhora T3/T4)
  - Iodo 150-300mcg/dia (APENAS se deficiência confirmada)
  - Levotiroxina se TSH >4.5 ou sintomático

### T3 Reverso (Item 17)
- **Razão T3/rT3:** <10 = bloqueio funcional
- **Condutas:**
  - Selênio 200-400mcg/dia (clearance de rT3)
  - NAC 600mg 2x/dia (detox)
  - Reduzir T4, adicionar T3 direto

### SHBG Homens (Item 7)
- **Range MFI:** 20-50 nmol/L (ótimo) vs 10-70 (laboratorial)
- **SHBG <20:** Síndrome metabólica, resistência insulínica
  - Berberina 500mg 3x/dia
  - Inositol 2-4g/dia
  - Dieta low-carb + jejum intermitente

---

## Troubleshooting

### Erro: "permission denied"
```bash
chmod +x EXECUTE_BATCH_FINAL_2B.sh
```

### Erro: "database plenya_db does not exist"
```bash
docker compose down
docker compose up -d
# Aguardar 10 segundos
docker compose exec db psql -U plenya_user -l
```

### Erro: "connection refused"
```bash
docker compose restart db
# Aguardar 5 segundos
docker compose exec db psql -U plenya_user -d plenya_db -c "SELECT 1;"
```

### Verificar se items já foram enriquecidos
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT name,
       CASE WHEN clinical_context IS NOT NULL THEN 'Enriquecido' ELSE 'Pendente' END as status
FROM score_items
WHERE id = 'bf77b326-caa5-46fd-b607-70a089918780';
"
```

---

## Próximos Passos Após Execução

1. ✅ Verificar enrichment no banco (queries acima)
2. ✅ Testar visualização no frontend:
   ```bash
   cd apps/web
   pnpm dev
   # Acessar: http://localhost:3000/scores
   ```
3. ✅ Validar conteúdo clínico com especialistas
4. ✅ Gerar relatório de impacto para stakeholders

---

## Suporte

Caso encontre problemas:
1. Verificar logs do Docker: `docker compose logs db`
2. Verificar estrutura da tabela: `docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"`
3. Verificar UUIDs dos items: comparar com `batch_final_2_exames_B.json`

---

**Última atualização:** 2026-01-28
**Status:** ✅ Pronto para execução
**Total de items:** 45
**Tempo estimado:** 5-10 segundos
