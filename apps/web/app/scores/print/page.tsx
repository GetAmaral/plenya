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
      <div className="hidden print:block mb-8">
        <div className="text-center border-b pb-4">
          <h1 className="text-3xl font-bold">Escore Plenya</h1>
          <p className="text-sm text-muted-foreground mt-2">
            Sistema de Pontuação de Saúde Funcional e Longevidade
          </p>
          <p className="text-xs text-muted-foreground mt-1">
            Gerado em: {new Date().toLocaleDateString('pt-BR', {
              day: '2-digit',
              month: 'long',
              year: 'numeric',
              hour: '2-digit',
              minute: '2-digit'
            })}
          </p>
        </div>
      </div>

      {/* Conteúdo Principal */}
      <div className="container mx-auto px-4 py-8 print:px-0 print:py-0">
        <div className="space-y-8">
          {groups.map((group, groupIndex) => (
            <div key={group.id} className="break-inside-avoid">
              {/* Grupo */}
              <div className="mb-4">
                <div className="bg-primary text-primary-foreground rounded-lg px-6 py-4 print:rounded-none">
                  <div className="flex items-center gap-3">
                    <div className="flex items-center justify-center w-8 h-8 rounded-full bg-primary-foreground/20 text-sm font-bold">
                      {groupIndex + 1}
                    </div>
                    <h2 className="text-xl font-bold">{group.name}</h2>
                  </div>
                </div>
              </div>

              {/* Subgrupos */}
              {group.subgroups && group.subgroups.length > 0 && (
                <div className="space-y-6 ml-4 print:ml-2">
                  {group.subgroups.map((subgroup, subgroupIndex) => (
                    <div key={subgroup.id} className="break-inside-avoid">
                      {/* Subgrupo */}
                      <div className="mb-3">
                        <div className="bg-muted rounded-lg px-5 py-3 print:rounded-none border-l-4 border-primary">
                          <div className="flex items-center gap-2">
                            <div className="flex items-center justify-center w-6 h-6 rounded-full bg-primary/10 text-xs font-semibold text-primary">
                              {groupIndex + 1}.{subgroupIndex + 1}
                            </div>
                            <h3 className="text-lg font-semibold">{subgroup.name}</h3>
                          </div>
                        </div>
                      </div>

                      {/* Itens */}
                      {subgroup.items && subgroup.items.length > 0 && (
                        <div className="space-y-4 ml-4 print:ml-2">
                          {subgroup.items
                            .filter(item => !item.parentItemId) // Apenas itens raiz
                            .map((item) => (
                              <div key={item.id} className="break-inside-avoid">
                                {/* Item */}
                                <div className="border rounded-lg print:rounded-none">
                                  <div className="bg-card px-4 py-3 border-b">
                                    <div className="flex items-start justify-between gap-4">
                                      <div className="flex-1">
                                        <h4 className="font-medium text-sm leading-tight">{item.name}</h4>
                                        {item.unit && (
                                          <p className="text-xs text-muted-foreground mt-1">
                                            Unidade: {item.unit}
                                            {item.unitConversion && ` | ${item.unitConversion}`}
                                          </p>
                                        )}
                                      </div>
                                      <div className="text-right shrink-0">
                                        <div className="text-sm font-bold text-primary">
                                          {item.points} pts
                                        </div>
                                      </div>
                                    </div>
                                  </div>

                                  {/* Níveis */}
                                  {item.levels && item.levels.length > 0 && (
                                    <div className="p-3 bg-muted/30">
                                      <div className="flex flex-wrap gap-2">
                                        {item.levels
                                          .sort((a, b) => a.level - b.level)
                                          .map((level) => {
                                            const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                            const range = formatLevelRange(level)

                                            return (
                                              <div
                                                key={level.id}
                                                className={`inline-flex items-center gap-1.5 rounded-full border px-3 py-1 text-xs font-medium ${style.bg} ${style.text} ${style.border}`}
                                                title={`${level.name}${range ? ` (${range})` : ''}${level.definition ? ` - ${level.definition}` : ''}`}
                                              >
                                                <span className="font-bold">N{level.level}:</span>
                                                <span>{level.name}</span>
                                                {range && (
                                                  <span className="font-mono text-[10px] opacity-75">
                                                    {range}
                                                  </span>
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
                                      <div className="p-3 space-y-3">
                                        {item.childItems.map((childItem) => (
                                          <div key={childItem.id} className="pl-3 border-l-2 border-primary/30">
                                            <div className="mb-2">
                                              <div className="flex items-start justify-between gap-2">
                                                <h5 className="font-medium text-xs text-muted-foreground">
                                                  {childItem.name}
                                                </h5>
                                                <span className="text-xs font-semibold text-primary shrink-0">
                                                  {childItem.points} pts
                                                </span>
                                              </div>
                                              {childItem.unit && (
                                                <p className="text-xs text-muted-foreground/80 mt-0.5">
                                                  {childItem.unit}
                                                </p>
                                              )}
                                            </div>

                                            {/* Níveis do item filho */}
                                            {childItem.levels && childItem.levels.length > 0 && (
                                              <div className="flex flex-wrap gap-1.5">
                                                {childItem.levels
                                                  .sort((a, b) => a.level - b.level)
                                                  .map((level) => {
                                                    const style = LEVEL_STYLES[level.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6]
                                                    const range = formatLevelRange(level)

                                                    return (
                                                      <div
                                                        key={level.id}
                                                        className={`inline-flex items-center gap-1 rounded-full border px-2.5 py-0.5 text-[10px] font-medium ${style.bg} ${style.text} ${style.border}`}
                                                        title={`${level.name}${range ? ` (${range})` : ''}${level.definition ? ` - ${level.definition}` : ''}`}
                                                      >
                                                        <span className="font-bold">N{level.level}:</span>
                                                        <span>{level.name}</span>
                                                        {range && (
                                                          <span className="font-mono text-[9px] opacity-75">
                                                            {range}
                                                          </span>
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
      <div className="hidden print:block mt-8 pt-4 border-t text-center text-xs text-muted-foreground">
        <p>Plenya - Sistema de Prontuário Médico Eletrônico</p>
        <p className="mt-1">Este documento contém informações confidenciais e deve ser tratado de acordo com a LGPD</p>
      </div>
    </div>
  )
}
