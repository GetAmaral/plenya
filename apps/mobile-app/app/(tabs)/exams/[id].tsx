import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';

export default function LabBatchDetailScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const batch = useQuery(options.patientMeLabBatchOptions(String(id ?? '')));

  if (batch.isLoading) return <Spinner centered />;
  if (!batch.data) {
    return (
      <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
        <EmptyState title="Resultado não encontrado" />
      </SafeAreaView>
    );
  }

  const b = batch.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <CardHeader>
            <Text variant="title">{b.laboratoryName}</Text>
          </CardHeader>
          <Text variant="caption">Coleta: {formatDate(b.collectionDate)}</Text>
          {b.resultDate && (
            <Text variant="caption">Resultado: {formatDate(b.resultDate)}</Text>
          )}
          {b.requestingDoctor && (
            <Text variant="caption">Solicitante: {b.requestingDoctor}</Text>
          )}
          {b.observations && (
            <Text variant="body" className="mt-2 italic">
              {b.observations}
            </Text>
          )}
        </Card>

        <View className="gap-2">
          {b.values.map((v) => (
            <Card key={v.id}>
              <View className="flex-row items-center justify-between">
                <View className="flex-1">
                  <Text variant="body" className="font-semibold">
                    {v.testName}
                  </Text>
                  <Text variant="caption">{v.testCode}</Text>
                </View>
                <Text variant="title">
                  {v.value !== undefined && v.value !== null
                    ? `${v.value}${v.unit ? ` ${v.unit}` : ''}`
                    : v.valueText || '—'}
                </Text>
              </View>
            </Card>
          ))}
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
