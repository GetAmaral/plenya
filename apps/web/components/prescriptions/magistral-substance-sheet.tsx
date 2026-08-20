'use client'

import { useEffect, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import {
  AlertTriangle,
  BookOpen,
  ChevronLeft,
  Droplets,
  FlaskConical,
  Loader2,
  Search,
  Sparkles,
} from 'lucide-react'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Sheet, SheetContent, SheetTitle } from '@/components/ui/sheet'
import { Skeleton } from '@/components/ui/skeleton'
import {
  getMagistralEvidence,
  searchMagistralComponents,
  toBullets,
  type MagistralComponent,
} from '@/lib/api/magistral-components'

interface MagistralSubstanceSheetProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  /** Termo já digitado no campo, para o painel abrir buscando o que a pessoa começou a escrever. */
  initialQuery?: string
  /** Insere na fórmula: devolve a substância e a dose escolhida no painel. */
  onPick: (component: MagistralComponent, dose: number | null) => void
}

const fmt = (v?: number | null) => (v === null || v === undefined ? '' : String(v).replace('.', ','))

/**
 * Painel de consulta ao catálogo magistral.
 *
 * Substitui o antigo dropdown de texto corrido, que era ilegível: a indicação de uma substância
 * ocupava a lista inteira e escondia as outras opções. Aqui a lista mostra o essencial em duas
 * linhas e o detalhe abre no mesmo painel, com dose, faixa, sinalizadores, indicação, posologia e
 * os trechos das aulas de onde aquilo saiu. A dose pode ser ajustada antes de entrar na fórmula.
 */
export function MagistralSubstanceSheet({
  open,
  onOpenChange,
  initialQuery = '',
  onPick,
}: MagistralSubstanceSheetProps) {
  const [term, setTerm] = useState(initialQuery)
  const [debounced, setDebounced] = useState(initialQuery)
  const [selected, setSelected] = useState<MagistralComponent | null>(null)
  const [dose, setDose] = useState('')

  useEffect(() => {
    if (open) {
      setTerm(initialQuery)
      setDebounced(initialQuery)
      setSelected(null)
    }
  }, [open, initialQuery])

  useEffect(() => {
    const t = setTimeout(() => setDebounced(term), 250)
    return () => clearTimeout(t)
  }, [term])

  const { data: results = [], isFetching } = useQuery({
    queryKey: ['magistral-sheet', debounced],
    queryFn: () => searchMagistralComponents(debounced),
    enabled: open,
  })

  const choose = (c: MagistralComponent) => {
    setSelected(c)
    setDose(c.usualDose ? String(c.usualDose) : '')
  }

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent
        side="right"
        className="flex w-full flex-col gap-0 p-0 sm:max-w-xl"
        onOpenAutoFocus={(e) => e.preventDefault()}
      >
        {/* Radix exige um título no diálogo para leitor de tela. O visual já tem o nome da
            substância no cabeçalho, então este fica só para a acessibilidade. */}
        <SheetTitle className="sr-only">Catálogo magistral</SheetTitle>
        {selected ? (
          <SubstanceDetail
            component={selected}
            dose={dose}
            onDoseChange={setDose}
            onBack={() => setSelected(null)}
            onUse={() => {
              onPick(selected, dose ? Number(dose) : null)
              onOpenChange(false)
            }}
          />
        ) : (
          <>
            <div className="border-b p-5 pr-12">
              <h2 className="flex items-center gap-2 text-lg font-semibold">
                <FlaskConical className="h-5 w-5 text-muted-foreground" />
                Catálogo magistral
              </h2>
              <p className="mt-1 text-sm text-muted-foreground">
                Busque por substância ou pelo que você quer tratar.
              </p>
              <div className="relative mt-3">
                <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                <Input
                  autoFocus
                  value={term}
                  onChange={(e) => setTerm(e.target.value)}
                  placeholder="coenzima, sono, magnésio..."
                  className="pl-9"
                />
                {isFetching && (
                  <Loader2 className="absolute right-3 top-1/2 h-4 w-4 -translate-y-1/2 animate-spin text-muted-foreground" />
                )}
              </div>
            </div>

            <div className="flex-1 overflow-y-auto p-3">
              {isFetching && results.length === 0 && (
                <div className="space-y-2">
                  <Skeleton className="h-20 w-full" />
                  <Skeleton className="h-20 w-full" />
                </div>
              )}

              {!isFetching && results.length === 0 && (
                <p className="p-6 text-center text-sm text-muted-foreground">
                  Nada no catálogo com esse termo. Você pode digitar a substância direto na fórmula.
                </p>
              )}

              <ul className="space-y-2">
                {results.map((c) => (
                  <li key={c.id}>
                    <button
                      type="button"
                      onClick={() => choose(c)}
                      className="w-full rounded-lg border p-3 text-left transition hover:border-primary/40 hover:bg-accent"
                    >
                      <div className="flex items-baseline justify-between gap-3">
                        <span className="font-medium">{c.name}</span>
                        <span className="shrink-0 text-sm text-muted-foreground">
                          {c.usualDose ? `${fmt(c.usualDose)} ${c.defaultUnit}` : c.defaultUnit}
                        </span>
                      </div>
                      {/* Na lista, os dois primeiros tópicos dizem mais que duas linhas de
                          parágrafo cortado no meio de uma frase. */}
                      {toBullets(c.indicationBullets).length > 0 ? (
                        <ul className="mt-1 space-y-0.5">
                          {toBullets(c.indicationBullets)
                            .slice(0, 2)
                            .map((b) => (
                              <li
                                key={b}
                                className="flex gap-1.5 text-xs leading-snug text-muted-foreground"
                              >
                                <span className="text-primary/60">·</span>
                                <span className="truncate">{b}</span>
                              </li>
                            ))}
                          {toBullets(c.indicationBullets).length > 2 && (
                            <li className="pl-3 text-[11px] text-muted-foreground/70">
                              +{toBullets(c.indicationBullets).length - 2}
                            </li>
                          )}
                        </ul>
                      ) : (
                        c.indications && (
                          <p className="mt-1 line-clamp-2 text-xs leading-snug text-muted-foreground">
                            {c.indications}
                          </p>
                        )
                      )}
                      <div className="mt-2 flex flex-wrap gap-1">
                        <Flags component={c} compact />
                      </div>
                    </button>
                  </li>
                ))}
              </ul>
            </div>
          </>
        )}
      </SheetContent>
    </Sheet>
  )
}

