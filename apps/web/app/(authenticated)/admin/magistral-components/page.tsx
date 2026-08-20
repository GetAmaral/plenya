'use client'

import { useState } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { BookOpen, Check, Loader2, Pin, Plus, Save, Search, Trash2 } from 'lucide-react'
import { toast } from 'sonner'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Checkbox } from '@/components/ui/checkbox'
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
  confirmMagistralComponent,
  getMagistralEvidence,
  pinMagistralEvidence,
  searchMagistralComponents,
  updateMagistralComponent,
  type MagistralComponent,
} from '@/lib/api/magistral-components'

interface Incompatibility {
  id: string
  severity: 'info' | 'warn' | 'avoid'
  mechanism: string
  note?: string
  source?: string
  componentA?: { id: string; name: string }
  componentB?: { id: string; name: string }
}

/**
 * Curadoria do catálogo magistral.
 *
 * Existe por um motivo concreto: a calculadora de cápsula só opina com a DENSIDADE APARENTE
 * cadastrada, e esse número não vem de base pública nenhuma — varia por lote e por farmácia.
 * Aqui é onde ele entra, junto dos sinalizadores que geram os avisos de compatibilidade.
 *
 * Campo vazio significa "não cadastrado" e continua vazio: em nenhum lugar um branco vira zero.
 */
