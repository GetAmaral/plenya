'use client'

import { useState, useEffect, useRef, useCallback } from 'react'
import type { Book, Rendition, NavItem, Location } from 'epubjs'
import { useAuthStore } from '@/lib/auth-store'
import { Card, CardContent } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Download,
  AlertCircle,
  ChevronLeft,
  ChevronRight,
  BookOpen,
  List,
  ZoomIn,
  ZoomOut,
} from 'lucide-react'
import { cn } from '@/lib/utils'

interface EpubViewerProps {
  url: string
  title?: string
}

const FONT_SIZES = [80, 90, 100, 110, 120, 140, 160]
const DEFAULT_FONT_IDX = 2 // 100%

export function EpubViewer({ url, title }: EpubViewerProps) {
  const viewerRef = useRef<HTMLDivElement>(null)
  const bookRef = useRef<Book | null>(null)
  const renditionRef = useRef<Rendition | null>(null)
  const blobUrlRef = useRef<string | null>(null)

  // Fetch state
  const [blobUrl, setBlobUrl] = useState<string | null>(null)
  const [fetchLoading, setFetchLoading] = useState(true)
  const [fetchError, setFetchError] = useState<string | null>(null)

  // Reader state
  const [isLoaded, setIsLoaded] = useState(false)
  const [loadError, setLoadError] = useState(false)
  const [toc, setToc] = useState<NavItem[]>([])
  const [showToc, setShowToc] = useState(false)

  // Pagination
  const [currentPage, setCurrentPage] = useState(1)
  const [totalPages, setTotalPages] = useState(0)
  const [pageInput, setPageInput] = useState('')

  // Zoom
  const [fontIdx, setFontIdx] = useState(DEFAULT_FONT_IDX)

  const accessToken = useAuthStore((s) => s.accessToken)

  // Step 1: Fetch EPUB with auth and create blob URL
  useEffect(() => {
    setFetchLoading(true)
    setFetchError(null)
    setBlobUrl(null)
    setIsLoaded(false)
    setLoadError(false)

    if (blobUrlRef.current) {
      URL.revokeObjectURL(blobUrlRef.current)
      blobUrlRef.current = null
    }

    const controller = new AbortController()

    fetch(url, {
      headers: accessToken ? { Authorization: `Bearer ${accessToken}` } : {},
      signal: controller.signal,
    })
      .then((res) => {
        if (!res.ok) throw new Error(`Erro ${res.status}`)
        return res.blob()
      })
      .then((blob) => {
        const objUrl = URL.createObjectURL(blob)
        blobUrlRef.current = objUrl
        setBlobUrl(objUrl)
        setFetchLoading(false)
      })
      .catch((err) => {
        if (err.name === 'AbortError') return
        setFetchError(err.message || 'Erro ao carregar EPUB')
        setFetchLoading(false)
      })

    return () => {
      controller.abort()
      if (blobUrlRef.current) {
        URL.revokeObjectURL(blobUrlRef.current)
        blobUrlRef.current = null
      }
    }
  }, [url, accessToken])

  // Step 2: Initialize epub.js when blob URL and viewer are ready
  useEffect(() => {
    if (!blobUrl || !viewerRef.current) return

    if (bookRef.current) {
      bookRef.current.destroy()
      bookRef.current = null
      renditionRef.current = null
    }

    setIsLoaded(false)
    setLoadError(false)
    setToc([])
    setCurrentPage(1)
    setTotalPages(0)
    setFontIdx(DEFAULT_FONT_IDX)

    let cancelled = false

    import('epubjs').then(({ default: Epub }) => {
      if (cancelled || !viewerRef.current) return

      const book = Epub(blobUrl, { openAs: 'epub' }) as Book
      bookRef.current = book

      const rendition = book.renderTo(viewerRef.current, {
        width: '100%',
        height: '100%',
        spread: 'none',
        allowScriptedContent: true,
      } as any)
      renditionRef.current = rendition

      // Track page changes
      rendition.on('relocated', (location: Location) => {
        if (cancelled || bookRef.current !== book) return
        const page = location.start.location + 1
        setCurrentPage(page > 0 ? page : 1)
      })

      book.loaded.navigation
        .then(({ toc: navItems }) => {
          if (cancelled || bookRef.current !== book) return
          setToc(navItems)
          setIsLoaded(true)
          rendition.display()
        })
        .then(() => {
          if (cancelled || bookRef.current !== book) return
          // Generate locations for page X of Y (1600 chars per "page")
          return book.locations.generate(1600)
        })
        .then(() => {
          if (cancelled || bookRef.current !== book) return
          setTotalPages(book.locations.total)
        })
        .catch(() => {
          if (cancelled || bookRef.current !== book) return
          setLoadError(true)
        })
    })

    return () => {
      cancelled = true
      if (bookRef.current) {
        bookRef.current.destroy()
        bookRef.current = null
        renditionRef.current = null
      }
    }
  }, [blobUrl])

  // Apply font size when it changes
  useEffect(() => {
    if (!renditionRef.current || !isLoaded) return
    renditionRef.current.themes.fontSize(`${FONT_SIZES[fontIdx]}%`)
  }, [fontIdx, isLoaded])

  const handlePrev = useCallback(() => renditionRef.current?.prev(), [])
  const handleNext = useCallback(() => renditionRef.current?.next(), [])

  const handleTocClick = useCallback((href: string) => {
    renditionRef.current?.display(href)
    setShowToc(false)
  }, [])

  const handleZoomIn = useCallback(() => {
    setFontIdx((i) => Math.min(i + 1, FONT_SIZES.length - 1))
  }, [])

  const handleZoomOut = useCallback(() => {
    setFontIdx((i) => Math.max(i - 1, 0))
  }, [])

  const handlePageInputKeyDown = useCallback(
    (e: React.KeyboardEvent<HTMLInputElement>) => {
      if (e.key !== 'Enter') return
      const n = parseInt(pageInput)
      if (!isNaN(n) && n >= 1 && n <= totalPages && bookRef.current) {
        const cfi = bookRef.current.locations.cfiFromLocation(n - 1)
        if (cfi) renditionRef.current?.display(cfi)
      }
      setPageInput('')
    },
    [pageInput, totalPages]
  )

  // ── Loading ──
  if (fetchLoading) {
    return (
      <Card className="w-full">
        <CardContent className="py-12">
          <Skeleton className="h-[700px] w-full" />
        </CardContent>
      </Card>
    )
  }

  // ── Fetch error ──
  if (fetchError) {
    return (
      <Card className="w-full">
        <CardContent className="py-8 flex flex-col items-center gap-4">
          <AlertCircle className="h-10 w-10 text-destructive" />
          <p className="text-sm text-destructive text-center">{fetchError}</p>
          <Button variant="outline" asChild>
            <a href={url} download>
              <Download className="mr-2 h-4 w-4" />
              Download EPUB
            </a>
          </Button>
        </CardContent>
      </Card>
    )
  }

  return (
    <div className="w-full rounded-lg border overflow-hidden flex flex-col" style={{ height: '85vh' }}>
      {/* ── Top toolbar ── */}
      <div className="flex items-center gap-2 px-3 py-2 border-b bg-muted/30 shrink-0 flex-wrap">
        {/* TOC toggle */}
        <Button
          variant={showToc ? 'secondary' : 'outline'}
          size="icon"
          className="h-8 w-8 shrink-0"
          onClick={() => setShowToc((v) => !v)}
          title="Índice de capítulos"
        >
          <List className="h-4 w-4" />
        </Button>

        {/* Title */}
        <div className="flex items-center gap-1.5 text-sm text-muted-foreground min-w-0 flex-1">
          <BookOpen className="h-4 w-4 shrink-0" />
          <span className="font-medium truncate">{title || 'EPUB'}</span>
        </div>

        {/* Prev / Next */}
        <div className="flex items-center gap-1 shrink-0">
          <Button variant="outline" size="icon" className="h-8 w-8" onClick={handlePrev} title="Página anterior">
            <ChevronLeft className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="icon" className="h-8 w-8" onClick={handleNext} title="Próxima página">
            <ChevronRight className="h-4 w-4" />
          </Button>
        </div>

        {/* Page X of Y */}
        {totalPages > 0 && (
          <div className="flex items-center gap-1.5 text-sm shrink-0">
            <span className="text-muted-foreground">Página</span>
            <Input
              className="h-8 w-14 text-center text-sm px-1"
              value={pageInput !== '' ? pageInput : currentPage}
              onFocus={() => setPageInput(String(currentPage))}
              onBlur={() => setPageInput('')}
              onChange={(e) => setPageInput(e.target.value)}
              onKeyDown={handlePageInputKeyDown}
              title="Digite a página e pressione Enter"
            />
            <span className="text-muted-foreground">de {totalPages}</span>
          </div>
        )}

        {/* Zoom */}
        <div className="flex items-center gap-1 shrink-0">
          <Button variant="outline" size="icon" className="h-8 w-8" onClick={handleZoomOut} disabled={fontIdx === 0} title="Diminuir texto">
            <ZoomOut className="h-4 w-4" />
          </Button>
          <span className="text-xs text-muted-foreground w-10 text-center">
            {FONT_SIZES[fontIdx]}%
          </span>
          <Button variant="outline" size="icon" className="h-8 w-8" onClick={handleZoomIn} disabled={fontIdx === FONT_SIZES.length - 1} title="Aumentar texto">
            <ZoomIn className="h-4 w-4" />
          </Button>
        </div>

        {/* Download */}
        <Button variant="outline" size="sm" asChild className="h-8 shrink-0">
          <a href={url} download>
            <Download className="mr-1 h-3 w-3" />
            Download
          </a>
        </Button>
      </div>

      {/* ── Content area (TOC sidebar + reader) ── */}
      <div className="flex flex-1 min-h-0">
        {/* TOC Sidebar */}
        {showToc && (
          <div className="w-64 border-r bg-muted/10 flex flex-col shrink-0 overflow-hidden">
            <div className="px-3 py-2 border-b text-xs font-semibold text-muted-foreground uppercase tracking-wide">
              Capítulos
            </div>
            <div className="flex-1 overflow-y-auto">
              <TocList items={toc} onSelect={handleTocClick} depth={0} />
            </div>
          </div>
        )}

        {/* Viewer */}
        <div className="flex-1 min-w-0 relative">
          {loadError && (
            <div className="absolute inset-0 flex flex-col items-center justify-center gap-4">
              <AlertCircle className="h-10 w-10 text-destructive" />
              <p className="text-sm text-destructive">Não foi possível renderizar o livro.</p>
              <Button variant="outline" asChild>
                <a href={url} download>
                  <Download className="mr-2 h-4 w-4" />
                  Download EPUB
                </a>
              </Button>
            </div>
          )}
          <div
            ref={viewerRef}
            className="w-full h-full bg-white"
            style={{ display: loadError ? 'none' : 'block' }}
          />
        </div>
      </div>
    </div>
  )
}

function TocList({
  items,
  onSelect,
  depth,
}: {
  items: NavItem[]
  onSelect: (href: string) => void
  depth: number
}) {
  return (
    <>
      {items.map((item, idx) => (
        <div key={idx}>
          <button
            onClick={() => onSelect(item.href)}
            className={cn(
              'w-full text-left px-3 py-1.5 text-sm hover:bg-muted/60 transition-colors truncate',
              depth > 0 && 'text-muted-foreground'
            )}
            style={{ paddingLeft: `${12 + depth * 12}px` }}
            title={item.label}
          >
            {item.label}
          </button>
          {item.subitems && item.subitems.length > 0 && (
            <TocList items={item.subitems} onSelect={onSelect} depth={depth + 1} />
          )}
        </div>
      ))}
    </>
  )
}
