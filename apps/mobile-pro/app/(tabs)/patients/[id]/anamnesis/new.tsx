import { useState } from 'react';
import { ScrollView, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, FormField, Input, Text, useToast } from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';
import { useEnsureSelectedPatient } from '../../../../../features/patients/useEnsureSelectedPatient';

export default function NewAnamnesisScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const queryClient = useQueryClient();
  const toast = useToast();
  const [title, setTitle] = useState('');
  const [freeText, setFreeText] = useState('');
  const [saving, setSaving] = useState(false);

  async function handleSave() {
    if (!title.trim()) {
      toast.show('Título obrigatório', 'warning');
      return;
    }
    setSaving(true);
    try {
      await options.anamnesisMutations.create({
        patientId,
        title: title.trim(),
        freeText: freeText.trim() || undefined,
      });
      await queryClient.invalidateQueries({
        queryKey: queryKeys.patients.anamnesis(patientId),
      });
      toast.show('Anamnese registrada', 'success');
      router.back();
    } catch (err) {
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error');
    } finally {
      setSaving(false);
    }
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <FormField label="Título" required>
          <Input
            value={title}
            onChangeText={setTitle}
            placeholder="Ex: Primeira consulta"
            editable={!saving}
            autoCapitalize="sentences"
          />
        </FormField>

        <FormField label="Anamnese (texto livre)" helper="História clínica, queixas, observações">
          <View className="rounded-lg border border-border bg-background p-3">
            <TextInput
              value={freeText}
              onChangeText={setFreeText}
              multiline
              numberOfLines={12}
              textAlignVertical="top"
              placeholder="Comece a escrever..."
              placeholderTextColor="#9CA3AF"
              editable={!saving}
              style={{ minHeight: 220, fontSize: 15, color: '#0f172a' }}
            />
          </View>
        </FormField>

        <Button onPress={handleSave} loading={saving} fullWidth size="lg">
          Salvar anamnese
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
