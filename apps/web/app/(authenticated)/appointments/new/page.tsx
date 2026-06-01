'use client';

/**
 * /appointments/new
 *
 * Form pra criar consulta:
 *  1. Patient (busca/select)
 *  2. Doctor (default = self se doctor; admin/secretary escolhe)
 *  3. Type + duração
 *  4. Date picker
 *  5. SlotPicker (livre = clicável)
 *  6. Reason + observações
 *
 * Submit → POST + redirect /appointments/:id (sync Google + Daily acontece async).
 */
import { useEffect, useMemo, useState } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { format } from 'date-fns';
import { Loader2, Calendar as CalendarIcon } from 'lucide-react';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { useSelectedPatient } from '@/lib/use-selected-patient';
import { usePatients } from '@/lib/api/patient-api';
import {
  APPOINTMENT_TYPE_DEFAULT_DURATION,
  APPOINTMENT_TYPE_LABELS,
  useCreateAppointment,
  useDoctors,
  type AppointmentType,
  type CalendarSlot,
} from '@/lib/api/calendar-api';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { PageHeader } from '@/components/layout/page-header';
import { SlotPicker } from '@/components/calendar/slot-picker';

const TODAY_YMD = format(new Date(), 'yyyy-MM-dd');

export default function NewAppointmentPage() {
  useRequireAuth();
  const router = useRouter();
  const search = useSearchParams();
  const { user } = useAuthStore();

  // Continuum (Fase 3): query params pra ancorar a consulta a um marco do programa.
  const { selectedPatient, selectedPatientId } = useSelectedPatient();
  const prefilledPatientId = search?.get('patientId') ?? '';
  const continuumItemId = search?.get('continuumItemId') ?? '';
  const prefilledSpecialty = search?.get('specialty') ?? ''; // doctor|nutritionist|psychologist|physicalEducator
  const prefilledDate = search?.get('date') ?? ''; // YYYY-MM-DD vindo do slot-click do calendário

  const isDoctor = isGranted(user, 'doctor');
  const isStaff =
    isGranted(user, 'admin') || isGranted(user, 'secretary') || isGranted(user, 'manager');

  const { data: patients, isLoading: loadingPatients } = usePatients();
  const { data: doctors } = useDoctors();
  const createMutation = useCreateAppointment();

  const [patientId, setPatientId] = useState<string>(
    prefilledPatientId || selectedPatientId || '',
  );
  const [patientSearch, setPatientSearch] = useState('');
  const [doctorId, setDoctorId] = useState<string>('');
  const [type, setType] = useState<AppointmentType>('initial_assessment');
  const [presetKey, setPresetKey] = useState<string>('initial_assessment');
  const [date, setDate] = useState<string>(prefilledDate || TODAY_YMD);
  const [duration, setDuration] = useState<number>(
    APPOINTMENT_TYPE_DEFAULT_DURATION.initial_assessment,
  );
  const [slot, setSlot] = useState<CalendarSlot | null>(null);
  const [reason, setReason] = useState('');
  const [patientNotes, setPatientNotes] = useState('');

  // Pré-seleciona doctor = self se for doctor.
  useEffect(() => {
    if (!doctorId && isDoctor && user?.id) {
      setDoctorId(user.id);
    } else if (!doctorId && isStaff && doctors && doctors.length > 0) {
      setDoctorId(doctors[0]!.id);
    }
  }, [doctorId, isDoctor, isStaff, doctors, user?.id]);

  // Pré-seleciona o paciente do contexto (selectedPatient) como padrão — mesmo
  // padrão dos demais itens de prontuário. Só preenche se nada veio na URL e o
  // usuário ainda não escolheu (não sobrescreve troca manual). Cobre o caso do
  // store rehidratar depois do primeiro render.
  useEffect(() => {
    if (!patientId && selectedPatientId) {
      setPatientId(selectedPatientId);
    }
  }, [patientId, selectedPatientId]);

  // Opções de consulta = preset de tipo + duração. Teleconsulta tem dois presets
  // (30 e 60 min): ambos gravam type=telemedicine, mudando só a duração — sem
  // criar tipo novo no backend.
  const CONSULTATION_OPTIONS: {
    key: string;
    type: AppointmentType;
    duration: number;
    label: string;
  }[] = [
    { key: 'initial_assessment', type: 'initial_assessment', duration: APPOINTMENT_TYPE_DEFAULT_DURATION.initial_assessment, label: APPOINTMENT_TYPE_LABELS.initial_assessment },
    { key: 'follow_up', type: 'follow_up', duration: APPOINTMENT_TYPE_DEFAULT_DURATION.follow_up, label: APPOINTMENT_TYPE_LABELS.follow_up },
    { key: 'telemedicine', type: 'telemedicine', duration: APPOINTMENT_TYPE_DEFAULT_DURATION.telemedicine, label: APPOINTMENT_TYPE_LABELS.telemedicine },
    { key: 'telemedicine_60', type: 'telemedicine', duration: 60, label: APPOINTMENT_TYPE_LABELS.telemedicine },
    { key: 'procedure', type: 'procedure', duration: APPOINTMENT_TYPE_DEFAULT_DURATION.procedure, label: APPOINTMENT_TYPE_LABELS.procedure },
    { key: 'results_review', type: 'results_review', duration: APPOINTMENT_TYPE_DEFAULT_DURATION.results_review, label: APPOINTMENT_TYPE_LABELS.results_review },
  ];

  // Troca de preset: define tipo + duração default e reseta slot.
  const handlePresetChange = (key: string) => {
    const opt = CONSULTATION_OPTIONS.find((o) => o.key === key);
    if (!opt) return;
    setPresetKey(key);
    setType(opt.type);
    setDuration(opt.duration);
    setSlot(null);
  };

  // Quando date muda, reset slot.
  const handleDateChange = (newDate: string) => {
    setDate(newDate);
    setSlot(null);
  };

  // Filtragem client-side dos pacientes
  const filteredPatients = useMemo(() => {
    const list = patients ?? [];
    const base = !patientSearch.trim()
      ? list.slice(0, 50)
      : list
          .filter(
            (p) =>
              p.name.toLowerCase().includes(patientSearch.toLowerCase()) ||
              (p.email ?? '').toLowerCase().includes(patientSearch.toLowerCase()) ||
              (p.cpf ?? '').includes(patientSearch),
          )
          .slice(0, 50);

    // Garante que o paciente já escolhido apareça na lista (senão o trigger do
    // Select fica em branco quando ele está fora da fatia/filtro).
    if (patientId && !base.some((p) => p.id === patientId)) {
      const chosen =
        list.find((p) => p.id === patientId) ??
        (selectedPatient?.id === patientId ? selectedPatient : null);
      if (chosen) return [chosen, ...base];
    }
    return base;
  }, [patients, patientSearch, patientId, selectedPatient]);

  const canSubmit =
    !!patientId && !!doctorId && !!slot && reason.trim().length > 0 && !createMutation.isPending;

  const handleSubmit = async () => {
    if (!canSubmit || !slot) return;
    try {
      const result = await createMutation.mutateAsync({
        patientId,
        doctorId,
        scheduledAt: slot.startUtc,
        durationMinutes: duration,
        type,
        reason: reason.trim(),
        patientNotes: patientNotes.trim() || undefined,
        continuumItemId: continuumItemId || undefined,
      });
      toast.success('Consulta agendada!', {
        description: 'Sincronização com Google e Daily.co ocorre em alguns segundos.',
      });
      router.push(`/appointments/${result.id}`);
    } catch (err) {
      toast.error('Erro ao agendar', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Consultas', href: '/appointments' }, { label: 'Nova' }]}
        title="Nova consulta"
        description={
          continuumItemId
            ? 'Ancorando consulta a marco do Continuum do paciente.'
            : 'Selecione paciente, médico e horário disponível.'
        }
      />
      {continuumItemId && (
        <div className="rounded-md border border-blue-200 bg-blue-50 px-3 py-2 text-xs text-blue-800">
          Esta consulta vai ancorar um marco do Continuum. Quando ela for marcada como
          "completada", o item da timeline atualiza automaticamente.
        </div>
      )}

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Coluna esquerda: form */}
        <div className="space-y-6 lg:col-span-2">
          {/* Patient + Doctor */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Paciente e médico</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="patient-search">Paciente</Label>
                <Input
                  id="patient-search"
                  placeholder="Buscar por nome, email ou CPF..."
                  value={patientSearch}
                  onChange={(e) => setPatientSearch(e.target.value)}
                />
                {loadingPatients ? (
                  <p className="text-sm italic text-muted-foreground">Carregando pacientes...</p>
                ) : (
                  <Select value={patientId} onValueChange={setPatientId}>
                    <SelectTrigger>
                      <SelectValue placeholder="Selecione paciente" />
                    </SelectTrigger>
                    <SelectContent>
                      {filteredPatients.length === 0 ? (
                        <div className="p-2 text-sm text-muted-foreground">
                          Nenhum paciente encontrado.
                        </div>
                      ) : (
                        filteredPatients.map((p) => (
                          <SelectItem key={p.id} value={p.id}>
                            {p.name}
                            {p.email ? ` — ${p.email}` : ''}
                          </SelectItem>
                        ))
                      )}
                    </SelectContent>
                  </Select>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="doctor-select">Médico</Label>
                <Select value={doctorId} onValueChange={setDoctorId}>
                  <SelectTrigger id="doctor-select">
                    <SelectValue placeholder="Selecione médico" />
                  </SelectTrigger>
                  <SelectContent>
                    {(doctors ?? []).map((d) => (
                      <SelectItem key={d.id} value={d.id}>
                        {d.name}
                        {d.specialty ? ` — ${d.specialty}` : ''}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </CardContent>
          </Card>

          {/* Type + Date */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Tipo e data</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid gap-4 sm:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="type-select">Tipo de consulta</Label>
                  <Select value={presetKey} onValueChange={handlePresetChange}>
                    <SelectTrigger id="type-select">
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      {CONSULTATION_OPTIONS.map((o) => (
                        <SelectItem key={o.key} value={o.key}>
                          {o.label} ({o.duration} min)
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="duration-input">Duração (min)</Label>
                  <Input
                    id="duration-input"
                    type="number"
                    min={15}
                    max={480}
                    step={5}
                    value={duration}
                    onChange={(e) => {
                      setDuration(Number(e.target.value));
                      setSlot(null);
                    }}
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="date-input">Data</Label>
                <Input
                  id="date-input"
                  type="date"
                  min={TODAY_YMD}
                  value={date}
                  onChange={(e) => handleDateChange(e.target.value)}
                />
              </div>
            </CardContent>
          </Card>

          {/* Slot picker */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <CalendarIcon className="h-4 w-4" />
                Horários disponíveis
              </CardTitle>
            </CardHeader>
            <CardContent>
              <SlotPicker
                doctorId={doctorId || undefined}
                date={date || undefined}
                type={type}
                overrideDuration={duration !== APPOINTMENT_TYPE_DEFAULT_DURATION[type] ? duration : undefined}
                selectedSlotUtc={slot?.startUtc}
                onSelect={(s) => setSlot(s)}
              />
            </CardContent>
          </Card>

          {/* Reason + notes */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Motivo e observações</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="reason">Motivo *</Label>
                <Textarea
                  id="reason"
                  placeholder="Ex: Consulta de rotina, retorno pós-cirurgia..."
                  value={reason}
                  onChange={(e) => setReason(e.target.value)}
                  rows={3}
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="patient-notes">Observações do paciente</Label>
                <Textarea
                  id="patient-notes"
                  placeholder="Informações relevantes informadas pelo paciente..."
                  value={patientNotes}
                  onChange={(e) => setPatientNotes(e.target.value)}
                  rows={3}
                />
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Coluna direita: resumo + submit */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Resumo</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              <div>
                <p className="text-xs text-muted-foreground">Tipo</p>
                <p className="font-medium">{APPOINTMENT_TYPE_LABELS[type]}</p>
              </div>
              <div>
                <p className="text-xs text-muted-foreground">Duração</p>
                <p className="font-medium">{duration} min</p>
              </div>
              <div>
                <p className="text-xs text-muted-foreground">Horário</p>
                <p className="font-medium">
                  {slot
                    ? format(new Date(slot.startUtc), 'dd/MM/yyyy HH:mm')
                    : 'Selecionar slot'}
                </p>
              </div>
              <Button
                className="w-full"
                disabled={!canSubmit}
                onClick={handleSubmit}
              >
                {createMutation.isPending && (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                )}
                Agendar
              </Button>
              {type === 'telemedicine' && slot && (
                <p className="text-xs text-muted-foreground">
                  Sala Daily.co será criada automaticamente.
                </p>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
