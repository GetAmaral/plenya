'use client';

/**
 * Card de Alergias do paciente (P2a). Display + gestão (adicionar, inativar,
 * "sem alergias conhecidas"). Usado na capa do prontuário e na síntese do
 * workspace de consulta. Escopado por patientId (path).
 */
import { useState } from 'react';
import { toast } from 'sonner';
import { AlertTriangle, Plus, ShieldCheck, X, Loader2 } from 'lucide-react';

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
import { cn } from '@/lib/utils';
import {
  useAllergies,
  useCreateAllergy,
  useDeleteAllergy,
  type AllergySeverity,
  type AllergySubstanceType,
} from '@/lib/api/clinical-skeleton';

const SEVERITY_META: Record<AllergySeverity, { label: string; chip: string }> = {
  anaphylaxis: { label: 'Anafilaxia', chip: 'border-red-300 bg-red-100 text-red-900' },
  severe: { label: 'Grave', chip: 'border-red-300 bg-red-50 text-red-800' },
  moderate: { label: 'Moderada', chip: 'border-orange-300 bg-orange-50 text-orange-800' },
  mild: { label: 'Leve', chip: 'border-yellow-300 bg-yellow-50 text-yellow-800' },
};

const TYPE_LABEL: Record<AllergySubstanceType, string> = {
  drug: 'Medicamento',
  food: 'Alimento',
  environmental: 'Ambiental',
  other: 'Outro',
};

const selectCls =
  'w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none focus:border-primary';

export function AllergiesCard({ patientId, className }: { patientId: string; className?: string }) {
  const { data: allergies = [], isLoading } = useAllergies(patientId);
  const createAllergy = useCreateAllergy(patientId);
  const deleteAllergy = useDeleteAllergy(patientId);

  const [open, setOpen] = useState(false);
  const [substance, setSubstance] = useState('');
  const [substanceType, setSubstanceType] = useState<AllergySubstanceType>('drug');
  const [severity, setSeverity] = useState<AllergySeverity>('moderate');
  const [reaction, setReaction] = useState('');

  const active = allergies.filter((a) => a.status === 'active');
  const nka = active.find((a) => a.noKnownAllergies);
  const real = active.filter((a) => !a.noKnownAllergies);
  const hasSevere = real.some((a) => a.severity === 'severe' || a.severity === 'anaphylaxis');

  async function handleAdd() {
    if (!substance.trim()) {
      toast.error('Informe a substância');
      return;
    }
    try {
      await createAllergy.mutateAsync({
        substance: substance.trim(),
        substanceType,
        severity,
        reaction: reaction.trim() || undefined,
      });
      toast.success('Alergia registrada');
      setOpen(false);
      setSubstance('');
      setReaction('');
      setSeverity('moderate');
      setSubstanceType('drug');
    } catch (err) {
      toast.error('Erro ao registrar alergia', {
        description: err instanceof Error ? err.message : undefined,
      });
    }
  }

  async function markNKA() {
    try {
      await createAllergy.mutateAsync({ noKnownAllergies: true });
      toast.success('Registrado: sem alergias conhecidas');
    } catch (err) {
      toast.error('Erro ao registrar', { description: err instanceof Error ? err.message : undefined });
    }
  }

  async function inactivate(id: string) {
    try {
      await deleteAllergy.mutateAsync(id);
    } catch (err) {
      toast.error('Erro ao remover', { description: err instanceof Error ? err.message : undefined });
    }
  }

  return (
    <Card className={cn(hasSevere && 'border-red-300', className)}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <AlertTriangle className={cn('h-4 w-4', hasSevere ? 'text-red-600' : 'text-amber-600')} />
          Alergias
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
        ) : nka && real.length === 0 ? (
          <p className="flex items-center gap-2 text-sm text-emerald-700">
            <ShieldCheck className="h-4 w-4" />
            Sem alergias conhecidas
          </p>
        ) : real.length === 0 ? (
          <div className="space-y-2">
            <p className="text-sm text-muted-foreground">Alergias não avaliadas.</p>
            <Button variant="ghost" size="sm" onClick={markNKA} disabled={createAllergy.isPending}>
              <ShieldCheck className="mr-1 h-3.5 w-3.5" />
              Marcar "sem alergias conhecidas"
            </Button>
          </div>
        ) : (
          <div className="flex flex-wrap gap-2">
            {real.map((a) => {
              const meta = SEVERITY_META[a.severity];
              return (
                <span
                  key={a.id}
                  className={cn('inline-flex items-center gap-1.5 rounded-md border px-2 py-1 text-xs', meta.chip)}
                  title={`${TYPE_LABEL[a.substanceType]}${a.reaction ? ` — ${a.reaction}` : ''} (${meta.label})`}
                >
                  <span className="font-medium">{a.substance}</span>
                  <span className="opacity-70">· {meta.label}</span>
                  <button
                    type="button"
                    onClick={() => inactivate(a.id)}
                    className="ml-0.5 rounded-full hover:bg-black/10"
                    aria-label={`Inativar alergia ${a.substance}`}
                  >
                    <X className="h-3 w-3" />
                  </button>
                </span>
              );
            })}
          </div>
        )}
      </CardContent>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Registrar alergia</DialogTitle>
          </DialogHeader>
          <div className="space-y-3 py-2">
            <div className="space-y-1">
              <Label htmlFor="alg-substance">Substância / agente</Label>
              <Input
                id="alg-substance"
                value={substance}
                onChange={(e) => setSubstance(e.target.value)}
                placeholder="Ex.: Dipirona, Penicilina, Frutos do mar"
                autoFocus
              />
            </div>
            <div className="grid grid-cols-2 gap-3">
              <div className="space-y-1">
                <Label>Tipo</Label>
                <select className={selectCls} value={substanceType} onChange={(e) => setSubstanceType(e.target.value as AllergySubstanceType)}>
                  <option value="drug">Medicamento</option>
                  <option value="food">Alimento</option>
                  <option value="environmental">Ambiental</option>
                  <option value="other">Outro</option>
                </select>
              </div>
              <div className="space-y-1">
                <Label>Gravidade</Label>
                <select className={selectCls} value={severity} onChange={(e) => setSeverity(e.target.value as AllergySeverity)}>
                  <option value="mild">Leve</option>
                  <option value="moderate">Moderada</option>
                  <option value="severe">Grave</option>
                  <option value="anaphylaxis">Anafilaxia</option>
                </select>
              </div>
            </div>
            <div className="space-y-1">
              <Label htmlFor="alg-reaction">Reação (opcional)</Label>
              <Input
                id="alg-reaction"
                value={reaction}
                onChange={(e) => setReaction(e.target.value)}
                placeholder="Ex.: urticária, angioedema, broncoespasmo"
              />
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancelar
            </Button>
            <Button onClick={handleAdd} disabled={createAllergy.isPending}>
              {createAllergy.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Registrar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}
