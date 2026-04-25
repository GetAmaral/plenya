import { useState } from 'react';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { isValidCpf, stripCpf, formatCpf } from '@plenya/domain';

export default function NewPatientScreen() {
  const queryClient = useQueryClient();
  const toast = useToast();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [cpf, setCpf] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [saving, setSaving] = useState(false);

  async function handleSave() {
    if (!name.trim()) {
      toast.show('Nome é obrigatório', 'warning');
      return;
    }
    const cpfDigits = stripCpf(cpf);
    if (cpfDigits && !isValidCpf(cpfDigits)) {
      toast.show('CPF inválido', 'warning');
      return;
    }

    setSaving(true);
    try {
      const created = await options.patientMutations.create({
        name: name.trim(),
        email: email.trim() || undefined,
        phone: phone.trim() || undefined,
        cpfMasked: cpfDigits ? formatCpf(cpfDigits) : undefined,
        birthDate: birthDate.trim() || undefined,
      });
      await queryClient.invalidateQueries({ queryKey: queryKeys.patients.all() });
      toast.show('Paciente cadastrado', 'success');
      router.replace(`/(tabs)/patients/${created.id}` as never);
    } catch (err) {
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error');
    } finally {
      setSaving(false);
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="caption">
          Apenas dados básicos. Anamnese, exames e prescrições são adicionados depois pelo
          prontuário do paciente.
        </Text>

        <FormField label="Nome completo" required>
          <Input
            value={name}
            onChangeText={setName}
            placeholder="Ex: Maria Silva"
            autoCapitalize="words"
            editable={!saving}
          />
        </FormField>

        <View className="flex-row gap-3">
          <View className="flex-1">
            <FormField label="Email">
              <Input
                value={email}
                onChangeText={setEmail}
                placeholder="email@exemplo.com"
                autoCapitalize="none"
                keyboardType="email-address"
                editable={!saving}
              />
            </FormField>
          </View>
        </View>

        <FormField label="Telefone">
          <Input
            value={phone}
            onChangeText={setPhone}
            placeholder="(11) 98765-4321"
            keyboardType="phone-pad"
            editable={!saving}
          />
        </FormField>

        <FormField label="CPF" helper="Criptografado em repouso (LGPD)">
          <Input
            value={cpf}
            onChangeText={(v) => setCpf(stripCpf(v).slice(0, 11))}
            placeholder="00000000000"
            keyboardType="number-pad"
            maxLength={11}
            editable={!saving}
          />
        </FormField>

        <FormField label="Data de nascimento" helper="Formato: AAAA-MM-DD">
          <Input
            value={birthDate}
            onChangeText={setBirthDate}
            placeholder="1985-03-15"
            keyboardType="numbers-and-punctuation"
            editable={!saving}
            maxLength={10}
          />
        </FormField>

        <Button onPress={handleSave} loading={saving} fullWidth size="lg">
          Cadastrar paciente
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
