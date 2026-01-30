'use client'

import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
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
import { Plus, FileText, Calendar } from 'lucide-react'
import {
  getAllLabRequests,
  createLabRequest,
  type LabRequest
} from '@/lib/api/lab-requests'
import {
  getAllLabRequestTemplates,
  getLabRequestTemplateById,
  type LabRequestTemplate
} from '@/lib/api/lab-request-templates'
import { getPatients, type Patient } from '@/lib/api/patients'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { PageHeader } from '@/components/layout/page-header'

export default function LabRequestsPage() {
  const router = useRouter()
  const [showCreateForm, setShowCreateForm] = useState(false)

  // Fetch lab requests
  const { data: requestsData, isLoading: requestsLoading } = useQuery({
    queryKey: ['lab-requests'],
    queryFn: () => getAllLabRequests({ limit: 50, offset: 0 })
  })

  const requests = requestsData?.data || []

  return (
    <div className="container mx-auto py-8 space-y-8">
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
            <LabRequestCard key={request.id} request={request} />
          ))}
        </div>
      )}
    </div>
  )
}

function CreateLabRequestForm({ onSuccess }: { onSuccess: () => void }) {
  const [selectedPatientId, setSelectedPatientId] = useState('')
  const [selectedTemplateId, setSelectedTemplateId] = useState<string>('')
  const [date, setDate] = useState(format(new Date(), 'yyyy-MM-dd'))
  const [exams, setExams] = useState('')
  const [notes, setNotes] = useState('')

  const queryClient = useQueryClient()

  // Fetch patients
  const { data: patients = [] } = useQuery({
    queryKey: ['patients'],
    queryFn: () => getPatients()
  })

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

    if (!selectedPatientId) {
      toast.error('Selecione um paciente')
      return
    }

    if (!exams.trim()) {
      toast.error('Adicione pelo menos um exame')
      return
    }

    createMutation.mutate({
      patientId: selectedPatientId,
      date,
      exams: exams.trim(),
      notes: notes.trim() || undefined
    })
  }

  return (
    <Card className="p-6 mb-6">
      <h2 className="text-xl font-semibold mb-4">Novo Pedido de Exames</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <div>
            <Label>Paciente *</Label>
            <Select value={selectedPatientId} onValueChange={setSelectedPatientId}>
              <SelectTrigger>
                <SelectValue placeholder="Selecione o paciente" />
              </SelectTrigger>
              <SelectContent>
                {patients.map((patient) => (
                  <SelectItem key={patient.id} value={patient.id}>
                    {patient.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div>
            <Label>Data *</Label>
            <Input
              type="date"
              value={date}
              onChange={(e) => setDate(e.target.value)}
              required
            />
          </div>
        </div>

        <div>
          <Label>Template (opcional)</Label>
          <Select value={selectedTemplateId} onValueChange={handleTemplateSelect}>
            <SelectTrigger>
              <SelectValue placeholder="Selecione um template ou digite manualmente" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="">Nenhum template</SelectItem>
              {sortedTemplates.map((template) => (
                <SelectItem key={template.id} value={template.id}>
                  {template.name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <p className="text-xs text-muted-foreground mt-1">
            Os exames do template serão inseridos no campo abaixo em ordem alfabética
          </p>
        </div>

        <div>
          <Label>Exames Solicitados * (um por linha)</Label>
          <Textarea
            value={exams}
            onChange={(e) => setExams(e.target.value)}
            placeholder="Digite os nomes dos exames, um por linha. Exemplo:&#10;Hemograma Completo&#10;Glicemia de Jejum&#10;Colesterol Total"
            rows={12}
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
      </form>
    </Card>
  )
}

function LabRequestCard({ request }: { request: LabRequest }) {
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

        <Button variant="outline" size="sm">
          <FileText className="h-4 w-4" />
        </Button>
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
