# VCM (MCV) - Exemplos de Uso no Frontend

Este documento mostra como o conteúdo enriquecido do VCM pode ser exibido no frontend web e mobile.

---

## 1. Card de Resultado Laboratorial

**Contexto:** Visualização rápida do resultado no dashboard do paciente

```typescript
// Componente: LabResultCard.tsx
interface LabResultCardProps {
  itemName: string;
  value: number;
  unit: string;
  referenceRange: [number, number];
  clinicalRelevance: string;
}

// Exemplo de dados retornados da API
{
  "itemName": "VCM (MCV)",
  "value": 72,
  "unit": "fL",
  "referenceRange": [80, 100],
  "status": "low", // calculado: value < referenceRange[0]
  "clinicalRelevance": "**VCM (Volume Corpuscular Médio) - Classificação de Anemias**\n\nO VCM é uma medida fundamental no hemograma..."
}
```

**UI Display:**

```
┌─────────────────────────────────────────────────┐
│ 🔴 VCM (MCV)                                    │
│                                                 │
│ 72 fL                                           │
│ ▼ Abaixo do normal (80-100 fL)                 │
│                                                 │
│ ⚠️ Anemia Microcítica                          │
│ Suas hemácias estão menores que o normal       │
│                                                 │
│ [Ver Explicação Completa] [Ver Conduta]        │
└─────────────────────────────────────────────────┘
```

---

## 2. Modal de Explicação para Paciente

**Contexto:** Paciente clica em "Ver Explicação Completa"

```typescript
// Componente: PatientExplanationModal.tsx
<Dialog>
  <DialogTitle>
    VCM (MCV) - O Que É?
  </DialogTitle>
  <DialogContent>
    {/* Renderizar patient_explanation com Markdown */}
    <ReactMarkdown>
      {scoreItem.patient_explanation}
    </ReactMarkdown>

    {/* Resultado atual do paciente */}
    <Alert severity="warning">
      <AlertTitle>Seu Resultado: 72 fL</AlertTitle>
      Suas hemácias estão menores que o normal (80-100 fL).
      Isso indica anemia microcítica.
    </Alert>

    {/* Call to action */}
    <Button onClick={scheduleAppointment}>
      Agendar Consulta com Médico
    </Button>
  </DialogContent>
</Dialog>
```

**Visualização:**

```
╔═══════════════════════════════════════════════════╗
║ VCM (MCV) - O Que É?                       [X]    ║
╠═══════════════════════════════════════════════════╣
║                                                   ║
║ O VCM (Volume Corpuscular Médio) mede o tamanho  ║
║ médio das suas hemácias (glóbulos vermelhos).    ║
║                                                   ║
║ Valores Normais:                                  ║
║ • Normal: 80 a 100 fL                            ║
║ • Abaixo de 80: hemácias menores                 ║
║ • Acima de 100: hemácias maiores                 ║
║                                                   ║
║ ┌───────────────────────────────────────────┐    ║
║ │ ⚠️ Seu Resultado: 72 fL                  │    ║
║ │                                           │    ║
║ │ Suas hemácias estão menores que o normal │    ║
║ │ (80-100 fL). Isso indica anemia          │    ║
║ │ microcítica.                              │    ║
║ └───────────────────────────────────────────┘    ║
║                                                   ║
║ Por Que o Tamanho Importa?                       ║
║                                                   ║
║ Hemácias Pequenas (VCM baixo):                   ║
║ Geralmente significa falta de ferro...           ║
║                                                   ║
║ [Artigos Científicos] [Agendar Consulta]         ║
╚═══════════════════════════════════════════════════╝
```

---

## 3. Painel Médico - Conduta Clínica

**Contexto:** Médico visualiza resultado e acessa protocolo de conduta

```typescript
// Componente: ClinicalConductPanel.tsx
interface ConductPanelProps {
  scoreItem: ScoreItem;
  patientResult: LabResult;
}

// Lógica de determinação automática de seção relevante
function getRelevantConduct(vcmValue: number, conduct: string): string {
  if (vcmValue < 80) {
    // Extrair seção "ANEMIA MICROCÍTICA"
    return extractSection(conduct, "ANEMIA MICROCÍTICA");
  } else if (vcmValue > 100) {
    // Extrair seção "ANEMIA MACROCÍTICA"
    return extractSection(conduct, "ANEMIA MACROCÍTICA");
  } else {
    // Extrair seção "ANEMIA NORMOCÍTICA"
    return extractSection(conduct, "ANEMIA NORMOCÍTICA");
  }
}
```

