'use client'

import { useRef } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'

import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Checkbox } from '@/components/ui/checkbox'
import { useFormNavigation } from '@/lib/use-form-navigation'
import {
  medicationDefinitionSchema,
  type MedicationDefinitionInput,
} from '@/lib/validations/medication-definition'

interface MedicationDefinitionFormProps {
  defaultValues?: Partial<MedicationDefinitionInput>
  onSubmit: (data: MedicationDefinitionInput) => void
  isSubmitting?: boolean
  submitLabel?: string
}

export function MedicationDefinitionForm({
  defaultValues,
  onSubmit,
  isSubmitting,
  submitLabel = 'Salvar',
}: MedicationDefinitionFormProps) {
  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const form = useForm<MedicationDefinitionInput>({
    resolver: zodResolver(medicationDefinitionSchema),
    defaultValues: {
      commonName: '',
      activeIngredient: '',
      category: 'simple',
      validityDays: 30,
      maxPerPrescription: 10,
      maxTreatmentDays: 60,
      requiresDigitalSignature: false,
      requiresSNCR: false,
      anvisaCode: '',
      ...defaultValues,
    },
  })

  return (
    <form ref={formRef} onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
      <Card>
        <CardHeader>
          <CardTitle>Identificação</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <Label htmlFor="commonName">Nome Comercial *</Label>
            <Input
              id="commonName"
              {...form.register('commonName')}
              placeholder="Ex: Dipirona 500mg"
            />
            {form.formState.errors.commonName && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.commonName.message}
              </p>
            )}
          </div>

          <div>
            <Label htmlFor="activeIngredient">Princípio Ativo (DCB) *</Label>
            <Input
              id="activeIngredient"
              {...form.register('activeIngredient')}
              placeholder="Ex: Dipirona Sódica"
            />
            {form.formState.errors.activeIngredient && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.activeIngredient.message}
              </p>
            )}
          </div>

          <div>
            <Label htmlFor="anvisaCode">Código ANVISA (opcional)</Label>
            <Input
              id="anvisaCode"
              {...form.register('anvisaCode')}
              placeholder="Ex: 1234567890123"
            />
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Categoria e Regulamentação</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <Label htmlFor="category">Categoria *</Label>
            <Select
              value={form.watch('category')}
              onValueChange={(value) =>
                form.setValue('category', value as MedicationDefinitionInput['category'])
              }
            >
              <SelectTrigger id="category">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="simple">Simples (Receita Branca)</SelectItem>
                <SelectItem value="c1">C1 - Controle Especial</SelectItem>
                <SelectItem value="c5">C5 - Psicotrópico</SelectItem>
                <SelectItem value="antibiotic">Antimicrobiano (RDC 471)</SelectItem>
                <SelectItem value="glp1">GLP-1 Agonista</SelectItem>
              </SelectContent>
            </Select>
            {form.formState.errors.category && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.category.message}
              </p>
            )}
          </div>

          <div className="space-y-3">
            <div className="flex items-center space-x-2">
              <Checkbox
                id="requiresSNCR"
                checked={form.watch('requiresSNCR')}
                onCheckedChange={(checked) =>
                  form.setValue('requiresSNCR', checked as boolean)
                }
              />
              <Label htmlFor="requiresSNCR" className="font-normal cursor-pointer">
                Requer registro no SNCR (Sistema Nacional de Controle de Receitas)
              </Label>
            </div>

            <div className="flex items-center space-x-2">
              <Checkbox
                id="requiresDigitalSignature"
                checked={form.watch('requiresDigitalSignature')}
                onCheckedChange={(checked) =>
                  form.setValue('requiresDigitalSignature', checked as boolean)
                }
              />
              <Label htmlFor="requiresDigitalSignature" className="font-normal cursor-pointer">
                Requer assinatura digital ICP-Brasil
              </Label>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Regras de Prescrição</CardTitle>
        </CardHeader>
        <CardContent className="grid grid-cols-3 gap-4">
          <div>
            <Label htmlFor="validityDays">Validade (dias) *</Label>
            <Input
              id="validityDays"
              type="number"
              {...form.register('validityDays')}
              placeholder="30"
            />
            {form.formState.errors.validityDays && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.validityDays.message}
              </p>
            )}
            <p className="text-xs text-muted-foreground mt-1">
              Validade padrão da prescrição
            </p>
          </div>

          <div>
            <Label htmlFor="maxPerPrescription">Máx. por Prescrição *</Label>
            <Input
              id="maxPerPrescription"
              type="number"
              {...form.register('maxPerPrescription')}
              placeholder="10"
            />
            {form.formState.errors.maxPerPrescription && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.maxPerPrescription.message}
              </p>
            )}
            <p className="text-xs text-muted-foreground mt-1">
              Quantidade máxima permitida
            </p>
          </div>

          <div>
            <Label htmlFor="maxTreatmentDays">Máx. Tratamento (dias) *</Label>
            <Input
              id="maxTreatmentDays"
              type="number"
              {...form.register('maxTreatmentDays')}
              placeholder="60"
            />
            {form.formState.errors.maxTreatmentDays && (
              <p className="text-sm text-destructive mt-1">
                {form.formState.errors.maxTreatmentDays.message}
              </p>
            )}
            <p className="text-xs text-muted-foreground mt-1">
              Duração máxima do tratamento
            </p>
          </div>
        </CardContent>
      </Card>

      <div className="flex justify-end gap-3">
        <Button
          type="button"
          variant="outline"
          onClick={() => window.history.back()}
          disabled={isSubmitting}
        >
          Cancelar
        </Button>
        <Button type="submit" disabled={isSubmitting}>
          {isSubmitting ? 'Salvando...' : submitLabel}
        </Button>
      </div>
    </form>
  )
}
