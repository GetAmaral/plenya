'use client';

/**
 * /continuum — panorama da equipe (Fase 7).
 *
 * 3 visões via tabs:
 *  - Alertas: items missed + pending nos próximos 7 dias (default — visão de ação)
 *  - Por paciente: kanban com semana atual, progresso e próximo marco
 *  - Por semana: heatmap dia × especialidade pra programação
 */
import { useMemo, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { addDays, format, parseISO } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import {
  AlertTriangle,
  Calendar,
  ChevronLeft,
  ChevronRight,
  Loader2,
  Package,
  Users,
  Workflow,
} from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { PageHeader } from '@/components/layout/page-header';
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from '@/components/ui/tabs';
import {
  useContinuumDashboardAlerts,
  useContinuumDashboardPatients,
  useContinuumDashboardWeek,
  ITEM_STATUS_LABELS,
  ITEM_STATUS_COLORS,
  ITEM_TYPE_LABELS,
  SPECIALTY_LABELS,
  type AlertRow,
  type PerPatientRow,
  type PerWeekItem,
  type ContinuumItemSpecialty,
} from '@/lib/api/continuum-api';
import { cn } from '@/lib/utils';

export default function ContinuumDashboardPage() {
  useRequireAuth();
  const [tab, setTab] = useState('alerts');

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[{ label: 'Continuum' }]}
        title="Continuum Plenya"
        description="Acompanhamento contínuo se constrói no tempo. Aqui a equipe vê o que precisa de ação, quem está em curso e o que vem na semana."
      />

      <Tabs value={tab} onValueChange={setTab} className="space-y-4">
        <TabsList>
          <TabsTrigger value="alerts">
            <AlertTriangle className="mr-1 h-4 w-4" />
            Alertas
          </TabsTrigger>
          <TabsTrigger value="patients">
            <Users className="mr-1 h-4 w-4" />
            Por paciente
          </TabsTrigger>
          <TabsTrigger value="week">
            <Calendar className="mr-1 h-4 w-4" />
            Por semana
          </TabsTrigger>
        </TabsList>

        <TabsContent value="alerts">
          <AlertsView />
        </TabsContent>
        <TabsContent value="patients">
          <PerPatientView />
        </TabsContent>
        <TabsContent value="week">
          <PerWeekView />
        </TabsContent>
      </Tabs>

      {/* Atalhos pra módulos relacionados */}
      <div className="grid gap-3 md:grid-cols-3">
        <Link href="/continuum/templates">
          <Card className="transition hover:border-primary">
            <CardContent className="flex items-center gap-3 py-4">
              <Workflow className="h-5 w-5 text-primary" />
              <div>
                <p className="text-sm font-semibold">Templates Continuum</p>
                <p className="text-xs text-muted-foreground">Semestral, Anual, customizados</p>
              </div>
            </CardContent>
          </Card>
        </Link>
        <Link href="/continuum/box-templates">
          <Card className="transition hover:border-primary">
            <CardContent className="flex items-center gap-3 py-4">
              <Package className="h-5 w-5 text-primary" />
              <div>
                <p className="text-sm font-semibold">Templates de Box</p>
                <p className="text-xs text-muted-foreground">Boas-vindas, mensal, reavaliação</p>
              </div>
            </CardContent>
          </Card>
        </Link>
        <Link href="/continuum/boxes">
          <Card className="transition hover:border-primary">
            <CardContent className="flex items-center gap-3 py-4">
              <Package className="h-5 w-5 text-primary" />
              <div>
                <p className="text-sm font-semibold">Boxes em curso</p>
                <p className="text-xs text-muted-foreground">Logística — a despachar / em trânsito</p>
              </div>
            </CardContent>
          </Card>
        </Link>
      </div>
    </div>
  );
}

// =====================================================
// Alertas
// =====================================================

