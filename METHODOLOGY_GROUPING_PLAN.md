# Plano de Implementação: Agrupamento por Metodologia em Health Scores

**Data:** 2026-02-16
**Objetivo:** Reorganizar a visualização de scores do paciente por metodologia quando houver subscription ativa

---

## 📋 Resumo Executivo

### Hierarquia Atual
```
ScoreGroup → ScoreSubgroup → ScoreItem
```

### Nova Hierarquia (com subscription ativa)
```
Method → MethodLetter → MethodPillar → ScoreGroup → ScoreSubgroup → ScoreItem
```

### Decisões de Design

1. ✅ **Items em múltiplos pilares**: Aparecem em cada pilar E contam pontos para todos
2. ✅ **Items sem pilar**: Letter virtual "Outros" no final com pilar "Items não incluídos nos pilares do método"
3. ✅ **Toggle de visualização**: Botão para alternar entre "View Metodológica" ↔ "View Tradicional"
4. ✅ **Score total**: Mantém lógica atual (cada item conta 1x só no total geral)
5. ✅ **Filtro "Apenas calculados"**: Mantém tudo visível sempre
6. ✅ **Deep linking**: Auto-abrir toda hierarquia Method → Letter → Pillar → Group → Subgroup

---

## 🔧 FASE 1: Backend (~4-6h)

### 1.1. `/apps/api/internal/repository/score_snapshot_repository.go`

**Linha ~48** - Adicionar preloads para MethodPillars:

```go
.Preload("ItemResults.Item.MethodPillars", "deleted_at IS NULL", func(db *gorm.DB) *gorm.DB {
    return db.Order("\"order\" ASC")
})
.Preload("ItemResults.Item.MethodPillars.Letter", "deleted_at IS NULL")
.Preload("ItemResults.Item.MethodPillars.Letter.Method", "deleted_at IS NULL")
```

**Linha ~94-96** - Repetir os mesmos preloads para `GetLatestByPatientID`

---

### 1.2. `/apps/api/internal/repository/subscription_repository.go`

**Adicionar método novo** (final do arquivo):

```go
// GetActiveByPatientID returns the active subscription with full method hierarchy
func (r *SubscriptionRepository) GetActiveByPatientID(patientID uuid.UUID) (*models.PatientSubscription, error) {
    var subscription models.PatientSubscription
    err := r.db.
        Preload("SubscriptionPlan").
        Preload("SubscriptionPlan.Method").
        Preload("SubscriptionPlan.Method.Letters", func(db *gorm.DB) *gorm.DB {
            return db.Where("deleted_at IS NULL").Order("\"order\" ASC")
        }).
        Preload("SubscriptionPlan.Method.Letters.Pillars", func(db *gorm.DB) *gorm.DB {
            return db.Where("deleted_at IS NULL").Order("\"order\" ASC")
        }).
        Where("patient_id = ? AND status = ?", patientID, models.SubscriptionActive).
        First(&subscription).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // No active subscription is not an error
        }
        return nil, err
    }
    return &subscription, nil
}
```

---

### 1.3. `/apps/api/internal/services/subscription_service.go`

**Adicionar método wrapper**:

```go
func (s *SubscriptionService) GetActiveByPatientID(patientID uuid.UUID) (*models.PatientSubscription, error) {
    return s.repo.GetActiveByPatientID(patientID)
}
```

---

### 1.4. `/apps/api/internal/handlers/subscription_handler.go`

**Adicionar handler**:

```go
// GetActivePatientSubscription godoc
// @Summary Get active subscription for patient
// @Description Returns the active subscription with full method hierarchy
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} models.PatientSubscription
// @Success 204 "No active subscription"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id}/subscriptions/active [get]
func (h *SubscriptionHandler) GetActivePatientSubscription(c *fiber.Ctx) error {
    patientID, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
            Error:   "invalid patient id",
            Message: "patient id must be a valid UUID",
        })
    }

    subscription, err := h.service.GetActiveByPatientID(patientID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
            Error:   "failed to fetch active subscription",
            Message: err.Error(),
        })
    }

    if subscription == nil {
        return c.SendStatus(fiber.StatusNoContent)
    }

    return c.JSON(subscription)
}
```

