'use client'

import type { AdministrationSpec, ChosenWord } from '@plenya/domain'
import { cn } from '@/lib/utils'

type Answers = Record<number, number>

interface AdminProps {
  spec: AdministrationSpec
  answers: Answers
  onChange: (next: Answers) => void
  compact?: boolean
  /** Palavras sorteadas para esta aplicação (word_recall). Resolvidas pelo renderer e
   *  compartilhadas entre imediato e tardio. */
  chosenWords?: ChosenWord[]
}

/** Roteia para o widget de administração certo conforme o tipo. */
export function ScaleAdminInput({ spec, answers, onChange, compact, chosenWords }: AdminProps) {
  if (spec.type === 'word_recall') {
    // Fallback defensivo: se o renderer não resolveu o set, usa a 1ª palavra de cada categoria.
    const words: ChosenWord[] =
      chosenWords && chosenWords.length > 0
        ? chosenWords
        : spec.categories.map((c) => ({ category: c.category, word: c.words[0] }))
    return <WordRecallAdmin words={words} answers={answers} onChange={onChange} compact={compact} />
  }
  return <DigitSpanAdmin spec={spec} answers={answers} onChange={onChange} compact={compact} />
}

// ---------------------------------------------------------------------------
// Evocação de palavras (Dubois)
// ---------------------------------------------------------------------------

const RECALL_OPTIONS = [
  { value: 0, label: 'Não lembrou' },
  { value: 1, label: 'Com dica' },
  { value: 2, label: 'Espontânea' },
]

function WordRecallAdmin({
  words,
  answers,
  onChange,
  compact,
}: {
  words: ChosenWord[]
  answers: Answers
  onChange: (next: Answers) => void
  compact?: boolean
}) {
  const set = (idx: number, value: number) => {
    const next = { ...answers }
    if (next[idx] === value) delete next[idx]
    else next[idx] = value
    onChange(next)
  }

  return (
    <div className="space-y-3">
      {/* Referência de apresentação */}
      <div className="rounded-lg border border-dashed bg-muted/40 p-3">
        <p className="mb-1.5 text-[11px] font-semibold uppercase tracking-wide text-muted-foreground">
          Apresente ao paciente
        </p>
        <div className="flex flex-wrap gap-x-4 gap-y-1 text-sm">
          {words.map((w, i) => (
            <span key={i} className="text-foreground">
              <span className="font-semibold">{w.word}</span>
              <span className="text-muted-foreground"> ({w.category})</span>
            </span>
          ))}
        </div>
      </div>

      {/* Evocação por palavra */}
      <div className="space-y-2">
        {words.map((w, idx) => (
          <div key={idx} className="rounded-lg border bg-card p-2.5">
            <p className="mb-2 text-[13px] text-foreground">
              <span className="font-semibold">{w.word}</span>
              <span className="text-muted-foreground"> · {w.category}</span>
            </p>
            <div className={cn('grid gap-1.5', compact ? 'grid-cols-3' : 'grid-cols-3')}>
              {RECALL_OPTIONS.map((opt) => {
                const selected = answers[idx] === opt.value
                return (
                  <button
                    key={opt.value}
                    type="button"
                    data-level-button={idx === 0 ? true : undefined}
                    aria-pressed={selected}
                    onClick={() => set(idx, opt.value)}
                    className={cn(
                      'rounded-md border px-2 py-1.5 text-center text-xs transition-all',
                      selected
                        ? opt.value === 0
                          ? 'border-muted-foreground bg-muted text-foreground'
                          : 'border-primary bg-primary text-primary-foreground'
                        : 'border-border bg-background hover:border-primary/50'
                    )}
                  >
                    {opt.label}
                  </button>
                )
              })}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

// ---------------------------------------------------------------------------
// Span de dígitos
// ---------------------------------------------------------------------------

function DigitSpanAdmin({
  spec,
  answers,
  onChange,
  compact,
}: {
  spec: Extract<AdministrationSpec, { type: 'digit_span' }>
  answers: Answers
  onChange: (next: Answers) => void
  compact?: boolean
}) {
  const lengths = Object.keys(spec.sequencesByLength)
    .map(Number)
    .sort((a, b) => a - b)

  const judge = (len: number, ok: 0 | 1) => {
    const next = { ...answers }
    if (next[len] === ok) delete next[len]
    else next[len] = ok
    onChange(next)
  }

  return (
    <div className="space-y-2">
      {lengths.map((len) => {
        const seqs = spec.sequencesByLength[len]
        const verdict = answers[len]
        return (
          <div
            key={len}
            className={cn(
              'flex flex-wrap items-center gap-x-3 gap-y-1.5 rounded-lg border p-2.5',
              verdict === 1 && 'border-primary/40 bg-primary/5'
            )}
          >
            <span className="text-[11px] font-semibold text-muted-foreground">{len} dígitos</span>
            <span className="flex flex-wrap gap-x-3 font-mono text-sm tracking-widest text-foreground">
              {seqs.map((s, i) => (
                <span key={i}>{s.join(' ')}</span>
              ))}
            </span>
            <div className="ml-auto flex gap-1.5">
              <button
                type="button"
                data-level-button={len === lengths[0] ? true : undefined}
                aria-pressed={verdict === 1}
                onClick={() => judge(len, 1)}
                className={cn(
                  'rounded-md border px-2.5 py-1 text-xs transition-all',
                  verdict === 1
                    ? 'border-primary bg-primary text-primary-foreground'
                    : 'border-border bg-background hover:border-primary/50'
                )}
              >
                Acertou
              </button>
              <button
                type="button"
                aria-pressed={verdict === 0}
                onClick={() => judge(len, 0)}
                className={cn(
                  'rounded-md border px-2.5 py-1 text-xs transition-all',
                  verdict === 0
                    ? 'border-muted-foreground bg-muted text-foreground'
                    : 'border-border bg-background hover:border-primary/50'
                )}
              >
                Errou
              </button>
            </div>
          </div>
        )
      })}
    </div>
  )
}
