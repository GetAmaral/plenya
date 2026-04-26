import { useMemo } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options, type PatientWorkoutSession } from '@plenya/api-client';
import { Card, EmptyState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';

/**
 * Histórico — lista cronológica + bar chart de "sessões por semana"
 * (últimas 8 semanas) feito com Views (sem libs de SVG/chart).
 */

function weekKey(d: Date): string {
  const tmp = new Date(d);
  const day = tmp.getDay() || 7; // segunda-feira como início (ISO)
  tmp.setDate(tmp.getDate() - (day - 1));
  tmp.setHours(0, 0, 0, 0);
  return tmp.toISOString().slice(0, 10);
}

function FrequencyChart({ sessions }: { sessions: PatientWorkoutSession[] }) {
  const buckets = useMemo(() => {
    const counts = new Map<string, number>();
    const today = new Date();
    today.setHours(0, 0, 0, 0);
    const weeks: string[] = [];
    for (let i = 7; i >= 0; i--) {
      const d = new Date(today);
      d.setDate(d.getDate() - i * 7);
      const k = weekKey(d);
      weeks.push(k);
      counts.set(k, 0);
    }
    sessions.forEach((s) => {
      if (!s.completedAt) return;
      const k = weekKey(new Date(s.completedAt));
      if (counts.has(k)) counts.set(k, (counts.get(k) ?? 0) + 1);
    });
    return weeks.map((k) => ({ key: k, count: counts.get(k) ?? 0 }));
  }, [sessions]);

  const max = Math.max(1, ...buckets.map((b) => b.count));

  return (
    <View>
      <View className="h-24 flex-row items-end gap-1.5">
        {buckets.map((b) => (
          <View key={b.key} className="flex-1 items-center justify-end">
            <View
              className="w-full rounded-sm bg-primary"
              style={{ height: `${(b.count / max) * 100}%`, minHeight: 2 }}
            />
            {b.count > 0 && (
              <Text variant="caption" className="mt-1 text-[10px]">
                {b.count}
              </Text>
            )}
          </View>
        ))}
      </View>
      <View className="mt-1 flex-row justify-between">
        <Text variant="caption" className="text-[10px]">
          8 sem atrás
        </Text>
        <Text variant="caption" className="text-[10px]">
          agora
        </Text>
      </View>
    </View>
  );
}

export default function TrainingHistoryScreen() {
  const sessions = useQuery(options.patientMeWorkoutSessionsOptions(60));

  if (sessions.isLoading) return <Spinner centered />;

  const data = sessions.data ?? [];
  const completedCount = data.filter((s) => s.completedAt).length;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Card>
          <Text variant="title">Frequência</Text>
          <Text variant="caption">{completedCount} treinos concluídos no histórico</Text>
          <View className="mt-2">
            <FrequencyChart sessions={data} />
          </View>
        </Card>

        {data.length === 0 ? (
          <EmptyState
            title="Sem sessões registradas"
            description="Inicie um treino pra começar seu histórico."
          />
        ) : (
          data.map((s) => (
            <Pressable
              key={s.id}
              onPress={() => router.push(`/(tabs)/training/sessions/${s.id}` as never)}
              accessibilityRole="button"
              accessibilityLabel={`Abrir sessão de ${formatDate(s.scheduledDate)}`}
            >
              <Card>
                <View className="flex-row items-center justify-between">
                  <View className="flex-1">
                    <Text variant="body" className="font-semibold">
                      {s.planSession?.name ?? 'Sessão'}
                    </Text>
                    <Text variant="caption">
                      {formatDate(s.scheduledDate)} · {s.logs?.length ?? 0} sets
                    </Text>
                  </View>
                  <Text
                    variant="caption"
                    className={s.completedAt ? 'font-semibold text-primary' : ''}
                  >
                    {s.completedAt ? '✓ concluído' : 'em andamento'}
                  </Text>
                </View>
              </Card>
            </Pressable>
          ))
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
