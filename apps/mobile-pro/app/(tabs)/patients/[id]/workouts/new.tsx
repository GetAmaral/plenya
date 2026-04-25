import { useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys, type ExerciseSummary } from '@plenya/api-client';
import {
  Button,
  Card,
  FormField,
  Input,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';
import { useEnsureSelectedPatient } from '../../../../../features/patients/useEnsureSelectedPatient';

interface DraftExercise {
  exerciseId: string;
  exerciseName: string;
  sets: number;
  reps: string;
  rest: string;
}

export default function NewWorkoutPlanScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const queryClient = useQueryClient();
  const toast = useToast();

  const [name, setName] = useState('');
  const [notes, setNotes] = useState('');
  const [sessionName, setSessionName] = useState('Sessão A');
  const [exercises, setExercises] = useState<DraftExercise[]>([]);
  const [exercisePickerOpen, setExercisePickerOpen] = useState(false);
  const [search, setSearch] = useState('');

  const exerciseList = useQuery(options.exercisesListOptions({ search }));

  const create = useMutation({
    mutationFn: () =>
      options.workoutPlanMutations.create({
        name: name.trim(),
        status: 'active',
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
                  })),
                },
              ]
            : undefined,
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: queryKeys.patients.workoutPlans(patientId),
      });
      toast.show('Plano criado', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error'),
  });

  function addExercise(ex: ExerciseSummary) {
    setExercises((prev) => [
      ...prev,
      {
        exerciseId: ex.id,
        exerciseName: ex.name,
        sets: 3,
        reps: '10',
        rest: '60s',
      },
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

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <FormField label="Nome do plano" required>
          <Input
            value={name}
            onChangeText={setName}
            placeholder="Ex: Hipertrofia 12 semanas"
            editable={!create.isPending}
          />
        </FormField>

        <FormField label="Observações">
          <Input
            value={notes}
            onChangeText={setNotes}
            placeholder="Notas gerais sobre o plano"
            multiline
            numberOfLines={3}
            editable={!create.isPending}
          />
        </FormField>

        <FormField label="Nome da sessão">
          <Input
            value={sessionName}
            onChangeText={setSessionName}
            placeholder="Sessão A"
            editable={!create.isPending}
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
                  disabled={create.isPending}
                >
                  Remover
                </Button>
              </View>
              <View className="mt-2 flex-row gap-2">
                <View className="flex-1">
                  <Text variant="caption">Séries</Text>
                  <Input
                    value={String(ex.sets)}
                    onChangeText={(v) =>
                      updateExercise(idx, { sets: Number(v) || 0 })
                    }
                    keyboardType="number-pad"
                    editable={!create.isPending}
                  />
                </View>
                <View className="flex-1">
                  <Text variant="caption">Reps</Text>
                  <Input
                    value={ex.reps}
                    onChangeText={(v) => updateExercise(idx, { reps: v })}
                    placeholder="10"
                    editable={!create.isPending}
                  />
                </View>
                <View className="flex-1">
                  <Text variant="caption">Descanso</Text>
                  <Input
                    value={ex.rest}
                    onChangeText={(v) => updateExercise(idx, { rest: v })}
                    placeholder="60s"
                    editable={!create.isPending}
                  />
                </View>
              </View>
            </Card>
          ))}

          <Button
            variant="outline"
            size="sm"
            onPress={() => setExercisePickerOpen(true)}
            disabled={create.isPending}
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
            create.mutate();
          }}
          loading={create.isPending}
          fullWidth
          size="lg"
        >
          Criar plano
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
