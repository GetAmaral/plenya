import { Stack } from 'expo-router';

export default function PatientsLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Pacientes' }} />
      <Stack.Screen name="new" options={{ title: 'Novo paciente' }} />
      <Stack.Screen name="[id]/index" options={{ title: 'Paciente' }} />
      <Stack.Screen name="[id]/edit" options={{ title: 'Editar paciente' }} />
      <Stack.Screen name="[id]/anamnesis/index" options={{ title: 'Anamnese' }} />
      <Stack.Screen name="[id]/anamnesis/new" options={{ title: 'Nova anamnese' }} />
      <Stack.Screen name="[id]/anamnesis/[anamnesisId]" options={{ title: 'Anamnese' }} />
      <Stack.Screen name="[id]/labs/index" options={{ title: 'Exames' }} />
      <Stack.Screen name="[id]/labs/[labResultId]" options={{ title: 'Exame' }} />
      <Stack.Screen name="[id]/prescriptions/index" options={{ title: 'Prescrições' }} />
      <Stack.Screen name="[id]/prescriptions/[prescriptionId]" options={{ title: 'Prescrição' }} />
      <Stack.Screen name="[id]/scores" options={{ title: 'Escores' }} />
      <Stack.Screen name="[id]/workouts/index" options={{ title: 'Treinos' }} />
      <Stack.Screen name="[id]/workouts/new" options={{ title: 'Novo plano' }} />
      <Stack.Screen name="[id]/assessments/index" options={{ title: 'Avaliações' }} />
      <Stack.Screen name="[id]/assessments/new" options={{ title: 'Nova avaliação' }} />
      <Stack.Screen name="[id]/assessments/[assessmentId]" options={{ title: 'Avaliação' }} />
      <Stack.Screen name="[id]/fitness-tests" options={{ title: 'Testes físicos' }} />
      <Stack.Screen name="[id]/posture" options={{ title: 'Avaliação postural' }} />
      <Stack.Screen name="[id]/periodization" options={{ title: 'Periodização' }} />
    </Stack>
  );
}
