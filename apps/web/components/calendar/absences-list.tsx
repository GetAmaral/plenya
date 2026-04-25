'use client';

/**
 * AbsencesList — lista de ausências (férias, congresso, feriado) + form pra
 * adicionar novas. Slot picker exclui slots dentro de ausência ativa.
 *
 * Layout: futuras em destaque + passadas colapsáveis.
 */
import { useMemo, useState } from 'react';
import { format, isAfter } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { Plus, Trash2, ChevronDown, ChevronRight } from 'lucide-react';
import { toast } from 'sonner';

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Card } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';

import {
  useCreateAbsence,
  useDeleteAbsence,
  useDoctorAbsences,
  type DoctorAbsence,
} from '@/lib/api/calendar-api';

interface AbsencesListProps {
  doctorId: string;
  readOnly?: boolean;
}

export function AbsencesList({ doctorId, readOnly = false }: AbsencesListProps) {
  const { data: absences, isLoading } = useDoctorAbsences(doctorId);
  const createMutation = useCreateAbsence(doctorId);
  const deleteMutation = useDeleteAbsence(doctorId);

  const [open, setOpen] = useState(false);
  const [showPast, setShowPast] = useState(false);

  const { future, past } = useMemo(() => {
    const now = new Date();
    const list = absences ?? [];
    return {
      future: list.filter((a) => isAfter(new Date(a.endAt), now)),
      past: list.filter((a) => !isAfter(new Date(a.endAt), now)),
    };
  }, [absences]);

  const handleDelete = async (a: DoctorAbsence) => {
    if (!confirm(`Remover ausência "${a.reason || 'sem motivo'}"?`)) return;
    try {
      await deleteMutation.mutateAsync(a.id);
      toast.success('Ausência removida');
    } catch (err) {
      toast.error('Erro ao remover', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  };

  if (isLoading) {
    return (
      <div className="space-y-2">
        <Skeleton className="h-12 w-full" />
        <Skeleton className="h-12 w-full" />
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <p className="text-sm text-muted-foreground">
          Períodos de indisponibilidade pontual (férias, congressos, feriados).
        </p>
        {!readOnly && (
          <Button size="sm" onClick={() => setOpen(true)}>
            <Plus className="mr-2 h-4 w-4" />
            Nova ausência
          </Button>
        )}
      </div>

      {/* Futuras */}
      <div className="space-y-2">
        <h4 className="text-sm font-semibold">Próximas ({future.length})</h4>
        {future.length === 0 ? (
          <p className="text-sm italic text-muted-foreground">Nenhuma ausência agendada.</p>
        ) : (
          future.map((a) => (
            <AbsenceRow
              key={a.id}
              absence={a}
              readOnly={readOnly}
              onDelete={() => handleDelete(a)}
            />
          ))
        )}
      </div>

      {/* Passadas (colapsáveis) */}
      {past.length > 0 && (
        <div className="space-y-2">
          <button
            type="button"
            onClick={() => setShowPast((s) => !s)}
            className="flex items-center gap-1 text-sm font-semibold text-muted-foreground hover:text-foreground"
          >
            {showPast ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
            Histórico ({past.length})
          </button>
          {showPast &&
            past.map((a) => (
              <AbsenceRow
                key={a.id}
                absence={a}
                readOnly={readOnly}
                onDelete={() => handleDelete(a)}
              />
            ))}
        </div>
      )}

      <NewAbsenceModal
        open={open}
        onClose={() => setOpen(false)}
        onSubmit={async (input) => {
          try {
            await createMutation.mutateAsync(input);
            toast.success('Ausência criada');
            setOpen(false);
          } catch (err) {
            toast.error('Erro ao criar', {
              description: err instanceof Error ? err.message : undefined,
            });
          }
        }}
      />
    </div>
  );
}

function AbsenceRow({
  absence,
  readOnly,
  onDelete,
}: {
  absence: DoctorAbsence;
  readOnly: boolean;
  onDelete: () => void;
}) {
  const start = new Date(absence.startAt);
  const end = new Date(absence.endAt);
  return (
    <Card className="flex items-center justify-between p-3">
      <div>
        <div className="text-sm font-medium">
          {format(start, "dd 'de' MMM yyyy, HH:mm", { locale: ptBR })}
          {' → '}
          {format(end, "dd 'de' MMM yyyy, HH:mm", { locale: ptBR })}
        </div>
        {absence.reason && (
          <div className="text-xs text-muted-foreground">{absence.reason}</div>
        )}
      </div>
      {!readOnly && (
        <Button
          size="sm"
          variant="ghost"
          onClick={onDelete}
          className="text-rose-600 hover:text-rose-700"
        >
          <Trash2 className="h-4 w-4" />
        </Button>
      )}
    </Card>
  );
}

// =====================================================
// NewAbsenceModal
// =====================================================

interface NewAbsenceModalProps {
  open: boolean;
  onClose: () => void;
  onSubmit: (input: { startAt: string; endAt: string; reason: string }) => void;
}

function NewAbsenceModal({ open, onClose, onSubmit }: NewAbsenceModalProps) {
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [reason, setReason] = useState('');

  const handleSave = () => {
    if (!startDate || !endDate) {
      toast.error('Informe datas de início e fim');
      return;
    }
    // Datetime-local sem TZ → converte como local time → toISOString (UTC)
    const startAt = new Date(startDate).toISOString();
    const endAt = new Date(endDate).toISOString();
    if (new Date(endAt) <= new Date(startAt)) {
      toast.error('Data final deve ser após início');
      return;
    }
    onSubmit({ startAt, endAt, reason });
    setStartDate('');
    setEndDate('');
    setReason('');
  };

  return (
    <Dialog open={open} onOpenChange={(v) => (!v ? onClose() : null)}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Nova ausência</DialogTitle>
          <DialogDescription>
            Os horários no intervalo serão removidos do slot picker.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4 py-2">
          <div className="space-y-1.5">
            <Label htmlFor="absence-start">Início</Label>
            <Input
              id="absence-start"
              type="datetime-local"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>
          <div className="space-y-1.5">
            <Label htmlFor="absence-end">Fim</Label>
            <Input
              id="absence-end"
              type="datetime-local"
              value={endDate}
              onChange={(e) => setEndDate(e.target.value)}
            />
          </div>
          <div className="space-y-1.5">
            <Label htmlFor="absence-reason">Motivo (opcional)</Label>
            <Input
              id="absence-reason"
              type="text"
              placeholder="Ex: Férias, Congresso SBEM"
              value={reason}
              onChange={(e) => setReason(e.target.value)}
            />
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
            Cancelar
          </Button>
          <Button onClick={handleSave}>Criar</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
