import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text, useToast } from '@plenya/ui-mobile';

/**
 * Detalhe do plano — mostra cada PlanSession (Treino A/B/C) com botão "Iniciar"
 * que cria uma WorkoutSession e abre o logger.
 */
export default function PlanDetailScreen() {
  const { planId } = useLocalSearchParams<{ planId: string }>();
  const qc = useQueryClient();
  const toast = useToast();
  const plan = useQuery(options.patientMeWorkoutPlanOptions(String(planId ?? '')));

  const start = useMutation({
    mutationFn: (planSessionId: string) =>
      options.patientMeMutations.startWorkoutSession({
        planId: String(planId),
        planSessionId,
        scheduledDate: new Date().toISOString(),
      }),
    onSuccess: (sess) => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.workoutSessions() });
      router.push(`/(tabs)/training/sessions/${sess.id}` as never);
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  if (plan.isLoading) return <Spinner centered />;
  if (!plan.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <EmptyState title="Plano não encontrado" />
      </SafeAreaView>
    );
  }

  const p = plan.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Card>
          <CardHeader>
            <Text variant="title">{p.name}</Text>
          </CardHeader>
          <Text variant="caption">
            {p.objective} · {p.intensity} · {p.weeklyFrequency}×/sem
          </Text>
        </Card>

        <Text variant="caption" className="font-semibold uppercase mt-2">
          Sessões
        </Text>

        {(p.sessions ?? []).map((s) => (
          <Card key={s.id}>
            <View className="flex-row items-center justify-between gap-3">
              <View className="flex-1">
                <Text variant="body" className="font-semibold">
                  {s.name}
                </Text>
                <Text variant="caption">{s.exercises?.length ?? 0} exercícios</Text>
              </View>
              <Pressable
                onPress={() => start.mutate(s.id)}
                disabled={start.isPending}
                className="rounded-lg bg-primary px-4 py-2"
                accessibilityRole="button"
                accessibilityLabel={`Iniciar ${s.name}`}
              >
                <Text className="font-semibold text-primary-foreground">
                  {start.isPending ? 'Abrindo…' : 'Iniciar'}
                </Text>
              </Pressable>
            </View>

            {(s.exercises ?? []).slice(0, 5).map((ex) => (
              <View key={ex.id} className="mt-2 border-l-2 border-muted pl-3">
                <Text variant="caption">
                  {ex.exercise?.name ?? 'Exercício'} · {ex.sets}×{ex.reps}
                </Text>
              </View>
            ))}
            {(s.exercises ?? []).length > 5 && (
              <Text variant="caption" className="mt-1 italic">
                + {(s.exercises ?? []).length - 5} exercícios
              </Text>
            )}
          </Card>
        ))}
      </ScrollView>
    </SafeAreaView>
  );
}
