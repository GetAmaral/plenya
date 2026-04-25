'use client';

/**
 * /calendario
 *
 * Visão estilo Google Calendar: Day / Week (default) / Month placeholder.
 * - Doctor logado vê própria agenda; admin/secretary/manager: dropdown.
 * - Click em bloco → drawer com info + link pra detail page.
 * - Navegação prev/next + "Hoje".
 *
 * Polling 15s (via useAppointments) mantém grid atualizado em real-time.
 */
import { useEffect, useMemo, useState } from 'react';
import { useRouter } from 'next/navigation';
import { addDays, format, startOfDay } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { ChevronLeft, ChevronRight, Plus } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { useAppointments, useDoctors } from '@/lib/api/calendar-api';

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { PageHeader } from '@/components/layout/page-header';

import { CalendarGrid } from '@/components/calendar/calendar-grid';
import {
  APPOINTMENT_STATUS_COLORS,
  APPOINTMENT_STATUS_LABELS,
  APPOINTMENT_TYPE_LABELS,
  type Appointment,
} from '@/lib/api/calendar-api';

type ViewMode = 'day' | 'week';

export default function CalendarioPage() {
  useRequireAuth();
  const router = useRouter();
  const { user } = useAuthStore();

  const isAdminOrManager =
    isGranted(user, 'admin') ||
    isGranted(user, 'manager') ||
    isGranted(user, 'secretary');
  const isDoctor = isGranted(user, 'doctor');

  const { data: doctors } = useDoctors();

  const [view, setView] = useState<ViewMode>('week');
  const [reference, setReference] = useState<Date>(startOfDay(new Date()));
  const [selectedDoctorId, setSelectedDoctorId] = useState<string | undefined>(undefined);
  const [drawerAppt, setDrawerAppt] = useState<Appointment | null>(null);

  // Default doctor filter
  useEffect(() => {
    if (selectedDoctorId !== undefined) return;
    if (isDoctor && user?.id) {
      setSelectedDoctorId(user.id);
    } else if (isAdminOrManager) {
      setSelectedDoctorId(''); // 'todos'
    }
  }, [isDoctor, isAdminOrManager, user?.id, selectedDoctorId]);

  // Janela de fetch
  const { dateFrom, dateTo } = useMemo(() => {
    if (view === 'day') {
      const start = startOfDay(reference);
      return { dateFrom: start.toISOString(), dateTo: addDays(start, 1).toISOString() };
    }
    // week (Sunday → Saturday)
    const start = startOfDay(addDays(reference, -reference.getDay()));
    return { dateFrom: start.toISOString(), dateTo: addDays(start, 7).toISOString() };
  }, [view, reference]);

  const { data: appointments = [], isLoading } = useAppointments({
    doctorId: selectedDoctorId || undefined,
    dateFrom,
    dateTo,
    limit: 200,
  });

  const handlePrev = () => {
    setReference(view === 'day' ? addDays(reference, -1) : addDays(reference, -7));
  };
  const handleNext = () => {
    setReference(view === 'day' ? addDays(reference, 1) : addDays(reference, 7));
  };
  const handleToday = () => setReference(startOfDay(new Date()));

  const headerLabel = useMemo(() => {
    if (view === 'day') {
      return format(reference, "EEEE, dd 'de' MMM yyyy", { locale: ptBR });
    }
    const start = addDays(reference, -reference.getDay());
    const end = addDays(start, 6);
    return `${format(start, 'dd MMM', { locale: ptBR })} - ${format(end, 'dd MMM yyyy', { locale: ptBR })}`;
  }, [view, reference]);

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[{ label: 'Calendário' }]}
        title="Calendário"
        description={headerLabel}
        actions={[
          {
            label: 'Nova consulta',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => router.push('/appointments/new'),
            variant: 'default',
          },
        ]}
      />

      {/* Toolbar */}
      <div className="flex flex-wrap items-center gap-3">
        <div className="flex items-center gap-1 rounded-md border p-0.5">
          <Button
            size="sm"
            variant={view === 'day' ? 'default' : 'ghost'}
            onClick={() => setView('day')}
          >
            Dia
          </Button>
          <Button
            size="sm"
            variant={view === 'week' ? 'default' : 'ghost'}
            onClick={() => setView('week')}
          >
            Semana
          </Button>
        </div>

        <div className="flex items-center gap-1">
          <Button variant="outline" size="sm" onClick={handlePrev}>
            <ChevronLeft className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="sm" onClick={handleToday}>
            Hoje
          </Button>
          <Button variant="outline" size="sm" onClick={handleNext}>
            <ChevronRight className="h-4 w-4" />
          </Button>
        </div>

        {isAdminOrManager && (
          <div className="ml-auto w-full sm:w-64">
            <Select
              value={selectedDoctorId ?? ''}
              onValueChange={(v) => setSelectedDoctorId(v === '__all' ? '' : v)}
            >
              <SelectTrigger>
                <SelectValue placeholder="Filtrar médico" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="__all">Todos os médicos</SelectItem>
                {(doctors ?? []).map((d) => (
                  <SelectItem key={d.id} value={d.id}>
                    {d.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        )}
      </div>

      {/* Grid */}
      {isLoading ? (
        <Card className="h-[600px] animate-pulse" />
      ) : (
        <CalendarGrid
          view={view}
          referenceDate={reference}
          appointments={appointments}
          onSelectAppointment={(a) => setDrawerAppt(a)}
        />
      )}

      {/* Drawer */}
      <Sheet open={!!drawerAppt} onOpenChange={(v) => (!v ? setDrawerAppt(null) : null)}>
        <SheetContent>
          {drawerAppt && (
            <>
              <SheetHeader>
                <SheetTitle>{drawerAppt.patient?.name ?? 'Consulta'}</SheetTitle>
                <SheetDescription>
                  {APPOINTMENT_TYPE_LABELS[drawerAppt.type]}
                </SheetDescription>
              </SheetHeader>
              <div className="mt-6 space-y-4 text-sm">
                <Badge
                  variant="outline"
                  className={APPOINTMENT_STATUS_COLORS[drawerAppt.status]}
                >
                  {APPOINTMENT_STATUS_LABELS[drawerAppt.status]}
                </Badge>
                <div>
                  <p className="text-xs uppercase text-muted-foreground">Data/Hora</p>
                  <p>
                    {format(new Date(drawerAppt.scheduledAt), "dd 'de' MMM yyyy 'às' HH:mm", {
                      locale: ptBR,
                    })}
                  </p>
                </div>
                <div>
                  <p className="text-xs uppercase text-muted-foreground">Duração</p>
                  <p>{drawerAppt.durationMinutes} min</p>
                </div>
                {drawerAppt.doctor && (
                  <div>
                    <p className="text-xs uppercase text-muted-foreground">Médico</p>
                    <p>{drawerAppt.doctor.name}</p>
                  </div>
                )}
                {drawerAppt.reason && (
                  <div>
                    <p className="text-xs uppercase text-muted-foreground">Motivo</p>
                    <p>{drawerAppt.reason}</p>
                  </div>
                )}
                <Button
                  className="w-full"
                  onClick={() => router.push(`/appointments/${drawerAppt.id}`)}
                >
                  Abrir página completa
                </Button>
              </div>
            </>
          )}
        </SheetContent>
      </Sheet>
    </div>
  );
}

