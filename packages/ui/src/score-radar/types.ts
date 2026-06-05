export type RadarLetter = {
  code: string
  name: string
  score: number
  color: string
  /** Glifo exibido (default = code). Ex.: EN mostra A/C/T/S enquanto code segue A/G/I/R
   *  para a paleta e o casamento letra↔pilar. */
  label?: string
}

export type RadarPillar = {
  letter: string
  name: string
  score: number
}

export type AgirRadarData = { letters: RadarLetter[]; pillars: RadarPillar[] }

/**
 * Shape mínima do snapshot que `buildAgir` consome — estruturalmente compatível com o
 * `PatientScoreSnapshot` do EMR e com o snapshot por-pilar do escore-light (Fase 3).
 */
export type AgirSnapshotInput = {
  itemResults?: Array<{
    status: string
    actualPoints: number
    maxPoints: number
    item?: {
      methodPillars?: Array<{
        id: string
        name: string
        order?: number
        letter?: { code: string; name: string; color?: string; order?: number }
      }>
    }
  }>
}
