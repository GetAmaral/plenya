'use client'

import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { Loader2, Plus, Save, Search, Trash2 } from 'lucide-react'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
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
import { PageHeader } from '@/components/layout/page-header'
import { useRequireAuth } from '@/lib/use-auth'
import { apiClient } from '@/lib/api-client'
import {
  createFormulaTemplate,
  deleteFormulaTemplate,
  listFormulaTemplates,
  updateFormulaTemplate,
  type DoseRule,
  type DoseBand,
  type FormulaTemplate,
  type FormulaTemplateComponent,
} from '@/lib/api/magistral-templates'
import { PHARMACEUTICAL_FORMS, USAGE_OPTIONS } from '@/lib/validations/compounded-prescription'
import { CATEGORY_OPTIONS_SHORT } from '@/lib/validations/prescription'

interface LabTest {
  id: string
  code: string
  name: string
  unit?: string
}

const novaFormula = (): FormulaTemplate => ({
  id: '',
  name: '',
  pharmaceuticalForm: 'cápsula',
  usageType: 'internal',
  route: 'oral',
  vehicle: 'Excipiente qsp 1 cápsula',
  quantityToDispense: 60,
  quantityUnit: 'cápsulas',
  posology: '',
  usageCount: 0,
  isActive: true,
  components: [{ substance: '', quantity: 0, unit: 'mg', category: 'simple' }],
})

/**
 * Fórmulas-base e suas regras de dose.
 *
 * É aqui que a dose dinâmica se cadastra, e o formulário força o que a regra precisa ter: piso e
 * teto. Sem os dois, a regra não salva — é o que garante que peso errado no prontuário ou exame
 * em unidade trocada não consigam sugerir dose absurda lá na frente.
 */
