'use client';

/**
 * /configuracoes/atendimento-ia
 *
 * Tela dedicada de configuração do Atendimento IA (recepcionista virtual):
 *  - Modo baseline global (Manual · Copiloto · Auto)
 *  - Janela de horário do Auto (grade semanal estilo Google; faixas podem cruzar a meia-noite)
 *  - Tempos: X (resposta/debounce), Y (fallback Copiloto→Auto), Z (alerta no Manual)
 *
 * Quem configura: admin, secretary, manager, doctor (backend RequireRole no /reception).
 * Plano: docs/emr/plano-atendimento-ia-global.md.
 */
import { useEffect, useMemo, useState } from 'react';
import { Plus, Trash2, Loader2, Clock, Save } from 'lucide-react';
import { toast } from 'sonner';

import { useRequireAuth } from '@/lib/use-auth';
import { cn } from '@/lib/utils';
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';
import { Skeleton } from '@/components/ui/skeleton';
import { PageHeader } from '@/components/layout/page-header';
import {
  type AutomationMode,
  type ScheduleTimeRange,
  type WeeklySchedule,
  type UpdateReceptionSettingsPayload,
  useReceptionSettings,
  useUpdateReceptionSettings,
} from '@/lib/api/conversations-api';

const DAYS: { key: string; label: string }[] = [
  { key: 'mon', label: 'Segunda' },
  { key: 'tue', label: 'Terça' },
  { key: 'wed', label: 'Quarta' },
  { key: 'thu', label: 'Quinta' },
  { key: 'fri', label: 'Sexta' },
  { key: 'sat', label: 'Sábado' },
  { key: 'sun', label: 'Domingo' },
];

const MODES: { value: AutomationMode; label: string; desc: string }[] = [
  { value: 'off', label: 'Manual', desc: 'A IA não age; a equipe responde tudo.' },
  { value: 'copilot', label: 'Copiloto', desc: 'A IA rascunha; a equipe revisa e envia.' },
  { value: 'auto', label: 'Auto', desc: 'A IA responde sozinha após o tempo de resposta.' },
];

// Exemplo do plano: noite (18:00→08:00) seg-sex; dia todo no fim de semana.
const EXAMPLE_SCHEDULE: WeeklySchedule = {
  mon: [{ start: '18:00', end: '08:00' }],
  tue: [{ start: '18:00', end: '08:00' }],
  wed: [{ start: '18:00', end: '08:00' }],
  thu: [{ start: '18:00', end: '08:00' }],
  fri: [{ start: '18:00', end: '08:00' }],
  sat: [{ start: '00:00', end: '00:00' }],
  sun: [{ start: '00:00', end: '00:00' }],
};

function cloneSchedule(s: WeeklySchedule): WeeklySchedule {
  const out: WeeklySchedule = {};
  for (const { key } of DAYS) out[key] = (s[key] ?? []).map((r) => ({ ...r }));
  return out;
}

