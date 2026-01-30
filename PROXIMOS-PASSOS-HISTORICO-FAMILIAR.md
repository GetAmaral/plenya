# Próximos Passos - Histórico Familiar de Doenças

**Status Atual:** 24/24 items enriquecidos (100%)
**Campos Preenchidos:** `clinicalRelevance`, `conduct`
**Data:** 27 de Janeiro de 2026

---

## Fase 1: Completar Campos Existentes (1-2 horas)

### 1.1 PatientExplanation (Linguagem Acessível)
**Objetivo:** Criar versão simplificada para pacientes

**Conteúdo Sugerido:**
- ~200 palavras por item
- Linguagem 8ª série (sem jargão técnico)
- Estrutura:
  1. O que significa ter histórico familiar desta doença? (2-3 frases)
  2. Quanto isso aumenta meu risco? (números simples: "2-3 vezes maior")
  3. O que posso fazer para prevenir? (3-5 ações concretas)
  4. Mensagem positiva sobre modificabilidade

**Exemplo - Diabetes:**
```
Se você tem parentes próximos com diabetes, seu risco de desenvolver a doença é 2 a 3 vezes maior que a população geral. Mas isso NÃO significa que você terá diabetes - estudos mostram que você pode reduzir seu risco em até 70% com mudanças no estilo de vida.

O que você pode fazer:
- Manter peso saudável (perder 7% do peso se necessário)
- Exercitar-se 30 minutos por dia, 5 vezes por semana
- Comer mais vegetais, grãos integrais, evitar açúcar e bebidas açucaradas
- Fazer exames de glicemia e insulina anualmente a partir dos 25 anos

Lembre-se: seus genes não são seu destino. Suas escolhas diárias podem "desligar" genes que aumentariam o risco de diabetes.
```

**Script de Geração:**
```python
# Pode ser gerado com IA usando prompt específico baseado em clinicalRelevance+conduct
# Critérios: clareza, positividade, acionabilidade, sem medo
```

---

### 1.2 Codes Programáticos
**Objetivo:** Identificadores únicos para queries e integrações

**Padrão Sugerido:**
```
FAM_HIST_DIABETES
FAM_HIST_DYSLIPIDEMIA
FAM_HIST_CVD
FAM_HIST_RENAL
FAM_HIST_AUTOIMMUNE
FAM_HIST_VIRAL_CHRONIC
FAM_HIST_HYPERTENSION
FAM_HIST_OBESITY
```

**Script de Atualização:**
```sql
-- Atualizar codes para os 24 items
UPDATE score_items SET code = 'FAM_HIST_DIABETES'
WHERE name = 'Diabetes mellitus / pré-diabetes / resistência insulínica';

UPDATE score_items SET code = 'FAM_HIST_DYSLIPIDEMIA'
WHERE name = 'Dislipidemia';

-- ... repetir para todas as 8 doenças
```

**Utilidade:**
- Queries simplificadas: `WHERE code = 'FAM_HIST_DIABETES'`
- Integrações externas via API
- Exportação de dados estruturados

---

## Fase 2: Referências Científicas (2-3 horas)

### 2.1 Vincular Artigos Científicos
**Objetivo:** Adicionar 3-5 referências por doença

**Estrutura:**
Usar tabela `articles` + relacionamento many-to-many `article_score_items`

**Artigos Sugeridos por Doença:**

**Diabetes:**
1. Diabetes Prevention Program Research Group. NEJM 2002;346(6):393-403. DOI:10.1056/NEJMoa012512
2. Knowler WC et al. Lancet 2009;374(9702):1677-1686. DOI:10.1016/S0140-6736(09)61457-4
3. Ley SH et al. Diabetologia 2014;57(7):1276-1286. DOI:10.1007/s00125-014-3245-x

**Cardiovascular:**
1. Estruch R et al. (PREDIMED). NEJM 2018;378(25):e34. DOI:10.1056/NEJMoa1800389
2. Khera AV et al. Nature Genetics 2016;48(6):584-590. DOI:10.1038/ng.3552
3. Marenberg ME et al. Circulation 1994;89(1):215-220. DOI:10.1161/01.CIR.89.1.215

**Obesidade:**
1. Qi Q et al. AJCN 2012;96(5):1129-1136. DOI:10.3945/ajcn.112.041707
2. Locke AE et al. (GIANT). Nature 2015;518(7538):197-206. DOI:10.1038/nature14177
3. Wilding JPH et al. NEJM 2021;384(11):989-1002. DOI:10.1056/NEJMoa2032183

