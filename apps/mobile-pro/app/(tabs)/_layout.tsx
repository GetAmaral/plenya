import { Redirect, Tabs } from 'expo-router';
import { useAutoLogout } from '../../lib/security/autoLogout';
import { useLogout } from '../../features/auth/useLogout';
import { useAuthStore } from '../../features/auth/authStore';

export default function TabsLayout() {
  const accessToken = useAuthStore((s) => s.accessToken);
  const biometricUnlocked = useAuthStore((s) => s.biometricUnlocked);
  const logout = useLogout();
  useAutoLogout(logout, 5 * 60 * 1000);

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
      <Tabs.Screen name="agenda" options={{ title: 'Agenda' }} />
      <Tabs.Screen name="leads" options={{ title: 'Leads', headerShown: false }} />
      <Tabs.Screen name="profile" options={{ title: 'Perfil' }} />
    </Tabs>
  );
}
