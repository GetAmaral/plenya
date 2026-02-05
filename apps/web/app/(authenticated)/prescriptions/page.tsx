'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useQuery } from '@tanstack/react-query'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import {
  Plus,
  FileText,
  Download,
  Edit,
  Copy,
  Trash2,
  Eye,
  AlertCircle,
  CheckCircle2,
  Clock,
  XCircle,
} from 'lucide-react'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'

import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { listPrescriptions } from '@/lib/api/prescriptions'
import type { Prescription, PrescriptionStatus } from '@/lib/api/prescriptions'

export default function PrescriptionsPage() {
  const router = useRouter()
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()
  const [statusFilter, setStatusFilter] = useState<PrescriptionStatus | 'all'>('all')

  // Query prescriptions
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['prescriptions', selectedPatient?.id, statusFilter],
    queryFn: () =>
      listPrescriptions({
        patientId: selectedPatient?.id,
        status: statusFilter === 'all' ? undefined : statusFilter,
        limit: 100,
      }),
    enabled: !!selectedPatient?.id,
  })

  if (loadingPatient || isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-center space-y-4">
          <FileText className="h-12 w-12 animate-pulse text-muted-foreground mx-auto" />
          <p className="text-muted-foreground">Carregando prescrições...</p>
        </div>
      </div>
    )
  }

  const prescriptions = data?.data || []
  const hasSignedPrescriptions = prescriptions.some((p) => p.signedPDFPath)

  return (
    <div className="container mx-auto py-8">
      <SelectedPatientHeader />

      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-3xl font-bold">Prescrições Digitais</h1>
          <p className="text-muted-foreground mt-1">
            {prescriptions.length} prescrição(ões) encontrada(s)
          </p>
        </div>
        <Button onClick={() => router.push('/prescriptions/new')}>
          <Plus className="mr-2 h-4 w-4" />
          Nova Prescrição
        </Button>
      </div>

      {/* Info Alert */}
      {hasSignedPrescriptions && (
        <Alert className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Prescrições Assinadas</AlertTitle>
          <AlertDescription>
            Prescrições com PDF assinado não podem ser editadas. Para fazer alterações, use a
            função "Duplicar" e edite a cópia.
          </AlertDescription>
        </Alert>
      )}

      {/* Filters */}
      <Card className="mb-6">
        <CardHeader>
          <CardTitle>Filtros</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex gap-4">
            <div className="w-64">
              <Select
                value={statusFilter}
                onValueChange={(value) => setStatusFilter(value as any)}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">Todos</SelectItem>
                  <SelectItem value="active">Ativos</SelectItem>
                  <SelectItem value="expired">Expirados</SelectItem>
                  <SelectItem value="used">Usados</SelectItem>
                  <SelectItem value="cancelled">Cancelados</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Table */}
      {prescriptions.length === 0 ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <FileText className="h-12 w-12 text-muted-foreground mb-4" />
            <p className="text-muted-foreground text-center">
              Nenhuma prescrição encontrada para este paciente.
            </p>
            <Button onClick={() => router.push('/prescriptions/new')} className="mt-4">
              <Plus className="mr-2 h-4 w-4" />
              Criar Primeira Prescrição
            </Button>
          </CardContent>
        </Card>
      ) : (
        <Card>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Medicamentos</TableHead>
                <TableHead>Categorias</TableHead>
                <TableHead>Data</TableHead>
                <TableHead>Validade</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>PDF</TableHead>
                <TableHead className="text-right">Ações</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {prescriptions.map((prescription) => {
                const medicationCount = prescription.medications?.length || 0
                const uniqueCategories = Array.from(
                  new Set(prescription.medications?.map((m) => m.category) || [])
                )

                return (
                  <TableRow key={prescription.id}>
                    {/* Medications */}
                    <TableCell>
                      <div className="font-medium">
                        {medicationCount} {medicationCount === 1 ? 'medicamento' : 'medicamentos'}
                      </div>
                      {prescription.medications && prescription.medications.length > 0 && (
                        <div className="text-sm text-muted-foreground mt-1">
                          {prescription.medications
                            .slice(0, 2)
                            .map((m) => m.medicationName)
                            .join(', ')}
                          {prescription.medications.length > 2 && '...'}
                        </div>
                      )}
                    </TableCell>

                    {/* Categories */}
                    <TableCell>
                      <div className="flex flex-wrap gap-1">
                        {uniqueCategories.map((category) => (
                          <CategoryBadge key={category} category={category} />
                        ))}
                      </div>
                    </TableCell>

                    {/* Date */}
                    <TableCell>
                      {format(new Date(prescription.prescriptionDate), 'dd/MM/yyyy', {
                        locale: ptBR,
                      })}
                    </TableCell>

                    {/* Valid Until */}
                    <TableCell>
                      {format(new Date(prescription.validUntil), 'dd/MM/yyyy', { locale: ptBR })}
                    </TableCell>

                    {/* Status */}
                    <TableCell>
                      <StatusBadge
                        status={prescription.status}
                        isUsed={prescription.isUsed}
                        validUntil={prescription.validUntil}
                      />
                    </TableCell>

                    {/* PDF Status */}
                    <TableCell>
                      {prescription.signedPDFPath ? (
                        <Badge variant="outline" className="bg-green-50 text-green-700">
                          <CheckCircle2 className="mr-1 h-3 w-3" />
                          Assinado
                        </Badge>
                      ) : (
                        <Badge variant="outline" className="bg-gray-50">
                          <Clock className="mr-1 h-3 w-3" />
                          Rascunho
                        </Badge>
                      )}
                    </TableCell>

                    {/* Actions */}
                    <TableCell className="text-right">
                      <PrescriptionActions
                        prescription={prescription}
                        onRefetch={refetch}
                        router={router}
                      />
                    </TableCell>
                  </TableRow>
                )
              })}
            </TableBody>
          </Table>
        </Card>
      )}
    </div>
  )
}

