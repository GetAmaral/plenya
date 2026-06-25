'use client'

import { useState, useEffect, useRef, type Dispatch, type SetStateAction } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip'
import { HelpCircle, CheckCircle2, AlertCircle, ChevronDown, ChevronRight, Plus } from 'lucide-react'
import { cn } from '@/lib/utils'
import { calcAge } from '@/lib/format-date'
import { toast } from 'sonner'
import type { AnamnesisTemplate, AnamnesisTemplateItem } from '@/lib/api/anamnesis-templates'
import type { ScoreGroup, ScoreSubgroup, ScoreItem, ScoreLevel } from '@/lib/api/score-api'
import type { AnamnesisItemFormValue } from './AnamnesisTemplateItemsForm'
import type { Patient } from '@/lib/auth-store'
import { AnamnesisItemHistory } from './AnamnesisItemHistory'
import { ScaleWidget } from './ScaleWidget'
import { getScaleDef, pickWordRecallSet, type ChosenWord } from '@plenya/domain'

// Evaluates whether a numeric value satisfies a ScoreLevel's operator/limits
function evaluatesTrue(value: number, level: ScoreLevel): boolean {
  const lower = level.lowerLimit != null ? parseFloat(level.lowerLimit) : null
  const upper = level.upperLimit != null ? parseFloat(level.upperLimit) : null
  switch (level.operator) {
    case '=':       return lower !== null && value === lower
    case '>':       return lower !== null && value > lower
    case '>=':      return lower !== null && value >= lower
    case '<':       return lower !== null && value < lower
    case '<=':      return lower !== null && value <= lower
    case 'between': return lower !== null && upper !== null && value >= lower && value <= upper
    default: return false
  }
}

// Returns the level number that matches a numeric value, or undefined if none
function detectLevel(value: number, levels: ScoreLevel[]): number | undefined {
  return levels.find(l => evaluatesTrue(value, l))?.level
}

// Conversões entre as respostas do widget (chaves numéricas) e o JSON persistido (chaves string)
function toStringKeys(answers: Record<number, number>): Record<string, number> {
  const out: Record<string, number> = {}
  for (const [k, v] of Object.entries(answers)) out[k] = v
  return out
}
function toNumberKeys(answers: Record<string, number> | undefined): Record<number, number> | undefined {
  if (!answers) return undefined
  const out: Record<number, number> = {}
  for (const [k, v] of Object.entries(answers)) out[Number(k)] = v
  return out
}

// Resolve o conjunto de palavras do teste de evocação (Dubois) UMA vez por anamnese:
// reaproveita as palavras já persistidas (anamnese carregada) ou sorteia uma forma nova.
// Imediato e tardio compartilham o mesmo conjunto (mesmas categorias).
function resolveWordRecallSet(
  template: AnamnesisTemplate,
  initialValues: AnamnesisItemFormValue[],
): ChosenWord[] | null {
  if (!template.items) return null
  let spec: Extract<NonNullable<ReturnType<typeof getScaleDef>>['administration'], { type: 'word_recall' }> | null = null
  const itemIds: string[] = []
  for (const ti of template.items) {
    const def = getScaleDef(ti.scoreItem?.anamneseItemCode)
    if (def?.administration?.type === 'word_recall') {
      if (!spec) spec = def.administration
      if (ti.scoreItem) itemIds.push(ti.scoreItem.id)
    }
  }
  if (!spec) return null
  for (const id of itemIds) {
    const words = initialValues.find((v) => v.scoreItemId === id)?.scaleResponses?.words
    if (words && words.length === spec.categories.length) {
      return spec.categories.map((c, i) => ({ category: c.category, word: words[i] }))
    }
  }
  return pickWordRecallSet(spec, (size) => Math.floor(Math.random() * size))
}

// Calcula idade a partir da data de nascimento.
// Delega para calcAge (TZ-safe: birthDate é data pura em meia-noite UTC; ler em
// horário local jogava -1 dia/idade perto do aniversário em BRT).
function calculateAge(birthDate: string): number {
  return calcAge(birthDate) ?? 0
}

// Verifica se um ScoreItem é aplicável ao paciente atual
function itemAppliesToPatient(scoreItem: ScoreItem, patient: Patient | null | undefined): boolean {
  if (!patient) return true

  const age = patient.age || calculateAge(patient.birthDate)

  // Filtro de gênero
  if (scoreItem.gender && scoreItem.gender !== 'not_applicable') {
    if (scoreItem.gender !== patient.gender) return false
  }

  // Filtro de faixa etária
  if (scoreItem.ageRangeMin !== undefined && age < scoreItem.ageRangeMin) return false
  if (scoreItem.ageRangeMax !== undefined && age > scoreItem.ageRangeMax) return false

  // Filtro de pós-menopausa (apenas para mulheres com dado disponível)
  if (scoreItem.postMenopause !== undefined) {
    if (patient.gender !== 'female') return false
    if (patient.menopause !== undefined && scoreItem.postMenopause !== patient.menopause) return false
  }

  return true
}

