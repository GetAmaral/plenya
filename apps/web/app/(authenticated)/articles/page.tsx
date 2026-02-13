'use client'

import { useState, useRef } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Skeleton } from '@/components/ui/skeleton'
import {
  useArticles,
  useSearchArticles,
  ArticleFilters as Filters,
} from '@/lib/api/article-api'
import { ArticleCard } from '@/components/articles/ArticleCard'
import { ArticleUploadDialog } from '@/components/articles/ArticleUploadDialog'
import { ArticleFilters } from '@/components/articles/ArticleFilters'
import { SemanticSearchDialog } from '@/components/articles/SemanticSearchDialog'
import {
  Upload,
  Search,
  X,
  ChevronLeft,
  ChevronRight,
  FileText,
  Heart,
  Sparkles,
} from 'lucide-react'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { PageHeader } from '@/components/layout/page-header'

type ViewMode = 'all' | 'favorites' | 'search'

export default function ArticlesPage() {
  const router = useRouter()
  const [viewMode, setViewMode] = useState<ViewMode>('all')
  const [searchQuery, setSearchQuery] = useState('')
  const [searchInput, setSearchInput] = useState('')
  const [page, setPage] = useState(1)
  const [pageSize, setPageSize] = useState(20)
  const [filters, setFilters] = useState<Filters>({})
  const [showUploadDialog, setShowUploadDialog] = useState(false)
  const [showSemanticSearchDialog, setShowSemanticSearchDialog] = useState(false)

  const searchFormRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef: searchFormRef })

  // Query baseado no modo
  const articlesQuery = useArticles(page, pageSize, filters)
  const favoritesQuery = useArticles(page, pageSize, { ...filters, favorite: true })
  const searchQuery_result = useSearchArticles(searchQuery, page, pageSize)

  const currentQuery =
    viewMode === 'search'
      ? searchQuery_result
      : viewMode === 'favorites'
      ? favoritesQuery
      : articlesQuery

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault()
    if (searchInput.trim()) {
      setSearchQuery(searchInput.trim())
      setViewMode('search')
      setPage(1)
    }
  }

  const handleClearSearch = () => {
    setSearchInput('')
    setSearchQuery('')
    setViewMode('all')
    setPage(1)
  }

  const handleFiltersChange = (newFilters: Filters) => {
    setFilters(newFilters)
    setPage(1)
  }

  const handlePageChange = (newPage: number) => {
    setPage(newPage)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  const totalPages = currentQuery.data?.pagination.totalPages || 0

  return (
    <div className="container mx-auto py-8 px-4 space-y-8">
      {/* Header */}
      <PageHeader
        breadcrumbs={[{ label: 'Artigos' }]}
        title="Artigos Científicos"
        description={
          articlesQuery.data
            ? `${articlesQuery.data.pagination.total} artigos na biblioteca`
            : 'Gerencie sua biblioteca de artigos médicos'
        }
        actions={[
          {
            label: 'Importar',
            icon: <Upload className="h-4 w-4" />,
            onClick: () => setShowUploadDialog(true),
            variant: 'default',
          },
        ]}
      />

      <div>

        {/* Search & Filters Bar */}
        <div className="flex flex-col lg:flex-row gap-4">
          {/* Search */}
          <form ref={searchFormRef} onSubmit={handleSearch} className="flex-1">
            <div className="relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Buscar por título, autores, abstract ou revista..."
                value={searchInput}
                onChange={(e) => setSearchInput(e.target.value)}
                className="pl-10 pr-10"
              />
              {searchInput && (
                <button
                  type="button"
                  onClick={handleClearSearch}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
                >
                  <X className="h-4 w-4" />
                </button>
              )}
            </div>
          </form>

          {/* View Mode Tabs */}
          <div className="flex gap-2 flex-wrap">
            <Button
              variant={viewMode === 'all' ? 'default' : 'outline'}
              onClick={() => {
                setViewMode('all')
                setPage(1)
              }}
              className="flex items-center gap-2"
            >
              <FileText className="h-4 w-4" />
              Todos
              {articlesQuery.data && (
                <span className="text-xs opacity-70">
                  ({articlesQuery.data.pagination.total})
                </span>
              )}
            </Button>

            <Button
              variant={viewMode === 'favorites' ? 'default' : 'outline'}
              onClick={() => {
                setViewMode('favorites')
                setPage(1)
              }}
              className="flex items-center gap-2"
            >
              <Heart className="h-4 w-4" />
              Favoritos
            </Button>

            <Button
              variant="outline"
              onClick={() => setShowSemanticSearchDialog(true)}
              className="flex items-center gap-2 border-purple-300 text-purple-700 hover:bg-purple-50 hover:text-purple-800"
            >
              <Sparkles className="h-4 w-4" />
              Busca Semântica (RAG)
            </Button>
          </div>
        </div>

        {/* Active Search Indicator */}
        {viewMode === 'search' && searchQuery && (
          <div className="mt-4 flex items-center gap-2 text-sm">
            <span className="text-muted-foreground">
              Buscando por: <strong>"{searchQuery}"</strong>
            </span>
            <Button
              variant="ghost"
              size="sm"
              onClick={handleClearSearch}
              className="h-6"
            >
              <X className="h-3 w-3 mr-1" />
              Limpar
            </Button>
          </div>
        )}
      </div>

      {/* Content */}
      <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
        {/* Filters Sidebar */}
        <div className="lg:col-span-1">
          <ArticleFilters filters={filters} onFiltersChange={handleFiltersChange} />
        </div>

        {/* Articles Grid */}
        <div className="lg:col-span-3">
          {/* Loading */}
          {currentQuery.isLoading && (
            <div className="grid gap-4 sm:grid-cols-1 md:grid-cols-2">
              {Array.from({ length: 6 }).map((_, i) => (
                <Skeleton key={i} className="h-64 rounded-lg" />
              ))}
            </div>
          )}

          {/* Error */}
          {currentQuery.isError && (
            <div className="text-center py-12">
              <p className="text-destructive">
                Erro ao carregar artigos. Tente novamente.
              </p>
            </div>
          )}

          {/* Empty State */}
          {currentQuery.isSuccess && currentQuery.data.articles.length === 0 && (
            <div className="text-center py-12">
              <FileText className="h-12 w-12 mx-auto text-muted-foreground mb-4" />
              <h3 className="text-lg font-semibold mb-2">Nenhum artigo encontrado</h3>
              <p className="text-muted-foreground mb-4">
                {viewMode === 'search'
                  ? 'Tente ajustar sua busca ou filtros.'
                  : viewMode === 'favorites'
                  ? 'Você ainda não marcou nenhum artigo como favorito.'
                  : 'Comece importando um PDF ou criando um artigo manualmente.'}
              </p>
              {viewMode === 'all' && (
                <Button onClick={() => setShowUploadDialog(true)}>
                  <Upload className="mr-2 h-4 w-4" />
                  Importar Primeiro Artigo
                </Button>
              )}
            </div>
          )}

          {/* Articles List */}
          {currentQuery.isSuccess && currentQuery.data.articles.length > 0 && (
            <>
              {/* Results Count */}
              <div className="mb-4 flex items-center justify-between">
                <p className="text-sm text-muted-foreground">
                  {currentQuery.data.pagination.total} artigo
                  {currentQuery.data.pagination.total !== 1 ? 's' : ''} encontrado
                  {currentQuery.data.pagination.total !== 1 ? 's' : ''}
                </p>

                <Select
                  value={pageSize.toString()}
                  onValueChange={(value) => {
                    setPageSize(parseInt(value))
                    setPage(1)
                  }}
                >
                  <SelectTrigger className="w-32">
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="10">10 por página</SelectItem>
                    <SelectItem value="20">20 por página</SelectItem>
                    <SelectItem value="50">50 por página</SelectItem>
                    <SelectItem value="100">100 por página</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              {/* Grid */}
              <div className="grid gap-4 sm:grid-cols-1 md:grid-cols-2">
                {currentQuery.data.articles.map((article) => (
                  <ArticleCard key={article.id} article={article} />
                ))}
              </div>

              {/* Pagination */}
              {totalPages > 1 && (
                <div className="mt-8 flex items-center justify-center gap-2">
                  <Button
                    variant="outline"
                    size="icon"
                    onClick={() => handlePageChange(page - 1)}
                    disabled={page === 1}
                  >
                    <ChevronLeft className="h-4 w-4" />
                  </Button>

                  <div className="flex items-center gap-1">
                    {Array.from({ length: Math.min(5, totalPages) }, (_, i) => {
                      let pageNum: number
                      if (totalPages <= 5) {
                        pageNum = i + 1
                      } else if (page <= 3) {
                        pageNum = i + 1
                      } else if (page >= totalPages - 2) {
                        pageNum = totalPages - 4 + i
                      } else {
                        pageNum = page - 2 + i
                      }

                      return (
                        <Button
                          key={pageNum}
                          variant={page === pageNum ? 'default' : 'outline'}
                          size="icon"
                          onClick={() => handlePageChange(pageNum)}
                        >
                          {pageNum}
                        </Button>
                      )
                    })}
                  </div>

                  <Button
                    variant="outline"
                    size="icon"
                    onClick={() => handlePageChange(page + 1)}
                    disabled={page === totalPages}
                  >
                    <ChevronRight className="h-4 w-4" />
                  </Button>
                </div>
              )}
            </>
          )}
        </div>
      </div>

      {/* Semantic Search Dialog */}
      <SemanticSearchDialog
        open={showSemanticSearchDialog}
        onOpenChange={setShowSemanticSearchDialog}
        onSelectArticle={(articleId) => router.push(`/articles/${articleId}`)}
      />

      {/* Upload Dialog */}
      <ArticleUploadDialog
        open={showUploadDialog}
        onOpenChange={setShowUploadDialog}
      />
    </div>
  )
}
