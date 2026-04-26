import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text } from '@plenya/ui-mobile';

export default function TrainingScreen() {
  const plans = useQuery(options.patientMeWorkoutPlansOptions());

  if (plans.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Pressable
          onPress={() => router.push('/(tabs)/training/history' as never)}
          accessibilityRole="button"
          accessibilityLabel="Ver histórico de treinos"
        >
          <Card>
            <Text variant="title">Histórico</Text>
            <Text variant="caption">Suas últimas sessões e frequência semanal.</Text>
          </Card>
        </Pressable>

        <Text variant="caption" className="font-semibold uppercase mt-2">
          Planos ativos
        </Text>

        {(plans.data ?? []).length === 0 ? (
          <EmptyState
            title="Sem planos ativos"
            description="Quando seu profissional liberar um plano, ele aparece aqui."
          />
        ) : (
          (plans.data ?? []).map((p) => (
            <Pressable
              key={p.id}
              onPress={() => router.push(`/(tabs)/training/${p.id}` as never)}
              accessibilityRole="button"
              accessibilityLabel={`Abrir plano ${p.name}`}
            >
              <Card>
                <CardHeader>
                  <Text variant="title">{p.name}</Text>
                </CardHeader>
                <Text variant="caption">
                  {p.objective} · {p.intensity} · {p.weeklyFrequency}×/sem ·{' '}
                  {p.totalSessions} sessões
                </Text>
              </Card>
            </Pressable>
          ))
        )}

        <View className="h-4" />
      </ScrollView>
    </SafeAreaView>
  );
}
