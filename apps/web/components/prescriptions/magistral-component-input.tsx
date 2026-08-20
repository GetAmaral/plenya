'use client'

import { useEffect, useRef, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { Search } from 'lucide-react'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { MagistralSubstanceSheet } from '@/components/prescriptions/magistral-substance-sheet'
import {
  searchMagistralComponents,
  type MagistralComponent,
} from '@/lib/api/magistral-components'

interface MagistralComponentInputProps {
  /** Nome do campo — mantém o input identificável no DOM como os demais do formulário. */
  name?: string
  value: string
  onChange: (value: string) => void
  onSelect: (component: MagistralComponent, dose?: number | null) => void
  placeholder?: string
}

/**
 * Campo de substância com dois caminhos, de propósito:
 *
 *   · o ATALHO — digitar e escolher numa lista de uma linha por item (nome + dose). É o caminho
 *     de quem já sabe o que quer e não precisa ler nada;
 *   · a CONSULTA — a lupa abre o painel lateral com indicação, faixa de dose, sinalizadores e os
 *     trechos das aulas. É o caminho de quem está decidindo.
 *
 * A lista curta existe porque a versão anterior despejava a indicação inteira dentro do dropdown:
 * uma substância ocupava a lista toda e escondia as demais. Texto longo agora só no painel.
 */
export function MagistralComponentInput({
  name,
  value,
  onChange,
  onSelect,
  placeholder,
}: MagistralComponentInputProps) {
  const [open, setOpen] = useState(false)
  const [sheetOpen, setSheetOpen] = useState(false)
  const [term, setTerm] = useState('')
  const containerRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const t = setTimeout(() => setTerm(value), 250)
    return () => clearTimeout(t)
  }, [value])

  const { data: suggestions = [] } = useQuery({
    queryKey: ['magistral-components', term],
    queryFn: () => searchMagistralComponents(term, { allowEmpty: false }),
    enabled: open && term.trim().length >= 2,
  })

  useEffect(() => {
    const onClickOutside = (e: MouseEvent) => {
      if (containerRef.current && !containerRef.current.contains(e.target as Node)) setOpen(false)
    }
    document.addEventListener('mousedown', onClickOutside)
    return () => document.removeEventListener('mousedown', onClickOutside)
  }, [])

  const visible = open && !sheetOpen && suggestions.length > 0

  return (
    <div ref={containerRef} className="relative">
      <div className="flex gap-1">
        <Input
          name={name}
          value={value}
          placeholder={placeholder ?? 'Substância'}
          onChange={(e) => {
            onChange(e.target.value)
            setOpen(true)
          }}
          onFocus={() => setOpen(true)}
          autoComplete="off"
        />
        <Button
          type="button"
          variant="outline"
          size="icon"
          className="shrink-0"
          title="Abrir o catálogo"
          onClick={() => {
            setOpen(false)
            setSheetOpen(true)
          }}
        >
          <Search className="h-4 w-4" />
        </Button>
      </div>

      {visible && (
        <ul className="absolute z-50 mt-1 max-h-64 w-full overflow-auto rounded-md border bg-popover p-1 shadow-md">
          {suggestions.map((s) => (
            <li key={s.id}>
              <button
                type="button"
                className="flex w-full items-baseline justify-between gap-3 rounded px-2 py-1.5 text-left text-sm hover:bg-accent"
                onClick={() => {
                  onSelect(s)
                  setOpen(false)
                }}
              >
                <span className="truncate">{s.name}</span>
                <span className="shrink-0 text-xs text-muted-foreground">
                  {s.usualDose
                    ? `${String(s.usualDose).replace('.', ',')} ${s.defaultUnit}`
                    : s.defaultUnit}
                </span>
              </button>
            </li>
          ))}
          <li className="border-t pt-1">
            <button
              type="button"
              className="flex w-full items-center gap-2 rounded px-2 py-1.5 text-left text-xs text-muted-foreground hover:bg-accent"
              onClick={() => {
                setOpen(false)
                setSheetOpen(true)
              }}
            >
              <Search className="h-3.5 w-3.5" />
              Ver indicações, faixa de dose e de onde saiu
            </button>
          </li>
        </ul>
      )}

      <MagistralSubstanceSheet
        open={sheetOpen}
        onOpenChange={setSheetOpen}
        initialQuery={value}
        onPick={(component, dose) => onSelect(component, dose)}
      />
    </div>
  )
}
