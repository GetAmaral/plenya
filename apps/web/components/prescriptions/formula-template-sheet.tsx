'use client'

import { useEffect, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { ChevronLeft, FlaskConical, Loader2, Search, Sparkles } from 'lucide-react'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Sheet, SheetContent, SheetTitle } from '@/components/ui/sheet'
import { Skeleton } from '@/components/ui/skeleton'
import { toBullets } from '@/lib/api/magistral-components'
import {
  listFormulaTemplates,
  suggestDoses,
  type DoseBandView,
  type DoseSuggestion,
  type FormulaTemplate,
} from '@/lib/api/magistral-templates'

interface FormulaTemplateSheetProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  patientId?: string
  /** Aplica a fórmula-base no formulário, com as doses que ficaram na tela. */
  onApply: (
    template: FormulaTemplate,
    doses: Record<string, number>,
    suggested: Record<string, number>
  ) => void
}

const fmt = (v?: number | null) => (v === null || v === undefined ? '' : String(v).replace('.', ','))

/** Como a faixa se lê: "até 20", "20 a 30", "acima de 60". */
const faixaTexto = (b: DoseBandView) => {
  if (b.lowerBound === null || b.lowerBound === undefined) return `até ${fmt(b.upperBound)}`
  if (b.upperBound === null || b.upperBound === undefined) return `acima de ${fmt(b.lowerBound)}`
  return `${fmt(b.lowerBound)} a ${fmt(b.upperBound)}`
}

/**
 * Fórmulas-base: escolher a fórmula pronta e, antes de aplicar, ver o que as regras sugerem para
 * ESTE paciente.
 *
 * A sugestão nunca entra sozinha. Cada componente mostra a dose da base, o que a regra sugeriu e
 * a frase de onde isso saiu ("peso 82,5 kg da consulta × 4 mg/kg"). Aceitar é um clique; ignorar
 * é não fazer nada.
 */
