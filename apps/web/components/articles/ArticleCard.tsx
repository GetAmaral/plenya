'use client'

import { useState } from 'react'
import Link from 'next/link'
import { formatDistanceToNow } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  Article,
  ARTICLE_TYPES,
  useToggleFavorite,
  useSetRating,
  useDeleteArticle,
  useBookChapters,
  articleApi,
} from '@/lib/api/article-api'
import {
  Star,
  Heart,
  Download,
  ExternalLink,
  MoreVertical,
  Trash2,
  Edit,
  FileText,
  Calendar,
  BookOpen,
  ChevronDown,
  ChevronUp,
  List,
} from 'lucide-react'
import { toast } from 'sonner'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'

interface ArticleCardProps {
  article: Article
}

// ─── sub-componente: lista de capítulos ───────────────────────────────────────

function ChapterList({ bookId }: { bookId: string }) {
  const [showAll, setShowAll] = useState(false)
  const { data, isLoading, isError } = useBookChapters(bookId)

  if (isLoading) {
    return (
      <div className="space-y-1.5 mt-1">
        {[1, 2, 3].map((i) => (
          <Skeleton key={i} className="h-5 w-full" />
        ))}
      </div>
    )
  }

  if (isError || !data?.chapters?.length) {
    return (
      <p className="text-xs text-muted-foreground mt-1 italic">
        Nenhum capítulo encontrado.
      </p>
    )
  }

  const chapters = data.chapters
  const visible = showAll ? chapters : chapters.slice(0, 5)

  return (
    <div className="mt-1 space-y-0.5">
      {visible.map((ch) => (
        <Link
          key={ch.id}
          href={`/articles/${ch.id}`}
          className="flex items-start gap-2 rounded px-1 py-0.5 text-xs text-muted-foreground hover:bg-muted hover:text-foreground transition-colors group/ch"
        >
          <span className="flex-shrink-0 tabular-nums text-muted-foreground/60 group-hover/ch:text-muted-foreground w-5 text-right">
            {ch.chapterNumber}.
          </span>
          <span className="line-clamp-1 flex-1">
            {ch.chapterTitle || ch.title}
          </span>
        </Link>
      ))}

      {chapters.length > 5 && (
        <button
          onClick={(e) => { e.preventDefault(); setShowAll(!showAll) }}
          className="mt-1 text-xs text-primary hover:underline pl-7"
        >
          {showAll
            ? 'Ver menos'
            : `+ ${chapters.length - 5} capítulo${chapters.length - 5 !== 1 ? 's' : ''}`}
        </button>
      )}
    </div>
  )
}

// ─── card principal ───────────────────────────────────────────────────────────

