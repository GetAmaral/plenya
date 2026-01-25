'use client'

import { useState, useMemo, useCallback, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import ReactFlow, {
  Background,
  MiniMap,
  Node,
  Edge,
  useNodesState,
  useEdgesState,
  useReactFlow,
  BackgroundVariant,
  ReactFlowProvider,
} from 'reactflow'
import 'reactflow/dist/style.css'
import {
  ArrowLeft,
  Download,
  ZoomIn,
  ZoomOut,
  Maximize2,
  Minimize2,
  ChevronDown,
  ChevronRight,
  Search,
} from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  useAllScoreGroupTrees,
  useDeleteScoreGroup,
  useDeleteScoreSubgroup,
  useDeleteScoreItem,
  useDeleteScoreLevel,
  ScoreGroup,
  ScoreSubgroup,
  ScoreItem,
  ScoreLevel
} from '@/lib/api/score-api'
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
import { toast } from 'sonner'
import { GroupNode } from '@/components/scores/mindmap/GroupNode'
import { SubgroupNode } from '@/components/scores/mindmap/SubgroupNode'
import { ItemNode } from '@/components/scores/mindmap/ItemNode'
import { LevelNode } from '@/components/scores/mindmap/LevelNode'
import { MindmapSearch } from '@/components/scores/mindmap/MindmapSearch'
import { ScoreGroupDialog } from '@/components/scores/ScoreGroupDialog'
import { ScoreSubgroupDialog } from '@/components/scores/ScoreSubgroupDialog'
import { ScoreItemDialog } from '@/components/scores/ScoreItemDialog'
import { ScoreLevelDialog } from '@/components/scores/ScoreLevelDialog'
import { buildMindmapLayout } from '@/components/scores/mindmap/useMindmapLayout'
import { exportMindmapToPNG } from '@/components/scores/mindmap/exportMindmap'

const nodeTypes = {
  groupNode: GroupNode,
  subgroupNode: SubgroupNode,
  itemNode: ItemNode,
  levelNode: LevelNode,
}

