'use client'

import { useState, useEffect } from 'react'
import { Sparkles, Search, Loader2, FileText, TrendingUp } from 'lucide-react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { ScrollArea } from '@/components/ui/scroll-area'
import { useSemanticSearch, type ArticleSearchResult } from '@/lib/api/semantic-search-api'
import { cn } from '@/lib/utils'

interface SemanticSearchDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  onSelectArticle?: (articleId: string) => void
}

export function SemanticSearchDialog({
  open,
  onOpenChange,
  onSelectArticle,
}: SemanticSearchDialogProps) {
  const [query, setQuery] = useState('')
  const [searchQuery, setSearchQuery] = useState('') // Debounced query

  // Debounce search query (500ms)
  useEffect(() => {
    const timer = setTimeout(() => {
      setSearchQuery(query)
    }, 500)
    return () => clearTimeout(timer)
  }, [query])

  const { data, isLoading, error } = useSemanticSearch(searchQuery, {
    enabled: searchQuery.length >= 3,
  })

  const handleSelectArticle = (articleId: string) => {
    onSelectArticle?.(articleId)
    onOpenChange(false)
    setQuery('')
    setSearchQuery('')
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[80vh]">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2">
            <Sparkles className="h-5 w-5 text-purple-500" />
            Busca Semântica (RAG)
          </DialogTitle>
          <DialogDescription>
            Busque artigos por significado, não apenas palavras exatas. Use linguagem natural.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* Search Input */}
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Descreva o que procura (ex: 'tratamento de hipertensão em idosos')"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              className="pl-9"
              autoFocus
            />
          </div>

          {/* Hints */}
          {query.length > 0 && query.length < 3 && (
            <p className="text-sm text-muted-foreground">
              Digite pelo menos 3 caracteres para buscar
            </p>
          )}

          {/* Loading State */}
          {isLoading && (
            <div className="flex items-center justify-center py-8">
              <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
              <span className="ml-2 text-sm text-muted-foreground">
                Buscando artigos relevantes...
              </span>
            </div>
          )}

          {/* Error State */}
          {error && (
            <div className="rounded-md border border-destructive/50 bg-destructive/10 p-4">
              <p className="text-sm text-destructive">
                Erro ao buscar artigos. Tente novamente.
              </p>
            </div>
          )}

          {/* Results */}
          {data && !isLoading && (
            <div className="space-y-2">
              <div className="flex items-center justify-between">
                <p className="text-sm text-muted-foreground">
                  {data.count} {data.count === 1 ? 'resultado encontrado' : 'resultados encontrados'}
                </p>
              </div>

              {data.count === 0 ? (
                <div className="rounded-md border border-dashed p-8 text-center">
                  <FileText className="mx-auto h-8 w-8 text-muted-foreground" />
                  <p className="mt-2 text-sm text-muted-foreground">
                    Nenhum artigo encontrado para "{data.query}"
                  </p>
                  <p className="mt-1 text-xs text-muted-foreground">
                    Tente usar termos diferentes ou mais genéricos
                  </p>
                </div>
              ) : (
                <ScrollArea className="h-[400px] pr-4">
                  <div className="space-y-3">
                    {data.results.map((result: ArticleSearchResult, index: number) => (
                      <ArticleResultCard
                        key={result.article.id}
                        result={result}
                        rank={index + 1}
                        onSelect={() => handleSelectArticle(result.article.id)}
                      />
                    ))}
                  </div>
                </ScrollArea>
              )}
            </div>
          )}

          {/* Empty State */}
          {!query && !isLoading && (
            <div className="rounded-md border border-dashed p-8 text-center">
              <Sparkles className="mx-auto h-8 w-8 text-muted-foreground" />
              <p className="mt-2 text-sm font-medium">Busca Inteligente</p>
              <p className="mt-1 text-xs text-muted-foreground">
                Use linguagem natural. Exemplo: "tratamento de diabetes tipo 2"
              </p>
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  )
}

// Article Result Card Component
interface ArticleResultCardProps {
  result: ArticleSearchResult
  rank: number
  onSelect: () => void
}

function ArticleResultCard({ result, rank, onSelect }: ArticleResultCardProps) {
  const { article, similarity, chunkText } = result
  const similarityPercent = Math.round(similarity * 100)

  return (
    <button
      onClick={onSelect}
      className={cn(
        'w-full rounded-lg border bg-card p-4 text-left transition-all hover:bg-accent hover:shadow-md',
        'focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2'
      )}
    >
      <div className="space-y-2">
        {/* Header */}
        <div className="flex items-start justify-between gap-2">
          <div className="flex-1 space-y-1">
            <div className="flex items-center gap-2">
              <Badge variant="secondary" className="text-xs">
                #{rank}
              </Badge>
              <h4 className="font-medium line-clamp-2">{article.title}</h4>
            </div>
            {article.authors && (
              <p className="text-xs text-muted-foreground line-clamp-1">
                {article.authors}
              </p>
            )}
          </div>

          {/* Similarity Score */}
          <div className="flex flex-col items-end gap-1">
            <div className="flex items-center gap-1 text-xs font-medium text-purple-600">
              <TrendingUp className="h-3 w-3" />
              {similarityPercent}%
            </div>
            <span className="text-xs text-muted-foreground">relevância</span>
          </div>
        </div>

        {/* Matching Chunk */}
        {chunkText && (
          <div className="rounded-md bg-muted/50 p-2">
            <p className="text-xs text-muted-foreground line-clamp-3">
              "{chunkText}"
            </p>
          </div>
        )}

        {/* Metadata */}
        <div className="flex items-center gap-2 text-xs text-muted-foreground">
          {article.year && <span>{article.year}</span>}
          {article.journal && (
            <>
              <span>•</span>
              <span className="line-clamp-1">{article.journal}</span>
            </>
          )}
        </div>
      </div>
    </button>
  )
}
