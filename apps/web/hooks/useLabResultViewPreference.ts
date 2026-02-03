import { useCallback, useMemo } from 'react'
import { useAuthStore } from '@/store/auth-store'
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

  // Obter preferência salva (default: alpha-asc)
  const currentOrderBy = useMemo<LabResultsOrderByPreference>(() => {
    return (
      (user?.preferences as any)?.labResultsOrderBy || { type: 'alpha-asc' }
    )
  }, [user])

  // Salvar preferência (com debounce embutido na API)
  const setOrderBy = useCallback(
    async (orderBy: LabResultsOrderByPreference) => {
      try {
        const newPreferences = {
          ...(user?.preferences || {}),
          labResultsOrderBy: orderBy,
        }

        const updatedUser = await userApi.updatePreferences(newPreferences)
        updateUser(updatedUser)
      } catch (error) {
        console.error('Failed to save order by preference:', error)
      }
    },
    [user, updateUser]
  )

  return { currentOrderBy, setOrderBy }
}
