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
import type { MethodLetter } from '@plenya/types'
import {
  CreateMethodLetterDTO,
  UpdateMethodLetterDTO,
  useCreateMethodLetter,
  useUpdateMethodLetter,
} from '@/lib/api/method-api'

interface MethodLetterDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  methodId: string
  letter?: MethodLetter
}

export function MethodLetterDialog({
  open,
  onOpenChange,
  methodId,
  letter,
}: MethodLetterDialogProps) {
  const isEditing = !!letter

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createLetter = useCreateMethodLetter()
  const updateLetter = useUpdateMethodLetter()

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateMethodLetterDTO>({
    defaultValues: {
      code: letter?.code || '',
      name: letter?.name || '',
      description: letter?.description || '',
      clinicalRelevance: letter?.clinicalRelevance || '',
      patientExplanation: letter?.patientExplanation || '',
      conduct: letter?.conduct || '',
      color: letter?.color || '#10B981',
      icon: letter?.icon || '',
      order: letter?.order || 0,
    },
  })

  // Reset form when letter changes
  useEffect(() => {
    if (letter) {
      reset({
        code: letter.code,
        name: letter.name,
        description: letter.description || '',
        clinicalRelevance: letter.clinicalRelevance || '',
        patientExplanation: letter.patientExplanation || '',
        conduct: letter.conduct || '',
        color: letter.color || '#10B981',
        icon: letter.icon || '',
        order: letter.order,
      })
    } else {
      reset({
        code: '',
        name: '',
        description: '',
        clinicalRelevance: '',
        patientExplanation: '',
        conduct: '',
        color: '#10B981',
        icon: '',
        order: 0,
      })
    }
  }, [letter, reset])

  const onSubmit = async (data: CreateMethodLetterDTO) => {
    try {
      // Convert empty strings to undefined for optional fields
      const payload = {
        ...data,
        description: data.description || undefined,
        clinicalRelevance: data.clinicalRelevance || undefined,
        patientExplanation: data.patientExplanation || undefined,
        conduct: data.conduct || undefined,
        color: data.color || undefined,
        icon: data.icon || undefined,
      }

      if (isEditing) {
        await updateLetter.mutateAsync({
          id: letter.id,
          data: payload as UpdateMethodLetterDTO,
        })
        toast.success('Letra atualizada com sucesso')
      } else {
        await createLetter.mutateAsync({
          methodId,
          data: payload,
        })
        toast.success('Letra criada com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      const errorMessage =
        error?.message ||
        (isEditing ? 'Erro ao atualizar letra' : 'Erro ao criar letra')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createLetter.isPending || updateLetter.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Letra' : 'Nova Letra da Metodologia'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações da letra'
              : 'Crie uma nova letra para organizar pilares'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          {/* Basic Info */}
          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <Label htmlFor="code">Código *</Label>
              <Input
                id="code"
                placeholder="Ex: A"
                {...register('code', {
                  required: 'Código é obrigatório',
                  minLength: {
                    value: 1,
                    message: 'Código deve ter no mínimo 1 caractere',
                  },
                  maxLength: {
                    value: 10,
                    message: 'Código deve ter no máximo 10 caracteres',
                  },
                })}
              />
              {errors.code && (
                <p className="text-sm text-destructive">{errors.code.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="icon">Ícone/Emoji</Label>
              <Input
                id="icon"
                placeholder="🥗"
                {...register('icon', {
                  maxLength: {
                    value: 50,
                    message: 'Ícone deve ter no máximo 50 caracteres',
                  },
                })}
              />
              {errors.icon && (
                <p className="text-sm text-destructive">{errors.icon.message}</p>
              )}
              <p className="text-xs text-muted-foreground">
                Emoji ou ícone Unicode
              </p>
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
            <Label htmlFor="name">Nome Completo *</Label>
            <Input
              id="name"
              placeholder="Ex: Alimentação e Atividade Física"
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
            <Label htmlFor="color">Cor do Tema</Label>
            <div className="flex gap-2">
              <Input
                id="color"
                type="color"
                className="w-16 h-10 p-1 cursor-pointer"
                {...register('color', {
                  pattern: {
                    value: /^#[0-9A-Fa-f]{6}$/,
                    message: 'Cor inválida (use formato #RRGGBB)',
                  },
                })}
              />
              <Input
                type="text"
                placeholder="#10B981"
                className="flex-1"
                {...register('color')}
              />
            </div>
            {errors.color && (
              <p className="text-sm text-destructive">{errors.color.message}</p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descrição</Label>
            <Textarea
              id="description"
              placeholder="Descreva a letra..."
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
                rows={2}
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
                rows={2}
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
                rows={2}
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