**Script de Inserção:**
```sql
-- Criar artigos
INSERT INTO articles (title, authors, journal, year, doi, category, summary)
VALUES
  ('Reduction in the Incidence of Type 2 Diabetes with Lifestyle Intervention',
   'Diabetes Prevention Program Research Group',
   'New England Journal of Medicine',
   2002,
   '10.1056/NEJMoa012512',
   'diabetes_prevention',
   'Estudo landmark demonstrando redução de 58% na incidência de diabetes tipo 2 com intervenção lifestyle (dieta + exercício) vs placebo.');

-- Vincular artigo a items
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
  (SELECT id FROM articles WHERE doi = '10.1056/NEJMoa012512'),
  id
FROM score_items
WHERE name = 'Diabetes mellitus / pré-diabetes / resistência insulínica';
```

---

### 2.2 Campo scientificReferences
**Alternativa:** Adicionar campo texto com referências formatadas

```sql
ALTER TABLE score_items ADD COLUMN scientific_references TEXT;
```

**Formato Sugerido (AMA style):**
```
1. Diabetes Prevention Program Research Group. Reduction in the incidence of type 2 diabetes with lifestyle intervention or metformin. N Engl J Med. 2002;346(6):393-403. doi:10.1056/NEJMoa012512

2. Knowler WC, Fowler SE, Hamman RF, et al. 10-year follow-up of diabetes incidence and weight loss in the Diabetes Prevention Program Outcomes Study. Lancet. 2009;374(9702):1677-1686. doi:10.1016/S0140-6736(09)61457-4

3. Ley SH, Hamdy O, Mohan V, Hu FB. Prevention and management of type 2 diabetes: dietary components and nutritional strategies. Diabetologia. 2014;57(7):1276-1286. doi:10.1007/s00125-014-3245-x
```

---

## Fase 3: Frontend - Visualização (3-5 horas)

### 3.1 Genograma Familiar Interativo
**Biblioteca:** React Flow ou D3.js

**Features:**
- Adicionar parentes com doenças (pais, irmãos, avós, tios, filhos)
- Ícones por doença (coração para CV, gota para diabetes, etc)
- Cálculo automático de risco baseado em:
  - Número de parentes afetados
  - Grau de parentesco
  - Idade de diagnóstico (precoce = maior risco)

**Exemplo de Interface:**
```
         Avô     Avó
          |       |
      ┌───┴───┬───┴───┐
      │       │       │
    Pai*   Tio    Tia
      │
   Paciente
```

*Diabetes aos 50 anos → Risco aumentado 2-3x

---

### 3.2 Cards de Doença Familiar
**Layout:**

```jsx
<Card>
  <CardHeader>
    <Icon type="diabetes" /> Diabetes Tipo 2
    <Badge variant="high-risk">Risco Alto</Badge>
  </CardHeader>

  <CardContent>
    <Section title="Seu Histórico">
      <FamilyTree>
        - Pai: Diagnosticado aos 52 anos
        - Avô paterno: Diagnosticado aos 60 anos
      </FamilyTree>
      <RiskCalculator>
        Risco estimado: 5-6x maior que população geral
      </RiskCalculator>
    </Section>

    <Section title="O que Isso Significa">
      <PatientExplanation>
        {item.patientExplanation}
      </PatientExplanation>
    </Section>

    <Section title="Recomendações Personalizadas">
      <Recommendations>
        ✓ Iniciar rastreamento aos 42 anos (10 anos antes do pai)
        ✓ HbA1c + glicemia + insulina anualmente
        ✓ Considerar teste genético TCF7L2/FTO
      </Recommendations>
    </Section>

    <Accordion title="Detalhes Científicos" (expandir)>
      <ClinicalRelevance>
        {item.clinicalRelevance}
      </ClinicalRelevance>
    </Accordion>

    <Accordion title="Protocolo Completo" (expandir)>
      <Conduct>
        {item.conduct}
      </Conduct>
    </Accordion>

    <ScientificReferences>
      {item.articles.map(...)}
    </ScientificReferences>
  </CardContent>
</Card>
```

---

### 3.3 Calculadora de Risco Personalizado
**Inputs:**
- Idade atual
- Doenças familiares (checkbox)
  - Por doença: número de parentes, grau, idade diagnóstico
- Fatores de estilo de vida (IMC, exercício, dieta, tabagismo)

**Output:**
- Risco estimado comparado à população geral
- Idade sugerida para iniciar rastreamento
- Exames recomendados com frequência
- Intervenções prioritárias

