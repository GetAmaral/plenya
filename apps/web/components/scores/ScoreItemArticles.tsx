'use client'

import Link from 'next/link'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { FileText, ChevronRight, Star } from 'lucide-react'
import { Article, ARTICLE_TYPES } from '@/lib/api/article-api'

interface ScoreItemArticlesProps {
  articles: Article[]
  isLoading?: boolean
}

export function ScoreItemArticles({ articles, isLoading }: ScoreItemArticlesProps) {
  if (isLoading) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="text-lg flex items-center gap-2">
            <FileText className="h-5 w-5" />
            Artigos Científicos
          </CardTitle>
          <CardDescription>Carregando artigos...</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-20 bg-muted animate-pulse rounded-md" />
            ))}
          </div>
        </CardContent>
      </Card>
    )
  }

  if (!articles || articles.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="text-lg flex items-center gap-2">
            <FileText className="h-5 w-5" />
            Artigos Científicos
          </CardTitle>
          <CardDescription>
            Nenhum artigo vinculado a este score item
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex flex-col items-center justify-center py-8 text-center">
            <FileText className="h-12 w-12 text-muted-foreground/50 mb-3" />
            <p className="text-sm text-muted-foreground">
              Este score item ainda não possui artigos científicos vinculados.
            </p>
            <p className="text-xs text-muted-foreground mt-1">
              Vincule artigos para fundamentar este item.
            </p>
          </div>
        </CardContent>
      </Card>
    )
  }

  // Ordenar por data de publicação (mais recente primeiro)
  const sortedArticles = [...articles].sort((a, b) => {
    const dateA = new Date(a.publishDate).getTime()
    const dateB = new Date(b.publishDate).getTime()
    return dateB - dateA
  })

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <div>
            <CardTitle className="text-lg flex items-center gap-2">
              <FileText className="h-5 w-5" />
              Artigos Científicos
            </CardTitle>
            <CardDescription>
              {articles.length} artigo{articles.length !== 1 ? 's' : ''} vinculado{articles.length !== 1 ? 's' : ''} a este score item
            </CardDescription>
          </div>
          <Link
            href="/articles"
            className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-1"
          >
            Ver todos os artigos
            <ChevronRight className="h-4 w-4" />
          </Link>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-3">
          {sortedArticles.map((article) => {
            const publishYear = new Date(article.publishDate).getFullYear()
            const articleTypeLabel = ARTICLE_TYPES.find(
              (t) => t.value === article.articleType
            )?.label || article.articleType

            return (
              <Link
                key={article.id}
                href={`/articles/${article.id}`}
                className="block"
              >
                <div className="p-4 rounded-lg border border-border hover:border-primary/50 hover:bg-muted/50 transition-all">
                  <div className="space-y-2">
                    {/* Título e badges */}
                    <div className="flex items-start justify-between gap-3">
                      <h3 className="font-medium text-sm line-clamp-2 flex-1">
                        {article.title}
                      </h3>
                      <div className="flex items-center gap-1.5 shrink-0">
                        {article.favorite && (
                          <Star className="h-3.5 w-3.5 fill-yellow-400 text-yellow-400" />
                        )}
                        <Badge variant="outline" className="text-xs">
                          {articleTypeLabel}
                        </Badge>
                      </div>
                    </div>

                    {/* Autores */}
                    <p className="text-xs text-muted-foreground line-clamp-1">
                      {article.authors}
                    </p>

                    {/* Journal e ano */}
                    <div className="flex items-center justify-between gap-2">
                      <p className="text-xs text-muted-foreground line-clamp-1 flex-1">
                        {article.journal} • {publishYear}
                      </p>

                      {/* Rating */}
                      {article.rating !== undefined && article.rating > 0 && (
                        <div className="flex items-center gap-1 shrink-0">
                          {Array.from({ length: article.rating }).map((_, i) => (
                            <Star
                              key={i}
                              className="h-3 w-3 fill-yellow-400 text-yellow-400"
                            />
                          ))}
                        </div>
                      )}
                    </div>

                    {/* DOI/PMID */}
                    {(article.doi || article.pmid) && (
                      <div className="flex items-center gap-2">
                        {article.doi && (
                          <Badge variant="secondary" className="text-xs font-mono">
                            DOI: {article.doi}
                          </Badge>
                        )}
                        {article.pmid && (
                          <Badge variant="secondary" className="text-xs font-mono">
                            PMID: {article.pmid}
                          </Badge>
                        )}
                      </div>
                    )}
                  </div>
                </div>
              </Link>
            )
          })}
        </div>

        {/* Link para biblioteca completa */}
        <div className="mt-6 pt-4 border-t">
          <Link
            href="/articles"
            className="inline-flex items-center gap-2 text-sm text-primary hover:underline"
          >
            <FileText className="h-4 w-4" />
            Ver biblioteca completa de artigos
          </Link>
        </div>
      </CardContent>
    </Card>
  )
}
