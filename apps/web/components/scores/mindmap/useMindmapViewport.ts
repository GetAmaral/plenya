import { useCallback, useMemo } from 'react'
import { Viewport } from 'reactflow'
import { useAuthStore } from '@/lib/auth-store'
import { userApi } from '@/lib/api/user-api'

const MINDMAP_VIEWPORT_KEY = 'mindmapViewport'
const MINDMAP_EXPANDED_NODES_KEY = 'mindmapExpandedNodes'

// Simple debounce implementation
function debounce<T extends (...args: any[]) => any>(
  func: T,
  wait: number
): (...args: Parameters<T>) => void {
  let timeout: NodeJS.Timeout | null = null
  return function (...args: Parameters<T>) {
    if (timeout) clearTimeout(timeout)
    timeout = setTimeout(() => func(...args), wait)
  }
}

interface MindmapViewportPreferences {
  x: number
  y: number
  zoom: number
}

/**
 * Hook para gerenciar automaticamente o salvamento e carregamento
 * do viewport e estado dos nodes do mindmap nas preferências do usuário
 */
export function useMindmapViewport() {
  const { user, updateUser, accessToken } = useAuthStore()

  // Obter viewport salvo do usuário
  const initialViewport = useMemo(() => {
    if (!user) return null

    const savedViewport = user.preferences?.[MINDMAP_VIEWPORT_KEY] as MindmapViewportPreferences | undefined

    if (savedViewport && savedViewport.x !== undefined && savedViewport.y !== undefined && savedViewport.zoom !== undefined) {
      return savedViewport
    }

    return null
  }, [user])

  // Obter nodes expandidos salvos do usuário
  const initialExpandedNodes = useMemo(() => {
    if (!user) return {}

    const savedExpandedNodes = user.preferences?.[MINDMAP_EXPANDED_NODES_KEY] as string[] | undefined

    if (savedExpandedNodes && Array.isArray(savedExpandedNodes)) {
      // Converter array para Record<string, boolean>
      return savedExpandedNodes.reduce((acc, nodeId) => {
        acc[nodeId] = true
        return acc
      }, {} as Record<string, boolean>)
    }

    return {}
  }, [user])

  // Salvar viewport e expanded nodes com debounce de 1 segundo
  const saveState = useCallback(
    debounce(async (viewport?: Viewport, expandedNodes?: Record<string, boolean>) => {
      if (!user || !accessToken) {
        return
      }

      try {
        const newPreferences = {
          ...(user.preferences || {}),
        }

        // Atualizar viewport se fornecido
        if (viewport) {
          newPreferences[MINDMAP_VIEWPORT_KEY] = {
            x: viewport.x,
            y: viewport.y,
            zoom: viewport.zoom,
          }
        }

        // Atualizar expanded nodes se fornecido
        if (expandedNodes) {
          // Converter Record<string, boolean> para array apenas dos expandidos
          const expandedNodeIds = Object.keys(expandedNodes).filter(nodeId => expandedNodes[nodeId])
          newPreferences[MINDMAP_EXPANDED_NODES_KEY] = expandedNodeIds
        }

        // Atualizar no backend
        const updatedUser = await userApi.updatePreferences(newPreferences)

        // Atualizar estado local
        updateUser(updatedUser)
      } catch (error) {
        console.error('[MINDMAP] Failed to save state:', error)
        // Silenciar erro para não interromper a experiência do usuário
      }
    }, 1000),
    [user, accessToken, updateUser]
  )

  // Funções para salvar estado
  const onViewportChange = useCallback((viewport: Viewport) => {
    saveState(viewport, undefined)
  }, [saveState])

  const onExpandedNodesChange = useCallback((expandedNodes: Record<string, boolean>) => {
    saveState(undefined, expandedNodes)
  }, [saveState])

  return {
    onViewportChange,
    onExpandedNodesChange,
    initialViewport,
    initialExpandedNodes,
  }
}
