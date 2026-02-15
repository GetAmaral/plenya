'use client'

import { useMemo, useState } from 'react'
import { Search, X } from 'lucide-react'
import type { ScoreItem } from '@plenya/types'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { normalizeForSearch } from '@/lib/utils'
import { DraggableScoreItem } from './DraggableScoreItem'

interface UnassignedItemsPanelProps {
  items: ScoreItem[]
}

interface GroupedItems {
  groupId: string
  groupName: string
  subgroups: {
    subgroupId: string
    subgroupName: string
    items: ScoreItem[]
  }[]
}

export function UnassignedItemsPanel({ items }: UnassignedItemsPanelProps) {
  const [searchQuery, setSearchQuery] = useState('')

  // Agrupar items por grupo e subgrupo (mesma lógica do ScoreTreeView)
  const groupedItems = useMemo(() => {
    const grouped: Record<string, GroupedItems> = {}

    items.forEach(item => {
      if (!item.subgroup?.group) return

      const groupId = item.subgroup.group.id
      const groupName = item.subgroup.group.name
      const subgroupId = item.subgroup.id
      const subgroupName = item.subgroup.name

      if (!grouped[groupId]) {
        grouped[groupId] = {
          groupId,
          groupName,
          subgroups: []
        }
      }

      let subgroup = grouped[groupId].subgroups.find(s => s.subgroupId === subgroupId)
      if (!subgroup) {
        subgroup = {
          subgroupId,
          subgroupName,
          items: []
        }
        grouped[groupId].subgroups.push(subgroup)
      }

      subgroup.items.push(item)
    })

    // Ordenar grupos, subgrupos e items por order (mesma lógica do ScoreTreeView)
    return Object.values(grouped).map(group => ({
      ...group,
      subgroups: group.subgroups
        .map(subgroup => ({
          ...subgroup,
          items: subgroup.items.sort((a, b) => a.order - b.order)
        }))
        .sort((a, b) => {
          // Ordenar subgrupos pela ordem do primeiro item
          const aOrder = a.items[0]?.subgroup?.order ?? 0
          const bOrder = b.items[0]?.subgroup?.order ?? 0
          return aOrder - bOrder
        })
    })).sort((a, b) => {
      // Ordenar grupos pela ordem do primeiro subgrupo
      const aOrder = a.subgroups[0]?.items[0]?.subgroup?.group?.order ?? 0
      const bOrder = b.subgroups[0]?.items[0]?.subgroup?.group?.order ?? 0
      return aOrder - bOrder
    })
  }, [items])

  // Filtrar items pela busca (usando normalizeForSearch - mesma função do ScoreSearch)
  const filteredGroups = useMemo(() => {
    if (!searchQuery.trim()) return groupedItems

    const query = normalizeForSearch(searchQuery)

    return groupedItems
      .map(group => ({
        ...group,
        subgroups: group.subgroups
          .map(subgroup => ({
            ...subgroup,
            items: subgroup.items.filter(item =>
              normalizeForSearch(item.name).includes(query) ||
              normalizeForSearch(item.unit || '').includes(query) ||
              normalizeForSearch(subgroup.subgroupName).includes(query) ||
              normalizeForSearch(group.groupName).includes(query)
            )
          }))
          .filter(subgroup => subgroup.items.length > 0)
      }))
      .filter(group => group.subgroups.length > 0)
  }, [groupedItems, searchQuery])

  const totalItems = items.length
  const filteredItemsCount = filteredGroups.reduce((acc, group) =>
    acc + group.subgroups.reduce((subAcc, subgroup) =>
      subAcc + subgroup.items.length, 0
    ), 0
  )

  if (totalItems === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Itens Não Atribuídos</CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-sm text-muted-foreground text-center py-8">
            Todos os itens estão atribuídos a pilares.
          </p>
        </CardContent>
      </Card>
    )
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between gap-4">
          <div>
            <CardTitle>Itens Não Atribuídos</CardTitle>
            <p className="text-sm text-muted-foreground mt-1">
              {totalItems} {totalItems === 1 ? 'item disponível' : 'itens disponíveis'} para atribuição
            </p>
          </div>

          {/* Search (mesma UI do ScoreSearch) */}
          <div className="flex items-center gap-2 flex-1 max-w-md">
            <div className="relative flex-1">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Buscar items..."
                className="pl-8 pr-8"
              />
              {searchQuery && (
                <Button
                  variant="ghost"
                  size="sm"
                  className="absolute right-0 top-0 h-full px-2"
                  onClick={() => setSearchQuery('')}
                >
                  <X className="h-4 w-4" />
                </Button>
              )}
            </div>
            {searchQuery && (
              <Badge variant="outline">
                {filteredItemsCount} {filteredItemsCount === 1 ? 'resultado' : 'resultados'}
              </Badge>
            )}
          </div>
        </div>
      </CardHeader>

      <CardContent className="max-h-[400px] overflow-y-auto">
        {filteredGroups.length === 0 ? (
          <p className="text-sm text-muted-foreground text-center py-8">
            Nenhum resultado encontrado para "{searchQuery}"
          </p>
        ) : (
          <div className="space-y-2">
            {/* Estrutura igual ao ScoreTreeView: Grupo → Subgrupo → Items */}
            {filteredGroups.map(group => (
              <div key={group.groupId} className="rounded-lg border">
                {/* Group Header (mesma UI do ScoreTreeView) */}
                <div className="px-3 py-2 bg-muted/50">
                  <h3 className="text-sm font-semibold text-left">{group.groupName}</h3>
                  <p className="text-xs text-muted-foreground mt-0.5 text-left">
                    {group.subgroups.reduce((acc, s) => acc + s.items.length, 0)} itens
                  </p>
                </div>

                {/* Subgroups Accordion (mesma UI do ScoreTreeView) */}
                <div className="p-2">
                  <Accordion type="multiple" className="space-y-1.5">
                    {group.subgroups.map(subgroup => (
                      <AccordionItem
                        key={subgroup.subgroupId}
                        value={subgroup.subgroupId}
                        className="border rounded-md"
                      >
                        <AccordionTrigger className="px-2.5 hover:no-underline py-2 text-left">
                          <div className="flex items-center gap-2">
                            <span className="text-sm font-medium text-left">{subgroup.subgroupName}</span>
                            <Badge variant="outline" className="text-[10px] px-1.5 py-0">
                              {subgroup.items.length}
                            </Badge>
                          </div>
                        </AccordionTrigger>

                        <AccordionContent className="px-2.5 pb-2">
                          <div className="space-y-1.5 mt-1">
                            {subgroup.items.map(item => (
                              <DraggableScoreItem
                                key={item.id}
                                item={item}
                              />
                            ))}
                          </div>
                        </AccordionContent>
                      </AccordionItem>
                    ))}
                  </Accordion>
                </div>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
