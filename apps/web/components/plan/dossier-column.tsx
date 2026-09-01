'use client';

/**
 * O prontuário compilado, visível.
 *
 * O dossiê já existia e ninguém via: era um JSON que só a API devolvia. Ele traz o que a máquina
 * deriva das três fontes do prontuário — réguas por exame com a escala aplicável AQUELE paciente,
 * achados classificados e ordenados por peso, condutas, prescrições e o último pedido de exames.
 *
 * É o congelado, não o vivo. A diferença é o ponto: o dossiê vivo muda quando o prontuário muda, e
 * escrever um documento contra chão que se move foi o que motivou congelar.
 */

import { useMemo, useState } from 'react';
import { AlertTriangle, ArrowDown, ArrowUp, Loader2, Minus, RefreshCw } from 'lucide-react';
import type { PlanDossier, PlanFinding, PlanRuler } from '@plenya/types';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { formatDate } from '@/lib/format-date';
import { cn } from '@/lib/utils';

/** Cor por nível do escore, na mesma rampa da régua (0 pior, 5 melhor). */
function corDoNivel(nivel?: number) {
  if (nivel == null) return 'bg-muted text-muted-foreground';
  if (nivel <= 1) return 'bg-red-100 text-red-900 dark:bg-red-950 dark:text-red-200';
  if (nivel <= 2) return 'bg-amber-100 text-amber-900 dark:bg-amber-950 dark:text-amber-200';
  if (nivel <= 3) return 'bg-sky-100 text-sky-900 dark:bg-sky-950 dark:text-sky-200';
  return 'bg-emerald-100 text-emerald-900 dark:bg-emerald-950 dark:text-emerald-200';
}

/**
 * A direção vale mais que a magnitude, e é onde o leitor mais erra — por isso ela aparece como
 * seta e não só como número.
 */
function Direcao({ trend }: { trend?: string }) {
  if (trend === 'worsening') return <ArrowUp className="h-3 w-3 text-red-600" aria-label="piorando" />;
  if (trend === 'improving') return <ArrowDown className="h-3 w-3 text-emerald-600" aria-label="melhorando" />;
  if (trend === 'stable') return <Minus className="h-3 w-3 text-muted-foreground" aria-label="estável" />;
  return null;
}

function LinhaAchado({ f }: { f: PlanFinding }) {
  const perdidos = f.pointsLost ?? 0;
  return (
    <li className="flex items-start gap-2 border-b py-1.5 last:border-0">
      <span
        className={cn(
          'mt-0.5 w-6 shrink-0 rounded px-1 text-center text-[10px] font-semibold leading-4',
          corDoNivel(f.level),
        )}
        title={f.level != null ? `nível ${f.level}` : 'sem nível'}
      >
        {f.level ?? '—'}
      </span>
      <span className="min-w-0 flex-1">
        <span className="flex items-center gap-1">
          <span className="truncate text-xs font-medium">{f.name}</span>
          <Direcao trend={f.trend} />
          {f.stale && (
            <span className="text-[10px] text-muted-foreground" title="medido há muito tempo">
              antigo
            </span>
          )}
        </span>
        <span className="text-[11px] text-muted-foreground">
          {f.text ?? f.value ?? '—'} {f.unit ?? ''}
          {f.date ? ` · ${formatDate(f.date)}` : ''}
          {perdidos > 0 ? ` · ${perdidos.toLocaleString('pt-BR')} pt` : ''}
        </span>
      </span>
    </li>
  );
}

function ListaVazia({ children }: { children: React.ReactNode }) {
  return <p className="px-1 py-6 text-center text-xs text-muted-foreground">{children}</p>;
}

export interface DossierColumnProps {
  dossier?: PlanDossier;
  frozenAt?: string;
  carregando?: boolean;
  /** Motivos pelos quais o prontuário andou desde o congelamento. Vazio = está em dia. */
  motivosDeEnvelhecimento?: string[];
  onRefresh?: () => void;
  refrescando?: boolean;
}