/** Sinalizadores como chips: o que muda a decisão de forma farmacêutica e de associação. */
function Flags({ component: c, compact = false }: { component: MagistralComponent; compact?: boolean }) {
  const chips: { label: string; className: string }[] = []
  if (c.eutecticFormer) chips.push({ label: 'forma eutético', className: 'bg-amber-50 text-amber-700' })
  if (c.hygroscopic) chips.push({ label: 'higroscópico', className: 'bg-sky-50 text-sky-700' })
  if (c.oxidizing) chips.push({ label: 'oxidante', className: 'bg-amber-50 text-amber-700' })
  if (c.oxidationSensitive)
    chips.push({ label: 'sensível a oxidação', className: 'bg-amber-50 text-amber-700' })
  if (c.photosensitive) chips.push({ label: 'fotossensível', className: 'bg-amber-50 text-amber-700' })
  if (typeof c.bitterness === 'number' && c.bitterness >= 2)
    chips.push({ label: 'sabor marcante', className: 'bg-amber-50 text-amber-700' })
  if (!c.bulkDensity) chips.push({ label: 'sem densidade', className: 'bg-muted text-muted-foreground' })
  else if (c.densitySource === 'classe')
    chips.push({ label: 'densidade aproximada', className: 'bg-muted text-muted-foreground' })
  if (c.evidenceStatus === 'suggested')
    chips.push({ label: 'a conferir', className: 'bg-sky-50 text-sky-700' })
  if (c.evidenceStatus === 'confirmed')
    chips.push({ label: 'conferido', className: 'bg-green-50 text-green-700' })

  const shown = compact ? chips.slice(0, 3) : chips
  return (
    <>
      {shown.map((chip) => (
        <Badge key={chip.label} variant="outline" className={`text-[10px] ${chip.className}`}>
          {chip.label}
        </Badge>
      ))}
      {compact && chips.length > shown.length && (
        <Badge variant="outline" className="text-[10px] text-muted-foreground">
          +{chips.length - shown.length}
        </Badge>
      )}
    </>
  )
}

