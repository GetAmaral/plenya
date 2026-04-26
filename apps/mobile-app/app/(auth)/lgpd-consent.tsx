import { useState } from 'react';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQueryClient } from '@tanstack/react-query';
import { Button, Card, Text, useToast } from '@plenya/ui-mobile';
import { options } from '@plenya/api-client';
import { useAuthStore } from '../../features/auth/authStore';

export default function LGPDConsentScreen() {
  const queryClient = useQueryClient();
  const markLgpdAccepted = useAuthStore((s) => s.markLgpdAccepted);
  const toast = useToast();
  const [loading, setLoading] = useState(false);

  async function accept() {
    setLoading(true);
    try {
      await options.meMutations.acceptLgpdConsent();
      await markLgpdAccepted();
      await queryClient.invalidateQueries({ queryKey: ['plenya', 'me'] });
      router.replace('/(auth)/biometric-unlock');
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Falha ao registrar aceite';
      toast.show(message, 'error');
    } finally {
      setLoading(false);
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background">
      <ScrollView contentContainerClassName="gap-4 p-6">
        <Text variant="heading">Termo de uso e LGPD</Text>
        <Text variant="caption">
          Antes de acessar seus dados de saúde no Plenya, precisamos do seu aceite.
        </Text>

        <Card>
          <Text variant="title" className="mb-2">
            Como tratamos seus dados
          </Text>
          <Text variant="body" className="mb-2">
            • Seus dados de saúde são acessíveis apenas para você e a equipe da clínica
            que cuida de você, com audit log de cada operação.
          </Text>
          <Text variant="body" className="mb-2">
            • CPF e RG são criptografados em repouso (AES-256). Nunca expomos esses campos
            no app sem necessidade.
          </Text>
          <Text variant="body" className="mb-2">
            • Tokens de acesso ficam no Keychain/Keystore do device. Cache local é criptografado
            com chave derivada do hardware.
          </Text>
          <Text variant="body" className="mb-2">
            • Você pode revogar o acesso deste device em "Perfil → Dispositivos ativos" a qualquer
            momento.
          </Text>
          <Text variant="body">
            • Crash reports não contêm PHI — sanitizamos nome, CPF, dados clínicos antes de enviar.
          </Text>
        </Card>

        <Card>
          <Text variant="title" className="mb-2">
            Seus deveres
          </Text>
          <Text variant="body" className="mb-2">
            • Não compartilhe sua sessão. Cada device é uma sessão revogável.
          </Text>
          <Text variant="body" className="mb-2">
            • Mantenha biometria e senha ativas no device.
          </Text>
          <Text variant="body">
            • Reporte qualquer suspeita de acesso indevido para suporte@plenyasaude.com.br.
          </Text>
        </Card>

        <View className="gap-2">
          <Button onPress={accept} loading={loading} fullWidth size="lg">
            Aceitar e continuar
          </Button>
          <Button
            variant="ghost"
            onPress={() => router.replace('/(auth)/login')}
            disabled={loading}
            fullWidth
            size="md"
          >
            Sair sem aceitar
          </Button>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
