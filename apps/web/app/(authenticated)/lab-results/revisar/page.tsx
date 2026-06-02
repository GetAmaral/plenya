'use client';

/**
 * Results inbox — "Exames a revisar" (P3 frente 1).
 * Fila cross-patient de lotes de resultados aguardando ciência clínica.
 * Críticos (level 0/1) no topo. Dar ciência grava reviewed_at/by (auditado).
 * NÃO usa useRequireSelectedPatient — é cross-patient por design.
 */

import { useState } from 'react';
import Link from 'next/link';
import { AlertTriangle, ChevronDown, ChevronRight, Loader2, Check, FileText, ExternalLink } from 'lucide-react';
import { toast } from 'sonner';
import { useRequireAuth } from '@/lib/use-auth';
import { Card, CardContent } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { cn } from '@/lib/utils';
import { getScoreLevelColors } from '@/lib/utils/score-level-colors';
import {
  useLabInbox,
  useLabInboxDetail,
  useAcknowledgeBatch,
  type LabInboxItem,
} from '@/lib/api/lab-inbox';

const LEVEL_LABEL: Record<number, string> = {
  0: 'Crítico',
  1: 'Muito alterado',
  2: 'Subótimo',
  3: 'Limítrofe',
  4: 'Bom',
  5: 'Ótimo',
  6: 'N/A',
};

function fmtDate(iso?: string) {
  if (!iso) return '—';
  try {
    return new Date(iso).toLocaleDateString('pt-BR');
  } catch {
    return '—';
  }
}

