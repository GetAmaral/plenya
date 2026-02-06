import { useCallback, useMemo } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { useAuthStore } from '@/lib/auth-store'
import { userApi } from '@/lib/api/user-api'

export type LabResultsOrderByType = 'alpha-asc' | 'alpha-desc' | 'view'

export interface LabResultsOrderByPreference {
  type: LabResultsOrderByType
  viewId?: string
}

/**
 * Hook para gerenciar a preferência de ordenação de resultados laboratoriais
 */
export function useLabResultViewPreference() {
  const { user, updateUser } = useAuthStore()
  const queryClient = useQueryClient()

  // Obter preferência salva (default: alpha-asc)
  const currentOrderBy = useMemo<LabResultsOrderByPreference>(() => {
    return (user?.preferences as any)?.labResultsOrderBy || { type: 'alpha-asc' }
  }, [user?.preferences])

  // Salvar preferência (com debounce embutido na API)
  const setOrderBy = useCallback(
    async (orderBy: LabResultsOrderByPreference) => {
      try {
        const newPreferences = {
          ...(user?.preferences || {}),
          labResultsOrderBy: orderBy,
        }

        const updatedUser = await userApi.updatePreferences(newPreferences)

        // Atualizar Zustand store
        updateUser({ ...updatedUser })

        // CRÍTICO: Atualizar cache da query para evitar sobrescrita pelo useEffect
        queryClient.setQueryData(['user', 'me'], updatedUser)
      } catch (error) {
        console.error('Failed to save order by preference:', error)
      }
    },
    [user, updateUser, queryClient]
  )

  return { currentOrderBy, setOrderBy }
}
