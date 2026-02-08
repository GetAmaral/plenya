"use client"

import { useEffect, useRef } from "react"
import { useQuery } from "@tanstack/react-query"
import { useRouter } from "next/navigation"
import { toast } from "sonner"
import { FileCheck, Loader2 } from "lucide-react"
import { useProcessingJobStore } from "@/lib/processing-job-store"
import { useAuthStore } from "@/lib/auth-store"
import { Button } from "@/components/ui/button"

type JobStatus = "pending" | "processing" | "completed" | "failed"

interface ProcessingJob {
  id: string
  labResultBatchId: string
  type: string
  status: JobStatus
  progressStep?: number
  progressMessage?: string
  errorMessage?: string
  createdAt: string
}

const STEP_MESSAGES = [
  "", // step 0
  "Enviando PDF ao servidor",
  "Upload completo",
  "Extraindo conteúdo do PDF",
  "Analisando com IA",
  "Análise concluída",
  "Salvando resultados",
]

export function GlobalProcessingMonitor() {
  const router = useRouter()
  const accessToken = useAuthStore((state) => state.accessToken)
  const { activeJobs, removeJob, updateJobStep } = useProcessingJobStore()

  // Track which jobs have shown completion toast
  const completedJobsRef = useRef<Set<string>>(new Set())
  // Track last step shown for each job
  const lastStepRef = useRef<Map<string, number>>(new Map())

  // Poll all active jobs
  const { data: jobs } = useQuery<ProcessingJob[]>({
    queryKey: ["processing-jobs-global", activeJobs.map(j => j.id)],
    queryFn: async () => {
      if (activeJobs.length === 0) return []

      const results = await Promise.all(
        activeJobs.map(async (job) => {
          try {
            const res = await fetch(
              `${process.env.NEXT_PUBLIC_API_URL}/api/v1/processing-jobs/${job.id}`,
              {
                headers: {
                  Authorization: `Bearer ${accessToken}`,
                },
              }
            )
            if (!res.ok) return null
            return res.json()
          } catch {
            return null
          }
        })
      )

      return results.filter((j): j is ProcessingJob => j !== null)
    },
    refetchInterval: activeJobs.length > 0 ? 2000 : false,
    enabled: activeJobs.length > 0 && !!accessToken,
  })

  useEffect(() => {
    if (!jobs) return

    jobs.forEach((job) => {
      const currentStep = job.progressStep || 0
      const lastStep = lastStepRef.current.get(job.id) || 0

      // Show progress toast for new steps
      if (currentStep > lastStep && currentStep > 0 && currentStep <= 6) {
        const message = STEP_MESSAGES[currentStep] || job.progressMessage || "Processando..."

        toast.loading(message, {
          id: `processing-${job.id}`,
          icon: <Loader2 className="h-4 w-4 animate-spin" />,
          description: `Etapa ${currentStep} de 6`,
        })

        lastStepRef.current.set(job.id, currentStep)
        updateJobStep(job.id, currentStep)
      }

      // Handle completion
      if (job.status === "completed" && !completedJobsRef.current.has(job.id)) {
        completedJobsRef.current.add(job.id)
        lastStepRef.current.delete(job.id)

        // Dismiss loading toast
        toast.dismiss(`processing-${job.id}`)

        // Show permanent success toast with actions
        toast.success("Processamento concluído!", {
          id: `completed-${job.id}`,
          icon: <FileCheck className="h-4 w-4" />,
          description: "Laudo médico processado com sucesso",
          duration: Infinity, // Permanent
          action: {
            label: "Ver",
            onClick: () => {
              router.push("/lab-results")
              toast.dismiss(`completed-${job.id}`)
            },
          },
          cancel: {
            label: "Fechar",
            onClick: () => {
              toast.dismiss(`completed-${job.id}`)
            },
          },
        })

        removeJob(job.id)
      }

      // Handle failure
      if (job.status === "failed" && !completedJobsRef.current.has(job.id)) {
        completedJobsRef.current.add(job.id)
        lastStepRef.current.delete(job.id)

        // Dismiss loading toast
        toast.dismiss(`processing-${job.id}`)

        toast.error("Erro no processamento", {
          id: `failed-${job.id}`,
          description: job.errorMessage || "Falha ao processar o PDF",
          duration: 10000,
        })

        removeJob(job.id)
      }
    })
  }, [jobs, removeJob, updateJobStep, router])

  return null // This component only handles side effects
}
