'use client'

import { useState, useRef, useEffect } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { usePatientGuard } from '@/lib/use-patient-guard'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Card } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Badge } from '@/components/ui/badge'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Plus, FileText, Calendar, Pencil, Check, Copy, Shield } from 'lucide-react'
import { ScrollArea } from '@/components/ui/scroll-area'
import { cn } from '@/lib/utils'
import {
  getLabRequestsByPatientId,
  createLabRequest,
  updateLabRequest,
  generateLabRequestPdf,
  openLabRequestPdf,
  importLabRequestExams,
  type LabRequest
} from '@/lib/api/lab-requests'
import {
  getAllLabRequestTemplates,
  getLabRequestTemplateById,
  type LabRequestTemplate
} from '@/lib/api/lab-request-templates'
import { applyTemplate, normalizeName } from '@/lib/lab-request-apply'
import { toast } from 'sonner'
import { format, parseISO } from 'date-fns'
import { PageHeader } from '@/components/layout/page-header'

export default function LabRequestsPage() {
  usePatientGuard(); // Restrict access to staff only
  const router = useRouter()
  const { selectedPatient, isLoading: patientLoading } = useRequireSelectedPatient()
  const [showCreateForm, setShowCreateForm] = useState(false)
  const [editingRequest, setEditingRequest] = useState<LabRequest | null>(null)
  const [duplicatingRequest, setDuplicatingRequest] = useState<LabRequest | null>(null)

  // Pedidos SÓ do paciente selecionado (isolamento de prontuário — diretriz do EMR: nunca
  // listar dados de todos os pacientes). Endpoint dedicado por paciente, não o /lab-requests global.
  const { data: requestsData, isLoading: requestsLoading } = useQuery({
    queryKey: ['lab-requests', selectedPatient?.id],
    queryFn: () => getLabRequestsByPatientId(selectedPatient!.id),
    enabled: !!selectedPatient // Só busca quando tiver paciente selecionado
  })

  const requests = requestsData || []

  if (patientLoading) {
    return <div className="container mx-auto py-8">Carregando...</div>
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      <SelectedPatientHeader />

      <PageHeader
        breadcrumbs={[{ label: 'Pedidos de Exames' }]}
        title="Pedidos de Exames"
        description={`${requests.length} pedido${requests.length !== 1 ? 's' : ''} de exames`}
        actions={[
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setShowCreateForm(!showCreateForm),
            variant: 'default',
          },
        ]}
      />

      {showCreateForm && (
        <CreateLabRequestForm onSuccess={() => setShowCreateForm(false)} />
      )}

      {editingRequest && (
        <EditLabRequestForm
          request={editingRequest}
          onSuccess={() => setEditingRequest(null)}
          onCancel={() => setEditingRequest(null)}
        />
      )}

      {duplicatingRequest && (
        <DuplicateLabRequestForm
          request={duplicatingRequest}
          onSuccess={() => setDuplicatingRequest(null)}
          onCancel={() => setDuplicatingRequest(null)}
        />
      )}

      {requestsLoading ? (
        <div className="text-center py-12">Carregando pedidos...</div>
      ) : requests.length === 0 ? (
        <Card className="p-12 text-center">
          <FileText className="h-12 w-12 mx-auto text-muted-foreground mb-4" />
          <p className="text-muted-foreground mb-4">Nenhum pedido de exame cadastrado ainda</p>
          <Button onClick={() => setShowCreateForm(true)}>
            <Plus className="h-4 w-4 mr-2" />
            Criar Primeiro Pedido
          </Button>
        </Card>
      ) : (
        <div className="grid gap-4">
          {requests.map((request) => (
            <LabRequestCard
              key={request.id}
              request={request}
              onEdit={() => setEditingRequest(request)}
              onDuplicate={() => setDuplicatingRequest(request)}
            />
          ))}
        </div>
      )}
    </div>
  )
}

