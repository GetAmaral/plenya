'use client'

import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import {
  Upload,
  CheckCircle2,
  AlertCircle,
  Calendar,
  Trash2,
  RefreshCw,
  Shield,
  AlertTriangle,
} from 'lucide-react'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import { Switch } from '@/components/ui/switch'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'

import { listCertificates, uploadCertificate, deleteCertificate, toggleCertificateActive } from '@/lib/api/certificates'
import { listByRole } from '@/lib/api/users'
import { useAuthStore, isGranted } from '@/lib/auth-store'

export default function CertificatesPage() {
  const [isUploadDialogOpen, setIsUploadDialogOpen] = useState(false)
  const [selectedDoctorId, setSelectedDoctorId] = useState<string>('')
  const [password, setPassword] = useState('')
  const [file, setFile] = useState<File | null>(null)

  const queryClient = useQueryClient()
  const { user } = useAuthStore()
  const isAdmin = user ? isGranted(user, 'admin') : false

  // Query: Lista de certificados
  const { data: certificates, isLoading: loadingCertificates } = useQuery({
    queryKey: ['certificates'],
    queryFn: listCertificates,
  })

  // Query: Lista de médicos (apenas para admin)
  const { data: doctors, isLoading: loadingDoctors } = useQuery({
    queryKey: ['doctors'],
    queryFn: () => listByRole('doctor'),
    enabled: isAdmin,
  })

  // Mutation: Upload de certificado
  const uploadMutation = useMutation({
    mutationFn: async () => {
      if (!selectedDoctorId || !file || !password) {
        throw new Error('Todos os campos são obrigatórios')
      }

      const formData = new FormData()
      formData.append('doctorId', selectedDoctorId)
      formData.append('password', password)
      formData.append('certificate', file)

      return uploadCertificate(formData)
    },
    onSuccess: (data) => {
      toast.success('Certificado enviado com sucesso!', {
        description: data.message || 'O certificado foi validado e salvo com segurança.',
      })
      // Reset form
      setFile(null)
      setPassword('')
      setSelectedDoctorId('')
      setIsUploadDialogOpen(false)
      // Invalidar cache
      queryClient.invalidateQueries({ queryKey: ['certificates'] })
    },
    onError: (error: any) => {
      // Verificar se é erro de CPF divergente (status 409)
      if (error.response?.status === 409 && error.response?.data?.error === 'CPF_MISMATCH') {
        const data = error.response.data
        const userCPF = data.userCPF || 'não cadastrado'
        const certCPF = data.certificateCPF || 'não identificado'

        toast.error('CPF do certificado não corresponde!', {
          description: (
            <>
              <br />
              CPF do usuário: {userCPF}
              <br />
              <br />
              CPF do certificado: {certCPF}
            </>
          ),
          duration: 8000,
        })
      } else {
        toast.error('Erro ao enviar certificado', {
          description: error.response?.data?.error || error.message,
        })
      }
    },
  })

  // Mutation: Deletar certificado
  const deleteMutation = useMutation({
    mutationFn: (userId: string) => deleteCertificate(userId),
    onSuccess: (data) => {
      toast.success('Certificado removido', {
        description: data.message,
      })
      queryClient.invalidateQueries({ queryKey: ['certificates'] })
    },
    onError: (error: any) => {
      toast.error('Erro ao remover certificado', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  // Mutation: Toggle certificado ativo/inativo
  const toggleMutation = useMutation({
    mutationFn: (userId: string) => toggleCertificateActive(userId),
    onSuccess: (data) => {
      toast.success(data.message)
      queryClient.invalidateQueries({ queryKey: ['certificates'] })
    },
    onError: (error: any) => {
      toast.error('Erro ao alterar status do certificado', {
        description: error.response?.data?.error || error.message,
      })
    },
  })

  const handleDelete = (userId: string, userName: string) => {
    if (confirm(`Tem certeza que deseja remover o certificado de ${userName}?`)) {
      deleteMutation.mutate(userId)
    }
  }

  const handleToggle = (userId: string) => {
    toggleMutation.mutate(userId)
  }

  const handleUpload = () => {
    uploadMutation.mutate()
  }

  if (loadingCertificates) {
    return (
      <div className="container mx-auto py-8">
        <div className="flex items-center justify-center">
          <RefreshCw className="h-8 w-8 animate-spin text-muted-foreground" />
        </div>
      </div>
    )
  }

  return (
    <div className="container mx-auto py-8 max-w-6xl">
      {/* Header */}
      <div className="flex items-center justify-between mb-6">
        <div>
          <h1 className="text-3xl font-bold">Certificados Digitais</h1>
          <p className="text-muted-foreground mt-1">
            Gerenciar certificados ICP-Brasil para assinatura digital de prescrições
          </p>
        </div>
        {isAdmin && (
          <Button onClick={() => setIsUploadDialogOpen(true)}>
            <Upload className="mr-2 h-4 w-4" />
            Adicionar Certificado
          </Button>
        )}
      </div>

      {/* Info Alert */}
      <Alert className="mb-6">
        <Shield className="h-4 w-4" />
        <AlertTitle>Certificados ICP-Brasil</AlertTitle>
        <AlertDescription>
          {isAdmin
            ? 'Gerencie os certificados digitais A1 dos médicos. Certificados expirados devem ser renovados para continuar assinando prescrições.'
            : 'Seu certificado digital permite assinar prescrições com validade jurídica. Renove antes de expirar.'}
        </AlertDescription>
      </Alert>

      {/* Lista de Certificados */}
      <Card>
        <CardHeader>
          <CardTitle>Certificados Cadastrados</CardTitle>
          <CardDescription>
            {certificates?.length || 0} certificado(s) no sistema
          </CardDescription>
        </CardHeader>
        <CardContent>
          {certificates && certificates.length > 0 ? (
            <Table>
              <TableHeader>
                <TableRow>
                  {isAdmin && <TableHead>Médico</TableHead>}
                  <TableHead>Titular do Certificado</TableHead>
                  <TableHead>CPF do Certificado</TableHead>
                  <TableHead>Email</TableHead>
                  <TableHead>Válido Até</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead className="text-right">Ações</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {certificates.map((cert) => (
                  <TableRow key={cert.userId}>
                    {isAdmin && <TableCell className="font-medium">{cert.userName}</TableCell>}
                    <TableCell>
                      {cert.certificateName ? (
                        <span className="font-medium">{cert.certificateName}</span>
                      ) : (
                        <span className="text-xs text-muted-foreground">
                          Serial: {cert.certificateSerial}
                        </span>
                      )}
                    </TableCell>
                    <TableCell>
                      {cert.certificateCPF ? (
                        <span className="font-mono text-sm">{cert.certificateCPF}</span>
                      ) : (
                        <span className="text-xs text-muted-foreground">Não extraído</span>
                      )}
                    </TableCell>
                    <TableCell>{cert.userEmail}</TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Calendar className="h-4 w-4 text-muted-foreground" />
                        {format(new Date(cert.validUntil), 'dd/MM/yyyy', { locale: ptBR })}
                      </div>
                      <p className="text-xs text-muted-foreground">
                        {cert.daysUntilExpiry > 0
                          ? `${cert.daysUntilExpiry} dias restantes`
                          : `Expirado há ${Math.abs(cert.daysUntilExpiry)} dias`}
                      </p>
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center gap-3">
                        {isAdmin && (
                          <Switch
                            checked={cert.certificateActive}
                            onCheckedChange={() => handleToggle(cert.userId)}
                            disabled={toggleMutation.isPending}
                          />
                        )}
                        <div className="flex flex-col gap-1">
                          {cert.certificateActive ? (
                            cert.isExpired ? (
                              <Badge variant="destructive" className="gap-1 w-fit">
                                <AlertCircle className="h-3 w-3" />
                                Expirado
                              </Badge>
                            ) : cert.needsRenewal ? (
                              <Badge variant="outline" className="gap-1 bg-yellow-50 text-yellow-700 w-fit">
                                <AlertTriangle className="h-3 w-3" />
                                Renovar
                              </Badge>
                            ) : (
                              <span className="text-sm text-green-600 font-medium">Ativo</span>
                            )
                          ) : (
                            <span className="text-sm text-muted-foreground">Desativado</span>
                          )}
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className="text-right">
                      <div className="flex items-center justify-end gap-2">
                        {(isAdmin || cert.userId === user?.id) && (
                          <>
                            <Button
                              variant="ghost"
                              size="sm"
                              onClick={() => {
                                setSelectedDoctorId(cert.userId)
                                setIsUploadDialogOpen(true)
                              }}
                            >
                              <RefreshCw className="h-4 w-4" />
                            </Button>
                            <Button
                              variant="ghost"
                              size="sm"
                              className="text-destructive"
                              onClick={() => handleDelete(cert.userId, cert.userName)}
                            >
                              <Trash2 className="h-4 w-4" />
                            </Button>
                          </>
                        )}
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          ) : (
            <div className="text-center py-12 text-muted-foreground">
              <Shield className="h-12 w-12 mx-auto mb-4 opacity-50" />
              <p>Nenhum certificado cadastrado.</p>
              {isAdmin && (
                <Button onClick={() => setIsUploadDialogOpen(true)} className="mt-4">
                  <Upload className="mr-2 h-4 w-4" />
                  Adicionar Primeiro Certificado
                </Button>
              )}
            </div>
          )}
        </CardContent>
      </Card>

      {/* Dialog de Upload */}
      <Dialog open={isUploadDialogOpen} onOpenChange={setIsUploadDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>
              {selectedDoctorId ? 'Substituir Certificado' : 'Adicionar Certificado'}
            </DialogTitle>
            <DialogDescription>
              Envie o arquivo .pfx ou .p12 do certificado ICP-Brasil e a senha de proteção.
            </DialogDescription>
          </DialogHeader>

          <div className="space-y-4">
            {/* Select Médico (apenas se admin e não for substituição) */}
            {isAdmin && !selectedDoctorId && (
              <div className="space-y-2">
                <Label htmlFor="doctor">Médico *</Label>
                <Select value={selectedDoctorId} onValueChange={setSelectedDoctorId}>
                  <SelectTrigger id="doctor">
                    <SelectValue placeholder="Selecione o médico" />
                  </SelectTrigger>
                  <SelectContent>
                    {loadingDoctors && (
                      <SelectItem value="loading" disabled>
                        Carregando médicos...
                      </SelectItem>
                    )}
                    {doctors?.map((doctor) => (
                      <SelectItem key={doctor.id} value={doctor.id}>
                        {doctor.name} ({doctor.email})
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            )}

            {/* Arquivo */}
            <div className="space-y-2">
              <Label htmlFor="file">Arquivo do Certificado *</Label>
              <Input
                id="file"
                type="file"
                accept=".pfx,.p12"
                onChange={(e) => setFile(e.target.files?.[0] || null)}
              />
              {file && (
                <p className="text-sm text-muted-foreground">
                  <CheckCircle2 className="inline h-3 w-3 mr-1" />
                  {file.name}
                </p>
              )}
            </div>

            {/* Senha */}
            <div className="space-y-2">
              <Label htmlFor="password">Senha do Certificado *</Label>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Digite a senha do certificado"
              />
            </div>

            {/* Botões */}
            <div className="flex justify-end gap-2 pt-4">
              <Button
                variant="outline"
                onClick={() => {
                  setIsUploadDialogOpen(false)
                  setSelectedDoctorId('')
                  setFile(null)
                  setPassword('')
                }}
              >
                Cancelar
              </Button>
              <Button
                onClick={handleUpload}
                disabled={
                  !selectedDoctorId || !file || !password || uploadMutation.isPending
                }
              >
                {uploadMutation.isPending ? (
                  <>
                    <RefreshCw className="mr-2 h-4 w-4 animate-spin" />
                    Enviando...
                  </>
                ) : (
                  <>
                    <Upload className="mr-2 h-4 w-4" />
                    {selectedDoctorId ? 'Substituir' : 'Adicionar'}
                  </>
                )}
              </Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  )
}
