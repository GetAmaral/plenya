import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams } from 'expo-router';
import { useQuery } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { Card, CardHeader, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../lib/security/screenCapture';

export default function PatientDetailScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patient = useQuery(options.patientOptions(id ?? ''));

  if (patient.isLoading) return <Spinner centered />;
  if (!patient.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <View className="p-6">
          <Text variant="caption">Paciente não encontrado</Text>
        </View>
      </SafeAreaView>
    );
  }

  const p = patient.data;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-3 p-4">
        <Text variant="heading">{p.name}</Text>
        <Text variant="caption">
          {p.birthDate ? `Nasc. ${formatDate(p.birthDate)}` : 'Data de nascimento não informada'}
        </Text>

        <Card>
          <CardHeader>
            <Text variant="title">Contato</Text>
          </CardHeader>
          <Text variant="body">{p.phone ?? '—'}</Text>
          <Text variant="caption">{p.email ?? '—'}</Text>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Documento</Text>
          </CardHeader>
          <Text variant="body">CPF: {p.cpfMasked ?? '—'}</Text>
        </Card>

        <Card>
          <CardHeader>
            <Text variant="title">Próximos passos</Text>
          </CardHeader>
          <Text variant="caption">
            Abas de anamnese, exames, prescrições e escores serão adicionadas no Sprint 2.
          </Text>
        </Card>
      </ScrollView>
    </SafeAreaView>
  );
}
