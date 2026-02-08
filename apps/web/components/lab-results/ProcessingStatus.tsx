"use client"

import { useEffect, useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { Loader2, CheckCircle2, AlertCircle, Clock, Upload, FileText, Brain, Database, Check } from "lucide-react"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Badge } from "@/components/ui/badge"
import { Progress } from "@/components/ui/progress"
import { useAuthStore } from "@/lib/auth-store"
import { cn } from "@/lib/utils"

interface ProcessingStatusProps {
  jobId: string
  onCompleted: () => void
  onFailed: (error: string) => void
}

type JobStatus = "pending" | "processing" | "completed" | "failed"

interface ProcessingJob {
  id: string
  labResultBatchId: string
  type: string
  status: JobStatus
  progressStep?: number
  progressMessage?: string
  errorMessage?: string
  attempts: number
  createdAt: string
  startedAt?: string
  completedAt?: string
}

interface ProcessingStep {
  number: number
  label: string
  icon: React.ComponentType<{ className?: string }>
}

const PROCESSING_STEPS: ProcessingStep[] = [
  { number: 1, label: "Enviando PDF ao servidor", icon: Upload },
  { number: 2, label: "Upload completo", icon: CheckCircle2 },
  { number: 3, label: "Extraindo conteúdo do PDF", icon: FileText },
  { number: 4, label: "Analisando com IA", icon: Brain },
  { number: 5, label: "Análise concluída", icon: Check },
  { number: 6, label: "Salvando resultados", icon: Database },
]

export function ProcessingStatus({ jobId, onCompleted, onFailed }: ProcessingStatusProps) {
  const accessToken = useAuthStore((state) => state.accessToken)
  const [pollingEnabled, setPollingEnabled] = useState(true)

  const { data: job } = useQuery<ProcessingJob>({
    queryKey: ["processing-job", jobId],
    queryFn: async () => {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/processing-jobs/${jobId}`, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      })
      if (!res.ok) throw new Error("Failed to fetch job")
      return res.json()
    },
    refetchInterval: pollingEnabled ? 1000 : false, // Poll every 1s for better UX
    enabled: !!jobId,
  })

  useEffect(() => {
    if (!job) return

    if (job.status === "completed") {
      setPollingEnabled(false)
      onCompleted()
    } else if (job.status === "failed") {
      setPollingEnabled(false)
      onFailed(job.errorMessage || "Erro desconhecido")
    }
  }, [job, onCompleted, onFailed])

  if (!job) return null

  // Calculate progress percentage
  const currentStep = job.progressStep || 0
  const progressPercentage = job.status === "completed" ? 100 : (currentStep / 6) * 100

  // Failed state
  if (job.status === "failed") {
    return (
      <Alert className="bg-red-50 border-red-200">
        <div className="flex items-start gap-3">
          <AlertCircle className="h-5 w-5 text-red-600 mt-0.5" />
          <div className="flex-1">
            <div className="flex items-center gap-2 flex-wrap">
              <span className="font-medium text-red-600">Erro na interpretação</span>
              {job.attempts > 1 && (
                <Badge variant="outline" className="text-xs">
                  Tentativa {job.attempts}
                </Badge>
              )}
            </div>
            {job.errorMessage && (
              <AlertDescription className="mt-2 text-sm">
                {job.errorMessage}
              </AlertDescription>
            )}
          </div>
        </div>
      </Alert>
    )
  }

  // Completed state
  if (job.status === "completed") {
    return (
      <Alert className="bg-green-50 border-green-200">
        <div className="flex items-start gap-3">
          <CheckCircle2 className="h-5 w-5 text-green-600 mt-0.5" />
          <div className="flex-1">
            <span className="font-medium text-green-600">Processamento concluído!</span>
            <p className="text-sm text-muted-foreground mt-1">
              {job.progressMessage || "Resultados preenchidos automaticamente. Revise e confirme os dados."}
            </p>
          </div>
        </div>
      </Alert>
    )
  }

  // Processing/Pending state - show detailed steps
  return (
    <Alert className="bg-blue-50 border-blue-200">
      <div className="space-y-4">
        {/* Header */}
        <div className="flex items-start gap-3">
          <Loader2 className="h-5 w-5 text-blue-600 animate-spin mt-0.5" />
          <div className="flex-1">
            <div className="flex items-center gap-2 flex-wrap">
              <span className="font-medium text-blue-600">
                {job.progressMessage || "Processando laudo médico..."}
              </span>
              {job.attempts > 1 && (
                <Badge variant="outline" className="text-xs">
                  Tentativa {job.attempts}
                </Badge>
              )}
            </div>
            <p className="text-sm text-muted-foreground mt-1">
              Aguarde enquanto processamos seu PDF
            </p>
          </div>
        </div>

        {/* Progress Bar */}
        <div className="space-y-2">
          <Progress value={progressPercentage} className="h-2" />
          <div className="text-xs text-muted-foreground text-right">
            {Math.round(progressPercentage)}% completo
          </div>
        </div>

        {/* Steps List */}
        <div className="space-y-2">
          {PROCESSING_STEPS.map((step) => {
            const isCompleted = currentStep > step.number
            const isCurrent = currentStep === step.number
            const isPending = currentStep < step.number

            const Icon = step.icon

            return (
              <div
                key={step.number}
                className={cn(
                  "flex items-center gap-3 py-2 px-3 rounded-lg transition-colors",
                  isCurrent && "bg-blue-100/50",
                  isCompleted && "bg-green-50/50"
                )}
              >
                <div
                  className={cn(
                    "flex items-center justify-center w-6 h-6 rounded-full text-xs font-medium shrink-0",
                    isCompleted && "bg-green-600 text-white",
                    isCurrent && "bg-blue-600 text-white",
                    isPending && "bg-gray-200 text-gray-500"
                  )}
                >
                  {isCompleted ? (
                    <CheckCircle2 className="h-4 w-4" />
                  ) : isCurrent ? (
                    <Loader2 className="h-4 w-4 animate-spin" />
                  ) : (
                    <Icon className="h-4 w-4" />
                  )}
                </div>

                <div className="flex-1 min-w-0">
                  <p
                    className={cn(
                      "text-sm font-medium truncate",
                      isCompleted && "text-green-700",
                      isCurrent && "text-blue-700",
                      isPending && "text-gray-500"
                    )}
                  >
                    {step.label}
                  </p>

                  {/* Show custom message for step 5 (exam count) */}
                  {isCurrent && step.number === 5 && job.progressMessage && job.progressMessage.includes("exames identificados") && (
                    <p className="text-xs text-blue-600 mt-0.5">
                      {job.progressMessage.split(" - ")[1]}
                    </p>
                  )}
                </div>

                {isCompleted && (
                  <CheckCircle2 className="h-4 w-4 text-green-600 shrink-0" />
                )}
              </div>
            )
          })}
        </div>
      </div>
    </Alert>
  )
}
