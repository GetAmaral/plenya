# BATCH 4 - GENÉTICA: PROMPT PARA ENRIQUECIMENTO

**Status:** ⏳ Preparado para execução após conclusão do Batch 3

---

## MISSÃO

Enriquecer **81 genes** do grupo Genética com conteúdo científico de medicina genômica funcional integrativa.

## ESTRUTURA DOS GENES

### Metabolismo (28 genes)
- **Diabetes/Glicemia:** ABCC8, CDKAL1, GCK (MODY2), HNF1A (MODY3), HNF1B (MODY5), HNF4A (MODY1), HHEX, IGF2BP2, INS VNTR, IRS1, KCNJ11, SLC30A8, TCF7L2, PPARG
- **Obesidade:** FTO, MC4R, LEPR, POMC, PPARA, PPARGC1A
- **Vitaminas/Nutrientes:** BCO1 (Vit A), FADS1/FADS2 (Ômega-3), MCM6 (Lactose), MTHFR (Homocisteína), SLC23A1 (Vit C), VDR (Vit D)
- **Gorduras:** FABP2

### Cardiovascular (19 genes)
- **Hipertensão:** ACE, AGT, AGTR1, ADD1, CYP11B2, GNB3, NOS3
- **Colesterol/Lipídios:** ABCA1, APOA1, APOA5, APOC2, APOE, GPIHBP1, LCAT, LDLR, LIPC, LMF1, LPL, PCSK9

### Neurodegeneração (11 genes)
- **Alzheimer:** APP, APOE, PSEN1, PSEN2
- **Parkinson:** LRRK2, PARK2, PARK7/DJ-1, PINK1, SNCA
- **Demências:** C9orf72, GRN, MAPT

### Detoxificação (13 genes)
- **Fase I:** CYP1A1, CYP1A2, CYP2A6
- **Fase II:** GSTM1, GSTP1, GSTT1, NAT2, EPHX1
- **Antioxidantes:** CAT, GPX1, SOD2
- **Álcool:** ADH1B, ALDH2

### Imunidade (5 genes)
- **Inflamação:** CRP, IL1B, IL6, TNF
- **Autoimune:** HLA-DQ2/DQ8

### Performance (4 genes)
- **Músculo:** ACTN3
- **Osso/Tendão:** COL1A1, COL5A1, ESR1

### Outros (1 gene)
- ALPL (Hipofosfatasia)

---

## ESTRATÉGIA DE EXECUÇÃO

### Abordagem sem confirmações:

```bash
# 1. Gerar arquivo SQL único com todos os 81 genes
cat > batch4_genetica_enrichment.sql <<'EOF'
-- Todos os 81 genes aqui
EOF

# 2. Executar UMA ÚNICA VEZ
cat batch4_genetica_enrichment.sql | docker compose exec -T db psql -U plenya_user -d plenya_db
```

---

## BUSCAR ARTIGOS GENÉTICOS

Primeiro, execute para ver artigos disponíveis:

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT id, title FROM articles
WHERE title ILIKE '%genética%'
   OR title ILIKE '%genômica%'
   OR title ILIKE '%SNP%'
   OR title ILIKE '%polimorfismo%'
   OR title ILIKE '%nutrigenômica%'
   OR title ILIKE '%farmacogenômica%'
   OR title ILIKE '%epigenética%'
LIMIT 50;"
```

---

## TEMPLATE SQL (para cada gene)

```sql
-- ============================================================
-- Gene: [NOME] - Subgrupo: [SUBGRUPO]
-- ============================================================

UPDATE score_items
SET
  clinical_relevance = 'Texto técnico sobre o gene (150-300 palavras).

ESTRUTURA:
1. Função do gene e proteína codificada
2. SNP/Variante específica e frequência populacional
3. Genótipos e alelos de risco
4. Impacto funcional (expressão, atividade enzimática, etc.)
5. Associações clínicas e evidências (ORs, RRs)
6. Interações gene-ambiente (nutrigenômica, farmacogenômica)
7. Abordagem funcional integrativa

EXEMPLO (MTHFR):
"O gene MTHFR codifica a enzima metilenotetrahidrofolato redutase, essencial no ciclo do folato e metabolismo da homocisteína. O SNP rs1801133 (C677T) resulta em substituição Ala222Val, reduzindo atividade enzimática em ~35% (CT) a ~70% (TT) sob condições de baixo folato. Frequência alélica T: ~30% caucasianos, 10-15% africanos, 40-50% asiáticos. Genótipo TT associa-se a hiperhomocisteinemia (HR 2.0-3.0), risco cardiovascular aumentado (OR 1.16 para IAM), defeitos de tubo neural (OR 1.8), depressão (OR 1.36) e pior resposta a antidepressivos. A interação gene-nutriente é crítica: suplementação com 5-MTHF, B12 e B6 normaliza homocisteína. Na medicina funcional, investigamos também status de B2 (cofator), estresse oxidativo e polimorfismos combinados (MTR, MTRR, CBS). Teste genético orienta dose e forma de suplementação..."',

  patient_explanation = 'Texto para paciente (100-200 palavras).

