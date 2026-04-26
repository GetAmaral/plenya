import { Stack } from 'expo-router';

export default function TrainingLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Biblioteca' }} />
      <Stack.Screen name="exercises/[id]" options={{ title: 'Exercício' }} />
      <Stack.Screen name="workout-plans/[id]/index" options={{ title: 'Plano de treino' }} />
      <Stack.Screen name="workout-plans/[id]/edit" options={{ title: 'Editar plano' }} />
    </Stack>
  );
}
