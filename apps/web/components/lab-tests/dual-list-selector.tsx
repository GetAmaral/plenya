'use client'

import { useState, useMemo } from 'react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card } from '@/components/ui/card'
import { Search, ChevronRight, ChevronLeft, X } from 'lucide-react'
import { LabTestDefinition } from '@/lib/api/lab-request-templates'
import { ScrollArea } from '@/components/ui/scroll-area'

interface DualListSelectorProps {
  availableTests: LabTestDefinition[]
  selectedTests: LabTestDefinition[]
  onSelectionChange: (selected: LabTestDefinition[]) => void
}

export function DualListSelector({
  availableTests,
  selectedTests,
  onSelectionChange
}: DualListSelectorProps) {
  const [searchLeft, setSearchLeft] = useState('')
  const [searchRight, setSearchRight] = useState('')
  const [highlightedLeft, setHighlightedLeft] = useState<string | null>(null)
  const [highlightedRight, setHighlightedRight] = useState<string | null>(null)

  // Filter tests that are NOT yet selected
  const unselectedTests = useMemo(() => {
    const selectedIds = new Set(selectedTests.map(t => t.id))
    return availableTests.filter(t => !selectedIds.has(t.id))
  }, [availableTests, selectedTests])

  // Filter by search
  const filteredLeft = useMemo(() => {
    if (!searchLeft) return unselectedTests
    const search = searchLeft.toLowerCase()
    return unselectedTests.filter(t =>
      t.name.toLowerCase().includes(search) ||
      t.code.toLowerCase().includes(search)
    )
  }, [unselectedTests, searchLeft])

  const filteredRight = useMemo(() => {
    if (!searchRight) return selectedTests
    const search = searchRight.toLowerCase()
    return selectedTests.filter(t =>
      t.name.toLowerCase().includes(search) ||
      t.code.toLowerCase().includes(search)
    )
  }, [selectedTests, searchRight])

  // Add test to selected
  const addTest = (test: LabTestDefinition) => {
    onSelectionChange([...selectedTests, test])
    setHighlightedLeft(null)
  }

  // Remove test from selected
  const removeTest = (test: LabTestDefinition) => {
    onSelectionChange(selectedTests.filter(t => t.id !== test.id))
    setHighlightedRight(null)
  }

  // Handle Enter key
  const handleKeyDownLeft = (e: React.KeyboardEvent, test: LabTestDefinition) => {
    if (e.key === 'Enter') {
      e.preventDefault()
      addTest(test)
    }
  }

  const handleKeyDownRight = (e: React.KeyboardEvent, test: LabTestDefinition) => {
    if (e.key === 'Enter') {
      e.preventDefault()
      removeTest(test)
    }
  }

  return (
    <div className="grid grid-cols-[1fr_auto_1fr] gap-4">
      {/* Left column - Available tests */}
      <Card className="p-4">
        <div className="space-y-4">
          <div>
            <h3 className="text-sm font-medium mb-2">Exames Disponíveis</h3>
            <div className="relative">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Buscar exames..."
                value={searchLeft}
                onChange={(e) => setSearchLeft(e.target.value)}
                className="pl-8"
              />
            </div>
          </div>

          <ScrollArea className="h-[500px]">
            <div className="space-y-1">
              {filteredLeft.length === 0 ? (
                <p className="text-sm text-muted-foreground text-center py-8">
                  {searchLeft ? 'Nenhum exame encontrado' : 'Todos os exames já foram selecionados'}
                </p>
              ) : (
                filteredLeft.map((test) => (
                  <div
                    key={test.id}
                    className={`
                      p-2 rounded-md cursor-pointer transition-colors text-sm
                      hover:bg-accent
                      ${highlightedLeft === test.id ? 'bg-accent' : ''}
                    `}
                    onClick={() => addTest(test)}
                    onKeyDown={(e) => handleKeyDownLeft(e, test)}
                    onMouseEnter={() => setHighlightedLeft(test.id)}
                    onMouseLeave={() => setHighlightedLeft(null)}
                    tabIndex={0}
                  >
                    <div className="font-medium">{test.name}</div>
                    {test.tussCode && (
                      <div className="text-xs text-muted-foreground">TUSS: {test.tussCode}</div>
                    )}
                  </div>
                ))
              )}
            </div>
          </ScrollArea>

          <p className="text-xs text-muted-foreground">
            {filteredLeft.length} exame{filteredLeft.length !== 1 ? 's' : ''} disponível{filteredLeft.length !== 1 ? 'is' : ''}
          </p>
        </div>
      </Card>

      {/* Middle - Action buttons */}
      <div className="flex flex-col justify-center gap-2">
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            if (highlightedLeft) {
              const test = filteredLeft.find(t => t.id === highlightedLeft)
              if (test) addTest(test)
            }
          }}
          disabled={!highlightedLeft}
          title="Adicionar selecionado (ou clique duplo)"
        >
          <ChevronRight className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            if (highlightedRight) {
              const test = filteredRight.find(t => t.id === highlightedRight)
              if (test) removeTest(test)
            }
          }}
          disabled={!highlightedRight}
          title="Remover selecionado (ou clique duplo)"
        >
          <ChevronLeft className="h-4 w-4" />
        </Button>
      </div>

      {/* Right column - Selected tests */}
      <Card className="p-4">
        <div className="space-y-4">
          <div>
            <h3 className="text-sm font-medium mb-2">Exames Selecionados</h3>
            <div className="relative">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Filtrar selecionados..."
                value={searchRight}
                onChange={(e) => setSearchRight(e.target.value)}
                className="pl-8"
              />
            </div>
          </div>

          <ScrollArea className="h-[500px]">
            <div className="space-y-1">
              {filteredRight.length === 0 ? (
                <p className="text-sm text-muted-foreground text-center py-8">
                  {searchRight ? 'Nenhum exame encontrado' : 'Nenhum exame selecionado'}
                </p>
              ) : (
                filteredRight.map((test) => (
                  <div
                    key={test.id}
                    className={`
                      p-2 rounded-md cursor-pointer transition-colors text-sm
                      hover:bg-destructive/10 group
                      ${highlightedRight === test.id ? 'bg-destructive/10' : ''}
                    `}
                    onClick={() => removeTest(test)}
                    onKeyDown={(e) => handleKeyDownRight(e, test)}
                    onMouseEnter={() => setHighlightedRight(test.id)}
                    onMouseLeave={() => setHighlightedRight(null)}
                    tabIndex={0}
                  >
                    <div className="flex items-start justify-between gap-2">
                      <div className="flex-1">
                        <div className="font-medium">{test.name}</div>
                        {test.tussCode && (
                          <div className="text-xs text-muted-foreground">TUSS: {test.tussCode}</div>
                        )}
                      </div>
                      <X className="h-4 w-4 text-muted-foreground group-hover:text-destructive flex-shrink-0 mt-0.5" />
                    </div>
                  </div>
                ))
              )}
            </div>
          </ScrollArea>

          <p className="text-xs text-muted-foreground">
            {filteredRight.length} exame{filteredRight.length !== 1 ? 's' : ''} selecionado{filteredRight.length !== 1 ? 's' : ''}
          </p>
        </div>
      </Card>
    </div>
  )
}
