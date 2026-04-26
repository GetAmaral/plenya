import { useMemo, useState } from 'react';
import {
  FlatList,
  Pressable,
  RefreshControl,
  ScrollView,
  View,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link, router } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import {
  options,
  appointmentTypeLabels,
  type Appointment,
  type AppointmentStatus,
  type AppointmentType,
} from '@plenya/api-client';
import {
  Button,
  Card,
  EmptyState,
  ErrorState,
  Sheet,
  Spinner,
  Text,
} from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

type RangeMode = 'today' | 'week';
type ViewMode = 'list' | 'grid';

function startOfToday(): Date {
  const d = new Date();
  d.setHours(0, 0, 0, 0);
  return d;
}

function endOfRange(mode: RangeMode): Date {
  const d = startOfToday();
  if (mode === 'today') {
    d.setHours(23, 59, 59, 999);
  } else {
    d.setDate(d.getDate() + 7);
    d.setHours(23, 59, 59, 999);
  }
  return d;
}

const STATUS_LABEL: Record<AppointmentStatus, string> = {
  scheduled: 'Agendada',
  confirmed: 'Confirmada',
  completed: 'Concluída',
  cancelled: 'Cancelada',
  no_show: 'Faltou',
};

const STATUS_BG: Record<AppointmentStatus, string> = {
  scheduled: 'bg-secondary',
  confirmed: 'bg-emerald-600',
  completed: 'bg-muted',
  cancelled: 'bg-destructive',
  no_show: 'bg-amber-500',
};

const STATUS_TEXT: Record<AppointmentStatus, string> = {
  scheduled: 'text-secondary-foreground',
  confirmed: 'text-white',
  completed: 'text-muted-foreground',
  cancelled: 'text-white',
  no_show: 'text-white',
};

const TYPE_ICON: Record<AppointmentType, string> = {
  initial_assessment: '◔',
  follow_up: '↻',
  telemedicine: '📹',
  procedure: '⚙',
  results_review: '📋',
};

const HOUR_START = 7; // 7h
const HOUR_END = 21; // 21h
const PX_PER_HOUR = 60;