export default function ConfigAtendimentoIAPage() {
  useRequireAuth();
  const { data, isLoading } = useReceptionSettings();
  const update = useUpdateReceptionSettings();

  const [baseline, setBaseline] = useState<AutomationMode>('off');
  const [scheduleEnabled, setScheduleEnabled] = useState(false);
  const [schedule, setSchedule] = useState<WeeklySchedule>({});
  const [debounceSeconds, setDebounceSeconds] = useState(30);
  const [copilotFallbackMinutes, setCopilotFallbackMinutes] = useState(10);
  const [offAlertMinutes, setOffAlertMinutes] = useState(20);
  const [loaded, setLoaded] = useState(false);

  // Hidrata o estado local quando a config chega (uma vez).
  useEffect(() => {
    if (!data || loaded) return;
    setBaseline(data.baselineMode);
    setScheduleEnabled(data.scheduleEnabled);
    setSchedule(cloneSchedule(data.schedule ?? {}));
    setDebounceSeconds(data.debounceSeconds);
    setCopilotFallbackMinutes(data.copilotFallbackMinutes);
    setOffAlertMinutes(data.offAlertMinutes);
    setLoaded(true);
  }, [data, loaded]);

  const addRange = (day: string) => {
    setSchedule((prev) => ({
      ...prev,
      [day]: [...(prev[day] ?? []), { start: '18:00', end: '08:00' }],
    }));
  };

  const removeRange = (day: string, idx: number) => {
    setSchedule((prev) => ({
      ...prev,
      [day]: (prev[day] ?? []).filter((_, i) => i !== idx),
    }));
  };

  const setRange = (day: string, idx: number, patch: Partial<ScheduleTimeRange>) => {
    setSchedule((prev) => ({
      ...prev,
      [day]: (prev[day] ?? []).map((r, i) => (i === idx ? { ...r, ...patch } : r)),
    }));
  };

  const setFullDay = (day: string) => {
    setSchedule((prev) => ({ ...prev, [day]: [{ start: '00:00', end: '00:00' }] }));
  };

  // Normaliza: remove dias vazios pra não poluir o JSON.
  const cleanedSchedule = useMemo(() => {
    const out: WeeklySchedule = {};
    for (const { key } of DAYS) {
      const ranges = (schedule[key] ?? []).filter((r) => r.start && r.end);
      if (ranges.length) out[key] = ranges;
    }
    return out;
  }, [schedule]);

  const handleSave = () => {
    // Blinda contra campo vazio/NaN (limpar o input vira NaN) batendo nas regras do backend.
    const clamp = (v: number, lo: number, hi: number, fallback: number) =>
      Number.isFinite(v) ? Math.min(hi, Math.max(lo, Math.round(v))) : fallback;
    const payload: UpdateReceptionSettingsPayload = {
      baselineMode: baseline,
      scheduleEnabled,
      schedule: cleanedSchedule,
      debounceSeconds: clamp(debounceSeconds, 0, 3600, 30),
      copilotFallbackMinutes: clamp(copilotFallbackMinutes, 1, 1440, 10),
      offAlertMinutes: clamp(offAlertMinutes, 1, 1440, 20),
    };
    update.mutate(payload, {
      onSuccess: () => toast.success('Atendimento IA salvo.'),
      onError: (e) =>
        toast.error(e instanceof Error ? e.message : 'Falha ao salvar a configuração'),
    });
  };

  return (
    <div className="container mx-auto max-w-4xl space-y-6 py-8">
      <PageHeader
        breadcrumbs={[{ label: 'Configurações' }, { label: 'Atendimento IA' }]}
        title="Atendimento IA"
        description="Recepcionista virtual: modo global, janela de horário do Auto e tempos de resposta."
      />

      {isLoading || !loaded ? (
        <div className="space-y-4">
          <Skeleton className="h-40 w-full" />
          <Skeleton className="h-64 w-full" />
        </div>
      ) : (
        <>
          {/* Modo baseline global */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Modo global</CardTitle>
              <CardDescription>
                Vale para toda conversa nova. Conversas individuais podem sobrepor por 24h.
                Quando a janela de horário abaixo está ativa, o modo sobe para Auto automaticamente.
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-2">
              {MODES.map((m) => (
                <button
                  key={m.value}
                  type="button"
                  onClick={() => setBaseline(m.value)}
                  className={cn(
                    'flex w-full items-start gap-3 rounded-lg border p-3 text-left transition-colors',
                    baseline === m.value
                      ? 'border-violet-400 bg-violet-50'
                      : 'border-border hover:bg-muted/50'
                  )}
                >
                  <span
                    className={cn(
                      'mt-0.5 flex h-4 w-4 shrink-0 items-center justify-center rounded-full border',
                      baseline === m.value ? 'border-violet-600' : 'border-muted-foreground'
                    )}
                  >
                    {baseline === m.value && <span className="h-2 w-2 rounded-full bg-violet-600" />}
                  </span>
                  <span>
                    <span className="block text-sm font-medium">{m.label}</span>
                    <span className="block text-xs text-muted-foreground">{m.desc}</span>
                  </span>
                </button>
              ))}
            </CardContent>
          </Card>

          {/* Janela de horário do Auto */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between gap-3">
                <div>
                  <CardTitle className="text-base">Janela de horário do Auto</CardTitle>
                  <CardDescription>
                    Dentro da janela, o modo vira Auto. Fora dela, vale o modo global acima.
                    Faixas que viram a noite (ex.: 18:00 às 08:00) são aceitas.
                  </CardDescription>
                </div>
                <Switch
                  checked={scheduleEnabled}
                  onCheckedChange={(v) => setScheduleEnabled(!!v)}
                  aria-label="Ativar janela de horário"
                />
              </div>
            </CardHeader>
            <CardContent className={cn('space-y-3', !scheduleEnabled && 'opacity-50')}>
              <div className="flex justify-end">
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  disabled={!scheduleEnabled}
                  onClick={() => setSchedule(cloneSchedule(EXAMPLE_SCHEDULE))}
                  title="18:00 às 08:00 de seg a sex, dia todo no fim de semana"
                >
                  Aplicar exemplo
                </Button>
              </div>
              <div className="space-y-2">
                {DAYS.map(({ key, label }) => {
                  const ranges = schedule[key] ?? [];
                  return (
                    <div
                      key={key}
                      className="flex flex-col gap-2 rounded-lg border border-border p-3 sm:flex-row sm:items-start"
                    >
                      <div className="flex w-24 shrink-0 items-center gap-1.5 text-sm font-medium">
                        <Clock className="h-3.5 w-3.5 text-muted-foreground" /> {label}
                      </div>
                      <div className="flex-1 space-y-2">
                        {ranges.length === 0 ? (
                          <p className="text-xs italic text-muted-foreground">Sem janela (segue o modo global)</p>
                        ) : (
                          ranges.map((r, i) => {
                            const allDay = r.start === '00:00' && r.end === '00:00';
                            return (
                              <div key={i} className="flex flex-wrap items-center gap-2">
                                <Input
                                  type="time"
                                  value={r.start}
                                  disabled={!scheduleEnabled}
                                  onChange={(e) => setRange(key, i, { start: e.target.value })}
                                  className="h-8 w-28"
                                  aria-label={`${label} início ${i + 1}`}
                                />
                                <span className="text-xs text-muted-foreground">até</span>
                                <Input
                                  type="time"
                                  value={r.end}
                                  disabled={!scheduleEnabled}
                                  onChange={(e) => setRange(key, i, { end: e.target.value })}
                                  className="h-8 w-28"
                                  aria-label={`${label} fim ${i + 1}`}
                                />
                                {allDay && (
                                  <span className="rounded bg-emerald-100 px-1.5 py-0.5 text-[10px] font-medium text-emerald-900">
                                    dia todo
                                  </span>
                                )}
                                <button
                                  type="button"
                                  disabled={!scheduleEnabled}
                                  onClick={() => removeRange(key, i)}
                                  className="rounded p-1 text-rose-600 hover:bg-rose-50 disabled:opacity-50"
                                  aria-label={`Remover faixa ${i + 1} de ${label}`}
                                >
                                  <Trash2 className="h-3.5 w-3.5" />
                                </button>
                              </div>
                            );
                          })
                        )}
                        <div className="flex gap-2">
                          <Button
                            type="button"
                            variant="ghost"
                            size="sm"
                            className="h-7 px-2 text-xs"
                            disabled={!scheduleEnabled}
                            onClick={() => addRange(key)}
                          >
                            <Plus className="mr-1 h-3 w-3" /> Faixa
                          </Button>
                          <Button
                            type="button"
                            variant="ghost"
                            size="sm"
                            className="h-7 px-2 text-xs"
                            disabled={!scheduleEnabled}
                            onClick={() => setFullDay(key)}
                          >
                            Dia todo
                          </Button>
                        </div>
                      </div>
                    </div>
                  );
                })}
              </div>
            </CardContent>
          </Card>

          {/* Tempos X/Y/Z */}
          <Card>
            <CardHeader>
              <CardTitle className="text-base">Tempos de resposta</CardTitle>
              <CardDescription>
                Controlam quando a IA rascunha, envia, escala e alerta.
              </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-4 sm:grid-cols-3">
              <div className="space-y-1.5">
                <Label htmlFor="x">Tempo de resposta (X)</Label>
                <div className="flex items-center gap-2">
                  <Input
                    id="x"
                    type="number"
                    min={0}
                    max={3600}
                    value={debounceSeconds}
                    onChange={(e) => setDebounceSeconds(Number(e.target.value))}
                    className="h-9"
                  />
                  <span className="text-xs text-muted-foreground">seg</span>
                </div>
                <p className="text-xs text-muted-foreground">
                  Espera após a última mensagem antes de rascunhar/enviar. Reseta a cada nova mensagem.
                </p>
              </div>
              <div className="space-y-1.5">
                <Label htmlFor="y">Fallback do Copiloto (Y)</Label>
                <div className="flex items-center gap-2">
                  <Input
                    id="y"
                    type="number"
                    min={1}
                    max={1440}
                    value={copilotFallbackMinutes}
                    onChange={(e) => setCopilotFallbackMinutes(Number(e.target.value))}
                    className="h-9"
                  />
                  <span className="text-xs text-muted-foreground">min</span>
                </div>
                <p className="text-xs text-muted-foreground">
                  No Copiloto, se ninguém responder em Y, a IA assume e envia.
                </p>
              </div>
              <div className="space-y-1.5">
                <Label htmlFor="z">Alerta no Manual (Z)</Label>
                <div className="flex items-center gap-2">
                  <Input
                    id="z"
                    type="number"
                    min={1}
                    max={1440}
                    value={offAlertMinutes}
                    onChange={(e) => setOffAlertMinutes(Number(e.target.value))}
                    className="h-9"
                  />
                  <span className="text-xs text-muted-foreground">min</span>
                </div>
                <p className="text-xs text-muted-foreground">
                  No Manual, se uma mensagem ficar Z sem resposta, o admin é alertado.
                </p>
              </div>
            </CardContent>
          </Card>

          <div className="flex justify-end gap-2">
            <Button onClick={handleSave} disabled={update.isPending}>
              {update.isPending ? (
                <Loader2 className="mr-1 h-4 w-4 animate-spin" />
              ) : (
                <Save className="mr-1 h-4 w-4" />
              )}
              Salvar
            </Button>
          </div>
        </>
      )}
    </div>
  );
}
