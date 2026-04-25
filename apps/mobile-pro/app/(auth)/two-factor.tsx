import { useState } from 'react';
import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { options } from '@plenya/api-client';
import { useAuthStore } from '../../features/auth/authStore';

export default function TwoFactorScreen() {
  const { email, pendingToken } = useLocalSearchParams<{
    email: string;
    pendingToken: string;
  }>();
  const setTokens = useAuthStore((s) => s.setTokens);
  const setUser = useAuthStore((s) => s.setUser);
  const toast = useToast();
  const [code, setCode] = useState('');
  const [loading, setLoading] = useState(false);

  async function handleSubmit() {
    if (code.length !== 6 || !email || !pendingToken) return;
    setLoading(true);
    try {
      const res = await options.authMutations.verifyTwoFactor({
        email,
        code,
        pendingToken,
      });
      await setTokens(res.accessToken, res.refreshToken);
      setUser({ ...res.user, has2FAEnabled: true });
      router.replace('/(auth)/biometric-unlock');
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Código inválido';
      toast.show(message, 'error');
    } finally {
      setLoading(false);
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 justify-center px-6">
        <Text variant="heading" className="mb-1">
          Verificação em 2 etapas
        </Text>
        <Text variant="caption" className="mb-8">
          Digite o código de 6 dígitos do seu aplicativo autenticador.
        </Text>

        <FormField label="Código" required>
          <Input
            value={code}
            onChangeText={(v) => setCode(v.replace(/\D/g, '').slice(0, 6))}
            placeholder="000000"
            keyboardType="number-pad"
            maxLength={6}
            editable={!loading}
          />
        </FormField>

        <View className="mt-4">
          <Button onPress={handleSubmit} loading={loading} fullWidth size="lg">
            Verificar
          </Button>
        </View>
      </View>
    </SafeAreaView>
  );
}
