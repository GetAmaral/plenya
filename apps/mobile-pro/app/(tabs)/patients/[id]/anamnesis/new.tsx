import { useState } from 'react';
import { Pressable, ScrollView, TextInput, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { api, options, queryKeys, type AnamnesisTemplate } from '@plenya/api-client';
import {
  Button,
  Card,
  FormField,
  Input,
  Sheet,
  Spinner,
  Text,
  useToast,
} from '@plenya/ui-mobile';
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
  const [templatesOpen, setTemplatesOpen] = useState(false);

  const templates = useQuery(options.anamnesisTemplatesOptions());

  const loadTemplate = useMutation({
    mutationFn: (templateId: string) =>
      api.get<AnamnesisTemplate>(`/api/v1/anamnesis-templates/${templateId}`),
    onSuccess: (tpl) => {
      const items = tpl.items ?? [];
      const draft = items
        .slice()
        .sort((a, b) => a.order - b.order)
        .map((it) => `${it.prompt}\n— `)
        .join('\n\n');
      if (!title.trim()) setTitle(tpl.name);
      setFreeText((prev) => (prev ? `${prev}\n\n${draft}` : draft));
      setTemplatesOpen(false);
      toast.show(`Template "${tpl.name}" aplicado`, 'success');
    },
    onError: () => toast.show('Falha ao carregar template', 'error'),
  });

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
        <Button
          variant="outline"
          size="sm"
          onPress={() => setTemplatesOpen(true)}
          fullWidth
        >
          Aplicar template
        </Button>

        <FormField label="Título" required>
          <Input
            testID="title"
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
              testID="freeText"
              value={freeText}
              onChangeText={setFreeText}
              multiline
              numberOfLines={14}
              textAlignVertical="top"
              placeholder="Comece a escrever..."
              placeholderTextColor="#9CA3AF"
              editable={!saving}
              style={{ minHeight: 260, fontSize: 15, color: '#0f172a' }}
            />
          </View>
        </FormField>

        <Button onPress={handleSave} loading={saving} fullWidth size="lg">
          Salvar anamnese
        </Button>
      </ScrollView>

      <Sheet open={templatesOpen} onClose={() => setTemplatesOpen(false)}>
        <Text variant="title" className="mb-2">
          Escolher template
        </Text>
        <Text variant="caption" className="mb-3">
          O conteúdo é inserido como rascunho — você edita à vontade.
        </Text>

        {templates.isLoading ? (
          <Spinner />
        ) : (
          <View className="gap-2">
            {(templates.data ?? []).map((tpl) => (
              <Pressable key={tpl.id} onPress={() => loadTemplate.mutate(tpl.id)}>
                <Card>
                  <Text variant="body">{tpl.name}</Text>
                  {tpl.description && (
                    <Text variant="caption">{tpl.description}</Text>
                  )}
                </Card>
              </Pressable>
            ))}
            {(templates.data ?? []).length === 0 && (
              <Text variant="caption">Nenhum template cadastrado.</Text>
            )}
          </View>
        )}
      </Sheet>
    </SafeAreaView>
  );
}
