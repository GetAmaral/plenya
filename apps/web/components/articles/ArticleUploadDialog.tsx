'use client'

import { useState, useCallback } from 'react'
import { useDropzone } from 'react-dropzone'
import { useQueryClient } from '@tanstack/react-query'
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
import { Switch } from '@/components/ui/switch'
import { Label } from '@/components/ui/label'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Upload,
  FileText,
  X,
  CheckCircle2,
  AlertCircle,
  AlertTriangle,
  BookOpen,
  Loader2,
} from 'lucide-react'
import { articleApi } from '@/lib/api/article-api'

// ─── tipos ────────────────────────────────────────────────────────────────────

type FileStatus = 'pending' | 'uploading' | 'done' | 'duplicate' | 'error'

interface FileEntry {
  id: string
  file: File
  asBook: boolean
  status: FileStatus
  progress: number
  error?: string
  chapterCount?: number
  existingTitle?: string // para status 'duplicate'
}

interface ArticleUploadDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
}

// ─── helpers ──────────────────────────────────────────────────────────────────

const ALLOWED_EXTENSIONS = ['.pdf', '.epub', '.txt', '.md']
const MAX_SIZE = 200 * 1024 * 1024 // 200 MB

const getFileExt = (file: File) =>
  '.' + (file.name.split('.').pop()?.toLowerCase() ?? '')

const isLikelyBook = (file: File) => {
  const ext = getFileExt(file)
  return (
    ext === '.epub' ||
    (file.size > 5 * 1024 * 1024 &&
      (ext === '.pdf' || ext === '.txt' || ext === '.md'))
  )
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return Math.round((bytes / Math.pow(k, i)) * 100) / 100 + ' ' + sizes[i]
}

// ─── componente principal ─────────────────────────────────────────────────────