---

### 1.5. `/apps/api/cmd/server/main.go`

**Adicionar rota** (dentro do grupo de patients):

```go
patients.Get("/:id/subscriptions/active", subscriptionHandler.GetActivePatientSubscription)
```

---

### 1.6. Gerar OpenAPI

```bash
pnpm generate
```

---

## 🎨 FASE 2: Frontend - API Layer (~2-3h)

### 2.1. `/apps/web/lib/api/subscription-api.ts`

**Adicionar hook** (após linha ~47):

```typescript
export function useActivePatientSubscription(patientId: string | undefined) {
  return useQuery({
    queryKey: [...subscriptionKeys.byPatient(patientId!), 'active'],
    queryFn: async () => {
      try {
        const response = await apiClient.get<PatientSubscription>(
          `/api/v1/patients/${patientId}/subscriptions/active`
        )
        return response
      } catch (error: any) {
        if (error.status === 204) {
          return null // No active subscription
        }
        throw error
      }
    },
    enabled: !!patientId,
    staleTime: 5 * 60 * 1000,
    retry: false,
  })
}
```

---

### 2.2. `/apps/web/lib/api/health-score-api.ts`

**Atualizar interface** `PatientScoreItemResult` (linhas ~66-78):

```typescript
item?: {
  id: string
  name: string
  unit?: string
  order?: number
  parentItemId?: string | null
  subgroupId?: string
  subgroup?: {
    id: string
    name: string
    order?: number
  }
  methodPillars?: Array<{
    id: string
    letterId: string
    name: string
    order: number
    letter?: {
      id: string
      methodId: string
      code: string
      name: string
      order: number
      icon?: string
      color?: string
      method?: {
        id: string
        name: string
        shortName: string
      }
    }
  }>
}
```

---

### 2.3. `packages/types/src/index.ts`

**Verificar/adicionar types** para Method, MethodLetter, MethodPillar se necessário.

---

## 🧩 FASE 3: Frontend - Lógica de Agrupamento (~10-14h)

### Arquivo principal: `/apps/web/app/(authenticated)/health-scores/[id]/page.tsx`

### 3.1. Imports e Estado

**No início do arquivo, adicionar imports**:

```typescript
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import { useActivePatientSubscription } from "@/lib/api/subscription-api"
```

**Dentro do component, após linha 107**:

```typescript
const { data: activeSubscription } = useActivePatientSubscription(selectedPatient?.id)
```

**Estado para controle de visualização (após linha 118)**:

```typescript
const [viewMode, setViewMode] = useState<'traditional' | 'methodology'>('methodology')
```

**Estados para acordeons da metodologia**:

```typescript
const [openMethods, setOpenMethods] = useState<string[]>([])
const [openLetters, setOpenLetters] = useState<string[]>([])
const [openPillars, setOpenPillars] = useState<string[]>([])
```

---

### 3.2. Função de Agrupamento por Metodologia

**Adicionar antes da função `itemsByGroup` (antes da linha 185)**:

