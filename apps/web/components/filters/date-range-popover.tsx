'use client';

import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { CalendarIcon, X } from 'lucide-react';
import type { DateRange } from 'react-day-picker';

import { Button } from '@/components/ui/button';
import { Calendar } from '@/components/ui/calendar';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { cn } from '@/lib/utils';

interface DateRangePopoverProps {
  value: DateRange | undefined;
  onChange: (range: DateRange | undefined) => void;
  /** Rótulo do botão quando vazio (default "Período"). */
  label?: string;
  className?: string;
}

/**
 * Popover com seletor de intervalo de datas (react-day-picker mode="range").
 * Fechar = aplicar; clicar no X reseta.
 *
 * @example
 *   const [range, setRange] = useState<DateRange | undefined>();
 *   <DateRangePopover value={range} onChange={setRange} label="Criados em" />
 */
export function DateRangePopover({
  value,
  onChange,
  label = 'Período',
  className,
}: DateRangePopoverProps) {
  const hasValue = !!(value?.from || value?.to);

  function renderLabel() {
    if (!value?.from) return label;
    if (value.to) {
      return `${format(value.from, 'dd/MM/yy', { locale: ptBR })} – ${format(
        value.to,
        'dd/MM/yy',
        { locale: ptBR },
      )}`;
    }
    return format(value.from, 'dd/MM/yyyy', { locale: ptBR });
  }

  return (
    <div className={cn('inline-flex items-center', className)}>
      <Popover>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            className={cn(
              'h-9 gap-2 font-normal',
              hasValue && 'border-primary/50 bg-primary/5 pr-2',
            )}
          >
            <CalendarIcon className="h-3.5 w-3.5 opacity-70" />
            <span>{renderLabel()}</span>
          </Button>
        </PopoverTrigger>
        <PopoverContent align="start" className="w-auto p-0">
          <Calendar
            mode="range"
            numberOfMonths={2}
            selected={value}
            onSelect={onChange}
            locale={ptBR}
            defaultMonth={value?.from}
          />
        </PopoverContent>
      </Popover>
      {hasValue && (
        <Button
          variant="ghost"
          size="sm"
          aria-label="Limpar período"
          className="-ml-2 h-9 w-9 p-0"
          onClick={() => onChange(undefined)}
        >
          <X className="h-3.5 w-3.5" />
        </Button>
      )}
    </div>
  );
}
