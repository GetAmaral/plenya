import { useEffect } from 'react';
import { Redirect, Tabs } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { useAutoLogout } from '../../lib/security/autoLogout';
import { useLogout } from '../../features/auth/useLogout';
import { useAuthStore } from '../../features/auth/authStore';
import { registerDeviceForPush } from '../../lib/push/registerDevice';
import { usePushNotificationRouter } from '../../lib/push/handler';

/**
 * 5 tabs do app paciente: Início / Treino / Exames / Mensagens / Perfil.
 *
 * Diferenças vs mobile-pro:
 *   - 30min de auto-logout (vs 5min) — paciente não manipula dados de outros.
 *   - Sem tab "Pacientes" / "Leads" / "Agenda staff".
 */
export default function TabsLayout() {
  const accessToken = useAuthStore((s) => s.accessToken);
  const biometricUnlocked = useAuthStore((s) => s.biometricUnlocked);
  const logout = useLogout();
  useAutoLogout(logout); // 30min default no app paciente
  usePushNotificationRouter();

  const unread = useQuery({
    ...options.patientMeMessagesUnreadOptions(),
    enabled: Boolean(accessToken && biometricUnlocked),
  });

  useEffect(() => {
    if (accessToken && biometricUnlocked) {
      registerDeviceForPush().catch(() => {
        /* push opt-in: falha não bloqueia */
      });
    }
  }, [accessToken, biometricUnlocked]);

  if (!accessToken) return <Redirect href="/(auth)/login" />;
  if (!biometricUnlocked) return <Redirect href="/(auth)/biometric-unlock" />;

  const unreadCount = unread.data?.unread ?? 0;
  const unreadBadge =
    unreadCount > 0 ? String(unreadCount > 99 ? '99+' : unreadCount) : undefined;

  return (
    <Tabs
      screenOptions={{
        headerShown: true,
        tabBarActiveTintColor: '#417e8e',
      }}
    >
      <Tabs.Screen name="index" options={{ title: 'Início' }} />
      <Tabs.Screen name="training" options={{ title: 'Treino', headerShown: false }} />
      <Tabs.Screen name="exams" options={{ title: 'Exames', headerShown: false }} />
      <Tabs.Screen
        name="messages"
        options={{
          title: 'Mensagens',
          headerShown: false,
          tabBarBadge: unreadBadge,
        }}
      />
      <Tabs.Screen name="profile" options={{ title: 'Perfil', headerShown: false }} />
    </Tabs>
  );
}
