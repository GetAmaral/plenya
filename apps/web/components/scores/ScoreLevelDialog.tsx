'use client'

import { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
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
  ScoreLevel,
  CreateScoreLevelDTO,
  UpdateScoreLevelDTO,
  useCreateScoreLevel,
  useUpdateScoreLevel,
} from '@/lib/api/score-api'

interface ScoreLevelDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  itemId: string
  level?: ScoreLevel
}

const OPERATORS = [
  { value: '=', label: 'Igual a (=)' },
  { value: '>', label: 'Maior que (>)' },
  { value: '>=', label: 'Maior ou igual (≥)' },
  { value: '<', label: 'Menor que (<)' },
  { value: '<=', label: 'Menor ou igual (≤)' },
  { value: 'between', label: 'Entre (intervalo)' },
] as const

const LEVEL_OPTIONS = [
  { value: 0, label: 'Nível 0 - Crítico (pior)', color: 'text-red-600' },
  { value: 1, label: 'Nível 1 - Muito Baixo/Alto', color: 'text-orange-600' },
  { value: 2, label: 'Nível 2 - Subótimo', color: 'text-yellow-600' },
  { value: 3, label: 'Nível 3 - Limítrofe', color: 'text-blue-600' },
  { value: 4, label: 'Nível 4 - Bom', color: 'text-green-600' },
  { value: 5, label: 'Nível 5 - Ótimo (melhor)', color: 'text-emerald-600' },
  { value: 6, label: 'Nível 6 - Reservado', color: 'text-gray-600' },
]

