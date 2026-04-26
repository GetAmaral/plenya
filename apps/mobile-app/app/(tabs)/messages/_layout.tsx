import { Stack } from 'expo-router';

export default function MessagesLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Mensagens' }} />
    </Stack>
  );
}
