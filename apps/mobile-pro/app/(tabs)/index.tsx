import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { Card, CardHeader, Text } from '@plenya/ui-mobile';

export default function DashboardScreen() {
  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="heading">Olá, doutor(a)</Text>
        <Text variant="caption">
          Bem-vindo ao Plenya Pro. Esta é a tela inicial do profissional — será preenchida com
          atalhos para pacientes selecionados, agenda do dia e notificações críticas.
        </Text>

        <View className="gap-3">
          <Card>
            <CardHeader>
              <Text variant="title">Paciente selecionado</Text>
            </CardHeader>
            <Text variant="caption">Nenhum paciente ativo</Text>
          </Card>

          <Card>
            <CardHeader>
              <Text variant="title">Agenda do dia</Text>
            </CardHeader>
            <Text variant="caption">Sem consultas agendadas</Text>
          </Card>

          <Card>
            <CardHeader>
              <Text variant="title">Leads novos</Text>
            </CardHeader>
            <Text variant="caption">Conecte-se para ver</Text>
          </Card>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
