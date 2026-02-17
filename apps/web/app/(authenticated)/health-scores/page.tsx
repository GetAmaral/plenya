"use client"

import { useState } from "react"
import { useRequireAuth } from "@/lib/use-auth"
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient"
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader"
import { PageHeader } from "@/components/layout/page-header"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Activity, Plus, AlertCircle } from "lucide-react"
import { useLatestHealthScore, useHealthScores } from "@/lib/api/health-score-api"
import { ScoreSnapshotSummary } from "@/components/health-scores/ScoreSnapshotSummary"
import { ScoreHistoryTable } from "@/components/health-scores/ScoreHistoryTable"
import { CalculateScoreDialog } from "@/components/health-scores/CalculateScoreDialog"
import { ScoreRadarChart } from "@/components/health-scores/ScoreRadarChart"
import { ScoreEvolutionChart } from "@/components/health-scores/ScoreEvolutionChart"
import { ScoreMethodologyAccordion } from "@/components/health-scores/ScoreMethodologyAccordion"
import { Skeleton } from "@/components/ui/skeleton"
import { Alert, AlertDescription } from "@/components/ui/alert"

export default function HealthScoresPage() {
  useRequireAuth()
  const { selectedPatient } = useRequireSelectedPatient()
  const [showCalculateDialog, setShowCalculateDialog] = useState(false)

  const {
    data: latestSnapshot,
    isLoading: isLoadingLatest,
    error: latestError,
  } = useLatestHealthScore(selectedPatient?.id || "")

  const {
    data: snapshots,
    isLoading: isLoadingHistory,
  } = useHealthScores(selectedPatient?.id || "")

  return (
    <div className="flex flex-col gap-6">
      {/* Header */}
      <SelectedPatientHeader />

      <PageHeader
        title="Escores de Saúde"
        description="Acompanhe a evolução dos indicadores de saúde do paciente"
      >
        <Button onClick={() => setShowCalculateDialog(true)}>
          <Plus className="mr-2 h-4 w-4" />
          Calcular Escore
        </Button>
      </PageHeader>

      {/* Latest Snapshot Summary */}
      {isLoadingLatest ? (
        <Card>
          <CardHeader>
            <Skeleton className="h-8 w-64" />
          </CardHeader>
          <CardContent>
            <Skeleton className="h-32 w-full" />
          </CardContent>
        </Card>
      ) : latestError ? (
        <Alert>
          <AlertCircle className="h-4 w-4" />
          <AlertDescription>
            Nenhum escore calculado ainda. Clique em &quot;Calcular Escore&quot; para gerar o
            primeiro snapshot.
          </AlertDescription>
        </Alert>
      ) : latestSnapshot ? (
        <>
          {/* Card de Score Total (sem os cards de detalhes) */}
          <ScoreSnapshotSummary snapshot={latestSnapshot} showDetailCards={false}>
            <ScoreRadarChart snapshot={latestSnapshot} />
          </ScoreSnapshotSummary>

          {/* Accordion de metodologia (igual à página de detalhes) */}
          <ScoreMethodologyAccordion snapshot={latestSnapshot} />

          {/* Evolution Chart */}
          {snapshots && snapshots.length > 1 && (
            <ScoreEvolutionChart snapshots={snapshots} />
          )}
        </>
      ) : null}

      {/* History Table */}
      <Card>
        <CardHeader>
          <CardTitle>Histórico de Avaliações</CardTitle>
          <CardDescription>
            Todos os snapshots de escores calculados anteriormente
          </CardDescription>
        </CardHeader>
        <CardContent>
          {isLoadingHistory ? (
            <div className="space-y-2">
              <Skeleton className="h-12 w-full" />
              <Skeleton className="h-12 w-full" />
              <Skeleton className="h-12 w-full" />
            </div>
          ) : snapshots && snapshots.length > 0 ? (
            <ScoreHistoryTable snapshots={snapshots} />
          ) : (
            <p className="text-center text-muted-foreground py-8">
              Nenhum snapshot encontrado
            </p>
          )}
        </CardContent>
      </Card>

      {/* Calculate Dialog */}
      <CalculateScoreDialog
        patientId={selectedPatient?.id || ""}
        open={showCalculateDialog}
        onOpenChange={setShowCalculateDialog}
      />
    </div>
  )
}
