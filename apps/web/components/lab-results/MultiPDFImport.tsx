"use client"

import { useCallback, useState } from "react"
import { useDropzone } from "react-dropzone"
import { useRouter } from "next/navigation"
import { toast } from "sonner"
import { Upload, FileText, CheckCircle2, Loader2, AlertCircle, X } from "lucide-react"
import { Button } from "@/components/ui/button"
import { useAuthStore } from "@/lib/auth-store"
import { useProcessingJobStore } from "@/lib/processing-job-store"
import { labResultBatchApi } from "@/lib/api/lab-result-batch-api"

type ItemStatus = "pending" | "creating" | "uploading" | "queued" | "error"

interface Item {
  file: File
  status: ItemStatus
  error?: string
}

const MAX_SIZE = 20 * 1024 * 1024

/**
 * Importa VÁRIOS laudos em PDF de uma vez: cada PDF vira seu próprio lote.
 * O laboratório e a data de coleta são extraídos do próprio PDF (backend), e a
 * interpretação + classificação rodam em segundo plano.
 */
export function MultiPDFImport() {
  const router = useRouter()
  const accessToken = useAuthStore((s) => s.accessToken)
  const { addJob } = useProcessingJobStore()
  const [items, setItems] = useState<Item[]>([])
  const [running, setRunning] = useState(false)
  const [done, setDone] = useState(false)

  const onDrop = useCallback((accepted: File[]) => {
    const valid: Item[] = []
    for (const file of accepted) {
      if (file.type !== "application/pdf") {
        toast.error(`${file.name}: apenas PDFs são permitidos`)
        continue
      }
      if (file.size > MAX_SIZE) {
        toast.error(`${file.name}: arquivo muito grande (max 20MB)`)
        continue
      }
      valid.push({ file, status: "pending" })
    }
    setItems((prev) => [...prev, ...valid])
  }, [])

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: { "application/pdf": [".pdf"] },
    multiple: true,
    disabled: running,
  })

  const setStatus = (idx: number, status: ItemStatus, error?: string) =>
    setItems((prev) => prev.map((it, i) => (i === idx ? { ...it, status, error } : it)))

  const removeItem = (idx: number) =>
    setItems((prev) => prev.filter((_, i) => i !== idx))

  const uploadOne = async (item: Item, idx: number): Promise<boolean> => {
    // 1) cria o lote (lab/data placeholder — o backend preenche a partir do PDF)
    setStatus(idx, "creating")
    let batchId: string
    try {
      const batch = await labResultBatchApi.create({
        laboratoryName: "Importado via PDF",
        collectionDate: new Date().toISOString(),
        status: "pending",
        labResults: [{ testName: "Processando...", testType: "pending" }],
      } as any)
      batchId = batch.id
    } catch (e: any) {
      setStatus(idx, "error", e?.message || "falha ao criar lote")
      return false
    }

    // 2) upload do PDF → job de interpretação (processa em background)
    setStatus(idx, "uploading")
    try {
      const formData = new FormData()
      formData.append("file", item.file)
      const resp = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/v1/lab-result-batches/${batchId}/upload-pdf`,
        { method: "POST", headers: { Authorization: `Bearer ${accessToken}` }, body: formData },
      )
      if (!resp.ok) {
        const err = await resp.json().catch(() => ({}))
        throw new Error(err.message || "falha no upload")
      }
      const data = await resp.json()
      addJob(data.jobId, batchId)
      setStatus(idx, "queued")
      return true
    } catch (e: any) {
      // Rollback best-effort do lote órfão (delete é admin-only; ignora se não puder).
      labResultBatchApi.delete(batchId).catch(() => {})
      setStatus(idx, "error", e?.message || "falha no upload")
      return false
    }
  }

  const handleImportAll = async () => {
    const pending = items.map((it, i) => ({ it, i })).filter(({ it }) => it.status === "pending")
    if (pending.length === 0) return
    setRunning(true)
    let ok = 0
    // Sequencial para não sobrecarregar o worker/OCR.
    for (const { it, i } of pending) {
      if (await uploadOne(it, i)) ok++
    }
    setRunning(false)
    setDone(true)
    toast.success(`${ok} laudo(s) em processamento`, {
      description: "Cada PDF virou um lote. A interpretação roda em segundo plano.",
    })
  }

  const statusLabel = (s: ItemStatus) => {
    switch (s) {
      case "creating": return "criando lote…"
      case "uploading": return "enviando…"
      case "queued": return "em processamento"
      case "error": return "erro"
      default: return "pronto p/ enviar"
    }
  }

  const pendingCount = items.filter((it) => it.status === "pending").length

  return (
    <div className="space-y-4">
      <div
        {...getRootProps()}
        className={`border-2 border-dashed rounded-lg p-8 text-center cursor-pointer transition-colors ${
          isDragActive ? "border-primary bg-primary/5" : "border-border hover:border-primary/50"
        } ${running ? "opacity-50 cursor-not-allowed" : ""}`}
      >
        <input {...getInputProps()} />
        <Upload className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
        <p className="font-medium mb-2">
          {isDragActive ? "Solte os PDFs aqui" : "Arraste vários PDFs ou clique para selecionar"}
        </p>
        <p className="text-sm text-muted-foreground">
          Cada PDF vira um lote. Laboratório e data de coleta são lidos do próprio laudo. Máximo 20MB cada.
        </p>
      </div>

      {items.length > 0 && (
        <div className="space-y-2">
          {items.map((it, idx) => (
            <div key={idx} className="flex items-center gap-3 border rounded-md p-3">
              <FileText className="h-5 w-5 text-primary shrink-0" />
              <div className="flex-1 min-w-0">
                <p className="font-medium truncate">{it.file.name}</p>
                <p className="text-xs text-muted-foreground">
                  {(it.file.size / 1024 / 1024).toFixed(2)} MB · {statusLabel(it.status)}
                  {it.error ? ` (${it.error})` : ""}
                </p>
              </div>
              {it.status === "queued" && <CheckCircle2 className="h-5 w-5 text-green-600" />}
              {(it.status === "creating" || it.status === "uploading") && (
                <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
              )}
              {it.status === "error" && <AlertCircle className="h-5 w-5 text-red-600" />}
              {it.status === "pending" && !running && (
                <Button variant="ghost" size="icon" onClick={() => removeItem(idx)}>
                  <X className="h-4 w-4" />
                </Button>
              )}
            </div>
          ))}
        </div>
      )}

      <div className="flex justify-end gap-2">
        {done && (
          <Button variant="outline" onClick={() => router.push("/lab-results")}>
            Ver lotes
          </Button>
        )}
        <Button onClick={handleImportAll} disabled={running || pendingCount === 0}>
          {running ? (
            <>
              <Loader2 className="mr-2 h-4 w-4 animate-spin" /> Importando…
            </>
          ) : (
            <>
              <Upload className="mr-2 h-4 w-4" /> Importar {pendingCount > 0 ? pendingCount : ""} laudo(s)
            </>
          )}
        </Button>
      </div>
    </div>
  )
}
