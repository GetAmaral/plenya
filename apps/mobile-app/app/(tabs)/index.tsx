import { useMemo } from 'react';
import { Pressable, RefreshControl, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate, formatDateTime } from '@plenya/domain';

export default function HomeScreen() {
  const profile = useQuery(options.patientMeProfileOptions());
  const appts = useQuery(options.patientMeAppointmentsOptions('upcoming'));
  const checkIn = useQuery(options.patientMeCheckInTodayOptions());
  const scores = useQuery(options.patientMeScoresOptions());
  const batches = useQuery(options.patientMeLabBatchesOptions());
  const continuum = useQuery(options.patientMeContinuumOptions());
  const qc = useQueryClient();

  const nextAppt = useMemo(() => (appts.data ?? [])[0], [appts.data]);
  const lastScore = useMemo(() => (scores.data ?? [])[0], [scores.data]);
  const lastBatch = useMemo(() => (batches.data ?? [])[0], [batches.data]);

  const greeting = useMemo(() => {
    const h = new Date().getHours();
    if (h < 12) return 'Bom dia';
    if (h < 18) return 'Boa tarde';
    return 'Boa noite';
  }, []);

  const refreshing =
    (appts.isRefetching && !appts.isLoading) ||
    (checkIn.isRefetching && !checkIn.isLoading) ||
    (scores.isRefetching && !scores.isLoading);

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

        {nextAppt ? (
          <Pressable
            onPress={() => router.push(`/appointments/${nextAppt.id}` as never)}
          >
            <Card>
              <CardHeader>
                <Text variant="title">Próxima consulta</Text>
              </CardHeader>
              <View>
                <Text variant="body">
                  {nextAppt.doctorName || 'Profissional'} ·{' '}
                  {nextAppt.isTelemedicine ? 'Telemedicina' : 'Presencial'}
                </Text>
                <Text variant="caption">{formatDateTime(nextAppt.scheduledAt)}</Text>
                {!nextAppt.patientConfirmedAt && (
                  <Text variant="caption" className="mt-1 font-semibold">
                    Toque pra confirmar presença →
                  </Text>
                )}
              </View>
            </Card>
          </Pressable>
        ) : (
          <Card>
            <CardHeader>
              <Text variant="title">Próxima consulta</Text>
            </CardHeader>
            {appts.isLoading ? (
              <Spinner />
            ) : (
              <EmptyState
                title="Sem consultas marcadas"
                description="Entre em contato com sua clínica para agendar."
              />
            )}
          </Card>
        )}

        {lastScore && (
          <Pressable onPress={() => router.push('/(tabs)/exams' as never)}>
            <Card>
              <CardHeader>
                <Text variant="title">Seu último escore</Text>
              </CardHeader>
              <Text variant="heading">
                {Math.round(lastScore.totalScorePercentage)}%
              </Text>
              <Text variant="caption">
                {formatDate(lastScore.createdAt)} · {lastScore.itemsEvaluatedCount} itens
              </Text>
            </Card>
          </Pressable>
        )}

        {lastBatch && (
          <Pressable onPress={() => router.push(`/exams/${lastBatch.id}` as never)}>
            <Card>
              <CardHeader>
                <Text variant="title">Último exame</Text>
              </CardHeader>
              <Text variant="body">{lastBatch.laboratoryName}</Text>
              <Text variant="caption">
                Coleta: {formatDate(lastBatch.collectionDate)} · {lastBatch.resultsCount}{' '}
                resultados
              </Text>
            </Card>
          </Pressable>
        )}

        {continuum.data && continuum.data.items?.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Sua jornada</Text>
            </CardHeader>
            <Text variant="caption" className="mb-2">
              {continuum.data.templateName} · {continuum.data.status}
            </Text>
            {continuum.data.items.slice(0, 3).map((it) => (
              <View key={it.id} className="border-l-2 border-primary pl-3 py-1">
                <Text variant="body">{it.title}</Text>
                <Text variant="caption">
                  {it.status}
                  {it.expectedDate ? ` · ${formatDate(it.expectedDate)}` : ''}
                </Text>
              </View>
            ))}
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
