import { FlatList, RefreshControl, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, posturalViewTypeLabels } from '@plenya/api-client';
import { posturalKeysFor } from '@plenya/api-client/options/physicalAssessments';
import { Card, EmptyState, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

function Deviation({ label, value }: { label: string; value?: number }) {
  if (value === undefined || value === null) return null;
  return (
    <View className="rounded-full bg-muted px-2 py-0.5">
      <Text variant="caption">
        {label} {value.toFixed(1)}°
      </Text>
    </View>
  );
}

export default function PatientPostureScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const list = useQuery(options.patientPosturalAssessmentsOptions(patientId));
  const { refreshing, onRefresh } = useRefresh([posturalKeysFor.byPatient(patientId)]);

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
            title="Sem avaliações posturais"
            description="Nenhuma avaliação postural cadastrada para este paciente."
          />
        }
        renderItem={({ item }) => (
          <Card>
            <View className="flex-row items-center justify-between">
              <View className="flex-1">
                <Text variant="title">{formatDate(item.assessmentDate)}</Text>
                <Text variant="caption">
                  {posturalViewTypeLabels[item.viewType] ?? item.viewType}
                </Text>
              </View>
              <View
                className={`rounded-full px-2 py-0.5 ${
                  item.severeDeviations > 0 ? 'bg-destructive' : 'bg-emerald-600'
                }`}
              >
                <Text className="text-xs font-semibold text-white">
                  {item.posturalClassification || '—'} · {item.posturalScore}
                </Text>
              </View>
            </View>
            <View className="mt-2 flex-row flex-wrap gap-1.5">
              <Deviation label="Ombro" value={item.shoulderDeviation} />
              <Deviation label="Quadril" value={item.hipDeviation} />
              <Deviation label="Cabeça lateral" value={item.headLateralDeviation} />
              <Deviation label="FHP" value={item.fhp} />
              <Deviation label="Cifose" value={item.thoracicKyphosis} />
              <Deviation label="Lordose" value={item.lumbarLordosis} />
              <Deviation label="Joelho" value={item.kneeAngle} />
            </View>
            {item.severeDeviations > 0 && (
              <Text variant="caption" className="mt-2 text-destructive">
                ⚠ {item.severeDeviations} desvio{item.severeDeviations === 1 ? '' : 's'} severo{item.severeDeviations === 1 ? '' : 's'}
              </Text>
            )}
            {item.notes && (
              <Text variant="caption" className="mt-1 italic">
                {item.notes}
              </Text>
            )}
          </Card>
        )}
      />
    </SafeAreaView>
  );
}
