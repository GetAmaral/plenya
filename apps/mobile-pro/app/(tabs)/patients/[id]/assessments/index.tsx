import { FlatList, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link, useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { physicalAssessmentsKeysFor } from '@plenya/api-client/options/physicalAssessments';
import { Button, Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../../features/patients/useEnsureSelectedPatient';

export default function PatientAssessmentsScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const list = useQuery(options.patientPhysicalAssessmentsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([
    physicalAssessmentsKeysFor.byPatient(patientId),
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
      <View className="px-4 pt-3">
        <Link href={`/(tabs)/patients/${patientId}/assessments/new`} asChild>
          <Button fullWidth>Nova avaliação</Button>
        </Link>
      </View>

      <FlatList
        data={list.data ?? []}
        keyExtractor={(item) => item.id}
        contentContainerClassName="gap-2 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
        ListEmptyComponent={
          <EmptyState
            title="Sem avaliações físicas"
            description="Toque em 'Nova avaliação' para registrar a primeira."
          />
        }
        renderItem={({ item }) => (
          <Link
            href={`/(tabs)/patients/${patientId}/assessments/${item.id}`}
            asChild
          >
            <Card>
              <Text variant="title">Avaliação · {formatDate(item.performedAt)}</Text>
              <View className="mt-1 flex-row gap-3">
                {item.weightKg !== undefined && (
                  <Text variant="caption">{item.weightKg} kg</Text>
                )}
                {item.heightCm !== undefined && (
                  <Text variant="caption">{item.heightCm} cm</Text>
                )}
                {item.bmi !== undefined && <Text variant="caption">IMC {item.bmi}</Text>}
                {item.bloodPressureSystolic !== undefined &&
                  item.bloodPressureDiastolic !== undefined && (
                    <Text variant="caption">
                      PA {item.bloodPressureSystolic}/{item.bloodPressureDiastolic}
                    </Text>
                  )}
              </View>
            </Card>
          </Link>
        )}
      />
    </SafeAreaView>
  );
}
