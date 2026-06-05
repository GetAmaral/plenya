'use client'

import { useEffect, useMemo, useRef, useState } from 'react'
import Link from 'next/link'
import { useParams } from 'next/navigation'
import { toast } from 'sonner'
import { ArrowLeft, Save, Pencil, ExternalLink, Loader2 } from 'lucide-react'
import { usePatientGuard } from '@/lib/use-patient-guard'
import { useAuthStore, isGranted } from '@/lib/auth-store'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { PageHeader } from '@/components/layout/page-header'
import { ScoreVersionDialog } from '@/components/scores/ScoreVersionDialog'
import { VersionItemsEditor } from '@/components/scores/VersionItemsEditor'
import { useAllScoreItems, type ScoreItem } from '@/lib/api/score-api'
import { useScoreVersion, useSetVersionItems } from '@/lib/api/score-version-api'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001'

export default function ScoreVersionEditorPage() {
  usePatientGuard() // staff only
  const { user } = useAuthStore()
  const isAdmin = user ? isGranted(user, 'admin') : false

  const params = useParams<{ id: string }>()
  const id = params?.id ?? ''

  const { data: version, isLoading: loadingVersion, error } = useScoreVersion(id)
  const { data: allItems, isLoading: loadingItems } = useAllScoreItems()
  const setItems = useSetVersionItems()

  const [editOpen, setEditOpen] = useState(false)
  const [selected, setSelected] = useState<ScoreItem[] | null>(null)
  const [baseline, setBaseline] = useState<ScoreItem[]>([])
  const initRef = useRef<string | null>(null)

  // Inicializa a seleção UMA vez por versão (não re-clobra edições não salvas em refetches).
  // Resolve cada item pelo pool completo (allItems) — o scoreItem embutido no GET /:id não traz
  // subgroup/group preloadados, então usamos o pool (que traz) e caímos no embed só como fallback.
  useEffect(() => {
    if (!version?.id || !allItems || initRef.current === version.id) return
    initRef.current = version.id
    const byId = new Map(allItems.map((i) => [i.id, i]))
    const items = [...(version.items ?? [])]
      .sort((a, b) => a.displayOrder - b.displayOrder)
      .map((vi) => byId.get(vi.scoreItemId) ?? vi.scoreItem)
      .filter((si): si is ScoreItem => !!si)
    setSelected(items)
    setBaseline(items)
  }, [version, allItems])

  const dirty = useMemo(() => {
    if (selected === null) return false
    if (selected.length !== baseline.length) return true
    return selected.some((it, i) => it.id !== baseline[i]?.id)
  }, [selected, baseline])

  const handleSave = async () => {
    if (!selected) return
    try {
      await setItems.mutateAsync({
        id,
        items: selected.map((it, idx) => ({ scoreItemId: it.id, displayOrder: idx })),
      })
      setBaseline(selected)
      toast.success('Itens da versão salvos')
    } catch (e: any) {
      toast.error(e?.message || 'Erro ao salvar itens')
    }
  }

  const handleDiscard = () => {
    if (selected === null) return
    // Restaura fielmente a composição + ordem salvas (baseline guarda os objetos ScoreItem).
    setSelected([...baseline])
  }

  if (loadingVersion || loadingItems) {
    return (
      <div className="flex items-center justify-center py-20 text-muted-foreground">
        <Loader2 className="mr-2 h-5 w-5 animate-spin" /> Carregando…
      </div>
    )
  }

  if (error || !version) {
    return (
      <div className="space-y-4">
        <Link href="/scores/versions" className="inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground">
          <ArrowLeft className="h-4 w-4" /> Versões de Escore
        </Link>
        <div className="rounded-lg border border-destructive/30 bg-destructive/5 p-6 text-sm text-destructive">
          Versão não encontrada.
        </div>
      </div>
    )
  }

  return (
    <div className="space-y-4">
      <Link
        href="/scores/versions"
        className="inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground"
      >
        <ArrowLeft className="h-4 w-4" /> Versões de Escore
      </Link>

      <div className="flex flex-wrap items-center justify-between gap-3">
        <div className="flex items-center gap-3 min-w-0">
          <h1 className="text-xl font-semibold truncate">{version.name}</h1>
          <Badge variant="outline" className="font-mono">{version.slug}</Badge>
          {!version.active && <Badge variant="secondary">inativa</Badge>}
        </div>
        <div className="flex items-center gap-2">
          <a
            href={`${API_URL}/api/v1/score-light/config/${version.slug}`}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-1 text-sm text-muted-foreground hover:text-primary"
            title="Abre o config publicado (estado salvo) em JSON"
          >
            <ExternalLink className="h-4 w-4" /> Config publicado
          </a>
          {isAdmin && (
            <Button variant="outline" size="sm" onClick={() => setEditOpen(true)} className="gap-2">
              <Pencil className="h-4 w-4" /> Editar metadados
            </Button>
          )}
        </div>
      </div>

      <PageHeader
        title={`Versão · ${version.name}`}
        description={
          dirty
            ? 'Alterações não salvas. Clique em Salvar para aplicar.'
            : 'Selecione e ordene os itens que compõem esta versão.'
        }
        actions={
          isAdmin
            ? [
                {
                  label: 'Descartar',
                  icon: <ArrowLeft className="h-4 w-4" />,
                  variant: 'outline' as const,
                  onClick: handleDiscard,
                  disabled: !dirty || setItems.isPending,
                },
                {
                  label: setItems.isPending ? 'Salvando…' : 'Salvar',
                  icon: <Save className="h-4 w-4" />,
                  onClick: handleSave,
                  disabled: !dirty || setItems.isPending,
                  loading: setItems.isPending,
                },
              ]
            : []
        }
      />

      {!isAdmin && (
        <div className="rounded-lg border bg-muted/40 p-3 text-sm text-muted-foreground">
          Apenas administradores podem editar a composição de versões. Você está em modo leitura.
        </div>
      )}

      <VersionItemsEditor
        availableItems={allItems ?? []}
        selectedItems={selected ?? []}
        onSelectionChange={setSelected}
        disabled={!isAdmin}
      />

      <ScoreVersionDialog open={editOpen} onOpenChange={setEditOpen} version={version} />
    </div>
  )
}
