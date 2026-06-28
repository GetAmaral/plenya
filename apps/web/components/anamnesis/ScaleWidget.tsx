'use client'

import { useEffect, useMemo, useState } from 'react'
import {
  type ScaleDef,
  type ChosenWord,
  questionOptions,
  scaleTotal,
  scaleAnsweredCount,
  isScaleComplete,
} from '@plenya/domain'
import { cn } from '@/lib/utils'
import { ScaleAdminInput } from './ScaleAdminWidgets'

export interface ScaleWidgetProps {
  def: ScaleDef
  /** Classifica o total no nível do score item (operador/limites do banco). */
  classify: (total: number) => number | undefined
  /** Nome do nível (para exibir a classificação). */
  levelName?: (level: number) => string | undefined
  /** Respostas iniciais por índice de pergunta (F3 reidrata daqui). */
  initialAnswers?: Record<number, number>
  /** Palavras sorteadas (word_recall) — resolvidas pelo renderer e iguais p/ imediato+tardio. */
  adminWords?: ChosenWord[]
  /** Dispara quando completa (level classificado) ou limpa (undefined). */
  onResult: (result: {
    level: number | undefined
    total: number
    answers: Record<number, number>
    words?: string[]
  }) => void
  compact?: boolean
}

/**
 * Widget genérico de escala somada (Tier A): renderiza cada pergunta com suas opções,
 * soma e classifica automaticamente pelo nível do score item — a mesma experiência do site.
 * A classificação vem dos `levels` do banco (via `classify`), não de regra hardcoded.
 */
export function ScaleWidget({
  def,
  classify,
  levelName,
  initialAnswers,
  adminWords,
  onResult,
  compact = false,
}: ScaleWidgetProps) {
  const [answers, setAnswers] = useState<Record<number, number>>(initialAnswers ?? {})

  const admin = def.kind === 'administered' ? def.administration : undefined
  const isWordRecall = admin?.type === 'word_recall'
  const answeredCount = scaleAnsweredCount(answers)
  // word_recall só classifica quando todas as palavras foram julgadas; demais, com 1+ resposta.
  const complete = isWordRecall
    ? answeredCount === admin.categories.length
    : isScaleComplete(def, answers)
  const total = useMemo(() => scaleTotal(def, answers), [def, answers])
  // ASEX e afins classificam por regra de item (não pelo total) via def.classifyAnswers.
  const level = complete
    ? def.classifyAnswers
      ? def.classifyAnswers(answers)
      : classify(total)
    : undefined
  const words = isWordRecall ? adminWords?.map((w) => w.word) : undefined

  // Reporta resultado sempre que o conjunto de respostas muda.
  useEffect(() => {
    onResult({ level: complete ? level : undefined, total, answers, words })
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [answers, complete, total])

  const choose = (qIdx: number, value: number) => {
    setAnswers((prev) => {
      // toggle: clicar de novo na mesma opção desmarca
      if (prev[qIdx] === value) {
        const next = { ...prev }
        delete next[qIdx]
        return next
      }
      return { ...prev, [qIdx]: value }
    })
  }

  const clear = () => setAnswers({})

  return (
    <div className="space-y-3">
      {def.instructions && (
        <p className="text-xs leading-relaxed text-muted-foreground">{def.instructions}</p>
      )}

      {admin ? (
        <ScaleAdminInput
          spec={admin}
          answers={answers}
          onChange={setAnswers}
          compact={compact}
          chosenWords={adminWords}
        />
      ) : (
      <div className="space-y-2.5">
        {(def.questions ?? []).map((q, qIdx) => {
          const opts = questionOptions(def, qIdx)
          return (
            <div key={qIdx} className="rounded-lg border bg-card p-3">
              <p className="mb-2 text-[13px] leading-snug text-foreground">
                <span className="mr-1.5 font-semibold text-primary">{qIdx + 1}.</span>
                {q.text}
              </p>
              <div className={cn('grid gap-1.5', compact ? 'grid-cols-2' : 'grid-cols-2 sm:grid-cols-4')}>
                {opts.map((opt) => {
                  const selected = answers[qIdx] === opt.value
                  return (
                    <button
                      key={opt.value}
                      type="button"
                      data-level-button={qIdx === 0 ? true : undefined}
                      aria-pressed={selected}
                      onClick={() => choose(qIdx, opt.value)}
                      className={cn(
                        'rounded-md border px-2.5 py-1.5 text-left text-xs leading-snug transition-all',
                        selected
                          ? 'border-primary bg-primary text-primary-foreground'
                          : 'border-border bg-background hover:border-primary/50'
                      )}
                    >
                      <span className="mr-1 font-semibold tabular-nums opacity-70">{opt.value}</span>
                      {opt.label}
                    </button>
                  )
                })}
              </div>
            </div>
          )
        })}
      </div>
      )}

      {/* Sem bloco de resultado aqui: o total e o nível classificado ficam no cabeçalho
          do item (renderer), junto de todos os níveis auto-marcados. Aqui só as perguntas. */}
      {answeredCount > 0 && (
        <div className="flex justify-end border-t pt-2">
          <button
            type="button"
            onClick={clear}
            className="text-xs text-muted-foreground underline underline-offset-4 transition hover:text-foreground"
          >
            limpar respostas
          </button>
        </div>
      )}
    </div>
  )
}
