import { useEffect, useState } from 'react'

/**
 * Hook para gerenciar preferências de visualização do Health Score Detail
 *
 * Salva estado de:
 * - Grupos abertos/fechados
 * - Subgrupos abertos/fechados
 * - Item selecionado
 * - Filtro "apenas avaliados"
 *
 * Usa sessionStorage (limpa ao fechar aba, mantém durante navegação)
 */

interface HealthScoreViewPreferences {
  openGroups: string[]
  openSubgroups: string[]
  selectedItemId: string | null
  showOnlyEvaluated: boolean
}

const STORAGE_KEY_PREFIX = 'health-score-view-prefs'

function getStorageKey(snapshotId: string): string {
  return `${STORAGE_KEY_PREFIX}-${snapshotId}`
}

function loadPreferences(snapshotId: string): HealthScoreViewPreferences | null {
  if (typeof window === 'undefined') return null

  try {
    const stored = sessionStorage.getItem(getStorageKey(snapshotId))
    if (!stored) return null

    return JSON.parse(stored)
  } catch (error) {
    console.warn('Failed to load health score view preferences:', error)
    return null
  }
}

function savePreferences(snapshotId: string, prefs: HealthScoreViewPreferences): void {
  if (typeof window === 'undefined') return

  try {
    sessionStorage.setItem(getStorageKey(snapshotId), JSON.stringify(prefs))
  } catch (error) {
    console.warn('Failed to save health score view preferences:', error)
  }
}

export function useHealthScoreViewPreferences(snapshotId: string) {
  // Flag para controlar se já carregou o estado inicial
  const [hasLoadedInitialState, setHasLoadedInitialState] = useState(false)

  // Estado de grupos abertos (array de group IDs)
  const [openGroups, setOpenGroups] = useState<string[]>([])

  // Estado de subgrupos abertos (array de subgroup IDs)
  const [openSubgroups, setOpenSubgroups] = useState<string[]>([])

  // Item selecionado (item ID)
  const [selectedItemId, setSelectedItemId] = useState<string | null>(null)

  // Filtro de visualização
  const [showOnlyEvaluated, setShowOnlyEvaluated] = useState(true)

  // Carregar preferências salvas ao montar
  useEffect(() => {
    const saved = loadPreferences(snapshotId)
    if (saved) {
      setOpenGroups(saved.openGroups || [])
      setOpenSubgroups(saved.openSubgroups || [])
      setSelectedItemId(saved.selectedItemId || null)
      setShowOnlyEvaluated(saved.showOnlyEvaluated ?? true)
    }
    // Marca que já carregou o estado inicial
    setHasLoadedInitialState(true)
  }, [snapshotId])

  // Salvar preferências sempre que mudam (MAS SÓ DEPOIS DE CARREGAR)
  useEffect(() => {
    // Não salva até ter carregado o estado inicial (evita sobrescrever com valores vazios)
    if (!hasLoadedInitialState) return

    const prefs: HealthScoreViewPreferences = {
      openGroups,
      openSubgroups,
      selectedItemId,
      showOnlyEvaluated,
    }
    savePreferences(snapshotId, prefs)
  }, [hasLoadedInitialState, snapshotId, openGroups, openSubgroups, selectedItemId, showOnlyEvaluated])

  return {
    // Estados
    openGroups,
    openSubgroups,
    selectedItemId,
    showOnlyEvaluated,

    // Setters
    setOpenGroups,
    setOpenSubgroups,
    setSelectedItemId,
    setShowOnlyEvaluated,
  }
}

/**
 * Limpa preferências de visualização de um snapshot específico
 */
export function clearHealthScoreViewPreferences(snapshotId: string): void {
  if (typeof window === 'undefined') return

  try {
    sessionStorage.removeItem(getStorageKey(snapshotId))
  } catch (error) {
    console.warn('Failed to clear health score view preferences:', error)
  }
}

/**
 * Limpa TODAS as preferências de health score (útil para limpeza geral)
 */
export function clearAllHealthScoreViewPreferences(): void {
  if (typeof window === 'undefined') return

  try {
    const keys = Object.keys(sessionStorage)
    keys.forEach(key => {
      if (key.startsWith(STORAGE_KEY_PREFIX)) {
        sessionStorage.removeItem(key)
      }
    })
  } catch (error) {
    console.warn('Failed to clear all health score view preferences:', error)
  }
}
