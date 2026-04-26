import { useEffect, useState } from 'react';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import {
  Button,
  ErrorState,
  FormField,
  Input,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../lib/security/screenCapture';
import { useEnsureSelectedPatient } from '../../../../features/patients/useEnsureSelectedPatient';

export default function EditPatientScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const queryClient = useQueryClient();
  const toast = useToast();
  const patient = useQuery(options.patientOptions(patientId));

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [address, setAddress] = useState('');

  useEffect(() => {
    if (patient.data) {
      setName(patient.data.name ?? '');
      setEmail(patient.data.email ?? '');
      setPhone(patient.data.phone ?? '');
      // addressJson é um campo legado serializado; preferimos string simples editável
      setAddress(patient.data.addressJson ?? '');
    }
  }, [patient.data]);

  const update = useMutation({
    mutationFn: () =>
      options.patientMutations.update(patientId, {
        name: name.trim() || undefined,
        email: email.trim() || undefined,
        phone: phone.trim() || undefined,
        // address API field — só inclui se mudou (evita sobrescrever JSON estruturado)
        ...(address.trim() && address.trim() !== (patient.data?.addressJson ?? '')
          ? { addressJson: address.trim() }
          : {}),
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.patients.detail(patientId) });
      queryClient.invalidateQueries({ queryKey: queryKeys.patients.all() });
      toast.show('Paciente atualizado', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error'),
  });

  if (patient.isLoading) return <Spinner centered />;
  if (patient.isError) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => patient.refetch()} />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="heading">Editar paciente</Text>
        <Text variant="caption">
          Atualizar nome, contato e endereço. CPF/RG não editáveis no app.
        </Text>

        <FormField label="Nome">
          <Input
            value={name}
            onChangeText={setName}
            autoCapitalize="words"
            editable={!update.isPending}
          />
        </FormField>

        <FormField label="Email">
          <Input
            value={email}
            onChangeText={setEmail}
            autoCapitalize="none"
            keyboardType="email-address"
            editable={!update.isPending}
          />
        </FormField>

        <FormField label="Telefone">
          <Input
            value={phone}
            onChangeText={setPhone}
            keyboardType="phone-pad"
            editable={!update.isPending}
          />
        </FormField>

        <FormField label="Endereço">
          <Input
            value={address}
            onChangeText={setAddress}
            multiline
            numberOfLines={3}
            editable={!update.isPending}
          />
        </FormField>

        <Button
          onPress={() => update.mutate()}
          loading={update.isPending}
          fullWidth
          size="lg"
        >
          Salvar alterações
        </Button>

        <Button
          variant="ghost"
          onPress={() => router.back()}
          disabled={update.isPending}
          fullWidth
        >
          Cancelar
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
