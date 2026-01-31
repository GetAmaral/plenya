'use client'

import { useEffect, useMemo, useRef, useState } from 'react'
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
import {
  ScoreItem,
  CreateScoreItemDTO,
  UpdateScoreItemDTO,
  useCreateScoreItem,
  useUpdateScoreItem,
  useAllScoreGroupTrees,
  useScoreSubgroup,
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

  // Fetch all groups with subgroups
  const { data: allGroups } = useAllScoreGroupTrees()

  // State for selected subgroup
  const [selectedSubgroupId, setSelectedSubgroupId] = useState(item?.subgroupId || subgroupId)
  const [selectedParentItemId, setSelectedParentItemId] = useState(item?.parentItemId || '')

  // Fetch current subgroup to get group info
  const { data: currentSubgroup } = useScoreSubgroup(selectedSubgroupId)

  // Get current group name
  const currentGroupName = useMemo(() => {
    if (!allGroups || !currentSubgroup) return ''
    const group = allGroups.find(g => g.id === currentSubgroup.groupId)
    return group?.name || ''
  }, [allGroups, currentSubgroup])

  // Check if current group is "Exames"
  const isExamsGroup = currentGroupName === 'Exames'

  // Get all available items for parentItem selection (from the same subgroup, excluding current item)
  const availableParentItems = useMemo(() => {
    if (!allGroups) return []

    // Find the selected subgroup
    for (const group of allGroups) {
      if (!group.subgroups) continue
      for (const subgroup of group.subgroups) {
        if (subgroup.id === selectedSubgroupId) {
          // Return items excluding the current item being edited
          return (subgroup.items || []).filter(i => !item || i.id !== item.id)
        }
      }
    }
    return []
  }, [allGroups, selectedSubgroupId, item])

  const {
    register,
    handleSubmit,
    reset,
    setValue,
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
      subgroupId: selectedSubgroupId,
      parentItemId: item?.parentItemId,
    },
  })

  // Reset form when item changes
  useEffect(() => {
    if (item) {
      setSelectedSubgroupId(item.subgroupId)
      setSelectedParentItemId(item.parentItemId || '')
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
        parentItemId: item.parentItemId,
      })
    } else {
      setSelectedSubgroupId(subgroupId)
      setSelectedParentItemId('')
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
        parentItemId: undefined,
      })
    }
  }, [item, subgroupId, reset])

  // Update form value when selectedSubgroupId changes
  useEffect(() => {
    setValue('subgroupId', selectedSubgroupId)
  }, [selectedSubgroupId, setValue])

  // Update form value when selectedParentItemId changes
  useEffect(() => {
    setValue('parentItemId', selectedParentItemId || undefined)
  }, [selectedParentItemId, setValue])

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
            subgroupId: selectedSubgroupId,
          } as UpdateScoreItemDTO,
        })
        toast.success('Item atualizado com sucesso')
      } else {
        await createItem.mutateAsync({
          ...payload,
          subgroupId: selectedSubgroupId,
        })
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
            <Label htmlFor="subgroup">Subgrupo *</Label>
            <Select
              value={selectedSubgroupId}
              onValueChange={setSelectedSubgroupId}
            >
              <SelectTrigger id="subgroup">
                <SelectValue placeholder="Selecione o subgrupo" />
              </SelectTrigger>
              <SelectContent>
                {allGroups?.map((group) =>
                  group.subgroups?.map((subgroup) => (
                    <SelectItem key={subgroup.id} value={subgroup.id}>
                      {group.name} &gt; {subgroup.name}
                    </SelectItem>
                  ))
                )}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label htmlFor="parentItem">Item Pai (opcional)</Label>
            <Select
              value={selectedParentItemId}
              onValueChange={setSelectedParentItemId}
            >
              <SelectTrigger id="parentItem">
                <SelectValue placeholder="Nenhum (item raiz)" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">Nenhum (item raiz)</SelectItem>
                {availableParentItems.map((parentItem) => (
                  <SelectItem key={parentItem.id} value={parentItem.id}>
                    {parentItem.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <p className="text-xs text-muted-foreground">
              Selecione um item pai para criar uma hierarquia
            </p>
          </div>

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

          {isExamsGroup && (
            <>
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
            </>
          )}

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
