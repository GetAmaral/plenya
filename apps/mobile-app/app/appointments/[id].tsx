import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, Stack, useLocalSearchParams } from 'expo-router';
import * as WebBrowser from 'expo-web-browser';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, Card, CardHeader, EmptyState, Spinner, Text, useToast } from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

export default function AppointmentDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const qc = useQueryClient();
  const toast = useToast();
  const appt = useQuery(options.patientMeAppointmentOptions(String(id ?? '')));

  const confirm = useMutation({
    mutationFn: () => options.patientMeMutations.confirmAppointment(String(id)),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.appointments() });
      toast.show('Consulta confirmada!', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro ao confirmar', 'error'),
  });

  const cancel = useMutation({
    mutationFn: () => options.patientMeMutations.requestCancelAppointment(String(id)),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.appointments() });
      toast.show('Solicitação enviada à clínica.', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  const reschedule = useMutation({
    mutationFn: () => options.patientMeMutations.requestRescheduleAppointment(String(id)),
    onSuccess: () => toast.show('Solicitação enviada à clínica.', 'success'),
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  if (appt.isLoading) return <Spinner centered />;
  if (!appt.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <EmptyState title="Consulta não encontrada" />
      </SafeAreaView>
    );
  }

  const a = appt.data;
  const isFuture = a.minutesUntilStart > -120; // até 2h após começa, ainda mostra
  const canJoin =
    a.isTelemedicine &&
    a.dailyRoomUrl &&
    a.minutesUntilStart <= 30 &&
    a.minutesUntilStart >= -120;

  async function joinRoom() {
    if (!a.dailyRoomUrl) return;
    try {
      await WebBrowser.openBrowserAsync(a.dailyRoomUrl);
    } catch (err) {
      toast.show(err instanceof Error ? err.message : 'Não foi possível abrir', 'error');
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <Stack.Screen options={{ title: 'Consulta' }} />
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <CardHeader>
            <Text variant="title">{a.doctorName || 'Profissional'}</Text>
          </CardHeader>
          <Text variant="body">{formatDateTime(a.scheduledAt)}</Text>
          <Text variant="caption">
            {a.isTelemedicine ? 'Telemedicina' : 'Presencial'} · {a.durationMinutes}min ·{' '}
            {a.type}
          </Text>
          {a.notes && (
            <Text variant="body" className="mt-2 italic">
              {a.notes}
            </Text>
          )}
        </Card>

        {canJoin && (
          <Button onPress={joinRoom} size="lg" fullWidth>
            Entrar na sala de telemedicina
          </Button>
        )}

        {isFuture && a.status !== 'cancelled' && (
          <View className="gap-2">
            {!a.patientConfirmedAt && (
              <Button onPress={() => confirm.mutate()} loading={confirm.isPending} fullWidth>
                Confirmar presença
              </Button>
            )}
            <Button
              onPress={() => reschedule.mutate()}
              loading={reschedule.isPending}
              variant="outline"
              fullWidth
            >
              Solicitar remarcação
            </Button>
            <Button
              onPress={() => cancel.mutate()}
              loading={cancel.isPending}
              variant="ghost"
              fullWidth
            >
              Solicitar cancelamento
            </Button>
          </View>
        )}

        <Button onPress={() => router.back()} variant="ghost" fullWidth>
          Voltar
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