export default function AgendaScreen() {
  const [mode, setMode] = useState<RangeMode>('today');
  const [view, setView] = useState<ViewMode>('list');
  const [doctorPickerOpen, setDoctorPickerOpen] = useState(false);
  const [selectedDoctorIds, setSelectedDoctorIds] = useState<string[]>([]);

  const me = useQuery(options.meOptions());
  const doctors = useQuery({
    ...options.doctorsListOptions(),
    enabled: doctorPickerOpen,
  });

  const fromDate = useMemo(() => startOfToday(), []);
  const toDate = useMemo(() => endOfRange(mode), [mode]);

  const query = useQuery(
    options.appointmentsListOptions({
      limit: 200,
      dateFrom: fromDate.toISOString(),
      dateTo: toDate.toISOString(),
      doctorIds: selectedDoctorIds.length > 0 ? selectedDoctorIds : undefined,
    }),
  );

  const filtered = useMemo(() => {
    return (query.data ?? [])
      .slice()
      .sort((a, b) => a.scheduledAt.localeCompare(b.scheduledAt));
  }, [query.data]);

  const days = useMemo(() => {
    const groups = new Map<string, Appointment[]>();
    filtered.forEach((a) => {
      const key = a.scheduledAt.slice(0, 10);
      if (!groups.has(key)) groups.set(key, []);
      groups.get(key)!.push(a);
    });
    return Array.from(groups.entries()).sort(([a], [b]) => a.localeCompare(b));
  }, [filtered]);

  const doctorFilterLabel =
    selectedDoctorIds.length === 0
      ? 'Todos profissionais'
      : `${selectedDoctorIds.length} profissional${selectedDoctorIds.length === 1 ? '' : 'is'}`;

  const meIsDoctor = me.data?.role === 'doctor' || me.data?.role === 'nurse';

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="px-4 pt-3">
        <Button
          onPress={() => router.push('/(tabs)/agenda/new' as never)}
          fullWidth
          size="sm"
        >
          + Nova consulta
        </Button>
      </View>

      <View className="flex-row gap-2 px-4 pt-3">
        <Pressable
          onPress={() => setMode('today')}
          className={`flex-1 items-center rounded-full py-2 ${
            mode === 'today' ? 'bg-primary' : 'bg-muted'
          }`}
        >
          <Text
            className={`text-sm font-semibold ${
              mode === 'today' ? 'text-primary-foreground' : 'text-foreground'
            }`}
          >
            Hoje
          </Text>
        </Pressable>
        <Pressable
          onPress={() => setMode('week')}
          className={`flex-1 items-center rounded-full py-2 ${
            mode === 'week' ? 'bg-primary' : 'bg-muted'
          }`}
        >
          <Text
            className={`text-sm font-semibold ${
              mode === 'week' ? 'text-primary-foreground' : 'text-foreground'
            }`}
          >
            7 dias
          </Text>
        </Pressable>
      </View>

      <View className="flex-row gap-2 px-4 pt-2">
        <Pressable
          onPress={() => setView(view === 'list' ? 'grid' : 'list')}
          className="flex-1 rounded-full bg-muted py-1.5"
        >
          <Text className="text-center text-xs font-semibold">
            {view === 'list' ? '📋 Lista' : '🗓 Grade'} (toque pra alternar)
          </Text>
        </Pressable>
        {!meIsDoctor && (
          <Pressable
            onPress={() => setDoctorPickerOpen(true)}
            className="flex-1 rounded-full bg-muted py-1.5"
          >
            <Text className="text-center text-xs font-semibold">
              👤 {doctorFilterLabel}
            </Text>
          </Pressable>
        )}
      </View>

      {query.isLoading ? (
        <Spinner centered />
      ) : query.isError ? (
        <ErrorState onRetry={() => query.refetch()} />
      ) : filtered.length === 0 ? (
        <EmptyState
          title="Sem consultas"
          description={
            mode === 'today'
              ? 'Você não tem consultas hoje.'
              : 'Nenhuma consulta nos próximos 7 dias.'
          }
        />
      ) : view === 'list' ? (
        <FlatList
          data={filtered}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8 pt-3"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
          renderItem={({ item }) => <AppointmentRow item={item} />}
        />
      ) : (
        <ScrollView
          className="flex-1"
          contentContainerClassName="gap-4 px-4 pb-8 pt-3"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
        >
          {days.map(([dateKey, dayItems]) => (
            <DayGrid key={dateKey} dateKey={dateKey} items={dayItems} />
          ))}
        </ScrollView>
      )}

      <Sheet open={doctorPickerOpen} onClose={() => setDoctorPickerOpen(false)}>
        <Text variant="title" className="mb-2">
          Filtrar por profissional
        </Text>
        <Pressable
          onPress={() => setSelectedDoctorIds([])}
          className={`mb-2 rounded-md px-3 py-2 ${
            selectedDoctorIds.length === 0 ? 'bg-primary' : 'bg-muted'
          }`}
        >
          <Text
            className={`font-semibold ${
              selectedDoctorIds.length === 0 ? 'text-primary-foreground' : 'text-foreground'
            }`}
          >
            Todos profissionais
          </Text>
        </Pressable>
        {(doctors.data ?? []).map((d) => {
          const checked = selectedDoctorIds.includes(d.id);
          return (
            <Pressable
              key={d.id}
              onPress={() =>
                setSelectedDoctorIds((prev) =>
                  prev.includes(d.id) ? prev.filter((x) => x !== d.id) : [...prev, d.id],
                )
              }
              className={`mb-1 flex-row items-center justify-between rounded-md px-3 py-2 ${
                checked ? 'bg-primary' : 'bg-muted'
              }`}
            >
              <View>
                <Text
                  variant="body"
                  className={checked ? 'text-primary-foreground' : 'text-foreground'}
                >
                  {d.name}
                </Text>
                <Text
                  variant="caption"
                  className={checked ? 'text-primary-foreground' : 'text-muted-foreground'}
                >
                  {d.email}
                </Text>
              </View>
              {checked && (
                <Text className="text-base text-primary-foreground">✓</Text>
              )}
            </Pressable>
          );
        })}
        <Button onPress={() => setDoctorPickerOpen(false)} fullWidth className="mt-3">
          Aplicar
        </Button>
      </Sheet>
    </SafeAreaView>
  );
}

