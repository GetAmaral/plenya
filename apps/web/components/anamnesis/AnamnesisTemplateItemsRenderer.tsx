'use client'

import { useState, useEffect, useRef, type Dispatch, type SetStateAction } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Badge } from '@/components/ui/badge'
import { Checkbox } from '@/components/ui/checkbox'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { HelpCircle, CheckCircle2, AlertCircle, ChevronDown, ChevronRight } from 'lucide-react'
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

// Chips inline densos (visão compacta): cinza neutro quando NÃO selecionado; ao selecionar,
// assume a cor do nível como no impresso (N0 vermelho → N5 esmeralda).
const LEVEL_CHIP_IDLE = 'border-gray-200 bg-gray-50 text-gray-600 hover:bg-gray-100'
const LEVEL_CHIP_SELECTED: Record<number, string> = {
  0: 'border-red-500 bg-red-100 text-red-900 ring-1 ring-red-400',
  1: 'border-orange-500 bg-orange-100 text-orange-900 ring-1 ring-orange-400',
  2: 'border-yellow-500 bg-yellow-100 text-yellow-900 ring-1 ring-yellow-400',
  3: 'border-blue-500 bg-blue-100 text-blue-900 ring-1 ring-blue-400',
  4: 'border-green-500 bg-green-100 text-green-900 ring-1 ring-green-400',
  5: 'border-emerald-500 bg-emerald-100 text-emerald-900 ring-1 ring-emerald-400',
  6: 'border-gray-500 bg-gray-100 text-gray-900 ring-1 ring-gray-400',
}

// Organize items by group and subgroup
interface OrganizedItem {
  templateItem: AnamnesisTemplateItem
  scoreItem: ScoreItem
  depth: number
}
interface OrganizedData {
  groups: Map<string, {
    group: ScoreGroup
    subgroups: Map<string, {
      subgroup: ScoreSubgroup
      items: OrganizedItem[]
    }>
  }>
}

// Aninha os itens de um subgrupo por parentItemId (mesma lógica do pôster/print do score):
// o pai é mantido e seus filhos vêm logo abaixo, com profundidade crescente para indentação.
// Se o pai não estiver presente no subgrupo (ex.: filtrado por demografia), o filho vira raiz.
function nestByParent(items: OrganizedItem[]): OrganizedItem[] {
  const byId = new Map<string, OrganizedItem>()
  items.forEach((it) => byId.set(it.scoreItem.id, it))
  const childrenOf = new Map<string, OrganizedItem[]>()
  const roots: OrganizedItem[] = []
  for (const it of items) {
    const pid = it.scoreItem.parentItemId
    if (pid && byId.has(pid)) {
      const list = childrenOf.get(pid) ?? []
      list.push(it)
      childrenOf.set(pid, list)
    } else {
      roots.push(it)
    }
  }
  const byOrder = (a: OrganizedItem, b: OrganizedItem) => a.templateItem.order - b.templateItem.order
  const out: OrganizedItem[] = []
  const visit = (it: OrganizedItem, depth: number) => {
    out.push({ ...it, depth })
    ;(childrenOf.get(it.scoreItem.id) ?? []).slice().sort(byOrder).forEach((k) => visit(k, depth + 1))
  }
  roots.sort(byOrder).forEach((r) => visit(r, 0))
  return out
}

function organizeTemplateItems(template: AnamnesisTemplate, patient?: Patient | null): OrganizedData {
  const groups = new Map()

  if (!template.items) {
    return { groups }
  }

  const sortedItems = [...template.items].sort((a, b) => a.order - b.order)

  sortedItems.forEach((templateItem) => {
    const scoreItem = templateItem.scoreItem

    if (!scoreItem || !scoreItem.subgroup) {
      return
    }

    // Filtro demográfico (idade/gênero/menopausa) — aplica a pais e filhos.
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
    subgroupData.items.push({ templateItem, scoreItem, depth: 0 })
  })

  // Reordena cada subgrupo como árvore pai→filho, atribuindo profundidade p/ indentação.
  for (const groupData of groups.values()) {
    for (const subgroupData of groupData.subgroups.values()) {
      subgroupData.items = nestByParent(subgroupData.items)
    }
  }

  return { groups }
}

