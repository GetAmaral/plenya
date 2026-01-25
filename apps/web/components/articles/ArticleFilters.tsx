'use client'

import { useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { ArticleFilters as Filters, ARTICLE_TYPES } from '@/lib/api/article-api'
import { Filter, X, Star } from 'lucide-react'

interface ArticleFiltersProps {
  filters: Filters
  onFiltersChange: (filters: Filters) => void
}

export function ArticleFilters({
  filters,
  onFiltersChange,
}: ArticleFiltersProps) {
  const [isExpanded, setIsExpanded] = useState(false)

  const hasActiveFilters =
    filters.journal ||
    filters.specialty ||
    filters.articleType ||
    filters.favorite !== undefined ||
    filters.rating !== undefined ||
    filters.publishedAfter ||
    filters.publishedBefore

  const handleClearFilters = () => {
    onFiltersChange({})
  }

  return (
    <Card>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <CardTitle className="text-base flex items-center gap-2">
            <Filter className="h-4 w-4" />
            Filtros
            {hasActiveFilters && (
              <span className="text-xs font-normal text-muted-foreground">
                ({Object.keys(filters).length} ativos)
              </span>
            )}
          </CardTitle>

          <div className="flex items-center gap-2">
            {hasActiveFilters && (
              <Button
                variant="ghost"
                size="sm"
                onClick={handleClearFilters}
                className="h-8"
              >
                <X className="mr-1 h-3 w-3" />
                Limpar
              </Button>
            )}

            <Button
              variant="outline"
              size="sm"
              onClick={() => setIsExpanded(!isExpanded)}
              className="h-8"
            >
              {isExpanded ? 'Ocultar' : 'Expandir'}
            </Button>
          </div>
        </div>
      </CardHeader>

      {isExpanded && (
        <CardContent className="space-y-4">
          {/* Journal */}
          <div className="space-y-2">
            <Label htmlFor="filter-journal">Revista</Label>
            <Input
              id="filter-journal"
              placeholder="Ex: Nature, Science..."
              value={filters.journal || ''}
              onChange={(e) =>
                onFiltersChange({
                  ...filters,
                  journal: e.target.value || undefined,
                })
              }
            />
          </div>

          {/* Specialty */}
          <div className="space-y-2">
            <Label htmlFor="filter-specialty">Especialidade</Label>
            <Input
              id="filter-specialty"
              placeholder="Ex: Cardiologia, Neurologia..."
              value={filters.specialty || ''}
              onChange={(e) =>
                onFiltersChange({
                  ...filters,
                  specialty: e.target.value || undefined,
                })
              }
            />
          </div>

          {/* Article Type */}
          <div className="space-y-2">
            <Label htmlFor="filter-type">Tipo de Artigo</Label>
            <Select
              value={filters.articleType || 'all'}
              onValueChange={(value) =>
                onFiltersChange({
                  ...filters,
                  articleType: value === 'all' ? undefined : (value as any),
                })
              }
            >
              <SelectTrigger id="filter-type">
                <SelectValue placeholder="Todos os tipos" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">Todos os tipos</SelectItem>
                {ARTICLE_TYPES.map((type) => (
                  <SelectItem key={type.value} value={type.value}>
                    {type.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Rating */}
          <div className="space-y-2">
            <Label htmlFor="filter-rating">Avaliação Mínima</Label>
            <Select
              value={filters.rating?.toString() || 'all'}
              onValueChange={(value) =>
                onFiltersChange({
                  ...filters,
                  rating: value === 'all' ? undefined : parseInt(value),
                })
              }
            >
              <SelectTrigger id="filter-rating">
                <SelectValue placeholder="Qualquer avaliação" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">Qualquer avaliação</SelectItem>
                {[5, 4, 3, 2, 1].map((rating) => (
                  <SelectItem key={rating} value={rating.toString()}>
                    <div className="flex items-center gap-1">
                      {Array.from({ length: rating }).map((_, i) => (
                        <Star
                          key={i}
                          className="h-3 w-3 fill-yellow-400 text-yellow-400"
                        />
                      ))}
                      <span className="ml-1">ou mais</span>
                    </div>
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Favorite Toggle */}
          <div className="flex items-center justify-between">
            <Label htmlFor="filter-favorite" className="cursor-pointer">
              Apenas favoritos
            </Label>
            <Switch
              id="filter-favorite"
              checked={filters.favorite || false}
              onCheckedChange={(checked) =>
                onFiltersChange({
                  ...filters,
                  favorite: checked ? true : undefined,
                })
              }
            />
          </div>

          {/* Date Range */}
          <div className="space-y-3">
            <Label>Período de Publicação</Label>

            <div className="grid grid-cols-2 gap-3">
              <div className="space-y-1.5">
                <Label htmlFor="filter-date-after" className="text-xs text-muted-foreground">
                  De
                </Label>
                <Input
                  id="filter-date-after"
                  type="date"
                  value={filters.publishedAfter || ''}
                  onChange={(e) =>
                    onFiltersChange({
                      ...filters,
                      publishedAfter: e.target.value || undefined,
                    })
                  }
                />
              </div>

              <div className="space-y-1.5">
                <Label htmlFor="filter-date-before" className="text-xs text-muted-foreground">
                  Até
                </Label>
                <Input
                  id="filter-date-before"
                  type="date"
                  value={filters.publishedBefore || ''}
                  onChange={(e) =>
                    onFiltersChange({
                      ...filters,
                      publishedBefore: e.target.value || undefined,
                    })
                  }
                />
              </div>
            </div>
          </div>
        </CardContent>
      )}
    </Card>
  )
}
