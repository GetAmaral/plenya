import { useEffect, useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  options,
  queryKeys,
  type ExerciseSummary,
  type WorkoutPlanSession,
} from '@plenya/api-client';
import {
  Button,
  Card,
  ErrorState,
  FormField,
  Input,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';

interface DraftExercise {
  exerciseId: string;
  exerciseName: string;
  sets: number;
  reps: string;
  rest: string;
  notes?: string;
}

const STATUSES: Array<{ key: 'draft' | 'active' | 'completed' | 'archived'; label: string }> = [
  { key: 'draft', label: 'Rascunho' },
  { key: 'active', label: 'Ativo' },
  { key: 'completed', label: 'Concluído' },
  { key: 'archived', label: 'Arquivado' },
];

export default function EditWorkoutPlanScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const planId = id ?? '';
  const queryClient = useQueryClient();
  const toast = useToast();

  const plan = useQuery(options.workoutPlanOptions(planId));

  const [name, setName] = useState('');
  const [notes, setNotes] = useState('');
  const [status, setStatus] =
    useState<'draft' | 'active' | 'completed' | 'archived'>('active');
  const [sessionName, setSessionName] = useState('Sessão A');
  const [exercises, setExercises] = useState<DraftExercise[]>([]);
  const [exercisePickerOpen, setExercisePickerOpen] = useState(false);
  const [search, setSearch] = useState('');

  useEffect(() => {
    if (!plan.data) return;
    setName(plan.data.name);
    setNotes(plan.data.notes ?? '');
    setStatus(plan.data.status);
    const first: WorkoutPlanSession | undefined = plan.data.sessions.slice().sort(
      (a, b) => a.order - b.order,
    )[0];
    if (first) {
      setSessionName(first.name || 'Sessão A');
      setExercises(
        first.exercises.map((ex) => ({
          exerciseId: ex.exerciseId,
          exerciseName: ex.exerciseName,
          sets: ex.sets,
          reps: ex.reps,
          rest: ex.rest,
          notes: ex.notes,
        })),
      );
    }
  }, [plan.data]);

  const exerciseList = useQuery(options.exercisesListOptions({ search }));

  const update = useMutation({
    mutationFn: () =>
      options.workoutPlanMutations.update(planId, {
        name: name.trim(),
        status,
        notes: notes.trim() || undefined,
        sessions:
          exercises.length > 0
            ? [
                {
                  name: sessionName.trim() || 'Sessão A',
                  order: 1,
                  exercises: exercises.map((e) => ({
                    exerciseId: e.exerciseId,
                    sets: e.sets,
                    reps: e.reps,
                    rest: e.rest,
                    notes: e.notes,
                  })),
                },
              ]
            : undefined,
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.workoutPlans.detail(planId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.workoutPlans.all() });
      toast.show('Plano atualizado', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error'),
  });

  function addExercise(ex: ExerciseSummary) {
    setExercises((prev) => [
      ...prev,
      { exerciseId: ex.id, exerciseName: ex.name, sets: 3, reps: '10', rest: '60s' },
    ]);
    setExercisePickerOpen(false);
    setSearch('');
  }

  function updateExercise(index: number, patch: Partial<DraftExercise>) {
    setExercises((prev) => prev.map((e, i) => (i === index ? { ...e, ...patch } : e)));
  }

  function removeExercise(index: number) {
    setExercises((prev) => prev.filter((_, i) => i !== index));
  }

  if (plan.isLoading) return <Spinner centered />;
  if (plan.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => plan.refetch()} />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="heading">Editar plano</Text>

        <FormField label="Nome" required>
          <Input value={name} onChangeText={setName} editable={!update.isPending} />
        </FormField>

        <FormField label="Status">
          <View className="flex-row flex-wrap gap-2">
            {STATUSES.map((s) => (
              <Pressable
                key={s.key}
                onPress={() => setStatus(s.key)}
                className={`rounded-full px-3 py-1.5 ${
                  status === s.key ? 'bg-primary' : 'bg-muted'
                }`}
              >
                <Text
                  className={`text-xs font-semibold ${
                    status === s.key ? 'text-primary-foreground' : 'text-foreground'
                  }`}
                >
                  {s.label}
                </Text>
              </Pressable>
            ))}
          </View>
        </FormField>

        <FormField label="Observações">
          <Input
            value={notes}
            onChangeText={setNotes}
            multiline
            numberOfLines={3}
            editable={!update.isPending}
          />
        </FormField>

        <FormField label="Nome da sessão">
          <Input
            value={sessionName}
            onChangeText={setSessionName}
            editable={!update.isPending}
          />
        </FormField>

        <View className="gap-2">
          <Text variant="caption">Exercícios ({exercises.length})</Text>
          {exercises.map((ex, idx) => (
            <Card key={`${ex.exerciseId}-${idx}`}>
              <View className="flex-row items-center justify-between">
                <Text variant="title" className="flex-1 pr-2">
                  {ex.exerciseName}
                </Text>
                <Button
                  size="sm"
                  variant="ghost"
                  onPress={() => removeExercise(idx)}
                  disabled={update.isPending}
                >
                  Remover
                </Button>
              </View>
              <View className="mt-2 flex-row gap-2">
                <View className="flex-1">
                  <Text variant="caption">Séries</Text>
                  <Input
                    value={String(ex.sets)}
                    onChangeText={(v) => updateExercise(idx, { sets: Number(v) || 0 })}
                    keyboardType="number-pad"
                    editable={!update.isPending}
                  />
                </View>
                <View className="flex-1">
                  <Text variant="caption">Reps</Text>
                  <Input
                    value={ex.reps}
                    onChangeText={(v) => updateExercise(idx, { reps: v })}
                    editable={!update.isPending}
                  />
                </View>
                <View className="flex-1">
                  <Text variant="caption">Descanso</Text>
                  <Input
                    value={ex.rest}
                    onChangeText={(v) => updateExercise(idx, { rest: v })}
                    editable={!update.isPending}
                  />
                </View>
              </View>
            </Card>
          ))}

          <Button
            variant="outline"
            size="sm"
            onPress={() => setExercisePickerOpen(true)}
            disabled={update.isPending}
          >
            + Adicionar exercício
          </Button>
        </View>

        <Button
          onPress={() => {
            if (!name.trim()) {
              toast.show('Nome obrigatório', 'warning');
              return;
            }
            update.mutate();
          }}
          loading={update.isPending}
          fullWidth
          size="lg"
        >
          Salvar alterações
        </Button>
      </ScrollView>

      <Sheet open={exercisePickerOpen} onClose={() => setExercisePickerOpen(false)}>
        <Text variant="title" className="mb-2">
          Escolher exercício
        </Text>
        <Input
          placeholder="Buscar..."
          value={search}
          onChangeText={setSearch}
          autoCapitalize="none"
        />
        <View className="mt-3 max-h-96">
          {exerciseList.isLoading ? (
            <Spinner />
          ) : (
            <View className="gap-1">
              {(exerciseList.data ?? []).slice(0, 30).map((ex) => (
                <Pressable
                  key={ex.id}
                  onPress={() => addExercise(ex)}
                  className="rounded-md bg-muted px-3 py-2"
                >
                  <Text variant="body">{ex.name}</Text>
                  <Text variant="caption">{ex.category ?? '—'}</Text>
                </Pressable>
              ))}
            </View>
          )}
        </View>
      </Sheet>
    </SafeAreaView>
  );
}
