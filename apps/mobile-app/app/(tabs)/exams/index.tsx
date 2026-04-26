import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, EmptyState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';

/**
 * Tela Exames — duas seções: Lab batches + Escores.
 * Pull-to-refresh implícito via TanStack staleTime; botão de refresh entra
 * só se o usuário pedir.
 */

function StatusBadge({ status }: { status: string }) {
  const map: Record<string, { label: string; cls: string }> = {
    pending: { label: 'Pendente', cls: 'bg-muted' },
    completed: { label: 'Concluído', cls: 'bg-primary' },
    in_progress: { label: 'Em análise', cls: 'bg-amber-500' },
  };
  const cfg = map[status] ?? { label: status, cls: 'bg-muted' };
  return (
    <View className={`rounded-full px-2 py-0.5 ${cfg.cls}`}>
      <Text className="text-xs text-primary-foreground">{cfg.label}</Text>
    </View>
  );
}

export default function ExamsScreen() {
  const batches = useQuery(options.patientMeLabBatchesOptions());
  const scores = useQuery(options.patientMeScoresOptions());

  if (batches.isLoading && scores.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-6 p-4">
        <View className="gap-2">
          <Text variant="caption" className="font-semibold uppercase">
            Exames laboratoriais
          </Text>
          {(batches.data ?? []).length === 0 ? (
            <EmptyState
              title="Sem exames cadastrados"
              description="Quando sua clínica registrar resultados, eles aparecem aqui."
            />
          ) : (
            (batches.data ?? []).map((b) => (
              <Pressable
                key={b.id}
                onPress={() => router.push(`/exams/${b.id}` as never)}
                accessibilityRole="button"
                accessibilityLabel={`Abrir exame ${b.laboratoryName}`}
              >
                <Card>
                  <View className="flex-row items-start justify-between gap-2">
                    <View className="flex-1">
                      <Text variant="body" className="font-semibold">
                        {b.laboratoryName}
                      </Text>
                      <Text variant="caption">
                        Coleta: {formatDate(b.collectionDate)}
                      </Text>
                      <Text variant="caption">{b.resultsCount} resultados</Text>
                    </View>
                    <StatusBadge status={b.status} />
                  </View>
                </Card>
              </Pressable>
            ))
          )}
        </View>

        <View className="gap-2">
          <Text variant="caption" className="font-semibold uppercase">
            Escores
          </Text>
          {(scores.data ?? []).length === 0 ? (
            <EmptyState
              title="Sem escores ainda"
              description="Sua avaliação Plenya aparece aqui quando estiver pronta."
            />
          ) : (
            (scores.data ?? []).map((s) => (
              <Card key={s.id}>
                <CardHeader>
                  <Text variant="title">
                    {Math.round(s.totalScorePercentage)}%
                  </Text>
                </CardHeader>
                <Text variant="caption">
                  {s.source === 'light' ? 'Avaliação Light' : 'Avaliação completa'} ·{' '}
                  {formatDate(s.createdAt)} · {s.itemsEvaluatedCount} itens
                </Text>
                <View className="mt-2 gap-1">
                  {s.groupResults.map((g) => (
                    <View
                      key={g.groupId}
                      className="flex-row items-center justify-between"
                    >
                      <Text variant="caption" className="flex-1">
                        {g.groupName}
                      </Text>
                      <Text variant="caption" className="font-semibold">
                        {Math.round(g.scorePercentage)}%
                      </Text>
                    </View>
                  ))}
                </View>
              </Card>
            ))
          )}
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
