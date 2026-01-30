'use client'

import { useState, useEffect, useMemo } from 'react'
import { Document, Page, pdfjs } from 'react-pdf'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import {
  ChevronLeft,
  ChevronRight,
  ZoomIn,
  ZoomOut,
  Download,
  Maximize2,
  Loader2,
} from 'lucide-react'
import { useAuthStore } from '@/lib/auth-store'
import 'react-pdf/dist/Page/AnnotationLayer.css'
import 'react-pdf/dist/Page/TextLayer.css'

interface PDFViewerProps {
  url: string
  title?: string
}

export function PDFViewer({ url, title }: PDFViewerProps) {
  const [numPages, setNumPages] = useState<number>(0)
  const [pageNumber, setPageNumber] = useState<number>(1)
  const [scale, setScale] = useState<number>(1.0)
  const [isLoading, setIsLoading] = useState(true)
  const accessToken = useAuthStore((state) => state.accessToken)

  // Configure PDF.js worker on client side only - use legacy build for Node.js compatibility
  useEffect(() => {
    pdfjs.GlobalWorkerOptions.workerSrc = `//unpkg.com/pdfjs-dist@${pdfjs.version}/legacy/build/pdf.worker.min.mjs`
  }, [])

  // Memoize the file config to prevent unnecessary reloads
  const fileConfig = useMemo(
    () => ({
      url,
      httpHeaders: accessToken
        ? { Authorization: `Bearer ${accessToken}` }
        : undefined,
    }),
    [url, accessToken]
  )

  function onDocumentLoadSuccess({ numPages }: { numPages: number }) {
    setNumPages(numPages)
    setIsLoading(false)
  }

  function onDocumentLoadError(error: Error) {
    console.error('Error loading PDF:', error)
    setIsLoading(false)
  }

  const changePage = (offset: number) => {
    setPageNumber((prevPageNumber) => {
      const newPage = prevPageNumber + offset
      return Math.min(Math.max(1, newPage), numPages)
    })
  }

  const previousPage = () => changePage(-1)
  const nextPage = () => changePage(1)

  const zoomIn = () => setScale((prev) => Math.min(prev + 0.2, 3.0))
  const zoomOut = () => setScale((prev) => Math.max(prev - 0.2, 0.5))

  const handleDownload = () => {
    window.open(url, '_blank')
  }

  return (
    <Card className="w-full">
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle className="text-lg">
            {title || 'Visualizador de PDF'}
          </CardTitle>

          <div className="flex items-center gap-2">
            <Button variant="outline" size="sm" onClick={zoomOut}>
              <ZoomOut className="h-4 w-4" />
            </Button>
            <span className="text-sm text-muted-foreground min-w-[60px] text-center">
              {Math.round(scale * 100)}%
            </span>
            <Button variant="outline" size="sm" onClick={zoomIn}>
              <ZoomIn className="h-4 w-4" />
            </Button>

            <div className="h-4 w-px bg-border mx-2" />

            <Button variant="outline" size="sm" onClick={handleDownload}>
              <Download className="mr-2 h-4 w-4" />
              Download
            </Button>
          </div>
        </div>
      </CardHeader>

      <CardContent>
        {/* PDF Document */}
        <div className="border rounded-lg bg-muted/30 overflow-hidden">
          <div className="flex items-center justify-center min-h-[600px] relative">
            {isLoading && (
              <div className="absolute inset-0 flex items-center justify-center bg-background/80 z-10">
                <div className="flex flex-col items-center gap-2">
                  <Loader2 className="h-8 w-8 animate-spin text-primary" />
                  <p className="text-sm text-muted-foreground">
                    Carregando PDF...
                  </p>
                </div>
              </div>
            )}

            <div className="overflow-auto max-h-[700px] w-full flex justify-center p-4">
              <Document
                file={fileConfig}
                onLoadSuccess={onDocumentLoadSuccess}
                onLoadError={onDocumentLoadError}
                loading={
                  <div className="flex items-center justify-center py-12">
                    <Loader2 className="h-8 w-8 animate-spin text-primary" />
                  </div>
                }
                error={
                  <div className="text-center py-12">
                    <p className="text-destructive mb-2">
                      Erro ao carregar PDF
                    </p>
                    <p className="text-sm text-muted-foreground">
                      Verifique se o arquivo existe e tente novamente.
                    </p>
                  </div>
                }
              >
                <Page
                  pageNumber={pageNumber}
                  scale={scale}
                  renderTextLayer={true}
                  renderAnnotationLayer={true}
                  className="shadow-lg"
                />
              </Document>
            </div>
          </div>
        </div>

        {/* Navigation */}
        {!isLoading && numPages > 0 && (
          <div className="flex items-center justify-between mt-4">
            <Button
              variant="outline"
              size="sm"
              onClick={previousPage}
              disabled={pageNumber <= 1}
            >
              <ChevronLeft className="mr-2 h-4 w-4" />
              Anterior
            </Button>

            <div className="flex items-center gap-2">
              <span className="text-sm text-muted-foreground">
                Página{' '}
                <input
                  type="number"
                  min={1}
                  max={numPages}
                  value={pageNumber}
                  onChange={(e) => {
                    const page = parseInt(e.target.value)
                    if (page >= 1 && page <= numPages) {
                      setPageNumber(page)
                    }
                  }}
                  className="w-16 px-2 py-1 text-center border rounded mx-1"
                />{' '}
                de {numPages}
              </span>
            </div>

            <Button
              variant="outline"
              size="sm"
              onClick={nextPage}
              disabled={pageNumber >= numPages}
            >
              Próxima
              <ChevronRight className="ml-2 h-4 w-4" />
            </Button>
          </div>
        )}
      </CardContent>
    </Card>
  )
}
