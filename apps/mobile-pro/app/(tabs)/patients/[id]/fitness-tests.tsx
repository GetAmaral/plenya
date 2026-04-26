import { FlatList, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { fitnessKeysFor } from '@plenya/api-client/options/physicalAssessments';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

function MetricChip({ label, value }: { label: string; value: number | undefined }) {
  if (value === undefined || value === null) return null;
  return (
    <View className="rounded-full bg-muted px-2 py-0.5">
      <Text variant="caption" className="text-foreground">
        {label} {value}
      </Text>
    </View>
  );
}

export default function PatientFitnessTestsScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const list = useQuery(options.patientFitnessTestsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([fitnessKeysFor.byPatient(patientId)]);

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
            title="Sem testes físicos"
            description="Nenhum teste físico (abdominal, flexão, prancha, burpee, FRT) registrado."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <View className="flex-row items-center justify-between">
              <Text variant="title">{formatDate(item.assessmentDate)}</Text>
              {item.overallClassification && (
                <View className="rounded-full bg-primary px-2 py-0.5">
                  <Text className="text-xs font-semibold text-primary-foreground">
                    {item.overallClassification} · {item.overallScore}
                  </Text>
                </View>
              )}
            </View>
            <View className="mt-2 flex-row flex-wrap gap-1.5">
              <MetricChip label="Abdominal" value={item.abdominalReps} />
              <MetricChip label="Flexão" value={item.pushupReps} />
              <MetricChip label="Prancha" value={item.plankSeconds} />
              <MetricChip label="Burpee" value={item.burpeeCycles} />
              <MetricChip label="FRT" value={item.frtReps} />
            </View>
            {item.notes && (
              <Text variant="caption" className="mt-2 italic">
                {item.notes}
              </Text>
            )}
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
