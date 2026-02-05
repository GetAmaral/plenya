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
    const orderBy = (user?.preferences as any)?.labResultsOrderBy || { type: 'alpha-asc' }
    console.log('[useLabResultViewPreference] currentOrderBy recalculated:', orderBy, 'user.preferences:', user?.preferences)
    return orderBy
  }, [user?.preferences])

  // Salvar preferência (com debounce embutido na API)
  const setOrderBy = useCallback(
    async (orderBy: LabResultsOrderByPreference) => {
      console.log('[useLabResultViewPreference] setOrderBy called with:', orderBy)
      try {
        const newPreferences = {
          ...(user?.preferences || {}),
          labResultsOrderBy: orderBy,
        }

        console.log('[useLabResultViewPreference] Saving preferences:', newPreferences)
        const updatedUser = await userApi.updatePreferences(newPreferences)
        console.log('[useLabResultViewPreference] Preferences saved, updating user:', updatedUser)

        // Atualizar Zustand store
        updateUser({ ...updatedUser })

        // CRÍTICO: Atualizar cache da query para evitar sobrescrita pelo useEffect
        queryClient.setQueryData(['user', 'me'], updatedUser)
      } catch (error) {
        console.error('[useLabResultViewPreference] Failed to save order by preference:', error)
      }
    },
    [user, updateUser, queryClient]
  )

  return { currentOrderBy, setOrderBy }
}
