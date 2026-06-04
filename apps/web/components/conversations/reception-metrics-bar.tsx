'use client';

/**
 * ReceptionMetricsBar — faixa compacta com as métricas do recepcionista virtual
 * (últimos 30 dias). Discreta; aparece no topo da Central de Conversas.
 */
import { Bot } from 'lucide-react';
import { useReceptionMetrics } from '@/lib/api/conversations-api';

function Stat({ label, value }: { label: string; value: number | string }) {
  return (
    <span className="inline-flex items-baseline gap-1">
      <strong className="text-foreground">{value}</strong>
      <span>{label}</span>
    </span>
  );
}

export function ReceptionMetricsBar() {
  const { data, isLoading } = useReceptionMetrics(30);
  if (isLoading || !data) return null;
  if (data.autoReplies === 0 && data.handoffs === 0 && data.conversationsWithBot === 0) {
    return null; // nada a mostrar ainda
  }

  return (
    <div className="flex flex-wrap items-center gap-x-4 gap-y-1 rounded-lg border border-violet-200 bg-violet-50/60 px-3 py-2 text-xs text-muted-foreground">
      <span className="inline-flex items-center gap-1 font-medium text-violet-800">
        <Bot className="h-3.5 w-3.5" /> Recepção IA · 30 dias
      </span>
      <Stat label="respostas" value={data.autoReplies} />
      <Stat label="conversas atendidas" value={data.conversationsWithBot} />
      <Stat label="handoffs" value={data.handoffs} />
      <Stat label="convertidos" value={data.convertedAfterBot} />
      {!data.globallyEnabled && (
        <span className="rounded bg-zinc-200 px-1 text-[10px] text-zinc-600" title="Auto-envio desligado globalmente">
          auto-envio off
        </span>
      )}
    </div>
  );
}
