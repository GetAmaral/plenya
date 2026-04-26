import { useState } from 'react';
import { KeyboardAvoidingView, Platform, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useMutation } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';

export default function PasswordScreen() {
  const toast = useToast();
  const [next, setNext] = useState('');
  const [confirm, setConfirm] = useState('');

  const save = useMutation({
    mutationFn: () =>
      options.patientMeMutations.setPassword({ password: next }),
    onSuccess: () => {
      toast.show('Senha atualizada', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  function handleSave() {
    if (next.length < 8) {
      toast.show('Senha precisa de ao menos 8 caracteres', 'error');
      return;
    }
    if (next !== confirm) {
      toast.show('Confirmação não bate', 'error');
      return;
    }
    save.mutate();
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        <ScrollView contentContainerClassName="gap-4 p-4">
          <Text variant="caption">
            Defina ou redefina sua senha de acesso. Mínimo 8 caracteres.
          </Text>

          <FormField label="Nova senha" required>
            <Input
              value={next}
              onChangeText={setNext}
              secureTextEntry
              textContentType="newPassword"
            />
          </FormField>
          <FormField label="Confirmar nova senha" required>
            <Input
              value={confirm}
              onChangeText={setConfirm}
              secureTextEntry
              textContentType="newPassword"
            />
          </FormField>

          <View className="gap-2 mt-2">
            <Button onPress={handleSave} loading={save.isPending} fullWidth size="lg">
              Salvar nova senha
            </Button>
            <Button onPress={() => router.back()} variant="ghost" fullWidth>
              Cancelar
            </Button>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
