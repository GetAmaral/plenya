'use client';

import { Check, ChevronDown } from 'lucide-react';
import { useState } from 'react';

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

export interface SingleSelectOption<T extends string = string> {
  value: T;
  label: string;
}

interface SingleSelectPillProps<T extends string = string> {
  label: string;
  options: ReadonlyArray<SingleSelectOption<T>>;
  /** Valor selecionado. `null`/undefined = nenhum. */
  value: T | null | undefined;
  onChange: (value: T | null) => void;
  /** Largura do popover (default 240px). */
  contentWidthClass?: string;
  /** Esconder caixa de busca (default false). Útil pra listas curtas. */
  hideSearch?: boolean;
  className?: string;
}

/**
 * Botão pill com popover de seleção única (combobox). Mostra o label do
 * valor escolhido ou só o `label` placeholder.
 *
 * @example
 *   <SingleSelectPill label="Atribuído a" options={ASSIGNEE_OPTIONS}
 *     value={assignee} onChange={setAssignee} />
 */
export function SingleSelectPill<T extends string = string>({
  label,
  options,
  value,
  onChange,
  contentWidthClass = 'w-60',
  hideSearch = false,
  className,
}: SingleSelectPillProps<T>) {
  const [open, setOpen] = useState(false);
  const selected = options.find((o) => o.value === value);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          size="sm"
          className={cn(
            'h-9 justify-between gap-2 font-normal',
            selected && 'border-primary/50 bg-primary/5',
            className,
          )}
        >
          <span className="truncate">
            {selected ? `${label}: ${selected.label}` : label}
          </span>
          <ChevronDown className="h-3.5 w-3.5 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent align="start" className={cn(contentWidthClass, 'p-0')}>
        <Command>
          {!hideSearch && <CommandInput placeholder={`Filtrar ${label.toLowerCase()}…`} />}
          <CommandList>
            <CommandEmpty>Nenhuma opção.</CommandEmpty>
            <CommandGroup>
              {options.map((opt) => {
                const isSelected = value === opt.value;
                return (
                  <CommandItem
                    key={opt.value}
                    value={opt.label}
                    onSelect={() => {
                      onChange(isSelected ? null : opt.value);
                      setOpen(false);
                    }}
                  >
                    <div
                      className={cn(
                        'mr-2 flex h-4 w-4 items-center justify-center rounded-full border',
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
