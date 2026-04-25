import { useMemo, useState } from 'react';
import { FlatList, Pressable, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { options, type Appointment } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

type RangeMode = 'today' | 'week';

function toIsoDate(d: Date): string {
  return d.toISOString();
}

function startOfToday(): Date {
  const d = new Date();
  d.setHours(0, 0, 0, 0);
  return d;
}

function endOfToday(): Date {
  const d = new Date();
  d.setHours(23, 59, 59, 999);
  return d;
}

function endOfWeek(): Date {
  const d = startOfToday();
  d.setDate(d.getDate() + 7);
  d.setHours(23, 59, 59, 999);
  return d;
}

const STATUS_LABEL: Record<Appointment['status'], string> = {
  scheduled: 'Agendada',
  confirmed: 'Confirmada',
  completed: 'Concluída',
  cancelled: 'Cancelada',
  no_show: 'Faltou',
};

const STATUS_COLOR: Record<Appointment['status'], string> = {
  scheduled: 'bg-secondary',
  confirmed: 'bg-emerald-600',
  completed: 'bg-muted',
  cancelled: 'bg-destructive',
  no_show: 'bg-amber-500',
};

const STATUS_TEXT_COLOR: Record<Appointment['status'], string> = {
  scheduled: 'text-secondary-foreground',
  confirmed: 'text-white',
  completed: 'text-muted-foreground',
  cancelled: 'text-white',
  no_show: 'text-white',
};

export default function AgendaScreen() {
  const [mode, setMode] = useState<RangeMode>('today');

  const { from, to } = useMemo(() => {
    if (mode === 'today') {
      return { from: toIsoDate(startOfToday()), to: toIsoDate(endOfToday()) };
    }
    return { from: toIsoDate(startOfToday()), to: toIsoDate(endOfWeek()) };
  }, [mode]);

  const query = useQuery(options.appointmentsByRangeOptions(from, to));

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
          data={(query.data ?? []).slice().sort((a, b) => a.startAt.localeCompare(b.startAt))}
          keyExtractor={(item) => item.id}
          contentContainerClassName="gap-2 px-4 pb-8"
          refreshControl={
            <RefreshControl refreshing={query.isRefetching} onRefresh={() => query.refetch()} />
          }
          ListEmptyComponent={
            <EmptyState
              title="Sem consultas"
              description={
                mode === 'today' ? 'Você não tem consultas hoje.' : 'Nenhuma consulta nos próximos 7 dias.'
              }
            />
          }
          renderItem={({ item }) => (
            <Card>
              <View className="flex-row items-center justify-between">
                <Text variant="title">{item.patientName}</Text>
                <View className={`rounded-full px-2.5 py-0.5 ${STATUS_COLOR[item.status]}`}>
                  <Text className={`text-xs font-semibold ${STATUS_TEXT_COLOR[item.status]}`}>
                    {STATUS_LABEL[item.status]}
                  </Text>
                </View>
              </View>
              <Text variant="caption">
                {formatDateTime(item.startAt)}
                {item.kind ? ` · ${item.kind}` : ''}
              </Text>
              {item.notes && (
                <Text variant="caption" className="mt-1 italic">
                  {item.notes}
                </Text>
              )}
            </Card>
          )}
        />
      )}
    </SafeAreaView>
  );
}
