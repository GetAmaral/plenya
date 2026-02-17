'use client'

import { useState, useEffect, useCallback, startTransition } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import { usePatientGuard } from '@/lib/use-patient-guard'
import { Plus, Network, Search, Printer, FileImage, ChevronsDown, ChevronsUp, Minimize2, Loader2, Eye } from 'lucide-react'
import { useAllScoreGroupTrees } from '@/lib/api/score-api'
import { ScoreTreeView } from '@/components/scores/ScoreTreeView'
import { ScoreGroupDialog } from '@/components/scores/ScoreGroupDialog'
import { ScoreSearch, SearchResult } from '@/components/scores/ScoreSearch'
import { PageHeader } from '@/components/layout/page-header'
import { useAuthStore } from '@/lib/auth-store'

export default function ScoresPage() {
  usePatientGuard(); // Restrict access to staff only
  const router = useRouter()
  const searchParams = useSearchParams()
  const { accessToken } = useAuthStore()
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [searchOpen, setSearchOpen] = useState(false)
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({})
  const [openSubgroups, setOpenSubgroups] = useState<string[]>([]) // Para abrir accordions na busca
  const [expandClinicalTexts, setExpandClinicalTexts] = useState(false)
  const [isExpanding, setIsExpanding] = useState(false)
  const [isDownloadingPoster, setIsDownloadingPoster] = useState(false)
  const [isMounted, setIsMounted] = useState(false)

  const { data: scoreGroups, isLoading, error } = useAllScoreGroupTrees()

  // Prevenir hydration errors - montar apenas no cliente
  useEffect(() => {
    setIsMounted(true)
  }, [])

  // Auto-focus em item quando vem da URL (?item=xxx)
  useEffect(() => {
    const itemId = searchParams.get('item')
    if (!itemId || !scoreGroups || !isMounted) return

    // Encontrar o item nos dados
    let foundItem: any = null
    let foundSubgroup: any = null
    let foundGroup: any = null

    for (const group of scoreGroups) {
      for (const subgroup of group.subgroups || []) {
        const item = subgroup.items?.find((i: any) => i.id === itemId)
        if (item) {
          foundItem = item
          foundSubgroup = subgroup
          foundGroup = group
          break
        }
      }
      if (foundItem) break
    }

    if (foundItem && foundSubgroup && foundGroup) {
      // Expandir todos os nodes necessários
      const nodesToExpand: Record<string, boolean> = {
        [`group-${foundGroup.id}`]: true,
        [`subgroup-${foundSubgroup.id}`]: true,
        [`item-${foundItem.id}`]: true,
      }

      setExpandedNodes(nodesToExpand)
      setOpenSubgroups([foundSubgroup.id])

      // Scroll para o elemento
      requestAnimationFrame(() => {
        requestAnimationFrame(() => {
          const tryScrollToElement = (attempts = 0, maxAttempts = 15) => {
            const element = document.getElementById(`item-${itemId}`)

            if (element) {
              const rect = element.getBoundingClientRect()
              const isVisible = rect.height > 0 && rect.width > 0

              if (isVisible) {
                element.scrollIntoView({ behavior: 'smooth', block: 'center' })
                element.classList.add('ring-2', 'ring-primary', 'ring-offset-2', 'transition-all')
                setTimeout(() => {
                  element.classList.remove('ring-2', 'ring-primary', 'ring-offset-2')
                }, 3000)
              } else if (attempts < maxAttempts) {
                setTimeout(() => tryScrollToElement(attempts + 1, maxAttempts), 100)
              }
            } else if (attempts < maxAttempts) {
              setTimeout(() => tryScrollToElement(attempts + 1, maxAttempts), 100)
            }
          }

          setTimeout(() => tryScrollToElement(), 300)
        })
      })
    }
  }, [searchParams, scoreGroups, isMounted])

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
      // Também abrir o accordion do subgrupo
      setOpenSubgroups(prev => [...prev, result.subgroupId!])
    }

    if (result.itemId) {
      nodesToExpand[`item-${result.itemId}`] = true
    }

    setExpandedNodes(nodesToExpand)
    setSearchOpen(false)

    // Aguardar React processar as mudanças de estado ANTES de iniciar scroll
    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        // Agora o React já processou os estados e começou a renderizar

        // Função helper para fazer scroll e destacar elemento
        const scrollToElement = (el: HTMLElement) => {
          el.scrollIntoView({ behavior: 'smooth', block: 'center' })
          el.classList.add('ring-2', 'ring-primary', 'ring-offset-2', 'transition-all')
          setTimeout(() => {
            el.classList.remove('ring-2', 'ring-primary', 'ring-offset-2')
          }, 2000)
        }

        // Aguardar abertura do accordion com retry
        const tryScrollToElement = (attempts = 0, maxAttempts = 15) => {
          const element = document.getElementById(result.id)

          if (element) {
            // Verificar se o elemento está realmente visível (accordion aberto)
            const rect = element.getBoundingClientRect()
            const isVisible = rect.height > 0 && rect.width > 0

            if (isVisible) {
              scrollToElement(element)
            } else if (attempts < maxAttempts) {
              // Elemento existe mas ainda não está visível (accordion abrindo)
              setTimeout(() => tryScrollToElement(attempts + 1, maxAttempts), 100)
            }
          } else if (attempts < maxAttempts) {
            // Elemento ainda não existe no DOM
            setTimeout(() => tryScrollToElement(attempts + 1, maxAttempts), 100)
          }
        }

        // Pequeno delay adicional para garantir que accordion começou a abrir
        setTimeout(() => tryScrollToElement(), 300)
      })
    })
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

  // Download Poster PDF
  const handleDownloadPoster = useCallback(async () => {
    if (!accessToken || isDownloadingPoster) return

    setIsDownloadingPoster(true)
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/score-groups/poster-pdf`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${accessToken}`,
        },
      })

      if (!response.ok) {
        // Tenta ler a mensagem de erro do backend
        let errorMessage = 'Erro ao gerar PDF'
        try {
          const errorData = await response.json()
          errorMessage = errorData.error || errorMessage
        } catch {
          errorMessage = `Erro ${response.status}: ${response.statusText}`
        }
        throw new Error(errorMessage)
      }

      const blob = await response.blob()
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = 'escore-plenya-poster.pdf'
      document.body.appendChild(a)
      a.click()
      window.URL.revokeObjectURL(url)
      document.body.removeChild(a)
    } catch (error) {
      console.error('Erro ao baixar PDF:', error)
      const errorMessage = error instanceof Error ? error.message : 'Erro desconhecido'
      alert(`Erro ao gerar PDF:\n\n${errorMessage}`)
    } finally {
      setIsDownloadingPoster(false)
    }
  }, [accessToken, isDownloadingPoster])

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
    <div className="space-y-4">
      {/* Loading Overlay durante geração de PDF */}
      {isDownloadingPoster && (
        <div className="fixed inset-0 z-[100] bg-background/95 backdrop-blur-sm flex items-center justify-center">
          <div className="flex flex-col items-center gap-4">
            <Loader2 className="h-16 w-16 animate-spin text-primary" />
            <div className="text-center">
              <p className="text-xl font-semibold">Gerando PDF do Pôster</p>
              <p className="text-sm text-muted-foreground mt-2">
                Isso pode levar alguns segundos...
              </p>
            </div>
          </div>
        </div>
      )}

      {/* Search Modal */}
      {searchOpen && (
        <div className="fixed inset-0 z-50 bg-black/50 flex items-start justify-center pt-20">
          <div className="w-full max-w-2xl">
            <ScoreSearch
              scoreGroups={scoreGroups || []}
              isOpen={searchOpen}
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
        actions={isMounted ? [
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
                label: isDownloadingPoster ? 'Gerando...' : 'Pôster',
                icon: isDownloadingPoster ? <Loader2 className="h-4 w-4 animate-spin" /> : <FileImage className="h-4 w-4" />,
                onClick: handleDownloadPoster,
                disabled: isDownloadingPoster,
              },
            ],
          },
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setIsCreateDialogOpen(true),
            variant: 'default',
          },
        ] : []}
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
          openSubgroups={openSubgroups}
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