function SubstanceDetail({
  component: c,
  dose,
  onDoseChange,
  onBack,
  onUse,
}: {
  component: MagistralComponent
  dose: string
  onDoseChange: (v: string) => void
  onBack: () => void
  onUse: () => void
}) {
  const { data: evidence = [], isLoading } = useQuery({
    queryKey: ['magistral-evidence', c.id],
    queryFn: () => getMagistralEvidence(c.id),
  })
  const [expanded, setExpanded] = useState<string | null>(null)

  return (
    <>
      <div className="border-b p-5 pr-12">
        <Button variant="ghost" size="sm" className="-ml-2 mb-2" onClick={onBack}>
          <ChevronLeft className="mr-1 h-4 w-4" />
          Catálogo
        </Button>
        <h2 className="text-xl font-semibold">{c.name}</h2>
        {c.synonyms && <p className="text-sm text-muted-foreground">{c.synonyms}</p>}
        <div className="mt-2 flex flex-wrap gap-1">
          <Flags component={c} />
        </div>
      </div>

      <div className="flex-1 space-y-5 overflow-y-auto p-5">
        {/* Dose: o número grande é o que o médico usa; a faixa dá o contexto sem precisar ler texto. */}
        <div className="rounded-lg border p-4">
          <Label className="text-xs uppercase tracking-wide text-muted-foreground">
            Dose usual
          </Label>
          <div className="mt-1 flex items-baseline gap-2">
            <span className="text-3xl font-semibold">
              {c.usualDose ? fmt(c.usualDose) : '—'}
            </span>
            <span className="text-lg text-muted-foreground">{c.defaultUnit}</span>
          </div>
          {(c.minDose || c.maxDose) && (
            <div className="mt-3">
              <DoseRange min={c.minDose} max={c.maxDose} usual={c.usualDose} unit={c.defaultUnit} />
            </div>
          )}
          {!c.usualDose && (
            <p className="mt-1 text-xs text-muted-foreground">
              Sem dose cadastrada. O que você usar aqui pode virar o padrão desta substância.
            </p>
          )}
        </div>

        <BulletSection
          icon={<Sparkles className="h-4 w-4 text-muted-foreground" />}
          title="Para que serve"
          bullets={toBullets(c.indicationBullets)}
          fullText={c.indications}
        />

        <BulletSection
          title="Posologia de referência"
          bullets={toBullets(c.doseBullets)}
          fullText={c.doseReference}
        />

        {c.notes && (
          <section>
            <h3 className="flex items-center gap-2 text-sm font-semibold">
              <AlertTriangle className="h-4 w-4 text-muted-foreground" />
              Observações
            </h3>
            <p className="mt-1 text-sm leading-relaxed text-muted-foreground">{c.notes}</p>
          </section>
        )}

        <section className="rounded-lg border border-dashed p-3">
          <p className="flex items-start gap-2 text-xs text-muted-foreground">
            <Droplets className="mt-0.5 h-3.5 w-3.5 shrink-0" />
            {!c.bulkDensity ? (
              <span>
                Densidade aparente não cadastrada: enquanto ela faltar, a calculadora de cápsula
                não opina sobre esta fórmula.
              </span>
            ) : c.densitySource === 'classe' ? (
              <span>
                Densidade {String(c.bulkDensity).replace('.', ',')} g/mL, aproximada pela classe do
                pó. Serve para ordem de grandeza; o valor do lote quem informa é a farmácia.
              </span>
            ) : (
              <span>
                Densidade {String(c.bulkDensity).replace('.', ',')} g/mL
                {c.densitySource ? ` · ${c.densitySource}` : ''}.
              </span>
            )}
          </p>
        </section>

        <section>
          <h3 className="flex items-center gap-2 text-sm font-semibold">
            <BookOpen className="h-4 w-4 text-muted-foreground" />
            De onde saiu
          </h3>
          {isLoading && <Skeleton className="mt-2 h-16 w-full" />}
          {!isLoading && evidence.length === 0 && (
            <p className="mt-1 text-sm text-muted-foreground">
              Sem trecho do material ligado a esta substância. O conteúdo veio de pesquisa externa.
            </p>
          )}
          <ul className="mt-2 space-y-2">
            {evidence.slice(0, 6).map((e) => (
              <li key={e.id} className="rounded-lg border">
                <button
                  type="button"
                  className="flex w-full items-start justify-between gap-3 p-3 text-left"
                  onClick={() => setExpanded(expanded === e.id ? null : e.id)}
                >
                  <span className="min-w-0">
                    <span className="block text-sm font-medium">{e.article?.title}</span>
                    <span className="text-xs text-muted-foreground">
                      {e.article?.journal}
                      {e.similarity ? ` · ${(Number(e.similarity) * 100).toFixed(0)}% de proximidade` : ''}
                    </span>
                  </span>
                  <span className="shrink-0 text-xs text-muted-foreground">
                    {expanded === e.id ? 'fechar' : 'ver trecho'}
                  </span>
                </button>
                {expanded === e.id && (
                  <p className="whitespace-pre-wrap border-t p-3 text-xs leading-relaxed text-muted-foreground">
                    {e.excerpt}
                  </p>
                )}
              </li>
            ))}
          </ul>
        </section>
      </div>

      {/* Ação: a dose entra editável, então dá para ajustar antes de inserir sem voltar à fórmula. */}
      <div className="flex items-end gap-3 border-t p-4">
        <div className="w-32 space-y-1">
          <Label className="text-xs">Dose</Label>
          <div className="flex items-center gap-2">
            <Input
              type="number"
              step="any"
              value={dose}
              onChange={(e) => onDoseChange(e.target.value)}
              placeholder="—"
            />
            <span className="text-sm text-muted-foreground">{c.defaultUnit}</span>
          </div>
        </div>
        <Button className="flex-1" onClick={onUse}>
          Usar nesta fórmula
        </Button>
      </div>
    </>
  )
}

