import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { useLogout } from '../../../features/auth/useLogout';

export default function ProfileScreen() {
  const profile = useQuery(options.patientMeProfileOptions());
  const logout = useLogout();

  if (profile.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <CardHeader>
            <Text variant="title">{profile.data?.name ?? 'Paciente'}</Text>
          </CardHeader>
          {profile.data?.email && <Text variant="caption">{profile.data.email}</Text>}
          {profile.data?.phone && <Text variant="caption">{profile.data.phone}</Text>}
        </Card>

        <Text variant="caption" className="italic">
          Avatar, senha, 2FA, sessões ativas, preferências de notificação,
          export LGPD e delete chegam na Sprint 4.
        </Text>

        <View className="mt-4">
          <Button onPress={logout} variant="outline" fullWidth>
            Sair
          </Button>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
