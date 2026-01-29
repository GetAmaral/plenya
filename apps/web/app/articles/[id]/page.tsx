'use client'

import { use, useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import dynamic from 'next/dynamic'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Separator } from '@/components/ui/separator'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {
  useArticle,
  useToggleFavorite,
  useSetRating,
  useDeleteArticle,
  ARTICLE_TYPES,
  articleApi,
} from '@/lib/api/article-api'
import { ArticleScoreItems } from '@/components/articles/ArticleScoreItems'

// Dynamically import PDFViewer with no SSR to avoid DOMMatrix errors
const PDFViewer = dynamic(
  () => import('@/components/articles/PDFViewer').then((mod) => mod.PDFViewer),
  {
    ssr: false,
    loading: () => (
      <Card className="w-full">
        <CardContent className="py-12">
          <div className="flex items-center justify-center">
            <Skeleton className="h-[600px] w-full" />
          </div>
        </CardContent>
      </Card>
    ),
  }
)
import {
  ArrowLeft,
  Heart,
  Star,
  Download,
  Edit,
  Trash2,
  ExternalLink,
  Calendar,
  BookOpen,
  FileText,
  Hash,
  Tag,
  ChevronDown,
  ChevronUp,
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
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'

interface PageProps {
  params: Promise<{ id: string }>
}

export default function ArticleDetailPage({ params }: PageProps) {
  const { id } = use(params)
  const router = useRouter()
  const { data: article, isLoading, isError } = useArticle(id)
  const toggleFavorite = useToggleFavorite()
  const setRating = useSetRating()
  const deleteArticle = useDeleteArticle()
  const [isFullContentExpanded, setIsFullContentExpanded] = useState(false)

  const handleToggleFavorite = async () => {
    if (!article) return

    try {
      await toggleFavorite.mutateAsync(article.id)
      toast.success(
        article.favorite ? 'Removido dos favoritos' : 'Adicionado aos favoritos'
      )
    } catch (error: any) {
      toast.error('Erro ao atualizar favorito')
    }
  }

  const handleSetRating = async (rating: number) => {
    if (!article) return

    try {
      await setRating.mutateAsync({ id: article.id, rating })
      toast.success(`Avaliação: ${rating} estrela${rating !== 1 ? 's' : ''}`)
    } catch (error: any) {
      toast.error('Erro ao avaliar artigo')
    }
  }

  const handleDelete = async () => {
    if (!article) return

    try {
      await deleteArticle.mutateAsync(article.id)
      toast.success('Artigo deletado com sucesso')
      router.push('/articles')
    } catch (error: any) {
      toast.error('Erro ao deletar artigo')
    }
  }

  const handleDownload = () => {
    if (!article) return
    const url = articleApi.getDownloadUrl(article.id)
    window.open(url, '_blank')
  }

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 px-4 max-w-4xl">
        <Skeleton className="h-8 w-32 mb-8" />
        <Skeleton className="h-12 w-3/4 mb-4" />
        <Skeleton className="h-6 w-1/2 mb-8" />
        <Skeleton className="h-64" />
      </div>
    )
  }

  if (isError || !article) {
    return (
      <div className="container mx-auto py-8 px-4 max-w-4xl">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4">Artigo não encontrado</h1>
          <Button asChild>
            <Link href="/articles">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Voltar para artigos
            </Link>
          </Button>
        </div>
      </div>
    )
  }

  const articleType = ARTICLE_TYPES.find((t) => t.value === article.articleType)
  const publishDate = new Date(article.publishDate)

  return (
    <div className="container mx-auto py-8 px-4 max-w-4xl">
      {/* Back Button */}
      <Button variant="ghost" asChild className="mb-6">
        <Link href="/articles">
          <ArrowLeft className="mr-2 h-4 w-4" />
          Voltar para artigos
        </Link>
      </Button>

      {/* Header */}
      <div className="mb-8">
        <div className="flex flex-wrap items-center gap-2 mb-4">
          <Badge variant="outline">{articleType?.label || article.articleType}</Badge>
          {article.specialty && <Badge variant="secondary">{article.specialty}</Badge>}
          {article.doi && (
            <Badge variant="outline" className="font-mono text-xs">
              DOI: {article.doi}
            </Badge>
          )}
          {article.pmid && (
            <Badge variant="outline" className="font-mono text-xs">
              PMID: {article.pmid}
            </Badge>
          )}
        </div>

        <h1 className="text-3xl font-bold mb-4">{article.title}</h1>

        <div className="flex flex-wrap items-center gap-4 text-muted-foreground mb-6">
          <span className="flex items-center gap-1.5">
            <BookOpen className="h-4 w-4" />
            {article.journal}
          </span>
          <span className="flex items-center gap-1.5">
            <Calendar className="h-4 w-4" />
            {format(publishDate, "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
          </span>
          {article.fileSize && (
            <span className="flex items-center gap-1.5">
              <FileText className="h-4 w-4" />
              {(article.fileSize / (1024 * 1024)).toFixed(1)} MB
            </span>
          )}
        </div>

        {/* Actions */}
        <div className="flex flex-wrap items-center gap-3">
          <Button
            variant={article.favorite ? 'default' : 'outline'}
            onClick={handleToggleFavorite}
            disabled={toggleFavorite.isPending}
          >
            <Heart
              className={`mr-2 h-4 w-4 ${article.favorite ? 'fill-current' : ''}`}
            />
            {article.favorite ? 'Favorito' : 'Adicionar aos favoritos'}
          </Button>

          {article.internalLink && (
            <Button variant="outline" onClick={handleDownload}>
              <Download className="mr-2 h-4 w-4" />
              Download PDF
            </Button>
          )}

          {article.originalLink && (
            <Button variant="outline" asChild>
              <a
                href={article.originalLink}
                target="_blank"
                rel="noopener noreferrer"
              >
                <ExternalLink className="mr-2 h-4 w-4" />
                Ver publicação original
              </a>
            </Button>
          )}

          <Button variant="outline" asChild>
            <Link href={`/articles/${article.id}/edit`}>
              <Edit className="mr-2 h-4 w-4" />
              Editar
            </Link>
          </Button>

          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="outline" className="text-destructive">
                <Trash2 className="mr-2 h-4 w-4" />
                Deletar
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
                <AlertDialogDescription>
                  Tem certeza que deseja deletar este artigo? Esta ação não pode ser
                  desfeita e o arquivo PDF será removido permanentemente.
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
        </div>

        {/* Rating */}
        <div className="mt-6">
          <p className="text-sm text-muted-foreground mb-2">Sua avaliação:</p>
          <div className="flex items-center gap-1">
            {[1, 2, 3, 4, 5].map((star) => (
              <button
                key={star}
                onClick={() => handleSetRating(star)}
                disabled={setRating.isPending}
                className="hover:scale-110 transition-transform disabled:opacity-50"
              >
                <Star
                  className={`h-6 w-6 ${
                    article.rating && star <= article.rating
                      ? 'fill-yellow-400 text-yellow-400'
                      : 'text-muted-foreground'
                  }`}
                />
              </button>
            ))}
          </div>
        </div>
      </div>

      <Separator className="my-8" />

      {/* PDF Viewer */}
      {article.internalLink && (
        <>
          <PDFViewer
            url={articleApi.getDownloadUrl(article.id)}
            title={article.title}
          />
          <Separator className="my-8" />
        </>
      )}

      {/* Content */}
      <div className="space-y-8">
        {/* Authors */}
        <Card>
          <CardHeader>
            <CardTitle className="text-lg">Autores</CardTitle>
          </CardHeader>
          <CardContent>
            <p>{article.authors}</p>
          </CardContent>
        </Card>

        {/* Abstract */}
        {article.abstract && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">Resumo</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="whitespace-pre-wrap">{article.abstract}</p>
            </CardContent>
          </Card>
        )}

        {/* Keywords & MeSH Terms */}
        {((article.keywords && article.keywords.length > 0) ||
          (article.meshTerms && article.meshTerms.length > 0)) && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Tag className="h-5 w-5" />
                Palavras-chave e Termos
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              {article.keywords && article.keywords.length > 0 && (
                <div>
                  <p className="text-sm font-medium mb-2">Keywords:</p>
                  <div className="flex flex-wrap gap-2">
                    {article.keywords.map((keyword, idx) => (
                      <Badge key={idx} variant="secondary">
                        {keyword}
                      </Badge>
                    ))}
                  </div>
                </div>
              )}

              {article.meshTerms && article.meshTerms.length > 0 && (
                <div>
                  <p className="text-sm font-medium mb-2">MeSH Terms:</p>
                  <div className="flex flex-wrap gap-2">
                    {article.meshTerms.map((term, idx) => (
                      <Badge key={idx} variant="outline">
                        {term}
                      </Badge>
                    ))}
                  </div>
                </div>
              )}
            </CardContent>
          </Card>
        )}

        {/* Score Items Vinculados */}
        {article.scoreItems && article.scoreItems.length > 0 && (
          <ArticleScoreItems scoreItems={article.scoreItems} />
        )}

        {/* Notes */}
        {article.notes && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">Notas Pessoais</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="whitespace-pre-wrap">{article.notes}</p>
            </CardContent>
          </Card>
        )}

        {/* Full Content (if available) */}
        {article.fullContent && (
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle className="text-lg">Texto Completo Extraído do PDF</CardTitle>
                  <CardDescription>
                    {article.fullContent.length.toLocaleString()} caracteres extraídos
                  </CardDescription>
                </div>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setIsFullContentExpanded(!isFullContentExpanded)}
                >
                  {isFullContentExpanded ? (
                    <>
                      <ChevronUp className="mr-2 h-4 w-4" />
                      Colapsar
                    </>
                  ) : (
                    <>
                      <ChevronDown className="mr-2 h-4 w-4" />
                      Expandir Texto Completo
                    </>
                  )}
                </Button>
              </div>
            </CardHeader>
            <CardContent>
              {isFullContentExpanded ? (
                <div className="max-h-[600px] overflow-y-auto border rounded-md p-4 bg-muted/30">
                  <p className="whitespace-pre-wrap text-sm font-mono">
                    {article.fullContent}
                  </p>
                </div>
              ) : (
                <div className="border rounded-md p-4 bg-muted/30">
                  <p className="whitespace-pre-wrap text-sm font-mono text-muted-foreground">
                    {article.fullContent.substring(0, 500)}...
                  </p>
                  <p className="text-xs text-muted-foreground mt-2">
                    Mostrando primeiros 500 caracteres. Clique em "Expandir" para ver o texto completo.
                  </p>
                </div>
              )}
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  )
}