```typescript
// Organize items by Method → Letter → Pillar → Group → Subgroup hierarchy
const itemsByMethodology = useMemo(() => {
  if (!snapshot?.itemResults || !activeSubscription?.subscriptionPlan?.method) {
    return null
  }

  const method = activeSubscription.subscriptionPlan.method
  if (!method.letters || method.letters.length === 0) return null

  // Map to track which items have been assigned to pillars
  const itemsInPillars = new Set<string>()

  // Build regular letters structure
  const regularLetters = method.letters.map(letter => {
    const letterPillars = (letter.pillars || []).map(pillar => {
      // Find all items that belong to this pillar
      const pillarItems = snapshot.itemResults.filter(itemResult =>
        itemResult.item?.methodPillars?.some(mp => mp.id === pillar.id)
      )

      // Track these items as assigned
      pillarItems.forEach(item => itemsInPillars.add(item.id))

      // Group by ScoreGroup → ScoreSubgroup
      const groups = new Map<string, {
        groupId: string
        groupName: string
        groupScore: number
        groupOrder: number
        subgroups: Map<string, {
          subgroupId: string
          subgroupName: string
          subgroupScore: number
          subgroupOrder: number
          items: ItemWithChildren[]
        }>
      }>()

      pillarItems.forEach(item => {
        const groupResult = snapshot.groupResults?.find(gr => gr.groupId === item.groupId)
        const subgroupId = item.item?.subgroupId || "unknown"
        const subgroupName = item.item?.subgroup?.name || "Sem subgrupo"

        if (!groups.has(item.groupId)) {
          groups.set(item.groupId, {
            groupId: item.groupId,
            groupName: groupResult?.group?.name || "Sem grupo",
            groupScore: groupResult?.scorePercentage || 0,
            groupOrder: groupResult?.group?.order ?? 9999,
            subgroups: new Map(),
          })
        }

        const group = groups.get(item.groupId)!

        if (!group.subgroups.has(subgroupId)) {
          group.subgroups.set(subgroupId, {
            subgroupId,
            subgroupName,
            subgroupScore: 0,
            subgroupOrder: item.item?.subgroup?.order ?? 9999,
            items: [],
          })
        }

        group.subgroups.get(subgroupId)!.items.push(item as ItemWithChildren)
      })

      // Organize hierarchies and calculate subgroup scores
      groups.forEach(group => {
        group.subgroups.forEach(subgroup => {
          subgroup.items = organizeItemsHierarchy(subgroup.items)

          const evaluatedItems = pillarItems.filter(
            item => item.item?.subgroupId === subgroup.subgroupId && item.status === "evaluated"
          )

          if (evaluatedItems.length > 0) {
            const totalActual = evaluatedItems.reduce((sum, item) => sum + item.actualPoints, 0)
            const totalMax = evaluatedItems.reduce((sum, item) => sum + item.maxPoints, 0)
            subgroup.subgroupScore = totalMax > 0 ? (totalActual / totalMax) * 100 : 0
          }
        })
      })

      // Calculate pillar score (items can be in multiple pillars, so count for each)
      const pillarEvaluatedItems = pillarItems.filter(item => item.status === "evaluated")
      const pillarScore = pillarEvaluatedItems.length > 0
        ? (pillarEvaluatedItems.reduce((sum, item) => sum + item.actualPoints, 0) /
           pillarEvaluatedItems.reduce((sum, item) => sum + item.maxPoints, 0)) * 100
        : 0

      return {
        pillar,
        pillarScore,
        itemCount: pillarItems.length,
        evaluatedCount: pillarEvaluatedItems.length,
        groups: Array.from(groups.values())
          .sort((a, b) => a.groupOrder - b.groupOrder)
          .map(g => ({
            ...g,
            subgroups: Array.from(g.subgroups.values()).sort((a, b) => a.subgroupOrder - b.subgroupOrder)
          })),
      }
    }).filter(p => p.groups.length > 0)

    return {
      letter,
      pillars: letterPillars,
    }
  }).filter(l => l.pillars.length > 0)

  // Build virtual "Outros" letter for unassigned items
  const unassignedItems = snapshot.itemResults.filter(item => !itemsInPillars.has(item.id))

  let virtualLetter = null
  if (unassignedItems.length > 0) {
    // Group unassigned items by Group → Subgroup
    const groups = new Map<string, {
      groupId: string
      groupName: string
      groupScore: number
      groupOrder: number
      subgroups: Map<string, {
        subgroupId: string
        subgroupName: string
        subgroupScore: number
        subgroupOrder: number
        items: ItemWithChildren[]
      }>
    }>()

    unassignedItems.forEach(item => {
      const groupResult = snapshot.groupResults?.find(gr => gr.groupId === item.groupId)
      const subgroupId = item.item?.subgroupId || "unknown"
      const subgroupName = item.item?.subgroup?.name || "Sem subgrupo"

      if (!groups.has(item.groupId)) {
        groups.set(item.groupId, {
          groupId: item.groupId,
          groupName: groupResult?.group?.name || "Sem grupo",
          groupScore: groupResult?.scorePercentage || 0,
          groupOrder: groupResult?.group?.order ?? 9999,
          subgroups: new Map(),
        })
      }

      const group = groups.get(item.groupId)!

      if (!group.subgroups.has(subgroupId)) {
        group.subgroups.set(subgroupId, {
          subgroupId,
          subgroupName,
          subgroupScore: 0,
          subgroupOrder: item.item?.subgroup?.order ?? 9999,
          items: [],
        })
      }

      group.subgroups.get(subgroupId)!.items.push(item as ItemWithChildren)
    })

    // Organize hierarchies and calculate scores
    groups.forEach(group => {
      group.subgroups.forEach(subgroup => {
        subgroup.items = organizeItemsHierarchy(subgroup.items)

        const evaluatedItems = unassignedItems.filter(
          item => item.item?.subgroupId === subgroup.subgroupId && item.status === "evaluated"
        )

        if (evaluatedItems.length > 0) {
          const totalActual = evaluatedItems.reduce((sum, item) => sum + item.actualPoints, 0)
          const totalMax = evaluatedItems.reduce((sum, item) => sum + item.maxPoints, 0)
          subgroup.subgroupScore = totalMax > 0 ? (totalActual / totalMax) * 100 : 0
        }
      })
    })

    const evaluatedUnassigned = unassignedItems.filter(item => item.status === "evaluated")
    const virtualPillarScore = evaluatedUnassigned.length > 0
      ? (evaluatedUnassigned.reduce((sum, item) => sum + item.actualPoints, 0) /
         evaluatedUnassigned.reduce((sum, item) => sum + item.maxPoints, 0)) * 100
      : 0

    virtualLetter = {
      letter: {
        id: 'virtual-outros',
        code: '•',
        name: 'Outros',
        order: 9999,
      },
      pillars: [{
        pillar: {
          id: 'virtual-unassigned',
          name: 'Items não incluídos nos pilares do método',
          order: 0,
        },
        pillarScore: virtualPillarScore,
        itemCount: unassignedItems.length,
        evaluatedCount: evaluatedUnassigned.length,
        groups: Array.from(groups.values())
          .sort((a, b) => a.groupOrder - b.groupOrder)
          .map(g => ({
            ...g,
            subgroups: Array.from(g.subgroups.values()).sort((a, b) => a.subgroupOrder - b.subgroupOrder)
          })),
      }],
    }
  }

  return {
    method,
    letters: virtualLetter ? [...regularLetters, virtualLetter] : regularLetters,
  }
}, [snapshot, activeSubscription])

// Decide which view to show
const shouldShowMethodologyView = viewMode === 'methodology' && itemsByMethodology !== null
```