**Exemplo:**
```
╔══════════════════════════════════════════╗
║  SEU RISCO PERSONALIZADO - DIABETES      ║
╠══════════════════════════════════════════╣
║  Risco genético:        5-6x             ║
║  Risco atual (com IMC): 8x               ║
║                                          ║
║  Iniciar screening:     Imediatamente    ║
║  Frequência:            Anual            ║
║                                          ║
║  SE você perder 7% peso + exercitar:     ║
║  Risco reduz para:      2-3x             ║
║  Redução absoluta:      70%              ║
╚══════════════════════════════════════════╝
```

---

## Fase 4: Alertas e Automações (2-3 horas)

### 4.1 Sistema de Alertas
**Regras:**

```typescript
interface FamilyHistoryAlert {
  condition: string;
  patientAge: number;
  startScreeningAge: number;
  tests: string[];
  frequency: string;
  urgency: 'immediate' | 'soon' | 'scheduled';
}

function generateAlerts(patient: Patient): FamilyHistoryAlert[] {
  const alerts: FamilyHistoryAlert[] = [];

  // Exemplo: Pai com IAM aos 52, paciente tem 42
  if (hasFamilyHistory('cardiovascular') && patient.age >= 42) {
    alerts.push({
      condition: 'Doença Cardiovascular',
      patientAge: patient.age,
      startScreeningAge: 42, // 10 anos antes do pai
      tests: ['Perfil lipídico avançado', 'hs-CRP', 'Homocisteína', 'CAC score'],
      frequency: 'Anual',
      urgency: 'immediate'
    });
  }

  return alerts;
}
```

**UI:**
```
🔔 Alertas de Rastreamento

┌─────────────────────────────────────────┐
│ ⚠️  AÇÃO NECESSÁRIA                     │
│                                         │
│ Doença Cardiovascular                   │
│ Você atingiu a idade recomendada para   │
│ iniciar rastreamento (42 anos).         │
│                                         │
│ ✓ Solicitar: CAC Score + Perfil Lipídico│
│ 📅 Agendar consulta cardiologista       │
└─────────────────────────────────────────┘
```

---

### 4.2 Templates de Solicitação de Exames
**Auto-preencher pedidos médicos:**

```typescript
function generateLabRequest(familyHistory: FamilyHistory[]): LabRequest {
  let tests = ['Hemograma', 'Função Renal', 'Função Hepática']; // Baseline

  if (familyHistory.includes('diabetes')) {
    tests.push(
      'Glicemia de Jejum',
      'HbA1c',
      'Insulina Basal',
      'Peptídeo-C',
      'TOTG com insulina 0-30-60-120min'
    );
  }

  if (familyHistory.includes('cardiovascular')) {
    tests.push(
      'Perfil Lipídico Avançado (LDL, HDL, VLDL, Triglicerídeos, não-HDL, Apo B, Apo A1)',
      'Lp(a)',
      'hs-CRP',
      'Homocisteína',
      'Fibrinogênio'
    );
  }

  return {
    tests,
    justification: generateJustification(familyHistory),
    priority: calculatePriority(familyHistory)
  };
}
```

**Justificativa Automática:**
```
JUSTIFICATIVA CLÍNICA:
Paciente com histórico familiar positivo para:
- Diabetes Tipo 2 (pai aos 52 anos, avô paterno aos 60 anos)
- Doença Cardiovascular (pai IAM aos 54 anos)

Risco estimado: 5-6x para DM2, 3-4x para DCV
Rastreamento precoce indicado conforme protocolo de medicina preventiva personalizada.

Referências:
- Diabetes Prevention Program (NEJM 2002;346:393-403)
- AHA/ACC Guidelines on Primary Prevention (Circulation 2019)
```

---

### 4.3 Relatório PDF para Paciente
**Conteúdo:**