export function FormulaTemplateSheet({
  open,
  onOpenChange,
  patientId,
  onApply,
}: FormulaTemplateSheetProps) {
  const [term, setTerm] = useState('')
  const [debounced, setDebounced] = useState('')
  const [selected, setSelected] = useState<FormulaTemplate | null>(null)
  const [doses, setDoses] = useState<Record<string, number>>({})
  const [accepted, setAccepted] = useState<Record<string, number>>({})

  useEffect(() => {
    if (open) {
      setSelected(null)
      setTerm('')
      setDebounced('')
    }
  }, [open])

  useEffect(() => {
    const t = setTimeout(() => setDebounced(term), 250)
    return () => clearTimeout(t)
  }, [term])

  const { data: templates = [], isFetching } = useQuery({
    queryKey: ['formula-templates', debounced],
    queryFn: () => listFormulaTemplates(debounced),
    enabled: open,
  })

  const { data: suggestion, isFetching: loadingSuggestion } = useQuery<DoseSuggestion>({
    queryKey: ['dose-suggestion', selected?.id, patientId],
    queryFn: () => suggestDoses(selected!.id, patientId!),
    enabled: !!selected && !!patientId,
  })

  const choose = (t: FormulaTemplate) => {
    setSelected(t)
    setAccepted({})
    const base: Record<string, number> = {}
    for (const c of t.components ?? []) base[c.substance] = c.quantity
    setDoses(base)
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent side="right" className="flex w-full flex-col gap-0 p-0 sm:max-w-xl">
        <SheetTitle className="sr-only">Fórmulas-base</SheetTitle>

        {!selected ? (
          <>
            <div className="border-b p-5 pr-12">
              <h2 className="flex items-center gap-2 text-lg font-semibold">
                <FlaskConical className="h-5 w-5 text-muted-foreground" />
                Fórmulas-base
              </h2>
              <p className="mt-1 text-sm text-muted-foreground">
                Suas fórmulas prontas. Busque pelo nome ou pelo que ela trata.
              </p>
              <div className="relative mt-3">
                <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                <Input
                  autoFocus
                  value={term}
                  onChange={(e) => setTerm(e.target.value)}
                  placeholder="sono, resistência insulínica, cabelo..."
                  className="pl-9"
                />
                {isFetching && (
                  <Loader2 className="absolute right-3 top-1/2 h-4 w-4 -translate-y-1/2 animate-spin text-muted-foreground" />
                )}
              </div>
            </div>

            <div className="flex-1 overflow-y-auto p-3">
              {isFetching && templates.length === 0 && <Skeleton className="h-24 w-full" />}
              {!isFetching && templates.length === 0 && (
                <p className="p-6 text-center text-sm text-muted-foreground">
                  Nenhuma fórmula-base ainda. Monte uma fórmula e use “Salvar como fórmula-base”.
                </p>
              )}
              <ul className="space-y-2">
                {templates.map((t) => (
                  <li key={t.id}>
                    <button
                      type="button"
                      onClick={() => choose(t)}
                      className="w-full rounded-lg border p-3 text-left transition hover:border-primary/40 hover:bg-accent"
                    >
                      <div className="flex items-baseline justify-between gap-3">
                        <span className="font-medium">{t.name}</span>
                        <span className="shrink-0 text-xs text-muted-foreground">
                          {t.pharmaceuticalForm} · {t.components?.length ?? 0} componentes
                        </span>
                      </div>
                      {toBullets(t.indicationBullets).length > 0 ? (
                        <ul className="mt-1 space-y-0.5">
                          {toBullets(t.indicationBullets)
                            .slice(0, 2)
                            .map((b) => (
                              <li key={b} className="flex gap-1.5 text-xs text-muted-foreground">
                                <span className="text-primary/60">·</span>
                                <span className="truncate">{b}</span>
                              </li>
                            ))}
                        </ul>
                      ) : (
                        t.indication && (
                          <p className="mt-1 line-clamp-2 text-xs text-muted-foreground">
                            {t.indication}
                          </p>
                        )
                      )}
                      <div className="mt-2 flex flex-wrap gap-1">
                        {t.usageCount > 0 && (
                          <Badge variant="outline" className="text-[10px] text-muted-foreground">
                            {t.usageCount} receita(s)
                          </Badge>
                        )}
                        {/* Fórmula que ninguém conferiu tem que dizer isso na lista, não na
                            tela de dentro: é aqui que ela é escolhida. */}
                        {!t.lastReview && (
                          <Badge
                            variant="outline"
                            className="border-sky-200 bg-sky-50 text-[10px] text-sky-700"
                          >
                            a conferir
                          </Badge>
                        )}
                      </div>
                    </button>
                  </li>
                ))}
              </ul>
            </div>
          </>
        ) : (
          <>
            <div className="border-b p-5 pr-12">
              <Button
                variant="ghost"
                size="sm"
                className="-ml-2 mb-2"
                onClick={() => setSelected(null)}
              >
                <ChevronLeft className="mr-1 h-4 w-4" />
                Fórmulas-base
              </Button>
              <h2 className="text-xl font-semibold">{selected.name}</h2>
              <p className="text-sm text-muted-foreground">
                {selected.pharmaceuticalForm} ·{' '}
                {selected.usageType === 'external' ? 'uso externo' : 'uso interno'} · aviar{' '}
                {fmt(selected.quantityToDispense)} {selected.quantityUnit}
              </p>
            </div>

            <div className="flex-1 space-y-4 overflow-y-auto p-5">
              {selected.indication && (
                <section>
                  <h3 className="flex items-center gap-2 text-sm font-semibold">
                    <Sparkles className="h-4 w-4 text-muted-foreground" />
                    Para que serve
                  </h3>
                  <p className="mt-1 text-sm leading-relaxed text-muted-foreground">
                    {selected.indication}
                  </p>
                </section>
              )}

              <section>
                <div className="flex items-center justify-between">
                  <h3 className="text-sm font-semibold">Componentes e doses</h3>
                  {loadingSuggestion && (
                    <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
                  )}
                </div>

                {suggestion?.weightKg !== undefined && (
                  <p className="mt-1 text-xs text-muted-foreground">
                    Peso usado nas regras: {fmt(suggestion.weightKg)} kg · {suggestion.weightSource}
                    {suggestion.weightDate
                      ? ` · ${suggestion.weightDate.split('-').reverse().join('/')}`
                      : ''}
                  </p>
                )}

                <ul className="mt-3 space-y-3">
                  {(selected.components ?? []).map((c) => {
                    const sug = suggestion?.items.find((i) => i.substance === c.substance)
                    const atual = doses[c.substance] ?? c.quantity
                    return (
                      <li key={c.substance} className="rounded-lg border p-3">
                        <div className="flex items-baseline justify-between gap-3">
                          <span className="font-medium">{c.substance}</span>
                          <span className="flex items-center gap-2">
                            <Input
                              type="number"
                              step="any"
                              className="h-8 w-24 text-right"
                              value={atual}
                              onChange={(e) =>
                                setDoses((prev) => ({
                                  ...prev,
                                  [c.substance]: Number(e.target.value),
                                }))
                              }
                            />
                            <span className="text-sm text-muted-foreground">{c.unit}</span>
                          </span>
                        </div>

                        {sug?.suggested !== undefined && sug?.suggested !== null && (
                          <div className="mt-2 rounded-md border border-dashed p-2">
                            <div className="flex flex-wrap items-center justify-between gap-2">
                              <span className="text-xs">
                                Regra sugere{' '}
                                <strong>
                                  {fmt(sug.suggested)} {c.unit}
                                </strong>
                                {sug.clamped && (
                                  <Badge
                                    variant="outline"
                                    className="ml-2 bg-amber-50 text-[10px] text-amber-700"
                                  >
                                    limitada pela trava
                                  </Badge>
                                )}
                              </span>
                              <Button
                                type="button"
                                size="sm"
                                variant={accepted[c.substance] ? 'default' : 'outline'}
                                onClick={() => {
                                  setDoses((prev) => ({ ...prev, [c.substance]: sug.suggested! }))
                                  setAccepted((prev) => ({ ...prev, [c.substance]: sug.suggested! }))
                                }}
                              >
                                {accepted[c.substance] ? 'aceita' : 'usar sugestão'}
                              </Button>
                            </div>
                            <p className="mt-1 text-[11px] leading-snug text-muted-foreground">
                              {sug.basis}
                            </p>

                            {/* A escada inteira, não só o degrau que pegou: vendo a régua o
                                médico julga a conduta, e não apenas o número que ela cuspiu. */}
                            {!!sug.bands?.length && (
                              <ul className="mt-2 grid gap-px overflow-hidden rounded border bg-border text-[11px]">
                                {sug.bands.map((b, bi) => (
                                  <li
                                    key={bi}
                                    className={`flex items-center justify-between gap-2 px-2 py-1 ${
                                      b.active
                                        ? 'bg-primary/10 font-medium text-foreground'
                                        : 'bg-background text-muted-foreground'
                                    }`}
                                  >
                                    <span>
                                      {faixaTexto(b)}
                                      {b.label ? ` · ${b.label}` : ''}
                                    </span>
                                    <span className="tabular-nums">
                                      {fmt(b.dose)} {c.unit}
                                    </span>
                                  </li>
                                ))}
                              </ul>
                            )}
                          </div>
                        )}

                        {sug?.suggested === undefined && sug?.reason && (
                          <p className="mt-2 text-[11px] text-muted-foreground">{sug.reason}</p>
                        )}

                        {sug?.lastPrescribed !== undefined && (
                          <p className="mt-1 text-[11px] text-muted-foreground">
                            Última assinada para este paciente: {fmt(sug.lastPrescribed)} {c.unit}
                          </p>
                        )}
                      </li>
                    )
                  })}
                </ul>
              </section>

              <p className="text-xs text-muted-foreground">
                A sugestão não entra na receita sozinha: ela preenche o campo acima, e o que vale é
                o que estiver aí quando você aplicar.
              </p>
            </div>

            <div className="flex items-center gap-3 border-t p-4">
              <div className="flex-1 text-xs text-muted-foreground">
                <Label className="text-xs">Posologia da base</Label>
                <p>{selected.posology || 'sem posologia definida'}</p>
              </div>
              <Button
                onClick={() => {
                  onApply(selected, doses, accepted)
                  onOpenChange(false)
                }}
              >
                Usar esta fórmula
              </Button>
            </div>
          </>
        )}
      </SheetContent>
    </Sheet>
  )
}
