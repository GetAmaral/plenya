'use client';

/**
 * WorkingHoursEditor — editor visual dos blocos de atendimento semanais.
 *
 * Estrutura:
 *   - 7 colunas (Dom..Sáb), cada uma com lista de blocos
 *   - Cada bloco mostra HH:MM - HH:MM (slotDuration min)
 *   - Click "+" abre modal para criar
 *   - Click bloco abre modal para editar
 *   - Botão "Copiar de segunda" replica blocos de Seg em Ter-Sex
 *
 * RBAC: backend handler verifica admin/manager/próprio doctor.
 */
import { useEffect, useMemo, useState } from 'react';
import { Plus, Trash2, Copy } from 'lucide-react';
import { toast } from 'sonner';

import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Card } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';

import {
  WEEKDAY_LABELS_SHORT,
  minutesToTimeString,
  timeStringToMinutes,
  useCreateWorkingHours,
  useDeleteWorkingHours,
  useUpdateWorkingHours,
  useWorkingHours,
  type WorkingHours,
} from '@/lib/api/calendar-api';

interface WorkingHoursEditorProps {
  doctorId: string;
  /** Quando true, desabilita mutações (caller é doctor olhando outro). */
  readOnly?: boolean;
}

interface BlockModalState {
  open: boolean;
  weekday: number;
  block: WorkingHours | null;
}

const DEFAULT_BLOCK = { startTime: '08:00', endTime: '12:00', slotDuration: 30 };

