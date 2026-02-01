'use client'

import { useState, useEffect, useMemo } from 'react'
import { Search, X, ChevronRight } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Badge } from '@/components/ui/badge'
import { cn, normalizeForSearch } from '@/lib/utils'
import { ScoreGroup } from '@/lib/api/score-api'

export interface SearchResult {
  type: 'group' | 'subgroup' | 'item' | 'level'
  id: string
  name: string
  path: string[] // Caminho completo: [groupName, subgroupName?, itemName?]
  groupId: string
  subgroupId?: string
  itemId?: string
  levelId?: string
}

interface ScoreSearchProps {
  scoreGroups: ScoreGroup[]
  isOpen: boolean
  onClose: () => void
  onResultClick: (result: SearchResult) => void
}

export function ScoreSearch({ scoreGroups, isOpen, onClose, onResultClick }: ScoreSearchProps) {
  const [searchQuery, setSearchQuery] = useState('')
  const [selectedIndex, setSelectedIndex] = useState(0)

  // Buscar em todos os níveis
  const searchResults = useMemo(() => {
    if (!searchQuery.trim() || !scoreGroups) return []

    const results: SearchResult[] = []
    const query = normalizeForSearch(searchQuery)

    scoreGroups.forEach(group => {
      // Buscar no grupo
      if (normalizeForSearch(group.name).includes(query)) {
        results.push({
          type: 'group',
          id: `group-${group.id}`,
          name: group.name,
          path: [group.name],
          groupId: group.id,
        })
      }

      // Buscar nos subgrupos
      group.subgroups?.forEach(subgroup => {
        if (normalizeForSearch(subgroup.name).includes(query)) {
          results.push({
            type: 'subgroup',
            id: `subgroup-${subgroup.id}`,
            name: subgroup.name,
            path: [group.name, subgroup.name],
            groupId: group.id,
            subgroupId: subgroup.id,
          })
        }

        // Buscar nos items
        subgroup.items?.forEach(item => {
          if (normalizeForSearch(item.name).includes(query) || normalizeForSearch(item.unit || '').includes(query)) {
            results.push({
              type: 'item',
              id: `item-${item.id}`,
              name: item.name,
              path: [group.name, subgroup.name, item.name],
              groupId: group.id,
              subgroupId: subgroup.id,
              itemId: item.id,
            })
          }

          // Buscar nos níveis
          item.levels?.forEach(level => {
            if (normalizeForSearch(level.name).includes(query)) {
              results.push({
                type: 'level',
                id: `level-${level.id}`,
                name: level.name,
                path: [group.name, subgroup.name, item.name, `Nível ${level.level}`],
                groupId: group.id,
                subgroupId: subgroup.id,
                itemId: item.id,
                levelId: level.id,
              })
            }
          })
        })
      })
    })

    return results
  }, [searchQuery, scoreGroups])

  // Reset selected index quando mudar a busca
  useEffect(() => {
    setSelectedIndex(0)
  }, [searchQuery])

  // Navegação por teclado
  useEffect(() => {
    if (!isOpen) return

    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'ArrowDown') {
        e.preventDefault()
        setSelectedIndex(prev => Math.min(prev + 1, searchResults.length - 1))
      } else if (e.key === 'ArrowUp') {
        e.preventDefault()
        setSelectedIndex(prev => Math.max(prev - 1, 0))
      } else if (e.key === 'Enter' && searchResults[selectedIndex]) {
        e.preventDefault()
        onResultClick(searchResults[selectedIndex])
      } else if (e.key === 'Escape') {
        e.preventDefault()
        onClose()
      }
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [isOpen, searchResults, selectedIndex, onResultClick, onClose])

  if (!isOpen) return null

  const getTypeColor = (type: SearchResult['type']) => {
    switch (type) {
      case 'group': return 'bg-primary text-primary-foreground'
      case 'subgroup': return 'bg-blue-500 text-white'
      case 'item': return 'bg-purple-500 text-white'
      case 'level': return 'bg-green-500 text-white'
    }
  }

  const getTypeLabel = (type: SearchResult['type']) => {
    switch (type) {
      case 'group': return 'Grupo'
      case 'subgroup': return 'Subgrupo'
      case 'item': return 'Item'
      case 'level': return 'Nível'
    }
  }

  return (
    <div className="w-full bg-background border rounded-lg shadow-2xl">
      {/* Header */}
      <div className="flex items-center gap-2 p-3 border-b">
        <Search className="h-4 w-4 text-muted-foreground" />
        <Input
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          placeholder="Buscar em grupos, subgrupos, items e níveis..."
          className="flex-1 border-0 focus-visible:ring-0 px-0"
          autoFocus
        />
        <Badge variant="outline" className="shrink-0">
          {searchResults.length} {searchResults.length === 1 ? 'resultado' : 'resultados'}
        </Badge>
        <Button variant="ghost" size="icon" onClick={onClose}>
          <X className="h-4 w-4" />
        </Button>
      </div>

      {/* Results */}
      <div className="max-h-[500px] overflow-y-auto">
        {searchQuery.trim() === '' ? (
          <div className="p-6 text-center text-sm text-muted-foreground">
            Digite para buscar em todos os níveis...
          </div>
        ) : searchResults.length === 0 ? (
          <div className="p-6 text-center text-sm text-muted-foreground">
            Nenhum resultado encontrado para "{searchQuery}"
          </div>
        ) : (
          <div className="divide-y">
            {searchResults.map((result, index) => (
              <button
                key={result.id}
                onClick={() => onResultClick(result)}
                onMouseEnter={() => setSelectedIndex(index)}
                className={cn(
                  "w-full text-left p-3 hover:bg-accent transition-colors",
                  index === selectedIndex && "bg-accent"
                )}
              >
                <div className="flex items-start gap-3">
                  <Badge className={cn("shrink-0 mt-0.5", getTypeColor(result.type))}>
                    {getTypeLabel(result.type)}
                  </Badge>
                  <div className="flex-1 min-w-0">
                    <div className="font-medium text-sm mb-1 truncate">
                      {result.name}
                    </div>
                    <div className="flex items-center gap-1 text-xs text-muted-foreground flex-wrap">
                      {result.path.map((segment, i) => (
                        <div key={i} className="flex items-center gap-1">
                          {i > 0 && <ChevronRight className="h-3 w-3 shrink-0" />}
                          <span className="truncate max-w-[140px]">{segment}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
              </button>
            ))}
          </div>
        )}
      </div>

      {/* Footer with hints */}
      {searchResults.length > 0 && (
        <div className="p-2 border-t bg-muted/50 text-xs text-muted-foreground flex items-center justify-between">
          <div className="flex gap-3">
            <span>↑↓ Navegar</span>
            <span>Enter Selecionar</span>
            <span>Esc Fechar</span>
          </div>
          <span>{selectedIndex + 1} / {searchResults.length}</span>
        </div>
      )}
    </div>
  )
}