function CreateLabRequestForm({ onSuccess }: { onSuccess: () => void }) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>('')
  const [date, setDate] = useState(format(new Date(), 'yyyy-MM-dd'))
  const [exams, setExams] = useState('')
  const [notes, setNotes] = useState('')
  const [coveredCodes, setCoveredCodes] = useState<Set<string>>(new Set())
  const [importing, setImporting] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)
  const { selectedPatient } = useRequireSelectedPatient()

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates (sorted alphabetically)
  const { data: templates = [] } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  const sortedTemplates = [...templates].sort(
    (a, b) => (a.displayOrder ?? 0) - (b.displayOrder ?? 0) || a.name.localeCompare(b.name)
  )

  // Create mutation
  const createMutation = useMutation({
    mutationFn: createLabRequest,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-requests'] })
      toast.success('Pedido criado com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao criar pedido')
    }
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setExams('')
      setSelectedTemplateId('')
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getLabRequestTemplateById(templateId)
      if (template.labTests && template.labTests.length > 0) {
        // Filtra por sexo do paciente (ex.: PSA só ♂) e remove os já cobertos por pedido externo.
        const examsText = applyTemplate(template.labTests, selectedPatient?.gender, coveredCodes)
        setExams(examsText)
      } else {
        toast.info('Este template não tem exames cadastrados ainda')
      }
    } catch (error) {
      toast.error('Erro ao carregar template')
    }
  }

  // Importa um pedido externo (foto/PDF), reconhece os exames e remove do pedido os já solicitados.
  const handleImportFile = async (file: File) => {
    setImporting(true)
    try {
      const { items } = await importLabRequestExams(file)
      const matched = items.filter((i) => i.matched && i.code)
      if (matched.length === 0) {
        toast.info('Nenhum exame do nosso catálogo foi reconhecido no pedido importado')
        return
      }
      const codes = matched.map((i) => i.code as string)
      const names = new Set(matched.map((i) => normalizeName(i.name || '')))
      const prevExams = exams
      // Remove os exames casados E as linhas de justificativa (#) que pertencem a eles.
      let droppingJust = false // dentro das linhas # de um exame que acabou de ser removido?
      const kept = exams
        .split('\n')
        .filter((line) => {
          const t = line.trim()
          if (t.startsWith('#')) return !droppingJust // justificativa segue o destino do exame de cima
          if (t === '') {
            droppingJust = false // linha em branco = novo bloco/página
            return true
          }
          if (names.has(normalizeName(line))) {
            droppingJust = true // exame removido → suas justificativas saem junto
            return false
          }
          droppingJust = false
          return true
        })
        .join('\n')
      setExams(kept)
      setCoveredCodes((prev) => new Set([...prev, ...codes]))
      toast.success(
        `${matched.length} exame(s) já solicitado(s) removido(s): ${matched.map((i) => i.name).join(', ')}`,
        {
          action: {
            label: 'Desfazer',
            onClick: () => {
              setExams(prevExams)
              setCoveredCodes((prev) => {
                const n = new Set(prev)
                codes.forEach((c) => n.delete(c))
                return n
              })
            },
          },
        }
      )
    } catch (e: any) {
      toast.error('Erro ao importar pedido: ' + (e?.message || ''))
    } finally {
      setImporting(false)
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    if (!exams.trim()) {
      toast.error('Adicione pelo menos um exame')
      return
    }

    const payload = {
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined,
      labRequestTemplateId: selectedTemplateId || undefined
    }

    // Backend pega patientId automaticamente do selectedPatient no JWT
    createMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <div className="flex items-center justify-between mb-4 gap-3">
        <h2 className="text-xl font-semibold">Novo Pedido de Exames</h2>
        <input
          ref={fileInputRef}
          type="file"
          accept="application/pdf,image/*"
          className="hidden"
          onChange={(e) => {
            const f = e.target.files?.[0]
            if (f) handleImportFile(f)
            e.currentTarget.value = ''
          }}
        />
        <Button
          type="button"
          variant="outline"
          size="sm"
          disabled={importing}
          onClick={() => fileInputRef.current?.click()}
          title="Importe um pedido de outro médico (foto/PDF); os exames já solicitados serão removidos do seu pedido"
        >
          {importing ? 'Importando…' : 'Importar pedido existente'}
        </Button>
      </div>
      <form ref={formRef} onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Coluna Principal - Formulário */}
          <div className="lg:col-span-2 space-y-4">
            <div>
              <Label>Data *</Label>
              <Input
                type="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                required
              />
            </div>

            <div>
              <Label>Exames Solicitados * (um por linha)</Label>
              <Textarea
                value={exams}
                onChange={(e) => setExams(e.target.value)}
                placeholder="Digite os nomes dos exames, um por linha. Comece a linha com # para justificar o exame acima (aparece no PDF sob o exame). Exemplo:&#10;Ressonância de crânio&#10;# Sinais de demência, investigação diagnóstica&#10;Hemograma Completo"
                rows={16}
                required
              />
              <p className="text-xs text-muted-foreground mt-1">
                {exams.split('\n').filter(line => line.trim() && !line.trim().startsWith('#')).length} exame(s)
              </p>
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
                {createMutation.isPending ? 'Criando...' : 'Criar Pedido'}
              </Button>
            </div>
          </div>

          {/* Coluna Lateral - Templates */}
          <div className="lg:col-span-1">
            <Label className="mb-3 block">Templates (opcional)</Label>
            <ScrollArea className="h-[600px] pr-4">
              <div className="space-y-2">
                <div
                  onClick={() => handleTemplateSelect('')}
                  className={cn(
                    "relative p-3 rounded-lg cursor-pointer transition-all",
                    "border-2",
                    selectedTemplateId === ''
                      ? "border-primary bg-primary/5"
                      : "border-border hover:border-primary/50 hover:bg-accent"
                  )}
                >
                  <div className="flex items-center justify-between">
                    <span className="font-medium text-sm">Nenhum template</span>
                    {selectedTemplateId === '' && (
                      <Check className="h-4 w-4 text-primary" />
                    )}
                  </div>
                </div>
                {sortedTemplates.map((template) => (
                  <div
                    key={template.id}
                    onClick={() => handleTemplateSelect(template.id)}
                    className={cn(
                      "relative p-3 rounded-lg cursor-pointer transition-all",
                      "border-2",
                      selectedTemplateId === template.id
                        ? "border-primary bg-primary/5"
                        : "border-border hover:border-primary/50 hover:bg-accent"
                    )}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium text-sm">{template.name}</div>
                        {template.description && (
                          <div className="text-xs text-muted-foreground mt-1">
                            {template.description}
                          </div>
                        )}
                      </div>
                      {selectedTemplateId === template.id && (
                        <Check className="h-4 w-4 text-primary shrink-0 mt-0.5" />
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
            {selectedTemplateId && (
              <p className="text-xs text-muted-foreground mt-3">
                Os exames do template foram inseridos acima em ordem alfabética
              </p>
            )}
          </div>
        </div>
      </form>
    </Card>
  )
}

function EditLabRequestForm({
  request,
  onSuccess,
  onCancel
}: {
  request: LabRequest
  onSuccess: () => void
  onCancel: () => void
}) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>(
    request.labRequestTemplateId || ''
  )
  // Extrair apenas a data (antes do T) para evitar problemas de timezone
  const [date, setDate] = useState(request.date.split('T')[0])
  const [exams, setExams] = useState(request.exams)
  const [notes, setNotes] = useState(request.notes || '')
  const { selectedPatient } = useRequireSelectedPatient()

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates (sorted alphabetically)
  const { data: templates = [] } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  const sortedTemplates = [...templates].sort(
    (a, b) => (a.displayOrder ?? 0) - (b.displayOrder ?? 0) || a.name.localeCompare(b.name)
  )

  // Update mutation
  const updateMutation = useMutation({
    mutationFn: (data: { date: string; exams: string; notes?: string }) =>
      updateLabRequest(request.id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-requests'] })
      toast.success('Pedido atualizado com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao atualizar pedido')
    }
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSelectedTemplateId('')
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getLabRequestTemplateById(templateId)
      if (template.labTests && template.labTests.length > 0) {
        // Filtra por sexo do paciente ao carregar o template (PSA só ♂, CA-125 só ♀, etc.).
        const examsText = applyTemplate(template.labTests, selectedPatient?.gender, new Set<string>())
        setExams(examsText)
      } else {
        toast.info('Este template não tem exames cadastrados ainda')
      }
    } catch (error) {
      toast.error('Erro ao carregar template')
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    if (!exams.trim()) {
      toast.error('Adicione pelo menos um exame')
      return
    }

    const payload = {
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined,
      labRequestTemplateId: selectedTemplateId || undefined
    }

    updateMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6">
      <h2 className="text-xl font-semibold mb-4">Editar Pedido de Exames</h2>
      <form ref={formRef} onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Coluna Principal - Formulário */}
          <div className="lg:col-span-2 space-y-4">
            <div>
              <Label>Data *</Label>
              <Input
                type="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                required
              />
            </div>

            <div>
              <Label>Exames Solicitados * (um por linha)</Label>
              <Textarea
                value={exams}
                onChange={(e) => setExams(e.target.value)}
                placeholder="Digite os nomes dos exames, um por linha. Comece a linha com # para justificar o exame acima (aparece no PDF sob o exame). Exemplo:&#10;Ressonância de crânio&#10;# Sinais de demência, investigação diagnóstica&#10;Hemograma Completo"
                rows={16}
                required
              />
              <p className="text-xs text-muted-foreground mt-1">
                {exams.split('\n').filter(line => line.trim() && !line.trim().startsWith('#')).length} exame(s)
              </p>
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

          {/* Coluna Lateral - Templates */}
          <div className="lg:col-span-1">
            <Label className="mb-3 block">Templates (opcional)</Label>
            <ScrollArea className="h-[600px] pr-4">
              <div className="space-y-2">
                <div
                  onClick={() => handleTemplateSelect('')}
                  className={cn(
                    "relative p-3 rounded-lg cursor-pointer transition-all",
                    "border-2",
                    selectedTemplateId === ''
                      ? "border-primary bg-primary/5"
                      : "border-border hover:border-primary/50 hover:bg-accent"
                  )}
                >
                  <div className="flex items-center justify-between">
                    <span className="font-medium text-sm">Nenhum template</span>
                    {selectedTemplateId === '' && (
                      <Check className="h-4 w-4 text-primary" />
                    )}
                  </div>
                </div>
                {sortedTemplates.map((template) => (
                  <div
                    key={template.id}
                    onClick={() => handleTemplateSelect(template.id)}
                    className={cn(
                      "relative p-3 rounded-lg cursor-pointer transition-all",
                      "border-2",
                      selectedTemplateId === template.id
                        ? "border-primary bg-primary/5"
                        : "border-border hover:border-primary/50 hover:bg-accent"
                    )}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium text-sm">{template.name}</div>
                        {template.description && (
                          <div className="text-xs text-muted-foreground mt-1">
                            {template.description}
                          </div>
                        )}
                      </div>
                      {selectedTemplateId === template.id && (
                        <Check className="h-4 w-4 text-primary shrink-0 mt-0.5" />
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
            {selectedTemplateId && (
              <p className="text-xs text-muted-foreground mt-3">
                Os exames do template foram inseridos acima em ordem alfabética
              </p>
            )}
          </div>
        </div>
      </form>
    </Card>
  )
}

function DuplicateLabRequestForm({
  request,
  onSuccess,
  onCancel
}: {
  request: LabRequest
  onSuccess: () => void
  onCancel: () => void
}) {
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>(
    request.labRequestTemplateId || ''
  )
  // Always use today's date when form is shown (not when component mounts)
  const [date, setDate] = useState(() => format(new Date(), 'yyyy-MM-dd'))
  const [exams, setExams] = useState(request.exams)
  const [notes, setNotes] = useState(request.notes || '')
  const { selectedPatient } = useRequireSelectedPatient()

  // Reset date to today whenever the form is shown
  useEffect(() => {
    setDate(format(new Date(), 'yyyy-MM-dd'))
  }, [request.id])

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates (sorted alphabetically)
  const { data: templates = [] } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  const sortedTemplates = [...templates].sort(
    (a, b) => (a.displayOrder ?? 0) - (b.displayOrder ?? 0) || a.name.localeCompare(b.name)
  )

  // Create mutation (duplicate is a create operation)
  const createMutation = useMutation({
    mutationFn: createLabRequest,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-requests'] })
      toast.success('Pedido duplicado com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao duplicar pedido')
    }
  })

  // Handle template selection
  const handleTemplateSelect = async (templateId: string) => {
    if (!templateId) {
      setSelectedTemplateId('')
      return
    }

    setSelectedTemplateId(templateId)

    try {
      const template = await getLabRequestTemplateById(templateId)
      if (template.labTests && template.labTests.length > 0) {
        // Filtra por sexo do paciente ao carregar o template (PSA só ♂, CA-125 só ♀, etc.).
        const examsText = applyTemplate(template.labTests, selectedPatient?.gender, new Set<string>())
        setExams(examsText)
      } else {
        toast.info('Este template não tem exames cadastrados ainda')
      }
    } catch (error) {
      toast.error('Erro ao carregar template')
    }
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    if (!exams.trim()) {
      toast.error('Adicione pelo menos um exame')
      return
    }

    const payload = {
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined,
      labRequestTemplateId: selectedTemplateId || undefined
    }

    createMutation.mutate(payload)
  }

  return (
    <Card className="p-6 mb-6 border-ocean-200 bg-ocean-50/50">
      <h2 className="text-xl font-semibold mb-4 flex items-center gap-2">
        <Copy className="h-5 w-5 text-ocean-700" />
        Duplicar Pedido de Exames
      </h2>
      <form ref={formRef} onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Coluna Principal - Formulário */}
          <div className="lg:col-span-2 space-y-4">
            <div>
              <Label>Data *</Label>
              <Input
                type="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                required
              />
            </div>

            <div>
              <Label>Exames Solicitados * (um por linha)</Label>
              <Textarea
                value={exams}
                onChange={(e) => setExams(e.target.value)}
                placeholder="Digite os nomes dos exames, um por linha. Comece a linha com # para justificar o exame acima (aparece no PDF sob o exame). Exemplo:&#10;Ressonância de crânio&#10;# Sinais de demência, investigação diagnóstica&#10;Hemograma Completo"
                rows={16}
                required
              />
              <p className="text-xs text-muted-foreground mt-1">
                {exams.split('\n').filter(line => line.trim() && !line.trim().startsWith('#')).length} exame(s)
              </p>
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
              <Button type="submit" disabled={createMutation.isPending}>
                {createMutation.isPending ? 'Duplicando...' : 'Criar Cópia'}
              </Button>
            </div>
          </div>

          {/* Coluna Lateral - Templates */}
          <div className="lg:col-span-1">
            <Label className="mb-3 block">Templates (opcional)</Label>
            <ScrollArea className="h-[600px] pr-4">
              <div className="space-y-2">
                <div
                  onClick={() => handleTemplateSelect('')}
                  className={cn(
                    "relative p-3 rounded-lg cursor-pointer transition-all",
                    "border-2",
                    selectedTemplateId === ''
                      ? "border-primary bg-primary/5"
                      : "border-border hover:border-primary/50 hover:bg-accent"
                  )}
                >
                  <div className="flex items-center justify-between">
                    <span className="font-medium text-sm">Nenhum template</span>
                    {selectedTemplateId === '' && (
                      <Check className="h-4 w-4 text-primary" />
                    )}
                  </div>
                </div>
                {sortedTemplates.map((template) => (
                  <div
                    key={template.id}
                    onClick={() => handleTemplateSelect(template.id)}
                    className={cn(
                      "relative p-3 rounded-lg cursor-pointer transition-all",
                      "border-2",
                      selectedTemplateId === template.id
                        ? "border-primary bg-primary/5"
                        : "border-border hover:border-primary/50 hover:bg-accent"
                    )}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium text-sm">{template.name}</div>
                        {template.description && (
                          <div className="text-xs text-muted-foreground mt-1">
                            {template.description}
                          </div>
                        )}
                      </div>
                      {selectedTemplateId === template.id && (
                        <Check className="h-4 w-4 text-primary shrink-0 mt-0.5" />
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
            {selectedTemplateId && (
              <p className="text-xs text-muted-foreground mt-3">
                Os exames do template foram inseridos acima em ordem alfabética
              </p>
            )}
          </div>
        </div>
      </form>
    </Card>
  )
}

function LabRequestCard({
  request,
  onEdit,
  onDuplicate
}: {
  request: LabRequest
  onEdit: () => void
  onDuplicate: () => void
}) {
  const queryClient = useQueryClient()

  // Parse exams into array (ignora linhas de justificativa, que começam com #)
  const examsArray = request.exams
    .split('\n')
    .map(line => line.trim())
    .filter(line => line && !line.startsWith('#'))
  const examsCount = examsArray.length

  const generatePdfMutation = useMutation({
    mutationFn: () => generateLabRequestPdf(request.id),
    onSuccess: async (data) => {
      queryClient.invalidateQueries({ queryKey: ['lab-requests'] })
      toast.success('PDF gerado com sucesso!')

      // Abre o PDF recém-gerado via download autenticado (rota estática /uploads foi removida)
      if (data.pdfUrl) {
        try {
          await openLabRequestPdf(request.id)
        } catch {
          toast.error('PDF gerado, mas não foi possível abrir. Tente o botão novamente.')
        }
      }
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao gerar PDF')
    }
  })

  const handleGeneratePdf = async () => {
    if (request.pdfUrl) {
      // PDF já existe: abrir via download autenticado
      try {
        await openLabRequestPdf(request.id)
      } catch {
        toast.error('Não foi possível abrir o PDF')
      }
    } else {
      // Gerar PDF
      generatePdfMutation.mutate()
    }
  }

  return (
    <Card className="p-4">
      <div className="flex items-start justify-between gap-4">
        <div className="flex-1 min-w-0">
          <div className="flex items-center gap-3 mb-3">
            <Calendar className="h-4 w-4 text-muted-foreground shrink-0" />
            <span className="text-sm font-medium">
              {(() => {
                // Extrair apenas a data (antes do T) para evitar problemas de timezone
                const dateOnly = request.date.split('T')[0]
                return format(parseISO(dateOnly), 'dd/MM/yyyy')
              })()}
            </span>
            {request.pdfUrl && (
              <div className="flex gap-2">
                <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-md text-xs font-medium bg-green-100 text-green-800">
                  <Check className="h-3 w-3" />
                  PDF Gerado
                </span>
                {request.signedAt && (
                  <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-md text-xs font-medium bg-blue-100 text-blue-800">
                    <Shield className="h-3 w-3" />
                    Assinado Digitalmente
                  </span>
                )}
              </div>
            )}
          </div>

          <div className="space-y-2">
            <div>
              <span className="text-sm font-medium text-foreground mb-2 block">
                {examsCount} exame{examsCount !== 1 ? 's' : ''}:
              </span>
              <div className="flex flex-wrap gap-1.5">
                {examsArray.map((exam, index) => (
                  <Badge
                    key={index}
                    variant="secondary"
                    className="bg-muted text-muted-foreground hover:bg-muted/80 font-normal"
                  >
                    {exam}
                  </Badge>
                ))}
              </div>
            </div>

            {request.notes && (
              <p className="text-sm text-muted-foreground italic border-l-2 border-muted pl-3">
                {request.notes}
              </p>
            )}
          </div>
        </div>

        <div className="flex gap-2 shrink-0">
          {!request.pdfUrl && (
            <Button
              variant="outline"
              size="sm"
              onClick={onEdit}
              title="Editar pedido"
            >
              <Pencil className="h-4 w-4" />
            </Button>
          )}
          <Button
            variant="outline"
            size="sm"
            onClick={onDuplicate}
            title="Duplicar pedido"
          >
            <Copy className="h-4 w-4" />
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={handleGeneratePdf}
            disabled={generatePdfMutation.isPending}
            title={request.pdfUrl ? 'Visualizar PDF' : 'Gerar PDF'}
          >
            <FileText className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </Card>
  )
}