interface AnamnesisTemplateItemsRendererProps {
  template: AnamnesisTemplate
  initialValues?: AnamnesisItemFormValue[]
  onChange: (values: AnamnesisItemFormValue[]) => void
  compact?: boolean // For regular form (smaller UI)
  focusScoreItemId?: string | null
  patient?: Patient | null
  // ID do paciente p/ o histórico por item. Quando ausente, cai em patient?.id.
  // Útil quando o filtro demográfico usa `patient` mas o id confiável vem de outra fonte
  // (ex.: appt.patientId no workspace da consulta).
  patientId?: string | null
}

// Score level color classes - must be complete strings for Tailwind to detect them
// For full-size buttons (non-compact mode)
const LEVEL_SELECTED_CLASSES: Record<number, string> = {
  0: 'bg-red-100 text-red-900 border-red-500 shadow-md scale-105',
  1: 'bg-orange-100 text-orange-900 border-orange-500 shadow-md scale-105',
  2: 'bg-yellow-100 text-yellow-900 border-yellow-500 shadow-md scale-105',
  3: 'bg-blue-100 text-blue-900 border-blue-500 shadow-md scale-105',
  4: 'bg-green-100 text-green-900 border-green-500 shadow-md scale-105',
  5: 'bg-emerald-100 text-emerald-900 border-emerald-500 shadow-md scale-105',
  6: 'bg-gray-100 text-gray-900 border-gray-500 shadow-md scale-105',
}

const LEVEL_HOVER_CLASSES: Record<number, string> = {
  0: 'bg-background border-border hover:bg-red-200',
  1: 'bg-background border-border hover:bg-orange-200',
  2: 'bg-background border-border hover:bg-yellow-200',
  3: 'bg-background border-border hover:bg-blue-200',
  4: 'bg-background border-border hover:bg-green-200',
  5: 'bg-background border-border hover:bg-emerald-200',
  6: 'bg-background border-border hover:bg-gray-200',
}

// For compact mode buttons
const LEVEL_COMPACT_SELECTED_CLASSES: Record<number, string> = {
  0: 'border-red-500 bg-red-100 text-red-900',
  1: 'border-orange-500 bg-orange-100 text-orange-900',
  2: 'border-yellow-500 bg-yellow-100 text-yellow-900',
  3: 'border-blue-500 bg-blue-100 text-blue-900',
  4: 'border-green-500 bg-green-100 text-green-900',
  5: 'border-emerald-500 bg-emerald-100 text-emerald-900',
  6: 'border-gray-500 bg-gray-100 text-gray-900',
}

// Organize items by group and subgroup
interface OrganizedData {
  groups: Map<string, {
    group: ScoreGroup
    subgroups: Map<string, {
      subgroup: ScoreSubgroup
      items: { templateItem: AnamnesisTemplateItem; scoreItem: ScoreItem }[]
    }>
  }>
}

function organizeTemplateItems(template: AnamnesisTemplate, patient?: Patient | null): OrganizedData {
  const groups = new Map()

  if (!template.items) {
    return { groups }
  }

  const sortedItems = [...template.items].sort((a, b) => a.order - b.order)

  // Primeiro pass: identificar quais parentItemIds têm filhos no template
  // (pais com filhos presentes não devem aparecer — só os filhos corretos)
  const parentIdsWithChildrenInTemplate = new Set<string>()
  sortedItems.forEach((templateItem) => {
    const parentId = templateItem.scoreItem?.parentItemId
    if (parentId) {
      parentIdsWithChildrenInTemplate.add(parentId)
    }
  })

  sortedItems.forEach((templateItem) => {
    const scoreItem = templateItem.scoreItem

    if (!scoreItem || !scoreItem.subgroup) {
      return
    }

    // Suprimir o item pai se ele tiver filhos no template (exibir apenas os filhos)
    if (parentIdsWithChildrenInTemplate.has(scoreItem.id)) {
      return
    }

    // Para itens filhos (com parentItemId): aplicar filtro demográfico
    // Para itens sem hierarquia: sempre exibir (filtro demográfico também se aplicar)
    if (!itemAppliesToPatient(scoreItem, patient)) {
      return
    }

    const subgroup = scoreItem.subgroup
    const group = subgroup.group

    if (!group) {
      return
    }

    if (!groups.has(group.id)) {
      groups.set(group.id, {
        group,
        subgroups: new Map(),
      })
    }

    const groupData = groups.get(group.id)

    if (!groupData.subgroups.has(subgroup.id)) {
      groupData.subgroups.set(subgroup.id, {
        subgroup,
        items: [],
      })
    }

    const subgroupData = groupData.subgroups.get(subgroup.id)
    subgroupData.items.push({ templateItem, scoreItem })
  })

  return { groups }
}

