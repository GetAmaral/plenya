'use client'

import { useState, useRef } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
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
import { Plus, FileText, Calendar, Pencil } from 'lucide-react'
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  getAllLabRequests,
  createLabRequest,
  updateLabRequest,
  type LabRequest
} from '@/lib/api/lab-requests'
import {
  getAllLabRequestTemplates,
  getLabRequestTemplateById,
  type LabRequestTemplate
} from '@/lib/api/lab-request-templates'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { PageHeader } from '@/components/layout/page-header'

export default function LabRequestsPage() {
  const router = useRouter()
  const { selectedPatient, isLoading: patientLoading } = useRequireSelectedPatient()
  const [showCreateForm, setShowCreateForm] = useState(false)
  const [editingRequest, setEditingRequest] = useState<LabRequest | null>(null)

  // Fetch lab requests (backend filtra automaticamente pelo selectedPatient)
  const { data: requestsData, isLoading: requestsLoading } = useQuery({
    queryKey: ['lab-requests', selectedPatient?.id],
    queryFn: () => getAllLabRequests({ limit: 50, offset: 0 }),
    enabled: !!selectedPatient // Só busca quando tiver paciente selecionado
  })

  const requests = requestsData?.data || []

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

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates (sorted alphabetically)
  const { data: templates = [] } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  const sortedTemplates = [...templates].sort((a, b) => a.name.localeCompare(b.name))

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
        // Sort tests alphabetically by name
        const sortedTests = [...template.labTests].sort((a, b) =>
          a.name.localeCompare(b.name)
        )

        // Create text with one exam per line
        const examsText = sortedTests.map(test => test.name).join('\n')
        setExams(examsText)
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

    // Backend pega patientId automaticamente do selectedPatient no JWT
    createMutation.mutate({
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined
    })
  }

  return (
    <Card className="p-6 mb-6">
      <h2 className="text-xl font-semibold mb-4">Novo Pedido de Exames</h2>
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
                placeholder="Digite os nomes dos exames, um por linha. Exemplo:&#10;Hemograma Completo&#10;Glicemia de Jejum&#10;Colesterol Total"
                rows={16}
                required
              />
              <p className="text-xs text-muted-foreground mt-1">
                {exams.split('\n').filter(line => line.trim()).length} exame(s)
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
              <RadioGroup value={selectedTemplateId} onValueChange={handleTemplateSelect}>
                <div className="space-y-2">
                  <div className="flex items-center space-x-2 p-3 rounded-md hover:bg-accent cursor-pointer border border-transparent hover:border-border">
                    <RadioGroupItem value="" id="no-template" />
                    <Label
                      htmlFor="no-template"
                      className="flex-1 cursor-pointer font-normal"
                    >
                      Nenhum template
                    </Label>
                  </div>
                  {sortedTemplates.map((template) => (
                    <div
                      key={template.id}
                      className="flex items-center space-x-2 p-3 rounded-md hover:bg-accent cursor-pointer border border-transparent hover:border-border"
                    >
                      <RadioGroupItem value={template.id} id={`template-${template.id}`} />
                      <Label
                        htmlFor={`template-${template.id}`}
                        className="flex-1 cursor-pointer font-normal"
                      >
                        <div>
                          <div className="font-medium">{template.name}</div>
                          {template.description && (
                            <div className="text-xs text-muted-foreground mt-1">
                              {template.description}
                            </div>
                          )}
                        </div>
                      </Label>
                    </div>
                  ))}
                </div>
              </RadioGroup>
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
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>('')
  const [date, setDate] = useState(format(new Date(request.date), 'yyyy-MM-dd'))
  const [exams, setExams] = useState(request.exams)
  const [notes, setNotes] = useState(request.notes || '')

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const queryClient = useQueryClient()

  // Fetch templates (sorted alphabetically)
  const { data: templates = [] } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  const sortedTemplates = [...templates].sort((a, b) => a.name.localeCompare(b.name))

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
        const sortedTests = [...template.labTests].sort((a, b) =>
          a.name.localeCompare(b.name)
        )
        const examsText = sortedTests.map(test => test.name).join('\n')
        setExams(examsText)
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

    updateMutation.mutate({
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined
    })
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
                placeholder="Digite os nomes dos exames, um por linha. Exemplo:&#10;Hemograma Completo&#10;Glicemia de Jejum&#10;Colesterol Total"
                rows={16}
                required
              />
              <p className="text-xs text-muted-foreground mt-1">
                {exams.split('\n').filter(line => line.trim()).length} exame(s)
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
              <RadioGroup value={selectedTemplateId} onValueChange={handleTemplateSelect}>
                <div className="space-y-2">
                  <div className="flex items-center space-x-2 p-3 rounded-md hover:bg-accent cursor-pointer border border-transparent hover:border-border">
                    <RadioGroupItem value="" id="edit-no-template" />
                    <Label
                      htmlFor="edit-no-template"
                      className="flex-1 cursor-pointer font-normal"
                    >
                      Nenhum template
                    </Label>
                  </div>
                  {sortedTemplates.map((template) => (
                    <div
                      key={template.id}
                      className="flex items-center space-x-2 p-3 rounded-md hover:bg-accent cursor-pointer border border-transparent hover:border-border"
                    >
                      <RadioGroupItem value={template.id} id={`edit-template-${template.id}`} />
                      <Label
                        htmlFor={`edit-template-${template.id}`}
                        className="flex-1 cursor-pointer font-normal"
                      >
                        <div>
                          <div className="font-medium">{template.name}</div>
                          {template.description && (
                            <div className="text-xs text-muted-foreground mt-1">
                              {template.description}
                            </div>
                          )}
                        </div>
                      </Label>
                    </div>
                  ))}
                </div>
              </RadioGroup>
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

function LabRequestCard({ request, onEdit }: { request: LabRequest; onEdit: () => void }) {
  const examsCount = request.exams.split('\n').filter(line => line.trim()).length

  return (
    <Card className="p-4">
      <div className="flex items-start justify-between">
        <div className="flex-1">
          <div className="flex items-center gap-3 mb-2">
            <Calendar className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm text-muted-foreground">
              {format(new Date(request.date), 'dd/MM/yyyy')}
            </span>
          </div>

          <h3 className="font-semibold">
            {request.patient?.name || 'Paciente'}
          </h3>

          <p className="text-sm text-muted-foreground mt-2">
            {examsCount} exame{examsCount !== 1 ? 's' : ''} solicitado{examsCount !== 1 ? 's' : ''}
          </p>

          {request.notes && (
            <p className="text-sm text-muted-foreground mt-2 italic">
              {request.notes}
            </p>
          )}
        </div>

        <div className="flex gap-2">
          <Button variant="outline" size="sm" onClick={onEdit}>
            <Pencil className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="sm">
            <FileText className="h-4 w-4" />
          </Button>
        </div>
      </div>

      <details className="mt-4">
        <summary className="cursor-pointer text-sm font-medium text-muted-foreground hover:text-foreground">
          Ver exames solicitados
        </summary>
        <div className="mt-2 p-3 bg-muted rounded-md">
          <pre className="text-sm whitespace-pre-wrap">{request.exams}</pre>
        </div>
      </details>
    </Card>
  )
}
