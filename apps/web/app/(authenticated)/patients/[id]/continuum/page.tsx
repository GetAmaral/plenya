'use client';

/**
 * /patients/[id]/continuum — timeline da inscrição Continuum do paciente.
 *
 * - Se não há inscrição ativa: botão "Iniciar Continuum" abre dialog com
 *   seletor de template + startDate + coordenador médico.
 * - Se há inscrição: mostra header (template, semana atual, progresso) + items
 *   agrupados por semana com status badges. Filtro por especialidade/tipo.
 *
 * Plano integrado e logística de box ficam pra Fases 4 e 5.
 */
import { useMemo, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { addDays, differenceInWeeks, format } from 'date-fns';
import { formatDate } from '@/lib/format-date';
import { toast } from 'sonner';
import {
  ArrowLeft,
  Loader2,
  Plus,
  Calendar as CalendarIcon,
  Package,
  Workflow,
  Target,
  CheckCircle2,
  AlertTriangle,
} from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { PageHeader } from '@/components/layout/page-header';
import {
  usePatientContinuums,
  useContinuumEnrollment,
  useEnrollPatientContinuum,
  useUpdateContinuumItem,
  useContinuumTemplates,
  ITEM_STATUS_LABELS,
  ITEM_STATUS_COLORS,
  ITEM_TYPE_LABELS,
  SPECIALTY_LABELS,
  type PatientContinuumItem,
  type ContinuumItemStatus,
} from '@/lib/api/continuum-api';
import { useDoctors } from '@/lib/api/calendar-api';
import { cn } from '@/lib/utils';
import { IntegratedPlanEditor } from '@/components/continuum/integrated-plan-editor';
import { BoxCard } from '@/components/continuum/box-card';
import { useContinuumBoxes } from '@/lib/api/continuum-api';

const TYPE_ICON = {
  appointment: CalendarIcon,
  box: Package,
  reassessment: Target,
  milestone: CheckCircle2,
  custom: Workflow,
} as const;

export default function PatientContinuumPage() {
  const params = useParams<{ id: string }>();
  const router = useRouter();
  useRequireAuth();
  const { user } = useAuthStore();
  const canManage =
    isGranted(user, 'admin') || isGranted(user, 'manager') || isGranted(user, 'secretary');
  // Plano integrado pode ser editado por qualquer profissional além do staff.
  const canEditPlan =
    canManage ||
    isGranted(user, 'doctor') ||
    isGranted(user, 'nutritionist') ||
    isGranted(user, 'psychologist') ||
    isGranted(user, 'physicalEducator');

  const { data: enrollments = [], isLoading } = usePatientContinuums(params.id);
  // Pega a primeira ativa (assumindo 1:1 paciente↔continuum ativo na V1).
  const activeEnrollmentSummary = enrollments.find((e) => e.status === 'active') ?? enrollments[0];
  const { data: enrollment } = useContinuumEnrollment(activeEnrollmentSummary?.id);

  const [enrollOpen, setEnrollOpen] = useState(false);

  if (isLoading) {
    return (
      <div className="flex h-96 items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[
          { label: 'Pacientes', href: '/patients' },
          { label: enrollment?.patient?.name ?? 'Paciente', href: `/patients/${params.id}` },
          { label: 'Continuum' },
        ]}
        title="Continuum"
        description={
          enrollment
            ? `${enrollment.patient?.name ?? ''} — jornada em curso`
            : 'Paciente ainda não está em um programa Continuum.'
        }
        actions={
          !enrollment && canManage
            ? [
                {
                  label: 'Iniciar Continuum',
                  icon: <Plus className="h-4 w-4" />,
                  onClick: () => setEnrollOpen(true),
                  variant: 'default',
                },
              ]
            : []
        }
      />

      {!enrollment ? (
        <Card>
          <CardContent className="py-10 text-center">
            <Workflow className="mx-auto mb-3 h-10 w-10 text-muted-foreground" />
            <p className="text-sm text-muted-foreground">
              Inicie o Continuum pra desenhar a jornada — Médico, Nutricionista, Psicólogo
              e Educador Físico em rotação semanal pelos próximos meses.
            </p>
          </CardContent>
        </Card>
      ) : (
        <>
          <IntegratedPlanEditor enrollment={enrollment} canEdit={canEditPlan} />
          <ContinuumTimeline enrollmentId={enrollment.id} canManage={canManage} />
        </>
      )}

      {enrollOpen && (
        <EnrollDialog
          patientId={params.id}
          open={enrollOpen}
          onClose={() => setEnrollOpen(false)}
        />
      )}
    </div>
  );
}

