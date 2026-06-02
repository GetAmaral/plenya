'use client';

/**
 * Painel de sinais vitais da consulta (P2a). Mostra a última aferição e permite
 * registrar uma nova, ancorada ao appointment. Usado no workspace de consulta.
 */
import { useState } from 'react';
import { toast } from 'sonner';
import { HeartPulse, Plus, Loader2 } from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useVitals, useCreateVitals } from '@/lib/api/clinical-skeleton';

type FieldKey =
  | 'systolicBp'
  | 'diastolicBp'
  | 'heartRate'
  | 'respRate'
  | 'temperature'
  | 'spo2'
  | 'weight'
  | 'height'
  | 'waistCircumference';

const FIELDS: { key: FieldKey; label: string; step?: string; unit: string }[] = [
  { key: 'systolicBp', label: 'PA sistólica', unit: 'mmHg' },
  { key: 'diastolicBp', label: 'PA diastólica', unit: 'mmHg' },
  { key: 'heartRate', label: 'FC', unit: 'bpm' },
  { key: 'respRate', label: 'FR', unit: 'irpm' },
  { key: 'temperature', label: 'Temp.', step: '0.1', unit: '°C' },
  { key: 'spo2', label: 'SpO₂', unit: '%' },
  { key: 'weight', label: 'Peso', step: '0.1', unit: 'kg' },
  { key: 'height', label: 'Altura', step: '0.1', unit: 'cm' },
  { key: 'waistCircumference', label: 'Cintura', step: '0.1', unit: 'cm' },
];

export function VitalsPanel({ patientId, appointmentId }: { patientId: string; appointmentId: string }) {
  const { data: vitals = [] } = useVitals(patientId);
  const createVitals = useCreateVitals(patientId);
  const latest = vitals[0];

  const [open, setOpen] = useState(false);
  const [form, setForm] = useState<Record<FieldKey, string>>({
    systolicBp: '', diastolicBp: '', heartRate: '', respRate: '',
    temperature: '', spo2: '', weight: '', height: '', waistCircumference: '',
  });

  async function handleSave() {
    const payload: Record<string, number | string> = { appointmentId };
    let any = false;
    for (const f of FIELDS) {
      const raw = form[f.key];
      if (raw !== '') {
        const n = Number(raw);
        if (!Number.isNaN(n)) {
          payload[f.key] = n;
          any = true;
        }
      }
    }
    if (!any) {
      toast.error('Preencha ao menos um sinal vital');
      return;
    }
    try {
      await createVitals.mutateAsync(payload as never);
      toast.success('Sinais vitais registrados');
      setOpen(false);
      setForm({
        systolicBp: '', diastolicBp: '', heartRate: '', respRate: '',
        temperature: '', spo2: '', weight: '', height: '', waistCircumference: '',
      });
    } catch (err) {
      toast.error('Erro ao registrar vitais', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  }

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-sm">
          <HeartPulse className="h-4 w-4 text-rose-600" />
          Sinais vitais
        </CardTitle>
        <Button variant="outline" size="sm" onClick={() => setOpen(true)}>
          <Plus className="mr-1 h-3.5 w-3.5" />
          Aferir
        </Button>
      </CardHeader>
      <CardContent>
        {latest ? (
          <div className="grid grid-cols-2 gap-x-4 gap-y-1 text-sm">
            {(latest.systolicBp || latest.diastolicBp) && (
              <Vital label="PA" value={`${latest.systolicBp ?? '—'}/${latest.diastolicBp ?? '—'}`} unit="mmHg" />
            )}
            {latest.heartRate != null && <Vital label="FC" value={latest.heartRate} unit="bpm" />}
            {latest.temperature != null && <Vital label="Temp." value={latest.temperature} unit="°C" />}
            {latest.spo2 != null && <Vital label="SpO₂" value={latest.spo2} unit="%" />}
            {latest.weight != null && <Vital label="Peso" value={latest.weight} unit="kg" />}
            {latest.bmi != null && <Vital label="IMC" value={latest.bmi} unit="" />}
          </div>
        ) : (
          <p className="text-sm text-muted-foreground">Sem aferição nesta consulta.</p>
        )}
      </CardContent>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Aferir sinais vitais</DialogTitle>
          </DialogHeader>
          <div className="grid grid-cols-2 gap-3 py-2 sm:grid-cols-3">
            {FIELDS.map((f) => (
              <div key={f.key} className="space-y-1">
                <Label className="text-xs">{f.label} <span className="text-muted-foreground">({f.unit})</span></Label>
                <Input
                  type="number"
                  step={f.step ?? '1'}
                  inputMode="decimal"
                  value={form[f.key]}
                  onChange={(e) => setForm((p) => ({ ...p, [f.key]: e.target.value }))}
                />
              </div>
            ))}
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancelar
            </Button>
            <Button onClick={handleSave} disabled={createVitals.isPending}>
              {createVitals.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Salvar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}

function Vital({ label, value, unit }: { label: string; value: string | number; unit: string }) {
  return (
    <div className="flex items-baseline justify-between gap-2">
      <span className="text-xs text-muted-foreground">{label}</span>
      <span className="font-medium">
        {value}
        {unit && <span className="ml-0.5 text-xs text-muted-foreground">{unit}</span>}
      </span>
    </div>
  );
}
