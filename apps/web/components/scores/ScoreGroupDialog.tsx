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
import {
  ScoreGroup,
  CreateScoreGroupDTO,
  UpdateScoreGroupDTO,
  useCreateScoreGroup,
  useUpdateScoreGroup,
} from '@/lib/api/score-api'

interface ScoreGroupDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  group?: ScoreGroup
}

export function ScoreGroupDialog({
  open,
  onOpenChange,
  group,
}: ScoreGroupDialogProps) {
  const isEditing = !!group

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createGroup = useCreateScoreGroup()
  const updateGroup = useUpdateScoreGroup()

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateScoreGroupDTO>({
    defaultValues: {
      name: group?.name || '',
      order: group?.order || 0,
    },
  })

  // Reset form when group changes
  useEffect(() => {
    if (group) {
      reset({
        name: group.name,
        order: group.order,
      })
    } else {
      reset({
        name: '',
        order: 0,
      })
    }
  }, [group, reset])

  const onSubmit = async (data: CreateScoreGroupDTO) => {
    try {
      if (isEditing) {
        await updateGroup.mutateAsync({
          id: group.id,
          data: data as UpdateScoreGroupDTO,
        })
        toast.success('Grupo atualizado com sucesso')
      } else {
        await createGroup.mutateAsync(data)
        toast.success('Grupo criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      // Extract error message from API error
      const errorMessage = error?.message || (isEditing ? 'Erro ao atualizar grupo' : 'Erro ao criar grupo')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createGroup.isPending || updateGroup.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Grupo' : 'Novo Grupo de Escores'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do grupo de escores'
              : 'Crie um novo grupo para organizar os escores clínicos'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Nome do Grupo *</Label>
            <Input
              id="name"
              placeholder="Ex: Hemograma Completo"
              {...register('name', {
                required: 'Nome é obrigatório',
                minLength: {
                  value: 2,
                  message: 'Nome deve ter no mínimo 2 caracteres',
                },
                maxLength: {
                  value: 200,
                  message: 'Nome deve ter no máximo 200 caracteres',
                },
              })}
            />
            {errors.name && (
              <p className="text-sm text-destructive">{errors.name.message}</p>
            )}
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
