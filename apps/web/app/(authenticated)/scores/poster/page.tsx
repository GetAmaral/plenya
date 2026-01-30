'use client'

import { useAllScoreGroupTrees } from '@/lib/api/score-api'
import { Button } from '@/components/ui/button'
import { Printer, Loader2, Download, ZoomIn, ZoomOut, Maximize2 } from 'lucide-react'
import { useEffect, useState } from 'react'

const LEVEL_STYLES = {
  0: {
    bg: 'bg-red-100',
    text: 'text-red-900',
    border: 'border-red-500',
  },
  1: {
    bg: 'bg-orange-100',
    text: 'text-orange-900',
    border: 'border-orange-500',
  },
  2: {
    bg: 'bg-yellow-100',
    text: 'text-yellow-900',
    border: 'border-yellow-500',
  },
  3: {
    bg: 'bg-blue-100',
    text: 'text-blue-900',
    border: 'border-blue-500',
  },
  4: {
    bg: 'bg-green-100',
    text: 'text-green-900',
    border: 'border-green-500',
  },
  5: {
    bg: 'bg-emerald-100',
    text: 'text-emerald-900',
    border: 'border-emerald-500',
  },
  6: {
    bg: 'bg-gray-100',
    text: 'text-gray-900',
    border: 'border-gray-500',
  },
}

function formatLevelRange(level: any) {
  if (level.operator === 'between' && level.lowerLimit && level.upperLimit) {
    return `${level.lowerLimit} - ${level.upperLimit}`
  }
  if (level.operator === '=') {
    return `= ${level.lowerLimit || level.upperLimit || ''}`
  }
  if (level.operator === '>') {
    return `> ${level.lowerLimit || level.upperLimit || ''}`
  }
  if (level.operator === '>=') {
    return `≥ ${level.lowerLimit || level.upperLimit || ''}`
  }
  if (level.operator === '<') {
    return `< ${level.upperLimit || level.lowerLimit || ''}`
  }
  if (level.operator === '<=') {
    return `≤ ${level.upperLimit || level.lowerLimit || ''}`
  }
  return ''
}

export default function ScorePosterPage() {
  const { data: groups = [], isLoading } = useAllScoreGroupTrees()
  const [zoom, setZoom] = useState(100)

  useEffect(() => {
    document.title = 'Escore Plenya - Pôster 60x300cm'
  }, [])

  const handlePrint = () => {
    window.print()
  }

  const handleSavePDF = () => {
    window.print()
  }

  const handleZoomIn = () => {
    setZoom(prev => Math.min(prev + 10, 200))
  }

  const handleZoomOut = () => {
    setZoom(prev => Math.max(prev - 10, 50))
  }

  const handleFitView = () => {
    setZoom(100)
  }

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <Loader2 className="h-12 w-12 animate-spin text-primary" />
          <p className="text-lg text-muted-foreground">Carregando escore...</p>
        </div>
      </div>
    )
  }

  return (
    <>
      <style jsx global>{`
        /* Configurações para impressão em papel 60x300cm */
        @page {
          size: 3000mm 600mm; /* 300cm x 60cm */
          margin: 0;
        }

        @media print {
          body {
            margin: 0;
            padding: 0;
            overflow: visible;
          }

          .print-container {
            width: 3000mm !important;
            height: 600mm !important;
            transform: none !important;
            overflow: visible !important;
          }

          .no-print {
            display: none !important;
          }
        }

        /* Configurações de visualização na tela */
        .screen-preview {
          transform: scale(${zoom / 100});
          transform-origin: top left;
          transition: transform 0.2s;
        }
      `}</style>

      {/* Cabeçalho com controles (oculto na impressão) */}
      <div className="no-print fixed top-0 left-0 right-0 z-50 bg-background border-b shadow-md">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold">Escore Plenya - Pôster 60x300cm</h1>
            <p className="text-sm text-muted-foreground">
              Otimizado para impressão em papel grande (banner horizontal)
            </p>
          </div>
          <div className="flex items-center gap-2">
            <div className="flex items-center gap-1 border rounded-lg px-2 py-1">
              <Button onClick={handleZoomOut} size="sm" variant="ghost">
                <ZoomOut className="h-4 w-4" />
              </Button>
              <span className="text-sm font-medium min-w-[60px] text-center">
                {zoom}%
              </span>
              <Button onClick={handleZoomIn} size="sm" variant="ghost">
                <ZoomIn className="h-4 w-4" />
              </Button>
              <Button onClick={handleFitView} size="sm" variant="ghost">
                <Maximize2 className="h-4 w-4" />
              </Button>
            </div>
            <Button onClick={handleSavePDF} size="lg" variant="default">
              <Download className="mr-2 h-5 w-5" />
              Salvar PDF
            </Button>
            <Button onClick={handlePrint} size="lg" variant="outline">
              <Printer className="mr-2 h-5 w-5" />
              Imprimir
            </Button>
          </div>
        </div>
      </div>

      {/* Container de visualização (com scroll) */}
      <div className="no-print pt-20 overflow-auto bg-gray-100 min-h-screen p-8">
        <div className="screen-preview">
          <PosterContent groups={groups} />
        </div>
      </div>

      {/* Versão para impressão (renderizada fora da tela) */}
      <div className="hidden print:block">
        <PosterContent groups={groups} />
      </div>
    </>
  )
}

