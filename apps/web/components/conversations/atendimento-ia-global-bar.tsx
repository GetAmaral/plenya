'use client';

/**
 * AtendimentoIAGlobalBar — controle GLOBAL do Atendimento IA no topo da Central de Conversas.
 * Um seletor de três estados (Manual · Copiloto · Auto) que define o modo baseline global +
 * um banner do estado efetivo (quando a janela de horário eleva pra Auto, mostra "Auto até HH:MM").
 * A config detalhada (janela semanal + tempos X/Y/Z) fica numa tela dedicada
 * (/configuracoes/atendimento-ia), não aqui. Plano: docs/emr/plano-atendimento-ia-global.md.
 *
 * Rótulo Manual = enum `off` (não muda o valor armazenado).
 */
import Link from 'next/link';
import { Loader2, Settings2 } from 'lucide-react';
import { toast } from 'sonner';

import { cn } from '@/lib/utils';
import {
  type AutomationMode,
  type WeeklySchedule,
  useReceptionSettings,
  useUpdateReceptionSettings,
} from '@/lib/api/conversations-api';

const MODES: { value: AutomationMode; label: string; title: string }[] = [
  { value: 'off', label: 'Manual', title: 'IA não age; a equipe responde tudo' },
  { value: 'copilot', label: 'Copiloto', title: 'IA rascunha; a equipe revisa e envia' },
  { value: 'auto', label: 'Auto', title: 'IA responde sozinha após o tempo de resposta' },
];

const DAY_KEYS = ['sun', 'mon', 'tue', 'wed', 'thu', 'fri', 'sat'] as const;

function hhmmToMinutes(s: string): number | null {
  const m = /^(\d{1,2}):(\d{2})$/.exec(s.trim());
  if (!m) return null;
  const h = Number(m[1]);
  const min = Number(m[2]);
  if (h < 0 || h > 23 || min < 0 || min > 59) return null;
  return h * 60 + min;
}

/**
 * Replica isAutoActiveAt do backend: dado o agora, devolve o horário "HH:MM" em que a janela
 * ativa termina (ou 'all-day' se faixa cobre o dia todo, ou null se não há janela ativa).
 * Trata faixas que cruzam a meia-noite (start > end), inclusive o spill do dia anterior.
 */
function autoWindowEnd(schedule: WeeklySchedule, now: Date): string | 'all-day' | null {
  const hm = now.getHours() * 60 + now.getMinutes();
  const today = DAY_KEYS[now.getDay()];
  const yest = DAY_KEYS[(now.getDay() + 6) % 7];

  for (const r of schedule[today] ?? []) {
    const st = hhmmToMinutes(r.start);
    const en = hhmmToMinutes(r.end);
    if (st == null || en == null) continue;
    if (st < en) {
      if (hm >= st && hm < en) return r.end;
    } else if (st > en) {
      if (hm >= st) return r.end; // termina no end do dia seguinte
    } else {
      return 'all-day';
    }
  }
  for (const r of schedule[yest] ?? []) {
    const st = hhmmToMinutes(r.start);
    const en = hhmmToMinutes(r.end);
    if (st == null || en == null) continue;
    if (st > en && hm < en) return r.end; // madrugada herdada do dia anterior
  }
  return null;
}

export function AtendimentoIAGlobalBar() {
  const { data, isLoading } = useReceptionSettings();
  const update = useUpdateReceptionSettings();

  if (isLoading || !data) {
    return (
      <div className="flex items-center gap-2 rounded-lg border border-border bg-card px-3 py-2 text-xs text-muted-foreground">
        <Loader2 className="h-3.5 w-3.5 animate-spin" /> Atendimento IA…
      </div>
    );
  }

  const baseline = data.baselineMode;
  const effective = data.effectiveMode;
  // Banner: quando a janela elevou pra Auto (efetivo=auto mas baseline≠auto), mostra "até HH:MM".
  const elevatedByWindow = data.autoActiveNow && effective === 'auto' && baseline !== 'auto';
  const windowEnd = elevatedByWindow ? autoWindowEnd(data.schedule ?? {}, new Date()) : null;

  const choose = (mode: AutomationMode) => {
    if (mode === baseline) return;
    update.mutate(
      { baselineMode: mode },
      {
        onError: (e) =>
          toast.error(e instanceof Error ? e.message : 'Falha ao mudar o Atendimento IA'),
      }
    );
  };

  return (
    <div className="flex flex-wrap items-center gap-x-3 gap-y-2 rounded-lg border border-border bg-card px-3 py-2">
      <span className="text-xs font-medium text-foreground">Atendimento IA</span>

      <div className="inline-flex overflow-hidden rounded-md border border-border" role="group" aria-label="Modo global do Atendimento IA">
        {MODES.map((opt) => (
          <button
            key={opt.value}
            type="button"
            onClick={() => choose(opt.value)}
            disabled={update.isPending}
            title={opt.title}
            aria-pressed={baseline === opt.value}
            className={cn(
              'px-2.5 py-1 text-xs font-medium transition-colors disabled:opacity-50',
              baseline === opt.value
                ? 'bg-primary text-primary-foreground'
                : 'bg-background text-muted-foreground hover:bg-ocean-50'
            )}
          >
            {baseline === opt.value && update.isPending ? (
              <Loader2 className="h-3 w-3 animate-spin" />
            ) : (
              opt.label
            )}
          </button>
        ))}
      </div>

      {/* Banner de estado efetivo */}
      {elevatedByWindow ? (
        <span className="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2 py-0.5 text-[11px] font-medium text-emerald-900" title="A janela de horário está elevando o modo para Auto">
          Auto{windowEnd && windowEnd !== 'all-day' ? ` até ${windowEnd}` : ''}
        </span>
      ) : (
        baseline !== 'off' && (
          <span
            className={cn(
              'inline-flex items-center rounded-full px-2 py-0.5 text-[11px] font-medium',
              effective === 'auto'
                ? 'bg-emerald-100 text-emerald-900'
                : 'bg-violet-100 text-violet-900'
            )}
          >
            {effective === 'auto' ? 'Respondendo sozinha' : 'Rascunhando p/ revisão'}
          </span>
        )
      )}

      {data.scheduleEnabled && !elevatedByWindow && (
        <span className="text-[11px] text-muted-foreground" title="Janela de horário do Auto habilitada (fora da janela vale o modo acima)">
          janela ativa
        </span>
      )}

      <Link
        href="/configuracoes/atendimento-ia"
        className="ml-auto inline-flex items-center gap-1 rounded-md border border-border bg-background px-2 py-1 text-[11px] font-medium text-muted-foreground transition-colors hover:bg-muted"
        title="Configurar janela de horário e tempos de resposta"
      >
        <Settings2 className="h-3.5 w-3.5" /> Configurar
      </Link>
    </div>
  );
}
