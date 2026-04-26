import { useEffect, useState } from 'react';
import { Pressable, ScrollView, Switch, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { options, queryKeys } from '@plenya/api-client';
import { Card, FormField, Spinner, Text, useToast } from '@plenya/ui-mobile';

const HOURS = Array.from({ length: 24 }, (_, i) => i);
const MINUTES_STEP = [0, 15, 30, 45];

function pad(n: number) {
  return String(n).padStart(2, '0');
}

export default function NotificationsScreen() {
  const qc = useQueryClient();
  const toast = useToast();
  const prefs = useQuery(options.patientMeNotificationPreferencesOptions());

  const [appointmentReminder, setApptR] = useState(true);
  const [messageAlert, setMsgA] = useState(true);
  const [workoutReminder, setWorkoutR] = useState(true);
  const [hour, setHour] = useState(7);
  const [minute, setMinute] = useState(0);

  useEffect(() => {
    if (!prefs.data) return;
    setApptR(prefs.data.appointmentReminder);
    setMsgA(prefs.data.messageAlert);
    setWorkoutR(prefs.data.workoutReminder);
    const [h, m] = (prefs.data.workoutReminderTime || '07:00').split(':');
    setHour(Number(h) || 7);
    setMinute(Number(m) || 0);
  }, [prefs.data]);

  const patch = useMutation({
    mutationFn: (body: Parameters<typeof options.patientMeMutations.patchNotificationPreferences>[0]) =>
      options.patientMeMutations.patchNotificationPreferences(body),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: queryKeys.patientMe.notificationPreferences() });
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Erro', 'error'),
  });

  function update(partial: Parameters<typeof patch.mutate>[0]) {
    patch.mutate(partial);
  }

  if (prefs.isLoading) return <Spinner centered />;

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Card>
          <View className="flex-row items-center justify-between gap-3">
            <View className="flex-1">
              <Text variant="body" className="font-semibold">
                Lembrete de consulta
              </Text>
              <Text variant="caption">~1h antes do horário marcado.</Text>
            </View>
            <Switch
              value={appointmentReminder}
              onValueChange={(v) => {
                setApptR(v);
                update({ appointmentReminder: v });
              }}
            />
          </View>
        </Card>

        <Card>
          <View className="flex-row items-center justify-between gap-3">
            <View className="flex-1">
              <Text variant="body" className="font-semibold">
                Mensagens da clínica
              </Text>
              <Text variant="caption">Notificar quando receber resposta.</Text>
            </View>
            <Switch
              value={messageAlert}
              onValueChange={(v) => {
                setMsgA(v);
                update({ messageAlert: v });
              }}
            />
          </View>
        </Card>

        <Card>
          <View className="flex-row items-center justify-between gap-3">
            <View className="flex-1">
              <Text variant="body" className="font-semibold">
                Lembrete de treino
              </Text>
              <Text variant="caption">Toque diário no horário escolhido.</Text>
            </View>
            <Switch
              value={workoutReminder}
              onValueChange={(v) => {
                setWorkoutR(v);
                update({ workoutReminder: v });
              }}
            />
          </View>

          {workoutReminder && (
            <View className="mt-3">
              <FormField label={`Horário: ${pad(hour)}:${pad(minute)}`}>
                <View className="flex-row gap-2">
                  <View className="flex-1">
                    <Text variant="caption">Hora</Text>
                    <ScrollView
                      horizontal
                      showsHorizontalScrollIndicator={false}
                      contentContainerClassName="gap-1"
                    >
                      {HOURS.map((h) => (
                        <Pressable
                          key={h}
                          onPress={() => {
                            setHour(h);
                            update({ workoutReminderTime: `${pad(h)}:${pad(minute)}` });
                          }}
                          className={`min-w-10 items-center rounded-md border px-2 py-2 ${
                            h === hour
                              ? 'border-primary bg-primary'
                              : 'border-border bg-card'
                          }`}
                          accessibilityRole="button"
                          accessibilityLabel={`Selecionar ${pad(h)}h`}
                        >
                          <Text
                            className={
                              h === hour
                                ? 'text-primary-foreground font-semibold'
                                : ''
                            }
                          >
                            {pad(h)}
                          </Text>
                        </Pressable>
                      ))}
                    </ScrollView>
                  </View>
                </View>
                <View className="mt-2">
                  <Text variant="caption">Minutos</Text>
                  <View className="flex-row gap-2">
                    {MINUTES_STEP.map((m) => (
                      <Pressable
                        key={m}
                        onPress={() => {
                          setMinute(m);
                          update({ workoutReminderTime: `${pad(hour)}:${pad(m)}` });
                        }}
                        className={`flex-1 items-center rounded-md border py-2 ${
                          m === minute
                            ? 'border-primary bg-primary'
                            : 'border-border bg-card'
                        }`}
                        accessibilityRole="button"
                        accessibilityLabel={`Selecionar ${pad(m)} minutos`}
                      >
                        <Text
                          className={
                            m === minute
                              ? 'text-primary-foreground font-semibold'
                              : ''
                          }
                        >
                          {pad(m)}
                        </Text>
                      </Pressable>
                    ))}
                  </View>
                </View>
              </FormField>
            </View>
          )}
        </Card>

        {patch.isPending && (
          <Text variant="caption" className="text-center italic">
            Salvando…
          </Text>
        )}
      </ScrollView>
    </SafeAreaView>
  );
}