// "?" de explicação do item. O texto (patientExplanation) já vem com o template — nada de fetch.
// Popover abre instantaneamente no clique (sem o delay de hover do tooltip).
function HelpPopover({ text, size = 'sm' }: { text: string; size?: 'sm' | 'md' }) {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <button
          type="button"
          aria-label="Explicação"
          className="shrink-0 text-muted-foreground/60 transition-colors hover:text-foreground"
        >
          <HelpCircle className={size === 'md' ? 'h-4 w-4' : 'h-3.5 w-3.5'} />
        </button>
      </PopoverTrigger>
      <PopoverContent align="start" className="max-w-sm text-sm leading-snug">
        {text}
      </PopoverContent>
    </Popover>
  )
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

  // --- Estado da visão compacta: accordion de grupos (todos abertos por padrão) ---
  const [expandedGroups, setExpandedGroups] = useState<Set<string>>(
    () => new Set(organized.groups.keys()) // todos os grupos abertos por padrão
  )
  const [expandedItems, setExpandedItems] = useState<Set<string>>(new Set()) // só escalas colapsam
  const [obsOpen, setObsOpen] = useState<Set<string>>(new Set())
  const [histOpen, setHistOpen] = useState<Set<string>>(new Set())

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

  // Pré-seleção do Nível 5 por padrão: itens marcados como defaultLevel5 (histórico de doença /
  // uso de medicação) já chegam preenchidos como "sem doença"/"sem uso". Só preenche itens
  // aplicáveis ao paciente e ainda SEM valor (não sobrescreve resposta existente). Roda uma vez.
  const didInitDefaults = useRef(false)
  useEffect(() => {
    if (didInitDefaults.current) return
    didInitDefaults.current = true
    setValues((prev) => {
      const next = new Map(prev)
      let changed = false
      for (const ti of template.items ?? []) {
        const si = ti.scoreItem as ScoreItem | undefined
        if (!si || !si.defaultLevel5) continue
        if (next.has(si.id)) continue
        if (!itemAppliesToPatient(si, patient)) continue
        const levels = si.levels ?? []
        if (levels.length === 0) continue
        const top = levels.some((l) => l.level === 5) ? 5 : Math.max(...levels.map((l) => l.level))
        next.set(si.id, { scoreItemId: si.id, selectedLevel: top, order: ti.order ?? 0 })
        changed = true
      }
      if (changed) requestAnimationFrame(() => onChange(Array.from(next.values())))
      return changed ? next : prev
    })
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])


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
          // valor final da escala = medida crua (numeric_value); selected_level é a classificação.
          numericValue: total,
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

  // --- Visão compacta (densa): grupos abertos por padrão, níveis inline coloridos como o impresso ---
  const compactView = (
    <div className="space-y-2">
      {Array.from(organized.groups.values()).map(({ group, subgroups }) => {
        const subArr = Array.from(subgroups.values())
        const allItems = subArr.flatMap((s) => s.items)
        const filled = allItems.filter(({ scoreItem }) => values.has(scoreItem.id)).length
        const total = allItems.length
        const gOpen = expandedGroups.has(group.id)

        return (
          <div key={group.id} className="overflow-hidden rounded-lg border bg-card">
            <button
              type="button"
              onClick={() => toggleSetMember(setExpandedGroups, group.id)}
              className="flex w-full items-center gap-2 bg-muted/40 px-3 py-1.5 text-left"
            >
              <span className="text-[13px] font-semibold text-foreground">{group.name}</span>
              <span className="ml-auto shrink-0 rounded-full bg-background px-1.5 py-0.5 text-[10px] text-muted-foreground">
                {filled}/{total}
              </span>
              <ChevronDown
                className={cn(
                  'h-3.5 w-3.5 shrink-0 text-muted-foreground transition-transform',
                  gOpen && 'rotate-180'
                )}
              />
            </button>

            {gOpen && (
              <div className="divide-y divide-border/70">
                {subArr.map(({ subgroup, items }) => {
                  const hasMaxSelect = subgroup.maxSelect > 0
                  const selectedCount = items.filter(
                    ({ scoreItem }) => values.get(scoreItem.id)?.selectedLevel !== undefined
                  ).length

                  return (
                    <div key={subgroup.id}>
                      {subArr.length > 1 && (
                        <div className="flex items-center justify-between border-b border-border/70 bg-muted/20 px-3 py-1">
                          <span className="text-[10px] font-semibold uppercase tracking-wide text-muted-foreground">
                            {subgroup.name}
                          </span>
                          {hasMaxSelect && (
                            <span className="text-[10px] text-muted-foreground">
                              {selectedCount}/{subgroup.maxSelect}
                            </span>
                          )}
                        </div>
                      )}

                      {items.map(({ templateItem, scoreItem, depth }) => {
                        const cur = values.get(scoreItem.id)
                        const levels = scoreItem.levels || []
                        const scaleDef = getScaleDef(scoreItem.anamneseItemCode)
                        const isScale = !!scaleDef && (scaleDef.kind === 'sum' || scaleDef.kind === 'administered' || scaleDef.kind === 'custom')
                        const indentStyle = depth > 0 ? { marginLeft: depth * 14 } : undefined
                        const rowCls = cn(
                          'px-3 py-1.5 last:border-b-0',
                          depth > 0 && 'border-l-2 border-l-primary/25'
                        )

                        // maxSelect → linha com checkbox
                        if (hasMaxSelect) {
                          const isSel = cur?.selectedLevel !== undefined
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => { if (el) itemRefs.current.set(scoreItem.id, el) }}
                              style={indentStyle}
                              className={cn('flex items-center gap-2', rowCls, isSel && 'bg-primary/5')}
                            >
                              <Checkbox
                                checked={isSel}
                                onCheckedChange={() =>
                                  handleCheckboxSelect(scoreItem.id, templateItem.order, subgroup.id, subgroup.maxSelect, items)
                                }
                                className="h-4 w-4 shrink-0"
                              />
                              <span className="text-[12.5px] leading-tight text-foreground">{scoreItem.name}</span>
                            </div>
                          )
                        }

                        // Escala (PHQ-9, GAD-7, Dubois…) → linha colapsável (widget é grande)
                        if (isScale) {
                          const open = expandedItems.has(scoreItem.id)
                          const showHistScale = histOpen.has(scoreItem.id)
                          const scaleTotalVal = cur?.scaleResponses?.total ?? cur?.numericValue
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => { if (el) itemRefs.current.set(scoreItem.id, el) }}
                              style={indentStyle}
                              className={cn(rowCls, 'px-0 py-0', depth > 0 && 'border-l-2 border-l-primary/25')}
                            >
                              <div className="flex w-full flex-wrap items-center gap-2 px-3 py-1.5">
                                <button
                                  type="button"
                                  onClick={() => toggleSetMember(setExpandedItems, scoreItem.id)}
                                  className="flex min-w-0 items-center gap-2 text-left"
                                  title="Abrir/fechar perguntas"
                                >
                                  <span className="min-w-0 text-[12.5px] leading-tight text-foreground">
                                    {depth > 0 && <span className="mr-1 text-muted-foreground/50">└</span>}
                                    {scoreItem.name}
                                    <span className="ml-1 rounded bg-muted px-1 text-[9px] text-muted-foreground">escala</span>
                                  </span>
                                  <ChevronRight className={cn('h-3.5 w-3.5 shrink-0 text-muted-foreground/60 transition-transform', open && 'rotate-90')} />
                                </button>
                                {scaleTotalVal !== undefined && (
                                  <span className="shrink-0 rounded-full border border-primary/30 bg-primary/5 px-2 py-0.5 text-[11px] font-bold tabular-nums text-foreground">
                                    {scaleTotalVal}/{scaleDef!.maxScore}
                                  </span>
                                )}
                                {/* Todos os níveis, como qualquer item — o nível certo é auto-marcado pelo cálculo da escala (não clicável). */}
                                <div className="ml-auto flex flex-wrap items-center justify-end gap-1">
                                  {levels.map((level) => {
                                    const sel = cur?.selectedLevel === level.level
                                    return (
                                      <span
                                        key={level.id}
                                        title={`N${level.level} · ${level.name}`}
                                        className={cn(
                                          'rounded-md border px-1.5 py-0.5 text-[11px] leading-tight',
                                          sel ? LEVEL_CHIP_SELECTED[level.level] ?? LEVEL_CHIP_SELECTED[6] : LEVEL_CHIP_IDLE
                                        )}
                                      >
                                        <span className="font-bold">N{level.level}</span>
                                        {level.name && <span className="ml-1">{level.name}</span>}
                                      </span>
                                    )
                                  })}
                                  <button
                                    type="button"
                                    onClick={() => toggleSetMember(setHistOpen, scoreItem.id)}
                                    title="Histórico do item"
                                    className={cn('shrink-0 rounded-md border border-dashed px-1.5 py-0.5 text-[11px] text-muted-foreground hover:bg-muted', showHistScale && 'border-solid text-foreground')}
                                  >
                                    hist
                                  </button>
                                </div>
                              </div>
                              {open && (
                                <div className="px-3 pb-3">
                                  <ScaleWidget
                                    def={scaleDef!}
                                    compact
                                    classify={(t) => detectLevel(t, levels)}
                                    levelName={(lv) => levels.find((l) => l.level === lv)?.name}
                                    initialAnswers={toNumberKeys(cur?.scaleResponses?.answers)}
                                    adminWords={scaleDef!.administration?.type === 'word_recall' ? wordRecallSet ?? undefined : undefined}
                                    onResult={(r) => handleScaleResult(scoreItem.id, r, templateItem.order)}
                                  />
                                </div>
                              )}
                              {showHistScale && (
                                <div className="px-3 pb-3">
                                  <AnamnesisItemHistory
                                    patientId={historyPatientId}
                                    scoreItemId={scoreItem.id}
                                    levels={levels}
                                    unit={scoreItem.unit}
                                    anamneseItemCode={scoreItem.anamneseItemCode}
                                    enabled={showHistScale}
                                  />
                                </div>
                              )}
                            </div>
                          )
                        }

                        // Item padrão → tudo inline (sem abrir): nome + chips coloridos + obs/histórico
                        const showObs = obsOpen.has(scoreItem.id) || !!cur?.textValue
                        const showHist = histOpen.has(scoreItem.id)

                        return (
                          <div
                            key={templateItem.id}
                            ref={(el) => { if (el) itemRefs.current.set(scoreItem.id, el) }}
                            style={indentStyle}
                            className={cn(rowCls, cur && 'bg-emerald-50/30')}
                          >
                            <div className="flex flex-wrap items-center gap-x-2 gap-y-1">
                              <span className="text-[12.5px] font-medium leading-tight text-foreground">
                                {depth > 0 && <span className="mr-1 text-muted-foreground/50">└</span>}
                                {scoreItem.name}
                                {scoreItem.unit && <span className="ml-1 text-[10px] font-normal text-muted-foreground">({scoreItem.unit})</span>}
                              </span>

                              {scoreItem.patientExplanation && (
                                <HelpPopover text={scoreItem.patientExplanation} />
                              )}

                              {scoreItem.unit && (
                                <Input
                                  type="number"
                                  step="any"
                                  value={cur?.numericValue ?? ''}
                                  onChange={(e) => handleNumericChange(scoreItem.id, e.target.value, templateItem.order, levels)}
                                  placeholder={scoreItem.unit}
                                  className="h-7 w-24 text-xs"
                                />
                              )}

                              <div className="ml-auto flex flex-wrap items-center justify-end gap-1">
                                {levels.map((level) => {
                                  const sel = cur?.selectedLevel === level.level
                                  return (
                                    <button
                                      key={level.id}
                                      type="button"
                                      data-level-button
                                      onClick={() => handleLevelSelect(scoreItem.id, level, templateItem.order)}
                                      title={`N${level.level} · ${level.name}`}
                                      className={cn(
                                        'rounded-md border px-1.5 py-0.5 text-[11px] leading-tight transition-all',
                                        sel ? LEVEL_CHIP_SELECTED[level.level] ?? LEVEL_CHIP_SELECTED[6] : LEVEL_CHIP_IDLE
                                      )}
                                    >
                                      <span className="font-bold">N{level.level}</span>
                                      {level.name && <span className="ml-1">{level.name}</span>}
                                    </button>
                                  )
                                })}

                                <button
                                  type="button"
                                  onClick={() => toggleSetMember(setObsOpen, scoreItem.id)}
                                  title="Observação"
                                  className={cn('rounded-md border border-dashed px-1.5 py-0.5 text-[11px] text-muted-foreground hover:bg-muted', showObs && 'border-solid text-foreground')}
                                >
                                  obs
                                </button>
                                <button
                                  type="button"
                                  onClick={() => toggleSetMember(setHistOpen, scoreItem.id)}
                                  title="Histórico do item"
                                  className={cn('rounded-md border border-dashed px-1.5 py-0.5 text-[11px] text-muted-foreground hover:bg-muted', showHist && 'border-solid text-foreground')}
                                >
                                  hist
                                </button>
                              </div>
                            </div>

                            {showObs && (
                              <Textarea
                                value={cur?.textValue || ''}
                                onChange={(e) => handleTextChange(scoreItem.id, e.target.value, templateItem.order)}
                                placeholder="Observações…"
                                rows={2}
                                className="mt-1.5 resize-none text-sm"
                              />
                            )}

                            {showHist && (
                              <div className="mt-1.5">
                                <AnamnesisItemHistory
                                  patientId={historyPatientId}
                                  scoreItemId={scoreItem.id}
                                  levels={levels}
                                  unit={scoreItem.unit}
                                  anamneseItemCode={scoreItem.anamneseItemCode}
                                  enabled={showHist}
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
    <>
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
                      <div className="flex items-center gap-2 p-3 bg-ocean-50 text-ocean-800 rounded-lg text-sm">
                        <AlertCircle className="h-4 w-4 shrink-0" />
                        <p>Máximo atingido. Desmarque um item para selecionar outro.</p>
                      </div>
                    )}

                    <div className={cn('space-y-3', hasMaxSelect && 'space-y-2')}>
                      {items.map(({ templateItem, scoreItem, depth }) => {
                        const currentValue = values.get(scoreItem.id)
                        const levels = scoreItem.levels || []
                        const scaleDef = getScaleDef(scoreItem.anamneseItemCode)
                        const hasLevelSelected = currentValue?.selectedLevel !== undefined
                        const hasNumericValue = currentValue?.numericValue !== undefined
                        const isFilled = !!(hasLevelSelected || hasNumericValue || currentValue?.textValue)
                        const isSelected = hasMaxSelect && hasLevelSelected
                        const indentStyle = depth > 0 ? { marginLeft: depth * 24 } : undefined

                        // Compact layout for maxSelect items
                        if (hasMaxSelect) {
                          return (
                            <div
                              key={templateItem.id}
                              ref={(el) => {
                                if (el) itemRefs.current.set(scoreItem.id, el)
                              }}
                              style={indentStyle}
                              className={cn(
                                'flex items-center gap-3 p-3 border rounded-lg transition-all',
                                depth > 0 && 'border-l-4 border-l-primary/30',
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
                                <HelpPopover text={scoreItem.patientExplanation} size="md" />
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
                            style={indentStyle}
                            className={cn(
                              'border-2 transition-all',
                              depth > 0 && 'border-l-4 border-l-primary/30',
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
                                      {depth > 0 && <span className="mr-1 text-muted-foreground/50">└</span>}
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
                                  <HelpPopover text={scoreItem.patientExplanation} size="md" />
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
    </>
  )
}
