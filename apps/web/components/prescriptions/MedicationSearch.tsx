'use client'

import { useState, useEffect } from 'react'
import { useQuery } from '@tanstack/react-query'
import { Check, ChevronsUpDown, Loader2, Search } from 'lucide-react'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Badge } from '@/components/ui/badge'

import { searchMedications } from '@/lib/api/medication-definitions'
import type { MedicationDefinition } from '@/lib/api/medication-definitions'

interface MedicationSearchProps {
  value?: MedicationDefinition
  onSelect: (medication: MedicationDefinition) => void
  placeholder?: string
}

export function MedicationSearch({
  value,
  onSelect,
  placeholder = 'Buscar medicamento...',
}: MedicationSearchProps) {
  const [open, setOpen] = useState(false)
  const [searchQuery, setSearchQuery] = useState('')
  const [debouncedQuery, setDebouncedQuery] = useState('')

  // Debounce search query (300ms)
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedQuery(searchQuery)
    }, 300)

    return () => clearTimeout(timer)
  }, [searchQuery])

  // Query medications
  const { data: medications = [], isLoading } = useQuery({
    queryKey: ['medication-search', debouncedQuery],
    queryFn: () => searchMedications(debouncedQuery, 10),
    enabled: debouncedQuery.length >= 2,
  })

  const handleSelect = (medication: MedicationDefinition) => {
    onSelect(medication)
    setOpen(false)
    setSearchQuery('')
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value ? (
            <div className="flex items-center gap-2 truncate">
              <span className="truncate">{value.commonName}</span>
              <CategoryBadge category={value.category} />
            </div>
          ) : (
            <span className="text-muted-foreground">{placeholder}</span>
          )}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[400px] p-0" align="start">
        <Command shouldFilter={false}>
          <CommandInput
            placeholder="Digite nome ou princípio ativo..."
            value={searchQuery}
            onValueChange={setSearchQuery}
          />
          <CommandList>
            {isLoading && (
              <div className="flex items-center justify-center py-6">
                <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
              </div>
            )}

            {!isLoading && searchQuery.length < 2 && (
              <CommandEmpty>
                <Search className="h-8 w-8 text-muted-foreground mb-2 mx-auto" />
                <p className="text-sm text-muted-foreground">
                  Digite ao menos 2 caracteres para buscar
                </p>
              </CommandEmpty>
            )}

            {!isLoading && searchQuery.length >= 2 && medications.length === 0 && (
              <CommandEmpty>Nenhum medicamento encontrado.</CommandEmpty>
            )}

            {!isLoading && medications.length > 0 && (
              <CommandGroup>
                {medications.map((medication) => (
                  <CommandItem
                    key={medication.id}
                    value={medication.id}
                    onSelect={() => handleSelect(medication)}
                    className="flex items-start gap-2 py-3"
                  >
                    <Check
                      className={cn(
                        'mt-1 h-4 w-4',
                        value?.id === medication.id ? 'opacity-100' : 'opacity-0'
                      )}
                    />
                    <div className="flex-1 space-y-1">
                      <div className="flex items-center gap-2">
                        <span className="font-medium">{medication.commonName}</span>
                        <CategoryBadge category={medication.category} />
                      </div>
                      <p className="text-sm text-muted-foreground">
                        {medication.activeIngredient}
                      </p>
                      <div className="flex items-center gap-3 text-xs text-muted-foreground">
                        <span>Validade: {medication.validityDays} dias</span>
                        {medication.requiresSNCR && (
                          <Badge variant="outline" className="text-xs">
                            SNCR
                          </Badge>
                        )}
                        {medication.requiresDigitalSignature && (
                          <Badge variant="outline" className="text-xs">
                            Assinatura Digital
                          </Badge>
                        )}
                      </div>
                    </div>
                  </CommandItem>
                ))}
              </CommandGroup>
            )}
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}

function CategoryBadge({ category }: { category: string }) {
  const variants: Record<string, { label: string; className: string }> = {
    simple: { label: 'Simples', className: 'bg-blue-50 text-blue-700 text-xs' },
    c1: { label: 'C1', className: 'bg-orange-50 text-orange-700 text-xs' },
    c5: { label: 'C5', className: 'bg-red-50 text-red-700 text-xs' },
    antibiotic: { label: 'Antibiótico', className: 'bg-purple-50 text-purple-700 text-xs' },
    glp1: { label: 'GLP-1', className: 'bg-green-50 text-green-700 text-xs' },
  }

  const variant = variants[category] || variants.simple

  return (
    <Badge variant="outline" className={variant.className}>
      {variant.label}
    </Badge>
  )
}
