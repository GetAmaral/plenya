'use client';

import { X } from 'lucide-react';
import type { DateRange } from 'react-day-picker';

import { DateRangePopover } from '@/components/filters/date-range-popover';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';

interface PatientsFilterBarProps {
  dateRange: DateRange | undefined;
  onDateRangeChange: (range: DateRange | undefined) => void;

  onlyActivePlan: boolean;
  onOnlyActivePlanChange: (next: boolean) => void;

  hasActiveFilters: boolean;
  onClear: () => void;
}

/**
 * Linha de filtros simplificada da listagem de Pacientes.
 */
export function PatientsFilterBar({
  dateRange,
  onDateRangeChange,
  onlyActivePlan,
  onOnlyActivePlanChange,
  hasActiveFilters,
  onClear,
}: PatientsFilterBarProps) {
  return (
    <div className="flex flex-wrap items-center gap-3">
      <DateRangePopover
        value={dateRange}
        onChange={onDateRangeChange}
        label="Cadastrados em"
      />
      <div className="inline-flex items-center gap-2">
        <Switch
          id="only-active-plan"
          checked={onlyActivePlan}
          onCheckedChange={onOnlyActivePlanChange}
        />
        <Label htmlFor="only-active-plan" className="cursor-pointer text-sm font-normal">
          Apenas com plano ativo
        </Label>
      </div>
      {hasActiveFilters && (
        <Button
          variant="ghost"
          size="sm"
          onClick={onClear}
          className="h-9 gap-1 text-muted-foreground"
        >
          <X className="h-3.5 w-3.5" />
          Limpar filtros
        </Button>
      )}
    </div>
  );
}
