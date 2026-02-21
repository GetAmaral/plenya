'use client'

import { use, useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

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
  useBookChapters,
  ARTICLE_TYPES,
  articleApi,
} from '@/lib/api/article-api'
import { ArticleScoreItems } from '@/components/articles/ArticleScoreItems'
import { RelatedScoreItems } from '@/components/articles/RelatedScoreItems'
import { PageHeader } from '@/components/layout/page-header'
import { FileViewer } from '@/components/articles/FileViewer'

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
  Tag,
  ChevronDown,
  ChevronUp,
  Database,
  CheckCircle2,
  Clock,
  XCircle,
  Loader2,
  List,
  Hash,
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

// ─── lista completa de capítulos (usada na página de detalhe do livro) ─────────

function BookChapterList({ bookId }: { bookId: string }) {
  const { data, isLoading, isError } = useBookChapters(bookId)

  if (isLoading) {
    return (
      <div className="space-y-2">
        {Array.from({ length: 5 }).map((_, i) => (
          <Skeleton key={i} className="h-10 w-full" />
        ))}
      </div>
    )
  }

  if (isError || !data?.chapters?.length) {
    return (
      <p className="text-sm text-muted-foreground italic">
        Nenhum capítulo encontrado.
      </p>
    )
  }

  return (
    <div className="divide-y divide-border rounded-md border overflow-hidden">
      {data.chapters.map((ch) => (
        <Link
          key={ch.id}
          href={`/articles/${ch.id}`}
          className="flex items-center gap-3 px-4 py-3 hover:bg-muted transition-colors group"
        >
          <span className="flex-shrink-0 w-7 h-7 rounded-full bg-primary/10 text-primary text-xs font-semibold flex items-center justify-center group-hover:bg-primary group-hover:text-primary-foreground transition-colors">
            {ch.chapterNumber}
          </span>
          <span className="flex-1 text-sm font-medium line-clamp-1 group-hover:text-primary transition-colors">
            {ch.chapterTitle || ch.title}
          </span>
          {ch.fileSize && (
            <span className="text-xs text-muted-foreground flex-shrink-0">
              {Math.round((ch.fileSize / 1024))} KB
            </span>
          )}
          <ChevronDown className="h-4 w-4 text-muted-foreground flex-shrink-0 -rotate-90 group-hover:text-primary transition-colors" />
        </Link>
      ))}
    </div>
  )
}

// ─── página principal ─────────────────────────────────────────────────────────