export function ArticleUploadDialog({ open, onOpenChange }: ArticleUploadDialogProps) {
  const [files, setFiles] = useState<FileEntry[]>([])
  const [isProcessing, setIsProcessing] = useState(false)
  const queryClient = useQueryClient()

  const updateFile = useCallback((id: string, patch: Partial<FileEntry>) => {
    setFiles(prev => prev.map(e => (e.id === id ? { ...e, ...patch } : e)))
  }, [])

  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      const newEntries: FileEntry[] = []
      const rejected: string[] = []

      for (const file of acceptedFiles) {
        const ext = getFileExt(file)
        if (!ALLOWED_EXTENSIONS.includes(ext)) {
          rejected.push(file.name)
          continue
        }
        if (file.size > MAX_SIZE) {
          rejected.push(`${file.name} (muito grande)`)
          continue
        }
        // Ignorar duplicatas já na lista
        if (files.some(e => e.file.name === file.name && e.file.size === file.size)) {
          continue
        }
        newEntries.push({
          id: `${file.name}-${file.size}-${Math.random()}`,
          file,
          asBook: getFileExt(file) === '.epub',
          status: 'pending',
          progress: 0,
        })
      }

      if (rejected.length > 0) {
        toast.error(`${rejected.length} arquivo(s) rejeitado(s)`, {
          description: rejected.slice(0, 3).join(', ') + (rejected.length > 3 ? '…' : ''),
        })
      }
      if (newEntries.length > 0) {
        setFiles(prev => [...prev, ...newEntries])
      }
    },
    [files],
  )

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      'application/pdf': ['.pdf'],
      'application/epub+zip': ['.epub'],
      'text/plain': ['.txt'],
      'text/markdown': ['.md'],
      'text/x-markdown': ['.md'],
    },
    multiple: true,
    disabled: isProcessing,
  })

  const removeFile = (id: string) =>
    setFiles(prev => prev.filter(e => e.id !== id))

  const toggleBook = (id: string, value: boolean) =>
    setFiles(prev => prev.map(e => (e.id === id ? { ...e, asBook: value } : e)))

  // ── upload sequencial ──────────────────────────────────────────────────────

  const handleUpload = async () => {
    const pending = files.filter(e => e.status === 'pending')
    if (pending.length === 0) return

    setIsProcessing(true)
    let successCount = 0
    let duplicateCount = 0
    let errorCount = 0

    for (const entry of pending) {
      updateFile(entry.id, { status: 'uploading', progress: 0 })

      // Simular progresso enquanto o servidor processa
      const step = entry.asBook ? 5 : 10
      const intervalMs = entry.asBook ? 500 : 200
      const interval = setInterval(() => {
        setFiles(prev =>
          prev.map(e => {
            if (e.id !== entry.id) return e
            const next = Math.min(e.progress + step, 85)
            return { ...e, progress: next }
          }),
        )
      }, intervalMs)

      try {
        const result = await articleApi.upload(entry.file, entry.asBook)
        clearInterval(interval)
        updateFile(entry.id, {
          status: 'done',
          progress: 100,
          chapterCount: result.totalChapters,
        })
        successCount++
      } catch (err: any) {
        clearInterval(interval)
        // 409 = arquivo duplicado (não é um erro real, é um aviso)
        if (err.message === 'DUPLICATE_FILE' || (err as any).data?.error === 'DUPLICATE_FILE') {
          updateFile(entry.id, {
            status: 'duplicate',
            progress: 0,
            existingTitle: (err as any).data?.existingTitle,
          })
          duplicateCount++
        } else {
          const msg = err.message || 'Erro desconhecido'
          updateFile(entry.id, { status: 'error', progress: 0, error: msg })
          errorCount++
        }
      }
    }

    queryClient.invalidateQueries({ queryKey: ['articles'] })
    setIsProcessing(false)

    // Toast de resumo — só mostra se houve algo novo
    if (successCount > 0 && errorCount === 0 && duplicateCount === 0) {
      toast.success(
        successCount === 1
          ? '1 arquivo importado com sucesso!'
          : `${successCount} arquivos importados com sucesso!`,
      )
    } else if (successCount > 0 || errorCount > 0 || duplicateCount > 0) {
      const parts: string[] = []
      if (successCount > 0) parts.push(`${successCount} importado(s)`)
      if (duplicateCount > 0) parts.push(`${duplicateCount} já existia(m)`)
      if (errorCount > 0) parts.push(`${errorCount} com erro`)
      if (errorCount > 0) {
        toast.warning(parts.join(' • '))
      } else {
        toast.info(parts.join(' • '))
      }
    }
  }

  const handleClose = () => {
    if (isProcessing) return
    onOpenChange(false)
    setTimeout(() => setFiles([]), 300)
  }

  // ── estado derivado ────────────────────────────────────────────────────────

  const pendingCount = files.filter(e => e.status === 'pending').length
  const doneCount = files.filter(e => e.status === 'done').length
  const isTerminal = (s: FileStatus) => s === 'done' || s === 'error' || s === 'duplicate'
  const allFinished = files.length > 0 && files.every(e => isTerminal(e.status))

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="max-w-2xl flex flex-col max-h-[90vh]">
        <DialogHeader className="flex-shrink-0">
          <DialogTitle>Importar Artigos e Livros</DialogTitle>
          <DialogDescription>
            Selecione um ou mais arquivos. Cada arquivo é processado de forma
            independente. EPUBs e arquivos grandes podem ser importados como livros.
          </DialogDescription>
        </DialogHeader>

        <div className="flex-1 overflow-hidden flex flex-col gap-4 min-h-0">
          {/* Dropzone — oculto durante processamento */}
          {!isProcessing && (
            <div
              {...getRootProps()}
              className={`
                border-2 border-dashed rounded-lg p-6 text-center cursor-pointer
                transition-colors duration-200 flex-shrink-0
                ${isDragActive
                  ? 'border-primary bg-primary/5'
                  : 'border-border hover:border-primary/50 hover:bg-muted/50'}
              `}
            >
              <input {...getInputProps()} />
              <Upload className="mx-auto h-10 w-10 text-muted-foreground mb-3" />
              <p className="text-sm font-medium mb-1">
                {isDragActive
                  ? 'Solte os arquivos aqui'
                  : 'Arraste um ou mais arquivos — ou clique para selecionar'}
              </p>
              <p className="text-xs text-muted-foreground">
                Múltiplos arquivos • Artigos até 50 MB • Livros até 200 MB • PDF, EPUB, TXT, MD
              </p>
            </div>
          )}

          {/* Lista de arquivos */}
          {files.length > 0 && (
            <ScrollArea className="flex-1 min-h-0 max-h-[50vh]">
              <div className="space-y-2 pr-3">
                {files.map(entry => (
                  <FileEntryRow
                    key={entry.id}
                    entry={entry}
                    isProcessing={isProcessing}
                    onRemove={() => removeFile(entry.id)}
                    onToggleBook={val => toggleBook(entry.id, val)}
                  />
                ))}
              </div>
            </ScrollArea>
          )}

          {/* Dica quando lista está vazia */}
          {files.length === 0 && (
            <div className="bg-muted/50 rounded-lg p-4 flex-shrink-0">
              <h4 className="font-medium text-sm mb-2">O que será extraído automaticamente:</h4>
              <ul className="text-sm text-muted-foreground space-y-1">
                <li>• <strong>DOI e PMID</strong> (se presentes no PDF)</li>
                <li>• <strong>Título</strong> e <strong>Abstract</strong></li>
                <li>• <strong>Texto completo</strong> do documento</li>
                <li>• <strong>Metadados adicionais</strong> via CrossRef API</li>
              </ul>
            </div>
          )}
        </div>

        {/* Rodapé */}
        <div className="flex items-center justify-between gap-3 pt-3 flex-shrink-0 border-t">
          <p className="text-sm text-muted-foreground">
            {files.length === 0
              ? 'Nenhum arquivo selecionado'
              : allFinished
              ? `${doneCount} de ${files.length} concluído(s)`
              : isProcessing
              ? `Processando… ${doneCount} de ${files.length}`
              : `${pendingCount} arquivo(s) para importar`}
          </p>
          <div className="flex gap-2">
            <Button variant="outline" onClick={handleClose} disabled={isProcessing}>
              {allFinished ? 'Fechar' : 'Cancelar'}
            </Button>
            {!allFinished && (
              <Button
                onClick={handleUpload}
                disabled={pendingCount === 0 || isProcessing}
              >
                {isProcessing ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Processando…
                  </>
                ) : (
                  <>
                    <Upload className="mr-2 h-4 w-4" />
                    {pendingCount === 1
                      ? 'Importar 1 arquivo'
                      : `Importar ${pendingCount} arquivos`}
                  </>
                )}
              </Button>
            )}
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}

