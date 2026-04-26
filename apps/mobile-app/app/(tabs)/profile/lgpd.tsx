import { useState } from 'react';
import { Alert, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import * as FileSystem from 'expo-file-system';
import * as Sharing from 'expo-sharing';
import { useMutation } from '@tanstack/react-query';
import { api, options } from '@plenya/api-client';
import { Button, Card, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { useLogout } from '../../../features/auth/useLogout';

/**
 * LGPD: export JSON via share sheet + solicitação de exclusão de conta.
 * Exclusão é "request" — backend cria pedido pra equipe processar (LGPD prevê
 * 15 dias úteis); paciente é deslogado em seguida.
 */

export default function LgpdScreen() {
  const toast = useToast();
  const logout = useLogout();
  const [reason, setReason] = useState('');

  const exportData = useMutation({
    mutationFn: async () => {
      const baseDir = FileSystem.documentDirectory;
      if (!baseDir) {
        throw new Error('Armazenamento local indisponível.');
      }
      const data = await api.get<unknown>('/api/v1/patient/me/lgpd/export');
      const path = `${baseDir}plenya-meus-dados.json`;
      await FileSystem.writeAsStringAsync(path, JSON.stringify(data, null, 2), {
        encoding: FileSystem.EncodingType.UTF8,
      });
      return path;
    },
    onSuccess: async (path) => {
      if (await Sharing.isAvailableAsync()) {
        await Sharing.shareAsync(path, {
          mimeType: 'application/json',
          dialogTitle: 'Meus dados Plenya',
        });
      } else {
        toast.show(`Salvo em: ${path}`, 'success');
      }
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  const requestDelete = useMutation({
    mutationFn: () =>
      options.patientMeMutations.lgpdRequestDelete({
        reason: reason || undefined,
      }),
    onSuccess: async () => {
      toast.show('Solicitação enviada à equipe.', 'success');
      await logout();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  function confirmDelete() {
    Alert.alert(
      'Solicitar exclusão',
      'Sua conta será desativada e seus dados clínicos arquivados conforme o prazo legal. Tem certeza?',
      [
        { text: 'Cancelar', style: 'cancel' },
        {
          text: 'Solicitar exclusão',
          style: 'destructive',
          onPress: () => requestDelete.mutate(),
        },
      ],
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <Text variant="title">Exportar meus dados</Text>
          <Text variant="caption" className="mb-3">
            Baixe um arquivo JSON com todos os seus dados Plenya. Você pode salvar
            no celular ou enviar pra outro app.
          </Text>
          <Button
            onPress={() => exportData.mutate()}
            loading={exportData.isPending}
            fullWidth
          >
            Exportar (JSON)
          </Button>
        </Card>

        <Card>
          <Text variant="title">Excluir minha conta</Text>
          <Text variant="caption" className="mb-3">
            Você pode solicitar a exclusão da sua conta a qualquer momento. A
            equipe processa em até 15 dias úteis e respeitando obrigações legais
            de retenção de prontuário.
          </Text>
          <FormField label="Motivo (opcional)">
            <Input
              value={reason}
              onChangeText={setReason}
              multiline
              placeholder="Ajuda a melhorar a Plenya"
            />
          </FormField>
          <View className="mt-3">
            <Button
              onPress={confirmDelete}
              loading={requestDelete.isPending}
              variant="destructive"
              fullWidth
            >
              Solicitar exclusão
            </Button>
          </View>
        </Card>
      </ScrollView>
    </SafeAreaView>
  );
}
