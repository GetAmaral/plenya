'use client'

import { useState, useMemo } from 'react'
import { Label } from '@/components/ui/label'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { Input } from '@/components/ui/input'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Separator } from '@/components/ui/separator'
import { Search, ChevronDown, ChevronRight } from 'lucide-react'
import { useAllScoreGroupTrees } from '@/lib/api/score-api'

interface ArticleScoreItemsSelectorProps {
  selectedItemIds: string[]
  onChange: (itemIds: string[]) => void
}

export function ArticleScoreItemsSelector({
  selectedItemIds,
  onChange,
}: ArticleScoreItemsSelectorProps) {
  const { data: groups = [], isLoading } = useAllScoreGroupTrees()
  const [searchQuery, setSearchQuery] = useState('')
  const [expandedGroups, setExpandedGroups] = useState<Record<string, boolean>>({})
  const [expandedSubgroups, setExpandedSubgroups] = useState<Record<string, boolean>>({})

  // Coletar todos os items em uma lista plana para contagem
  const allItems = useMemo(() => {
    const items: Array<{ id: string; name: string; groupName: string; subgroupName: string }> = []
    groups.forEach((group) => {
      group.subgroups?.forEach((subgroup) => {
        subgroup.items?.forEach((item) => {
          items.push({
            id: item.id,
            name: item.name,
            groupName: group.name,
            subgroupName: subgroup.name,
          })
        })
      })
    })
    return items
  }, [groups])

  // Filtrar items por busca
  const filteredGroups = useMemo(() => {
    if (!searchQuery) return groups

    return groups
      .map((group) => ({
        ...group,
        subgroups: group.subgroups
          ?.map((subgroup) => ({
            ...subgroup,
            items: subgroup.items?.filter(
              (item) =>
                item.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                subgroup.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                group.name.toLowerCase().includes(searchQuery.toLowerCase())
            ),
          }))
          .filter((subgroup) => subgroup.items && subgroup.items.length > 0),
      }))
      .filter((group) => group.subgroups && group.subgroups.length > 0)
  }, [groups, searchQuery])

  const toggleGroup = (groupId: string) => {
    setExpandedGroups((prev) => ({ ...prev, [groupId]: !prev[groupId] }))
  }

  const toggleSubgroup = (subgroupId: string) => {
    setExpandedSubgroups((prev) => ({ ...prev, [subgroupId]: !prev[subgroupId] }))
  }

  const toggleItem = (itemId: string) => {
    if (selectedItemIds.includes(itemId)) {
      onChange(selectedItemIds.filter((id) => id !== itemId))
    } else {
      onChange([...selectedItemIds, itemId])
    }
  }

  const toggleAllInSubgroup = (subgroupId: string, itemIds: string[]) => {
    const allSelected = itemIds.every((id) => selectedItemIds.includes(id))
    if (allSelected) {
      onChange(selectedItemIds.filter((id) => !itemIds.includes(id)))
    } else {
      const newIds = itemIds.filter((id) => !selectedItemIds.includes(id))
      onChange([...selectedItemIds, ...newIds])
    }
  }

  if (isLoading) {
    return (
      <div className="flex items-center justify-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
      </div>
    )
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <Label>Score Items Vinculados</Label>
        <Badge variant="secondary">
          {selectedItemIds.length} de {allItems.length} selecionados
        </Badge>
      </div>

      {/* Search */}
      <div className="relative">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input
          placeholder="Buscar items..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="pl-9"
        />
      </div>

      {/* Items Tree */}
      <ScrollArea className="h-[400px] rounded-md border">
        <div className="p-4 space-y-3">
          {filteredGroups.length === 0 ? (
            <p className="text-sm text-muted-foreground text-center py-8">
              Nenhum item encontrado
            </p>
          ) : (
            filteredGroups.map((group) => (
              <div key={group.id} className="space-y-2">
                {/* Group Header */}
                <button
                  type="button"
                  onClick={() => toggleGroup(group.id)}
                  className="flex items-center gap-2 w-full text-left hover:bg-muted/50 rounded-md px-2 py-1.5 transition-colors"
                >
                  {expandedGroups[group.id] ? (
                    <ChevronDown className="h-4 w-4 text-muted-foreground" />
                  ) : (
                    <ChevronRight className="h-4 w-4 text-muted-foreground" />
                  )}
                  <span className="font-semibold text-sm">{group.name}</span>
                  <Badge variant="outline" className="ml-auto text-xs">
                    {group.subgroups?.reduce(
                      (acc, sg) => acc + (sg.items?.length || 0),
                      0
                    )}{' '}
                    items
                  </Badge>
                </button>

                {/* Subgroups */}
                {expandedGroups[group.id] && group.subgroups && (
                  <div className="ml-6 space-y-2">
                    {group.subgroups.map((subgroup) => {
                      const subgroupItemIds = subgroup.items?.map((item) => item.id) || []
                      const allSelected = subgroupItemIds.every((id) =>
                        selectedItemIds.includes(id)
                      )
                      const someSelected = subgroupItemIds.some((id) =>
                        selectedItemIds.includes(id)
                      )

                      return (
                        <div key={subgroup.id} className="space-y-2">
                          {/* Subgroup Header */}
                          <div className="flex items-center gap-2">
                            <Checkbox
                              checked={allSelected}
                              onCheckedChange={() =>
                                toggleAllInSubgroup(subgroup.id, subgroupItemIds)
                              }
                              className={someSelected && !allSelected ? 'opacity-50' : ''}
                            />
                            <button
                              type="button"
                              onClick={() => toggleSubgroup(subgroup.id)}
                              className="flex items-center gap-2 flex-1 text-left hover:bg-muted/30 rounded-md px-2 py-1 transition-colors"
                            >
                              {expandedSubgroups[subgroup.id] ? (
                                <ChevronDown className="h-3 w-3 text-muted-foreground" />
                              ) : (
                                <ChevronRight className="h-3 w-3 text-muted-foreground" />
                              )}
                              <span className="font-medium text-sm">{subgroup.name}</span>
                              <Badge variant="secondary" className="ml-auto text-xs">
                                {subgroup.items?.length || 0}
                              </Badge>
                            </button>
                          </div>

                          {/* Items */}
                          {expandedSubgroups[subgroup.id] && subgroup.items && (
                            <div className="ml-6 space-y-1">
                              {subgroup.items.map((item) => (
                                <div
                                  key={item.id}
                                  className="flex items-center gap-2 hover:bg-muted/30 rounded-md px-2 py-1.5 transition-colors"
                                >
                                  <Checkbox
                                    checked={selectedItemIds.includes(item.id)}
                                    onCheckedChange={() => toggleItem(item.id)}
                                  />
                                  <div className="flex-1">
                                    <p className="text-sm">{item.name}</p>
                                    {item.unit && (
                                      <p className="text-xs text-muted-foreground">
                                        {item.unit}
                                      </p>
                                    )}
                                  </div>
                                  {item.points > 0 && (
                                    <Badge variant="outline" className="text-xs">
                                      {item.points}pt
                                    </Badge>
                                  )}
                                </div>
                              ))}
                            </div>
                          )}
                        </div>
                      )
                    })}
                  </div>
                )}

                <Separator className="my-2" />
              </div>
            ))
          )}
        </div>
      </ScrollArea>

      <p className="text-xs text-muted-foreground">
        Selecione os score items que este artigo científico suporta ou referencia
      </p>
    </div>
  )
}
