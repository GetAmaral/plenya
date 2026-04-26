import { Stack } from 'expo-router';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function ExamsLayout() {
  useScreenCaptureProtection();
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Exames e escores' }} />
      <Stack.Screen name="[id]" options={{ title: 'Resultado' }} />
    </Stack>
  );
}
