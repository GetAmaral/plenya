'use client';

/**
 * AutomationToggle — controle do recepcionista virtual por conversa.
 * Off (IA não age) · Copiloto (IA sugere, humano envia) · Automático (IA assume
 * após o fallback de tempo). Reativar limpa eventual pausa de handoff.
 */
import { toast } from 'sonner';
import { Loader2 } from 'lucide-react';

import { cn } from '@/lib/utils';
import {
  type AutomationMode,
  type ConversationOwnerType,
  useConversationAutomation,
  useSetConversationAutomation,
} from '@/lib/api/conversations-api';

// Rótulos alinhados ao controle global (Manual = enum `off`).
const OPTIONS: { value: AutomationMode; label: string; title: string }[] = [
  { value: 'off', label: 'Manual', title: 'IA não age nesta conversa; você responde' },
  { value: 'copilot', label: 'Copiloto', title: 'IA rascunha; você revisa e envia' },
  { value: 'auto', label: 'Auto', title: 'IA responde sozinha após o tempo de resposta' },
];

export function AutomationToggle({
  ownerType,
  ownerId,
}: {
  ownerType: ConversationOwnerType;
  ownerId: string;
}) {
  const { data, isLoading } = useConversationAutomation(ownerType, ownerId);
  const setMode = useSetConversationAutomation(ownerType, ownerId);

  const current = data?.mode ?? 'off';
  const paused = data?.pausedUntil && new Date(data.pausedUntil) > new Date();
  const autoButGloballyOff = current === 'auto' && data && !data.globallyEnabled;

  const choose = (mode: AutomationMode) => {
    if (mode === current) return;
    setMode.mutate(
      { mode, fallbackMinutes: data?.fallbackMinutes ?? 0 },
      {
        onError: (e) =>
          toast.error(e instanceof Error ? e.message : 'Falha ao mudar modo da IA'),
      }
    );
  };

  return (
    <div className="hidden items-center gap-2 md:flex" title="Atendimento IA desta conversa (sobrepõe o global por 24h)">
      <span className="text-[10px] uppercase tracking-wide text-muted-foreground">IA</span>
      <div className="inline-flex overflow-hidden rounded-md border border-border">
        {OPTIONS.map((opt) => (
          <button
            key={opt.value}
            type="button"
            onClick={() => choose(opt.value)}
            disabled={isLoading || setMode.isPending}
            title={opt.title}
            className={cn(
              'px-2 py-1 text-[11px] font-medium transition-colors disabled:opacity-50',
              current === opt.value
                ? 'bg-violet-600 text-white'
                : 'bg-white text-muted-foreground hover:bg-violet-50'
            )}
          >
            {opt.value === current && setMode.isPending ? (
              <Loader2 className="h-3 w-3 animate-spin" />
            ) : (
              opt.label
            )}
          </button>
        ))}
      </div>
      {paused && (
        <span className="rounded bg-amber-100 px-1 text-[10px] text-amber-800" title="Pausada após handoff; reative para retomar">
          pausada
        </span>
      )}
      {autoButGloballyOff && (
        <span className="rounded bg-zinc-100 px-1 text-[10px] text-zinc-600" title="Auto-envio desligado globalmente (RECEPTION_BOT_ENABLED)">
          global off
        </span>
      )}
    </div>
  );
}
