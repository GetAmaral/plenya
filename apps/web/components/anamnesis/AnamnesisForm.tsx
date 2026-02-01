'use client'

import { useState, useRef } from 'react'
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
import { Check } from 'lucide-react'
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

interface CreateAnamnesisFormProps {
  onSuccess: () => void
}

export function CreateAnamnesisForm({ onSuccess }: CreateAnamnesisFormProps) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>('')
  const [consultationDate, setConsultationDate] = useState(
    format(new Date(), "yyyy-MM-dd'T'HH:mm")
  )
  const [summary, setSummary] = useState('')
  const [content, setContent] = useState('')
  const [visibility, setVisibility] = useState<'all' | 'medicalOnly' | 'psychOnly'>('all')
  const [notes, setNotes] = useState('')

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
      toast.error(error.response?.data?.error || 'Erro ao criar anamnese')
    },
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSummary('')
      setContent('')
      setSelectedTemplateId('')
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getAnamnesisTemplateById(templateId)
      // Set summary with template name
      if (!summary) {
        setSummary(`Anamnese - ${template.name}`)
      }
      // You could also pre-fill content based on template items if needed
    } catch (error) {
      toast.error('Erro ao carregar template')
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Validate that there's actual text content
    if (!hasTextContent(summary)) {
      toast.error('O resumo é obrigatório')
      return
    }

    // Convert to RFC3339 format
    const dateObj = new Date(consultationDate)
    const rfc3339Date = dateObj.toISOString()

    // Extract plain text from HTML for search/indexing
    const summaryPlain = htmlToPlainText(summary)
    const contentPlain = content ? htmlToPlainText(content) : undefined

    const payload: CreateAnamnesisRequest = {
      consultationDate: rfc3339Date,
      summary: summaryPlain || undefined, // Plain text
      summaryHtml: summary.trim() || undefined, // HTML
      content: contentPlain || undefined, // Plain text
      contentHtml: content.trim() || undefined, // HTML
      visibility,
      notes: notes.trim() || undefined,
      anamnesisTemplateId: selectedTemplateId || undefined,
    }

    console.log('Creating anamnesis with:', payload)

    // Backend gets patientId automatically from selectedPatient in JWT
    createMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <h2 className="text-xl font-semibold mb-4">Nova Anamnese</h2>
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
              <Label>Resumo *</Label>
              <RichTextEditor
                value={summary}
                onChange={setSummary}
                placeholder="Resumo breve da anamnese..."
                minHeight="120px"
              />
              <p className="text-xs text-muted-foreground mt-1">
                Resumo formatado da anamnese (obrigatório)
              </p>
            </div>

            <div>
              <Label>Conteúdo Detalhado (opcional)</Label>
              <RichTextEditor
                value={content}
                onChange={setContent}
                placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                minHeight="300px"
              />
              <p className="text-xs text-muted-foreground mt-1">
                Anamnese completa com formatação rica
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
  onSuccess: () => void
  onCancel: () => void
}

export function EditAnamnesisForm({ anamnesis, onSuccess, onCancel }: EditAnamnesisFormProps) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>(
    anamnesis.anamnesisTemplateId || ''
  )
  const [consultationDate, setConsultationDate] = useState(
    format(new Date(anamnesis.consultationDate), "yyyy-MM-dd'T'HH:mm")
  )
  // Use HTML for editing (fallback to plain text for backward compatibility)
  const [summary, setSummary] = useState(anamnesis.summaryHtml || anamnesis.summary || '')
  const [content, setContent] = useState(anamnesis.contentHtml || anamnesis.content || '')
  const [visibility, setVisibility] = useState<'all' | 'medicalOnly' | 'psychOnly'>(
    anamnesis.visibility
  )
  const [notes, setNotes] = useState(anamnesis.notes || '')

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

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
      toast.error(error.response?.data?.error || 'Erro ao atualizar anamnese')
    },
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSelectedTemplateId('')
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getAnamnesisTemplateById(templateId)
      // Update summary if empty
      if (!summary) {
        setSummary(`Anamnese - ${template.name}`)
      }
    } catch (error) {
      toast.error('Erro ao carregar template')
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Validate that there's actual text content
    if (!hasTextContent(summary)) {
      toast.error('O resumo é obrigatório')
      return
    }

    // Convert to RFC3339 format
    const dateObj = new Date(consultationDate)
    const rfc3339Date = dateObj.toISOString()

    // Extract plain text from HTML for search/indexing
    const summaryPlain = htmlToPlainText(summary)
    const contentPlain = content ? htmlToPlainText(content) : undefined

    const payload = {
      consultationDate: rfc3339Date,
      summary: summaryPlain || undefined, // Plain text
      summaryHtml: summary.trim() || undefined, // HTML
      content: contentPlain || undefined, // Plain text
      contentHtml: content.trim() || undefined, // HTML
      visibility,
      notes: notes.trim() || undefined,
      anamnesisTemplateId: selectedTemplateId || undefined,
    }

    console.log('Updating anamnesis with:', payload)

    updateMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <h2 className="text-xl font-semibold mb-4">Editar Anamnese</h2>
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
              <Label>Resumo *</Label>
              <RichTextEditor
                value={summary}
                onChange={setSummary}
                placeholder="Resumo breve da anamnese..."
                minHeight="120px"
              />
              <p className="text-xs text-muted-foreground mt-1">
                Resumo formatado da anamnese (obrigatório)
              </p>
            </div>

            <div>
              <Label>Conteúdo Detalhado (opcional)</Label>
              <RichTextEditor
                value={content}
                onChange={setContent}
                placeholder="Conteúdo detalhado da anamnese com formatação completa..."
                minHeight="300px"
              />
              <p className="text-xs text-muted-foreground mt-1">
                Anamnese completa com formatação rica
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
