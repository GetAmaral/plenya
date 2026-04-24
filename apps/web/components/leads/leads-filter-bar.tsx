'use client';

import { X } from 'lucide-react';
import type { DateRange } from 'react-day-picker';

import { DateRangePopover } from '@/components/filters/date-range-popover';
import { MultiSelectChips, type MultiSelectOption } from '@/components/filters/multi-select-chips';
import {
  SingleSelectPill,
  type SingleSelectOption,
} from '@/components/filters/single-select-pill';
import { Button } from '@/components/ui/button';
import {
  STATUS_LABELS,
  SOURCE_LABELS,
  type LeadSource,
  type LeadStatus,
  type StaffUser,
} from '@/lib/api/leads-api';

const STATUS_OPTIONS: ReadonlyArray<MultiSelectOption<LeadStatus>> = (
  Object.keys(STATUS_LABELS) as LeadStatus[]
).map((value) => ({ value, label: STATUS_LABELS[value] }));

const SOURCE_OPTIONS: ReadonlyArray<MultiSelectOption<LeadSource>> = (
  Object.keys(SOURCE_LABELS) as LeadSource[]
).map((value) => ({ value, label: SOURCE_LABELS[value] }));

export interface AssigneeChoice {
  /** "me" / "unassigned" / staffUserId */
  value: string;
  label: string;
}

interface LeadsFilterBarProps {
  statuses: LeadStatus[];
  onStatusesChange: (next: LeadStatus[]) => void;

  sources: LeadSource[];
  onSourcesChange: (next: LeadSource[]) => void;

  assignee: string | null;
  onAssigneeChange: (next: string | null) => void;
  staffUsers: StaffUser[];
  currentUserId: string | undefined;

  dateRange: DateRange | undefined;
  onDateRangeChange: (range: DateRange | undefined) => void;

  hasActiveFilters: boolean;
  onClear: () => void;
}

/**
 * Linha de chips estilo Linear/Front App pra filtrar a lista de Leads.
 * Inclui Status (multi), Origem (multi), Atribuído a (single), Período (range).
 */
export function LeadsFilterBar({
  statuses,
  onStatusesChange,
  sources,
  onSourcesChange,
  assignee,
  onAssigneeChange,
  staffUsers,
  currentUserId,
  dateRange,
  onDateRangeChange,
  hasActiveFilters,
  onClear,
}: LeadsFilterBarProps) {
  const assigneeOptions: SingleSelectOption[] = [
    ...(currentUserId ? [{ value: 'me', label: 'Atribuídos a mim' }] : []),
    { value: 'unassigned', label: 'Sem atribuição' },
    ...staffUsers
      .filter((u) => u.id !== currentUserId)
      .map<SingleSelectOption>((u) => ({ value: u.id, label: u.name })),
  ];

  return (
    <div className="flex flex-wrap items-center gap-2">
      <MultiSelectChips<LeadStatus>
        label="Status"
        options={STATUS_OPTIONS}
        selected={statuses}
        onChange={onStatusesChange}
        hideSearch
      />
      <MultiSelectChips<LeadSource>
        label="Origem"
        options={SOURCE_OPTIONS}
        selected={sources}
        onChange={onSourcesChange}
        hideSearch
      />
      <SingleSelectPill
        label="Atribuído a"
        options={assigneeOptions}
        value={assignee}
        onChange={onAssigneeChange}
      />
      <DateRangePopover
        value={dateRange}
        onChange={onDateRangeChange}
        label="Criados em"
      />
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
