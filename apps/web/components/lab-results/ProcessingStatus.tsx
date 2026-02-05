"use client"

import { useEffect, useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { Loader2, CheckCircle2, AlertCircle, Clock } from "lucide-react"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Badge } from "@/components/ui/badge"
import { useAuthStore } from "@/lib/auth-store"

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
  errorMessage?: string
  attempts: number
  createdAt: string
  startedAt?: string
  completedAt?: string
}

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
    refetchInterval: pollingEnabled ? 3000 : false,
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

  const statusConfig = {
    pending: {
      icon: Clock,
      label: "Aguardando processamento",
      color: "text-yellow-600",
      bgColor: "bg-yellow-50 border-yellow-200",
    },
    processing: {
      icon: Loader2,
      label: "Interpretando PDF com IA",
      color: "text-blue-600",
      bgColor: "bg-blue-50 border-blue-200",
      animate: true,
    },
    completed: {
      icon: CheckCircle2,
      label: "Interpretação concluída",
      color: "text-green-600",
      bgColor: "bg-green-50 border-green-200",
    },
    failed: {
      icon: AlertCircle,
      label: "Erro na interpretação",
      color: "text-red-600",
      bgColor: "bg-red-50 border-red-200",
    },
  }

  if (!job) return null

  const config = statusConfig[job.status as keyof typeof statusConfig]
  const Icon = config.icon

  return (
    <Alert className={config.bgColor}>
      <div className="flex items-start gap-3">
        <Icon className={`h-5 w-5 ${config.color} ${config.animate ? "animate-spin" : ""} mt-0.5`} />
        <div className="flex-1">
          <div className="flex items-center gap-2 flex-wrap">
            <span className={`font-medium ${config.color}`}>{config.label}</span>
            {job.attempts > 1 && (
              <Badge variant="outline" className="text-xs">
                Tentativa {job.attempts}
              </Badge>
            )}
          </div>

          {job.status === "processing" && (
            <p className="text-sm text-muted-foreground mt-1">
              Extraindo texto via OCR e interpretando com Claude AI...
            </p>
          )}

          {job.status === "completed" && (
            <p className="text-sm text-muted-foreground mt-1">
              Resultados preenchidos automaticamente. Revise e confirme os dados.
            </p>
          )}

          {job.status === "failed" && job.errorMessage && (
            <AlertDescription className="mt-2 text-sm">
              {job.errorMessage}
            </AlertDescription>
          )}
        </div>
      </div>
    </Alert>
  )
}