---

### 3.3. Auto-open Hierarquia para Selected Item

**Modificar o useEffect existente (linhas ~146-164)**:

```typescript
// Auto-open all levels that contain the selected item
useEffect(() => {
  if (selectedItemId && snapshot?.itemResults) {
    const selectedItem = snapshot.itemResults.find(ir => ir.id === selectedItemId)
    if (selectedItem) {
      const groupId = selectedItem.groupId
      const subgroupId = selectedItem.item?.subgroupId

      // Traditional view - open group and subgroup
      if (groupId && !openGroups.includes(groupId)) {
        setOpenGroups(prev => [...prev, groupId])
      }
      if (subgroupId && !openSubgroups.includes(subgroupId)) {
        setOpenSubgroups(prev => [...prev, subgroupId])
      }

      // Methodology view - also open method, letter, pillar
      if (shouldShowMethodologyView && itemsByMethodology) {
        // Find which letter and pillar contain this item
        itemsByMethodology.letters.forEach(letterData => {
          letterData.pillars.forEach(pillarData => {
            const hasItem = pillarData.groups.some(g =>
              g.subgroups.some(sg =>
                sg.items.some(item => item.id === selectedItemId)
              )
            )

            if (hasItem) {
              // Open method
              if (!openMethods.includes(itemsByMethodology.method.id)) {
                setOpenMethods(prev => [...prev, itemsByMethodology.method.id])
              }
              // Open letter
              if (!openLetters.includes(letterData.letter.id)) {
                setOpenLetters(prev => [...prev, letterData.letter.id])
              }
              // Open pillar
              if (!openPillars.includes(pillarData.pillar.id)) {
                setOpenPillars(prev => [...prev, pillarData.pillar.id])
              }
            }
          })
        })
      }
    }
  }
}, [selectedItemId, snapshot, shouldShowMethodologyView, itemsByMethodology, openGroups, openSubgroups, openMethods, openLetters, openPillars])
```

