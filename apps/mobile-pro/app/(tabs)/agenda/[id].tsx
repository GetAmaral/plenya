import { Linking, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  options,
  appointmentTypeLabels,
  queryKeys,
  type AppointmentStatus,
  type AppointmentType,
} from '@plenya/api-client';
import {
  Button,
  Card,
  CardHeader,
  ErrorState,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { formatDateTime } from '@plenya/domain';

const STATUS_LABEL: Record<AppointmentStatus, string> = {
  scheduled: 'Agendada',
  confirmed: 'Confirmada',
  completed: 'Concluída',
  cancelled: 'Cancelada',
  no_show: 'Faltou',
};

export default function AppointmentDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const appointmentId = id ?? '';
  const queryClient = useQueryClient();
  const toast = useToast();

  const query = useQuery(options.appointmentDetailOptions(appointmentId));

  const confirm = useMutation({
    mutationFn: () => options.appointmentMutations.confirm(appointmentId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.appointments.all() });
      toast.show('Consulta confirmada', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao confirmar', 'error'),
  });

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }
  if (!query.data) return null;

  const a = query.data;
  const isTelemedicine: AppointmentType = a.type;
  const canJoinCall =
    a.type === 'telemedicine' &&
    a.dailyRoomUrl &&
    (a.status === 'scheduled' || a.status === 'confirmed');
  const canConfirm = a.status === 'scheduled';

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <View>
          <Text variant="heading">
            {a.patientName ?? 'Paciente'}
          </Text>
          <Text variant="caption">
            {appointmentTypeLabels[isTelemedicine]} · {a.durationMinutes}min · {STATUS_LABEL[a.status]}
          </Text>
        </View>

        <Card>
          <CardHeader>
            <Text variant="title">Quando</Text>
          </CardHeader>
          <Text variant="body">{formatDateTime(a.scheduledAt)}</Text>
          {a.confirmedAt && (
            <Text variant="caption">Confirmada em {formatDateTime(a.confirmedAt)}</Text>
          )}
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Motivo</Text>
          </CardHeader>
          <Text variant="body">{a.reason}</Text>
          {a.patientNotes && (
            <View className="mt-2">
              <Text variant="caption" className="font-semibold">
                Notas do paciente
              </Text>
              <Text variant="body">{a.patientNotes}</Text>
            </View>
          )}
        </Card>

        {canJoinCall && (
          <Button
            onPress={() => {
              if (!a.dailyRoomUrl) return;
              Linking.openURL(a.dailyRoomUrl).catch(() =>
                toast.show('Falha ao abrir a sala', 'error'),
              );
            }}
            fullWidth
            size="lg"
          >
            Entrar na teleconsulta
          </Button>
        )}

        {canConfirm && (
          <Button
            onPress={() => confirm.mutate()}
            loading={confirm.isPending}
            variant="outline"
            fullWidth
          >
            Confirmar consulta
          </Button>
        )}

        {a.patientId && (
          <Button
            variant="ghost"
            onPress={() => router.push(`/(tabs)/patients/${a.patientId}` as never)}
            fullWidth
          >
            Abrir prontuário do paciente
          </Button>
        )}

        {(a.doctorNotes || a.diagnosis) && (
          <Card>
            <CardHeader>
              <Text variant="title">Pós-consulta</Text>
            </CardHeader>
            {a.doctorNotes && <Text variant="body">{a.doctorNotes}</Text>}
            {a.diagnosis && (
              <View className="mt-2">
                <Text variant="caption" className="font-semibold">
                  Diagnóstico
                </Text>
                <Text variant="body">{a.diagnosis}</Text>
              </View>
            )}
          </Card>
        )}

        {a.cancelledAt && (
          <Card className="border-destructive">
            <Text variant="title" className="text-destructive">
              Cancelada
            </Text>
            <Text variant="caption">{formatDateTime(a.cancelledAt)}</Text>
            {a.cancellationReason && (
              <Text variant="body" className="mt-1">
                {a.cancellationReason}
              </Text>
            )}
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
