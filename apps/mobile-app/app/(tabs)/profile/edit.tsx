import { useEffect, useState } from 'react';
import { KeyboardAvoidingView, Platform, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, FormField, Input, Spinner, Text, useToast } from '@plenya/ui-mobile';

export default function EditProfileScreen() {
  const qc = useQueryClient();
  const toast = useToast();
  const profile = useQuery(options.patientMeProfileOptions());
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [address, setAddress] = useState('');
  const [emergencyPhone, setEmergencyPhone] = useState('');

  useEffect(() => {
    if (!profile.data) return;
    setPhone(profile.data.phone ?? '');
    setEmail(profile.data.email ?? '');
  }, [profile.data]);

  const save = useMutation({
    mutationFn: () =>
      options.patientMeMutations.updateProfile({
        phone: phone || undefined,
        email: email || undefined,
        address: address || undefined,
        emergencyPhone: emergencyPhone || undefined,
      }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.profile() });
      toast.show('Dados atualizados', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  if (profile.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : undefined}
        className="flex-1"
      >
        <ScrollView contentContainerClassName="gap-4 p-4">
          <Text variant="caption">
            Dados de identificação (nome, CPF, data de nascimento) só podem ser
            alterados pela sua clínica.
          </Text>

          <FormField label="Telefone">
            <Input
              value={phone}
              onChangeText={setPhone}
              keyboardType="phone-pad"
              placeholder="(11) 99999-9999"
            />
          </FormField>
          <FormField label="Email">
            <Input
              value={email}
              onChangeText={setEmail}
              keyboardType="email-address"
              autoCapitalize="none"
              placeholder="seu@email.com.br"
            />
          </FormField>
          <FormField label="Endereço">
            <Input
              value={address}
              onChangeText={setAddress}
              placeholder="Rua, número, bairro, cidade"
            />
          </FormField>
          <FormField label="Contato de emergência">
            <Input
              value={emergencyPhone}
              onChangeText={setEmergencyPhone}
              keyboardType="phone-pad"
              placeholder="(11) 99999-9999"
            />
          </FormField>

          <View className="gap-2 mt-2">
            <Button
              onPress={() => save.mutate()}
              loading={save.isPending}
              fullWidth
              size="lg"
            >
              Salvar
            </Button>
            <Button onPress={() => router.back()} variant="ghost" fullWidth>
              Cancelar
            </Button>
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}
