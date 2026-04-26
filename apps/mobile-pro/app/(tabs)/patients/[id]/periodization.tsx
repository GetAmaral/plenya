import { FlatList, RefreshControl } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { periodizationKeysFor } from '@plenya/api-client/options/periodizations';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

const FRAMEWORK_LABEL: Record<string, string> = {
  bompa: 'Bompa',
  linear: 'Linear',
  undulating: 'Ondulatória',
  block: 'Em blocos',
};

export default function PatientPeriodizationScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const list = useQuery(options.patientPeriodizationsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([
    periodizationKeysFor.byPatient(patientId),
  ]);

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
            title="Sem periodização"
            description="Nenhum framework de periodização cadastrado pra este paciente."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <Text variant="title">
              {FRAMEWORK_LABEL[item.framework] ?? item.framework}
            </Text>
            <Text variant="caption">
              Início {formatDate(item.startDate)} · {item.totalWeeks} semanas
              {item.status ? ` · ${item.status}` : ''}
            </Text>
            {item.goal && (
              <Text variant="body" className="mt-1">
                {item.goal}
              </Text>
            )}
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
