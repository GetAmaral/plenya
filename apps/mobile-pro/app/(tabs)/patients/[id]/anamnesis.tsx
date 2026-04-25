import { FlatList, RefreshControl } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatRelative } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

export default function PatientAnamnesisScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const list = useQuery(options.patientAnamnesisOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([queryKeys.patients.anamnesis(patientId)]);

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
            title="Nenhuma anamnese"
            description="Este paciente ainda não tem anamnese registrada."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <Text variant="title">{item.title || 'Anamnese'}</Text>
            <Text variant="caption">Atualizada {formatRelative(item.updatedAt)}</Text>
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
