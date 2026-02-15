'use client'

import { useState, useEffect, useRef } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip'
import { HelpCircle, CheckCircle2, AlertCircle } from 'lucide-react'
import { cn } from '@/lib/utils'
import { toast } from 'sonner'
import type { AnamnesisTemplate, AnamnesisTemplateItem } from '@/lib/api/anamnesis-templates'
import type { ScoreGroup, ScoreSubgroup, ScoreItem, ScoreLevel } from '@/lib/api/score-api'
import type { AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'

interface AnamnesisTemplateItemsRendererProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
  compact?: boolean // For regular form (smaller UI)
  focusScoreItemId?: string | null
}

// Score level color classes - must be complete strings for Tailwind to detect them
// For full-size buttons (non-compact mode)
const LEVEL_SELECTED_CLASSES: Record<number, string> = {
  0: 'bg-red-100 text-red-900 border-red-500 shadow-md scale-105',
  1: 'bg-orange-100 text-orange-900 border-orange-500 shadow-md scale-105',
  2: 'bg-yellow-100 text-yellow-900 border-yellow-500 shadow-md scale-105',
  3: 'bg-blue-100 text-blue-900 border-blue-500 shadow-md scale-105',
  4: 'bg-green-100 text-green-900 border-green-500 shadow-md scale-105',
  5: 'bg-emerald-100 text-emerald-900 border-emerald-500 shadow-md scale-105',
  6: 'bg-gray-100 text-gray-900 border-gray-500 shadow-md scale-105',
}

const LEVEL_HOVER_CLASSES: Record<number, string> = {
  0: 'bg-background border-border hover:bg-red-200',
  1: 'bg-background border-border hover:bg-orange-200',
  2: 'bg-background border-border hover:bg-yellow-200',
  3: 'bg-background border-border hover:bg-blue-200',
  4: 'bg-background border-border hover:bg-green-200',
  5: 'bg-background border-border hover:bg-emerald-200',
  6: 'bg-background border-border hover:bg-gray-200',
}

// For compact mode buttons
const LEVEL_COMPACT_SELECTED_CLASSES: Record<number, string> = {
  0: 'border-red-500 bg-red-100 text-red-900',
  1: 'border-orange-500 bg-orange-100 text-orange-900',
  2: 'border-yellow-500 bg-yellow-100 text-yellow-900',
  3: 'border-blue-500 bg-blue-100 text-blue-900',
  4: 'border-green-500 bg-green-100 text-green-900',
  5: 'border-emerald-500 bg-emerald-100 text-emerald-900',
  6: 'border-gray-500 bg-gray-100 text-gray-900',
}

// Organize items by group and subgroup
interface OrganizedData {
  groups: Map<string, {
    group: ScoreGroup
    subgroups: Map<string, {
      subgroup: ScoreSubgroup
      items: { templateItem: AnamnesisTemplateItem; scoreItem: ScoreItem }[]
    }>
  }>
}

function organizeTemplateItems(template: AnamnesisTemplate): OrganizedData {
  const groups = new Map()

  if (!template.items) {
    return { groups }
  }

  const sortedItems = [...template.items].sort((a, b) => a.order - b.order)

  sortedItems.forEach((templateItem) => {
    const scoreItem = templateItem.scoreItem

    if (!scoreItem || !scoreItem.subgroup) {
      return
    }

    const subgroup = scoreItem.subgroup
    const group = subgroup.group

    if (!group) {
      return
    }

    if (!groups.has(group.id)) {
      groups.set(group.id, {
        group,
        subgroups: new Map(),
      })
    }

    const groupData = groups.get(group.id)

    if (!groupData.subgroups.has(subgroup.id)) {
      groupData.subgroups.set(subgroup.id, {
        subgroup,
        items: [],
      })
    }

    const subgroupData = groupData.subgroups.get(subgroup.id)
    subgroupData.items.push({ templateItem, scoreItem })
  })

  return { groups }
}