export function ArticleCard({ article }: ArticleCardProps) {
  const [showDeleteDialog, setShowDeleteDialog] = useState(false)
  const [showChapters, setShowChapters] = useState(false)

  const toggleFavorite = useToggleFavorite()
  const setRating = useSetRating()
  const deleteArticle = useDeleteArticle()

  const isBook = article.sourceType === 'book'
  const articleType = ARTICLE_TYPES.find((t) => t.value === article.articleType)

  const handleToggleFavorite = async () => {
    try {
      await toggleFavorite.mutateAsync(article.id)
      toast.success(article.favorite ? 'Removido dos favoritos' : 'Adicionado aos favoritos')
    } catch {
      toast.error('Erro ao atualizar favorito')
    }
  }

  const handleSetRating = async (rating: number) => {
    try {
      await setRating.mutateAsync({ id: article.id, rating })
      toast.success(`Avaliação: ${rating} estrela${rating !== 1 ? 's' : ''}`)
    } catch {
      toast.error('Erro ao avaliar')
    }
  }

  const handleDelete = async () => {
    try {
      await deleteArticle.mutateAsync(article.id)
      toast.success(isBook ? 'Livro deletado com sucesso' : 'Artigo deletado com sucesso')
      setShowDeleteDialog(false)
    } catch {
      toast.error('Erro ao deletar')
    }
  }

  const handleDownload = () => {
    const url = articleApi.getDownloadUrl(article.id)
    window.open(url, '_blank')
  }

  const publishDate = new Date(article.publishDate)
  const formattedDate = formatDistanceToNow(publishDate, {
    addSuffix: true,
    locale: ptBR,
  })

  return (
    <>
      <Card
        className={`group hover:shadow-lg transition-all duration-200 ${
          isBook
            ? 'border-primary/20 bg-primary/[0.02] hover:border-primary/40'
            : ''
        }`}
      >
        <CardHeader className="pb-3">
          <div className="flex items-start justify-between gap-3">
            <div className="flex-1 min-w-0">
              <Link
                href={`/articles/${article.id}`}
                className="block group-hover:text-primary transition-colors"
              >
                {/* Ícone de livro antes do título */}
                {isBook && (
                  <div className="flex items-center gap-1.5 mb-1">
                    <BookOpen className="h-4 w-4 text-primary flex-shrink-0" />
                    <span className="text-xs font-medium text-primary uppercase tracking-wide">
                      Livro
                    </span>
                  </div>
                )}
                <h3 className="font-semibold text-base line-clamp-2 leading-tight">
                  {article.title}
                </h3>
              </Link>

              <div className="flex flex-wrap items-center gap-2 mt-2">
                {isBook ? (
                  <>
                    <Badge
                      variant="outline"
                      className="font-normal border-primary/30 text-primary"
                    >
                      <BookOpen className="h-3 w-3 mr-1" />
                      Livro
                    </Badge>
                    {article.totalChapters && (
                      <Badge variant="secondary" className="font-normal">
                        <List className="h-3 w-3 mr-1" />
                        {article.totalChapters} capítulo{article.totalChapters !== 1 ? 's' : ''}
                      </Badge>
                    )}
                  </>
                ) : (
                  <>
                    <Badge variant="outline" className="font-normal">
                      {articleType?.label || article.articleType}
                    </Badge>
                    {article.specialty && (
                      <Badge variant="secondary" className="font-normal">
                        {article.specialty}
                      </Badge>
                    )}
                    {article.doi && (
                      <Badge variant="outline" className="font-mono text-xs">
                        DOI
                      </Badge>
                    )}
                    {article.pmid && (
                      <Badge variant="outline" className="font-mono text-xs">
                        PMID
                      </Badge>
                    )}
                  </>
                )}
              </div>
            </div>

            <div className="flex items-center gap-1 flex-shrink-0">
              <Button
                variant="ghost"
                size="icon"
                onClick={handleToggleFavorite}
                disabled={toggleFavorite.isPending}
                className={article.favorite ? 'text-red-500' : 'text-muted-foreground'}
              >
                <Heart className={`h-4 w-4 ${article.favorite ? 'fill-current' : ''}`} />
              </Button>

              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="ghost" size="icon">
                    <MoreVertical className="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuItem asChild>
                    <Link href={`/articles/${article.id}`}>
                      <FileText className="mr-2 h-4 w-4" />
                      Ver detalhes
                    </Link>
                  </DropdownMenuItem>

                  <DropdownMenuItem asChild>
                    <Link href={`/articles/${article.id}/edit`}>
                      <Edit className="mr-2 h-4 w-4" />
                      Editar
                    </Link>
                  </DropdownMenuItem>

                  {article.internalLink && (
                    <DropdownMenuItem onClick={handleDownload}>
                      <Download className="mr-2 h-4 w-4" />
                      {isBook ? 'Download do arquivo' : 'Download PDF'}
                    </DropdownMenuItem>
                  )}

                  {article.originalLink && (
                    <DropdownMenuItem asChild>
                      <a
                        href={article.originalLink}
                        target="_blank"
                        rel="noopener noreferrer"
                      >
                        <ExternalLink className="mr-2 h-4 w-4" />
                        Abrir publicação
                      </a>
                    </DropdownMenuItem>
                  )}

                  <DropdownMenuSeparator />

                  <DropdownMenuItem
                    onClick={() => setShowDeleteDialog(true)}
                    className="text-destructive focus:text-destructive"
                  >
                    <Trash2 className="mr-2 h-4 w-4" />
                    Deletar
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          </div>
        </CardHeader>

        <CardContent className="pb-3">
          {/* Authors & Journal/Date */}
          <div className="space-y-1 text-sm text-muted-foreground mb-3">
            {article.authors && article.authors !== 'Desconhecido' && (
              <p className="line-clamp-1">
                <span className="font-medium">Autores:</span> {article.authors}
              </p>
            )}
            <div className="flex items-center gap-4 flex-wrap">
              {!isBook && (
                <span className="flex items-center gap-1.5">
                  <BookOpen className="h-3.5 w-3.5" />
                  {article.journal}
                </span>
              )}
              <span className="flex items-center gap-1.5">
                <Calendar className="h-3.5 w-3.5" />
                {formattedDate}
              </span>
            </div>
          </div>

          {/* Abstract */}
          {article.abstract && (
            <p className="text-sm text-muted-foreground line-clamp-3">
              {article.abstract}
            </p>
          )}

          {/* Keywords (somente artigos) */}
          {!isBook && article.keywords && article.keywords.length > 0 && (
            <div className="flex flex-wrap gap-1 mt-3">
              {article.keywords.slice(0, 5).map((keyword, idx) => (
                <Badge key={idx} variant="secondary" className="text-xs font-normal">
                  {keyword}
                </Badge>
              ))}
              {article.keywords.length > 5 && (
                <Badge variant="secondary" className="text-xs font-normal">
                  +{article.keywords.length - 5}
                </Badge>
              )}
            </div>
          )}

          {/* Seção de capítulos para livros */}
          {isBook && article.totalChapters && article.totalChapters > 0 && (
            <div className="mt-3">
              <button
                onClick={() => setShowChapters(!showChapters)}
                className="flex items-center gap-1.5 text-xs font-medium text-primary hover:text-primary/80 transition-colors"
              >
                {showChapters ? (
                  <ChevronUp className="h-3.5 w-3.5" />
                ) : (
                  <ChevronDown className="h-3.5 w-3.5" />
                )}
                {showChapters ? 'Ocultar capítulos' : `Ver ${article.totalChapters} capítulo${article.totalChapters !== 1 ? 's' : ''}`}
              </button>

              {showChapters && (
                <div className="mt-2 border-l-2 border-primary/20 pl-3">
                  <ChapterList bookId={article.id} />
                </div>
              )}
            </div>
          )}
        </CardContent>

        <CardFooter className="pt-3 border-t">
          <div className="flex items-center justify-between w-full">
            {/* Rating */}
            <div className="flex items-center gap-1">
              {[1, 2, 3, 4, 5].map((star) => (
                <button
                  key={star}
                  onClick={() => handleSetRating(star)}
                  disabled={setRating.isPending}
                  className="hover:scale-110 transition-transform disabled:opacity-50"
                >
                  <Star
                    className={`h-4 w-4 ${
                      article.rating && star <= article.rating
                        ? 'fill-yellow-400 text-yellow-400'
                        : 'text-muted-foreground'
                    }`}
                  />
                </button>
              ))}
            </div>

            {/* File Size */}
            {article.fileSize && (
              <span className="text-xs text-muted-foreground">
                {(article.fileSize / (1024 * 1024)).toFixed(1)} MB
              </span>
            )}
          </div>
        </CardFooter>
      </Card>

      {/* Delete Confirmation Dialog */}
      <AlertDialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              {isBook ? (
                <>
                  Tem certeza que deseja deletar o livro{' '}
                  <strong>"{article.title}"</strong>? Todos os{' '}
                  {article.totalChapters} capítulos também serão removidos
                  permanentemente.
                </>
              ) : (
                <>
                  Tem certeza que deseja deletar o artigo{' '}
                  <strong>"{article.title}"</strong>? Esta ação não pode ser
                  desfeita e o arquivo PDF será removido permanentemente.
                </>
              )}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={handleDelete}
              className="bg-destructive hover:bg-destructive/90"
            >
              {deleteArticle.isPending ? 'Deletando...' : 'Deletar'}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
