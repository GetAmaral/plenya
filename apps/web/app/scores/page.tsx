'use client'

import { useState, useEffect, useCallback } from 'react'
import { useRouter } from 'next/navigation'
import { Plus, Network, Search } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { useAllScoreGroupTrees } from '@/lib/api/score-api'
import { ScoreTreeView } from '@/components/scores/ScoreTreeView'
import { ScoreGroupDialog } from '@/components/scores/ScoreGroupDialog'
import { ScoreSearch, SearchResult } from '@/components/scores/ScoreSearch'

export default function ScoresPage() {
  const router = useRouter()
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [searchOpen, setSearchOpen] = useState(false)
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({})

  const { data: scoreGroups, isLoading, error } = useAllScoreGroupTrees()

  // Atalho Ctrl+F para abrir busca
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 'f') {
        e.preventDefault()
        setSearchOpen(true)
      }
    }

    window.addEventListener('keydown', handleKeyDown)
    return () => window.removeEventListener('keydown', handleKeyDown)
  }, [])

  // Expandir nodes quando clicar em resultado de busca
  const handleSearchResultClick = useCallback((result: SearchResult) => {
    const nodesToExpand: Record<string, boolean> = { ...expandedNodes }

    // Expandir grupo
    if (result.groupId) {
      nodesToExpand[`group-${result.groupId}`] = true
    }

    // Expandir subgrupo
    if (result.subgroupId) {
      nodesToExpand[`subgroup-${result.subgroupId}`] = true
    }

    // Expandir item
    if (result.itemId) {
      nodesToExpand[`item-${result.itemId}`] = true
    }

    setExpandedNodes(nodesToExpand)
    setSearchOpen(false)

    // Aguardar um pouco mais para garantir que o accordion expandiu
    setTimeout(() => {
      const element = document.getElementById(result.id)
      if (element) {
        // Scroll suave até o elemento
        element.scrollIntoView({ behavior: 'smooth', block: 'center' })

        // Highlight temporário com transição
        element.classList.add('ring-2', 'ring-primary', 'ring-offset-2', 'transition-all')

        setTimeout(() => {
          element.classList.remove('ring-2', 'ring-primary', 'ring-offset-2')
        }, 2000)
      } else {
        console.warn('Elemento não encontrado:', result.id)
      }
    }, 300)
  }, [expandedNodes])

  const handleSearchToggle = useCallback(() => {
    setSearchOpen(prev => !prev)
  }, [])

  if (error) {
    return (
      <div className="container mx-auto py-8">
        <div className="rounded-lg border border-destructive bg-destructive/10 p-4">
          <h3 className="font-semibold text-destructive">Erro ao carregar escores</h3>
          <p className="text-sm text-muted-foreground mt-1">
            {error instanceof Error ? error.message : 'Erro desconhecido'}
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="container mx-auto py-8 space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Gestão de Escores</h1>
          <p className="text-muted-foreground mt-1">
            Gerencie os critérios de estratificação de risco
          </p>
        </div>

        <div className="flex gap-2">
          <Button onClick={handleSearchToggle} variant="outline">
            <Search className="mr-2 h-4 w-4" />
            Procurar
            <kbd className="ml-2 pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium text-muted-foreground opacity-100">
              Ctrl+F
            </kbd>
          </Button>

          <Button
            variant="outline"
            onClick={() => router.push('/scores/mindmap')}
          >
            <Network className="mr-2 h-4 w-4" />
            Visualizar Mindmap
          </Button>

          <Button onClick={() => setIsCreateDialogOpen(true)}>
            <Plus className="mr-2 h-4 w-4" />
            Novo Grupo
          </Button>
        </div>
      </div>

      {/* Search Component */}
      <ScoreSearch
        scoreGroups={scoreGroups || []}
        isOpen={searchOpen}
        onClose={() => setSearchOpen(false)}
        onResultClick={handleSearchResultClick}
      />

      {/* Content */}
      {isLoading ? (
        <div className="flex items-center justify-center py-12">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary" />
        </div>
      ) : !scoreGroups || scoreGroups.length === 0 ? (
        <div className="rounded-lg border border-dashed p-12 text-center">
          <h3 className="text-lg font-semibold">Nenhum grupo de escores cadastrado</h3>
          <p className="text-sm text-muted-foreground mt-1">
            Comece criando um novo grupo de escores
          </p>
          <Button
            onClick={() => setIsCreateDialogOpen(true)}
            className="mt-4"
          >
            <Plus className="mr-2 h-4 w-4" />
            Criar Primeiro Grupo
          </Button>
        </div>
      ) : (
        <ScoreTreeView groups={scoreGroups} expandedNodes={expandedNodes} />
      )}

      {/* Create Group Dialog */}
      <ScoreGroupDialog
        open={isCreateDialogOpen}
        onOpenChange={setIsCreateDialogOpen}
      />
    </div>
  )
}
