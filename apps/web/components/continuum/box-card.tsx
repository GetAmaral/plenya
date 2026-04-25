'use client';

/**
 * BoxCard — card individual de PatientContinuumBox com fluxo de status,
 * tracking, transportadora e notas. Usado pela página /continuum/boxes
 * (cross-paciente) e pode ser embutido na timeline do paciente.
 */
import { useState } from 'react';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { Loader2, Package, Truck, CheckCircle2 } from 'lucide-react';
import { toast } from 'sonner';

import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  BOX_STATUS_LABELS,
  BOX_STATUS_COLORS,
  useUpdateContinuumBox,
  type ContinuumBox,
  type BoxStatus,
} from '@/lib/api/continuum-api';
import { cn } from '@/lib/utils';

interface Props {
  box: ContinuumBox;
  canManage: boolean;
  showPatient?: boolean; // hide quando já está no contexto do paciente
}

const CARRIERS = ['Correios', 'Loggi', 'Jadlog', 'Sequoia', 'Outro'];

export function BoxCard({ box, canManage, showPatient = true }: Props) {
  const update = useUpdateContinuumBox();
  const [tracking, setTracking] = useState(box.trackingCode ?? '');
  const [carrier, setCarrier] = useState(box.carrier ?? '');
  const [editingTracking, setEditingTracking] = useState(false);

  const setStatus = async (status: BoxStatus) => {
    try {
      await update.mutateAsync({ id: box.id, payload: { status } });
      toast.success(`Box marcado como ${BOX_STATUS_LABELS[status].toLowerCase()}`);
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao atualizar');
    }
  };

  const saveTracking = async () => {
    try {
      await update.mutateAsync({
        id: box.id,
        payload: { trackingCode: tracking, carrier },
      });
      toast.success('Tracking salvo');
      setEditingTracking(false);
    } catch (e: any) {
      toast.error(e?.response?.data?.error ?? 'Falha ao salvar');
    }
  };

  return (
    <Card className={cn(box.status === 'delivered' && 'bg-emerald-50/30')}>
      <CardContent className="space-y-3 p-4">
        <div className="flex flex-wrap items-start justify-between gap-2">
          <div className="space-y-1">
            <div className="flex items-center gap-2">
              <Package className="h-4 w-4 text-muted-foreground" />
              <p className="text-sm font-semibold">{box.name}</p>
              <Badge variant="outline" className={cn('text-xs', BOX_STATUS_COLORS[box.status])}>
                {BOX_STATUS_LABELS[box.status]}
              </Badge>
            </div>
            <p className="text-xs text-muted-foreground">
              {showPatient && (
                <>
                  {box.patientName} ·{' '}
                </>
              )}
              Semana {box.weekOffset + 1} ·{' '}
              {format(new Date(box.expectedDate), 'dd/MM/yyyy', { locale: ptBR })}
            </p>
          </div>
          {canManage && (
            <Select
              value={box.status}
              onValueChange={(v) => setStatus(v as BoxStatus)}
            >
              <SelectTrigger className="h-8 w-36 text-xs">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {Object.entries(BOX_STATUS_LABELS).map(([k, v]) => (
                  <SelectItem key={k} value={k} className="text-xs">
                    {v}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          )}
        </div>

        {box.contents && (
          <details className="text-xs text-muted-foreground">
            <summary className="cursor-pointer font-medium text-foreground hover:underline">
              Conteúdo
            </summary>
            <pre className="mt-1 whitespace-pre-wrap font-sans text-xs">{box.contents}</pre>
          </details>
        )}

        {/* Tracking */}
        <div className="rounded-md border bg-muted/30 p-2">
          <div className="flex items-center justify-between gap-2">
            <div className="flex items-center gap-2 text-xs">
              <Truck className="h-3.5 w-3.5 text-muted-foreground" />
              {!editingTracking ? (
                <span>
                  {box.trackingCode ? (
                    <>
                      <span className="font-mono">{box.trackingCode}</span>
                      {box.carrier && (
                        <span className="ml-2 text-muted-foreground">· {box.carrier}</span>
                      )}
                    </>
                  ) : (
                    <span className="text-muted-foreground">Sem código de rastreio</span>
                  )}
                </span>
              ) : (
                <div className="flex flex-wrap items-center gap-1">
                  <Input
                    className="h-7 w-40 text-xs"
                    placeholder="BR123..."
                    value={tracking}
                    onChange={(e) => setTracking(e.target.value)}
                  />
                  <Select value={carrier} onValueChange={setCarrier}>
                    <SelectTrigger className="h-7 w-28 text-xs">
                      <SelectValue placeholder="Trans." />
                    </SelectTrigger>
                    <SelectContent>
                      {CARRIERS.map((c) => (
                        <SelectItem key={c} value={c} className="text-xs">
                          {c}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              )}
            </div>
            {canManage && (
              <Button
                size="sm"
                variant="ghost"
                className="h-7 text-xs"
                onClick={() => {
                  if (editingTracking) {
                    saveTracking();
                  } else {
                    setEditingTracking(true);
                  }
                }}
                disabled={update.isPending}
              >
                {update.isPending && editingTracking ? (
                  <Loader2 className="h-3 w-3 animate-spin" />
                ) : editingTracking ? (
                  'Salvar'
                ) : (
                  'Editar'
                )}
              </Button>
            )}
          </div>
        </div>

        {/* Timeline mini */}
        {(box.preparedAt || box.shippedAt || box.deliveredAt) && (
          <div className="flex flex-wrap gap-3 text-xs text-muted-foreground">
            {box.preparedAt && (
              <span>
                Preparado{' '}
                {format(new Date(box.preparedAt), 'dd/MM', { locale: ptBR })}
              </span>
            )}
            {box.shippedAt && (
              <span>
                Enviado{' '}
                {format(new Date(box.shippedAt), 'dd/MM', { locale: ptBR })}
              </span>
            )}
            {box.deliveredAt && (
              <span className="font-medium text-emerald-700">
                <CheckCircle2 className="mr-1 inline h-3 w-3" />
                Entregue{' '}
                {format(new Date(box.deliveredAt), 'dd/MM', { locale: ptBR })}
              </span>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  );
}
