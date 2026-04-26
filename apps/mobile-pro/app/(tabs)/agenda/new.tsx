import { useEffect, useMemo, useState } from 'react';
import { Pressable, ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { router } from 'expo-router';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import {
  api,
  options,
  queryKeys,
  appointmentTypeLabels,
  appointmentTypeDefaultDuration,
  type AppointmentType,
  type PatientSummary,
  type StaffUser,
} from '@plenya/api-client';
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

const TYPES: AppointmentType[] = [
  'initial_assessment',
  'follow_up',
  'telemedicine',
  'procedure',
  'results_review',
];

function pad(n: number) {
  return String(n).padStart(2, '0');
}

function defaultDate(): string {
  const d = new Date();
  d.setDate(d.getDate() + 1);
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}`;
}

function isValidDate(s: string): boolean {
  return /^\d{4}-\d{2}-\d{2}$/.test(s) && !Number.isNaN(new Date(`${s}T12:00:00`).getTime());
}

function isValidTime(s: string): boolean {
  return /^\d{2}:\d{2}$/.test(s);
}

function buildIsoLocal(date: string, time: string): string {
  // Constrói ISO com offset local do device (RFC3339 aceito pelo backend)
  const [y, m, d] = date.split('-').map(Number);
  const [hh, mm] = time.split(':').map(Number);
  const dt = new Date(y, m - 1, d, hh, mm);
  const tzMin = -dt.getTimezoneOffset();
  const sign = tzMin >= 0 ? '+' : '-';
  const offH = pad(Math.floor(Math.abs(tzMin) / 60));
  const offM = pad(Math.abs(tzMin) % 60);
  return `${y}-${pad(m)}-${pad(d)}T${pad(hh)}:${pad(mm)}:00${sign}${offH}:${offM}`;
}

export default function NewAppointmentScreen() {
  const queryClient = useQueryClient();
  const toast = useToast();

  const me = useQuery(options.meOptions());
  const doctors = useQuery(options.doctorsListOptions());

  const [type, setType] = useState<AppointmentType>('follow_up');
  const [duration, setDuration] = useState<string>(
    String(appointmentTypeDefaultDuration.follow_up),
  );
  const [date, setDate] = useState(defaultDate());
  const [time, setTime] = useState('09:00');
  const [reason, setReason] = useState('');
  const [patientNotes, setPatientNotes] = useState('');

  const [patientPickerOpen, setPatientPickerOpen] = useState(false);
  const [patientSearch, setPatientSearch] = useState('');
  const [selectedPatient, setSelectedPatient] = useState<PatientSummary | null>(null);

  const [doctorPickerOpen, setDoctorPickerOpen] = useState(false);
  const [selectedDoctor, setSelectedDoctor] = useState<StaffUser | null>(null);

  const patientsQ = useQuery({
    queryKey: ['plenya', 'patients', 'list-search', patientSearch],
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (patientSearch) qs.set('search', patientSearch);
      qs.set('limit', '30');
      return api.get<{ items: PatientSummary[] }>(
        `/api/v1/patients?${qs.toString()}`,
        { signal },
      );
    },
    enabled: patientPickerOpen,
  });

  const create = useMutation({
    mutationFn: () => {
      if (!selectedPatient || !selectedDoctor) {
        throw new Error('Selecione paciente e profissional');
      }
      if (!isValidDate(date) || !isValidTime(time)) {
        throw new Error('Data ou hora inválida');
      }
      const dur = Number(duration);
      if (!Number.isInteger(dur) || dur < 15 || dur > 480) {
        throw new Error('Duração deve estar entre 15 e 480 min');
      }
      if (!reason.trim()) {
        throw new Error('Motivo obrigatório');
      }
      return options.appointmentMutations.create({
        patientId: selectedPatient.id,
        doctorId: selectedDoctor.id,
        scheduledAt: buildIsoLocal(date, time),
        durationMinutes: dur,
        type,
        reason: reason.trim(),
        patientNotes: patientNotes.trim() || undefined,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.appointments.all() });
      toast.show('Consulta agendada', 'success');
      router.back();
    },
    onError: (err) =>
      toast.show(err instanceof Error ? err.message : 'Falha ao agendar', 'error'),
  });

  // Auto-seleciona doctor quando o usuário logado for médico
  const meIsDoctor = useMemo(
    () => me.data?.role === 'doctor' || me.data?.role === 'nurse',
    [me.data?.role],
  );
  useEffect(() => {
    if (meIsDoctor && !selectedDoctor && me.data) {
      setSelectedDoctor({
        id: me.data.id,
        name: me.data.name,
        email: me.data.email,
        role: me.data.role,
      });
    }
  }, [meIsDoctor, me.data, selectedDoctor]);

  function selectType(t: AppointmentType) {
    setType(t);
    setDuration(String(appointmentTypeDefaultDuration[t]));
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['bottom']}>
      <ScrollView contentContainerClassName="gap-4 p-4">
        <Text variant="heading">Nova consulta</Text>

        <FormField label="Paciente" required>
          <Pressable
            onPress={() => setPatientPickerOpen(true)}
            className="rounded-lg border border-border bg-background p-3"
          >
            <Text variant="body">
              {selectedPatient?.name ?? 'Selecionar paciente...'}
            </Text>
          </Pressable>
        </FormField>

        <FormField label="Profissional" required>
          <Pressable
            onPress={() => setDoctorPickerOpen(true)}
            className="rounded-lg border border-border bg-background p-3"
          >
            <Text variant="body">
              {selectedDoctor?.name ?? 'Selecionar profissional...'}
            </Text>
          </Pressable>
        </FormField>

        <FormField label="Tipo">
          <View className="flex-row flex-wrap gap-2">
            {TYPES.map((t) => (
              <Pressable
                key={t}
                onPress={() => selectType(t)}
                className={`rounded-full px-3 py-1.5 ${
                  type === t ? 'bg-primary' : 'bg-muted'
                }`}
              >
                <Text
                  className={`text-xs font-semibold ${
                    type === t ? 'text-primary-foreground' : 'text-foreground'
                  }`}
                >
                  {appointmentTypeLabels[t]}
                </Text>
              </Pressable>
            ))}
          </View>
        </FormField>

        <View className="flex-row gap-3">
          <View className="flex-1">
            <FormField label="Data (AAAA-MM-DD)" required>
              <Input
                value={date}
                onChangeText={setDate}
                placeholder="2026-05-10"
                autoCapitalize="none"
                editable={!create.isPending}
              />
            </FormField>
          </View>
          <View className="flex-1">
            <FormField label="Hora (HH:MM)" required>
              <Input
                value={time}
                onChangeText={setTime}
                placeholder="14:00"
                keyboardType="number-pad"
                editable={!create.isPending}
              />
            </FormField>
          </View>
        </View>

        <FormField label="Duração (min)">
          <Input
            value={duration}
            onChangeText={setDuration}
            keyboardType="number-pad"
            editable={!create.isPending}
          />
        </FormField>

        <FormField label="Motivo" required>
          <Input
            value={reason}
            onChangeText={setReason}
            placeholder="Ex: Retorno avaliação cardiológica"
            multiline
            numberOfLines={2}
            editable={!create.isPending}
          />
        </FormField>

        <FormField label="Notas do paciente">
          <Input
            value={patientNotes}
            onChangeText={setPatientNotes}
            placeholder="Observações que o paciente trouxe"
            multiline
            numberOfLines={3}
            editable={!create.isPending}
          />
        </FormField>

        <Button
          onPress={() => create.mutate()}
          loading={create.isPending}
          fullWidth
          size="lg"
        >
          Agendar consulta
        </Button>
      </ScrollView>

      <Sheet open={patientPickerOpen} onClose={() => setPatientPickerOpen(false)}>
        <Text variant="title" className="mb-2">
          Escolher paciente
        </Text>
        <Input
          value={patientSearch}
          onChangeText={setPatientSearch}
          placeholder="Buscar paciente..."
          autoCapitalize="none"
        />
        <View className="mt-3 max-h-96">
          {patientsQ.isLoading ? (
            <Spinner />
          ) : (
            <View className="gap-1">
              {(patientsQ.data?.items ?? []).slice(0, 30).map((p) => (
                  <Pressable
                    key={p.id}
                    onPress={() => {
                      setSelectedPatient(p);
                      setPatientPickerOpen(false);
                    }}
                    className="rounded-md bg-muted px-3 py-2"
                  >
                    <Text variant="body">{p.name}</Text>
                    <Text variant="caption">{p.email ?? p.phone ?? '—'}</Text>
                  </Pressable>
                ))}
            </View>
          )}
        </View>
      </Sheet>

      <Sheet open={doctorPickerOpen} onClose={() => setDoctorPickerOpen(false)}>
        <Text variant="title" className="mb-2">
          Escolher profissional
        </Text>
        {doctors.isLoading ? (
          <Spinner />
        ) : (
          <View className="gap-1">
            {(doctors.data ?? []).map((d) => (
              <Pressable
                key={d.id}
                onPress={() => {
                  setSelectedDoctor(d);
                  setDoctorPickerOpen(false);
                }}
                className="rounded-md bg-muted px-3 py-2"
              >
                <Card>
                  <Text variant="body">{d.name}</Text>
                  <Text variant="caption">{d.email}</Text>
                </Card>
              </Pressable>
            ))}
          </View>
        )}
      </Sheet>
    </SafeAreaView>
  );
}