---

### 3.4. Renderização com Toggle

**Substituir o bloco de renderização de acordeons (linhas ~622-737)** com:

```typescript
{/* View Mode Toggle - only show if methodology view is available */}
{itemsByMethodology && (
  <Card className="bg-gradient-to-r from-primary/5 to-primary/10 border-primary/20">
    <CardContent className="pt-6">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-lg font-semibold flex items-center gap-2">
            <span className="text-2xl">🎯</span>
            {itemsByMethodology.method.name}
          </h3>
          <p className="text-sm text-muted-foreground mt-1">
            Metodologia clínica vinculada à sua assinatura
          </p>
        </div>

        <ToggleGroup type="single" value={viewMode} onValueChange={(v) => v && setViewMode(v as any)}>
          <ToggleGroupItem value="methodology" aria-label="Visualização por Metodologia">
            <span className="text-xs">Por Metodologia</span>
          </ToggleGroupItem>
          <ToggleGroupItem value="traditional" aria-label="Visualização Tradicional">
            <span className="text-xs">Tradicional</span>
          </ToggleGroupItem>
        </ToggleGroup>
      </div>
    </CardContent>
  </Card>
)}

{/* Conditional Rendering: Methodology vs Traditional */}
{shouldShowMethodologyView ? (
  // METHODOLOGY VIEW
  <Accordion type="multiple" className="space-y-3" value={openLetters} onValueChange={setOpenLetters}>
    {itemsByMethodology!.letters.map((letterData) => (
      <AccordionItem
        key={letterData.letter.id}
        value={letterData.letter.id}
        className="rounded-lg border-2 transition-all"
        style={{ borderColor: letterData.letter.color || 'hsl(var(--border))' }}
      >
        {/* LETTER HEADER */}
        <div className="relative bg-gradient-to-r from-primary/10 to-primary/5 rounded-t-lg">
          <AccordionTrigger className="py-4 px-4 hover:no-underline">
            <div className="flex items-center gap-3 flex-1">
              {letterData.letter.icon && (
                <span className="text-3xl">{letterData.letter.icon}</span>
              )}
              <div className="text-left">
                <h2 className="text-xl font-bold">
                  {letterData.letter.code} - {letterData.letter.name}
                </h2>
                <p className="text-xs text-muted-foreground mt-1">
                  {letterData.pillars.length} {letterData.pillars.length === 1 ? 'pilar' : 'pilares'}
                </p>
              </div>
            </div>
          </AccordionTrigger>
        </div>

        {/* PILLARS */}
        <AccordionContent className="px-2 pb-2 pt-2">
          <Accordion type="multiple" className="space-y-2" value={openPillars} onValueChange={setOpenPillars}>
            {letterData.pillars.map((pillarData) => (
              <AccordionItem
                key={pillarData.pillar.id}
                value={pillarData.pillar.id}
                className="border-2 rounded-md transition-all ml-4"
              >
                {/* PILLAR HEADER */}
                <div className="relative bg-muted/30">
                  <AccordionTrigger className="py-3 px-3 hover:no-underline">
                    <div className="flex items-center w-full pr-4">
                      <div className="flex items-center gap-2 flex-1">
                        <span className="text-base font-semibold">{pillarData.pillar.name}</span>
                        <Badge variant="outline" className="text-xs">
                          {pillarData.evaluatedCount}/{pillarData.itemCount}
                        </Badge>
                      </div>
                      <div className="w-20 flex justify-end">
                        {pillarData.pillarScore > 0 && (
                          <Badge className={`${getScoreColor(pillarData.pillarScore)} text-sm`}>
                            {pillarData.pillarScore.toFixed(1)}%
                          </Badge>
                        )}
                      </div>
                    </div>
                  </AccordionTrigger>
                </div>

                {/* GROUPS within PILLAR */}
                <AccordionContent className="px-2 pb-2">
                  <Accordion type="multiple" className="space-y-1.5" value={openGroups} onValueChange={setOpenGroups}>
                    {pillarData.groups.map((group) => (
                      <AccordionItem
                        key={group.groupId}
                        value={group.groupId}
                        className="border rounded-md transition-all ml-2"
                      >
                        {/* GROUP HEADER */}
                        <div className="relative bg-muted/50">
                          <AccordionTrigger className="py-2 px-3 hover:no-underline">
                            <div className="flex items-center w-full pr-4">
                              <span className="text-sm font-medium flex-1 text-left">{group.groupName}</span>
                              <div className="w-16 flex justify-end">
                                {group.groupScore > 0 && (
                                  <Badge className={`${getScoreColor(group.groupScore)} text-xs`}>
                                    {group.groupScore.toFixed(1)}%
                                  </Badge>
                                )}
                              </div>
                            </div>
                          </AccordionTrigger>
                        </div>

                        {/* SUBGROUPS */}
                        <AccordionContent className="px-2 pb-2">
                          <Accordion type="multiple" className="space-y-1" value={openSubgroups} onValueChange={setOpenSubgroups}>
                            {group.subgroups.map((subgroup) => (
                              <AccordionItem
                                key={subgroup.subgroupId}
                                value={subgroup.subgroupId}
                                className="border rounded-sm transition-all ml-2"
                              >
                                {/* SUBGROUP HEADER */}
                                <AccordionTrigger className="py-2 px-2.5 text-left hover:no-underline">
                                  <div className="flex items-center w-full pr-4">
                                    <div className="flex items-center gap-2 flex-1">
                                      <span className="text-sm font-medium text-left">{subgroup.subgroupName}</span>
                                      <Badge variant="outline" className="text-[10px] px-1.5 py-0">
                                        {subgroup.items.length}
                                      </Badge>
                                    </div>
                                    <div className="w-16 flex justify-end">
                                      {subgroup.subgroupScore > 0 && (
                                        <Badge className={`${getScoreColor(subgroup.subgroupScore)} text-xs`}>
                                          {subgroup.subgroupScore.toFixed(1)}%
                                        </Badge>
                                      )}
                                    </div>
                                  </div>
                                </AccordionTrigger>

                                {/* ITEMS */}
                                <AccordionContent className="px-2.5 pb-2">
                                  <div className="space-y-1.5 mt-1">
                                    {filterItemsRecursive(subgroup.items).map((item) =>
                                      renderItemWithChildren(item, 0)
                                    )}
                                  </div>
                                </AccordionContent>
                              </AccordionItem>
                            ))}
                          </Accordion>
                        </AccordionContent>
                      </AccordionItem>
                    ))}
                  </Accordion>
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
        </AccordionContent>
      </AccordionItem>
    ))}
  </Accordion>
) : (
  // TRADITIONAL VIEW - Keep existing code unchanged
  <Accordion type="multiple" className="space-y-2" value={openGroups} onValueChange={setOpenGroups}>
    {itemsByGroup.map((group, index) => (
      // ... existing rendering code from lines 629-736 ...
    ))}
  </Accordion>
)}
```

