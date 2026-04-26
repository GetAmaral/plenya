import { Stack } from 'expo-router';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function ProfileLayout() {
  useScreenCaptureProtection();
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Meu perfil' }} />
      <Stack.Screen name="notifications" options={{ title: 'Notificações' }} />
      <Stack.Screen name="edit" options={{ title: 'Editar dados' }} />
      <Stack.Screen name="password" options={{ title: 'Trocar senha' }} />
      <Stack.Screen name="sessions" options={{ title: 'Sessões ativas' }} />
      <Stack.Screen name="lgpd" options={{ title: 'Privacidade (LGPD)' }} />
    </Stack>
  );
}
