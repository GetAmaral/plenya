'use client'

import { useState, useRef } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Badge } from '@/components/ui/badge'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible'
import {
  Calendar,
  FileText,
  Eye,
  EyeOff,
  Lock,
  StickyNote,
  Sparkles,
  ChevronDown,
  ChevronUp,
  Check,
  ClipboardList,
  Save,
} from 'lucide-react'
import { cn } from '@/lib/utils'
import {
  createAnamnesis,
  type CreateAnamnesisRequest,
} from '@/lib/api/anamnesis'
import {
  getAllAnamnesisTemplates,
  getAnamnesisTemplateById,
  type AnamnesisTemplate,
} from '@/lib/api/anamnesis-templates'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import { RichTextEditor } from '@/components/ui/rich-text-editor'
import { htmlToPlainText } from '@/lib/html-utils'
import { AnamnesisFullscreenTemplateItems } from './AnamnesisFullscreenTemplateItems'
import type { AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'

interface AnamnesisFullscreenFormProps {
  onSuccess: () => void
  onCancel: () => void
}

export function AnamnesisFullscreenForm({ onSuccess, onCancel }: AnamnesisFullscreenFormProps) {
  // Form state
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>('')
  const [selectedTemplate, setSelectedTemplate] = useState<AnamnesisTemplate | null>(null)
  const [consultationDate, setConsultationDate] = useState(
    format(new Date(), "yyyy-MM-dd'T'HH:mm")
  )
  const [summary, setSummary] = useState('')
  const [content, setContent] = useState('')
  const [visibility, setVisibility] = useState<'all' | 'medicalOnly' | 'psychOnly' | 'authorOnly'>('all')
  const [notes, setNotes] = useState('')
  const [templateItems, setTemplateItems] = useState<AnamnesisItemFormValue[]>([])
  const [richTextEnabled, setRichTextEnabled] = useState(false)

  // UI state
  const [templateDialogOpen, setTemplateDialogOpen] = useState(false)
  const [summaryExpanded, setSummaryExpanded] = useState(true)
  const [contentExpanded, setContentExpanded] = useState(false)
  const [detailsExpanded, setDetailsExpanded] = useState(false)
  const [itemsExpanded, setItemsExpanded] = useState(false)

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates
  const { data: templates = [] } = useQuery({
    queryKey: ['anamnesis-templates'],
    queryFn: () => getAllAnamnesisTemplates(false),
  })

  const sortedTemplates = [...templates].sort((a, b) => a.name.localeCompare(b.name))

  // Create mutation
  const createMutation = useMutation({
    mutationFn: createAnamnesis,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis'] })
      toast.success('Anamnese criada com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      if (error.response?.status === 403) {
        toast.error('Você não tem permissão para criar esta anamnese')
      } else {
        toast.error(error.response?.data?.error || 'Erro ao criar anamnese')
      }
    },
  })

  // Toggle rich text
  const toggleRichText = () => {
    if (richTextEnabled) {
      setSummary(htmlToPlainText(summary))
      if (content) {
        setContent(htmlToPlainText(content))
      }
    }
    setRichTextEnabled(!richTextEnabled)
  }

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSummary('')
      setContent('')
      setSelectedTemplateId('')
      setSelectedTemplate(null)
      setTemplateItems([])
      setTemplateDialogOpen(false)
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getAnamnesisTemplateById(templateId, true)
      setSelectedTemplate(template)
      if (!summary) {
        setSummary(`Anamnese - ${template.name}`)
      }
      setTemplateDialogOpen(false)
      toast.success(`Template "${template.name}" selecionado`)
    } catch (error) {
      console.error('Erro ao carregar template:', error)
      toast.error('Erro ao carregar template')
      setSelectedTemplate(null)
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    const dateObj = new Date(consultationDate)
    const rfc3339Date = dateObj.toISOString()

    const summaryPlain = richTextEnabled ? htmlToPlainText(summary) : summary
    const summaryHtml = richTextEnabled ? summary.trim() : undefined
    const contentPlain = content ? (richTextEnabled ? htmlToPlainText(content) : content) : undefined
    const contentHtml = richTextEnabled ? content.trim() : undefined

    const payload: CreateAnamnesisRequest = {
      consultationDate: rfc3339Date,
      summary: summaryPlain || undefined,
      summaryHtml: summaryHtml || undefined,
      content: contentPlain || undefined,
      contentHtml: contentHtml || undefined,
      visibility,
      notes: notes.trim() || undefined,
      anamnesisTemplateId: selectedTemplateId || undefined,
      items: templateItems.length > 0 ? templateItems.map((item) => ({
        scoreItemId: item.scoreItemId,
        numericValue: item.numericValue,
        textValue: item.textValue,
        order: item.order,
      })) : undefined,
    }

    createMutation.mutate(payload)
  }

  const getVisibilityIcon = (vis: typeof visibility) => {
    switch (vis) {
      case 'all':
        return Eye
      case 'medicalOnly':
      case 'psychOnly':
        return Lock
      default:
        return EyeOff
    }
  }

  const VisibilityIcon = getVisibilityIcon(visibility)

  return (
    <div className="h-full flex">
      {/* Main form - scrollable */}
      <ScrollArea className="flex-1">
        <form ref={formRef} onSubmit={handleSubmit} className="max-w-4xl mx-auto p-8 space-y-6">
          {/* Quick Info Bar */}
          <Card className="bg-primary/5 border-primary/20">
            <CardContent className="p-4">
              <div className="flex items-center justify-between gap-4 flex-wrap">
                <div className="flex items-center gap-4">
                  <div className="flex items-center gap-2">
                    <Calendar className="h-4 w-4 text-muted-foreground" />
                    <span className="text-sm font-medium">
                      {format(new Date(consultationDate), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
                    </span>
                  </div>
                  {selectedTemplate && (
                    <div className="flex items-center gap-2">
                      <FileText className="h-4 w-4 text-muted-foreground" />
                      <span className="text-sm font-medium">{selectedTemplate.name}</span>
                    </div>
                  )}
                </div>
                <div className="flex items-center gap-2">
                  <VisibilityIcon className="h-4 w-4 text-muted-foreground" />
                  <Badge variant="outline" className="gap-1.5">
                    {visibility === 'all' && 'Todos'}
                    {visibility === 'medicalOnly' && 'Apenas Médicos'}
                    {visibility === 'psychOnly' && 'Apenas Psicólogos'}
                    {visibility === 'authorOnly' && 'Apenas Autor'}
                  </Badge>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Section 1: Data e Template */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base flex items-center gap-2">
                <Calendar className="h-5 w-5" />
                Data e Template
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <Label className="text-base">Data da Consulta *</Label>
                <Input
                  type="datetime-local"
                  value={consultationDate}
                  onChange={(e) => setConsultationDate(e.target.value)}
                  required
                  className="text-base h-12 mt-2"
                />
              </div>

              <div>
                <Label className="text-base">Template (opcional)</Label>
                <Dialog open={templateDialogOpen} onOpenChange={setTemplateDialogOpen}>
                  <DialogTrigger asChild>
                    <Button
                      type="button"
                      variant="outline"
                      className="w-full h-12 mt-2 text-base justify-between"
                    >
                      <span className="flex items-center gap-2">
                        <FileText className="h-4 w-4" />
                        {selectedTemplate ? selectedTemplate.name : 'Selecionar template'}
                      </span>
                      <ClipboardList className="h-4 w-4" />
                    </Button>
                  </DialogTrigger>
                  <DialogContent className="max-w-2xl max-h-[80vh]">
                    <DialogHeader>
                      <DialogTitle>Selecionar Template</DialogTitle>
                      <DialogDescription>
                        Escolha um template para preencher automaticamente
                      </DialogDescription>
                    </DialogHeader>
                    <ScrollArea className="h-[400px] pr-4">
                      <div className="space-y-2">
                        <button
                          type="button"
                          onClick={() => handleTemplateSelect('')}
                          className={cn(
                            'w-full p-4 rounded-lg border-2 text-left transition-all',
                            selectedTemplateId === ''
                              ? 'border-primary bg-primary/5'
                              : 'border-border hover:border-primary/50'
                          )}
                        >
                          <div className="flex items-center justify-between">
                            <span className="font-medium">Nenhum template</span>
                            {selectedTemplateId === '' && <Check className="h-5 w-5 text-primary" />}
                          </div>
                        </button>
                        {sortedTemplates.map((template) => (
                          <button
                            key={template.id}
                            type="button"
                            onClick={() => handleTemplateSelect(template.id)}
                            className={cn(
                              'w-full p-4 rounded-lg border-2 text-left transition-all',
                              selectedTemplateId === template.id
                                ? 'border-primary bg-primary/5'
                                : 'border-border hover:border-primary/50'
                            )}
                          >
                            <div className="flex items-start justify-between gap-2">
                              <div className="flex-1">
                                <div className="font-medium text-base">{template.name}</div>
                                <div className="text-sm text-muted-foreground mt-1">{template.area}</div>
                              </div>
                              {selectedTemplateId === template.id && (
                                <Check className="h-5 w-5 text-primary flex-shrink-0" />
                              )}
                            </div>
                          </button>
                        ))}
                      </div>
                    </ScrollArea>
                  </DialogContent>
                </Dialog>
              </div>
            </CardContent>
          </Card>

          {/* Section 2: Resumo */}
          <Collapsible open={summaryExpanded} onOpenChange={setSummaryExpanded}>
            <Card>
              <CollapsibleTrigger asChild>
                <CardHeader className="cursor-pointer hover:bg-accent/50 transition-colors">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-base flex items-center gap-2">
                      <FileText className="h-5 w-5" />
                      Resumo
                    </CardTitle>
                    {summaryExpanded ? (
                      <ChevronUp className="h-5 w-5 text-muted-foreground" />
                    ) : (
                      <ChevronDown className="h-5 w-5 text-muted-foreground" />
                    )}
                  </div>
                </CardHeader>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <CardContent className="space-y-4">
                  {richTextEnabled ? (
                    <RichTextEditor
                      editorId="fullscreen-summary"
                      value={summary}
                      onChange={setSummary}
                      placeholder="Resumo breve da anamnese..."
                      minHeight="150px"
                    />
                  ) : (
                    <Textarea
                      value={summary}
                      onChange={(e) => setSummary(e.target.value)}
                      placeholder="Resumo breve da anamnese..."
                      rows={5}
                      className="text-base resize-none"
                    />
                  )}
                  <p className="text-sm text-muted-foreground">
                    {richTextEnabled ? 'Resumo com formatação rica' : 'Resumo em texto simples'}
                  </p>
                </CardContent>
              </CollapsibleContent>
            </Card>
          </Collapsible>

          {/* Section 3: Conteúdo Detalhado */}
          <Collapsible open={contentExpanded} onOpenChange={setContentExpanded}>
            <Card>
              <CollapsibleTrigger asChild>
                <CardHeader className="cursor-pointer hover:bg-accent/50 transition-colors">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-base flex items-center gap-2">
                      <FileText className="h-5 w-5" />
                      Conteúdo Detalhado
                    </CardTitle>
                    {contentExpanded ? (
                      <ChevronUp className="h-5 w-5 text-muted-foreground" />
                    ) : (
                      <ChevronDown className="h-5 w-5 text-muted-foreground" />
                    )}
                  </div>
                </CardHeader>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <CardContent className="space-y-4">
                  {richTextEnabled ? (
                    <RichTextEditor
                      editorId="fullscreen-content"
                      value={content}
                      onChange={setContent}
                      placeholder="Conteúdo detalhado da anamnese..."
                      minHeight="300px"
                    />
                  ) : (
                    <Textarea
                      value={content}
                      onChange={(e) => setContent(e.target.value)}
                      placeholder="Conteúdo detalhado da anamnese..."
                      rows={10}
                      className="text-base resize-none"
                    />
                  )}
                  <p className="text-sm text-muted-foreground">
                    {richTextEnabled ? 'Conteúdo com formatação rica' : 'Conteúdo em texto simples'}
                  </p>
                </CardContent>
              </CollapsibleContent>
            </Card>
          </Collapsible>

          {/* Section 4: Detalhes Adicionais */}
          <Collapsible open={detailsExpanded} onOpenChange={setDetailsExpanded}>
            <Card>
              <CollapsibleTrigger asChild>
                <CardHeader className="cursor-pointer hover:bg-accent/50 transition-colors">
                  <div className="flex items-center justify-between">
                    <CardTitle className="text-base flex items-center gap-2">
                      <StickyNote className="h-5 w-5" />
                      Detalhes Adicionais
                    </CardTitle>
                    {detailsExpanded ? (
                      <ChevronUp className="h-5 w-5 text-muted-foreground" />
                    ) : (
                      <ChevronDown className="h-5 w-5 text-muted-foreground" />
                    )}
                  </div>
                </CardHeader>
              </CollapsibleTrigger>
              <CollapsibleContent>
                <CardContent className="space-y-4">
                  <div>
                    <Label className="text-base">Visibilidade *</Label>
                    <Select value={visibility} onValueChange={(v: any) => setVisibility(v)}>
                      <SelectTrigger className="h-12 mt-2 text-base">
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="all" className="text-base">
                          <div className="flex items-center gap-2">
                            <Eye className="h-4 w-4" />
                            Todos os Profissionais
                          </div>
                        </SelectItem>
                        <SelectItem value="medicalOnly" className="text-base">
                          <div className="flex items-center gap-2">
                            <Lock className="h-4 w-4" />
                            Apenas Médicos
                          </div>
                        </SelectItem>
                        <SelectItem value="psychOnly" className="text-base">
                          <div className="flex items-center gap-2">
                            <Lock className="h-4 w-4" />
                            Apenas Psicólogos
                          </div>
                        </SelectItem>
                        <SelectItem value="authorOnly" className="text-base">
                          <div className="flex items-center gap-2">
                            <EyeOff className="h-4 w-4" />
                            Apenas Autor
                          </div>
                        </SelectItem>
                      </SelectContent>
                    </Select>
                  </div>

                  <div>
                    <Label className="text-base">Observações (opcional)</Label>
                    <Textarea
                      value={notes}
                      onChange={(e) => setNotes(e.target.value)}
                      placeholder="Observações adicionais..."
                      rows={4}
                      className="text-base resize-none mt-2"
                    />
                  </div>
                </CardContent>
              </CollapsibleContent>
            </Card>
          </Collapsible>

          {/* Section 5: Template Items */}
          {selectedTemplate && selectedTemplate.items && selectedTemplate.items.length > 0 && (
            <Collapsible open={itemsExpanded} onOpenChange={setItemsExpanded}>
              <Card>
                <CollapsibleTrigger asChild>
                  <CardHeader className="cursor-pointer hover:bg-accent/50 transition-colors">
                    <div className="flex items-center justify-between">
                      <CardTitle className="text-base flex items-center gap-2">
                        <ClipboardList className="h-5 w-5" />
                        Items do Template
                        {templateItems.length > 0 && (
                          <Badge variant="secondary" className="ml-2">
                            {templateItems.length} preenchido{templateItems.length !== 1 ? 's' : ''}
                          </Badge>
                        )}
                      </CardTitle>
                      {itemsExpanded ? (
                        <ChevronUp className="h-5 w-5 text-muted-foreground" />
                      ) : (
                        <ChevronDown className="h-5 w-5 text-muted-foreground" />
                      )}
                    </div>
                  </CardHeader>
                </CollapsibleTrigger>
                <CollapsibleContent>
                  <CardContent>
                    <AnamnesisFullscreenTemplateItems
                      key={selectedTemplate.id}
                      template={selectedTemplate}
                      initialValues={templateItems}
                      onChange={setTemplateItems}
                    />
                  </CardContent>
                </CollapsibleContent>
              </Card>
            </Collapsible>
          )}

          {/* Action Buttons */}
          <Card className="bg-muted/50">
            <CardContent className="p-6">
              <div className="flex items-center justify-between gap-4">
                <Button
                  type="button"
                  variant={richTextEnabled ? 'default' : 'outline'}
                  onClick={toggleRichText}
                  className="gap-2 h-12 px-6"
                >
                  <Sparkles className="h-4 w-4" />
                  {richTextEnabled ? 'Formatação Ativada' : 'Ativar Formatação'}
                </Button>
                <div className="flex gap-3">
                  <Button
                    type="button"
                    variant="outline"
                    onClick={onCancel}
                    className="h-12 px-6 text-base"
                  >
                    Cancelar
                  </Button>
                  <Button
                    type="submit"
                    disabled={createMutation.isPending}
                    className="h-12 px-8 text-base gap-2"
                  >
                    <Save className="h-4 w-4" />
                    {createMutation.isPending ? 'Salvando...' : 'Salvar Anamnese'}
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        </form>
      </ScrollArea>
    </div>
  )
}