export default function ArticleDetailPage({ params }: PageProps) {
  const { id } = use(params)
  const router = useRouter()
  const { data: article, isLoading, isError } = useArticle(id)
  const toggleFavorite = useToggleFavorite()
  const setRating = useSetRating()
  const deleteArticle = useDeleteArticle()
  const [isFullContentExpanded, setIsFullContentExpanded] = useState(false)

  // Para capítulos: carrega o livro pai para breadcrumb
  const parentId = article?.parentArticleId ?? ''
  const { data: parentBook } = useArticle(parentId)

  const isBook = article?.sourceType === 'book'
  const isChapter = article?.sourceType === 'book_chapter'

  const handleToggleFavorite = async () => {
    if (!article) return
    try {
      await toggleFavorite.mutateAsync(article.id)
      toast.success(article.favorite ? 'Removido dos favoritos' : 'Adicionado aos favoritos')
    } catch {
      toast.error('Erro ao atualizar favorito')
    }
  }

  const handleSetRating = async (rating: number) => {
    if (!article) return
    try {
      await setRating.mutateAsync({ id: article.id, rating })
      toast.success(`Avaliação: ${rating} estrela${rating !== 1 ? 's' : ''}`)
    } catch {
      toast.error('Erro ao avaliar')
    }
  }

  const handleDelete = async () => {
    if (!article) return
    try {
      await deleteArticle.mutateAsync(article.id)
      toast.success(isBook ? 'Livro deletado' : 'Artigo deletado')
      if (isChapter && article.parentArticleId) {
        router.push(`/articles/${article.parentArticleId}`)
      } else {
        router.push('/articles')
      }
    } catch {
      toast.error('Erro ao deletar')
    }
  }

  const handleDownload = () => {
    if (!article) return
    window.open(articleApi.getDownloadUrl(article.id), '_blank')
  }

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 px-4 max-w-4xl space-y-4">
        <Skeleton className="h-8 w-32" />
        <Skeleton className="h-12 w-3/4" />
        <Skeleton className="h-6 w-1/2" />
        <Skeleton className="h-64" />
      </div>
    )
  }

  if (isError || !article) {
    return (
      <div className="container mx-auto py-8 px-4 max-w-4xl text-center">
        <h1 className="text-2xl font-bold mb-4">Não encontrado</h1>
        <Button asChild>
          <Link href="/articles">
            <ArrowLeft className="mr-2 h-4 w-4" />
            Voltar para artigos
          </Link>
        </Button>
      </div>
    )
  }

  const articleType = ARTICLE_TYPES.find((t) => t.value === article.articleType)
  const publishDate = new Date(article.publishDate)

  // ── breadcrumbs dinâmicos ──────────────────────────────────────────────────
  const breadcrumbs = isChapter
    ? [
        { label: 'Artigos', href: '/articles' },
        ...(parentBook
          ? [{ label: parentBook.title, href: `/articles/${parentBook.id}` }]
          : [{ label: 'Livro', href: article.parentArticleId ? `/articles/${article.parentArticleId}` : '/articles' }]),
        { label: `Cap. ${article.chapterNumber}` },
      ]
    : [{ label: 'Artigos', href: '/articles' }, { label: article.title }]

  // ── header description ─────────────────────────────────────────────────────
  const headerDescription = isBook
    ? `${article.totalChapters} capítulo${article.totalChapters !== 1 ? 's' : ''} • ${format(publishDate, "yyyy", { locale: ptBR })}`
    : isChapter
    ? `${parentBook?.title ?? 'Livro'} • ${format(publishDate, "yyyy", { locale: ptBR })}`
    : `${article.journal} • ${format(publishDate, "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}`

  return (
    <div className="container mx-auto py-8 px-4 max-w-4xl">
      {/* Header */}
      <PageHeader
        title={
          isChapter
            ? `Cap. ${article.chapterNumber}: ${article.chapterTitle || article.title}`
            : article.title
        }
        description={headerDescription}
        breadcrumbs={breadcrumbs}
      >
        <div className="flex flex-wrap items-center gap-2">
          {/* Botão de voltar ao livro para capítulos */}
          {isChapter && article.parentArticleId && (
            <Button variant="outline" asChild>
              <Link href={`/articles/${article.parentArticleId}`}>
                <ArrowLeft className="mr-2 h-4 w-4" />
                Voltar ao livro
              </Link>
            </Button>
          )}

          <Button
            variant={article.favorite ? 'default' : 'outline'}
            onClick={handleToggleFavorite}
            disabled={toggleFavorite.isPending}
          >
            <Heart className={`mr-2 h-4 w-4 ${article.favorite ? 'fill-current' : ''}`} />
            {article.favorite ? 'Favorito' : 'Favoritar'}
          </Button>

          {article.internalLink && (
            <Button variant="outline" onClick={handleDownload}>
              <Download className="mr-2 h-4 w-4" />
              Download
            </Button>
          )}

          {article.originalLink && (
            <Button variant="outline" asChild>
              <a href={article.originalLink} target="_blank" rel="noopener noreferrer">
                <ExternalLink className="mr-2 h-4 w-4" />
                Ver original
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
                  {isBook ? (
                    <>
                      Deletar o livro <strong>"{article.title}"</strong> também
                      remove todos os seus {article.totalChapters} capítulos permanentemente.
                    </>
                  ) : (
                    <>
                      Deletar <strong>"{article.title}"</strong>? Esta ação não pode
                      ser desfeita.
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
        </div>
      </PageHeader>

      {/* Badges de classificação */}
      <div className="mb-6">
        <div className="flex flex-wrap items-center gap-2 mb-4">
          {isBook ? (
            <>
              <Badge variant="outline" className="border-primary/40 text-primary gap-1">
                <BookOpen className="h-3 w-3" />
                Livro
              </Badge>
              {article.totalChapters && (
                <Badge variant="secondary" className="gap-1">
                  <List className="h-3 w-3" />
                  {article.totalChapters} capítulo{article.totalChapters !== 1 ? 's' : ''}
                </Badge>
              )}
            </>
          ) : isChapter ? (
            <>
              <Badge variant="outline" className="gap-1">
                <Hash className="h-3 w-3" />
                Capítulo {article.chapterNumber}
              </Badge>
              {parentBook && (
                <Badge variant="secondary" className="gap-1">
                  <BookOpen className="h-3 w-3" />
                  {parentBook.title}
                </Badge>
              )}
            </>
          ) : (
            <>
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
            </>
          )}

          {article.fileSize && (
            <span className="text-xs text-muted-foreground flex items-center gap-1">
              <FileText className="h-3.5 w-3.5" />
              {(article.fileSize / (1024 * 1024)).toFixed(1)} MB
            </span>
          )}
        </div>

        {/* Rating */}
        <div className="flex items-center gap-2">
          <span className="text-sm text-muted-foreground">Avaliação:</span>
          {[1, 2, 3, 4, 5].map((star) => (
            <button
              key={star}
              onClick={() => handleSetRating(star)}
              disabled={setRating.isPending}
              className="hover:scale-110 transition-transform disabled:opacity-50"
            >
              <Star
                className={`h-5 w-5 ${
                  article.rating && star <= article.rating
                    ? 'fill-yellow-400 text-yellow-400'
                    : 'text-muted-foreground'
                }`}
              />
            </button>
          ))}
        </div>
      </div>

      <Separator className="my-6" />

      {/* ── LIVRO: capítulos em destaque ─────────────────────────────────────── */}
      {isBook && (
        <div className="mb-8">
          <Card className="border-primary/20">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <List className="h-5 w-5 text-primary" />
                Capítulos
              </CardTitle>
              <CardDescription>
                {article.totalChapters} capítulo{article.totalChapters !== 1 ? 's' : ''} indexados e disponíveis para busca semântica
              </CardDescription>
            </CardHeader>
            <CardContent>
              <BookChapterList bookId={article.id} />
            </CardContent>
          </Card>
        </div>
      )}

      {/* File Viewer — PDF, EPUB, TXT, MD (artigos e livros, não capítulos) */}
      {article.internalLink && !isChapter && (
        <>
          <FileViewer
            internalLink={article.internalLink}
            downloadUrl={articleApi.getDownloadUrl(article.id)}
            title={article.title}
            fullContent={article.fullContent ?? undefined}
          />
          <Separator className="my-8" />
        </>
      )}

      {/* Content sections */}
      <div className="space-y-8">
        {/* Abstract */}
        {article.abstract && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">
                {isBook ? 'Descrição' : 'Resumo'}
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="whitespace-pre-wrap">{article.abstract}</p>
            </CardContent>
          </Card>
        )}

        {/* Autores (apenas artigos e capítulos com autor definido) */}
        {!isBook && article.authors && article.authors !== 'Desconhecido' && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg">Autores</CardTitle>
            </CardHeader>
            <CardContent>
              <p>{article.authors}</p>
            </CardContent>
          </Card>
        )}

        {/* Keywords & MeSH — apenas artigos */}
        {!isBook && !isChapter &&
          ((article.keywords && article.keywords.length > 0) ||
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
                      <Badge key={idx} variant="secondary">{keyword}</Badge>
                    ))}
                  </div>
                </div>
              )}
              {article.meshTerms && article.meshTerms.length > 0 && (
                <div>
                  <p className="text-sm font-medium mb-2">MeSH Terms:</p>
                  <div className="flex flex-wrap gap-2">
                    {article.meshTerms.map((term, idx) => (
                      <Badge key={idx} variant="outline">{term}</Badge>
                    ))}
                  </div>
                </div>
              )}
            </CardContent>
          </Card>
        )}

        {/* Embedding Status — não aplicável para livros (capítulos têm o seu) */}
        {!isBook && (
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Database className="h-5 w-5" />
                Status de Embedding (RAG)
              </CardTitle>
              <CardDescription>Processamento vetorial para busca semântica</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center justify-between">
                <span className="text-sm font-medium">Status:</span>
                <div className="flex items-center gap-2">
                  {(article as any).embeddingStatus === 'completed' && (
                    <>
                      <CheckCircle2 className="h-4 w-4 text-green-600" />
                      <Badge className="bg-green-600">Completo</Badge>
                    </>
                  )}
                  {(article as any).embeddingStatus === 'processing' && (
                    <>
                      <Loader2 className="h-4 w-4 text-blue-600 animate-spin" />
                      <Badge className="bg-blue-600">Processando</Badge>
                    </>
                  )}
                  {(article as any).embeddingStatus === 'failed' && (
                    <>
                      <XCircle className="h-4 w-4 text-red-600" />
                      <Badge variant="destructive">Falhou</Badge>
                    </>
                  )}
                  {(article as any).embeddingStatus === 'pending' && (
                    <>
                      <Clock className="h-4 w-4 text-yellow-600" />
                      <Badge variant="secondary">Pendente</Badge>
                    </>
                  )}
                </div>
              </div>

              {(article as any).chunkCount > 0 && (
                <div className="flex items-center justify-between">
                  <span className="text-sm font-medium">Chunks gerados:</span>
                  <Badge variant="outline">{(article as any).chunkCount} chunks</Badge>
                </div>
              )}

              {(article as any).lastEmbeddedAt && (
                <div className="flex items-center justify-between">
                  <span className="text-sm font-medium">Processado em:</span>
                  <span className="text-sm text-muted-foreground">
                    {format(new Date((article as any).lastEmbeddedAt), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
                  </span>
                </div>
              )}

              {(article as any).embeddingStatus === 'completed' && (
                <div className="pt-2 border-t">
                  <p className="text-xs text-muted-foreground">
                    {isChapter
                      ? 'Este capítulo está disponível para busca semântica via RAG.'
                      : 'Este artigo foi processado e está disponível para busca semântica via RAG.'}
                  </p>
                </div>
              )}
            </CardContent>
          </Card>
        )}

        {/* Para livros: nota explicativa sobre embedding */}
        {isBook && (
          <Card className="border-muted">
            <CardContent className="pt-6">
              <p className="text-sm text-muted-foreground flex items-start gap-2">
                <Database className="h-4 w-4 mt-0.5 flex-shrink-0" />
                O embedding e a busca semântica funcionam por capítulo. Cada capítulo
                é indexado individualmente para maior precisão na busca semântica.
              </p>
            </CardContent>
          </Card>
        )}

        {/* Score Items — artigos e capítulos */}
        {!isBook && article.scoreItems && article.scoreItems.length > 0 && (
          <ArticleScoreItems scoreItems={article.scoreItems} />
        )}

        {/* Related Score Items — artigos e capítulos */}
        {!isBook && (
          <RelatedScoreItems
            articleId={article.id}
            articleTitle={article.title}
            limit={10}
          />
        )}

        {/* Notas pessoais */}
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

        {/* Texto completo — apenas artigos e capítulos (livros não têm fullContent) */}
        {article.fullContent && (
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle className="text-lg">
                    {isChapter ? 'Texto do Capítulo' : 'Texto Completo'}
                  </CardTitle>
                  <CardDescription>
                    {article.fullContent.length.toLocaleString()} caracteres
                  </CardDescription>
                </div>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setIsFullContentExpanded(!isFullContentExpanded)}
                >
                  {isFullContentExpanded ? (
                    <><ChevronUp className="mr-2 h-4 w-4" />Colapsar</>
                  ) : (
                    <><ChevronDown className="mr-2 h-4 w-4" />Expandir</>
                  )}
                </Button>
              </div>
            </CardHeader>
            <CardContent>
              {isFullContentExpanded ? (
                <div className="max-h-[600px] overflow-y-auto border rounded-md p-4 bg-muted/30">
                  <p className="whitespace-pre-wrap text-sm font-mono">{article.fullContent}</p>
                </div>
              ) : (
                <div className="border rounded-md p-4 bg-muted/30">
                  <p className="whitespace-pre-wrap text-sm font-mono text-muted-foreground">
                    {article.fullContent.substring(0, 500)}...
                  </p>
                  <p className="text-xs text-muted-foreground mt-2">
                    Mostrando primeiros 500 caracteres.
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
