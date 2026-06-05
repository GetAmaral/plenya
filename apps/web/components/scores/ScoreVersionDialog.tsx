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
import { Switch } from '@/components/ui/switch'
import {
  ScoreVersion,
  CreateScoreVersionDTO,
  UpdateScoreVersionDTO,
  useCreateScoreVersion,
  useUpdateScoreVersion,
} from '@/lib/api/score-version-api'

interface ScoreVersionDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  version?: ScoreVersion
}

export function ScoreVersionDialog({ open, onOpenChange, version }: ScoreVersionDialogProps) {
  const isEditing = !!version

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createVersion = useCreateScoreVersion()
  const updateVersion = useUpdateScoreVersion()

  const {
    register,
    handleSubmit,
    reset,
    watch,
    setValue,
    formState: { errors },
  } = useForm<CreateScoreVersionDTO>({
    defaultValues: {
      name: version?.name || '',
      slug: version?.slug || '',
      description: version?.description || '',
      siteIntro: version?.siteIntro || '',
      order: version?.order || 0,
      active: version?.active ?? true,
    },
  })

  const active = watch('active')

  useEffect(() => {
    reset(
      version
        ? {
            name: version.name,
            slug: version.slug,
            description: version.description || '',
            siteIntro: version.siteIntro || '',
            order: version.order,
            active: version.active,
          }
        : { name: '', slug: '', description: '', siteIntro: '', order: 0, active: true }
    )
  }, [version, reset])

  const onSubmit = async (data: CreateScoreVersionDTO) => {
    // Normaliza '' -> undefined nos opcionais.
    const payload: CreateScoreVersionDTO = {
      name: data.name.trim(),
      slug: data.slug.trim(),
      description: data.description?.trim() || undefined,
      siteIntro: data.siteIntro?.trim() || undefined,
      order: data.order ?? 0,
      active: data.active,
    }
    try {
      if (isEditing) {
        await updateVersion.mutateAsync({ id: version.id, data: payload as UpdateScoreVersionDTO })
        toast.success('Versão atualizada com sucesso')
      } else {
        await createVersion.mutateAsync(payload)
        toast.success('Versão criada com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error: any) {
      toast.error(error?.message || (isEditing ? 'Erro ao atualizar versão' : 'Erro ao criar versão'))
    }
  }

  const isSubmitting = createVersion.isPending || updateVersion.isPending

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>{isEditing ? 'Editar Versão' : 'Nova Versão de Escore'}</DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite os metadados da versão. Os itens são configurados na tela de itens.'
              : 'Crie uma versão (ex: Triagem, Light) que seleciona e ordena um subconjunto de itens para o site.'}
          </DialogDescription>
        </DialogHeader>

        <form ref={formRef} onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Nome *</Label>
            <Input
              id="name"
              placeholder="Ex: Escore Triagem"
              {...register('name', {
                required: 'Nome é obrigatório',
                minLength: { value: 2, message: 'Mínimo 2 caracteres' },
                maxLength: { value: 120, message: 'Máximo 120 caracteres' },
              })}
            />
            {errors.name && <p className="text-sm text-destructive">{errors.name.message}</p>}
          </div>

          <div className="space-y-2">
            <Label htmlFor="slug">Slug *</Label>
            <Input
              id="slug"
              placeholder="ex: triagem"
              {...register('slug', {
                required: 'Slug é obrigatório',
                minLength: { value: 2, message: 'Mínimo 2 caracteres' },
                maxLength: { value: 60, message: 'Máximo 60 caracteres' },
                pattern: {
                  value: /^[a-z0-9-]+$/,
                  message: 'Use apenas letras minúsculas, números e hífen',
                },
              })}
            />
            {errors.slug && <p className="text-sm text-destructive">{errors.slug.message}</p>}
            <p className="text-xs text-muted-foreground">
              Identifica a URL pública do config (<code>/score-light/config/{watch('slug') || 'slug'}</code>) e o
              arquivo estático do site. {isEditing && 'Trocar o slug de uma versão já publicada quebra o site, evite.'}
            </p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descrição</Label>
            <Textarea
              id="description"
              rows={2}
              placeholder="Uso interno: para que serve esta versão"
              {...register('description')}
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="siteIntro">Texto de introdução (site)</Label>
            <Textarea
              id="siteIntro"
              rows={2}
              placeholder="Texto exibido ao público no início do questionário"
              {...register('siteIntro')}
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="order">Ordem</Label>
              <Input
                id="order"
                type="number"
                placeholder="0"
                {...register('order', {
                  valueAsNumber: true,
                  min: { value: 0, message: 'Mínimo 0' },
                  max: { value: 9999, message: 'Máximo 9999' },
                })}
              />
              {errors.order && <p className="text-sm text-destructive">{errors.order.message}</p>}
            </div>

            <div className="space-y-2">
              <Label htmlFor="active">Ativa</Label>
              <div className="flex h-10 items-center gap-2">
                <Switch
                  id="active"
                  checked={active}
                  onCheckedChange={(v) => setValue('active', v, { shouldValidate: true })}
                />
                <span className="text-sm text-muted-foreground">{active ? 'Sim' : 'Não'}</span>
              </div>
            </div>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={isSubmitting}>
              Cancelar
            </Button>
            <Button type="submit" disabled={isSubmitting}>
              {isSubmitting ? (isEditing ? 'Salvando...' : 'Criando...') : isEditing ? 'Salvar' : 'Criar'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