export default function FormulasBasePage() {
  useRequireAuth()
  const queryClient = useQueryClient()
  const [term, setTerm] = useState('')
  const [editing, setEditing] = useState<FormulaTemplate | null>(null)

  const { data: templates = [], isFetching } = useQuery({
    queryKey: ['formula-templates-admin', term],
    queryFn: () => listFormulaTemplates(term),
  })

  // A regra por exame aponta para o `code` do catálogo de exames — nenhum código clínico
  // hardcoded no front nem no back.
  const { data: labTests = [] } = useQuery({
    queryKey: ['lab-test-definitions-lite'],
    queryFn: () => apiClient.get<LabTest[]>('/api/v1/lab-tests/definitions'),
  })

  const save = useMutation({
    mutationFn: (t: FormulaTemplate) => {
      const body = {
        name: t.name,
        indication: t.indication,
        indicationBullets: t.indicationBullets,
        pharmaceuticalForm: t.pharmaceuticalForm,
        usageType: t.usageType,
        route: t.route,
        vehicle: t.vehicle,
        quantityToDispense: t.quantityToDispense,
        quantityUnit: t.quantityUnit,
        posology: t.posology,
        duration: t.duration,
        instructions: t.instructions,
        notes: t.notes,
        components: (t.components ?? []).filter((c) => c.substance.trim()),
      }
      return t.id ? updateFormulaTemplate(t.id, body) : createFormulaTemplate(body)
    },
    onSuccess: () => {
      toast.success('Fórmula-base salva')
      setEditing(null)
      queryClient.invalidateQueries({ queryKey: ['formula-templates-admin'] })
      queryClient.invalidateQueries({ queryKey: ['formula-templates'] })
    },
    onError: (e: any) => toast.error('Não deu para salvar', { description: e?.message }),
  })

  const remove = useMutation({
    mutationFn: deleteFormulaTemplate,
    onSuccess: () => {
      toast.success('Fórmula-base removida')
      queryClient.invalidateQueries({ queryKey: ['formula-templates-admin'] })
    },
  })

  const patch = (changes: Partial<FormulaTemplate>) =>
    setEditing((prev) => (prev ? { ...prev, ...changes } : prev))

  const patchComponent = (index: number, changes: Partial<FormulaTemplateComponent>) =>
    setEditing((prev) => {
      if (!prev) return prev
      const components = [...(prev.components ?? [])]
      components[index] = { ...components[index], ...changes }
      return { ...prev, components }
    })

  const patchRule = (index: number, changes: Partial<DoseRule>) =>
    setEditing((prev) => {
      if (!prev) return prev
      const components = [...(prev.components ?? [])]
      const atual = components[index].rule ?? { kind: 'fixed', minDose: 0, maxDose: 0 }
      components[index] = { ...components[index], rule: { ...atual, ...changes } }
      return { ...prev, components }
    })

  const patchBand = (index: number, bandIndex: number, changes: Partial<DoseBand>) =>
    setEditing((prev) => {
      if (!prev) return prev
      const components = [...(prev.components ?? [])]
      const rule = components[index].rule
      if (!rule) return prev
      const bands = [...(rule.bands ?? [])]
      bands[bandIndex] = { ...bands[bandIndex], ...changes }
      components[index] = { ...components[index], rule: { ...rule, bands } }
      return { ...prev, components }
    })

  const addBand = (index: number) =>
    setEditing((prev) => {
      if (!prev) return prev
      const components = [...(prev.components ?? [])]
      const rule = components[index].rule
      if (!rule) return prev
      const bands = [...(rule.bands ?? []), { lowerBound: null, upperBound: null, dose: 0, label: '' }]
      components[index] = { ...components[index], rule: { ...rule, bands } }
      return { ...prev, components }
    })

  const removeBand = (index: number, bandIndex: number) =>
    setEditing((prev) => {
      if (!prev) return prev
      const components = [...(prev.components ?? [])]
      const rule = components[index].rule
      if (!rule) return prev
      const bands = (rule.bands ?? []).filter((_: DoseBand, k: number) => k !== bandIndex)
      components[index] = { ...components[index], rule: { ...rule, bands } }
      return { ...prev, components }
    })

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Admin', href: '/admin' }, { label: 'Fórmulas-base' }]}
        title="Fórmulas-base"
        description="As fórmulas prontas que você reusa, e as regras que sugerem dose a partir do prontuário"
      />

      <Card>
        <CardHeader className="flex flex-row items-center justify-between gap-3 space-y-0">
          <div>
            <CardTitle>Suas fórmulas</CardTitle>
            <CardDescription>Mais usadas primeiro.</CardDescription>
          </div>
          <Button onClick={() => setEditing(novaFormula())}>
            <Plus className="mr-2 h-4 w-4" />
            Nova fórmula-base
          </Button>
        </CardHeader>
        <CardContent className="space-y-3">
          <div className="flex items-center gap-2">
            <Search className="h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Buscar por nome ou indicação..."
              value={term}
              onChange={(e) => setTerm(e.target.value)}
            />
            {isFetching && <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />}
          </div>

          <div className="divide-y rounded-lg border">
            {templates.map((t) => (
              <div key={t.id} className="flex flex-wrap items-center justify-between gap-2 p-3">
                <div className="min-w-0">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="font-medium">{t.name}</span>
                    <Badge variant="outline" className="text-xs">
                      {t.pharmaceuticalForm}
                    </Badge>
                    {(t.components ?? []).some((c) => c.rule) && (
                      <Badge variant="outline" className="bg-sky-50 text-xs text-sky-700">
                        com regra de dose
                      </Badge>
                    )}
                  </div>
                  {t.indication && (
                    <p className="mt-1 line-clamp-2 text-sm text-muted-foreground">{t.indication}</p>
                  )}
                  <p className="text-xs text-muted-foreground">
                    {t.components?.length ?? 0} componentes · {t.usageCount} receita(s)
                  </p>
                </div>
                <div className="flex gap-2">
                  <Button variant="outline" size="sm" onClick={() => setEditing({ ...t })}>
                    Editar
                  </Button>
                  <Button
                    variant="ghost"
                    size="sm"
                    className="text-destructive"
                    onClick={() => remove.mutate(t.id)}
                  >
                    <Trash2 className="h-4 w-4" />
                  </Button>
                </div>
              </div>
            ))}
            {templates.length === 0 && (
              <p className="p-6 text-center text-sm text-muted-foreground">
                Nenhuma fórmula-base ainda.
              </p>
            )}
          </div>
        </CardContent>
      </Card>

      {editing && (
        <Card>
          <CardHeader>
            <CardTitle>{editing.id ? editing.name : 'Nova fórmula-base'}</CardTitle>
            <CardDescription>
              A indicação é o que faz esta fórmula aparecer quando você busca pelo que quer tratar.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-5">
            <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
              <div className="space-y-2">
                <Label>Nome *</Label>
                <Input
                  value={editing.name}
                  placeholder="Fórmula do sono"
                  onChange={(e) => patch({ name: e.target.value })}
                />
              </div>
              <div className="space-y-2">
                <Label>Forma farmacêutica *</Label>
                <Select
                  value={editing.pharmaceuticalForm}
                  onValueChange={(v) => {
                    const preset = PHARMACEUTICAL_FORMS.find((f) => f.value === v)
                    patch({
                      pharmaceuticalForm: v,
                      ...(preset
                        ? {
                            quantityUnit: preset.quantityUnit,
                            usageType: preset.usageType,
                            route: preset.route,
                            vehicle: preset.vehicle,
                          }
                        : {}),
                    })
                  }}
                >
                  <SelectTrigger>
                    <SelectValue />
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
                  value={editing.usageType}
                  onValueChange={(v) => patch({ usageType: v as 'internal' | 'external' })}
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

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
              <div className="space-y-2">
                <Label>Indicação em tópicos</Label>
                <Textarea
                  rows={4}
                  placeholder={'Um por linha:\ninsônia de manutenção\nhigiene do sono'}
                  value={editing.indicationBullets ?? ''}
                  onChange={(e) => patch({ indicationBullets: e.target.value })}
                />
              </div>
              <div className="space-y-2">
                <Label>Indicação (texto)</Label>
                <Textarea
                  rows={4}
                  value={editing.indication ?? ''}
                  onChange={(e) => patch({ indication: e.target.value })}
                />
              </div>
            </div>

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-4">
              <div className="space-y-2">
                <Label>Veículo</Label>
                <Input value={editing.vehicle ?? ''} onChange={(e) => patch({ vehicle: e.target.value })} />
              </div>
              <div className="space-y-2">
                <Label>Aviar</Label>
                <Input
                  type="number"
                  step="any"
                  value={editing.quantityToDispense}
                  onChange={(e) => patch({ quantityToDispense: Number(e.target.value) })}
                />
              </div>
              <div className="space-y-2">
                <Label>Unidade</Label>
                <Input
                  value={editing.quantityUnit}
                  onChange={(e) => patch({ quantityUnit: e.target.value })}
                />
              </div>
              <div className="space-y-2">
                <Label>Posologia</Label>
                <Input
                  value={editing.posology ?? ''}
                  placeholder="1 cápsula ao deitar"
                  onChange={(e) => patch({ posology: e.target.value })}
                />
              </div>
            </div>

            <div className="space-y-3">
              <Label className="text-sm font-semibold">Componentes e regras de dose</Label>
              {(editing.components ?? []).map((c, i) => (
                <div key={i} className="space-y-3 rounded-lg border p-3">
                  <div className="grid grid-cols-1 gap-2 sm:grid-cols-12">
                    <Input
                      className="sm:col-span-5"
                      placeholder="Substância"
                      value={c.substance}
                      onChange={(e) => patchComponent(i, { substance: e.target.value })}
                    />
                    <Input
                      className="sm:col-span-2"
                      type="number"
                      step="any"
                      placeholder="Dose"
                      value={c.quantity || ''}
                      onChange={(e) => patchComponent(i, { quantity: Number(e.target.value) })}
                    />
                    <Input
                      className="sm:col-span-2"
                      placeholder="un"
                      value={c.unit}
                      onChange={(e) => patchComponent(i, { unit: e.target.value })}
                    />
                    <Select
                      value={c.category ?? 'simple'}
                      onValueChange={(v) => patchComponent(i, { category: v as never })}
                    >
                      <SelectTrigger className="sm:col-span-2">
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        {CATEGORY_OPTIONS_SHORT.map((o) => (
                          <SelectItem key={o.value} value={o.value}>
                            {o.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      className="text-destructive sm:col-span-1"
                      onClick={() =>
                        patch({ components: (editing.components ?? []).filter((_, j) => j !== i) })
                      }
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>

                  {!c.rule ? (
                    <Button
                      type="button"
                      variant="outline"
                      size="sm"
                      onClick={() =>
                        patchRule(i, { kind: 'fixed', minDose: c.quantity || 1, maxDose: c.quantity || 1 })
                      }
                    >
                      <Plus className="mr-2 h-4 w-4" />
                      Regra de dose
                    </Button>
                  ) : (
                    <div className="space-y-3 rounded-md border border-dashed p-3">
                      <div className="flex flex-wrap items-center gap-2">
                        <Select
                          value={c.rule.kind}
                          onValueChange={(v) =>
                            // O default que aparece no select precisa existir no estado: sem
                            // isto, "menor que" ficava só na tela e o payload ia sem a comparação.
                            patchRule(i, {
                              kind: v as DoseRule['kind'],
                              ...(v === 'lab_threshold' && !c.rule?.labOperator
                                ? { labOperator: 'lt' as const }
                                : {}),
                              // Faixa nenhuma não é regra: a primeira já nasce na tela.
                              ...(v === 'lab_band' && !c.rule?.bands?.length
                                ? { bands: [{ upperBound: null, lowerBound: null, dose: 0, label: '' }] }
                                : {}),
                            })
                          }
                        >
                          <SelectTrigger className="w-56">
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            <SelectItem value="fixed">Dose fixa</SelectItem>
                            <SelectItem value="per_kg">Por peso (mg/kg)</SelectItem>
                            <SelectItem value="lab_threshold">Conforme exame (um corte)</SelectItem>
                            <SelectItem value="lab_band">Conforme exame (faixas)</SelectItem>
                          </SelectContent>
                        </Select>
                        <Button
                          type="button"
                          variant="ghost"
                          size="sm"
                          className="text-destructive"
                          onClick={() => patchComponent(i, { rule: null })}
                        >
                          remover regra
                        </Button>
                      </div>

                      {c.rule.kind === 'fixed' && (
                        <div className="w-40 space-y-1">
                          <Label className="text-xs">Dose fixa</Label>
                          <Input
                            type="number"
                            step="any"
                            value={c.rule.fixedDose ?? ''}
                            onChange={(e) => patchRule(i, { fixedDose: Number(e.target.value) })}
                          />
                        </div>
                      )}

                      {c.rule.kind === 'per_kg' && (
                        <div className="w-40 space-y-1">
                          <Label className="text-xs">Dose por kg</Label>
                          <Input
                            type="number"
                            step="any"
                            value={c.rule.perKg ?? ''}
                            onChange={(e) => patchRule(i, { perKg: Number(e.target.value) })}
                          />
                        </div>
                      )}

                      {c.rule.kind === 'lab_threshold' && (
                        <div className="grid grid-cols-1 gap-2 sm:grid-cols-5">
                          <div className="space-y-1 sm:col-span-2">
                            <Label className="text-xs">Exame</Label>
                            <Select
                              value={c.rule.labCode ?? ''}
                              onValueChange={(v) =>
                                patchRule(i, {
                                  labCode: v,
                                  labUnit: labTests.find((t) => t.code === v)?.unit ?? '',
                                })
                              }
                            >
                              <SelectTrigger>
                                <SelectValue placeholder="Escolha o exame" />
                              </SelectTrigger>
                              <SelectContent>
                                {labTests.slice(0, 300).map((t) => (
                                  <SelectItem key={t.code} value={t.code}>
                                    {t.name}
                                  </SelectItem>
                                ))}
                              </SelectContent>
                            </Select>
                          </div>
                          <div className="space-y-1">
                            <Label className="text-xs">Comparação</Label>
                            <Select
                              value={c.rule.labOperator ?? 'lt'}
                              onValueChange={(v) => patchRule(i, { labOperator: v as never })}
                            >
                              <SelectTrigger>
                                <SelectValue />
                              </SelectTrigger>
                              <SelectContent>
                                <SelectItem value="lt">menor que</SelectItem>
                                <SelectItem value="lte">menor ou igual</SelectItem>
                                <SelectItem value="gt">maior que</SelectItem>
                                <SelectItem value="gte">maior ou igual</SelectItem>
                              </SelectContent>
                            </Select>
                          </div>
                          <div className="space-y-1">
                            <Label className="text-xs">Limiar</Label>
                            <Input
                              type="number"
                              step="any"
                              value={c.rule.labThreshold ?? ''}
                              onChange={(e) => patchRule(i, { labThreshold: Number(e.target.value) })}
                            />
                          </div>
                          <div className="space-y-1">
                            <Label className="text-xs">Dose se sim</Label>
                            <Input
                              type="number"
                              step="any"
                              value={c.rule.doseIfTrue ?? ''}
                              onChange={(e) => patchRule(i, { doseIfTrue: Number(e.target.value) })}
                            />
                          </div>
                          <div className="space-y-1 sm:col-span-1">
                            <Label className="text-xs">Dose se não</Label>
                            <Input
                              type="number"
                              step="any"
                              value={c.rule.doseIfFalse ?? ''}
                              onChange={(e) => patchRule(i, { doseIfFalse: Number(e.target.value) })}
                            />
                          </div>
                        </div>
                      )}


                      {c.rule.kind === 'lab_band' && (
                        <div className="space-y-2">
                          <div className="grid grid-cols-1 gap-2 sm:grid-cols-3">
                            <div className="space-y-1 sm:col-span-2">
                              <Label className="text-xs">Exame</Label>
                              <Select
                                value={c.rule.labCode ?? ''}
                                onValueChange={(v) =>
                                  // A unidade vem junto do exame escolhido, não digitada: é ela
                                  // que faz o motor recusar resultado gravado em outra unidade.
                                  patchRule(i, {
                                    labCode: v,
                                    labUnit: labTests.find((t) => t.code === v)?.unit ?? '',
                                  })
                                }
                              >
                                <SelectTrigger>
                                  <SelectValue placeholder="Escolha o exame" />
                                </SelectTrigger>
                                <SelectContent>
                                  {labTests.slice(0, 300).map((t) => (
                                    <SelectItem key={t.code} value={t.code}>
                                      {t.name}
                                      {t.unit ? ` (${t.unit})` : ''}
                                    </SelectItem>
                                  ))}
                                </SelectContent>
                              </Select>
                            </div>
                            <div className="space-y-1">
                              <Label className="text-xs">Unidade da regra</Label>
                              <Input
                                value={c.rule.labUnit ?? ''}
                                placeholder="ng/mL"
                                onChange={(e) => patchRule(i, { labUnit: e.target.value })}
                              />
                            </div>
                          </div>

                          <div className="space-y-2 rounded-md bg-muted/40 p-2">
                            {(c.rule.bands ?? []).map((b, bi) => (
                              <div key={bi} className="grid grid-cols-2 gap-2 sm:grid-cols-9">
                                <div className="space-y-1 sm:col-span-2">
                                  <Label className="text-[11px]">Acima de</Label>
                                  <Input
                                    type="number"
                                    step="any"
                                    placeholder="sem piso"
                                    value={b.lowerBound ?? ''}
                                    onChange={(e) =>
                                      patchBand(i, bi, {
                                        lowerBound: e.target.value === '' ? null : Number(e.target.value),
                                      })
                                    }
                                  />
                                </div>
                                <div className="space-y-1 sm:col-span-2">
                                  <Label className="text-[11px]">Até (inclusive)</Label>
                                  <Input
                                    type="number"
                                    step="any"
                                    placeholder="sem teto"
                                    value={b.upperBound ?? ''}
                                    onChange={(e) =>
                                      patchBand(i, bi, {
                                        upperBound: e.target.value === '' ? null : Number(e.target.value),
                                      })
                                    }
                                  />
                                </div>
                                <div className="space-y-1 sm:col-span-2">
                                  <Label className="text-[11px]">Dose</Label>
                                  <Input
                                    type="number"
                                    step="any"
                                    value={b.dose || ''}
                                    onChange={(e) => patchBand(i, bi, { dose: Number(e.target.value) })}
                                  />
                                </div>
                                <div className="space-y-1 sm:col-span-2">
                                  <Label className="text-[11px]">Como chamar</Label>
                                  <Input
                                    value={b.label ?? ''}
                                    placeholder="insuficiência"
                                    onChange={(e) => patchBand(i, bi, { label: e.target.value })}
                                  />
                                </div>
                                <div className="flex items-end">
                                  <Button
                                    type="button"
                                    variant="ghost"
                                    size="sm"
                                    className="text-destructive"
                                    onClick={() => removeBand(i, bi)}
                                  >
                                    remover
                                  </Button>
                                </div>
                              </div>
                            ))}
                            <Button
                              type="button"
                              variant="outline"
                              size="sm"
                              onClick={() => addBand(i)}
                            >
                              <Plus className="mr-2 h-4 w-4" />
                              faixa
                            </Button>
                            <p className="text-[11px] text-muted-foreground">
                              A faixa vai de <em>acima de</em> até <em>até</em>, inclusive o valor
                              do topo, igual às faixas do escore. Deixar em branco é sem piso ou
                              sem teto. Faixas não podem se sobrepor; buraco entre elas é
                              permitido e o motor responde dizendo que não há conduta cadastrada.
                            </p>
                          </div>
                        </div>
                      )}

                      <div className="grid grid-cols-1 gap-2 sm:grid-cols-5">
                        <div className="space-y-1">
                          <Label className="text-xs">Piso *</Label>
                          <Input
                            type="number"
                            step="any"
                            value={c.rule.minDose || ''}
                            onChange={(e) => patchRule(i, { minDose: Number(e.target.value) })}
                          />
                        </div>
                        <div className="space-y-1">
                          <Label className="text-xs">Teto *</Label>
                          <Input
                            type="number"
                            step="any"
                            value={c.rule.maxDose || ''}
                            onChange={(e) => patchRule(i, { maxDose: Number(e.target.value) })}
                          />
                        </div>
                        <div className="space-y-1">
                          <Label className="text-xs">Arredondar para</Label>
                          <Input
                            type="number"
                            step="any"
                            placeholder="500"
                            value={c.rule.roundTo ?? ''}
                            onChange={(e) =>
                              patchRule(i, {
                                roundTo: e.target.value === '' ? undefined : Number(e.target.value),
                              })
                            }
                          />
                        </div>
                        <div className="space-y-1">
                          <Label className="text-xs">Dado válido por (dias)</Label>
                          <Input
                            type="number"
                            value={c.rule.maxDataAgeDays ?? 365}
                            onChange={(e) => patchRule(i, { maxDataAgeDays: Number(e.target.value) })}
                          />
                        </div>
                        <div className="space-y-1">
                          <Label className="text-xs">Observação</Label>
                          <Input
                            value={c.rule.note ?? ''}
                            placeholder="reavaliar em 90 dias"
                            onChange={(e) => patchRule(i, { note: e.target.value })}
                          />
                        </div>
                      </div>
                      <p className="text-xs text-muted-foreground">
                        Piso e teto são obrigatórios: é o que impede peso ou exame errado no
                        prontuário de virar dose absurda. Dado mais velho que o prazo não sugere
                        nada.
                      </p>
                    </div>
                  )}
                </div>
              ))}

              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={() =>
                  patch({
                    components: [
                      ...(editing.components ?? []),
                      { substance: '', quantity: 0, unit: 'mg', category: 'simple' },
                    ],
                  })
                }
              >
                <Plus className="mr-2 h-4 w-4" />
                Adicionar componente
              </Button>
            </div>

            <div className="flex gap-3">
              <Button onClick={() => save.mutate(editing)} disabled={save.isPending}>
                {save.isPending ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <Save className="mr-2 h-4 w-4" />
                )}
                Salvar
              </Button>
              <Button variant="ghost" onClick={() => setEditing(null)}>
                Cancelar
              </Button>
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  )
}
