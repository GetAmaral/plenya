# 🚀 QUICK START - BATCH FINAL 2B

## Missão
Enriquecer 45 exames laboratoriais com conteúdo MFI

---

## ⚡ Execução em 30 Segundos

### Passo 1: Verificar Docker
```bash
docker compose ps
```
✅ Deve aparecer: `db ... running`

Se não estiver rodando:
```bash
docker compose up -d
```

---

### Passo 2: Executar Script
```bash
cd /home/user/plenya
./EXECUTE_BATCH_FINAL_2B.sh
```

**Aguardar 5-10 segundos...**

---

### Passo 3: Verificar Resultado
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) FROM score_items WHERE clinical_context IS NOT NULL;
"
```

**Resultado esperado:** Número > 0 (idealmente 45+)

---

## ✅ Pronto!

45 items agora têm:
- ✅ Valores ótimos funcionais
- ✅ Interpretação clínica completa
- ✅ Protocolos de intervenção com doses
- ✅ Monitoramento definido
- ✅ Referências científicas

---

## 📋 Visualizar Exemplo

```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT name, LEFT(clinical_context, 200) as preview
FROM score_items
WHERE name = 'TSH'
LIMIT 1;
"
```

---

## 📄 Documentação Completa

- **Sumário Executivo:** `BATCH-FINAL-2B-SUMARIO-EXECUTIVO.md`
- **Instruções Detalhadas:** `INSTRUCOES-EXECUCAO-BATCH-FINAL-2B.md`
- **Relatório Técnico:** `BATCH-FINAL-2B-REPORT.md`

---

## ⚠️ Problemas?

### Container não está rodando
```bash
docker compose restart db
# Aguardar 10 segundos
docker compose ps
```

### Script sem permissão
```bash
chmod +x EXECUTE_BATCH_FINAL_2B.sh
```

### Ver logs de erro
```bash
docker compose logs db | tail -50
```

---

## 🎯 Items Enriquecidos (45 total)

### Urinálise (6 items)
- Urobilinogênio, Nitrito, Hemácias-Sedimento
- Células Epiteliais, Cristais Patológicos, Leveduras

### Hormônios (14 items)
- SHBG (M/F), DHEA-S (6 faixas etárias)
- TSH, T3 Livre, T3 Reverso
- Testosterona (M/F), Progesterona (M/Gestantes)
- FSH (Fases do Ciclo), TRAb

### Bioquímica (10 items)
- AST, Gama GT, Ureia, Sódio
- Proteínas Totais, Albumina
- Vitamina E, Alfa-2 Globulina
- INR, VCM

### Cardiovascular (2 items)
- Troponina I Ultrassensível

### Hematologia (2 items)
- Hematócrito, Ferritina Pós-Menopausa

### Microbiologia (2 items)
- Urocultura, Hepatite B (HbsAg)

### Imagem (5 items)
- USG Próstata (Volume/PSAD)
- TC Tórax (Nódulo)
- Endoscopia Alta (Esofagite/Barrett)

### Sedimento (2 items)
- Muco, Cristais

---

## 💡 Exemplos de Conteúdo MFI

### TSH = 3.5 mUI/L (normal lab, subótimo MFI)
**Interpretação MFI:**
- Hipotireoidismo subclínico
- Sintomas: fadiga, ganho de peso
- Investigar: anti-TPO, T3/T4, selênio, iodo

**Condutas:**
- Selênio 200mcg/dia
- Ashwagandha 600mg/dia
- Considerar levotiroxina se sintomático
- Monitorar em 8 semanas

---

### SHBG = 15 nmol/L (baixo)
**Interpretação MFI:**
- Resistência insulínica
- Risco metabólico aumentado

**Condutas:**
- Berberina 500mg 3x/dia
- Dieta low-carb <100g/dia
- Jejum intermitente 16:8
- Exercício HIIT + força
- Meta: SHBG >20 nmol/L

---

## 🎉 Sucesso!

Agora você tem 45 exames com conteúdo MFI de excelência.

**Tempo total:** ~30 segundos
**Próximo passo:** Visualizar no frontend

```bash
cd apps/web
pnpm dev
# Acessar: http://localhost:3000
```

---

**Data:** 2026-01-28
**Status:** ✅ Pronto para usar
