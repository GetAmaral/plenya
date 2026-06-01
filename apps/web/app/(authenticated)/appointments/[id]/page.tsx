'use client';

/**
 * /appointments/[id]
 *
 * Detail view com:
 *  - Header (paciente + tipo + status badge)
 *  - Info (data/hora + duração + médico + motivo)
 *  - Ações: Confirmar / Cancelar / Reagendar
 *  - Daily.co embed (se type=telemedicine)
 *  - Badges de notificação enviada
 *
 * Polling 15s pra status; sala telemed é aberta on-demand via meeting_token
 * (HIGH H9: sala Daily.co é privacy=private; URL crua não funciona).
 */
import { useEffect, useMemo, useRef, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import {
  ArrowLeft,
  Calendar as CalendarIcon,
  CalendarClock,
  Clock,
  Stethoscope,
  Video,
  CheckCircle2,
  XCircle,
  Mail,
  MessageSquare,
  Loader2,
  AlertTriangle,
} from 'lucide-react';
import { ScheduleRecallDialog } from '@/components/recepcao/schedule-recall-dialog';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { Textarea } from '@/components/ui/textarea';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { PageHeader } from '@/components/layout/page-header';

import { DailyCoEmbed } from '@/components/calendar/daily-co-embed';
import { SlotPicker } from '@/components/calendar/slot-picker';
import {
  APPOINTMENT_STATUS_COLORS,
  APPOINTMENT_STATUS_LABELS,
  APPOINTMENT_TYPE_DEFAULT_DURATION,
  APPOINTMENT_TYPE_LABELS,
  useAppointment,
  useCancelAppointment,
  useConfirmAppointment,
  useTelemedToken,
  useUpdateAppointment,
  type CalendarSlot,
} from '@/lib/api/calendar-api';

export default function AppointmentDetailPage() {
  useRequireAuth();
  const router = useRouter();
  const params = useParams<{ id: string }>();
  const appointmentId = params?.id ?? '';

  const { data: appt, isLoading } = useAppointment(appointmentId);
  const confirmMutation = useConfirmAppointment(appointmentId);
  const cancelMutation = useCancelAppointment(appointmentId);
  const updateMutation = useUpdateAppointment(appointmentId);
  // HIGH H9 — sala Daily.co é privacy=private. Pedimos um meeting_token de
  // owner (médico = is_owner=true, screenshare=true) sob demanda. Token novo
  // a cada clique em "Abrir sala" — janela curta limita superfície de abuso.
  const telemedTokenMutation = useTelemedToken(appointmentId);
  const [joinUrl, setJoinUrl] = useState<string | null>(null);
  const [joinError, setJoinError] = useState<string | null>(null);

  const [showCancel, setShowCancel] = useState(false);
  const [showReschedule, setShowReschedule] = useState(false);
  const [showRecall, setShowRecall] = useState(false);
  const [showVideo, setShowVideo] = useState(true);

  const [cancelReason, setCancelReason] = useState('');

  // Reschedule modal state
  const [reschedDate, setReschedDate] = useState('');
  const [reschedSlot, setReschedSlot] = useState<CalendarSlot | null>(null);

  // Teleconsulta em janela separada — médico segue navegando no prontuário.
  const callWindowRef = useRef<Window | null>(null);
  const [callOpen, setCallOpen] = useState(false);
  const [popupBlocked, setPopupBlocked] = useState(false);

  // Reseta o painel quando o médico fecha a janela da teleconsulta.
  useEffect(() => {
    if (!callOpen) return;
    const t = setInterval(() => {
      if (!callWindowRef.current || callWindowRef.current.closed) {
        callWindowRef.current = null;
        setCallOpen(false);
      }
    }, 2000);
    return () => clearInterval(t);
  }, [callOpen]);

  // Abre a teleconsulta numa janela separada. O window.open é chamado JÁ no
  // gesto do clique (antes do await do token) — senão o bloqueador de popup
  // barra a janela. Se for bloqueado, cai pro embed inline.
  const openTelemedWindow = async () => {
    setJoinError(null);
    setPopupBlocked(false);
    const features =
      'width=1180,height=800,menubar=no,toolbar=no,location=no,status=no';
    const win = window.open('', 'plenya-teleconsulta', features);
    if (win) {
      win.document.write(
        '<!doctype html><title>Teleconsulta Plenya</title>' +
          '<body style="margin:0;font-family:Arial,Helvetica,sans-serif;display:flex;' +
          'align-items:center;justify-content:center;height:100vh;color:#555">' +
          'Conectando à teleconsulta…</body>',
      );
    }
    try {
      const res = await telemedTokenMutation.mutateAsync();
      setJoinUrl(res.joinUrl);
      if (win && !win.closed) {
        win.location.href = res.joinUrl;
        win.focus();
        callWindowRef.current = win;
        setCallOpen(true);
      } else {
        setPopupBlocked(true);
        setShowVideo(true);
      }
    } catch (err) {
      if (win && !win.closed) win.close();
      setJoinError(
        err instanceof Error ? err.message : 'Erro ao gerar acesso à sala',
      );
    }
  };

  const isActionable = useMemo(() => {
    if (!appt) return false;
    return appt.status === 'scheduled' || appt.status === 'confirmed';
  }, [appt]);

  const handleConfirm = async () => {
    try {
      await confirmMutation.mutateAsync();
      toast.success('Consulta confirmada');
    } catch (err) {
      toast.error('Erro ao confirmar', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  const handleCancel = async () => {
    if (!cancelReason.trim()) {
      toast.error('Informe o motivo do cancelamento');
      return;
    }
    try {
      await cancelMutation.mutateAsync({ reason: cancelReason.trim() });
      toast.success('Consulta cancelada');
      setShowCancel(false);
      setCancelReason('');
    } catch (err) {
      toast.error('Erro ao cancelar', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  const handleReschedule = async () => {
    if (!reschedSlot) {
      toast.error('Selecione um novo horário');
      return;
    }
    try {
      await updateMutation.mutateAsync({ scheduledAt: reschedSlot.startUtc });
      toast.success('Consulta reagendada');
      setShowReschedule(false);
      setReschedSlot(null);
      setReschedDate('');
    } catch (err) {
      const status = (err as { status?: number }).status;
      toast.error('Erro ao reagendar', {
        description:
          status === 409
            ? 'Esse horário já foi reservado para o médico. Escolha outro.'
            : err instanceof Error
              ? err.message
              : undefined,
      });
    }
  };

  if (isLoading || !appt) {
    return (
      <div className="container mx-auto space-y-4 py-8">
        <Skeleton className="h-12 w-64" />
        <Skeleton className="h-64 w-full" />
      </div>
    );
  }

  const start = new Date(appt.scheduledAt);
  const statusColor = APPOINTMENT_STATUS_COLORS[appt.status];

  return (
    <div className="container mx-auto space-y-6 py-8">
      <PageHeader
        breadcrumbs={[
          { label: 'Consultas', href: '/appointments' },
          { label: appt.patient?.name ?? 'Detalhes' },
        ]}
        title={appt.patient?.name ?? 'Consulta'}
        description={`${APPOINTMENT_TYPE_LABELS[appt.type]} — ${format(start, "dd 'de' MMM yyyy 'às' HH:mm", { locale: ptBR })}`}
      />

      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" onClick={() => router.push('/appointments')}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          Voltar
        </Button>
        <Badge variant="outline" className={statusColor}>
          {APPOINTMENT_STATUS_LABELS[appt.status]}
        </Badge>
        {appt.confirmationSentAt && (
          <Badge variant="outline" className="border-emerald-200 bg-emerald-50 text-emerald-900">
            <Mail className="mr-1 h-3 w-3" />
            Confirmação enviada
          </Badge>
        )}
        {appt.reminderSentAt && (
          <Badge variant="outline" className="border-amber-200 bg-amber-50 text-amber-900">
            <MessageSquare className="mr-1 h-3 w-3" />
            Lembrete enviado
          </Badge>
        )}
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Coluna principal */}
        <div className="space-y-6 lg:col-span-2">
          {/* Daily.co embed (telemedicine) — HIGH H9: sala privacy=private,
              precisa de meeting_token. Pedimos token sob demanda ao clicar. */}
          {appt.type === 'telemedicine' && (
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2 text-base">
                  <Video className="h-4 w-4" />
                  Sala de Teleconsulta
                </CardTitle>
              </CardHeader>
              <CardContent>
                {popupBlocked && joinUrl && showVideo ? (
                  // Fallback: navegador bloqueou a janela separada → embute inline.
                  <div className="space-y-2">
                    <p className="text-xs text-muted-foreground">
                      Seu navegador bloqueou a janela separada. Permita pop-ups deste
                      site para abrir a teleconsulta em janela própria e continuar
                      navegando pelo prontuário.
                    </p>
                    <DailyCoEmbed
                      roomURL={joinUrl}
                      onLeave={() => {
                        setShowVideo(false);
                        setJoinUrl(null);
                        setPopupBlocked(false);
                      }}
                    />
                  </div>
                ) : callOpen ? (
                  // Sala aberta em janela separada — EMR livre pra navegação.
                  <div className="space-y-3">
                    <p className="text-sm text-muted-foreground">
                      A teleconsulta está aberta em outra janela. Você pode continuar
                      navegando pelo prontuário aqui.
                    </p>
                    <div className="flex flex-wrap gap-2">
                      <Button
                        variant="outline"
                        onClick={() => callWindowRef.current?.focus()}
                      >
                        <Video className="mr-2 h-4 w-4" />
                        Trazer janela pra frente
                      </Button>
                      <Button
                        variant="ghost"
                        onClick={() => {
                          callWindowRef.current?.close();
                          callWindowRef.current = null;
                          setCallOpen(false);
                        }}
                      >
                        Encerrar
                      </Button>
                    </div>
                  </div>
                ) : (
                  <div className="space-y-3">
                    {joinError && (
                      <div className="flex items-start gap-2 rounded border border-destructive/30 bg-destructive/5 p-3 text-sm text-destructive">
                        <AlertTriangle className="mt-0.5 h-4 w-4 shrink-0" />
                        <span>{joinError}</span>
                      </div>
                    )}
                    <Button onClick={openTelemedWindow} disabled={telemedTokenMutation.isPending}>
                      {telemedTokenMutation.isPending ? (
                        <>
                          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                          Gerando acesso...
                        </>
                      ) : (
                        <>
                          <Video className="mr-2 h-4 w-4" />
                          Abrir sala em nova janela
                        </>
                      )}
                    </Button>
                    <p className="text-xs text-muted-foreground">
                      A sala abre em uma janela separada, para você continuar no
                      prontuário durante a consulta. Disponível de 3 horas antes a 3
                      horas depois do horário agendado.
                    </p>
                  </div>
                )}
              </CardContent>
            </Card>
          )}

          {/* Info */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Informações</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <InfoRow icon={CalendarIcon} label="Data/Hora">
                {format(start, "EEEE, dd 'de' MMM yyyy 'às' HH:mm", { locale: ptBR })}
              </InfoRow>
              <InfoRow icon={Clock} label="Duração">
                {appt.durationMinutes} min
              </InfoRow>
              <InfoRow icon={Stethoscope} label="Médico">
                {appt.doctor?.name ?? appt.doctorId}
              </InfoRow>
              <InfoRow icon={CalendarIcon} label="Motivo">
                {appt.reason}
              </InfoRow>
              {appt.patientNotes && (
                <InfoRow icon={MessageSquare} label="Observações do paciente">
                  {appt.patientNotes}
                </InfoRow>
              )}
              {appt.cancellationReason && (
                <InfoRow icon={AlertTriangle} label="Motivo do cancelamento">
                  {appt.cancellationReason}
                </InfoRow>
              )}
              {appt.externalCalendarEventId ? (
                <InfoRow icon={CalendarIcon} label="Google Event">
                  Sincronizado com Google Calendar
                </InfoRow>
              ) : null}
            </CardContent>
          </Card>
        </div>

        {/* Coluna lateral: ações */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Ações</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              {appt.status === 'scheduled' && isActionable && (
                <Button
                  className="w-full"
                  onClick={handleConfirm}
                  disabled={confirmMutation.isPending}
                >
                  {confirmMutation.isPending && (
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                  )}
                  <CheckCircle2 className="mr-2 h-4 w-4" />
                  Confirmar
                </Button>
              )}
              {isActionable && (
                <>
                  <Button
                    variant="outline"
                    className="w-full"
                    onClick={() => setShowReschedule(true)}
                  >
                    <CalendarIcon className="mr-2 h-4 w-4" />
                    Reagendar
                  </Button>
                  <Button
                    variant="outline"
                    className="w-full text-rose-600 hover:bg-rose-50 hover:text-rose-700"
                    onClick={() => setShowCancel(true)}
                  >
                    <XCircle className="mr-2 h-4 w-4" />
                    Cancelar
                  </Button>
                </>
              )}
              {!isActionable && (
                <p className="text-sm italic text-muted-foreground">
                  Sem ações de agenda para consulta {APPOINTMENT_STATUS_LABELS[appt.status].toLowerCase()}.
                </p>
              )}
              {appt.patient && (
                <Button
                  variant="outline"
                  className="w-full"
                  onClick={() => setShowRecall(true)}
                >
                  <CalendarClock className="mr-2 h-4 w-4" />
                  Agendar retorno
                </Button>
              )}
            </CardContent>
          </Card>

          {appt.patient && (
            <Card>
              <CardHeader>
                <CardTitle className="text-base">Paciente</CardTitle>
              </CardHeader>
              <CardContent className="space-y-2 text-sm">
                <p className="font-medium">{appt.patient.name}</p>
                {appt.patient.email && (
                  <p className="text-muted-foreground">{appt.patient.email}</p>
                )}
                {appt.patient.phone && (
                  <p className="text-muted-foreground">{appt.patient.phone}</p>
                )}
                <Button
                  variant="link"
                  className="h-auto p-0"
                  onClick={() => router.push(`/patients/${appt.patient!.id}`)}
                >
                  Abrir prontuário →
                </Button>
              </CardContent>
            </Card>
          )}
        </div>
      </div>

      {/* Cancel dialog */}
      <Dialog open={showCancel} onOpenChange={setShowCancel}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Cancelar consulta?</DialogTitle>
            <DialogDescription>
              Esta ação remove o evento do Google Calendar e a sala Daily.co. O paciente
              será notificado.
            </DialogDescription>
          </DialogHeader>
          <div className="space-y-2 py-2">
            <Label htmlFor="cancel-reason">Motivo</Label>
            <Textarea
              id="cancel-reason"
              placeholder="Ex: Imprevisto do paciente, conflito de agenda..."
              value={cancelReason}
              onChange={(e) => setCancelReason(e.target.value)}
              rows={3}
            />
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setShowCancel(false)}>
              Voltar
            </Button>
            <Button
              variant="destructive"
              onClick={handleCancel}
              disabled={cancelMutation.isPending}
            >
              {cancelMutation.isPending && (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              )}
              Cancelar consulta
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Reschedule dialog */}
      <Dialog open={showReschedule} onOpenChange={setShowReschedule}>
        <DialogContent className="sm:max-w-2xl">
          <DialogHeader>
            <DialogTitle>Reagendar consulta</DialogTitle>
            <DialogDescription>
              Selecione um novo horário disponível para o mesmo médico/tipo.
            </DialogDescription>
          </DialogHeader>
          <div className="space-y-4 py-2">
            <div className="space-y-2">
              <Label htmlFor="resched-date">Nova data</Label>
              <Input
                id="resched-date"
                type="date"
                min={format(new Date(), 'yyyy-MM-dd')}
                value={reschedDate}
                onChange={(e) => {
                  setReschedDate(e.target.value);
                  setReschedSlot(null);
                }}
              />
            </div>
            <SlotPicker
              doctorId={appt.doctorId}
              date={reschedDate || undefined}
              type={appt.type}
              overrideDuration={
                appt.durationMinutes !== APPOINTMENT_TYPE_DEFAULT_DURATION[appt.type]
                  ? appt.durationMinutes
                  : undefined
              }
              selectedSlotUtc={reschedSlot?.startUtc}
              onSelect={setReschedSlot}
            />
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setShowReschedule(false)}>
              Voltar
            </Button>
            <Button
              onClick={handleReschedule}
              disabled={updateMutation.isPending || !reschedSlot}
            >
              {updateMutation.isPending && (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              )}
              Confirmar reagendamento
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {appt.patient && (
        <ScheduleRecallDialog
          open={showRecall}
          onOpenChange={setShowRecall}
          patientId={appt.patient.id}
          doctorId={appt.doctorId}
          sourceAppointmentId={appointmentId}
          patientName={appt.patient.name}
        />
      )}
    </div>
  );
}

function InfoRow({
  icon: Icon,
  label,
  children,
}: {
  icon: React.ComponentType<{ className?: string }>;
  label: string;
  children: React.ReactNode;
}) {
  return (
    <div className="flex items-start gap-3">
      <Icon className="mt-0.5 h-4 w-4 shrink-0 text-muted-foreground" />
      <div className="flex-1">
        <p className="text-xs uppercase text-muted-foreground">{label}</p>
        <div className="text-sm">{children}</div>
      </div>
    </div>
  );
}
