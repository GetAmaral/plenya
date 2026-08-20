'use client'

import { useEffect, useState } from 'react'
import { useMutation, useQuery } from '@tanstack/react-query'
import { AlertTriangle, Info, Loader2, Pill, Save, TriangleAlert } from 'lucide-react'
import { toast } from 'sonner'

import { Button } from '@/components/ui/button'
import {
  checkMagistralFormula,
  saveMagistralDefaultDose,
  toBullets,
  type MagistralAlert,
  type MagistralCheckRequest,
} from '@/lib/api/magistral-components'
import type { FormulaFormData } from '@/lib/validations/compounded-prescription'

interface FormulaCheckPanelProps {
  formula: FormulaFormData
  /** Marca (ou desmarca) que a dose daquela substância é do elemento, não do insumo. */
  onMarkElemental?: (substance: string, elemental: boolean) => void
  /** Aplica no componente a categoria de receita que a substância carrega. */
  onSetCategory?: (substance: string, category: string) => void
}

const LEVEL_STYLE: Record<MagistralAlert['level'], string> = {
  avoid: 'border-destructive/40 bg-destructive/5 text-destructive',
  warn: 'border-amber-300 bg-amber-50 text-amber-800',
  info: 'border-sky-200 bg-sky-50 text-sky-800',
}

/**
 * Painel de conferência da fórmula: incompatibilidade, palatabilidade e tamanho de cápsula.
 *
 * Duas regras de projeto visíveis aqui:
 *   · AVISA, NÃO BLOQUEIA — mesma escolha do alerta de alergia que já existe no EMR;
 *   · quando falta densidade, a calculadora NÃO opina e diz de quem falta. Número inventado numa
 *     receita é pior que silêncio.
 */
