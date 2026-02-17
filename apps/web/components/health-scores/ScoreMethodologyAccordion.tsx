"use client"

import { useState, useMemo, useRef } from "react"
import { useRouter } from "next/navigation"
import type { PatientScoreItemResult, PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { useActivePatientSubscription } from "@/lib/api/subscription-api"
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient"
import { Card, CardContent } from "@/components/ui/card"
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { FlaskConical, MoreVertical, Edit } from "lucide-react"

interface ScoreMethodologyAccordionProps {
  snapshot: PatientScoreSnapshot
}

// Estilos de cor por nível (igual à página de detalhes)
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

// Helper para organizar items em hierarquia pai-filho - OUTSIDE component
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

  // Sort all levels recursively by order field
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

export function ScoreMethodologyAccordion({ snapshot }: ScoreMethodologyAccordionProps) {
  const { selectedPatient } = useRequireSelectedPatient()
  const { data: activeSubscription } = useActivePatientSubscription(selectedPatient?.id)
  const router = useRouter()
  const [viewMode, setViewMode] = useState<'traditional' | 'methodology'>('methodology')
  const [openLetters, setOpenLetters] = useState<string[]>([])
  const [openPillars, setOpenPillars] = useState<string[]>([])
  const [openGroups, setOpenGroups] = useState<string[]>([])
  const [openSubgroups, setOpenSubgroups] = useState<string[]>([])
  const [selectedItemId, setSelectedItemId] = useState<string | null>(null)
  const [showOnlyEvaluated, setShowOnlyEvaluated] = useState(false)
  const itemRefs = useRef<Map<string, HTMLDivElement>>(new Map())

  // Toggle item selection
  const handleItemClick = (itemId: string) => {
    setSelectedItemId(prev => prev === itemId ? null : itemId)
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

  // Get edit link based on data source
  const getEditLink = (item: ItemWithChildren): string | null => {
    if (item.status !== "evaluated") return null

    if (item.dataSource === "lab_result" && item.labResult?.labResultBatch?.id) {
      return `/lab-results/${item.labResult.labResultBatch.id}/edit#${item.labResultId}`
    }

    if (item.dataSource === "anamnesis_item" && item.anamnesisItem?.anamnesis?.id && item.itemId) {
      return `/anamnesis?edit=${item.anamnesisItem.anamnesis.id}&focusItem=${item.itemId}`
    }

    return null
  }

  // Handle edit click
  const handleEditClick = (item: ItemWithChildren) => {
    setSelectedItemId(item.id)
    const editLink = getEditLink(item)
    if (editLink) {
      router.push(editLink)
    }
  }

  const getBadgeScoreColor = (percentage: number) => {
    if (percentage >= 86) return "bg-green-500 text-white"
    if (percentage >= 71) return "bg-blue-500 text-white"
    if (percentage >= 51) return "bg-yellow-500 text-white"
    if (percentage >= 31) return "bg-orange-500 text-white"
    return "bg-red-500 text-white"
  }

  // Render item with children recursively
  const renderItemWithChildren = (item: ItemWithChildren, depth: number = 0): JSX.Element[] => {
    const elements: JSX.Element[] = []
    const levelStyle = item.levelNumber !== null && item.levelNumber !== undefined
      ? LEVEL_STYLES[item.levelNumber as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
      : null

    const style = {
      marginLeft: `${depth * 20}px`,
      paddingLeft: depth > 0 ? '8px' : '0',
      borderLeft: depth > 0 ? '2px solid hsl(var(--border))' : 'none',
    }

    const parentScore = calculateParentScore(item)
    const isSelected = selectedItemId === item.id

    elements.push(
      <div
        key={item.id}
        ref={(el) => {
          if (el) itemRefs.current.set(item.id, el)
        }}
        style={style}
        className="rounded-lg transition-all flex items-start gap-1"
      >
        <div className="flex-1 min-w-0">
          <div
            className={`text-sm border rounded-md p-2.5 bg-card cursor-pointer transition-all ${
              isSelected ? 'border-primary bg-primary/5 ring-2 ring-primary/20' : 'hover:border-primary/50'
            }`}
            onClick={() => handleItemClick(item.id)}
          >
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

              <div className="flex items-start gap-1 flex-shrink-0">
                <div className="flex flex-col items-end gap-1 w-16">
                  {parentScore && (
                    <Badge className={`${getBadgeScoreColor(parentScore.percentage)} text-[9px] whitespace-nowrap`}>
                      {parentScore.percentage.toFixed(1)}%
                    </Badge>
                  )}
                  {item.status === "evaluated" && levelStyle && (
                    <>
                      {item.dataSource === "anamnesis_item" ? (
                        <span className={`text-xs px-2 py-0.5 rounded-full border-2 font-bold whitespace-nowrap ${levelStyle.bg} ${levelStyle.text} ${levelStyle.border}`}>
                          N{item.levelNumber}: {item.levelMatched?.name}
                        </span>
                      ) : (
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
                  {item.status === "no_data_available" && item.maxPoints > 0 && (
                    <Badge variant="outline" className="bg-orange-100 text-orange-800 border-orange-300 text-xs">
                      Sem dados
                    </Badge>
                  )}
                </div>

                {getEditLink(item) && (
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="sm" className="h-6 w-6 p-0">
                        <MoreVertical className="h-3 w-3" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuItem
                        onClick={() => handleEditClick(item)}
                        className="cursor-pointer"
                      >
                        <Edit className="mr-2 h-4 w-4" />
                        Editar
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    )

    if (item.children && item.children.length > 0) {
      item.children.forEach((child) => {
        elements.push(...renderItemWithChildren(child, depth + 1))
      })
    }

    return elements
  }

  // Organize items by Method → Letter → Pillar → Group → Subgroup hierarchy
  const itemsByMethodology = useMemo(() => {
    if (!snapshot?.itemResults || !activeSubscription?.subscriptionPlan?.method) {
      return null
    }

    const method = activeSubscription.subscriptionPlan.method
    if (!method.letters || method.letters.length === 0) return null

    const itemsInPillars = new Set<string>()

    const regularLetters = method.letters.map(letter => {
      const letterPillars = (letter.pillars || []).map(pillar => {
        const pillarItems = snapshot.itemResults.filter(itemResult =>
          itemResult.item?.methodPillars?.some(mp => mp.id === pillar.id)
        )

        pillarItems.forEach(item => itemsInPillars.add(item.id))

        const groups = new Map<string, any>()

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

        groups.forEach(group => {
          group.subgroups.forEach((subgroup: any) => {
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

    return {
      method,
      letters: regularLetters,
    }
  }, [snapshot, activeSubscription])

  // Organize items by group → subgroup (Traditional view)
  const itemsByGroup = useMemo(() => {
    if (!snapshot?.itemResults || !snapshot?.groupResults) return []

    const groups = new Map<string, any>()

    snapshot.itemResults.forEach(item => {
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

    groups.forEach(group => {
      group.subgroups.forEach((subgroup: any) => {
        subgroup.items = organizeItemsHierarchy(subgroup.items)

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

    return Array.from(groups.values())
      .sort((a, b) => a.groupOrder - b.groupOrder)
      .map(g => ({
        ...g,
        subgroups: Array.from(g.subgroups.values()).sort((a, b) => a.subgroupOrder - b.subgroupOrder)
      }))
  }, [snapshot])

  const shouldShowMethodologyView = viewMode === 'methodology' && itemsByMethodology !== null

  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "bg-green-600 text-white border-green-600"
    if (percentage >= 71) return "bg-blue-500 text-white border-blue-500"
    if (percentage >= 51) return "bg-yellow-500 text-white border-yellow-500"
    if (percentage >= 31) return "bg-orange-500 text-white border-orange-500"
    return "bg-red-500 text-white border-red-500"
  }

  return (
    <div className="space-y-4">
      {/* Toggle */}
      {itemsByMethodology && (
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <h3 className="text-lg font-semibold">
                  {itemsByMethodology.method.name}
                </h3>
                <p className="text-sm text-muted-foreground mt-1">
                  Metodologia clínica vinculada à sua assinatura
                </p>
              </div>

              <div className="flex items-center gap-4">
                {/* Show only evaluated switch */}
                <div className="flex items-center gap-2">
                  <Label htmlFor="show-only-evaluated-accordion" className="text-xs cursor-pointer whitespace-nowrap">
                    {showOnlyEvaluated ? "Apenas calculados" : "Exibir todos"}
                  </Label>
                  <Switch
                    id="show-only-evaluated-accordion"
                    checked={showOnlyEvaluated}
                    onCheckedChange={setShowOnlyEvaluated}
                  />
                </div>

                <ToggleGroup type="single" value={viewMode} onValueChange={(v) => v && setViewMode(v as any)}>
                  <ToggleGroupItem value="methodology" aria-label="Visualização por Metodologia">
                    <span className="text-xs">Por Metodologia</span>
                  </ToggleGroupItem>
                  <ToggleGroupItem value="traditional" aria-label="Visualização Tradicional">
                    <span className="text-xs">Por Grupos</span>
                  </ToggleGroupItem>
                </ToggleGroup>
              </div>
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
              style={{ borderColor: (letterData.letter as any).color || 'hsl(var(--border))' }}
            >
              <div className="relative bg-gradient-to-r from-primary/10 to-primary/5 rounded-t-lg">
                <AccordionTrigger className="py-4 px-4 hover:no-underline">
                  <div className="flex items-center gap-3 flex-1">
                    <div className="text-left">
                      <h2 className="text-xl font-bold">
                        {(letterData.letter as any).code} - {(letterData.letter as any).name}
                      </h2>
                      <p className="text-xs text-muted-foreground mt-1">
                        {letterData.pillars.length} {letterData.pillars.length === 1 ? 'pilar' : 'pilares'}
                      </p>
                    </div>
                  </div>
                </AccordionTrigger>
              </div>

              <AccordionContent className="px-2 pb-2 pt-2">
                <Accordion type="multiple" className="space-y-2" value={openPillars} onValueChange={setOpenPillars}>
                  {letterData.pillars.map((pillarData) => (
                    <AccordionItem
                      key={pillarData.pillar.id}
                      value={pillarData.pillar.id}
                      className="border-2 rounded-md transition-all ml-4"
                    >
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

                      <AccordionContent className="px-2 pb-2">
                        <Accordion type="multiple" className="space-y-1.5" value={openGroups} onValueChange={setOpenGroups}>
                            {pillarData.groups.map((group) => (
                              <AccordionItem
                                key={group.groupId}
                                value={group.groupId}
                                className="border rounded-md transition-all ml-2"
                              >
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

                                <AccordionContent className="px-2 pb-2">
                                  <Accordion type="multiple" className="space-y-1" value={openSubgroups} onValueChange={setOpenSubgroups}>
                                    {group.subgroups.map((subgroup: any) => (
                                      <AccordionItem
                                        key={subgroup.subgroupId}
                                        value={subgroup.subgroupId}
                                        className="border rounded-sm transition-all ml-2"
                                      >
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
        // TRADITIONAL VIEW
        <Accordion type="multiple" className="space-y-2" value={openGroups} onValueChange={setOpenGroups}>
          {itemsByGroup.map((group) => (
            <AccordionItem
              key={group.groupId}
              value={group.groupId}
              className="rounded-lg border-2 transition-all"
            >
              <div className="relative bg-gradient-to-r from-muted/50 to-muted/30">
                <AccordionTrigger className="py-3 px-4 hover:no-underline">
                  <div className="flex items-center w-full pr-4">
                    <h3 className="text-base font-semibold flex-1 text-left">{group.groupName}</h3>
                    <div className="w-20 flex justify-end">
                      {group.groupScore > 0 && (
                        <Badge className={`${getScoreColor(group.groupScore)} text-sm`}>
                          {group.groupScore.toFixed(1)}%
                        </Badge>
                      )}
                    </div>
                  </div>
                </AccordionTrigger>
              </div>

              <AccordionContent className="px-2 pb-2">
                <Accordion type="multiple" className="space-y-1" value={openSubgroups} onValueChange={setOpenSubgroups}>
                  {group.subgroups.map((subgroup: any) => (
                    <AccordionItem
                      key={subgroup.subgroupId}
                      value={subgroup.subgroupId}
                      className="border rounded-sm transition-all ml-2"
                    >
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
      )}
    </div>
  )
}
