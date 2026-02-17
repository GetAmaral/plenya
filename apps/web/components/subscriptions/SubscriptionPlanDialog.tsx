'use client'

import { useEffect, useRef } from 'react'
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
import type { SubscriptionPlan } from '@plenya/types'
import {
  CreateSubscriptionPlanDTO,
  UpdateSubscriptionPlanDTO,
  useCreateSubscriptionPlan,
  useUpdateSubscriptionPlan,
} from '@/lib/api/subscription-plan-api'
import { useAllMethods } from '@/lib/api/method-api'

interface SubscriptionPlanDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  plan?: SubscriptionPlan
  initialData?: Partial<FormData>
}

interface FormData {
  name: string
  description: string
  features: string
  price: number
  currency: string
  billingCycle: 'monthly' | 'quarterly' | 'yearly' | 'one_time'
  methodId?: string
  trialPeriodDays: number
  order: number
  isActive: boolean
}

export function SubscriptionPlanDialog({
  open,
  onOpenChange,
  plan,
  initialData,
}: SubscriptionPlanDialogProps) {
  const isEditing = !!plan

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const { data: methods } = useAllMethods()
  const createMutation = useCreateSubscriptionPlan()
  const updateMutation = useUpdateSubscriptionPlan()

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<FormData>({
    defaultValues: {
      name: plan?.name || initialData?.name || '',
      description: plan?.description || initialData?.description || '',
      features: plan?.features || initialData?.features || '',
      price: plan?.price || initialData?.price || 0,
      currency: plan?.currency || initialData?.currency || 'BRL',
      billingCycle: plan?.billingCycle || initialData?.billingCycle || 'monthly',
      methodId: plan?.methodId || initialData?.methodId || undefined,
      trialPeriodDays: plan?.trialPeriodDays || initialData?.trialPeriodDays || 0,
      order: plan?.order || initialData?.order || 0,
      isActive: plan?.isActive ?? initialData?.isActive ?? true,
    },
  })

  const billingCycle = watch('billingCycle')
  const methodId = watch('methodId')
  const isActive = watch('isActive')

  // Reset form when plan or initialData changes
  useEffect(() => {
    if (plan) {
      reset({
        name: plan.name,
        description: plan.description || '',
        features: plan.features || '',
        price: plan.price,
        currency: plan.currency,
        billingCycle: plan.billingCycle,
        methodId: plan.methodId || undefined,
        trialPeriodDays: plan.trialPeriodDays,
        order: plan.order,
        isActive: plan.isActive,
      })
    } else {
      reset({
        name: initialData?.name || '',
        description: initialData?.description || '',
        features: initialData?.features || '',
        price: initialData?.price || 0,
        currency: initialData?.currency || 'BRL',
        billingCycle: initialData?.billingCycle || 'monthly',
        methodId: initialData?.methodId || undefined,
        trialPeriodDays: initialData?.trialPeriodDays || 0,
        order: initialData?.order || 0,
        isActive: initialData?.isActive ?? true,
      })
    }
  }, [plan, initialData, reset])

  const onSubmit = async (data: FormData) => {
    try {
      // Validate features JSON
      if (data.features) {
        try {
          JSON.parse(data.features)
        } catch {
          toast.error('Features deve ser um JSON válido')
          return
        }
      }

      // Prepare payload
      const payload = {
        name: data.name,
        description: data.description || '',
        features: data.features || '[]',
        price: data.price,
        currency: data.currency,
        billingCycle: data.billingCycle,
        methodId: data.methodId || undefined,
        trialPeriodDays: data.trialPeriodDays,
        order: data.order,
      }

      if (isEditing) {
        await updateMutation.mutateAsync({
          id: plan.id,
          data: payload as UpdateSubscriptionPlanDTO,
        })
        toast.success('Plano atualizado com sucesso')
      } else {
        await createMutation.mutateAsync(payload as CreateSubscriptionPlanDTO)
        toast.success('Plano criado com sucesso')

        // If created as active, show additional success message
        if (data.isActive) {
          toast.info('O plano foi criado como ativo e já está disponível para uso')
        }
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      const errorMessage =
        error?.message ||
        (isEditing ? 'Erro ao atualizar plano' : 'Erro ao criar plano')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createMutation.isPending || updateMutation.isPending

  // Generate sample features JSON
  const getSampleFeatures = () => {
    return JSON.stringify([
      'Consultas ilimitadas',
      'Acesso ao app mobile',
      'Relatórios mensais',
      'Suporte prioritário',
    ], null, 2)
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Plano de Subscription' : 'Novo Plano de Subscription'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do plano'
              : 'Crie um novo plano de subscription que pacientes podem assinar'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Basic Info */}
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="name">Nome do Plano *</Label>
              <Input
                id="name"
                placeholder="Ex: Plano Premium AGIR"
                {...register('name', {
                  required: 'Nome do plano é obrigatório',
                  minLength: { value: 3, message: 'Mínimo 3 caracteres' },
                  maxLength: { value: 200, message: 'Máximo 200 caracteres' },
                })}
              />
              {errors.name && (
                <p className="text-sm text-destructive">{errors.name.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="methodId">Metodologia (opcional)</Label>
              <Select value={methodId || 'none'} onValueChange={(value) => setValue('methodId', value === 'none' ? undefined : value)}>
                <SelectTrigger>
                  <SelectValue placeholder="Selecione uma metodologia" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="none">Sem metodologia específica</SelectItem>
                  {methods?.map((method) => (
                    <SelectItem key={method.id} value={method.id}>
                      {method.shortName} - {method.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <p className="text-xs text-muted-foreground">
                Vincule o plano a uma metodologia específica (opcional)
              </p>
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descrição *</Label>
            <Textarea
              id="description"
              placeholder="Descreva o que diferencia este plano..."
              rows={2}
              {...register('description', {
                required: 'Descrição é obrigatória',
                minLength: { value: 10, message: 'Mínimo 10 caracteres' },
              })}
            />
            {errors.description && (
              <p className="text-sm text-destructive">{errors.description.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label htmlFor="features">Features (JSON) *</Label>
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={() => setValue('features', getSampleFeatures())}
                className="text-xs"
              >
                Usar exemplo
              </Button>
            </div>
            <Textarea
              id="features"
              placeholder='["Feature 1", "Feature 2", "Feature 3"]'
              rows={4}
              {...register('features', {
                required: 'Features são obrigatórias',
              })}
            />
            {errors.features && (
              <p className="text-sm text-destructive">{errors.features.message}</p>
            )}
            <p className="text-xs text-muted-foreground">
              Array JSON com as features do plano. Ex: ["Feature 1", "Feature 2"]
            </p>
          </div>

          {/* Pricing */}
          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <Label htmlFor="price">Preço *</Label>
              <Input
                id="price"
                type="number"
                step="0.01"
                min="0"
                placeholder="299.90"
                {...register('price', {
                  required: 'Preço é obrigatório',
                  valueAsNumber: true,
                  min: { value: 0, message: 'Preço deve ser maior ou igual a 0' },
                })}
              />
              {errors.price && (
                <p className="text-sm text-destructive">{errors.price.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="currency">Moeda *</Label>
              <Select value={watch('currency')} onValueChange={(value) => setValue('currency', value)}>
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="BRL">BRL (Real)</SelectItem>
                  <SelectItem value="USD">USD (Dólar)</SelectItem>
                  <SelectItem value="EUR">EUR (Euro)</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="space-y-2">
              <Label htmlFor="billingCycle">Ciclo de Cobrança *</Label>
              <Select value={billingCycle} onValueChange={(value) => setValue('billingCycle', value as any)}>
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="monthly">Mensal</SelectItem>
                  <SelectItem value="quarterly">Trimestral</SelectItem>
                  <SelectItem value="yearly">Anual</SelectItem>
                  <SelectItem value="one_time">Pagamento Único</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          {/* Additional Settings */}
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="trialPeriodDays">Período de Trial (dias)</Label>
              <Input
                id="trialPeriodDays"
                type="number"
                min="0"
                placeholder="7"
                {...register('trialPeriodDays', {
                  valueAsNumber: true,
                  min: { value: 0, message: 'Mínimo 0 dias' },
                })}
              />
              {errors.trialPeriodDays && (
                <p className="text-sm text-destructive">{errors.trialPeriodDays.message}</p>
              )}
              <p className="text-xs text-muted-foreground">
                Número de dias de teste grátis (0 = sem trial)
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="order">Ordem de Exibição</Label>
              <Input
                id="order"
                type="number"
                min="0"
                placeholder="0"
                {...register('order', {
                  valueAsNumber: true,
                  min: { value: 0, message: 'Mínimo 0' },
                })}
              />
              {errors.order && (
                <p className="text-sm text-destructive">{errors.order.message}</p>
              )}
              <p className="text-xs text-muted-foreground">
                Ordem de exibição na lista (menor primeiro)
              </p>
            </div>
          </div>

          {/* Is Active */}
          <div className="flex items-center space-x-2 p-4 bg-muted/50 rounded-lg">
            <Checkbox
              id="isActive"
              checked={isActive}
              onCheckedChange={(checked) => setValue('isActive', checked as boolean)}
            />
            <div className="grid gap-1.5 leading-none">
              <Label
                htmlFor="isActive"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 cursor-pointer"
              >
                Plano Ativo
              </Label>
              <p className="text-sm text-muted-foreground">
                Apenas planos ativos ficam disponíveis para novos subscriptions
              </p>
            </div>
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
