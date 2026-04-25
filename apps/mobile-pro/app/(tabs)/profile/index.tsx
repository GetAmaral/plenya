import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { useLogout } from '../../../features/auth/useLogout';

export default function ProfileScreen() {
  const me = useQuery(options.meOptions());
  const sessions = useQuery(options.meSessionsOptions());
  const logout = useLogout();

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
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Dispositivos ativos</Text>
          </CardHeader>
          <View className="gap-2">
            {(sessions.data ?? []).map((s) => (
              <View key={s.id}>
                <Text variant="body">
                  {s.device} · {s.platform}
                  {s.current ? ' (este)' : ''}
                </Text>
                <Text variant="caption">
                  {s.appVariant} v{s.appVersion ?? '—'}
                </Text>
              </View>
            ))}
            {(sessions.data ?? []).length === 0 && <Text variant="caption">Nenhuma sessão</Text>}
          </View>
        </Card>

        <Button variant="destructive" onPress={logout} fullWidth>
          Sair
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
