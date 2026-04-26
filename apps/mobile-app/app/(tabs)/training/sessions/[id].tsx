import { useMemo, useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  options,
  queryKeys,
  type PatientWorkoutLog,
  type PatientWorkoutSessionTemplate,
} from '@plenya/api-client';
import {
  Button,
  Card,
  EmptyState,
  Input,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';

/**
 * Logger inline — para cada exercício, mostra cada set como linha editável
 * (reps + peso + RPE). Tap "Salvar set" upserta no backend.
 *
 * Estado local mantém os valores em digitação; backend é a fonte da verdade
 * pra "qual set já foi salvo" (via session.logs).
 */

type LogDraft = {
  reps?: string;
  weight?: string;
  rpe?: string;
  notes?: string;
};

function findLog(
  logs: PatientWorkoutLog[] | undefined,
  planExerciseId: string,
  setNumber: number,
): PatientWorkoutLog | undefined {
  return (logs ?? []).find(
    (l) => l.planExerciseId === planExerciseId && l.setNumber === setNumber,
  );
}

function ExerciseBlock({
  ex,
  sessionId,
  logs,
}: {
  ex: PatientWorkoutSessionTemplate['exercises'][number];
  sessionId: string;
  logs?: PatientWorkoutLog[];
}) {
  const qc = useQueryClient();
  const toast = useToast();
  const [drafts, setDrafts] = useState<Record<number, LogDraft>>({});

  const log = useMutation({
    mutationFn: ({ setNumber, draft }: { setNumber: number; draft: LogDraft }) =>
      options.patientMeMutations.logSet(sessionId, {
        planExerciseId: ex.id,
        exerciseId: ex.exerciseId,
        setNumber,
        reps: draft.reps ? Number(draft.reps) : undefined,
        weight: draft.weight ? Number(draft.weight.replace(',', '.')) : undefined,
        rpe: draft.rpe ? Number(draft.rpe) : undefined,
        notes: draft.notes,
      }),
    onSuccess: () => {
      qc.invalidateQueries({
        queryKey: queryKeys.patientMe.workoutSessionDetail(sessionId),
      });
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  const sets = Array.from({ length: ex.sets }, (_, i) => i + 1);

  return (
    <Card>
      <Text variant="body" className="font-semibold">
        {ex.exercise?.name ?? 'Exercício'}
      </Text>
      <Text variant="caption">
        Alvo: {ex.sets}×{ex.reps} · descanso {ex.restBetweenSetsSec}s
      </Text>
      {ex.notes && (
        <Text variant="caption" className="italic mt-1">
          {ex.notes}
        </Text>
      )}

      <View className="gap-2 mt-3">
        {sets.map((n) => {
          const existing = findLog(logs, ex.id, n);
          const draft =
            drafts[n] ??
            (existing
              ? {
                  reps: existing.reps?.toString() ?? '',
                  weight: existing.weight?.toString() ?? '',
                  rpe: existing.rpe?.toString() ?? '',
                }
              : {});
          const done = Boolean(existing);
          return (
            <View
              key={n}
              className={`rounded-lg border p-3 ${
                done ? 'border-primary bg-primary/10' : 'border-border bg-card'
              }`}
            >
              <View className="flex-row items-center justify-between mb-2">
                <Text variant="caption" className="font-semibold">
                  Set {n}
                </Text>
                {done && (
                  <Text variant="caption" className="font-semibold text-primary">
                    ✓ salvo
                  </Text>
                )}
              </View>
              <View className="flex-row gap-2">
                <View className="flex-1">
                  <Text variant="caption">Reps</Text>
                  <Input
                    value={draft.reps ?? ''}
                    onChangeText={(v) =>
                      setDrafts((prev) => ({ ...prev, [n]: { ...draft, reps: v } }))
                    }
                    keyboardType="number-pad"
                    placeholder="0"
                  />
                </View>
                <View className="flex-1">
                  <Text variant="caption">Peso (kg)</Text>
                  <Input
                    value={draft.weight ?? ''}
                    onChangeText={(v) =>
                      setDrafts((prev) => ({ ...prev, [n]: { ...draft, weight: v } }))
                    }
                    keyboardType="decimal-pad"
                    placeholder="0"
                  />
                </View>
                <View className="w-16">
                  <Text variant="caption">RPE</Text>
                  <Input
                    value={draft.rpe ?? ''}
                    onChangeText={(v) =>
                      setDrafts((prev) => ({ ...prev, [n]: { ...draft, rpe: v } }))
                    }
                    keyboardType="number-pad"
                    placeholder="—"
                  />
                </View>
              </View>
              <Pressable
                onPress={() => log.mutate({ setNumber: n, draft })}
                disabled={log.isPending}
                className="mt-2 rounded-lg bg-primary py-2"
                accessibilityRole="button"
                accessibilityLabel={`Salvar set ${n}`}
              >
                <Text className="text-center font-semibold text-primary-foreground">
                  {done ? 'Atualizar set' : 'Salvar set'}
                </Text>
              </Pressable>
            </View>
          );
        })}
      </View>
    </Card>
  );
}

export default function WorkoutSessionScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const qc = useQueryClient();
  const toast = useToast();
  const sess = useQuery(options.patientMeWorkoutSessionOptions(String(id ?? '')));
  const [finalNotes, setFinalNotes] = useState('');

  const complete = useMutation({
    mutationFn: () =>
      options.patientMeMutations.completeWorkoutSession(String(id), {
        notes: finalNotes || undefined,
      }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.workoutSessions() });
      toast.show('Treino concluído! 💪', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  const exercises = useMemo(
    () => sess.data?.planSession?.exercises ?? [],
    [sess.data?.planSession?.exercises],
  );

  if (sess.isLoading) return <Spinner centered />;
  if (!sess.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <EmptyState title="Sessão não encontrada" />
      </SafeAreaView>
    );
  }

  const completed = Boolean(sess.data.completedAt);

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Card>
          <Text variant="title">{sess.data.planSession?.name ?? 'Sessão'}</Text>
          <Text variant="caption">
            {completed ? 'Concluído ✓' : 'Em andamento'} · {sess.data.logs?.length ?? 0}{' '}
            sets registrados
          </Text>
        </Card>

        {exercises.map((ex) => (
          <ExerciseBlock
            key={ex.id}
            ex={ex}
            sessionId={String(id)}
            logs={sess.data?.logs}
          />
        ))}

        {!completed && (
          <Card>
            <Text variant="caption">Notas finais (opcional)</Text>
            <Input
              value={finalNotes}
              onChangeText={setFinalNotes}
              multiline
              placeholder="Como foi o treino?"
            />
            <Button
              onPress={() => complete.mutate()}
              loading={complete.isPending}
              size="lg"
              fullWidth
              className="mt-3"
            >
              Finalizar treino
            </Button>
          </Card>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
