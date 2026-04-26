import { Stack } from 'expo-router';

export default function TrainingLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Meus treinos' }} />
    </Stack>
  );
}
