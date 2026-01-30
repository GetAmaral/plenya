'use client'

import { use, useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { useForm, Controller } from 'react-hook-form'
import { format } from 'date-fns'
import { toast } from 'sonner'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Badge } from '@/components/ui/badge'
import {
  useArticle,
  useUpdateArticle,
  useAddScoreItems,
  useRemoveScoreItems,
  useUploadArticle,
  ARTICLE_TYPES,
  UpdateArticleDTO,
} from '@/lib/api/article-api'
import { ArticleScoreItemsSelector } from '@/components/articles/ArticleScoreItemsSelector'
import { ArrowLeft, Save, Loader2, Upload, FileText, X } from 'lucide-react'
import { PageHeader } from '@/components/layout/page-header'

interface PageProps {
  params: Promise<{ id: string }>
}

export default function EditArticlePage({ params }: PageProps) {
  const { id } = use(params)
  const router = useRouter()
  const { data: article, isLoading, isError } = useArticle(id)
  const updateMutation = useUpdateArticle()
  const addScoreItemsMutation = useAddScoreItems()
  const removeScoreItemsMutation = useRemoveScoreItems()
  const uploadMutation = useUploadArticle()

  const [selectedScoreItemIds, setSelectedScoreItemIds] = useState<string[]>([])
  const [selectedFile, setSelectedFile] = useState<File | null>(null)

  // Atualizar selectedScoreItemIds quando article mudar
  useEffect(() => {
    if (article?.scoreItems) {
      setSelectedScoreItemIds(article.scoreItems.map((item) => item.id))
    }
  }, [article])

  const {
    register,
    handleSubmit,
    control,
    formState: { errors, isDirty },
  } = useForm<UpdateArticleDTO>({
    defaultValues: article
      ? {
          title: article.title,
          authors: article.authors,
          journal: article.journal,
          publishDate: format(new Date(article.publishDate), 'yyyy-MM-dd'),
          language: article.language,
          doi: article.doi || '',
          pmid: article.pmid || '',
          issn: article.issn || '',
          abstract: article.abstract || '',
          fullContent: article.fullContent || '',
          notes: article.notes || '',
          originalLink: article.originalLink || '',
          articleType: article.articleType,
          specialty: article.specialty || '',
        }
      : undefined,
    values: article
      ? {
          title: article.title,
          authors: article.authors,
          journal: article.journal,
          publishDate: format(new Date(article.publishDate), 'yyyy-MM-dd'),
          language: article.language,
          doi: article.doi || '',
          pmid: article.pmid || '',
          issn: article.issn || '',
          abstract: article.abstract || '',
          fullContent: article.fullContent || '',
          notes: article.notes || '',
          originalLink: article.originalLink || '',
          articleType: article.articleType,
          specialty: article.specialty || '',
        }
      : undefined,
  })

  const [keywordsInput, setKeywordsInput] = useState(
    article?.keywords?.join(', ') || ''
  )
  const [meshTermsInput, setMeshTermsInput] = useState(
    article?.meshTerms?.join(', ') || ''
  )

  const onSubmit = async (data: UpdateArticleDTO) => {
    try {
      // Parse keywords and meshTerms
      const keywords =
        keywordsInput
          .split(',')
          .map((k) => k.trim())
          .filter((k) => k.length > 0) || undefined

      const meshTerms =
        meshTermsInput
          .split(',')
          .map((m) => m.trim())
          .filter((m) => m.length > 0) || undefined

      // Convert publishDate from yyyy-MM-dd to ISO datetime
      let publishDate = data.publishDate
      if (publishDate && typeof publishDate === 'string' && !publishDate.includes('T')) {
        // Add time component to make it a valid ISO datetime
        publishDate = `${publishDate}T00:00:00Z`
      }

      // Clean up empty strings
      const cleanData: UpdateArticleDTO = {
        ...data,
        publishDate,
        doi: data.doi || undefined,
        pmid: data.pmid || undefined,
        issn: data.issn || undefined,
        abstract: data.abstract || undefined,
        fullContent: data.fullContent || undefined,
        notes: data.notes || undefined,
        originalLink: data.originalLink || undefined,
        specialty: data.specialty || undefined,
        keywords: keywords && keywords.length > 0 ? keywords : undefined,
        meshTerms: meshTerms && meshTerms.length > 0 ? meshTerms : undefined,
      }

      // Atualizar artigo
      await updateMutation.mutateAsync({ id, data: cleanData })

      // Upload de novo PDF se selecionado
      if (selectedFile) {
        const formData = new FormData()
        formData.append('file', selectedFile)
        // Adicionar metadados básicos ao FormData
        formData.append('title', data.title)
        formData.append('authors', data.authors)
        formData.append('journal', data.journal)
        formData.append('publishDate', publishDate || '')
        formData.append('language', data.language)
        formData.append('articleType', data.articleType)

        await uploadMutation.mutateAsync(formData)
      }

      // Atualizar score items se houve mudanças
      const originalItemIds = article?.scoreItems?.map((item) => item.id) || []
      const itemsToAdd = selectedScoreItemIds.filter(
        (itemId) => !originalItemIds.includes(itemId)
      )
      const itemsToRemove = originalItemIds.filter(
        (itemId) => !selectedScoreItemIds.includes(itemId)
      )

      if (itemsToAdd.length > 0) {
        await addScoreItemsMutation.mutateAsync({ id, scoreItemIds: itemsToAdd })
      }

      if (itemsToRemove.length > 0) {
        await removeScoreItemsMutation.mutateAsync({ id, scoreItemIds: itemsToRemove })
      }

      toast.success('Artigo atualizado com sucesso!')
      router.push(`/articles/${id}`)
    } catch (error: any) {
      toast.error('Erro ao atualizar artigo', {
        description: error.response?.data?.error || error.message,
      })
    }
  }

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 px-4 max-w-4xl">
        <Skeleton className="h-8 w-32 mb-8" />
        <Skeleton className="h-12 w-64 mb-8" />
        <div className="space-y-6">
          <Skeleton className="h-64" />
          <Skeleton className="h-64" />
          <Skeleton className="h-64" />
        </div>
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

  return (
    <div className="container mx-auto py-8 px-4 max-w-4xl">
      {/* Header */}
      <PageHeader
        title="Editar Artigo"
        description="Atualize as informações do artigo científico"
        breadcrumbs={[
          { label: "Artigos", href: "/articles" },
          { label: article.title, href: `/articles/${id}` },
          { label: "Editar" }
        ]}
      />

      {/* Form */}
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        {/* Basic Information */}
        <Card>
          <CardHeader>
            <CardTitle>Informações Básicas</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            {/* Title */}
            <div className="space-y-2">
              <Label htmlFor="title">
                Título <span className="text-destructive">*</span>
              </Label>
              <Input
                id="title"
                {...register('title', {
                  required: 'Título é obrigatório',
                  minLength: { value: 3, message: 'Mínimo 3 caracteres' },
                  maxLength: { value: 1000, message: 'Máximo 1000 caracteres' },
                })}
                className={errors.title ? 'border-destructive' : ''}
              />
              {errors.title && (
                <p className="text-sm text-destructive">{errors.title.message}</p>
              )}
            </div>

            {/* Authors */}
            <div className="space-y-2">
              <Label htmlFor="authors">
                Autores <span className="text-destructive">*</span>
              </Label>
              <Input
                id="authors"
                placeholder="Ex: Silva JA, Santos MF, Oliveira PR"
                {...register('authors', {
                  required: 'Autores são obrigatórios',
                  minLength: { value: 1, message: 'Mínimo 1 caractere' },
                  maxLength: { value: 2000, message: 'Máximo 2000 caracteres' },
                })}
                className={errors.authors ? 'border-destructive' : ''}
              />
              {errors.authors && (
                <p className="text-sm text-destructive">
                  {errors.authors.message}
                </p>
              )}
            </div>

            {/* Journal & Date */}
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="journal">
                  Revista <span className="text-destructive">*</span>
                </Label>
                <Input
                  id="journal"
                  placeholder="Ex: Nature Medicine"
                  {...register('journal', {
                    required: 'Revista é obrigatória',
                    minLength: { value: 1, message: 'Mínimo 1 caractere' },
                    maxLength: { value: 500, message: 'Máximo 500 caracteres' },
                  })}
                  className={errors.journal ? 'border-destructive' : ''}
                />
                {errors.journal && (
                  <p className="text-sm text-destructive">
                    {errors.journal.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="publishDate">
                  Data de Publicação <span className="text-destructive">*</span>
                </Label>
                <Input
                  id="publishDate"
                  type="date"
                  {...register('publishDate', {
                    required: 'Data é obrigatória',
                  })}
                  className={errors.publishDate ? 'border-destructive' : ''}
                />
                {errors.publishDate && (
                  <p className="text-sm text-destructive">
                    {errors.publishDate.message}
                  </p>
                )}
              </div>
            </div>

            {/* Type & Language */}
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="articleType">
                  Tipo de Artigo <span className="text-destructive">*</span>
                </Label>
                <Controller
                  name="articleType"
                  control={control}
                  rules={{ required: 'Tipo é obrigatório' }}
                  render={({ field }) => (
                    <Select value={field.value || ''} onValueChange={field.onChange}>
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        {ARTICLE_TYPES.map((type) => (
                          <SelectItem key={type.value} value={type.value}>
                            {type.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  )}
                />
                {errors.articleType && (
                  <p className="text-sm text-destructive">
                    {errors.articleType.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="language">
                  Idioma <span className="text-destructive">*</span>
                </Label>
                <Input
                  id="language"
                  placeholder="Ex: en, pt-BR"
                  {...register('language', {
                    required: 'Idioma é obrigatório',
                    minLength: { value: 2, message: 'Mínimo 2 caracteres' },
                    maxLength: { value: 10, message: 'Máximo 10 caracteres' },
                  })}
                  className={errors.language ? 'border-destructive' : ''}
                />
                {errors.language && (
                  <p className="text-sm text-destructive">
                    {errors.language.message}
                  </p>
                )}
              </div>
            </div>

            {/* Specialty */}
            <div className="space-y-2">
              <Label htmlFor="specialty">Especialidade</Label>
              <Input
                id="specialty"
                placeholder="Ex: Cardiologia, Neurologia"
                {...register('specialty', {
                  maxLength: { value: 200, message: 'Máximo 200 caracteres' },
                })}
              />
            </div>
          </CardContent>
        </Card>

        {/* Identifiers */}
        <Card>
          <CardHeader>
            <CardTitle>Identificadores</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-3 gap-4">
              <div className="space-y-2">
                <Label htmlFor="doi">DOI</Label>
                <Input
                  id="doi"
                  placeholder="10.1038/..."
                  {...register('doi', {
                    maxLength: { value: 255, message: 'Máximo 255 caracteres' },
                  })}
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="pmid">PMID</Label>
                <Input
                  id="pmid"
                  placeholder="38355808"
                  {...register('pmid', {
                    maxLength: { value: 50, message: 'Máximo 50 caracteres' },
                  })}
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="issn">ISSN</Label>
                <Input
                  id="issn"
                  placeholder="1546-170X"
                  {...register('issn', {
                    maxLength: { value: 20, message: 'Máximo 20 caracteres' },
                  })}
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label htmlFor="originalLink">Link da Publicação Original</Label>
              <Input
                id="originalLink"
                type="url"
                placeholder="https://..."
                {...register('originalLink', {
                  maxLength: { value: 2048, message: 'Máximo 2048 caracteres' },
                })}
              />
            </div>
          </CardContent>
        </Card>

        {/* PDF Upload */}
        <Card>
          <CardHeader>
            <CardTitle>Arquivo PDF</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            {/* Current PDF Info */}
            {article.internalLink && !selectedFile && (
              <div className="flex items-center gap-3 p-3 bg-muted/50 rounded-lg border">
                <FileText className="h-8 w-8 text-primary" />
                <div className="flex-1">
                  <p className="text-sm font-medium">PDF atual</p>
                  <p className="text-xs text-muted-foreground">
                    {article.fileSize
                      ? `${(article.fileSize / (1024 * 1024)).toFixed(2)} MB`
                      : 'Tamanho desconhecido'}
                  </p>
                </div>
                <Badge variant="secondary">Disponível</Badge>
              </div>
            )}

            {/* New File Selected */}
            {selectedFile && (
              <div className="flex items-center gap-3 p-3 bg-primary/10 rounded-lg border border-primary">
                <Upload className="h-8 w-8 text-primary" />
                <div className="flex-1">
                  <p className="text-sm font-medium">Novo arquivo selecionado</p>
                  <p className="text-xs text-muted-foreground">
                    {selectedFile.name} ({(selectedFile.size / (1024 * 1024)).toFixed(2)} MB)
                  </p>
                </div>
                <Button
                  type="button"
                  variant="ghost"
                  size="sm"
                  onClick={() => setSelectedFile(null)}
                >
                  <X className="h-4 w-4" />
                </Button>
              </div>
            )}

            {/* Upload Button */}
            <div className="space-y-2">
              <Label htmlFor="pdfFile">
                {article.internalLink ? 'Substituir PDF' : 'Upload de PDF'}
              </Label>
              <div className="flex gap-2">
                <Input
                  id="pdfFile"
                  type="file"
                  accept=".pdf,application/pdf"
                  onChange={(e) => {
                    const file = e.target.files?.[0]
                    if (file) {
                      if (file.type !== 'application/pdf') {
                        toast.error('Apenas arquivos PDF são permitidos')
                        e.target.value = ''
                        return
                      }
                      if (file.size > 50 * 1024 * 1024) {
                        toast.error('Arquivo muito grande. Máximo: 50MB')
                        e.target.value = ''
                        return
                      }
                      setSelectedFile(file)
                    }
                  }}
                  className="cursor-pointer"
                />
              </div>
              <p className="text-xs text-muted-foreground">
                Formato: PDF | Tamanho máximo: 50MB
                {article.internalLink && ' | Substituir o arquivo atual'}
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Content */}
        <Card>
          <CardHeader>
            <CardTitle>Conteúdo</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            {/* Abstract */}
            <div className="space-y-2">
              <Label htmlFor="abstract">Resumo (Abstract)</Label>
              <Textarea
                id="abstract"
                rows={6}
                placeholder="Digite o resumo do artigo..."
                {...register('abstract')}
              />
            </div>

            {/* Full Content */}
            <div className="space-y-2">
              <Label htmlFor="fullContent">
                Texto Completo Extraído do PDF
              </Label>
              <Textarea
                id="fullContent"
                rows={10}
                placeholder="Texto completo extraído automaticamente do PDF..."
                className="font-mono text-xs"
                {...register('fullContent')}
              />
              <p className="text-xs text-muted-foreground">
                Este campo é preenchido automaticamente ao fazer upload do PDF, mas pode ser editado manualmente
              </p>
            </div>

            {/* Notes */}
            <div className="space-y-2">
              <Label htmlFor="notes">Notas Pessoais</Label>
              <Textarea
                id="notes"
                rows={4}
                placeholder="Suas anotações sobre o artigo..."
                {...register('notes')}
              />
            </div>
          </CardContent>
        </Card>

        {/* Keywords & MeSH */}
        <Card>
          <CardHeader>
            <CardTitle>Palavras-chave e Termos</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="keywords">Keywords</Label>
              <Input
                id="keywords"
                placeholder="cardiovascular, hypertension, risk factors (separadas por vírgula)"
                value={keywordsInput}
                onChange={(e) => setKeywordsInput(e.target.value)}
              />
              <p className="text-xs text-muted-foreground">
                Separe as palavras-chave por vírgula
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="meshTerms">MeSH Terms</Label>
              <Input
                id="meshTerms"
                placeholder="Cardiovascular Diseases, Hypertension (separados por vírgula)"
                value={meshTermsInput}
                onChange={(e) => setMeshTermsInput(e.target.value)}
              />
              <p className="text-xs text-muted-foreground">
                Separe os termos MeSH por vírgula
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Score Items */}
        <Card>
          <CardHeader>
            <CardTitle>Vinculação com Score Items</CardTitle>
          </CardHeader>
          <CardContent>
            <ArticleScoreItemsSelector
              selectedItemIds={selectedScoreItemIds}
              onChange={setSelectedScoreItemIds}
            />
          </CardContent>
        </Card>

        {/* Actions */}
        <div className="flex items-center justify-between">
          <Button variant="outline" type="button" asChild>
            <Link href={`/articles/${id}`}>Cancelar</Link>
          </Button>

          <Button
            type="submit"
            disabled={!isDirty || updateMutation.isPending}
          >
            {updateMutation.isPending ? (
              <>
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                Salvando...
              </>
            ) : (
              <>
                <Save className="mr-2 h-4 w-4" />
                Salvar Alterações
              </>
            )}
          </Button>
        </div>
      </form>
    </div>
  )
}
