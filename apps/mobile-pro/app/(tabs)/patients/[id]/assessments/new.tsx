import { useState } from 'react';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, useLocalSearchParams } from 'expo-router';
import { useQueryClient } from '@tanstack/react-query';
import { options } from '@plenya/api-client';
import { physicalAssessmentsKeysFor } from '@plenya/api-client/options/physicalAssessments';
import {
  Button,
  FormField,
  Input,
  PhotoPicker,
  Text,
  useToast,
  type SelectedPhoto,
} from '@plenya/ui-mobile';
import { useScreenCaptureProtection } from '../../../../../lib/security/screenCapture';
import { useEnsureSelectedPatient } from '../../../../../features/patients/useEnsureSelectedPatient';
import { usePickAndUpload } from '../../../../../features/uploads/usePickAndUpload';

interface PhotoState extends SelectedPhoto {
  url: string;
  path: string;
}

function parseNumber(value: string): number | undefined {
  const cleaned = value.replace(',', '.').trim();
  if (!cleaned) return undefined;
  const n = Number(cleaned);
  return Number.isFinite(n) ? n : undefined;
}

export default function NewAssessmentScreen() {
  useScreenCaptureProtection();
  const { id } = useLocalSearchParams<{ id: string }>();
  const patientId = id ?? '';
  useEnsureSelectedPatient(patientId);

  const queryClient = useQueryClient();
  const toast = useToast();
  const pickAndUpload = usePickAndUpload();

  const [weight, setWeight] = useState('');
  const [height, setHeight] = useState('');
  const [systolic, setSystolic] = useState('');
  const [diastolic, setDiastolic] = useState('');
  const [notes, setNotes] = useState('');
  const [photos, setPhotos] = useState<PhotoState[]>([]);
  const [saving, setSaving] = useState(false);

  async function handleAddPhoto() {
    try {
      const result = await pickAndUpload();
      if (!result) return;
      setPhotos((prev) => [
        ...prev,
        { uri: result.localUri, url: result.url, path: result.path, mimeType: 'image/jpeg' },
      ]);
    } catch (err) {
      toast.show(err instanceof Error ? err.message : 'Falha no upload', 'error');
    }
  }

  function handleRemovePhoto(index: number) {
    setPhotos((prev) => prev.filter((_, i) => i !== index));
  }

  async function handleSave() {
    const w = parseNumber(weight);
    const h = parseNumber(height);
    if (!w && !h && photos.length === 0) {
      toast.show('Informe ao menos peso, altura ou uma foto', 'warning');
      return;
    }

    const bmi = w && h ? Number((w / Math.pow(h / 100, 2)).toFixed(1)) : undefined;

    setSaving(true);
    try {
      await options.physicalAssessmentMutations.create({
        weightKg: w,
        heightCm: h,
        bmi,
        bloodPressureSystolic: parseNumber(systolic),
        bloodPressureDiastolic: parseNumber(diastolic),
        notes: notes.trim() || undefined,
        photoUrls: photos.map((p) => p.url),
      });
      await queryClient.invalidateQueries({
        queryKey: physicalAssessmentsKeysFor.byPatient(patientId),
      });
      toast.show('Avaliação registrada', 'success');
      router.back();
    } catch (err) {
      toast.show(err instanceof Error ? err.message : 'Falha ao salvar', 'error');
    } finally {
      setSaving(false);
    }
  }

  const previewBmi = (() => {
    const w = parseNumber(weight);
    const h = parseNumber(height);
    if (!w || !h) return null;
    return (w / Math.pow(h / 100, 2)).toFixed(1);
  })();

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="heading">Nova avaliação</Text>

        <View className="flex-row gap-3">
          <View className="flex-1">
            <FormField label="Peso (kg)">
              <Input
                value={weight}
                onChangeText={setWeight}
                keyboardType="decimal-pad"
                placeholder="70.5"
                editable={!saving}
              />
            </FormField>
          </View>
          <View className="flex-1">
            <FormField label="Altura (cm)">
              <Input
                value={height}
                onChangeText={setHeight}
                keyboardType="decimal-pad"
                placeholder="172"
                editable={!saving}
              />
            </FormField>
          </View>
        </View>

        {previewBmi && (
          <Text variant="caption">IMC calculado: {previewBmi}</Text>
        )}

        <View className="flex-row gap-3">
          <View className="flex-1">
            <FormField label="PA sistólica">
              <Input
                value={systolic}
                onChangeText={setSystolic}
                keyboardType="number-pad"
                placeholder="120"
                editable={!saving}
              />
            </FormField>
          </View>
          <View className="flex-1">
            <FormField label="PA diastólica">
              <Input
                value={diastolic}
                onChangeText={setDiastolic}
                keyboardType="number-pad"
                placeholder="80"
                editable={!saving}
              />
            </FormField>
          </View>
        </View>

        <FormField label="Observações">
          <Input
            value={notes}
            onChangeText={setNotes}
            placeholder="Notas clínicas relevantes"
            multiline
            numberOfLines={4}
            editable={!saving}
          />
        </FormField>

        <PhotoPicker
          label="Fotos da avaliação"
          photos={photos}
          onAdd={handleAddPhoto}
          onRemove={handleRemovePhoto}
          max={6}
          disabled={saving}
        />

        <Button onPress={handleSave} loading={saving} fullWidth size="lg">
          Salvar avaliação
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
