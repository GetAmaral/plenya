"use client"

import { useParams, useRouter } from "next/navigation"
import { useMemo, useState, useEffect, useRef } from "react"
import { useRequireAuth } from "@/lib/use-auth"
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient"
import { useHealthScoreDetail, useCalculateHealthScore } from "@/lib/api/health-score-api"
import type { PatientScoreItemResult } from "@/lib/api/health-score-api"
import { useActivePatientSubscription } from "@/lib/api/subscription-api"
import { useHealthScoreViewPreferences } from "@/lib/use-health-score-view-preferences"
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
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import { ArrowLeft, Activity, CheckCircle2, XCircle, MinusCircle, AlertCircle, FlaskConical, MoreVertical, Edit, RefreshCw } from "lucide-react"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"
import { toast } from "sonner"

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
  const { selectedPatient } = useRequireSelectedPatient()
  const params = useParams()
  const router = useRouter()
  const snapshotId = params.id as string

  // Preferências de visualização persistidas
  const {
    openGroups,
    setOpenGroups,
    openSubgroups,
    setOpenSubgroups,
    selectedItemId,
    setSelectedItemId,
    showOnlyEvaluated,
    setShowOnlyEvaluated,
  } = useHealthScoreViewPreferences(snapshotId)

  const { data: snapshot, isLoading, error } = useHealthScoreDetail(snapshotId)
  const { data: activeSubscription } = useActivePatientSubscription(selectedPatient?.id)
  const calculateMutation = useCalculateHealthScore()

  // View mode state (methodology vs traditional)
  const [viewMode, setViewMode] = useState<'traditional' | 'methodology'>('methodology')

  // States for methodology view accordions
  const [openLetters, setOpenLetters] = useState<string[]>([])
  const [openPillars, setOpenPillars] = useState<string[]>([])

  // Refs for scrolling to selected item
  const itemRefs = useRef<Map<string, HTMLDivElement>>(new Map())
  const hasScrolledToItem = useRef(false)

  const handleRecalculate = async () => {
    if (!selectedPatient) return

    try {
      toast.info("Recalculando score...")
      const newSnapshot = await calculateMutation.mutateAsync({
        patientId: selectedPatient.id,
        notes: `Recálculo após edição de dados - ${format(new Date(), "dd/MM/yyyy HH:mm", { locale: ptBR })}`,
      })
      toast.success("Score recalculado com sucesso!")
      // Navigate to the new snapshot
      router.push(`/health-scores/${newSnapshot.id}`)
    } catch (error) {
      console.error("Failed to recalculate score:", error)
      toast.error("Erro ao recalcular score")
    }
  }

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

        // Methodology view - also open letter and pillar
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
  }, [selectedItemId, snapshot, shouldShowMethodologyView, itemsByMethodology, openGroups, openSubgroups, openLetters, openPillars, setOpenGroups, setOpenSubgroups])

  // Auto-scroll to selected item when page loads
  useEffect(() => {
    if (selectedItemId && !hasScrolledToItem.current && itemRefs.current.size > 0) {
      const itemElement = itemRefs.current.get(selectedItemId)
      if (itemElement) {
        hasScrolledToItem.current = true
        setTimeout(() => {
          itemElement.scrollIntoView({ behavior: 'smooth', block: 'center' })
        }, 500) // Increased delay to allow accordion to open first
      }
    }
  }, [selectedItemId, itemRefs.current.size])

  // Toggle item selection
  const handleItemClick = (itemId: string) => {
    setSelectedItemId(prev => prev === itemId ? null : itemId)
  }

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

    // Convert groups to array and sort by group.order ASC (same as ScoreTreeView)
    return Array.from(groups.values()).sort((a, b) => {
      const orderA = snapshot.groupResults?.find(gr => gr.groupId === a.groupId)?.group?.order ?? 0
      const orderB = snapshot.groupResults?.find(gr => gr.groupId === b.groupId)?.group?.order ?? 0
      return orderA - orderB
    })
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

  // Get edit link based on data source
  const getEditLink = (item: ItemWithChildren): string | null => {
    if (item.status !== "evaluated") return null

    if (item.dataSource === "lab_result" && item.labResult?.labResultBatch?.id) {
      // Include lab result ID as hash to scroll and focus on that specific field
      return `/lab-results/${item.labResult.labResultBatch.id}/edit#${item.labResultId}`
    }

    if (item.dataSource === "anamnesis_item" && item.anamnesisItem?.anamnesis?.id && item.itemId) {
      // Include anamnesis ID and score item ID as query params
      return `/anamnesis?edit=${item.anamnesisItem.anamnesis.id}&focusItem=${item.itemId}`
    }

    return null
  }

  // Handle edit click - save item as selected and ensure groups/subgroups are open
  const handleEditClick = (item: ItemWithChildren) => {
    // Save item as selected
    setSelectedItemId(item.id)

    // Ensure group and subgroup are open
    const groupId = item.groupId
    const subgroupId = item.item?.subgroupId

    if (groupId && !openGroups.includes(groupId)) {
      setOpenGroups(prev => [...prev, groupId])
    }

    if (subgroupId && !openSubgroups.includes(subgroupId)) {
      setOpenSubgroups(prev => [...prev, subgroupId])
    }

    // Navigate to edit page
    const editLink = getEditLink(item)
    if (editLink) {
      router.push(editLink)
    }
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

    // Check if this item is selected
    const isSelected = selectedItemId === item.id

    // Render current item (use unique PatientScoreItemResult.id as key)
    elements.push(
      <div
        key={item.id}
        ref={(el) => {
          if (el) itemRefs.current.set(item.id, el)
        }}
        style={style}
        className="rounded-lg transition-all flex items-start gap-1"
      >
        {/* Card with item content */}
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

              {/* Badges column - aligned right with fixed width */}
              <div className="flex items-start gap-1 flex-shrink-0">
                <div className="flex flex-col items-end gap-1 w-16">
                  {/* Parent score badge (if has children) - smallest size */}
                  {parentScore && (
                    <Badge className={`${getScoreColor(parentScore.percentage)} text-[9px] whitespace-nowrap`}>
                      {parentScore.percentage.toFixed(1)}%
                    </Badge>
                  )}
                  {/* Item evaluation badges */}
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

                {/* Edit menu button (only for evaluated items) */}
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
        title={snapshot.displayTitle ?? `Snapshot - ${format(new Date(snapshot.calculatedAt), "dd/MM/yyyy HH:mm", { locale: ptBR })}`}
        breadcrumbs={[{ label: "Escores de Saúde", href: "/health-scores" }]}
      >
        <div className="flex items-center gap-2">
          <Button
            variant="outline"
            onClick={handleRecalculate}
            disabled={calculateMutation.isPending}
          >
            <RefreshCw className={`mr-2 h-4 w-4 ${calculateMutation.isPending ? 'animate-spin' : ''}`} />
            Recalcular Score
          </Button>
          <Button variant="outline" onClick={() => router.back()}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            Voltar
          </Button>
        </div>
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

      {/* View Mode Toggle - only show if methodology view is available */}
      {itemsByMethodology && (
        <Card className="bg-gradient-to-r from-primary/5 to-primary/10 border-primary/20">
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
                  <Label htmlFor="show-only-evaluated-detail" className="text-xs cursor-pointer whitespace-nowrap">
                    {showOnlyEvaluated ? "Apenas calculados" : "Exibir todos"}
                  </Label>
                  <Switch
                    id="show-only-evaluated-detail"
                    checked={showOnlyEvaluated}
                    onCheckedChange={setShowOnlyEvaluated}
                  />
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
              {/* LETTER HEADER */}
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
        // TRADITIONAL VIEW
        <Accordion
        type="multiple"
        className="space-y-2"
        value={openGroups}
        onValueChange={setOpenGroups}
      >
        {itemsByGroup.map((group, index) => (
          <AccordionItem
            key={group.groupId}
            value={group.groupId}
            className="rounded-lg border transition-all"
          >
            {/* Group Header - now collapsible with relative positioning */}
            <div className="relative bg-muted/50">
              <AccordionTrigger className="hover:no-underline py-2 px-3 [&>svg]:order-last w-full">
                <div className="flex items-center gap-3 flex-1" style={{ paddingRight: index === 0 ? '240px' : '120px' }}>
                  {/* Group name - takes remaining space */}
                  <div className="flex-1">
                    <h2 className="text-sm font-semibold text-left">{group.groupName}</h2>
                    <p className="text-xs text-muted-foreground mt-0.5 text-left">
                      {group.subgroups.size} subgrupos • {
                        (() => {
                          const totalItems = Array.from(group.subgroups.values()).reduce((acc, sg) => acc + sg.items.length, 0)
                          const evaluatedItems = snapshot.itemResults?.filter(ir => ir.groupId === group.groupId && ir.status === "evaluated").length || 0
                          return `${evaluatedItems}/${totalItems} itens calculados`
                        })()
                      }
                    </p>
                  </div>
                </div>
              </AccordionTrigger>

              {/* Group score badge - positioned absolutely before the expand arrow */}
              <div className="absolute right-10 top-1/2 -translate-y-1/2 flex items-center gap-2 pointer-events-none">
                {group.groupScore === 0 ? (
                  <span className="text-red-500 font-bold text-lg pointer-events-auto">—</span>
                ) : (
                  <Badge className={`${getScoreColor(group.groupScore)} text-sm pointer-events-auto`}>
                    {group.groupScore.toFixed(1)}%
                  </Badge>
                )}
              </div>
            </div>

            {/* Subgroups */}
            <AccordionContent className="px-0 pb-0">
              {group.subgroups.size > 0 && (
                <div className="p-2">
                  <Accordion
                    type="multiple"
                    className="space-y-1.5"
                    value={openSubgroups}
                    onValueChange={setOpenSubgroups}
                  >
                    {Array.from(group.subgroups.values())
                      .sort((a, b) => {
                        // Get subgroup order from first item's subgroup
                        const orderA = a.items[0]?.item?.subgroup?.order ?? 0
                        const orderB = b.items[0]?.item?.subgroup?.order ?? 0
                        return orderA - orderB
                      })
                      .map((subgroup) => (
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
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
      )}
    </div>
  )
}
