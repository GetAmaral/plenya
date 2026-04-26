import { useState } from 'react';
import { KeyboardAvoidingView, Platform, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { options } from '@plenya/api-client';

/**
 * Magic-link request — paciente informa email, recebe link clicável.
 * Resposta é sempre 204 (anti-enumeração) — não confirmamos se email existe.
 *
 * O link recebido abre `plenya-app://consume-link?token=...` (ou o universal
 * link `https://app.plenyasaude.com.br/consume-link?token=...` quando ativo)
 * → ConsumeLinkScreen troca pelo JWT.
 */
export default function MagicLinkScreen() {
  const toast = useToast();
  const [email, setEmail] = useState('');
  const [loading, setLoading] = useState(false);
  const [sent, setSent] = useState(false);

  async function handleSubmit() {
    if (!email) return;
    setLoading(true);
    try {
      await options.patientAuthMutations.requestMagicLink({ email });
      setSent(true);
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Erro ao enviar';
      toast.show(message, 'error');
    } finally {
      setLoading(false);
    }
  }

  if (sent) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <View className="flex-1 justify-center px-6 gap-4">
          <Text variant="display">Verifique seu email</Text>
          <Text variant="body">
            Se o email <Text className="font-semibold">{email}</Text> está cadastrado,
            enviamos um link de acesso. Toque no link no seu celular pra entrar.
          </Text>
          <Text variant="caption">
            Não chegou em alguns minutos? Verifique a caixa de spam ou tente outra vez.
          </Text>
          <Button onPress={() => router.back()} fullWidth size="lg">
            Voltar
          </Button>
        </View>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background">
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        <View className="flex-1 justify-center px-6 gap-4">
          <Text variant="display">Receber link</Text>
          <Text variant="caption">
            Enviaremos um link único pro seu email. Toque pra entrar — sem senha.
          </Text>

          <FormField label="Email" required>
            <Input
              value={email}
              onChangeText={setEmail}
              placeholder="seu@email.com.br"
              autoCapitalize="none"
              keyboardType="email-address"
              textContentType="emailAddress"
              editable={!loading}
            />
          </FormField>

          <Button onPress={handleSubmit} loading={loading} fullWidth size="lg">
            Enviar link
          </Button>
          <Button onPress={() => router.back()} variant="ghost" fullWidth>
            Cancelar
          </Button>
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
