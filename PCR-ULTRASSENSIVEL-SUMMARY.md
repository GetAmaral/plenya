# PCR Ultrassensível - Sumário Executivo

## ✅ Status: CONCLUÍDO

**Item:** PCR ultrassensível (hs-CRP)
**ID:** `9b8eb642-4b9c-48e2-b15d-ee2aeba78c83`
**Data:** 2026-01-28

---

## 📊 Métricas de Enriquecimento

| Campo | Tamanho |
|-------|---------|
| Clinical Relevance | 1.668 caracteres |
| Patient Explanation | 1.683 caracteres |
| Conduct | 2.472 caracteres |
| **Total** | **5.823 caracteres** |
| Artigos vinculados | **12 artigos** |
| Novos artigos inseridos | **3 artigos** |

---

## 🔬 Evidências Científicas (2024-2025)

### 1. ACC Scientific Statement 2025
- hsCRP preditor mais forte que LDL-colesterol em pacientes com estatinas
- Recomenda hsCRP ≥2 mg/L como "risk enhancer"

### 2. UK Biobank Study 2024
- 448.653 indivíduos analisados
- Associação forte com MACE, morte CV e mortalidade geral
- Integração ao SCORE2 melhora performance preditiva

### 3. StatPearls 2024
- Fisiopatologia: IL-6, cascata complemento
- Protocolo: 2 medições com intervalo de 2 semanas

---

## 📈 Valores de Referência

| Valor | Interpretação |
|-------|---------------|
| **<1 mg/L** | Baixo risco cardiovascular |
| **1-3 mg/L** | Risco cardiovascular moderado |
| **>3 mg/L** | Alto risco cardiovascular |
| **≥2 mg/L** | Risco residual inflamatório (prevenção secundária) |
| **>10 mg/L** | Inflamação aguda (excluir da avaliação CV) |

---

## 📚 Artigos Inseridos

1. **Inflammation and Cardiovascular Disease: 2025 ACC Scientific Statement**
   - JACC, 2025
   - DOI: 10.1016/j.jacc.2025.08.047
   - PubMed: 41020749

2. **C-reactive protein and cardiovascular risk in the general population**
   - European Heart Journal, 2024
   - DOI: 10.1093/eurheartj/ehaf937

3. **C-Reactive Protein: Clinical Relevance and Interpretation**
   - StatPearls, 2024
   - PubMed: 30020613

---

## 💡 Destaques do Conteúdo

### Clinical Relevance
- Fisiopatologia (IL-6, sistema complemento)
- Evidências UK Biobank e ACC 2025
- Interpretação clínica por faixas de risco
- Fatores modificadores (tabagismo, obesidade, exercício)

### Patient Explanation
- Linguagem clara e acessível
- Significado dos valores com analogias
- Ações práticas para melhorar os níveis
- Empoderamento do paciente

### Conduct
- Protocolo de solicitação (2 dosagens, intervalo 2 semanas)
- Condutas específicas por faixa de risco
- Abordagem terapêutica não farmacológica e farmacológica
- Monitoramento e metas
- Referências: ESC 2024, ACC 2025, AHA/ACC

---

## 🎯 Diferencial Clínico

✅ **Superioridade Preditiva:** hsCRP > LDL-C em pacientes com estatinas
✅ **Valores Atualizados:** Inclui ≥2 mg/L para risco residual
✅ **Integração SCORE2:** Melhora performance de modelos preditivos
✅ **Abordagem Completa:** Prevenção primária + secundária

---

## 🚀 Execução

**Script:** `/home/user/plenya/scripts/enrich_hscrp_item.py`

**Comando:**
```bash
docker run --rm --network plenya_plenya-network \
  -v /home/user/plenya/scripts:/scripts python:3.11-slim \
  bash -c "pip install psycopg2-binary && python3 /scripts/enrich_hscrp_item.py"
```

**Resultado:** ✅ Sucesso

---

## 📖 Fontes

- [ACC Statement 2025](https://www.jacc.org/doi/10.1016/j.jacc.2025.08.047)
- [European Heart Journal 2024](https://academic.oup.com/eurheartj/advance-article/doi/10.1093/eurheartj/ehaf937/8377304)
- [StatPearls - C-Reactive Protein](https://www.ncbi.nlm.nih.gov/books/NBK441843/)
- [Johns Hopkins - Assessing CV Risk with CRP](https://www.hopkinsmedicine.org/health/treatment-tests-and-therapies/assessing-cardiovascular-risk-with-c-reactive-protein)
- [ACC - hsCRP Risk Assessment Tool](https://www.acc.org/latest-in-cardiology/articles/2025/12/01/01/prioritizing-health-hscrp)

---

**Gerado em:** 2026-01-29
**Próximos passos:** Revisar com cardiologista, validar laboratório, testar frontend
