'use client'

import { useState, useEffect } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { Badge } from '@/components/ui/badge'
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip'
import { HelpCircle } from 'lucide-react'
import { cn } from '@/lib/utils'
import type { AnamnesisTemplate, AnamnesisTemplateItem } from '@/lib/api/anamnesis-templates'
import type { ScoreGroup, ScoreSubgroup, ScoreItem, ScoreLevel } from '@/lib/api/score-api'

// Type for the form values
export interface AnamnesisItemFormValue {
  scoreItemId: string
  numericValue?: number
  textValue?: string
  order: number
}

interface AnamnesisTemplateItemsFormProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
}

// Group colors (same as score poster)
const GROUP_COLORS: Record<string, { bg: string; text: string; border: string }> = {
  default: { bg: 'bg-gray-50', text: 'text-gray-900', border: 'border-gray-200' },
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

  // Sort by order
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

    // Get or create group
    if (!groups.has(group.id)) {
      groups.set(group.id, {
        group,
        subgroups: new Map(),
      })
    }

    const groupData = groups.get(group.id)

    // Get or create subgroup
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

export function AnamnesisTemplateItemsForm({
  template,
  initialValues = [],
  onChange,
}: AnamnesisTemplateItemsFormProps) {
  const [values, setValues] = useState<Map<string, AnamnesisItemFormValue>>(() => {
    // Initialize state from initialValues on mount
    const newValues = new Map<string, AnamnesisItemFormValue>()
    initialValues.forEach((val) => {
      newValues.set(val.scoreItemId, val)
    })
    return newValues
  })

  // Reset values when template changes (template.id is in the component key)
  // No useEffect needed - the component will remount due to key prop

  // Organize data
  const organized = organizeTemplateItems(template)

  // Handle level selection
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

      // Notify parent immediately
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
        // Remove if both text and level are empty
        newValues.delete(scoreItemId)
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: existing?.numericValue,
          textValue: text || undefined,
          order,
        })
      }

      // Notify parent immediately
      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  if (!template.items || template.items.length === 0) {
    return (
      <div className="text-center text-muted-foreground py-8">
        Este template não possui items configurados
      </div>
    )
  }

  return (
    <TooltipProvider delayDuration={300}>
      <div className="space-y-6">
      {Array.from(organized.groups.values()).map(({ group, subgroups }) => {
        const colorScheme = GROUP_COLORS.default

        return (
          <Card key={group.id} className={cn('overflow-hidden', colorScheme.border)}>
            <CardHeader className={cn(colorScheme.bg, 'pb-3')}>
              <CardTitle className={cn('text-lg', colorScheme.text)}>
                {group.name}
              </CardTitle>
            </CardHeader>
            <CardContent className="pt-4 space-y-4">
              {Array.from(subgroups.values()).map(({ subgroup, items }) => (
                <div key={subgroup.id} className="space-y-3">
                  <h4 className="font-medium text-sm text-muted-foreground">
                    {subgroup.name}
                  </h4>

                  <div className="space-y-4">
                    {items.map(({ templateItem, scoreItem }) => {
                      const currentValue = values.get(scoreItem.id)
                      const levels = scoreItem.levels || []

                      return (
                        <div
                          key={templateItem.id}
                          className="border rounded-lg p-4 space-y-3 bg-card"
                        >
                          <div className="flex items-start justify-between gap-2">
                            <div className="flex items-center gap-2 flex-1">
                              <Label className="text-sm font-medium">
                                {scoreItem.name}
                                {scoreItem.unit && (
                                  <span className="text-muted-foreground ml-1">
                                    ({scoreItem.unit})
                                  </span>
                                )}
                              </Label>
                              {scoreItem.patientExplanation && (
                                <Tooltip>
                                  <TooltipTrigger asChild>
                                    <button
                                      type="button"
                                      className="text-muted-foreground hover:text-foreground transition-colors"
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
                          </div>

                          {/* Levels */}
                          {levels.length > 0 && (
                            <div className="space-y-2">
                              <Label className="text-xs text-muted-foreground">
                                Níveis (opcional):
                              </Label>
                              <div className="flex flex-wrap gap-2">
                                {levels.map((level) => {
                                  const isSelected = currentValue?.numericValue === level.level

                                  return (
                                    <button
                                      key={level.id}
                                      type="button"
                                      onClick={() =>
                                        handleLevelSelect(scoreItem.id, level, templateItem.order)
                                      }
                                      className={cn(
                                        'px-3 py-1.5 rounded-md border-2 text-sm font-medium transition-all',
                                        isSelected
                                          ? 'border-primary bg-primary text-primary-foreground'
                                          : 'border-border hover:border-primary/50 bg-background'
                                      )}
                                    >
                                      {level.name} ({level.level})
                                    </button>
                                  )
                                })}
                              </div>
                            </div>
                          )}

                          {/* Text input */}
                          <div className="space-y-2">
                            <Label className="text-xs text-muted-foreground">
                              Observações (opcional):
                            </Label>
                            <Textarea
                              value={currentValue?.textValue || ''}
                              onChange={(e) =>
                                handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                              }
                              placeholder="Digite observações adicionais..."
                              rows={2}
                              className="text-sm"
                            />
                          </div>

                          {/* Status indicator */}
                          {currentValue && (currentValue.numericValue !== undefined || currentValue.textValue) && (
                            <div className="flex gap-2">
                              {currentValue.numericValue !== undefined && (
                                <Badge variant="secondary" className="text-xs">
                                  Nível: {currentValue.numericValue}
                                </Badge>
                              )}
                              {currentValue.textValue && (
                                <Badge variant="secondary" className="text-xs">
                                  Com observações
                                </Badge>
                              )}
                            </div>
                          )}
                        </div>
                      )
                    })}
                  </div>
                </div>
              ))}
            </CardContent>
          </Card>
        )
      })}
      </div>
    </TooltipProvider>
  )
}
