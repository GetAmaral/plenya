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
import { Checkbox } from '@/components/ui/checkbox'
import type { Method } from '@plenya/types'
import {
  CreateMethodDTO,
  UpdateMethodDTO,
  useCreateMethod,
  useUpdateMethod,
} from '@/lib/api/method-api'

interface MethodDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  method?: Method
}

export function MethodDialog({
  open,
  onOpenChange,
  method,
}: MethodDialogProps) {
  const isEditing = !!method

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createMethod = useCreateMethod()
  const updateMethod = useUpdateMethod()

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<CreateMethodDTO>({
    defaultValues: {
      name: method?.name || '',
      shortName: method?.shortName || '',
      description: method?.description || '',
      version: method?.version || '1.0',
      color: method?.color || '#6366F1',
      order: method?.order || 0,
      isDefault: method?.isDefault || false,
    },
  })

  // Reset form when method changes
  useEffect(() => {
    if (method) {
      reset({
        name: method.name,
        shortName: method.shortName,
        description: method.description || '',
        version: method.version || '1.0',
        color: method.color || '#6366F1',
        order: method.order,
        isDefault: method.isDefault || false,
      })
    } else {
      reset({
        name: '',
        shortName: '',
        description: '',
        version: '1.0',
        color: '#6366F1',
        order: 0,
        isDefault: false,
      })
    }
  }, [method, reset])

  const onSubmit = async (data: CreateMethodDTO) => {
    try {
      // Convert empty strings to undefined for optional fields
      const payload = {
        ...data,
        description: data.description || undefined,
        version: data.version || undefined,
        color: data.color || undefined,
      }

      if (isEditing) {
        await updateMethod.mutateAsync({
          id: method.id,
          data: payload as UpdateMethodDTO,
        })
        toast.success('Metodologia atualizada com sucesso')
      } else {
        await createMethod.mutateAsync(payload)
        toast.success('Metodologia criada com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      const errorMessage =
        error?.message ||
        (isEditing ? 'Erro ao atualizar metodologia' : 'Erro ao criar metodologia')
      toast.error(errorMessage)
    }
  }

  const isSubmitting = createMethod.isPending || updateMethod.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Metodologia' : 'Nova Metodologia Clínica'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações da metodologia'
              : 'Crie uma nova metodologia para organizar itens de escore'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="name">Nome Completo *</Label>
              <Input
                id="name"
                placeholder="Ex: AGIR - Protocolo de Saúde Integrativa"
                {...register('name', {
                  required: 'Nome é obrigatório',
                  minLength: {
                    value: 3,
                    message: 'Nome deve ter no mínimo 3 caracteres',
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
              <Label htmlFor="shortName">Sigla *</Label>
              <Input
                id="shortName"
                placeholder="Ex: AGIR"
                {...register('shortName', {
                  required: 'Sigla é obrigatória',
                  minLength: {
                    value: 2,
                    message: 'Sigla deve ter no mínimo 2 caracteres',
                  },
                  maxLength: {
                    value: 20,
                    message: 'Sigla deve ter no máximo 20 caracteres',
                  },
                })}
              />
              {errors.shortName && (
                <p className="text-sm text-destructive">{errors.shortName.message}</p>
              )}
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descrição</Label>
            <Textarea
              id="description"
              placeholder="Descreva a metodologia..."
              rows={3}
              {...register('description', {
                maxLength: {
                  value: 1000,
                  message: 'Descrição deve ter no máximo 1000 caracteres',
                },
              })}
            />
            {errors.description && (
              <p className="text-sm text-destructive">{errors.description.message}</p>
            )}
          </div>

          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <Label htmlFor="version">Versão</Label>
              <Input
                id="version"
                placeholder="1.0"
                {...register('version', {
                  maxLength: {
                    value: 20,
                    message: 'Versão deve ter no máximo 20 caracteres',
                  },
                })}
              />
              {errors.version && (
                <p className="text-sm text-destructive">{errors.version.message}</p>
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
                  placeholder="#6366F1"
                  className="flex-1"
                  {...register('color')}
                />
              </div>
              {errors.color && (
                <p className="text-sm text-destructive">{errors.color.message}</p>
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

          <div className="flex items-center space-x-2 p-4 bg-muted/50 rounded-lg">
            <Checkbox
              id="isDefault"
              checked={watch('isDefault')}
              onCheckedChange={(checked) => setValue('isDefault', checked as boolean)}
            />
            <div className="grid gap-1.5 leading-none">
              <Label
                htmlFor="isDefault"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 cursor-pointer"
              >
                Metodologia Padrão
              </Label>
              <p className="text-sm text-muted-foreground">
                Ao marcar como padrão, todas as outras metodologias serão desmarcadas automaticamente
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
