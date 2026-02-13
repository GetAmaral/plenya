'use client'

import { Sparkles, ExternalLink, Plus, TrendingUp, Loader2 } from 'lucide-react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import { useRecommendedArticles, type ArticleSearchResult } from '@/lib/api/semantic-search-api'
import { cn } from '@/lib/utils'
import Link from 'next/link'

interface ArticleRecommendationsProps {
  scoreItemId: string | undefined
  scoreItemName?: string
  onLinkArticle?: (articleId: string) => void
  limit?: number
  className?: string
}

/**
 * Widget que mostra recomendações automáticas de artigos para um ScoreItem
 * Usa busca semântica (RAG) para encontrar artigos relevantes
 */
export function ArticleRecommendations({
  scoreItemId,
  scoreItemName,
  onLinkArticle,
  limit = 3,
  className,
}: ArticleRecommendationsProps) {
  const { data, isLoading, error } = useRecommendedArticles(scoreItemId, { limit })

  if (!scoreItemId) {
    return null
  }

  return (
    <Card className={cn('border-purple-200 bg-purple-50/50', className)}>
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between gap-2">
          <div className="space-y-1">
            <CardTitle className="flex items-center gap-2 text-base">
              <Sparkles className="h-4 w-4 text-purple-500" />
              Artigos Recomendados (RAG)
            </CardTitle>
            <CardDescription className="text-xs">
              Sugestões baseadas em similaridade semântica
              {scoreItemName && (
                <span className="block mt-1 font-medium text-purple-700">
                  para: {scoreItemName}
                </span>
              )}
            </CardDescription>
          </div>
          <Badge variant="secondary" className="text-xs">
            Auto
          </Badge>
        </div>
      </CardHeader>

      <CardContent>
        {/* Loading State */}
        {isLoading && (
          <div className="flex items-center justify-center py-6">
            <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            <span className="ml-2 text-sm text-muted-foreground">
              Buscando recomendações...
            </span>
          </div>
        )}

        {/* Error State */}
        {error && (
          <div className="rounded-md bg-destructive/10 p-3 text-sm text-destructive">
            Erro ao carregar recomendações
          </div>
        )}

        {/* Empty State */}
        {data && data.count === 0 && !isLoading && (
          <div className="text-center py-6">
            <p className="text-sm text-muted-foreground">
              Nenhum artigo encontrado para este parâmetro
            </p>
            <p className="text-xs text-muted-foreground mt-1">
              Tente adicionar mais informações ao campo "Relevância Clínica"
            </p>
          </div>
        )}

        {/* Results */}
        {data && data.count > 0 && !isLoading && (
          <div className="space-y-3">
            {data.results.map((result: ArticleSearchResult, index: number) => (
              <div key={result.article.id}>
                {index > 0 && <Separator className="my-3" />}
                <RecommendedArticleCard
                  result={result}
                  onLink={onLinkArticle}
                />
              </div>
            ))}

            {/* View More Link */}
            {data.count > limit && (
              <div className="pt-2 text-center">
                <p className="text-xs text-muted-foreground">
                  +{data.count - limit} artigos adicionais encontrados
                </p>
              </div>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  )
}

// Recommended Article Card Component
interface RecommendedArticleCardProps {
  result: ArticleSearchResult
  onLink?: (articleId: string) => void
}

function RecommendedArticleCard({ result, onLink }: RecommendedArticleCardProps) {
  const { article, similarity } = result
  const similarityPercent = Math.round(similarity * 100)

  return (
    <div className="space-y-2">
      {/* Header with Similarity */}
      <div className="flex items-start justify-between gap-2">
        <div className="flex-1 min-w-0">
          <h4 className="font-medium text-sm line-clamp-2 leading-tight">
            {article.title}
          </h4>
          {article.authors && (
            <p className="text-xs text-muted-foreground mt-1 line-clamp-1">
              {article.authors}
            </p>
          )}
        </div>

        {/* Similarity Badge */}
        <Badge
          variant="secondary"
          className={cn(
            'shrink-0 text-xs font-medium',
            similarityPercent >= 80 && 'bg-green-100 text-green-700',
            similarityPercent >= 70 && similarityPercent < 80 && 'bg-blue-100 text-blue-700',
            similarityPercent < 70 && 'bg-gray-100 text-gray-700'
          )}
        >
          <TrendingUp className="h-3 w-3 mr-1" />
          {similarityPercent}%
        </Badge>
      </div>

      {/* Metadata */}
      {(article.year || article.journal) && (
        <div className="flex items-center gap-2 text-xs text-muted-foreground">
          {article.year && <span>{article.year}</span>}
          {article.journal && (
            <>
              <span>•</span>
              <span className="line-clamp-1">{article.journal}</span>
            </>
          )}
        </div>
      )}

      {/* Actions */}
      <div className="flex items-center gap-2">
        {onLink && (
          <Button
            size="sm"
            variant="default"
            onClick={() => onLink(article.id)}
            className="h-7 text-xs"
          >
            <Plus className="h-3 w-3 mr-1" />
            Vincular
          </Button>
        )}
        <Button
          size="sm"
          variant="ghost"
          asChild
          className="h-7 text-xs"
        >
          <Link href={`/articles/${article.id}`} target="_blank">
            <ExternalLink className="h-3 w-3 mr-1" />
            Ver Artigo
          </Link>
        </Button>
      </div>
    </div>
  )
}