function AlertsView() {
  const router = useRouter();
  const { data: alerts = [], isLoading } = useContinuumDashboardAlerts(7);

  const grouped = useMemo(() => {
    const missed = alerts.filter((a) => a.severity === 'missed');
    const dueSoon = alerts.filter((a) => a.severity === 'due-soon');
    return { missed, dueSoon };
  }, [alerts]);

  if (isLoading) {
    return (
      <Card className="flex h-40 items-center justify-center">
        <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
      </Card>
    );
  }
  if (alerts.length === 0) {
    return (
      <Card>
        <CardContent className="py-10 text-center text-sm text-muted-foreground">
          Nenhum alerta — equipe em dia.
        </CardContent>
      </Card>
    );
  }

  return (
    <div className="space-y-4">
      {grouped.missed.length > 0 && (
        <Card className="border-red-200">
          <CardHeader className="pb-2">
            <CardTitle className="text-base text-red-700">
              <AlertTriangle className="mr-1 inline h-4 w-4" />
              Atrasados ({grouped.missed.length})
            </CardTitle>
            <CardDescription>Items que passaram da janela esperada.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            {grouped.missed.map((a) => (
              <AlertItemRow key={a.id} alert={a} onClick={() => router.push(`/patients/${a.patientId}/continuum`)} />
            ))}
          </CardContent>
        </Card>
      )}
      {grouped.dueSoon.length > 0 && (
        <Card>
          <CardHeader className="pb-2">
            <CardTitle className="text-base">Nos próximos 7 dias ({grouped.dueSoon.length})</CardTitle>
          </CardHeader>
          <CardContent className="space-y-2">
            {grouped.dueSoon.map((a) => (
              <AlertItemRow key={a.id} alert={a} onClick={() => router.push(`/patients/${a.patientId}/continuum`)} />
            ))}
          </CardContent>
        </Card>
      )}
    </div>
  );
}

function AlertItemRow({ alert: a, onClick }: { alert: AlertRow; onClick: () => void }) {
  return (
    <button
      type="button"
      onClick={onClick}
      className={cn(
        'flex w-full items-center gap-3 rounded-md border bg-background px-3 py-2 text-left transition hover:bg-accent',
        a.severity === 'missed' && 'border-red-200 bg-red-50/40',
      )}
    >
      <div className="min-w-[160px] flex-1">
        <p className="text-sm font-medium">{a.patientName}</p>
        <p className="text-xs text-muted-foreground">
          {a.title}
          {a.specialty && ` · ${SPECIALTY_LABELS[a.specialty]}`}
        </p>
      </div>
      <Badge variant="outline" className={cn('text-xs', ITEM_STATUS_COLORS[a.status])}>
        {ITEM_STATUS_LABELS[a.status]}
      </Badge>
      <span className="shrink-0 text-xs text-muted-foreground">
        {format(parseISO(a.expectedDate), "dd/MM", { locale: ptBR })}
      </span>
    </button>
  );
}

// =====================================================
// Por paciente
// =====================================================

function PerPatientView() {
  const router = useRouter();
  const { data: rows = [], isLoading } = useContinuumDashboardPatients();

  if (isLoading) {
    return (
      <Card className="flex h-40 items-center justify-center">
        <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
      </Card>
    );
  }
  if (rows.length === 0) {
    return (
      <Card>
        <CardContent className="py-10 text-center text-sm text-muted-foreground">
          Nenhum paciente com Continuum ativo.
        </CardContent>
      </Card>
    );
  }

  return (
    <div className="grid gap-3 md:grid-cols-2">
      {rows.map((r) => (
        <PerPatientCard
          key={r.continuumId}
          row={r}
          onOpen={() => router.push(`/patients/${r.patientId}/continuum`)}
        />
      ))}
    </div>
  );
}