**UI Display para VCM = 72 fL:**

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ 🩺 Conduta Clínica - Anemia Microcítica     ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

📋 INVESTIGAÇÃO INICIAL:
☐ Ferritina sérica (< 30 ng/mL = deficiência ferro)
☐ Saturação de transferrina (< 20% = deficiência)
☐ Proteína C reativa (avaliar inflamação)
☐ Hemoglobina A2 e F (talassemia)

[+ Solicitar Exames]

─────────────────────────────────────────────────

💊 CONDUTA SUGERIDA:

Se deficiência de ferro confirmada:
• Sulfato ferroso 300mg VO 1-2x/dia
• Investigar causa (sangramento GI, dieta)
• Reavaliar hemograma em 4-8 semanas

[+ Prescrever Tratamento]

─────────────────────────────────────────────────

📅 SEGUIMENTO:
Semana 4-8: reavaliar hemograma
Expectativa: ↑ Hb > 1 g/dL

[Agendar Follow-up]

─────────────────────────────────────────────────

📚 ARTIGOS CIENTÍFICOS (3)
• Mean Corpuscular Volume - StatPearls (2024)
• Anemia - StatPearls (2024)
• Normal and Abnormal CBC - StatPearls (2024)

[Ver Referências]
```

---

## 4. Widget de Timeline de Evolução

**Contexto:** Visualizar evolução do VCM ao longo do tempo com contexto clínico

```typescript
// Componente: VCMTrendChart.tsx
<Card>
  <CardHeader>
    <CardTitle>Evolução do VCM</CardTitle>
  </CardHeader>
  <CardContent>
    {/* Gráfico de linha mostrando evolução */}
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={vcmHistory}>
        <Line dataKey="value" stroke="#3b82f6" />
        {/* Faixa de referência 80-100 */}
        <ReferenceLine y={80} stroke="green" strokeDasharray="3 3" />
        <ReferenceLine y={100} stroke="green" strokeDasharray="3 3" />
        {/* Área de macrocitose marcada (>110) */}
        <ReferenceArea y1={110} y2={200} fill="red" fillOpacity={0.1} />
      </LineChart>
    </ResponsiveContainer>

    {/* Contexto clínico baseado na tendência */}
    <Alert>
      <TrendingUpIcon />
      <AlertDescription>
        Melhora progressiva após início de sulfato ferroso.
        Reavaliar em 4 semanas conforme protocolo.
      </AlertDescription>
    </Alert>
  </CardContent>
</Card>
```

**Visualização:**

```
┌─────────────────────────────────────────────────┐
│ Evolução do VCM                                 │
├─────────────────────────────────────────────────┤
│ fL                                              │
│ 100 ┬─────────────────────────────────────     │
│     │          ┌───── Faixa Normal ─────┐      │
│  90 ┤- - - - - - - - - - - - - - - - - -       │
│  80 ┤- - - - - - - - - - - - - - - - - -       │
│     │                        ╱                  │
│  70 ┤                    ╱                      │
│     │                 ╱                         │
│  60 ┤              ●                            │
│     └─┬──────┬──────┬──────┬──────┬──────      │
│      Jan    Fev    Mar    Abr    Mai           │
│                                                 │
│ ┌───────────────────────────────────────────┐  │
│ │ 📈 Melhora progressiva após início de     │  │
│ │ sulfato ferroso. Reavaliar em 4 semanas. │  │
│ └───────────────────────────────────────────┘  │
└─────────────────────────────────────────────────┘
```

---

## 5. Notificação Push Mobile

**Contexto:** Resultado de exame disponível com resumo inteligente

```typescript
// Serviço de notificação
interface LabResultNotification {
  title: string;
  body: string;
  data: {
    itemId: string;
    value: number;
    status: "normal" | "low" | "high";
    patientExplanationPreview: string;
  };
}

