'use client'

import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { AlertTriangle, Check, Loader2, UserCheck } from 'lucide-react'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Skeleton } from '@/components/ui/skeleton'
import { PageHeader } from '@/components/layout/page-header'
import { useRequireAuth } from '@/lib/use-auth'
import {
  curateMedicationSubstance,
  getMedicationReviewQueue,
  type ReviewQueueItem,
} from '@/lib/api/medication-definitions'

const CATEGORIES: { value: ReviewQueueItem['category']; label: string; hint: string }[] = [
  { value: 'simple', label: 'Receita simples', hint: 'validade 30 dias' },
  { value: 'antibiotic', label: 'Antimicrobiano', hint: 'validade 10 dias' },
  { value: 'c1', label: 'Controle especial (C1)', hint: 'receita retida, exige assinatura' },
  { value: 'c5', label: 'Anabolizante (C5)', hint: 'receita retida, exige assinatura' },
  { value: 'glp1', label: 'GLP-1', hint: 'validade 90 dias' },
  { value: 'a_b', label: 'Notificação A/B', hint: 'fora do receituário do EMR' },
]

/**
 * Conferência do catálogo importado da ANVISA.
 *
 * O import classifica ~78% das apresentações com regra defensável e marca o resto — o que a
 * Lista de Preços não permite afirmar. Esta tela é onde essa dúvida vira decisão: uma linha
 * por SUBSTÂNCIA (não por apresentação, senão seriam 5.900 decisões em vez de ~1.000), na
 * ordem em que importa conferir.
 *
 * Confirmar e corrigir são a mesma operação: as duas carimbam a curadoria, que é o que impede
 * o reimport mensal da ANVISA de desfazer a decisão do médico.
 */
export default function MedicationReviewPage() {
  useRequireAuth()
  const queryClient = useQueryClient()
  const [overrides, setOverrides] = useState<Record<string, ReviewQueueItem['category']>>({})

  const { data, isLoading } = useQuery({
    queryKey: ['medication-review-queue'],
    queryFn: () => getMedicationReviewQueue({ limit: 25 }),
  })

  const curate = useMutation({
    mutationFn: curateMedicationSubstance,
    onSuccess: (res) => {
      toast.success(`${res.activeIngredient}: ${res.updated} apresentação(ões) conferidas`)
      queryClient.invalidateQueries({ queryKey: ['medication-review-queue'] })
      queryClient.invalidateQueries({ queryKey: ['medication-search'] })
    },
    onError: (error: any) => {
      toast.error('Não deu para salvar', { description: error?.message })
    },
  })

  const items = data?.items ?? []

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[
          { label: 'Admin', href: '/admin' },
          { label: 'Medicamentos', href: '/admin/medication-definitions' },
          { label: 'Conferência' },
        ]}
        title="Conferência do catálogo"
        description={
          data
            ? `${data.total} substâncias aguardando conferência`
            : 'Substâncias que a lista da ANVISA não permite classificar com segurança'
        }
      />

      <Card className="border-amber-200 bg-amber-50/50">
        <CardContent className="pt-6 text-sm text-amber-900">
          A lista da ANVISA traz a tarja, não as listas da Portaria 344. Quando ela publica a
          apresentação <strong>sem tarja nenhuma</strong>, o import assume receita simples e
          marca aqui — é o caso onde errar significa emitir receita simples para algo que
          precisava de controle especial. Conferir uma substância vale para todas as
          apresentações dela e sobrevive à atualização mensal da lista.
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Fila</CardTitle>
        </CardHeader>
        <CardContent className="space-y-3">
          {isLoading && (
            <>
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-20 w-full" />
            </>
          )}

          {!isLoading && items.length === 0 && (
            <p className="py-8 text-center text-muted-foreground">
              Nada pendente. O catálogo está conferido.
            </p>
          )}

          {items.map((item) => {
            const chosen = overrides[item.activeIngredient] ?? item.category
            const changed = chosen !== item.category
            const busy = curate.isPending && curate.variables?.activeIngredient === item.activeIngredient

            return (
              <div
                key={item.activeIngredient}
                className="flex flex-col gap-3 rounded-lg border p-3 sm:flex-row sm:items-center sm:justify-between"
              >
                <div className="min-w-0 flex-1 space-y-1">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="font-medium">{item.activeIngredient}</span>
                    {item.usedByPatients && (
                      <Badge variant="outline" className="gap-1 bg-sky-50 text-xs text-sky-700">
                        <UserCheck className="h-3 w-3" />
                        já usada por paciente
                      </Badge>
                    )}
                    {item.categorySource === 'cmed_fallback' && (
                      <Badge variant="outline" className="gap-1 bg-amber-50 text-xs text-amber-700">
                        <AlertTriangle className="h-3 w-3" />
                        deduzida sem tarja
                      </Badge>
                    )}
                  </div>

                  <p className="truncate text-sm text-muted-foreground" title={item.sampleProducts}>
                    {item.sampleProducts}
                  </p>

                  <p className="text-xs text-muted-foreground">
                    {item.presentations} apresentação(ões) ·{' '}
                    {item.stripe ? `tarja ${item.stripe.replace('_', ' ')}` : 'sem tarja publicada'}
                    {item.therapeuticClass ? ` · ${item.therapeuticClass}` : ''}
                  </p>
                </div>

                <div className="flex shrink-0 items-center gap-2">
                  <Select
                    value={chosen}
                    onValueChange={(v) =>
                      setOverrides((prev) => ({
                        ...prev,
                        [item.activeIngredient]: v as ReviewQueueItem['category'],
                      }))
                    }
                  >
                    <SelectTrigger className="w-full sm:w-[220px]">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      {CATEGORIES.map((c) => (
                        <SelectItem key={c.value} value={c.value}>
                          {c.label}
                          <span className="ml-2 text-xs text-muted-foreground">{c.hint}</span>
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>

                  <Button
                    size="sm"
                    variant={changed ? 'default' : 'outline'}
                    disabled={busy}
                    onClick={() =>
                      curate.mutate({
                        activeIngredient: item.activeIngredient,
                        category: chosen,
                        // Tarja preta continua fora do receituário; as demais voltam a ser
                        // prescritíveis quando o médico reclassifica.
                        isPrescribable: chosen !== 'a_b',
                      })
                    }
                  >
                    {busy ? (
                      <Loader2 className="h-4 w-4 animate-spin" />
                    ) : (
                      <>
                        <Check className="mr-1 h-4 w-4" />
                        {changed ? 'Corrigir' : 'Confirmar'}
                      </>
                    )}
                  </Button>
                </div>
              </div>
            )
          })}
        </CardContent>
      </Card>
    </div>
  )
}
