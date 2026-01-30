# Implementação de Novos Campos - Score Items e Levels

**Data:** 25 de Janeiro de 2026
**Status:** ✅ **COMPLETO**

---

## Resumo Executivo

Foram adicionados três novos campos essenciais aos modelos `ScoreItem` e `ScoreLevel` para enriquecer o sistema de estratificação de risco com informações clínicas, educativas e orientações de conduta.

### Novos Campos Implementados

| Campo | Tipo | Finalidade |
|-------|------|------------|
| **clinical_relevance** | TEXT | Explicação técnica da relevância clínica para profissionais de saúde |
| **patient_explanation** | TEXT | Explicação em linguagem simples e acessível para pacientes |
| **conduct** | TEXT | Orientações de conduta clínica recomendada |

---

## Alterações no Banco de Dados

### Tabela: `score_items`

**Colunas adicionadas:**
```sql
clinical_relevance  TEXT
patient_explanation TEXT
conduct             TEXT
```

**Status:** ✅ Criadas automaticamente via GORM auto-migration

### Tabela: `score_levels`

**Colunas adicionadas:**
```sql
clinical_relevance  TEXT
patient_explanation TEXT
conduct             TEXT
```

**Status:** ✅ Criadas manualmente via ALTER TABLE

---

## Alterações no Código

### 1. `/apps/api/internal/models/score_item.go`

Adicionados três novos campos após `UnitConversion`:

```go
// Relevância clínica - explicação técnica para profissionais de saúde
// @example Valores baixos de hemoglobina indicam anemia...
ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

// Explicação para o paciente - linguagem simples e acessível
// @example Hemoglobina é a proteína que transporta oxigênio...
PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

// Conduta clínica recomendada
// @example Investigar causa da anemia...
Conduct *string `gorm:"type:text" json:"conduct,omitempty"`
```

### 2. `/apps/api/internal/models/score_level.go`

Adicionados os mesmos três campos após `Operator`:

```go
// Relevância clínica - explicação técnica para profissionais de saúde
// @example FEVE entre 55-70% indica função cardíaca normal...
ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

// Explicação para o paciente - linguagem simples e acessível
// @example Seu coração está bombeando sangue de forma eficiente...
PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

// Conduta clínica recomendada
// @example Manter acompanhamento regular...
Conduct *string `gorm:"type:text" json:"conduct,omitempty"`
```

---

## Exemplos de Uso - Dados Reais

### Score Item: NT-proBNP (<50 anos)

**ID:** `49c88f04-ab34-4d19-8b60-64765b6fc8f0`

```json
{
  "name": "NT-proBNP (<50 anos)",
  "clinicalRelevance": "NT-proBNP (N-terminal pro-B-type natriuretic peptide) é um biomarcador cardíaco que reflete o estresse da parede ventricular. Elevação indica sobrecarga de volume ou pressão no coração, sendo fundamental para diagnóstico e prognóstico de insuficiência cardíaca. Níveis baixos em pacientes sem sintomas têm alto valor preditivo negativo para IC.",

  "patientExplanation": "Este exame mede uma proteína que o coração libera quando está trabalhando com dificuldade. Valores normais indicam que seu coração está funcionando bem, sem sobrecarga.",

  "conduct": "Valores normais (<125 pg/mL em <50 anos): baixa probabilidade de IC, dispensar ecocardiograma se assintomático. Valores intermediários: avaliar sintomas, fatores de risco, considerar ecocardiograma. Valores elevados: investigar IC com ecocardiograma, ECG, raio-X de tórax. Seguir algoritmo ESC 2023 para diagnóstico de IC."
}
```

### Score Level: NT-proBNP >1800 pg/mL (Level 0 - Risco Alto)

**ID:** `b6c0866d-6b5b-4fc1-997f-ddc8f5ee6dd8`

