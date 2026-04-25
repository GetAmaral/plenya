import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';

export default function AssessmentDetailScreen() {
  useScreenCaptureProtection();
  const { assessmentId } = useLocalSearchParams<{ assessmentId: string }>();
  const query = useQuery(options.physicalAssessmentDetailOptions(assessmentId ?? ''));

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }

  const a = query.data;
  if (!a) return null;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Text variant="heading">Avaliação física</Text>
        <Text variant="caption">{formatDate(a.performedAt)}</Text>

        <Card>
          <CardHeader>
            <Text variant="title">Composição corporal</Text>
          </CardHeader>
          <View className="gap-1">
            {a.weightKg !== undefined && <Text variant="body">Peso: {a.weightKg} kg</Text>}
            {a.heightCm !== undefined && <Text variant="body">Altura: {a.heightCm} cm</Text>}
            {a.bmi !== undefined && <Text variant="body">IMC: {a.bmi}</Text>}
          </View>
        </Card>

        {a.bloodPressureSystolic !== undefined && a.bloodPressureDiastolic !== undefined && (
          <Card>
            <CardHeader>
              <Text variant="title">Pressão arterial</Text>
            </CardHeader>
            <Text variant="body">
              {a.bloodPressureSystolic}/{a.bloodPressureDiastolic} mmHg
            </Text>
          </Card>
        )}

        {a.acsmTagsStructured && a.acsmTagsStructured.length > 0 && (
          <Card>
            <CardHeader>
              <Text variant="title">Tags ACSM</Text>
            </CardHeader>
            <View className="flex-row flex-wrap gap-2">
              {a.acsmTagsStructured.map((tag, idx) => (
                <View
                  key={`${tag.label}-${idx}`}
                  className="rounded-full px-3 py-1"
                  style={{ backgroundColor: tag.color || '#94a3b8' }}
                >
                  <Text className="text-xs font-semibold text-white">{tag.label}</Text>
                </View>
              ))}
            </View>
          </Card>
        )}

        {a.notes && (
          <Card>
            <CardHeader>
              <Text variant="title">Observações</Text>
            </CardHeader>
            <Text variant="body">{a.notes}</Text>
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
