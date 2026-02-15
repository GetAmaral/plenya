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
import type { MethodPillar } from '@plenya/types'
import {
  CreateMethodPillarDTO,
  UpdateMethodPillarDTO,
  useCreateMethodPillar,
  useUpdateMethodPillar,
} from '@/lib/api/method-api'

interface MethodPillarDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  letterId: string
  pillar?: MethodPillar
}

export function MethodPillarDialog({
  open,
  onOpenChange,
  letterId,
  pillar,
}: MethodPillarDialogProps) {
  const isEditing = !!pillar

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createPillar = useCreateMethodPillar()
  const updatePillar = useUpdateMethodPillar()

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateMethodPillarDTO>({
    defaultValues: {
      name: pillar?.name || '',
      description: pillar?.description || '',
      clinicalRelevance: pillar?.clinicalRelevance || '',
      patientExplanation: pillar?.patientExplanation || '',
      conduct: pillar?.conduct || '',
      order: pillar?.order || 0,
    },
  })

  // Reset form when pillar changes
  useEffect(() => {
    if (pillar) {
      reset({
        name: pillar.name,
        description: pillar.description || '',
        clinicalRelevance: pillar.clinicalRelevance || '',
        patientExplanation: pillar.patientExplanation || '',
        conduct: pillar.conduct || '',
        order: pillar.order,
      })
    } else {
      reset({
        name: '',
        description: '',
        clinicalRelevance: '',
        patientExplanation: '',
        conduct: '',
        order: 0,
      })
    }
  }, [pillar, reset])

  const onSubmit = async (data: CreateMethodPillarDTO) => {
    try {
      // Convert empty strings to undefined for optional fields
      const payload = {
        ...data,
        description: data.description || undefined,
        clinicalRelevance: data.clinicalRelevance || undefined,
        patientExplanation: data.patientExplanation || undefined,
        conduct: data.conduct || undefined,
      }

      if (isEditing) {
        await updatePillar.mutateAsync({
          id: pillar.id,
          data: payload as UpdateMethodPillarDTO,
        })
        toast.success('Pilar atualizado com sucesso')
      } else {
        await createPillar.mutateAsync({
          letterId,
          data: payload,
        })
        toast.success('Pilar criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      const errorMessage =
        error?.message ||
        (isEditing ? 'Erro ao atualizar pilar' : 'Erro ao criar pilar')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createPillar.isPending || updatePillar.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Pilar' : 'Novo Pilar'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do pilar'
              : 'Crie um novo pilar para organizar itens de escore'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Basic Info */}
          <div className="grid grid-cols-3 gap-4">
            <div className="col-span-2 space-y-2">
              <Label htmlFor="name">Nome do Pilar *</Label>
              <Input
                id="name"
                placeholder="Ex: Avaliação Nutricional"
                {...register('name', {
                  required: 'Nome é obrigatório',
                  minLength: {
                    value: 3,
                    message: 'Nome deve ter no mínimo 3 caracteres',
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

            <div className="space-y-2">
              <Label htmlFor="order">Ordem</Label>
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
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descrição</Label>
            <Textarea
              id="description"
              placeholder="Descreva o pilar..."
              rows={2}
              {...register('description')}
            />
          </div>

          {/* Clinical Enrichment Fields */}
          <div className="space-y-4 pt-4 border-t">
            <h3 className="font-medium text-sm">Enriquecimento Clínico (Opcional)</h3>

            <div className="space-y-2">
              <Label htmlFor="clinicalRelevance">Relevância Clínica</Label>
              <Textarea
                id="clinicalRelevance"
                placeholder="Explicação técnica para profissionais..."
                rows={3}
                {...register('clinicalRelevance')}
              />
              <p className="text-xs text-muted-foreground">
                Explicação técnica para profissionais de saúde
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="patientExplanation">Explicação para Paciente</Label>
              <Textarea
                id="patientExplanation"
                placeholder="Explicação em linguagem simples..."
                rows={3}
                {...register('patientExplanation')}
              />
              <p className="text-xs text-muted-foreground">
                Linguagem acessível para pacientes
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="conduct">Conduta Clínica</Label>
              <Textarea
                id="conduct"
                placeholder="Recomendações de conduta..."
                rows={3}
                {...register('conduct')}
              />
              <p className="text-xs text-muted-foreground">
                Recomendações de ação clínica
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