```json
{
  "level": 0,
  "name": ">1800",
  "clinicalRelevance": "NT-proBNP >1800 pg/mL em pacientes <50 anos indica alta probabilidade de insuficiência cardíaca aguda ou descompensada. Sugere sobrecarga ventricular significativa, com necessidade de avaliação urgente. Diagnóstico diferencial inclui embolia pulmonar, cor pulmonale, sepse.",

  "patientExplanation": "Este resultado indica que seu coração pode estar sob grande pressão. É importante procurar avaliação médica urgente para investigar a causa e iniciar tratamento adequado.",

  "conduct": "URGENTE: Realizar ecocardiograma em <48h. ECG, raio-X de tórax, função renal, eletrólitos. Avaliar sinais/sintomas de IC descompensada (dispneia, edema, ortopneia). Considerar internação se sintomático. Iniciar diurético se congestão presente. Encaminhar ao cardiologista."
}
```

### Score Level: NT-proBNP <50 pg/mL (Level 5 - Baixo Risco)

**ID:** `2627cf54-7494-4841-a438-d7de334a5d65`

```json
{
  "level": 5,
  "name": "<50",
  "clinicalRelevance": "NT-proBNP <50 pg/mL em pacientes <50 anos tem valor preditivo negativo >95% para insuficiência cardíaca. Este nível baixo indica ausência de estresse ventricular significativo, tornando IC altamente improvável. Permite descartar IC em pacientes com dispneia de causa incerta.",

  "patientExplanation": "Este resultado é excelente! Significa que a probabilidade de problemas no bombeamento do seu coração é muito baixa. Seu coração está funcionando bem.",

  "conduct": "IC improvável. Dispensar ecocardiograma se paciente assintomático. Investigar outras causas de sintomas se presentes (anemia, doença pulmonar, ansiedade, descondicionamento físico). Manter acompanhamento de fatores de risco cardiovascular. Retorno conforme rotina."
}
```

---

## Estatísticas do Banco de Dados

### Contadores

| Entidade | Total |
|----------|-------|
| **Score Items** | 772 |
| **Score Levels** | 3.028 |
| **Items com novos campos preenchidos** | 1 (exemplo) |
| **Levels com novos campos preenchidos** | 2 (exemplos) |

### Campos a Preencher

- **Score Items:** 771 itens aguardando preenchimento dos novos campos
- **Score Levels:** 3.026 níveis aguardando preenchimento dos novos campos

---

## Casos de Uso

### 1. Interface do Médico

Quando o médico visualiza um resultado de exame:

```
┌─────────────────────────────────────────┐
│ NT-proBNP: 1.950 pg/mL                 │
│ Nível de Risco: 0 (Alto)               │
│                                         │
│ 📊 RELEVÂNCIA CLÍNICA:                 │
│ NT-proBNP >1800 pg/mL indica alta      │
│ probabilidade de IC aguda...           │
│                                         │
│ 🏥 CONDUTA RECOMENDADA:                │
│ URGENTE: Realizar eco em <48h...       │
└─────────────────────────────────────────┘
```

### 2. Interface do Paciente

Quando o paciente visualiza seu resultado:

```
┌─────────────────────────────────────────┐
│ Seu Resultado: NT-proBNP                │
│                                         │
│ 💬 O QUE SIGNIFICA:                    │
│ Este resultado indica que seu coração  │
│ pode estar sob grande pressão...       │
│                                         │
│ 🩺 PRÓXIMOS PASSOS:                    │
│ É importante procurar avaliação        │
│ médica urgente...                      │
└─────────────────────────────────────────┘
```

### 3. Geração de Relatórios

O sistema pode gerar laudos automáticos com:
- Interpretação técnica (clinical_relevance)
- Orientações ao paciente (patient_explanation)
- Recomendações de seguimento (conduct)

---

## Benefícios Implementados

### Para Médicos
✅ **Padronização:** Orientações clínicas consistentes baseadas em evidências
✅ **Agilidade:** Acesso rápido a condutas recomendadas sem consultar guidelines
✅ **Educação:** Contexto clínico para cada parâmetro e nível de risco
✅ **Segurança:** Alertas claros sobre situações que requerem ação urgente

### Para Pacientes
✅ **Compreensão:** Explicações em linguagem simples sobre seus exames
✅ **Transparência:** Entendimento dos motivos das condutas médicas
✅ **Empoderamento:** Maior participação nas decisões sobre sua saúde
✅ **Redução de ansiedade:** Clareza sobre o significado dos resultados

### Para o Sistema
✅ **Escalabilidade:** Campos reutilizáveis em toda a plataforma
✅ **Integridade:** Dados estruturados no banco de dados
✅ **Auditoria:** Registro das orientações fornecidas
✅ **Compliance:** Documentação das bases para decisões clínicas