// Exemplo de notificação para VCM baixo
{
  "title": "Novo resultado: VCM (MCV)",
  "body": "72 fL - Abaixo do normal. Pode indicar falta de ferro.",
  "data": {
    "itemId": "a14322a8-07d5-480c-9131-cfdd3f0b7c21",
    "value": 72,
    "status": "low",
    "patientExplanationPreview": "O VCM mede o tamanho médio das suas hemácias. Valor baixo geralmente significa falta de ferro..."
  }
}
```

**Visualização Mobile:**

```
╔══════════════════════════════════╗
║ 🔔 Novo Resultado de Exame       ║
╠══════════════════════════════════╣
║                                  ║
║ VCM (MCV)                        ║
║ 72 fL                            ║
║ ▼ Abaixo do normal               ║
║                                  ║
║ Pode indicar falta de ferro.    ║
║ Seu médico já foi notificado.   ║
║                                  ║
║ [Ver Detalhes]   [Dispensar]    ║
╚══════════════════════════════════╝
```

---

## 6. Comparação de Resultados (Antes/Depois)

**Contexto:** Médico avalia eficácia do tratamento

```typescript
// Componente: BeforeAfterComparison.tsx
<ComparisonView>
  <Column title="Antes do Tratamento">
    <LabValue value={72} unit="fL" status="low" />
    <Badge color="red">Microcítica</Badge>
    <Text>Ferritina: 15 ng/mL</Text>
    <Text>Data: 15/01/2026</Text>
  </Column>

  <Divider>
    <ArrowRightIcon />
    <Text>8 semanas de sulfato ferroso</Text>
  </Divider>

  <Column title="Depois do Tratamento">
    <LabValue value={86} unit="fL" status="normal" />
    <Badge color="green">Normalizado</Badge>
    <Text>Ferritina: 68 ng/mL</Text>
    <Text>Data: 12/03/2026</Text>
  </Column>
</ComparisonView>

<Alert color="success">
  ✅ Resposta adequada ao tratamento
  (aumento VCM > 10 fL + normalização ferritina)
</Alert>
```

**Visualização:**

```
┌──────────────┐          ┌──────────────┐
│ ANTES        │   ───→   │ DEPOIS       │
├──────────────┤   8sem   ├──────────────┤
│              │          │              │
│  🔴 72 fL    │          │  🟢 86 fL    │
│  Microcítica │          │  Normal      │
│              │          │              │
│ Ferritina:   │          │ Ferritina:   │
│  15 ng/mL    │          │  68 ng/mL    │
│              │          │              │
│ 15/01/2026   │          │ 12/03/2026   │
└──────────────┘          └──────────────┘

┌────────────────────────────────────────┐
│ ✅ Resposta adequada ao tratamento     │
│ (aumento VCM > 10 fL + normalização)   │
└────────────────────────────────────────┘
```

---

## 7. Busca por Artigos Científicos

**Contexto:** Médico ou paciente acessa referências científicas vinculadas

```typescript
// Componente: RelatedArticles.tsx
<ArticlesList>
  {scoreItem.articles.map(article => (
    <ArticleCard key={article.id}>
      <ArticleTitle>{article.title}</ArticleTitle>
      <ArticleMeta>
        <Journal>{article.journal}</Journal>
        <Year>{article.publishDate.getFullYear()}</Year>
        <Badge>{article.articleType}</Badge>
      </ArticleMeta>
      <ArticleAbstract>
        {article.abstract.substring(0, 200)}...
      </ArticleAbstract>
      <ArticleActions>
        <Button variant="outline" onClick={() => window.open(article.originalLink)}>
          Ver no PubMed
        </Button>
        <Button variant="ghost">
          <BookmarkIcon />
        </Button>
      </ArticleActions>
    </ArticleCard>
  ))}