---

## ✅ CHECKLIST DE IMPLEMENTAÇÃO

### Backend
- [ ] Adicionar preloads de MethodPillars em `score_snapshot_repository.go`
- [ ] Criar `GetActiveByPatientID` em `subscription_repository.go`
- [ ] Adicionar método no `subscription_service.go`
- [ ] Criar handler `GetActivePatientSubscription` em `subscription_handler.go`
- [ ] Adicionar rota em `main.go`
- [ ] Testar endpoint com Postman/curl
- [ ] Rodar `pnpm generate` para atualizar OpenAPI e types

### Frontend
- [ ] Adicionar hook `useActivePatientSubscription` em `subscription-api.ts`
- [ ] Atualizar types do `PatientScoreItemResult` em `health-score-api.ts`
- [ ] Adicionar imports no componente de health-scores
- [ ] Adicionar query de subscription ativa
- [ ] Adicionar estados `viewMode`, `openLetters`, `openPillars`
- [ ] Implementar função `itemsByMethodology`
- [ ] Criar letter virtual "Outros"
- [ ] Atualizar useEffect de auto-open para metodologia
- [ ] Implementar renderização condicional com toggle
- [ ] Testar com paciente COM subscription ativa
- [ ] Testar com paciente SEM subscription
- [ ] Testar toggle entre views
- [ ] Testar deep linking (scroll to item)
- [ ] Testar filtros