export function ScoreLevelDialog({
  open,
  onOpenChange,
  itemId,
  level,
}: ScoreLevelDialogProps) {
  const isEditing = !!level

  const createLevel = useCreateScoreLevel()
  const updateLevel = useUpdateScoreLevel()

  const [selectedOperator, setSelectedOperator] = useState<string>(
    level?.operator || 'between'
  )

  const {
    register,
    handleSubmit,
    reset,
    setValue,
    watch,
    formState: { errors },
  } = useForm<CreateScoreLevelDTO>({
    defaultValues: {
      level: level?.level ?? 5,
      name: level?.name || '',
      lowerLimit: level?.lowerLimit || '',
      upperLimit: level?.upperLimit || '',
      operator: level?.operator || 'between',
      clinicalRelevance: level?.clinicalRelevance || '',
      patientExplanation: level?.patientExplanation || '',
      conduct: level?.conduct || '',
      itemId: itemId,
    },
  })

  const watchOperator = watch('operator')

  // Reset form when level changes
  useEffect(() => {
    if (level) {
      reset({
        level: level.level,
        name: level.name,
        lowerLimit: level.lowerLimit || '',
        upperLimit: level.upperLimit || '',
        operator: level.operator,
        clinicalRelevance: level.clinicalRelevance || '',
        patientExplanation: level.patientExplanation || '',
        conduct: level.conduct || '',
        itemId: level.itemId,
      })
      setSelectedOperator(level.operator)
    } else {
      reset({
        level: 5,
        name: '',
        lowerLimit: '',
        upperLimit: '',
        operator: 'between',
        clinicalRelevance: '',
        patientExplanation: '',
        conduct: '',
        itemId: itemId,
      })
      setSelectedOperator('between')
    }
  }, [level, itemId, reset])

  const onSubmit = async (data: CreateScoreLevelDTO) => {
    try {
      // Convert empty strings to undefined for optional fields
      const payload = {
        ...data,
        lowerLimit: data.lowerLimit || undefined,
        upperLimit: data.upperLimit || undefined,
        clinicalRelevance: data.clinicalRelevance || undefined,
        patientExplanation: data.patientExplanation || undefined,
        conduct: data.conduct || undefined,
      }

      if (isEditing) {
        await updateLevel.mutateAsync({
          id: level.id,
          data: {
            level: payload.level,
            name: payload.name,
            lowerLimit: payload.lowerLimit,
            upperLimit: payload.upperLimit,
            operator: payload.operator,
            clinicalRelevance: payload.clinicalRelevance,
            patientExplanation: payload.patientExplanation,
            conduct: payload.conduct,
          } as UpdateScoreLevelDTO,
        })
        toast.success('Nível atualizado com sucesso')
      } else {
        await createLevel.mutateAsync(payload)
        toast.success('Nível criado com sucesso')
      }
      onOpenChange(false)
      reset()
    } catch (error) {
      toast.error(
        isEditing ? 'Erro ao atualizar nível' : 'Erro ao criar nível'
      )
    }
  }

  const isSubmitting = createLevel.isPending || updateLevel.isPending

  // Show both limit fields only for 'between' operator
  const showBothLimits = watchOperator === 'between'
  const showSingleLimit = !showBothLimits

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>
            {isEditing ? 'Editar Nível' : 'Novo Nível de Estratificação'}
          </DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Edite as informações do nível de risco'
              : 'Defina um novo nível de estratificação de risco (0=pior, 5=melhor)'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="level">Nível *</Label>
              <Select
                value={watch('level')?.toString()}
                onValueChange={(value) => setValue('level', parseInt(value))}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Selecione o nível" />
                </SelectTrigger>
                <SelectContent>
                  {LEVEL_OPTIONS.map((option) => (
                    <SelectItem key={option.value} value={option.value.toString()}>
                      <span className={option.color}>{option.label}</span>
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {errors.level && (
                <p className="text-sm text-destructive">{errors.level.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="operator">Operador *</Label>
              <Select
                value={watchOperator}
                onValueChange={(value) => {
                  setValue('operator', value as any)
                  setSelectedOperator(value)
                }}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Selecione o operador" />
                </SelectTrigger>
                <SelectContent>
                  {OPERATORS.map((op) => (
                    <SelectItem key={op.value} value={op.value}>
                      {op.label}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {errors.operator && (
                <p className="text-sm text-destructive">{errors.operator.message}</p>
              )}
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="name">Descrição do Nível *</Label>
            <Input
              id="name"
              placeholder="Ex: 55 a 70 (Ótimo)"
              {...register('name', {
                required: 'Descrição é obrigatória',
                minLength: {
                  value: 1,
                  message: 'Descrição deve ter no mínimo 1 caractere',
                },
                maxLength: {
                  value: 500,
                  message: 'Descrição deve ter no máximo 500 caracteres',
                },
              })}
            />
            {errors.name && (
              <p className="text-sm text-destructive">{errors.name.message}</p>
            )}
          </div>

          {showBothLimits && (
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="lowerLimit">Limite Inferior</Label>
                <Input
                  id="lowerLimit"
                  placeholder="Ex: 55"
                  {...register('lowerLimit')}
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="upperLimit">Limite Superior</Label>
                <Input
                  id="upperLimit"
                  placeholder="Ex: 70"
                  {...register('upperLimit')}
                />
              </div>
            </div>
          )}

          {showSingleLimit && (
            <div className="space-y-2">
              <Label htmlFor="lowerLimit">Valor de Referência</Label>
              <Input
                id="lowerLimit"
                placeholder="Ex: 40"
                {...register('lowerLimit')}
              />
              <p className="text-xs text-muted-foreground">
                Digite o valor usado na comparação com o operador selecionado
              </p>
            </div>
          )}

          <div className="border-t pt-4 space-y-4">
            <h3 className="text-sm font-semibold text-muted-foreground">
              Informações Clínicas (Opcional)
            </h3>

            <div className="space-y-2">
              <Label htmlFor="clinicalRelevance">Relevância Clínica</Label>
              <Textarea
                id="clinicalRelevance"
                placeholder="Explicação técnica para profissionais de saúde sobre este nível..."
                rows={3}
                {...register('clinicalRelevance')}
              />
              <p className="text-xs text-muted-foreground">
                Explicação técnica da relevância clínica deste nível de risco
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="patientExplanation">Explicação para o Paciente</Label>
              <Textarea
                id="patientExplanation"
                placeholder="Explicação em linguagem simples sobre o que este resultado significa..."
                rows={3}
                {...register('patientExplanation')}
              />
              <p className="text-xs text-muted-foreground">
                Explicação em linguagem simples para o paciente entender este resultado
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="conduct">Conduta Clínica</Label>
              <Textarea
                id="conduct"
                placeholder="Orientações de conduta para este nível de resultado..."
                rows={3}
                {...register('conduct')}
              />
              <p className="text-xs text-muted-foreground">
                Orientações de conduta clínica recomendada para este nível
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
