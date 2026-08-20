'use client'

import { useFieldArray, type Control, type UseFormReturn } from 'react-hook-form'
import { AlertCircle, BookMarked, Plus, Trash2 } from 'lucide-react'

import { Badge } from '@/components/ui/badge'
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
import { Textarea } from '@/components/ui/textarea'
import { FormulaCheckPanel } from '@/components/prescriptions/formula-check-panel'
import { MagistralComponentInput } from '@/components/prescriptions/magistral-component-input'
import { CATEGORY_OPTIONS, CATEGORY_OPTIONS_SHORT } from '@/lib/validations/prescription'
import {
  COMPONENT_UNITS,
  emptyFormulaComponent,
  formulaHasControlled,
  highestCategory,
  PHARMACEUTICAL_FORMS,
  USAGE_OPTIONS,
  type CompoundedPrescriptionFormData,
} from '@/lib/validations/compounded-prescription'

interface FormulaCardProps {
  /** Guarda a fórmula montada como base reutilizável. */
  onSaveAsTemplate?: () => void
  savingTemplate?: boolean
  form: UseFormReturn<CompoundedPrescriptionFormData>
  /** Passado separado do form para o useFieldArray aninhado não re-renderizar tudo a cada tecla. */
  control: Control<CompoundedPrescriptionFormData>
  index: number
  canRemove: boolean
  onRemove: () => void
}

/**
 * Uma fórmula magistral: cabeçalho (forma, uso, aviamento) + lista de componentes.
 *
 * Escolher a forma farmacêutica preenche num clique unidade de aviamento, uso, via e veículo —
 * é o que separa "receita magistral em 30 segundos" de "preencher 8 campos por fórmula".
 */
