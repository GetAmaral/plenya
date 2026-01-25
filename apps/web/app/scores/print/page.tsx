'use client'

import { useAllScoreGroupTrees } from '@/lib/api/score-api'
import { Button } from '@/components/ui/button'
import { Printer, Loader2 } from 'lucide-react'
import { useEffect } from 'react'

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

export default function ScorePrintPage() {
  const { data: groups = [], isLoading } = useAllScoreGroupTrees()

  useEffect(() => {
    // Atualizar título da página
    document.title = 'Escore Plenya - Versão para Impressão'
  }, [])

  const handlePrint = () => {
    window.print()
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
    <div className="min-h-screen bg-background">
      {/* Cabeçalho com botão de impressão (oculto na impressão) */}
      <div className="print:hidden sticky top-0 z-50 bg-background border-b">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-bold">Escore Plenya</h1>
            <p className="text-sm text-muted-foreground">Versão para Impressão</p>
          </div>
          <Button onClick={handlePrint} size="lg">
            <Printer className="mr-2 h-5 w-5" />
            Imprimir / Salvar PDF
          </Button>
        </div>
      </div>

      {/* Cabeçalho da impressão (visível apenas na impressão) */}
      <div className="hidden print:block mb-1">
        <div className="text-center border-b pb-0.5">
          <h1 className="text-sm font-bold">Escore Plenya - Sistema de Pontuação de Saúde</h1>
          <p className="text-[7pt] text-muted-foreground">
            Gerado em: {new Date().toLocaleDateString('pt-BR', { day: '2-digit', month: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit' })}
          </p>
        </div>
      </div>

      {/* Conteúdo Principal */}
      <div className="container mx-auto px-4 py-8 print:px-0 print:py-0">
        <div className="space-y-2 print:space-y-1">
          {groups.map((group, groupIndex) => (
            <div key={group.id}>
              {/* Grupo */}
              <div className="mb-1 print:mb-0.5">
                <div className="bg-primary text-primary-foreground rounded-lg px-3 py-1.5 print:rounded-none print:px-2 print:py-1">
                  <div className="flex items-center gap-2">
                    <div className="flex items-center justify-center w-5 h-5 rounded-full bg-primary-foreground/20 text-[10px] font-bold print:w-4 print:h-4">
                      {groupIndex + 1}
                    </div>
                    <h2 className="text-base font-bold print:text-sm">{group.name}</h2>
                  </div>
                </div>
              </div>

              {/* Subgrupos */}
              {group.subgroups && group.subgroups.length > 0 && (
                <div className="space-y-3 ml-2 print:space-y-1 print:ml-1">
                  {group.subgroups.map((subgroup, subgroupIndex) => (
                    <div key={subgroup.id}>
                      {/* Subgrupo */}
                      <div className="mb-1 print:mb-0.5">
                        <div className="bg-muted rounded-lg px-3 py-1.5 print:rounded-none border-l-2 border-primary print:px-2 print:py-1">
                          <div className="flex items-center gap-1.5">
                            <div className="flex items-center justify-center w-4 h-4 rounded-full bg-primary/10 text-[9px] font-semibold text-primary print:w-3 print:h-3">
                              {groupIndex + 1}.{subgroupIndex + 1}
                            </div>
                            <h3 className="text-sm font-semibold print:text-xs">{subgroup.name}</h3>
                          </div>
                        </div>
                      </div>

                      {/* Itens */}
                      {subgroup.items && subgroup.items.length > 0 && (
                        <div className="space-y-2 ml-2 print:space-y-1 print:ml-1">
                          {subgroup.items
                            .filter(item => !item.parentItemId) // Apenas itens raiz
                            .map((item) => (
                              <div key={item.id} className="break-inside-avoid">
                                {/* Item */}
                                <div className="border rounded-lg print:rounded-none print:border-gray-300">
                                  <div className="bg-card px-2 py-1.5 border-b print:px-1.5 print:py-1">
                                    <div className="flex items-start justify-between gap-2">
                                      <div className="flex-1">
                                        <h4 className="font-medium text-xs leading-tight print:text-[9pt]">
                                          {item.name}
                                          {item.unit && (
                                            <span className="text-[10px] text-muted-foreground ml-1.5 print:text-[7pt]">
                                              ({item.unit}{item.unitConversion && ` | ${item.unitConversion}`})
                                            </span>
                                          )}
                                        </h4>
                                      </div>
                                      {item.points > 0 && (
                                        <div className="text-right shrink-0">
                                          <div className="text-[10px] font-bold text-primary print:text-[8pt]">
                                            {item.points}pt
                                          </div>
                                        </div>
                                      )}
                                    </div>
                                  </div>

                                  {/* Níveis */}
                                  {item.levels && item.levels.length > 0 && (
                                    <div className="p-2 bg-muted/30 print:p-1">
                                      <div className="flex flex-wrap gap-1 print:gap-0.5">
                                        {item.levels
                                          .sort((a, b) => a.level - b.level)
                                          .map((level) => {
                                            const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                            const range = formatLevelRange(level)
                                            const hasValues = level.lowerLimit != null || level.upperLimit != null

                                            return (
                                              <div
                                                key={level.id}
                                                className={`inline-flex items-center gap-1 rounded-full border px-2 py-0.5 text-[10px] font-medium ${style.bg} ${style.text} ${style.border} print:px-1.5 print:py-0 print:text-[7pt]`}
                                                title={`${level.name}${range ? ` (${range})` : ''}${level.definition ? ` - ${level.definition}` : ''}`}
                                              >
                                                <span className="font-bold">N{level.level}:</span>
                                                {hasValues ? (
                                                  <span className="font-mono text-[9px] print:text-[6.5pt]">{range}</span>
                                                ) : (
                                                  <span className="print:text-[6.5pt]">{level.name}</span>
                                                )}
                                              </div>
                                            )
                                          })}
                                      </div>
                                    </div>
                                  )}

                                  {/* Itens Filhos (hierárquicos) */}
                                  {item.childItems && item.childItems.length > 0 && (
                                    <div className="border-t bg-muted/10">
                                      <div className="p-1.5 space-y-1.5 print:p-1 print:space-y-1">
                                        {item.childItems.map((childItem) => (
                                          <div key={childItem.id} className="pl-2 border-l-2 border-primary/30 print:pl-1.5">
                                            <div className="mb-1 print:mb-0.5">
                                              <div className="flex items-start justify-between gap-2">
                                                <h5 className="font-medium text-[10px] text-muted-foreground print:text-[8pt]">
                                                  {childItem.name}
                                                  {childItem.unit && (
                                                    <span className="text-[9px] text-muted-foreground/80 ml-1.5 print:text-[7pt]">
                                                      ({childItem.unit})
                                                    </span>
                                                  )}
                                                </h5>
                                                {childItem.points > 0 && (
                                                  <span className="text-[10px] font-semibold text-primary shrink-0 print:text-[8pt]">
                                                    {childItem.points}pt
                                                  </span>
                                                )}
                                              </div>
                                            </div>

                                            {/* Níveis do item filho */}
                                            {childItem.levels && childItem.levels.length > 0 && (
                                              <div className="flex flex-wrap gap-1 print:gap-0.5">
                                                {childItem.levels
                                                  .sort((a, b) => a.level - b.level)
                                                  .map((level) => {
                                                    const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                                    const range = formatLevelRange(level)
                                                    const hasValues = level.lowerLimit != null || level.upperLimit != null

                                                    return (
                                                      <div
                                                        key={level.id}
                                                        className={`inline-flex items-center gap-1 rounded-full border px-2 py-0.5 text-[10px] font-medium ${style.bg} ${style.text} ${style.border} print:px-1.5 print:py-0 print:text-[7pt]`}
                                                        title={`${level.name}${range ? ` (${range})` : ''}${level.definition ? ` - ${level.definition}` : ''}`}
                                                      >
                                                        <span className="font-bold">N{level.level}:</span>
                                                        {hasValues ? (
                                                          <span className="font-mono text-[9px] print:text-[6.5pt]">{range}</span>
                                                        ) : (
                                                          <span className="print:text-[6.5pt]">{level.name}</span>
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
                              </div>
                            ))}
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Rodapé (visível apenas na impressão) */}
      <div className="hidden print:block mt-2 pt-1 border-t text-center text-[8pt] text-muted-foreground">
        <p>Plenya - Sistema de Prontuário Médico Eletrônico</p>
        <p className="mt-0.5">Este documento contém informações confidenciais e deve ser tratado de acordo com a LGPD</p>
      </div>
    </div>
  )
}