```markdown
# SEU HISTÓRICO FAMILIAR E RECOMENDAÇÕES PERSONALIZADAS

## Resumo do Seu Perfil
- Idade: 42 anos
- Histórico Familiar: Diabetes Tipo 2, Doença Cardiovascular

## Genograma da Família
[Árvore visual com ícones de doenças]

## Doenças Identificadas

### 1. Diabetes Tipo 2
**Seu Histórico:**
- Pai: Diagnosticado aos 52 anos
- Avô paterno: Diagnosticado aos 60 anos

**O Que Isso Significa:**
[patientExplanation do item]

**Seu Risco:** 5-6x maior que população geral

**O Que Fazer:**
✓ Exames anuais: HbA1c, glicemia, insulina
✓ Perder 7% do peso (se necessário)
✓ Exercitar 150min/semana
✓ Dieta mediterrânea baixo IG

---

## Calendário de Rastreamento Personalizado

| Exame | Frequência | Próxima Data |
|-------|-----------|--------------|
| HbA1c + Glicemia + Insulina | Anual | 15/02/2026 |
| Perfil Lipídico Avançado | Anual | 15/02/2026 |
| CAC Score | Único aos 45 anos | 2029 |

---

## Recursos Adicionais
- Vídeos educativos sobre prevenção
- Receitas mediterrâneas
- Apps de exercício recomendados

---

Gerado em: 27/01/2026
Sistema: Plenya EMR
```

---

## Fase 5: Integração com Genética (futuro)

### 5.1 Import de Testes Genéticos
**Providers:**
- 23andMe
- AncestryDNA
- Genera (Brasil)
- meuDNA (Brasil)

**Workflow:**
1. Paciente faz teste genético comercial
2. Faz download do arquivo raw (23andMe raw data)
3. Upload no sistema Plenya
4. Parser identifica SNPs relevantes:
   - rs7903146 (TCF7L2 - diabetes)
   - rs9939609 (FTO - obesidade)
   - rs1333049 (9p21 - cardiovascular)
   - APOE genotype (ε2/ε3/ε4)
   - rs662799 (APOA5 - triglicerídeos)
5. Gera escore de risco genético poligênico

**Output:**
```
🧬 SEU PERFIL GENÉTICO

TCF7L2 rs7903146: CT (1 cópia de risco)
→ Risco diabetes aumentado em 1,4x
→ Combinado com histórico familiar: 7-8x

FTO rs9939609: AA (2 cópias de risco)
→ Tendência a IMC 3-4 kg/m² maior
→ ATENÇÃO: Atividade física reduz este efeito em 40%

APOE: ε3/ε4 (1 cópia ε4)
→ Maior resposta do LDL a gorduras saturadas
→ RECOMENDAÇÃO: Reduzir saturadas para <5% calorias

Escore Poligênico Global: 75/100 (quartil superior)
```

---

### 5.2 Recomendações Ajustadas por Genótipo
**Personalização Nutricional:**

```typescript
function personalizeNutrition(genotype: Genotype): NutritionPlan {
  let plan: NutritionPlan = baseMediterraneanDiet();

  // APOE ε4: restringir saturadas
  if (genotype.APOE.includes('ε4')) {
    plan.saturatedFat.max = '5%'; // vs 7% padrão
    plan.notes.push('Seu genótipo APOE ε4 responde mais ao colesterol dietário');
  }

  // FTO alta: enfatizar exercício
  if (genotype.FTO === 'AA') {
    plan.exercise.importance = 'CRÍTICO';
    plan.notes.push('Exercício reduz 40% do efeito genético do FTO');
  }

  // TCF7L2 alta: low-carb pode ser superior
  if (genotype.TCF7L2 === 'TT') {
    plan.carbs.type = 'low-carb';
    plan.notes.push('Seu genótipo pode responder melhor a dieta low-carb');
  }

  return plan;
}
```

---

## Priorização Sugerida

### Curto Prazo (1-2 semanas)
1. ✓ PatientExplanation para todos os 24 items
2. ✓ Codes programáticos
3. ✓ 3 referências científicas por doença

### Médio Prazo (1 mês)
4. ✓ Frontend: Cards de doença familiar com expansão
5. ✓ Calculadora de risco personalizado
6. ✓ Alertas de rastreamento

### Longo Prazo (3-6 meses)
7. ○ Genograma interativo
8. ○ Templates de solicitação automática
9. ○ Relatórios PDF personalizados
10. ○ Integração com testes genéticos

---

## Métricas de Sucesso

### Fase 1 (Completude)
- [ ] 100% items com patientExplanation
- [ ] 100% items com codes
- [ ] 100% items com ≥3 referências

### Fase 2 (Engajamento)
- [ ] 80% pacientes com histórico familiar preenchido
- [ ] 60% pacientes visualizam explicações
- [ ] 40% pacientes expandem detalhes científicos

### Fase 3 (Ação)
- [ ] 70% pacientes com rastreamento agendado
- [ ] 50% pacientes iniciam intervenções lifestyle
- [ ] 30% redução em diagnósticos tardios

---

**Última Atualização:** 27 de Janeiro de 2026
**Responsável:** Equipe de Produto Plenya
**Revisão:** Trimestral
