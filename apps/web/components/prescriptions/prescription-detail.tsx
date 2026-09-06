'use client'

import { useRouter } from 'next/navigation'
import { Copy, FileCheck, FileText } from 'lucide-react'
import { toast } from 'sonner'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { SendDocumentEmailButton } from '@/components/clinical/send-document-email-button'
import { SendDocumentWhatsAppButton } from '@/components/clinical/send-document-whatsapp-button'
import { formatDate, formatDateOnly, formatDateTime } from '@/lib/format-date'
import {
  openPrescriptionDraftPdf,
  openPrescriptionPdf,
  type Prescription,
} from '@/lib/api/prescriptions'
import { CATEGORY_OPTIONS } from '@/lib/validations/prescription'

const categoryLabel = (category?: string) =>
  CATEGORY_OPTIONS.find((c) => c.value === category)?.label ?? category ?? '-'

/** Número em pt-BR sem zeros à toa: 0.25 → "0,25", 300 → "300". Mesma regra do PDF. */
const formatQuantity = (value: number) => String(value).replace('.', ',')

const componentQuantity = (quantity: number, unit: string) =>
  unit === '%' ? `${formatQuantity(quantity)}%` : `${formatQuantity(quantity)} ${unit}`

/**
 * Visualização somente leitura de uma prescrição.
 *
 * O corpo disto vivia dentro de `[id]/edit/page.tsx`, no ramo "já assinada" — e por isso a rota
 * `/prescriptions/[id]`, linkada pelo olho da listagem, dava 404. Agora a tela de visualização é
 * a dona deste conteúdo e a edição só cuida de rascunho.
 */
