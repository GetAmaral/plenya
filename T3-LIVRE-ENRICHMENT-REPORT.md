# Relatório de Enriquecimento: T3 Livre

**Data:** 2026-01-28
**Item ID:** `d164eacf-a0d7-48f2-899d-3f0d57ec7cc3`
**Nome:** T3 Livre
**Categoria:** Exames > Laboratoriais

---

## ✅ Status: CONCLUÍDO COM SUCESSO

---

## 📊 Resumo Executivo

O item **T3 Livre** foi enriquecido com sucesso no banco de dados, incluindo:

1. **3 artigos científicos** de referência salvos e vinculados
2. **Conteúdo clínico completo** em português brasileiro
3. **Campos preenchidos:** clinical_relevance, patient_explanation, conduct
4. **Data de revisão:** 2026-01-29 00:02:08

---

## 📚 Artigos Científicos Vinculados

### 1. The clinical evaluation of patients with subclinical hyperthyroidism and free T3 toxicosis
- **Autores:** Figge J, Leinung M, Goodman AD et al.
- **Journal:** American Journal of Medicine
- **Ano:** 1994
- **PMID:** 8154510
- **Resumo:** Estudo clássico sobre a avaliação de pacientes com hipertireoidismo subclínico e T3 toxicosis, estabelecendo critérios diagnósticos e padrões de interpretação para T3 livre em contextos de função tireoidiana limítrofe.

### 2. Clinical usage recommendations and analytic performance goals for total and free triiodothyronine measurements
- **Autores:** Demers LM, Spencer CA
- **Journal:** Clinical Chemistry
- **Ano:** 1995
- **PMID:** 8565219
- **Resumo:** Recomendações para uso clínico e objetivos de desempenho analítico para medições de T3 total e livre, incluindo padronização de métodos laboratoriais e valores de referência.

### 3. To test or not to test? Clinical utility and considerations for triiodothyronine (T3) testing
- **Autores:** Association for Diagnostics & Laboratory Medicine
- **Journal:** ADLMConnect
- **Ano:** 2024
- **Resumo:** Revisão contemporânea sobre utilidade clínica do teste de T3, incluindo indicações precisas, interpretação em diferentes contextos clínicos e frequência de T3 toxicosis (1.6% em pacientes com TSH suprimido).

---

## 📝 Conteúdo Clínico Gerado

### Clinical Relevance (1.861 caracteres)
Conteúdo técnico médico sobre:
- Definição e fisiologia do T3 livre
- Diferença entre T4 e T3, conversão periférica
- Indicações para dosagem (T3 toxicosis)
- Interpretação de valores elevados e reduzidos
- Valores de referência (2.0-4.4 pg/mL)
- Frequência de T3 toxicosis (1.6% dos casos com TSH suprimido)
- Condições que afetam a conversão periférica

### Patient Explanation (1.365 caracteres)
Explicação acessível incluindo:
- O que é T3 livre em linguagem simples
- Função do hormônio (metabolismo)
- Sintomas de valores alterados
- Valores normais de referência
- Importância da interpretação contextual

### Conduct (1.568 caracteres)
Protocolo clínico detalhado:
- **T3 elevado:** Confirmação com TSH/T4, anticorpos, cintilografia, USG, encaminhamento
- **T3 reduzido:** Investigação de doença aguda, revisão de medicações, diagnóstico diferencial
- **Reavaliação:** Prazos para seguimento (4-6 semanas em tratamento, 8-12 semanas em limítrofe)

---

## 🔍 Validação dos Dados

### Tamanho dos Campos
```
clinical_relevance:    1.861 caracteres ✓
patient_explanation:   1.365 caracteres ✓
conduct:               1.568 caracteres ✓
```

### Relações no Banco
```
Total de artigos vinculados: 12
(3 novos + 9 pré-existentes das aulas MFI)
```

### Timestamp
```
last_review: 2026-01-29 00:02:08.559192
updated_at:  2026-01-29 00:02:08.559192
```

---

## 🎯 Informações Técnicas

### Valores de Referência
- **Normal:** 2.0 - 4.4 pg/mL
- **Método:** Varia por laboratório

### Indicações Clínicas
1. TSH suprimido com T4 livre normal (suspeita de T3 toxicosis)
2. Sintomas de hipertireoidismo com testes iniciais limítrofes
3. Monitoramento de tratamento de hipertireoidismo
4. Avaliação de conversão periférica de T4

### Condições Associadas
- **T3 elevado:** T3 toxicosis, doença de Graves, bócio multinodular tóxico, adenoma tóxico
- **T3 reduzido:** Hipotireoidismo, síndrome do doente eutireoideo, medicamentos (propranolol, amiodarona, corticoides)

---

## 🚀 Próximos Passos

Este item está **100% completo** e pronto para uso no sistema:

- ✅ Artigos científicos salvos
- ✅ Relações many-to-many criadas
- ✅ Conteúdo clínico em PT-BR
- ✅ Explicação para pacientes
- ✅ Protocolo de conduta clínica
- ✅ Data de revisão registrada

---

## 📁 Arquivos Gerados

1. `/home/user/plenya/scripts/enrich_t3_livre.py` - Script Python (não utilizado devido a limitações do container)
2. `/home/user/plenya/scripts/t3_livre_enrichment.sql` - SQL de enriquecimento
3. `/home/user/plenya/T3-LIVRE-ENRICHMENT-REPORT.md` - Este relatório

---

## 🔗 Fontes de Referência

Durante a pesquisa, foram consultadas as seguintes fontes:

- [Triiodothyronine: Reference Range - Medscape](https://emedicine.medscape.com/article/2089598-overview)
- [The clinical evaluation of patients with subclinical hyperthyroidism - PubMed](https://pubmed.ncbi.nlm.nih.gov/8154510/)
- [Clinical utility and considerations for T3 testing - ADLM](https://myadlm.org/science-and-research/scientific-shorts/2024/clinical-utility-and-considerations-for-triiodothyronine-testing)
- [Thyroid Function Tests - American Thyroid Association](https://www.thyroid.org/thyroid-function-tests/)

---

**Enriquecimento executado via Docker Compose**
**Banco de dados: PostgreSQL 17**
**Executado em: 2026-01-28 23:56 - 00:02 UTC**
