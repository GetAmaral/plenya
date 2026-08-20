'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useMutation, useQuery } from '@tanstack/react-query'
import { toast } from 'sonner'
import { formatDate, formatDateOnly, isPastDateOnly } from '@/lib/format-date'
import {
  Plus,
  FileText,
  FlaskConical,
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

import { useRequireSelectedPatient } from '@/lib/use-require-selected-patient'
import { SelectedPatientHeader } from '@/components/patients/SelectedPatientHeader'
import { SendDocumentEmailButton } from '@/components/clinical/send-document-email-button'
import { SendDocumentWhatsAppButton } from '@/components/clinical/send-document-whatsapp-button'
import { deletePrescription, listPrescriptions, openPrescriptionPdf } from '@/lib/api/prescriptions'
import type { Prescription, PrescriptionStatus } from '@/lib/api/prescriptions'

export default function PrescriptionsPage() {
  const router = useRouter()
  const { selectedPatient, isLoading: loadingPatient } = useRequireSelectedPatient()
  const [statusFilter, setStatusFilter] = useState<PrescriptionStatus | 'all' | 'used'>('all')

  // Query prescriptions
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['prescriptions', selectedPatient?.id, statusFilter],
    queryFn: () =>
      listPrescriptions({
        patientId: selectedPatient?.id,
        // 'used' não é status no banco: dispensação é a flag isUsed. Mandar 'used' para a API
        // devolvia lista vazia com a tela de "nenhuma prescrição".
        status: statusFilter === 'all' || statusFilter === 'used' ? undefined : statusFilter,
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

  // "Usados" é dispensação, não status: filtra aqui, com o que a API já devolveu.
  const prescriptions = (data || []).filter((p) => statusFilter !== 'used' || p.isUsed)
  const hasSignedPrescriptions = prescriptions.some((p) => p.signedPdfPath)

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
        <div className="flex flex-col gap-2 sm:flex-row">
          <Button variant="outline" onClick={() => router.push('/prescriptions/new?tipo=manipulado')}>
            <FlaskConical className="mr-2 h-4 w-4" />
            Manipulado
          </Button>
          <Button onClick={() => router.push('/prescriptions/new')}>
            <Plus className="mr-2 h-4 w-4" />
            Nova prescrição
          </Button>
        </div>
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
                <TableHead>Itens</TableHead>
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
                // Manipulado e industrializado convivem na mesma lista: a coluna mostra o que a
                // receita tem de fato, senão toda receita magistral apareceria como "0 medicamentos".
                const isCompounded = prescription.type === 'compounded'
                const itemCount = isCompounded
                  ? prescription.formulas?.length || 0
                  : prescription.medications?.length || 0
                const itemLabel = isCompounded
                  ? itemCount === 1
                    ? 'fórmula'
                    : 'fórmulas'
                  : itemCount === 1
                    ? 'medicamento'
                    : 'medicamentos'
                const itemNames = isCompounded
                  ? (prescription.formulas ?? []).map((f, i) => f.name || `Fórmula ${i + 1}`)
                  : (prescription.medications ?? []).map((m) => m.medicationName)
                const uniqueCategories = Array.from(
                  new Set(
                    isCompounded
                      ? (prescription.formulas ?? []).map((f) => f.highestCategory)
                      : (prescription.medications ?? []).map((m) => m.category)
                  )
                )

                return (
                  <TableRow key={prescription.id}>
                    {/* Medications */}
                    <TableCell>
                      <div className="flex flex-wrap items-center gap-2">
                        <span className="font-medium">
                          {itemCount} {itemLabel}
                        </span>
                        {isCompounded && (
                          <Badge variant="outline" className="bg-gold-50 text-xs text-gold-600">
                            Manipulado
                          </Badge>
                        )}
                      </div>
                      {itemNames.length > 0 && (
                        <div className="mt-1 text-sm text-muted-foreground">
                          {itemNames.slice(0, 2).join(', ')}
                          {itemNames.length > 2 && '...'}
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
                      {formatDate(prescription.prescriptionDate, 'dd/MM/yyyy')}
                    </TableCell>

                    {/* Valid Until */}
                    <TableCell>
                      {/* DATE pura: formatDate mostraria o dia anterior em BRT. */}
                      {formatDateOnly(prescription.validUntil)}
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
                      {prescription.signedPdfPath ? (
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
  const isExpired = isPastDateOnly(validUntil)

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
  const isSigned = !!prescription.signedPdfPath
  const canEdit = !isSigned

  const remove = useMutation({
    mutationFn: () => deletePrescription(prescription.id),
    onSuccess: () => {
      toast.success('Rascunho excluído')
      onRefetch()
    },
    onError: (error: any) =>
      toast.error('Erro ao excluir', {
        description: error.response?.data?.error || error.message,
      }),
  })

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

      {/* Download do PDF — pelo endpoint autenticado. O caminho em signedPdfPath aponta para o
          /uploads estático, que deixou de ser servido no hardening: abrir direto dava 404. */}
      {isSigned && (
        <Button
          variant="ghost"
          size="sm"
          title="Baixar PDF"
          onClick={() =>
            openPrescriptionPdf(prescription.id).catch((e) =>
              toast.error(e?.message ?? 'Falha ao baixar o PDF')
            )
          }
        >
          <Download className="h-4 w-4" />
        </Button>
      )}

      {/* Enviar: um botão por canal, sempre por ato explícito. Assinar não manda nada. */}
      {isSigned && prescription.patientId && (
        <>
          <SendDocumentWhatsAppButton
            patientId={prescription.patientId}
            docType="prescription"
            docId={prescription.id}
          />
          <SendDocumentEmailButton
            patientId={prescription.patientId}
            docType="prescription"
            docId={prescription.id}
          />
        </>
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

      {/* Excluir rascunho. O botão existia sem onClick nenhum: clicar não fazia nada. */}
      {canEdit && (
        <AlertDialog>
          <AlertDialogTrigger asChild>
            <Button variant="ghost" size="sm" className="text-destructive" title="Excluir rascunho">
              <Trash2 className="h-4 w-4" />
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Excluir este rascunho?</AlertDialogTitle>
              <AlertDialogDescription>
                O rascunho será removido da lista. Receita assinada nunca é excluída.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Cancelar</AlertDialogCancel>
              <AlertDialogAction
                onClick={() => remove.mutate()}
                className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
              >
                Excluir
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      )}
    </div>
  )
}
