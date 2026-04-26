import { useEffect, useRef, useState } from 'react';
import { View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { Button, Spinner, Text } from '@plenya/ui-mobile';
import { options } from '@plenya/api-client';
import { useAuthStore } from '../../features/auth/authStore';

/**
 * ConsumeLinkScreen — atendida via deep link `plenya-app://consume-link?token=xxx`
 * (custom scheme) ou Universal Link `https://app.plenyasaude.com.br/consume-link?token=xxx`.
 *
 * Aceita também ?type=invite pra usar o endpoint de convite (1ª senha) em vez
 * do magic-link normal.
 *
 * Idempotente: se for renderizado 2x (deep link + cold start), o segundo run
 * vê accessToken já setado e redireciona.
 */
export default function ConsumeLinkScreen() {
  const params = useLocalSearchParams<{ token?: string; type?: string }>();
  const setTokens = useAuthStore((s) => s.setTokens);
  const setUser = useAuthStore((s) => s.setUser);
  const [error, setError] = useState<string | null>(null);
  const consumed = useRef(false);

  useEffect(() => {
    if (consumed.current) return;
    if (!params.token) {
      setError('Link inválido — token ausente.');
      return;
    }
    consumed.current = true;

    (async () => {
      try {
        const fn =
          params.type === 'invite'
            ? options.patientAuthMutations.consumeInvite
            : options.patientAuthMutations.consumeMagicLink;
        const res = await fn({ token: String(params.token) });
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
        setError(err instanceof Error ? err.message : 'Link inválido ou expirado');
      }
    })();
  }, [params.token, params.type, setTokens, setUser]);

  return (
    <SafeAreaView className="flex-1 bg-background">
      <View className="flex-1 items-center justify-center px-6 gap-4">
        {error ? (
          <>
            <Text variant="title">Não foi possível entrar</Text>
            <Text variant="body" className="text-center">
              {error}
            </Text>
            <Button onPress={() => router.replace('/(auth)/login')} fullWidth size="lg">
              Voltar pro login
            </Button>
          </>
        ) : (
          <>
            <Spinner />
            <Text variant="caption">Validando link…</Text>
          </>
        )}
      </View>
    </SafeAreaView>
  );
}