export default function MagistralComponentsAdminPage() {
  useRequireAuth()
  const queryClient = useQueryClient()
  const [term, setTerm] = useState('')
  const [editing, setEditing] = useState<MagistralComponent | null>(null)

  const { data: components = [], isFetching } = useQuery({
    queryKey: ['magistral-admin', term],
    // Busca vazia lista o repertório por uso; a partir de 2 caracteres, filtra.
    queryFn: () => searchMagistralComponents(term),
  })

  const { data: pairs = [] } = useQuery({
    queryKey: ['magistral-incompatibilities'],
    queryFn: () => apiClient.get<Incompatibility[]>('/api/v1/magistral-components/incompatibilities'),
  })

  const save = useMutation({
    mutationFn: (c: MagistralComponent) => updateMagistralComponent(c.id, c),
    onSuccess: () => {
      toast.success('Componente atualizado')
      setEditing(null)
      queryClient.invalidateQueries({ queryKey: ['magistral-admin'] })
    },
    onError: (e: any) => toast.error('Não deu para salvar', { description: e?.message }),
  })

  const confirm = useMutation({
    mutationFn: confirmMagistralComponent,
    onSuccess: () => {
      toast.success('Marcado como conferido')
      setEditing(null)
      queryClient.invalidateQueries({ queryKey: ['magistral-admin'] })
    },
    onError: (e: any) => toast.error('Não deu para confirmar', { description: e?.message }),
  })

  const removePair = useMutation({
    mutationFn: (id: string) =>
      apiClient.delete(`/api/v1/magistral-components/incompatibilities/${id}`),
    onSuccess: () => {
      toast.success('Par removido')
      queryClient.invalidateQueries({ queryKey: ['magistral-incompatibilities'] })
    },
  })

  const [newPair, setNewPair] = useState({ a: '', b: '', severity: 'warn', mechanism: '' })
  const createPair = useMutation({
    mutationFn: () =>
      apiClient.post('/api/v1/magistral-components/incompatibilities', {
        componentAId: newPair.a,
        componentBId: newPair.b,
        severity: newPair.severity,
        mechanism: newPair.mechanism,
      }),
    onSuccess: () => {
      toast.success('Par cadastrado')
      setNewPair({ a: '', b: '', severity: 'warn', mechanism: '' })
      queryClient.invalidateQueries({ queryKey: ['magistral-incompatibilities'] })
    },
    onError: (e: any) => toast.error('Não deu para cadastrar', { description: e?.message }),
  })

  const num = (v?: number) => (v === undefined || v === null ? '' : String(v))

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Admin', href: '/admin' }, { label: 'Catálogo magistral' }]}
        title="Catálogo magistral"
        description="Faixa de dose, densidade aparente e sinalizadores que alimentam os avisos da fórmula"
      />

      <Card className="border-amber-200 bg-amber-50/50">
        <CardContent className="pt-6 text-sm text-amber-900">
          A <strong>densidade aparente</strong> não existe em base pública e muda com o lote e a
          compactação da farmácia. Enquanto ela não estiver cadastrada, a calculadora de cápsula
          diz que não sabe em vez de estimar. Cadastre o valor que a sua farmácia de confiança
          informar.
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Componentes</CardTitle>
          <CardDescription>
            Busque pelo nome. As substâncias entram sozinhas no catálogo quando você as prescreve.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-3">
          <div className="flex items-center gap-2">
            <Search className="h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Buscar substância..."
              value={term}
              onChange={(e) => setTerm(e.target.value)}
            />
            {isFetching && <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />}
          </div>

          <div className="divide-y rounded-lg border">
            {components.map((c) => (
              <div key={c.id} className="flex flex-wrap items-center justify-between gap-2 p-3">
                <div className="min-w-0">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="font-medium">{c.name}</span>
                    {c.bulkDensity ? (
                      <Badge
                        variant="outline"
                        className={`text-xs ${c.densitySource === 'classe' ? 'text-muted-foreground' : ''}`}
                        title={c.densitySource === 'classe' ? 'aproximada pela classe do pó' : c.densitySource}
                      >
                        {String(c.bulkDensity).replace('.', ',')} g/mL
                        {c.densitySource === 'classe' ? ' aprox.' : ''}
                      </Badge>
                    ) : (
                      <Badge variant="outline" className="bg-amber-50 text-xs text-amber-700">
                        sem densidade
                      </Badge>
                    )}
                    {c.evidenceStatus === 'suggested' && (
                      <Badge variant="outline" className="bg-sky-50 text-xs text-sky-700">
                        sugerido, a conferir
                      </Badge>
                    )}
                    {c.evidenceStatus === 'confirmed' && (
                      <Badge variant="outline" className="bg-green-50 text-xs text-green-700">
                        conferido
                      </Badge>
                    )}
                    {c.eutecticFormer && <Badge variant="outline" className="text-xs">eutético</Badge>}
                    {c.hygroscopic && <Badge variant="outline" className="text-xs">higroscópico</Badge>}
                    {c.oxidizing && <Badge variant="outline" className="text-xs">oxidante</Badge>}
                    {c.oxidationSensitive && (
                      <Badge variant="outline" className="text-xs">sensível a oxidação</Badge>
                    )}
                  </div>
                  {c.indications && (
                    <p className="mt-1 line-clamp-2 text-sm text-muted-foreground">{c.indications}</p>
                  )}
                  <p className="text-xs text-muted-foreground">
                    {c.usualDose ? `dose usual ${String(c.usualDose).replace('.', ',')} ${c.defaultUnit}` : 'sem dose cadastrada'}
                    {' · '}
                    {c.usageCount} uso(s) · origem {c.source}
                  </p>
                </div>
                <Button variant="outline" size="sm" onClick={() => setEditing({ ...c })}>
                  Editar
                </Button>
              </div>
            ))}
            {components.length === 0 && (
              <p className="p-6 text-center text-sm text-muted-foreground">
                Nada encontrado.
              </p>
            )}
          </div>
        </CardContent>
      </Card>

      {editing && (
        <Card>
          <CardHeader>
            <CardTitle>{editing.name}</CardTitle>
            <CardDescription>Deixe em branco o que você não souber — vazio é silêncio, não zero.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 gap-4 sm:grid-cols-4">
              <div className="space-y-2">
                <Label>Unidade padrão</Label>
                <Input
                  value={editing.defaultUnit}
                  onChange={(e) => setEditing({ ...editing, defaultUnit: e.target.value })}
                />
              </div>
              <div className="space-y-2">
                <Label>Dose usual</Label>
                <Input
                  type="number"
                  step="any"
                  value={num(editing.usualDose)}
                  onChange={(e) =>
                    setEditing({ ...editing, usualDose: e.target.value ? Number(e.target.value) : undefined })
                  }
                />
              </div>
              <div className="space-y-2">
                <Label>Dose mínima</Label>
                <Input
                  type="number"
                  step="any"
                  value={num(editing.minDose)}
                  onChange={(e) =>
                    setEditing({ ...editing, minDose: e.target.value ? Number(e.target.value) : undefined })
                  }
                />
              </div>
              <div className="space-y-2">
                <Label>Dose máxima</Label>
                <Input
                  type="number"
                  step="any"
                  value={num(editing.maxDose)}
                  onChange={(e) =>
                    setEditing({ ...editing, maxDose: e.target.value ? Number(e.target.value) : undefined })
                  }
                />
              </div>
            </div>

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-3">
              <div className="space-y-2">
                <Label>Densidade aparente (g/mL)</Label>
                <Input
                  type="number"
                  step="any"
                  placeholder="ex: 0,55"
                  value={num(editing.bulkDensity)}
                  onChange={(e) =>
                    setEditing({ ...editing, bulkDensity: e.target.value ? Number(e.target.value) : undefined })
                  }
                />
                <p className="text-xs text-muted-foreground">
                  {editing.densitySource === 'classe'
                    ? 'Hoje é uma aproximação pela classe do pó. Ao salvar um valor aqui, ele passa a valer como informado por você.'
                    : 'Sem este número a calculadora de cápsula não opina.'}
                </p>
              </div>
              <div className="space-y-2">
                <Label>Amargor</Label>
                <Select
                  value={editing.bitterness === undefined ? 'none' : String(editing.bitterness)}
                  onValueChange={(v) =>
                    setEditing({ ...editing, bitterness: v === 'none' ? undefined : Number(v) })
                  }
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="none">Não avaliado</SelectItem>
                    <SelectItem value="0">Sem amargor</SelectItem>
                    <SelectItem value="1">Leve</SelectItem>
                    <SelectItem value="2">Marcante</SelectItem>
                    <SelectItem value="3">Intolerável em sachê</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="space-y-2">
                <Label>Sinônimos (busca)</Label>
                <Input
                  value={editing.synonyms ?? ''}
                  onChange={(e) => setEditing({ ...editing, synonyms: e.target.value })}
                />
              </div>
            </div>

            <div className="flex flex-wrap gap-4">
              {(
                [
                  ['eutecticFormer', 'Forma mistura eutética'],
                  ['hygroscopic', 'Higroscópico'],
                  ['oxidizing', 'Oxidante'],
                  ['oxidationSensitive', 'Sensível a oxidação'],
                  ['photosensitive', 'Fotossensível'],
                ] as const
              ).map(([key, label]) => (
                <label key={key} className="flex items-center gap-2 text-sm">
                  <Checkbox
                    checked={!!editing[key]}
                    onCheckedChange={(v) => setEditing({ ...editing, [key]: !!v })}
                  />
                  {label}
                </label>
              ))}
            </div>

            <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
              <div className="space-y-2">
                <Label>Indicações em tópicos</Label>
                <Textarea
                  rows={5}
                  placeholder={'Um por linha:\nuso de estatina\ninsuficiência cardíaca'}
                  value={editing.indicationBullets ?? ''}
                  onChange={(e) => setEditing({ ...editing, indicationBullets: e.target.value })}
                />
                <p className="text-xs text-muted-foreground">
                  Uma linha por tópico. É o que aparece no painel durante a prescrição.
                </p>
              </div>
              <div className="space-y-2">
                <Label>Posologia em tópicos</Label>
                <Textarea
                  rows={5}
                  placeholder={'Um por linha:\n100 mg/dia com refeição gordurosa'}
                  value={editing.doseBullets ?? ''}
                  onChange={(e) => setEditing({ ...editing, doseBullets: e.target.value })}
                />
              </div>
            </div>

            <div className="space-y-2">
              <Label>Indicações (texto completo)</Label>
              <Textarea
                rows={3}
                placeholder="Para que você usa esta substância..."
                value={editing.indications ?? ''}
                onChange={(e) => setEditing({ ...editing, indications: e.target.value })}
              />
            </div>

            <div className="space-y-2">
              <Label>Posologia de referência (texto completo)</Label>
              <Textarea
                rows={2}
                placeholder="Ex: 1 cápsula ao deitar, iniciar com dose baixa..."
                value={editing.doseReference ?? ''}
                onChange={(e) => setEditing({ ...editing, doseReference: e.target.value })}
              />
            </div>

            <div className="space-y-2">
              <Label>Observações</Label>
              <Textarea
                rows={2}
                value={editing.notes ?? ''}
                onChange={(e) => setEditing({ ...editing, notes: e.target.value })}
              />
            </div>

            <EvidencePanel componentId={editing.id} />

            <div className="flex gap-3">
              <Button onClick={() => save.mutate(editing)} disabled={save.isPending}>
                {save.isPending ? (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <Save className="mr-2 h-4 w-4" />
                )}
                Salvar
              </Button>
              {editing.evidenceStatus === 'suggested' && (
                <Button
                  variant="outline"
                  onClick={() => confirm.mutate(editing.id)}
                  disabled={confirm.isPending}
                >
                  <Check className="mr-2 h-4 w-4" />
                  Conferido, sem mudar nada
                </Button>
              )}
              <Button variant="ghost" onClick={() => setEditing(null)}>
                Cancelar
              </Button>
            </div>
          </CardContent>
        </Card>
      )}

      <Card>
        <CardHeader>
          <CardTitle>Incompatibilidades curadas</CardTitle>
          <CardDescription>
            Pares que não saem de sinalizador. Poucas regras que nunca erram valem mais que muitas
            que gritam à toa.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-3">
          {pairs.map((p) => (
            <div key={p.id} className="flex items-start justify-between gap-3 rounded-lg border p-3">
              <div>
                {/* div, não p: o Badge renderiza uma div e <div> dentro de <p> quebra a
                    hidratação do React. */}
                <div className="flex flex-wrap items-center gap-2 text-sm font-medium">
                  <span>
                    {p.componentA?.name} × {p.componentB?.name}
                  </span>
                  <Badge variant="outline" className="text-xs">
                    {p.severity}
                  </Badge>
                </div>
                <p className="text-xs text-muted-foreground">{p.mechanism}</p>
                {p.source && <p className="text-xs text-muted-foreground">Fonte: {p.source}</p>}
              </div>
              <Button
                variant="ghost"
                size="sm"
                className="text-destructive"
                onClick={() => removePair.mutate(p.id)}
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          ))}
          {pairs.length === 0 && (
            <p className="text-sm text-muted-foreground">Nenhum par cadastrado ainda.</p>
          )}

          <div className="grid grid-cols-1 gap-2 rounded-lg border border-dashed p-3 sm:grid-cols-5">
            <Select value={newPair.a} onValueChange={(v) => setNewPair({ ...newPair, a: v })}>
              <SelectTrigger>
                <SelectValue placeholder="Substância A" />
              </SelectTrigger>
              <SelectContent>
                {components.map((c) => (
                  <SelectItem key={c.id} value={c.id}>
                    {c.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <Select value={newPair.b} onValueChange={(v) => setNewPair({ ...newPair, b: v })}>
              <SelectTrigger>
                <SelectValue placeholder="Substância B" />
              </SelectTrigger>
              <SelectContent>
                {components.map((c) => (
                  <SelectItem key={c.id} value={c.id}>
                    {c.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <Select
              value={newPair.severity}
              onValueChange={(v) => setNewPair({ ...newPair, severity: v })}
            >
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="info">Observação</SelectItem>
                <SelectItem value="warn">Desaconselhado</SelectItem>
                <SelectItem value="avoid">Não associar</SelectItem>
              </SelectContent>
            </Select>
            <Input
              placeholder="Mecanismo"
              value={newPair.mechanism}
              onChange={(e) => setNewPair({ ...newPair, mechanism: e.target.value })}
            />
            <Button
              onClick={() => createPair.mutate()}
              disabled={!newPair.a || !newPair.b || !newPair.mechanism || createPair.isPending}
            >
              <Plus className="mr-2 h-4 w-4" />
              Adicionar
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}

/**
 * Trechos do RAG que sustentam a indicação — a aula e o pedaço exato, para conferir sem reabrir
 * o material inteiro. Evidência é anexada, nunca gerada: nada aqui alimenta cálculo do sistema.
 */
function EvidencePanel({ componentId }: { componentId: string }) {
  const queryClient = useQueryClient()
  const { data: evidence = [], isLoading } = useQuery({
    queryKey: ['magistral-evidence', componentId],
    queryFn: () => getMagistralEvidence(componentId),
  })

  const pin = useMutation({
    mutationFn: ({ id, pinned }: { id: string; pinned: boolean }) =>
      pinMagistralEvidence(id, pinned),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['magistral-evidence', componentId] }),
  })

  if (isLoading) {
    return <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
  }
  if (evidence.length === 0) {
    return (
      <p className="text-xs text-muted-foreground">
        Nenhum trecho do RAG ligado a esta substância.
      </p>
    )
  }

  return (
    <div className="space-y-2">
      <Label className="flex items-center gap-2">
        <BookOpen className="h-4 w-4" />
        De onde saiu ({evidence.length} trecho{evidence.length > 1 ? 's' : ''})
      </Label>
      <div className="space-y-2">
        {evidence.map((e) => (
          <div key={e.id} className="rounded-lg border p-3">
            <div className="flex items-start justify-between gap-2">
              <div className="min-w-0">
                <p className="text-sm font-medium">{e.article?.title ?? 'Artigo'}</p>
                <p className="text-xs text-muted-foreground">
                  {e.article?.journal}
                  {e.similarity ? ` · similaridade ${Number(e.similarity).toFixed(2)}` : ''}
                </p>
              </div>
              <Button
                variant={e.pinned ? 'default' : 'ghost'}
                size="sm"
                onClick={() => pin.mutate({ id: e.id, pinned: !e.pinned })}
                title={e.pinned ? 'Soltar' : 'Fixar como o trecho que sustenta a indicação'}
              >
                <Pin className="h-4 w-4" />
              </Button>
            </div>
            <p className="mt-2 whitespace-pre-wrap text-xs text-muted-foreground">{e.excerpt}</p>
          </div>
        ))}
      </div>
    </div>
  )
}
