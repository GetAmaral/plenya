'use client'

import { useEffect, useRef, useState } from 'react'
import { useForm } from 'react-hook-form'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { toast } from 'sonner'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Checkbox } from '@/components/ui/checkbox'
import { Card, CardContent } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { PatientSubscription, SubscriptionPlan } from '@plenya/types'
import {
  CreateSubscriptionDTO,
  UpdateSubscriptionDTO,
  useCreateSubscription,
  useUpdateSubscription,
} from '@/lib/api/subscription-api'
import { useActiveSubscriptionPlans } from '@/lib/api/subscription-plan-api'

interface SubscriptionDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  subscription?: PatientSubscription
  patientId?: string
}

interface SubscriptionFormData {
  subscriptionPlanId: string
  status: 'active' | 'inactive' | 'cancelled' | 'expired' | 'suspended' | 'trial'
  autoRenew: boolean
  startDate: string
  endDate?: string
  trialEndDate?: string
  discountPercent: number
  discountReason?: string
  customPrice?: number
  customTrialDays?: number
  notes?: string
}

const billingCycleLabels = {
  monthly: 'Mensal',
  quarterly: 'Trimestral',
  yearly: 'Anual',
  one_time: 'Pagamento Único',
}

// Helper to convert date to YYYY-MM-DD format
const formatDateForInput = (date: string | undefined | null): string => {
  if (!date) return ''
  try {
    const d = new Date(date)
    return d.toISOString().split('T')[0]
  } catch {
    return ''
  }
}