function MindmapContent() {
  const router = useRouter()
  const { data: scoreGroups, isLoading, error } = useAllScoreGroupTrees()
  const { zoomIn, zoomOut, fitView, setViewport, getZoom, getNodes } = useReactFlow()

  // Estado de expansão: { nodeId: boolean }
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({})
  const [zoomLevel, setZoomLevel] = useState(0.6)
  const [initialized, setInitialized] = useState(false)
  const [searchOpen, setSearchOpen] = useState(false)

  // Estado para dialogs de edição
  const [editingGroup, setEditingGroup] = useState<ScoreGroup | null>(null)
  const [editingSubgroup, setEditingSubgroup] = useState<{ subgroup: ScoreSubgroup; groupId: string } | null>(null)
  const [editingItem, setEditingItem] = useState<{ item: ScoreItem; subgroupId: string } | null>(null)
  const [editingLevel, setEditingLevel] = useState<{ level: ScoreLevel; itemId: string } | null>(null)

  // Estado para dialogs de criação
  const [creatingSubgroup, setCreatingSubgroup] = useState<{ groupId: string } | null>(null)
  const [creatingItem, setCreatingItem] = useState<{ subgroupId: string } | null>(null)
  const [creatingLevel, setCreatingLevel] = useState<{ itemId: string } | null>(null)

  // Estado para confirmação de exclusão
  const [deletingGroup, setDeletingGroup] = useState<ScoreGroup | null>(null)
  const [deletingSubgroup, setDeletingSubgroup] = useState<ScoreSubgroup | null>(null)
  const [deletingItem, setDeletingItem] = useState<ScoreItem | null>(null)
  const [deletingLevel, setDeletingLevel] = useState<ScoreLevel | null>(null)

  // Inicializar com todos os grupos expandidos quando carregar
  useEffect(() => {
    if (scoreGroups && scoreGroups.length > 0 && !initialized) {
      const initial: Record<string, boolean> = {}
      scoreGroups.forEach(group => {
        initial[`group-${group.id}`] = true // Grupos expandidos por padrão
      })
      setExpandedNodes(initial)
      setInitialized(true)
    }
  }, [scoreGroups, initialized])

  // Toggle expansão de um node
  const toggleNode = useCallback((nodeId: string) => {
    setExpandedNodes(prev => ({
      ...prev,
      [nodeId]: !prev[nodeId]
    }))
  }, [])

  // Handlers para edição via context menu
  const handleEditGroup = useCallback((group: ScoreGroup) => {
    setEditingGroup(group)
  }, [])

  const handleEditSubgroup = useCallback((subgroup: ScoreSubgroup, groupId: string) => {
    setEditingSubgroup({ subgroup, groupId })
  }, [])

  const handleEditItem = useCallback((item: ScoreItem, subgroupId: string) => {
    setEditingItem({ item, subgroupId })
  }, [])

  const handleEditLevel = useCallback((level: ScoreLevel, itemId: string) => {
    setEditingLevel({ level, itemId })
  }, [])

  // Handlers para criação via context menu
  const handleAddSubgroup = useCallback((groupId: string) => {
    setCreatingSubgroup({ groupId })
  }, [])

  const handleAddItem = useCallback((subgroupId: string) => {
    setCreatingItem({ subgroupId })
  }, [])

  const handleAddLevel = useCallback((itemId: string) => {
    setCreatingLevel({ itemId })
  }, [])

  // Handlers para exclusão via context menu
  const handleDeleteGroup = useCallback((group: ScoreGroup) => {
    setDeletingGroup(group)
  }, [])

  const handleDeleteSubgroup = useCallback((subgroup: ScoreSubgroup) => {
    setDeletingSubgroup(subgroup)
  }, [])

  const handleDeleteItem = useCallback((item: ScoreItem) => {
    setDeletingItem(item)
  }, [])

  const handleDeleteLevel = useCallback((level: ScoreLevel) => {
    setDeletingLevel(level)
  }, [])

  // Delete mutations
  const deleteGroup = useDeleteScoreGroup()
  const deleteSubgroup = useDeleteScoreSubgroup()
  const deleteItem = useDeleteScoreItem()
  const deleteLevel = useDeleteScoreLevel()

  // Confirmar exclusões
  const confirmDeleteGroup = useCallback(async () => {
    if (!deletingGroup) return
    try {
      await deleteGroup.mutateAsync(deletingGroup.id)
      toast.success('Grupo excluído com sucesso')
      setDeletingGroup(null)
    } catch (error) {
      toast.error('Erro ao excluir grupo')
    }
  }, [deletingGroup, deleteGroup])

  const confirmDeleteSubgroup = useCallback(async () => {
    if (!deletingSubgroup) return
    try {
      await deleteSubgroup.mutateAsync(deletingSubgroup.id)
      toast.success('Subgrupo excluído com sucesso')
      setDeletingSubgroup(null)
    } catch (error) {
      toast.error('Erro ao excluir subgrupo')
    }
  }, [deletingSubgroup, deleteSubgroup])

  const confirmDeleteItem = useCallback(async () => {
    if (!deletingItem) return
    try {
      await deleteItem.mutateAsync(deletingItem.id)
      toast.success('Item excluído com sucesso')
      setDeletingItem(null)
    } catch (error) {
      toast.error('Erro ao excluir item')
    }
  }, [deletingItem, deleteItem])

  const confirmDeleteLevel = useCallback(async () => {
    if (!deletingLevel) return
    try {
      await deleteLevel.mutateAsync(deletingLevel.id)
      toast.success('Nível excluído com sucesso')
      setDeletingLevel(null)
    } catch (error) {
      toast.error('Erro ao excluir nível')
    }
  }, [deletingLevel, deleteLevel])

  // Build nodes and edges com estado de expansão
  const { nodes: calculatedNodes, edges: calculatedEdges} = useMemo(() => {
    if (!scoreGroups || scoreGroups.length === 0) {
      return { nodes: [], edges: [] }
    }
    const layout = buildMindmapLayout(scoreGroups, expandedNodes, zoomLevel)

    // Adicionar callbacks aos nodes
    return {
      nodes: layout.nodes.map(node => {
        const baseData = {
          ...node.data,
          onToggle: () => toggleNode(node.id),
          isExpanded: expandedNodes[node.id] || false,
        }

        // Adicionar callbacks baseado no tipo
        if (node.type === 'groupNode') {
          return {
            ...node,
            data: {
              ...baseData,
              onEdit: () => handleEditGroup(node.data as ScoreGroup),
              onAdd: () => handleAddSubgroup(node.data.id),
              onDelete: () => handleDeleteGroup(node.data as ScoreGroup),
            }
          }
        } else if (node.type === 'subgroupNode') {
          return {
            ...node,
            data: {
              ...baseData,
              onEdit: (groupId: string) => handleEditSubgroup(node.data as ScoreSubgroup, groupId),
              onAdd: () => handleAddItem(node.data.id),
              onDelete: () => handleDeleteSubgroup(node.data as ScoreSubgroup),
            }
          }
        } else if (node.type === 'itemNode') {
          return {
            ...node,
            data: {
              ...baseData,
              onEdit: (subgroupId: string) => handleEditItem(node.data as ScoreItem, subgroupId),
              onAdd: () => handleAddLevel(node.data.id),
              onDelete: () => handleDeleteItem(node.data as ScoreItem),
            }
          }
        } else if (node.type === 'levelNode') {
          return {
            ...node,
            data: {
              ...baseData,
              onEdit: (itemId: string) => handleEditLevel(node.data as ScoreLevel, itemId),
              onDelete: () => handleDeleteLevel(node.data as ScoreLevel),
            }
          }
        }

        return { ...node, data: baseData }
      }),
      edges: layout.edges
    }
  }, [
    scoreGroups,
    expandedNodes,
    zoomLevel,
    toggleNode,
    handleEditGroup,
    handleEditSubgroup,
    handleEditItem,
    handleEditLevel,
    handleAddSubgroup,
    handleAddItem,
    handleAddLevel,
    handleDeleteGroup,
    handleDeleteSubgroup,
    handleDeleteItem,
    handleDeleteLevel,
  ])

  const [nodes, setNodes, onNodesChange] = useNodesState(calculatedNodes)
  const [edges, setEdges, onEdgesChange] = useEdgesState(calculatedEdges)

  // Atualizar nodes quando mudar expansão ou zoom
  useEffect(() => {
    setNodes(calculatedNodes)
    setEdges(calculatedEdges)
  }, [calculatedNodes, calculatedEdges, setNodes, setEdges])

  // Monitorar zoom level
  useEffect(() => {
    const interval = setInterval(() => {
      const zoom = getZoom()
      setZoomLevel(zoom)
    }, 100)
    return () => clearInterval(interval)
  }, [getZoom])

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

  // Expandir/colapsar tudo
  const handleExpandAll = useCallback(() => {
    if (!scoreGroups) return
    const allExpanded: Record<string, boolean> = {}
    scoreGroups.forEach(group => {
      allExpanded[`group-${group.id}`] = true
      group.subgroups?.forEach(subgroup => {
        allExpanded[`subgroup-${subgroup.id}`] = true
        subgroup.items?.forEach(item => {
          allExpanded[`item-${item.id}`] = true
        })
      })
    })
    setExpandedNodes(allExpanded)
  }, [scoreGroups])

  const handleCollapseAll = useCallback(() => {
    setExpandedNodes({})
  }, [])

  const handleExport = useCallback(() => {
    exportMindmapToPNG()
  }, [])

  const handleFitView = useCallback(() => {
    fitView({ padding: 0.2, duration: 300 })
  }, [fitView])

  const handleZoomIn = useCallback(() => {
    zoomIn({ duration: 300 })
  }, [zoomIn])

  const handleZoomOut = useCallback(() => {
    zoomOut({ duration: 300 })
  }, [zoomOut])

  const handleResetView = useCallback(() => {
    setViewport({ x: 0, y: 0, zoom: 0.6 }, { duration: 300 })
  }, [setViewport])

  // Expandir e focar em resultado de busca
  const handleSearchResultClick = useCallback((result: any) => {
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

    // Aguardar renderização e então centralizar no node
    setTimeout(() => {
      const nodes = getNodes()
      const targetNode = nodes.find(n => n.id === result.id)

      if (targetNode) {
        // Centralizar viewport no node
        const x = -targetNode.position.x + window.innerWidth / 2 - 168 // 168 = metade da largura do card
        const y = -targetNode.position.y + window.innerHeight / 2 - 50  // 50 = metade da altura aproximada

        setViewport({ x, y, zoom: 1.0 }, { duration: 500 })
      }

      // Fechar busca
      setSearchOpen(false)
    }, 100)
  }, [expandedNodes, getNodes, setViewport])

  const handleSearchToggle = useCallback(() => {
    setSearchOpen(prev => !prev)
  }, [])

  if (isLoading) {
    return (
      <div className="h-screen flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary mx-auto mb-4" />
          <p className="text-muted-foreground">Carregando mindmap...</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="h-screen flex items-center justify-center">
        <div className="text-center max-w-md">
          <h3 className="text-lg font-semibold text-destructive mb-2">
            Erro ao carregar dados
          </h3>
          <p className="text-sm text-muted-foreground mb-4">
            {error instanceof Error ? error.message : 'Erro desconhecido'}
          </p>
          <Button onClick={() => router.push('/scores')}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            Voltar
          </Button>
        </div>
      </div>
    )
  }

  if (!scoreGroups || scoreGroups.length === 0) {
    return (
      <div className="h-screen flex items-center justify-center">
        <div className="text-center max-w-md">
          <h3 className="text-lg font-semibold mb-2">Nenhum dado encontrado</h3>
          <p className="text-sm text-muted-foreground mb-4">
            Não há grupos de escores cadastrados para visualizar
          </p>
          <Button onClick={() => router.push('/scores')}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            Voltar para Gestão
          </Button>
        </div>
      </div>
    )
  }

  return (
    <div className="h-screen flex flex-col">
      {/* Header */}
      <div className="border-b bg-background px-6 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => router.push('/scores')}
            >
              <ArrowLeft className="mr-2 h-4 w-4" />
              Voltar
            </Button>
            <div>
              <h1 className="text-2xl font-bold">Mindmap de Escores</h1>
              <p className="text-sm text-muted-foreground">
                Visualização hierárquica de {scoreGroups.length} grupos · Zoom: {Math.round(zoomLevel * 100)}%
              </p>
            </div>
          </div>

          <div className="flex gap-2">
            <Button onClick={handleSearchToggle} variant="outline" size="sm">
              <Search className="mr-2 h-4 w-4" />
              Procurar
              <kbd className="ml-2 pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium text-muted-foreground opacity-100">
                Ctrl+F
              </kbd>
            </Button>
            <Button onClick={handleExpandAll} variant="outline" size="sm">
              <ChevronDown className="mr-2 h-4 w-4" />
              Expandir Tudo
            </Button>
            <Button onClick={handleCollapseAll} variant="outline" size="sm">
              <ChevronRight className="mr-2 h-4 w-4" />
              Recolher Tudo
            </Button>
            <Button onClick={handleExport} variant="outline" size="sm">
              <Download className="mr-2 h-4 w-4" />
              Exportar PNG
            </Button>
          </div>
        </div>
      </div>

      {/* Search Box */}
      <MindmapSearch
        scoreGroups={scoreGroups || []}
        isOpen={searchOpen}
        onClose={() => setSearchOpen(false)}
        onResultClick={handleSearchResultClick}
      />

      {/* React Flow Canvas */}
      <div className="flex-1 relative">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          fitView
          fitViewOptions={{
            padding: 0.2,
            minZoom: 0.5,  // Zoom mínimo do fitView = 50%
            maxZoom: 1.0,  // Zoom máximo do fitView = 100%
          }}
          minZoom={0.1}
          maxZoom={2.5}
          defaultViewport={{ x: 0, y: 0, zoom: 0.7 }}
          defaultEdgeOptions={{
            type: 'smoothstep',
            animated: true,
            style: { strokeWidth: 2 },
          }}
        >
          <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
          <MiniMap
            nodeStrokeWidth={3}
            zoomable
            pannable
            className="bg-background border"
          />
        </ReactFlow>

        {/* Custom Controls Panel */}
        <div className="absolute top-4 right-4 flex flex-col gap-2">
          <Button
            variant="secondary"
            size="icon"
            onClick={handleZoomIn}
            title="Aumentar zoom"
          >
            <ZoomIn className="h-4 w-4" />
          </Button>
          <Button
            variant="secondary"
            size="icon"
            onClick={handleZoomOut}
            title="Diminuir zoom"
          >
            <ZoomOut className="h-4 w-4" />
          </Button>
          <Button
            variant="secondary"
            size="icon"
            onClick={handleFitView}
            title="Ajustar visualização"
          >
            <Maximize2 className="h-4 w-4" />
          </Button>
          <Button
            variant="secondary"
            size="icon"
            onClick={handleResetView}
            title="Reposicionar"
          >
            <Minimize2 className="h-4 w-4" />
          </Button>
        </div>

        {/* Zoom indicator */}
        <div className="absolute bottom-4 left-4 bg-background/90 backdrop-blur border rounded-lg px-3 py-2">
          <div className="text-xs text-muted-foreground">
            {zoomLevel < 0.25 && '🔭 Visão Geral - Grupos'}
            {zoomLevel >= 0.25 && zoomLevel < 0.6 && '📊 Visão Média - Grupos + Subgrupos'}
            {zoomLevel >= 0.6 && zoomLevel < 1.0 && '🔍 Visão Detalhada - Items'}
            {zoomLevel >= 1.0 && '🎯 Visão Completa - Todos Níveis'}
          </div>
        </div>
      </div>

      {/* Edit Dialogs */}
      {editingGroup && (
        <ScoreGroupDialog
          open={!!editingGroup}
          onOpenChange={(open) => !open && setEditingGroup(null)}
          group={editingGroup}
        />
      )}

      {editingSubgroup && (
        <ScoreSubgroupDialog
          open={!!editingSubgroup}
          onOpenChange={(open) => !open && setEditingSubgroup(null)}
          groupId={editingSubgroup.groupId}
          subgroup={editingSubgroup.subgroup}
        />
      )}

      {editingItem && (
        <ScoreItemDialog
          open={!!editingItem}
          onOpenChange={(open) => !open && setEditingItem(null)}
          subgroupId={editingItem.subgroupId}
          item={editingItem.item}
        />
      )}

      {editingLevel && (
        <ScoreLevelDialog
          open={!!editingLevel}
          onOpenChange={(open) => !open && setEditingLevel(null)}
          itemId={editingLevel.itemId}
          level={editingLevel.level}
        />
      )}

      {/* Create Dialogs */}
      {creatingSubgroup && (
        <ScoreSubgroupDialog
          open={!!creatingSubgroup}
          onOpenChange={(open) => !open && setCreatingSubgroup(null)}
          groupId={creatingSubgroup.groupId}
        />
      )}

      {creatingItem && (
        <ScoreItemDialog
          open={!!creatingItem}
          onOpenChange={(open) => !open && setCreatingItem(null)}
          subgroupId={creatingItem.subgroupId}
        />
      )}

      {creatingLevel && (
        <ScoreLevelDialog
          open={!!creatingLevel}
          onOpenChange={(open) => !open && setCreatingLevel(null)}
          itemId={creatingLevel.itemId}
        />
      )}

      {/* Delete Confirmation Dialogs */}
      <AlertDialog open={!!deletingGroup} onOpenChange={(open) => !open && setDeletingGroup(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar Exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir o grupo "{deletingGroup?.name}"?
              {deletingGroup?.subgroups && deletingGroup.subgroups.length > 0 && (
                <span className="block mt-2 text-destructive font-medium">
                  Atenção: Este grupo contém {deletingGroup.subgroups.length} subgrupo(s) que também serão excluídos.
                </span>
              )}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={confirmDeleteGroup} className="bg-destructive text-destructive-foreground hover:bg-destructive/90">
              Excluir
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={!!deletingSubgroup} onOpenChange={(open) => !open && setDeletingSubgroup(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar Exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir o subgrupo "{deletingSubgroup?.name}"?
              {deletingSubgroup?.items && deletingSubgroup.items.length > 0 && (
                <span className="block mt-2 text-destructive font-medium">
                  Atenção: Este subgrupo contém {deletingSubgroup.items.length} item(s) que também serão excluídos.
                </span>
              )}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={confirmDeleteSubgroup} className="bg-destructive text-destructive-foreground hover:bg-destructive/90">
              Excluir
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={!!deletingItem} onOpenChange={(open) => !open && setDeletingItem(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar Exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir o item "{deletingItem?.name}"?
              {deletingItem?.levels && deletingItem.levels.length > 0 && (
                <span className="block mt-2 text-destructive font-medium">
                  Atenção: Este item contém {deletingItem.levels.length} nível(is) que também serão excluídos.
                </span>
              )}
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={confirmDeleteItem} className="bg-destructive text-destructive-foreground hover:bg-destructive/90">
              Excluir
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={!!deletingLevel} onOpenChange={(open) => !open && setDeletingLevel(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar Exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir o nível {deletingLevel?.level} - "{deletingLevel?.name}"?
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={confirmDeleteLevel} className="bg-destructive text-destructive-foreground hover:bg-destructive/90">
              Excluir
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}

export default function MindmapPage() {
  return (
    <ReactFlowProvider>
      <MindmapContent />
    </ReactFlowProvider>
  )
}
