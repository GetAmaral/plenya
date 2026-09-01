'use client';

/**
 * O que este paciente já fez, na hora de montar o pedido.
 *
 * Existe para impedir um erro concreto: o painel do catálogo quase nunca tem resultado próprio (o
 * laboratório reporta os analitos), então quem confere olhando só o painel conclui "nunca fez" e
 * pede de novo o exame que o paciente acabou de fazer. A conta certa vem pronta do servidor, que
 * resolve o painel pelos filhos.
 */
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { AlertTriangle, ChevronDown, ChevronRight, FlaskConical } from 'lucide-react';

import { apiClient } from '@/lib/api-client';
import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';

type Via = 'never' | 'own' | 'children';

interface CoverageEntry {
  code: string;
  name: string;
  lastDoneAt?: string;
  daysAgo?: number;
  via: Via;
  childrenDone?: number;
  childrenDoneEver?: number;
  childrenTotal?: number;
}

/** Janela padrão: um ano. É o horizonte em que repetir um exame costuma ser desnecessário. */
const JANELA_DIAS = 365;

export function LabCoverageCard({ patientId }: { patientId: string }) {
  const [aberto, setAberto] = useState(false);

  // doneOnly corta ~80% do corpo: o catálogo inteiro são 528 entradas e este cartão só usa as feitas.
  const { data, isLoading, isError } = useQuery({
    queryKey: ['lab-coverage', patientId],
    enabled: !!patientId,
    queryFn: () =>
      apiClient.get<{ entries: CoverageEntry[] }>(
        `/api/v1/patients/${patientId}/lab-coverage?doneOnly=true`,
      ),
  });

  const feitos = (data?.entries ?? [])
    .filter((e) => e.via !== 'never' && (e.daysAgo ?? Infinity) <= JANELA_DIAS)
    .sort((a, b) => (a.daysAgo ?? 0) - (b.daysAgo ?? 0));

  if (isLoading) return null;

  // Falha de rede ou 403 (o endpoint é de clínico; secretaria abre a tela mas não a cobertura) NÃO
  // pode sair igual a "não fez nada": some do jeito errado, e o médico volta a pedir exame repetido
  // achando que a conferência disse que não havia nada. Diz o que houve.
  if (isError) {
    return (
      <Card className="mb-6 border-amber-300 bg-amber-50 p-4 text-sm text-amber-900">
        <span className="flex items-center gap-2">
          <AlertTriangle className="h-4 w-4" />
          Não foi possível carregar o que o paciente já fez. Confira o histórico de exames antes de
          pedir de novo.
        </span>
      </Card>
    );
  }

  if (feitos.length === 0) return null;

  return (
    <Card className="mb-6 p-4">
      <button
        className="flex w-full items-center gap-2 text-left"
        onClick={() => setAberto((v) => !v)}
      >
        {aberto ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
        <FlaskConical className="h-4 w-4 text-muted-foreground" />
        <span className="font-medium">Já feito no último ano</span>
        <Badge variant="secondary">{feitos.length}</Badge>
        <span className="ml-auto text-xs text-muted-foreground">
          confira antes de repetir
        </span>
      </button>

      {aberto && (
        <div className="mt-3 max-h-72 space-y-1 overflow-y-auto">
          {feitos.map((e) => (
            <div
              key={e.code}
              className="flex items-baseline justify-between gap-3 border-b py-1.5 text-sm last:border-0"
            >
              {/* min-w-0 no item do flex + truncate num bloco: truncate num <span> inline não
                  recorta (overflow/text-overflow não valem em caixa inline), e o catálogo tem nome
                  de 82 caracteres que invadiria a coluna da direita. */}
              <span className="min-w-0 flex-1">
                <span className="block truncate">{e.name}</span>
                {/* A proporção é a DESTA coleta, não a de sempre: "Rotina de urina" tem 14 analitos
                    já vistos mas só 8 vieram na coleta mais recente, e mostrar 14 faria o médico
                    deixar de pedir os outros 6. */}
                {e.via === 'children' && e.childrenTotal ? (
                  <span className="text-xs text-muted-foreground">
                    {e.childrenDone}/{e.childrenTotal} analitos nesta coleta
                    {e.childrenDoneEver && e.childrenDoneEver > (e.childrenDone ?? 0)
                      ? ` · ${e.childrenDoneEver} já vistos antes`
                      : ''}
                  </span>
                ) : null}
              </span>
              <span
                className={cn(
                  'shrink-0 tabular-nums',
                  (e.daysAgo ?? 0) <= 90 ? 'font-medium text-emerald-700' : 'text-muted-foreground',
                )}
              >
                há {e.daysAgo} dias
              </span>
            </div>
          ))}
        </div>
      )}
    </Card>
  );
}
