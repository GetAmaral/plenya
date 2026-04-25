import { Stack } from 'expo-router';

export default function PatientsLayout() {
  return (
    <Stack screenOptions={{ headerBackTitle: 'Voltar' }}>
      <Stack.Screen name="index" options={{ title: 'Pacientes' }} />
      <Stack.Screen name="[id]/index" options={{ title: 'Paciente' }} />
      <Stack.Screen name="[id]/anamnesis" options={{ title: 'Anamnese' }} />
      <Stack.Screen name="[id]/labs" options={{ title: 'Exames' }} />
      <Stack.Screen name="[id]/prescriptions" options={{ title: 'Prescrições' }} />
      <Stack.Screen name="[id]/scores" options={{ title: 'Escores' }} />
    </Stack>
  );
}
