import { RefreshControl, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useQueries, useQuery } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, Card, CardHeader, ErrorState, Spinner, Text } from '@plenya/ui-mobile';
import { formatDate } from '@plenya/domain';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useRefresh } from '../../../../features/patients/usePatientRefresh';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';
import { SubsectionLink } from '../../../../components/patient/SubsectionLink';

export default function PatientOverviewScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const patient = useQuery(options.patientOptions(patientId));

  const counts = useQueries({
    queries: [
      {
        ...options.patientAnamnesisOptions(patientId),
        select: (data: unknown) => (Array.isArray(data) ? data.length : 0),
      },
      {
        ...options.patientLabResultsOptions(patientId),
        select: (data: unknown) => (Array.isArray(data) ? data.length : 0),
      },
      {
        ...options.patientPrescriptionsOptions(patientId),
        select: (data: unknown) => (Array.isArray(data) ? data.length : 0),
      },
    ],
  });

  const { refreshing, onRefresh } = useRefresh([
    queryKeys.patients.detail(patientId),
    queryKeys.patients.anamnesis(patientId),
    queryKeys.patients.labResults(patientId),
    queryKeys.patients.prescriptions(patientId),
    queryKeys.patients.scores(patientId),
  ]);

  if (patient.isLoading) return <Spinner centered />;
  if (patient.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => patient.refetch()} />
      </SafeAreaView>
    );
  }

  const p = patient.data;
  if (!p) return null;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView
        contentContainerClassName="gap-3 p-4"
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={onRefresh} />}
      >
        <View className="flex-row items-start justify-between gap-3">
          <View className="flex-1">
            <Text variant="heading">{p.name}</Text>
            <Text variant="caption">
              {p.birthDate ? `Nasc. ${formatDate(p.birthDate)}` : 'Data de nascimento não informada'}
            </Text>
          </View>
          <Button
            variant="outline"
            size="sm"
            onPress={() => router.push(`/(tabs)/patients/${patientId}/edit` as never)}
          >
            Editar
          </Button>
        </View>

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
          {p.rgMasked && <Text variant="caption">RG: {p.rgMasked}</Text>}
        </Card>

        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/anamnesis`}
          title="Anamnese"
          description="Histórico clínico do paciente"
          count={counts[0].data}
        />
        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/labs`}
          title="Exames laboratoriais"
          description="Resultados com valores e referências"
          count={counts[1].data}
        />
        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/prescriptions`}
          title="Prescrições"
          description="Receitas ativas e históricas"
          count={counts[2].data}
        />
        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/scores`}
          title="Escores de saúde"
          description="Estratificação por grupos e itens clínicos"
        />
        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/workouts`}
          title="Planos de treino"
          description="Sessões e exercícios prescritos"
        />
        <SubsectionLink
          href={`/(tabs)/patients/${patientId}/assessments`}
          title="Avaliações físicas"
          description="Composição corporal, PA, tags ACSM"
        />
      </ScrollView>
    </SafeAreaView>
  );
}
