'use client';

import { Check, ChevronDown } from 'lucide-react';
import { useState } from 'react';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { cn } from '@/lib/utils';

export interface MultiSelectOption<T extends string = string> {
  value: T;
  label: string;
}

interface MultiSelectChipsProps<T extends string = string> {
  /** Rótulo curto exibido no botão quando nada selecionado (ex: "Status"). */
  label: string;
  options: ReadonlyArray<MultiSelectOption<T>>;
  selected: ReadonlyArray<T>;
  onChange: (next: T[]) => void;
  /** Largura do popover (default 240px). */
  contentWidthClass?: string;
  /** Esconder caixa de busca interna (default false). */
  hideSearch?: boolean;
  className?: string;
}

/**
 * Botão estilo Linear/Front: clica e abre popover com checkbox de múltiplos valores.
 * Mostra contador de selecionados quando >0.
 *
 * @example
 *   <MultiSelectChips label="Status" options={STATUS_OPTIONS}
 *     selected={statuses} onChange={setStatuses} />
 */
export function MultiSelectChips<T extends string = string>({
  label,
  options,
  selected,
  onChange,
  contentWidthClass = 'w-60',
  hideSearch = false,
  className,
}: MultiSelectChipsProps<T>) {
  const [open, setOpen] = useState(false);
  const hasSelection = selected.length > 0;

  function toggle(value: T) {
    if (selected.includes(value)) {
      onChange(selected.filter((v) => v !== value));
    } else {
      onChange([...selected, value]);
    }
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          size="sm"
          className={cn(
            'h-9 justify-between gap-2 font-normal',
            hasSelection && 'border-primary/50 bg-primary/5',
            className,
          )}
        >
          <span className="flex items-center gap-2">
            <span>{label}</span>
            {hasSelection && (
              <Badge variant="secondary" className="rounded px-1.5 py-0 text-xs">
                {selected.length}
              </Badge>
            )}
          </span>
          <ChevronDown className="h-3.5 w-3.5 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent align="start" className={cn(contentWidthClass, 'p-0')}>
        <Command>
          {!hideSearch && <CommandInput placeholder={`Filtrar ${label.toLowerCase()}…`} />}
          <CommandList>
            <CommandEmpty>Nenhuma opção.</CommandEmpty>
            <CommandGroup>
              {options.map((opt) => {
                const isSelected = selected.includes(opt.value);
                return (
                  <CommandItem
                    key={opt.value}
                    value={opt.label}
                    onSelect={() => toggle(opt.value)}
                  >
                    <div
                      className={cn(
                        'mr-2 flex h-4 w-4 items-center justify-center rounded-sm border',
                        isSelected
                          ? 'border-primary bg-primary text-primary-foreground'
                          : 'border-input',
                      )}
                    >
                      {isSelected && <Check className="h-3 w-3" />}
                    </div>
                    <span>{opt.label}</span>
                  </CommandItem>
                );
              })}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
