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
  ScoreItem,
  CreateScoreItemDTO,
  UpdateScoreItemDTO,
  useCreateScoreItem,
  useUpdateScoreItem,
} from '@/lib/api/score-api'

interface ScoreItemDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  subgroupId: string
  item?: ScoreItem
}

export function ScoreItemDialog({
  open,
  onOpenChange,
  subgroupId,
  item,
}: ScoreItemDialogProps) {
  const isEditing = !!item

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createItem = useCreateScoreItem()
  const updateItem = useUpdateScoreItem()

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateScoreItemDTO>({
    defaultValues: {
      name: item?.name || '',
      unit: item?.unit || '',
      unitConversion: item?.unitConversion || '',
      clinicalRelevance: item?.clinicalRelevance || '',
      patientExplanation: item?.patientExplanation || '',
      conduct: item?.conduct || '',
      points: item?.points || 0,
      order: item?.order || 0,
      subgroupId: subgroupId,
    },
  })

  // Reset form when item changes
  useEffect(() => {
    if (item) {
      reset({
        name: item.name,
        unit: item.unit || '',
        unitConversion: item.unitConversion || '',
        clinicalRelevance: item.clinicalRelevance || '',
        patientExplanation: item.patientExplanation || '',
        conduct: item.conduct || '',
        points: item.points,
        order: item.order,
        subgroupId: item.subgroupId,
      })
    } else {
      reset({
        name: '',
        unit: '',
        unitConversion: '',
        clinicalRelevance: '',
        patientExplanation: '',
        conduct: '',
        points: 0,
        order: 0,
        subgroupId: subgroupId,
      })
    }
  }, [item, subgroupId, reset])

  const onSubmit = async (data: CreateScoreItemDTO) => {
    try {
      // Convert empty strings to undefined for optional fields
      const payload = {
        ...data,
        unit: data.unit || undefined,
        unitConversion: data.unitConversion || undefined,
        clinicalRelevance: data.clinicalRelevance || undefined,
        patientExplanation: data.patientExplanation || undefined,
        conduct: data.conduct || undefined,
      }

      if (isEditing) {
        await updateItem.mutateAsync({
          id: item.id,
          data: {
            name: payload.name,
            unit: payload.unit,
            unitConversion: payload.unitConversion,
            clinicalRelevance: payload.clinicalRelevance,
            patientExplanation: payload.patientExplanation,
            conduct: payload.conduct,
            points: payload.points,
            order: payload.order,
          } as UpdateScoreItemDTO,
        })
        toast.success('Item atualizado com sucesso')
      } else {
        await createItem.mutateAsync(payload)
        toast.success('Item criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error) {
      toast.error(
        isEditing ? 'Erro ao atualizar item' : 'Erro ao criar item'
      )
    }
  }

  const isSubmitting = createItem.isPending || updateItem.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Item' : 'Novo Item de Escore'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do item de escore'
              : 'Crie um novo item de escore (parâmetro clínico)'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Nome do Parâmetro *</Label>
            <Input
              id="name"
              placeholder="Ex: Hemoglobina - Homens"
              {...register('name', {
                required: 'Nome é obrigatório',
                minLength: {
                  value: 2,
                  message: 'Nome deve ter no mínimo 2 caracteres',
                },
                maxLength: {
                  value: 300,
                  message: 'Nome deve ter no máximo 300 caracteres',
                },
              })}
            />
            {errors.name && (
              <p className="text-sm text-destructive">{errors.name.message}</p>
            )}
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="unit">Unidade de Medida</Label>
              <Input
                id="unit"
                placeholder="Ex: g/dL"
                {...register('unit', {
                  maxLength: {
                    value: 50,
                    message: 'Unidade deve ter no máximo 50 caracteres',
                  },
                })}
              />
              {errors.unit && (
                <p className="text-sm text-destructive">{errors.unit.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="points">Pontos Máximos *</Label>
              <Input
                id="points"
                type="number"
                step="0.1"
                placeholder="0"
                {...register('points', {
                  valueAsNumber: true,
                  required: 'Pontos é obrigatório',
                  min: {
                    value: 0,
                    message: 'Pontos deve ser maior ou igual a 0',
                  },
                  max: {
                    value: 100,
                    message: 'Pontos deve ser menor ou igual a 100',
                  },
                })}
              />
              {errors.points && (
                <p className="text-sm text-destructive">{errors.points.message}</p>
              )}
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="unitConversion">Conversão de Unidade</Label>
            <Textarea
              id="unitConversion"
              placeholder="Ex: 1 g/dL = 10 g/L"
              rows={2}
              {...register('unitConversion')}
            />
            <p className="text-xs text-muted-foreground">
              Informação opcional sobre conversão entre unidades
            </p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="order">Ordem de Exibição</Label>
            <Input
              id="order"
              type="number"
              placeholder="0"
              {...register('order', {
                valueAsNumber: true,
                min: {
                  value: 0,
                  message: 'Ordem deve ser maior ou igual a 0',
                },
                max: {
                  value: 9999,
                  message: 'Ordem deve ser menor que 10000',
                },
              })}
            />
            {errors.order && (
              <p className="text-sm text-destructive">{errors.order.message}</p>
            )}
            <p className="text-xs text-muted-foreground">
              Deixe 0 para ordenação automática
            </p>
          </div>

          <div className="border-t pt-4 space-y-4">
            <h3 className="text-sm font-semibold text-muted-foreground">
              Informações Clínicas (Opcional)
            </h3>

            <div className="space-y-2">
              <Label htmlFor="clinicalRelevance">Relevância Clínica</Label>
              <Textarea
                id="clinicalRelevance"
                placeholder="Explicação técnica para profissionais de saúde..."
                rows={3}
                {...register('clinicalRelevance')}
              />
              <p className="text-xs text-muted-foreground">
                Explicação técnica da relevância clínica deste parâmetro
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="patientExplanation">Explicação para o Paciente</Label>
              <Textarea
                id="patientExplanation"
                placeholder="Explicação em linguagem simples e acessível..."
                rows={3}
                {...register('patientExplanation')}
              />
              <p className="text-xs text-muted-foreground">
                Explicação em linguagem simples para o paciente entender
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="conduct">Conduta Clínica</Label>
              <Textarea
                id="conduct"
                placeholder="Orientações de conduta clínica recomendada..."
                rows={3}
                {...register('conduct')}
              />
              <p className="text-xs text-muted-foreground">
                Orientações de conduta clínica baseadas em evidências
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