export function DossierColumn({
  dossier,
  frozenAt,
  carregando,
  motivosDeEnvelhecimento,
  onRefresh,
  refrescando,
}: DossierColumnProps) {
  const [busca, setBusca] = useState('');

  const reguas = useMemo(() => {
    const todas = Object.entries(dossier?.rulers ?? {}) as [string, PlanRuler][];
    const q = busca.trim().toLowerCase();
    const filtradas = q
      ? todas.filter(([code, r]) => (r.name ?? '').toLowerCase().includes(q) || code.toLowerCase().includes(q))
      : todas;
    return filtradas.sort((a, b) => (a[1].name ?? '').localeCompare(b[1].name ?? '', 'pt-BR'));
  }, [dossier?.rulers, busca]);

  const movendo = dossier?.moving ?? [];
  const bem = dossier?.strong ?? [];
  const condutas = dossier?.carePlan ?? [];
  const vitais = dossier?.vitals ?? [];

  if (carregando) {
    return (
      <div className="flex h-40 items-center justify-center">
        <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
      </div>
    );
  }
  if (!dossier) {
    return <ListaVazia>Sem dossiê congelado para este plano.</ListaVazia>;
  }

  const envelheceu = (motivosDeEnvelhecimento?.length ?? 0) > 0;

  return (
    <div className="flex h-full flex-col gap-2">
      <div className="flex items-baseline justify-between gap-2">
        <h2 className="text-sm font-semibold">Prontuário compilado</h2>
        {frozenAt && (
          <span className="text-[11px] text-muted-foreground" title="o dossiê não muda sozinho">
            de {formatDate(frozenAt)}
          </span>
        )}
      </div>

      {envelheceu && (
        <div className="rounded-md border border-amber-300 bg-amber-50 p-2 text-[11px] text-amber-900 dark:border-amber-800 dark:bg-amber-950 dark:text-amber-200">
          <span className="flex items-center gap-1 font-medium">
            <AlertTriangle className="h-3 w-3" />
            O prontuário andou desde este dossiê
          </span>
          <span className="mt-0.5 block">{motivosDeEnvelhecimento!.join(' · ')}</span>
          {onRefresh && (
            <Button
              size="sm"
              variant="outline"
              className="mt-2 h-7 w-full text-[11px]"
              onClick={onRefresh}
              disabled={refrescando}
            >
              {refrescando ? (
                <Loader2 className="mr-1 h-3 w-3 animate-spin" />
              ) : (
                <RefreshCw className="mr-1 h-3 w-3" />
              )}
              Atualizar e ver o que mudou
            </Button>
          )}
        </div>
      )}

      <Tabs defaultValue="movendo" className="flex min-h-0 flex-1 flex-col">
        <TabsList className="h-8 w-full justify-start gap-0.5 overflow-x-auto">
          <TabsTrigger value="movendo" className="h-6 px-2 text-[11px]">
            Se movendo {movendo.length > 0 && <span className="ml-1 opacity-60">{movendo.length}</span>}
          </TabsTrigger>
          <TabsTrigger value="bem" className="h-6 px-2 text-[11px]">
            Está bem {bem.length > 0 && <span className="ml-1 opacity-60">{bem.length}</span>}
          </TabsTrigger>
          <TabsTrigger value="reguas" className="h-6 px-2 text-[11px]">
            Réguas <span className="ml-1 opacity-60">{Object.keys(dossier.rulers ?? {}).length}</span>
          </TabsTrigger>
          <TabsTrigger value="condutas" className="h-6 px-2 text-[11px]">
            Condutas {condutas.length > 0 && <span className="ml-1 opacity-60">{condutas.length}</span>}
          </TabsTrigger>
        </TabsList>

        <ScrollArea className="mt-2 min-h-0 flex-1 pr-2">
          <TabsContent value="movendo" className="mt-0">
            <p className="pb-1 text-[11px] text-muted-foreground">
              Ordenado pelos pontos que o item tira do escore.
            </p>
            {movendo.length === 0 ? (
              <ListaVazia>Nada em movimento.</ListaVazia>
            ) : (
              <ul>{movendo.map((f, i) => <LinhaAchado key={`${f.code}-${i}`} f={f} />)}</ul>
            )}
          </TabsContent>

          <TabsContent value="bem" className="mt-0">
            <p className="pb-1 text-[11px] text-muted-foreground">
              Nível 4 e 5, ordenado pelo peso do item: em nível bom ninguém perde ponto, então peso é
              o único sinal.
            </p>
            {bem.length === 0 ? (
              <ListaVazia>Nada classificado como bom ainda.</ListaVazia>
            ) : (
              <ul>{bem.map((f, i) => <LinhaAchado key={`${f.code}-${i}`} f={f} />)}</ul>
            )}
          </TabsContent>

          <TabsContent value="reguas" className="mt-0">
            <Input
              value={busca}
              onChange={(e) => setBusca(e.target.value)}
              placeholder="Buscar exame"
              className="mb-2 h-7 text-xs"
            />
            {reguas.length === 0 ? (
              <ListaVazia>Nenhuma régua para esta busca.</ListaVazia>
            ) : (
              <ul>
                {reguas.map(([code, r]) => {
                  const ultimo = r.history?.[r.history.length - 1];
                  return (
                    <li key={code} className="flex items-baseline justify-between gap-2 border-b py-1.5 last:border-0">
                      <span className="min-w-0 flex-1 truncate text-xs">{r.name}</span>
                      <span className="shrink-0 text-[11px] text-muted-foreground">
                        {ultimo?.text ?? '—'} {r.unit ?? ''}
                      </span>
                    </li>
                  );
                })}
              </ul>
            )}
          </TabsContent>

          <TabsContent value="condutas" className="mt-0">
            {condutas.length === 0 ? (
              <ListaVazia>
                Nenhuma conduta registrada no plano de cuidado. Sem elas, a metade prescritiva do
                deck não tem de onde sair.
              </ListaVazia>
            ) : (
              <ul>
                {condutas.map((c) => (
                  <li key={c.id} className="border-b py-1.5 last:border-0">
                    <span className="flex items-center gap-1">
                      <Badge variant="outline" className="h-4 px-1 text-[10px]">
                        {c.letterCode}
                      </Badge>
                      {c.priority === 'high' && (
                        <span className="text-[10px] font-medium text-amber-700">prioridade</span>
                      )}
                    </span>
                    <span className="mt-0.5 block text-xs">{c.recommendation}</span>
                  </li>
                ))}
              </ul>
            )}
          </TabsContent>

          {vitais.length > 0 && (
            <p className="mt-3 border-t pt-2 text-[11px] text-muted-foreground">
              Última aferição: {vitais[0].systolicBp ?? '—'}/{vitais[0].diastolicBp ?? '—'} mmHg
              {vitais[0].weight ? ` · ${vitais[0].weight} kg` : ''}
              {vitais[0].measuredAt ? ` · ${formatDate(vitais[0].measuredAt)}` : ''}
            </p>
          )}
        </ScrollArea>
      </Tabs>
    </div>
  );
}
