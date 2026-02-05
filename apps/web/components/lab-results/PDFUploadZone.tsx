"use client"

import { useCallback, useState } from "react"
import { useDropzone } from "react-dropzone"
import { toast } from "sonner"
import { Upload, FileText, X, Loader2, CheckCircle2 } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Progress } from "@/components/ui/progress"
import { useAuthStore } from "@/lib/auth-store"

interface PDFUploadZoneProps {
  batchId?: string
  onUploadSuccess: (jobId: string) => void
  onBeforeUpload?: () => Promise<string> // Retorna o batchId (usado para criar batch antes do upload)
  disabled?: boolean
}

export function PDFUploadZone({ batchId, onUploadSuccess, onBeforeUpload, disabled }: PDFUploadZoneProps) {
  const accessToken = useAuthStore((state) => state.accessToken)
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const [isUploading, setIsUploading] = useState(false)
  const [uploadProgress, setUploadProgress] = useState(0)

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = acceptedFiles[0]
    if (!file) return

    if (file.size > 20 * 1024 * 1024) {
      toast.error("Arquivo muito grande (max 20MB)")
      return
    }

    if (file.type !== "application/pdf") {
      toast.error("Apenas PDFs são permitidos")
      return
    }

    setSelectedFile(file)
    setUploadProgress(0)
  }, [])

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: { "application/pdf": [".pdf"] },
    maxFiles: 1,
    disabled: disabled || isUploading,
  })

  const handleUpload = async () => {
    if (!selectedFile) return

    try {
      setIsUploading(true)

      // Se não tem batchId, criar batch primeiro
      let effectiveBatchId = batchId
      if (!effectiveBatchId && onBeforeUpload) {
        try {
          effectiveBatchId = await onBeforeUpload()
        } catch (error: any) {
          toast.error("Erro ao criar batch", { description: error.message })
          setIsUploading(false)
          return
        }
      }

      if (!effectiveBatchId) {
        toast.error("Erro: Batch ID não disponível")
        setIsUploading(false)
        return
      }

      const formData = new FormData()
      formData.append("file", selectedFile)

      // Simular progresso durante upload
      const progressInterval = setInterval(() => {
        setUploadProgress((prev) => Math.min(prev + 15, 90))
      }, 200)

      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/lab-result-batches/${effectiveBatchId}/upload-pdf`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: formData,
      })

      clearInterval(progressInterval)
      setUploadProgress(100)

      if (!response.ok) {
        const error = await response.json()
        throw new Error(error.message || "Upload failed")
      }

      const data = await response.json()

      toast.success("PDF enviado!", {
        description: "Iniciando interpretação automática...",
      })

      onUploadSuccess(data.jobId)
    } catch (error: any) {
      toast.error("Erro no upload", { description: error.message })
      setUploadProgress(0)
    } finally {
      setIsUploading(false)
    }
  }

  const handleRemove = () => {
    setSelectedFile(null)
    setUploadProgress(0)
  }

  return (
    <div className="space-y-4">
      {!selectedFile && (
        <div
          {...getRootProps()}
          className={`
            border-2 border-dashed rounded-lg p-8 text-center cursor-pointer
            transition-colors
            ${isDragActive ? "border-primary bg-primary/5" : "border-border hover:border-primary/50"}
            ${disabled ? "opacity-50 cursor-not-allowed" : ""}
          `}
        >
          <input {...getInputProps()} />
          <Upload className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
          <p className="font-medium mb-2">
            {isDragActive ? "Solte o PDF aqui" : "Arraste PDF ou clique para selecionar"}
          </p>
          <p className="text-sm text-muted-foreground">Máximo: 20MB</p>
        </div>
      )}

      {selectedFile && (
        <div className="border rounded-lg p-4">
          <div className="flex items-start gap-4">
            <FileText className="h-10 w-10 text-primary flex-shrink-0" />
            <div className="flex-1 min-w-0">
              <div className="flex items-start justify-between gap-2">
                <div className="flex-1">
                  <p className="font-medium truncate">{selectedFile.name}</p>
                  <p className="text-sm text-muted-foreground">
                    {(selectedFile.size / 1024 / 1024).toFixed(2)} MB
                  </p>
                </div>
                {!isUploading && uploadProgress === 0 && (
                  <Button variant="ghost" size="icon" onClick={handleRemove}>
                    <X className="h-4 w-4" />
                  </Button>
                )}
              </div>

              {uploadProgress > 0 && (
                <div className="mt-3 space-y-2">
                  <Progress value={uploadProgress} className="h-2" />
                  {uploadProgress === 100 ? (
                    <div className="flex items-center gap-2 text-sm text-green-600">
                      <CheckCircle2 className="h-4 w-4" />
                      Upload concluído!
                    </div>
                  ) : (
                    <div className="flex items-center gap-2 text-sm">
                      <Loader2 className="h-4 w-4 animate-spin" />
                      Enviando... {uploadProgress}%
                    </div>
                  )}
                </div>
              )}
            </div>
          </div>

          {!isUploading && uploadProgress === 0 && (
            <div className="mt-4 flex justify-end">
              <Button onClick={handleUpload}>
                <Upload className="mr-2 h-4 w-4" />
                Interpretar PDF com IA
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  )
}
