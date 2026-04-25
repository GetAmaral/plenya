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
  pending: 'Pendente',
  completed: 'Concluído',
};

export default function PatientLabsScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';

  const list = useQuery(options.patientLabResultsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([queryKeys.patients.labResults(patientId)]);

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
            title="Sem exames"
            description="Nenhum resultado laboratorial registrado."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <View className="flex-row items-center justify-between">
              <Text variant="title">{item.labName ?? 'Exame'}</Text>
              <Text variant="caption">{STATUS_LABEL[item.status] ?? item.status}</Text>
            </View>
            <Text variant="caption">Coletado em {formatDate(item.collectedAt)}</Text>
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