export function FormulaCard({
  form,
  control,
  index,
  canRemove,
  onRemove,
  onSaveAsTemplate,
  savingTemplate,
}: FormulaCardProps) {
  const { fields, append, remove } = useFieldArray({
    control,
    name: `formulas.${index}.components`,
  })

  const formula = form.watch(`formulas.${index}`)
  const controlled = formula ? formulaHasControlled(formula) : false
  const category = formula ? highestCategory(formula) : 'simple'

  const applyPharmaceuticalForm = (value: string) => {
    const previous = PHARMACEUTICAL_FORMS.find(
      (f) => f.value === form.getValues(`formulas.${index}.pharmaceuticalForm`)
    )
    const preset = PHARMACEUTICAL_FORMS.find((f) => f.value === value)
    form.setValue(`formulas.${index}.pharmaceuticalForm`, value)
    if (!preset) return

    // Componente de creme se escreve em %, o de cápsula em mg. Só troca a unidade de quem ainda
    // está na unidade padrão da forma anterior — escolha manual permanece.
    if (previous && previous.componentUnit !== preset.componentUnit) {
      const components = form.getValues(`formulas.${index}.components`) ?? []
      components.forEach((c, ci) => {
        if (c.unit === previous.componentUnit) {
          form.setValue(`formulas.${index}.components.${ci}.unit`, preset.componentUnit)
        }
      })
    }
    form.setValue(`formulas.${index}.quantityUnit`, preset.quantityUnit)
    form.setValue(`formulas.${index}.usageType`, preset.usageType)
    form.setValue(`formulas.${index}.route`, preset.route)
    // Veículo só é sobrescrito enquanto ainda for um dos padrões. `isDirty` não serve aqui:
    // acrescentar uma fórmula ao field array já marca os campos como sujos, e o veículo da
    // fórmula nova ficaria travado no padrão da forma anterior.
    const current = (form.getValues(`formulas.${index}.vehicle`) || '').trim()
    const isPreset = current === '' || PHARMACEUTICAL_FORMS.some((f) => f.vehicle === current)
    if (isPreset) form.setValue(`formulas.${index}.vehicle`, preset.vehicle)
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex flex-wrap items-center justify-between gap-2">
          <CardTitle className="flex flex-wrap items-center gap-2">
            Fórmula {index + 1}
            {controlled && (
              <Badge variant="outline" className="bg-amber-50 text-xs text-amber-700">
                {CATEGORY_OPTIONS.find((c) => c.value === category)?.label ?? category}
              </Badge>
            )}
          </CardTitle>
          <div className="flex items-center gap-1">
            {onSaveAsTemplate && (
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={onSaveAsTemplate}
                disabled={savingTemplate}
                title="Guardar esta fórmula como base reutilizável"
              >
                <BookMarked className="mr-2 h-4 w-4" />
                Salvar como base
              </Button>
            )}
            {canRemove && (
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={onRemove}
                className="text-destructive hover:text-destructive"
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Remover
              </Button>
            )}
          </div>
        </div>
      </CardHeader>

      <CardContent className="space-y-4">
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <div className="space-y-2">
            <Label>Nome da fórmula (opcional)</Label>
            <Input placeholder="Ex: Fórmula do sono" {...form.register(`formulas.${index}.name`)} />
          </div>

          <div className="space-y-2">
            <Label>Tipo de receita</Label>
            {/* Um seletor por FÓRMULA, não por componente: a receita é do conjunto, e repetir a
                pergunta em cada substância só dava chance de divergirem entre si. */}
            <Select
              value={category}
              onValueChange={(v) => {
                const comps = form.getValues(`formulas.${index}.components`) ?? []
                comps.forEach((_, ci) =>
                  form.setValue(`formulas.${index}.components.${ci}.category`, v as never)
                )
              }}
            >
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {CATEGORY_OPTIONS_SHORT.map((c) => (
                  <SelectItem key={c.value} value={c.value}>
                    {c.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label>Forma farmacêutica *</Label>
            <Select
              value={form.watch(`formulas.${index}.pharmaceuticalForm`)}
              onValueChange={applyPharmaceuticalForm}
            >
              <SelectTrigger>
                <SelectValue placeholder="Escolha" />
              </SelectTrigger>
              <SelectContent>
                {PHARMACEUTICAL_FORMS.map((f) => (
                  <SelectItem key={f.value} value={f.value}>
                    {f.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label>Uso *</Label>
            <Select
              value={form.watch(`formulas.${index}.usageType`)}
              onValueChange={(v) =>
                form.setValue(`formulas.${index}.usageType`, v as 'internal' | 'external')
              }
            >
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {USAGE_OPTIONS.map((u) => (
                  <SelectItem key={u.value} value={u.value}>
                    {u.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>

        {/* Componentes */}
        <div className="space-y-3 rounded-lg border p-3">
          <div className="flex items-center justify-between">
            <Label className="text-sm font-semibold">Componentes</Label>
            <span className="text-xs text-muted-foreground">{fields.length}/20</span>
          </div>

          {fields.map((field, ci) => (
            <div key={field.id} className="grid grid-cols-1 gap-2 sm:grid-cols-12">
              <div className="sm:col-span-5">
                <MagistralComponentInput
                  name={`formulas.${index}.components.${ci}.substance`}
                  value={form.watch(`formulas.${index}.components.${ci}.substance`) ?? ''}
                  onChange={(v) => {
                    form.setValue(`formulas.${index}.components.${ci}.substance`, v)
                    // Digitou por cima da sugestão: o vínculo com o catálogo deixa de valer.
                    form.setValue(`formulas.${index}.components.${ci}.magistralComponentId`, undefined)
                  }}
                  onSelect={(comp, dose) => {
                    form.setValue(`formulas.${index}.components.${ci}.substance`, comp.name)
                    form.setValue(`formulas.${index}.components.${ci}.magistralComponentId`, comp.id)
                    if (comp.defaultUnit) {
                      form.setValue(`formulas.${index}.components.${ci}.unit`, comp.defaultUnit)
                    }
                    // Mineral se prescreve em ELEMENTO. O sinal vem do catálogo e não do
                    // percentual: quelato de cobre não tem percentual fixo (muda de fornecedor) e
                    // mesmo assim a dose é do elemento. Enquanto isso era deduzido de
                    // `elementalPercent`, justamente cobre, boro, manganês e vanádio entravam na
                    // receita como se a dose fosse do pó.
                    // Só `doseAsElemental` decide. Deduzir de `elementalPercent` marcava também
                    // o palmitato de ascorbila, que tem percentual de ativo conhecido (43%) mas
                    // se prescreve pela substância nomeada: "100 mg" viravam 233 mg de insumo.
                    if (comp.doseAsElemental) {
                      form.setValue(`formulas.${index}.components.${ci}.asElemental`, true)
                    }
                    // Dose escolhida no painel manda; sem ela, a dose usual só preenche campo
                    // vazio — sugestão não sobrescreve decisão já tomada.
                    const atual = Number(
                      form.getValues(`formulas.${index}.components.${ci}.quantity`)
                    )
                    if (dose) {
                      form.setValue(`formulas.${index}.components.${ci}.quantity`, dose)
                    } else if (comp.usualDose && !atual) {
                      form.setValue(`formulas.${index}.components.${ci}.quantity`, comp.usualDose)
                    }
                  }}
                />
              </div>
              <div className="sm:col-span-2">
                <Input
                  type="number"
                  step="any"
                  min="0"
                  placeholder="Qtd"
                  {...form.register(`formulas.${index}.components.${ci}.quantity`)}
                />
              </div>
              <div className="sm:col-span-2">
                <Select
                  value={form.watch(`formulas.${index}.components.${ci}.unit`)}
                  onValueChange={(v) =>
                    form.setValue(`formulas.${index}.components.${ci}.unit`, v)
                  }
                >
                  <SelectTrigger>
                    <SelectValue placeholder="un" />
                  </SelectTrigger>
                  <SelectContent>
                    {COMPONENT_UNITS.map((u) => (
                      <SelectItem key={u} value={u}>
                        {u}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              {/* Mineral quelado/diluído: a dose é do ELEMENTO, e isso precisa estar dito na
                  linha — é o que a farmácia converte para a massa do insumo. */}
              <div className="flex items-center sm:col-span-2">
                {form.watch(`formulas.${index}.components.${ci}.asElemental`) ? (
                  <Badge
                    variant="outline"
                    className="cursor-pointer border-sky-200 bg-sky-50 text-[11px] text-sky-700"
                    onClick={() =>
                      form.setValue(`formulas.${index}.components.${ci}.asElemental`, false)
                    }
                    title="A dose é do elemento. Clique para passar a contar como insumo."
                  >
                    dose do elemento
                  </Badge>
                ) : (
                  <Badge
                    variant="outline"
                    className="cursor-pointer text-[11px] text-muted-foreground"
                    onClick={() =>
                      form.setValue(`formulas.${index}.components.${ci}.asElemental`, true)
                    }
                    title="A dose é do insumo. Clique para marcar que é do elemento."
                  >
                    dose do insumo
                  </Badge>
                )}
              </div>
              <div className="flex items-center justify-end gap-1 sm:col-span-1">
                {fields.length > 1 && (
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    onClick={() => remove(ci)}
                    className="text-destructive hover:text-destructive"
                    title="Remover componente"
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                )}
              </div>
            </div>
          ))}

          {fields.length < 20 && (
            <Button
              type="button"
              variant="outline"
              size="sm"
              className="w-full"
              onClick={() =>
                append(
                  emptyFormulaComponent(
                    PHARMACEUTICAL_FORMS.find(
                      (f) => f.value === form.getValues(`formulas.${index}.pharmaceuticalForm`)
                    )?.componentUnit
                  )
                )
              }
            >
              <Plus className="mr-2 h-4 w-4" />
              Adicionar componente
            </Button>
          )}

          {formula && (
            <FormulaCheckPanel
              formula={formula}
              onSetCategory={(substancia, categoria) => {
                const alvo = (form.getValues(`formulas.${index}.components`) ?? []).findIndex(
                  (c) => c.substance?.trim().toLowerCase() === substancia.toLowerCase()
                )
                if (alvo >= 0) {
                  form.setValue(
                    `formulas.${index}.components.${alvo}.category`,
                    categoria as never,
                    { shouldDirty: true }
                  )
                }
              }}
              onMarkElemental={(substancia, elemento) => {
                const alvo = (form.getValues(`formulas.${index}.components`) ?? []).findIndex(
                  (c) => c.substance?.trim().toLowerCase() === substancia.toLowerCase()
                )
                if (alvo >= 0) {
                  form.setValue(`formulas.${index}.components.${alvo}.asElemental`, elemento)
                }
              }}
            />
          )}

          {form.formState.errors.formulas?.[index]?.components && (
            <p className="text-sm text-destructive">
              {form.formState.errors.formulas[index]?.components?.message ??
                'Revise os componentes desta fórmula.'}
            </p>
          )}
        </div>

        <div className="space-y-2">
          <Label>Veículo / excipiente</Label>
          <Input
            placeholder="Ex: Excipiente qsp 1 cápsula"
            {...form.register(`formulas.${index}.vehicle`)}
          />
        </div>

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
          <div className="space-y-2">
            <Label>Aviar *</Label>
            <Input
              type="number"
              step="any"
              min="0"
              placeholder="60"
              {...form.register(`formulas.${index}.quantityToDispense`)}
            />
          </div>
          <div className="space-y-2">
            <Label>Unidade *</Label>
            <Input
              placeholder="cápsulas"
              {...form.register(`formulas.${index}.quantityUnit`)}
            />
          </div>
          <div className="space-y-2">
            <Label>Duração (dias)</Label>
            <Input type="number" min="0" {...form.register(`formulas.${index}.duration`)} />
          </div>
        </div>

        {controlled && (
          <div className="space-y-2">
            <Label>Quantidade por extenso *</Label>
            <Input
              placeholder="Ex: sessenta"
              {...form.register(`formulas.${index}.quantityInWords`)}
            />
            <p className="flex items-start gap-1 text-xs text-amber-700">
              <AlertCircle className="mt-0.5 h-3 w-3 shrink-0" />
              Obrigatória porque esta fórmula tem substância de controle especial.
            </p>
          </div>
        )}

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div className="space-y-2">
            <Label>Posologia *</Label>
            <Input
              placeholder="Ex: 1 cápsula ao deitar"
              {...form.register(`formulas.${index}.posology`)}
            />
          </div>
          <div className="space-y-2">
            <Label>Orientações desta fórmula (opcional)</Label>
            <Textarea
              rows={2}
              placeholder="Ex: manter refrigerado"
              {...form.register(`formulas.${index}.instructions`)}
            />
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