export function SubscriptionDialog({
  open,
  onOpenChange,
  subscription,
  patientId,
}: SubscriptionDialogProps) {
  const isEditing = !!subscription

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const { data: plans, isLoading: loadingPlans } = useActiveSubscriptionPlans()
  const createMutation = useCreateSubscription()
  const updateMutation = useUpdateSubscription()

  const [selectedPlan, setSelectedPlan] = useState<SubscriptionPlan | undefined>(undefined)

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<SubscriptionFormData>({
    defaultValues: {
      subscriptionPlanId: subscription?.subscriptionPlanId || '',
      status: subscription?.status || 'active',
      autoRenew: subscription?.autoRenew ?? true,
      startDate: formatDateForInput(subscription?.startDate) || new Date().toISOString().split('T')[0],
      endDate: formatDateForInput(subscription?.endDate),
      trialEndDate: formatDateForInput(subscription?.trialEndDate),
      discountPercent: subscription?.discountPercent || 0,
      discountReason: subscription?.discountReason || '',
      customPrice: subscription?.customPrice || undefined,
      customTrialDays: subscription?.customTrialDays || undefined,
      notes: subscription?.notes || '',
    },
  })

  const autoRenew = watch('autoRenew')
  const status = watch('status')
  const subscriptionPlanId = watch('subscriptionPlanId')
  const discountPercent = watch('discountPercent')
  const customPrice = watch('customPrice')

  // Update selected plan when plan ID changes
  useEffect(() => {
    if (subscriptionPlanId && plans) {
      const plan = plans.find((p) => p.id === subscriptionPlanId)
      setSelectedPlan(plan)
    } else {
      setSelectedPlan(undefined)
    }
  }, [subscriptionPlanId, plans])

  // Reset form when subscription changes
  useEffect(() => {
    if (subscription) {
      reset({
        subscriptionPlanId: subscription.subscriptionPlanId,
        status: subscription.status,
        autoRenew: subscription.autoRenew,
        startDate: formatDateForInput(subscription.startDate),
        endDate: formatDateForInput(subscription.endDate),
        trialEndDate: formatDateForInput(subscription.trialEndDate),
        discountPercent: subscription.discountPercent,
        discountReason: subscription.discountReason || '',
        customPrice: subscription.customPrice || undefined,
        customTrialDays: subscription.customTrialDays || undefined,
        notes: subscription.notes || '',
      })
    } else {
      reset({
        subscriptionPlanId: '',
        status: 'active',
        autoRenew: true,
        startDate: new Date().toISOString().split('T')[0],
        endDate: '',
        trialEndDate: '',
        discountPercent: 0,
        discountReason: '',
        customPrice: undefined,
        customTrialDays: undefined,
        notes: '',
      })
    }
  }, [subscription, reset])

  const onSubmit = async (data: SubscriptionFormData) => {
    try {
      // Validate required fields
      if (!data.subscriptionPlanId) {
        toast.error('Selecione um plano de subscription')
        return
      }

      // Convert empty strings to undefined
      const payload = {
        ...data,
        endDate: data.endDate || undefined,
        trialEndDate: data.trialEndDate || undefined,
        discountReason: data.discountReason || undefined,
        customPrice: data.customPrice || undefined,
        customTrialDays: data.customTrialDays || undefined,
        notes: data.notes || undefined,
      }

      if (isEditing) {
        await updateMutation.mutateAsync({
          id: subscription.id,
          data: payload as UpdateSubscriptionDTO,
        })
        toast.success('Subscription atualizado com sucesso')
      } else {
        if (!patientId) {
          toast.error('ID do paciente não fornecido')
          return
        }
        await createMutation.mutateAsync({
          patientId,
          data: { ...payload, patientId } as CreateSubscriptionDTO,
        })
        toast.success('Subscription criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      const errorMessage =
        error?.message ||
        (isEditing ? 'Erro ao atualizar subscription' : 'Erro ao criar subscription')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createMutation.isPending || updateMutation.isPending

  const formatCurrency = (value: number, currency: string = 'BRL') => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency,
    }).format(value)
  }

  // Calculate final price preview
  const calculateFinalPrice = () => {
    if (!selectedPlan) return null
    if (customPrice) return customPrice
    return selectedPlan.price * (1 - (discountPercent || 0) / 100)
  }

  const finalPrice = calculateFinalPrice()

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Subscription' : 'Novo Subscription'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do subscription'
              : 'Crie um novo pacote de acompanhamento para o paciente'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Plan Selector */}
          <div className="space-y-2">
            <Label htmlFor="subscriptionPlanId">Plano de Subscription *</Label>
            <Select
              value={subscriptionPlanId}
              onValueChange={(value) => setValue('subscriptionPlanId', value)}
              disabled={loadingPlans}
            >
              <SelectTrigger>
                <SelectValue placeholder={loadingPlans ? 'Carregando planos...' : 'Selecione um plano'} />
              </SelectTrigger>
              <SelectContent>
                {plans?.map((plan) => (
                  <SelectItem key={plan.id} value={plan.id}>
                    {plan.name} - {formatCurrency(plan.price, plan.currency)} /{' '}
                    {billingCycleLabels[plan.billingCycle]}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {errors.subscriptionPlanId && (
              <p className="text-sm text-destructive">{errors.subscriptionPlanId.message}</p>
            )}
          </div>

          {/* Plan Preview */}
          {selectedPlan && (
            <Card className="bg-muted/50">
              <CardContent className="pt-4 space-y-2">
                <div className="flex items-center justify-between">
                  <h4 className="font-medium">{selectedPlan.name}</h4>
                  {selectedPlan.method && (
                    <Badge variant="outline" style={{ borderColor: selectedPlan.method.color || undefined }}>
                      {selectedPlan.method.shortName}
                    </Badge>
                  )}
                </div>
                <p className="text-sm text-muted-foreground">{selectedPlan.description}</p>
                <div className="grid grid-cols-2 gap-2 text-xs">
                  <div>
                    <span className="text-muted-foreground">Preço base:</span>{' '}
                    <span className="font-medium">{formatCurrency(selectedPlan.price, selectedPlan.currency)}</span>
                  </div>
                  <div>
                    <span className="text-muted-foreground">Trial:</span>{' '}
                    <span className="font-medium">{selectedPlan.trialPeriodDays} dias</span>
                  </div>
                </div>
                {finalPrice !== null && finalPrice !== selectedPlan.price && (
                  <div className="pt-2 border-t">
                    <div className="flex items-baseline gap-2">
                      <span className="text-sm text-muted-foreground">Preço final:</span>
                      <span className="text-lg font-bold text-green-600">
                        {formatCurrency(finalPrice, selectedPlan.currency)}
                      </span>
                    </div>
                  </div>
                )}
              </CardContent>
            </Card>
          )}

          {/* Status */}
          <div className="space-y-2">
            <Label htmlFor="status">Status *</Label>
            <Select value={status} onValueChange={(value) => setValue('status', value as any)}>
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="active">Ativo</SelectItem>
                <SelectItem value="trial">Trial</SelectItem>
                <SelectItem value="inactive">Inativo</SelectItem>
                <SelectItem value="suspended">Suspenso</SelectItem>
                <SelectItem value="cancelled">Cancelado</SelectItem>
                <SelectItem value="expired">Expirado</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* Dates */}
          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <Label htmlFor="startDate">Data de Início *</Label>
              <Input
                id="startDate"
                type="date"
                {...register('startDate', { required: 'Data de início é obrigatória' })}
              />
              {errors.startDate && (
                <p className="text-sm text-destructive">{errors.startDate.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="endDate">Data de Término</Label>
              <Input id="endDate" type="date" {...register('endDate')} />
            </div>

            <div className="space-y-2">
              <Label htmlFor="trialEndDate">Fim do Trial</Label>
              <Input id="trialEndDate" type="date" {...register('trialEndDate')} />
            </div>
          </div>

          {/* Auto Renew */}
          <div className="flex items-center space-x-2 p-4 bg-muted/50 rounded-lg">
            <Checkbox
              id="autoRenew"
              checked={autoRenew}
              onCheckedChange={(checked) => setValue('autoRenew', checked as boolean)}
            />
            <div className="grid gap-1.5 leading-none">
              <Label
                htmlFor="autoRenew"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 cursor-pointer"
              >
                Renovação Automática
              </Label>
              <p className="text-sm text-muted-foreground">
                Renova automaticamente no final do período
              </p>
            </div>
          </div>

          {/* Customization Section */}
          <div className="space-y-4 p-4 border rounded-lg">
            <h3 className="font-medium text-sm">Customizações Opcionais</h3>

            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="discountPercent">Desconto (%)</Label>
                <Input
                  id="discountPercent"
                  type="number"
                  step="0.01"
                  min="0"
                  max="100"
                  placeholder="10"
                  {...register('discountPercent', {
                    valueAsNumber: true,
                    min: { value: 0, message: 'Mínimo 0%' },
                    max: { value: 100, message: 'Máximo 100%' },
                  })}
                />
                {errors.discountPercent && (
                  <p className="text-sm text-destructive">{errors.discountPercent.message}</p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="customPrice">Preço Customizado</Label>
                <Input
                  id="customPrice"
                  type="number"
                  step="0.01"
                  min="0"
                  placeholder="Deixe vazio para usar preço do plano"
                  {...register('customPrice', {
                    valueAsNumber: true,
                    min: { value: 0, message: 'Preço deve ser maior ou igual a 0' },
                  })}
                />
                {errors.customPrice && (
                  <p className="text-sm text-destructive">{errors.customPrice.message}</p>
                )}
                <p className="text-xs text-muted-foreground">
                  Sobrescreve o preço do plano (ignorando desconto)
                </p>
              </div>
            </div>

            {(discountPercent > 0 || customPrice) && (
              <div className="space-y-2">
                <Label htmlFor="discountReason">Motivo do Desconto/Customização</Label>
                <Textarea
                  id="discountReason"
                  placeholder="Ex: Promoção de lançamento, acordo especial..."
                  rows={2}
                  {...register('discountReason')}
                />
              </div>
            )}

            <div className="space-y-2">
              <Label htmlFor="customTrialDays">Período de Trial Customizado (dias)</Label>
              <Input
                id="customTrialDays"
                type="number"
                min="0"
                placeholder={`Padrão: ${selectedPlan?.trialPeriodDays || 0} dias`}
                {...register('customTrialDays', {
                  valueAsNumber: true,
                  min: { value: 0, message: 'Mínimo 0 dias' },
                })}
              />
              {errors.customTrialDays && (
                <p className="text-sm text-destructive">{errors.customTrialDays.message}</p>
              )}
            </div>
          </div>

          {/* Notes */}
          <div className="space-y-2">
            <Label htmlFor="notes">Observações Internas</Label>
            <Textarea
              id="notes"
              placeholder="Anotações para uso interno..."
              rows={2}
              {...register('notes')}
            />
          </div>

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={isSubmitting}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={isSubmitting}>
              {isSubmitting
                ? isEditing
                  ? 'Salvando...'
                  : 'Criando...'
                : isEditing
                ? 'Salvar'
                : 'Criar'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
