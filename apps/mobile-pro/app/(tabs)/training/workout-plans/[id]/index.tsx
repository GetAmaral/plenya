import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Link, router, useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';

const STATUS_LABEL: Record<string, string> = {
  draft: 'Rascunho',
  active: 'Ativo',
  completed: 'Concluído',
  archived: 'Arquivado',
};

export default function WorkoutPlanDetailScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const planId = id ?? '';
  const query = useQuery(options.workoutPlanOptions(planId));

  if (query.isLoading) return <Spinner centered />;
  if (query.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => query.refetch()} />
      </SafeAreaView>
    );
  }

  const plan = query.data;
  if (!plan) return null;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <View className="flex-row items-start justify-between gap-3">
          <View className="flex-1">
            <Text variant="heading">{plan.name}</Text>
            <Text variant="caption">
              {STATUS_LABEL[plan.status] ?? plan.status} · iniciado em {formatDate(plan.startDate)}
              {plan.endDate ? ` · término ${formatDate(plan.endDate)}` : ''}
            </Text>
          </View>
          <Button
            variant="outline"
            size="sm"
            onPress={() => router.push(`/(tabs)/training/workout-plans/${planId}/edit` as never)}
          >
            Editar
          </Button>
        </View>

        {plan.notes && (
          <Card>
            <Text variant="body">{plan.notes}</Text>
          </Card>
        )}

        {plan.sessions
          .slice()
          .sort((a, b) => a.order - b.order)
          .map((session) => (
            <Card key={session.id}>
              <CardHeader>
                <Text variant="title">{session.name || `Sessão ${session.order}`}</Text>
              </CardHeader>
              <View className="gap-2">
                {session.exercises.map((ex, idx) => (
                  <Link
                    key={`${session.id}-${idx}`}
                    href={`/(tabs)/training/exercises/${ex.exerciseId}`}
                    asChild
                  >
                    <View className="rounded-md bg-muted px-3 py-2">
                      <Text variant="body">{ex.exerciseName}</Text>
                      <Text variant="caption">
                        {ex.sets} séries × {ex.reps} reps · descanso {ex.rest}
                      </Text>
                      {ex.notes && (
                        <Text variant="caption" className="mt-0.5 italic">
                          {ex.notes}
                        </Text>
                      )}
                    </View>
                  </Link>
                ))}
              </View>
            </Card>
          ))}
      </ScrollView>
    </SafeAreaView>
  );
}
