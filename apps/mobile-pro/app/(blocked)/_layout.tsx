import { Stack } from 'expo-router';

export default function BlockedLayout() {
  return (
    <Stack
      screenOptions={{
        headerShown: false,
        animation: 'fade',
        contentStyle: { backgroundColor: '#0f172a' },
      }}
    />
  );
}
