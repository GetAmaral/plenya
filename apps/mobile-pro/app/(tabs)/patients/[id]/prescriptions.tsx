import { FlatList, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';

const STATUS_LABEL: Record<string, string> = {
  draft: 'Rascunho',
  signed: 'Assinada',
  cancelled: 'Cancelada',
};

const STATUS_COLOR: Record<string, string> = {
  draft: 'text-muted-foreground',
  signed: 'text-emerald-600',
  cancelled: 'text-destructive',
};

export default function PatientPrescriptionsScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';

  const list = useQuery(options.patientPrescriptionsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([queryKeys.patients.prescriptions(patientId)]);

  if (list.isLoading) return <Spinner centered />;
  if (list.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => list.refetch()} />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <FlatList
        data={list.data ?? []}
        keyExtractor={(item) => item.id}
        contentContainerClassName="gap-2 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
        ListEmptyComponent={
          <EmptyState
            title="Sem prescrições"
            description="Nenhuma receita emitida para este paciente."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <View className="flex-row items-center justify-between">
              <Text variant="title">Receita</Text>
              <Text className={`text-xs font-semibold ${STATUS_COLOR[item.status] ?? ''}`}>
                {STATUS_LABEL[item.status] ?? item.status}
              </Text>
            </View>
            <Text variant="caption">Emitida em {formatDate(item.issuedAt)}</Text>
            {item.signedAt && (
              <Text variant="caption">Assinada em {formatDate(item.signedAt)}</Text>
            )}
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
