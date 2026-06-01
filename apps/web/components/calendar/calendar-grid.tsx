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
 * Click em área vazia de um dia → onSlotClick (página cria consulta no dia).
 *
 * Drag-and-drop (enableDragDrop): arrastar um bloco reagenda a consulta.
 *   - Só consultas 'scheduled'/'confirmed' são arrastáveis.
 *   - PointerSensor com activationConstraint distance:6px → clique (<6px) continua
 *     abrindo o detalhe e o clique no vazio continua criando; só >6px vira arrasto.
 *   - delta.y → minutos (snap 15min); coluna de destino (over) → novo dia.
 *   - Reagendamento otimista (useRescheduleAppointment) com rollback + toast no 409.
 *
 * Performance: simples client-side render — O(N appointments). Pra clinics
 * com >500 consultas/semana otimizar com virtualização.
 */
import { useMemo, useRef, type MutableRefObject } from 'react';
import { addDays, format, isSameDay, startOfDay } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import {
  DndContext,
  MouseSensor,
  TouchSensor,
  useSensor,
  useSensors,
  useDraggable,
  useDroppable,
  type DragEndEvent,
} from '@dnd-kit/core';
import { CSS } from '@dnd-kit/utilities';
import { toast } from 'sonner';
import { cn } from '@/lib/utils';
import {
  APPOINTMENT_TYPE_COLORS,
  APPOINTMENT_TYPE_LABELS,
  useRescheduleAppointment,
  type Appointment,
} from '@/lib/api/calendar-api';
import { doctorBlockClass } from '@/lib/calendar/doctor-colors';

const HOUR_START = 8;
const HOUR_END = 20;
const HOUR_HEIGHT_PX = 56; // px por hora
const TOTAL_HEIGHT = (HOUR_END - HOUR_START) * HOUR_HEIGHT_PX;
const SNAP_MINUTES = 15;

/** Só consultas ainda "vivas" podem ser arrastadas pra reagendar. */
function canDragStatus(status: Appointment['status']): boolean {
  return status === 'scheduled' || status === 'confirmed';
}

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
  /** Habilita arrastar blocos pra reagendar (drag-and-drop). */
  enableDragDrop?: boolean;
  /**
   * Decide se a consulta pode ser arrastada pelo usuário atual. Espelha a
   * autorização do backend (médico só reagenda a própria; admin/secretária/
   * gerente reagendam qualquer uma) pra evitar o "pulo + rollback + 403".
   * Default: tudo arrastável.
   */
  canDragAppointment?: (a: Appointment) => boolean;
}

