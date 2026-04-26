import { Stack } from 'expo-router';

export default function AgendaLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Agenda' }} />
      <Stack.Screen name="new" options={{ title: 'Nova consulta' }} />
      <Stack.Screen name="[id]" options={{ title: 'Consulta' }} />
    </Stack>
  );
}
