import { useMemo, useState } from 'react';
import { FlatList, Pressable, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import {
  options,
  appointmentTypeLabels,
  type Appointment,
  type AppointmentStatus,
  type AppointmentType,
} from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

type RangeMode = 'today' | 'week';

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

export default function AgendaScreen() {
  const [mode, setMode] = useState<RangeMode>('today');
  const query = useQuery(options.appointmentsListOptions({ limit: 200 }));

  const filtered = useMemo(() => {
    const fromMs = startOfToday().getTime();
    const toMs = endOfRange(mode).getTime();
    const items = (query.data ?? []).filter((a) => {
      const ts = new Date(a.scheduledAt).getTime();
      return ts >= fromMs && ts <= toMs;
    });
    return items.sort((a, b) => a.scheduledAt.localeCompare(b.scheduledAt));
  }, [query.data, mode]);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <View className="flex-row gap-2 p-4">
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

      {query.isLoading ? (
        <Spinner centered />
      ) : query.isError ? (
        <ErrorState onRetry={() => query.refetch()} />
      ) : (
        <FlatList
          data={filtered}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
          ListEmptyComponent={
            <EmptyState
              title="Sem consultas"
              description={
                mode === 'today'
                  ? 'Você não tem consultas hoje.'
                  : 'Nenhuma consulta nos próximos 7 dias.'
              }
            />
          }
          renderItem={({ item }) => <AppointmentRow item={item} />}
        />
      )}
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
