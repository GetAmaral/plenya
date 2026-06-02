'use client';

/**
 * CarePlanCard — plano de cuidado AGIR (P3 frente 3). Agrupa recomendações por pilar
 * A·G·I·R (Atividade · Gestão · Integração · Ritmo), cada uma ancorável a um biomarcador.
 * Multidisciplinar. Usado na capa do paciente e no workspace de consulta.
 * O botão "Gerar relatório" emite o relatório longitudinal assinado no portal (frente 3b).
 */

import { useState } from 'react';
import { Target, Plus, Loader2, Trash2, Check, FileText, RotateCcw } from 'lucide-react';
import { toast } from 'sonner';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { cn } from '@/lib/utils';
import {
  useCarePlan,
  useCreateCarePlanItem,
  useUpdateCarePlanItem,
  useDeleteCarePlanItem,
  useGenerateCarePlanReport,
  type AgirLetter,
  type CarePlanPriority,
} from '@/lib/api/care-plan';

const selectCls =
  'w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none focus:border-primary';

const LETTERS: { code: AgirLetter; name: string; dot: string; chip: string }[] = [
  { code: 'A', name: 'Atividade', dot: 'bg-amber-500', chip: 'bg-amber-100 text-amber-900 border-amber-200' },
  { code: 'G', name: 'Gestão', dot: 'bg-teal-600', chip: 'bg-teal-100 text-teal-900 border-teal-200' },
  { code: 'I', name: 'Integração', dot: 'bg-sky-600', chip: 'bg-sky-100 text-sky-900 border-sky-200' },
  { code: 'R', name: 'Ritmo', dot: 'bg-emerald-600', chip: 'bg-emerald-100 text-emerald-900 border-emerald-200' },
];

const PRIORITY_META: Record<CarePlanPriority, { label: string; cls: string }> = {
  high: { label: 'Alta', cls: 'bg-red-100 text-red-900 border-red-200' },
  medium: { label: 'Média', cls: 'bg-amber-100 text-amber-900 border-amber-200' },
  low: { label: 'Baixa', cls: 'bg-slate-100 text-slate-700 border-slate-200' },
};

