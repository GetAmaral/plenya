"use client"

import { useParams, useRouter } from "next/navigation"
import { useMemo, useState } from "react"
import { useRequireAuth } from "@/lib/use-auth"
import { useHealthScoreDetail } from "@/lib/api/health-score-api"
import type { PatientScoreItemResult } from "@/lib/api/health-score-api"
import { PageHeader } from "@/components/layout/page-header"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Skeleton } from "@/components/ui/skeleton"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import { ArrowLeft, Activity, CheckCircle2, XCircle, MinusCircle, AlertCircle, FlaskConical } from "lucide-react"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"

// Estilos de cor por nível (igual à página de anamnese)
const LEVEL_STYLES = {
  0: { bg: 'bg-red-100', text: 'text-red-900', border: 'border-red-500' },
  1: { bg: 'bg-orange-100', text: 'text-orange-900', border: 'border-orange-500' },
  2: { bg: 'bg-yellow-100', text: 'text-yellow-900', border: 'border-yellow-500' },
  3: { bg: 'bg-blue-100', text: 'text-blue-900', border: 'border-blue-500' },
  4: { bg: 'bg-green-100', text: 'text-green-900', border: 'border-green-500' },
  5: { bg: 'bg-emerald-100', text: 'text-emerald-900', border: 'border-emerald-500' },
  6: { bg: 'bg-gray-100', text: 'text-gray-900', border: 'border-gray-500' },
} as const

interface ItemWithChildren extends PatientScoreItemResult {
  children?: ItemWithChildren[]
}

// Helper to organize items in hierarchy (parent → children) - OUTSIDE component
function organizeItemsHierarchy(items: ItemWithChildren[]): ItemWithChildren[] {
  const itemMap = new Map<string, ItemWithChildren>()
  const rootItems: ItemWithChildren[] = []

  // First, create map of all items using their unique ID (PatientScoreItemResult.id)
  items.forEach(item => {
    itemMap.set(item.id, { ...item, children: [] })
  })

  // Then, organize hierarchy based on ScoreItem.parentItemId
  items.forEach(item => {
    const itemWithChildren = itemMap.get(item.id)!
    const parentItemId = item.item?.parentItemId

    if (parentItemId) {
      // This item has a parent - find the PatientScoreItemResult that references this parent ScoreItem
      const parentResult = items.find(i => i.itemId === parentItemId)
      if (parentResult) {
        const parent = itemMap.get(parentResult.id)
        if (parent) {
          parent.children = parent.children || []
          parent.children.push(itemWithChildren)
        } else {
          rootItems.push(itemWithChildren)
        }
      } else {
        // Parent not found in results, add as root
        rootItems.push(itemWithChildren)
      }
    } else {
      // No parent - this is a root item
      rootItems.push(itemWithChildren)
    }
  })

  // Sort all levels recursively by order field (same as ScoreTreeView)
  const sortItemsRecursive = (items: ItemWithChildren[]) => {
    items.sort((a, b) => (a.item?.order || 0) - (b.item?.order || 0))
    items.forEach(item => {
      if (item.children && item.children.length > 0) {
        sortItemsRecursive(item.children)
      }
    })
  }

  sortItemsRecursive(rootItems)

  return rootItems
}