/** Faixa de dose desenhada: onde a dose usual cai entre o mínimo e o máximo cadastrados. */
function DoseRange({
  min,
  max,
  usual,
  unit,
}: {
  min?: number
  max?: number
  usual?: number
  unit: string
}) {
  const lo = min ?? usual ?? 0
  const hi = max ?? usual ?? 0
  const span = hi - lo
  const pos = usual !== undefined && span > 0 ? Math.min(100, Math.max(0, ((usual - lo) / span) * 100)) : null

  return (
    <div>
      <div className="relative h-1.5 rounded-full bg-muted">
        <div className="absolute inset-y-0 left-0 right-0 rounded-full bg-primary/20" />
        {pos !== null && (
          <div
            className="absolute top-1/2 h-3 w-3 -translate-y-1/2 rounded-full border-2 border-background bg-primary"
            style={{ left: `calc(${pos}% - 6px)` }}
          />
        )}
      </div>
      <div className="mt-1 flex justify-between text-xs text-muted-foreground">
        <span>{min ? `${fmt(min)} ${unit}` : ''}</span>
        <span>{max ? `${fmt(max)} ${unit}` : ''}</span>
      </div>
    </div>
  )
}

/**
 * Seção em tópicos com o texto completo a um clique.
 *
 * O tópico é para decidir; o texto é para conferir. Mostrar os dois ao mesmo tempo faria o painel
 * virar a parede de texto que ele veio substituir.
 */
function BulletSection({
  icon,
  title,
  bullets,
  fullText,
}: {
  icon?: React.ReactNode
  title: string
  bullets: string[]
  fullText?: string
}) {
  const [showText, setShowText] = useState(false)
  if (bullets.length === 0 && !fullText) return null

  return (
    <section>
      <div className="flex items-center justify-between gap-2">
        <h3 className="flex items-center gap-2 text-sm font-semibold">
          {icon}
          {title}
        </h3>
        {fullText && bullets.length > 0 && (
          <button
            type="button"
            className="text-xs text-muted-foreground underline-offset-2 hover:underline"
            onClick={() => setShowText((v) => !v)}
          >
            {showText ? 'ver tópicos' : 'ver texto completo'}
          </button>
        )}
      </div>

      {showText || bullets.length === 0 ? (
        <p className="mt-1 text-sm leading-relaxed text-muted-foreground">{fullText}</p>
      ) : (
        <ul className="mt-2 space-y-1.5">
          {bullets.map((b) => (
            <li key={b} className="flex gap-2 text-sm leading-snug">
              <span className="mt-1.5 h-1.5 w-1.5 shrink-0 rounded-full bg-primary/50" />
              <span>{b}</span>
            </li>
          ))}
        </ul>
      )}
    </section>
  )
}
