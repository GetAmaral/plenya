import { Stack } from 'expo-router';

export default function PatientsLayout() {
  return (
    <Stack>
      <Stack.Screen name="index" options={{ title: 'Pacientes' }} />
      <Stack.Screen name="[id]" options={{ title: 'Paciente' }} />
    </Stack>
  );
}