function CategoryBadge({ category }: { category: string }) {
  const variants: Record<string, { label: string; className: string }> = {
    simple: { label: 'Simples', className: 'bg-blue-50 text-blue-700' },
    c1: { label: 'C1', className: 'bg-orange-50 text-orange-700' },
    c5: { label: 'C5', className: 'bg-red-50 text-red-700' },
    antibiotic: { label: 'Antibiótico', className: 'bg-purple-50 text-purple-700' },
    glp1: { label: 'GLP-1', className: 'bg-green-50 text-green-700' },
  }

  const variant = variants[category] || variants.simple

  return (
    <Badge variant="outline" className={variant.className}>
      {variant.label}
    </Badge>
  )
}

function StatusBadge({
  status,
  isUsed,
  validUntil,
}: {
  status: string
  isUsed: boolean
  validUntil: string
}) {
  const isExpired = new Date(validUntil) < new Date()

  if (isUsed) {
    return (
      <Badge variant="outline" className="bg-gray-50">
        <CheckCircle2 className="mr-1 h-3 w-3" />
        Dispensado
      </Badge>
    )
  }

  if (isExpired) {
    return (
      <Badge variant="outline" className="bg-red-50 text-red-700">
        <XCircle className="mr-1 h-3 w-3" />
        Expirado
      </Badge>
    )
  }

  if (status === 'cancelled') {
    return (
      <Badge variant="outline" className="bg-gray-50">
        <XCircle className="mr-1 h-3 w-3" />
        Cancelado
      </Badge>
    )
  }

  return (
    <Badge variant="outline" className="bg-green-50 text-green-700">
      <CheckCircle2 className="mr-1 h-3 w-3" />
      Ativo
    </Badge>
  )
}

function PrescriptionActions({
  prescription,
  onRefetch,
  router,
}: {
  prescription: Prescription
  onRefetch: () => void
  router: any
}) {
  const isSigned = !!prescription.signedPDFPath
  const canEdit = !isSigned

  return (
    <div className="flex items-center justify-end gap-2">
      {/* View/Details */}
      <Button
        variant="ghost"
        size="sm"
        onClick={() => router.push(`/prescriptions/${prescription.id}`)}
      >
        <Eye className="h-4 w-4" />
      </Button>

      {/* Download PDF (if signed) */}
      {isSigned && prescription.signedPDFPath && (
        <Button
          variant="ghost"
          size="sm"
          onClick={() => window.open(prescription.signedPDFPath!, '_blank')}
        >
          <Download className="h-4 w-4" />
        </Button>
      )}

      {/* Edit (if not signed) */}
      {canEdit && (
        <Button
          variant="ghost"
          size="sm"
          onClick={() => router.push(`/prescriptions/${prescription.id}/edit`)}
        >
          <Edit className="h-4 w-4" />
        </Button>
      )}

      {/* Duplicate (if signed) */}
      {isSigned && (
        <Button
          variant="ghost"
          size="sm"
          onClick={() => router.push(`/prescriptions/${prescription.id}/duplicate`)}
        >
          <Copy className="h-4 w-4" />
        </Button>
      )}

      {/* Delete (only if not signed) */}
      {canEdit && (
        <Button variant="ghost" size="sm" className="text-destructive">
          <Trash2 className="h-4 w-4" />
        </Button>
      )}
    </div>
  )
}