export function CarePlanCard({
  patientId,
  className,
}: {
  patientId: string;
  className?: string;
}) {
  const { data = [], isLoading } = useCarePlan(patientId);
  const createItem = useCreateCarePlanItem(patientId);
  const updateItem = useUpdateCarePlanItem(patientId);
  const deleteItem = useDeleteCarePlanItem(patientId);
  const genReport = useGenerateCarePlanReport(patientId);

  const [open, setOpen] = useState(false);
  const [letter, setLetter] = useState<AgirLetter>('A');
  const [recommendation, setRecommendation] = useState('');
  const [rationale, setRationale] = useState('');
  const [target, setTarget] = useState('');
  const [labTestCode, setLabTestCode] = useState('');
  const [priority, setPriority] = useState<CarePlanPriority>('medium');

  const reset = () => {
    setLetter('A');
    setRecommendation('');
    setRationale('');
    setTarget('');
    setLabTestCode('');
    setPriority('medium');
  };

  const handleAdd = async () => {
    if (!recommendation.trim()) {
      toast.error('Informe a recomendação');
      return;
    }
    try {
      await createItem.mutateAsync({
        letterCode: letter,
        recommendation: recommendation.trim(),
        rationale: rationale.trim() || undefined,
        target: target.trim() || undefined,
        labTestCode: labTestCode.trim() || undefined,
        priority,
      });
      toast.success('Recomendação adicionada');
      setOpen(false);
      reset();
    } catch (err) {
      toast.error('Falha ao adicionar', { description: err instanceof Error ? err.message : undefined });
    }
  };

  const handleReport = async () => {
    try {
      await genReport.mutateAsync();
      toast.success('Relatório gerado e publicado no portal', {
        description: 'Disponível em Documentos do paciente',
      });
    } catch (err) {
      toast.error('Falha ao gerar relatório', { description: err instanceof Error ? err.message : undefined });
    }
  };

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <Target className="h-4 w-4 text-primary" />
          Plano de cuidado AGIR
        </CardTitle>
        <div className="flex items-center gap-2">
          <Button variant="ghost" size="sm" onClick={handleReport} disabled={genReport.isPending || data.length === 0} title="Gerar relatório longitudinal assinado no portal">
            {genReport.isPending ? <Loader2 className="mr-1 h-4 w-4 animate-spin" /> : <FileText className="mr-1 h-4 w-4" />}
            Relatório
          </Button>
          <Button variant="outline" size="sm" onClick={() => setOpen(true)}>
            <Plus className="mr-1 h-4 w-4" />
            Recomendação
          </Button>
        </div>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
        ) : data.length === 0 ? (
          <p className="text-sm text-muted-foreground">
            Nenhuma recomendação. Monte o plano por pilar (Atividade · Gestão · Integração · Ritmo).
          </p>
        ) : (
          <div className="space-y-4">
            {LETTERS.map((L) => {
              const items = data.filter((i) => i.letterCode === L.code);
              if (items.length === 0) return null;
              return (
                <div key={L.code}>
                  <div className="mb-2 flex items-center gap-2">
                    <span className={cn('flex h-5 w-5 items-center justify-center rounded-full text-[11px] font-bold text-white', L.dot)}>
                      {L.code}
                    </span>
                    <span className="text-sm font-semibold">{L.name}</span>
                  </div>
                  <ul className="space-y-2">
                    {items.map((item) => (
                      <li key={item.id} className="rounded-md border p-2">
                        <div className="flex items-start gap-2">
                          <div className="min-w-0 flex-1">
                            <p className="text-sm">{item.recommendation}</p>
                            <div className="mt-1 flex flex-wrap items-center gap-2">
                              <Badge variant="secondary" className={cn('text-[10px]', PRIORITY_META[item.priority].cls)}>
                                {PRIORITY_META[item.priority].label}
                              </Badge>
                              {item.labTestCode && (
                                <Badge variant="outline" className="text-[10px]">{item.labTestCode}</Badge>
                              )}
                              {item.target && <span className="text-xs text-muted-foreground">Meta: {item.target}</span>}
                              {item.authorName && <span className="text-xs text-muted-foreground">· {item.authorName}</span>}
                            </div>
                            {item.rationale && <p className="mt-1 text-xs text-muted-foreground">{item.rationale}</p>}
                          </div>
                          <div className="flex shrink-0 items-center gap-1">
                            <Button
                              variant="ghost"
                              size="icon"
                              className="h-7 w-7"
                              title="Marcar como atingida"
                              onClick={() =>
                                updateItem.mutate({ id: item.id, payload: { status: 'achieved' } })
                              }
                            >
                              <Check className="h-4 w-4" />
                            </Button>
                            <Button
                              variant="ghost"
                              size="icon"
                              className="h-7 w-7"
                              title="Remover"
                              onClick={() => deleteItem.mutate(item.id)}
                            >
                              <Trash2 className="h-4 w-4" />
                            </Button>
                          </div>
                        </div>
                      </li>
                    ))}
                  </ul>
                </div>
              );
            })}
          </div>
        )}
      </CardContent>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="max-h-[90vh] overflow-y-auto sm:max-w-lg">
          <DialogHeader>
            <DialogTitle>Nova recomendação</DialogTitle>
          </DialogHeader>
          <div className="space-y-3">
            <div>
              <Label htmlFor="cp-letter">Pilar AGIR</Label>
              <select id="cp-letter" className={selectCls} value={letter} onChange={(e) => setLetter(e.target.value as AgirLetter)}>
                {LETTERS.map((L) => (
                  <option key={L.code} value={L.code}>
                    {L.code} — {L.name}
                  </option>
                ))}
              </select>
            </div>
            <div>
              <Label htmlFor="cp-rec">Recomendação</Label>
              <Textarea id="cp-rec" rows={3} value={recommendation} onChange={(e) => setRecommendation(e.target.value)} />
            </div>
            <div>
              <Label htmlFor="cp-prio">Prioridade</Label>
              <select id="cp-prio" className={selectCls} value={priority} onChange={(e) => setPriority(e.target.value as CarePlanPriority)}>
                <option value="high">Alta</option>
                <option value="medium">Média</option>
                <option value="low">Baixa</option>
              </select>
            </div>
            <div>
              <Label htmlFor="cp-target">Meta (opcional)</Label>
              <Input id="cp-target" value={target} onChange={(e) => setTarget(e.target.value)} placeholder="Ex.: LDL < 70 mg/dL" />
            </div>
            <div>
              <Label htmlFor="cp-lab">Biomarcador / código (opcional)</Label>
              <Input id="cp-lab" value={labTestCode} onChange={(e) => setLabTestCode(e.target.value)} placeholder="Ex.: LDL" />
            </div>
            <div>
              <Label htmlFor="cp-rationale">Justificativa (opcional)</Label>
              <Textarea id="cp-rationale" rows={2} value={rationale} onChange={(e) => setRationale(e.target.value)} />
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancelar
            </Button>
            <Button onClick={handleAdd} disabled={createItem.isPending}>
              {createItem.isPending && <Loader2 className="mr-1 h-4 w-4 animate-spin" />}
              Adicionar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}
