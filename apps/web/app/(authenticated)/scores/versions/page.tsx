'use client'

import { useState } from 'react'
import Link from 'next/link'
import { toast } from 'sonner'
import { Plus, Layers, Pencil, Trash2, ListOrdered, Loader2 } from 'lucide-react'
import { usePatientGuard } from '@/lib/use-patient-guard'
import { useAuthStore, isGranted } from '@/lib/auth-store'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { PageHeader } from '@/components/layout/page-header'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'
import { ScoreVersionDialog } from '@/components/scores/ScoreVersionDialog'
import {
  type ScoreVersion,
  useScoreVersions,
  useDeleteScoreVersion,
} from '@/lib/api/score-version-api'

export default function ScoreVersionsPage() {
  usePatientGuard() // staff only
  const { user } = useAuthStore()
  const isAdmin = user ? isGranted(user, 'admin') : false

  const { data: versions, isLoading, error } = useScoreVersions()
  const deleteVersion = useDeleteScoreVersion()

  const [createOpen, setCreateOpen] = useState(false)
  const [editVersion, setEditVersion] = useState<ScoreVersion | undefined>(undefined)
  const [toDelete, setToDelete] = useState<ScoreVersion | undefined>(undefined)

  const handleDelete = async () => {
    if (!toDelete) return
    try {
      await deleteVersion.mutateAsync(toDelete.id)
      toast.success('Versão excluída')
      setToDelete(undefined)
    } catch (e: any) {
      toast.error(e?.message || 'Erro ao excluir versão')
    }
  }

  const sorted = [...(versions ?? [])].sort((a, b) => a.order - b.order || a.name.localeCompare(b.name))

  return (
    <div className="space-y-4">
      <PageHeader
        title="Versões de Escore"
        description={versions ? `${versions.length} versão(ões)` : undefined}
        actions={
          isAdmin
            ? [
                {
                  label: 'Nova versão',
                  icon: <Plus className="h-4 w-4" />,
                  onClick: () => setCreateOpen(true),
                },
              ]
            : []
        }
      />

      <p className="text-sm text-muted-foreground">
        Cada versão (ex: Triagem, Light) seleciona e ordena um subconjunto dos itens de escore para publicar no site.
        Os itens permanecem na fonte única; a versão só os escolhe e ordena.
      </p>

      {isLoading && (
        <div className="flex items-center justify-center py-16 text-muted-foreground">
          <Loader2 className="mr-2 h-5 w-5 animate-spin" /> Carregando…
        </div>
      )}

      {error && (
        <div className="rounded-lg border border-destructive/30 bg-destructive/5 p-6 text-sm text-destructive">
          Erro ao carregar versões.
        </div>
      )}

      {!isLoading && !error && sorted.length === 0 && (
        <div className="rounded-lg border border-dashed p-10 text-center text-sm text-muted-foreground">
          <Layers className="mx-auto mb-2 h-6 w-6 opacity-50" />
          Nenhuma versão criada ainda.
        </div>
      )}

      {sorted.length > 0 && (
        <div className="divide-y rounded-lg border">
          {sorted.map((v) => (
            <div key={v.id} className="flex flex-wrap items-center justify-between gap-3 p-4">
              <div className="flex items-center gap-3 min-w-0">
                <Layers className="h-5 w-5 shrink-0 text-muted-foreground" />
                <div className="min-w-0">
                  <div className="flex items-center gap-2">
                    <span className="font-medium truncate">{v.name}</span>
                    <Badge variant="outline" className="font-mono">{v.slug}</Badge>
                    {!v.active && <Badge variant="secondary">inativa</Badge>}
                  </div>
                  {v.description && (
                    <p className="text-xs text-muted-foreground truncate">{v.description}</p>
                  )}
                </div>
              </div>

              <div className="flex items-center gap-2 shrink-0">
                <Button asChild variant="default" size="sm" className="gap-2">
                  <Link href={`/scores/versions/${v.id}`}>
                    <ListOrdered className="h-4 w-4" /> Configurar itens
                  </Link>
                </Button>
                {isAdmin && (
                  <>
                    <Button variant="outline" size="sm" onClick={() => setEditVersion(v)} className="gap-2">
                      <Pencil className="h-4 w-4" /> Editar
                    </Button>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => setToDelete(v)}
                      className="gap-2 text-muted-foreground hover:text-destructive"
                    >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </>
                )}
              </div>
            </div>
          ))}
        </div>
      )}

      {/* Criar */}
      <ScoreVersionDialog open={createOpen} onOpenChange={setCreateOpen} />

      {/* Editar metadados */}
      <ScoreVersionDialog
        open={!!editVersion}
        onOpenChange={(o) => !o && setEditVersion(undefined)}
        version={editVersion}
      />

      {/* Excluir */}
      <AlertDialog open={!!toDelete} onOpenChange={(o) => !o && setToDelete(undefined)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Excluir versão?</AlertDialogTitle>
            <AlertDialogDescription>
              A versão <strong>{toDelete?.name}</strong> ({toDelete?.slug}) será removida. Os itens de escore
              em si não são afetados, apenas o vínculo desta versão. Esta ação não pode ser desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel disabled={deleteVersion.isPending}>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={(e) => {
                e.preventDefault()
                handleDelete()
              }}
              disabled={deleteVersion.isPending}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            >
              {deleteVersion.isPending ? 'Excluindo…' : 'Excluir'}
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}
