import { Redirect } from 'expo-router';
import { View } from 'react-native';
import { Spinner } from '@plenya/ui-mobile';
import { useAuthStore } from '../features/auth/authStore';

export default function Index() {
  const hydrated = useAuthStore((s) => s.hydrated);
  const accessToken = useAuthStore((s) => s.accessToken);
  const biometricUnlocked = useAuthStore((s) => s.biometricUnlocked);

  if (!hydrated) {
    return (
      <View className="flex-1 items-center justify-center bg-background">
        <Spinner />
      </View>
    );
  }

  if (!accessToken) return <Redirect href="/(auth)/login" />;
  if (!useAuthStore.getState().lgpdAccepted) return <Redirect href="/(auth)/lgpd-consent" />;
  if (!biometricUnlocked) return <Redirect href="/(auth)/biometric-unlock" />;
  return <Redirect href="/(tabs)" />;
}
