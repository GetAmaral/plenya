'use client'

import { Sparkles, ExternalLink, TrendingUp, Loader2, Activity } from 'lucide-react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import { useRelatedScoreItems, type ScoreItemSimilarity } from '@/lib/api/semantic-search-api'
import { cn } from '@/lib/utils'
import Link from 'next/link'

interface RelatedScoreItemsProps {
  articleId: string | undefined
  articleTitle?: string
  limit?: number
  className?: string
}

/**
 * Componente que mostra quais ScoreItems (parâmetros clínicos) um artigo cobre
 * Usa busca semântica (RAG) para descobrir relações automáticas
 */
export function RelatedScoreItems({
  articleId,
  articleTitle,
  limit = 10,
  className,
}: RelatedScoreItemsProps) {
  const { data, isLoading, error } = useRelatedScoreItems(articleId, { limit })

  if (!articleId) {
    return null
  }

  return (
    <Card className={cn('border-purple-200 bg-purple-50/50', className)}>
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between gap-2">
          <div className="space-y-1">
            <CardTitle className="flex items-center gap-2 text-base">
              <Sparkles className="h-4 w-4 text-purple-500" />
              Parâmetros Clínicos Relacionados
            </CardTitle>
            <CardDescription className="text-xs">
              ScoreItems cobertos por este artigo (descoberta automática via RAG)
              {articleTitle && (
                <span className="block mt-1 font-medium text-purple-700 line-clamp-1">
                  {articleTitle}
                </span>
              )}
            </CardDescription>
          </div>
          <Badge variant="secondary" className="text-xs shrink-0">
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
              Analisando artigo...
            </span>
          </div>
        )}

        {/* Error State */}
        {error && (
          <div className="rounded-md bg-destructive/10 p-3 text-sm text-destructive">
            Erro ao analisar artigo
          </div>
        )}

        {/* Empty State */}
        {data && data.count === 0 && !isLoading && (
          <div className="text-center py-6">
            <Activity className="mx-auto h-8 w-8 text-muted-foreground" />
            <p className="mt-2 text-sm text-muted-foreground">
              Nenhum parâmetro clínico relacionado encontrado
            </p>
            <p className="text-xs text-muted-foreground mt-1">
              Este artigo pode não cobrir parâmetros específicos do sistema de escores
            </p>
          </div>
        )}

        {/* Results */}
        {data && data.count > 0 && !isLoading && (
          <div className="space-y-3">
            <div className="text-sm text-muted-foreground">
              {data.count} {data.count === 1 ? 'parâmetro encontrado' : 'parâmetros encontrados'}
            </div>

            <div className="space-y-2">
              {data.results.map((result: ScoreItemSimilarity, index: number) => (
                <div key={result.scoreItem.id}>
                  {index > 0 && <Separator className="my-2" />}
                  <RelatedScoreItemCard result={result} />
                </div>
              ))}
            </div>

            {/* Info Footer */}
            <div className="pt-2 rounded-md bg-purple-100/50 p-3">
              <p className="text-xs text-purple-700">
                💡 <strong>Dica:</strong> Estes parâmetros podem se beneficiar deste artigo como referência científica.
                Considere vincular os mais relevantes.
              </p>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  )
}

// Related ScoreItem Card Component
interface RelatedScoreItemCardProps {
  result: ScoreItemSimilarity
}

function RelatedScoreItemCard({ result }: RelatedScoreItemCardProps) {
  const { scoreItem, similarity } = result
  const similarityPercent = Math.round(similarity * 100)

  // Build breadcrumb path (Group > Subgroup > Item)
  const breadcrumb = [
    scoreItem.subgroup?.group?.name,
    scoreItem.subgroup?.name,
    scoreItem.name,
  ]
    .filter(Boolean)
    .join(' › ')

  return (
    <div className="space-y-2 rounded-md border bg-card p-3 hover:bg-accent/50 transition-colors">
      {/* Header with Similarity */}
      <div className="flex items-start justify-between gap-2">
        <div className="flex-1 min-w-0">
          <h4 className="font-medium text-sm leading-tight">
            {scoreItem.name}
          </h4>
          {breadcrumb && (
            <p className="text-xs text-muted-foreground mt-1 line-clamp-1">
              {breadcrumb}
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
      <div className="flex items-center gap-2 text-xs text-muted-foreground">
        {scoreItem.unit && (
          <>
            <span>Unidade: {scoreItem.unit}</span>
            <span>•</span>
          </>
        )}
        {scoreItem.gender && scoreItem.gender !== 'not_applicable' && (
          <>
            <span>
              {scoreItem.gender === 'male' ? 'Homens' : 'Mulheres'}
            </span>
            <span>•</span>
          </>
        )}
        {scoreItem.ageRangeMin !== null || scoreItem.ageRangeMax !== null ? (
          <span>
            Idade: {scoreItem.ageRangeMin ?? 0}-{scoreItem.ageRangeMax ?? '∞'}
          </span>
        ) : null}
      </div>

      {/* Clinical Relevance Snippet */}
      {scoreItem.clinicalRelevance && (
        <div className="rounded-md bg-muted/50 p-2">
          <p className="text-xs text-muted-foreground line-clamp-2">
            {scoreItem.clinicalRelevance}
          </p>
        </div>
      )}

      {/* Action */}
      <div className="flex items-center gap-2 pt-1">
        <Button
          size="sm"
          variant="ghost"
          asChild
          className="h-7 text-xs"
        >
          <Link
            href={`/scores/${scoreItem.subgroup?.groupId}/${scoreItem.subgroupId}/${scoreItem.id}`}
            target="_blank"
          >
            <ExternalLink className="h-3 w-3 mr-1" />
            Ver Parâmetro
          </Link>
        </Button>
      </div>
    </div>
  )
}
