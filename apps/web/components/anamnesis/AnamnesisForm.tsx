'use client'

import { useState, useRef, useEffect } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Card } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Check, Type, Sparkles } from 'lucide-react'
import { ScrollArea } from '@/components/ui/scroll-area'
import { cn } from '@/lib/utils'
import {
  createAnamnesis,
  updateAnamnesis,
  type Anamnesis,
  type CreateAnamnesisRequest,
} from '@/lib/api/anamnesis'
import {
  getAllAnamnesisTemplates,
  getAnamnesisTemplateById,
  type AnamnesisTemplate,
} from '@/lib/api/anamnesis-templates'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { RichTextEditor } from '@/components/ui/rich-text-editor'
import { htmlToPlainText, hasTextContent } from '@/lib/html-utils'
import { AnamnesisTemplateItemsForm, type AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'

interface CreateAnamnesisFormProps {
  onSuccess: () => void
}

export function CreateAnamnesisForm({ onSuccess }: CreateAnamnesisFormProps) {
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
  const focusScoreItemId: string | null | undefined = undefined // Not used in create form

  // Rich text editor toggle (single toggle for both fields)
  const [richTextEnabled, setRichTextEnabled] = useState(false)

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Toggle rich text for both fields
  const toggleRichText = () => {
    if (richTextEnabled) {
      // Converting from rich text to plain text for both fields
      setSummary(htmlToPlainText(summary))
      if (content) {
        setContent(htmlToPlainText(content))
      }
    }
    setRichTextEnabled(!richTextEnabled)
  }

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

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSummary('')
      setContent('')
      setSelectedTemplateId('')
      setSelectedTemplate(null)
      setTemplateItems([])
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getAnamnesisTemplateById(templateId, true) // withItems = true
      setSelectedTemplate(template)
      // Set summary with template name
      if (!summary) {
        setSummary(`Anamnese - ${template.name}`)
      }
    } catch (error) {
      console.error('Erro ao carregar template:', error)
      toast.error('Erro ao carregar template')
      setSelectedTemplate(null)
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Convert to RFC3339 format
    const dateObj = new Date(consultationDate)
    const rfc3339Date = dateObj.toISOString()

    // Prepare data based on whether rich text is enabled
    const summaryPlain = richTextEnabled ? htmlToPlainText(summary) : summary
    const summaryHtml = richTextEnabled ? summary.trim() : undefined
    const contentPlain = content ? (richTextEnabled ? htmlToPlainText(content) : content) : undefined
    const contentHtml = richTextEnabled ? content.trim() : undefined

    const payload: CreateAnamnesisRequest = {
      consultationDate: rfc3339Date,
      summary: summaryPlain || undefined, // Plain text
      summaryHtml: summaryHtml || undefined, // HTML (only if rich text enabled)
      content: contentPlain || undefined, // Plain text
      contentHtml: contentHtml || undefined, // HTML (only if rich text enabled)
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

    // Backend gets patientId automatically from selectedPatient in JWT
    createMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Nova Anamnese</h2>
        <Button
          type="button"
          variant={richTextEnabled ? "default" : "outline"}
          size="sm"
          onClick={toggleRichText}
          className="gap-2"
        >
          <Sparkles className="h-4 w-4" />
          {richTextEnabled ? 'Formatação Ativada' : 'Ativar Formatação'}
        </Button>
      </div>
      <form ref={formRef} onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Main Column - Form */}
          <div className="lg:col-span-2 space-y-4">
            <div>
              <Label>Data da Consulta *</Label>
              <Input
                type="datetime-local"
                value={consultationDate}
                onChange={(e) => setConsultationDate(e.target.value)}
                required
              />
            </div>

            <div>
              <Label>Resumo (opcional)</Label>
              {richTextEnabled ? (
                <RichTextEditor editorId="create-summary"
                  value={summary}
                  onChange={setSummary}
                  placeholder="Resumo breve da anamnese..."
                  minHeight="120px"
                />
              ) : (
                <Textarea
                  value={summary}
                  onChange={(e) => setSummary(e.target.value)}
                  placeholder="Resumo breve da anamnese..."
                  rows={4}
                />
              )}
              <p className="text-xs text-muted-foreground mt-1">
                {richTextEnabled
                  ? 'Resumo formatado da anamnese'
                  : 'Resumo em texto simples'}
              </p>
            </div>

            <div>
              <Label>Conteúdo Detalhado (opcional)</Label>
              {richTextEnabled ? (
                <RichTextEditor editorId="create-content"
                  value={content}
                  onChange={setContent}
                  placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                  minHeight="300px"
                />
              ) : (
                <Textarea
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                  rows={8}
                />
              )}
              <p className="text-xs text-muted-foreground mt-1">
                {richTextEnabled
                  ? 'Anamnese completa com formatação rica'
                  : 'Anamnese completa em texto simples'}
              </p>
            </div>

            <div>
              <Label>Visibilidade *</Label>
              <Select value={visibility} onValueChange={(v: any) => setVisibility(v)}>
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">Todos os Profissionais</SelectItem>
                  <SelectItem value="medicalOnly">Apenas Médicos</SelectItem>
                  <SelectItem value="psychOnly">Apenas Psicólogos</SelectItem>
                  <SelectItem value="authorOnly">Apenas Autor</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div>
              <Label>Observações (opcional)</Label>
              <Textarea
                value={notes}
                onChange={(e) => setNotes(e.target.value)}
                placeholder="Observações adicionais..."
                rows={3}
              />
            </div>

            {/* Template Items Form */}
            {selectedTemplate && selectedTemplate.items && selectedTemplate.items.length > 0 && (
              <div className="pt-4 border-t">
                <h3 className="text-lg font-semibold mb-4">
                  Items do Template: {selectedTemplate.name}
                </h3>
                <AnamnesisTemplateItemsForm
                  key={selectedTemplate.id}
                  template={selectedTemplate}
                  initialValues={templateItems}
                  onChange={setTemplateItems}
                  focusScoreItemId={focusScoreItemId}
                />
              </div>
            )}

            <div className="flex justify-end gap-2 pt-4 border-t">
              <Button type="button" variant="outline" onClick={onSuccess}>
                Cancelar
              </Button>
              <Button type="submit" disabled={createMutation.isPending}>
                {createMutation.isPending ? 'Criando...' : 'Criar Anamnese'}
              </Button>
            </div>
          </div>

          {/* Sidebar - Templates */}
          <div className="lg:col-span-1">
            <Label className="mb-3 block">Templates (opcional)</Label>
            <ScrollArea className="h-[600px] pr-4">
              <div className="space-y-2">
                <div
                  onClick={() => handleTemplateSelect('')}
                  className={cn(
                    'relative p-3 rounded-lg cursor-pointer transition-all',
                    'border-2',
                    selectedTemplateId === ''
                      ? 'border-primary bg-primary/5'
                      : 'border-border hover:border-primary/50 hover:bg-accent'
                  )}
                >
                  <div className="flex items-center justify-between">
                    <span className="font-medium text-sm">Nenhum template</span>
                    {selectedTemplateId === '' && <Check className="h-4 w-4 text-primary" />}
                  </div>
                </div>
                {sortedTemplates.map((template) => (
                  <div
                    key={template.id}
                    onClick={() => handleTemplateSelect(template.id)}
                    className={cn(
                      'relative p-3 rounded-lg cursor-pointer transition-all',
                      'border-2',
                      selectedTemplateId === template.id
                        ? 'border-primary bg-primary/5'
                        : 'border-border hover:border-primary/50 hover:bg-accent'
                    )}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium text-sm">{template.name}</div>
                        <div className="text-xs text-muted-foreground mt-1">{template.area}</div>
                      </div>
                      {selectedTemplateId === template.id && (
                        <Check className="h-4 w-4 text-primary flex-shrink-0 mt-0.5" />
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
            {selectedTemplateId && (
              <p className="text-xs text-muted-foreground mt-3">
                Template selecionado aplicado ao resumo
              </p>
            )}
          </div>
        </div>
      </form>
    </Card>
  )
}

interface EditAnamnesisFormProps {
  anamnesis: Anamnesis
  focusScoreItemId?: string | null
  onSuccess: () => void
  onCancel: () => void
}

export function EditAnamnesisForm({ anamnesis, focusScoreItemId, onSuccess, onCancel }: EditAnamnesisFormProps) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>(
    anamnesis.anamnesisTemplateId || ''
  )
  const [selectedTemplate, setSelectedTemplate] = useState<AnamnesisTemplate | null>(null)
  const [consultationDate, setConsultationDate] = useState(
    format(new Date(anamnesis.consultationDate), "yyyy-MM-dd'T'HH:mm")
  )
  // Use HTML for editing (fallback to plain text for backward compatibility)
  const [summary, setSummary] = useState(anamnesis.summaryHtml || anamnesis.summary || '')
  const [content, setContent] = useState(anamnesis.contentHtml || anamnesis.content || '')
  const [visibility, setVisibility] = useState<'all' | 'medicalOnly' | 'psychOnly' | 'authorOnly'>(
    anamnesis.visibility
  )
  const [notes, setNotes] = useState(anamnesis.notes || '')

  // Initialize template items from existing anamnesis
  const [templateItems, setTemplateItems] = useState<AnamnesisItemFormValue[]>(
    anamnesis.items?.map(item => ({
      scoreItemId: item.scoreItemId,
      numericValue: item.numericValue,
      textValue: item.textValue,
      order: item.order,
    })) || []
  )

  // Rich text editor toggle - initialize based on whether HTML exists
  const [richTextEnabled, setRichTextEnabled] = useState(!!anamnesis.summaryHtml || !!anamnesis.contentHtml)

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Load template on mount if template is selected
  useEffect(() => {
    if (anamnesis.anamnesisTemplateId) {
      getAnamnesisTemplateById(anamnesis.anamnesisTemplateId, true)
        .then(setSelectedTemplate)
        .catch(() => toast.error('Erro ao carregar template'))
    }
  }, [anamnesis.anamnesisTemplateId])

  // Toggle rich text for both fields
  const toggleRichText = () => {
    if (richTextEnabled) {
      // Converting from rich text to plain text for both fields
      setSummary(htmlToPlainText(summary))
      if (content) {
        setContent(htmlToPlainText(content))
      }
    }
    setRichTextEnabled(!richTextEnabled)
  }

  // Fetch templates
  const { data: templates = [] } = useQuery({
    queryKey: ['anamnesis-templates'],
    queryFn: () => getAllAnamnesisTemplates(false),
  })

  const sortedTemplates = [...templates].sort((a, b) => a.name.localeCompare(b.name))

  // Update mutation
  const updateMutation = useMutation({
    mutationFn: (data: any) => updateAnamnesis(anamnesis.id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis'] })
      toast.success('Anamnese atualizada com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      if (error.response?.status === 403) {
        toast.error('Você não tem permissão para editar esta anamnese')
      } else {
        toast.error(error.response?.data?.error || 'Erro ao atualizar anamnese')
      }
    },
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSelectedTemplateId('')
      setSelectedTemplate(null)
      // Preserve existing items when deselecting template
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getAnamnesisTemplateById(templateId, true) // withItems = true
      setSelectedTemplate(template)

      // Update summary if empty
      if (!summary) {
        setSummary(`Anamnese - ${template.name}`)
      }

      // Preserve existing items and merge with new template
      // Items from new template go first (in order), then existing items not in template
      if (template.items) {
        const templateScoreItemIds = new Set(template.items.map(ti => ti.scoreItemId))
        const preservedItems = templateItems.filter(item => !templateScoreItemIds.has(item.scoreItemId))
        // Keep existing values, new template items will be empty initially
        setTemplateItems([...templateItems.filter(item => templateScoreItemIds.has(item.scoreItemId)), ...preservedItems])
      }
    } catch (error) {
      toast.error('Erro ao carregar template')
      setSelectedTemplate(null)
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Convert to RFC3339 format
    const dateObj = new Date(consultationDate)
    const rfc3339Date = dateObj.toISOString()

    // Prepare data based on whether rich text is enabled
    const summaryPlain = richTextEnabled ? htmlToPlainText(summary) : summary
    const summaryHtml = richTextEnabled ? summary.trim() : undefined
    const contentPlain = content ? (richTextEnabled ? htmlToPlainText(content) : content) : undefined
    const contentHtml = richTextEnabled ? content.trim() : undefined

    const payload = {
      consultationDate: rfc3339Date,
      summary: summaryPlain || undefined, // Plain text
      summaryHtml: summaryHtml || undefined, // HTML (only if rich text enabled)
      content: contentPlain || undefined, // Plain text
      contentHtml: contentHtml || undefined, // HTML (only if rich text enabled)
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

    updateMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-semibold">Editar Anamnese</h2>
        <Button
          type="button"
          variant={richTextEnabled ? "default" : "outline"}
          size="sm"
          onClick={toggleRichText}
          className="gap-2"
        >
          <Sparkles className="h-4 w-4" />
          {richTextEnabled ? 'Formatação Ativada' : 'Ativar Formatação'}
        </Button>
      </div>
      <form ref={formRef} onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Main Column - Form */}
          <div className="lg:col-span-2 space-y-4">
            <div>
              <Label>Data da Consulta *</Label>
              <Input
                type="datetime-local"
                value={consultationDate}
                onChange={(e) => setConsultationDate(e.target.value)}
                required
              />
            </div>

            <div>
              <Label>Resumo (opcional)</Label>
              {richTextEnabled ? (
                <RichTextEditor editorId="edit-summary"
                  value={summary}
                  onChange={setSummary}
                  placeholder="Resumo breve da anamnese..."
                  minHeight="120px"
                />
              ) : (
                <Textarea
                  value={summary}
                  onChange={(e) => setSummary(e.target.value)}
                  placeholder="Resumo breve da anamnese..."
                  rows={4}
                />
              )}
              <p className="text-xs text-muted-foreground mt-1">
                {richTextEnabled
                  ? 'Resumo formatado da anamnese'
                  : 'Resumo em texto simples'}
              </p>
            </div>

            <div>
              <Label>Conteúdo Detalhado (opcional)</Label>
              {richTextEnabled ? (
                <RichTextEditor editorId="edit-content"
                  value={content}
                  onChange={setContent}
                  placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                  minHeight="300px"
                />
              ) : (
                <Textarea
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                  rows={8}
                />
              )}
              <p className="text-xs text-muted-foreground mt-1">
                {richTextEnabled
                  ? 'Anamnese completa com formatação rica'
                  : 'Anamnese completa em texto simples'}
              </p>
            </div>

            <div>
              <Label>Visibilidade *</Label>
              <Select value={visibility} onValueChange={(v: any) => setVisibility(v)}>
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">Todos os Profissionais</SelectItem>
                  <SelectItem value="medicalOnly">Apenas Médicos</SelectItem>
                  <SelectItem value="psychOnly">Apenas Psicólogos</SelectItem>
                  <SelectItem value="authorOnly">Apenas Autor</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div>
              <Label>Observações (opcional)</Label>
              <Textarea
                value={notes}
                onChange={(e) => setNotes(e.target.value)}
                placeholder="Observações adicionais..."
                rows={3}
              />
            </div>

            {/* Template Items Form */}
            {selectedTemplate && selectedTemplate.items && selectedTemplate.items.length > 0 && (
              <div className="pt-4 border-t">
                <h3 className="text-lg font-semibold mb-4">
                  Items do Template: {selectedTemplate.name}
                </h3>
                <AnamnesisTemplateItemsForm
                  key={selectedTemplate.id}
                  template={selectedTemplate}
                  initialValues={templateItems}
                  onChange={setTemplateItems}
                  focusScoreItemId={focusScoreItemId}
                />
              </div>
            )}

            <div className="flex justify-end gap-2 pt-4 border-t">
              <Button type="button" variant="outline" onClick={onCancel}>
                Cancelar
              </Button>
              <Button type="submit" disabled={updateMutation.isPending}>
                {updateMutation.isPending ? 'Salvando...' : 'Salvar Alterações'}
              </Button>
            </div>
          </div>

          {/* Sidebar - Templates */}
          <div className="lg:col-span-1">
            <Label className="mb-3 block">Templates (opcional)</Label>
            <ScrollArea className="h-[600px] pr-4">
              <div className="space-y-2">
                <div
                  onClick={() => handleTemplateSelect('')}
                  className={cn(
                    'relative p-3 rounded-lg cursor-pointer transition-all',
                    'border-2',
                    selectedTemplateId === ''
                      ? 'border-primary bg-primary/5'
                      : 'border-border hover:border-primary/50 hover:bg-accent'
                  )}
                >
                  <div className="flex items-center justify-between">
                    <span className="font-medium text-sm">Nenhum template</span>
                    {selectedTemplateId === '' && <Check className="h-4 w-4 text-primary" />}
                  </div>
                </div>
                {sortedTemplates.map((template) => (
                  <div
                    key={template.id}
                    onClick={() => handleTemplateSelect(template.id)}
                    className={cn(
                      'relative p-3 rounded-lg cursor-pointer transition-all',
                      'border-2',
                      selectedTemplateId === template.id
                        ? 'border-primary bg-primary/5'
                        : 'border-border hover:border-primary/50 hover:bg-accent'
                    )}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium text-sm">{template.name}</div>
                        <div className="text-xs text-muted-foreground mt-1">{template.area}</div>
                      </div>
                      {selectedTemplateId === template.id && (
                        <Check className="h-4 w-4 text-primary flex-shrink-0 mt-0.5" />
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
            {selectedTemplateId && (
              <p className="text-xs text-muted-foreground mt-3">
                Template selecionado aplicado ao resumo
              </p>
            )}
          </div>
        </div>
      </form>
    </Card>
  )
}