</ArticlesList>
```

**Visualização:**

```
┌─────────────────────────────────────────────────┐
│ 📚 Artigos Científicos Relacionados (3)         │
├─────────────────────────────────────────────────┤
│                                                 │
│ ┌─────────────────────────────────────────┐    │
│ │ Mean Corpuscular Volume - StatPearls    │    │
│ │ NCBI Bookshelf · 2024 · Review          │    │
│ │                                         │    │
│ │ Revisão abrangente sobre VCM como      │    │
│ │ medida crítica para identificar a      │    │
│ │ causa subjacente de anemia...          │    │
│ │                                         │    │
│ │ [Ver no PubMed] [🔖]                   │    │
│ └─────────────────────────────────────────┘    │
│                                                 │
│ ┌─────────────────────────────────────────┐    │
│ │ Anemia - StatPearls                     │    │
│ │ NCBI Bookshelf · 2024 · Review          │    │
│ │                                         │    │
│ │ Classificação completa de anemia       │    │
│ │ baseada em VCM. Detalha                │    │
│ │ fisiopatologia...                      │    │
│ │                                         │    │
│ │ [Ver no PubMed] [🔖]                   │    │
│ └─────────────────────────────────────────┘    │
│                                                 │
│ [Ver Todos os Artigos]                          │
└─────────────────────────────────────────────────┘
```

---

## 8. API Endpoints para Frontend

**Exemplo de chamadas à API:**

```typescript
// 1. Buscar score_item com conteúdo enriquecido
GET /api/v1/score-items/a14322a8-07d5-480c-9131-cfdd3f0b7c21
Response:
{
  "id": "a14322a8-07d5-480c-9131-cfdd3f0b7c21",
  "name": "VCM (MCV)",
  "clinical_relevance": "**VCM (Volume Corpuscular Médio)...",
  "patient_explanation": "**VCM - O Que É e Por Que É Importante**...",
  "conduct": "**Conduta Clínica Baseada em VCM**...",
  "updated_at": "2026-01-28T16:38:01.177564Z"
}

// 2. Buscar artigos vinculados ao item
GET /api/v1/score-items/a14322a8-07d5-480c-9131-cfdd3f0b7c21/articles
Response:
{
  "articles": [
    {
      "id": "818ca563-3e28-4fa2-bdd0-74f010f81a89",
      "title": "Mean Corpuscular Volume - StatPearls",
      "journal": "NCBI Bookshelf",
      "article_type": "review",
      "publish_date": "2024-01-01",
      "original_link": "https://www.ncbi.nlm.nih.gov/books/NBK545275/",
      "abstract": "Revisão abrangente sobre VCM..."
    },
    // ... outros artigos
  ]
}

// 3. Buscar resultado do paciente com interpretação
GET /api/v1/patients/:patientId/lab-results?item_id=a14322a8-07d5-480c-9131-cfdd3f0b7c21
Response:
{
  "result": {
    "value": 72,
    "unit": "fL",
    "date": "2026-01-15T10:00:00Z",
    "status": "low", // calculado automaticamente
    "interpretation": "Anemia microcítica",
    "score_item": {
      "name": "VCM (MCV)",
      "patient_explanation": "...",
      "reference_range": [80, 100]
    }
  }
}
```

---

## 9. Inteligência de Negócio

**Exemplo: Dashboard de Qualidade Assistencial**

```typescript
// Métricas baseadas em VCM enriquecido
interface QualityMetrics {
  anemias_detected: number;
  microcytic_anemia_treatment_rate: number; // % que receberam sulfato ferroso
  follow_up_adherence: number; // % que retornaram em 4-8 semanas
  normalization_rate: number; // % que normalizaram VCM
}

// Query analítica
SELECT
  COUNT(*) FILTER (WHERE value < 80) as microcytic_cases,
  COUNT(*) FILTER (WHERE value BETWEEN 80 AND 100) as normocytic_cases,
  COUNT(*) FILTER (WHERE value > 100) as macrocytic_cases,
  AVG(CASE WHEN value < 80 THEN
    EXTRACT(EPOCH FROM (next_result_date - first_result_date)) / 86400
  END) as avg_days_to_follow_up
FROM lab_results
WHERE score_item_id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
  AND date >= NOW() - INTERVAL '1 year';
```

---

## 10. Componente Completo de Resultado de Exame

**Arquivo:** `LabResultDetailView.tsx`

```typescript
import { useState } from 'react';
import { Card, Tabs, Alert, Button } from '@/components/ui';
import { ReactMarkdown } from 'react-markdown';