export function AnamnesisTemplateItemsRenderer({
  template,
  initialValues = [],
  onChange,
  compact = false,
  focusScoreItemId,
  patient,
  patientId,
}: AnamnesisTemplateItemsRendererProps) {
  const historyPatientId = patientId ?? patient?.id ?? null
  const [values, setValues] = useState<Map<string, AnamnesisItemFormValue>>(() => {
    const newValues = new Map<string, AnamnesisItemFormValue>()
    initialValues.forEach((val) => {
      newValues.set(val.scoreItemId, val)
    })
    return newValues
  })
  const [focusedItemId, setFocusedItemId] = useState<string | null>(null)
  // Conjunto de palavras do Dubois, resolvido uma vez (persistido ou sorteado) e compartilhado
  // entre imediato e tardio nesta anamnese.
  const [wordRecallSet] = useState<ChosenWord[] | null>(() => resolveWordRecallSet(template, initialValues))

  const organized = organizeTemplateItems(template, patient)
  const itemRefs = useRef<Map<string, HTMLDivElement>>(new Map())
  const hasFocused = useRef(false)

  // --- Estado da visão compacta (mobile): accordion de grupos + itens expansíveis ---
  const [expandedGroups, setExpandedGroups] = useState<Set<string>>(
    () => new Set(Array.from(organized.groups.keys()).slice(0, 1)) // primeiro grupo aberto
  )
  const [expandedItems, setExpandedItems] = useState<Set<string>>(new Set())
  const [obsOpen, setObsOpen] = useState<Set<string>>(new Set())

  const toggleSetMember = (
    setter: Dispatch<SetStateAction<Set<string>>>,
    id: string
  ) =>
    setter((prev) => {
      const next = new Set(prev)
      next.has(id) ? next.delete(id) : next.add(id)
      return next
    })

  // Ao focar um item (deep link do health score), abre o grupo e o próprio item.
  useEffect(() => {
    if (!compact || !focusScoreItemId) return
    for (const [gid, g] of organized.groups) {
      for (const sg of g.subgroups.values()) {
        if (sg.items.some((it) => it.scoreItem.id === focusScoreItemId)) {
          setExpandedGroups((prev) => new Set(prev).add(gid))
          setExpandedItems((prev) => new Set(prev).add(focusScoreItemId))
          return
        }
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [compact, focusScoreItemId])

  // Chip que resume o valor atual do item na linha colapsada.
  const currentChip = (cur: AnamnesisItemFormValue | undefined, scoreItem: ScoreItem) => {
    if (cur?.selectedLevel !== undefined) {
      const lvl = cur.selectedLevel
      const name = (scoreItem.levels || []).find((l) => l.level === lvl)?.name
      return (
        <span
          className={cn(
            'whitespace-nowrap rounded-full border-[1.5px] px-2.5 py-0.5 text-[11px] font-bold',
            LEVEL_COMPACT_SELECTED_CLASSES[lvl] || LEVEL_COMPACT_SELECTED_CLASSES[6]
          )}
        >
          N{lvl}
          {name ? ` · ${name}` : ''}
        </span>
      )
    }
    if (cur?.numericValue !== undefined) {
      return (
        <span className="whitespace-nowrap rounded-full border-[1.5px] border-border bg-muted px-2.5 py-0.5 text-[11px] font-semibold text-foreground">
          {cur.numericValue}
          {scoreItem.unit ? ` ${scoreItem.unit}` : ''}
        </span>
      )
    }
    return (
      <span className="whitespace-nowrap rounded-full border-[1.5px] border-dashed border-border px-2.5 py-0.5 text-[11px] font-medium text-muted-foreground">
        preencher
      </span>
    )
  }

  // Auto-scroll and focus on specific item when coming from health score edit
  useEffect(() => {
    if (focusScoreItemId && !hasFocused.current) {
      hasFocused.current = true

      setTimeout(() => {
        const itemElement = itemRefs.current.get(focusScoreItemId)
        if (itemElement) {
          itemElement.scrollIntoView({ behavior: 'smooth', block: 'center' })

          // Find first focusable input in the item and focus it
          const input = itemElement.querySelector<HTMLInputElement | HTMLButtonElement>(
            'input, button[data-level-button]'
          )
          if (input) {
            setTimeout(() => {
              input.focus()
              if (input instanceof HTMLInputElement) {
                input.select()
              }
            }, 400)
          }
        }
      }, 600)
    }
  }, [focusScoreItemId])

  // Handle level selection (for items without maxSelect)
  const handleLevelSelect = (scoreItemId: string, level: ScoreLevel, order: number) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      if (existing && existing.selectedLevel === level.level) {
        // Deselect if clicking the same level
        if (!existing.textValue && existing.numericValue === undefined) {
          newValues.delete(scoreItemId)
        } else {
          newValues.set(scoreItemId, { ...existing, selectedLevel: undefined })
        }
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: existing?.numericValue,
          selectedLevel: level.level,
          textValue: existing?.textValue,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  // Resultado de um widget de escala (PHQ-9, GAD-7, …): seta o nível classificado
  // e guarda o detalhe por pergunta (scaleResponses). `level` undefined = incompleta/limpa.
  const handleScaleResult = (
    scoreItemId: string,
    result: { level: number | undefined; total: number; answers: Record<number, number>; words?: string[] },
    order: number,
  ) => {
    const { level, total, answers, words } = result
    const answeredCount = Object.keys(answers).length
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)
      if (level === undefined) {
        // Sem classificação (incompleta/limpa): mantém só se houver respostas parciais ou obs.
        if (answeredCount === 0 && !existing?.textValue) {
          newValues.delete(scoreItemId)
        } else {
          newValues.set(scoreItemId, {
            ...existing,
            scoreItemId,
            selectedLevel: undefined,
            numericValue: undefined,
            scaleResponses:
              answeredCount > 0 ? { answers: toStringKeys(answers), total, words } : undefined,
            order,
          })
        }
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          selectedLevel: level,
          textValue: existing?.textValue,
          scaleResponses: { answers: toStringKeys(answers), total, words },
          order,
        })
      }
      requestAnimationFrame(() => onChange(Array.from(newValues.values())))
      return newValues
    })
  }

  // Handle text change
  const handleTextChange = (scoreItemId: string, text: string, order: number) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      if (!text && existing?.numericValue === undefined && existing?.selectedLevel === undefined) {
        newValues.delete(scoreItemId)
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: existing?.numericValue,
          selectedLevel: existing?.selectedLevel,
          textValue: text || undefined,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  // Handle checkbox selection for maxSelect subgroups
  const handleCheckboxSelect = (
    scoreItemId: string,
    order: number,
    subgroupId: string,
    maxSelect: number,
    items: { templateItem: AnamnesisTemplateItem; scoreItem: ScoreItem }[]
  ) => {
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)

      // Get all selected items in this subgroup
      const subgroupItemIds = items.map(({ scoreItem }) => scoreItem.id)
      const selectedInSubgroup = Array.from(newValues.entries())
        .filter(([id]) => subgroupItemIds.includes(id))
        .filter(([, val]) => val.selectedLevel !== undefined)

      if (existing && existing.selectedLevel !== undefined) {
        // Deselect
        newValues.delete(scoreItemId)
      } else {
        // Check if we can select more
        if (selectedInSubgroup.length >= maxSelect) {
          // Show toast warning
          toast.warning(`Você pode selecionar no máximo ${maxSelect} ${maxSelect === 1 ? 'item' : 'itens'}`, {
            duration: 3000,
          })
          return prev
        }
        // Select with selectedLevel = 1 (represents "selected")
        newValues.set(scoreItemId, {
          scoreItemId,
          selectedLevel: 1,
          textValue: existing?.textValue,
          order,
        })
      }

      requestAnimationFrame(() => {
        onChange(Array.from(newValues.values()))
      })

      return newValues
    })
  }

  // Handle numeric input change (for items with unit — real measured value)
  const handleNumericChange = (scoreItemId: string, raw: string, order: number, levels: ScoreLevel[]) => {
    const num = raw === '' ? undefined : parseFloat(raw)
    const autoLevel = num !== undefined && !isNaN(num) ? detectLevel(num, levels) : undefined
    setValues((prev) => {
      const newValues = new Map(prev)
      const existing = newValues.get(scoreItemId)
      if (num === undefined && !existing?.textValue && existing?.selectedLevel === undefined) {
        newValues.delete(scoreItemId)
      } else {
        newValues.set(scoreItemId, {
          scoreItemId,
          numericValue: num,
          selectedLevel: autoLevel ?? existing?.selectedLevel,
          textValue: existing?.textValue,
          order,
        })
      }
      requestAnimationFrame(() => onChange(Array.from(newValues.values())))
      return newValues
    })
  }

  if (!template.items || template.items.length === 0) {
    return (
      <div className="text-center text-muted-foreground py-8">
        <p className={compact ? 'text-sm' : 'text-base'}>
          Este template não possui items configurados
        </p>
      </div>
    )
  }

  // --- Visão compacta (mobile): accordion denso + item expansível + histórico ---
  const compactView = (
    <div className="space-y-2.5">
      {Array.from(organized.groups.values()).map(({ group, subgroups }) => {
        const subArr = Array.from(subgroups.values())
        const allItems = subArr.flatMap((s) => s.items)
        const filled = allItems.filter(({ scoreItem }) => values.has(scoreItem.id)).length
        const total = allItems.length
        const gOpen = expandedGroups.has(group.id)

        return (
          <div key={group.id} className="overflow-hidden rounded-xl border bg-card">
            <button
              type="button"
              onClick={() => toggleSetMember(setExpandedGroups, group.id)}
              className="flex w-full items-center gap-2.5 px-3.5 py-3 text-left"
            >
              <span className="text-sm font-semibold text-foreground">{group.name}</span>
              <span className="ml-auto shrink-0 rounded-full bg-muted px-2 py-0.5 text-[11px] text-muted-foreground">
                {filled}/{total}
              </span>
              <ChevronDown
                className={cn(
                  'h-4 w-4 shrink-0 text-muted-foreground transition-transform',
                  gOpen && 'rotate-180'
                )}
              />
            </button>

            {gOpen && (
              <div className="border-t">
                {subArr.map(({ subgroup, items }) => {
                  const hasMaxSelect = subgroup.maxSelect > 0
                  const selectedCount = items.filter(
                    ({ scoreItem }) => values.get(scoreItem.id)?.selectedLevel !== undefined
                  ).length

                  return (
                    <div key={subgroup.id}>
                      {subArr.length > 1 && (
                        <div className="flex items-center justify-between bg-muted/40 px-3.5 py-1.5">
                          <span className="text-[11px] font-semibold uppercase tracking-wide text-muted-foreground">
                            {subgroup.name}
                          </span>
                          {hasMaxSelect && (
                            <span className="text-[10px] text-muted-foreground">
                              {selectedCount}/{subgroup.maxSelect}
                            </span>
                          )}
                        </div>
                      )}

                      {items.map(({ templateItem, scoreItem }) => {
                        const cur = values.get(scoreItem.id)
                        const levels = scoreItem.levels || []
                        const scaleDef = getScaleDef(scoreItem.anamneseItemCode)

                        // maxSelect → linha com checkbox (sem expandir)
                        if (hasMaxSelect) {
                          const isSel = cur?.selectedLevel !== undefined
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => {
                                if (el) itemRefs.current.set(scoreItem.id, el)
                              }}
                              className={cn(
                                'flex items-center gap-2.5 border-b px-3.5 py-2.5 last:border-b-0',
                                isSel && 'bg-primary/5'
                              )}
                            >
                              <Checkbox
                                checked={isSel}
                                onCheckedChange={() =>
                                  handleCheckboxSelect(
                                    scoreItem.id,
                                    templateItem.order,
                                    subgroup.id,
                                    subgroup.maxSelect,
                                    items
                                  )
                                }
                                className="h-5 w-5 shrink-0"
                              />
                              <span className="text-[13px] leading-tight text-foreground">
                                {scoreItem.name}
                              </span>
                            </div>
                          )
                        }

                        // Item padrão → linha colapsável
                        const itemOpen = expandedItems.has(scoreItem.id)
                        const showObs = obsOpen.has(scoreItem.id) || !!cur?.textValue

                        return (
                          <div
                            key={templateItem.id}
                            ref={(el) => {
                              if (el) itemRefs.current.set(scoreItem.id, el)
                            }}
                            className={cn('border-b last:border-b-0', itemOpen && 'bg-amber-50/40')}
                          >
                            <button
                              type="button"
                              onClick={() => toggleSetMember(setExpandedItems, scoreItem.id)}
                              className="flex w-full items-center gap-2.5 px-3.5 py-2.5 text-left"
                            >
                              <span className="min-w-0 flex-1 text-[13px] leading-tight text-foreground">
                                {scoreItem.name}
                                {scoreItem.unit && (
                                  <span className="ml-1 text-[11px] text-muted-foreground">
                                    ({scoreItem.unit})
                                  </span>
                                )}
                              </span>
                              <span className="shrink-0">{currentChip(cur, scoreItem)}</span>
                              <ChevronRight
                                className={cn(
                                  'h-3.5 w-3.5 shrink-0 text-muted-foreground/60 transition-transform',
                                  itemOpen && 'rotate-90'
                                )}
                              />
                            </button>

                            {itemOpen && (
                              <div className="space-y-3 px-3.5 pb-3.5">
                                {scaleDef && (scaleDef.kind === 'sum' || scaleDef.kind === 'administered' || scaleDef.kind === 'custom') ? (
                                  <ScaleWidget
                                    def={scaleDef}
                                    compact
                                    classify={(t) => detectLevel(t, levels)}
                                    levelName={(lv) => levels.find((l) => l.level === lv)?.name}
                                    initialAnswers={toNumberKeys(cur?.scaleResponses?.answers)}
                                    adminWords={
                                      scaleDef.administration?.type === 'word_recall'
                                        ? wordRecallSet ?? undefined
                                        : undefined
                                    }
                                    onResult={(r) =>
                                      handleScaleResult(scoreItem.id, r, templateItem.order)
                                    }
                                  />
                                ) : (
                                  <>
                                {scoreItem.unit && (
                                  <Input
                                    type="number"
                                    step="any"
                                    value={cur?.numericValue ?? ''}
                                    onChange={(e) =>
                                      handleNumericChange(
                                        scoreItem.id,
                                        e.target.value,
                                        templateItem.order,
                                        levels
                                      )
                                    }
                                    placeholder={`Valor em ${scoreItem.unit}`}
                                    className="h-9 text-sm"
                                  />
                                )}

                                {levels.length > 0 && (
                                  <div className="flex flex-wrap gap-2">
                                    {levels.map((level) => {
                                      const sel = cur?.selectedLevel === level.level
                                      return (
                                        <button
                                          key={level.id}
                                          type="button"
                                          data-level-button
                                          onClick={() =>
                                            handleLevelSelect(scoreItem.id, level, templateItem.order)
                                          }
                                          className={cn(
                                            'rounded-lg border-2 px-3 py-1.5 text-xs font-medium transition-all',
                                            sel
                                              ? LEVEL_COMPACT_SELECTED_CLASSES[level.level] ||
                                                  LEVEL_COMPACT_SELECTED_CLASSES[6]
                                              : 'border-border bg-background hover:border-primary/50'
                                          )}
                                        >
                                          {level.name} ({level.level})
                                        </button>
                                      )
                                    })}
                                  </div>
                                )}
                                  </>
                                )}

                                {showObs ? (
                                  <Textarea
                                    value={cur?.textValue || ''}
                                    onChange={(e) =>
                                      handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                                    }
                                    placeholder="Observações…"
                                    rows={2}
                                    className="resize-none text-sm"
                                  />
                                ) : (
                                  <button
                                    type="button"
                                    onClick={() => toggleSetMember(setObsOpen, scoreItem.id)}
                                    className="flex items-center gap-1.5 text-xs font-medium text-primary"
                                  >
                                    <Plus className="h-3.5 w-3.5" /> Adicionar observação
                                  </button>
                                )}

                                {scoreItem.patientExplanation && (
                                  <p className="text-[11px] leading-snug text-muted-foreground">
                                    {scoreItem.patientExplanation}
                                  </p>
                                )}

                                <AnamnesisItemHistory
                                  patientId={historyPatientId}
                                  scoreItemId={scoreItem.id}
                                  levels={levels}
                                  unit={scoreItem.unit}
                                  enabled={itemOpen}
                                />
                              </div>
                            )}
                          </div>
                        )
                      })}
                    </div>
                  )
                })}
              </div>
            )}
          </div>
        )
      })}
    </div>
  )

  return (
    <TooltipProvider delayDuration={300}>
      {compact ? compactView : (
      <div className="space-y-6">
        {Array.from(organized.groups.values()).map(({ group, subgroups }) => (
          <Card key={group.id} className="overflow-hidden border-2">
            <CardHeader className={cn('bg-muted/50', compact ? 'pb-3' : 'pb-4')}>
              <CardTitle className={cn(compact ? 'text-base' : 'text-lg', 'font-semibold flex items-center gap-2')}>
                {group.name}
                <Badge variant="outline" className="ml-auto">
                  {Array.from(subgroups.values()).reduce(
                    (acc, { items }) => acc + items.filter(({ scoreItem }) => values.has(scoreItem.id)).length,
                    0
                  )} / {Array.from(subgroups.values()).reduce((acc, { items }) => acc + items.length, 0)}
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent className={cn(compact ? 'pt-4 space-y-4' : 'pt-6 space-y-6')}>
              {Array.from(subgroups.values()).map(({ subgroup, items }) => {
                const hasMaxSelect = subgroup.maxSelect > 0

                // Count selected items in this subgroup
                const selectedInSubgroup = items.filter(({ scoreItem }) => {
                  const val = values.get(scoreItem.id)
                  return val?.selectedLevel !== undefined
                })
                const selectedCount = selectedInSubgroup.length

                return (
                  <div key={subgroup.id} className="space-y-3">
                    <div className="flex items-center justify-between border-b pb-2">
                      <h4 className={cn(
                        'font-semibold text-primary',
                        compact ? 'text-sm' : 'text-base'
                      )}>
                        {subgroup.name}
                      </h4>
                      {hasMaxSelect && (
                        <Badge variant={selectedCount >= subgroup.maxSelect ? 'default' : 'outline'} className="gap-1.5">
                          {selectedCount} / {subgroup.maxSelect}
                        </Badge>
                      )}
                    </div>

                    {hasMaxSelect && selectedCount >= subgroup.maxSelect && (
                      <div className="flex items-center gap-2 p-3 bg-blue-50 text-blue-900 rounded-lg text-sm">
                        <AlertCircle className="h-4 w-4 shrink-0" />
                        <p>Máximo atingido. Desmarque um item para selecionar outro.</p>
                      </div>
                    )}

                    <div className={cn('space-y-3', hasMaxSelect && 'space-y-2')}>
                      {items.map(({ templateItem, scoreItem }) => {
                        const currentValue = values.get(scoreItem.id)
                        const levels = scoreItem.levels || []
                        const scaleDef = getScaleDef(scoreItem.anamneseItemCode)
                        const hasLevelSelected = currentValue?.selectedLevel !== undefined
                        const hasNumericValue = currentValue?.numericValue !== undefined
                        const isFilled = !!(hasLevelSelected || hasNumericValue || currentValue?.textValue)
                        const isSelected = hasMaxSelect && hasLevelSelected

                        // Compact layout for maxSelect items
                        if (hasMaxSelect) {
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => {
                                if (el) itemRefs.current.set(scoreItem.id, el)
                              }}
                              className={cn(
                                'flex items-center gap-3 p-3 border rounded-lg transition-all',
                                isSelected ? 'border-primary bg-primary/5' : 'border-border'
                              )}
                            >
                              <Checkbox
                                checked={isSelected}
                                onCheckedChange={() =>
                                  handleCheckboxSelect(
                                    scoreItem.id,
                                    templateItem.order,
                                    subgroup.id,
                                    subgroup.maxSelect,
                                    items
                                  )
                                }
                                className="h-5 w-5 shrink-0"
                              />
                              <Label className={cn(
                                'font-medium shrink-0 cursor-pointer min-w-[200px]',
                                compact ? 'text-sm' : 'text-base'
                              )}>
                                {scoreItem.name}
                              </Label>
                              <Input
                                value={currentValue?.textValue || ''}
                                onChange={(e) =>
                                  handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                                }
                                placeholder="Observações..."
                                className={cn('flex-1', compact ? 'h-9 text-sm' : 'h-10')}
                              />
                              {scoreItem.patientExplanation && (
                                <Tooltip>
                                  <TooltipTrigger asChild>
                                    <button
                                      type="button"
                                      className="text-muted-foreground hover:text-foreground transition-colors shrink-0"
                                    >
                                      <HelpCircle className="h-4 w-4" />
                                    </button>
                                  </TooltipTrigger>
                                  <TooltipContent className="max-w-md">
                                    <p className="text-sm">{scoreItem.patientExplanation}</p>
                                  </TooltipContent>
                                </Tooltip>
                              )}
                            </div>
                          )
                        }

                        // Regular layout for non-maxSelect items
                        const isThisItemFocused = focusedItemId === scoreItem.id

                        return (
                          <Card
                            key={templateItem.id}
                            ref={(el) => {
                              if (el) itemRefs.current.set(scoreItem.id, el)
                            }}
                            className={cn(
                              'border-2 transition-all',
                              isThisItemFocused ? 'border-primary bg-primary/5' : 'border-border'
                            )}
                            onFocus={() => setFocusedItemId(scoreItem.id)}
                            onBlur={() => setFocusedItemId(null)}
                            tabIndex={-1}
                          >
                            <CardContent className={cn(compact ? 'p-4 space-y-3' : 'p-5 space-y-4')}>
                              {/* Item header */}
                              <div className="flex items-start justify-between gap-3">
                                <div className="flex items-center gap-3 flex-1">
                                  {isFilled && !compact && (
                                    <CheckCircle2 className="h-5 w-5 text-primary shrink-0" />
                                  )}
                                  <div className="flex-1">
                                    <Label className={cn(
                                      'font-semibold cursor-pointer',
                                      compact ? 'text-sm' : 'text-base'
                                    )}>
                                      {scoreItem.name}
                                      {scoreItem.unit && (
                                        <span className="text-muted-foreground ml-2 font-normal">
                                          ({scoreItem.unit})
                                        </span>
                                      )}
                                    </Label>
                                  </div>
                                </div>
                                {scoreItem.patientExplanation && (
                                  <Tooltip>
                                    <TooltipTrigger asChild>
                                      <button
                                        type="button"
                                        className="text-muted-foreground hover:text-foreground transition-colors p-1"
                                      >
                                        <HelpCircle className={compact ? 'h-4 w-4' : 'h-5 w-5'} />
                                      </button>
                                    </TooltipTrigger>
                                    <TooltipContent className="max-w-md">
                                      <p className="text-sm">{scoreItem.patientExplanation}</p>
                                    </TooltipContent>
                                  </Tooltip>
                                )}
                              </div>

                              {/* Escala (PHQ-9, GAD-7, …): widget pergunta-a-pergunta */}
                              {scaleDef && (scaleDef.kind === 'sum' || scaleDef.kind === 'administered' || scaleDef.kind === 'custom') ? (
                                <ScaleWidget
                                  def={scaleDef}
                                  classify={(t) => detectLevel(t, levels)}
                                  levelName={(lv) => levels.find((l) => l.level === lv)?.name}
                                  initialAnswers={toNumberKeys(currentValue?.scaleResponses?.answers)}
                                  adminWords={
                                    scaleDef.administration?.type === 'word_recall'
                                      ? wordRecallSet ?? undefined
                                      : undefined
                                  }
                                  onResult={(r) =>
                                    handleScaleResult(scoreItem.id, r, templateItem.order)
                                  }
                                />
                              ) : (
                                <>
                              {/* Numeric input for items with unit */}
                              {scoreItem.unit && (
                                <div className="space-y-1">
                                  <Label className="text-xs font-medium text-muted-foreground">
                                    Valor medido ({scoreItem.unit}):
                                  </Label>
                                  <Input
                                    type="number"
                                    step="any"
                                    value={currentValue?.numericValue ?? ''}
                                    onChange={(e) => handleNumericChange(scoreItem.id, e.target.value, templateItem.order, levels)}
                                    placeholder={`Digite em ${scoreItem.unit}`}
                                    className={cn(compact ? 'h-9 text-sm' : 'h-10')}
                                  />
                                </div>
                              )}

                              {/* Levels */}
                              {levels.length > 0 && (
                                <div className="space-y-2">
                                  <Label className="text-xs font-medium text-muted-foreground">
                                    {scoreItem.unit
                                      ? (compact ? 'Nível (auto ou manual):' : 'Nível (auto-detectado ou selecione manualmente):')
                                      : (compact ? 'Níveis (opcional):' : 'Selecione o nível:')}
                                  </Label>
                                  <div className={cn(
                                    compact ? 'flex flex-wrap gap-2' : 'grid grid-cols-2 sm:grid-cols-3 gap-3'
                                  )}>
                                    {levels.map((level) => {
                                      const isLevelSelected = currentValue?.selectedLevel === level.level
                                      const selectedClass = LEVEL_SELECTED_CLASSES[level.level] || LEVEL_SELECTED_CLASSES[6]
                                      const hoverClass = LEVEL_HOVER_CLASSES[level.level] || LEVEL_HOVER_CLASSES[6]
                                      const compactSelectedClass = LEVEL_COMPACT_SELECTED_CLASSES[level.level] || LEVEL_COMPACT_SELECTED_CLASSES[6]

                                      return (
                                        <button
                                          key={level.id}
                                          type="button"
                                          data-level-button
                                          onClick={() =>
                                            handleLevelSelect(scoreItem.id, level, templateItem.order)
                                          }
                                          className={cn(
                                            'rounded-lg border-2 text-sm font-medium transition-all',
                                            compact
                                              ? 'px-3 py-1.5'
                                              : 'p-4 border-3 font-bold flex flex-col items-center gap-1 min-h-[80px] justify-center',
                                            isLevelSelected
                                              ? compact
                                                ? compactSelectedClass
                                                : selectedClass
                                              : compact
                                                ? 'border-border hover:border-primary/50 bg-background'
                                                : hoverClass
                                          )}
                                        >
                                          {compact ? (
                                            `${level.name} (${level.level})`
                                          ) : (
                                            <>
                                              <span className="text-lg">N{level.level}</span>
                                              <span className="text-xs text-center leading-tight">
                                                {level.name}
                                              </span>
                                            </>
                                          )}
                                        </button>
                                      )
                                    })}
                                  </div>
                                </div>
                              )}
                                </>
                              )}

                              {/* Text input */}
                              <div className="space-y-2">
                                <Label className="text-xs font-medium text-muted-foreground">
                                  Observações{compact ? ' (opcional)' : ' adicionais'}:
                                </Label>
                                <Textarea
                                  value={currentValue?.textValue || ''}
                                  onChange={(e) =>
                                    handleTextChange(scoreItem.id, e.target.value, templateItem.order)
                                  }
                                  placeholder={compact ? 'Digite observações adicionais...' : 'Digite observações sobre este item...'}
                                  rows={compact ? 2 : 3}
                                  className={cn('resize-none', compact ? 'text-sm' : 'text-base')}
                                />
                              </div>

                              {/* Status badges */}
                              {isFilled && (
                                <div className="flex gap-2 flex-wrap pt-2 border-t">
                                  {currentValue.selectedLevel !== undefined && (
                                    <Badge variant={compact ? 'secondary' : 'default'} className="text-xs">
                                      {compact ? `Nível: ${currentValue.selectedLevel}` : `Nível selecionado: ${currentValue.selectedLevel}`}
                                    </Badge>
                                  )}
                                  {currentValue.numericValue !== undefined && (
                                    <Badge variant="secondary" className="text-xs">
                                      {compact
                                        ? `${currentValue.numericValue} ${scoreItem.unit ?? ''}`
                                        : `Valor: ${currentValue.numericValue} ${scoreItem.unit ?? ''}`}
                                    </Badge>
                                  )}
                                  {currentValue.textValue && (
                                    <Badge variant="secondary" className="text-xs">
                                      Com observações
                                    </Badge>
                                  )}
                                </div>
                              )}
                            </CardContent>
                          </Card>
                        )
                      })}
                    </div>
                  </div>
                )
              })}
            </CardContent>
          </Card>
        ))}
      </div>
      )}
    </TooltipProvider>
  )
}
