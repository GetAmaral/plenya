import { useCallback, useEffect, useState } from 'react';
import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { Button, Text, useToast } from '@plenya/ui-mobile';
import { authenticateBiometric, isBiometricAvailable } from '../../lib/security/biometric';
import { useAuthStore } from '../../features/auth/authStore';

export default function BiometricUnlockScreen() {
  const markUnlocked = useAuthStore((s) => s.markBiometricUnlocked);
  const toast = useToast();
  const [available, setAvailable] = useState<boolean | null>(null);

  const attemptAuth = useCallback(async () => {
    const res = await authenticateBiometric('Desbloqueie o Plenya Pro');
    if (res.ok) {
      markUnlocked();
      router.replace('/(tabs)');
    } else if (res.reason === 'unavailable') {
      markUnlocked();
      router.replace('/(tabs)');
    } else if (res.reason === 'failed') {
      toast.show('Autenticação falhou. Tente novamente.', 'error');
    }
  }, [markUnlocked, toast]);

  useEffect(() => {
    (async () => {
      const ok = await isBiometricAvailable();
      setAvailable(ok);
      if (!ok) {
        markUnlocked();
        router.replace('/(tabs)');
      } else {
        attemptAuth();
      }
    })();
  }, [attemptAuth, markUnlocked]);

  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 items-center justify-center px-6">
        <Text variant="heading" className="mb-2 text-center">
          Proteção biométrica
        </Text>
        <Text variant="caption" className="mb-8 text-center">
          Use Face ID, Touch ID ou sua digital para acessar dados clínicos.
        </Text>
        {available && (
          <Button onPress={attemptAuth} size="lg" fullWidth>
            Desbloquear
          </Button>
        )}
      </View>
    </SafeAreaView>
  );
}
