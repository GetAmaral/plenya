'use client';

/**
 * SlotPicker — grid de slots livres pra selecionar horário de consulta.
 *
 * Recebe doctorId + date + type. Quando date muda, refetch automático.
 * Cada slot livre vira um botão clicável; selecionado fica destacado.
 *
 * Backend já intersecciona working_hours ∩ ¬absence ∩ ¬google_busy ∩ ¬appointment.
 * Portanto se um slot aparece, está livre. Slots não retornados = indisponíveis.
 */
import { formatDate } from '@/lib/format-date';
import { Skeleton } from '@/components/ui/skeleton';
import { cn } from '@/lib/utils';
import {
  useCalendarSlots,
  type AppointmentType,
  type CalendarSlot,
} from '@/lib/api/calendar-api';

interface SlotPickerProps {
  doctorId: string | undefined;
  /** YYYY-MM-DD (sem TZ — interpretado em America/Sao_Paulo no backend) */
  date: string | undefined;
  type: AppointmentType;
  overrideDuration?: number;
  /** ISO UTC do slot atualmente selecionado (compara startUtc) */
  selectedSlotUtc?: string;
  onSelect: (slot: CalendarSlot) => void;
}

export function SlotPicker({
  doctorId,
  date,
  type,
  overrideDuration,
  selectedSlotUtc,
  onSelect,
}: SlotPickerProps) {
  const { data, isLoading, isError, error } = useCalendarSlots(
    doctorId,
    date,
    type,
    overrideDuration,
  );

  if (!doctorId || !date) {
    return (
      <p className="text-sm italic text-muted-foreground">
        Selecione médico e data para ver slots disponíveis.
      </p>
    );
  }

  if (isLoading) {
    return (
      <div className="grid grid-cols-3 gap-2 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8">
        {Array.from({ length: 16 }).map((_, i) => (
          <Skeleton key={i} className="h-10 w-full" />
        ))}
      </div>
    );
  }

  if (isError) {
    return (
      <p className="text-sm text-rose-600">
        Erro ao carregar slots: {error instanceof Error ? error.message : 'desconhecido'}
      </p>
    );
  }

  const slots = data?.slots ?? [];

  if (slots.length === 0) {
    return (
      <p className="rounded border bg-muted/40 p-4 text-sm italic text-muted-foreground">
        Nenhum horário disponível nesta data. Tente outra data ou ajuste os horários
        de atendimento em Configurações &gt; Agenda.
      </p>
    );
  }

  return (
    <div className="grid grid-cols-3 gap-2 sm:grid-cols-4 md:grid-cols-6 lg:grid-cols-8">
      {slots.map((slot) => {
        const isSelected = selectedSlotUtc === slot.startUtc;
        const startLabel = formatDate(slot.startUtc, 'HH:mm');
        return (
          <button
            key={slot.startUtc}
            type="button"
            onClick={() => onSelect(slot)}
            className={cn(
              'rounded-md border px-3 py-2 text-sm font-medium transition-colors',
              isSelected
                ? 'border-primary bg-primary text-primary-foreground'
                : 'border-input bg-background hover:bg-accent hover:text-accent-foreground',
            )}
          >
            {startLabel}
          </button>
        );
      })}
    </div>
  );
}
