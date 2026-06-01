'use client';

/**
 * /configuracoes/agenda
 *
 * Editor de horário padrão + ausências do médico.
 * - Doctor: edita as próprias.
 * - Admin/Manager: dropdown pra escolher médico.
 *
 * Tabs: "Horário padrão" + "Ausências".
 */
import { useEffect, useState } from 'react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Skeleton } from '@/components/ui/skeleton';
import { Label } from '@/components/ui/label';
import { PageHeader } from '@/components/layout/page-header';

import { WorkingHoursEditor } from '@/components/calendar/working-hours-editor';
import { AbsencesList } from '@/components/calendar/absences-list';
import { useDoctors } from '@/lib/api/calendar-api';

export default function ConfigAgendaPage() {
  useRequireAuth();
  const { user } = useAuthStore();
  const { data: doctors, isLoading: loadingDoctors } = useDoctors();

  // admin/manager/secretary podem escolher e configurar a agenda de qualquer
  // médico; doctor configura a própria. A secretária é quem mais opera a agenda.
  const canSelectDoctor =
    isGranted(user, 'admin') ||
    isGranted(user, 'manager') ||
    isGranted(user, 'secretary');
  const isDoctor = isGranted(user, 'doctor');

  const [selectedDoctorId, setSelectedDoctorId] = useState<string | undefined>(undefined);

  // Se é doctor (e não admin), pré-seleciona ele mesmo.
  useEffect(() => {
    if (selectedDoctorId) return;
    if (isDoctor && user?.id) {
      setSelectedDoctorId(user.id);
    } else if (canSelectDoctor && doctors && doctors.length > 0) {
      setSelectedDoctorId(doctors[0]!.id);
    }
  }, [isDoctor, canSelectDoctor, doctors, user?.id, selectedDoctorId]);

  // ReadOnly: doctor olhando outro doctor (não permite por enquanto pelo backend, mas guardamos a UI).
  const readOnly =
    !canSelectDoctor && isDoctor && !!selectedDoctorId && selectedDoctorId !== user?.id;

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Configurações' }, { label: 'Agenda' }]}
        title="Agenda"
        description="Configure horários padrão de atendimento e períodos de ausência."
      />

      {/* Doctor picker (admin/manager/secretary) */}
      {canSelectDoctor && (
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Médico</CardTitle>
          </CardHeader>
          <CardContent>
            {loadingDoctors ? (
              <Skeleton className="h-10 w-64" />
            ) : (
              <div className="max-w-md space-y-2">
                <Label htmlFor="doctor-select">Selecione o médico</Label>
                <Select
                  value={selectedDoctorId}
                  onValueChange={setSelectedDoctorId}
                >
                  <SelectTrigger id="doctor-select">
                    <SelectValue placeholder="Escolher médico" />
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
            )}
          </CardContent>
        </Card>
      )}

      {selectedDoctorId ? (
        <Tabs defaultValue="working-hours" className="space-y-4">
          <TabsList>
            <TabsTrigger value="working-hours">Horário padrão</TabsTrigger>
            <TabsTrigger value="absences">Ausências</TabsTrigger>
          </TabsList>

          <TabsContent value="working-hours">
            <Card>
              <CardContent className="pt-6">
                <WorkingHoursEditor doctorId={selectedDoctorId} readOnly={readOnly} />
              </CardContent>
            </Card>
          </TabsContent>

          <TabsContent value="absences">
            <Card>
              <CardContent className="pt-6">
                <AbsencesList doctorId={selectedDoctorId} readOnly={readOnly} />
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
      ) : (
        <Card>
          <CardContent className="py-8 text-center text-sm text-muted-foreground">
            Selecione um médico para configurar a agenda.
          </CardContent>
        </Card>
      )}
    </div>
  );
}
