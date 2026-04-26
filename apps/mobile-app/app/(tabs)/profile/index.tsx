import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Button, Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { useLogout } from '../../../features/auth/useLogout';

function Row({ title, subtitle, onPress }: { title: string; subtitle?: string; onPress: () => void }) {
  return (
    <Pressable onPress={onPress} accessibilityRole="button" accessibilityLabel={title}>
      <Card>
        <Text variant="title">{title}</Text>
        {subtitle && <Text variant="caption">{subtitle}</Text>}
      </Card>
    </Pressable>
  );
}

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

        <Row
          title="Editar dados"
          subtitle="Telefone, email, endereço, contato de emergência"
          onPress={() => router.push('/(tabs)/profile/edit' as never)}
        />
        <Row
          title="Trocar senha"
          subtitle="Definir uma nova senha de acesso"
          onPress={() => router.push('/(tabs)/profile/password' as never)}
        />
        <Row
          title="Sessões ativas"
          subtitle="Dispositivos conectados à sua conta"
          onPress={() => router.push('/(tabs)/profile/sessions' as never)}
        />
        <Row
          title="Notificações"
          subtitle="Lembretes de consulta, mensagens e treino"
          onPress={() => router.push('/(tabs)/profile/notifications' as never)}
        />
        <Row
          title="Privacidade (LGPD)"
          subtitle="Exportar seus dados ou solicitar exclusão"
          onPress={() => router.push('/(tabs)/profile/lgpd' as never)}
        />

        <View className="mt-4">
          <Button onPress={logout} variant="outline" fullWidth>
            Sair
          </Button>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