---

## 🎯 Ordem de Execução Recomendada

### Dia 1: Backend
1. Modificar `score_snapshot_repository.go` (preloads)
2. Criar método em `subscription_repository.go`
3. Adicionar service layer
4. Criar handler
5. Adicionar rota
6. Testar endpoint com Postman
7. Rodar `pnpm generate`

### Dia 2: Frontend API + Lógica
1. Adicionar hook `useActivePatientSubscription`
2. Atualizar types
3. Adicionar estados no componente
4. Implementar função `itemsByMethodology`
5. Testar agrupamento no console.log

### Dia 3: Frontend UI
1. Implementar toggle de visualização
2. Implementar renderização de metodologia
3. Ajustar auto-open para hierarquia completa
4. Styling e polimento

### Dia 4: Testes
1. Teste completo com subscription
2. Teste sem subscription
3. Teste deep linking
4. Teste filtros
5. Fix de bugs e ajustes finais

---

## 🚨 Desafios Conhecidos e Soluções

### 1. Items em Múltiplos Pilares
**Problema:** Item pode estar em vários pilares
**Solução:** Renderizar em cada pilar e contar pontos para todos (scores de pilares podem somar > 100%, mas score total do paciente é único)

### 2. Performance com Dados Grandes
**Problema:** Muitos loops aninhados
**Solução:** `useMemo` para cache, considerar virtualization se necessário

### 3. Gerenciamento de Estado Complexo
**Problema:** 6 níveis de hierarquia com open/close
**Solução:** Estados separados (`openLetters`, `openPillars`, etc.) seguindo padrão existente

### 4. Backwards Compatibility
**Problema:** Pacientes sem subscription devem ver view tradicional
**Solução:** Renderização condicional limpa, código tradicional permanece intacto

---

## 📊 Estimativa de Esforço

- **Backend:** 4-6 horas
- **Frontend API:** 2-3 horas
- **Frontend Lógica:** 6-8 horas
- **Frontend UI:** 4-6 horas
- **Testes:** 4-6 horas
- **Total:** 20-29 horas (2.5-4 dias)

---

## 📝 Notas Importantes

1. **Score Total do Paciente:** Mantém a lógica atual - cada item conta apenas uma vez no total geral, mesmo que apareça em múltiplos pilares
2. **Filtro "Apenas Calculados":** Na view metodológica, mantém todos os níveis visíveis sempre
3. **Toggle Padrão:** Quando paciente tem subscription, abre automaticamente em "Por Metodologia"
4. **Letter Virtual:** Sempre aparece por último se houver items não associados a pilares

---

**Autor:** Claude Code
**Revisão:** v1.0
**Status:** Pronto para implementação
