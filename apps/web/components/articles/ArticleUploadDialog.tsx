'use client'

import { useState, useCallback } from 'react'
import { useDropzone } from 'react-dropzone'
import { toast } from 'sonner'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { Progress } from '@/components/ui/progress'
import { Upload, FileText, X, CheckCircle2, AlertCircle } from 'lucide-react'
import { useUploadArticle } from '@/lib/api/article-api'

interface ArticleUploadDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function ArticleUploadDialog({
  open,
  onOpenChange,
}: ArticleUploadDialogProps) {
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const [uploadProgress, setUploadProgress] = useState(0)
  const uploadMutation = useUploadArticle()

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = acceptedFiles[0]
    if (file) {
      // Validar tamanho (max 50MB)
      if (file.size > 50 * 1024 * 1024) {
        toast.error('Arquivo muito grande. Máximo: 50MB')
        return
      }

      // Validar tipo
      if (file.type !== 'application/pdf') {
        toast.error('Apenas arquivos PDF são permitidos')
        return
      }

      setSelectedFile(file)
      setUploadProgress(0)
    }
  }, [])

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      'application/pdf': ['.pdf'],
    },
    maxFiles: 1,
    disabled: uploadMutation.isPending,
  })

  const handleUpload = async () => {
    if (!selectedFile) return

    try {
      // Simular progresso
      const interval = setInterval(() => {
        setUploadProgress((prev) => {
          if (prev >= 90) {
            clearInterval(interval)
            return 90
          }
          return prev + 10
        })
      }, 200)

      await uploadMutation.mutateAsync(selectedFile)

      clearInterval(interval)
      setUploadProgress(100)

      toast.success('PDF importado com sucesso!', {
        description: 'Metadados foram extraídos automaticamente.',
      })

      setTimeout(() => {
        onOpenChange(false)
        setSelectedFile(null)
        setUploadProgress(0)
      }, 1500)
    } catch (error: any) {
      toast.error('Erro ao fazer upload', {
        description: error.response?.data?.error || error.message,
      })
      setUploadProgress(0)
    }
  }

  const handleRemoveFile = () => {
    setSelectedFile(null)
    setUploadProgress(0)
  }

  const formatFileSize = (bytes: number) => {
    if (bytes === 0) return '0 Bytes'
    const k = 1024
    const sizes = ['Bytes', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i]
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Upload de Artigo Científico</DialogTitle>
          <DialogDescription>
            Faça upload de um PDF e os metadados serão extraídos automaticamente
            (DOI, PMID, título, abstract, autores).
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* Dropzone */}
          {!selectedFile && (
            <div
              {...getRootProps()}
              className={`
                border-2 border-dashed rounded-lg p-8 text-center cursor-pointer
                transition-colors duration-200
                ${
                  isDragActive
                    ? 'border-primary bg-primary/5'
                    : 'border-border hover:border-primary/50 hover:bg-muted/50'
                }
                ${uploadMutation.isPending ? 'opacity-50 cursor-not-allowed' : ''}
              `}
            >
              <input {...getInputProps()} />
              <Upload className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
              <p className="text-base font-medium mb-2">
                {isDragActive
                  ? 'Solte o arquivo PDF aqui'
                  : 'Arraste um PDF ou clique para selecionar'}
              </p>
              <p className="text-sm text-muted-foreground">
                Máximo: 50MB • Apenas arquivos PDF
              </p>
            </div>
          )}

          {/* Selected File */}
          {selectedFile && (
            <div className="border rounded-lg p-4">
              <div className="flex items-start gap-4">
                <div className="flex-shrink-0">
                  <FileText className="h-10 w-10 text-primary" />
                </div>

                <div className="flex-1 min-w-0">
                  <div className="flex items-start justify-between gap-2">
                    <div className="flex-1 min-w-0">
                      <p className="font-medium truncate">{selectedFile.name}</p>
                      <p className="text-sm text-muted-foreground">
                        {formatFileSize(selectedFile.size)}
                      </p>
                    </div>

                    {!uploadMutation.isPending && uploadProgress === 0 && (
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={handleRemoveFile}
                        className="flex-shrink-0"
                      >
                        <X className="h-4 w-4" />
                      </Button>
                    )}
                  </div>

                  {/* Progress */}
                  {uploadProgress > 0 && (
                    <div className="mt-3 space-y-2">
                      <Progress value={uploadProgress} className="h-2" />
                      <div className="flex items-center gap-2 text-sm">
                        {uploadProgress === 100 ? (
                          <>
                            <CheckCircle2 className="h-4 w-4 text-green-600" />
                            <span className="text-green-600 font-medium">
                              Upload concluído!
                            </span>
                          </>
                        ) : (
                          <span className="text-muted-foreground">
                            Enviando e extraindo metadados... {uploadProgress}%
                          </span>
                        )}
                      </div>
                    </div>
                  )}

                  {/* Error */}
                  {uploadMutation.isError && (
                    <div className="mt-3 flex items-center gap-2 text-sm text-destructive">
                      <AlertCircle className="h-4 w-4" />
                      <span>Erro ao fazer upload do arquivo</span>
                    </div>
                  )}
                </div>
              </div>
            </div>
          )}

          {/* Info */}
          <div className="bg-muted/50 rounded-lg p-4">
            <h4 className="font-medium text-sm mb-2">
              O que será extraído automaticamente:
            </h4>
            <ul className="text-sm text-muted-foreground space-y-1">
              <li>• <strong>DOI e PMID</strong> (se presentes no PDF)</li>
              <li>• <strong>Título</strong> e <strong>Abstract</strong></li>
              <li>• <strong>Texto completo</strong> do documento</li>
              <li>
                • <strong>Metadados adicionais</strong> via CrossRef API (autores,
                journal, keywords)
              </li>
            </ul>
          </div>

          {/* Actions */}
          <div className="flex justify-end gap-3">
            <Button
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={uploadMutation.isPending}
            >
              Cancelar
            </Button>
            <Button
              onClick={handleUpload}
              disabled={!selectedFile || uploadMutation.isPending || uploadProgress > 0}
            >
              {uploadMutation.isPending ? (
                <>
                  <span className="animate-spin mr-2">⏳</span>
                  Processando...
                </>
              ) : (
                <>
                  <Upload className="mr-2 h-4 w-4" />
                  Importar Artigo
                </>
              )}
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
