import { useMemo } from 'react';
import { Pressable, RefreshControl, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

/**
 * Home do app paciente — Sprint 1 mostra próxima consulta + check-in CTA.
 * Continuum, último score e último exame entram na Sprint 2.
 */
export default function HomeScreen() {
  const profile = useQuery(options.patientMeProfileOptions());
  const appts = useQuery(options.patientMeAppointmentsOptions('upcoming'));
  const checkIn = useQuery(options.patientMeCheckInTodayOptions());
  const qc = useQueryClient();

  const nextAppt = useMemo(() => (appts.data ?? [])[0], [appts.data]);

  const greeting = useMemo(() => {
    const h = new Date().getHours();
    if (h < 12) return 'Bom dia';
    if (h < 18) return 'Boa tarde';
    return 'Boa noite';
  }, []);

  const refreshing =
    (appts.isRefetching && !appts.isLoading) ||
    (checkIn.isRefetching && !checkIn.isLoading);

  function refreshAll() {
    qc.invalidateQueries({ queryKey: queryKeys.patientMe.all() });
  }

  if (profile.isLoading) return <Spinner centered />;

  const firstName = profile.data?.name?.split(' ')[0] ?? '';
  const todayCheckedIn = Boolean(checkIn.data?.checkIn);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView
        contentContainerClassName="gap-4 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={refreshAll} />}
      >
        <View>
          <Text variant="heading">
            {greeting}, {firstName || 'paciente'}
          </Text>
          <Text variant="caption">
            {new Date().toLocaleDateString('pt-BR', {
              weekday: 'long',
              day: 'numeric',
              month: 'long',
            })}
          </Text>
        </View>

        <Pressable onPress={() => router.push('/check-in' as never)}>
          <Card>
            <CardHeader>
              <Text variant="title">
                {todayCheckedIn ? 'Check-in feito ✓' : 'Como você está hoje?'}
              </Text>
            </CardHeader>
            <Text variant="caption">
              {todayCheckedIn
                ? 'Toque pra fazer outro check-in se quiser.'
                : 'Em ~10s seu profissional acompanha sua semana.'}
            </Text>
          </Card>
        </Pressable>

        <Pressable onPress={() => router.push('/(tabs)/messages' as never)}>
          <Card>
            <CardHeader>
              <Text variant="title">Próxima consulta</Text>
            </CardHeader>
            {appts.isLoading ? (
              <Spinner />
            ) : nextAppt ? (
              <View>
                <Text variant="body">
                  {nextAppt.doctorName || 'Profissional'} ·{' '}
                  {nextAppt.isTelemedicine ? 'Telemedicina' : 'Presencial'}
                </Text>
                <Text variant="caption">{formatDateTime(nextAppt.scheduledAt)}</Text>
              </View>
            ) : (
              <EmptyState
                title="Sem consultas marcadas"
                description="Entre em contato com sua clínica para agendar."
              />
            )}
          </Card>
        </Pressable>
      </ScrollView>
    </SafeAreaView>
  );
}