ESTRUTURA:
1. O que é o gene (analogia simples)
2. O que a variante faz
3. Como isso afeta você
4. O que pode ser feito
5. Mensagem positiva

EXEMPLO (MTHFR):
"O gene MTHFR funciona como uma 'fábrica' que processa vitaminas do complexo B (especialmente ácido fólico) para que seu corpo possa usá-las. Algumas pessoas nascem com uma variante desse gene que faz a fábrica trabalhar mais devagar. Se você tem essa variante e não consome folato suficiente, pode acumular uma substância chamada homocisteína no sangue, o que aumenta risco de problemas cardiovasculares, dificulta a gravidez e pode afetar o humor. A boa notícia: isso é completamente manejável! Basta usar a forma 'pronta' de folato (5-MTHF ou metilfolato) em vez do ácido fólico comum, junto com vitaminas B12 e B6. É como dar o produto final para a fábrica lenta - ela não precisa processar, apenas usar. Com suplementação adequada, seus níveis de homocisteína normalizam e os riscos diminuem significativamente..."',

  conduct = 'Orientações de conduta (80-150 palavras).

ESTRUTURA:
1. Quando testar
2. Método de teste
3. Interpretação de resultados
4. Exames complementares
5. Intervenções por genótipo
6. Monitoramento

EXEMPLO (MTHFR):
"Testar quando: histórico familiar cardiovascular precoce, trombose, abortos recorrentes, depressão refratária, defeitos tubo neural. Método: PCR em tempo real para rs1801133 (C677T). Interpretar: CC = normal, CT = heterozigoto (35% ↓atividade), TT = homozigoto (70% ↓atividade). Complementar com: homocisteína plasmática (VR <10 µmol/L), vitamina B12, folato eritrocitário, B6, B2. Conduta por genótipo: CC (folato normal 400mcg/dia), CT (5-MTHF 800-1000mcg/dia), TT (5-MTHF 1000-2000mcg/dia). Sempre associar: B12 metilada 1000-2000mcg/dia, B6 (P5P) 50-100mg/dia, B2 10mg/dia. Evitar: ácido fólico sintético (pode mascarar deficiência B12). Monitorar homocisteína a cada 3-6 meses até normalização. Meta: <7 µmol/L..."',

  updated_at = NOW()
WHERE id = '[UUID-DO-GENE]';