export function PrescriptionDetail({ prescription }: { prescription: Prescription }) {
  const router = useRouter()
  const isSigned = !!prescription.signedPdfPath
  const isCompounded = prescription.type === 'compounded'

  return (
    <>
      {isSigned ? (
        <Alert className="mb-6 border-ocean-500 bg-ocean-50 dark:bg-petrol/30">
          <FileCheck className="h-4 w-4 text-ocean-700" />
          <AlertTitle className="text-ocean-700">
            {prescription.signatureMode === 'manual'
              ? 'Receita emitida para assinatura em papel'
              : 'Prescrição com assinatura digital'}
          </AlertTitle>
          <AlertDescription className="text-ocean-700">
            Prescrição emitida não pode ser alterada. Para mudar algo, duplique e emita uma nova.
          </AlertDescription>
        </Alert>
      ) : (
        <Alert className="mb-6">
          <FileText className="h-4 w-4" />
          <AlertTitle>Rascunho</AlertTitle>
          <AlertDescription>
            Esta prescrição ainda não foi assinada e não vale como receita.
          </AlertDescription>
        </Alert>
      )}

      {isCompounded && (
        <div className="mb-6 space-y-4">
          {(prescription.formulas ?? []).map((formula, index) => (
            <Card key={formula.id}>
              <CardHeader>
                <div className="flex flex-wrap items-center justify-between gap-2">
                  <CardTitle>
                    {formula.name || `Fórmula ${index + 1}`}{' '}
                    <span className="font-normal text-muted-foreground">
                      {formula.pharmaceuticalForm}
                    </span>
                  </CardTitle>
                  <div className="flex gap-2">
                    <Badge variant="outline">
                      {formula.usageType === 'external' ? 'Uso externo' : 'Uso interno'}
                    </Badge>
                    <Badge variant="outline">{categoryLabel(formula.highestCategory)}</Badge>
                  </div>
                </div>
              </CardHeader>
              <CardContent className="space-y-3">
                <ul className="space-y-1">
                  {formula.components.map((c) => (
                    <li key={c.id} className="flex items-baseline justify-between gap-3 text-sm">
                      <span>
                        {c.substance}
                        {c.note ? (
                          <span className="ml-1 text-xs italic text-muted-foreground">
                            {c.note}
                          </span>
                        ) : null}
                      </span>
                      <span className="font-medium">{componentQuantity(c.quantity, c.unit)}</span>
                    </li>
                  ))}
                  {formula.vehicle && (
                    <li className="text-sm text-muted-foreground">{formula.vehicle}</li>
                  )}
                </ul>

                <div className="grid grid-cols-1 gap-3 border-t pt-3 sm:grid-cols-2">
                  <div>
                    <Label className="text-muted-foreground">Aviar</Label>
                    <p className="font-medium">
                      {formatQuantity(formula.quantityToDispense)} {formula.quantityUnit}
                      {formula.quantityInWords ? ` (${formula.quantityInWords})` : ''}
                    </p>
                  </div>
                  <div>
                    <Label className="text-muted-foreground">Posologia</Label>
                    <p className="font-medium">
                      {formula.posology}
                      {formula.duration
                        ? ` · por ${formula.duration} ${formula.duration === 1 ? 'dia' : 'dias'}`
                        : ' · uso contínuo'}
                    </p>
                  </div>
                  {formula.instructions && (
                    <div className="sm:col-span-2">
                      <Label className="text-muted-foreground">Orientações</Label>
                      <p className="font-medium">{formula.instructions}</p>
                    </div>
                  )}
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      <div className={isCompounded ? 'hidden' : 'mb-6 space-y-4'}>
        {prescription.medications.map((med, index) => (
          <Card key={med.id}>
            <CardHeader>
              <div className="flex items-center justify-between gap-2">
                <CardTitle>Medicamento {index + 1}</CardTitle>
                <Badge variant="outline">{categoryLabel(med.category)}</Badge>
              </div>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div>
                  <Label className="text-muted-foreground">Nome</Label>
                  <p className="font-medium">{med.medicationName}</p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Princípio ativo</Label>
                  <p className="font-medium">{med.activeIngredient}</p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Concentração</Label>
                  <p className="font-medium">{med.concentration || '-'}</p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Duração</Label>
                  {/* Duração em branco é uso contínuo, e a tela precisa DIZER isso: "0 dias" lia
                      como erro de digitação, e o silêncio não distingue esquecimento de decisão. */}
                  <p className="font-medium">
                    {med.duration ? `${med.duration} ${med.duration === 1 ? 'dia' : 'dias'}` : 'Uso contínuo'}
                  </p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Quantidade</Label>
                  <p className="font-medium">
                    {med.quantity ? (
                      <>
                        {med.quantity}
                        {med.quantityInWords ? ` (${med.quantityInWords})` : ''}
                      </>
                    ) : (
                      'Não especificada'
                    )}
                  </p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Posologia</Label>
                  <p className="font-medium">
                    {med.dosage} {med.frequency}
                    {med.route ? ` · via ${med.route}` : ''}
                  </p>
                </div>
                {med.instructions && (
                  <div className="sm:col-span-2">
                    <Label className="text-muted-foreground">Instruções</Label>
                    <p className="font-medium">{med.instructions}</p>
                  </div>
                )}
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {prescription.generalInstructions && (
        <Card className="mb-6">
          <CardHeader>
            <CardTitle>Instruções gerais</CardTitle>
          </CardHeader>
          <CardContent>
            <p>{prescription.generalInstructions}</p>
          </CardContent>
        </Card>
      )}

      <Card>
        <CardHeader>
          <CardTitle>{isSigned ? 'Informações da assinatura' : 'Situação'}</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div>
              <Label className="text-muted-foreground">Data da prescrição</Label>
              <p className="font-medium">{formatDate(prescription.prescriptionDate)}</p>
            </div>
            <div>
              <Label className="text-muted-foreground">Válida até</Label>
              {/* valid_until é DATE pura: formatDate a leria como meia-noite UTC e mostraria
                  o dia anterior em BRT. */}
              <p className="font-medium">{formatDateOnly(prescription.validUntil)}</p>
            </div>
            {prescription.signedAt && (
              <div>
                <Label className="text-muted-foreground">Assinada em</Label>
                <p className="font-medium">{formatDateTime(prescription.signedAt)}</p>
              </div>
            )}
            {prescription.sncrNumber && (
              <div>
                <Label className="text-muted-foreground">SNCR</Label>
                <p className="font-medium">{prescription.sncrNumber}</p>
              </div>
            )}
          </div>

          <div className="flex flex-col gap-3 pt-2 sm:flex-row">
            {isSigned && (
              <Button
                variant="outline"
                className="flex-1"
                onClick={() =>
                  openPrescriptionPdf(prescription.id).catch((e) =>
                    toast.error(e?.message ?? 'Falha ao baixar o PDF')
                  )
                }
              >
                <FileCheck className="mr-2 h-4 w-4" />
                Baixar PDF
              </Button>
            )}
            {/* Envio é ato explícito e por canal: assinar gera o PDF e publica no portal, e não
                dispara e-mail nem WhatsApp. Quem manda é quem clica, e cada canal tem seu botão. */}
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
            {/* Conferir ANTES de assinar. Sem isto a única forma de ver a receita era assiná-la,
                e assinatura é ato médico com certificado, que não se desfaz. O rascunho sai no
                layout de assinatura manual, com o campo em branco, para não passar por válido se
                o arquivo sair do EMR. */}
            {!isSigned && (
              <>
                <Button
                  variant="outline"
                  className="flex-1"
                  onClick={() =>
                    openPrescriptionDraftPdf(prescription.id).catch((e) =>
                      toast.error(e?.message ?? 'Falha ao gerar o rascunho')
                    )
                  }
                >
                  <FileText className="mr-2 h-4 w-4" />
                  Ver rascunho em PDF
                </Button>
                <Button
                  variant="outline"
                  className="flex-1"
                  onClick={() => router.push(`/prescriptions/${prescription.id}/edit`)}
                >
                  Editar rascunho
                </Button>
              </>
            )}
            <Button
              className="flex-1"
              onClick={() => router.push(`/prescriptions/new?from=${prescription.id}`)}
            >
              <Copy className="mr-2 h-4 w-4" />
              Repetir receita
            </Button>
          </div>

          <Button
            variant="ghost"
            className="w-full"
            onClick={() => router.push('/prescriptions')}
          >
            Voltar para prescrições
          </Button>
        </CardContent>
      </Card>
    </>
  )
}
