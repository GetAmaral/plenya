import { Stack } from 'expo-router';

export default function ConversationsLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Conversas' }} />
      <Stack.Screen name="[ownerType]/[ownerId]" options={{ title: 'Conversa' }} />
    </Stack>
  );
}
