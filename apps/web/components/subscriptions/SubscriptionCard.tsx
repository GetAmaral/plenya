'use client'

import { useState } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { SubscriptionStatusBadge } from './SubscriptionStatusBadge'
import { useCancelSubscription, useRenewSubscription } from '@/lib/api/subscription-api'
import { toast } from 'sonner'
import { Calendar, DollarSign, RefreshCw, X, Edit, TrendingUp } from 'lucide-react'
import type { PatientSubscription } from '@plenya/types'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'
import { Textarea } from '@/components/ui/textarea'
import { Label } from '@/components/ui/label'

interface SubscriptionCardProps {
  subscription: PatientSubscription
  onEdit?: () => void
  showActions?: boolean
}

interface PlanSnapshot {
  name: string
  price: number
  features: string
  billingCycle: string
  currency: string
}

const billingCycleLabels = {
  monthly: 'Mensal',
  quarterly: 'Trimestral',
  yearly: 'Anual',
  one_time: 'Pagamento Único',
}

const parsePlanSnapshot = (snapshot: string): PlanSnapshot | null => {
  try {
    return JSON.parse(snapshot)
  } catch {
    return null
  }
}

export function SubscriptionCard({ subscription, onEdit, showActions = true }: SubscriptionCardProps) {
  const [cancelDialogOpen, setCancelDialogOpen] = useState(false)
  const [renewDialogOpen, setRenewDialogOpen] = useState(false)
  const [cancellationReason, setCancellationReason] = useState('')

  const cancelMutation = useCancelSubscription()
  const renewMutation = useRenewSubscription()

  const handleCancel = async () => {
    try {
      await cancelMutation.mutateAsync({
        id: subscription.id,
        data: { cancellationReason: cancellationReason || undefined },
      })
      toast.success('Subscription cancelado com sucesso')
      setCancelDialogOpen(false)
      setCancellationReason('')
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao cancelar subscription')
    }
  }

  const handleRenew = async () => {
    try {
      await renewMutation.mutateAsync({ id: subscription.id })
      toast.success('Subscription renovado com sucesso')
      setRenewDialogOpen(false)
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao renovar subscription')
    }
  }

  const formatDate = (date: string | undefined) => {
    if (!date) return '-'
    return new Date(date).toLocaleDateString('pt-BR')
  }

  // Parse plan snapshot to get original plan data
  const planData = parsePlanSnapshot(subscription.planSnapshot)

  const formatCurrency = (value: number, currency?: string) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: currency || planData?.currency || 'BRL',
    }).format(value)
  }

  // Calculate final price: custom price takes precedence, then apply discount to snapshot price
  const snapshotPrice = planData?.price || 0
  const finalPrice = subscription.customPrice
    ? subscription.customPrice
    : snapshotPrice * (1 - subscription.discountPercent / 100)

  // Has discount if custom price is set or discount percent > 0
  const hasDiscount = subscription.customPrice !== undefined || subscription.discountPercent > 0

  // Parse features from snapshot
  const features = planData?.features ? JSON.parse(planData.features) : null

  return (
    <>
      <Card className="h-full">
        <CardHeader>
          <div className="flex items-start justify-between gap-4">
            <div className="flex-1">
              <div className="flex items-center gap-2 mb-2">
                <CardTitle className="text-xl">
                  {planData?.name || 'Plano Indisponível'}
                </CardTitle>
                <SubscriptionStatusBadge status={subscription.status} />
              </div>
              {subscription.patient?.name && (
                <CardDescription className="text-sm font-medium">
                  Paciente: {subscription.patient.name}
                </CardDescription>
              )}
              {planData?.name && (
                <CardDescription className="text-sm text-muted-foreground">
                  Plano: {planData.name}
                </CardDescription>
              )}
            </div>
            {subscription.subscriptionPlan?.method && (
              <Badge variant="outline" style={{ borderColor: subscription.subscriptionPlan.method.color || undefined }}>
                {subscription.subscriptionPlan.method.shortName}
              </Badge>
            )}
          </div>
        </CardHeader>

        <CardContent className="space-y-4">
          {/* Pricing */}
          <div className="flex items-baseline gap-2">
            {hasDiscount && !subscription.customPrice && (
              <span className="text-sm text-muted-foreground line-through">
                {formatCurrency(snapshotPrice, planData?.currency)}
              </span>
            )}
            <span className="text-2xl font-bold">{formatCurrency(finalPrice, planData?.currency)}</span>
            {planData?.billingCycle && (
              <span className="text-sm text-muted-foreground">
                / {billingCycleLabels[planData.billingCycle as keyof typeof billingCycleLabels] || planData.billingCycle}
              </span>
            )}
          </div>

          {hasDiscount && (
            <div className="flex items-center gap-2 text-sm text-green-600">
              <TrendingUp className="h-4 w-4" />
              <span>
                {subscription.customPrice !== undefined
                  ? `Preço customizado: ${formatCurrency(subscription.customPrice, planData?.currency)}`
                  : `${subscription.discountPercent}% de desconto`}
                {subscription.discountReason && ` - ${subscription.discountReason}`}
              </span>
            </div>
          )}

          {/* Dates */}
          <div className="grid grid-cols-2 gap-3 text-sm">
            <div className="flex items-center gap-2">
              <Calendar className="h-4 w-4 text-muted-foreground" />
              <div>
                <div className="text-xs text-muted-foreground">Início</div>
                <div className="font-medium">{formatDate(subscription.startDate)}</div>
              </div>
            </div>

            {subscription.trialEndDate && (
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-blue-500" />
                <div>
                  <div className="text-xs text-muted-foreground">Fim Trial</div>
                  <div className="font-medium text-blue-600">{formatDate(subscription.trialEndDate)}</div>
                </div>
              </div>
            )}

            {subscription.endDate && (
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-muted-foreground" />
                <div>
                  <div className="text-xs text-muted-foreground">Término</div>
                  <div className="font-medium">{formatDate(subscription.endDate)}</div>
                </div>
              </div>
            )}

            {subscription.nextBillingDate && (
              <div className="flex items-center gap-2">
                <DollarSign className="h-4 w-4 text-green-500" />
                <div>
                  <div className="text-xs text-muted-foreground">Próx. Cobrança</div>
                  <div className="font-medium text-green-600">{formatDate(subscription.nextBillingDate)}</div>
                </div>
              </div>
            )}
          </div>

          {/* Features */}
          {features && (
            <div className="pt-3 border-t">
              <div className="text-xs text-muted-foreground mb-2">Incluído no plano:</div>
              <div className="flex flex-wrap gap-2">
                {Object.entries(features).map(([key, value]) => (
                  <Badge key={key} variant="secondary" className="text-xs">
                    {key}: {String(value)}
                  </Badge>
                ))}
              </div>
            </div>
          )}

          {/* Auto Renew */}
          {subscription.autoRenew && subscription.status === 'active' && (
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <RefreshCw className="h-4 w-4" />
              <span>Renovação automática habilitada</span>
            </div>
          )}

          {/* Renewal Count */}
          {subscription.renewalCount > 0 && (
            <div className="text-xs text-muted-foreground">
              Renovações: {subscription.renewalCount}
            </div>
          )}

          {/* Actions */}
          {showActions && (
            <div className="flex gap-2 pt-3 border-t">
              {onEdit && (
                <Button variant="outline" size="sm" onClick={onEdit} className="flex-1">
                  <Edit className="h-4 w-4 mr-1" />
                  Editar
                </Button>
              )}

              {subscription.status === 'active' && (
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setCancelDialogOpen(true)}
                  className="flex-1 hover:bg-destructive hover:text-destructive-foreground"
                >
                  <X className="h-4 w-4 mr-1" />
                  Cancelar
                </Button>
              )}

              {(subscription.status === 'expired' || subscription.status === 'cancelled') && (
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setRenewDialogOpen(true)}
                  className="flex-1 hover:bg-green-500 hover:text-white"
                >
                  <RefreshCw className="h-4 w-4 mr-1" />
                  Renovar
                </Button>
              )}
            </div>
          )}

          {/* Notes */}
          {subscription.notes && (
            <div className="pt-3 border-t">
              <div className="text-xs text-muted-foreground mb-1">Observações:</div>
              <p className="text-sm">{subscription.notes}</p>
            </div>
          )}

          {/* Cancellation Reason */}
          {subscription.cancellationReason && (
            <div className="pt-3 border-t">
              <div className="text-xs text-muted-foreground mb-1">Motivo do cancelamento:</div>
              <p className="text-sm text-red-600">{subscription.cancellationReason}</p>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Cancel Dialog */}
      <AlertDialog open={cancelDialogOpen} onOpenChange={setCancelDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Cancelar Subscription</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja cancelar "{planData?.name || 'este plano'}"? Esta ação não pode ser desfeita facilmente.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <div className="space-y-2 py-4">
            <Label htmlFor="cancellationReason">Motivo do cancelamento (opcional)</Label>
            <Textarea
              id="cancellationReason"
              placeholder="Ex: Questões financeiras, mudança de plano..."
              value={cancellationReason}
              onChange={(e) => setCancellationReason(e.target.value)}
              rows={3}
            />
          </div>
          <AlertDialogFooter>
            <AlertDialogCancel>Voltar</AlertDialogCancel>
            <AlertDialogAction
              onClick={handleCancel}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
              disabled={cancelMutation.isPending}
            >
              {cancelMutation.isPending ? 'Cancelando...' : 'Confirmar Cancelamento'}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      {/* Renew Dialog */}
      <AlertDialog open={renewDialogOpen} onOpenChange={setRenewDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Renovar Subscription</AlertDialogTitle>
            <AlertDialogDescription>
              Deseja renovar o subscription "{planData?.name || 'este plano'}"? O status será alterado para ativo.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={handleRenew}
              className="bg-green-500 hover:bg-green-600"
              disabled={renewMutation.isPending}
            >
              {renewMutation.isPending ? 'Renovando...' : 'Confirmar Renovação'}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
