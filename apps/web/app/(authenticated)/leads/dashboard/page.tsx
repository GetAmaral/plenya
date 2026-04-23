'use client';

import { useState } from 'react';
import Link from 'next/link';
import { ArrowLeft, BarChart3 } from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { Badge } from '@/components/ui/badge';
import { useRequireAuth } from '@/lib/use-auth';
import { usePatientGuard } from '@/lib/use-patient-guard';
import {
  useLeadStats,
  SOURCE_LABELS,
  STATUS_LABELS,
  STATUS_COLORS,
  type LeadSource,
  type LeadStatus,
  type StatsPeriodKey,
} from '@/lib/api/leads-api';

const STATUS_ORDER: LeadStatus[] = [
  'new',
  'contacted',
  'qualified',
  'converted',
  'lost',
  'unsubscribed',
];

const SOURCE_ORDER: LeadSource[] = [
  'light_claim',
  'contact_form',
  'whatsapp_inbound',
  'newsletter',
  'manual',
];

function formatDuration(seconds?: number): string {
  if (!seconds || seconds <= 0) return '—';
  const days = Math.floor(seconds / 86400);
  const hours = Math.floor((seconds % 86400) / 3600);
  const mins = Math.floor((seconds % 3600) / 60);
  if (days > 0) return `${days}d ${hours}h`;
  if (hours > 0) return `${hours}h ${mins}min`;
  return `${mins}min`;
}

export default function LeadDashboardPage() {
  useRequireAuth();
  usePatientGuard();

  const [period, setPeriod] = useState<StatsPeriodKey>('30d');
  const [customFrom, setCustomFrom] = useState('');
  const [customTo, setCustomTo] = useState('');

  const { data: stats, isLoading } = useLeadStats(period, customFrom, customTo);

  const maxByDay = Math.max(...(stats?.byDay.map((d) => d.count) ?? [1]), 1);
  const maxBySource = Math.max(...Object.values(stats?.bySource ?? { x: 1 }), 1);

  return (
    <div className="space-y-6 p-6">
      <Link
        href="/leads"
        className="inline-flex items-center text-sm text-muted-foreground hover:text-foreground"
      >
        <ArrowLeft className="mr-1 h-4 w-4" /> Lista de leads
      </Link>

      <div className="flex flex-wrap items-center justify-between gap-3">
        <h1 className="flex items-center gap-2 text-2xl font-light">
          <BarChart3 className="h-5 w-5" /> Dashboard de Leads
        </h1>
        <div className="flex items-center gap-2">
          <select
            value={period}
            onChange={(e) => setPeriod(e.target.value as StatsPeriodKey)}
            className="h-9 rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm focus:outline-none focus:ring-1 focus:ring-ring"
          >
            <option value="7d">Últimos 7 dias</option>
            <option value="30d">Últimos 30 dias</option>
            <option value="90d">Últimos 90 dias</option>
            <option value="custom">Período customizado…</option>
          </select>
          {period === 'custom' && (
            <>
              <input
                type="date"
                value={customFrom}
                onChange={(e) => setCustomFrom(e.target.value)}
                className="h-9 rounded-md border border-input bg-transparent px-3 py-1 text-sm"
              />
              <span className="text-sm">→</span>
              <input
                type="date"
                value={customTo}
                onChange={(e) => setCustomTo(e.target.value)}
                className="h-9 rounded-md border border-input bg-transparent px-3 py-1 text-sm"
              />
            </>
          )}
        </div>
      </div>

      {isLoading || !stats ? (
        <div className="space-y-4">
          <Skeleton className="h-32 w-full" />
          <Skeleton className="h-64 w-full" />
        </div>
      ) : (
        <>
          {/* Totais por status */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Totais — {stats.period.label}</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-7">
                <div className="rounded-md border p-4">
                  <div className="text-xs uppercase text-muted-foreground">Total</div>
                  <div className="mt-1 text-3xl font-light tabular-nums">
                    {stats.totals.all ?? 0}
                  </div>
                </div>
                {STATUS_ORDER.map((s) => (
                  <div key={s} className={`rounded-md border p-4 ${STATUS_COLORS[s]}`}>
                    <div className="text-xs uppercase">{STATUS_LABELS[s]}</div>
                    <div className="mt-1 text-3xl font-light tabular-nums">
                      {stats.totals[s] ?? 0}
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>

          {/* Conversão + tempos */}
          <div className="grid gap-4 md:grid-cols-3">
            <Card>
              <CardHeader>
                <CardTitle className="text-sm">Taxa de conversão</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-light tabular-nums">
                  {(stats.conversionRate * 100).toFixed(1)}%
                </div>
                <p className="mt-1 text-xs text-muted-foreground">
                  Convertidos / total no período
                </p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle className="text-sm">Tempo médio até 1º contato</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-light">
                  {formatDuration(stats.avgTimeToContact)}
                </div>
                <p className="mt-1 text-xs text-muted-foreground">
                  Captura → primeira mensagem enviada
                </p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader>
                <CardTitle className="text-sm">Tempo médio até conversão</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="text-4xl font-light">
                  {formatDuration(stats.avgTimeToConvert)}
                </div>
                <p className="mt-1 text-xs text-muted-foreground">
                  Captura → status convertido
                </p>
              </CardContent>
            </Card>
          </div>

          {/* Origem (bar chart) */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Por origem</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                {SOURCE_ORDER.map((source) => {
                  const count = stats.bySource[source] ?? 0;
                  const pct = (count / maxBySource) * 100;
                  return (
                    <div key={source} className="flex items-center gap-3">
                      <Badge variant="outline" className="w-44 justify-start">
                        {SOURCE_LABELS[source]}
                      </Badge>
                      <div className="flex-1 overflow-hidden rounded bg-muted">
                        <div
                          className="h-6 bg-primary/70"
                          style={{ width: `${pct}%` }}
                        />
                      </div>
                      <span className="w-10 text-right tabular-nums text-sm">{count}</span>
                    </div>
                  );
                })}
              </div>
            </CardContent>
          </Card>

          {/* Ao longo do tempo (sparkline) */}
          {stats.byDay.length > 0 && (
            <Card>
              <CardHeader>
                <CardTitle className="text-base">Captura ao longo do tempo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="flex h-32 items-end gap-px">
                  {stats.byDay.map((d) => (
                    <div
                      key={d.date}
                      title={`${d.date}: ${d.count}`}
                      className="flex-1 bg-primary/40 hover:bg-primary"
                      style={{
                        height: `${(d.count / maxByDay) * 100}%`,
                        minHeight: d.count > 0 ? '2px' : '0',
                      }}
                    />
                  ))}
                </div>
                <div className="mt-2 flex justify-between text-xs text-muted-foreground">
                  <span>{stats.byDay[0]?.date}</span>
                  <span>{stats.byDay[stats.byDay.length - 1]?.date}</span>
                </div>
              </CardContent>
            </Card>
          )}
        </>
      )}
    </div>
  );
}