export function CalendarGrid({
  view,
  referenceDate,
  appointments,
  onSelectAppointment,
  colorByDoctor = false,
  onSlotClick,
  enableDragDrop = false,
  canDragAppointment,
}: CalendarGridProps) {
  const days = useMemo(() => {
    if (view === 'day') return [startOfDay(referenceDate)];
    // Week starts Sunday
    const weekStart = startOfDay(addDays(referenceDate, -referenceDate.getDay()));
    return Array.from({ length: 7 }, (_, i) => addDays(weekStart, i));
  }, [view, referenceDate]);

  const today = startOfDay(new Date());

  const reschedule = useRescheduleAppointment();
  // Flag que "engole" o clique sintético disparado logo após um arrasto, pra
  // não abrir o detalhe (clique no bloco) nem criar consulta (clique no vazio).
  const justDraggedRef = useRef(false);

  const sensors = useSensors(
    // Desktop: 6px separa clique (abre detalhe / cria) de arrasto (reagenda).
    useSensor(MouseSensor, { activationConstraint: { distance: 6 } }),
    // Touch: press-and-hold (250ms) — assim um flick de scroll sobre um bloco
    // rola a página normalmente e só vira arrasto num toque deliberado.
    useSensor(TouchSensor, { activationConstraint: { delay: 250, tolerance: 8 } }),
  );

  function handleDragEnd(event: DragEndEvent) {
    justDraggedRef.current = true;
    // Reseta após o ciclo de eventos atual (o click sintético vem logo depois).
    window.setTimeout(() => {
      justDraggedRef.current = false;
    }, 0);

    const { active, over, delta } = event;
    const appt = active.data.current?.appt as Appointment | undefined;
    if (!appt) return;

    const originalStart = new Date(appt.scheduledAt);

    // delta.y (px) → minutos.
    const origMinutesOfDay = originalStart.getHours() * 60 + originalStart.getMinutes();
    const rawShift = (delta.y / HOUR_HEIGHT_PX) * 60;

    // Dia de destino: coluna sob o drop (over); senão mantém o dia original.
    const overYmd = over?.data.current?.ymd as string | undefined;
    const sameDay = !overYmd || overYmd === format(originalStart, 'yyyy-MM-dd');

    // No-op: arrasto curto sem mudar de dia e sem deslocamento real (≥ ½ slot).
    // Cobre também consultas fora da grade de 15min (ex.: 10:07) — um clique que
    // passou de 6px não deve "normalizar" o horário silenciosamente.
    const snappedShiftMin = Math.round(rawShift / SNAP_MINUTES) * SNAP_MINUTES;
    if (sameDay && snappedShiftMin === 0) return;

    // Snap no novo horário absoluto (00/15/30/45) e clampa na grade visível.
    let newMinutesOfDay = Math.round((origMinutesOfDay + rawShift) / SNAP_MINUTES) * SNAP_MINUTES;
    const minMinute = HOUR_START * 60;
    const maxMinute = HOUR_END * 60 - appt.durationMinutes;
    newMinutesOfDay = Math.max(minMinute, Math.min(newMinutesOfDay, Math.max(minMinute, maxMinute)));

    let newDate: Date;
    if (overYmd) {
      const [y, m, d] = overYmd.split('-').map(Number);
      newDate = new Date(y, m - 1, d, 0, 0, 0, 0);
    } else {
      newDate = startOfDay(originalStart);
    }
    newDate.setHours(Math.floor(newMinutesOfDay / 60), newMinutesOfDay % 60, 0, 0);

    // No-op: não mudou nada.
    if (newDate.getTime() === originalStart.getTime()) return;

    reschedule.mutate(
      { id: appt.id, scheduledAt: newDate.toISOString() },
      {
        onSuccess: () => {
          toast.success('Consulta reagendada', {
            description: format(newDate, "EEE, dd/MM 'às' HH:mm", { locale: ptBR }),
          });
        },
        onError: (err) => {
          const status = (err as { status?: number }).status;
          toast.error('Erro ao reagendar', {
            description:
              status === 409
                ? 'Esse horário já foi reservado para o médico. Escolha outro.'
                : status === 403
                  ? 'Você não pode reagendar a consulta de outro profissional.'
                  : err instanceof Error
                    ? err.message
                    : undefined,
          });
        },
      },
    );
  }

  return (
    <DndContext sensors={sensors} onDragEnd={handleDragEnd}>
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

          {/* Hours column */}
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

          {/* Day columns */}
          {days.map((d) => {
            const dayAppts = appointments.filter((a) =>
              isSameDay(new Date(a.scheduledAt), d),
            );
            return (
              <DayColumn
                key={d.toISOString()}
                day={d}
                appointments={dayAppts}
                colorByDoctor={colorByDoctor}
                onSelectAppointment={onSelectAppointment}
                onSlotClick={onSlotClick}
                enableDragDrop={enableDragDrop}
                canDragAppointment={canDragAppointment}
                justDraggedRef={justDraggedRef}
              />
            );
          })}
        </div>
      </div>
    </DndContext>
  );
}

interface DayColumnProps {
  day: Date;
  appointments: Appointment[];
  colorByDoctor: boolean;
  onSelectAppointment: (a: Appointment) => void;
  onSlotClick?: (dateYmd: string) => void;
  enableDragDrop: boolean;
  canDragAppointment?: (a: Appointment) => boolean;
  justDraggedRef: MutableRefObject<boolean>;
}

