import { Stack } from 'expo-router';

export default function LeadsLayout() {
  return (
    <Stack>
      <Stack.Screen name="index" options={{ title: 'Leads' }} />
      <Stack.Screen name="[id]" options={{ title: 'Lead' }} />
    </Stack>
  );
}