// ─── linha por arquivo ────────────────────────────────────────────────────────

interface FileEntryRowProps {
  entry: FileEntry
  isProcessing: boolean
  onRemove: () => void
  onToggleBook: (val: boolean) => void
}

function FileEntryRow({ entry, isProcessing, onRemove, onToggleBook }: FileEntryRowProps) {
  const { file, asBook, status, progress, error, chapterCount } = entry
  const showBookToggle = status === 'pending' && !isProcessing && isLikelyBook(file)

  return (
    <div
      className={`
        border rounded-lg p-3 transition-colors
        ${status === 'done' ? 'border-green-200 bg-green-50/50 dark:border-green-800 dark:bg-green-950/20' : ''}
        ${status === 'duplicate' ? 'border-amber-200 bg-amber-50/50 dark:border-amber-800 dark:bg-amber-950/20' : ''}
        ${status === 'error' ? 'border-destructive/30 bg-destructive/5' : ''}
      `}
    >
      <div className="flex items-start gap-3">
        {/* Ícone de status */}
        <div className="flex-shrink-0 mt-0.5">
          {status === 'done' ? (
            <CheckCircle2 className="h-5 w-5 text-green-600" />
          ) : status === 'duplicate' ? (
            <AlertTriangle className="h-5 w-5 text-amber-500" />
          ) : status === 'error' ? (
            <AlertCircle className="h-5 w-5 text-destructive" />
          ) : status === 'uploading' ? (
            <Loader2 className="h-5 w-5 text-primary animate-spin" />
          ) : asBook ? (
            <BookOpen className="h-5 w-5 text-primary" />
          ) : (
            <FileText className="h-5 w-5 text-muted-foreground" />
          )}
        </div>

        {/* Conteúdo */}
        <div className="flex-1 min-w-0">
          <div className="flex items-start justify-between gap-2">
            <div className="min-w-0">
              <p className="text-sm font-medium truncate">{file.name}</p>
              <p className="text-xs text-muted-foreground">
                {formatFileSize(file.size)}
                {status === 'done' && chapterCount && (
                  <span className="ml-2 text-green-600">• {chapterCount} capítulos</span>
                )}
                {status === 'done' && !chapterCount && (
                  <span className="ml-2 text-green-600">• importado</span>
                )}
              </p>
            </div>
            {status === 'pending' && !isProcessing && (
              <Button
                variant="ghost"
                size="icon"
                className="h-6 w-6 flex-shrink-0"
                onClick={onRemove}
              >
                <X className="h-3.5 w-3.5" />
              </Button>
            )}
          </div>

          {/* Toggle livro */}
          {showBookToggle && (
            <div className="mt-2 flex items-center gap-2">
              <Switch
                id={`book-${entry.id}`}
                checked={asBook}
                onCheckedChange={onToggleBook}
                className="scale-90"
              />
              <Label
                htmlFor={`book-${entry.id}`}
                className="text-xs text-muted-foreground cursor-pointer"
              >
                Importar como livro (divide em capítulos)
              </Label>
            </div>
          )}

          {/* Barra de progresso */}
          {status === 'uploading' && (
            <div className="mt-2">
              <Progress value={progress} className="h-1.5" />
              <p className="text-xs text-muted-foreground mt-1">
                {asBook ? 'Processando capítulos…' : `Enviando… ${progress}%`}
              </p>
            </div>
          )}

          {/* Aviso de duplicata */}
          {status === 'duplicate' && (
            <p className="text-xs text-amber-700 dark:text-amber-400 mt-1">
              Já importado
              {entry.existingTitle ? ` como "${entry.existingTitle}"` : ' anteriormente'}
            </p>
          )}

          {/* Mensagem de erro real */}
          {status === 'error' && error && (
            <p className="text-xs text-destructive mt-1 truncate" title={error}>
              {error}
            </p>
          )}
        </div>
      </div>
    </div>
  )
}