function ResultRows({ batchId }: { batchId: string }) {
  const { data, isLoading } = useLabInboxDetail(batchId);
  if (isLoading) {
    return (
      <div className="flex items-center gap-2 px-4 py-3 text-sm text-muted-foreground">
        <Loader2 className="h-4 w-4 animate-spin" /> Carregando resultados…
      </div>
    );
  }
  const rows = data?.labResults ?? [];
  if (rows.length === 0) {
    return <p className="px-4 py-3 text-sm text-muted-foreground">Sem resultados detalhados.</p>;
  }
  return (
    <div className="overflow-x-auto px-2 pb-3">
      <table className="w-full text-sm">
        <thead>
          <tr className="text-left text-xs uppercase text-muted-foreground">
            <th className="px-2 py-1 font-medium">Exame</th>
            <th className="px-2 py-1 font-medium">Resultado</th>
            <th className="px-2 py-1 font-medium">Nível</th>
          </tr>
        </thead>
        <tbody>
          {rows.map((r) => {
            const lvl = r.level;
            const colors = lvl != null ? getScoreLevelColors(lvl) : null;
            const value =
              r.resultNumeric != null
                ? `${r.resultNumeric}${r.unit ? ' ' + r.unit : ''}`
                : r.resultText ?? '—';
            return (
              <tr key={r.id} className="border-t">
                <td className="px-2 py-1.5">{r.labTestDefinition?.name ?? r.testName}</td>
                <td className="px-2 py-1.5 font-medium">{value}</td>
                <td className="px-2 py-1.5">
                  {lvl != null && colors ? (
                    <span
                      className={cn(
                        'inline-block rounded border px-2 py-0.5 text-xs',
                        colors.bg,
                        colors.text,
                        colors.border,
                      )}
                    >
                      {LEVEL_LABEL[lvl] ?? lvl}
                    </span>
                  ) : (
                    <span className="text-xs text-muted-foreground">não classificado</span>
                  )}
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}

function InboxRow({ item }: { item: LabInboxItem }) {
  const [open, setOpen] = useState(false);
  const acknowledge = useAcknowledgeBatch();

  const handleAck = async () => {
    try {
      await acknowledge.mutateAsync(item.id);
      toast.success('Ciência registrada', { description: `${item.patientName} — ${item.laboratoryName}` });
    } catch (err) {
      toast.error('Falha ao dar ciência', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  return (
    <Card className={cn('overflow-hidden', item.isCritical && 'border-red-300')}>
      <CardContent className="p-0">
        <div className="flex items-center gap-3 px-4 py-3">
          <button
            type="button"
            onClick={() => setOpen((v) => !v)}
            className="flex flex-1 items-center gap-3 text-left"
          >
            {open ? (
              <ChevronDown className="h-4 w-4 shrink-0 text-muted-foreground" />
            ) : (
              <ChevronRight className="h-4 w-4 shrink-0 text-muted-foreground" />
            )}
            <div className="min-w-0 flex-1">
              <div className="flex items-center gap-2">
                <span className="truncate font-medium">{item.patientName || 'Paciente'}</span>
                {item.isCritical && (
                  <Badge variant="secondary" className="border-red-200 bg-red-100 text-red-900">
                    <AlertTriangle className="mr-1 h-3 w-3" /> Crítico
                  </Badge>
                )}
              </div>
              <p className="truncate text-sm text-muted-foreground">
                {item.laboratoryName} · coleta {fmtDate(item.collectionDate)} · resultado {fmtDate(item.resultDate)}
              </p>
            </div>
            <div className="hidden shrink-0 text-right text-xs text-muted-foreground sm:block">
              <div>
                {item.abnormalCount} alterado{item.abnormalCount === 1 ? '' : 's'} de {item.totalResults}
              </div>
              {item.worstLevel != null && <div>pior nível: {LEVEL_LABEL[item.worstLevel] ?? item.worstLevel}</div>}
            </div>
          </button>
          <div className="flex shrink-0 items-center gap-2">
            <Link href={`/patients/${item.patientId}`} title="Abrir prontuário">
              <Button variant="ghost" size="sm">
                <ExternalLink className="h-4 w-4" />
              </Button>
            </Link>
            <Button size="sm" onClick={handleAck} disabled={acknowledge.isPending}>
              {acknowledge.isPending ? (
                <Loader2 className="mr-1 h-4 w-4 animate-spin" />
              ) : (
                <Check className="mr-1 h-4 w-4" />
              )}
              Dar ciência
            </Button>
          </div>
        </div>
        {open && (
          <div className="border-t bg-muted/30">
            <ResultRows batchId={item.id} />
          </div>
        )}
      </CardContent>
    </Card>
  );
}

export default function LabResultsInboxPage() {
  useRequireAuth();
  const { data, isLoading, isError } = useLabInbox();

  const items = data ?? [];
  const criticalCount = items.filter((i) => i.isCritical).length;

  return (
    <div className="mx-auto max-w-4xl space-y-4 p-4 sm:p-6">
      <header className="space-y-1">
        <div className="flex items-center gap-2">
          <FileText className="h-5 w-5 text-primary" />
          <h1 className="text-xl font-semibold">Exames a revisar</h1>
        </div>
        <p className="text-sm text-muted-foreground">
          Lotes de resultados aguardando ciência clínica. Críticos no topo. Dar ciência fica registrado na auditoria.
        </p>
      </header>

      {!isLoading && items.length > 0 && (
        <div className="flex items-center gap-3 text-sm">
          <Badge variant="secondary">{items.length} a revisar</Badge>
          {criticalCount > 0 && (
            <Badge variant="secondary" className="border-red-200 bg-red-100 text-red-900">
              <AlertTriangle className="mr-1 h-3 w-3" /> {criticalCount} crítico{criticalCount === 1 ? '' : 's'}
            </Badge>
          )}
        </div>
      )}

      {isLoading ? (
        <div className="flex items-center gap-2 py-10 text-muted-foreground">
          <Loader2 className="h-5 w-5 animate-spin" /> Carregando fila…
        </div>
      ) : isError ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            Não foi possível carregar a fila de exames.
          </CardContent>
        </Card>
      ) : items.length === 0 ? (
        <Card>
          <CardContent className="flex flex-col items-center gap-2 py-12 text-center">
            <Check className="h-8 w-8 text-emerald-500" />
            <p className="font-medium">Nenhum exame pendente de revisão</p>
            <p className="text-sm text-muted-foreground">Todos os resultados disponíveis já receberam ciência.</p>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-2">
          {items.map((item) => (
            <InboxRow key={item.id} item={item} />
          ))}
        </div>
      )}
    </div>
  );
}
