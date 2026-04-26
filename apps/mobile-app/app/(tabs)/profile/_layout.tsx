import { Stack } from 'expo-router';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function ProfileLayout() {
  useScreenCaptureProtection();
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Meu perfil' }} />
    </Stack>
  );
}