function PosterContent({ groups }: { groups: any[] }) {
  return (
    <div className="print-container bg-white" style={{ width: '3000mm', height: '600mm' }}>
      {/* Header do pôster */}
      <div className="border-b-8 border-primary bg-gradient-to-r from-primary to-primary/80 text-primary-foreground px-16 py-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-[80px] font-black tracking-tight leading-none">
              ESCORE PLENYA
            </h1>
            <p className="text-[32px] font-semibold mt-2 opacity-90">
              Sistema Completo de Pontuação de Saúde e Longevidade
            </p>
          </div>
          <div className="text-right">
            <div className="text-[24px] font-medium">
              {new Date().toLocaleDateString('pt-BR', {
                day: '2-digit',
                month: 'long',
                year: 'numeric'
              })}
            </div>
            <div className="text-[20px] opacity-80 mt-1">
              {groups.length} Grupos · {groups.reduce((acc, g) => acc + (g.subgroups?.length || 0), 0)} Subgrupos
            </div>
          </div>
        </div>
      </div>

      {/* Conteúdo em layout horizontal (colunas) */}
      <div className="grid grid-cols-6 gap-6 p-8" style={{ height: 'calc(600mm - 180px)' }}>
        {groups.map((group, groupIndex) => (
          <div key={group.id} className="flex flex-col h-full">
            {/* Cabeçalho do Grupo */}
            <div className="bg-primary text-primary-foreground rounded-2xl px-6 py-4 mb-4 shadow-lg">
              <div className="flex items-center gap-3">
                <div className="flex items-center justify-center w-12 h-12 rounded-full bg-primary-foreground/20 text-2xl font-black">
                  {groupIndex + 1}
                </div>
                <h2 className="text-[28px] font-black leading-tight">{group.name}</h2>
              </div>
            </div>

            {/* Subgrupos em scroll vertical */}
            <div className="flex-1 overflow-visible space-y-4">
              {group.subgroups && group.subgroups.length > 0 && (
                <>
                  {group.subgroups.map((subgroup: any, subgroupIndex: number) => (
                    <div key={subgroup.id} className="bg-gradient-to-br from-muted to-muted/50 rounded-xl p-4 border-2 border-primary/20 shadow-md">
                      {/* Cabeçalho do Subgrupo */}
                      <div className="mb-3 pb-2 border-b-2 border-primary/30">
                        <div className="flex items-center gap-2">
                          <div className="flex items-center justify-center w-8 h-8 rounded-full bg-primary/10 text-[16px] font-black text-primary">
                            {groupIndex + 1}.{subgroupIndex + 1}
                          </div>
                          <h3 className="text-[20px] font-bold">{subgroup.name}</h3>
                        </div>
                      </div>

                      {/* Itens */}
                      {subgroup.items && subgroup.items.length > 0 && (
                        <div className="space-y-3">
                          {subgroup.items
                            .filter((item: any) => !item.parentItemId)
                            .map((item: any) => (
                              <div key={item.id} className="border-2 rounded-lg overflow-hidden bg-white shadow-sm">
                                {/* Cabeçalho do Item */}
                                <div className="bg-card px-3 py-2 border-b-2 border-muted">
                                  <div className="flex items-start justify-between gap-2">
                                    <div className="flex-1">
                                      <h4 className="font-bold text-[16px] leading-tight">
                                        {item.name}
                                      </h4>
                                      {item.unit && (
                                        <span className="text-[13px] text-muted-foreground mt-1 block">
                                          {item.unit}{item.unitConversion && ` | ${item.unitConversion}`}
                                        </span>
                                      )}
                                    </div>
                                    {item.points > 0 && (
                                      <div className="shrink-0 bg-primary/10 rounded-full px-3 py-1">
                                        <div className="text-[16px] font-black text-primary">
                                          {item.points}pt
                                        </div>
                                      </div>
                                    )}
                                  </div>
                                </div>

                                {/* Níveis */}
                                {item.levels && item.levels.length > 0 && (
                                  <div className="p-3 bg-muted/20">
                                    <div className="flex flex-wrap gap-2">
                                      {item.levels
                                        .sort((a: any, b: any) => a.level - b.level)
                                        .map((level: any) => {
                                          const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                          const range = formatLevelRange(level)
                                          const hasValues = level.lowerLimit != null || level.upperLimit != null

                                          return (
                                            <div
                                              key={level.id}
                                              className={`inline-flex items-center gap-1.5 rounded-full border-2 px-3 py-1 text-[14px] font-bold ${style.bg} ${style.text} ${style.border} shadow-sm`}
                                            >
                                              <span className="font-black">N{level.level}:</span>
                                              {hasValues ? (
                                                <span className="font-mono text-[13px]">{range}</span>
                                              ) : (
                                                <span className="text-[13px]">{level.name}</span>
                                              )}
                                            </div>
                                          )
                                        })}
                                    </div>
                                  </div>
                                )}

                                {/* Itens Filhos */}
                                {item.childItems && item.childItems.length > 0 && (
                                  <div className="border-t-2 bg-muted/10">
                                    <div className="p-2 space-y-2">
                                      {item.childItems.map((childItem: any) => (
                                        <div key={childItem.id} className="pl-3 border-l-4 border-primary/30">
                                          <div className="flex items-start justify-between gap-2 mb-1">
                                            <h5 className="font-semibold text-[14px] text-muted-foreground">
                                              {childItem.name}
                                              {childItem.unit && (
                                                <span className="text-[12px] text-muted-foreground/80 ml-1.5">
                                                  ({childItem.unit})
                                                </span>
                                              )}
                                            </h5>
                                            {childItem.points > 0 && (
                                              <span className="text-[14px] font-bold text-primary shrink-0">
                                                {childItem.points}pt
                                              </span>
                                            )}
                                          </div>

                                          {/* Níveis do item filho */}
                                          {childItem.levels && childItem.levels.length > 0 && (
                                            <div className="flex flex-wrap gap-1.5">
                                              {childItem.levels
                                                .sort((a: any, b: any) => a.level - b.level)
                                                .map((level: any) => {
                                                  const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                                  const range = formatLevelRange(level)
                                                  const hasValues = level.lowerLimit != null || level.upperLimit != null

                                                  return (
                                                    <div
                                                      key={level.id}
                                                      className={`inline-flex items-center gap-1 rounded-full border-2 px-2.5 py-0.5 text-[13px] font-bold ${style.bg} ${style.text} ${style.border}`}
                                                    >
                                                      <span className="font-black">N{level.level}:</span>
                                                      {hasValues ? (
                                                        <span className="font-mono text-[12px]">{range}</span>
                                                      ) : (
                                                        <span className="text-[12px]">{level.name}</span>
                                                      )}
                                                    </div>
                                                  )
                                                })}
                                            </div>
                                          )}
                                        </div>
                                      ))}
                                    </div>
                                  </div>
                                )}
                              </div>
                            ))}
                        </div>
                      )}
                    </div>
                  ))}
                </>
              )}
            </div>
          </div>
        ))}
      </div>

      {/* Rodapé */}
      <div className="absolute bottom-0 left-0 right-0 border-t-4 border-primary bg-gradient-to-r from-primary/5 to-primary/10 px-16 py-6">
        <div className="flex items-center justify-between">
          <div>
            <p className="text-[24px] font-bold">Plenya - Sistema de Prontuário Médico Eletrônico</p>
            <p className="text-[16px] text-muted-foreground mt-1">
              Documento confidencial - LGPD aplicável
            </p>
          </div>
          <div className="text-right">
            <div className="text-[20px] font-semibold">
              plenya.com.br
            </div>
            <div className="text-[14px] text-muted-foreground mt-1">
              © {new Date().getFullYear()} Todos os direitos reservados
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