interface LabResultDetailViewProps {
  result: {
    value: number;
    unit: string;
    date: Date;
    scoreItem: {
      id: string;
      name: string;
      clinical_relevance: string;
      patient_explanation: string;
      conduct: string;
      articles: Article[];
    };
    referenceRange: [number, number];
  };
  userRole: 'patient' | 'doctor';
}

export function LabResultDetailView({ result, userRole }: LabResultDetailViewProps) {
  const [activeTab, setActiveTab] = useState('explanation');

  const getStatus = () => {
    if (result.value < result.referenceRange[0]) return 'low';
    if (result.value > result.referenceRange[1]) return 'high';
    return 'normal';
  };

  const getClassification = () => {
    if (result.value < 80) return 'Anemia Microcítica';
    if (result.value > 100) return 'Macrocitose';
    return 'Normocítico';
  };

  return (
    <div className="space-y-6">
      {/* Cabeçalho com valor */}
      <Card>
        <div className="flex justify-between items-center">
          <div>
            <h2 className="text-2xl font-bold">{result.scoreItem.name}</h2>
            <p className="text-gray-500">
              {new Date(result.date).toLocaleDateString('pt-BR')}
            </p>
          </div>
          <div className="text-right">
            <div className="text-4xl font-bold">{result.value} {result.unit}</div>
            <div className="text-sm text-gray-500">
              Referência: {result.referenceRange[0]}-{result.referenceRange[1]} {result.unit}
            </div>
          </div>
        </div>

        {/* Status visual */}
        <Alert severity={getStatus() === 'normal' ? 'success' : 'warning'} className="mt-4">
          {getStatus() === 'normal'
            ? '✅ Dentro da faixa normal'
            : `⚠️ ${getClassification()}`}
        </Alert>
      </Card>

      {/* Tabs com conteúdo */}
      <Tabs value={activeTab} onValueChange={setActiveTab}>
        <TabsList>
          <TabsTrigger value="explanation">Explicação</TabsTrigger>
          {userRole === 'doctor' && (
            <>
              <TabsTrigger value="clinical">Relevância Clínica</TabsTrigger>
              <TabsTrigger value="conduct">Conduta</TabsTrigger>
            </>
          )}
          <TabsTrigger value="articles">Artigos ({result.scoreItem.articles.length})</TabsTrigger>
        </TabsList>

        <TabsContent value="explanation">
          <Card>
            <ReactMarkdown>{result.scoreItem.patient_explanation}</ReactMarkdown>
          </Card>
        </TabsContent>

        {userRole === 'doctor' && (
          <>
            <TabsContent value="clinical">
              <Card>
                <ReactMarkdown>{result.scoreItem.clinical_relevance}</ReactMarkdown>
              </Card>
            </TabsContent>

            <TabsContent value="conduct">
              <Card>
                <ReactMarkdown>{result.scoreItem.conduct}</ReactMarkdown>
                <div className="mt-6 space-x-4">
                  <Button>Solicitar Exames Complementares</Button>
                  <Button variant="outline">Prescrever Tratamento</Button>
                </div>
              </Card>
            </TabsContent>
          </>
        )}

        <TabsContent value="articles">
          <div className="space-y-4">
            {result.scoreItem.articles.map(article => (
              <Card key={article.id}>
                <h3 className="font-bold">{article.title}</h3>
                <p className="text-sm text-gray-500">
                  {article.journal} · {article.publishDate} · {article.articleType}
                </p>
                <p className="mt-2">{article.abstract}</p>
                <Button
                  variant="link"
                  onClick={() => window.open(article.originalLink, '_blank')}
                >
                  Ver artigo completo →
                </Button>
              </Card>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
```

---

## Conclusão

O conteúdo enriquecido do VCM permite criar experiências ricas tanto para pacientes quanto para médicos:

1. **Pacientes:** Compreendem seus resultados com linguagem acessível
2. **Médicos:** Acessam protocolos baseados em evidências imediatamente
3. **Sistema:** Rastreia fontes científicas e mantém qualidade assistencial

Todo o conteúdo está pronto para consumo via API REST e pode ser renderizado em web ou mobile com Markdown.