export function WorkingHoursEditor({ doctorId, readOnly = false }: WorkingHoursEditorProps) {
  const { data: blocks, isLoading } = useWorkingHours(doctorId);
  const createMutation = useCreateWorkingHours(doctorId);
  const updateMutation = useUpdateWorkingHours(doctorId);
  const deleteMutation = useDeleteWorkingHours(doctorId);

  const [modal, setModal] = useState<BlockModalState>({
    open: false,
    weekday: 1,
    block: null,
  });

  const blocksByDay = useMemo(() => {
    const out: Record<number, WorkingHours[]> = { 0: [], 1: [], 2: [], 3: [], 4: [], 5: [], 6: [] };
    (blocks ?? []).forEach((b) => {
      out[b.weekday]?.push(b);
    });
    return out;
  }, [blocks]);

  const handleAdd = (weekday: number) => {
    setModal({ open: true, weekday, block: null });
  };

  const handleEdit = (block: WorkingHours) => {
    setModal({ open: true, weekday: block.weekday, block });
  };

  const handleDelete = async (block: WorkingHours) => {
    if (!confirm('Remover esse bloco de horário?')) return;
    try {
      await deleteMutation.mutateAsync(block.id);
      toast.success('Bloco removido');
    } catch (err) {
      toast.error('Erro ao remover bloco', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  const handleCopyFromMonday = async () => {
    const mondayBlocks = blocksByDay[1] ?? [];
    if (mondayBlocks.length === 0) {
      toast.error('Nenhum bloco em Segunda para copiar');
      return;
    }
    if (!confirm(`Copiar ${mondayBlocks.length} bloco(s) de Segunda para Ter-Sex?`)) return;
    try {
      for (const wd of [2, 3, 4, 5]) {
        for (const b of mondayBlocks) {
          await createMutation.mutateAsync({
            weekday: wd,
            startMinute: b.startMinute,
            endMinute: b.endMinute,
            slotDuration: b.slotDuration,
          });
        }
      }
      toast.success('Blocos copiados de Segunda para Ter-Sex');
    } catch (err) {
      toast.error('Erro ao copiar', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 gap-3 md:grid-cols-2 lg:grid-cols-7">
        {Array.from({ length: 7 }).map((_, i) => (
          <Skeleton key={i} className="h-40 w-full" />
        ))}
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex flex-wrap items-center justify-between gap-2">
        <p className="text-sm text-muted-foreground">
          Configure os blocos semanais de atendimento. O slot picker usa esses blocos
          como base disponível.
        </p>
        {!readOnly && (
          <Button variant="outline" size="sm" onClick={handleCopyFromMonday}>
            <Copy className="mr-2 h-3.5 w-3.5" />
            Copiar de Segunda
          </Button>
        )}
      </div>

      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-7">
        {WEEKDAY_LABELS_SHORT.map((label, idx) => (
          <Card key={idx} className="flex min-h-[180px] flex-col p-3">
            <div className="mb-3 flex items-center justify-between">
              <h4 className="text-sm font-semibold">{label}</h4>
              {!readOnly && (
                <Button
                  size="sm"
                  variant="ghost"
                  className="h-6 w-6 p-0"
                  onClick={() => handleAdd(idx)}
                  title="Adicionar bloco"
                >
                  <Plus className="h-3.5 w-3.5" />
                </Button>
              )}
            </div>
            <div className="flex-1 space-y-1.5">
              {(blocksByDay[idx] ?? []).length === 0 ? (
                <p className="text-xs italic text-muted-foreground">Sem atendimento</p>
              ) : (
                blocksByDay[idx]!.map((b) => (
                  <div
                    key={b.id}
                    className="group flex items-center justify-between rounded border bg-muted/40 px-2 py-1.5 text-xs"
                  >
                    <button
                      type="button"
                      disabled={readOnly}
                      onClick={() => handleEdit(b)}
                      className="flex-1 text-left disabled:cursor-default"
                    >
                      <div className="font-medium">
                        {minutesToTimeString(b.startMinute)} - {minutesToTimeString(b.endMinute)}
                      </div>
                      <div className="text-muted-foreground">{b.slotDuration} min</div>
                    </button>
                    {!readOnly && (
                      <button
                        type="button"
                        onClick={() => handleDelete(b)}
                        className="ml-2 opacity-0 transition-opacity group-hover:opacity-100"
                        title="Remover"
                      >
                        <Trash2 className="h-3 w-3 text-rose-600" />
                      </button>
                    )}
                  </div>
                ))
              )}
            </div>
          </Card>
        ))}
      </div>

      <BlockModal
        modal={modal}
        onClose={() => setModal({ open: false, weekday: 1, block: null })}
        onSubmit={async (input) => {
          try {
            if (modal.block) {
              await updateMutation.mutateAsync({
                id: modal.block.id,
                payload: input,
              });
              toast.success('Bloco atualizado');
            } else {
              await createMutation.mutateAsync(input);
              toast.success('Bloco criado');
            }
            setModal({ open: false, weekday: 1, block: null });
          } catch (err) {
            toast.error('Erro ao salvar', {
              description: err instanceof Error ? err.message : undefined,
            });
          }
        }}
      />
    </div>
  );
}

// =====================================================
// BlockModal
// =====================================================

interface BlockModalProps {
  modal: BlockModalState;
  onClose: () => void;
  onSubmit: (input: {
    weekday: number;
    startMinute: number;
    endMinute: number;
    slotDuration: number;
  }) => void;
}

function BlockModal({ modal, onClose, onSubmit }: BlockModalProps) {
  const [startTime, setStartTime] = useState(DEFAULT_BLOCK.startTime);
  const [endTime, setEndTime] = useState(DEFAULT_BLOCK.endTime);
  const [slotDuration, setSlotDuration] = useState<number>(DEFAULT_BLOCK.slotDuration);

  // Reset quando modal abre/altera
  useEffect(() => {
    if (modal.open) {
      if (modal.block) {
        setStartTime(minutesToTimeString(modal.block.startMinute));
        setEndTime(minutesToTimeString(modal.block.endMinute));
        setSlotDuration(modal.block.slotDuration);
      } else {
        setStartTime(DEFAULT_BLOCK.startTime);
        setEndTime(DEFAULT_BLOCK.endTime);
        setSlotDuration(DEFAULT_BLOCK.slotDuration);
      }
    }
  }, [modal.open, modal.block?.id]);

  const handleSave = () => {
    const startMinute = timeStringToMinutes(startTime);
    const endMinute = timeStringToMinutes(endTime);
    if (endMinute <= startMinute) {
      toast.error('Hora final deve ser após hora inicial');
      return;
    }
    if (slotDuration <= 0 || slotDuration > endMinute - startMinute) {
      toast.error('Duração de slot inválida');
      return;
    }
    onSubmit({
      weekday: modal.weekday,
      startMinute,
      endMinute,
      slotDuration,
    });
  };

  return (
    <Dialog open={modal.open} onOpenChange={(v) => (!v ? onClose() : null)}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {modal.block ? 'Editar bloco' : 'Novo bloco'} — {WEEKDAY_LABELS_SHORT[modal.weekday]}
          </DialogTitle>
          <DialogDescription>
            Defina o intervalo de atendimento e a granularidade dos slots.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4 py-2">
          <div className="grid grid-cols-2 gap-3">
            <div className="space-y-1.5">
              <Label htmlFor="start">Início</Label>
              <Input
                id="start"
                type="time"
                value={startTime}
                onChange={(e) => setStartTime(e.target.value)}
              />
            </div>
            <div className="space-y-1.5">
              <Label htmlFor="end">Fim</Label>
              <Input
                id="end"
                type="time"
                value={endTime}
                onChange={(e) => setEndTime(e.target.value)}
              />
            </div>
          </div>
          <div className="space-y-1.5">
            <Label htmlFor="slot">Granularidade dos slots</Label>
            <select
              id="slot"
              value={slotDuration}
              onChange={(e) => setSlotDuration(Number(e.target.value))}
              className="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"
            >
              <option value={15}>15 minutos</option>
              <option value={30}>30 minutos</option>
              <option value={60}>60 minutos</option>
            </select>
            <p className="text-xs text-muted-foreground">
              Define de quanto em quanto tempo um novo horário começa. Consultas mais longas
              aparecem como slots maiores no agendamento, mas sempre alinhados a essa granularidade.
            </p>
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
            Cancelar
          </Button>
          <Button onClick={handleSave}>Salvar</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

