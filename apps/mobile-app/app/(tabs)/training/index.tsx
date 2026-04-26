import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, EmptyState, Spinner, Text } from '@plenya/ui-mobile';

export default function TrainingScreen() {
  const plans = useQuery(options.patientMeWorkoutPlansOptions());

  if (plans.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        {(plans.data ?? []).length === 0 ? (
          <EmptyState
            title="Sem planos ativos"
            description="Quando seu profissional liberar um plano, ele aparece aqui."
          />
        ) : (
          (plans.data ?? []).map((p) => (
            <Card key={p.id}>
              <Text variant="title">{p.name}</Text>
              <Text variant="caption">
                {p.objective} · {p.intensity} · {p.weeklyFrequency}×/sem ·{' '}
                {p.totalSessions} sessões
              </Text>
            </Card>
          ))
        )}
        <View className="h-4" />
        <Text variant="caption" className="text-center italic">
          Logger de execução chega na Sprint 3.
        </Text>
      </ScrollView>
    </SafeAreaView>
  );
}
