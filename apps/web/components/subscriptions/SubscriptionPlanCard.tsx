'use client'

import { useState } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import {
  useDeleteSubscriptionPlan,
  useActivateSubscriptionPlan,
  useDeactivateSubscriptionPlan,
} from '@/lib/api/subscription-plan-api'
import { toast } from 'sonner'
import { Check, Edit, Trash2, Hash, Calendar, Copy } from 'lucide-react'
import type { SubscriptionPlan } from '@plenya/types'
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

interface SubscriptionPlanCardProps {
  plan: SubscriptionPlan
  onEdit?: (plan: SubscriptionPlan) => void
  onDuplicate?: (plan: SubscriptionPlan) => void
  showActions?: boolean
}

const billingCycleLabels = {
  monthly: 'Mensal',
  quarterly: 'Trimestral',
  yearly: 'Anual',
  one_time: 'Pagamento Único',
}

export function SubscriptionPlanCard({ plan, onEdit, onDuplicate, showActions = true }: SubscriptionPlanCardProps) {
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false)

  const deleteMutation = useDeleteSubscriptionPlan()
  const activateMutation = useActivateSubscriptionPlan()
  const deactivateMutation = useDeactivateSubscriptionPlan()

  const handleDelete = async () => {
    try {
      await deleteMutation.mutateAsync(plan.id)
      toast.success('Plano excluído com sucesso')
      setDeleteDialogOpen(false)
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao excluir plano')
    }
  }

  const handleActivate = async () => {
    try {
      await activateMutation.mutateAsync(plan.id)
      toast.success('Plano ativado com sucesso')
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao ativar plano')
    }
  }

  const handleDeactivate = async () => {
    try {
      await deactivateMutation.mutateAsync(plan.id)
      toast.success('Plano desativado com sucesso')
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao desativar plano')
    }
  }

  const formatCurrency = (value: number) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: plan.currency || 'BRL',
    }).format(value)
  }

  // Parse features JSON safely
  let features: string[] = []
  try {
    if (plan.features) {
      const parsed = JSON.parse(plan.features)
      if (Array.isArray(parsed)) {
        features = parsed
      } else if (typeof parsed === 'object') {
        // If it's an object, convert to key: value strings
        features = Object.entries(parsed).map(([key, value]) => `${key}: ${value}`)
      }
    }
  } catch (error) {
    console.error('Error parsing features:', error)
  }

  const isSubmitting = deleteMutation.isPending || activateMutation.isPending || deactivateMutation.isPending

  return (
    <>
      <Card className="h-full">
        <CardHeader>
          <div className="flex items-start justify-between gap-4">
            <div className="flex-1">
              <div className="flex items-center gap-2 mb-2">
                <CardTitle className="text-xl">{plan.name}</CardTitle>
                <Badge variant={plan.isActive ? 'default' : 'secondary'}>
                  {plan.isActive ? 'Ativo' : 'Inativo'}
                </Badge>
                <Badge variant="outline" className="flex items-center gap-1">
                  <Hash className="h-3 w-3" />
                  {plan.order}
                </Badge>
              </div>
              {plan.description && (
                <CardDescription className="text-sm">{plan.description}</CardDescription>
              )}
            </div>
            {plan.method && (
              <Badge variant="outline" style={{ borderColor: plan.method.color || undefined }}>
                {plan.method.shortName}
              </Badge>
            )}
          </div>
        </CardHeader>

        <CardContent className="space-y-4">
          {/* Pricing */}
          <div className="flex items-baseline gap-2">
            <span className="text-2xl font-bold">{formatCurrency(plan.price)}</span>
            <span className="text-sm text-muted-foreground">/ {billingCycleLabels[plan.billingCycle]}</span>
          </div>

          {/* Trial Period */}
          {plan.trialPeriodDays > 0 && (
            <div className="flex items-center gap-2 text-sm text-blue-600">
              <Calendar className="h-4 w-4" />
              <span>{plan.trialPeriodDays} dias de teste grátis</span>
            </div>
          )}

          {/* Features */}
          {features.length > 0 && (
            <div className="pt-3 border-t">
              <div className="text-xs text-muted-foreground mb-2">Incluído no plano:</div>
              <div className="space-y-1.5">
                {features.map((feature, index) => (
                  <div key={index} className="flex items-start gap-2 text-sm">
                    <Check className="h-4 w-4 text-green-500 mt-0.5 flex-shrink-0" />
                    <span>{feature}</span>
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* Metadata */}
          <div className="pt-3 border-t text-xs text-muted-foreground">
            <div>Moeda: {plan.currency}</div>
            <div>Criado em: {new Date(plan.createdAt).toLocaleDateString('pt-BR')}</div>
          </div>

          {/* Actions */}
          {showActions && (
            <div className="space-y-2 pt-3 border-t">
              <div className="flex gap-2">
                {onEdit && (
                  <Button variant="outline" size="sm" onClick={() => onEdit(plan)} className="flex-1">
                    <Edit className="h-4 w-4 mr-1" />
                    Editar
                  </Button>
                )}

                {onDuplicate && (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => onDuplicate(plan)}
                    className="flex-1 hover:bg-blue-500 hover:text-white"
                  >
                    <Copy className="h-4 w-4 mr-1" />
                    Duplicar
                  </Button>
                )}
              </div>

              <div className="flex gap-2">
                {plan.isActive ? (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={handleDeactivate}
                    className="flex-1 hover:bg-orange-500 hover:text-white"
                    disabled={isSubmitting}
                  >
                    Desativar
                  </Button>
                ) : (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={handleActivate}
                    className="flex-1 hover:bg-green-500 hover:text-white"
                    disabled={isSubmitting}
                  >
                    Ativar
                  </Button>
                )}

                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setDeleteDialogOpen(true)}
                  className="flex-1 hover:bg-destructive hover:text-destructive-foreground"
                  disabled={isSubmitting}
                >
                  <Trash2 className="h-4 w-4 mr-1" />
                  Excluir
                </Button>
              </div>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Delete Dialog */}
      <AlertDialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Excluir Plano</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir "{plan.name}"? Esta ação não pode ser desfeita.
              Subscriptions existentes não serão afetados, mas novos subscriptions não poderão usar este plano.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={handleDelete}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
              disabled={deleteMutation.isPending}
            >
              {deleteMutation.isPending ? 'Excluindo...' : 'Confirmar Exclusão'}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
