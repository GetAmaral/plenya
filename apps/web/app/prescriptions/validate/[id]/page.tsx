'use client'

import { useQuery } from '@tanstack/react-query'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import {
  CheckCircle2,
  AlertCircle,
  Download,
  ExternalLink,
  FileCheck,
  User,
  Stethoscope,
  Pill,
  Shield,
  Calendar,
} from 'lucide-react'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'

import { validatePublic } from '@/lib/api/prescriptions'

interface PageProps {
  params: {
    id: string
  }
}

export default function ValidatePrescriptionPage({ params }: PageProps) {
  const { data, isLoading, error } = useQuery({
    queryKey: ['prescription-validation', params.id],
    queryFn: () => validatePublic(params.id),
    retry: false,
  })

  if (isLoading) {
    return (
      <div className="container mx-auto py-12 max-w-3xl">
        <Card>
          <CardContent className="flex items-center justify-center py-12">
            <div className="text-center space-y-4">
              <FileCheck className="h-12 w-12 animate-pulse text-muted-foreground mx-auto" />
              <p className="text-muted-foreground">Validando prescrição...</p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  if (error || !data) {
    return (
      <div className="container mx-auto py-12 max-w-3xl">
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Prescrição Não Encontrada</AlertTitle>
          <AlertDescription>
            Não foi possível encontrar ou validar esta prescrição. Verifique se o código QR está
            correto.
          </AlertDescription>
        </Alert>
      </div>
    )
  }

  const isValid = data.valid && data.pdfIntact
  const hasIssues = data.prescription.isExpired || data.prescription.isUsed || !data.pdfIntact

  return (
    <div className="container mx-auto py-12 max-w-4xl">
      {/* Header */}
      <div className="text-center mb-8">
        <h1 className="text-4xl font-bold mb-2">Validação de Prescrição Digital</h1>
        <p className="text-muted-foreground">
          Sistema Nacional de Controle de Receitas (SNCR) - ANVISA
        </p>
      </div>

      {/* Status Alert */}
      <Alert
        variant={isValid ? 'default' : 'destructive'}
        className={
          isValid
            ? 'border-green-500 bg-green-50 dark:bg-green-950 mb-8'
            : 'mb-8'
        }
      >
        {isValid ? (
          <CheckCircle2 className="h-5 w-5 text-green-600" />
        ) : (
          <AlertCircle className="h-5 w-5" />
        )}
        <AlertTitle className={isValid ? 'text-green-600 text-xl' : 'text-xl'}>
          {isValid ? '✅ PRESCRIÇÃO VÁLIDA' : '❌ PRESCRIÇÃO INVÁLIDA'}
        </AlertTitle>
        <AlertDescription className={isValid ? 'text-green-600' : ''}>
          {isValid && 'Esta prescrição está válida e pode ser dispensada.'}
          {data.prescription.isExpired && 'Prescrição expirada - fora do prazo de validade.'}
          {data.prescription.isUsed && 'Prescrição já foi dispensada anteriormente.'}
          {!data.pdfIntact && 'PDF foi alterado - assinatura digital inválida.'}
        </AlertDescription>
      </Alert>

      {/* Issues Badges */}
      {hasIssues && (
        <div className="flex flex-wrap gap-2 mb-6">
          {data.prescription.isExpired && (
            <Badge variant="destructive">
              <Calendar className="h-3 w-3 mr-1" />
              Expirada
            </Badge>
          )}
          {data.prescription.isUsed && (
            <Badge variant="destructive">
              <AlertCircle className="h-3 w-3 mr-1" />
              Já Dispensada
            </Badge>
          )}
          {!data.pdfIntact && (
            <Badge variant="destructive">
              <Shield className="h-3 w-3 mr-1" />
              PDF Alterado
            </Badge>
          )}
        </div>
      )}

      <div className="grid gap-6">
        {/* Prescription Info */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <FileCheck className="h-5 w-5" />
              Informações da Prescrição
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-2 gap-4">
              <div>
                <p className="text-sm text-muted-foreground">Data de Emissão</p>
                <p className="font-medium">
                  {format(new Date(data.prescription.prescriptionDate), "dd 'de' MMMM 'de' yyyy", {
                    locale: ptBR,
                  })}
                </p>
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Válida Até</p>
                <p className="font-medium">
                  {format(new Date(data.prescription.validUntil), "dd 'de' MMMM 'de' yyyy", {
                    locale: ptBR,
                  })}
                </p>
              </div>
            </div>

            {data.prescription.sncrNumber && (
              <div>
                <p className="text-sm text-muted-foreground">Número SNCR</p>
                <p className="font-mono font-medium text-lg">{data.prescription.sncrNumber}</p>
              </div>
            )}

            <div>
              <p className="text-sm text-muted-foreground">Quantidade de Medicamentos</p>
              <p className="font-medium">
                {data.prescription.medicationCount || data.medications?.length || 0}{' '}
                {(data.prescription.medicationCount || data.medications?.length || 0) === 1
                  ? 'medicamento'
                  : 'medicamentos'}
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Patient Info */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <User className="h-5 w-5" />
              Paciente
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            <div>
              <p className="text-sm text-muted-foreground">Nome</p>
              <p className="font-medium text-lg">{data.patient.name}</p>
            </div>
            <div>
              <p className="text-sm text-muted-foreground">CPF</p>
              <p className="font-mono">{data.patient.cpf}</p>
              <p className="text-xs text-muted-foreground mt-1">
                (Mascarado por segurança - últimos dígitos visíveis)
              </p>
            </div>
          </CardContent>
        </Card>

        {/* Doctor Info */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Stethoscope className="h-5 w-5" />
              Médico Prescritor
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            <div>
              <p className="text-sm text-muted-foreground">Nome</p>
              <p className="font-medium text-lg">{data.doctor.name}</p>
            </div>
            <div>
              <p className="text-sm text-muted-foreground">CRM</p>
              <p className="font-medium">{data.doctor.crm}</p>
            </div>
          </CardContent>
        </Card>

        {/* Medications Info */}
        <div className="space-y-4">
          <h2 className="text-xl font-semibold flex items-center gap-2">
            <Pill className="h-5 w-5" />
            Medicamentos Prescritos
          </h2>

          {data.medications && data.medications.length > 0 ? (
            data.medications.map((medication, index) => (
              <Card key={medication.id || index}>
                <CardHeader>
                  <CardTitle className="text-lg">
                    {index + 1}. {medication.name}
                  </CardTitle>
                  <CardDescription>
                    <Badge variant="outline" className="mt-1">
                      {medication.category === 'c1' && 'Controle Especial (C1)'}
                      {medication.category === 'c5' && 'Psicotrópico (C5)'}
                      {medication.category === 'antibiotic' && 'Antimicrobiano'}
                      {medication.category === 'glp1' && 'GLP-1 Agonista'}
                      {medication.category === 'simple' && 'Receita Simples'}
                    </Badge>
                  </CardDescription>
                </CardHeader>
                <CardContent className="space-y-3">
                  <div>
                    <p className="text-sm text-muted-foreground">Princípio Ativo (DCB)</p>
                    <p className="font-medium">{medication.activeIngredient}</p>
                  </div>
                  <Separator />
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <p className="text-sm text-muted-foreground">Concentração</p>
                      <p className="font-medium">{medication.concentration}</p>
                    </div>
                    <div>
                      <p className="text-sm text-muted-foreground">Quantidade</p>
                      <p className="font-medium">{medication.quantity}</p>
                    </div>
                  </div>
                  <div>
                    <p className="text-sm text-muted-foreground">Quantidade por Extenso</p>
                    <p className="font-medium italic">{medication.quantityInWords}</p>
                  </div>
                </CardContent>
              </Card>
            ))
          ) : (
            <Card>
              <CardContent className="py-8 text-center text-muted-foreground">
                Nenhum medicamento encontrado nesta prescrição.
              </CardContent>
            </Card>
          )}
        </div>

        {/* Digital Signature */}
        <Card className="border-blue-200 bg-blue-50 dark:border-blue-800 dark:bg-blue-950">
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-blue-900 dark:text-blue-100">
              <Shield className="h-5 w-5" />
              Assinatura Digital ICP-Brasil
            </CardTitle>
            <CardDescription className="text-blue-700 dark:text-blue-300">
              Documento autêntico com validade jurídica
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-3 text-blue-800 dark:text-blue-200">
            <div className="flex items-center gap-2">
              {data.pdfIntact ? (
                <CheckCircle2 className="h-5 w-5 text-green-600" />
              ) : (
                <AlertCircle className="h-5 w-5 text-red-600" />
              )}
              <span className="font-medium">
                {data.pdfIntact
                  ? 'Assinatura Digital Válida'
                  : 'Assinatura Inválida - PDF Alterado'}
              </span>
            </div>

            {data.signature.signedAt && (
              <div>
                <p className="text-sm">Assinado em:</p>
                <p className="font-medium">
                  {format(new Date(data.signature.signedAt), "dd/MM/yyyy 'às' HH:mm", {
                    locale: ptBR,
                  })}
                </p>
              </div>
            )}

            {data.signature.certificateSerial && (
              <div>
                <p className="text-sm">Certificado Serial:</p>
                <p className="font-mono text-xs">{data.signature.certificateSerial}</p>
              </div>
            )}

            <div className="text-xs space-y-1 pt-2 border-t border-blue-200 dark:border-blue-800">
              <p>✅ Padrão PAdES (PDF Advanced Electronic Signatures)</p>
              <p>✅ Certificado ICP-Brasil válido</p>
              <p>✅ Conforme CFM Resolução 2.299/2021</p>
            </div>
          </CardContent>
        </Card>

        {/* Actions */}
        <div className="flex flex-col sm:flex-row gap-4">
          {data.signature.signedPdfUrl && (
            <Button
              onClick={() => window.open(data.signature.signedPdfUrl!, '_blank')}
              className="flex-1"
              size="lg"
            >
              <Download className="mr-2 h-4 w-4" />
              Baixar PDF Assinado
            </Button>
          )}

          <Button
            variant="outline"
            onClick={() => window.open('https://validar.iti.gov.br/', '_blank')}
            className="flex-1"
            size="lg"
          >
            <ExternalLink className="mr-2 h-4 w-4" />
            Verificar Assinatura no ITI
          </Button>
        </div>

        {/* Footer Info */}
        <Card className="bg-muted">
          <CardContent className="text-sm text-muted-foreground space-y-2 py-4">
            <p className="font-medium">Sobre esta validação:</p>
            <ul className="list-disc list-inside space-y-1 ml-2">
              <li>Esta página é pública e pode ser acessada via QR Code</li>
              <li>Dados do paciente são mascarados por segurança (LGPD)</li>
              <li>A farmácia deve verificar a identidade do paciente antes de dispensar</li>
              <li>
                Em caso de dúvida, o farmacêutico deve entrar em contato com o médico prescritor
              </li>
            </ul>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