export function FormulaCheckPanel({ formula, onMarkElemental, onSetCategory }: FormulaCheckPanelProps) {
  // Debounce: o painel acompanha a digitação sem bater na API a cada tecla.
  const [debounced, setDebounced] = useState<MagistralCheckRequest | null>(null)

  // Sem useMemo com `[formula]` de dependência: o `watch` do react-hook-form devolve SEMPRE a
  // mesma referência do objeto interno, então o memo nunca invalidava e o painel ficava preso ao
  // primeiro render (fórmula vazia). O payload é barato de montar a cada render.
  const payload = ((): MagistralCheckRequest | null => {
    const components = (formula.components ?? [])
      .filter((c) => (c.substance || '').trim().length >= 2)
      .map((c) => ({
        magistralComponentId: c.magistralComponentId,
        substance: c.substance.trim(),
        quantity: Number(c.quantity) || 0,
        unit: c.unit,
        asElemental: c.asElemental,
      }))
    if (components.length === 0 || !formula.pharmaceuticalForm) return null
    return {
      pharmaceuticalForm: formula.pharmaceuticalForm,
      components,
      // A conta de tomadas por dia é do backend: era a mesma regra escrita duas vezes, e duas
      // implementações da mesma regra divergem com o tempo.
      posology: formula.posology,
      vehicle: formula.vehicle,
    }
  })()

  const payloadKey = payload ? JSON.stringify(payload) : ''
  useEffect(() => {
    const t = setTimeout(() => setDebounced(payload), 500)
    return () => clearTimeout(t)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [payloadKey])

  // A chave da query é o payload JÁ debounced, não o atual: com a chave no payload vivo, o
  // TanStack disparava a busca com o valor ANTIGO e não refazia quando o debounce chegava —
  // o painel ficava um passo atrás do que estava na tela.
  const debouncedKey = debounced ? JSON.stringify(debounced) : ''
  const { data, isFetching } = useQuery({
    queryKey: ['magistral-check', debouncedKey],
    queryFn: () => checkMagistralFormula(debounced as MagistralCheckRequest),
    enabled: !!debounced,
  })

  const saveDose = useMutation({
    mutationFn: saveMagistralDefaultDose,
    onSuccess: (c) =>
      toast.success(`Dose padrão de ${c.name} salva`, {
        description: 'Da próxima vez ela vem preenchida.',
      }),
    onError: (e: any) => toast.error('Não deu para salvar', { description: e?.message }),
  })

  if (!payload) return null

  const alerts = data?.alerts ?? []
  const capsule = data?.capsule
  // Só oferece guardar padrão para substância que o médico já dosou e o catálogo não conhece.
  const missingDose = (data?.components ?? []).filter((c) => !c.hasDose)

  return (
    <div className="space-y-2 rounded-lg border border-dashed p-3">
      <div className="flex items-center gap-2 text-sm font-semibold">
        Conferência da fórmula
        {isFetching && <Loader2 className="h-3 w-3 animate-spin text-muted-foreground" />}
      </div>

      {alerts.length === 0 && !isFetching && (
        <p className="text-xs text-muted-foreground">
          Nada a apontar nas substâncias catalogadas. O catálogo cobre só o que já foi cadastrado,
          então ausência de aviso não é atestado de compatibilidade.
        </p>
      )}

      {alerts.map((a, i) => (
        <div key={i} className={`flex gap-2 rounded-md border p-2 text-xs ${LEVEL_STYLE[a.level]}`}>
          {a.level === 'info' ? (
            <Info className="mt-0.5 h-3.5 w-3.5 shrink-0" />
          ) : a.level === 'avoid' ? (
            <TriangleAlert className="mt-0.5 h-3.5 w-3.5 shrink-0" />
          ) : (
            <AlertTriangle className="mt-0.5 h-3.5 w-3.5 shrink-0" />
          )}
          <span className="flex-1">
            {a.message}
            {/* O alerta de correção é o único acionável: resolver é dizer o que a dose significa. */}
            {a.kind === 'correction' && onMarkElemental && (
              <Button
                type="button"
                size="sm"
                variant="outline"
                className="ml-2 h-6 px-2 text-[11px]"
                onClick={() => onMarkElemental(a.substances[0], true)}
              >
                é dose do elemento
              </Button>
            )}
          </span>
        </div>
      ))}

      {capsule && (
        <div className="flex gap-2 rounded-md border bg-muted/40 p-2 text-xs">
          <Pill className="mt-0.5 h-3.5 w-3.5 shrink-0 text-muted-foreground" />
          <div>
            <p>{capsule.explanation}</p>
            {capsule.decided && capsule.size && (
              <p className="mt-1 font-semibold">
                {capsule.capsulesPerDose} cápsula(s) tamanho {capsule.size}
                {capsule.sizeIfCompacted ? ` (ou ${capsule.sizeIfCompacted})` : ''}
              </p>
            )}
            <p className="mt-1 text-[11px] text-muted-foreground">
              Auxílio de consultório. Quem especifica a manipulação é a farmácia.
            </p>
          </div>
        </div>
      )}

      {/* Posologia de referência do catálogo, ao lado da fórmula que está sendo montada. Quando
          veio do material e ainda não foi conferida, a tela diz isso em vez de apresentar como
          fato do consultório. */}
      {(data?.components ?? [])
        .filter(
          (c) =>
            c.doseReference ||
            c.indications ||
            c.indicationBullets ||
            c.preferredName ||
            c.rawMaterialMg ||
            c.regulatoryNote ||
            (c.defaultCategory && c.defaultCategory !== 'simple')
        )
        .map((c) => (
          <div key={`ref-${c.substance}`} className="rounded-md border bg-muted/30 p-2 text-xs">
            <span className="font-medium">{c.substance}</span>
            {c.evidenceStatus === 'suggested' && (
              <span className="ml-2 rounded bg-sky-100 px-1 py-0.5 text-[10px] text-sky-700">
                do material, a conferir
              </span>
            )}
            {/* Aqui é lembrete durante a montagem da fórmula, não leitura: dois tópicos de
                indicação e dois de posologia. O texto inteiro fica no painel do catálogo. */}
            <div className="mt-1 grid gap-1 sm:grid-cols-2">
              <ul className="space-y-0.5">
                {toBullets(c.indicationBullets)
                  .slice(0, 2)
                  .map((b) => (
                    <li key={b} className="flex gap-1.5 text-muted-foreground">
                      <span className="text-primary/60">·</span>
                      <span>{b}</span>
                    </li>
                  ))}
              </ul>
              <ul className="space-y-0.5">
                {toBullets(c.doseBullets)
                  .slice(0, 2)
                  .map((b) => (
                    <li key={b} className="flex gap-1.5 text-muted-foreground">
                      <span className="text-primary/60">→</span>
                      <span>{b}</span>
                    </li>
                  ))}
              </ul>
            </div>
            {toBullets(c.indicationBullets).length === 0 && c.indications && (
              <p className="mt-0.5 line-clamp-2 text-muted-foreground">{c.indications}</p>
            )}

            {/* Conversão elemento → insumo: é o que a farmácia pesa e o que ocupa a cápsula. */}
            {c.rawMaterialMg && (
              <p className="mt-1 rounded bg-sky-50 px-2 py-1 text-[11px] text-sky-800">
                Dose do elemento equivale a{' '}
                <strong>{Math.round(c.rawMaterialMg)} mg do insumo</strong>
                {c.elementalPercent
                  ? ` (${String(c.elementalPercent).replace('.', ',')}% de ativo)`
                  : ''}
                .
              </p>
            )}

            {/* Categoria de receita e restrição de finalidade. Sem isto, "testosterona 50 mg"
                numa fórmula magistral sai como receita simples, quando a Portaria 344/98 pede
                Controle Especial em duas vias. */}
            {c.defaultCategory && c.defaultCategory !== 'simple' && (
              <p className="mt-1 rounded bg-destructive/5 px-2 py-1 text-[11px] text-destructive">
                Substância de <strong>{c.defaultCategory.toUpperCase()}</strong>: a receita precisa
                sair como controlada.
                {onSetCategory && (
                  <Button
                    type="button"
                    size="sm"
                    variant="outline"
                    className="ml-2 h-6 px-2 text-[11px]"
                    onClick={() => onSetCategory(c.substance, c.defaultCategory!)}
                  >
                    marcar {c.defaultCategory.toUpperCase()}
                  </Button>
                )}
              </p>
            )}
            {c.regulatoryNote && (
              <p className="mt-1 rounded bg-amber-50 px-2 py-1 text-[11px] text-amber-900">
                {c.regulatoryNote}
              </p>
            )}

            {/* A forma que ele usa. A nota traz o que a literatura sustenta, inclusive quando não
                sustenta — é dele a decisão, mas com a evidência à vista. */}
            {c.preferredName && (
              <p className="mt-1 rounded bg-amber-50 px-2 py-1 text-[11px] text-amber-900">
                Você costuma usar <strong>{c.preferredName}</strong>.
                {c.preferenceNote ? ` ${c.preferenceNote}` : ''}
              </p>
            )}
          </div>
        ))}

      {missingDose.length > 0 && (
        <div className="flex flex-wrap gap-2 pt-1">
          {missingDose.map((c) => {
            const typed = (formula.components ?? []).find(
              (fc) => fc.substance.trim().toLowerCase() === c.substance.toLowerCase()
            )
            if (!typed || !Number(typed.quantity)) return null
            return (
              <Button
                key={c.substance}
                type="button"
                variant="outline"
                size="sm"
                disabled={saveDose.isPending}
                onClick={() =>
                  saveDose.mutate({
                    substance: c.substance,
                    dose: Number(typed.quantity),
                    unit: typed.unit,
                  })
                }
              >
                <Save className="mr-1 h-3 w-3" />
                Salvar {String(typed.quantity).replace('.', ',')} {typed.unit} como padrão de{' '}
                {c.substance}
              </Button>
            )
          })}
        </div>
      )}
    </div>
  )
}
