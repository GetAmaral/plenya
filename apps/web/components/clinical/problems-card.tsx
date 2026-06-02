'use client';

/**
 * Lista de problemas do paciente (P2b). Descrição texto-livre + CID-10 opcional
 * (autocomplete no catálogo curado). Resolver / remover. Usado na capa e no
 * workspace de consulta.
 */
import { useState } from 'react';
import { toast } from 'sonner';
import { ClipboardList, Plus, X, Loader2, CheckCircle2 } from 'lucide-react';

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
  useProblems,
  useCreateProblem,
  useUpdateProblem,
  useDeleteProblem,
  useCIDSearch,
  type CIDCode,
} from '@/lib/api/clinical-skeleton';

function CIDAutocomplete({
  selected,
  onSelect,
  onClear,
}: {
  selected: CIDCode | null;
  onSelect: (c: CIDCode) => void;
  onClear: () => void;
}) {
  const [q, setQ] = useState('');
  const { data: results = [], isFetching } = useCIDSearch(q, !selected);

  if (selected) {
    return (
      <div className="flex items-center justify-between rounded-md border bg-muted/40 px-3 py-2 text-sm">
        <span>
          <span className="font-mono font-medium">{selected.code}</span> — {selected.description}
        </span>
        <button type="button" onClick={onClear} aria-label="Remover CID">
          <X className="h-4 w-4 text-muted-foreground" />
        </button>
      </div>
    );
  }
  return (
    <div className="relative">
      <Input
        value={q}
        onChange={(e) => setQ(e.target.value)}
        placeholder="Buscar CID-10 (ex.: hipertensão, E11, diabetes)"
      />
      {q.trim().length >= 2 && (
        <div className="absolute z-50 mt-1 max-h-56 w-full overflow-y-auto rounded-md border bg-popover shadow-md">
          {isFetching ? (
            <div className="flex items-center justify-center p-3">
              <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
            </div>
          ) : results.length === 0 ? (
            <p className="p-3 text-xs text-muted-foreground">Nenhum CID no catálogo. Use a descrição livre.</p>
          ) : (
            results.map((c) => (
              <button
                key={c.code}
                type="button"
                onClick={() => { onSelect(c); setQ(''); }}
                className="flex w-full items-start gap-2 px-3 py-2 text-left text-sm hover:bg-accent"
              >
                <span className="font-mono font-medium">{c.code}</span>
                <span className="text-muted-foreground">{c.description}</span>
              </button>
            ))
          )}
        </div>
      )}
    </div>
  );
}

export function ProblemsCard({ patientId, className }: { patientId: string; className?: string }) {
  const { data: problems = [], isLoading } = useProblems(patientId);
  const createProblem = useCreateProblem(patientId);
  const updateProblem = useUpdateProblem(patientId);
  const deleteProblem = useDeleteProblem(patientId);

  const [open, setOpen] = useState(false);
  const [description, setDescription] = useState('');
  const [cid, setCid] = useState<CIDCode | null>(null);
  const [onsetDate, setOnsetDate] = useState('');

  async function handleAdd() {
    if (!description.trim()) {
      toast.error('Informe a descrição do problema');
      return;
    }
    try {
      await createProblem.mutateAsync({
        description: description.trim(),
        cidCode: cid?.code,
        onsetDate: onsetDate || undefined,
      });
      toast.success('Problema adicionado');
      setOpen(false);
      setDescription('');
      setCid(null);
      setOnsetDate('');
    } catch (err) {
      toast.error('Erro ao adicionar problema', { description: err instanceof Error ? err.message : undefined });
    }
  }

  async function resolve(id: string) {
    try {
      await updateProblem.mutateAsync({ id, payload: { status: 'resolved' } });
    } catch (err) {
      toast.error('Erro ao resolver', { description: err instanceof Error ? err.message : undefined });
    }
  }
  async function remove(id: string) {
    try {
      await deleteProblem.mutateAsync(id);
    } catch (err) {
      toast.error('Erro ao remover', { description: err instanceof Error ? err.message : undefined });
    }
  }

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <ClipboardList className="h-4 w-4 text-primary" />
          Problemas ativos
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
        ) : problems.length === 0 ? (
          <p className="text-sm text-muted-foreground">Nenhum problema ativo registrado.</p>
        ) : (
          <ul className="space-y-2">
            {problems.map((p) => (
              <li key={p.id} className="flex items-start justify-between gap-2 text-sm">
                <div className="min-w-0 flex-1">
                  <span className="font-medium">{p.description}</span>
                  {p.cidCode && (
                    <Badge variant="outline" className="ml-2 font-mono text-[10px]">
                      {p.cidCode}
                    </Badge>
                  )}
                  {p.onsetDate && (
                    <span className="ml-2 text-xs text-muted-foreground">desde {p.onsetDate}</span>
                  )}
                </div>
                <div className="flex shrink-0 items-center gap-1">
                  <button
                    type="button"
                    onClick={() => resolve(p.id)}
                    title="Marcar como resolvido"
                    className="rounded p-1 text-emerald-600 hover:bg-emerald-50"
                  >
                    <CheckCircle2 className="h-4 w-4" />
                  </button>
                  <button
                    type="button"
                    onClick={() => remove(p.id)}
                    title="Remover"
                    className="rounded p-1 text-muted-foreground hover:bg-accent"
                  >
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
            <DialogTitle>Adicionar problema</DialogTitle>
          </DialogHeader>
          <div className="space-y-3 py-2">
            <div className="space-y-1">
              <Label htmlFor="prob-desc">Descrição</Label>
              <Input
                id="prob-desc"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Ex.: Hipertensão arterial sistêmica"
                autoFocus
              />
            </div>
            <div className="space-y-1">
              <Label>CID-10 (opcional)</Label>
              <CIDAutocomplete
                selected={cid}
                onSelect={(c) => {
                  setCid(c);
                  if (!description.trim()) setDescription(c.description);
                }}
                onClear={() => setCid(null)}
              />
            </div>
            <div className="space-y-1">
              <Label htmlFor="prob-onset">Início (opcional)</Label>
              <Input id="prob-onset" type="date" value={onsetDate} onChange={(e) => setOnsetDate(e.target.value)} />
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setOpen(false)}>Cancelar</Button>
            <Button onClick={handleAdd} disabled={createProblem.isPending}>
              {createProblem.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              Adicionar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}
