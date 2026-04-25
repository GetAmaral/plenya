'use client';

/**
 * /continuum/boxes — visão cross-paciente da equipe de logística.
 * Lista boxes ordenados por expectedDate, filtráveis por status.
 * Default: planned + preparing (o que precisa de ação).
 */
import { useState } from 'react';
import { Loader2, Package } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { useAuthStore, isGranted } from '@/lib/auth-store';
import { Card, CardContent } from '@/components/ui/card';
import {
  ToggleGroup,
  ToggleGroupItem,
} from '@/components/ui/toggle-group';
import { PageHeader } from '@/components/layout/page-header';
import {
  useContinuumBoxes,
  useContinuumBoxCounts,
  BOX_STATUS_LABELS,
  type BoxStatus,
} from '@/lib/api/continuum-api';
import { BoxCard } from '@/components/continuum/box-card';

const STATUS_TABS: { key: 'pending' | BoxStatus | 'all'; label: string; statuses: BoxStatus[] }[] = [
  { key: 'pending', label: 'A despachar', statuses: ['planned', 'preparing'] },
  { key: 'shipped', label: 'Em trânsito', statuses: ['shipped'] },
  { key: 'delivered', label: 'Entregues', statuses: ['delivered'] },
  { key: 'all', label: 'Todos', statuses: [] },
];

export default function ContinuumBoxesPage() {
  useRequireAuth();
  const { user } = useAuthStore();
  const canManage =
    isGranted(user, 'admin') || isGranted(user, 'manager') || isGranted(user, 'secretary');

  const [tab, setTab] = useState<string>('pending');
  const tabConfig = STATUS_TABS.find((t) => t.key === tab) ?? STATUS_TABS[0];
  const { data: boxes = [], isLoading } = useContinuumBoxes(
    tabConfig.statuses.length > 0 ? tabConfig.statuses : undefined,
  );
  const { data: counts } = useContinuumBoxCounts();

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[{ label: 'Continuum', href: '/continuum' }, { label: 'Boxes' }]}
        title="Boxes Plenya"
        description="Visão da equipe de logística — planejados, em trânsito e entregues."
      />

      {/* Counts */}
      {counts && (
        <div className="grid grid-cols-2 gap-2 md:grid-cols-5">
          {(Object.keys(BOX_STATUS_LABELS) as BoxStatus[]).map((s) => (
            <Card key={s}>
              <CardContent className="py-3">
                <p className="text-xs uppercase text-muted-foreground">
                  {BOX_STATUS_LABELS[s]}
                </p>
                <p className="text-2xl font-semibold">{counts[s] ?? 0}</p>
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      <ToggleGroup
        type="single"
        value={tab}
        onValueChange={(v) => v && setTab(v)}
        className="justify-start"
      >
        {STATUS_TABS.map((t) => (
          <ToggleGroupItem key={t.key} value={t.key} size="sm">
            {t.label}
          </ToggleGroupItem>
        ))}
      </ToggleGroup>

      {isLoading ? (
        <Card className="flex h-40 items-center justify-center">
          <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
        </Card>
      ) : boxes.length === 0 ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            <Package className="mx-auto mb-2 h-8 w-8" />
            Nenhum box nesta visão.
          </CardContent>
        </Card>
      ) : (
        <div className="grid gap-3 md:grid-cols-2">
          {boxes.map((b) => (
            <BoxCard key={b.id} box={b} canManage={canManage} />
          ))}
        </div>
      )}
    </div>
  );
}