function DayColumn({
  day,
  appointments,
  colorByDoctor,
  onSelectAppointment,
  onSlotClick,
  enableDragDrop,
  canDragAppointment,
  justDraggedRef,
}: DayColumnProps) {
  const ymd = format(day, 'yyyy-MM-dd');
  const { setNodeRef, isOver } = useDroppable({ id: `day-${ymd}`, data: { ymd } });

  return (
    <div
      ref={setNodeRef}
      className={cn(
        'relative border-r',
        onSlotClick && 'cursor-pointer',
        isOver && 'bg-primary/5',
      )}
      style={{ height: TOTAL_HEIGHT }}
      onClick={() => {
        if (justDraggedRef.current) return; // engole o clique pós-arrasto
        onSlotClick?.(ymd);
      }}
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
      {appointments.map((a) => (
        <AppointmentBlock
          key={a.id}
          appointment={a}
          colorByDoctor={colorByDoctor}
          draggable={
            enableDragDrop &&
            canDragStatus(a.status) &&
            (canDragAppointment ? canDragAppointment(a) : true)
          }
          onSelect={onSelectAppointment}
          justDraggedRef={justDraggedRef}
        />
      ))}
    </div>
  );
}

interface AppointmentBlockProps {
  appointment: Appointment;
  colorByDoctor: boolean;
  draggable: boolean;
  onSelect: (a: Appointment) => void;
  justDraggedRef: MutableRefObject<boolean>;
}

function AppointmentBlock({
  appointment: a,
  colorByDoctor,
  draggable,
  onSelect,
  justDraggedRef,
}: AppointmentBlockProps) {
  const { attributes, listeners, setNodeRef, transform, isDragging } = useDraggable({
    id: a.id,
    data: { appt: a },
    disabled: !draggable,
  });

  const start = new Date(a.scheduledAt);
  const startMinutesFromMidnight = start.getHours() * 60 + start.getMinutes();
  const startMinutesFromGrid = startMinutesFromMidnight - HOUR_START * 60;
  if (startMinutesFromGrid < 0) return null;
  const top = (startMinutesFromGrid / 60) * HOUR_HEIGHT_PX;
  const height = Math.max(20, (a.durationMinutes / 60) * HOUR_HEIGHT_PX - 2);
  const baseClass = colorByDoctor
    ? cn('border-l-4', doctorBlockClass(a.doctorId))
    : APPOINTMENT_TYPE_COLORS[a.type];
  const patientName = a.patient?.name ?? 'Paciente';
  const typeLabel = APPOINTMENT_TYPE_LABELS[a.type];
  const doctorName = a.doctor?.name;
  const titleParts = [`${format(start, 'HH:mm')} ${patientName}`, typeLabel];
  if (doctorName) titleParts.push(`Dr(a). ${doctorName}`);

  // Só conecta listeners/ref do dnd-kit quando o bloco é arrastável.
  const dragHandlers = draggable ? { ...listeners, ...attributes } : {};

  return (
    <button
      ref={draggable ? setNodeRef : undefined}
      type="button"
      {...dragHandlers}
      onClick={(e) => {
        e.stopPropagation();
        if (justDraggedRef.current) return; // não abre detalhe após arrastar
        onSelect(a);
      }}
      className={cn(
        'absolute left-1 right-1 overflow-hidden rounded border px-1.5 py-1 text-left text-xs shadow-sm transition-colors',
        baseClass,
        a.status === 'cancelled' && 'opacity-50 line-through',
        draggable && 'cursor-grab active:cursor-grabbing',
        isDragging && 'z-30 cursor-grabbing opacity-90 shadow-lg ring-2 ring-primary/50',
      )}
      style={{ top, height, transform: CSS.Translate.toString(transform) }}
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
}