---

## Integração com Articles (Evidências)

Os novos campos complementam a relação many-to-many com `articles`:

```
ScoreItem/ScoreLevel
│
├── clinical_relevance (explicação técnica)
├── patient_explanation (explicação leiga)
├── conduct (orientações)
│
└── articles[] (referências científicas que embasam as orientações)
```

**Exemplo:**
- **Conduct:** "Seguir algoritmo ESC 2023 para diagnóstico de IC"
- **Article:** ESC Heart Failure Algorithms 2023 (PDF linkado)

---

## Próximos Passos

### Curto Prazo (Urgente)
1. ✅ Campos criados no banco de dados
2. ✅ Modelos Go atualizados
3. ⏳ **Criar endpoints de atualização em massa** para preencher campos
4. ⏳ **Implementar interface de edição** no frontend
5. ⏳ **Exibir novos campos** nas páginas de detalhes de score items/levels

### Médio Prazo
1. **Popular dados usando IA:** Usar LLM para gerar sugestões de preenchimento baseadas em:
   - Nome do parâmetro
   - Unidade de medida
   - Níveis de referência
   - Artigos científicos linkados

2. **Revisão médica:** Validar conteúdos gerados por IA com médicos especialistas

3. **Tradução multilíngue:** Suporte para português, inglês, espanhol

### Longo Prazo
1. **IA generativa em tempo real:** Personalizar explicações baseadas em:
   - Perfil do paciente (idade, escolaridade, condições prévias)
   - Contexto clínico específico
   - Preferências de comunicação

2. **Versionamento:** Manter histórico de alterações nas orientações

3. **Feedback loop:** Coletar feedback de médicos e pacientes sobre qualidade das explicações

---

## Comandos Úteis

### Verificar estrutura das tabelas
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_items"
docker compose exec db psql -U plenya_user -d plenya_db -c "\d score_levels"
```

### Contar registros com campos preenchidos
```sql
-- Score Items com campos preenchidos
SELECT COUNT(*) FROM score_items
WHERE clinical_relevance IS NOT NULL
   OR patient_explanation IS NOT NULL
   OR conduct IS NOT NULL;

-- Score Levels com campos preenchidos
SELECT COUNT(*) FROM score_levels
WHERE clinical_relevance IS NOT NULL
   OR patient_explanation IS NOT NULL
   OR conduct IS NOT NULL;
```

### Buscar items/levels vazios para preencher
```sql
-- Items sem clinical_relevance
SELECT id, name FROM score_items
WHERE clinical_relevance IS NULL
LIMIT 10;

-- Levels sem patient_explanation
SELECT id, level, name FROM score_levels
WHERE patient_explanation IS NULL
LIMIT 10;
```

---

## Arquivos Modificados

### Backend
- `/apps/api/internal/models/score_item.go` - Adicionados 3 campos
- `/apps/api/internal/models/score_level.go` - Adicionados 3 campos

### Banco de Dados
- `score_items` - 3 novas colunas (TEXT)
- `score_levels` - 3 novas colunas (TEXT)

### Documentação
- `/home/user/plenya/NEW-FIELDS-IMPLEMENTATION-REPORT.md` - Este arquivo

---

## Conclusão

A implementação dos campos `clinical_relevance`, `patient_explanation` e `conduct` representa um avanço significativo na capacidade do sistema Plenya EMR de fornecer:

1. **Orientação clínica baseada em evidências**
2. **Educação em saúde para pacientes**
3. **Padronização de condutas médicas**
4. **Rastreabilidade de decisões clínicas**

Com 772 score items e 3.028 score levels disponíveis, o próximo desafio é popular esses campos com conteúdo de alta qualidade, idealmente combinando:
- Inteligência artificial para geração inicial
- Revisão por especialistas médicos
- Integração com artigos científicos já cadastrados no sistema

---

**Status Final:** ✅ **IMPLEMENTAÇÃO COMPLETA**
**Próximo Sprint:** Popular campos com conteúdo clínico de qualidade

---

*Plenya EMR - Sistema de Prontuário Eletrônico Baseado em Evidências*
*Versão: 2026.01*