function PerPatientCard({ row: r, onOpen }: { row: PerPatientRow; onOpen: () => void }) {
  const pct = r.totalItems > 0 ? Math.round((r.completedItems / r.totalItems) * 100) : 0;
  return (
    <Card className="cursor-pointer transition hover:border-primary" onClick={onOpen}>
      <CardHeader className="pb-2">
        <div className="flex items-center justify-between">
          <CardTitle className="text-base">{r.patientName}</CardTitle>
          {r.missedItems > 0 && (
            <Badge variant="outline" className="border-red-200 bg-red-50 text-red-700">
              {r.missedItems} atrasados
            </Badge>
          )}
        </div>
        <CardDescription>
          Semana {r.currentWeek} de {r.durationWeeks}
          {r.coordinatorName && ` · Coordenador: ${r.coordinatorName}`}
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-3">
        {/* Progress bar */}
        <div>
          <div className="flex justify-between text-xs text-muted-foreground">
            <span>Progresso</span>
            <span>
              {r.completedItems}/{r.totalItems} · {pct}%
            </span>
          </div>
          <div className="mt-1 h-2 overflow-hidden rounded-full bg-muted">
            <div
              className="h-full bg-primary transition-all"
              style={{ width: `${pct}%` }}
            />
          </div>
        </div>
        {/* Status pills */}
        <div className="flex flex-wrap gap-2 text-xs">
          <span className="rounded-md bg-emerald-50 px-2 py-0.5 text-emerald-700">
            ✓ {r.completedItems} concluídos
          </span>
          <span className="rounded-md bg-blue-50 px-2 py-0.5 text-blue-700">
            {r.scheduledItems} agendados
          </span>
          <span className="rounded-md bg-slate-100 px-2 py-0.5 text-slate-700">
            {r.pendingItems} pendentes
          </span>
        </div>
        {/* Próximo marco */}
        {r.nextItemTitle && r.nextItemDate && (
          <div className="rounded-md border bg-muted/30 px-2 py-1.5 text-xs">
            <p className="font-medium">Próximo: {r.nextItemTitle}</p>
            <p className="text-muted-foreground">
              {format(parseISO(r.nextItemDate), "dd 'de' MMM", { locale: ptBR })}
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  );
}

// =====================================================
// Por semana
// =====================================================

const WEEKDAYS = ['Dom', 'Seg', 'Ter', 'Qua', 'Qui', 'Sex', 'Sáb'];
const SPECIALTY_KEYS: ContinuumItemSpecialty[] = [
  'doctor',
  'nutritionist',
  'psychologist',
  'physicalEducator',
];

function PerWeekView() {
  const router = useRouter();
  // Início = domingo da semana atual (frontend só envia data, backend reusa default).
  const [weekStart, setWeekStart] = useState(() => {
    const now = new Date();
    const offset = now.getDay();
    const d = addDays(now, -offset);
    return format(d, 'yyyy-MM-dd');
  });
  const { data, isLoading } = useContinuumDashboardWeek(weekStart);

  const items = data?.items ?? [];
  const start = parseISO(weekStart);

  // Index por (dia 0..6, especialidade ou 'box')
  type Cell = { count: number; items: PerWeekItem[] };
  const grid: Record<string, Cell> = {};
  for (let d = 0; d < 7; d++) {
    for (const s of SPECIALTY_KEYS) grid[`${d}-${s}`] = { count: 0, items: [] };
    grid[`${d}-box`] = { count: 0, items: [] };
  }
  for (const it of items) {
    const day = Math.floor(
      (parseISO(it.expectedDate).getTime() - start.getTime()) / (24 * 3600 * 1000),
    );
    if (day < 0 || day > 6) continue;
    let key = '';
    if (it.type === 'appointment' && it.specialty) key = `${day}-${it.specialty}`;
    else if (it.type === 'box') key = `${day}-box`;
    else continue;
    const cell = grid[key];
    if (!cell) continue;
    cell.count += 1;
    cell.items.push(it);
  }

  const shiftWeek = (delta: number) => {
    const next = addDays(start, delta * 7);
    setWeekStart(format(next, 'yyyy-MM-dd'));
  };

  return (
    <div className="space-y-3">
      <div className="flex items-center gap-2">
        <Button variant="outline" size="sm" onClick={() => shiftWeek(-1)}>
          <ChevronLeft className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          size="sm"
          onClick={() => {
            const now = new Date();
            const offset = now.getDay();
            setWeekStart(format(addDays(now, -offset), 'yyyy-MM-dd'));
          }}
        >
          Hoje
        </Button>
        <Button variant="outline" size="sm" onClick={() => shiftWeek(1)}>
          <ChevronRight className="h-4 w-4" />
        </Button>
        <p className="ml-2 text-sm text-muted-foreground">
          {format(start, "dd 'de' MMM", { locale: ptBR })} —{' '}
          {format(addDays(start, 6), "dd 'de' MMM yyyy", { locale: ptBR })}
        </p>
      </div>

      {isLoading ? (
        <Card className="flex h-40 items-center justify-center">
          <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
        </Card>
      ) : (
        <Card>
          <CardContent className="overflow-x-auto p-0">
            <table className="w-full text-xs">
              <thead>
                <tr className="border-b">
                  <th className="px-3 py-2 text-left">Eixo</th>
                  {WEEKDAYS.map((d, i) => (
                    <th key={i} className="px-2 py-2 text-center">
                      <p>{d}</p>
                      <p className="font-normal text-muted-foreground">
                        {format(addDays(start, i), 'dd/MM')}
                      </p>
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {SPECIALTY_KEYS.map((spec) => (
                  <tr key={spec} className="border-b last:border-0">
                    <td className="px-3 py-2 font-medium">
                      {SPECIALTY_LABELS[spec]}
                    </td>
                    {Array.from({ length: 7 }).map((_, day) => {
                      const cell = grid[`${day}-${spec}`];
                      return <HeatCell key={day} cell={cell} router={router} />;
                    })}
                  </tr>
                ))}
                <tr>
                  <td className="px-3 py-2 font-medium">Box Plenya</td>
                  {Array.from({ length: 7 }).map((_, day) => {
                    const cell = grid[`${day}-box`];
                    return <HeatCell key={day} cell={cell} router={router} />;
                  })}
                </tr>
              </tbody>
            </table>
          </CardContent>
        </Card>
      )}
    </div>
  );
}

function HeatCell({
  cell,
  router,
}: {
  cell: { count: number; items: PerWeekItem[] };
  router: ReturnType<typeof useRouter>;
}) {
  const intensity = Math.min(cell.count / 5, 1);
  const bg =
    cell.count === 0
      ? 'bg-muted/30'
      : intensity < 0.4
      ? 'bg-blue-100'
      : intensity < 0.7
      ? 'bg-blue-200'
      : 'bg-blue-300';
  return (
    <td className="p-1 text-center">
      <div
        className={cn('group relative rounded-md py-2 transition', bg, cell.count > 0 && 'cursor-pointer hover:ring-2 hover:ring-primary/30')}
        onClick={() => {
          if (cell.count === 0) return;
          // Abre primeiro paciente — UX simples; futuramente popover com lista.
          const it = cell.items[0]!;
          router.push(`/patients/${it.patientId}/continuum`);
        }}
      >
        <p className="text-sm font-semibold">{cell.count > 0 ? cell.count : '—'}</p>
        {cell.count > 0 && (
          <div className="absolute left-1/2 top-full z-10 hidden -translate-x-1/2 translate-y-1 whitespace-nowrap rounded-md border bg-popover px-2 py-1 text-xs shadow-md group-hover:block">
            {cell.items.slice(0, 5).map((it) => (
              <p key={it.id} className="text-left">
                {it.patientName}
              </p>
            ))}
            {cell.items.length > 5 && <p>+{cell.items.length - 5}</p>}
          </div>
        )}
      </div>
    </td>
  );
}
