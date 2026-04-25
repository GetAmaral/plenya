import { useMemo } from 'react';
import { Pressable, RefreshControl, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import {
  appointmentTypeLabels,
  options,
  queryKeys,
  type Appointment,
} from '@plenya/api-client';
import { Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { formatDateTime, initials } from '@plenya/domain';

function isUpcoming(a: Appointment): boolean {
  if (a.status === 'cancelled' || a.status === 'completed' || a.status === 'no_show') {
    return false;
  }
  return new Date(a.scheduledAt).getTime() >= Date.now() - 60 * 60 * 1000;
}

export default function DashboardScreen() {
  const me = useQuery(options.meOptions());
  const appts = useQuery(options.appointmentsListOptions({ limit: 100 }));
  const leadsDash = useQuery(options.leadsDashboardOptions());
  const unread = useQuery(options.unreadCountOptions());
  const queryClient = useQueryClient();

  const nextAppt = useMemo(() => {
    const list = (appts.data ?? [])
      .filter(isUpcoming)
      .sort((a, b) => a.scheduledAt.localeCompare(b.scheduledAt));
    return list[0];
  }, [appts.data]);

  const greeting = useMemo(() => {
    const hour = new Date().getHours();
    if (hour < 12) return 'Bom dia';
    if (hour < 18) return 'Boa tarde';
    return 'Boa noite';
  }, []);

  const refreshing =
    (appts.isRefetching && !appts.isLoading) ||
    (leadsDash.isRefetching && !leadsDash.isLoading);

  function refreshAll() {
    queryClient.invalidateQueries({ queryKey: queryKeys.appointments.all() });
    queryClient.invalidateQueries({ queryKey: queryKeys.leads.dashboard() });
    queryClient.invalidateQueries({ queryKey: queryKeys.notifications.unread() });
  }

  if (me.isLoading) return <Spinner centered />;

  const firstName = me.data?.name?.split(' ')[0] ?? 'doutor(a)';

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView
        contentContainerClassName="gap-4 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={refreshAll} />}
      >
        <View>
          <Text variant="heading">
            {greeting}, {firstName}
          </Text>
          <Text variant="caption">{new Date().toLocaleDateString('pt-BR', {
            weekday: 'long',
            day: 'numeric',
            month: 'long',
          })}</Text>
        </View>

        <Pressable onPress={() => router.push('/(tabs)/notifications')}>
          <Card>
            <View className="flex-row items-center justify-between">
              <View className="flex-1">
                <Text variant="title">Notificações</Text>
                <Text variant="caption">
                  {unread.data?.count
                    ? `${unread.data.count} não lida${unread.data.count === 1 ? '' : 's'}`
                    : 'Nada novo'}
                </Text>
              </View>
              {(unread.data?.count ?? 0) > 0 && (
                <View className="h-7 min-w-7 items-center justify-center rounded-full bg-primary px-2">
                  <Text className="text-xs font-semibold text-primary-foreground">
                    {unread.data!.count > 99 ? '99+' : unread.data!.count}
                  </Text>
                </View>
              )}
            </View>
          </Card>
        </Pressable>

        <Pressable
          onPress={() =>
            nextAppt
              ? router.push(`/(tabs)/agenda/${nextAppt.id}`)
              : router.push('/(tabs)/agenda')
          }
        >
          <Card>
            <CardHeader>
              <Text variant="title">Próxima consulta</Text>
            </CardHeader>
            {nextAppt ? (
              <View>
                <Text variant="body">{nextAppt.patientName ?? 'Paciente'}</Text>
                <Text variant="caption">
                  {formatDateTime(nextAppt.scheduledAt)} ·{' '}
                  {appointmentTypeLabels[nextAppt.type]}
                </Text>
              </View>
            ) : (
              <Text variant="caption">Sem consultas no horizonte.</Text>
            )}
          </Card>
        </Pressable>

        {me.data?.selectedPatientId && (
          <Pressable
            onPress={() =>
              router.push(`/(tabs)/patients/${me.data!.selectedPatientId}` as never)
            }
          >
            <Card>
              <CardHeader>
                <Text variant="title">Paciente em atendimento</Text>
              </CardHeader>
              <View className="flex-row items-center gap-3">
                <View className="h-12 w-12 items-center justify-center rounded-full bg-primary">
                  <Text className="text-base font-semibold text-primary-foreground">
                    {initials(me.data?.selectedPatient?.name ?? 'P')}
                  </Text>
                </View>
                <View className="flex-1">
                  <Text variant="body">
                    {me.data?.selectedPatient?.name ?? 'Paciente selecionado'}
                  </Text>
                  <Text variant="caption">Toque pra abrir o prontuário</Text>
                </View>
              </View>
            </Card>
          </Pressable>
        )}

        <Pressable onPress={() => router.push('/(tabs)/leads')}>
          <Card>
            <CardHeader>
              <Text variant="title">Leads</Text>
            </CardHeader>
            {leadsDash.data ? (
              <View>
                <Text variant="body">
                  {leadsDash.data.newInLast24h} nova{leadsDash.data.newInLast24h === 1 ? '' : 's'} em 24h
                </Text>
                <Text variant="caption">
                  {Object.entries(leadsDash.data.bySource)
                    .filter(([, v]) => v > 0)
                    .map(([k, v]) => `${v} ${k}`)
                    .join(' · ') || 'Sem leads na fila'}
                </Text>
              </View>
            ) : (
              <Text variant="caption">Carregando…</Text>
            )}
          </Card>
        </Pressable>
      </ScrollView>
    </SafeAreaView>
  );
}