export default function HealthScoreDetailPage() {
  useRequireAuth()
  const params = useParams()
  const router = useRouter()
  const snapshotId = params.id as string
  const [showOnlyEvaluated, setShowOnlyEvaluated] = useState(false)

  const { data: snapshot, isLoading, error } = useHealthScoreDetail(snapshotId)

  // Organize items by group → subgroup → items with hierarchy
  const itemsByGroup = useMemo(() => {
    if (!snapshot?.itemResults || !snapshot?.groupResults) return []

    const groups = new Map<string, {
      groupId: string
      groupName: string
      groupScore: number
      subgroups: Map<string, {
        subgroupId: string
        subgroupName: string
        subgroupScore: number
        items: ItemWithChildren[]
      }>
    }>()

    // First, organize items by group and subgroup
    snapshot.itemResults.forEach((item) => {
      const groupResult = snapshot.groupResults?.find((gr) => gr.groupId === item.groupId)
      const subgroupId = item.item?.subgroupId || "unknown"
      const subgroupName = item.item?.subgroup?.name || "Sem subgrupo"

      if (!groups.has(item.groupId)) {
        groups.set(item.groupId, {
          groupId: item.groupId,
          groupName: groupResult?.group?.name || "Sem grupo",
          groupScore: groupResult?.scorePercentage || 0,
          subgroups: new Map(),
        })
      }

      const group = groups.get(item.groupId)!

      if (!group.subgroups.has(subgroupId)) {
        group.subgroups.set(subgroupId, {
          subgroupId,
          subgroupName,
          subgroupScore: 0,
          items: [],
        })
      }

      group.subgroups.get(subgroupId)!.items.push(item as ItemWithChildren)
    })

    // Organize items hierarchy and calculate subgroup scores
    groups.forEach((group) => {
      group.subgroups.forEach((subgroup) => {
        subgroup.items = organizeItemsHierarchy(subgroup.items)

        // Calculate subgroup score (only evaluated items)
        const evaluatedItems = snapshot.itemResults.filter(
          item => item.item?.subgroupId === subgroup.subgroupId && item.status === "evaluated"
        )

        if (evaluatedItems.length > 0) {
          const totalActual = evaluatedItems.reduce((sum, item) => sum + item.actualPoints, 0)
          const totalMax = evaluatedItems.reduce((sum, item) => sum + item.maxPoints, 0)
          subgroup.subgroupScore = totalMax > 0 ? (totalActual / totalMax) * 100 : 0
        }
      })
    })

    return Array.from(groups.values()).sort((a, b) =>
      a.groupName.localeCompare(b.groupName)
    )
  }, [snapshot])

  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "bg-green-500 text-white"
    if (percentage >= 71) return "bg-blue-500 text-white"
    if (percentage >= 51) return "bg-yellow-500 text-white"
    if (percentage >= 31) return "bg-orange-500 text-white"
    return "bg-red-500 text-white"
  }

  // Calculate score of a parent item based on its children (recursively)
  const calculateParentScore = (item: ItemWithChildren): { actualPoints: number, maxPoints: number, percentage: number } | null => {
    if (!item.children || item.children.length === 0) {
      return null // Not a parent
    }

    let totalActual = 0
    let totalMax = 0

    const sumChildrenPoints = (children: ItemWithChildren[]) => {
      children.forEach(child => {
        if (child.status === "evaluated") {
          totalActual += child.actualPoints
          totalMax += child.maxPoints
        }
        // Recursively sum grandchildren
        if (child.children && child.children.length > 0) {
          sumChildrenPoints(child.children)
        }
      })
    }

    sumChildrenPoints(item.children)

    if (totalMax === 0) return null

    return {
      actualPoints: totalActual,
      maxPoints: totalMax,
      percentage: (totalActual / totalMax) * 100
    }
  }

  // Filter items recursively (keep parent if any child has data)
  const filterItemsRecursive = (items: ItemWithChildren[]): ItemWithChildren[] => {
    if (!showOnlyEvaluated) return items

    return items
      .map(item => {
        // First, recursively filter children
        const filteredChildren = item.children ? filterItemsRecursive(item.children) : []

        // Include item if:
        // 1. It's evaluated (has data), OR
        // 2. It has children that passed the filter (at least one child with data)
        if (item.status === "evaluated" || filteredChildren.length > 0) {
          return {
            ...item,
            children: filteredChildren
          }
        }

        // Exclude this item (not evaluated and no children with data)
        return null
      })
      .filter((item): item is ItemWithChildren => item !== null)
  }

  // Render item with children recursively (with indentation - exactly like ScoreTreeView)
  const renderItemWithChildren = (item: ItemWithChildren, depth: number = 0): JSX.Element[] => {
    const elements: JSX.Element[] = []
    const levelStyle = item.levelNumber !== null && item.levelNumber !== undefined
      ? LEVEL_STYLES[item.levelNumber as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
      : null

    // Same style as ScoreTreeView
    const style = {
      marginLeft: `${depth * 20}px`,
      paddingLeft: depth > 0 ? '8px' : '0',
      borderLeft: depth > 0 ? '2px solid hsl(var(--border))' : 'none',
    }

    // Calculate parent score if item has children
    const parentScore = calculateParentScore(item)

    // Render current item (use unique PatientScoreItemResult.id as key)
    elements.push(
      <div
        key={item.id}
        style={style}
        className="rounded-lg transition-all flex items-start gap-1"
      >
        {/* Card with item content */}
        <div className="flex-1 min-w-0">
          <div className="text-sm border rounded-md p-2.5 bg-card">
            <div className="flex items-start gap-2">
              <div className="flex-1 min-w-0">
                <span className="font-semibold">{item.item?.name}</span>
                {item.status === "evaluated" && (
                  <div className="text-xs text-muted-foreground mt-1">
                    {item.actualPoints.toFixed(1)} / {item.maxPoints.toFixed(1)} pts
                    ({((item.actualPoints / item.maxPoints) * 100).toFixed(0)}%)
                  </div>
                )}
                {item.status === "no_data_available" && item.maxPoints > 0 && (
                  <div className="text-xs text-muted-foreground mt-1">
                    0 / {item.maxPoints.toFixed(1)} pts
                  </div>
                )}
                {item.status === "not_applicable" && item.notEvaluatedReason && (
                  <p className="text-xs text-muted-foreground italic mt-1">
                    {item.notEvaluatedReason}
                  </p>
                )}
              </div>

              {/* Badges column - aligned right with fixed width */}
              <div className="flex flex-col items-end gap-1 w-16 flex-shrink-0">
                {item.status === "evaluated" && levelStyle && (
                  <>
                    {item.dataSource === "anamnesis_item" ? (
                      // Anamnesis: Show level badge
                      <span className={`text-xs px-2 py-0.5 rounded-full border-2 font-bold whitespace-nowrap ${levelStyle.bg} ${levelStyle.text} ${levelStyle.border}`}>
                        N{item.levelNumber}: {item.levelMatched?.name}
                      </span>
                    ) : (
                      // Lab Result: Show numeric value with color based on level
                      <span className={`text-xs px-2 py-0.5 rounded-full border-2 font-bold whitespace-nowrap ${levelStyle.bg} ${levelStyle.text} ${levelStyle.border} flex items-center gap-1`}>
                        <FlaskConical className="h-3 w-3" />
                        {item.valueUsed?.toFixed(2)} {item.item?.unit}
                      </span>
                    )}
                  </>
                )}
                {item.status === "not_applicable" && (
                  <Badge variant="outline" className="bg-gray-100 text-gray-600 border-gray-300 text-xs">
                    Não aplicável
                  </Badge>
                )}
                {/* Only show "Sem dados" badge if item has points (maxPoints > 0) */}
                {item.status === "no_data_available" && item.maxPoints > 0 && (
                  <Badge variant="outline" className="bg-orange-100 text-orange-800 border-orange-300 text-xs">
                    Sem dados
                  </Badge>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    )

    // Render children recursively with increased depth (same as ScoreTreeView)
    if (item.children && item.children.length > 0) {
      item.children.forEach((child) => {
        elements.push(...renderItemWithChildren(child, depth + 1))
      })
    }

    return elements
  }

  if (isLoading) {
    return (
      <div className="flex flex-col gap-6">
        <Skeleton className="h-12 w-full" />
        <Skeleton className="h-64 w-full" />
      </div>
    )
  }

  if (error || !snapshot) {
    return (
      <div className="flex flex-col gap-6">
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertDescription>
            Erro ao carregar detalhes do snapshot. O snapshot pode ter sido deletado.
          </AlertDescription>
        </Alert>
        <Button onClick={() => router.back()}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          Voltar
        </Button>
      </div>
    )
  }

  const evaluatedCount = snapshot.itemResults?.filter((ir) => ir.status === "evaluated").length || 0
  const noDataCount = snapshot.itemResults?.filter((ir) => ir.status === "no_data_available").length || 0
  const notApplicableCount = snapshot.itemResults?.filter((ir) => ir.status === "not_applicable").length || 0

  return (
    <div className="flex flex-col gap-6">
      {/* Header */}
      <PageHeader
        title="Detalhes do Snapshot"
        description={`Calculado em ${format(new Date(snapshot.calculatedAt), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}`}
      >
        <Button variant="outline" onClick={() => router.back()}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          Voltar
        </Button>
      </PageHeader>

      {/* Summary Cards */}
      <div className="grid gap-4 md:grid-cols-4">
        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Score Total
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex items-center gap-2">
              <Activity className="h-4 w-4 text-muted-foreground" />
              <Badge className={getScoreColor(snapshot.totalScorePercentage)} variant="default">
                {snapshot.totalScorePercentage.toFixed(1)}%
              </Badge>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Itens Avaliados
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex items-center gap-2">
              <CheckCircle2 className="h-4 w-4 text-green-600" />
              <span className="text-2xl font-bold">{evaluatedCount}</span>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Sem Dados
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex items-center gap-2">
              <XCircle className="h-4 w-4 text-orange-500" />
              <span className="text-2xl font-bold">{noDataCount}</span>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-sm font-medium text-muted-foreground">
              Não Aplicáveis
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex items-center gap-2">
              <MinusCircle className="h-4 w-4 text-gray-400" />
              <span className="text-2xl font-bold">{notApplicableCount}</span>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Notes */}
      {snapshot.notes && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Observações</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground">{snapshot.notes}</p>
          </CardContent>
        </Card>
      )}

      {/* Results by Group → Subgroup → Items (same structure as ScoreTreeView) */}
      <div className="space-y-2">
        {itemsByGroup.map((group, index) => (
          <div key={group.groupId} className="rounded-lg border transition-all">
            {/* Group Header */}
            <div className="flex items-center px-3 py-2 bg-muted/50 gap-3">
              <div className="flex-1">
                <h2 className="text-sm font-semibold text-left">{group.groupName}</h2>
                <p className="text-xs text-muted-foreground mt-0.5 text-left">
                  {group.subgroups.size} subgrupos • {Array.from(group.subgroups.values()).reduce((acc, sg) => acc + sg.items.length, 0)} itens
                </p>
              </div>

              {/* Show only evaluated switch (only on first group, controls all) */}
              {index === 0 && (
                <div className="flex items-center gap-2">
                  <Label htmlFor="show-only-evaluated" className="text-xs cursor-pointer whitespace-nowrap">
                    {showOnlyEvaluated ? "Apenas calculados" : "Exibir todos"}
                  </Label>
                  <Switch
                    id="show-only-evaluated"
                    checked={showOnlyEvaluated}
                    onCheckedChange={setShowOnlyEvaluated}
                  />
                </div>
              )}

              {/* Group score badge - aligned right with fixed width */}
              <div className="w-16 flex justify-end">
                <Badge className={`${getScoreColor(group.groupScore)} text-sm`}>
                  {group.groupScore.toFixed(1)}%
                </Badge>
              </div>
            </div>

            {/* Subgroups */}
            {group.subgroups.size > 0 && (
              <div className="p-2">
                <Accordion type="multiple" className="space-y-1.5">
                  {Array.from(group.subgroups.values()).map((subgroup) => (
                    <AccordionItem
                      key={subgroup.subgroupId}
                      value={subgroup.subgroupId}
                      className="border rounded-md transition-all"
                    >
                      <AccordionTrigger className="hover:no-underline py-2 px-2.5 text-left">
                        <div className="flex items-center w-full pr-4">
                          <div className="flex items-center gap-2 flex-1">
                            <span className="text-sm font-medium text-left">{subgroup.subgroupName}</span>
                            <Badge variant="outline" className="text-[10px] px-1.5 py-0">
                              {subgroup.items.length}
                            </Badge>
                          </div>
                          {/* Subgroup score badge - aligned right with fixed width */}
                          <div className="w-16 flex justify-end">
                            {subgroup.subgroupScore > 0 && (
                              <Badge className={`${getScoreColor(subgroup.subgroupScore)} text-xs`}>
                                {subgroup.subgroupScore.toFixed(1)}%
                              </Badge>
                            )}
                          </div>
                        </div>
                      </AccordionTrigger>

                      <AccordionContent className="px-2.5 pb-2">
                        <div className="space-y-1.5 mt-1">
                          {filterItemsRecursive(subgroup.items).map((item) => renderItemWithChildren(item, 0))}
                        </div>
                      </AccordionContent>
                    </AccordionItem>
                  ))}
                </Accordion>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}