export function AnamnesisTemplateItemsRenderer({
  template,
  initialValues = [],
  onChange,
  compact = false,
  focusScoreItemId,
}: AnamnesisTemplateItemsRendererProps) {
  const [values, setValues] = useState<Map<string, AnamnesisItemFormValue>>(() => {
    const newValues = new Map<string, AnamnesisItemFormValue>()
    initialValues.forEach((val) => {
      newValues.set(val.scoreItemId, val)
    })
    return newValues
  })
  const [focusedItemId, setFocusedItemId] = useState<string | null>(null)

  const organized = organizeTemplateItems(template)
  const itemRefs = useRef<Map<string, HTMLDivElement>>(new Map())
  const hasFocused = useRef(false)

  // Auto-scroll and focus on specific item when coming from health score edit
  useEffect(() => {
    if (focusScoreItemId && !hasFocused.current) {
      hasFocused.current = true

      setTimeout(() => {
        const itemElement = itemRefs.current.get(focusScoreItemId)
        if (itemElement) {
          itemElement.scrollIntoView({ behavior: 'smooth', block: 'center' })

          // Find first focusable input in the item and focus it
          const input = itemElement.querySelector<HTMLInputElement | HTMLButtonElement>(
            'input, button[data-level-button]'
          )
          if (input) {
            setTimeout(() => {
              input.focus()
              if (input instanceof HTMLInputElement) {
                input.select()
              }
            }, 400)
          }
        }
      }, 600)
    }
  }, [focusScoreItemId])

  // Handle level selection (for items without maxSelect)
  const handleLevelSelect = (scoreItemId: string, level: ScoreLevel, order: number) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      if (existing && existing.numericValue === level.level) {
        // Deselect if clicking the same level
        if (!existing.textValue) {
          newValues.delete(scoreItemId)
        } else {
          newValues.set(scoreItemId, { ...existing, numericValue: undefined })
        }
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: level.level,
          textValue: existing?.textValue,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  // Handle text change
  const handleTextChange = (scoreItemId: string, text: string, order: number) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      if (!text && !existing?.numericValue) {
        newValues.delete(scoreItemId)
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: existing?.numericValue,
          textValue: text || undefined,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  // Handle checkbox selection for maxSelect subgroups
  const handleCheckboxSelect = (
    scoreItemId: string,
    order: number,
    subgroupId: string,
    maxSelect: number,
    items: { templateItem: AnamnesisTemplateItem; scoreItem: ScoreItem }[]
  ) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      // Get all selected items in this subgroup
      const subgroupItemIds = items.map(({ scoreItem }) => scoreItem.id)
      const selectedInSubgroup = Array.from(newValues.entries())
        .filter(([id]) => subgroupItemIds.includes(id))
        .filter(([, val]) => val.numericValue !== undefined)

      if (existing && existing.numericValue !== undefined) {
        // Deselect
        newValues.delete(scoreItemId)
      } else {
        // Check if we can select more
        if (selectedInSubgroup.length >= maxSelect) {
          // Show toast warning
          toast.warning(`Você pode selecionar no máximo ${maxSelect} ${maxSelect === 1 ? 'item' : 'itens'}`, {
            duration: 3000,
          })
          return prev
        }
        // Select with numericValue = 1 (represents "selected")
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: 1,
          textValue: existing?.textValue,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  if (!template.items || template.items.length === 0) {
    return (
      <div className="text-center text-muted-foreground py-8">
        <p className={compact ? 'text-sm' : 'text-base'}>
          Este template não possui items configurados
        </p>
      </div>
    )
  }

  return (
    <TooltipProvider delayDuration={300}>
      <div className="space-y-6">
        {Array.from(organized.groups.values()).map(({ group, subgroups }) => (
          <Card key={group.id} className="overflow-hidden border-2">
            <CardHeader className={cn('bg-muted/50', compact ? 'pb-3' : 'pb-4')}>
              <CardTitle className={cn(compact ? 'text-base' : 'text-lg', 'font-semibold flex items-center gap-2')}>
                {group.name}
                <Badge variant="outline" className="ml-auto">
                  {Array.from(subgroups.values()).reduce(
                    (acc, { items }) => acc + items.filter(({ scoreItem }) => values.has(scoreItem.id)).length,
                    0
                  )} / {Array.from(subgroups.values()).reduce((acc, { items }) => acc + items.length, 0)}
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent className={cn(compact ? 'pt-4 space-y-4' : 'pt-6 space-y-6')}>
              {Array.from(subgroups.values()).map(({ subgroup, items }) => {
                const hasMaxSelect = subgroup.maxSelect > 0

                // Count selected items in this subgroup
                const selectedInSubgroup = items.filter(({ scoreItem }) => {
                  const val = values.get(scoreItem.id)
                  return val?.numericValue !== undefined
                })
                const selectedCount = selectedInSubgroup.length

                return (
                  <div key={subgroup.id} className="space-y-3">
                    <div className="flex items-center justify-between border-b pb-2">
                      <h4 className={cn(
                        'font-semibold text-primary',
                        compact ? 'text-sm' : 'text-base'
                      )}>
                        {subgroup.name}
                      </h4>
                      {hasMaxSelect && (
                        <Badge variant={selectedCount >= subgroup.maxSelect ? 'default' : 'outline'} className="gap-1.5">
                          {selectedCount} / {subgroup.maxSelect}
                        </Badge>
                      )}
                    </div>

                    {hasMaxSelect && selectedCount >= subgroup.maxSelect && (
                      <div className="flex items-center gap-2 p-3 bg-blue-50 text-blue-900 rounded-lg text-sm">
                        <AlertCircle className="h-4 w-4 flex-shrink-0" />
                        <p>Máximo atingido. Desmarque um item para selecionar outro.</p>
                      </div>
                    )}

                    <div className={cn('space-y-3', hasMaxSelect && 'space-y-2')}>
                      {items.map(({ templateItem, scoreItem }) => {
                        const currentValue = values.get(scoreItem.id)
                        const levels = scoreItem.levels || []
                        const hasLevelSelected = currentValue?.numericValue !== undefined
                        const isFilled = !!(hasLevelSelected || currentValue?.textValue)
                        const isSelected = hasMaxSelect && hasLevelSelected

                        // Compact layout for maxSelect items
                        if (hasMaxSelect) {
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => {
                                if (el) itemRefs.current.set(scoreItem.id, el)
                              }}
                              className={cn(
                                'flex items-center gap-3 p-3 border rounded-lg transition-all',
                                isSelected ? 'border-primary bg-primary/5' : 'border-border'
                              )}
                            >
                              <Checkbox
                                checked={isSelected}
                                onCheckedChange={() =>
                                  handleCheckboxSelect(
                                    scoreItem.id,
                                    templateItem.order,
                                    subgroup.id,
                                    subgroup.maxSelect,
                                    items
                                  )
                                }
                                className="h-5 w-5 flex-shrink-0"
                              />
                              <Label className={cn(
                                'font-medium flex-shrink-0 cursor-pointer min-w-[200px]',
                                compact ? 'text-sm' : 'text-base'
                              )}>
                                {scoreItem.name}
                              </Label>
                              <Input
                                value={currentValue?.textValue || ''}
                                onChange={(e) =>
                                  handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                                }
                                placeholder="Observações..."
                                className={cn('flex-1', compact ? 'h-9 text-sm' : 'h-10')}
                              />
                              {scoreItem.patientExplanation && (
                                <Tooltip>
                                  <TooltipTrigger asChild>
                                    <button
                                      type="button"
                                      className="text-muted-foreground hover:text-foreground transition-colors flex-shrink-0"
                                    >
                                      <HelpCircle className="h-4 w-4" />
                                    </button>
                                  </TooltipTrigger>
                                  <TooltipContent className="max-w-md">
                                    <p className="text-sm">{scoreItem.patientExplanation}</p>
                                  </TooltipContent>
                                </Tooltip>
                              )}
                            </div>
                          )
                        }

                        // Regular layout for non-maxSelect items
                        const isThisItemFocused = focusedItemId === scoreItem.id

                        return (
                          <Card
                            key={templateItem.id}
                            ref={(el) => {
                              if (el) itemRefs.current.set(scoreItem.id, el)
                            }}
                            className={cn(
                              'border-2 transition-all',
                              isThisItemFocused ? 'border-primary bg-primary/5' : 'border-border'
                            )}
                            onFocus={() => setFocusedItemId(scoreItem.id)}
                            onBlur={() => setFocusedItemId(null)}
                            tabIndex={-1}
                          >
                            <CardContent className={cn(compact ? 'p-4 space-y-3' : 'p-5 space-y-4')}>
                              {/* Item header */}
                              <div className="flex items-start justify-between gap-3">
                                <div className="flex items-center gap-3 flex-1">
                                  {isFilled && !compact && (
                                    <CheckCircle2 className="h-5 w-5 text-primary flex-shrink-0" />
                                  )}
                                  <div className="flex-1">
                                    <Label className={cn(
                                      'font-semibold cursor-pointer',
                                      compact ? 'text-sm' : 'text-base'
                                    )}>
                                      {scoreItem.name}
                                      {scoreItem.unit && (
                                        <span className="text-muted-foreground ml-2 font-normal">
                                          ({scoreItem.unit})
                                        </span>
                                      )}
                                    </Label>
                                  </div>
                                </div>
                                {scoreItem.patientExplanation && (
                                  <Tooltip>
                                    <TooltipTrigger asChild>
                                      <button
                                        type="button"
                                        className="text-muted-foreground hover:text-foreground transition-colors p-1"
                                      >
                                        <HelpCircle className={compact ? 'h-4 w-4' : 'h-5 w-5'} />
                                      </button>
                                    </TooltipTrigger>
                                    <TooltipContent className="max-w-md">
                                      <p className="text-sm">{scoreItem.patientExplanation}</p>
                                    </TooltipContent>
                                  </Tooltip>
                                )}
                              </div>

                              {/* Levels */}
                              {levels.length > 0 && (
                                <div className="space-y-2">
                                  <Label className="text-xs font-medium text-muted-foreground">
                                    {compact ? 'Níveis (opcional):' : 'Selecione o nível:'}
                                  </Label>
                                  <div className={cn(
                                    compact ? 'flex flex-wrap gap-2' : 'grid grid-cols-2 sm:grid-cols-3 gap-3'
                                  )}>
                                    {levels.map((level) => {
                                      const isLevelSelected = currentValue?.numericValue === level.level
                                      const selectedClass = LEVEL_SELECTED_CLASSES[level.level] || LEVEL_SELECTED_CLASSES[6]
                                      const hoverClass = LEVEL_HOVER_CLASSES[level.level] || LEVEL_HOVER_CLASSES[6]
                                      const compactSelectedClass = LEVEL_COMPACT_SELECTED_CLASSES[level.level] || LEVEL_COMPACT_SELECTED_CLASSES[6]

                                      return (
                                        <button
                                          key={level.id}
                                          type="button"
                                          data-level-button
                                          onClick={() =>
                                            handleLevelSelect(scoreItem.id, level, templateItem.order)
                                          }
                                          className={cn(
                                            'rounded-lg border-2 text-sm font-medium transition-all',
                                            compact
                                              ? 'px-3 py-1.5'
                                              : 'p-4 border-3 font-bold flex flex-col items-center gap-1 min-h-[80px] justify-center',
                                            isLevelSelected
                                              ? compact
                                                ? compactSelectedClass
                                                : selectedClass
                                              : compact
                                                ? 'border-border hover:border-primary/50 bg-background'
                                                : hoverClass
                                          )}
                                        >
                                          {compact ? (
                                            `${level.name} (${level.level})`
                                          ) : (
                                            <>
                                              <span className="text-lg">N{level.level}</span>
                                              <span className="text-xs text-center leading-tight">
                                                {level.name}
                                              </span>
                                            </>
                                          )}
                                        </button>
                                      )
                                    })}
                                  </div>
                                </div>
                              )}

                              {/* Text input */}
                              <div className="space-y-2">
                                <Label className="text-xs font-medium text-muted-foreground">
                                  Observações{compact ? ' (opcional)' : ' adicionais'}:
                                </Label>
                                <Textarea
                                  value={currentValue?.textValue || ''}
                                  onChange={(e) =>
                                    handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                                  }
                                  placeholder={compact ? 'Digite observações adicionais...' : 'Digite observações sobre este item...'}
                                  rows={compact ? 2 : 3}
                                  className={cn('resize-none', compact ? 'text-sm' : 'text-base')}
                                />
                              </div>

                              {/* Status badges */}
                              {isFilled && (
                                <div className="flex gap-2 pt-2 border-t">
                                  {currentValue.numericValue !== undefined && (
                                    <Badge variant={compact ? 'secondary' : 'default'} className="text-xs">
                                      {compact ? `Nível: ${currentValue.numericValue}` : `Nível selecionado: ${currentValue.numericValue}`}
                                    </Badge>
                                  )}
                                  {currentValue.textValue && (
                                    <Badge variant="secondary" className="text-xs">
                                      Com observações
                                    </Badge>
                                  )}
                                </div>
                              )}
                            </CardContent>
                          </Card>
                        )
                      })}
                    </div>
                  </div>
                )
              })}
            </CardContent>
          </Card>
        ))}
      </div>
    </TooltipProvider>
  )
}
