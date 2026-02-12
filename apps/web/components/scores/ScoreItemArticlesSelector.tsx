'use client'

import { useState, useMemo } from 'react'
import { Label } from '@/components/ui/label'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { Input } from '@/components/ui/input'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Search, Star } from 'lucide-react'
import { useArticles, ARTICLE_TYPES } from '@/lib/api/article-api'

interface ScoreItemArticlesSelectorProps {
  selectedArticleIds: string[]
  onChange: (articleIds: string[]) => void
}

export function ScoreItemArticlesSelector({
  selectedArticleIds,
  onChange,
}: ScoreItemArticlesSelectorProps) {
  const { data, isLoading } = useArticles(1, 100) // Carregar primeiros 100 artigos
  const articles = data?.articles || []
  const [searchQuery, setSearchQuery] = useState('')

  // Filtrar artigos por busca
  const filteredArticles = useMemo(() => {
    if (!searchQuery) return articles

    const query = searchQuery.toLowerCase()
    return articles.filter(
      (article) =>
        article.title.toLowerCase().includes(query) ||
        article.authors.toLowerCase().includes(query) ||
        article.journal.toLowerCase().includes(query) ||
        (article.doi && article.doi.toLowerCase().includes(query)) ||
        (article.pmid && article.pmid.toLowerCase().includes(query))
    )
  }, [articles, searchQuery])

  // Ordenar artigos alfabeticamente por título
  const sortedArticles = useMemo(() => {
    return [...filteredArticles].sort((a, b) => a.title.localeCompare(b.title))
  }, [filteredArticles])

  const toggleArticle = (articleId: string) => {
    if (selectedArticleIds.includes(articleId)) {
      onChange(selectedArticleIds.filter((id) => id !== articleId))
    } else {
      onChange([...selectedArticleIds, articleId])
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
        <Label>Artigos Científicos Vinculados</Label>
        <Badge variant="secondary">
          {selectedArticleIds.length} de {articles.length} selecionados
        </Badge>
      </div>

      {/* Search */}
      <div className="relative">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input
          placeholder="Buscar por título, autores, journal, DOI ou PMID..."
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="pl-9"
        />
      </div>

      {/* Articles List */}
      <ScrollArea className="h-[400px] rounded-md border">
        <div className="p-4 space-y-2">
          {sortedArticles.length === 0 ? (
            <p className="text-sm text-muted-foreground text-center py-8">
              Nenhum artigo encontrado
            </p>
          ) : (
            sortedArticles.map((article) => {
              const publishYear = new Date(article.publishDate).getFullYear()
              const articleTypeLabel = ARTICLE_TYPES.find(
                (t) => t.value === article.articleType
              )?.label || article.articleType

              return (
                <div
                  key={article.id}
                  className="flex items-start gap-3 hover:bg-muted/30 rounded-md p-3 transition-colors border border-transparent hover:border-muted"
                >
                  <Checkbox
                    checked={selectedArticleIds.includes(article.id)}
                    onCheckedChange={() => toggleArticle(article.id)}
                    className="mt-0.5"
                  />
                  <div className="flex-1 min-w-0 space-y-1">
                    {/* Título e badges */}
                    <div className="flex items-start justify-between gap-2">
                      <p className="text-sm font-medium line-clamp-2 flex-1">
                        {article.title}
                      </p>
                      <div className="flex items-center gap-1.5 shrink-0">
                        {article.favorite && (
                          <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
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
                        <div className="flex items-center gap-0.5 shrink-0">
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
                      <div className="flex items-center gap-2 flex-wrap">
                        {article.doi && (
                          <Badge variant="secondary" className="text-xs font-mono">
                            DOI: {article.doi.length > 20 ? article.doi.slice(0, 20) + '...' : article.doi}
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
              )
            })
          )}
        </div>
      </ScrollArea>

      <p className="text-xs text-muted-foreground">
        Selecione os artigos científicos que fundamentam este score item
      </p>
    </div>
  )
}
