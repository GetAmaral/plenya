import { useEffect, useState } from 'react';
import { ScrollView, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, ErrorState, FormField, Input, Spinner, Text, useToast } from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';
import { useEnsureSelectedPatient } from '../../../../../features/patients/useEnsureSelectedPatient';

export default function EditAnamnesisScreen() {
  useScreenCaptureProtection();
  const { id, anamnesisId } = useLocalSearchParams<{ id: string; anamnesisId: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const queryClient = useQueryClient();
  const toast = useToast();

  const detail = useQuery(options.anamnesisDetailOptions(anamnesisId ?? ''));
  const [title, setTitle] = useState('');
  const [freeText, setFreeText] = useState('');
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (detail.data) {
      setTitle(detail.data.title ?? '');
      setFreeText(detail.data.freeText ?? '');
    }
  }, [detail.data]);

  if (detail.isLoading) return <Spinner centered />;
  if (detail.isError || !detail.data) {
    return (
      <SafeAreaView className="flex-1 bg-background">
        <ErrorState onRetry={() => detail.refetch()} />
      </SafeAreaView>
    );
  }

  async function handleSave() {
    if (!anamnesisId) return;
    setSaving(true);
    try {
      await options.anamnesisMutations.update(anamnesisId, {
        title: title.trim(),
        freeText: freeText.trim() || undefined,
      });
      await Promise.all([
        queryClient.invalidateQueries({ queryKey: queryKeys.anamnesis.detail(anamnesisId) }),
        queryClient.invalidateQueries({ queryKey: queryKeys.patients.anamnesis(patientId) }),
      ]);
      toast.show('Anamnese atualizada', 'success');
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
          <Input value={title} onChangeText={setTitle} editable={!saving} />
        </FormField>

        <FormField label="Anamnese (texto livre)">
          <View className="rounded-lg border border-border bg-background p-3">
            <TextInput
              value={freeText}
              onChangeText={setFreeText}
              multiline
              textAlignVertical="top"
              placeholderTextColor="#9CA3AF"
              editable={!saving}
              style={{ minHeight: 240, fontSize: 15, color: '#0f172a' }}
            />
          </View>
        </FormField>

        <Button onPress={handleSave} loading={saving} fullWidth size="lg">
          Salvar alterações
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
