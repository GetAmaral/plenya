'use client';

/**
 * CalendarGrid — visão tipo Google Calendar (Day/Week).
 *
 * Layout:
 *   - Linhas: hours (default 8h-20h, configurable HOURS_RANGE)
 *   - Colunas: dias visíveis
 *   - Cada appointment = bloco posicionado por top/height calculados
 *
 * Click em bloco → onSelect callback (página abre drawer/modal).
 *
 * Performance: simples client-side render — O(N appointments). Pra clinics
 * com >500 consultas/semana otimizar com virtualização.
 */
import { useMemo } from 'react';
import { addDays, format, isSameDay, startOfDay } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { cn } from '@/lib/utils';
import {
  APPOINTMENT_TYPE_COLORS,
  APPOINTMENT_TYPE_LABELS,
  type Appointment,
} from '@/lib/api/calendar-api';
import { doctorBlockClass } from '@/lib/calendar/doctor-colors';

const HOUR_START = 8;
const HOUR_END = 20;
const HOUR_HEIGHT_PX = 56; // px por hora
const TOTAL_HEIGHT = (HOUR_END - HOUR_START) * HOUR_HEIGHT_PX;

interface CalendarGridProps {
  /** Modo: 'day' (1 col) ou 'week' (7 cols). */
  view: 'day' | 'week';
  /** Data de referência. Em modo 'week', semana de domingo. */
  referenceDate: Date;
  appointments: Appointment[];
  onSelectAppointment: (a: Appointment) => void;
  /**
   * Quando true, blocos recebem cor por médico (em vez de cor por tipo).
   * Útil quando há multi-doctor view. Caller passa true se selecionou >1.
   */
  colorByDoctor?: boolean;
  /** Clique em área vazia de um dia → callback com a data (yyyy-MM-dd). */
  onSlotClick?: (dateYmd: string) => void;
}

export function CalendarGrid({
  view,
  referenceDate,
  appointments,
  onSelectAppointment,
  colorByDoctor = false,
  onSlotClick,
}: CalendarGridProps) {
  const days = useMemo(() => {
    if (view === 'day') return [startOfDay(referenceDate)];
    // Week starts Sunday
    const weekStart = startOfDay(addDays(referenceDate, -referenceDate.getDay()));
    return Array.from({ length: 7 }, (_, i) => addDays(weekStart, i));
  }, [view, referenceDate]);

  const today = startOfDay(new Date());

  return (
    <div className="overflow-x-auto rounded-lg border bg-background">
      <div
        className="grid"
        style={{
          gridTemplateColumns: `60px repeat(${days.length}, minmax(120px, 1fr))`,
        }}
      >
        {/* Header row */}
        <div className="border-b border-r bg-muted/40" />
        {days.map((d) => {
          const isToday = isSameDay(d, today);
          return (
            <div
              key={d.toISOString()}
              className={cn(
                'border-b border-r p-2 text-center text-sm',
                isToday && 'bg-primary/10 font-semibold',
              )}
            >
              <div className="text-xs uppercase text-muted-foreground">
                {format(d, 'EEE', { locale: ptBR })}
              </div>
              <div className={cn('text-base', isToday && 'text-primary')}>
                {format(d, 'd MMM', { locale: ptBR })}
              </div>
            </div>
          );
        })}

        {/* Hours column + day columns */}
        <div className="relative border-r" style={{ height: TOTAL_HEIGHT }}>
          {Array.from({ length: HOUR_END - HOUR_START }, (_, i) => HOUR_START + i).map(
            (h) => (
              <div
                key={h}
                className="absolute left-0 right-0 border-t text-right pr-1 text-[10px] text-muted-foreground"
                style={{ top: (h - HOUR_START) * HOUR_HEIGHT_PX }}
              >
                {String(h).padStart(2, '0')}:00
              </div>
            ),
          )}
        </div>

        {days.map((d) => {
          const dayAppts = appointments.filter((a) =>
            isSameDay(new Date(a.scheduledAt), d),
          );
          return (
            <div
              key={d.toISOString()}
              className={cn('relative border-r', onSlotClick && 'cursor-pointer')}
              style={{ height: TOTAL_HEIGHT }}
              onClick={() => onSlotClick?.(format(d, 'yyyy-MM-dd'))}
              title={onSlotClick ? 'Clique para agendar neste dia' : undefined}
            >
              {/* Hour gridlines */}
              {Array.from({ length: HOUR_END - HOUR_START }, (_, i) => i).map((i) => (
                <div
                  key={i}
                  className="absolute left-0 right-0 border-t border-stone-100 dark:border-stone-800"
                  style={{ top: i * HOUR_HEIGHT_PX }}
                />
              ))}
              {/* Appointments */}
              {dayAppts.map((a) => {
                const start = new Date(a.scheduledAt);
                const startMinutesFromMidnight = start.getHours() * 60 + start.getMinutes();
                const startMinutesFromGrid = startMinutesFromMidnight - HOUR_START * 60;
                if (startMinutesFromGrid < 0) return null;
                const top = (startMinutesFromGrid / 60) * HOUR_HEIGHT_PX;
                const height = Math.max(
                  20,
                  (a.durationMinutes / 60) * HOUR_HEIGHT_PX - 2,
                );
                const baseClass = colorByDoctor
                  ? cn('border-l-4', doctorBlockClass(a.doctorId))
                  : APPOINTMENT_TYPE_COLORS[a.type];
                const patientName = a.patient?.name ?? 'Paciente';
                const typeLabel = APPOINTMENT_TYPE_LABELS[a.type];
                const doctorName = a.doctor?.name;
                const titleParts = [`${format(start, 'HH:mm')} ${patientName}`, typeLabel];
                if (doctorName) titleParts.push(`Dr(a). ${doctorName}`);
                return (
                  <button
                    key={a.id}
                    type="button"
                    onClick={(e) => {
                      e.stopPropagation();
                      onSelectAppointment(a);
                    }}
                    className={cn(
                      'absolute left-1 right-1 overflow-hidden rounded border px-1.5 py-1 text-left text-xs shadow-sm transition-colors',
                      baseClass,
                      a.status === 'cancelled' && 'opacity-50 line-through',
                    )}
                    style={{ top, height }}
                    title={titleParts.join(' · ')}
                  >
                    <div className="truncate font-semibold">
                      {format(start, 'HH:mm')} {patientName}
                    </div>
                    <div className="flex items-center justify-between gap-1 text-[10px] opacity-80">
                      <span className="truncate">{typeLabel}</span>
                      {colorByDoctor && doctorName && (
                        <span className="truncate font-medium">{doctorName.split(' ')[0]}</span>
                      )}
                    </div>
                  </button>
                );
              })}
            </div>
          );
        })}
      </div>
    </div>
  );
}
