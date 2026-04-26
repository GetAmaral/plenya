import { useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router, Stack } from 'expo-router';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Button, Card, FormField, Input, Text, useToast } from '@plenya/ui-mobile';

/**
 * Check-in diário — 5 sliders (1-5) + sono em horas + notas.
 * UX simples: linha de botões 1-5 (sem dependência extra de slider lib).
 * MVP cumpre o "10s" do plano; refina visual em iteração futura.
 */

function Scale({
  value,
  onChange,
  min = 1,
  max = 5,
  labels,
}: {
  value: number;
  onChange: (v: number) => void;
  min?: number;
  max?: number;
  labels?: [string, string];
}) {
  const items = Array.from({ length: max - min + 1 }, (_, i) => i + min);
  return (
    <View>
      <View className="flex-row gap-2">
        {items.map((n) => (
          <Pressable
            key={n}
            onPress={() => onChange(n)}
            className={`flex-1 items-center justify-center rounded-lg border py-3 ${
              value === n ? 'border-primary bg-primary' : 'border-border bg-card'
            }`}
            accessibilityRole="button"
            accessibilityLabel={`Selecionar ${n}`}
          >
            <Text
              className={`text-base font-semibold ${
                value === n ? 'text-primary-foreground' : 'text-foreground'
              }`}
            >
              {n}
            </Text>
          </Pressable>
        ))}
      </View>
      {labels && (
        <View className="mt-1 flex-row justify-between">
          <Text variant="caption">{labels[0]}</Text>
          <Text variant="caption">{labels[1]}</Text>
        </View>
      )}
    </View>
  );
}

export default function CheckInScreen() {
  const qc = useQueryClient();
  const toast = useToast();
  const [energy, setEnergy] = useState(3);
  const [pain, setPain] = useState(0);
  const [painLocation, setPainLocation] = useState('');
  const [mood, setMood] = useState(3);
  const [sleepHours, setSleepHours] = useState('7');
  const [sleepQuality, setSleepQuality] = useState(3);
  const [stress, setStress] = useState(3);
  const [notes, setNotes] = useState('');

  const create = useMutation({
    mutationFn: () =>
      options.patientMeMutations.createCheckIn({
        energy,
        pain,
        painLocation: pain > 0 && painLocation ? painLocation : undefined,
        mood,
        sleepHours: Number(sleepHours.replace(',', '.')) || 0,
        sleepQuality,
        stress,
        notes: notes || undefined,
      }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.checkIns() });
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.checkInToday() });
      toast.show('Check-in registrado!', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <Stack.Screen options={{ title: 'Como você está hoje?' }} />
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <FormField label="Energia">
            <Scale value={energy} onChange={setEnergy} labels={['Exausto', 'Ótimo']} />
          </FormField>
        </Card>

        <Card>
          <FormField label="Dor">
            <Scale value={pain} onChange={setPain} min={0} labels={['Nenhuma', 'Severa']} />
          </FormField>
          {pain > 0 && (
            <FormField label="Onde dói?">
              <Input
                value={painLocation}
                onChangeText={setPainLocation}
                placeholder="Ex: lombar, joelho direito…"
              />
            </FormField>
          )}
        </Card>

        <Card>
          <FormField label="Humor">
            <Scale value={mood} onChange={setMood} labels={['Péssimo', 'Ótimo']} />
          </FormField>
        </Card>

        <Card>
          <FormField label="Horas de sono">
            <Input
              value={sleepHours}
              onChangeText={setSleepHours}
              keyboardType="decimal-pad"
              placeholder="7"
            />
          </FormField>
          <FormField label="Qualidade do sono">
            <Scale value={sleepQuality} onChange={setSleepQuality} labels={['Ruim', 'Excelente']} />
          </FormField>
        </Card>

        <Card>
          <FormField label="Estresse">
            <Scale value={stress} onChange={setStress} labels={['Tranquilo', 'Muito alto']} />
          </FormField>
        </Card>

        <Card>
          <FormField label="Notas (opcional)">
            <Input
              value={notes}
              onChangeText={setNotes}
              multiline
              placeholder="Algo que queira registrar?"
            />
          </FormField>
        </Card>

        <Button
          onPress={() => create.mutate()}
          loading={create.isPending}
          fullWidth
          size="lg"
        >
          Registrar check-in
        </Button>
      </ScrollView>
    </SafeAreaView>
  );
}
