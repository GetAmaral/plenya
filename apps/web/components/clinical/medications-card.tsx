'use client';

/**
 * Medicações em uso do paciente (P2b) — reconciliação. Distinto de prescrição
 * emitida. Usado na capa e no workspace de consulta.
 */
import { useState } from 'react';
import { toast } from 'sonner';
import { Pill, Plus, X, Loader2, PauseCircle } from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import {
  useMedicationsInUse,
  useCreateMedication,
  useUpdateMedication,
  useDeleteMedication,
  type MedicationSource,
} from '@/lib/api/clinical-skeleton';

const SOURCE_LABEL: Record<MedicationSource, string> = {
  prescribed_here: 'Prescrito aqui',
  external: 'Externo',
  patient_reported: 'Relato do paciente',
};

const selectCls =
  'w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none focus:border-primary';

export function MedicationsCard({ patientId, className }: { patientId: string; className?: string }) {
  const { data: meds = [], isLoading } = useMedicationsInUse(patientId);
  const createMed = useCreateMedication(patientId);
  const updateMed = useUpdateMedication(patientId);
  const deleteMed = useDeleteMedication(patientId);

  const [open, setOpen] = useState(false);
  const [form, setForm] = useState({
    medicationName: '',
    activeIngredient: '',
    dosage: '',
    frequency: '',
    route: '',
    source: 'patient_reported' as MedicationSource,
  });

  function set<K extends keyof typeof form>(k: K, v: (typeof form)[K]) {
    setForm((p) => ({ ...p, [k]: v }));
  }

  async function handleAdd() {
    if (!form.medicationName.trim()) {
      toast.error('Informe o nome da medicação');
      return;
    }
    try {
      await createMed.mutateAsync({
        medicationName: form.medicationName.trim(),
        activeIngredient: form.activeIngredient.trim() || undefined,
        dosage: form.dosage.trim() || undefined,
        frequency: form.frequency.trim() || undefined,
        route: form.route.trim() || undefined,
        source: form.source,
      });
      toast.success('Medicação adicionada');
      setOpen(false);
      setForm({ medicationName: '', activeIngredient: '', dosage: '', frequency: '', route: '', source: 'patient_reported' });
    } catch (err) {
      toast.error('Erro ao adicionar medicação', { description: err instanceof Error ? err.message : undefined });
    }
  }

  async function suspend(id: string) {
    try {
      await updateMed.mutateAsync({ id, payload: { status: 'suspended' } });
    } catch (err) {
      toast.error('Erro ao suspender', { description: err instanceof Error ? err.message : undefined });
    }
  }
  async function remove(id: string) {
    try {
      await deleteMed.mutateAsync(id);
    } catch (err) {
      toast.error('Erro ao remover', { description: err instanceof Error ? err.message : undefined });
    }
  }

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <Pill className="h-4 w-4 text-primary" />
          Medicações em uso
        </CardTitle>
        <Button variant="outline" size="sm" onClick={() => setOpen(true)}>
          <Plus className="mr-1 h-3.5 w-3.5" />
          Adicionar
        </Button>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <div className="flex h-12 items-center justify-center">
            <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
          </div>
        ) : meds.length === 0 ? (
          <p className="text-sm text-muted-foreground">Nenhuma medicação em uso registrada.</p>
        ) : (
          <ul className="space-y-2">
            {meds.map((m) => (
              <li key={m.id} className="flex items-start justify-between gap-2 text-sm">
                <div className="min-w-0 flex-1">
                  <span className="font-medium">{m.medicationName}</span>
                  {(m.dosage || m.frequency) && (
                    <span className="ml-2 text-xs text-muted-foreground">
                      {[m.dosage, m.frequency].filter(Boolean).join(' · ')}
                    </span>
                  )}
                  <Badge variant="outline" className="ml-2 text-[10px]">{SOURCE_LABEL[m.source]}</Badge>
                </div>
                <div className="flex shrink-0 items-center gap-1">
                  <button type="button" onClick={() => suspend(m.id)} title="Suspender" className="rounded p-1 text-amber-600 hover:bg-amber-50">
                    <PauseCircle className="h-4 w-4" />
                  </button>
                  <button type="button" onClick={() => remove(m.id)} title="Remover" className="rounded p-1 text-muted-foreground hover:bg-accent">
                    <X className="h-4 w-4" />
                  </button>
                </div>
              </li>
            ))}
          </ul>
        )}
      </CardContent>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Adicionar medicação em uso</DialogTitle>
          </DialogHeader>
          <div className="space-y-3 py-2">
            <div className="space-y-1">
              <Label htmlFor="med-name">Medicação</Label>
              <Input id="med-name" value={form.medicationName} onChange={(e) => set('medicationName', e.target.value)} placeholder="Ex.: Losartana" autoFocus />
            </div>
            <div className="grid grid-cols-2 gap-3">
              <div className="space-y-1">
                <Label htmlFor="med-ai">Princípio ativo</Label>
                <Input id="med-ai" value={form.activeIngredient} onChange={(e) => set('activeIngredient', e.target.value)} placeholder="Losartana potássica" />
              </div>
              <div className="space-y-1">
                <Label htmlFor="med-dose">Dose</Label>
                <Input id="med-dose" value={form.dosage} onChange={(e) => set('dosage', e.target.value)} placeholder="50 mg" />
              </div>
            </div>
            <div className="grid grid-cols-2 gap-3">
              <div className="space-y-1">
                <Label htmlFor="med-freq">Frequência</Label>
                <Input id="med-freq" value={form.frequency} onChange={(e) => set('frequency', e.target.value)} placeholder="1x/dia" />
              </div>
              <div className="space-y-1">
                <Label>Origem</Label>
                <select className={selectCls} value={form.source} onChange={(e) => set('source', e.target.value as MedicationSource)}>
                  <option value="patient_reported">Relato do paciente</option>
                  <option value="external">Externo</option>
                  <option value="prescribed_here">Prescrito aqui</option>
                </select>
              </div>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setOpen(false)}>Cancelar</Button>
            <Button onClick={handleAdd} disabled={createMed.isPending}>
              {createMed.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Adicionar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}
