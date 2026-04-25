'use client';

/**
 * DoctorMultiSelect — popover com checkbox por médico.
 *
 * - Usado no /calendario pra escolher quais agendas exibir simultaneamente
 *   (secretária, manager, admin ou doctor curiosa sobre colega).
 * - Persistência da seleção fica no caller (localStorage por user).
 * - Mostra contagem ("3 médicos") no trigger e badges com cores derivadas
 *   pra reforçar a relação visual com os blocos do grid.
 */
import { useMemo, useState } from 'react';
import { Check, ChevronDown, Users } from 'lucide-react';

import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import { doctorColorClasses, doctorDotClass } from '@/lib/calendar/doctor-colors';

interface DoctorOption {
  id: string;
  name: string;
}

interface DoctorMultiSelectProps {
  doctors: DoctorOption[];
  selectedIds: string[];
  onChange: (ids: string[]) => void;
  className?: string;
}

export function DoctorMultiSelect({
  doctors,
  selectedIds,
  onChange,
  className,
}: DoctorMultiSelectProps) {
  const [open, setOpen] = useState(false);
  const [search, setSearch] = useState('');

  const filtered = useMemo(() => {
    const q = search.trim().toLowerCase();
    if (!q) return doctors;
    return doctors.filter((d) => d.name.toLowerCase().includes(q));
  }, [doctors, search]);

  const selectedSet = useMemo(() => new Set(selectedIds), [selectedIds]);

  const toggle = (id: string) => {
    const next = new Set(selectedSet);
    if (next.has(id)) next.delete(id);
    else next.add(id);
    onChange(Array.from(next));
  };

  const selectAll = () => onChange(doctors.map((d) => d.id));
  const clearAll = () => onChange([]);

  const triggerLabel =
    selectedIds.length === 0
      ? 'Nenhum médico'
      : selectedIds.length === doctors.length
      ? 'Todos os médicos'
      : `${selectedIds.length} médico${selectedIds.length > 1 ? 's' : ''}`;

  return (
    <div className={cn('flex flex-col gap-2', className)}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button variant="outline" className="justify-between" size="sm">
            <span className="flex items-center gap-2">
              <Users className="h-4 w-4" />
              {triggerLabel}
            </span>
            <ChevronDown className="ml-2 h-4 w-4 opacity-60" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-72 p-0" align="end">
          <div className="border-b p-2">
            <Input
              autoFocus
              placeholder="Buscar médico..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="h-8"
            />
          </div>
          <div className="flex items-center justify-between border-b px-2 py-1.5 text-xs">
            <button
              type="button"
              onClick={selectAll}
              className="text-primary hover:underline"
            >
              Selecionar todos
            </button>
            <button
              type="button"
              onClick={clearAll}
              className="text-muted-foreground hover:underline"
            >
              Limpar
            </button>
          </div>
          <div className="max-h-72 overflow-y-auto p-1">
            {filtered.length === 0 ? (
              <p className="px-3 py-4 text-center text-xs text-muted-foreground">
                Nenhum médico encontrado
              </p>
            ) : (
              filtered.map((d) => {
                const checked = selectedSet.has(d.id);
                return (
                  <button
                    key={d.id}
                    type="button"
                    onClick={() => toggle(d.id)}
                    className="flex w-full items-center gap-2 rounded-sm px-2 py-1.5 text-sm hover:bg-accent"
                  >
                    <Checkbox checked={checked} className="pointer-events-none" />
                    <span
                      className={cn('h-2 w-2 rounded-full', doctorDotClass(d.id))}
                      aria-hidden
                    />
                    <span className="flex-1 truncate text-left">{d.name}</span>
                    {checked && <Check className="h-3.5 w-3.5 text-primary" />}
                  </button>
                );
              })
            )}
          </div>
        </PopoverContent>
      </Popover>

      {selectedIds.length > 0 && selectedIds.length < doctors.length && (
        <div className="flex flex-wrap gap-1">
          {selectedIds.map((id) => {
            const d = doctors.find((x) => x.id === id);
            if (!d) return null;
            return (
              <Badge
                key={id}
                variant="outline"
                className={cn('px-2 py-0 text-[11px]', doctorColorClasses(id).badge)}
              >
                {d.name}
              </Badge>
            );
          })}
        </div>
      )}
    </div>
  );
}
