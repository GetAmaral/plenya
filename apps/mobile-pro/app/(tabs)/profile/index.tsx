import { useState } from 'react';
import { ScrollView, Share, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import * as Application from 'expo-application';
import * as Updates from 'expo-updates';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import {
  Button,
  Card,
  CardHeader,
  Dialog,
  FormField,
  Input,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { useLogout } from '../../../features/auth/useLogout';

const APP_VERSION = Application.nativeApplicationVersion ?? '0.0.0';
const BUILD_VERSION = Application.nativeBuildVersion ?? '—';
const UPDATE_ID = Updates.updateId ? Updates.updateId.slice(0, 8) : 'embedded';
const UPDATE_CHANNEL = Updates.channel ?? 'development';

export default function ProfileScreen() {
  const queryClient = useQueryClient();
  const toast = useToast();
  const logout = useLogout();

  const me = useQuery(options.meOptions());
  const sessions = useQuery(options.meSessionsOptions());

  const [pwdOpen, setPwdOpen] = useState(false);
  const [twoFAOpen, setTwoFAOpen] = useState(false);
  const [currentPwd, setCurrentPwd] = useState('');
  const [newPwd, setNewPwd] = useState('');
  const [confirmPwd, setConfirmPwd] = useState('');
  const [twoFAPwd, setTwoFAPwd] = useState('');

  const revokeSession = useMutation({
    mutationFn: (sessionId: string) => options.meMutations.revokeSession(sessionId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.meSessions() });
      toast.show('Dispositivo revogado', 'success');
    },
    onError: (err) => toast.show(err instanceof Error ? err.message : 'Falha', 'error'),
  });

  const changePwd = useMutation({
    mutationFn: () =>
      options.meMutations.changePassword({
        currentPassword: currentPwd,
        newPassword: newPwd,
      }),
    onSuccess: () => {
      setPwdOpen(false);
      setCurrentPwd('');
      setNewPwd('');
      setConfirmPwd('');
      toast.show('Senha alterada', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao trocar senha', 'error'),
  });

  const disable2FA = useMutation({
    mutationFn: () => options.meMutations.disable2FA({ password: twoFAPwd }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.me() });
      setTwoFAOpen(false);
      setTwoFAPwd('');
      toast.show('2FA desabilitado', 'success');
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao desabilitar 2FA', 'error'),
  });

  const exportData = useMutation({
    mutationFn: () => options.meMutations.exportData(),
    onSuccess: async (data) => {
      const json = JSON.stringify(data, null, 2);
      await Share.share({
        title: 'Plenya — meus dados',
        message: json,
      });
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao exportar', 'error'),
  });

  function handleSubmitPwd() {
    if (newPwd.length < 8) {
      toast.show('Nova senha precisa de ao menos 8 caracteres', 'warning');
      return;
    }
    if (newPwd !== confirmPwd) {
      toast.show('Confirmação não bate', 'warning');
      return;
    }
    changePwd.mutate();
  }

  if (me.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Card>
          <CardHeader>
            <Text variant="title">{me.data?.name ?? 'Usuário'}</Text>
          </CardHeader>
          <Text variant="caption">{me.data?.email}</Text>
          <Text variant="caption">Role: {me.data?.role}</Text>
          <Text variant="caption">2FA: {me.data?.has2FAEnabled ? 'ativo' : 'inativo'}</Text>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Segurança</Text>
          </CardHeader>
          <View className="gap-2">
            <Button variant="outline" size="sm" onPress={() => setPwdOpen(true)}>
              Trocar senha
            </Button>
            {me.data?.has2FAEnabled && (
              <Button
                variant="outline"
                size="sm"
                onPress={() => setTwoFAOpen(true)}
              >
                Desabilitar 2FA
              </Button>
            )}
          </View>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Dispositivos ativos</Text>
          </CardHeader>
          <View className="gap-2">
            {(sessions.data ?? []).map((s) => (
              <View key={s.id} className="flex-row items-start justify-between gap-2">
                <View className="flex-1">
                  <Text variant="body">
                    {s.device || 'Dispositivo'} · {s.platform}
                    {s.current ? ' (este)' : ''}
                  </Text>
                  <Text variant="caption">
                    {s.appVariant} v{s.appVersion ?? '—'}
                  </Text>
                </View>
                {!s.current && (
                  <Button
                    size="sm"
                    variant="ghost"
                    onPress={() => revokeSession.mutate(s.id)}
                    loading={revokeSession.isPending}
                  >
                    Revogar
                  </Button>
                )}
              </View>
            ))}
            {(sessions.data ?? []).length === 0 && (
              <Text variant="caption">Nenhuma sessão</Text>
            )}
          </View>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Privacidade (LGPD)</Text>
          </CardHeader>
          <Text variant="caption" className="mb-2">
            Você pode exportar todos os seus dados pessoais armazenados pelo Plenya.
          </Text>
          <Button
            variant="outline"
            size="sm"
            onPress={() => exportData.mutate()}
            loading={exportData.isPending}
          >
            Exportar meus dados
          </Button>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Sobre o app</Text>
          </CardHeader>
          <Text variant="body">
            Versão {APP_VERSION} (build {BUILD_VERSION})
          </Text>
          <Text variant="caption">
            Canal {UPDATE_CHANNEL} · OTA {UPDATE_ID}
          </Text>
        </Card>

        <Button variant="destructive" onPress={logout} fullWidth>
          Sair
        </Button>
      </ScrollView>

      <Dialog open={pwdOpen} onClose={() => setPwdOpen(false)}>
        <Text variant="title" className="mb-3">
          Trocar senha
        </Text>
        <View className="gap-3">
          <FormField label="Senha atual" required>
            <Input
              secureTextEntry
              value={currentPwd}
              onChangeText={setCurrentPwd}
              editable={!changePwd.isPending}
              autoCapitalize="none"
              textContentType="password"
            />
          </FormField>
          <FormField label="Nova senha" required helper="Mínimo 8 caracteres">
            <Input
              secureTextEntry
              value={newPwd}
              onChangeText={setNewPwd}
              editable={!changePwd.isPending}
              autoCapitalize="none"
              textContentType="newPassword"
            />
          </FormField>
          <FormField label="Confirmar nova senha" required>
            <Input
              secureTextEntry
              value={confirmPwd}
              onChangeText={setConfirmPwd}
              editable={!changePwd.isPending}
              autoCapitalize="none"
              textContentType="newPassword"
            />
          </FormField>
        </View>
        <View className="mt-4 flex-row gap-2">
          <Button
            variant="outline"
            onPress={() => setPwdOpen(false)}
            disabled={changePwd.isPending}
            className="flex-1"
          >
            Cancelar
          </Button>
          <Button
            onPress={handleSubmitPwd}
            loading={changePwd.isPending}
            className="flex-1"
          >
            Trocar
          </Button>
        </View>
      </Dialog>

      <Dialog open={twoFAOpen} onClose={() => setTwoFAOpen(false)}>
        <Text variant="title" className="mb-3">
          Desabilitar 2FA
        </Text>
        <Text variant="body" className="mb-3">
          Confirme sua senha pra remover a verificação em 2 etapas. Reduz a segurança da
          sua conta.
        </Text>
        <FormField label="Senha" required>
          <Input
            secureTextEntry
            value={twoFAPwd}
            onChangeText={setTwoFAPwd}
            editable={!disable2FA.isPending}
            autoCapitalize="none"
            textContentType="password"
          />
        </FormField>
        <View className="mt-4 flex-row gap-2">
          <Button
            variant="outline"
            onPress={() => setTwoFAOpen(false)}
            disabled={disable2FA.isPending}
            className="flex-1"
          >
            Cancelar
          </Button>
          <Button
            variant="destructive"
            onPress={() => disable2FA.mutate()}
            loading={disable2FA.isPending}
            className="flex-1"
          >
            Desabilitar
          </Button>
        </View>
      </Dialog>
    </SafeAreaView>
  );
}
