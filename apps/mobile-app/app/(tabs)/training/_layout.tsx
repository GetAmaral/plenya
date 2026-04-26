import { Stack } from 'expo-router';

export default function TrainingLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Meus treinos' }} />
      <Stack.Screen name="[planId]/index" options={{ title: 'Plano de treino' }} />
      <Stack.Screen name="sessions/[id]" options={{ title: 'Sessão' }} />
      <Stack.Screen name="history" options={{ title: 'Histórico' }} />
    </Stack>
  );
}
