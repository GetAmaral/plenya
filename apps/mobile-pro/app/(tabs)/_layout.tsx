import { useEffect } from 'react';
import { Redirect, Tabs } from 'expo-router';
import { useAutoLogout } from '../../lib/security/autoLogout';
import { useLogout } from '../../features/auth/useLogout';
import { useAuthStore } from '../../features/auth/authStore';
import { registerDeviceForPush } from '../../lib/push/registerDevice';

export default function TabsLayout() {
  const accessToken = useAuthStore((s) => s.accessToken);
  const biometricUnlocked = useAuthStore((s) => s.biometricUnlocked);
  const logout = useLogout();
  useAutoLogout(logout, 5 * 60 * 1000);

  useEffect(() => {
    if (accessToken && biometricUnlocked) {
      registerDeviceForPush().catch(() => {
        /* no-op: push é opt-in, falha não bloqueia uso */
      });
    }
  }, [accessToken, biometricUnlocked]);

  if (!accessToken) return <Redirect href="/(auth)/login" />;
  if (!biometricUnlocked) return <Redirect href="/(auth)/biometric-unlock" />;

  return (
    <Tabs
      screenOptions={{
        headerShown: true,
        tabBarActiveTintColor: '#10b981',
      }}
    >
      <Tabs.Screen name="index" options={{ title: 'Início' }} />
      <Tabs.Screen name="patients" options={{ title: 'Pacientes', headerShown: false }} />
      <Tabs.Screen name="training" options={{ title: 'Treino', headerShown: false }} />
      <Tabs.Screen name="leads" options={{ title: 'Leads', headerShown: false }} />
      <Tabs.Screen name="profile" options={{ title: 'Perfil' }} />
    </Tabs>
  );
}
