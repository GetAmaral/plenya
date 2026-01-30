'use client'

import { useState, useEffect, useCallback, startTransition } from 'react'
import { useRouter } from 'next/navigation'
import { Plus, Network, Search, Printer, FileImage, ChevronsDown, ChevronsUp, Minimize2, Loader2, Eye } from 'lucide-react'
import { useAllScoreGroupTrees } from '@/lib/api/score-api'
import { ScoreTreeView } from '@/components/scores/ScoreTreeView'
import { ScoreGroupDialog } from '@/components/scores/ScoreGroupDialog'
import { ScoreSearch, SearchResult } from '@/components/scores/ScoreSearch'
import { PageHeader } from '@/components/layout/page-header'

export default function ScoresPage() {
  const router = useRouter()
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [searchOpen, setSearchOpen] = useState(false)
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({})
  const [expandClinicalTexts, setExpandClinicalTexts] = useState(false)
  const [isExpanding, setIsExpanding] = useState(false)

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

    if (result.groupId) {
      nodesToExpand[`group-${result.groupId}`] = true
    }

    if (result.subgroupId) {
      nodesToExpand[`subgroup-${result.subgroupId}`] = true
    }

    if (result.itemId) {
      nodesToExpand[`item-${result.itemId}`] = true
    }

    setExpandedNodes(nodesToExpand)
    setSearchOpen(false)

    setTimeout(() => {
      const element = document.getElementById(result.id)
      if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'center' })
        element.classList.add('ring-2', 'ring-primary', 'ring-offset-2', 'transition-all')
        setTimeout(() => {
          element.classList.remove('ring-2', 'ring-primary', 'ring-offset-2')
        }, 2000)
      }
    }, 300)
  }, [expandedNodes])

  const handleSearchToggle = useCallback(() => {
    setSearchOpen(prev => !prev)
  }, [])

  // Expandir tudo (com textos clínicos)
  const handleExpandAll = useCallback(() => {
    if (!scoreGroups || isExpanding) return
    setIsExpanding(true)

    setTimeout(() => {
      startTransition(() => {
        const allNodes: Record<string, boolean> = {}
        scoreGroups.forEach(group => {
          allNodes[`group-${group.id}`] = true
          group.subgroups?.forEach(subgroup => {
            allNodes[`subgroup-${subgroup.id}`] = true
            subgroup.items?.forEach(item => {
              allNodes[`item-${item.id}`] = true
            })
          })
        })
        setExpandedNodes(allNodes)
        setExpandClinicalTexts(true)
      })

      setTimeout(() => {
        setIsExpanding(false)
      }, 500)
    }, 100)
  }, [scoreGroups, isExpanding])

  // Expandir tudo (sem textos clínicos)
  const handleExpandAllWithoutTexts = useCallback(() => {
    if (!scoreGroups || isExpanding) return
    setIsExpanding(true)

    setTimeout(() => {
      startTransition(() => {
        const allNodes: Record<string, boolean> = {}
        scoreGroups.forEach(group => {
          allNodes[`group-${group.id}`] = true
          group.subgroups?.forEach(subgroup => {
            allNodes[`subgroup-${subgroup.id}`] = true
            subgroup.items?.forEach(item => {
              allNodes[`item-${item.id}`] = true
            })
          })
        })
        setExpandedNodes(allNodes)
        setExpandClinicalTexts(false)
      })

      setTimeout(() => {
        setIsExpanding(false)
      }, 300)
    }, 100)
  }, [scoreGroups, isExpanding])

  // Recolher tudo
  const handleCollapseAll = useCallback(() => {
    if (isExpanding) return
    setIsExpanding(true)

    setTimeout(() => {
      startTransition(() => {
        setExpandedNodes({})
        setExpandClinicalTexts(false)
      })
      setTimeout(() => {
        setIsExpanding(false)
      }, 100)
    }, 50)
  }, [isExpanding])

  const totalGroups = scoreGroups?.length || 0
  const totalItems = scoreGroups?.reduce((acc, group) => {
    return acc + (group.subgroups?.reduce((subAcc, subgroup) => {
      return subAcc + (subgroup.items?.length || 0)
    }, 0) || 0)
  }, 0) || 0

  if (error) {
    return (
      <div className="flex items-center justify-center h-96">
        <div className="text-center">
          <p className="text-destructive">Erro ao carregar escores</p>
          <p className="text-sm text-muted-foreground mt-2">
            {error.message || 'Erro desconhecido'}
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="space-y-8">
      {/* Search Modal */}
      {searchOpen && (
        <div className="fixed inset-0 z-50 bg-black/50 flex items-start justify-center pt-20">
          <div className="w-full max-w-2xl">
            <ScoreSearch
              scoreGroups={scoreGroups || []}
              onResultClick={handleSearchResultClick}
              onClose={() => setSearchOpen(false)}
            />
          </div>
        </div>
      )}

      {/* Page Header */}
      <PageHeader
        breadcrumbs={[{ label: 'Escores' }]}
        title="Escores de Saúde"
        description={`${totalGroups} grupos · ${totalItems} critérios de avaliação`}
        actions={[
          {
            label: 'Expandir',
            icon: isExpanding ? <Loader2 className="h-4 w-4 animate-spin" /> : <ChevronsDown className="h-4 w-4" />,
            disabled: isExpanding || isLoading,
            items: [
              {
                label: 'Expandir Tudo',
                icon: <ChevronsDown className="h-4 w-4" />,
                onClick: handleExpandAll,
                disabled: isExpanding || isLoading,
              },
              {
                label: 'Sem Textos',
                icon: <ChevronsUp className="h-4 w-4" />,
                onClick: handleExpandAllWithoutTexts,
                disabled: isExpanding || isLoading,
              },
              {
                label: 'Recolher',
                icon: <Minimize2 className="h-4 w-4" />,
                onClick: handleCollapseAll,
                disabled: isExpanding || isLoading,
              },
            ],
          },
          {
            label: 'Buscar',
            icon: <Search className="h-4 w-4" />,
            onClick: handleSearchToggle,
          },
          {
            label: 'Visualizar',
            icon: <Eye className="h-4 w-4" />,
            items: [
              {
                label: 'Mindmap',
                icon: <Network className="h-4 w-4" />,
                onClick: () => router.push('/scores/mindmap'),
              },
              {
                label: 'Impressão',
                icon: <Printer className="h-4 w-4" />,
                onClick: () => router.push('/scores/print'),
              },
              {
                label: 'Pôster',
                icon: <FileImage className="h-4 w-4" />,
                onClick: () => router.push('/scores/poster'),
              },
            ],
          },
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setIsCreateDialogOpen(true),
            variant: 'default',
          },
        ]}
      />

      {/* Content */}
      {isLoading && (
        <div className="flex items-center justify-center h-64">
          <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
        </div>
      )}

      {!isLoading && scoreGroups && scoreGroups.length > 0 && (
        <ScoreTreeView
          groups={scoreGroups}
          expandedNodes={expandedNodes}
          expandClinicalTexts={expandClinicalTexts}
        />
      )}

      {!isLoading && (!scoreGroups || scoreGroups.length === 0) && (
        <div className="flex flex-col items-center justify-center h-96 text-center">
          <Network className="h-16 w-16 text-muted-foreground mb-4" />
          <h3 className="text-lg font-semibold mb-2">Nenhum escore cadastrado</h3>
          <p className="text-sm text-muted-foreground mb-4">
            Comece criando seu primeiro grupo de escores
          </p>
          <Button onClick={() => setIsCreateDialogOpen(true)} className="gap-2">
            <Plus className="h-4 w-4" />
            Criar Primeiro Grupo
          </Button>
        </div>
      )}

      {/* Create Dialog */}
      <ScoreGroupDialog
        open={isCreateDialogOpen}
        onOpenChange={setIsCreateDialogOpen}
        onSuccess={() => {
          setIsCreateDialogOpen(false)
        }}
      />
    </div>
  )
}
