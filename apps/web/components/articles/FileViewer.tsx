'use client'

import dynamic from 'next/dynamic'
import { Skeleton } from '@/components/ui/skeleton'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Download, FileText } from 'lucide-react'

const PDFViewer = dynamic(
  () => import('@/components/articles/PDFViewer').then((mod) => mod.PDFViewer),
  {
    ssr: false,
    loading: () => (
      <Card className="w-full">
        <CardContent className="py-12">
          <Skeleton className="h-[600px] w-full" />
        </CardContent>
      </Card>
    ),
  }
)

const EpubViewer = dynamic(
  () => import('@/components/articles/EpubViewer').then((mod) => mod.EpubViewer),
  {
    ssr: false,
    loading: () => (
      <Card className="w-full">
        <CardContent className="py-12">
          <Skeleton className="h-[700px] w-full" />
        </CardContent>
      </Card>
    ),
  }
)

interface FileViewerProps {
  internalLink: string
  downloadUrl: string
  title?: string
  fullContent?: string
}

function getExtension(link: string): string {
  const lower = link.toLowerCase()
  if (lower.endsWith('.epub')) return 'epub'
  if (lower.endsWith('.pdf')) return 'pdf'
  if (lower.endsWith('.md')) return 'md'
  if (lower.endsWith('.txt')) return 'txt'
  return 'unknown'
}

export function FileViewer({ internalLink, downloadUrl, title, fullContent }: FileViewerProps) {
  const ext = getExtension(internalLink)

  if (ext === 'pdf') {
    return <PDFViewer url={downloadUrl} title={title} />
  }

  if (ext === 'epub') {
    return <EpubViewer url={downloadUrl} title={title} />
  }

  if ((ext === 'txt' || ext === 'md') && fullContent) {
    // TXT/MD já são exibidos na seção "Texto Completo" da página —
    // aqui só mostramos o botão de download para não duplicar o conteúdo.
    return (
      <Card className="w-full">
        <CardHeader>
          <CardTitle className="text-lg flex items-center gap-2">
            <FileText className="h-5 w-5" />
            {title || 'Arquivo de Texto'}
          </CardTitle>
        </CardHeader>
        <CardContent>
          <p className="text-sm text-muted-foreground mb-4">
            O conteúdo completo está disponível na seção &quot;Texto Completo&quot; abaixo.
          </p>
          <Button variant="outline" asChild>
            <a href={downloadUrl} download>
              <Download className="mr-2 h-4 w-4" />
              Download {ext.toUpperCase()}
            </a>
          </Button>
        </CardContent>
      </Card>
    )
  }

  return null
}