function AppointmentRow({ item }: { item: Appointment }) {
  return (
    <Link href={`/(tabs)/agenda/${item.id}`} asChild>
      <Card>
        <View className="flex-row items-center justify-between">
          <View className="flex-1 pr-2">
            <Text variant="title">
              {TYPE_ICON[item.type]} {item.patientName ?? 'Paciente'}
            </Text>
            <Text variant="caption">
              {formatDateTime(item.scheduledAt)} · {appointmentTypeLabels[item.type]} ·{' '}
              {item.durationMinutes}min
              {item.doctorName ? ` · ${item.doctorName}` : ''}
            </Text>
            {item.reason && (
              <Text variant="caption" className="mt-0.5 italic">
                {item.reason}
              </Text>
            )}
          </View>
          <View className={`rounded-full px-2.5 py-0.5 ${STATUS_BG[item.status]}`}>
            <Text className={`text-xs font-semibold ${STATUS_TEXT[item.status]}`}>
              {STATUS_LABEL[item.status]}
            </Text>
          </View>
        </View>
      </Card>
    </Link>
  );
}

function DayGrid({ dateKey, items }: { dateKey: string; items: Appointment[] }) {
  const dayLabel = useMemo(() => {
    const d = new Date(`${dateKey}T12:00:00`);
    return d.toLocaleDateString('pt-BR', {
      weekday: 'long',
      day: 'numeric',
      month: 'long',
    });
  }, [dateKey]);

  const totalHours = HOUR_END - HOUR_START;
  const containerHeight = totalHours * PX_PER_HOUR;

  return (
    <View>
      <Text variant="title" className="mb-2 capitalize">
        {dayLabel}
      </Text>
      <View
        className="rounded-lg border border-border bg-card"
        style={{ height: containerHeight }}
      >
        {Array.from({ length: totalHours + 1 }).map((_, i) => {
          const hour = HOUR_START + i;
          return (
            <View
              key={hour}
              className="flex-row border-b border-border"
              style={{
                position: 'absolute',
                top: i * PX_PER_HOUR,
                left: 0,
                right: 0,
                height: 1,
              }}
            >
              <Text
                variant="caption"
                style={{ position: 'absolute', left: 4, top: -7, backgroundColor: 'transparent' }}
              >
                {String(hour).padStart(2, '0')}:00
              </Text>
            </View>
          );
        })}

        {items.map((a) => {
          const start = new Date(a.scheduledAt);
          const startHour = start.getHours() + start.getMinutes() / 60;
          if (startHour < HOUR_START || startHour > HOUR_END) return null;
          const top = (startHour - HOUR_START) * PX_PER_HOUR;
          const height = Math.max((a.durationMinutes / 60) * PX_PER_HOUR, 28);
          return (
            <Link key={a.id} href={`/(tabs)/agenda/${a.id}`} asChild>
              <Pressable
                style={{
                  position: 'absolute',
                  top,
                  left: 56,
                  right: 8,
                  height,
                }}
                className={`overflow-hidden rounded-md px-2 py-1 ${
                  STATUS_BG[a.status]
                } border border-border`}
              >
                <Text
                  className={`text-xs font-semibold ${STATUS_TEXT[a.status]}`}
                  numberOfLines={1}
                >
                  {TYPE_ICON[a.type]} {a.patientName ?? 'Paciente'}
                </Text>
                <Text
                  className={`text-[10px] ${STATUS_TEXT[a.status]}`}
                  numberOfLines={1}
                >
                  {String(start.getHours()).padStart(2, '0')}:
                  {String(start.getMinutes()).padStart(2, '0')} ·{' '}
                  {appointmentTypeLabels[a.type]}
                </Text>
              </Pressable>
            </Link>
          );
        })}
      </View>
    </View>
  );
}