-- Linkar artigos científicos
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '[UUID-DO-GENE]'
FROM articles a
WHERE a.title ILIKE '%[termo-relevante]%'
LIMIT 5
ON CONFLICT DO NOTHING;
```

---

## TÓPICOS ESSENCIAIS POR CATEGORIA

### Genes de Metabolismo

**Foco:**
- Risco de diabetes tipo 2
- Sensibilidade à insulina
- Metabolismo de nutrientes
- MODY (diabetes monogênico)
- Nutrigenômica (resposta a dietas)

**Intervenções:**
- Dieta personalizada (low-carb, mediterrânea, jejum)
- Exercício (resistência vs aeróbico)
- Suplementação (cromo, ácido alfa-lipóico, berberina, inositol)
- Monitoramento glicêmico

### Genes Cardiovasculares

**Foco:**
- Metabolismo de lipoproteínas
- Regulação da pressão arterial
- Risco de IAM/AVC
- Farmacogenômica (estatinas, anti-hipertensivos)

**Intervenções:**
- Padrão alimentar (ômega-3, gorduras, sódio)
- Exercício aeróbico
- Suplementação (CoQ10, ômega-3, magnésio)
- Escolha de estatina baseada em APOE

### Genes de Neurodegeneração

**Foco:**
- Risco de Alzheimer/Parkinson
- Idade de início
- Progressão
- Proteção neuronal

**Intervenções:**
- Neuroproteção (curcumina, resveratrol, ômega-3)
- Estimulação cognitiva
- Exercício físico
- Controle de fatores de risco (diabetes, hipertensão)
- Dieta MIND/mediterrânea

### Genes de Detoxificação

**Foco:**
- Metabolismo de xenobióticos
- Capacidade antioxidante
- Metabolismo de cafeína/álcool
- Risco de toxicidade medicamentosa

**Intervenções:**
- Redução de exposições
- Suporte detox (glutationa, NAC, ALA)
- Ajuste de doses medicamentosas
- Timing de cafeína

### Genes de Imunidade

**Foco:**
- Resposta inflamatória
- Risco autoimune (celíaca)
- Níveis de citocinas

**Intervenções:**
- Dieta anti-inflamatória
- Eliminação de gatilhos (glúten para HLA-DQ2/8)
- Modulação imune (ômega-3, curcumina, probióticos)

### Genes de Performance

**Foco:**
- Tipo de fibra muscular
- Risco de lesões
- Densidade óssea
- Resposta ao treino

**Intervenções:**
- Treino personalizado (força vs resistência)
- Suplementação (colágeno, vitamina D, cálcio)
- Prevenção de lesões

---

## LINGUAGEM E ESTILO

### Clinical Relevance (Técnico)

**Use:**
- Nomenclatura genética correta (rs números, alelos, genótipos)
- Frequências alélicas por etnia
- Odds ratios, hazard ratios, valores de p
- Mecanismos moleculares
- Interações gene-gene e gene-ambiente
- Evidências de GWAS, metanálises

**Exemplo de terminologia:**
- "SNP rs1801133 (C677T) no éxon 4"
- "Substituição Ala222Val"
- "Frequência do alelo T: 0.32 (CEU), 0.12 (YRI)"
- "OR 1.68 (IC95% 1.23-2.31) para IAM em homozigotos TT"
- "Interação gene-nutriente: ↑ atividade com riboflavina"

### Patient Explanation (Acessível)

**Use:**
- Analogias ("gene como receita", "enzima como fábrica")
- Linguagem empática
- Foco em ações, não em fatalismo
- Genética como "tendência", não "destino"
- Empoderamento através do conhecimento

**Evite:**
- Determinismo genético
- Linguagem alarmista
- Jargão sem explicação
- Fatalismos ("você vai ter...")

### Conduct (Prático)

**Inclua sempre:**
- Quando testar (indicações)
- Como testar (método)
- Como interpretar (por genótipo)
- Exames complementares
- Intervenções específicas (doses)
- Follow-up

---

## ARTIGOS CIENTÍFICOS

### Buscar por temas:

1. **Nutrigenômica geral**
2. **Farmacogenômica**
3. **Epigenética**
4. **Medicina Genômica Funcional**
5. **Testes genéticos na prática clínica**

### Linkar 3-5 artigos por gene

Priorizar artigos de:
- Medicina Funcional Integrativa
- Nutrigenômica aplicada
- Guidelines de interpretação
- Casos clínicos

---

## QUALIDADE MÍNIMA

### Clinical Relevance
- ✅ Mínimo 900 caracteres
- ✅ Nomenclatura genética correta
- ✅ Frequências populacionais
- ✅ Evidências quantitativas (ORs, etc.)
- ✅ Mecanismos moleculares
- ✅ Interações gene-ambiente

### Patient Explanation
- ✅ Mínimo 600 caracteres
- ✅ Linguagem acessível
- ✅ Analogias práticas
- ✅ Foco em ações
- ✅ Tom empoderador

### Conduct
- ✅ Mínimo 500 caracteres
- ✅ Indicações de teste claras
- ✅ Interpretação por genótipo
- ✅ Doses específicas
- ✅ Follow-up definido

---

## CHECKLIST FINAL

Antes de executar, verificar:

- [ ] Arquivo SQL único gerado (`batch4_genetica_enrichment.sql`)
- [ ] 81 genes incluídos
- [ ] Todos os 7 subgrupos cobertos
- [ ] Artigos buscados e IDs coletados
- [ ] Nomenclatura genética correta (rs números)
- [ ] Textos em português BR
- [ ] Sem emojis
- [ ] Tamanhos adequados
- [ ] SQL testado em 1-2 genes primeiro

---

## EXECUÇÃO

```bash
# Gerar e executar
cat batch4_genetica_enrichment.sql | docker compose exec -T db psql -U plenya_user -d plenya_db > batch4_output.log 2>&1

# Verificar sucesso
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
  COUNT(*) FILTER (WHERE LENGTH(clinical_relevance) > 100) as enriquecidos,
  COUNT(*) as total
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups g ON sg.group_id = g.id
WHERE g.name = 'Genética';"
```

---

**Status:** Pronto para execução! 🧬
**Estimativa:** 15-20 minutos para 81 genes
**Próximo passo:** Aguardar conclusão do Batch 3 e iniciar
