import { useState } from 'react';
import { KeyboardAvoidingView, Platform, Pressable, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { options } from '@plenya/api-client';
import { useAuthStore } from '../../features/auth/authStore';

export default function LoginScreen() {
  const setTokens = useAuthStore((s) => s.setTokens);
  const setUser = useAuthStore((s) => s.setUser);
  const toast = useToast();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);

  async function handleSubmit() {
    if (!email || !password) return;
    setLoading(true);
    try {
      const res = await options.patientAuthMutations.loginPassword({ email, password });
      await setTokens(res.accessToken, res.refreshToken);
      setUser({
        id: res.user.id,
        name: res.user.name,
        email: res.user.email,
        role: res.user.roles[0] ?? 'patient',
        has2FAEnabled: res.user.twoFactorEnabled,
      });
      router.replace(
        useAuthStore.getState().lgpdAccepted
          ? '/(auth)/biometric-unlock'
          : '/(auth)/lgpd-consent',
      );
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Erro ao entrar';
      toast.show(message, 'error');
    } finally {
      setLoading(false);
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background">
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        <View className="flex-1 justify-center px-6">
          <Text variant="display" className="mb-1">
            Plenya
          </Text>
          <Text variant="caption" className="mb-8">
            Sua jornada de saúde, no seu bolso.
          </Text>

          <View className="gap-4">
            <FormField label="Email" required>
              <Input
                testID="email"
                value={email}
                onChangeText={setEmail}
                placeholder="seu@email.com.br"
                autoCapitalize="none"
                keyboardType="email-address"
                textContentType="emailAddress"
                editable={!loading}
              />
            </FormField>

            <FormField label="Senha" required>
              <Input
                testID="password"
                value={password}
                onChangeText={setPassword}
                secureTextEntry
                placeholder="••••••••"
                textContentType="password"
                editable={!loading}
              />
            </FormField>

            <Button onPress={handleSubmit} loading={loading} fullWidth size="lg">
              Entrar
            </Button>

            <Pressable
              onPress={() => router.push('/(auth)/magic-link')}
              accessibilityRole="button"
              accessibilityLabel="Entrar com link mágico por email"
            >
              <Text variant="caption" className="text-center underline">
                Esqueci a senha · receber link por email
              </Text>
            </Pressable>
          </View>
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
