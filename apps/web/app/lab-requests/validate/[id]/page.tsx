'use client'

import { use } from 'react'
import { useQuery } from '@tanstack/react-query'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import { CheckCircle2, AlertCircle, Shield, Calendar, FileText, Download } from 'lucide-react'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { apiClient } from '@/lib/api-client'

interface ValidationData {
  valid: boolean
  labRequest: {
    id: string
    requestDate: string
    examCount: number
    exams: string[]
  }
  patient: {
    name: string
    cpf: string
  }
  doctor: {
    name: string
    crm: string
  }
  signature?: {
    signedAt: string
    certificateSerial: string
    signedPdfUrl: string
    certificateName?: string
    certificateCPF?: string
    signedPdfHash?: string
  }
  error?: string
}

async function validateLabRequest(id: string): Promise<ValidationData> {
  return await apiClient.get<ValidationData>(`/api/v1/lab-requests/validate/${id}`)
}

export default function ValidateLabRequestPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params)

  const { data, isLoading, error } = useQuery({
    queryKey: ['lab-request-validation', id],
    queryFn: () => validateLabRequest(id),
  })

  if (isLoading) {
    return (
      <div className="container mx-auto py-8 max-w-2xl">
        <div className="flex items-center justify-center">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
        </div>
      </div>
    )
  }

  if (error || !data) {
    return (
      <div className="container mx-auto py-8 max-w-2xl">
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Erro ao validar pedido</AlertTitle>
          <AlertDescription>
            Não foi possível validar este pedido de exames. Verifique o link e tente novamente.
          </AlertDescription>
        </Alert>
      </div>
    )
  }

  return (
    <div className="container mx-auto py-8 max-w-2xl">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <FileText className="h-6 w-6" />
            Validação de Pedido de Exames
          </CardTitle>
          <CardDescription>
            Verificação pública de autenticidade
          </CardDescription>
        </CardHeader>

        <CardContent className="space-y-6">
          {/* Status */}
          <Alert variant={data.valid ? "default" : "destructive"}>
            {data.valid ? (
              <CheckCircle2 className="h-4 w-4" />
            ) : (
              <AlertCircle className="h-4 w-4" />
            )}
            <AlertTitle>
              {data.valid ? "✅ PEDIDO VÁLIDO" : "❌ PEDIDO INVÁLIDO"}
            </AlertTitle>
            <AlertDescription>
              {data.error || (data.valid 
                ? "Este pedido de exames é autêntico e foi assinado digitalmente."
                : "Não foi possível verificar a autenticidade deste pedido."
              )}
            </AlertDescription>
          </Alert>

          {data.valid && (
            <>
              {/* Informações do Paciente */}
              <div className="space-y-2">
                <h3 className="font-semibold text-sm text-muted-foreground uppercase">
                  Paciente
                </h3>
                <div className="rounded-lg border p-4 bg-muted/50">
                  <p className="font-medium">{data.patient.name}</p>
                  <p className="text-sm text-muted-foreground">
                    CPF: {data.patient.cpf}
                  </p>
                </div>
              </div>

              {/* Informações do Médico */}
              <div className="space-y-2">
                <h3 className="font-semibold text-sm text-muted-foreground uppercase">
                  Médico Solicitante
                </h3>
                <div className="rounded-lg border p-4 bg-muted/50">
                  <p className="font-medium">{data.doctor.name}</p>
                  <p className="text-sm text-muted-foreground">
                    {data.doctor.crm}
                  </p>
                </div>
              </div>

              {/* Informações do Pedido */}
              <div className="space-y-2">
                <h3 className="font-semibold text-sm text-muted-foreground uppercase">
                  Pedido
                </h3>
                <div className="rounded-lg border p-4 bg-muted/50 space-y-3">
                  <div className="flex items-center gap-2">
                    <Calendar className="h-4 w-4 text-muted-foreground" />
                    <span className="text-sm">
                      Solicitado em:{' '}
                      {format(new Date(data.labRequest.requestDate), "dd 'de' MMMM 'de' yyyy", {
                        locale: ptBR,
                      })}
                    </span>
                  </div>
                  <div className="text-sm">
                    <Badge variant="outline" className="font-mono">
                      {data.labRequest.examCount} exame(s)
                    </Badge>
                  </div>

                  {/* Lista de Exames (Accordion) */}
                  {data.labRequest.exams && data.labRequest.exams.length > 0 && (
                    <Accordion type="single" collapsible className="w-full">
                      <AccordionItem value="exams" className="border-0">
                        <AccordionTrigger className="text-sm font-medium hover:no-underline py-2">
                          Ver lista de exames
                        </AccordionTrigger>
                        <AccordionContent>
                          <ul className="space-y-1 mt-2">
                            {data.labRequest.exams.map((exam, index) => (
                              <li key={index} className="text-sm flex items-start gap-2">
                                <span className="text-muted-foreground">•</span>
                                <span>{exam}</span>
                              </li>
                            ))}
                          </ul>
                        </AccordionContent>
                      </AccordionItem>
                    </Accordion>
                  )}
                </div>
              </div>

              {/* Assinatura Digital */}
              {data.signature && (
                <div className="space-y-2">
                  <h3 className="font-semibold text-sm text-muted-foreground uppercase">
                    Assinatura Digital ICP-Brasil
                  </h3>
                  <div className="rounded-lg border-2 p-4 bg-green-50 border-green-500">
                    <div className="flex items-start gap-3">
                      <div className="flex-shrink-0">
                        <div className="bg-green-600 text-white px-2 py-1 rounded text-xs font-bold">
                          ICP-BRASIL
                        </div>
                      </div>
                      <div className="flex-1 space-y-3">
                        <div>
                          <p className="text-sm font-bold text-green-900 flex items-center gap-2">
                            <CheckCircle2 className="h-4 w-4" />
                            Assinatura Digital Válida
                          </p>
                          <p className="text-xs text-green-700 mt-1">
                            Documento assinado com certificado digital padrão ICP-Brasil (PAdES)
                          </p>
                        </div>

                        {data.signature.certificateName && (
                          <div className="bg-white/50 rounded p-2">
                            <p className="text-xs font-semibold text-green-900">Titular do Certificado:</p>
                            <p className="text-xs text-green-800">{data.signature.certificateName}</p>
                            {data.signature.certificateCPF && (
                              <p className="text-xs text-green-700 font-mono mt-1">
                                CPF: {data.signature.certificateCPF}
                              </p>
                            )}
                          </div>
                        )}

                        <div className="text-xs text-green-700 space-y-1">
                          <p>
                            Assinado em:{' '}
                            {format(
                              new Date(data.signature.signedAt),
                              "dd/MM/yyyy 'às' HH:mm",
                              { locale: ptBR }
                            )}
                          </p>
                          <p className="font-mono">Serial: {data.signature.certificateSerial}</p>
                          {data.signature.signedPdfHash && (
                            <p className="font-mono break-all">
                              Hash SHA-256: {data.signature.signedPdfHash.substring(0, 32)}...
                            </p>
                          )}
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              )}

              {/* Ações */}
              <div className="flex flex-col sm:flex-row gap-2 pt-4">
                {data.signature?.signedPdfUrl && (
                  <Button
                    className="flex-1"
                    onClick={() => {
                      // Construct full API URL for the PDF
                      const apiBaseUrl = process.env.NEXT_PUBLIC_API_URL?.replace('/api/v1', '') || 'http://localhost:3001'
                      const pdfUrl = data.signature!.signedPdfUrl.startsWith('http')
                        ? data.signature!.signedPdfUrl
                        : `${apiBaseUrl}${data.signature!.signedPdfUrl}`
                      window.open(pdfUrl, '_blank')
                    }}
                  >
                    <Download className="mr-2 h-4 w-4" />
                    Baixar PDF Assinado
                  </Button>
                )}
                <Button
                  variant="outline"
                  className="flex-1"
                  onClick={() => window.open('https://validar.iti.gov.br/', '_blank')}
                >
                  <Shield className="mr-2 h-4 w-4" />
                  Verificar no ITI
                </Button>
              </div>

              {/* Tutorial de Validação no ITI */}
              <Alert>
                <Shield className="h-4 w-4" />
                <AlertTitle>Como verificar a assinatura no ITI</AlertTitle>
                <AlertDescription className="text-xs space-y-2">
                  <p>
                    Para validar a autenticidade da assinatura digital no site oficial do
                    Instituto Nacional de Tecnologia da Informação (ITI):
                  </p>
                  <ol className="list-decimal list-inside space-y-1 ml-2">
                    <li>Clique em "Baixar PDF Assinado" acima</li>
                    <li>Acesse <span className="font-mono bg-muted px-1">https://validar.iti.gov.br</span></li>
                    <li>Faça upload do arquivo PDF baixado</li>
                    <li>O sistema mostrará: "✅ Assinatura válida - ICP-Brasil"</li>
                  </ol>
                  <p className="mt-2 text-muted-foreground">
                    Conforme <strong>Resolução CFM 2.299/2021</strong>, documentos médicos eletrônicos
                    devem ser assinados com certificado ICP-Brasil NGS2 (Nível de Garantia de Segurança 2).
                  </p>
                </AlertDescription>
              </Alert>

              {/* Info sobre Plenya EMR */}
              <div className="text-center text-xs text-muted-foreground pt-4 border-t">
                <p>Emitido por: <strong>Plenya EMR</strong> - Sistema de Prontuário Eletrônico</p>
                <p className="mt-1">Documento gerado digitalmente • Validação online disponível 24/7</p>
              </div>
            </>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