// =====================================================
// Timeline (separado pra hot-reload + memo)
// =====================================================

function ContinuumTimeline({
  enrollmentId,
  canManage,
}: {
  enrollmentId: string;
  canManage: boolean;
}) {
  const router = useRouter();
  const { data: enrollment } = useContinuumEnrollment(enrollmentId);
  const update = useUpdateContinuumItem(enrollmentId);
  // Carrega todos boxes (cross-paciente) e indexa por continuumItemId pra
  // exibir card detalhado nos items type=box. List é leve (joins simples).
  const { data: allBoxes = [] } = useContinuumBoxes();
  const boxesByItemId = useMemo(() => {
    const m: Record<string, (typeof allBoxes)[number]> = {};
    for (const b of allBoxes) m[b.continuumItemId] = b;
    return m;
  }, [allBoxes]);
  const [filterType, setFilterType] = useState<string>('all');
  const [filterStatus, setFilterStatus] = useState<string>('all');

  const items = enrollment?.items ?? [];
  const filtered = useMemo(() => {
    return items.filter((it) => {
      if (filterType !== 'all' && it.type !== filterType) return false;
      if (filterStatus !== 'all' && it.status !== filterStatus) return false;
      return true;
    });
  }, [items, filterType, filterStatus]);

  const grouped = useMemo(() => {
    const map = new Map<number, PatientContinuumItem[]>();
    for (const it of filtered) {
      const arr = map.get(it.weekOffset) ?? [];
      arr.push(it);
      map.set(it.weekOffset, arr);
    }
    return Array.from(map.entries()).sort((a, b) => a[0] - b[0]);
  }, [filtered]);

  const stats = useMemo(() => {
    const total = items.length;
    const completed = items.filter((i) => i.status === 'completed').length;
    const missed = items.filter((i) => i.status === 'missed').length;
    const scheduled = items.filter((i) => i.status === 'scheduled').length;
    return { total, completed, missed, scheduled };
  }, [items]);

  if (!enrollment) return null;

  const startDate = new Date(enrollment.startDate);
  const endDate = new Date(enrollment.endDate);
  const totalWeeks = Math.max(1, Math.round((endDate.getTime() - startDate.getTime()) / (7 * 24 * 3600 * 1000)));
  const currentWeek = Math.max(0, Math.min(totalWeeks, differenceInWeeks(new Date(), startDate)));

  const handleQuickStatus = async (item: PatientContinuumItem, status: ContinuumItemStatus) => {
    try {
      await update.mutateAsync({
        itemId: item.id,
        payload: {
          status,
          ...(status === 'completed' ? { completedAt: new Date().toISOString() } : {}),
        },
      });
      toast.success('Item atualizado');
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao atualizar');
    }
  };

  return (
    <>
      {/* Header da inscrição */}
      <Card>
        <CardContent className="grid gap-4 py-6 md:grid-cols-4">
          <div>
            <p className="text-xs uppercase text-muted-foreground">Período</p>
            <p className="text-sm font-medium">
              {formatDate(startDate, "dd 'de' MMM yyyy")} —{' '}
              {formatDate(endDate, "dd 'de' MMM yyyy")}
            </p>
            <p className="text-xs text-muted-foreground">{totalWeeks} semanas</p>
          </div>
          <div>
            <p className="text-xs uppercase text-muted-foreground">Semana atual</p>
            <p className="text-2xl font-semibold">{currentWeek + 1}</p>
            <p className="text-xs text-muted-foreground">de {totalWeeks}</p>
          </div>
          <div>
            <p className="text-xs uppercase text-muted-foreground">Concluído</p>
            <p className="text-2xl font-semibold">
              {stats.completed}/{stats.total}
            </p>
            <p className="text-xs text-muted-foreground">
              {stats.total > 0 ? Math.round((stats.completed / stats.total) * 100) : 0}%
            </p>
          </div>
          <div>
            <p className="text-xs uppercase text-muted-foreground">Atrasados</p>
            <p
              className={cn(
                'text-2xl font-semibold',
                stats.missed > 0 ? 'text-red-600' : '',
              )}
            >
              {stats.missed}
            </p>
            <p className="text-xs text-muted-foreground">{stats.scheduled} agendados</p>
          </div>
        </CardContent>
      </Card>

      {/* Filtros */}
      <div className="flex flex-wrap gap-2">
        <Select value={filterType} onValueChange={setFilterType}>
          <SelectTrigger className="h-9 w-44">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Todos os tipos</SelectItem>
            {Object.entries(ITEM_TYPE_LABELS).map(([k, v]) => (
              <SelectItem key={k} value={k}>
                {v}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        <Select value={filterStatus} onValueChange={setFilterStatus}>
          <SelectTrigger className="h-9 w-44">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Todos os status</SelectItem>
            {Object.entries(ITEM_STATUS_LABELS).map(([k, v]) => (
              <SelectItem key={k} value={k}>
                {v}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      {/* Timeline */}
      <div className="space-y-4">
        {grouped.length === 0 ? (
          <Card>
            <CardContent className="py-10 text-center text-sm text-muted-foreground">
              Nenhum item bate com o filtro.
            </CardContent>
          </Card>
        ) : (
          grouped.map(([week, weekItems]) => (
            <Card key={week}>
              <CardHeader className="pb-2">
                <div className="flex items-center justify-between">
                  <CardTitle className="text-sm font-medium">
                    Semana {week + 1}
                    {week === currentWeek && (
                      <Badge variant="secondary" className="ml-2">
                        Atual
                      </Badge>
                    )}
                  </CardTitle>
                  <span className="text-xs text-muted-foreground">
                    {formatDate(addDays(startDate, week * 7), "dd 'de' MMM")}
                  </span>
                </div>
              </CardHeader>
              <CardContent className="space-y-2">
                {weekItems.map((it) => {
                  const Icon = TYPE_ICON[it.type] ?? Workflow;
                  const isLate = it.status === 'missed';
                  const linkedBox = it.type === 'box' ? boxesByItemId[it.id] : undefined;
                  if (linkedBox) {
                    return (
                      <BoxCard
                        key={it.id}
                        box={linkedBox}
                        canManage={canManage}
                        showPatient={false}
                      />
                    );
                  }
                  return (
                    <div
                      key={it.id}
                      className={cn(
                        'flex flex-wrap items-center gap-3 rounded-md border bg-background px-3 py-2',
                        isLate && 'border-red-200 bg-red-50/50',
                      )}
                    >
                      <Icon
                        className={cn(
                          'h-4 w-4 shrink-0',
                          isLate ? 'text-red-600' : 'text-muted-foreground',
                        )}
                      />
                      <div className="min-w-[180px] flex-1">
                        <p className="text-sm font-medium">{it.title}</p>
                        <p className="text-xs text-muted-foreground">
                          {ITEM_TYPE_LABELS[it.type]}
                          {it.specialty && ` · ${SPECIALTY_LABELS[it.specialty]}`}
                          {' · '}
                          {formatDate(it.expectedDate, "dd/MM")}
                          {isLate && (
                            <>
                              {' · '}
                              <span className="font-medium text-red-600">Atrasado</span>
                            </>
                          )}
                        </p>
                      </div>
                      <Badge variant="outline" className={cn('text-xs', ITEM_STATUS_COLORS[it.status])}>
                        {ITEM_STATUS_LABELS[it.status]}
                      </Badge>
                      {/* Botão Agendar — só pra items appointment ainda não ancorados. */}
                      {canManage &&
                        it.type === 'appointment' &&
                        (it.status === 'pending' || it.status === 'missed') &&
                        !it.appointmentId && (
                          <Button
                            size="sm"
                            variant="outline"
                            className="h-8 text-xs"
                            onClick={() => {
                              const params = new URLSearchParams({
                                patientId: enrollment.patientId,
                                continuumItemId: it.id,
                              });
                              if (it.specialty) params.set('specialty', it.specialty);
                              router.push(`/appointments/new?${params.toString()}`);
                            }}
                          >
                            Agendar
                          </Button>
                        )}
                      {/* Quando já ancorado, link rápido pra abrir a consulta. */}
                      {it.appointmentId && (
                        <Button
                          size="sm"
                          variant="ghost"
                          className="h-8 text-xs"
                          onClick={() => router.push(`/appointments/${it.appointmentId}`)}
                        >
                          Abrir consulta
                        </Button>
                      )}
                      {canManage && (
                        <Select
                          value={it.status}
                          onValueChange={(v) =>
                            handleQuickStatus(it, v as ContinuumItemStatus)
                          }
                        >
                          <SelectTrigger className="h-8 w-32 text-xs">
                            <SelectValue />
                          </SelectTrigger>
                          <SelectContent>
                            {Object.entries(ITEM_STATUS_LABELS).map(([k, v]) => (
                              <SelectItem key={k} value={k} className="text-xs">
                                {v}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      )}
                    </div>
                  );
                })}
              </CardContent>
            </Card>
          ))
        )}
      </div>
    </>
  );
}

// =====================================================
// Enroll Dialog
// =====================================================

function EnrollDialog({
  patientId,
  open,
  onClose,
}: {
  patientId: string;
  open: boolean;
  onClose: () => void;
}) {
  const enroll = useEnrollPatientContinuum(patientId);
  const { data: templates = [] } = useContinuumTemplates();
  const { data: doctors = [] } = useDoctors();

  const [templateId, setTemplateId] = useState('');
  const [startDate, setStartDate] = useState(format(new Date(), 'yyyy-MM-dd'));
  const [coordinatorId, setCoordinatorId] = useState('');

  const handleEnroll = async () => {
    if (!templateId) {
      toast.error('Escolha um template');
      return;
    }
    try {
      await enroll.mutateAsync({
        templateId,
        startDate,
        coordinatorDoctorId: coordinatorId || undefined,
      });
      toast.success('Continuum iniciado');
      onClose();
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao iniciar');
    }
  };

  return (
    <Dialog open={open} onOpenChange={(v) => !v && onClose()}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Iniciar Continuum</DialogTitle>
          <DialogDescription>
            Define o programa, a data de início e o coordenador. A timeline é montada
            no momento — secretária ancora as consultas reais conforme a equipe agenda.
          </DialogDescription>
        </DialogHeader>
        <div className="space-y-3">
          <div>
            <label className="mb-1 block text-xs font-medium">Template</label>
            <Select value={templateId} onValueChange={setTemplateId}>
              <SelectTrigger>
                <SelectValue placeholder="Escolha um programa" />
              </SelectTrigger>
              <SelectContent>
                {templates.map((t: { id: string; name: string; durationWeeks: number }) => (
                  <SelectItem key={t.id} value={t.id}>
                    {t.name} ({t.durationWeeks} sem)
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
          <div>
            <label className="mb-1 block text-xs font-medium">Data de início</label>
            <Input
              type="date"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>
          <div>
            <label className="mb-1 block text-xs font-medium">
              Coordenador (opcional — recebe alertas de atraso)
            </label>
            <Select value={coordinatorId} onValueChange={setCoordinatorId}>
              <SelectTrigger>
                <SelectValue placeholder="Sem coordenador" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="">Sem coordenador</SelectItem>
                {doctors.map((d: { id: string; name: string }) => (
                  <SelectItem key={d.id} value={d.id}>
                    {d.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
            Cancelar
          </Button>
          <Button onClick={handleEnroll} disabled={enroll.isPending || !templateId}>
            {enroll.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : 'Iniciar Continuum'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
