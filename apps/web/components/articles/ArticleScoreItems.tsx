'use client'

import Link from 'next/link'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Network, ChevronRight } from 'lucide-react'
import { ScoreItem } from '@/lib/api/article-api'

interface ArticleScoreItemsProps {
  scoreItems: ScoreItem[]
}

export function ArticleScoreItems({ scoreItems }: ArticleScoreItemsProps) {
  if (!scoreItems || scoreItems.length === 0) {
    return null
  }

  // Agrupar items por grupo e subgrupo
  const groupedItems = scoreItems.reduce((acc, item) => {
    if (!item.subgroup) return acc

    const groupId = item.subgroup.group?.id || 'ungrouped'
    const groupName = item.subgroup.group?.name || 'Sem Grupo'
    const subgroupId = item.subgroup.id
    const subgroupName = item.subgroup.name

    if (!acc[groupId]) {
      acc[groupId] = {
        id: groupId,
        name: groupName,
        subgroups: {}
      }
    }

    if (!acc[groupId].subgroups[subgroupId]) {
      acc[groupId].subgroups[subgroupId] = {
        id: subgroupId,
        name: subgroupName,
        items: []
      }
    }

    acc[groupId].subgroups[subgroupId].items.push(item)

    return acc
  }, {} as Record<string, {
    id: string
    name: string
    subgroups: Record<string, {
      id: string
      name: string
      items: ScoreItem[]
    }>
  }>)

  const groups = Object.values(groupedItems)

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <div>
            <CardTitle className="text-lg flex items-center gap-2">
              <Network className="h-5 w-5" />
              Score Items Vinculados
            </CardTitle>
            <CardDescription>
              Este artigo está vinculado a {scoreItems.length} item{scoreItems.length !== 1 ? 's' : ''} do sistema de escores
            </CardDescription>
          </div>
          <Link
            href="/scores"
            className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-1"
          >
            Ver todos os escores
            <ChevronRight className="h-4 w-4" />
          </Link>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-6">
          {groups.map((group) => (
            <div key={group.id} className="space-y-3">
              {/* Cabeçalho do Grupo */}
              <div className="flex items-center gap-2">
                <div className="h-1 w-1 rounded-full bg-primary" />
                <h3 className="font-semibold text-base">{group.name}</h3>
              </div>

              {/* Subgrupos */}
              <div className="ml-3 space-y-4">
                {Object.values(group.subgroups).map((subgroup) => (
                  <div key={subgroup.id} className="space-y-2">
                    {/* Cabeçalho do Subgrupo */}
                    <div className="flex items-center gap-2 pl-3 border-l-2 border-muted">
                      <h4 className="font-medium text-sm text-muted-foreground">
                        {subgroup.name}
                      </h4>
                      <Badge variant="secondary" className="text-xs">
                        {subgroup.items.length} {subgroup.items.length === 1 ? 'item' : 'itens'}
                      </Badge>
                    </div>

                    {/* Items */}
                    <div className="ml-3 space-y-1.5">
                      {subgroup.items.map((item) => (
                        <div
                          key={item.id}
                          className="flex items-center justify-between py-2 px-3 rounded-md hover:bg-muted/50 transition-colors border border-transparent hover:border-muted"
                        >
                          <div className="flex items-center gap-3">
                            <div className="h-1.5 w-1.5 rounded-full bg-muted-foreground/50" />
                            <div>
                              <p className="text-sm font-medium">{item.name}</p>
                              {item.description && (
                                <p className="text-xs text-muted-foreground line-clamp-1">
                                  {item.description}
                                </p>
                              )}
                            </div>
                          </div>
                          <div className="flex items-center gap-2">
                            {item.unit && (
                              <Badge variant="outline" className="text-xs font-mono">
                                {item.unit}
                              </Badge>
                            )}
                            <Badge variant="default" className="text-xs">
                              {item.points}pt
                            </Badge>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>

        {/* Link para mindmap */}
        <div className="mt-6 pt-4 border-t">
          <Link
            href="/scores/mindmap"
            className="inline-flex items-center gap-2 text-sm text-primary hover:underline"
          >
            <Network className="h-4 w-4" />
            Visualizar no Mindmap de Escores
          </Link>
        </div>
      </CardContent>
    </Card>
  )
}
