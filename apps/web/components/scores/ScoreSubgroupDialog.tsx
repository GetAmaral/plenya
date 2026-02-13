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
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  ScoreSubgroup,
  CreateScoreSubgroupDTO,
  UpdateScoreSubgroupDTO,
  useCreateScoreSubgroup,
  useUpdateScoreSubgroup,
} from '@/lib/api/score-api'

interface ScoreSubgroupDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  groupId: string
  subgroup?: ScoreSubgroup
}

export function ScoreSubgroupDialog({
  open,
  onOpenChange,
  groupId,
  subgroup,
}: ScoreSubgroupDialogProps) {
  const isEditing = !!subgroup

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createSubgroup = useCreateScoreSubgroup()
  const updateSubgroup = useUpdateScoreSubgroup()

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<CreateScoreSubgroupDTO>({
    defaultValues: {
      name: subgroup?.name || '',
      order: subgroup?.order || 0,
      maxSelect: subgroup?.maxSelect || 0,
      groupId: groupId,
    },
  })

  // Reset form when subgroup changes
  useEffect(() => {
    if (subgroup) {
      reset({
        name: subgroup.name,
        order: subgroup.order,
        maxSelect: subgroup.maxSelect || 0,
        groupId: subgroup.groupId,
      })
    } else {
      reset({
        name: '',
        order: 0,
        maxSelect: 0,
        groupId: groupId,
      })
    }
  }, [subgroup, groupId, reset])

  const onSubmit = async (data: CreateScoreSubgroupDTO) => {
    try {
      if (isEditing) {
        await updateSubgroup.mutateAsync({
          id: subgroup.id,
          data: {
            name: data.name,
            order: data.order,
            maxSelect: data.maxSelect,
          } as UpdateScoreSubgroupDTO,
        })
        toast.success('Subgrupo atualizado com sucesso')
      } else {
        await createSubgroup.mutateAsync(data)
        toast.success('Subgrupo criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      // Extract error message from API error
      const errorMessage = error?.message || (isEditing ? 'Erro ao atualizar subgrupo' : 'Erro ao criar subgrupo')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createSubgroup.isPending || updateSubgroup.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Subgrupo' : 'Novo Subgrupo'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do subgrupo'
              : 'Crie um novo subgrupo para organizar os itens de escore'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Nome do Subgrupo *</Label>
            <Input
              id="name"
              placeholder="Ex: Série Vermelha"
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

          <div className="space-y-2">
            <Label htmlFor="maxSelect">Máximo de Itens Selecionáveis</Label>
            <Select
              value={String(watch('maxSelect') || 0)}
              onValueChange={(value) =>
                setValue('maxSelect', Number(value), { shouldValidate: true })
              }
            >
              <SelectTrigger>
                <SelectValue placeholder="Selecione..." />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="0">Não se aplica</SelectItem>
                <SelectItem value="1">1</SelectItem>
                <SelectItem value="2">2</SelectItem>
                <SelectItem value="3">3</SelectItem>
                <SelectItem value="4">4</SelectItem>
                <SelectItem value="5">5</SelectItem>
                <SelectItem value="6">6</SelectItem>
              </SelectContent>
            </Select>
            {errors.maxSelect && (
              <p className="text-sm text-destructive">
                {errors.maxSelect.message}
              </p>
            )}
            <p className="text-xs text-muted-foreground">
              0 = multi-select ilimitado. 1+ = limite de seleções permitidas
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
